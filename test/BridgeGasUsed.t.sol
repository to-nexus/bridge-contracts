// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeTest} from "./Bridge.t.sol";
import {console} from "forge-std/Test.sol";

contract BridgeGasUsedTest is BridgeTest {
    function setUp() public override {
        super.setUp();
        // enableGasUsedLog = true;
    }

    function testLogBridgeGasused() public {
        uint reserve = 100000 * 1e18;
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, reserve);
        vm.prank(OWNER);
        cross.approve(address(bridgeBSC), reserve);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), reserve);

        // cycle1
        if (enableGasUsedLog) console.log("bridge coin cycle1");
        deposit(false, amount, 5);
        withdraw(false, amount, 5);

        // cycle2
        if (enableGasUsedLog) console.log("\nbridge coin cycle2");
        deposit(false, amount, 5);
        withdraw(false, amount, 5);

        // cycle3
        if (enableGasUsedLog) console.log("\nbridge coin cycle3");
        deposit(false, amount, 5);
        withdraw(false, amount, 5);

        // owner key bridge
        (uint index, bool ok) = bscBridge(address(cross), OWNER, OWNER, amount, 0, 0);
        if (ok) {
            if (enableGasUsedLog) console.log("owner bridge token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
            bscIncrementIndex();
            ok = crossFinalize(index, address(NATIVE_TOKEN), OWNER, amount, 3);
            if (ok && enableGasUsedLog) console.log("owner finalize coin", uint(vm.lastCallGas().gasTotalUsed));
        }

        (uint value, uint gas, uint ex) = crossCalcFee(NATIVE_TOKEN, amount);
        assertTrue(value + gas + ex <= amount);

        assertTrue(value > 0);
        (index, ok) = crossBridge(address(NATIVE_TOKEN), OWNER, OWNER, value, gas, ex);
        if (ok) {
            if (enableGasUsedLog) console.log("owner bridge coin", uint(vm.lastCallGas().gasTotalUsed));
            crossIncrementIndex();
            ok = bscFinalize(index, address(cross), OWNER, amount, 3);
            if (ok && enableGasUsedLog) {
                console.log("owner finalize token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
            }
        }

        // test token
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, reserve);
        vm.prank(OWNER);
        testTokenBSC.approve(address(bridgeBSC), reserve);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), reserve);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), reserve);
        vm.prank(OWNER);
        testTokenCross.approve(address(bridgeCross), reserve);

        // cycle1
        if (enableGasUsedLog) console.log("bridge token cycle1");
        depositToken(false, amount, 5);
        withdrawToken(false, amount, 5);

        // cycle2
        if (enableGasUsedLog) console.log("\nbridge token cycle2");
        depositToken(false, amount, 5);
        withdrawToken(false, amount, 5);

        // cycle3
        if (enableGasUsedLog) console.log("\nbridge token cycle3");
        depositToken(false, amount, 5);
        withdrawToken(false, amount, 5);

        // owner
        (index, ok) = bscBridge(address(testTokenBSC), OWNER, OWNER, amount, 0, 0);
        if (ok) {
            bscIncrementIndex();
            crossFinalize(index, address(testTokenCross), OWNER, amount, 3);
            if (ok && enableGasUsedLog) console.log("owner finalize token (mint)", uint(vm.lastCallGas().gasTotalUsed));
        }

        (value, gas, ex) = crossCalcFee(testTokenCross, amount);
        assertTrue(value + gas + ex <= amount);
        (index, ok) = crossBridge(address(testTokenCross), OWNER, OWNER, value, gas, ex);
        if (ok) {
            if (enableGasUsedLog) console.log("owner bridge token (burn)", uint(vm.lastCallGas().gasTotalUsed));
            crossIncrementIndex();
            bscFinalize(index, address(testTokenBSC), OWNER, amount, 3);
        }
    }
}
