// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../src/SwapBridgeRouter.sol";
import "../Bridge.t.sol";
import "../token/TestToken.sol";
import "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

contract SwapBridgeTest is BridgeTest {
    event SwapBridgeInitiated(
        uint indexed toChainID, uint indexed index, address indexed from, address to, address[] path, uint[] amounts
    );
    event BridgeInitiated(
        uint indexed toChainID,
        uint indexed index,
        IERC20 fromToken,
        IERC20 toToken,
        address indexed from,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes extraData,
        uint time
    );

    IPancakeRouter02 public router;
    SwapBridgeRouter public swapBridgeRouter;
    TestToken public tokenUSD;
    address public WETH;

    address[] public wethToCrossPath;
    address[] public tokenUsdToCrossPath;

    function setUp() public override {
        bscNodeURL = vm.envString("BSC_RPC_URL"); // use bsc-mainnet from env
        super.setUp();
        vm.selectFork(bscForkID);

        router = IPancakeRouter02(address(0x10ED43C718714eb63d5aA57B78B54704E256024E)); // BSC Mainnet Pancake Swap Router
        swapBridgeRouter =
            new SwapBridgeRouter(OWNER, address(bridgeBSC), address(router), address(cross), CROSS_CHAIN_ID);
        WETH = router.WETH();

        vm.startPrank(OWNER);
        // Deploy test token
        tokenUSD = new TestToken("USD Token", "USD", 18);

        // Mint test token
        tokenUSD.mint(OWNER, 1000000000 ether);
        vm.deal(OWNER, 1000000000 ether);

        // Create swap pool
        tokenUSD.approve(address(router), type(uint).max);
        cross.approve(address(router), type(uint).max);

        router.addLiquidity(
            address(tokenUSD),
            address(cross),
            10000000 ether,
            100000000 ether,
            9000000 ether,
            90000000 ether,
            OWNER,
            block.timestamp + 300
        );

        router.addLiquidityETH{value: 10000 ether}(
            address(cross), 60000000 ether, 54000000 ether, 9000 ether, OWNER, block.timestamp + 300
        );

        wethToCrossPath = [WETH, address(cross)];
        tokenUsdToCrossPath = [address(tokenUSD), address(cross)];

        swapBridgeRouter.setTokenPath(address(tokenUSD), tokenUsdToCrossPath);
        swapBridgeRouter.setTokenPath(WETH, wethToCrossPath);

        vm.stopPrank();
    }

    function testSwapExactTokensForTokensBridge() public {
        _swapExactTokensForTokensBridge(100 ether);
    }

    function testSwapTokensForExactTokensBridge() public {
        _swapTokensForExactTokensBridge(100 ether);
    }

    function testSwapExactETHForTokensBridge() public {
        _swapExactETHForTokensBridge(100 ether);
    }

    function testSwapTokensForExactETHBridge() public {
        _swapTokensForExactETHBridge(100 ether);
    }

    function testSwapExactTokensForETHBridge() public {
        _swapExactTokensForETHBridge(1000 ether);
    }

    function testSwapETHForExactTokensBridge() public {
        _swapETHForExactTokensBridge(100 ether);
    }

    function testSwapExactTokensForCrossTokensBridge() public {
        _swapExactTokensForCrossTokensBridge(100 ether);
    }

    function testSwapTokensForExactCrossTokensBridge() public {
        _swapTokensForExactCrossTokensBridge(100 ether);
    }

    function testSwapExactETHForCrossTokensBridge() public {
        _swapExactETHForCrossTokensBridge(100 ether);
    }

    function testSwapETHForExactCrossTokensBridge() public {
        _swapETHForExactCrossTokensBridge(100 ether);
    }

    function testFuzzSwapExactTokensForTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapExactTokensForTokensBridge(x);
    }

    function testFuzzSwapTokensForExactTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapTokensForExactTokensBridge(x);
    }

    function testFuzzSwapExactETHForTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapExactETHForTokensBridge(x);
    }

    function testFuzzSwapTokensForExactETHBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapTokensForExactETHBridge(x);
    }

    function testFuzzSwapExactTokensForETHBridge(uint x) public {
        vm.assume(x > 1000 ether && x < 10000 ether);
        _swapExactTokensForETHBridge(x);
    }

    function testFuzzSwapETHForExactTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapETHForExactTokensBridge(x);
    }

    function testFuzzSwapExactTokensForCrossTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapExactTokensForCrossTokensBridge(x);
    }

    function testFuzzSwapTokensForExactCrossTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapTokensForExactCrossTokensBridge(x);
    }

    function testFuzzSwapExactETHForCrossTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapExactETHForCrossTokensBridge(x);
    }

    function testFuzzSwapETHForExactCrossTokensBridge(uint x) public {
        vm.assume(x > 100 ether && x < 1000 ether);
        _swapETHForExactCrossTokensBridge(x);
    }

    function _swapExactTokensForTokensBridge(uint amount) private {
        vm.assume(tokenUSD.balanceOf(OWNER) >= amount);

        // token 충전
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amount);

        address[] memory path = new address[](2);
        path[0] = address(tokenUSD);
        path[1] = address(cross);

        // getSwapBridgeOut 테스트
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount, path);
        assertTrue(amounts[0] == amount);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeExactTokensForTokens(
                CROSS_CHAIN_ID, USER, amount, amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapTokensForExactTokensBridge(uint bridgeValue) private {
        // Prepare for swap
        address[] memory path = new address[](2);
        path[0] = address(tokenUSD);
        path[1] = address(cross);

        (uint[] memory amounts, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        vm.assume(tokenUSD.balanceOf(OWNER) >= amounts[0]);
        // Charge token
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);

            // Swap bridge
            swapBridgeRouter.swapBridgeTokensForExactTokens(
                CROSS_CHAIN_ID, USER, bridgeValue, amounts[0], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapExactETHForTokensBridge(uint amount) private {
        vm.assume(OWNER.balance >= amount);
        // Charge token
        vm.prank(OWNER);
        payable(USER).transfer(amount);

        address[] memory path = new address[](2);
        path[0] = WETH;
        path[1] = address(cross);

        // Test getSwapBridgeOut
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount, path);
        assertTrue(amounts[0] == amount);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeExactETHForTokens{value: amount}(
                CROSS_CHAIN_ID, USER, amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapTokensForExactETHBridge(uint bridgeValue) private {
        // Prepare for swap
        address[] memory path = new address[](2);
        path[0] = address(cross);
        path[1] = WETH;

        (uint[] memory amounts, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        vm.assume(cross.balanceOf(OWNER) >= amounts[0]);
        // Charge token
        vm.prank(OWNER);
        cross.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            cross.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                NATIVE_TOKEN,
                weth,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeTokensForExactETH(
                CROSS_CHAIN_ID, USER, bridgeValue, amounts[0], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapExactTokensForETHBridge(uint amount) private {
        vm.assume(cross.balanceOf(OWNER) >= amount);
        // Charge token
        vm.prank(OWNER);
        cross.transfer(USER, amount);

        address[] memory path = new address[](2);
        path[0] = address(cross);
        path[1] = WETH;

        // Test getSwapBridgeOut
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeOut(CROSS_CHAIN_ID, amount, path);
        assertTrue(amounts[0] == amount);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            cross.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                NATIVE_TOKEN,
                weth,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeExactTokensForETH(
                CROSS_CHAIN_ID, USER, amounts[0], amounts[1], networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapETHForExactTokensBridge(uint bridgeValue) private {
        // Prepare for swap
        address[] memory path = new address[](2);
        path[0] = WETH;
        path[1] = address(cross);

        (uint[] memory amounts, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeIn(CROSS_CHAIN_ID, bridgeValue, path);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        vm.assume(OWNER.balance >= amounts[0]);
        // Charge token
        vm.prank(OWNER);
        payable(USER).transfer(amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, path, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeETHForExactTokens{value: amounts[0]}(
                CROSS_CHAIN_ID, USER, bridgeValue, networkFee, exFee, path, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapExactTokensForCrossTokensBridge(uint amount) private {
        vm.assume(tokenUSD.balanceOf(OWNER) >= amount);
        // Charge token
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amount);

        // Test getSwapBridgeOut
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeOutCross(address(tokenUSD), amount);
        assertTrue(amounts[0] == amount);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, tokenUsdToCrossPath, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeExactTokensForCrossTokens(
                address(tokenUSD), USER, amount, amounts[1], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapTokensForExactCrossTokensBridge(uint bridgeValue) private {
        (uint[] memory amounts, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeInCross(address(tokenUSD), bridgeValue);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        vm.assume(tokenUSD.balanceOf(OWNER) >= amounts[0]);
        // Charge token
        vm.prank(OWNER);
        tokenUSD.transfer(USER, amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);
            tokenUSD.approve(address(swapBridgeRouter), type(uint).max);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, tokenUsdToCrossPath, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeTokensForExactCrossTokens(
                address(tokenUSD), USER, bridgeValue, amounts[0], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapExactETHForCrossTokensBridge(uint amount) private {
        vm.assume(OWNER.balance >= amount);
        // Charge token
        vm.prank(OWNER);
        payable(USER).transfer(amount);

        // Test getSwapBridgeOut
        (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee,) =
            swapBridgeRouter.getSwapBridgeOutCross(WETH, amount);
        assertTrue(amounts[0] == amount);
        assertTrue(amounts[1] > 0);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, wethToCrossPath, amounts);

            // Swap bridge
            swapBridgeRouter.swapBridgeExactETHForCrossTokens{value: amount}(
                WETH, USER, amounts[1], networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }

    function _swapETHForExactCrossTokensBridge(uint bridgeValue) private {
        (uint[] memory amounts, uint networkFee, uint exFee,) = swapBridgeRouter.getSwapBridgeInCross(WETH, bridgeValue);
        assertTrue(amounts[0] > 0);
        assertTrue(amounts[1] == bridgeValue);

        vm.assume(OWNER.balance >= amounts[0]);
        // Charge token
        vm.prank(OWNER);
        payable(USER).transfer(amounts[0]);

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        {
            vm.startPrank(USER);

            vm.expectEmit(true, true, true, true, address(bridgeBSC), 1);
            emit BridgeInitiated(
                CROSS_CHAIN_ID,
                expectedIndex,
                cross,
                NATIVE_TOKEN,
                address(swapBridgeRouter),
                USER,
                bridgeValue,
                networkFee,
                exFee,
                new bytes(0),
                block.timestamp
            );
            vm.expectEmit(true, true, true, true, address(swapBridgeRouter), 1);
            emit SwapBridgeInitiated(CROSS_CHAIN_ID, expectedIndex, USER, USER, wethToCrossPath, amounts);
            // Swap bridge
            swapBridgeRouter.swapBridgeETHForExactCrossTokens{value: amounts[0]}(
                WETH, USER, bridgeValue, networkFee, exFee, uint(block.timestamp + 300)
            );
            vm.stopPrank();
        }
    }
}
