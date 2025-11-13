// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeBot} from "../src/BridgeBot.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeBotSetConfig
 * @notice Script to configure BridgeBot
 */
contract BridgeBotSetConfig is Script {
    /**
     * @notice Set bridge configuration
     * @param _bridgeBotAddress Address of the BridgeBot contract
     * @param _tokenAddress Token address to bridge
     * @param _recipient Recipient address on target chain
     * @param _toChainID Target chain ID
     * @param _interval Execution interval in seconds
     * @param _lastExecuted Last execution timestamp (0 = allow immediate execution)
     */
    function setConfig(
        address _bridgeBotAddress,
        address _tokenAddress,
        address _recipient,
        uint _toChainID,
        uint _interval,
        uint _lastExecuted
    ) public {
        require(_bridgeBotAddress != address(0), "Invalid BridgeBot address");
        require(_tokenAddress != address(0), "Invalid token address");
        require(_recipient != address(0), "Invalid recipient address");
        require(_interval > 0, "Invalid interval");

        console.log("Setting BridgeBot configuration:");
        console.log("  BridgeBot:", _bridgeBotAddress);
        console.log("  Token:", _tokenAddress);
        console.log("  Recipient:", _recipient);
        console.log("  Target Chain ID:", _toChainID);
        console.log("  Interval:", _interval, "seconds");
        console.log("  Last Executed:", _lastExecuted);

        BridgeBot bridgeBot = BridgeBot(payable(_bridgeBotAddress));

        vm.startBroadcast();

        bridgeBot.setConfig(_tokenAddress, _recipient, _toChainID, _interval, _lastExecuted);

        vm.stopBroadcast();

        console.log("\n=== Configuration Set Successfully ===");
        console.log("Config details:");
        console.log("  Token:", _tokenAddress);
        console.log("  Recipient:", _recipient);
        console.log("  Target Chain:", _toChainID);
        console.log("  Interval:", _interval, "seconds");
        console.log("  Interval (hours):", _interval / 3600);
        console.log("  Enabled: true");

        if (_lastExecuted == 0) console.log("  Status: Ready for immediate execution");
        else console.log("  Last Executed:", _lastExecuted);
    }
}
