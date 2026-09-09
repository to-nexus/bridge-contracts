// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ICrossMintableERC20V2Code} from "./ICrossMintableERC20V2Code.sol";

/**
 * @title IHyperMintableERC20Code
 * @notice `ICrossMintableERC20V2Code` (the full Cross factory surface: legacy bridge-only
 *         creation path, explicit name/symbol/minter creation, CREATE2 address prediction,
 *         shared beacon, and owner-managed pass-through setters) extended with the HyperCore
 *         link-delegation surface for tokens created by this factory, plus a compat-named
 *         creation entrypoint.
 * @dev `HyperMintableERC20Code` inherits `CrossMintableERC20V2Code` rather than redeclaring an
 *      equivalent `tokenAdmin()`/`beacon()`/`computeTokenAddress*` surface, so this interface
 *      inherits the same functions from `ICrossMintableERC20V2Code` rather than restating
 *      them. There is no `tokenAdmin()` here — every token this factory creates gets
 *      `address(this)` (this factory) as `defaultAdmin()`, exactly like
 *      `CrossMintableERC20V2Code`'s own tokens.
 */
interface IHyperMintableERC20Code is ICrossMintableERC20V2Code {
    /**
     * @notice Emitted, alongside `ICrossMintableERC20V2Code.CrossMintableERC20Created`, for every
     *         token this factory creates
     * @dev Preserved for ABI/tooling continuity: existing off-chain consumers of this factory
     *      already watch this topic, so it must keep being emitted. `_create` (inherited,
     *      unmodified, from `CrossMintableERC20V2Code`) calls the overridable `_emitCreated`
     *      hook, which this factory overrides to emit BOTH
     *      `ICrossMintableERC20V2Code.CrossMintableERC20Created` and then this event — in that
     *      order, in the same transaction — for every token this factory creates, whether via
     *      `createCrossMintableERC20`, `createMintableERC20`, or this factory's own
     *      `createHyperMintableERC20`. The two events describe ONE token creation, not two:
     *      consumers subscribed to both topics MUST deduplicate by transaction hash or by the
     *      created `tokenAddress`, never count a single creation twice.
     * @param remoteChainID Chain ID of the remote token this local token wraps
     * @param remoteToken Address of the remote token
     * @param tokenAddress Address of the newly deployed `HyperMintableERC20`
     */
    event HyperMintableERC20Created(uint indexed remoteChainID, address indexed remoteToken, address tokenAddress);

    /**
     * @notice Delegates to `token.setHyperCoreDeployer(finalizer)` on a token this factory created
     * @dev Reverts with `CrossMintableERC20V2CodeUnknownToken` for any other address. Restricted
     *      to `Const.ADMIN_ROLE`.
     * @param token Address of a `HyperMintableERC20` this factory deployed
     * @param finalizer Address that will sign the matching HyperCore `finalizeEvmContract` action
     */
    function setHyperCoreDeployer(address token, address finalizer) external;

    /**
     * @notice Delegates to `token.setCoreTokenIndex(index)` on a token this factory created
     * @dev Reverts with `CrossMintableERC20V2CodeUnknownToken` for any other address. Restricted
     *      to `Const.ADMIN_ROLE`.
     * @param token Address of a `HyperMintableERC20` this factory deployed
     * @param index HyperCore spot token index
     */
    function setCoreTokenIndex(address token, uint64 index) external;

    /**
     * @notice Deploys a `HyperMintableERC20` with caller-chosen name/symbol and an explicit minter
     * @dev Thin compat wrapper kept ONLY so this selector — already relied on by this
     *      UUPS-upgradeable factory's existing operational tooling — does not disappear; it
     *      delegates to the exact same inherited creation path as
     *      `ICrossMintableERC20V2Code.createMintableERC20`. Same CREATE2 salt derivation
     *      (`remoteChainID`, `remoteToken`) and the same `BeaconProxy`-over-`beacon()` deployment
     *      shape as `createCrossMintableERC20` — only the initcode's name/symbol/minter inputs
     *      differ. Restricted to `Const.ADMIN_ROLE`.
     * @param remoteChainID Chain ID of the remote token this local token wraps
     * @param remoteToken Address of the remote token
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals Token decimals
     * @param minter Address granted `MINTER_ROLE` on the created token
     * @return tokenAddress Address of the newly deployed `HyperMintableERC20`
     */
    function createHyperMintableERC20(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external returns (address tokenAddress);

    /**
     * @notice Returns whether `token` was deployed by this factory
     * @dev Compat alias for the inherited `isCrossMintableERC20` — kept so this selector, relied
     *      on by existing operational tooling, does not disappear.
     * @param token Address to check
     * @return True if this factory created `token`
     */
    function isHyperMintableERC20(address token) external view returns (bool);
}
