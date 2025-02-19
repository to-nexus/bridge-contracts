// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeStandard} from "../interface/IBridgeStandard.sol";
import {IBridgeTokenInfo} from "../interface/IBridgeTokenInfo.sol";
import {Indexer} from "./Indexer.sol";
import {TokenManager} from "./TokenManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";

/**
 * @title BridgeStandard
 * @notice This abstract contract provides a standard implementation for cross-chain bridging functionality.
 * It inherits from several OpenZeppelin upgradeable contracts for security and functionality.
 * It also utilizes Indexer, TokenManager, and ValidatorManager for managing bridge requests and validators.
 */
abstract contract BridgeStandard is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    TokenManager,
    Indexer,
    ValidatorManager,
    IBridgeStandard
{
    using EnumerableSet for EnumerableSet.UintSet;

    error BridgeStandardInvalidIndex(uint expected, uint actual);
    error BridgeStandardInvalidAmount(
        uint minimumValue, uint expectedGas, uint expectedService, uint actualValue, uint actualGas, uint actualService
    );
    error BridgeStandardInvalidSignatures(uint index);
    error BridgeStandardInvalidPermitValue(address token, uint value);
    error BridgeStandardDuplicateIndex(uint index);
    error BridgeStandardNotExistingIndex(uint index);
    error BridgeStandardCanNotZeroAddress(string name);
    error BridgeStandardFailedPermit(address token);

    event BridgeInitiated(
        uint indexed index,
        IERC20 indexed token,
        IERC20 pairToken,
        address indexed from,
        address to,
        uint value,
        uint time,
        bool permit,
        bytes[] extraData
    );
    event BridgeFinalized(uint indexed index, IERC20 indexed token, address indexed to, uint value, uint time);
    event BridgeFinalizeReverted(uint indexed index);
    event BridgeFeeCharged(
        uint indexed index, IERC20 indexed token, address indexed account, uint gasFee, uint serviceFee
    );

    uint internal constant _exrate = 100; // cross : xcross, 1 : 100
    bytes32 private constant FINALIZE_TYPEHASH =
        keccak256("Finalize(uint256 index,address token,address to,uint256 value,bytes[] extraData)");

    IBridgeTokenInfo public bridgeTokenInfo;
    uint private _initializedAt;
    address payable private _rewardWallet;

    mapping(uint => FinalizeArguments) private _revertedArguments;
    mapping(uint => bytes) private _revertedReason;
    EnumerableSet.UintSet private _revertedIndex;

    uint[44] private __gap;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract.
     * @param rewardWallet_ The address of the reward wallet.
     * @param _bridgeTokenInfo The address of the BridgeTokenInfo contract.
     */
    function __BridgeStandard_init(address rewardWallet_, address _bridgeTokenInfo) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __Validator_init();

        require(address(rewardWallet_) != address(0), BridgeStandardCanNotZeroAddress("rewardWallet"));
        // require(address(_bridgeTokenInfo) != address(0), CanNotZeroAddress("_bridgeTokenInfo")); // allow zero address

        _initializedAt = block.number;
        bridgeTokenInfo = IBridgeTokenInfo(_bridgeTokenInfo);
        _rewardWallet = payable(rewardWallet_);
    }

    /**
     * @notice Initiates a bridge transaction.
     * @param token The address of the token to bridge.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge transaction was initiated successfully.
     */
    function bridge(IERC20 token, uint value, uint gas, uint service, bytes[] calldata extraData)
        public
        payable
        returns (bool)
    {
        return bridgeTo(token, _msgSender(), value, gas, service, extraData);
    }

    /**
     * @notice Initiates a bridge transaction to a specific recipient.
     * @param token The address of the token to bridge.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge transaction was initiated successfully.
     */
    function bridgeTo(IERC20 token, address to, uint value, uint gas, uint service, bytes[] calldata extraData)
        public
        payable
        whenNotPaused
        checkValidToken(address(token))
        nonReentrant
        returns (bool)
    {
        (uint minimum, uint _gas, uint _service, bool ok) = _checkAmount(token, value, gas, service);
        require(ok, BridgeStandardInvalidAmount(minimum, _gas, _service, value, gas, service));
        _bridge(token, _msgSender(), to, value, _gas, _service, false, extraData);
        return true;
    }

    /**
     * @notice Initiates a bridge transaction using a permit.
     * @param token The address of the token to bridge.
     * @param account The address of the account initiating the transaction.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param permitArgs The permit arguments.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge transaction was initiated successfully.
     */
    function permitBridge(
        IERC20 token,
        address account,
        uint value,
        uint gas,
        uint service,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) public payable returns (bool) {
        return permitBridgeTo(token, account, account, value, gas, service, permitArgs, extraData);
    }

    /**
     * @notice Initiates a bridge transaction to a specific recipient using a permit.
     * @param token The address of the token to bridge.
     * @param from The address of the account initiating the transaction.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param permitArgs The permit arguments.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge transaction was initiated successfully.
     */
    function permitBridgeTo(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) public payable whenNotPaused checkValidToken(address(token)) nonReentrant returns (bool) {
        (uint minimum, uint _gas, uint _service, bool ok) = _checkAmount(token, value, gas, service);
        require(ok, BridgeStandardInvalidAmount(minimum, _gas, _service, value, gas, service));
        _permitBridge(token, from, to, value, _gas, _service, permitArgs, extraData);
        return true;
    }

    /**
     * @notice Finalizes a bridge transaction.
     * @param index The index of the bridge transaction.
     * @param token The address of the token.
     * @param to The address of the recipient.
     * @param value The amount of tokens.
     * @param extraData Additional data for the bridge transaction.
     * @param sigs The signatures of the validators.
     * @return True if the bridge transaction was finalized successfully.
     */
    function finalize(uint index, IERC20 token, address to, uint value, bytes[] calldata extraData, bytes[] memory sigs)
        public
        payable
        whenNotPaused
        checkValidToken(address(token))
        nonReentrant
        returns (bool)
    {
        uint nextIndex = nextFinalizeIndex();
        require(index == nextIndex, BridgeStandardInvalidIndex(nextIndex, index));

        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, address(token), to, value, extraData));
        require(_validate(h, sigs), BridgeStandardInvalidSignatures(index));

        _incrementFinalizeIndex();
        (bool ok, bytes memory reason) = _finalizeBridge(token, to, value);
        if (ok) {
            emit BridgeFinalized(index, token, to, value, block.timestamp);
        } else {
            require(_revertedIndex.add(index), BridgeStandardDuplicateIndex(index));
            _revertedReason[index] = reason;
            _revertedArguments[index] = FinalizeArguments(index, token, to, value, extraData);
            emit BridgeFinalizeReverted(index);
        }
        return true;
    }

    /**
     * @notice Finalizes multiple bridge transactions.
     * @param args An array of FinalizeArguments.
     * @param sigs A 2D array of signatures.
     * @return True if all bridge transactions were finalized successfully.
     */
    function finalizeBatch(FinalizeArguments[] calldata args, bytes[][] memory sigs) external payable returns (bool) {
        for (uint i = 0; i < args.length; i++) {
            finalize(args[i].index, args[i].token, args[i].to, args[i].value, args[i].extraData, sigs[i]);
        }
        return true;
    }

    /**
     * @notice Retries a reverted finalize transaction.
     * @param index The index of the reverted transaction.
     * @return True if the retry was successful.
     */
    function retryFinalize(uint index) public whenNotPaused nonReentrant returns (bool) {
        require(_revertedIndex.contains(index), BridgeStandardNotExistingIndex(index));
        FinalizeArguments memory args = _revertedArguments[index];

        (bool ok, bytes memory reason) = _finalizeBridge(args.token, args.to, args.value);
        require(ok, string(reason));

        _revertedIndex.remove(index);
        delete(_revertedReason[index]);
        delete(_revertedArguments[index]);

        emit BridgeFinalized(args.index, args.token, args.to, args.value, block.timestamp);
        return true;
    }

    /**
     * @notice Retries multiple reverted finalize transactions.
     * @param indexes An array of indexes of the reverted transactions.
     * @return True if all retries were successful.
     */
    function retryFinalizeBatch(uint[] memory indexes) external returns (bool) {
        for (uint i = 0; i < indexes.length; i++) {
            retryFinalize(indexes[i]);
        }
        return true;
    }

    /**
     * @notice Internal function to handle permit and bridge logic.
     * @param token The address of the token to bridge.
     * @param from The address of the account initiating the transaction.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param permitArgs The permit arguments.
     * @param extraData Additional data for the bridge transaction.
     */
    function _permitBridge(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) private {
        require(permitArgs.value >= value + gas + service, BridgeStandardInvalidPermitValue(address(token), value));
        _permitCall(permitArgs);
        _bridge(token, from, to, value, gas, service, true, extraData);
    }

    /**
     * @notice Internal function to handle bridge logic.
     * @param token The address of the token to bridge.
     * @param from The address of the account initiating the transaction.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param permit Whether the transaction uses a permit.
     * @param extraData Additional data for the bridge transaction.
     */
    function _bridge(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        bool permit,
        bytes[] calldata extraData
    ) private {
        _initiateBridge(token, from, value, gas + service);
        _emitInitiate(nextInitiateIndex(), token, from, to, value, gas, service, permit, extraData);
        _incrementInitiateIndex();
    }

    /**
     * @notice Internal function to emit the BridgeInitiated event.
     * @param index The index of the bridge transaction.
     * @param token The address of the token to bridge.
     * @param from The address of the account initiating the transaction.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param gas The gas fee for the transaction.
     * @param service The service fee for the transaction.
     * @param permit Whether the transaction uses a permit.
     * @param extraData Additional data for the bridge transaction.
     */
    function _emitInitiate(
        uint index,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        bool permit,
        bytes[] calldata extraData
    ) internal {
        emit BridgeInitiated(
            index, token, IERC20(getPairToken(address(token))), from, to, value, block.timestamp, permit, extraData
        );

        emit BridgeFeeCharged(index, token, from, gas, service);
    }

    /**
     * @notice Internal function to check the validity of the amounts.
     * @param token The address of the token.
     * @param value The amount of tokens.
     * @param gas The gas fee.
     * @param service The service fee.
     * @return minimum The minimum amount of tokens required.
     * @return _gas The calculated gas fee.
     * @return _service The calculated service fee.
     * @return ok True if the amounts are valid.
     */
    function _checkAmount(IERC20 token, uint value, uint gas, uint service)
        private
        view
        returns (uint minimum, uint _gas, uint _service, bool ok)
    {
        if (address(bridgeTokenInfo) == address(0)) return (0, 0, 0, true);
        (minimum, _gas, _service) = bridgeTokenInfo.calculate(address(token), value);
        if ((gas >= _gas) && (service >= _service) && (value >= minimum)) ok = true;
    }

    /**
     * @notice Internal function to execute a permit call for ERC20 tokens.
     * @param permitArgs The permit arguments struct containing necessary data for the permit call.
     * This includes the token, account, value, deadline, v, r, and s parameters.
     * @dev Reverts if the permit call fails or if the token is not a valid ERC20Permit contract.
     */
    function _permitCall(PermitArguments memory permitArgs) private {
        bytes memory data = abi.encodeCall(
            permitArgs.token.permit,
            (
                permitArgs.account,
                address(this),
                permitArgs.value,
                permitArgs.deadline,
                permitArgs.v,
                permitArgs.r,
                permitArgs.s
            )
        );

        IERC20Permit token = permitArgs.token;
        uint returnSize;
        uint returnValue;

        assembly ("memory-safe") {
            let success := call(gas(), token, 0, add(data, 0x20), mload(data), 0, 0x20)
            if iszero(success) {
                let ptr := mload(0x40)
                returndatacopy(ptr, 0, returndatasize())
                revert(ptr, returndatasize())
            }
            returnSize := returndatasize()
            returnValue := mload(0)
        }

        if (returnSize == 0 ? address(permitArgs.token).code.length == 0 : returnValue != 1) {
            revert BridgeStandardFailedPermit(address(permitArgs.token));
        }
    }

    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal virtual;
    function _finalizeBridge(IERC20 token, address to, uint value) internal virtual returns (bool, bytes memory);

    /**
     * db    db d888888b d88888b db   d8b   db      d88888b db    db d8b   db  .o88b. d888888b d888888b  .d88b.  d8b   db
     * 88    88   `88'   88'     88   I8I   88      88'     88    88 888o  88 d8P  Y8 `~~88~~'   `88'   .8P  Y8. 888o  88
     * Y8    8P    88    88ooooo 88   I8I   88      88ooo   88    88 88V8o 88 8P         88       88    88    88 88V8o 88
     * `8b  d8'    88    88~~~~~ Y8   I8I   88      88~~~   88    88 88 V8o88 8b         88       88    88    88 88 V8o88
     *  `8bd8'    .88.   88.     `8b d8'8b d8'      88      88b  d88 88  V888 Y8b  d8    88      .88.   `8b  d8' 88  V888
     *    YP    Y888888P Y88888P  `8b8' `8d8'       YP      ~Y8888P' VP   V8P  `Y88P'    YP    Y888888P  `Y88P'  VP   V8P
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    function initializedAt() external view returns (uint) {
        return _initializedAt;
    }

    function rewardWallet() public view returns (address payable) {
        return _rewardWallet;
    }

    function revertedArguments(uint index) public view returns (FinalizeArguments memory) {
        return _revertedArguments[index];
    }

    function revertedReason(uint index) public view returns (bytes memory) {
        return _revertedReason[index];
    }

    // BridgeTokenInfo
    function calculate(IERC20 token, uint value) external view returns (uint minimum, uint gas, uint service) {
        (minimum, gas, service) = bridgeTokenInfo.calculate(address(token), value);
    }

    function denominator() external view returns (uint) {
        return bridgeTokenInfo.denominator();
    }

    function getTokenInfo(IERC20 token) public view returns (IBridgeTokenInfo.TokenInfo memory) {
        return bridgeTokenInfo.getTokenInfo(address(token));
    }

    function tokenInfoLength() public view returns (uint) {
        return bridgeTokenInfo.tokensLength();
    }

    function tokenInfoByIndex(uint index) external view returns (IBridgeTokenInfo.TokenInfo memory) {
        return bridgeTokenInfo.tokenInfoByIndex(index);
    }

    function allTokenInfo() external view returns (IBridgeTokenInfo.TokenInfo[] memory) {
        return bridgeTokenInfo.allTokenInfo();
    }

    /**
     * .d8888. d88888b d888888b      d88888b db    db d8b   db  .o88b. d888888b d888888b  .d88b.  d8b   db
     * 88'  YP 88'     `~~88~~'      88'     88    88 888o  88 d8P  Y8 `~~88~~'   `88'   .8P  Y8. 888o  88
     * `8bo.   88ooooo    88         88ooo   88    88 88V8o 88 8P         88       88    88    88 88V8o 88
     *   `Y8b. 88~~~~~    88         88~~~   88    88 88 V8o88 8b         88       88    88    88 88 V8o88
     * db   8D 88.        88         88      88b  d88 88  V888 Y8b  d8    88      .88.   `8b  d8' 88  V888
     * `8888Y' Y88888P    YP         YP      ~Y8888P' VP   V8P  `Y88P'    YP    Y888888P  `Y88P'  VP   V8P
     */
    function addToken(IERC20 token, IERC20 pair) public onlyOwner {
        _addTokenPair(address(token), address(pair));
    }

    function removeToken(IERC20 token) external onlyOwner {
        _removeTokenPair(address(token));
    }

    function pauseToken(IERC20 token) external onlyOwner {
        _pauseToken(address(token));
    }

    function unpauseToken(IERC20 token) external onlyOwner {
        _unpauseToken(address(token));
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function changeRewardWallet(address payable rewardWallet_) external onlyOwner {
        _rewardWallet = rewardWallet_;
    }

    function setTokenInfo(IBridgeTokenInfo _bridgeTokenInfo) external onlyOwner {
        require(address(_bridgeTokenInfo) != address(0), BridgeStandardCanNotZeroAddress("bridgeTokenInfo"));
        bridgeTokenInfo = _bridgeTokenInfo;
    }

    function removeTokenInfo() external onlyOwner {
        bridgeTokenInfo = IBridgeTokenInfo(address(0));
    }

    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
