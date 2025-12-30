// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardTransientUpgradeable} from
    "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {BridgeRegistry} from "./abstract/BridgeRegistry.sol";
import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {IBaseBridge} from "./interface/IBaseBridge.sol";

import {IBridgeExecutor} from "./interface/IBridgeExecutor.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {Const} from "./lib/Const.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title BaseBridge
 * @notice Core implementation of cross-chain bridge functionality
 * @dev Provides token bridging between chains with multi-signature verification
 * Key features:
 * - Token bridging with transaction finalization
 * - Multi-signature security with threshold validation
 * - Fee management and safety mechanisms
 * - Support for native tokens and ERC20 tokens
 */
contract BaseBridge is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardTransientUpgradeable,
    BridgeRegistry,
    ValidatorManager,
    IBaseBridge
{
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    error BaseBridgeInvalidValue(uint expected, uint actual);
    error BaseBridgeInvalidIndex(uint expected, uint actual);
    error BaseBridgeInvalidBalance();
    error BaseBridgeInvalidAmount();
    error BaseBridgeInvalidPermitToken(address expected, address actual);
    error BaseBridgeFailedPermitBatch(uint index, bytes reason);
    error BaseBridgeCanNotZeroAddress();
    error BaseBridgeCanNotZeroValue();
    error BaseBridgeVerifierNotSet();
    error BaseBridgeNotExistIndex(uint index);
    error BaseBridgeNotExistToken(address token);
    error BaseBridgeNotExpired(uint delayExpiration, uint timestamp);
    error BaseBridgeBurnFailed(IERC20 token, address from, uint value);
    error BaseBridgeNotMatchLength();
    error BaseBridgeFailedCall();
    error BaseBridgeMismatchPermitAccount();
    error BaseBridgeFailedRelease(Const.FinalizeStatus status);
    error BaseBridgeOnlyExecutor();

    /**
     * @notice Emitted when a bridge operation is initiated
     * @param toChainID Destination chain ID
     * @param index Unique identifier for this bridge operation
     * @param fromToken Source token being bridged
     * @param toToken Target token on destination chain
     * @param from Address initiating the bridge
     * @param to Recipient address on destination chain
     * @param value Amount of tokens being bridged
     * @param networkFee Network fee for the operation
     * @param exFee Exchange fee for the operation
     * @param time Timestamp of initiation
     * @param extraData Additional operation data
     */
    event BridgeInitiated(
        uint indexed toChainID,
        uint indexed index,
        IERC20 fromToken,
        IERC20 toToken,
        address indexed from,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes extraData,
        uint time
    );

    /**
     * @notice Emitted when a bridge operation is finalized
     * @param fromChainID Source chain ID
     * @param index Unique identifier for this operation
     * @param toToken Token on destination chain
     * @param to Recipient address
     * @param value Amount of tokens received
     * @param time Timestamp of finalization
     */
    event BridgeFinalized(
        uint indexed fromChainID, uint indexed index, IERC20 indexed toToken, address to, uint value, uint time
    );

    /**
     * @notice Emitted when a bridge operation enters pending state
     * @param fromChainID Source chain ID
     * @param index Unique identifier for this operation
     * @param toToken Token on destination chain
     * @param to Recipient address
     * @param value Amount of tokens to be received
     * @param time Timestamp when operation entered pending state
     * @param status Status of the pending operation
     */
    event BridgePending(
        uint indexed fromChainID,
        uint indexed index,
        IERC20 indexed toToken,
        address to,
        uint value,
        uint time,
        Const.FinalizeStatus status
    );

    /**
     * @notice Emitted when a manual release is performed
     * @param remoteChainID The ID of the chain the operation was performed on
     * @param index The index of the operation
     */
    event ManualReleased(uint indexed remoteChainID, uint indexed index);

    /**
     * @notice Emitted when a pending operation is removed without being processed
     * @param remoteChainID The ID of the chain the operation was removed from
     * @param index The index of the operation that was removed without processing
     */
    event PendingRemoved(uint indexed remoteChainID, uint indexed index);

    /**
     * @notice Emitted when delay expiration is set for a pending operation
     * @param fromChainID Source chain ID
     * @param index Unique identifier for the operation
     * @param expiration Timestamp when the verification delay expires
     */
    event VerificationDelayExpirationSet(uint indexed fromChainID, uint indexed index, uint expiration);

    /**
     * @notice Emitted when the bridge verifier is set
     * @param bridgeVerifier The address of the new bridge verifier
     */
    event BridgeVerifierSet(address indexed bridgeVerifier);

    /**
     * @notice Emitted when the wallet is set
     * @param dev The address of the new wallet
     */
    event DevSet(address indexed dev);

    /**
     * @notice Emitted when the bridge executor is set
     * @param bridgeExecutor The address of the new bridge executor
     */
    event BridgeExecutorSet(address indexed bridgeExecutor);

    /**
     * @notice Emitted when extra call is executed via BridgeExecutor
     * @param fromChainID Source chain ID
     * @param index Unique identifier for the operation
     * @param success Whether the extra call succeeded
     */
    event ExtraCallExecuted(uint indexed fromChainID, uint indexed index, bool success);

    /// @dev Hash for finalize operation signature verification
    bytes32 private constant FINALIZE_TYPEHASH = keccak256(
        "FinalizeBridge(uint256 fromChainID,uint256 index,address toToken,address to,uint256 value,bytes extraData)"
    );

    /// @dev Fee management contract
    IBridgeVerifier public bridgeVerifier;

    /// @dev dev wallet
    address payable private _dev;

    /// @dev Block number when contract was initialized
    uint private _initializedAt;

    /// @dev Bridge executor contract for handling extradata operations
    IBridgeExecutor public bridgeExecutor;

    /// @dev Storage gap for future upgrades
    uint[46] private __gap;

    /**
     * @notice Contract constructor
     * @dev Disables direct initialization of implementations
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the bridge contract
     * @param owner Owner address
     * @param dev_ Reward wallet address
     * @param threshold_ Required validator signatures
     */
    function initialize(address owner, address payable dev_, uint8 threshold_) external initializer {
        __BaseBridge_init(owner, dev_, threshold_);
    }

    /**
     * @notice Internal initialization function
     * @param owner Owner address
     * @param dev_ Reward wallet address
     * @param threshold_ Required validator signatures
     */
    function __BaseBridge_init(address owner, address payable dev_, uint8 threshold_) internal onlyInitializing {
        require(owner != address(0), BaseBridgeCanNotZeroAddress());
        require(dev_ != address(0), BaseBridgeCanNotZeroAddress());
        require(threshold_ != 0, BaseBridgeCanNotZeroValue());

        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuardTransient_init();
        __RoleManager_init(owner);
        __Validator_init(threshold_);
        __BridgeRegistry_init();

        _dev = dev_;
        emit DevSet(dev_);

        _initializedAt = block.number;
    }

    /**
     * @notice Bridges a token to a remote chain
     * @param toChainID Target chain ID
     * @param fromToken Token to bridge
     * @param to Recipient address
     * @param value Amount to bridge
     * @param networkFee Network fee for transaction
     * @param exFee Exchange fee
     * @param extraData Additional data
     * @return success Operation status
     */
    function bridgeToken(
        uint toChainID,
        IERC20 fromToken,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes calldata extraData
    ) public payable whenNotPaused onlyValidToken(toChainID, address(fromToken)) nonReentrant returns (bool) {
        require(to != address(0), BaseBridgeCanNotZeroAddress());
        (networkFee, exFee) = _checkInitiateAmount(toChainID, fromToken, value, networkFee, exFee);
        _executeBridge(
            BridgeTokenArguments({
                toChainID: toChainID,
                fromToken: fromToken,
                from: _msgSender(),
                to: to,
                value: value,
                networkFee: networkFee,
                exFee: exFee,
                extraData: extraData
            })
        );
        return true;
    }

    /**
     * @notice Bridges tokens using permit functionality
     * @dev extraData parameter is intentionally ignored to prevent front-running attacks
     * @param toChainID Target chain ID
     * @param fromToken Token to bridge
     * @param to Recipient address (must match permit account)
     * @param value Amount to bridge
     * @param networkFee Network fee
     * @param exFee Exchange fee
     * @param permitArgs Permit signature parameters
     * @return success Operation status
     */
    function permitBridgeToken(
        uint toChainID,
        IERC20 fromToken,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes calldata, // extraData ignored to prevent front-running attacks
        PermitArguments calldata permitArgs
    ) public payable whenNotPaused onlyValidToken(toChainID, address(fromToken)) nonReentrant returns (bool) {
        require(
            address(fromToken) == address(permitArgs.token),
            BaseBridgeInvalidPermitToken(address(fromToken), address(permitArgs.token))
        );

        // The permitBridgeToken function doesn't restrict msg.sender, allowing anyone to initiate the bridge operation.
        // However, it requires that 'to' matches permitArgs.account to ensure that the recipient on the target chain
        // is the same as the account that signed the permit. This restriction provides protection against front-running
        // by enforcing that only the intended recipient can receive the bridged tokens.
        require(to == permitArgs.account, BaseBridgeMismatchPermitAccount());

        (networkFee, exFee) = _checkInitiateAmount(toChainID, fromToken, value, networkFee, exFee);
        require(
            permitArgs.value >= value + networkFee + exFee,
            BaseBridgeInvalidValue(value + networkFee + exFee, permitArgs.value)
        );

        bridgeVerifier.safePermit(
            fromToken,
            permitArgs.account,
            address(this),
            permitArgs.value,
            permitArgs.deadline,
            permitArgs.v,
            permitArgs.r,
            permitArgs.s
        );

        _executeBridge(
            BridgeTokenArguments({
                toChainID: toChainID,
                fromToken: fromToken,
                from: permitArgs.account,
                to: to,
                value: value,
                networkFee: networkFee,
                exFee: exFee,
                extraData: "" // Empty: extraData not supported in permit flow
            })
        );
        return true;
    }

    /**
     * @notice Processes multiple permit-based bridge operations
     * @param args Array of bridge operation details
     * @param permitArgs Array of permit parameters
     */
    function permitBridgeTokenBatch(BridgeTokenArguments[] calldata args, PermitArguments[] calldata permitArgs)
        external
    {
        require(args.length == permitArgs.length, BaseBridgeNotMatchLength());
        for (uint i = 0; i < args.length; ++i) {
            try this.permitBridgeToken(
                args[i].toChainID,
                args[i].fromToken,
                args[i].to,
                args[i].value,
                args[i].networkFee,
                args[i].exFee,
                args[i].extraData,
                permitArgs[i]
            ) {} catch (bytes memory reason) {
                revert BaseBridgeFailedPermitBatch(i, reason);
            }
        }
    }

    /**
     * @notice Finalizes a bridge operation on target chain
     * @param args Finalization parameters
     * @param v Signature v values array
     * @param r Signature r values array
     * @param s Signature s values array
     * @return success Operation status
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
        require(msg.value == 0, BaseBridgeInvalidValue(0, msg.value));

        uint index = _useFinalizeIndex(args.fromChainID);
        require(args.index == index, BaseBridgeInvalidIndex(index, args.index));

        bytes32 messageHash = keccak256(
            abi.encode(
                FINALIZE_TYPEHASH,
                args.fromChainID,
                args.index,
                address(args.toToken),
                args.to,
                args.value,
                keccak256(args.extraData)
            )
        );
        _validateSignature(messageHash, v, r, s);

        Const.FinalizeStatus status;
        bool delay;
        {
            (status, delay) = _checkFinalizeAmount(args.fromChainID, args.toToken, args.value, false);
            if (status == Const.FinalizeStatus.Success) {
                status =
                    _finalizeBridge(args.fromChainID, args.index, args.toToken, args.to, args.value, args.extraData);
            }
        }

        if (status == Const.FinalizeStatus.Success) {
            emit BridgeFinalized(args.fromChainID, args.index, args.toToken, args.to, args.value, block.timestamp);
        } else {
            _setPendingArguments(args, status, delay);
            emit BridgePending(args.fromChainID, args.index, args.toToken, args.to, args.value, block.timestamp, status);
        }
        return true;
    }

    /**
     * @notice Finalizes multiple bridge operations
     * @param args Array of finalization parameters
     * @param v Array of signature v values arrays
     * @param r Array of signature r values arrays
     * @param s Array of signature s values arrays
     * @return success Operation status
     */
    function finalizeBridgeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool) {
        uint length = args.length;
        require(v.length == length && r.length == length && s.length == length, BaseBridgeNotMatchLength());

        for (uint i = 0; i < length; ++i) {
            finalizeBridge(args[i], v[i], r[i], s[i]);
        }
        return true;
    }

    /**
     * @notice Retries a pending bridge finalization
     * @dev Attempts to complete a previously failed finalization.
     * If this function reverts due to permanent failures (e.g., recipient contract
     * rejecting transfers), use manualReleasePendingWithRecipient to specify
     * an alternative recipient address for recovery.
     * @param remoteChainID Chain ID of the source chain
     * @param index Index of the pending operation
     */
    function releasePending(uint remoteChainID, uint index) public whenNotPaused nonReentrant {
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));

        PendingData memory pending = _pendingData[remoteChainID][index];

        (Const.FinalizeStatus status,) =
            _checkFinalizeAmount(remoteChainID, pending.args.toToken, pending.args.value, true);

        require(status == Const.FinalizeStatus.Success, string(abi.encode(uint(status))));
        require(
            pending.delayExpiration == 0 || pending.delayExpiration < block.timestamp,
            BaseBridgeNotExpired(pending.delayExpiration, block.timestamp)
        );

        _releasePending(remoteChainID, index, address(0));
    }

    /**
     * @notice Attempts to process a pending bridge finalization
     * @dev Retries a previously failed finalization
     * - Validates pending operation exists
     * - Verifies token is not paused
     * - Checks verification delay has expired
     * - Updates pending amounts
     * - Processes token transfer/minting
     * @param remoteChainIDs Chain IDs of the source chains
     * @param indexes Indexes of the pending operations
     */
    function releasePendingBatch(uint[] memory remoteChainIDs, uint[] memory indexes) external {
        require(remoteChainIDs.length == indexes.length, BaseBridgeNotMatchLength());
        for (uint i = 0; i < indexes.length; ++i) {
            releasePending(remoteChainIDs[i], indexes[i]);
        }
    }

    /**
     * @notice Manually releases a pending operation regardless of pause or safety deadline
     * @dev Verifier-only function to force process a pending operation
     * - Bypasses token pause and safety deadline checks
     * - Verifies pending operation exists
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     */
    function manualReleasePending(uint remoteChainID, uint index) external onlyRole(Const.VERIFIER_ROLE) {
        manualReleasePendingWithRecipient(remoteChainID, index, address(0));
    }

    /**
     * @notice Manually release a pending operation regardless of pause or safety deadline with a custom recipient
     * @dev Verifier-only function to force release a pending operation
     * - Bypasses token pause and safety deadline checks
     * - Verifies pending operation exists
     * - Allows overriding the original recipient address
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     * @param recipient Custom recipient address that will receive the tokens instead of the original recipient
     */
    function manualReleasePendingWithRecipient(uint remoteChainID, uint index, address recipient)
        public
        onlyRole(Const.VERIFIER_ROLE)
        nonReentrant
    {
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));
        _releasePending(remoteChainID, index, recipient);
        emit ManualReleased(remoteChainID, index);
    }

    /**
     * @notice Removes a pending operation without processing
     * @dev Only callable by admin role
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     */
    function removePending(uint remoteChainID, uint index) external onlyRole(Const.ADMIN_ROLE) nonReentrant {
        _removePendingArguments(remoteChainID, index);
        emit PendingRemoved(remoteChainID, index);
    }

    /**
     * @notice Locks a pending operation to prevent automatic processing
     * @dev Sets the safety deadline to maximum uint value
     * - Requires Verifier role (Const.VERIFIER_ROLE)
     * - Verifies pending operation exists
     * - Emits PendingDataLocked event
     * @param remoteChainID Chain ID of the pending operation
     * @param index Index of the pending operation
     */
    function setVerificationDelayExpiration(uint remoteChainID, uint index, uint delay)
        external
        onlyRole(Const.VERIFIER_ROLE)
    {
        require(_pendingIndex[remoteChainID].contains(index), BaseBridgeNotExistIndex(index));
        uint expiration = block.timestamp + delay;
        _pendingData[remoteChainID][index].delayExpiration = expiration;
        emit VerificationDelayExpirationSet(remoteChainID, index, expiration);
    }

    /**
     * @notice Internal function to execute a bridge operation
     * @dev Handles token deposit/burn and fee collection
     * @param args Bridge operation details
     */
    function _executeBridge(BridgeTokenArguments memory args) internal {
        uint index;
        IERC20 remoteToken;
        {
            index = _useInitiateIndex(args.toChainID);
            remoteToken = IERC20(_tokenPairs[args.toChainID][address(args.fromToken)].remoteToken);
        }

        _initiateBridge(args.toChainID, args.fromToken, args.from, args.value, args.networkFee + args.exFee);

        emit BridgeInitiated(
            args.toChainID,
            index,
            args.fromToken,
            remoteToken,
            args.from,
            args.to,
            args.value,
            args.networkFee,
            args.exFee,
            args.extraData,
            block.timestamp
        );
    }

    /**
     * @notice Processes a pending bridge operation
     * @dev Internal function to complete pending operations
     * - Removes pending data
     * - Finalizes the bridge operation
     * - Emits finalization event
     * @param remoteChainID Source chain identifier
     * @param index Index of the pending operation
     * @param recipient Recipient address
     */
    function _releasePending(uint remoteChainID, uint index, address recipient) private {
        FinalizeArguments memory args = _removePendingArguments(remoteChainID, index);
        if (recipient == address(0)) recipient = args.to;
        (Const.FinalizeStatus status) =
            _finalizeBridge(args.fromChainID, args.index, args.toToken, recipient, args.value, args.extraData);
        require(status == Const.FinalizeStatus.Success, BaseBridgeFailedRelease(status));
        emit BridgeFinalized(args.fromChainID, args.index, args.toToken, recipient, args.value, block.timestamp);
    }

    /**
     * @notice Validates and adjusts fee amounts for bridge operations
     * @dev Checks if provided fees meet minimum requirements
     * - Verifies minimum token amount
     * - Validates network fee is sufficient
     * - Validates exchange fee is sufficient
     * - Returns adjusted fee values
     * @param remoteChainID Chain ID for fee calculation
     * @param token Token being bridged
     * @param value Amount being bridged
     * @param networkFee Provided network fee
     * @param exFee Provided exchange fee
     * @return _networkFee Adjusted network fee
     * @return _exFee Adjusted exchange fee
     */
    function _checkInitiateAmount(uint remoteChainID, IERC20 token, uint value, uint networkFee, uint exFee)
        private
        view
        returns (uint _networkFee, uint _exFee)
    {
        TokenPair memory tokenPair = _tokenPairs[remoteChainID][address(token)];
        if (!tokenPair.isOrigin) require(tokenPair.minted >= value, BaseBridgeInvalidValue(value, tokenPair.minted));
        require(address(bridgeVerifier) != address(0), BaseBridgeVerifierNotSet());

        uint minimumValue;
        (minimumValue, _networkFee, _exFee) = bridgeVerifier.calculateFee(remoteChainID, token, value);
        require(
            value != 0 && value >= minimumValue && networkFee >= _networkFee && exFee >= _exFee,
            BaseBridgeInvalidAmount()
        );
    }

    /**
     * @notice Validates finalization amount against thresholds
     * @dev Checks if amount requires extra verification
     * @param fromChainID Source chain ID
     * @param token Token address
     * @param value Amount to verify
     * @return status Success status
     * @return delay Whether verification delay should be applied
     */
    function _checkFinalizeAmount(uint fromChainID, IERC20 token, uint value, bool retry)
        internal
        virtual
        returns (Const.FinalizeStatus status, bool delay)
    {
        require(address(bridgeVerifier) != address(0), BaseBridgeVerifierNotSet());

        // Check finalize pause status
        if (_tokenFinalizePaused[fromChainID][address(token)]) return (Const.FinalizeStatus.TokenPaused, false);

        // Skip validation if this is a retry - validation was already completed in the initial attempt
        if (!retry) {
            status = bridgeVerifier.validateBridgeTokenValue(token, value);
            if (status != Const.FinalizeStatus.Success) delay = true;
        } else {
            status = Const.FinalizeStatus.Success;
        }
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
        if (address(token) == Const.NATIVE_TOKEN) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, BaseBridgeInvalidValue(value + fee, msg.value));
            if (fee != 0) {
                bool success = _safeCall(_dev, fee, "");
                require(success, BaseBridgeFailedCall());
            }
        } else {
            require(msg.value == 0, BaseBridgeInvalidValue(0, msg.value));
            token.safeTransferFrom(from, address(this), value + fee);
            if (fee != 0) token.safeTransfer(_dev, fee);

            if (!_tokenPairs[remoteChainID][address(token)].isOrigin) {
                require(
                    ICrossMintableERC20(address(token)).burn(address(this), value), // Burn the wrapped tokens on the source chain
                    BaseBridgeBurnFailed(token, from, value)
                );
            }
        }
        _depositToken(remoteChainID, address(token), value);
    }

    /**
     * @notice Executes token transfer or minting for bridge finalization
     * @dev Handles different token types (native, mintable, regular)
     * - If extraData is provided and bridgeExecutor is set, delegates to BridgeExecutor
     * - On executor failure, reverts to normal token transfer to user
     * @param fromChainID Source chain ID
     * @param index Finalize operation index
     * @param toToken Destination token
     * @param to Recipient address
     * @param value Amount to transfer/mint
     * @param extraData Additional data for bridge executor (contract address + calldata)
     * @return status Success status
     */
    function _finalizeBridge(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes memory extraData
    ) internal returns (Const.FinalizeStatus status) {
        bool isOrigin = _tokenPairs[fromChainID][address(toToken)].isOrigin;

        // Check if extraData is provided and long enough (>= 24 bytes for contract address and method ID)
        address executor = address(bridgeExecutor);
        if (extraData.length >= 24 && executor != address(0) && executor.code.length != 0) {
            // Parse target contract address from extraData
            address targetContract;
            assembly {
                targetContract := shr(96, mload(add(extraData, 32)))
            }

            // Check if target is whitelisted (NO try/catch - use low-level staticcall and safe bool parsing)
            bool isWhitelisted;
            {
                (bool ok, bytes memory returndata) =
                    executor.staticcall(abi.encodeCall(IBridgeExecutor.isWhitelistedTarget, (targetContract)));
                if (ok && returndata.length >= 32) {
                    assembly {
                        isWhitelisted := mload(add(returndata, 32))
                    }
                }
            }
            if (isWhitelisted) {
                bool isERC20 = address(toToken) != Const.NATIVE_TOKEN;

                // For ERC20, transfer/mint to Executor first
                // For Native token, send directly via msg.value in executeExtraCall
                if (isERC20) {
                    status = _transferOrMintToken(toToken, executor, value, isOrigin, false);
                    if (status != Const.FinalizeStatus.Success) return status;
                }

                // Call executeExtraCall and check result
                // Returns (bool success, uint remaining)
                bool success;
                uint remaining;
                {
                    (bool ok, bytes memory returndata) = executor.call{value: isERC20 ? 0 : value}(
                        abi.encodeCall(
                            IBridgeExecutor.executeExtraCall, (fromChainID, index, toToken, to, value, extraData)
                        )
                    );
                    if (ok && returndata.length >= 64) {
                        assembly {
                            success := mload(add(returndata, 32))
                            remaining := mload(add(returndata, 64))
                        }
                    }
                }
                emit ExtraCallExecuted(fromChainID, index, success);

                // 1. remaining == 0: executor가 이미 user에게 전송 완료
                if (remaining == 0) {
                    _withdrawToken(fromChainID, address(toToken), value);
                    return Const.FinalizeStatus.Success;
                }

                // 2. remaining > 0: executor에서 bridge로 회수
                if (isERC20) {
                    if (isOrigin) toToken.safeTransferFrom(executor, address(this), remaining);
                    else ICrossMintableERC20(address(toToken)).burn(executor, remaining);
                }
                // Native: executor가 이미 bridge로 전송함

                // 3. remaining을 user에게 전송
                status = _transferOrMintToken(toToken, to, remaining, isOrigin, false);
                if (status == Const.FinalizeStatus.Success) _withdrawToken(fromChainID, address(toToken), value);
                return status;
            }
        }

        // Normal token transfer flow (no extraData or not whitelisted)
        status = _transferOrMintToken(toToken, to, value, isOrigin, false);
        if (status == Const.FinalizeStatus.Success) _withdrawToken(fromChainID, address(toToken), value);
        return status;
    }

    /**
     * @notice Transfers or mints tokens to recipient
     * @param toToken Token to transfer/mint
     * @param to Recipient address
     * @param value Amount to transfer/mint
     * @param isOrigin Whether token is origin token
     * @param forceTransfer If true, always use transfer (token already minted)
     * @return status Success status
     */
    function _transferOrMintToken(IERC20 toToken, address to, uint value, bool isOrigin, bool forceTransfer)
        private
        returns (Const.FinalizeStatus status)
    {
        if (address(toToken) == Const.NATIVE_TOKEN) {
            if (!_safeCall(payable(to), value, "")) return Const.FinalizeStatus.TransferFailed;
            return Const.FinalizeStatus.Success;
        }
        if (value == 0) return Const.FinalizeStatus.Success;

        bytes memory data;
        if (forceTransfer || isOrigin) {
            data = abi.encodeCall(IERC20.transfer, (to, value));
            status = Const.FinalizeStatus.TransferFailed;
        } else {
            data = abi.encodeCall(ICrossMintableERC20.mint, (to, value));
            status = Const.FinalizeStatus.MintFailed;
        }
        return _safeCall(payable(address(toToken)), 0, data) ? Const.FinalizeStatus.Success : status;
    }

    /**
     * @notice Safely calls an external contract
     * @dev Handles external call with value transfer and error handling
     * - Verifies balance is sufficient
     * - Makes external call with provided value and data
     * - Validates return value if applicable
     * @param recipient Target address
     * @param amount Native token amount to send
     * @param data Call data (ERC20 token transfer or ERC20 token minting)
     * @return success Success status
     */
    function _safeCall(address payable recipient, uint amount, bytes memory data) private returns (bool success) {
        require(address(this).balance >= amount, BaseBridgeInvalidBalance());

        bytes memory returndata;
        (success, returndata) = recipient.call{value: amount, gas: 100000}(data);
        if (!success) return (false);

        if (amount == 0) {
            if (returndata.length == 0) {
                if (address(recipient).code.length == 0) return false;
            } else {
                bool returnValue;
                assembly {
                    returnValue := mload(add(returndata, 32))
                }
                return returnValue;
            }
        }
        return true;
    }

    /**
     * ░█░█░▀█▀░█▀▀░█░█░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▄▀░░█░░█▀▀░█▄█░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░░▀░░▀▀▀░▀▀▀░▀░▀░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
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
     * @notice Returns the reward wallet address
     * @dev Used for fee collection
     * @return Address of the reward wallet
     */
    function dev() external view returns (address) {
        return _dev;
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
    function setPause(bool set) external onlyRole(Const.OPERATOR_ROLE) {
        if (set) _pause();
        else _unpause();
    }

    /**
     * @notice Updates the reward wallet address
     * @dev Changes the destination for collected fees
     * @param dev_ New reward wallet address
     */
    function setDev(address payable dev_) external onlyRole(Const.ADMIN_ROLE) {
        require(dev_ != address(0), BaseBridgeCanNotZeroAddress());
        _dev = dev_;
        emit DevSet(dev_);
    }

    /**
     * @notice Sets the bridge verifier contract
     * @dev Updates the contract used for bridge verification
     * @param _bridgeVerifier New bridge verifier address
     */
    function setBridgeVerifier(IBridgeVerifier _bridgeVerifier) external onlyRole(Const.ADMIN_ROLE) {
        require(address(_bridgeVerifier) != address(0), BaseBridgeCanNotZeroAddress());
        bridgeVerifier = _bridgeVerifier;
        emit BridgeVerifierSet(address(_bridgeVerifier));
    }

    /**
     * @notice Sets the bridge executor contract
     * @dev Updates the contract used for executing bridge operations with extradata
     * @param _bridgeExecutor New bridge executor address
     */
    function setBridgeExecutor(IBridgeExecutor _bridgeExecutor) external onlyRole(Const.ADMIN_ROLE) {
        bridgeExecutor = _bridgeExecutor;
        emit BridgeExecutorSet(address(_bridgeExecutor));
    }

    /**
     * @notice Receives native tokens from BridgeExecutor when extradata call fails
     * @dev Only BridgeExecutor can send native tokens to this contract
     */
    receive() external payable {
        require(msg.sender == address(bridgeExecutor), BaseBridgeOnlyExecutor());
    }

    /**
     * @notice Authorizes an upgrade to a new implementation
     * @dev Controls UUPS upgradability
     * @param _newImplementation Address of the new implementation
     */
    function _authorizeUpgrade(address _newImplementation) internal override onlyRole(Const.ADMIN_ROLE) {}
}
