// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";
import {IHyperMintableERC20} from "./IHyperMintableERC20.sol";

import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PermitUpgradeable} from
    "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {StorageSlot} from "@openzeppelin/contracts/utils/StorageSlot.sol";

/**
 * @title HyperMintableERC20
 * @notice `CrossMintableERC20V2`-shaped wrapped token with a hardened HyperCore link on top,
 *         deployed as a `BeaconProxy` so its logic can be upgraded without changing its address
 *         (a HyperCore link is permanently bound to the EVM address that Core finalized).
 * @dev Deploying a bridge-wrapped token on HyperEVM with the ordinary `CrossMintableERC20V2Code`
 *      factory leaves the token with no way to link to a HyperCore spot asset: of the three
 *      `finalizeEvmContract` variants Core offers, `create{nonce}` requires an EOA CREATE
 *      deployment (the factory deploys via CREATE2), `firstStorageSlot` needs the finalizer at
 *      storage slot 0 (occupied by OZ `ERC20Upgradeable`'s namespaced balances), and only
 *      `customStorageSlot` — which reads the finalizer from `keccak256("HyperCore deployer")` —
 *      is usable. This contract exposes that slot through `setHyperCoreDeployer` /
 *      `hyperCoreDeployer`.
 *
 *      `HYPERCORE_DEPLOYER_SLOT` is a plain `keccak256("HyperCore deployer")` digest, exactly as
 *      Hyperliquid's documentation specifies it — it is NOT ERC-7201 masked (`-1 & ~0xff`).
 *      Because it is a keccak256 output, collision with Solidity's sequential slots (0, 1, 2, ...)
 *      or mapping/array derived slots is negligible, the same argument ERC-7201 itself relies on.
 *
 *      Deployed behind a `BeaconProxy`: every token created by a given `HyperMintableERC20Code`
 *      factory shares one `UpgradeableBeacon`, so a single beacon upgrade fixes a bug in every
 *      token at once (Beacon's own single point of concentration — see the deploy scripts for
 *      the recommended multisig/timelock beacon owner). Own state lives in ERC-7201 namespaced
 *      storage so it never collides with `ERC20Upgradeable` / `AccessControlUpgradeable` slots
 *      or with the ERC-1967 beacon/implementation slots the proxy itself uses.
 */
contract HyperMintableERC20 is
    ERC20Upgradeable,
    ERC20PermitUpgradeable,
    IHyperMintableERC20,
    AccessControlDefaultAdminRulesUpgradeable
{
    error HyperMintableERC20CoreTokenIndexNotSet();

    /// @dev keccak256("HyperCore deployer") — the ONLY slot Core's `customStorageSlot` variant
    /// of `finalizeEvmContract` reads. Plain digest, not ERC-7201 masked.
    bytes32 public constant HYPERCORE_DEPLOYER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

    /// @dev Top byte of every HyperCore system address (`0x20...` + 19-byte token index).
    uint160 public constant CORE_SYSTEM_ADDRESS_PREFIX = uint160(0x20) << 152;

    /// @dev HyperEVM's Core read precompile: `tokenInfo(uint32) -> CoreTokenInfo`. Used only by
    /// `setCoreTokenIndex` to obtain a forgery-proof, on-chain proof of the link (§6.4).
    address internal constant CORE_TOKEN_INFO_PRECOMPILE = 0x000000000000000000000000000000000000080C;

    /// @custom:storage-location erc7201:nexus.storage.HyperMintableERC20
    struct HyperMintableERC20Storage {
        // Packed into a single slot: 8 + 1 + 1 + 1 + 20 = 31 bytes.
        uint64 coreTokenIndex;
        bool coreTokenIndexSet;
        int8 coreExtraWeiDecimals;
        uint8 decimals;
        address factoryLinker;
    }

    // keccak256(abi.encode(uint256(keccak256("nexus.storage.HyperMintableERC20")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant HyperMintableERC20StorageLocation =
        0xc65afee183f44952959f664e5c2b8a5e4916480c324e2787b2215d744391d400;

    function _getHyperMintableERC20Storage() private pure returns (HyperMintableERC20Storage storage $) {
        assembly {
            $.slot := HyperMintableERC20StorageLocation
        }
    }

    /// @dev Implementation contract is never initialized directly — only `BeaconProxy` instances
    /// pointing at it are (see `HyperMintableERC20Code._initCode`).
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes a token instance. Called once, atomically, from the `BeaconProxy`
     *         constructor via `HyperMintableERC20Code._initCode`.
     * @param initialOwner Default admin of the token; always has link authority regardless of role
     * @param initialMinter Address granted `MINTER_ROLE` if non-zero (normally the bridge)
     * @param initialLinker Address granted `LINKER_ROLE` if non-zero (normally the creating factory);
     *        stored as `factoryLinker`
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals_ ERC20 decimals
     */
    function initialize(
        address initialOwner,
        address initialMinter,
        address initialLinker,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) public initializer {
        __ERC20_init(name_, symbol_);
        __ERC20Permit_init(name_);
        __AccessControlDefaultAdminRules_init(0, initialOwner);

        if (initialMinter != address(0)) {
            // Grant the minter role to the bridge
            _grantRole(Const.MINTER_ROLE, initialMinter);
        }

        HyperMintableERC20Storage storage $ = _getHyperMintableERC20Storage();
        $.factoryLinker = initialLinker;
        if (initialLinker != address(0)) {
            // Grant the linker role to the creating factory
            _grantRole(Const.LINKER_ROLE, initialLinker);
        }

        $.decimals = decimals_;
    }

    /// @dev Reverts with the same ABI `onlyRole(Const.LINKER_ROLE)` would (`IAccessControl`'s
    /// `AccessControlUnauthorizedAccount`), but gates on `isLinkAuthority` instead of raw role
    /// membership so a `LINKER_ROLE` grant outside the two authorized principals confers nothing.
    modifier onlyLinkAuthority() {
        if (!isLinkAuthority(_msgSender())) {
            revert IAccessControl.AccessControlUnauthorizedAccount(_msgSender(), Const.LINKER_ROLE);
        }
        _;
    }

    /// @inheritdoc IHyperMintableERC20
    function isLinkAuthority(address account) public view returns (bool) {
        // Zero is never an authority — after default-admin renunciation `defaultAdmin()`
        // itself becomes zero, and `factoryLinker` is zero for a standalone deployment.
        if (account == address(0)) return false;
        if (account == defaultAdmin()) return true;
        return account == _getHyperMintableERC20Storage().factoryLinker && hasRole(Const.LINKER_ROLE, account);
    }

    /// @inheritdoc IHyperMintableERC20
    function factoryLinker() external view returns (address) {
        return _getHyperMintableERC20Storage().factoryLinker;
    }

    function mint(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    function decimals() public view override(ERC20Upgradeable, ICrossMintableERC20) returns (uint8) {
        return _getHyperMintableERC20Storage().decimals;
    }

    function nonces(address owner_) public view override(ERC20PermitUpgradeable, ICrossMintableERC20) returns (uint) {
        return super.nonces(owner_);
    }

    /// @inheritdoc IHyperMintableERC20
    function setHyperCoreDeployer(address finalizer) external onlyLinkAuthority {
        require(!_getHyperMintableERC20Storage().coreTokenIndexSet, HyperCoreLinkAlreadyFinalized());

        StorageSlot.AddressSlot storage slot = StorageSlot.getAddressSlot(HYPERCORE_DEPLOYER_SLOT);
        address previous = slot.value;
        slot.value = finalizer;
        emit HyperCoreDeployerSet(previous, finalizer);
    }

    /// @inheritdoc IHyperMintableERC20
    function hyperCoreDeployer() external view returns (address) {
        return StorageSlot.getAddressSlot(HYPERCORE_DEPLOYER_SLOT).value;
    }

    /// @inheritdoc IHyperMintableERC20
    /// @dev See §6.4 of the hardening spec: fail-closed in this exact order —
    ///      `CoreTokenIndexAlreadySet` -> `CoreTokenIndexOutOfRange` -> `CoreTokenInfoUnavailable`
    ///      -> `CoreLinkNotFinalized` -> `CoreDecimalsMismatch`. The precompile read cannot be
    ///      forged by the caller: it reflects Core's own state.
    function setCoreTokenIndex(uint64 index) external onlyLinkAuthority {
        HyperMintableERC20Storage storage $ = _getHyperMintableERC20Storage();

        require(!$.coreTokenIndexSet, CoreTokenIndexAlreadySet());
        require(index <= type(uint32).max, CoreTokenIndexOutOfRange(index));

        (bool ok, bytes memory ret) = CORE_TOKEN_INFO_PRECOMPILE.staticcall(abi.encode(uint32(index)));
        require(ok && ret.length != 0, CoreTokenInfoUnavailable(index));

        CoreTokenInfo memory info = abi.decode(ret, (CoreTokenInfo));
        require(info.evmContract == address(this), CoreLinkNotFinalized(index, info.evmContract));
        require(
            int(uint(decimals())) == int(uint(info.weiDecimals)) + int(info.evmExtraWeiDecimals),
            CoreDecimalsMismatch(decimals(), info.weiDecimals, info.evmExtraWeiDecimals)
        );

        $.coreTokenIndex = index;
        $.coreExtraWeiDecimals = info.evmExtraWeiDecimals;
        $.coreTokenIndexSet = true;
        emit CoreTokenIndexSet(index, info.evmExtraWeiDecimals);
    }

    /// @inheritdoc IHyperMintableERC20
    function coreTokenIndex() external view returns (uint64) {
        return _getHyperMintableERC20Storage().coreTokenIndex;
    }

    /// @inheritdoc IHyperMintableERC20
    function isCoreTokenIndexSet() external view returns (bool) {
        return _getHyperMintableERC20Storage().coreTokenIndexSet;
    }

    /// @inheritdoc IHyperMintableERC20
    function coreExtraWeiDecimals() external view returns (int8) {
        return _getHyperMintableERC20Storage().coreExtraWeiDecimals;
    }

    /// @inheritdoc IHyperMintableERC20
    function coreUnit() public view returns (uint) {
        int8 extra = _getHyperMintableERC20Storage().coreExtraWeiDecimals;
        if (extra <= 0) return 1;
        return 10 ** uint(uint8(extra));
    }

    /// @inheritdoc IHyperMintableERC20
    function coreTransferableAmount(uint amount) public view returns (uint) {
        uint unit = coreUnit();
        return amount - (amount % unit);
    }

    /// @inheritdoc IHyperMintableERC20
    function coreSystemAddress() public view returns (address) {
        require(_getHyperMintableERC20Storage().coreTokenIndexSet, HyperMintableERC20CoreTokenIndexNotSet());
        return address(CORE_SYSTEM_ADDRESS_PREFIX | uint160(_getHyperMintableERC20Storage().coreTokenIndex));
    }

    /// @inheritdoc IHyperMintableERC20
    function transferToCore(uint amount) external returns (uint sent) {
        // Internal `_transfer` only — no external call, so no reentrancy surface. No
        // dedicated event: the ERC20 `Transfer(to == coreSystemAddress())` already says it all.
        sent = coreTransferableAmount(amount);
        _transfer(_msgSender(), coreSystemAddress(), sent);
    }

    /// @inheritdoc IHyperMintableERC20
    function transferToCoreFor(address coreRecipient, uint amount) external returns (uint sent) {
        require(coreRecipient != address(0), CoreRecipientZero());
        address systemAddress = coreSystemAddress();
        require(coreRecipient != systemAddress, CoreRecipientIsSystemAddress());

        sent = coreTransferableAmount(amount);
        require(sent != 0, CoreAmountBelowOneCoreWei());

        _transfer(_msgSender(), coreRecipient, sent);
        _transfer(coreRecipient, systemAddress, sent);
    }
}
