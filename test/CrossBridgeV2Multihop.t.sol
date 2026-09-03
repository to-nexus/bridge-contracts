// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {BSCBridge} from "../src/BSCBridge.sol";
import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {PriceFeed} from "../src/PriceFeed.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";

import {Const} from "../src/lib/Const.sol";
import {ForwardLib} from "../src/lib/ForwardLib.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";

import {CrossMintableERC20Code} from "../src/token/CrossMintableERC20Code.sol";
import {MockTargetContract} from "./BridgeExecutor.t.sol";
import {CrossBridgeV2ForwardTest} from "./CrossBridgeV2Forward.t.sol";

import {TestToken} from "./token/TestToken.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Vm} from "forge-std/Vm.sol";

/**
 * @title CrossBridgeV2MultihopTest
 * @notice M-2 (plan spec §14.5) + M-3 (plan spec §14.3/§6.1): genuine three-bridge
 * A -> B -> C end-to-end coverage and `crossSupply()` pass-through accounting.
 * @dev "Chain C" is a THIRD real, independently deployed `BSCBridge` instance (its own
 * `BridgeVerifier`/`PriceFeed`/`BridgeExecutor`/target contract) — not a mock, not a
 * fabricated signed message. It is deployed on the SAME underlying fork as `crossForkID`
 * (a separate CONTRACT instance representing a separate real chain, exactly as
 * `bridgeBSC` and `bridgeCross` already coexist as two independent contracts before this
 * round). "Chain A" for the native E2E test is the fixture's real `bridgeBSC`, driven
 * through its genuine `bridgeToken` entrypoint (`_hop1NativeAndFinalize`, inherited).
 * For the ERC20-both-origin E2E, chain A is represented the same way the pre-existing,
 * already-reviewed D1 tests represent an origin-chain message (a validly signed
 * `_signAndFinalize` call) — hub-origin assets have no real chain-A contract to
 * originate from in this fixture (BSC's own registered tokens are not hub-origin), and a
 * validator-signed finalize IS the system's actual trust mechanism for "chain A said
 * so"; hop-2's finalize on chain C is what matters for genuineness and IS fully real.
 */
contract CrossBridgeV2MultihopTest is CrossBridgeV2ForwardTest {
    uint internal constant CHAIN_C_ID = 90999;
    uint internal constant CHAIN_A_ID = 90501;

    address internal chainCOwner;
    BSCBridge internal bridgeChainC;
    BridgeVerifier internal bridgeVerifierChainC;
    PriceFeed internal priceFeedChainC;
    BridgeExecutor internal bridgeExecutorChainC;
    MockTargetContract internal mockTargetChainC;
    ICrossMintableERC20Code internal crossMintableERC20CodeChainC;
    TestToken internal chainCLocalCrossToken;

    // H-1: a REAL, independently deployed chain-A leaf bridge (mirrors `bridgeChainC`'s
    // deployment pattern) holding a WRAPPED representation of the hub-origin ERC20, so
    // the ERC20 A -> B -> C E2E test can genuinely originate hop-1 on chain A instead of
    // faking it via a signed message.
    address internal chainAOwner;
    BSCBridge internal bridgeChainA;
    BridgeVerifier internal bridgeVerifierChainA;
    PriceFeed internal priceFeedChainA;
    ICrossMintableERC20Code internal crossMintableERC20CodeChainA;
    TestToken internal chainALocalCrossToken;

    function setUp() public virtual override {
        super.setUp();

        vm.selectFork(crossForkID);
        chainCOwner = makeAddr("chainC-owner");

        // Deploy chain C's bridge exactly like a real leaf bridge (mirrors
        // `BSCTest.setUp()`'s own BSCBridge deployment pattern), registered symmetrically
        // to CROSS_CHAIN_ID.
        chainCLocalCrossToken = new TestToken("Chain C Cross", "CCROSS", 18);
        BSCBridge bridgeChainCImpl = new BSCBridge();
        ERC1967Proxy bridgeChainCProxy = new ERC1967Proxy(address(bridgeChainCImpl), bytes(""));
        bridgeChainC = BSCBridge(payable(address(bridgeChainCProxy)));
        bridgeChainC.initializeBSCBridge(
            chainCOwner, payable(chainCOwner), threshold, CROSS_CHAIN_ID, address(chainCLocalCrossToken), 0
        );

        vm.startPrank(chainCOwner);
        bridgeChainC.grantRole(Const.ADMIN_ROLE, chainCOwner);
        bridgeChainC.grantRole(Const.EDITOR_ROLE, chainCOwner);
        bridgeChainC.grantRole(Const.OPERATOR_ROLE, chainCOwner);
        bytes32[] memory roles = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            roles[i] = Const.VALIDATOR_ROLE;
        }
        bridgeChainC.grantRoleBatch(roles, VALIDATORS);

        // Fee table: deliberately minimal (zero default price/fee/thresholds, no
        // per-chain gas price configured) — every fee/minimum resolves to 0 except the
        // 1-token-unit minimum floor, keeping this leg's arithmetic simple and
        // orthogonal to the (already-covered-elsewhere) fee-calculation logic itself.
        PriceFeed priceFeedChainCImpl = new PriceFeed();
        ERC1967Proxy priceFeedChainCProxy = new ERC1967Proxy(address(priceFeedChainCImpl), bytes(""));
        priceFeedChainC = PriceFeed(address(priceFeedChainCProxy));
        priceFeedChainC.initialize(chainCOwner, DOLLAR_DECIMALS);
        priceFeedChainC.grantRole(Const.PRICER_ROLE, chainCOwner);

        bridgeVerifierChainC = new BridgeVerifier(
            chainCOwner, address(bridgeChainC), address(priceFeedChainC), 200_000, 0, 0, 0, 0, 0, 2 hours
        );
        bridgeChainC.setBridgeVerifier(bridgeVerifierChainC);

        crossMintableERC20CodeChainC =
            ICrossMintableERC20Code(address(new CrossMintableERC20Code(address(bridgeChainC))));
        bridgeChainC.setCrossMintableERC20Code(crossMintableERC20CodeChainC);

        // Native pair: chain C's own native coin <-> CROSS's native coin. Both
        // non-origin (mirrors how CROSS/BSC each treat their OWN native relative to the
        // counterpart chain — native never mints/burns regardless of this flag; it only
        // controls deposited/minted BOOKKEEPING, spec §6.1).
        bridgeChainC.registerToken(CROSS_CHAIN_ID, false, address(NATIVE_TOKEN), address(NATIVE_TOKEN));
        vm.deal(address(bridgeChainC), 1_000 ether);

        bridgeExecutorChainC = new BridgeExecutor(chainCOwner, address(bridgeChainC));
        bridgeChainC.setBridgeExecutor(IBridgeExecutor(address(bridgeExecutorChainC)));
        mockTargetChainC = new MockTargetContract();
        bridgeExecutorChainC.addWhitelistTarget(address(mockTargetChainC));
        vm.stopPrank();

        // CROSS's own registration for CHAIN_C_ID's native pair: isOrigin=true, so
        // `_checkInitiateAmount`'s `minted >= value` precondition (which only applies to
        // NOT-origin pairs) never gates outbound forwards to chain C — no pre-seeding
        // needed. (Native pass-through itself is unconditional either way per D1.)
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(CHAIN_C_ID, true, address(NATIVE_TOKEN), address(NATIVE_TOKEN));

        // H-1: deploy chain A's bridge exactly like a real leaf bridge (same pattern as
        // chain C above), registered symmetrically to CROSS_CHAIN_ID. No native pair, no
        // executor/target: this fixture only needs chain A to originate a REAL ERC20
        // `bridgeToken` call whose wrapped representation of the hub-origin ERC20 is
        // created per-test via `createToken` (the hub-origin ERC20 itself is deployed
        // fresh inside `test_e2e_erc20_bothOrigin_ABC`, not here).
        chainAOwner = makeAddr("chainA-owner");
        chainALocalCrossToken = new TestToken("Chain A Cross", "ACROSS", 18);
        BSCBridge bridgeChainAImpl = new BSCBridge();
        ERC1967Proxy bridgeChainAProxy = new ERC1967Proxy(address(bridgeChainAImpl), bytes(""));
        bridgeChainA = BSCBridge(payable(address(bridgeChainAProxy)));
        bridgeChainA.initializeBSCBridge(
            chainAOwner, payable(chainAOwner), threshold, CROSS_CHAIN_ID, address(chainALocalCrossToken), 0
        );

        vm.startPrank(chainAOwner);
        bridgeChainA.grantRole(Const.ADMIN_ROLE, chainAOwner);
        bridgeChainA.grantRole(Const.EDITOR_ROLE, chainAOwner);
        bridgeChainA.grantRole(Const.OPERATOR_ROLE, chainAOwner);
        bytes32[] memory chainARoles = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            chainARoles[i] = Const.VALIDATOR_ROLE;
        }
        bridgeChainA.grantRoleBatch(chainARoles, VALIDATORS);

        // Same deliberately-minimal fee table as chain C's PriceFeed/BridgeVerifier.
        PriceFeed priceFeedChainAImpl = new PriceFeed();
        ERC1967Proxy priceFeedChainAProxy = new ERC1967Proxy(address(priceFeedChainAImpl), bytes(""));
        priceFeedChainA = PriceFeed(address(priceFeedChainAProxy));
        priceFeedChainA.initialize(chainAOwner, DOLLAR_DECIMALS);
        priceFeedChainA.grantRole(Const.PRICER_ROLE, chainAOwner);

        bridgeVerifierChainA = new BridgeVerifier(
            chainAOwner, address(bridgeChainA), address(priceFeedChainA), 200_000, 0, 0, 0, 0, 0, 2 hours
        );
        bridgeChainA.setBridgeVerifier(bridgeVerifierChainA);

        crossMintableERC20CodeChainA =
            ICrossMintableERC20Code(address(new CrossMintableERC20Code(address(bridgeChainA))));
        bridgeChainA.setCrossMintableERC20Code(crossMintableERC20CodeChainA);
        vm.stopPrank();
    }

    // ----------------------------------------------------------------
    // Helpers
    // ----------------------------------------------------------------

    /// @dev Generic sign+finalize against `bridgeChainC` (mirrors
    /// `CrossBridgeV2ForwardTest._signAndFinalize`, which is hardcoded to `bridgeCross`).
    function _signAndFinalizeChainC(
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

        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, keccak256(extraData)));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeChainC.domainSeparator(), h);

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
        ok = bridgeChainC.finalizeBridgeBatch(args, vArray, rArray, sArray);
    }

    /// @dev Builds hop-3's extraData (target = `mockTargetChainC`, a genuine callback
    /// that actually consumes the forwarded native value on chain C).
    function _hop3TargetExtraData(uint value3) internal view returns (bytes memory) {
        bytes memory call_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, value3, bytes("hop3")
        );
        return abi.encodePacked(address(mockTargetChainC), call_);
    }

    /// @dev M-1: builds hop-2's extraData for the ERC20 A -> B -> C E2E (target =
    /// `mockTargetChainC`), using `handleBridgeCallbackWithReturn` rather than the
    /// native test's plain `handleBridgeCallback` so the returned `bytes32` hash gives
    /// an independent, deterministic proof (surfaced via `ExtraCallExecuted`'s
    /// `returnData`) that the target itself actually processed this exact
    /// (token, user, value, data) call -- not just that its ERC20 balance moved.
    function _hop2TargetExtraDataERC20(address token3, uint value3, bytes memory tag)
        internal
        view
        returns (bytes memory)
    {
        bytes memory call_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallbackWithReturn.selector, token3, USER, value3, tag
        );
        return abi.encodePacked(address(mockTargetChainC), call_);
    }

    /// @dev Expected `handleBridgeCallbackWithReturn` return value for the given args,
    /// mirroring the mock's own `keccak256(abi.encode(token, user, amount, data))`.
    function _expectedTargetReturnHash(address token3, uint value3, bytes memory tag) internal view returns (bytes32) {
        return keccak256(abi.encode(token3, USER, value3, tag));
    }

    /// @dev Generic sign+finalize against `bridgeChainA` (H-1: mirrors
    /// `_signAndFinalizeChainC`, used only to finalize the bootstrap leg's inbound
    /// mint on chain A — hop-1 of the measured route itself is a genuine `bridgeToken`
    /// call, not a signed finalize).
    function _signAndFinalizeChainA(
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

        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, keccak256(extraData)));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeChainA.domainSeparator(), h);

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
        ok = bridgeChainA.finalizeBridgeBatch(args, vArray, rArray, sArray);
    }

    // ----------------------------------------------------------------
    // H-1: BridgeInitiated / BridgeFinalized event-log decoding helpers (generic over
    // the emitting bridge, unlike `CrossBridgeV2ForwardTest`'s ForwardLib-specific ones)
    // ----------------------------------------------------------------

    bytes32 internal constant BRIDGE_INITIATED_TOPIC0 = keccak256(
        "BridgeInitiated(uint256,uint256,address,address,address,address,uint256,uint256,uint256,bytes,uint256)"
    );
    bytes32 internal constant BRIDGE_FINALIZED_TOPIC0 =
        keccak256("BridgeFinalized(uint256,uint256,address,address,uint256,uint256)");
    bytes32 internal constant EXTRA_CALL_EXECUTED_TOPIC0 =
        keccak256("ExtraCallExecuted(uint256,uint256,address,bytes4,bool,uint256,bytes)");

    /// @dev Finds the (first) `BridgeInitiated` log emitted BY `emitter` among `logs`,
    /// and decodes its fields. Used to capture hop-1's REAL, on-chain arguments off of
    /// chain A's `bridgeToken` call rather than hardcoding them.
    function _findBridgeInitiated(Vm.Log[] memory logs, address emitter)
        internal
        pure
        returns (
            bool found,
            uint toChainID,
            uint index,
            address fromToken,
            address toToken,
            address from,
            address to,
            uint value,
            uint networkFee,
            uint exFee,
            bytes memory extraData
        )
    {
        for (uint i = 0; i < logs.length; i++) {
            if (
                logs[i].emitter == emitter && logs[i].topics.length == 4 && logs[i].topics[0] == BRIDGE_INITIATED_TOPIC0
            ) {
                toChainID = uint(logs[i].topics[1]);
                index = uint(logs[i].topics[2]);
                from = address(uint160(uint(logs[i].topics[3])));
                (fromToken, toToken, to, value, networkFee, exFee, extraData,) =
                    abi.decode(logs[i].data, (address, address, address, uint, uint, uint, bytes, uint));
                return (true, toChainID, index, fromToken, toToken, from, to, value, networkFee, exFee, extraData);
            }
        }
    }

    /// @dev Finds the (first) `BridgeFinalized` log emitted BY `emitter` among `logs`.
    function _findBridgeFinalized(Vm.Log[] memory logs, address emitter)
        internal
        pure
        returns (bool found, uint fromChainID, uint index, address toToken, address to, uint value)
    {
        for (uint i = 0; i < logs.length; i++) {
            if (
                logs[i].emitter == emitter && logs[i].topics.length == 4 && logs[i].topics[0] == BRIDGE_FINALIZED_TOPIC0
            ) {
                fromChainID = uint(logs[i].topics[1]);
                index = uint(logs[i].topics[2]);
                toToken = address(uint160(uint(logs[i].topics[3])));
                (to, value,) = abi.decode(logs[i].data, (address, uint, uint));
                return (true, fromChainID, index, toToken, to, value);
            }
        }
    }

    /// @dev M-1: finds the (first) `ExtraCallExecuted` log emitted BY `emitter` among
    /// `logs`, and decodes its fields -- used to independently confirm chain C's
    /// `BridgeExecutor`-mediated target call actually ran (as opposed to merely
    /// observing balance movement).
    function _findExtraCallExecuted(Vm.Log[] memory logs, address emitter)
        internal
        pure
        returns (
            bool found,
            address targetContract,
            bytes4 methodID,
            bool success,
            uint consumed,
            bytes memory returnData
        )
    {
        for (uint i = 0; i < logs.length; i++) {
            if (
                logs[i].emitter == emitter && logs[i].topics.length == 4
                    && logs[i].topics[0] == EXTRA_CALL_EXECUTED_TOPIC0
            ) {
                targetContract = address(uint160(uint(logs[i].topics[3])));
                (methodID, success, consumed, returnData) = abi.decode(logs[i].data, (bytes4, bool, uint, bytes));
                return (true, targetContract, methodID, success, consumed, returnData);
            }
        }
    }

    // ----------------------------------------------------------------
    // M-2: genuine A -> B -> C end-to-end
    // ----------------------------------------------------------------

    /// @notice Native A(BSC) -> B(CROSS, forwards) -> C(bridgeChainC, finalizes AND
    /// calls a real target contract). Every hop is a REAL bridge contract call; nothing
    /// here is a mock finalize.
    function test_e2e_native_ABC() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;

        bytes memory hop3ExtraData = _hop3TargetExtraData(value2);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop3ExtraData);

        uint hop2Index = bridgeCross.getNextInitiateIndex(CHAIN_C_ID);
        uint depositedCBefore = bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited;
        uint targetBalBefore = address(mockTargetChainC).balance;

        // hop-1 (real): BSC -> CROSS, staging + completing the forward to chain C.
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "hop-1's forward must succeed on chain B (no D3 fallback)");
        assertEq(
            bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited,
            depositedCBefore + value2,
            "chain B's ledger for the CHAIN_C_ID pair must record exactly the forwarded principal"
        );

        // hop-2 (real): finalize on chain C, which calls the real target contract.
        assertTrue(
            _signAndFinalizeChainC(CROSS_CHAIN_ID, hop2Index, address(NATIVE_TOKEN), USER, value2, hop3ExtraData, 5)
        );

        assertEq(
            address(mockTargetChainC).balance,
            targetBalBefore + value2,
            "chain C's real target contract must have received and consumed the forwarded value"
        );
        assertEq(USER.balance, 0, "chain C target call succeeded: no fallback payout to USER on chain C");
    }

    /// @notice Same A -> B -> C route, but chain C's target reverts — chain C's OWN
    /// EXISTING (unmodified) fallback must pay USER directly there, exactly like any
    /// ordinary (non-forward) extra-call failure would.
    function test_e2e_chainC_targetRevert_fallback() public {
        vm.selectFork(crossForkID);
        mockTargetChainC.setShouldRevert(true);

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory hop3ExtraData = _hop3TargetExtraData(value2);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop3ExtraData);
        uint hop2Index = bridgeCross.getNextInitiateIndex(CHAIN_C_ID);

        _hop1NativeAndFinalize(ctxValue, extraData);
        vm.selectFork(crossForkID);
        assertEq(
            USER.balance, 0, "hop-1 forward must still succeed (chain C's target failure is a LATER, separate hop)"
        );

        uint targetBalBefore = address(mockTargetChainC).balance;
        assertTrue(
            _signAndFinalizeChainC(CROSS_CHAIN_ID, hop2Index, address(NATIVE_TOKEN), USER, value2, hop3ExtraData, 5)
        );

        assertEq(address(mockTargetChainC).balance, targetBalBefore, "reverting target must not receive/keep funds");
        assertEq(USER.balance, value2, "chain C's own fallback must pay USER directly when its target reverts");
    }

    /// @notice ERC20 both-origin A -> B -> C: the asset originates ON THE HUB (CROSS)
    /// itself (spec §6.1's "both isOrigin=true" pattern — see D1's existing
    /// `test_D1_erc20_bothOrigin_success` for the two-chain version this generalizes).
    /// @dev H-1 fix: hop-1 is now a GENUINE `bridgeToken` call on a real chain-A leaf
    /// bridge (`bridgeChainA`), holding a real wrapped representation of the hub-origin
    /// ERC20 — not a fabricated signed message. Since registration alone seeds no
    /// inventory, a REAL bootstrap leg (hub -> chain A, finalized there) runs FIRST to
    /// mint USER the wrapped token and seed both ledgers; only then does the measured
    /// A -> B -> C route run, starting with chain A's own `bridgeToken` call. Hop-1's
    /// arguments for the hub's finalize are read off the REAL `BridgeInitiated` event
    /// chain A emits (never hardcoded), following the exact field-mapping the plan
    /// pins down: the hub's `toToken` is the event's `toToken` (== `remoteToken`, the
    /// HUB's own local token address) — NOT the event's `fromToken` (chain A's own
    /// wrapped address). Chain C's leg remains a REAL finalize on a real bridge holding
    /// a real wrapped token, as before.
    function test_e2e_erc20_bothOrigin_ABC() public {
        vm.selectFork(crossForkID);
        TestToken hubToken = new TestToken("Hub Origin", "HUB", 18);
        hubToken.mint(address(bridgeCross), 1000 ether);

        vm.startPrank(CrossOWNER);
        bridgeCross.registerToken(CHAIN_A_ID, true, address(hubToken), address(0x9501));
        bridgeCross.registerToken(CHAIN_C_ID, true, address(hubToken), address(0x9502));
        vm.stopPrank();

        // Chain A and chain C each need their OWN local (wrapped) representation of
        // hubToken to actually mint/burn against real inventory there.
        vm.prank(chainAOwner);
        address chainAHubToken = bridgeChainA.createToken(CROSS_CHAIN_ID, address(hubToken), "AHUB", 18);
        vm.prank(chainCOwner);
        address chainCHubToken = bridgeChainC.createToken(CROSS_CHAIN_ID, address(hubToken), "CHUB", 18);

        // ------------------------------------------------------------------
        // STEP 0 (H-1 CRITICAL PREREQUISITE): bootstrap leg hub -> chain A. Registration
        // alone seeds no inventory: USER holds no wrapped token on chain A, chain A's
        // `minted[hub]` is 0, and the hub's `deposited[chainA]` is 0, so the measured
        // route below would have nothing real to originate from without this. A REAL
        // `bridgeToken` call on the hub, finalized on chain A, seeds BOTH ledgers.
        // ------------------------------------------------------------------
        uint bootstrapAmount = 100 ether;
        (, uint bootFee, uint bootEx) = bridgeVerifierCross.calculateFee(CHAIN_A_ID, IERC20(hubToken), bootstrapAmount);
        uint bootTotal = bootstrapAmount + bootFee + bootEx;

        uint hubDepositedABeforeBootstrap = bridgeCross.getTokenPair(CHAIN_A_ID, address(hubToken)).deposited;
        uint chainAMintedBeforeBootstrap = bridgeChainA.getTokenPair(CROSS_CHAIN_ID, chainAHubToken).minted;

        hubToken.mint(USER, bootTotal);
        vm.prank(USER);
        hubToken.approve(address(bridgeCrossV2), bootTotal);
        uint bootIndex = bridgeCross.getNextInitiateIndex(CHAIN_A_ID);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken(CHAIN_A_ID, IERC20(hubToken), USER, bootstrapAmount, bootFee, bootEx, "");

        assertEq(
            bridgeCross.getTokenPair(CHAIN_A_ID, address(hubToken)).deposited,
            hubDepositedABeforeBootstrap + bootstrapAmount,
            "bootstrap leg must seed hub deposited[chainA]"
        );

        assertTrue(
            _signAndFinalizeChainA(CROSS_CHAIN_ID, bootIndex, chainAHubToken, USER, bootstrapAmount, "", 5),
            "bootstrap leg finalize on chain A must succeed"
        );

        assertEq(
            bridgeChainA.getTokenPair(CROSS_CHAIN_ID, chainAHubToken).minted,
            chainAMintedBeforeBootstrap + bootstrapAmount,
            "bootstrap leg must seed chain A minted[hub]"
        );
        assertEq(
            IERC20(chainAHubToken).balanceOf(USER),
            bootstrapAmount,
            "bootstrap leg must mint the wrapped token to USER on chain A"
        );

        // ------------------------------------------------------------------
        // STEP 1 (measured route): real chain A -> hub -> chain C, starting with a REAL
        // `bridgeToken` call on chain A against the wrapped inventory just bootstrapped.
        // ------------------------------------------------------------------
        uint value2 = 10 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(hubToken), value2);
        uint ctxValue = value2 + fee2 + ex2;
        require(ctxValue <= bootstrapAmount, "bootstrap inventory must cover the measured route's hop-1 principal");

        // M-1: hop-2's own extraData now carries a REAL chain-C target call (mirrors
        // the native E2E's `_hop3TargetExtraData`), so this route proves through to
        // "chain A `bridgeToken` -> ... -> chain C target method call", not merely
        // destination delivery. `mockTargetChainC` is already whitelisted on
        // `bridgeExecutorChainC` in `setUp()` (shared with the native E2E).
        bytes memory targetTag = bytes("hop2-erc20");
        bytes memory hop2TargetExtraData = _hop2TargetExtraDataERC20(chainCHubToken, value2, targetTag);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop2TargetExtraData);

        (, uint feeA1, uint exA1) = bridgeVerifierChainA.calculateFee(CROSS_CHAIN_ID, IERC20(chainAHubToken), ctxValue);
        uint totalA1 = ctxValue + feeA1 + exA1;

        uint hop2IndexExpected = bridgeCross.getNextInitiateIndex(CHAIN_C_ID);
        uint depositedCBefore = bridgeCross.getTokenPair(CHAIN_C_ID, address(hubToken)).deposited;
        uint hubDepositedABefore = bridgeCross.getTokenPair(CHAIN_A_ID, address(hubToken)).deposited;
        uint chainAMintedBefore = bridgeChainA.getTokenPair(CROSS_CHAIN_ID, chainAHubToken).minted;
        uint chainAUserBalBefore = IERC20(chainAHubToken).balanceOf(USER);

        // L-1: snapshot the REAL ERC20 balances of all three bridges around the
        // MEASURED route (bootstrap-leg seeding is already behind us at this point, so
        // these snapshots capture only the measured route's own net movement).
        uint chainABridgeBalBefore = IERC20(chainAHubToken).balanceOf(address(bridgeChainA));
        uint hubBridgeBalBefore = hubToken.balanceOf(address(bridgeCross));

        vm.prank(USER);
        IERC20(chainAHubToken).approve(address(bridgeChainA), totalA1);

        // hop-1 (real): chain A's own `bridgeToken` entrypoint, capturing its emitted
        // `BridgeInitiated` rather than assuming its arguments.
        vm.recordLogs();
        vm.prank(USER);
        bridgeChainA.bridgeToken(CROSS_CHAIN_ID, IERC20(chainAHubToken), USER, ctxValue, feeA1, exA1, extraData);

        (
            bool foundHop1Init,
            uint evToChainID,
            uint evIndex,
            address evFromToken,
            address evToToken,
            ,
            address evTo,
            uint evValue,
            ,
            ,
            bytes memory evExtraData
        ) = _findBridgeInitiated(vm.getRecordedLogs(), address(bridgeChainA));
        assertTrue(foundHop1Init, "chain A must emit a real BridgeInitiated for hop-1");
        assertEq(evToChainID, CROSS_CHAIN_ID, "hop-1's own destination chain (from chain A's perspective) is the hub");
        assertEq(evTo, USER);
        assertEq(evValue, ctxValue, "event's value must be the fee-exclusive principal chain A submitted");
        assertEq(keccak256(evExtraData), keccak256(extraData), "event's extraData must be exactly what was submitted");

        // The plan's explicit field-mapping table, pinned down as an assertion: the
        // hub's finalize `toToken` must be the event's `toToken` (== `remoteToken`, the
        // HUB's own local token), and must NOT be the event's `fromToken` (chain A's
        // own wrapped address) — the two are deliberately distinct addresses here.
        assertTrue(evFromToken != evToToken, "sanity: chain A's wrapped token and the hub's local token must differ");
        assertEq(evFromToken, chainAHubToken, "event's fromToken must be chain A's own wrapped token");
        assertEq(evToToken, address(hubToken), "event's toToken (remoteToken) must be the HUB's local token");

        assertEq(
            IERC20(chainAHubToken).balanceOf(USER),
            chainAUserBalBefore - totalA1,
            "hop-1 origination must burn the wrapped token from USER on chain A"
        );
        assertEq(
            bridgeChainA.getTokenPair(CROSS_CHAIN_ID, chainAHubToken).minted,
            chainAMintedBefore - ctxValue,
            "hop-1 origination must decrement chain A's minted[hub] by ctxValue"
        );

        // L-1: chain A's OWN real ERC20 balance (of its wrapped `chainAHubToken`) must
        // net to zero across hop-1 origination -- USER's payment (principal + chain-A
        // fee) flows in, chain A's own fee flows out to `_dev`, and the wrapped
        // principal is burned, consistent with the mint/burn (not custody) model for a
        // non-origin pair.
        assertEq(
            IERC20(chainAHubToken).balanceOf(address(bridgeChainA)),
            chainABridgeBalBefore,
            "chain A bridge's own wrapped-token balance must be unchanged (transfer-in nets against fee-out + burn)"
        );

        // hop-1 finalize on the hub, using ONLY the captured event's arguments (per the
        // field-mapping table: fromChainID is chain A's OWN chain ID, known
        // independently -- it is NOT part of the event).
        vm.recordLogs();
        assertTrue(_signAndFinalize(CHAIN_A_ID, evIndex, evToToken, evTo, evValue, evExtraData, 5));
        Vm.Log[] memory hubLogs = vm.getRecordedLogs();

        vm.selectFork(crossForkID);
        assertEq(hubToken.balanceOf(USER), 0, "hop-1's forward must succeed on the hub (no D3 fallback)");
        assertEq(
            bridgeCross.getTokenPair(CHAIN_A_ID, address(hubToken)).deposited,
            hubDepositedABefore - ctxValue,
            "hub deposited[chainA] must decrease by ctxValue (net of the earlier bootstrap seeding)"
        );
        assertEq(
            bridgeCross.getTokenPair(CHAIN_C_ID, address(hubToken)).deposited,
            depositedCBefore + value2,
            "hub deposited[chainC] must increase by the forwarded principal"
        );

        // L-1: the hub's own REAL hubToken balance, across this single finalize call
        // (hop-1's own finalize immediately followed by hop-2's forward-initiate,
        // synchronously in the same transaction), must decrease by EXACTLY the fee
        // (networkFee2 + exFee2) that leaks out to `_dev` on the re-initiate to chain
        // C -- the full ctxValue is pulled out to the executor and pulled straight back
        // in by `bridgeTokenForwarded`'s own `_initiateBridge`, so only the fee is a
        // genuine net outflow. This is the same view as the §14.3 invariant: the
        // ledger delta (deposited[chainA] -ctxValue, deposited[chainC] +value2) is
        // consistent with the real balance delta only once the fee outflow is
        // accounted for (ctxValue - value2 = fee2 + ex2).
        assertEq(
            hubToken.balanceOf(address(bridgeCross)),
            hubBridgeBalBefore - (fee2 + ex2),
            "hub bridge's real ERC20 balance must decrease by exactly the leaked network+ex fee"
        );

        (bool foundHubFinalized,,,,,) = _findBridgeFinalized(hubLogs, address(bridgeCrossV2));
        assertTrue(foundHubFinalized, "hub must emit BridgeFinalized for hop-1's own finalize");
        (bool foundHop2Init,,,,,,,,,,) = _findBridgeInitiated(hubLogs, address(bridgeCrossV2));
        assertTrue(foundHop2Init, "hub must emit BridgeInitiated for hop-2 (the forward to chain C)");
        (bool foundForwardInitiated,,,,,,,,,) = _findForwardInitiated(hubLogs);
        assertTrue(foundForwardInitiated, "hub must emit ForwardInitiated for the forward to chain C");

        // hop-2 (real): finalize on chain C. With M-1's fix, this now goes all the way
        // through chain C's real `BridgeExecutor` into `mockTargetChainC`'s target
        // method, rather than stopping at a plain mint to USER.
        uint chainCBridgeBalBefore = IERC20(chainCHubToken).balanceOf(address(bridgeChainC));
        uint chainCMintedBefore = bridgeChainC.getTokenPair(CROSS_CHAIN_ID, chainCHubToken).minted;
        uint chainCTargetBalBefore = IERC20(chainCHubToken).balanceOf(address(mockTargetChainC));
        uint chainCUserBalBefore = IERC20(chainCHubToken).balanceOf(USER);

        vm.recordLogs();
        assertTrue(
            _signAndFinalizeChainC(
                CROSS_CHAIN_ID, hop2IndexExpected, chainCHubToken, USER, value2, hop2TargetExtraData, 5
            )
        );
        Vm.Log[] memory chainCLogs = vm.getRecordedLogs();

        (bool foundChainCFinalized,,,,,) = _findBridgeFinalized(chainCLogs, address(bridgeChainC));
        assertTrue(foundChainCFinalized, "chain C must emit BridgeFinalized for hop-2's own finalize");

        // M-1: the real chain-C target actually received/consumed the ERC20.
        (bool foundExtraCall,, bytes4 extraCallMethodID, bool extraCallSuccess, uint consumed, bytes memory returnData)
        = _findExtraCallExecuted(chainCLogs, address(bridgeChainC));
        assertTrue(foundExtraCall, "chain C must emit ExtraCallExecuted for hop-2's target call");
        assertTrue(extraCallSuccess, "chain C's target call must succeed");
        assertEq(
            extraCallMethodID,
            MockTargetContract.handleBridgeCallbackWithReturn.selector,
            "ExtraCallExecuted's methodID must match the target method encoded into hop-2's extraData"
        );
        assertEq(
            keccak256(returnData),
            keccak256(abi.encode(_expectedTargetReturnHash(chainCHubToken, value2, targetTag))),
            "ExtraCallExecuted's returnData must be the target's own deterministic proof of the exact call it processed"
        );

        uint remaining = IERC20(chainCHubToken).balanceOf(USER) - chainCUserBalBefore;
        assertEq(
            consumed + remaining,
            value2,
            "consumed by the target plus remaining sent to USER must equal hop-2's principal"
        );

        assertEq(
            IERC20(chainCHubToken).balanceOf(address(mockTargetChainC)),
            chainCTargetBalBefore + consumed,
            "chain C's real target contract must have received and consumed the forwarded ERC20"
        );

        // L-1: chain C's minted[hub] ledger and chain C bridge's own real ERC20
        // balance, cross-checked against each other (§14.3 invariant view): the bridge
        // mints the wrapped principal to itself then hands it straight to the executor
        // for the target call, so its OWN balance nets back to zero even though the
        // ledger records the full mint.
        assertEq(
            bridgeChainC.getTokenPair(CROSS_CHAIN_ID, chainCHubToken).minted,
            chainCMintedBefore + value2,
            "chain C minted[hub] must increase by hop-2's forwarded principal"
        );
        assertEq(
            IERC20(chainCHubToken).balanceOf(address(bridgeChainC)),
            chainCBridgeBalBefore,
            "chain C bridge's own real ERC20 balance must be unchanged (mint-to-self nets against hand-off to the executor)"
        );
    }

    /// @notice M-1: same ERC20 A -> B -> C route as `test_e2e_erc20_bothOrigin_ABC`,
    /// but chain C's target reverts -- chain C's OWN EXISTING (unmodified) fallback
    /// must mint/pay the ERC20 to USER directly there, exactly like the native
    /// `test_e2e_chainC_targetRevert_fallback` already proves for native value.
    /// @dev Bootstrap + hop-1 mirror `test_e2e_erc20_bothOrigin_ABC` exactly (that test
    /// already pins down the event field-mapping and every intermediate ledger delta
    /// for this shared setup); this test only adds the assertions specific to the
    /// target-revert fallback path.
    function test_e2e_erc20_chainC_targetRevert_fallback() public {
        vm.selectFork(crossForkID);
        mockTargetChainC.setShouldRevert(true);

        TestToken hubToken = new TestToken("Hub Origin", "HUB", 18);
        hubToken.mint(address(bridgeCross), 1000 ether);

        vm.startPrank(CrossOWNER);
        bridgeCross.registerToken(CHAIN_A_ID, true, address(hubToken), address(0x9501));
        bridgeCross.registerToken(CHAIN_C_ID, true, address(hubToken), address(0x9502));
        vm.stopPrank();

        vm.prank(chainAOwner);
        address chainAHubToken = bridgeChainA.createToken(CROSS_CHAIN_ID, address(hubToken), "AHUB", 18);
        vm.prank(chainCOwner);
        address chainCHubToken = bridgeChainC.createToken(CROSS_CHAIN_ID, address(hubToken), "CHUB", 18);

        // Bootstrap leg hub -> chain A (identical to the success test's STEP 0): seeds
        // USER's wrapped inventory on chain A so the measured route below has
        // something real to originate from.
        uint bootstrapAmount = 100 ether;
        (, uint bootFee, uint bootEx) = bridgeVerifierCross.calculateFee(CHAIN_A_ID, IERC20(hubToken), bootstrapAmount);
        uint bootTotal = bootstrapAmount + bootFee + bootEx;
        hubToken.mint(USER, bootTotal);
        vm.prank(USER);
        hubToken.approve(address(bridgeCrossV2), bootTotal);
        uint bootIndex = bridgeCross.getNextInitiateIndex(CHAIN_A_ID);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken(CHAIN_A_ID, IERC20(hubToken), USER, bootstrapAmount, bootFee, bootEx, "");
        assertTrue(
            _signAndFinalizeChainA(CROSS_CHAIN_ID, bootIndex, chainAHubToken, USER, bootstrapAmount, "", 5),
            "bootstrap leg finalize on chain A must succeed"
        );

        // Measured route: chain A -> hub -> chain C, hop-2 targeting the (now
        // reverting) mock target.
        uint value2 = 10 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(hubToken), value2);
        uint ctxValue = value2 + fee2 + ex2;
        require(ctxValue <= bootstrapAmount, "bootstrap inventory must cover the measured route's hop-1 principal");
        bytes memory targetTag = bytes("hop2-erc20-revert");
        bytes memory hop2TargetExtraData = _hop2TargetExtraDataERC20(chainCHubToken, value2, targetTag);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop2TargetExtraData);

        (, uint feeA1, uint exA1) = bridgeVerifierChainA.calculateFee(CROSS_CHAIN_ID, IERC20(chainAHubToken), ctxValue);
        uint totalA1 = ctxValue + feeA1 + exA1;
        uint hop1IndexExpected = bridgeChainA.getNextInitiateIndex(CROSS_CHAIN_ID);
        uint hop2IndexExpected = bridgeCross.getNextInitiateIndex(CHAIN_C_ID);

        vm.prank(USER);
        IERC20(chainAHubToken).approve(address(bridgeChainA), totalA1);
        vm.prank(USER);
        bridgeChainA.bridgeToken(CROSS_CHAIN_ID, IERC20(chainAHubToken), USER, ctxValue, feeA1, exA1, extraData);

        // hop-1's own finalize on the hub -- `evToToken`/`evTo`/`evValue`/`evExtraData`
        // are `hubToken`/`USER`/`ctxValue`/`extraData` by construction (chain A's
        // `bridgeToken` call above), exactly as the sibling success test independently
        // derives (and asserts) off the real `BridgeInitiated` event.
        assertTrue(_signAndFinalize(CHAIN_A_ID, hop1IndexExpected, address(hubToken), USER, ctxValue, extraData, 5));

        // hop-2: finalize on chain C, whose target now reverts.
        uint targetBalBefore = IERC20(chainCHubToken).balanceOf(address(mockTargetChainC));
        assertTrue(
            _signAndFinalizeChainC(
                CROSS_CHAIN_ID, hop2IndexExpected, chainCHubToken, USER, value2, hop2TargetExtraData, 5
            )
        );

        assertEq(
            IERC20(chainCHubToken).balanceOf(address(mockTargetChainC)),
            targetBalBefore,
            "reverting target must not receive/keep any ERC20"
        );
        vm.selectFork(crossForkID);
        assertEq(
            IERC20(chainCHubToken).balanceOf(USER),
            value2,
            "chain C's own existing fallback must mint/pay the ERC20 to USER directly when the target reverts"
        );
    }

    /// @notice §14.5 operational case: one `finalizeBridgeBatch` containing BOTH a
    /// forward item (to chain C) and an ordinary (non-forward) item must process both
    /// correctly in the same transaction.
    function test_batch_mixedForwardAndOrdinary() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValueForward = value2 + fee2 + ex2;
        bytes memory forwardExtraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, "");

        uint ordinaryAmount = 2 ether;

        vm.selectFork(bscForkID);
        (, uint netF, uint exF) =
            bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValueForward);
        uint indexForward = nextIndexBSC;
        vm.deal(USER, ctxValueForward + netF + exF);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxValueForward + netF + exF}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValueForward, netF, exF, forwardExtraData
        );
        bscIncrementIndex();

        (, uint netO, uint exO) =
            bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ordinaryAmount);
        uint indexOrdinary = nextIndexBSC;
        vm.deal(USER, ordinaryAmount + netO + exO);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ordinaryAmount + netO + exO}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ordinaryAmount, netO, exO, ""
        );
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        bytes32 hForward = keccak256(
            abi.encode(
                FINALIZE_TYPEHASH,
                BSC_CHAIN_ID,
                indexForward,
                address(NATIVE_TOKEN),
                USER,
                ctxValueForward,
                keccak256(forwardExtraData)
            )
        );
        bytes32 hOrdinary = keccak256(
            abi.encode(
                FINALIZE_TYPEHASH,
                BSC_CHAIN_ID,
                indexOrdinary,
                address(NATIVE_TOKEN),
                USER,
                ordinaryAmount,
                keccak256(bytes(""))
            )
        );
        bytes32 hashForward = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hForward);
        bytes32 hashOrdinary = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), hOrdinary);

        uint8[] memory vF = new uint8[](5);
        bytes32[] memory rF = new bytes32[](5);
        bytes32[] memory sF = new bytes32[](5);
        uint8[] memory vO = new uint8[](5);
        bytes32[] memory rO = new bytes32[](5);
        bytes32[] memory sO = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            (vF[i], rF[i], sF[i]) = vm.sign(VALIDATOR_PKs[i], hashForward);
            (vO[i], rO[i], sO[i]) = vm.sign(VALIDATOR_PKs[i], hashOrdinary);
        }

        IBridgeRegistry.FinalizeArguments[] memory args = new IBridgeRegistry.FinalizeArguments[](2);
        args[0] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexForward,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ctxValueForward,
            extraData: forwardExtraData
        });
        args[1] = IBridgeRegistry.FinalizeArguments({
            fromChainID: BSC_CHAIN_ID,
            index: indexOrdinary,
            toToken: IERC20(NATIVE_TOKEN),
            to: USER,
            value: ordinaryAmount,
            extraData: ""
        });
        uint8[][] memory vArray = new uint8[][](2);
        bytes32[][] memory rArray = new bytes32[][](2);
        bytes32[][] memory sArray = new bytes32[][](2);
        vArray[0] = vF;
        rArray[0] = rF;
        sArray[0] = sF;
        vArray[1] = vO;
        rArray[1] = rO;
        sArray[1] = sO;

        uint depositedCBefore = bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited;
        assertTrue(bridgeCross.finalizeBridgeBatch(args, vArray, rArray, sArray));

        assertEq(USER.balance, ordinaryAmount, "the ordinary item must pay USER directly; the forward item must not");
        assertEq(
            bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited,
            depositedCBefore + value2,
            "the forward item must have initiated hop-2 in the SAME batch as the ordinary item"
        );
    }

    /// @notice §14.5: measures the actual nested extraData length (chain A's top-level
    /// extraData, embedding chain B's forward call which itself embeds chain C's own
    /// extraData) against `_maxExtraDataLength`'s boundary — success at exactly the
    /// configured max, revert one byte over.
    function test_nestedExtraData_maxLengthBoundary() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), value2);
        bytes memory hop3ExtraData = _hop3TargetExtraData(value2);
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, value2, fee2, ex2, hop3ExtraData);
        uint ctxValue = value2 + fee2 + ex2;

        uint measuredLength = extraData.length;
        emit log_named_uint("nested extraData length (hop-1 extraData embedding hop-2+hop-3)", measuredLength);

        // At exactly the measured length: must succeed. The length check that matters
        // for hop-1's OWN extraData argument lives on the ORIGINATING bridge (BSC),
        // since that's where it's first submitted as a `bridgeToken` argument.
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setMaxExtraDataLength(measuredLength);
        _hop1NativeAndFinalize(ctxValue, extraData);
        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "extraData at exactly the configured max length must still succeed");

        // One byte over: bridgeToken itself reverts at hop-1's ORIGINATION.
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setMaxExtraDataLength(measuredLength - 1);
        (, uint netA1, uint exA1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValue);
        vm.deal(USER, ctxValue + netA1 + exA1);
        vm.prank(USER);
        vm.expectRevert(); // BaseBridgeExtraDataTooLong
        bridgeBSC.bridgeToken{value: ctxValue + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValue, netA1, exA1, extraData
        );
    }

    /// @notice §14.5 gas measurement: total gas of a finalize batch item that performs a
    /// forward (including the `ForwardLib` DELEGATECALL overhead), against an equivalent
    /// ordinary (non-forward) item's gas, to isolate the forward-specific incremental
    /// cost. Reported plainly for the implementation report's `_postCallGasReserve`
    /// recommendation — `_postCallGasReserve` itself budgets BaseBridge's POST-executor-
    /// call bookkeeping (approve-clear, `_withdrawToken`, event emission), which does
    /// NOT scale with forward complexity, so this measurement informs the report's
    /// recommendation rather than asserting a specific reserve value here.
    function test_gasUsed_forwardVsOrdinary() public {
        vm.selectFork(crossForkID);
        uint amount = 1 ether;

        // Ordinary (non-forward) baseline.
        vm.selectFork(bscForkID);
        (, uint netO, uint exO) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), amount);
        uint indexOrdinary = nextIndexBSC;
        vm.deal(USER, amount + netO + exO);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: amount + netO + exO}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, amount, netO, exO, ""
        );
        bscIncrementIndex();
        vm.selectFork(crossForkID);
        assertTrue(_signAndFinalize(BSC_CHAIN_ID, indexOrdinary, address(NATIVE_TOKEN), USER, amount, "", 5));
        uint ordinaryGas = vm.lastCallGas().gasTotalUsed;

        // Forward (to chain C) with the same principal.
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), amount);
        uint ctxValue = amount + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, amount, fee2, ex2, "");

        vm.selectFork(bscForkID);
        (, uint netA1, uint exA1) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValue);
        uint indexForward = nextIndexBSC;
        vm.deal(USER, ctxValue + netA1 + exA1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxValue + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValue, netA1, exA1, extraData
        );
        bscIncrementIndex();
        vm.selectFork(crossForkID);
        assertTrue(_signAndFinalize(BSC_CHAIN_ID, indexForward, address(NATIVE_TOKEN), USER, ctxValue, extraData, 5));
        uint forwardGas = vm.lastCallGas().gasTotalUsed;

        emit log_named_uint("finalizeBridgeBatch gas - ordinary item", ordinaryGas);
        emit log_named_uint("finalizeBridgeBatch gas - forward item (incl. ForwardLib DELEGATECALL)", forwardGas);
        assertTrue(forwardGas > ordinaryGas, "a forward item must cost more gas than an ordinary item");
        uint incrementalForwardGas = forwardGas - ordinaryGas;
        emit log_named_uint("forward-specific incremental gas (delta)", incrementalForwardGas);

        // Sanity bound: comfortably within a single block's gas limit with generous
        // headroom for batching multiple items.
        assertLt(
            forwardGas, 1_000_000, "a single forward finalize item should stay well under typical block gas budgets"
        );
    }

    // ----------------------------------------------------------------
    // M-3: crossSupply() pass-through accounting (plan spec §6.1/§14.3)
    // ----------------------------------------------------------------

    /// @notice BSC -> CROSS -> chainC pass-through: `crossSupply()` increases by
    /// `ctxValue` (the FULL BSC-leg amount finalized to CROSS), not merely
    /// `hop2Principal` (the smaller amount forwarded onward) — this is the DEFINITION of
    /// `crossSupply()` (net CROSS inflow from BSC, spec §6.1), not drift, and is pinned
    /// here so a future reader does not "fix" it as a bug.
    function test_crossSupply_passThrough_increases() public {
        vm.selectFork(crossForkID);
        uint crossSupplyBefore = bridgeCross.crossSupply();
        uint depositedCBefore = bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited;

        uint hop2Principal = 1 ether;
        (, uint networkFee2, uint exFee2) =
            bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), hop2Principal);
        uint ctxValue = hop2Principal + networkFee2 + exFee2;
        bytes memory extraData = _forwardExtraData(CHAIN_C_ID, USER, hop2Principal, networkFee2, exFee2, "");

        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "forward must succeed for this accounting comparison to be meaningful");
        assertEq(
            bridgeCross.crossSupply(),
            crossSupplyBefore + ctxValue,
            "crossSupply() must increase by ctxValue (the FULL BSC-leg amount) -- intentional definition, spec section 6.1, not drift"
        );
        assertEq(
            bridgeCross.getTokenPair(CHAIN_C_ID, Const.NATIVE_TOKEN).deposited,
            depositedCBefore + hop2Principal,
            "chain C's ledger changes by hop2Principal ONLY (fees stay on chain B)"
        );
    }

    /// @notice CROSS -> BSC withdrawal decreases `crossSupply()` by exactly the
    /// principal withdrawn (fees are a separate `_dev` payout, not part of the
    /// `deposited`/`minted` accounting `_depositToken` touches).
    function test_crossSupply_returnToBSC_decreases() public {
        vm.selectFork(crossForkID);
        // Seed minted[BSC][NATIVE] first via an ordinary (non-forward) inbound finalize.
        _hop1NativeAndFinalize(10 ether, "");
        vm.selectFork(crossForkID);
        uint crossSupplyBefore = bridgeCross.crossSupply();
        assertTrue(crossSupplyBefore >= 10 ether);

        uint withdrawPrincipal = 2 ether;
        (, uint fee, uint ex) =
            bridgeVerifierCross.calculateFee(BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), withdrawPrincipal);
        uint total = withdrawPrincipal + fee + ex;
        vm.deal(USER, total);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken{value: total}(
            BSC_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, withdrawPrincipal, fee, ex, ""
        );

        assertEq(
            bridgeCross.crossSupply(),
            crossSupplyBefore - withdrawPrincipal,
            "crossSupply() must decrease by exactly the withdrawn PRINCIPAL (fee is a separate _dev payout)"
        );
    }

    /// @notice A simple CROSS<->chainC route (no BSC leg at all) must leave
    /// `crossSupply()` (which reads ONLY the BSC_CHAIN_ID pair) completely unchanged.
    function test_crossSupply_unchanged_onNonBscRoute() public {
        vm.selectFork(crossForkID);
        uint crossSupplyBefore = bridgeCross.crossSupply();

        uint amount = 1 ether;
        (, uint fee, uint ex) = bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), amount);
        uint total = amount + fee + ex;
        vm.deal(USER, total);
        vm.prank(USER);
        bridgeCrossV2.bridgeToken{value: total}(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), USER, amount, fee, ex, "");

        assertEq(
            bridgeCross.crossSupply(),
            crossSupplyBefore,
            "a CROSS<->chainC route must not touch the BSC-pair-only crossSupply()"
        );
    }

    /// @notice `crossSupplyLimit`'s circuit breaker (checked at hop-1's OWN finalize,
    /// independent of and prior to the forward attempt) includes pass-through volume:
    /// a second BSC->CROSS->chainC pass-through that would push `crossSupply()` over the
    /// configured limit must go pending with `CrossSupplyLimitExceeded`.
    function test_crossSupplyLimit_includesPassThrough() public {
        vm.selectFork(crossForkID);
        uint tightLimit = bridgeCross.crossSupply() + 5 ether;
        vm.prank(CrossOWNER);
        bridgeCrossV2.setCrossSupplyLimit(tightLimit);

        uint hop2PrincipalOk = 1 ether;
        (, uint fee2a, uint ex2a) =
            bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), hop2PrincipalOk);
        uint ctxValueOk = hop2PrincipalOk + fee2a + ex2a;
        bytes memory extraDataOk = _forwardExtraData(CHAIN_C_ID, USER, hop2PrincipalOk, fee2a, ex2a, "");
        _hop1NativeAndFinalize(ctxValueOk, extraDataOk);
        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "first pass-through must succeed under the (still-headroom) limit");

        uint hop2PrincipalOver = 10 ether;
        (, uint fee2b, uint ex2b) =
            bridgeVerifierCross.calculateFee(CHAIN_C_ID, IERC20(Const.NATIVE_TOKEN), hop2PrincipalOver);
        uint ctxValueOver = hop2PrincipalOver + fee2b + ex2b;
        bytes memory extraDataOver = _forwardExtraData(CHAIN_C_ID, USER, hop2PrincipalOver, fee2b, ex2b, "");

        vm.selectFork(bscForkID);
        (, uint netA1, uint exA1) =
            bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), ctxValueOver);
        uint index = nextIndexBSC;
        vm.deal(USER, ctxValueOver + netA1 + exA1);
        vm.prank(USER);
        bridgeBSC.bridgeToken{value: ctxValueOver + netA1 + exA1}(
            CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, ctxValueOver, netA1, exA1, extraDataOver
        );
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        _signAndFinalize(BSC_CHAIN_ID, index, address(NATIVE_TOKEN), USER, ctxValueOver, extraDataOver, 5);
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.CrossSupplyLimitExceeded,
            "pass-through volume must be included in the crossSupplyLimit circuit breaker"
        );
    }
}
