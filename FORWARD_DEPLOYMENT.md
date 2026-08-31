# `CrossBridgeV2` forwarded-entrypoint deployment procedure

한국어 문서: [FORWARD_DEPLOYMENT_KO.md](FORWARD_DEPLOYMENT_KO.md)

This document covers deploying/upgrading a CROSS-chain hub bridge to `CrossBridgeV2`
(the multi-hop `bridgeTokenForwarded` entrypoint, plan spec §7). It assumes the reader
is already familiar with the general UUPS upgrade flow for this repo's bridges; it only
covers what is NEW or DIFFERENT for this feature.

**Scope note.** Nothing here sends a transaction on your behalf. The verification step
(step 1) is a read-only `forge script` invocation — it never calls
`vm.broadcast`/`vm.startBroadcast`. The actual `upgradeToAndCall` transaction, executor
whitelist configuration, and any other state-changing step are yours to run through your
existing deployment tooling, following the sequence below.

## Why this procedure exists

`_forwardBegin`'s `ForwardLib.setCtx()` call runs **before** the executor
whitelist/approve gates in `_finalizeBridge` (plan spec §7.5). A `ForwardLib` link error
(wrong address, missing code, wrong code) is therefore **not** absorbed by the normal D3
fallback — it makes `bridgeTokenForwarded`'s staging call revert, which reverts the
**entire finalize batch** the forwarded item is part of. Because `finalizeIndex` is
strictly sequential (`BaseBridge.sol:391-392`), a batch that keeps reverting **blocks the
whole finalize queue for that source chain**, not just the one forwarded item. No funds
are lost, but it is an availability incident, and the fix is not obvious from the error
alone unless you already know this document exists. Get the link right before upgrading.

## Step 1 — Four-step integrity check (read-only)

Run this **before** pointing any proxy at the new implementation:

```bash
forge script script/ForwardLibVerify.s.sol \
  --sig 'run(address,address)' <newImplementation> <forwardLibAddress> \
  --rpc-url <rpc>
```

It checks, and reverts (non-zero exit code) if any of these fail:

1. `forwardLibAddress.code.length > 0`.
2. `forwardLibAddress`'s runtime codehash matches this repo's compiled `ForwardLib`
   artifact (`out/ForwardLib.sol/ForwardLib.json`) — with the library's own
   compiler-inserted self-address guard patched in, so this is a like-for-like
   comparison, not a false mismatch.
3. Every `ForwardLib` relocation the compiler recorded for `CrossBridgeV2`
   (`out/CrossBridgeV2.sol/CrossBridgeV2.json`'s `deployedBytecode.linkReferences`), read
   directly out of `newImplementation`'s **actual deployed runtime bytecode**, contains
   exactly `forwardLibAddress`.
4. The build artifact's unlinked runtime template, with every one of those relocations
   patched to `forwardLibAddress` (and `CrossBridgeV2`'s own UUPS self-address immutable
   patched to `newImplementation`), hashes to **exactly** `newImplementation`'s actual
   runtime codehash — proving the entire deployed runtime matches what the compiler
   produced, not just the 20 bytes at each relocation.

This intentionally does **not** scan the runtime for the library address as a raw byte
pattern anywhere in the code (see the script's own doc comment for why: a pattern scan
can false-positive on an unrelated constant, and can never prove every relocation was
patched). Reading the compiler-declared relocation offsets is the only way to answer "is
every relocation correctly patched, and is the rest of the code exactly what was
compiled". `test/ForwardLibVerify.t.sol` unit-tests this logic directly (correct link
passes; wrong library address, right-address-but-wrong-code, and an implementation that
merely *contains* the right address at the right offsets while being otherwise unrelated
bytecode are all rejected).

If this script reverts, **stop** — do not run the upgrade transaction. Re-deploy
`ForwardLib` (or `CrossBridgeV2`, depending on which side is wrong) and re-run this
check.

## Step 2 — Executor configuration

Once step 1 passes and the upgrade transaction has been sent:

1. Add the (now-upgraded) `CrossBridgeV2` proxy address to the `BridgeExecutor` target
   whitelist: `bridgeExecutor.addWhitelistTarget(address(bridgeProxy))`.
2. Enable method-selector checking for that target and register **only**
   `bridgeTokenForwarded`'s selector:
   ```solidity
   bridgeExecutor.setMethodCheckEnabled(address(bridgeProxy), true);
   bridgeExecutor.addWhitelistMethods(address(bridgeProxy), [CrossBridgeV2.bridgeTokenForwarded.selector]);
   ```
   Method-check is defense-in-depth, not the only defense (see plan spec §11 — `setDev`,
   `bridgeToken`, `releasePending`, etc. are already independently blocked by role checks
   and the shared reentrancy guard even with method-check disabled), but enable it
   anyway: it is cheap and it turns a future whitelist-target misconfiguration into a
   loud revert instead of a silent extra attack surface.

## Step 3 — Post-upgrade smoke test

Before relying on the forward path for real user funds, drive one **small** real
forwarded finalize end-to-end (A → B → C, smallest denomination your fee tables allow)
and confirm:

- `ForwardInitiated` is emitted with the expected `toChainID`/`value`/fee fields.
- The hop-2 finalize on chain C actually completes (target contract call succeeds, or —
  if there is no hop-3 target — the recipient receives funds on chain C).
- No `ForwardFailed` event was emitted for this item.

This is a manual/operational step, not automated by any script in this repo — automating
it would mean sending a real cross-chain transaction pair as part of a "read-only" gate,
which is out of scope for `ForwardLibVerify.s.sol` by design (see its own header comment:
that script sends no transactions at all).

## Rollback — the failure mode determines the procedure

**These are not interchangeable.** Using the wrong one for the situation either does
nothing (leaves the real problem in place) or wastes time better spent re-upgrading.

| Failure | Procedure |
|---|---|
| Forward *logic* misbehaves post-upgrade with a normally-linked `ForwardLib` (stale fee data, D1 violation, unexpected `ForwardFailed`, etc.) | Remove the bridge address from the executor's target whitelist. Every subsequent finalize with a `bridgeTokenForwarded`-shaped `extraData` falls back through `ForwardFailed(ForwardNotAttempted)` to the existing single-hop payout on chain B. This is the lightest rollback — no re-upgrade needed. **First response for anything that isn't a link/code problem.** |
| **`ForwardLib` link or code failure** (codesize 0, codehash mismatch, the library itself reverting) | **Removing the executor whitelist entry does NOT help.** `_forwardBegin`'s `ForwardLib.setCtx()` call happens **before** `_finalizeBridge`'s executor-whitelist check, so the batch keeps reverting identically whether or not the bridge address is whitelisted — the finalize queue for that source chain stays blocked either way. The only fix is: `setPause(true)` (stop new finalize attempts from failing loudly while you work), then **re-upgrade immediately to the previous, known-good implementation.** Storage layout is unchanged between `CrossBridge`/`CrossBridgeV2` revisions (plan spec §8), so this is a lossless revert. |

If you are not sure which case you are in, **diagnose read-only first, before sending any
transaction:**

1. Re-run step 1's `forge script script/ForwardLibVerify.s.sol --sig 'run(address,address)'`
   against the CURRENTLY-live implementation and its configured `ForwardLib` address. This
   costs nothing and queues nothing — it sends no transaction — and directly distinguishes
   the two cases: the four checks it runs (codesize, codehash, per-relocation address, and
   the wholesale patched-runtime hash) are exactly what a link/code failure trips.
2. If the verifier **reverts**, you are in the second case (link/code failure) —
   `setPause(true)` and re-upgrade to the previous implementation immediately, per the
   table above. Do not attempt a whitelist-removal-and-resubmit first: as the table
   explains, that does nothing for this failure mode and just queues another finalize
   batch that reverts the same way, adding a second failed transaction on top of the
   first.
3. If the verifier **passes**, the link itself is fine and you are in the first case
   (ordinary forward-logic misbehavior) — remove the executor whitelist entry as
   described above; no re-upgrade needed.
4. Only if the read-only diagnostic is inconclusive (e.g. you cannot reproduce the exact
   on-chain implementation/library addresses locally) should you fall back to sending a
   real, small production finalize retry to observe the failure directly — and even then,
   prefer retrying with the whitelist removed first (case 1's lightest rollback) before
   ever re-upgrading, since a link failure will keep reverting regardless of whitelist
   state and a wasted whitelist-removal retry at least costs nothing beyond one gas fee,
   while a wrong re-upgrade costs an extra round trip.

## Indexer / off-chain tooling requirement

`ForwardInitiated` and `ForwardFailed` are **defined in `ForwardLib`**, but because
`CrossBridgeV2` reaches `ForwardLib`'s functions via `DELEGATECALL`, the `emitter`
address on both events is the **bridge proxy address**, not `ForwardLib`'s own address.
Any indexer or off-chain monitoring tool that decodes bridge events by ABI must **merge
`ForwardLib`'s event ABI into the bridge's own ABI** before it can decode these two
events. If this merge is skipped, `ForwardInitiated`/`ForwardFailed` will simply not
decode — the raw log topics/data are still on-chain and correct, but nothing surfaces
them, which silently removes the entire D3 (hop-2 failure) observability story described
in plan spec §10. Confirm this merge is in place as part of any upgrade to
`CrossBridgeV2`, not just once at initial rollout — a tooling redeploy that regenerates
the bridge ABI from `CrossBridgeV2`'s own artifact alone will regress this silently.
