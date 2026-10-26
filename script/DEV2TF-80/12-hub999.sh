#!/usr/bin/env bash
# [DEV2TF-385] ONE 허브에 체인 999 등록 + gasPrice[999] — ONE→999 를 ONE→56(BSC) 과 «완전히 같은» 설정으로 연다
#
#   ./12-hub999.sh --check       [--token 0x…]   # 사전 점검만 (조회) — 다음에 할 단계를 알려 준다
#   ./12-hub999.sh --step gas    [--token 0x…]   # tx ① updateGasPriceBatch([999],[gasPrice[56]])  → 서명 · 전송
#   ./12-hub999.sh --step pair   [--token 0x…]   # ① 이 온체인에 반영된 것을 확인한 뒤 tx ② registerToken(999, true, native, TOKEN999)
#   ./12-hub999.sh --verify      [--token 0x…]   # 전송 후 확인 (조회)
#
# ★★ ② 가 올라가면 ONE↔999 전송이 «열린다». 999 쪽(09-hype-setup.sh --verify)과 big block OFF 가 끝난 뒤에만 한다.
# ★ 왜 두 파일인가: gorge send 는 영수증 조회가 시간 초과되면 «다음 tx 를 계속 보낸다». 한 파일에 담으면 ① 이 실패했는데
#   ② 가 성공해 수수료 0 경로가 열릴 수 있다. 그래서 ① 의 «온체인 반영» 을 확인해야만 ② 파일이 만들어진다.
# ★ BSC 와 «완전히» 같게: gasPrice[999] = gasPrice[56], ONE PriceFeed native[999] = native[56],
#   그리고 calculateFee(999, ONE) == calculateFee(56, ONE) 를 강제한다. native[999] 는 PRICER(C)가 먼저 넣는다.
# ★ TOKEN999 기본값: chains.env HYPE_TOKEN_FACTORY 의 tokenForPair(612055, 0x…01). --token 으로 직접 줄 수 있다.
# ★ CROSS 는 forge 포크가 불안정해(테스트넷 실측) forge script 가 아니라 gorge write 로 만든다.
set -uo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/../gorge/mainnet/_lib.sh"   # 운영 주소·RPC 는 레포 밖(DEV2TF-80 첨부)에 둔다

ACTION="check"; STEP=""; TOKEN=""
while [ $# -gt 0 ]; do
  case "$1" in
    --check)  ACTION="check"; shift ;;
    --verify) ACTION="verify"; shift ;;
    --step)   ACTION="build"; STEP="$2"; shift 2 ;;
    --token)  TOKEN="$2"; shift 2 ;;
    -h|--help) sed -n '2,15p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
    *) die "모르는 옵션: $1" ;;
  esac
done
case "$ACTION:$STEP" in check:|verify:|build:gas|build:pair) ;; *) die "--step gas|pair 를 지정해라." ;; esac

require_env
NATIVE=0x0000000000000000000000000000000000000001
ZERO=0x0000000000000000000000000000000000000000
ONE_ETH=1000000000000000000
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
first(){ printf '%s' "$1" | head -1 | awk '{print $1}'; }
nth(){ printf '%s' "$1" | sed -n "${2}p" | awk '{print $1}'; }   # 여러 줄 반환의 n 번째 값 (cast 주석 [1e9] 제거)
# cast 는 큰 수 뒤에 주석([6.12e5])을 붙인다 — 바깥 대괄호만 벗기고 원소마다 첫 단어만 취한다
ids_of(){ q "$1" "$2" 'allChainIDs()(uint256[])' | sed 's/^\[//; s/\]$//' | tr ',' '\n' | awk 'NF{print $1}' | sort -n | paste -sd, -; }
pair(){ # rpc bridge chainId token → "local remote isOrigin paused"
  q "$1" "$2" "getTokenPair(uint256,address)((address,address,bool,bool,uint256,uint256,uint256))" "$3" "$4" \
    | tr -d '()' | awk -F', ' '{print $1, $2, $3, $4}'
}
fee3(){ q "$@" | awk 'NF{printf "%s%s", (n++?",":""), $1}'; }   # calculateFee → "min,networkFee,exFee"

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

# ── BSC 경로(ONE→56) 기준값 — 999 경로를 이것과 «완전히» 같게 연다 ────────────
GP999="$(first "$(q "$R_READ" "$NEWV" 'getGasPrice(uint256)(uint256)' 56)")"
NP56="$(nth "$(q "$R_READ" "$PF_ONE" 'getNativeTokenPrice(uint256)(bool,uint256,uint256)' 56)" 2)"
FEE56="$(fee3 "$R_READ" "$NEWV" "$FEESIG" 56 "$NATIVE" "$ONE_ETH")"
printf '%s' "$GP999" | grep -Eq '^[1-9][0-9]*$' || die "ONE gasPrice[56] 조회값이 이상하다 ('$GP999')"
info "기준(ONE→56) — gasPrice[56] $GP999 · native[56] $NP56 · calculateFee(56, 1 ONE) = $FEE56"

gas999(){ first "$(q "$1" "$NEWV" 'getGasPrice(uint256)(uint256)' 999)"; }
# gasPrice[999] 상태 — READ·SEND 둘 다 반영(set) / 둘 다 미반영(unset) / 한쪽만(syncing: RPC 동기화 대기)
gas_state(){
  local a b; a="$(gas999 "$R_READ")"; b="$(gas999 "$R_SEND")"
  if   [ "$a" = "$GP999" ] && [ "$b" = "$GP999" ]; then echo set
  elif [ "$a" != "$GP999" ] && [ "$b" != "$GP999" ]; then echo unset
  else echo syncing; fi
}

# ★ 가격 — BSC 경로와 «같은 값» 이어야 한다. 가격이 없으면 계약이 실패를 버리고 networkFee 0 을 쓰므로
#   (CalcAmountLib.sol:38 · BridgeVerifier.sol:546) «수수료가 같다» 만으로는 정상을 증명하지 못한다 → 가격 자체를 본다.
price_gate(){
  local np56 np999 tp
  np56="$(q "$R_READ" "$PF_ONE" 'getNativeTokenPrice(uint256)(bool,uint256,uint256)' 56)"
  [ "$(first "$np56")" = "true" ] && gate_ok "ONE PriceFeed native[56] = $NP56 (기준)" || gate_ng "ONE PriceFeed native[56] 가 없다 — 기준값을 정할 수 없다"
  np999="$(q "$R_READ" "$PF_ONE" 'getNativeTokenPrice(uint256)(bool,uint256,uint256)' 999)"
  [ "$(first "$np999")" = "true" ] && [ "$(nth "$np999" 2)" = "$NP56" ] \
    && gate_ok "ONE PriceFeed native[999] = native[56] = $NP56" \
    || gate_ng "ONE PriceFeed($PF_ONE) native[999] = '$(printf '%s' "$np999" | tr '\n' ' ')' ≠ native[56] $NP56.
       PRICER(C) 가 updateNativeTokenPrice([999],[$NP56],[0]) 을 먼저 넣어야 한다 (BSC 경로와 같은 값)"
  [ "$(first "$(q "$R_READ" "$PF_ONE" 'getTokenPriceInDollars(address)(bool,uint256,uint256)' "$NATIVE")")" = "true" ] \
    && gate_ok "ONE PriceFeed — ONE(native) 가격 있음" || gate_ng "ONE PriceFeed($PF_ONE) 에 ONE(native) 가격이 없다"
  tp="$(q "$R999" "$PF999" 'getTokenPriceInDollars(address)(bool,uint256,uint256)' "$TOKEN")"
  [ "$(first "$tp")" = "true" ] && [ "$(nth "$tp" 2)" = "1000000" ] && gate_ok "999 PriceFeed — TOKEN999 가격 1.00 USD (BSC ONE 토큰과 동일)" \
    || gate_ng "999 PriceFeed($PF999) TOKEN999 가격이 1.00 USD 가 아니다 ('$(printf '%s' "$tp" | tr '\n' ' ')') — 09 묶음 결과 확인"
  [ "$(printf '%s' "$FEE56" | cut -d, -f2)" != "0" ] && gate_ok "기준 ONE→56 networkFee > 0 ($FEE56)" || gate_ng "기준 ONE→56 networkFee 가 0 이다 ($FEE56) — 기준 경로 자체가 이상하다"
}

# ── 전송 후 확인 ─────────────────────────────────────────────────────────
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
  g="$(gas999 "$R_READ")"
  [ "$g" = "$GP999" ] && gate_ok "ONE gasPrice[999] = $g (= gasPrice[56])" || gate_ng "gasPrice[999] = $g (기대 $GP999)"
  price_gate
  f999="$(fee3 "$R_READ" "$NEWV" "$FEESIG" 999 "$NATIVE" "$ONE_ETH")"
  [ "$(printf '%s' "$f999" | cut -d, -f2)" != "0" ] && gate_ok "ONE→999 networkFee > 0" || gate_ng "ONE→999 networkFee 가 0 이다 ($f999)"
  [ "$f999" = "$FEE56" ] && gate_ok "calculateFee(999, 1 ONE) = $f999 = calculateFee(56) — BSC 경로와 동일" \
    || gate_ng "calculateFee(999)=$f999 ≠ calculateFee(56)=$FEE56"
  f2="$(fee3 "$R999" "$V999" "$FEESIG" 612055 "$TOKEN" "$ONE_ETH")"
  info "999→ONE calculateFee(1 ONE) = $f2  (networkFee 0 — BSC→ONE 과 같은 패턴)"
  [ "$(printf '%s' "$f2" | cut -d, -f2)" = "0" ] && gate_ok "999→ONE networkFee = 0" || gate_ng "999→ONE networkFee ≠ 0 ($f2)"
  gate_end
  hdr "다음"; echo "    ★ 여기서 ONE↔999 전송이 열렸다 — 시각을 기록하고 관계자에게 공유. 이후 DEV2TF-386 · 387"
  exit 0
fi

# ── 사전 점검 (--check · 두 단계 모두) ──────────────────────────────────────
hdr "사전 점검 — 999 가 준비됐는가 / ONE 권한 / 가격 (BSC 경로와 동일한가)"
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

require_role "$BRIDGE" "$(cast keccak 'EDITOR_ROLE')" "ONE 브릿지 EDITOR (registerToken)"
require_role "$NEWV"   "$(cast keccak 'PRICER_ROLE')" "ONE verifier PRICER (updateGasPriceBatch)"
ids="$(ids_of "$R_READ" "$BRIDGE")"
[ "$ids" = "56,1000,8217" ] && gate_ok "ONE allChainIDs = [$ids] (999 없음)" \
  || gate_ng "ONE allChainIDs = [$ids] (기대 56,1000,8217) — 이미 등록됐거나 상태가 다르다"
read -r l r o p <<<"$(pair "$R_READ" "$BRIDGE" 999 "$NATIVE")"
[ "$(lc "$l")" = "$ZERO" ] && gate_ok "ONE 에 (999, native) 페어 없음" || gate_ng "ONE 에 (999, native) 페어가 이미 있다: local=$l remote=$r"

price_gate

GS="$(gas_state)"; g_now="$(gas999 "$R_READ")"; g_send="$(gas999 "$R_SEND")"
info "현재 ONE gasPrice[999] — READ $g_now · SEND $g_send (넣을 값 $GP999 = gasPrice[56]) → $GS"
gate_end

if [ "$ACTION" = "check" ]; then
  hdr "다음"
  case "$GS" in
    set)     echo "    ① gasPrice 반영 확인됨 (READ·SEND) → $0 --step pair" ;;
    unset)   echo "    $0 --step gas" ;;
    syncing) echo "    ① 이 한쪽 RPC 에만 보인다 — RPC 동기화 대기. 잠시 뒤 $0 --check 를 다시 (gas 를 다시 보내지 않는다)" ;;
  esac
  exit 0
fi

nl="$(cast nonce "$SENDER" --block latest  --rpc-url "$R_SEND")"
np="$(cast nonce "$SENDER" --block pending --rpc-url "$R_SEND")"
[ "$nl" = "$np" ] || die "ONE latest($nl) ≠ pending($np) — 대기 중인 tx 가 있다. 정리 후 다시."

if [ "$STEP" = "gas" ]; then
  case "$GS" in
    set)     die "gasPrice[999] 는 이미 $GP999 다 — --step pair 로 간다." ;;
    syncing) die "gasPrice[999] 가 한쪽 RPC 에만 반영됐다 (READ $g_now · SEND $g_send) — 동기화 대기 후 --check. 다시 보내지 않는다." ;;
  esac
  TO="$NEWV"; SIG='updateGasPriceBatch(uint256[],uint256[])'; ARGS=("[999]" "[$GP999]"); F="$WORK/tx-$PFX-H1.json"
  hdr "ONE — ① updateGasPriceBatch([999],[$GP999])   (아직 전송 경로는 열리지 않는다)"
else
  # ① 이 «온체인에 반영» 됐는지 READ·SEND 양쪽에서 본다 — 영수증만으로는 부족하다 (gorge 는 시간 초과 시 다음으로 넘어간다)
  [ "$GS" = "set" ] \
    || die "gasPrice[999] 가 READ·SEND 양쪽에 $GP999 로 보이지 않는다 (READ $g_now · SEND $g_send · $GS) — ① 전송·반영 확인 후 다시."
  f999="$(fee3 "$R_READ" "$NEWV" "$FEESIG" 999 "$NATIVE" "$ONE_ETH")"
  [ "$(printf '%s' "$f999" | cut -d, -f2)" != "0" ] || die "ONE→999 networkFee 가 0 이다 ($f999) — 가격 확인"
  [ "$f999" = "$FEE56" ] || die "calculateFee(999)=$f999 ≠ calculateFee(56)=$FEE56 — BSC 경로와 다르다. 가격·gasPrice 확인"
  ok "calculateFee(999, 1 ONE) = $f999 = calculateFee(56) — 페어를 열어도 된다"
  TO="$BRIDGE"; SIG='registerToken(uint256,bool,address,address)'; ARGS=(999 true "$NATIVE" "$TOKEN"); F="$WORK/tx-$PFX-H2.json"
  hdr "ONE — ② registerToken(999, true, native, TOKEN999)   ★ 이 tx 로 ONE↔999 가 열린다"
fi

rm -f "$F"
gorge_run write $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --out "$F" --final "$TO" "$SIG" "${ARGS[@]}" \
  || die "gorge write 실패: $SIG (위 gorge 출력 참고)"

hdr "tx 파일 구조 점검"
CD="$(cast calldata "$SIG" "${ARGS[@]}")"
python3 - "$F" "$nl" "$SENDER" "$TO" "$CD" <<'PY' || die "tx 파일 구조가 예상과 다르다 — 서명하지 마라."
import json,sys
f,n0,sender,to,cd=sys.argv[1],int(sys.argv[2]),*[a.lower() for a in sys.argv[3:]]
num=lambda v: int(str(v),16) if str(v).startswith('0x') else int(v)
d=json.load(open(f)); txs=d['transactions']; bad=[]
if len(txs)!=1: bad.append(f"건수 {len(txs)} != 1")
for i,t in enumerate(txs[:1]):
    x=t['transaction']
    if str(x.get('from','')).lower()!=sender: bad.append(f"from {x.get('from')}")
    if num(x['nonce'])!=n0: bad.append(f"nonce {x['nonce']} != {n0}")
    if str(x.get('to','')).lower()!=to: bad.append(f"to {x.get('to')} != {to}")
    if str(x.get('input') or '').lower()!=cd: bad.append("calldata 불일치")
    if num(x.get('value') or 0)!=0: bad.append("value != 0")
    if x.get('chainId') is not None and num(x['chainId'])!=612055: bad.append(f"tx chainId {x['chainId']}")
if num(d.get('chainId'))!=612055: bad.append(f"chainId {d.get('chainId')}")
print("    건수", len(txs), "· nonce", n0)
if bad: print("\n".join("    ✘ "+b for b in bad)); sys.exit(1)
PY
ok "건수 1 · 발신자 · nonce · 대상 · calldata · value 0 · chainId 612055"

finish_build "$F"
if [ "$STEP" = "gas" ]; then
  echo "      전송·영수증 확인 뒤:  $0 --check   (gasPrice 반영 확인)  →  $0 --step pair"
else
  echo "      전송 뒤:  $0 --verify"
fi
