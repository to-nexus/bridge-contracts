// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";
import {IHyperMintableERC20} from "./IHyperMintableERC20.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {ERC20, ERC20Permit, IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import {StorageSlot} from "@openzeppelin/contracts/utils/StorageSlot.sol";

/**
 * @title HyperMintableERC20
 * @notice `CrossMintableERC20V2`-shaped wrapped token with a HyperCore link slot on top.
 * @dev Deploying a bridge-wrapped token on HyperEVM with the ordinary `CrossMintableERC20V2Code`
 *      factory leaves the token with no way to link to a HyperCore spot asset: of the three
 *      `finalizeEvmContract` variants Core offers, `create{nonce}` requires an EOA CREATE
 *      deployment (the factory deploys via CREATE2), `firstStorageSlot` needs the finalizer at
 *      storage slot 0 (occupied by OZ `ERC20._balances`), and only `customStorageSlot` — which
 *      reads the finalizer from `keccak256("HyperCore deployer")` — is usable. This contract
 *      exposes that slot through `setHyperCoreDeployer` / `hyperCoreDeployer`.
 *
 *      `HYPERCORE_DEPLOYER_SLOT` is a plain `keccak256("HyperCore deployer")` digest, exactly as
 *      Hyperliquid's documentation specifies it — it is NOT ERC-7201 masked (`-1 & ~0xff`).
 *      Because it is a keccak256 output, collision with Solidity's sequential slots (0, 1, 2, ...)
 *      or mapping/array derived slots is negligible, the same argument ERC-7201 itself relies on.
 */
contract HyperMintableERC20 is ERC20, ERC20Permit, IHyperMintableERC20, AccessControlDefaultAdminRules {
    error HyperMintableERC20CoreTokenIndexNotSet();

    /// @dev keccak256("HyperCore deployer") — the ONLY slot Core's `customStorageSlot` variant
    /// of `finalizeEvmContract` reads. Plain digest, not ERC-7201 masked.
    bytes32 public constant HYPERCORE_DEPLOYER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

    /// @dev Top byte of every HyperCore system address (`0x20...` + 19-byte token index).
    uint160 public constant CORE_SYSTEM_ADDRESS_PREFIX = uint160(0x20) << 152;

    uint8 private immutable _decimals;

    /// @dev Packed with `_coreTokenIndexSet` into a single storage slot.
    uint64 private _coreTokenIndex;
    /// @dev Index 0 is a valid HyperCore index (e.g. USDC), so it cannot double as an
    /// "unset" sentinel — this flag is the actual source of truth.
    bool private _coreTokenIndexSet;

    /// @notice The one address that may hold role-based link authority (the creating factory).
    /// @dev Immutable: a `LINKER_ROLE` grant to any other account confers nothing, so an
    ///      accidental or malicious grant cannot widen who may write the HyperCore slots.
    address public immutable factoryLinker;

    /**
     * @param initialOwner Default admin of the token; always has link authority regardless of role
     * @param initialMinter Address granted `MINTER_ROLE` if non-zero (normally the bridge)
     * @param initialLinker Address granted `LINKER_ROLE` if non-zero (normally the creating factory);
     *        stored as the immutable `factoryLinker`
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals_ ERC20 decimals
     */
    constructor(
        address initialOwner,
        address initialMinter,
        address initialLinker,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) ERC20(name_, symbol_) ERC20Permit(name_) AccessControlDefaultAdminRules(0, initialOwner) {
        if (initialMinter != address(0)) {
            // Grant the minter role to the bridge
            _grantRole(Const.MINTER_ROLE, initialMinter);
        }

        factoryLinker = initialLinker;
        if (initialLinker != address(0)) {
            // Grant the linker role to the creating factory
            _grantRole(Const.LINKER_ROLE, initialLinker);
        }

        _decimals = decimals_;
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
        return account == factoryLinker && hasRole(Const.LINKER_ROLE, account);
    }

    function mint(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    function decimals() public view override(ERC20, ICrossMintableERC20) returns (uint8) {
        return _decimals;
    }

    function nonces(address owner_) public view override(ERC20Permit, ICrossMintableERC20) returns (uint) {
        return super.nonces(owner_);
    }

    /// @inheritdoc IHyperMintableERC20
    function setHyperCoreDeployer(address finalizer) external onlyLinkAuthority {
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
    function setCoreTokenIndex(uint64 index) external onlyLinkAuthority {
        emit CoreTokenIndexSet(_coreTokenIndex, _coreTokenIndexSet, index);
        _coreTokenIndex = index;
        _coreTokenIndexSet = true;
    }

    /// @inheritdoc IHyperMintableERC20
    function coreTokenIndex() external view returns (uint64) {
        return _coreTokenIndex;
    }

    /// @inheritdoc IHyperMintableERC20
    function isCoreTokenIndexSet() external view returns (bool) {
        return _coreTokenIndexSet;
    }

    /// @inheritdoc IHyperMintableERC20
    function coreSystemAddress() public view returns (address) {
        require(_coreTokenIndexSet, HyperMintableERC20CoreTokenIndexNotSet());
        return address(CORE_SYSTEM_ADDRESS_PREFIX | uint160(_coreTokenIndex));
    }

    /// @inheritdoc IHyperMintableERC20
    function transferToCore(uint amount) external returns (bool) {
        // Internal `_transfer` only — no external call, so no reentrancy surface. No
        // dedicated event: the ERC20 `Transfer(to == coreSystemAddress())` already says it all.
        _transfer(_msgSender(), coreSystemAddress(), amount);
        return true;
    }
}
