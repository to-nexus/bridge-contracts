// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

import {ICrossCheck} from "./interface/ICrossCheck.sol";
import {RoleManager} from "./abstract/RoleManager.sol";
import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {Const} from "./lib/Const.sol";
import {CrossCheckMixin} from "./lib/CrossCheckMixin.sol";
import {CrossCheckStorage} from "./abstract/CrossCheckStorage.sol";

/**
 * @title CrossCheck
 * @notice Contract responsible for validating and storing check blocks for the Cross chain block verification
 * @dev This contract extends the RoleManager and ValidatorManager contracts to provide access control and
 *      signature validation for check blocks
 */
contract CrossCheck is Initializable, UUPSUpgradeable, PausableUpgradeable, RoleManager, ValidatorManager, CrossCheckStorage, ICrossCheck {
    // ////////// errors //////////
    error CrossCheckInvalidMaxCheckBlocks(uint256 maxCheckBlocks);
    error CrossCheckInvalidChainID(uint256 chainID);
    error CrossCheckNotNextBlock(uint256 nonce, uint256 start);
    error CrossCheckBlockNotFound(uint256 blockNumber);

    // ////////// events //////////

    /**
     * @notice Emitted when the contract is initialized
     * @param chainID The ID of the Cross chain being monitored
     * @param blocksPerCheck Number of blocks covered by each check block
     * @param validatorThreshold Number of validator signatures required for data validation
     */
    event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint256 maxCheckBlocks, uint8 validatorThreshold);

    /**
     * @notice Emitted when the maximum number of check blocks is updated
     * @param maxCheckBlocks The new maximum number of check blocks
     */
    event CrossCheckMaxCheckBlocksUpdated(uint256 maxCheckBlocks);

    // ////////// constants //////////

    /**
     * @dev TypeHash used for EIP-712 signature verification
     */
    bytes32 internal constant SUBMIT_TYPEHASH = keccak256("CheckBlockArg(uint256 nonce,uint256 start,uint256 end,bytes32 rootHash,uint256 chainID)");

    /**
     * @dev Role identifiers used for access control
     */
    bytes32 public constant OPERATOR_ROLE = Const.OPERATOR_ROLE;
    bytes32 public constant VALIDATOR_ROLE = Const.VALIDATOR_ROLE;

    /**
     * @dev Chain ID of the Cross chain
     */
    uint256 private constant _crossChainID = CrossCheckMixin.crossChainID;

    /**
     * @dev Number of blocks covered by each check block
     */
    uint256 private constant _blocksPerCheck = CrossCheckMixin.blocksPerCheck;

    // ////////// storage variables //////////

    // see {CrossCheckStorage}

    /**
     * @notice Maximum number of check blocks that can be stored
     */
    uint256 public maxCheckBlocks = 100;

    /**
     * @dev storage gap
     * 3 slots at CrossCheckStorage
     * 1 slot here
     */
    uint[46] private __gap;

    /**
     * @dev Direct initialization is disabled
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Contract initializer
     * @param maxCheckBlocks_ Maximum number of check blocks to store
     * @param validatorThreshold_ Number of validator signatures required for data validation
     */
    function initialize(uint256 maxCheckBlocks_, uint8 validatorThreshold_) external initializer {
        __UUPSUpgradeable_init();
        __Pausable_init();
        __RoleManager_init(msg.sender);
        __Validator_init(validatorThreshold_);

        _setMaxCheckBlocks(maxCheckBlocks_);

        emit CrossCheckInitialized(_crossChainID, _blocksPerCheck, maxCheckBlocks_, validatorThreshold_);
    }

    /**
     * @dev see {ICrossCheck-submitCheckpoint}
     */
    function submitCheckpoint(
        bytes calldata data,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) external onlyRole(VALIDATOR_ROLE) whenNotPaused {
        address proposer = _msgSender();
        CheckBlockArg memory _block = abi.decode(data, (CheckBlockArg));

        // check chain ID
        if (_block.chainID != _crossChainID) {
            revert CrossCheckInvalidChainID(_block.chainID);
        }
        // check validity
        if (_block.end <= _block.start || _block.rootHash == bytes32(0)) {
            revert CrossCheckInvalidData();
        }

        // check next nonce and start block number
        (, uint256 _nextStart) = getNextBlockInfo();
        uint256 _nextEnd = 0;
        unchecked {
            _nextEnd = _nextStart + _blocksPerCheck - 1;
        }
        if (nextNonce != _block.nonce || _nextStart != _block.start || _nextEnd != _block.end) {
            revert CrossCheckNotNextBlock(_block.nonce, _block.start);
        }

        // check signature
        bytes32 typeHash = keccak256(abi.encode(SUBMIT_TYPEHASH, _block.nonce, _block.start, _block.end, _block.rootHash, _block.chainID));
        _validateSignature(typeHash, v, r, s);

        // add new check block
        _addCheckBlock(proposer, _block.nonce, _block.start, _block.end, _block.rootHash);

        // set the next nonce to the new one
        unchecked {
            ++nextNonce;
        }

        emit CheckBlockAdded(proposer, _block.nonce, _block.start, _block.end, _block.rootHash);

        // prune check blocks if necessary
        _pruneCheckBlocks();
    }

    /**
     * @dev see {ICrossCheck-verifyBlock}
     * @dev The merkle tree which created the proof must has its leaves sorted
     */
    function verifyBlock(uint256 blockNumber, bytes32[] calldata proof, bytes32 blockHash) external view returns (bool) {
        CheckBlock memory _block = getCheckBlockByBlockNumber(blockNumber);
        if (_block.createdAt == 0) {
            revert CrossCheckBlockNotFound(blockNumber);
        }
        return MerkleProof.verifyCalldata(proof, _block.rootHash, blockHash);
    }

    // ////////// view functions //////////

    /**
     * @dev see {EIP712Upgradeable-_domainSeparatorV4}
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    // ////////// admin functions //////////

    /**
     * @notice Updates the maximum number of check blocks to store
     * @param maxCheckBlocks_ Maximum number of check blocks to store
     */
    function setMaxCheckBlocks(uint256 maxCheckBlocks_) external onlyRole(OPERATOR_ROLE) {
        _setMaxCheckBlocks(maxCheckBlocks_);
    }

    /**
     * @dev see {PausableUpgradeable-pause}
     */
    function pause() external onlyRole(OPERATOR_ROLE) {
        _pause();
    }

    /**
     * @dev see {PausableUpgradeable-unpause}
     */
    function unpause() external onlyRole(OPERATOR_ROLE) {
        _unpause();
    }

    // ////////// internal and private functions //////////

    /**
     * @dev Prunes check blocks if necessary
     */
    function _pruneCheckBlocks() private {
        uint256 _old = firstNonce;
        while (nextNonce - firstNonce > maxCheckBlocks) {
            _removeCheckBlock(firstNonce);
            unchecked {
                ++firstNonce;
            }
        }
        if (firstNonce > _old) {
            emit CheckBlockPruned(_old, firstNonce - _old);
        }
    }

    /**
     * @dev Sets the maximum number of check blocks to store
     * @param maxCheckBlocks_ Maximum number of check blocks to store
     */
    function _setMaxCheckBlocks(uint256 maxCheckBlocks_) private {
        if (maxCheckBlocks_ < 10) {
            revert CrossCheckInvalidMaxCheckBlocks(maxCheckBlocks_);
        }
        maxCheckBlocks = maxCheckBlocks_;
        emit CrossCheckMaxCheckBlocksUpdated(maxCheckBlocks_);
    }

    /**
     * @dev see {UUPSUpgradeable-_authorizeUpgrade}
     */
    function _authorizeUpgrade(address) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
