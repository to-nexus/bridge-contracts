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
     * @notice The starting block number of the latest check block
     */
    uint256 public latestBlock;

    /**
     * @notice The number of submitted check blocks
     */
    uint256 public numCheckBlocks;

    /**
     * @notice Mapping from start block number to corresponding check block
     * @dev Key is the starting block number of the check range
     */
    mapping(uint256 => CheckBlock) internal _checkBlocks;

    /**
     * @dev storage gap
     */
    uint[47] private __gap;

    /**
     * @notice Retrieves check block for a specific block number
     * @param blockNumber The starting block number to query
     * @return nonce Sequential identifier of this check block
     * @return start The starting block number this check block covers
     * @return end The ending block number this check block covers
     * @return createdAt The block timestamp when this check block was created
     * @return rootHash The hash representing the root of the block data
     * @return proposer The address that proposed this check block
     */
    function getCheckBlock(
        uint256 blockNumber
    ) external view returns (uint256 nonce, uint256 start, uint256 end, uint256 createdAt, bytes32 rootHash, address proposer) {
        CheckBlock storage _block = _checkBlocks[blockNumber];
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
        unchecked {
            ++numCheckBlocks;
        }
    }
}
