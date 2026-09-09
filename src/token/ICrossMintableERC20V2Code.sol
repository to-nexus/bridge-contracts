// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ICrossMintableERC20Code} from "./ICrossMintableERC20Code.sol";

/**
 * @title ICrossMintableERC20V2Code
 * @notice `ICrossMintableERC20Code` (the legacy, selector-frozen bridge path) extended with the
 *         `CrossMintableERC20V2Code` factory's own surface: an explicit name/symbol/minter
 *         creation path, CREATE2 address prediction for both paths, the shared beacon, and the
 *         owner-only pass-through management functions that stand in for a per-token admin now
 *         that every created token's `defaultAdmin()` is the factory itself.
 */
interface ICrossMintableERC20V2Code is ICrossMintableERC20Code {
    /**
     * @notice `initialize`'s `tokenImplementation` argument was the zero address
     */
    error CrossMintableERC20V2CodeZeroAddress();

    /**
     * @notice `token` was not deployed by this factory
     * @dev Gates the token-scoped pass-through setters (`grantTokenRole`/`revokeTokenRole`/
     *      `beginTokenDefaultAdminTransfer`) to tokens this factory itself created.
     * @param token The address that is not a token this factory created
     */
    error CrossMintableERC20V2CodeUnknownToken(address token);

    /**
     * @notice This factory already created a token for `(remoteChainID, remoteToken)`
     * @dev Enforced identically by both creation paths so crossing paths, changing the
     *      name/symbol, or changing the minter cannot bypass the one-token-per-pair guard.
     * @param remoteChainID Chain ID of the remote token
     * @param remoteToken Address of the remote token
     * @param existingToken The token this factory already created for this pair
     */
    error CrossMintableERC20V2CodePairAlreadyCreated(uint remoteChainID, address remoteToken, address existingToken);

    /**
     * @notice Emitted when a `CrossMintableERC20V2` is deployed by this factory, by either
     *         creation path
     * @dev Emitted from the overridable `_emitCreated` hook. A subclass factory
     *      that already ships its own legacy creation event (see
     *      `IHyperMintableERC20Code.HyperMintableERC20Created`) overrides that hook to emit BOTH
     *      events, in the same transaction, for the SAME token creation — a consumer subscribed
     *      to both topics MUST deduplicate by transaction hash or by `tokenAddress`, never treat
     *      them as two separate creations.
     * @param remoteChainID Chain ID of the remote token this local token wraps
     * @param remoteToken Address of the remote token
     * @param tokenAddress Address of the newly deployed `CrossMintableERC20V2`
     */
    event CrossMintableERC20Created(
        uint indexed remoteChainID, address indexed remoteToken, address indexed tokenAddress
    );

    /**
     * @notice Deploys a `CrossMintableERC20V2` with caller-chosen name/symbol and an explicit
     *         minter, as a `BeaconProxy` over this factory's shared beacon
     * @dev Same CREATE2 salt derivation (`remoteChainID`, `remoteToken`) and the same
     *      `BeaconProxy`-over-`beacon()` deployment shape as `createCrossMintableERC20` — only
     *      the initcode's name/symbol/minter inputs differ. Restricted to `Const.ADMIN_ROLE`.
     *      Reverts with `CrossMintableERC20V2CodePairAlreadyCreated` if this factory already
     *      created a token for `(remoteChainID, remoteToken)` — via either creation path.
     * @param remoteChainID Chain ID of the remote token this local token wraps
     * @param remoteToken Address of the remote token
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals Token decimals
     * @param minter Address granted `MINTER_ROLE` on the created token if non-zero
     * @return tokenAddress Address of the newly deployed `CrossMintableERC20V2`
     */
    function createMintableERC20(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external returns (address tokenAddress);

    /**
     * @notice Predicts the CREATE2 address `createCrossMintableERC20` would deploy to
     * @dev Shares the exact same `_initCode(...)` builder as `createCrossMintableERC20` (via the
     *      same derived name/symbol), so the prediction and the real deployment can never
     *      structurally diverge. `minter` is the address that will actually call
     *      `createCrossMintableERC20` (normally the bridge) — supplying the wrong address does
     *      NOT revert, it simply returns a different address than what will actually be deployed
     *      once the real minter creates the token.
     * @param remoteChainID Chain ID of the remote token this local token would wrap
     * @param remoteToken Address of the remote token
     * @param symbol Token symbol (name is derived as "Cross Bridge <symbol>", symbol as "<symbol>x")
     * @param decimals Token decimals
     * @param minter Address that will call `createCrossMintableERC20` (the bridge)
     * @return Predicted `CrossMintableERC20V2` address
     */
    function computeTokenAddress(
        uint remoteChainID,
        address remoteToken,
        string memory symbol,
        uint8 decimals,
        address minter
    ) external view returns (address);

    /**
     * @notice Predicts the CREATE2 address `createMintableERC20` would deploy to
     * @dev Shares the exact same `_initCode(...)` builder as `createMintableERC20`, so the
     *      prediction and the real deployment can never structurally diverge. `minter` here is
     *      the exact `minter` ARGUMENT that will be passed to `createMintableERC20` — unlike
     *      `computeTokenAddress`, it is NOT the factory caller. Supplying a different value than
     *      the eventual creation uses does NOT revert; it simply returns a different address.
     * @param remoteChainID Chain ID of the remote token this local token would wrap
     * @param remoteToken Address of the remote token
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals Token decimals
     * @param minter Address that will be passed as `minter` to `createMintableERC20` and granted
     *        `MINTER_ROLE` (normally the bridge)
     * @return Predicted `CrossMintableERC20V2` address
     */
    function computeTokenAddressWithName(
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) external view returns (address);

    /**
     * @notice The `UpgradeableBeacon` every `CrossMintableERC20V2` created by this factory points
     *         at
     * @dev Created and owned by this factory itself inside `initialize` — there is no
     *      separate beacon deployment or ownership hand-off step. Only its address (never the
     *      implementation it currently resolves to) is embedded in a token's initcode, so
     *      upgrading it never moves future CREATE2 predictions.
     * @return Beacon address baked into every created token's `BeaconProxy` initcode
     */
    function beacon() external view returns (address);

    /**
     * @notice Returns whether `token` was deployed by this factory
     * @param token Address to check
     * @return True if this factory created `token`
     */
    function isCrossMintableERC20(address token) external view returns (bool);

    /**
     * @notice Returns the token this factory created for `(remoteChainID, remoteToken)`, or the
     *         zero address if none exists yet
     * @dev Backs the one-token-per-pair guard: both creation paths consult and update the
     *      same mapping, so a duplicate cannot be created by crossing paths, changing the name,
     *      or changing the minter.
     * @param remoteChainID Chain ID of the remote token
     * @param remoteToken Address of the remote token
     * @return The token address already created for this pair by this factory, or `address(0)`
     */
    function tokenForPair(uint remoteChainID, address remoteToken) external view returns (address);

    /**
     * @notice Upgrades the shared token beacon — reflected in EVERY token this factory created,
     *         immediately
     * @dev Restricted to `Const.ADMIN_ROLE`. There is no staged/partial rollout: verify the new
     *      implementation thoroughly before calling this.
     * @param newImplementation New `CrossMintableERC20V2` logic address
     */
    function upgradeBeacon(address newImplementation) external;

    /**
     * @notice Grants `role` on `token` on behalf of this factory (the token's `defaultAdmin()`)
     * @dev Restricted to `Const.ADMIN_ROLE` and to tokens this factory created
     *      (`CrossMintableERC20V2CodeUnknownToken` otherwise). Cannot be used for
     *      `DEFAULT_ADMIN_ROLE` itself — the token (via `AccessControlDefaultAdminRules`) only
     *      allows that through the two-step transfer below.
     * @param token Address of a `CrossMintableERC20V2` this factory deployed
     * @param role Role identifier to grant
     * @param account Address to grant `role` to
     */
    function grantTokenRole(address token, bytes32 role, address account) external;

    /**
     * @notice Revokes `role` on `token` on behalf of this factory (the token's `defaultAdmin()`)
     * @dev Restricted to `Const.ADMIN_ROLE` and to tokens this factory created
     *      (`CrossMintableERC20V2CodeUnknownToken` otherwise).
     * @param token Address of a `CrossMintableERC20V2` this factory deployed
     * @param role Role identifier to revoke
     * @param account Address to revoke `role` from
     */
    function revokeTokenRole(address token, bytes32 role, address account) external;

    /**
     * @notice Escape hatch, but scoped strictly to TOKEN-LEVEL administration: begins
     *         transferring `token`'s `DEFAULT_ADMIN_ROLE` (and, with it, its role-management
     *         authority — `grantTokenRole`/`revokeTokenRole` above stop applying to `token` once
     *         accepted) away from this factory to `newAdmin`, on behalf of this factory (the
     *         token's current `defaultAdmin()`)
     * @dev Restricted to `Const.ADMIN_ROLE` and to tokens this factory created
     *      (`CrossMintableERC20V2CodeUnknownToken` otherwise). Only begins the transfer — the
     *      token's own `AccessControlDefaultAdminRules` two-step schedule still applies:
     *      `newAdmin` must separately call `acceptDefaultAdminTransfer()` directly on the token
     *      once the (zero, in this factory's case) delay has elapsed.
     *
     *      What this does NOT do: detach `token` from this factory's shared beacon. `token` is a
     *      `BeaconProxy` permanently pointed at `beacon()` — there is no way, in this
     *      architecture, to move a single already-deployed `BeaconProxy` to a different beacon or
     *      to freeze its implementation. So this factory's `ADMIN_ROLE` holder can still replace
     *      `token`'s logic (along with every other token this factory created) via
     *      `upgradeBeacon`, even after `token`'s `defaultAdmin()`/role-management authority has
     *      moved away. "Escape hatch" here means an escape from role management only, never from
     *      beacon-level control.
     * @param token Address of a `CrossMintableERC20V2` this factory deployed
     * @param newAdmin Address to begin transferring `token`'s `DEFAULT_ADMIN_ROLE` to
     */
    function beginTokenDefaultAdminTransfer(address token, address newAdmin) external;
}
