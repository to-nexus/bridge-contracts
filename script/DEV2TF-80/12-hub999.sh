#!/usr/bin/env bash
# [DEV2TF-385] ONE 허브에 체인 999 등록 + gasPrice[999] — tx 파일 하나 (2건) · 서명 1회
#
#   ./12-hub999.sh --check  [--token 0x…]    # 사전 점검만 (조회)
#   ./12-hub999.sh          [--token 0x…]    # tx 빌드 — ① updateGasPriceBatch([999],[1e9])  ② registerToken(999, true, native, TOKEN999)
#   ./12-hub999.sh --verify [--token 0x…]    # 전송 후 확인 (조회)
#
# ★★ ② 가 올라가면 ONE↔999 전송이 «열린다». 999 쪽(09-hype-setup.sh --verify)과 big block OFF 가 끝난 뒤에만 한다.
# ★ 순서: gasPrice 를 «먼저», 페어를 «나중에». 한 파일이어도 원자적이지 않다 — 페어가 먼저 채굴되고 gasPrice 가
#   늦거나 실패하면 수수료 0 경로가 열린다. ① 이 revert 하면 gorge send 가 멈추고 묻는다 → ② 를 보내지 않는다(n).
# ★ 두 호출은 서로 의존하지 않는다 (updateGasPriceBatch 는 체인 등록 여부를 보지 않는다 — BridgeVerifier.sol:569).
# ★ TOKEN999 기본값: chains.env HYPE_TOKEN_FACTORY 의 tokenForPair(612055, 0x…01). --token 으로 직접 줄 수 있다.
# ★ CROSS 는 forge 포크가 불안정해(테스트넷 실측) forge script 가 아니라 gorge write 2건을 한 파일에 담는다.
set -uo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/../gorge/mainnet/_lib.sh"   # 운영 주소·RPC 는 레포 밖(DEV2TF-80 첨부)에 둔다

ACTION="build"; TOKEN=""
while [ $# -gt 0 ]; do
  case "$1" in
    --check)  ACTION="check"; shift ;;
    --verify) ACTION="verify"; shift ;;
    --token)  TOKEN="$2"; shift 2 ;;
    -h|--help) sed -n '2,13p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
    *) die "모르는 옵션: $1" ;;
  esac
done

require_env
NATIVE=0x0000000000000000000000000000000000000001
ZERO=0x0000000000000000000000000000000000000000
GP999=1000000000                       # ONE→999 (DEV2TF-381 결정 8)
FEESIG='calculateFee(uint256,address,uint256)(uint256,uint256,uint256)'
lc(){ printf '%s' "$1" | tr 'A-Z' 'a-z'; }
isaddr(){ printf '%s' "$1" | grep -Eq '^0x[0-9a-fA-F]{40}$'; }
# 조회 실패를 «빈 값» 으로 흘리지 않는다 — 실패하면 그 자리에서 멈춘다
q(){ # rpc to sig [args…]
  local rpc="$1" to="$2" sig="$3" out; shift 3
  out="$(cast call "$to" "$sig" "$@" --rpc-url "$rpc" 2>&1)" || die "조회 실패: $sig @ $to ($rpc)
     $out"
  printf '%s' "$out"
}
qaddr(){ local a; a="$(q "$@" | head -1)"; isaddr "$a" || die "주소가 아니다: $3 → $a"; printf '%s' "$a"; }
# cast 는 큰 수 뒤에 주석([6.12e5])을 붙인다 — 바깥 대괄호만 벗기고 원소마다 첫 단어만 취한다
ids_of(){ q "$1" "$2" 'allChainIDs()(uint256[])' | sed 's/^\[//; s/\]$//' | tr ',' '\n' | awk 'NF{print $1}' | sort -n | paste -sd, -; }
pair(){ # rpc bridge chainId token → "local remote isOrigin paused"
  q "$1" "$2" "getTokenPair(uint256,address)((address,address,bool,bool,uint256,uint256,uint256))" "$3" "$4" \
    | tr -d '()' | awk -F', ' '{print $1, $2, $3, $4}'
}
first(){ printf '%s' "$1" | head -1 | awk '{print $1}'; }

# ── 체인 확인 — 변수 대입만으로는 «어느 체인» 인지 보장되지 않는다 ─────────────
use_chain hype
R999="$R_READ"; P999="${HYPE_BRIDGE:-}"
[ "$(cast chain-id --rpc-url "$R999" 2>/dev/null)" = "999" ] || die "HYPE RPC 가 chainId 999 가 아니다 ($R999)"
use_chain one
for r in "$R_READ" "$R_SEND"; do
  [ "$(cast chain-id --rpc-url "$r" 2>/dev/null)" = "612055" ] || die "ONE RPC 가 chainId 612055 가 아니다 ($r)"
done
ok "chainId — ONE READ·SEND 612055 · HYPE 999"

# ── 999 쪽 값 ───────────────────────────────────────────────────────────
isaddr "$P999" || die "chains.env 의 HYPE_BRIDGE 가 비어 있거나 주소가 아니다 (09-hype-setup.sh --verify 출력의 chains.env 줄을 먼저 옮겨 적을 것)."
if [ -z "$TOKEN" ]; then
  isaddr "${HYPE_TOKEN_FACTORY:-}" || die "HYPE_TOKEN_FACTORY 가 비어 있다. --token 으로 지정하거나 chains.env 를 채워라."
  TOKEN="$(qaddr "$R999" "$HYPE_TOKEN_FACTORY" "tokenForPair(uint256,address)(address)" 612055 "$NATIVE")"
fi
isaddr "$TOKEN" && [ "$(lc "$TOKEN")" != "$ZERO" ] || die "TOKEN999 를 찾지 못했다 ($TOKEN)."
V999="$(qaddr "$R999" "$P999" 'bridgeVerifier()(address)')"
PF999="$(qaddr "$R999" "$V999" 'priceFeed()(address)')"

# ── ONE 쪽 값 — 신 verifier 는 «기록(state.env) 필수» + 온체인 READ·SEND 양쪽과 일치 ──
NEWV="$(state_need ONE_NEW_VERIFIER)"
for r in "$R_READ" "$R_SEND"; do
  v="$(qaddr "$r" "$BRIDGE" 'bridgeVerifier()(address)')"
  [ "$(lc "$v")" = "$(lc "$NEWV")" ] || die "ONE bridgeVerifier($v @ $r) 가 ONE_NEW_VERIFIER($NEWV) 와 다르다 — 4번(D) 상태 확인"
done
PF_ONE="$(qaddr "$R_READ" "$NEWV" 'priceFeed()(address)')"
ok "ONE bridgeVerifier = ONE_NEW_VERIFIER ($NEWV) · PriceFeed $PF_ONE"

if [ "$ACTION" = "verify" ]; then
  hdr "ONE — 전송 후 확인 (CROSS RPC 는 직후 값이 낡을 수 있다 — 영수증 이벤트와 함께 본다)"
  ids="$(ids_of "$R_READ" "$BRIDGE")"
  [ "$ids" = "56,999,1000,8217" ] && gate_ok "ONE allChainIDs = [$ids]" || gate_ng "ONE allChainIDs = [$ids] (기대 56,999,1000,8217)"
  read -r l r o p <<<"$(pair "$R_READ" "$BRIDGE" 999 "$NATIVE")"
  [ "$(lc "$l")" = "$(lc "$NATIVE")" ] && [ "$(lc "$r")" = "$(lc "$TOKEN")" ] && [ "$o" = "true" ] && [ "$p" = "false" ] \
    && gate_ok "ONE 페어 local=native remote=TOKEN999 isOrigin=true paused=false" \
    || gate_ng "ONE 페어가 기대와 다르다: local=$l remote=$r isOrigin=$o paused=$p"
  read -r l r o p <<<"$(pair "$R999" "$P999" 612055 "$TOKEN")"
  [ "$(lc "$l")" = "$(lc "$TOKEN")" ] && [ "$(lc "$r")" = "$(lc "$NATIVE")" ] && [ "$o" = "false" ] && [ "$p" = "false" ] \
    && gate_ok "999 페어와 짝 (999 isOrigin=false · ONE isOrigin=true)" \
    || gate_ng "999 페어와 짝이 안 맞는다: local=$l remote=$r isOrigin=$o paused=$p"
  g="$(first "$(q "$R_READ" "$NEWV" 'getGasPrice(uint256)(uint256)' 999)")"
  [ "$g" = "$GP999" ] && gate_ok "ONE verifier gasPrice[999] = $g" || gate_ng "gasPrice[999] = $g (기대 $GP999)"
  # calculateFee → (minimumValue, networkFee, exFeeRate). ONE→999 는 > 0, 999→ONE 은 gasPrice 0 이라 == 0 이 정상
  f1="$(q "$R_READ" "$NEWV" "$FEESIG" 999 "$NATIVE" 1000000000000000000 | sed -n 2p | awk '{print $1}')"
  f2="$(q "$R999" "$V999" "$FEESIG" 612055 "$TOKEN" 1000000000000000000 | sed -n 2p | awk '{print $1}')"
  [ -n "$f1" ] && [ "$f1" != "0" ] && gate_ok "ONE→999 networkFee = $f1 (> 0)" || gate_ng "ONE→999 networkFee = '$f1' — 0 이면 가격 공급 확인"
  [ "$f2" = "0" ] && gate_ok "999→ONE networkFee = 0 (gasPrice 0 — BSC→ONE 과 같은 패턴)" || gate_ng "999→ONE networkFee = '$f2' (기대 0)"
  gate_end
  hdr "다음"; echo "    ★ 여기서 ONE↔999 전송이 열렸다 — 시각을 기록하고 관계자에게 공유. 이후 DEV2TF-386 · 387"
  exit 0
fi

# ── 사전 점검 ───────────────────────────────────────────────────────────
hdr "사전 점검 — 999 가 준비됐는가 / ONE 권한 / 가격"
info "999 브릿지 $P999 · TOKEN999 $TOKEN · V999 $V999 · PF999 $PF999"
read -r l r o p <<<"$(pair "$R999" "$P999" 612055 "$TOKEN")"
[ "$(lc "$l")" = "$(lc "$TOKEN")" ] && [ "$(lc "$r")" = "$(lc "$NATIVE")" ] && [ "$o" = "false" ] && [ "$p" = "false" ] \
  && gate_ok "999 페어 local=TOKEN999 remote=native isOrigin=false paused=false" \
  || gate_ng "999 페어가 없거나 다르다: local=$l remote=$r isOrigin=$o paused=$p"
ids9="$(ids_of "$R999" "$P999")"
[ "$ids9" = "612055" ] && gate_ok "999 allChainIDs = [612055]" || gate_ng "999 allChainIDs = [$ids9] (기대 612055)"
[ "$(first "$(q "$R999" "$P999" 'paused()(bool)')")" = "false" ] && gate_ok "999 paused=false" || gate_ng "999 브릿지가 paused"
E999="$(qaddr "$R999" "$P999" 'bridgeExecutor()(address)')"
[ "$(lc "$E999")" != "$ZERO" ] && [ "$(cast codesize "$E999" --rpc-url "$R999")" -gt 0 ] \
  && gate_ok "999 executor 연결 ($E999)" || gate_ng "999 executor 미연결 또는 코드 없음 ($E999)"
bb="$(cast rpc eth_usingBigBlocks "$SENDER" --rpc-url "$R999" 2>/dev/null | tr -d '"')"
[ "$bb" = "false" ] && gate_ok "999 big block OFF" || gate_ng "999 big block 상태 '$bb' — 08-bigblock.sh --off 먼저"

# ONE→999 는 ONE→56(BSC) 과 같은 설정으로 연다 (DEV2TF-381 결정 8) — 기준값을 온체인에서 읽어 대조
g56="$(first "$(q "$R_READ" "$NEWV" 'getGasPrice(uint256)(uint256)' 56)")"
[ "$g56" = "$GP999" ] && gate_ok "gasPrice[999] 로 넣을 값 $GP999 = ONE→56 gasPrice[56]" \
  || gate_ng "ONE gasPrice[56]=$g56 ≠ 넣을 값 $GP999 — BSC 경로와 다르다. 값을 확정한 뒤 진행"
require_role "$BRIDGE" "$(cast keccak 'EDITOR_ROLE')" "ONE 브릿지 EDITOR (registerToken)"
require_role "$NEWV"   "$(cast keccak 'PRICER_ROLE')" "ONE verifier PRICER (updateGasPriceBatch)"
ids="$(ids_of "$R_READ" "$BRIDGE")"
[ "$ids" = "56,1000,8217" ] && gate_ok "ONE allChainIDs = [$ids] (999 없음)" \
  || gate_ng "ONE allChainIDs = [$ids] (기대 56,1000,8217) — 이미 등록됐거나 상태가 다르다"
read -r l r o p <<<"$(pair "$R_READ" "$BRIDGE" 999 "$NATIVE")"
[ "$(lc "$l")" = "$ZERO" ] && gate_ok "ONE 에 (999, native) 페어 없음" || gate_ng "ONE 에 (999, native) 페어가 이미 있다: local=$l remote=$r"

# ★ 가격 3종 — 하나라도 없으면 tx 는 성공하지만 수수료·최소금액이 틀어진다
#   ONE→999 networkFee = CalcAmountLib.calculateTokenAmountForNetworkFee: 목적지(999) 네이티브 가격 «과» 출발 토큰(ONE) 가격 둘 다 필요
p1="$(first "$(q "$R_READ" "$PF_ONE" 'getNativeTokenPrice(uint256)(bool,uint256,uint256)' 999)")"
[ "$p1" = "true" ] && gate_ok "ONE PriceFeed — 999 네이티브(HYPE) 가격 있음" \
  || gate_ng "ONE PriceFeed($PF_ONE) 에 999 네이티브 가격이 없다 → ONE→999 networkFee 0.
       PRICER(C) 가 updateNativeTokenPrice([999],[HYPE 가격],[시각]) 을 먼저 넣어야 한다"
p2="$(first "$(q "$R_READ" "$PF_ONE" 'getTokenPriceInDollars(address)(bool,uint256,uint256)' "$NATIVE")")"
[ "$p2" = "true" ] && gate_ok "ONE PriceFeed — ONE(native) 가격 있음" \
  || gate_ng "ONE PriceFeed($PF_ONE) 에 ONE(native) 가격이 없다 → ONE→999 networkFee 0"
p3="$(first "$(q "$R999" "$PF999" 'getTokenPriceInDollars(address)(bool,uint256,uint256)' "$TOKEN")")"
[ "$p3" = "true" ] && gate_ok "999 PriceFeed — TOKEN999 가격 있음" \
  || gate_ng "999 PriceFeed($PF999) 에 TOKEN999 가격이 없다 → 999→ONE 최소금액이 defaultTokenPrice(10000 = 0.01 USD) 기준 500 ONE.
       PRICER 가 updatePrice([TOKEN999],[ONE 가격],[시각]) 을 먼저 넣어야 한다"
gate_end
[ "$ACTION" = "check" ] && exit 0

# ── 빌드 (한 파일 2건 — ① gasPrice 먼저 ② 페어 나중) ─────────────────────────
nl="$(cast nonce "$SENDER" --block latest  --rpc-url "$R_SEND")"
np="$(cast nonce "$SENDER" --block pending --rpc-url "$R_SEND")"
[ "$nl" = "$np" ] || die "ONE latest($nl) ≠ pending($np) — 대기 중인 tx 가 있다. 정리 후 다시."

hdr "ONE — ① updateGasPriceBatch([999],[$GP999])  ② registerToken(999, true, native, TOKEN999)"
F="$WORK/tx-$PFX-H1.json"; rm -f "$F"
gorge_run write $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --out "$F" \
  "$NEWV" "updateGasPriceBatch(uint256[],uint256[])" "[999]" "[$GP999]" \
  || die "gorge write 실패: updateGasPriceBatch (위 gorge 출력 참고)"
gorge_run write $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --out "$F" --append --final \
  "$BRIDGE" "registerToken(uint256,bool,address,address)" 999 true "$NATIVE" "$TOKEN" \
  || die "gorge write 실패: registerToken (위 gorge 출력 참고)"

hdr "tx 파일 구조 점검"
C1="$(cast calldata 'updateGasPriceBatch(uint256[],uint256[])' "[999]" "[$GP999]")"
C2="$(cast calldata 'registerToken(uint256,bool,address,address)' 999 true "$NATIVE" "$TOKEN")"
python3 - "$F" "$nl" "$SENDER" "$NEWV" "$C1" "$BRIDGE" "$C2" <<'PY' || die "tx 파일 구조가 예상과 다르다 — 서명하지 마라."
import json,sys
f,n0,sender,to1,c1,to2,c2=sys.argv[1],int(sys.argv[2]),*[a.lower() for a in sys.argv[3:]]
num=lambda v: int(str(v),16) if str(v).startswith('0x') else int(v)
d=json.load(open(f)); txs=d['transactions']; bad=[]
if len(txs)!=2: bad.append(f"건수 {len(txs)} != 2")
want=[(to1,c1),(to2,c2)]
for i,t in enumerate(txs[:2]):
    x=t['transaction']
    if str(x.get('from','')).lower()!=sender: bad.append(f"[{i}] from {x.get('from')}")
    if num(x['nonce'])!=n0+i: bad.append(f"[{i}] nonce {x['nonce']} != {n0+i}")
    if str(x.get('to','')).lower()!=want[i][0]: bad.append(f"[{i}] to {x.get('to')} != {want[i][0]}")
    data=str(x.get('input') or x.get('data') or '').lower()
    if data!=want[i][1]: bad.append(f"[{i}] calldata 불일치")
    if num(x.get('value') or 0)!=0: bad.append(f"[{i}] value != 0")
    cid=x.get('chainId', d.get('chainId'))
    if cid is not None and num(cid)!=612055: bad.append(f"[{i}] chainId {cid}")
if num(d.get('chainId'))!=612055: bad.append(f"chainId {d.get('chainId')}")
print("    건수", len(txs), "· nonce", n0, "~", n0+len(txs)-1, "· ① gasPrice → ② registerToken")
if bad: print("\n".join("    ✘ "+b for b in bad)); sys.exit(1)
PY
ok "건수 2 · 발신자 · 연속 nonce · 대상 · calldata · value 0 · chainId 612055"

finish_build "$F"
cat <<EOF

  ★ ① 이 revert 하면 gorge send 가 «계속할지» 묻는다 → n (② 페어를 열지 않는다).
  ★ 전송 뒤:  $0 --verify
EOF
