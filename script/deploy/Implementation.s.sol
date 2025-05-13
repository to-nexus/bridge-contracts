// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BSCBridge} from "../../src/BSCBridge.sol";
import {BaseBridge} from "../../src/BaseBridge.sol";
import {CrossBridge} from "../../src/CrossBridge.sol";
import {PriceFeed} from "../../src/PriceFeed.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {Const} from "../../src/lib/Const.sol";
import {Script, console} from "forge-std/Script.sol";

contract ImplementationScript is Script {
    function setUp() public {}

    function deployImpl() public {
        deployPriceFeedImpl();
        deployBaseBridgeImpl();
    }

    function deployBSCImpl() public {
        deployPriceFeedImpl();
        deployBSCBridgeImpl();
    }

    function deployPriceFeedImpl() public {
        vm.broadcast();
        PriceFeed impl = new PriceFeed();
        console.log("PriceFeed Implementation will be deployed to:", address(impl));
    }

    function deployBSCBridgeImpl() public {
        vm.broadcast();
        BSCBridge impl = new BSCBridge();
        console.log("BSCBridge Implementation will be deployed to:", address(impl));
    }

    function deployBaseBridgeImpl() public {
        vm.broadcast();
        BaseBridge impl = new BaseBridge();
        console.log("BaseBridge Implementation will be deployed to:", address(impl));
    }
}
