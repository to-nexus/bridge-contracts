// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier, IBridgeVerifier} from "../src/BridgeVerifier.sol";
import {IPriceFeed, PriceFeed} from "../src/PriceFeed.sol";

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {console} from "forge-std/Test.sol";
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

        // Perform the rest of the test
        (Const.FinalizeStatus status) = bridgeVerifierCross.validateBridgeTokenValue(testTokenCross, 20 ether);

        assertTrue(status == Const.FinalizeStatus.VerificationAmountThresholdExceeded);
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
}
