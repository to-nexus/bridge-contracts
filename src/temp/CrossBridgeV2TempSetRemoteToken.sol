// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossBridgeV2} from "../CrossBridgeV2.sol";
import {Const} from "../lib/Const.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title CrossBridgeV2TempSetRemoteToken
 * @notice ⚠️ TEMPORARY / TESTNET-ONLY / ONE-SHOT / DISPOSABLE IMPLEMENTATION ⚠️
 *
 * This contract exists to fix a single stuck value:
 * `_tokenPairs[998][Const.NATIVE_TOKEN].remoteToken` on the CROSS testnet bridge proxy
 * is permanently pinned to a decommissioned `CrossMintableERC20V2` address
 * (`0x3E0217c3926106b7E585B5341439f3150c6cab7c`) with no existing admin function able to
 * change it (plan spec §1). It is meant to sit behind the proxy for a few minutes,
 * correct that one field via `setRemoteToken`, and then be replaced again.
 *
 * - Lives ONLY on branch `temp/cross-bridge-setremotetoken`. Never merged into `dev`.
 * - Never deployed to mainnet.
 * - **Restore target after use**: `0xe52bdBd35e0Db7b2efbaF17628ed0CD9D2bA8919`
 *   (the currently-deployed `CrossBridgeV2` implementation — see
 *   `script/temp/CrossBridgeV2Temp.s.sol:restore`). Re-upgrading the proxy back to this
 *   address fully removes `setRemoteToken` from the ABI; no new build is needed for the
 *   restore step since that implementation is already deployed and linked.
 *
 * @dev Inherits `CrossBridgeV2` and adds **zero storage variables** — this is the
 * entire safety argument for storage-layout compatibility (plan spec §6/§8): with no new
 * state variables declared here, the compiler-derived storage layout of this contract is
 * IDENTICAL to `CrossBridgeV2`'s (which is itself layout-compatible with `CrossBridge`).
 * Do not add state variables to this contract for any reason.
 *
 * The target pair (`remoteChainID = 998`, `localToken = Const.NATIVE_TOKEN`) is fixed as
 * a `private constant`, making it impossible for `setRemoteToken` to touch any other
 * token pair.
 */
contract CrossBridgeV2TempSetRemoteToken is CrossBridgeV2 {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @dev HyperEVM testnet chain ID — the only remote chain this one-shot fix targets.
    uint private constant TARGET_REMOTE_CHAIN_ID = 998;

    /// @dev Native CROSS token — the only local token this one-shot fix targets.
    address private constant TARGET_LOCAL_TOKEN = Const.NATIVE_TOKEN;

    /// @notice Thrown when the target pair (998, NATIVE_TOKEN) is not registered.
    error TempPairNotRegistered();

    /// @notice Thrown when `newRemoteToken` is the zero address.
    error TempZeroRemoteToken();

    /**
     * @notice Emitted when the target pair's `remoteToken` is overwritten.
     * @param remoteChainID Always `TARGET_REMOTE_CHAIN_ID` (998).
     * @param localToken Always `TARGET_LOCAL_TOKEN` (native CROSS).
     * @param oldRemoteToken The `remoteToken` value before this call.
     * @param newRemoteToken The `remoteToken` value after this call.
     */
    event TempRemoteTokenSet(
        uint indexed remoteChainID, address indexed localToken, address indexed oldRemoteToken, address newRemoteToken
    );

    /**
     * @notice One-shot fix: overwrites `_tokenPairs[998][NATIVE_TOKEN].remoteToken`.
     * @dev Deliberately NOT `whenNotPaused` — this contract sits behind the proxy only
     * while diagnosing/fixing a stuck value, and requiring the bridge to be unpaused to
     * call this would make the fix self-locking if the bridge is ever paused during the
     * window (plan spec §11). No other field of the pair struct is touched.
     * @param newRemoteToken The corrected HyperEVM-side token address.
     */
    function setRemoteToken(address newRemoteToken) external onlyRole(Const.ADMIN_ROLE) {
        require(newRemoteToken != address(0), TempZeroRemoteToken());
        require(_tokens[TARGET_REMOTE_CHAIN_ID].contains(TARGET_LOCAL_TOKEN), TempPairNotRegistered());

        TokenPair storage pair = _tokenPairs[TARGET_REMOTE_CHAIN_ID][TARGET_LOCAL_TOKEN];
        address oldRemoteToken = pair.remoteToken;
        pair.remoteToken = newRemoteToken;

        emit TempRemoteTokenSet(TARGET_REMOTE_CHAIN_ID, TARGET_LOCAL_TOKEN, oldRemoteToken, newRemoteToken);
    }
}
