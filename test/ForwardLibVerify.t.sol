// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";

import {ForwardLibVerify} from "../script/ForwardLibVerify.s.sol";
import {CrossBridgeV2} from "../src/CrossBridgeV2.sol";
import {ForwardLib} from "../src/lib/ForwardLib.sol";

/**
 * @title ForwardLibVerifyHarness
 * @notice Inherits `ForwardLibVerify` SOLELY to reuse its internal bytecode-construction
 * helpers (`_hexToBytes`, `_collectForwardLibOffsets`, `_writeAddress`,
 * `_patchAllImmutableReferences`) for building synthetic test scenarios below. The thing
 * actually under test in every assertion is a SEPARATE, plain `ForwardLibVerify`
 * instance's `verify(...)` — never this harness's own (inherited but unused) copy.
 */
contract ForwardLibVerifyHarness is ForwardLibVerify {
    /// @dev Builds the REAL, correctly-linked `CrossBridgeV2` runtime bytecode a genuine
    /// deployment at `implAddr` linked against `libAddr` would have: the build
    /// artifact's own template, with every `ForwardLib` link-reference offset patched to
    /// `libAddr` and `CrossBridgeV2`'s own `UUPSUpgradeable.__self` immutable patched to
    /// `implAddr`.
    function buildGenuineImplRuntime(address implAddr, address libAddr) external view returns (bytes memory) {
        string memory json = vm.readFile(HOST_ARTIFACT_PATH);
        bytes memory template = _hexToBytes(vm.parseJsonString(json, ".deployedBytecode.object"));

        uint[] memory offsets = _collectForwardLibOffsets(json);
        for (uint i = 0; i < offsets.length; i++) {
            _writeAddress(template, offsets[i], libAddr);
        }
        _patchAllImmutableReferences(template, json, ".deployedBytecode.immutableReferences", implAddr);
        return template;
    }

    /// @dev Like `buildGenuineImplRuntime`, but everything OUTSIDE the link-reference and
    /// self-address-immutable slots is overwritten with `filler` instead of copied from
    /// the real artifact — i.e. "the right address sits at every relocation the compiler
    /// declared, but the rest of the code is unrelated". Used to prove the wholesale
    /// runtime-hash check (not just a per-offset address check) is load-bearing.
    function buildCoincidentalImplRuntime(address implAddr, address libAddr, bytes1 filler)
        external
        view
        returns (bytes memory runtime)
    {
        string memory json = vm.readFile(HOST_ARTIFACT_PATH);
        bytes memory template = _hexToBytes(vm.parseJsonString(json, ".deployedBytecode.object"));

        runtime = new bytes(template.length);
        for (uint i = 0; i < runtime.length; i++) {
            runtime[i] = filler;
        }

        uint[] memory offsets = _collectForwardLibOffsets(json);
        for (uint i = 0; i < offsets.length; i++) {
            _writeAddress(runtime, offsets[i], libAddr);
        }
        _patchAllImmutableReferences(runtime, json, ".deployedBytecode.immutableReferences", implAddr);
    }

    /// @dev Reconstructs the runtime bytecode a genuine `ForwardLib` deployment at
    /// `libSelfAddr` would have: the library's own build artifact template, with its
    /// compiler-inserted "reject direct CALL" self-address guard patched to
    /// `libSelfAddr`. Mirrors `_verifyLibCode`'s internal expectation exactly, so tests
    /// can assert `ForwardLibVerify`'s exact revert data against it.
    function buildPatchedLibRuntime(address libSelfAddr) external view returns (bytes memory template) {
        string memory json = vm.readFile(LIB_ARTIFACT_JSON_PATH);
        template = vm.getDeployedCode(LIB_ARTIFACT);
        _patchAllImmutableReferences(template, json, ".deployedBytecode.immutableReferences", libSelfAddr);
    }

    /// @dev M-1: reads the ACTUALLY-linked `ForwardLib` address straight out of `impl`'s
    /// real deployed runtime bytecode, at the offset(s) the `CrossBridgeV2` build
    /// artifact declares as `ForwardLib` link references — the address is discovered,
    /// never assumed by the caller. Reuses `verify(...)`'s own offset-discovery logic
    /// (`_collectForwardLibOffsets`/`_readAddress`) so this is exactly what `verify`
    /// itself would read, not a second, potentially-diverging implementation of it.
    function discoverLinkedLibAddress(address impl) external view returns (address addr) {
        string memory json = vm.readFile(HOST_ARTIFACT_PATH);
        uint[] memory offsets = _collectForwardLibOffsets(json);
        require(offsets.length > 0, "no ForwardLib link references found in the build artifact");
        addr = _readAddress(impl.code, offsets[0]);
    }
}

/**
 * @title ForwardLibVerifyTest
 * @notice Self-tests for `script/ForwardLibVerify.s.sol`, per plan spec §7.5 / issue
 * plan H-1 item 4: the verifier itself needs unit coverage, not just a happy-path smoke
 * run, since a bug in the verifier is exactly as dangerous as no verifier at all.
 * @dev Four required cases: (1) genuinely correct link -> passes, (2) wrong library
 * address -> rejected, (3) right library address but different code deployed there ->
 * rejected (codehash mismatch), (4) an implementation that happens to contain the
 * expected address at the right offsets but is otherwise unrelated bytecode -> rejected
 * (proves the wholesale-hash check, not just the per-offset check, is load-bearing).
 */
contract ForwardLibVerifyTest is Test {
    ForwardLibVerify internal verifier;
    ForwardLibVerifyHarness internal harness;
    address internal lib;

    function setUp() public {
        verifier = new ForwardLibVerify();
        harness = new ForwardLibVerifyHarness();
        lib = _deployForwardLib();
    }

    /// @dev `ForwardLib` is a Solidity `library`, so it cannot be instantiated with
    /// `new ForwardLib()` — the compiler rejects `new` for library types even though
    /// external libraries ARE ordinary deployed contracts under the hood. Deploying via
    /// raw `CREATE` of `type(ForwardLib).creationCode` sidesteps that restriction.
    function _deployForwardLib() internal returns (address addr) {
        bytes memory bytecode = type(ForwardLib).creationCode;
        assembly {
            addr := create(0, add(bytecode, 32), mload(bytecode))
        }
        require(addr != address(0), "ForwardLib deployment failed");
    }

    /// @notice M-1: a REAL, genuinely compiled-and-linked `CrossBridgeV2` (not a
    /// synthetic runtime built from the verifier's own patching helpers) must pass
    /// `verify(...)` against its ACTUALLY linked `ForwardLib` address, discovered from
    /// the artifact-declared relocation offsets rather than assumed by the test.
    /// @dev `test_verify_correctlyLinked_passes` below is circular: it feeds `verify`
    /// exactly the kind of runtime `verify` itself expects, built by the SAME patching
    /// helpers `verify` relies on, so it cannot prove the immutable-patching assumptions
    /// (or forge's actual linker/deployment output) are correct in the first place. This
    /// test closes that gap with a genuinely deployed host contract.
    function test_verify_realDeployedImplementation_passes() public {
        CrossBridgeV2 impl = new CrossBridgeV2();
        address actualLib = harness.discoverLinkedLibAddress(address(impl));
        assertTrue(actualLib.code.length > 0, "discovered library address must have code");

        // Must not revert.
        verifier.verify(address(impl), actualLib);
    }

    /// @notice (1) Correctly linked implementation + correct library address -> passes.
    function test_verify_correctlyLinked_passes() public {
        address implAddr = makeAddr("cross-bridge-v2-impl");
        bytes memory runtime = harness.buildGenuineImplRuntime(implAddr, address(lib));
        vm.etch(implAddr, runtime);

        // Must not revert.
        verifier.verify(implAddr, address(lib));
    }

    /// @notice (2) Implementation is genuinely linked against `lib`, but the caller
    /// passes a DIFFERENT address as the expected library -> rejected.
    function test_verify_wrongLibraryAddress_reverts() public {
        address implAddr = makeAddr("cross-bridge-v2-impl-2");
        bytes memory runtime = harness.buildGenuineImplRuntime(implAddr, address(lib));
        vm.etch(implAddr, runtime);

        address wrongLib = _deployForwardLib();
        assertTrue(wrongLib != lib, "test requires two distinct library addresses");

        vm.expectRevert();
        verifier.verify(implAddr, wrongLib);
    }

    /// @notice (3) The implementation IS linked against `claimedLib`'s address, and that
    /// address has code — but the code deployed there is NOT `ForwardLib` (different
    /// codehash). Must be rejected by the library-codehash check, independent of
    /// whether the implementation's link references "look" consistent.
    function test_verify_libraryAddressHasWrongCode_reverts() public {
        // An address with SOME code, but not ForwardLib's.
        address claimedLib = makeAddr("not-actually-forward-lib");
        vm.etch(claimedLib, hex"6080604052600080fd"); // trivial non-ForwardLib runtime

        address implAddr = makeAddr("cross-bridge-v2-impl-3");
        bytes memory runtime = harness.buildGenuineImplRuntime(implAddr, claimedLib);
        vm.etch(implAddr, runtime);

        bytes32 expectedLibCodehash = keccak256(harness.buildPatchedLibRuntime(claimedLib));

        vm.expectRevert(
            abi.encodeWithSelector(
                ForwardLibVerify.ForwardLibVerifyLibCodehashMismatch.selector, expectedLibCodehash, claimedLib.codehash
            )
        );
        verifier.verify(implAddr, claimedLib);
    }

    /// @notice (4) `libAddr`'s bytes sit at every offset the compiler declared as a
    /// `ForwardLib` relocation, and the self-address immutable is correctly patched too
    /// — but everything else in the "implementation" is arbitrary filler, not real
    /// compiled `CrossBridgeV2` code. A verifier that only checked the 20 bytes at each
    /// offset (a pattern scan) would be fooled by this; the wholesale runtime-hash check
    /// must still reject it.
    function test_verify_coincidentalAddressMatch_reverts() public {
        address implAddr = makeAddr("cross-bridge-v2-impl-4");
        bytes memory runtime = harness.buildCoincidentalImplRuntime(implAddr, address(lib), 0xAB);
        vm.etch(implAddr, runtime);

        try verifier.verify(implAddr, address(lib)) {
            fail();
        } catch (bytes memory reason) {
            assertEq(
                bytes4(reason),
                ForwardLibVerify.ForwardLibVerifyPatchedRuntimeMismatch.selector,
                "must be rejected specifically by the wholesale runtime-hash check"
            );
        }
    }
}
