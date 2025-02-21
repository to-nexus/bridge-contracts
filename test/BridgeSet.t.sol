// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";

contract BridgeSetTest is BridgeTest {
    function test_set_validator() public {
        vm.startPrank(OWNER);

        vm.selectFork(ethereumChainID);
        bridgeEthereum.removeValidator(VALIDATOR5);
        bool ok = bridgeEthereum.isValidator(VALIDATOR5);
        vm.assertFalse(ok);
        bridgeEthereum.setValidator(VALIDATOR5);
        ok = bridgeEthereum.isValidator(VALIDATOR5);
        vm.assertTrue(ok);

        vm.selectFork(crossChainID);
        bridgeCross.removeValidator(VALIDATOR5);
        ok = bridgeCross.isValidator(VALIDATOR5);
        vm.assertFalse(ok);
        bridgeCross.setValidator(VALIDATOR5);
        ok = bridgeCross.isValidator(VALIDATOR5);
        vm.assertTrue(ok);

        vm.stopPrank();
    }

    function test_set_pause() public {
        uint amount = 1000 ether;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        bridgeEthereum.pause();
        vm.selectFork(crossChainID);
        vm.prank(OWNER);
        bridgeCross.pause();

        vm.selectFork(ethereumChainID);
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

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        bridgeEthereum.unpause();
        vm.selectFork(crossChainID);
        vm.prank(OWNER);
        bridgeCross.unpause();

        vm.prank(USER);
        deposit(false, amount, 5);
        vm.prank(USER);
        withdraw(false, amount, 5);
    }

    function test_set_threshold() public {
        uint amount = 1000;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        vm.prank(USER);
        deposit(false, amount, 3);

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        bridgeEthereum.changeThreshold(4);
        vm.selectFork(crossChainID);
        vm.prank(OWNER);
        bridgeCross.changeThreshold(4);

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        finalizeRevertCross = true;
        vm.prank(USER);
        uint index = deposit(true, amount, 3);

        vm.selectFork(crossChainID);
        crossFinalize(index, address(coin), amount, 4);

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        bridgeEthereum.changeThreshold(1);
        vm.selectFork(crossChainID);
        vm.prank(OWNER);
        bridgeCross.changeThreshold(1);

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        vm.prank(USER);
        deposit(false, amount, 1);
    }
}
