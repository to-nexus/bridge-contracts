// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {Test, console} from "forge-std/Test.sol";

contract SigTest is Test {
    bytes32 hash;
    uint8 v;
    bytes32 r;
    bytes32 s;
    bytes sig;

    uint length = 50;
    bytes32[] hash_arr = new bytes32[](length);
    uint8[] v_arr = new uint8[](length);
    bytes32[] r_arr = new bytes32[](length);
    bytes32[] s_arr = new bytes32[](length);
    bytes[] sig_arr = new bytes[](length);

    uint internal constant OWNER_PK = uint(bytes32("owner"));
    address public OWNER;

    function setUp() public {
        OWNER = vm.addr(OWNER_PK);

        hash = bytes32("hash");
        (v, r, s) = vm.sign(OWNER_PK, hash);
        sig = abi.encodePacked(r, s, v);

        for (uint i = 0; i < length; ++i) {
            hash_arr[i] = hash;
            v_arr[i] = v;
            r_arr[i] = r;
            s_arr[i] = s;
            sig_arr[i] = sig;
        }
    }

    /**
     * [PASS] test_sigarray() (gas: 20222)
     * [PASS] test_sigarray_many() (gas: 802639)
     * [PASS] test_sigvrs() (gas: 17380)
     * [PASS] test_sigvrs_many() (gas: 578317)
     */
    function test_sigvrs() public view {
        sigvrs(hash, v, r, s);
    }

    function test_sigarray() public view {
        sigarray(hash, sig);
    }

    function test_sigvrs_many() public view {
        sigvrsBatch(hash_arr, v_arr, r_arr, s_arr);
    }

    function test_sigarray_many() public view {
        sigarrayBatch(hash_arr, sig_arr);
    }

    function sigvrs(bytes32 _hash, uint8 _v, bytes32 _r, bytes32 _s) public view {
        address recovered = ECDSA.recover(_hash, _v, _r, _s);
        assertEq(OWNER, recovered);
    }

    function sigarray(bytes32 _hash, bytes memory _sig) public view {
        address recovered = ECDSA.recover(_hash, _sig);
        assertEq(OWNER, recovered);
    }

    function sigvrsBatch(bytes32[] memory _hash, uint8[] memory _v, bytes32[] memory _r, bytes32[] memory _s)
        public
        view
    {
        for (uint i = 0; i < _hash.length; ++i) {
            sigvrs(_hash[i], _v[i], _r[i], _s[i]);
        }
    }

    function sigarrayBatch(bytes32[] memory _hash, bytes[] memory _sig) public view {
        for (uint i = 0; i < _hash.length; ++i) {
            sigarray(_hash[i], _sig[i]);
        }
    }
}
