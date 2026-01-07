// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

import {IBridgeExecutor} from "./interface/IBridgeExecutor.sol";
import {Const} from "./lib/Const.sol";

/**
 * @title BridgeExecutor
 * @notice Executes bridge operations with extra data
 * @dev This contract handles complex bridge finalizations that require calling external contracts
 *
 * Key features:
 * - Parses extraData to extract target contract address and calldata
 * - Handles both native token and ERC20 token operations
 * - For native tokens: receives value from bridge
 * - For ERC20 tokens: pulls tokens from bridge via transferFrom
 * - Returns consumed amount; sends remaining to user directly
 *
 * Security:
 * - MAY revert - on revert, entire bridge finalization rolls back safely
 * - Target contracts must be whitelisted
 * - Optional method selector whitelisting per target
 */
contract BridgeExecutor is AccessControl, ReentrancyGuardTransient, IBridgeExecutor {
    using SafeERC20 for IERC20;

    /// @notice Thrown when target contract is not in the whitelist
    error BETargetNotWhitelisted();

    /// @notice Thrown when method selector is not whitelisted for the target
    error BEMethodNotWhitelisted();

    /// @notice Thrown when an invalid address (zero address) is provided
    error BEInvalidAddress();

    /// @notice Thrown when trying to add a target that is already whitelisted
    error BEAlreadyWhitelisted();

    /// @notice Thrown when trying to remove a target that is not whitelisted
    error BENotWhitelisted();

    /// @notice Thrown when recipient address is invalid for token recovery
    error BEInvalidRecipient();

    /// @notice Thrown when native token transfer fails during recovery
    error BEFailedToReturnNative();

    /// @notice Thrown when gas reserve value is out of valid range (50k-1M)
    error BEInvalidGasReserve();

    /// @notice Thrown when extraData is too short (< 24 bytes)
    error BEInvalidExtraData();

    /// @notice Thrown when msg.value doesn't match expected value for native token
    error BEInvalidValue();

    /// @notice Thrown when ERC20 transferFrom fails
    error BETransferFromFailed();

    /// @notice Thrown when value is zero
    error BEZeroValue();

    /// @notice Thrown when target contract call fails
    error BETargetCallFailed();

    /**
     * @notice Target whitelist configuration
     * @param isWhitelisted Whether the target is whitelisted
     * @param methodCheckEnabled Whether method selector checking is enabled for this target
     */
    struct Target {
        bool isWhitelisted;
        bool methodCheckEnabled;
    }

    /**
     * @notice Emitted when a target contract is added to whitelist
     * @param target Target contract address
     */
    event TargetWhitelisted(address indexed target);

    /**
     * @notice Emitted when a target contract is removed from whitelist
     * @param target Target contract address
     */
    event TargetRemovedFromWhitelist(address indexed target);

    /**
     * @notice Emitted when a method is added to a target's whitelist
     * @param target Target contract address
     * @param methodID Function selector
     */
    event MethodWhitelisted(address indexed target, bytes4 indexed methodID);

    /**
     * @notice Emitted when a method is removed from a target's whitelist
     * @param target Target contract address
     * @param methodID Function selector
     */
    event MethodRemovedFromWhitelist(address indexed target, bytes4 indexed methodID);

    /**
     * @notice Emitted when method check is toggled for a target
     * @param target Target contract address
     * @param enabled Whether method check is enabled
     */
    event MethodCheckEnabled(address indexed target, bool enabled);

    /**
     * @notice Emitted when post-call gas reserve is changed
     * @param oldValue Previous gas reserve value
     * @param newValue New gas reserve value
     */
    event PostCallGasReserveSet(uint oldValue, uint newValue);

    /// @dev Target whitelist configuration mapping
    mapping(address => Target) private _targets;

    /// @dev Mapping of whitelisted method selectors per target
    mapping(address => mapping(bytes4 => bool)) private _whitelistedMethods;

    /// @dev Post-call gas reserve (adjustable by admin)
    uint private _postCallGasReserve = 150_000;

    /// @dev Maximum return data size to prevent OOG from unbounded returndata
    uint private constant MAX_RETURN_DATA_SIZE = 1024;

    /**
     * @notice Constructor
     * @param owner Address of the admin owner
     */
    constructor(address owner, address executor) {
        require(owner != address(0), BEInvalidAddress());
        _grantRole(DEFAULT_ADMIN_ROLE, owner);
        _grantRole(Const.ADMIN_ROLE, owner);

        if (executor != address(0)) _grantRole(Const.EXECUTOR_ROLE, executor);
    }

    /**
     * @notice Executes extra call with bridge finalization data
     * @dev This function MAY revert. On revert, entire finalization is rolled back.
     *      - For ERC20: pulls tokens from msg.sender via transferFrom
     *      - For Native: receives value via msg.value
     *      - Returns consumed amount; sends remaining to user directly
     */
    function executeExtraCall(
        uint, /* fromChainID */
        uint, /* index */
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable onlyRole(Const.EXECUTOR_ROLE) nonReentrant returns (uint consumed, bytes memory returnData) {
        // value == 0 → revert
        require(value > 0, BEZeroValue());

        // Parse extraData: target (20 bytes) + methodID (4 bytes) + calldata
        require(extraData.length >= 24, BEInvalidExtraData());

        address targetContract;
        bytes4 methodID;
        assembly {
            targetContract := shr(96, calldataload(extraData.offset))
            methodID := calldataload(add(extraData.offset, 20))
        }

        // Whitelist checks (revert if not valid)
        Target storage target = _targets[targetContract];
        require(target.isWhitelisted, BETargetNotWhitelisted());
        if (target.methodCheckEnabled) require(_whitelistedMethods[targetContract][methodID], BEMethodNotWhitelisted());

        bool isNative = address(toToken) == Const.NATIVE_TOKEN;

        // Pull tokens from bridge
        if (isNative) {
            require(msg.value == value, BEInvalidValue());
        } else {
            require(msg.value == 0, BEInvalidValue());
            // Pull tokens from bridge via transferFrom
            toToken.safeTransferFrom(msg.sender, address(this), value);
        }

        // Record balance before call
        uint balBefore = isNative ? address(this).balance - msg.value : toToken.balanceOf(address(this));

        // Approve target for ERC20
        if (!isNative) toToken.forceApprove(targetContract, value);

        // Reserve minimum gas for post-call logic
        uint gasReserve = _postCallGasReserve;
        uint remainingGas = gasleft() < gasReserve ? 0 : gasleft() - gasReserve;

        // Call target contract with bounded returndata to prevent OOG
        bool success;
        (success, returnData) = _safeCall(targetContract, remainingGas, isNative ? value : 0, extraData[20:]);

        // Clear approval (always for ERC20)
        if (!isNative) toToken.forceApprove(targetContract, 0);

        // Target call failed → revert (bridge will fallback to user)
        require(success, BETargetCallFailed());

        // Calculate consumed based on balance change
        uint balAfter = isNative ? address(this).balance : toToken.balanceOf(address(this));

        if (isNative) {
            // For native: target may return ETH, so check if balance increased
            uint returned = balAfter > balBefore ? balAfter - balBefore : 0;
            consumed = returned >= value ? 0 : value - returned;
        } else {
            consumed = balBefore > balAfter ? balBefore - balAfter : 0;
        }

        // Send remaining to user directly
        uint remaining = value > consumed ? value - consumed : 0;
        if (remaining > 0) {
            if (isNative) {
                (bool ok,) = payable(to).call{value: remaining}("");
                require(ok, BEFailedToReturnNative());
            } else {
                toToken.safeTransfer(to, remaining);
            }
        }
    }

    // =========================
    // Admin functions
    // =========================

    /**
     * @notice Add a target contract to the whitelist
     * @dev Only callable by admin
     * @param target Target contract address to whitelist
     */
    function addWhitelistTarget(address target) external onlyRole(Const.ADMIN_ROLE) {
        require(target != address(0), BEInvalidAddress());
        require(!_targets[target].isWhitelisted, BEAlreadyWhitelisted());
        _targets[target].isWhitelisted = true;
        emit TargetWhitelisted(target);
    }

    /**
     * @notice Remove a target contract from the whitelist
     * @dev Only callable by admin
     * @param target Target contract address to remove
     */
    function removeWhitelistTarget(address target) external onlyRole(Const.ADMIN_ROLE) {
        require(_targets[target].isWhitelisted, BENotWhitelisted());
        _targets[target].isWhitelisted = false;
        emit TargetRemovedFromWhitelist(target);
    }

    /**
     * @notice Enable/disable method check for a target
     * @dev Only callable by admin
     * @param target Target contract address
     * @param enabled Whether to enable method check
     */
    function setMethodCheckEnabled(address target, bool enabled) external onlyRole(Const.ADMIN_ROLE) {
        require(target != address(0), BEInvalidAddress());
        _targets[target].methodCheckEnabled = enabled;
        emit MethodCheckEnabled(target, enabled);
    }

    /**
     * @notice Add multiple method selectors to the whitelist for a target
     * @dev Only callable by admin
     * @param target Target contract address
     * @param methodIDs Array of function selectors to whitelist
     */
    function addWhitelistMethods(address target, bytes4[] calldata methodIDs) external onlyRole(Const.ADMIN_ROLE) {
        require(target != address(0), BEInvalidAddress());
        for (uint i = 0; i < methodIDs.length; ++i) {
            _whitelistedMethods[target][methodIDs[i]] = true;
            emit MethodWhitelisted(target, methodIDs[i]);
        }
    }

    /**
     * @notice Remove multiple method selectors from the whitelist for a target
     * @dev Only callable by admin
     * @param target Target contract address
     * @param methodIDs Array of function selectors to remove
     */
    function removeWhitelistMethods(address target, bytes4[] calldata methodIDs) external onlyRole(Const.ADMIN_ROLE) {
        require(target != address(0), BEInvalidAddress());
        for (uint i = 0; i < methodIDs.length; ++i) {
            _whitelistedMethods[target][methodIDs[i]] = false;
            emit MethodRemovedFromWhitelist(target, methodIDs[i]);
        }
    }

    /**
     * @notice Set the post-call gas reserve value
     * @dev Only callable by admin
     * @param value New gas reserve amount
     */
    function setPostCallGasReserve(uint value) external onlyRole(Const.ADMIN_ROLE) {
        require(value >= 50_000 && value <= 1_000_000, BEInvalidGasReserve());
        uint oldValue = _postCallGasReserve;
        _postCallGasReserve = value;
        emit PostCallGasReserveSet(oldValue, value);
    }

    /**
     * @notice Emergency function to recover stuck tokens
     * @dev Only callable by ADMIN_ROLE
     * @param token Token to recover (address(1) for native token, i.e., Const.NATIVE_TOKEN)
     * @param amount Amount to recover
     * @param recipient Address to send recovered tokens to
     */
    function recoverToken(address token, uint amount, address recipient) external onlyRole(Const.ADMIN_ROLE) {
        require(recipient != address(0) && token != address(0), BEInvalidRecipient());
        if (token == Const.NATIVE_TOKEN) {
            (bool ok,) = payable(recipient).call{value: amount}("");
            require(ok, BEFailedToReturnNative());
        } else {
            IERC20(token).safeTransfer(recipient, amount);
        }
    }

    // =========================
    // View functions
    // =========================

    /**
     * @notice Check if a target contract is whitelisted
     * @param target Target contract address to check
     * @return True if target is whitelisted
     */
    function isWhitelistedTarget(address target) external view returns (bool) {
        return _targets[target].isWhitelisted;
    }

    /**
     * @notice Check if method check is enabled for a target
     * @param target Target contract address to check
     * @return True if method check is enabled
     */
    function isMethodCheckEnabled(address target) external view returns (bool) {
        return _targets[target].methodCheckEnabled;
    }

    /**
     * @notice Check if a method selector is whitelisted for a target
     * @param target Target contract address
     * @param methodID Function selector
     * @return True if method is whitelisted
     */
    function isWhitelistedMethod(address target, bytes4 methodID) external view returns (bool) {
        return _whitelistedMethods[target][methodID];
    }

    /**
     * @notice Get the current post-call gas reserve value
     * @return The gas reserve amount
     */
    function postCallGasReserve() external view returns (uint) {
        return _postCallGasReserve;
    }

    /**
     * @notice Allows contract to receive native tokens
     */
    receive() external payable {}

    // =========================
    // Internal helpers
    // =========================

    /**
     * @dev Makes external call with bounded return data size to prevent OOG
     * @param target Target contract address
     * @param gasLimit Gas limit for the call
     * @param value Native value to send
     * @param data Calldata to send
     * @return success True if call succeeded
     * @return returnData Return data (capped at MAX_RETURN_DATA_SIZE)
     */
    function _safeCall(address target, uint gasLimit, uint value, bytes memory data)
        private
        returns (bool success, bytes memory returnData)
    {
        uint maxSize = MAX_RETURN_DATA_SIZE;
        assembly {
            // Allocate memory for returnData
            returnData := mload(0x40)

            // Make the call
            success := call(gasLimit, target, value, add(data, 0x20), mload(data), 0, 0)

            // Cap returndata size
            let size := returndatasize()
            if gt(size, maxSize) { size := maxSize }

            // Store length and copy limited data
            mstore(returnData, size)
            returndatacopy(add(returnData, 0x20), 0, size)

            // Update free memory pointer (32-byte aligned)
            mstore(0x40, add(add(returnData, 0x20), and(add(size, 31), not(31))))
        }
    }
}
