// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console} from "forge-std/Script.sol";

/**
 * @title ForwardLibVerify
 * @notice READ-ONLY deployment-integrity gate for `CrossBridge`'s `ForwardLib` link.
 * Sends NO transactions — there is no
 * `vm.broadcast`/`vm.startBroadcast` anywhere in this file. It only reads on-chain code
 * and the local build artifact, and reverts if anything is wrong.
 * @dev Usage (run against an ALREADY-DEPLOYED, but not-yet-promoted, `CrossBridge`
 * implementation, BEFORE pointing the proxy at it):
 *
 *   forge script script/ForwardLibVerify.s.sol \
 *     --sig 'run(address,address)' <newImplementation> <forwardLibAddress> \
 *     --rpc-url <rpc>
 *
 * See `FORWARD_DEPLOYMENT.md` for the full 4-step deployment procedure this script
 * is step 1-3 of (step 4 is an on-chain smoke test, documented but not automated here —
 * automating it would require sending a real forwarded finalize transaction, which is
 * out of scope for a read-only gate).
 *
 * Checks performed (all read-only):
 *   1. `lib.code.length > 0` — the configured `ForwardLib` address actually has code.
 *   2. `lib.codehash` matches the `ForwardLib` build artifact's runtime codehash.
 *   3. Every `ForwardLib` link-reference offset recorded in the `CrossBridge` build
 *      artifact, when read directly out of `impl`'s ACTUAL deployed runtime bytecode,
 *      contains exactly `lib`'s address (rejects a partially-relinked or mismatched
 *      implementation).
 *   4. The build artifact's unlinked runtime template, with every one of those offsets
 *      patched to `lib`, hashes to EXACTLY `impl.codehash` — proving the entire deployed
 *      runtime (not just the 20 bytes at each link offset) matches what the compiler
 *      produced for `CrossBridge` with `ForwardLib` correctly linked. This is strictly
 *      stronger than (3) alone: (3) only proves "an address sits at that offset", not
 *      that the rest of the bytecode wasn't tampered with or built from a different
 *      source.
 *
 * A deliberate design choice: this does NOT scan the runtime bytecode for the `lib`
 * address as a raw byte pattern anywhere. A pattern scan can (a) false-positive on an
 * unrelated constant that happens to equal the address, and (b) can never prove that
 * every relocation was actually patched (only that the address appears somewhere).
 * Reading the compiler-declared link-reference offsets is the only sound way to answer
 * "is EVERY relocation correctly patched, and is the rest of the code exactly what was
 * compiled".
 */
contract ForwardLibVerify is Script {
    /// @notice The library artifact's fully-qualified name, resolved via `vm.getDeployedCode`.
    string internal constant LIB_ARTIFACT = "ForwardLib.sol:ForwardLib";

    /// @notice Path to the host contract's build artifact, read via `vm.readFile`.
    string internal constant HOST_ARTIFACT_PATH = "out/CrossBridge.sol/CrossBridge.json";

    /// @notice Path to the library's own build artifact, read via `vm.readFile`.
    string internal constant LIB_ARTIFACT_JSON_PATH = "out/ForwardLib.sol/ForwardLib.json";

    /// @notice The only external library `CrossBridge` is expected to link against.
    string internal constant EXPECTED_LIB_NAME = "ForwardLib";

    /// @notice Fully-qualified library name as the compiler records it under
    /// `metadata.settings.libraries` when the artifact is built pre-linked.
    string internal constant LIB_FQN = "src/lib/ForwardLib.sol:ForwardLib";

    error ForwardLibVerifyLibHasNoCode(address lib);
    error ForwardLibVerifyLibCodehashMismatch(bytes32 expected, bytes32 actual);
    error ForwardLibVerifyNoLinkReferenceFound();
    /// @notice Artifact was built pre-linked but declares no `settings.libraries` entry for `ForwardLib`.
    error ForwardLibVerifyPrelinkedLibraryNotDeclared();
    /// @notice Artifact was built pre-linked to a DIFFERENT library address than the one being verified.
    error ForwardLibVerifyPrelinkedAddressMismatch(address expected, address declared);
    error ForwardLibVerifyUnexpectedLinkReferenceContract(string file, string name);
    error ForwardLibVerifyLinkReferenceOutOfRange(uint offset, uint length, uint runtimeLength);
    error ForwardLibVerifyLinkedAddressMismatch(uint offset, address expected, address actual);
    error ForwardLibVerifyRuntimeLengthMismatch(uint artifactLength, uint implLength);
    error ForwardLibVerifyPatchedRuntimeMismatch(bytes32 expectedImplCodehash, bytes32 patchedArtifactCodehash);

    /**
     * @notice Entry point for `forge script --sig 'run(address,address)'`.
     * @param impl The already-deployed (not-yet-promoted) `CrossBridge` implementation.
     * @param lib The `ForwardLib` address `impl` is expected to be linked against.
     */
    function run(address impl, address lib) external view {
        verify(impl, lib);
        console.log("ForwardLib link integrity OK");
        console.log("  impl:", impl);
        console.log("  lib: ", lib);
    }

    /**
     * @notice Runs all four checks; reverts on the first failure. Exposed as a plain
     * function (not folded into `run`) so `test/ForwardLibVerify.t.sol` can unit-test
     * this logic directly against synthetic addresses/bytecode, independent of any real
     * `forge script` invocation.
     */
    function verify(address impl, address lib) public view {
        _verifyLibCode(lib);
        _verifyLinkedImplementation(impl, lib);
    }

    /**
     * @dev Checks (1) and (2): the library itself has code and matches the build
     * artifact.
     *
     * Solidity's EXTERNAL library functions carry a compiler-inserted guard that
     * rejects a direct `CALL` (as opposed to the intended `DELEGATECALL`) by comparing
     * `ADDRESS` against the library's own deployment address. That comparison constant
     * is filled in by the library's OWN constructor at deploy time — it is NOT present
     * in `vm.getDeployedCode`'s static artifact bytecode (which instead has a zeroed
     * placeholder there, recorded in the artifact as `deployedBytecode.immutableReferences`).
     * This means two different, byte-for-byte-identical deployments of `ForwardLib` at
     * two different addresses legitimately have two different runtime codehashes — so a
     * raw `lib.codehash == keccak256(vm.getDeployedCode(...))` comparison would ALWAYS
     * fail for a correctly deployed library. This patches that self-address placeholder
     * to `lib`'s own address (mirroring what `lib`'s constructor did at deploy time)
     * before hashing, exactly analogous to how `_verifyLinkedImplementation` patches
     * `CrossBridge`'s `ForwardLib` link-reference offsets.
     */
    function _verifyLibCode(address lib) internal view {
        if (lib.code.length == 0) revert ForwardLibVerifyLibHasNoCode(lib);

        string memory json = vm.readFile(LIB_ARTIFACT_JSON_PATH);
        bytes memory template = vm.getDeployedCode(LIB_ARTIFACT);
        _patchAllImmutableReferences(template, json, ".deployedBytecode.immutableReferences", lib);

        bytes32 expected = keccak256(template);
        bytes32 actual = lib.codehash;
        if (actual != expected) revert ForwardLibVerifyLibCodehashMismatch(expected, actual);
    }

    /**
     * @dev Patches EVERY immutable slot recorded under the JSON object at `basePath`
     * (regardless of its key names — Solidity files these under an internal AST-node-id
     * or symbol name that varies by contract and compiler version) to `addr`,
     * RIGHT-aligned within each slot, exactly as an `address(this) == <constant>`
     * comparison compiles (the `ADDRESS` opcode pushes a right-aligned 32-byte word, so
     * the compiled comparison constant is right-aligned too — unlike the LEFT-aligned
     * raw 20-byte `PUSH20` operand used for external-library link references).
     *
     * Two independent things in this codebase compile to this pattern, both of which
     * must be patched before a wholesale runtime-hash comparison can succeed:
     *   - `ForwardLib`'s own compiler-inserted "reject direct CALL" guard, which compares
     *     `ADDRESS` against the library's own deployment address.
     *   - `CrossBridge`'s inherited `UUPSUpgradeable.__self` (`address(this)` captured
     *     as a Solidity `immutable`), used by its `onlyProxy`/`notDelegated` checks.
     * Both are ordinary Solidity/Yul immutables from the EVM's point of view; only the
     * caller decides which `addr` they should be patched to (the library's own address
     * for the first case, `impl`'s own address for the second).
     */
    function _patchAllImmutableReferences(
        bytes memory template,
        string memory json,
        string memory basePath,
        address addr
    ) internal view {
        if (!vm.keyExistsJson(json, basePath)) return; // nothing to patch; template already final.

        string[] memory names = vm.parseJsonKeys(json, basePath);
        for (uint n = 0; n < names.length; n++) {
            string memory entryPath = string.concat(basePath, ".", _jsonKey(names[n]));
            uint[2][] memory entries = abi.decode(vm.parseJson(json, entryPath), (uint[2][]));
            for (uint e = 0; e < entries.length; e++) {
                uint length = entries[e][0];
                uint start = entries[e][1];
                require(start + length <= template.length, "immutable reference out of range");
                _writeAddressRightAligned(template, start, length, addr);
            }
        }
    }

    /// @dev Checks (3) and (4) against `impl`'s ACTUAL on-chain runtime bytecode.
    function _verifyLinkedImplementation(address impl, address lib) internal view {
        string memory json = vm.readFile(HOST_ARTIFACT_PATH);

        // NOTE: deliberately NOT `vm.parseJsonBytes` here — on the installed forge-std
        // toolchain, `parseJsonBytes`/`vm.parseBytes` fail to parse a hex string this
        // long (~22KB, CrossBridge being close to the EIP-170 limit). Reading the raw
        // JSON string and hex-decoding it ourselves sidesteps that toolchain limitation.
        bytes memory unlinkedTemplate = _hexToBytes(vm.parseJsonString(json, ".deployedBytecode.object"));
        bytes memory implRuntime = impl.code;

        require(
            unlinkedTemplate.length == implRuntime.length,
            ForwardLibVerifyRuntimeLengthMismatch(unlinkedTemplate.length, implRuntime.length)
        );

        uint[] memory offsets = _collectForwardLibOffsets(json);

        for (uint i = 0; i < offsets.length; i++) {
            uint offset = offsets[i];
            if (offset + 20 > implRuntime.length) {
                revert ForwardLibVerifyLinkReferenceOutOfRange(offset, 20, implRuntime.length);
            }

            // (3) direct extraction check against the ACTUAL deployed runtime — proves
            // this specific relocation was patched to `lib`.
            address linked = _readAddress(implRuntime, offset);
            if (linked != lib) revert ForwardLibVerifyLinkedAddressMismatch(offset, lib, linked);

            // Patch the unlinked template in place for the wholesale check below.
            _writeAddress(unlinkedTemplate, offset, lib);
        }
        // offsets.length == 0 means CrossBridge was already fully linked at build time
        // (built with `--libraries`), so there are no relocations left to extract and
        // step (3)'s per-relocation proof cannot run. Recover the equivalent guarantee
        // from the artifact's own metadata: the compiler records the address it linked
        // against under `settings.libraries`, and that MUST be the library we are
        // verifying. Without this, step (4) would happily confirm that impl matches an
        // artifact linked to some OTHER library address.
        if (offsets.length == 0) _verifyPrelinkedDeclaration(json, lib);

        // CrossBridge also carries its OWN self-address immutable (inherited from
        // `UUPSUpgradeable`, see `_patchAllImmutableReferences`'s doc) — patch that to
        // `impl`'s own address too, or the wholesale hash below can never match.
        _patchAllImmutableReferences(unlinkedTemplate, json, ".deployedBytecode.immutableReferences", impl);

        // (4) wholesale check — the fully-patched template must hash to EXACTLY impl's
        // actual runtime codehash, proving nothing else in the bytecode diverged.
        bytes32 implCodehash = keccak256(implRuntime);
        bytes32 patchedCodehash = keccak256(unlinkedTemplate);
        if (patchedCodehash != implCodehash) {
            revert ForwardLibVerifyPatchedRuntimeMismatch(implCodehash, patchedCodehash);
        }
    }

    /**
     * @dev Verifies a pre-linked artifact declares `lib` as the address it was linked
     * against. Used only when `deployedBytecode.linkReferences` is empty, which happens
     * when the artifact was compiled with `--libraries`/`foundry.toml` `libraries`.
     *
     * The compiler records this under `metadata.settings.libraries` keyed by the
     * fully-qualified library name.
     */
    function _verifyPrelinkedDeclaration(string memory json, address lib) internal view {
        string memory path = string.concat(".metadata.settings.libraries", _jsonKey(LIB_FQN));

        if (!vm.keyExistsJson(json, path)) revert ForwardLibVerifyPrelinkedLibraryNotDeclared();

        address declared = vm.parseAddress(vm.parseJsonString(json, path));
        if (declared != lib) revert ForwardLibVerifyPrelinkedAddressMismatch(lib, declared);
    }

    /**
     * @dev Reads every `{start, length}` entry filed under `ForwardLib` in the build
     * artifact's `deployedBytecode.linkReferences`, across whichever source-file key the
     * compiler filed it under. Reverts if any link reference is filed under a name other
     * than `ForwardLib` (an unexpected external library would mean this script's
     * assumptions are stale) or if none is found (`length == 0`) is the caller's signal
     * that `CrossBridge` was already fully linked at build time.
     */
    function _collectForwardLibOffsets(string memory json) internal pure returns (uint[] memory offsets) {
        string memory basePath = ".deployedBytecode.linkReferences";
        string[] memory files = vm.parseJsonKeys(json, basePath);

        uint total;
        uint[][] memory perFile = new uint[][](files.length);

        for (uint f = 0; f < files.length; f++) {
            string memory contractsPath = string.concat(basePath, _jsonKey(files[f]));
            string[] memory names = vm.parseJsonKeys(json, contractsPath);

            for (uint c = 0; c < names.length; c++) {
                if (keccak256(bytes(names[c])) != keccak256(bytes(EXPECTED_LIB_NAME))) {
                    revert ForwardLibVerifyUnexpectedLinkReferenceContract(files[f], names[c]);
                }
            }

            if (names.length == 0) continue;

            string memory entryPath = string.concat(contractsPath, ".", EXPECTED_LIB_NAME);
            // Each entry is `{length, start}` (alphabetical field order, matching
            // Foundry's JSON-to-ABI struct decoding convention); layout-equivalent to
            // `uint[2][]` since both fields are static uints.
            uint[2][] memory entries = abi.decode(vm.parseJson(json, entryPath), (uint[2][]));

            uint[] memory fileOffsets = new uint[](entries.length);
            for (uint e = 0; e < entries.length; e++) {
                // entries[e] = [length, start]; ForwardLib addresses are always 20 bytes.
                fileOffsets[e] = entries[e][1];
            }
            perFile[f] = fileOffsets;
            total += fileOffsets.length;
        }

        // total == 0 is NOT an error: it means the artifact was compiled with
        // `--libraries` (or a `libraries` entry in foundry.toml), so the compiler already
        // baked the address in and emitted no relocations. `_verifyLinkedImplementation`
        // detects the empty result and switches to the pre-linked verification path.
        if (total == 0) return new uint[](0);

        offsets = new uint[](total);
        uint w;
        for (uint f = 0; f < perFile.length; f++) {
            for (uint e = 0; e < perFile[f].length; e++) {
                offsets[w++] = perFile[f][e];
            }
        }
    }

    /// @dev `vm.parseJsonKeys`'s path segments must be escaped for keys containing `.` or
    /// `/` (source file paths do); wrapping in single quotes is the jq/forge-std convention.
    function _jsonKey(string memory key) internal pure returns (string memory) {
        return string.concat("['", key, "']");
    }

    /// @dev Reads a 20-byte address out of `data` starting at byte offset `offset`.
    function _readAddress(bytes memory data, uint offset) internal pure returns (address addr) {
        assembly ("memory-safe") {
            addr := shr(96, mload(add(add(data, 32), offset)))
        }
    }

    /// @dev Overwrites 20 bytes of `data` starting at byte offset `offset` with `addr`.
    function _writeAddress(bytes memory data, uint offset, address addr) internal pure {
        assembly ("memory-safe") {
            let ptr := add(add(data, 32), offset)
            let word := mload(ptr)
            // Keep only the bottom 12 bytes of the existing word (0xff * 12, i.e. 24 hex
            // digits — zero-extended by the assembler to a full 32-byte value), then OR
            // in `addr` left-aligned into the top 20 bytes.
            mstore(ptr, or(and(word, 0xffffffffffffffffffffffff), shl(96, addr)))
        }
    }

    /**
     * @dev Overwrites a `slotLength`-byte window of `data` starting at byte offset
     * `offset` with `addr`, RIGHT-aligned within that window (upper bytes zeroed) — the
     * layout an `ADDRESS == <constant>` comparison compiles to, as opposed to
     * `_writeAddress`'s LEFT-aligned raw `PUSH20` operand layout. Implemented with plain
     * indexed byte writes (not assembly) since this only ever runs over a handful of
     * bytes and correctness matters far more than gas here.
     */
    function _writeAddressRightAligned(bytes memory data, uint offset, uint slotLength, address addr) internal pure {
        require(slotLength >= 20 && slotLength <= 32, "bad immutable slot length");
        uint zeroBytes = slotLength - 20;
        for (uint i = 0; i < zeroBytes; i++) {
            data[offset + i] = 0;
        }
        bytes20 packed = bytes20(addr);
        for (uint i = 0; i < 20; i++) {
            data[offset + zeroBytes + i] = packed[i];
        }
    }

    /**
     * @dev Decodes a `0x`-prefixed hex string into `bytes`, implemented by hand instead
     * of via `vm.parseBytes`/`vm.parseJsonBytes` — on the installed forge-std toolchain
     * those cheatcodes fail on a hex string as long as a full contract's runtime
     * bytecode (see the call site in `_verifyLinkedImplementation`).
     *
     * Deliberately LENIENT on unrecognized characters (decoded as the nibble `0`)
     * rather than reverting: Solidity's UNLINKED artifact bytecode represents each
     * external-library relocation not as zero bytes but as a `__$<hash>$__`-style
     * placeholder string, which is not valid hex. Every one of those placeholder
     * regions is later fully overwritten by `_writeAddress` at the exact offsets
     * `_collectForwardLibOffsets` reads from the SAME artifact's `linkReferences`, so
     * decoding the placeholder as zero bytes here (rather than rejecting it) is safe:
     * the caller must patch those bytes before drawing any conclusion from them, and
     * `_verifyLinkedImplementation`'s length check still catches a genuinely malformed
     * artifact.
     */
    function _hexToBytes(string memory s) internal pure returns (bytes memory out) {
        bytes memory b = bytes(s);
        uint start = (b.length >= 2 && b[0] == "0" && (b[1] == "x" || b[1] == "X")) ? 2 : 0;
        require((b.length - start) % 2 == 0, "odd-length hex string");
        uint len = (b.length - start) / 2;
        out = new bytes(len);
        for (uint i = 0; i < len; i++) {
            out[i] = bytes1(_hexNibble(b[start + 2 * i]) * 16 + _hexNibble(b[start + 2 * i + 1]));
        }
    }

    /// @dev Decodes a single ASCII hex character (`0-9`, `a-f`, `A-F`) to its 0-15 value;
    /// any other character (e.g. from a `__$...$__` link placeholder) decodes to `0`.
    function _hexNibble(bytes1 c) internal pure returns (uint8) {
        uint8 ch = uint8(c);
        if (ch >= 0x30 && ch <= 0x39) return ch - 0x30; // '0'-'9'
        if (ch >= 0x61 && ch <= 0x66) return ch - 0x61 + 10; // 'a'-'f'
        if (ch >= 0x41 && ch <= 0x46) return ch - 0x41 + 10; // 'A'-'F'
        return 0;
    }
}
