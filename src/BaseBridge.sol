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
import {IBaseBridge} from "./interface/IBaseBridge.sol";
import {IBridgeFeeManager} from "./interface/IBridgeFeeManager.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title BaseBridge
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
 * - bridgeFeeManager: Contract managing bridge fees
 * - _dev: Reward wallet address for collecting fees
 * - _initializedAt: Block number when contract was initialized
 *
 * @dev Bridge process flow:
 * 1. User calls bridgeToken/permitBridgeToken
 * 2. Tokens are deposited or burned in contract
 * 3. BridgeInitiated event is emitted
 * 4. Validators provide signatures
 * 5. finalizeBridge mints/transfers tokens on target chain
 */
contract BaseBridge is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    BridgeRegistry,
    ValidatorManager,
    IBaseBridge
{
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    error BaseBridgeInvalidValueUnit(IERC20 token, uint value);
    error BaseBridgeInvalidValue(uint expected, uint actual);
    error BaseBridgeInvalidIndex(uint expected, uint actual);
    error BaseBridgeInvalidMsgValue(uint expected, uint actual);
    error BaseBridgeInvalidBalance(uint expected, uint actual);
    error BaseBridgeInvalidAmount();
    error BaseBridgeInvalidPermitToken(address expected, address actual);
    error BaseBridgeCanNotZeroMsgValue();
    error BaseBridgeCanNotZeroAddress();
    error BaseBridgeNotExistIndex(uint index);
    error BaseBridgeNotExistToken(address token);
    error BaseBridgeNotExpired(uint delayExpiration, uint timestamp);
    error BaseBridgeBurnFailed(IERC20 token, address from, uint value);
    error BaseBridgeNotMatchLength();
    error BaseBridgeFailedCall(string reason);

    /**
     * @notice Emitted when a bridge operation is initiated
     * @param toChainID The ID of the destination chain
     * @param index Unique identifier for this bridge operation
     * @param fromToken Token being bridged from the source chain
     * @param toToken Corresponding token on the destination chain
     * @param from Address initiating the bridge operation
     * @param to Recipient address on the destination chain
     * @param value Amount of tokens being bridged
     * @param gasFee Gas fee for the bridge operation
     * @param exFee Exchange fee for the bridge operation
     * @param time Timestamp when the operation was initiated
     * @param permit Whether the bridge operation used a permit
     * @param extraData Additional data for the bridge operation
     */
    event BridgeInitiated(
        uint indexed toChainID,
        uint indexed index,
        IERC20 fromToken,
        IERC20 toToken,
        address indexed from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        uint time,
        bool permit,
        bytes extraData
    );

    /**
     * @notice Emitted when a bridge operation is finalized on the destination chain
     * @param fromChainID The ID of the source chain
     * @param index Unique identifier for this bridge operation
     * @param toToken Token that was bridged to the destination chain
     * @param to Recipient address receiving the tokens
     * @param value Amount of tokens received
     * @param time Timestamp when the finalization occurred
     */
    event BridgeFinalized(
        uint indexed fromChainID, uint indexed index, IERC20 indexed toToken, address to, uint value, uint time
    );

    /**
     * @notice Emitted when a bridge operation is finalized on the destination chain
     * @param fromChainID The ID of the source chain
     * @param index Unique identifier for this bridge operation
     * @param toToken Token that was bridged to the destination chain
     * @param to Recipient address receiving the tokens
     * @param value Amount of tokens received
     * @param time Timestamp when the finalization occurred
     */
    event BridgePending(
        uint indexed fromChainID, uint indexed index, IERC20 indexed toToken, address to, uint value, uint time
    );

    /**
     * @notice Emitted when a pending bridge operation is manually locked by an admin
     * @param fromChainID The ID of the source chain
     * @param index Unique identifier for the locked bridge operation
     */
    event VerificationDelaySet(uint indexed fromChainID, uint indexed index, uint delay);

    /// @dev Native token representation address
    address internal constant NATIVE_TOKEN = address(1);

    /// @dev Hash for finalize operation signature verification
    bytes32 private constant FINALIZE_TYPEHASH = keccak256(
        "FinalizeBridge(uint256 fromChainID,uint256 index,address toToken,address to,uint256 value,bytes extraData)"
    );

    /// @dev Fee station contract for managing bridge fees
    IBridgeFeeManager public bridgeFeeManager;

    /// @dev Reward wallet for collecting fees
    address payable private _dev;

    /// @dev Block number when contract was initialized
    uint private _initializedAt;

    /// @dev Storage gap for future upgrades
    uint[47] private __gap;

    /**
     * @notice Contract constructor
     * @dev Disables direct initialization of implementations
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Allows contract to receive native tokens
     * @dev Reverts if sent value is zero
     */
    receive() external payable {
        require(msg.value != 0, BaseBridgeCanNotZeroMsgValue());
    }

    /**
     * @notice Initializes the bridge contract
     * @dev One-time initializer for the bridge contract
     * - Sets up initial state including validator threshold and reward wallet
     * - Calls internal initialization function
     * @param _threshold Required number of validator signatures
     * @param dev_ Address of reward wallet
     */
    function initialize(uint8 _threshold, address dev_) external virtual initializer {
        __BaseBridge_init(_threshold, dev_);
    }

    /**
     * @notice Initializes the bridge contract
     * @dev Sets up initial state including validator threshold and reward wallet
     * @param _threshold Required number of validator signatures
     * @param dev_ Address of reward wallet
     */
    function __BaseBridge_init(uint8 _threshold, address dev_) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __Validator_init(_threshold);
        __BridgeRegistry_init();

        require(address(dev_) != address(0), BaseBridgeCanNotZeroAddress());
        _dev = payable(dev_);

        _initializedAt = block.number;
    }

    /**
     * @notice Bridges a token to a remote chain
     * @dev Core function for initiating a bridge transfer
     * - Validates token and amounts
     * - Collects fees
     * - Handles token deposit/burn
     * - Emits events for tracking
     * @param toChainID Target chain identifier
     * @param fromToken Token address to bridge
     * @param to Recipient address on target chain
     * @param value Amount of tokens to bridge
     * @param gasFee Gas fee for transaction on target chain
     * @param exFee Exchange fee for the bridge service
     * @param extraData Additional data for the bridge operation
     * @return success Boolean indicating successful initiation
     */
    function bridgeToken(
        uint toChainID,
        IERC20 fromToken,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData
    ) public payable whenNotPaused onlyValidToken(toChainID, address(fromToken)) nonReentrant returns (bool) {
        (gasFee, exFee) = _checkAmount(toChainID, fromToken, value, gasFee, exFee);
        _executeBridge(toChainID, fromToken, _msgSender(), to, value, gasFee, exFee, false, extraData);
        return true;
    }

    /**
     * @notice Bridges tokens using permit functionality
     * @dev Allows bridging with permit-based approval
     * - Validates permit signature
     * - Processes bridge operation
     * - Handles token transfer and fees
     * @param toChainID Target chain identifier
     * @param fromToken Token address to bridge
     * @param to Recipient address on target chain
     * @param value Amount of tokens to bridge
     * @param gasFee Gas fee for target chain
     * @param exFee Exchange fee
     * @param extraData Additional bridge data
     * @param permitArgs Permit signature and parameters
     * @return success Boolean indicating successful initiation
     */
    function permitBridgeToken(
        uint toChainID,
        IERC20 fromToken,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData,
        PermitArguments calldata permitArgs
    ) public payable whenNotPaused onlyValidToken(toChainID, address(fromToken)) nonReentrant returns (bool) {
        require(
            address(fromToken) == address(permitArgs.token),
            BaseBridgeInvalidPermitToken(address(fromToken), address(permitArgs.token))
        );
        (gasFee, exFee) = _checkAmount(toChainID, fromToken, value, gasFee, exFee);

        require(
            permitArgs.value >= value + gasFee + exFee, BaseBridgeInvalidValue(value + gasFee + exFee, permitArgs.value)
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
            (bool success, bytes memory reason) = _safeCall(payable(address(permitArgs.token)), 0, data);
            require(success, BaseBridgeFailedCall(string(reason)));
        }

        _executeBridge(toChainID, fromToken, permitArgs.account, to, value, gasFee, exFee, true, extraData);
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
        require(args.length == permitArgs.length, BaseBridgeNotMatchLength());
        for (uint i = 0; i < args.length; ++i) {
            permitBridgeToken(
                args[i].toChainID,
                args[i].fromToken,
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
        require(
            _tokens[args.fromChainID].contains(address(args.toToken)), BaseBridgeNotExistToken(address(args.toToken))
        );
        require(msg.value == 0, BaseBridgeInvalidMsgValue(0, msg.value));

        uint index = _useFinalizeIndex(args.fromChainID);
        require(args.index == index, BaseBridgeInvalidIndex(index, args.index));

        _validateSignature(
            keccak256(
                abi.encode(FINALIZE_TYPEHASH, args.index, address(args.toToken), args.to, args.value, args.extraData)
            ),
            v,
            r,
            s
        );

        bool ok;
        bytes memory reason;
        bool delay;
        {
            TokenPair memory tokenPair = _tokenPairs[args.fromChainID][address(args.toToken)];
            if (tokenPair.paused) {
                reason = "token paused";
            } else if (tokenPair.verificationAmountThreshold != 0 && tokenPair.verificationAmountThreshold < args.value)
            {
                reason = "over safety limit";
                delay = true;
            } else {
                (ok, reason) = _finalizeBridge(args.fromChainID, args.toToken, args.to, args.value);
            }
        }

        if (ok) {
            emit BridgeFinalized(args.fromChainID, args.index, args.toToken, args.to, args.value, block.timestamp);
        } else {
            _setPendingArguments(args, reason, delay);
            emit BridgePending(args.fromChainID, args.index, args.toToken, args.to, args.value, block.timestamp);
        }
        return true;
    }

    /**
     * @notice Finalizes multiple bridges on the local chain in a batch
     * @dev Processes multiple finalize operations in a single transaction
     * @param args An array of FinalizeArguments containing the details of each bridge
     * @param v An array of arrays of v values for the signatures
     * @param r An array of arrays of r values for the signatures
     * @param s An array of arrays of s values for the signatures
     * @return Boolean indicating whether all bridges were finalized successfully
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
    function releasePending(uint remoteChainID, uint index) public whenNotPaused nonReentrant returns (bool) {
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));

        PendingData memory pending = _pendingData[remoteChainID][index];
        TokenPair memory tokenPair = _tokenPairs[remoteChainID][address(pending.args.toToken)];
        require(!tokenPair.paused, RegistryTokenPaused(address(pending.args.toToken)));
        require(
            pending.delayExpiration == 0 || pending.delayExpiration < block.timestamp,
            BaseBridgeNotExpired(pending.delayExpiration, block.timestamp)
        );

        return _releasePending(remoteChainID, index);
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
    function releaseExpiredPending(uint maxCount) external {
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
                    !_tokenPairs[pending.args.fromChainID][address(pending.args.toToken)].paused
                        && pending.delayExpiration != 0 && pending.delayExpiration < block.timestamp
                ) {
                    // Process the pending operation
                    _releasePending(chainID, index);
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
     * - Emits PendingDataLocked event
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     */
    function setVerificationDelay(uint remoteChainID, uint index, uint delay) external onlyAdmin {
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));
        _pendingData[remoteChainID][index].delayExpiration = block.timestamp + delay;
        emit VerificationDelaySet(remoteChainID, index, delay);
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
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));
        return _releasePending(remoteChainID, index);
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
        int conversionRatio = _conversionRatio[remoteChainID][address(token)];
        if (conversionRatio > 1) require(value % uint(conversionRatio) == 0, BaseBridgeInvalidValueUnit(token, value));

        if (address(token) == NATIVE_TOKEN) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, BaseBridgeInvalidMsgValue(value + fee, msg.value));
            if (fee != 0) {
                (bool success, bytes memory reason) = _safeCall(_dev, fee, "");
                require(success, BaseBridgeFailedCall(string(reason)));
            }
        } else {
            require(msg.value == 0, BaseBridgeInvalidMsgValue(0, msg.value));

            token.safeTransferFrom(from, address(this), value + fee);
            if (fee != 0) token.safeTransfer(_dev, fee);

            if (_tokenPairs[remoteChainID][address(token)].isOrigin) {
                _depositToken(remoteChainID, address(token), value);
            } else {
                require(
                    ICrossMintableERC20(address(token)).burn(address(this), value), // Burn the wrapped tokens on the source chain
                    BaseBridgeBurnFailed(token, from, value)
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
        int conversionRatio = _conversionRatio[remoteChainID][address(token)];
        if (conversionRatio > 1) value = value * uint(conversionRatio);
        else if (conversionRatio < -1) value = value / uint(-conversionRatio);

        if (address(token) == NATIVE_TOKEN) {
            return _safeCall(payable(to), value, "");
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
            remoteChainID, index, token, remoteToken, from, to, value, gasFee, exFee, block.timestamp, permit, extraData
        );
    }

    /**
     * @notice Transfers ERC20 tokens during bridge finalization
     * @dev Handles transfer of original tokens on destination chain
     * - Updates withdrawal tracking
     * - Transfers tokens to recipient
     * @param remoteChainID Source chain identifier
     * @param token Token to transfer
     * @param to Recipient address
     * @param value Amount to transfer
     * @return ok Success status
     * @return reason Error message if failed
     */
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

    /**
     * @notice Mints wrapped tokens during bridge finalization
     * @dev Handles minting of CrossMintable tokens on destination chain
     * @param token Token to mint
     * @param to Recipient address
     * @param value Amount to mint
     * @return ok Success status
     * @return reason Error message if failed
     */
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

    /**
     * @notice Processes a pending bridge operation
     * @dev Internal function to complete pending operations
     * - Removes pending data
     * - Finalizes the bridge operation
     * - Emits finalization event
     * @param remoteChainID Source chain identifier
     * @param index Index of the pending operation
     * @return success Boolean indicating successful processing
     */
    function _releasePending(uint remoteChainID, uint index) private returns (bool) {
        FinalizeArguments memory args = _removePendingArguments(remoteChainID, index);

        (bool ok, bytes memory reason) = _finalizeBridge(args.fromChainID, args.toToken, args.to, args.value);
        require(ok, string(reason));

        emit BridgeFinalized(args.fromChainID, args.index, args.toToken, args.to, args.value, block.timestamp);
        return true;
    }

    /**
     * @notice Safely calls an external contract
     * @dev Handles external call with value transfer and error handling
     * - Verifies balance is sufficient
     * - Makes external call with provided value and data
     * - Validates return value if applicable
     * @param recipient Target address
     * @param amount Native token amount to send
     * @param data Call data
     * @return success Success status
     * @return reason Error message if failed
     */
    function _safeCall(address payable recipient, uint amount, bytes memory data)
        private
        returns (bool success, bytes memory reason)
    {
        require(address(this).balance >= amount, BaseBridgeInvalidBalance(amount, address(this).balance));

        bytes memory returndata;
        (success, returndata) = recipient.call{value: amount}(data);
        if (!success) {
            if (returndata.length > 0) return (false, returndata);
            else return (false, "reverted");
        }

        if (amount == 0) {
            if (returndata.length == 0) {
                if (address(recipient).code.length == 0) return (false, "not code");
            } else {
                bool returnValue;
                assembly {
                    returnValue := mload(add(returndata, 32))
                }
                if (returnValue != true) return (false, "return false");
            }
        }

        return (true, "");
    }

    /**
     * ░█░█░▀█▀░█▀▀░█░█░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▄▀░░█░░█▀▀░█▄█░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░░▀░░▀▀▀░▀▀▀░▀░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Returns the domain separator for EIP-712 signatures
     * @dev Used for signature verification in finalizeBridge
     * @return The domain separator hash
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /**
     * @notice Returns the block number at which the contract was initialized
     * @dev Used for tracking contract deployment time
     * @return The block number at initialization
     */
    function initializedAt() external view returns (uint) {
        return _initializedAt;
    }

    /**
     * @notice Estimates the fee for bridging a token
     * @dev Queries the fee manager for calculated fee values
     * @param remoteChainID Target chain identifier
     * @param token Token to bridge
     * @param value Amount to bridge
     * @return minimumValue Minimum amount of token required for bridging
     * @return gasFee Gas fee for transaction on target chain
     * @return exFee Exchange fee for the bridge service
     */
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint gasFee, uint exFee)
    {
        (minimumValue, gasFee, exFee) = bridgeFeeManager.calculateFee(remoteChainID, token, value);
    }

    /**
     * @notice Returns the reward wallet address
     * @dev Used for fee collection
     * @return Address of the reward wallet
     */
    function dev() external view returns (address) {
        return _dev;
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
        if (address(bridgeFeeManager) == address(0)) return (0, 0);
        uint minimumValue;
        (minimumValue, _gasFee, _exFee) = bridgeFeeManager.calculateFee(remoteChainID, token, value);
        require(value >= minimumValue && gasFee >= _gasFee && exFee >= _exFee, BaseBridgeInvalidAmount());
    }

    /**
     * ░█▀▀░█▀▀░▀█▀░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▀█░█▀▀░░█░░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░▀▀▀░▀▀▀░░▀░░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Pauses or unpauses the contract
     * @dev Controls whether bridge operations can be executed
     * @param set True to pause, false to unpause
     */
    function setPause(bool set) external onlyOwner {
        if (set) _pause();
        else _unpause();
    }

    /**
     * @notice Updates the reward wallet address
     * @dev Changes the destination for collected fees
     * @param dev_ New reward wallet address
     */
    function setDev(address payable dev_) external onlyOwner {
        require(dev_ != address(0), BaseBridgeCanNotZeroAddress());
        _dev = dev_;
    }

    /**
     * @notice Sets the fee manager contract
     * @dev Updates the contract used for fee calculations
     * @param _bridgeFeeManager New fee manager address
     */
    function setFeeManager(IBridgeFeeManager _bridgeFeeManager) external onlyAdmin {
        bridgeFeeManager = _bridgeFeeManager;
    }

    /**
     * @notice Authorizes an upgrade to a new implementation
     * @dev Controls UUPS upgradability
     * @param _newImplementation Address of the new implementation
     */
    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
