// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IPriceFeed} from "../src/PriceFeed.sol";
import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";

/**
 * @title BridgeTokenPriceUnavailableTest
 * @notice `TokenPriceUnavailable` fail-closed behaviour, and revalidation of
 * not-yet-validated pending records on release.
 * @dev Uses the shared BSC/Cross fork fixture (`BridgeTest`), overriding
 * `bridgeVerifierCross`'s thresholds and `priceFeedCross`'s price data per test the same
 * way `BridgeTokenMonitoring.t.sol` and `BridgeValueLimitWhitelist.t.sol` do.
 */
contract BridgeTokenPriceUnavailableTest is BridgeTest {
    function setUp() public override {
        super.setUp();

        vm.selectFork(bscForkID);
        vm.deal(USER, 10_000 ether);
        vm.prank(OWNER);
        cross.transfer(USER, 10_000 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);

        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), type(uint).max);
    }

    /**
     * @dev The real `PriceFeed` contract's `updatePrice` rejects a zero price outright
     * (`PriceFeedCanNotZeroValue`), so a genuinely malicious/misconfigured `IPriceFeed`
     * that reports `exist == true, price == 0` can't be produced through it. Simulate
     * that implementation directly instead, the same way `BridgeTokenMonitoring.t.sol`
     * simulates a dead price feed with `vm.mockCallRevert`.
     */
    function _zeroPrice(address token) internal {
        vm.selectFork(crossForkID);
        vm.mockCall(
            address(priceFeedCross),
            abi.encodeWithSelector(IPriceFeed.getTokenPriceInDollars.selector, token),
            abi.encode(true, uint(0), uint(0))
        );
    }

    // ============================================================
    // TokenPriceUnavailable fail-closed
    // ============================================================

    /**
     * @notice T7: With monitoring active, a token whose price feed explicitly reports
     * `exist == true, price == 0` finalizes to a `TokenPriceUnavailable` pending instead
     * of silently passing every amount check.
     */
    function test_finalize_tokenPriceUnavailable_whenMonitoringActive() public {
        _zeroPrice(address(weth));

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(1); // any nonzero value activates monitoring

        uint amount = 5 ether;
        uint index = depositETH(true, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPriceUnavailable,
            "zero-priced token must fail closed while monitoring is active"
        );
    }

    /**
     * @notice T21: The value-limit whitelist check runs BEFORE the unit-price check
     * (fixed ordering decision, see BridgeVerifier.validateBridgeTokenValue's NatSpec) -
     * a whitelisted recipient bypasses fail-closed too, even with monitoring active and
     * the token's price unavailable.
     */
    function test_finalize_whitelistedRecipient_bypassesTokenPriceUnavailable() public {
        _zeroPrice(address(weth));

        address wl = makeAddr("wl_price_unavailable");
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(1); // activate monitoring
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        vm.stopPrank();

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(NATIVE_TOKEN), USER, wl, 5 ether, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint balBefore = weth.balanceOf(wl);
        crossFinalize(index, address(weth), wl, 5 ether, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None,
            "whitelisted recipient must not be parked, even with price unavailable"
        );
        assertEq(weth.balanceOf(wl), balBefore + 5 ether);
    }

    /**
     * @notice T8: Regression pin for R4/AC-5 - under the CURRENT on-chain config
     * (both thresholds 0, monitoring inactive), a zero-priced token still finalizes
     * normally. The `monitoringActive` gate must not fire when monitoring is off.
     */
    function test_finalize_zeroPricedToken_succeeds_whenMonitoringInactive() public {
        _zeroPrice(address(weth));

        // bridgeVerifierCross's thresholds are 0/0 by default in this fixture (mirrors
        // the live on-chain config) - no explicit setter calls here on purpose.
        uint amount = 5 ether;
        depositETH(false, amount, threshold); // asserts full success balance deltas itself
    }

    /**
     * @notice T9: A normally-priced token whose small transfer amount floors its score
     * to 0 must NOT be confused with `TokenPriceUnavailable` (which is keyed off the
     * unit price being zero, not the scaled score).
     */
    function test_finalize_smallAmount_floorsScoreToZero_stillSucceeds() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(1); // monitoring active, weth price left untouched (nonzero)

        uint amount = 1; // 1 wei: unitPrice * 1 / 1e18 floors to a score of 0
        depositETH(false, amount, threshold); // must succeed immediately, not fail-closed
    }

    // ============================================================
    // Revalidation of not-yet-validated pending records
    // ============================================================

    /**
     * @notice T10: A record parked as `TokenPaused` (validation never ran) must be
     * revalidated on `releasePending`. Once unpaused, a transfer that would have
     * exceeded the verification threshold now reverts instead of silently releasing.
     */
    function test_releasePending_revalidates_tokenPaused_andReverts_whenExceedsThreshold() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(5 * 1e6); // $5
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, true); // finalize paused
        vm.stopPrank();

        uint amount = 60 ether; // NATIVE_TOKEN price = $0.1 -> score = $6 > $5 threshold
        uint index = deposit(true, amount, threshold);

        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPaused);

        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);

        vm.expectRevert();
        bridgeCross.releasePending(BSC_CHAIN_ID, index);

        // Record stays pending, unreleased.
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPaused);
    }

    /**
     * @notice T11: The same `TokenPaused` scenario but below the verification threshold:
     * once unpaused, `releasePending` succeeds - and the transfer's score is counted
     * exactly once (it was never recorded while parked, since validation never ran).
     */
    function test_releasePending_revalidates_tokenPaused_andSucceeds_whenBelowThreshold() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(5 * 1e6); // $5
        bridgeVerifierCross.setPeriodTotalValueThreshold(100 * 1e6); // large headroom, just to exercise accounting
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, true);
        vm.stopPrank();

        uint amount = 30 ether; // score = $3 < $5 threshold
        uint index = deposit(true, amount, threshold);

        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPaused);
        assertEq(
            bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN),
            0,
            "score must not be recorded while parked - validation never ran"
        );

        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);

        uint balBefore = USER.balance;
        bridgeCross.releasePending(BSC_CHAIN_ID, index);

        assertEq(USER.balance, balBefore + amount, "release must deliver the pending amount");
        assertTrue(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None);
        assertEq(
            bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN),
            3 * 1e6,
            "score must be counted exactly once, by the revalidation on release"
        );
    }

    /**
     * @notice T12: A record whose status came FROM validation having already run and
     * made a decision (`VerificationAmountThresholdExceeded`) is NOT revalidated on
     * release: it stays blocked until the 24h verification delay expires, then releases
     * normally - the existing delay-then-release behavior is unchanged.
     */
    function test_releasePending_doesNotRevalidate_verificationThresholdExceeded() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(5 * 1e6); // $5, token NOT paused this time

        uint amount = 60 ether; // score = $6 > $5 threshold -> validation runs and decides
        uint index = deposit(true, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status
                == Const.FinalizeStatus.VerificationAmountThresholdExceeded
        );

        // Before the delay expires, release must still revert (existing behavior).
        vm.expectRevert();
        bridgeCross.releasePending(BSC_CHAIN_ID, index);

        vm.warp(block.timestamp + 24 hours + 1);

        uint balBefore = USER.balance;
        bridgeCross.releasePending(BSC_CHAIN_ID, index);
        assertEq(USER.balance, balBefore + amount, "release must succeed once the 24h delay has passed");
    }

    /**
     * @notice T20: Fail-closed is not bypassed by delay expiry. A `TokenPriceUnavailable`
     * pending stays blocked past its 24h delay window as long as the price is still
     * unavailable, and only releases once the price is restored.
     */
    function test_releasePending_tokenPriceUnavailable_notBypassedByDelayExpiry() public {
        _zeroPrice(address(weth));

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        // Activates monitoring, but comfortably above the ~$50 score this transfer
        // scores once the price is restored below - this test isn't about the
        // single-transfer threshold, just about TokenPriceUnavailable itself.
        bridgeVerifierCross.setVerificationAmountThreshold(1000 * 1e6);

        uint amount = 5 ether;
        uint index = depositETH(true, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPriceUnavailable
        );

        vm.warp(block.timestamp + 24 hours + 1);

        // Price is still unavailable: release must still revert, not auto-succeed.
        vm.expectRevert();
        bridgeCross.releasePending(BSC_CHAIN_ID, index);
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPriceUnavailable,
            "fail-closed pending must survive delay expiry"
        );

        // Restore the price by clearing the mock: the feed's real, never-actually-changed
        // price for weth (set in the base fixture) takes over again, and release now
        // succeeds.
        vm.clearMockedCalls();

        uint balBefore = weth.balanceOf(USER);
        bridgeCross.releasePending(BSC_CHAIN_ID, index);
        assertEq(weth.balanceOf(USER), balBefore + amount);
    }

    /**
     * @notice T22: A record whose status came from validation deciding
     * `PeriodTotalValueThresholdExceeded` (the accumulator WAS touched) is likewise not
     * revalidated on release - proving the "already validated" bucket isn't limited to
     * the single-transfer threshold, and that release doesn't double-count its score.
     */
    function test_releasePending_doesNotRevalidate_periodTotalThresholdExceeded() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(0); // isolate the period-total path
        bridgeVerifierCross.setPeriodTotalValueThreshold(5 * 1e6); // $5
        vm.stopPrank();

        uint amount = 60 ether; // score = $6 > $5 period threshold, and IS recorded
        uint index = deposit(true, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status
                == Const.FinalizeStatus.PeriodTotalValueThresholdExceeded
        );
        uint scoreAfterFinalize = bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN);
        assertEq(scoreAfterFinalize, 6 * 1e6, "the triggering transfer's own score IS recorded");

        vm.warp(block.timestamp + 24 hours + 1);

        uint balBefore = USER.balance;
        bridgeCross.releasePending(BSC_CHAIN_ID, index);
        assertEq(USER.balance, balBefore + amount);

        assertEq(
            bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN),
            scoreAfterFinalize,
            "release must not double-count the score via a second validation run"
        );
    }
}
