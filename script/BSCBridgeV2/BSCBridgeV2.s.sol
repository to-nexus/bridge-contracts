// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BSCBridgeV2} from "../../src/BSCBridgeV2.sol";

import {Const} from "../../src/lib/Const.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";
import {Upgrades} from "lib/openzeppelin-foundry-upgrades/src/Upgrades.sol";

contract BSCBridgeV2Script is Script {
    event BridgeInitiated(
        uint indexed toChainID,
        uint indexed index,
        IERC20 fromToken,
        IERC20 toToken,
        address indexed from,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes extraData,
        uint time
    );

    address public deadWallet;

    function setUp() public {
        deadWallet = address(0x000000000000000000000000000000000000dEaD);
    }

    function deployBSCBridgeV2Impl() public returns (address) {
        vm.broadcast();
        BSCBridgeV2 impl = new BSCBridgeV2();
        console.log("BSCBridgeV2Logic will be deployed to", address(impl));
        return address(impl);
    }

    function upgradeBSCBridgeV2(BSCBridgeV2 proxy, address impl, uint crossChainID, IERC20 cross) public {
        console.log("Legacy Implementation", Upgrades.getImplementationAddress(address(proxy)));

        console.log("new Implementation", impl);
        vm.broadcast();
        proxy.upgradeToAndCall(impl, abi.encodeCall(BSCBridgeV2.reinitializeBSCBridge, (crossChainID, cross)));

        console.log("Upgraded Implementation", Upgrades.getImplementationAddress(address(proxy)));
    }
}
