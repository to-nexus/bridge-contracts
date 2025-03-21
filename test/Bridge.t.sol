// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {EthereumChainTest} from "./chain/EthereumChain.t.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {console} from "forge-std/Test.sol";

contract BridgeTest is EthereumChainTest {
    bool enableGasUsedLog;

    function setUp() public virtual override {
        super.setUp();
    }

    // cross token bridge (ethereum chain -> cross chain)
    function deposit(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        vm.selectFork(ethereumForkID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        bool ok;
        (index, ok) = ethereumBridge(address(cross), USER, USER, amount, 0, 0);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
            ethereumIncrementIndex();
            ok = crossFinalize(index, address(NATIVE_TOKEN), USER, amount, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize coin", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(cross));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(ethereumForkID);
            assertEq(userTokenBalance - amount, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossForkID);
            assertEq(userCoinBalance + amount, USER.balance);
            assertEq(bridgeCoinBalance - amount, address(bridgeCross).balance);
        }
    }

    // xcross bridge (cross chain -> ethereum chain)
    function withdraw(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        (uint value, uint gas, uint ex) = crossCalcFee(NATIVE_TOKEN, amount);
        assertTrue(value + gas + ex <= amount);

        vm.selectFork(ethereumForkID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;
        uint rewardWalletBalance = REWARD.balance;

        assertTrue(value > 0);

        bool ok;
        (index, ok) = crossBridge(address(NATIVE_TOKEN), USER, USER, value, gas, ex);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge coin", uint(vm.lastCallGas().gasTotalUsed));
            crossIncrementIndex();
            ok = ethereumFinalize(index, address(cross), USER, value, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(cross));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(ethereumForkID);
            assertEq(userTokenBalance + value, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance - value, cross.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossForkID);
            assertEq(userCoinBalance - (value + gas + ex), USER.balance);
            assertEq(bridgeCoinBalance + value, address(bridgeCross).balance);
            assertEq(rewardWalletBalance + gas + ex, REWARD.balance);
        }
    }

    // erc20 token bridge (ethereum chain -> cross chain)
    function depositToken(bool isRevert, uint amount, uint validatorNum) public {
        vm.selectFork(ethereumForkID);
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        (uint index, bool ok) = ethereumBridge(address(testTokenEthereum), USER, USER, amount, 0, 0);
        if (ok) {
            ethereumIncrementIndex();
            crossFinalize(index, address(testTokenCross), USER, amount, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize token (mint)", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(testTokenCross));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(testTokenEthereum));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(ethereumForkID);
            assertEq(userEthereumBalance - amount, testTokenEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance + amount, testTokenEthereum.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossForkID);
            assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
        }
    }

    // erc20 token bridge (cross chain -> ethereum chain)
    function withdrawToken(bool isRevert, uint amount, uint validatorNum) public {
        (uint value, uint gas, uint ex) = crossCalcFee(IERC20(address(testTokenCross)), amount);
        assertTrue(value + gas + ex <= amount);
        uint total = value + gas + ex;
        assertTrue(total <= amount);

        vm.selectFork(ethereumForkID);
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);
        uint rewardWalletBalance = testTokenCross.balanceOf(REWARD);

        (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, ex);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge token (burn)", uint(vm.lastCallGas().gasTotalUsed));
            crossIncrementIndex();
            ethereumFinalize(index, address(testTokenEthereum), USER, value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(testTokenCross));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(testTokenEthereum));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(ethereumForkID);
            assertEq(userEthereumBalance + value, testTokenEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance - value, testTokenEthereum.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossForkID);
            assertEq(userCrossBalance - total, testTokenCross.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + ex, testTokenCross.balanceOf(REWARD));
        }
    }

    // eth NATIVE_TOKEN bridge (ethereum chain -> cross chain)
    function depositETH(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        vm.selectFork(crossForkID);
        uint userTokenBalance = weth.balanceOf(USER);
        vm.selectFork(ethereumForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeEthereum).balance;

        bool ok;
        (index, ok) = ethereumBridge(address(NATIVE_TOKEN), USER, USER, amount, 0, 0);
        if (ok) {
            ethereumIncrementIndex();
            crossFinalize(index, address(weth), USER, amount, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(weth));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(NATIVE_TOKEN));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(crossForkID);
            assertEq(userTokenBalance + amount, weth.balanceOf(USER));
            vm.selectFork(ethereumForkID);
            assertEq(userCoinBalance - amount, USER.balance);
            assertEq(bridgeCoinBalance + amount, address(bridgeEthereum).balance);
        }
    }

    // eth NATIVE_TOKEN bridge (cross chain -> ethereum chain)
    function withdrawETH(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        (uint value, uint gas, uint ex) = crossCalcFee(weth, amount);
        assertTrue(value + gas + ex <= amount);

        vm.selectFork(crossForkID);
        uint userTokenBalance = weth.balanceOf(USER);
        uint rewardWalletBalance = weth.balanceOf(REWARD);
        vm.selectFork(ethereumForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeEthereum).balance;

        bool ok;
        (index, ok) = crossBridge(address(weth), USER, USER, value, gas, ex);
        if (ok) {
            crossIncrementIndex();
            ethereumFinalize(index, address(NATIVE_TOKEN), USER, value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            uint crossBridged = bridgeCross.bridgeNetQty(ETHEREUM_CHAIN_ID, address(weth));
            vm.selectFork(ethereumForkID);
            uint ethereumBridged = bridgeEthereum.bridgeNetQty(CROSS_CHAIN_ID, address(NATIVE_TOKEN));
            assertEq(crossBridged, ethereumBridged);

            vm.selectFork(crossForkID);
            assertEq(userTokenBalance - (value + gas + ex), weth.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + ex, weth.balanceOf(REWARD));
            vm.selectFork(ethereumForkID);
            assertEq(userCoinBalance + value, USER.balance);
            assertEq(bridgeCoinBalance - value, address(bridgeEthereum).balance);
        }
    }
}
