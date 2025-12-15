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

    error BETargetNotWhitelisted();
    error BEMethodNotWhitelisted();
    error BEInvalidAddress();
    error BEInvalidValue();
    error BEInsufficientBalance();
    error BEAlreadyWhitelisted();
    error BENotWhitelisted();
    error BEInvalidRecipient();
    error BEFailedToReturnNative();
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
        address, // to (unused, kept for interface compatibility)
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
            if (msg.value != 0) _trySendNative(msg.sender, msg.value);
            return (false, 0);
        }

        bool isNative = address(toToken) == Const.NATIVE_TOKEN;
        Target storage target = _targets[targetContract];

        // Reject non-whitelisted targets (no revert)
        if (!target.isWhitelisted) {
            if (isNative && msg.value != 0) _trySendNative(msg.sender, msg.value);
            return (false, 0);
        }

        // Optional: methodID check per target (no revert)
        if (target.methodCheckEnabled && !_whitelistedMethods[targetContract][methodID]) {
            if (isNative && msg.value != 0) _trySendNative(msg.sender, msg.value);
            return (false, 0);
        }

        // Record balance before call
        uint balBefore;
        if (isNative) {
            // For native: balance before = current balance - msg.value (what we just received)
            balBefore = address(this).balance - msg.value;
            if (msg.value != value) {
                // Refund whatever was sent (best-effort) and fail
                if (msg.value != 0) _trySendNative(msg.sender, msg.value);
                return (false, 0);
            }
        } else {
            // Any unexpected msg.value is refunded to the bridge (best-effort)
            if (msg.value != 0) _trySendNative(msg.sender, msg.value);

            (bool okBal, uint bal) = _tryBalanceOf(toToken, address(this));
            if (!okBal || bal < value) return (false, 0);
            balBefore = bal;

            if (!_tryForceApprove(toToken, targetContract, value)) return (false, 0);
        }

        // Reserve minimum gas for post-extracall logic
        uint gasReserve = _postCallGasReserve;
        uint remainingGas = gasleft() < gasReserve ? 0 : gasleft() - gasReserve;

        // Call target contract
        bytes memory returnData;
        (success, returnData) = targetContract.call{gas: remainingGas, value: isNative ? value : 0}(extraData[20:]);

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
        uint consumed = balBefore < balAfter ? 0 : balBefore - balAfter;
        // For native, we need to account for the msg.value we received
        if (isNative) {
            // We received msg.value, so actual consumed = msg.value - (balAfter - balBefore)
            // If balAfter > balBefore, target returned some
            consumed = balAfter >= balBefore ? value - (balAfter - balBefore) : value;
        }
        remaining = value > consumed ? value - consumed : 0;

        // Return remaining to bridge
        if (remaining > 0) {
            if (isNative) _trySendNative(msg.sender, remaining);
            else _tryForceApprove(toToken, msg.sender, remaining);
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

    function _trySendNative(address recipient, uint amount) private returns (bool ok) {
        if (amount == 0) return true;
        (ok,) = payable(recipient).call{value: amount}("");
    }

    function _tryBalanceOf(IERC20 token, address account) private view returns (bool ok, uint bal) {
        bytes memory returndata;
        (ok, returndata) = address(token).staticcall(abi.encodeCall(IERC20.balanceOf, (account)));
        if (!ok || returndata.length < 32) return (false, 0);
        bal = abi.decode(returndata, (uint));
        return (true, bal);
    }

    function _callOptionalReturnBool(IERC20 token, bytes memory data) private returns (bool) {
        (bool ok, bytes memory returndata) = address(token).call(data);
        if (!ok) return false;

        // Tokens may return no data, assume success if target is a contract
        if (returndata.length == 0) return address(token).code.length != 0;

        if (returndata.length >= 32) return abi.decode(returndata, (bool));
        return false;
    }

    function _tryApprove(IERC20 token, address spender, uint amount) private returns (bool) {
        return _callOptionalReturnBool(token, abi.encodeCall(IERC20.approve, (spender, amount)));
    }

    function _tryForceApprove(IERC20 token, address spender, uint amount) private returns (bool) {
        // Try direct approve first
        if (_tryApprove(token, spender, amount)) return true;
        // Some tokens require setting to 0 before changing allowance
        if (!_tryApprove(token, spender, 0)) return false;
        return _tryApprove(token, spender, amount);
    }
}
