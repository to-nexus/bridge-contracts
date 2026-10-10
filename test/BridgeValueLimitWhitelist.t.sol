// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {console} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

/**
 * @title BridgeValueLimitWhitelistE2ETest
 * @notice E2E tests verifying the `BaseBridge` / `CrossBridge` wiring of the
 * `BridgeVerifier` value-limit whitelist through an actual finalize, reusing the same
 * fork harness as `BridgeCrossSupplyLimit.t.sol`.
 */
contract BridgeValueLimitWhitelistE2ETest is BridgeTest {
    function setUp() public virtual override {
        super.setUp();

        // Ensure USER has enough tokens for tests (mirrors BridgeCrossSupplyLimitTest.setUp()).
        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);
        cross.transfer(USER, 1000 ether);
        vm.stopPrank();

        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);

        // Make the single-transfer verification threshold trivially low so any deposit
        // in these tests would normally be delayed to pending.
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(1);
    }

    /**
     * @notice T12: A finalize to a whitelisted recipient is immediately
     * `BridgeFinalized` (not pending) even though it exceeds the single-transfer
     * threshold, and the verifier's accumulator is left untouched.
     */
    function test_finalize_whitelisted_recipient_e2e() public {
        address wl = makeAddr("wl_e2e_recipient");

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint recipientBalanceBefore = wl.balance;

        // Use vm.recordLogs() rather than vm.expectEmit(): the latter arms for the very
        // next external call, which is fragile once helpers make incidental staticcalls
        // first (see the comment above `_assertMislinkRevertsWholeBatch` in
        // `test/CrossBridgeForward.t.sol` for a documented instance of that trap).
        vm.recordLogs();
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold);
        Vm.Log[] memory logs = vm.getRecordedLogs();

        // The literal success signal off-chain consumers (scanner, validator) index:
        // `BridgeFinalized(fromChainID, index, toToken, to, value, timestamp)`.
        bytes32 bridgeFinalizedTopic0 = keccak256("BridgeFinalized(uint256,uint256,address,address,uint256,uint256)");
        bool foundBridgeFinalized;
        for (uint i = 0; i < logs.length; i++) {
            if (logs[i].topics.length == 4 && logs[i].topics[0] == bridgeFinalizedTopic0) {
                assertEq(logs[i].topics[1], bytes32(BSC_CHAIN_ID), "fromChainID mismatch");
                assertEq(logs[i].topics[2], bytes32(index), "index mismatch");
                assertEq(logs[i].topics[3], bytes32(uint(uint160(address(NATIVE_TOKEN)))), "toToken mismatch");

                (address decodedTo, uint decodedValue, uint decodedTime) =
                    abi.decode(logs[i].data, (address, uint, uint));
                assertEq(decodedTo, wl, "to mismatch");
                assertEq(decodedValue, amount, "value mismatch");
                assertEq(decodedTime, block.timestamp, "timestamp mismatch");

                foundBridgeFinalized = true;
            }
        }
        assertTrue(foundBridgeFinalized, "BridgeFinalized must be emitted for the whitelisted finalize");

        // Immediately finalized: no pending entry was ever created for this index.
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None,
            "whitelisted finalize must not be pending"
        );
        assertEq(wl.balance, recipientBalanceBefore + amount, "recipient must receive funds immediately");

        // The bypass must not have touched the period accumulator (RQ3).
        assertEq(bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN), 0);
    }

    /**
     * @notice T13: A whitelisted recipient still falls to pending with
     * `CrossSupplyLimitExceeded` when the transfer would exceed `crossSupplyLimit` -
     * the whitelist only exempts the verifier's own thresholds, not the cross supply cap.
     */
    function test_finalize_whitelisted_still_respects_crossSupplyLimit() public {
        address wl = makeAddr("wl_e2e_supply_limit");

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        uint supplyLimit = 10 ether;
        bridgeCross.setCrossSupplyLimit(supplyLimit + CROSS_FOUNDATION_INITIAL_SUPPLY);
        vm.stopPrank();

        uint amount = 20 ether; // exceeds crossSupplyLimit

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.CrossSupplyLimitExceeded,
            "crossSupplyLimit must still be enforced for whitelisted recipients"
        );
    }

    /**
     * @notice T14: A whitelisted recipient still falls to pending with `TokenPaused`
     * when finalize is paused for that token pair - the whitelist only exempts the
     * verifier's own thresholds, not the token-pause circuit breaker.
     */
    function test_finalize_whitelisted_still_respects_tokenPause() public {
        address wl = makeAddr("wl_e2e_token_pause");

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, true);
        vm.stopPrank();

        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPaused,
            "token finalize pause must still be enforced for whitelisted recipients"
        );
    }

    /**
     * @notice A whitelisted recipient does NOT bypass the single-transfer threshold
     * when the finalize item carries non-empty extraData: with extraData driving a
     * possible executor call, the signed `to` is not necessarily who ends up holding
     * the value, so the exemption is withheld and the item parks exactly like a
     * non-whitelisted one would (not a `TokenPaused`/`CrossSupplyLimitExceeded`-style
     * reason - this is the threshold itself firing).
     */
    function test_finalize_whitelisted_nonEmptyExtraData_overThreshold_parksInsteadOfBypassing() public {
        address wl = makeAddr("wl_nonempty_extradata");

        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        vm.recordLogs();
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold, bytes("nonempty"));
        Vm.Log[] memory logs = vm.getRecordedLogs();

        bytes32 bypassedTopic0 = keccak256("ValueLimitWhitelistBypassed(address,address,uint256)");
        for (uint i = 0; i < logs.length; i++) {
            assertFalse(logs[i].topics[0] == bypassedTopic0, "non-empty extraData must never bypass");
        }

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status
                == Const.FinalizeStatus.VerificationAmountThresholdExceeded,
            "whitelisted recipient must still hit the single-transfer threshold once extraData is non-empty"
        );
        assertEq(bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN), 0, "a parked item must not touch the accumulator");
    }

    /**
     * @notice With both thresholds actually active and set high enough that the
     * transfer clears them on its own merits, a whitelisted recipient with non-empty
     * extraData settles immediately (as any ordinary transfer would) - but, unlike the
     * empty-extraData case, through the verifier's normal accounting path: no bypass
     * event, and the period accumulator records the score exactly as it would for a
     * non-whitelisted recipient.
     */
    function test_finalize_whitelisted_nonEmptyExtraData_withinLimits_settlesWithoutBypass() public {
        address wl = makeAddr("wl_nonempty_extradata_within_limits");

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        // Raise both thresholds well above what this transfer can score, so a
        // non-exempt evaluation still succeeds - isolating "non-empty extraData loses
        // the bypass" from "the transfer happens to be too large".
        bridgeVerifierCross.setVerificationAmountThreshold(type(uint192).max);
        bridgeVerifierCross.setPeriodTotalValueThreshold(type(uint192).max);
        vm.stopPrank();

        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint scoreBefore = bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN);
        uint recipientBalanceBefore = wl.balance;

        vm.recordLogs();
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold, bytes("nonempty"));
        Vm.Log[] memory logs = vm.getRecordedLogs();

        bytes32 bypassedTopic0 = keccak256("ValueLimitWhitelistBypassed(address,address,uint256)");
        for (uint i = 0; i < logs.length; i++) {
            assertFalse(logs[i].topics[0] == bypassedTopic0, "non-empty extraData must never bypass");
        }

        assertEq(wl.balance, recipientBalanceBefore + amount, "the transfer must still settle on its own merits");
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.None,
            "a within-limits transfer must not be pending"
        );
        assertGt(
            bridgeVerifierCross.getTokenCurrentScore(NATIVE_TOKEN),
            scoreBefore,
            "a non-exempt settlement must record its score in the period accumulator"
        );
    }

    /**
     * @notice A record parked for a reason unrelated to the amount (here,
     * `TokenPaused`) always has its `extraData` cleared once pending (the bridge never
     * retains executor-routing data across a park/release cycle). So even though the
     * ORIGINAL finalize attempt carried non-empty extraData and therefore evaluated
     * without the whitelist exemption, `releasePending` re-validates against the
     * recipient alone and the whitelist applies there - this transfer would otherwise
     * still fail the single-transfer threshold on release.
     */
    function test_releasePending_whitelisted_originallyNonEmptyExtraData_exemptsOnRelease() public {
        address wl = makeAddr("wl_release_after_nonempty_extradata");

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        bridgeVerifierCross.setVerificationAmountThreshold(1); // trivially low: gates anyone not exempt
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, true);
        vm.stopPrank();

        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), wl, amount, threshold, bytes("nonempty"));
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.TokenPaused,
            "must park on the pause, independent of amount or extraData"
        );
        assertEq(bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).args.extraData.length, 0, "extraData must be cleared once pending");

        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);

        uint recipientBalanceBefore = wl.balance;
        bridgeCross.releasePending(BSC_CHAIN_ID, index);
        assertEq(wl.balance, recipientBalanceBefore + amount, "release must settle via the whitelist exemption");
    }
}
