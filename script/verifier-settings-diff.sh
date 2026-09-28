#!/usr/bin/env bash
# 구 / 신 BridgeVerifier 의 설정을 슬롯 원시값으로 전수 대조한다. 읽기 전용.
#
# BridgeVerifier 는 프록시가 아니라 생성자 배포다. 재배포하면 설정이 자동으로
# 따라오지 않으므로, 브릿지를 갈아끼우기 전에 "빠뜨린 설정이 없는가" 를 증명해야 한다.
# 함수 반환값 비교(calculateFee 등)는 커버리지를 스스로 증명하지 못하므로 쓰지 않는다.
#
# 실행하면(직접 호출) 지금까지와 동일하게 동작한다. source 하면 vsd_cmp / mslot 등
# 함수만 정의되고 아무것도 실행되지 않는다 — script/test/verifier-settings-diff.test.sh 가
# 이 방식으로 비교 로직을 직접 호출해서 검증한다.
set -euo pipefail

DIFFS=0
ERRORS=0

usage() {
  cat <<'USAGE'
BridgeVerifier 설정 전수 대조 (읽기 전용)

  ./script/verifier-settings-diff.sh \
      --rpc-url <RPC> --bridge <브릿지프록시> --old <구 베리파이어> --new <새 베리파이어>

옵션
  --rpc-url <URL>      대상 체인 RPC                                   (필수)
  --bridge  <주소>     브릿지 프록시 — 등록 체인·토큰을 여기서 열거한다  (필수)
  --old     <주소>     현재 쓰고 있는 BridgeVerifier                    (필수)
  --new     <주소>     새로 배포한 BridgeVerifier                       (필수)
  --token   <주소>     추가로 비교할 토큰 (여러 번 지정 가능)            (선택)

대조 항목
  슬롯 1~8   스칼라 설정 8개
  슬롯 10    _gasPrice — 브릿지에 등록된 체인 전부
  슬롯 9     _exFeeRate — 등록된 모든 체인의 토큰 쌍 양쪽 전부
  슬롯 16    _minimumTokenValueOf — 토큰별 최소 전송금액 오버라이드 (같은 토큰 집합)
  슬롯 13    역할 12종의 멤버 목록
  슬롯 11/12 _tokenCurrentVolume · _tokenMovementHistory  → [주의] 로만 표시
             (런타임 누적치라 옮길 수 없다. 임계값이 0 이면 영향 없음)

종료 코드
  0  불일치 없음 (조회 실패도 없음)   1  불일치 있음 / 조회 실패 있음 / 인자 오류
USAGE
}

# sto / rolemembers / dec 는 실패 시 반드시 non-zero 를 반환한다 (set -o pipefail 이
# 파이프 왼쪽 명령의 실패를 오른쪽까지 전파한다). 빈 슬롯의 정상 값(0x000…0)과 조회
# 실패를 내용으로 구분하지 않는다 — 성공/실패는 오직 종료 상태로만 구분한다. 호출부는
# `if ! o="$(sto …)"; then vsd_err …; continue; fi` 형태로 실패를 명시적으로 처리한다.
sto() { cast storage "$1" "$2" --rpc-url "$RPC" 2>/dev/null | tail -1; }
rolemembers() { cast call "$1" "getRoleMembers(bytes32)(address[])" "$2" --rpc-url "$RPC" 2>/dev/null | head -1 | tr 'A-Z' 'a-z'; }
dec() { cast to-dec "$1" 2>/dev/null; }
mslot() { # key(32B hex, 0x 없이) baseSlot -> keccak(key . base)
  cast keccak "0x$1$(printf '%064x' "$2")"
}

# vsd_abi_array <raw> <words-per-record> — cast 가 돌려준 동적 배열 응답(예:
# allTokenPairs)을 파싱 전에 구조부터 검증한다. 통과하면 헤더 2워드(offset·length)를
# 뗀 레코드 본문만 stdout 으로 돌려주고 0 을 반환한다. 실패하면 stdout 은 비우고
# 사유를 stderr 로 낸 뒤 1 을 반환한다. 호출부는
#   if body="$(vsd_abi_array "$raw" 7 2>/dev/null)"; then …
#   else reason="$(vsd_abi_array "$raw" 7 2>&1 1>/dev/null)"; vsd_err … "$reason"; fi
# 형태로 성공/실패를 명시적으로 나눠 처리한다 — vsd_cmp/sto 와 같은 규율이다.
#
# 검사 순서: hex 문자만 → 64(워드) 배수 길이 → 최소 2워드(헤더) → offset==0x20 →
# 선언된 length 가 실제 워드 수와 정합.
#
# ★ 마지막 검사에서 선언된 length 워드를 bash 정수로 변환하지 않는다. malformed
# 응답은 거대한 uint256 length 를 선언할 수 있고, 그걸 $((16#…)) 로 바꾸면 bash 의
# signed 64비트 산술이 오버플로해 운 나쁘면 우연히 일치해버릴 수 있다. 대신 방향을
# 뒤집는다 — 실제 응답의 총 워드 수(작고 신뢰 가능한 수)에서 기대 레코드 개수를
# 역산하고, 그 값만 32바이트 hex 워드로 인코딩해 선언된 length 워드와 문자열로
# 비교한다. 신뢰할 수 없는 값은 끝까지 산술에 넣지 않는다.
vsd_abi_array() {
  local raw="$1" words_per_record="$2"
  raw="${raw#0x}"
  if ! [[ "$raw" =~ ^[0-9a-fA-F]*$ ]]; then
    echo "비-hex 응답" >&2
    return 1
  fi
  raw="$(printf '%s' "$raw" | tr 'A-Z' 'a-z')"
  local total_chars=${#raw}
  if [ $((total_chars % 64)) -ne 0 ]; then
    echo "워드 경계 불일치 (len=${total_chars})" >&2
    return 1
  fi
  local total_words=$((total_chars / 64))
  if [ "$total_words" -lt 2 ]; then
    echo "헤더 없음 (len=${total_chars}, 최소 128 필요)" >&2
    return 1
  fi
  local offset_word="${raw:0:64}"
  local expect_offset
  expect_offset="$(printf '%064x' 32)"
  if [ "$offset_word" != "$expect_offset" ]; then
    echo "offset 이상 (offset=0x${offset_word})" >&2
    return 1
  fi
  local length_word="${raw:64:64}"
  local record_words=$((total_words - 2))
  if [ $((record_words % words_per_record)) -ne 0 ]; then
    echo "워드 경계 불일치 (레코드 워드 수=${record_words}, words-per-record=${words_per_record})" >&2
    return 1
  fi
  local expected=$((record_words / words_per_record))
  local expect_length_word
  expect_length_word="$(printf '%064x' "$expected")"
  if [ "$length_word" != "$expect_length_word" ]; then
    echo "길이 불일치 (선언=0x${length_word}, 실제 레코드 수로 역산한 기대값=${expected})" >&2
    return 1
  fi
  printf '%s' "${raw:128}"
  return 0
}

# vsd_err <label> <사유> — 조회/변환 실패를 [ERR ] 로 찍고 ERRORS 를 올린다. vsd_cmp
# 와 달리 "값이 다르다" 가 아니라 "값을 얻지 못했다" 를 뜻한다 — 문구로 구분해 두지 않으면
# "불일치 0 + 오류 N" 이 통과로 오독될 위험이 있다 (이 오독 자체가 이번 라운드의 결함이었다).
vsd_err() {
  local label="$1" reason="$2"
  printf "  [ERR ] %-40s %s\n" "$label" "$reason"
  ERRORS=$((ERRORS + 1))
}

# vsd_cmp <label> <old> <new> — 값이 같으면 [OK] 를 찍고 0 을 반환한다. 다르면
# [DIFF] 를 찍고 DIFFS 를 올린 뒤 1 을 반환한다.
#
# ★ 이 스크립트는 set -euo pipefail 이다. vsd_cmp 를 맨 명령으로 호출하면 첫
# 불일치에서 스크립트가 죽어 나머지 설정을 대조하지 못한다. 호출부(vsd_main)는
# 반드시 반환값을 소비한다 (`vsd_cmp … || true`, 또는 `if ! vsd_cmp …`).
vsd_cmp() {
  local label="$1" old="$2" new="$3"
  if [ "$old" = "$new" ]; then
    printf "  [OK  ] %-40s %s\n" "$label" "$old"
    return 0
  fi
  printf "  [DIFF] %-40s old=%s new=%s\n" "$label" "$old" "$new"
  DIFFS=$((DIFFS + 1))
  return 1
}

vsd_main() {
  RPC=""; BRIDGE=""; OLD=""; NEW=""; EXTRA_TOKENS=""
  DIFFS=0
  ERRORS=0

  while [ $# -gt 0 ]; do
    case "$1" in
      --rpc-url) RPC="$2"; shift 2;;
      --bridge)  BRIDGE="$2"; shift 2;;
      --old)     OLD="$2"; shift 2;;
      --new)     NEW="$2"; shift 2;;
      --token)
        # 여기서 걸러두지 않으면 오타난 주소가 뒤쪽 grep 에서 말없이 탈락한다 —
        # 운영자는 그 토큰을 검사했다고 믿지만 실제로는 대조되지 않은 채 통과가 찍힌다.
        # 조회가 시작되기 전, 인자 오류로 즉시 알린다.
        if ! [[ "$2" =~ ^0x[0-9a-fA-F]{40}$ ]]; then
          echo "--token 값이 올바른 주소 형식이 아니다: $2" >&2; usage; return 1
        fi
        EXTRA_TOKENS="$EXTRA_TOKENS $2"; shift 2;;
      -h|--help) usage; return 0;;
      *) echo "알 수 없는 옵션: $1" >&2; usage; return 1;;
    esac
  done
  [ -n "$RPC" ] && [ -n "$BRIDGE" ] && [ -n "$OLD" ] && [ -n "$NEW" ] || {
    echo "--rpc-url / --bridge / --old / --new 는 모두 필수다." >&2; usage; return 1; }
  [ "$(echo "$OLD" | tr 'A-Z' 'a-z')" != "$(echo "$NEW" | tr 'A-Z' 'a-z')" ] || {
    echo "--old 와 --new 가 같은 주소다. 확인하라." >&2; return 1; }

  echo "==================================================================="
  echo " BridgeVerifier 설정 전수 대조"
  echo "   RPC     : $RPC   (chainId $(cast chain-id --rpc-url "$RPC" 2>&1 | head -1))"
  echo "   브릿지   : $BRIDGE"
  echo "   구(old) : $OLD"
  echo "   신(new) : $NEW"
  echo "==================================================================="
  echo
  echo "[1] 스칼라 설정 (슬롯 1~8)"
  names=(_ priceFeed _finalizeBridgeGas _defaultExFeeRate _defaultTokenPrice _minimumTokenValue _verificationAmountThreshold _periodTotalValueThreshold _timeWindow)
  for i in 1 2 3 4 5 6 7 8; do
    label="slot$i ${names[$i]}"
    if ! o="$(sto "$OLD" "$i")"; then vsd_err "$label" "old 슬롯 조회 실패"; continue; fi
    if ! n="$(sto "$NEW" "$i")"; then vsd_err "$label" "new 슬롯 조회 실패"; continue; fi
    if ! od="$(dec "$o")"; then vsd_err "$label" "old 값 변환 실패 ($o)"; continue; fi
    if ! nd="$(dec "$n")"; then vsd_err "$label" "new 값 변환 실패 ($n)"; continue; fi
    vsd_cmp "$label" "$od" "$nd" || true
  done

  # allChainIDs 호출 자체의 성공/실패를 먼저 가린다 — 호출은 성공했는데 결과가 정상적으로
  # 빈 배열인 경우(등록 체인 0개)는 오류가 아니다. 뒤에서 grep 이 "매치 없음"으로 실패하는
  # 것과, 여기서 cast call 자체가 실패하는 것을 섞으면 다시 예전 결함으로 돌아간다.
  CHAINS=""
  CHAINS_OK=1
  if raw_chains="$(cast call "$BRIDGE" "allChainIDs()(uint[])" --rpc-url "$RPC" 2>/dev/null)"; then
    # 주의: cast 는 큰 수에 "11155111 [1.115e7]" 처럼 주석을 붙인다. `tr ',' '\n' | sed
    # 's/\[[^]]*\]//g'` 로 대괄호를 지우던 예전 방식은 콤마로 자른 첫 줄에 바깥 여는
    # 대괄호가 남는 바람에(예: "[56 [5.6e1]") \[ 가 바깥 [ 에 매치해 그 줄 전체가
    # 사라졌다(단일 원소 배열도 콤마가 없어 같은 이유로 통째로 사라짐).
    # 대신 주석 패턴("[<숫자>e<지수>]")만 먼저 지운다 — 주석 안에는 ']' 가 없으므로
    # 이 치환은 주석에만 매치하고 바깥 대괄호는 건드리지 않는다. 남은 것은 바깥 대괄호와
    # 콤마로 구분된 정수뿐이므로 숫자 토큰만 뽑으면 원소 수·구분자·주석 유무에 무관하게
    # 동작한다.
    # 아래 grep 의 || true 는 "매치 없음"(정상적으로 빈 목록) 만 삼킨다 — cast call 자체의
    # 실패는 위 if 에서 이미 걸러졌으므로 여기 섞이지 않는다.
    CHAINS=$(printf '%s' "$raw_chains" | sed 's/\[[0-9.]*e[0-9+-]*\]//g' | grep -oE '[0-9]+' || true)
    # raw_chains 가 "[]"(진짜 빈 배열)가 아닌데 체인이 0개로 파싱되면 파싱 이상이다.
    # 성공한 cast call 을 "정상적으로 빈 목록"으로 오인해 §2/§3 을 조용히 스킵하고
    # 통과를 찍는 것이 이번 라운드의 결함이었다 — 같은 실수를 파싱 쪽에서 반복하지 않는다.
    if [ -z "$CHAINS" ] && [ "$(printf '%s' "$raw_chains" | tr -d '[:space:]')" != "[]" ]; then
      CHAINS_OK=0
      vsd_err "allChainIDs" "파싱 이상 — raw=$raw_chains 에서 체인을 0개로 추출함. §2 gasPrice · §3 토큰 대조 스킵"
    fi
  else
    CHAINS_OK=0
    vsd_err "allChainIDs" "브릿지 체인 열거 실패 — §2 gasPrice · §3 토큰 대조 스킵"
  fi

  echo
  if [ "$CHAINS_OK" -eq 1 ]; then
    echo "[2] _gasPrice (슬롯 10) — 등록 체인 $(echo "$CHAINS" | wc -w | tr -d ' ')개"
    for id in $CHAINS; do
      label="chain $id"
      sl="$(mslot "$(printf '%064x' "$id")" 10)"
      if ! o="$(sto "$OLD" "$sl")"; then vsd_err "$label" "old 조회 실패"; continue; fi
      if ! n="$(sto "$NEW" "$sl")"; then vsd_err "$label" "new 조회 실패"; continue; fi
      if ! od="$(dec "$o")"; then vsd_err "$label" "old 값 변환 실패 ($o)"; continue; fi
      if ! nd="$(dec "$n")"; then vsd_err "$label" "new 값 변환 실패 ($n)"; continue; fi
      vsd_cmp "$label" "$od" "$nd" || true
    done
  else
    echo "[2] _gasPrice (슬롯 10) — 체인 열거 실패로 스킵"
  fi

  # 등록 체인의 토큰 쌍에서 주소를 전부 뽑는다 (양쪽 다). allChainIDs 자체가 실패했으면
  # §3 은 통째로 스킵한다 — 의존 데이터가 없는 채로 "토큰 0개" 를 정상 결과로 보고하지 않는다.
  TOKENS=""
  if [ "$CHAINS_OK" -eq 1 ]; then
    # TokenPair 는 (address,address,bool,bool,uint,uint,uint) = 7워드.
    # 동적 배열이라 앞에 offset·length 2워드가 붙는다. 그 뒤부터 7워드씩 끊어
    # 0·1번 워드(= 양쪽 토큰 주소)만 뽑는다. uint 값이 주소처럼 보이는 것에 속지 않는다.
    for id in $CHAINS; do
      if ! raw="$(cast call "$BRIDGE" "allTokenPairs(uint256)" "$id" --rpc-url "$RPC" 2>/dev/null)"; then
        vsd_err "allTokenPairs(chain $id)" "토큰 쌍 열거 실패 — 이 체인의 토큰 대조 스킵"
        continue
      fi
      # 파싱 전에 구조부터 검증한다 — '0x'(malformed)를 "토큰 없음"으로 삼켜 슬롯 16
      # 대조를 통째로 건너뛴 채 통과하던 결함이 이번 라운드의 발단이었다. 정상적인
      # 빈 배열(offset=32, length=0, 2워드)은 오류가 아니다 — 0개로 보고하고 통과한다.
      if ! body="$(vsd_abi_array "$raw" 7 2>/dev/null)"; then
        # ★ set -e 주의: 이 재호출은 if 조건이 아니라 맨 대입문이다 — `|| true` 없이
        # 두면 vsd_abi_array 가 다시 non-zero 를 반환하는 순간 스크립트가 vsd_err
        # 호출 전에 죽는다(맨 명령으로 쓴 실패한 커맨드 서브스티튜션도 set -e 를
        # 발동시킨다 — if 안에서 썼던 첫 호출과 다르다).
        reason="$(vsd_abi_array "$raw" 7 2>&1 1>/dev/null || true)"
        vsd_err "allTokenPairs(chain $id)" "$reason — 이 체인의 토큰 대조 스킵"
        continue
      fi
      TOKENS="$TOKENS $(echo "$body" | fold -w64 | awk '
        {
          i=(NR-1)%7
          if (i==0 || i==1) {
            a=substr($0,25)
            if (substr($0,1,24)=="000000000000000000000000" && a!="0000000000000000000000000000000000000000") print "0x" a
          }
        }')"
    done
    TOKENS=$(echo "$TOKENS $EXTRA_TOKENS" | tr ' ' '\n' | tr 'A-Z' 'a-z' | grep -E '^0x[0-9a-f]{40}$' | sort -u || true)
  fi

  echo
  if [ "$CHAINS_OK" -eq 1 ]; then
    # grep -c 는 매치가 0 건이어도 "0" 을 출력하고 종료코드 1 을 낸다. `|| echo 0` 을 쓰면
    # 0 건일 때 "0" 이 두 번 찍힌다. 종료코드만 삼키면 된다.
    echo "[3] 토큰별 설정 — _exFeeRate (슬롯 9) · _minimumTokenValueOf (슬롯 16) — 토큰 $(echo "$TOKENS" | grep -c . || true)개"
  else
    echo "[3] 토큰별 설정 — 체인 열거 실패로 스킵됨 (allChainIDs 참고)"
  fi
  NOTES=""
  for t in $TOKENS; do
    key="$(printf '%024d%s' 0 "$(echo "$t" | sed 's/^0x//')")"
    s9="$(mslot "$key" 9)"; s11="$(mslot "$key" 11)"; s12="$(mslot "$key" 12)"
    # 슬롯 16 = _minimumTokenValueOf. 15 가 아니다 — _valueLimitWhitelist 가
    # EnumerableSet 이라 14(_values)·15(_positions) 두 칸을 쓴다.
    # 값 0 은 "미설정 → 글로벌 사용" 이므로 구/신이 같이 0 이면 정상이다.
    s16="$(mslot "$key" 16)"
    if ! o9="$(sto "$OLD" "$s9")";   then vsd_err "$t exFee" "old 조회 실패"; continue; fi
    if ! n9="$(sto "$NEW" "$s9")";   then vsd_err "$t exFee" "new 조회 실패"; continue; fi
    if ! o11="$(sto "$OLD" "$s11")"; then vsd_err "$t volume" "old 누적치 조회 실패"; continue; fi
    if ! n11="$(sto "$NEW" "$s11")"; then vsd_err "$t volume" "new 누적치 조회 실패"; continue; fi
    if ! o12="$(sto "$OLD" "$s12")"; then vsd_err "$t history" "old 누적치 조회 실패"; continue; fi
    if ! n12="$(sto "$NEW" "$s12")"; then vsd_err "$t history" "new 누적치 조회 실패"; continue; fi
    if ! o16="$(sto "$OLD" "$s16")"; then vsd_err "$t minValueOf" "old 조회 실패"; continue; fi
    if ! n16="$(sto "$NEW" "$s16")"; then vsd_err "$t minValueOf" "new 조회 실패"; continue; fi
    if ! od9="$(dec "$o9")";   then vsd_err "$t exFee" "old 값 변환 실패 ($o9)"; continue; fi
    if ! nd9="$(dec "$n9")";   then vsd_err "$t exFee" "new 값 변환 실패 ($n9)"; continue; fi
    if ! od16="$(dec "$o16")"; then vsd_err "$t minValueOf" "old 값 변환 실패 ($o16)"; continue; fi
    if ! nd16="$(dec "$n16")"; then vsd_err "$t minValueOf" "new 값 변환 실패 ($n16)"; continue; fi
    vsd_cmp "$t exFee" "$od9" "$nd9" || true
    vsd_cmp "$t minValueOf" "$od16" "$nd16" || true
    if [ "$o11" != "$n11" ] || [ "$o12" != "$n12" ]; then
      od11="$(dec "$o11" 2>/dev/null || echo '?')"; nd11="$(dec "$n11" 2>/dev/null || echo '?')"
      NOTES="$NOTES\n    $t  누적 volume old=$od11 new=$nd11"
    fi
  done

  echo
  echo "[4] 역할 멤버 (슬롯 13) — 12종"
  ROLES="DEFAULT_ADMIN:0x0000000000000000000000000000000000000000000000000000000000000000
ADMIN:0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775
EDITOR:0x21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c
OPERATOR:0x97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929
VALIDATOR:0x21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926
PRICER:0xc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a
VERIFIER:0x0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09
BRIDGE:0x52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f
MINTER:0x9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6
EXECUTOR:0xd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63
INITIATOR:0x6b8b15f1c11543d8280deaa7c24d12fffba6a357e4428e8c43e4234790186bff
LINKER:0x733bac3dca102687aa08c854c5f9067fc424f98fd8e90e41ad6b73aecc59a4fd"
  for entry in $ROLES; do
    name="${entry%%:*}"; hash="${entry##*:}"
    if ! o="$(rolemembers "$OLD" "$hash")"; then vsd_err "$name" "old 역할 조회 실패"; continue; fi
    if ! n="$(rolemembers "$NEW" "$hash")"; then vsd_err "$name" "new 역할 조회 실패"; continue; fi
    if [ "$o" = "$n" ]; then
      [ "$o" != "[]" ] && [ -n "$o" ] && printf "  [OK  ] %-14s %s\n" "$name" "$o"
    else
      printf "  [DIFF] %-14s\n         old=%s\n         new=%s\n" "$name" "$o" "$n"
      DIFFS=$((DIFFS + 1))
    fi
  done

  if [ -n "$NOTES" ]; then
    echo
    echo "[주의] 모니터링 누적치는 재배포 시 넘어오지 않는다 (옮길 setter 가 없다)."
    printf "%b\n" "$NOTES"
    # 이 판단은 비교 대상이 아니라 안내 문구일 뿐이지만, 조회 실패를 "0" 으로 오인해
    # "영향 없음" 이라고 잘못 안내하면 안 되므로 여기도 실패를 구분해 처리한다.
    t6="?"; t7="?"
    if t6o="$(sto "$OLD" 6)" && t6="$(dec "$t6o")" \
       && t7o="$(sto "$OLD" 7)" && t7="$(dec "$t7o")"; then
      if [ "$t6" = "0" ] && [ "$t7" = "0" ]; then
        echo "    → 구 베리파이어의 두 임계값이 모두 0 이므로 모니터링이 꺼져 있다. 영향 없음."
      else
        echo "    ★★ 임계값이 0 이 아니다 (amount=$t6, period=$t7)."
        echo "       재배포 순간 기간합산 윈도가 비워진다. 진행 전에 담당자와 협의할 것."
      fi
    else
      vsd_err "monitoring threshold(old)" "임계값(슬롯 6/7) 조회 실패 — 영향 여부 자동 판단 불가"
      echo "    ★★ 임계값 조회에 실패해 영향 여부를 자동으로 판단할 수 없다. 담당자와 확인할 것."
    fi
  fi

  echo
  echo "==================================================================="
  echo " 설정 불일치 총합: $DIFFS   조회 실패 총합: $ERRORS"
  if [ "$ERRORS" -gt 0 ]; then
    echo " → ★★ 검증되지 않음 — 조회 실패 $ERRORS 건. 이 결과로는 통과 여부를 판단할 수 없다."
    echo "    (불일치 0 이어도 통과가 아니다. 조회 실패 원인을 해결한 뒤 다시 돌려라.)"
  elif [ "$DIFFS" -eq 0 ]; then
    echo " → 통과. 브릿지를 새 베리파이어로 연결해도 된다."
  else
    echo " → ★ 중단. 빠진 설정을 채운 뒤 다시 돌려라."
  fi
  echo "==================================================================="
  [ "$DIFFS" -eq 0 ] && [ "$ERRORS" -eq 0 ]
}

if [ "${BASH_SOURCE[0]}" = "$0" ]; then
  vsd_main "$@"
fi
