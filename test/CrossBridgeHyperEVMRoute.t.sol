// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {BSCBridge} from "../src/BSCBridge.sol";
import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {PriceFeed} from "../src/PriceFeed.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";

import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {HyperMintableERC20} from "../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {MockTargetContract} from "./BridgeExecutor.t.sol";
import {CrossBridgeMultihopTest} from "./CrossBridgeMultihop.t.sol";

import {TestToken} from "./token/TestToken.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/**
 * @title CrossBridgeHyperEVMRouteTest
 * @notice The PRODUCTION route, end to end and in one user action:
 *
 *     BSC --(bridgeToken)--> CROSS --(forward)--> HyperEVM --> target contract
 *
 * @dev Why this exists on top of `CrossBridgeMultihopTest`'s generic A->B->C tests:
 * those pin the MECHANISM using synthetic chain ids (90501 / 90999). This one pins the
 * REAL ROUTE with the REAL chain ids the deployment actually uses, so that a change to
 * chain-id handling, pair registration, or the native pass-through invariant that happens to be fine for
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
contract CrossBridgeHyperEVMRouteTest is CrossBridgeMultihopTest {
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
    HyperMintableERC20Code internal crossMintableERC20CodeHyper;
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

        // `HyperMintableERC20Code` in place of the V1 `CrossMintableERC20Code` — same
        // `ICrossMintableERC20Code` surface, but its tokens carry the HyperCore link slot
        // (see the wrapped-token tests below). Since `HyperMintableERC20Code` now inherits
        // `CrossMintableERC20V2Code`, the factory creates and owns its own token beacon
        // INSIDE `initialize` — there is no separate beacon deploy step, and no `tokenAdmin` parameter
        // any more (every created token's `defaultAdmin()` is the factory itself). `hyperOwner`
        // is the factory's `ADMIN_ROLE`, consistent with every other role in this fixture being
        // centralized on `hyperOwner`. `initialize` is inherited (not redeclared) from
        // `CrossMintableERC20V2Code`, so it is referenced via that declaring contract here —
        // `abi.encodeCall`'s magic member lookup only resolves functions declared directly on
        // the named type, not merely-inherited ones; the selector only depends on the function
        // signature, so this still dispatches correctly against `codeImplHyper`'s actual
        // (inherited) `initialize`. The factory itself is a UUPS `ERC1967Proxy`, atomically
        // initialized here exactly as a real deploy script would.
        HyperMintableERC20 tokenImplHyper = new HyperMintableERC20();
        HyperMintableERC20Code codeImplHyper = new HyperMintableERC20Code();
        ERC1967Proxy codeProxyHyper = new ERC1967Proxy(
            address(codeImplHyper),
            abi.encodeCall(CrossMintableERC20V2Code.initialize, (hyperOwner, address(bridgeHyper), address(tokenImplHyper)))
        );
        crossMintableERC20CodeHyper = HyperMintableERC20Code(address(codeProxyHyper));
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

        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, fromChainID, index, token, to, value, keccak256(extraData)));
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

        // ---- leg 1+2 assertions: CROSS forwarded, no fallback payout ----
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
        bytes memory extraData = _forwardExtraData(HYPEREVM_TESTNET_CHAIN_ID, USER, value2, fee2, ex2, hop3);

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

    // =====================================================================
    // Wrapped ERC20 path: proves
    // `BridgeRegistry.createToken -> HyperMintableERC20Code -> HyperMintableERC20`
    // is fully compatible end to end through a real bridge instance, on top of the native
    // route already exercised above.
    // =====================================================================

    /// @dev Creates a fresh wrapped `HyperMintableERC20` for an arbitrary CROSS-side remote
    /// token, through the real `bridgeHyper.createToken` entry point (EDITOR_ROLE = hyperOwner).
    function _createWrappedToken() internal returns (address tokenAddress, address remoteToken) {
        vm.selectFork(crossForkID);
        remoteToken = makeAddr("hyperevm-wrapped-remote");

        vm.prank(hyperOwner);
        tokenAddress = bridgeHyper.createToken(CROSS_CHAIN_ID, remoteToken, "TT", 18);
    }

    /// `createToken` deploys a real `HyperMintableERC20` tracked by the factory,
    /// and its automatic `registerToken(isOrigin=false)` call is visible through `getTokenPair`.
    function test_createToken_wrapped_registeredAndTrackedByFactory() public {
        (address tokenAddress, address remoteToken) = _createWrappedToken();

        // Real HyperMintableERC20, tracked by the factory that created it.
        assertTrue(crossMintableERC20CodeHyper.isHyperMintableERC20(tokenAddress));
        HyperMintableERC20 wrapped = HyperMintableERC20(tokenAddress);
        assertEq(wrapped.name(), "Cross Bridge TT");
        assertEq(wrapped.symbol(), "TTx");
        assertEq(wrapped.decimals(), 18);

        // createToken's internal registerToken(isOrigin=false) call registered the pair.
        IBridgeRegistry.TokenPair memory pair = bridgeHyper.getTokenPair(CROSS_CHAIN_ID, tokenAddress);
        assertEq(pair.localToken, tokenAddress);
        assertEq(pair.remoteToken, remoteToken);
        assertFalse(pair.isOrigin);
    }

    /// The bridge holds MINTER_ROLE on the token it created: a validator-signed finalize
    /// mints to the recipient, and bridging back out burns via the same wrapped token —
    /// proving the whole factory-created token is usable end to end through the bridge.
    function test_wrappedToken_finalizeMintsAndReverseBridgeBurns() public {
        (address tokenAddress,) = _createWrappedToken();
        HyperMintableERC20 wrapped = HyperMintableERC20(tokenAddress);
        assertTrue(wrapped.hasRole(Const.MINTER_ROLE, address(bridgeHyper)));

        address recipient = makeAddr("hyperevm-wrapped-recipient");
        uint value = 10 ether;
        uint index = bridgeHyper.getNextFinalizeIndex(CROSS_CHAIN_ID);

        assertTrue(
            _signAndFinalizeHyper(CROSS_CHAIN_ID, index, tokenAddress, recipient, value, bytes(""), 5),
            "inbound finalize must mint the wrapped token to recipient"
        );
        assertEq(wrapped.balanceOf(recipient), value);
        assertEq(bridgeHyper.getTokenPair(CROSS_CHAIN_ID, tokenAddress).minted, value);

        // Reverse direction: bridging the wrapped token back out burns it (capped by `minted`).
        vm.selectFork(crossForkID);
        uint value2 = 4 ether;
        (, uint networkFee, uint exFee) = bridgeVerifierHyper.calculateFee(CROSS_CHAIN_ID, IERC20(tokenAddress), value2);
        uint supplyBefore = wrapped.totalSupply();

        vm.prank(recipient);
        wrapped.approve(address(bridgeHyper), value2 + networkFee + exFee);

        vm.prank(recipient);
        assertTrue(
            bridgeHyper.bridgeToken(CROSS_CHAIN_ID, IERC20(tokenAddress), USER, value2, networkFee, exFee, bytes(""))
        );

        assertEq(wrapped.balanceOf(recipient), value - value2 - networkFee - exFee, "burn + fee left the sender");
        assertEq(wrapped.totalSupply(), supplyBefore - value2, "only the bridged value is burned, not the fee");
        assertEq(
            bridgeHyper.getTokenPair(CROSS_CHAIN_ID, tokenAddress).minted,
            value - value2,
            "minted accounting must fall by the burned amount"
        );
    }

    /// The factory-delegated `setHyperCoreDeployer` path works on a token created through
    /// the real bridge, and the RAW storage slot (not just the getter) reflects the finalizer.
    function test_factoryPath_setHyperCoreDeployer_writesRawSlot() public {
        (address tokenAddress,) = _createWrappedToken();
        address finalizer = makeAddr("hyperevm-finalizer");

        vm.prank(hyperOwner);
        crossMintableERC20CodeHyper.setHyperCoreDeployer(tokenAddress, finalizer);

        bytes32 raw = vm.load(tokenAddress, HyperMintableERC20(tokenAddress).HYPERCORE_DEPLOYER_SLOT());
        assertEq(raw, bytes32(uint(uint160(finalizer))));
        assertEq(HyperMintableERC20(tokenAddress).hyperCoreDeployer(), finalizer);
    }
}
