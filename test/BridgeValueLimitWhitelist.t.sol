// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {console} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

/**
 * @title BridgeValueLimitWhitelistE2ETest
 * @notice E2E tests (T12-T14) verifying the `BaseBridge` / `CrossBridge` wiring of the
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
}
