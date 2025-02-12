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

import {IBridgeFeeManager} from "../interface/IBridgeFeeManager.sol";
import {IBridgeStandard} from "../interface/IBridgeStandard.sol";
import {Indexer} from "./Indexer.sol";
import {TokenManager} from "./TokenManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";

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
    error BridgeStandardInvalidFee(uint expectedGas, uint expectedService, uint actualGas, uint actualService);
    error BridgeStandardInvalidSignature(uint index);
    error BridgeStandardDuplicateIndex(uint index);
    error BridgeStandardNotExistingIndex(uint index);
    error BridgeStandardCanNotZeroAddress(string name);

    event BridgeInitiated(
        uint indexed index,
        IERC20 indexed token,
        IERC20 pairToken,
        address indexed from,
        address to,
        uint value,
        uint gasFee,
        uint serviceFee,
        uint time,
        bytes[] extraData
    );
    event BridgeFinalized(uint indexed index, IERC20 indexed token, address indexed to, uint value, uint time);
    event BridgeFinalizeReverted(uint indexed index);

    uint internal constant _exrate = 100; // cross : xcross, 1 : 100
    bytes32 private constant FINALIZE_TYPEHASH =
        keccak256("Finalize(uint256 index,address token,address to,uint256 value)");

    IBridgeFeeManager public BridgeFeeManager;
    uint private _initializedAt;
    address payable private _rewardWallet;

    mapping(uint => FinalizeArguments) private _revertedArguments;
    mapping(uint => bytes) private _revertedReason;
    EnumerableSet.UintSet private _revertedIndex;

    uint[45] private __gap;

    constructor() {
        _disableInitializers();
    }

    function __BridgeStandard_init(address rewardWallet_, address _BridgeFeeManager) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __Validator_init();

        require(address(rewardWallet_) != address(0), BridgeStandardCanNotZeroAddress("rewardWallet"));
        // require(address(_BridgeFeeManager) != address(0), CanNotZeroAddress("BridgeFeeManager")); // allow zero address

        _initializedAt = block.timestamp;
        BridgeFeeManager = IBridgeFeeManager(_BridgeFeeManager);
        _rewardWallet = payable(rewardWallet_);
    }

    /**
     * @notice Initiates a bridge transfer
     * @param token The token to be bridged
     * @param value The amount of tokens to bridge
     * @param gas The gas fee for the bridge transaction
     * @param service The service fee for the bridge transaction
     * @param extraData Additional data for the bridge transaction
     */
    function bridge(IERC20 token, uint value, uint gas, uint service, bytes[] calldata extraData)
        public
        payable
        returns (bool)
    {
        return bridgeTo(token, _msgSender(), value, gas, service, extraData);
    }

    /**
     * @notice Initiates a bridge transfer to a specific address
     * @param token The token to be bridged
     * @param to The recipient address on the other chain
     * @param value The amount of tokens to bridge
     * @param gas The gas fee for the bridge transaction
     * @param service The service fee for the bridge transaction
     * @param extraData Additional data for the bridge transaction
     */
    function bridgeTo(IERC20 token, address to, uint value, uint gas, uint service, bytes[] calldata extraData)
        public
        payable
        whenNotPaused
        checkValidToken(address(token))
        nonReentrant
        returns (bool)
    {
        (uint _gas, uint _service, bool ok) = _checkFee(token, value, gas, service);
        require(ok, BridgeStandardInvalidFee(_gas, _service, gas, service));
        _bridge(token, _msgSender(), to, value, _gas, _service, extraData);
        return true;
    }

    function permitBridge(
        IERC20 token,
        address account,
        uint value,
        uint gas,
        uint service,
        uint deadline,
        bytes memory permitSig,
        bytes[] calldata extraData
    ) public payable returns (bool) {
        return permitBridgeTo(token, account, account, value, gas, service, deadline, permitSig, extraData);
    }

    function permitBridgeTo(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        uint deadline,
        bytes memory permitSig,
        bytes[] calldata extraData
    ) public payable whenNotPaused checkValidToken(address(token)) nonReentrant returns (bool) {
        (uint _gas, uint _service, bool ok) = _checkFee(token, value, gas, service);
        require(ok, BridgeStandardInvalidFee(_gas, _service, gas, service));
        _permitBridge(token, from, to, value, _gas, _service, deadline, permitSig, extraData);
        return true;
    }

    /**
     * @notice Finalizes a bridge transfer
     * @param index The index of the bridge transfer
     * @param token The token to be bridged
     * @param to The recipient address on the other chain
     * @param value The amount of tokens to bridge
     * @param extraData Additional data for the bridge transaction
     * @param sigs The signatures for the bridge transfer
     */
    function finalize(uint index, IERC20 token, address to, uint value, bytes[] calldata extraData, bytes[] memory sigs)
        public
        payable
        whenNotPaused
        checkValidToken(address(token))
        onlyValidator
        nonReentrant
        returns (bool)
    {
        uint nextIndex = nextFinalizeIndex();
        require(index == nextIndex, BridgeStandardInvalidIndex(nextIndex, index));

        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, to, value, extraData));
        require(_validate(h, sigs), BridgeStandardInvalidSignature(index));

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
     * @notice Finalizes multiple bridge transfers
     * @param args An array of FinalizeArguments containing the details of each bridge transfer
     */
    function finalizeBatch(FinalizeArguments[] calldata args, bytes[][] memory sigs) external payable returns (bool) {
        for (uint i = 0; i < args.length; i++) {
            finalize(args[i].index, args[i].token, args[i].to, args[i].value, args[i].extraData, sigs[i]);
        }
        return true;
    }

    /**
     * @notice Retries the finalization of a bridge transfer that was previously reverted
     * @param index The index of the bridge transfer to retry
     */
    function retryFinalize(uint index) external whenNotPaused onlyValidator nonReentrant returns (bool) {
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

    function _permitBridge(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        uint deadline,
        bytes memory permitSig,
        bytes[] calldata extraData
    ) private {
        require(permitSig.length == 65, "invalid permit signature");
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly ("memory-safe") {
            r := mload(add(permitSig, 0x20))
            s := mload(add(permitSig, 0x40))
            v := byte(0, mload(add(permitSig, 0x60)))
        }
        IERC20Permit(address(token)).permit(from, address(this), value + gas + service, deadline, v, r, s);
        _bridge(token, from, to, value, gas, service, extraData);
    }

    function _bridge(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        bytes[] calldata extraData
    ) private {
        _initiateBridge(token, from, value, gas + service);
        _emitInitiate(nextInitiateIndex(), token, from, to, value, gas, service, extraData);
        _incrementInitiateIndex();
    }

    function _emitInitiate(
        uint index,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        bytes[] calldata extraData
    ) internal {
        emit BridgeInitiated(
            index,
            token,
            IERC20(getPairToken(address(token))),
            from,
            to,
            value,
            gas,
            service,
            block.timestamp,
            extraData
        );
    }

    function _checkFee(IERC20 token, uint value, uint gas, uint service)
        private
        view
        returns (uint gas_, uint service_, bool ok)
    {
        if (address(BridgeFeeManager) == address(0)) return (0, 0, true);

        (uint _gas, uint _service) = BridgeFeeManager.calculateFee(address(token), value);
        if (gas < _gas) return (_gas, _service, false);
        else if (service < _service) return (_gas, _service, false);
        return (_gas, _service, true);
    }

    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal virtual;
    function _finalizeBridge(IERC20 token, address to, uint value) internal virtual returns (bool, bytes memory);
    // ----- View Functions -----

    function initializedAt() external view returns (uint) {
        return _initializedAt;
    }

    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    function rewardWallet() public view returns (address payable) {
        return _rewardWallet;
    }

    /**
     * @notice Returns the denominator used in fee calculations
     * @return The denominator value for fee calculations
     */
    function denominator() external view returns (uint) {
        return BridgeFeeManager.denominator();
    }

    /**
     * @notice Calculates the gas and service fees for a given token and value
     * @param token The token for which to calculate the fees
     * @param value The value for which to calculate the fees
     * @return gas The calculated gas fee
     * @return service The calculated service fee
     */
    function calculateFee(IERC20 token, uint value) external view returns (uint gas, uint service) {
        (gas, service) = BridgeFeeManager.calculateFee(address(token), value);
    }

    function getTokenFee(IERC20 token) public view returns (IBridgeFeeManager.FeeInfo memory) {
        return BridgeFeeManager.getTokenFee(address(token));
    }

    function BridgeFeeManagerLength() public view returns (uint) {
        return BridgeFeeManager.tokensLength();
    }

    function feeInfoByIndex(uint index) external view returns (IBridgeFeeManager.FeeInfo memory) {
        return BridgeFeeManager.feeInfoByIndex(index);
    }

    function allFeeInfo() external view returns (IBridgeFeeManager.FeeInfo[] memory) {
        return BridgeFeeManager.allFeeInfo();
    }

    function revertedArguments(uint index) public view returns (FinalizeArguments memory) {
        return _revertedArguments[index];
    }

    function revertedReason(uint index) public view returns (bytes memory) {
        return _revertedReason[index];
    }

    // ----- Set Functions -----

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

    // uups
    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
