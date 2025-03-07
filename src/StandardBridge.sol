// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {BridgeRegistry} from "./abstract/BridgeRegistry.sol";
import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {IBridgeFeeStation} from "./interface/IBridgeFeeStation.sol";
import {IStandardBridge} from "./interface/IStandardBridge.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title StandardBridge
 * @notice Core implementation of cross-chain bridge functionality
 * @dev This contract provides the following key features:
 * - Token bridging (bridgeToken)
 * - Permit-based token bridging (permitBridgeToken)
 * - Bridge finalization (finalizeBridge)
 * - Batch processing capabilities
 * - Fee management
 * - Pending operation handling
 * - Safety mechanisms for high-value transfers
 *
 * @dev Key state variables:
 * - bridgeFeeStation: Contract managing bridge fees
 * - _nexus: Reward wallet address for collecting fees
 * - _initializedAt: Block number when contract was initialized
 *
 * @dev Bridge process flow:
 * 1. User calls bridgeToken/permitBridgeToken
 * 2. Tokens are deposited or burned in contract
 * 3. BridgeInitiated event is emitted
 * 4. Validators provide signatures
 * 5. finalizeBridge mints/transfers tokens on target chain
 */
contract StandardBridge is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    BridgeRegistry,
    ValidatorManager,
    IStandardBridge
{
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    error StandardInvalidValueUnit(IERC20 token, uint value);
    error StandardInvalidValue(uint expected, uint actual);
    error StandardInvalidIndex(uint expected, uint actual);
    error StandardInvalidMsgValue(uint expected, uint actual);
    error StandardInvalidBalance(uint expected, uint actual);
    error StandardInvalidAmount();
    error StandardInvalidPermitToken(address expected, address actual);
    error StandardCanNotZeroMsgValue();
    error StandardCanNotZeroAddress();
    error StandardNotExistIndex(uint index);
    error StandardNotExistToken(address token);
    error StandardNotExist();
    error StandardNotExpiredSafeDeadline(uint safeDeadline, uint timestamp);
    error StandardBurnFailed(IERC20 token, address from, uint value);
    error StandardNotMatchLength();
    error StandardFailedCall();

    event BridgeInitiated(
        uint indexed remoteChainID,
        uint indexed index,
        IERC20 indexed localToken,
        IERC20 remoteToken,
        address from,
        address to,
        uint value,
        uint time,
        bool permit,
        bytes extraData
    );
    event BridgeFinalized(
        uint indexed remoteChainID, uint indexed index, IERC20 token, address indexed to, uint value, uint time
    );
    event BridgeFinalizePending(uint indexed index);
    event BridgeFeeCharged(uint indexed index, IERC20 indexed token, address indexed account, uint gasFee, uint exFee);
    event PendingLocked(uint indexed remoteChainID, uint indexed index);
    event NexusSet(address indexed wallet);
    event FeeStationSet(IBridgeFeeStation indexed feeStation);

    /// @dev Native token representation address
    address internal constant NATIVE_TOKEN = address(1);

    /// @dev Hash for finalize operation signature verification
    bytes32 private constant FINALIZE_TYPEHASH = keccak256(
        "Finalize(uint256 remoteChainID,uint256 index,address token,address to,uint256 value,bytes extraData)"
    );

    /// @dev Fee station contract for managing bridge fees
    IBridgeFeeStation public bridgeFeeStation;

    /// @dev Reward wallet for collecting fees
    address payable private _nexus;

    /// @dev Block number when contract was initialized
    uint private _initializedAt;

    /// @dev Storage gap for future upgrades
    uint[47] private __gap;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Allows contract to receive native tokens
     * @dev Reverts if sent value is zero
     */
    receive() external payable {
        require(msg.value != 0, StandardCanNotZeroMsgValue());
    }

    /**
     * @notice Initializes the bridge contract
     * @dev One-time initializer for the bridge contract
     * - Sets up initial state including validator threshold and reward wallet
     * - Calls internal initialization function
     * @param _threshold Required number of validator signatures
     * @param nexus_ Address of reward wallet
     */
    function initialize(uint8 _threshold, address nexus_) external virtual initializer {
        __StandardBridge_init(_threshold, nexus_);
    }

    /**
     * @notice Initializes the bridge contract
     * @dev Sets up initial state including validator threshold and reward wallet
     * @param _threshold Required number of validator signatures
     * @param nexus_ Address of reward wallet
     */
    function __StandardBridge_init(uint8 _threshold, address nexus_) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __Validator_init(_threshold);
        __BridgeRegistry_init();

        require(address(nexus_) != address(0), StandardCanNotZeroAddress());
        _nexus = payable(nexus_);

        _initializedAt = block.number;
    }

    /**
     * @notice Bridges a token to a remote chain
     * @dev Core function for initiating a bridge transfer
     * - Validates token and amounts
     * - Collects fees
     * - Handles token deposit/burn
     * - Emits events for tracking
     * @param remoteChainID Target chain identifier
     * @param token Token address to bridge
     * @param to Recipient address on target chain
     * @param value Amount of tokens to bridge
     * @param gasFee Gas fee for transaction on target chain
     * @param exFee Exchange fee for the bridge service
     * @param extraData Additional data for the bridge operation
     * @return success Boolean indicating successful initiation
     */
    function bridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData
    ) public payable whenNotPaused onlyValidToken(remoteChainID, address(token)) nonReentrant returns (bool) {
        (gasFee, exFee) = _checkAmount(remoteChainID, token, value, gasFee, exFee);
        _executeBridge(remoteChainID, token, _msgSender(), to, value, gasFee, exFee, false, extraData);
        return true;
    }

    /**
     * @notice Bridges tokens using permit functionality
     * @dev Allows bridging with permit-based approval
     * - Validates permit signature
     * - Processes bridge operation
     * - Handles token transfer and fees
     * @param remoteChainID Target chain identifier
     * @param token Token address to bridge
     * @param to Recipient address on target chain
     * @param value Amount of tokens to bridge
     * @param gasFee Gas fee for target chain
     * @param exFee Exchange fee
     * @param extraData Additional bridge data
     * @param permitArgs Permit signature and parameters
     * @return success Boolean indicating successful initiation
     */
    function permitBridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData,
        PermitArguments calldata permitArgs
    ) public payable whenNotPaused onlyValidToken(remoteChainID, address(token)) nonReentrant returns (bool) {
        require(
            address(token) == address(permitArgs.token),
            StandardInvalidPermitToken(address(token), address(permitArgs.token))
        );
        (gasFee, exFee) = _checkAmount(remoteChainID, token, value, gasFee, exFee);

        require(
            permitArgs.value >= value + gasFee + exFee, StandardInvalidValue(value + gasFee + exFee, permitArgs.value)
        );

        {
            bytes memory data = abi.encodeCall(
                IERC20Permit.permit,
                (
                    permitArgs.account,
                    address(this),
                    permitArgs.value,
                    permitArgs.deadline,
                    permitArgs.v,
                    permitArgs.r,
                    permitArgs.s
                )
            );
            _safeCall(payable(address(permitArgs.token)), 0, data);
        }

        _executeBridge(remoteChainID, token, permitArgs.account, to, value, gasFee, exFee, true, extraData);
        return true;
    }

    /**
     * @notice Processes multiple permit-based bridge operations in batch
     * @dev Handles multiple permit bridge requests in a single transaction
     * - Validates all permit signatures
     * - Processes each bridge operation
     * - Tracks success/failure for each operation
     * @param args Array of bridge operation details
     * @param permitArgs Array of permit details matching bridge operations
     */
    function permitBridgeTokenBatch(BridgeTokenArguments[] calldata args, PermitArguments[] calldata permitArgs)
        external
        payable
    {
        require(args.length == permitArgs.length, StandardNotMatchLength());
        for (uint i = 0; i < args.length; ++i) {
            permitBridgeToken(
                args[i].remoteChainID,
                args[i].token,
                args[i].to,
                args[i].value,
                args[i].gasFee,
                args[i].exFee,
                args[i].extraData,
                permitArgs[i]
            );
        }
    }

    /**
     * @notice Finalizes a bridge operation on the target chain
     * @dev Processes the completion of a bridge transfer
     * - Validates validator signatures
     * - Handles token minting/transfer
     * - Manages pending operations if needed
     * @param args Finalization arguments including token and amount
     * @param v Array of signature v values
     * @param r Array of signature r values
     * @param s Array of signature s values
     * @return success Boolean indicating successful finalization
     */
    function finalizeBridge(FinalizeArguments calldata args, uint8[] memory v, bytes32[] memory r, bytes32[] memory s)
        public
        payable
        whenNotPaused
        nonReentrant
        returns (bool)
    {
        require(_tokens[args.remoteChainID].contains(address(args.token)), StandardNotExistToken(address(args.token)));
        require(msg.value == 0, StandardInvalidMsgValue(0, msg.value));

        uint index = _useFinalizeIndex(args.remoteChainID);
        require(args.index == index, StandardInvalidIndex(index, args.index));

        _validateSignature(
            keccak256(
                abi.encode(FINALIZE_TYPEHASH, args.index, address(args.token), args.to, args.value, args.extraData)
            ),
            v,
            r,
            s
        );

        bool ok;
        bytes memory reason;
        bool delay;
        {
            TokenPair memory tokenPair = _tokenPairs[args.remoteChainID][address(args.token)];
            if (tokenPair.paused) {
                reason = "token paused";
            } else if (tokenPair.safetyLimit != 0 && tokenPair.safetyLimit < args.value) {
                reason = "over safety limit";
                delay = true;
            } else {
                (ok, reason) = _finalizeBridge(args.remoteChainID, args.token, args.to, args.value);
            }
        }

        if (ok) {
            emit BridgeFinalized(args.remoteChainID, args.index, args.token, args.to, args.value, block.timestamp);
        } else {
            _setPendingArguments(args, reason, delay);
            emit BridgeFinalizePending(args.index);
        }

        return true;
    }

    /**
     * @notice Finalizes multiple bridges on the local chain in a batch.
     * @param args An array of FinalizeArguments containing the details of each bridge.
     * @param v An array of arrays of v values for the signatures.
     * @param r An array of arrays of r values for the signatures.
     * @param s An array of arrays of s values for the signatures.
     * @return A boolean indicating whether all bridges were finalized successfully.
     */
    function finalizeBridgeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool) {
        for (uint i = 0; i < args.length; ++i) {
            finalizeBridge(args[i], v[i], r[i], s[i]);
        }
        return true;
    }

    /**
     * @notice Retries a pending bridge finalization
     * @dev Attempts to complete a previously failed finalization
     * - Validates pending operation exists
     * - Updates pending amounts
     * - Processes token transfer/minting
     * @param remoteChainID Chain ID of the source chain
     * @param index Index of the pending operation
     * @return success Boolean indicating successful retry
     */
    function retryFinalizeBridge(uint remoteChainID, uint index) public whenNotPaused nonReentrant returns (bool) {
        require(_pendingIndex[remoteChainID].contains(index), StandardNotExistIndex(index));

        PendingData memory pending = _pendingData[remoteChainID][index];
        TokenPair memory tokenPair = _tokenPairs[remoteChainID][address(pending.args.token)];
        require(!tokenPair.paused, RegistryTokenPaused(address(pending.args.token)));
        require(
            pending.safeDeadline == 0 || pending.safeDeadline < block.timestamp,
            StandardNotExpiredSafeDeadline(pending.safeDeadline, block.timestamp)
        );

        return _processPending(remoteChainID, index);
    }

    /**
     * @notice Retries finalizing multiple bridges on the local chain in a batch.
     * @param remoteChainID The ID of the remote chain.
     * @param indexes An array of indexes of the bridges to retry.
     * @return A boolean indicating whether all bridges were retried successfully.
     */
    function retryFinalizeBridgeBatch(uint remoteChainID, uint[] memory indexes) external returns (bool) {
        for (uint i = 0; i < indexes.length; ++i) {
            retryFinalizeBridge(remoteChainID, indexes[i]);
        }
        return true;
    }

    /**
     * @notice Processes pending operations that have exceeded their safety deadline
     * @dev Automatically processes expired pending operations up to a maximum count
     * - Checks if any pending operations have expired
     * - Iterates through all chains and their pending operations
     * - Processes operations that have exceeded their safety deadline
     * - Stops after processing maxCount operations
     * - Skips paused chains and tokens
     * @param maxCount Maximum number of operations to process
     */
    function processExpiredPending(uint maxCount) external {
        // Skip if no pending operations have expired
        if (_latestExpiredPendingTime < block.timestamp) return;

        uint count = 0;
        uint[] memory chainIDs = _chains.values();
        for (uint i = 0; i < chainIDs.length; ++i) {
            uint chainID = chainIDs[i];
            // Skip paused chains
            if (_chainData[chainID].paused) continue;

            uint[] memory indexes = _pendingIndex[chainID].values();
            for (uint j = 0; j < indexes.length; j++) {
                uint index = indexes[j];

                PendingData memory pending = _pendingData[chainID][index];
                // Process operations where:
                // 1. The token is not paused
                // 2. The operation has a safety deadline (not manually locked)
                // 3. The safety deadline has passed
                if (
                    !_tokenPairs[pending.args.remoteChainID][address(pending.args.token)].paused
                        && pending.safeDeadline != 0 && pending.safeDeadline < block.timestamp
                ) {
                    // Process the pending operation
                    _processPending(chainID, index);
                    if (++count >= maxCount) return;
                }
            }
        }
    }

    /**
     * @notice Locks a pending operation to prevent automatic processing
     * @dev Sets the safety deadline to maximum uint value
     * - Requires Admin role (ADMIN_ROLE)
     * - Verifies pending operation exists
     * - Emits PendingLocked event
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     */
    function lockPending(uint remoteChainID, uint index) external onlyAdmin {
        require(_pendingIndex[remoteChainID].contains(index), StandardNotExistIndex(index));
        _pendingData[remoteChainID][index].safeDeadline = type(uint).max;
        emit PendingLocked(remoteChainID, index);
    }

    /**
     * @notice Manually processes a pending operation regardless of pause or safety deadline
     * @dev Admin-only function to force process a pending operation
     * - Bypasses token pause and safety deadline checks
     * - Requires Admin role (ADMIN_ROLE)
     * - Verifies pending operation exists
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     * @return success Boolean indicating successful processing
     */
    function manualProcessPending(uint remoteChainID, uint index) external onlyAdmin returns (bool) {
        require(_pendingIndex[remoteChainID].contains(index), StandardNotExistIndex(index));
        return _processPending(remoteChainID, index);
    }

    /**
     * @notice Internal function to handle bridge initiation
     * @dev Processes token deposit or burn for bridging
     * - Handles native token transfers
     * - Manages ERC20 token transfers
     * - Updates deposit tracking
     * @param remoteChainID Target chain identifier
     * @param token Token being bridged
     * @param from Source address
     * @param value Amount being bridged
     * @param fee Total fees to collect
     */
    function _initiateBridge(uint remoteChainID, IERC20 token, address from, uint value, uint fee) internal virtual {
        int exchangeRate = _exchangeRate[remoteChainID][address(token)];
        if (exchangeRate > 1) require(value % uint(exchangeRate) == 0, StandardInvalidValueUnit(token, value));

        if (address(token) == NATIVE_TOKEN) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, StandardInvalidMsgValue(value + fee, msg.value));
            if (fee != 0) _safeCall(_nexus, fee, "");
        } else {
            require(msg.value == 0, StandardInvalidMsgValue(0, msg.value));

            token.safeTransferFrom(from, address(this), value + fee);
            if (fee != 0) token.safeTransfer(_nexus, fee);

            if (_tokenPairs[remoteChainID][address(token)].isOrigin) {
                _depositToken(remoteChainID, address(token), value);
            } else {
                require(
                    ICrossMintableERC20(address(token)).burn(address(this), value), // Burn the wrapped tokens on the source chain
                    StandardBurnFailed(token, from, value)
                );
            }
        }
    }

    /**
     * @notice Internal function to complete bridge finalization
     * @dev Handles token distribution on target chain
     * - Validates safety limits
     * - Processes token rates
     * - Handles native/ERC20 token transfers
     * @param remoteChainID Source chain identifier
     * @param token Token to distribute
     * @param to Recipient address
     * @param value Amount to distribute
     * @return ok Success status
     * @return reason Error message if failed
     */
    function _finalizeBridge(uint remoteChainID, IERC20 token, address to, uint value)
        internal
        virtual
        returns (bool ok, bytes memory reason)
    {
        int exchangeRate = _exchangeRate[remoteChainID][address(token)];
        if (exchangeRate > 1) value = value * uint(exchangeRate);
        else if (exchangeRate < -1) value = value / uint(-exchangeRate);

        if (address(token) == NATIVE_TOKEN) {
            _safeCall(payable(to), value, "");
            ok = true;
        } else if (value != 0) {
            if (_tokenPairs[remoteChainID][address(token)].isOrigin) {
                return _transferERC20(remoteChainID, token, to, value);
            } else {
                return _mintCrossMintableERC20(token, to, value);
            }
        }
    }

    /**
     * @notice Executes the bridge operation
     * @dev Core internal function that handles the bridge process
     * - Initiates token deposit/burn
     * - Generates operation index
     * - Emits events for tracking
     * @param remoteChainID Target chain identifier
     * @param token Token being bridged
     * @param from Source address
     * @param to Destination address
     * @param value Amount being bridged
     * @param gasFee Gas fee for transaction
     * @param exFee Exchange fee
     * @param permit Whether permit was used
     * @param extraData Additional operation data
     */
    function _executeBridge(
        uint remoteChainID,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bool permit,
        bytes calldata extraData
    ) private {
        _initiateBridge(remoteChainID, token, from, value, gasFee + exFee);

        uint index;
        IERC20 remoteToken;
        {
            index = _useInitiateIndex(remoteChainID);
            remoteToken = IERC20(_tokenPairs[remoteChainID][address(token)].remoteToken);
        }

        emit BridgeInitiated(
            remoteChainID, index, token, remoteToken, from, to, value, block.timestamp, permit, extraData
        );
        emit BridgeFeeCharged(index, token, from, gasFee, exFee);
    }

    function _transferERC20(uint remoteChainID, IERC20 token, address to, uint value)
        private
        returns (bool ok, bytes memory reason)
    {
        try IERC20(token).transfer(to, value) returns (bool success) {
            ok = success;
            if (success) _withdrawToken(remoteChainID, address(token), value);
            else reason = "transfer failed";
        } catch (bytes memory lowLevelData) {
            reason = lowLevelData;
        }
    }

    function _mintCrossMintableERC20(IERC20 token, address to, uint value)
        private
        returns (bool ok, bytes memory reason)
    {
        try ICrossMintableERC20(address(token)).mint(to, value) returns (bool success) {
            ok = success;
            if (!success) reason = "mint failed";
        } catch (bytes memory lowLevelData) {
            reason = lowLevelData;
        }
    }

    // pending 중인 요청 처리
    function _processPending(uint remoteChainID, uint index) private returns (bool) {
        FinalizeArguments memory args = _removePendingArguments(remoteChainID, index);

        (bool ok, bytes memory reason) = _finalizeBridge(args.remoteChainID, args.token, args.to, args.value);
        require(ok, string(reason));

        emit BridgeFinalized(remoteChainID, args.index, args.token, args.to, args.value, block.timestamp);
        return true;
    }

    function _safeCall(address payable recipient, uint amount, bytes memory data) private {
        require(address(this).balance >= amount, StandardInvalidBalance(amount, address(this).balance));
        (bool success, bytes memory returndata) = recipient.call{value: amount}(data);
        if (!success) {
            if (returndata.length > 0) {
                assembly ("memory-safe") {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert StandardFailedCall();
            }
        }

        if (amount == 0) {
            // 성공 시 추가 검증
            if (returndata.length == 0) {
                // 반환 데이터가 없으면 컨트랙트 코드 존재 여부 확인
                require(address(recipient).code.length != 0, StandardFailedCall());
            } else {
                // 반환 데이터가 있으면 값이 1인지 확인
                bool returnValue;
                assembly {
                    returnValue := mload(add(returndata, 32))
                }
                require(returnValue == true, StandardFailedCall());
            }
        }
    }

    /**
     * ░█░█░▀█▀░█▀▀░█░█░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▄▀░░█░░█▀▀░█▄█░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░░▀░░▀▀▀░▀▀▀░▀░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Returns the domain separator.
     * @return The domain separator.
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /**
     * @notice Returns the block number at which the contract was initialized.
     * @return The block number at which the contract was initialized.
     */
    function initializedAt() external view returns (uint) {
        return _initializedAt;
    }

    /**
     * @notice Estimates the fee for bridging a token.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to bridge.
     * @param value The amount of the token to bridge.
     * @return minimumAmount The minimum amount of the token to bridge.
     * @return gasFee The gas fee for the bridge.
     * @return exFee The extra fee for the bridge.
     */
    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee)
    {
        (minimumAmount, gasFee, exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
    }

    function nexus() external view returns (address) {
        return _nexus;
    }

    /**
     * @notice Validates and adjusts fee amounts for bridge operations
     * @dev Checks if provided fees meet minimum requirements
     * - Verifies minimum token amount
     * - Validates gas fee is sufficient
     * - Validates exchange fee is sufficient
     * - Returns adjusted fee values
     * @param remoteChainID Chain ID for fee calculation
     * @param token Token being bridged
     * @param value Amount being bridged
     * @param gasFee Provided gas fee
     * @param exFee Provided exchange fee
     * @return _gasFee Adjusted gas fee
     * @return _exFee Adjusted exchange fee
     */
    function _checkAmount(uint remoteChainID, IERC20 token, uint value, uint gasFee, uint exFee)
        private
        view
        returns (uint _gasFee, uint _exFee)
    {
        if (address(bridgeFeeStation) == address(0)) return (0, 0);
        uint minimumAmount;
        (minimumAmount, _gasFee, _exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
        require(value >= minimumAmount && gasFee >= _gasFee && exFee >= _exFee, StandardInvalidAmount());
    }

    /**
     * ░█▀▀░█▀▀░▀█▀░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▀█░█▀▀░░█░░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░▀▀▀░▀▀▀░░▀░░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Pauses the contract.
     */
    function setPause(bool set) external onlyOwner {
        if (set) _pause();
        else _unpause();
    }

    /**
     * @notice Sets the reward wallet.
     * @param nexus_ The address of the reward wallet.
     */
    function setNexus(address payable nexus_) external onlyOwner {
        require(nexus_ != address(0), StandardCanNotZeroAddress());
        _nexus = nexus_;
        emit NexusSet(_nexus);
    }

    /**
     * @notice Sets the fee station.
     * @param _bridgeFeeStation The address of the fee station.
     */
    function setFeeStation(IBridgeFeeStation _bridgeFeeStation) external onlyAdmin {
        bridgeFeeStation = _bridgeFeeStation;
        emit FeeStationSet(bridgeFeeStation);
    }

    /**
     * @notice Authorizes an upgrade to a new implementation.
     * @param _newImplementation The address of the new implementation.
     */
    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
