// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {Const} from "../../src/lib/Const.sol";
import {CrossBridgeV2TempSetRemoteToken} from "../../src/temp/CrossBridgeV2TempSetRemoteToken.sol";

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {Upgrades} from "lib/openzeppelin-foundry-upgrades/src/Upgrades.sol";

import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossBridgeV2Temp
 * @notice ⚠️ TEMPORARY / TESTNET-ONLY / ONE-SHOT ⚠️ Deployment + runbook script for
 * `CrossBridgeV2TempSetRemoteToken` (plan spec §7 items 8-9, §15). Lives ONLY on branch
 * `temp/cross-bridge-setremotetoken`; never merged into `dev`.
 *
 * See the 한글 실행 런북 at the bottom of this file for the full step-by-step procedure.
 */
contract CrossBridgeV2Temp is Script {
    /// @notice HyperEVM testnet chain ID — the only remote chain this fix targets.
    uint private constant HYPER_CHAIN_ID = 998;

    /// @notice CROSS chain ID (plan spec §6 "온체인 상수"). Used by `preflightDrained` when
    /// querying the HyperEVM-side registry for CROSS's lane counters.
    uint private constant CROSS_CHAIN_ID = 612044;

    /// @notice The CROSS bridge proxy (plan spec §6). Fixed as a constant — this is a
    /// one-shot script, so an operator-supplied address is only an unnecessary typo
    /// surface (round-2 issue plan H1/M1).
    address private constant CROSS_BRIDGE = 0xb81d6e000000000000000000000000000000C0de;

    /// @notice The HyperEVM bridge (plan spec §6), used by `preflightDrained` to read the
    /// HyperEVM side directly instead of trusting operator-supplied values (round-2 issue
    /// plan H1).
    address private constant HYPER_BRIDGE = 0xeF7F79dA2281704647306267E7E1C72848768c22;

    /// @notice The implementation to restore to after the fix — the currently-deployed
    /// `CrossBridgeV2` (plan spec §6/§15). No new build needed to restore: this address
    /// is already deployed and correctly linked against `FORWARD_LIB`.
    address internal constant ORIGINAL_IMPL = 0xe52bdBd35e0Db7b2efbaF17628ed0CD9D2bA8919;

    /// @notice `ForwardLib` address `CrossBridgeV2`/the temp impl must be linked against
    /// (plan spec §6/§7-10).
    address internal constant FORWARD_LIB = 0x2d0c06B1D999fDA0D7a51906A3cAA94518e8Fe0f;

    /// @notice The known current `remoteToken` value this run is expected to replace
    /// (plan spec §6). Fixed as a constant — an operator-supplied `expectedOld` would let
    /// the fail-closed check be defeated by accidentally passing whatever the current
    /// value happens to be (round-2 issue plan M1).
    /// @dev UPDATED for the second remap round. The first round moved the pair from
    /// `0x3E0217c3926106b7E585B5341439f3150c6cab7c` to the value below; this round moves
    /// it off that value onto the new beacon-proxy token. This constant MUST match the
    /// pair's live `remoteToken` at execution time or `preflightRemoteToken` fails closed.
    address private constant EXPECTED_OLD_REMOTE_TOKEN = 0xBB37E106055AbAF98e4f039B6d000bF50646a410;

    /// @notice Thrown by `preflightDrained` when CROSS's next-initiate index for 998
    /// does not match HyperEVM's next-finalize index for 612044 (an in-flight
    /// CROSS→Hyper bridge would still carry the OLD `remoteToken` in its signed payload).
    error NotDrainedCrossToHyper(uint crossInitiate, uint hyperFinalize);

    /// @notice Thrown by `preflightDrained` when HyperEVM's next-initiate index for
    /// 612044 does not match CROSS's next-finalize index for 998.
    error NotDrainedHyperToCross(uint hyperInitiate, uint crossFinalize);

    /// @notice Thrown by `preflightDrained` when either side has non-empty pending
    /// operations for the 998/612044 lane.
    error PendingNotEmpty(uint crossPending, uint hyperPending);

    /// @notice Thrown by `preflightRemoteToken` when the CURRENT `remoteToken` is
    /// neither the expected stuck old value nor the already-corrected `T_new` value —
    /// an unexplained third value means something else touched this pair; do not proceed
    /// blindly.
    error PreflightRemoteTokenUnexpectedValue(address current, address expectedOld, address tNew);

    /// @notice Thrown by `preflightRemoteToken` when `tNew` itself is unusable — zero, or
    /// equal to `EXPECTED_OLD_REMOTE_TOKEN` (i.e. not actually a NEW value). Checked
    /// BEFORE reading the current on-chain value so a bad `tNew` can never fall through
    /// into the idempotent-skip branch and silently defeat fail-closed (round-2 issue
    /// plan M1).
    error InvalidNewRemoteToken(address tNew);

    function setUp() public {}

    // ============================================================
    // Deploy
    // ============================================================

    /**
     * @notice Deploys `CrossBridgeV2TempSetRemoteToken`.
     * @dev Must be built with `ForwardLib` linked to `FORWARD_LIB` (see the runbook's
     * build command below) — mirrors `CrossBridgeV2Deploy.s.sol:deployCrossBridgeV2Impl`.
     * Run `ForwardLibVerify.s.sol` against the returned address BEFORE calling
     * `switchTo`.
     * @return impl The deployed temp implementation address.
     */
    function deployTemp() public returns (address impl) {
        vm.broadcast();
        CrossBridgeV2TempSetRemoteToken implementation = new CrossBridgeV2TempSetRemoteToken();
        impl = address(implementation);

        console.log("CrossBridgeV2TempSetRemoteToken deployed to:", impl);
        console.log("  runtime size:", impl.code.length);
        console.log("");
        console.log("Next: run ForwardLibVerify.s.sol against this impl BEFORE switchTo().");
        console.log("  forge script script/ForwardLibVerify.s.sol \\");
        console.log("    --sig 'run(address,address)' %s %s \\", impl, FORWARD_LIB);
        console.log("    --rpc-url <rpc>");
    }

    // ============================================================
    // Observation (read-only, no broadcast)
    // ============================================================

    /**
     * @notice Dumps the on-chain observations an operator should compare across the
     * transition window (plan spec §15 steps 0/6/10 "스냅샷"/"소진 재검증"/"comparator
     * 실행"): the target pair's full 7-field struct, the surrounding registry counters
     * for the 998 lane, `crossSupplyLimit`, the proxy's current implementation, and the
     * bridge's native balance. Call this at each checkpoint in the runbook and diff the
     * output by eye — this function only observes; it makes no pass/fail judgement
     * (`preflightDrained`/`preflightRemoteToken` do that for the two conditions that
     * actually gate progression).
     * @param proxy The CROSS bridge proxy address.
     */
    function snapshot(address proxy) public view {
        IBridgeRegistry registry = IBridgeRegistry(proxy);
        IBridgeRegistry.TokenPair memory pair = registry.getTokenPair(HYPER_CHAIN_ID, Const.NATIVE_TOKEN);

        console.log("=== CrossBridgeV2Temp snapshot ===");
        console.log("block.number:", block.number);
        console.log("block.timestamp:", block.timestamp);
        console.log("proxy:", proxy);
        console.log("current implementation:", Upgrades.getImplementationAddress(proxy));
        console.log("");
        console.log("-- target pair (998, NATIVE_TOKEN) --");
        console.log("localToken:", pair.localToken);
        console.log("remoteToken:", pair.remoteToken);
        console.log("isOrigin:", pair.isOrigin);
        console.log("paused:", pair.paused);
        console.log("pendingAmount:", pair.pendingAmount);
        console.log("deposited:", pair.deposited);
        console.log("minted:", pair.minted);
        console.log("");
        console.log("-- 998 lane counters --");
        console.log("getNextInitiateIndex(998):", registry.getNextInitiateIndex(HYPER_CHAIN_ID));
        console.log("getNextFinalizeIndex(998):", registry.getNextFinalizeIndex(HYPER_CHAIN_ID));
        console.log("allPendingIndex(998).length:", registry.allPendingIndex(HYPER_CHAIN_ID).length);
        console.log("");
        console.log("-- global --");
        console.log("native balance:", proxy.balance);
        console.log("===================================");
    }

    // ============================================================
    // Preflight gates
    // ============================================================

    /**
     * @notice Verifies the cross-drain condition (plan spec §15 "교차 소진 검증") BEFORE
     * entering the transition window: no in-flight CROSS<->HyperEVM bridge exists on the
     * 998/612044 lane. Reads BOTH chains itself via `vm.createFork`/`vm.selectFork` — the
     * six values are never operator-supplied, since a hand-entered value (stale copy-paste,
     * typo) could be made to pass the gate (round-2 issue plan H1). Reverts on the first
     * failing condition; logs all six either way.
     * @param crossRpc RPC URL for the CROSS chain.
     * @param hyperRpc RPC URL for the HyperEVM chain.
     */
    function preflightDrained(string memory crossRpc, string memory hyperRpc) public {
        uint crossFork = vm.createFork(crossRpc);
        uint hyperFork = vm.createFork(hyperRpc);

        vm.selectFork(crossFork);
        IBridgeRegistry crossRegistry = IBridgeRegistry(CROSS_BRIDGE);
        uint crossInit = crossRegistry.getNextInitiateIndex(HYPER_CHAIN_ID);
        uint crossFin = crossRegistry.getNextFinalizeIndex(HYPER_CHAIN_ID);
        uint crossPending = crossRegistry.allPendingIndex(HYPER_CHAIN_ID).length;

        vm.selectFork(hyperFork);
        IBridgeRegistry hyperRegistry = IBridgeRegistry(HYPER_BRIDGE);
        uint hyperInit = hyperRegistry.getNextInitiateIndex(CROSS_CHAIN_ID);
        uint hyperFin = hyperRegistry.getNextFinalizeIndex(CROSS_CHAIN_ID);
        uint hyperPending = hyperRegistry.allPendingIndex(CROSS_CHAIN_ID).length;

        console.log("=== preflightDrained ===");
        console.log("crossInit:", crossInit);
        console.log("hyperFin:", hyperFin);
        console.log("hyperInit:", hyperInit);
        console.log("crossFin:", crossFin);
        console.log("crossPending:", crossPending);
        console.log("hyperPending:", hyperPending);

        _assertDrained(crossInit, hyperFin, hyperInit, crossFin, crossPending, hyperPending);

        console.log("DRAINED OK");
    }

    /**
     * @notice Pure comparison logic for `preflightDrained`, split out so it is testable
     * without a dual-RPC fork setup (round-2 issue plan H1). Exposed to tests via a
     * harness contract that inherits `CrossBridgeV2Temp` (mirrors
     * `ForwardLibVerifyHarness is ForwardLibVerify` in `test/ForwardLibVerify.t.sol`).
     * @param crossInit `CROSS.getNextInitiateIndex(998)`.
     * @param hyperFin `HYPE.getNextFinalizeIndex(612044)`.
     * @param hyperInit `HYPE.getNextInitiateIndex(612044)`.
     * @param crossFin `CROSS.getNextFinalizeIndex(998)`.
     * @param crossPending `CROSS.allPendingIndex(998).length`.
     * @param hyperPending `HYPE.allPendingIndex(612044).length`.
     */
    function _assertDrained(
        uint crossInit,
        uint hyperFin,
        uint hyperInit,
        uint crossFin,
        uint crossPending,
        uint hyperPending
    ) internal pure {
        if (crossInit != hyperFin) revert NotDrainedCrossToHyper(crossInit, hyperFin);
        if (hyperInit != crossFin) revert NotDrainedHyperToCross(hyperInit, crossFin);
        if (crossPending != 0 || hyperPending != 0) revert PendingNotEmpty(crossPending, hyperPending);
    }

    /**
     * @notice 3-way branch on the CURRENT `remoteToken` value (plan spec §10 error
     * table): proceed if it is still the expected stuck old value, skip (idempotent, no
     * revert) if it is already `tNew`, or revert if it is a third, unexplained value.
     * @dev `tNew` is validated BEFORE the current on-chain value is even read (round-2
     * issue plan M1) — if this check ran after, `tNew == 0` or `tNew ==
     * EXPECTED_OLD_REMOTE_TOKEN` could fall into the idempotent-skip branch (`current ==
     * tNew`) and silently defeat fail-closed instead of reverting.
     * @param tNew The corrected HyperEVM-side token address this run intends to set.
     * @return shouldProceed True if the caller should go on to call `setRemoteToken`;
     * false if the value is already correct (idempotent skip, plan spec §10/§12).
     */
    function preflightRemoteToken(address tNew) public view returns (bool shouldProceed) {
        if (tNew == address(0) || tNew == EXPECTED_OLD_REMOTE_TOKEN) revert InvalidNewRemoteToken(tNew);

        address current = IBridgeRegistry(CROSS_BRIDGE).getTokenPair(HYPER_CHAIN_ID, Const.NATIVE_TOKEN).remoteToken;

        console.log("=== preflightRemoteToken ===");
        console.log("current remoteToken:", current);
        console.log("expectedOld:        ", EXPECTED_OLD_REMOTE_TOKEN);
        console.log("tNew:                ", tNew);

        if (current == tNew) {
            console.log("preflightRemoteToken: already T_new - idempotent skip, no call needed");
            return false;
        }

        if (current != EXPECTED_OLD_REMOTE_TOKEN) {
            revert PreflightRemoteTokenUnexpectedValue(current, EXPECTED_OLD_REMOTE_TOKEN, tNew);
        }

        console.log("preflightRemoteToken: OK - current == expectedOld, safe to call setRemoteToken(tNew)");
        return true;
    }

    // ============================================================
    // Switch / restore
    // ============================================================

    /**
     * @notice Upgrades the proxy to the temp implementation. Run `ForwardLibVerify` and
     * `preflightDrained`/`preflightRemoteToken` BEFORE calling this.
     * @param proxy The CROSS bridge proxy address.
     * @param tempImpl The `CrossBridgeV2TempSetRemoteToken` address from `deployTemp()`.
     */
    function switchTo(address proxy, address tempImpl) public {
        console.log("Current implementation:", Upgrades.getImplementationAddress(proxy));
        console.log("Switching to temp implementation:", tempImpl);

        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(tempImpl, "");

        console.log("Switched implementation:", Upgrades.getImplementationAddress(proxy));
    }

    /**
     * @notice Restores the proxy to `ORIGINAL_IMPL`. No new build needed — that
     * implementation is already deployed and linked. This is the highest-priority
     * recovery step (plan spec §15 "9 원복 실패 → 최우선 사고"): retry immediately on
     * failure.
     * @param proxy The CROSS bridge proxy address.
     */
    function restore(address proxy) public {
        console.log("Current implementation:", Upgrades.getImplementationAddress(proxy));
        console.log("Restoring to original implementation:", ORIGINAL_IMPL);

        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(ORIGINAL_IMPL, "");

        address restored = Upgrades.getImplementationAddress(proxy);
        console.log("Restored implementation:", restored);
        require(restored == ORIGINAL_IMPL, "restore: implementation slot did not land on ORIGINAL_IMPL");
    }
}

/**
 * ============================================================================
 * 한글 실행 런북 (plan spec §15)
 * ============================================================================
 *
 * 전제: 없음. 회계 사전 정리를 하지 않는다 (plan spec §11-1). 필요한 전제는 1 단계의
 * 교차 소진 검증 통과뿐이다.
 *
 * 온체인 상수 (plan spec §6):
 *   CROSS 브리지 프록시    : 0xb81d6e000000000000000000000000000000C0de
 *   원복 대상 impl         : 0xe52bdBd35e0Db7b2efbaF17628ed0CD9D2bA8919  (ORIGINAL_IMPL)
 *   ForwardLib             : 0x2d0c06B1D999fDA0D7a51906A3cAA94518e8Fe0f  (FORWARD_LIB)
 *   HyperEVM 브리지        : 0xeF7F79dA2281704647306267E7E1C72848768c22
 *   CROSS / HyperEVM chainId: 612044 / 998
 *
 * | # | 체인    | 동작                                              | 권한   |
 * |---|---------|---------------------------------------------------|--------|
 * | 0 | 양쪽    | 사전 스냅샷 (`snapshot`, HyperEVM 쪽은 별도 조회)  | —      |
 * | 1 | 양쪽    | 교차 소진 검증 (`preflightDrained`)                | —      |
 * | 2 | HyperEVM| `HyperMintableERC20Code` 배포                      | —      |
 * | 3 | HyperEVM| `setCrossMintableERC20Code(신규 팩토리)`           | ADMIN  |
 * | 4 | HyperEVM| `createToken(612044, 0x…01, <SYM>, 18)` (big block)| EDITOR |
 * | 5 | CROSS   | 임시 impl 배포(`deployTemp`) + `ForwardLibVerify`  | —      |
 * | 6 | 양쪽    | 소진 재검증 (`preflightDrained`) — 창 시작         | —      |
 * | 7 | CROSS   | `switchTo(proxy, tempImpl)`                        | ADMIN  |
 * | 8 | CROSS   | `preflightRemoteToken` 분기 후 `setRemoteToken`    | ADMIN  |
 * | 9 | CROSS   | `restore(proxy)`                                   | ADMIN  |
 * | 10| 양쪽    | `snapshot` 재실행 (Group A 불변 확인)              | —      |
 * | 11| 양쪽    | 양방향 브리지 e2e — CROSS→Hyper 가 T_new 를 민팅   | —      |
 *
 * 6~9 는 한 세션에 연속 실행 (목표 2~3분), 그 동안 브리지 트랜잭션을 보내지 않는다.
 *
 * 배포 명령 (5단계). `forge script` 는 자체적으로 컴파일·링크하므로, 배포 프로파일과
 * `--libraries` 는 이 명령 자체에 붙여야 한다 — 별도로 `forge build --libraries` 를 먼저
 * 돌려도 뒤따르는 `forge script` 가 같은 프로파일·같은 링크로 만들어진다는 보장이 없다
 * (round-2 issue plan M2). 배포 프로파일은 메타데이터 해시를 제외해 `--libraries` 설정
 * 변화가 바이트코드에 영향을 주지 않게 한다:
 *
 *   FOUNDRY_PROFILE=deploy forge script script/temp/CrossBridgeV2Temp.s.sol:CrossBridgeV2Temp \
 *     --sig "deployTemp()" --rpc-url <CROSS_RPC_URL> --broadcast \
 *     --libraries src/lib/ForwardLib.sol:ForwardLib:0x2d0c06b1d999fda0d7a51906a3caa94518e8fe0f
 *
 * (`switchTo`/`restore` 는 이미 배포되어 링크가 끝난 impl 주소를 프록시에 연결만 하므로
 * 프로파일·`--libraries` 가 필요 없다 — 새 바이트코드를 만들지 않기 때문이다.)
 *
 * 링크 무결성 검증 (필수 게이트, 5단계 직후 — 읽기 전용, 트랜잭션 없음):
 *
 *   forge script script/ForwardLibVerify.s.sol \
 *     --sig 'run(address,address)' <TEMP_IMPL_ADDR> 0x2d0c06b1d999fda0d7a51906a3caa94518e8fe0f \
 *     --rpc-url <CROSS_RPC_URL>
 *
 * 원복 (9단계, 실패 시 최우선 사고 — 자금 위험은 없으나 즉시 재시도):
 *
 *   forge script script/temp/CrossBridgeV2Temp.s.sol:CrossBridgeV2Temp \
 *     --sig "restore(address)" 0xb81d6e000000000000000000000000000000C0de \
 *     --rpc-url <CROSS_RPC_URL> --broadcast
 *
 * 롤백 대응표 (plan spec §15):
 *   1·6 소진 실패      : 미처리 건이 정상 finalize 되기를 기다렸다 재검증 (pause 를 안
 *                        걸었으므로 finalize 는 정상 진행)
 *   5 빌드/검증 실패   : 프록시 미변경. 중단
 *   7 업그레이드 실패  : impl 슬롯 확인 → 원본이면 중단
 *   8 실패             : 9 로 즉시 원복 후 분석
 *   8 후 값 오류       : 임시 impl 이 붙어 있으면 재호출로 교정
 *   9 원복 실패        : 최우선 사고. 재시도. 자금 위험은 없으나 원복까지 계속 시도
 *   창 중 예상 밖 트래픽: 선택적으로 양쪽 `setChainPause` 사용 가능 (필수 아님)
 */
