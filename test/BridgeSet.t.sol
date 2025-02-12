// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeStandardTest} from "./BridgeStandard.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test, console} from "forge-std/Test.sol";

contract BridgeSetTest is Test, BridgeStandardTest {
    function test_set_validator() public {
        vm.startPrank(OWNER);

        bridgeEthereum.b.removeValidator(VALIDATOR5);
        bridgeCross.b.removeValidator(VALIDATOR5);

        bool ok = bridgeEthereum.b.isValidator(VALIDATOR5);
        vm.assertFalse(ok);
        ok = bridgeCross.b.isValidator(VALIDATOR5);
        vm.assertFalse(ok);

        bridgeEthereum.b.setValidator(VALIDATOR5);
        bridgeCross.b.setValidator(VALIDATOR5);

        ok = bridgeEthereum.b.isValidator(VALIDATOR5);
        vm.assertTrue(ok);
        ok = bridgeCross.b.isValidator(VALIDATOR5);
        vm.assertTrue(ok);

        vm.stopPrank();
    }

    function test_set_pause() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        deposit(false, amount, 5);

        vm.prank(OWNER);
        bridgeEthereum.b.pause();
        vm.prank(OWNER);
        bridgeCross.b.pause();

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        bridgeRevert = true;
        vm.prank(USER);
        deposit(true, amount, 5);

        bridgeRevert = true;
        vm.prank(USER);
        withdraw(true, amount, 5);

        vm.prank(OWNER);
        bridgeEthereum.b.unpause();
        vm.prank(OWNER);
        bridgeCross.b.unpause();

        vm.prank(USER);
        deposit(false, amount, 5);
        vm.prank(USER);
        withdraw(false, amount, 5);
    }

    function test_set_threshold() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        vm.prank(USER);
        deposit(false, amount, 3);

        vm.prank(OWNER);
        bridgeEthereum.b.changeThreshold(4);
        vm.prank(OWNER);
        bridgeCross.b.changeThreshold(4);

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        finalizeRevert = true;
        vm.prank(USER);
        uint index = deposit(true, amount, 3);
        finalize(bridgeEthereum.otherBridge, index, address(xcross), amount, 4);

        vm.prank(OWNER);
        bridgeEthereum.b.changeThreshold(1);
        vm.prank(OWNER);
        bridgeCross.b.changeThreshold(1);

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        vm.prank(USER);
        deposit(false, amount, 1);
    }
}
