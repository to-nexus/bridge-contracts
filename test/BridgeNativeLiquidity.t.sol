// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";

/// @notice A recipient that always rejects incoming native value, used to exercise the
/// unchanged revert-on-failure behavior of the `_dev` fee payout call site.
contract RevertingDev {
    receive() external payable {
        revert("RevertingDev: no thanks");
    }
}

/**
 * @title BridgeNativeLiquidityTest
 * @notice A bridge short on native liquidity must park an `InsufficientLiquidity` pending
 * instead of reverting the whole finalize batch item, while the unrelated `_dev` fee
 * payout call site on initiate keeps reverting as before.
 */
contract BridgeNativeLiquidityTest is BridgeTest {
    function setUp() public override {
        super.setUp();

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 10_000 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);

        vm.selectFork(crossForkID);
        vm.deal(USER, 10_000 ether);
    }

    /**
     * @notice T23: A finalize that would pay out more native CROSS than the bridge
     * currently holds does not revert - it parks an `InsufficientLiquidity` pending, and
     * the finalize index still advances (a following finalize is not blocked).
     */
    function test_finalize_insufficientNativeLiquidity_parksPending_indexAdvances() public {
        vm.selectFork(crossForkID);
        vm.deal(address(bridgeCross), 1 ether); // force the bridge's native balance low

        uint amount = 5 ether; // exceeds the 1 ether the bridge now holds
        uint index = deposit(true, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InsufficientLiquidity,
            "short liquidity must park a pending record, not revert the batch item"
        );

        // The finalize index still advanced - prove it by successfully finalizing the
        // NEXT index for an amount the still-low balance can actually cover.
        uint index2 = deposit(false, 0.5 ether, threshold);
        assertEq(index2, index + 1, "the next finalize index must not be stuck behind the pending one");
    }

    /**
     * @notice T24: Once the bridge is topped up, `releasePending` succeeds immediately -
     * `delayExpiration == 0` because the amount/threshold check had already succeeded
     * before liquidity was found to be short, so there is no 24h wait.
     */
    function test_releasePending_succeedsImmediately_onceLiquidityRestored() public {
        vm.selectFork(crossForkID);
        vm.deal(address(bridgeCross), 1 ether);

        uint amount = 5 ether;
        uint index = deposit(true, amount, threshold);
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InsufficientLiquidity);
        assertEq(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).delayExpiration, 0, "no delay for a settlement-stage failure");

        // Top up the bridge's native balance.
        vm.deal(address(bridgeCross), 100 ether);

        uint balBefore = USER.balance;
        bridgeCross.releasePending(BSC_CHAIN_ID, index); // no warp needed: delayExpiration is 0
        assertEq(USER.balance, balBefore + amount);
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None);
    }

    /**
     * @notice T25: While liquidity is still short, `releasePending` reverts - and the
     * pending record is NOT lost. `_releasePending` removes it, then re-attempts
     * settlement and reverts on failure, rolling the removal back with it.
     */
    function test_releasePending_stillReverts_whenLiquidityStillShort_pendingSurvives() public {
        vm.selectFork(crossForkID);
        vm.deal(address(bridgeCross), 1 ether);

        uint amount = 5 ether;
        uint index = deposit(true, amount, threshold);
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InsufficientLiquidity);

        // Still short on liquidity: release must revert.
        vm.expectRevert();
        bridgeCross.releasePending(BSC_CHAIN_ID, index);

        // The pending record must still be there, unchanged, ready for a later retry.
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InsufficientLiquidity,
            "a failed release must not remove the pending record (rollback)"
        );
        assertEq(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).args.value, amount);
    }

    /**
     * @notice The `_dev` fee payout on `_initiateBridge` still reverts the whole initiate
     * when the call fails. Only the finalize-side native payout degrades to a pending
     * record; skipping a fee silently is not an acceptable outcome here.
     */
    function test_initiateBridge_devFeePayment_stillReverts_onFailedCall() public {
        vm.selectFork(crossForkID);
        RevertingDev evilDev = new RevertingDev();
        vm.prank(CrossOWNER);
        bridgeCross.setDev(payable(address(evilDev)));

        (uint value, uint gas, uint ex) = crossCalcFee(NATIVE_TOKEN, 10 ether);
        require(gas + ex > 0, "test requires a nonzero fee to exercise the _dev callback");

        uint bridgeBalanceBefore = address(bridgeCross).balance;

        bridgeRevertCross = true;
        (, bool ok) = crossBridge(address(NATIVE_TOKEN), USER, USER, value, gas, ex);

        assertFalse(ok, "a failed _dev fee payment must still revert the whole initiate");
        assertEq(address(bridgeCross).balance, bridgeBalanceBefore, "a reverted initiate must not move any funds");
    }
}
