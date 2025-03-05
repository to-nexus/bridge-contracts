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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"bridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainIDs\",\"type\":\"uint256\"}],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_nexus\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nexus\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"pendingArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"pendingReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"_crossMintableERC20Factory\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"nexus_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"setSafetyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"success\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"reason\",\"type\":\"string[]\"}],\"name\":\"BridgeTokenBatchProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20FactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RewardWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RegistryExistFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"localTokenRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remoteTokenRate\",\"type\":\"uint256\"}],\"name\":\"RegistryInvalidRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroLocalRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroRemoteRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedSendValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidExFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidGasFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMinAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidPendingAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotExistFeeStation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"StandardNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"f30589c3": "allValidators()",
		"47666cb1": "bridgeFeeStation()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"9118b5eb": "bridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[])",
		"b7f3358d": "changeThreshold(uint8)",
		"0b43c02c": "clearPending(uint256)",
		"79b53840": "createToken(uint256,address,uint256,uint256,uint256,string,uint8)",
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
		"a3f5c1d2": "nexus()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"d2ff130d": "pauseToken(uint256,address)",
		"5c975abb": "paused()",
		"1a9a379f": "pendingArguments(uint256,uint256)",
		"3d3e68c2": "pendingReason(uint256,uint256)",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"4227fd2d": "registerToken(uint256,bool,address,address,uint256,uint256,uint256)",
		"d7c82f32": "removeFeeStation()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
		"3960e787": "retryFinalizeBridge(uint256,uint256)",
		"030372c3": "retryFinalizeBridgeBatch(uint256,uint256[])",
		"1a1aebbb": "setCrossMintableERC20Factory(address)",
		"54db0126": "setFeeStation(address)",
		"5958621e": "setRewardWallet(address)",
		"39a621f3": "setSafetyLimit(uint256,address,uint256)",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516162456100f95f395f818161372e01528181613757015261391601526162455ff3fe60806040526004361061035e575f3560e01c8063814914b5116101bd578063ae766389116100f2578063d605665b11610092578063f30589c31161006d578063f30589c3146109ee578063f450964314610a0f578063f698da2514610a2e578063facd743b14610a42575f5ffd5b8063d605665b146109a8578063d7c82f32146109bb578063f2fde38b146109cf575f5ffd5b8063cbae5958116100cd578063cbae595814610937578063cf56118e14610956578063d2ff130d1461096a578063d5717fc514610989575f5ffd5b8063ae766389146108ca578063aed1d40314610904578063b7f3358d14610918575f5ffd5b80639118b5eb1161015d57806396ce07951161013857806396ce07951461084a578063a3f5c1d21461085e578063ad3cb1cc1461087b578063ae6893f8146108ab575f5ffd5b80639118b5eb1461080457806391cf6d3e146108175780639300c9261461082b575f5ffd5b806384d58d421161019857806384d58d421461077757806388d67d6d146107965780638da5cb5b146107a95780638f517c17146107e5575f5ffd5b8063814914b5146107105780638456cb591461073c57806384b0196e14610750575f5ffd5b806347666cb1116102935780635b605f5c116102335780637101fcd31161020e5780637101fcd314610692578063715018a6146106b157806379214874146106c557806379b53840146106f1575f5ffd5b80635b605f5c146106305780635c975abb1461065c5780635fd262de1461067f575f5ffd5b80635187599d1161026e5780635187599d146105b157806352d1902d146105d057806354db0126146105f25780635958621e14610611575f5ffd5b806347666cb1146105545780634d5d00561461058b5780634f1ef2861461059e575f5ffd5b80633960e787116102fe5780633f4ba83a116102d95780633f4ba83a146104e257806340a141ff146104f65780634227fd2d1461051557806342cde4e814610534575f5ffd5b80633960e7871461047857806339a621f3146104975780633d3e68c2146104b6575f5ffd5b80631938e0f2116103395780631938e0f2146103fb5780631a1aebbb1461040e5780631a9a379f1461042d5780631d40f0d814610459575f5ffd5b8063030372c3146103895780630b43c02c146103bd5780631327d3d8146103dc575f5ffd5b3661038557345f03610383576040516365d14ce560e11b815260040160405180910390fd5b005b5f5ffd5b348015610394575f5ffd5b506103a86103a3366004614cb8565b610a61565b60405190151581526020015b60405180910390f35b3480156103c8575f5ffd5b506103836103d7366004614d60565b610aa5565b3480156103e7575f5ffd5b506103836103f6366004614d8b565b610b03565b6103a8610409366004614e9a565b610b11565b348015610419575f5ffd5b50610383610428366004614d8b565b610db9565b348015610438575f5ffd5b5061044c610447366004614f53565b610e69565b6040516103b49190614fa1565b348015610464575f5ffd5b50610383610473366004614ffd565b610fb2565b348015610483575f5ffd5b506103a8610492366004614f53565b610fec565b3480156104a2575f5ffd5b506103836104b13660046150a2565b611190565b3480156104c1575f5ffd5b506104d56104d0366004614f53565b6111f6565b6040516103b491906150d7565b3480156104ed575f5ffd5b506103836112a4565b348015610501575f5ffd5b50610383610510366004614d8b565b6112b6565b348015610520575f5ffd5b5061038361052f3660046150f6565b6112c0565b34801561053f575f5ffd5b505f5460405160ff90911681526020016103b4565b34801561055f575f5ffd5b50606454610573906001600160a01b031681565b6040516001600160a01b0390911681526020016103b4565b6103a86105993660046151a3565b6112e0565b6103836105ac3660046152a4565b61157f565b3480156105bc575f5ffd5b506103836105cb366004615303565b61159a565b3480156105db575f5ffd5b506105e46116a7565b6040519081526020016103b4565b3480156105fd575f5ffd5b5061038361060c366004614d8b565b6116c3565b34801561061c575f5ffd5b5061038361062b366004614d8b565b61173b565b34801561063b575f5ffd5b5061064f61064a366004614d60565b6117b3565b6040516103b491906153b2565b348015610667575f5ffd5b505f5160206161ae5f395f51905f525460ff166103a8565b6103a861068d366004615400565b61191c565b34801561069d575f5ffd5b506103836106ac366004614ffd565b611a0f565b3480156106bc575f5ffd5b50610383611a25565b3480156106d0575f5ffd5b506106e46106df366004614d60565b611a36565b6040516103b491906154c2565b3480156106fc575f5ffd5b5061057361070b3660046154d4565b611a52565b34801561071b575f5ffd5b5061072f61072a366004615567565b611b6b565b6040516103b4919061558a565b348015610747575f5ffd5b50610383611c0d565b34801561075b575f5ffd5b50610764611c1d565b6040516103b49796959493929190615599565b348015610782575f5ffd5b50610383610791366004615567565b611cc6565b6103a86107a43660046156d2565b611dbf565b3480156107b4575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610573565b3480156107f0575f5ffd5b50603254610573906001600160a01b031681565b61038361081236600461580e565b611e5a565b348015610822575f5ffd5b506066546105e4565b348015610836575f5ffd5b50610383610845366004614ffd565b6121c4565b348015610855575f5ffd5b506105e46121fb565b348015610869575f5ffd5b506065546001600160a01b0316610573565b348015610886575f5ffd5b506104d5604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156108b6575f5ffd5b506105e46108c5366004614d60565b61226b565b3480156108d5575f5ffd5b506108e96108e43660046150a2565b612287565b604080519384526020840192909252908201526060016103b4565b34801561090f575f5ffd5b506105e4612314565b348015610923575f5ffd5b5061038361093236600461584c565b61231f565b348015610942575f5ffd5b50610573610951366004614d60565b61236f565b348015610961575f5ffd5b506106e461237b565b348015610975575f5ffd5b50610383610984366004615567565b612387565b348015610994575f5ffd5b506105e46109a3366004614d60565b612484565b6103836109b6366004615865565b6124a0565b3480156109c6575f5ffd5b50610383612846565b3480156109da575f5ffd5b506103836109e9366004614d8b565b6128b3565b3480156109f9575f5ffd5b50610a026128ed565b6040516103b491906158fe565b348015610a1a575f5ffd5b50610383610a29366004615567565b6128f9565b348015610a39575f5ffd5b506105e46129dd565b348015610a4d575f5ffd5b506103a8610a5c366004614d8b565b6129e6565b5f805b8251811015610a9957610a9084848381518110610a8357610a8361593e565b6020026020010151610fec565b50600101610a64565b50600190505b92915050565b610aad6129f2565b5f818152603560205260408120610ac690600301612a4d565b90505f5b8151811015610afe57610af683838381518110610ae957610ae961593e565b6020026020010151612a60565b600101610aca565b505050565b610b0e816001612b0a565b50565b5f610b1a612bc7565b610b22612bf7565b610b3c8535610b376060880160408901614d8b565b612c2e565b610b4c6060870160408801614d8b565b90610b7b576040516353b2527760e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610ba6576040516329e2b03f60e21b815260048101929092526024820152604401610b72565b50610bb390508535612484565b602086013514610bc38635612484565b86602001359091610bf057604051631351db4160e31b815260048101929092526024820152604401610b72565b50610c8290507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020870135610c2c6060890160408a01614d8b565b610c3c60808a0160608b01614d8b565b60808a0135610c4e60a08c018c615952565b604051602001610c6497969594939291906159bc565b60405160208183030381529060405280519060200120858585612c45565b610ca085355f90815260356020526040902060020180546001019055565b5f80610cd18735610cb760608a0160408b01614d8b565b610cc760808b0160608c01614d8b565b8a60800135612e3e565b915091508115610d5c57610ceb6080880160608901614d8b565b6001600160a01b0316602088013588357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610d2c60608c0160408d01614d8b565b604080516001600160a01b03909216825260808d01356020830152429082015260600160405180910390a4610d95565b610d668782613025565b6040516020880135907f40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2905f90a25b600192505050610db160015f5160206161ce5f395f51905f5255565b949350505050565b610dc16129f2565b6032546001600160a01b03168015610df857604051639ad61dbd60e01b81526001600160a01b039091166004820152602401610b72565b506001600160a01b038116610e2057604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517f18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9905f90a250565b610eaf6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f8381526035602090815260408083208584526005908101835292819020815160c0810183528154815260018201549381019390935260028101546001600160a01b03908116928401929092526003810154909116606083015260048101546080830152918201805491929160a084019190610f2a90615a09565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5690615a09565b8015610fa15780601f10610f7857610100808354040283529160200191610fa1565b820191905f5260205f20905b815481529060010190602001808311610f8457829003601f168201915b505050505081525050905092915050565b5f5b8151811015610fe857610fe0828281518110610fd257610fd261593e565b60200260200101515f612b0a565b600101610fb4565b5050565b5f610ff5612bc7565b610ffd612bf7565b5f6110088484610e69565b5f858152603760209081526040808320848201516001600160a01b031684529091529020600581015491925090610100900460ff16156110b6576080820151600782015460408401518792909190808210156110955760405163134c17e760e21b815260048101949094526001600160a01b03909216602484015260448301526064820152608401610b72565b505050508160800151816007015f8282546110b09190615a55565b90915550505b5f5f6110d087856040015186606001518760800151612e3e565b915091508181906110f45760405162461bcd60e51b8152600401610b7291906150d7565b506110ff8787612a60565b83606001516001600160a01b03168460200151887f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b876040015188608001514260405161116a939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a46001945050505050610a9f60015f5160206161ce5f395f51905f5255565b6111986129f2565b6111a28383612c2e565b82906111cd5760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f9283526037602090815260408085206001600160a01b039094168552929052912060040155565b5f828152603560209081526040808320848452600601909152902080546060919061122090615a09565b80601f016020809104026020016040519081016040528092919081815260200182805461124c90615a09565b80156112975780601f1061126e57610100808354040283529160200191611297565b820191905f5260205f20905b81548152906001019060200180831161127a57829003601f168201915b5050505050905092915050565b6112ac6129f2565b6112b4613128565b565b610b0e815f612b0a565b6112c86129f2565b6112d787878787878787613181565b50505050505050565b5f6112e9612bc7565b89896112f58282612c2e565b81906113205760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168452909152902060050154819060ff1615611375576040516338384f6f60e11b81526001600160a01b039091166004820152602401610b72565b5061137e612bf7565b61138b8c8c8b8b8b6134b0565b90985096508661139b898b615a68565b6113a59190615a68565b60408501351015876113b78a8c615a68565b6113c19190615a68565b856040013590916113ee576040516365efbabf60e11b815260048101929092526024820152604401610b72565b505f90506113ff6020860186614d8b565b6001600160a01b031663d505accf61141d6040880160208901614d8b565b30604089013560608a013561143860a08c0160808d0161584c565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a088013560c482015260c088013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031660e09490941b9390931790925292505f91506114bf90870187614d8b565b90505f5f60205f8551602087015f875af1806114e0576040513d5f823e3d81fd5b50505f513d915081156114f65780600114611510565b6115036020890189614d8b565b6001600160a01b03163b15155b61152d57604051631c92cad760e01b815260040160405180910390fd5b505050506115568c8c8660200160208101906115499190614d8b565b8d8d8d8d60018e8e6135e7565b6001925061157060015f5160206161ce5f395f51905f5255565b50509998505050505050505050565b611587613723565b611590826137c7565b610fe882826137cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156115de5750825b90505f826001600160401b031660011480156115f95750303b155b905081158015611607575080155b156116255760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561164f57845460ff60401b1916600160401b1785555b611659878761388b565b83156112d757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b5f6116b061390b565b505f51602061618e5f395f51905f525b90565b6116cb6129f2565b6001600160a01b0381166116f257604051630fb9363360e41b815260040160405180910390fd5b606480546001600160a01b0319166001600160a01b0383169081179091556040517f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444905f90a250565b6117436129f2565b6001600160a01b03811661176a57604051630fb9363360e41b815260040160405180910390fd5b606580546001600160a01b0319166001600160a01b0383169081179091556040517f5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b905f90a250565b5f818152603660205260408120606091906117cd90612a4d565b90505f81516001600160401b038111156117e9576117e9614c56565b60405190808252806020026020018201604052801561182257816020015b61180f614bb3565b8152602001906001900390816118075790505b5090505f5b82518110156119145760375f8681526020019081526020015f205f8483815181106118545761185461593e565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f208251610120810184528154851681526001820154909416918401919091526002810154918301919091526003810154606083015260048101546080830152600581015460ff808216151560a08501526101009182900416151560c0840152600682015460e08401526007909101549082015282518390839081106119015761190161593e565b6020908102919091010152600101611827565b509392505050565b5f611925612bc7565b88886119318282612c2e565b819061195c5760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168452909152902060050154819060ff16156119b1576040516338384f6f60e11b81526001600160a01b039091166004820152602401610b72565b506119ba612bf7565b5f5f6119c98d8d8c8c8c6134b0565b915091506119e58d8d6119d93390565b8e8e87875f8f8f6135e7565b600194505050611a0160015f5160206161ce5f395f51905f5255565b505098975050505050505050565b611a1c6104736001612a4d565b610b0e816121c4565b611a2d6129f2565b6112b45f613954565b5f818152603560205260409020606090610a9f90600301612a4d565b5f611a5b6129f2565b6032546001600160a01b0316611a84576040516315aeca0d60e11b815260040160405180910390fd5b6032546040516bffffffffffffffffffffffff1960608a901b1660208201526001600160a01b0390911690634804a041906034016040516020818303038152906040528051906020012085604051602001611adf9190615a92565b60405160208183030381529060405286866040518563ffffffff1660e01b8152600401611b0f9493929190615ab3565b6020604051808303815f875af1158015611b2b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b4f9190615af2565b9050611b60885f838a8a8a8a613181565b979650505050505050565b611b73614bb3565b505f8281526037602090815260408083206001600160a01b038086168552908352928190208151610120810183528154851681526001820154909416928401929092526002820154908301526003810154606083015260048101546080830152600581015460ff808216151560a08501526101009182900416151560c0840152600682015460e08401526007909101549082015292915050565b611c156129f2565b6112b46139c4565b5f60608082808083815f51602061616e5f395f51905f528054909150158015611c4857506001810154155b611c8c5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610b72565b611c94613a0c565b611c9c613acc565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b611cce6129f2565b5f828152603660205260409020611ce59082613b0a565b8190611d105760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168452909152902060050154819060ff16611d6457604051636c508f9f60e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168085529252808320600501805460ff1916905551909184917fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e99190a35050565b5f805b85811015611e4d57611e44878783818110611ddf57611ddf61593e565b9050602002810190611df19190615b0d565b868381518110611e0357611e0361593e565b6020026020010151868481518110611e1d57611e1d61593e565b6020026020010151868581518110611e3757611e3761593e565b6020026020010151610b11565b50600101611dc2565b5060019695505050505050565b5f816001600160401b03811115611e7357611e73614c56565b604051908082528060200260200182016040528015611e9c578160200160208202803683370190505b5090505f826001600160401b03811115611eb857611eb8614c56565b604051908082528060200260200182016040528015611eeb57816020015b6060815260200190600190039081611ed65790505b5090505f5b838110156121815730635fd262de868684818110611f1057611f1061593e565b9050602002810190611f229190615b2b565b35878785818110611f3557611f3561593e565b9050602002810190611f479190615b2b565b611f58906040810190602001614d8b565b888886818110611f6a57611f6a61593e565b9050602002810190611f7c9190615b2b565b611f8d906060810190604001614d8b565b898987818110611f9f57611f9f61593e565b9050602002810190611fb19190615b2b565b606001358a8a88818110611fc757611fc761593e565b9050602002810190611fd99190615b2b565b608001358b8b89818110611fef57611fef61593e565b90506020028101906120019190615b2b565b60a001358c8c8a8181106120175761201761593e565b90506020028101906120299190615b2b565b6120379060c0810190615952565b6040518963ffffffff1660e01b815260040161205a989796959493929190615b3f565b6020604051808303815f875af1925050508015612094575060408051601f3d908101601f1916820190925261209191810190615b86565b60015b612153576120a0615ba1565b806308c379a0036120e357506120b4615bb9565b806120bf575061211b565b808383815181106120d2576120d261593e565b602002602001018190525050612179565b634e487b710361211b576120f5615c33565b90612100575061211b565b61210981613b2b565b8383815181106120d2576120d261593e565b3d808015612144576040519150601f19603f3d011682016040523d82523d5f602084013e612149565b606091505b5061210981613b67565b5060018382815181106121685761216861593e565b911515602092830291909101909101525b600101611ef0565b505f15157f9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b2883836040516121b6929190615c50565b60405180910390a250505050565b5f5b8151811015610fe8576121f38282815181106121e4576121e461593e565b60200260200101516001612b0a565b6001016121c6565b606454604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015612242573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122669190615cf2565b905090565b5f818152603560205260408120600190810154610a9f91615a68565b6064805460405163ae76638960e01b8152600481018690526001600160a01b038581166024830152604482018590525f938493849392169163ae7663899101606060405180830381865afa1580156122e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123059190615d09565b91989097509095509350505050565b5f6122666001613b7a565b6123276129f2565b5f805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f610a9f600183613b83565b60606122666033612a4d565b61238f6129f2565b5f8281526036602052604090206123a69082613b0a565b81906123d15760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168452909152902060050154819060ff1615612426576040516338384f6f60e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b0385168085529252808320600501805460ff1916600117905551909184917ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a39190a35050565b5f81815260356020526040812060020154610a9f906001615a68565b8281146124c05760405163214485c960e01b815260040160405180910390fd5b5f836001600160401b038111156124d9576124d9614c56565b604051908082528060200260200182016040528015612502578160200160208202803683370190505b5090505f846001600160401b0381111561251e5761251e614c56565b60405190808252806020026020018201604052801561255157816020015b606081526020019060019003908161253c5790505b5090505f5b858110156128005730634d5d00568888848181106125765761257661593e565b90506020028101906125889190615b2b565b3589898581811061259b5761259b61593e565b90506020028101906125ad9190615b2b565b6125be906040810190602001614d8b565b8a8a868181106125d0576125d061593e565b90506020028101906125e29190615b2b565b6125f3906060810190604001614d8b565b8b8b878181106126055761260561593e565b90506020028101906126179190615b2b565b606001358c8c8881811061262d5761262d61593e565b905060200281019061263f9190615b2b565b608001358d8d898181106126555761265561593e565b90506020028101906126679190615b2b565b60a001358e8e8a81811061267d5761267d61593e565b905060200281019061268f9190615b2b565b61269d9060c0810190615952565b8e8e8c8181106126af576126af61593e565b905060e002016040518a63ffffffff1660e01b81526004016126d999989796959493929190615d34565b6020604051808303815f875af1925050508015612713575060408051601f3d908101601f1916820190925261271091810190615b86565b60015b6127d25761271f615ba1565b806308c379a0036127625750612733615bb9565b8061273e575061279a565b808383815181106127515761275161593e565b6020026020010181905250506127f8565b634e487b710361279a57612774615c33565b9061277f575061279a565b61278881613b2b565b8383815181106127515761275161593e565b3d8080156127c3576040519150601f19603f3d011682016040523d82523d5f602084013e6127c8565b606091505b5061278881613b67565b5060018382815181106127e7576127e761593e565b911515602092830291909101909101525b600101612556565b50600115157f9a646b7804bc3bd0be428ae57a316a81e7d597a26e3e44099233ce9e756c9b288383604051612836929190615c50565b60405180910390a2505050505050565b61284e6129f2565b6064546001600160a01b031661287757604051637bda535160e01b815260040160405180910390fd5b606480546001600160a01b03191690556040515f907f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444908290a2565b6128bb6129f2565b6001600160a01b0381166128e457604051631e4fbdf760e01b81525f6004820152602401610b72565b610b0e81613954565b60606122666001612a4d565b6129016129f2565b5f8281526036602052604090206129189082613b8e565b81906129435760405163153096f360e11b81526001600160a01b039091166004820152602401610b72565b505f8281526037602090815260408083206001600160a01b038516808552925280832080546001600160a01b03199081168255600182018054909116905560028101849055600381018490556004810184905560058101805461ffff191690556006810184905560070183905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f612266613ba2565b5f610a9f600183613b0a565b33612a247f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146112b45760405163118cdaa760e01b8152336004820152602401610b72565b60605f612a5983613bab565b9392505050565b5f828152603560205260409020600301612a7a8183613c04565b8290612a9c57604051637f11bea960e01b8152600401610b7291815260200190565b505f8281526003820160205260408120612ab591614c0c565b5f8281526002808301602052604082208281556001810183905590810180546001600160a01b031990811690915560038201805490911690556004810182905590612b036005830182614c0c565b5050505050565b612b126129f2565b8015612b5457612b23600183613c0f565b8290612b4e576040516329a04e7760e21b81526001600160a01b039091166004820152602401610b72565b50612b8c565b612b5f600183613b8e565b8290612b8a5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610b72565b505b604051811515906001600160a01b038416907f763b63b30e91c843bb39e4379603697003d3b7c1f192619cd782fa33bdc44396905f90a35050565b5f5160206161ae5f395f51905f525460ff16156112b45760405163d93c066560e01b815260040160405180910390fd5b5f5160206161ce5f395f51905f52805460011901612c2857604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f828152603660205260408120612a599083613b0a565b8251825181148015612c575750815181145b835183518392612c8b576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610b72565b50505f5482915060ff16811015612cb857604051632fcba65760e11b8152600401610b7291815260200190565b505f80826001600160401b03811115612cd357612cd3614c56565b604051908082528060200260200182016040528015612cfc578160200160208202803683370190505b5090505f5b83811015612e09575f612d6c888381518110612d1f57612d1f61593e565b6020026020010151888481518110612d3957612d3961593e565b6020026020010151888581518110612d5357612d5361593e565b6020026020010151612d648d613c23565b929190613c4f565b9050612d77816129e6565b8190612da25760405163845a09e760e01b81526001600160a01b039091166004820152602401610b72565b505f805b8451811015612df257848181518110612dc157612dc161593e565b60200260200101516001600160a01b0316836001600160a01b031603612dea5760019150612df2565b600101612da6565b5080612dff578460010194505b5050600101612d01565b505f54829060ff16811015612e3457604051632fcba65760e11b8152600401610b7291815260200190565b5050505050505050565b5f8481526037602090815260408083206001600160a01b038088168552908352818420825161012081018452815483168152600182015490921693820193909352600283015491810191909152600382015460608281019190915260048301546080830152600583015460ff80821615801560a086015261010092839004909116151560c0850152600685015460e08501526007909401549083015291612f0f57505060408051808201909152600f81526e1d1bdad95b881a5cc81c185d5cd959608a1b60208201525f915061301c565b608081015115801590612f255750838160800151105b15612f4e575f6040518060600160405280602281526020016161ee60229139925092505061301c565b604081015115612f8c578060400151600114612f7a576040810151612f739085615e03565b9350612f8c565b6060810151612f899085615e2e565b93505b5f196001600160a01b03871601612fc157612fa78585613c7b565b6001925060405180602001604052805f815250915061301a565b831561301a575f8781526037602090815260408083206001600160a01b038a168452909152902060050154610100900460ff161561300f5761300587878787613d11565b925092505061301c565b613005868686613e85565b505b94509492505050565b81355f908152603560209081526040909120908301356130486003830182613fef565b819061306a576040516307a5c91d60e31b8152600401610b7291815260200190565b505f818152600583016020526040902084906130868282615f5e565b50505f81815260068301602052604090206130a18482615ffe565b5083355f908152603760205260408082209082906130c59060608901908901614d8b565b6001600160a01b0316815260208101919091526040015f206005810154909150610100900460ff1615612b03578460800135816007015f8282546131099190615a68565b90915550505050505050565b60015f5160206161ce5f395f51905f5255565b613130613ffa565b5f5160206161ae5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001612364565b5f878152603560205260409020613199603389613fef565b156131a2578781555b6001600160a01b0386166131c957604051636ca1fdd760e01b815260040160405180910390fd5b5f8881526036602052604090206131e09087613c0f565b869061320b576040516311dd05f360e31b81526001600160a01b039091166004820152602401610b72565b508284146132e957835f036132335760405163f0006e8f60e01b815260040160405180910390fd5b825f0361325357604051639a5ea84b60e01b815260040160405180910390fd5b828411156132a4578260011480156132735750613271600a856160b8565b155b8484909161329d57604051635c66012760e01b815260048101929092526024820152604401610b72565b50506132e9565b8360011480156132bc57506132ba600a846160b8565b155b848490916132e657604051635c66012760e01b815260048101929092526024820152604401610b72565b50505b604051806101200160405280876001600160a01b03168152602001866001600160a01b031681526020018581526020018481526020018381526020015f1515815260200188151581526020015f81526020015f81525060375f8a81526020019081526020015f205f886001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015f6101000a81548160ff02191690831515021790555060c08201518160050160016101000a81548160ff02191690831515021790555060e082015181600601556101008201518160070155905050846001600160a01b0316866001600160a01b0316897fef354a4d67e88f3c656af891265f876391922dd3618eee9734a68467629be5e98787878d60405161349e9493929190938452602084019290925260408301521515606082015260800190565b60405180910390a45050505050505050565b6064545f9081906001600160a01b03166134ce57505f9050806135dd565b6064805460405163ae76638960e01b8152600481018a90526001600160a01b038981166024830152604482018990525f9392169163ae7663899101606060405180830381865afa158015613524573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135489190615d09565b9094509250905080868181101561357b5760405163428d097560e01b815260048101929092526024820152604401610b72565b5083905085818110156135aa576040516307cd920960e51b815260048101929092526024820152604401610b72565b5082905084818110156135d9576040516379a9e62b60e11b815260048101929092526024820152604401610b72565b5050505b9550959350505050565b6135fd8a8a8a896135f8898b615a68565b614029565b5f5f6136088c61226b565b915060375f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b031690508a6001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7848e8e8e428d8d8d60405161369d9897969594939291906160cb565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a6040516136f4929190918252602082015260400190565b60405180910390a45f8c8152603560205260409020600190810180549091019055505050505050505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806137a957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661379d5f51602061618e5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156112b45760405163703e46dd60e11b815260040160405180910390fd5b610b0e6129f2565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613829575060408051601f3d908101601f1916820190925261382691810190615cf2565b60015b61385157604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610b72565b5f51602061618e5f395f51905f52811461388157604051632a87526960e21b815260048101829052602401610b72565b610afe838361426c565b6138936142c1565b61389c3361430a565b6138a461431b565b6138ac614323565b6138b4614333565b6138bd82614343565b6001600160a01b0381166138e457604051630fb9363360e41b815260040160405180910390fd5b606580546001600160a01b0319166001600160a01b03929092169190911790555043606655565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112b45760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6139cc612bc7565b5f5160206161ae5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613169565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061616e5f395f51905f5291613a4a90615a09565b80601f0160208091040260200160405190810160405280929190818152602001828054613a7690615a09565b8015613ac15780601f10613a9857610100808354040283529160200191613ac1565b820191905f5260205f20905b815481529060010190602001808311613aa457829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061616e5f395f51905f5291613a4a90615a09565b6001600160a01b0381165f9081526001830160205260408120541515612a59565b6040516b02830b734b19031b7b2329d160a51b6020820152602c8101829052606090604c015b6040516020818303038152906040529050919050565b606081604051602001613b519190616115565b5f610a9f825490565b5f612a5983836143a8565b5f612a59836001600160a01b0384166143ce565b5f6122666144b1565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613bf857602002820191905f5260205f20905b815481526020019060010190808311613be4575b50505050509050919050565b5f612a5983836143ce565b5f612a59836001600160a01b038416614524565b5f610a9f613c2f613ba2565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613c5f88888888614570565b925092509250613c6f8282614638565b50909695505050505050565b804747821115613ca7576040516371e4d07760e11b815260048101929092526024820152604401610b72565b50505f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114613cf3576040519150601f19603f3d011682016040523d82523d5f602084013e613cf8565b606091505b509150915081613d0b57613d0b816146f0565b50505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613d80575060408051601f3d908101601f19168201909252613d7d91810190615b86565b60015b613e1857613d8c615ba1565b806308c379a003613db25750613da0615bb9565b80613dab5750613de0565b905061301c565b634e487b7103613de057613dc4615c33565b90613dcf5750613de0565b613dd881613b2b565b91505061301c565b3d808015613e09576040519150601f19603f3d011682016040523d82523d5f602084013e613e0e565b606091505b50613dd881613b67565b8015613e43576001925060405180602001604052805f8152509150613e3e878786614719565b61301a565b6040518060400160405280601f81526020017f5374616e646172644272696467653a207472616e73666572206661696c65640081525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613ef4575060408051601f3d908101601f19168201909252613ef191810190615b86565b60015b613f8c57613f00615ba1565b806308c379a003613f265750613f14615bb9565b80613f1f5750613f54565b9050613fe7565b634e487b7103613f5457613f38615c33565b90613f435750613f54565b613f4c81613b2b565b915050613fe7565b3d808015613f7d576040519150601f19603f3d011682016040523d82523d5f602084013e613f82565b606091505b50613f4c81613b67565b8015613fac576001925060405180602001604052805f8152509150613fe5565b6040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b5f612a598383614524565b5f5160206161ae5f395f51905f525460ff166112b457604051638dfc202b60e01b815260040160405180910390fd5b5f8581526037602090815260408083206001600160a01b038816845290915290206002015460018111156140975761406181846160b8565b8590849015614094576040516348d53e0760e01b81526001600160a01b0390921660048301526024820152604401610b72565b50505b5f196001600160a01b0386160161410a576140b28284615a68565b34146140be8385615a68565b3490916140e7576040516329e2b03f60e21b815260048101929092526024820152604401610b72565b5050811561410557606554614105906001600160a01b031683613c7b565b614264565b5f348015614134576040516329e2b03f60e21b815260048101929092526024820152604401610b72565b50614158905084306141468587615a68565b6001600160a01b0389169291906147ad565b811561417857606554614178906001600160a01b03878116911684614814565b5f8681526037602090815260408083206001600160a01b0389168452909152902060050154610100900460ff16156141b557614105868685614845565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af11580156141ff573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142239190615b86565b8585859091926142605760405163601bc95b60e11b81526001600160a01b0393841660048201529290911660248301526044820152606401610b72565b5050505b505050505050565b61427582614883565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156142b957610afe82826148e6565b610fe8614958565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166112b457604051631afcd79f60e31b815260040160405180910390fd5b6143126142c1565b610b0e81614977565b6112b46142c1565b61432b6142c1565b6112b461497f565b61433b6142c1565b6112b461499f565b61434b6142c1565b614393604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506149a7565b5f805460ff191660ff92909216919091179055565b5f825f0182815481106143bd576143bd61593e565b905f5260205f200154905092915050565b5f81815260018301602052604081205480156144a8575f6143f0600183615a55565b85549091505f9061440390600190615a55565b9050808214614462575f865f0182815481106144215761442161593e565b905f5260205f200154905080875f0184815481106144415761444161593e565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806144735761447361613a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a9f565b5f915050610a9f565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6144db6149b9565b6144e3614a21565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f81815260018301602052604081205461456957508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a9f565b505f610a9f565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156145a957505f9150600390508261462e565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156145fa573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661462557505f92506001915082905061462e565b92505f91508190505b9450945094915050565b5f82600381111561464b5761464b61614e565b03614654575050565b60018260038111156146685761466861614e565b036146865760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561469a5761469a61614e565b036146bb5760405163fce698f760e01b815260048101829052602401610b72565b60038260038111156146cf576146cf61614e565b03610fe8576040516335e2f38360e21b815260048101829052602401610b72565b8051156147005780518082602001fd5b60405163fd5478bd60e01b815260040160405180910390fd5b5f8381526037602090815260408083206001600160a01b0386168452909152902060078101546147499083615a68565b8160060154101584848490919261478c57604051634252ea9360e01b815260048101939093526001600160a01b0390911660248301526044820152606401610b72565b50505081816006015f8282546147a29190615a55565b909155505050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052613d0b9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614a63565b6040516001600160a01b03838116602483015260448201839052610afe91859182169063a9059cbb906064016147e2565b5f8381526037602090815260408083206001600160a01b038616845290915281206006018054839290614879908490615a68565b9091555050505050565b806001600160a01b03163b5f036148b857604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610b72565b5f51602061618e5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516149029190616162565b5f60405180830381855af49150503d805f811461493a576040519150601f19603f3d011682016040523d82523d5f602084013e61493f565b606091505b509150915061494f858383614acf565b95945050505050565b34156112b45760405163b398979f60e01b815260040160405180910390fd5b6128bb6142c1565b6149876142c1565b5f5160206161ae5f395f51905f52805460ff19169055565b6131156142c1565b6149af6142c1565b610fe88282614b2b565b5f5f51602061616e5f395f51905f52816149d1613a0c565b8051909150156149e957805160209091012092915050565b815480156149f8579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061616e5f395f51905f5281614a39613acc565b805190915015614a5157805160209091012092915050565b600182015480156149f8579392505050565b5f5f60205f8451602086015f885af180614a82576040513d5f823e3d81fd5b50505f513d91508115614a99578060011415614aa6565b6001600160a01b0384163b155b15613d0b57604051635274afe760e01b81526001600160a01b0385166004820152602401610b72565b606082614ae457614adf82614b8a565b612a59565b8151158015614afb57506001600160a01b0384163b155b15614b2457604051639996b31560e01b81526001600160a01b0385166004820152602401610b72565b5092915050565b614b336142c1565b5f51602061616e5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614b6c8482615ffe565b5060038101614b7b8382615ffe565b505f8082556001909101555050565b805115614b9a5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518061012001604052805f6001600160a01b031681526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f151581526020015f151581526020015f81526020015f81525090565b508054614c1890615a09565b5f825580601f10614c27575050565b601f0160209004905f5260205f2090810190610b0e91905b80821115614c52575f8155600101614c3f565b5090565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f191681016001600160401b0381118282101715614c8f57614c8f614c56565b6040525050565b5f6001600160401b03821115614cae57614cae614c56565b5060051b60200190565b5f5f60408385031215614cc9575f5ffd5b8235915060208301356001600160401b03811115614ce5575f5ffd5b8301601f81018513614cf5575f5ffd5b8035614d0081614c96565b604051614d0d8282614c6a565b80915082815260208101915060208360051b850101925087831115614d30575f5ffd5b6020840193505b82841015614d52578335825260209384019390910190614d37565b809450505050509250929050565b5f60208284031215614d70575f5ffd5b5035919050565b6001600160a01b0381168114610b0e575f5ffd5b5f60208284031215614d9b575f5ffd5b8135612a5981614d77565b803560ff81168114614db6575f5ffd5b919050565b5f82601f830112614dca575f5ffd5b8135614dd581614c96565b604051614de28282614c6a565b80915082815260208101915060208360051b860101925085831115614e05575f5ffd5b602085015b83811015614e2957614e1b81614da6565b835260209283019201614e0a565b5095945050505050565b5f82601f830112614e42575f5ffd5b8135614e4d81614c96565b604051614e5a8282614c6a565b80915082815260208101915060208360051b860101925085831115614e7d575f5ffd5b602085015b83811015614e29578035835260209283019201614e82565b5f5f5f5f60808587031215614ead575f5ffd5b84356001600160401b03811115614ec2575f5ffd5b850160c08188031215614ed3575f5ffd5b935060208501356001600160401b03811115614eed575f5ffd5b614ef987828801614dbb565b93505060408501356001600160401b03811115614f14575f5ffd5b614f2087828801614e33565b92505060608501356001600160401b03811115614f3b575f5ffd5b614f4787828801614e33565b91505092959194509250565b5f5f60408385031215614f64575f5ffd5b50508035926020909101359150565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60208152815160208201526020820151604082015260018060a01b03604083015116606082015260018060a01b036060830151166080820152608082015160a08201525f60a083015160c080840152610db160e0840182614f73565b5f6020828403121561500d575f5ffd5b81356001600160401b03811115615022575f5ffd5b8201601f81018413615032575f5ffd5b803561503d81614c96565b60405161504a8282614c6a565b80915082815260208101915060208360051b85010192508683111561506d575f5ffd5b6020840193505b8284101561509857833561508781614d77565b825260209384019390910190615074565b9695505050505050565b5f5f5f606084860312156150b4575f5ffd5b8335925060208401356150c681614d77565b929592945050506040919091013590565b602081525f612a596020830184614f73565b8015158114610b0e575f5ffd5b5f5f5f5f5f5f5f60e0888a03121561510c575f5ffd5b87359650602088013561511e816150e9565b9550604088013561512e81614d77565b9450606088013561513e81614d77565b9699959850939660808101359560a0820135955060c0909101359350915050565b5f5f83601f84011261516f575f5ffd5b5081356001600160401b03811115615185575f5ffd5b60208301915083602082850101111561519c575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c08112156151bd575f5ffd5b8a35995060208b01356151cf81614d77565b985060408b01356151df81614d77565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b0381111561520e575f5ffd5b61521a8d828e0161515f565b90955093505060e060df1982011215615231575f5ffd5b5060e08a0190509295985092959850929598565b5f5f6001600160401b0384111561525e5761525e614c56565b50604051601f8401601f1916602001906152788282614c6a565b80925084815285858501111561528c575f5ffd5b848460208301375f6020868301015250509392505050565b5f5f604083850312156152b5575f5ffd5b82356152c081614d77565b915060208301356001600160401b038111156152da575f5ffd5b8301601f810185136152ea575f5ffd5b6152f985823560208401615245565b9150509250929050565b5f5f60408385031215615314575f5ffd5b61531d83614da6565b9150602083013561532d81614d77565b809150509250929050565b60018060a01b03815116825260018060a01b03602082015116602083015260408101516040830152606081015160608301526080810151608083015260a081015161538760a084018215159052565b5060c081015161539b60c084018215159052565b5060e0818101519083015261010090810151910152565b602080825282518282018190525f918401906040840190835b818110156153f5576153de838551615338565b6020939093019261012092909201916001016153cb565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615417575f5ffd5b88359750602089013561542981614d77565b9650604089013561543981614d77565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615468575f5ffd5b6154748b828c0161515f565b999c989b5096995094979396929594505050565b5f8151808452602084019350602083015f5b828110156154b857815186526020958601959091019060010161549a565b5093949350505050565b602081525f612a596020830184615488565b5f5f5f5f5f5f5f60e0888a0312156154ea575f5ffd5b8735965060208801356154fc81614d77565b955060408801359450606088013593506080880135925060a08801356001600160401b0381111561552b575f5ffd5b8801601f81018a1361553b575f5ffd5b61554a8a823560208401615245565b92505061555960c08901614da6565b905092959891949750929550565b5f5f60408385031215615578575f5ffd5b82359150602083013561532d81614d77565b6101208101610a9f8284615338565b60ff60f81b8816815260e060208201525f6155b760e0830189614f73565b82810360408401526155c98189614f73565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506155fa8185615488565b9a9950505050505050505050565b5f5f83601f840112615618575f5ffd5b5081356001600160401b0381111561562e575f5ffd5b6020830191508360208260051b850101111561519c575f5ffd5b5f82601f830112615657575f5ffd5b813561566281614c96565b60405161566f8282614c6a565b80915082815260208101915060208360051b860101925085831115615692575f5ffd5b602085015b83811015614e295780356001600160401b038111156156b4575f5ffd5b6156c3886020838a0101614e33565b84525060209283019201615697565b5f5f5f5f5f608086880312156156e6575f5ffd5b85356001600160401b038111156156fb575f5ffd5b61570788828901615608565b90965094505060208601356001600160401b03811115615725575f5ffd5b8601601f81018813615735575f5ffd5b803561574081614c96565b60405161574d8282614c6a565b80915082815260208101915060208360051b85010192508a831115615770575f5ffd5b602084015b838110156157b05780356001600160401b03811115615792575f5ffd5b6157a18d602083890101614dbb565b84525060209283019201615775565b50955050505060408601356001600160401b038111156157ce575f5ffd5b6157da88828901615648565b92505060608601356001600160401b038111156157f5575f5ffd5b61580188828901615648565b9150509295509295909350565b5f5f6020838503121561581f575f5ffd5b82356001600160401b03811115615834575f5ffd5b61584085828601615608565b90969095509350505050565b5f6020828403121561585c575f5ffd5b612a5982614da6565b5f5f5f5f60408587031215615878575f5ffd5b84356001600160401b0381111561588d575f5ffd5b61589987828801615608565b90955093505060208501356001600160401b038111156158b7575f5ffd5b8501601f810187136158c7575f5ffd5b80356001600160401b038111156158dc575f5ffd5b87602060e0830284010111156158f0575f5ffd5b949793965060200194505050565b602080825282518282018190525f918401906040840190835b818110156153f55783516001600160a01b0316835260209384019390920191600101615917565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112615967575f5ffd5b8301803591506001600160401b03821115615980575f5ffd5b60200191503681900382131561519c575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f906159fc9083018486615994565b9998505050505050505050565b600181811c90821680615a1d57607f821691505b602082108103615a3b57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610a9f57610a9f615a41565b80820180821115610a9f57610a9f615a41565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f612a59600d830184615a7b565b848152608060208201525f615acb6080830186614f73565b8281036040840152615add8186614f73565b91505060ff8316606083015295945050505050565b5f60208284031215615b02575f5ffd5b8151612a5981614d77565b5f823560be19833603018112615b21575f5ffd5b9190910192915050565b5f823560de19833603018112615b21575f5ffd5b8881526001600160a01b03888116602083015287166040820152606081018690526080810185905260a0810184905260e060c082018190525f906155fa9083018486615994565b5f60208284031215615b96575f5ffd5b8151612a59816150e9565b5f60033d11156116c05760045f5f3e505f5160e01c90565b5f60443d1015615bc65790565b6040513d600319016004823e80513d60248201116001600160401b0382111715615bef57505090565b80820180516001600160401b03811115615c0a575050505090565b3d8401600319018282016020011115615c24575050505090565b61191460208285010185614c6a565b5f5f60233d1115615c4c57602060045f3e50505f516001905b9091565b604080825283519082018190525f9060208501906060840190835b81811015615c8b5783511515835260209384019390920191600101615c6b565b50508381036020850152809150845180825260208201925060208160051b830101602087015f5b83811015615ce457601f19858403018652615cce838351614f73565b6020968701969093509190910190600101615cb2565b509098975050505050505050565b5f60208284031215615d02575f5ffd5b5051919050565b5f5f5f60608486031215615d1b575f5ffd5b5050815160208301516040909301519094929350919050565b8981526001600160a01b03898116602083015288166040820152606081018790526080810186905260a081018590526101c060c082018190525f90615d7c9083018587615994565b90508235615d8981614d77565b6001600160a01b031660e08301526020830135615da581614d77565b6001600160a01b03166101008301526040830135610120830152606083013561014083015260ff615dd860808501614da6565b1661016083015260a083013561018083015260c0909201356101a09091015298975050505050505050565b8082028115828204841417610a9f57610a9f615a41565b634e487b7160e01b5f52601260045260245ffd5b5f82615e3c57615e3c615e1a565b500490565b80546001600160a01b0319166001600160a01b0392909216919091179055565b601f821115610afe57805f5260205f20601f840160051c81016020851015615e865750805b601f840160051c820191505b81811015612b03575f8155600101615e92565b6001600160401b03831115615ebc57615ebc614c56565b615ed083615eca8354615a09565b83615e61565b5f601f841160018114615f01575f8515615eea5750838201355b5f19600387901b1c1916600186901b178355612b03565b5f83815260208120601f198716915b82811015615f305786850135825560209485019460019092019101615f10565b5086821015615f4c575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81358155602082013560018201556040820135615f7a81614d77565b615f878160028401615e41565b506060820135615f9681614d77565b615fa38160038401615e41565b506080820135600482015560a082013536839003601e19018112615fc5575f5ffd5b820180356001600160401b03811115615fdc575f5ffd5b602082019150803603821315615ff0575f5ffd5b613d0b818360058601615ea5565b81516001600160401b0381111561601757616017614c56565b61602b816160258454615a09565b84615e61565b6020601f82116001811461605d575f83156160465750848201515b5f19600385901b1c1916600184901b178455612b03565b5f84815260208120601f198516915b8281101561608c578785015182556020948501946001909201910161606c565b50848210156160a957868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f826160c6576160c6615e1a565b500690565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f906155fa9083018486615994565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f612a596011830184615a7b565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f612a598284615a7b56fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0076616c75652069732067726561746572207468616e20736166657479206c696d6974a26469706673582212201f18748915f56621a6fc741767617068d6ed9a6d65eacc7c77af4d288ff10fac64736f6c634300081c0033",
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

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BranchBridge.Contract.AllPendingIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BranchBridge *BranchBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BranchBridge.Contract.AllPendingIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256)[])
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256)[])
func (_BranchBridge *BranchBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256)[])
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256))
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256))
func (_BranchBridge *BranchBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,uint256,uint256,uint256,bool,bool,uint256,uint256))
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

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_BranchBridge *BranchBridgeCaller) Nexus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "nexus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_BranchBridge *BranchBridgeSession) Nexus() (common.Address, error) {
	return _BranchBridge.Contract.Nexus(&_BranchBridge.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) Nexus() (common.Address, error) {
	return _BranchBridge.Contract.Nexus(&_BranchBridge.CallOpts)
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

// PendingArguments is a free data retrieval call binding the contract method 0x1a9a379f.
//
// Solidity: function pendingArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCaller) PendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "pendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryFinalizeArguments)).(*IBridgeRegistryFinalizeArguments)

	return out0, err

}

// PendingArguments is a free data retrieval call binding the contract method 0x1a9a379f.
//
// Solidity: function pendingArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeSession) PendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _BranchBridge.Contract.PendingArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// PendingArguments is a free data retrieval call binding the contract method 0x1a9a379f.
//
// Solidity: function pendingArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCallerSession) PendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _BranchBridge.Contract.PendingArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// PendingReason is a free data retrieval call binding the contract method 0x3d3e68c2.
//
// Solidity: function pendingReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCaller) PendingReason(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (string, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "pendingReason", remoteChainID, index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PendingReason is a free data retrieval call binding the contract method 0x3d3e68c2.
//
// Solidity: function pendingReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeSession) PendingReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.PendingReason(&_BranchBridge.CallOpts, remoteChainID, index)
}

// PendingReason is a free data retrieval call binding the contract method 0x3d3e68c2.
//
// Solidity: function pendingReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCallerSession) PendingReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.PendingReason(&_BranchBridge.CallOpts, remoteChainID, index)
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

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainIDs) returns()
func (_BranchBridge *BranchBridgeTransactor) ClearPending(opts *bind.TransactOpts, remoteChainIDs *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "clearPending", remoteChainIDs)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainIDs) returns()
func (_BranchBridge *BranchBridgeSession) ClearPending(remoteChainIDs *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.ClearPending(&_BranchBridge.TransactOpts, remoteChainIDs)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainIDs) returns()
func (_BranchBridge *BranchBridgeTransactorSession) ClearPending(remoteChainIDs *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.ClearPending(&_BranchBridge.TransactOpts, remoteChainIDs)
}

// CreateToken is a paid mutator transaction binding the contract method 0x79b53840.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, localTokenRate, remoteTokenRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x79b53840.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, localTokenRate, remoteTokenRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x79b53840.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, localTokenRate, remoteTokenRate, safetyLimit, symbol, decimals)
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
// Solidity: function initialize(uint8 _threshold, address _nexus) returns()
func (_BranchBridge *BranchBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, _nexus common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "initialize", _threshold, _nexus)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address _nexus) returns()
func (_BranchBridge *BranchBridgeSession) Initialize(_threshold uint8, _nexus common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _nexus)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address _nexus) returns()
func (_BranchBridge *BranchBridgeTransactorSession) Initialize(_threshold uint8, _nexus common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _nexus)
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

// RegisterToken is a paid mutator transaction binding the contract method 0x4227fd2d.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4227fd2d.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4227fd2d.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, localTokenRate *big.Int, remoteTokenRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, localTokenRate, remoteTokenRate, safetyLimit)
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
// Solidity: function setRewardWallet(address nexus_) returns()
func (_BranchBridge *BranchBridgeTransactor) SetRewardWallet(opts *bind.TransactOpts, nexus_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setRewardWallet", nexus_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address nexus_) returns()
func (_BranchBridge *BranchBridgeSession) SetRewardWallet(nexus_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, nexus_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address nexus_) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetRewardWallet(nexus_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, nexus_)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeTransactor) SetSafetyLimit(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setSafetyLimit", remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetSafetyLimit(&_BranchBridge.TransactOpts, remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetSafetyLimit(&_BranchBridge.TransactOpts, remoteChainID, token, safetyLimit)
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

// BranchBridgeBridgeFinalizePendingIterator is returned from FilterBridgeFinalizePending and is used to iterate over the raw logs and unpacked data for BridgeFinalizePending events raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizePendingIterator struct {
	Event *BranchBridgeBridgeFinalizePending // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeFinalizePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFinalizePending)
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
		it.Event = new(BranchBridgeBridgeFinalizePending)
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
func (it *BranchBridgeBridgeFinalizePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFinalizePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFinalizePending represents a BridgeFinalizePending event raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizePending struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizePending is a free log retrieval operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalizePending(opts *bind.FilterOpts, index []*big.Int) (*BranchBridgeBridgeFinalizePendingIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizePendingIterator{contract: _BranchBridge.contract, event: "BridgeFinalizePending", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizePending is a free log subscription operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalizePending(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalizePending, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFinalizePending)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
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

// ParseBridgeFinalizePending is a log parse operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFinalizePending(log types.Log) (*BranchBridgeBridgeFinalizePending, error) {
	event := new(BranchBridgeBridgeFinalizePending)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
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
	LocalToken      common.Address
	RemoteToken     common.Address
	LocalTokenRate  *big.Int
	RemoteTokenRate *big.Int
	SafetyLimit     *big.Int
	IsOrigin        bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xef354a4d67e88f3c656af891265f876391922dd3618eee9734a68467629be5e9.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, bool isOrigin)
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

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xef354a4d67e88f3c656af891265f876391922dd3618eee9734a68467629be5e9.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, bool isOrigin)
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xef354a4d67e88f3c656af891265f876391922dd3618eee9734a68467629be5e9.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, uint256 localTokenRate, uint256 remoteTokenRate, uint256 safetyLimit, bool isOrigin)
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
