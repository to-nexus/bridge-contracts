// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BridgeExceptionTest is BridgeTest {
    function test_deposit_with_insufficient_balance() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        bridgeRevertEthereum = true;
        deposit(true, amount + 1000, 5);
    }

    function test_deposit_with_insufficient_validator_signature() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        finalizeRevertCross = true;
        deposit(true, amount, 1);
    }

    function test_withdraw_with_insufficient_balance() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        uint transferValue = amount;
        vm.prank(OWNER);
        cross.transfer(USER, transferValue);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);

        bridgeRevertCross = true;
        withdraw(true, USER.balance + 1000, 5);
    }

    function test_withdraw_with_insufficient_validator_signature() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);

        vm.prank(USER);

        finalizeRevertEthereum = true;
        withdraw(true, amount, 1);
    }

    function test_deposit_with_insufficient_allowance() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount / 2);

        bridgeRevertEthereum = true;
        deposit(true, amount, 5);
    }

    function test_deposit_withdraw_at_token_paused() public {
        // token pause
        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        bridgeEthereum.pauseToken(CROSS_CHAIN_ID, address(cross));

        uint amount = 1000 ether;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        bridgeRevertEthereum = true;
        deposit(true, amount, 5);

        // token unpause
        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        bridgeEthereum.unpauseToken(CROSS_CHAIN_ID, address(cross));

        deposit(false, amount, 5);

        // token pause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.pauseToken(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN));

        bridgeRevertCross = true;
        withdraw(true, amount, 5);

        // token unpause
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.unpauseToken(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN));

        withdraw(false, amount * 10, 5);
    }

    function test_deposit_with_over_safety_limit() public {
        uint amount = 1000 ether;

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setSafetyLimit(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN), amount - 1);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        uint index = deposit(true, amount, 5);

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        vm.expectRevert();
        bridgeCross.retryFinalizeBridge(ETHEREUM_CHAIN_ID, index);

        vm.prank(CrossOWNER);
        bridgeCross.setSafetyLimit(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN), amount);

        vm.prank(CrossOWNER);
        bridgeCross.retryFinalizeBridge(ETHEREUM_CHAIN_ID, index);
    }
}
