// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ICrossMintableERC20Code} from "./ICrossMintableERC20Code.sol";

/**
 * @title IHyperMintableERC20Code
 * @notice `ICrossMintableERC20Code` extended with the HyperCore link-delegation surface for
 *         tokens created by this factory, plus an explicit name/symbol/minter creation path.
 */
interface IHyperMintableERC20Code is ICrossMintableERC20Code {
    /**
     * @notice Emitted when a `HyperMintableERC20` is deployed by this factory
     * @param remoteChainID Chain ID of the remote token this local token wraps
     * @param remoteToken Address of the remote token
     * @param tokenAddress Address of the newly deployed `HyperMintableERC20`
     */
    event HyperMintableERC20Created(uint indexed remoteChainID, address indexed remoteToken, address tokenAddress);

    /**
     * @notice Delegates to `token.setHyperCoreDeployer(finalizer)` on a token this factory created
     * @dev Reverts with `HyperMintableERC20CodeUnknownToken` for any other address. Restricted
     *      to `Const.ADMIN_ROLE`.
     * @param token Address of a `HyperMintableERC20` this factory deployed
     * @param finalizer Address that will sign the matching HyperCore `finalizeEvmContract` action
     */
    function setHyperCoreDeployer(address token, address finalizer) external;

    /**
     * @notice Delegates to `token.setCoreTokenIndex(index)` on a token this factory created
     * @dev Reverts with `HyperMintableERC20CodeUnknownToken` for any other address. Restricted
     *      to `Const.ADMIN_ROLE`.
     * @param token Address of a `HyperMintableERC20` this factory deployed
     * @param index HyperCore spot token index
     */
    function setCoreTokenIndex(address token, uint64 index) external;

    /**
     * @notice Predicts the CREATE2 address `createCrossMintableERC20` would deploy to
     * @dev Shares the exact same `_initCode(...)` builder as `createCrossMintableERC20` (via the
     *      same derived name/symbol), so the prediction and the real deployment can never
     *      structurally diverge. `minter` is the address that will actually call
     *      `createCrossMintableERC20` (normally the bridge) — supplying the wrong address does
     *      NOT revert, it simply returns a different address than what will actually be deployed
     *      once the real minter creates the token. This is intentional: it allows precomputing
     *      an address before the eventual minter has even been granted `Const.BRIDGE_ROLE`.
     * @param remoteChainID Chain ID of the remote token this local token would wrap
     * @param remoteToken Address of the remote token
     * @param symbol Token symbol (name is derived as "Cross Bridge <symbol>", symbol as "<symbol>x")
     * @param decimals Token decimals
     * @param minter Address that will call `createCrossMintableERC20` (the bridge)
     * @return Predicted `HyperMintableERC20` address
     */
    function computeTokenAddress(
        uint remoteChainID,
        address remoteToken,
        string memory symbol,
        uint8 decimals,
        address minter
    ) external view returns (address);

    /**
     * @notice Deploys a `HyperMintableERC20` with caller-chosen name/symbol and an explicit minter
     * @dev Same CREATE2 salt derivation (`remoteChainID`, `remoteToken`) and the same
     *      `BeaconProxy`-over-`beacon()` deployment shape as `createCrossMintableERC20` — only
     *      the initcode's name/symbol/minter inputs differ. Restricted to `Const.ADMIN_ROLE`.
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
     * @notice Predicts the CREATE2 address `createHyperMintableERC20` would deploy to
     * @dev Shares the exact same `_initCode(...)` builder as `createHyperMintableERC20`, so the
     *      prediction and the real deployment can never structurally diverge. `minter` here is the
     *      exact `minter` ARGUMENT that will be passed to `createHyperMintableERC20` and receive
     *      `MINTER_ROLE` — unlike `computeTokenAddress`, it is NOT the factory caller. Supplying a
     *      different value than the eventual creation uses does NOT revert; it simply returns a
     *      different address than what will actually be deployed.
     * @param remoteChainID Chain ID of the remote token this local token would wrap
     * @param remoteToken Address of the remote token
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals Token decimals
     * @param minter Address that will be passed as `minter` to `createHyperMintableERC20` and
     *        granted `MINTER_ROLE` (normally the bridge)
     * @return Predicted `HyperMintableERC20` address
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
     * @notice Immutable-in-effect `defaultAdmin()` passed to every token this factory creates
     * @dev As the token's current default admin, this address has effective link authority
     *      (`isLinkAuthority`) over `setHyperCoreDeployer` / `setCoreTokenIndex` without holding
     *      `LINKER_ROLE` itself — that role is instead granted to this factory (`address(this)`
     *      in the created token's `factoryLinker`). Stored (not `immutable`) because this
     *      contract is itself a UUPS proxy: the logic contract's constructor only disables
     *      initializers, so per-deployment values live in storage and survive a logic upgrade.
     * @return Configured token admin address
     */
    function tokenAdmin() external view returns (address);

    /**
     * @notice The `UpgradeableBeacon` every `HyperMintableERC20` created by this factory points at
     * @dev Also stored (not `immutable`) for the same reason as `tokenAdmin` — see there.
     * @return Beacon address baked into every created token's `BeaconProxy` initcode
     */
    function beacon() external view returns (address);

    /**
     * @notice Returns whether `token` was deployed by this factory
     * @param token Address to check
     * @return True if this factory created `token`
     */
    function isHyperMintableERC20(address token) external view returns (bool);
}
