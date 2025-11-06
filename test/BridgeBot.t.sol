// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeBot} from "../src/BridgeBot.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {PriceFeed} from "../src/PriceFeed.sol";
import {Const} from "../src/lib/Const.sol";

import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {IBridgeVerifier} from "../src/interface/IBridgeVerifier.sol";
import {TestToken} from "./token/TestToken.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test, console} from "forge-std/Test.sol";

contract BridgeBotTest is Test {
    BridgeBot public bridgeBot;
    MockBridge public mockBridge;
    BridgeVerifier public bridgeVerifier;
    PriceFeed public priceFeed;
    TestToken public testToken;

    address public owner = makeAddr("owner");
    address public editor = makeAddr("editor");
    address public user = makeAddr("user");
    address public recipient = makeAddr("recipient");
    address public executor = makeAddr("executor");

    uint public constant TEST_CHAIN_ID = 97; // BSC testnet
    uint public constant INITIAL_SUPPLY = 1000000 ether;
    uint public constant BRIDGE_AMOUNT = 100 ether;
    uint public constant DAILY_INTERVAL = 86400; // 24 hours

    event BridgeExecuted(
        address indexed tokenAddress,
        uint amount,
        address indexed recipient,
        uint toChainID,
        address executor,
        uint timestamp
    );

    function setUp() public {
        vm.startPrank(owner);

        // Deploy test token
        testToken = new TestToken("Test Token", "TEST", 18);
        testToken.mint(owner, INITIAL_SUPPLY);

        // Deploy price feed
        address priceFeedImpl = address(new PriceFeed());
        ERC1967Proxy priceFeedProxy = new ERC1967Proxy(priceFeedImpl, bytes(""));
        priceFeed = PriceFeed(address(priceFeedProxy));
        priceFeed.initialize(owner, 18);

        // Mock bridge deployment (using a simple contract for testing)
        // We need to create mockBridge first, then create bridgeVerifier with mockBridge address
        mockBridge = new MockBridge(address(0)); // Temporary bridge verifier

        // Deploy bridge verifier
        bridgeVerifier = new BridgeVerifier(
            owner,
            address(mockBridge), // bridge address
            address(priceFeed),
            200000, // finalizeBridgeGas
            0, // verificationAmountThreshold
            0, // defaultTokenPrice
            0, // periodTotalValueThreshold
            0, // tokenScoreThreshold
            0, // tokenCurrentVolumeThreshold
            2 hours // timeWindow
        );

        // Update mock bridge with correct bridge verifier
        mockBridge = new MockBridge(address(bridgeVerifier));

        // Deploy bridge bot (with 0 admin delay for testing)
        bridgeBot = new BridgeBot(address(mockBridge), owner, editor, executor, 0);

        console.log("BridgeBot address:", address(bridgeBot));

        // Fund the bridge bot
        testToken.transfer(address(bridgeBot), BRIDGE_AMOUNT * 10);
        vm.deal(address(bridgeBot), 10 ether);

        vm.stopPrank();
    }

    function testSetConfig() public {
        vm.startPrank(owner);

        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        BridgeBot.BridgeConfig memory config = bridgeBot.getConfig();

        assertEq(config.tokenAddress, address(testToken));
        assertEq(config.recipient, recipient);
        assertEq(config.toChainID, TEST_CHAIN_ID);
        assertEq(config.interval, DAILY_INTERVAL);
        assertEq(config.lastExecuted, 0); // 0 means immediate execution allowed
        assertTrue(config.enabled);

        vm.stopPrank();
    }

    function testExecuteBridge() public {
        vm.startPrank(owner);
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);
        vm.stopPrank();

        // Check if can execute
        (bool canExecute, uint nextAvailableTime) = bridgeBot.canExecuteBridge();
        assertTrue(canExecute);
        assertEq(nextAvailableTime, 0);

        // Execute bridge
        vm.startPrank(executor);
        vm.expectEmit(true, true, true, true);
        emit BridgeExecuted(address(testToken), BRIDGE_AMOUNT, recipient, TEST_CHAIN_ID, executor, block.timestamp);

        bridgeBot.executeBridge(BRIDGE_AMOUNT);
        vm.stopPrank();

        // Check config updated
        BridgeBot.BridgeConfig memory config = bridgeBot.getConfig();
        assertEq(config.lastExecuted, block.timestamp);

        // Should not be able to execute again immediately
        (canExecute,) = bridgeBot.canExecuteBridge();
        assertFalse(canExecute);
    }

    function testExecuteBridgeAfterInterval() public {
        vm.startPrank(owner);
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);
        vm.stopPrank();

        // Execute first time
        vm.prank(executor);
        bridgeBot.executeBridge(BRIDGE_AMOUNT);

        // Should not be able to execute before interval
        (bool canExecute,) = bridgeBot.canExecuteBridge();
        assertFalse(canExecute);

        // Move time forward by interval
        vm.warp(block.timestamp + DAILY_INTERVAL);

        // Should be able to execute now
        (canExecute,) = bridgeBot.canExecuteBridge();
        assertTrue(canExecute);

        // Execute again
        vm.prank(executor);
        bridgeBot.executeBridge(BRIDGE_AMOUNT);
    }

    function testExecuteBridgeWithNativeToken() public {
        vm.startPrank(owner);
        bridgeBot.setConfig(Const.NATIVE_TOKEN, recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);
        vm.stopPrank();

        uint initialBalance = address(bridgeBot).balance;

        // Execute bridge
        vm.prank(executor);
        bridgeBot.executeBridgeNative(1 ether);

        // Check balance decreased
        assertTrue(address(bridgeBot).balance < initialBalance);
    }

    function testUpdateConfig() public {
        vm.startPrank(owner);

        // Set initial config
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // Update config with specific lastExecuted time
        address newRecipient = makeAddr("newRecipient");
        uint newInterval = DAILY_INTERVAL * 2;
        uint updateTime = block.timestamp;
        bridgeBot.setConfig(address(testToken), newRecipient, TEST_CHAIN_ID, newInterval, updateTime);

        BridgeBot.BridgeConfig memory config = bridgeBot.getConfig();
        assertEq(config.recipient, newRecipient);
        assertEq(config.interval, newInterval);
        assertEq(config.lastExecuted, updateTime);

        vm.stopPrank();
    }

    function testToggleConfig() public {
        vm.startPrank(owner);

        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // Disable config
        bridgeBot.setEnabled(false);

        BridgeBot.BridgeConfig memory config = bridgeBot.getConfig();
        assertFalse(config.enabled);

        // Try to execute (should fail)
        vm.stopPrank();
        vm.prank(executor);
        vm.expectRevert("Config not available");
        bridgeBot.executeBridge(BRIDGE_AMOUNT);

        // Enable config
        vm.prank(owner);
        bridgeBot.setEnabled(true);

        config = bridgeBot.getConfig();
        assertTrue(config.enabled);
    }

    function testWithdrawToken() public {
        uint withdrawAmount = BRIDGE_AMOUNT;
        uint initialBalance = testToken.balanceOf(user);
        vm.prank(owner);
        bridgeBot.withdrawToken(address(testToken), user, withdrawAmount);

        assertEq(testToken.balanceOf(user), initialBalance + withdrawAmount);
    }

    function testWithdrawNative() public {
        vm.prank(owner);

        uint withdrawAmount = 1 ether;
        uint initialBalance = user.balance;

        bridgeBot.withdrawNative(payable(user), withdrawAmount);

        assertEq(user.balance, initialBalance + withdrawAmount);
    }

    function testWithdrawAllTokens() public {
        uint botBalance = testToken.balanceOf(address(bridgeBot));
        uint initialUserBalance = testToken.balanceOf(user);
        vm.prank(owner);
        bridgeBot.withdrawAllTokens(address(testToken), user);

        assertEq(testToken.balanceOf(address(bridgeBot)), 0);
        assertEq(testToken.balanceOf(user), initialUserBalance + botBalance);
    }

    function testWithdrawAllNative() public {
        vm.prank(owner);

        uint botBalance = address(bridgeBot).balance;
        uint initialUserBalance = user.balance;

        bridgeBot.withdrawAllNative(payable(user));

        assertEq(address(bridgeBot).balance, 0);
        assertEq(user.balance, initialUserBalance + botBalance);
    }

    function testOnlyEditorFunctions() public {
        // User without editor role tries to set config
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // User without editor role tries to toggle config
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.setEnabled(false);
    }

    function testOnlyExecutorCanBridge() public {
        vm.prank(owner);
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // User without executor role tries to execute bridge
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.executeBridge(BRIDGE_AMOUNT);
    }

    function testOnlyAdminFunctions() public {
        // User without admin role tries to withdraw token
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.withdrawToken(address(testToken), user, BRIDGE_AMOUNT);

        // User without admin role tries to withdraw native
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.withdrawNative(payable(user), 1 ether);
    }

    function testRoleManagement() public {
        bytes32 defaultAdminRole = bridgeBot.DEFAULT_ADMIN_ROLE();
        bytes32 editorRole = bridgeBot.EDITOR_ROLE();
        bytes32 executorRole = bridgeBot.EXECUTOR_ROLE();

        // Check initial roles
        assertTrue(bridgeBot.hasRole(defaultAdminRole, owner));
        assertTrue(bridgeBot.hasRole(editorRole, owner));
        assertTrue(bridgeBot.hasRole(editorRole, editor));
        assertTrue(bridgeBot.hasRole(executorRole, owner));
        assertTrue(bridgeBot.hasRole(executorRole, executor));

        // Grant executor role to user
        vm.prank(owner);
        bridgeBot.grantRole(editorRole, user);

        assertTrue(bridgeBot.hasRole(editorRole, user));

        // Revoke executor role from user
        vm.prank(owner);
        bridgeBot.revokeRole(editorRole, user);

        assertFalse(bridgeBot.hasRole(editorRole, user));
    }

    function testEditorCanManageConfigs() public {
        // Editor can set config
        vm.prank(editor);
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        BridgeBot.BridgeConfig memory config = bridgeBot.getConfig();
        assertEq(config.tokenAddress, address(testToken));

        // Editor can toggle config
        vm.prank(editor);
        bridgeBot.setEnabled(false);

        config = bridgeBot.getConfig();
        assertFalse(config.enabled);
    }

    function testExecuteBridgeWrongTokenType() public {
        // Setup ERC20 config
        vm.prank(owner);
        bridgeBot.setConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // Try to execute native bridge with ERC20 config
        vm.prank(executor);
        vm.expectRevert("Use executeBridge for ERC20 tokens");
        bridgeBot.executeBridgeNative(1 ether);

        // Setup Native config
        vm.prank(owner);
        bridgeBot.setConfig(Const.NATIVE_TOKEN, recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        // Try to execute ERC20 bridge with Native config
        vm.prank(executor);
        vm.expectRevert("Use executeBridgeNative for native tokens");
        bridgeBot.executeBridge(1 ether);
    }
}

/**
 * @dev Mock Bridge contract for testing
 */
contract MockBridge {
    BridgeVerifier private _bridgeVerifier;

    constructor(address bridgeVerifierAddr) {
        _bridgeVerifier = BridgeVerifier(bridgeVerifierAddr);
    }

    function bridgeVerifier() public view returns (IBridgeVerifier) {
        return IBridgeVerifier(address(_bridgeVerifier));
    }

    function bridgeToken(uint, IERC20 token, address, uint amount, uint gasFee, uint exFee, bytes calldata)
        external
        payable
        returns (bool)
    {
        if (address(token) != address(1)) token.transferFrom(msg.sender, address(this), amount + gasFee + exFee);
        return true;
    }

    function chainID() external pure returns (uint) {
        return 97;
    }
}
