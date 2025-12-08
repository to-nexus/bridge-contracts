// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeBot} from "../src/BridgeBot.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeBotDeploy
 * @notice Deployment script for BridgeBot contract
 */
contract BridgeBotDeploy is Script {
    /**
     * @notice Deploy BridgeBot contract
     * @param _bridgeAddress Address of the bridge contract
     * @param _ownerAddress Address of the owner/admin
     * @param _editorAddress Address of the editor (can be same as owner)
     * @param _executorAddress Address of the executor (can be same as owner)
     * @param _adminDelay Delay for admin role changes (in seconds, e.g., 86400 = 1 day)
     * @return bridgeBot The deployed BridgeBot contract
     */
    function deploy(
        address _bridgeAddress,
        address _ownerAddress,
        address _editorAddress,
        address _executorAddress,
        uint48 _adminDelay
    ) public returns (BridgeBot) {
        require(_bridgeAddress != address(0), "Invalid bridge address");
        require(_ownerAddress != address(0), "Invalid owner address");
        require(_editorAddress != address(0), "Invalid editor address");
        require(_executorAddress != address(0), "Invalid executor address");

        console.log("Deploying BridgeBot with parameters:");
        console.log("  Bridge:", _bridgeAddress);
        console.log("  Owner:", _ownerAddress);
        console.log("  Editor:", _editorAddress);
        console.log("  Executor:", _executorAddress);
        console.log("  Admin Delay:", _adminDelay);

        vm.startBroadcast();

        BridgeBot bridgeBot =
            new BridgeBot(_bridgeAddress, _ownerAddress, _editorAddress, _executorAddress, _adminDelay);

        console.log("\n=== Deployment Complete ===");
        console.log("BridgeBot deployed at:", address(bridgeBot));
        console.log("\nNext steps:");
        console.log("1. Verify contract on block explorer");
        console.log("2. Fund the contract with tokens/native currency");
        console.log("3. Set bridge configuration using setConfig()");

        vm.stopBroadcast();

        return bridgeBot;
    }
}
