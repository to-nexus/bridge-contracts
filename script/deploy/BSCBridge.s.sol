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
        address bscBridge = deployBSCBridgeProxy();
        _setupBridge(bscBridge);
    }

    function deployBSCBridgeProxy() public returns (address) {
        crossChainID = vm.envUint(BridgeScript.BRIDGE_CROSS_CHAIN_ID);
        cross = vm.envAddress(BRIDGE_CROSS);
        crossInitialSupply = vm.envUint(BRIDGE_CROSS_INITIAL_SUPPLY) * 1 ether;

        console.log("crossChainID", crossChainID);
        console.log("cross", cross);
        console.log("crossInitialSupply", crossInitialSupply);

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
