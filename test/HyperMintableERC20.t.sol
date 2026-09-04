// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeExecutor} from "../src/BridgeExecutor.sol";
import {Const} from "../src/lib/Const.sol";
import {HyperMintableERC20} from "../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {IHyperMintableERC20} from "../src/token/IHyperMintableERC20.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {Errors} from "@openzeppelin/contracts/utils/Errors.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/Test.sol";

/**
 * @title HyperMintableERC20Test
 * @notice T1-T18 + T23-T28 (proxied setup) plus A1-A27 of the HyperCore link hardening +
 *         proxy + name/symbol test plan (spec §14).
 * @dev `test/CrossBridgeV2HyperEVMRoute.t.sol` covers T19-T22, the end-to-end path through a
 *      real bridge instance.
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
    // A1/A2 actually verify independent, forge-computed values (`cast index-erc7201 <id>`),
    // not merely that the contract agrees with itself.
    bytes32 internal constant TOKEN_STORAGE_SLOT = 0xc65afee183f44952959f664e5c2b8a5e4916480c324e2787b2215d744391d400;
    bytes32 internal constant FACTORY_STORAGE_SLOT = 0xd55591ae453b55973909532d0252f41e2a7c543bd5a578ad3871068e0c28b400;
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
    address public tokenAdmin; // initial defaultAdmin() of every created token; not granted LINKER_ROLE — derives link authority dynamically via isLinkAuthority()
    address public bridge; // BRIDGE_ROLE on the factory, MINTER_ROLE on created tokens
    address public user;
    address public finalizer;
    address public beaconOwner; // UpgradeableBeacon owner() — recommended multisig/timelock in production

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
        beaconOwner = makeAddr("beaconOwner");

        tokenImplementation = new HyperMintableERC20();
        beacon = new UpgradeableBeacon(address(tokenImplementation), beaconOwner);
        codeImplementation = new HyperMintableERC20Code();

        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(HyperMintableERC20Code.initialize, (factoryOwner, tokenAdmin, bridge, address(beacon)))
        );
        code = HyperMintableERC20Code(address(proxy));
        console.log("HyperMintableERC20Code", address(code));

        vm.prank(bridge);
        token = HyperMintableERC20(code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS));
        console.log("HyperMintableERC20", address(token));
    }

    function _deployCode(address factoryOwnerAddr, address tokenAdminAddr, address bridgeAddr)
        internal
        returns (HyperMintableERC20Code freshCode)
    {
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(codeImplementation),
            abi.encodeCall(
                HyperMintableERC20Code.initialize, (factoryOwnerAddr, tokenAdminAddr, bridgeAddr, address(beacon))
            )
        );
        freshCode = HyperMintableERC20Code(address(proxy));
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
    /// default to the realistic TONE link (request.md R4): `weiDecimals=8`,
    /// `evmExtraWeiDecimals=10`, summing to `DECIMALS` (18) so `setCoreTokenIndex`'s
    /// decimals cross-check passes unless the test deliberately breaks it.
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

    /// @dev AC-1 bonus: `Const.LINKER_ROLE` matches its documented digest.
    function test_constLinkerRole() public pure {
        assertEq(Const.LINKER_ROLE, keccak256("LINKER_ROLE"));
    }

    /// T1. The slot constant is a raw `keccak256("HyperCore deployer")`, never an ERC-7201
    /// masked digest — a typo here silently breaks the only usable `finalizeEvmContract` path.
    function test_T1_hypercoreDeployerSlot_isRawKeccak() public view {
        assertEq(token.HYPERCORE_DEPLOYER_SLOT(), keccak256("HyperCore deployer"));
        assertEq(token.HYPERCORE_DEPLOYER_SLOT(), 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f);
    }

    /// T2. `setHyperCoreDeployer` writes the RAW slot, not just something the getter agrees with.
    function test_T2_setHyperCoreDeployer_rawSlot() public {
        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);

        bytes32 raw = vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT());
        assertEq(raw, bytes32(uint(uint160(finalizer))));
        assertEq(token.hyperCoreDeployer(), finalizer);
    }

    /// T3. No aliasing in either direction between the link slot and ordinary ERC20/AccessControl storage.
    function test_T3_noStorageAliasing() public {
        vm.prank(bridge);
        token.mint(user, INITIAL_SUPPLY);

        vm.prank(user);
        token.approve(finalizer, 123 ether);

        vm.prank(user);
        token.transfer(tokenAdmin, 1 ether);

        uint deadline = block.timestamp + 1 days;
        (uint8 v, bytes32 r, bytes32 s) = _signPermit(USER_PK, user, finalizer, 5 ether, deadline);
        token.permit(user, finalizer, 5 ether, deadline, v, r, s);

        vm.prank(tokenAdmin);
        token.grantRole(Const.MINTER_ROLE, user);

        // Direction 1: heavy ordinary activity must not disturb the (still-unset) link slot.
        bytes32 rawAfterActivity = vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT());
        assertEq(rawAfterActivity, bytes32(0));

        // Direction 2: writing the link slot must not disturb balances/allowance/nonces/roles.
        uint balBefore = token.balanceOf(user);
        uint allowBefore = token.allowance(user, finalizer);
        uint noncesBefore = token.nonces(user);
        bool hasRoleBefore = token.hasRole(Const.MINTER_ROLE, user);

        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);

        assertEq(token.balanceOf(user), balBefore);
        assertEq(token.allowance(user, finalizer), allowBefore);
        assertEq(token.nonces(user), noncesBefore);
        assertEq(token.hasRole(Const.MINTER_ROLE, user), hasRoleBefore);
    }

    /// T4. Finalizer can be recalled back to zero.
    function test_T4_resetFinalizerToZero() public {
        vm.startPrank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);
        token.setHyperCoreDeployer(address(0));
        vm.stopPrank();

        assertEq(token.hyperCoreDeployer(), address(0));
        assertEq(vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT()), bytes32(0));
    }

    // ---------------------------------------------------------------------
    // Permissions
    // ---------------------------------------------------------------------

    /// T5. Factory (ADMIN_ROLE) path succeeds, token defaultAdmin direct path succeeds, third party reverts.
    function test_T5_setHyperCoreDeployer_viaFactory_andDirectly_thirdPartyReverts() public {
        vm.prank(factoryOwner);
        code.setHyperCoreDeployer(address(token), finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        address finalizer2 = makeAddr("finalizer2");
        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer2);
        assertEq(token.hyperCoreDeployer(), finalizer2);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);
    }

    /// T6. Token admin revoking the factory's LINKER_ROLE blocks the factory path (fallback still works).
    function test_T6_revokeFactoryLinkerRole_blocksFactoryPath() public {
        vm.prank(tokenAdmin);
        token.revokeRole(Const.LINKER_ROLE, address(code));

        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(code), Const.LINKER_ROLE
            )
        );
        code.setHyperCoreDeployer(address(token), finalizer);

        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);
    }

    /// T7. Factory setters require ADMIN_ROLE.
    function test_T7_factorySetters_requireAdminRole() public {
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

    /// T8. Regression: deployer != initialOwner must not revert the factory's atomic proxy
    /// initialization (the trap the V2 factory falls into by using public `grantRole` instead
    /// of `_grantRole`).
    function test_T8_deployerNotInitialOwner_doesNotRevert() public {
        address deployer = makeAddr("someoneElse");
        vm.prank(deployer);
        HyperMintableERC20Code freshCode = _deployCode(factoryOwner, tokenAdmin, bridge);

        assertTrue(freshCode.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertFalse(freshCode.hasRole(Const.ADMIN_ROLE, deployer));
        assertEq(freshCode.tokenAdmin(), tokenAdmin);
    }

    // ---------------------------------------------------------------------
    // Core system address
    // ---------------------------------------------------------------------

    /// T9. index 200 -> 0x20000000000000000000000000000000000000c8 exactly.
    function test_T9_coreSystemAddress_index200() public {
        _mockValidCoreLink(address(token), 200);
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(200);

        address expected = 0x20000000000000000000000000000000000000C8;
        assertEq(token.coreSystemAddress(), expected);
    }

    /// T10. index 0 is settable and valid (not an "unset" sentinel).
    function test_T10_coreTokenIndexZero_isValid() public {
        _mockValidCoreLink(address(token), 0);
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(0);

        assertTrue(token.isCoreTokenIndexSet());
        assertEq(token.coreTokenIndex(), 0);
        assertEq(token.coreSystemAddress(), address(uint160(0x20) << 152));
    }

    /// T11. Unset index makes both `coreSystemAddress` and `transferToCore` revert.
    function test_T11_unsetIndex_reverts() public {
        vm.expectRevert(HyperMintableERC20.HyperMintableERC20CoreTokenIndexNotSet.selector);
        token.coreSystemAddress();

        vm.prank(user);
        vm.expectRevert(HyperMintableERC20.HyperMintableERC20CoreTokenIndexNotSet.selector);
        token.transferToCore(1);
    }

    /// T12. `transferToCore` moves balance to the system address and emits `Transfer`; insufficient balance reverts.
    function test_T12_transferToCore() public {
        _mockValidCoreLink(address(token), 200);
        vm.prank(tokenAdmin);
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

    /// T13. Fuzz: for any in-range index, the high byte is always 0x20 and the low 19 bytes are the index.
    function testFuzz_T13_coreSystemAddress(uint32 index) public {
        _mockValidCoreLink(address(token), index);
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(index);

        uint160 sys = uint160(token.coreSystemAddress());
        assertEq(sys >> 152, uint160(0x20));
        assertEq(sys & ((uint160(1) << 152) - 1), uint160(index));
    }

    // ---------------------------------------------------------------------
    // ERC20 / factory
    // ---------------------------------------------------------------------

    /// T14. mint/burn work normally; missing MINTER_ROLE reverts.
    function test_T14_mintBurn() public {
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

    /// T15. EIP-2612 permit signature verification, domain name = "Cross Bridge <SYM>".
    function test_T15_permit() public {
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

    /// T16. `computeTokenAddress` prediction equals the actual `createCrossMintableERC20` result.
    function test_T16_computeTokenAddress_matchesActual() public view {
        address predicted = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predicted, address(token));
    }

    /// T16a. Same (remoteChainID, remoteToken) but different symbol/decimals predicts a different address.
    function test_T16a_differentSymbolOrDecimals_changesPredictedAddress() public view {
        address predicted = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        address predictedOtherSymbol =
            code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, "OTHER", DECIMALS, bridge);
        address predictedOtherDecimals = code.computeTokenAddress(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, 6, bridge);

        assertTrue(predicted != predictedOtherSymbol);
        assertTrue(predicted != predictedOtherDecimals);
    }

    /// T16b. `isHyperMintableERC20` is true right after creation; setter on an unknown address reverts.
    function test_T16b_isHyperMintableERC20_andUnknownTokenReverts() public {
        assertTrue(code.isHyperMintableERC20(address(token)));
        assertFalse(code.isHyperMintableERC20(address(0xDEAD)));

        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(HyperMintableERC20Code.HyperMintableERC20CodeUnknownToken.selector, address(0xDEAD))
        );
        code.setHyperCoreDeployer(address(0xDEAD), finalizer);
    }

    /// T16c. Factory replacement: factory B's ADMIN cannot touch a token factory A created;
    /// factory A's ADMIN can; the token's own defaultAdmin always can.
    function test_T16c_factoryReplacement_scopedToCreatingFactory() public {
        address factoryBOwner = makeAddr("factoryBOwner");
        HyperMintableERC20Code codeB = _deployCode(factoryBOwner, tokenAdmin, bridge);

        vm.prank(factoryBOwner);
        vm.expectRevert(
            abi.encodeWithSelector(HyperMintableERC20Code.HyperMintableERC20CodeUnknownToken.selector, address(token))
        );
        codeB.setHyperCoreDeployer(address(token), finalizer);

        vm.prank(factoryOwner);
        code.setHyperCoreDeployer(address(token), finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        address finalizer2 = makeAddr("finalizer2");
        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer2);
        assertEq(token.hyperCoreDeployer(), finalizer2);
    }

    /// T16d. Determinism holds even when the address that computes the prediction differs from
    /// the address that actually calls `createCrossMintableERC20` — as long as `minter` matches
    /// the real creator (the bridge).
    function test_T16d_determinism_computeCallerDiffersFromCreateCaller() public {
        uint remoteChainID2 = REMOTE_CHAIN_ID + 1;

        address predictedByOwner = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(user);
        address predictedByOther = code.computeTokenAddress(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedByOwner, predictedByOther);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID2, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(actual, predictedByOwner);
    }

    /// T16e. Transferring the factory's OWN default admin (2-step) does not change what
    /// `defaultAdmin()` newly created tokens end up with, nor the predicted address — both are
    /// pinned to the factory's `tokenAdmin` storage.
    function test_T16e_factoryDefaultAdminTransfer_doesNotAffectTokenAdminOrPrediction() public {
        address newFactoryOwner = makeAddr("newFactoryOwner");
        uint remoteChainID3 = REMOTE_CHAIN_ID + 2;

        address predictedBefore = code.computeTokenAddress(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        vm.prank(factoryOwner);
        code.beginDefaultAdminTransfer(newFactoryOwner);
        vm.warp(block.timestamp + 1);
        vm.prank(newFactoryOwner);
        code.acceptDefaultAdminTransfer();

        assertEq(code.defaultAdmin(), newFactoryOwner);
        assertEq(code.tokenAdmin(), tokenAdmin);

        address predictedAfter = code.computeTokenAddress(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedAfter, predictedBefore);

        vm.prank(bridge);
        address newToken = code.createCrossMintableERC20(remoteChainID3, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(newToken, predictedBefore);
        assertEq(HyperMintableERC20(newToken).defaultAdmin(), tokenAdmin);
    }

    /// T16f. Negative: predicting with the wrong `minter` yields an address that differs from
    /// what the real bridge-driven creation actually deploys.
    function test_T16f_wrongMinterArg_predictedDiffersFromActual() public {
        uint remoteChainID4 = REMOTE_CHAIN_ID + 3;
        address wrongMinter = makeAddr("wrongMinter");

        address predictedWrong = code.computeTokenAddress(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS, wrongMinter);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID4, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertTrue(predictedWrong != actual);
    }

    /// T17. Redeploying with the same salt + same args reverts (CREATE2 idempotency).
    function test_T17_redeploySameSaltAndArgs_reverts() public {
        vm.prank(bridge);
        vm.expectRevert(Errors.FailedDeployment.selector);
        code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS);
    }

    /// T18. EIP-170: runtime code of both LOGIC implementations (what the 24,576-byte limit
    /// actually constrains now that both token and factory are proxies) stays under the limit;
    /// the proxies themselves are trivially small.
    function test_T18_runtimeCodeSize_underEip170Limit() public view {
        assertLt(address(tokenImplementation).code.length, 24_576);
        assertLt(address(codeImplementation).code.length, 24_576);
        assertLt(address(token).code.length, 24_576);
        assertLt(address(code).code.length, 24_576);
    }

    // ---------------------------------------------------------------------
    // Link authority (H1: exactly two principals — current defaultAdmin() and factoryLinker())
    // ---------------------------------------------------------------------

    /// T23. After a completed default-admin transfer, the *old* admin's link authority is gone:
    /// both link setters revert with `AccessControlUnauthorizedAccount(oldAdmin, LINKER_ROLE)`.
    function test_T23_oldAdminLosesLinkAuthorityAfterTransfer() public {
        address newAdmin = makeAddr("newAdmin");

        vm.prank(tokenAdmin);
        token.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        token.acceptDefaultAdminTransfer();

        assertEq(token.defaultAdmin(), newAdmin);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, tokenAdmin, Const.LINKER_ROLE
            )
        );
        token.setHyperCoreDeployer(finalizer);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, tokenAdmin, Const.LINKER_ROLE
            )
        );
        token.setCoreTokenIndex(1);
    }

    /// T24. The *new* admin gains link authority immediately, with no separate grant needed, and
    /// can actually write the raw slot / set the Core index.
    function test_T24_newAdminGetsLinkAuthorityImmediately() public {
        address newAdmin = makeAddr("newAdmin");

        vm.prank(tokenAdmin);
        token.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        token.acceptDefaultAdminTransfer();

        vm.prank(newAdmin);
        token.setHyperCoreDeployer(finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);
        bytes32 raw = vm.load(address(token), token.HYPERCORE_DEPLOYER_SLOT());
        assertEq(raw, bytes32(uint(uint160(finalizer))));

        _mockValidCoreLink(address(token), 7);
        vm.prank(newAdmin);
        token.setCoreTokenIndex(7);
        assertEq(token.coreTokenIndex(), 7);
    }

    /// T25. A third party (never granted any role) reverts both before AND after the admin
    /// transfer.
    function test_T25_thirdParty_revertsBeforeAndAfterTransfer() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);

        address newAdmin = makeAddr("newAdmin");
        vm.prank(tokenAdmin);
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

    /// T26. Granting `LINKER_ROLE` to a third party flips `hasRole` but not `isLinkAuthority` —
    /// R6 (link authority is exactly two principals) is enforced in code, not just documented.
    function test_T26_thirdPartyGrantedLinkerRole_stillReverts() public {
        vm.prank(tokenAdmin);
        token.grantRole(Const.LINKER_ROLE, user);

        assertTrue(token.hasRole(Const.LINKER_ROLE, user));
        assertFalse(token.isLinkAuthority(user));

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.LINKER_ROLE)
        );
        token.setHyperCoreDeployer(finalizer);
    }

    /// T27. `isLinkAuthority` matches actual success/failure for five cases: old admin, new
    /// admin, factory, a third party holding only the role, and the zero address.
    function test_T27_isLinkAuthority_matchesActualOutcome() public {
        address newAdmin = makeAddr("newAdmin");
        vm.prank(tokenAdmin);
        token.beginDefaultAdminTransfer(newAdmin);
        vm.warp(block.timestamp + 1);
        vm.prank(newAdmin);
        token.acceptDefaultAdminTransfer();

        vm.prank(newAdmin);
        token.grantRole(Const.LINKER_ROLE, user);

        assertFalse(token.isLinkAuthority(tokenAdmin)); // old admin
        assertTrue(token.isLinkAuthority(newAdmin)); // new admin
        assertTrue(token.isLinkAuthority(address(code))); // factory (factoryLinker)
        assertFalse(token.isLinkAuthority(user)); // role-only third party
        assertFalse(token.isLinkAuthority(address(0))); // zero address
    }

    /// T28. Revoking the factory's `LINKER_ROLE` blocks the factory path even though
    /// `factoryLinker` cannot itself be reassigned; re-granting it restores the factory path.
    function test_T28_revokeAndRegrantFactoryLinkerRole() public {
        assertEq(token.factoryLinker(), address(code));

        vm.prank(tokenAdmin);
        token.revokeRole(Const.LINKER_ROLE, address(code));

        vm.prank(factoryOwner);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(code), Const.LINKER_ROLE
            )
        );
        code.setHyperCoreDeployer(address(token), finalizer);

        vm.prank(tokenAdmin);
        token.grantRole(Const.LINKER_ROLE, address(code));

        vm.prank(factoryOwner);
        code.setHyperCoreDeployer(address(token), finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);
    }

    // =====================================================================
    // A1-A2. ERC-7201 / ERC-1967 slot hygiene
    // =====================================================================

    /// A1. The token's and factory's namespaced storage slots never collide with each other,
    /// with the raw HyperCore deployer slot, or with the ERC-1967 beacon/implementation slots
    /// the proxies themselves rely on. Also checks the ERC-7201 masking convention (low byte
    /// zeroed) actually holds for both computed constants.
    function test_A1_namespacedSlots_noCollision() public view {
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

    /// A2. Token state actually lives at the raw `TOKEN_STORAGE_SLOT`, packed into a single
    /// word exactly as documented (spec §8): `coreTokenIndex` (uint64) | `coreTokenIndexSet`
    /// (bool) | `coreExtraWeiDecimals` (int8) | `decimals` (uint8) | `factoryLinker` (address).
    function test_A2_tokenStorage_rawSlotPacking() public {
        _mockValidCoreLink(address(token), 5);
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(5);

        bytes32 raw = vm.load(address(token), TOKEN_STORAGE_SLOT);
        uint word = uint(raw);

        // Solidity packs the FIRST-declared struct member into the LOW-order end of the slot:
        // coreTokenIndex (bits 0-63) | coreTokenIndexSet (64-71) | coreExtraWeiDecimals (72-79)
        // | decimals (80-87) | factoryLinker (88-247).
        uint64 decodedIndex = uint64(word);
        bool decodedSet = uint8(word >> 64) != 0;
        int8 decodedExtra = int8(uint8(word >> 72));
        uint8 decodedDecimals = uint8(word >> 80);
        address decodedFactoryLinker = address(uint160(word >> 88));

        assertEq(decodedFactoryLinker, token.factoryLinker());
        assertEq(decodedDecimals, token.decimals());
        assertEq(decodedExtra, token.coreExtraWeiDecimals());
        assertEq(decodedSet, token.isCoreTokenIndexSet());
        assertEq(decodedIndex, token.coreTokenIndex());
    }

    // =====================================================================
    // A3-A9. setCoreTokenIndex validation + one-shot lock
    // =====================================================================

    /// A3. A correctly linked index succeeds and records the exact index + evmExtraWeiDecimals.
    function test_A3_setCoreTokenIndex_validLink_succeeds() public {
        _mockValidCoreLink(address(token), 2895);

        vm.expectEmit(true, false, false, true, address(token));
        emit IHyperMintableERC20.CoreTokenIndexSet(2895, 10);

        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(2895);

        assertTrue(token.isCoreTokenIndexSet());
        assertEq(token.coreTokenIndex(), 2895);
        assertEq(token.coreExtraWeiDecimals(), 10);
    }

    /// A4. Unlinked index (`evmContract == address(0)`) is rejected.
    function test_A4_setCoreTokenIndex_unlinked_reverts() public {
        _mockCoreTokenInfo(3, address(0), 8, 10);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreLinkNotFinalized.selector, uint64(3), address(0))
        );
        token.setCoreTokenIndex(3);
    }

    /// A5. Index linked to a DIFFERENT contract is rejected — this is exactly the spot-pair
    /// vs. token-index mix-up from request.md R4 (index 2902 resolves to a different token).
    function test_A5_setCoreTokenIndex_linkedToOtherContract_reverts() public {
        address otherToken = makeAddr("otherHyperCoreToken");
        _mockCoreTokenInfo(2902, otherToken, 8, 10);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreLinkNotFinalized.selector, uint64(2902), otherToken)
        );
        token.setCoreTokenIndex(2902);
    }

    /// A6. Precompile revert (e.g. wildly out-of-range index) is surfaced as `CoreTokenInfoUnavailable`.
    function test_A6_setCoreTokenIndex_precompileReverts_reverts() public {
        vm.mockCallRevert(CORE_TOKEN_INFO_PRECOMPILE, abi.encode(uint32(999_999)), bytes("precompile revert"));

        vm.prank(tokenAdmin);
        vm.expectRevert(abi.encodeWithSelector(IHyperMintableERC20.CoreTokenInfoUnavailable.selector, uint64(999_999)));
        token.setCoreTokenIndex(999_999);
    }

    /// A7. `index > type(uint32).max` is rejected before ever reaching the precompile.
    function test_A7_setCoreTokenIndex_indexOutOfUint32Range_reverts() public {
        uint64 tooLarge = uint64(type(uint32).max) + 1;

        vm.prank(tokenAdmin);
        vm.expectRevert(abi.encodeWithSelector(IHyperMintableERC20.CoreTokenIndexOutOfRange.selector, tooLarge));
        token.setCoreTokenIndex(tooLarge);
    }

    /// A8. `decimals() != weiDecimals + evmExtraWeiDecimals` is rejected.
    function test_A8_setCoreTokenIndex_decimalsMismatch_reverts() public {
        // 5 + 5 = 10 != DECIMALS (18).
        _mockCoreTokenInfo(9, address(token), 5, 5);

        vm.prank(tokenAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(IHyperMintableERC20.CoreDecimalsMismatch.selector, DECIMALS, uint8(5), int8(5))
        );
        token.setCoreTokenIndex(9);
    }

    /// A9. A second call — even with the exact same (already-linked) index — is rejected once
    /// the first has succeeded.
    function test_A9_setCoreTokenIndex_secondCall_reverts() public {
        _mockValidCoreLink(address(token), 11);
        vm.startPrank(tokenAdmin);
        token.setCoreTokenIndex(11);

        vm.expectRevert(IHyperMintableERC20.CoreTokenIndexAlreadySet.selector);
        token.setCoreTokenIndex(11);

        _mockValidCoreLink(address(token), 12);
        vm.expectRevert(IHyperMintableERC20.CoreTokenIndexAlreadySet.selector);
        token.setCoreTokenIndex(12);
        vm.stopPrank();
    }

    // =====================================================================
    // A10-A11. setHyperCoreDeployer lock
    // =====================================================================

    /// A10. Once `setCoreTokenIndex` has succeeded, `setHyperCoreDeployer` is permanently locked.
    function test_A10_setHyperCoreDeployer_lockedAfterIndexSet() public {
        _mockValidCoreLink(address(token), 13);
        vm.startPrank(tokenAdmin);
        token.setCoreTokenIndex(13);

        vm.expectRevert(IHyperMintableERC20.HyperCoreLinkAlreadyFinalized.selector);
        token.setHyperCoreDeployer(finalizer);
        vm.stopPrank();
    }

    /// A11. Before the index is set, `setHyperCoreDeployer` behaves exactly as before —
    /// including recalling a mistaken finalizer back to the zero address.
    function test_A11_setHyperCoreDeployer_worksBeforeIndexSet_includingRecallToZero() public {
        assertFalse(token.isCoreTokenIndexSet());

        vm.startPrank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);
        assertEq(token.hyperCoreDeployer(), finalizer);

        token.setHyperCoreDeployer(address(0));
        assertEq(token.hyperCoreDeployer(), address(0));
        vm.stopPrank();

        assertFalse(token.isCoreTokenIndexSet());
    }

    // =====================================================================
    // A12-A17. Rounded transferToCore / transferToCoreFor
    // =====================================================================

    /// A12. `transferToCore` rounds `amount` down to the nearest multiple of `coreUnit()` and
    /// returns the actually-sent amount; the remainder stays with the caller.
    function test_A12_transferToCore_roundsDown() public {
        _mockValidCoreLink(address(token), 14); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(tokenAdmin);
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

    /// A17. When `evmExtraWeiDecimals <= 0`, `coreUnit()` is `1` and nothing is ever rounded.
    function test_A17_coreUnit_isOneWhenExtraNonPositive() public {
        _mockCoreTokenInfo(15, address(token), 18, 0);
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(15);

        assertEq(token.coreUnit(), 1);
        assertEq(token.coreTransferableAmount(1234567), 1234567);
    }

    /// A13/A14. `transferToCoreFor` emits `Transfer(coreRecipient -> systemAddress)` for the
    /// exact rounded amount, and `coreRecipient`'s EVM balance nets to exactly zero (it only
    /// ever passes through).
    function test_A13_A14_transferToCoreFor_eventAndZeroNetBalance() public {
        _mockValidCoreLink(address(token), 16);
        vm.prank(tokenAdmin);
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

    /// A15. A `sent == 0` (amount rounds to zero) is rejected.
    function test_A15_transferToCoreFor_zeroSent_reverts() public {
        _mockValidCoreLink(address(token), 17); // coreUnit() == 1e10
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(17);

        vm.prank(bridge);
        token.mint(user, 10 ether);

        vm.prank(user);
        vm.expectRevert(IHyperMintableERC20.CoreAmountBelowOneCoreWei.selector);
        token.transferToCoreFor(makeAddr("coreRecipient"), 1e10 - 1);
    }

    /// A16. Zero `coreRecipient` and `coreRecipient == coreSystemAddress()` are both rejected.
    function test_A16_transferToCoreFor_invalidRecipients_reverts() public {
        _mockValidCoreLink(address(token), 18);
        vm.prank(tokenAdmin);
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
    // A18-A21. Beacon / factory upgrades
    // =====================================================================

    /// A18. Beacon upgrade preserves every token's existing storage (balances, roles, link
    /// state) and applies the new logic (`version()`) immediately.
    function test_A18_beaconUpgrade_preservesState() public {
        vm.prank(bridge);
        token.mint(user, 10 ether);
        vm.prank(user);
        token.approve(finalizer, 5 ether);
        vm.prank(tokenAdmin);
        token.setHyperCoreDeployer(finalizer);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(beaconOwner);
        beacon.upgradeTo(address(newImpl));

        assertEq(token.balanceOf(user), 10 ether);
        assertEq(token.allowance(user, finalizer), 5 ether);
        assertEq(token.hyperCoreDeployer(), finalizer);
        assertEq(HyperMintableERC20Mock(address(token)).version(), 2);
    }

    /// A19. A single beacon upgrade is reflected by every token sharing that beacon at once.
    function test_A19_beaconUpgrade_affectsAllTokensAtOnce() public {
        vm.prank(bridge);
        address token2Address = code.createCrossMintableERC20(REMOTE_CHAIN_ID + 100, REMOTE_TOKEN, "T2", 6);
        HyperMintableERC20 token2 = HyperMintableERC20(token2Address);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(beaconOwner);
        beacon.upgradeTo(address(newImpl));

        assertEq(HyperMintableERC20Mock(address(token)).version(), 2);
        assertEq(HyperMintableERC20Mock(address(token2)).version(), 2);
    }

    /// A20. Factory (UUPS) upgrade preserves storage: `tokenAdmin`, `beacon`, roles, and the
    /// `isHyperMintableERC20` registry all survive.
    function test_A20_factoryUpgrade_preservesState() public {
        assertTrue(code.isHyperMintableERC20(address(token)));

        HyperMintableERC20CodeMock newImpl = new HyperMintableERC20CodeMock();
        vm.prank(factoryOwner);
        code.upgradeToAndCall(address(newImpl), bytes(""));

        assertEq(code.tokenAdmin(), tokenAdmin);
        assertEq(code.beacon(), address(beacon));
        assertTrue(code.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertTrue(code.hasRole(Const.BRIDGE_ROLE, bridge));
        assertTrue(code.isHyperMintableERC20(address(token)));
        assertEq(HyperMintableERC20CodeMock(address(code)).version(), 2);
    }

    /// A21. Upgrade authority is gated: only the beacon owner can upgrade the beacon, only the
    /// factory's `ADMIN_ROLE` can upgrade the factory.
    function test_A21_upgradeAuthority_gated() public {
        HyperMintableERC20Mock newTokenImpl = new HyperMintableERC20Mock();
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user));
        beacon.upgradeTo(address(newTokenImpl));

        HyperMintableERC20CodeMock newCodeImpl = new HyperMintableERC20CodeMock();
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.upgradeToAndCall(address(newCodeImpl), bytes(""));
    }

    // =====================================================================
    // A22-A22c. Initialization safety
    // =====================================================================

    /// A22. The token LOGIC contract cannot be initialized directly.
    function test_A22_tokenImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        tokenImplementation.initialize(tokenAdmin, bridge, address(code), "X", "X", 18);
    }

    /// A22b. The factory LOGIC contract cannot be initialized directly.
    function test_A22b_codeImplementation_directInitializeReverts() public {
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        codeImplementation.initialize(factoryOwner, tokenAdmin, bridge, address(beacon));
    }

    /// A22c. The factory PROXY is initialized atomically at construction (no separate
    /// initialize transaction is possible or needed), and cannot be initialized a second time.
    function test_A22c_factoryProxy_atomicInitialization() public {
        // `code` (from setUp) is already live with its constructor-time state — no second
        // transaction occurred between deployment and this assertion.
        assertEq(code.tokenAdmin(), tokenAdmin);
        assertEq(code.beacon(), address(beacon));
        assertTrue(code.hasRole(Const.ADMIN_ROLE, factoryOwner));

        vm.expectRevert(Initializable.InvalidInitialization.selector);
        code.initialize(factoryOwner, tokenAdmin, bridge, address(beacon));
    }

    // =====================================================================
    // A23-A26. Explicit name/symbol/minter creation path
    // =====================================================================

    /// A23. `createHyperMintableERC20` deploys with caller-chosen name/symbol and grants
    /// `MINTER_ROLE` to the explicit `minter` argument.
    function test_A23_createHyperMintableERC20_arbitraryNameAndMinter() public {
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

    /// A24. `createHyperMintableERC20` is gated to `ADMIN_ROLE`.
    function test_A24_createHyperMintableERC20_requiresAdminRole() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, user, Const.ADMIN_ROLE)
        );
        code.createHyperMintableERC20(REMOTE_CHAIN_ID + 201, REMOTE_TOKEN, "X", "X", 18, user);
    }

    /// A25. `computeTokenAddressWithName` prediction equals the actual deployment.
    function test_A25_computeTokenAddressWithName_matchesActual() public {
        uint remoteChainID5 = REMOTE_CHAIN_ID + 202;
        address explicitMinter = makeAddr("explicitMinter2");

        address predicted =
            code.computeTokenAddressWithName(remoteChainID5, REMOTE_TOKEN, "Name X", "SYMX", 7, explicitMinter);

        vm.prank(factoryOwner);
        address actual =
            code.createHyperMintableERC20(remoteChainID5, REMOTE_TOKEN, "Name X", "SYMX", 7, explicitMinter);

        assertEq(predicted, actual);
    }

    /// A26. CREATE2 address prediction survives a beacon logic upgrade (R-10): the initcode
    /// only ever embeds the beacon's own fixed address, never the implementation it currently
    /// resolves to.
    function test_A26_predictedAddress_unchangedAfterBeaconUpgrade() public {
        uint remoteChainID6 = REMOTE_CHAIN_ID + 203;

        address predictedBefore = code.computeTokenAddress(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);

        HyperMintableERC20Mock newImpl = new HyperMintableERC20Mock();
        vm.prank(beaconOwner);
        beacon.upgradeTo(address(newImpl));

        address predictedAfter = code.computeTokenAddress(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS, bridge);
        assertEq(predictedBefore, predictedAfter);

        vm.prank(bridge);
        address actual = code.createCrossMintableERC20(remoteChainID6, REMOTE_TOKEN, SYMBOL, DECIMALS);

        assertEq(actual, predictedBefore);
        assertEq(HyperMintableERC20Mock(actual).version(), 2);
    }

    // =====================================================================
    // A28. End-to-end `BridgeExecutor` extra-call path for `transferToCoreFor`
    // =====================================================================

    /// A28. `transferToCoreFor` exercised through the REAL `BridgeExecutor.executeExtraCall`
    /// path (not a direct call), per AC-7: the extra-call's whole point is that Core credits
    /// the USER, never the executor. `value` is deliberately not a multiple of `coreUnit()` so
    /// a genuine truncation occurs; `to` is set to the executor itself so the truncated dust
    /// that `executeExtraCall` returns to `to` is a self-transfer (stays with the executor),
    /// keeping the executor's own balance change equal to exactly `sent` and leaving `user`'s
    /// balance untouched by anything except the pass-through (which nets to zero).
    function test_A28_transferToCoreFor_viaBridgeExecutor_creditsUserNotExecutor() public {
        _mockValidCoreLink(address(token), 28); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(tokenAdmin);
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

    /// A29. The PRODUCTION shape of A28: `to` is the USER, which is what the bridge actually
    /// passes. A28 deliberately set `to` to the executor so the dust self-transfer could not
    /// disturb its "user nets to zero" assertion — that left one thing unproven, namely that
    /// the truncated remainder is genuinely handed back to the user rather than stranded in
    /// the executor. This case closes exactly that gap: the pass-through still credits Core to
    /// the user (`Transfer(user -> systemAddr)`), and on top of it `executeExtraCall` returns
    /// `remaining == value - consumed` to `to`, so the user ends up **up by the dust** and the
    /// executor ends up with nothing.
    function test_A29_transferToCoreFor_viaBridgeExecutor_dustReturnedToUser() public {
        _mockValidCoreLink(address(token), 29); // weiDecimals 8, evmExtraWeiDecimals 10 -> coreUnit() == 1e10
        vm.prank(tokenAdmin);
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

        // Core still credits the USER, exactly as in A28 — `to` does not affect that leg.
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
}

/// @dev Minimal "upgraded" token implementation used only by beacon-upgrade tests (A18, A19,
/// A26): identical storage layout and behavior to `HyperMintableERC20`, plus one new function
/// so the tests can observe that an upgrade actually took effect.
contract HyperMintableERC20Mock is HyperMintableERC20 {
    function version() external pure returns (uint) {
        return 2;
    }
}

/// @dev Minimal "upgraded" factory implementation used only by the factory-upgrade test (A20).
contract HyperMintableERC20CodeMock is HyperMintableERC20Code {
    function version() external pure returns (uint) {
        return 2;
    }
}
