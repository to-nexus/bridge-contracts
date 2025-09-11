// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BSCTest} from "./chain/BSC.t.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {console} from "forge-std/Test.sol";

contract BridgeTest is BSCTest {
    bool enableGasUsedLog;

    function setUp() public virtual override {
        super.setUp();
    }

    // cross token bridge (bsc -> cross chain)
    function deposit(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        vm.selectFork(bscForkID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        bool ok;
        (index, ok) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
            bscIncrementIndex();
            ok = crossFinalize(index, address(NATIVE_TOKEN), USER, amount, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize coin", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(bscForkID);
            assertEq(userTokenBalance - amount, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeBSC)));
            vm.selectFork(crossForkID);
            assertEq(userCoinBalance + amount, USER.balance);
            assertEq(bridgeCoinBalance - amount, address(bridgeCross).balance);
        }
    }

    // xcross bridge (cross chain -> bsc)
    function withdraw(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        (uint value, uint gas, uint ex) = crossCalcFee(NATIVE_TOKEN, amount);
        assertTrue(value + gas + ex <= amount);

        vm.selectFork(bscForkID);
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));
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
            ok = bscFinalize(index, address(cross), USER, value, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(bscForkID);
            assertEq(userTokenBalance + value, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance - value, cross.balanceOf(address(bridgeBSC)));
            vm.selectFork(crossForkID);
            assertEq(userCoinBalance - (value + gas + ex), USER.balance);
            assertEq(bridgeCoinBalance + value, address(bridgeCross).balance);
            assertEq(rewardWalletBalance + gas + ex, REWARD.balance);
        }
    }

    // erc20 token bridge (bsc -> cross chain)
    function depositToken(bool isRevert, uint amount, uint validatorNum) public {
        vm.selectFork(bscForkID);
        uint userBSCBalance = testTokenBSC.balanceOf(USER);
        uint bridgeBSCBalance = testTokenBSC.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        (uint index, bool ok) = bscBridge(address(testTokenBSC), USER, USER, amount, 0, 0);
        if (ok) {
            bscIncrementIndex();
            crossFinalize(index, address(testTokenCross), USER, amount, validatorNum);
            if (ok && enableGasUsedLog) console.log("finalize token (mint)", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(bscForkID);
            assertEq(userBSCBalance - amount, testTokenBSC.balanceOf(USER));
            assertEq(bridgeBSCBalance + amount, testTokenBSC.balanceOf(address(bridgeBSC)));
            vm.selectFork(crossForkID);
            assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
        }
    }

    // erc20 token bridge (cross chain -> bsc)
    function withdrawToken(bool isRevert, uint amount, uint validatorNum) public {
        (uint value, uint gas, uint ex) = crossCalcFee(IERC20(address(testTokenCross)), amount);
        assertTrue(value + gas + ex <= amount);
        uint total = value + gas + ex;
        assertTrue(total <= amount);

        vm.selectFork(bscForkID);
        uint userBSCBalance = testTokenBSC.balanceOf(USER);
        uint bridgeBSCBalance = testTokenBSC.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);
        uint rewardWalletBalance = testTokenCross.balanceOf(REWARD);

        (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, ex);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge token (burn)", uint(vm.lastCallGas().gasTotalUsed));
            crossIncrementIndex();
            bscFinalize(index, address(testTokenBSC), USER, value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(bscForkID);
            assertEq(userBSCBalance + value, testTokenBSC.balanceOf(USER));
            assertEq(bridgeBSCBalance - value, testTokenBSC.balanceOf(address(bridgeBSC)));
            vm.selectFork(crossForkID);
            assertEq(userCrossBalance - total, testTokenCross.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + ex, testTokenCross.balanceOf(REWARD));
        }
    }

    // eth NATIVE_TOKEN bridge (bsc -> cross chain)
    function depositETH(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        vm.selectFork(crossForkID);
        uint userTokenBalance = weth.balanceOf(USER);
        vm.selectFork(bscForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeBSC).balance;

        bool ok;
        (index, ok) = bscBridge(address(NATIVE_TOKEN), USER, USER, amount, 0, 0);
        if (ok) {
            bscIncrementIndex();
            crossFinalize(index, address(weth), USER, amount, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            assertEq(userTokenBalance + amount, weth.balanceOf(USER));
            vm.selectFork(bscForkID);
            assertEq(userCoinBalance - amount, USER.balance);
            assertEq(bridgeCoinBalance + amount, address(bridgeBSC).balance);
        }
    }

    // eth NATIVE_TOKEN bridge (cross chain -> bsc)
    function withdrawETH(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        (uint value, uint gas, uint ex) = crossCalcFee(weth, amount);
        assertTrue(value + gas + ex <= amount);

        vm.selectFork(crossForkID);
        uint userTokenBalance = weth.balanceOf(USER);
        uint rewardWalletBalance = weth.balanceOf(REWARD);
        vm.selectFork(bscForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeBSC).balance;

        bool ok;
        (index, ok) = crossBridge(address(weth), USER, USER, value, gas, ex);
        if (ok) {
            crossIncrementIndex();
            bscFinalize(index, address(NATIVE_TOKEN), USER, value, validatorNum);
        }

        if (!isRevert) {
            vm.selectFork(crossForkID);
            assertEq(userTokenBalance - (value + gas + ex), weth.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + ex, weth.balanceOf(REWARD));
            vm.selectFork(bscForkID);
            assertEq(userCoinBalance + value, USER.balance);
            assertEq(bridgeCoinBalance - value, address(bridgeBSC).balance);
        }
    }

    // burn cross token to dead wallet (bsc -> cross chain)
    function burnCrossToDeadWallet(
        bool isRevert,
        address from,
        address deadWallet,
        uint amount,
        bool alreadyTransferred
    ) public returns (uint index) {
        vm.selectFork(bscForkID);
        uint bscUserBalance = cross.balanceOf(from);
        uint bscBridgeBalance = cross.balanceOf(address(bridgeBSC));
        uint bscDeadWalletBalance = cross.balanceOf(deadWallet);
        vm.selectFork(crossForkID);
        uint crossDeadWalletBalance = deadWallet.balance;
        uint crossBridgeBalance = address(bridgeCross).balance;

        bool ok;
        (index, ok) = bscBurnCrossToDeadWallet(from, deadWallet, amount, alreadyTransferred);
        if (ok) {
            if (enableGasUsedLog) console.log("bridge token (transfer)", uint(vm.lastCallGas().gasTotalUsed));
            bscIncrementIndex();
            ok = crossFinalize(index, address(NATIVE_TOKEN), deadWallet, amount, 5);
            if (ok && enableGasUsedLog) console.log("finalize coin", uint(vm.lastCallGas().gasTotalUsed));
        }

        if (!isRevert) {
            vm.selectFork(bscForkID);
            if (alreadyTransferred) {
                assertEq(bscUserBalance, cross.balanceOf(from));
            } else {
                assertEq(bscUserBalance - amount, cross.balanceOf(from));
                assertEq(bscDeadWalletBalance + amount, cross.balanceOf(deadWallet));
            }
            assertEq(bscBridgeBalance, cross.balanceOf(address(bridgeBSC)));
            vm.selectFork(crossForkID);
            assertEq(crossDeadWalletBalance + amount, deadWallet.balance);
            assertEq(crossBridgeBalance - amount, address(bridgeCross).balance);
        }
    }
}
