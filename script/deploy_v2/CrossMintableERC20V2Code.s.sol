// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../../src/BaseBridge.sol";
import {Const} from "../../src/lib/Const.sol";
import {CrossMintableERC20V2Code, ICrossMintableERC20Code} from "../../src/token/CrossMintableERC20V2Code.sol";
import {Script, console} from "forge-std/Script.sol";

contract CrossMintableERC20V2CodeScript is Script {
    function setUp() public {}

    function deployCrossMintableERC20V2Code(BaseBridge bridge, address admin) public {
        vm.startBroadcast();
        CrossMintableERC20V2Code code = new CrossMintableERC20V2Code(admin, address(bridge));
        console.log("CrossMintableERC20V2Code will be deployed to:", address(code));

        bridge.setCrossMintableERC20Code(ICrossMintableERC20Code(address(code)));
        console.log("CrossMintableERC20V2Code set to bridge");
        vm.stopBroadcast();
    }
}
