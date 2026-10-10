// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../src/BaseBridge.sol";
import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {Const} from "../src/lib/Const.sol";

import {BridgeExecutorTest, MockTargetContract} from "./BridgeExecutor.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

/**
 * @title BridgeInvalidRecipientTest
 * @notice No path can settle bridged value to the bridge contract itself. Initiate-side
 * entrypoints reject `to == address(this)` up front; the settlement chokepoint inside
 * `_finalizeBridge` catches it regardless of which finalize-time path (ordinary
 * finalize, `releasePending`, `manualReleasePendingWithRecipient`'s replacement
 * recipient) or token branch (native, origin transfer, wrapped mint, executor) reaches
 * it.
 * @dev Builds on `BridgeExecutorTest` for the already-wired `bridgeExecutorCross` /
 * `mockTargetCross` fixture the executor-branch test needs.
 */
contract BridgeInvalidRecipientTest is BridgeExecutorTest {
    function setUp() public override {
        super.setUp();

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, 1_000 ether);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), type(uint).max);

        vm.prank(OWNER);
        testTokenBSC.transfer(USER, 1_000 ether);
        vm.prank(USER);
        testTokenBSC.approve(address(bridgeBSC), type(uint).max);
    }

    // ----------------------------------------------------------------
    // Initiate-side guards
    // ----------------------------------------------------------------

    /// @notice `bridgeToken` rejects `to == address(this)` for a native transfer, before
    /// any fee or balance check even runs.
    function test_bridgeToken_native_toSelf_reverts() public {
        vm.selectFork(bscForkID);
        vm.expectRevert(BaseBridge.BaseBridgeInvalidRecipient.selector);
        bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(Const.NATIVE_TOKEN), address(bridgeBSC), 1, 0, 0, NULLDATA);
    }

    /// @notice Same guard, ERC20 branch.
    function test_bridgeToken_erc20_toSelf_reverts() public {
        vm.selectFork(bscForkID);
        vm.expectRevert(BaseBridge.BaseBridgeInvalidRecipient.selector);
        bridgeBSC.bridgeToken(CROSS_CHAIN_ID, cross, address(bridgeBSC), 1, 0, 0, NULLDATA);
    }

    /// @notice `permitBridgeTokenBatch`'s per-item guard fires even with `to` matching
    /// `permitArgs.account` (so the existing mismatch check alone would not have caught
    /// it) and before `safePermit` is ever reached - no real signature is needed to
    /// observe the revert.
    function test_permitBridgeTokenBatch_toSelf_reverts() public {
        vm.selectFork(bscForkID);
        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments({
            toChainID: CROSS_CHAIN_ID,
            fromToken: cross,
            from: address(bridgeBSC),
            to: address(bridgeBSC),
            value: 1,
            networkFee: 0,
            exFee: 0,
            extraData: NULLDATA
        });
        IBaseBridge.PermitArguments[] memory permitArgs = new IBaseBridge.PermitArguments[](1);
        permitArgs[0] = IBaseBridge.PermitArguments({
            token: IERC20Permit(address(cross)),
            account: address(bridgeBSC),
            value: 1,
            deadline: type(uint).max,
            v: 0,
            r: bytes32(0),
            s: bytes32(0)
        });

        vm.prank(VALIDATOR1); // holds INITIATOR_ROLE per the BSC fixture
        vm.expectRevert(BaseBridge.BaseBridgeInvalidRecipient.selector);
        bridgeBSC.permitBridgeTokenBatch(args, permitArgs);
    }

    // ----------------------------------------------------------------
    // Settlement chokepoint (`_finalizeBridge`) - every branch that reaches it
    // ----------------------------------------------------------------

    /// @notice Native finalize settling to the bridge itself parks as
    /// `InvalidRecipient` instead of crediting the bridge's own balance, AND the
    /// finalize index still advances - the very next index finalizes normally.
    function test_finalize_native_toSelf_parksAndIndexAdvances() public {
        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), address(bridgeCross), amount, threshold);
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InvalidRecipient,
            "settling native value to the bridge itself must park, not succeed"
        );

        vm.selectFork(bscForkID);
        (uint index2,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();
        assertEq(index2, index + 1, "the finalize index must not be stuck behind the parked item");

        vm.selectFork(crossForkID);
        uint balBefore = USER.balance;
        assertTrue(crossFinalize(index2, address(NATIVE_TOKEN), USER, amount, threshold));
        assertEq(USER.balance, balBefore + amount, "the next index must finalize normally");
    }

    /// @notice Origin-side ERC20 finalize (plain `transfer`, not mint) settling to the
    /// bridge itself also parks as `InvalidRecipient`.
    function test_finalize_originERC20Transfer_toSelf_parks() public {
        uint amount = 50 ether;

        // Establish USER's wrapped TT balance on CROSS first (ordinary deposit), then
        // bridge it back so the finalize lands on BSC against `testTokenBSC` - the
        // isOrigin=true, plain-transfer pair.
        depositToken(false, amount, threshold);

        vm.selectFork(crossForkID);
        (uint value, uint gas, uint ex) = crossCalcFee(IERC20(address(testTokenCross)), amount);
        vm.prank(USER);
        testTokenCross.approve(address(bridgeCross), type(uint).max);
        (uint index,) = crossBridge(address(testTokenCross), USER, USER, value, gas, ex);
        crossIncrementIndex();

        vm.selectFork(bscForkID);
        bscFinalize(index, address(testTokenBSC), address(bridgeBSC), value, threshold);
        assertTrue(
            bridgeBSC.getPendingArguments(CROSS_CHAIN_ID, index).status == Const.FinalizeStatus.InvalidRecipient,
            "origin-side transfer settling to the bridge itself must park"
        );
    }

    /// @notice Wrapped-side ERC20 finalize (mint, not transfer) settling to the bridge
    /// itself also parks as `InvalidRecipient`.
    function test_finalize_wrappedERC20Mint_toSelf_parks() public {
        uint amount = 50 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(testTokenBSC), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(testTokenCross), address(bridgeCross), amount, threshold);
        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InvalidRecipient,
            "wrapped-side mint settling to the bridge itself must park"
        );
    }

    /// @notice The guard sits ahead of extraData parsing entirely: even when extraData
    /// resolves to a genuinely whitelisted executor target, `to == address(this)` still
    /// parks before the executor is ever invoked - the target's balance never moves.
    function test_finalize_executorBranch_toSelf_parksWithoutInvokingExecutor() public {
        uint amount = 1_000 ether;

        vm.selectFork(bscForkID);
        vm.deal(USER, amount);
        bytes memory calldata_ = abi.encodeWithSelector(
            MockTargetContract.handleBridgeCallback.selector, address(1), USER, amount, bytes("test data")
        );
        bytes memory extraData = abi.encodePacked(address(mockTargetCross), calldata_);

        (uint value, uint gas, uint service) = bscCalcFee(IERC20(Const.NATIVE_TOKEN), amount);
        (uint index,) = bscBridge(Const.NATIVE_TOKEN, USER, USER, value, gas, service, extraData);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        uint targetBalBefore = address(mockTargetCross).balance;
        crossFinalize(index, address(NATIVE_TOKEN), address(bridgeCross), value, threshold, extraData);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InvalidRecipient,
            "the recipient guard must fire ahead of the executor dispatch"
        );
        assertEq(address(mockTargetCross).balance, targetBalBefore, "the executor must never have been invoked");
    }

    // ----------------------------------------------------------------
    // Recovery paths for a parked `InvalidRecipient` record
    // ----------------------------------------------------------------

    /// @notice `releasePending` on an `InvalidRecipient` record reverts (the record was
    /// parked for who it was going to, not for its amount, and releasing re-plays the
    /// SAME bad recipient) - but the pending record itself survives the failed attempt,
    /// ready for `manualReleasePendingWithRecipient`.
    function test_releasePending_onInvalidRecipient_reverts_recordSurvives() public {
        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), address(bridgeCross), amount, threshold);

        vm.expectRevert(
            abi.encodeWithSelector(BaseBridge.BaseBridgeFailedRelease.selector, Const.FinalizeStatus.InvalidRecipient)
        );
        bridgeCross.releasePending(BSC_CHAIN_ID, index);

        assertTrue(
            bridgeCross.getPendingArguments(BSC_CHAIN_ID, index).status == Const.FinalizeStatus.InvalidRecipient,
            "a failed release must not remove the pending record"
        );
    }

    /// @notice `manualReleasePendingWithRecipient` with the bridge itself as the
    /// replacement recipient reverts exactly the same way; with an ordinary EOA
    /// replacement it succeeds and clears the record.
    function test_manualReleasePendingWithRecipient_toSelf_reverts_toEOA_succeeds() public {
        uint amount = 5 ether;

        vm.selectFork(bscForkID);
        (uint index,) = bscBridge(address(cross), USER, USER, amount, 0, 0);
        bscIncrementIndex();

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), address(bridgeCross), amount, threshold);

        vm.prank(CrossOWNER); // holds VERIFIER_ROLE per the CROSS fixture
        vm.expectRevert(
            abi.encodeWithSelector(BaseBridge.BaseBridgeFailedRelease.selector, Const.FinalizeStatus.InvalidRecipient)
        );
        bridgeCross.manualReleasePendingWithRecipient(BSC_CHAIN_ID, index, address(bridgeCross));

        address rescueRecipient = makeAddr("rescue_recipient");
        uint balBefore = rescueRecipient.balance;
        vm.prank(CrossOWNER);
        bridgeCross.manualReleasePendingWithRecipient(BSC_CHAIN_ID, index, rescueRecipient);
        assertEq(rescueRecipient.balance, balBefore + amount, "an EOA replacement recipient must still succeed");
    }
}
