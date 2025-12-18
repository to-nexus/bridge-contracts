// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../../src/BaseBridge.sol";
import {CrossBridge} from "../../src/CrossBridge.sol";
import {PriceFeed} from "../../src/PriceFeed.sol";

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "lib/openzeppelin-foundry-upgrades/src/Upgrades.sol";

import {Const} from "../../src/lib/Const.sol";
import {Script, console} from "forge-std/Script.sol";

contract ImplementationScript is Script {
    function setUp() public {}

    function deployPriceFeedImpl() public {
        vm.broadcast();
        PriceFeed impl = new PriceFeed();
        console.log("PriceFeed Implementation will be deployed to:", address(impl));
    }

    function deployBaseBridgeImpl() public {
        vm.broadcast();
        BaseBridge impl = new BaseBridge();
        console.log("BaseBridge Implementation will be deployed to:", address(impl));
    }

    function deployCrossBridgeImpl() public {
        vm.broadcast();
        CrossBridge impl = new CrossBridge();
        console.log("CrossBridge Implementation will be deployed to:", address(impl));
    }

    function upgradeBridgeImplWithoutCall(address proxy, address impl) public {
        console.log("Legacy Implementation", Upgrades.getImplementationAddress(address(proxy)));

        console.log("new Implementation", impl);
        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(impl, "");

        console.log("Upgraded Implementation", Upgrades.getImplementationAddress(address(proxy)));
    }
}
