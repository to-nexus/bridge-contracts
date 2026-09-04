// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {HyperMintableERC20} from "./HyperMintableERC20.sol";
import {IHyperMintableERC20} from "./IHyperMintableERC20.sol";
import {IHyperMintableERC20Code} from "./IHyperMintableERC20Code.sol";
import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

/**
 * @title HyperMintableERC20Code
 * @notice Factory that deploys `HyperMintableERC20` tokens via CREATE2 (as `BeaconProxy`
 *         instances sharing one `UpgradeableBeacon`) and, unlike `CrossMintableERC20V2Code`,
 *         keeps the initcode deterministic enough to predict a token's address before it is
 *         created — even across a beacon logic upgrade, since the initcode only ever embeds the
 *         beacon's own (fixed) address, never the implementation it currently points at.
 * @dev Determinism rests on every initcode input being either fixed storage or an explicit
 *      argument: `tokenAdmin`/`beacon` (storage here, see `tokenAdmin`/`beacon` docs), `minter`
 *      (explicit argument to `computeTokenAddress`/`computeTokenAddressWithName`, `_msgSender()`
 *      in the real call), `address(this)`, `name`/`symbol` and `decimals`.
 *      `CrossMintableERC20V2Code` instead bakes in the mutable `defaultAdmin()` and the
 *      per-caller `_msgSender()`, which is why it cannot offer address prediction.
 *
 *      This factory is itself deployed as a UUPS (`ERC1967Proxy`) singleton — a single instance
 *      is expected, so there is no need for a beacon here.
 */
contract HyperMintableERC20Code is
    UUPSUpgradeable,
    AccessControlDefaultAdminRulesUpgradeable,
    IHyperMintableERC20Code
{
    error HyperMintableERC20CodeZeroAddress();
    error HyperMintableERC20CodeUnknownToken(address token);

    /// @custom:storage-location erc7201:nexus.storage.HyperMintableERC20Code
    struct HyperMintableERC20CodeStorage {
        // Becomes `defaultAdmin()` of every token this factory creates. As the token's
        // *current* default admin it has effective link authority (`isLinkAuthority`) over
        // `setHyperCoreDeployer` / `setCoreTokenIndex` dynamically — it is never granted
        // `LINKER_ROLE` itself, and does not need to be: authority tracks whoever currently
        // holds `defaultAdmin()` on the token, following that token's own
        // `beginDefaultAdminTransfer` / `acceptDefaultAdminTransfer` automatically. This
        // factory, by contrast, IS granted `LINKER_ROLE` on each token it creates (see
        // `_initCode`'s `address(this)` argument) and is that token's `factoryLinker` — the
        // token's admin can `revokeRole(LINKER_ROLE, factory)` to cut this factory's access
        // off, and `grantRole` it back later to restore it; a `LINKER_ROLE` grant to any other
        // address confers no authority (see `HyperMintableERC20.isLinkAuthority`).
        // Fixed at `initialize` — never changes out from under a CREATE2 address prediction.
        // Use a multisig or timelock address here (see `script/HyperMintableERC20Code.s.sol`).
        address tokenAdmin;
        // `UpgradeableBeacon` every created token's `BeaconProxy` points at. Only its address is
        // embedded in a token's initcode, never the implementation it currently resolves to, so
        // an upgrade of the beacon's implementation never moves future CREATE2 predictions.
        address beacon;
        // Tokens this factory has deployed, gating the setter delegation below.
        mapping(address => bool) isHyperMintableERC20;
    }

    // keccak256(abi.encode(uint256(keccak256("nexus.storage.HyperMintableERC20Code")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant HyperMintableERC20CodeStorageLocation =
        0xd55591ae453b55973909532d0252f41e2a7c543bd5a578ad3871068e0c28b400;

    function _getHyperMintableERC20CodeStorage() private pure returns (HyperMintableERC20CodeStorage storage $) {
        assembly {
            $.slot := HyperMintableERC20CodeStorageLocation
        }
    }

    /// @dev Implementation contract is never initialized directly — only the `ERC1967Proxy`
    /// pointing at it is (see deploy scripts).
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the factory. Called once, atomically, from the `ERC1967Proxy`
     *         constructor.
     * @param initialOwner Granted `Const.ADMIN_ROLE` (and, via `AccessControlDefaultAdminRules`,
     *        `DEFAULT_ADMIN_ROLE`) on this factory
     * @param initialTokenAdmin `defaultAdmin()` passed to every created token (effective link
     *        authority via `isLinkAuthority`, without holding `LINKER_ROLE` itself); must be non-zero
     * @param initialBridge Granted `Const.BRIDGE_ROLE` if non-zero
     * @param beacon_ `UpgradeableBeacon` every created token's `BeaconProxy` will point at; must
     *        be non-zero
     */
    function initialize(address initialOwner, address initialTokenAdmin, address initialBridge, address beacon_)
        public
        initializer
    {
        require(initialTokenAdmin != address(0), HyperMintableERC20CodeZeroAddress());
        require(beacon_ != address(0), HyperMintableERC20CodeZeroAddress());

        __UUPSUpgradeable_init();
        __AccessControlDefaultAdminRules_init(0, initialOwner);

        HyperMintableERC20CodeStorage storage $ = _getHyperMintableERC20CodeStorage();
        $.tokenAdmin = initialTokenAdmin;
        $.beacon = beacon_;

        // `_grantRole`, not the public `grantRole` — the latter is what makes
        // `CrossMintableERC20V2Code`'s constructor revert whenever the deployer account isn't
        // `initialOwner` itself (public `grantRole` is access-controlled and checks the caller,
        // not just `initialOwner`). `_grantRole` bypasses that check entirely, as intended here
        // (the account broadcasting the proxy deployment need not be `initialOwner`).
        _grantRole(Const.ADMIN_ROLE, initialOwner);

        if (initialBridge != address(0)) _grantRole(Const.BRIDGE_ROLE, initialBridge);
    }

    function _authorizeUpgrade(address) internal override onlyRole(Const.ADMIN_ROLE) {}

    /// @inheritdoc IHyperMintableERC20Code
    function tokenAdmin() external view returns (address) {
        return _getHyperMintableERC20CodeStorage().tokenAdmin;
    }

    /// @inheritdoc IHyperMintableERC20Code
    function beacon() external view returns (address) {
        return _getHyperMintableERC20CodeStorage().beacon;
    }

    /// @inheritdoc IHyperMintableERC20Code
    function isHyperMintableERC20(address token) external view returns (bool) {
        return _getHyperMintableERC20CodeStorage().isHyperMintableERC20[token];
    }

    /**
     * @dev Shared by both creation paths and both address-prediction views so the predicted and
     *      actual initcode can never structurally diverge. Builds a `BeaconProxy(beacon, initData)`
     *      initcode — the logic implementation address never appears in it (R-10): only the
     *      beacon's own (fixed) address does, so a beacon upgrade never moves future predictions.
     */
    function _initCode(string memory name_, string memory symbol_, uint8 decimals_, address minter)
        internal
        view
        returns (bytes memory)
    {
        HyperMintableERC20CodeStorage storage $ = _getHyperMintableERC20CodeStorage();
        bytes memory initData = abi.encodeCall(
            HyperMintableERC20.initialize,
            (
                $.tokenAdmin, // Initial owner -> becomes defaultAdmin(); effective link authority via isLinkAuthority
                minter, // Initial minter
                address(this), // Initial linker -> this factory; granted LINKER_ROLE and stored as factoryLinker
                name_,
                symbol_,
                decimals_
            )
        );
        return abi.encodePacked(type(BeaconProxy).creationCode, abi.encode($.beacon, initData));
    }

    /// @dev Shared CREATE2 deploy + registry-mark + event for both creation paths.
    function _create(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        address minter
    ) internal returns (address tokenAddress) {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        tokenAddress = Create2.deploy(0, salt, _initCode(name_, symbol_, decimals_, minter));

        _getHyperMintableERC20CodeStorage().isHyperMintableERC20[tokenAddress] = true;
        emit HyperMintableERC20Created(remoteChainID, remoteToken, tokenAddress);
    }

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

    /// @inheritdoc IHyperMintableERC20Code
    function createHyperMintableERC20(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external onlyRole(Const.ADMIN_ROLE) returns (address tokenAddress) {
        tokenAddress = _create(remoteChainID, remoteToken, name_, symbol_, decimals, minter);
    }

    /// @inheritdoc IHyperMintableERC20Code
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

    /// @inheritdoc IHyperMintableERC20Code
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

    /// @inheritdoc IHyperMintableERC20Code
    function setHyperCoreDeployer(address token, address finalizer) external onlyRole(Const.ADMIN_ROLE) {
        require(
            _getHyperMintableERC20CodeStorage().isHyperMintableERC20[token], HyperMintableERC20CodeUnknownToken(token)
        );
        IHyperMintableERC20(token).setHyperCoreDeployer(finalizer);
    }

    /// @inheritdoc IHyperMintableERC20Code
    function setCoreTokenIndex(address token, uint64 index) external onlyRole(Const.ADMIN_ROLE) {
        require(
            _getHyperMintableERC20CodeStorage().isHyperMintableERC20[token], HyperMintableERC20CodeUnknownToken(token)
        );
        IHyperMintableERC20(token).setCoreTokenIndex(index);
    }
}
