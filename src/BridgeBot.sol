// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {Const} from "./lib/Const.sol";

/**
 * @title BridgeBot
 * @notice Bot contract for periodic token bridging
 * @dev Provides automated token bridging functionality based on configured intervals
 * - Periodic bridge execution
 * - Admin-only configuration management and withdrawals
 * - Role-based bridge execution capability
 */
contract BridgeBot is AccessControlDefaultAdminRules, ReentrancyGuard {
    using SafeERC20 for IERC20;

    error BridgeBotCanNotZeroAddress();
    error BridgeBotCanNotZeroValue();
    error BridgeBotNotTimeYet(uint nextAvailableTime);
    error BridgeBotInsufficientBalance(uint required, uint available);
    error BridgeBotBridgeFailed();

    /**
     * @notice Bridge configuration structure
     * @param tokenAddress Token address to bridge
     * @param recipient Recipient wallet address
     * @param toChainID Target chain ID
     * @param interval Bridge execution interval in seconds
     * @param lastExecuted Last execution timestamp
     * @param enabled Activation status
     */
    struct BridgeConfig {
        address tokenAddress;
        address recipient;
        uint toChainID;
        uint interval;
        uint lastExecuted;
        bool enabled;
    }

    /**
     * @notice Bridge execution event
     * @param tokenAddress Token address
     * @param amount Bridged amount
     * @param recipient Recipient address
     * @param toChainID Target chain ID
     * @param executor Transaction executor
     * @param timestamp Execution timestamp
     */
    event BridgeExecuted(
        address indexed tokenAddress,
        uint amount,
        address indexed recipient,
        uint toChainID,
        address executor,
        uint timestamp
    );

    /**
     * @notice Bridge configuration updated event
     */
    event ConfigSet(BridgeConfig config);

    /**
     * @notice Bridge configuration enabled/disabled event
     */
    event ConfigToggled(bool enabled);

    /**
     * @notice Token/Native withdrawal event
     * @param token Token address (Const.NATIVE_TOKEN for native coin)
     * @param to Recipient address
     * @param amount Withdrawal amount
     */
    event Withdrawn(address indexed token, address indexed to, uint amount);

    /// @dev Bridge contract
    BaseBridge public immutable bridge;

    /// @dev Bridge configuration
    BridgeConfig public config;

    /**
     * @notice Contract constructor
     * @param _bridge Bridge contract address
     * @param _owner Initial admin address
     * @param _editor Initial editor address
     * @param _executor Initial executor address
     * @param _adminDelay Delay for admin role changes (in seconds)
     */
    constructor(address _bridge, address _owner, address _editor, address _executor, uint48 _adminDelay)
        AccessControlDefaultAdminRules(_adminDelay, _owner)
    {
        require(_bridge != address(0), BridgeBotCanNotZeroAddress());
        require(_editor != address(0), BridgeBotCanNotZeroAddress());
        require(_executor != address(0), BridgeBotCanNotZeroAddress());
        bridge = BaseBridge(payable(_bridge));

        // Grant roles
        _grantRole(Const.EDITOR_ROLE, _owner);
        _grantRole(Const.EDITOR_ROLE, _editor);
        _grantRole(Const.EXECUTOR_ROLE, _owner);
        _grantRole(Const.EXECUTOR_ROLE, _executor);
    }

    /**
     * @notice Set bridge configuration
     * @param tokenAddress Token address
     * @param recipient Recipient address
     * @param toChainID Target chain ID
     * @param interval Execution interval in seconds
     * @param lastExecuted Last execution timestamp (0 = allow immediate execution)
     */
    function setConfig(address tokenAddress, address recipient, uint toChainID, uint interval, uint lastExecuted)
        external
        onlyRole(Const.EDITOR_ROLE)
    {
        require(tokenAddress != address(0), BridgeBotCanNotZeroAddress());
        require(recipient != address(0), BridgeBotCanNotZeroAddress());
        require(interval > 0, BridgeBotCanNotZeroValue());

        config.tokenAddress = tokenAddress;
        config.recipient = recipient;
        config.toChainID = toChainID;
        config.interval = interval;
        config.lastExecuted = lastExecuted;
        config.enabled = true;

        emit ConfigSet(config);
    }

    /**
     * @notice Enable/disable bridge configuration
     * @param enabled Enable status
     */
    function setEnabled(bool enabled) external onlyRole(Const.EDITOR_ROLE) {
        config.enabled = enabled;
        emit ConfigToggled(enabled);
    }

    /**
     * @notice Check if bridge can be executed
     * @return canExecute Whether execution is possible
     * @return nextAvailableTime Next available execution time
     */
    function canExecuteBridge() public view returns (bool canExecute, uint nextAvailableTime) {
        if (!config.enabled || config.tokenAddress == address(0)) return (false, 0);

        // First execution case (lastExecuted == 0)
        if (config.lastExecuted == 0) return (true, 0);

        uint currentPeriod = block.timestamp - (block.timestamp % config.interval);
        uint lastExecutedPeriod = config.lastExecuted - (config.lastExecuted % config.interval);

        canExecute = currentPeriod > lastExecutedPeriod;
        nextAvailableTime = canExecute ? 0 : lastExecutedPeriod + config.interval;
    }

    /**
     * @notice Execute token bridge (role restricted)
     * @dev Automatically detects token type (native or ERC20) based on config
     * @param amount Amount to bridge
     */
    function executeBridge(uint amount) external onlyRole(Const.EXECUTOR_ROLE) nonReentrant {
        if (config.tokenAddress == Const.NATIVE_TOKEN) _executeBridgeNativeInternal(amount);
        else _executeBridgeERC20Internal(amount);
    }

    /**
     * @notice Internal ERC20 bridge execution function
     * @param amount Amount to bridge
     */
    function _executeBridgeERC20Internal(uint amount) internal {
        require(amount > 0, BridgeBotCanNotZeroValue());
        require(config.enabled && config.tokenAddress != address(0), "Config not available");

        (bool canExecute, uint nextAvailableTime) = canExecuteBridge();
        require(canExecute, BridgeBotNotTimeYet(nextAvailableTime));

        // Check balance
        uint balance = IERC20(config.tokenAddress).balanceOf(address(this));

        // Calculate fees
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint gasFee, uint exFee) =
            bridgeVerifier.calculateFee(config.toChainID, IERC20(config.tokenAddress), amount);

        uint totalRequired = amount + gasFee + exFee;
        require(balance >= totalRequired, BridgeBotInsufficientBalance(totalRequired, balance));
        require(amount >= minimumValue, "Amount below minimum");

        // Token approval - check and approve max if needed
        IERC20 token = IERC20(config.tokenAddress);
        uint currentAllowance = token.allowance(address(this), address(bridge));
        if (currentAllowance < totalRequired) token.forceApprove(address(bridge), type(uint).max);

        // Execute bridge
        bool success = bridge.bridgeToken(
            config.toChainID, IERC20(config.tokenAddress), config.recipient, amount, gasFee, exFee, ""
        );

        require(success, BridgeBotBridgeFailed());

        // Update last execution time
        config.lastExecuted = block.timestamp;

        emit BridgeExecuted(
            config.tokenAddress, amount, config.recipient, config.toChainID, msg.sender, block.timestamp
        );
    }

    /**
     * @notice Internal native token bridge execution function
     * @param amount Amount to bridge
     */
    function _executeBridgeNativeInternal(uint amount) internal {
        require(amount > 0, BridgeBotCanNotZeroValue());
        require(config.enabled, "Config not available");

        (bool canExecute, uint nextAvailableTime) = canExecuteBridge();
        require(canExecute, BridgeBotNotTimeYet(nextAvailableTime));

        // Check balance
        uint balance = address(this).balance;

        // Calculate fees
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint gasFee, uint exFee) =
            bridgeVerifier.calculateFee(config.toChainID, IERC20(Const.NATIVE_TOKEN), amount);

        uint totalRequired = amount + gasFee + exFee;
        require(balance >= totalRequired, BridgeBotInsufficientBalance(totalRequired, balance));
        require(amount >= minimumValue, "Amount below minimum");

        // Execute bridge with native token
        bool success = bridge.bridgeToken{value: totalRequired}(
            config.toChainID, IERC20(Const.NATIVE_TOKEN), config.recipient, amount, gasFee, exFee, ""
        );

        require(success, BridgeBotBridgeFailed());

        // Update last execution time
        config.lastExecuted = block.timestamp;

        emit BridgeExecuted(Const.NATIVE_TOKEN, amount, config.recipient, config.toChainID, msg.sender, block.timestamp);
    }

    /**
     * @notice Withdraw ERC20 tokens (admin only)
     * @param token Token address
     * @param to Recipient address
     * @param amount Amount to withdraw
     */
    function withdrawToken(address token, address to, uint amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(token != address(0), BridgeBotCanNotZeroAddress());
        require(to != address(0), BridgeBotCanNotZeroAddress());
        require(amount > 0, BridgeBotCanNotZeroValue());

        IERC20(token).safeTransfer(to, amount);
        emit Withdrawn(token, to, amount);
    }

    /**
     * @notice Withdraw native tokens (admin only)
     * @param to Recipient address
     * @param amount Amount to withdraw
     */
    function withdrawNative(address payable to, uint amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(to != address(0), BridgeBotCanNotZeroAddress());
        require(amount > 0, BridgeBotCanNotZeroValue());
        require(address(this).balance >= amount, "Insufficient balance");

        (bool success,) = to.call{value: amount}("");
        require(success, "Transfer failed");

        emit Withdrawn(Const.NATIVE_TOKEN, to, amount);
    }

    /**
     * @notice Withdraw all ERC20 tokens (admin only)
     * @param token Token address
     * @param to Recipient address
     */
    function withdrawAllTokens(address token, address to) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(token != address(0), BridgeBotCanNotZeroAddress());
        require(to != address(0), BridgeBotCanNotZeroAddress());

        uint balance = IERC20(token).balanceOf(address(this));
        if (balance > 0) {
            IERC20(token).safeTransfer(to, balance);
            emit Withdrawn(token, to, balance);
        }
    }

    /**
     * @notice Withdraw all native tokens (admin only)
     * @param to Recipient address
     */
    function withdrawAllNative(address payable to) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(to != address(0), BridgeBotCanNotZeroAddress());

        uint balance = address(this).balance;
        if (balance > 0) {
            (bool success,) = to.call{value: balance}("");
            require(success, "Transfer failed");
            emit Withdrawn(Const.NATIVE_TOKEN, to, balance);
        }
    }

    /**
     * @notice Get bridge configuration
     * @return Bridge configuration
     */
    function getConfig() external view returns (BridgeConfig memory) {
        return config;
    }

    /**
     * @notice Receive native tokens
     */
    receive() external payable {}

    /**
     * @notice Fallback function
     */
    fallback() external payable {}
}
