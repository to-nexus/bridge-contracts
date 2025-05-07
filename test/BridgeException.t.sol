// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BridgeExceptionTest is BridgeTest {
    function testDepositWithInsufficientBalance() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        deposit(true, amount + 1000, 5);
    }

    function testDepositWithInsufficientValidatorSignature() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        finalizeRevertCross = true;
        deposit(true, amount, 1);
    }

    function testWithdrawWithInsufficientBalance() public {
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

    function testWithdrawWithInsufficientValidatorSignature() public {
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

    function testDepositWithInsufficientAllowance() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount / 2);

        bridgeRevertBSC = true;
        deposit(true, amount, 5);
    }

    function testDepositWithdrawAtTokenPaused() public {
        // token pause
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), true);

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
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), false);

        deposit(false, amount, 5);

        // token pause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), true);

        bridgeRevertCross = true;
        withdraw(true, amount, 5);

        // token unpause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false);

        withdraw(false, amount, 5);
    }
}
