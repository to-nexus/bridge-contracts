// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test, console} from "forge-std/Test.sol";

import {SwapV3Receiver} from "../src/examples/SwapV3Receiver.sol";
import {IBridgeReceiver} from "../src/interface/IBridgeReceiver.sol";

/**
 * @title SwapV3ReceiverTest
 * @notice Tests for SwapV3Receiver functionality
 */
contract SwapV3ReceiverTest is Test {
    SwapV3Receiver public receiver;
    MockBridge public bridge;
    MockSwapRouter public swapRouter;
    MockToken public tokenIn;
    MockToken public tokenOut;

    address public user = address(0x1);
    uint public constant BRIDGE_AMOUNT = 1000 ether;
    uint public constant SWAP_RATE = 95; // 1000 tokenIn = 950 tokenOut (5% loss)

    function setUp() public {
        bridge = new MockBridge();
        swapRouter = new MockSwapRouter();
        tokenIn = new MockToken("Token In", "TIN");
        tokenOut = new MockToken("Token Out", "TOUT");

        receiver = new SwapV3Receiver(address(bridge), address(swapRouter));

        // Mint tokens
        tokenIn.mint(address(bridge), BRIDGE_AMOUNT * 10);
        tokenOut.mint(address(swapRouter), BRIDGE_AMOUNT * 10);

        // Setup swap rate
        swapRouter.setSwapRate(SWAP_RATE);
    }

    function testSuccessfulSwap() public {
        // Arrange
        uint24 fee = 3000;
        uint amountOutMinimum = 900 ether; // 10% slippage tolerance

        bytes memory extraData = abi.encode(user, address(tokenOut), fee, amountOutMinimum);

        // Act
        bridge.finalize(address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        // Assert
        uint expectedOut = BRIDGE_AMOUNT * SWAP_RATE / 100;
        assertEq(tokenOut.balanceOf(user), expectedOut);
        assertEq(tokenIn.balanceOf(address(receiver)), 0); // All swapped
    }

    function testSwapWithDifferentFees() public {
        uint24[3] memory fees = [uint24(500), uint24(3000), uint24(10000)];

        for (uint i = 0; i < fees.length; i++) {
            address testUser = address(uint160(1000 + i));
            bytes memory extraData = abi.encode(testUser, address(tokenOut), fees[i], 900 ether);

            bridge.finalize(address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

            uint expectedOut = BRIDGE_AMOUNT * SWAP_RATE / 100;
            assertEq(tokenOut.balanceOf(testUser), expectedOut);
        }
    }

    function testSwapFailsReturnsOriginalToken() public {
        // Arrange - Set very high minimum output (will fail)
        uint24 fee = 3000;
        uint amountOutMinimum = 1000 ether; // Higher than possible output
        bytes memory extraData = abi.encode(user, address(tokenOut), fee, amountOutMinimum);

        // Act
        bridge.finalize(address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        // Assert - User receives original token
        assertEq(tokenIn.balanceOf(user), BRIDGE_AMOUNT);
        assertEq(tokenOut.balanceOf(user), 0);
    }

    function testOnlyBridgeCanCall() public {
        bytes memory extraData = abi.encode(user, address(tokenOut), uint24(3000), 900 ether);

        // Should fail when called directly
        vm.expectRevert(SwapV3Receiver.SwapV3InvalidBridge.selector);
        receiver.onBridgeReceived(1, 1, tokenIn, BRIDGE_AMOUNT, extraData);
    }

    function testReplayAttackPrevention() public {
        uint24 fee = 3000;
        uint amountOutMinimum = 900 ether;
        bytes memory extraData = abi.encode(user, address(tokenOut), fee, amountOutMinimum);

        // First call succeeds with chainID=1, index=1
        bridge.finalizeWithChainId(1, 1, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        // Second call with same fromChainID and index should fail
        vm.expectRevert(SwapV3Receiver.SwapV3AlreadyProcessed.selector);
        bridge.finalizeWithChainId(1, 1, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);
    }

    function testDifferentChainIdsDontConflict() public {
        uint24 fee = 3000;
        uint amountOutMinimum = 900 ether;
        bytes memory extraData = abi.encode(user, address(tokenOut), fee, amountOutMinimum);

        // Call with different chain IDs should both succeed
        bridge.finalizeWithChainId(1, 1, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);
        bridge.finalizeWithChainId(2, 1, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        // User should receive twice the output
        uint expectedOut = (BRIDGE_AMOUNT * SWAP_RATE / 100) * 2;
        assertEq(tokenOut.balanceOf(user), expectedOut);
    }

    function testSwapRouterFailureReturnsTokens() public {
        // Arrange - Make router revert
        swapRouter.setShouldRevert(true);

        uint24 fee = 3000;
        uint amountOutMinimum = 900 ether;
        bytes memory extraData = abi.encode(user, address(tokenOut), fee, amountOutMinimum);

        // Act
        bridge.finalize(address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        // Assert - Original tokens returned
        assertEq(tokenIn.balanceOf(user), BRIDGE_AMOUNT);
        assertEq(tokenOut.balanceOf(user), 0);
    }

    function testSetSwapRouterOnlyOwner() public {
        address newRouter = address(0x999);

        // Non-owner cannot set
        vm.prank(user);
        vm.expectRevert();
        receiver.setSwapRouter(newRouter);

        // Owner can set
        receiver.setSwapRouter(newRouter);
        assertEq(address(receiver.swapRouter()), newRouter);
    }

    function testEmergencyWithdrawOnlyOwner() public {
        // Send some tokens to receiver
        tokenIn.mint(address(receiver), 100 ether);

        // Non-owner cannot withdraw
        vm.prank(user);
        vm.expectRevert();
        receiver.emergencyWithdraw(tokenIn, 100 ether);

        // Owner can withdraw
        address owner = receiver.owner();
        uint balanceBefore = tokenIn.balanceOf(owner);
        receiver.emergencyWithdraw(tokenIn, 100 ether);
        assertEq(tokenIn.balanceOf(owner), balanceBefore + 100 ether);
    }

    function testMultipleUsersIndependently() public {
        address user1 = address(0x101);
        address user2 = address(0x102);
        address user3 = address(0x103);

        uint24 fee = 3000;
        uint amountOutMinimum = 900 ether;

        // User 1
        bytes memory extraData1 = abi.encode(user1, address(tokenOut), fee, amountOutMinimum);
        bridge.finalizeWithChainId(1, 1, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData1);

        // User 2
        bytes memory extraData2 = abi.encode(user2, address(tokenOut), fee, amountOutMinimum);
        bridge.finalizeWithChainId(1, 2, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData2);

        // User 3
        bytes memory extraData3 = abi.encode(user3, address(tokenOut), fee, amountOutMinimum);
        bridge.finalizeWithChainId(1, 3, address(receiver), tokenIn, BRIDGE_AMOUNT, extraData3);

        // All users receive their swapped tokens
        uint expectedOut = BRIDGE_AMOUNT * SWAP_RATE / 100;
        assertEq(tokenOut.balanceOf(user1), expectedOut);
        assertEq(tokenOut.balanceOf(user2), expectedOut);
        assertEq(tokenOut.balanceOf(user3), expectedOut);
    }

    function testZeroAddressValidation() public {
        // Bridge cannot be zero
        vm.expectRevert(SwapV3Receiver.SwapV3InvalidBridge.selector);
        new SwapV3Receiver(address(0), address(swapRouter));

        // SwapRouter cannot be zero
        vm.expectRevert(SwapV3Receiver.SwapV3InvalidRouter.selector);
        new SwapV3Receiver(address(bridge), address(0));
    }

    function testReturnsCorrectSelector() public {
        bytes memory extraData = abi.encode(user, address(tokenOut), uint24(3000), 900 ether);

        bytes4 selector = bridge.finalizeAndReturnSelector(address(receiver), tokenIn, BRIDGE_AMOUNT, extraData);

        assertEq(selector, IBridgeReceiver.onBridgeReceived.selector);
    }
}

/**
 * @notice Mock bridge for testing
 */
contract MockBridge {
    uint public currentChainId = 1;
    uint public currentIndex = 1;

    function finalize(address recipient, MockToken token, uint amount, bytes memory extraData) external {
        finalizeWithChainId(currentChainId, currentIndex, recipient, token, amount, extraData);
        currentIndex++;
    }

    function finalizeWithChainId(
        uint fromChainID,
        uint index,
        address recipient,
        MockToken token,
        uint amount,
        bytes memory extraData
    ) public {
        // Transfer tokens to recipient
        token.transfer(recipient, amount);

        // Call onBridgeReceived
        IBridgeReceiver(recipient).onBridgeReceived(fromChainID, index, IERC20(address(token)), amount, extraData);
    }

    function finalizeAndReturnSelector(address recipient, MockToken token, uint amount, bytes memory extraData)
        external
        returns (bytes4)
    {
        token.transfer(recipient, amount);
        return IBridgeReceiver(recipient).onBridgeReceived(
            currentChainId, currentIndex, IERC20(address(token)), amount, extraData
        );
    }
}

/**
 * @notice Mock Uniswap V3 SwapRouter
 */
contract MockSwapRouter {
    uint public swapRate = 100; // Default 1:1
    bool public shouldRevert = false;

    function setSwapRate(uint _rate) external {
        swapRate = _rate;
    }

    function setShouldRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function exactInputSingle(ExactInputSingleParams calldata params) external returns (uint amountOut) {
        require(!shouldRevert, "Mock swap failed");

        // Transfer tokenIn from caller
        IERC20(params.tokenIn).transferFrom(msg.sender, address(this), params.amountIn);

        // Calculate output amount
        amountOut = params.amountIn * swapRate / 100;

        // Check minimum
        require(amountOut >= params.amountOutMinimum, "Insufficient output");

        // Transfer tokenOut to recipient
        MockToken tokenOut = MockToken(params.tokenOut);
        tokenOut.transfer(params.recipient, amountOut);
    }

    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint deadline;
        uint amountIn;
        uint amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }
}

/**
 * @notice Mock ERC20 token
 */
contract MockToken is IERC20 {
    string public name;
    string public symbol;
    uint8 public constant decimals = 18;
    uint public totalSupply;

    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

    constructor(string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
    }

    function mint(address to, uint amount) external {
        balanceOf[to] += amount;
        totalSupply += amount;
    }

    function transfer(address to, uint amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        return true;
    }

    function transferFrom(address from, address to, uint amount) external returns (bool) {
        allowance[from][msg.sender] -= amount;
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        return true;
    }

    function approve(address spender, uint amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        return true;
    }
}
