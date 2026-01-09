// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2} from "../src/token/CrossMintableERC20V2.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/Test.sol";

contract CrossMintableERC20V2Test is Test {
    CrossMintableERC20V2Code public crossMintableERC20Code;
    CrossMintableERC20V2 public crossMintableERC20;

    address public owner;
    address public minter; // minter is the bridge

    uint internal constant OWNER_PK = uint(bytes32("owner"));
    uint internal constant MINTER_PK = uint(bytes32("minter"));

    uint internal constant INITIAL_SUPPLY = 1_000_000_000 ether;

    string internal constant SYMBOL = "TEST";
    uint8 internal constant DECIMALS = 18;

    function setUp() public {
        owner = vm.addr(OWNER_PK);
        minter = vm.addr(MINTER_PK);

        vm.prank(owner);
        crossMintableERC20Code = new CrossMintableERC20V2Code(owner, minter);
        console.log("crossMintableERC20Code", address(crossMintableERC20Code));

        vm.prank(minter);
        crossMintableERC20 =
            CrossMintableERC20V2(crossMintableERC20Code.createCrossMintableERC20(1, address(1), SYMBOL, DECIMALS));
        console.log("crossMintableERC20", address(crossMintableERC20));

        vm.prank(owner);
        crossMintableERC20.grantRole(Const.MINTER_ROLE, minter);
    }

    function test_mint_burn() public {
        // mint
        vm.prank(minter);
        crossMintableERC20.mint(owner, INITIAL_SUPPLY);
        assertEq(crossMintableERC20.balanceOf(owner), INITIAL_SUPPLY);

        // burn
        vm.prank(minter);
        crossMintableERC20.burn(owner, INITIAL_SUPPLY);
        assertEq(crossMintableERC20.balanceOf(owner), 0);
    }
}
