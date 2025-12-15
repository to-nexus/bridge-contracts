// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {SwapBridgeRouter} from "../src/SwapBridgeRouter.sol";

import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {IQuoterV2, ISwapBridgeRouter, ISwapRouter, IWETH9} from "../src/interface/ISwapBridgeRouter.sol";
import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {TestToken} from "./token/TestToken.sol";

/**
 * @title MockWETH9
 * @notice Mock WETH9 contract for testing
 */
contract MockWETH9 is TestToken {
    constructor() TestToken("Wrapped Ether", "WETH", 18) {}

    function deposit() external payable {
        _mint(msg.sender, msg.value);
    }

    function withdraw(uint amount) external {
        _burn(msg.sender, amount);
        (bool success,) = msg.sender.call{value: amount}("");
        require(success, "WETH: withdraw failed");
    }

    receive() external payable {
        _mint(msg.sender, msg.value);
    }
}

/**
 * @title MockUniswapV3Router
 * @notice Mock Uniswap V3 Router for testing swap functionality
 * @dev Implements ISwapRouter and provides WETH9() via IPeripheryImmutableState
 */
contract MockUniswapV3Router is ISwapRouter {
    address public immutable WETH9;
    uint public swapRate = 1e18; // 1:1 default rate

    constructor(address _weth9) {
        WETH9 = _weth9;
    }

    function setSwapRate(uint _rate) external {
        swapRate = _rate;
    }

    // IUniswapV3SwapCallback implementation (required by ISwapRouter)
    function uniswapV3SwapCallback(int, int, bytes calldata) external pure override {
        // Mock implementation - do nothing
    }

    function exactInputSingle(ExactInputSingleParams calldata params)
        external
        payable
        override
        returns (uint amountOut)
    {
        // Transfer tokenIn from sender
        IERC20(params.tokenIn).transferFrom(msg.sender, address(this), params.amountIn);

        // Calculate output amount
        amountOut = params.amountIn * swapRate / 1e18;
        require(amountOut >= params.amountOutMinimum, "Too little received");

        // Transfer tokenOut to recipient
        IERC20(params.tokenOut).transfer(params.recipient, amountOut);
    }

    function exactInput(ExactInputParams calldata params) external payable override returns (uint amountOut) {
        // Extract tokenIn and tokenOut from path
        address tokenIn = address(bytes20(params.path[0:20]));
        address tokenOut = address(bytes20(params.path[params.path.length - 20:]));

        // Transfer tokenIn from sender
        IERC20(tokenIn).transferFrom(msg.sender, address(this), params.amountIn);

        // Calculate output amount
        amountOut = params.amountIn * swapRate / 1e18;
        require(amountOut >= params.amountOutMinimum, "Too little received");

        // Transfer tokenOut to recipient
        IERC20(tokenOut).transfer(params.recipient, amountOut);
    }

    function exactOutputSingle(ExactOutputSingleParams calldata params)
        external
        payable
        override
        returns (uint amountIn)
    {
        // Calculate required input amount
        amountIn = params.amountOut * 1e18 / swapRate;
        require(amountIn <= params.amountInMaximum, "Too much requested");

        // Transfer tokenIn from sender
        IERC20(params.tokenIn).transferFrom(msg.sender, address(this), amountIn);

        // Transfer tokenOut to recipient
        IERC20(params.tokenOut).transfer(params.recipient, params.amountOut);
    }

    function exactOutput(ExactOutputParams calldata params) external payable override returns (uint amountIn) {
        // For exactOutput, path is reversed: tokenOut is first, tokenIn is last
        address tokenOut = address(bytes20(params.path[0:20]));
        address tokenIn = address(bytes20(params.path[params.path.length - 20:]));

        // Calculate required input amount
        amountIn = params.amountOut * 1e18 / swapRate;
        require(amountIn <= params.amountInMaximum, "Too much requested");

        // Transfer tokenIn from sender
        IERC20(tokenIn).transferFrom(msg.sender, address(this), amountIn);

        // Transfer tokenOut to recipient
        IERC20(tokenOut).transfer(params.recipient, params.amountOut);
    }

    /// @notice Fund the router with tokens for testing
    function fund(address token, uint amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
    }
}

/**
 * @title MockUniswapV3Quoter
 * @notice Mock Uniswap V3 QuoterV2 for testing quote functionality
 */
contract MockUniswapV3Quoter is IQuoterV2 {
    uint public swapRate = 1e18; // 1:1 default rate

    function setSwapRate(uint _rate) external {
        swapRate = _rate;
    }

    function quoteExactInputSingle(QuoteExactInputSingleParams memory params)
        external
        view
        override
        returns (uint amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint gasEstimate)
    {
        amountOut = params.amountIn * swapRate / 1e18;
        sqrtPriceX96After = 0;
        initializedTicksCrossed = 0;
        gasEstimate = 100000;
    }

    function quoteExactInput(bytes memory, /* path */ uint amountIn)
        external
        view
        override
        returns (
            uint amountOut,
            uint160[] memory sqrtPriceX96AfterList,
            uint32[] memory initializedTicksCrossedList,
            uint gasEstimate
        )
    {
        amountOut = amountIn * swapRate / 1e18;
        sqrtPriceX96AfterList = new uint160[](1);
        initializedTicksCrossedList = new uint32[](1);
        gasEstimate = 100000;
    }

    function quoteExactOutputSingle(QuoteExactOutputSingleParams memory params)
        external
        view
        override
        returns (uint amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint gasEstimate)
    {
        amountIn = params.amount * 1e18 / swapRate;
        sqrtPriceX96After = 0;
        initializedTicksCrossed = 0;
        gasEstimate = 100000;
    }

    function quoteExactOutput(bytes memory, /* path */ uint amountOut)
        external
        view
        override
        returns (
            uint amountIn,
            uint160[] memory sqrtPriceX96AfterList,
            uint32[] memory initializedTicksCrossedList,
            uint gasEstimate
        )
    {
        amountIn = amountOut * 1e18 / swapRate;
        sqrtPriceX96AfterList = new uint160[](1);
        initializedTicksCrossedList = new uint32[](1);
        gasEstimate = 100000;
    }
}

contract SwapBridgeRouterTest is BridgeTest {
    SwapBridgeRouter public swapBridgeRouterBSC;
    MockUniswapV3Router public mockSwapRouterBSC;
    MockUniswapV3Quoter public mockQuoterBSC;
    MockWETH9 public mockWETHBSC;

    // Test token for swap output
    IERC20 public swapOutputTokenBSC;

    function setUp() public override {
        super.setUp();

        // Deploy mock contracts on BSC
        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);

        // Deploy Mock WETH
        mockWETHBSC = new MockWETH9();

        // Deploy Mock Uniswap Router
        mockSwapRouterBSC = new MockUniswapV3Router(address(mockWETHBSC));

        // Deploy Mock Uniswap Quoter
        mockQuoterBSC = new MockUniswapV3Quoter();

        // Create an additional test token for swap output (different from cross)
        swapOutputTokenBSC = IERC20(address(new TestToken("Swap Output Token", "SOT", 18)));

        // Register swap output token with bridge for CROSS chain
        bridgeBSC.registerToken(CROSS_CHAIN_ID, true, address(swapOutputTokenBSC), address(testTokenCross));

        // Mint tokens for testing
        TestToken(address(swapOutputTokenBSC)).mint(OWNER, 1_000_000 * 1e18);
        TestToken(address(mockWETHBSC)).mint(OWNER, 1_000_000 * 1e18);

        // Fund mock swap router with output tokens
        swapOutputTokenBSC.approve(address(mockSwapRouterBSC), 1_000_000 * 1e18);
        mockSwapRouterBSC.fund(address(swapOutputTokenBSC), 500_000 * 1e18);

        // Fund mock swap router with WETH for exact output tests
        mockWETHBSC.approve(address(mockSwapRouterBSC), 1_000_000 * 1e18);
        mockSwapRouterBSC.fund(address(mockWETHBSC), 500_000 * 1e18);

        // Deploy SwapBridgeRouter
        swapBridgeRouterBSC =
            new SwapBridgeRouter(address(mockSwapRouterBSC), address(mockQuoterBSC), address(bridgeBSC));

        // Set token price for swapOutputTokenBSC
        address[] memory tokens = new address[](1);
        uint[] memory prices = new uint[](1);
        uint[] memory pricesAt = new uint[](1);
        tokens[0] = address(swapOutputTokenBSC);
        prices[0] = 1; // $1
        pricesAt[0] = 0;
        priceFeedBSC.updatePrice(tokens, prices, pricesAt);

        vm.stopPrank();

        // Give USER some tokens
        vm.prank(OWNER);
        cross.transfer(USER, 10_000 * 1e18);
        vm.prank(OWNER);
        TestToken(address(mockWETHBSC)).transfer(USER, 10_000 * 1e18);
    }

    /**
     * @notice Test swapExactInputSingleAndBridge
     */
    function test_swapExactInputSingleAndBridge() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter to spend CROSS tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeUserCrossBalance = cross.balanceOf(USER);
        uint beforeBridgeOutputBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        // Build params
        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap and bridge
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User CROSS balance should decrease");
        assertGt(
            swapOutputTokenBSC.balanceOf(address(bridgeBSC)),
            beforeBridgeOutputBalance,
            "Bridge should receive output tokens"
        );
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapExactInputAndBridge (multi-hop)
     */
    function test_swapExactInputAndBridge() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        // Build path: CROSS -> (fee) -> SwapOutputToken
        bytes memory path = abi.encodePacked(address(cross), uint24(3000), address(swapOutputTokenBSC));

        ISwapBridgeRouter.ExactInputAndBridgeParams memory params = ISwapBridgeRouter.ExactInputAndBridgeParams({
            path: path,
            amountIn: amountIn,
            amountOutMinimum: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        // Execute
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapExactInputAndBridge(params, block.timestamp + 1 hours);

        // Verify
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User CROSS balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapExactOutputSingleAndBridge
     */
    function test_swapExactOutputSingleAndBridge() public {
        uint amountOut = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.ExactOutputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactOutputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: amountOut,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapExactOutputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify: with 1:1 rate, amountIn should equal amountOut
        assertEq(amountIn, amountOut, "Amount in should equal amount out with 1:1 rate");
        // User should have been refunded excess tokens
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn");
    }

    /**
     * @notice Test swapExactOutputAndBridge (multi-hop)
     */
    function test_swapExactOutputAndBridge() public {
        uint amountOut = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        // Build path (reversed for exactOutput): SwapOutputToken -> (fee) -> CROSS
        bytes memory path = abi.encodePacked(address(swapOutputTokenBSC), uint24(3000), address(cross));

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.ExactOutputAndBridgeParams memory params = ISwapBridgeRouter.ExactOutputAndBridgeParams({
            path: path,
            amountOut: amountOut,
            amountInMaximum: amountInMaximum,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapExactOutputAndBridge(params, block.timestamp + 1 hours);

        // Verify
        assertEq(amountIn, amountOut, "Amount in should equal amount out with 1:1 rate");
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn");
    }

    /**
     * @notice Test swapExactInputSingleETHAndBridge
     */
    function test_swapExactInputSingleETHAndBridge() public {
        uint amountIn = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute with ETH
        vm.prank(USER);
        uint amountOut =
            swapBridgeRouterBSC.swapExactInputSingleETHAndBridge{value: amountIn}(params, block.timestamp + 1 hours);

        // Verify
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User ETH balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapExactInputETHAndBridge (multi-hop with ETH)
     */
    function test_swapExactInputETHAndBridge() public {
        uint amountIn = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        // Build path: WETH -> (fee) -> SwapOutputToken
        bytes memory path = abi.encodePacked(address(mockWETHBSC), uint24(3000), address(swapOutputTokenBSC));

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.ExactInputAndBridgeParams memory params = ISwapBridgeRouter.ExactInputAndBridgeParams({
            path: path,
            amountIn: amountIn,
            amountOutMinimum: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountOut =
            swapBridgeRouterBSC.swapExactInputETHAndBridge{value: amountIn}(params, block.timestamp + 1 hours);

        // Verify
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User ETH balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapExactOutputSingleETHAndBridge
     */
    function test_swapExactOutputSingleETHAndBridge() public {
        uint amountOut = 0.5 ether;
        uint amountInMaximum = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.ExactOutputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactOutputSingleAndBridgeParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: amountOut,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapExactOutputSingleETHAndBridge{value: amountInMaximum}(
            params, block.timestamp + 1 hours
        );

        // Verify: with 1:1 rate, amountIn should equal amountOut, and excess should be refunded
        assertEq(amountIn, amountOut, "Amount in should equal amount out with 1:1 rate");
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User should be refunded excess ETH");
    }

    /**
     * @notice Test swapExactOutputETHAndBridge (multi-hop)
     */
    function test_swapExactOutputETHAndBridge() public {
        uint amountOut = 0.5 ether;
        uint amountInMaximum = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        // Build path (reversed for exactOutput): SwapOutputToken -> (fee) -> WETH
        bytes memory path = abi.encodePacked(address(swapOutputTokenBSC), uint24(3000), address(mockWETHBSC));

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.ExactOutputAndBridgeParams memory params = ISwapBridgeRouter.ExactOutputAndBridgeParams({
            path: path,
            amountOut: amountOut,
            amountInMaximum: amountInMaximum,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn =
            swapBridgeRouterBSC.swapExactOutputETHAndBridge{value: amountInMaximum}(params, block.timestamp + 1 hours);

        // Verify
        assertEq(amountIn, amountOut, "Amount in should equal amount out with 1:1 rate");
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User should be refunded excess ETH");
    }

    /**
     * @notice Test getExpectedBridgeAmount view function
     */
    function test_getExpectedBridgeAmount() public {
        uint totalAmount = 1000 * 1e18;

        vm.selectFork(bscForkID);

        (bool ok, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, totalAmount);

        // Verify ok status
        assertTrue(ok, "Should return ok");

        // Verify fee breakdown sums to total
        assertEq(bridgeValue + networkFee + exFee, totalAmount, "Fees should sum to total");

        // In test setup, networkFee and exFeeRate are 0, so bridgeValue should equal totalAmount
        assertEq(bridgeValue, totalAmount, "Bridge value should equal total (no fees configured)");
        assertEq(networkFee, 0, "Network fee should be 0 (not configured)");
        assertEq(exFee, 0, "Exchange fee should be 0 (not configured)");
    }

    /**
     * @notice Test calculateBridgeFees view function
     */
    function test_calculateBridgeFees() public {
        uint value = 1000 * 1e18;

        vm.selectFork(bscForkID);

        (uint minimumValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.calculateBridgeFees(CROSS_CHAIN_ID, swapOutputTokenBSC, value);

        // In test setup, fees are not configured, so all should be 0
        assertEq(minimumValue, 0, "Minimum value should be 0 (not configured)");
        assertEq(networkFee, 0, "Network fee should be 0 (not configured)");
        assertEq(exFee, 0, "Exchange fee should be 0 (not configured)");
    }

    /**
     * @notice Test deadline check
     */
    function test_revert_deadlineExpired() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Set deadline to past
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouter.SBRDeadlineExpired.selector);
        swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp - 1);
    }

    /**
     * @notice Test with different swap rate
     */
    function test_swapWithDifferentRate() public {
        uint amountIn = 1000 * 1e18;
        uint swapRate = 2 * 1e18; // 1 input = 2 output

        vm.selectFork(bscForkID);

        // Set swap rate
        mockSwapRouterBSC.setSwapRate(swapRate);

        // Approve
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: amountIn * 2, // Expect 2x output
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify: should get 2x output
        assertEq(amountOut, amountIn * 2, "Should get 2x output with 2:1 rate");
    }

    /**
     * @notice Test SwapAndBridge event includes correct initiateIndex
     */
    function test_swapAndBridgeEvent_emitsInitiateIndex() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Get the expected initiateIndex before the swap
        uint expectedInitiateIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);

        // Approve SwapBridgeRouter to spend CROSS tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Expect the SwapAndBridge event with correct initiateIndex
        vm.expectEmit(true, true, true, true);
        emit ISwapBridgeRouter.SwapAndBridge(
            USER,
            address(cross),
            address(swapOutputTokenBSC),
            amountIn,
            amountIn, // 1:1 rate
            CROSS_CHAIN_ID,
            expectedInitiateIndex,
            amountIn, // bridgeValue (no fees)
            0, // networkFee
            0 // exFee
        );

        // Execute swap and bridge
        vm.prank(USER);
        swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify the initiateIndex was incremented
        assertEq(
            bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID),
            expectedInitiateIndex + 1,
            "Initiate index should be incremented"
        );
    }

    // ============ Quote Function Tests ============

    /**
     * @notice Test getAmountSwapBridgeOut - single pool quote
     */
    function test_getAmountSwapBridgeOut() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);

        // Verify all return values
        assertTrue(ok, "Quote should succeed");

        // With 1:1 swap rate: swapAmountOut = amountIn
        assertEq(swapAmountOut, amountIn, "Swap amount out should equal input with 1:1 rate");

        // With no fees configured: bridgeValue = swapAmountOut, networkFee = 0, exFee = 0
        assertEq(bridgeValue, swapAmountOut, "Bridge value should equal swap output (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(bridgeValue + networkFee + exFee, swapAmountOut, "Fees should sum to swap output");

        // Verify bridge value calculation matches getExpectedBridgeAmount
        (bool expectedOk, uint expectedBridgeValue, uint expectedNetworkFee, uint expectedExFee) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, swapAmountOut);
        assertTrue(expectedOk, "Expected calculation should succeed");
        assertEq(bridgeValue, expectedBridgeValue, "Bridge value should match expected");
        assertEq(networkFee, expectedNetworkFee, "Network fee should match expected");
        assertEq(exFee, expectedExFee, "Exchange fee should match expected");
    }

    /**
     * @notice Test getAmountSwapBridgeOut with different swap rate
     */
    function test_getAmountSwapBridgeOut_differentRate() public {
        uint amountIn = 1000 * 1e18;
        uint swapRate = 2 * 1e18; // 1:2 rate
        uint expectedSwapOut = amountIn * 2; // 2000 * 1e18

        vm.selectFork(bscForkID);

        // Set swap rate on quoter
        mockQuoterBSC.setSwapRate(swapRate);

        (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);

        // Verify
        assertTrue(ok, "Quote should succeed");
        assertEq(swapAmountOut, expectedSwapOut, "Swap output should be 2x with 1:2 rate");

        // With no fees configured: bridgeValue = swapAmountOut
        assertEq(bridgeValue, expectedSwapOut, "Bridge value should equal swap output (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(bridgeValue + networkFee + exFee, swapAmountOut, "Fees should sum to swap output");

        // Reset swap rate
        mockQuoterBSC.setSwapRate(1e18);
    }

    /**
     * @notice Test getAmountSwapBridgeOutMultihop - multi-hop quote
     */
    function test_getAmountSwapBridgeOutMultihop() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Build path: CROSS -> (fee) -> SwapOutputToken
        bytes memory path = abi.encodePacked(address(cross), uint24(3000), address(swapOutputTokenBSC));

        (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getAmountSwapBridgeOutMultihop(CROSS_CHAIN_ID, path, amountIn);

        // Verify all return values
        assertTrue(ok, "Quote should succeed");

        // With 1:1 swap rate: swapAmountOut = amountIn
        assertEq(swapAmountOut, amountIn, "Swap amount out should equal input with 1:1 rate");

        // With no fees configured: bridgeValue = swapAmountOut
        assertEq(bridgeValue, amountIn, "Bridge value should equal input (no fees, 1:1 rate)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(bridgeValue + networkFee + exFee, swapAmountOut, "Fees should sum to swap output");
    }

    /**
     * @notice Test getAmountSwapBridgeIn - single pool reverse quote
     */
    function test_getAmountSwapBridgeIn() public {
        uint desiredBridgeValue = 500 * 1e18;

        vm.selectFork(bscForkID);

        (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue);

        // Verify all return values
        assertTrue(ok, "Quote should succeed");

        // With no fees: swapAmountOut = desiredBridgeValue
        assertEq(swapAmountOut, desiredBridgeValue, "Swap output should equal bridge value (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(desiredBridgeValue + networkFee + exFee, swapAmountOut, "Bridge value + fees should equal swap output");

        // With 1:1 rate: amountIn = swapAmountOut = desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value (1:1 rate, no fees)");
        assertEq(amountIn, swapAmountOut, "Amount in should equal swap output with 1:1 rate");

        // Cross-verify: if we use this amountIn for getAmountSwapBridgeOut, we should get back the desired bridgeValue
        (bool okVerify, uint swapOutVerify, uint bridgeValueVerify,,) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(okVerify, "Verification quote should succeed");
        assertEq(swapOutVerify, swapAmountOut, "Swap output should match");
        assertEq(bridgeValueVerify, desiredBridgeValue, "Bridge value should match desired");
    }

    /**
     * @notice Test getAmountSwapBridgeInMultihop - multi-hop reverse quote
     */
    function test_getAmountSwapBridgeInMultihop() public {
        uint desiredBridgeValue = 500 * 1e18;

        vm.selectFork(bscForkID);

        // Build path (reversed for exactOutput): SwapOutputToken -> (fee) -> CROSS
        bytes memory path = abi.encodePacked(address(swapOutputTokenBSC), uint24(3000), address(cross));

        (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getAmountSwapBridgeInMultihop(CROSS_CHAIN_ID, path, desiredBridgeValue);

        // Verify all return values
        assertTrue(ok, "Quote should succeed");

        // With no fees: swapAmountOut = desiredBridgeValue
        assertEq(swapAmountOut, desiredBridgeValue, "Swap output should equal bridge value (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(desiredBridgeValue + networkFee + exFee, swapAmountOut, "Bridge value + fees should equal swap output");

        // With 1:1 rate: amountIn = swapAmountOut = desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value (1:1 rate, no fees)");
    }

    /**
     * @notice Test getAmountSwapBridgeIn with zero bridge value
     * @dev If minimumValue is 0, then zero bridgeValue is acceptable
     */
    function test_getAmountSwapBridgeIn_zeroBridgeValue() public {
        // Use zero bridge value
        uint desiredBridgeValue = 0;

        vm.selectFork(bscForkID);

        (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue);

        // With zero bridgeValue and zero minimumValue, the quote succeeds with zero amounts
        assertTrue(ok, "Quote should succeed when minimumValue is 0");
        assertEq(amountIn, 0, "Amount in should be 0");
        assertEq(swapAmountOut, 0, "Swap output should be 0");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
    }

    /**
     * @notice Test quote consistency - verify getAmountSwapBridgeIn and getAmountSwapBridgeOut are inverses
     */
    function test_quoteConsistency() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // First get the output for a given input
        (bool ok1, uint swapAmountOut, uint bridgeValue,,) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(ok1, "First quote should succeed");

        // Now get the required input for that bridge value
        (bool ok2, uint requiredAmountIn, uint swapAmountOut2,,) = swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, bridgeValue
        );
        assertTrue(ok2, "Second quote should succeed");

        // With 1:1 rate, the required input should equal the swap output needed
        // and swap outputs should match
        assertEq(swapAmountOut, swapAmountOut2, "Swap outputs should match");
        assertEq(requiredAmountIn, swapAmountOut2, "Required input should equal swap output with 1:1 rate");
    }

    // ============ Enhanced Existing Test Verifications ============

    /**
     * @notice Test swapExactInputSingleAndBridge with full return value verification
     */
    function test_swapExactInputSingleAndBridge_fullVerification() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Get expected values first
        (bool expectedOk, uint expectedSwapOut, uint expectedBridgeValue, uint expectedNetworkFee, uint expectedExFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(expectedOk, "Quote should succeed");

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeUserCrossBalance = cross.balanceOf(USER);
        uint beforeBridgeOutputBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        // Build params
        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap and bridge
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify swap output matches expected
        assertEq(amountOut, expectedSwapOut, "Swap output should match quote");

        // Verify user balance decreased correctly
        assertEq(
            cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User CROSS balance should decrease by amountIn"
        );

        // Verify bridge received the expected total amount (bridgeValue + networkFee + exFee)
        uint bridgeBalanceIncrease = swapOutputTokenBSC.balanceOf(address(bridgeBSC)) - beforeBridgeOutputBalance;
        assertEq(
            bridgeBalanceIncrease,
            expectedBridgeValue + expectedNetworkFee + expectedExFee,
            "Bridge should receive swap output"
        );

        // Verify fee breakdown
        assertEq(expectedBridgeValue + expectedNetworkFee + expectedExFee, amountOut, "Fees should sum to swap output");
    }

    /**
     * @notice Test quoter immutable state
     */
    function test_quoterImmutable() public {
        vm.selectFork(bscForkID);

        // Verify quoter is set correctly
        assertEq(address(swapBridgeRouterBSC.quoter()), address(mockQuoterBSC), "Quoter should be set correctly");
    }

    // ============ Tests with Fees Configured ============

    /**
     * @notice Test getAmountSwapBridgeOut with exchange fee configured
     * @dev Sets 1% exchange fee (100 / 10000 = 1%)
     */
    function test_getAmountSwapBridgeOut_withExFee() public {
        uint amountIn = 1000 * 1e18;
        uint exFeeRate = 100; // 1% (denominator is 10000)
        uint denominator = 10000;

        vm.selectFork(bscForkID);

        // Set exchange fee rate for swapOutputToken
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);

        // Verify quote succeeded
        assertTrue(ok, "Quote should succeed");

        // With 1:1 swap rate: swapAmountOut = amountIn
        assertEq(swapAmountOut, amountIn, "Swap amount out should equal input with 1:1 rate");

        // Verify fee calculation
        // bridgeValue = swapAmountOut * denominator / (denominator + exFeeRate)
        uint expectedBridgeValue = (swapAmountOut * denominator) / (denominator + exFeeRate);
        uint expectedExFee = (expectedBridgeValue * exFeeRate) / denominator;

        assertEq(bridgeValue, expectedBridgeValue, "Bridge value should match calculated");
        assertEq(networkFee, 0, "Network fee should be 0 (gas price not set)");
        assertEq(exFee, expectedExFee, "Exchange fee should be 1% of bridge value");

        // Verify total approximately equals swapAmountOut (allow 1 wei rounding error)
        assertApproxEqAbs(
            bridgeValue + networkFee + exFee, swapAmountOut, 1, "Fees should sum to swap output (1 wei tolerance)"
        );

        // Verify exFee is approximately 1% of bridgeValue
        assertApproxEqRel(exFee, bridgeValue / 100, 0.01e18, "Exchange fee should be ~1% of bridge value");

        // Clean up: reset fee rate
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test getAmountSwapBridgeIn with exchange fee configured
     * @dev Sets 1% exchange fee and verifies reverse calculation
     */
    function test_getAmountSwapBridgeIn_withExFee() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint exFeeRate = 100; // 1% (denominator is 10000)
        uint denominator = 10000;

        vm.selectFork(bscForkID);

        // Set exchange fee rate for swapOutputToken
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue);

        // Verify quote succeeded
        assertTrue(ok, "Quote should succeed");

        // Calculate expected values
        // exFee = desiredBridgeValue * exFeeRate / denominator
        uint expectedExFee = (desiredBridgeValue * exFeeRate) / denominator;
        uint expectedSwapAmountOut = desiredBridgeValue + expectedExFee + 0; // networkFee = 0

        assertEq(swapAmountOut, expectedSwapAmountOut, "Swap output should be bridgeValue + fees");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, expectedExFee, "Exchange fee should be 1% of bridge value");

        // With 1:1 rate: amountIn = swapAmountOut
        assertEq(amountIn, swapAmountOut, "Amount in should equal swap output with 1:1 rate");

        // Cross-verify: using this amountIn should give back the desired bridgeValue
        (bool okVerify,, uint bridgeValueVerify,,) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(okVerify, "Verification quote should succeed");
        assertEq(bridgeValueVerify, desiredBridgeValue, "Bridge value should match desired");

        // Clean up: reset fee rate
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test getAmountSwapBridgeOut with exchange fee only
     * @dev Network fee requires complex PriceFeed setup, so we test exFee separately
     */
    function test_getAmountSwapBridgeOut_feeBreakdown() public {
        uint amountIn = 1000 * 1e18;
        uint exFeeRate = 200; // 2%
        uint denominator = 10000;

        vm.selectFork(bscForkID);

        // Set exchange fee rate
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);

        assertTrue(ok, "Quote should succeed");
        assertEq(swapAmountOut, amountIn, "Swap output should equal input with 1:1 rate");

        // Verify exchange fee is approximately 2% of bridgeValue
        assertGt(exFee, 0, "Exchange fee should be positive");
        assertLt(bridgeValue, swapAmountOut, "Bridge value should be less than swap output");
        assertEq(networkFee, 0, "Network fee should be 0 (gas price not configured)");

        // bridgeValue + exFee + networkFee should equal swapAmountOut
        assertApproxEqAbs(
            bridgeValue + exFee + networkFee, swapAmountOut, 1, "bridgeValue + fees should equal swapAmountOut"
        );

        // Verify the ratio: exFee / bridgeValue ≈ exFeeRate / denominator
        uint actualRatio = (exFee * denominator) / bridgeValue;
        assertApproxEqAbs(actualRatio, exFeeRate, 1, "Fee ratio should match exFeeRate");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test actual swap+bridge execution with fees configured
     */
    function test_swapExactInputSingleAndBridge_withFees() public {
        uint amountIn = 1000 * 1e18;
        uint exFeeRate = 100; // 1%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Get expected values with fees
        (bool expectedOk, uint expectedSwapOut, uint expectedBridgeValue, uint expectedNetworkFee, uint expectedExFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(expectedOk, "Quote should succeed");

        // Verify expected values have fees applied
        assertGt(expectedExFee, 0, "Expected exchange fee should be positive");
        assertEq(expectedNetworkFee, 0, "Expected network fee should be 0 (gas price not configured)");
        assertLt(expectedBridgeValue, expectedSwapOut, "Expected bridge value should be less than swap output");

        // Approve and execute
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeBridgeBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        ISwapBridgeRouter.ExactInputSingleAndBridgeParams memory params = ISwapBridgeRouter
            .ExactInputSingleAndBridgeParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.SwapAndBridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapExactInputSingleAndBridge(params, block.timestamp + 1 hours);

        // Verify swap output matches quote
        assertEq(amountOut, expectedSwapOut, "Swap output should match quote");

        // Bridge receives totalAmount but sends fees to _dev, so bridge balance only increases by bridgeValue
        // See BaseBridge._initiateBridge: token.safeTransferFrom(from, this, value+fee); token.safeTransfer(_dev, fee);
        uint bridgeBalanceIncrease = swapOutputTokenBSC.balanceOf(address(bridgeBSC)) - beforeBridgeBalance;
        assertEq(
            bridgeBalanceIncrease,
            expectedBridgeValue,
            "Bridge balance should increase by bridgeValue only (fees go to _dev)"
        );

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test quote consistency with fees - verify In and Out are inverses
     * @dev getAmountSwapBridgeOut: given amountIn, what bridgeValue do I get?
     *      getAmountSwapBridgeIn: given bridgeValue, what amountIn do I need?
     *      These should be inverses: In(Out(x).bridgeValue) should give back x
     */
    function test_quoteConsistency_withFees() public {
        uint amountIn = 1000 * 1e18;
        uint exFeeRate = 100; // 1%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Step 1: Get bridgeValue for given amountIn
        (bool ok1, uint swapAmountOut1, uint bridgeValue1, uint networkFee1, uint exFee1) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(ok1, "First quote should succeed");
        assertGt(exFee1, 0, "Exchange fee should be positive");

        // With 1:1 rate: swapAmountOut = amountIn
        assertEq(swapAmountOut1, amountIn, "Swap output should equal input with 1:1 rate");

        // Step 2: Get required amountIn for the same bridgeValue
        (bool ok2, uint requiredAmountIn, uint swapAmountOut2, uint networkFee2, uint exFee2) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, bridgeValue1);
        assertTrue(ok2, "Second quote should succeed");

        // Key verification: requiredAmountIn should equal original amountIn
        // Because: if we need bridgeValue1, we need swapAmountOut2 = bridgeValue1 + fees
        // And with 1:1 rate, requiredAmountIn = swapAmountOut2
        // swapAmountOut2 = bridgeValue1 + networkFee2 + exFee2
        assertEq(swapAmountOut2, bridgeValue1 + networkFee2 + exFee2, "Swap output should be bridgeValue + fees");

        // With 1:1 rate: requiredAmountIn = swapAmountOut2
        assertEq(requiredAmountIn, swapAmountOut2, "Required input should equal swap output with 1:1 rate");

        // The fees calculated for same bridgeValue should match
        assertEq(networkFee2, networkFee1, "Network fees should match for same bridgeValue");
        assertEq(exFee2, exFee1, "Exchange fees should match for same bridgeValue");

        // Round-trip verification with rounding tolerance
        // Due to integer division: bridgeValue = swapAmountOut * 10000 / 10100
        // Then: swapAmountOut2 = bridgeValue + fees may lose 1 wei
        assertApproxEqAbs(requiredAmountIn, amountIn, 1, "Round-trip should return original amountIn (1 wei tolerance)");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test different fee rates
     */
    function test_getAmountSwapBridgeOut_variousFeeRates() public {
        uint amountIn = 1000 * 1e18;
        uint denominator = 10000;
        uint feeRate1 = 50; // 0.5%
        uint feeRate2 = 500; // 5%

        vm.selectFork(bscForkID);

        // Test with 0.5% fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, feeRate1);

        (bool ok1,, uint bridgeValue1,, uint exFee1) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(ok1, "Quote should succeed");
        assertApproxEqAbs(bridgeValue1 + exFee1, amountIn, 1, "Total should equal input");

        // Verify fee ratio for 0.5%
        uint actualRatio1 = (exFee1 * denominator) / bridgeValue1;
        assertApproxEqAbs(actualRatio1, feeRate1, 1, "Fee ratio should match 0.5%");

        // Test with 5% fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, feeRate2);

        (bool ok2,, uint bridgeValue2,, uint exFee2) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(ok2, "Quote should succeed");
        assertApproxEqAbs(bridgeValue2 + exFee2, amountIn, 1, "Total should equal input");

        // Verify fee ratio for 5%
        uint actualRatio2 = (exFee2 * denominator) / bridgeValue2;
        assertApproxEqAbs(actualRatio2, feeRate2, 1, "Fee ratio should match 5%");

        // Higher fee rate should result in lower bridgeValue
        assertLt(bridgeValue2, bridgeValue1, "Higher fee should result in lower bridge value");
        assertGt(exFee2, exFee1, "Higher fee rate should result in higher fee");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }
}
