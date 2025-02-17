// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeEthereum} from "../src/BridgeEthereum.sol";
import {CrossToken} from "../test/token/CrossToken.sol";
import {Multicall} from "./Multicall.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {CrossMintableERC20, CrossMintableERC20Code} from "../src/CrossMintableERC20.sol";
import {Script, console} from "forge-std/Script.sol";

contract BridgeEthereumScript is Script {
    Multicall public multicall;
    BridgeEthereum public bridgeEthereum;

    address public VALIDATOR1;
    address public VALIDATOR2;
    address public VALIDATOR3;
    address public VALIDATOR4;
    address public VALIDATOR5;
    address public OWNER;
    address public REWARD;
    address[] public VALIDATORS;

    CrossToken public cross;
    IERC20 public xcross;

    function setUp() public {
        VALIDATOR1 = address(0x4631605A0F7E27a29b42e7814f26aa713016afcd);
        VALIDATOR2 = address(0x45D12F118b8706cB7b2A3d09C6F66632Fcf467A8);
        VALIDATOR3 = address(0x8190E3891F232DA7AD15c53c7464f939D7145Dcf);
        VALIDATOR4 = address(0x8A71e0F79c1F5806d372F14FfaDDF03Ba5F7Adc5);
        VALIDATOR5 = address(0xaAca7edC8a89eF5d7Dd62d29f06f3726041Ba08E);
        OWNER = address(0x14d95B8ecc31875F409c1fe5Cd58B3B5CDDfdDFD);
        REWARD = address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8);
        xcross = IERC20(address(1));
        VALIDATORS = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4, VALIDATOR5];
    }

    function run() public {
        vm.startBroadcast();

        cross = new CrossToken(OWNER, 1000000000);
        multicall = new Multicall();

        BridgeEthereum bridgeEthereumImpl = new BridgeEthereum();
        bridgeEthereum = BridgeEthereum(address(new ERC1967Proxy(address(bridgeEthereumImpl), bytes(""))));
        bridgeEthereum.initialize(IERC20(address(cross)), REWARD, address(0));

        bridgeEthereum.setValidators(VALIDATORS);
        bridgeEthereum.addToken(IERC20(address(cross)), xcross);

        vm.stopBroadcast();
    }
}
