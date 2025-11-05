// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeBot} from "../src/BridgeBot.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {PriceFeed} from "../src/PriceFeed.sol";

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

    address public constant NATIVE_TOKEN = address(1);
    address public owner = address(0x1);
    address public user = address(0x2);
    address public recipient = address(0x3);
    address public executor = address(0x4);

    uint public constant TEST_CHAIN_ID = 97; // BSC testnet
    uint public constant INITIAL_SUPPLY = 1000000 ether;
    uint public constant BRIDGE_AMOUNT = 100 ether;
    uint public constant DAILY_INTERVAL = 86400; // 24 hours

    event BridgeExecuted(
        uint indexed configId,
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

        // Deploy bridge bot (with 1 day admin delay)
        bridgeBot = new BridgeBot(address(mockBridge), owner, executor, 1 days);

        // Fund the bridge bot
        testToken.transfer(address(bridgeBot), BRIDGE_AMOUNT * 10);
        vm.deal(address(bridgeBot), 10 ether);

        vm.stopPrank();
    }

    function testAddBridgeConfig() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);

        assertEq(config.tokenAddress, address(testToken));
        assertEq(config.recipient, recipient);
        assertEq(config.toChainID, TEST_CHAIN_ID);
        assertEq(config.interval, DAILY_INTERVAL);
        assertEq(config.lastExecuted, 0);
        assertTrue(config.enabled);

        vm.stopPrank();
    }

    function testExecuteBridge() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        vm.stopPrank();

        // Check if can execute
        (bool canExecute, uint nextAvailableTime) = bridgeBot.canExecuteBridge(configId);
        assertTrue(canExecute);
        assertEq(nextAvailableTime, 0);

        // Execute bridge
        vm.startPrank(executor);
        vm.expectEmit(true, true, true, true);
        emit BridgeExecuted(
            configId, address(testToken), BRIDGE_AMOUNT, recipient, TEST_CHAIN_ID, executor, block.timestamp
        );

        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);
        vm.stopPrank();

        // Check config updated
        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);
        assertEq(config.lastExecuted, block.timestamp);

        // Should not be able to execute again immediately
        (canExecute,) = bridgeBot.canExecuteBridge(configId);
        assertFalse(canExecute);
    }

    function testExecuteBridgeAfterInterval() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        vm.stopPrank();

        // Execute first time
        vm.prank(executor);
        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);

        // Should not be able to execute immediately
        (bool canExecute,) = bridgeBot.canExecuteBridge(configId);
        assertFalse(canExecute);

        // Fast forward to next day
        vm.warp(block.timestamp + DAILY_INTERVAL);

        // Should be able to execute again
        (canExecute,) = bridgeBot.canExecuteBridge(configId);
        assertTrue(canExecute);

        // Execute again
        vm.prank(executor);
        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);

        // Verify execution
        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);
        assertEq(config.lastExecuted, block.timestamp);
    }

    function testExecuteBridgeWithNativeToken() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(NATIVE_TOKEN, recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        vm.stopPrank();

        // Execute bridge with native token
        vm.prank(executor);
        bridgeBot.executeBridgeNative(configId, 1 ether);

        // Verify execution
        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);
        assertEq(config.lastExecuted, block.timestamp);
    }

    function testExecuteBridgeWrongTokenType() public {
        vm.startPrank(owner);

        // Add ERC20 config
        uint erc20ConfigId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        // Add native config
        uint nativeConfigId = bridgeBot.addBridgeConfig(NATIVE_TOKEN, recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        vm.stopPrank();

        // Try to execute ERC20 config with executeBridgeNative - should fail
        vm.prank(executor);
        vm.expectRevert("Use executeBridge for ERC20 tokens");
        bridgeBot.executeBridgeNative(erc20ConfigId, 1 ether);

        // Try to execute native config with executeBridge - should fail
        vm.prank(executor);
        vm.expectRevert("Use executeBridgeNative for native tokens");
        bridgeBot.executeBridge(nativeConfigId, 1 ether);
    }

    function testUpdateBridgeConfig() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        // Update config
        bridgeBot.updateBridgeConfig(
            configId,
            NATIVE_TOKEN,
            user,
            56, // BSC mainnet
            3600, // 1 hour
            0 // lastExecuted (0 = use current block.timestamp)
        );

        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);
        assertEq(config.tokenAddress, NATIVE_TOKEN);
        assertEq(config.recipient, user);
        assertEq(config.toChainID, 56);
        assertEq(config.interval, 3600);

        vm.stopPrank();
    }

    function testToggleBridgeConfig() public {
        vm.startPrank(owner);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        // Disable config
        bridgeBot.toggleBridgeConfig(configId, false);

        BridgeBot.BridgeConfig memory config = bridgeBot.getBridgeConfig(configId);
        assertFalse(config.enabled);

        // Should not be able to execute disabled config
        (bool canExecute,) = bridgeBot.canExecuteBridge(configId);
        assertFalse(canExecute);

        vm.stopPrank();
    }

    function testWithdrawTokens() public {
        uint withdrawAmount = 50 ether;

        vm.startPrank(owner);

        uint balanceBefore = testToken.balanceOf(owner);
        bridgeBot.withdrawToken(address(testToken), owner, withdrawAmount);
        uint balanceAfter = testToken.balanceOf(owner);

        assertEq(balanceAfter - balanceBefore, withdrawAmount);

        vm.stopPrank();
    }

    function testWithdrawNative() public {
        uint withdrawAmount = 1 ether;

        vm.startPrank(owner);

        uint balanceBefore = owner.balance;
        bridgeBot.withdrawNative(payable(owner), withdrawAmount);
        uint balanceAfter = owner.balance;

        assertEq(balanceAfter - balanceBefore, withdrawAmount);

        vm.stopPrank();
    }

    function testOnlyEditorFunctions() public {
        // Non-editor should not be able to call editor functions
        vm.startPrank(user);

        vm.expectRevert();
        bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        vm.expectRevert();
        bridgeBot.updateBridgeConfig(0, address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL, 0);

        vm.expectRevert();
        bridgeBot.toggleBridgeConfig(0, false);

        vm.stopPrank();
    }

    function testOnlyAdminFunctions() public {
        // Non-admin should not be able to call admin functions
        vm.startPrank(user);

        vm.expectRevert();
        bridgeBot.withdrawToken(address(testToken), user, 1 ether);

        vm.expectRevert();
        bridgeBot.withdrawNative(payable(user), 1 ether);

        vm.stopPrank();
    }

    function testEditorCanManageConfigs() public {
        // Grant editor role to user
        vm.prank(owner);
        bridgeBot.grantRole(bridgeBot.EDITOR_ROLE(), user);

        // Editor should be able to manage configs
        vm.startPrank(user);

        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);
        assertEq(configId, 0);

        bridgeBot.updateBridgeConfig(configId, address(testToken), user, TEST_CHAIN_ID, 3600, 0);
        bridgeBot.toggleBridgeConfig(configId, false);

        // But editor cannot withdraw
        vm.expectRevert();
        bridgeBot.withdrawToken(address(testToken), user, 1 ether);

        vm.stopPrank();
    }

    function testRoleManagement() public {
        // Test initial roles
        console.log("Owner address:", owner);
        console.log("msg.sender:", msg.sender);
        console.log("Has DEFAULT_ADMIN_ROLE:", bridgeBot.hasRole(bridgeBot.DEFAULT_ADMIN_ROLE(), owner));
        console.log("DEFAULT_ADMIN:", bridgeBot.defaultAdmin());
        console.log("EDITOR_ROLE admin:", uint(bridgeBot.getRoleAdmin(bridgeBot.EDITOR_ROLE())));
        console.log("DEFAULT_ADMIN_ROLE value:", uint(bridgeBot.DEFAULT_ADMIN_ROLE()));

        assertTrue(bridgeBot.hasRole(bridgeBot.DEFAULT_ADMIN_ROLE(), owner));
        assertTrue(bridgeBot.hasRole(bridgeBot.EDITOR_ROLE(), owner));
        assertTrue(bridgeBot.hasRole(bridgeBot.EXECUTOR_ROLE(), owner));
        assertTrue(bridgeBot.hasRole(bridgeBot.EXECUTOR_ROLE(), executor));

        // Test granting editor role
        console.log("About to grant EDITOR_ROLE, msg.sender will be:", owner);
        vm.prank(owner);
        bridgeBot.grantRole(bridgeBot.EDITOR_ROLE(), user);
        assertTrue(bridgeBot.hasRole(bridgeBot.EDITOR_ROLE(), user));

        // Test revoking editor role
        vm.prank(owner);
        bridgeBot.revokeRole(bridgeBot.EDITOR_ROLE(), user);
        assertFalse(bridgeBot.hasRole(bridgeBot.EDITOR_ROLE(), user));

        // Test granting executor role
        vm.prank(owner);
        bridgeBot.grantRole(bridgeBot.EXECUTOR_ROLE(), user);
        assertTrue(bridgeBot.hasRole(bridgeBot.EXECUTOR_ROLE(), user));

        // Test revoking executor role
        vm.prank(owner);
        bridgeBot.revokeRole(bridgeBot.EXECUTOR_ROLE(), executor);
        assertFalse(bridgeBot.hasRole(bridgeBot.EXECUTOR_ROLE(), executor));

        // Test only admin can manage roles
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.grantRole(bridgeBot.EDITOR_ROLE(), address(0x5));

        vm.prank(user);
        vm.expectRevert();
        bridgeBot.grantRole(bridgeBot.EXECUTOR_ROLE(), address(0x5));

        vm.prank(user);
        vm.expectRevert();
        bridgeBot.revokeRole(bridgeBot.EXECUTOR_ROLE(), owner);
    }

    function testOnlyExecutorCanBridge() public {
        vm.prank(owner);
        uint configId = bridgeBot.addBridgeConfig(address(testToken), recipient, TEST_CHAIN_ID, DAILY_INTERVAL);

        // Non-executor should not be able to execute bridge
        vm.prank(user);
        vm.expectRevert();
        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);

        // Executor should be able to execute bridge
        vm.prank(executor);
        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);

        // Owner (who has executor role) should also be able to execute bridge
        vm.warp(block.timestamp + DAILY_INTERVAL + 1);
        vm.prank(owner);
        bridgeBot.executeBridge(configId, BRIDGE_AMOUNT);
    }
}

// Mock bridge contract for testing
contract MockBridge {
    IBridgeVerifier private _bridgeVerifier;

    constructor(address bridgeVerifierAddr) {
        _bridgeVerifier = IBridgeVerifier(bridgeVerifierAddr);
    }

    function bridgeToken(uint, IERC20, address, uint, uint, uint, bytes calldata) external payable returns (bool) {
        return true;
    }

    function bridgeVerifier() external view returns (IBridgeVerifier) {
        return _bridgeVerifier;
    }

    // Mock other required functions
    function finalizeBridge(
        IBridgeRegistry.FinalizeArguments calldata,
        uint8[] memory,
        bytes32[] memory,
        bytes32[] memory
    ) external payable returns (bool) {
        return true;
    }

    function permitBridgeToken(
        uint,
        IERC20,
        address,
        uint,
        uint,
        uint,
        bytes calldata,
        IBaseBridge.PermitArguments calldata
    ) external payable returns (bool) {
        return true;
    }

    function permitBridgeTokenBatch(IBaseBridge.BridgeTokenArguments[] calldata, IBaseBridge.PermitArguments[] calldata)
        external
        payable
    {}

    function finalizeBridgeBatch(
        IBridgeRegistry.FinalizeArguments[] calldata,
        uint8[][] memory,
        bytes32[][] memory,
        bytes32[][] memory
    ) external payable returns (bool) {
        return true;
    }

    function releasePending(uint, uint) external {}

    function releasePendingBatch(uint[] memory, uint[] memory) external {}

    function domainSeparator() external pure returns (bytes32) {
        return bytes32(0);
    }

    function initializedAt() external pure returns (uint) {
        return 0;
    }
}
