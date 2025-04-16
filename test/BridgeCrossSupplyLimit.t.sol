// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {console} from "forge-std/Test.sol";

/**
 * @title BridgeCrossSupplyLimitTest
 * @notice Test for CrossBridge supply limit functionality
 */
contract BridgeCrossSupplyLimitTest is BridgeTest {
    function setUp() public virtual override {
        super.setUp();

        // Ensure USER has enough tokens for tests
        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);
        cross.transfer(USER, 1000 ether);
        vm.stopPrank();

        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);
    }

    // Test setting cross supply limit
    function test_set_cross_supply_limit() public {
        vm.selectFork(crossForkID);

        // Set cross supply limit as admin
        vm.startPrank(CrossOWNER);
        bridgeCross.setCrossSupplyLimit(1000 ether + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        // Try to set cross supply limit as non-admin (should revert)
        vm.startPrank(USER);
        vm.expectRevert();
        bridgeCross.setCrossSupplyLimit(500 ether + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();
    }

    // Test cross supply limit exceeded
    function test_cross_supply_limit_exceeded() public {
        // Setup: First set a low cross supply limit
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        uint supplyLimit = 10 ether;
        bridgeCross.setCrossSupplyLimit(supplyLimit + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        // First bridge: should succeed (below limit)
        uint smallAmount = 5 ether;
        uint index = deposit(false, smallAmount, threshold);

        // Second bridge: should fail (exceeds limit)
        uint largeAmount = 10 ether;

        // Handle the bsc side
        vm.selectFork(bscForkID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));

        (index,) = bscBridge(address(cross), USER, USER, largeAmount, 0, 0);
        bscIncrementIndex();

        // Move to cross chain for finalization - should fail due to limit
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;

        // Try to finalize - should fail with CrossSupplyLimitExceeded
        crossFinalize(index, address(NATIVE_TOKEN), USER, largeAmount, threshold);

        // Verify that finalization was not successful due to limit
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.CrossSupplyLimitExceeded,
            "Finalization should fail due to supply limit"
        );

        // Check that token balances didn't change on finalize attempt
        vm.selectFork(bscForkID);
        assertEq(userTokenBalance - largeAmount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + largeAmount, cross.balanceOf(address(bridgeBSC)));

        // Check that coin balances didn't change on finalize attempt
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance, USER.balance);

        // Test manual release pending
        vm.prank(CrossOWNER);
        bridgeCross.manualReleasePending(BSC_CHAIN_ID, index);
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None);
    }

    // Test using whole supply limit
    function test_using_full_cross_supply_limit() public {
        // Setup: Set a specific cross supply limit
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        uint supplyLimit = 100 ether;
        bridgeCross.setCrossSupplyLimit(supplyLimit + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        // First deposit exactly at the limit
        deposit(false, supplyLimit, threshold);

        // Next deposit should fail (exceeds limit)
        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, 1 ether, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        // Try to finalize - should fail
        crossFinalize(index, address(NATIVE_TOKEN), USER, 1 ether, threshold);

        // Verify that finalization was not successful due to limit
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.CrossSupplyLimitExceeded,
            "Finalization should fail due to supply limit"
        );
    }

    // Test updating cross supply limit
    function test_update_cross_supply_limit() public {
        // Setup: First set a low cross supply limit
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        uint supplyLimit = 10 ether;
        bridgeCross.setCrossSupplyLimit(supplyLimit + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        // Bridge amount near limit
        uint amount = 9 ether;
        deposit(false, amount, threshold);

        // Increase limit as admin
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        uint newSupplyLimit = 100 ether;
        bridgeCross.setCrossSupplyLimit(newSupplyLimit + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        // Bridge additional amount should now succeed
        deposit(false, 50 ether, threshold);
    }
}
