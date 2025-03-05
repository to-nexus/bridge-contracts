// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IBridgeFeeStation} from "../interface/IBridgeFeeStation.sol";
import {IStandardBridge} from "../interface/IStandardBridge.sol";
import {ICrossMintableERC20} from "../token/ICrossMintableERC20.sol";
import {BridgeRegistry} from "./BridgeRegistry.sol";
import {ValidatorManager} from "./ValidatorManager.sol";

/**
 * @title Standard
 * @notice This abstract contract provides a standard implementation for cross-chain bridging functionality.
 * It inherits from several OpenZeppelin upgradeable contracts for security and functionality.
 */
abstract contract StandardBridge is
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    ValidatorManager,
    BridgeRegistry,
    IStandardBridge
{
    using EnumerableSet for EnumerableSet.UintSet;
    using SafeERC20 for IERC20;

    error StandardInvalidValueUnit(IERC20 token, uint value);
    error StandardInvalidValue(uint expected, uint actual);
    error StandardInvalidIndex(uint expected, uint actual);
    error StandardInvalidMsgValue(uint expected, uint actual);
    error StandardInvalidBalance(uint expected, uint actual);
    error StandardInvalidMinAmount(uint expected, uint actual);
    error StandardInvalidGasFee(uint expected, uint actual);
    error StandardInvalidExFee(uint expected, uint actual);
    error StandardInvalidPendingAmount(uint remoteChainID, address token, uint pendingAmount, uint value);
    error StandardCanNotZeroMsgValue();
    error StandardCanNotZeroAddress();
    error StandardNotExistToken(address token);
    error StandardNotExistFeeStation();
    error StandardBurnFailed(IERC20 token, address from, uint value);
    error StandardNotMatchLength();
    error StandardFailedPermit();
    error StandardFailedSendValue();

    event BridgeInitiated(
        uint indexed remoteChainID,
        uint indexed index,
        IERC20 indexed localToken,
        IERC20 remoteToken,
        address from,
        address to,
        uint value,
        uint time,
        bool permit,
        bytes extraData
    );
    event BridgeFinalized(
        uint indexed remoteChainID, uint indexed index, IERC20 token, address indexed to, uint value, uint time
    );
    event BridgeFinalizePending(uint indexed index);
    event BridgeFeeCharged(uint indexed index, IERC20 indexed token, address indexed account, uint gasFee, uint exFee);
    event RewardWalletSet(address indexed wallet);
    event FeeStationSet(IBridgeFeeStation indexed feeStation);
    event BridgeTokenBatchProcessed(bool indexed permit, bool[] success, string[] reason);

    address internal constant NATIVE_TOKEN = address(1);
    bytes32 private constant FINALIZE_TYPEHASH = keccak256(
        "Finalize(uint256 remoteChainID,uint256 index,address token,address to,uint256 value,bytes extraData)"
    );

    IBridgeFeeStation public bridgeFeeStation;
    address payable private _nexus;
    uint private _initializedAt;

    uint[47] private __gap;

    constructor() {
        _disableInitializers();
    }

    receive() external payable {
        require(msg.value != 0, StandardCanNotZeroMsgValue());
    }

    function __StandardBridge_init(uint8 _threshold, address nexus_) internal onlyInitializing {
        __Ownable_init(_msgSender());
        __UUPSUpgradeable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __Validator_init(_threshold);

        require(address(nexus_) != address(0), StandardCanNotZeroAddress());
        _nexus = payable(nexus_);

        _initializedAt = block.number;
    }

    /**
     * ░█▀▀░█░█░█▀▀░█▀▀░█░█░▀█▀░█▀▀░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░█▀▀░▄▀▄░█▀▀░█░░░█░█░░█░░█▀▀░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░▀▀▀░▀░▀░▀▀▀░▀▀▀░▀▀▀░░▀░░▀▀▀░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Bridges a token to a remote chain.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to bridge.
     * @param to The address to receive the token on the remote chain.
     * @param value The amount of the token to bridge.
     * @param gasFee The gas fee for the bridge.
     * @param exFee The extra fee for the bridge.
     * @param extraData Additional data for the bridge.
     * @return A boolean indicating whether the bridge was successful.
     */
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

    /**
     * @notice Bridges multiple tokens to remote chains in a batch.
     * @param args An array of BridgeTokenArguments containing the details of each bridge.
     */
    function bridgeTokenBatch(BridgeTokenArguments[] calldata args) external payable {
        bool[] memory success = new bool[](args.length);
        string[] memory reason = new string[](args.length);
        for (uint i = 0; i < args.length; ++i) {
            try this.bridgeToken(
                args[i].remoteChainID,
                args[i].token,
                args[i].to,
                args[i].value,
                args[i].gasFee,
                args[i].exFee,
                args[i].extraData
            ) {
                success[i] = true;
            } catch Error(string memory _reason) {
                reason[i] = _reason;
            } catch Panic(uint errorCode) {
                reason[i] = _handleRevert(errorCode);
            } catch (bytes memory lowLevelData) {
                reason[i] = _handleRevert(lowLevelData);
            }
        }

        emit BridgeTokenBatchProcessed(false, success, reason);
    }

    /**
     * @notice Bridges a token to a remote chain with a permit.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to bridge.
     * @param to The address to receive the token on the remote chain.
     * @param value The amount of the token to bridge.
     * @param gasFee The gas fee for the bridge.
     * @param exFee The extra fee for the bridge.
     * @param extraData Additional data for the bridge.
     * @param permitArgs The permit arguments for the token.
     * @return A boolean indicating whether the bridge was successful.
     */
    function permitBridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData,
        PermitArguments calldata permitArgs
    ) public payable whenNotPaused onlyValidToken(remoteChainID, address(token)) nonReentrant returns (bool) {
        (gasFee, exFee) = _checkAmount(remoteChainID, token, value, gasFee, exFee);

        {
            require(
                permitArgs.value >= value + gasFee + exFee,
                StandardInvalidValue(value + gasFee + exFee, permitArgs.value)
            );

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

            IERC20Permit _token = permitArgs.token;
            uint returnSize;
            uint returnValue;

            assembly ("memory-safe") {
                let success := call(gas(), _token, 0, add(data, 0x20), mload(data), 0, 0x20)
                if iszero(success) {
                    let ptr := mload(0x40)
                    returndatacopy(ptr, 0, returndatasize())
                    revert(ptr, returndatasize())
                }
                returnSize := returndatasize()
                returnValue := mload(0)
            }

            require(
                returnSize == 0 ? address(permitArgs.token).code.length != 0 : returnValue == 1, StandardFailedPermit()
            );
        }

        _executeBridge(remoteChainID, token, permitArgs.account, to, value, gasFee, exFee, true, extraData);
        return true;
    }

    /**
     * @notice Bridges multiple tokens to remote chains with permits in a batch.
     * @param args An array of BridgeTokenArguments containing the details of each bridge.
     * @param permitArgs An array of PermitArguments containing the permit details for each token.
     */
    function permitBridgeTokenBatch(BridgeTokenArguments[] calldata args, PermitArguments[] calldata permitArgs)
        external
        payable
    {
        require(args.length == permitArgs.length, StandardNotMatchLength());
        bool[] memory success = new bool[](args.length);
        string[] memory reason = new string[](args.length);
        for (uint i = 0; i < args.length; ++i) {
            try this.permitBridgeToken(
                args[i].remoteChainID,
                args[i].token,
                args[i].to,
                args[i].value,
                args[i].gasFee,
                args[i].exFee,
                args[i].extraData,
                permitArgs[i]
            ) {
                success[i] = true;
            } catch Error(string memory _reason) {
                reason[i] = _reason;
            } catch Panic(uint errorCode) {
                reason[i] = _handleRevert(errorCode);
            } catch (bytes memory lowLevelData) {
                reason[i] = _handleRevert(lowLevelData);
            }
        }

        emit BridgeTokenBatchProcessed(true, success, reason);
    }

    /**
     * @notice Finalizes a bridge on the local chain.
     * @param args The finalize arguments for the bridge.
     * @param v The array of v values for the signatures.
     * @param r The array of r values for the signatures.
     * @param s The array of s values for the signatures.
     * @return A boolean indicating whether the bridge was finalized successfully.
     */
    function finalizeBridge(FinalizeArguments calldata args, uint8[] memory v, bytes32[] memory r, bytes32[] memory s)
        public
        payable
        whenNotPaused
        nonReentrant
        returns (bool)
    {
        require(_tokenContains(args.remoteChainID, address(args.token)), StandardNotExistToken(address(args.token)));
        require(msg.value == 0, StandardInvalidMsgValue(0, msg.value));
        require(
            args.index == getNextFinalizeIndex(args.remoteChainID),
            StandardInvalidIndex(getNextFinalizeIndex(args.remoteChainID), args.index)
        );

        _validateSignature(
            keccak256(
                abi.encode(FINALIZE_TYPEHASH, args.index, address(args.token), args.to, args.value, args.extraData)
            ),
            v,
            r,
            s
        );
        _incrementFinalizeIndex(args.remoteChainID);

        (bool ok, string memory reason) = _finalizeBridge(args.remoteChainID, args.token, args.to, args.value);
        if (ok) {
            emit BridgeFinalized(args.remoteChainID, args.index, args.token, args.to, args.value, block.timestamp);
        } else {
            _setPendingArguments(args, reason);
            emit BridgeFinalizePending(args.index);
        }

        return true;
    }

    /**
     * @notice Finalizes multiple bridges on the local chain in a batch.
     * @param args An array of FinalizeArguments containing the details of each bridge.
     * @param v An array of arrays of v values for the signatures.
     * @param r An array of arrays of r values for the signatures.
     * @param s An array of arrays of s values for the signatures.
     * @return A boolean indicating whether all bridges were finalized successfully.
     */
    function finalizeBridgeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool) {
        for (uint i = 0; i < args.length; ++i) {
            finalizeBridge(args[i], v[i], r[i], s[i]);
        }
        return true;
    }

    /**
     * @notice Retries finalizing a bridge on the local chain.
     * @param remoteChainID The ID of the remote chain.
     * @param index The index of the bridge to retry.
     * @return A boolean indicating whether the bridge was retried successfully.
     */
    function retryFinalizeBridge(uint remoteChainID, uint index) public whenNotPaused nonReentrant returns (bool) {
        FinalizeArguments memory args = pendingArguments(remoteChainID, index);
        TokenPair storage tokenPair = _tokenPairs[remoteChainID][address(args.token)];
        if (tokenPair.isOrigin) {
            require(
                tokenPair.pendingAmount >= args.value,
                StandardInvalidPendingAmount(remoteChainID, address(args.token), tokenPair.pendingAmount, args.value)
            );
            tokenPair.pendingAmount -= args.value;
        }

        (bool ok, string memory reason) = _finalizeBridge(remoteChainID, args.token, args.to, args.value);
        require(ok, reason);

        _removePendingArguments(remoteChainID, index);

        emit BridgeFinalized(remoteChainID, args.index, args.token, args.to, args.value, block.timestamp);
        return true;
    }

    /**
     * @notice Retries finalizing multiple bridges on the local chain in a batch.
     * @param remoteChainID The ID of the remote chain.
     * @param indexes An array of indexes of the bridges to retry.
     * @return A boolean indicating whether all bridges were retried successfully.
     */
    function retryFinalizeBridgeBatch(uint remoteChainID, uint[] memory indexes) external returns (bool) {
        for (uint i = 0; i < indexes.length; ++i) {
            retryFinalizeBridge(remoteChainID, indexes[i]);
        }
        return true;
    }

    function _initiateBridge(uint remoteChainID, IERC20 token, address from, uint value, uint fee) internal virtual {
        uint localTokenRate = _tokenPairs[remoteChainID][address(token)].localTokenRate;
        if (localTokenRate > 1) require(value % localTokenRate == 0, StandardInvalidValueUnit(token, value));

        if (address(token) == NATIVE_TOKEN) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, StandardInvalidMsgValue(value + fee, msg.value));
            if (fee != 0) sendValue(_nexus, fee); // Send fees to the reward wallet
        } else {
            require(msg.value == 0, StandardInvalidMsgValue(0, msg.value));

            token.safeTransferFrom(from, address(this), value + fee);
            if (fee != 0) token.safeTransfer(_nexus, fee);

            if (_tokenPairs[remoteChainID][address(token)].isOrigin) {
                _depositToken(remoteChainID, address(token), value);
            } else {
                require(
                    ICrossMintableERC20(address(token)).burn(address(this), value), // Burn the wrapped tokens on the source chain
                    StandardBurnFailed(token, from, value)
                );
            }
        }
    }

    function _finalizeBridge(uint remoteChainID, IERC20 token, address to, uint value)
        internal
        virtual
        returns (bool ok, string memory reason)
    {
        TokenPair memory tokenPair = _tokenPairs[remoteChainID][address(token)];
        if (tokenPair.paused) return (false, "token is paused");
        if (tokenPair.safetyLimit != 0 && tokenPair.safetyLimit < value) {
            return (false, "value is greater than safety limit");
        }

        if (tokenPair.localTokenRate != 0) {
            if (tokenPair.localTokenRate != 1) value = value * tokenPair.localTokenRate;
            else value = value / tokenPair.remoteTokenRate;
        }

        if (address(token) == NATIVE_TOKEN) {
            sendValue(payable(to), value);
            ok = true;
            reason = "";
        } else if (value != 0) {
            if (_tokenPairs[remoteChainID][address(token)].isOrigin) {
                return _transferERC20(remoteChainID, token, to, value);
            } else {
                return _mintCrossMintableERC20(token, to, value);
            }
        }
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

        uint index;
        IERC20 remoteToken;
        {
            index = getNextInitiateIndex(remoteChainID);
            remoteToken = IERC20(_tokenPairs[remoteChainID][address(token)].remoteToken);
        }

        emit BridgeInitiated(
            remoteChainID, index, token, remoteToken, from, to, value, block.timestamp, permit, extraData
        );
        emit BridgeFeeCharged(index, token, from, gasFee, exFee);

        _incrementInitiateIndex(remoteChainID);
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
                reason = "StandardBridge: transfer failed";
            }
        } catch Error(string memory _reason) {
            reason = _reason;
        } catch Panic(uint errorCode) {
            reason = _handleRevert(errorCode);
        } catch (bytes memory lowLevelData) {
            reason = _handleRevert(lowLevelData);
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
                reason = "StandardBridge: mint failed";
            }
        } catch Error(string memory _reason) {
            reason = _reason;
        } catch Panic(uint errorCode) {
            reason = _handleRevert(errorCode);
        } catch (bytes memory lowLevelData) {
            reason = _handleRevert(lowLevelData);
        }
    }

    function sendValue(address payable recipient, uint amount) private {
        require(address(this).balance >= amount, StandardInvalidBalance(amount, address(this).balance));
        (bool success, bytes memory returndata) = recipient.call{value: amount}("");
        if (!success) _revert(returndata);
    }

    function _revert(bytes memory returndata) private pure {
        if (returndata.length > 0) {
            assembly ("memory-safe") {
                let returndata_size := mload(returndata)
                revert(add(32, returndata), returndata_size)
            }
        } else {
            revert StandardFailedSendValue();
        }
    }

    /**
     * ░█░█░▀█▀░█▀▀░█░█░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▄▀░░█░░█▀▀░█▄█░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░░▀░░▀▀▀░▀▀▀░▀░▀░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Returns the domain separator.
     * @return The domain separator.
     */
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /**
     * @notice Returns the block number at which the contract was initialized.
     * @return The block number at which the contract was initialized.
     */
    function initializedAt() external view returns (uint) {
        return _initializedAt;
    }

    /**
     * @notice Estimates the fee for bridging a token.
     * @param remoteChainID The ID of the remote chain.
     * @param token The address of the token to bridge.
     * @param value The amount of the token to bridge.
     * @return minimumAmount The minimum amount of the token to bridge.
     * @return gasFee The gas fee for the bridge.
     * @return exFee The extra fee for the bridge.
     */
    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee)
    {
        (minimumAmount, gasFee, exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
    }

    /**
     * @notice Returns the denominator used for fee calculations.
     * @return The denominator used for fee calculations.
     */
    function denominator() external view returns (uint) {
        return bridgeFeeStation.denominator();
    }

    /**
     * @notice Returns the address of the reward wallet.
     * @return The address of the reward wallet.
     */
    function nexus() public view returns (address payable) {
        return _nexus;
    }

    function _checkAmount(uint remoteChainID, IERC20 token, uint value, uint gasFee, uint exFee)
        private
        view
        returns (uint _gasFee, uint _exFee)
    {
        if (address(bridgeFeeStation) == address(0)) return (0, 0);
        uint minimumAmount;
        (minimumAmount, _gasFee, _exFee) = bridgeFeeStation.estimateFee(remoteChainID, token, value);
        require(value >= minimumAmount, StandardInvalidMinAmount(minimumAmount, value));
        require(gasFee >= _gasFee, StandardInvalidGasFee(_gasFee, gasFee));
        require(exFee >= _exFee, StandardInvalidExFee(_exFee, exFee));
    }

    function _handleRevert(uint errorCode) private pure returns (string memory reason) {
        return (string(abi.encodePacked("Panic code: ", errorCode)));
    }

    function _handleRevert(bytes memory lowLevelData) private pure returns (string memory reason) {
        return (string(abi.encodePacked("Low-level error: ", lowLevelData)));
    }

    /**
     * ░█▀▀░█▀▀░▀█▀░░░█▀▀░█░█░█▀█░█▀▀░▀█▀░▀█▀░█▀█░█▀█░█▀▀
     * ░▀▀█░█▀▀░░█░░░░█▀▀░█░█░█░█░█░░░░█░░░█░░█░█░█░█░▀▀█
     * ░▀▀▀░▀▀▀░░▀░░░░▀░░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀░▀▀▀
     */

    /**
     * @notice Pauses the contract.
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpauses the contract.
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    /**
     * @notice Sets the reward wallet.
     * @param nexus_ The address of the reward wallet.
     */
    function setRewardWallet(address payable nexus_) external onlyOwner {
        require(nexus_ != address(0), StandardCanNotZeroAddress());
        _nexus = nexus_;
        emit RewardWalletSet(_nexus);
    }

    /**
     * @notice Sets the fee station.
     * @param _bridgeFeeStation The address of the fee station.
     */
    function setFeeStation(IBridgeFeeStation _bridgeFeeStation) external onlyOwner {
        require(address(_bridgeFeeStation) != address(0), StandardCanNotZeroAddress());
        bridgeFeeStation = _bridgeFeeStation;
        emit FeeStationSet(bridgeFeeStation);
    }

    /**
     * @notice Removes the fee station.
     */
    function removeFeeStation() external onlyOwner {
        require(address(bridgeFeeStation) != address(0), StandardNotExistFeeStation());
        bridgeFeeStation = IBridgeFeeStation(address(0));
        emit FeeStationSet(bridgeFeeStation);
    }

    /**
     * @notice Authorizes an upgrade to a new implementation.
     * @param _newImplementation The address of the new implementation.
     */
    function _authorizeUpgrade(address _newImplementation) internal override onlyOwner {}
}
