// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

interface ICrossCheck {
    /**
     * @dev Input struct for checkpoint submission
     * @param nonce Sequential number for this checkpoint
     * @param start Starting block number for this checkpoint
     * @param end Ending block number for this checkpoint
     * @param rootHash Hash of the block data for the checkpoint range
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
     * @notice Submits a new checkpoint with validator signatures
     * @dev Verifies signatures, builds a check block and updates current block
     * @param data ABI-encoded CheckBlockArg struct with checkpoint details
     * @param v Signature v values array
     * @param r Signature r values array
     * @param s Signature s values array
     */
    function submitCheckpoint(bytes calldata data, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) external;
}
