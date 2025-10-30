// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";

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
     * @param configId Configuration ID
     * @param tokenAddress Token address
     * @param amount Bridged amount
     * @param recipient Recipient address
     * @param toChainID Target chain ID
     * @param executor Transaction executor
     * @param timestamp Execution timestamp
     */
    event BridgeExecuted(
        uint indexed configId,
        address indexed tokenAddress,
        uint amount,
        address indexed recipient,
        uint toChainID,
        address executor,
        uint timestamp
    );

    /**
     * @notice Bridge configuration added event
     */
    event BridgeConfigAdded(uint indexed configId, BridgeConfig config);

    /**
     * @notice Bridge configuration updated event
     */
    event BridgeConfigUpdated(uint indexed configId, BridgeConfig config);

    /**
     * @notice Bridge configuration enabled/disabled event
     */
    event BridgeConfigToggled(uint indexed configId, bool enabled);

    /**
     * @notice Token withdrawal event
     */
    event TokenWithdrawn(address indexed token, address indexed to, uint amount);

    /**
     * @notice Native token withdrawal event
     */
    event NativeWithdrawn(address indexed to, uint amount);

    /// @dev Native token address constant
    address public constant NATIVE_TOKEN = address(1);

    /// @dev Role for executing bridge operations
    bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");

    /// @dev Bridge contract
    BaseBridge public immutable bridge;

    /// @dev Bridge configurations
    mapping(uint => BridgeConfig) public bridgeConfigs;

    /// @dev Next configuration ID
    uint public nextConfigId;

    /**
     * @notice Contract constructor
     * @param _bridge Bridge contract address
     * @param _owner Initial admin address
     * @param _executor Initial executor address
     * @param _adminDelay Delay for admin role changes (in seconds)
     */
    constructor(address _bridge, address _owner, address _executor, uint48 _adminDelay)
        AccessControlDefaultAdminRules(_adminDelay, _owner)
    {
        require(_bridge != address(0), BridgeBotCanNotZeroAddress());
        require(_executor != address(0), BridgeBotCanNotZeroAddress());
        bridge = BaseBridge(payable(_bridge));

        // Set up executor role
        _grantRole(EXECUTOR_ROLE, _owner);
        _grantRole(EXECUTOR_ROLE, _executor);
    }

    /**
     * @notice Add bridge configuration
     * @param tokenAddress Token address
     * @param recipient Recipient address
     * @param toChainID Target chain ID
     * @param interval Execution interval in seconds
     * @return configId Configuration ID
     */
    function addBridgeConfig(address tokenAddress, address recipient, uint toChainID, uint interval)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (uint configId)
    {
        require(tokenAddress != address(0), BridgeBotCanNotZeroAddress());
        require(recipient != address(0), BridgeBotCanNotZeroAddress());
        require(interval > 0, BridgeBotCanNotZeroValue());

        configId = nextConfigId++;

        bridgeConfigs[configId] = BridgeConfig({
            tokenAddress: tokenAddress,
            recipient: recipient,
            toChainID: toChainID,
            interval: interval,
            lastExecuted: 0,
            enabled: true
        });

        emit BridgeConfigAdded(configId, bridgeConfigs[configId]);
    }

    /**
     * @notice Update bridge configuration
     * @param configId Configuration ID
     * @param tokenAddress Token address
     * @param recipient Recipient address
     * @param toChainID Target chain ID
     * @param interval Execution interval in seconds
     */
    function updateBridgeConfig(uint configId, address tokenAddress, address recipient, uint toChainID, uint interval)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        require(bridgeConfigs[configId].tokenAddress != address(0), "Config not exists");
        require(tokenAddress != address(0), BridgeBotCanNotZeroAddress());
        require(recipient != address(0), BridgeBotCanNotZeroAddress());
        require(interval > 0, BridgeBotCanNotZeroValue());

        bridgeConfigs[configId].tokenAddress = tokenAddress;
        bridgeConfigs[configId].recipient = recipient;
        bridgeConfigs[configId].toChainID = toChainID;
        bridgeConfigs[configId].interval = interval;

        emit BridgeConfigUpdated(configId, bridgeConfigs[configId]);
    }

    /**
     * @notice Enable/disable bridge configuration
     * @param configId Configuration ID
     * @param enabled Enable status
     */
    function toggleBridgeConfig(uint configId, bool enabled) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(bridgeConfigs[configId].tokenAddress != address(0), "Config not exists");

        bridgeConfigs[configId].enabled = enabled;
        emit BridgeConfigToggled(configId, enabled);
    }

    /**
     * @notice Check if bridge can be executed
     * @param configId Configuration ID
     * @return canExecute Whether execution is possible
     * @return nextAvailableTime Next available execution time
     */
    function canExecuteBridge(uint configId) public view returns (bool canExecute, uint nextAvailableTime) {
        BridgeConfig memory config = bridgeConfigs[configId];

        if (!config.enabled || config.tokenAddress == address(0)) return (false, 0);

        // First execution case (lastExecuted == 0)
        if (config.lastExecuted == 0) return (true, 0);

        uint currentPeriod = block.timestamp - (block.timestamp % config.interval);
        uint lastExecutedPeriod = config.lastExecuted - (config.lastExecuted % config.interval);

        canExecute = currentPeriod > lastExecutedPeriod;
        nextAvailableTime = canExecute ? 0 : lastExecutedPeriod + config.interval;
    }

    /**
     * @notice Execute bridge (role restricted)
     * @param configId Configuration ID
     * @param amount Amount to bridge
     */
    function executeBridge(uint configId, uint amount) external onlyRole(EXECUTOR_ROLE) nonReentrant {
        _executeBridgeInternal(configId, amount);
    }

    /**
     * @notice Internal bridge execution function
     * @param configId Configuration ID
     * @param amount Amount to bridge
     */
    function _executeBridgeInternal(uint configId, uint amount) internal {
        require(amount > 0, BridgeBotCanNotZeroValue());

        BridgeConfig storage config = bridgeConfigs[configId];
        require(config.enabled && config.tokenAddress != address(0), "Config not available");

        (bool canExecute, uint nextAvailableTime) = canExecuteBridge(configId);
        require(canExecute, BridgeBotNotTimeYet(nextAvailableTime));

        // Check balance
        uint balance;
        if (config.tokenAddress == NATIVE_TOKEN) balance = address(this).balance;
        else balance = IERC20(config.tokenAddress).balanceOf(address(this));

        // Calculate fees
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint gasFee, uint exFee) =
            bridgeVerifier.calculateFee(config.toChainID, IERC20(config.tokenAddress), amount);

        uint totalRequired = amount + gasFee + exFee;
        require(balance >= totalRequired, BridgeBotInsufficientBalance(totalRequired, balance));
        require(amount >= minimumValue, "Amount below minimum");

        // Token approval (for ERC20)
        if (config.tokenAddress != NATIVE_TOKEN) {
            IERC20 token = IERC20(config.tokenAddress);
            if (token.allowance(address(this), address(bridge)) < totalRequired) {
                token.approve(address(bridge), type(uint).max);
            }
        }

        // Execute bridge
        bool success;
        if (config.tokenAddress == NATIVE_TOKEN) {
            success = bridge.bridgeToken{value: totalRequired}(
                config.toChainID, IERC20(config.tokenAddress), config.recipient, amount, gasFee, exFee, ""
            );
        } else {
            success = bridge.bridgeToken(
                config.toChainID, IERC20(config.tokenAddress), config.recipient, amount, gasFee, exFee, ""
            );
        }

        require(success, BridgeBotBridgeFailed());

        // Update last execution time
        config.lastExecuted = block.timestamp;

        emit BridgeExecuted(
            configId, config.tokenAddress, amount, config.recipient, config.toChainID, msg.sender, block.timestamp
        );
    }

    /**
     * @notice Execute multiple bridge configurations in batch
     * @param configIds Array of configuration IDs
     * @param amounts Array of bridge amounts corresponding to each configuration
     */
    function executeBridgeBatch(uint[] calldata configIds, uint[] calldata amounts)
        external
        onlyRole(EXECUTOR_ROLE)
        nonReentrant
    {
        require(configIds.length == amounts.length, "Array length mismatch");

        for (uint i = 0; i < configIds.length; i++) {
            (bool canExecute,) = canExecuteBridge(configIds[i]);
            if (canExecute && amounts[i] > 0) _executeBridgeInternal(configIds[i], amounts[i]);
        }
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
        emit TokenWithdrawn(token, to, amount);
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

        emit NativeWithdrawn(to, amount);
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
            emit TokenWithdrawn(token, to, balance);
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
            emit NativeWithdrawn(to, balance);
        }
    }

    /**
     * @notice Get bridge configuration
     * @param configId Configuration ID
     * @return config Bridge configuration
     */
    function getBridgeConfig(uint configId) external view returns (BridgeConfig memory config) {
        return bridgeConfigs[configId];
    }

    /**
     * @notice Get executable bridge configurations
     * @param maxConfigs Maximum number of configurations to retrieve
     * @return executableConfigs Array of executable configuration IDs
     */
    function getExecutableConfigs(uint maxConfigs) external view returns (uint[] memory executableConfigs) {
        uint[] memory temp = new uint[](maxConfigs);
        uint count = 0;

        for (uint i = 0; i < nextConfigId && count < maxConfigs; i++) {
            (bool canExecute,) = canExecuteBridge(i);
            if (canExecute) {
                temp[count] = i;
                count++;
            }
        }

        executableConfigs = new uint[](count);
        for (uint i = 0; i < count; i++) {
            executableConfigs[i] = temp[i];
        }
    }

    /**
     * @notice Grant executor role to an address (default admin only)
     * @param account Address to grant executor role
     */
    function grantExecutorRole(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(account != address(0), BridgeBotCanNotZeroAddress());
        grantRole(EXECUTOR_ROLE, account);
    }

    /**
     * @notice Revoke executor role from an address (default admin only)
     * @param account Address to revoke executor role
     */
    function revokeExecutorRole(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(EXECUTOR_ROLE, account);
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
