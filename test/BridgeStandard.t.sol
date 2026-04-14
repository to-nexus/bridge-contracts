// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {BaseBridge} from "../src/BaseBridge.sol";
import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {TestToken} from "./token/TestToken.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

contract BaseBridgeTest is BridgeTest {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    // ----- Test -----
    function test_depositWithdraw() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        deposit(false, amount, 5);
        withdraw(false, amount, 5);
    }

    function test_depositWithdraw_eth() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        depositETH(false, amount, 5);
        withdrawETH(false, amount, 5);
    }

    function test_depositWithdrawToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);
        depositToken(false, amount, 5);

        vm.selectFork(crossForkID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);
        withdrawToken(false, amount, 5);
    }

    function test_fuzz_depositWithdraw(uint amount) public {
        vm.selectFork(bscForkID);
        vm.assume(amount < (cross.balanceOf(OWNER) / 1e18) && amount != 0);
        amount = amount * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

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

        vm.selectFork(bscForkID);
        uint dealValue = amount - USER.balance;
        vm.deal(USER, dealValue);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        depositETH(false, amount, 5);
        withdrawETH(false, amount, 5);
    }

    function test_fuzz_depositWithdrawToken(uint amount) public {
        vm.selectFork(bscForkID);
        vm.assume(amount < (testTokenBSC.balanceOf(OWNER) / 1e18) && amount != 0);
        amount = amount * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);
        depositToken(false, amount, 5);

        vm.selectFork(crossForkID);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);
        withdrawToken(false, amount, 5);
    }

    function test_pending_finalize() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        depositToken(true, amount, 5); // should finalize revert

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);

        {
            // token stop
            vm.selectFork(bscForkID);
            vm.prank(OWNER);
            TestToken(address(testTokenBSC)).stop();

            (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, service);
            assertTrue(ok);
            crossIncrementIndex();
            vm.selectFork(bscForkID);
            uint before = testTokenBSC.balanceOf(USER);
            bscFinalize(index, address(testTokenBSC), USER, value, 5);
            assertEq(before, testTokenBSC.balanceOf(USER));

            vm.selectFork(bscForkID);
            IBaseBridge.PendingData memory pendingArgs = bridgeBSC.getPendingArguments(CROSS_CHAIN_ID, index);
            assertEq(pendingArgs.delayExpiration, 0);
            assertEq(pendingArgs.args.index, index);
            assertEq(address(pendingArgs.args.toToken), address(testTokenBSC));
            assertEq(pendingArgs.args.to, USER);
            assertEq(pendingArgs.args.value, value);
            assertTrue(Const.FinalizeStatus.Success != pendingArgs.status);

            // token start
            vm.prank(OWNER);
            vm.selectFork(bscForkID);
            TestToken(address(testTokenBSC)).start();

            before = testTokenBSC.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeBSC.releasePending(CROSS_CHAIN_ID, index);

            assertEq(before + value, testTokenBSC.balanceOf(USER));
        }
    }

    function test_pending_finalize_with_revert_transfer() public {
        uint amount = 1000 * 1e18;

        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        depositToken(true, amount, 5); // should finalize revert

        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), amount);

        {
            // token set revert transfer
            vm.selectFork(bscForkID);
            vm.prank(OWNER);
            TestToken(address(testTokenBSC)).setRevertTransfer(true);

            (uint value, uint gas, uint service) = crossCalcFee(IERC20(address(testTokenCross)), amount);
            uint total = value + gas + service;
            assertTrue(total <= amount);
            (uint index, bool ok) = crossBridge(address(testTokenCross), USER, USER, value, gas, service);
            assertTrue(ok);
            crossIncrementIndex();
            vm.selectFork(bscForkID);
            uint before = testTokenBSC.balanceOf(USER);
            bscFinalize(index, address(testTokenBSC), USER, value, 5);
            assertEq(before, testTokenBSC.balanceOf(USER));

            vm.selectFork(bscForkID);
            IBaseBridge.PendingData memory pendingArgs = bridgeBSC.getPendingArguments(CROSS_CHAIN_ID, index);
            assertEq(pendingArgs.delayExpiration, 0);
            assertEq(pendingArgs.args.index, index);
            assertEq(address(pendingArgs.args.toToken), address(testTokenBSC));
            assertEq(pendingArgs.args.to, USER);
            assertEq(pendingArgs.args.value, value);
            assertTrue(Const.FinalizeStatus.Success != pendingArgs.status);

            // token unset revert transfer
            vm.prank(OWNER);
            vm.selectFork(bscForkID);
            TestToken(address(testTokenBSC)).setRevertTransfer(false);

            before = testTokenBSC.balanceOf(USER);

            vm.prank(VALIDATOR1);
            bridgeBSC.releasePending(CROSS_CHAIN_ID, index);

            assertEq(before + value, testTokenBSC.balanceOf(USER));
        }
    }

    function test_permit_deposit() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeBSC)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, NULLDATA);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;
        vm.prank(VALIDATOR1);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);

        crossFinalize(nextIndexCross, address(NATIVE_TOKEN), USER, amount, 5);
        bscIncrementIndex();

        vm.selectFork(bscForkID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeBSC)));
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance + (amount), USER.balance);
        assertEq(bridgeCoinBalance - (amount), address(bridgeCross).balance);
    }

    function test_permit_deposit_before_same_approve() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeBSC)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        // approve before permit
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, NULLDATA);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;
        vm.prank(VALIDATOR1);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);

        crossFinalize(nextIndexCross, address(NATIVE_TOKEN), USER, amount, 5);
        bscIncrementIndex();

        vm.selectFork(bscForkID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeBSC)));
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance + (amount), USER.balance);
        assertEq(bridgeCoinBalance - (amount), address(bridgeCross).balance);
    }

    function test_permit_depositToken() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        assertTrue(testTokenBSC.allowance(USER, address(bridgeBSC)) == 0);

        // initiate
        uint userBSCBalance = testTokenBSC.balanceOf(USER);
        uint bridgeBSCBalance = testTokenBSC.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCrossBalance = testTokenCross.balanceOf(USER);

        uint index;
        IBaseBridge.PermitArguments memory permitArgs;
        {
            // make permit sig
            index = nextIndexCross;

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(testTokenBSC)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, type(uint).max));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(testTokenBSC)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);

            permitArgs =
                IBaseBridge.PermitArguments(IERC20Permit(address(testTokenBSC)), USER, amount, type(uint).max, v, r, s);
        }

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, testTokenBSC, USER, USER, amount, 0, 0, NULLDATA);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;
        vm.prank(VALIDATOR1);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);

        bscIncrementIndex();
        crossFinalize(index, address(testTokenCross), USER, amount, 5);

        vm.selectFork(bscForkID);
        assertEq(userBSCBalance - amount, testTokenBSC.balanceOf(USER));
        assertEq(bridgeBSCBalance + amount, testTokenBSC.balanceOf(address(bridgeBSC)));
        vm.selectFork(crossForkID);
        assertEq(userCrossBalance + amount, testTokenCross.balanceOf(USER));
    }

    /**
     * @notice Test that permitBridgeTokenBatch passes caller-supplied extraData
     * @dev When non-empty extraData is provided, the emitted event should contain the same extraData
     */
    function test_permit_extraData_passed() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        bytes memory extraData = abi.encodePacked(address(0xDEAD), bytes4(0x12345678), bytes("payload"));

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        address expectedRemoteToken = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(cross)).remoteToken;

        vm.expectEmit(true, true, true, true);
        emit BaseBridge.BridgeInitiated(
            CROSS_CHAIN_ID,
            expectedIndex,
            cross,
            IERC20(expectedRemoteToken),
            USER,
            USER,
            amount,
            0,
            0,
            extraData,
            block.timestamp
        );

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, extraData);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;
        vm.prank(VALIDATOR1);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);
    }

    /**
     * @notice Test that permitBridgeTokenBatch passes extraData from args
     * @dev When BridgeTokenArguments contain non-empty extraData, emitted events should contain the same extraData
     */
    function test_permit_batch_extraData_passed() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        bytes memory extraData = abi.encodePacked(address(0xDEAD), bytes4(0x12345678));

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, extraData);

        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;

        uint expectedIndex = bridgeBSC.getNextInitiateIndex(CROSS_CHAIN_ID);
        address expectedRemoteToken = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(cross)).remoteToken;

        vm.expectEmit(true, true, true, true);
        emit BaseBridge.BridgeInitiated(
            CROSS_CHAIN_ID,
            expectedIndex,
            cross,
            IERC20(expectedRemoteToken),
            USER,
            USER,
            amount,
            0,
            0,
            extraData,
            block.timestamp
        );

        vm.prank(VALIDATOR1);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);
    }

    /**
     * @notice Test that permitBridgeTokenBatch reverts when extraData exceeds maxExtraDataLength
     */
    function test_permit_extraData_length_validation() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);

        vm.prank(OWNER);
        bridgeBSC.setMaxExtraDataLength(32);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        bytes memory tooLongExtraData = new bytes(33);

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, tooLongExtraData);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;

        vm.prank(VALIDATOR1);
        vm.expectRevert(BaseBridge.BaseBridgeExtraDataTooLong.selector);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);

        vm.prank(OWNER);
        bridgeBSC.setMaxExtraDataLength(0);
    }

    /**
     * @notice Test that permitBridgeTokenBatch reverts when called by unauthorized address
     */
    function test_permit_revert_unauthorized_caller() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, "");
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;

        vm.prank(USER);
        vm.expectRevert();
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);
    }

    function test_permit_deposit_batch() public {
        // not allow fail
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeBSC)) == 0);

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeBSC));
        vm.selectFork(crossForkID);
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross).balance;

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.PermitArguments memory permitArgs2;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, OWNER, address(bridgeBSC), amount, nonce, deadline));
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
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);
        assertTrue(cross.balanceOf(OWNER) < beforeOwnerBalance); // owner balance should change

        crossFinalize(nextIndexCross, address(NATIVE_TOKEN), USER, amount, 5);
        bscIncrementIndex();

        vm.selectFork(bscForkID);
        assertEq(userTokenBalance - amount, cross.balanceOf(USER));
        assertEq(bridgeTokenBalance + amount * 2, cross.balanceOf(address(bridgeBSC)));
        vm.selectFork(crossForkID);
        assertEq(userCoinBalance + (amount), USER.balance);
        assertEq(bridgeCoinBalance - (amount), address(bridgeCross).balance);
    }

    function test_permit_deposit_batch_with_fail() public {
        // allow fail
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertTrue(cross.allowance(USER, address(bridgeBSC)) == 0);

        IBaseBridge.PermitArguments memory permitArgs;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(USER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
            bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);
            permitArgs = IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        }

        IBaseBridge.PermitArguments memory permitArgs2;
        {
            uint deadline = type(uint).max;
            // make permit sig

            vm.selectFork(bscForkID);
            uint nonce = IERC20Permit(address(cross)).nonces(OWNER);
            bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, OWNER, address(bridgeBSC), amount, nonce, deadline));
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
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);
        assertEq(cross.balanceOf(OWNER), beforeOwnerBalance); // owner balance should not change
    }

    function test_tokenPair_depositedAndMinted() public {
        uint amount = 1000 * 1e18;

        // ===== Original token test (CROSS token) =====
        vm.selectFork(bscForkID);

        // Check initial deposited value
        IBridgeRegistry.TokenPair memory pairBefore = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(cross));

        // Setup for deposit
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        // Perform deposit
        deposit(false, amount, 5);

        // Check if deposited increased
        vm.selectFork(bscForkID);
        IBridgeRegistry.TokenPair memory pairAfter = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(cross));
        assertEq(pairAfter.deposited - pairBefore.deposited, amount, "Deposit amount delta check");

        // ===== Wrapped token test =====
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        testTokenBSC.transfer(USER, amount);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), amount);

        // Check initial minted value
        vm.selectFork(crossForkID);
        IBridgeRegistry.TokenPair memory wrappedPairBefore =
            bridgeCross.getTokenPair(BSC_CHAIN_ID, address(testTokenCross));

        // Deposit original token (which should mint wrapped token)
        depositToken(false, amount, 5);

        // Check if minted increased
        vm.selectFork(crossForkID);
        IBridgeRegistry.TokenPair memory wrappedPairAfter =
            bridgeCross.getTokenPair(BSC_CHAIN_ID, address(testTokenCross));
        assertEq(wrappedPairAfter.minted - wrappedPairBefore.minted, amount, "Mint amount delta check");
    }

    function test_tokenPair_ethDeposited() public {
        uint amount = 1000 * 1e18;

        // ETH deposit/withdraw effect on deposited value test
        vm.selectFork(bscForkID);

        // Check initial deposited value for ETH
        IBridgeRegistry.TokenPair memory pairBefore = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(NATIVE_TOKEN));

        // Setup for deposit
        vm.deal(USER, amount);
        vm.selectFork(crossForkID);
        vm.prank(USER);
        weth.approve(address(bridgeCross), amount);

        // Perform ETH deposit
        depositETH(false, amount, 5);

        // Check if deposited increased
        vm.selectFork(bscForkID);
        IBridgeRegistry.TokenPair memory pairAfter = bridgeBSC.getTokenPair(CROSS_CHAIN_ID, address(NATIVE_TOKEN));
        assertEq(pairAfter.deposited - pairBefore.deposited, amount, "ETH deposit amount delta check");
    }
}
