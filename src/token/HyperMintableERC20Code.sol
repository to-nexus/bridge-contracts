// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {CrossMintableERC20V2Code} from "./CrossMintableERC20V2Code.sol";
import {IHyperMintableERC20} from "./IHyperMintableERC20.sol";
import {IHyperMintableERC20Code} from "./IHyperMintableERC20Code.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

/**
 * @title HyperMintableERC20Code
 * @notice Factory that deploys `HyperMintableERC20` tokens via CREATE2 (as `BeaconProxy`
 *         instances sharing one `UpgradeableBeacon`) and, like its `CrossMintableERC20V2Code`
 *         parent, keeps the initcode deterministic enough to predict a token's address before it
 *         is created — even across a beacon logic upgrade, since the initcode only ever embeds
 *         the beacon's own (fixed) address, never the implementation it currently points at.
 * @dev Inherits `CrossMintableERC20V2Code` rather than redeclaring an equivalent factory:
 *      `beacon`/`isCrossMintableERC20`/`tokenForPair`/`initialize`/`_create`/the legacy
 *      `createCrossMintableERC20` path/`createMintableERC20`/both address-prediction views/the
 *      four owner-managed pass-through setters are ALL inherited unchanged. This contract
 *      overrides ONLY `_initCode`, to build a `HyperMintableERC20` (not `CrossMintableERC20V2`)
 *      initcode with the extra `initialLinker` argument, and adds the HyperCore link-delegation
 *      surface on top.
 *
 *      Every token this factory creates gets `address(this)` (this factory) as `defaultAdmin()`,
 *      via `_initCode` below — there is no separately stored `tokenAdmin` (see
 *      `IHyperMintableERC20Code`). A token's `defaultAdmin()` and `factoryLinker` therefore
 *      always converge on this SAME factory address, so `isLinkAuthority` collapses to one
 *      principal for every token this factory creates. A token whose `defaultAdmin()` is instead
 *      a separately-set EOA was not created through this factory's `_initCode`, and
 *      `isLinkAuthority` on it checks both principals independently since they need not match.
 *
 *      This factory is itself deployed as a UUPS (`ERC1967Proxy`) singleton — a single instance
 *      is expected, so there is no need for a beacon here (mirrors `CrossMintableERC20V2Code`).
 */
contract HyperMintableERC20Code is CrossMintableERC20V2Code, IHyperMintableERC20Code {
    // `CrossMintableERC20V2CodeZeroAddress` / `CrossMintableERC20V2CodeUnknownToken` /
    // `CrossMintableERC20V2CodePairAlreadyCreated` are declared on `ICrossMintableERC20V2Code`
    // and inherited from `CrossMintableERC20V2Code` — not redeclared here. Likewise `initialize`,
    // `_create`, `beacon()`, `isCrossMintableERC20()`, `tokenForPair()`, `upgradeBeacon()`,
    // `grantTokenRole()`, `revokeTokenRole()`, `beginTokenDefaultAdminTransfer()`,
    // `createCrossMintableERC20()`, `createMintableERC20()`, `computeTokenAddress()`, and
    // `computeTokenAddressWithName()` are ALL inherited unchanged — see `CrossMintableERC20V2Code`.

    // No own constructor: `CrossMintableERC20V2Code`'s constructor already calls
    // `_disableInitializers()`. Implementation contract is never initialized directly — only the
    // `ERC1967Proxy` pointing at it is (see deploy scripts).

    /**
     * @dev Overrides the inherited `_initCode` to build a `HyperMintableERC20` initcode instead
     *      of a `CrossMintableERC20V2` one: same `BeaconProxy(beacon, initData)` shape and the
     *      same `address(this)` default-admin argument, plus the extra
     *      `initialLinker = address(this)` argument that grants
     *      THIS factory `LINKER_ROLE` on the created token (stored as its `factoryLinker`).
     *      `_create` (inherited, unmodified) calls this virtual function, so both creation paths
     *      — the inherited `createCrossMintableERC20` and this contract's own
     *      `createHyperMintableERC20` — produce `HyperMintableERC20` `BeaconProxy` instances
     *      through the exact same code path, and so do both inherited address-prediction views.
     *
     *      Reads `beacon()` via `this.` (an external self-call, view-compatible) rather than the
     *      parent's storage getter directly: that getter is `private` to
     *      `CrossMintableERC20V2Code` and therefore not reachable from this derived contract's
     *      own function bodies.
     */
    function _initCode(string memory name_, string memory symbol_, uint8 decimals_, address minter)
        internal
        view
        override
        returns (bytes memory)
    {
        bytes memory initData = abi.encodeCall(
            IHyperMintableERC20.initialize,
            (
                address(this), // Initial owner -> becomes defaultAdmin()
                minter, // Initial minter
                address(this), // Initial linker -> this factory; granted LINKER_ROLE and stored as factoryLinker
                name_,
                symbol_,
                decimals_
            )
        );
        return abi.encodePacked(type(BeaconProxy).creationCode, abi.encode(this.beacon(), initData));
    }

    /**
     * @dev Overrides the inherited `_emitCreated` hook so every token this factory creates —
     *      through ANY of the three entry points that eventually call `_create` (the inherited
     *      `createCrossMintableERC20`, the inherited `createMintableERC20`, and this contract's
     *      own `createHyperMintableERC20`) — emits BOTH the base `CrossMintableERC20Created` (via
     *      `super._emitCreated`, unchanged deploy/registry logic) AND the legacy
     *      `HyperMintableERC20Created` topic that existing off-chain tooling for this factory
     *      already watches. Both events describe the exact SAME token creation in the SAME
     *      transaction — see `IHyperMintableERC20Code.HyperMintableERC20Created` for the required
     *      consumer-side dedup rule (by tx hash or created token address), never process both as
     *      two creations.
     */
    function _emitCreated(uint remoteChainID, address remoteToken, address tokenAddress) internal override {
        super._emitCreated(remoteChainID, remoteToken, tokenAddress);
        emit HyperMintableERC20Created(remoteChainID, remoteToken, tokenAddress);
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
    function isHyperMintableERC20(address token) external view returns (bool) {
        return this.isCrossMintableERC20(token);
    }

    /// @inheritdoc IHyperMintableERC20Code
    function setHyperCoreDeployer(address token, address finalizer)
        external
        onlyRole(Const.ADMIN_ROLE)
        onlyKnownToken(token)
    {
        IHyperMintableERC20(token).setHyperCoreDeployer(finalizer);
    }

    /// @inheritdoc IHyperMintableERC20Code
    function setCoreTokenIndex(address token, uint64 index) external onlyRole(Const.ADMIN_ROLE) onlyKnownToken(token) {
        IHyperMintableERC20(token).setCoreTokenIndex(index);
    }
}
