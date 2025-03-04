// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeRegistry} from "../interface/IBridgeRegistry.sol";
import {ICrossMintableERC20Factory} from "../token/ICrossMintableERC20Factory.sol";

abstract contract BridgeRegistry is OwnableUpgradeable, IBridgeRegistry {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Math for uint;

    error RegistryNotExistToken(address token);
    error RegistryNotExistIndex(uint index);
    error RegistryExistFactory(address factory);
    error RegistryExistIndex(uint index);
    error RegistryExistToken(address token);
    error RegistryZeroAddress();
    error RegistryZeroLocalRate();
    error RegistryZeroRemoteRate();
    error RegistryTokenPaused(address token);
    error RegistryNotPaused(address token);
    error RegistryInvalidRate(uint localTokenRate, uint remoteTokenRate);
    error RegistryBalanceLow(uint remoteChainID, address token, uint value);
    error RegistryFactoryNotSet();

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
    event CrossMintableERC20FactorySet(address indexed factory);

    ICrossMintableERC20Factory public crossMintableERC20Factory;

    EnumerableSet.UintSet private _chains;

    mapping(uint => Chain) internal _chainData; // chain id : Chain
    mapping(uint => EnumerableSet.AddressSet) internal _tokens; // chain id : tokens
    mapping(uint => mapping(address => TokenPair)) internal _tokenPairs; // chain id : token address : TokenPair

    uint[44] private __gap;

    modifier onlyValidToken(uint remoteChainID, address token) {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        require(!_tokenPairs[remoteChainID][token].paused, RegistryTokenPaused(token));
        _;
    }

    // external nonpayable
    function createToken(
        uint remoteChainID,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        string memory symbol,
        uint8 decimals
    ) external onlyOwner returns (address tokenAddress) {
        require(address(crossMintableERC20Factory) != address(0), RegistryFactoryNotSet());
        tokenAddress = crossMintableERC20Factory.createCrossMintableERC20(
            keccak256(abi.encodePacked(remoteToken)),
            string(abi.encodePacked("Cross Bridge ", symbol)), // @DEV: name을 외부입력을 받을건지 확인
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
        require(_tokens[remoteChainID].remove(token), RegistryNotExistToken(token));
        delete (_tokenPairs[remoteChainID][token]);

        emit TokenPairUnregistered(remoteChainID, token);
    }

    function pauseToken(uint remoteChainID, address token) external onlyOwner {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        require(!_tokenPairs[remoteChainID][token].paused, RegistryTokenPaused(token));
        _tokenPairs[remoteChainID][token].paused = true;

        emit TokenPairPaused(remoteChainID, token);
    }

    function unpauseToken(uint remoteChainID, address token) external onlyOwner {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        require(_tokenPairs[remoteChainID][token].paused, RegistryNotPaused(token));
        _tokenPairs[remoteChainID][token].paused = false;

        emit TokenPairUnpaused(remoteChainID, token);
    }

    function setCrossMintableERC20Factory(ICrossMintableERC20Factory _crossMintableERC20Factory) external onlyOwner {
        require(
            address(crossMintableERC20Factory) == address(0), RegistryExistFactory(address(crossMintableERC20Factory))
        );
        require(address(_crossMintableERC20Factory) != address(0), RegistryZeroAddress());

        crossMintableERC20Factory = _crossMintableERC20Factory;
        emit CrossMintableERC20FactorySet(address(crossMintableERC20Factory));
    }

    // external view
    function allChainIDs() external view returns (uint[] memory) {
        return _chains.values();
    }

    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory) {
        address[] memory tokens = _tokens[remoteChainID].values();
        TokenPair[] memory _pairs = new TokenPair[](tokens.length);
        for (uint i = 0; i < tokens.length; ++i) {
            _pairs[i] = _tokenPairs[remoteChainID][tokens[i]];
        }
        return _pairs;
    }

    function allRevertedIndex(uint remoteChainID) external view returns (uint[] memory) {
        return _chainData[remoteChainID].reverted.index.values();
    }

    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory) {
        return _tokenPairs[remoteChainID][token];
    }

    // public view
    function getNextInitiateIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].initiateIndex + 1;
    }

    function getNextFinalizeIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].finalizeIndex + 1;
    }

    function revertedArguments(uint remoteChainID, uint index) public view returns (FinalizeArguments memory) {
        return _chainData[remoteChainID].reverted.data[index];
    }

    function revertedReason(uint remoteChainID, uint index) public view returns (string memory) {
        return _chainData[remoteChainID].reverted.reason[index];
    }

    // internal nonpayable
    function _incrementInitiateIndex(uint remoteChainID) internal {
        unchecked {
            ++_chainData[remoteChainID].initiateIndex;
        }
    }

    function _incrementFinalizeIndex(uint remoteChainID) internal {
        unchecked {
            ++_chainData[remoteChainID].finalizeIndex;
        }
    }

    function _depositToken(uint remoteChainID, address token, uint value) internal {
        _tokenPairs[remoteChainID][token].deposited += value;
    }

    function _withdrawToken(uint remoteChainID, address token, uint value) internal {
        (bool ok, uint deposited) = _tokenPairs[remoteChainID][token].deposited.trySub(value);
        require(ok, RegistryBalanceLow(remoteChainID, token, value));

        _tokenPairs[remoteChainID][token].deposited = deposited;
    }

    function _setRevertedArguments(FinalizeArguments calldata args, string memory reason) internal {
        Chain storage chain = _chainData[args.remoteChainID];
        uint index = args.index;
        require(chain.reverted.index.add(index), RegistryExistIndex(index));

        chain.reverted.data[index] = args;
        chain.reverted.reason[index] = reason;
    }

    function _removeRevertedArguments(uint remoteChainID, uint index) internal {
        Chain storage chain = _chainData[remoteChainID];
        require(chain.reverted.index.remove(index), RegistryNotExistIndex(index));
        delete (chain.reverted.reason[index]);
        delete (chain.reverted.data[index]);
    }

    // private functions
    function _registerToken(
        uint remoteChainID,
        bool isOrigin,
        address localToken,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate
    ) private {
        Chain storage chain = _chainData[remoteChainID];

        if (_chains.add(remoteChainID)) chain.remoteChainID = remoteChainID;

        require(localToken != address(0), RegistryZeroAddress());
        require(_tokens[remoteChainID].add(localToken), RegistryExistToken(localToken));

        if (localTokenRate != remoteTokenRate) {
            require(localTokenRate != 0, RegistryZeroLocalRate());
            require(remoteTokenRate != 0, RegistryZeroRemoteRate());

            if (localTokenRate > remoteTokenRate) {
                require(
                    remoteTokenRate == 1 && localTokenRate % 10 == 0,
                    RegistryInvalidRate(localTokenRate, remoteTokenRate)
                );
            } else {
                require(
                    localTokenRate == 1 && remoteTokenRate % 10 == 0,
                    RegistryInvalidRate(localTokenRate, remoteTokenRate)
                );
            }
        }

        _tokenPairs[remoteChainID][localToken] = IBridgeRegistry.TokenPair({
            localToken: localToken,
            remoteToken: remoteToken,
            localTokenRate: localTokenRate,
            remoteTokenRate: remoteTokenRate,
            isOrigin: isOrigin,
            paused: false,
            deposited: 0
        });

        emit TokenPairRegistered(remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate);
    }
}
