// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {HyperMintableERC20} from "./HyperMintableERC20.sol";
import {IHyperMintableERC20} from "./IHyperMintableERC20.sol";
import {IHyperMintableERC20Code} from "./IHyperMintableERC20Code.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

/**
 * @title HyperMintableERC20Code
 * @notice Factory that deploys `HyperMintableERC20` tokens via CREATE2 and, unlike
 *         `CrossMintableERC20V2Code`, keeps the initcode deterministic enough to predict a
 *         token's address before it is created.
 * @dev Determinism rests on every initcode input being either immutable or an explicit
 *      argument: `tokenAdmin` (immutable here), `minter` (explicit argument to
 *      `computeTokenAddress`, `_msgSender()` in the real call), `address(this)`, `symbol` and
 *      `decimals`. `CrossMintableERC20V2Code` instead bakes in the mutable `defaultAdmin()` and
 *      the per-caller `_msgSender()`, which is why it cannot offer address prediction.
 */
contract HyperMintableERC20Code is AccessControlDefaultAdminRules, IHyperMintableERC20Code {
    error HyperMintableERC20CodeZeroAddress();
    error HyperMintableERC20CodeUnknownToken(address token);

    /// @dev Becomes `defaultAdmin()` of every token this factory creates. As the token's
    /// *current* default admin it has effective link authority (`isLinkAuthority`) over
    /// `setHyperCoreDeployer` / `setCoreTokenIndex` dynamically — it is never granted
    /// `LINKER_ROLE` itself, and does not need to be: authority tracks whoever currently holds
    /// `defaultAdmin()` on the token, following that token's own
    /// `beginDefaultAdminTransfer` / `acceptDefaultAdminTransfer` automatically. This factory,
    /// by contrast, IS granted `LINKER_ROLE` on each token it creates (see `_initCode`'s
    /// `address(this)` argument) and is that token's immutable `factoryLinker` — the token's
    /// admin can `revokeRole(LINKER_ROLE, factory)` to cut this factory's access off, and
    /// `grantRole` it back later to restore it; a `LINKER_ROLE` grant to any other address
    /// confers no authority (see `HyperMintableERC20.isLinkAuthority`).
    /// Immutable so it can never change out from under a CREATE2 address prediction; use a
    /// multisig or timelock address here (see `script/HyperMintableERC20Code.s.sol`).
    address public immutable tokenAdmin;

    /// @dev Tokens this factory has deployed, gating the setter delegation below.
    mapping(address => bool) public isHyperMintableERC20;

    /**
     * @param initialOwner Granted `Const.ADMIN_ROLE` on this factory
     * @param initialTokenAdmin Immutable `defaultAdmin()` passed to every created token (effective
     *        link authority via `isLinkAuthority`, without holding `LINKER_ROLE`); must be non-zero
     * @param initialBridge Granted `Const.BRIDGE_ROLE` if non-zero
     */
    constructor(address initialOwner, address initialTokenAdmin, address initialBridge)
        AccessControlDefaultAdminRules(0, initialOwner)
    {
        require(initialTokenAdmin != address(0), HyperMintableERC20CodeZeroAddress());
        tokenAdmin = initialTokenAdmin;

        // `_grantRole`, not the public `grantRole` — the latter is what makes
        // `CrossMintableERC20V2Code`'s constructor revert whenever the deployer account isn't
        // `initialOwner` itself (public `grantRole` is access-controlled and checks the caller,
        // not just `initialOwner`). `_grantRole` bypasses that check entirely, as intended here.
        _grantRole(Const.ADMIN_ROLE, initialOwner);

        if (initialBridge != address(0)) _grantRole(Const.BRIDGE_ROLE, initialBridge);
    }

    /**
     * @dev Shared by `createCrossMintableERC20` and `computeTokenAddress` so the predicted and
     *      actual initcode can never structurally diverge.
     */
    function _initCode(string memory symbol, uint8 decimals, address minter) internal view returns (bytes memory) {
        return abi.encodePacked(
            type(HyperMintableERC20).creationCode,
            abi.encode(
                tokenAdmin, // Initial owner -> becomes defaultAdmin(); effective link authority via isLinkAuthority
                minter, // Initial minter
                address(this), // Initial linker -> this factory; granted LINKER_ROLE and stored as factoryLinker
                string(abi.encodePacked("Cross Bridge ", symbol)),
                string(abi.encodePacked(symbol, "x")),
                decimals
            )
        );
    }

    function createCrossMintableERC20(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        onlyRole(Const.BRIDGE_ROLE)
        returns (address tokenAddress)
    {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        tokenAddress = Create2.deploy(0, salt, _initCode(symbol, decimals, _msgSender()));

        isHyperMintableERC20[tokenAddress] = true;
        emit HyperMintableERC20Created(remoteChainID, remoteToken, tokenAddress);
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
        return Create2.computeAddress(salt, keccak256(_initCode(symbol, decimals, minter)), address(this));
    }

    /// @inheritdoc IHyperMintableERC20Code
    function setHyperCoreDeployer(address token, address finalizer) external onlyRole(Const.ADMIN_ROLE) {
        require(isHyperMintableERC20[token], HyperMintableERC20CodeUnknownToken(token));
        IHyperMintableERC20(token).setHyperCoreDeployer(finalizer);
    }

    /// @inheritdoc IHyperMintableERC20Code
    function setCoreTokenIndex(address token, uint64 index) external onlyRole(Const.ADMIN_ROLE) {
        require(isHyperMintableERC20[token], HyperMintableERC20CodeUnknownToken(token));
        IHyperMintableERC20(token).setCoreTokenIndex(index);
    }
}
