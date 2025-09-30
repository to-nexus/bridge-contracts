// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {BSCBridgeV2} from "../src/BSCBridgeV2.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {Const} from "../src/lib/Const.sol";

import {PriceFeed} from "../src/PriceFeed.sol";

import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {SettingTest} from "./chain/Setting.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {console} from "forge-std/Test.sol";

contract BSCBridgeV2Test is BridgeTest {
    BSCBridgeV2 internal bscBridgeV2Impl;
    BSCBridgeV2 internal bridgeBSCV2;

    address internal deadWallet = address(0x000000000000000000000000000000000000dEaD);

    function setUp() public override {
        super.setUp();

        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);
        bridgeBSC.upgradeToAndCall(
            address(new BSCBridgeV2()), abi.encodeCall(BSCBridgeV2.reinitializeBSCBridge, (CROSS_CHAIN_ID, cross))
        );
        vm.stopPrank();

        bridgeBSCV2 = BSCBridgeV2(address(bridgeBSC));
    }

    function test_burn_cross_to_dead_wallet() public {
        vm.selectFork(bscForkID);
        burnCrossToDeadWallet(false, OWNER, deadWallet, 1000 ether, true);
    }

    function test_burn_cross_to_dead_wallet_with_transfer() public {
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.approve(address(bridgeBSCV2), 1000 ether);
        burnCrossToDeadWallet(false, OWNER, deadWallet, 1000 ether, false);
    }

    function test_fuzz_burn_cross_to_dead_wallet(uint amount) public {
        vm.assume(amount > 0);

        vm.selectFork(crossForkID);
        vm.assume(amount < address(bridgeCross).balance);

        // disable verification
        vm.startPrank(CrossOWNER);
        bridgeCross.setCrossSupplyLimit(INITIAL_SUPPLY);
        bridgeVerifierCross.setVerificationAmountThreshold(0);
        bridgeVerifierCross.setPeriodTotalValueThreshold(0);
        vm.stopPrank();

        vm.selectFork(bscForkID);
        burnCrossToDeadWallet(false, OWNER, deadWallet, amount, true);
    }

    function test_fuzz_burn_cross_to_dead_wallet_with_transfer(uint amount) public {
        vm.assume(amount > 0);

        vm.selectFork(crossForkID);
        vm.assume(amount < address(bridgeCross).balance);

        // disable verification
        vm.startPrank(CrossOWNER);
        bridgeCross.setCrossSupplyLimit(INITIAL_SUPPLY);
        bridgeVerifierCross.setVerificationAmountThreshold(0);
        bridgeVerifierCross.setPeriodTotalValueThreshold(0);
        vm.stopPrank();

        vm.selectFork(bscForkID);
        vm.assume(amount < cross.balanceOf(OWNER));

        vm.prank(OWNER);
        cross.approve(address(bridgeBSCV2), amount);
        burnCrossToDeadWallet(false, OWNER, deadWallet, amount, false);
    }
}
