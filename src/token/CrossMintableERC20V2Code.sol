// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {CrossMintableERC20V2} from "./CrossMintableERC20V2.sol";
import {ICrossMintableERC20Code} from "./ICrossMintableERC20Code.sol";
import {ICrossMintableERC20V2Code} from "./ICrossMintableERC20V2Code.sol";

import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

/**
 * @title CrossMintableERC20V2Code
 * @notice Factory that deploys `CrossMintableERC20V2` tokens via CREATE2, as `BeaconProxy`
 *         instances all sharing one `UpgradeableBeacon` that THIS factory creates and owns
 *         itself in `initialize` (D8) — a single beacon upgrade fixes every token this factory
 *         created at once, instead of the costly new-token + pair-remap + liquidity-move
 *         migration the old non-upgradeable design required (already paid once on HyperEVM).
 * @dev Every created token's `defaultAdmin()` is this factory itself (`address(this)`), never a
 *      separately stored `tokenAdmin` (D7) — the factory's own address is already fixed (it is
 *      the CREATE2 deployer of every token it creates), so there is nothing new to keep in sync
 *      when this factory's own ownership changes hands. This is also why this factory, unlike
 *      `HyperMintableERC20Code`, needs no `tokenAdmin` storage slot at all.
 *
 *      A consequence of the above: trust concentrates entirely on this factory's `ADMIN_ROLE`
 *      (R9). That one role can (1) upgrade the shared beacon, rewriting every created token's
 *      logic at once, (2) manage any created token's roles through the pass-through setters
 *      below (including minting new supply out of thin air by granting `MINTER_ROLE`), and
 *      (3) upgrade this factory itself (UUPS). Splitting these three into separate principals is
 *      not possible by construction (a single `defaultAdmin()`/beacon owner is the point, D9) —
 *      the only available defense is making that one `ADMIN_ROLE` holder a multisig/timelock
 *      (see the deploy script). `beginTokenDefaultAdminTransfer` is an escape hatch, but only
 *      within token-level administration: it moves a token's `defaultAdmin()`/role-management
 *      authority to a new admin. It does NOT detach the token from this factory's shared beacon —
 *      the token remains a `BeaconProxy` permanently pointed at `beacon()`, so this factory's
 *      `ADMIN_ROLE` holder can still replace that token's logic (along with every other token this
 *      factory created) via `upgradeBeacon`. See `ICrossMintableERC20V2Code.beginTokenDefaultAdminTransfer`.
 *
 *      This factory is itself deployed as a UUPS (`ERC1967Proxy`) singleton — a single instance
 *      is expected, so there is no need for a beacon here (mirrors this factory's
 *      `HyperMintableERC20Code` sibling).
 *
 *      One token per remote pair (R10): CREATE2's salt is derived only from
 *      `(remoteChainID, remoteToken)`, but the initcode built by `_initCode` also embeds
 *      name/symbol/decimals/minter — so salt collision alone cannot stop a second, differently
 *      named token for the same pair from landing at a different address. `tokenForPair` closes
 *      that gap explicitly, shared by both creation paths so crossing them cannot bypass it
 *      either. This guard is scoped to tokens THIS factory created; a duplicate against a token
 *      a V1 factory already created is a separate, operational-only concern (deploy script
 *      preflight, see `script/CrossMintableERC20V2Code.s.sol`).
 */
contract CrossMintableERC20V2Code is
    UUPSUpgradeable,
    AccessControlDefaultAdminRulesUpgradeable,
    ICrossMintableERC20V2Code
{
    // `CrossMintableERC20V2CodeZeroAddress` / `CrossMintableERC20V2CodeUnknownToken` /
    // `CrossMintableERC20V2CodePairAlreadyCreated` are declared on `ICrossMintableERC20V2Code`
    // (U1) and inherited from there — not redeclared here.

    /// @custom:storage-location erc7201:nexus.storage.CrossMintableERC20V2Code
    struct CrossMintableERC20V2CodeStorage {
        // `UpgradeableBeacon` this factory creates and takes ownership of inside `initialize`
        // (D8) — there is no separate beacon deployment / ownership hand-off step. Only its
        // address is embedded in a token's initcode, never the implementation it currently
        // resolves to, so an upgrade of the beacon's implementation never moves future CREATE2
        // predictions.
        address beacon;
        // Tokens this factory has deployed, gating the token-scoped pass-through setters below.
        mapping(address => bool) isCrossMintableERC20;
        // keccak256(abi.encode(remoteChainID, remoteToken)) -> token this factory created for
        // that pair. Enforces at most one token per remote pair across BOTH creation paths (R10)
        // — see `_create`.
        mapping(bytes32 => address) tokenForPair;
    }

    // keccak256(abi.encode(uint256(keccak256("nexus.storage.CrossMintableERC20V2Code")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CrossMintableERC20V2CodeStorageLocation =
        0x2b6c7dfd4e9330f706f3d4045fa01a8903ed43939ff7d252d6806f39e1487c00;

    function _getCrossMintableERC20V2CodeStorage() private pure returns (CrossMintableERC20V2CodeStorage storage $) {
        assembly {
            $.slot := CrossMintableERC20V2CodeStorageLocation
        }
    }

    /// @dev Implementation contract is never initialized directly — only the `ERC1967Proxy`
    /// pointing at it is (see deploy scripts).
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the factory: grants roles, then creates and takes ownership of the
     *         shared token beacon (D8) — no separate beacon deployment / hand-off step exists.
     * @param initialOwner Granted `Const.ADMIN_ROLE` (and, via `AccessControlDefaultAdminRules`,
     *        `DEFAULT_ADMIN_ROLE`) on this factory
     * @param initialBridge Granted `Const.BRIDGE_ROLE` if non-zero
     * @param tokenImplementation Initial `CrossMintableERC20V2` logic address the newly created
     *        `UpgradeableBeacon` (owned by this factory) will point at; must be non-zero
     */
    function initialize(address initialOwner, address initialBridge, address tokenImplementation) public initializer {
        require(tokenImplementation != address(0), CrossMintableERC20V2CodeZeroAddress());

        __UUPSUpgradeable_init();
        __AccessControlDefaultAdminRules_init(0, initialOwner);

        // The beacon this factory will forever create tokens against (R8) — owned by this
        // factory itself, so `upgradeBeacon` below is the only way to change its implementation.
        _getCrossMintableERC20V2CodeStorage().beacon = address(new UpgradeableBeacon(tokenImplementation, address(this)));

        // `_grantRole`, not the public `grantRole` — the latter is what makes the legacy
        // (non-upgradeable) `CrossMintableERC20V2Code`'s constructor revert whenever the deployer
        // account isn't `initialOwner` itself (public `grantRole` is access-controlled and checks
        // the caller, not just `initialOwner`). `_grantRole` bypasses that check entirely, as
        // intended here (the account broadcasting the proxy deployment need not be
        // `initialOwner`) (R6).
        _grantRole(Const.ADMIN_ROLE, initialOwner);

        if (initialBridge != address(0)) _grantRole(Const.BRIDGE_ROLE, initialBridge);
    }

    function _authorizeUpgrade(address) internal override onlyRole(Const.ADMIN_ROLE) {}

    /// @dev Gates the token-scoped owner pass-through setters to tokens this factory created —
    /// mirrors `HyperMintableERC20Code`'s `HyperMintableERC20CodeUnknownToken` gate.
    modifier onlyKnownToken(address token) {
        require(
            _getCrossMintableERC20V2CodeStorage().isCrossMintableERC20[token], CrossMintableERC20V2CodeUnknownToken(token)
        );
        _;
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function beacon() external view returns (address) {
        return _getCrossMintableERC20V2CodeStorage().beacon;
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function isCrossMintableERC20(address token) external view returns (bool) {
        return _getCrossMintableERC20V2CodeStorage().isCrossMintableERC20[token];
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function tokenForPair(uint remoteChainID, address remoteToken) external view returns (address) {
        return _getCrossMintableERC20V2CodeStorage().tokenForPair[_pairKey(remoteChainID, remoteToken)];
    }

    /// @dev Single definition of the pair key so `_create`'s guard and `tokenForPair`'s getter
    /// can never disagree on how a pair is identified.
    function _pairKey(uint remoteChainID, address remoteToken) internal pure returns (bytes32) {
        return keccak256(abi.encode(remoteChainID, remoteToken));
    }

    /**
     * @dev Shared by both creation paths and both address-prediction views so the predicted and
     *      actual initcode can never structurally diverge (R8). Builds a
     *      `BeaconProxy(beacon, initData)` initcode: only the beacon's own (fixed) address is
     *      embedded, never the implementation it currently points at, and the token's initial
     *      owner is always `address(this)` — this factory's own (CREATE2/proxy-fixed) address —
     *      never a mutable value like `defaultAdmin()`. So neither a beacon upgrade nor a change
     *      of who currently holds this factory's `ADMIN_ROLE`/`defaultAdmin()` ever moves a
     *      future CREATE2 prediction (AC-9).
     */
    function _initCode(string memory name_, string memory symbol_, uint8 decimals_, address minter)
        internal
        view
        returns (bytes memory)
    {
        bytes memory initData =
            abi.encodeCall(CrossMintableERC20V2.initialize, (address(this), minter, name_, symbol_, decimals_));
        return abi.encodePacked(
            type(BeaconProxy).creationCode, abi.encode(_getCrossMintableERC20V2CodeStorage().beacon, initData)
        );
    }

    /**
     * @dev Shared CREATE2 deploy + pair guard + registry-mark + event for both creation paths
     *      (R10). `_initCode` bakes name/symbol/decimals/minter into the CREATE2 initcode, so the
     *      salt (derived only from `remoteChainID`/`remoteToken`) alone cannot stop a second,
     *      differently-named/minter'd token for the same pair from landing at a DIFFERENT
     *      address — hence the explicit `tokenForPair` check below, keyed identically regardless
     *      of which creation path is used, so crossing paths cannot bypass it either. This is an
     *      intentional behavior change from the old factory (which allowed exactly that by
     *      varying only the symbol) — a bug fix, since more than one token per pair fragments
     *      liquidity/accounting.
     */
    function _create(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        address minter
    ) internal returns (address tokenAddress) {
        CrossMintableERC20V2CodeStorage storage $ = _getCrossMintableERC20V2CodeStorage();
        bytes32 pairKey = _pairKey(remoteChainID, remoteToken);
        require(
            $.tokenForPair[pairKey] == address(0),
            CrossMintableERC20V2CodePairAlreadyCreated(remoteChainID, remoteToken, $.tokenForPair[pairKey])
        );

        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        tokenAddress = Create2.deploy(0, salt, _initCode(name_, symbol_, decimals_, minter));

        $.isCrossMintableERC20[tokenAddress] = true;
        $.tokenForPair[pairKey] = tokenAddress;
        emit CrossMintableERC20Created(remoteChainID, remoteToken, tokenAddress);
    }

    /// @inheritdoc ICrossMintableERC20Code
    function createCrossMintableERC20(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        onlyRole(Const.BRIDGE_ROLE)
        returns (address tokenAddress)
    {
        tokenAddress = _create(
            remoteChainID,
            remoteToken,
            string(abi.encodePacked("Cross Bridge ", symbol)),
            string(abi.encodePacked(symbol, "x")),
            decimals,
            _msgSender()
        );
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function createMintableERC20(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external onlyRole(Const.ADMIN_ROLE) returns (address tokenAddress) {
        tokenAddress = _create(remoteChainID, remoteToken, name_, symbol_, decimals, minter);
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function computeTokenAddress(
        uint remoteChainID,
        address remoteToken,
        string memory symbol,
        uint8 decimals,
        address minter
    ) external view returns (address) {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        return Create2.computeAddress(
            salt,
            keccak256(
                _initCode(
                    string(abi.encodePacked("Cross Bridge ", symbol)),
                    string(abi.encodePacked(symbol, "x")),
                    decimals,
                    minter
                )
            ),
            address(this)
        );
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function computeTokenAddressWithName(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external view returns (address) {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        return Create2.computeAddress(salt, keccak256(_initCode(name_, symbol_, decimals, minter)), address(this));
    }

    // =====================================================================
    // Owner-managed pass-through surface (R9) — all `ADMIN_ROLE`; the three
    // that take a `token` argument are additionally gated to tokens this
    // factory created.
    // =====================================================================

    /// @inheritdoc ICrossMintableERC20V2Code
    function upgradeBeacon(address newImplementation) external onlyRole(Const.ADMIN_ROLE) {
        UpgradeableBeacon(_getCrossMintableERC20V2CodeStorage().beacon).upgradeTo(newImplementation);
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function grantTokenRole(address token, bytes32 role, address account)
        external
        onlyRole(Const.ADMIN_ROLE)
        onlyKnownToken(token)
    {
        // `this` (the factory) is `token`'s `defaultAdmin()` (D7), so this call succeeds exactly
        // as if `defaultAdmin()` had called `grantRole` on the token directly.
        CrossMintableERC20V2(token).grantRole(role, account);
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function revokeTokenRole(address token, bytes32 role, address account)
        external
        onlyRole(Const.ADMIN_ROLE)
        onlyKnownToken(token)
    {
        CrossMintableERC20V2(token).revokeRole(role, account);
    }

    /// @inheritdoc ICrossMintableERC20V2Code
    function beginTokenDefaultAdminTransfer(address token, address newAdmin)
        external
        onlyRole(Const.ADMIN_ROLE)
        onlyKnownToken(token)
    {
        CrossMintableERC20V2(token).beginDefaultAdminTransfer(newAdmin);
    }
}
