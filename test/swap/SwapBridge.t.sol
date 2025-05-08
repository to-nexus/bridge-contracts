// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../Bridge.t.sol";

import "../../src/SwapBridgeRouter.sol";
import "../token/TestToken.sol";
import "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

contract SwapBridgeTest is BridgeTest {
    event SwapBridgeInitiated(
        uint indexed toChainID, uint indexed index, address indexed from, address to, uint bridgeValue, address[] path
    );

    IPancakeRouter02 public router;
    SwapBridgeRouter public swapBridgeRouter;
    TestToken public tokenUSD;
    address public WETH;

    address[] public wethToCrossPath;
    address[] public tokenUsdToCrossPath;

    function setUp() public override {
        bscNodeURL = "https://bsc-mainnet.crosstoken.io/2272489872e4f1475ff25d57ce93b51989f933c7";
        super.setUp();
        vm.selectFork(bscForkID);

        router = IPancakeRouter02(address(0x10ED43C718714eb63d5aA57B78B54704E256024E));
        swapBridgeRouter =
            new SwapBridgeRouter(OWNER, address(bridgeBSC), address(router), address(cross), CROSS_CHAIN_ID);
        WETH = router.WETH();

        vm.startPrank(OWNER);
        // 테스트용 토큰 배포
        tokenUSD = new TestToken("USD Token", "USD", 18);

        // 테스트용 토큰 발행
        tokenUSD.mint(OWNER, 1000000000 * 10 ** 18);
        vm.deal(OWNER, 100000 ether);

        // 스왑풀 생성
        tokenUSD.approve(address(router), type(uint).max);
        cross.approve(address(router), type(uint).max);

        router.addLiquidity(
            address(tokenUSD), // $1
            address(cross), // $0.1
            10000 * 10 ** 18, // 10000 tokenUSD
            100000 * 10 ** 18, // 100000 cross
            9000 * 10 ** 18, // 최소 9000 tokenUSD
            90000 * 10 ** 18, // 최소 90000 cross
            OWNER,
            block.timestamp + 300
        );

        router.addLiquidityETH{value: 10000 ether}( // $5000
            address(cross), // $0.1
            50000000 * 10 ** 18, // 50000000 cross
            5000000 * 10 ** 18, // 5000000 cross
            1000 * 10 ** 18, // 최소 1000 bnb
            OWNER,
            block.timestamp + 300
        );

        wethToCrossPath = [WETH, address(cross)];
        tokenUsdToCrossPath = [address(tokenUSD), address(cross)];

        swapBridgeRouter.setTokenPath(address(tokenUSD), tokenUsdToCrossPath);
        swapBridgeRouter.setTokenPath(WETH, wethToCrossPath);

        vm.stopPrank();
    }

    function testSwapExactTokensForTokensBridge() public {
        uint amount0 = 100 ether;
        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amount0);

        address[] memory path = new address[](2);
        path[0] = address(tokenUSD);
        path[1] = address(cross);

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount0, path);
        assertTrue(amounts[0] == amount0);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapExactTokensForTokensBridge(
                CROSS_CHAIN_ID, USER, amount0, amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapTokensForExactTokensBridge() public {
        uint bridgeValue = 100 ether;

        // 스왑 준비
        address[] memory path = new address[](2);
        path[0] = address(tokenUSD);
        path[1] = address(cross);

        (uint[] memory amounts, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapTokensForExactTokensBridge(
                CROSS_CHAIN_ID, USER, bridgeValue, amounts[0], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapExactETHForTokensBridge() public {
        uint amount0 = 100 ether;
        // token 충전
        vm.prank(OWNER);
        payable(USER).transfer(amount0);

        address[] memory path = new address[](2);
        path[0] = WETH;
        path[1] = address(cross);

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount0, path);
        assertTrue(amounts[0] == amount0);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapExactETHForTokensBridge{value: amount0}(
                CROSS_CHAIN_ID, USER, amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapTokensForExactETHBridge() public {
        uint bridgeValue = 100 ether;

        // 스왑 준비
        address[] memory path = new address[](2);
        path[0] = address(cross);
        path[1] = WETH;

        (uint[] memory amounts, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        // token 충전
        vm.prank(OWNER);
        cross.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            cross.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapTokensForExactETHBridge(
                CROSS_CHAIN_ID, USER, bridgeValue, amounts[0], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapExactTokensForETHBridge() public {
        uint amount0 = 1000 ether;
        // token 충전
        vm.prank(OWNER);
        cross.transfer(USER, amount0);

        address[] memory path = new address[](2);
        path[0] = address(cross);
        path[1] = WETH;

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount0, path);
        assertTrue(amounts[0] == amount0);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            cross.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapExactTokensForETHBridge(
                CROSS_CHAIN_ID, USER, amounts[0], amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapETHForExactTokensBridge() public {
        uint bridgeValue = 100 ether;

        // 스왑 준비
        address[] memory path = new address[](2);
        path[0] = WETH;
        path[1] = address(cross);

        (uint[] memory amounts, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);
        payable(USER).transfer(amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, path);
            // 스왑 브릿지
            swapBridgeRouter.swapETHForExactTokensBridge{value: amounts[0]}(
                CROSS_CHAIN_ID, USER, bridgeValue, networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapExactTokensForCrossTokensBridge() public {
        uint amount0 = 100 ether;
        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amount0);

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeOutCross(address(tokenUSD), amount0);
        assertTrue(amounts[0] == amount0);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, tokenUsdToCrossPath);
            // 스왑 브릿지
            swapBridgeRouter.swapExactTokensForCrossTokensBridge(
                address(tokenUSD), USER, amount0, amounts[1], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapTokensForExactCrossTokensBridge() public {
        uint bridgeValue = 100 ether;

        (uint[] memory amounts, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeInCross(address(tokenUSD), bridgeValue);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, tokenUsdToCrossPath);
            // 스왑 브릿지
            swapBridgeRouter.swapTokensForExactCrossTokensBridge(
                address(tokenUSD), USER, bridgeValue, amounts[0], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapExactETHForCrossTokensBridge() public {
        uint amount0 = 100 ether;
        // token 충전
        vm.prank(OWNER);
        payable(USER).transfer(amount0);

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouter.getSwapBridgeOutCross(WETH, amount0);
        assertTrue(amounts[0] == amount0);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, wethToCrossPath);
            // 스왑 브릿지
            swapBridgeRouter.swapExactETHForCrossTokensBridge{value: amount0}(
                WETH, USER, amounts[1], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function testSwapETHForExactCrossTokensBridge() public {
        uint bridgeValue = 100 ether;

        (uint[] memory amounts, uint networkFee, uint exFee) = swapBridgeRouter.getSwapBridgeInCross(WETH, bridgeValue);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);
        payable(USER).transfer(amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, bridgeValue, wethToCrossPath);
            // 스왑 브릿지
            swapBridgeRouter.swapETHForExactCrossTokensBridge{value: amounts[0]}(
                WETH, USER, bridgeValue, networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }
}
