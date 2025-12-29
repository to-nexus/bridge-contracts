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
 * - For native tokens: forwards value to target contract
 * - For ERC20 tokens: pulls tokens from bridge and forwards to target
 * - Always performs balance verification and returns remaining tokens
 *
 * Security:
 * - Only callable by authorized bridge contracts
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
     * @notice Emitted when an extra call execution is performed
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param targetContract Contract that was called
     * @param methodID Method ID of the called contract
     * @param success Whether the execution was successful
     * @param remaining Amount of tokens remaining (returned to user)
     * @param returnData Call response data (success: return value, failure: revert reason)
     */
    event ExtraCallExecuted(
        uint indexed fromChainID,
        uint indexed index,
        address indexed targetContract,
        bytes4 methodID,
        bool success,
        uint remaining,
        bytes returnData
    );

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
    uint private _postCallGasReserve = 200_000;

    /// @dev Native transfer gas limit (sufficient for receive/fallback)
    uint private constant NATIVE_TRANSFER_GAS = 30_000;

    /// @dev Maximum return data size to prevent OOG from unbounded returndata
    uint private constant MAX_RETURN_DATA_SIZE = 1024;

    /**
     * @notice Constructor
     * @param owner Address of the admin owner
     * @param bridge Address of the bridge contract that can call this executor
     */
    constructor(address owner, address bridge) {
        require(owner != address(0) && bridge != address(0), BEInvalidAddress());
        _grantRole(DEFAULT_ADMIN_ROLE, owner);
        _grantRole(Const.EXECUTOR_ROLE, bridge);
    }

    /**
     * @notice Executes extra call with bridge finalization data
     * @dev This function MUST NOT revert. All failures return (false, 0) or (true/false, remaining).
     * Always performs balance verification and returns remaining tokens to bridge.
     */
    function executeExtraCall(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable onlyRole(Const.EXECUTOR_ROLE) nonReentrant returns (bool success, uint remaining) {
        // IMPORTANT: This function is invoked during bridge finalization.
        // It must NOT revert. All validation failures should return (false, 0).

        // value == 0 → skip without event
        if (value == 0) return (false, 0);

        address targetContract;
        bytes4 methodID;
        // extraData >= 24 bytes: 20 (address) + 4 (selector)
        if (extraData.length >= 24) {
            assembly {
                targetContract := shr(96, calldataload(extraData.offset))
                methodID := calldataload(add(extraData.offset, 20))
            }
        } else {
            // Refund any unexpected native value (best-effort) and fail
            if (msg.value != 0) _trySendNative(msg.sender, msg.value, NATIVE_TRANSFER_GAS);
            // Send tokens directly to user on failure
            return _failAndRefund(toToken, value, to);
        }

        bool isNative = address(toToken) == Const.NATIVE_TOKEN;
        Target storage target = _targets[targetContract];

        // Reject non-whitelisted targets (no revert)
        if (!target.isWhitelisted) {
            // Send tokens directly to user on failure
            return _failAndRefund(toToken, value, to);
        }

        // Optional: methodID check per target (no revert)
        if (target.methodCheckEnabled && !_whitelistedMethods[targetContract][methodID]) {
            // Send tokens directly to user on failure
            return _failAndRefund(toToken, value, to);
        }

        // Record balance before call
        uint balBefore;
        if (isNative) {
            // For native: balance before = current balance - msg.value (what we just received)
            balBefore = address(this).balance - msg.value;
            if (msg.value != value) {
                // Refund whatever was sent (best-effort) and fail
                _trySendNative(msg.sender, msg.value, NATIVE_TRANSFER_GAS);
                return (false, 0);
            }
        } else {
            // Any unexpected msg.value is refunded to the bridge (best-effort)
            if (msg.value != 0) _trySendNative(msg.sender, msg.value, NATIVE_TRANSFER_GAS);

            (bool okBal, uint bal) = _tryBalanceOf(toToken, address(this));
            // Send actual balance to user on failure
            if (!okBal || bal < value) return _failAndRefund(toToken, bal, to);
            balBefore = bal;

            // Send tokens to user if target approval fails
            if (!_tryForceApprove(toToken, targetContract, value)) return _failAndRefund(toToken, value, to);
        }

        // Reserve minimum gas for post-extracall logic
        uint gasReserve = _postCallGasReserve;
        uint remainingGas = gasleft() < gasReserve ? 0 : gasleft() - gasReserve;

        // Call target contract with bounded returndata to prevent OOG
        bytes memory returnData;
        (success, returnData) = _safeCall(targetContract, remainingGas, isNative ? value : 0, extraData[20:]);

        // Clear approval (always for ERC20) - best-effort (no revert)
        if (!isNative) _tryForceApprove(toToken, targetContract, 0);

        // Calculate remaining based on balance change
        uint balAfter;
        if (isNative) {
            balAfter = address(this).balance;
        } else {
            (bool okBal, uint bal) = _tryBalanceOf(toToken, address(this));
            balAfter = okBal ? bal : 0;
        }

        // consumed = what was actually spent
        uint consumed;
        if (isNative) {
            uint returned = balAfter > balBefore ? balAfter - balBefore : 0;
            consumed = returned >= value ? 0 : value - returned;
        } else {
            consumed = balBefore > balAfter ? balBefore - balAfter : 0;
        }
        remaining = value > consumed ? value - consumed : 0;

        // Return remaining to user (try direct), fallback to bridge
        if (remaining > 0) {
            if (isNative) {
                uint reserve = _postCallGasReserve + NATIVE_TRANSFER_GAS;
                uint gasForTransfer = gasleft() > reserve ? gasleft() - reserve : NATIVE_TRANSFER_GAS;
                if (_trySendNative(to, remaining, gasForTransfer)) remaining = 0; // user received directly

                else _trySendNative(msg.sender, remaining, NATIVE_TRANSFER_GAS);
            } else {
                if (_tryTransfer(toToken, to, remaining)) remaining = 0; // user received directly

                else _tryForceApprove(toToken, msg.sender, remaining);
            }
        }

        emit ExtraCallExecuted(fromChainID, index, targetContract, methodID, success, remaining, returnData);
        return (success, remaining);
    }

    // =========================
    // Admin functions
    // =========================

    /**
     * @notice Add a target contract to the whitelist
     * @dev Only callable by admin
     * @param target Target contract address to whitelist
     */
    function addWhitelistTarget(address target) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
    function removeWhitelistTarget(address target) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
    function setMethodCheckEnabled(address target, bool enabled) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
    function addWhitelistMethods(address target, bytes4[] calldata methodIDs) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
    function removeWhitelistMethods(address target, bytes4[] calldata methodIDs)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
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
    function setPostCallGasReserve(uint value) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(value >= 50_000 && value <= 1_000_000, BEInvalidGasReserve());
        uint oldValue = _postCallGasReserve;
        _postCallGasReserve = value;
        emit PostCallGasReserveSet(oldValue, value);
    }

    /**
     * @notice Emergency function to recover stuck tokens
     * @dev Only callable by admin
     * @param token Token to recover (address(1) for native token, i.e., Const.NATIVE_TOKEN)
     * @param amount Amount to recover
     * @param recipient Address to send recovered tokens to
     */
    function recoverToken(address token, uint amount, address recipient) external onlyRole(DEFAULT_ADMIN_ROLE) {
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
    // Internal helpers (no revert)
    // =========================

    /**
     * @dev Attempts to send native tokens without reverting
     * @param recipient Address to send native tokens to
     * @param amount Amount of native tokens to send
     * @param gasLimit Gas limit for the call
     * @return ok True if transfer succeeded
     */
    function _trySendNative(address recipient, uint amount, uint gasLimit) private returns (bool ok) {
        if (amount == 0) return true;
        (ok,) = payable(recipient).call{value: amount, gas: gasLimit}("");
    }

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

    /**
     * @dev Attempts to get token balance without reverting
     * @param token ERC20 token to query
     * @param account Account to check balance for
     * @return ok True if call succeeded
     * @return bal Token balance
     */
    function _tryBalanceOf(IERC20 token, address account) private view returns (bool ok, uint bal) {
        bytes memory returndata;
        (ok, returndata) = address(token).staticcall(abi.encodeCall(IERC20.balanceOf, (account)));
        if (!ok || returndata.length < 32) return (false, 0);
        bal = abi.decode(returndata, (uint));
        return (true, bal);
    }

    /**
     * @dev Calls token function and safely parses optional bool return
     * @param token ERC20 token to call
     * @param data Encoded function call data
     * @return True if call succeeded and returned true (or no data with code)
     */
    function _callOptionalReturnBool(IERC20 token, bytes memory data) private returns (bool) {
        (bool ok, bytes memory returndata) = address(token).call(data);
        if (!ok) return false;

        // Tokens may return no data, assume success if target is a contract
        if (returndata.length == 0) return address(token).code.length != 0;

        if (returndata.length >= 32) return abi.decode(returndata, (bool));
        return false;
    }

    /**
     * @dev Attempts to approve token spending without reverting
     * @param token ERC20 token to approve
     * @param spender Address to approve spending for
     * @param amount Amount to approve
     * @return True if approval succeeded
     */
    function _tryApprove(IERC20 token, address spender, uint amount) private returns (bool) {
        return _callOptionalReturnBool(token, abi.encodeCall(IERC20.approve, (spender, amount)));
    }

    /**
     * @dev Attempts to force approve (handles tokens that require 0 allowance first)
     * @param token ERC20 token to approve
     * @param spender Address to approve spending for
     * @param amount Amount to approve
     * @return True if approval succeeded
     */
    function _tryForceApprove(IERC20 token, address spender, uint amount) private returns (bool) {
        // Try direct approve first
        if (_tryApprove(token, spender, amount)) return true;
        // Some tokens require setting to 0 before changing allowance
        if (!_tryApprove(token, spender, 0)) return false;
        return _tryApprove(token, spender, amount);
    }

    /**
     * @dev Attempts to transfer ERC20 tokens without reverting
     * @param token ERC20 token to transfer
     * @param to Recipient address
     * @param amount Amount to transfer
     * @return True if transfer succeeded
     */
    function _tryTransfer(IERC20 token, address to, uint amount) private returns (bool) {
        return _callOptionalReturnBool(token, abi.encodeCall(IERC20.transfer, (to, amount)));
    }

    /**
     * @dev Attempts to send tokens directly to user on early failure
     * @notice On early validation failure, tries to send tokens directly to user.
     *         If direct send fails, returns remaining for bridge to handle recovery.
     * @param token ERC20 token (or native token address)
     * @param value Amount to send
     * @param to User address to receive tokens
     * @return success Always false (early failure)
     * @return remaining 0 if sent to user, value if bridge needs to recover
     */
    function _failAndRefund(IERC20 token, uint value, address to) private returns (bool, uint) {
        if (value == 0) return (false, 0);

        if (address(token) == Const.NATIVE_TOKEN) {
            // Native: try sending directly to user
            uint reserve = _postCallGasReserve + NATIVE_TRANSFER_GAS;
            uint userGas = gasleft() > reserve ? gasleft() - reserve : NATIVE_TRANSFER_GAS;

            if (_trySendNative(to, value, userGas)) {
                // Success: user received tokens
                return (false, 0);
            }
            // Fallback: send to bridge for recovery
            _trySendNative(msg.sender, value, NATIVE_TRANSFER_GAS);
            return (false, value);
        } else {
            // ERC20: try transferring directly to user
            if (_tryTransfer(token, to, value)) {
                // Success: user received tokens
                return (false, 0);
            }
            // Fallback: set allowance for bridge to recover
            _tryForceApprove(token, msg.sender, value);
            return (false, value);
        }
    }
}
