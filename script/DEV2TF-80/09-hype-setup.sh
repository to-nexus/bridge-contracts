#!/usr/bin/env bash
# [DEV2TF-382 · 383 · 384 · 435 STEP 1·2] HyperEVM 999 구축을 tx 파일 «하나»로 — 서명 1회
#
#   ./09-hype-setup.sh --check                                      # 사전 점검만 (조회)
#   ./09-hype-setup.sh --mode all --name "<이름>" --symbol "<심볼>"  # 9·10·11 + 435(토큰·999 페어) — 16건
#   ./09-hype-setup.sh --mode infra                                 # 9·10·11 만 — 14건
#   ./09-hype-setup.sh --verify --bridge 0x… [--token 0x…]          # 전송 후 온체인 재검증 (조회)
#
# 내용은 HypeSetup.s.sol 에 있다. gorge script 가 forge 시뮬레이션(브로드캐스트 없음)으로 tx 를 모으고,
# 시뮬레이션 끝의 require(역할·설정·주소·링크 가능 상태·바이트코드)가 하나라도 틀리면 파일이 «안 만들어진다».
#
# ★ big block 은 이 파일 밖이다 (L1 액션). 빌드 전에 ON, 전송·검증 후 08-bigblock.sh --off.
# ★ 주소는 A 의 nonce 로 정해진다. 빌드 → 전송 사이에 A 로 다른 tx 를 보내지 않는다. 보냈으면 다시 빌드.
# ★ 12번(ONE 허브 등록)은 체인이 달라 따로 — 12-hub999.sh
set -uo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/../gorge/mainnet/_lib.sh"   # 운영 주소·RPC 는 레포 밖(DEV2TF-80 첨부)에 둔다

MODE=""; NAME=""; SYMBOL=""; ACTION="build"; VBRIDGE=""; VTOKEN=""
while [ $# -gt 0 ]; do
  case "$1" in
    --mode)   MODE="$2"; shift 2 ;;
    --name)   NAME="$2"; shift 2 ;;
    --symbol) SYMBOL="$2"; shift 2 ;;
    --check)  ACTION="check"; shift ;;
    --verify) ACTION="verify"; shift ;;
    --bridge) VBRIDGE="$2"; shift 2 ;;
    --token)  VTOKEN="$2"; shift 2 ;;
    -h|--help) sed -n '2,15p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
    *) die "모르는 옵션: $1" ;;
  esac
done

require_env; use_chain hype
SCRIPT="script/DEV2TF-80/HypeSetup.s.sol:HypeSetup"
lc(){ printf '%s' "$1" | tr 'A-Z' 'a-z'; }

[ -z "${FOUNDRY_LIBRARIES:-}" ] || die "FOUNDRY_LIBRARIES 가 설정돼 있다 ($FOUNDRY_LIBRARIES).  unset FOUNDRY_LIBRARIES"
cd "$REPO" || die "레포 루트로 이동 실패"

# ── HypeSetup.s.sol 에 넘길 환경 — 운영 주소는 레포에 두지 않고 chains.env 에서만 읽는다 ──
for k in HYPE_BIMPL HYPE_BIMPL_CODEHASH HYPE_PROLE_ADMIN HYPE_PROLE_PRICER \
         HYPE_BROLE_DEFAULT_ADMIN HYPE_BROLE_ADMIN HYPE_BROLE_VALIDATOR HYPE_BROLE_EDITOR HYPE_BROLE_OPERATOR HYPE_BROLE_VERIFIER \
         HYPE_VROLE_DEFAULT_ADMIN HYPE_VROLE_ADMIN HYPE_VROLE_EDITOR HYPE_VROLE_PRICER \
         HYPE_EROLE_DEFAULT_ADMIN HYPE_EROLE_ADMIN; do
  eval "v=\${$k-__UNSET__}"
  [ "$v" != "__UNSET__" ] && [ -n "$v" ] || die "chains.env 에 $k 가 없다 (2026-10-26 이후 첨부의 chains.env 가 필요하다)."
  export "$k=$v"
done
eval "v=\${HYPE_BROLE_INITIATOR-__UNSET__}"
[ "$v" != "__UNSET__" ] || die "chains.env 에 HYPE_BROLE_INITIATOR 가 없다 (빈 값으로 «있어야» 한다)."
export HYPE_BROLE_INITIATOR="$v"
# dev 는 ONE · BSC 브릿지와 같은 수취 주소를 쓴다 — 두 체인에서 읽어 일치할 때만 넘긴다
DEV_ONE="$(cast call "$ONE_BRIDGE" 'dev()(address)' --rpc-url "$ONE_RPC_READ" 2>/dev/null)"
DEV_BSC="$(cast call "$BSC_BRIDGE" 'dev()(address)' --rpc-url "$BSC_RPC_READ" 2>/dev/null)"
[ -n "$DEV_ONE" ] && [ "$(lc "$DEV_ONE")" = "$(lc "$DEV_BSC")" ] \
  || die "ONE dev($DEV_ONE) 와 BSC dev($DEV_BSC) 가 다르거나 조회 실패 — 999 dev 를 정할 수 없다."
export HYPE_DEV="$DEV_ONE" SENDER

# ── 검증 (전송 후) ───────────────────────────────────────────────────────
if [ "$ACTION" = "verify" ]; then
  [ -n "$VBRIDGE" ] || die "--bridge <999 브릿지 프록시> 를 지정해라."
  hdr "$CH_NAME — 전송 후 온체인 재검증 (조회만 · tx 없음)"
  out="$(forge script "$SCRIPT" --sig "verify(address,address)" "$VBRIDGE" "${VTOKEN:-0x0000000000000000000000000000000000000000}" \
        --rpc-url "$R_READ" 2>&1)"; rc=$?
  printf '%s\n' "$out" | sed -n '/== Logs ==/,$p' | sed 's/^/    /'
  if [ $rc -ne 0 ] || ! printf '%s\n' "$out" | grep -q "VERIFY OK"; then
    printf '%s\n' "$out" | grep -E "Error|revert|Revert" | tail -5 | sed 's/^/      /'
    die "검증 실패 — 위 메시지의 항목을 먼저 확인할 것."
  fi
  ok "역할 · 설정 · 연결 · 링크 가능 상태 · 런타임 바이트코드 전부 일치"
  pick(){ printf '%s\n' "$out" | grep -E "^[[:space:]]+$1[[:space:]]+=" | head -1 | awk '{print $NF}'; }
  state_put HYPE_NEW_VERIFIER "$(pick 'BridgeVerifier \(V999\)')"
  state_put HYPE_NEW_EXECUTOR "$(pick 'BridgeExecutor \(E999\)')"
  [ -n "$VTOKEN" ] && state_put HYPE_TOKEN999 "$VTOKEN"
  hdr "다음"
  cat <<EOF
    1) 위 «chains.env» 줄을 script/gorge/mainnet/chains.env 에 옮겨 적는다 → DEV2TF-80 첨부 갱신
    2) big block OFF
         ./script/gorge/mainnet/08-bigblock.sh --off --prepare   →  gorgef sign712  →  --send
         ./script/gorge/mainnet/08-bigblock.sh --status           # false
    3) ./script/gorge/mainnet/01-preflight.sh --chain hype
    4) 12번 — ./script/DEV2TF-80/12-hub999.sh --check
EOF
  exit 0
fi

# ── 사전 점검 (빌드 전 · --check) ────────────────────────────────────────
hdr "$CH_NAME — 사전 점검"
[ "$(lc "$SENDER")" = "$(lc "$HYPE_BROLE_DEFAULT_ADMIN")" ] \
  || die "SENDER($SENDER) 가 HYPE_BROLE_DEFAULT_ADMIN($HYPE_BROLE_DEFAULT_ADMIN) 이 아니다 — 운영 키 A 로 실행해라."
ok "SENDER = 운영 키 A · dev = $HYPE_DEV (ONE · BSC 와 동일)"
cid="$(cast chain-id --rpc-url "$R_SEND" 2>/dev/null)"
[ "$cid" = "999" ] && ok "chainId 999" || die "chainId 가 999 가 아니다 ($cid) — RPC 확인 (사무실/DevOps IP 전용)"

role="$(curl -s -m 20 -X POST -H 'Content-Type: application/json' \
  --data "{\"type\":\"userRole\",\"user\":\"$SENDER\"}" "$HYPE_API/info")"
case "$role" in *'"user"'*) ok "HyperCore 계정 활성 ($role)" ;; *) die "HyperCore 계정이 아니다 ($role) — big block 을 켤 수 없다" ;; esac

bb="$(cast rpc eth_usingBigBlocks "$SENDER" --rpc-url "$R_SEND" 2>/dev/null | tr -d '"')"
[ "$bb" = "true" ] && ok "usingBigBlocks=true" \
  || die "big block 이 꺼져 있다 ($bb). 3M 를 넘는 배포 3건이 들어가지 않는다.
     ./script/gorge/mainnet/08-bigblock.sh --on --prepare  →  gorgef sign712  →  --send"

nl="$(cast nonce "$SENDER" --block latest  --rpc-url "$R_SEND")"
np="$(cast nonce "$SENDER" --block pending --rpc-url "$R_SEND")"
[ "$nl" = "$np" ] && ok "nonce $nl (latest = pending, 대기 tx 없음)" \
  || die "latest($nl) ≠ pending($np) — 대기 중인 tx 가 있다. 정리 후 다시."

bal="$(cast balance "$SENDER" --rpc-url "$R_SEND")"
python3 -c "import sys; sys.exit(0 if int('$bal') >= 10**16 else 1)" \
  && ok "잔고 $(cast to-unit "$bal" ether) HYPE (예상 소모 약 0.004 HYPE)" \
  || die "잔고 부족 — $(cast to-unit "$bal" ether) HYPE (0.01 HYPE 이상 필요)"

ch="$(cast keccak "$(cast code "$HYPE_BIMPL" --rpc-url "$R_SEND")")"
[ "$(lc "$ch")" = "$(lc "$HYPE_BIMPL_CODEHASH")" ] && ok "BaseBridge 구현 $HYPE_BIMPL = 0번 검수본 (codehash)" \
  || die "BaseBridge 구현 codehash 불일치 ($ch ≠ HYPE_BIMPL_CODEHASH)"

[ -z "${HYPE_BRIDGE:-}" ] || warn "chains.env 의 HYPE_BRIDGE 가 이미 차 있다 ($HYPE_BRIDGE). 이미 배포했다면 다시 보내지 않는다."

[ "$ACTION" = "check" ] && { ok "사전 점검 통과"; exit 0; }

# ── 빌드 ────────────────────────────────────────────────────────────────
case "$MODE" in
  all)   [ -n "$NAME" ] && [ -n "$SYMBOL" ] || die "--mode all 은 --name 과 --symbol 이 필요하다."
         SIG="all(string,string)"; ARGS=("$NAME" "$SYMBOL"); EXPECT=16 ;;
  infra) SIG="infra()"; ARGS=(); EXPECT=14 ;;
  *)     die "--mode all|infra 를 지정해라." ;;
esac

hdr "$CH_NAME — gorge script ($SIG)"
[ "$MODE" = "all" ] && info "토큰 이름 «$NAME» · 심볼 «$SYMBOL» · decimals 18 · minter = 브릿지  (생성 후 바꿀 수 없다)"
F="$WORK/tx-$PFX-S.json"; rm -f "$F"
gorge_run script $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --profile deploy --out "$F" \
  "$SCRIPT" "$SIG" ${ARGS[@]+"${ARGS[@]}"} \
  || die "gorge script 실패 — 위 출력의 require 메시지가 원인이다. tx 파일은 만들어지지 않았다."
printf '%s\n' "$GORGE_OUT" | grep -q "VERIFY OK" || die "시뮬레이션 검증(VERIFY OK) 이 출력되지 않았다 — 서명하지 마라."

hdr "tx 파일 구조 점검"
python3 - "$F" "$EXPECT" "$nl" "$SENDER" <<'PY' || die "tx 파일 구조가 예상과 다르다 — 서명하지 마라."
import json,sys
f,exp,n0,sender=sys.argv[1],int(sys.argv[2]),int(sys.argv[3]),sys.argv[4].lower()
d=json.load(open(f)); txs=d['transactions']; bad=[]
if len(txs)!=exp: bad.append(f"건수 {len(txs)} != {exp}")
for i,t in enumerate(txs):
    x=t['transaction']
    if str(x.get('from','')).lower()!=sender: bad.append(f"[{i}] from {x.get('from')}")
    if int(str(x['nonce']),16 if str(x['nonce']).startswith('0x') else 10)!=n0+i: bad.append(f"[{i}] nonce {x['nonce']} != {n0+i}")
    g=int(str(x['gas']),16 if str(x['gas']).startswith('0x') else 10)
    if g>30_000_000: bad.append(f"[{i}] gas {g} > big block 30M")
cid=d.get('chainId'); cid=int(str(cid),16) if str(cid).startswith('0x') else int(cid)
if cid!=999: bad.append(f"chainId {cid}")
print("    건수", len(txs), "· nonce", n0, "~", n0+len(txs)-1, "· chainId", cid)
if bad: print("\n".join("    ✘ "+b for b in bad)); sys.exit(1)
PY
ok "건수 · 발신자 · 연속 nonce · chainId · 가스 한도(≤30M) 확인"

finish_build "$F"
cat <<EOF

  ★ 서명 화면에서 tx 마다 «Confirm signing» 을 묻는다 — 위 표와 대조하며 y.
  ★ 전송 뒤 (같은 출력의 P999 · TOKEN999 주소로):
      $0 --verify --bridge <P999>$( [ "$MODE" = "all" ] && printf ' --token <TOKEN999>' )
EOF
