// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeRegistry} from "../interface/IBridgeRegistry.sol";
import {CrossMintableERC20Factory, CrossMintableERC20FactoryCode} from "../token/CrossMintableERC20Factory.sol";

abstract contract BridgeRegistry is OwnableUpgradeable, IBridgeRegistry {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Math for uint;

    error BridgeRegistryNotExistIndex(uint index);
    error BridgeRegistryNotExistToken(address token);
    error BridgeRegistryNotExistChain(uint chainID);
    error BridgeRegistryAleadyExistToken(address token);
    error BridgeRegistryAleadyExistChain(uint chainID);
    error BridgeRegistryAleadyExistIndex(uint index);
    error BridgeRegistryAleadyExistFactory(CrossMintableERC20Factory factory);
    error BridgeRegistryCanNotZeroAddress(string name);
    error BridgeRegistryCanNotZeroValue(string name);
    error BridgeRegistryInvalidToken(address token);
    error BridgeRegistryBalanceLow(uint remoteChainID, address token, uint value);
    error BridgeRegistryTokenPaused(address token);
    error BridgeRegistryAleadyPaused(address token);
    error BridgeRegistryNotPaused(address token);
    error BridgeRegistryFactoryNotSet();
    error BridgeRegistryInvalidRate(uint localTokenRate, uint remoteTokenRate);

    event TokenPairRegistered(
        uint indexed remoteChainID,
        bool isOrigin,
        address indexed localToken,
        address indexed remoteToken,
        uint localTokenRate,
        uint remoteTokenRate
    );
    event TokenPairUnregistered(uint indexed remoteChainID, address indexed localToken);
    event TokenPairPaused(uint indexed remoteChainID, address indexed token);
    event TokenPairUnpaused(uint indexed remoteChainID, address indexed token);
    event CrossMintableERC20FactorySet(CrossMintableERC20Factory indexed factory);

    CrossMintableERC20Factory public crossMintableERC20Factory;

    EnumerableSet.UintSet private _chains;
    mapping(uint => Chain) internal _chainData; // chain id : Chain
    mapping(uint => mapping(address => ExchangeRate)) internal _exchangeRate;

    uint[47] private __gap;

    function __BridgeRegistry_init(address _crossMintableERC20FactoryCode) internal {
        if (_crossMintableERC20FactoryCode != address(0)) _setERC20Factory(_crossMintableERC20FactoryCode);
    }

    modifier onlyValidToken(uint remoteChainID, address _token) {
        require(_chainData[remoteChainID].tokens.contains(_token), BridgeRegistryInvalidToken(_token));
        require(!_chainData[remoteChainID].tokenPairs[_token].paused, BridgeRegistryTokenPaused(_token));
        _;
    }

    function allChainIDs() external view returns (uint[] memory) {
        return _chains.values();
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

    function allRevertedIndex(uint remoteChainID) external view returns (uint[] memory) {
        return _chainData[remoteChainID].reverted.index.values();
    }

    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory) {
        return _chainData[remoteChainID].tokenPairs[token];
    }

    function getNextInitiateIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].index.initiate + 1;
    }

    function getNextFinalizeIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].index.finalize + 1;
    }

    function revertedArguments(uint remoteChainID, uint index) public view returns (FinalizeArguments memory) {
        return _chainData[remoteChainID].reverted.data[index];
    }

    function revertedReason(uint remoteChainID, uint index) public view returns (string memory) {
        return _chainData[remoteChainID].reverted.reason[index];
    }

    function _incrementInitiateIndex(uint remoteChainID) internal {
        unchecked {
            ++_chainData[remoteChainID].index.initiate;
        }
    }

    function _incrementFinalizeIndex(uint remoteChainID) internal {
        unchecked {
            ++_chainData[remoteChainID].index.finalize;
        }
    }

    function _isOriginToken(uint remoteChainID, address token) internal view returns (bool) {
        return _chainData[remoteChainID].tokenPairs[token].isOrigin;
    }

    function _getRemoteToken(uint remoteChainID, address localToken) internal view returns (address remoteToken) {
        return _chainData[remoteChainID].tokenPairs[localToken].remoteToken;
    }

    function _depositToken(uint remoteChainID, address token, uint value) internal {
        _chainData[remoteChainID].tokenPairs[token].deposited += value;
    }

    function _withdrawToken(uint remoteChainID, address token, uint value) internal {
        Chain storage chain = _chainData[remoteChainID];
        (bool ok, uint deposited) = chain.tokenPairs[token].deposited.trySub(value);
        require(ok, BridgeRegistryBalanceLow(chain.remoteChainID, token, value));

        chain.tokenPairs[token].deposited = deposited;
    }

    function _setRevertedArguments(FinalizeArguments calldata args, string memory reason) internal {
        Chain storage chain = _chainData[args.remoteChainID];
        uint index = args.index;
        require(chain.reverted.index.add(index), BridgeRegistryAleadyExistIndex(index));

        chain.reverted.data[index] = args;
        chain.reverted.reason[index] = reason;
    }

    function _removeRevertedArguments(uint remoteChainID, uint index) internal {
        Chain storage chain = _chainData[remoteChainID];
        require(chain.reverted.index.remove(index), BridgeRegistryNotExistIndex(index));
        delete (chain.reverted.reason[index]);
        delete (chain.reverted.data[index]);
    }

    // set function
    function createToken(
        uint remoteChainID,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        string memory symbol,
        uint8 decimals
    ) external onlyOwner returns (address tokenAddress) {
        require(address(crossMintableERC20Factory) != address(0), BridgeRegistryFactoryNotSet());
        tokenAddress = crossMintableERC20Factory.createCrossMintableERC20(
            keccak256(abi.encodePacked(remoteToken)),
            string(abi.encodePacked("Cross Bridge ", symbol)),
            symbol,
            decimals
        );
        _registerToken(remoteChainID, false, tokenAddress, remoteToken, localTokenRate, remoteTokenRate);
    }

    function registerToken(
        uint remoteChainID,
        bool isOrigin,
        address localToken,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate
    ) external onlyOwner {
        _registerToken(remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate);
    }

    function unregisterToken(uint remoteChainID, address token) external onlyOwner {
        require(_chains.contains(remoteChainID), BridgeRegistryNotExistChain(remoteChainID));
        require(token != address(0), BridgeRegistryCanNotZeroAddress("token"));

        Chain storage chain = _chainData[remoteChainID];
        require(chain.tokens.remove(token), BridgeRegistryNotExistToken(token));
        delete (chain.tokenPairs[token]);

        emit TokenPairUnregistered(remoteChainID, token);
    }

    function pauseToken(uint remoteChainID, address token) external onlyOwner {
        Chain storage chain = _chainData[remoteChainID];
        require(chain.tokens.contains(token), BridgeRegistryNotExistToken(token));
        require(!chain.tokenPairs[token].paused, BridgeRegistryAleadyPaused(token));
        chain.tokenPairs[token].paused = true;

        emit TokenPairPaused(remoteChainID, token);
    }

    function unpauseToken(uint remoteChainID, address token) external onlyOwner {
        Chain storage chain = _chainData[remoteChainID];
        require(chain.tokens.contains(token), BridgeRegistryNotExistToken(token));
        require(chain.tokenPairs[token].paused, BridgeRegistryNotPaused(token));
        chain.tokenPairs[token].paused = false;

        emit TokenPairUnpaused(remoteChainID, token);
    }

    function setERC20Factory(address factoryCode) external onlyOwner {
        _setERC20Factory(factoryCode);
    }

    function _setERC20Factory(address factoryCode) private {
        require(factoryCode != address(0), BridgeRegistryCanNotZeroAddress("factoryCode"));
        require(
            address(crossMintableERC20Factory) == address(0),
            BridgeRegistryAleadyExistFactory(crossMintableERC20Factory)
        );

        bytes memory bytecode =
            abi.encodePacked(CrossMintableERC20FactoryCode(factoryCode).code(), abi.encode(address(this)));
        crossMintableERC20Factory =
            CrossMintableERC20Factory(Create2.deploy(0, keccak256(abi.encodePacked(block.chainid)), bytecode));

        emit CrossMintableERC20FactorySet(crossMintableERC20Factory);
    }

    function _registerToken(
        uint remoteChainID,
        bool isOrigin,
        address localToken,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate
    ) private {
        if (_chains.contains(remoteChainID)) {
            require(_chains.add(remoteChainID), BridgeRegistryAleadyExistChain(remoteChainID));
        }

        Chain storage chain = _chainData[remoteChainID];
        require(localToken != address(0), BridgeRegistryCanNotZeroAddress("localToken"));
        require(chain.tokens.add(localToken), BridgeRegistryAleadyExistToken(localToken));

        chain.tokenPairs[localToken] = IBridgeRegistry.TokenPair({
            localToken: localToken,
            remoteToken: remoteToken,
            isOrigin: isOrigin,
            paused: false,
            deposited: 0
        });

        if (localTokenRate != remoteTokenRate) {
            require(localTokenRate != 0, BridgeRegistryCanNotZeroValue("remoteTokenRate"));
            require(remoteTokenRate != 0, BridgeRegistryCanNotZeroValue("remoteTokenRate"));

            if (localTokenRate > remoteTokenRate) {
                require(
                    remoteTokenRate == 1 && localTokenRate % 10 == 0,
                    BridgeRegistryInvalidRate(localTokenRate, remoteTokenRate)
                );
            } else {
                require(
                    localTokenRate == 1 && remoteTokenRate % 10 == 0,
                    BridgeRegistryInvalidRate(localTokenRate, remoteTokenRate)
                );
            }

            _exchangeRate[remoteChainID][localToken] =
                ExchangeRate({localTokenRate: localTokenRate, remoteTokenRate: remoteTokenRate});
        }

        emit TokenPairRegistered(remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate);
    }
}
