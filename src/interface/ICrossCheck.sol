// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

interface ICrossCheck {
    /**
     * @notice Input struct for check block submission
     * @param nonce Sequential number for this check block
     * @param start Starting block number for this check block
     * @param end Ending block number for this check block
     * @param rootHash Hash of the block data for the check block range
     * @param chainID ID of the Cross chain being monitored
     */
    struct CheckBlockArg {
        uint256 nonce;
        uint256 start;
        uint256 end;
        bytes32 rootHash;
        uint256 chainID;
    }

    /**
     * @notice Submits a new check block with validator signatures
     * @dev Verifies signatures, builds a check block and inserts it
     * @param data ABI-encoded CheckBlockArg struct with check block details
     * @param v Signature v values array
     * @param r Signature r values array
     * @param s Signature s values array
     */
    function submitCheckpoint(bytes calldata data, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) external;

    /**
     * @notice Checks if the given blockHash can be proved to be part of a check block
     * @param blockNumber The starting block number to query the check block
     * @param proof The merkle proof
     * @param blockHash The hash of the block to be verified
     * @return isProved True if it is successfully proved
     */
    function verifyBlock(uint256 blockNumber, bytes32[] calldata proof, bytes32 blockHash) external view returns (bool);
}
