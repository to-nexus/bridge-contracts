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
        address indexed localToken,
        address indexed remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        uint safetyLimit,
        bool isOrigin
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
        require(_tokenContains(remoteChainID, token), RegistryNotExistToken(token));
        require(!_tokenPairs[remoteChainID][token].paused, RegistryTokenPaused(token));
        _;
    }

    function _tokenContains(uint remoteChainID, address token) internal view returns (bool) {
        return _tokens[remoteChainID].contains(token);
    }

    /**
     * @notice Creates a new token on the local chain.
     * @param remoteChainID The ID of the remote chain.
     * @param remoteToken The address of the token on the remote chain.
     * @param localTokenRate The rate of the local token.
     * @param remoteTokenRate The rate of the remote token.
     * @param symbol The symbol of the token.
     * @param decimals The decimals of the token.
     * @return tokenAddress The address of the created token.
     */
    function createToken(
        uint remoteChainID,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        uint safetyLimit,
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
        _registerToken(remoteChainID, false, tokenAddress, remoteToken, localTokenRate, remoteTokenRate, safetyLimit);
    }

    /**
     * @notice Registers a token pair.
     * @param remoteChainID The ID of the remote chain.
     * @param isOrigin Whether the token is the origin token.
     * @param localToken The address of the local token.
     * @param remoteToken The address of the remote token.
     * @param localTokenRate The rate of the local token.
     * @param remoteTokenRate The rate of the remote token.
     */
    function registerToken(
        uint remoteChainID,
        bool isOrigin,
        address localToken,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        uint safetyLimit
    ) external onlyOwner {
        _registerToken(remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate, safetyLimit);
    }

    /**
     * @notice Unregisters a token pair.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to unregister.
     */
    function unregisterToken(uint remoteChainID, address token) external onlyOwner {
        require(_tokens[remoteChainID].remove(token), RegistryNotExistToken(token));
        delete (_tokenPairs[remoteChainID][token]);

        emit TokenPairUnregistered(remoteChainID, token);
    }

    /**
     * @notice Pauses a token pair.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to pause.
     */
    function pauseToken(uint remoteChainID, address token) external onlyOwner {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        require(!_tokenPairs[remoteChainID][token].paused, RegistryTokenPaused(token));
        _tokenPairs[remoteChainID][token].paused = true;

        emit TokenPairPaused(remoteChainID, token);
    }

    /**
     * @notice Unpauses a token pair.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to unpause.
     */
    function unpauseToken(uint remoteChainID, address token) external onlyOwner {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        require(_tokenPairs[remoteChainID][token].paused, RegistryNotPaused(token));
        _tokenPairs[remoteChainID][token].paused = false;

        emit TokenPairUnpaused(remoteChainID, token);
    }

    /**
     * @notice Sets the CrossMintableERC20Factory.
     * @param _crossMintableERC20Factory The address of the CrossMintableERC20Factory.
     */
    function setCrossMintableERC20Factory(ICrossMintableERC20Factory _crossMintableERC20Factory) external onlyOwner {
        require(
            address(crossMintableERC20Factory) == address(0), RegistryExistFactory(address(crossMintableERC20Factory))
        );
        require(address(_crossMintableERC20Factory) != address(0), RegistryZeroAddress());

        crossMintableERC20Factory = _crossMintableERC20Factory;
        emit CrossMintableERC20FactorySet(address(crossMintableERC20Factory));
    }

    function clearPending(uint remoteChainIDs) external onlyOwner {
        uint[] memory indexes = _chainData[remoteChainIDs].pending.index.values();
        for (uint i = 0; i < indexes.length; ++i) {
            _removePendingArguments(remoteChainIDs, indexes[i]);
        }
    }

    function setSafetyLimit(uint remoteChainID, address token, uint safetyLimit) external onlyOwner {
        require(_tokenContains(remoteChainID, token), RegistryNotExistToken(token));
        _tokenPairs[remoteChainID][token].safetyLimit = safetyLimit;
    }

    /**
     * @notice Returns all chain IDs.
     * @return An array of all chain IDs.
     */
    function allChainIDs() external view returns (uint[] memory) {
        return _chains.values();
    }

    /**
     * @notice Returns all token pairs for a given remote chain ID.
     * @param remoteChainID The ID of the remote chain.
     * @return An array of all token pairs.
     */
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory) {
        address[] memory tokens = _tokens[remoteChainID].values();
        TokenPair[] memory _pairs = new TokenPair[](tokens.length);
        for (uint i = 0; i < tokens.length; ++i) {
            _pairs[i] = _tokenPairs[remoteChainID][tokens[i]];
        }
        return _pairs;
    }

    /**
     * @notice Returns all pending indexes for a given remote chain ID.
     * @param remoteChainID The ID of the remote chain.
     * @return An array of all pending indexes.
     */
    function allPendingIndex(uint remoteChainID) external view returns (uint[] memory) {
        return _chainData[remoteChainID].pending.index.values();
    }

    /**
     * @notice Returns the token pair for a given remote chain ID and token address.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token.
     * @return The token pair.
     */
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory) {
        return _tokenPairs[remoteChainID][token];
    }

    /**
     * @notice Returns the next initiate index for a given remote chain ID.
     * @param remoteChainID The ID of the remote chain.
     * @return The next initiate index.
     */
    function getNextInitiateIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].initiateIndex + 1;
    }

    /**
     * @notice Returns the next finalize index for a given remote chain ID.
     * @param remoteChainID The ID of the remote chain.
     * @return The next finalize index.
     */
    function getNextFinalizeIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].finalizeIndex + 1;
    }

    /**
     * @notice Returns the pending arguments for a given remote chain ID and index.
     * @param remoteChainID The ID of the remote chain.
     * @param index The index of the pending arguments.
     * @return The pending arguments.
     */
    function pendingArguments(uint remoteChainID, uint index) public view returns (FinalizeArguments memory) {
        return _chainData[remoteChainID].pending.data[index];
    }

    /**
     * @notice Returns the pending reason for a given remote chain ID and index.
     * @param remoteChainID The ID of the remote chain.
     * @param index The index of the pending reason.
     * @return The pending reason.
     */
    function pendingReason(uint remoteChainID, uint index) public view returns (string memory) {
        return _chainData[remoteChainID].pending.reason[index];
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
        TokenPair storage tokenPair = _tokenPairs[remoteChainID][token];
        require(tokenPair.deposited >= value + tokenPair.pendingAmount, RegistryBalanceLow(remoteChainID, token, value));
        tokenPair.deposited -= value;
    }

    function _setPendingArguments(FinalizeArguments calldata args, string memory reason) internal {
        Chain storage chain = _chainData[args.remoteChainID];
        uint index = args.index;
        require(chain.pending.index.add(index), RegistryExistIndex(index));

        chain.pending.data[index] = args;
        chain.pending.reason[index] = reason;
        TokenPair storage tokenPair = _tokenPairs[args.remoteChainID][address(args.token)];
        if (tokenPair.isOrigin) tokenPair.pendingAmount += args.value;
    }

    function _removePendingArguments(uint remoteChainID, uint index) internal {
        PendingData storage pending = _chainData[remoteChainID].pending;
        require(pending.index.remove(index), RegistryNotExistIndex(index));

        delete (pending.reason[index]);
        delete (pending.data[index]);
    }

    // private functions
    function _registerToken(
        uint remoteChainID,
        bool isOrigin,
        address localToken,
        address remoteToken,
        uint localTokenRate,
        uint remoteTokenRate,
        uint safetyLimit
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
            safetyLimit: safetyLimit,
            isOrigin: isOrigin,
            paused: false,
            deposited: 0,
            pendingAmount: 0
        });

        emit TokenPairRegistered(
            remoteChainID, localToken, remoteToken, localTokenRate, remoteTokenRate, safetyLimit, isOrigin
        );
    }
}
