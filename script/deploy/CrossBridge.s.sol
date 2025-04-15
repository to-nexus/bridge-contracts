// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BSCBridge} from "../../src/BSCBridge.sol";
import {BridgeVerifier} from "../../src/BridgeVerifier.sol";
import {BaseBridge, CrossBridge} from "../../src/CrossBridge.sol";

import {PriceFeed} from "../../src/PriceFeed.sol";
import {CrossMintableERC20Code} from "../../src/token/CrossMintableERC20Code.sol";
import {BridgeScript} from "./Bridge.s.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

contract CrossBridgeScript is Script, BridgeScript {
    CrossBridge crossBridge;

    string CROSS_SUPPLY_LIMIT = "CROSS_SUPPLY_LIMIT";

    uint crossSupplyLimit;

    function setUp() public override {
        super.setUp();

        crossBridge = CrossBridge(address(0xb81d6e000000000000000000000000000000C0de));

        crossSupplyLimit = vm.envUint(CROSS_SUPPLY_LIMIT) * 1 ether;
        console.log("crossSupplyLimit", crossSupplyLimit);
    }

    /**
     * @notice Setup Cross Bridge contract and initialize
     * command
     * forge script ./script/deploy/Bridge.sol:BridgeScript \
     * -f $RPC_URL \
     * --sender $SENDER \
     * --sig "setupCrossBridge()"
     * $ARGUMENTS
     */
    function setupCrossBridge() public {
        bscChainID = vm.envUint(BridgeScript.BRIDGE_BSC_CHAIN_ID);
        console.log("bscChainID", bscChainID);

        vm.broadcast();
        crossBridge.initializeCrossBridge(owner, dev, threshold, bscChainID, cross, crossInitialSupply);

        _setupBridge(address(crossBridge));

        vm.broadcast();
        crossBridge.setCrossSupplyLimit(crossSupplyLimit);
    }

    function initializeCrossBridge() public {
        vm.broadcast();
        crossBridge.initializeCrossBridge(owner, dev, threshold, bscChainID, cross, crossInitialSupply);
    }
}
