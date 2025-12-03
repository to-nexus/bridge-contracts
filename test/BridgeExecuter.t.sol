// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {BridgeExecuter} from "../src/BridgeExecuter.sol";
import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {IBridgeExecuter} from "../src/interface/IBridgeExecuter.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {ICrossMintableERC20} from "../src/token/ICrossMintableERC20.sol";

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";

/**
 * @title MockTargetContract
 * @notice Mock contract for testing BridgeExecuter functionality
 */
contract MockTargetContract {
    event Received(address token, address from, uint value, bytes data);
    event NativeReceived(address from, uint value, bytes data);

    bool public shouldRevert;

    function setShouldRevert(bool _shouldRevert) external {
        shouldRevert = _shouldRevert;
    }

    function handleBridgeCallback(address token, address user, uint amount, bytes calldata data) external payable {
        if (shouldRevert) revert("MockTargetContract: intentional revert");

        if (token == address(1)) {
            // Native token
            require(msg.value == amount, "MockTargetContract: incorrect value");
            emit NativeReceived(user, amount, data);
        } else {
            // ERC20 token - pull tokens from executer
            require(
                IERC20(token).transferFrom(msg.sender, address(this), amount), "MockTargetContract: transfer failed"
            );
            emit Received(token, user, amount, data);
        }
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

        // Pull tokenIn from caller (BridgeExecuter)
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

contract BridgeExecuterTest is BridgeTest {
    BridgeExecuter public bridgeExecuterCross;
    BridgeExecuter public bridgeExecuterBSC;
    MockTargetContract public mockTargetCross;
    MockTargetContract public mockTargetBSC;
    MockSwap public mockSwapCross;
    MockSwap public mockSwapBSC;

    function setUp() public override {
        super.setUp();

        // Deploy BridgeExecuter for each chain
        vm.selectFork(crossForkID);
        bridgeExecuterCross = new BridgeExecuter(CrossOWNER, address(bridgeCross));
        mockTargetCross = new MockTargetContract();
        mockSwapCross = new MockSwap();

        vm.selectFork(bscForkID);
        bridgeExecuterBSC = new BridgeExecuter(OWNER, address(bridgeBSC));
        mockTargetBSC = new MockTargetContract();
        mockSwapBSC = new MockSwap();

        // Set BridgeExecuter on bridges
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setBridgeExecuter(IBridgeExecuter(address(bridgeExecuterCross)));

        // Whitelist mock contracts
        vm.prank(CrossOWNER);
        bridgeExecuterCross.addWhitelistTarget(address(mockTargetCross));
        vm.prank(CrossOWNER);
        bridgeExecuterCross.addWhitelistTarget(address(mockSwapCross));

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setBridgeExecuter(IBridgeExecuter(address(bridgeExecuterBSC)));

        // Whitelist mock contracts
        vm.prank(OWNER);
        bridgeExecuterBSC.addWhitelistTarget(address(mockTargetBSC));
        vm.prank(OWNER);
        bridgeExecuterBSC.addWhitelistTarget(address(mockSwapBSC));
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

        uint index = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
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

        uint index = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fall back to sending tokens to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = USER.balance;
        uint beforeTargetBalance = address(mockTargetCross).balance;
        uint beforeExecuterBalance = address(bridgeExecuterCross).balance;
        uint beforeBridgeBalance = address(bridgeCross).balance;

        crossFinalize(index, address(NATIVE_TOKEN), USER, value, 5, extraData);

        // Verify tokens were sent to user (fallback behavior)
        assertEq(USER.balance, beforeUserBalance + value, "User should receive tokens");
        assertEq(address(mockTargetCross).balance, beforeTargetBalance, "Target should not receive tokens");
        assertEq(address(bridgeExecuterCross).balance, beforeExecuterBalance, "Executer should return all tokens");
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize - should fall back to minting to user
        vm.selectFork(crossForkID);
        uint beforeUserBalance = testTokenCross.balanceOf(USER);
        uint beforeTargetBalance = testTokenCross.balanceOf(address(mockTargetCross));
        uint beforeExecuterBalance = testTokenCross.balanceOf(address(bridgeExecuterCross));
        uint beforeBridgeBalance = testTokenCross.balanceOf(address(bridgeCross));

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify tokens were minted to user (fallback behavior)
        assertEq(testTokenCross.balanceOf(USER), beforeUserBalance + value, "User should receive tokens");
        assertEq(
            testTokenCross.balanceOf(address(mockTargetCross)), beforeTargetBalance, "Target should not receive tokens"
        );
        assertEq(
            testTokenCross.balanceOf(address(bridgeExecuterCross)),
            beforeExecuterBalance,
            "Executer should not hold tokens (burned)"
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

        uint index = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
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
        uint swapIndex =
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
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
        assertFalse(bridgeExecuterCross.isWhitelistedTarget(testTarget));

        // Add to whitelist
        vm.prank(CrossOWNER);
        bridgeExecuterCross.addWhitelistTarget(testTarget);
        assertTrue(bridgeExecuterCross.isWhitelistedTarget(testTarget));

        // Remove from whitelist
        vm.prank(CrossOWNER);
        bridgeExecuterCross.removeWhitelistTarget(testTarget);
        assertFalse(bridgeExecuterCross.isWhitelistedTarget(testTarget));
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
        uint swapIndex =
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

        uint index = bscBridge(address(testTokenBSC), USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        // Finalize
        vm.selectFork(crossForkID);
        uint beforeWethBalance = weth.balanceOf(USER);

        crossFinalize(index, address(testTokenCross), USER, value, 5, extraData);

        // Verify user received 2x weth (due to 2:1 rate)
        assertEq(weth.balanceOf(USER), beforeWethBalance + expectedOutput);
    }

    // Helper function to bridge with extradata
    function bscBridge(
        address token,
        address from,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes memory extraData
    ) internal returns (uint index) {
        index = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        vm.prank(from);
        if (token == Const.NATIVE_TOKEN) {
            bridgeBSC.bridgeToken{value: value + networkFee + exFee}(
                CROSS_CHAIN_ID, IERC20(token), to, value, networkFee, exFee, extraData
            );
        } else {
            bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(token), to, value, networkFee, exFee, extraData);
        }
    }

    function crossFinalize(
        uint index,
        address token,
        address to,
        uint value,
        uint validatorCount,
        bytes memory extraData
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
        bridgeCross.finalizeBridge(args, v, r, s);
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

        bytes32 structHash = keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, extraData));
        bytes32 digest = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), structHash);

        for (uint i = 0; i < validatorCount; ++i) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], digest);
        }
    }
}
