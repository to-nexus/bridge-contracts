// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeBot} from "../src/BridgeBot.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeBotDeploy
 * @notice BridgeBot 배포 스크립트
 * @dev 함수 파라미터를 통해 설정값을 받아 BridgeBot을 배포합니다
 *
 * 사용 방법:
 * forge script script/BridgeBotDeploy.s.sol:BridgeBotDeploy \
 *   --sig "deploy(address,address,address,address,uint48)" \
 *   <BRIDGE_ADDRESS> <OWNER_ADDRESS> <EDITOR_ADDRESS> <EXECUTOR_ADDRESS> <ADMIN_DELAY> \
 *   --rpc-url $RPC_URL --broadcast --verify
 *
 * 예시:
 * forge script script/BridgeBotDeploy.s.sol:BridgeBotDeploy \
 *   --sig "deploy(address,address,address,address,uint48)" \
 *   0x192FFaCAf26A53b439C33F3d2C166ceB0A26bf2d \
 *   0xB9032595eC0465f43de9CF68c1E230888a5d16b6 \
 *   0xB9032595eC0465f43de9CF68c1E230888a5d16b6 \
 *   0xB9032595eC0465f43de9CF68c1E230888a5d16b6 \
 *   86400 \
 *   --rpc-url $RPC_URL --broadcast
 */
contract BridgeBotDeploy is Script {
    BridgeBot public bridgeBot;

    /**
     * @notice BridgeBot 배포
     * @param _bridgeAddress BaseBridge 컨트랙트 주소
     * @param _ownerAddress Admin/Owner 주소
     * @param _editorAddress Editor 주소
     * @param _executorAddress Executor 주소
     * @param _adminDelay Admin role 변경 지연 시간 (초 단위)
     */
    function deploy(
        address _bridgeAddress,
        address _ownerAddress,
        address _editorAddress,
        address _executorAddress,
        uint48 _adminDelay
    ) public {
        require(_bridgeAddress != address(0), "Invalid bridge address");
        require(_ownerAddress != address(0), "Invalid owner address");
        require(_editorAddress != address(0), "Invalid editor address");
        require(_executorAddress != address(0), "Invalid executor address");

        console.log("========================================");
        console.log("Deployment Configuration:");
        console.log("Bridge Address:", _bridgeAddress);
        console.log("Owner Address:", _ownerAddress);
        console.log("Editor Address:", _editorAddress);
        console.log("Executor Address:", _executorAddress);
        console.log("Admin Delay:", _adminDelay, "seconds");
        console.log("========================================\n");

        vm.startBroadcast();

        // BridgeBot 배포
        bridgeBot = new BridgeBot(_bridgeAddress, _ownerAddress, _editorAddress, _executorAddress, _adminDelay);

        vm.stopBroadcast();

        console.log("\n========================================");
        console.log("BridgeBot deployed successfully!");
        console.log("========================================");
        console.log("BridgeBot Address:", address(bridgeBot));
        console.log("Bridge:", _bridgeAddress);
        console.log("Owner:", _ownerAddress);
        console.log("Editor:", _editorAddress);
        console.log("Executor:", _executorAddress);
        console.log("Admin Delay:", _adminDelay, "seconds");
        console.log("========================================\n");
    }
}
