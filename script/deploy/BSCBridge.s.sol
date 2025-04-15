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

contract BSCBridgeScript is Script, BridgeScript {
    string IMPLEMENTATION = "IMPLEMENTATION";

    address impl;

    function setUp() public override {
        super.setUp();
        
        // load env variables
        impl = vm.envAddress(IMPLEMENTATION);

        // print env variables
        console.log("impl", impl);
    }

    /**
     * @notice Setup BSC Bridge contract and initialize
     * command
     * forge script ./script/deploy/Bridge.sol:BridgeScript \
     * -f $RPC_URL \
     * --sender $SENDER \
     * --sig "setupBSCBridge()"
     */
    function setupBSCBridge() public {
        crossChainID = vm.envUint(BridgeScript.BRIDGE_CROSS_CHAIN_ID);
        console.log("crossChainID", crossChainID);

        BSCBridge bscBridge = BSCBridge(deployBSCBridgeProxy());

        _setupBridge(address(bscBridge));
    }

    function deployBSCBridgeProxy() public returns (address) {
        vm.broadcast();
        address proxy = address(
            new ERC1967Proxy(
                impl,
                abi.encodeCall(
                    BSCBridge.initializeBSCBridge, (owner, dev, threshold, crossChainID, cross, crossInitialSupply)
                )
            )
        );
        console.log("BSCBridgeProxy will be deployed to", proxy);
        return proxy;
    }
}
