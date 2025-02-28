// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BranchBridgeMetaData contains all meta data concerning the BranchBridge contract.
var BranchBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allRevertedIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"bridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossMintableERC20FactoryCode\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"internalType\":\"structIStandardBridge.PermitBridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryCode\",\"type\":\"address\"}],\"name\":\"setERC20Factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"success\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"reason\",\"type\":\"string[]\"}],\"name\":\"BridgeTokenFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20FactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RewardWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryAleadyExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"BridgeRegistryAleadyExistFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryAleadyExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryAleadyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeRegistryCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeRegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeRegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryInvalidRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeRegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeRegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"StandardBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"StandardBridgeFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBridgeInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"StandardBridgeNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permitAccount\",\"type\":\"address\"}],\"name\":\"StandardBridgeNotMatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"d80e3950": "allRevertedIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"f30589c3": "allValidators()",
		"47666cb1": "bridgeFeeStation()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"9118b5eb": "bridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[])",
		"b7f3358d": "changeThreshold(uint8)",
		"1003e37f": "createToken(uint256,address,uint256,uint256,string,uint8)",
		"8f517c17": "crossMintableERC20Factory()",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"ae766389": "estimateFee(uint256,address,uint256)",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"814914b5": "getTokenPair(uint256,address)",
		"e79805d3": "initialize(uint8,address,address)",
		"91cf6d3e": "initializedAt()",
		"facd743b": "isValidator(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"d2ff130d": "pauseToken(uint256,address)",
		"5c975abb": "paused()",
		"3b3bdcb1": "permitBridgeToken(uint256,address,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"9afead46": "permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))[])",
		"52d1902d": "proxiableUUID()",
		"57847893": "registerToken(uint256,bool,address,address,uint256,uint256)",
		"d7c82f32": "removeFeeStation()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
		"3960e787": "retryFinalizeBridge(uint256,uint256)",
		"030372c3": "retryFinalizeBridgeBatch(uint256,uint256[])",
		"952a95de": "revertedArguments(uint256,uint256)",
		"8415a385": "revertedReason(uint256,uint256)",
		"fb75b2c7": "rewardWallet()",
		"f516b879": "setERC20Factory(address)",
		"54db0126": "setFeeStation(address)",
		"5958621e": "setRewardWallet(address)",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"84d58d42": "unpauseToken(uint256,address)",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516162656100f95f395f818161323b0152818161326401526133a801526162655ff3fe60806040526004361061032a575f3560e01c80638da5cb5b116101a3578063cbae5958116100f2578063f2fde38b11610092578063f516b8791161006d578063f516b87914610a0a578063f698da2514610a29578063facd743b14610a3d578063fb75b2c714610a5c575f5ffd5b8063f2fde38b146109ab578063f30589c3146109ca578063f4509643146109eb575f5ffd5b8063d5717fc5116100cd578063d5717fc51461093a578063d7c82f3214610959578063d80e39501461096d578063e79805d31461098c575f5ffd5b8063cbae5958146108db578063cf56118e146108fa578063d2ff130d1461091b575f5ffd5b806396ce07951161015d578063ae6893f811610138578063ae6893f81461084f578063ae7663891461086e578063aed1d403146108a8578063b7f3358d146108bc575f5ffd5b806396ce0795146107f85780639afead461461080c578063ad3cb1cc1461081f575f5ffd5b80638da5cb5b1461072c5780638f517c17146107685780639118b5eb1461078657806391cf6d3e146107995780639300c926146107ad578063952a95de146107cc575f5ffd5b806354db012611610279578063715018a6116102195780638456cb59116101f45780638456cb59146106bf57806384b0196e146106d357806384d58d42146106fa57806388d67d6d14610719575f5ffd5b8063715018a6146105b8578063814914b5146105cc5780638415a38514610693575f5ffd5b80635b605f5c116102545780635b605f5c146105375780635c975abb146105635780635fd262de146105865780637101fcd314610599575f5ffd5b806354db0126146104da57806357847893146104f95780635958621e14610518575f5ffd5b80633b3bdcb1116102e457806342cde4e8116102bf57806342cde4e81461046557806347666cb1146104865780634f1ef286146104a557806352d1902d146104b8575f5ffd5b80633b3bdcb11461041f5780633f4ba83a1461043257806340a141ff14610446575f5ffd5b8063030372c3146103445780631003e37f146103785780631327d3d8146103af5780631938e0f2146103ce5780631d40f0d8146103e15780633960e78714610400575f5ffd5b3661034057345f0361033e5761033e614bdc565b005b5f5ffd5b34801561034f575f5ffd5b5061036361035e366004614c77565b610a79565b60405190151581526020015b60405180910390f35b348015610383575f5ffd5b50610397610392366004614db1565b610abd565b6040516001600160a01b03909116815260200161036f565b3480156103ba575f5ffd5b5061033e6103c9366004614e3a565b610bd2565b6103636103dc366004614f34565b610be0565b3480156103ec575f5ffd5b5061033e6103fb366004614fed565b610eec565b34801561040b575f5ffd5b5061036361041a366004615088565b610f26565b61036361042d3660046150e5565b61101d565b34801561043d575f5ffd5b5061033e61118b565b348015610451575f5ffd5b5061033e610460366004614e3a565b61119d565b348015610470575f5ffd5b5060365460405160ff909116815260200161036f565b348015610491575f5ffd5b50606754610397906001600160a01b031681565b61033e6104b336600461519b565b6111a7565b3480156104c3575f5ffd5b506104cc6111c2565b60405190815260200161036f565b3480156104e5575f5ffd5b5061033e6104f4366004614e3a565b6111de565b348015610504575f5ffd5b5061033e610513366004615207565b61127a565b348015610523575f5ffd5b5061033e610532366004614e3a565b611298565b348015610542575f5ffd5b50610556610551366004615268565b611330565b60405161036f91906152be565b34801561056e575f5ffd5b505f5160206161f05f395f51905f525460ff16610363565b61036361059436600461530b565b611489565b3480156105a4575f5ffd5b5061033e6105b3366004614fed565b611598565b3480156105c3575f5ffd5b5061033e6115ae565b3480156105d7575f5ffd5b506106866105e6366004615393565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9182526003602090815260408084206001600160a01b039384168552600901825292839020835160a08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260020154608082015290565b60405161036f91906153c1565b34801561069e575f5ffd5b506106b26106ad366004615088565b6115bf565b60405161036f91906153fd565b3480156106ca575f5ffd5b5061033e61166d565b3480156106de575f5ffd5b506106e761167d565b60405161036f9796959493929190615449565b348015610705575f5ffd5b5061033e610714366004615393565b611726565b610363610727366004615582565b61181a565b348015610737575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610397565b348015610773575f5ffd5b505f54610397906001600160a01b031681565b61033e6107943660046156be565b6118b5565b3480156107a4575f5ffd5b506069546104cc565b3480156107b8575f5ffd5b5061033e6107c7366004614fed565b611c1f565b3480156107d7575f5ffd5b506107eb6107e6366004615088565b611c56565b60405161036f91906156fc565b348015610803575f5ffd5b506104cc611d99565b61033e61081a3660046156be565b611e09565b34801561082a575f5ffd5b506106b2604051806040016040528060058152602001640352e302e360dc1b81525081565b34801561085a575f5ffd5b506104cc610869366004615268565b6121c5565b348015610879575f5ffd5b5061088d610888366004615760565b6121e1565b6040805193845260208401929092529082015260600161036f565b3480156108b3575f5ffd5b506104cc612270565b3480156108c7575f5ffd5b5061033e6108d6366004615795565b61227b565b3480156108e6575f5ffd5b506103976108f5366004615268565b6122cc565b348015610905575f5ffd5b5061090e6122d8565b60405161036f91906157ae565b348015610926575f5ffd5b5061033e610935366004615393565b6122e4565b348015610945575f5ffd5b506104cc610954366004615268565b6123df565b348015610964575f5ffd5b5061033e6123fb565b348015610978575f5ffd5b5061090e610987366004615268565b61248b565b348015610997575f5ffd5b5061033e6109a63660046157c0565b6124a7565b3480156109b6575f5ffd5b5061033e6109c5366004614e3a565b6125b7565b3480156109d5575f5ffd5b506109de6125f1565b60405161036f9190615806565b3480156109f6575f5ffd5b5061033e610a05366004615393565b6125fd565b348015610a15575f5ffd5b5061033e610a24366004614e3a565b612723565b348015610a34575f5ffd5b506104cc612734565b348015610a48575f5ffd5b50610363610a57366004614e3a565b61273d565b348015610a67575f5ffd5b506068546001600160a01b0316610397565b5f805b8251811015610ab157610aa884848381518110610a9b57610a9b615846565b6020026020010151610f26565b50600101610a7c565b50600190505b92915050565b5f610ac6612749565b5f546001600160a01b0316610aee57604051630b1e6f9360e21b815260040160405180910390fd5b5f546040516bffffffffffffffffffffffff19606089901b1660208201526001600160a01b0390911690634804a041906034016040516020818303038152906040528051906020012085604051602001610b489190615871565b60405160208183030381529060405286866040518563ffffffff1660e01b8152600401610b789493929190615892565b6020604051808303815f875af1158015610b94573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb891906158d1565b9050610bc8875f838989896127a4565b9695505050505050565b610bdd816001612ad4565b50565b5f610be9612b91565b8435610bfb6060870160408801614e3a565b5f828152600360205260409020610c159060070182612bc1565b8190610c4557604051632d2b991760e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff1615610ca45760405163270c009360e21b81526001600160a01b039091166004820152602401610c3c565b50610cad612be5565b5f348015610cd757604051630970ff1560e01b815260048101929092526024820152604401610c3c565b50610ce4905087356123df565b602088013514610cf488356123df565b88602001359091610d2157604051635522005360e01b815260048101929092526024820152604401610c3c565b50610db390507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020890135610d5d60608b0160408c01614e3a565b610d6d60808c0160608d01614e3a565b60808c0135610d7f60a08e018e6158ec565b604051602001610d959796959493929190615956565b60405160208183030381529060405280519060200120878787612c1c565b610dd187355f90815260036020526040902060020180546001019055565b5f80610e028935610de860608c0160408d01614e3a565b610df860808d0160608e01614e3a565b8c60800135612e0d565b915091508115610e8d57610e1c60808a0160608b01614e3a565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610e5d60608e0160408f01614e3a565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610ec6565b610e978982612efd565b60405160208a0135907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610ee260015f5160206162105f395f51905f5255565b5050949350505050565b5f5b8151811015610f2257610f1a828281518110610f0c57610f0c615846565b60200260200101515f612ad4565b600101610eee565b5050565b5f610f2f612b91565b610f37612be5565b5f610f428484611c56565b90505f5f610f5e86846040015185606001518660800151612e0d565b91509150818190610f825760405162461bcd60e51b8152600401610c3c91906153fd565b50610f8d8686612f93565b82606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8660400151876080015142604051610ff8939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a460019350505050610ab760015f5160206162105f395f51905f5255565b5f611026612b91565b5f8b81526003602052604090208b908b906110449060070182612bc1565b819061106f57604051632d2b991760e11b81526001600160a01b039091166004820152602401610c3c565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156110ce5760405163270c009360e21b81526001600160a01b039091166004820152602401610c3c565b506110d7612be5565b6110e76040850160208601614e3a565b6001600160a01b038c81169116148b6111066040870160208801614e3a565b9091611138576040516326e544d360e11b81526001600160a01b03928316600482015291166024820152604401610c3c565b50505f5f6111498f8f8d8d8d613037565b9150915061115f8f8f8f8f8f87878f8f8f613150565b60019450505061117b60015f5160206162105f395f51905f5255565b50509a9950505050505050505050565b611193612749565b61119b6131d7565b565b610bdd815f612ad4565b6111af613230565b6111b8826132d4565b610f2282826132dc565b5f6111cb61339d565b505f5160206161d05f395f51905f525b90565b6111e6612749565b6001600160a01b0381166112315760405163799d19f960e11b81526020600482015260116024820152702fb13934b233b2a332b2a9ba30ba34b7b760791b6044820152606401610c3c565b606780546001600160a01b0319166001600160a01b0383169081179091556040517f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444905f90a250565b611282612749565b6112908686868686866127a4565b505050505050565b6112a0612749565b6001600160a01b0381166112e75760405163799d19f960e11b815260206004820152600d60248201526c72657761726457616c6c65745f60981b6044820152606401610c3c565b606880546001600160a01b0319166001600160a01b0383169081179091556040517f5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b905f90a250565b5f81815260036020526040812060609161134c600783016133e6565b90505f81516001600160401b0381111561136857611368614bf0565b6040519080825280602002602001820160405280156113bf57816020015b6040805160a0810182525f808252602080830182905292820181905260608201819052608082015282525f199092019101816113865790505b5090505f5b825181101561148057836009015f8483815181106113e4576113e4615846565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160a08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b909304909116151560608201526002909101546080820152825183908390811061146d5761146d615846565b60209081029190910101526001016113c4565b50949350505050565b5f611492612b91565b5f898152600360205260409020899089906114b09060070182612bc1565b81906114db57604051632d2b991760e11b81526001600160a01b039091166004820152602401610c3c565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff161561153a5760405163270c009360e21b81526001600160a01b039091166004820152602401610c3c565b50611543612be5565b5f5f6115528d8d8c8c8c613037565b9150915061156e8d8d6115623390565b8e8e87875f8f8f6133f2565b60019450505061158a60015f5160206162105f395f51905f5255565b505098975050505050505050565b6115a56103fb60346133e6565b610bdd81611c1f565b6115b6612749565b61119b5f613438565b5f82815260036020908152604080832084845260040190915290208054606091906115e9906159a3565b80601f0160208091040260200160405190810160405280929190818152602001828054611615906159a3565b80156116605780601f1061163757610100808354040283529160200191611660565b820191905f5260205f20905b81548152906001019060200180831161164357829003601f168201915b5050505050905092915050565b611675612749565b61119b6134a8565b5f60608082808083815f5160206161b05f395f51905f5280549091501580156116a857506001810154155b6116ec5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610c3c565b6116f46134f0565b6116fc6135b0565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61172e612749565b5f8281526003602052604090206117486007820183612bc1565b8290611773576040516339fcd1cd60e11b81526001600160a01b039091166004820152602401610c3c565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff166117c65760405163f06fe74560e01b81526001600160a01b039091166004820152602401610c3c565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b191690555185917fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e991a3505050565b5f805b858110156118a85761189f87878381811061183a5761183a615846565b905060200281019061184c91906159db565b86838151811061185e5761185e615846565b602002602001015186848151811061187857611878615846565b602002602001015186858151811061189257611892615846565b6020026020010151610be0565b5060010161181d565b5060019695505050505050565b5f816001600160401b038111156118ce576118ce614bf0565b6040519080825280602002602001820160405280156118f7578160200160208202803683370190505b5090505f826001600160401b0381111561191357611913614bf0565b60405190808252806020026020018201604052801561194657816020015b60608152602001906001900390816119315790505b5090505f5b83811015611bdc5730635fd262de86868481811061196b5761196b615846565b905060200281019061197d91906159f9565b3587878581811061199057611990615846565b90506020028101906119a291906159f9565b6119b3906040810190602001614e3a565b8888868181106119c5576119c5615846565b90506020028101906119d791906159f9565b6119e8906060810190604001614e3a565b8989878181106119fa576119fa615846565b9050602002810190611a0c91906159f9565b606001358a8a88818110611a2257611a22615846565b9050602002810190611a3491906159f9565b608001358b8b89818110611a4a57611a4a615846565b9050602002810190611a5c91906159f9565b60a001358c8c8a818110611a7257611a72615846565b9050602002810190611a8491906159f9565b611a929060c08101906158ec565b6040518963ffffffff1660e01b8152600401611ab5989796959493929190615a0d565b6020604051808303815f875af1925050508015611aef575060408051601f3d908101601f19168201909252611aec91810190615a54565b60015b611bae57611afb615a6f565b806308c379a003611b3e5750611b0f615a87565b80611b1a5750611b76565b80838381518110611b2d57611b2d615846565b602002602001018190525050611bd4565b634e487b7103611b7657611b50615b09565b90611b5b5750611b76565b611b64816135ee565b838381518110611b2d57611b2d615846565b3d808015611b9f576040519150601f19603f3d011682016040523d82523d5f602084013e611ba4565b606091505b50611b648161362a565b506001838281518110611bc357611bc3615846565b911515602092830291909101909101525b60010161194b565b505f15157fb8748b102128effc65c1bc368e51a5d18b212136f6f6097a5730a990abca24308383604051611c11929190615b26565b60405180910390a250505050565b5f5b8151811015610f2257611c4e828281518110611c3f57611c3f615846565b60200260200101516001612ad4565b600101611c21565b611c9c6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f8381526003602081815260408084208685528301825292839020835160c0810185528154815260018201549281019290925260028101546001600160a01b03908116948301949094529182015490921660608301526004810154608083015260058101805460a084019190611d11906159a3565b80601f0160208091040260200160405190810160405280929190818152602001828054611d3d906159a3565b8015611d885780601f10611d5f57610100808354040283529160200191611d88565b820191905f5260205f20905b815481529060010190602001808311611d6b57829003601f168201915b505050505081525050905092915050565b606754604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015611de0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e049190615bc8565b905090565b5f816001600160401b03811115611e2257611e22614bf0565b604051908082528060200260200182016040528015611e4b578160200160208202803683370190505b5090505f826001600160401b03811115611e6757611e67614bf0565b604051908082528060200260200182016040528015611e9a57816020015b6060815260200190600190039081611e855790505b5090505f5b8381101561218f5730633b3bdcb1868684818110611ebf57611ebf615846565b9050602002810190611ed19190615bdf565b35878785818110611ee457611ee4615846565b9050602002810190611ef69190615bdf565b611f07906040810190602001614e3a565b888886818110611f1957611f19615846565b9050602002810190611f2b9190615bdf565b611f3c906060810190604001614e3a565b898987818110611f4e57611f4e615846565b9050602002810190611f609190615bdf565b611f71906080810190606001614e3a565b8a8a88818110611f8357611f83615846565b9050602002810190611f959190615bdf565b608001358b8b89818110611fab57611fab615846565b9050602002810190611fbd9190615bdf565b60a001358c8c8a818110611fd357611fd3615846565b9050602002810190611fe59190615bdf565b60c001358d8d8b818110611ffb57611ffb615846565b905060200281019061200d9190615bdf565b61201b9060e08101906158ec565b8f8f8d81811061202d5761202d615846565b905060200281019061203f9190615bdf565b610100016040518b63ffffffff1660e01b81526004016120689a99989796959493929190615bf4565b6020604051808303815f875af19250505080156120a2575060408051601f3d908101601f1916820190925261209f91810190615a54565b60015b612161576120ae615a6f565b806308c379a0036120f157506120c2615a87565b806120cd5750612129565b808383815181106120e0576120e0615846565b602002602001018190525050612187565b634e487b710361212957612103615b09565b9061210e5750612129565b612117816135ee565b8383815181106120e0576120e0615846565b3d808015612152576040519150601f19603f3d011682016040523d82523d5f602084013e612157565b606091505b506121178161362a565b50600183828151811061217657612176615846565b911515602092830291909101909101525b600101611e9f565b50600115157fb8748b102128effc65c1bc368e51a5d18b212136f6f6097a5730a990abca24308383604051611c11929190615b26565b5f818152600360205260408120600190810154610ab791615ce1565b60675460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa15801561223d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122619190615cf4565b91989097509095509350505050565b5f611e04603461363d565b612283612749565b6036805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f610ab7603483613646565b6060611e0460016133e6565b6122ec612749565b5f8281526003602052604090206123066007820183612bc1565b8290612331576040516339fcd1cd60e11b81526001600160a01b039091166004820152602401610c3c565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff161561238557604051637cd2bb6b60e11b81526001600160a01b039091166004820152602401610c3c565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b1916600160a81b1790555185917ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a391a3505050565b5f81815260036020526040812060020154610ab7906001615ce1565b612403612749565b6067546001600160a01b031661244f57604051632930b4a160e01b815260206004820152601060248201526f313934b233b2a332b2a9ba30ba34b7b760811b6044820152606401610c3c565b606780546001600160a01b03191690556040515f907f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444908290a2565b5f818152600360205260409020606090610ab7906005016133e6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156124eb5750825b90505f826001600160401b031660011480156125065750303b155b905081158015612514575080155b156125325760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561255c57845460ff60401b1916600160401b1785555b612567888888613651565b83156125ad57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6125bf612749565b6001600160a01b0381166125e857604051631e4fbdf760e01b81525f6004820152602401610c3c565b610bdd81613438565b6060611e0460346133e6565b612605612749565b6126106001836136fa565b82906126325760405163c125a76960e01b8152600401610c3c91815260200190565b506001600160a01b03811661267257604051633bba86ab60e01b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610c3c565b5f82815260036020526040902061268c6007820183613711565b82906126b7576040516339fcd1cd60e11b81526001600160a01b039091166004820152602401610c3c565b506001600160a01b0382165f81815260098301602052604080822080546001600160a01b03191681556001810180546001600160b01b03191690556002018290555185917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d91a3505050565b61272b612749565b610bdd81613725565b5f611e046138b1565b5f610ab7603483612bc1565b3361277b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461119b5760405163118cdaa760e01b8152336004820152602401610c3c565b6127af6001876136fa565b156127e3576127bf6001876138ba565b86906127e157604051636974200760e11b8152600401610c3c91815260200190565b505b5f8681526003602052604090206001600160a01b03851661283457604051633bba86ab60e01b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610c3c565b61284160078201866138c5565b859061286c57604051635538ae3160e01b81526001600160a01b039091166004820152602401610c3c565b506040805160a0810182526001600160a01b0380881680835287821660208085019182528b15158587019081525f606087018181526080880182815295825260098a0190935296909620945185549085166001600160a01b03199091161785559051600185018054965192511515600160a81b0260ff60a81b19931515600160a01b026001600160a81b031990981692909516919091179590951716919091179092559051600290910155828214612a7657825f036129605760405163079a0dfd60e41b815260206004820152600f60248201526e72656d6f7465546f6b656e5261746560881b6044820152606401610c3c565b815f036129a25760405163079a0dfd60e41b815260206004820152600f60248201526e72656d6f7465546f6b656e5261746560881b6044820152606401610c3c565b818311156129f3578160011480156129c257506129c0600a84615d33565b155b838390916129ec5760405163c2df5c2760e01b815260048101929092526024820152604401610c3c565b5050612a38565b826001148015612a0b5750612a09600a83615d33565b155b83839091612a355760405163c2df5c2760e01b815260048101929092526024820152604401610c3c565b50505b60408051808201825284815260208082018581525f8b8152600483528481206001600160a01b038b1682529092529290209051815590516001909101555b604080518715158152602081018590529081018390526001600160a01b03808616919087169089907f4336c2d0a5a91f2abb5c43ad295d612b03503ba19e54f7a9be8d9f249a242e3a9060600160405180910390a450505050505050565b612adc612749565b8015612b1e57612aed6034836138c5565b8290612b18576040516329a04e7760e21b81526001600160a01b039091166004820152602401610c3c565b50612b56565b612b29603483613711565b8290612b545760405163fdbc594760e01b81526001600160a01b039091166004820152602401610c3c565b505b604051811515906001600160a01b038416907f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d905f90a35050565b5f5160206161f05f395f51905f525460ff161561119b5760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f90815260018301602052604081205415155b9392505050565b5f5160206162105f395f51905f52805460011901612c1657604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b8251825181148015612c2e5750815181145b835183518392612c62576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610c3c565b505060365482915060ff16811015612c9057604051632fcba65760e11b8152600401610c3c91815260200190565b505f80826001600160401b03811115612cab57612cab614bf0565b604051908082528060200260200182016040528015612cd4578160200160208202803683370190505b5090505f5b83811015612de1575f612d44888381518110612cf757612cf7615846565b6020026020010151888481518110612d1157612d11615846565b6020026020010151888581518110612d2b57612d2b615846565b6020026020010151612d3c8d6138d9565b929190613905565b9050612d4f8161273d565b8190612d7a5760405163845a09e760e01b81526001600160a01b039091166004820152602401610c3c565b505f805b8451811015612dca57848181518110612d9957612d99615846565b60200260200101516001600160a01b0316836001600160a01b031603612dc25760019150612dca565b600101612d7e565b5080612dd7578460010194505b5050600101612cd9565b50603654829060ff168110156125ad57604051632fcba65760e11b8152600401610c3c91815260200190565b5f8481526004602090815260408083206001600160a01b03871684528252808320815180830190925280548083526001909101549282019290925260609115612e7e578051600114612e6c578051612e659085615d46565b9350612e7e565b6020810151612e7b9085615d5d565b93505b5f196001600160a01b03871601612ebc57612ea26001600160a01b03861685613931565b6001925060405180602001604052805f8152509150612ef2565b8315612ef257612ecc87876139c3565b15612ee757612edd878787876139fb565b9250925050612ef4565b612edd868686613b78565b505b94509492505050565b81355f90815260036020908152604090912090830135612f2060058301826138ba565b8190612f4257604051633b7f553d60e21b8152600401610c3c91815260200190565b505f81815260038301602052604090208490612f5e8282615e8d565b50505f8181526004830160205260409020612f798482615f2d565b5050505050565b60015f5160206162105f395f51905f5255565b5f828152600360205260409020612fad6005820183613ce2565b8290612fcf57604051635a8e93f960e11b8152600401610c3c91815260200190565b505f8281526004820160205260408120612fe891614b92565b5f828152600380830160205260408220828155600181018390556002810180546001600160a01b031990811690915591810180549092169091556004810182905590612f796005830182614b92565b6067545f9081906001600160a01b031661305557505f905080613146565b60675460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa1580156130ab573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130cf9190615cf4565b909450925090508086108015906130e65750828510155b80156130f25750818410155b81848489898990919293949561313e576040516378f2c0e960e01b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610c3c565b505050505050505b9550959350505050565b8361315b8688615ce1565b6131659190615ce1565b816040013510158987909161319e57604051638d392d0360e01b81526001600160a01b0390921660048301526024820152604401610c3c565b506131b890506131b336839003830183615fe7565b613ced565b6131cb8a8a8a8a8a8a8a60018b8b6133f2565b50505050505050505050565b6131df613dea565b5f5160206161f05f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016122c1565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806132b657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166132aa5f5160206161d05f395f51905f52546001600160a01b031690565b6001600160a01b031614155b1561119b5760405163703e46dd60e11b815260040160405180910390fd5b610bdd612749565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613336575060408051601f3d908101601f1916820190925261333391810190615bc8565b60015b61335e57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610c3c565b5f5160206161d05f395f51905f52811461338e57604051632a87526960e21b815260048101829052602401610c3c565b6133988383613e19565b505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461119b5760405163703e46dd60e11b815260040160405180910390fd5b60605f612bde83613e6e565b6134088a8a8a89613403898b615ce1565b613ec7565b61341a8a8a8a8a8a8a8a8a8a8a6140ef565b5f8a81526003602052604090206001908101805490910190556131cb565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6134b0612b91565b5f5160206161f05f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613218565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206161b05f395f51905f529161352e906159a3565b80601f016020809104026020016040519081016040528092919081815260200182805461355a906159a3565b80156135a55780601f1061357c576101008083540402835291602001916135a5565b820191905f5260205f20905b81548152906001019060200180831161358857829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206161b05f395f51905f529161352e906159a3565b6040516b02830b734b19031b7b2329d160a51b6020820152602c8101829052606090604c015b6040516020818303038152906040529050919050565b6060816040516020016136149190616067565b5f610ab7825490565b5f612bde83836141e7565b61365961420d565b61366233614256565b61366b83614267565b613674816142cd565b61367c6142e5565b6136846142ed565b61368c6142fd565b6001600160a01b0382166136d25760405163799d19f960e11b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610c3c565b50606880546001600160a01b0319166001600160a01b03929092169190911790555043606955565b5f8181526001830160205260408120541515612bde565b5f612bde836001600160a01b03841661430d565b6001600160a01b03811661376a57604051633bba86ab60e01b815260206004820152600b60248201526a666163746f7279436f646560a81b6044820152606401610c3c565b5f546001600160a01b031680156137a05760405163783b8aef60e01b81526001600160a01b039091166004820152602401610c3c565b505f816001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156137dd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613804919081019061608c565b604080513060208201520160408051601f198184030181529082905261382d929160200161610b565b60408051601f198184030181528282524660208401529250613868915f910160405160208183030381529060405280519060200120836143f0565b5f80546001600160a01b0319166001600160a01b039290921691821781556040517f18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b99190a25050565b5f611e04614484565b5f612bde83836144f7565b5f612bde836001600160a01b0384166144f7565b5f610ab76138e56138b1565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f61391588888888614543565b925092509250613925828261460b565b50909695505050505050565b8047101561395b5760405163cf47918160e01b815247600482015260248101829052604401610c3c565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f81146139a5576040519150601f19603f3d011682016040523d82523d5f602084013e6139aa565b606091505b5091509150816139bd576139bd816146c3565b50505050565b5f9182526003602090815260408084206001600160a01b0393909316845260099092019052902060010154600160a01b900460ff1690565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613a6a575060408051601f3d908101601f19168201909252613a6791810190615a54565b60015b613b0b57613a76615a6f565b806308c379a003613a9f5750613a8a615a87565b80613a955750613ad0565b5f92509050612ef4565b634e487b7103613ad057613ab1615b09565b90613abc5750613ad0565b5f9250613ac8816135ee565b915050612ef4565b3d808015613af9576040519150601f19603f3d011682016040523d82523d5f602084013e613afe565b606091505b505f9250613ac88161362a565b8015613b36576001925060405180602001604052805f8152509150613b318787866146ec565b612ef2565b505060408051808201909152601f81527f5374616e646172644272696467653a207472616e73666572206661696c65640060208201525f969095509350505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613be7575060408051601f3d908101601f19168201909252613be491810190615a54565b60015b613c7f57613bf3615a6f565b806308c379a003613c195750613c07615a87565b80613c125750613c47565b9050613cda565b634e487b7103613c4757613c2b615b09565b90613c365750613c47565b613c3f816135ee565b915050613cda565b3d808015613c70576040519150601f19603f3d011682016040523d82523d5f602084013e613c75565b606091505b50613c3f8161362a565b8015613c9f576001925060405180602001604052805f8152509150613cd8565b6040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b5f612bde838361430d565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180613d99576040513d5f823e3d81fd5b50505f513d91508115613db0578060011415613dbe565b84516001600160a01b03163b155b15612f79578451604051630251f0a360e31b81526001600160a01b039091166004820152602401610c3c565b5f5160206161f05f395f51905f525460ff1661119b57604051638dfc202b60e01b815260040160405180910390fd5b613e228261478e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613e665761339882826147f1565b610f22614863565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613ebb57602002820191905f5260205f20905b815481526020019060010190808311613ea7575b50505050509050919050565b5f8581526004602090815260408083206001600160a01b03881684529091529020548015801590613ef9575080600114155b15613f3e57613f088184615d33565b8590849015613f3b57604051635acebbd760e11b81526001600160a01b0390921660048301526024820152604401610c3c565b50505b5f196001600160a01b03861601613fb157613f598284615ce1565b3414613f658385615ce1565b349091613f8e57604051630970ff1560e01b815260048101929092526024820152604401610c3c565b50508115613fac57606854613fac906001600160a01b031683613931565b611290565b5f348015613fdb57604051630970ff1560e01b815260048101929092526024820152604401610c3c565b50613fff90508430613fed8587615ce1565b6001600160a01b038916929190614882565b811561401f5760685461401f906001600160a01b038781169116846148e9565b61402986866139c3565b1561403957613fac86868561491a565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af1158015614083573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906140a79190615a54565b8585859091926140e45760405163615f082960e11b81526001600160a01b0393841660048201529290911660248301526044820152606401610c3c565b505050505050505050565b5f5f6140fa8c6121c5565b915061412f8c8c5f9182526003602090815260408084206001600160a01b039384168552600901909152909120600101541690565b90508a6001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7848e8e8e428d8d8d60405161417a98979695949392919061611f565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a6040516141d1929190918252602082015260400190565b60405180910390a4505050505050505050505050565b5f825f0182815481106141fc576141fc615846565b905f5260205f200154905092915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661119b57604051631afcd79f60e31b815260040160405180910390fd5b61425e61420d565b610bdd8161495b565b61426f61420d565b6142b7604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250614963565b6036805460ff191660ff92909216919091179055565b6001600160a01b03811615610bdd57610bdd81613725565b61119b61420d565b6142f561420d565b61119b614975565b61430561420d565b61119b614995565b5f81815260018301602052604081205480156143e7575f61432f600183616169565b85549091505f9061434290600190616169565b90508082146143a1575f865f01828154811061436057614360615846565b905f5260205f200154905080875f01848154811061438057614380615846565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806143b2576143b261617c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610ab7565b5f915050610ab7565b5f8347101561441b5760405163cf47918160e01b815247600482015260248101859052604401610c3c565b81515f0361443c57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d15198115161561445d576040513d5f823e3d81fd5b6001600160a01b038116612bde5760405163b06ebf3d60e01b815260040160405180910390fd5b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6144ae61499d565b6144b6614a05565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f81815260018301602052604081205461453c57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610ab7565b505f610ab7565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561457c57505f91506003905082614601565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156145cd573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166145f857505f925060019150829050614601565b92505f91508190505b9450945094915050565b5f82600381111561461e5761461e616190565b03614627575050565b600182600381111561463b5761463b616190565b036146595760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561466d5761466d616190565b0361468e5760405163fce698f760e01b815260048101829052602401610c3c565b60038260038111156146a2576146a2616190565b03610f22576040516335e2f38360e21b815260048101829052602401610c3c565b8051156146d35780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f8381526003602090815260408083206001600160a01b03861684526009810190925282206002015490919081906147249085614a47565b84549193509150858584614764576040516376ff3ecd60e01b815260048101939093526001600160a01b0390911660248301526044820152606401610c3c565b5050506001600160a01b039094165f90815260099092016020525060409020600201919091555050565b806001600160a01b03163b5f036147c357604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610c3c565b5f5160206161d05f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161480d91906161a4565b5f60405180830381855af49150503d805f8114614845576040519150601f19603f3d011682016040523d82523d5f602084013e61484a565b606091505b509150915061485a858383614a6b565b95945050505050565b341561119b5760405163b398979f60e01b815260040160405180910390fd5b6040516001600160a01b0384811660248301528381166044830152606482018390526139bd9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614ac7565b6040516001600160a01b0383811660248301526044820183905261339891859182169063a9059cbb906064016148b7565b5f8381526003602090815260408083206001600160a01b038616845260090190915281206002018054839290614951908490615ce1565b9091555050505050565b6125bf61420d565b61496b61420d565b610f228282614b33565b61497d61420d565b5f5160206161f05f395f51905f52805460ff19169055565b612f8061420d565b5f5f5160206161b05f395f51905f52816149b56134f0565b8051909150156149cd57805160209091012092915050565b815480156149dc579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206161b05f395f51905f5281614a1d6135b0565b805190915015614a3557805160209091012092915050565b600182015480156149dc579392505050565b5f5f83831115614a5b57505f905080614a64565b50600190508183035b9250929050565b606082614a8057614a7b826146c3565b612bde565b8151158015614a9757506001600160a01b0384163b155b15614ac057604051639996b31560e01b81526001600160a01b0385166004820152602401610c3c565b5080612bde565b5f5f60205f8451602086015f885af180614ae6576040513d5f823e3d81fd5b50505f513d91508115614afd578060011415614b0a565b6001600160a01b0384163b155b156139bd57604051635274afe760e01b81526001600160a01b0385166004820152602401610c3c565b614b3b61420d565b5f5160206161b05f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614b748482615f2d565b5060038101614b838382615f2d565b505f8082556001909101555050565b508054614b9e906159a3565b5f825580601f10614bad575050565b601f0160209004905f5260205f2090810190610bdd91905b80821115614bd8575f8155600101614bc5565b5090565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b0382111715614c2357614c23614bf0565b60405250565b601f8201601f191681016001600160401b0381118282101715614c4e57614c4e614bf0565b6040525050565b5f6001600160401b03821115614c6d57614c6d614bf0565b5060051b60200190565b5f5f60408385031215614c88575f5ffd5b8235915060208301356001600160401b03811115614ca4575f5ffd5b8301601f81018513614cb4575f5ffd5b8035614cbf81614c55565b604051614ccc8282614c29565b80915082815260208101915060208360051b850101925087831115614cef575f5ffd5b6020840193505b82841015614d11578335825260209384019390910190614cf6565b809450505050509250929050565b6001600160a01b0381168114610bdd575f5ffd5b5f6001600160401b03821115614d4b57614d4b614bf0565b50601f01601f191660200190565b5f614d6383614d33565b604051614d708282614c29565b809250848152858585011115614d84575f5ffd5b848460208301375f6020868301015250509392505050565b803560ff81168114614dac575f5ffd5b919050565b5f5f5f5f5f5f60c08789031215614dc6575f5ffd5b863595506020870135614dd881614d1f565b9450604087013593506060870135925060808701356001600160401b03811115614e00575f5ffd5b8701601f81018913614e10575f5ffd5b614e1f89823560208401614d59565b925050614e2e60a08801614d9c565b90509295509295509295565b5f60208284031215614e4a575f5ffd5b8135612bde81614d1f565b5f82601f830112614e64575f5ffd5b8135614e6f81614c55565b604051614e7c8282614c29565b80915082815260208101915060208360051b860101925085831115614e9f575f5ffd5b602085015b83811015614ec357614eb581614d9c565b835260209283019201614ea4565b5095945050505050565b5f82601f830112614edc575f5ffd5b8135614ee781614c55565b604051614ef48282614c29565b80915082815260208101915060208360051b860101925085831115614f17575f5ffd5b602085015b83811015614ec3578035835260209283019201614f1c565b5f5f5f5f60808587031215614f47575f5ffd5b84356001600160401b03811115614f5c575f5ffd5b850160c08188031215614f6d575f5ffd5b935060208501356001600160401b03811115614f87575f5ffd5b614f9387828801614e55565b93505060408501356001600160401b03811115614fae575f5ffd5b614fba87828801614ecd565b92505060608501356001600160401b03811115614fd5575f5ffd5b614fe187828801614ecd565b91505092959194509250565b5f60208284031215614ffd575f5ffd5b81356001600160401b03811115615012575f5ffd5b8201601f81018413615022575f5ffd5b803561502d81614c55565b60405161503a8282614c29565b80915082815260208101915060208360051b85010192508683111561505d575f5ffd5b6020840193505b82841015610bc857833561507781614d1f565b825260209384019390910190615064565b5f5f60408385031215615099575f5ffd5b50508035926020909101359150565b5f5f83601f8401126150b8575f5ffd5b5081356001600160401b038111156150ce575f5ffd5b602083019150836020828501011115614a64575f5ffd5b5f5f5f5f5f5f5f5f5f5f8a8c036101e0811215615100575f5ffd5b8b359a5060208c013561511281614d1f565b995060408c013561512281614d1f565b985060608c013561513281614d1f565b975060808c0135965060a08c0135955060c08c0135945060e08c01356001600160401b03811115615161575f5ffd5b61516d8e828f016150a8565b90955093505060e060ff1982011215615184575f5ffd5b506101008b0190509295989b9194979a5092959850565b5f5f604083850312156151ac575f5ffd5b82356151b781614d1f565b915060208301356001600160401b038111156151d1575f5ffd5b8301601f810185136151e1575f5ffd5b6151f085823560208401614d59565b9150509250929050565b8015158114610bdd575f5ffd5b5f5f5f5f5f5f60c0878903121561521c575f5ffd5b86359550602087013561522e816151fa565b9450604087013561523e81614d1f565b9350606087013561524e81614d1f565b9598949750929560808101359460a0909101359350915050565b5f60208284031215615278575f5ffd5b5035919050565b80516001600160a01b03908116835260208083015190911690830152604080820151151590830152606080820151151590830152608090810151910152565b602080825282518282018190525f918401906040840190835b81811015615300576152ea83855161527f565b6020939093019260a092909201916001016152d7565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615322575f5ffd5b88359750602089013561533481614d1f565b9650604089013561534481614d1f565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615373575f5ffd5b61537f8b828c016150a8565b999c989b5096995094979396929594505050565b5f5f604083850312156153a4575f5ffd5b8235915060208301356153b681614d1f565b809150509250929050565b60a08101610ab7828461527f565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f612bde60208301846153cf565b5f8151808452602084019350602083015f5b8281101561543f578151865260209586019590910190600101615421565b5093949350505050565b60ff60f81b8816815260e060208201525f61546760e08301896153cf565b828103604084015261547981896153cf565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506154aa818561540f565b9a9950505050505050505050565b5f5f83601f8401126154c8575f5ffd5b5081356001600160401b038111156154de575f5ffd5b6020830191508360208260051b8501011115614a64575f5ffd5b5f82601f830112615507575f5ffd5b813561551281614c55565b60405161551f8282614c29565b80915082815260208101915060208360051b860101925085831115615542575f5ffd5b602085015b83811015614ec35780356001600160401b03811115615564575f5ffd5b615573886020838a0101614ecd565b84525060209283019201615547565b5f5f5f5f5f60808688031215615596575f5ffd5b85356001600160401b038111156155ab575f5ffd5b6155b7888289016154b8565b90965094505060208601356001600160401b038111156155d5575f5ffd5b8601601f810188136155e5575f5ffd5b80356155f081614c55565b6040516155fd8282614c29565b80915082815260208101915060208360051b85010192508a831115615620575f5ffd5b602084015b838110156156605780356001600160401b03811115615642575f5ffd5b6156518d602083890101614e55565b84525060209283019201615625565b50955050505060408601356001600160401b0381111561567e575f5ffd5b61568a888289016154f8565b92505060608601356001600160401b038111156156a5575f5ffd5b6156b1888289016154f8565b9150509295509295909350565b5f5f602083850312156156cf575f5ffd5b82356001600160401b038111156156e4575f5ffd5b6156f0858286016154b8565b90969095509350505050565b60208152815160208201526020820151604082015260018060a01b03604083015116606082015260018060a01b036060830151166080820152608082015160a08201525f60a083015160c08084015261575860e08401826153cf565b949350505050565b5f5f5f60608486031215615772575f5ffd5b83359250602084013561578481614d1f565b929592945050506040919091013590565b5f602082840312156157a5575f5ffd5b612bde82614d9c565b602081525f612bde602083018461540f565b5f5f5f606084860312156157d2575f5ffd5b6157db84614d9c565b925060208401356157eb81614d1f565b915060408401356157fb81614d1f565b809150509250925092565b602080825282518282018190525f918401906040840190835b818110156153005783516001600160a01b031683526020938401939092019160010161581f565b634e487b7160e01b5f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f612bde600d83018461585a565b848152608060208201525f6158aa60808301866153cf565b82810360408401526158bc81866153cf565b91505060ff8316606083015295945050505050565b5f602082840312156158e1575f5ffd5b8151612bde81614d1f565b5f5f8335601e19843603018112615901575f5ffd5b8301803591506001600160401b0382111561591a575f5ffd5b602001915036819003821315614a64575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615996908301848661592e565b9998505050505050505050565b600181811c908216806159b757607f821691505b6020821081036159d557634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be198336030181126159ef575f5ffd5b9190910192915050565b5f823560de198336030181126159ef575f5ffd5b8881526001600160a01b03888116602083015287166040820152606081018690526080810185905260a0810184905260e060c082018190525f906154aa908301848661592e565b5f60208284031215615a64575f5ffd5b8151612bde816151fa565b5f60033d11156111db5760045f5f3e505f5160e01c90565b5f60443d1015615a945790565b6040513d600319016004823e80513d60248201116001600160401b0382111715615abd57505090565b80820180516001600160401b03811115615ad8575050505090565b3d8401600319018282016020011115615af2575050505090565b615b0160208285010185614c29565b509392505050565b5f5f60233d1115615b2257602060045f3e50505f516001905b9091565b604080825283519082018190525f9060208501906060840190835b81811015615b615783511515835260209384019390920191600101615b41565b50508381036020850152809150845180825260208201925060208160051b830101602087015f5b83811015615bba57601f19858403018652615ba48383516153cf565b6020968701969093509190910190600101615b88565b509098975050505050505050565b5f60208284031215615bd8575f5ffd5b5051919050565b5f82356101de198336030181126159ef575f5ffd5b8a81526001600160a01b038a811660208301528981166040830152881660608201526080810187905260a0810186905260c081018590526101e060e082018190525f90615c44908301858761592e565b90508235615c5181614d1f565b6001600160a01b03166101008301526020830135615c6e81614d1f565b6001600160a01b03166101208301526040830135610140830152606083013561016083015260ff615ca160808501614d9c565b1661018083015260a08301356101a083015260c0909201356101c0909101529998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610ab757610ab7615ccd565b5f5f5f60608486031215615d06575f5ffd5b5050815160208301516040909301519094929350919050565b634e487b7160e01b5f52601260045260245ffd5b5f82615d4157615d41615d1f565b500690565b8082028115828204841417610ab757610ab7615ccd565b5f82615d6b57615d6b615d1f565b500490565b80546001600160a01b0319166001600160a01b0392909216919091179055565b601f82111561339857805f5260205f20601f840160051c81016020851015615db55750805b601f840160051c820191505b81811015612f79575f8155600101615dc1565b6001600160401b03831115615deb57615deb614bf0565b615dff83615df983546159a3565b83615d90565b5f601f841160018114615e30575f8515615e195750838201355b5f19600387901b1c1916600186901b178355612f79565b5f83815260208120601f198716915b82811015615e5f5786850135825560209485019460019092019101615e3f565b5086821015615e7b575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81358155602082013560018201556040820135615ea981614d1f565b615eb68160028401615d70565b506060820135615ec581614d1f565b615ed28160038401615d70565b506080820135600482015560a082013536839003601e19018112615ef4575f5ffd5b820180356001600160401b03811115615f0b575f5ffd5b602082019150803603821315615f1f575f5ffd5b6139bd818360058601615dd4565b81516001600160401b03811115615f4657615f46614bf0565b615f5a81615f5484546159a3565b84615d90565b6020601f821160018114615f8c575f8315615f755750848201515b5f19600385901b1c1916600184901b178455612f79565b5f84815260208120601f198516915b82811015615fbb5787850151825560209485019460019092019101615f9b565b5084821015615fd857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60e0828403128015615ff8575f5ffd5b5060405161600581614c04565b823561601081614d1f565b8152602083013561602081614d1f565b6020820152604083810135908201526060808401359082015261604560808401614d9c565b608082015260a0838101359082015260c0928301359281019290925250919050565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f612bde601183018461585a565b5f6020828403121561609c575f5ffd5b81516001600160401b038111156160b1575f5ffd5b8201601f810184136160c1575f5ffd5b80516160cc81614d33565b6040516160d98282614c29565b8281528660208486010111156160ed575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b5f615758616119838661585a565b8461585a565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f906154aa908301848661592e565b81810381811115610ab757610ab7615ccd565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f612bde828461585a56fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122019a0b85e40de20905bf4d550cb9234f3d98b631656af6f237638e5580b5764d764736f6c634300081c0033",
}

// BranchBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BranchBridgeMetaData.ABI instead.
var BranchBridgeABI = BranchBridgeMetaData.ABI

// Deprecated: Use BranchBridgeMetaData.Sigs instead.
// BranchBridgeFuncSigs maps the 4-byte function signature to its string representation.
var BranchBridgeFuncSigs = BranchBridgeMetaData.Sigs

// BranchBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BranchBridgeMetaData.Bin instead.
var BranchBridgeBin = BranchBridgeMetaData.Bin

// DeployBranchBridge deploys a new Ethereum contract, binding an instance of BranchBridge to it.
func DeployBranchBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BranchBridge, error) {
	parsed, err := BranchBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BranchBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BranchBridge{BranchBridgeCaller: BranchBridgeCaller{contract: contract}, BranchBridgeTransactor: BranchBridgeTransactor{contract: contract}, BranchBridgeFilterer: BranchBridgeFilterer{contract: contract}}, nil
}

// BranchBridge is an auto generated Go binding around an Ethereum contract.
type BranchBridge struct {
	BranchBridgeCaller     // Read-only binding to the contract
	BranchBridgeTransactor // Write-only binding to the contract
	BranchBridgeFilterer   // Log filterer for contract events
}

// BranchBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BranchBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BranchBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BranchBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BranchBridgeSession struct {
	Contract     *BranchBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BranchBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BranchBridgeCallerSession struct {
	Contract *BranchBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BranchBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BranchBridgeTransactorSession struct {
	Contract     *BranchBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BranchBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BranchBridgeRaw struct {
	Contract *BranchBridge // Generic contract binding to access the raw methods on
}

// BranchBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BranchBridgeCallerRaw struct {
	Contract *BranchBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BranchBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BranchBridgeTransactorRaw struct {
	Contract *BranchBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBranchBridge creates a new instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridge(address common.Address, backend bind.ContractBackend) (*BranchBridge, error) {
	contract, err := bindBranchBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BranchBridge{BranchBridgeCaller: BranchBridgeCaller{contract: contract}, BranchBridgeTransactor: BranchBridgeTransactor{contract: contract}, BranchBridgeFilterer: BranchBridgeFilterer{contract: contract}}, nil
}

// NewBranchBridgeCaller creates a new read-only instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeCaller(address common.Address, caller bind.ContractCaller) (*BranchBridgeCaller, error) {
	contract, err := bindBranchBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeCaller{contract: contract}, nil
}

// NewBranchBridgeTransactor creates a new write-only instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BranchBridgeTransactor, error) {
	contract, err := bindBranchBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTransactor{contract: contract}, nil
}

// NewBranchBridgeFilterer creates a new log filterer instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BranchBridgeFilterer, error) {
	contract, err := bindBranchBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeFilterer{contract: contract}, nil
}

// bindBranchBridge binds a generic wrapper to an already deployed contract.
func bindBranchBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BranchBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BranchBridge *BranchBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BranchBridge.Contract.BranchBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BranchBridge *BranchBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.Contract.BranchBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BranchBridge *BranchBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BranchBridge.Contract.BranchBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BranchBridge *BranchBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BranchBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BranchBridge *BranchBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BranchBridge *BranchBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BranchBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BranchBridge.Contract.UPGRADEINTERFACEVERSION(&_BranchBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BranchBridge.Contract.UPGRADEINTERFACEVERSION(&_BranchBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _BranchBridge.Contract.AllChainIDs(&_BranchBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BranchBridge.Contract.AllChainIDs(&_BranchBridge.CallOpts)
}

// AllRevertedIndex is a free data retrieval call binding the contract method 0xd80e3950.
//
// Solidity: function allRevertedIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeCaller) AllRevertedIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allRevertedIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllRevertedIndex is a free data retrieval call binding the contract method 0xd80e3950.
//
// Solidity: function allRevertedIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeSession) AllRevertedIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BranchBridge.Contract.AllRevertedIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// AllRevertedIndex is a free data retrieval call binding the contract method 0xd80e3950.
//
// Solidity: function allRevertedIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeCallerSession) AllRevertedIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BranchBridge.Contract.AllRevertedIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeSession) AllValidators() ([]common.Address, error) {
	return _BranchBridge.Contract.AllValidators(&_BranchBridge.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeCallerSession) AllValidators() ([]common.Address, error) {
	return _BranchBridge.Contract.AllValidators(&_BranchBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeCaller) BridgeFeeStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "bridgeFeeStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeSession) BridgeFeeStation() (common.Address, error) {
	return _BranchBridge.Contract.BridgeFeeStation(&_BranchBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) BridgeFeeStation() (common.Address, error) {
	return _BranchBridge.Contract.BridgeFeeStation(&_BranchBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeCaller) CrossMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "crossMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeSession) CrossMintableERC20Factory() (common.Address, error) {
	return _BranchBridge.Contract.CrossMintableERC20Factory(&_BranchBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) CrossMintableERC20Factory() (common.Address, error) {
	return _BranchBridge.Contract.CrossMintableERC20Factory(&_BranchBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) Denominator() (*big.Int, error) {
	return _BranchBridge.Contract.Denominator(&_BranchBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) Denominator() (*big.Int, error) {
	return _BranchBridge.Contract.Denominator(&_BranchBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeSession) DomainSeparator() ([32]byte, error) {
	return _BranchBridge.Contract.DomainSeparator(&_BranchBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _BranchBridge.Contract.DomainSeparator(&_BranchBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BranchBridge *BranchBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BranchBridge *BranchBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BranchBridge.Contract.Eip712Domain(&_BranchBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BranchBridge *BranchBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BranchBridge.Contract.Eip712Domain(&_BranchBridge.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumAmount *big.Int
		GasFee        *big.Int
		ExFee         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BranchBridge.Contract.EstimateFee(&_BranchBridge.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BranchBridge.Contract.EstimateFee(&_BranchBridge.CallOpts, remoteChainID, token, value)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextFinalizeIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextFinalizeIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextInitiateIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextInitiateIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) InitializedAt() (*big.Int, error) {
	return _BranchBridge.Contract.InitializedAt(&_BranchBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _BranchBridge.Contract.InitializedAt(&_BranchBridge.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeSession) IsValidator(validator common.Address) (bool, error) {
	return _BranchBridge.Contract.IsValidator(&_BranchBridge.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _BranchBridge.Contract.IsValidator(&_BranchBridge.CallOpts, validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeSession) Owner() (common.Address, error) {
	return _BranchBridge.Contract.Owner(&_BranchBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) Owner() (common.Address, error) {
	return _BranchBridge.Contract.Owner(&_BranchBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeSession) Paused() (bool, error) {
	return _BranchBridge.Contract.Paused(&_BranchBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeCallerSession) Paused() (bool, error) {
	return _BranchBridge.Contract.Paused(&_BranchBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _BranchBridge.Contract.ProxiableUUID(&_BranchBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BranchBridge.Contract.ProxiableUUID(&_BranchBridge.CallOpts)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCaller) RevertedArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "revertedArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryFinalizeArguments)).(*IBridgeRegistryFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _BranchBridge.Contract.RevertedArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCallerSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _BranchBridge.Contract.RevertedArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCaller) RevertedReason(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (string, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "revertedReason", remoteChainID, index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.RevertedReason(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCallerSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.RevertedReason(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeCaller) RewardWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "rewardWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeSession) RewardWallet() (common.Address, error) {
	return _BranchBridge.Contract.RewardWallet(&_BranchBridge.CallOpts)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) RewardWallet() (common.Address, error) {
	return _BranchBridge.Contract.RewardWallet(&_BranchBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeSession) Threshold() (uint8, error) {
	return _BranchBridge.Contract.Threshold(&_BranchBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeCallerSession) Threshold() (uint8, error) {
	return _BranchBridge.Contract.Threshold(&_BranchBridge.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BranchBridge.Contract.ValidatorByIndex(&_BranchBridge.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BranchBridge.Contract.ValidatorByIndex(&_BranchBridge.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) ValidatorLength() (*big.Int, error) {
	return _BranchBridge.Contract.ValidatorLength(&_BranchBridge.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) ValidatorLength() (*big.Int, error) {
	return _BranchBridge.Contract.ValidatorLength(&_BranchBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) BridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "bridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeTokenBatch is a paid mutator transaction binding the contract method 0x9118b5eb.
//
// Solidity: function bridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args) payable returns()
func (_BranchBridge *BranchBridgeTransactor) BridgeTokenBatch(opts *bind.TransactOpts, args []IStandardBridgeBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "bridgeTokenBatch", args)
}

// BridgeTokenBatch is a paid mutator transaction binding the contract method 0x9118b5eb.
//
// Solidity: function bridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args) payable returns()
func (_BranchBridge *BranchBridgeSession) BridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeTokenBatch(&_BranchBridge.TransactOpts, args)
}

// BridgeTokenBatch is a paid mutator transaction binding the contract method 0x9118b5eb.
//
// Solidity: function bridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args) payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) BridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeTokenBatch(&_BranchBridge.TransactOpts, args)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.ChangeThreshold(&_BranchBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.ChangeThreshold(&_BranchBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x1003e37f.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, localTokenRate, remoteTokenRate, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x1003e37f.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, localTokenRate, remoteTokenRate, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x1003e37f.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, localTokenRate, remoteTokenRate, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0xe79805d3.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode) returns()
func (_BranchBridge *BranchBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "initialize", _threshold, _rewardWallet, _crossMintableERC20FactoryCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xe79805d3.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode) returns()
func (_BranchBridge *BranchBridgeSession) Initialize(_threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet, _crossMintableERC20FactoryCode)
}

// Initialize is a paid mutator transaction binding the contract method 0xe79805d3.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode) returns()
func (_BranchBridge *BranchBridgeTransactorSession) Initialize(_threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet, _crossMintableERC20FactoryCode)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeSession) Pause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Pause(&_BranchBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Pause(&_BranchBridge.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) PauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "pauseToken", remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.PauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.PauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x3b3bdcb1.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, from, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x3b3bdcb1.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x3b3bdcb1.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x9afead46.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))[] args) payable returns()
func (_BranchBridge *BranchBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IStandardBridgePermitBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "permitBridgeTokenBatch", args)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x9afead46.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))[] args) payable returns()
func (_BranchBridge *BranchBridgeSession) PermitBridgeTokenBatch(args []IStandardBridgePermitBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeTokenBatch(&_BranchBridge.TransactOpts, args)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x9afead46.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))[] args) payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) PermitBridgeTokenBatch(args []IStandardBridgePermitBridgeTokenArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeTokenBatch(&_BranchBridge.TransactOpts, args)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x57847893.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate) returns()
func (_BranchBridge *BranchBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x57847893.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate) returns()
func (_BranchBridge *BranchBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x57847893.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveFeeStation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeFeeStation")
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeSession) RemoveFeeStation() (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveFeeStation(&_BranchBridge.TransactOpts)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveFeeStation() (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveFeeStation(&_BranchBridge.TransactOpts)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidator(&_BranchBridge.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidator(&_BranchBridge.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidators(&_BranchBridge.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidators(&_BranchBridge.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BranchBridge.Contract.RenounceOwnership(&_BranchBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BranchBridge.Contract.RenounceOwnership(&_BranchBridge.TransactOpts)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.ResetValidators(&_BranchBridge.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.ResetValidators(&_BranchBridge.TransactOpts, validators)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridge", remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, remoteChainID, indexes)
}

// SetERC20Factory is a paid mutator transaction binding the contract method 0xf516b879.
//
// Solidity: function setERC20Factory(address factoryCode) returns()
func (_BranchBridge *BranchBridgeTransactor) SetERC20Factory(opts *bind.TransactOpts, factoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setERC20Factory", factoryCode)
}

// SetERC20Factory is a paid mutator transaction binding the contract method 0xf516b879.
//
// Solidity: function setERC20Factory(address factoryCode) returns()
func (_BranchBridge *BranchBridgeSession) SetERC20Factory(factoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetERC20Factory(&_BranchBridge.TransactOpts, factoryCode)
}

// SetERC20Factory is a paid mutator transaction binding the contract method 0xf516b879.
//
// Solidity: function setERC20Factory(address factoryCode) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetERC20Factory(factoryCode common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetERC20Factory(&_BranchBridge.TransactOpts, factoryCode)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactor) SetFeeStation(opts *bind.TransactOpts, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setFeeStation", _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetFeeStation(&_BranchBridge.TransactOpts, _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetFeeStation(&_BranchBridge.TransactOpts, _bridgeFeeStation)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeTransactor) SetRewardWallet(opts *bind.TransactOpts, rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setRewardWallet", rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, rewardWallet_)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidator(&_BranchBridge.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidator(&_BranchBridge.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidators(&_BranchBridge.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidators(&_BranchBridge.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.TransferOwnership(&_BranchBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.TransferOwnership(&_BranchBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeSession) Unpause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Unpause(&_BranchBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Unpause(&_BranchBridge.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) UnpauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unpauseToken", remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnpauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnpauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnregisterToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnregisterToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.UpgradeToAndCall(&_BranchBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.UpgradeToAndCall(&_BranchBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeSession) Receive() (*types.Transaction, error) {
	return _BranchBridge.Contract.Receive(&_BranchBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BranchBridge.Contract.Receive(&_BranchBridge.TransactOpts)
}

// BranchBridgeBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the BranchBridge contract.
type BranchBridgeBridgeFeeChargedIterator struct {
	Event *BranchBridgeBridgeFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeBridgeFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFeeCharged represents a BridgeFeeCharged event raised by the BranchBridge contract.
type BranchBridgeBridgeFeeCharged struct {
	Index   *big.Int
	Token   common.Address
	Account common.Address
	GasFee  *big.Int
	ExFee   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeCharged is a free log retrieval operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*BranchBridgeBridgeFeeChargedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFeeChargedIterator{contract: _BranchBridge.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFeeCharged)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeFeeCharged is a log parse operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFeeCharged(log types.Log) (*BranchBridgeBridgeFeeCharged, error) {
	event := new(BranchBridgeBridgeFeeCharged)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeFinalizeRevertedIterator is returned from FilterBridgeFinalizeReverted and is used to iterate over the raw logs and unpacked data for BridgeFinalizeReverted events raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizeRevertedIterator struct {
	Event *BranchBridgeBridgeFinalizeReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFinalizeReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeBridgeFinalizeReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFinalizeReverted represents a BridgeFinalizeReverted event raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizeReverted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeReverted is a free log retrieval operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalizeReverted(opts *bind.FilterOpts, index []*big.Int) (*BranchBridgeBridgeFinalizeRevertedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizeRevertedIterator{contract: _BranchBridge.contract, event: "BridgeFinalizeReverted", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeReverted is a free log subscription operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalizeReverted(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalizeReverted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFinalizeReverted)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeFinalizeReverted is a log parse operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFinalizeReverted(log types.Log) (*BranchBridgeBridgeFinalizeReverted, error) {
	event := new(BranchBridgeBridgeFinalizeReverted)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizedIterator struct {
	Event *BranchBridgeBridgeFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeBridgeFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFinalized represents a BridgeFinalized event raised by the BranchBridge contract.
type BranchBridgeBridgeFinalized struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Token         common.Address
	To            common.Address
	Value         *big.Int
	Time          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (*BranchBridgeBridgeFinalizedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizedIterator{contract: _BranchBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalized, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFinalized)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeFinalized is a log parse operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFinalized(log types.Log) (*BranchBridgeBridgeFinalized, error) {
	event := new(BranchBridgeBridgeFinalized)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BranchBridge contract.
type BranchBridgeBridgeInitiatedIterator struct {
	Event *BranchBridgeBridgeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeBridgeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeInitiated represents a BridgeInitiated event raised by the BranchBridge contract.
type BranchBridgeBridgeInitiated struct {
	RemoteChainID *big.Int
	Index         *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Time          *big.Int
	Permit        bool
	ExtraData     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address indexed localToken, address remoteToken, address from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (*BranchBridgeBridgeInitiatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeInitiatedIterator{contract: _BranchBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address indexed localToken, address remoteToken, address from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeInitiated, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeInitiated)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeInitiated is a log parse operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address indexed localToken, address remoteToken, address from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeInitiated(log types.Log) (*BranchBridgeBridgeInitiated, error) {
	event := new(BranchBridgeBridgeInitiated)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeTokenFailedIterator is returned from FilterBridgeTokenFailed and is used to iterate over the raw logs and unpacked data for BridgeTokenFailed events raised by the BranchBridge contract.
type BranchBridgeBridgeTokenFailedIterator struct {
	Event *BranchBridgeBridgeTokenFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeBridgeTokenFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeTokenFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeBridgeTokenFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeBridgeTokenFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeTokenFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeTokenFailed represents a BridgeTokenFailed event raised by the BranchBridge contract.
type BranchBridgeBridgeTokenFailed struct {
	Permit  bool
	Success []bool
	Reason  []string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokenFailed is a free log retrieval operation binding the contract event 0xb8748b102128effc65c1bc368e51a5d18b212136f6f6097a5730a990abca2430.
//
// Solidity: event BridgeTokenFailed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeTokenFailed(opts *bind.FilterOpts, permit []bool) (*BranchBridgeBridgeTokenFailedIterator, error) {

	var permitRule []interface{}
	for _, permitItem := range permit {
		permitRule = append(permitRule, permitItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeTokenFailed", permitRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeTokenFailedIterator{contract: _BranchBridge.contract, event: "BridgeTokenFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeTokenFailed is a free log subscription operation binding the contract event 0xb8748b102128effc65c1bc368e51a5d18b212136f6f6097a5730a990abca2430.
//
// Solidity: event BridgeTokenFailed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeTokenFailed(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeTokenFailed, permit []bool) (event.Subscription, error) {

	var permitRule []interface{}
	for _, permitItem := range permit {
		permitRule = append(permitRule, permitItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeTokenFailed", permitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeTokenFailed)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeTokenFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeTokenFailed is a log parse operation binding the contract event 0xb8748b102128effc65c1bc368e51a5d18b212136f6f6097a5730a990abca2430.
//
// Solidity: event BridgeTokenFailed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeTokenFailed(log types.Log) (*BranchBridgeBridgeTokenFailed, error) {
	event := new(BranchBridgeBridgeTokenFailed)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeTokenFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeCrossMintableERC20FactorySetIterator is returned from FilterCrossMintableERC20FactorySet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20FactorySet events raised by the BranchBridge contract.
type BranchBridgeCrossMintableERC20FactorySetIterator struct {
	Event *BranchBridgeCrossMintableERC20FactorySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeCrossMintableERC20FactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeCrossMintableERC20FactorySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeCrossMintableERC20FactorySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeCrossMintableERC20FactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeCrossMintableERC20FactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeCrossMintableERC20FactorySet represents a CrossMintableERC20FactorySet event raised by the BranchBridge contract.
type BranchBridgeCrossMintableERC20FactorySet struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20FactorySet is a free log retrieval operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_BranchBridge *BranchBridgeFilterer) FilterCrossMintableERC20FactorySet(opts *bind.FilterOpts, factory []common.Address) (*BranchBridgeCrossMintableERC20FactorySetIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeCrossMintableERC20FactorySetIterator{contract: _BranchBridge.contract, event: "CrossMintableERC20FactorySet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20FactorySet is a free log subscription operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_BranchBridge *BranchBridgeFilterer) WatchCrossMintableERC20FactorySet(opts *bind.WatchOpts, sink chan<- *BranchBridgeCrossMintableERC20FactorySet, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeCrossMintableERC20FactorySet)
				if err := _BranchBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossMintableERC20FactorySet is a log parse operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_BranchBridge *BranchBridgeFilterer) ParseCrossMintableERC20FactorySet(log types.Log) (*BranchBridgeCrossMintableERC20FactorySet, error) {
	event := new(BranchBridgeCrossMintableERC20FactorySet)
	if err := _BranchBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BranchBridge contract.
type BranchBridgeEIP712DomainChangedIterator struct {
	Event *BranchBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the BranchBridge contract.
type BranchBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BranchBridge *BranchBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BranchBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeEIP712DomainChangedIterator{contract: _BranchBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BranchBridge *BranchBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BranchBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeEIP712DomainChanged)
				if err := _BranchBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BranchBridge *BranchBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BranchBridgeEIP712DomainChanged, error) {
	event := new(BranchBridgeEIP712DomainChanged)
	if err := _BranchBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeFeeStationSetIterator is returned from FilterFeeStationSet and is used to iterate over the raw logs and unpacked data for FeeStationSet events raised by the BranchBridge contract.
type BranchBridgeFeeStationSetIterator struct {
	Event *BranchBridgeFeeStationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeFeeStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeFeeStationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeFeeStationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeFeeStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeFeeStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeFeeStationSet represents a FeeStationSet event raised by the BranchBridge contract.
type BranchBridgeFeeStationSet struct {
	FeeStation common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeStationSet is a free log retrieval operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_BranchBridge *BranchBridgeFilterer) FilterFeeStationSet(opts *bind.FilterOpts, feeStation []common.Address) (*BranchBridgeFeeStationSetIterator, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeFeeStationSetIterator{contract: _BranchBridge.contract, event: "FeeStationSet", logs: logs, sub: sub}, nil
}

// WatchFeeStationSet is a free log subscription operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_BranchBridge *BranchBridgeFilterer) WatchFeeStationSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeFeeStationSet, feeStation []common.Address) (event.Subscription, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeFeeStationSet)
				if err := _BranchBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeStationSet is a log parse operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_BranchBridge *BranchBridgeFilterer) ParseFeeStationSet(log types.Log) (*BranchBridgeFeeStationSet, error) {
	event := new(BranchBridgeFeeStationSet)
	if err := _BranchBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BranchBridge contract.
type BranchBridgeInitializedIterator struct {
	Event *BranchBridgeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeInitialized represents a Initialized event raised by the BranchBridge contract.
type BranchBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BranchBridge *BranchBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BranchBridgeInitializedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeInitializedIterator{contract: _BranchBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BranchBridge *BranchBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BranchBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeInitialized)
				if err := _BranchBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BranchBridge *BranchBridgeFilterer) ParseInitialized(log types.Log) (*BranchBridgeInitialized, error) {
	event := new(BranchBridgeInitialized)
	if err := _BranchBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BranchBridge contract.
type BranchBridgeOwnershipTransferredIterator struct {
	Event *BranchBridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the BranchBridge contract.
type BranchBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BranchBridge *BranchBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BranchBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeOwnershipTransferredIterator{contract: _BranchBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BranchBridge *BranchBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BranchBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeOwnershipTransferred)
				if err := _BranchBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BranchBridge *BranchBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BranchBridgeOwnershipTransferred, error) {
	event := new(BranchBridgeOwnershipTransferred)
	if err := _BranchBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BranchBridge contract.
type BranchBridgePausedIterator struct {
	Event *BranchBridgePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgePaused represents a Paused event raised by the BranchBridge contract.
type BranchBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BranchBridge *BranchBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BranchBridgePausedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BranchBridgePausedIterator{contract: _BranchBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BranchBridge *BranchBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BranchBridgePaused) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgePaused)
				if err := _BranchBridge.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BranchBridge *BranchBridgeFilterer) ParsePaused(log types.Log) (*BranchBridgePaused, error) {
	event := new(BranchBridgePaused)
	if err := _BranchBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeRewardWalletSetIterator is returned from FilterRewardWalletSet and is used to iterate over the raw logs and unpacked data for RewardWalletSet events raised by the BranchBridge contract.
type BranchBridgeRewardWalletSetIterator struct {
	Event *BranchBridgeRewardWalletSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeRewardWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeRewardWalletSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeRewardWalletSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeRewardWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeRewardWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeRewardWalletSet represents a RewardWalletSet event raised by the BranchBridge contract.
type BranchBridgeRewardWalletSet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardWalletSet is a free log retrieval operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address indexed wallet)
func (_BranchBridge *BranchBridgeFilterer) FilterRewardWalletSet(opts *bind.FilterOpts, wallet []common.Address) (*BranchBridgeRewardWalletSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "RewardWalletSet", walletRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeRewardWalletSetIterator{contract: _BranchBridge.contract, event: "RewardWalletSet", logs: logs, sub: sub}, nil
}

// WatchRewardWalletSet is a free log subscription operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address indexed wallet)
func (_BranchBridge *BranchBridgeFilterer) WatchRewardWalletSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeRewardWalletSet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "RewardWalletSet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeRewardWalletSet)
				if err := _BranchBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardWalletSet is a log parse operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address indexed wallet)
func (_BranchBridge *BranchBridgeFilterer) ParseRewardWalletSet(log types.Log) (*BranchBridgeRewardWalletSet, error) {
	event := new(BranchBridgeRewardWalletSet)
	if err := _BranchBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BranchBridge contract.
type BranchBridgeThresholdChangedIterator struct {
	Event *BranchBridgeThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeThresholdChanged represents a ThresholdChanged event raised by the BranchBridge contract.
type BranchBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BranchBridge *BranchBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BranchBridgeThresholdChangedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeThresholdChangedIterator{contract: _BranchBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BranchBridge *BranchBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BranchBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeThresholdChanged)
				if err := _BranchBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdChanged is a log parse operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BranchBridge *BranchBridgeFilterer) ParseThresholdChanged(log types.Log) (*BranchBridgeThresholdChanged, error) {
	event := new(BranchBridgeThresholdChanged)
	if err := _BranchBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairPausedIterator is returned from FilterTokenPairPaused and is used to iterate over the raw logs and unpacked data for TokenPairPaused events raised by the BranchBridge contract.
type BranchBridgeTokenPairPausedIterator struct {
	Event *BranchBridgeTokenPairPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeTokenPairPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeTokenPairPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeTokenPairPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairPaused represents a TokenPairPaused event raised by the BranchBridge contract.
type BranchBridgeTokenPairPaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairPaused is a free log retrieval operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairPaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BranchBridgeTokenPairPausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairPausedIterator{contract: _BranchBridge.contract, event: "TokenPairPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairPaused is a free log subscription operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairPaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairPaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairPaused)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairPaused is a log parse operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairPaused(log types.Log) (*BranchBridgeTokenPairPaused, error) {
	event := new(BranchBridgeTokenPairPaused)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BranchBridge contract.
type BranchBridgeTokenPairRegisteredIterator struct {
	Event *BranchBridgeTokenPairRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeTokenPairRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the BranchBridge contract.
type BranchBridgeTokenPairRegistered struct {
	RemoteChainID   *big.Int
	IsOrigin        bool
	LocalToken      common.Address
	RemoteToken     common.Address
	LocalTokenRate  *big.Int
	RemoteTokenRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0x4336c2d0a5a91f2abb5c43ad295d612b03503ba19e54f7a9be8d9f249a242e3a.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BranchBridgeTokenPairRegisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairRegisteredIterator{contract: _BranchBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0x4336c2d0a5a91f2abb5c43ad295d612b03503ba19e54f7a9be8d9f249a242e3a.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairRegistered)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairRegistered is a log parse operation binding the contract event 0x4336c2d0a5a91f2abb5c43ad295d612b03503ba19e54f7a9be8d9f249a242e3a.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*BranchBridgeTokenPairRegistered, error) {
	event := new(BranchBridgeTokenPairRegistered)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairUnpausedIterator is returned from FilterTokenPairUnpaused and is used to iterate over the raw logs and unpacked data for TokenPairUnpaused events raised by the BranchBridge contract.
type BranchBridgeTokenPairUnpausedIterator struct {
	Event *BranchBridgeTokenPairUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeTokenPairUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeTokenPairUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeTokenPairUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairUnpaused represents a TokenPairUnpaused event raised by the BranchBridge contract.
type BranchBridgeTokenPairUnpaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnpaused is a free log retrieval operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairUnpaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BranchBridgeTokenPairUnpausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairUnpausedIterator{contract: _BranchBridge.contract, event: "TokenPairUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnpaused is a free log subscription operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairUnpaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairUnpaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairUnpaused)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairUnpaused is a log parse operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairUnpaused(log types.Log) (*BranchBridgeTokenPairUnpaused, error) {
	event := new(BranchBridgeTokenPairUnpaused)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the BranchBridge contract.
type BranchBridgeTokenPairUnregisteredIterator struct {
	Event *BranchBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeTokenPairUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the BranchBridge contract.
type BranchBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*BranchBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairUnregisteredIterator{contract: _BranchBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairUnregistered)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairUnregistered is a log parse operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*BranchBridgeTokenPairUnregistered, error) {
	event := new(BranchBridgeTokenPairUnregistered)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BranchBridge contract.
type BranchBridgeUnpausedIterator struct {
	Event *BranchBridgeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeUnpaused represents a Unpaused event raised by the BranchBridge contract.
type BranchBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BranchBridge *BranchBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BranchBridgeUnpausedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeUnpausedIterator{contract: _BranchBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BranchBridge *BranchBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeUnpaused)
				if err := _BranchBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BranchBridge *BranchBridgeFilterer) ParseUnpaused(log types.Log) (*BranchBridgeUnpaused, error) {
	event := new(BranchBridgeUnpaused)
	if err := _BranchBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BranchBridge contract.
type BranchBridgeUpgradedIterator struct {
	Event *BranchBridgeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeUpgraded represents a Upgraded event raised by the BranchBridge contract.
type BranchBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BranchBridge *BranchBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BranchBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeUpgradedIterator{contract: _BranchBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BranchBridge *BranchBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BranchBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeUpgraded)
				if err := _BranchBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BranchBridge *BranchBridgeFilterer) ParseUpgraded(log types.Log) (*BranchBridgeUpgraded, error) {
	event := new(BranchBridgeUpgraded)
	if err := _BranchBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the BranchBridge contract.
type BranchBridgeValidatorSetIterator struct {
	Event *BranchBridgeValidatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BranchBridgeValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeValidatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BranchBridgeValidatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BranchBridgeValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeValidatorSet represents a ValidatorSet event raised by the BranchBridge contract.
type BranchBridgeValidatorSet struct {
	Validator common.Address
	Status    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) FilterValidatorSet(opts *bind.FilterOpts, validator []common.Address, status []bool) (*BranchBridgeValidatorSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ValidatorSet", validatorRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeValidatorSetIterator{contract: _BranchBridge.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeValidatorSet, validator []common.Address, status []bool) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ValidatorSet", validatorRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeValidatorSet)
				if err := _BranchBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSet is a log parse operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) ParseValidatorSet(log types.Log) (*BranchBridgeValidatorSet, error) {
	event := new(BranchBridgeValidatorSet)
	if err := _BranchBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
