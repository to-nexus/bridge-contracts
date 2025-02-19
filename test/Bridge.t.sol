// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {EthereumChainTest} from "./chain/EthereumChain.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BridgeTest is EthereumChainTest {
    function setUp() public override {
        super.setUp();
    }

    // cross token bridge (ethereum chain -> cross chain)
    function deposit(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        vm.selectFork(ethereumChainID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        bool ok;
        (index, ok) = ethereumBridge(address(cross), amount, 0, 0);
        if (ok) {
            ethereumIncrementIndex();
            crossFinalize(index, address(xcross), amount, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(ethereumChainID);
            assertEq(userTokenBalance - amount, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossChainID);
            assertEq(userCoinBalance + (amount * exrate), USER.balance);
            assertEq(bridgeCoinBalance - (amount * exrate), address(bridgeCross).balance);
        }
    }

    // xcross bridge (cross chain -> ethereum chain)
    function withdraw(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        assertTrue(amount / exrate > 0);
        (uint value, uint gas, uint service) = crossCalcFee(xcross, amount);
        assertTrue(value + gas + service <= amount);

        vm.selectFork(ethereumChainID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;
        uint rewardWalletBalance = REWARD.balance;

        value = (value / exrate) * exrate;
        assertTrue(value > 0);

        bool ok;
        (index, ok) = crossBridge(address(xcross), value, gas, service);
        if (ok) {
            crossIncrementIndex();
            ethereumFinalize(index, address(cross), value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(ethereumChainID);
            assertEq(userTokenBalance + (value / exrate), cross.balanceOf(USER));
            assertEq(bridgeTokenBalance - (value / exrate), cross.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossChainID);
            assertEq(userCoinBalance - (value + gas + service), USER.balance);
            assertEq(bridgeCoinBalance + value, address(bridgeCross).balance);
            assertEq(rewardWalletBalance + gas + service, REWARD.balance);
        }
    }

    // erc20 token bridge (ethereum chain -> cross chain)
    function depositToken(bool isRevert, uint amount, uint validatorNum) public {
        vm.selectFork(ethereumChainID);
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        (uint index, bool ok) = ethereumBridge(address(testTokenEthereum), amount, 0, 0);
        if (ok) {
            ethereumIncrementIndex();
            crossFinalize(index, address(testTokenCross), amount, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(ethereumChainID);
            assertEq(userEthereumBalance - amount, testTokenEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance + amount, testTokenEthereum.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossChainID);
            assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
        }
    }

    // erc20 token bridge (cross chain -> ethereum chain)
    function withdrawToken(bool isRevert, uint amount, uint validatorNum) public {
        (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
        assertTrue(value + gas + service <= amount);
        uint total = value + gas + service;
        assertTrue(total <= amount);

        vm.selectFork(ethereumChainID);
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossChainID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);
        uint rewardWalletBalance = testTokenCross.balanceOf(REWARD);

        (uint index, bool ok) = crossBridge(address(testTokenCross), value, gas, service);
        if (ok) {
            crossIncrementIndex();
            ethereumFinalize(index, address(testTokenEthereum), value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(ethereumChainID);
            assertEq(userEthereumBalance + value, testTokenEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance - value, testTokenEthereum.balanceOf(address(bridgeEthereum)));
            vm.selectFork(crossChainID);
            assertEq(userCrossBalance - total, testTokenCross.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + service, testTokenCross.balanceOf(REWARD));
        }
    }
}
