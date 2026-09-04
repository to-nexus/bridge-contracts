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
     * @notice Shape of the tuple returned by the HyperCore read precompile's `tokenInfo(uint32)`
     * @dev `evmContract` is zero for an unlinked index and is the ONLY field that can prove a
     *      link — see `setCoreTokenIndex`. `evmExtraWeiDecimals` satisfies
     *      `EVM decimals == weiDecimals + evmExtraWeiDecimals`, range `[-2, 18]`.
     */
    struct CoreTokenInfo {
        string name;
        uint64[] spots;
        uint64 deployerTradingFeeShare;
        address deployer;
        address evmContract;
        uint8 szDecimals;
        uint8 weiDecimals;
        int8 evmExtraWeiDecimals;
    }

    /**
     * @notice Emitted when the HyperCore finalizer address stored in the link slot changes
     * @param previousDeployer Finalizer address that was previously stored
     * @param newDeployer Finalizer address now stored
     */
    event HyperCoreDeployerSet(address indexed previousDeployer, address indexed newDeployer);

    /**
     * @notice Emitted once, when `setCoreTokenIndex` succeeds and permanently finalizes the link
     * @param index HyperCore spot token index now recorded (permanent — `setCoreTokenIndex` only
     *        ever succeeds once)
     * @param evmExtraWeiDecimals Core's `evmExtraWeiDecimals` for this token, recorded alongside
     *        the index for `coreUnit()` / `coreTransferableAmount()`
     */
    event CoreTokenIndexSet(uint64 indexed index, int8 evmExtraWeiDecimals);

    /// @notice `setCoreTokenIndex` was already called successfully once; it never succeeds again
    error CoreTokenIndexAlreadySet();

    /// @notice `index` exceeds the precompile's `uint32` domain
    error CoreTokenIndexOutOfRange(uint64 index);

    /// @notice The Core read precompile reverted or returned no data for `index`
    error CoreTokenInfoUnavailable(uint64 index);

    /// @notice Core has not (yet, or ever) linked `index` to this contract's address
    error CoreLinkNotFinalized(uint64 index, address evmContract);

    /// @notice `decimals() != weiDecimals + evmExtraWeiDecimals` as reported by Core for `index`
    error CoreDecimalsMismatch(uint8 decimals, uint8 weiDecimals, int8 evmExtraWeiDecimals);

    /// @notice `setHyperCoreDeployer` is permanently locked once `setCoreTokenIndex` has succeeded
    error HyperCoreLinkAlreadyFinalized();

    /// @notice `transferToCoreFor` was called with a zero `coreRecipient`
    error CoreRecipientZero();

    /// @notice `transferToCoreFor` was called with `coreRecipient == coreSystemAddress()`
    error CoreRecipientIsSystemAddress();

    /// @notice The rounded-down amount to move is zero (below one Core wei)
    error CoreAmountBelowOneCoreWei();

    /**
     * @notice Writes `finalizer` to the `keccak256("HyperCore deployer")` storage slot
     * @dev Zero is an allowed value: it is the only way to recall a mistakenly set finalizer
     *      before Core has finalized the link. Restricted to `isLinkAuthority` — see there for
     *      why holding `Const.LINKER_ROLE` alone is not sufficient. Reverts with
     *      `HyperCoreLinkAlreadyFinalized` once `setCoreTokenIndex` has succeeded — the slot is
     *      meaningless (and unsafe to keep changing) after the Core link is permanent.
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
     * @notice Records this token's HyperCore spot asset index — succeeds at most once
     * @dev Verifies the link against the HyperCore read precompile before recording anything:
     *      the index must already be unset (`CoreTokenIndexAlreadySet` otherwise), in range for
     *      the precompile's `uint32` domain (`CoreTokenIndexOutOfRange`), resolvable
     *      (`CoreTokenInfoUnavailable`), actually linked to `address(this)`
     *      (`CoreLinkNotFinalized`), and decimals-consistent (`CoreDecimalsMismatch`). Restricted
     *      to `isLinkAuthority`.
     * @param index HyperCore spot token index
     */
    function setCoreTokenIndex(uint64 index) external;

    /**
     * @notice Returns the HyperCore spot token index recorded via `setCoreTokenIndex`
     * @return Recorded Core token index (meaningless if `isCoreTokenIndexSet()` is false)
     */
    function coreTokenIndex() external view returns (uint64);

    /**
     * @notice Returns whether `setCoreTokenIndex` has succeeded (index is permanently recorded)
     * @return True if a Core token index has been recorded
     */
    function isCoreTokenIndexSet() external view returns (bool);

    /**
     * @notice Returns this token's `evmExtraWeiDecimals` as reported by Core when the link was
     *         finalized
     * @dev Meaningless (reads as `0`) while `isCoreTokenIndexSet()` is false.
     * @return Core's `evmExtraWeiDecimals` for this token
     */
    function coreExtraWeiDecimals() external view returns (int8);

    /**
     * @notice Returns the smallest EVM-wei increment that maps to a whole Core wei
     * @dev `10 ** evmExtraWeiDecimals` when positive, `1` when zero or negative (Core is at
     *      least as fine-grained as the EVM side, so no amount needs rounding).
     * @return The EVM-wei granularity of one Core wei
     */
    function coreUnit() external view returns (uint);

    /**
     * @notice Rounds `amount` down to the nearest multiple of `coreUnit()`
     * @param amount Amount to round down
     * @return The largest multiple of `coreUnit()` that is `<= amount`
     */
    function coreTransferableAmount(uint amount) external view returns (uint);

    /**
     * @notice Derives this token's HyperCore system address from its recorded token index
     * @dev Reverts with `HyperMintableERC20CoreTokenIndexNotSet` while no index is recorded
     * @return HyperCore system address `0x20...` + token index
     */
    function coreSystemAddress() external view returns (address);

    /**
     * @notice Convenience transfer of `amount` from the caller to this token's Core system
     *         address, rounded down to `coreUnit()` so no dust is burned
     * @dev Not a standard — Hyperliquid credits Core based on the `Transfer.from` of a plain
     *      ERC20 `transfer` to the system address, so any standard-compliant transfer works too;
     *      this function is a convenience wrapper only, and any guard placed on it can be
     *      bypassed by calling `transfer` directly. Plain internal `_transfer` — no external
     *      call, so no reentrancy surface. The rounding remainder stays with the caller.
     * @param amount Amount to move to the Core system address, before rounding
     * @return sent The actually transferred (rounded-down) amount
     */
    function transferToCore(uint amount) external returns (uint sent);

    /**
     * @notice Moves `amount` (rounded down to `coreUnit()`) from the caller to Core, crediting
     *         `coreRecipient` on Core instead of the caller
     * @dev Core credits based on the `from` of the ERC20 `Transfer` it observes at the system
     *      address, so a single `transferToCore` can only ever credit the caller. This performs
     *      `_transfer(caller, coreRecipient, sent)` then `_transfer(coreRecipient, coreSystemAddress(), sent)`
     *      in the same transaction: `coreRecipient`'s EVM balance strictly nets to zero, so no
     *      role gate is needed — the only reachable effect is crediting someone else's Core
     *      balance from the caller's own tokens. The rounding remainder stays with the caller.
     * @param coreRecipient Address Core will credit (its EVM balance does not change)
     * @param amount Amount to move to Core, before rounding
     * @return sent The actually transferred (rounded-down) amount
     */
    function transferToCoreFor(address coreRecipient, uint amount) external returns (uint sent);
}
