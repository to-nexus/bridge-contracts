// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Test, console} from "forge-std/Test.sol";
// import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
// import {TestToken} from "./TestToken.sol";
// import {BridgeEthereum} from "../src/BridgeEthereum.sol";
// import {BridgeCross} from "../src/BridgeCross.sol";
// import {BridgeStandard} from "../src/BridgeStandard.sol";
// import {BridgeFeeManager, IBridgeFeeManager} from "../src/BridgeFeeManager.sol";
// import {CrossMintableERC20} from "../src/CrossMintableERC20.sol";
import {BridgeStandardTest} from "./BridgeStandard.t.sol";

contract BridgeExceptionTest is Test, BridgeStandardTest {
    function test_deposit_with_insufficient_balance() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        bridgeRevert = true;
        deposit(true, amount + 1000, 5);
    }

    function test_deposit_with_insufficient_validator_signature() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        finalizeRevert = true;
        deposit(true, amount, 1);
    }

    function test_withdraw_with_insufficient_balance() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        deposit(false, amount, 5);

        vm.prank(USER);

        bridgeRevert = true;
        withdraw(true, USER.balance + 1000, 5);
    }

    function test_withdraw_with_insufficient_validator_signature() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        deposit(false, amount, 5);

        vm.prank(USER);

        finalizeRevert = true;
        withdraw(true, amount, 1);
    }

    function test_deposit_with_insufficient_allowance() public {
        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount / 2);

        bridgeRevert = true;
        deposit(true, amount, 5);
    }

    function test_deposit_withdraw_at_token_paused() public {
        // token pause
        vm.prank(OWNER);
        bridgeEthereum.b.pauseToken(IERC20(address(cross)));

        uint amount = 1000;

        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeEthereum.b), amount);

        bridgeRevert = true;
        deposit(true, amount, 5);

        // token unpause
        vm.prank(OWNER);
        bridgeEthereum.b.unpauseToken(IERC20(address(cross)));

        deposit(false, amount, 5);

        // token pause
        vm.prank(OWNER);
        bridgeCross.b.pauseToken(xcross);

        bridgeRevert = true;
        withdraw(true, amount, 5);

        // token unpause
        vm.prank(OWNER);
        bridgeCross.b.unpauseToken(xcross);

        withdraw(false, amount * 10, 5);
    }
}
