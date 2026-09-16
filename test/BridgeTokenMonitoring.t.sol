// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier, IBridgeVerifier} from "../src/BridgeVerifier.sol";
import {IPriceFeed, PriceFeed} from "../src/PriceFeed.sol";

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {console} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
/**
 * @title BridgeVerifierTokenValueTest
 * @notice Test contract for the validateBridgeTokenValue function of BridgeVerifier
 * @dev Tests various scenarios for token value validation including thresholds and time windows
 */

contract BridgeVerifierTokenValueTest is BridgeTest {
    // Test variables
    uint private constant HIGH_TOKEN_PRICE = 10000 * (10 ** 6); // High token price (10000 USD)
    uint private constant TEST_VERIFICATION_AMOUNT_THRESHOLD = 100_000 * 1e6; // Verification threshold (set lower than token price)
    uint private constant TEST_PERIOD_TOTAL_VALUE_THRESHOLD = 500_000 * 1e6; // Period total value threshold
    uint private constant TEST_TIME_WINDOW = 1 hours; // Time window

    function setUp() public override {
        super.setUp();

        // Utilize components already deployed in CrossChain
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeCross.grantRole(PRICER_ROLE, VALIDATOR1); // for test

        // Modify existing bridgeVerifierCross settings
        bridgeVerifierCross.setVerificationAmountThreshold(TEST_VERIFICATION_AMOUNT_THRESHOLD);
        bridgeVerifierCross.setPeriodTotalValueThreshold(TEST_PERIOD_TOTAL_VALUE_THRESHOLD);
        bridgeVerifierCross.setTimeWindow(TEST_TIME_WINDOW);

        // Grant Const.BRIDGE_ROLE to CrossOWNER for testing
        bridgeVerifierCross.grantRole(Const.BRIDGE_ROLE, CrossOWNER);

        // Set token price very high
        address[] memory tokens = new address[](2);
        uint[] memory prices = new uint[](2);
        uint[] memory pricesAt = new uint[](2);

        tokens[0] = address(weth);
        prices[0] = HIGH_TOKEN_PRICE;
        pricesAt[0] = 0;

        tokens[1] = address(testTokenCross);
        prices[1] = HIGH_TOKEN_PRICE;
        pricesAt[1] = 0;

        // Cannot call prank again while already in prank mode, so stop and restart
        vm.stopPrank();
        vm.prank(VALIDATOR1);
        priceFeedCross.updatePrice(tokens, prices, pricesAt);
    }

    /**
     * @notice Test basic validation with amount below thresholds
     * @dev Verifies that validation passes for normal amounts
     */
    function test_validateBridgeTokenValue_below_threshold() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        uint smallAmount = 1 ether;
        (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(IERC20(address(weth)), smallAmount);

        assertTrue(status == Const.FinalizeStatus.Success);

        vm.stopPrank();
    }

    /**
     * @notice Test validation failure when amount exceeds verification threshold
     * @dev Verifies that validation fails when a single transfer exceeds the set threshold
     */
    function test_validateBridgeTokenValue_exceed_verification_threshold() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        uint score = bridgeVerifierCross.getTokenCurrentScore(testTokenCross);

        // Perform the rest of the test
        (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether);

        assertTrue(status == Const.FinalizeStatus.VerificationAmountThresholdExceeded);
        assertTrue(score == bridgeVerifierCross.getTokenCurrentScore(testTokenCross));

        vm.stopPrank();
    }

    /**
     * @notice Test validation failure when total volume exceeds period threshold
     * @dev Verifies that validation fails when multiple transfers exceed the period total value threshold
     */
    function test_validateBridgeTokenValue_exceed_period_total_threshold() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        // Make multiple transfers to exceed the period total value threshold
        for (uint i = 0; i < 5; i++) {
            (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether);
            assertTrue(status == Const.FinalizeStatus.Success);
        }

        // The last transfer should fail as it exceeds the period total value threshold
        (Const.FinalizeStatus s) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether);

        assertTrue(s == Const.FinalizeStatus.PeriodTotalValueThresholdExceeded);

        vm.stopPrank();
    }

    /**
     * @notice Test validation after time window has passed
     * @dev Verifies that old transfers are removed from the time window tracking
     */
    function test_validateBridgeTokenValue_time_window_expiration() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        // Accumulate transfers just below the period threshold
        uint transferAmount = 10 ether;
        uint numTransfers = 4; // Total 50 ether, below threshold

        for (uint i = 0; i < numTransfers; i++) {
            (Const.FinalizeStatus status) =
                bridgeVerifierCross.validateBridgeTokenValue(IERC20(address(weth)), transferAmount);
            assertTrue(status == Const.FinalizeStatus.Success);
        }

        // Move time forward past the time window
        vm.warp(block.timestamp + TEST_TIME_WINDOW + 1 hours);

        // This transfer should pass since previous transfers are now outside the time window
        (Const.FinalizeStatus s) = bridgeVerifierCross.validateBridgeTokenValue(IERC20(address(weth)), transferAmount);

        assertTrue(s == Const.FinalizeStatus.Success);

        vm.stopPrank();
    }

    /**
     * @notice Test zero token value validation
     * @dev Verifies behavior when token value is zero
     */
    function test_validateBridgeTokenValue_zero_value() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(IERC20(address(weth)), 0);

        assertTrue(status == Const.FinalizeStatus.Success);

        vm.stopPrank();
    }

    /**
     * @notice Test validation with different tokens
     * @dev Verifies that validation properly tracks thresholds separately for different tokens
     */
    function test_validateBridgeTokenValue_multiple_tokens() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        // Make multiple transfers to exceed the period total value threshold
        for (uint i = 0; i < 5; i++) {
            (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether);
            assertTrue(status == Const.FinalizeStatus.Success);
        }

        // The last transfer should fail as it exceeds the period total value threshold
        (Const.FinalizeStatus s) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether);

        assertTrue(s == Const.FinalizeStatus.PeriodTotalValueThresholdExceeded);

        vm.stopPrank();
    }

    // ============================================================
    // Value-limit whitelist (T1-T11, T15-T16)
    // ============================================================

    /**
     * @notice T1: A whitelisted recipient bypasses the single-transfer verification
     * threshold, and doing so leaves the token's accumulated score untouched.
     */
    function test_validateBridgeTokenValue_whitelist_bypasses_verification_threshold() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_verification");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        uint scoreBefore = bridgeVerifierCross.getTokenCurrentScore(testTokenCross);

        Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);

        assertTrue(status == Const.FinalizeStatus.Success);
        assertEq(bridgeVerifierCross.getTokenCurrentScore(testTokenCross), scoreBefore);

        vm.stopPrank();
    }

    /**
     * @notice T2: Repeated whitelisted transfers are never recorded in the period
     * accumulator or movement history.
     */
    function test_validateBridgeTokenValue_whitelist_not_accumulated() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_not_accumulated");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        for (uint i = 0; i < 5; i++) {
            Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
            assertTrue(status == Const.FinalizeStatus.Success);
        }

        assertEq(bridgeVerifierCross.getTokenCurrentScore(testTokenCross), 0);
        assertEq(bridgeVerifierCross.getTokenMovementHistory(testTokenCross).length, 0);

        vm.stopPrank();
    }

    /**
     * @notice T3: Whitelisted transfers interleaved with non-whitelisted ones do not
     * affect the non-whitelisted transfers' own accumulation.
     */
    function test_validateBridgeTokenValue_whitelist_does_not_affect_others() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_others");
        address other = makeAddr("other_recipient");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, other);
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, other);
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, other);

        // Only the 3 non-whitelisted 10-ether transfers should have accumulated:
        // 3 * 100_000e6 = 300_000e6.
        assertEq(bridgeVerifierCross.getTokenCurrentScore(testTokenCross), 300_000 * 1e6);

        vm.stopPrank();
    }

    /**
     * @notice T4: A whitelisted recipient bypasses the period-total threshold even when
     * the accumulator is already full.
     */
    function test_validateBridgeTokenValue_whitelist_bypasses_period_threshold() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_period");
        address other = makeAddr("other_period");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        // Fill the period accumulator to exactly the threshold (5 * 10 ether = 500_000e6).
        for (uint i = 0; i < 5; i++) {
            Const.FinalizeStatus fillStatus =
                bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, other);
            assertTrue(fillStatus == Const.FinalizeStatus.Success);
        }

        // A further non-whitelisted transfer now exceeds the period threshold.
        Const.FinalizeStatus blocked = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, other);
        assertTrue(blocked == Const.FinalizeStatus.PeriodTotalValueThresholdExceeded);

        // The whitelisted recipient still succeeds despite the full period accumulator.
        Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 10 ether, wl);
        assertTrue(status == Const.FinalizeStatus.Success);

        vm.stopPrank();
    }

    /**
     * @notice T5: Once removed from the whitelist, the same recipient is subject to the
     * full thresholds again.
     */
    function test_validateBridgeTokenValue_whitelist_removed_reapplies_limits() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_removed");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        Const.FinalizeStatus bypassed = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
        assertTrue(bypassed == Const.FinalizeStatus.Success);

        bridgeVerifierCross.removeValueLimitWhitelist(wl);

        Const.FinalizeStatus limited = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
        assertTrue(limited == Const.FinalizeStatus.VerificationAmountThresholdExceeded);

        vm.stopPrank();
    }

    /**
     * @notice T6: The legacy 2-arg overload always evaluates as non-whitelisted
     * (fail-safe), regardless of whether some address is whitelisted.
     */
    function test_validateBridgeTokenValue_legacy_two_arg_is_not_whitelisted() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_legacy");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        // The legacy overload has no recipient parameter, so it always internally uses
        // to = address(0), which is never whitelisted.
        Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether);
        assertTrue(status == Const.FinalizeStatus.VerificationAmountThresholdExceeded);

        vm.stopPrank();
    }

    /**
     * @notice T7: The zero address can never be whitelisted.
     */
    function test_validateBridgeTokenValue_zero_address_never_whitelisted() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierCanNotZeroValue.selector, "account"));
        bridgeVerifierCross.addValueLimitWhitelist(address(0));

        vm.stopPrank();
    }

    /**
     * @notice T8: A whitelisted transfer succeeds even when the price feed is dead,
     * because the bypass branch returns before any price lookup.
     */
    function test_validateBridgeTokenValue_whitelist_works_without_price_feed() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_no_price_feed");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);
        vm.stopPrank();

        // Simulate a dead/unavailable price feed: any call reaching the price lookup reverts.
        vm.mockCallRevert(
            address(priceFeedCross),
            abi.encodeWithSelector(IPriceFeed.getTokenPriceInDollars.selector, address(testTokenCross)),
            "price feed down"
        );

        vm.startPrank(CrossOWNER);

        // The whitelisted call succeeds without ever touching the price feed.
        Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);
        assertTrue(status == Const.FinalizeStatus.Success);

        // Sanity check: a non-whitelisted call under the same broken price feed reverts,
        // proving the bypass really is what avoids the price feed dependency above.
        vm.expectRevert("price feed down");
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, address(0));

        vm.stopPrank();
    }

    /**
     * @notice T9: Only ADMIN_ROLE may add or remove whitelist entries.
     */
    function test_validateBridgeTokenValue_whitelist_admin_only() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_admin_only");
        address nonAdmin = makeAddr("non_admin");

        vm.prank(nonAdmin);
        vm.expectRevert();
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.prank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.prank(nonAdmin);
        vm.expectRevert();
        bridgeVerifierCross.removeValueLimitWhitelist(wl);
    }

    /**
     * @notice T10: Duplicate add and missing remove each revert with a dedicated error
     * (no silent no-op).
     */
    function test_validateBridgeTokenValue_whitelist_duplicate_and_missing() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_dup");
        address notRegistered = makeAddr("wl_missing");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierWhitelistAlreadyAdded.selector, wl));
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.expectRevert(
            abi.encodeWithSelector(BridgeVerifier.BridgeVerifierWhitelistNotFound.selector, notRegistered)
        );
        bridgeVerifierCross.removeValueLimitWhitelist(notRegistered);

        vm.stopPrank();
    }

    /**
     * @notice T11: The full-list getter and length reflect add/remove operations.
     */
    function test_validateBridgeTokenValue_whitelist_enumeration() public {
        vm.selectFork(crossForkID);
        address wl1 = makeAddr("wl_enum_1");
        address wl2 = makeAddr("wl_enum_2");

        vm.startPrank(CrossOWNER);
        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 0);

        bridgeVerifierCross.addValueLimitWhitelist(wl1);
        bridgeVerifierCross.addValueLimitWhitelist(wl2);
        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 2);

        address[] memory all = bridgeVerifierCross.getValueLimitWhitelist();
        assertEq(all.length, 2);
        assertTrue(_contains(all, wl1) && _contains(all, wl2));

        bridgeVerifierCross.removeValueLimitWhitelist(wl1);
        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 1);
        all = bridgeVerifierCross.getValueLimitWhitelist();
        assertEq(all.length, 1);
        assertEq(all[0], wl2);

        vm.stopPrank();
    }

    /**
     * @notice T15: Batch add/remove are atomic - a duplicate/missing entry anywhere in
     * the batch reverts the whole transaction, with no partial application.
     */
    function test_validateBridgeTokenValue_whitelist_batch_is_atomic() public {
        vm.selectFork(crossForkID);
        address a1 = makeAddr("wl_batch_1");
        address a2 = makeAddr("wl_batch_2");
        address a3 = makeAddr("wl_batch_3");

        vm.startPrank(CrossOWNER);

        // a2 is already registered, so an add-batch containing it must revert entirely.
        bridgeVerifierCross.addValueLimitWhitelist(a2);
        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 1);

        address[] memory addBatch = new address[](3);
        addBatch[0] = a1;
        addBatch[1] = a2;
        addBatch[2] = a3;

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierWhitelistAlreadyAdded.selector, a2));
        bridgeVerifierCross.addValueLimitWhitelistBatch(addBatch);

        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 1);
        assertFalse(bridgeVerifierCross.isValueLimitWhitelisted(a1));
        assertFalse(bridgeVerifierCross.isValueLimitWhitelisted(a3));

        // a4 was never registered, so a remove-batch containing it must also revert entirely.
        address a4 = makeAddr("wl_batch_4");
        address[] memory removeBatch = new address[](2);
        removeBatch[0] = a2;
        removeBatch[1] = a4;

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierWhitelistNotFound.selector, a4));
        bridgeVerifierCross.removeValueLimitWhitelistBatch(removeBatch);

        assertEq(bridgeVerifierCross.getValueLimitWhitelistLength(), 1);
        assertTrue(bridgeVerifierCross.isValueLimitWhitelisted(a2));

        vm.stopPrank();
    }

    /**
     * @notice T16: Pagination boundary conditions - normal window, partial tail window,
     * offset past the end (empty, not a revert), and limit == 0 (empty). Concatenating
     * pages must reproduce the full-list getter's result.
     */
    function test_validateBridgeTokenValue_whitelist_pagination() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        address[] memory accounts = new address[](5);
        for (uint i = 0; i < 5; i++) {
            accounts[i] = makeAddr(string(abi.encodePacked("wl_page_", vm.toString(i))));
        }
        bridgeVerifierCross.addValueLimitWhitelistBatch(accounts);

        address[] memory full = bridgeVerifierCross.getValueLimitWhitelist();
        assertEq(full.length, 5);

        // Normal window in the middle.
        address[] memory mid = bridgeVerifierCross.getValueLimitWhitelist(1, 2);
        assertEq(mid.length, 2);
        assertEq(mid[0], full[1]);
        assertEq(mid[1], full[2]);

        // offset + limit > length: only the remaining entries are returned.
        address[] memory tail = bridgeVerifierCross.getValueLimitWhitelist(3, 10);
        assertEq(tail.length, 2);
        assertEq(tail[0], full[3]);
        assertEq(tail[1], full[4]);

        // offset >= length: empty array, not a revert.
        address[] memory emptyAtLength = bridgeVerifierCross.getValueLimitWhitelist(5, 10);
        assertEq(emptyAtLength.length, 0);
        address[] memory emptyPastLength = bridgeVerifierCross.getValueLimitWhitelist(100, 10);
        assertEq(emptyPastLength.length, 0);

        // limit == 0: empty array.
        address[] memory emptyLimit = bridgeVerifierCross.getValueLimitWhitelist(0, 0);
        assertEq(emptyLimit.length, 0);

        // Concatenating pages reproduces the full-list getter's result.
        address[] memory page0 = bridgeVerifierCross.getValueLimitWhitelist(0, 3);
        address[] memory page1 = bridgeVerifierCross.getValueLimitWhitelist(3, 3);
        assertEq(page0.length, 3);
        assertEq(page1.length, 2);
        for (uint i = 0; i < 3; i++) {
            assertEq(page0[i], full[i]);
        }
        for (uint i = 0; i < 2; i++) {
            assertEq(page1[i], full[3 + i]);
        }

        vm.stopPrank();
    }

    // ============================================================
    // Event verification, positive and negative (T17-T20)
    // ============================================================

    /**
     * @notice T17: A successful bypass emits `ValueLimitWhitelistBypassed` exactly once
     * with the exact (token, to, value) arguments.
     */
    function test_validateBridgeTokenValue_bypass_event_emitted_with_exact_args() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_event");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.expectEmit(true, true, false, true);
        emit BridgeVerifier.ValueLimitWhitelistBypassed(testTokenCross, wl, 20 ether);
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether, wl);

        vm.stopPrank();
    }

    /**
     * @notice T18: A normal (non-whitelisted) successful call never emits
     * `ValueLimitWhitelistBypassed`.
     */
    function test_validateBridgeTokenValue_bypass_event_not_emitted_when_not_whitelisted() public {
        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        vm.recordLogs();
        Const.FinalizeStatus status = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 1 ether, address(0));
        assertTrue(status == Const.FinalizeStatus.Success);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        bytes32 bypassedTopic0 = keccak256("ValueLimitWhitelistBypassed(address,address,uint256)");
        for (uint i = 0; i < logs.length; i++) {
            assertTrue(logs[i].topics.length == 0 || logs[i].topics[0] != bypassedTopic0);
        }

        vm.stopPrank();
    }

    /**
     * @notice T19: The legacy 2-arg overload never emits `ValueLimitWhitelistBypassed`,
     * even when some address is whitelisted (pairs with T6).
     */
    function test_validateBridgeTokenValue_bypass_event_not_emitted_on_legacy_two_arg() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_legacy_event");

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.recordLogs();
        bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 1 ether);
        Vm.Log[] memory logs = vm.getRecordedLogs();

        bytes32 bypassedTopic0 = keccak256("ValueLimitWhitelistBypassed(address,address,uint256)");
        for (uint i = 0; i < logs.length; i++) {
            assertTrue(logs[i].topics.length == 0 || logs[i].topics[0] != bypassedTopic0);
        }

        vm.stopPrank();
    }

    /**
     * @notice T20: Single add/remove emit their respective events with exact arguments;
     * a successful batch add emits exactly one `Added` event per input address, in
     * input order.
     */
    function test_validateBridgeTokenValue_whitelist_admin_events() public {
        vm.selectFork(crossForkID);
        address wl = makeAddr("wl_admin_event");

        vm.startPrank(CrossOWNER);

        vm.expectEmit(true, false, false, true);
        emit BridgeVerifier.ValueLimitWhitelistAdded(wl);
        bridgeVerifierCross.addValueLimitWhitelist(wl);

        vm.expectEmit(true, false, false, true);
        emit BridgeVerifier.ValueLimitWhitelistRemoved(wl);
        bridgeVerifierCross.removeValueLimitWhitelist(wl);

        address b1 = makeAddr("wl_batch_evt_1");
        address b2 = makeAddr("wl_batch_evt_2");
        address b3 = makeAddr("wl_batch_evt_3");
        address[] memory batch = new address[](3);
        batch[0] = b1;
        batch[1] = b2;
        batch[2] = b3;

        vm.recordLogs();
        bridgeVerifierCross.addValueLimitWhitelistBatch(batch);
        Vm.Log[] memory logs = vm.getRecordedLogs();

        bytes32 addedTopic0 = keccak256("ValueLimitWhitelistAdded(address)");
        address[] memory addedInOrder = new address[](3);
        uint count;
        for (uint i = 0; i < logs.length; i++) {
            if (logs[i].topics.length > 0 && logs[i].topics[0] == addedTopic0) {
                addedInOrder[count] = address(uint160(uint(logs[i].topics[1])));
                count++;
            }
        }
        assertEq(count, 3);
        assertEq(addedInOrder[0], b1);
        assertEq(addedInOrder[1], b2);
        assertEq(addedInOrder[2], b3);

        // Batch removal: exactly one Removed event per input address, in input order.
        // `EnumerableSet.remove` uses swap-and-pop, so the set's internal enumeration
        // order changes as the loop runs - assert on the EMITTED event order (which
        // follows the input array), not on set enumeration order.
        vm.recordLogs();
        bridgeVerifierCross.removeValueLimitWhitelistBatch(batch);
        Vm.Log[] memory removeLogs = vm.getRecordedLogs();

        bytes32 removedTopic0 = keccak256("ValueLimitWhitelistRemoved(address)");
        address[] memory removedInOrder = new address[](3);
        uint removedCount;
        for (uint i = 0; i < removeLogs.length; i++) {
            if (removeLogs[i].topics.length > 0 && removeLogs[i].topics[0] == removedTopic0) {
                removedInOrder[removedCount] = address(uint160(uint(removeLogs[i].topics[1])));
                removedCount++;
            }
        }
        assertEq(removedCount, 3);
        assertEq(removedInOrder[0], b1);
        assertEq(removedInOrder[1], b2);
        assertEq(removedInOrder[2], b3);

        vm.stopPrank();
    }

    function _contains(address[] memory arr, address needle) private pure returns (bool) {
        for (uint i = 0; i < arr.length; i++) {
            if (arr[i] == needle) return true;
        }
        return false;
    }
}
