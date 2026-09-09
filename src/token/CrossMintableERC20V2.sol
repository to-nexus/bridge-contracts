// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";

import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PermitUpgradeable} from
    "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";

/**
 * @title CrossMintableERC20V2
 * @notice Upgradeable wrapped bridge token, deployed as a `BeaconProxy` so every token created by
 *         a `CrossMintableERC20V2Code` factory shares one `UpgradeableBeacon` — a single beacon
 *         upgrade fixes a defect in every already-deployed token at once. Replaces the earlier
 *         non-upgradeable design, where the same fix required a costly migration (deploy a new
 *         token, remap the pair, move liquidity over) — a real cost paid once already on
 *         HyperEVM, where `67e18` of the old token remain stranded.
 * @dev `immutable` state cannot be used behind a proxy — it is baked into the logic contract's
 *      own bytecode, not the proxy's storage — so `decimals` lives in ERC-7201 namespaced
 *      storage here instead of the `immutable` field the old design used.
 *
 *      This token has no `bridge`/admin storage of its own: `initialize`'s `initialOwner`
 *      becomes `defaultAdmin()` via `AccessControlDefaultAdminRules`, and the creating factory
 *      always passes its OWN address (`address(this)` from the factory's `_initCode`) for that
 *      parameter — see `CrossMintableERC20V2Code` for why that is deliberately not a separately
 *      stored `tokenAdmin`.
 */
contract CrossMintableERC20V2 is
    ERC20Upgradeable,
    ERC20PermitUpgradeable,
    ICrossMintableERC20,
    AccessControlDefaultAdminRulesUpgradeable
{
    /// @custom:storage-location erc7201:nexus.storage.CrossMintableERC20V2
    struct CrossMintableERC20V2Storage {
        uint8 decimals;
    }

    // keccak256(abi.encode(uint256(keccak256("nexus.storage.CrossMintableERC20V2")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CrossMintableERC20V2StorageLocation =
        0x9581ab046f3e11ba7ecf3d09df64f89f69c3a2a7a7e6f470b2498e2d7b1f5700;

    function _getCrossMintableERC20V2Storage() private pure returns (CrossMintableERC20V2Storage storage $) {
        assembly {
            $.slot := CrossMintableERC20V2StorageLocation
        }
    }

    /// @dev Implementation contract is never initialized directly — only `BeaconProxy` instances
    /// pointing at it are (see `CrossMintableERC20V2Code._initCode`).
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes a token instance. Called once, atomically, from the `BeaconProxy`
     *         constructor via `CrossMintableERC20V2Code._initCode`.
     * @dev `virtual` so `HyperMintableERC20` can override this exact signature to permanently
     *      revert — closing the linker-less init path — while reusing the actual initialization
     *      logic via `__CrossMintableERC20V2_init` under its own, differently-shaped public
     *      initializer.
     * @param initialOwner Becomes `defaultAdmin()` of this token. The creating factory always
     *        passes its own address here — there is no separately stored token admin.
     * @param initialMinter Address granted `MINTER_ROLE` if non-zero (normally the bridge)
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals_ ERC20 decimals
     */
    function initialize(
        address initialOwner,
        address initialMinter,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) public virtual initializer {
        __CrossMintableERC20V2_init(initialOwner, initialMinter, name_, symbol_, decimals_);
    }

    /**
     * @notice Extracted initialization body, callable from a subclass's own differently-shaped
     *         public initializer (see `HyperMintableERC20.initialize`, the 6-arg overload).
     * @dev `internal onlyInitializing` — mirrors `BaseBridge.__BaseBridge_init`'s pattern: only
     *      callable from within another `initializer`-guarded function, never directly.
     */
    function __CrossMintableERC20V2_init(
        address initialOwner,
        address initialMinter,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) internal onlyInitializing {
        __ERC20_init(name_, symbol_);
        __ERC20Permit_init(name_);
        __AccessControlDefaultAdminRules_init(0, initialOwner);

        if (initialMinter != address(0)) {
            // Grant the minter role to the bridge
            _grantRole(Const.MINTER_ROLE, initialMinter);
        }

        _getCrossMintableERC20V2Storage().decimals = decimals_;
    }

    function mint(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _mint(_account, _amount);
        return true;
    }

    function burn(address _account, uint _amount) external onlyRole(Const.MINTER_ROLE) returns (bool) {
        _burn(_account, _amount);
        return true;
    }

    /// @dev `virtual`, though `HyperMintableERC20` does NOT override this: it deliberately reuses
    /// this exact implementation (and this contract's ERC-7201 storage) for `decimals` rather than
    /// keeping its own copy — `virtual` only for consistency/future subclasses.
    function decimals() public view virtual override(ERC20Upgradeable, ICrossMintableERC20) returns (uint8) {
        return _getCrossMintableERC20V2Storage().decimals;
    }

    function nonces(address owner_)
        public
        view
        virtual
        override(ERC20PermitUpgradeable, ICrossMintableERC20)
        returns (uint)
    {
        return super.nonces(owner_);
    }
}
