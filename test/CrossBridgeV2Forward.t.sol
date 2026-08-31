// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {BSCBridge} from "../src/BSCBridge.sol";
import {BSCBridgeV2} from "../src/BSCBridgeV2.sol";
import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {CrossBridgeV2} from "../src/CrossBridgeV2.sol";
import {BridgeRegistry} from "../src/abstract/BridgeRegistry.sol";
import {ICrossBridgeV2} from "../src/interface/ICrossBridgeV2.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";
import {ForwardLib} from "../src/lib/ForwardLib.sol";

import {BridgeExecutorTest, MockRevertingApproveMintableToken, MockTargetContract} from "./BridgeExecutor.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Vm} from "forge-std/Vm.sol";

/**
 * @title EvilForwardExecutor
 * @notice Stand-in `bridgeExecutor` (and, optionally, `_dev`) used ONLY to prove
 * `ForwardLib`'s call-scoped reentrancy guard blocks a nested call into
 * `bridgeTokenForwarded` made from within an in-flight forward (via the native fee
 * payout to `_dev` mid-`_initiateBridge`), while the legitimate OUTER forward call
 * still completes successfully. It mimics just enough of `IBridgeExecutor` to drive
 * `_finalizeBridge`'s executor path.
 */
contract EvilForwardExecutor {
    address public bridge;
    uint public bridgeTokenReentryChainID;
    bool public shouldReenterForward;
    bool public shouldReenterBridgeToken;
    bool public shouldReenterFinalizeBridgeBatch;
    bool public shouldReenterReleasePending;
    bool public reentryAttempted;
    bool public reentrySucceeded;
    bytes public reentryRevertData;

    constructor(address bridge_) {
        bridge = bridge_;
    }

    function setShouldReenterForward(bool v) external {
        shouldReenterForward = v;
    }

    function setShouldReenterBridgeToken(bool v, uint chainID) external {
        shouldReenterBridgeToken = v;
        bridgeTokenReentryChainID = chainID;
    }

    /// @dev M-1b regression: forward mid-flight reentry into `finalizeBridgeBatch`
    /// itself (an empty, harmless batch) must be blocked by the SAME pre-existing
    /// shared `nonReentrant` guard `bridgeToken` reentry is blocked by — proving the
    /// guard covers `finalizeBridgeBatch` too, not just `bridgeToken`.
    function setShouldReenterFinalizeBridgeBatch(bool v) external {
        shouldReenterFinalizeBridgeBatch = v;
    }

    /// @dev M-1b regression: forward mid-flight reentry into `releasePending` (with a
    /// nonexistent index — any revert reason is fine, we only care that the call is
    /// blocked by the shared guard BEFORE it could reach `BaseBridgeNotExistIndex`).
    function setShouldReenterReleasePending(bool v) external {
        shouldReenterReleasePending = v;
    }

    function isWhitelistedTarget(address) external pure returns (bool) {
        return true;
    }

    function executeExtraCall(uint, uint, IERC20 toToken, address, uint value, bytes calldata extraData)
        external
        payable
        returns (uint consumed, bytes memory returnData)
    {
        bool isNative = address(toToken) == Const.NATIVE_TOKEN;
        address target = address(bytes20(extraData[:20]));
        (bool ok, bytes memory ret) = target.call{value: isNative ? value : 0}(extraData[20:]);
        require(ok, "target call failed");
        consumed = value;
        returnData = ret;
    }

    // Fires when the bridge forwards the native `_dev` fee mid-`_initiateBridge`, i.e.
    // while the outer `bridgeTokenForwarded` call still holds the forward guard.
    receive() external payable {
        if (shouldReenterForward) {
            reentryAttempted = true;
            shouldReenterForward = false; // avoid infinite recursion
            try ICrossBridgeV2(bridge).bridgeTokenForwarded(0, address(0xdead), 0, 0, 0, "") {
                reentrySucceeded = true;
            } catch (bytes memory reason) {
                reentryRevertData = reason;
            }
        } else if (shouldReenterBridgeToken) {
            reentryAttempted = true;
            shouldReenterBridgeToken = false;
            // toChainID/token must be a REGISTERED pair so the call reaches the shared
            // `nonReentrant` guard (past `onlyValidToken`) rather than failing earlier.
            try CrossBridgeV2(payable(bridge)).bridgeToken(
                bridgeTokenReentryChainID, IERC20(Const.NATIVE_TOKEN), address(0xdead), 0, 0, 0, ""
            ) {
                reentrySucceeded = true;
            } catch (bytes memory reason) {
                reentryRevertData = reason;
            }
        } else if (shouldReenterFinalizeBridgeBatch) {
            reentryAttempted = true;
            shouldReenterFinalizeBridgeBatch = false;
            // Empty batch: the shared `nonReentrant` guard (already held by the outer
            // `finalizeBridgeBatch` call this reentry happens inside of) must reject
            // this before the empty-array body would otherwise trivially succeed.
            try CrossBridgeV2(payable(bridge)).finalizeBridgeBatch(
                new IBridgeRegistry.FinalizeArguments[](0), new uint8[][](0), new bytes32[][](0), new bytes32[][](0)
            ) returns (bool) {
                reentrySucceeded = true;
            } catch (bytes memory reason) {
                reentryRevertData = reason;
            }
        } else if (shouldReenterReleasePending) {
            reentryAttempted = true;
            shouldReenterReleasePending = false;
            try CrossBridgeV2(payable(bridge)).releasePending(999999999, 999999999) {
                reentrySucceeded = true;
            } catch (bytes memory reason) {
                reentryRevertData = reason;
            }
        }
    }
}

/**
 * @title NonConsumingExecutor
 * @notice Stand-in `bridgeExecutor` used ONLY to reproduce `ForwardFailureCode.ForwardCallNotConsumed`:
 * returns a well-formed (>= 64 byte) SUCCESS response without ever actually calling the
 * target (`bridgeTokenForwarded`), so the staged forward context is never consumed.
 * @dev Plan spec §10 / test/M-1: this is a deliberately non-compliant/malicious executor
 * shape — a spec-compliant `BridgeExecutor` can never produce this outcome, since it
 * always calls the target before returning success. It exists purely as defense-in-depth
 * coverage for "what if `bridgeExecutor` were misconfigured or compromised".
 */
contract NonConsumingExecutor {
    function isWhitelistedTarget(address) external pure returns (bool) {
        return true;
    }

    function executeExtraCall(uint, uint, IERC20, address, uint value, bytes calldata)
        external
        payable
        returns (uint consumed, bytes memory returnData)
    {
        // Never touches `extraData`'s target at all.
        consumed = value;
        returnData = "";
    }
}

/**
 * @title MaliciousNativeReceiver
 * @notice A native-token finalize recipient (`to`) whose receive callback attempts to
 * reenter both `bridgeTokenForwarded` and `bridgeToken`. Used for the §14.1 regression
 * proving BaseBridge's ordinary `_safeCall(to, value, "")` native payout callback (gas
 * 100k) — ATTACKER-CONTROLLED since `to` is a caller-supplied address — cannot reach
 * either entrypoint: `bridgeTokenForwarded` because `msg.sender` here is this receiver,
 * not `bridgeExecutor`; `bridgeToken` because the shared `nonReentrant` guard is already
 * held by the outer `finalizeBridgeBatch` call this callback fires inside of.
 */
contract MaliciousNativeReceiver {
    address public bridge;
    uint public bridgeTokenReentryChainID;
    bool public forwardAttempted;
    bool public forwardSucceeded;
    bytes public forwardRevertData;
    bool public bridgeTokenAttempted;
    bool public bridgeTokenSucceeded;
    bytes public bridgeTokenRevertData;

    constructor(address bridge_, uint bridgeTokenReentryChainID_) {
        bridge = bridge_;
        bridgeTokenReentryChainID = bridgeTokenReentryChainID_;
    }

    receive() external payable {
        forwardAttempted = true;
        try ICrossBridgeV2(bridge).bridgeTokenForwarded(0, address(0xdead), 0, 0, 0, "") {
            forwardSucceeded = true;
        } catch (bytes memory reason) {
            forwardRevertData = reason;
        }

        bridgeTokenAttempted = true;
        try CrossBridgeV2(payable(bridge)).bridgeToken(
            bridgeTokenReentryChainID, IERC20(Const.NATIVE_TOKEN), address(0xdead), 0, 0, 0, ""
        ) {
            bridgeTokenSucceeded = true;
        } catch (bytes memory reason) {
            bridgeTokenRevertData = reason;
        }
    }
}

/**
 * @title RevertingReceiver
 * @notice Rejects any native ETH sent to it. Used as a `_dev` stand-in to force
 * `_initiateBridge`'s native fee payout (mid-`bridgeTokenForwarded`) to revert, taking
 * down the entire forwarded call — the vehicle for the M-1 `BridgeInitiated`-rollback
 * regression (plan spec §12 atomicity).
 */
contract RevertingReceiver {
    receive() external payable {
        revert("RevertingReceiver: no thanks");
    }
}

/**
 * @title CrossBridgeV2ForwardTest
 * @notice Forward (multi-hop) entrypoint test suite per plan spec §14.1 (guards /
 * access control), §14.2 (D1 pass-through invariant), and §14.6 (code-size gates).
 * @dev Fixture upgrade order follows §14.7: `CrossBridge` is deployed and initialized
 * exactly as in the base `BridgeExecutorTest` fixture, and is only THEN upgraded to
 * `CrossBridgeV2` — mirroring the real upgrade path rather than deploying V2 fresh.
 */
contract CrossBridgeV2ForwardTest is BridgeExecutorTest {
    CrossBridgeV2 internal bridgeCrossV2;

    function setUp() public virtual override {
        super.setUp();

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);
        bridgeCross.upgradeToAndCall(address(new CrossBridgeV2()), bytes(""));
        bridgeCrossV2 = CrossBridgeV2(payable(address(bridgeCross)));

        // The forwarded entrypoint's own target is the bridge itself.
        bridgeExecutorCross.addWhitelistTarget(address(bridgeCrossV2));

        // Keep monitoring thresholds out of the way for synthetic/unconfigured tokens
        // and chain IDs used by the D1 tests below.
        bridgeVerifierCross.setVerificationAmountThreshold(0);
        bridgeVerifierCross.setPeriodTotalValueThreshold(0);
        vm.stopPrank();
    }

    // ----------------------------------------------------------------
    // Helpers
    // ----------------------------------------------------------------

    function _forwardExtraData(uint toChainID2, address to2, uint value2, uint networkFee2, uint exFee2, bytes memory hop3)
        internal
        view
        returns (bytes memory)
    {
        bytes memory call_ = abi.encodeWithSelector(
            CrossBridgeV2.bridgeTokenForwarded.selector, toChainID2, to2, value2, networkFee2, exFee2, hop3
        );
        return abi.encodePacked(address(bridgeCrossV2), call_);
    }

    /// @dev Mirrors `CrossChainTest.crossFinalize`'s signing logic but is generic over
    /// `fromChainID`/`token`, so it can also finalize from the synthetic chain IDs the
    /// D1 tests register.
    function _signAndFinalize(
        uint fromChainID,
        uint index,
        address token,
        address to,
        uint value,
        bytes memory extraData,
        uint sigCount
    ) internal returns (bool ok) {
        vm.selectFork(crossForkID);
        if (sigCount > threshold) sigCount = threshold;

        bytes32 h =
            keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, keccak256(extraData)));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), h);

        uint8[] memory v = new uint8[](sigCount);
        bytes32[] memory r = new bytes32[](sigCount);
        bytes32[] memory s = new bytes32[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], hash);
        }

        IBridgeRegistry.FinalizeArguments[] memory args = new IBridgeRegistry.FinalizeArguments[](1);
        args[0] = IBridgeRegistry.FinalizeArguments({
            fromChainID: fromChainID,
            index: index,
            toToken: IERC20(token),
            to: to,
            value: value,
            extraData: extraData
        });
        uint8[][] memory vArray = new uint8[][](1);
        bytes32[][] memory rArray = new bytes32[][](1);
        bytes32[][] memory sArray = new bytes32[][](1);
        vArray[0] = v;
        rArray[0] = r;
        sArray[0] = s;
        ok = bridgeCross.finalizeBridgeBatch(args, vArray, rArray, sArray);
    }

    /// @dev Drives hop-1 (BSC -> CROSS, native) for exactly `ctxValue` principal and
    /// finalizes it on CROSS with `extraData`, WITHOUT asserting the forward's own
    /// outcome — callers assert on that themselves.
    function _hop1NativeAndFinalize(uint ctxValue, bytes memory extraData) internal {
        vm.selectFork(bscForkID);
        (, uint networkFee1, uint exFee1) =
            bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValue);
        uint total = ctxValue + networkFee1 + exFee1;
        vm.deal(USER, total);
        uint index = nextIndexBSC;
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: total}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValue, networkFee1, exFee1, extraData
        );
        bscIncrementIndex();

        _signAndFinalize(BSC_CHAIN_ID, index, address(NATIVE_TOKEN), USER, ctxValue, extraData, 5);
    }

    /// @dev Stages a valid forward context by driving a real BSC -> CROSS native hop-1
    /// finalize whose extraData targets `bridgeTokenForwarded` with the given hop-2
    /// params. Uses exact verifier-computed fees on both legs so the forwarded call's
    /// `value + networkFee + exFee == ctxValue` invariant holds by construction.
    /// @dev For an isOrigin=true pair, `_withdrawToken` decrements `deposited`, so any
    /// finalize FROM that chain needs a prior real deposit at least as large as what
    /// will be withdrawn, or it underflows. Drives one ordinary (non-forward) bridge-out
    /// to seed `deposited[chainID][token]`.
    function _bumpDeposited(uint chainID, IERC20 token, uint amount) internal {
        (, uint fee, uint exFee) = bridgeVerifierCross.calculateFee(chainID, token, amount);
        deal(address(token), USER, amount + fee + exFee);
        vm.prank(USER);
        token.approve(address(bridgeCrossV2), amount + fee + exFee);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken(chainID, token, USER, amount, fee, exFee, "");
    }

    function _stageNativeForward(uint value2, uint networkFee2, uint exFee2, bytes memory hop3) internal {
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, networkFee2, exFee2, hop3);
        uint ctxValue = value2 + networkFee2 + exFee2;
        _hop1NativeAndFinalize(ctxValue, extraData);
    }

    // ----------------------------------------------------------------
    // M-1: event-log decoding helpers
    // ----------------------------------------------------------------

    bytes32 internal constant FORWARD_FAILED_TOPIC0 = keccak256("ForwardFailed(uint256,uint256,uint8,bytes32)");
    bytes32 internal constant FORWARD_INITIATED_TOPIC0 =
        keccak256("ForwardInitiated(uint256,uint256,uint256,address,address,uint256,uint256,uint256,bytes32)");

    /// @dev Finds the (first) `ForwardFailed` log emitted BY THE BRIDGE (DELEGATECALL
    /// means `ForwardLib`'s events carry the bridge's own address as emitter — plan spec
    /// §9) among `logs`, and decodes its fields.
    function _findForwardFailed(Vm.Log[] memory logs)
        internal
        view
        returns (bool found, uint fromChainID, uint finalizeIndex, ForwardLib.ForwardFailureCode code, bytes32 reasonHash)
    {
        for (uint i = 0; i < logs.length; i++) {
            if (
                logs[i].emitter == address(bridgeCrossV2) && logs[i].topics.length == 3
                    && logs[i].topics[0] == FORWARD_FAILED_TOPIC0
            ) {
                fromChainID = uint(logs[i].topics[1]);
                finalizeIndex = uint(logs[i].topics[2]);
                (uint8 codeRaw, bytes32 rh) = abi.decode(logs[i].data, (uint8, bytes32));
                return (true, fromChainID, finalizeIndex, ForwardLib.ForwardFailureCode(codeRaw), rh);
            }
        }
    }

    /// @dev Finds the (first) `ForwardInitiated` log emitted BY THE BRIDGE among `logs`,
    /// and decodes its fields.
    function _findForwardInitiated(Vm.Log[] memory logs)
        internal
        view
        returns (
            bool found,
            uint fromChainID,
            uint finalizeIndex,
            uint toChainID,
            address fromToken,
            address to,
            uint value,
            uint networkFee,
            uint exFee,
            bytes32 extraDataHash
        )
    {
        for (uint i = 0; i < logs.length; i++) {
            if (
                logs[i].emitter == address(bridgeCrossV2) && logs[i].topics.length == 4
                    && logs[i].topics[0] == FORWARD_INITIATED_TOPIC0
            ) {
                fromChainID = uint(logs[i].topics[1]);
                finalizeIndex = uint(logs[i].topics[2]);
                toChainID = uint(logs[i].topics[3]);
                (fromToken, to, value, networkFee, exFee, extraDataHash) =
                    abi.decode(logs[i].data, (address, address, uint, uint, uint, bytes32));
                return (true, fromChainID, finalizeIndex, toChainID, fromToken, to, value, networkFee, exFee, extraDataHash);
            }
        }
    }

    /// @dev True if ANY `ForwardFailed` log (emitted by the bridge) is present.
    function _hasForwardFailed(Vm.Log[] memory logs) internal view returns (bool found) {
        (found,,,,) = _findForwardFailed(logs);
    }

    /// @dev Discovers the `ForwardLib` address forge auto-linked for `CrossBridgeV2` in
    /// this test run. Foundry predeploys/reuses a SINGLE external-library instance per
    /// test run for a given library, so this is the SAME address baked into
    /// `bridgeCrossV2`'s real implementation set up in `setUp()`. Offset derived from
    /// `out/CrossBridgeV2.sol/CrossBridgeV2.json`'s `deployedBytecode.linkReferences`
    /// (stable for this round — `CrossBridgeV2.sol`'s logic is unchanged this round; if
    /// it ever changes, re-derive via `script/ForwardLibVerify.s.sol`'s own JSON-based
    /// approach instead of a hardcoded offset).
    function _discoverForwardLibAddress() internal returns (address addr) {
        bytes memory code = address(new CrossBridgeV2()).code;
        uint offset = 8638;
        assembly ("memory-safe") {
            addr := shr(96, mload(add(add(code, 32), offset)))
        }
    }

    // ----------------------------------------------------------------
    // §14.1 Guards / access control
    // ----------------------------------------------------------------

    function test_forward_directCallByEOA_reverts() public {
        vm.selectFork(crossForkID);
        vm.prank(USER);
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardNotExecutor.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, 1, 0, 0, "");
    }

    function test_forward_directCallByArbitraryContract_reverts() public {
        vm.selectFork(crossForkID);
        vm.prank(address(mockTargetCross));
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardNotExecutor.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, 1, 0, 0, "");
    }

    function test_forward_executorOutsideFinalizeContext_reverts() public {
        vm.selectFork(crossForkID);
        vm.prank(address(bridgeExecutorCross));
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardContextInactive.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, 1, 0, 0, "");
    }

    /// @notice Validation priority: the msg.sender check fires before the paused check,
    /// even though both would independently fail.
    function test_forward_priority_senderCheckBeforePaused() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(true);

        vm.prank(USER); // not the executor
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardNotExecutor.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, 1, 0, 0, "");
    }

    /// @notice Validation priority regression for the "no modifiers" design (§7.3):
    /// with BOTH paused AND an inactive context, the caller sees the ctx error, NOT
    /// `Pausable`'s error — proving the ordering is msg.sender -> ctx -> paused ->
    /// token, implemented in the function body rather than via modifiers.
    function test_forward_priority_contextCheckBeforePaused() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(true);

        vm.prank(address(bridgeExecutorCross));
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardContextInactive.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, 1, 0, 0, "");
    }

    /// @notice A staged context is consumed exactly once: a legitimate forward succeeds,
    /// then a sequential second call (as the executor, with no freshly staged context)
    /// reverts with the same "inactive" error — proving there is no residual/replayable
    /// context (regression for the early-return `_forwardEnd`/clear placement).
    function test_forward_successThenSequentialReplay_reverts() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        _stageNativeForward(value2, networkFee2, exFee2, "");

        // sanity: the staged forward actually succeeded (no fallback payout to USER)
        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "forward should have consumed the finalize payout, not paid USER");

        vm.prank(address(bridgeExecutorCross));
        vm.expectRevert(ICrossBridgeV2.BaseBridgeForwardContextInactive.selector);
        bridgeCrossV2.bridgeTokenForwarded(BSC_CHAIN_ID, USER, value2, networkFee2, exFee2, "");
    }

    /// @notice Nested/reentrant call into `bridgeTokenForwarded` while the forward guard
    /// is still held (triggered via the native `_dev` fee payout mid-`_initiateBridge`)
    /// reverts with `ForwardReentrantCall`, while the legitimate OUTER call still
    /// completes successfully — proving the guard is a pure defense-in-depth layer, not
    /// load-bearing for the outer happy path.
    function test_forward_nestedReentrancy_reverts() public {
        vm.selectFork(crossForkID);
        EvilForwardExecutor evil = new EvilForwardExecutor(address(bridgeCrossV2));
        vm.startPrank(CrossOWNER);
        bridgeCrossV2.setBridgeExecutor(IBridgeExecutor(address(evil)));
        bridgeCrossV2.setDev(payable(address(evil)));
        vm.stopPrank();
        evil.setShouldReenterForward(true);

        uint value2 = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(networkFee2 + exFee2 > 0, "test requires a nonzero fee to trigger the _dev callback");
        _stageNativeForward(value2, networkFee2, exFee2, "");

        assertTrue(evil.reentryAttempted(), "nested call was never attempted");
        assertFalse(evil.reentrySucceeded(), "nested call must not succeed");
        assertEq(bytes4(evil.reentryRevertData()), ForwardLib.ForwardReentrantCall.selector);
    }

    /// @notice Regression for pre-existing behaviour: while a finalize (and hence any
    /// forward nested inside it) is executing, `finalizeBridgeBatch`'s own SHARED
    /// `nonReentrant` guard is held for the whole call, so a nested call into the
    /// unrelated `bridgeToken` entrypoint is blocked by that pre-existing guard, not by
    /// anything `ForwardLib` added.
    function test_forward_bridgeTokenReentrancy_blockedByExistingGuard() public {
        vm.selectFork(crossForkID);
        EvilForwardExecutor evil = new EvilForwardExecutor(address(bridgeCrossV2));
        vm.startPrank(CrossOWNER);
        bridgeCrossV2.setBridgeExecutor(IBridgeExecutor(address(evil)));
        bridgeCrossV2.setDev(payable(address(evil)));
        vm.stopPrank();
        evil.setShouldReenterBridgeToken(true, BSC_CHAIN_ID);

        uint value2 = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(networkFee2 + exFee2 > 0, "test requires a nonzero fee to trigger the _dev callback");
        _stageNativeForward(value2, networkFee2, exFee2, "");

        assertTrue(evil.reentryAttempted());
        assertFalse(evil.reentrySucceeded());
        // ReentrancyGuardReentrantCall(), from ReentrancyGuardTransientUpgradeable (unchanged, pre-existing guard)
        assertEq(bytes4(evil.reentryRevertData()), bytes4(keccak256("ReentrancyGuardReentrantCall()")));
    }

    /// @notice Priority case for this round: a batch with TWO forwarded items must have
    /// BOTH succeed, proving the forward guard is call-scoped (acquired/released per
    /// `bridgeTokenForwarded` invocation), not transaction/batch-scoped.
    function test_forward_batchOfTwoForwardedItems_bothSucceed() public {
        vm.selectFork(crossForkID);
        uint hop2InitiateIndexBefore = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);

        uint value2a = 1 ether;
        uint value2b = 2 ether;
        (, uint fee2a, uint ex2a) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2a);
        (, uint fee2b, uint ex2b) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2b);
        uint ctxA = value2a + fee2a + ex2a;
        uint ctxB = value2b + fee2b + ex2b;
        bytes memory extraDataA = _forwardExtraData(BSC_CHAIN_ID, USER, value2a, fee2a, ex2a, "");
        bytes memory extraDataB = _forwardExtraData(BSC_CHAIN_ID, USER, value2b, fee2b, ex2b, "");

        // Two independent hop-1 initiations from BSC (sequential indices).
        vm.selectFork(bscForkID);
        (, uint netA1, uint exA1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxA);
        (, uint netB1, uint exB1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxB);
        uint indexA = nextIndexBSC;
        vm.deal(USER, ctxA + netA1 + exA1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxA + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxA, netA1, exA1, extraDataA
        );
        bscIncrementIndex();

        uint indexB = nextIndexBSC;
        vm.deal(USER, ctxB + netB1 + exB1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxB + netB1 + exB1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxB, netB1, exB1, extraDataB
        );
        bscIncrementIndex();

        // One finalize batch containing both items.
        vm.selectFork(crossForkID);
        bytes32 hA = keccak256(
            abi.encode(FINALIZE_TYPEHASH, BSC_CHAIN_ID, indexA, address(NATIVE_TOKEN), USER, ctxA, keccak256(extraDataA))
        );
        bytes32 hB = keccak256(
            abi.encode(FINALIZE_TYPEHASH, BSC_CHAIN_ID, indexB, address(NATIVE_TOKEN), USER, ctxB, keccak256(extraDataB))
        );
        bytes32 hashA = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hA);
        bytes32 hashB = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hB);

        uint8[] memory vA = new uint8[](3);
        bytes32[] memory rA = new bytes32[](3);
        bytes32[] memory sA = new bytes32[](3);
        uint8[] memory vB = new uint8[](3);
        bytes32[] memory rB = new bytes32[](3);
        bytes32[] memory sB = new bytes32[](3);
        for (uint i = 0; i < 3; i++) {
            (vA[i], rA[i], sA[i]) = vm.sign(VALIDATOR_PKs[i], hashA);
            (vB[i], rB[i], sB[i]) = vm.sign(VALIDATOR_PKs[i], hashB);
        }

        IBridgeRegistry.FinalizeArguments[] memory args = new IBridgeRegistry.FinalizeArguments[](2);
        args[0] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexA,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxA,
            extraData: extraDataA
        });
        args[1] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexB,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxB,
            extraData: extraDataB
        });
        uint8[][] memory vArray = new uint8[][](2);
        bytes32[][] memory rArray = new bytes32[][](2);
        bytes32[][] memory sArray = new bytes32[][](2);
        vArray[0] = vA;
        rArray[0] = rA;
        sArray[0] = sA;
        vArray[1] = vB;
        rArray[1] = rB;
        sArray[1] = sB;

        assertTrue(bridgeCross.finalizeBridgeBatch(args, vArray, rArray, sArray));

        // Neither forwarded item fell back to a direct USER payout on CROSS...
        assertEq(USER.balance, 0, "neither forwarded item should fall back to a direct USER payout");
        // ...and BOTH actually initiated a new hop-2 bridge from CROSS (guard released
        // between items, so the second was not spuriously blocked).
        assertEq(bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID), hop2InitiateIndexBefore + 2);
    }

    /// @notice A forward failure (revert) in one batch item must not leave the guard
    /// held for the next item — proving the guard's EIP-1153 auto-release-on-revert
    /// lifecycle, not just its acquire/release on the happy path.
    function test_forward_batchRevertThenNextItemStillSucceeds() public {
        vm.selectFork(crossForkID);

        // Item 1: deliberately mismatched amount -> ExecutorCallReverted -> guard
        // auto-released by the EIP-1153 rollback of the reverted inner call.
        uint value2a = 1 ether;
        (, uint fee2a, uint ex2a) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2a);
        uint ctxA = value2a + fee2a + ex2a;
        bytes memory extraDataA = _forwardExtraData(BSC_CHAIN_ID, USER, value2a, fee2a, ex2a, "");

        // Item 2: valid forward.
        uint value2b = 2 ether;
        (, uint fee2b, uint ex2b) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2b);
        uint ctxB = value2b + fee2b + ex2b;
        bytes memory extraDataB = _forwardExtraData(BSC_CHAIN_ID, USER, value2b, fee2b, ex2b, "");

        vm.selectFork(bscForkID);
        (, uint netA1, uint exA1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxA + 1);
        uint indexA = nextIndexBSC;
        vm.deal(USER, ctxA + 1 + netA1 + exA1);
        vm.prank(USER);
        // hop-1 finalizes ctxA + 1 wei so hop-2's exact-equality check fails by 1 wei.
        bridgeBSC.bridgeToken{value: ctxA + 1 + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxA + 1, netA1, exA1, extraDataA
        );
        bscIncrementIndex();

        (, uint netB1, uint exB1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxB);
        uint indexB = nextIndexBSC;
        vm.deal(USER, ctxB + netB1 + exB1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxB + netB1 + exB1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxB, netB1, exB1, extraDataB
        );
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        bytes32 hA = keccak256(
            abi.encode(
                FINALIZE_TYPEHASH, BSC_CHAIN_ID, indexA, address(NATIVE_TOKEN), USER, ctxA + 1, keccak256(extraDataA)
            )
        );
        bytes32 hB = keccak256(
            abi.encode(FINALIZE_TYPEHASH, BSC_CHAIN_ID, indexB, address(NATIVE_TOKEN), USER, ctxB, keccak256(extraDataB))
        );
        bytes32 hashA = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hA);
        bytes32 hashB = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hB);

        uint8[] memory vA = new uint8[](3);
        bytes32[] memory rA = new bytes32[](3);
        bytes32[] memory sA = new bytes32[](3);
        uint8[] memory vB = new uint8[](3);
        bytes32[] memory rB = new bytes32[](3);
        bytes32[] memory sB = new bytes32[](3);
        for (uint i = 0; i < 3; i++) {
            (vA[i], rA[i], sA[i]) = vm.sign(VALIDATOR_PKs[i], hashA);
            (vB[i], rB[i], sB[i]) = vm.sign(VALIDATOR_PKs[i], hashB);
        }

        IBridgeRegistry.FinalizeArguments[] memory args = new IBridgeRegistry.FinalizeArguments[](2);
        args[0] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexA,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxA + 1,
            extraData: extraDataA
        });
        args[1] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexB,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxB,
            extraData: extraDataB
        });
        uint8[][] memory vArray = new uint8[][](2);
        bytes32[][] memory rArray = new bytes32[][](2);
        bytes32[][] memory sArray = new bytes32[][](2);
        vArray[0] = vA;
        rArray[0] = rA;
        sArray[0] = sA;
        vArray[1] = vB;
        rArray[1] = rB;
        sArray[1] = sB;

        assertTrue(bridgeCross.finalizeBridgeBatch(args, vArray, rArray, sArray));

        // Item 1's forward reverted, so it fell back to a direct USER payout...
        assertEq(USER.balance, ctxA + 1, "item 1 should have fallen back to USER");
        // ...but item 2, right after it in the same batch, still forwarded successfully.
    }

    // ----------------------------------------------------------------
    // §14.2 D1 pass-through invariant
    // ----------------------------------------------------------------

    function test_D1_erc20_bothOrigin_success() public {
        vm.selectFork(crossForkID);
        TestToken myToken = new TestToken("Origin", "ORG", 18);
        myToken.mint(address(bridgeCross), 1000 ether);

        uint chainX = 90001;
        uint chainY = 90002;
        vm.startPrank(CrossOWNER);
        bridgeCross.registerToken(chainX, true, address(myToken), address(0x1111));
        bridgeCross.registerToken(chainY, true, address(myToken), address(0x2222));
        vm.stopPrank();

        uint value2 = 10 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainY, IERC20(address(myToken)), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(chainY, USER, value2, fee2, ex2, "");

        // chainX's pair is isOrigin=true, so the hop-1 finalize's `_withdrawToken` will
        // decrement `deposited[chainX][myToken]` — seed it first.
        _bumpDeposited(chainX, IERC20(address(myToken)), ctxValue * 2);

        assertTrue(_signAndFinalize(chainX, 1, address(myToken), USER, ctxValue, extraData, 5));

        assertEq(myToken.balanceOf(address(bridgeExecutorCross)), 0, "executor should not retain any ORG token");
        assertEq(bridgeCross.getTokenPair(chainY, address(myToken)).deposited, value2);
        assertEq(myToken.balanceOf(USER), 0, "D1 both-origin success must NOT fall back to a direct USER payout");
    }

    function test_D1_erc20_fromNotOrigin_reverts() public {
        vm.selectFork(crossForkID);
        uint chainX = 90003; // from-chain pair: NOT origin (wrapped)
        uint chainY = 90004; // to-chain pair: origin

        vm.startPrank(CrossOWNER);
        address wrappedAddr = bridgeCross.createToken(chainX, address(0x3333), "WRP2", 18);
        bridgeCross.registerToken(chainY, true, wrappedAddr, address(0x4444));
        vm.stopPrank();

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainY, IERC20(wrappedAddr), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(chainY, USER, value2, fee2, ex2, "");

        assertTrue(_signAndFinalize(chainX, 1, wrappedAddr, USER, ctxValue, extraData, 5));

        // D1 violation -> bridgeTokenForwarded reverts -> normal finalize fallback pays USER.
        assertEq(IERC20(wrappedAddr).balanceOf(USER), ctxValue);
    }

    function test_D1_erc20_toNotOrigin_reverts() public {
        vm.selectFork(crossForkID);
        uint chainX = 90005; // from-chain pair: origin
        uint chainY = 90006; // to-chain pair: NOT origin (wrapped)

        vm.startPrank(CrossOWNER);
        address wrappedAddr = bridgeCross.createToken(chainY, address(0x5555), "WRP3", 18);
        bridgeCross.registerToken(chainX, true, wrappedAddr, address(0x6666));
        vm.stopPrank();

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainY, IERC20(wrappedAddr), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(chainY, USER, value2, fee2, ex2, "");

        // chainX's pair is isOrigin=true, so the hop-1 finalize's `_withdrawToken` will
        // decrement `deposited[chainX][wrappedAddr]` — seed it first. This also funds
        // the bridge with real balance for the (origin) hop-1 finalize's plain transfer.
        _bumpDeposited(chainX, IERC20(wrappedAddr), ctxValue * 2);

        assertTrue(_signAndFinalize(chainX, 1, wrappedAddr, USER, ctxValue, extraData, 5));

        // D1 violation -> bridgeTokenForwarded reverts -> normal finalize fallback pays USER.
        assertEq(IERC20(wrappedAddr).balanceOf(USER), ctxValue);
    }

    function test_D1_native_bypassed_success() public {
        // Native is registered isOrigin=false relative to BSC on CROSS (§7.3's
        // generalization rationale), yet pass-through succeeds unconditionally because
        // D1 bypasses native regardless of isOrigin.
        vm.selectFork(crossForkID);
        assertFalse(bridgeCross.getTokenPair(BSC_CHAIN_ID, Const.NATIVE_TOKEN).isOrigin);

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        _stageNativeForward(value2, fee2, ex2, "");

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "native forward should have succeeded, not fallen back");
    }

    /// @notice L-1 fix (prior review): the ORIGINAL version of this test measured
    /// `minted[BSC_CHAIN_ID][NATIVE_TOKEN]` at test start, but that pair's baseline
    /// already carries `CROSS_FOUNDATION_INITIAL_SUPPLY` (50,000,000 ether) from fixture
    /// setup — so `value2 = mintedAvailable + 1 ether` was itself astronomically large,
    /// and the observed fallback could not be confidently attributed to the intended
    /// `minted >= value` gate specifically (some other gate tripping on such a huge
    /// value would have looked identical). Fixed by using a FRESH synthetic destination
    /// chain with an explicitly pinned, small `minted` baseline, and by asserting the
    /// specific failure code (`ExecutorCallReverted`, meaning the revert happened INSIDE
    /// `bridgeTokenForwarded` itself — exactly what `_checkInitiateAmount`'s
    /// `require(minted >= value)` produces) rather than only the fallback payout.
    function test_D1_native_insufficientMinted_fallsBack() public {
        vm.selectFork(crossForkID);
        uint chainZ = 90007;
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(chainZ, false, address(NATIVE_TOKEN), address(0x7777));

        // Seed a SMALL, explicitly-pinned `minted[chainZ][NATIVE]` baseline via one
        // ordinary (non-forward) finalize FROM chainZ, isolated from every other chain's
        // (especially BSC_CHAIN_ID's foundation-supply-inflated) minted counter.
        uint seedAmount = 5 ether;
        assertTrue(_signAndFinalize(chainZ, 1, address(NATIVE_TOKEN), USER, seedAmount, "", 5));
        uint mintedAvailable = bridgeCross.getTokenPair(chainZ, Const.NATIVE_TOKEN).minted;
        assertEq(mintedAvailable, seedAmount, "pinned pre-forward minted[toChainID] baseline");

        uint value2 = mintedAvailable + 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainZ, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(chainZ, USER, value2, fee2, ex2, "");

        uint userBalanceBefore = USER.balance;
        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(
            USER.balance,
            userBalanceBefore + ctxValue,
            "insufficient minted[toChainID] must fall back to USER, not under-collateralize"
        );

        (bool found,,, ForwardLib.ForwardFailureCode code,) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(
            code == ForwardLib.ForwardFailureCode.ExecutorCallReverted,
            "insufficient minted must fail INSIDE bridgeTokenForwarded (ExecutorCallReverted), pinning the cause to the minted>=value gate specifically"
        );
    }

    // ----------------------------------------------------------------
    // L-2: native D1 — all four isOrigin[from] x isOrigin[to] combinations
    // ----------------------------------------------------------------

    /// @dev Registers a fresh (native) token pair for `chainID` with the given
    /// `isOrigin` flag and a throwaway remote-token placeholder address.
    function _registerNativePair(uint chainID, bool isOrigin) internal {
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(chainID, isOrigin, address(NATIVE_TOKEN), address(uint160(0xA000 + chainID)));
    }

    /// @dev Seeds `deposited[chainID][NATIVE]` via a real outbound CROSS -> chainID
    /// native bridge (needed before any finalize FROM an isOrigin=true chain, whose
    /// `_withdrawToken` decrements `deposited` and would otherwise underflow).
    function _bumpDepositedNative(uint chainID, uint amount) internal {
        (, uint fee, uint exFee) = bridgeVerifierCross.calculateFee(chainID, IERC20(Const.NATIVE_TOKEN), amount);
        uint total = amount + fee + exFee;
        vm.deal(USER, total);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken{value: total}(chainID, IERC20(Const.NATIVE_TOKEN), USER, amount, fee, exFee, "");
    }

    /// @dev Seeds `minted[chainID][NATIVE]` via one ordinary (non-forward) finalize FROM
    /// chainID (needed before forwarding TO an isOrigin=false chain, whose
    /// `_checkInitiateAmount` requires `minted >= value`).
    function _bumpMintedNative(uint chainID, uint amount) internal {
        assertTrue(_signAndFinalize(chainID, 1, address(NATIVE_TOKEN), USER, amount, "", 5));
    }

    /// @dev Runs one full isOrigin[from] x isOrigin[to] combination: registers a fresh
    /// `fromChainID`/`toChainID` pair with the given flags, seeds whatever precondition
    /// each flag requires, then asserts the native forward succeeds (D1 bypasses native
    /// unconditionally — spec §14.2) regardless of the combination.
    /// @dev M-2: also snapshots the bridge's real native balance and both pairs'
    /// `minted`/`deposited` around the forward, asserts the EXACT per-pair delta this
    /// specific isOrigin combination must produce (spec §6.1's `_depositToken`/
    /// `_withdrawToken` rules), and asserts the §14.3 accounting invariant tying those
    /// ledger deltas to the bridge's real balance change:
    /// `Δbalance == −Σ(!isOrigin)Δminted + Σ(isOrigin)Δdeposited`. Native pass-through
    /// round-trips through the executor with net zero balance effect on its own; the
    /// only REAL balance movement is hop-2's network/ex fee paid out to `_dev`, and the
    /// invariant below is what ties that observable movement back to the ledger.
    function _runD1NativeCombo(uint fromChainID, bool fromIsOrigin, uint toChainID, bool toIsOrigin) internal {
        vm.selectFork(crossForkID);
        _registerNativePair(fromChainID, fromIsOrigin);
        _registerNativePair(toChainID, toIsOrigin);

        if (fromIsOrigin) _bumpDepositedNative(fromChainID, 100 ether);
        uint value2 = 1 ether;
        if (!toIsOrigin) _bumpMintedNative(toChainID, value2 * 2);

        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(toChainID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(toChainID, USER, value2, fee2, ex2, "");

        // M-2: snapshot immediately before the forward — after the seeding above, whose
        // own balance/ledger effects must not pollute this combination's measured delta.
        uint balBefore = address(bridgeCross).balance;
        IBridgeRegistry.TokenPair memory fromPairBefore = bridgeCross.getTokenPair(fromChainID, Const.NATIVE_TOKEN);
        IBridgeRegistry.TokenPair memory toPairBefore = bridgeCross.getTokenPair(toChainID, Const.NATIVE_TOKEN);

        uint userBalanceBefore = USER.balance;
        assertTrue(_signAndFinalize(fromChainID, 1, address(NATIVE_TOKEN), USER, ctxValue, extraData, 5));

        assertEq(
            USER.balance,
            userBalanceBefore,
            "native forward must succeed regardless of isOrigin[from]/isOrigin[to] combination, not fall back"
        );

        // M-2: exact per-pair ledger deltas for this specific isOrigin combination.
        IBridgeRegistry.TokenPair memory fromPairAfter = bridgeCross.getTokenPair(fromChainID, Const.NATIVE_TOKEN);
        IBridgeRegistry.TokenPair memory toPairAfter = bridgeCross.getTokenPair(toChainID, Const.NATIVE_TOKEN);

        // Σ(!isOrigin)Δminted and Σ(isOrigin)Δdeposited, accumulated across BOTH pairs —
        // exactly the two sums the §14.3 invariant below is stated in terms of.
        int mintedDeltaSum;
        int depositedDeltaSum;

        if (fromIsOrigin) {
            assertEq(
                fromPairAfter.deposited,
                fromPairBefore.deposited - ctxValue,
                "from-pair (isOrigin) deposited must decrease by ctxValue (hop-1's own _withdrawToken)"
            );
            assertEq(fromPairAfter.minted, fromPairBefore.minted, "from-pair (isOrigin) minted must be unchanged");
            depositedDeltaSum -= int(ctxValue);
        } else {
            assertEq(
                fromPairAfter.minted,
                fromPairBefore.minted + ctxValue,
                "from-pair (wrapped) minted must increase by ctxValue (hop-1's own _withdrawToken)"
            );
            assertEq(fromPairAfter.deposited, fromPairBefore.deposited, "from-pair (wrapped) deposited must be unchanged");
            mintedDeltaSum += int(ctxValue);
        }

        if (toIsOrigin) {
            assertEq(
                toPairAfter.deposited,
                toPairBefore.deposited + value2,
                "to-pair (isOrigin) deposited must increase by value2 (hop-2's own _depositToken)"
            );
            assertEq(toPairAfter.minted, toPairBefore.minted, "to-pair (isOrigin) minted must be unchanged");
            depositedDeltaSum += int(value2);
        } else {
            assertEq(
                toPairAfter.minted,
                toPairBefore.minted - value2,
                "to-pair (wrapped) minted must decrease by value2 (hop-2's own _depositToken)"
            );
            assertEq(toPairAfter.deposited, toPairBefore.deposited, "to-pair (wrapped) deposited must be unchanged");
            mintedDeltaSum -= int(value2);
        }

        // M-2 / spec §14.3 invariant.
        int balDelta = int(address(bridgeCross).balance) - int(balBefore);
        int expectedBalDelta = -mintedDeltaSum + depositedDeltaSum;
        assertEq(
            balDelta,
            expectedBalDelta,
            "Sec 14.3 invariant: bridge balance delta must equal -sum(!isOrigin minted delta) + sum(isOrigin deposited delta)"
        );
    }

    /// @notice isOrigin[from]=true, isOrigin[to]=true.
    function test_D1_native_isOriginCombination_TT_succeeds() public {
        _runD1NativeCombo(90101, true, 90102, true);
    }

    /// @notice isOrigin[from]=true, isOrigin[to]=false.
    function test_D1_native_isOriginCombination_TF_succeeds() public {
        _runD1NativeCombo(90103, true, 90104, false);
    }

    /// @notice isOrigin[from]=false, isOrigin[to]=true.
    function test_D1_native_isOriginCombination_FT_succeeds() public {
        _runD1NativeCombo(90105, false, 90106, true);
    }

    /// @notice isOrigin[from]=false, isOrigin[to]=false.
    function test_D1_native_isOriginCombination_FF_succeeds() public {
        _runD1NativeCombo(90107, false, 90108, false);
    }

    // ----------------------------------------------------------------
    // §14.6 Code-size gates (deployment blockers if these fail)
    // ----------------------------------------------------------------

    function test_size_crossBridgeV2_underEIP170Limit() public {
        assertLe(address(new CrossBridgeV2()).code.length, 24576);
    }

    function test_size_bscBridge_underEIP170Limit() public {
        vm.selectFork(bscForkID);
        assertLe(address(new BSCBridge()).code.length, 24576);
    }

    function test_size_bscBridgeV2_underEIP170Limit() public {
        vm.selectFork(bscForkID);
        assertLe(address(new BSCBridgeV2()).code.length, 24576);
    }

    // ----------------------------------------------------------------
    // H-1: ForwardLib link-integrity — mislink recovery (plan spec §7.5)
    // ----------------------------------------------------------------

    /// @notice Dedicated link-integrity happy-path smoke test. Distinct in PURPOSE from
    /// `test_D1_native_bypassed_success` (which documents the D1 bypass rule itself):
    /// this one exists purely to anchor "a correctly linked ForwardLib forwards
    /// successfully", as the positive counterpart to the mislink tests below.
    function test_correctlyLinkedLibrary_forwardSucceeds() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        _stageNativeForward(value2, fee2, ex2, "");

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "correctly linked ForwardLib: forward must succeed, not fall back");
    }

    /// @dev Shared body for both mislink variants below: stages a forward, but with
    /// `ForwardLib`'s linked address tampered with FIRST, and asserts the entire
    /// finalize batch reverts — D3's normal per-item fallback does NOT absorb this
    /// (plan spec §7.5, since `_forwardBegin`'s `ForwardLib.setCtx()` call runs before
    /// the executor whitelist/approve gates) — no D3 payout, no finalize-index
    /// progress, i.e. genuinely nothing committed.
    function _assertMislinkRevertsWholeBatch(bytes memory tamperedLibCode) internal {
        vm.selectFork(crossForkID);
        address libAddr = _discoverForwardLibAddress();
        bytes memory originalLibCode = libAddr.code;
        vm.etch(libAddr, tamperedLibCode);

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.selectFork(bscForkID);
        (, uint netA1, uint exA1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValue);
        uint index = nextIndexBSC;
        vm.deal(USER, ctxValue + netA1 + exA1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxValue + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValue, netA1, exA1, extraData
        );
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint userBalanceBefore = USER.balance;
        uint finalizeIndexBefore = bridgeCross.getNextFinalizeIndex(BSC_CHAIN_ID);

        // Inlined (not via `_signAndFinalize`): `vm.expectRevert()` arms for the VERY
        // NEXT external call, and `_signAndFinalize` itself makes an earlier, harmless
        // `domainSeparator()` staticcall before `finalizeBridgeBatch` — which would
        // itself be (wrongly) treated as "the next call" and fail the expectation since
        // it doesn't revert. Compute the signature first, THEN arm `expectRevert()`
        // immediately before the one call that must actually revert.
        bytes32 h = keccak256(
            abi.encode(FINALIZE_TYPEHASH, BSC_CHAIN_ID, index, address(NATIVE_TOKEN), USER, ctxValue, keccak256(extraData))
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), h);
        uint8[] memory v = new uint8[](5);
        bytes32[] memory r = new bytes32[](5);
        bytes32[] memory s = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], hash);
        }
        IBridgeRegistry.FinalizeArguments[] memory args = new IBridgeRegistry.FinalizeArguments[](1);
        args[0] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: index,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxValue,
            extraData: extraData
        });
        uint8[][] memory vArray = new uint8[][](1);
        bytes32[][] memory rArray = new bytes32[][](1);
        bytes32[][] memory sArray = new bytes32[][](1);
        vArray[0] = v;
        rArray[0] = r;
        sArray[0] = s;

        vm.expectRevert();
        bridgeCross.finalizeBridgeBatch(args, vArray, rArray, sArray);

        assertEq(USER.balance, userBalanceBefore, "reverted batch must not pay USER via the normal D3 fallback");
        assertEq(
            bridgeCross.getNextFinalizeIndex(BSC_CHAIN_ID),
            finalizeIndexBefore,
            "reverted batch must not advance the finalize index"
        );

        // Isolation (plan requirement): restore the tampered library's real code within
        // THIS test — forge's own per-test state isolation already prevents this from
        // leaking into other test functions regardless, but restore explicitly anyway.
        vm.etch(libAddr, originalLibCode);
    }

    /// @notice Mislink variant 1: the linked address has no code at all.
    function test_mislinkedLibrary_emptyCode_revertsWholeBatch() public {
        _assertMislinkRevertsWholeBatch("");
    }

    /// @notice Mislink variant 2: the linked address has code, but it unconditionally
    /// reverts (`PUSH0 PUSH0 REVERT`, reverting with empty data).
    function test_mislinkedLibrary_alwaysReverts_revertsWholeBatch() public {
        _assertMislinkRevertsWholeBatch(hex"5f5ffd");
    }

    // ----------------------------------------------------------------
    // M-1: failure-code / event observability contract (plan spec §10)
    // ----------------------------------------------------------------

    function test_forwardFailed_notAttempted_targetNotWhitelisted() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeExecutorCross.removeWhitelistTarget(address(bridgeCrossV2));

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, ctxValue, "not-attempted forward must fall back to USER");

        (bool found,,, ForwardLib.ForwardFailureCode code, bytes32 reasonHash) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(code == ForwardLib.ForwardFailureCode.ForwardNotAttempted);
        assertEq(reasonHash, bytes32(0), "reasonHash must be zero when the executor was never invoked");
    }

    /// @dev NOTE on which "approve failure" shape actually produces `ForwardNotAttempted`:
    /// `_finalizeBridge` captures approve's outcome via a raw low-level `.call(...)`,
    /// whose `ok` reflects whether the CALL ITSELF reverted — NOT the ABI-decoded
    /// boolean return value. An ERC20 whose `approve` merely RETURNS `false` (without
    /// reverting) therefore does NOT hit this path: the low-level call still succeeds
    /// (`fwd.ok = true`), so the executor IS still invoked with a never-actually-granted
    /// allowance, and `executeExtraCall`'s `transferFrom` reverts instead — producing
    /// `ExecutorCallReverted`, not `ForwardNotAttempted`. Only a genuinely REVERTING
    /// `approve` (like `MockRevertingApproveMintableToken` here) makes the low-level
    /// call itself fail and reproduces `ForwardNotAttempted` via this gate. This is
    /// existing, already-reviewed `BaseBridge` behavior (round 1 found zero logic
    /// defects) — out of scope to change this round; this test targets the scenario
    /// that actually reaches `ForwardNotAttempted`, not the plan's literal wording.
    function test_forwardFailed_notAttempted_erc20ApproveFails() public {
        vm.selectFork(crossForkID);
        MockRevertingApproveMintableToken badToken = new MockRevertingApproveMintableToken(address(bridgeCrossV2));
        uint chainA = 90201;
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(chainA, false, address(badToken), address(0x9201));

        uint value2 = 1 ether;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, 0, 0, "");

        vm.recordLogs();
        assertTrue(_signAndFinalize(chainA, 1, address(badToken), USER, value2, extraData, 5));

        assertEq(badToken.balanceOf(USER), value2, "approve-failure forward must fall back to minting USER the token");

        (bool found,,, ForwardLib.ForwardFailureCode code, bytes32 reasonHash) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(code == ForwardLib.ForwardFailureCode.ForwardNotAttempted);
        assertEq(reasonHash, bytes32(0), "reasonHash must be zero when the executor was never invoked");
    }

    function test_forwardFailed_executorCallReverted_feeMismatch() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.recordLogs();
        // Finalize with ctxValue+1 (one wei off) so bridgeTokenForwarded's exact-equality
        // check (value + networkFee + exFee == ctxValue) fails, reverting INSIDE the
        // forward call itself.
        _hop1NativeAndFinalize(ctxValue + 1, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, ctxValue + 1, "fee-mismatch forward must fall back to USER");

        (bool found,,, ForwardLib.ForwardFailureCode code, bytes32 reasonHash) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(code == ForwardLib.ForwardFailureCode.ExecutorCallReverted);
        // L-1: `reasonHash` must be the EXACT hash of the outer executor-frame revert
        // data, not merely nonzero. Here the fee-mismatch causes `bridgeTokenForwarded`
        // ITSELF (the target) to revert; `BridgeExecutor.executeExtraCall` (BridgeExecutor.sol:199)
        // then replaces that with its own `BETargetCallFailed()` before its call frame
        // reverts, so that substituted selector is what the outer frame's revert data
        // actually is.
        assertEq(
            reasonHash,
            keccak256(abi.encodeWithSelector(BridgeExecutor.BETargetCallFailed.selector)),
            "reasonHash must be the exact hash of BridgeExecutor's BETargetCallFailed() substitution"
        );
    }

    function test_forwardFailed_executorCallReverted_methodNotWhitelisted() public {
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeExecutorCross.setMethodCheckEnabled(address(bridgeCrossV2), true);
        // Deliberately do NOT whitelist bridgeTokenForwarded's selector.

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, ctxValue, "method-whitelist rejection must fall back to USER");

        (bool found,,, ForwardLib.ForwardFailureCode code, bytes32 reasonHash) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(
            code == ForwardLib.ForwardFailureCode.ExecutorCallReverted,
            "method-whitelist rejection happens INSIDE executeExtraCall (executor WAS invoked) -> ExecutorCallReverted, not ForwardNotAttempted"
        );
        // L-1: here the EXECUTOR ITSELF reverts (the method-whitelist gate rejects the
        // call before the target is ever reached), so unlike the fee-mismatch case above
        // there is no `BETargetCallFailed()` substitution -- the outer frame's revert
        // data is `BEMethodNotWhitelisted()` directly.
        assertEq(
            reasonHash,
            keccak256(abi.encodeWithSelector(BridgeExecutor.BEMethodNotWhitelisted.selector)),
            "reasonHash must be the exact hash of the executor's own BEMethodNotWhitelisted() revert"
        );
    }

    function test_forwardFailed_callNotConsumed_mockExecutor() public {
        vm.selectFork(crossForkID);
        NonConsumingExecutor evilExec = new NonConsumingExecutor();
        vm.prank(CrossOWNER);
        bridgeCrossV2.setBridgeExecutor(IBridgeExecutor(address(evilExec)));

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(
            USER.balance,
            0,
            "ForwardCallNotConsumed must NOT fall back to a direct USER payout (executor claimed success)"
        );

        (bool found,,, ForwardLib.ForwardFailureCode code,) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(code == ForwardLib.ForwardFailureCode.ForwardCallNotConsumed);
    }

    function test_forwardInitiated_allFieldsMatchArguments() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        bytes memory hop3 = bytes("hop3-marker");
        uint expectedFinalizeIndex = nextIndexBSC;

        vm.recordLogs();
        _stageNativeForward(value2, fee2, ex2, hop3);

        (
            bool found,
            uint fromChainID,
            uint finalizeIndex,
            uint toChainID,
            address fromToken,
            address to,
            uint value,
            uint networkFee,
            uint exFee,
            bytes32 extraDataHash
        ) = _findForwardInitiated(vm.getRecordedLogs());

        assertTrue(found, "ForwardInitiated must be emitted");
        assertEq(fromChainID, BSC_CHAIN_ID);
        assertEq(finalizeIndex, expectedFinalizeIndex);
        assertEq(toChainID, BSC_CHAIN_ID); // hop-2's own destination chain in this helper
        assertEq(fromToken, address(NATIVE_TOKEN));
        assertEq(to, USER);
        assertEq(value, value2);
        assertEq(networkFee, fee2);
        assertEq(exFee, ex2);
        assertEq(extraDataHash, keccak256(hop3), "extraDataHash must hash bridgeTokenForwarded's OWN extraData arg (hop3)");
    }

    /// @notice Regression: a NON-forward extraData target failing must never emit
    /// `ForwardFailed` (only `_forwardBegin`-matched items ever call `_forwardEnd`).
    function test_ordinaryExtraCallFailure_doesNotEmitForwardFailed() public {
        vm.selectFork(crossForkID);
        mockTargetCross.setShouldRevert(true);

        uint amount = 1 ether;
        bytes memory calldata_ =
            abi.encodeWithSelector(MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes(""));
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);
        uint index = nextIndexBSC;
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: amount}(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, amount, 0, 0, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        vm.recordLogs();
        assertTrue(_signAndFinalize(BSC_CHAIN_ID, index, address(NATIVE_TOKEN), USER, amount, extraData, 5));

        assertEq(USER.balance, amount, "ordinary (non-forward) extra-call failure should still fall back normally");
        assertFalse(_hasForwardFailed(vm.getRecordedLogs()), "a NON-forward extraData target must never emit ForwardFailed");
    }

    /// @notice `bridgeTokenForwarded`'s own native fee payout (mid-`_initiateBridge`)
    /// reverting must roll back EVERYTHING hop-2 did (spec §12 atomicity) — no
    /// `BridgeInitiated`, no initiate-index progress — and the ORIGINAL finalize must
    /// still fall back to paying the chain B recipient.
    function test_forwardFailure_rollsBackBridgeInitiated_fallsBackToUser() public {
        vm.selectFork(crossForkID);
        RevertingReceiver evilDev = new RevertingReceiver();
        vm.prank(CrossOWNER);
        bridgeCrossV2.setDev(payable(address(evilDev)));

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(fee2 + ex2 > 0, "test requires a nonzero fee to trigger the _dev callback");

        uint hop2IndexBefore = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(BSC_CHAIN_ID, USER, value2, fee2, ex2, "");

        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, ctxValue, "hop-2 failure must fall back to paying USER on chain B");
        assertEq(
            bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID),
            hop2IndexBefore,
            "hop-2's initiate index must not have advanced (fully rolled back)"
        );

        Vm.Log[] memory logs = vm.getRecordedLogs();
        bytes32 bridgeInitiatedTopic0 = keccak256(
            "BridgeInitiated(uint256,uint256,address,address,address,address,uint256,uint256,uint256,bytes,uint256)"
        );
        for (uint i = 0; i < logs.length; i++) {
            assertFalse(
                logs[i].emitter == address(bridgeCrossV2) && logs[i].topics.length > 0
                    && logs[i].topics[0] == bridgeInitiatedTopic0,
                "BridgeInitiated must NOT have been emitted for the rolled-back hop-2"
            );
        }
    }

    // ----------------------------------------------------------------
    // M-1b: §14.1 named regression cases required to literally claim acceptance
    // criterion 3
    // ----------------------------------------------------------------

    function test_forward_reentry_finalizeBridgeBatch_reverts() public {
        vm.selectFork(crossForkID);
        EvilForwardExecutor evil = new EvilForwardExecutor(address(bridgeCrossV2));
        vm.startPrank(CrossOWNER);
        bridgeCrossV2.setBridgeExecutor(IBridgeExecutor(address(evil)));
        bridgeCrossV2.setDev(payable(address(evil)));
        vm.stopPrank();
        evil.setShouldReenterFinalizeBridgeBatch(true);

        uint value2 = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(networkFee2 + exFee2 > 0, "test requires a nonzero fee to trigger the _dev callback");
        _stageNativeForward(value2, networkFee2, exFee2, "");

        assertTrue(evil.reentryAttempted(), "nested call was never attempted");
        assertFalse(evil.reentrySucceeded(), "nested finalizeBridgeBatch reentry must not succeed");
        assertEq(bytes4(evil.reentryRevertData()), bytes4(keccak256("ReentrancyGuardReentrantCall()")));
    }

    function test_forward_reentry_releasePending_reverts() public {
        vm.selectFork(crossForkID);
        EvilForwardExecutor evil = new EvilForwardExecutor(address(bridgeCrossV2));
        vm.startPrank(CrossOWNER);
        bridgeCrossV2.setBridgeExecutor(IBridgeExecutor(address(evil)));
        bridgeCrossV2.setDev(payable(address(evil)));
        vm.stopPrank();
        evil.setShouldReenterReleasePending(true);

        uint value2 = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        require(networkFee2 + exFee2 > 0, "test requires a nonzero fee to trigger the _dev callback");
        _stageNativeForward(value2, networkFee2, exFee2, "");

        assertTrue(evil.reentryAttempted(), "nested call was never attempted");
        assertFalse(evil.reentrySucceeded(), "nested releasePending reentry must not succeed");
        assertEq(bytes4(evil.reentryRevertData()), bytes4(keccak256("ReentrancyGuardReentrantCall()")));
    }

    /// @notice A malicious/compromised native-recipient `to` (attacker-controlled: `to`
    /// is a caller-supplied address on an ordinary, non-forward finalize) cannot reach
    /// EITHER `bridgeTokenForwarded` (blocked: `msg.sender` is the receiver, not
    /// `bridgeExecutor`) or `bridgeToken` (blocked: the shared `nonReentrant` guard is
    /// already held by the enclosing `finalizeBridgeBatch` call) from its native
    /// receiver callback.
    function test_nativeReceiverCallback_bridgeTokenForwarded_reverts() public {
        vm.selectFork(crossForkID);
        MaliciousNativeReceiver evilReceiver = new MaliciousNativeReceiver(address(bridgeCrossV2), BSC_CHAIN_ID);

        uint amount = 1 ether;
        vm.selectFork(bscForkID);
        vm.deal(USER, amount);
        uint index = nextIndexBSC;
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: amount}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), address(evilReceiver), amount, 0, 0, ""
        );
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        assertTrue(_signAndFinalize(BSC_CHAIN_ID, index, address(NATIVE_TOKEN), address(evilReceiver), amount, "", 5));

        assertTrue(evilReceiver.forwardAttempted(), "bridgeTokenForwarded reentry was never attempted");
        assertFalse(evilReceiver.forwardSucceeded(), "bridgeTokenForwarded reentry must not succeed");
        assertEq(bytes4(evilReceiver.forwardRevertData()), ICrossBridgeV2.BaseBridgeForwardNotExecutor.selector);

        assertTrue(evilReceiver.bridgeTokenAttempted(), "bridgeToken reentry was never attempted");
        assertFalse(evilReceiver.bridgeTokenSucceeded(), "bridgeToken reentry must not succeed (existing guard unchanged)");
        assertEq(bytes4(evilReceiver.bridgeTokenRevertData()), bytes4(keccak256("ReentrancyGuardReentrantCall()")));
    }

    /// @dev Drives one ordinary finalize whose extraData targets `bridgeCrossV2` itself
    /// with `selector`, and asserts the call falls back to a normal USER payout — i.e.
    /// the executor-invoked call reverted rather than succeeding.
    function _assertAdminSelectorBlockedViaExecutor(bytes4 selector, bytes memory args) internal {
        bytes memory extraData = abi.encodePacked(address(bridgeCrossV2), selector, args);

        uint amount = 1 ether;
        vm.selectFork(bscForkID);
        vm.deal(USER, amount);
        uint index = nextIndexBSC;
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: amount}(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, amount, 0, 0, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint balBefore = USER.balance;
        assertTrue(_signAndFinalize(BSC_CHAIN_ID, index, address(NATIVE_TOKEN), USER, amount, extraData, 5));
        assertEq(
            USER.balance,
            balBefore + amount,
            "admin/guarded selector reached via the executor must revert and fall back to USER"
        );
    }

    /// @notice Even with `methodCheckEnabled == false` for the bridge as an executor
    /// target (the executor's OWN selector whitelist providing no protection), the
    /// bridge's PRE-EXISTING role checks and shared reentrancy guard independently
    /// block `setDev`/`bridgeToken`/`releasePending` when reached via the executor
    /// (plan spec §11's on-chain defense claim, §14.1).
    function test_methodCheckDisabled_adminSelectors_revert() public {
        vm.selectFork(crossForkID);
        assertFalse(bridgeExecutorCross.isMethodCheckEnabled(address(bridgeCrossV2)));

        address devBefore = bridgeCross.dev();

        _assertAdminSelectorBlockedViaExecutor(bridgeCrossV2.setDev.selector, abi.encode(address(0xdead)));
        _assertAdminSelectorBlockedViaExecutor(
            bridgeCrossV2.bridgeToken.selector,
            abi.encode(BSC_CHAIN_ID, address(NATIVE_TOKEN), USER, uint(1), uint(0), uint(0), bytes(""))
        );
        _assertAdminSelectorBlockedViaExecutor(bridgeCrossV2.releasePending.selector, abi.encode(BSC_CHAIN_ID, uint(1)));

        assertEq(bridgeCross.dev(), devBefore, "setDev must not have taken effect via the executor path");
    }

    /// @notice `_validateToken`'s visibility change (`private` -> `internal`, needed so
    /// `bridgeTokenForwarded` can call it directly without a modifier) must not have
    /// altered its behavior for its EXISTING callers. `permitBridgeTokenBatch`'s
    /// `onlyValidToken` modifier calls this SAME shared internal function — a pure
    /// visibility change cannot make that call path diverge, so covering `bridgeToken`'s
    /// three error cases here is a complete regression demonstration by construction.
    function test_validateToken_errorsPreserved() public {
        vm.selectFork(crossForkID);
        vm.deal(USER, 10 ether);

        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(BridgeRegistry.RegistryNotExistToken.selector, address(0xBEEF)));
        bridgeCrossV2.bridgeToken{value: 1 ether}(BSC_CHAIN_ID, IERC20(address(0xBEEF)), USER, 1 ether, 0, 0, "");

        vm.prank(CrossOWNER);
        bridgeCross.setChainPause(BSC_CHAIN_ID, true);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(BridgeRegistry.RegistryChainPaused.selector, BSC_CHAIN_ID));
        bridgeCrossV2.bridgeToken{value: 1 ether}(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, 1 ether, 0, 0, "");
        vm.prank(CrossOWNER);
        bridgeCross.setChainPause(BSC_CHAIN_ID, false);

        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), true, true);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(BridgeRegistry.RegistryTokenPaused.selector, address(NATIVE_TOKEN)));
        bridgeCrossV2.bridgeToken{value: 1 ether}(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, 1 ether, 0, 0, "");
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(NATIVE_TOKEN), false, false);
    }
}
