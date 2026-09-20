// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IBaseBridge} from "../src/interface/IBaseBridge.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

/**
 * @title BridgeSafePermitTest
 * @notice A front-run permit signature must not block the bridge's own gasless deposit
 * flow, and `safePermit` must reject any caller other than a bridge holding
 * `BRIDGE_ROLE`.
 */
contract BridgeSafePermitTest is BridgeTest {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    /**
     * @notice T14: A third party who front-runs the user's permit signature in the
     * mempool (submitting it directly to the token before the bridge does) burns the
     * nonce but also creates the allowance the bridge needs - so the bridge's own
     * `permitBridgeTokenBatch` call must still succeed via `safePermit`'s early return,
     * not revert on the now-consumed signature.
     */
    function test_safePermit_frontRunDoesNotBlockBridge() public {
        uint amount = 1000 * 1e18;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        assertEq(cross.allowance(USER, address(bridgeBSC)), 0);

        uint deadline = type(uint).max;
        uint nonce = IERC20Permit(address(cross)).nonces(USER);
        bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeBSC), amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);

        // Attacker front-runs: submits the exact same signature directly to the token,
        // burning USER's nonce and creating the allowance the bridge would have created
        // itself.
        address attacker = makeAddr("permit_frontrunner");
        vm.prank(attacker);
        IERC20Permit(address(cross)).permit(USER, address(bridgeBSC), amount, deadline, v, r, s);
        assertEq(cross.allowance(USER, address(bridgeBSC)), amount, "front-run must have created the allowance");

        // The bridge's own permit-based deposit, using the SAME (now-consumed)
        // signature, must still succeed - it never needs to call permit() again because
        // the allowance is already sufficient.
        IBaseBridge.PermitArguments memory permitArgs =
            IBaseBridge.PermitArguments(IERC20Permit(address(cross)), USER, amount, deadline, v, r, s);
        IBaseBridge.BridgeTokenArguments[] memory args = new IBaseBridge.BridgeTokenArguments[](1);
        args[0] = IBaseBridge.BridgeTokenArguments(CROSS_CHAIN_ID, cross, USER, USER, amount, 0, 0, NULLDATA);
        IBaseBridge.PermitArguments[] memory permitArgsArray = new IBaseBridge.PermitArguments[](1);
        permitArgsArray[0] = permitArgs;

        uint userBalanceBefore = cross.balanceOf(USER);
        vm.prank(VALIDATOR1); // holds INITIATOR_ROLE per the BSC fixture
        bridgeBSC.permitBridgeTokenBatch(args, permitArgsArray);

        assertEq(cross.balanceOf(USER), userBalanceBefore - amount, "deposit must have gone through");
    }

    /**
     * @notice T15: `safePermit` is restricted to `BRIDGE_ROLE`. A direct call from
     * anyone else (the bridge itself is the only holder, granted at the verifier's
     * construction) must revert before even inspecting the signature.
     */
    function test_safePermit_directCall_requiresBridgeRole() public {
        vm.selectFork(bscForkID);
        vm.expectRevert();
        bridgeVerifierBSC.safePermit(cross, USER, address(bridgeBSC), 1 ether, type(uint).max, 0, bytes32(0), bytes32(0));
    }
}
