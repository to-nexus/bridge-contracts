// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BridgeExceptionTest is BridgeTest {
    function test_deposit_with_insufficient_balance() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        deposit(true, amount + 1000, 5);
    }

    function test_deposit_with_insufficient_validator_signature() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        finalizeRevertCross = true;
        deposit(true, amount, 1);
    }

    function test_withdraw_with_insufficient_balance() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        uint transferValue = amount;
        vm.prank(OWNER);
        cross.transfer(USER, transferValue);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        deposit(false, amount, 5);

        bridgeRevertCross = true;
        withdraw(true, USER.balance + 1000, 5);
    }

    function test_withdraw_with_insufficient_validator_signature() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        deposit(false, amount, 5);

        vm.prank(USER);

        finalizeRevertBSC = true;
        withdraw(true, amount, 1);
    }

    function test_deposit_with_insufficient_allowance() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount / 2);

        bridgeRevertBSC = true;
        deposit(true, amount, 5);
    }

    function test_deposit_withdraw_at_token_paused() public {
        // token pause
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), true, false);

        uint amount = 1000 ether;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        deposit(true, amount, 5);

        // token unpause
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), false, false);

        deposit(false, amount, 5);

        // token pause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), true, false);

        bridgeRevertCross = true;
        withdraw(true, amount, 5);

        // token unpause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);

        withdraw(false, amount, 5);
    }

    function test_deposit_with_zero_recipient() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        vm.selectFork(bscForkID);
        vm.prank(USER);
        vm.expectRevert();
        bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(address(cross)), address(0), amount, 0, 0, "");
    }

    function test_finalize_at_token_paused() public {
        uint amount = 1000 ether;

        // First deposit tokens normally (BSC -> Cross)
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        // Set finalize pause on Cross chain for NATIVE_TOKEN before deposit
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, true);

        // Initiate on BSC
        vm.selectFork(bscForkID);
        vm.prank(USER);
        bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(address(cross)), USER, amount, 0, 0, "");
        bscIncrementIndex();

        // Finalize on Cross should go to pending due to finalize pause
        vm.selectFork(crossForkID);
        uint userBalanceBefore = USER.balance;
        crossFinalize(1, address(NATIVE_TOKEN), USER, amount, 5);

        // User should NOT receive tokens (pending)
        assertEq(USER.balance, userBalanceBefore, "User should not receive tokens when finalize paused");

        // Unpause finalize
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);

        // Release pending - now user should receive
        bridgeCross.releasePending(BSC_CHAIN_ID, 1);
        assertEq(USER.balance, userBalanceBefore + amount, "User should receive tokens after unpause");
    }

    function test_finalize_at_chain_paused() public {
        uint amount = 1000 ether;

        // First deposit tokens normally (BSC -> Cross)
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        // Initiate on BSC
        vm.selectFork(bscForkID);
        vm.prank(USER);
        bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(address(cross)), USER, amount, 0, 0, "");
        bscIncrementIndex();

        // Set chain pause on Cross chain before finalize
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setChainPause(BSC_CHAIN_ID, true);

        // Finalize on Cross should revert due to chain pause
        finalizeRevertCross = true;
        crossFinalize(1, address(NATIVE_TOKEN), USER, amount, 5);
    }
}
