// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {CrossMintableERC20V2} from "../../src/token/CrossMintableERC20V2.sol";
import {Script, console} from "forge-std/Script.sol";

contract CrossMintableERC20V2Script is Script {
    function setUp() public {}

    function deployCrossMintableERC20V2(
        address initialOwner,
        address initialMinter,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) public {
        vm.broadcast();
        CrossMintableERC20V2 token = new CrossMintableERC20V2(initialOwner, initialMinter, name, symbol, decimals);
        console.log("CrossMintableERC20V2 will be deployed to:", address(token));
    }
}
