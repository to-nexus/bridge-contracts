// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {ERC165Checker} from "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeFeeStation} from "../interface/IBridgeFeeStation.sol";
import {IStandardBridge} from "../interface/IStandardBridge.sol";
import {CrossMintableERC20, ICrossMintableERC20} from "../token/CrossMintableERC20.sol";
import {ChainManager} from "./ChainManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";

/**
 * @title StandardBridge
 * @notice This abstract contract provides a standard implementation for cross-chain bridging functionality.
 * It inherits from several OpenZeppelin upgradeable contracts for security and functionality.
 */
abstract contract StandardBridge is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    ChainManager,
    ValidatorManager,
    IStandardBridge
{
    using EnumerableSet for EnumerableSet.UintSet;
    using Address for address payable;
    using SafeERC20 for IERC20;

    error errStandardBridgeInvalidIndex(uint expected, uint actual);
    error errStandardBridgeInvalidAmount(
        uint minimumValue, uint expectedGas, uint expectedEx, uint actualValue, uint actualGas, uint actualEx
    );
    error errStandardBridgeInvalidValue(uint expected, uint actual);
    error errStandardBridgeInvalidPermitValue(address token, uint value);
    error errStandardBridgeCanNotZeroAddress(string name);
    error errStandardBridgeFailedPermit(address token);
    error errStandardBridgeBurnFailed(address token, address from, uint value);

    event BridgeInitiated(
        uint indexed remoteChainID,
        uint indexed index,
        IERC20 localToken,
        IERC20 remoteToken,
        address indexed from,
        address to,
        uint value,
        uint time,
        bool permit,
        bytes extraData
    );
    event BridgeFinalized(uint indexed index, IERC20 indexed token, address indexed to, uint value, uint time);
    event BridgeFinalizeReverted(uint indexed index);
    event BridgeFeeCharged(uint indexed index, IERC20 indexed token, address indexed account, uint gasFee, uint exFee);

    address internal constant NATIVE_TOKEN = address(1);
    bytes32 private constant FINALIZE_TYPEHASH =
        keccak256("Finalize(uint256 index,address token,address to,uint256 value,bytes extraData)");

    IBridgeFeeStation public bridgeFeeStation;
    address payable private _rewardWallet;
    uint private _initializedAt;

    uint[47] private __gap;

    constructor() {
        _disableInitializers();
    }

    receive() external payable {
        assert(msg.value != 0); // Ensure the receive function only accepts non-zero value
    }

    function __StandardBridge_init(
        uint8 _threshold,
        address rewardWallet_,
        address _crossMintableERC20FactoryCode,
        address _bridgeFeeStation
    ) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __Validator_init(_threshold);
        __ChainManager_init(_crossMintableERC20FactoryCode);
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        if (_bridgeFeeStation != address(0)) bridgeFeeStation = IBridgeFeeStation(_bridgeFeeStation);

        require(address(rewardWallet_) != address(0), errStandardBridgeCanNotZeroAddress("rewardWallet"));
        _rewardWallet = payable(rewardWallet_);

        _initializedAt = block.number;
    }

    function bridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData
    ) public payable whenNotPaused onlyValidToken(remoteChainID, address(token)) nonReentrant returns (bool) {
        (uint _gasFee, uint _exFee) = _checkAmount(remoteChainID, token, value, gasFee, exFee);
        _executeBridge(remoteChainID, token, _msgSender(), to, value, _gasFee, _exFee, false, extraData);
        return true;
    }

    function permitBridgeToken(
        uint remoteChainID,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        PermitArguments memory permitArgs,
        bytes calldata extraData
    ) external payable whenNotPaused onlyValidToken(remoteChainID, address(token)) nonReentrant returns (bool) {
        (uint _gasFee, uint _exFee) = _checkAmount(remoteChainID, token, value, gasFee, exFee);
        _executePermitBridge(remoteChainID, token, from, to, value, _gasFee, _exFee, permitArgs, extraData);
        return true;
    }

    function finalizeBridge(
        uint fromChainID,
        uint index,
        IERC20 token,
        address to,
        uint value,
        bytes calldata extraData,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) public payable whenNotPaused onlyValidToken(fromChainID, address(token)) nonReentrant returns (bool) {
        require(msg.value == 0, errStandardBridgeInvalidValue(0, msg.value));
        require(
            index == getNextFinalizeIndex(fromChainID),
            errStandardBridgeInvalidIndex(getNextFinalizeIndex(fromChainID), index)
        );

        _validateSignature(
            keccak256(abi.encode(FINALIZE_TYPEHASH, index, address(token), to, value, extraData)), v, r, s
        );
        _incrementFinalizeIndex(fromChainID);

        (bool ok, string memory reason) = _finalizeBridge(fromChainID, token, to, value);
        if (ok) {
            emit BridgeFinalized(index, token, to, value, block.timestamp);
        } else {
            _setRevertedArguments(fromChainID, index, token, to, value, extraData, reason);
            emit BridgeFinalizeReverted(index);
        }

        return true;
    }

    function finalizeBridgeBatch(
        uint fromChainID,
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool) {
        for (uint i = 0; i < args.length; ++i) {
            finalizeBridge(
                fromChainID,
                args[i].index,
                args[i].token,
                args[i].to,
                args[i].value,
                args[i].extraData,
                v[i],
                r[i],
                s[i]
            );
        }
        return true;
    }

    function retryFinalizeBridge(uint fromChainID, uint index) public whenNotPaused nonReentrant returns (bool) {
        FinalizeArguments memory args = _getRevertedArguments(fromChainID, index);

        (bool ok, string memory reason) = _finalizeBridge(fromChainID, args.token, args.to, args.value);
        require(ok, reason);

        _removeRevertedArguments(fromChainID, index);

        emit BridgeFinalized(args.index, args.token, args.to, args.value, block.timestamp);
        return true;
    }

    function retryFinalizeBridgeBatch(uint fromChainID, uint[] memory indexes) external returns (bool) {
        for (uint i = 0; i < indexes.length; ++i) {
            retryFinalizeBridge(fromChainID, indexes[i]);
        }
        return true;
    }

    function _initiateBridge(uint remoteChainID, IERC20 token, address from, uint value, uint fee) internal virtual {
        if (address(token) == NATIVE_TOKEN) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, errStandardBridgeInvalidValue(value + fee, msg.value));
            if (fee != 0) _rewardWallet.sendValue(fee); // Send fees to the reward wallet
        } else {
            require(msg.value == 0, errStandardBridgeInvalidValue(0, msg.value));

            token.safeTransferFrom(from, address(this), value + fee);
            if (fee != 0) token.safeTransfer(_rewardWallet, fee);

            if (_isOriginToken(remoteChainID, address(token))) {
                _depositToken(remoteChainID, address(token), value);
            } else {
                require(
                    ICrossMintableERC20(address(token)).burn(address(this), value), // Burn the wrapped tokens on the source chain
                    errStandardBridgeBurnFailed(address(token), from, value)
                );
            }
        }
    }

    function _finalizeBridge(uint remoteChainID, IERC20 token, address to, uint value)
        internal
        virtual
        returns (bool ok, string memory reason)
    {
        if (address(token) == NATIVE_TOKEN) {
            payable(to).sendValue(value);
            ok = true;
            reason = "";
        } else if (value != 0) {
            if (_isOriginToken(remoteChainID, address(token))) return _transferERC20(remoteChainID, token, to, value);
            else return _mintCrossMintableERC20(token, to, value);
        }
    }

    function _executePermitBridge(
        uint remoteChainID,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        PermitArguments memory permitArgs,
        bytes calldata extraData
    ) private {
        require(permitArgs.value >= value + gasFee + exFee, errStandardBridgeInvalidPermitValue(address(token), value));
        _executePermit(permitArgs);
        _executeBridge(remoteChainID, token, from, to, value, gasFee, exFee, true, extraData);
    }

    function _executeBridge(
        uint remoteChainID,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bool permit,
        bytes calldata extraData
    ) private {
        _initiateBridge(remoteChainID, token, from, value, gasFee + exFee);
        _emitInitiate(remoteChainID, token, from, to, value, gasFee, exFee, permit, extraData);
        _incrementInitiateIndex(remoteChainID);
    }

    function _emitInitiate(
        uint remoteChainID,
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bool permit,
        bytes calldata extraData
    ) private {
        uint index;
        IERC20 remoteToken;
        {
            index = getNextInitiateIndex(remoteChainID);
            remoteToken = IERC20(_getRemoteToken(remoteChainID, address(token)));
        }

        emit BridgeInitiated(
            remoteChainID, index, token, remoteToken, from, to, value, block.timestamp, permit, extraData
        );
        emit BridgeFeeCharged(index, token, from, gasFee, exFee);
    }

    function _checkAmount(uint remoteChainID, IERC20 token, uint value, uint gasFee, uint exFee)
        private
        view
        returns (uint _gasFee, uint _exFee)
    {
        if (address(bridgeFeeStation) == address(0)) return (0, 0);
        uint minimumAmount;
        (minimumAmount, _gasFee, _exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
        require(
            (value >= minimumAmount) && (gasFee >= _gasFee) && (exFee >= _exFee),
            errStandardBridgeInvalidAmount(minimumAmount, _gasFee, _exFee, value, gasFee, exFee)
        );
    }

    function _executePermit(PermitArguments memory permitArgs) private {
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
            revert errStandardBridgeFailedPermit(address(permitArgs.token));
        }
    }

    function _transferERC20(uint remoteChainID, IERC20 token, address to, uint value)
        private
        returns (bool ok, string memory reason)
    {
        try IERC20(token).transfer(to, value) returns (bool success) {
            if (success) {
                ok = true;
                reason = "";
                _withdrawToken(remoteChainID, address(token), value);
            } else {
                ok = false;
                reason = "StandardBridge: transfer failed";
            }
        } catch Error(string memory _reason) {
            ok = false;
            reason = _reason;
        } catch Panic(uint _errorCode) {
            ok = false;
            reason = string(abi.encodePacked("Panic code: ", _errorCode));
        } catch (bytes memory _lowLevelData) {
            ok = false;
            reason = string(abi.encodePacked("Low-level error: ", _lowLevelData));
        }
    }

    function _mintCrossMintableERC20(IERC20 token, address to, uint value)
        private
        returns (bool ok, string memory reason)
    {
        try ICrossMintableERC20(address(token)).mint(to, value) returns (bool success) {
            if (success) {
                ok = true;
                reason = "";
            } else {
                ok = false;
                reason = "StandardBridge: mint failed";
            }
        } catch Error(string memory _reason) {
            ok = false;
            reason = _reason;
        } catch Panic(uint _errorCode) {
            ok = false;
            reason = string(abi.encodePacked("Panic code: ", _errorCode));
        } catch (bytes memory _lowLevelData) {
            ok = false;
            reason = string(abi.encodePacked("Low-level error: ", _lowLevelData));
        }
    }

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

    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee)
    {
        (minimumAmount, gasFee, exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
    }

    function denominator() external view returns (uint) {
        return bridgeFeeStation.denominator();
    }

    /**
     * .d8888. d88888b d888888b      d88888b db    db d8b   db  .o88b. d888888b d888888b  .d88b.  d8b   db
     * 88'  YP 88'     `~~88~~'      88'     88    88 888o  88 d8P  Y8 `~~88~~'   `88'   .8P  Y8. 888o  88
     * `8bo.   88ooooo    88         88ooo   88    88 88V8o 88 8P         88       88    88    88 88V8o 88
     *   `Y8b. 88~~~~~    88         88~~~   88    88 88 V8o88 8b         88       88    88    88 88 V8o88
     * db   8D 88.        88         88      88b  d88 88  V888 Y8b  d8    88      .88.   `8b  d8' 88  V888
     * `8888Y' Y88888P    YP         YP      ~Y8888P' VP   V8P  `Y88P'    YP    Y888888P  `Y88P'  VP   V8P
     */
    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function setRewardWallet(address payable rewardWallet_) external onlyOwner {
        require(rewardWallet_ != address(0), errStandardBridgeCanNotZeroAddress("rewardWallet_"));
        _rewardWallet = rewardWallet_;
    }

    function setFeeStation(IBridgeFeeStation _bridgeFeeStation) external onlyOwner {
        require(address(_bridgeFeeStation) != address(0), errStandardBridgeCanNotZeroAddress("_bridgeFeeStation"));
        bridgeFeeStation = _bridgeFeeStation;
    }

    function removeFeeStation() external onlyOwner {
        bridgeFeeStation = IBridgeFeeStation(address(0));
    }

    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
