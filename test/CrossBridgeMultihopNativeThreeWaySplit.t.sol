// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Const} from "../src/lib/Const.sol";

import {CrossBridgeMultihopTest} from "./CrossBridgeMultihop.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title CrossBridgeMultihopNativeThreeWaySplitTest
 * @notice Positive regression for the forwarded entrypoint's strict native equality
 * check: with BOTH a real network fee and a real exchange fee configured for the hop-2
 * leg (so `value_ + currentNetworkFee + exFee_ == ctxValue` is a genuine three-way
 * split, not one term collapsing to zero), the full A -> B -> C native route still
 * succeeds end to end, settles the forward with no remainder, and leaves the forward
 * context clean for the next item.
 */
contract CrossBridgeMultihopNativeThreeWaySplitTest is CrossBridgeMultihopTest {
    function test_e2e_native_ABC_threeWaySplit_networkFeeAndExFeeBothNonzero() public {
        vm.selectFork(crossForkID);

        // Give the CROSS -> chain C leg a real, priced network fee (unset in the base
        // fixture, where that lane's gas price defaults to 0 and so always prices to a
        // zero network fee) on top of the already-nonzero default exchange-fee rate, so
        // the settlement below genuinely exercises all three terms.
        vm.startPrank(CrossOWNER);
        uint[] memory chainIDs = new uint[](1);
        uint[] memory prices = new uint[](1);
        uint[] memory pricesAt = new uint[](1);
        chainIDs[0] = CHAIN_C_ID;
        prices[0] = 100_000;
        pricesAt[0] = 0;
        priceFeedCross.updateNativeTokenPrice(chainIDs, prices, pricesAt);
        bridgeVerifierCross.updateGasPrice(CHAIN_C_ID, 1 gwei);
        vm.stopPrank();

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(fee2 > 0 && ex2 > 0, "test requires a genuine three-way split");
        uint ctxValue = value2 + fee2 + ex2;

        bytes memory hop3ExtraData = _hop3TargetExtraData(value2);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop3ExtraData);

        uint hop2Index = bridgeCross.getNextInitiateIndex(CHAIN_C_ID);
        uint hop1Index = nextIndexBSC;
        uint depositedCBefore = bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited;
        uint targetBalBefore = address(mockTargetChainC).balance;

        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "hop-1's forward must succeed on chain B (no fallback payout)");
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, hop1Index).status == Const.FinalizeStatus.None,
            "remaining == 0: the forwarded finalize item must not be pending"
        );
        assertEq(
            bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited,
            depositedCBefore + value2,
            "chain B's ledger for the CHAIN_C_ID pair must record exactly the forwarded principal, split three ways under the hood"
        );

        assertTrue(
            _signAndFinalizeChainC(CROSS_CHAIN_ID, hop2Index, address(NATIVE_TOKEN), USER, value2, hop3ExtraData, 5)
        );
        assertEq(
            address(mockTargetChainC).balance,
            targetBalBefore + value2,
            "chain C's real target contract must have received and consumed the forwarded value"
        );

        // Forward context released cleanly: an ordinary (non-forward) deposit right
        // after is unaffected by the just-completed forward.
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 1 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);
        deposit(false, 1 ether, threshold);
    }
}
