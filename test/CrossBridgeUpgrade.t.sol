// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {CrossBridge} from "../src/CrossBridge.sol";
import {CrossBridge} from "../src/CrossBridge.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";

import {BridgeExecutorTest} from "./BridgeExecutor.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/**
 * @title CrossBridgeUpgradeTest
 * @notice Upgrade storage-layout integrity and
 * initializer-sealing regression suite. Fixture follows the same upgrade order as
 * `CrossBridgeForwardTest` (`CrossBridge` deployed/initialized exactly as in the base
 * fixture, upgraded to `CrossBridge` INSIDE each test rather than in `setUp` — so
 * every test can capture pre-upgrade state to compare against).
 */
contract CrossBridgeUpgradeTest is BridgeExecutorTest {
    /// @dev `_bscChainID`'s storage slot, identical in `CrossBridge` and
    /// `CrossBridge` — confirmed via `forge inspect CrossBridge storageLayout` and
    /// `forge inspect CrossBridge storageLayout` (both report slot 101). `_bscChainID`
    /// is `private` with no getter, so this is the only direct way to compare it
    /// byte-for-byte across the upgrade; the behavioral check below additionally proves
    /// it BEHAVIORALLY via `crossSupply()`.
    bytes32 internal constant BSC_CHAIN_ID_SLOT = bytes32(uint(101));

    function setUp() public virtual override {
        super.setUp();

        // `deposit()` (BridgeTest) moves `cross` ERC20 FROM USER on the BSC side, so
        // USER needs a funded + approved balance first (mirrors
        // BridgeCrossSupplyLimitTest.setUp()'s identical need).
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 1000 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);
    }

    /// @dev Generic sign+finalize helper (unlike `CrossChainTest.crossFinalize`, which
    /// hardcodes `fromChainID = BSC_CHAIN_ID`), needed for the synthetic
    /// distinguishing-pair setup in `test_upgrade_preservesStorage`.
    function _signAndFinalizeGeneric(
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

    /// @notice Every piece of pre-upgrade state (roles, token pairs, initiate/finalize
    /// index progress, `crossSupplyLimit`, and the private `_bscChainID`) survives an
    /// upgrade to `CrossBridge` unchanged — proving the identical storage-layout
    /// redeclaration is correct.
    function test_upgrade_preservesStorage() public {
        vm.selectFork(crossForkID);

        vm.startPrank(CrossOWNER);
        bridgeCross.grantRole(Const.VERIFIER_ROLE, USER);
        uint newLimit = 777_777 ether + CROSS_FOUNDATION_INITIAL_SUPPLY;
        bridgeCross.setCrossSupplyLimit(newLimit);

        uint chainQ = 90401;
        TestToken extraToken = new TestToken("Extra", "EXTRA", 18);
        extraToken.mint(address(bridgeCross), 1000 ether);
        bridgeCross.registerToken(chainQ, true, address(extraToken), address(0x9401));

        // A DIFFERENT chain's synthetic native-like pair with a DIFFERENT `minted`
        // value than BSC_CHAIN_ID's, to rule out `crossSupply()` (post-upgrade)
        // accidentally reading the wrong `_bscChainID` / the wrong pair entirely.
        uint chainOther = chainQ + 1;
        bridgeCross.registerToken(chainOther, false, address(NATIVE_TOKEN), address(0x9402));
        vm.stopPrank();

        // Bump initiate/finalize index progress for BSC_CHAIN_ID via one ordinary
        // native round trip (also advances `minted[BSC_CHAIN_ID][NATIVE]`, i.e.
        // `crossSupply()`, away from its pure-foundation-supply baseline).
        deposit(false, 5 ether, threshold);

        assertTrue(_signAndFinalizeGeneric(chainOther, 1, address(NATIVE_TOKEN), USER, 3 ether, "", 5));
        uint otherChainMintedBefore = bridgeCross.getTokenPair(chainOther, Const.NATIVE_TOKEN).minted;
        assertEq(otherChainMintedBefore, 3 ether);

        uint crossSupplyBefore = bridgeCross.crossSupply();
        assertTrue(
            otherChainMintedBefore != crossSupplyBefore, "test requires two DISTINGUISHABLE minted values pre-upgrade"
        );

        bool hasRoleBefore = bridgeCross.hasRole(Const.VERIFIER_ROLE, USER);
        IBridgeRegistry.TokenPair memory pairBefore = bridgeCross.getTokenPair(chainQ, address(extraToken));
        uint nextInitiateBefore = bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID);
        uint nextFinalizeBefore = bridgeCross.getNextFinalizeIndex(BSC_CHAIN_ID);
        uint crossSupplyLimitBefore = bridgeCross.crossSupplyLimit();
        bytes32 bscChainIDSlotBefore = vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT);

        CrossBridge newImpl = new CrossBridge();
        vm.prank(CrossOWNER);
        bridgeCross.upgradeToAndCall(address(newImpl), bytes(""));
        CrossBridge bridgeCrossV2 = CrossBridge(payable(address(bridgeCross)));

        // (1) Slot comparison for the private, getter-less `_bscChainID`.
        bytes32 bscChainIDSlotAfter = vm.load(address(bridgeCross), BSC_CHAIN_ID_SLOT);
        assertEq(
            bscChainIDSlotAfter, bscChainIDSlotBefore, "_bscChainID storage slot must be byte-identical post-upgrade"
        );

        // (2) Behavioral proof: crossSupply() reads _tokenPairs[_bscChainID][NATIVE], so
        // it returning the SAME value (not the OTHER chain's DIFFERENT value) after the
        // upgrade proves `_bscChainID` itself round-tripped correctly, not just that its
        // raw slot bits happened to match.
        assertEq(
            bridgeCrossV2.crossSupply(),
            crossSupplyBefore,
            "crossSupply() must read the SAME _bscChainID pair post-upgrade"
        );
        assertEq(
            bridgeCross.getTokenPair(chainOther, Const.NATIVE_TOKEN).minted,
            otherChainMintedBefore,
            "an unrelated chain's pair must be untouched by the upgrade"
        );

        assertEq(bridgeCross.hasRole(Const.VERIFIER_ROLE, USER), hasRoleBefore, "roles must survive the upgrade");

        IBridgeRegistry.TokenPair memory pairAfter = bridgeCross.getTokenPair(chainQ, address(extraToken));
        assertEq(pairAfter.localToken, pairBefore.localToken);
        assertEq(pairAfter.remoteToken, pairBefore.remoteToken);
        assertEq(pairAfter.isOrigin, pairBefore.isOrigin);
        assertEq(pairAfter.deposited, pairBefore.deposited);
        assertEq(pairAfter.minted, pairBefore.minted);

        assertEq(
            bridgeCross.getNextInitiateIndex(BSC_CHAIN_ID),
            nextInitiateBefore,
            "initiate index must survive the upgrade"
        );
        assertEq(
            bridgeCross.getNextFinalizeIndex(BSC_CHAIN_ID),
            nextFinalizeBefore,
            "finalize index must survive the upgrade"
        );
        assertEq(bridgeCrossV2.crossSupplyLimit(), crossSupplyLimitBefore, "crossSupplyLimit must survive the upgrade");
    }

    /// @notice Post-upgrade, `initialize(...)` (inherited from `BaseBridge`, overridden
    /// in `CrossBridge` as a permanent revert stub) is sealed shut.
    function test_upgrade_initializeSealed() public {
        vm.selectFork(crossForkID);
        CrossBridge newImpl = new CrossBridge();
        vm.prank(CrossOWNER);
        bridgeCross.upgradeToAndCall(address(newImpl), bytes(""));
        CrossBridge bridgeCrossV2 = CrossBridge(payable(address(bridgeCross)));

        vm.expectRevert(CrossBridge.Disabled.selector);
        bridgeCrossV2.initialize(CrossOWNER, REWARD, threshold);
    }

    /// @notice `CrossBridge` does not inherit `CrossBridge`, so `initializeCrossBridge`
    /// does not exist on it at all post-upgrade — not merely reverting, genuinely absent
    /// (a raw call with its selector fails to resolve to any function and reverts via
    /// the EVM's default no-matching-selector/no-fallback behavior).
    function test_upgrade_initializeCrossBridgeGone() public {
        vm.selectFork(crossForkID);
        CrossBridge newImpl = new CrossBridge();
        vm.prank(CrossOWNER);
        bridgeCross.upgradeToAndCall(address(newImpl), bytes(""));

        bytes memory data = abi.encodeWithSignature(
            "initializeCrossBridge(address,address,uint8,uint256,address,uint256)",
            CrossOWNER,
            REWARD,
            threshold,
            BSC_CHAIN_ID,
            address(cross),
            CROSS_FOUNDATION_INITIAL_SUPPLY
        );
        (bool ok,) = address(bridgeCross).call(data);
        assertFalse(ok, "initializeCrossBridge selector must not exist on CrossBridge");
    }

    /// @notice `crossSupplyLimit`'s circuit-breaker (unchanged CrossBridge logic,
    /// only inherited via CrossBridge's override of `_checkFinalizeAmount`) still
    /// enforces correctly post-upgrade.
    function test_upgrade_crossSupplyLimitStillEnforced() public {
        vm.selectFork(crossForkID);
        CrossBridge newImpl = new CrossBridge();
        vm.prank(CrossOWNER);
        bridgeCross.upgradeToAndCall(address(newImpl), bytes(""));
        CrossBridge bridgeCrossV2 = CrossBridge(payable(address(bridgeCross)));

        vm.prank(CrossOWNER);
        bridgeCrossV2.setCrossSupplyLimit(10 ether + CROSS_FOUNDATION_INITIAL_SUPPLY);

        // Under the limit: succeeds normally.
        deposit(false, 5 ether, threshold);

        // Exceeds the limit: falls to pending with CrossSupplyLimitExceeded, exactly
        // like the pre-upgrade CrossBridge behavior (BridgeCrossSupplyLimit.t.sol).
        vm.selectFork(bscForkID);
        (uint index2,) = bscBridge(address(cross), USER, USER, 10 ether, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index2, address(NATIVE_TOKEN), USER, 10 ether, threshold);
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index2).status
                == Const.FinalizeStatus.CrossSupplyLimitExceeded,
            "crossSupplyLimit must still be enforced post-upgrade"
        );
    }
}
