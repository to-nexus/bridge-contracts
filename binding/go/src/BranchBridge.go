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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allRevertedIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"bridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"_crossMintableERC20Factory\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"success\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"reason\",\"type\":\"string[]\"}],\"name\":\"BridgeTokenBatchProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20FactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RewardWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RegistryExistFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"RegistryInvalidRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroLocalRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroRemoteRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedSendValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidExFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidGasFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMinAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotExistFeeStation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
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
		"5187599d": "initialize(uint8,address)",
		"91cf6d3e": "initializedAt()",
		"facd743b": "isValidator(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"d2ff130d": "pauseToken(uint256,address)",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
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
		"1a1aebbb": "setCrossMintableERC20Factory(address)",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615f726100f95f395f81816134f50152818161351e01526136e20152615f725ff3fe60806040526004361061032a575f3560e01c806384d58d42116101a3578063b7f3358d116100f2578063d80e395011610092578063f45096431161006d578063f450964314610a36578063f698da2514610a55578063facd743b14610a69578063fb75b2c714610a88575f5ffd5b8063d80e3950146109d7578063f2fde38b146109f6578063f30589c314610a15575f5ffd5b8063d2ff130d116100cd578063d2ff130d14610972578063d5717fc514610991578063d605665b146109b0578063d7c82f32146109c3575f5ffd5b8063b7f3358d14610913578063cbae595814610932578063cf56118e14610951575f5ffd5b80639300c9261161015d578063ad3cb1cc11610138578063ad3cb1cc14610876578063ae6893f8146108a6578063ae766389146108c5578063aed1d403146108ff575f5ffd5b80639300c92614610817578063952a95de1461083657806396ce079514610862575f5ffd5b806384d58d421461076357806388d67d6d146107825780638da5cb5b146107955780638f517c17146107d15780639118b5eb146107f057806391cf6d3e14610803575f5ffd5b80635187599d116102795780635fd262de11610219578063814914b5116101f4578063814914b51461061a5780638415a385146106fc5780638456cb591461072857806384b0196e1461073c575f5ffd5b80635fd262de146105d45780637101fcd3146105e7578063715018a614610606575f5ffd5b8063578478931161025457806357847893146105475780635958621e146105665780635b605f5c146105855780635c975abb146105b1575f5ffd5b80635187599d146104e757806352d1902d1461050657806354db012614610528575f5ffd5b80633960e787116102e457806342cde4e8116102bf57806342cde4e81461048257806347666cb1146104a25780634d5d0056146104c15780634f1ef286146104d4575f5ffd5b80633960e787146104305780633f4ba83a1461044f57806340a141ff14610463575f5ffd5b8063030372c3146103555780631003e37f146103895780631327d3d8146103c05780631938e0f2146103df5780631a1aebbb146103f25780631d40f0d814610411575f5ffd5b3661035157345f0361034f576040516365d14ce560e11b815260040160405180910390fd5b005b5f5ffd5b348015610360575f5ffd5b5061037461036f366004614a4b565b610aa5565b60405190151581526020015b60405180910390f35b348015610394575f5ffd5b506103a86103a3366004614b7b565b610ae9565b6040516001600160a01b039091168152602001610380565b3480156103cb575f5ffd5b5061034f6103da366004614c04565b610c00565b6103746103ed366004614cfe565b610c0e565b3480156103fd575f5ffd5b5061034f61040c366004614c04565b610f12565b34801561041c575f5ffd5b5061034f61042b366004614db7565b610fc2565b34801561043b575f5ffd5b5061037461044a366004614e52565b610ffc565b34801561045a575f5ffd5b5061034f6110f3565b34801561046e575f5ffd5b5061034f61047d366004614c04565b611105565b34801561048d575f5ffd5b505f5460405160ff9091168152602001610380565b3480156104ad575f5ffd5b506063546103a8906001600160a01b031681565b6103746104cf366004614eaf565b61110f565b61034f6104e2366004614f51565b6113c2565b3480156104f2575f5ffd5b5061034f610501366004614fb0565b6113dd565b348015610511575f5ffd5b5061051a6114eb565b604051908152602001610380565b348015610533575f5ffd5b5061034f610542366004614c04565b611507565b348015610552575f5ffd5b5061034f610561366004614ff2565b61157f565b348015610571575f5ffd5b5061034f610580366004614c04565b61159d565b348015610590575f5ffd5b506105a461059f366004615053565b611615565b60405161038091906150bd565b3480156105bc575f5ffd5b505f516020615efd5f395f51905f525460ff16610374565b6103746105e236600461510a565b611793565b3480156105f2575f5ffd5b5061034f610601366004614db7565b61189a565b348015610611575f5ffd5b5061034f6118b0565b348015610625575f5ffd5b506106ef610634366004615192565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f8281526037602090815260408083206001600160a01b03808616855290835292819020815160e08101835281548516815260018201549094169284019290925260028201549083015260038101546060830152600481015460ff8082161515608085015261010090910416151560a08301526005015460c082015292915050565b60405161038091906151b5565b348015610707575f5ffd5b5061071b610716366004614e52565b6118c1565b60405161038091906151f1565b348015610733575f5ffd5b5061034f61196f565b348015610747575f5ffd5b5061075061197f565b604051610380979695949392919061523d565b34801561076e575f5ffd5b5061034f61077d366004615192565b611a28565b610374610790366004615376565b611b27565b3480156107a0575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166103a8565b3480156107dc575f5ffd5b506032546103a8906001600160a01b031681565b61034f6107fe3660046154b2565b611bc2565b34801561080e575f5ffd5b5060655461051a565b348015610822575f5ffd5b5061034f610831366004614db7565b611f2c565b348015610841575f5ffd5b50610855610850366004614e52565b611f63565b60405161038091906154f0565b34801561086d575f5ffd5b5061051a6120aa565b348015610881575f5ffd5b5061071b604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156108b1575f5ffd5b5061051a6108c0366004615053565b61211a565b3480156108d0575f5ffd5b506108e46108df366004615554565b612136565b60408051938452602084019290925290820152606001610380565b34801561090a575f5ffd5b5061051a6121c5565b34801561091e575f5ffd5b5061034f61092d366004615589565b6121d0565b34801561093d575f5ffd5b506103a861094c366004615053565b612220565b34801561095c575f5ffd5b5061096561222c565b60405161038091906155a2565b34801561097d575f5ffd5b5061034f61098c366004615192565b612238565b34801561099c575f5ffd5b5061051a6109ab366004615053565b61233c565b61034f6109be3660046155b4565b612358565b3480156109ce575f5ffd5b5061034f6126fe565b3480156109e2575f5ffd5b506109656109f1366004615053565b61276b565b348015610a01575f5ffd5b5061034f610a10366004614c04565b612787565b348015610a20575f5ffd5b50610a296127c1565b604051610380919061564d565b348015610a41575f5ffd5b5061034f610a50366004615192565b6127cd565b348015610a60575f5ffd5b5061051a6128a3565b348015610a74575f5ffd5b50610374610a83366004614c04565b6128ac565b348015610a93575f5ffd5b506064546001600160a01b03166103a8565b5f805b8251811015610add57610ad484848381518110610ac757610ac761568d565b6020026020010151610ffc565b50600101610aa8565b50600190505b92915050565b5f610af26128b8565b6032546001600160a01b0316610b1b576040516315aeca0d60e11b815260040160405180910390fd5b6032546040516bffffffffffffffffffffffff19606089901b1660208201526001600160a01b0390911690634804a041906034016040516020818303038152906040528051906020012085604051602001610b7691906156b8565b60405160208183030381529060405286866040518563ffffffff1660e01b8152600401610ba694939291906156d9565b6020604051808303815f875af1158015610bc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be69190615718565b9050610bf6875f83898989612913565b9695505050505050565b610c0b816001612c18565b50565b5f610c17612cd5565b8435610c296060870160408801614c04565b5f828152603660205260409020610c409082612d05565b8190610c705760405163153096f360e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526037602090815260408083206001600160a01b03851684529091529020600401548190610100900460ff1615610cca576040516338384f6f60e11b81526001600160a01b039091166004820152602401610c67565b50610cd3612d29565b5f348015610cfd576040516329e2b03f60e21b815260048101929092526024820152604401610c67565b50610d0a9050873561233c565b602088013514610d1a883561233c565b88602001359091610d4757604051631351db4160e31b815260048101929092526024820152604401610c67565b50610dd990507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020890135610d8360608b0160408c01614c04565b610d9360808c0160608d01614c04565b60808c0135610da560a08e018e615733565b604051602001610dbb979695949392919061579d565b60405160208183030381529060405280519060200120878787612d60565b610df787355f90815260356020526040902060020180546001019055565b5f80610e288935610e0e60608c0160408d01614c04565b610e1e60808d0160608e01614c04565b8c60800135612f59565b915091508115610eb357610e4260808a0160608b01614c04565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610e8360608e0160408f01614c04565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610eec565b610ebd89826130aa565b60405160208a0135907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610f0860015f516020615f1d5f395f51905f5255565b5050949350505050565b610f1a6128b8565b6032546001600160a01b03168015610f5157604051639ad61dbd60e01b81526001600160a01b039091166004820152602401610c67565b506001600160a01b038116610f7957604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517f18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9905f90a250565b5f5b8151811015610ff857610ff0828281518110610fe257610fe261568d565b60200260200101515f612c18565b600101610fc4565b5050565b5f611005612cd5565b61100d612d29565b5f6110188484611f63565b90505f5f61103486846040015185606001518660800151612f59565b915091508181906110585760405162461bcd60e51b8152600401610c6791906151f1565b506110638686613140565b82606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b86604001518760800151426040516110ce939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a460019350505050610ae360015f516020615f1d5f395f51905f5255565b6110fb6128b8565b6111036131e4565b565b610c0b815f612c18565b5f611118612cd5565b5f8a81526036602052604090208a908a906111339082612d05565b819061115e5760405163153096f360e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b03851684529091529020600401548190610100900460ff16156111b8576040516338384f6f60e11b81526001600160a01b039091166004820152602401610c67565b506111c1612d29565b6111ce8c8c8b8b8b61323d565b9098509650866111de898b6157fe565b6111e891906157fe565b60408501351015876111fa8a8c6157fe565b61120491906157fe565b85604001359091611231576040516365efbabf60e11b815260048101929092526024820152604401610c67565b505f90506112426020860186614c04565b6001600160a01b031663d505accf6112606040880160208901614c04565b30604089013560608a013561127b60a08c0160808d01615589565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a088013560c482015260c088013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031660e09490941b9390931790925292505f915061130290870187614c04565b90505f5f60205f8551602087015f875af180611323576040513d5f823e3d81fd5b50505f513d915081156113395780600114611353565b6113466020890189614c04565b6001600160a01b03163b15155b61137057604051631c92cad760e01b815260040160405180910390fd5b505050506113998c8c86602001602081019061138c9190614c04565b8d8d8d8d60018e8e613374565b600192506113b360015f516020615f1d5f395f51905f5255565b50509998505050505050505050565b6113ca6134ea565b6113d38261358e565b610ff88282613596565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156114215750825b90505f826001600160401b0316600114801561143c5750303b155b90508115801561144a575080155b156114685760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561149257845460ff60401b1916600160401b1785555b61149c8787613657565b83156114e257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f6114f46136d7565b505f516020615edd5f395f51905f525b90565b61150f6128b8565b6001600160a01b03811661153657604051630fb9363360e41b815260040160405180910390fd5b606380546001600160a01b0319166001600160a01b0383169081179091556040517f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444905f90a250565b6115876128b8565b611595868686868686612913565b505050505050565b6115a56128b8565b6001600160a01b0381166115cc57604051630fb9363360e41b815260040160405180910390fd5b606480546001600160a01b0319166001600160a01b0383169081179091556040517f5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b905f90a250565b5f8181526036602052604081206060919061162f90613720565b90505f81516001600160401b0381111561164b5761164b6149e9565b6040519080825280602002602001820160405280156116b057816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f199092019101816116695790505b5090505f5b825181101561178b5760375f8681526020019081526020015f205f8483815181106116e2576116e261568d565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549094169184019190915260028101549183019190915260038101546060830152600481015460ff8082161515608085015261010090910416151560a08301526005015460c082015282518390839081106117785761177861568d565b60209081029190910101526001016116b5565b509392505050565b5f61179c612cd5565b5f898152603660205260409020899089906117b79082612d05565b81906117e25760405163153096f360e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b03851684529091529020600401548190610100900460ff161561183c576040516338384f6f60e11b81526001600160a01b039091166004820152602401610c67565b50611845612d29565b5f5f6118548d8d8c8c8c61323d565b915091506118708d8d6118643390565b8e8e87875f8f8f613374565b60019450505061188c60015f516020615f1d5f395f51905f5255565b505098975050505050505050565b6118a761042b6001613720565b610c0b81611f2c565b6118b86128b8565b6111035f61372c565b5f82815260356020908152604080832084845260040190915290208054606091906118eb90615811565b80601f016020809104026020016040519081016040528092919081815260200182805461191790615811565b80156119625780601f1061193957610100808354040283529160200191611962565b820191905f5260205f20905b81548152906001019060200180831161194557829003601f168201915b5050505050905092915050565b6119776128b8565b61110361379c565b5f60608082808083815f516020615ebd5f395f51905f5280549091501580156119aa57506001810154155b6119ee5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610c67565b6119f66137e4565b6119fe6138a4565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b611a306128b8565b5f828152603660205260409020611a479082612d05565b8190611a725760405163153096f360e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b03851684529091529020600401548190610100900460ff16611acb57604051636c508f9f60e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b0385168085529252808320600401805461ff001916905551909184917fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e99190a35050565b5f805b85811015611bb557611bac878783818110611b4757611b4761568d565b9050602002810190611b599190615849565b868381518110611b6b57611b6b61568d565b6020026020010151868481518110611b8557611b8561568d565b6020026020010151868581518110611b9f57611b9f61568d565b6020026020010151610c0e565b50600101611b2a565b5060019695505050505050565b5f816001600160401b03811115611bdb57611bdb6149e9565b604051908082528060200260200182016040528015611c04578160200160208202803683370190505b5090505f826001600160401b03811115611c2057611c206149e9565b604051908082528060200260200182016040528015611c5357816020015b6060815260200190600190039081611c3e5790505b5090505f5b83811015611ee95730635fd262de868684818110611c7857611c7861568d565b9050602002810190611c8a9190615867565b35878785818110611c9d57611c9d61568d565b9050602002810190611caf9190615867565b611cc0906040810190602001614c04565b888886818110611cd257611cd261568d565b9050602002810190611ce49190615867565b611cf5906060810190604001614c04565b898987818110611d0757611d0761568d565b9050602002810190611d199190615867565b606001358a8a88818110611d2f57611d2f61568d565b9050602002810190611d419190615867565b608001358b8b89818110611d5757611d5761568d565b9050602002810190611d699190615867565b60a001358c8c8a818110611d7f57611d7f61568d565b9050602002810190611d919190615867565b611d9f9060c0810190615733565b6040518963ffffffff1660e01b8152600401611dc298979695949392919061587b565b6020604051808303815f875af1925050508015611dfc575060408051601f3d908101601f19168201909252611df9918101906158c2565b60015b611ebb57611e086158dd565b806308c379a003611e4b5750611e1c6158f5565b80611e275750611e83565b80838381518110611e3a57611e3a61568d565b602002602001018190525050611ee1565b634e487b7103611e8357611e5d61596f565b90611e685750611e83565b611e71816138e2565b838381518110611e3a57611e3a61568d565b3d808015611eac576040519150601f19603f3d011682016040523d82523d5f602084013e611eb1565b606091505b50611e718161391e565b506001838281518110611ed057611ed061568d565b911515602092830291909101909101525b600101611c58565b505f15157f9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b288383604051611f1e92919061598c565b60405180910390a250505050565b5f5b8151811015610ff857611f5b828281518110611f4c57611f4c61568d565b60200260200101516001612c18565b600101611f2e565b611fa96040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f8381526035602090815260408083208584526003908101835292819020815160c0810183528154815260018201549381019390935260028101546001600160a01b0390811692840192909252928301541660608201526004820154608082015260058201805491929160a08401919061202290615811565b80601f016020809104026020016040519081016040528092919081815260200182805461204e90615811565b80156120995780601f1061207057610100808354040283529160200191612099565b820191905f5260205f20905b81548152906001019060200180831161207c57829003601f168201915b505050505081525050905092915050565b606354604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa1580156120f1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121159190615a2e565b905090565b5f818152603560205260408120600190810154610ae3916157fe565b60635460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa158015612192573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121b69190615a45565b91989097509095509350505050565b5f6121156001613931565b6121d86128b8565b5f805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f610ae360018361393a565b60606121156033613720565b6122406128b8565b5f8281526036602052604090206122579082612d05565b81906122825760405163153096f360e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b03851684529091529020600401548190610100900460ff16156122dc576040516338384f6f60e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b0385168085529252808320600401805461ff00191661010017905551909184917ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a39190a35050565b5f81815260356020526040812060020154610ae39060016157fe565b8281146123785760405163214485c960e01b815260040160405180910390fd5b5f836001600160401b03811115612391576123916149e9565b6040519080825280602002602001820160405280156123ba578160200160208202803683370190505b5090505f846001600160401b038111156123d6576123d66149e9565b60405190808252806020026020018201604052801561240957816020015b60608152602001906001900390816123f45790505b5090505f5b858110156126b85730634d5d005688888481811061242e5761242e61568d565b90506020028101906124409190615867565b358989858181106124535761245361568d565b90506020028101906124659190615867565b612476906040810190602001614c04565b8a8a868181106124885761248861568d565b905060200281019061249a9190615867565b6124ab906060810190604001614c04565b8b8b878181106124bd576124bd61568d565b90506020028101906124cf9190615867565b606001358c8c888181106124e5576124e561568d565b90506020028101906124f79190615867565b608001358d8d8981811061250d5761250d61568d565b905060200281019061251f9190615867565b60a001358e8e8a8181106125355761253561568d565b90506020028101906125479190615867565b6125559060c0810190615733565b8e8e8c8181106125675761256761568d565b905060e002016040518a63ffffffff1660e01b815260040161259199989796959493929190615a70565b6020604051808303815f875af19250505080156125cb575060408051601f3d908101601f191682019092526125c8918101906158c2565b60015b61268a576125d76158dd565b806308c379a00361261a57506125eb6158f5565b806125f65750612652565b808383815181106126095761260961568d565b6020026020010181905250506126b0565b634e487b71036126525761262c61596f565b906126375750612652565b612640816138e2565b8383815181106126095761260961568d565b3d80801561267b576040519150601f19603f3d011682016040523d82523d5f602084013e612680565b606091505b506126408161391e565b50600183828151811061269f5761269f61568d565b911515602092830291909101909101525b60010161240e565b50600115157f9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b2883836040516126ee92919061598c565b60405180910390a2505050505050565b6127066128b8565b6063546001600160a01b031661272f57604051637bda535160e01b815260040160405180910390fd5b606380546001600160a01b03191690556040515f907f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444908290a2565b5f818152603560205260409020606090610ae390600501613720565b61278f6128b8565b6001600160a01b0381166127b857604051631e4fbdf760e01b81525f6004820152602401610c67565b610c0b8161372c565b60606121156001613720565b6127d56128b8565b5f8281526036602052604090206127ec9082613945565b81906128175760405163153096f360e11b81526001600160a01b039091166004820152602401610c67565b505f8281526037602090815260408083206001600160a01b038516808552925280832080546001600160a01b031990811682556001820180549091169055600281018490556003810184905560048101805461ffff1916905560050183905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f612115613959565b5f610ae3600183612d05565b336128ea7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146111035760405163118cdaa760e01b8152336004820152602401610c67565b5f86815260356020526040902061292b603388613962565b15612934578681555b6001600160a01b03851661295b57604051636ca1fdd760e01b815260040160405180910390fd5b5f878152603660205260409020612972908661396d565b859061299d576040516311dd05f360e31b81526001600160a01b039091166004820152602401610c67565b50818314612a7b57825f036129c55760405163f0006e8f60e01b815260040160405180910390fd5b815f036129e557604051639a5ea84b60e01b815260040160405180910390fd5b81831115612a3657816001148015612a055750612a03600a84615b53565b155b83839091612a2f57604051635c66012760e01b815260048101929092526024820152604401610c67565b5050612a7b565b826001148015612a4e5750612a4c600a83615b53565b155b83839091612a7857604051635c66012760e01b815260048101929092526024820152604401610c67565b50505b6040518060e00160405280866001600160a01b03168152602001856001600160a01b0316815260200184815260200183815260200187151581526020015f151581526020015f81525060375f8981526020019081526020015f205f876001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301556080820151816004015f6101000a81548160ff02191690831515021790555060a08201518160040160016101000a81548160ff02191690831515021790555060c08201518160050155905050836001600160a01b0316856001600160a01b0316887f4336c2d0a5a91f2abb5c43ad295d612b03503ba19e54f7a9be8d9f249a242e3a898787604051612c079392919092151583526020830191909152604082015260600190565b60405180910390a450505050505050565b612c206128b8565b8015612c6257612c3160018361396d565b8290612c5c576040516329a04e7760e21b81526001600160a01b039091166004820152602401610c67565b50612c9a565b612c6d600183613945565b8290612c985760405163fdbc594760e01b81526001600160a01b039091166004820152602401610c67565b505b604051811515906001600160a01b038416907f763b63b30e91c843bb39e4379603697003d3b7c1f192619cd782fa33bdc44396905f90a35050565b5f516020615efd5f395f51905f525460ff16156111035760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f90815260018301602052604081205415155b9392505050565b5f516020615f1d5f395f51905f52805460011901612d5a57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b8251825181148015612d725750815181145b835183518392612da6576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610c67565b50505f5482915060ff16811015612dd357604051632fcba65760e11b8152600401610c6791815260200190565b505f80826001600160401b03811115612dee57612dee6149e9565b604051908082528060200260200182016040528015612e17578160200160208202803683370190505b5090505f5b83811015612f24575f612e87888381518110612e3a57612e3a61568d565b6020026020010151888481518110612e5457612e5461568d565b6020026020010151888581518110612e6e57612e6e61568d565b6020026020010151612e7f8d613981565b9291906139ad565b9050612e92816128ac565b8190612ebd5760405163845a09e760e01b81526001600160a01b039091166004820152602401610c67565b505f805b8451811015612f0d57848181518110612edc57612edc61568d565b60200260200101516001600160a01b0316836001600160a01b031603612f055760019150612f0d565b600101612ec1565b5080612f1a578460010194505b5050600101612e1c565b505f54829060ff16811015612f4f57604051632fcba65760e11b8152600401610c6791815260200190565b5050505050505050565b5f8481526037602090815260408083206001600160a01b038088168552908352818420825160e08101845281548316815260018201549092169382019390935260028301549181018290526003830154606082810191909152600484015460ff8082161515608085015261010090910416151560a083015260059093015460c08201529015613016578060400151600114613004576040810151612ffd9085615b66565b9350613016565b60608101516130139085615b7d565b93505b5f196001600160a01b0387160161304b5761303185856139d9565b6001925060405180602001604052805f815250915061309f565b831561309f575f8781526037602090815260408083206001600160a01b038a16845290915290206004015460ff16156130945761308a87878787613a6f565b92509250506130a1565b61308a868686613bec565b505b94509492505050565b81355f908152603560209081526040909120908301356130cd6005830182613962565b81906130ef576040516307a5c91d60e31b8152600401610c6791815260200190565b505f8181526003830160205260409020849061310b8282615cad565b50505f81815260048301602052604090206131268482615d4d565b5050505050565b60015f516020615f1d5f395f51905f5255565b5f82815260356020526040902061315a6005820183613d56565b829061317c57604051637f11bea960e01b8152600401610c6791815260200190565b505f82815260048201602052604081206131959161499f565b5f828152600380830160205260408220828155600181018390556002810180546001600160a01b031990811690915591810180549092169091556004810182905590613126600583018261499f565b6131ec613d61565b5f516020615efd5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001612215565b6063545f9081906001600160a01b031661325b57505f90508061336a565b60635460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa1580156132b1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132d59190615a45565b909450925090508086818110156133085760405163428d097560e01b815260048101929092526024820152604401610c67565b508390508581811015613337576040516307cd920960e51b815260048101929092526024820152604401610c67565b508290508481811015613366576040516379a9e62b60e11b815260048101929092526024820152604401610c67565b5050505b9550959350505050565b5f8a81526037602090815260408083206001600160a01b03808e16855290835292819020815160e08101835281548516815260018201549094169284019290925260028201549083015260038101546060830152600481015460ff8082161515608085015261010090910416151560a08301526005015460c08201526134068b8b8b8a6134018a8c6157fe565b613d90565b5f5f6134118d61211a565b9150826020015190508b6001600160a01b0316828e7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7848f8f8f428e8e8e604051613463989796959493929190615e07565b60405180910390a48a6001600160a01b03168c6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8b8b6040516134ba929190918252602082015260400190565b60405180910390a45f8d815260356020526040902060019081018054909101905550505050505050505050505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061357057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166135645f516020615edd5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156111035760405163703e46dd60e11b815260040160405180910390fd5b610c0b6128b8565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156135f0575060408051601f3d908101601f191682019092526135ed91810190615a2e565b60015b61361857604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610c67565b5f516020615edd5f395f51905f52811461364857604051632a87526960e21b815260048101829052602401610c67565b6136528383613fcd565b505050565b61365f614022565b6136683361406b565b61367061407c565b613678614084565b613680614094565b613689826140a4565b6001600160a01b0381166136b057604051630fb9363360e41b815260040160405180910390fd5b606480546001600160a01b0319166001600160a01b03929092169190911790555043606555565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146111035760405163703e46dd60e11b815260040160405180910390fd5b60605f612d2283614109565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6137a4612cd5565b5f516020615efd5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613225565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615ebd5f395f51905f529161382290615811565b80601f016020809104026020016040519081016040528092919081815260200182805461384e90615811565b80156138995780601f1061387057610100808354040283529160200191613899565b820191905f5260205f20905b81548152906001019060200180831161387c57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615ebd5f395f51905f529161382290615811565b6040516b02830b734b19031b7b2329d160a51b6020820152602c8101829052606090604c015b6040516020818303038152906040529050919050565b6060816040516020016139089190615e51565b5f610ae3825490565b5f612d228383614162565b5f612d22836001600160a01b038416614188565b5f61211561426b565b5f612d2283836142de565b5f612d22836001600160a01b0384166142de565b5f610ae361398d613959565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6139bd8888888861432a565b9250925092506139cd82826143f2565b50909695505050505050565b804747821115613a05576040516371e4d07760e11b815260048101929092526024820152604401610c67565b50505f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114613a51576040519150601f19603f3d011682016040523d82523d5f602084013e613a56565b606091505b509150915081613a6957613a69816144aa565b50505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613ade575060408051601f3d908101601f19168201909252613adb918101906158c2565b60015b613b7f57613aea6158dd565b806308c379a003613b135750613afe6158f5565b80613b095750613b44565b5f925090506130a1565b634e487b7103613b4457613b2561596f565b90613b305750613b44565b5f9250613b3c816138e2565b9150506130a1565b3d808015613b6d576040519150601f19603f3d011682016040523d82523d5f602084013e613b72565b606091505b505f9250613b3c8161391e565b8015613baa576001925060405180602001604052805f8152509150613ba58787866144d3565b61309f565b505060408051808201909152601f81527f5374616e646172644272696467653a207472616e73666572206661696c65640060208201525f969095509350505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613c5b575060408051601f3d908101601f19168201909252613c58918101906158c2565b60015b613cf357613c676158dd565b806308c379a003613c8d5750613c7b6158f5565b80613c865750613cbb565b9050613d4e565b634e487b7103613cbb57613c9f61596f565b90613caa5750613cbb565b613cb3816138e2565b915050613d4e565b3d808015613ce4576040519150601f19603f3d011682016040523d82523d5f602084013e613ce9565b606091505b50613cb38161391e565b8015613d13576001925060405180602001604052805f8152509150613d4c565b6040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b5f612d228383614188565b5f516020615efd5f395f51905f525460ff1661110357604051638dfc202b60e01b815260040160405180910390fd5b5f8581526037602090815260408083206001600160a01b03881684529091529020600201546001811115613dfe57613dc88184615b53565b8590849015613dfb576040516348d53e0760e01b81526001600160a01b0390921660048301526024820152604401610c67565b50505b5f196001600160a01b03861601613e7157613e1982846157fe565b3414613e2583856157fe565b349091613e4e576040516329e2b03f60e21b815260048101929092526024820152604401610c67565b50508115613e6c57606454613e6c906001600160a01b0316836139d9565b611595565b5f348015613e9b576040516329e2b03f60e21b815260048101929092526024820152604401610c67565b50613ebf90508430613ead85876157fe565b6001600160a01b038916929190614575565b8115613edf57606454613edf906001600160a01b038781169116846145dc565b5f8681526037602090815260408083206001600160a01b038916845290915290206004015460ff1615613f1757613e6c86868561460d565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af1158015613f61573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f8591906158c2565b858585909192613fc25760405163601bc95b60e11b81526001600160a01b0393841660048201529290911660248301526044820152606401610c67565b505050505050505050565b613fd68261464b565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561401a5761365282826146ae565b610ff8614720565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661110357604051631afcd79f60e31b815260040160405180910390fd5b614073614022565b610c0b8161473f565b611103614022565b61408c614022565b611103614747565b61409c614022565b611103614767565b6140ac614022565b6140f4604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061476f565b5f805460ff191660ff92909216919091179055565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561415657602002820191905f5260205f20905b815481526020019060010190808311614142575b50505050509050919050565b5f825f0182815481106141775761417761568d565b905f5260205f200154905092915050565b5f8181526001830160205260408120548015614262575f6141aa600183615e76565b85549091505f906141bd90600190615e76565b905080821461421c575f865f0182815481106141db576141db61568d565b905f5260205f200154905080875f0184815481106141fb576141fb61568d565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061422d5761422d615e89565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610ae3565b5f915050610ae3565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f614295614781565b61429d6147e9565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f81815260018301602052604081205461432357508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610ae3565b505f610ae3565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561436357505f915060039050826143e8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156143b4573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166143df57505f9250600191508290506143e8565b92505f91508190505b9450945094915050565b5f82600381111561440557614405615e9d565b0361440e575050565b600182600381111561442257614422615e9d565b036144405760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561445457614454615e9d565b036144755760405163fce698f760e01b815260048101829052602401610c67565b600382600381111561448957614489615e9d565b03610ff8576040516335e2f38360e21b815260048101829052602401610c67565b8051156144ba5780518082602001fd5b60405163fd5478bd60e01b815260040160405180910390fd5b5f8381526037602090815260408083206001600160a01b03861684529091528120600501548190614504908461482b565b90925090508484848461454357604051634252ea9360e01b815260048101939093526001600160a01b0390911660248301526044820152606401610c67565b5050505f9485526037602090815260408087206001600160a01b03909616875294905292909320600501919091555050565b6040516001600160a01b038481166024830152838116604483015260648201839052613a699186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061484f565b6040516001600160a01b0383811660248301526044820183905261365291859182169063a9059cbb906064016145aa565b5f8381526037602090815260408083206001600160a01b0386168452909152812060050180548392906146419084906157fe565b9091555050505050565b806001600160a01b03163b5f0361468057604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610c67565b5f516020615edd5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516146ca9190615eb1565b5f60405180830381855af49150503d805f8114614702576040519150601f19603f3d011682016040523d82523d5f602084013e614707565b606091505b50915091506147178583836148bb565b95945050505050565b34156111035760405163b398979f60e01b815260040160405180910390fd5b61278f614022565b61474f614022565b5f516020615efd5f395f51905f52805460ff19169055565b61312d614022565b614777614022565b610ff88282614917565b5f5f516020615ebd5f395f51905f52816147996137e4565b8051909150156147b157805160209091012092915050565b815480156147c0579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615ebd5f395f51905f52816148016138a4565b80519091501561481957805160209091012092915050565b600182015480156147c0579392505050565b5f5f8383111561483f57505f905080614848565b50600190508183035b9250929050565b5f5f60205f8451602086015f885af18061486e576040513d5f823e3d81fd5b50505f513d91508115614885578060011415614892565b6001600160a01b0384163b155b15613a6957604051635274afe760e01b81526001600160a01b0385166004820152602401610c67565b6060826148d0576148cb82614976565b612d22565b81511580156148e757506001600160a01b0384163b155b1561491057604051639996b31560e01b81526001600160a01b0385166004820152602401610c67565b5092915050565b61491f614022565b5f516020615ebd5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026149588482615d4d565b50600381016149678382615d4d565b505f8082556001909101555050565b8051156149865780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5080546149ab90615811565b5f825580601f106149ba575050565b601f0160209004905f5260205f2090810190610c0b91905b808211156149e5575f81556001016149d2565b5090565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f191681016001600160401b0381118282101715614a2257614a226149e9565b6040525050565b5f6001600160401b03821115614a4157614a416149e9565b5060051b60200190565b5f5f60408385031215614a5c575f5ffd5b8235915060208301356001600160401b03811115614a78575f5ffd5b8301601f81018513614a88575f5ffd5b8035614a9381614a29565b604051614aa082826149fd565b80915082815260208101915060208360051b850101925087831115614ac3575f5ffd5b6020840193505b82841015614ae5578335825260209384019390910190614aca565b809450505050509250929050565b6001600160a01b0381168114610c0b575f5ffd5b5f5f6001600160401b03841115614b2057614b206149e9565b50604051601f8401601f191660200190614b3a82826149fd565b809250848152858585011115614b4e575f5ffd5b848460208301375f6020868301015250509392505050565b803560ff81168114614b76575f5ffd5b919050565b5f5f5f5f5f5f60c08789031215614b90575f5ffd5b863595506020870135614ba281614af3565b9450604087013593506060870135925060808701356001600160401b03811115614bca575f5ffd5b8701601f81018913614bda575f5ffd5b614be989823560208401614b07565b925050614bf860a08801614b66565b90509295509295509295565b5f60208284031215614c14575f5ffd5b8135612d2281614af3565b5f82601f830112614c2e575f5ffd5b8135614c3981614a29565b604051614c4682826149fd565b80915082815260208101915060208360051b860101925085831115614c69575f5ffd5b602085015b83811015614c8d57614c7f81614b66565b835260209283019201614c6e565b5095945050505050565b5f82601f830112614ca6575f5ffd5b8135614cb181614a29565b604051614cbe82826149fd565b80915082815260208101915060208360051b860101925085831115614ce1575f5ffd5b602085015b83811015614c8d578035835260209283019201614ce6565b5f5f5f5f60808587031215614d11575f5ffd5b84356001600160401b03811115614d26575f5ffd5b850160c08188031215614d37575f5ffd5b935060208501356001600160401b03811115614d51575f5ffd5b614d5d87828801614c1f565b93505060408501356001600160401b03811115614d78575f5ffd5b614d8487828801614c97565b92505060608501356001600160401b03811115614d9f575f5ffd5b614dab87828801614c97565b91505092959194509250565b5f60208284031215614dc7575f5ffd5b81356001600160401b03811115614ddc575f5ffd5b8201601f81018413614dec575f5ffd5b8035614df781614a29565b604051614e0482826149fd565b80915082815260208101915060208360051b850101925086831115614e27575f5ffd5b6020840193505b82841015610bf6578335614e4181614af3565b825260209384019390910190614e2e565b5f5f60408385031215614e63575f5ffd5b50508035926020909101359150565b5f5f83601f840112614e82575f5ffd5b5081356001600160401b03811115614e98575f5ffd5b602083019150836020828501011115614848575f5ffd5b5f5f5f5f5f5f5f5f5f898b036101c0811215614ec9575f5ffd5b8a35995060208b0135614edb81614af3565b985060408b0135614eeb81614af3565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115614f1a575f5ffd5b614f268d828e01614e72565b90955093505060e060df1982011215614f3d575f5ffd5b5060e08a0190509295985092959850929598565b5f5f60408385031215614f62575f5ffd5b8235614f6d81614af3565b915060208301356001600160401b03811115614f87575f5ffd5b8301601f81018513614f97575f5ffd5b614fa685823560208401614b07565b9150509250929050565b5f5f60408385031215614fc1575f5ffd5b614fca83614b66565b91506020830135614fda81614af3565b809150509250929050565b8015158114610c0b575f5ffd5b5f5f5f5f5f5f60c08789031215615007575f5ffd5b86359550602087013561501981614fe5565b9450604087013561502981614af3565b9350606087013561503981614af3565b9598949750929560808101359460a0909101359350915050565b5f60208284031215615063575f5ffd5b5035919050565b80516001600160a01b03908116835260208083015190911690830152604080820151908301526060808201519083015260808082015115159083015260a08181015115159083015260c090810151910152565b602080825282518282018190525f918401906040840190835b818110156150ff576150e983855161506a565b6020939093019260e092909201916001016150d6565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615121575f5ffd5b88359750602089013561513381614af3565b9650604089013561514381614af3565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615172575f5ffd5b61517e8b828c01614e72565b999c989b5096995094979396929594505050565b5f5f604083850312156151a3575f5ffd5b823591506020830135614fda81614af3565b60e08101610ae3828461506a565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f612d2260208301846151c3565b5f8151808452602084019350602083015f5b82811015615233578151865260209586019590910190600101615215565b5093949350505050565b60ff60f81b8816815260e060208201525f61525b60e08301896151c3565b828103604084015261526d81896151c3565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061529e8185615203565b9a9950505050505050505050565b5f5f83601f8401126152bc575f5ffd5b5081356001600160401b038111156152d2575f5ffd5b6020830191508360208260051b8501011115614848575f5ffd5b5f82601f8301126152fb575f5ffd5b813561530681614a29565b60405161531382826149fd565b80915082815260208101915060208360051b860101925085831115615336575f5ffd5b602085015b83811015614c8d5780356001600160401b03811115615358575f5ffd5b615367886020838a0101614c97565b8452506020928301920161533b565b5f5f5f5f5f6080868803121561538a575f5ffd5b85356001600160401b0381111561539f575f5ffd5b6153ab888289016152ac565b90965094505060208601356001600160401b038111156153c9575f5ffd5b8601601f810188136153d9575f5ffd5b80356153e481614a29565b6040516153f182826149fd565b80915082815260208101915060208360051b85010192508a831115615414575f5ffd5b602084015b838110156154545780356001600160401b03811115615436575f5ffd5b6154458d602083890101614c1f565b84525060209283019201615419565b50955050505060408601356001600160401b03811115615472575f5ffd5b61547e888289016152ec565b92505060608601356001600160401b03811115615499575f5ffd5b6154a5888289016152ec565b9150509295509295909350565b5f5f602083850312156154c3575f5ffd5b82356001600160401b038111156154d8575f5ffd5b6154e4858286016152ac565b90969095509350505050565b60208152815160208201526020820151604082015260018060a01b03604083015116606082015260018060a01b036060830151166080820152608082015160a08201525f60a083015160c08084015261554c60e08401826151c3565b949350505050565b5f5f5f60608486031215615566575f5ffd5b83359250602084013561557881614af3565b929592945050506040919091013590565b5f60208284031215615599575f5ffd5b612d2282614b66565b602081525f612d226020830184615203565b5f5f5f5f604085870312156155c7575f5ffd5b84356001600160401b038111156155dc575f5ffd5b6155e8878288016152ac565b90955093505060208501356001600160401b03811115615606575f5ffd5b8501601f81018713615616575f5ffd5b80356001600160401b0381111561562b575f5ffd5b87602060e08302840101111561563f575f5ffd5b949793965060200194505050565b602080825282518282018190525f918401906040840190835b818110156150ff5783516001600160a01b0316835260209384019390920191600101615666565b634e487b7160e01b5f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f612d22600d8301846156a1565b848152608060208201525f6156f160808301866151c3565b828103604084015261570381866151c3565b91505060ff8316606083015295945050505050565b5f60208284031215615728575f5ffd5b8151612d2281614af3565b5f5f8335601e19843603018112615748575f5ffd5b8301803591506001600160401b03821115615761575f5ffd5b602001915036819003821315614848575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f906157dd9083018486615775565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610ae357610ae36157ea565b600181811c9082168061582557607f821691505b60208210810361584357634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be1983360301811261585d575f5ffd5b9190910192915050565b5f823560de1983360301811261585d575f5ffd5b8881526001600160a01b03888116602083015287166040820152606081018690526080810185905260a0810184905260e060c082018190525f9061529e9083018486615775565b5f602082840312156158d2575f5ffd5b8151612d2281614fe5565b5f60033d11156115045760045f5f3e505f5160e01c90565b5f60443d10156159025790565b6040513d600319016004823e80513d60248201116001600160401b038211171561592b57505090565b80820180516001600160401b03811115615946575050505090565b3d8401600319018282016020011115615960575050505090565b61178b602082850101856149fd565b5f5f60233d111561598857602060045f3e50505f516001905b9091565b604080825283519082018190525f9060208501906060840190835b818110156159c757835115158352602093840193909201916001016159a7565b50508381036020850152809150845180825260208201925060208160051b830101602087015f5b83811015615a2057601f19858403018652615a0a8383516151c3565b60209687019690935091909101906001016159ee565b509098975050505050505050565b5f60208284031215615a3e575f5ffd5b5051919050565b5f5f5f60608486031215615a57575f5ffd5b5050815160208301516040909301519094929350919050565b8981526001600160a01b03898116602083015288166040820152606081018790526080810186905260a081018590526101c060c082018190525f90615ab89083018587615775565b90508235615ac581614af3565b6001600160a01b031660e08301526020830135615ae181614af3565b6001600160a01b03166101008301526040830135610120830152606083013561014083015260ff615b1460808501614b66565b1661016083015260a083013561018083015260c0909201356101a09091015298975050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82615b6157615b61615b3f565b500690565b8082028115828204841417610ae357610ae36157ea565b5f82615b8b57615b8b615b3f565b500490565b80546001600160a01b0319166001600160a01b0392909216919091179055565b601f82111561365257805f5260205f20601f840160051c81016020851015615bd55750805b601f840160051c820191505b81811015613126575f8155600101615be1565b6001600160401b03831115615c0b57615c0b6149e9565b615c1f83615c198354615811565b83615bb0565b5f601f841160018114615c50575f8515615c395750838201355b5f19600387901b1c1916600186901b178355613126565b5f83815260208120601f198716915b82811015615c7f5786850135825560209485019460019092019101615c5f565b5086821015615c9b575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81358155602082013560018201556040820135615cc981614af3565b615cd68160028401615b90565b506060820135615ce581614af3565b615cf28160038401615b90565b506080820135600482015560a082013536839003601e19018112615d14575f5ffd5b820180356001600160401b03811115615d2b575f5ffd5b602082019150803603821315615d3f575f5ffd5b613a69818360058601615bf4565b81516001600160401b03811115615d6657615d666149e9565b615d7a81615d748454615811565b84615bb0565b6020601f821160018114615dac575f8315615d955750848201515b5f19600385901b1c1916600184901b178455613126565b5f84815260208120601f198516915b82811015615ddb5787850151825560209485019460019092019101615dbb565b5084821015615df857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f9061529e9083018486615775565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f612d2260118301846156a1565b81810381811115610ae357610ae36157ea565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f612d2282846156a156fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220d42787de89ba9860c215f96c4d4d4528510d68808dd7dbb45b46f17ceeeadec464736f6c634300081c0033",
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,bool,bool,uint256)[])
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,bool,bool,uint256)[])
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,bool,bool,uint256))
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,bool,bool,uint256))
func (_BranchBridge *BranchBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,bool,bool,uint256))
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

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet) returns()
func (_BranchBridge *BranchBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, _rewardWallet common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "initialize", _threshold, _rewardWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet) returns()
func (_BranchBridge *BranchBridgeSession) Initialize(_threshold uint8, _rewardWallet common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet) returns()
func (_BranchBridge *BranchBridgeTransactorSession) Initialize(_threshold uint8, _rewardWallet common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet)
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

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BranchBridge *BranchBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BranchBridge *BranchBridgeSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeTokenBatch(&_BranchBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeTokenBatch(&_BranchBridge.TransactOpts, args, permitArgs)
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

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_BranchBridge *BranchBridgeTransactor) SetCrossMintableERC20Factory(opts *bind.TransactOpts, _crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setCrossMintableERC20Factory", _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_BranchBridge *BranchBridgeSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetCrossMintableERC20Factory(&_BranchBridge.TransactOpts, _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetCrossMintableERC20Factory(&_BranchBridge.TransactOpts, _crossMintableERC20Factory)
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

// BranchBridgeBridgeTokenBatchProcessedIterator is returned from FilterBridgeTokenBatchProcessed and is used to iterate over the raw logs and unpacked data for BridgeTokenBatchProcessed events raised by the BranchBridge contract.
type BranchBridgeBridgeTokenBatchProcessedIterator struct {
	Event *BranchBridgeBridgeTokenBatchProcessed // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeTokenBatchProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeTokenBatchProcessed)
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
		it.Event = new(BranchBridgeBridgeTokenBatchProcessed)
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
func (it *BranchBridgeBridgeTokenBatchProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeTokenBatchProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeTokenBatchProcessed represents a BridgeTokenBatchProcessed event raised by the BranchBridge contract.
type BranchBridgeBridgeTokenBatchProcessed struct {
	Permit  bool
	Success []bool
	Reason  []string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokenBatchProcessed is a free log retrieval operation binding the contract event 0x9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b28.
//
// Solidity: event BridgeTokenBatchProcessed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeTokenBatchProcessed(opts *bind.FilterOpts, permit []bool) (*BranchBridgeBridgeTokenBatchProcessedIterator, error) {

	var permitRule []interface{}
	for _, permitItem := range permit {
		permitRule = append(permitRule, permitItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeTokenBatchProcessed", permitRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeTokenBatchProcessedIterator{contract: _BranchBridge.contract, event: "BridgeTokenBatchProcessed", logs: logs, sub: sub}, nil
}

// WatchBridgeTokenBatchProcessed is a free log subscription operation binding the contract event 0x9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b28.
//
// Solidity: event BridgeTokenBatchProcessed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeTokenBatchProcessed(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeTokenBatchProcessed, permit []bool) (event.Subscription, error) {

	var permitRule []interface{}
	for _, permitItem := range permit {
		permitRule = append(permitRule, permitItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeTokenBatchProcessed", permitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeTokenBatchProcessed)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeTokenBatchProcessed", log); err != nil {
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

// ParseBridgeTokenBatchProcessed is a log parse operation binding the contract event 0x9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b28.
//
// Solidity: event BridgeTokenBatchProcessed(bool indexed permit, bool[] success, string[] reason)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeTokenBatchProcessed(log types.Log) (*BranchBridgeBridgeTokenBatchProcessed, error) {
	event := new(BranchBridgeBridgeTokenBatchProcessed)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeTokenBatchProcessed", log); err != nil {
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

// BranchBridgeValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the BranchBridge contract.
type BranchBridgeValidatorUpdatedIterator struct {
	Event *BranchBridgeValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *BranchBridgeValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeValidatorUpdated)
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
		it.Event = new(BranchBridgeValidatorUpdated)
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
func (it *BranchBridgeValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeValidatorUpdated represents a ValidatorUpdated event raised by the BranchBridge contract.
type BranchBridgeValidatorUpdated struct {
	Validator common.Address
	Status    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0x763b63b30e91c843bb39e4379603697003d3b7c1f192619cd782fa33bdc44396.
//
// Solidity: event ValidatorUpdated(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, validator []common.Address, status []bool) (*BranchBridgeValidatorUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ValidatorUpdated", validatorRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeValidatorUpdatedIterator{contract: _BranchBridge.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0x763b63b30e91c843bb39e4379603697003d3b7c1f192619cd782fa33bdc44396.
//
// Solidity: event ValidatorUpdated(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *BranchBridgeValidatorUpdated, validator []common.Address, status []bool) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ValidatorUpdated", validatorRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeValidatorUpdated)
				if err := _BranchBridge.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
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

// ParseValidatorUpdated is a log parse operation binding the contract event 0x763b63b30e91c843bb39e4379603697003d3b7c1f192619cd782fa33bdc44396.
//
// Solidity: event ValidatorUpdated(address indexed validator, bool indexed status)
func (_BranchBridge *BranchBridgeFilterer) ParseValidatorUpdated(log types.Log) (*BranchBridgeValidatorUpdated, error) {
	event := new(BranchBridgeValidatorUpdated)
	if err := _BranchBridge.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
