// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IBridgeStandard} from "../src/interface/IBridgeStandard.sol";

import {BridgeTest} from "./Bridge.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract BridgeStandardTest is BridgeTest {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    // ----- Test -----
    function test_depositWithdraw() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);
        withdraw(false, amount * 10, 5);
    }

    function test_depositWithdraw_eth() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumChainID);
        vm.deal(USER, amount);
        vm.selectFork(crossChainID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        depositETH(false, amount, 5);
        withdrawETH(false, amount, 5);
    }

    function test_depositWithdrawToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenEthereum.approve(address(bridgeEthereum), amount);
        depositToken(false, amount, 5);

        vm.selectFork(crossChainID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);
        withdrawToken(false, amount, 5);
    }

    function test_reverted_finalize() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenEthereum.approve(address(bridgeEthereum), amount);

        depositToken(true, amount, 5); // should finalize revert

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);

        {
            // token stop
            vm.selectFork(ethereumChainID);
            vm.prank(OWNER);
            TestToken(address(testTokenEthereum)).stop();

            (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = crossBridge(address(testTokenCross), value, gas, service);
            assertTrue(ok);
            crossIncrementIndex();
            ethereumFinalize(index, address(testTokenEthereum), value, 5);

            vm.selectFork(ethereumChainID);
            bytes memory reason = bridgeEthereum.revertedReason(index);
            assertNotEq(bytes(""), reason);

            // token start
            vm.prank(OWNER);
            vm.selectFork(ethereumChainID);
            TestToken(address(testTokenEthereum)).start();

            uint before = testTokenEthereum.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeEthereum.retryFinalize(index);

            assertEq(before + value, testTokenEthereum.balanceOf(USER));
        }
    }

    function test_permit_deposit() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeEthereum)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBridgeStandard.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumChainID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBridgeStandard.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        vm.prank(VALIDATOR1);
        assertTrue(bridgeEthereum.bridgeCross(USER, amount, permitArgs, NULLDATA));

        crossFinalize(nextIndexCross, address(coin), amount, 5);
        ethereumIncrementIndex();

        vm.selectFork(ethereumChainID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum)));
        vm.selectFork(crossChainID);
        assertEq(userCoinBalance + (amount * exrate), USER.balance);
        assertEq(bridgeCoinBalance - (amount * exrate), address(bridgeCross).balance);
    }

    function test_permit_depositToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumChainID);
        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        assertTrue(testTokenEthereum.allowance(USER, address(bridgeEthereum)) == 0);

        // initiate
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        uint index;
        IBridgeStandard.PermitArguments memory permitArgs;
        {
            // make permit sig
            index = nextIndexCross;

            vm.selectFork(ethereumChainID);
            uint nonce = IERC20Permit(address(testTokenEthereum)).nonces(USER);
            bytes32 h =
                keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, type(uint).max));
            bytes32 hash =
                MessageHashUtils.toTypedDataHash(IERC20Permit(address(testTokenEthereum)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);

            permitArgs = IBridgeStandard.PermitArguments(
                IERC20Permit(address(testTokenEthereum)), USER, amount, type(uint).max, v, r, s
            );
        }

        vm.prank(VALIDATOR1);
        bool ok = bridgeEthereum.permitBridge(testTokenEthereum, USER, amount, 0, 0, permitArgs, NULLDATA);
        assertTrue(ok);

        ethereumIncrementIndex();
        crossFinalize(index, address(testTokenCross), amount, 5);

        vm.selectFork(ethereumChainID);
        assertEq(userEthereumBalance - amount, testTokenEthereum.balanceOf(USER));
        assertEq(bridgeEthereumBalance + amount, testTokenEthereum.balanceOf(address(bridgeEthereum)));
        vm.selectFork(crossChainID);
        assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
    }
}
