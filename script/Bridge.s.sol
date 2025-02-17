// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeEthereum} from "../src/BridgeEthereum.sol";

import {CrossMintableERC20, CrossMintableERC20Code} from "../src/CrossMintableERC20.sol";
import {IBridgeStandard} from "../src/interface/IBridgeStandard.sol";
import {CrossToken} from "../test/token/CrossToken.sol";
import {Multicall} from "./Multicall.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Script, console} from "forge-std/Script.sol";

contract BridgeEthereumScript is Script {
    bytes32 public constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    BridgeEthereum public bridgeEthereum;
    IERC20 public cross;

    address public USER;
    uint public USER_PK;

    function setUp() public {
        USER = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        USER_PK = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        bridgeEthereum = BridgeEthereum(address(0xF2620CcF88EE16040f1bC6882cF41597CC8AAbd3));
        cross = IERC20(address(0x33981fEb1022A1dEd330456905583ca0b9E9fd08));
    }

    function run() public {
        vm.startBroadcast();

        uint amount = 10 ether;
        uint nonce = IERC20Permit(address(cross)).nonces(USER);
        bytes32 h = keccak256(abi.encode(PERMIT_TYPEHASH, USER, address(bridgeEthereum), amount, nonce, type(uint).max));
        bytes32 hash = MessageHashUtils.toTypedDataHash(IERC20Permit(address(cross)).DOMAIN_SEPARATOR(), h);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(USER_PK, hash);

        bridgeEthereum.bridgeCross(
            USER,
            amount,
            IBridgeStandard.PermitArguments(IERC20Permit(address(cross)), USER, amount, type(uint).max, v, r, s),
            new bytes[](0)
        );

        vm.stopBroadcast();
    }
}
