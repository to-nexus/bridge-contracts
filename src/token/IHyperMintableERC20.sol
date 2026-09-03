// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";

/**
 * @title IHyperMintableERC20
 * @notice `ICrossMintableERC20` extended with the HyperCore link surface.
 * @dev A `HyperMintableERC20` writes its HyperCore finalizer address to the fixed storage
 *      slot `keccak256("HyperCore deployer")` so that Core's `finalizeEvmContract` action can
 *      pick it up via the `customStorageSlot` variant (see `HyperMintableERC20` for why the
 *      other two variants — `create{nonce}` and `firstStorageSlot` — are not usable here).
 */
interface IHyperMintableERC20 is ICrossMintableERC20 {
    /**
     * @notice Emitted when the HyperCore finalizer address stored in the link slot changes
     * @param previousDeployer Finalizer address that was previously stored
     * @param newDeployer Finalizer address now stored
     */
    event HyperCoreDeployerSet(address indexed previousDeployer, address indexed newDeployer);

    /**
     * @notice Emitted when the linked HyperCore spot token index changes
     * @param previousIndex Core token index that was previously stored (meaningless if `previousIsSet` is false)
     * @param previousIsSet Whether an index had been set before this call
     * @param newIndex Core token index now stored
     */
    event CoreTokenIndexSet(uint64 previousIndex, bool previousIsSet, uint64 newIndex);

    /**
     * @notice Writes `finalizer` to the `keccak256("HyperCore deployer")` storage slot
     * @dev Zero is an allowed value: it is the only way to recall a mistakenly set finalizer
     *      before Core has finalized the link. Restricted to `isLinkAuthority` — see there for
     *      why holding `Const.LINKER_ROLE` alone is not sufficient.
     * @param finalizer Address that will sign the matching HyperCore `finalizeEvmContract` action
     */
    function setHyperCoreDeployer(address finalizer) external;

    /**
     * @notice Returns the immutable factory address that is one of the two link authorities
     * @dev Set once at construction from the deploying factory (zero for a standalone
     *      deployment). A `LINKER_ROLE` grant to any other address confers no authority — see
     *      `isLinkAuthority`.
     * @return The creating factory's address, or zero if none was set at construction
     */
    function factoryLinker() external view returns (address);

    /**
     * @notice Returns whether `account` currently has effective authority over the HyperCore
     *         link slots (`setHyperCoreDeployer` / `setCoreTokenIndex`)
     * @dev Effective authority is narrower than `hasRole(Const.LINKER_ROLE, account)`: it is
     *      exactly the current `defaultAdmin()` (regardless of role membership — this tracks
     *      `beginDefaultAdminTransfer` / `acceptDefaultAdminTransfer` automatically) plus
     *      `factoryLinker`, and only while `factoryLinker` still holds `Const.LINKER_ROLE`. A
     *      third party that is granted `Const.LINKER_ROLE` has the role but not this authority;
     *      the zero address never has it, even after default-admin renunciation.
     * @param account Address to check
     * @return True if `account` may call `setHyperCoreDeployer` / `setCoreTokenIndex` right now
     */
    function isLinkAuthority(address account) external view returns (bool);

    /**
     * @notice Reads the current HyperCore finalizer address from the link slot
     * @return Address currently stored in `keccak256("HyperCore deployer")`
     */
    function hyperCoreDeployer() external view returns (address);

    /**
     * @notice Records this token's HyperCore spot asset index
     * @dev Index 0 is a valid Core index (e.g. USDC), so a separate "is set" flag is kept
     *      rather than treating 0 as an unset sentinel. Restricted to `isLinkAuthority`.
     * @param index HyperCore spot token index
     */
    function setCoreTokenIndex(uint64 index) external;

    /**
     * @notice Returns the HyperCore spot token index recorded via `setCoreTokenIndex`
     * @return Recorded Core token index (meaningless if `isCoreTokenIndexSet()` is false)
     */
    function coreTokenIndex() external view returns (uint64);

    /**
     * @notice Returns whether `setCoreTokenIndex` has been called at least once
     * @return True if a Core token index has been recorded
     */
    function isCoreTokenIndexSet() external view returns (bool);

    /**
     * @notice Derives this token's HyperCore system address from its recorded token index
     * @dev Reverts with `HyperMintableERC20CoreTokenIndexNotSet` while no index is recorded
     * @return HyperCore system address `0x20...` + token index
     */
    function coreSystemAddress() external view returns (address);

    /**
     * @notice Convenience transfer of `amount` from the caller to this token's Core system address
     * @dev Plain internal `_transfer` — no external call, so no reentrancy surface
     * @param amount Amount to move to the Core system address
     * @return True on success
     */
    function transferToCore(uint amount) external returns (bool);
}
