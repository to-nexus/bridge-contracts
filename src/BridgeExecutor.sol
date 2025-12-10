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
 *
 * Security:
 * - Only callable by authorized bridge contracts
 * - Target contracts must be whitelisted (to be implemented)
 */
contract BridgeExecutor is AccessControl, ReentrancyGuardTransient, IBridgeExecutor {
    using SafeERC20 for IERC20;

    error BETargetNotWhitelisted();
    error BEInvalidAddress();
    error BEInvalidValue();
    error BEInsufficientBalance();
    error BEAlreadyWhitelisted();
    error BENotWhitelisted();
    error BEInvalidRecipient();
    error BEFailedToReturnNative();

    /**
     * @notice Emitted when an extra call execution is performed
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param toToken Token being bridged
     * @param to Original recipient address
     * @param value Amount of tokens
     * @param targetContract Contract that was called
     * @param methodID Method ID of the called contract
     * @param success Whether the execution was successful
     * @param returnData Call response data (success: return value, failure: revert reason)
     */
    event ExtraCallExecuted(
        uint indexed fromChainID,
        uint indexed index,
        IERC20 indexed toToken,
        address to,
        uint value,
        address targetContract,
        bytes4 methodID,
        bool success,
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

    /// @dev Mapping of whitelisted target contracts
    mapping(address => bool) private _whitelistedTargets;

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

    function executeExtraCall(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable onlyRole(Const.EXECUTOR_ROLE) nonReentrant returns (bool success) {
        // Skip if value is 0
        if (value == 0) return false;

        address targetContract;
        bytes4 methodID;
        assembly {
            targetContract := shr(96, calldataload(extraData.offset))
            methodID := calldataload(add(extraData.offset, 20))
        }
        require(_whitelistedTargets[targetContract], BETargetNotWhitelisted());

        bool isNative = address(toToken) == Const.NATIVE_TOKEN;

        // Validate balance and approve if needed
        if (isNative) {
            require(msg.value == value, BEInvalidValue());
        } else {
            require(toToken.balanceOf(address(this)) >= value, BEInsufficientBalance());
            toToken.forceApprove(targetContract, value);
        }

        // Reserve minimum gas for post-extracall logic
        uint remainingGas = gasleft() < 200000 ? 0 : gasleft() - 200000;

        // Call target contract
        bytes memory returnData;
        (success, returnData) = targetContract.call{gas: remainingGas, value: isNative ? value : 0}(extraData[20:]);

        // Clear approval (always for ERC20)
        if (!isNative) toToken.forceApprove(targetContract, 0);

        // Handle failure: return tokens to bridge
        if (!success) {
            if (isNative) {
                (bool sent,) = msg.sender.call{value: value}("");
                require(sent, BEFailedToReturnNative());
            } else {
                // Approve bridge to burn (for wrapped) or transfer back is handled by bridge
                toToken.forceApprove(msg.sender, value);
            }
        }

        emit ExtraCallExecuted(fromChainID, index, toToken, to, value, targetContract, methodID, success, returnData);
        return success;
    }

    /**
     * @notice Add a target contract to the whitelist
     * @dev Only callable by admin
     * @param target Target contract address to whitelist
     */
    function addWhitelistTarget(address target) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(target != address(0), BEInvalidAddress());
        require(!_whitelistedTargets[target], BEAlreadyWhitelisted());
        _whitelistedTargets[target] = true;
        emit TargetWhitelisted(target);
    }

    /**
     * @notice Remove a target contract from the whitelist
     * @dev Only callable by admin
     * @param target Target contract address to remove
     */
    function removeWhitelistTarget(address target) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_whitelistedTargets[target], BENotWhitelisted());
        _whitelistedTargets[target] = false;
        emit TargetRemovedFromWhitelist(target);
    }

    /**
     * @notice Check if a target contract is whitelisted
     * @param target Target contract address to check
     * @return True if target is whitelisted
     */
    function isWhitelistedTarget(address target) external view returns (bool) {
        return _whitelistedTargets[target];
    }

    /**
     * @notice Emergency function to recover stuck tokens
     * @dev Only callable by admin
     * @param token Token to recover (address(0) for native token)
     * @param amount Amount to recover
     * @param recipient Address to send recovered tokens to
     */
    function recoverToken(address token, uint amount, address recipient) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(recipient != address(0) && token != address(0), BEInvalidRecipient());
        if (token == Const.NATIVE_TOKEN) payable(recipient).transfer(amount);
        else IERC20(token).safeTransfer(recipient, amount);
    }

    /**
     * @notice Allows contract to receive native tokens
     */
    receive() external payable {}
}
