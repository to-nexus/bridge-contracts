// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IChainManager {
    struct FinalizeArguments {
        uint index;
        IERC20 token;
        address to;
        uint value;
        bytes extraData;
    }

    struct Reverted {
        mapping(uint => FinalizeArguments) data; // index : reverted arguments
        mapping(uint => string) reason; // index : reverted reason
        EnumerableSet.UintSet index;
    }

    struct Index {
        uint initiate;
        uint finalize;
    }

    struct TokenPair {
        address localToken;
        address remoteToken;
        bool isOrigin;
        bool paused;
        uint deposited;
    }

    struct Chain {
        uint remoteChainID;
        Index index;
        Reverted reverted;
        EnumerableSet.AddressSet tokens;
        mapping(address => TokenPair) tokenPairs; // localToken : TokenPair
    }

    function getNextInitiateIndex(uint remoteChainID) external view returns (uint);
    function getNextFinalizeIndex(uint remoteChainID) external view returns (uint);
    function revertedArguments(uint remoteChainID, uint index) external view returns (FinalizeArguments memory);
    function revertedReason(uint remoteChainID, uint index) external view returns (string memory);
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory);
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory);
}
