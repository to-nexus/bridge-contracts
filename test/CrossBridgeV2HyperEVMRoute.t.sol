// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {BSCBridge} from "../src/BSCBridge.sol";
import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {PriceFeed} from "../src/PriceFeed.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {Const} from "../src/lib/Const.sol";

import {CrossBridgeV2MultihopTest} from "./CrossBridgeV2Multihop.t.sol";
import {MockTargetContract} from "./BridgeExecutor.t.sol";
import {CrossMintableERC20Code} from "../src/token/CrossMintableERC20Code.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/**
 * @title CrossBridgeV2HyperEVMRouteTest
 * @notice The PRODUCTION route, end to end and in one user action:
 *
 *     BSC --(bridgeToken)--> CROSS --(forward)--> HyperEVM --> target contract
 *
 * @dev Why this exists on top of `CrossBridgeV2MultihopTest`'s generic A->B->C tests:
 * those pin the MECHANISM using synthetic chain ids (90501 / 90999). This one pins the
 * REAL ROUTE with the REAL chain ids the deployment actually uses, so that a change to
 * chain-id handling, pair registration, or the D1 invariant that happens to be fine for
 * a small synthetic id but wrong for HyperEVM's is caught here.
 *
 * The user performs exactly ONE action: `bridgeBSC.bridgeToken(...)` on BSC. Everything
 * after that — the CROSS finalize, the forward, the HyperEVM finalize, and the HyperEVM
 * target call — is driven by validator signatures and the bridge's own logic, which is
 * precisely what "한번에 bridge" means operationally.
 *
 * The HyperEVM leaf is a REAL, independently deployed bridge instance (its own
 * `BridgeVerifier` / `PriceFeed` / `BridgeExecutor` / target), deployed with the same
 * pattern the fixture already uses for chain A and chain C. It is a separate CONTRACT
 * instance standing in for a separate real chain, exactly as `bridgeBSC` and
 * `bridgeCross` already coexist.
 */
contract CrossBridgeV2HyperEVMRouteTest is CrossBridgeV2MultihopTest {
    /// @dev HyperEVM testnet. Mainnet is 999; the routing logic is identical and is
    /// covered by `test_e2e_bsc_cross_hyperevm_mainnetChainId`.
    uint internal constant HYPEREVM_TESTNET_CHAIN_ID = 998;
    uint internal constant HYPEREVM_MAINNET_CHAIN_ID = 999;

    address internal hyperOwner;
    BSCBridge internal bridgeHyper;
    BridgeVerifier internal bridgeVerifierHyper;
    PriceFeed internal priceFeedHyper;
    BridgeExecutor internal bridgeExecutorHyper;
    MockTargetContract internal mockTargetHyper;
    ICrossMintableERC20Code internal crossMintableERC20CodeHyper;
    TestToken internal hyperLocalCrossToken;

    function setUp() public virtual override {
        super.setUp();

        vm.selectFork(crossForkID);
        hyperOwner = makeAddr("hyperevm-owner");

        hyperLocalCrossToken = new TestToken("HyperEVM Cross", "HCROSS", 18);
        BSCBridge implHyper = new BSCBridge();
        ERC1967Proxy proxyHyper = new ERC1967Proxy(address(implHyper), bytes(""));
        bridgeHyper = BSCBridge(payable(address(proxyHyper)));
        bridgeHyper.initializeBSCBridge(
            hyperOwner, payable(hyperOwner), threshold, CROSS_CHAIN_ID, address(hyperLocalCrossToken), 0
        );

        vm.startPrank(hyperOwner);
        bridgeHyper.grantRole(Const.ADMIN_ROLE, hyperOwner);
        bridgeHyper.grantRole(Const.EDITOR_ROLE, hyperOwner);
        bridgeHyper.grantRole(Const.OPERATOR_ROLE, hyperOwner);
        bytes32[] memory roles = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            roles[i] = Const.VALIDATOR_ROLE;
        }
        bridgeHyper.grantRoleBatch(roles, VALIDATORS);

        PriceFeed implPF = new PriceFeed();
        ERC1967Proxy proxyPF = new ERC1967Proxy(address(implPF), bytes(""));
        priceFeedHyper = PriceFeed(address(proxyPF));
        priceFeedHyper.initialize(hyperOwner, DOLLAR_DECIMALS);
        priceFeedHyper.grantRole(Const.PRICER_ROLE, hyperOwner);

        bridgeVerifierHyper = new BridgeVerifier(
            hyperOwner, address(bridgeHyper), address(priceFeedHyper), 200_000, 0, 0, 0, 0, 0, 2 hours
        );
        bridgeHyper.setBridgeVerifier(bridgeVerifierHyper);

        crossMintableERC20CodeHyper = ICrossMintableERC20Code(address(new CrossMintableERC20Code(address(bridgeHyper))));
        bridgeHyper.setCrossMintableERC20Code(crossMintableERC20CodeHyper);

        // HyperEVM's native coin <-> CROSS's native coin, registered symmetrically to
        // CROSS_CHAIN_ID (the leaf never stores its OWN chain id — routing is decided by
        // the HUB-side pair registration below).
        bridgeHyper.registerToken(CROSS_CHAIN_ID, false, address(NATIVE_TOKEN), address(NATIVE_TOKEN));
        vm.deal(address(bridgeHyper), 1_000 ether);

        bridgeExecutorHyper = new BridgeExecutor(hyperOwner, address(bridgeHyper));
        bridgeHyper.setBridgeExecutor(IBridgeExecutor(address(bridgeExecutorHyper)));
        mockTargetHyper = new MockTargetContract();
        bridgeExecutorHyper.addWhitelistTarget(address(mockTargetHyper));
        vm.stopPrank();

        // The hub's registration is what makes chain id 998 route to this leaf.
        // isOrigin=true so `_checkInitiateAmount`'s `minted >= value` precondition (which
        // only applies to NOT-origin pairs) never gates outbound forwards.
        vm.startPrank(CrossOWNER);
        bridgeCross.registerToken(HYPEREVM_TESTNET_CHAIN_ID, true, address(NATIVE_TOKEN), address(NATIVE_TOKEN));
        bridgeCross.registerToken(HYPEREVM_MAINNET_CHAIN_ID, true, address(NATIVE_TOKEN), address(NATIVE_TOKEN));
        vm.stopPrank();
    }

    /// @dev hop-3 payload: the HyperEVM-side target call that actually consumes value.
    function _hyperTargetExtraData(uint value3) internal view returns (bytes memory) {
        bytes memory call_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, value3, bytes("hyperevm")
        );
        return abi.encodePacked(address(mockTargetHyper), call_);
    }

    /// @dev Validator-signed finalize on the HyperEVM leaf (mirrors
    /// `_signAndFinalizeChainC`, retargeted at `bridgeHyper`).
    function _signAndFinalizeHyper(
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
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeHyper.domainSeparator(), h);

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
        ok = bridgeHyper.finalizeBridgeBatch(args, vArray, rArray, sArray);
    }

    /// @dev Drives the whole BSC -> CROSS -> HyperEVM route for `toChainID` and asserts
    /// every leg. Returns nothing; every meaningful check is an assert.
    function _runRoute(uint toChainID) internal {
        vm.selectFork(crossForkID);

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(toChainID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;

        bytes memory hop3 = _hyperTargetExtraData(value2);
        bytes memory extraData = _forwardExtraData(toChainID, USER, value2, fee2, ex2, hop3);

        uint hop2Index = bridgeCross.getNextInitiateIndex(toChainID);
        uint hubDepositedBefore = bridgeCross.getTokenPair(toChainID, Const.NATIVE_TOKEN).deposited;
        uint hubMintedBscBefore = bridgeCross.getTokenPair(BSC_CHAIN_ID, Const.NATIVE_TOKEN).minted;
        uint hyperMintedBefore = bridgeHyper.getTokenPair(CROSS_CHAIN_ID, Const.NATIVE_TOKEN).minted;
        uint targetBalBefore = address(mockTargetHyper).balance;

        // ---- The ONE user action: bridgeToken on BSC. ----
        // (helper also performs the validator-signed CROSS finalize that this emits)
        _hop1NativeAndFinalize(ctxValue, extraData);

        // ---- leg 1+2 assertions: CROSS forwarded, no D3 fallback ----
        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "forward must succeed on CROSS: no fallback payout to USER on the hub");
        assertEq(
            bridgeCross.getTokenPair(toChainID, Const.NATIVE_TOKEN).deposited,
            hubDepositedBefore + value2,
            "hub ledger for the HyperEVM pair must record exactly the forwarded principal"
        );
        assertEq(
            bridgeCross.getTokenPair(BSC_CHAIN_ID, Const.NATIVE_TOKEN).minted,
            hubMintedBscBefore + ctxValue,
            "hub's BSC-leg minted must rise by the full finalized amount (spec 6.1: pass-through counts)"
        );

        // ---- leg 3: finalize on HyperEVM, which calls the real target ----
        assertTrue(
            _signAndFinalizeHyper(CROSS_CHAIN_ID, hop2Index, address(NATIVE_TOKEN), USER, value2, hop3, 5),
            "HyperEVM finalize must succeed"
        );

        assertEq(
            address(mockTargetHyper).balance,
            targetBalBefore + value2,
            "HyperEVM's target contract must have received and consumed the forwarded value"
        );
        assertEq(USER.balance, 0, "HyperEVM target call succeeded: no fallback payout to USER there");
        assertEq(
            bridgeHyper.getTokenPair(CROSS_CHAIN_ID, Const.NATIVE_TOKEN).minted,
            hyperMintedBefore + value2,
            "HyperEVM ledger must record the delivered principal"
        );
    }

    /// @notice BSC -> CROSS -> HyperEVM testnet (998), one user action, target called.
    function test_e2e_bsc_cross_hyperevm_native() public {
        _runRoute(HYPEREVM_TESTNET_CHAIN_ID);
    }

    /// @notice Same route against HyperEVM mainnet's chain id (999).
    function test_e2e_bsc_cross_hyperevm_mainnetChainId() public {
        _runRoute(HYPEREVM_MAINNET_CHAIN_ID);
    }

    /// @notice If HyperEVM's target reverts, HyperEVM's OWN existing fallback pays USER
    /// there — the CROSS forward itself still succeeded and must NOT be rolled back.
    function test_e2e_bsc_cross_hyperevm_targetRevert_fallback() public {
        vm.selectFork(crossForkID);
        uint value2 = 1 ether;
        (, uint fee2, uint ex2) =
            bridgeVerifierCross.calculateFee(HYPEREVM_TESTNET_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;

        bytes memory hop3 = _hyperTargetExtraData(value2);
        bytes memory extraData =
            _forwardExtraData(HYPEREVM_TESTNET_CHAIN_ID, USER, value2, fee2, ex2, hop3);

        uint hop2Index = bridgeCross.getNextInitiateIndex(HYPEREVM_TESTNET_CHAIN_ID);

        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(USER.balance, 0, "the CROSS forward itself must still have succeeded");

        // Make HyperEVM's target revert, then finalize there.
        vm.prank(hyperOwner);
        mockTargetHyper.setShouldRevert(true);

        uint targetBalBefore = address(mockTargetHyper).balance;
        assertTrue(
            _signAndFinalizeHyper(CROSS_CHAIN_ID, hop2Index, address(NATIVE_TOKEN), USER, value2, hop3, 5),
            "finalize still succeeds; only the extra call fails"
        );

        assertEq(address(mockTargetHyper).balance, targetBalBefore, "reverting target must receive nothing");
        assertEq(USER.balance, value2, "HyperEVM's own normal-flow fallback must pay USER directly");
    }
}
