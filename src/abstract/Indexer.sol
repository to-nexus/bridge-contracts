// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IIndexer} from "../interface/IIndexer.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

abstract contract Indexer is IIndexer {
    uint internal _initiateIndex;
    uint internal _finalizeIndex;

    function nextInitiateIndex() public view returns (uint) {
        return _initiateIndex + 1;
    }

    function nextFinalizeIndex() public view returns (uint) {
        return _finalizeIndex + 1;
    }

    function _incrementInitiateIndex() internal {
        unchecked {
            ++_initiateIndex;
        }
    }

    function _incrementFinalizeIndex() internal {
        unchecked {
            ++_finalizeIndex;
        }
    }

    uint[48] private __gap;
}
