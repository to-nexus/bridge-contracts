// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IRoleManager} from "./IRoleManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IBridgeRegistry is IRoleManager {
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
        bool paused;
    }

    struct TokenPair {
        address localToken; // local token address
        address remoteToken; // remote token address
        bool paused; // whether the token is paused
        bool isOrigin; // whether the token is origin token
        uint safetyLimit; // safety limit of the token
        uint deposited; // deposited amount of the token
        uint pendingAmount; // pending amount of the token
    }

    struct PendingData {
        FinalizeArguments args;
        uint safeDeadline;
        bytes reason;
    }

    function allChainIDs() external view returns (uint[] memory);
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory);
    function allPendingIndex(uint remoteChainID) external view returns (uint[] memory);
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory);
    function getNextInitiateIndex(uint remoteChainID) external view returns (uint);
    function getNextFinalizeIndex(uint remoteChainID) external view returns (uint);
    function getPendingArguments(uint remoteChainID, uint index) external view returns (PendingData memory);
    function hasExpiredPending() external view returns (bool);
    function clearPending(uint remoteChainID) external;
}
