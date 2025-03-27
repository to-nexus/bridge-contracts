// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IBaseBridge} from "../src/interface/IBaseBridge.sol";

import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract BaseBridgeTest is BridgeTest {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    // ----- Test -----
    function test_depositWithdraw() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);
        withdraw(false, amount, 5);
    }

    function test_depositWithdraw_eth() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.deal(USER, amount);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        depositETH(false, amount, 5);
        withdrawETH(false, amount, 5);
    }

    function test_depositWithdrawToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenEthereum.approve(address(bridgeEthereum), amount);
        depositToken(false, amount, 5);

        vm.selectFork(crossForkID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);
        withdrawToken(false, amount, 5);
    }

    function test_fuzz_depositWithdraw(uint amount) public {
        vm.selectFork(ethereumForkID);
        vm.assume(amount < (cross.balanceOf(OWNER) / 1e18) && amount != 0);
        amount = amount * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum), amount);

        deposit(false, amount, 5);
        withdraw(false, amount, 5);
    }

    function test_fuzz_depositWithdraw_eth(uint amount) public {
        (uint minimum, uint gasFee, uint exFeeRate) = crossGetTokenFee(weth);
        uint denom = bridgeVerifierCross.denominator();
        minimum = minimum + gasFee + (minimum * exFeeRate / denom);
        if (minimum < 10) minimum = 1000;

        vm.assume(amount > minimum && amount != 0);
        vm.assume(amount < 1e24);

        vm.selectFork(ethereumForkID);
        uint dealValue = amount - USER.balance;
        vm.deal(USER, dealValue);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        depositETH(false, amount, 5);
        withdrawETH(false, amount, 5);
    }

    function test_fuzz_depositWithdrawToken(uint amount) public {
        vm.selectFork(ethereumForkID);
        vm.assume(amount < (testTokenEthereum.balanceOf(OWNER) / 1e18) && amount != 0);
        amount = amount * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenEthereum.approve(address(bridgeEthereum), amount);
        depositToken(false, amount, 5);

        vm.selectFork(crossForkID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);
        withdrawToken(false, amount, 5);
    }

    function test_pending_finalize() public {
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
            vm.selectFork(ethereumForkID);
            vm.prank(OWNER);
            TestToken(address(testTokenEthereum)).stop();

            (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, service);
            assertTrue(ok);
            crossIncrementIndex();
            vm.selectFork(ethereumForkID);
            uint before = testTokenEthereum.balanceOf(USER);
            ethereumFinalize(index, address(testTokenEthereum), USER, value, 5);
            assertEq(before, testTokenEthereum.balanceOf(USER));

            vm.selectFork(ethereumForkID);
            IBaseBridge.PendingData memory pendingArgs = bridgeEthereum.getPendingArguments(CROSS_CHAIN_ID, index);
            assertEq(pendingArgs.delayExpiration, 0);
            assertEq(pendingArgs.args.index, index);
            assertEq(address(pendingArgs.args.toToken), address(testTokenEthereum));
            assertEq(pendingArgs.args.to, USER);
            assertEq(pendingArgs.args.value, value);
            assertTrue(Const.FinalizeStatus.Success != pendingArgs.status);

            // token start
            vm.prank(OWNER);
            vm.selectFork(ethereumForkID);
            TestToken(address(testTokenEthereum)).start();

            before = testTokenEthereum.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeEthereum.releasePending(CROSS_CHAIN_ID, index);

            assertEq(before + value, testTokenEthereum.balanceOf(USER));
        }
    }

    function test_pending_finalize_with_revert_transfer() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        vm.prank(USER);
        testTokenEthereum.approve(address(bridgeEthereum), amount);

        depositToken(true, amount, 5); // should finalize revert

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);

        {
            // token set revert transfer
            vm.selectFork(ethereumForkID);
            vm.prank(OWNER);
            TestToken(address(testTokenEthereum)).setRevertTransfer(true);

            (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, service);
            assertTrue(ok);
            crossIncrementIndex();
            vm.selectFork(ethereumForkID);
            uint before = testTokenEthereum.balanceOf(USER);
            ethereumFinalize(index, address(testTokenEthereum), USER, value, 5);
            assertEq(before, testTokenEthereum.balanceOf(USER));

            vm.selectFork(ethereumForkID);
            IBaseBridge.PendingData memory pendingArgs = bridgeEthereum.getPendingArguments(CROSS_CHAIN_ID, index);
            assertEq(pendingArgs.delayExpiration, 0);
            assertEq(pendingArgs.args.index, index);
            assertEq(address(pendingArgs.args.toToken), address(testTokenEthereum));
            assertEq(pendingArgs.args.to, USER);
            assertEq(pendingArgs.args.value, value);
            assertTrue(Const.FinalizeStatus.Success != pendingArgs.status);

            // token unset revert transfer
            vm.prank(OWNER);
            vm.selectFork(ethereumForkID);
            TestToken(address(testTokenEthereum)).setRevertTransfer(false);

            before = testTokenEthereum.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeEthereum.releasePending(CROSS_CHAIN_ID, index);

            assertEq(before + value, testTokenEthereum.balanceOf(USER));
        }
    }

    function test_permit_deposit() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeEthereum)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        vm.prank(VALIDATOR1);
        assertTrue(bridgeEthereum.permitBridgeToken(CROSS_CHAIN_ID, cross, USER, amount, 0, 0, NULLDATA, permitArgs));

        crossFinalize(nextIndexCross, address(NATIVE_TOKEN), USER, amount, 5);
        ethereumIncrementIndex();

        vm.selectFork(ethereumForkID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum)));
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance + (amount), USER.balance);
        assertEq(bridgeCoinBalance - (amount), address(bridgeCross).balance);
    }

    function test_permit_depositToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        testTokenEthereum.transfer(USER, amount);
        assertTrue(testTokenEthereum.allowance(USER, address(bridgeEthereum)) == 0);

        // initiate
        uint userEthereumBalance = testTokenEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenEthereum.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        uint index;
        IBaseBridge.PermitArguments memory permitArgs;
        {
            // make permit sig
            index = nextIndexCross;

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(testTokenEthereum)).nonces(USER);
            bytes32 h =
                keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, type(uint).max));
            bytes32 hash =
                MessageHashUtils.toTypedDataHash(IERC20Permit(address(testTokenEthereum)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);

            permitArgs = IBaseBridge.PermitArguments(
                IERC20Permit(address(testTokenEthereum)), USER, amount, type(uint).max, v, r, s
            );
        }

        vm.prank(VALIDATOR1);
        bool ok = bridgeEthereum.permitBridgeToken(
            CROSS_CHAIN_ID, testTokenEthereum, USER, amount, 0, 0, NULLDATA, permitArgs
        );
        assertTrue(ok);

        ethereumIncrementIndex();
        crossFinalize(index, address(testTokenCross), USER, amount, 5);

        vm.selectFork(ethereumForkID);
        assertEq(userEthereumBalance - amount, testTokenEthereum.balanceOf(USER));
        assertEq(bridgeEthereumBalance + amount, testTokenEthereum.balanceOf(address(bridgeEthereum)));
        vm.selectFork(crossForkID);
        assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
    }

    function test_permit_deposit_batch() public {
        // not allow fail
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeEthereum)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.PermitArguments memory permitArgs2;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, OWNER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(OWNER_PK, hash);
            permitArgs2 = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), OWNER, amount, deadline, v, r, s);
        }

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](2);
        {
            args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, NULLDATA);
        }
        {
            args[1] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, OWNER, OWNER, amount, 0, 0, NULLDATA);
        }
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](2);
        {
            permitArgsArray[0] = permitArgs;
        }
        {
            permitArgsArray[1] = permitArgs2;
        }

        uint beforeOwnerBalance = cross.balanceOf(OWNER);
        vm.prank(VALIDATOR1);
        bridgeEthereum.permitBridgeTokenBatch(args, permitArgsArray);
        assertTrue(cross.balanceOf(OWNER) < beforeOwnerBalance); // owner balance should change

        crossFinalize(nextIndexCross, address(NATIVE_TOKEN), USER, amount, 5);
        ethereumIncrementIndex();

        vm.selectFork(ethereumForkID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount * 2, cross.balanceOf(address(bridgeEthereum)));
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance + (amount), USER.balance);
        assertEq(bridgeCoinBalance - (amount), address(bridgeCross).balance);
    }

    function test_permit_deposit_batch_with_fail() public {
        // allow fail
        uint amount = 1000 * 1e18;

        vm.selectFork(ethereumForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeEthereum)) == 0);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.PermitArguments memory permitArgs2;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(ethereumForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, OWNER, address(bridgeEthereum), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(OWNER_PK, hash);
            permitArgs2 =
                IBaseBridge.PermitArguments(IERC20Permit(address(cross)), OWNER, amount, deadline, v + 1, r, s); // invalid v
        }

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](2);
        {
            args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, NULLDATA);
        }
        {
            args[1] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, OWNER, OWNER, amount, 0, 0, NULLDATA);
        }
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](2);
        {
            permitArgsArray[0] = permitArgs;
        }
        {
            permitArgsArray[1] = permitArgs2;
        }

        uint beforeOwnerBalance = cross.balanceOf(OWNER);
        vm.prank(VALIDATOR1);
        vm.expectRevert();
        bridgeEthereum.permitBridgeTokenBatch(args, permitArgsArray);
        assertEq(cross.balanceOf(OWNER), beforeOwnerBalance); // owner balance should not change
    }
}
