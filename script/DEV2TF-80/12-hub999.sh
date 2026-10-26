#!/usr/bin/env bash
# [DEV2TF-385] ONE 허브에 체인 999 등록 + gasPrice[999] — tx 파일 하나 (2건) · 서명 1회
#
#   ./12-hub999.sh --check  [--token 0x…]    # 사전 점검만 (조회)
#   ./12-hub999.sh          [--token 0x…]    # tx 빌드 — registerToken(999, true, native, TOKEN999) + updateGasPriceBatch([999],[1e9])
#   ./12-hub999.sh --verify [--token 0x…]    # 전송 후 확인 (조회)
#
# ★★ 이 tx 가 올라가면 ONE↔999 전송이 «열린다». 999 쪽(script/DEV2TF-80/09-hype-setup.sh --verify)이 끝난 뒤에만 한다.
# ★ TOKEN999 기본값: chains.env HYPE_TOKEN_FACTORY 의 tokenForPair(612055, 0x…01). --token 으로 직접 줄 수 있다.
# ★ CROSS 는 forge 포크가 불안정해(테스트넷 실측) forge script 가 아니라 gorge write 2건을 한 파일에 담는다.
#   두 호출은 서로 의존하지 않는다 (updateGasPriceBatch 는 체인 등록 여부를 보지 않는다 — BridgeVerifier.sol:569).
set -uo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/../gorge/mainnet/_lib.sh"   # 운영 주소·RPC 는 레포 밖(DEV2TF-80 첨부)에 둔다

ACTION="build"; TOKEN=""
while [ $# -gt 0 ]; do
  case "$1" in
    --check)  ACTION="check"; shift ;;
    --verify) ACTION="verify"; shift ;;
    --token)  TOKEN="$2"; shift 2 ;;
    -h|--help) sed -n '2,11p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
    *) die "모르는 옵션: $1" ;;
  esac
done

require_env
NATIVE=0x0000000000000000000000000000000000000001
GP999=1000000000                       # ONE→999 (DEV2TF-381 결정 8)
lc(){ printf '%s' "$1" | tr 'A-Z' 'a-z'; }

# ── 999 쪽 값 ───────────────────────────────────────────────────────────
use_chain hype
R999="$R_READ"; P999="${HYPE_BRIDGE:-}"
[ -n "$P999" ] || die "chains.env 의 HYPE_BRIDGE 가 비어 있다 (script/DEV2TF-80/09-hype-setup.sh --verify 출력의 chains.env 줄을 먼저 옮겨 적을 것)."
if [ -z "$TOKEN" ]; then
  [ -n "${HYPE_TOKEN_FACTORY:-}" ] || die "HYPE_TOKEN_FACTORY 가 비어 있다. --token 으로 지정하거나 chains.env 를 채워라."
  TOKEN="$(cast call "$HYPE_TOKEN_FACTORY" "tokenForPair(uint256,address)(address)" 612055 "$NATIVE" --rpc-url "$R999")"
fi
[ "$(lc "$TOKEN")" != "0x0000000000000000000000000000000000000000" ] && [ -n "$TOKEN" ] || die "TOKEN999 를 찾지 못했다."
V999="$(rd "$R999" "$P999" 'bridgeVerifier()(address)')"
PF999="$(rd "$R999" "$V999" 'priceFeed()(address)')"

pair(){ # rpc bridge chainId token → "local remote isOrigin paused"
  cast call "$2" "getTokenPair(uint256,address)((address,address,bool,bool,uint256,uint256,uint256))" "$3" "$4" \
    --rpc-url "$1" 2>/dev/null | tr -d '()' | awk -F', ' '{print $1, $2, $3, $4}'
}

# ── ONE 쪽 값 ───────────────────────────────────────────────────────────
use_chain one
# 기준은 «온체인 bridgeVerifier()» 다 (4번 D 에서 이미 신 verifier 로 바뀌었다). state.env 가 있으면 대조만 한다.
NEWV="$(rd "$R_READ" "$BRIDGE" 'bridgeVerifier()(address)')"
[ -n "$NEWV" ] || die "ONE bridgeVerifier() 조회 실패"
SV="$(state_get ONE_NEW_VERIFIER)"
[ -z "$SV" ] || [ "$(lc "$SV")" = "$(lc "$NEWV")" ] \
  || die "ONE bridgeVerifier($NEWV) 가 state.env 의 ONE_NEW_VERIFIER($SV) 와 다르다 — 확인 후 진행"
PF_ONE="$(rd "$R_READ" "$NEWV" 'priceFeed()(address)')"

if [ "$ACTION" = "verify" ]; then
  hdr "ONE — 전송 후 확인 (CROSS RPC 는 직후 값이 낡을 수 있다 — 영수증 이벤트와 함께 본다)"
  ids="$(cast call "$BRIDGE" 'allChainIDs()(uint256[])' --rpc-url "$R_READ")"
  printf '%s' "$ids" | tr -d '[] ' | tr ',' '\n' | awk '{print $1}' | grep -qx 999 \
    && gate_ok "allChainIDs 에 999 ($ids)" || gate_ng "allChainIDs 에 999 가 없다 ($ids)"
  read -r l r o p <<<"$(pair "$R_READ" "$BRIDGE" 999 "$NATIVE")"
  [ "$(lc "$l")" = "$(lc "$NATIVE")" ] && [ "$(lc "$r")" = "$(lc "$TOKEN")" ] && [ "$o" = "true" ] && [ "$p" = "false" ] \
    && gate_ok "ONE 페어 local=native remote=$TOKEN isOrigin=true paused=false" \
    || gate_ng "ONE 페어가 기대와 다르다: local=$l remote=$r isOrigin=$o paused=$p"
  read -r l r o p <<<"$(pair "$R999" "$P999" 612055 "$TOKEN")"
  [ "$o" = "false" ] && [ "$(lc "$r")" = "$(lc "$NATIVE")" ] \
    && gate_ok "999 페어와 짝 (999 isOrigin=false · ONE isOrigin=true)" || gate_ng "999 페어와 짝이 안 맞는다 (999 isOrigin=$o remote=$r)"
  g="$(rd "$R_READ" "$NEWV" 'getGasPrice(uint256)(uint256)' 999 | awk '{print $1}')"
  [ "$g" = "$GP999" ] && gate_ok "ONE verifier gasPrice[999] = $g" || gate_ng "gasPrice[999] = $g (기대 $GP999)"
  FEESIG='calculateFee(uint256,address,uint256)(uint256,uint256,uint256)'
  info "ONE→999 calculateFee(1 ONE): $(cast call "$NEWV" "$FEESIG" 999 "$NATIVE" 1000000000000000000 --rpc-url "$R_READ" 2>&1 | tr '\n' ' ')"
  info "999→ONE calculateFee(1 ONE): $(cast call "$V999" "$FEESIG" 612055 "$TOKEN" 1000000000000000000 --rpc-url "$R999" 2>&1 | tr '\n' ' ')"
  info "  (networkFee 는 두 번째 값. 999→ONE 은 gasPrice 0 이라 0 이 정상 — BSC→ONE 과 같은 패턴)"
  gate_end
  hdr "다음"; echo "    ★ 여기서 ONE↔999 전송이 열렸다 — 시각을 기록하고 관계자에게 공유. 이후 DEV2TF-386 · 387"
  exit 0
fi

# ── 사전 점검 ───────────────────────────────────────────────────────────
hdr "사전 점검 — 999 가 준비됐는가 / ONE 권한 / 가격"
info "999 브릿지 $P999 · TOKEN999 $TOKEN · V999 $V999"
read -r l r o p <<<"$(pair "$R999" "$P999" 612055 "$TOKEN")"
[ "$(lc "$l")" = "$(lc "$TOKEN")" ] && [ "$(lc "$r")" = "$(lc "$NATIVE")" ] && [ "$o" = "false" ] && [ "$p" = "false" ] \
  && gate_ok "999 페어 local=TOKEN999 remote=native isOrigin=false paused=false" \
  || gate_ng "999 페어가 없거나 다르다: local=$l remote=$r isOrigin=$o paused=$p"
[ "$(rd "$R999" "$P999" 'paused()(bool)')" = "false" ] && gate_ok "999 paused=false" || gate_ng "999 브릿지가 paused"
[ "$(lc "$(rd "$R999" "$P999" 'bridgeExecutor()(address)')")" != "0x0000000000000000000000000000000000000000" ] \
  && gate_ok "999 executor 연결" || gate_ng "999 executor 미연결"
bb="$(cast rpc eth_usingBigBlocks "$SENDER" --rpc-url "$R999" 2>/dev/null | tr -d '"')"
[ "$bb" = "false" ] && gate_ok "999 big block OFF" || gate_ng "999 big block 이 아직 켜져 있다 ($bb) — 08-bigblock.sh --off 먼저"

info "ONE verifier $NEWV · PriceFeed $PF_ONE"
require_role "$BRIDGE" "$(cast keccak 'EDITOR_ROLE')" "ONE 브릿지 EDITOR (registerToken)"
require_role "$NEWV"   "$(cast keccak 'PRICER_ROLE')" "ONE verifier PRICER (updateGasPriceBatch)"
ids="$(cast call "$BRIDGE" 'allChainIDs()(uint256[])' --rpc-url "$R_READ")"
printf '%s' "$ids" | tr -d '[] ' | tr ',' '\n' | awk '{print $1}' | grep -qx 999 \
  && gate_ng "ONE allChainIDs 에 999 가 이미 있다 ($ids) — 이미 등록됐다" || gate_ok "ONE allChainIDs $ids (999 없음)"

# ★ 가격 — 없으면 tx 는 성공하지만 수수료·최소금액이 틀어진다 (CalcAmountLib.calculateTokenAmountForNetworkFee)
np="$(cast call "$PF_ONE" 'getNativeTokenPrice(uint256)(bool,uint256,uint256)' 999 --rpc-url "$R_READ" | head -1)"
[ "$np" = "true" ] && gate_ok "ONE PriceFeed 에 999 네이티브(HYPE) 가격 있음" \
  || gate_ng "ONE PriceFeed($PF_ONE) 에 999 네이티브 가격이 없다 → ONE→999 networkFee 가 0 이 된다.
       PRICER 가 updateNativeTokenPrice([999],[HYPE 가격],[시각]) 을 먼저 넣어야 한다"
tp="$(cast call "$PF999" 'getTokenPriceInDollars(address)(bool,uint256,uint256)' "$TOKEN" --rpc-url "$R999" | head -1)"
[ "$tp" = "true" ] && gate_ok "999 PriceFeed 에 TOKEN999 가격 있음" \
  || gate_ng "999 PriceFeed($PF999) 에 TOKEN999 가격이 없다 → 999→ONE 최소금액이 defaultTokenPrice(10000 = 0.01 USD) 기준 500 ONE 이 된다.
       PRICER 가 updatePrice([TOKEN999],[ONE 가격],[시각]) 을 먼저 넣어야 한다"
gate_end
[ "$ACTION" = "check" ] && exit 0

# ── 빌드 (한 파일 2건) ───────────────────────────────────────────────────
hdr "ONE — registerToken(999, true, native, TOKEN999) + updateGasPriceBatch([999],[$GP999])"
F="$WORK/tx-$PFX-H1.json"; rm -f "$F"
gorge_run write $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --out "$F" \
  "$BRIDGE" "registerToken(uint256,bool,address,address)" 999 true "$NATIVE" "$TOKEN" \
  || die "gorge write 실패: registerToken (위 gorge 출력 참고)"
gorge_run write $(gc_opt) --rpc-url "$R_SEND" --sender "$SENDER" --out "$F" --append --final \
  "$NEWV" "updateGasPriceBatch(uint256[],uint256[])" "[999]" "[$GP999]" \
  || die "gorge write 실패: updateGasPriceBatch (위 gorge 출력 참고)"
finish_build "$F"
echo "      전송 뒤:  $0 --verify"
