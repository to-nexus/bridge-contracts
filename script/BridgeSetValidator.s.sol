// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeStandard} from "../src/abstract/BridgeStandard.sol";
import {Script, console} from "forge-std/Script.sol";

contract BridgeEthereumScript is Script {
    BridgeStandard public bridge;

    address public VALIDATOR1;
    address public VALIDATOR2;
    address public VALIDATOR3;
    address public VALIDATOR4;
    address public VALIDATOR5;
    address[] public VALIDATORS;

    function setUp() public {
        VALIDATOR1 = address(0x4631605A0F7E27a29b42e7814f26aa713016afcd);
        VALIDATOR2 = address(0x45D12F118b8706cB7b2A3d09C6F66632Fcf467A8);
        VALIDATOR3 = address(0x8190E3891F232DA7AD15c53c7464f939D7145Dcf);
        VALIDATOR4 = address(0x8A71e0F79c1F5806d372F14FfaDDF03Ba5F7Adc5);
        VALIDATOR5 = address(0xaAca7edC8a89eF5d7Dd62d29f06f3726041Ba08E);
        VALIDATORS = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4, VALIDATOR5];
        // bridge = BridgeStandard(address(0x8d4B69F0308293ed37a154369E5A2c91A13CCD65)); // ethereum
        bridge = BridgeStandard(address(0x761b68081fb50359540ac89dCE0fD1FF2C5C5cC4)); // cross
    }

    function run() public {
        vm.startBroadcast();

        bridge.setValidators(VALIDATORS);

        vm.stopBroadcast();
    }
}
