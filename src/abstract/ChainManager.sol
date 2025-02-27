// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IChainManager} from "../interface/IChainManager.sol";
import {ChainLib} from "../lib/ChainLib.sol";
import {CrossMintableERC20Factory, CrossMintableERC20FactoryCode} from "../token/CrossMintableERC20Factory.sol";

abstract contract ChainManager is OwnableUpgradeable, IChainManager {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Math for uint;
    using ChainLib for Chain;

    error errChainManagerTokenFactoryNotSet();
    error errChainManagerChainAleadyExist(uint chainID);
    error errChainManagerChainNotExist(uint chainID);
    error errChainManagerInvalidToken(address token);
    error errChainManagerTokenPaused(address token);
    error errChainManagerInsufficientBalance(uint remoteChainID, address token, uint value);

    event TokenPairRegistered(
        uint indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken
    );
    event TokenPairUnregistered(uint indexed remoteChainID, address indexed localToken);
    event TokenPairPaused(uint indexed remoteChainID, address indexed token);
    event TokenPairUnpaused(uint indexed remoteChainID, address indexed token);
    event ChainSet(uint indexed remoteChainID);

    CrossMintableERC20Factory public crossMintableERC20Factory;
    EnumerableSet.UintSet private _chains;

    mapping(uint => Chain) internal _chainData; // chain id : Chain

    uint[47] private __gap;

    function __ChainManager_init(address _crossMintableERC20FactoryCode) internal {
        if (_crossMintableERC20FactoryCode == address(0)) return;

        bytes memory bytecode = abi.encodePacked(
            CrossMintableERC20FactoryCode(_crossMintableERC20FactoryCode).code(), abi.encode(address(this))
        );
        address addr = Create2.deploy(0, keccak256(abi.encodePacked(block.chainid)), bytecode);
        require(addr != address(0), "zero");

        crossMintableERC20Factory = CrossMintableERC20Factory(addr);
    }

    modifier onlyValidToken(uint remoteChainID, address _token) {
        require(_chainData[remoteChainID].tokens.contains(_token), errChainManagerInvalidToken(_token));
        require(!_chainData[remoteChainID].tokenPairs[_token].paused, errChainManagerTokenPaused(_token));
        _;
    }

    // index
    function getNextInitiateIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].getNextInitiateIndex();
    }

    function getNextFinalizeIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].getNextFinalizeIndex();
    }

    function _incrementInitiateIndex(uint remoteChainID) internal {
        _chainData[remoteChainID].incrementInitiateIndex();
    }

    function _incrementFinalizeIndex(uint remoteChainID) internal {
        _chainData[remoteChainID].incrementFinalizeIndex();
    }

    // reverted args
    function revertedArguments(uint remoteChainID, uint index) public view returns (FinalizeArguments memory) {
        return _chainData[remoteChainID].getRevertedArguments(index);
    }

    function revertedReason(uint remoteChainID, uint index) public view returns (string memory) {
        return _chainData[remoteChainID].getRevertedReason(index);
    }

    function _getRevertedArguments(uint remoteChainID, uint index)
        internal
        view
        returns (FinalizeArguments memory args)
    {
        return _chainData[remoteChainID].getRevertedArguments(index);
    }

    function _setRevertedArguments(
        uint remoteChainID,
        uint index,
        IERC20 token,
        address to,
        uint value,
        bytes calldata extraData,
        string memory reason
    ) internal {
        _chainData[remoteChainID].setRevertedData(
            FinalizeArguments({index: index, token: token, to: to, value: value, extraData: extraData}), reason
        );
    }

    function _removeRevertedArguments(uint remoteChainID, uint index) internal {
        _chainData[remoteChainID].removeRevertedData(index);
    }

    // token
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory) {
        return _chainData[remoteChainID].tokenPairs[token];
    }

    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory) {
        Chain storage chain = _chainData[remoteChainID];
        address[] memory tokens = chain.tokens.values();
        TokenPair[] memory _pairs = new TokenPair[](tokens.length);
        for (uint i = 0; i < tokens.length; ++i) {
            _pairs[i] = chain.tokenPairs[tokens[i]];
        }
        return _pairs;
    }

    function createToken(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        onlyOwner
        returns (address tokenAddress)
    {
        require(address(crossMintableERC20Factory) != address(0), errChainManagerTokenFactoryNotSet());
        tokenAddress = crossMintableERC20Factory.createCrossMintableERC20(
            keccak256(abi.encodePacked(remoteToken)),
            string(abi.encodePacked("Cross Bridge ", symbol)),
            symbol,
            decimals
        );
        registerToken(remoteChainID, false, tokenAddress, remoteToken); // Register the new token with the bridge
    }

    function setChain(uint remoteChainID) public onlyOwner {
        require(_chains.add(remoteChainID), errChainManagerChainAleadyExist(remoteChainID));
    }

    function registerToken(uint remoteChainID, bool isOrigin, address localToken, address remoteToken)
        public
        onlyOwner
    {
        require(_chains.contains(remoteChainID), errChainManagerChainNotExist(remoteChainID));
        _chainData[remoteChainID].registerToken(isOrigin, localToken, remoteToken);
        emit TokenPairRegistered(remoteChainID, isOrigin, localToken, remoteToken);
    }

    function unregisterToken(uint remoteChainID, address token) external onlyOwner {
        require(_chains.contains(remoteChainID), errChainManagerChainNotExist(remoteChainID));
        _chainData[remoteChainID].unregisterToken(token);
        emit TokenPairUnregistered(remoteChainID, token);
    }

    function pauseToken(uint remoteChainID, address token) external onlyOwner {
        _chainData[remoteChainID].pauseToken(token);
        emit TokenPairPaused(remoteChainID, token);
    }

    function unpauseToken(uint remoteChainID, address token) external onlyOwner {
        _chainData[remoteChainID].unpauseToken(token);
        emit TokenPairUnpaused(remoteChainID, token);
    }

    function _depositToken(uint remoteChainID, address token, uint value) internal {
        _chainData[remoteChainID].depositToken(token, value);
    }

    function _withdrawToken(uint remoteChainID, address token, uint value) internal {
        _chainData[remoteChainID].withdrawToken(token, value);
    }

    function _isOriginToken(uint remoteChainID, address token) internal view returns (bool) {
        return _chainData[remoteChainID].tokenPairs[token].isOrigin;
    }

    function _getRemoteToken(uint remoteChainID, address localToken) internal view returns (address remoteToken) {
        return _chainData[remoteChainID].tokenPairs[localToken].remoteToken;
    }
}
