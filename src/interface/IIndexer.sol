// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface IIndexer {
    function nextInitiateIndex() external view returns (uint);
    function nextFinalizeIndex() external view returns (uint);
}
