// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {CrossMintableERC20V2} from "./CrossMintableERC20V2.sol";
import {ICrossMintableERC20Code} from "./ICrossMintableERC20Code.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin/contracts/access/extensions/AccessControlDefaultAdminRules.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

contract CrossMintableERC20V2Code is AccessControlDefaultAdminRules, ICrossMintableERC20Code {
    error ERC20CodeNotBridge(address account);

    constructor(address initialOwner, address initialBridge) AccessControlDefaultAdminRules(0, initialOwner) {
        grantRole(Const.BRIDGE_ROLE, initialBridge);
    }

    function createCrossMintableERC20(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        onlyRole(Const.BRIDGE_ROLE)
        returns (address tokenAddress)
    {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        bytes memory bytecode = abi.encodePacked(
            type(CrossMintableERC20V2).creationCode,
            abi.encode(
                defaultAdmin(), // Initial owner
                _msgSender(), // Initial minter
                string(abi.encodePacked("Cross Bridge ", symbol)),
                abi.encodePacked(symbol, "x"),
                decimals
            )
        ); // Combine creation code and constructor arguments
        tokenAddress = Create2.deploy(0, salt, bytecode); // Deploy the wrapped token using Create2
    }
}
