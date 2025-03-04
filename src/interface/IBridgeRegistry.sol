// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IBridgeRegistry {
    struct FinalizeArguments {
        uint remoteChainID;
        uint index;
        IERC20 token;
        address to;
        uint value;
        bytes extraData;
    }

    struct Chain {
        uint remoteChainID;
        uint initiateIndex;
        uint finalizeIndex;
        Reverted reverted;
    }

    struct TokenPair {
        address localToken;
        address remoteToken;
        uint localTokenRate;
        uint remoteTokenRate;
        bool isOrigin;
        bool paused;
        uint deposited;
    }

    struct Reverted {
        mapping(uint => FinalizeArguments) data; // index : reverted arguments
        mapping(uint => string) reason; // index : reverted reason
        EnumerableSet.UintSet index;
    }

    function allChainIDs() external view returns (uint[] memory);
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory);
    function allRevertedIndex(uint remoteChainID) external view returns (uint[] memory);
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory);
    function getNextInitiateIndex(uint remoteChainID) external view returns (uint);
    function getNextFinalizeIndex(uint remoteChainID) external view returns (uint);
    function revertedArguments(uint remoteChainID, uint index) external view returns (FinalizeArguments memory);
    function revertedReason(uint remoteChainID, uint index) external view returns (string memory);
}
