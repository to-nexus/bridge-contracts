// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2} from "../src/token/CrossMintableERC20V2.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {ICrossMintableERC20V2Code} from "../src/token/ICrossMintableERC20V2Code.sol";

import {CrossMintableERC20V2CodeScript} from "../script/CrossMintableERC20V2Code.s.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import {Errors} from "@openzeppelin/contracts/utils/Errors.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/Test.sol";

/**
 * @title CrossMintableERC20V2Test
 * @notice Covers the BeaconProxy + UUPS factory design end to end: the legacy creation path, the
 *         beacon-proxy shape of created tokens, address-prediction stability, factory
 *         upgradeability, per-pair uniqueness across every creation variant, owner-managed
 *         token pass-through, and the DEFAULT_ADMIN_ROLE-vs-ADMIN_ROLE ownership transfer
 *         runbook.
 * @dev Most of this file is deliberately scoped to `CrossMintableERC20V2` +
 *      `CrossMintableERC20V2Code` in isolation — `bridge` is a bare address (as in
 *      `HyperMintableERC20Test`), not a real `BaseBridge` instance. The last two groups of tests
 *      step outside that isolation on purpose: one deploys a real (non-forked,
 *      locally-initialized) `BaseBridge` proxy to prove a cross-factory duplicate-pair gap at
 *      the `BridgeRegistry` level, and another instantiates `CrossMintableERC20V2CodeScript`
 *      directly (no `forge script` invocation needed — `preflightCheckDuplicateRemoteToken` is a
 *      plain `view` function) against a minimal mock registry that only implements
 *      `allTokenPairs`.
 */
contract CrossMintableERC20V2Test is Test {
    /// @dev Declared locally (not imported from IERC20) purely so `vm.expectEmit` can match the
    /// inherited OZ ERC20 `Transfer` event by signature.
    event Transfer(address indexed from, address indexed to, uint value);

    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    // ERC-7201 / ERC-1967 slot literals — the test's OWN copies (never read from a getter), the
    // same independence discipline `HyperMintableERC20Test` uses for its namespaced-slot check.
    bytes32 internal constant TOKEN_STORAGE_SLOT = 0x9581ab046f3e11ba7ecf3d09df64f89f69c3a2a7a7e6f470b2498e2d7b1f5700;
    bytes32 internal constant FACTORY_STORAGE_SLOT = 0x2b6c7dfd4e9330f706f3d4045fa01a8903ed43939ff7d252d6806f39e1487c00;
    bytes32 internal constant ERC1967_IMPLEMENTATION_SLOT =
        0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
    bytes32 internal constant ERC1967_BEACON_SLOT = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;

    CrossMintableERC20V2 public tokenImplementation;
    CrossMintableERC20V2Code public codeImplementation;
    CrossMintableERC20V2Code public code;
    CrossMintableERC20V2 public token;

    uint internal constant FACTORY_OWNER_PK = uint(bytes32("factoryOwner"));
    uint internal constant BRIDGE_PK = uint(bytes32("bridge"));
    uint internal constant USER_PK = uint(bytes32("user"));

    address public factoryOwner; // factory ADMIN_ROLE + DEFAULT_ADMIN_ROLE holder
    address public bridge; // BRIDGE_ROLE on the factory, MINTER_ROLE on tokens created via the legacy path
    address public user;

    uint internal constant REMOTE_CHAIN_ID = 998;
    address internal constant REMOTE_TOKEN = address(0xBEEF);
    string internal constant SYMBOL = "TT";
    uint8 internal constant DECIMALS = 18;

    uint internal constant INITIAL_SUPPLY = 1_000_000_000 ether;

    function setUp() public {
        factoryOwner = vm.addr(FACTORY_OWNER_PK);
        bridge = vm.addr(BRIDGE_PK);
        user = vm.addr(USER_PK);

        tokenImplementation = new CrossMintableERC20V2();
        codeImplementation = new CrossMintableERC20V2Code();

        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(CrossMintableERC20V2Code.initialize, (factoryOwner, bridge, address(tokenImplementation)))
        );
        code = CrossMintableERC20V2Code(address(proxy));
        console.log("CrossMintableERC20V2Code", address(code));

        vm.prank(bridge);
        token = CrossMintableERC20V2(code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS));
        console.log("CrossMintableERC20V2", address(token));
    }

    function _deployCode(address factoryOwnerAddr, address bridgeAddr) internal returns (CrossMintableERC20V2Code freshCode) {
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(
                CrossMintableERC20V2Code.initialize, (factoryOwnerAddr, bridgeAddr, address(tokenImplementation))
            )
        );
        freshCode = CrossMintableERC20V2Code(address(proxy));
    }

    function _signPermit(uint pk, address owner, address spender, uint value, uint deadline)
        internal
        view
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, owner, spender, value, token.nonces(owner), deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(token.DOMAIN_SEPARATOR(), h);
        (v, r, s) = vm.sign(pk, hash);
    }

    // =====================================================================
    // Legacy invariants
    // =====================================================================

    /// `createCrossMintableERC20`'s selector is frozen at `0xf88d3d42` — the whole point of
    /// keeping its signature untouched across this conversion.
    function test_legacySelector_isFrozen() public pure {
        assertEq(ICrossMintableERC20Code.createCrossMintableERC20.selector, bytes4(0xf88d3d42));
    }

    /// Legacy path derives name/symbol exactly as before: "Cross Bridge <SYM>" / "<SYM>x".
    function test_legacyPath_derivedNameAndSymbol() public view {
        assertEq(token.name(), string(abi.encodePacked("Cross Bridge ", SYMBOL)));
        assertEq(token.symbol(), string(abi.encodePacked(SYMBOL, "x")));
        assertEq(token.decimals(), DECIMALS);
    }

    /// The legacy path's `minter` is the actual caller (`_msgSender()`, the bridge), which
    /// receives `MINTER_ROLE` on the created token.
    function test_legacyPath_callerGetsMinterRole() public view {
        assertTrue(token.hasRole(Const.MINTER_ROLE, bridge));
    }

    /// Every token's `defaultAdmin()` is the FACTORY PROXY address itself — not a
    /// separately stored `tokenAdmin`. There is no `tokenAdmin()` getter to compare against.
    function test_tokenDefaultAdmin_isFactoryItself() public view {
        assertEq(token.defaultAdmin(), address(code));
    }

    /// Missing `BRIDGE_ROLE` reverts the legacy path.
    function test_legacyPath_requiresBridgeRole() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.BRIDGE_ROLE)
        );
        code.createCrossMintableERC20(REMOTE_CHAIN_ID + 1, REMOTE_TOKEN, SYMBOL, DECIMALS);
    }

    // =====================================================================
    // BeaconProxy shape
    // =====================================================================

    /// The created token really is a `BeaconProxy`: its ERC-1967 beacon slot holds this
    /// factory's beacon address, and it has no ERC-1967 implementation slot of its own set
    /// (`BeaconProxy` never writes that slot — only the beacon does).
    function test_createdToken_isRealBeaconProxy() public view {
        bytes32 beaconSlotValue = vm.load(address(token), ERC1967_BEACON_SLOT);
        assertEq(address(uint160(uint(beaconSlotValue))), code.beacon());
        assertEq(vm.load(address(token), ERC1967_IMPLEMENTATION_SLOT), bytes32(0));
    }

    /// Beacon upgrade is reflected in an ALREADY-issued token immediately, while its balances,
    /// roles and `decimals` (ERC-7201 storage) are preserved untouched.
    function test_beaconUpgrade_reflectsOnExistingToken_preservingState() public {
        vm.prank(bridge);
        token.mint(user, 10 ether);
        vm.prank(factoryOwner);
        code.grantTokenRole(address(token), Const.MINTER_ROLE, user);

        CrossMintableERC20V2Mock newImpl = new CrossMintableERC20V2Mock();
        vm.prank(factoryOwner);
        code.upgradeBeacon(address(newImpl));

        assertEq(token.balanceOf(user), 10 ether);
        assertEq(token.decimals(), DECIMALS);
        assertTrue(token.hasRole(Const.MINTER_ROLE, user));
        assertEq(CrossMintableERC20V2Mock(address(token)).version(), 2);
    }

    /// A beacon upgrade never moves a future CREATE2 prediction — the initcode only ever
    /// embeds the beacon's own (fixed) address, never the implementation it currently resolves to.
    function test_beaconUpgrade_doesNotMovePrediction() public {
        uint remoteChainID2 = REMOTE_CHAIN_ID + 2;
        address predictedBefore = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        CrossMintableERC20V2Mock newImpl = new CrossMintableERC20V2Mock();
        vm.prank(factoryOwner);
        code.upgradeBeacon(address(newImpl));

        address predictedAfter = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedBefore, predictedAfter);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS);
        assertEq(actual, predictedBefore);
        assertEq(CrossMintableERC20V2Mock(actual).version(), 2);
    }

    /// The token LOGIC contract can never be initialized directly (`_disableInitializers`).
    function test_tokenImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        tokenImplementation.initialize(factoryOwner, bridge, "X", "X", 18);
    }

    /// The token PROXY cannot be initialized a second time.
    function test_tokenProxy_reinitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        token.initialize(factoryOwner, bridge, "X", "X", 18);
    }

    // =====================================================================
    // Address prediction
    // =====================================================================

    /// Legacy-path prediction matches the actual `createCrossMintableERC20` result.
    function test_legacyPath_predictionMatchesActual() public view {
        address predicted = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predicted, address(token));
    }

    /// New (`createMintableERC20`) path prediction matches the actual result.
    function test_newPath_predictionMatchesActual() public {
        uint remoteChainID3 = REMOTE_CHAIN_ID + 3;
        address minter = makeAddr("explicitMinter");

        address predicted =
            code.computeTokenAddressWithName(remoteChainID3, REMOTE_TOKEN, "Name X", "SYMX", 7, minter);

        vm.prank(factoryOwner);
        address actual = code.createMintableERC20(remoteChainID3, REMOTE_TOKEN, "Name X", "SYMX", 7, minter);

        assertEq(predicted, actual);
    }

    /// Transferring the factory's OWN `defaultAdmin()` (its `DEFAULT_ADMIN_ROLE`, distinct
    /// from `ADMIN_ROLE`) changes nothing about future predictions: `_initCode` always
    /// embeds `address(this)` (the factory's own fixed address), never `defaultAdmin()`.
    function test_addressPrediction_stableAcrossFactoryDefaultAdminTransfer() public {
        address newFactoryOwner = makeAddr("newFactoryOwner");
        uint remoteChainID4 = REMOTE_CHAIN_ID + 4;

        address predictedBefore = code.computeTokenAddress(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(factoryOwner);
        code.beginDefaultAdminTransfer(newFactoryOwner);
        vm.warp(block.timestamp + 1);
        vm.prank(newFactoryOwner);
        code.acceptDefaultAdminTransfer();

        assertEq(code.defaultAdmin(), newFactoryOwner);

        address predictedAfter = code.computeTokenAddress(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedAfter, predictedBefore);
    }

    /// Even after the factory's `defaultAdmin()` moves, a freshly
    /// created token's OWN `defaultAdmin()` is still the factory's (unchanged) proxy address —
    /// not the factory's new `defaultAdmin()`, and not the old one either. Prediction stays
    /// pinned to the same value before and after.
    function test_factoryDefaultAdminTransfer_doesNotAffectTokenDefaultAdminOrPrediction() public {
        address newFactoryOwner = makeAddr("newFactoryOwner13a");
        uint remoteChainID5 = REMOTE_CHAIN_ID + 5;

        address predictedBefore = code.computeTokenAddress(remoteChainID5, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(factoryOwner);
        code.beginDefaultAdminTransfer(newFactoryOwner);
        vm.warp(block.timestamp + 1);
        vm.prank(newFactoryOwner);
        code.acceptDefaultAdminTransfer();

        address predictedAfter = code.computeTokenAddress(remoteChainID5, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedAfter, predictedBefore);

        // `bridge` still holds BRIDGE_ROLE (unaffected by the factory's defaultAdmin transfer),
        // so the legacy path can still be exercised to confirm the actual deployed address.
        vm.prank(bridge);
        address newToken = code.createCrossMintableERC20(remoteChainID5, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(newToken, predictedBefore);
        assertEq(CrossMintableERC20V2(newToken).defaultAdmin(), address(code));
    }

    // =====================================================================
    // New explicit path
    // =====================================================================

    /// `createMintableERC20` reflects its arguments exactly (no derivation).
    function test_createMintableERC20_reflectsArguments() public {
        address minter = makeAddr("explicitMinter14");
        uint remoteChainID6 = REMOTE_CHAIN_ID + 6;

        vm.prank(factoryOwner);
        address tokenAddress =
            code.createMintableERC20(remoteChainID6, REMOTE_TOKEN, "My Custom Token", "MCT", 9, minter);

        CrossMintableERC20V2 customToken = CrossMintableERC20V2(tokenAddress);
        assertEq(customToken.name(), "My Custom Token");
        assertEq(customToken.symbol(), "MCT");
        assertEq(customToken.decimals(), 9);
        assertTrue(customToken.hasRole(Const.MINTER_ROLE, minter));
        assertTrue(code.isCrossMintableERC20(tokenAddress));
    }

    /// `createMintableERC20` is gated to `ADMIN_ROLE`.
    function test_createMintableERC20_requiresAdminRole() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.createMintableERC20(REMOTE_CHAIN_ID + 7, REMOTE_TOKEN, "X", "X", 18, user);
    }

    /// `minter == address(0)` on the new path means `MINTER_ROLE` is granted to nobody.
    function test_createMintableERC20_zeroMinter_grantsNoRole() public {
        vm.prank(factoryOwner);
        address tokenAddress = code.createMintableERC20(REMOTE_CHAIN_ID + 8, REMOTE_TOKEN, "X", "X", 18, address(0));

        CrossMintableERC20V2 customToken = CrossMintableERC20V2(tokenAddress);
        assertFalse(customToken.hasRole(Const.MINTER_ROLE, address(0)));
        assertFalse(customToken.hasRole(Const.MINTER_ROLE, factoryOwner));
        assertFalse(customToken.hasRole(Const.MINTER_ROLE, address(code)));
    }

    // =====================================================================
    // Factory upgradeability
    // =====================================================================

    /// The factory PROXY cannot be initialized a second time.
    function test_factoryProxy_reinitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        code.initialize(factoryOwner, bridge, address(tokenImplementation));
    }

    /// Deployer != `initialOwner` must not revert the factory's atomic proxy
    /// initialization — the trap the legacy (non-upgradeable) `CrossMintableERC20V2Code`'s
    /// constructor falls into by using public `grantRole` instead of `_grantRole`.
    function test_initialize_succeedsWhenDeployerIsNotInitialOwner() public {
        address deployer = makeAddr("someoneElse18");
        vm.prank(deployer);
        CrossMintableERC20V2Code freshCode = _deployCode(factoryOwner, bridge);

        assertTrue(freshCode.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertFalse(freshCode.hasRole(Const.ADMIN_ROLE, deployer));
    }

    /// The factory LOGIC contract can never be initialized directly.
    function test_codeImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        codeImplementation.initialize(factoryOwner, bridge, address(tokenImplementation));
    }

    /// `upgradeToAndCall` without `ADMIN_ROLE` reverts.
    function test_factoryUpgrade_requiresAdminRole() public {
        CrossMintableERC20V2CodeMock newCodeImpl = new CrossMintableERC20V2CodeMock();
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.upgradeToAndCall(address(newCodeImpl), bytes(""));
    }

    /// Factory (UUPS) upgrade preserves storage: `beacon`, `isCrossMintableERC20` and
    /// `tokenForPair` all survive.
    function test_factoryUpgrade_preservesState() public {
        assertTrue(code.isCrossMintableERC20(address(token)));
        address beaconBefore = code.beacon();
        address tokenForPairBefore = code.tokenForPair(REMOTE_CHAIN_ID, REMOTE_TOKEN);

        CrossMintableERC20V2CodeMock newImpl = new CrossMintableERC20V2CodeMock();
        vm.prank(factoryOwner);
        code.upgradeToAndCall(address(newImpl), bytes(""));

        assertEq(code.beacon(), beaconBefore);
        assertTrue(code.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertTrue(code.hasRole(Const.BRIDGE_ROLE, bridge));
        assertTrue(code.isCrossMintableERC20(address(token)));
        assertEq(code.tokenForPair(REMOTE_CHAIN_ID, REMOTE_TOKEN), tokenForPairBefore);
        assertEq(CrossMintableERC20V2CodeMock(address(code)).version(), 2);
    }

    // =====================================================================
    // Common
    // =====================================================================

    /// `CrossMintableERC20Created` carries the right fields on the legacy path.
    function test_creationEvent_legacyPath() public {
        uint remoteChainID9 = REMOTE_CHAIN_ID + 9;
        address predicted = code.computeTokenAddress(remoteChainID9, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.expectEmit(true, true, true, true, address(code));
        emit ICrossMintableERC20V2Code.CrossMintableERC20Created(remoteChainID9, REMOTE_TOKEN, predicted);

        vm.prank(bridge);
        code.createCrossMintableERC20(remoteChainID9, REMOTE_TOKEN, SYMBOL, DECIMALS);
    }

    /// `CrossMintableERC20Created` carries the right fields on the new explicit path.
    function test_creationEvent_newPath() public {
        uint remoteChainID10 = REMOTE_CHAIN_ID + 10;
        address minter = makeAddr("explicitMinter22b");
        address predicted =
            code.computeTokenAddressWithName(remoteChainID10, REMOTE_TOKEN, "X", "X", 18, minter);

        vm.expectEmit(true, true, true, true, address(code));
        emit ICrossMintableERC20V2Code.CrossMintableERC20Created(remoteChainID10, REMOTE_TOKEN, predicted);

        vm.prank(factoryOwner);
        code.createMintableERC20(remoteChainID10, REMOTE_TOKEN, "X", "X", 18, minter);
    }

    /// `isCrossMintableERC20` is true right after creation, false for an unrelated address.
    function test_isCrossMintableERC20_trueForCreatedToken() public view {
        assertTrue(code.isCrossMintableERC20(address(token)));
        assertFalse(code.isCrossMintableERC20(address(0xDEAD)));
    }

    /// Because `tokenForPair` is checked before CREATE2 ever runs, retrying the SAME pair —
    /// even with byte-for-byte identical arguments (what would otherwise be a raw CREATE2 salt
    /// collision) — is rejected by `CrossMintableERC20V2CodePairAlreadyCreated`, not
    /// `Errors.FailedDeployment`: the pair guard is strictly in front of the raw CREATE2 check.
    function test_samePairSameArgs_revertsWithPairAlreadyCreated_notRawCreate2Collision() public {
        vm.prank(bridge);
        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodePairAlreadyCreated.selector,
                REMOTE_CHAIN_ID,
                REMOTE_TOKEN,
                address(token)
            )
        );
        code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS);
    }

    // =====================================================================
    // One token per remote pair, enforced across every creation variant
    // =====================================================================

    /// Same pair, legacy path again but with a DIFFERENT symbol (would have produced a
    /// different CREATE2 address under a pair-guard-less factory) — still rejected.
    function test_samePair_differentSymbol_stillReverts() public {
        vm.prank(bridge);
        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodePairAlreadyCreated.selector,
                REMOTE_CHAIN_ID,
                REMOTE_TOKEN,
                address(token)
            )
        );
        code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, "OTHER", DECIMALS);
    }

    /// Same pair, crossing from the legacy path to the new explicit path with a
    /// different name/symbol/minter — still rejected.
    function test_samePair_crossPathDifferentNameAndMinter_stillReverts() public {
        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodePairAlreadyCreated.selector,
                REMOTE_CHAIN_ID,
                REMOTE_TOKEN,
                address(token)
            )
        );
        code.createMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, "Totally Different", "TD", 6, makeAddr("otherMinter"));
    }

    /// Same pair via the new path twice, second attempt varying only the minter — still
    /// rejected (guards a pair, not an (args) tuple).
    function test_samePair_newPathDifferentMinterOnly_stillReverts() public {
        uint remoteChainID11 = REMOTE_CHAIN_ID + 11;
        vm.startPrank(factoryOwner);
        address first = code.createMintableERC20(remoteChainID11, REMOTE_TOKEN, "X", "X", 18, makeAddr("minterA"));

        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodePairAlreadyCreated.selector,
                remoteChainID11,
                REMOTE_TOKEN,
                first
            )
        );
        code.createMintableERC20(remoteChainID11, REMOTE_TOKEN, "X", "X", 18, makeAddr("minterB"));
        vm.stopPrank();
    }

    /// A DIFFERENT `remoteToken` on the same `remoteChainID` is unaffected — the guard
    /// is keyed on the pair, not the chain alone.
    function test_differentRemoteToken_sameChain_succeeds() public {
        vm.prank(bridge);
        address second = code.createCrossMintableERC20(REMOTE_CHAIN_ID, address(0xCAFE), SYMBOL, DECIMALS);
        assertTrue(code.isCrossMintableERC20(second));
        assertEq(code.tokenForPair(REMOTE_CHAIN_ID, address(0xCAFE)), second);
    }

    // =====================================================================
    // Owner-managed token pass-through
    // =====================================================================

    /// The beacon's `owner()` is the factory PROXY itself, not an EOA — so upgrading it
    /// must go through `code.upgradeBeacon`, never a direct call, even from the `ADMIN_ROLE` EOA.
    function test_beaconOwner_isFactoryItself_directCallReverts() public {
        UpgradeableBeacon tokenBeacon = UpgradeableBeacon(code.beacon());
        assertEq(tokenBeacon.owner(), address(code));

        CrossMintableERC20V2Mock newImpl = new CrossMintableERC20V2Mock();
        vm.prank(factoryOwner);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, factoryOwner));
        tokenBeacon.upgradeTo(address(newImpl));
    }

    /// `grantTokenRole`/`revokeTokenRole` delegate to the token on the factory's behalf (the
    /// factory is the token's `defaultAdmin()`) and are gated to `ADMIN_ROLE` + known tokens.
    function test_grantAndRevokeTokenRole_delegatesAndIsGated() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.grantTokenRole(address(token), Const.MINTER_ROLE, user);

        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(ICrossMintableERC20V2Code.CrossMintableERC20V2CodeUnknownToken.selector, address(0xDEAD))
        );
        code.grantTokenRole(address(0xDEAD), Const.MINTER_ROLE, user);

        vm.prank(factoryOwner);
        code.grantTokenRole(address(token), Const.MINTER_ROLE, user);
        assertTrue(token.hasRole(Const.MINTER_ROLE, user));

        vm.prank(factoryOwner);
        code.revokeTokenRole(address(token), Const.MINTER_ROLE, user);
        assertFalse(token.hasRole(Const.MINTER_ROLE, user));
    }

    /// `beginTokenDefaultAdminTransfer` is the escape hatch: it only BEGINS the token's own
    /// two-step transfer — the new admin must still separately `acceptDefaultAdminTransfer()` on
    /// the token itself. Once accepted, the factory can no longer manage that token's roles.
    function test_beginTokenDefaultAdminTransfer_escapeHatch() public {
        address newTokenAdmin = makeAddr("newTokenAdmin");

        vm.prank(factoryOwner);
        code.beginTokenDefaultAdminTransfer(address(token), newTokenAdmin);

        // Still the factory until accepted.
        assertEq(token.defaultAdmin(), address(code));

        vm.warp(block.timestamp + 1);
        vm.prank(newTokenAdmin);
        token.acceptDefaultAdminTransfer();

        assertEq(token.defaultAdmin(), newTokenAdmin);

        // The factory (no longer the token's defaultAdmin) can no longer grant token roles.
        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, address(code), bytes32(0))
        );
        code.grantTokenRole(address(token), Const.MINTER_ROLE, user);

        // The new admin manages the token directly from here on.
        vm.prank(newTokenAdmin);
        token.grantRole(Const.MINTER_ROLE, user);
        assertTrue(token.hasRole(Const.MINTER_ROLE, user));
    }

    // =====================================================================
    // Ownership transfer runbook — DEFAULT_ADMIN_ROLE != ADMIN_ROLE
    // =====================================================================

    /// Before step 3 (grantRole(ADMIN_ROLE, newOwner)) the new owner cannot perform any
    /// ADMIN_ROLE-gated action, even after accepting DEFAULT_ADMIN_ROLE (step 1-2 alone are not
    /// enough). After step 4 (revokeRole(ADMIN_ROLE, oldOwner)) the OLD owner can no longer
    /// either. Probed with `createMintableERC20`, an ordinary ADMIN_ROLE-gated action.
    function test_ownershipTransferRunbook_gatesAdminRoleActions() public {
        address newOwner = makeAddr("newOwner2c");

        // Step 1-2: begin + accept DEFAULT_ADMIN_ROLE transfer.
        vm.prank(factoryOwner);
        code.beginDefaultAdminTransfer(newOwner);
        vm.warp(block.timestamp + 1);
        vm.prank(newOwner);
        code.acceptDefaultAdminTransfer();
        assertEq(code.defaultAdmin(), newOwner);

        // Before step 3: newOwner has DEFAULT_ADMIN_ROLE but NOT ADMIN_ROLE -> still reverts.
        vm.prank(newOwner);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, newOwner, Const.ADMIN_ROLE)
        );
        code.createMintableERC20(REMOTE_CHAIN_ID + 12, REMOTE_TOKEN, "X", "X", 18, address(0));

        // Step 3: newOwner (now DEFAULT_ADMIN_ROLE holder) grants itself ADMIN_ROLE.
        vm.prank(newOwner);
        code.grantRole(Const.ADMIN_ROLE, newOwner);

        // newOwner can now act.
        vm.prank(newOwner);
        code.createMintableERC20(REMOTE_CHAIN_ID + 12, REMOTE_TOKEN, "X", "X", 18, address(0));

        // oldOwner (factoryOwner) still holds ADMIN_ROLE at this point -> still succeeds.
        vm.prank(factoryOwner);
        code.createMintableERC20(REMOTE_CHAIN_ID + 13, REMOTE_TOKEN, "X", "X", 18, address(0));

        // Step 4: newOwner revokes the old owner's ADMIN_ROLE.
        vm.prank(newOwner);
        code.revokeRole(Const.ADMIN_ROLE, factoryOwner);

        // oldOwner can no longer act.
        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, factoryOwner, Const.ADMIN_ROLE
            )
        );
        code.createMintableERC20(REMOTE_CHAIN_ID + 14, REMOTE_TOKEN, "X", "X", 18, address(0));
    }

    // =====================================================================
    // Cross-factory duplicate registration + deploy-script preflight
    // =====================================================================

    /// @dev Deploys a real (non-forked) `BaseBridge` behind an `ERC1967Proxy`, granting
    /// `bridgeOwner` `ADMIN_ROLE` + `EDITOR_ROLE` on top of the `DEFAULT_ADMIN_ROLE` its own
    /// `initialize` grants — enough to call `setCrossMintableERC20Code` and `createToken`.
    function _deployBridge(address bridgeOwner) internal returns (BaseBridge bridgeInstance) {
        BaseBridge bridgeImpl = new BaseBridge();
        ERC1967Proxy proxy =
            new ERC1967Proxy(address(bridgeImpl), abi.encodeCall(BaseBridge.initialize, (bridgeOwner, payable(bridgeOwner), 1)));
        bridgeInstance = BaseBridge(address(proxy));

        vm.startPrank(bridgeOwner);
        bridgeInstance.grantRole(Const.ADMIN_ROLE, bridgeOwner);
        bridgeInstance.grantRole(Const.EDITOR_ROLE, bridgeOwner);
        vm.stopPrank();
    }

    /// `BridgeRegistry._registerToken` only guards duplicates by `localToken`
    /// (`src/abstract/BridgeRegistry.sol:468`), so replacing the bridge's factory and
    /// re-creating a token for the SAME `(remoteChainID, remoteToken)` is NOT blocked on-chain —
    /// a second pair is registered alongside the first instead of `RegistryExistToken`. This
    /// pins the exact gap the deploy script's preflight check below exists to catch
    /// operationally.
    function test_bridgeRegistry_allowsDuplicatePair_acrossFactoryReplacement() public {
        address bridgeOwner = makeAddr("bridgeOwner25");
        BaseBridge bridgeInstance = _deployBridge(bridgeOwner);

        CrossMintableERC20V2Code factory1 = _deployCode(factoryOwner, address(bridgeInstance));
        CrossMintableERC20V2Code factory2 = _deployCode(factoryOwner, address(bridgeInstance));

        uint remoteChainID25 = REMOTE_CHAIN_ID + 25;
        address remoteToken25 = makeAddr("remoteToken25");

        vm.startPrank(bridgeOwner);
        bridgeInstance.setCrossMintableERC20Code(ICrossMintableERC20Code(address(factory1)));
        address localToken1 = bridgeInstance.createToken(remoteChainID25, remoteToken25, "AAA", 18);

        bridgeInstance.setCrossMintableERC20Code(ICrossMintableERC20Code(address(factory2)));
        address localToken2 = bridgeInstance.createToken(remoteChainID25, remoteToken25, "BBB", 18);
        vm.stopPrank();

        assertTrue(localToken1 != localToken2);
        assertTrue(factory1.isCrossMintableERC20(localToken1));
        assertTrue(factory2.isCrossMintableERC20(localToken2));

        IBridgeRegistry.TokenPair[] memory pairs = bridgeInstance.allTokenPairs(remoteChainID25);
        assertEq(pairs.length, 2);
        assertEq(pairs[0].localToken, localToken1);
        assertEq(pairs[0].remoteToken, remoteToken25);
        assertEq(pairs[1].localToken, localToken2);
        assertEq(pairs[1].remoteToken, remoteToken25);
    }

    /// The deploy script's `preflightCheckDuplicateRemoteToken` reverts when the target
    /// `remoteChainID` already has a pair for `remoteToken` (regardless of `localToken`), and
    /// does NOT revert for an unrelated `remoteToken` on the same chain. Uses a minimal mock that
    /// implements only `allTokenPairs`, cast to `BaseBridge` — Solidity dispatches by selector,
    /// so the mock never needs to actually BE a `BaseBridge`.
    function test_scriptPreflight_revertsOnExistingRemoteToken() public {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        MockAllTokenPairsRegistry mockRegistry = new MockAllTokenPairsRegistry();

        address existingLocalToken26 = makeAddr("existingLocalToken26");
        address existingRemoteToken26 = makeAddr("existingRemoteToken26");
        IBridgeRegistry.TokenPair[] memory pairs = new IBridgeRegistry.TokenPair[](1);
        pairs[0] = IBridgeRegistry.TokenPair({
            localToken: existingLocalToken26,
            remoteToken: existingRemoteToken26,
            isOrigin: false,
            paused: false,
            pendingAmount: 0,
            deposited: 0,
            minted: 0
        });
        mockRegistry.setPairs(pairs);

        BaseBridge asBridge = BaseBridge(address(mockRegistry));

        vm.expectRevert(
            "preflightCheckDuplicateRemoteToken: remoteToken already registered for this remoteChainID under a different localToken"
        );
        scriptContract.preflightCheckDuplicateRemoteToken(asBridge, REMOTE_CHAIN_ID, existingRemoteToken26);

        // A different remoteToken on the same remoteChainID must NOT revert.
        scriptContract.preflightCheckDuplicateRemoteToken(asBridge, REMOTE_CHAIN_ID, makeAddr("unrelatedRemoteToken26"));
    }

    // =====================================================================
    // `revokeOldAdmin` identity checks and the view/broadcast split for the three
    // ownership-transfer status functions
    // =====================================================================

    /// `beginTransferStatus` / `acceptTransferStatus` / `revokeOldAdminStatus` return the
    /// right `TransferStage` as the underlying on-chain state is walked through every stage of a
    /// real 4-step transfer, using a factory dedicated to this test (`oldOwner27`/`newOwner27`)
    /// so it doesn't interact with the shared `setUp` fixture's `code`.
    function test_statusFunctions_matchStateAcrossStages() public {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        address oldOwner27 = makeAddr("oldOwner27");
        address newOwner27 = makeAddr("newOwner27");
        address wrongOwner27 = makeAddr("wrongOwner27");

        CrossMintableERC20V2Code fresh = _deployCode(oldOwner27, bridge);

        // --- beginTransferStatus ---
        // Fresh factory: defaultAdmin()==oldOwner27, nothing pending -> Ready.
        (CrossMintableERC20V2CodeScript.TransferStage stage, string memory reason) =
            scriptContract.beginTransferStatus(address(fresh), oldOwner27, newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Ready));

        // Wrong oldOwner argument (defaultAdmin() is neither the given oldOwner nor newOwner)
        // -> Inconsistent, with a non-empty diagnostic reason.
        (stage, reason) = scriptContract.beginTransferStatus(address(fresh), wrongOwner27, newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);

        // Begin for real: pendingDefaultAdmin()==newOwner27, defaultAdmin() still oldOwner27
        // -> AlreadyDone (already begun).
        vm.prank(oldOwner27);
        fresh.beginDefaultAdminTransfer(newOwner27);

        (stage,) = scriptContract.beginTransferStatus(address(fresh), oldOwner27, newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.AlreadyDone));

        // --- acceptTransferStatus ---
        // pendingDefaultAdmin()==newOwner27, not yet accepted -> Ready.
        (stage,) = scriptContract.acceptTransferStatus(address(fresh), newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Ready));

        // An unrelated address with nothing pending for it -> Inconsistent.
        (stage, reason) = scriptContract.acceptTransferStatus(address(fresh), wrongOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);

        // Accept for real: defaultAdmin()==newOwner27 now, but ADMIN_ROLE not granted yet
        // -> still Ready (only the grant sub-step remains).
        vm.warp(block.timestamp + 1);
        vm.prank(newOwner27);
        fresh.acceptDefaultAdminTransfer();

        (stage,) = scriptContract.acceptTransferStatus(address(fresh), newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Ready));

        // Grant ADMIN_ROLE: postcondition fully satisfied -> AlreadyDone.
        vm.prank(newOwner27);
        fresh.grantRole(Const.ADMIN_ROLE, newOwner27);

        (stage,) = scriptContract.acceptTransferStatus(address(fresh), newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.AlreadyDone));

        // --- revokeOldAdminStatus ---
        // defaultAdmin()==newOwner27, both oldOwner27 and newOwner27 hold ADMIN_ROLE -> Ready.
        (stage,) = scriptContract.revokeOldAdminStatus(address(fresh), oldOwner27, newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Ready));

        // Revoke for real: oldOwner27 now lacks ADMIN_ROLE -> Inconsistent, NEVER AlreadyDone
        // (see the tests below for why AlreadyDone must never be reachable here).
        vm.prank(newOwner27);
        fresh.revokeRole(Const.ADMIN_ROLE, oldOwner27);

        (stage, reason) = scriptContract.revokeOldAdminStatus(address(fresh), oldOwner27, newOwner27);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);
    }

    /// Pins the exact bug this identity check exists to prevent: after
    /// a real transfer completes (`realOldOwner28` -> `newOwner28`) with `realOldOwner28` still
    /// holding `ADMIN_ROLE` (the operator hasn't revoked it yet), calling
    /// `revokeOldAdminStatus(factory, wrongOldOwner28, newOwner28)` with a MISTYPED oldOwner that
    /// never held `ADMIN_ROLE` must be `Inconsistent`, never `AlreadyDone` -- a stateless
    /// "does oldOwner currently lack ADMIN_ROLE" check cannot tell "already revoked" apart from
    /// "wrong address", so treating it as AlreadyDone would silently leave the REAL old admin's
    /// privileged ADMIN_ROLE untouched while reporting success. The execution wrapper must revert
    /// and must not broadcast.
    function test_revokeOldAdminStatus_wrongOldOwner_isInconsistent_notAlreadyDone() public {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        address realOldOwner28 = makeAddr("realOldOwner28");
        address newOwner28 = makeAddr("newOwner28");
        address wrongOldOwner28 = makeAddr("wrongOldOwner28"); // typo'd address, never held ADMIN_ROLE

        CrossMintableERC20V2Code fresh = _deployCode(realOldOwner28, bridge);

        vm.prank(realOldOwner28);
        fresh.beginDefaultAdminTransfer(newOwner28);
        vm.warp(block.timestamp + 1);
        vm.startPrank(newOwner28);
        fresh.acceptDefaultAdminTransfer();
        fresh.grantRole(Const.ADMIN_ROLE, newOwner28);
        vm.stopPrank();

        assertEq(fresh.defaultAdmin(), newOwner28);
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, realOldOwner28));
        assertFalse(fresh.hasRole(Const.ADMIN_ROLE, wrongOldOwner28));

        (CrossMintableERC20V2CodeScript.TransferStage stage, string memory reason) =
            scriptContract.revokeOldAdminStatus(address(fresh), wrongOldOwner28, newOwner28);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);

        vm.expectRevert(bytes(reason));
        scriptContract.revokeOldAdmin(address(fresh), wrongOldOwner28, newOwner28);

        // The real old admin's ADMIN_ROLE must remain completely untouched -- no silent success.
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, realOldOwner28));
    }

    /// Structural guarantee, fuzzed rather than pinned to one state: for ANY
    /// oldOwner/newOwner addresses, `revokeOldAdminStatus` never returns `AlreadyDone`. Its
    /// source (`revokeOldAdminStatus`'s @dev in the script) has no code path that returns
    /// `AlreadyDone` at all -- this test is the operational proof that structural claim holds
    /// for arbitrary inputs, not just the specific states the other tests in this group happen
    /// to construct.
    function testFuzz_revokeOldAdminStatus_neverReturnsAlreadyDoneForAnyOwners(address fuzzOldOwner, address fuzzNewOwner)
        public
    {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();

        // `code` is the shared setUp fixture (defaultAdmin()==factoryOwner, factoryOwner holds
        // ADMIN_ROLE) -- its exact state is irrelevant to what's being proven here: AlreadyDone
        // must never come out, for any fuzzed address pair, regardless of the target's state.
        (CrossMintableERC20V2CodeScript.TransferStage stage,) =
            scriptContract.revokeOldAdminStatus(address(code), fuzzOldOwner, fuzzNewOwner);
        assertTrue(stage != CrossMintableERC20V2CodeScript.TransferStage.AlreadyDone);
    }

    /// `newOwner30` has accepted `DEFAULT_ADMIN_ROLE` (step 2) but step 3 (granting itself
    /// `ADMIN_ROLE`) was skipped -- revoking `oldOwner30`'s `ADMIN_ROLE` now would strand the
    /// factory with zero admins. Must be `Inconsistent`, and the execution wrapper must revert
    /// without broadcasting (oldOwner30 keeps ADMIN_ROLE).
    function test_revokeOldAdminStatus_newOwnerMissingAdminRole_isInconsistent() public {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        address oldOwner30 = makeAddr("oldOwner30");
        address newOwner30 = makeAddr("newOwner30");

        CrossMintableERC20V2Code fresh = _deployCode(oldOwner30, bridge);

        vm.prank(oldOwner30);
        fresh.beginDefaultAdminTransfer(newOwner30);
        vm.warp(block.timestamp + 1);
        vm.prank(newOwner30);
        fresh.acceptDefaultAdminTransfer(); // step 3 (grantRole ADMIN_ROLE to newOwner) deliberately skipped

        assertEq(fresh.defaultAdmin(), newOwner30);
        assertFalse(fresh.hasRole(Const.ADMIN_ROLE, newOwner30));
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, oldOwner30));

        (CrossMintableERC20V2CodeScript.TransferStage stage, string memory reason) =
            scriptContract.revokeOldAdminStatus(address(fresh), oldOwner30, newOwner30);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);

        vm.expectRevert(bytes(reason));
        scriptContract.revokeOldAdmin(address(fresh), oldOwner30, newOwner30);

        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, oldOwner30));
    }

    /// `oldOwner30a` pre-grants `ADMIN_ROLE` to
    /// `newOwner30a` BEFORE the `DEFAULT_ADMIN_ROLE` transfer is even begun/accepted -- so both
    /// addresses already hold `ADMIN_ROLE` (preconditions 2 and 3 both pass), but
    /// `defaultAdmin()` is still `oldOwner30a` (precondition 1 fails). This is exactly the gap
    /// precondition 1 (`defaultAdmin() == newOwner`) exists to close: without it, this state
    /// would read as `Ready` and the execution wrapper would broadcast a `revokeRole` call that
    /// the factory's role-admin check would then reject. Must be `Inconsistent`, and the
    /// execution wrapper must not broadcast at all.
    function test_revokeOldAdminStatus_preGrantedAdminRoleButTransferNotAccepted_isInconsistent()
        public
    {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        address oldOwner30a = makeAddr("oldOwner30a");
        address newOwner30a = makeAddr("newOwner30a");

        CrossMintableERC20V2Code fresh = _deployCode(oldOwner30a, bridge);

        vm.prank(oldOwner30a);
        fresh.grantRole(Const.ADMIN_ROLE, newOwner30a);

        assertEq(fresh.defaultAdmin(), oldOwner30a);
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, oldOwner30a));
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, newOwner30a));

        (CrossMintableERC20V2CodeScript.TransferStage stage, string memory reason) =
            scriptContract.revokeOldAdminStatus(address(fresh), oldOwner30a, newOwner30a);
        assertEq(uint(stage), uint(CrossMintableERC20V2CodeScript.TransferStage.Inconsistent));
        assertTrue(bytes(reason).length > 0);

        vm.expectRevert(bytes(reason));
        scriptContract.revokeOldAdmin(address(fresh), oldOwner30a, newOwner30a);

        // No broadcast happened -- state is exactly as before the (reverted) call.
        assertEq(fresh.defaultAdmin(), oldOwner30a);
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, oldOwner30a));
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, newOwner30a));
    }

    /// Full 4-step ownership transfer, executed by calling the script wrappers directly
    /// (no `forge script` CLI invocation needed, same rationale as the deploy-script preflight
    /// test above) with two DISTINCT real
    /// addresses for oldOwner/newOwner. No `vm.prank` is used anywhere in this test -- each
    /// wrapper sets its own broadcast sender explicitly (`vm.startBroadcast(oldOwner)` /
    /// `vm.startBroadcast(newOwner)`), so there is no prank/broadcast interaction to worry about.
    function test_fullOwnershipTransfer_viaExplicitBroadcastSenders_completes() public {
        CrossMintableERC20V2CodeScript scriptContract = new CrossMintableERC20V2CodeScript();
        address oldOwner31 = makeAddr("oldOwner31");
        address newOwner31 = makeAddr("newOwner31");

        CrossMintableERC20V2Code fresh = _deployCode(oldOwner31, bridge);

        scriptContract.beginOwnershipTransfer(address(fresh), oldOwner31, newOwner31);
        vm.warp(block.timestamp + 1);
        scriptContract.acceptOwnershipAndGrantAdmin(address(fresh), newOwner31);
        scriptContract.revokeOldAdmin(address(fresh), oldOwner31, newOwner31);

        assertEq(fresh.defaultAdmin(), newOwner31);
        assertTrue(fresh.hasRole(Const.ADMIN_ROLE, newOwner31));
        assertFalse(fresh.hasRole(Const.ADMIN_ROLE, oldOwner31));
    }
}

/// @dev Minimal "upgraded" token implementation used only by the beacon-upgrade tests:
/// identical storage layout and behavior to `CrossMintableERC20V2`, plus one new function so the
/// tests can observe that an upgrade actually took effect.
contract CrossMintableERC20V2Mock is CrossMintableERC20V2 {
    function version() external pure returns (uint) {
        return 2;
    }
}

/// @dev Minimal "upgraded" factory implementation used only by the factory-upgrade tests.
contract CrossMintableERC20V2CodeMock is CrossMintableERC20V2Code {
    function version() external pure returns (uint) {
        return 2;
    }
}

/// @dev Minimal mock used only by the deploy-script preflight test: implements just `allTokenPairs`, matching
/// `IBridgeRegistry`'s selector, so it can be cast to `BaseBridge` and passed to
/// `CrossMintableERC20V2CodeScript.preflightCheckDuplicateRemoteToken` without deploying a real
/// bridge — Solidity dispatches external calls by selector, not by actual runtime type.
contract MockAllTokenPairsRegistry {
    IBridgeRegistry.TokenPair[] internal _pairs;

    function setPairs(IBridgeRegistry.TokenPair[] memory pairs_) external {
        delete _pairs;
        for (uint i = 0; i < pairs_.length; ++i) {
            _pairs.push(pairs_[i]);
        }
    }

    function allTokenPairs(uint) external view returns (IBridgeRegistry.TokenPair[] memory) {
        return _pairs;
    }
}
