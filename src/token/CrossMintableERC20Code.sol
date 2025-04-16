// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// import {ERC20, IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20.sol";

import {CrossMintableERC20} from "./CrossMintableERC20.sol";
import {ICrossMintableERC20Code} from "./ICrossMintableERC20Code.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

contract CrossMintableERC20Code is ICrossMintableERC20Code {
    error ERC20CodeNotBridge(address account);

    // slot 0 : bridge
    address public bridge;

    constructor(address _bridge) {
        bridge = _bridge;
    }

    modifier onlyBridge() {
        require(msg.sender == bridge, ERC20CodeNotBridge(msg.sender));
        _;
    }

    function createCrossMintableERC20(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        onlyBridge
        returns (address tokenAddress)
    {
        bytes32 salt = keccak256(abi.encodePacked(remoteChainID, remoteToken));
        bytes memory bytecode = abi.encodePacked(
            type(CrossMintableERC20).creationCode,
            abi.encode(
                bridge, string(abi.encodePacked("Cross Bridge ", symbol)), abi.encodePacked(symbol, "x"), decimals
            )
        ); // Combine creation code and constructor arguments
        tokenAddress = Create2.deploy(0, salt, bytecode); // Deploy the wrapped token using Create2
    }
}
