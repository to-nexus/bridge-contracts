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
    event NewCheckBlock(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash);

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
    error CrossCheckBlockExists(uint256 start);

    /**
     * @notice Mapping from start block number to corresponding check block
     * @dev Key is the starting block number of the check range
     */
    mapping(uint256 => CheckBlock) internal _checkBlocks;

    /**
     * @notice Mapping from nonce to corresponding check block's start block number
     * @dev Key is the nonce of the check block
     */
    mapping(uint256 => uint256) internal _nonceToCheckBlocks;

    /**
     * @notice The starting block number of the latest check block
     */
    uint256 public latestBlock;

    /**
     * @notice The number of submitted check blocks
     */
    uint256 public numCheckBlocks;

    /**
     * @dev storage gap
     */
    uint[46] private __gap;

    /**
     * @notice Retrieves check block of a specific starting block number
     * @param startBlockNumber The starting block number to query
     * @return nonce Sequential identifier of this check block
     * @return start The starting block number this check block covers
     * @return end The ending block number this check block covers
     * @return createdAt The block timestamp when this check block was created
     * @return rootHash The hash representing the root of the block data
     * @return proposer The address that proposed this check block
     */
    function getCheckBlock(
        uint256 startBlockNumber
    ) public view returns (uint256 nonce, uint256 start, uint256 end, uint256 createdAt, bytes32 rootHash, address proposer) {
        CheckBlock storage _block = _checkBlocks[startBlockNumber];
        (nonce, start, end, createdAt, rootHash, proposer) = (
            uint256(_block.nonce),
            uint256(_block.start),
            uint256(_block.end),
            uint256(_block.createdAt),
            _block.rootHash,
            _block.proposer
        );
    }

    /**
     * @notice Retrieves check block of a specific nonce
     * @param nonce The nonce of the check block to query
     * @return nonce Sequential identifier of this check block
     * @return start The starting block number this check block covers
     * @return end The ending block number this check block covers
     * @return createdAt The block timestamp when this check block was created
     * @return rootHash The hash representing the root of the block data
     * @return proposer The address that proposed this check block
     */
    function getCheckBlockByNonce(uint256 nonce) external view returns (uint256, uint256, uint256, uint256, bytes32, address) {
        // short cut
        if (nonce == 0) {
            return getCheckBlock(0);
        }
        uint256 start = _nonceToCheckBlocks[nonce];
        if (start == 0) {
            // nonce is not found
            return (0, 0, 0, 0, bytes32(0), address(0));
        }
        return getCheckBlock(_nonceToCheckBlocks[nonce]);
    }

    /**
     * @notice Retrieves check block of a specific block number
     * @param blockNumber The block number to query
     * @return nonce Sequential identifier of this check block
     * @return start The starting block number this check block covers
     * @return end The ending block number this check block covers
     * @return createdAt The block timestamp when this check block was created
     * @return rootHash The hash representing the root of the block data
     * @return proposer The address that proposed this check block
     */
    function getCheckBlockByBlockNumber(uint256 blockNumber) external view returns (uint256, uint256, uint256, uint256, bytes32, address) {
        // short cut
        if (blockNumber == 0) {
            return getCheckBlock(0);
        }
        if (_checkBlocks[blockNumber].createdAt > 0) {
            return getCheckBlock(blockNumber);
        }

        // binary search using nonce
        uint256 low = 0;
        uint256 high = numCheckBlocks - 1;
        while (low <= high) {
            uint256 mid = (low + high) / 2;
            CheckBlock storage _block = _checkBlocks[_nonceToCheckBlocks[mid]];
            if (_block.start <= blockNumber && blockNumber <= _block.end) {
                return getCheckBlock(_block.start);
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
        return (0, 0, 0, 0, bytes32(0), address(0));
    }

    /**
     * @notice Calculates the next block number to be processed
     * @dev Returns the next nonce and starting block number
     */
    function getNextBlockInfo() public view returns (uint256 nextNonce, uint256 nextStart) {
        CheckBlock storage _block = _checkBlocks[latestBlock];
        if (_block.createdAt > 0) {
            unchecked {
                nextNonce = uint256(_block.nonce) + 1;
                nextStart = uint256(_block.end) + 1;
            }
        }
    }

    /**
     * @dev Validates and adds a new check block
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
        if (nonce > type(uint64).max || start > type(uint64).max || end > type(uint64).max) {
            revert CrossCheckInvalidData();
        }
        // check if the block is already added
        if (_checkBlocks[start].createdAt > 0) {
            revert CrossCheckBlockExists(start);
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
        _checkBlocks[start] = nextBlock;
        _nonceToCheckBlocks[nonce] = start;
        unchecked {
            ++numCheckBlocks;
        }
    }
}
