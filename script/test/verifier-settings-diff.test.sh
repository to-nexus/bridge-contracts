#!/usr/bin/env bash
# Pure-bash tests for script/verifier-settings-diff.sh -- no RPC, no forked chain.
#
# T1/T2 exercise the comparison primitives directly, by sourcing the target script
# (the source guard at its EOF means sourcing only defines functions, nothing runs).
# T3/T4 run the target script end-to-end behind a `cast` stub that fakes only
# `storage` / `call` / `chain-id` and fails loudly on anything it doesn't recognize;
# `keccak` and `to-dec` are delegated to the real `cast` binary so the slot-key math
# itself is never faked.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
TARGET="$REPO_ROOT/script/verifier-settings-diff.sh"
REAL_CAST="$(command -v cast || true)"

[ -f "$TARGET" ] || { echo "target script not found: $TARGET" >&2; exit 1; }
[ -n "$REAL_CAST" ] || { echo "cast binary not found on PATH -- required for T1 and for keccak/to-dec delegation" >&2; exit 1; }

FAILURES=0
pass() { echo "PASS: $1"; }
fail() {
  echo "FAIL: $1" >&2
  FAILURES=$((FAILURES + 1))
}

# Fixed test addresses, built by zero-padding a small decimal so the length is
# always a correct 40 hex chars regardless of how many digits the number has.
OLD_ADDR="0x$(printf '%040x' 161)" # ...a1
NEW_ADDR="0x$(printf '%040x' 178)" # ...b2
BRIDGE_ADDR="0x$(printf '%040x' 195)" # ...c3
TOKEN_ADDR="0x$(printf '%040x' 222)" # ...de
RPC_URL="http://stub-rpc.invalid:1"

# ---------------------------------------------------------------------------
# T1 -- mslot computes the slot-16 mapping key exactly, using the real `cast
# keccak` (no stub anywhere in this test). The expected hash is hardcoded, but
# it is re-derived here a second, independent way (`cast abi-encode`, which
# does not share any code with mslot's manual hex construction) so the
# constant itself cannot silently rot. It is the same formula --
# keccak256(abi.encode(token, uint(16))) -- that
# test_minimumTokenValueOf_lives_in_slot_16 in
# test/BridgeMinimumTokenValueOverride.t.sol checks against actual EVM storage
# for a different (dynamically deployed) token address: that test proves the
# mapping physically lives at slot 16, this test proves the bash tooling
# computes that slot's key correctly -- together they cover the claim end to
# end even though the two tests don't share a token address.
# ---------------------------------------------------------------------------
test_t1() {
  local name="T1 mslot / real cast keccak matches the slot-16 formula"
  local addr="0x1111111111111111111111111111111111111111"
  local expected="0xe7b8b5c9b1eccf490ba7eccfc874c5312753d50e634fc2abb2aadf2c68b8f2c9"

  local reencoded rehash
  reencoded="$("$REAL_CAST" abi-encode "f(address,uint256)" "$addr" 16)"
  rehash="$("$REAL_CAST" keccak "$reencoded")"
  if [ "$rehash" != "$expected" ]; then
    fail "$name (hardcoded expected constant no longer matches cast abi-encode -- update it)"
    return
  fi

  if (
    source "$TARGET"
    key="$(printf '%024d%s' 0 "$(echo "$addr" | sed 's/^0x//')")"
    actual="$(mslot "$key" 16)"
    [ "$actual" = "$expected" ]
  ); then
    pass "$name"
  else
    fail "$name"
  fi
}

# ---------------------------------------------------------------------------
# T2 -- vsd_cmp itself: 0 + no DIFFS change on a match, non-zero + DIFFS++ on a
# mismatch.
# ---------------------------------------------------------------------------
test_t2() {
  local name="T2 vsd_cmp / equal->0 no-op, mismatch->nonzero and DIFFS++"

  if (
    source "$TARGET"
    DIFFS=0

    rc=0
    vsd_cmp "same" "5" "5" >/dev/null || rc=$?
    [ "$rc" -eq 0 ] || { echo "expected 0 for a match, got $rc" >&2; exit 1; }
    [ "$DIFFS" -eq 0 ] || { echo "DIFFS moved on a match: $DIFFS" >&2; exit 1; }

    rc=0
    vsd_cmp "different" "5" "6" >/dev/null || rc=$?
    [ "$rc" -ne 0 ] || { echo "expected non-zero for a mismatch" >&2; exit 1; }
    [ "$DIFFS" -eq 1 ] || { echo "DIFFS must be exactly 1 after one mismatch, got $DIFFS" >&2; exit 1; }
  ); then
    pass "$name"
  else
    fail "$name"
  fi
}

# Writes the `cast` stub used by T3-T8. Static content -- everything that
# varies between calls (addresses, slot keys, the injected mismatch value, and
# -- for T5-T8 -- which lookup should fail) comes from VSD_* environment
# variables set by the caller, so the file only needs to be written once per
# test. VSD_FAIL_STORAGE / VSD_FAIL_CHAINS / VSD_FAIL_ROLES /
# VSD_FAIL_TOKENPAIRS default to unset (=0), which reproduces the exact T3/T4
# behavior -- a test that doesn't set them exercises no injected failure.
write_stub() {
  local path="$1"
  cat > "$path" <<'STUB'
#!/usr/bin/env bash
set -euo pipefail

fail_stub() { echo "STUB ERROR: $*" >&2; exit 99; }

case "$1" in
  keccak) shift; exec "$VSD_REAL_CAST" keccak "$@" ;;
  to-dec) shift; exec "$VSD_REAL_CAST" to-dec "$@" ;;
  chain-id)
    echo "1234"
    ;;
  storage)
    addr="$(echo "$2" | tr 'A-Z' 'a-z')"; slot="$3"
    if [ "${VSD_FAIL_STORAGE:-0}" = "1" ]; then
      echo "(stub) simulated cast storage failure" >&2
      exit 1
    fi
    case "$slot" in
      1) echo "0x0000000000000000000000000000000000000000000000000000000000000011" ;;
      2) echo "0x0000000000000000000000000000000000000000000000000000000000000022" ;;
      3) echo "0x0000000000000000000000000000000000000000000000000000000000000033" ;;
      4) echo "0x0000000000000000000000000000000000000000000000000000000000000044" ;;
      5) echo "0x0000000000000000000000000000000000000000000000000000000000000055" ;;
      6) echo "0x0000000000000000000000000000000000000000000000000000000000000000" ;;
      7) echo "0x0000000000000000000000000000000000000000000000000000000000000000" ;;
      8) echo "0x0000000000000000000000000000000000000000000000000000000000000088" ;;
      "$VSD_SLOT9") echo "0x0000000000000000000000000000000000000000000000000000000000000007" ;;
      "${VSD_SLOT10:-}") echo "0x0000000000000000000000000000000000000000000000000000000000000012" ;;
      "$VSD_SLOT11") echo "0x0000000000000000000000000000000000000000000000000000000000000000" ;;
      "$VSD_SLOT12") echo "0x0000000000000000000000000000000000000000000000000000000000000000" ;;
      "$VSD_SLOT16")
        case "$addr" in
          "$VSD_OLD") echo "0x0000000000000000000000000000000000000000000000000000000000000005" ;;
          "$VSD_NEW") echo "$VSD_MIN16_NEW" ;;
          *) fail_stub "unexpected address for slot16 storage: $addr" ;;
        esac
        ;;
      *)
        # Extra chain _gasPrice slots (base slot 10, keyed by chain id), for tests
        # that exercise more than one chain id and so can't fit in the single
        # VSD_SLOT10 case above. VSD_GASPRICE_MAP is a space-separated list of
        # "slot=value" entries; unset (all T3-T8) means this loop matches nothing
        # and falls through to fail_stub exactly as before.
        matched=0
        for kv in ${VSD_GASPRICE_MAP:-}; do
          k="${kv%%=*}"; v="${kv#*=}"
          if [ "$slot" = "$k" ]; then echo "$v"; matched=1; break; fi
        done
        [ "$matched" = "1" ] || fail_stub "unexpected storage slot: $slot (addr=$addr)"
        ;;
    esac
    ;;
  call)
    addr="$(echo "$2" | tr 'A-Z' 'a-z')"; sig="$3"
    case "$sig" in
      "allChainIDs()(uint[])")
        [ "$addr" = "$VSD_BRIDGE" ] || fail_stub "unexpected bridge address: $addr"
        if [ "${VSD_FAIL_CHAINS:-0}" = "1" ]; then
          echo "(stub) simulated allChainIDs failure" >&2
          exit 1
        fi
        echo "${VSD_CHAINS_LIST:-[]}"
        ;;
      "getRoleMembers(bytes32)(address[])")
        case "$addr" in
          "$VSD_OLD" | "$VSD_NEW") : ;;
          *) fail_stub "unexpected address for getRoleMembers: $addr" ;;
        esac
        if [ "${VSD_FAIL_ROLES:-0}" = "1" ]; then
          echo "(stub) simulated getRoleMembers failure" >&2
          exit 1
        fi
        echo "[]"
        ;;
      "allTokenPairs(uint256)")
        [ "$addr" = "$VSD_BRIDGE" ] || fail_stub "unexpected bridge address: $addr"
        if [ "${VSD_FAIL_TOKENPAIRS:-0}" = "1" ]; then
          echo "(stub) simulated allTokenPairs failure" >&2
          exit 1
        fi
        echo "${VSD_TOKENPAIRS_RESULT:-0x}"
        ;;
      *) fail_stub "unexpected call signature: $sig" ;;
    esac
    ;;
  *) fail_stub "unexpected cast subcommand: $1" ;;
esac
STUB
  chmod +x "$path"
}

# Slot keys for the one test token, computed once via the real `cast keccak`
# (same math as mslot -- this only fixes the *keys* the stub must recognize,
# not the compared values, so it doesn't fake the thing being tested).
TOKEN_KEY="$(printf '%024d%s' 0 "$(echo "$TOKEN_ADDR" | sed 's/^0x//')")"
SLOT9="$("$REAL_CAST" keccak "0x$TOKEN_KEY$(printf '%064x' 9)")"
SLOT11="$("$REAL_CAST" keccak "0x$TOKEN_KEY$(printf '%064x' 11)")"
SLOT12="$("$REAL_CAST" keccak "0x$TOKEN_KEY$(printf '%064x' 12)")"
SLOT16="$("$REAL_CAST" keccak "0x$TOKEN_KEY$(printf '%064x' 16)")"

# T8 needs allChainIDs to return one real chain id so the allTokenPairs call
# actually gets made (and then fails). That chain id's _gasPrice slot (base
# 10, keyed by the chain id itself rather than a token address) must also be
# recognized by the storage stub, or section [2] would hit the stub's
# catch-all "unexpected storage slot" instead of the failure T8 means to test.
CHAIN_ID_FOR_TESTS=1234
SLOT10="$("$REAL_CAST" keccak "0x$(printf '%064x' "$CHAIN_ID_FOR_TESTS")$(printf '%064x' 10)")"

# Second chain id, used only by T10 (two chains, one with a cast exponent
# annotation on it) to prove the round-3 chain-list parsing fix extracts
# every chain id, not just the ones a single hardcoded VSD_SLOT10 could cover.
CHAIN_ID_FOR_TESTS_2=5678
SLOT10_2="$("$REAL_CAST" keccak "0x$(printf '%064x' "$CHAIN_ID_FOR_TESTS_2")$(printf '%064x' 10)")"

# Synthetic allTokenPairs(uint256) ABI encoding for exactly one TokenPair
# (tokenA=TOKEN_ADDR, tokenB=zero address, the remaining 5 words zeroed),
# used by T9 to prove a single-element chain list drives a real, chain-sourced
# token comparison in section [3] -- not just a token added via --token.
# Layout: word0 offset(0x20) + word1 length(1) + 7 record words
#   (0=tokenA, 1=tokenB, 2..6=bool,bool,uint,uint,uint per TokenPair) = 9 words
#   total (576 hex chars). Round-4 fix: this fixture used to be one word
#   short (8 words / 512 chars, missing the record's FINAL word -- tokenB was
#   present as a zero word all along) and the old parser
#   silently accepted it -- vsd_abi_array now rejects that shape (see T13,
#   which reuses the old broken layout on purpose as a regression fixture).
TOKENPAIRS_ONE_PAIR="0x$(printf '%064x' 32)$(printf '%064x' 1)000000000000000000000000$(echo "$TOKEN_ADDR" | sed 's/^0x//')$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)"

# T13 fixture -- the exact pre-round-4 broken shape: 8 words (512 chars)
# instead of the correct 9. The 7-word TokenPair record carries only 6 words
# here, so what is missing is the record's LAST word, not tokenB (tokenB is
# present as a zero word). vsd_abi_array must now reject this as a
# word-count/record-boundary violation.
TOKENPAIRS_TRUNCATED="0x$(printf '%064x' 32)$(printf '%064x' 1)000000000000000000000000$(echo "$TOKEN_ADDR" | sed 's/^0x//')$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)"

# T14 fixture -- structurally a single complete 7-word TokenPair record (9
# words total, same body as TOKENPAIRS_ONE_PAIR) but the declared length word
# says 2 records instead of 1. vsd_abi_array must reject the length/record
# mismatch even though the word-boundary math alone is clean.
TOKENPAIRS_LENGTH_MISMATCH="0x$(printf '%064x' 32)$(printf '%064x' 2)000000000000000000000000$(echo "$TOKEN_ADDR" | sed 's/^0x//')$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)$(printf '%064x' 0)"

# T15 fixture -- a genuine empty array: offset(0x20) + length(0), exactly 2
# header words and nothing else. This must NOT be treated as an error.
TOKENPAIRS_EMPTY="0x$(printf '%064x' 32)$(printf '%064x' 0)"

# ---------------------------------------------------------------------------
# T3 -- end to end: every slot (including slot 16) matches -> exit 0.
# ---------------------------------------------------------------------------
test_t3() {
  local name="T3 end-to-end / everything matches -> exit 0"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_MIN16_NEW="0x0000000000000000000000000000000000000000000000000000000000000005" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" --token "$TOKEN_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -ne 0 ]; then
    fail "$name (expected exit 0, got $rc)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "설정 불일치 총합: 0"; then
    fail "$name (missing the zero-mismatch summary line)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T4 -- end to end: only slot 16 (minValueOf) differs -> exit non-zero, the
# mismatch is reported, sections after it in the script still ran (the
# set -e early-abort regression this whole refactor exists to prevent), and
# the total mismatch count is exactly the 1 injected.
# ---------------------------------------------------------------------------
test_t4() {
  local name="T4 end-to-end / slot-16 mismatch -> nonzero, later sections still run, DIFFS==1"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_MIN16_NEW="0x0000000000000000000000000000000000000000000000000000000000000009" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" --token "$TOKEN_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[DIFF\].*minValueOf"; then
    fail "$name (no [DIFF] line mentioning minValueOf -- the injected mismatch was not reported)"
    printf '%s\n' "$out" >&2
    return
  fi
  # Regression guard for the set -e early-abort bug: the role section runs
  # after the token loop that contains the injected slot-16 mismatch, so its
  # header must still appear in the output.
  if ! printf '%s\n' "$out" | grep -q "역할 멤버"; then
    fail "$name (role section did not run -- vsd_main likely aborted early on the mismatch)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "설정 불일치 총합: 1"; then
    fail "$name (mismatch total is not exactly 1 -- expected only the injected slot-16 diff)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T5 -- cast storage fails on both old and new (the exact fail-open bug: both
# sides failing used to compare equal and print [OK]). Assert non-zero exit,
# an [ERR ] line, and -- the sharpest regression guard -- that the pass banner
# ("-> 통과") never appears no matter how many lookups failed identically.
# ---------------------------------------------------------------------------
test_t5() {
  local name="T5 end-to-end / cast storage fails on both old and new -> nonzero, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_FAIL_STORAGE=1 \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite storage lookup failures)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \]"; then
    fail "$name (no [ERR ] line reported for the failed storage lookups)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T6 -- allChainIDs call fails. This used to silently become an empty chain
# list (section [2]/[3] quietly skipped, exit 0). Assert non-zero exit, an
# explicit skip notice, no pass banner, and that the role section (which does
# not depend on allChainIDs) still ran -- the abort-early regression guard.
# ---------------------------------------------------------------------------
test_t6() {
  local name="T6 end-to-end / allChainIDs call fails -> nonzero, dependent sections skipped, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_FAIL_CHAINS=1 \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite allChainIDs failure)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "체인 열거 실패"; then
    fail "$name (no skip notice for the failed chain enumeration)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "역할 멤버"; then
    fail "$name (role section did not run -- vsd_main likely aborted early on the allChainIDs failure)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T7 -- getRoleMembers fails for every role. Assert non-zero exit, an [ERR ]
# line, and no pass banner.
# ---------------------------------------------------------------------------
test_t7() {
  local name="T7 end-to-end / getRoleMembers call fails -> nonzero, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_FAIL_ROLES=1 \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite role lookup failures)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \]"; then
    fail "$name (no [ERR ] line reported for the failed role lookups)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T8 -- allTokenPairs fails for a chain that allChainIDs legitimately returned
# (chain enumeration itself succeeds -- only the per-chain token-pair fetch
# fails). Assert non-zero exit, an [ERR ] line, and no pass banner.
# ---------------------------------------------------------------------------
test_t8() {
  local name="T8 end-to-end / allTokenPairs call fails -> nonzero, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_FAIL_TOKENPAIRS=1 \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  # Single-element chain list on purpose: this used to be the exact shape that
  # tripped the chain-list parser's single-element bug (round-3 fix). Using
  # one element here doubles as a regression guard now that the parser
  # handles it correctly.
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite allTokenPairs failure)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \]"; then
    fail "$name (no [ERR ] line reported for the failed allTokenPairs call)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}


# ---------------------------------------------------------------------------
# T9 -- round-3 regression guard: a single-element chain list ("[1234]") used
# to vanish entirely under the old chain-list parser (no comma to split on,
# so the whole "[1234]" matched the bracket-stripping sed in one shot). Assert
# that chain's gasPrice (section [2]) AND a token sourced from that chain's
# own allTokenPairs (section [3], not via --token) both actually get compared,
# and the run exits 0.
# ---------------------------------------------------------------------------
test_t9() {
  local name="T9 end-to-end / single-element chain list -> that chain's gasPrice+token comparisons actually run, exit 0"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_MIN16_NEW="0x0000000000000000000000000000000000000000000000000000000000000005" \
    VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_TOKENPAIRS_RESULT="$TOKENPAIRS_ONE_PAIR" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -ne 0 ]; then
    fail "$name (expected exit 0, got $rc)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "chain $CHAIN_ID_FOR_TESTS"; then
    fail "$name (chain $CHAIN_ID_FOR_TESTS did not appear in section [2] -- single-element parsing regressed)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "$TOKEN_ADDR exFee"; then
    fail "$name (token sourced from allTokenPairs(chain $CHAIN_ID_FOR_TESTS) did not appear in section [3])"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T10 -- round-3 regression guard: when the FIRST element of allChainIDs'
# result carries a cast exponent annotation (e.g. "1234 [1.234e3]"), the old
# parser's bracket-stripping sed matched from that first element's outer '['
# through the annotation's ']' and deleted the whole first line -- dropping
# that chain while keeping the rest. Assert both chains are compared.
# ---------------------------------------------------------------------------
test_t10() {
  local name="T10 end-to-end / first chain has a cast exponent annotation -> both chains compared"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_MIN16_NEW="0x0000000000000000000000000000000000000000000000000000000000000005" \
    VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS [1.234e3], $CHAIN_ID_FOR_TESTS_2]" \
    VSD_GASPRICE_MAP="$SLOT10=0x0000000000000000000000000000000000000000000000000000000000000012 $SLOT10_2=0x0000000000000000000000000000000000000000000000000000000000000012" \
    VSD_TOKENPAIRS_RESULT="$TOKENPAIRS_EMPTY" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -ne 0 ]; then
    fail "$name (expected exit 0, got $rc)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "chain $CHAIN_ID_FOR_TESTS "; then
    fail "$name (annotated first chain $CHAIN_ID_FOR_TESTS did not appear -- it was dropped)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "chain $CHAIN_ID_FOR_TESTS_2"; then
    fail "$name (second chain $CHAIN_ID_FOR_TESTS_2 did not appear)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T11 -- allChainIDs succeeds (cast call itself does not fail) and returns a
# non-empty string that is NOT a genuine empty array "[]", but the chain-list
# parser extracts zero numeric tokens from it (garbage/unexpected content).
# That must be treated as a parse anomaly, not a legitimately-empty chain
# list: vsd_err fires, exit is non-zero, and the pass banner never appears.
# ---------------------------------------------------------------------------
test_t11() {
  local name="T11 end-to-end / non-empty allChainIDs result parses to zero chains -> vsd_err, nonzero, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_CHAINS_LIST="[abc]" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite the parse anomaly)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \].*allChainIDs"; then
    fail "$name (no [ERR ] line reported for the allChainIDs parse anomaly)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T12 -- round-4 regression guard: allTokenPairs returns the literal '0x'
# (malformed -- not a valid empty array, which needs offset+length = 2 words).
# The old parser silently treated this as "no tokens" and passed. Assert
# vsd_err fires, exit is non-zero, and the pass banner never appears.
# ---------------------------------------------------------------------------
test_t12() {
  local name="T12 end-to-end / allTokenPairs returns '0x' -> vsd_err, nonzero, no pass banner"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_SLOT10="$SLOT10" VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_TOKENPAIRS_RESULT="0x" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite the malformed '0x' response)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \].*allTokenPairs"; then
    fail "$name (no [ERR ] line reported for the malformed allTokenPairs response)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T13 -- truncated response: one word short of a full 7-word TokenPair record
# (the exact pre-round-4 broken TOKENPAIRS fixture shape). The old parser
# silently accepted this. Assert vsd_err fires and exit is non-zero.
# ---------------------------------------------------------------------------
test_t13() {
  local name="T13 end-to-end / truncated allTokenPairs response (record one word short) -> vsd_err, nonzero"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_SLOT10="$SLOT10" VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_TOKENPAIRS_RESULT="$TOKENPAIRS_TRUNCATED" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite the truncated response)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \].*allTokenPairs"; then
    fail "$name (no [ERR ] line reported for the truncated allTokenPairs response)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T14 -- length mismatch: declared length word says 2 records but only 1
# record's worth of data is present. Assert vsd_err fires and exit is
# non-zero.
# ---------------------------------------------------------------------------
test_t14() {
  local name="T14 end-to-end / allTokenPairs declares length 2 but carries 1 record -> vsd_err, nonzero"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_SLOT10="$SLOT10" VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_TOKENPAIRS_RESULT="$TOKENPAIRS_LENGTH_MISMATCH" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (pass banner text appeared despite the length mismatch)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "\[ERR \].*allTokenPairs"; then
    fail "$name (no [ERR ] line reported for the allTokenPairs length mismatch)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T15 -- a genuine empty array (offset=32, length=0, exactly 2 header words)
# must NOT be treated as an error: 0 tokens reported, run still passes.
# ---------------------------------------------------------------------------
test_t15() {
  local name="T15 end-to-end / genuine empty allTokenPairs array -> not an error, 0 tokens, passes"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT10="$SLOT10" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_CHAINS_LIST="[$CHAIN_ID_FOR_TESTS]" VSD_TOKENPAIRS_RESULT="$TOKENPAIRS_EMPTY" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -ne 0 ]; then
    fail "$name (expected exit 0, got $rc)"
    printf '%s\n' "$out" >&2
    return
  fi
  if printf '%s\n' "$out" | grep -q "\[ERR \].*allTokenPairs"; then
    fail "$name (a genuinely empty array was reported as an allTokenPairs error)"
    printf '%s\n' "$out" >&2
    return
  fi
  # `grep -c` prints "0" on zero matches AND still exits 1, so the header's old
  # `|| echo 0` printed the count twice ("0\n0개"). That path only became
  # reachable once a genuinely empty array stopped being an error, so the
  # header now swallows just the exit code (`|| true`). Pin it: the count must
  # appear exactly once.
  if [ "$(printf '%s\n' "$out" | grep -c '토큰 0개')" != "1" ]; then
    fail "$name (expected the token count to be printed exactly once)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "브릿지를 새 베리파이어로 연결해도 된다"; then
    fail "$name (expected the pass banner to appear)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T16 -- a mistyped --token value must be rejected at argument-parse time,
# before any lookup starts: exit non-zero as an argument error. No stub is
# needed since the script must fail before making any cast call.
# ---------------------------------------------------------------------------
test_t16() {
  local name="T16 argument parsing / invalid --token -> argument error, nonzero exit"
  local out rc
  local bad_token="0xZZ11111111111111111111111111111111111a"

  set +e
  out="$(
    "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" --token "$bad_token" 2>&1
  )"
  rc=$?
  set -e

  if [ "$rc" -eq 0 ]; then
    fail "$name (expected a non-zero exit, got 0)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q -- "--token 값이 올바른 주소 형식이 아니다"; then
    fail "$name (no argument-error message reported for the invalid --token value)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

# ---------------------------------------------------------------------------
# T17 -- a valid --token must actually appear in the compared set (section
# [3]), sourced purely via --token (no chains registered), proving the flag
# actually drives a comparison rather than merely being accepted.
# ---------------------------------------------------------------------------
test_t17() {
  local name="T17 end-to-end / valid --token actually appears in the compared set"
  local stubdir out rc
  stubdir="$(mktemp -d)"
  write_stub "$stubdir/cast"

  set +e
  out="$(
    VSD_REAL_CAST="$REAL_CAST" VSD_OLD="$OLD_ADDR" VSD_NEW="$NEW_ADDR" VSD_BRIDGE="$BRIDGE_ADDR" \
    VSD_SLOT9="$SLOT9" VSD_SLOT11="$SLOT11" VSD_SLOT12="$SLOT12" VSD_SLOT16="$SLOT16" \
    VSD_MIN16_NEW="0x0000000000000000000000000000000000000000000000000000000000000005" \
    PATH="$stubdir:$PATH" "$TARGET" --rpc-url "$RPC_URL" --bridge "$BRIDGE_ADDR" \
      --old "$OLD_ADDR" --new "$NEW_ADDR" --token "$TOKEN_ADDR" 2>&1
  )"
  rc=$?
  set -e
  rm -rf "$stubdir"

  if [ "$rc" -ne 0 ]; then
    fail "$name (expected exit 0, got $rc)"
    printf '%s\n' "$out" >&2
    return
  fi
  if ! printf '%s\n' "$out" | grep -q "$TOKEN_ADDR exFee"; then
    fail "$name (the --token address did not appear in the compared set)"
    printf '%s\n' "$out" >&2
    return
  fi
  pass "$name"
}

test_t1
test_t2
test_t3
test_t4
test_t5
test_t6
test_t7
test_t8
test_t9
test_t10
test_t11
test_t12
test_t13
test_t14
test_t15
test_t16
test_t17

echo
if [ "$FAILURES" -eq 0 ]; then
  echo "ALL PASS"
  exit 0
fi
echo "FAILURES: $FAILURES"
exit 1
