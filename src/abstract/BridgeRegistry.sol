// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeRegistry} from "../interface/IBridgeRegistry.sol";

import {Const} from "../lib/Const.sol";
import {ICrossMintableERC20Code} from "../token/ICrossMintableERC20Code.sol";
import {RoleManager} from "./RoleManager.sol";
/**
 * @title BridgeRegistry
 * @notice Registry managing token pairs and chain information for the bridge
 * @dev This contract provides the following key features:
 * - Token pair registration/deregistration
 * - Chain information management
 * - Token pause/unpause functionality
 * - CrossMintable token creation
 * - Pending operation tracking
 *
 * @dev Key structures:
 * TokenPair {
 *   localToken: Token address on this chain
 *   remoteToken: Corresponding token on remote chain
 *   verificationAmountThreshold: Maximum amount that can be processed automatically without additional verification
 *   isOrigin: Whether this is the origin chain
 *   paused: Bridge operations paused flag
 *   deposited: Total amount deposited
 *   pendingAmount: Amount in pending operations
 * }
 *
 * Chain {
 *   remoteChainID: ID of the remote chain
 *   initiateIndex: Current index for initiated bridges
 *   finalizeIndex: Current index for finalized bridges
 *   paused: Whether operations for this chain are paused
 * }
 *
 * PendingData {
 *   args: FinalizeArguments struct with operation details
 *   reason: Error message explaining pending status
 *   delayExpiration: Timestamp after which pending operation can be processed
 * }
 */

abstract contract BridgeRegistry is RoleManager, IBridgeRegistry {
    using EnumerableSet for EnumerableSet.UintSet;
    using EnumerableSet for EnumerableSet.AddressSet;
    using Math for uint;

    error RegistryNotExistChain(uint remoteChainID);
    error RegistryNotExistToken(address token);
    error RegistryNotExistIndex(uint index);
    error RegistryExistERC20Code(address code);
    error RegistryExistIndex(uint index);
    error RegistryExistToken(address token);
    error RegistryZeroAddress();
    error RegistryChainPaused(uint remoteChainID);
    error RegistryTokenPaused(address token);
    error RegistryFactoryNotSet();

    /**
     * @notice Emitted when a token pair is registered
     * @param remoteChainID ID of the remote chain
     * @param localToken Address of token on this chain
     * @param remoteToken Address of corresponding token on remote chain
     * @param isOrigin Whether this is the origin chain for the token
     */
    event TokenPairRegistered(
        uint indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin
    );

    /**
     * @notice Emitted when a token pair is unregistered
     * @param remoteChainID ID of the remote chain
     * @param localToken Address of token on this chain
     */
    event TokenPairUnregistered(uint indexed remoteChainID, address indexed localToken);

    /**
     * @notice Emitted when a token's pause status is changed
     * @param remoteChainID ID of the remote chain
     * @param token Address of the token
     * @param pause New pause status
     */
    event TokenPauseSet(uint indexed remoteChainID, address indexed token, bool pause);

    /**
     * @notice Emitted when a chain is paused or unpaused
     * @param remoteChainID ID of the remote chain
     * @param pause New pause status
     */
    event ChainPauseSet(uint indexed remoteChainID, bool pause);

    /**
     * @notice Emitted when the verification delay is set
     * @param delay New verification delay value
     */
    event VerificationDelaySet(uint delay);

    /**
     * @notice Emitted when the CrossMintableERC20Factory is set
     * @param code Address of the CrossMintableERC20 Code contract
     */
    event CrossMintableERC20CodeSet(address indexed code);

    /// @dev Factory contract for creating new CrossMintable tokens
    ICrossMintableERC20Code public crossMintableERC20Code;

    /// @dev Delay period for verification (24 hours)
    uint private _verificationDelay;

    /// @dev Set of registered chain IDs
    EnumerableSet.UintSet internal _chains;

    /// @dev Mapping from chain ID to Chain struct
    mapping(uint => Chain) internal _chainData;

    /// @dev Mapping from chain ID to set of registered token addresses
    mapping(uint => EnumerableSet.AddressSet) internal _tokens;

    /// @dev Mapping from chain ID and token address to TokenPair struct
    mapping(uint => mapping(address => TokenPair)) internal _tokenPairs;

    /// @dev Mapping from chain ID to set of pending operation indices
    mapping(uint => EnumerableSet.UintSet) internal _pendingIndex;

    /// @dev Mapping from chain ID and index to pending operation data
    mapping(uint => mapping(uint => PendingData)) internal _pendingData;

    /// @dev Storage gap for future upgradessetCrossMintableERC20Code
    uint[41] private __gap;

    /**
     * @notice Initializes the BridgeRegistry
     * @dev Sets up initial state
     * - Grants Admin role to contract owner using Const.EDITOR_ROLE identifier
     * - Sets verification delay to 24 hours
     */
    function __BridgeRegistry_init() internal onlyInitializing {
        _verificationDelay = 24 hours;
    }

    /**
     * @notice Modifier to check if token is registered and not paused
     * @dev Validates token status for bridge operations
     * @param remoteChainID Chain ID to check
     * @param token Token address to check
     */
    modifier onlyValidToken(uint remoteChainID, address token) {
        _validateToken(remoteChainID, token);
        _;
    }

    /**
     * @notice Creates a new token on the local chain
     * @dev Deploys a new CrossMintable token using the factory
     * - Generates deterministic salt from remote token
     * - Creates token with specified parameters
     * - Registers token pair automatically
     * @param remoteChainID Chain ID where the original token exists
     * @param remoteToken Address of the original token
     * @param symbol Token symbol
     * @param decimals Token decimals
     * @return tokenAddress Address of the newly created token
     */
    function createToken(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        returns (address tokenAddress)
    {
        require(address(crossMintableERC20Code) != address(0), RegistryFactoryNotSet());
        tokenAddress = crossMintableERC20Code.createCrossMintableERC20(remoteChainID, remoteToken, symbol, decimals);
        registerToken(remoteChainID, false, tokenAddress, remoteToken);
    }

    /**
     * @notice Registers a token pair for bridging
     * @dev Links local and remote tokens with specified parameters
     * - Validates token rates
     * - Sets up token pair configuration
     * - Handles both origin and wrapped tokens
     * @param remoteChainID Target chain identifier
     * @param isOrigin Whether this is the origin chain for the token
     * @param localToken Local token address
     * @param remoteToken Remote token address
     */
    function registerToken(uint remoteChainID, bool isOrigin, address localToken, address remoteToken)
        public
        onlyRole(Const.EDITOR_ROLE)
    {
        if (_chains.add(remoteChainID)) {
            _chainData[remoteChainID] =
                Chain({remoteChainID: remoteChainID, initiateIndex: 0, finalizeIndex: 0, paused: false});
        }

        require(localToken != address(0), RegistryZeroAddress());
        require(_tokens[remoteChainID].add(localToken), RegistryExistToken(localToken));

        _tokenPairs[remoteChainID][localToken] = IBridgeRegistry.TokenPair({
            localToken: localToken,
            remoteToken: remoteToken,
            isOrigin: isOrigin,
            paused: false,
            deposited: 0,
            pendingAmount: 0
        });

        emit TokenPairRegistered(remoteChainID, localToken, remoteToken, isOrigin);
    }

    /**
     * @notice Unregisters a token pair
     * @dev Removes token pair from registry
     * - Validates token exists
     * - Cleans up token pair data
     * - Emits unregistration event
     * @param remoteChainID Chain ID to unregister from
     * @param token Token address to unregister
     */
    function unregisterToken(uint remoteChainID, address token) external onlyRole(Const.ADMIN_ROLE) {
        require(_tokens[remoteChainID].remove(token), RegistryNotExistToken(token));
        delete (_tokenPairs[remoteChainID][token]);
        emit TokenPairUnregistered(remoteChainID, token);
    }

    /**
     * @notice Sets the CrossMintableERC20Factory contract
     * @dev Updates factory reference for creating new tokens
     * - Can only be set once
     * - Validates non-zero address
     * - Emits factory update event
     * @param _crossMintableERC20Code Address of the new factory
     */
    function setCrossMintableERC20Code(ICrossMintableERC20Code _crossMintableERC20Code)
        external
        onlyRole(Const.ADMIN_ROLE)
    {
        require(address(crossMintableERC20Code) == address(0), RegistryExistERC20Code(address(crossMintableERC20Code)));
        require(address(_crossMintableERC20Code) != address(0), RegistryZeroAddress());

        crossMintableERC20Code = _crossMintableERC20Code;
        emit CrossMintableERC20CodeSet(address(crossMintableERC20Code));
    }

    /**
     * @notice Pauses bridge operations for a specific chain
     * @dev Temporarily disables all bridge operations for the chain
     * - Validates chain exists and is not already paused
     * - Sets pause flag
     * - Emits pause event
     * @param remoteChainID Chain ID to pause
     * @param pause Whether to pause (true) or unpause (false)
     */
    function setChainPause(uint remoteChainID, bool pause) external onlyRole(Const.OPERATOR_ROLE) {
        require(_chains.contains(remoteChainID), RegistryNotExistChain(remoteChainID));
        _chainData[remoteChainID].paused = pause;
        emit ChainPauseSet(remoteChainID, pause);
    }

    /**
     * @notice Pauses bridging operations for a token pair
     * @dev Temporarily disables bridge operations
     * - Validates token exists and is not already paused
     * - Sets pause flag
     * - Emits pause event
     * @param remoteChainID Chain ID of the token pair
     * @param token Token address to pause
     * @param pause Whether to pause (true) or unpause (false)
     */
    function setTokenPause(uint remoteChainID, address token, bool pause) external onlyRole(Const.OPERATOR_ROLE) {
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        _tokenPairs[remoteChainID][token].paused = pause;
        emit TokenPauseSet(remoteChainID, token, pause);
    }

    /**
     * @notice Sets the verification delay
     * @dev Updates the delay period for verification
     * @param delay New delay value in seconds
     */
    function setVerificationDelay(uint delay) external onlyRole(Const.ADMIN_ROLE) {
        _verificationDelay = delay;
        emit VerificationDelaySet(delay);
    }

    /**
     * @notice Returns all registered chain IDs
     * @dev Retrieves array of chain IDs from storage
     * @return Array of chain IDs
     */
    function allChainIDs() external view returns (uint[] memory) {
        return _chains.values();
    }

    /**
     * @notice Returns all token pairs for a chain
     * @dev Retrieves and formats token pair data
     * - Gets all registered tokens for chain
     * - Maps tokens to their pair data
     * @param remoteChainID Chain ID to query
     * @return Array of TokenPair structs
     */
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory) {
        address[] memory tokens = _tokens[remoteChainID].values();
        TokenPair[] memory _pairs = new TokenPair[](tokens.length);
        for (uint i = 0; i < tokens.length; ++i) {
            _pairs[i] = _tokenPairs[remoteChainID][tokens[i]];
        }
        return _pairs;
    }

    /**
     * @notice Returns all pending operation indices
     * @dev Lists indices of incomplete bridge operations
     * @param remoteChainID Chain ID to query
     * @return Array of pending indices
     */
    function allPendingIndex(uint remoteChainID) external view returns (uint[] memory) {
        return _pendingIndex[remoteChainID].values();
    }

    /**
     * @notice Returns token pair information
     * @dev Retrieves specific token pair configuration
     * @param remoteChainID Chain ID of the pair
     * @param token Token address to query
     * @return TokenPair struct containing pair details
     */
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory) {
        return _tokenPairs[remoteChainID][token];
    }

    /**
     * @notice Gets next initiate index for a chain
     * @dev Returns incremented current initiate index
     * @param remoteChainID Chain ID to query
     * @return Next available index for bridge initiation
     */
    function getNextInitiateIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].initiateIndex + 1;
    }

    /**
     * @notice Gets next finalize index for a chain
     * @dev Returns incremented current finalize index
     * @param remoteChainID Chain ID to query
     * @return Next available index for bridge finalization
     */
    function getNextFinalizeIndex(uint remoteChainID) public view returns (uint) {
        return _chainData[remoteChainID].finalizeIndex + 1;
    }

    /**
     * @notice Returns pending bridge operation details
     * @dev Retrieves stored finalization arguments
     * @param remoteChainID Chain ID of pending operation
     * @param index Index of pending operation
     * @return PendingData struct with operation details
     */
    function getPendingArguments(uint remoteChainID, uint index) external view returns (PendingData memory) {
        return _pendingData[remoteChainID][index];
    }

    function isPending(uint remoteChainID, uint index) external view returns (bool) {
        return _pendingIndex[remoteChainID].contains(index);
    }

    /**
     * @notice Validates token registration and status
     * @dev Internal function to check token can be used for bridging
     * - Verifies token is registered
     * - Verifies chain is not paused
     * - Verifies token is not paused
     * @param remoteChainID Chain ID to check
     * @param token Token address to check
     */
    function _validateToken(uint remoteChainID, address token) private view {
        // check token is registered
        require(_tokens[remoteChainID].contains(token), RegistryNotExistToken(token));
        // check chain is not paused
        require(!_chainData[remoteChainID].paused, RegistryChainPaused(remoteChainID));
        // check token is not paused
        require(!_tokenPairs[remoteChainID][token].paused, RegistryTokenPaused(token));
    }

    /**
     * @notice Increments initiate index for a chain
     * @dev Internal function to update counter
     * @param remoteChainID Chain ID to update
     * @return index Current index value (after increment)
     */
    function _useInitiateIndex(uint remoteChainID) internal returns (uint index) {
        unchecked {
            index = ++_chainData[remoteChainID].initiateIndex;
        }
    }

    /**
     * @notice Increments finalize index for a chain
     * @dev Internal function to update counter
     * @param remoteChainID Chain ID to update
     * @return index Current index value (after increment)
     */
    function _useFinalizeIndex(uint remoteChainID) internal returns (uint index) {
        unchecked {
            index = ++_chainData[remoteChainID].finalizeIndex;
        }
    }

    /**
     * @notice Records token deposit
     * @dev Updates deposited amount for token pair
     * @param remoteChainID Chain ID of token pair
     * @param token Token address
     * @param value Amount deposited
     */
    function _depositToken(uint remoteChainID, address token, uint value) internal virtual {
        _tokenPairs[remoteChainID][token].deposited += value;
    }

    /**
     * @notice Processes token withdrawal
     * @dev Updates deposited amount and validates balance
     * - Checks sufficient balance including pending amounts
     * - Reduces deposited amount
     * @param remoteChainID Chain ID of token pair
     * @param token Token address
     * @param value Amount to withdraw
     */
    function _withdrawToken(uint remoteChainID, address token, uint value) internal virtual {
        _tokenPairs[remoteChainID][token].deposited -= value;
    }

    /**
     * @notice Records pending bridge operation
     * @dev Stores operation details and updates state
     * - Adds to pending index set
     * - Stores operation arguments and reason
     * - Updates pending amounts for origin tokens
     * - Updates verification delay expiration for delayed processing
     * @param args Finalization arguments to store
     * @param status Status of the pending operation
     * @param reason Error message explaining pending status
     * @param delay Whether to delay processing (verification delay)
     */
    function _setPendingArguments(
        FinalizeArguments calldata args,
        Const.FinalizeStatus status,
        bytes memory reason,
        bool delay
    ) internal {
        require(_pendingIndex[args.fromChainID].add(args.index), RegistryExistIndex(args.index));

        _pendingData[args.fromChainID][args.index] = PendingData({
            args: args,
            status: status,
            reason: reason,
            delayExpiration: delay ? block.timestamp + _verificationDelay : 0
        });

        TokenPair storage tokenPair = _tokenPairs[args.fromChainID][address(args.toToken)];
        tokenPair.pendingAmount += args.value;
    }

    /**
     * @notice Removes pending bridge operation
     * @dev Cleans up pending operation data
     * - Removes from pending index set
     * - Clears stored arguments and reason
     * - Updates pending amounts for origin tokens
     * @param remoteChainID Chain ID of operation
     * @param index Index of operation to remove
     * @return args The finalization arguments of the removed pending operation
     */
    function _removePendingArguments(uint remoteChainID, uint index) internal returns (FinalizeArguments memory args) {
        require(_pendingIndex[remoteChainID].remove(index), RegistryNotExistIndex(index));

        args = _pendingData[remoteChainID][index].args;
        TokenPair storage tokenPair = _tokenPairs[remoteChainID][address(args.toToken)];

        tokenPair.pendingAmount -= args.value;
        delete (_pendingData[remoteChainID][index]);
    }
}
