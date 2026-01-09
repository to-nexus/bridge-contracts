// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeExecutorDeploy
 * @notice BridgeExecutor 배포 스크립트
 * @dev Bridge finalization 시 extraData를 처리하는 BridgeExecutor를 배포합니다.
 *
 * .env 파일에서 읽어오는 환경 변수:
 * - OWNER (required): Admin 권한을 가질 주소
 * - BRIDGE (optional): EXECUTOR_ROLE을 부여할 Bridge 주소
 *
 * 사용법:
 * 1. 기본 배포:
 *    forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 *
 * 2. 배포 + 검증:
 *    forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 *      --rpc-url <RPC_URL> \
 *      --broadcast \
 *      --verify
 *
 * 3. Bridge에 Executor 연결:
 *    forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 *      --rpc-url <RPC_URL> \
 *      --sig "setBridgeExecutor(address)" <EXECUTOR_ADDRESS> \
 *      --broadcast
 */
contract BridgeExecutorDeploy is Script {
    /**
     * @notice 기본 실행 함수 - BridgeExecutor 배포
     */
    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // 환경 변수에서 주소 읽기
        address owner = vm.envAddress("OWNER");
        address bridge = vm.envOr("BRIDGE", address(0));

        require(owner != address(0), "OWNER is required");

        console.log("\n=== Configuration ===");
        console.log("Owner (Admin):", owner);
        console.log("Bridge (Executor):", bridge);

        vm.startBroadcast();

        // BridgeExecutor 배포
        BridgeExecutor executor = new BridgeExecutor(owner, bridge);
        console.log("\nBridgeExecutor deployed:", address(executor));

        vm.stopBroadcast();

        _printSummary(address(executor), owner, bridge);
    }

    /**
     * @notice 파라미터를 직접 전달하여 배포
     * @param _owner Admin 권한을 가질 주소
     * @param _bridge EXECUTOR_ROLE을 부여할 Bridge 주소 (0이면 부여 안함)
     * @return executor 배포된 BridgeExecutor 컨트랙트
     */
    function deploy(address _owner, address _bridge) external returns (BridgeExecutor executor) {
        require(_owner != address(0), "Invalid owner address");

        console.log("Deploying BridgeExecutor with parameters:");
        console.log("  Owner:", _owner);
        console.log("  Bridge:", _bridge);

        vm.startBroadcast();

        executor = new BridgeExecutor(_owner, _bridge);

        vm.stopBroadcast();

        console.log("\nBridgeExecutor deployed at:", address(executor));
        return executor;
    }

    /**
     * @notice Bridge에 Executor 연결
     * @dev ADMIN_ROLE 권한이 필요합니다
     * @param _bridge Bridge 주소
     * @param _executor 배포된 BridgeExecutor 주소
     */
    function setBridgeExecutor(address _bridge, address _executor) external {
        require(_bridge != address(0), "Invalid bridge address");
        require(_executor != address(0), "Invalid executor address");

        console.log("Setting BridgeExecutor on Bridge:");
        console.log("  Bridge:", _bridge);
        console.log("  Executor:", _executor);

        vm.startBroadcast();

        BaseBridge(payable(_bridge)).setBridgeExecutor(IBridgeExecutor(_executor));

        vm.stopBroadcast();

        console.log("\nBridgeExecutor set successfully!");
    }

    /**
     * @notice Executor에 화이트리스트 타겟 추가
     * @param _executor BridgeExecutor 주소
     * @param _target 화이트리스트에 추가할 타겟 컨트랙트 주소
     */
    function addWhitelistTarget(address _executor, address _target) external {
        require(_executor != address(0), "Invalid executor address");
        require(_target != address(0), "Invalid target address");

        console.log("Adding whitelist target:");
        console.log("  Executor:", _executor);
        console.log("  Target:", _target);

        vm.startBroadcast();

        BridgeExecutor(payable(_executor)).addWhitelistTarget(_target);

        vm.stopBroadcast();

        console.log("\nTarget whitelisted successfully!");
    }

    /**
     * @notice Executor에 화이트리스트 메서드 추가
     * @param _executor BridgeExecutor 주소
     * @param _target 타겟 컨트랙트 주소
     * @param _methodIDs 화이트리스트에 추가할 메서드 셀렉터 배열
     */
    function addWhitelistMethods(address _executor, address _target, bytes4[] calldata _methodIDs) external {
        require(_executor != address(0), "Invalid executor address");
        require(_target != address(0), "Invalid target address");

        console.log("Adding whitelist methods:");
        console.log("  Executor:", _executor);
        console.log("  Target:", _target);
        console.log("  Method count:", _methodIDs.length);

        vm.startBroadcast();

        BridgeExecutor(payable(_executor)).addWhitelistMethods(_target, _methodIDs);

        vm.stopBroadcast();

        console.log("\nMethods whitelisted successfully!");
    }

    /**
     * @notice 타겟에 대한 메서드 체크 활성화/비활성화
     * @param _executor BridgeExecutor 주소
     * @param _target 타겟 컨트랙트 주소
     * @param _enabled 활성화 여부
     */
    function setMethodCheckEnabled(address _executor, address _target, bool _enabled) external {
        require(_executor != address(0), "Invalid executor address");
        require(_target != address(0), "Invalid target address");

        console.log("Setting method check:");
        console.log("  Executor:", _executor);
        console.log("  Target:", _target);
        console.log("  Enabled:", _enabled);

        vm.startBroadcast();

        BridgeExecutor(payable(_executor)).setMethodCheckEnabled(_target, _enabled);

        vm.stopBroadcast();

        console.log("\nMethod check setting updated!");
    }

    function _printSummary(address executor, address owner, address bridge) internal pure {
        console.log("\n=== Deployment Summary ===");
        console.log("BridgeExecutor:", executor);
        console.log("Owner (Admin):", owner);
        if (bridge != address(0)) console.log("Bridge (EXECUTOR_ROLE granted):", bridge);

        console.log("\n=== Next Steps ===");
        console.log("1. Set executor on Bridge (requires ADMIN_ROLE on Bridge):");
        console.log("   forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \\");
        console.log("     --sig 'setBridgeExecutor(address,address)' <BRIDGE> <EXECUTOR> \\");
        console.log("     --rpc-url $RPC_URL --broadcast");

        console.log("\n2. Add whitelist targets (requires ADMIN_ROLE on Executor):");
        console.log("   forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \\");
        console.log("     --sig 'addWhitelistTarget(address,address)' <EXECUTOR> <TARGET> \\");
        console.log("     --rpc-url $RPC_URL --broadcast");

        console.log("\n3. (Optional) Enable method check for a target:");
        console.log("   forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \\");
        console.log("     --sig 'setMethodCheckEnabled(address,address,bool)' <EXECUTOR> <TARGET> true \\");
        console.log("     --rpc-url $RPC_URL --broadcast");

        console.log("\n4. (Optional) Add whitelist methods:");
        console.log("   forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \\");
        console.log("     --sig 'addWhitelistMethods(address,address,bytes4[])' <EXECUTOR> <TARGET> '[0x12345678]' \\");
        console.log("     --rpc-url $RPC_URL --broadcast");
    }
}

/*
 * ===================================
 * BridgeExecutorDeploy.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # 필수 환경 변수
 * # --------------------------------------------------
 *
 * # Admin 권한을 가질 주소
 * # DEFAULT_ADMIN_ROLE이 부여되어 whitelist 관리 등 가능
 * OWNER=0x...
 *
 * # BaseBridge 프록시 주소
 * # EXECUTOR_ROLE이 부여되어 executeExtraCall 호출 가능
 * BRIDGE=0x...
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # 1. BridgeExecutor 배포:
 * # forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --broadcast
 *
 * # 2. Bridge에 Executor 연결 (ADMIN_ROLE 필요):
 * # forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "setBridgeExecutor(address)" $EXECUTOR_ADDRESS \
 * #   --broadcast
 *
 * # 3. 화이트리스트 타겟 추가:
 * # forge script script/BridgeExecutorDeploy.s.sol:BridgeExecutorDeploy \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "addWhitelistTarget(address,address)" $EXECUTOR $TARGET \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 배포 후 확인 사항
 * # --------------------------------------------------
 *
 * # 1. Executor가 Bridge에 설정되었는지 확인
 * # cast call $BRIDGE "bridgeExecutor()" --rpc-url $RPC_URL
 *
 * # 2. 타겟이 화이트리스트에 있는지 확인
 * # cast call $EXECUTOR "isWhitelistedTarget(address)" $TARGET --rpc-url $RPC_URL
 *
 * # 3. 메서드 체크가 활성화되었는지 확인
 * # cast call $EXECUTOR "isMethodCheckEnabled(address)" $TARGET --rpc-url $RPC_URL
 *
 * # 4. 메서드가 화이트리스트에 있는지 확인
 * # cast call $EXECUTOR "isWhitelistedMethod(address,bytes4)" $TARGET $METHOD_ID --rpc-url $RPC_URL
 *
 * # 5. Gas Reserve 확인
 * # cast call $EXECUTOR "postCallGasReserve()" --rpc-url $RPC_URL
 */
