// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {SwapBridgeRouterCL} from "../src/SwapBridgeRouterCL.sol";
import {Script, console} from "forge-std/Script.sol";

interface IPeripheryState {
    function factory() external view returns (address);
    function WETH9() external view returns (address);
}

interface IBridgeLike {
    function bridgeVerifier() external view returns (address);
}

/**
 * @title SwapBridgeRouterCLDeploy
 * @notice SwapBridgeRouterCL 배포 스크립트 (Slipstream CL 용 신규 배포, 기존 SwapBridgeRouter 는 변경하지 않음)
 * @dev 생성자 인자만 있는 non-upgradeable 컨트랙트 (proxy / initializer / init data 없음, owner 없음).
 *      배포 전 체크: router.factory() == quoter.factory(), router.WETH9() == quoter.WETH9(), bridge 에 code 존재,
 *      bridge.bridgeVerifier() != 0.
 *
 * 환경 변수 (모두 필수):
 * - CL_SWAP_ROUTER : Slipstream SwapRouter   (testnet 612044: 0x5F6B5E31E329700176d531333E6B9fc1eC7BA107)
 * - CL_QUOTER      : Slipstream QuoterV2     (testnet 612044: 0x0Ab6d60565158F791CA08eFc6F95333F0f10851b)
 * - BRIDGE         : 허브 Bridge 프록시      (testnet 612044: 0xb81d6e000000000000000000000000000000C0de,
 *                    기존 SwapBridgeRouter 0xc9B2...CF5B 의 bridge() 값)
 *
 * 사용법 (시뮬레이션; 실제 배포는 --broadcast 와 서명자 지정이 필요하며 승인 후에만 실행):
 *   CL_SWAP_ROUTER=... CL_QUOTER=... BRIDGE=... \
 *   forge script script/SwapBridgeRouterCLDeploy.s.sol:SwapBridgeRouterCLDeploy --rpc-url <RPC_URL>
 *
 * 배포 후 확인:
 *   cast call $SBR_CL "swapRouter()(address)"; cast call $SBR_CL "quoter()(address)"
 *   cast call $SBR_CL "bridge()(address)";     cast call $SBR_CL "WETH9()(address)"
 *   cast call $SBR_CL "getAmountSwapBridgeOut(uint256,address,address,int24,uint256)" $TO_CHAIN_ID $TOKEN_IN $TOKEN_OUT $TICK_SPACING $AMOUNT_IN
 */
contract SwapBridgeRouterCLDeploy is Script {
    function run() external returns (SwapBridgeRouterCL) {
        return _deploy(vm.envAddress("CL_SWAP_ROUTER"), vm.envAddress("CL_QUOTER"), vm.envAddress("BRIDGE"));
    }

    /// @notice 파라미터를 직접 전달하여 배포
    function deploy(address swapRouter, address quoter, address bridge) external returns (SwapBridgeRouterCL) {
        return _deploy(swapRouter, quoter, bridge);
    }

    /// @notice 배포 전 체크만 수행 (broadcast 없음)
    function check(address swapRouter, address quoter, address bridge) public view {
        require(swapRouter != address(0) && quoter != address(0) && bridge != address(0), "zero address");
        require(swapRouter.code.length > 0, "CL_SWAP_ROUTER has no code");
        require(quoter.code.length > 0, "CL_QUOTER has no code");
        require(bridge.code.length > 0, "BRIDGE has no code");
        address factory = IPeripheryState(swapRouter).factory();
        require(factory != address(0) && factory == IPeripheryState(quoter).factory(), "router/quoter factory mismatch");
        address weth = IPeripheryState(swapRouter).WETH9();
        require(weth != address(0) && weth == IPeripheryState(quoter).WETH9(), "router/quoter WETH9 mismatch");
        require(IBridgeLike(bridge).bridgeVerifier() != address(0), "bridge has no verifier");
    }

    function _deploy(address swapRouter, address quoter, address bridge) internal returns (SwapBridgeRouterCL sbr) {
        check(swapRouter, quoter, bridge);
        console.log("Chain ID:", block.chainid);
        console.log("CL SwapRouter:", swapRouter);
        console.log("CL QuoterV2:", quoter);
        console.log("Bridge:", bridge);
        console.log("CL factory:", IPeripheryState(swapRouter).factory());

        vm.startBroadcast();
        sbr = new SwapBridgeRouterCL(swapRouter, quoter, bridge);
        vm.stopBroadcast();

        require(address(sbr.swapRouter()) == swapRouter && address(sbr.quoter()) == quoter, "immutables");
        require(address(sbr.bridge()) == bridge, "bridge immutable");
        console.log("SwapBridgeRouterCL:", address(sbr));
        console.log("WETH9 (from SwapRouter):", address(sbr.WETH9()));
        console.log("Next: set bridge-scanner swap_bridge_router to this address and move the old V3 wrapper");
        console.log("      to additional_swap_bridge_routers (in-flight X3 transfers keep being tracked)");
    }
}
