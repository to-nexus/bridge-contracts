// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {SwapBridgeRouter} from "../src/SwapBridgeRouter.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title SwapBridgeRouterDeploy
 * @notice SwapBridgeRouter 배포 스크립트
 * @dev Uniswap V3 swap과 bridge를 결합한 SwapBridgeRouter를 배포합니다.
 *      BridgeVerifier는 Bridge에서 동적으로 조회하고, WETH9는 SwapRouter에서 조회합니다.
 *
 * .env 파일에서 읽어오는 환경 변수:
 * - SWAP_ROUTER (required, 필수): Uniswap V3 SwapRouter 주소
 * - QUOTER (required, 필수): Uniswap V3 Quoter 주소
 * - BRIDGE (required, 필수): BaseBridge 프록시 주소
 *
 * 사용법:
 * 1. 기본 배포:
 *    forge script script/SwapBridgeRouterDeploy.s.sol:SwapBridgeRouterDeploy \
 *      --rpc-url <RPC_URL> \
 *      --broadcast
 *
 * 2. 배포 + 검증:
 *    forge script script/SwapBridgeRouterDeploy.s.sol:SwapBridgeRouterDeploy \
 *      --rpc-url <RPC_URL> \
 *      --broadcast \
 *      --verify
 */
contract SwapBridgeRouterDeploy is Script {
    /**
     * @notice 기본 실행 함수 - SwapBridgeRouter 배포
     */
    function run() external {
        address deployer = msg.sender;

        console.log("Deployer:", deployer);
        console.log("Chain ID:", block.chainid);

        // 환경 변수에서 주소 읽기 (모두 필수)
        address swapRouterAddr = vm.envAddress("SWAP_ROUTER");
        address quoterAddr = vm.envAddress("QUOTER");
        address bridgeAddr = vm.envAddress("BRIDGE");

        require(swapRouterAddr != address(0), "SWAP_ROUTER is required");
        require(quoterAddr != address(0), "QUOTER is required");
        require(bridgeAddr != address(0), "BRIDGE is required");

        console.log("\n=== Configuration ===");
        console.log("Uniswap V3 SwapRouter:", swapRouterAddr);
        console.log("Uniswap V3 Quoter:", quoterAddr);
        console.log("Bridge:", bridgeAddr);

        vm.startBroadcast();

        // SwapBridgeRouter 배포
        SwapBridgeRouter swapBridgeRouter = new SwapBridgeRouter(swapRouterAddr, quoterAddr, bridgeAddr);
        console.log("\nSwapBridgeRouter deployed:", address(swapBridgeRouter));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("SwapBridgeRouter:", address(swapBridgeRouter));
        console.log("SwapRouter:", swapRouterAddr);
        console.log("Quoter:", quoterAddr);
        console.log("Bridge:", bridgeAddr);
        console.log("WETH9 (from SwapRouter):", address(swapBridgeRouter.WETH9()));

        console.log("\n=== Next Steps ===");
        console.log("1. Verify contract on block explorer");
        console.log("2. Users can now use swap+bridge functions:");
        console.log("   - swapExactInputSingleAndBridge()");
        console.log("   - swapExactInputAndBridge()");
        console.log("   - swapExactOutputSingleAndBridge()");
        console.log("   - swapExactOutputAndBridge()");
        console.log("   - swapExactInputSingleETHAndBridge()");
        console.log("   - swapExactInputETHAndBridge()");
        console.log("   - swapExactOutputSingleETHAndBridge()");
        console.log("   - swapExactOutputETHAndBridge()");
        console.log("3. View functions for fee estimation:");
        console.log("   - getExpectedBridgeAmount()");
        console.log("   - calculateBridgeFees()");
        console.log("4. Quote functions for swap+bridge estimation:");
        console.log("   - getAmountSwapBridgeOut()");
        console.log("   - getAmountSwapBridgeOutMultihop()");
        console.log("   - getAmountSwapBridgeIn()");
        console.log("   - getAmountSwapBridgeInMultihop()");
    }

    /**
     * @notice 파라미터를 직접 전달하여 배포
     * @param _swapRouter Uniswap V3 SwapRouter 주소
     * @param _quoter Uniswap V3 Quoter 주소
     * @param _bridge BaseBridge 프록시 주소
     * @return swapBridgeRouter 배포된 SwapBridgeRouter 컨트랙트
     */
    function deploy(address _swapRouter, address _quoter, address _bridge)
        external
        returns (SwapBridgeRouter swapBridgeRouter)
    {
        require(_swapRouter != address(0), "Invalid swapRouter address");
        require(_quoter != address(0), "Invalid quoter address");
        require(_bridge != address(0), "Invalid bridge address");

        console.log("Deploying SwapBridgeRouter with parameters:");
        console.log("  SwapRouter:", _swapRouter);
        console.log("  Quoter:", _quoter);
        console.log("  Bridge:", _bridge);

        vm.startBroadcast();

        swapBridgeRouter = new SwapBridgeRouter(_swapRouter, _quoter, _bridge);

        console.log("\n=== Deployment Complete ===");
        console.log("SwapBridgeRouter deployed at:", address(swapBridgeRouter));
        console.log("Quoter:", address(swapBridgeRouter.quoter()));
        console.log("WETH9 (from SwapRouter):", address(swapBridgeRouter.WETH9()));

        vm.stopBroadcast();

        return swapBridgeRouter;
    }
}

/*
 * ===================================
 * SwapBridgeRouterDeploy.env.example
 * ===================================
 *
 * # --------------------------------------------------
 * # 필수 환경 변수
 * # --------------------------------------------------
 *
 * # Uniswap V3 SwapRouter 주소
 * # 각 체인별 공식 주소:
 * # - Ethereum Mainnet: 0xE592427A0AEce92De3Edee1F18E0157C05861564
 * # - Polygon: 0xE592427A0AEce92De3Edee1F18E0157C05861564
 * # - Arbitrum: 0xE592427A0AEce92De3Edee1F18E0157C05861564
 * # - Optimism: 0xE592427A0AEce92De3Edee1F18E0157C05861564
 * # - BSC: (Uniswap V3 not officially deployed, use PancakeSwap V3 or fork)
 * SWAP_ROUTER=0x...
 *
 * # Uniswap V3 Quoter 주소
 * # 각 체인별 공식 주소:
 * # - Ethereum Mainnet: 0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6
 * # - Polygon: 0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6
 * # - Arbitrum: 0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6
 * # - Optimism: 0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6
 * QUOTER=0x...
 *
 * # BaseBridge 프록시 주소
 * # 기존 배포된 Bridge 컨트랙트 주소를 사용합니다.
 * BRIDGE=0x...
 *
 * # --------------------------------------------------
 * # 자동으로 조회되는 값들
 * # --------------------------------------------------
 *
 * # BridgeVerifier: bridge.bridgeVerifier()에서 동적으로 조회
 * # WETH9: swapRouter.WETH9()에서 조회 (PeripheryImmutableState)
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # 1. 기본 배포 (환경 변수 사용):
 * # forge script script/SwapBridgeRouterDeploy.s.sol:SwapBridgeRouterDeploy \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --broadcast
 *
 * # 2. 배포 + Etherscan 검증:
 * # forge script script/SwapBridgeRouterDeploy.s.sol:SwapBridgeRouterDeploy \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --broadcast \
 * #   --verify \
 * #   --etherscan-api-key $ETHERSCAN_API_KEY
 *
 * # 3. 시뮬레이션 (배포 없이 테스트):
 * # forge script script/SwapBridgeRouterDeploy.s.sol:SwapBridgeRouterDeploy \
 * #   --rpc-url $RPC_URL
 *
 * # --------------------------------------------------
 * # 배포 후 확인 사항
 * # --------------------------------------------------
 *
 * # 1. 컨트랙트가 올바르게 배포되었는지 확인
 * # cast call $SWAP_BRIDGE_ROUTER "swapRouter()" --rpc-url $RPC_URL
 * # cast call $SWAP_BRIDGE_ROUTER "quoter()" --rpc-url $RPC_URL
 * # cast call $SWAP_BRIDGE_ROUTER "bridge()" --rpc-url $RPC_URL
 * # cast call $SWAP_BRIDGE_ROUTER "WETH9()" --rpc-url $RPC_URL
 *
 * # 2. Fee 계산 테스트
 * # cast call $SWAP_BRIDGE_ROUTER \
 * #   "getExpectedBridgeAmount(uint256,address,uint256)" \
 * #   $TO_CHAIN_ID $TOKEN_ADDRESS $AMOUNT \
 * #   --rpc-url $RPC_URL
 *
 * # 3. Quote 테스트 (swap+bridge 예상값 조회)
 * # cast call $SWAP_BRIDGE_ROUTER \
 * #   "getAmountSwapBridgeOut(uint256,address,address,uint24,uint256)" \
 * #   $TO_CHAIN_ID $TOKEN_IN $TOKEN_OUT $FEE $AMOUNT_IN \
 * #   --rpc-url $RPC_URL
 */
