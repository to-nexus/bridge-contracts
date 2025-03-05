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
        PendingData pending;
    }

    struct TokenPair {
        address localToken; // local token address
        address remoteToken; // remote token address
        uint localTokenRate; // rate of the local token
        uint remoteTokenRate; // rate of the remote token
        uint safetyLimit; // safety limit of the token
        bool paused; // whether the token is paused
        bool isOrigin; // whether the token is origin token
        uint deposited; // deposited amount of the token
        uint pendingAmount; // pending amount of the token
    }

    struct PendingData {
        EnumerableSet.UintSet index;
        mapping(uint => FinalizeArguments) data; // index : arguments
        mapping(uint => string) reason; // index : reason
    }

    function allChainIDs() external view returns (uint[] memory);
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory);
    function allPendingIndex(uint remoteChainID) external view returns (uint[] memory);
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory);
    function getNextInitiateIndex(uint remoteChainID) external view returns (uint);
    function getNextFinalizeIndex(uint remoteChainID) external view returns (uint);
    function pendingArguments(uint remoteChainID, uint index) external view returns (FinalizeArguments memory);
    function pendingReason(uint remoteChainID, uint index) external view returns (string memory);
    function clearPending(uint remoteChainID) external;
}
