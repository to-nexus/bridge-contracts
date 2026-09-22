// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {HyperMintableERC20} from "../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {ICrossMintableERC20V2Code} from "../src/token/ICrossMintableERC20V2Code.sol";
import {IHyperMintableERC20} from "../src/token/IHyperMintableERC20.sol";
import {IHyperMintableERC20Code} from "../src/token/IHyperMintableERC20Code.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {console} from "forge-std/Test.sol";

/**
 * @title HyperMintableERC20Test
 * @notice Covers HyperCore link hardening (raw-slot storage, one-shot linking, rounded Core
 *         transfers), the beacon-proxy factory shape shared with `CrossMintableERC20V2`, and the
 *         explicit name/symbol/minter creation path, using the proxied setup from `setUp`.
 * @dev `test/CrossBridgeHyperEVMRoute.t.sol` covers the end-to-end path through a real bridge
 *      instance instead of the bare `bridge` address used here.
 */
contract HyperMintableERC20Test is Test {
    /// @dev Declared locally (not imported from IERC20) purely so `vm.expectEmit` can match
    /// the inherited OZ ERC20 `Transfer` event by signature.
    event Transfer(address indexed from, address indexed to, uint value);

    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    /// @dev HyperEVM's Core read precompile, mirrored here (not imported) so a typo in the
    /// contract's own copy would show up as a test failure rather than being masked.
    address internal constant CORE_TOKEN_INFO_PRECOMPILE = 0x000000000000000000000000000000000000080C;

    // ERC-7201 / ERC-1967 slot literals — the test's OWN copies (never read from a getter) so
    // these tests actually verify independent, forge-computed values (`cast index-erc7201 <id>`),
    // not merely that the contract agrees with itself.
    bytes32 internal constant TOKEN_STORAGE_SLOT = 0xc65afee183f44952959f664e5c2b8a5e4916480c324e2787b2215d744391d400;
    /// @dev `HyperMintableERC20Code` has no namespaced storage of its own — it fully inherits
    /// `CrossMintableERC20V2Code`'s, so this is THAT slot (matches
    /// `CrossMintableERC20V2CodeStorageLocation` in `src/token/CrossMintableERC20V2Code.sol`),
    /// not a Hyper-specific one.
    bytes32 internal constant FACTORY_STORAGE_SLOT = 0x2b6c7dfd4e9330f706f3d4045fa01a8903ed43939ff7d252d6806f39e1487c00;
    bytes32 internal constant ERC1967_IMPLEMENTATION_SLOT =
        0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
    bytes32 internal constant ERC1967_BEACON_SLOT = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;

    HyperMintableERC20 public tokenImplementation;
    UpgradeableBeacon public beacon;
    HyperMintableERC20Code public codeImplementation;
    HyperMintableERC20Code public code;
    HyperMintableERC20 public token;

    uint internal constant FACTORY_OWNER_PK = uint(bytes32("factoryOwner"));
    uint internal constant TOKEN_ADMIN_PK = uint(bytes32("tokenAdmin"));
    uint internal constant BRIDGE_PK = uint(bytes32("bridge"));
    uint internal constant USER_PK = uint(bytes32("user"));
    uint internal constant FINALIZER_PK = uint(bytes32("finalizer"));

    address public factoryOwner; // factory ADMIN_ROLE + DEFAULT_ADMIN_ROLE holder
    /// @dev NOT the factory-created fixture `token`'s admin — that's `address(code)`. Kept as a
    /// plain funded EOA for the "link authority" test group below, which deploys its own
    /// STANDALONE token (via `_deployStandaloneToken`) with this as a deliberately DISTINCT admin
    /// from the factory — see those tests' doc comments.
    address public tokenAdmin;
    address public bridge; // BRIDGE_ROLE on the factory, MINTER_ROLE on created tokens
    address public user;
    address public finalizer;

    uint internal constant REMOTE_CHAIN_ID = 998; // HyperEVM testnet
    address internal constant REMOTE_TOKEN = address(0xBEEF);
    string internal constant SYMBOL = "TT";
    uint8 internal constant DECIMALS = 18;

    uint internal constant INITIAL_SUPPLY = 1_000_000_000 ether;

    function setUp() public {
        factoryOwner = vm.addr(FACTORY_OWNER_PK);
        tokenAdmin = vm.addr(TOKEN_ADMIN_PK);
        bridge = vm.addr(BRIDGE_PK);
        user = vm.addr(USER_PK);
        finalizer = vm.addr(FINALIZER_PK);

        tokenImplementation = new HyperMintableERC20();
        codeImplementation = new HyperMintableERC20Code();

        // `HyperMintableERC20Code` inherits `CrossMintableERC20V2Code`: `initialize` is the
        // inherited 3-arg form (no `tokenAdmin` param), and it creates + self-owns its own token
        // beacon internally — there is no separate beacon deploy/ownership-handoff step.
        // `initialize` is referenced via the DECLARING contract
        // (`CrossMintableERC20V2Code`) because `abi.encodeCall`'s magic member lookup only
        // resolves functions declared directly on the named type, not merely-inherited ones; the
        // selector only depends on the signature, so this still dispatches correctly against
        // `codeImplementation`'s actual (inherited) `initialize`.
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(CrossMintableERC20V2Code.initialize, (factoryOwner, bridge, address(tokenImplementation)))
        );
        code = HyperMintableERC20Code(address(proxy));
        beacon = UpgradeableBeacon(code.beacon()); // self-created + self-owned by the factory
        console.log("HyperMintableERC20Code", address(code));

        vm.prank(bridge);
        token = HyperMintableERC20(code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS));
        console.log("HyperMintableERC20", address(token));
    }

    function _deployCode(address factoryOwnerAddr, address bridgeAddr) internal returns (HyperMintableERC20Code freshCode) {
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(
                CrossMintableERC20V2Code.initialize, (factoryOwnerAddr, bridgeAddr, address(tokenImplementation))
            )
        );
        freshCode = HyperMintableERC20Code(address(proxy));
    }

    /// @dev Deploys a STANDALONE `HyperMintableERC20` (bypassing the factory) with deliberately
    /// DISTINCT `initialOwner`/`initialLinker` — needed by the "link authority" test group below
    /// to exercise the two-DISTINCT-principals design that a factory-created token (whose
    /// `defaultAdmin()` and `factoryLinker` are the SAME address, the factory itself) cannot
    /// exhibit. NOT registered with any factory (`isCrossMintableERC20` is false for it), so it
    /// is only reachable directly, never via a factory's delegated setters.
    function _deployStandaloneToken(address admin, address linker) internal returns (HyperMintableERC20 t) {
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(tokenImplementation),
            abi.encodeCall(
                IHyperMintableERC20.initialize,
                (admin, bridge, linker, "Standalone", "STA", DECIMALS)
            )
        );
        t = HyperMintableERC20(address(proxy));
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

    /// @dev Stubs the Core read precompile's `tokenInfo(uint32)` for `index`. Fixture values
    /// default to a realistic TONE-style link: `weiDecimals=8`, `evmExtraWeiDecimals=10`,
    /// summing to `DECIMALS` (18) so `setCoreTokenIndex`'s decimals cross-check passes unless
    /// the test deliberately breaks it.
    function _mockCoreTokenInfo(uint32 index, address evmContract, uint8 weiDecimals, int8 evmExtraWeiDecimals)
        internal
    {
        IHyperMintableERC20.CoreTokenInfo memory info = IHyperMintableERC20.CoreTokenInfo({
            name: "MOCK",
            spots: new uint64[](0),
            deployerTradingFeeShare: 0,
            deployer: address(0),
            evmContract: evmContract,
            szDecimals: 2,
            weiDecimals: weiDecimals,
            evmExtraWeiDecimals: evmExtraWeiDecimals
        });
        vm.mockCall(CORE_TOKEN_INFO_PRECOMPILE, abi.encode(index), abi.encode(info));
    }

    function _mockValidCoreLink(address tokenAddr, uint32 index) internal {
        _mockCoreTokenInfo(index, tokenAddr, 8, 10);
    }

    // ---------------------------------------------------------------------
    // Slot correctness (core)
    // ---------------------------------------------------------------------

    /// @dev `Const.LINKER_ROLE` matches its documented digest.
    function test_constLinkerRole() public pure {
        assertEq(Const.LINKER_ROLE, keccak256("LINKER_ROLE"));
    }

    /// The slot constant is a raw `keccak256("HyperCore deployer")`, never an ERC-7201
    /// masked digest — a typo here silently breaks the only usable `finalizeEvmContract` path.
    function test_hypercoreDeployerSlot_isRawKeccakOfLiteralString() public view {
        assertEq(token.HYPERCORE_DEPLOYER_SLOT(), keccak256("HyperCore deployer"));
        assertEq(token.HYPERCORE_DEPLOYER_SLOT(), 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f);
    }

    /// `setHyperCoreDeployer` writes the RAW slot, not just something the getter agrees with.
    function test_setHyperCoreDeployer_writesRawSlot() public {
        vm.prank(address(code));
        token.setHyperCoreDeployer(finalizer);

        bytes32 raw = vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT());
        assertEq(raw, bytes32(uint(uint160(finalizer))));
        assertEq(token.hyperCoreDeployer(), finalizer);
    }

    /// No aliasing in either direction between the link slot and ordinary ERC20/AccessControl storage.
    function test_hyperCoreDeployerSlot_doesNotAliasOrdinaryStorage() public {
        vm.prank(bridge);
        token.mint(user, INITIAL_SUPPLY);

        vm.prank(user);
        token.approve(finalizer, 123 ether);

        vm.prank(user);
        token.transfer(tokenAdmin, 1 ether);

        uint deadline = block.timestamp + 1 days;
        (uint8 v, bytes32 r, bytes32 s) = _signPermit(USER_PK, user, finalizer, 5 ether, deadline);
        token.permit(user, finalizer, 5 ether, deadline, v, r, s);

        vm.prank(address(code));
        token.grantRole(Const.MINTER_ROLE, user);

        // Direction 1: heavy ordinary activity must not disturb the (still-unset) link slot.
        bytes32 rawAfterActivity = vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT());
        assertEq(rawAfterActivity, bytes32(0));

        // Direction 2: writing the link slot must not disturb balances/allowance/nonces/roles.
        uint balBefore = token.balanceOf(user);
        uint allowBefore = token.allowance(user, finalizer);
        uint noncesBefore = token.nonces(user);
        bool hasRoleBefore = token.hasRole(Const.MINTER_ROLE, user);

        vm.prank(address(code));
        token.setHyperCoreDeployer(finalizer);

        assertEq(token.balanceOf(user), balBefore);
        assertEq(token.allowance(user, finalizer), allowBefore);
        assertEq(token.nonces(user), noncesBefore);
        assertEq(token.hasRole(Const.MINTER_ROLE, user), hasRoleBefore);
    }

    /// Finalizer can be recalled back to zero.
    function test_setHyperCoreDeployer_canResetToZero() public {
        vm.startPrank(address(code));
        token.setHyperCoreDeployer(finalizer);
        token.setHyperCoreDeployer(address(0));
        vm.stopPrank();

        assertEq(token.hyperCoreDeployer(), address(0));
        assertEq(vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT()), bytes32(0));
    }

    // ---------------------------------------------------------------------
    // Permissions
    // ---------------------------------------------------------------------

    /// Factory (ADMIN_ROLE) path succeeds, token defaultAdmin direct path succeeds, third party reverts.
    function test_setHyperCoreDeployer_viaFactoryOrDirectly_thirdPartyReverts() public {
        vm.prank(factoryOwner);
        code.setHyperCoreDeployer(address(token), finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        address finalizer2 = makeAddr("finalizer2");
        vm.prank(address(code));
        token.setHyperCoreDeployer(finalizer2);
        assertEq(token.hyperCoreDeployer(), finalizer2);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);
    }

    /// Token admin revoking the factory's LINKER_ROLE blocks the factory path (fallback still works).
    /// @dev Standalone token (see the note on `test_isLinkAuthority_oldAdminLosesAuthorityAfterDefaultAdminTransfer`
    /// below): on the factory-created fixture `token`,
    /// `address(code)` is ALSO `defaultAdmin()`, so revoking its `LINKER_ROLE` has no observable
    /// effect on `code.setHyperCoreDeployer` — that call reaches the token as `address(code)`,
    /// and the `defaultAdmin()` branch of `isLinkAuthority` alone already grants authority
    /// regardless of `LINKER_ROLE`. This scenario (revoking the role actually blocking the
    /// factory path) is only meaningful when the factory-linker identity is distinct from the
    /// admin, which is exactly what this standalone token sets up. `code.setHyperCoreDeployer`
    /// additionally requires the target to be a token `code` itself created (`isCrossMintableERC20`)
    /// — a standalone token deliberately is not, so it is called directly here (as `address(code)`
    /// would itself) instead.
    function test_revokeFactoryLinkerRole_blocksFactoryPath() public {
        HyperMintableERC20 standaloneToken = _deployStandaloneToken(tokenAdmin, address(code));

        vm.prank(tokenAdmin);
        standaloneToken.revokeRole(Const.LINKER_ROLE, address(code));

        vm.prank(address(code));
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(code), Const.LINKER_ROLE
            )
        );
        standaloneToken.setHyperCoreDeployer(finalizer);

        vm.prank(tokenAdmin);
        standaloneToken.setHyperCoreDeployer(finalizer);
        assertEq(standaloneToken.hyperCoreDeployer(), finalizer);
    }

    /// Factory setters require ADMIN_ROLE.
    function test_factorySetters_requireAdminRole() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.setHyperCoreDeployer(address(token), finalizer);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.setCoreTokenIndex(address(token), 5);
    }

    /// Deployer != initialOwner must not revert the factory's atomic proxy
    /// initialization (the trap the V2 factory falls into by using public `grantRole` instead
    /// of `_grantRole`).
    function test_initialize_succeedsWhenDeployerIsNotInitialOwner() public {
        address deployer = makeAddr("someoneElse");
        vm.prank(deployer);
        HyperMintableERC20Code freshCode = _deployCode(factoryOwner, bridge);

        assertTrue(freshCode.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertFalse(freshCode.hasRole(Const.ADMIN_ROLE, deployer));
        assertTrue(freshCode.beacon() != address(0), "factory must have created its own beacon");
    }

    // ---------------------------------------------------------------------
    // Core system address
    // ---------------------------------------------------------------------

    /// index 200 -> 0x20000000000000000000000000000000000000c8 exactly.
    function test_coreSystemAddress_index200MapsToExpectedPrecompileAddress() public {
        _mockValidCoreLink(address(token), 200);
        vm.prank(address(code));
        token.setCoreTokenIndex(200);

        address expected = 0x20000000000000000000000000000000000000C8;
        assertEq(token.coreSystemAddress(), expected);
    }

    /// index 0 is settable and valid (not an "unset" sentinel).
    function test_setCoreTokenIndex_zeroIsValidNotUnsetSentinel() public {
        _mockValidCoreLink(address(token), 0);
        vm.prank(address(code));
        token.setCoreTokenIndex(0);

        assertTrue(token.isCoreTokenIndexSet());
        assertEq(token.coreTokenIndex(), 0);
        assertEq(token.coreSystemAddress(), address(uint160(0x20) << 152));
    }

    /// Unset index makes both `coreSystemAddress` and `transferToCore` revert.
    function test_coreSystemAddressAndTransferToCore_revertWhenIndexUnset() public {
        vm.expectRevert(HyperMintableERC20.HyperMintableERC20CoreTokenIndexNotSet.selector);
        token.coreSystemAddress();

        vm.prank(user);
        vm.expectRevert(HyperMintableERC20.HyperMintableERC20CoreTokenIndexNotSet.selector);
        token.transferToCore(1);
    }

    /// `transferToCore` moves balance to the system address and emits `Transfer`; insufficient balance reverts.
    function test_transferToCore_movesBalanceAndRevertsOnInsufficientBalance() public {
        _mockValidCoreLink(address(token), 200);
        vm.prank(address(code));
        token.setCoreTokenIndex(200);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(user, systemAddr, 5 ether);

        vm.prank(user);
        assertEq(token.transferToCore(5 ether), 5 ether);

        assertEq(token.balanceOf(systemAddr), 5 ether);
        assertEq(token.balanceOf(user), 5 ether);

        uint balance = token.balanceOf(user);
        // +1 alone would round back down to exactly `balance` (coreUnit() == 1e10 here) and
        // succeed instead of reverting — go a full coreUnit over so the rounded-down `sent`
        // still exceeds the balance.
        uint attempted = balance + token.coreUnit();
        uint needed = token.coreTransferableAmount(attempted);
        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, user, balance, needed));
        vm.prank(user);
        token.transferToCore(attempted);
    }

    /// Fuzz: for any in-range index, the high byte is always 0x20 and the low 19 bytes are the index.
    function testFuzz_coreSystemAddress_highByteFixedLowBytesEqualIndex(uint32 index) public {
        _mockValidCoreLink(address(token), index);
        vm.prank(address(code));
        token.setCoreTokenIndex(index);

        uint160 sys = uint160(token.coreSystemAddress());
        assertEq(sys >> 152, uint160(0x20));
        assertEq(sys & ((uint160(1) << 152) - 1), uint160(index));
    }

    // ---------------------------------------------------------------------
    // ERC20 / factory
    // ---------------------------------------------------------------------

    /// mint/burn work normally; missing MINTER_ROLE reverts.
    function test_mintAndBurn_requireMinterRole() public {
        vm.prank(bridge);
        assertTrue(token.mint(user, INITIAL_SUPPLY));
        assertEq(token.balanceOf(user), INITIAL_SUPPLY);

        vm.prank(bridge);
        assertTrue(token.burn(user, INITIAL_SUPPLY));
        assertEq(token.balanceOf(user), 0);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.MINTER_ROLE)
        );
        token.mint(user, 1);
    }

    /// EIP-2612 permit signature verification, domain name = "Cross Bridge <SYM>".
    function test_permit_verifiesEip2612Signature() public {
        assertEq(token.name(), string(abi.encodePacked("Cross Bridge ", SYMBOL)));
        assertEq(token.symbol(), string(abi.encodePacked(SYMBOL, "x")));

        uint value = 100 ether;
        uint deadline = block.timestamp + 1 days;
        uint nonceBefore = token.nonces(user);
        (uint8 v, bytes32 r, bytes32 s) = _signPermit(USER_PK, user, finalizer, value, deadline);

        token.permit(user, finalizer, value, deadline, v, r, s);

        assertEq(token.allowance(user, finalizer), value);
        assertEq(token.nonces(user), nonceBefore + 1);
    }

    /// `computeTokenAddress` prediction equals the actual `createCrossMintableERC20` result.
    function test_computeTokenAddress_matchesActualCreatedAddress() public view {
        address predicted = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predicted, address(token));
    }

    /// Same (remoteChainID, remoteToken) but different symbol/decimals predicts a different address.
    function test_computeTokenAddress_changesWithSymbolOrDecimals() public view {
        address predicted = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        address predictedOtherSymbol =
            code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, "OTHER", DECIMALS, bridge);
        address predictedOtherDecimals = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, 6, bridge);

        assertTrue(predicted != predictedOtherSymbol);
        assertTrue(predicted != predictedOtherDecimals);
    }

    /// `isHyperMintableERC20` is true right after creation; setter on an unknown address reverts.
    function test_isHyperMintableERC20_trueForCreated_settersRevertForUnknown() public {
        assertTrue(code.isHyperMintableERC20(address(token)));
        assertFalse(code.isHyperMintableERC20(address(0xDEAD)));

        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodeUnknownToken.selector, address(0xDEAD)
            )
        );
        code.setHyperCoreDeployer(address(0xDEAD), finalizer);
    }

    /// Factory replacement: factory B's ADMIN cannot touch a token factory A created;
    /// factory A's ADMIN can; the token's own defaultAdmin (the creating factory itself)
    /// always can.
    function test_setHyperCoreDeployer_viaFactory_scopedToCreatingFactory() public {
        address factoryBOwner = makeAddr("factoryBOwner");
        HyperMintableERC20Code codeB = _deployCode(factoryBOwner, bridge);

        vm.prank(factoryBOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                ICrossMintableERC20V2Code.CrossMintableERC20V2CodeUnknownToken.selector, address(token)
            )
        );
        codeB.setHyperCoreDeployer(address(token), finalizer);

        vm.prank(factoryOwner);
        code.setHyperCoreDeployer(address(token), finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        address finalizer2 = makeAddr("finalizer2");
        vm.prank(address(code));
        token.setHyperCoreDeployer(finalizer2);
        assertEq(token.hyperCoreDeployer(), finalizer2);
    }

    /// Determinism holds even when the address that computes the prediction differs from
    /// the address that actually calls `createCrossMintableERC20` — as long as `minter` matches
    /// the real creator (the bridge).
    function test_computeTokenAddress_deterministicRegardlessOfCaller() public {
        uint remoteChainID2 = REMOTE_CHAIN_ID + 1;

        address predictedByOwner = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(user);
        address predictedByOther = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedByOwner, predictedByOther);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(actual, predictedByOwner);
    }

    /// Transferring the factory's OWN default admin (2-step) does not change what
    /// `defaultAdmin()` newly created tokens end up with, nor the predicted address — both are
    /// pinned to the factory's OWN (CREATE2-fixed) address, `address(code)`, not to whoever
    /// currently holds the factory's `ADMIN_ROLE`/`defaultAdmin()`.
    function test_factoryDefaultAdminTransfer_doesNotAffectTokenAdminOrPrediction() public {
        address newFactoryOwner = makeAddr("newFactoryOwner");
        uint remoteChainID3 = REMOTE_CHAIN_ID + 2;

        address predictedBefore = code.computeTokenAddress(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(factoryOwner);
        code.beginDefaultAdminTransfer(newFactoryOwner);
        vm.warp(block.timestamp + 1);
        vm.prank(newFactoryOwner);
        code.acceptDefaultAdminTransfer();

        assertEq(code.defaultAdmin(), newFactoryOwner);

        address predictedAfter = code.computeTokenAddress(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedAfter, predictedBefore);

        vm.prank(bridge);
        address newToken = code.createCrossMintableERC20(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(newToken, predictedBefore);
        assertEq(HyperMintableERC20(newToken).defaultAdmin(), address(code), "defaultAdmin() == the factory itself");
        assertTrue(HyperMintableERC20(newToken).isLinkAuthority(address(code)));
    }

    /// Negative: predicting with the wrong `minter` yields an address that differs from
    /// what the real bridge-driven creation actually deploys.
    function test_computeTokenAddress_wrongMinterYieldsDifferentAddress() public {
        uint remoteChainID4 = REMOTE_CHAIN_ID + 3;
        address wrongMinter = makeAddr("wrongMinter");

        address predictedWrong = code.computeTokenAddress(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS, wrongMinter);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertTrue(predictedWrong != actual);
    }

    /// Redeploying with the same salt + same args reverts.
    /// @dev Reverts with the inherited `CrossMintableERC20V2CodePairAlreadyCreated` guard:
    /// `_create` (fully inherited from `CrossMintableERC20V2Code`, not reimplemented) checks
    /// `tokenForPair` BEFORE ever reaching the CREATE2 deploy, so the redeploy never gets far
    /// enough to hit the raw `Errors.FailedDeployment` CREATE2-collision revert.
    function test_createCrossMintableERC20_duplicatePairReverts() public {
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

    /// EIP-170: runtime code of both LOGIC implementations (what the 24,576-byte limit
    /// actually constrains, since both token and factory are proxies) stays under the limit;
    /// the proxies themselves are trivially small.
    function test_runtimeCodeSize_underEip170Limit() public view {
        assertLt(address(tokenImplementation).code.length, 24_576);
        assertLt(address(codeImplementation).code.length, 24_576);
        assertLt(address(token).code.length, 24_576);
        assertLt(address(code).code.length, 24_576);
    }

    // ---------------------------------------------------------------------
    // Link authority (H1: exactly two principals — current defaultAdmin() and factoryLinker())
    // ---------------------------------------------------------------------

    /// After a completed default-admin transfer, the *old* admin's link authority is gone:
    /// both link setters revert with `AccessControlUnauthorizedAccount(oldAdmin, LINKER_ROLE)`.
    /// @dev Uses a STANDALONE token (`_deployStandaloneToken`), not the factory-created fixture
    /// `token`: a factory-created token's `defaultAdmin()` and `factoryLinker` are the
    /// SAME address (the factory itself), which retains link authority via the factoryLinker
    /// path even after a defaultAdmin transfer — so "old admin" and "factory" would not be
    /// distinguishable principals on `token`, and this old-admin-specifically-loses-authority
    /// assertion would not hold. A standalone token lets `tokenAdmin` (the admin) and
    /// `address(code)` (the linker) start out as two genuinely distinct principals instead.
    function test_isLinkAuthority_oldAdminLosesAuthorityAfterDefaultAdminTransfer() public {
        HyperMintableERC20 standaloneToken = _deployStandaloneToken(tokenAdmin, address(code));
        address newAdmin = makeAddr("newAdmin");

        vm.prank(tokenAdmin);
        standaloneToken.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        standaloneToken.acceptDefaultAdminTransfer();

        assertEq(standaloneToken.defaultAdmin(), newAdmin);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, tokenAdmin, Const.LINKER_ROLE
            )
        );
        standaloneToken.setHyperCoreDeployer(finalizer);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, tokenAdmin, Const.LINKER_ROLE
            )
        );
        standaloneToken.setCoreTokenIndex(1);
    }

    /// The *new* admin gains link authority immediately, with no separate grant needed, and
    /// can actually write the raw slot / set the Core index.
    /// @dev Standalone token — see the note on `test_isLinkAuthority_oldAdminLosesAuthorityAfterDefaultAdminTransfer`.
    function test_isLinkAuthority_newAdminGainsAuthorityImmediatelyAfterTransfer() public {
        HyperMintableERC20 standaloneToken = _deployStandaloneToken(tokenAdmin, address(code));
        address newAdmin = makeAddr("newAdmin");

        vm.prank(tokenAdmin);
        standaloneToken.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        standaloneToken.acceptDefaultAdminTransfer();

        vm.prank(newAdmin);
        standaloneToken.setHyperCoreDeployer(finalizer);
        assertEq(standaloneToken.hyperCoreDeployer(), finalizer);
        bytes32 raw = vm.load(address(standaloneToken), standaloneToken.HYPERCORE_DEPLOYER_SLOT());
        assertEq(raw, bytes32(uint(uint160(finalizer))));

        _mockValidCoreLink(address(standaloneToken), 7);
        vm.prank(newAdmin);
        standaloneToken.setCoreTokenIndex(7);
        assertEq(standaloneToken.coreTokenIndex(), 7);
    }

    /// A third party (never granted any role) reverts both before AND after the admin
    /// transfer.
    function test_setHyperCoreDeployer_thirdPartyRevertsBeforeAndAfterAdminTransfer() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);

        address newAdmin = makeAddr("newAdmin");
        vm.prank(address(code));
        token.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        token.acceptDefaultAdminTransfer();

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);
    }

    /// Granting `LINKER_ROLE` to a third party flips `hasRole` but not `isLinkAuthority` —
    /// link authority is exactly two principals, enforced in code, not just documented.
    function test_isLinkAuthority_falseForThirdPartyGrantedOnlyLinkerRole() public {
        vm.prank(address(code));
        token.grantRole(Const.LINKER_ROLE, user);

        assertTrue(token.hasRole(Const.LINKER_ROLE, user));
        assertFalse(token.isLinkAuthority(user));

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);
    }

    /// `isLinkAuthority` matches actual success/failure for five cases: old admin, new
    /// admin, factory, a third party holding only the role, and the zero address.
    /// @dev Standalone token — see the note on `test_isLinkAuthority_oldAdminLosesAuthorityAfterDefaultAdminTransfer`.
    function test_isLinkAuthority_matchesActualOutcome() public {
        HyperMintableERC20 standaloneToken = _deployStandaloneToken(tokenAdmin, address(code));
        address newAdmin = makeAddr("newAdmin");
        vm.prank(tokenAdmin);
        standaloneToken.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        standaloneToken.acceptDefaultAdminTransfer();

        vm.prank(newAdmin);
        standaloneToken.grantRole(Const.LINKER_ROLE, user);

        assertFalse(standaloneToken.isLinkAuthority(tokenAdmin)); // old admin
        assertTrue(standaloneToken.isLinkAuthority(newAdmin)); // new admin
        assertTrue(standaloneToken.isLinkAuthority(address(code))); // factory (factoryLinker)
        assertFalse(standaloneToken.isLinkAuthority(user)); // role-only third party
        assertFalse(standaloneToken.isLinkAuthority(address(0))); // zero address
    }

    /// Revoking the factory's `LINKER_ROLE` blocks the factory-linker authority path even
    /// though `factoryLinker` cannot itself be reassigned; re-granting it restores the path.
    /// @dev Standalone token — see the note on `test_isLinkAuthority_oldAdminLosesAuthorityAfterDefaultAdminTransfer`;
    /// additionally, this exercises the link functions
    /// DIRECTLY (as `address(code)` would call them itself) rather than through the factory's
    /// `setHyperCoreDeployer` delegation, since that path additionally requires the target to be
    /// a token the CALLING factory actually created (`isCrossMintableERC20`) — which a
    /// standalone token deliberately is not. On the factory-created fixture `token`, revoking
    /// `address(code)`'s `LINKER_ROLE` would have no observable effect anyway, since `address(code)`
    /// is ALSO `token`'s `defaultAdmin()` there and that branch alone already grants authority —
    /// this scenario is only meaningful when the factory-linker identity is distinct from the
    /// admin, exactly what this standalone token sets up.
    function test_isLinkAuthority_factoryPathRestoredAfterRegrantingLinkerRole() public {
        HyperMintableERC20 standaloneToken = _deployStandaloneToken(tokenAdmin, address(code));
        assertEq(standaloneToken.factoryLinker(), address(code));

        vm.prank(tokenAdmin);
        standaloneToken.revokeRole(Const.LINKER_ROLE, address(code));

        vm.prank(address(code));
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(code), Const.LINKER_ROLE
            )
        );
        standaloneToken.setHyperCoreDeployer(finalizer);

        vm.prank(tokenAdmin);
        standaloneToken.grantRole(Const.LINKER_ROLE, address(code));

        vm.prank(address(code));
        standaloneToken.setHyperCoreDeployer(finalizer);
        assertEq(standaloneToken.hyperCoreDeployer(), finalizer);
    }

    // =====================================================================
    // ERC-7201 / ERC-1967 slot hygiene
    // =====================================================================

    /// The token's and factory's namespaced storage slots never collide with each other,
    /// with the raw HyperCore deployer slot, or with the ERC-1967 beacon/implementation slots
    /// the proxies themselves rely on. Also checks the ERC-7201 masking convention (low byte
    /// zeroed) actually holds for both computed constants.
    function test_namespacedSlots_noCollisionAndErc7201MaskingHolds() public view {
        assertTrue(TOKEN_STORAGE_SLOT != FACTORY_STORAGE_SLOT);
        assertTrue(TOKEN_STORAGE_SLOT != token.HYPERCORE_DEPLOYER_SLOT());
        assertTrue(FACTORY_STORAGE_SLOT != token.HYPERCORE_DEPLOYER_SLOT());
        assertTrue(TOKEN_STORAGE_SLOT != ERC1967_IMPLEMENTATION_SLOT);
        assertTrue(TOKEN_STORAGE_SLOT != ERC1967_BEACON_SLOT);
        assertTrue(FACTORY_STORAGE_SLOT != ERC1967_IMPLEMENTATION_SLOT);
        assertTrue(FACTORY_STORAGE_SLOT != ERC1967_BEACON_SLOT);

        assertEq(uint(TOKEN_STORAGE_SLOT) & 0xff, 0);
        assertEq(uint(FACTORY_STORAGE_SLOT) & 0xff, 0);
    }

    /// Token state lives at the raw `TOKEN_STORAGE_SLOT`, packed into a single word:
    /// `coreTokenIndex` (uint64, bytes 0-7) | `coreTokenIndexSet` (bool, byte 8) |
    /// `coreExtraWeiDecimals` (int8, byte 9) | `factoryLinker` (address, bytes 10-29). `decimals`
    /// is not part of this struct/slot — it lives in `CrossMintableERC20V2`'s own ERC-7201
    /// storage instead (asserted via `token.decimals()` directly, not decoded from this slot).
    function test_tokenStorage_packsFieldsIntoSingleRawSlot() public {
        _mockValidCoreLink(address(token), 5);
        vm.prank(address(code));
        token.setCoreTokenIndex(5);

        bytes32 raw = vm.load(address(token), TOKEN_STORAGE_SLOT);
        uint word = uint(raw);

        // Solidity packs the FIRST-declared struct member into the LOW-order end of the slot:
        // coreTokenIndex (bits 0-63) | coreTokenIndexSet (64-71) | coreExtraWeiDecimals (72-79)
        // | factoryLinker (80-239) — i.e. bytes 10-29.
        uint64 decodedIndex = uint64(word);
        bool decodedSet = uint8(word >> 64) != 0;
        int8 decodedExtra = int8(uint8(word >> 72));
        address decodedFactoryLinker = address(uint160(word >> 80));

        assertEq(decodedFactoryLinker, token.factoryLinker());
        assertEq(decodedExtra, token.coreExtraWeiDecimals());
        assertEq(decodedSet, token.isCoreTokenIndexSet());
        assertEq(decodedIndex, token.coreTokenIndex());

        // `decimals()` itself lives in Cross's storage now, not this slot — assert it separately
        // via the inherited getter, and confirm it's plainly reachable (not derived from garbage).
        assertEq(token.decimals(), DECIMALS);
    }

    // =====================================================================
    // setCoreTokenIndex validation + one-shot lock
    // =====================================================================

    /// A correctly linked index succeeds and records the exact index + evmExtraWeiDecimals.
    function test_setCoreTokenIndex_validLink_succeeds() public {
        _mockValidCoreLink(address(token), 2895);

        vm.expectEmit(true, false, false, true, address(token));
        emit IHyperMintableERC20.CoreTokenIndexSet(2895, 10);

        vm.prank(address(code));
        token.setCoreTokenIndex(2895);

        assertTrue(token.isCoreTokenIndexSet());
        assertEq(token.coreTokenIndex(), 2895);
        assertEq(token.coreExtraWeiDecimals(), 10);
    }

    /// Unlinked index (`evmContract == address(0)`) is rejected.
    function test_setCoreTokenIndex_unlinked_reverts() public {
        _mockCoreTokenInfo(3, address(0), 8, 10);

        vm.prank(address(code));
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreLinkNotFinalized.selector, uint64(3), address(0))
        );
        token.setCoreTokenIndex(3);
    }

    /// Index linked to a DIFFERENT contract is rejected — this is exactly the spot-pair
    /// vs. token-index mix-up that a real Core link can hit (index 2902 resolves to a different
    /// token).
    function test_setCoreTokenIndex_linkedToOtherContract_reverts() public {
        address otherToken = makeAddr("otherHyperCoreToken");
        _mockCoreTokenInfo(2902, otherToken, 8, 10);

        vm.prank(address(code));
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreLinkNotFinalized.selector, uint64(2902), otherToken)
        );
        token.setCoreTokenIndex(2902);
    }

    /// Precompile revert (e.g. wildly out-of-range index) is surfaced as `CoreTokenInfoUnavailable`.
    function test_setCoreTokenIndex_precompileReverts_reverts() public {
        vm.mockCallRevert(CORE_TOKEN_INFO_PRECOMPILE, abi.encode(uint32(999_999)), bytes("precompile revert"));

        vm.prank(address(code));
        vm.expectRevert(abi.encodeWithSelector(IHyperMintableERC20.CoreTokenInfoUnavailable.selector, uint64(999_999)));
        token.setCoreTokenIndex(999_999);
    }

    /// `index > type(uint32).max` is rejected before ever reaching the precompile.
    function test_setCoreTokenIndex_indexOutOfUint32Range_reverts() public {
        uint64 tooLarge = uint64(type(uint32).max) + 1;

        vm.prank(address(code));
        vm.expectRevert(abi.encodeWithSelector(IHyperMintableERC20.CoreTokenIndexOutOfRange.selector, tooLarge));
        token.setCoreTokenIndex(tooLarge);
    }

    /// `decimals() != weiDecimals + evmExtraWeiDecimals` is rejected.
    function test_setCoreTokenIndex_decimalsMismatch_reverts() public {
        // 5 + 5 = 10 != DECIMALS (18).
        _mockCoreTokenInfo(9, address(token), 5, 5);

        vm.prank(address(code));
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreDecimalsMismatch.selector, DECIMALS, uint8(5), int8(5))
        );
        token.setCoreTokenIndex(9);
    }

    /// A second call — even with the exact same (already-linked) index — is rejected once
    /// the first has succeeded.
    function test_setCoreTokenIndex_revertsOnSecondCall() public {
        _mockValidCoreLink(address(token), 11);
        vm.startPrank(address(code));
        token.setCoreTokenIndex(11);

        vm.expectRevert(IHyperMintableERC20.CoreTokenIndexAlreadySet.selector);
        token.setCoreTokenIndex(11);

        _mockValidCoreLink(address(token), 12);
        vm.expectRevert(IHyperMintableERC20.CoreTokenIndexAlreadySet.selector);
        token.setCoreTokenIndex(12);
        vm.stopPrank();
    }

    // =====================================================================
    // setHyperCoreDeployer lock
    // =====================================================================

    /// Once `setCoreTokenIndex` has succeeded, `setHyperCoreDeployer` is permanently locked.
    function test_setHyperCoreDeployer_lockedAfterIndexSet() public {
        _mockValidCoreLink(address(token), 13);
        vm.startPrank(address(code));
        token.setCoreTokenIndex(13);

        vm.expectRevert(IHyperMintableERC20.HyperCoreLinkAlreadyFinalized.selector);
        token.setHyperCoreDeployer(finalizer);
        vm.stopPrank();
    }

    /// Before the index is set, `setHyperCoreDeployer` behaves exactly as before —
    /// including recalling a mistaken finalizer back to the zero address.
    function test_setHyperCoreDeployer_worksBeforeIndexSet_includingRecallToZero() public {
        assertFalse(token.isCoreTokenIndexSet());

        vm.startPrank(address(code));
        token.setHyperCoreDeployer(finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        token.setHyperCoreDeployer(address(0));
        assertEq(token.hyperCoreDeployer(), address(0));
        vm.stopPrank();

        assertFalse(token.isCoreTokenIndexSet());
    }

    // =====================================================================
    // Rounded transferToCore / transferToCoreFor
    // =====================================================================

    /// `transferToCore` rounds `amount` down to the nearest multiple of `coreUnit()` and
    /// returns the actually-sent amount; the remainder stays with the caller.
    function test_transferToCore_roundsDown() public {
        _mockValidCoreLink(address(token), 14); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(14);
        assertEq(token.coreUnit(), 1e10);

        vm.prank(bridge);
        token.mint(user, 10 ether);

        uint amount = 1e10 * 5 + 123; // 5 whole Core units plus dust below one Core wei
        vm.prank(user);
        uint sent = token.transferToCore(amount);

        assertEq(sent, 1e10 * 5);
        assertEq(token.balanceOf(token.coreSystemAddress()), 1e10 * 5);
        assertEq(token.balanceOf(user), 10 ether - (1e10 * 5));
    }

    /// When `evmExtraWeiDecimals <= 0`, `coreUnit()` is `1` and nothing is ever rounded.
    function test_coreUnit_isOneWhenExtraNonPositive() public {
        _mockCoreTokenInfo(15, address(token), 18, 0);
        vm.prank(address(code));
        token.setCoreTokenIndex(15);

        assertEq(token.coreUnit(), 1);
        assertEq(token.coreTransferableAmount(1234567), 1234567);
    }

    /// `transferToCoreFor` emits `Transfer(coreRecipient -> systemAddress)` for the
    /// exact rounded amount, and `coreRecipient`'s EVM balance nets to exactly zero (it only
    /// ever passes through).
    function test_transferToCoreFor_emitsEventAndNetsRecipientBalanceToZero() public {
        _mockValidCoreLink(address(token), 16);
        vm.prank(address(code));
        token.setCoreTokenIndex(16);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        address coreRecipient = makeAddr("coreRecipient");
        uint balBefore = token.balanceOf(coreRecipient);

        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(coreRecipient, systemAddr, 5 ether);

        vm.prank(user);
        uint sent = token.transferToCoreFor(coreRecipient, 5 ether);

        assertEq(sent, 5 ether);
        assertEq(token.balanceOf(coreRecipient), balBefore, "coreRecipient EVM balance must net to zero");
        assertEq(token.balanceOf(systemAddr), 5 ether);
        assertEq(token.balanceOf(user), 5 ether);
    }

    /// A `sent == 0` (amount rounds to zero) is rejected.
    function test_transferToCoreFor_zeroSent_reverts() public {
        _mockValidCoreLink(address(token), 17); // coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(17);

        vm.prank(bridge);
        token.mint(user, 10 ether);

        vm.prank(user);
        vm.expectRevert(IHyperMintableERC20.CoreAmountBelowOneCoreWei.selector);
        token.transferToCoreFor(makeAddr("coreRecipient"), 1e10 - 1);
    }

    /// Zero `coreRecipient` and `coreRecipient == coreSystemAddress()` are both rejected.
    function test_transferToCoreFor_invalidRecipients_reverts() public {
        _mockValidCoreLink(address(token), 18);
        vm.prank(address(code));
        token.setCoreTokenIndex(18);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        vm.prank(user);
        vm.expectRevert(IHyperMintableERC20.CoreRecipientZero.selector);
        token.transferToCoreFor(address(0), 1 ether);

        vm.prank(user);
        vm.expectRevert(IHyperMintableERC20.CoreRecipientIsSystemAddress.selector);
        token.transferToCoreFor(systemAddr, 1 ether);
    }

    // =====================================================================
    // Beacon / factory upgrades
    // =====================================================================

    /// Beacon upgrade preserves every token's existing storage (balances, roles, link
    /// state) and applies the new logic (`version()`) immediately.
    /// @dev The beacon is self-owned by the factory, so the upgrade goes through the
    /// factory's `upgradeBeacon` (`ADMIN_ROLE`-gated), not a direct `beacon.upgradeTo` from a
    /// separately-configured `beaconOwner` EOA.
    function test_beaconUpgrade_preservesState() public {
        vm.prank(bridge);
        token.mint(user, 10 ether);
        vm.prank(user);
        token.approve(finalizer, 5 ether);
        vm.prank(address(code));
        token.setHyperCoreDeployer(finalizer);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(factoryOwner);
        code.upgradeBeacon(address(newImpl));

        assertEq(token.balanceOf(user), 10 ether);
        assertEq(token.allowance(user, finalizer), 5 ether);
        assertEq(token.hyperCoreDeployer(), finalizer);
        assertEq(HyperMintableERC20Mock(address(token)).version(), 2);
    }

    /// A single beacon upgrade is reflected by every token sharing that beacon at once.
    function test_beaconUpgrade_affectsAllTokensAtOnce() public {
        vm.prank(bridge);
        address token2Address = code.createCrossMintableERC20(REMOTE_CHAIN_ID + 100, REMOTE_TOKEN, "T2", 6);
        HyperMintableERC20 token2 = HyperMintableERC20(token2Address);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(factoryOwner);
        code.upgradeBeacon(address(newImpl));

        assertEq(HyperMintableERC20Mock(address(token)).version(), 2);
        assertEq(HyperMintableERC20Mock(address(token2)).version(), 2);
    }

    /// Factory (UUPS) upgrade preserves storage: `beacon`, roles, and the
    /// `isHyperMintableERC20` registry all survive. (There is no separate `tokenAdmin`.)
    function test_factoryUpgrade_preservesState() public {
        assertTrue(code.isHyperMintableERC20(address(token)));

        HyperMintableERC20CodeMock newImpl = new HyperMintableERC20CodeMock();
        vm.prank(factoryOwner);
        code.upgradeToAndCall(address(newImpl), bytes(""));

        assertEq(code.beacon(), address(beacon));
        assertTrue(code.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertTrue(code.hasRole(Const.BRIDGE_ROLE, bridge));
        assertTrue(code.isHyperMintableERC20(address(token)));
        assertEq(HyperMintableERC20CodeMock(address(code)).version(), 2);
    }

    /// Upgrade authority is gated: only the factory's `ADMIN_ROLE` can upgrade the beacon
    /// (the beacon's `owner()` IS the factory, so `beacon.upgradeTo` from any EOA, even one
    /// that used to be a designated `beaconOwner`, fails; only `code.upgradeBeacon` from
    /// `ADMIN_ROLE` reaches it) or the factory itself.
    function test_upgradeAuthority_gated() public {
        HyperMintableERC20Mock newTokenImpl = new HyperMintableERC20Mock();
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user));
        beacon.upgradeTo(address(newTokenImpl));

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.upgradeBeacon(address(newTokenImpl));

        HyperMintableERC20CodeMock newCodeImpl = new HyperMintableERC20CodeMock();
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.upgradeToAndCall(address(newCodeImpl), bytes(""));
    }

    // =====================================================================
    // Initialization safety
    // =====================================================================

    /// The token LOGIC contract cannot be initialized directly.
    function test_tokenImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        tokenImplementation.initialize(tokenAdmin, bridge, address(code), "X", "X", 18);
    }

    /// The factory LOGIC contract cannot be initialized directly.
    function test_codeImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        codeImplementation.initialize(factoryOwner, bridge, address(tokenImplementation));
    }

    /// The factory PROXY is initialized atomically at construction (no separate
    /// initialize transaction is possible or needed), and cannot be initialized a second time.
    function test_factoryProxy_atomicInitialization() public {
        // `code` (from setUp) is already live with its constructor-time state — no second
        // transaction occurred between deployment and this assertion.
        assertEq(code.beacon(), address(beacon));
        assertTrue(code.hasRole(Const.ADMIN_ROLE, factoryOwner));

        vm.expectRevert(Initializable.InvalidInitialization.selector);
        code.initialize(factoryOwner, bridge, address(tokenImplementation));
    }

    // =====================================================================
    // Explicit name/symbol/minter creation path
    // =====================================================================

    /// `createHyperMintableERC20` deploys with caller-chosen name/symbol and grants
    /// `MINTER_ROLE` to the explicit `minter` argument.
    function test_createHyperMintableERC20_arbitraryNameAndMinter() public {
        address explicitMinter = makeAddr("explicitMinter");

        vm.prank(factoryOwner);
        address tokenAddress = code.createHyperMintableERC20(
            REMOTE_CHAIN_ID + 200, REMOTE_TOKEN, "My Custom Token", "MCT", 9, explicitMinter
        );

        HyperMintableERC20 customToken = HyperMintableERC20(tokenAddress);
        assertEq(customToken.name(), "My Custom Token");
        assertEq(customToken.symbol(), "MCT");
        assertEq(customToken.decimals(), 9);
        assertTrue(customToken.hasRole(Const.MINTER_ROLE, explicitMinter));
        assertTrue(code.isHyperMintableERC20(tokenAddress));
    }

    /// `createHyperMintableERC20` is gated to `ADMIN_ROLE`.
    function test_createHyperMintableERC20_requiresAdminRole() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.createHyperMintableERC20(REMOTE_CHAIN_ID + 201, REMOTE_TOKEN, "X", "X", 18, user);
    }

    /// `computeTokenAddressWithName` prediction equals the actual deployment.
    function test_computeTokenAddressWithName_matchesActual() public {
        uint remoteChainID5 = REMOTE_CHAIN_ID + 202;
        address explicitMinter = makeAddr("explicitMinter2");

        address predicted =
            code.computeTokenAddressWithName(remoteChainID5, REMOTE_TOKEN, "Name X", "SYMX", 7, explicitMinter);

        vm.prank(factoryOwner);
        address actual =
            code.createHyperMintableERC20(remoteChainID5, REMOTE_TOKEN, "Name X", "SYMX", 7, explicitMinter);

        assertEq(predicted, actual);
    }

    /// CREATE2 address prediction survives a beacon logic upgrade: the initcode
    /// only ever embeds the beacon's own fixed address, never the implementation it currently
    /// resolves to.
    function test_predictedAddress_unchangedAfterBeaconUpgrade() public {
        uint remoteChainID6 = REMOTE_CHAIN_ID + 203;

        address predictedBefore = code.computeTokenAddress(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(factoryOwner);
        code.upgradeBeacon(address(newImpl));

        address predictedAfter = code.computeTokenAddress(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedBefore, predictedAfter);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(actual, predictedBefore);
        assertEq(HyperMintableERC20Mock(actual).version(), 2);
    }

    // =====================================================================
    // End-to-end `BridgeExecutor` extra-call path for `transferToCoreFor`
    // =====================================================================

    /// `transferToCoreFor` exercised through the REAL `BridgeExecutor.executeExtraCall`
    /// path (not a direct call): the extra-call's whole point is that Core credits
    /// the USER, never the executor. `value` is deliberately not a multiple of `coreUnit()` so
    /// a genuine truncation occurs; `to` is set to the executor itself so the truncated dust
    /// that `executeExtraCall` returns to `to` is a self-transfer (stays with the executor),
    /// keeping the executor's own balance change equal to exactly `sent` and leaving `user`'s
    /// balance untouched by anything except the pass-through (which nets to zero).
    function test_transferToCoreFor_viaBridgeExecutor_creditsUserNotExecutor() public {
        _mockValidCoreLink(address(token), 28); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(28);
        address systemAddr = token.coreSystemAddress();

        address executorCaller = makeAddr("executorCaller"); // holds EXECUTOR_ROLE, mimics the bridge
        BridgeExecutor executor = new BridgeExecutor(factoryOwner, executorCaller);

        vm.startPrank(factoryOwner);
        executor.addWhitelistTarget(address(token));
        executor.setMethodCheckEnabled(address(token), true);
        bytes4[] memory methods = new bytes4[](1);
        methods[0] = IHyperMintableERC20.transferToCoreFor.selector;
        executor.addWhitelistMethods(address(token), methods);
        vm.stopPrank();

        uint value = 1e10 * 7 + 42; // 7 whole Core units plus dust below one Core wei
        uint expectedSent = 1e10 * 7;

        vm.prank(bridge);
        token.mint(executorCaller, value);
        vm.prank(executorCaller);
        token.approve(address(executor), value);

        bytes memory extraCalldata = abi.encodeWithSelector(IHyperMintableERC20.transferToCoreFor.selector, user, value);
        bytes memory extraData = abi.encodePacked(address(token), extraCalldata);

        uint userBalBefore = token.balanceOf(user);
        assertEq(token.balanceOf(address(executor)), 0, "executor starts with no balance");

        // The final leg of transferToCoreFor's pass-through: Transfer(user -> systemAddr).
        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(user, systemAddr, expectedSent);

        vm.prank(executorCaller);
        (uint consumed,) = executor.executeExtraCall(0, 0, IERC20(address(token)), address(executor), value, extraData);

        // `executeExtraCall` pulled the full `value` into the executor before the target call,
        // so the executor's balance change across the whole call is `value - sent` — i.e. its
        // balance dropped by exactly `sent` off the peak it briefly held, and only the truncated
        // dust (`value - sent`) is left behind (self-transferred back to it as `remaining`).
        assertEq(consumed, expectedSent, "consumed must equal the coreUnit-rounded sent amount");
        assertEq(token.balanceOf(user), userBalBefore, "user EVM balance must net to zero");
        assertEq(token.balanceOf(systemAddr), expectedSent);
        assertEq(
            token.balanceOf(address(executor)),
            value - expectedSent,
            "executor is left holding only the truncated remainder (down by exactly `sent` from the value it pulled in)"
        );
    }

    /// The PRODUCTION shape of the BridgeExecutor case above: `to` is the USER, which is what the
    /// bridge actually passes. That case deliberately set `to` to the executor so the dust self-transfer could not
    /// disturb its "user nets to zero" assertion — that left one thing unproven, namely that
    /// the truncated remainder is genuinely handed back to the user rather than stranded in
    /// the executor. This case closes exactly that gap: the pass-through still credits Core to
    /// the user (`Transfer(user -> systemAddr)`), and on top of it `executeExtraCall` returns
    /// `remaining == value - consumed` to `to`, so the user ends up **up by the dust** and the
    /// executor ends up with nothing.
    function test_transferToCoreFor_viaBridgeExecutor_dustReturnedToUser() public {
        _mockValidCoreLink(address(token), 29); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(29);
        address systemAddr = token.coreSystemAddress();

        address executorCaller = makeAddr("executorCaller29"); // holds EXECUTOR_ROLE, mimics the bridge
        BridgeExecutor executor = new BridgeExecutor(factoryOwner, executorCaller);

        vm.startPrank(factoryOwner);
        executor.addWhitelistTarget(address(token));
        executor.setMethodCheckEnabled(address(token), true);
        bytes4[] memory methods = new bytes4[](1);
        methods[0] = IHyperMintableERC20.transferToCoreFor.selector;
        executor.addWhitelistMethods(address(token), methods);
        vm.stopPrank();

        uint value = 1e10 * 7 + 42; // 7 whole Core units plus dust below one Core wei
        uint expectedSent = 1e10 * 7;
        uint expectedDust = value - expectedSent; // 42

        vm.prank(bridge);
        token.mint(executorCaller, value);
        vm.prank(executorCaller);
        token.approve(address(executor), value);

        bytes memory extraCalldata = abi.encodeWithSelector(IHyperMintableERC20.transferToCoreFor.selector, user, value);
        bytes memory extraData = abi.encodePacked(address(token), extraCalldata);

        uint userBalBefore = token.balanceOf(user);
        assertEq(token.balanceOf(address(executor)), 0, "executor starts with no balance");

        // Core still credits the USER, exactly as in the case above — `to` does not affect that leg.
        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(user, systemAddr, expectedSent);

        // `to = user`: the real bridge always passes the recipient here.
        vm.prank(executorCaller);
        (uint consumed,) = executor.executeExtraCall(0, 0, IERC20(address(token)), user, value, extraData);

        assertEq(consumed, expectedSent, "consumed must equal the coreUnit-rounded sent amount");
        assertEq(token.balanceOf(systemAddr), expectedSent, "Core-bound amount is the rounded value");

        // The point of this case: the dust is NOT stranded. `remaining = value - consumed` is
        // transferred to `to`, so the user is up by exactly the truncated remainder.
        assertEq(token.balanceOf(user), userBalBefore + expectedDust, "user receives the truncated remainder back");
        assertEq(token.balanceOf(address(executor)), 0, "executor keeps nothing when `to` is the user");
    }

    // =====================================================================
    // `transferToCoreFor` called DIRECTLY (no BridgeExecutor)
    // =====================================================================

    /// Direct call with an `amount` that is NOT a multiple of `coreUnit()`. The event/zero-net-
    /// balance case above pins the pass-through with an exact multiple, where there is no remainder to place at all;
    /// this case is the one that actually distinguishes where the remainder goes. It stays with
    /// the CALLER, exactly as `transferToCore` does — `coreRecipient` is a Core identity, so it
    /// receives a Core credit and its EVM balance still nets to zero. Only the executor path
    /// (the BridgeExecutor cases above) route the remainder onward, and they do so through `executeExtraCall`'s
    /// `remaining` refund to `to`, not through this function.
    function test_transferToCoreFor_directCall_dustStaysWithCaller() public {
        _mockValidCoreLink(address(token), 30); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(30);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        address coreRecipient = makeAddr("coreRecipient30");
        uint amount = 1e10 * 3 + 77; // 3 whole Core units plus dust below one Core wei
        uint expectedSent = 1e10 * 3;

        uint callerBalBefore = token.balanceOf(user);
        uint coreRecipientBalBefore = token.balanceOf(coreRecipient);

        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(coreRecipient, systemAddr, expectedSent);

        vm.prank(user);
        uint sent = token.transferToCoreFor(coreRecipient, amount);

        assertEq(sent, expectedSent, "sent must be the coreUnit-rounded amount");
        assertEq(token.balanceOf(user), callerBalBefore - expectedSent, "caller is debited only `sent`; the remainder stays with it");
        assertEq(
            token.balanceOf(coreRecipient),
            coreRecipientBalBefore,
            "coreRecipient EVM balance nets to zero even when `amount` truncates"
        );
        assertEq(token.balanceOf(systemAddr), expectedSent, "system address receives exactly sent");
    }

    /// `coreRecipient == caller` (self-aliasing). Both legs then touch the same account, so
    /// the call degenerates to exactly `transferToCore`: the caller is down by `sent` and keeps
    /// the remainder. Pinned because the aliased pass-through is the one shape where the two
    /// `_transfer` calls could interfere with each other.
    function test_transferToCoreFor_selfRecipient_retainsDust() public {
        _mockValidCoreLink(address(token), 31); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(address(code));
        token.setCoreTokenIndex(31);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        uint amount = 1e10 * 5 + 13; // 5 whole Core units plus dust below one Core wei
        uint expectedSent = 1e10 * 5;

        uint userBalBefore = token.balanceOf(user);

        vm.prank(user);
        uint sent = token.transferToCoreFor(user, amount);

        assertEq(sent, expectedSent, "sent must be the coreUnit-rounded amount");
        assertEq(token.balanceOf(user), userBalBefore - expectedSent, "caller/coreRecipient is down by sent only");
        assertEq(token.balanceOf(systemAddr), expectedSent, "system address receives exactly sent");
    }

    // =====================================================================
    // Dual creation events. Every path that reaches
    // `_create` must emit BOTH `CrossMintableERC20Created` and `HyperMintableERC20Created`,
    // EXACTLY ONCE each, for the exact same token creation — never zero, never twice. Consumers
    // subscribed to both topics must dedupe by tx hash / token address (see the NatSpec on both
    // events); these tests pin down the "exactly once" half of that contract precisely so a future
    // regression (e.g. someone re-inlining `_create` in a subclass) is caught immediately.
    // =====================================================================

    /// @dev Counts logs emitted BY `emitter` whose first topic (event selector) equals `topic0`.
    function _countMatchingLogs(Vm.Log[] memory logs, address emitter, bytes32 topic0) internal pure returns (uint count) {
        for (uint i = 0; i < logs.length; i++) {
            if (logs[i].emitter == emitter && logs[i].topics.length > 0 && logs[i].topics[0] == topic0) {
                count++;
            }
        }
    }

    /// @dev Asserts both creation events fired from `code`, targeting `(remoteChainIDX, REMOTE_TOKEN,
    /// tokenAddress)`, exactly once each — via `vm.expectEmit` (exact field match, in emission
    /// order: Cross event first per `HyperMintableERC20Code._emitCreated`) AND an independent
    /// `vm.recordLogs` count (catches an accidental duplicate/omission that a same-shaped
    /// `expectEmit` pair alone would not).
    function _expectBothCreationEvents(uint remoteChainIDX, address tokenAddress) internal {
        vm.expectEmit(true, true, true, true, address(code));
        emit ICrossMintableERC20V2Code.CrossMintableERC20Created(remoteChainIDX, REMOTE_TOKEN, tokenAddress);

        vm.expectEmit(true, true, false, true, address(code));
        emit IHyperMintableERC20Code.HyperMintableERC20Created(remoteChainIDX, REMOTE_TOKEN, tokenAddress);
    }

    /// Path (1): the inherited legacy bridge-only entrypoint (`createCrossMintableERC20`,
    /// `BRIDGE_ROLE`).
    function test_dualCreationEvents_bridgePath() public {
        uint remoteChainIDX = REMOTE_CHAIN_ID + 900;
        address predicted = code.computeTokenAddress(remoteChainIDX, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        _expectBothCreationEvents(remoteChainIDX, predicted);
        vm.recordLogs();
        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainIDX, REMOTE_TOKEN, SYMBOL, DECIMALS);
        assertEq(actual, predicted);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(
            _countMatchingLogs(logs, address(code), ICrossMintableERC20V2Code.CrossMintableERC20Created.selector),
            1,
            "CrossMintableERC20Created must fire exactly once"
        );
        assertEq(
            _countMatchingLogs(logs, address(code), IHyperMintableERC20Code.HyperMintableERC20Created.selector),
            1,
            "HyperMintableERC20Created must fire exactly once"
        );
    }

    /// Path (2): the inherited explicit creation entrypoint (`createMintableERC20`,
    /// `ADMIN_ROLE`) — reaches `_create` without going through either bridge-specific wrapper.
    function test_dualCreationEvents_explicitAdminPath() public {
        uint remoteChainIDX = REMOTE_CHAIN_ID + 901;
        address explicitMinter = makeAddr("explicitAdminPathMinter");
        address predicted =
            code.computeTokenAddressWithName(remoteChainIDX, REMOTE_TOKEN, "Explicit Admin Path", "EAP", 9, explicitMinter);

        _expectBothCreationEvents(remoteChainIDX, predicted);
        vm.recordLogs();
        vm.prank(factoryOwner);
        address actual =
            code.createMintableERC20(remoteChainIDX, REMOTE_TOKEN, "Explicit Admin Path", "EAP", 9, explicitMinter);
        assertEq(actual, predicted);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(
            _countMatchingLogs(logs, address(code), ICrossMintableERC20V2Code.CrossMintableERC20Created.selector),
            1,
            "CrossMintableERC20Created must fire exactly once"
        );
        assertEq(
            _countMatchingLogs(logs, address(code), IHyperMintableERC20Code.HyperMintableERC20Created.selector),
            1,
            "HyperMintableERC20Created must fire exactly once"
        );
    }

    /// Path (3): the Hyper-named compat wrapper (`createHyperMintableERC20`, `ADMIN_ROLE`)
    /// — the path most likely to be watched by legacy tooling that only knows the Hyper topic.
    function test_dualCreationEvents_hyperCompatWrapperPath() public {
        uint remoteChainIDX = REMOTE_CHAIN_ID + 902;
        address explicitMinter = makeAddr("hyperCompatWrapperMinter");
        address predicted =
            code.computeTokenAddressWithName(remoteChainIDX, REMOTE_TOKEN, "Hyper Compat Wrapper", "HCW", 11, explicitMinter);

        _expectBothCreationEvents(remoteChainIDX, predicted);
        vm.recordLogs();
        vm.prank(factoryOwner);
        address actual =
            code.createHyperMintableERC20(remoteChainIDX, REMOTE_TOKEN, "Hyper Compat Wrapper", "HCW", 11, explicitMinter);
        assertEq(actual, predicted);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(
            _countMatchingLogs(logs, address(code), ICrossMintableERC20V2Code.CrossMintableERC20Created.selector),
            1,
            "CrossMintableERC20Created must fire exactly once"
        );
        assertEq(
            _countMatchingLogs(logs, address(code), IHyperMintableERC20Code.HyperMintableERC20Created.selector),
            1,
            "HyperMintableERC20Created must fire exactly once"
        );
    }
}

/// @dev Minimal "upgraded" token implementation used only by the beacon-upgrade tests below:
/// identical storage layout and behavior to `HyperMintableERC20`, plus one new function
/// so the tests can observe that an upgrade actually took effect.
contract HyperMintableERC20Mock is HyperMintableERC20 {
    function version() external pure returns (uint) {
        return 2;
    }
}

/// @dev Minimal "upgraded" factory implementation used only by the factory-upgrade test below.
contract HyperMintableERC20CodeMock is HyperMintableERC20Code {
    function version() external pure returns (uint) {
        return 2;
    }
}
