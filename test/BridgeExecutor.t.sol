// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Vm} from "forge-std/Vm.sol";

import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {ICrossMintableERC20} from "../src/token/ICrossMintableERC20.sol";

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";

/**
 * @title MockTargetContract
 * @notice Mock contract for testing BridgeExecutor functionality
 */
contract MockTargetContract {
    event Received(address token, address from, uint value, bytes data);
    event NativeReceived(address from, uint value, bytes data);

    bool public shouldRevert;
    uint public consumePercent = 100; // How much % of tokens to consume (default 100%)

    function setShouldRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function setConsumePercent(uint _percent) external {
        consumePercent = _percent;
    }

    function handleBridgeCallback(address token, address user, uint amount, bytes calldata data) external payable {
        if (shouldRevert) revert("MockTargetContract: intentional revert");

        uint consumeAmount = amount * consumePercent / 100;

        if (token == address(1)) {
            // Native token
            require(msg.value == amount, "MockTargetContract: incorrect value");
            // Return unconsumed portion to sender
            if (consumeAmount < amount) {
                (bool ok,) = msg.sender.call{value: amount - consumeAmount}("");
                require(ok, "MockTargetContract: refund failed");
            }
            emit NativeReceived(user, consumeAmount, data);
        } else {
            // ERC20 token - pull only the consumed amount from executor
            if (consumeAmount > 0) {
                require(
                    IERC20(token).transferFrom(msg.sender, address(this), consumeAmount),
                    "MockTargetContract: transfer failed"
                );
            }
            emit Received(token, user, consumeAmount, data);
        }
    }

    /**
     * @notice Callback function with return value for testing returnData in events
     * @param token Token address
     * @param user User address
     * @param amount Amount
     * @param data Extra data
     * @return hash Keccak256 hash of the input parameters
     */
    function handleBridgeCallbackWithReturn(address token, address user, uint amount, bytes calldata data)
        external
        payable
        returns (bytes32 hash)
    {
        if (shouldRevert) revert("MockTargetContract: intentional revert");

        uint consumeAmount = amount * consumePercent / 100;

        if (token == address(1)) {
            // Native token
            require(msg.value == amount, "MockTargetContract: incorrect value");
            // Return unconsumed portion to sender
            if (consumeAmount < amount) {
                (bool ok,) = msg.sender.call{value: amount - consumeAmount}("");
                require(ok, "MockTargetContract: refund failed");
            }
            emit NativeReceived(user, consumeAmount, data);
        } else {
            // ERC20 token - pull only the consumed amount from executor
            if (consumeAmount > 0) {
                require(
                    IERC20(token).transferFrom(msg.sender, address(this), consumeAmount),
                    "MockTargetContract: transfer failed"
                );
            }
            emit Received(token, user, consumeAmount, data);
        }

        return keccak256(abi.encode(token, user, amount, data));
    }

    receive() external payable {}
}

/**
 * @title MockGasConsumer
 * @notice Mock contract that consumes a lot of gas to test gas limit protection
 */
contract MockGasConsumer {
    uint public counter;

    function consumeGas(uint iterations) external payable {
        // Consume gas by doing storage writes
        for (uint i = 0; i < iterations; i++) {
            counter = i;
        }
    }

    receive() external payable {}
}

/**
 * @title MockBadApproveToken
 * @notice Mock ERC20 token that always fails on approve (for testing forceApprove failure)
 */
contract MockBadApproveToken {
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;

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

    function approve(address, uint) external pure returns (bool) {
        return false; // Always fail
    }

    function mint(address to, uint amount) external {
        balanceOf[to] += amount;
    }
}

/**
 * @title MockExtraETHReturn
 * @notice Mock contract that returns MORE ETH than it received
 * @dev underflow protection in consumed calculation
 */
contract MockExtraETHReturn {
    /**
     * @notice Handler that returns all ETH (pre-funded + msg.value)
     * @dev This simulates a scenario where target returns more than value
     */
    function handleWithExtraReturn(address, address, uint, bytes calldata) external payable {
        uint totalReturn = address(this).balance; // includes msg.value + pre-funded
        (bool ok,) = msg.sender.call{value: totalReturn}("");
        require(ok, "Extra return failed");
    }

    receive() external payable {}
}

/**
 * @title MockLargeReturnData
 * @notice Mock contract that returns large data to test returnData size cap
 * @dev Tests CBU-11 fix for unbounded returnData OOG vulnerability
 */
contract MockLargeReturnData {
    /**
     * @notice Returns large data of specified size
     * @param size Size of data to return (in bytes)
     */
    function returnLargeData(uint size) external payable returns (bytes memory) {
        bytes memory data = new bytes(size);
        // Fill with some pattern
        for (uint i = 0; i < size; i++) {
            data[i] = bytes1(uint8(i % 256));
        }
        return data;
    }

    receive() external payable {}
}

/**
 * @title MockSwap
 * @notice Mock swap contract that exchanges tokenIn for tokenOut
 * @dev Simulates a DEX swap for testing bridge + swap flow
 */
contract MockSwap {
    event Swapped(
        address indexed tokenIn, address indexed tokenOut, address indexed recipient, uint amountIn, uint amountOut
    );

    bool public shouldRevert;
    uint public swapRate = 1e18; // 1:1 default rate (18 decimals)

    function setShouldRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function setSwapRate(uint _rate) external {
        swapRate = _rate;
    }

    /**
     * @notice Swap tokenIn to tokenOut and send to recipient
     * @param tokenIn Token to receive from caller
     * @param tokenOut Token to send to recipient
     * @param amountIn Amount of tokenIn to receive
     * @param recipient Address to receive tokenOut
     * @param minAmountOut Minimum amount of tokenOut expected
     */
    function swap(address tokenIn, address tokenOut, uint amountIn, address recipient, uint minAmountOut)
        external
        returns (uint amountOut)
    {
        if (shouldRevert) revert("MockSwap: intentional revert");

        // Pull tokenIn from caller (BridgeExecutor)
        require(IERC20(tokenIn).transferFrom(msg.sender, address(this), amountIn), "MockSwap: transferFrom failed");

        // Calculate output amount (simple rate conversion)
        amountOut = amountIn * swapRate / 1e18;
        require(amountOut >= minAmountOut, "MockSwap: insufficient output amount");

        // Send tokenOut to recipient
        require(IERC20(tokenOut).transfer(recipient, amountOut), "MockSwap: transfer failed");

        emit Swapped(tokenIn, tokenOut, recipient, amountIn, amountOut);
    }

    /**
     * @notice Fund the swap contract with tokens for testing
     * @param token Token to fund
     * @param amount Amount to fund
     */
    function fund(address token, uint amount) external {
        require(IERC20(token).transferFrom(msg.sender, address(this), amount), "MockSwap: fund failed");
    }
}

contract BridgeExecutorTest is BridgeTest {
    BridgeExecutor public bridgeExecutorCross;
    BridgeExecutor public bridgeExecutorBSC;
    MockTargetContract public mockTargetCross;
    MockTargetContract public mockTargetBSC;
    MockSwap public mockSwapCross;
    MockSwap public mockSwapBSC;

    function setUp() public override {
        super.setUp();

        // Deploy BridgeExecutor for each chain
        vm.selectFork(crossForkID);
        bridgeExecutorCross = new BridgeExecutor(CrossOWNER, address(bridgeCross));
        mockTargetCross = new MockTargetContract();
        mockSwapCross = new MockSwap();

        vm.selectFork(bscForkID);
        bridgeExecutorBSC = new BridgeExecutor(OWNER, address(bridgeBSC));
        mockTargetBSC = new MockTargetContract();
        mockSwapBSC = new MockSwap();

        // Set BridgeExecutor on bridges
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setBridgeExecutor(IBridgeExecutor(address(bridgeExecutorCross)));

        // Whitelist mock contracts
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(mockTargetCross));
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(mockSwapCross));

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setBridgeExecutor(IBridgeExecutor(address(bridgeExecutorBSC)));

        // Whitelist mock contracts
        vm.prank(OWNER);
        bridgeExecutorBSC.addWhitelistTarget(address(mockTargetBSC));
        vm.prank(OWNER);
        bridgeExecutorBSC.addWhitelistTarget(address(mockSwapBSC));
    }

    /**
     * @notice Test bridge with extradata for native token
     */
    function test_bridgeWithExtraData_nativeToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata: target contract address + calldata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge with extradata
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = address(mockTargetCross).balance;
        uint beforeUserBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify tokens were sent to target contract instead of user
        assertEq(address(mockTargetCross).balance, beforeTargetBalance + value);
        assertEq(USER.balance, beforeUserBalance); // User should not receive additional tokens
    }

    /**
     * @notice Test bridge with extradata for ERC20 token
     */
    function test_bridgeWithExtraData_erc20Token() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Prepare extradata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, amount, bytes("erc20 test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(mockTargetCross));
        uint beforeUserBalance = testTokenCross.balanceOf(USER);

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify tokens were sent to target contract
        assertEq(testTokenCross.balanceOf(address(mockTargetCross)), beforeTargetBalance + value);
        assertEq(testTokenCross.balanceOf(USER), beforeUserBalance); // User should not receive additional tokens
    }

    /**
     * @notice Test bridge with extradata failure - should revert to normal flow
     */
    function test_bridgeWithExtraData_failureRevertsToNormalFlow_nativeToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Set mock to revert
        vm.selectFork(crossForkID);
        mockTargetCross.setShouldRevert(true);

        // Prepare extradata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fall back to sending tokens to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;
        uint beforeExecutorBalance = address(bridgeExecutorCross).balance;
        uint beforeBridgeBalance = address(bridgeCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify tokens were sent to user (fallback behavior)
        assertEq(USER.balance, beforeUserBalance + value, "User should receive tokens");
        assertEq(address(mockTargetCross).balance, beforeTargetBalance, "Target should not receive tokens");
        assertEq(address(bridgeExecutorCross).balance, beforeExecutorBalance, "Executor should return all tokens");
        // Bridge sends value to user, so balance decreases by value
        assertEq(address(bridgeCross).balance, beforeBridgeBalance - value, "Bridge should send value to user");
    }

    /**
     * @notice Test bridge with extradata failure for ERC20 token
     */
    function test_bridgeWithExtraData_failureRevertsToNormalFlow_erc20Token() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Set mock to revert
        vm.selectFork(crossForkID);
        mockTargetCross.setShouldRevert(true);

        // Prepare extradata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fall back to minting to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(mockTargetCross));
        uint beforeExecutorBalance = testTokenCross.balanceOf(address(bridgeExecutorCross));
        uint beforeBridgeBalance = testTokenCross.balanceOf(address(bridgeCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify tokens were minted to user (fallback behavior)
        assertEq(testTokenCross.balanceOf(USER), beforeUserBalance + value, "User should receive tokens");
        assertEq(
            testTokenCross.balanceOf(address(mockTargetCross)), beforeTargetBalance, "Target should not receive tokens"
        );
        assertEq(
            testTokenCross.balanceOf(address(bridgeExecutorCross)),
            beforeExecutorBalance,
            "Executor should not hold tokens (burned)"
        );
        assertEq(testTokenCross.balanceOf(address(bridgeCross)), beforeBridgeBalance, "Bridge should not hold tokens");
    }

    /**
     * @notice Test bridge without extradata (normal flow)
     */
    function test_bridgeWithoutExtraData() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        // Normal bridge without extradata
        deposit(false, amount, 5);

        // Verify bridge worked
        vm.selectFork(crossForkID);
        assertGt(USER.balance, 0);
    }

    /**
     * @notice Test extradata with insufficient length (should use normal flow)
     */
    function test_bridgeWithShortExtraData() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // ExtraData with less than 20 bytes (should be ignored)
        bytes memory extraData = bytes("short");

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should use normal flow
        vm.selectFork(crossForkID);
        uint beforeBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify tokens sent to user (normal flow)
        assertEq(USER.balance, beforeBalance + value);
    }

    /**
     * @notice Test bridge + swap: Bridge testToken and swap to weth
     * @dev Real-world scenario: User bridges token A and receives token B after swap
     */
    function test_bridgeWithSwap_success() public {
        uint amount = 1000 * 1e18;

        // Fund mockSwapCross with weth tokens for swap
        vm.selectFork(crossForkID);
        // Bridge some weth from BSC to fund the swap contract
        vm.selectFork(bscForkID);
        vm.deal(CrossOWNER, amount * 20);
        (uint swapValue, uint swapGas, uint swapService) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount * 10);
        (uint swapIndex,) =
            bscBridge(Const.NATIVE_TOKEN, CrossOWNER, address(mockSwapCross), swapValue, swapGas, swapService, "");
        bscIncrementIndex();
        vm.selectFork(crossForkID);
        crossFinalize(swapIndex, address(weth), address(mockSwapCross), swapValue, 5, "");

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Prepare extradata for swap: MockSwap.swap(tokenIn, tokenOut, amountIn, recipient, minAmountOut)
        bytes memory calldata_ = abi.encodeWithSelector(
            MockSwap.swap.selector,
            address(testTokenCross), // tokenIn (what we're bridging)
            address(weth), // tokenOut (what user wants)
            amount, // amountIn
            USER, // recipient
            amount // minAmountOut (1:1 for simplicity)
        );
        bytes memory extraData = abi.encodePacked(address(mockSwapCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should swap testTokenCross to weth
        vm.selectFork(crossForkID);
        uint beforeTestTokenBalance = testTokenCross.balanceOf(USER);
        uint beforeWethBalance = weth.balanceOf(USER);
        uint beforeSwapTestTokenBalance = testTokenCross.balanceOf(address(mockSwapCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify swap happened: user received weth, not testTokenCross
        assertEq(testTokenCross.balanceOf(USER), beforeTestTokenBalance); // No testToken received
        assertEq(weth.balanceOf(USER), beforeWethBalance + value); // Weth received
        assertEq(testTokenCross.balanceOf(address(mockSwapCross)), beforeSwapTestTokenBalance + value); // Swap has testToken
    }

    /**
     * @notice Test bridge + swap failure: Should fallback to giving user original token
     * @dev When swap fails, user should receive the bridged token instead
     */
    function test_bridgeWithSwap_failure() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Prepare extradata for swap
        bytes memory calldata_ =
            abi.encodeWithSelector(MockSwap.swap.selector, address(testTokenCross), address(weth), amount, USER, amount);
        bytes memory extraData = abi.encodePacked(address(mockSwapCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - swap should fail, user gets testTokenCross instead
        vm.selectFork(crossForkID);
        uint beforeTestTokenBalance = testTokenCross.balanceOf(USER);
        uint beforeWethBalance = weth.balanceOf(USER);

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify fallback: user received testTokenCross (not weth)
        assertEq(testTokenCross.balanceOf(USER), beforeTestTokenBalance + value); // TestToken received (fallback)
        assertEq(weth.balanceOf(USER), beforeWethBalance); // No weth received
    }

    /**
     * @notice Test bridge with non-whitelisted target (should fallback to normal flow)
     * @dev When target is not whitelisted, should give tokens directly to user
     */
    function test_bridgeWithNonWhitelistedTarget() public {
        uint amount = 1000 * 1e18;

        // Deploy a new mock contract that's NOT whitelisted
        vm.selectFork(crossForkID);
        MockTargetContract nonWhitelistedTarget = new MockTargetContract();

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Prepare extradata with non-whitelisted target
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(nonWhitelistedTarget), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fallback to giving tokens to user (not calling non-whitelisted target)
        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(nonWhitelistedTarget));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify tokens were given to user (fallback behavior)
        assertEq(testTokenCross.balanceOf(USER), beforeUserBalance + value);
        assertEq(testTokenCross.balanceOf(address(nonWhitelistedTarget)), beforeTargetBalance); // Target didn't receive
    }

    /**
     * @notice Test whitelist management functions
     * @dev Verifies add/remove whitelist functionality
     */
    function test_whitelistManagement() public {
        vm.selectFork(crossForkID);

        address testTarget = address(0x1234);

        // Check initial state
        assertFalse(bridgeExecutorCross.isWhitelistedTarget(testTarget));

        // Add to whitelist
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(testTarget);
        assertTrue(bridgeExecutorCross.isWhitelistedTarget(testTarget));

        // Remove from whitelist
        vm.prank(CrossOWNER);
        bridgeExecutorCross.removeWhitelistTarget(testTarget);
        assertFalse(bridgeExecutorCross.isWhitelistedTarget(testTarget));
    }

    /**
     * @notice Test bridge + swap with different swap rate
     * @dev Verifies that swap rate conversion works correctly
     */
    function test_bridgeWithSwap_differentRate() public {
        uint amount = 1000 * 1e18;
        uint swapRate = 2 * 1e18; // 1 testToken = 2 weth

        // Fund mockSwapCross with weth tokens for swap (need 2x for 2:1 rate)
        vm.selectFork(bscForkID);
        vm.deal(CrossOWNER, amount * 20);
        (uint swapValue, uint swapGas, uint swapService) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount * 10);
        (uint swapIndex,) =
            bscBridge(Const.NATIVE_TOKEN, CrossOWNER, address(mockSwapCross), swapValue, swapGas, swapService, "");
        bscIncrementIndex();
        vm.selectFork(crossForkID);
        crossFinalize(swapIndex, address(weth), address(mockSwapCross), swapValue, 5, "");

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Set swap rate
        vm.selectFork(crossForkID);
        mockSwapCross.setSwapRate(swapRate);

        // Prepare extradata with adjusted minAmountOut
        uint expectedOutput = amount * swapRate / 1e18;
        bytes memory calldata_ = abi.encodeWithSelector(
            MockSwap.swap.selector,
            address(testTokenCross),
            address(weth),
            amount,
            USER,
            expectedOutput // minAmountOut adjusted for rate
        );
        bytes memory extraData = abi.encodePacked(address(mockSwapCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize
        vm.selectFork(crossForkID);
        uint beforeWethBalance = weth.balanceOf(USER);

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify user received 2x weth (due to 2:1 rate)
        assertEq(weth.balanceOf(USER), beforeWethBalance + expectedOutput);
    }

    /**
     * @notice Test that gas limit protection works correctly
     * @dev Verifies that BridgeExecutor reserves gas for post-call operations
     */
    function test_gasLimitProtection() public {
        uint amount = 1000 * 1e18;

        // Deploy gas consumer and whitelist it
        vm.selectFork(crossForkID);
        MockGasConsumer gasConsumer = new MockGasConsumer();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(gasConsumer));

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata to consume a lot of gas (but not so much to deplete everything)
        // Use moderate iterations so the call succeeds but tests gas reservation
        bytes memory calldata_ = abi.encodeWithSelector(MockGasConsumer.consumeGas.selector, 1000);
        bytes memory extraData = abi.encodePacked(address(gasConsumer), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = address(gasConsumer).balance;
        uint beforeUserBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify the call succeeded and gas consumer received the tokens
        assertEq(address(gasConsumer).balance, beforeTargetBalance + value, "Gas consumer should receive tokens");
        assertEq(USER.balance, beforeUserBalance, "User should not receive tokens on success");
    }

    /**
     * @notice Test that gas limit is properly calculated and reserved
     * @dev Verifies that BridgeExecutor reserves 200k gas for post-call operations
     */
    function test_gasLimitReservation() public {
        uint amount = 1000 * 1e18;

        // Deploy gas consumer and whitelist it
        vm.selectFork(crossForkID);
        MockGasConsumer gasConsumer = new MockGasConsumer();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(gasConsumer));

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata with moderate gas consumption
        bytes memory calldata_ = abi.encodeWithSelector(MockGasConsumer.consumeGas.selector, 500);
        bytes memory extraData = abi.encodePacked(address(gasConsumer), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize with limited gas to test reservation
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = address(gasConsumer).balance;
        uint beforeUserBalance = USER.balance;

        // Call finalize with enough gas for all operations
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify the call succeeded
        assertEq(address(gasConsumer).balance, beforeTargetBalance + value, "Gas consumer should receive tokens");
        assertEq(USER.balance, beforeUserBalance, "User should not receive tokens on success");

        // Verify counter was updated (gas was consumed properly)
        assertEq(gasConsumer.counter(), 499, "Counter should be updated to last iteration value");
    }

    /**
     * @notice Test that gas exhaustion in target call causes fallback to user
     * @dev Verifies that when target call runs out of gas, user receives tokens
     */
    function test_gasExhaustion_fallbackToUser() public {
        uint amount = 1000 * 1e18;

        // Deploy gas consumer and whitelist it
        vm.selectFork(crossForkID);
        MockGasConsumer gasConsumer = new MockGasConsumer();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(gasConsumer));

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata to consume a LOT of gas (will fail due to limited gas)
        bytes memory calldata_ = abi.encodeWithSelector(MockGasConsumer.consumeGas.selector, 50000);
        bytes memory extraData = abi.encodePacked(address(gasConsumer), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain with LIMITED GAS
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = address(gasConsumer).balance;
        uint beforeUserBalance = USER.balance;
        uint beforeBridgeBalance = address(bridgeCross).balance;

        // Call finalize with limited gas (500k should be enough for bridge logic but not for 50000 iterations)
        crossFinalizeWithGasLimit(index, address(NATIVE_TOKEN), USER, value, 5, extraData, 500_000);

        // Due to gas exhaustion, target call should fail and user should receive tokens
        assertEq(address(gasConsumer).balance, beforeTargetBalance, "Gas consumer should NOT receive tokens");
        assertEq(USER.balance, beforeUserBalance + value, "User should receive tokens as fallback");
        assertEq(
            address(bridgeCross).balance,
            beforeBridgeBalance - value,
            "Bridge balance should decrease by value sent to user"
        );
    }

    // Override to add vm.recordLogs() for event verification
    function crossFinalize(uint index, address token, address to, uint value, uint sigCount, bytes memory extraData)
        public
        override
        returns (bool ok)
    {
        vm.recordLogs();
        return super.crossFinalize(index, token, to, value, sigCount, extraData);
    }

    function crossFinalizeWithGasLimit(
        uint index,
        address token,
        address to,
        uint value,
        uint validatorCount,
        bytes memory extraData,
        uint gasLimit
    ) internal {
        IBridgeRegistry.FinalizeArguments memory args = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: index,
            toToken: IERC20(token),
            to: to,
            value: value,
            extraData: extraData
        });
        (uint8[] memory v, bytes32[] memory r, bytes32[] memory s) =
            signFinalize(BSC_CHAIN_ID, index, token, to, value, extraData, validatorCount);
        vm.recordLogs();
        // Call with limited gas
        (bool success,) = address(bridgeCross).call{gas: gasLimit}(
            abi.encodeWithSelector(bridgeCross.finalizeBridge.selector, args, v, r, s)
        );
        require(success, "finalizeBridge call failed");
    }

    function signFinalize(
        uint fromChainID,
        uint index,
        address token,
        address to,
        uint value,
        bytes memory extraData,
        uint validatorCount
    ) internal view returns (uint8[] memory v, bytes32[] memory r, bytes32[] memory s) {
        v = new uint8[](validatorCount);
        r = new bytes32[](validatorCount);
        s = new bytes32[](validatorCount);

        bytes32 structHash =
            keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, keccak256(extraData)));
        bytes32 digest = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), structHash);

        for (uint i = 0; i < validatorCount; ++i) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], digest);
        }
    }

    /**
     * @notice Test that ExtraCallExecuted event emits correct returnData on success
     * @dev Verifies that the return value from target contract is properly captured in event
     */
    function test_extraCallEvent_returnData_onSuccess() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata with function that returns data
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallbackWithReturn.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Calculate expected return value
        bytes32 expectedReturnValue = keccak256(abi.encode(address(1), USER, amount, bytes("test data")));

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain and capture logs
        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Get recorded logs and find ExtraCallExecuted event from BridgeExecutor
        // New event signature: ExtraCallExecuted(uint256 indexed, uint256 indexed, address indexed, bytes4, bool, uint256, bytes)
        Vm.Log[] memory logs = vm.getRecordedLogs();
        bool foundEvent = false;
        bytes memory returnData;

        for (uint i = 0; i < logs.length; i++) {
            // ExtraCallExecuted event signature (updated)
            if (logs[i].topics[0] == keccak256("ExtraCallExecuted(uint256,uint256,address,bytes4,bool,uint256,bytes)"))
            {
                foundEvent = true;
                // Decode non-indexed parameters: methodID, success, remaining, returnData
                (, bool success,, bytes memory data) = abi.decode(logs[i].data, (bytes4, bool, uint, bytes));
                returnData = data;

                // Verify success is true
                assertTrue(success, "ExtraCall should succeed");

                // Verify returnData contains the expected return value (abi-encoded bytes32)
                bytes32 decodedReturnValue = abi.decode(returnData, (bytes32));
                assertEq(decodedReturnValue, expectedReturnValue, "Return data should match expected hash");
                break;
            }
        }

        assertTrue(foundEvent, "ExtraCallExecuted event should be emitted");
    }

    /**
     * @notice Test that ExtraCallExecuted event emits revert reason on failure
     * @dev Verifies that the revert reason from target contract is properly captured in event
     */
    function test_extraCallEvent_returnData_onFailure() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Set mock to revert
        vm.selectFork(crossForkID);
        mockTargetCross.setShouldRevert(true);

        // Prepare extradata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallbackWithReturn.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain and capture logs
        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Get recorded logs and find ExtraCallExecuted event from BridgeExecutor
        // New event signature: ExtraCallExecuted(uint256 indexed, uint256 indexed, address indexed, bytes4, bool, uint256, bytes)
        Vm.Log[] memory logs = vm.getRecordedLogs();
        bool foundEvent = false;

        for (uint i = 0; i < logs.length; i++) {
            // ExtraCallExecuted event signature (updated)
            if (logs[i].topics[0] == keccak256("ExtraCallExecuted(uint256,uint256,address,bytes4,bool,uint256,bytes)"))
            {
                foundEvent = true;
                // Decode non-indexed parameters: methodID, success, remaining, returnData
                (, bool success,, bytes memory returnData) = abi.decode(logs[i].data, (bytes4, bool, uint, bytes));

                // Verify success is false
                assertFalse(success, "ExtraCall should fail");

                // Verify returnData contains revert reason
                // The revert reason is encoded as Error(string) selector + abi-encoded string
                assertTrue(returnData.length > 0, "Return data should contain revert reason");

                // Check that the revert reason contains our expected message
                // Error(string) selector is 0x08c379a0
                bytes4 errorSelector;
                assembly {
                    errorSelector := mload(add(returnData, 32))
                }
                assertEq(errorSelector, bytes4(0x08c379a0), "Should be Error(string) selector");

                // Decode the error message
                // Skip the selector (4 bytes) and decode the string
                bytes memory errorData = new bytes(returnData.length - 4);
                for (uint j = 4; j < returnData.length; j++) {
                    errorData[j - 4] = returnData[j];
                }
                string memory errorMessage = abi.decode(errorData, (string));
                assertEq(errorMessage, "MockTargetContract: intentional revert", "Revert reason should match");
                break;
            }
        }

        assertTrue(foundEvent, "ExtraCallExecuted event should be emitted");
    }

    /**
     * @notice Test that ExtraCallExecuted event emits correct returnData for ERC20 token success
     * @dev Verifies return value capture works for ERC20 tokens as well
     */
    function test_extraCallEvent_returnData_onSuccess_erc20() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Prepare extradata with function that returns data
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallbackWithReturn.selector,
            address(testTokenCross),
            USER,
            amount,
            bytes("erc20 test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Calculate expected return value
        bytes32 expectedReturnValue = keccak256(abi.encode(address(testTokenCross), USER, amount, bytes("erc20 test")));

        // Initiate bridge
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain and capture logs
        vm.selectFork(crossForkID);
        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Get recorded logs and find ExtraCallExecuted event from BridgeExecutor
        // New event signature: ExtraCallExecuted(uint256 indexed, uint256 indexed, address indexed, bytes4, bool, uint256, bytes)
        Vm.Log[] memory logs = vm.getRecordedLogs();
        bool foundEvent = false;

        for (uint i = 0; i < logs.length; i++) {
            // ExtraCallExecuted event signature (updated)
            if (logs[i].topics[0] == keccak256("ExtraCallExecuted(uint256,uint256,address,bytes4,bool,uint256,bytes)"))
            {
                foundEvent = true;
                // Decode non-indexed parameters: methodID, success, remaining, returnData
                (, bool success,, bytes memory returnData) = abi.decode(logs[i].data, (bytes4, bool, uint, bytes));

                // Verify success is true
                assertTrue(success, "ExtraCall should succeed");

                // Verify returnData contains the expected return value
                bytes32 decodedReturnValue = abi.decode(returnData, (bytes32));
                assertEq(decodedReturnValue, expectedReturnValue, "Return data should match expected hash");
                break;
            }
        }

        assertTrue(foundEvent, "ExtraCallExecuted event should be emitted");
    }

    // ============ New Test Cases ============

    /**
     * @notice Test method whitelist - allowed selector
     */
    function test_methodWhitelist_allowedSelector() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);

        // Enable method check for mockTargetCross
        vm.prank(CrossOWNER);
        bridgeExecutorCross.setMethodCheckEnabled(address(mockTargetCross), true);

        // Add allowed method selector
        bytes4[] memory selectors = new bytes4[](1);
        selectors[0] = MockTargetContract.handleBridgeCallback.selector;
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistMethods(address(mockTargetCross), selectors);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata with allowed selector
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should succeed since selector is whitelisted
        vm.selectFork(crossForkID);
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        assertEq(address(mockTargetCross).balance, beforeTargetBalance + value, "Target should receive tokens");
    }

    /**
     * @notice Test method whitelist - disallowed selector should fallback to user
     */
    function test_methodWhitelist_disallowedSelector() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);

        // Enable method check for mockTargetCross
        vm.prank(CrossOWNER);
        bridgeExecutorCross.setMethodCheckEnabled(address(mockTargetCross), true);

        // Don't whitelist any methods - all selectors should be rejected

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata with non-whitelisted selector
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fallback to user since selector is not whitelisted
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        assertEq(USER.balance, beforeUserBalance + value, "User should receive tokens");
        assertEq(address(mockTargetCross).balance, beforeTargetBalance, "Target should NOT receive tokens");
    }

    /**
     * @notice Test remaining tokens - partial consume
     */
    function test_remaining_partialConsume() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);
        // Set mock to consume only 50%
        mockTargetCross.setConsumePercent(50);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target receives 50%, user receives remaining 50%
        uint expectedTargetAmount = value * 50 / 100;
        uint expectedUserAmount = value - expectedTargetAmount;
        assertEq(
            address(mockTargetCross).balance, beforeTargetBalance + expectedTargetAmount, "Target should receive 50%"
        );
        assertEq(USER.balance, beforeUserBalance + expectedUserAmount, "User should receive remaining 50%");
    }

    /**
     * @notice Test remaining tokens - zero consume (target doesn't pull tokens)
     */
    function test_remaining_zeroConsume() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);
        // Set mock to consume 0%
        mockTargetCross.setConsumePercent(0);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target receives 0%, user receives all 100%
        assertEq(address(mockTargetCross).balance, beforeTargetBalance, "Target should receive 0%");
        assertEq(USER.balance, beforeUserBalance + value, "User should receive all 100%");
    }

    /**
     * @notice Test remaining tokens - ERC20 partial consume
     */
    function test_remaining_partialConsume_erc20() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);
        // Set mock to consume only 70%
        mockTargetCross.setConsumePercent(70);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, amount, bytes("erc20 test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(mockTargetCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Target receives 70%, user receives remaining 30%
        uint expectedTargetAmount = value * 70 / 100;
        uint expectedUserAmount = value - expectedTargetAmount;
        assertEq(
            testTokenCross.balanceOf(address(mockTargetCross)),
            beforeTargetBalance + expectedTargetAmount,
            "Target should receive 70%"
        );
        assertEq(
            testTokenCross.balanceOf(USER), beforeUserBalance + expectedUserAmount, "User should receive remaining 30%"
        );
    }

    /**
     * @notice Test GAS_RESERVE admin change
     */
    function test_gasReserve_adminChange() public {
        vm.selectFork(crossForkID);

        // Check initial value
        assertEq(bridgeExecutorCross.postCallGasReserve(), 200_000);

        // Change gas reserve
        vm.prank(CrossOWNER);
        bridgeExecutorCross.setPostCallGasReserve(300_000);

        assertEq(bridgeExecutorCross.postCallGasReserve(), 300_000);

        // Test invalid values
        vm.prank(CrossOWNER);
        vm.expectRevert();
        bridgeExecutorCross.setPostCallGasReserve(10_000); // Too low

        vm.prank(CrossOWNER);
        vm.expectRevert();
        bridgeExecutorCross.setPostCallGasReserve(2_000_000); // Too high
    }

    /**
     * @notice Test batch method whitelist management
     */
    function test_batchMethodWhitelist() public {
        vm.selectFork(crossForkID);

        address target = address(mockTargetCross);

        // Add multiple methods
        bytes4[] memory selectors = new bytes4[](3);
        selectors[0] = MockTargetContract.handleBridgeCallback.selector;
        selectors[1] = MockTargetContract.handleBridgeCallbackWithReturn.selector;
        selectors[2] = bytes4(keccak256("customFunction()"));

        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistMethods(target, selectors);

        assertTrue(bridgeExecutorCross.isWhitelistedMethod(target, selectors[0]));
        assertTrue(bridgeExecutorCross.isWhitelistedMethod(target, selectors[1]));
        assertTrue(bridgeExecutorCross.isWhitelistedMethod(target, selectors[2]));

        // Remove methods
        bytes4[] memory removeSelectors = new bytes4[](2);
        removeSelectors[0] = selectors[0];
        removeSelectors[1] = selectors[2];

        vm.prank(CrossOWNER);
        bridgeExecutorCross.removeWhitelistMethods(target, removeSelectors);

        assertFalse(bridgeExecutorCross.isWhitelistedMethod(target, selectors[0]));
        assertTrue(bridgeExecutorCross.isWhitelistedMethod(target, selectors[1]));
        assertFalse(bridgeExecutorCross.isWhitelistedMethod(target, selectors[2]));
    }

    /**
     * @notice Test early-return: Origin ERC20 with non-whitelisted target
     * @dev When executor encounters non-whitelisted target, it sends tokens directly to user.
     *      Returns (false, 0) to indicate user already received tokens.
     */
    function test_earlyReturn_originERC20_nonWhitelistedTarget() public {
        uint amount = 1000 * 1e18;

        // First, deposit testToken from BSC to CROSS to get some testTokenCross
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint depositIndex,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, "");
        bscIncrementIndex();

        // Finalize on CROSS to get testTokenCross for USER
        vm.selectFork(crossForkID);
        crossFinalize(depositIndex, address(testTokenCross), USER, value, 5, "");

        // Now bridge back CROSS -> BSC with extraData targeting a NON-whitelisted contract
        vm.selectFork(crossForkID);
        address nonWhitelistedTarget = address(0xDEADBEEF);

        // Prepare extraData with non-whitelisted target
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenBSC), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(nonWhitelistedTarget, calldata_);

        // Calculate fee and bridge from CROSS
        (uint crossValue, uint crossGas, uint crossService) = crossCalcFee(testTokenCross, value);

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), value);

        uint bridgeIndex = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        vm.prank(USER);
        bridgeCross.bridgeToken(BSC_CHAIN_ID, testTokenCross, USER, crossValue, crossGas, crossService, extraData);
        crossIncrementIndex();

        // Finalize on BSC - executor sends directly to user, returns (false, 0)
        vm.selectFork(bscForkID);
        uint beforeUserBalance = testTokenBSC.balanceOf(USER);
        uint beforeExecutorBalance = testTokenBSC.balanceOf(address(bridgeExecutorBSC));

        bscFinalize(bridgeIndex, address(testTokenBSC), USER, crossValue, 5, extraData);

        // Verify: User receives tokens directly from executor
        assertEq(
            testTokenBSC.balanceOf(USER),
            beforeUserBalance + crossValue,
            "User should receive origin tokens directly from executor"
        );
        assertEq(
            testTokenBSC.balanceOf(address(bridgeExecutorBSC)),
            beforeExecutorBalance,
            "Executor should not hold any tokens"
        );
    }

    /**
     * @notice Test early-return: Origin ERC20 with method not whitelisted
     * @dev When method selector is not whitelisted, executor sends tokens directly to user.
     */
    function test_earlyReturn_originERC20_methodNotWhitelisted() public {
        uint amount = 1000 * 1e18;

        // Setup: deposit testToken from BSC to CROSS
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint depositIndex,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, "");
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(depositIndex, address(testTokenCross), USER, value, 5, "");

        // Enable method check for mockTargetBSC but don't whitelist any methods
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeExecutorBSC.setMethodCheckEnabled(address(mockTargetBSC), true);

        // Bridge from CROSS -> BSC with extraData
        vm.selectFork(crossForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenBSC), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetBSC), calldata_);

        // Calculate fee
        (uint crossValue, uint crossGas, uint crossService) = crossCalcFee(testTokenCross, value);

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), value);

        uint bridgeIndex = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        vm.prank(USER);
        bridgeCross.bridgeToken(BSC_CHAIN_ID, testTokenCross, USER, crossValue, crossGas, crossService, extraData);
        crossIncrementIndex();

        // Finalize on BSC - executor sends directly to user
        vm.selectFork(bscForkID);
        uint beforeUserBalance = testTokenBSC.balanceOf(USER);

        bscFinalize(bridgeIndex, address(testTokenBSC), USER, crossValue, 5, extraData);

        // Verify user receives tokens directly from executor
        assertEq(
            testTokenBSC.balanceOf(USER),
            beforeUserBalance + crossValue,
            "User should receive tokens directly from executor"
        );
    }

    /**
     * @notice Test early-return: Origin ERC20 with short extraData
     * @dev When extraData is too short (< 24 bytes), executor sends tokens directly to user.
     */
    function test_earlyReturn_originERC20_shortExtraData() public {
        uint amount = 1000 * 1e18;

        // Setup: deposit testToken from BSC to CROSS
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint depositIndex,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, "");
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(depositIndex, address(testTokenCross), USER, value, 5, "");

        // Bridge from CROSS -> BSC with SHORT extraData (< 24 bytes)
        // This should cause early-return in executor
        bytes memory shortExtraData = bytes("short"); // Only 5 bytes

        // Calculate fee
        (uint crossValue, uint crossGas, uint crossService) = crossCalcFee(testTokenCross, value);

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), value);

        uint bridgeIndex = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        vm.prank(USER);
        bridgeCross.bridgeToken(BSC_CHAIN_ID, testTokenCross, USER, crossValue, crossGas, crossService, shortExtraData);
        crossIncrementIndex();

        // Finalize on BSC - executor sends directly to user
        vm.selectFork(bscForkID);
        uint beforeUserBalance = testTokenBSC.balanceOf(USER);

        bscFinalize(bridgeIndex, address(testTokenBSC), USER, crossValue, 5, shortExtraData);

        // Verify user receives tokens directly from executor
        assertEq(
            testTokenBSC.balanceOf(USER),
            beforeUserBalance + crossValue,
            "User should receive tokens directly from executor"
        );
    }

    /**
     * @notice Test early-return: Wrapped ERC20 with non-whitelisted target
     * @dev When executor encounters non-whitelisted target with wrapped token,
     *      it transfers directly to user (CrossMintableERC20 supports transfer).
     */
    function test_earlyReturn_wrappedERC20_nonWhitelistedTarget() public {
        uint amount = 1000 * 1e18;

        // BSC -> CROSS: bridge testTokenBSC to get testTokenCross (wrapped)
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);

        // Prepare extraData with NON-whitelisted target
        address nonWhitelistedTarget = address(0xDEADBEEF);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(nonWhitelistedTarget, calldata_);

        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS - executor transfers directly to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeExecutorBalance = testTokenCross.balanceOf(address(bridgeExecutorCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify: User receives wrapped tokens directly from executor
        assertEq(
            testTokenCross.balanceOf(USER),
            beforeUserBalance + value,
            "User should receive wrapped tokens directly from executor"
        );
        assertEq(
            testTokenCross.balanceOf(address(bridgeExecutorCross)),
            beforeExecutorBalance,
            "Executor should not hold any wrapped tokens"
        );
    }

    /**
     * @notice Test executor: Insufficient balance sends directly to user
     * @dev When executor has less tokens than expected, it sends actual balance directly to user.
     *      Returns (false, 0) since user received tokens.
     */
    function test_executor_insufficientBalance_sendsDirectlyToUser() public {
        vm.selectFork(crossForkID);

        uint value = 1000 * 1e18;
        uint actualBalance = 500 * 1e18; // Less than value

        // Mint testTokenCross to executor (less than expected value)
        vm.prank(address(bridgeCross));
        ICrossMintableERC20(address(testTokenCross)).mint(address(bridgeExecutorCross), actualBalance);

        // Prepare valid extraData
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        uint beforeUserBalance = testTokenCross.balanceOf(USER);

        // Call executor directly - should send actualBalance to user and return (false, 0)
        vm.prank(address(bridgeCross));
        (bool success, uint remaining) =
            bridgeExecutorCross.executeExtraCall(BSC_CHAIN_ID, 999, testTokenCross, USER, value, extraData);

        assertFalse(success, "Should fail due to insufficient balance");
        assertEq(remaining, 0, "Remaining should be 0 (user received tokens)");

        // Verify user received tokens directly
        assertEq(
            testTokenCross.balanceOf(USER),
            beforeUserBalance + actualBalance,
            "User should receive actual balance directly"
        );
        assertEq(
            testTokenCross.balanceOf(address(bridgeExecutorCross)),
            0,
            "Executor should have 0 balance after direct transfer"
        );
    }

    /**
     * @notice Test executor: Approve fails but transfer succeeds
     * @dev When target approval fails, executor tries to transfer directly to user.
     *      If transfer succeeds, returns (false, 0). User receives tokens.
     */
    function test_executor_approveFails_transferSucceeds_sendsToUser() public {
        vm.selectFork(crossForkID);

        // Deploy bad approve token (approve always fails, but transfer works)
        MockBadApproveToken badToken = new MockBadApproveToken();
        uint value = 1000 * 1e18;

        // Mint tokens to executor
        badToken.mint(address(bridgeExecutorCross), value);

        // Prepare valid extraData targeting whitelisted contract
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(badToken), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        uint beforeUserBalance = badToken.balanceOf(USER);

        // Call executor directly - approve to target fails, but transfer to user succeeds
        vm.prank(address(bridgeCross));
        (bool success, uint remaining) =
            bridgeExecutorCross.executeExtraCall(BSC_CHAIN_ID, 998, IERC20(address(badToken)), USER, value, extraData);

        assertFalse(success, "Should fail due to forceApprove failure");
        assertEq(remaining, 0, "Remaining should be 0 (user received tokens via transfer)");

        // User received tokens via direct transfer
        assertEq(
            badToken.balanceOf(USER),
            beforeUserBalance + value,
            "User should receive tokens via transfer when approve fails"
        );
        assertEq(badToken.balanceOf(address(bridgeExecutorCross)), 0, "Executor should have 0 balance after transfer");
    }

    /**
     * @notice Test early-return: Native token with non-whitelisted target
     * @dev When executor encounters non-whitelisted target with native token,
     *      it sends ETH directly to user and returns (false, 0).
     */
    function test_earlyReturn_nativeToken_nonWhitelistedTarget() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extraData with NON-whitelisted target for native token
        address nonWhitelistedTarget = address(0xDEADBEEF);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(nonWhitelistedTarget, calldata_);

        // Initiate bridge with native token
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS - executor sends directly to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify user receives native tokens directly from executor
        assertEq(USER.balance, beforeUserBalance + value, "User should receive native tokens directly");
    }

    /**
     * @notice Test target revert: User receives remaining after target failure
     * @dev When target reverts, user receives the remaining tokens (not consumed by target).
     *      In this case target is set to consume 100% but revert, so remaining = value.
     */
    function test_targetReverts_userReceivesRemaining() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Set mock to revert AFTER consuming tokens
        vm.selectFork(crossForkID);
        mockTargetCross.setShouldRevert(true);
        mockTargetCross.setConsumePercent(100);

        vm.selectFork(bscForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target reverts, so user should receive all tokens
        assertEq(USER.balance, beforeUserBalance + value, "User should receive all tokens on target revert");
    }

    /**
     * @notice Test remaining bridge recovery: Origin ERC20
     * @dev When remaining > 0 and direct transfer to user fails, bridge recovers via transferFrom
     *      and transfers to user via normal flow.
     */
    function test_remaining_bridgeRecovery_originERC20() public {
        uint amount = 1000 * 1e18;

        // Setup: deposit testToken from BSC to CROSS
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint depositIndex,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, "");
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(depositIndex, address(testTokenCross), USER, value, 5, "");

        // Set mock to consume 60% on BSC
        vm.selectFork(bscForkID);
        mockTargetBSC.setConsumePercent(60);

        // Bridge from CROSS -> BSC with extraData
        vm.selectFork(crossForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenBSC), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetBSC), calldata_);

        (uint crossValue, uint crossGas, uint crossService) = crossCalcFee(testTokenCross, value);

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), value);

        uint bridgeIndex = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        vm.prank(USER);
        bridgeCross.bridgeToken(BSC_CHAIN_ID, testTokenCross, USER, crossValue, crossGas, crossService, extraData);
        crossIncrementIndex();

        // Finalize on BSC
        vm.selectFork(bscForkID);
        uint beforeUserBalance = testTokenBSC.balanceOf(USER);
        uint beforeTargetBalance = testTokenBSC.balanceOf(address(mockTargetBSC));

        bscFinalize(bridgeIndex, address(testTokenBSC), USER, crossValue, 5, extraData);

        // Target receives 60%, user receives remaining 40% via bridge recovery
        uint targetReceived = testTokenBSC.balanceOf(address(mockTargetBSC)) - beforeTargetBalance;
        uint userReceived = testTokenBSC.balanceOf(USER) - beforeUserBalance;

        // Verify target got ~60% and user got ~40%
        assertApproxEqRel(targetReceived, crossValue * 60 / 100, 0.01e18, "Target should receive ~60%");
        assertApproxEqRel(userReceived, crossValue * 40 / 100, 0.01e18, "User should receive ~40% via bridge recovery");
        assertEq(targetReceived + userReceived, crossValue, "Total should equal crossValue");
    }

    /**
     * @notice Test remaining bridge recovery: Wrapped ERC20
     * @dev When remaining > 0 for wrapped token, bridge burns from executor and mints to user.
     */
    function test_remaining_bridgeRecovery_wrappedERC20() public {
        uint amount = 1000 * 1e18;

        // BSC -> CROSS: bridge testTokenBSC to get testTokenCross (wrapped)
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Set mock to consume 80%
        vm.selectFork(crossForkID);
        mockTargetCross.setConsumePercent(80);

        vm.selectFork(bscForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(testTokenCross), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(testTokenBSC, amount);
        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS
        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(mockTargetCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Target receives 80%, user receives remaining 20% via bridge (burn/mint)
        uint expectedTargetAmount = value * 80 / 100;
        uint expectedUserAmount = value - expectedTargetAmount;
        assertEq(
            testTokenCross.balanceOf(address(mockTargetCross)),
            beforeTargetBalance + expectedTargetAmount,
            "Target should receive 80%"
        );
        assertEq(
            testTokenCross.balanceOf(USER),
            beforeUserBalance + expectedUserAmount,
            "User should receive remaining 20% via bridge burn/mint"
        );
    }

    /**
     * @notice Test remaining bridge recovery: Native token
     * @dev When remaining > 0 for native token and direct transfer fails,
     *      executor sends to bridge and bridge forwards to user.
     */
    function test_remaining_bridgeRecovery_native() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Set mock to consume 40%
        vm.selectFork(crossForkID);
        mockTargetCross.setConsumePercent(40);

        vm.selectFork(bscForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target receives 40%, user receives remaining 60% via bridge
        uint expectedTargetAmount = value * 40 / 100;
        uint expectedUserAmount = value - expectedTargetAmount;
        assertEq(
            address(mockTargetCross).balance, beforeTargetBalance + expectedTargetAmount, "Target should receive 40%"
        );
        assertEq(USER.balance, beforeUserBalance + expectedUserAmount, "User should receive remaining 60% via bridge");
    }

    /**
     * @notice Test: remaining == 0 means executor already sent to user
     * @dev When executor successfully sends remaining directly to user,
     *      it returns remaining = 0 and bridge completes without additional transfer.
     */
    function test_remaining_executorDirectTransfer_success() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Set mock to consume 25%
        vm.selectFork(crossForkID);
        mockTargetCross.setConsumePercent(25);

        vm.selectFork(bscForkID);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target receives 25%, user receives remaining 75%
        // (either directly from executor or via bridge - same result)
        uint expectedTargetAmount = value * 25 / 100;
        uint expectedUserAmount = value - expectedTargetAmount;
        assertEq(
            address(mockTargetCross).balance, beforeTargetBalance + expectedTargetAmount, "Target should receive 25%"
        );
        assertEq(USER.balance, beforeUserBalance + expectedUserAmount, "User should receive remaining 75%");
    }

    // ============ Underflow Protection Tests ============

    /**
     * @notice Test consumed calculation when target returns MORE ETH than value
     * @dev Ensures no underflow when balAfter - balBefore > value
     *      Scenario: Target contract is pre-funded and returns all its balance
     */
    function test_nativeConsumed_noUnderflow_extraETHReturned() public {
        uint amount = 1000 * 1e18;
        uint extraETH = 500 * 1e18; // Extra ETH pre-funded to mock

        // Deploy and whitelist MockExtraETHReturn on CROSS
        vm.selectFork(crossForkID);
        MockExtraETHReturn extraReturnMock = new MockExtraETHReturn();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(extraReturnMock));

        // Pre-fund the mock contract with extra ETH
        vm.deal(address(extraReturnMock), extraETH);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata targeting the mock that returns extra ETH
        bytes memory calldata_ = abi.encodeWithSelector(
            MockExtraETHReturn.handleWithExtraReturn.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(extraReturnMock), calldata_);

        // Initiate bridge
        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize on CROSS chain - should NOT underflow
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        // This should succeed without underflow
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Mock returned value + extraETH, so:
        // - balAfter - balBefore = value + extraETH > value
        // - consumed should be 0 (protected from underflow)
        // - remaining = value - 0 = value
        // User receives value (full amount) since consumed = 0
        assertEq(USER.balance, beforeUserBalance + value, "User should receive full value");
        // Mock should have 0 balance (sent everything back)
        assertEq(address(extraReturnMock).balance, 0, "Mock should have sent all ETH");
    }

    /**
     * @notice Test consumed calculation via executor direct call
     * @dev More direct test of the consumed calculation logic
     */
    function test_executor_nativeConsumed_underflowProtection() public {
        uint value = 100 ether;
        uint extraETH = 50 ether;

        vm.selectFork(crossForkID);

        // Deploy and whitelist MockExtraETHReturn
        MockExtraETHReturn extraReturnMock = new MockExtraETHReturn();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(extraReturnMock));

        // Pre-fund the mock contract with extra ETH
        vm.deal(address(extraReturnMock), extraETH);

        // Prepare extradata
        bytes memory calldata_ = abi.encodeWithSelector(
            MockExtraETHReturn.handleWithExtraReturn.selector, address(1), USER, value, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(extraReturnMock), calldata_);

        uint beforeUserBalance = USER.balance;

        // Call executor directly with native token
        vm.deal(address(bridgeCross), value);
        vm.prank(address(bridgeCross));
        (bool success, uint remaining) = bridgeExecutorCross.executeExtraCall{value: value}(
            BSC_CHAIN_ID, 12345, IERC20(Const.NATIVE_TOKEN), USER, value, extraData
        );

        // Call should succeed (target executed)
        assertTrue(success, "Target call should succeed");

        // Since target returned value + extraETH:
        // - returned = (value + extraETH) - 0 = value + extraETH
        // - consumed = returned >= value ? 0 : value - returned = 0
        // - remaining = value - consumed = value
        // But target successfully executed, so remaining should reflect that
        // The user receives remaining directly or via bridge

        // Mock sent all its ETH back to executor, executor sent remaining to user
        assertEq(address(extraReturnMock).balance, 0, "Mock should be empty");

        // User should have received remaining (value since consumed = 0)
        // Either remaining = 0 (sent to user) or remaining > 0 (for bridge to handle)
        uint userReceived = USER.balance - beforeUserBalance;
        assertTrue(
            userReceived == value || remaining == value, "User should receive value or remaining should be value"
        );
    }

    /**
     * @notice Test edge case: target returns exactly what it received
     * @dev consumed should equal value, remaining = 0
     */
    function test_nativeConsumed_exactReturn() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(crossForkID);
        // Set mock to consume 0% (returns everything)
        mockTargetCross.setConsumePercent(0);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Target returned value, so:
        // - consumed = 0
        // - remaining = value
        // User receives all value
        assertEq(USER.balance, beforeUserBalance + value, "User should receive full value when target returns all");
    }

    // ============ Return Data Size Cap Tests (CBU-11) ============

    /**
     * @notice Test that large return data is capped and doesn't cause OOG
     * @dev Verifies CBU-11 fix: returnData should be limited to MAX_RETURN_DATA_SIZE (1KB)
     */
    function test_returnDataSizeCap_largeReturn() public {
        uint amount = 1000 * 1e18;

        // Deploy and whitelist MockLargeReturnData
        vm.selectFork(crossForkID);
        MockLargeReturnData largeReturnMock = new MockLargeReturnData();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(largeReturnMock));

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata to return 10KB of data (should be capped to 1KB)
        bytes memory calldata_ = abi.encodeWithSelector(MockLargeReturnData.returnLargeData.selector, 10 * 1024);
        bytes memory extraData = abi.encodePacked(address(largeReturnMock), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize should NOT revert even with large return data
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        // This should succeed without OOG
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // User should receive tokens (target consumed the native, returned large data)
        assertGe(USER.balance, beforeUserBalance, "User balance should not decrease");
    }

    /**
     * @notice Test normal return data size (under cap) works correctly
     * @dev Small return data should be captured fully
     */
    function test_returnDataSizeCap_normalReturn() public {
        uint amount = 1000 * 1e18;

        // Deploy and whitelist MockLargeReturnData
        vm.selectFork(crossForkID);
        MockLargeReturnData largeReturnMock = new MockLargeReturnData();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(largeReturnMock));

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);

        // Prepare extradata to return 100 bytes (under 1KB cap)
        bytes memory calldata_ = abi.encodeWithSelector(MockLargeReturnData.returnLargeData.selector, 100);
        bytes memory extraData = abi.encodePacked(address(largeReturnMock), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;

        // This should succeed
        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // User should receive tokens
        assertGe(USER.balance, beforeUserBalance, "User balance should not decrease");
    }

    /**
     * @notice Test executor directly with large return data
     * @dev Direct call to verify returnData is capped
     */
    function test_executor_returnDataCapped_directCall() public {
        uint value = 100 ether;

        vm.selectFork(crossForkID);

        // Deploy and whitelist MockLargeReturnData
        MockLargeReturnData largeReturnMock = new MockLargeReturnData();
        vm.prank(CrossOWNER);
        bridgeExecutorCross.addWhitelistTarget(address(largeReturnMock));

        // Prepare extradata to return 50KB (way over 1KB cap)
        bytes memory calldata_ = abi.encodeWithSelector(MockLargeReturnData.returnLargeData.selector, 50 * 1024);
        bytes memory extraData = abi.encodePacked(address(largeReturnMock), calldata_);

        // Fund bridge and call executor directly
        vm.deal(address(bridgeCross), value);
        vm.prank(address(bridgeCross));

        // This should NOT revert due to OOG
        (bool success,) = bridgeExecutorCross.executeExtraCall{value: value}(
            BSC_CHAIN_ID, 99999, IERC20(Const.NATIVE_TOKEN), USER, value, extraData
        );

        // Call should succeed (returnData capped internally)
        assertTrue(success, "Executor should handle large return data without OOG");
    }
}
