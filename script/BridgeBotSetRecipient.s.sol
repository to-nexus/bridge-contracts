// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeBot} from "../src/BridgeBot.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeBotSetRecipient
 * @notice Script to change ONLY the recipient of a deployed BridgeBot's config.
 * @dev BridgeBot.setConfig() overwrites the entire config and forces `enabled = true`.
 *      To change only the recipient safely, this script reads the current config via
 *      getConfig(), preserves every other field (token / toChainID / interval /
 *      lastExecuted), and restores the `enabled` flag afterwards if it was disabled.
 *      Requires the broadcast wallet to hold EDITOR_ROLE (or be the owner).
 */
contract BridgeBotSetRecipient is Script {
    /**
     * @notice Default entry point. Reads parameters from environment variables.
     * @dev Env vars: BRIDGE_BOT_ADDR (address), NEW_RECIPIENT (address)
     */
    function run() external {
        address bridgeBot = vm.envAddress("BRIDGE_BOT_ADDR");
        address newRecipient = vm.envAddress("NEW_RECIPIENT");
        _setRecipient(bridgeBot, newRecipient);
    }

    /**
     * @notice Explicit entry point for `--sig "setRecipient(address,address)"`.
     * @param _bridgeBot Address of the deployed BridgeBot contract
     * @param _newRecipient New recipient address to set on the config
     */
    function setRecipient(address _bridgeBot, address _newRecipient) public {
        _setRecipient(_bridgeBot, _newRecipient);
    }

    /**
     * @notice Explicit entry point for `--sig "setInterval(address,uint256)"`.
     * @param _bridgeBot Address of the deployed BridgeBot contract
     * @param _newInterval New execution interval in seconds (must be > 0)
     */
    function setInterval(address _bridgeBot, uint _newInterval) public {
        _setInterval(_bridgeBot, _newInterval);
    }

    /**
     * @notice Read current config, swap recipient only, broadcast, then verify.
     * @param _bridgeBot Address of the deployed BridgeBot contract
     * @param _newRecipient New recipient address to set on the config
     */
    function _setRecipient(address _bridgeBot, address _newRecipient) internal {
        require(_bridgeBot != address(0), "Invalid BridgeBot address");
        require(_newRecipient != address(0), "Invalid recipient address");

        BridgeBot bot = BridgeBot(payable(_bridgeBot));

        // Read existing config so we preserve every field except recipient.
        BridgeBot.BridgeConfig memory cur = bot.getConfig();
        require(cur.tokenAddress != address(0), "Config not initialized");

        console.log("Updating BridgeBot recipient:");
        console.log("  BridgeBot:     ", _bridgeBot);
        console.log("  Old recipient: ", cur.recipient);
        console.log("  New recipient: ", _newRecipient);
        console.log("  Preserved Token:      ", cur.tokenAddress);
        console.log("  Preserved ChainID:    ", cur.toChainID);
        console.log("  Preserved Interval:   ", cur.interval);
        console.log("  Preserved LastExecuted:", cur.lastExecuted);
        console.log("  Preserved Enabled:    ", cur.enabled);

        if (cur.recipient == _newRecipient) {
            console.log("\nRecipient already set to the target address. Nothing to do.");
            return;
        }

        vm.startBroadcast();

        // setConfig overwrites the whole config and forces enabled = true.
        bot.setConfig(cur.tokenAddress, _newRecipient, cur.toChainID, cur.interval, cur.lastExecuted);

        // Restore the original enabled state if it was disabled before.
        if (!cur.enabled) bot.setEnabled(false);

        vm.stopBroadcast();

        // Verify the on-chain result.
        BridgeBot.BridgeConfig memory updated = bot.getConfig();
        require(updated.recipient == _newRecipient, "Recipient update failed");
        require(updated.tokenAddress == cur.tokenAddress, "Token changed unexpectedly");
        require(updated.toChainID == cur.toChainID, "ChainID changed unexpectedly");
        require(updated.interval == cur.interval, "Interval changed unexpectedly");
        require(updated.enabled == cur.enabled, "Enabled state not preserved");

        console.log("\n=== Recipient Updated Successfully ===");
        console.log("  Recipient: ", updated.recipient);
        console.log("  Enabled:   ", updated.enabled);
    }

    /**
     * @notice Read current config, change interval only, broadcast, then verify.
     * @param _bridgeBot Address of the deployed BridgeBot contract
     * @param _newInterval New execution interval in seconds (must be > 0)
     */
    function _setInterval(address _bridgeBot, uint _newInterval) internal {
        require(_bridgeBot != address(0), "Invalid BridgeBot address");
        require(_newInterval > 0, "Invalid interval");

        BridgeBot bot = BridgeBot(payable(_bridgeBot));

        // Read existing config so we preserve every field except interval.
        BridgeBot.BridgeConfig memory cur = bot.getConfig();
        require(cur.tokenAddress != address(0), "Config not initialized");

        console.log("Updating BridgeBot interval:");
        console.log("  BridgeBot:    ", _bridgeBot);
        console.log("  Old interval: ", cur.interval);
        console.log("  New interval: ", _newInterval);
        console.log("  Preserved Token:      ", cur.tokenAddress);
        console.log("  Preserved Recipient:  ", cur.recipient);
        console.log("  Preserved ChainID:    ", cur.toChainID);
        console.log("  Preserved LastExecuted:", cur.lastExecuted);
        console.log("  Preserved Enabled:    ", cur.enabled);

        if (cur.interval == _newInterval) {
            console.log("\nInterval already set to the target value. Nothing to do.");
            return;
        }

        vm.startBroadcast();

        // setConfig overwrites the whole config and forces enabled = true.
        bot.setConfig(cur.tokenAddress, cur.recipient, cur.toChainID, _newInterval, cur.lastExecuted);

        // Restore the original enabled state if it was disabled before.
        if (!cur.enabled) bot.setEnabled(false);

        vm.stopBroadcast();

        // Verify the on-chain result.
        BridgeBot.BridgeConfig memory updated = bot.getConfig();
        require(updated.interval == _newInterval, "Interval update failed");
        require(updated.tokenAddress == cur.tokenAddress, "Token changed unexpectedly");
        require(updated.recipient == cur.recipient, "Recipient changed unexpectedly");
        require(updated.toChainID == cur.toChainID, "ChainID changed unexpectedly");
        require(updated.enabled == cur.enabled, "Enabled state not preserved");

        console.log("\n=== Interval Updated Successfully ===");
        console.log("  Interval: ", updated.interval);
        console.log("  Enabled:  ", updated.enabled);
    }
}
