// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

import {SwapBridgeRouter} from "../src/SwapBridgeRouter.sol";
import {SwapBridgeRouterCL} from "../src/SwapBridgeRouterCL.sol";
import {ICLQuoterV2, ICLSwapRouter} from "../src/interface/ICLSwapRouter.sol";
import {ISwapBridgeRouter} from "../src/interface/ISwapBridgeRouter.sol";
import {ISwapBridgeRouterCL} from "../src/interface/ISwapBridgeRouterCL.sol";

import {BridgeTest} from "./Bridge.t.sol";
import {MockWETH9} from "./SwapBridgeRouter.t.sol";
import {TestToken} from "./token/TestToken.sol";

/**
 * @title MockCLRouter
 * @notice Mock of the Slipstream SwapRouter: single-pool structs carry int24 tickSpacing. It has no V3
 * (uint24 fee) entry points, so a V3-encoded call (0x414bf389) reverts like the real CL router.
 */
contract MockCLRouter is ICLSwapRouter {
    address public immutable WETH9;
    uint public swapRate = 1e18;
    uint public partialConsumeRate = 1e18;
    int24 public lastTickSpacing;
    bytes public lastPath;
    address public reenterTarget;
    bytes public reenterData;

    /// @dev exactInput sends the input to the first pool (as the real router does) rather than parking it on the router
    address public constant POOL = address(uint160(uint(keccak256("MockCLRouter.POOL"))));
    /// @notice Per-token amount a multi-hop exactInput leaves on the router (a partially filled last pool)
    mapping(address => uint) public intermediateLeftover;
    /// @notice Per-token amount of the router's own balance a multi-hop exactInput moves out to the pools
    mapping(address => uint) public intermediateDrain;

    function setIntermediateLeftover(address token, uint amount) external {
        intermediateLeftover[token] = amount;
    }

    function setIntermediateDrain(address token, uint amount) external {
        intermediateDrain[token] = amount;
    }

    constructor(address _weth9) {
        WETH9 = _weth9;
    }

    receive() external payable {}

    function setSwapRate(uint r) external {
        swapRate = r;
    }

    function setPartialConsumeRate(uint r) external {
        partialConsumeRate = r;
    }

    function setReenter(address target, bytes calldata data) external {
        reenterTarget = target;
        reenterData = data;
    }

    function fund(address token, uint amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
    }

    function _maybeReenter() internal {
        if (reenterTarget != address(0)) {
            (bool ok, bytes memory ret) = reenterTarget.call(reenterData);
            if (!ok) {
                assembly {
                    revert(add(ret, 32), mload(ret))
                }
            }
        }
    }

    function exactInputSingle(ExactInputSingleParams calldata p) external payable override returns (uint amountOut) {
        lastTickSpacing = p.tickSpacing;
        _maybeReenter();
        uint consumed = p.sqrtPriceLimitX96 != 0 ? p.amountIn * partialConsumeRate / 1e18 : p.amountIn;
        IERC20(p.tokenIn).transferFrom(msg.sender, address(this), consumed);
        amountOut = consumed * swapRate / 1e18;
        require(amountOut >= p.amountOutMinimum, "Too little received");
        IERC20(p.tokenOut).transfer(p.recipient, amountOut);
    }

    function exactInput(ExactInputParams calldata p) external payable override returns (uint amountOut) {
        lastPath = p.path;
        address tokenIn = address(bytes20(p.path[0:20]));
        address tokenOut = address(bytes20(p.path[p.path.length - 20:]));
        IERC20(tokenIn).transferFrom(msg.sender, POOL, p.amountIn);
        uint hops = (p.path.length - 20) / 23;
        for (uint i = 1; i < hops; ++i) {
            address intermediate = address(bytes20(p.path[i * 23:i * 23 + 20]));
            uint drain = intermediateDrain[intermediate];
            if (drain > 0) IERC20(intermediate).transfer(POOL, drain);
            uint leftover = intermediateLeftover[intermediate];
            if (leftover > 0) TestToken(intermediate).mint(address(this), leftover);
        }
        amountOut = p.amountIn * swapRate / 1e18;
        require(amountOut >= p.amountOutMinimum, "Too little received");
        IERC20(tokenOut).transfer(p.recipient, amountOut);
    }

    function exactOutputSingle(ExactOutputSingleParams calldata p) external payable override returns (uint amountIn) {
        lastTickSpacing = p.tickSpacing;
        amountIn = p.amountOut * 1e18 / swapRate;
        require(amountIn <= p.amountInMaximum, "Too much requested");
        IERC20(p.tokenIn).transferFrom(msg.sender, address(this), amountIn);
        IERC20(p.tokenOut).transfer(p.recipient, p.amountOut);
    }

    function exactOutput(ExactOutputParams calldata p) external payable override returns (uint amountIn) {
        lastPath = p.path;
        address tokenOut = address(bytes20(p.path[0:20]));
        address tokenIn = address(bytes20(p.path[p.path.length - 20:]));
        amountIn = p.amountOut * 1e18 / swapRate;
        require(amountIn <= p.amountInMaximum, "Too much requested");
        IERC20(tokenIn).transferFrom(msg.sender, address(this), amountIn);
        IERC20(tokenOut).transfer(p.recipient, p.amountOut);
    }
}

/// @notice Mock Slipstream QuoterV2 (int24 tickSpacing); reverts for tick spacings flagged as "no pool"
contract MockCLQuoter is ICLQuoterV2 {
    uint public swapRate = 1e18;
    int24 public lastTickSpacing;
    mapping(int24 => bool) public noPool;

    function setNoPool(int24 ts, bool v) external {
        noPool[ts] = v;
    }

    function quoteExactInput(bytes memory, uint amountIn)
        external
        view
        override
        returns (uint, uint160[] memory, uint32[] memory, uint)
    {
        return (amountIn * swapRate / 1e18, new uint160[](0), new uint32[](0), 0);
    }

    function quoteExactInputSingle(QuoteExactInputSingleParams memory p)
        external
        override
        returns (uint, uint160, uint32, uint)
    {
        require(!noPool[p.tickSpacing], "no pool");
        lastTickSpacing = p.tickSpacing;
        return (p.amountIn * swapRate / 1e18, 0, 0, 0);
    }

    function quoteExactOutput(bytes memory, uint amountOut)
        external
        view
        override
        returns (uint, uint160[] memory, uint32[] memory, uint)
    {
        return (amountOut * 1e18 / swapRate, new uint160[](0), new uint32[](0), 0);
    }

    function quoteExactOutputSingle(QuoteExactOutputSingleParams memory p)
        external
        override
        returns (uint, uint160, uint32, uint)
    {
        require(!noPool[p.tickSpacing], "no pool");
        lastTickSpacing = p.tickSpacing;
        return (p.amount * 1e18 / swapRate, 0, 0, 0);
    }
}

/**
 * @title SwapBridgeRouterCLTest
 * @notice Mock-based tests of SwapBridgeRouterCL (Slipstream, int24 tickSpacing) against the real bridge stack of BridgeTest
 */
contract SwapBridgeRouterCLTest is BridgeTest {
    SwapBridgeRouterCL public sbr;
    MockCLRouter public clRouter;
    MockCLQuoter public clQuoter;
    MockWETH9 public mweth;
    IERC20 public outToken;

    bytes4 constant CL_EXACT_INPUT_SINGLE = 0xa026383e;
    bytes4 constant V3_EXACT_INPUT_SINGLE = 0x414bf389;

    function setUp() public override {
        super.setUp();
        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);
        mweth = new MockWETH9();
        clRouter = new MockCLRouter(address(mweth));
        clQuoter = new MockCLQuoter();
        outToken = IERC20(address(new TestToken("Swap Output Token", "SOT", 18)));
        bridgeBSC.registerToken(CROSS_CHAIN_ID, true, address(outToken), address(testTokenCross));
        TestToken(address(outToken)).mint(OWNER, 1_000_000 * 1e18);
        TestToken(address(mweth)).mint(OWNER, 1_000_000 * 1e18);
        outToken.approve(address(clRouter), type(uint).max);
        clRouter.fund(address(outToken), 500_000 * 1e18);
        sbr = new SwapBridgeRouterCL(address(clRouter), address(clQuoter), address(bridgeBSC));

        address[] memory tokens = new address[](1);
        uint[] memory prices = new uint[](1);
        uint[] memory pricesAt = new uint[](1);
        tokens[0] = address(outToken);
        prices[0] = 1;
        priceFeedBSC.updatePrice(tokens, prices, pricesAt);
        cross.transfer(USER, 10_000 * 1e18);
        vm.stopPrank();
        vm.deal(USER, 10_000 ether);
    }

    function _bp() internal view returns (ISwapBridgeRouterCL.BridgeParams memory) {
        return ISwapBridgeRouterCL.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""});
    }

    function _single(int24 ts, uint amountIn, uint minOut, uint160 limit)
        internal
        view
        returns (ISwapBridgeRouterCL.SwapBridgeExactInputSingleParams memory)
    {
        return ISwapBridgeRouterCL.SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(outToken),
            tickSpacing: ts,
            amountIn: amountIn,
            amountOutMinimum: minOut,
            sqrtPriceLimitX96: limit,
            bridgeParams: _bp()
        });
    }

    // ---------------------------------------------------------------- construction

    function test_cl_constructor_readsWETHAndRejectsZero() public {
        vm.selectFork(bscForkID);
        assertEq(address(sbr.WETH9()), address(mweth));
        assertEq(address(sbr.swapRouter()), address(clRouter));
        assertEq(address(sbr.quoter()), address(clQuoter));
        assertEq(address(sbr.bridge()), address(bridgeBSC));
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidAddress.selector);
        new SwapBridgeRouterCL(address(0), address(clQuoter), address(bridgeBSC));
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidAddress.selector);
        new SwapBridgeRouterCL(address(clRouter), address(0), address(bridgeBSC));
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidAddress.selector);
        new SwapBridgeRouterCL(address(clRouter), address(clQuoter), address(0));
    }

    // ---------------------------------------------------------------- selectors / ABI

    function test_cl_selectorsDifferFromV3Wrapper() public pure {
        assertEq(ICLSwapRouter.exactInputSingle.selector, CL_EXACT_INPUT_SINGLE);
        assertEq(ICLQuoterV2.quoteExactInputSingle.selector, bytes4(0x9e7defe6));
        assertEq(ISwapBridgeRouterCL.swapBridgeExactInputSingle.selector, bytes4(0x49d9a7f7));
        assertEq(ISwapBridgeRouterCL.getAmountSwapBridgeOut.selector, bytes4(0xdf7da91c));
        assertTrue(
            ISwapBridgeRouterCL.swapBridgeExactInputSingle.selector
                != ISwapBridgeRouter.swapBridgeExactInputSingle.selector
        );
        assertTrue(
            ISwapBridgeRouterCL.getAmountSwapBridgeIn.selector != ISwapBridgeRouter.getAmountSwapBridgeIn.selector
        );
        // multi-hop entry points keep the bytes path ABI
        assertEq(ISwapBridgeRouterCL.swapBridgeExactInput.selector, ISwapBridgeRouter.swapBridgeExactInput.selector);
    }

    /// @notice A V3-shaped call (uint24 fee struct) sent to the CL wrapper has no matching selector and reverts
    function test_cl_oldV3WrapperSelector_reverts() public {
        vm.selectFork(bscForkID);
        uint amountIn = 100 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory v3 = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(outToken),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });
        vm.prank(USER);
        (bool ok,) = address(sbr).call(
            abi.encodeWithSelector(ISwapBridgeRouter.swapBridgeExactInputSingle.selector, v3, block.timestamp + 1)
        );
        assertFalse(ok, "V3 wrapper selector must not be accepted by the CL wrapper");
        (ok,) = address(sbr).call(
            abi.encodeWithSelector(
                ISwapBridgeRouter.getAmountSwapBridgeOut.selector,
                CROSS_CHAIN_ID,
                address(cross),
                address(outToken),
                uint24(3000),
                amountIn
            )
        );
        assertFalse(ok, "V3 quote selector must not be accepted by the CL wrapper");
    }

    /// @notice The V3 router selector 0x414bf389 does not exist on the CL router (the CL wrapper never emits it)
    function test_cl_router_rejectsV3ExactInputSingle() public {
        vm.selectFork(bscForkID);
        (bool ok,) = address(clRouter).call(
            abi.encodeWithSelector(
                V3_EXACT_INPUT_SINGLE,
                address(cross),
                address(outToken),
                uint24(3000),
                USER,
                block.timestamp,
                uint(1),
                uint(0),
                uint160(0)
            )
        );
        assertFalse(ok);
    }

    // ---------------------------------------------------------------- exact input single

    function test_cl_exactInputSingle_forwardsTickSpacingAndBridges() public {
        vm.selectFork(bscForkID);
        uint amountIn = 1000 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        uint userBefore = cross.balanceOf(USER);
        uint bridgeBefore = outToken.balanceOf(address(bridgeBSC));

        // the router is called with the CL exactInputSingle selector and the exact tickSpacing
        vm.expectCall(
            address(clRouter),
            abi.encodeCall(
                ICLSwapRouter.exactInputSingle,
                (
                    ICLSwapRouter.ExactInputSingleParams({
                        tokenIn: address(cross),
                        tokenOut: address(outToken),
                        tickSpacing: 200,
                        recipient: address(sbr),
                        deadline: block.timestamp + 1 hours,
                        amountIn: amountIn,
                        amountOutMinimum: 0,
                        sqrtPriceLimitX96: 0
                    })
                )
            )
        );
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInputSingle(_single(200, amountIn, 0, 0), block.timestamp + 1 hours);

        assertEq(out, amountIn);
        assertEq(clRouter.lastTickSpacing(), int24(200));
        assertEq(cross.balanceOf(USER), userBefore - amountIn);
        assertGt(outToken.balanceOf(address(bridgeBSC)), bridgeBefore);
        assertEq(cross.balanceOf(address(sbr)), 0);
        assertEq(outToken.balanceOf(address(sbr)), 0);
        assertEq(cross.allowance(address(sbr), address(clRouter)), 0, "router allowance cleared");
    }

    /// @notice Negative tick spacings are passed through unchanged (int24 two's complement), never re-derived
    function test_cl_negativeTickSpacing_passedThrough() public {
        vm.selectFork(bscForkID);
        int24[3] memory values = [int24(-200), type(int24).min, type(int24).max];
        for (uint i = 0; i < values.length; ++i) {
            vm.prank(USER);
            cross.approve(address(sbr), 10 * 1e18);
            vm.prank(USER);
            sbr.swapBridgeExactInputSingle(_single(values[i], 10 * 1e18, 0, 0), block.timestamp + 1);
            assertEq(clRouter.lastTickSpacing(), values[i]);
        }
        // the int24 word is sign-extended in calldata
        bytes memory data = abi.encodeCall(sbr.swapBridgeExactInputSingle, (_single(-1, 1, 0, 0), 1));
        bytes32 word;
        assembly {
            word := mload(add(data, add(32, add(4, 0x80))))
        }
        assertEq(word, bytes32(type(uint).max), "tickSpacing -1 is encoded as 0xff..ff");
    }

    function test_cl_exactInputSingle_slippage_reverts() public {
        vm.selectFork(bscForkID);
        uint amountIn = 100 * 1e18;
        clRouter.setSwapRate(0.9e18);
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        vm.prank(USER);
        vm.expectRevert("Too little received");
        sbr.swapBridgeExactInputSingle(_single(200, amountIn, 90 * 1e18 + 1, 0), block.timestamp + 1);
        // at the boundary it passes
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInputSingle(_single(200, amountIn, 90 * 1e18, 0), block.timestamp + 1);
        assertEq(out, 90 * 1e18);
    }

    function test_cl_deadlineExpired_reverts() public {
        vm.selectFork(bscForkID);
        vm.prank(USER);
        cross.approve(address(sbr), 1e18);
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouterCL.SBRDeadlineExpired.selector);
        sbr.swapBridgeExactInputSingle(_single(200, 1e18, 0, 0), block.timestamp - 1);
    }

    function test_cl_partialFill_refundsUnspent() public {
        vm.selectFork(bscForkID);
        uint amountIn = 100 * 1e18;
        clRouter.setPartialConsumeRate(0.5e18);
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        uint before = cross.balanceOf(USER);
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInputSingle(_single(200, amountIn, 0, 1), block.timestamp + 1);
        assertEq(out, 50 * 1e18);
        assertEq(before - cross.balanceOf(USER), 50 * 1e18, "unspent half refunded");
        assertEq(cross.balanceOf(address(sbr)), 0);
    }

    function test_cl_sameToken_reverts() public {
        vm.selectFork(bscForkID);
        ISwapBridgeRouterCL.SwapBridgeExactInputSingleParams memory p = _single(200, 1e18, 0, 0);
        p.tokenOut = address(cross);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRCircularPath.selector, address(cross)));
        sbr.swapBridgeExactInputSingle(p, block.timestamp + 1);
    }

    function test_cl_reentrancy_reverts() public {
        vm.selectFork(bscForkID);
        clRouter.setReenter(
            address(sbr), abi.encodeCall(sbr.swapBridgeExactInputSingle, (_single(200, 1, 0, 0), block.timestamp + 1))
        );
        vm.prank(USER);
        cross.approve(address(sbr), 1e18);
        vm.prank(USER);
        vm.expectRevert(ReentrancyGuardTransient.ReentrancyGuardReentrantCall.selector);
        sbr.swapBridgeExactInputSingle(_single(200, 1e18, 0, 0), block.timestamp + 1);
    }

    // ---------------------------------------------------------------- ETH / exact output / multi-hop

    function test_cl_exactInputSingleETH() public {
        vm.selectFork(bscForkID);
        uint amountIn = 10 ether;
        ISwapBridgeRouterCL.SwapBridgeExactInputSingleParams memory p = _single(100, amountIn, 0, 0);
        p.tokenIn = address(mweth);
        uint before = USER.balance;
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInputSingleETH{value: amountIn}(p, block.timestamp + 1);
        assertEq(out, amountIn);
        assertEq(USER.balance, before - amountIn);
        assertEq(clRouter.lastTickSpacing(), int24(100));
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidValue.selector);
        sbr.swapBridgeExactInputSingleETH{value: amountIn - 1}(p, block.timestamp + 1);
    }

    function test_cl_exactOutputSingleETH_refundsExcessAndUnspent() public {
        vm.selectFork(bscForkID);
        // fund the router with WETH so exact-output can pay out tokenOut = outToken
        ISwapBridgeRouterCL.SwapBridgeExactOutputSingleParams memory p = ISwapBridgeRouterCL
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(mweth),
            tokenOut: address(outToken),
            tickSpacing: -50,
            amountOut: 100 * 1e18,
            amountInMaximum: 200 ether,
            sqrtPriceLimitX96: 0,
            bridgeParams: _bp()
        });
        (, uint nf, uint ef) = sbr.calculateBridgeFees(CROSS_CHAIN_ID, outToken, 100 * 1e18);
        uint before = USER.balance;
        vm.prank(USER);
        uint amountIn = sbr.swapBridgeExactOutputSingleETH{value: 250 ether}(p, block.timestamp + 1);
        assertEq(amountIn, 100 * 1e18 + nf + ef, "pays only the swap input at 1:1");
        assertEq(before - USER.balance, amountIn, "excess msg.value and unspent WETH refunded");
        assertEq(clRouter.lastTickSpacing(), int24(-50));
        assertEq(mweth.balanceOf(address(sbr)), 0);
        assertEq(address(sbr).balance, 0);
    }

    function test_cl_exactOutputSingle_slippage_reverts() public {
        vm.selectFork(bscForkID);
        ISwapBridgeRouterCL.SwapBridgeExactOutputSingleParams memory p = ISwapBridgeRouterCL
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(outToken),
            tickSpacing: 200,
            amountOut: 100 * 1e18,
            amountInMaximum: 199 * 1e18, // at a 1:0.5 rate the swap needs >= 200 input
            sqrtPriceLimitX96: 0,
            bridgeParams: _bp()
        });
        clRouter.setSwapRate(0.5e18);
        vm.prank(USER);
        cross.approve(address(sbr), type(uint).max);
        vm.prank(USER);
        vm.expectRevert("Too much requested");
        sbr.swapBridgeExactOutputSingle(p, block.timestamp + 1);
    }

    function test_cl_multihop_int24Path() public {
        vm.selectFork(bscForkID);
        uint amountIn = 100 * 1e18;
        // path fields are int24 tick spacings; a negative one is passed through as raw bytes
        bytes memory path = abi.encodePacked(address(cross), int24(-100), address(mweth), int24(200), address(outToken));
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInput(
            ISwapBridgeRouterCL.SwapBridgeExactInputParams({
                path: path,
                amountIn: amountIn,
                amountOutMinimum: amountIn,
                bridgeParams: _bp()
            }),
            block.timestamp + 1
        );
        assertEq(out, amountIn);
        assertEq(clRouter.lastPath(), path);
        // malformed path length
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidPath.selector);
        sbr.swapBridgeExactInput(
            ISwapBridgeRouterCL.SwapBridgeExactInputParams({
                path: abi.encodePacked(path, uint8(1)),
                amountIn: amountIn,
                amountOutMinimum: 0,
                bridgeParams: _bp()
            }),
            block.timestamp + 1
        );
    }

    // ---------------------------------------------------------------- quotes

    function test_cl_getAmountSwapBridgeOut_usesTickSpacing() public {
        vm.selectFork(bscForkID);
        (ISwapBridgeRouterCL.QuoteStatus st, uint swapOut, uint bv, uint nf, uint ef) =
            sbr.getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(outToken), -200, 1000 * 1e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.Success));
        assertEq(swapOut, 1000 * 1e18);
        assertEq(bv + nf + ef <= swapOut, true);
        assertEq(clQuoter.lastTickSpacing(), int24(-200));

        clQuoter.setNoPool(2000, true);
        (st,,,,) = sbr.getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(outToken), 2000, 1e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.InvalidSwap));
    }

    function test_cl_getAmountSwapBridgeIn_usesTickSpacing() public {
        vm.selectFork(bscForkID);
        (ISwapBridgeRouterCL.QuoteStatus st, uint amountIn, uint swapOut,,) =
            sbr.getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(outToken), 50, 100 * 1e18);
        assertEq(uint(st), uint(ISwapBridgeRouterCL.QuoteStatus.Success));
        assertEq(amountIn, swapOut);
        assertEq(clQuoter.lastTickSpacing(), int24(50));
    }

    /// @notice The V3 wrapper is a separate, unchanged contract: it keeps its uint24 fee API
    function test_cl_v3WrapperUnchanged() public pure {
        assertEq(SwapBridgeRouter.swapBridgeExactInputSingle.selector, bytes4(0x80413d2f));
        assertEq(SwapBridgeRouter.getAmountSwapBridgeOut.selector, bytes4(0x86f64142));
    }

    // ---------------------------------------------------------------- multi-hop partial fill (CBU-08 port) and multi-hop variants

    function _mh(bytes memory path, uint amountIn)
        internal
        view
        returns (ISwapBridgeRouterCL.SwapBridgeExactInputParams memory)
    {
        return ISwapBridgeRouterCL.SwapBridgeExactInputParams({
            path: path,
            amountIn: amountIn,
            amountOutMinimum: 0,
            bridgeParams: _bp()
        });
    }

    function _mhOut(bytes memory path, uint amountOut, uint maxIn)
        internal
        view
        returns (ISwapBridgeRouterCL.SwapBridgeExactOutputParams memory)
    {
        return ISwapBridgeRouterCL.SwapBridgeExactOutputParams({
            path: path,
            amountOut: amountOut,
            amountInMaximum: maxIn,
            bridgeParams: _bp()
        });
    }

    /// A partially filled last pool leaves the intermediate on the CL router and reverts
    function test_cl_multihop_partialFill_reverts() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        uint amountIn = 1000 * 1e18;
        uint leftover = 7 * 1e18;
        clRouter.setIntermediateLeftover(address(mid), leftover);
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        bytes memory path = abi.encodePacked(address(cross), int24(-200), address(mid), int24(200), address(outToken));
        uint userBefore = cross.balanceOf(USER);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRPartialFill.selector, address(mid), leftover));
        sbr.swapBridgeExactInput(_mh(path, amountIn), block.timestamp + 1 hours);
        assertEq(cross.balanceOf(USER), userBefore, "caller keeps the original tokens");
        assertEq(cross.balanceOf(address(sbr)), 0);
        assertEq(mid.balanceOf(address(clRouter)), 0, "revert unwinds the leftover");
    }

    function test_cl_multihopETH_partialFill_reverts() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        uint leftover = 3 * 1e17;
        clRouter.setIntermediateLeftover(address(mid), leftover);
        bytes memory path = abi.encodePacked(address(mweth), int24(2000), address(mid), int24(-1), address(outToken));
        uint userBefore = USER.balance;
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRPartialFill.selector, address(mid), leftover));
        sbr.swapBridgeExactInputETH{value: 1 ether}(_mh(path, 1 ether), block.timestamp + 1 hours);
        assertEq(USER.balance, userBefore, "caller keeps the ETH");
        assertEq(address(sbr).balance, 0);
    }

    /// A balance already on the CL router (no new leftover) is not a partial fill
    function test_cl_multihop_preExistingIntermediateBalance_succeeds() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        mid.mint(address(clRouter), 500 * 1e18);
        uint amountIn = 1000 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        bytes memory path = abi.encodePacked(address(cross), int24(200), address(mid), int24(200), address(outToken));
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInput(_mh(path, amountIn), block.timestamp + 1 hours);
        assertEq(out, amountIn);
        assertEq(mid.balanceOf(address(clRouter)), 500 * 1e18, "pre-existing intermediate balance untouched");
    }

    /// An intermediate balance that decreased during the swap is not a partial fill
    function test_cl_multihop_decreasedIntermediateBalance_succeeds() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        mid.mint(address(clRouter), 500 * 1e18);
        clRouter.setIntermediateDrain(address(mid), 200 * 1e18);
        uint amountIn = 1000 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        bytes memory path = abi.encodePacked(address(cross), int24(200), address(mid), int24(200), address(outToken));
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInput(_mh(path, amountIn), block.timestamp + 1 hours);
        assertEq(out, amountIn);
        assertEq(mid.balanceOf(address(clRouter)), 300 * 1e18);
        assertEq(cross.balanceOf(address(sbr)), 0);
    }

    /// Three hops: a drained intermediate does not mask a leftover on another; each leftover names its token
    function test_cl_threeHop_leftover_reverts() public {
        vm.selectFork(bscForkID);
        TestToken mid1 = new TestToken("Mid1", "M1", 18);
        TestToken mid2 = new TestToken("Mid2", "M2", 18);
        mid1.mint(address(clRouter), 500 * 1e18);
        clRouter.setIntermediateDrain(address(mid1), 200 * 1e18);
        clRouter.setIntermediateLeftover(address(mid2), 9 * 1e18);
        uint amountIn = 1000 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), amountIn);
        bytes memory path = abi.encodePacked(
            address(cross), int24(50), address(mid1), int24(-100), address(mid2), int24(2000), address(outToken)
        );
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRPartialFill.selector, address(mid2), 9 * 1e18));
        sbr.swapBridgeExactInput(_mh(path, amountIn), block.timestamp + 1 hours);

        clRouter.setIntermediateDrain(address(mid1), 0);
        clRouter.setIntermediateLeftover(address(mid2), 0);
        clRouter.setIntermediateLeftover(address(mid1), 4 * 1e18);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRPartialFill.selector, address(mid1), 4 * 1e18));
        sbr.swapBridgeExactInput(_mh(path, amountIn), block.timestamp + 1 hours);
    }

    /// Native-token multi-hop exact input with a negative int24 segment
    function test_cl_multihopETH_exactInput_negativeSegment() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        bytes memory path = abi.encodePacked(address(mweth), int24(-200), address(mid), int24(100), address(outToken));
        uint before = USER.balance;
        vm.prank(USER);
        uint out = sbr.swapBridgeExactInputETH{value: 2 ether}(_mh(path, 2 ether), block.timestamp + 1 hours);
        assertEq(out, 2 ether);
        assertEq(before - USER.balance, 2 ether);
        assertEq(clRouter.lastPath(), path, "path (incl. 0xffff38 = -200) forwarded unchanged");
        assertEq(mweth.balanceOf(address(sbr)), 0);
        assertEq(address(sbr).balance, 0);
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidValue.selector);
        sbr.swapBridgeExactInputETH{value: 1 ether}(_mh(path, 2 ether), block.timestamp + 1 hours);
    }

    /// ERC20 multi-hop exact output (reversed path) with a negative segment: pays the bridge fees, refunds unspent input
    function test_cl_multihop_exactOutput_negativeSegment() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        uint bridgeValue = 100 * 1e18;
        (, uint nf, uint ef) = sbr.calculateBridgeFees(CROSS_CHAIN_ID, outToken, bridgeValue);
        // exactOutput path: tokenOut -> ... -> tokenIn
        bytes memory path = abi.encodePacked(address(outToken), int24(-50), address(mid), int24(200), address(cross));
        uint maxIn = 300 * 1e18;
        vm.prank(USER);
        cross.approve(address(sbr), maxIn);
        uint before = cross.balanceOf(USER);
        uint idx = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        vm.prank(USER);
        uint spent = sbr.swapBridgeExactOutput(_mhOut(path, bridgeValue, maxIn), block.timestamp + 1 hours);
        assertEq(spent, bridgeValue + nf + ef, "1:1 mock rate");
        assertEq(before - cross.balanceOf(USER), spent, "unspent input refunded");
        assertEq(clRouter.lastPath(), path);
        assertEq(bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID), idx + 1);
        assertEq(cross.balanceOf(address(sbr)), 0);
        assertEq(outToken.balanceOf(address(sbr)), 0);
        // slippage: amountInMaximum below the required input
        vm.prank(USER);
        cross.approve(address(sbr), type(uint).max);
        vm.prank(USER);
        vm.expectRevert("Too much requested");
        sbr.swapBridgeExactOutput(_mhOut(path, bridgeValue, bridgeValue + nf + ef - 1), block.timestamp + 1 hours);
        // circular path is refused
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(SwapBridgeRouterCL.SBRCircularPath.selector, address(cross)));
        sbr.swapBridgeExactOutput(
            _mhOut(abi.encodePacked(address(cross), int24(-50), address(mid), int24(200), address(cross)), 1e18, maxIn),
            block.timestamp + 1 hours
        );
    }

    /// Native-token multi-hop exact output: excess msg.value and unspent WETH are refunded as ETH
    function test_cl_multihopETH_exactOutput_refunds() public {
        vm.selectFork(bscForkID);
        TestToken mid = new TestToken("Mid", "MID", 18);
        uint bridgeValue = 100 * 1e18;
        (, uint nf, uint ef) = sbr.calculateBridgeFees(CROSS_CHAIN_ID, outToken, bridgeValue);
        bytes memory path = abi.encodePacked(address(outToken), int24(200), address(mid), int24(-2000), address(mweth));
        uint before = USER.balance;
        vm.prank(USER);
        uint spent = sbr.swapBridgeExactOutputETH{value: 250 ether}(
            _mhOut(path, bridgeValue, 200 ether), block.timestamp + 1 hours
        );
        assertEq(spent, bridgeValue + nf + ef);
        assertEq(before - USER.balance, spent, "excess value and unspent WETH refunded");
        assertEq(clRouter.lastPath(), path);
        assertEq(mweth.balanceOf(address(sbr)), 0);
        assertEq(address(sbr).balance, 0);
        // path must end in WETH9
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouterCL.SBRInvalidAddress.selector);
        sbr.swapBridgeExactOutputETH{value: 200 ether}(
            _mhOut(
                abi.encodePacked(address(outToken), int24(200), address(mid), int24(-2000), address(cross)),
                bridgeValue,
                200 ether
            ),
            block.timestamp + 1 hours
        );
    }
}
