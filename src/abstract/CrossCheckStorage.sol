// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

/**
 * @title CrossCheckBlock
 * @notice Contract that defines the structure and events for the Cross chain verification checkpoint
 */
abstract contract CrossCheckBlock {
    /**
     * @notice Emitted when a new check block is created
     * @param proposer The address that proposed this check block
     * @param nonce Sequential identifier for the check block
     * @param start The starting block number this check covers
     * @param end The ending block number this check covers
     * @param rootHash The hash representing the root of the block data
     */
    event CheckBlockAdded(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash);

    /**
     * @notice Emitted when a check block is pruned
     * @param nonce The nonce of the first check block that is pruned
     * @param count The number of check blocks pruned
     */
    event CheckBlockPruned(uint256 indexed nonce, uint256 count);

    /**
     * @notice Structure representing a Cross chain verification checkpoint
     * @dev Optimized for gas usage using uint64 for numeric values (3 slots)
     */
    struct CheckBlock {
        uint64 nonce;
        uint64 start;
        uint64 end;
        uint64 createdAt;
        bytes32 rootHash;
        address proposer;
    }
}

/**
 * @title CrossCheckStorage
 * @notice Contract for managing information related to CrossCheckBlock
 */
abstract contract CrossCheckStorage is CrossCheckBlock {
    error CrossCheckInvalidData();
    error CrossCheckBlockExists(uint256 nonce);
    error CrossCheckNonceNotFound(uint256 nonce);

    /**
     * @notice Mapping from check block nonce to corresponding check block
     * @dev Key is the nonce of the check block
     */
    mapping(uint256 => CheckBlock) internal _checkBlocks;

    /**
     * @notice Nonce of the first check block
     */
    uint256 public firstNonce;

    /**
     * @notice Nonce of the last check block
     */
    uint256 public lastNonce;

    /**
     * @notice Maximum number of check blocks that can be stored
     */
    uint256 public maxCheckBlocks = 100;

    /**
     * @dev storage gap
     */
    uint[46] private __gap;

    /**
     * @notice Returns the number of check blocks
     * @return The number of check blocks
     */
    function numCheckBlocks() external view returns (uint256) {
        if (lastNonce == firstNonce) {
            if (_checkBlocks[lastNonce].createdAt == 0) {
                return 0;
            }
        }
        return lastNonce - firstNonce + 1;
    }

    /**
     * @notice Retrieves check block of a specific nonce
     * @param nonce The nonce of the check block to query
     * @return block The check block
     */
    function getCheckBlock(uint256 nonce) external view returns (CheckBlock memory) {
        return _checkBlocks[nonce];
    }

    /**
     * @notice Retrieves check block of a specific block number
     * @param blockNumber The block number to query
     * @return block The check block
     */
    function getCheckBlockByBlockNumber(uint256 blockNumber) public view returns (CheckBlock memory) {
        // short cut
        if (blockNumber == 0) {
            return _checkBlocks[0];
        }

        // binary search using nonce
        uint256 low = firstNonce;
        uint256 high = lastNonce;
        while (low <= high) {
            uint256 mid = (low + high) / 2;
            CheckBlock storage _block = _checkBlocks[mid];
            if (_block.start <= blockNumber && blockNumber <= _block.end) {
                return _block;
            } else if (blockNumber < _block.start) {
                unchecked {
                    high = mid - 1;
                }
            } else {
                unchecked {
                    low = mid + 1;
                }
            }
        }
        return CheckBlock(0, 0, 0, 0, bytes32(0), address(0));
    }

    /**
     * @notice Calculates the next block number to be processed
     * @dev Returns the next nonce and starting block number
     */
    function getNextBlockInfo() public view returns (uint256 nextNonce, uint256 nextStart) {
        CheckBlock storage _block = _checkBlocks[lastNonce];
        if (_block.createdAt > 0) {
            unchecked {
                nextNonce = lastNonce + 1;
                nextStart = uint256(_block.end) + 1;
            }
        }
    }

    /**
     * @dev Validates and adds a new check block
     * This function does not check if the check block is the last one
     * The caller is responsible to check it
     * @param proposer Address of the check block proposer
     * @param nonce Sequential number of this check block
     * @param start Starting block number for this check block
     * @param end Ending block number for this check block
     * @param rootHash Hash of the block data in the check range
     */
    function _addCheckBlock(address proposer, uint256 nonce, uint256 start, uint256 end, bytes32 rootHash) internal virtual {
        // check validity
        if (end <= start || rootHash == bytes32(0)) {
            revert CrossCheckInvalidData();
        }
        // check boundary
        if (nonce > type(uint64).max || end > type(uint64).max) {
            revert CrossCheckInvalidData();
        }
        // check if the block is already added
        if (_checkBlocks[nonce].createdAt > 0) {
            revert CrossCheckBlockExists(nonce);
        }

        CheckBlock memory nextBlock = CheckBlock({
            nonce: uint64(nonce),
            start: uint64(start),
            end: uint64(end),
            createdAt: uint64(block.timestamp),
            rootHash: rootHash,
            proposer: proposer
        });

        // insert new check block
        _checkBlocks[nonce] = nextBlock;
    }

    /**
     * @dev Removes a check block from the storage
     * This function does not check if the check block is the first one
     * The caller is responsible to check it
     * @param nonce The nonce of the check block to remove
     */
    function _removeCheckBlock(uint256 nonce) internal virtual {
        if (_checkBlocks[nonce].createdAt == 0) {
            revert CrossCheckNonceNotFound(nonce);
        }
        delete _checkBlocks[nonce];
    }
}
