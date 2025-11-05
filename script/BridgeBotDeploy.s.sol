// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeBot} from "../src/BridgeBot.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeBotExample
 * @notice BridgeBot 사용 예제 스크립트
 * @dev 이 스크립트는 BridgeBot을 배포하고 설정하는 예제를 보여줍니다
 */
contract BridgeBotExample is Script {
    // 상수 정의
    address public constant NATIVE_TOKEN = address(1);

    // 실제 주소들 (테스트넷/메인넷에 맞게 수정 필요)
    address public constant BRIDGE_ADDRESS = address(0x192FFaCAf26A53b439C33F3d2C166ceB0A26bf2d);
    address public constant CROSS_TOKEN = address(0x692a1dEC4C3703e7E932264a093ddFE92b95a678);
    address public constant OWNER = address(0xB9032595eC0465f43de9CF68c1E230888a5d16b6);
    address public constant RECIPIENT = address(0x8A088A4c67a374807B4592579EE92bdf8494E4fE);

    // 체인 ID
    uint public constant CROSS_CHAIN_ID = 612044;
    uint public constant BSC_CHAIN_ID = 56;

    // 시간 간격
    uint public constant DAILY_INTERVAL = 86400; // 24시간
    uint public constant HOURLY_INTERVAL = 3600; // 1시간

    BridgeBot public bridgeBot;
    BaseBridge public bridge;

    function setUp() public {
        bridge = BaseBridge(payable(BRIDGE_ADDRESS));
    }

    /**
     * @notice BridgeBot 배포 및 초기 설정
     */
    function deployAndSetup() public {
        vm.startBroadcast();

        // BridgeBot 배포 (with 1 day admin delay)
        bridgeBot = new BridgeBot(BRIDGE_ADDRESS, OWNER, OWNER, 1 days);

        console.log("BridgeBot deployed at:", address(bridgeBot));

        // 초기 자금 입금 (예제)
        if (CROSS_TOKEN != NATIVE_TOKEN) {
            // ERC20 토큰을 BridgeBot에 전송
            IERC20(CROSS_TOKEN).transfer(address(bridgeBot), 10000 ether);
            console.log("Transferred 10000 CROSS tokens to BridgeBot");
        }

        // 네이티브 토큰도 전송 (필요시)
        if (address(bridgeBot).balance < 1 ether) {
            payable(address(bridgeBot)).transfer(1 ether);
            console.log("Transferred 1 ETH to BridgeBot");
        }

        vm.stopBroadcast();
    }

    /**
     * @notice 브릿지 설정 추가 예제
     */
    function addBridgeConfigs() public {
        vm.startBroadcast();

        // 설정 1: 매일 CROSS 토큰을 Cross 체인으로 브릿지
        uint configId1 = bridgeBot.addBridgeConfig(
            CROSS_TOKEN, // 토큰 주소
            RECIPIENT, // 받는 주소
            CROSS_CHAIN_ID, // 대상 체인 ID
            DAILY_INTERVAL // 24시간 간격
        );

        console.log("Daily CROSS bridge config added with ID:", configId1);

        // 설정 2: 매시간 네이티브 토큰을 Cross 체인으로 브릿지
        uint configId2 = bridgeBot.addBridgeConfig(
            NATIVE_TOKEN, // 네이티브 토큰
            RECIPIENT, // 받는 주소
            CROSS_CHAIN_ID, // 대상 체인 ID
            HOURLY_INTERVAL // 1시간 간격
        );

        console.log("Hourly native token bridge config added with ID:", configId2);

        // 설정 3: 매일 CROSS 토큰을 BSC 체인으로 브릿지
        uint configId3 = bridgeBot.addBridgeConfig(
            CROSS_TOKEN, // 토큰 주소
            RECIPIENT, // 받는 주소
            BSC_CHAIN_ID, // BSC 체인 ID
            DAILY_INTERVAL // 24시간 간격
        );

        console.log("Daily CROSS to BSC bridge config added with ID:", configId3);

        vm.stopBroadcast();
    }

    /**
     * @notice 브릿지 실행 예제
     */
    function executeBridges() public {
        vm.startBroadcast();

        uint totalConfigs = bridgeBot.nextConfigId();
        console.log("Total configs:", totalConfigs);

        uint executedCount = 0;

        // 모든 설정을 순회하며 실행 가능한 것들만 실행
        for (uint configId = 0; configId < totalConfigs; configId++) {
            (bool canExecute,) = bridgeBot.canExecuteBridge(configId);

            if (canExecute) {
                // 개별 설정 정보 조회
                BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);

                console.log("Executing config ID:", configId);
                console.log("Token:", config.tokenAddress);
                console.log("To Chain:", config.toChainID);

                // 동적으로 amount 결정 (예제: CROSS 토큰은 3500 ether, 네이티브는 0.1 ether)
                uint amount;
                if (config.tokenAddress == CROSS_TOKEN) amount = 3500 ether;
                else if (config.tokenAddress == NATIVE_TOKEN) amount = 0.1 ether;
                else amount = 1000 ether; // 기본값

                console.log("Amount:", amount);

                // 브릿지 실행
                try bridgeBot.executeBridge(configId, amount) {
                    console.log("Bridge executed successfully");
                    executedCount++;
                } catch Error(string memory reason) {
                    console.log("Bridge execution failed:", reason);
                }
            }
        }

        console.log("Executed bridges:", executedCount);

        vm.stopBroadcast();
    }

    /**
     * @notice 설정 업데이트 예제
     */
    function updateConfig() public {
        vm.startBroadcast();

        uint configId = 0; // 업데이트할 설정 ID

        // 설정 업데이트
        bridgeBot.updateBridgeConfig(
            configId,
            CROSS_TOKEN, // 새로운 토큰 주소
            RECIPIENT, // 새로운 받는 주소
            CROSS_CHAIN_ID, // 새로운 체인 ID
            DAILY_INTERVAL * 2, // 새로운 간격 (48시간)
            0 // lastExecuted (0 = use current block.timestamp)
        );

        console.log("Config", configId, "updated");

        vm.stopBroadcast();
    }

    /**
     * @notice 설정 비활성화/활성화
     */
    function toggleConfig() public {
        vm.startBroadcast();

        uint configId = 0; // 토글할 설정 ID

        // 설정 비활성화
        bridgeBot.toggleBridgeConfig(configId, false);
        console.log("Config", configId, "disabled");

        // 설정 다시 활성화
        bridgeBot.toggleBridgeConfig(configId, true);
        console.log("Config", configId, "enabled");

        vm.stopBroadcast();
    }

    /**
     * @notice 자금 출금 예제
     */
    function withdrawFunds() public {
        vm.startBroadcast();

        // ERC20 토큰 출금
        bridgeBot.withdrawToken(CROSS_TOKEN, OWNER, 1000 ether);
        console.log("Withdrew 1000 CROSS tokens");

        // 네이티브 토큰 출금
        bridgeBot.withdrawNative(payable(OWNER), 0.5 ether);
        console.log("Withdrew 0.5 ETH");

        // 모든 토큰 출금
        bridgeBot.withdrawAllTokens(CROSS_TOKEN, OWNER);
        console.log("Withdrew all CROSS tokens");

        // 모든 네이티브 토큰 출금
        bridgeBot.withdrawAllNative(payable(OWNER));
        console.log("Withdrew all native tokens");

        vm.stopBroadcast();
    }

    /**
     * @notice 설정 조회 예제
     */
    function queryConfigs() public view {
        // 특정 설정 조회
        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(0);

        console.log("Config 0:");
        console.log("  Token:", config.tokenAddress);
        console.log("  Recipient:", config.recipient);
        console.log("  To Chain:", config.toChainID);
        console.log("  Interval:", config.interval);
        console.log("  Last Executed:", config.lastExecuted);
        console.log("  Enabled:", config.enabled);

        // 실행 가능 여부 확인
        (bool canExecute, uint nextAvailableTime) = bridgeBot.canExecuteBridge(0);
        console.log("  Can Execute:", canExecute);
        console.log("  Next Available Time:", nextAvailableTime);
    }

    /**
     * @notice 메인 실행 함수
     */
    function run() public {
        // 사용할 함수들을 주석 해제하여 실행

        // 1. BridgeBot 배포 및 설정
        // deployAndSetup();

        // 2. 브릿지 설정 추가
        // addBridgeConfigs();

        // 3. 브릿지 실행
        executeBridges(); // 개별 실행

        // 4. 설정 관리 (오너만 가능)
        // updateConfig();
        // toggleConfig();

        // 5. 자금 출금 (오너만 가능)
        // withdrawFunds();

        // 6. 설정 조회
        // queryConfigs();
    }
}
