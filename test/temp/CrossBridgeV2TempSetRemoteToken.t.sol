// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {CrossBridgeV2} from "../../src/CrossBridgeV2.sol";

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {Const} from "../../src/lib/Const.sol";
import {CrossBridgeV2TempSetRemoteToken} from "../../src/temp/CrossBridgeV2TempSetRemoteToken.sol";

import {CrossBridgeV2Temp} from "../../script/temp/CrossBridgeV2Temp.s.sol";
import {BridgeExecutorTest} from "../BridgeExecutor.t.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

/**
 * @notice Exposes `CrossBridgeV2Temp._assertDrained` for direct testing (round-2 issue
 * plan H1) — mirrors `ForwardLibVerifyHarness is ForwardLibVerify` in
 * `test/ForwardLibVerify.t.sol`.
 */
contract CrossBridgeV2TempHarness is CrossBridgeV2Temp {
    function assertDrained(
        uint crossInit,
        uint hyperFin,
        uint hyperInit,
        uint crossFin,
        uint crossPending,
        uint hyperPending
    ) external pure {
        _assertDrained(crossInit, hyperFin, hyperInit, crossFin, crossPending, hyperPending);
    }
}

/**
 * @title CrossBridgeV2TempSetRemoteTokenTest
 * @notice Plan spec §14 T1-T9 for the disposable `CrossBridgeV2TempSetRemoteToken`
 * (branch `temp/cross-bridge-setremotetoken`, never merged to `dev`), plus round-2 issue
 * plan T10-T17 for `CrossBridgeV2Temp._assertDrained`/`preflightRemoteToken`. Fixture
 * follows `CrossBridgeV2Upgrade.t.sol`'s pattern: inherits `BridgeExecutorTest`, upgrades
 * the proxy INSIDE each test (never in `setUp`) so pre-upgrade state is always
 * capturable.
 * @dev T9 (preflight 3-branch fixed) and T10-T17 are covered here directly against
 * `CrossBridgeV2Temp`/`CrossBridgeV2TempHarness` (the script contract instantiated like
 * `ForwardLibVerify.t.sol` does with `ForwardLibVerify`) rather than only via the
 * dual-RPC dry-run, since the branch/comparison logic is a simple pure read plus
 * comparisons.
 */
contract CrossBridgeV2TempSetRemoteTokenTest is BridgeExecutorTest {
    /// @dev Target pair fixed by the temp impl (plan spec §6 "온체인 상수"). The base
    /// fixture's real chain ID is `BSC_CHAIN_ID` (57), not 998 — `TARGET_REMOTE_CHAIN_ID`
    /// is hardcoded in the temp impl, so the (998, NATIVE_TOKEN) pair must be explicitly
    /// registered in every test below.
    uint internal constant TARGET_REMOTE_CHAIN_ID = 998;

    /// @dev Stand-in for the real testnet stuck value
    /// (`0x3E0217c3926106b7E585B5341439f3150c6cab7c`, plan spec §6) — an arbitrary
    /// nonzero address is sufficient for these tests, which only check the field is
    /// read/written correctly, not that specific mainnet value.
    address internal constant OLD_REMOTE_TOKEN = address(0xDEAD);

    /// @dev Stand-in for the corrected HyperEVM-side token address (`T_new`).
    address internal constant T_NEW = address(0xBEEF);

    /// @dev Mirrors `CrossBridgeV2Temp.CROSS_BRIDGE` (private in the script, so the
    /// literal is duplicated here). Since round-2 issue plan M1 hardcodes the target
    /// proxy inside `preflightRemoteToken` instead of taking it as an argument, no local
    /// fork fixture can live at this exact address — T9 below mocks
    /// `IBridgeRegistry.getTokenPair` at this address directly instead.
    address internal constant CROSS_BRIDGE_CONST = 0xb81d6e000000000000000000000000000000C0de;

    /// @dev Mirrors `CrossBridgeV2Temp.EXPECTED_OLD_REMOTE_TOKEN` (private in the script,
    /// so the literal is duplicated here) — the real testnet stuck value (plan spec §6),
    /// used directly since round-2 issue plan M1 removed the `expectedOld` argument.
    address internal constant EXPECTED_OLD_REMOTE_TOKEN = 0x3E0217c3926106b7E585B5341439f3150c6cab7c;

    /// @dev Standard ERC1967 implementation slot
    /// (`keccak256("eip1967.proxy.implementation") - 1`), used to observe the proxy's
    /// raw implementation slot directly — the same technique
    /// `CrossBridgeV2Upgrade.t.sol`'s `BSC_CHAIN_ID_SLOT` uses for a known private slot,
    /// rather than trusting a getter that could itself be part of what's under test.
    bytes32 internal constant ERC1967_IMPLEMENTATION_SLOT = bytes32(uint(keccak256("eip1967.proxy.implementation")) - 1);

    /// @dev `_bscChainID`'s slot — confirmed identical in `CrossBridge`/`CrossBridgeV2`
    /// via `forge inspect` in `CrossBridgeV2Upgrade.t.sol` (slot 101). Since the temp
    /// impl adds zero storage variables, it is identical here too (also confirmed
    /// separately via the shell-level `forge inspect storageLayout` diff, see test.log).
    bytes32 internal constant BSC_CHAIN_ID_SLOT = bytes32(uint(101));

    function setUp() public virtual override {
        super.setUp();

        // `deposit()` (BridgeTest) moves `cross` ERC20 FROM USER on the BSC side, so
        // USER needs a funded + approved balance first (mirrors
        // CrossBridgeV2Upgrade.t.sol's identical setUp need).
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 1000 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);
    }

    // -------------------------------------------------------------------
    // Fixture helpers
    // -------------------------------------------------------------------

    /// @dev Upgrades the CROSS proxy from the base fixture's `CrossBridge` to a fresh
    /// `CrossBridgeV2` instance, mirroring `CrossBridgeV2Upgrade.t.sol`. Returns BOTH a
    /// proxy-typed handle (`v2`, for calling `CrossBridgeV2` functions through the proxy)
    /// and the raw deployed implementation address (`implAddr`, needed separately to
    /// restore to it later — the proxy address itself is never a valid
    /// `upgradeToAndCall` target).
    function _upgradeToV2() internal returns (CrossBridgeV2 v2, address implAddr) {
        vm.selectFork(crossForkID);
        CrossBridgeV2 v2Impl = new CrossBridgeV2();
        implAddr = address(v2Impl);
        vm.prank(CrossOWNER);
        bridgeCross.upgradeToAndCall(implAddr, bytes(""));
        v2 = CrossBridgeV2(payable(address(bridgeCross)));
    }

    /// @dev `_upgradeToV2()` plus registering the synthetic (998, NATIVE_TOKEN) pair the
    /// temp impl targets.
    function _upgradeToV2AndRegisterPair() internal returns (CrossBridgeV2 v2, address implAddr) {
        (v2, implAddr) = _upgradeToV2();
        vm.prank(CrossOWNER);
        v2.registerToken(TARGET_REMOTE_CHAIN_ID, true, address(NATIVE_TOKEN), OLD_REMOTE_TOKEN);
    }

    /// @dev Upgrades the (already-V2) proxy to a fresh temp-impl deployment.
    function _upgradeToTemp() internal returns (CrossBridgeV2TempSetRemoteToken temp) {
        vm.selectFork(crossForkID);
        CrossBridgeV2TempSetRemoteToken tempImpl = new CrossBridgeV2TempSetRemoteToken();
        vm.prank(CrossOWNER);
        CrossBridgeV2(payable(address(bridgeCross))).upgradeToAndCall(address(tempImpl), bytes(""));
        temp = CrossBridgeV2TempSetRemoteToken(payable(address(bridgeCross)));
    }

    /// @dev Restores the proxy to `restoreTo`. In production this is the fixed
    /// `0xe52bdBd35e0Db7b2efbaF17628ed0CD9D2bA8919` address (see
    /// `script/temp/CrossBridgeV2Temp.s.sol:ORIGINAL_IMPL`); in this local-fork suite the
    /// SAME `CrossBridgeV2` instance deployed by `_upgradeToV2()` is used instead, so T3
    /// can assert the impl slot returns to that EXACT address.
    function _restore(address restoreTo) internal {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        CrossBridgeV2TempSetRemoteToken(payable(address(bridgeCross))).upgradeToAndCall(restoreTo, bytes(""));
    }

    function _implSlot() internal view returns (address) {
        return address(uint160(uint(vm.load(address(bridgeCross), ERC1967_IMPLEMENTATION_SLOT))));
    }

    /// @dev Mocks `IBridgeRegistry(CROSS_BRIDGE_CONST).getTokenPair(998,
    /// Const.NATIVE_TOKEN).remoteToken` to return `current`, for testing
    /// `CrossBridgeV2Temp.preflightRemoteToken` (T9) against its hardcoded `CROSS_BRIDGE`
    /// target. No contract is deployed at `CROSS_BRIDGE_CONST` in this local-fork suite —
    /// `vm.mockCall` intercepts the external view call directly, which is sufficient since
    /// `preflightRemoteToken` only reads the `remoteToken` field.
    function _mockCurrentRemoteToken(address current) internal {
        vm.mockCall(
            CROSS_BRIDGE_CONST,
            abi.encodeWithSelector(IBridgeRegistry.getTokenPair.selector, TARGET_REMOTE_CHAIN_ID, Const.NATIVE_TOKEN),
            abi.encode(
                IBridgeRegistry.TokenPair({
                    localToken: address(0),
                    remoteToken: current,
                    isOrigin: false,
                    paused: false,
                    pendingAmount: 0,
                    deposited: 0,
                    minted: 0
                })
            )
        );
    }

    // -------------------------------------------------------------------
    // T1 (최우선/AC-2): storage-layout identity
    // -------------------------------------------------------------------

    /// @notice T1: proves layout identity between `CrossBridgeV2` and the temp impl via
    /// known-slot / getter reads taken immediately before and after switching the proxy
    /// to the temp impl — the same technique `CrossBridgeV2Upgrade.t.sol` uses for
    /// `_bscChainID`. The SOURCE-level proof (the compiler's own storageLayout for both
    /// contracts, byte-for-byte compared field by field) is run separately in the shell
    /// per the round's verification checklist; its output is captured in `test.log`
    /// alongside this test run.
    function test_T1_layoutIdentical_knownSlots() public {
        (CrossBridgeV2 v2,) = _upgradeToV2AndRegisterPair();

        vm.prank(CrossOWNER);
        v2.grantRole(Const.VERIFIER_ROLE, USER);
        vm.prank(CrossOWNER);
        // Must stay above CROSS_FOUNDATION_INITIAL_SUPPLY: `crossSupply()` already sits at
        // that baseline before any deposit (mirrors CrossBridgeV2Upgrade.t.sol's
        // `newLimit`), otherwise `deposit()` below falls into a CrossSupplyLimitExceeded
        // pending instead of succeeding.
        v2.setCrossSupplyLimit(123_456 ether + CROSS_FOUNDATION_INITIAL_SUPPLY);

        deposit(false, 5 ether, threshold); // bumps BSC_CHAIN_ID initiate/finalize + minted

        bytes32 bscChainIDSlotBefore = vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT);
        uint crossSupplyLimitBefore = v2.crossSupplyLimit();
        uint crossSupplyBefore = v2.crossSupply();
        bool hasRoleBefore = v2.hasRole(Const.VERIFIER_ROLE, USER);
        IBridgeRegistry.TokenPair memory pairBefore = v2.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        uint nextInitiateBefore = v2.getNextInitiateIndex(BSC_CHAIN_ID);
        uint nextFinalizeBefore = v2.getNextFinalizeIndex(BSC_CHAIN_ID);

        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        assertEq(
            vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT),
            bscChainIDSlotBefore,
            "_bscChainID storage slot must be byte-identical across the temp-impl upgrade"
        );
        assertEq(temp.crossSupplyLimit(), crossSupplyLimitBefore, "crossSupplyLimit must survive");
        assertEq(temp.crossSupply(), crossSupplyBefore, "crossSupply() (reads the _bscChainID pair) must survive");
        assertEq(temp.hasRole(Const.VERIFIER_ROLE, USER), hasRoleBefore, "roles must survive");

        IBridgeRegistry.TokenPair memory pairAfter = temp.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(pairAfter.localToken, pairBefore.localToken, "localToken must survive");
        assertEq(pairAfter.remoteToken, pairBefore.remoteToken, "remoteToken must survive (not yet set)");
        assertEq(pairAfter.isOrigin, pairBefore.isOrigin, "isOrigin must survive");
        assertEq(pairAfter.paused, pairBefore.paused, "paused must survive");
        assertEq(pairAfter.pendingAmount, pairBefore.pendingAmount, "pendingAmount must survive");
        assertEq(pairAfter.deposited, pairBefore.deposited, "deposited must survive");
        assertEq(pairAfter.minted, pairBefore.minted, "minted must survive");

        assertEq(temp.getNextInitiateIndex(BSC_CHAIN_ID), nextInitiateBefore, "initiate index must survive");
        assertEq(temp.getNextFinalizeIndex(BSC_CHAIN_ID), nextFinalizeBefore, "finalize index must survive");
    }

    // -------------------------------------------------------------------
    // T2 (AC-1/AC-3): setRemoteToken changes ONLY remoteToken
    // -------------------------------------------------------------------

    function test_T2_setRemoteToken_onlyRemoteTokenChanges() public {
        _upgradeToV2AndRegisterPair();
        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        IBridgeRegistry.TokenPair memory before = temp.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(before.remoteToken, OLD_REMOTE_TOKEN);

        vm.expectEmit(true, true, true, true);
        emit CrossBridgeV2TempSetRemoteToken.TempRemoteTokenSet(
            TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN), OLD_REMOTE_TOKEN, T_NEW
        );
        vm.prank(CrossOWNER);
        temp.setRemoteToken(T_NEW);

        IBridgeRegistry.TokenPair memory aft = temp.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(aft.remoteToken, T_NEW, "remoteToken must be updated");
        assertEq(aft.localToken, before.localToken, "localToken must be unchanged");
        assertEq(aft.isOrigin, before.isOrigin, "isOrigin must be unchanged");
        assertEq(aft.paused, before.paused, "paused must be unchanged");
        assertEq(aft.pendingAmount, before.pendingAmount, "pendingAmount must be unchanged");
        assertEq(aft.deposited, before.deposited, "deposited must be unchanged");
        assertEq(aft.minted, before.minted, "minted must be unchanged");
    }

    // -------------------------------------------------------------------
    // T3 (AC-4): upgrade -> set -> restore round trip
    // -------------------------------------------------------------------

    /// @notice After restoring to the original `CrossBridgeV2` implementation address:
    /// (a) the ERC1967 implementation slot is back to that exact address, (b)
    /// `remoteToken` is now `T_NEW` (the write persists — restoring only swaps the
    /// implementation, it never touches storage), and (c) every other observation (the
    /// pair's other 6 fields, `crossSupplyLimit`, roles, `_bscChainID`'s slot, an
    /// unrelated chain's pair, BSC_CHAIN_ID index progress) is identical to its
    /// pre-switch snapshot.
    function test_T3_upgradeSetRestore_roundTrip() public {
        (CrossBridgeV2 v2, address v2Addr) = _upgradeToV2AndRegisterPair();

        vm.prank(CrossOWNER);
        v2.grantRole(Const.VERIFIER_ROLE, USER);
        vm.prank(CrossOWNER);
        // Must stay above CROSS_FOUNDATION_INITIAL_SUPPLY — see T1's identical note.
        v2.setCrossSupplyLimit(999 ether + CROSS_FOUNDATION_INITIAL_SUPPLY);

        uint chainOther = 90402;
        vm.prank(CrossOWNER);
        v2.registerToken(chainOther, false, address(NATIVE_TOKEN), address(0x9402));

        deposit(false, 5 ether, threshold);

        bytes32 bscChainIDSlotBefore = vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT);
        uint crossSupplyLimitBefore = v2.crossSupplyLimit();
        bool hasRoleBefore = v2.hasRole(Const.VERIFIER_ROLE, USER);
        uint nextInitiateBefore = v2.getNextInitiateIndex(BSC_CHAIN_ID);
        uint nextFinalizeBefore = v2.getNextFinalizeIndex(BSC_CHAIN_ID);
        IBridgeRegistry.TokenPair memory pairBefore = v2.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        IBridgeRegistry.TokenPair memory otherPairBefore = v2.getTokenPair(chainOther, address(NATIVE_TOKEN));

        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();
        vm.prank(CrossOWNER);
        temp.setRemoteToken(T_NEW);

        _restore(v2Addr);

        assertEq(_implSlot(), v2Addr, "impl slot must be restored to the original CrossBridgeV2 address");

        CrossBridgeV2 restored = CrossBridgeV2(payable(address(bridgeCross)));
        IBridgeRegistry.TokenPair memory pairAfter =
            restored.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(pairAfter.remoteToken, T_NEW, "remoteToken must remain T_NEW after restore");
        assertEq(pairAfter.localToken, pairBefore.localToken, "localToken must be unchanged");
        assertEq(pairAfter.isOrigin, pairBefore.isOrigin, "isOrigin must be unchanged");
        assertEq(pairAfter.paused, pairBefore.paused, "paused must be unchanged");
        assertEq(pairAfter.pendingAmount, pairBefore.pendingAmount, "pendingAmount must be unchanged");
        assertEq(pairAfter.deposited, pairBefore.deposited, "deposited must be unchanged");
        assertEq(pairAfter.minted, pairBefore.minted, "minted must be unchanged");

        assertEq(
            vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT), bscChainIDSlotBefore, "_bscChainID slot must survive"
        );
        assertEq(restored.crossSupplyLimit(), crossSupplyLimitBefore, "crossSupplyLimit must survive");
        assertEq(restored.hasRole(Const.VERIFIER_ROLE, USER), hasRoleBefore, "roles must survive");
        assertEq(restored.getNextInitiateIndex(BSC_CHAIN_ID), nextInitiateBefore, "initiate index must survive");
        assertEq(restored.getNextFinalizeIndex(BSC_CHAIN_ID), nextFinalizeBefore, "finalize index must survive");

        IBridgeRegistry.TokenPair memory otherPairAfter = restored.getTokenPair(chainOther, address(NATIVE_TOKEN));
        assertEq(otherPairAfter.remoteToken, otherPairBefore.remoteToken, "unrelated pair must be untouched");
        assertEq(otherPairAfter.deposited, otherPairBefore.deposited, "unrelated pair deposited must be untouched");
        assertEq(otherPairAfter.minted, otherPairBefore.minted, "unrelated pair minted must be untouched");
    }

    // -------------------------------------------------------------------
    // T4-T6 (AC-7): revert paths
    // -------------------------------------------------------------------

    /// @notice T4: a non-ADMIN caller reverts with `AccessControlUnauthorizedAccount`.
    function test_T4_setRemoteToken_revertsForNonAdmin() public {
        _upgradeToV2AndRegisterPair();
        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        vm.prank(USER);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, USER, Const.ADMIN_ROLE)
        );
        temp.setRemoteToken(T_NEW);
    }

    /// @notice T5: `newRemoteToken == address(0)` reverts with `TempZeroRemoteToken`.
    function test_T5_setRemoteToken_revertsOnZero() public {
        _upgradeToV2AndRegisterPair();
        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        vm.prank(CrossOWNER);
        vm.expectRevert(CrossBridgeV2TempSetRemoteToken.TempZeroRemoteToken.selector);
        temp.setRemoteToken(address(0));
    }

    /// @notice T6: the target pair not being registered reverts with
    /// `TempPairNotRegistered` — uses `_upgradeToV2()` WITHOUT registering the pair.
    function test_T6_setRemoteToken_revertsWhenPairNotRegistered() public {
        _upgradeToV2();
        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        vm.prank(CrossOWNER);
        vm.expectRevert(CrossBridgeV2TempSetRemoteToken.TempPairNotRegistered.selector);
        temp.setRemoteToken(T_NEW);
    }

    // -------------------------------------------------------------------
    // T7-T8 (AC-8): existing functionality + re-entry
    // -------------------------------------------------------------------

    /// @notice T7: with the temp impl attached, existing `CrossBridgeV2` behavior is
    /// unaffected — `initialize` stays sealed (`Disabled()`), token-pair lookups work,
    /// and the pause family (`setPause`/`paused()`) works normally. Also confirms
    /// `setRemoteToken` itself deliberately has NO `whenNotPaused` guard (plan spec §11):
    /// it must succeed even while the bridge is paused.
    function test_T7_existingFunctionalityWorksOnTempImpl() public {
        _upgradeToV2AndRegisterPair();
        CrossBridgeV2TempSetRemoteToken temp = _upgradeToTemp();

        vm.expectRevert(CrossBridgeV2.Disabled.selector);
        temp.initialize(CrossOWNER, REWARD, threshold);

        IBridgeRegistry.TokenPair memory pair = temp.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(pair.remoteToken, OLD_REMOTE_TOKEN, "pair lookup must work normally on temp impl");

        assertFalse(temp.paused(), "bridge must start unpaused");
        vm.prank(CrossOWNER);
        temp.setPause(true);
        assertTrue(temp.paused(), "setPause(true) must work on temp impl");

        vm.prank(CrossOWNER);
        temp.setRemoteToken(T_NEW);
        assertEq(
            temp.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN)).remoteToken,
            T_NEW,
            "setRemoteToken must succeed even while paused (no whenNotPaused guard)"
        );

        vm.prank(CrossOWNER);
        temp.setPause(false);
        assertFalse(temp.paused(), "setPause(false) must work on temp impl");

        // Ordinary bridge functionality (deposit/finalize) still works on the temp impl.
        deposit(false, 5 ether, threshold);
    }

    /// @notice T8: after restoring to the original impl, the proxy can be switched to a
    /// fresh temp-impl deployment again and `setRemoteToken` still works identically
    /// with the same value (idempotent re-entry, plan spec §12 "재실행 가능").
    function test_T8_reentryAfterRestore() public {
        (, address v2Addr) = _upgradeToV2AndRegisterPair();

        CrossBridgeV2TempSetRemoteToken temp1 = _upgradeToTemp();
        vm.prank(CrossOWNER);
        temp1.setRemoteToken(T_NEW);
        _restore(v2Addr);

        assertEq(_implSlot(), v2Addr, "first restore must land on the original impl");
        assertEq(
            CrossBridgeV2(payable(address(bridgeCross))).getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN))
                .remoteToken,
            T_NEW
        );

        CrossBridgeV2TempSetRemoteToken temp2 = _upgradeToTemp();
        vm.prank(CrossOWNER);
        temp2.setRemoteToken(T_NEW);
        assertEq(
            temp2.getTokenPair(TARGET_REMOTE_CHAIN_ID, address(NATIVE_TOKEN)).remoteToken,
            T_NEW,
            "second round trip must succeed identically (idempotent)"
        );

        _restore(v2Addr);
        assertEq(_implSlot(), v2Addr, "second restore must also land on the original impl");
    }

    // -------------------------------------------------------------------
    // T9 (AC-6): preflightRemoteToken 3-branch
    // -------------------------------------------------------------------

    /// @notice T9: `CrossBridgeV2Temp.preflightRemoteToken`'s 3-way branch on the
    /// CURRENT `remoteToken` value — proceed / idempotent-skip / revert. Round-2 issue
    /// plan M1 hardcoded the target proxy to the real `CROSS_BRIDGE` constant (removing
    /// the `proxy`/`expectedOld` arguments), so this can no longer be exercised against
    /// this suite's own local-fork proxy — `_mockCurrentRemoteToken` mocks the external
    /// `getTokenPair` read at that exact constant address instead, covering the full
    /// function (guard + read + 3-way branch) exactly as it will actually run.
    function test_T9_preflightRemoteToken_threeBranches() public {
        vm.selectFork(crossForkID);
        CrossBridgeV2Temp preflight = new CrossBridgeV2Temp();

        // Branch 1: current == EXPECTED_OLD_REMOTE_TOKEN -> proceed (true), no revert.
        _mockCurrentRemoteToken(EXPECTED_OLD_REMOTE_TOKEN);
        bool shouldProceed = preflight.preflightRemoteToken(T_NEW);
        assertTrue(shouldProceed, "current == EXPECTED_OLD_REMOTE_TOKEN must signal proceed");

        // Branch 2: current == tNew (already applied) -> idempotent skip (false), no revert.
        _mockCurrentRemoteToken(T_NEW);
        shouldProceed = preflight.preflightRemoteToken(T_NEW);
        assertFalse(shouldProceed, "current == tNew must signal idempotent skip");

        // Branch 3: current is a third, unexplained value -> revert.
        address thirdValue = address(0xC0FFEE);
        _mockCurrentRemoteToken(thirdValue);

        vm.expectRevert(
            abi.encodeWithSelector(
                CrossBridgeV2Temp.PreflightRemoteTokenUnexpectedValue.selector,
                thirdValue,
                EXPECTED_OLD_REMOTE_TOKEN,
                T_NEW
            )
        );
        preflight.preflightRemoteToken(T_NEW);
    }

    // -------------------------------------------------------------------
    // T10-T15 (round-2 issue plan H1): _assertDrained via CrossBridgeV2TempHarness
    // -------------------------------------------------------------------

    /// @notice T10: all four indices matched and both pending counts empty -> passes
    /// (no revert), using generic (non-regression) numbers.
    function test_T10_assertDrained_matchedAndEmptyPending_passes() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        harness.assertDrained(10, 10, 3, 3, 0, 0);
    }

    /// @notice T11: `crossInit != hyperFin` -> `NotDrainedCrossToHyper` (the risk
    /// direction for this transition: an in-flight CROSS->Hyper bridge).
    function test_T11_assertDrained_revertsOnCrossToHyperMismatch() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        vm.expectRevert(abi.encodeWithSelector(CrossBridgeV2Temp.NotDrainedCrossToHyper.selector, 5, 4));
        harness.assertDrained(5, 4, 1, 1, 0, 0);
    }

    /// @notice T12: `hyperInit != crossFin` -> `NotDrainedHyperToCross`.
    function test_T12_assertDrained_revertsOnHyperToCrossMismatch() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        vm.expectRevert(abi.encodeWithSelector(CrossBridgeV2Temp.NotDrainedHyperToCross.selector, 2, 1));
        harness.assertDrained(4, 4, 2, 1, 0, 0);
    }

    /// @notice T13: `crossPending != 0` -> `PendingNotEmpty`.
    function test_T13_assertDrained_revertsOnCrossPendingNotEmpty() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        vm.expectRevert(abi.encodeWithSelector(CrossBridgeV2Temp.PendingNotEmpty.selector, 1, 0));
        harness.assertDrained(4, 4, 1, 1, 1, 0);
    }

    /// @notice T14: `hyperPending != 0` -> `PendingNotEmpty`.
    function test_T14_assertDrained_revertsOnHyperPendingNotEmpty() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        vm.expectRevert(abi.encodeWithSelector(CrossBridgeV2Temp.PendingNotEmpty.selector, 0, 1));
        harness.assertDrained(4, 4, 1, 1, 0, 1);
    }

    /// @notice T15: real observed-state regression — `(4, 4, 1, 1, 0, 0)` (2026-09-03
    /// testnet observation, round-2 issue plan H1) must pass, pinning that this drained
    /// state stays a pass across future edits to the comparison logic.
    function test_T15_assertDrained_observedTestnetState_passes() public {
        CrossBridgeV2TempHarness harness = new CrossBridgeV2TempHarness();
        harness.assertDrained(4, 4, 1, 1, 0, 0);
    }

    // -------------------------------------------------------------------
    // T16-T17 (round-2 issue plan M1): preflightRemoteToken fail-closed tNew guard
    // -------------------------------------------------------------------

    /// @notice T16: `tNew == address(0)` -> `InvalidNewRemoteToken`, checked before any
    /// on-chain read (no mock needed — the call reverts before touching `CROSS_BRIDGE`).
    function test_T16_preflightRemoteToken_revertsOnZeroTNew() public {
        CrossBridgeV2Temp preflight = new CrossBridgeV2Temp();
        vm.expectRevert(abi.encodeWithSelector(CrossBridgeV2Temp.InvalidNewRemoteToken.selector, address(0)));
        preflight.preflightRemoteToken(address(0));
    }

    /// @notice T17: `tNew == EXPECTED_OLD_REMOTE_TOKEN` -> `InvalidNewRemoteToken` — a
    /// `tNew` equal to the OLD value is not a "new" value at all; without this guard it
    /// would otherwise fall into the idempotent-skip branch only if `current` also
    /// happened to equal it, silently defeating fail-closed for the operator's actual
    /// mistake (passing the old address as `tNew`).
    function test_T17_preflightRemoteToken_revertsOnTNewEqualsExpectedOld() public {
        CrossBridgeV2Temp preflight = new CrossBridgeV2Temp();
        vm.expectRevert(
            abi.encodeWithSelector(CrossBridgeV2Temp.InvalidNewRemoteToken.selector, EXPECTED_OLD_REMOTE_TOKEN)
        );
        preflight.preflightRemoteToken(EXPECTED_OLD_REMOTE_TOKEN);
    }
}
