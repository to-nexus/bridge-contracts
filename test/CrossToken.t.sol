// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {CrossToken} from "./token/CrossToken.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test, console} from "forge-std/Test.sol";

contract CrossTokenTest is Test {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    uint internal constant OWNER_PK = uint(bytes32("owner"));
    uint internal constant USER_PK = uint(bytes32("user"));

    address public OWNER;
    address public USER;
    CrossToken public cross;
    TestToken public ttoken;

    function setUp() public {
        OWNER = vm.addr(OWNER_PK);
        USER = vm.addr(USER_PK);
        cross = new CrossToken(OWNER, 1_000_000_000);
        ttoken = new TestToken("Test Token", "TT", 18);

        vm.prank(OWNER);
        ttoken.mint(OWNER, 1_000_000_000 * 1e18);
    }

    function test_permitTransfer() public {
        // create permit data
        address from = OWNER;
        address to = USER;
        uint value = 100000 * 1e18;
        uint deadline = type(uint).max;
        uint nonce;
        uint8 v;
        bytes32 r;
        bytes32 s;
        {
            // create permit data
            nonce = cross.nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, from, to, value, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(cross.DOMAIN_SEPARATOR(), h);
            (v, r, s) = vm.sign(OWNER_PK, hash);
        }

        {
            vm.startPrank(USER);
            cross.permit(from, to, value, deadline, v, r, s);
            cross.transferFrom(from, to, value);
            vm.stopPrank();

            assertEq(value, cross.balanceOf(USER));
        }
    }

    function test_permitTransfer2() public {
        // create permit data
        address from = OWNER;
        address to = USER;
        uint value = 100000 * 1e18;
        uint deadline = type(uint).max;
        uint nonce;
        uint8 v;
        bytes32 r;
        bytes32 s;
        {
            // create permit data
            nonce = ttoken.nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, from, to, value, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(ttoken.DOMAIN_SEPARATOR(), h);
            (v, r, s) = vm.sign(OWNER_PK, hash);
        }

        {
            vm.startPrank(USER);
            ttoken.permit(from, to, value, deadline, v, r, s);
            ttoken.transferFrom(from, to, value);
            vm.stopPrank();

            assertEq(value, ttoken.balanceOf(USER));
        }
    }
}
