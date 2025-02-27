// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// import {ERC20, IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20.sol";

import {CrossMintableERC20} from "./CrossMintableERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract CrossMintableERC20FactoryCode {
    function code() external pure returns (bytes memory) {
        return type(CrossMintableERC20Factory).creationCode;
    }
}

contract CrossMintableERC20Factory is Ownable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error errCrossMintableERC20FactoryNotBridge(address account);
    error errCrossMintableERC20FactoryNotExist(address token);

    // slot 0 : owner, slot 1 : bridge
    address public bridge;
    EnumerableSet.AddressSet private _tokens;

    constructor(address _bridge) Ownable(_bridge) {
        bridge = _bridge;
    }

    modifier onlyBridge() {
        require(_msgSender() == bridge, errCrossMintableERC20FactoryNotBridge(_msgSender()));
        _;
    }

    function createCrossMintableERC20(bytes32 salt, string memory name, string memory symbol, uint8 decimals)
        external
        onlyBridge
        returns (address tokenAddress)
    {
        bytes memory bytecode =
            abi.encodePacked(type(CrossMintableERC20).creationCode, abi.encode(bridge, name, symbol, decimals)); // Combine creation code and constructor arguments
        tokenAddress = Create2.deploy(0, salt, bytecode); // Deploy the wrapped token using Create2
    }

    function pause(CrossMintableERC20 token) external onlyOwner {
        require(_tokens.contains(address(token)), errCrossMintableERC20FactoryNotExist(address(token)));
        token.pause();
    }

    function unpause(CrossMintableERC20 token) external onlyOwner {
        require(_tokens.contains(address(token)), errCrossMintableERC20FactoryNotExist(address(token)));
        token.unpause();
    }
}
