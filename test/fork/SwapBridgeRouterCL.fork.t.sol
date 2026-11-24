// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test} from "forge-std/Test.sol";

import {SwapBridgeRouterCL} from "../../src/SwapBridgeRouterCL.sol";
import {IBaseBridge} from "../../src/interface/IBaseBridge.sol";
import {ICLSwapRouter} from "../../src/interface/ICLSwapRouter.sol";
import {ISwapBridgeRouterCL} from "../../src/interface/ISwapBridgeRouterCL.sol";

interface ICLFactoryLike {
    function getPool(address tokenA, address tokenB, int24 tickSpacing) external view returns (address);
}

interface IPeripheryLike {
    function factory() external view returns (address);
    function WETH9() external view returns (address);
}

interface ICLPoolLike {
    function tickSpacing() external view returns (int24);
}

/**
 * @title SwapBridgeRouterCLForkTest
 * @notice Deploys SwapBridgeRouterCL inside a LOCAL fork of ONE chain testnet (612044) and swaps + bridges through the
 * real Slipstream SwapRouter / QuoterV2 / pools and the real hub bridge. Nothing is broadcast: every write happens in the
 * in-memory fork. Funds are dealt inside the fork.
 *
 * Run (skipped unless CL_FORK_RPC is set):
 *   CL_FORK_RPC=https://testnet-archive.crosstoken.io:22001 forge test --match-contract SwapBridgeRouterCLForkTest
 * Optional CL_FORK_BLOCK pins the fork block.
 */
contract SwapBridgeRouterCLForkTest is Test {
    // testnet 612044 (verified by read-only calls)
    address constant CL_ROUTER = 0x5F6B5E31E329700176d531333E6B9fc1eC7BA107;
    address constant CL_QUOTER = 0x0Ab6d60565158F791CA08eFc6F95333F0f10851b;
    address constant CL_FACTORY = 0x38C3a5c6f5A1607f0438420B44f9a197940f9A3c;
    address constant HUB_BRIDGE = 0xb81d6e000000000000000000000000000000C0de; // bridge() of the deployed V3 SwapBridgeRouter
    address constant V3_SBR = 0xc9B23a5E14747Cc2Acc3Dca7237D3174CE1BCf5B;
    address constant WCROSS = 0xda37a3B60E6b97E8CB44b460130d40331efc566b;
    address constant ONEUSD = 0x9364ea6790f6E0EcFaa5164085f2a7de34EC55Fb;
    address constant USDTX = 0x9F85c7B5D7637E18f946cc8AF9C131318c6833d9;
    address constant POOL_ONEUSD_USDTX = 0x83c082f35EAb5BF910E241DA1e745f41362a8842; // tickSpacing 200
    address constant POOL_ONEUSD_WCROSS = 0x9593E81E023F302650F500B2851A5292FA574Ad3; // tickSpacing 2000
    uint constant BSC_TESTNET = 97; // USDTx has a hub pair to BSC testnet
    bytes4 constant V3_EXACT_INPUT_SINGLE = 0x414bf389;

    SwapBridgeRouterCL sbr;
    address user = makeAddr("clForkUser");
    address recipient = makeAddr("clForkRecipient");
    bool enabled;

    function setUp() public {
        string memory rpc = vm.envOr("CL_FORK_RPC", string(""));
        if (bytes(rpc).length == 0) return;
        uint blockNumber = vm.envOr("CL_FORK_BLOCK", uint(0));
        if (blockNumber == 0) vm.createSelectFork(rpc);
        else vm.createSelectFork(rpc, blockNumber);
        require(block.chainid == 612044, "not ONE testnet");
        enabled = true;
        sbr = new SwapBridgeRouterCL(CL_ROUTER, CL_QUOTER, HUB_BRIDGE);
        vm.deal(user, 1_000 ether);
    }

    modifier onlyFork() {
        if (!enabled) {
            vm.skip(true);
            return;
        }
        _;
    }

    function _bp() internal view returns (ISwapBridgeRouterCL.BridgeParams memory) {
        return ISwapBridgeRouterCL.BridgeParams({toChainID: BSC_TESTNET, recipient: recipient, extraData: ""});
    }

    function test_fork_wiring() public onlyFork {
        assertEq(IPeripheryLike(CL_ROUTER).factory(), CL_FACTORY);
        assertEq(IPeripheryLike(CL_QUOTER).factory(), CL_FACTORY);
        assertEq(address(sbr.WETH9()), WCROSS);
        assertEq(address(sbr.bridge()), HUB_BRIDGE);
        // the hub bridge is the one the deployed V3 wrapper uses
        (bool ok, bytes memory ret) = V3_SBR.staticcall(abi.encodeWithSignature("bridge()"));
        assertTrue(ok);
        assertEq(abi.decode(ret, (address)), HUB_BRIDGE);
        assertEq(ICLFactoryLike(CL_FACTORY).getPool(ONEUSD, USDTX, 200), POOL_ONEUSD_USDTX);
        assertEq(ICLPoolLike(POOL_ONEUSD_USDTX).tickSpacing(), int24(200));
        assertEq(ICLFactoryLike(CL_FACTORY).getPool(ONEUSD, WCROSS, 2000), POOL_ONEUSD_WCROSS);
    }

    function test_fork_quote_tickSpacing() public onlyFork {
        (ISwapBridgeRouterCL.QuoteStatus st, uint swapOut, uint bv, uint nf, uint ef) =
            sbr.getAmountSwapBridgeOut(BSC_TESTNET, ONEUSD, USDTX, 200, 100e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.Success));
        assertGt(swapOut, 99e18);
        assertEq(bv + nf + ef <= swapOut, true);
        // no pool at tick spacing 100 for this pair: the real quoter reverts -> InvalidSwap
        assertEq(ICLFactoryLike(CL_FACTORY).getPool(ONEUSD, USDTX, 100), address(0));
        (st,,,,) = sbr.getAmountSwapBridgeOut(BSC_TESTNET, ONEUSD, USDTX, 100, 100e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.InvalidSwap));
        // a negative tick spacing is not a pool either
        (st,,,,) = sbr.getAmountSwapBridgeOut(BSC_TESTNET, ONEUSD, USDTX, -200, 100e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.InvalidSwap));
    }

    function test_fork_exactInputSingle_swapAndBridge() public onlyFork {
        uint amountIn = 100e18;
        deal(ONEUSD, user, amountIn);
        (, uint quoted,,,) = sbr.getAmountSwapBridgeOut(BSC_TESTNET, ONEUSD, USDTX, 200, amountIn);
        uint idx = IBaseBridge(HUB_BRIDGE).getNextInitiateIndex(BSC_TESTNET);

        vm.startPrank(user);
        IERC20(ONEUSD).approve(address(sbr), amountIn);
        ISwapBridgeRouterCL.SwapBridgeExactInputSingleParams memory p = ISwapBridgeRouterCL
            .SwapBridgeExactInputSingleParams({
            tokenIn: ONEUSD,
            tokenOut: USDTX,
            tickSpacing: 200,
            amountIn: amountIn,
            amountOutMinimum: quoted + 1,
            sqrtPriceLimitX96: 0,
            bridgeParams: _bp()
        });
        // slippage: one wei above the quote reverts on the real router
        vm.expectRevert();
        sbr.swapBridgeExactInputSingle(p, block.timestamp + 600);

        p.amountOutMinimum = quoted;
        uint out = sbr.swapBridgeExactInputSingle(p, block.timestamp + 600);
        vm.stopPrank();

        assertEq(out, quoted, "execution matches the real quoter");
        assertEq(IERC20(ONEUSD).balanceOf(user), 0);
        assertEq(IBaseBridge(HUB_BRIDGE).getNextInitiateIndex(BSC_TESTNET), idx + 1, "one bridge transfer initiated");
        assertEq(IERC20(USDTX).balanceOf(address(sbr)), 0);
        assertEq(IERC20(ONEUSD).balanceOf(address(sbr)), 0);
    }

    function test_fork_nativeMultihop_int24Path() public onlyFork {
        // WCROSS -(2000)-> ONEUSD -(200)-> USDTx, native ONE in
        bytes memory path = abi.encodePacked(WCROSS, int24(2000), ONEUSD, int24(200), USDTX);
        uint amountIn = 100 ether;
        (ISwapBridgeRouterCL.QuoteStatus st, uint quoted,,,) =
            sbr.getAmountSwapBridgeOutMultihop(BSC_TESTNET, path, amountIn);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.Success));
        uint idx = IBaseBridge(HUB_BRIDGE).getNextInitiateIndex(BSC_TESTNET);
        uint before = user.balance;
        vm.prank(user);
        uint out = sbr.swapBridgeExactInputETH{value: amountIn}(
            ISwapBridgeRouterCL.SwapBridgeExactInputParams({
                path: path,
                amountIn: amountIn,
                amountOutMinimum: quoted,
                bridgeParams: _bp()
            }),
            block.timestamp + 600
        );
        assertEq(out, quoted);
        assertEq(before - user.balance, amountIn);
        assertEq(IBaseBridge(HUB_BRIDGE).getNextInitiateIndex(BSC_TESTNET), idx + 1);
        assertEq(IERC20(WCROSS).balanceOf(address(sbr)), 0);
    }

    function test_fork_exactOutputSingle_refundsUnspent() public onlyFork {
        uint bridgeValue = 50e18;
        (ISwapBridgeRouterCL.QuoteStatus st, uint amountInQuoted,,,) =
            sbr.getAmountSwapBridgeIn(BSC_TESTNET, ONEUSD, USDTX, 200, bridgeValue);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.Success));
        uint maxIn = amountInQuoted * 110 / 100;
        deal(ONEUSD, user, maxIn);
        vm.startPrank(user);
        IERC20(ONEUSD).approve(address(sbr), maxIn);
        uint spent = sbr.swapBridgeExactOutputSingle(
            ISwapBridgeRouterCL.SwapBridgeExactOutputSingleParams({
                tokenIn: ONEUSD,
                tokenOut: USDTX,
                tickSpacing: 200,
                amountOut: bridgeValue,
                amountInMaximum: maxIn,
                sqrtPriceLimitX96: 0,
                bridgeParams: _bp()
            }),
            block.timestamp + 600
        );
        vm.stopPrank();
        assertEq(spent, amountInQuoted, "spent the quoted input");
        assertEq(IERC20(ONEUSD).balanceOf(user), maxIn - spent, "unspent input refunded");
        assertEq(IERC20(ONEUSD).balanceOf(address(sbr)), 0);
    }

    /// @notice The real CL router has no V3 exactInputSingle (uint24 fee) entry point: the old selector reverts
    function test_fork_v3SelectorRejectedByCLRouter() public onlyFork {
        deal(ONEUSD, address(this), 1e18);
        IERC20(ONEUSD).approve(CL_ROUTER, 1e18);
        (bool ok,) = CL_ROUTER.call(
            abi.encodeWithSelector(
                V3_EXACT_INPUT_SINGLE,
                ONEUSD,
                USDTX,
                uint24(200),
                address(this),
                block.timestamp,
                uint(1e18),
                uint(0),
                uint160(0)
            )
        );
        assertFalse(ok, "0x414bf389 must not execute on the CL router");
        // the CL selector with the same values swaps
        uint out = ICLSwapRouter(CL_ROUTER).exactInputSingle(
            ICLSwapRouter.ExactInputSingleParams({
                tokenIn: ONEUSD,
                tokenOut: USDTX,
                tickSpacing: 200,
                recipient: address(this),
                deadline: block.timestamp,
                amountIn: 1e18,
                amountOutMinimum: 0,
                sqrtPriceLimitX96: 0
            })
        );
        assertGt(out, 0);
    }
}
