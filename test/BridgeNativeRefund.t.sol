// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {Const} from "../src/lib/Const.sol";

import {BridgeTest} from "./Bridge.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

/// @notice Calls `bridgeToken` on behalf of whoever funds it, and rejects any native
/// value sent back to it - used to prove a failed refund reverts the whole call rather
/// than silently dropping the surplus.
contract RevertingRefundCaller {
    function callBridgeToken(
        BaseBridge bridge,
        uint toChainID,
        IERC20 token,
        address to,
        uint value,
        uint networkFee,
        uint exFee
    ) external payable {
        bridge.bridgeToken{value: msg.value}(toChainID, token, to, value, networkFee, exFee, "");
    }

    receive() external payable {
        revert("RevertingRefundCaller: no thanks");
    }
}

/**
 * @title BridgeNativeRefundTest
 * @notice `bridgeToken`'s native branch treats `msg.value` as an upper bound: it must
 * cover `value + networkFee + exFee`, and any surplus comes back to the caller in the
 * same transaction. The other two payable entrypoints keep requiring exact native value
 * (or none at all), since neither has a single caller a surplus could be returned to.
 */
contract BridgeNativeRefundTest is BridgeTest {
    function setUp() public override {
        super.setUp();
    }

    /// @notice `msg.value` exactly equal to `value + networkFee + exFee` still succeeds,
    /// with no refund call made at all (the bridge's balance accounts for every wei).
    function test_bridgeToken_native_exactMsgValue_noRefund() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint ex) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint total = value + gas + ex;

        vm.deal(USER, total);
        uint userBalBefore = USER.balance;
        uint bridgeBalBefore = address(bridgeBSC).balance;

        vm.prank(USER);
        assertTrue(bridgeBSC.bridgeToken{value: total}(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, value, gas, ex, NULLDATA));

        assertEq(USER.balance, userBalBefore - total, "no refund: the exact amount is spent");
        assertEq(address(bridgeBSC).balance, bridgeBalBefore + total);
    }

    /// @notice `msg.value` above the required amount refunds exactly the difference,
    /// and the bridge's own balance only ever reflects the required amount, not the
    /// surplus.
    function test_bridgeToken_native_excessMsgValue_refundsExactDifference() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint ex) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint required = value + gas + ex;
        uint surplus = 0.37 ether;

        vm.deal(USER, required + surplus);
        uint userBalBefore = USER.balance;
        uint bridgeBalBefore = address(bridgeBSC).balance;

        vm.prank(USER);
        assertTrue(
            bridgeBSC.bridgeToken{value: required + surplus}(
                CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, value, gas, ex, NULLDATA
            )
        );

        assertEq(USER.balance, userBalBefore - required, "only the required amount should be net-spent");
        assertEq(address(bridgeBSC).balance, bridgeBalBefore + required, "the bridge must not retain the surplus");
    }

    /// @notice `msg.value` short of `value + networkFee + exFee` still reverts - the
    /// floor is a floor, not a suggestion.
    function test_bridgeToken_native_shortMsgValue_reverts() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint ex) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint required = value + gas + ex;

        vm.deal(USER, required);
        vm.prank(USER);
        vm.expectRevert(
            abi.encodeWithSelector(BaseBridge.BaseBridgeInvalidValue.selector, required, required - 1)
        );
        bridgeBSC.bridgeToken{value: required - 1}(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, value, gas, ex, NULLDATA);
    }

    /// @notice A fee rate drop between when the caller quoted `networkFee`/`exFee` and
    /// when the transaction actually executes now succeeds (instead of reverting the
    /// whole bridge over stale fee numbers), refunding the caller for the fee the quote
    /// over-estimated.
    function test_bridgeToken_native_feeDropBetweenQuoteAndExecution_succeedsWithRefund() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);

        // This fixture's default exchange-fee rate is 0 - set a real nonzero rate first
        // so there is an actual fee to drop.
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(IERC20(Const.NATIVE_TOKEN), 500); // 5% of denominator()

        (uint value, uint quotedGas, uint quotedEx) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        require(quotedEx > 0, "test requires a nonzero quoted exchange fee to demonstrate a drop");
        uint quotedTotal = value + quotedGas + quotedEx;

        // The exchange-fee rate drops to zero (an editor action) before execution -
        // `type(uint).max` is the verifier's documented fee-exemption sentinel; a plain
        // `0` would instead fall back to the default rate.
        vm.prank(OWNER);
        bridgeVerifierBSC.setExFeeRate(IERC20(Const.NATIVE_TOKEN), type(uint).max);

        (, uint actualGas, uint actualEx) = bridgeVerifierBSC.calculateFee(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), value);
        assertLt(actualEx, quotedEx, "the drop must actually lower the recomputed exchange fee");
        uint actualTotal = value + actualGas + actualEx;

        vm.deal(USER, quotedTotal);
        uint userBalBefore = USER.balance;

        // The caller still sends the STALE (higher) quote - the fee locals get
        // overwritten by `_checkInitiateAmount` to the fresh (lower) ones before the
        // refund is computed.
        vm.prank(USER);
        assertTrue(
            bridgeBSC.bridgeToken{value: quotedTotal}(
                CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, value, quotedGas, quotedEx, NULLDATA
            )
        );

        assertEq(USER.balance, userBalBefore - actualTotal, "only the freshly-recomputed (lower) cost is net-spent");
    }

    /// @notice The refund goes to `_msgSender()` (the caller), never to `to` (the
    /// destination-chain recipient), even when the two differ.
    function test_bridgeToken_refund_goesToCaller_notToRecipient() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint ex) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint required = value + gas + ex;
        uint surplus = 0.1 ether;

        address recipient = makeAddr("different_recipient");
        vm.deal(USER, required + surplus);
        uint userBalBefore = USER.balance;
        uint recipientBalBefore = recipient.balance;

        vm.prank(USER);
        assertTrue(
            bridgeBSC.bridgeToken{value: required + surplus}(
                CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), recipient, value, gas, ex, NULLDATA
            )
        );

        assertEq(USER.balance, userBalBefore - required, "the caller must receive the surplus");
        assertEq(recipient.balance, recipientBalBefore, "the destination-chain recipient must not receive any refund");
    }

    /// @notice A caller whose `receive()` rejects native value reverts the ENTIRE
    /// `bridgeToken` call when a refund is owed - the deposit/fee accounting already
    /// performed is rolled back with it, rather than silently dropping the surplus.
    function test_bridgeToken_refundRecipientRejectsNative_revertsWholeCall() public {
        uint amount = 10 ether;
        vm.selectFork(bscForkID);
        (uint value, uint gas, uint ex) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        uint required = value + gas + ex;

        RevertingRefundCaller caller = new RevertingRefundCaller();
        vm.deal(address(caller), required + 1 ether);

        vm.expectRevert(BaseBridge.BaseBridgeFailedCall.selector);
        caller.callBridgeToken{value: required + 1 ether}(
            bridgeBSC, CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), USER, value, gas, ex
        );
    }

    /// @notice The ERC20 branch is unaffected by the native refund change: it still
    /// requires `msg.value == 0` exactly as before.
    function test_bridgeToken_erc20_stillRequiresZeroMsgValue() public {
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 10 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);

        vm.deal(USER, 1 ether);
        vm.prank(USER);
        vm.expectRevert(abi.encodeWithSelector(BaseBridge.BaseBridgeInvalidValue.selector, 0, 1 ether));
        bridgeBSC.bridgeToken{value: 1 ether}(CROSS_CHAIN_ID, cross, USER, 1 ether, 0, 0, NULLDATA);
    }

    /// @notice `permitBridgeTokenBatch` rejects ANY non-zero `msg.value`, single item or
    /// not - it is permit-funded (ERC20) throughout and has no single caller to refund a
    /// surplus to.
    function test_permitBridgeTokenBatch_nonZeroMsgValue_reverts() public {
        vm.selectFork(bscForkID);
        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments({
            toChainID: CROSS_CHAIN_ID,
            fromToken: cross,
            from: USER,
            to: USER,
            value: 1 ether,
            networkFee: 0,
            exFee: 0,
            extraData: NULLDATA
        });
        IBaseBridge.PermitArguments[] memory permitArgs = new IBaseBridge.PermitArguments[](1);
        permitArgs[0] = IBaseBridge.PermitArguments({
            token: IERC20Permit(address(cross)),
            account: USER,
            value: 1 ether,
            deadline: type(uint).max,
            v: 0,
            r: bytes32(0),
            s: bytes32(0)
        });

        vm.deal(VALIDATOR1, 1 ether);
        vm.prank(VALIDATOR1); // holds INITIATOR_ROLE per the BSC fixture
        vm.expectRevert(abi.encodeWithSelector(BaseBridge.BaseBridgeInvalidValue.selector, 0, 1 ether));
        bridgeBSC.permitBridgeTokenBatch{value: 1 ether}(args, permitArgs);
    }

    /// @notice The same guard blocks a MULTI-item batch exactly the same way - closing
    /// the shape where, before this change, each native item could have independently
    /// floor-checked against the SAME `msg.value`, letting a batch re-use one funding
    /// source across several items.
    function test_permitBridgeTokenBatch_multiItem_cannotCarryNative() public {
        vm.selectFork(bscForkID);
        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](2);
        IBaseBridge.PermitArguments[] memory permitArgs = new IBaseBridge.PermitArguments[](2);
        for (uint i = 0; i < 2; i++) {
            args[i] = IBaseBridge.BridgeTokenArguments({
                toChainID: CROSS_CHAIN_ID,
                fromToken: IERC20(Const.NATIVE_TOKEN),
                from: USER,
                to: USER,
                value: 1 ether,
                networkFee: 0,
                exFee: 0,
                extraData: NULLDATA
            });
            permitArgs[i] = IBaseBridge.PermitArguments({
                token: IERC20Permit(Const.NATIVE_TOKEN),
                account: USER,
                value: 1 ether,
                deadline: type(uint).max,
                v: 0,
                r: bytes32(0),
                s: bytes32(0)
            });
        }

        vm.deal(VALIDATOR1, 1 ether);
        vm.prank(VALIDATOR1);
        vm.expectRevert(abi.encodeWithSelector(BaseBridge.BaseBridgeInvalidValue.selector, 0, 1 ether));
        bridgeBSC.permitBridgeTokenBatch{value: 1 ether}(args, permitArgs);
    }
}
