// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {IPriceFeed} from "../src/PriceFeed.sol";
import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title BridgeMinimumTokenValueOverrideTest
 * @notice Covers the per-token minimum-transfer-value override on `BridgeVerifier`.
 * @dev The override mirrors `_exFeeRate`'s shape (per-token entry, `0` falls back to the
 * global) but deliberately has NO exemption sentinel: `0` means "unset" and nothing else.
 *
 * Fixture values inherited from `CrossChain.t.sol` (DOLLAR_DECIMALS = 6):
 *   _minimumTokenValue  = 10_000   ($0.01, global)
 *   _defaultTokenPrice  = 10_000   ($0.01)
 *   weth price          = 10e6     ($10.00)  -> global minimum 0.001 weth
 *   testTokenCross      = 1e6      ($1.00)   -> global minimum 0.01  TT
 *   NATIVE_TOKEN        = 100_000  ($0.10)   -> global minimum 0.1   CROSS
 */
contract BridgeMinimumTokenValueOverrideTest is BridgeTest {
    /// @dev One dollar expressed in the price feed's `dollarDecimals`.
    uint private constant ONE_DOLLAR = 10 ** DOLLAR_DECIMALS;

    event MinimumTokenValueOfSet(address indexed token, uint minimumTokenValue);

    function _minimumOf(IERC20 token) private view returns (uint minimumValue) {
        (minimumValue,,) = bridgeVerifierCross.getTokenConfig(BSC_CHAIN_ID, token);
    }

    /**
     * @notice An unset override leaves the pre-existing global behaviour untouched.
     * @dev This is the regression guard for the change: every already-registered token
     * must keep the exact minimum it had before the override existed.
     */
    function test_unset_override_falls_back_to_global() public {
        vm.selectFork(crossForkID);

        assertEq(bridgeVerifierCross.getMinimumTokenValueOf(weth), 0, "weth override should start unset");
        assertEq(bridgeVerifierCross.getMinimumTokenValueOf(testTokenCross), 0, "TT override should start unset");

        // 10_000 ($0.01) / price, scaled by the token's 18 decimals.
        assertEq(_minimumOf(weth), 0.001 ether, "weth minimum should follow the global value");
        assertEq(_minimumOf(testTokenCross), 0.01 ether, "TT minimum should follow the global value");
        assertEq(_minimumOf(NATIVE_TOKEN), 0.1 ether, "native minimum should follow the global value");
    }

    /// @notice An override moves only its own token; every other token keeps the global.
    function test_override_applies_only_to_that_token() public {
        vm.selectFork(crossForkID);

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 5 * ONE_DOLLAR); // $5.00

        // TT is $1.00, so a $5 minimum is 5 tokens.
        assertEq(_minimumOf(testTokenCross), 5 ether, "TT should use its override");
        // Untouched tokens keep the global $0.01.
        assertEq(_minimumOf(weth), 0.001 ether, "weth must not be affected");
        assertEq(_minimumOf(NATIVE_TOKEN), 0.1 ether, "native must not be affected");
    }

    /// @notice Writing `0` clears the override and restores the global value.
    function test_zero_clears_override_and_restores_global() public {
        vm.selectFork(crossForkID);

        vm.startPrank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 5 * ONE_DOLLAR);
        assertEq(_minimumOf(testTokenCross), 5 ether, "override should apply first");

        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 0);
        vm.stopPrank();

        assertEq(bridgeVerifierCross.getMinimumTokenValueOf(testTokenCross), 0, "override should be cleared");
        assertEq(_minimumOf(testTokenCross), 0.01 ether, "minimum should fall back to the global value");
    }

    /**
     * @notice The getter reports the stored value, not the resolved one.
     * @dev An operator has to be able to tell "unset" apart from "set to the same number
     * as the global", because the two behave differently the moment the global changes.
     */
    function test_getter_reports_raw_value_not_resolved() public {
        vm.selectFork(crossForkID);

        // Set the override to exactly the global value.
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 10_000);
        assertEq(bridgeVerifierCross.getMinimumTokenValueOf(testTokenCross), 10_000, "raw override should be readable");

        // Move the global; the pinned token must not follow it.
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValue(5 * ONE_DOLLAR);

        assertEq(_minimumOf(testTokenCross), 0.01 ether, "pinned token keeps its own value");
        assertEq(_minimumOf(weth), 0.5 ether, "unpinned token follows the new global");
    }

    /**
     * @notice The operational scenario this override was added for.
     * @dev CROSS repriced from $1.00 to $0.15 pushes the native minimum from 5 to 33.33
     * tokens under a single global $5. A $0.75 override brings it back to 5 CROSS while
     * a dollar-priced token keeps its own $5 minimum.
     */
    function test_scenario_native_minimum_pinned_to_five_tokens() public {
        vm.selectFork(crossForkID);

        // Reprice: native $0.15, TT $1.00 (already $1). Global minimum $5.
        address[] memory tokens = new address[](1);
        uint[] memory prices = new uint[](1);
        uint[] memory pricesAt = new uint[](1);
        tokens[0] = address(NATIVE_TOKEN);
        prices[0] = 150_000; // $0.15
        pricesAt[0] = 0;

        vm.startPrank(CrossOWNER);
        priceFeedCross.updatePrice(tokens, prices, pricesAt);
        bridgeVerifierCross.setMinimumTokenValue(5 * ONE_DOLLAR);
        vm.stopPrank();

        // Without an override a global $5 is 33.33... CROSS.
        assertEq(_minimumOf(NATIVE_TOKEN), uint(5 * ONE_DOLLAR) * 1 ether / 150_000, "native follows global $5");
        assertApproxEqAbs(_minimumOf(NATIVE_TOKEN), 33.333333 ether, 0.00001 ether, "global $5 is ~33.33 CROSS");

        // Pin native to $0.75 so it lands back on 5 tokens.
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(NATIVE_TOKEN, 750_000); // $0.75

        assertEq(_minimumOf(NATIVE_TOKEN), 5 ether, "native minimum should be exactly 5 CROSS");
        assertEq(_minimumOf(testTokenCross), 5 ether, "TT keeps the global $5, which is 5 TT at $1");
    }

    /**
     * @notice The override is still divided by `_defaultTokenPrice` when the feed has no
     * price for the token.
     * @dev `_defaultTokenPrice` cannot be zero (guarded in both the constructor and
     * `setDefaultTokenPrice`), so the `tokenPrice == 0` branch in `getTokenConfig` is not
     * reachable through any setter. The reachable fallback is the default price, and the
     * override rides on top of it: a $5 override against the $0.01 default is 500 tokens,
     * which is exactly the blow-up an operator needs to see rather than a silent $5.
     */
    function test_override_uses_default_price_when_feed_has_none() public {
        vm.selectFork(crossForkID);

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 5 * ONE_DOLLAR);

        // Price feed reports "no price for this token" -> _defaultTokenPrice ($0.01) applies.
        vm.mockCall(
            address(priceFeedCross),
            abi.encodeWithSelector(IPriceFeed.getTokenPriceInDollars.selector, address(testTokenCross)),
            abi.encode(false, uint(0), uint(0))
        );

        assertEq(_minimumOf(testTokenCross), 500 ether, "$5 override at the $0.01 default price is 500 tokens");
    }

    /**
     * @notice Pins `_minimumTokenValueOf` to storage slot 16.
     * @dev `script/verifier-settings-diff.sh` proves that a redeployed verifier
     * carries every setting over, and it reads this mapping by slot number. If the slot
     * moves and that script is not updated, the proof silently stops covering this
     * setting -- so the number is asserted here rather than left to review.
     *
     * 16, not 15: `_valueLimitWhitelist` is an `EnumerableSet.AddressSet`, which occupies
     * two slots (14 `_values`, 15 `_positions`).
     */
    function test_minimumTokenValueOf_lives_in_slot_16() public {
        vm.selectFork(crossForkID);

        uint expected = 7 * ONE_DOLLAR;
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, expected);

        // keccak256(abi.encode(key, slot)) -- the same math `verifier-settings-diff.sh`
        // does with its `mslot` helper.
        bytes32 slot = keccak256(abi.encode(address(testTokenCross), uint(16)));
        assertEq(
            uint(vm.load(address(bridgeVerifierCross), slot)),
            expected,
            "slot 16 must hold the per-token minimum; update verifier-settings-diff.sh if this moves"
        );
    }

    /// @notice `_defaultTokenPrice` is guarded against zero, so the price can never be 0.
    function test_default_token_price_cannot_be_zero() public {
        vm.selectFork(crossForkID);

        vm.expectRevert(
            abi.encodeWithSelector(BridgeVerifier.BridgeVerifierCanNotZeroValue.selector, "defaultTokenPrice")
        );
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setDefaultTokenPrice(0);
    }

    /// @notice Only `EDITOR_ROLE` may set an override.
    function test_setter_requires_editor_role() public {
        vm.selectFork(crossForkID);

        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, USER, Const.EDITOR_ROLE)
        );
        vm.prank(USER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 1);
    }

    /// @notice The zero address is rejected, mirroring `setExFeeRate`.
    function test_setter_rejects_zero_address() public {
        vm.selectFork(crossForkID);

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierCanNotZeroValue.selector, "token"));
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(IERC20(address(0)), 1);
    }

    /// @notice The setter emits the token-scoped event.
    function test_setter_emits_event() public {
        vm.selectFork(crossForkID);

        vm.expectEmit(true, false, false, true, address(bridgeVerifierCross));
        emit MinimumTokenValueOfSet(address(testTokenCross), 5 * ONE_DOLLAR);

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOf(testTokenCross, 5 * ONE_DOLLAR);
    }

    /// @notice The batch setter applies every entry.
    function test_batch_sets_every_entry() public {
        vm.selectFork(crossForkID);

        IERC20[] memory tokens = new IERC20[](2);
        uint[] memory values = new uint[](2);
        tokens[0] = testTokenCross;
        values[0] = 5 * ONE_DOLLAR; // $5 at $1 -> 5 TT
        tokens[1] = weth;
        values[1] = 20 * ONE_DOLLAR; // $20 at $10 -> 2 weth

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOfBatch(tokens, values);

        assertEq(_minimumOf(testTokenCross), 5 ether, "TT override applied");
        assertEq(_minimumOf(weth), 2 ether, "weth override applied");
    }

    /**
     * @notice A non-`EDITOR_ROLE` caller is rejected even for an empty-array batch call.
     * @dev Without the direct `onlyRole` on the batch entrypoint, an empty array skips the
     * loop entirely and the per-item check inside `setMinimumTokenValueOf` never runs, so
     * the call would otherwise succeed as an unauthorized no-op.
     */
    function test_batch_requires_editor_role_even_when_empty() public {
        vm.selectFork(crossForkID);

        IERC20[] memory tokens = new IERC20[](0);
        uint[] memory values = new uint[](0);

        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, USER, Const.EDITOR_ROLE)
        );
        vm.prank(USER);
        bridgeVerifierCross.setMinimumTokenValueOfBatch(tokens, values);
    }

    /// @notice A length mismatch reverts the whole batch.
    function test_batch_rejects_length_mismatch() public {
        vm.selectFork(crossForkID);

        IERC20[] memory tokens = new IERC20[](2);
        uint[] memory values = new uint[](1);
        tokens[0] = testTokenCross;
        tokens[1] = weth;
        values[0] = 1;

        vm.expectRevert(BridgeVerifier.BridgeVerifierInvalidLength.selector);
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOfBatch(tokens, values);
    }

    /**
     * @notice The batch is atomic: one bad entry rolls back the entries before it.
     * @dev Same guarantee `addValueLimitWhitelistBatch` documents — there is no partial
     * application to reason about afterwards.
     */
    function test_batch_is_atomic_on_a_bad_entry() public {
        vm.selectFork(crossForkID);

        IERC20[] memory tokens = new IERC20[](2);
        uint[] memory values = new uint[](2);
        tokens[0] = testTokenCross;
        values[0] = 5 * ONE_DOLLAR;
        tokens[1] = IERC20(address(0)); // rejected
        values[1] = 1;

        vm.expectRevert(abi.encodeWithSelector(BridgeVerifier.BridgeVerifierCanNotZeroValue.selector, "token"));
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValueOfBatch(tokens, values);

        assertEq(bridgeVerifierCross.getMinimumTokenValueOf(testTokenCross), 0, "first entry must be rolled back");
    }
}
