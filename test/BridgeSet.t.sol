// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";

contract BridgeSetTest is BridgeTest {
    function test_set_validator() public {
        vm.startPrank(OWNER);

        vm.selectFork(ethereumForkID);
        bridgeEthereum.revokeRoleBatch(VALIDATOR_ROLE, VALIDATORS);
        bool ok = bridgeEthereum.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertFalse(ok);
        bridgeEthereum.grantRoleBatch(VALIDATOR_ROLE, VALIDATORS);
        ok = bridgeEthereum.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertTrue(ok);
        vm.stopPrank();

        vm.startPrank(CrossOWNER);
        vm.selectFork(crossForkID);
        bridgeCross.revokeRoleBatch(VALIDATOR_ROLE, VALIDATORS);
        ok = bridgeCross.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertFalse(ok);
        bridgeCross.grantRoleBatch(VALIDATOR_ROLE, VALIDATORS);
        ok = bridgeCross.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertTrue(ok);
        vm.stopPrank();
    }

    function test_set_pause() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        bridgeEthereum.setPause(true);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(true);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        bridgeRevertEthereum = true;
        vm.prank(USER);
        deposit(true, amount, 5);

        bridgeRevertCross = true;
        vm.prank(USER);
        withdraw(true, amount, 5);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        bridgeEthereum.setPause(false);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(false);

        vm.prank(USER);
        deposit(false, amount, 5);
        vm.prank(USER);
        withdraw(false, amount, 5);
    }

    function test_set_threshold() public {
        uint amount = 1000;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        vm.prank(USER);
        deposit(false, amount, 3);

        vm.selectFork(ethereumForkID);
        threshold = 4;
        vm.prank(OWNER);
        bridgeEthereum.changeThreshold(threshold);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.changeThreshold(threshold);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        finalizeRevertCross = true;
        vm.prank(USER);
        uint index = deposit(true, amount, 3);

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), USER, amount, 4);

        vm.selectFork(ethereumForkID);
        threshold = 1;
        vm.prank(OWNER);
        bridgeEthereum.changeThreshold(threshold);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.changeThreshold(threshold);

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        vm.prank(USER);
        deposit(false, amount, 1);
    }
}
