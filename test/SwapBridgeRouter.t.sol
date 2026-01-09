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
    uint public partialConsumeRate = 1e18; // 100% consume by default (1e18 = 100%)

    constructor(address _weth9) {
        WETH9 = _weth9;
    }

    function setSwapRate(uint _rate) external {
        swapRate = _rate;
    }

    /// @notice Set partial consumption rate (1e18 = 100%, 5e17 = 50%)
    function setPartialConsumeRate(uint _rate) external {
        partialConsumeRate = _rate;
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
        // Simulate partial consumption when sqrtPriceLimitX96 is non-zero
        uint consumedAmount =
            params.sqrtPriceLimitX96 != 0 ? params.amountIn * partialConsumeRate / 1e18 : params.amountIn;

        // Transfer tokenIn from sender (only consumed amount)
        IERC20(params.tokenIn).transferFrom(msg.sender, address(this), consumedAmount);

        // Calculate output amount based on consumed input
        amountOut = consumedAmount * swapRate / 1e18;
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
     * @notice Test swapBridgeExactInputSingle
     */
    function test_swapBridgeExactInputSingle() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter to spend CROSS tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeUserCrossBalance = cross.balanceOf(USER);
        uint beforeBridgeOutputBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        // Build params
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap and bridge
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

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
     * @notice Test swapBridgeExactInput (multi-hop)
     */
    function test_swapBridgeExactInput() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        // Build path: CROSS -> (fee) -> SwapOutputToken
        bytes memory path = abi.encodePacked(address(cross), uint24(3000), address(swapOutputTokenBSC));

        ISwapBridgeRouter.SwapBridgeExactInputParams memory params = ISwapBridgeRouter.SwapBridgeExactInputParams({
            path: path,
            amountIn: amountIn,
            amountOutMinimum: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        // Execute
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInput(params, block.timestamp + 1 hours);

        // Verify
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User CROSS balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapBridgeExactOutputSingle
     * @dev User specifies exact bridgeValue they want to receive
     */
    function test_swapBridgeExactOutputSingle() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // Verify: with 1:1 rate and no fees, amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value with 1:1 rate and no fees");
        // User should have been refunded excess tokens
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn");
    }

    /**
     * @notice Test swapBridgeExactOutput (multi-hop)
     * @dev User specifies exact bridgeValue they want to receive
     */
    function test_swapBridgeExactOutput() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        // Build path (reversed for exactOutput): SwapOutputToken -> (fee) -> CROSS
        bytes memory path = abi.encodePacked(address(swapOutputTokenBSC), uint24(3000), address(cross));

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.SwapBridgeExactOutputParams memory params = ISwapBridgeRouter.SwapBridgeExactOutputParams({
            path: path,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutput(params, block.timestamp + 1 hours);

        // Verify: with 1:1 rate and no fees, amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value with 1:1 rate and no fees");
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn");
    }

    /**
     * @notice Test swapBridgeExactInputSingleETH
     */
    function test_swapBridgeExactInputSingleETH() public {
        uint amountIn = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute with ETH
        vm.prank(USER);
        uint amountOut =
            swapBridgeRouterBSC.swapBridgeExactInputSingleETH{value: amountIn}(params, block.timestamp + 1 hours);

        // Verify
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User ETH balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapBridgeExactInputETH (multi-hop with ETH)
     */
    function test_swapBridgeExactInputETH() public {
        uint amountIn = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        // Build path: WETH -> (fee) -> SwapOutputToken
        bytes memory path = abi.encodePacked(address(mockWETHBSC), uint24(3000), address(swapOutputTokenBSC));

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.SwapBridgeExactInputParams memory params = ISwapBridgeRouter.SwapBridgeExactInputParams({
            path: path,
            amountIn: amountIn,
            amountOutMinimum: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputETH{value: amountIn}(params, block.timestamp + 1 hours);

        // Verify
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User ETH balance should decrease");
        assertEq(amountOut, amountIn, "Amount out should equal amount in with 1:1 rate");
    }

    /**
     * @notice Test swapBridgeExactOutputSingleETH
     * @dev User specifies exact bridgeValue they want to receive with ETH input
     */
    function test_swapBridgeExactOutputSingleETH() public {
        uint desiredBridgeValue = 0.5 ether;
        uint amountInMaximum = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingleETH{value: amountInMaximum}(
            params, block.timestamp + 1 hours
        );

        // Verify: with 1:1 rate and no fees, amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value with 1:1 rate and no fees");
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User should be refunded excess ETH");
    }

    /**
     * @notice Test swapBridgeExactOutputETH (multi-hop)
     * @dev User specifies exact bridgeValue they want to receive with ETH input
     */
    function test_swapBridgeExactOutputETH() public {
        uint desiredBridgeValue = 0.5 ether;
        uint amountInMaximum = 1 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        // Build path (reversed for exactOutput): SwapOutputToken -> (fee) -> WETH
        bytes memory path = abi.encodePacked(address(swapOutputTokenBSC), uint24(3000), address(mockWETHBSC));

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.SwapBridgeExactOutputParams memory params = ISwapBridgeRouter.SwapBridgeExactOutputParams({
            path: path,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn =
            swapBridgeRouterBSC.swapBridgeExactOutputETH{value: amountInMaximum}(params, block.timestamp + 1 hours);

        // Verify: with 1:1 rate and no fees, amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value with 1:1 rate and no fees");
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User should be refunded excess ETH");
    }

    /**
     * @notice Test getExpectedBridgeAmount view function
     */
    function test_getExpectedBridgeAmount() public {
        uint totalAmount = 1000 * 1e18;

        vm.selectFork(bscForkID);

        (ISwapBridgeRouter.QuoteStatus status, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, totalAmount);

        // Verify ok status
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Should return Success");

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

        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Set deadline to past
        vm.prank(USER);
        vm.expectRevert(SwapBridgeRouter.SBRDeadlineExpired.selector);
        swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp - 1);
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

        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: amountIn * 2, // Expect 2x output
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

        // Verify: should get 2x output
        assertEq(amountOut, amountIn * 2, "Should get 2x output with 2:1 rate");
    }

    /**
     * @notice Test SwapBridge event includes correct initiateIndex
     */
    function test_swapAndBridgeEvent_emitsInitiateIndex() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Get the expected initiateIndex before the swap
        uint expectedInitiateIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);

        // Approve SwapBridgeRouter to spend CROSS tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Expect the SwapBridge event with correct initiateIndex
        vm.expectEmit(true, true, true, true);
        emit ISwapBridgeRouter.SwapBridge(
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
        swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

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

        (ISwapBridgeRouter.QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );

        // Verify all return values
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // With 1:1 swap rate: swapAmountOut = amountIn
        assertEq(swapAmountOut, amountIn, "Swap amount out should equal input with 1:1 rate");

        // With no fees configured: bridgeValue = swapAmountOut, networkFee = 0, exFee = 0
        assertEq(bridgeValue, swapAmountOut, "Bridge value should equal swap output (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(bridgeValue + networkFee + exFee, swapAmountOut, "Fees should sum to swap output");

        // Verify bridge value calculation matches getExpectedBridgeAmount
        (
            ISwapBridgeRouter.QuoteStatus expectedStatus,
            uint expectedBridgeValue,
            uint expectedNetworkFee,
            uint expectedExFee
        ) = swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, swapAmountOut);
        assertTrue(expectedStatus == ISwapBridgeRouter.QuoteStatus.Success, "Expected calculation should succeed");
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

        (ISwapBridgeRouter.QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );

        // Verify
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");
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

        (ISwapBridgeRouter.QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getAmountSwapBridgeOutMultihop(CROSS_CHAIN_ID, path, amountIn);

        // Verify all return values
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

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

        (ISwapBridgeRouter.QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue
        );

        // Verify all return values
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // With no fees: swapAmountOut = desiredBridgeValue
        assertEq(swapAmountOut, desiredBridgeValue, "Swap output should equal bridge value (no fees)");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
        assertEq(desiredBridgeValue + networkFee + exFee, swapAmountOut, "Bridge value + fees should equal swap output");

        // With 1:1 rate: amountIn = swapAmountOut = desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value (1:1 rate, no fees)");
        assertEq(amountIn, swapAmountOut, "Amount in should equal swap output with 1:1 rate");

        // Cross-verify: if we use this amountIn for getAmountSwapBridgeOut, we should get back the desired bridgeValue
        (ISwapBridgeRouter.QuoteStatus statusVerify, uint swapOutVerify, uint bridgeValueVerify,,) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(statusVerify == ISwapBridgeRouter.QuoteStatus.Success, "Verification quote should succeed");
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

        (ISwapBridgeRouter.QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getAmountSwapBridgeInMultihop(CROSS_CHAIN_ID, path, desiredBridgeValue);

        // Verify all return values
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

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

        (ISwapBridgeRouter.QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue
        );

        // With zero bridgeValue and zero minimumValue, the quote succeeds with zero amounts
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed when minimumValue is 0");
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
        (ISwapBridgeRouter.QuoteStatus status1, uint swapAmountOut, uint bridgeValue,,) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(status1 == ISwapBridgeRouter.QuoteStatus.Success, "First quote should succeed");

        // Now get the required input for that bridge value
        (ISwapBridgeRouter.QuoteStatus status2, uint requiredAmountIn, uint swapAmountOut2,,) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, bridgeValue);
        assertTrue(status2 == ISwapBridgeRouter.QuoteStatus.Success, "Second quote should succeed");

        // With 1:1 rate, the required input should equal the swap output needed
        // and swap outputs should match
        assertEq(swapAmountOut, swapAmountOut2, "Swap outputs should match");
        assertEq(requiredAmountIn, swapAmountOut2, "Required input should equal swap output with 1:1 rate");
    }

    // ============ Enhanced Existing Test Verifications ============

    /**
     * @notice Test swapBridgeExactInputSingle with full return value verification
     */
    function test_swapBridgeExactInputSingle_fullVerification() public {
        uint amountIn = 1000 * 1e18;

        vm.selectFork(bscForkID);

        // Get expected values first
        (
            ISwapBridgeRouter.QuoteStatus expectedStatus,
            uint expectedSwapOut,
            uint expectedBridgeValue,
            uint expectedNetworkFee,
            uint expectedExFee
        ) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(expectedStatus == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeUserCrossBalance = cross.balanceOf(USER);
        uint beforeBridgeOutputBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        // Build params
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap and bridge
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

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

        (ISwapBridgeRouter.QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );

        // Verify quote succeeded
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

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

        (ISwapBridgeRouter.QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue
        );

        // Verify quote succeeded
        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

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
        (ISwapBridgeRouter.QuoteStatus statusVerify,, uint bridgeValueVerify,,) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(statusVerify == ISwapBridgeRouter.QuoteStatus.Success, "Verification quote should succeed");
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

        (ISwapBridgeRouter.QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );

        assertTrue(status == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");
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
    function test_swapBridgeExactInputSingle_withFees() public {
        uint amountIn = 1000 * 1e18;
        uint exFeeRate = 100; // 1%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Get expected values with fees
        (
            ISwapBridgeRouter.QuoteStatus expectedStatus,
            uint expectedSwapOut,
            uint expectedBridgeValue,
            uint expectedNetworkFee,
            uint expectedExFee
        ) = swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(expectedStatus == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // Verify expected values have fees applied
        assertGt(expectedExFee, 0, "Expected exchange fee should be positive");
        assertEq(expectedNetworkFee, 0, "Expected network fee should be 0 (gas price not configured)");
        assertLt(expectedBridgeValue, expectedSwapOut, "Expected bridge value should be less than swap output");

        // Approve and execute
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        uint beforeBridgeBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

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

        // Ensure no rounding dust is trapped in the router contract
        assertEq(
            swapOutputTokenBSC.balanceOf(address(swapBridgeRouterBSC)),
            0,
            "SwapBridgeRouter should not retain tokenOut dust after bridging"
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
        (ISwapBridgeRouter.QuoteStatus status1, uint swapAmountOut1, uint bridgeValue1, uint networkFee1, uint exFee1) =
        swapBridgeRouterBSC.getAmountSwapBridgeOut(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn
        );
        assertTrue(status1 == ISwapBridgeRouter.QuoteStatus.Success, "First quote should succeed");
        assertGt(exFee1, 0, "Exchange fee should be positive");

        // With 1:1 rate: swapAmountOut = amountIn
        assertEq(swapAmountOut1, amountIn, "Swap output should equal input with 1:1 rate");

        // Step 2: Get required amountIn for the same bridgeValue
        (
            ISwapBridgeRouter.QuoteStatus status2,
            uint requiredAmountIn,
            uint swapAmountOut2,
            uint networkFee2,
            uint exFee2
        ) = swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, bridgeValue1
        );
        assertTrue(status2 == ISwapBridgeRouter.QuoteStatus.Success, "Second quote should succeed");

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

        (ISwapBridgeRouter.QuoteStatus status1,, uint bridgeValue1,, uint exFee1) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(status1 == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");
        assertApproxEqAbs(bridgeValue1 + exFee1, amountIn, 1, "Total should equal input");

        // Verify fee ratio for 0.5%
        uint actualRatio1 = (exFee1 * denominator) / bridgeValue1;
        assertApproxEqAbs(actualRatio1, feeRate1, 1, "Fee ratio should match 0.5%");

        // Test with 5% fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, feeRate2);

        (ISwapBridgeRouter.QuoteStatus status2,, uint bridgeValue2,, uint exFee2) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, amountIn);
        assertTrue(status2 == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");
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

    // ============ swapExactBridgeValue Tests with Fees ============

    /**
     * @notice Test swapBridgeExactOutputSingle with fees configured
     * @dev Verifies user receives EXACTLY the specified bridgeValue
     */
    function test_swapBridgeExactOutputSingle_withFees() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;
        uint exFeeRate = 100; // 1%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Get expected amountIn from quote
        (ISwapBridgeRouter.QuoteStatus quoteStatus, uint expectedAmountIn,,, uint expectedExFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue);
        assertTrue(quoteStatus == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");
        assertGt(expectedExFee, 0, "Expected exchange fee should be positive");

        // Approve SwapBridgeRouter
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        uint beforeUserCrossBalance = cross.balanceOf(USER);
        uint beforeBridgeBalance = swapOutputTokenBSC.balanceOf(address(bridgeBSC));

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // Verify amountIn matches quote
        assertEq(amountIn, expectedAmountIn, "Actual amountIn should match quote");

        // Verify user spent correct amount
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should spend amountIn");

        // Verify bridge balance increase equals desiredBridgeValue (fees go to _dev)
        uint bridgeBalanceIncrease = swapOutputTokenBSC.balanceOf(address(bridgeBSC)) - beforeBridgeBalance;
        assertEq(bridgeBalanceIncrease, desiredBridgeValue, "Bridge should receive exactly desiredBridgeValue");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test swapBridgeExactOutputSingle verifies exact bridgeValue in event
     */
    function test_swapBridgeExactOutputSingle_eventVerification() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;
        uint exFeeRate = 100; // 1%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Get expected values from quote
        (
            ISwapBridgeRouter.QuoteStatus quoteStatus,
            uint expectedAmountIn,
            uint expectedSwapOutput,
            uint expectedNetworkFee,
            uint expectedExFee
        ) = swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue
        );
        assertTrue(quoteStatus == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // Get expected initiateIndex
        uint expectedInitiateIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);

        // Approve
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Expect event with exact bridgeValue
        vm.expectEmit(true, true, true, true);
        emit ISwapBridgeRouter.SwapBridge(
            USER,
            address(cross),
            address(swapOutputTokenBSC),
            expectedAmountIn,
            expectedSwapOutput,
            CROSS_CHAIN_ID,
            expectedInitiateIndex,
            desiredBridgeValue, // KEY: bridgeValue should be EXACTLY what user specified
            expectedNetworkFee,
            expectedExFee
        );

        // Execute
        vm.prank(USER);
        swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    /**
     * @notice Test swapBridgeExactOutputSingle refunds excess input
     */
    function test_swapExactBridgeValue_refundsExcess() public {
        uint desiredBridgeValue = 100 * 1e18;
        uint amountInMaximum = 1000 * 1e18; // Way more than needed

        vm.selectFork(bscForkID);

        // Approve
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // With 1:1 rate and no fees: amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value");

        // User should be refunded excess (amountInMaximum - amountIn)
        uint expectedRefund = amountInMaximum - amountIn;
        assertEq(
            cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn, excess refunded"
        );
        assertEq(expectedRefund, 900 * 1e18, "Refund should be 900 tokens");
    }

    /**
     * @notice Test swapBridgeExactOutputSingleETH refunds excess ETH
     */
    function test_swapExactBridgeValueETH_refundsExcess() public {
        uint desiredBridgeValue = 0.1 ether;
        uint amountInMaximum = 1 ether; // Way more than needed

        vm.selectFork(bscForkID);
        vm.deal(USER, 10 ether);

        uint beforeUserETHBalance = USER.balance;

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingleETH{value: amountInMaximum}(
            params, block.timestamp + 1 hours
        );

        // With 1:1 rate and no fees: amountIn should equal desiredBridgeValue
        assertEq(amountIn, desiredBridgeValue, "Amount in should equal bridge value");

        // User should be refunded excess ETH
        assertEq(USER.balance, beforeUserETHBalance - amountIn, "User should be refunded excess ETH");
    }

    /**
     * @notice Test swapExactBridgeValue with different swap rates
     */
    function test_swapExactBridgeValue_differentSwapRate() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;
        uint swapRate = 2 * 1e18; // 1 input = 2 output

        vm.selectFork(bscForkID);

        // Set swap rate (2:1 output/input ratio)
        mockSwapRouterBSC.setSwapRate(swapRate);
        mockQuoterBSC.setSwapRate(swapRate);

        // With 2:1 rate and no fees: to get 500 bridgeValue, need 250 input
        uint expectedAmountIn = desiredBridgeValue * 1e18 / swapRate; // 250 * 1e18

        // Approve
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        uint beforeUserCrossBalance = cross.balanceOf(USER);

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute
        vm.prank(USER);
        uint amountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // Verify: with 2:1 rate, should only need 250 input for 500 bridgeValue
        assertEq(amountIn, expectedAmountIn, "Amount in should be half of bridge value with 2:1 rate");
        assertEq(cross.balanceOf(USER), beforeUserCrossBalance - amountIn, "User should only spend amountIn");

        // Reset swap rate
        mockSwapRouterBSC.setSwapRate(1e18);
        mockQuoterBSC.setSwapRate(1e18);
    }

    /**
     * @notice Test swapExactBridgeValue consistency - quote matches execution
     */
    function test_swapExactBridgeValue_quoteMatchesExecution() public {
        uint desiredBridgeValue = 500 * 1e18;
        uint amountInMaximum = 1000 * 1e18;
        uint exFeeRate = 200; // 2%

        vm.selectFork(bscForkID);

        // Set exchange fee
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, exFeeRate);

        // Get quote first
        (
            ISwapBridgeRouter.QuoteStatus quoteStatus,
            uint quotedAmountIn,
            uint quotedSwapOutput,
            uint quotedNetworkFee,
            uint quotedExFee
        ) = swapBridgeRouterBSC.getAmountSwapBridgeIn(
            CROSS_CHAIN_ID, address(cross), address(swapOutputTokenBSC), 3000, desiredBridgeValue
        );
        assertTrue(quoteStatus == ISwapBridgeRouter.QuoteStatus.Success, "Quote should succeed");

        // Verify quote math: swapOutput = bridgeValue + networkFee + exFee
        assertEq(
            quotedSwapOutput,
            desiredBridgeValue + quotedNetworkFee + quotedExFee,
            "Swap output should equal bridgeValue + fees"
        );

        // Approve and execute
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountInMaximum);

        ISwapBridgeRouter.SwapBridgeExactOutputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactOutputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountOut: desiredBridgeValue,
            amountInMaximum: amountInMaximum,
            sqrtPriceLimitX96: 0,
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        vm.prank(USER);
        uint actualAmountIn = swapBridgeRouterBSC.swapBridgeExactOutputSingle(params, block.timestamp + 1 hours);

        // Verify execution matches quote
        assertEq(actualAmountIn, quotedAmountIn, "Actual amountIn should match quoted amountIn");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(swapOutputTokenBSC, 0);
    }

    // ============ QuoteStatus Tests ============

    /**
     * @notice Test getExpectedBridgeAmount returns NoPair for unregistered token
     */
    function test_getExpectedBridgeAmount_NoPair() public {
        vm.selectFork(bscForkID);

        // Use an unregistered token address
        address unregisteredToken = address(0x1234567890123456789012345678901234567890);
        uint totalAmount = 1000 * 1e18;

        (ISwapBridgeRouter.QuoteStatus status, uint bridgeValue, uint networkFee, uint exFee) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, IERC20(unregisteredToken), totalAmount);

        assertEq(uint(status), uint(ISwapBridgeRouter.QuoteStatus.NoPair), "Should return NoPair status");
        assertEq(bridgeValue, 0, "Bridge value should be 0");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
    }

    /**
     * @notice Test getExpectedBridgeAmount returns InsufficientForFee when amount <= networkFee
     */
    function test_getExpectedBridgeAmount_InsufficientForFee() public {
        vm.selectFork(bscForkID);

        // Get the current network fee
        (, uint networkFeeAmount,) = bridgeVerifierBSC.getTokenConfig(CROSS_CHAIN_ID, swapOutputTokenBSC);

        // If network fee is 0, set a high finalize gas to generate one
        if (networkFeeAmount == 0) {
            vm.prank(OWNER);
            bridgeVerifierBSC.setFinalizeBridgeGas(1e15); // High gas value

            // Check again
            (, networkFeeAmount,) = bridgeVerifierBSC.getTokenConfig(CROSS_CHAIN_ID, swapOutputTokenBSC);
        }

        // Skip test if still 0 (no gas price configured)
        if (networkFeeAmount == 0) return;

        // Try with amount equal to network fee (should fail since we need > networkFee)
        uint totalAmount = networkFeeAmount;

        (ISwapBridgeRouter.QuoteStatus status, uint bridgeValue, uint networkFee,) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, totalAmount);

        assertEq(
            uint(status),
            uint(ISwapBridgeRouter.QuoteStatus.InsufficientForFee),
            "Should return InsufficientForFee status"
        );
        assertEq(bridgeValue, 0, "Bridge value should be 0");
        assertTrue(networkFee > 0, "Network fee should be positive");
    }

    /**
     * @notice Test getExpectedBridgeAmount returns InsufficientValue when bridgeValue < minimumValue
     */
    function test_getExpectedBridgeAmount_InsufficientValue() public {
        vm.selectFork(bscForkID);

        // Save original minimum
        uint originalMinimum = bridgeVerifierBSC.getMinimumTokenValue();

        // Set a high minimum token value (in USD, with 8 decimals)
        // 1000 USD minimum = 1000 * 1e8
        vm.prank(OWNER);
        bridgeVerifierBSC.setMinimumTokenValue(1000 * 1e8); // $1000 minimum

        // Get the actual minimum value in tokens
        (uint minimumTokenAmount,,) = bridgeVerifierBSC.getTokenConfig(CROSS_CHAIN_ID, swapOutputTokenBSC);

        // Skip test if minimum is 0 (might happen if token has no price)
        if (minimumTokenAmount == 0) {
            vm.prank(OWNER);
            bridgeVerifierBSC.setMinimumTokenValue(originalMinimum);
            return;
        }

        // Try with amount that results in bridgeValue < minimumValue
        uint totalAmount = minimumTokenAmount - 1;

        (ISwapBridgeRouter.QuoteStatus status, uint bridgeValue,,) =
            swapBridgeRouterBSC.getExpectedBridgeAmount(CROSS_CHAIN_ID, swapOutputTokenBSC, totalAmount);

        assertEq(
            uint(status),
            uint(ISwapBridgeRouter.QuoteStatus.InsufficientValue),
            "Should return InsufficientValue status"
        );
        assertTrue(bridgeValue < minimumTokenAmount, "Bridge value should be below minimum");

        // Clean up
        vm.prank(OWNER);
        bridgeVerifierBSC.setMinimumTokenValue(originalMinimum);
    }

    /**
     * @notice Test getAmountSwapBridgeOut returns NoPair when token is not registered
     * @dev Quoter succeeds but bridge amount calculation returns NoPair
     */
    function test_getAmountSwapBridgeOut_NoPair() public {
        vm.selectFork(bscForkID);

        // Use a token that is not registered
        address unregisteredToken = address(0xdEADbeEF00000000000000000000000000000001);
        uint amountIn = 1000 * 1e18;

        (ISwapBridgeRouter.QuoteStatus status,, uint bridgeValue, uint networkFee, uint exFee) = swapBridgeRouterBSC
            .getAmountSwapBridgeOut(CROSS_CHAIN_ID, address(cross), unregisteredToken, 3000, amountIn);

        assertEq(uint(status), uint(ISwapBridgeRouter.QuoteStatus.NoPair), "Should return NoPair status");
        // Note: swapAmountOut may have value from quoter, but bridge values should be 0
        assertEq(bridgeValue, 0, "Bridge value should be 0");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
    }

    /**
     * @notice Test getAmountSwapBridgeIn returns NoPair when token is not registered
     */
    function test_getAmountSwapBridgeIn_NoPair() public {
        vm.selectFork(bscForkID);

        // Use a token that is not registered
        address unregisteredToken = address(0xdEADbeEF00000000000000000000000000000001);
        uint bridgeValue = 1000 * 1e18;

        (ISwapBridgeRouter.QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee) =
        swapBridgeRouterBSC.getAmountSwapBridgeIn(CROSS_CHAIN_ID, address(cross), unregisteredToken, 3000, bridgeValue);

        assertEq(uint(status), uint(ISwapBridgeRouter.QuoteStatus.NoPair), "Should return NoPair status");
        assertEq(amountIn, 0, "Amount in should be 0");
        assertEq(swapAmountOut, 0, "Swap amount out should be 0");
        assertEq(networkFee, 0, "Network fee should be 0");
        assertEq(exFee, 0, "Exchange fee should be 0");
    }

    // ============ Partial Swap Refund Tests ============

    /**
     * @notice Test that unspent ERC20 tokens are refunded when sqrtPriceLimitX96 causes early termination
     */
    function test_swapBridgeExactInputSingle_partialSwap_refundsUnspent() public {
        vm.selectFork(bscForkID);

        uint amountIn = 100 * 1e18;
        uint partialRate = 5e17; // 50% consumption

        // Set partial consumption rate
        mockSwapRouterBSC.setPartialConsumeRate(partialRate);

        // Record initial balances
        uint userBalanceBefore = cross.balanceOf(USER);

        // Approve tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        // Create params with non-zero sqrtPriceLimitX96 to trigger partial consumption
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0, // Accept any output for this test
            sqrtPriceLimitX96: 1, // Non-zero to trigger partial consumption
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

        // Verify only 50% was consumed for swap output
        uint expectedConsumed = amountIn * partialRate / 1e18; // 50 tokens
        assertEq(amountOut, expectedConsumed, "Output should be 50% of input");

        // Verify unspent 50% was refunded to user
        uint userBalanceAfter = cross.balanceOf(USER);
        assertEq(userBalanceBefore - userBalanceAfter, expectedConsumed, "User should only spend consumed amount");

        // Verify router has no leftover tokens
        assertEq(cross.balanceOf(address(swapBridgeRouterBSC)), 0, "Router should have no leftover tokens");

        // Reset partial consumption rate
        mockSwapRouterBSC.setPartialConsumeRate(1e18);
    }

    /**
     * @notice Test that unspent ETH (WETH) is refunded when sqrtPriceLimitX96 causes early termination
     */
    function test_swapBridgeExactInputSingleETH_partialSwap_refundsUnspent() public {
        vm.selectFork(bscForkID);

        uint amountIn = 1 ether;
        uint partialRate = 6e17; // 60% consumption

        // Set partial consumption rate
        mockSwapRouterBSC.setPartialConsumeRate(partialRate);

        // Fund user with ETH
        vm.deal(USER, 10 ether);

        // Record initial balance
        uint userEthBefore = USER.balance;

        // Create params with non-zero sqrtPriceLimitX96 to trigger partial consumption
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(mockWETHBSC),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0, // Accept any output for this test
            sqrtPriceLimitX96: 1, // Non-zero to trigger partial consumption
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap with ETH
        vm.prank(USER);
        uint amountOut =
            swapBridgeRouterBSC.swapBridgeExactInputSingleETH{value: amountIn}(params, block.timestamp + 1 hours);

        // Verify only 60% was consumed for swap output
        uint expectedConsumed = amountIn * partialRate / 1e18;
        assertEq(amountOut, expectedConsumed, "Output should be 60% of input");

        // Verify unspent 40% was refunded to user as ETH
        uint userEthAfter = USER.balance;
        assertEq(userEthBefore - userEthAfter, expectedConsumed, "User should only spend consumed ETH amount");

        // Verify router has no leftover WETH
        assertEq(mockWETHBSC.balanceOf(address(swapBridgeRouterBSC)), 0, "Router should have no leftover WETH");

        // Reset partial consumption rate
        mockSwapRouterBSC.setPartialConsumeRate(1e18);
    }

    /**
     * @notice Test that full consumption works correctly (no refund needed)
     */
    function test_swapBridgeExactInputSingle_fullConsumption_noRefund() public {
        vm.selectFork(bscForkID);

        uint amountIn = 100 * 1e18;

        // Ensure full consumption (default)
        mockSwapRouterBSC.setPartialConsumeRate(1e18);

        // Record initial balance
        uint userBalanceBefore = cross.balanceOf(USER);

        // Approve tokens
        vm.prank(USER);
        cross.approve(address(swapBridgeRouterBSC), amountIn);

        // Create params with zero sqrtPriceLimitX96 (full consumption)
        ISwapBridgeRouter.SwapBridgeExactInputSingleParams memory params = ISwapBridgeRouter
            .SwapBridgeExactInputSingleParams({
            tokenIn: address(cross),
            tokenOut: address(swapOutputTokenBSC),
            fee: 3000,
            amountIn: amountIn,
            amountOutMinimum: 0,
            sqrtPriceLimitX96: 0, // Zero means no price limit, full consumption
            bridgeParams: ISwapBridgeRouter.BridgeParams({toChainID: CROSS_CHAIN_ID, recipient: USER, extraData: ""})
        });

        // Execute swap
        vm.prank(USER);
        uint amountOut = swapBridgeRouterBSC.swapBridgeExactInputSingle(params, block.timestamp + 1 hours);

        // Verify full amount was consumed
        assertEq(amountOut, amountIn, "Output should equal input (1:1 rate)");

        // Verify user spent full amount
        uint userBalanceAfter = cross.balanceOf(USER);
        assertEq(userBalanceBefore - userBalanceAfter, amountIn, "User should spend full amount");

        // Verify router has no leftover tokens
        assertEq(cross.balanceOf(address(swapBridgeRouterBSC)), 0, "Router should have no leftover tokens");
    }
}
