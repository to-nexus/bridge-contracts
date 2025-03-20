// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity 0.8.28;

import {Initializable, UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

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
    error CrossCheckInvalidChainID(uint256 chainID);
    error CrossCheckNotNextBlock(uint256 nonce, uint256 start);

    // ////////// events //////////

    /**
     * @notice Emitted when the contract is initialized
     * @param chainID The ID of the Cross chain being monitored
     * @param blocksPerCheck Number of blocks covered by each check block
     * @param validatorThreshold Number of validator signatures required for data validation
     */
    event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint8 validatorThreshold);

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
     * @dev Direct initialization is disabled
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Contract initializer
     * @param validatorThreshold Number of validator signatures required for data validation
     */
    function initialize(uint8 validatorThreshold) external initializer {
        __UUPSUpgradeable_init();
        __Pausable_init();
        __RoleManager_init(msg.sender);
        __Validator_init(validatorThreshold);

        emit CrossCheckInitialized(_crossChainID, _blocksPerCheck, validatorThreshold);
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

        // check next nonce and start block number
        (uint256 nextNonce, uint256 nextStart) = getNextBlockInfo();
        uint256 nextEnd = 0;
        unchecked {
            nextEnd = nextStart + _blocksPerCheck - 1;
        }
        if (nextNonce != _block.nonce || nextStart != _block.start || nextEnd != _block.end) {
            revert CrossCheckNotNextBlock(_block.nonce, _block.start);
        }

        // check signature
        bytes32 typeHash = keccak256(abi.encode(SUBMIT_TYPEHASH, _block.nonce, _block.start, _block.end, _block.rootHash, _block.chainID));
        _validateSignature(typeHash, v, r, s);

        // add new check block
        _addCheckBlock(proposer, _block.nonce, _block.start, _block.end, _block.rootHash);

        // set the latest block number to the new one
        latestBlock = _block.start;

        emit NewCheckBlock(proposer, _block.nonce, _block.start, _block.end, _block.rootHash);
    }

    // ////////// view functions //////////

    /**
     * @dev see {EIP712Upgradeable-_domainSeparatorV4}
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    // ////////// admin functions //////////

    // TODO: add functions to prune or override existing check blocks

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
     * @dev see {UUPSUpgradeable-_authorizeUpgrade}
     */
    function _authorizeUpgrade(address) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
