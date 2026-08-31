// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title ICrossBridgeV2
 * @notice ABI surface unique to `CrossBridgeV2`: the executor-only forwarded (hop-2+)
 * bridge entrypoint used for multi-hop bridging (A -> B -> C), plus the errors it can
 * revert with.
 * @dev `bridgeTokenForwarded` is intentionally NOT part of `IBaseBridge` — it only
 * exists on relay-hub bridges such as `CrossBridgeV2`, not on `BaseBridge` itself.
 */
interface ICrossBridgeV2 {
    /**
     * @notice Thrown when `bridgeTokenForwarded` is called by anyone other than the
     * configured `bridgeExecutor`.
     */
    error BaseBridgeForwardNotExecutor();

    /**
     * @notice Thrown when `bridgeTokenForwarded` is called without an active staged
     * forward context (e.g. direct call, replay of an already-consumed context, or a
     * nested call after the context was already consumed).
     */
    error BaseBridgeForwardContextInactive();

    /**
     * @notice Thrown when the D1 pass-through invariant is violated for an ERC20
     * token: both the from-chain and to-chain token pairs must be origin pairs.
     */
    error BaseBridgeForwardNotOrigin();

    /**
     * @notice Thrown when `value + networkFee + exFee` does not exactly equal the
     * staged forward context's value.
     * @param computed The caller-supplied `value + networkFee + exFee`.
     * @param ctxValue The staged forward context's value.
     */
    error BaseBridgeForwardAmountMismatch(uint computed, uint ctxValue);

    /**
     * @notice Executor-only entrypoint used to initiate hop-2+ of a multi-hop bridge.
     * @dev No modifiers: `fromToken`/`fromChainID` are never parameters and are only
     * obtained from the trusted, transient forward context staged by `_forwardBegin`
     * during the current finalize's executor call. All validation happens in the
     * function body, in order, after the context has been consumed.
     * @param toChainID Destination chain ID for hop-2+
     * @param to Recipient address on the destination chain
     * @param value Amount to bridge onward (excluding fees)
     * @param networkFee Network fee for hop-2+, pre-specified by the originating chain
     * @param exFee Exchange fee for hop-2+, pre-specified by the originating chain
     * @param extraData Additional data for further hops
     */
    function bridgeTokenForwarded(
        uint toChainID,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes calldata extraData
    ) external payable;
}
