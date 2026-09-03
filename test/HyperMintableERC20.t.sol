// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../src/lib/Const.sol";
import {HyperMintableERC20} from "../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {IERC20Errors} from "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";

import {Errors} from "@openzeppelin/contracts/utils/Errors.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/Test.sol";

/**
 * @title HyperMintableERC20Test
 * @notice T1-T18 of the HyperMintableERC20 / HyperMintableERC20Code test plan (spec §14).
 * @dev `test/CrossBridgeV2HyperEVMRoute.t.sol` covers T19-T22, the end-to-end path through a
 *      real bridge instance.
 */
contract HyperMintableERC20Test is Test {
    /// @dev Declared locally (not imported from IERC20) purely so `vm.expectEmit` can match
    /// the inherited OZ ERC20 `Transfer` event by signature.
    event Transfer(address indexed from, address indexed to, uint value);

    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

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

        code = new HyperMintableERC20Code(factoryOwner, tokenAdmin, bridge);
        console.log("HyperMintableERC20Code", address(code));

        vm.prank(bridge);
        token = HyperMintableERC20(code.createCrossMintableERC20(REMOTE_CHAIN_ID, REMOTE_TOKEN, SYMBOL, DECIMALS));
        console.log("HyperMintableERC20", address(token));
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

    /// T8. Regression: deployer != initialOwner must not revert the factory constructor
    /// (the trap the V2 factory falls into by using public `grantRole` instead of `_grantRole`).
    function test_T8_deployerNotInitialOwner_doesNotRevert() public {
        address deployer = makeAddr("someoneElse");
        vm.prank(deployer);
        HyperMintableERC20Code freshCode = new HyperMintableERC20Code(factoryOwner, tokenAdmin, bridge);

        assertTrue(freshCode.hasRole(Const.ADMIN_ROLE, factoryOwner));
        assertFalse(freshCode.hasRole(Const.ADMIN_ROLE, deployer));
        assertEq(freshCode.tokenAdmin(), tokenAdmin);
    }

    // ---------------------------------------------------------------------
    // Core system address
    // ---------------------------------------------------------------------

    /// T9. index 200 -> 0x20000000000000000000000000000000000000c8 exactly.
    function test_T9_coreSystemAddress_index200() public {
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(200);

        address expected = 0x20000000000000000000000000000000000000C8;
        assertEq(token.coreSystemAddress(), expected);
    }

    /// T10. index 0 is settable and valid (not an "unset" sentinel).
    function test_T10_coreTokenIndexZero_isValid() public {
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
        vm.prank(tokenAdmin);
        token.setCoreTokenIndex(200);
        address systemAddr = token.coreSystemAddress();

        vm.prank(bridge);
        token.mint(user, 10 ether);

        vm.expectEmit(true, true, false, true, address(token));
        emit Transfer(user, systemAddr, 5 ether);

        vm.prank(user);
        assertTrue(token.transferToCore(5 ether));

        assertEq(token.balanceOf(systemAddr), 5 ether);
        assertEq(token.balanceOf(user), 5 ether);

        uint balance = token.balanceOf(user);
        uint attempted = balance + 1; // attempted amount and the `needed` error argument must match
        vm.expectRevert(
            abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, user, balance, attempted)
        );
        vm.prank(user);
        token.transferToCore(attempted);
    }

    /// T13. Fuzz: for any uint64 index, the high byte is always 0x20 and the low 19 bytes are the index.
    function testFuzz_T13_coreSystemAddress(uint64 index) public {
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
        HyperMintableERC20Code codeB = new HyperMintableERC20Code(factoryBOwner, tokenAdmin, bridge);

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
    /// pinned to the immutable `tokenAdmin`.
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

    /// T18. EIP-170: runtime code of the token and the factory both stay under the 24,576-byte limit.
    /// @dev `type(HyperMintableERC20).runtimeCode` cannot be used here: solc rejects
    /// `runtimeCode`/`creationCode` for any contract that (like this one, via `_decimals`)
    /// contains an immutable variable, since the generic runtime code has no fixed value for
    /// where the immutable gets inlined. `address(token).code.length` is the actual deployed
    /// runtime bytecode size and is what EIP-170 checks against, so it covers the same ground.
    function test_T18_runtimeCodeSize_underEip170Limit() public view {
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
    /// can actually write the raw slot.
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
    /// `factoryLinker` is immutable; re-granting it restores the factory path.
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
}
