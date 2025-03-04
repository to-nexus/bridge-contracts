// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// import {ERC20, IERC20} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20.sol";

import {CrossMintableERC20, ICrossMintableERC20} from "./CrossMintableERC20.sol";
import {ICrossMintableERC20Factory} from "./ICrossMintableERC20Factory.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract CrossMintableERC20FactoryCode {
    function code() external pure returns (bytes memory) {
        return type(CrossMintableERC20Factory).creationCode;
    }
}

contract CrossMintableERC20Factory is ICrossMintableERC20Factory {
    using EnumerableSet for EnumerableSet.AddressSet;

    error errCrossMintableERC20FactoryNotBridge(address account);
    error errCrossMintableERC20FactoryNotExist(address token);

    // slot 0 : bridge
    address public bridge;
    EnumerableSet.AddressSet private _tokens;

    constructor(address _bridge) {
        bridge = _bridge;
    }

    modifier onlyBridge() {
        require(msg.sender == bridge, errCrossMintableERC20FactoryNotBridge(msg.sender));
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

    function pause(ICrossMintableERC20 token) external onlyBridge {
        require(_tokens.contains(address(token)), errCrossMintableERC20FactoryNotExist(address(token)));
        token.pause();
    }

    function unpause(ICrossMintableERC20 token) external onlyBridge {
        require(_tokens.contains(address(token)), errCrossMintableERC20FactoryNotExist(address(token)));
        token.unpause();
    }

    function allTokens() external view returns (address[] memory) {
        return _tokens.values();
    }

    function tokenByIndex(uint index) external view returns (address) {
        return _tokens.at(index);
    }

    function tokensLength() external view returns (uint) {
        return _tokens.length();
    }
}
