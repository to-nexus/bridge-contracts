// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {TransientSlot} from "@openzeppelin/contracts/utils/TransientSlot.sol";

import {ICrossBridgeV2} from "../interface/ICrossBridgeV2.sol";

/**
 * @title ForwardLib
 * @notice External library backing the executor-only forwarded (multi-hop) bridge
 * entrypoint (`bridgeTokenForwarded`) on relay-hub bridges such as `CrossBridgeV2`.
 * @dev Deployed as a separate contract and DELEGATECALL'd by the host bridge: all
 * transient storage touched here lives in the HOST's transient space, and all events
 * emitted here are emitted with the HOST's address as `emitter`. Every function here
 * MUST stay `external` (never `internal`) — an `internal` library function is inlined
 * into the caller's bytecode by the optimizer instead of being linked and
 * DELEGATECALL'd, which would defeat the point of splitting this logic out.
 *
 * Transient context layout (ERC-7201-style fixed base + offsets; EIP-1153 storage, so
 * none of this occupies a persistent storage slot):
 *   CTX_BASE+0  active (bool)         CTX_BASE+1  fromChainID (uint)
 *   CTX_BASE+2  toToken (address)     CTX_BASE+3  value (uint)
 *   CTX_BASE+4  finalizeIndex (uint)
 * The forward reentrancy guard lives under its own, separately-derived base slot.
 */
library ForwardLib {
    using TransientSlot for *;

    // keccak256(abi.encode(uint256(keccak256("nexus.bridge.forward.ctx")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CTX_BASE = 0xf985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f400;

    bytes32 private constant CTX_ACTIVE = CTX_BASE;
    bytes32 private constant CTX_FROM_CHAIN_ID = bytes32(uint256(CTX_BASE) + 1);
    bytes32 private constant CTX_TOKEN = bytes32(uint256(CTX_BASE) + 2);
    bytes32 private constant CTX_VALUE = bytes32(uint256(CTX_BASE) + 3);
    bytes32 private constant CTX_FINALIZE_INDEX = bytes32(uint256(CTX_BASE) + 4);

    // keccak256(abi.encode(uint256(keccak256("nexus.bridge.forward.guard")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant GUARD_SLOT = 0x96c5fec86ccff8d1ee06d7ba5a223396fbf7f905e4727220bd20cf4e693d0500;

    /// @notice Nested/reentrant call while a forward reentrancy guard is already held.
    error ForwardReentrantCall();

    /// @notice Classification of a hop-2+ forward attempt that did not complete successfully.
    enum ForwardFailureCode {
        None,
        ForwardNotAttempted,
        ExecutorCallReverted,
        ForwardCallNotConsumed
    }

    /**
     * @notice Emitted (as the host bridge, via DELEGATECALL) when a forwarded hop-2+
     * bridge is successfully initiated from within a finalize's executor call.
     */
    event ForwardInitiated(
        uint indexed fromChainID,
        uint indexed finalizeIndex,
        uint indexed toChainID,
        address fromToken,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes32 extraDataHash
    );

    /**
     * @notice Emitted (as the host bridge, via DELEGATECALL) when a staged forward
     * context was not consumed by a successful `bridgeTokenForwarded` call.
     * @param reasonHash keccak256 of the executor-frame revert/return data, or 0 when
     * the executor was never invoked (`ForwardNotAttempted`).
     */
    event ForwardFailed(
        uint indexed fromChainID, uint indexed finalizeIndex, ForwardFailureCode code, bytes32 reasonHash
    );

    /**
     * @notice Acquires the call-scoped forward reentrancy guard.
     * @dev Reverts if already held (nested/reentrant call). Released by `releaseGuard`
     * on normal return; an EIP-1153 rollback on revert releases it automatically. This
     * guard is call-scoped, not transaction-scoped, so a batch with multiple forwarded
     * items can have all of them succeed.
     */
    function acquireGuard() external {
        if (GUARD_SLOT.asBoolean().tload()) revert ForwardReentrantCall();
        GUARD_SLOT.asBoolean().tstore(true);
    }

    /// @notice Releases the call-scoped forward reentrancy guard.
    function releaseGuard() external {
        GUARD_SLOT.asBoolean().tstore(false);
    }

    /**
     * @notice Stages the forward context for the finalize item about to be forwarded.
     * @dev Called from `_forwardBegin`, before the executor is invoked.
     */
    function setCtx(uint fromChainID, address token, uint value, uint finalizeIndex) external {
        CTX_FROM_CHAIN_ID.asUint256().tstore(fromChainID);
        CTX_TOKEN.asAddress().tstore(token);
        CTX_VALUE.asUint256().tstore(value);
        CTX_FINALIZE_INDEX.asUint256().tstore(finalizeIndex);
        CTX_ACTIVE.asBoolean().tstore(true);
    }

    /**
     * @notice Consumes (read-then-zero) the staged forward context.
     * @dev Reverts with `ICrossBridgeV2.BaseBridgeForwardContextInactive` if inactive.
     * Zeroing on read (rather than via a later explicit clear) means nested
     * reentrancy and a sequential second call are blocked by the very same mechanism:
     * only one `bridgeTokenForwarded` call per staged context can ever succeed.
     */
    function consumeCtx() external returns (uint fromChainID, address token, uint value, uint finalizeIndex) {
        if (!CTX_ACTIVE.asBoolean().tload()) revert ICrossBridgeV2.BaseBridgeForwardContextInactive();

        fromChainID = CTX_FROM_CHAIN_ID.asUint256().tload();
        token = CTX_TOKEN.asAddress().tload();
        value = CTX_VALUE.asUint256().tload();
        finalizeIndex = CTX_FINALIZE_INDEX.asUint256().tload();

        CTX_ACTIVE.asBoolean().tstore(false);
        CTX_FROM_CHAIN_ID.asUint256().tstore(0);
        CTX_TOKEN.asAddress().tstore(address(0));
        CTX_VALUE.asUint256().tstore(0);
        CTX_FINALIZE_INDEX.asUint256().tstore(0);
    }

    /// @notice Emits `ForwardInitiated` (as the host bridge, via DELEGATECALL).
    function emitInitiated(
        uint fromChainID,
        uint finalizeIndex,
        uint toChainID,
        address fromToken,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes32 extraDataHash
    ) external {
        emit ForwardInitiated(
            fromChainID, finalizeIndex, toChainID, fromToken, to, value, networkFee, exFee, extraDataHash
        );
    }

    /**
     * @notice Classifies a resolved forward attempt, clears any residual staged
     * context, and emits `ForwardFailed` (as the host bridge) when it did not succeed.
     * @dev Called from `_forwardEnd` on every exit path of `_finalizeBridge` for which
     * `_forwardBegin` returned true, including the early success `return` — so a
     * staged-but-unconsumed context never leaks into the next finalize batch item.
     * @param executorInvoked Whether the executor was actually called (false means the
     * whitelist or ERC20 approve gate rejected the item before the executor was ever
     * reached, so the classification is `ForwardNotAttempted`).
     * @param ok The outer executor-call frame's success flag (meaningless when
     * `executorInvoked` is false).
     * @param result The outer executor-call frame's return/revert data.
     */
    function endAndReport(uint fromChainID, uint finalizeIndex, bool executorInvoked, bool ok, bytes memory result)
        external
    {
        bool consumed = !CTX_ACTIVE.asBoolean().tload();

        // Always clear: a staged-but-unconsumed context must never leak into the next
        // finalize batch item.
        CTX_ACTIVE.asBoolean().tstore(false);
        CTX_FROM_CHAIN_ID.asUint256().tstore(0);
        CTX_TOKEN.asAddress().tstore(address(0));
        CTX_VALUE.asUint256().tstore(0);
        CTX_FINALIZE_INDEX.asUint256().tstore(0);

        ForwardFailureCode code;
        if (!executorInvoked) {
            code = ForwardFailureCode.ForwardNotAttempted;
        } else if (!ok) {
            code = ForwardFailureCode.ExecutorCallReverted;
        } else if (!consumed) {
            code = ForwardFailureCode.ForwardCallNotConsumed;
        } else {
            // Success: `ForwardInitiated` was already emitted by the forwarded call itself.
            return;
        }

        bytes32 reasonHash = executorInvoked ? keccak256(result) : bytes32(0);
        emit ForwardFailed(fromChainID, finalizeIndex, code, reasonHash);
    }
}
