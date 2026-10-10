// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Const} from "../src/lib/Const.sol";
import {ForwardLib} from "../src/lib/ForwardLib.sol";
import {IBridgeExecutor} from "../src/interface/IBridgeExecutor.sol";

import {CrossBridgeForwardTest} from "./CrossBridgeForward.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title OverSendingExecutor
 * @notice Stand-in `bridgeExecutor` used ONLY to simulate a compromised or misbehaving
 * executor that forwards MORE native value than the bridge actually handed it for the
 * call into the whitelisted target -- here, the bridge's own `bridgeTokenForwarded`. A
 * spec-compliant `BridgeExecutor` always forwards exactly the `value` it was given, so
 * this is pure defense-in-depth coverage for "what if `bridgeExecutor` were
 * misconfigured or compromised" rather than a reachable production path.
 */
contract OverSendingExecutor {
    function isWhitelistedTarget(address) external pure returns (bool) {
        return true;
    }

    function executeExtraCall(uint, uint, IERC20, address, uint value, bytes calldata extraData)
        external
        payable
        returns (uint consumed, bytes memory returnData)
    {
        address target = address(bytes20(extraData[:20]));
        (bool ok, bytes memory ret) = target.call{value: value + 1}(extraData[20:]);
        require(ok, "target call failed");
        consumed = value;
        returnData = ret;
    }

    receive() external payable {}
}

/**
 * @title CrossBridgeForwardInvalidRecipientTest
 * @notice `bridgeTokenForwarded` rejects `to == address(this)` the same way the ordinary
 * initiate entrypoints do, and independently re-checks the native amount it is funded
 * with rather than trusting whatever `bridgeExecutor` forwards.
 */
contract CrossBridgeForwardInvalidRecipientTest is CrossBridgeForwardTest {
    /// @notice Hop-2's OWN `to` (the ultimate recipient `bridgeTokenForwarded` is asked
    /// to deliver to) set to the bridge itself must revert inside `bridgeTokenForwarded`
    /// -- the executor call then fails, and chain B's existing fallback pays hop-1's
    /// signed recipient directly instead of initiating hop-2 at all.
    function test_bridgeTokenForwarded_toSelf_reverts_fallsBackToDirectPayout() public {
        vm.selectFork(crossForkID);
        uint chainW = 90031;
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(chainW, false, address(NATIVE_TOKEN), address(0xA031));

        uint seedAmount = 10 ether;
        assertTrue(_signAndFinalize(chainW, 1, address(NATIVE_TOKEN), USER, seedAmount, "", 5));

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainW, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;

        // Hop-2's `to` is the bridge's own address -- exactly the case
        // `bridgeTokenForwarded` must reject regardless of what hop-1 signed.
        bytes memory extraData = _forwardExtraData(chainW, address(bridgeCrossV2), value2, fee2, ex2, "");

        uint userBalanceBefore = USER.balance;
        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(
            USER.balance,
            userBalanceBefore + ctxValue,
            "a forward rejected for its own recipient must fall back to hop-1's signed recipient, not strand the value"
        );

        (bool found,,, ForwardLib.ForwardFailureCode code,) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(
            code == ForwardLib.ForwardFailureCode.ExecutorCallReverted,
            "the rejection happens INSIDE bridgeTokenForwarded itself, same as any other initiate-side guard"
        );
    }

    /// @notice Even if `bridgeExecutor` forwards more native value than the bridge gave
    /// it, `bridgeTokenForwarded` still only accepts the exact settled amount --
    /// `_initiateBridge`'s own floor check (relaxed to `>=` for the ordinary
    /// `bridgeToken` refund path) is not relied on here, since this call has no single
    /// caller to refund a surplus to.
    function test_bridgeTokenForwarded_executorOverSendingNative_reverts() public {
        vm.selectFork(crossForkID);
        uint chainW = 90032;
        vm.prank(CrossOWNER);
        bridgeCross.registerToken(chainW, false, address(NATIVE_TOKEN), address(0xA032));

        uint seedAmount = 10 ether;
        assertTrue(_signAndFinalize(chainW, 1, address(NATIVE_TOKEN), USER, seedAmount, "", 5));

        OverSendingExecutor overEx = new OverSendingExecutor();
        vm.deal(address(overEx), 1 ether); // spare balance to over-send with
        vm.prank(CrossOWNER);
        bridgeCross.setBridgeExecutor(IBridgeExecutor(address(overEx)));

        uint value2 = 1 ether;
        (, uint fee2, uint ex2) = bridgeVerifierCross.calculateFee(chainW, IERC20(Const.NATIVE_TOKEN), value2);
        uint ctxValue = value2 + fee2 + ex2;
        bytes memory extraData = _forwardExtraData(chainW, USER, value2, fee2, ex2, "");

        uint userBalanceBefore = USER.balance;
        vm.recordLogs();
        _hop1NativeAndFinalize(ctxValue, extraData);

        vm.selectFork(crossForkID);
        assertEq(
            USER.balance,
            userBalanceBefore + ctxValue,
            "an over-funded forward must fall back to direct payout, not let the surplus ride along"
        );

        (bool found,,, ForwardLib.ForwardFailureCode code,) = _findForwardFailed(vm.getRecordedLogs());
        assertTrue(found, "ForwardFailed must be emitted");
        assertTrue(code == ForwardLib.ForwardFailureCode.ExecutorCallReverted);
    }
}
