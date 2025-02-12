// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeEthereum} from "../src/BridgeEthereum.sol";
import {BridgeStandardTest} from "./BridgeStandard.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test, console} from "forge-std/Test.sol";

contract BridgeTest is Test, BridgeStandardTest {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    // ----- Test -----
    function test_depositWithdraw() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        deposit(false, amount, 5);
        withdraw(false, amount * 10, 5);
    }

    function test_depositWithdrawToken() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenAtEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenAtEthereum.approve(address(bridgeEthereum.b), amount);
        depositToken(false, amount, 5);

        vm.prank(USER);
        testTokenAtCross.approve(address(bridgeCross.b), amount);
        withdrawToken(false, amount, 5);
    }

    function test_reverted_finalize() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenAtEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenAtEthereum.approve(address(bridgeEthereum.b), amount);

        depositToken(true, amount, 5); // should finalize revert

        vm.prank(USER);
        testTokenAtCross.approve(address(bridgeCross.b), amount);

        {
            // token stop
            vm.prank(OWNER);
            TestToken(address(testTokenAtEthereum)).stop();

            (uint value, uint gas, uint service) = calcFee(IERC20(address(testTokenAtCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = bridge(bridgeCross, address(testTokenAtCross), value, gas, service);
            assertTrue(ok);
            bridgeCross.nextIndex++;
            finalize(bridgeCross.otherBridge, index, address(testTokenAtEthereum), value, 5);

            bytes memory reason = bridgeEthereum.b.revertedReason(index);
            assertNotEq(bytes(""), reason);

            // token start
            vm.prank(OWNER);
            TestToken(address(testTokenAtEthereum)).start();

            uint before = testTokenAtEthereum.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeEthereum.b.retryFinalize(index);

            assertEq(before + value, testTokenAtEthereum.balanceOf(USER));
        }
    }

    function test_permit_deposit() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeEthereum.b)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum.b));
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross.b).balance;

        uint deadline = type(uint).max;
        uint8 v;
        bytes32 r;
        bytes32 s;
        uint index;
        bool ok;
        {
            // make permit sig
            index = bridgeEthereum.nextIndex;

            uint nonce = cross.nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum.b), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(cross.DOMAIN_SEPARATOR(), h);
            (v, r, s) = vm.sign(USER_PK, hash);
        }

        vm.prank(VALIDATOR1);
        ok = BridgeEthereum(address(bridgeEthereum.b)).bridgeCross(
            USER, amount, deadline, abi.encodePacked(r, s, v), NULLDATA
        );
        assertTrue(ok);

        bridgeEthereum.nextIndex++;
        finalize(bridgeEthereum.otherBridge, index, address(xcross), amount, 5);

        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum.b)));
        assertEq(userCoinBalance + (amount * exrate), USER.balance);
        assertEq(bridgeCoinBalance - (amount * exrate), address(bridgeCross.b).balance);
    }

    function test_permit_depositToken() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenAtEthereum.transfer(USER, amount);
        assertTrue(testTokenAtEthereum.allowance(USER, address(bridgeEthereum.b)) == 0);

        // initiate
        uint userEthereumBalance = testTokenAtEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenAtEthereum.balanceOf(address(bridgeEthereum.b));
        uint userCrossBalance = testTokenAtCross.balanceOf(USER);

        uint index;
        bytes memory permitSig;
        {
            uint8 v;
            bytes32 r;
            bytes32 s;
            // make permit sig
            index = bridgeEthereum.nextIndex;

            uint nonce = IERC20Permit(address(testTokenAtEthereum)).nonces(USER);
            bytes32 h =
                keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum.b), amount, nonce, type(uint).max));
            bytes32 hash =
                MessageHashUtils.toTypedDataHash(IERC20Permit(address(testTokenAtEthereum)).DOMAIN_SEPARATOR(), h);
            (v, r, s) = vm.sign(USER_PK, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        vm.prank(VALIDATOR1);
        bool ok = BridgeEthereum(address(bridgeEthereum.b)).permitBridge(
            testTokenAtEthereum, USER, amount, 0, 0, type(uint).max, permitSig, NULLDATA
        );
        assertTrue(ok);

        bridgeEthereum.nextIndex++;
        finalize(bridgeEthereum.otherBridge, index, address(testTokenAtCross), amount, 5);

        assertEq(userEthereumBalance - amount, testTokenAtEthereum.balanceOf(USER));
        assertEq(bridgeEthereumBalance + amount, testTokenAtEthereum.balanceOf(address(bridgeEthereum.b)));
        assertEq(userCrossBalance + amount, testTokenAtCross.balanceOf(USER));
    }
}
