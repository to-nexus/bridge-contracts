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

// CrossBridgeMetaData contains all meta data concerning the CrossBridge contract.
var CrossBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasExpiredPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"nexus_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lockPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualProcessPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nexus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"processExpiredPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"resetRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"_crossMintableERC20Factory\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"nexus_\",\"type\":\"address\"}],\"name\":\"setNexus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"setSafetyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20FactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"NexusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"SafetyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RegistryExistFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleAlreadyGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"StandardInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"StandardNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"StandardNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StandardNotExpiredSafeDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"47666cb1": "bridgeFeeStation()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"b7f3358d": "changeThreshold(uint8)",
		"0b43c02c": "clearPending(uint256)",
		"d016d625": "createToken(uint256,address,int256,uint256,string,uint8)",
		"8f517c17": "crossMintableERC20Factory()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"ae766389": "estimateFee(uint256,address,uint256)",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"b33eb36e": "getPendingArguments(uint256,uint256)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"814914b5": "getTokenPair(uint256,address)",
		"3d507c5e": "hasExpiredPending()",
		"91d14854": "hasRole(bytes32,address)",
		"5187599d": "initialize(uint8,address)",
		"91cf6d3e": "initializedAt()",
		"f17f6cb7": "lockPending(uint256,uint256)",
		"7f4ab9f5": "manualProcessPending(uint256,uint256)",
		"a3f5c1d2": "nexus()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"1089fd58": "processExpiredPending(uint256)",
		"52d1902d": "proxiableUUID()",
		"1e7bf215": "registerToken(uint256,bool,address,address,int256,uint256)",
		"715018a6": "renounceOwnership()",
		"2d87b7ee": "resetRole(bytes32,address[])",
		"3960e787": "retryFinalizeBridge(uint256,uint256)",
		"030372c3": "retryFinalizeBridgeBatch(uint256,uint256[])",
		"1a1aebbb": "setCrossMintableERC20Factory(address)",
		"54db0126": "setFeeStation(address)",
		"2670b817": "setNexus(address)",
		"bedb86fb": "setPause(bool)",
		"6160751f": "setPauseChain(uint256,bool)",
		"4d3f0da9": "setPauseToken(uint256,address,bool)",
		"d4bf502a": "setRole(bytes32,address[],bool)",
		"39a621f3": "setSafetyLimit(uint256,address,uint256)",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615fa86100eb5f395f613a150152615fa85ff3fe6080604052600436106102c2575f3560e01c80637f4ab9f51161016f578063ae766389116100d8578063d4bf502a11610092578063f17f6cb71161006d578063f17f6cb714610973578063f2fde38b14610992578063f4509643146109b1578063f698da25146109d0575f5ffd5b8063d4bf502a14610922578063d5717fc514610941578063d605665b14610960575f5ffd5b8063ae7663891461084b578063b33eb36e14610885578063b7f3358d146108b1578063bedb86fb146108d0578063cf56118e146108ef578063d016d62514610903575f5ffd5b806391cf6d3e1161012957806391cf6d3e1461077357806391d1485414610787578063a3246ad3146107a6578063a3f5c1d2146107d2578063ad3cb1cc146107ef578063ae6893f81461082c575f5ffd5b80637f4ab9f514610601578063814914b51461062057806384b0196e1461070657806388d67d6d1461072d5780638da5cb5b146107405780638f517c1714610754575f5ffd5b806347666cb11161022b57806354db0126116101e55780635fd262de116101c05780635fd262de1461058f5780636160751f146105a2578063715018a6146105c157806379214874146105d5575f5ffd5b806354db0126146105215780635b605f5c146105405780635c975abb1461056c575f5ffd5b806347666cb1146104645780634d3f0da91461049b5780634d5d0056146104ba5780634f1ef286146104cd5780635187599d146104e057806352d1902d146104ff575f5ffd5b80632670b8171161027c5780632670b817146103b05780632d87b7ee146103cf5780633960e787146103ee57806339a621f31461040d5780633d507c5e1461042c57806342cde4e814610443575f5ffd5b8063030372c3146102ed5780630b43c02c146103215780631089fd58146103405780631938e0f21461035f5780631a1aebbb146103725780631e7bf21514610391575f5ffd5b366102e957345f036102e7576040516365d14ce560e11b815260040160405180910390fd5b005b5f5ffd5b3480156102f8575f5ffd5b5061030c610307366004614d40565b6109e4565b60405190151581526020015b60405180910390f35b34801561032c575f5ffd5b506102e761033b366004614de1565b610a28565b34801561034b575f5ffd5b506102e761035a366004614de1565b610a84565b61030c61036d366004614ed4565b610d6c565b34801561037d575f5ffd5b506102e761038c366004614fa1565b611139565b34801561039c575f5ffd5b506102e76103ab366004614fc9565b6111e9565b3480156103bb575f5ffd5b506102e76103ca366004614fa1565b611499565b3480156103da575f5ffd5b506102e76103e936600461508e565b611511565b3480156103f9575f5ffd5b5061030c6104083660046150d1565b611543565b348015610418575f5ffd5b506102e76104273660046150f1565b611848565b348015610437575f5ffd5b5060335442111561030c565b34801561044e575f5ffd5b5060645460405160ff9091168152602001610318565b34801561046f575f5ffd5b50609654610483906001600160a01b031681565b6040516001600160a01b039091168152602001610318565b3480156104a6575f5ffd5b506102e76104b5366004615126565b61193b565b61030c6104c83660046151a9565b611a79565b6102e76104db3660046152c1565b611c6c565b3480156104eb575f5ffd5b506102e76104fa366004615303565b611c87565b34801561050a575f5ffd5b50610513611d95565b604051908152602001610318565b34801561052c575f5ffd5b506102e761053b366004614fa1565b611db0565b34801561054b575f5ffd5b5061055f61055a366004614de1565b611e48565b604051610318919061538b565b348015610577575f5ffd5b505f516020615f335f395f51905f525460ff1661030c565b61030c61059d3660046153d8565b611fcd565b3480156105ad575f5ffd5b506102e76105bc366004615460565b612036565b3480156105cc575f5ffd5b506102e761212b565b3480156105e0575f5ffd5b506105f46105ef366004614de1565b61213e565b60405161031891906154bd565b34801561060c575f5ffd5b5061030c61061b3660046150d1565b612157565b34801561062b575f5ffd5b506106f961063a3660046154cf565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526039602090815260408084206001600160a01b039384168552825292839020835160e08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b60405161031891906154f2565b348015610711575f5ffd5b5061071a6121f2565b604051610318979695949392919061552e565b61030c61073b36600461565b565b61229b565b34801561074b575f5ffd5b50610483612336565b34801561075f575f5ffd5b50603254610483906001600160a01b031681565b34801561077e575f5ffd5b50609854610513565b348015610792575f5ffd5b5061030c6107a13660046154cf565b612364565b3480156107b1575f5ffd5b506107c56107c0366004614de1565b61237b565b604051610318919061578b565b3480156107dd575f5ffd5b506097546001600160a01b0316610483565b3480156107fa575f5ffd5b5061081f604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161031891906157cb565b348015610837575f5ffd5b50610513610846366004614de1565b612394565b348015610856575f5ffd5b5061086a6108653660046150f1565b6123b0565b60408051938452602084019290925290820152606001610318565b348015610890575f5ffd5b506108a461089f3660046150d1565b61243f565b60405161031891906157dd565b3480156108bc575f5ffd5b506102e76108cb36600461586f565b6125e9565b3480156108db575f5ffd5b506102e76108ea366004615888565b61263a565b3480156108fa575f5ffd5b506105f461265b565b34801561090e575f5ffd5b5061048361091d3660046158a3565b61266c565b34801561092d575f5ffd5b506102e761093c36600461592c565b61277c565b34801561094c575f5ffd5b5061051361095b366004614de1565b6127b3565b6102e761096e366004615977565b6127cf565b34801561097e575f5ffd5b506102e761098d3660046150d1565b612962565b34801561099d575f5ffd5b506102e76109ac366004614fa1565b612a39565b3480156109bc575f5ffd5b506102e76109cb3660046154cf565b612a73565b3480156109db575f5ffd5b50610513612b58565b5f805b8251811015610a1c57610a1384848381518110610a0657610a06615a10565b6020026020010151611543565b506001016109e7565b50600190505b92915050565b610a30612b61565b5f818152603b60205260408120610a4690612b93565b90505f5b8151811015610a7f57610a7683838381518110610a6957610a69615a10565b6020026020010151612b9f565b50600101610a4a565b505050565b426033541015610a915750565b5f80610a9d6035612b93565b90505f5b8151811015610d66575f828281518110610abd57610abd615a10565b6020908102919091018101515f818152603790925260409091206003015490915060ff1615610aec5750610d5e565b5f818152603b60205260408120610b0290612b93565b90505f5b8151811015610d5a575f828281518110610b2257610b22615a10565b6020908102919091018101515f868152603c83526040808220838352909352828120835161012081019094528054606085019081526001820154608086015260028201546001600160a01b0390811660a087015260038301541660c0860152600482015460e0860152600582018054949650929493919284928491610100850191610bac90615a24565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd890615a24565b8015610c235780601f10610bfa57610100808354040283529160200191610c23565b820191905f5260205f20905b815481529060010190602001808311610c0657829003601f168201915b505050505081525050815260200160068201548152602001600782018054610c4a90615a24565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7690615a24565b8015610cc15780601f10610c9857610100808354040283529160200191610cc1565b820191905f5260205f20905b815481529060010190602001808311610ca457829003601f168201915b5050509190925250508151515f90815260396020908152604080832085518201516001600160a01b0316845290915290206001015491925050600160a01b900460ff16158015610d145750602081015115155b8015610d235750428160200151105b15610d5057610d328583612dac565b5088610d3d89615a70565b98508810610d5057505050505050505050565b5050600101610b06565b5050505b600101610aa1565b50505050565b5f610d75612e7b565b610d7d612eab565b610da5610d906060870160408801614fa1565b86355f90815260386020526040902090612ee2565b610db56060870160408801614fa1565b90610de4576040516353b2527760e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610e0f576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b505084355f9081526037602052604081206002018054600101908190559050806020870135808214610e5d57604051631351db4160e31b815260048101929092526024820152604401610ddb565b50610eef90507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020880135610e9960608a0160408b01614fa1565b610ea960808b0160608c01614fa1565b60808b0135610ebb60a08d018d615a88565b604051602001610ed19796959493929190615af2565b60405160208183030381529060405280519060200120868686612f03565b85355f90815260396020526040808220606091839182918290610f17908d8701908e01614fa1565b6001600160a01b03908116825260208083019390935260409182015f20825160e08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c0830152909150610fc8576040518060400160405280600c81526020016b1d1bdad95b881c185d5cd95960a21b8152509250611051565b608081015115801590610fe2575089608001358160800151105b1561101c57604051806040016040528060118152602001701bdd995c881cd859995d1e481b1a5b5a5d607a1b815250925060019150611051565b61104b8a3561103160608d0160408e01614fa1565b61104160808e0160608f01614fa1565b8d60800135613153565b90945092505b5082156110d95761106860808a0160608b01614fa1565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b6110a960608e0160408f01614fa1565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4611113565b6110e489838361324a565b60405160208a0135907f40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2905f90a25b600194505050505061113160015f516020615f535f395f51905f5255565b949350505050565b611141612b61565b6032546001600160a01b0316801561117857604051639ad61dbd60e01b81526001600160a01b039091166004820152602401610ddb565b506001600160a01b0381166111a057604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517f18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9905f90a250565b6111f1612b61565b6111fc60358761340a565b1561125957604080516080810182528781525f6020808301828152838501838152606085018481528c855260379093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b03841661128057604051636ca1fdd760e01b815260040160405180910390fd5b5f8681526038602052604090206112979085613415565b84906112c2576040516311dd05f360e31b81526001600160a01b039091166004820152602401610ddb565b506040518060e00160405280856001600160a01b03168152602001846001600160a01b031681526020015f1515815260200186151581526020018281526020015f81526020015f81525060395f8881526020019081526020015f205f866001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c08201518160040155905050600182138061140457505f1982125b1561142e575f868152603a602090815260408083206001600160a01b038816845290915290208290555b826001600160a01b0316846001600160a01b0316877fffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a6685858a6040516114899392919092835260208301919091521515604082015260600190565b60405180910390a4505050505050565b6114a1612b61565b6001600160a01b0381166114c857604051630fb9363360e41b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0383169081179091556040517f489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f905f90a250565b6115338261152d5f5f8681526020019081526020015f20612b93565b5f61277c565b61153f8282600161277c565b5050565b5f61154c612e7b565b611554612eab565b5f838152603b6020526040902061156b9083613429565b829061158d57604051630e0b575b60e21b8152600401610ddb91815260200190565b505f838152603c60209081526040808320858452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061160990615a24565b80601f016020809104026020016040519081016040528092919081815260200182805461163590615a24565b80156116805780601f1061165757610100808354040283529160200191611680565b820191905f5260205f20905b81548152906001019060200180831161166357829003601f168201915b5050505050815250508152602001600682015481526020016007820180546116a790615a24565b80601f01602080910402602001604051908101604052809291908181526020018280546116d390615a24565b801561171e5780601f106116f55761010080835404028352916020019161171e565b820191905f5260205f20905b81548152906001019060200180831161170157829003601f168201915b505050919092525050505f85815260396020908152604080832084518201516001600160a01b03908116855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b8504811615801585850152600160a81b909504161515606084015260028101546080840152600381015460a08401526004015460c08301528351015192935091906117df576040516338384f6f60e11b81526001600160a01b039091166004820152602401610ddb565b50602082015115806117f45750428260200151105b826020015142909161182257604051636d60854360e01b815260048101929092526024820152604401610ddb565b505061182e8585612dac565b92505050610a2260015f516020615f535f395f51905f5255565b61185a6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916118955760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f8381526038602052604090206118ae9083612ee2565b82906118d95760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8381526039602090815260408083206001600160a01b038616808552908352928190206002018490555183815285917fd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c891015b60405180910390a3505050565b61194d6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916119885760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f8381526038602052604090206119a19083612ee2565b82906119cc5760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8381526039602090815260408083206001600160a01b038616845290915290206001015460ff600160a01b90910416151581151514610a7f575f8381526039602090815260408083206001600160a01b0386168085529252918290206001018054841515600160a01b0260ff60a01b19909116179055905184907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061192e90851515815260200190565b5f611a82612e7b565b8989611a8e8282613440565b611a96612eab565b611aa36020850185614fa1565b6001600160a01b038c81169116148b611abf6020870187614fa1565b9091611af15760405163a9212f4360e01b81526001600160a01b03928316600482015291166024820152604401610ddb565b5050611b008c8c8b8b8b613516565b909850965086611b10898b615b3f565b611b1a9190615b3f565b6040850135101587611b2c8a8c615b3f565b611b369190615b3f565b85604001359091611b63576040516365efbabf60e11b815260048101929092526024820152604401610ddb565b505f9050611b776040860160208701614fa1565b3060408701356060880135611b9260a08a0160808b0161586f565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a086013560c482015260c086013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031663d505accf60e01b179052909150611c2090611c1990870187614fa1565b5f836135fa565b50611c438c8c611c366040880160208901614fa1565b8d8d8d8d60018e8e613723565b60019250611c5d60015f516020615f535f395f51905f5255565b50509998505050505050505050565b611c74613858565b611c7d826138be565b61153f82826138c6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611ccb5750825b90505f826001600160401b03166001148015611ce65750303b155b905081158015611cf4575080155b15611d125760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611d3c57845460ff60401b1916600160401b1785555b611d468787613982565b8315611d8c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f611d9e613a0a565b505f516020615f135f395f51905f5290565b611dc26420a226a4a760d91b33612364565b6420a226a4a760d91b339091611dfd5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b5050609680546001600160a01b0319166001600160a01b0383169081179091556040517f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444905f90a250565b5f81815260386020526040812060609190611e6290612b93565b90505f81516001600160401b03811115611e7e57611e7e614cb2565b604051908082528060200260200182016040528015611ee357816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f19909201910181611e9c5790505b5090505f5b8251811015611fc55760395f8681526020019081526020015f205f848381518110611f1557611f15615a10565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c08201528251839083908110611fb257611fb2615a10565b6020908102919091010152600101611ee8565b509392505050565b5f611fd6612e7b565b8888611fe28282613440565b611fea612eab565b611ff78b8b8a8a8a613516565b909750955061200e8b8b338c8c8c8c5f8d8d613723565b6001925061202860015f516020615f535f395f51905f5255565b505098975050505050505050565b6120486420a226a4a760d91b33612364565b6420a226a4a760d91b3390916120835760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b506120919050603583613429565b82906120b35760405163ac7dbbfd60e01b8152600401610ddb91815260200190565b505f8281526037602052604090206003015460ff1615158115151461153f575f82815260376020908152604091829020600301805460ff1916841515908117909155915191825283917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a25050565b612133612b61565b61213c5f613a53565b565b5f818152603b60205260409020606090610a2290612b93565b5f61216a6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916121a55760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f838152603b602052604090206121be9083613429565b82906121e057604051630e0b575b60e21b8152600401610ddb91815260200190565b506121eb8383612dac565b9392505050565b5f60608082808083815f516020615ef35f395f51905f52805490915015801561221d57506001810154155b6122615760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610ddb565b612269613ac3565b612271613b83565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015612329576123208787838181106122bb576122bb615a10565b90506020028101906122cd9190615b52565b8683815181106122df576122df615a10565b60200260200101518684815181106122f9576122f9615a10565b602002602001015186858151811061231357612313615a10565b6020026020010151610d6c565b5060010161229e565b5060019695505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f8281526020819052604081206121eb9083612ee2565b5f818152602081905260409020606090610a2290612b93565b5f818152603760205260408120600190810154610a2291615b3f565b60965460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa15801561240c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124309190615b70565b91989097509095509350505050565b612447614bff565b5f838152603c6020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e08401526005820180548492918491610100850191906124c390615a24565b80601f01602080910402602001604051908101604052809291908181526020018280546124ef90615a24565b801561253a5780601f106125115761010080835404028352916020019161253a565b820191905f5260205f20905b81548152906001019060200180831161251d57829003601f168201915b50505050508152505081526020016006820154815260200160078201805461256190615a24565b80601f016020809104026020016040519081016040528092919081815260200182805461258d90615a24565b80156125d85780601f106125af576101008083540402835291602001916125d8565b820191905f5260205f20905b8154815290600101906020018083116125bb57829003601f168201915b505050505081525050905092915050565b6125f1612b61565b6064805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b612642612b61565b801561265357612650613bc1565b50565b612650613c1d565b60606126676035612b93565b905090565b6032545f906001600160a01b0316612697576040516315aeca0d60e11b815260040160405180910390fd5b6032546040516bffffffffffffffffffffffff19606089901b1660208201526001600160a01b0390911690634804a0419060340160405160208183030381529060405280519060200120856040516020016126f29190615bb2565b60405160208183030381529060405286866040518563ffffffff1660e01b81526004016127229493929190615bd3565b6020604051808303815f875af115801561273e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127629190615c12565b9050612772875f838989896111e9565b9695505050505050565b5f5b8251811015610d66576127ab8484838151811061279d5761279d615a10565b602002602001015184613c62565b60010161277e565b5f81815260376020526040812060020154610a22906001615b3f565b8281146127ef5760405163214485c960e01b815260040160405180910390fd5b5f5b8381101561295b5761295285858381811061280e5761280e615a10565b90506020028101906128209190615c2d565b3586868481811061283357612833615a10565b90506020028101906128459190615c2d565b612856906040810190602001614fa1565b87878581811061286857612868615a10565b905060200281019061287a9190615c2d565b61288b906060810190604001614fa1565b88888681811061289d5761289d615a10565b90506020028101906128af9190615c2d565b606001358989878181106128c5576128c5615a10565b90506020028101906128d79190615c2d565b608001358a8a888181106128ed576128ed615a10565b90506020028101906128ff9190615c2d565b60a001358b8b8981811061291557612915615a10565b90506020028101906129279190615c2d565b6129359060c0810190615a88565b8b8b8b81811061294757612947615a10565b905060e00201611a79565b506001016127f1565b5050505050565b6129746420a226a4a760d91b33612364565b6420a226a4a760d91b3390916129af5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f828152603b602052604090206129c89082613429565b81906129ea57604051630e0b575b60e21b8152600401610ddb91815260200190565b505f828152603c602090815260408083208484529091528082205f1960069091015551829184917fa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f69190a35050565b612a41612b61565b6001600160a01b038116612a6a57604051631e4fbdf760e01b81525f6004820152602401610ddb565b61265081613a53565b612a7b612b61565b5f828152603860205260409020612a929082613d92565b8190612abd5760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8281526039602090815260408083206001600160a01b03851680855290835281842080546001600160a01b03191681556001810180546001600160b01b03191690556002810185905560038101859055600401849055858452603a835281842081855290925280832083905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f612667613da6565b33612b6a612336565b6001600160a01b03161461213c5760405163118cdaa760e01b8152336004820152602401610ddb565b60605f6121eb83613daf565b612ba7614c25565b5f838152603b60205260409020612bbe9083613e08565b8290612be057604051637f11bea960e01b8152600401610ddb91815260200190565b505f838152603c60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190612c5890615a24565b80601f0160208091040260200160405190810160405280929190818152602001828054612c8490615a24565b8015612ccf5780601f10612ca657610100808354040283529160200191612ccf565b820191905f5260205f20905b815481529060010190602001808311612cb257829003601f168201915b505050919092525050505f848152603960209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615612d32578160800151816004015f828254612d2c9190615c41565b90915550505b5f848152603c602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181612d8b6005830182614c68565b5050600682015f9055600782015f612da39190614c68565b50505092915050565b5f5f612db88484612b9f565b90505f5f612dd7835f0151846040015185606001518660800151613153565b91509150818190612dfb5760405162461bcd60e51b8152600401610ddb91906157cb565b5082606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8660400151876080015142604051612e67939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b5f516020615f335f395f51905f525460ff161561213c5760405163d93c066560e01b815260040160405180910390fd5b5f516020615f535f395f51905f52805460011901612edc57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f90815260018301602052604081205415156121eb565b8251825181148015612f155750815181145b835183518392612f49576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610ddb565b505060645482915060ff16811015612f7757604051631aebd17960e11b8152600401610ddb91815260200190565b505f80826001600160401b03811115612f9257612f92614cb2565b604051908082528060200260200182016040528015612fbb578160200160208202803683370190505b5090505f5b8381101561311d575f61302b888381518110612fde57612fde615a10565b6020026020010151888481518110612ff857612ff8615a10565b602002602001015188858151811061301257613012615a10565b60200260200101516130238d613e13565b929190613e3f565b9050613043682b20a624a220aa27a960b91b82612364565b682b20a624a220aa27a960b91b8290916130825760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b505f9050805b858110156130d3578481815181106130a2576130a2615a10565b60200260200101516001600160a01b0316836001600160a01b0316036130cb57600191506130d3565b600101613088565b508061311357818486815181106130ec576130ec615a10565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b5050600101612fc0565b50606454829060ff1681101561314957604051631aebd17960e11b8152600401610ddb91815260200190565b5050505050505050565b5f848152603a602090815260408083206001600160a01b038716845290915281205460609060018113156131925761318b8185615c54565b93506131b1565b5f198112156131b1576131a481615c6b565b6131ae9085615c99565b93505b5f196001600160a01b038716016131e4576131db858560405180602001604052805f8152506135fa565b6001925061323f565b831561323f575f8781526039602090815260408083206001600160a01b038a168452909152902060010154600160a81b900460ff16156132345761322a87878787613e6b565b9250925050613241565b61322a868686613f61565b505b94509492505050565b82355f908152603b602090815260409091206132689185013561340a565b83602001359061328e576040516307a5c91d60e31b8152600401610ddb91815260200190565b505f81156132b7576034546132a39042615b3f565b9050603454426132b39190615b3f565b6033555b6040518060600160405280856132cc90615cac565b81526020808201849052604091820186905286355f908152603c82528281208883013582528252829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a08101518290600582019061335e9082615d76565b505050602082015160068201556040820151600782019061337f9082615d76565b50505083355f908152603960205260408082209082906133a59060608901908901614fa1565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff161561295b578460800135816004015f8282546133eb9190615b3f565b90915550505050505050565b60015f516020615f535f395f51905f5255565b5f6121eb8383614043565b5f6121eb836001600160a01b038416614043565b5f81815260018301602052604081205415156121eb565b5f8281526038602052604090206134579082612ee2565b81906134825760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f82815260376020526040902060030154829060ff16156134ba57604051636fc24b7960e11b8152600401610ddb91815260200190565b505f8281526039602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610a7f576040516338384f6f60e11b81526001600160a01b039091166004820152602401610ddb565b6096545f9081906001600160a01b031661353457505f9050806135f0565b60965460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa15801561358a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135ae9190615b70565b909450925090508086108015906135c55750828510155b80156135d15750818410155b6135ee5760405163dc48c7f160e01b815260040160405180910390fd5b505b9550959350505050565b814747821115613626576040516371e4d07760e11b815260048101929092526024820152604401610ddb565b50505f5f846001600160a01b031684846040516136439190615e30565b5f6040518083038185875af1925050503d805f811461367d576040519150601f19603f3d011682016040523d82523d5f602084013e613682565b606091505b5091509150816136b55780511561369c5780518082602001fd5b604051630daf67c960e11b815260040160405180910390fd5b835f0361295b5780515f036136f357846001600160a01b03163b5f036136ee57604051630daf67c960e11b815260040160405180910390fd5b61295b565b602081015160018115151461371b57604051630daf67c960e11b815260040160405180910390fd5b505050505050565b6137398a8a8a89613734898b615b3f565b61408f565b5f8a815260376020526040812060019081018054909101908190558190915060395f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b031690508a6001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7848e8e8e428d8d8d6040516137eb989796959493929190615e3b565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a604051613842929190918252602082015260400190565b60405180910390a4505050505050505050505050565b30610101620956d760881b0114806138a05750610101620956d760881b016138945f516020615f135f395f51905f52546001600160a01b031690565b6001600160a01b031614155b1561213c5760405163703e46dd60e11b815260040160405180910390fd5b612650612b61565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613920575060408051601f3d908101601f1916820190925261391d91810190615e85565b60015b61394857604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610ddb565b5f516020615f135f395f51905f52811461397857604051632a87526960e21b815260048101829052602401610ddb565b610a7f83836142e0565b61398a614335565b6139933361437e565b61399b61438f565b6139a3614397565b6139ab6143a7565b6139b4826143b7565b6139bc61441d565b6001600160a01b0381166139e357604051630fb9363360e41b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b03929092169190911790555043609855565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461213c5760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615ef35f395f51905f5291613b0190615a24565b80601f0160208091040260200160405190810160405280929190818152602001828054613b2d90615a24565b8015613b785780601f10613b4f57610100808354040283529160200191613b78565b820191905f5260205f20905b815481529060010190602001808311613b5b57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615ef35f395f51905f5291613b0190615a24565b613bc9612e7b565b5f516020615f335f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200161262f565b613c25614449565b5f516020615f335f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613c05565b613c6a612b61565b6001600160a01b038216613c9157604051635989b9d360e01b815260040160405180910390fd5b82613caf57604051630e1b39f960e31b815260040160405180910390fd5b8015613d06575f838152602081905260409020613ccc9083613415565b83839091613cff57604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610ddb565b5050613d53565b5f838152602081905260409020613d1d9083613d92565b83839091613d5057604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610ddb565b50505b80151583836001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f360405160405180910390a4505050565b5f6121eb836001600160a01b038416614478565b5f61266761455b565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613dfc57602002820191905f5260205f20905b815481526020019060010190808311613de8575b50505050509050919050565b5f6121eb8383614478565b5f610a22613e1f613da6565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613e4f888888886145ce565b925092509250613e5f8282614696565b50909695505050505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613eda575060408051601f3d908101601f19168201909252613ed791810190615e9c565b60015b613f14573d808015613f07576040519150601f19603f3d011682016040523d82523d5f602084013e613f0c565b606091505b509050613241565b8092508015613f2d57613f2887878661474e565b61323f565b6040518060400160405280600f81526020016e1d1c985b9cd9995c8819985a5b1959608a1b81525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613fd0575060408051601f3d908101601f19168201909252613fcd91810190615e9c565b60015b61400a573d808015613ffd576040519150601f19603f3d011682016040523d82523d5f602084013e614002565b606091505b50905061403b565b80925080614039576040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b81525091505b505b935093915050565b5f81815260018301602052604081205461408857508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a22565b505f610a22565b5f858152603a602090815260408083206001600160a01b038816845290915290205460018113156140fa576140c48184615eb7565b85908490156140f7576040516348d53e0760e01b81526001600160a01b0390921660048301526024820152604401610ddb565b50505b5f196001600160a01b0386160161417d576141158284615b3f565b34146141218385615b3f565b34909161414a576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b505081156141785760975460408051602081019091525f8152614178916001600160a01b03169084906135fa565b61371b565b5f3480156141a7576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b506141cb905084306141b98587615b3f565b6001600160a01b0389169291906147f9565b81156141eb576097546141eb906001600160a01b03878116911684614860565b5f8681526039602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff161561422a57614178868685614891565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af1158015614274573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142989190615e9c565b8585859091926142d55760405163601bc95b60e11b81526001600160a01b0393841660048201529290911660248301526044820152606401610ddb565b505050505050505050565b6142e9826148cf565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561432d57610a7f8282614932565b61153f6149a4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661213c57604051631afcd79f60e31b815260040160405180910390fd5b614386614335565b612650816149c3565b61213c614335565b61439f614335565b61213c6149cb565b6143af614335565b61213c6149eb565b6143bf614335565b614407604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506149f3565b6064805460ff191660ff92909216919091179055565b614425614335565b6144406420a226a4a760d91b614439612336565b6001613c62565b62015180603455565b5f516020615f335f395f51905f525460ff1661213c57604051638dfc202b60e01b815260040160405180910390fd5b5f8181526001830160205260408120548015614552575f61449a600183615c41565b85549091505f906144ad90600190615c41565b905080821461450c575f865f0182815481106144cb576144cb615a10565b905f5260205f200154905080875f0184815481106144eb576144eb615a10565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061451d5761451d615eca565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a22565b5f915050610a22565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f614585614a05565b61458d614a6d565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561460757505f9150600390508261468c565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614658573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661468357505f92506001915082905061468c565b92505f91508190505b9450945094915050565b5f8260038111156146a9576146a9615ede565b036146b2575050565b60018260038111156146c6576146c6615ede565b036146e45760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156146f8576146f8615ede565b036147195760405163fce698f760e01b815260048101829052602401610ddb565b600382600381111561472d5761472d615ede565b0361153f576040516335e2f38360e21b815260048101829052602401610ddb565b5f8381526039602090815260408083206001600160a01b03861684529091529020600481015461477e9083615b3f565b6003820154600483015486928692869290918210156147d6576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610ddb565b505050505081816003015f8282546147ee9190615c41565b909155505050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610d669186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614aaf565b6040516001600160a01b03838116602483015260448201839052610a7f91859182169063a9059cbb9060640161482e565b5f8381526039602090815260408083206001600160a01b0386168452909152812060030180548392906148c5908490615b3f565b9091555050505050565b806001600160a01b03163b5f0361490457604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610ddb565b5f516020615f135f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161494e9190615e30565b5f60405180830381855af49150503d805f8114614986576040519150601f19603f3d011682016040523d82523d5f602084013e61498b565b606091505b509150915061499b858383614b1b565b95945050505050565b341561213c5760405163b398979f60e01b815260040160405180910390fd5b612a41614335565b6149d3614335565b5f516020615f335f395f51905f52805460ff19169055565b6133f7614335565b6149fb614335565b61153f8282614b77565b5f5f516020615ef35f395f51905f5281614a1d613ac3565b805190915015614a3557805160209091012092915050565b81548015614a44579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615ef35f395f51905f5281614a85613b83565b805190915015614a9d57805160209091012092915050565b60018201548015614a44579392505050565b5f5f60205f8451602086015f885af180614ace576040513d5f823e3d81fd5b50505f513d91508115614ae5578060011415614af2565b6001600160a01b0384163b155b15610d6657604051635274afe760e01b81526001600160a01b0385166004820152602401610ddb565b606082614b3057614b2b82614bd6565b6121eb565b8151158015614b4757506001600160a01b0384163b155b15614b7057604051639996b31560e01b81526001600160a01b0385166004820152602401610ddb565b5092915050565b614b7f614335565b5f516020615ef35f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614bb88482615d76565b5060038101614bc78382615d76565b505f8082556001909101555050565b805115614be65780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614c12614c25565b81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614c7490615a24565b5f825580601f10614c83575050565b601f0160209004905f5260205f209081019061265091905b80821115614cae575f8155600101614c9b565b5090565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614ce857614ce8614cb2565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614d1657614d16614cb2565b604052919050565b5f6001600160401b03821115614d3657614d36614cb2565b5060051b60200190565b5f5f60408385031215614d51575f5ffd5b8235915060208301356001600160401b03811115614d6d575f5ffd5b8301601f81018513614d7d575f5ffd5b8035614d90614d8b82614d1e565b614cee565b8082825260208201915060208360051b850101925087831115614db1575f5ffd5b6020840193505b82841015614dd3578335825260209384019390910190614db8565b809450505050509250929050565b5f60208284031215614df1575f5ffd5b5035919050565b803560ff81168114614e08575f5ffd5b919050565b5f82601f830112614e1c575f5ffd5b8135614e2a614d8b82614d1e565b8082825260208201915060208360051b860101925085831115614e4b575f5ffd5b602085015b83811015614e6f57614e6181614df8565b835260209283019201614e50565b5095945050505050565b5f82601f830112614e88575f5ffd5b8135614e96614d8b82614d1e565b8082825260208201915060208360051b860101925085831115614eb7575f5ffd5b602085015b83811015614e6f578035835260209283019201614ebc565b5f5f5f5f60808587031215614ee7575f5ffd5b84356001600160401b03811115614efc575f5ffd5b850160c08188031215614f0d575f5ffd5b935060208501356001600160401b03811115614f27575f5ffd5b614f3387828801614e0d565b93505060408501356001600160401b03811115614f4e575f5ffd5b614f5a87828801614e79565b92505060608501356001600160401b03811115614f75575f5ffd5b614f8187828801614e79565b91505092959194509250565b6001600160a01b0381168114612650575f5ffd5b5f60208284031215614fb1575f5ffd5b81356121eb81614f8d565b8015158114612650575f5ffd5b5f5f5f5f5f5f60c08789031215614fde575f5ffd5b863595506020870135614ff081614fbc565b9450604087013561500081614f8d565b9350606087013561501081614f8d565b9598949750929560808101359460a0909101359350915050565b5f82601f830112615039575f5ffd5b8135615047614d8b82614d1e565b8082825260208201915060208360051b860101925085831115615068575f5ffd5b602085015b83811015614e6f57803561508081614f8d565b83526020928301920161506d565b5f5f6040838503121561509f575f5ffd5b8235915060208301356001600160401b038111156150bb575f5ffd5b6150c78582860161502a565b9150509250929050565b5f5f604083850312156150e2575f5ffd5b50508035926020909101359150565b5f5f5f60608486031215615103575f5ffd5b83359250602084013561511581614f8d565b929592945050506040919091013590565b5f5f5f60608486031215615138575f5ffd5b83359250602084013561514a81614f8d565b9150604084013561515a81614fbc565b809150509250925092565b5f5f83601f840112615175575f5ffd5b5081356001600160401b0381111561518b575f5ffd5b6020830191508360208285010111156151a2575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c08112156151c3575f5ffd5b8a35995060208b01356151d581614f8d565b985060408b01356151e581614f8d565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115615214575f5ffd5b6152208d828e01615165565b90955093505060e060df1982011215615237575f5ffd5b5060e08a0190509295985092959850929598565b5f5f6001600160401b0384111561526457615264614cb2565b50601f8301601f191660200161527981614cee565b91505082815283838301111561528d575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126152b2575f5ffd5b6121eb8383356020850161524b565b5f5f604083850312156152d2575f5ffd5b82356152dd81614f8d565b915060208301356001600160401b038111156152f7575f5ffd5b6150c7858286016152a3565b5f5f60408385031215615314575f5ffd5b61531d83614df8565b9150602083013561532d81614f8d565b809150509250929050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b818110156153cd576153b7838551615338565b6020939093019260e092909201916001016153a4565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b0312156153ef575f5ffd5b88359750602089013561540181614f8d565b9650604089013561541181614f8d565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115615440575f5ffd5b61544c8b828c01615165565b999c989b5096995094979396929594505050565b5f5f60408385031215615471575f5ffd5b82359150602083013561532d81614fbc565b5f8151808452602084019350602083015f5b828110156154b3578151865260209586019590910190600101615495565b5093949350505050565b602081525f6121eb6020830184615483565b5f5f604083850312156154e0575f5ffd5b82359150602083013561532d81614f8d565b60e08101610a228284615338565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61554c60e0830189615500565b828103604084015261555e8189615500565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061558f8185615483565b9a9950505050505050505050565b5f5f83601f8401126155ad575f5ffd5b5081356001600160401b038111156155c3575f5ffd5b6020830191508360208260051b85010111156151a2575f5ffd5b5f82601f8301126155ec575f5ffd5b81356155fa614d8b82614d1e565b8082825260208201915060208360051b86010192508583111561561b575f5ffd5b602085015b83811015614e6f5780356001600160401b0381111561563d575f5ffd5b61564c886020838a0101614e79565b84525060209283019201615620565b5f5f5f5f5f6080868803121561566f575f5ffd5b85356001600160401b03811115615684575f5ffd5b6156908882890161559d565b90965094505060208601356001600160401b038111156156ae575f5ffd5b8601601f810188136156be575f5ffd5b80356156cc614d8b82614d1e565b8082825260208201915060208360051b85010192508a8311156156ed575f5ffd5b602084015b8381101561572d5780356001600160401b0381111561570f575f5ffd5b61571e8d602083890101614e0d565b845250602092830192016156f2565b50955050505060408601356001600160401b0381111561574b575f5ffd5b615757888289016155dd565b92505060608601356001600160401b03811115615772575f5ffd5b61577e888289016155dd565b9150509295509295909350565b602080825282518282018190525f918401906040840190835b818110156153cd5783516001600160a01b03168352602093840193909201916001016157a4565b602081525f6121eb6020830184615500565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c0610120840152615848610140840182615500565b9050602084015160408401526040840151601f1984830301606085015261499b8282615500565b5f6020828403121561587f575f5ffd5b6121eb82614df8565b5f60208284031215615898575f5ffd5b81356121eb81614fbc565b5f5f5f5f5f5f60c087890312156158b8575f5ffd5b8635955060208701356158ca81614f8d565b9450604087013593506060870135925060808701356001600160401b038111156158f2575f5ffd5b8701601f81018913615902575f5ffd5b6159118982356020840161524b565b92505061592060a08801614df8565b90509295509295509295565b5f5f5f6060848603121561593e575f5ffd5b8335925060208401356001600160401b0381111561595a575f5ffd5b6159668682870161502a565b925050604084013561515a81614fbc565b5f5f5f5f6040858703121561598a575f5ffd5b84356001600160401b0381111561599f575f5ffd5b6159ab8782880161559d565b90955093505060208501356001600160401b038111156159c9575f5ffd5b8501601f810187136159d9575f5ffd5b80356001600160401b038111156159ee575f5ffd5b87602060e083028401011115615a02575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680615a3857607f821691505b602082108103615a5657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201615a8157615a81615a5c565b5060010190565b5f5f8335601e19843603018112615a9d575f5ffd5b8301803591506001600160401b03821115615ab6575f5ffd5b6020019150368190038213156151a2575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615b329083018486615aca565b9998505050505050505050565b80820180821115610a2257610a22615a5c565b5f823560be19833603018112615b66575f5ffd5b9190910192915050565b5f5f5f60608486031215615b82575f5ffd5b5050815160208301516040909301519094929350919050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f6121eb600d830184615b9b565b848152608060208201525f615beb6080830186615500565b8281036040840152615bfd8186615500565b91505060ff8316606083015295945050505050565b5f60208284031215615c22575f5ffd5b81516121eb81614f8d565b5f823560de19833603018112615b66575f5ffd5b81810381811115610a2257610a22615a5c565b8082028115828204841417610a2257610a22615a5c565b5f600160ff1b8201615c7f57615c7f615a5c565b505f0390565b634e487b7160e01b5f52601260045260245ffd5b5f82615ca757615ca7615c85565b500490565b5f60c08236031215615cbc575f5ffd5b615cc4614cc6565b82358152602080840135908201526040830135615ce081614f8d565b60408201526060830135615cf381614f8d565b60608201526080838101359082015260a08301356001600160401b03811115615d1a575f5ffd5b615d26368286016152a3565b60a08301525092915050565b601f821115610a7f57805f5260205f20601f840160051c81016020851015615d575750805b601f840160051c820191505b8181101561295b575f8155600101615d63565b81516001600160401b03811115615d8f57615d8f614cb2565b615da381615d9d8454615a24565b84615d32565b6020601f821160018114615dd5575f8315615dbe5750848201515b5f19600385901b1c1916600184901b17845561295b565b5f84815260208120601f198516915b82811015615e045787850151825560209485019460019092019101615de4565b5084821015615e2157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6121eb8284615b9b565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f9061558f9083018486615aca565b5f60208284031215615e95575f5ffd5b5051919050565b5f60208284031215615eac575f5ffd5b81516121eb81614fbc565b5f82615ec557615ec5615c85565b500690565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220e78cb77e6bf0cbf74f76f1149ab952699d3568bbfd99c860dab00147fea7bfc664736f6c634300081c0033",
}

// CrossBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeMetaData.ABI instead.
var CrossBridgeABI = CrossBridgeMetaData.ABI

// Deprecated: Use CrossBridgeMetaData.Sigs instead.
// CrossBridgeFuncSigs maps the 4-byte function signature to its string representation.
var CrossBridgeFuncSigs = CrossBridgeMetaData.Sigs

// CrossBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossBridgeMetaData.Bin instead.
var CrossBridgeBin = CrossBridgeMetaData.Bin

// DeployCrossBridge deploys a new Ethereum contract, binding an instance of CrossBridge to it.
func DeployCrossBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossBridge, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// CrossBridge is an auto generated Go binding around an Ethereum contract.
type CrossBridge struct {
	CrossBridgeCaller     // Read-only binding to the contract
	CrossBridgeTransactor // Write-only binding to the contract
	CrossBridgeFilterer   // Log filterer for contract events
}

// CrossBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossBridgeSession struct {
	Contract     *CrossBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossBridgeCallerSession struct {
	Contract *CrossBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrossBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossBridgeTransactorSession struct {
	Contract     *CrossBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossBridgeRaw struct {
	Contract *CrossBridge // Generic contract binding to access the raw methods on
}

// CrossBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossBridgeCallerRaw struct {
	Contract *CrossBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossBridgeTransactorRaw struct {
	Contract *CrossBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossBridge creates a new instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridge(address common.Address, backend bind.ContractBackend) (*CrossBridge, error) {
	contract, err := bindCrossBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// NewCrossBridgeCaller creates a new read-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeCaller(address common.Address, caller bind.ContractCaller) (*CrossBridgeCaller, error) {
	contract, err := bindCrossBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCaller{contract: contract}, nil
}

// NewCrossBridgeTransactor creates a new write-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossBridgeTransactor, error) {
	contract, err := bindCrossBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTransactor{contract: contract}, nil
}

// NewCrossBridgeFilterer creates a new log filterer instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossBridgeFilterer, error) {
	contract, err := bindCrossBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeFilterer{contract: contract}, nil
}

// bindCrossBridge binds a generic wrapper to an already deployed contract.
func bindCrossBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.CrossBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeFeeStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeFeeStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeFeeStation() (common.Address, error) {
	return _CrossBridge.Contract.BridgeFeeStation(&_CrossBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeFeeStation() (common.Address, error) {
	return _CrossBridge.Contract.BridgeFeeStation(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_CrossBridge *CrossBridgeCaller) CrossMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_CrossBridge *CrossBridgeSession) CrossMintableERC20Factory() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Factory(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) CrossMintableERC20Factory() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Factory(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossBridge *CrossBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

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
func (_CrossBridge *CrossBridgeSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _CrossBridge.Contract.EstimateFee(&_CrossBridge.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _CrossBridge.Contract.EstimateFee(&_CrossBridge.CallOpts, remoteChainID, token, value)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_CrossBridge *CrossBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasExpiredPending(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasExpiredPending")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasExpiredPending() (bool, error) {
	return _CrossBridge.Contract.HasExpiredPending(&_CrossBridge.CallOpts)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasExpiredPending() (bool, error) {
	return _CrossBridge.Contract.HasExpiredPending(&_CrossBridge.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Nexus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "nexus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_CrossBridge *CrossBridgeSession) Nexus() (common.Address, error) {
	return _CrossBridge.Contract.Nexus(&_CrossBridge.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Nexus() (common.Address, error) {
	return _CrossBridge.Contract.Nexus(&_CrossBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeSession) Owner() (common.Address, error) {
	return _CrossBridge.Contract.Owner(&_CrossBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Owner() (common.Address, error) {
	return _CrossBridge.Contract.Owner(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCallerSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) BridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "bridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeTransactor) ClearPending(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "clearPending", remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ClearPending(&_CrossBridge.TransactOpts, remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ClearPending(&_CrossBridge.TransactOpts, remoteChainID)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_CrossBridge *CrossBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "initialize", _threshold, nexus_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_CrossBridge *CrossBridgeSession) Initialize(_threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, _threshold, nexus_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) Initialize(_threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, _threshold, nexus_)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) LockPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "lockPending", remoteChainID, index)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) LockPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.LockPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) LockPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.LockPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactor) ManualProcessPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualProcessPending", remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualProcessPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualProcessPending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeTransactor) ProcessExpiredPending(opts *bind.TransactOpts, maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "processExpiredPending", maxCount)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeSession) ProcessExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ProcessExpiredPending(&_CrossBridge.TransactOpts, maxCount)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ProcessExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ProcessExpiredPending(&_CrossBridge.TransactOpts, maxCount)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceOwnership(&_CrossBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossBridge *CrossBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceOwnership(&_CrossBridge.TransactOpts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeTransactor) ResetRole(opts *bind.TransactOpts, role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "resetRole", role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ResetRole(&_CrossBridge.TransactOpts, role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ResetRole(&_CrossBridge.TransactOpts, role, newAccounts)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "retryFinalizeBridge", remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RetryFinalizeBridge(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RetryFinalizeBridge(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_CrossBridge *CrossBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_CrossBridge *CrossBridgeSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RetryFinalizeBridgeBatch(&_CrossBridge.TransactOpts, remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RetryFinalizeBridgeBatch(&_CrossBridge.TransactOpts, remoteChainID, indexes)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossMintableERC20Factory(opts *bind.TransactOpts, _crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossMintableERC20Factory", _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Factory(&_CrossBridge.TransactOpts, _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Factory(&_CrossBridge.TransactOpts, _crossMintableERC20Factory)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_CrossBridge *CrossBridgeTransactor) SetFeeStation(opts *bind.TransactOpts, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setFeeStation", _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_CrossBridge *CrossBridgeSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetFeeStation(&_CrossBridge.TransactOpts, _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetFeeStation(&_CrossBridge.TransactOpts, _bridgeFeeStation)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_CrossBridge *CrossBridgeTransactor) SetNexus(opts *bind.TransactOpts, nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setNexus", nexus_)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_CrossBridge *CrossBridgeSession) SetNexus(nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetNexus(&_CrossBridge.TransactOpts, nexus_)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetNexus(nexus_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetNexus(&_CrossBridge.TransactOpts, nexus_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPauseChain(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPauseChain", remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseChain(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseChain(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPauseToken", remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseToken(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPauseToken(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetRole(opts *bind.TransactOpts, role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setRole", role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetRole(&_CrossBridge.TransactOpts, role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetRole(&_CrossBridge.TransactOpts, role, accounts, set)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeTransactor) SetSafetyLimit(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setSafetyLimit", remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetSafetyLimit(&_CrossBridge.TransactOpts, remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetSafetyLimit(&_CrossBridge.TransactOpts, remoteChainID, token, safetyLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.TransferOwnership(&_CrossBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossBridge *CrossBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.TransferOwnership(&_CrossBridge.TransactOpts, newOwner)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// CrossBridgeBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the CrossBridge contract.
type CrossBridgeBridgeFeeChargedIterator struct {
	Event *CrossBridgeBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFeeCharged)
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
		it.Event = new(CrossBridgeBridgeFeeCharged)
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
func (it *CrossBridgeBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFeeCharged represents a BridgeFeeCharged event raised by the CrossBridge contract.
type CrossBridgeBridgeFeeCharged struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*CrossBridgeBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFeeChargedIterator{contract: _CrossBridge.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFeeCharged)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFeeCharged(log types.Log) (*CrossBridgeBridgeFeeCharged, error) {
	event := new(CrossBridgeBridgeFeeCharged)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeFinalizePendingIterator is returned from FilterBridgeFinalizePending and is used to iterate over the raw logs and unpacked data for BridgeFinalizePending events raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizePendingIterator struct {
	Event *CrossBridgeBridgeFinalizePending // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeFinalizePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFinalizePending)
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
		it.Event = new(CrossBridgeBridgeFinalizePending)
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
func (it *CrossBridgeBridgeFinalizePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFinalizePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFinalizePending represents a BridgeFinalizePending event raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizePending struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizePending is a free log retrieval operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFinalizePending(opts *bind.FilterOpts, index []*big.Int) (*CrossBridgeBridgeFinalizePendingIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFinalizePendingIterator{contract: _CrossBridge.contract, event: "BridgeFinalizePending", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizePending is a free log subscription operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFinalizePending(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFinalizePending, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFinalizePending)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFinalizePending(log types.Log) (*CrossBridgeBridgeFinalizePending, error) {
	event := new(CrossBridgeBridgeFinalizePending)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizedIterator struct {
	Event *CrossBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFinalized)
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
		it.Event = new(CrossBridgeBridgeFinalized)
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
func (it *CrossBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFinalized represents a BridgeFinalized event raised by the CrossBridge contract.
type CrossBridgeBridgeFinalized struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (*CrossBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFinalizedIterator{contract: _CrossBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFinalized, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFinalized)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFinalized(log types.Log) (*CrossBridgeBridgeFinalized, error) {
	event := new(CrossBridgeBridgeFinalized)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the CrossBridge contract.
type CrossBridgeBridgeInitiatedIterator struct {
	Event *CrossBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeInitiated)
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
		it.Event = new(CrossBridgeBridgeInitiated)
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
func (it *CrossBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeInitiated represents a BridgeInitiated event raised by the CrossBridge contract.
type CrossBridgeBridgeInitiated struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (*CrossBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeInitiatedIterator{contract: _CrossBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address indexed localToken, address remoteToken, address from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeInitiated, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeInitiated)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeInitiated(log types.Log) (*CrossBridgeBridgeInitiated, error) {
	event := new(CrossBridgeBridgeInitiated)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the CrossBridge contract.
type CrossBridgeChainPauseSetIterator struct {
	Event *CrossBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeChainPauseSet)
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
		it.Event = new(CrossBridgeChainPauseSet)
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
func (it *CrossBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeChainPauseSet represents a ChainPauseSet event raised by the CrossBridge contract.
type CrossBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*CrossBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeChainPauseSetIterator{contract: _CrossBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeChainPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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

// ParseChainPauseSet is a log parse operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) ParseChainPauseSet(log types.Log) (*CrossBridgeChainPauseSet, error) {
	event := new(CrossBridgeChainPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossMintableERC20FactorySetIterator is returned from FilterCrossMintableERC20FactorySet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20FactorySet events raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20FactorySetIterator struct {
	Event *CrossBridgeCrossMintableERC20FactorySet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeCrossMintableERC20FactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossMintableERC20FactorySet)
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
		it.Event = new(CrossBridgeCrossMintableERC20FactorySet)
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
func (it *CrossBridgeCrossMintableERC20FactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossMintableERC20FactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossMintableERC20FactorySet represents a CrossMintableERC20FactorySet event raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20FactorySet struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20FactorySet is a free log retrieval operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossMintableERC20FactorySet(opts *bind.FilterOpts, factory []common.Address) (*CrossBridgeCrossMintableERC20FactorySetIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossMintableERC20FactorySetIterator{contract: _CrossBridge.contract, event: "CrossMintableERC20FactorySet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20FactorySet is a free log subscription operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossMintableERC20FactorySet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossMintableERC20FactorySet, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossMintableERC20FactorySet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseCrossMintableERC20FactorySet(log types.Log) (*CrossBridgeCrossMintableERC20FactorySet, error) {
	event := new(CrossBridgeCrossMintableERC20FactorySet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChangedIterator struct {
	Event *CrossBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeEIP712DomainChanged)
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
		it.Event = new(CrossBridgeEIP712DomainChanged)
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
func (it *CrossBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeEIP712DomainChangedIterator{contract: _CrossBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeEIP712DomainChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*CrossBridgeEIP712DomainChanged, error) {
	event := new(CrossBridgeEIP712DomainChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeFeeStationSetIterator is returned from FilterFeeStationSet and is used to iterate over the raw logs and unpacked data for FeeStationSet events raised by the CrossBridge contract.
type CrossBridgeFeeStationSetIterator struct {
	Event *CrossBridgeFeeStationSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeFeeStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeFeeStationSet)
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
		it.Event = new(CrossBridgeFeeStationSet)
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
func (it *CrossBridgeFeeStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeFeeStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeFeeStationSet represents a FeeStationSet event raised by the CrossBridge contract.
type CrossBridgeFeeStationSet struct {
	FeeStation common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeStationSet is a free log retrieval operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_CrossBridge *CrossBridgeFilterer) FilterFeeStationSet(opts *bind.FilterOpts, feeStation []common.Address) (*CrossBridgeFeeStationSetIterator, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeFeeStationSetIterator{contract: _CrossBridge.contract, event: "FeeStationSet", logs: logs, sub: sub}, nil
}

// WatchFeeStationSet is a free log subscription operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_CrossBridge *CrossBridgeFilterer) WatchFeeStationSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeFeeStationSet, feeStation []common.Address) (event.Subscription, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeFeeStationSet)
				if err := _CrossBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseFeeStationSet(log types.Log) (*CrossBridgeFeeStationSet, error) {
	event := new(CrossBridgeFeeStationSet)
	if err := _CrossBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossBridge contract.
type CrossBridgeInitializedIterator struct {
	Event *CrossBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeInitialized)
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
		it.Event = new(CrossBridgeInitialized)
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
func (it *CrossBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeInitialized represents a Initialized event raised by the CrossBridge contract.
type CrossBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossBridgeInitializedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeInitializedIterator{contract: _CrossBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeInitialized)
				if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseInitialized(log types.Log) (*CrossBridgeInitialized, error) {
	event := new(CrossBridgeInitialized)
	if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeNexusSetIterator is returned from FilterNexusSet and is used to iterate over the raw logs and unpacked data for NexusSet events raised by the CrossBridge contract.
type CrossBridgeNexusSetIterator struct {
	Event *CrossBridgeNexusSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeNexusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeNexusSet)
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
		it.Event = new(CrossBridgeNexusSet)
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
func (it *CrossBridgeNexusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeNexusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeNexusSet represents a NexusSet event raised by the CrossBridge contract.
type CrossBridgeNexusSet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNexusSet is a free log retrieval operation binding the contract event 0x489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f.
//
// Solidity: event NexusSet(address indexed wallet)
func (_CrossBridge *CrossBridgeFilterer) FilterNexusSet(opts *bind.FilterOpts, wallet []common.Address) (*CrossBridgeNexusSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "NexusSet", walletRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeNexusSetIterator{contract: _CrossBridge.contract, event: "NexusSet", logs: logs, sub: sub}, nil
}

// WatchNexusSet is a free log subscription operation binding the contract event 0x489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f.
//
// Solidity: event NexusSet(address indexed wallet)
func (_CrossBridge *CrossBridgeFilterer) WatchNexusSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeNexusSet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "NexusSet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeNexusSet)
				if err := _CrossBridge.contract.UnpackLog(event, "NexusSet", log); err != nil {
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

// ParseNexusSet is a log parse operation binding the contract event 0x489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f.
//
// Solidity: event NexusSet(address indexed wallet)
func (_CrossBridge *CrossBridgeFilterer) ParseNexusSet(log types.Log) (*CrossBridgeNexusSet, error) {
	event := new(CrossBridgeNexusSet)
	if err := _CrossBridge.contract.UnpackLog(event, "NexusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossBridge contract.
type CrossBridgeOwnershipTransferredIterator struct {
	Event *CrossBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeOwnershipTransferred)
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
		it.Event = new(CrossBridgeOwnershipTransferred)
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
func (it *CrossBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the CrossBridge contract.
type CrossBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossBridge *CrossBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeOwnershipTransferredIterator{contract: _CrossBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossBridge *CrossBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeOwnershipTransferred)
				if err := _CrossBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*CrossBridgeOwnershipTransferred, error) {
	event := new(CrossBridgeOwnershipTransferred)
	if err := _CrossBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossBridge contract.
type CrossBridgePausedIterator struct {
	Event *CrossBridgePaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePaused)
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
		it.Event = new(CrossBridgePaused)
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
func (it *CrossBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePaused represents a Paused event raised by the CrossBridge contract.
type CrossBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossBridgePausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgePausedIterator{contract: _CrossBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossBridgePaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePaused(log types.Log) (*CrossBridgePaused, error) {
	event := new(CrossBridgePaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePendingLockedIterator is returned from FilterPendingLocked and is used to iterate over the raw logs and unpacked data for PendingLocked events raised by the CrossBridge contract.
type CrossBridgePendingLockedIterator struct {
	Event *CrossBridgePendingLocked // Event containing the contract specifics and raw log

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
func (it *CrossBridgePendingLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePendingLocked)
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
		it.Event = new(CrossBridgePendingLocked)
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
func (it *CrossBridgePendingLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePendingLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePendingLocked represents a PendingLocked event raised by the CrossBridge contract.
type CrossBridgePendingLocked struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingLocked is a free log retrieval operation binding the contract event 0xa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f6.
//
// Solidity: event PendingLocked(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterPendingLocked(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgePendingLockedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "PendingLocked", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgePendingLockedIterator{contract: _CrossBridge.contract, event: "PendingLocked", logs: logs, sub: sub}, nil
}

// WatchPendingLocked is a free log subscription operation binding the contract event 0xa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f6.
//
// Solidity: event PendingLocked(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchPendingLocked(opts *bind.WatchOpts, sink chan<- *CrossBridgePendingLocked, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "PendingLocked", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePendingLocked)
				if err := _CrossBridge.contract.UnpackLog(event, "PendingLocked", log); err != nil {
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

// ParsePendingLocked is a log parse operation binding the contract event 0xa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f6.
//
// Solidity: event PendingLocked(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) ParsePendingLocked(log types.Log) (*CrossBridgePendingLocked, error) {
	event := new(CrossBridgePendingLocked)
	if err := _CrossBridge.contract.UnpackLog(event, "PendingLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the CrossBridge contract.
type CrossBridgeRoleUpdatedIterator struct {
	Event *CrossBridgeRoleUpdated // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleUpdated)
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
		it.Event = new(CrossBridgeRoleUpdated)
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
func (it *CrossBridgeRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleUpdated represents a RoleUpdated event raised by the CrossBridge contract.
type CrossBridgeRoleUpdated struct {
	Account common.Address
	Role    [32]byte
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleUpdated(opts *bind.FilterOpts, account []common.Address, role [][32]byte, status []bool) (*CrossBridgeRoleUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleUpdatedIterator{contract: _CrossBridge.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleUpdated, account []common.Address, role [][32]byte, status []bool) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleUpdated)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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

// ParseRoleUpdated is a log parse operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_CrossBridge *CrossBridgeFilterer) ParseRoleUpdated(log types.Log) (*CrossBridgeRoleUpdated, error) {
	event := new(CrossBridgeRoleUpdated)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeSafetyLimitSetIterator is returned from FilterSafetyLimitSet and is used to iterate over the raw logs and unpacked data for SafetyLimitSet events raised by the CrossBridge contract.
type CrossBridgeSafetyLimitSetIterator struct {
	Event *CrossBridgeSafetyLimitSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeSafetyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeSafetyLimitSet)
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
		it.Event = new(CrossBridgeSafetyLimitSet)
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
func (it *CrossBridgeSafetyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeSafetyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeSafetyLimitSet represents a SafetyLimitSet event raised by the CrossBridge contract.
type CrossBridgeSafetyLimitSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	SafetyLimit   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSafetyLimitSet is a free log retrieval operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 safetyLimit)
func (_CrossBridge *CrossBridgeFilterer) FilterSafetyLimitSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeSafetyLimitSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeSafetyLimitSetIterator{contract: _CrossBridge.contract, event: "SafetyLimitSet", logs: logs, sub: sub}, nil
}

// WatchSafetyLimitSet is a free log subscription operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 safetyLimit)
func (_CrossBridge *CrossBridgeFilterer) WatchSafetyLimitSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeSafetyLimitSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeSafetyLimitSet)
				if err := _CrossBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
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

// ParseSafetyLimitSet is a log parse operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 safetyLimit)
func (_CrossBridge *CrossBridgeFilterer) ParseSafetyLimitSet(log types.Log) (*CrossBridgeSafetyLimitSet, error) {
	event := new(CrossBridgeSafetyLimitSet)
	if err := _CrossBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossBridge contract.
type CrossBridgeThresholdChangedIterator struct {
	Event *CrossBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeThresholdChanged)
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
		it.Event = new(CrossBridgeThresholdChanged)
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
func (it *CrossBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeThresholdChanged represents a ThresholdChanged event raised by the CrossBridge contract.
type CrossBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossBridgeThresholdChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeThresholdChangedIterator{contract: _CrossBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeThresholdChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseThresholdChanged(log types.Log) (*CrossBridgeThresholdChanged, error) {
	event := new(CrossBridgeThresholdChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairRegisteredIterator struct {
	Event *CrossBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairRegistered)
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
		it.Event = new(CrossBridgeTokenPairRegistered)
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
func (it *CrossBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	ExchangeRate  *big.Int
	SafetyLimit   *big.Int
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 exchangeRate, uint256 safetyLimit, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*CrossBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairRegisteredIterator{contract: _CrossBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 exchangeRate, uint256 safetyLimit, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairRegistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 exchangeRate, uint256 safetyLimit, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*CrossBridgeTokenPairRegistered, error) {
	event := new(CrossBridgeTokenPairRegistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregisteredIterator struct {
	Event *CrossBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairUnregistered)
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
		it.Event = new(CrossBridgeTokenPairUnregistered)
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
func (it *CrossBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*CrossBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairUnregisteredIterator{contract: _CrossBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairUnregistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*CrossBridgeTokenPairUnregistered, error) {
	event := new(CrossBridgeTokenPairUnregistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the CrossBridge contract.
type CrossBridgeTokenPauseSetIterator struct {
	Event *CrossBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPauseSet)
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
		it.Event = new(CrossBridgeTokenPauseSet)
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
func (it *CrossBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPauseSet represents a TokenPauseSet event raised by the CrossBridge contract.
type CrossBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPauseSetIterator{contract: _CrossBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPauseSet(log types.Log) (*CrossBridgeTokenPauseSet, error) {
	event := new(CrossBridgeTokenPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossBridge contract.
type CrossBridgeUnpausedIterator struct {
	Event *CrossBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUnpaused)
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
		it.Event = new(CrossBridgeUnpaused)
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
func (it *CrossBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUnpaused represents a Unpaused event raised by the CrossBridge contract.
type CrossBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossBridgeUnpausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUnpausedIterator{contract: _CrossBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUnpaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUnpaused(log types.Log) (*CrossBridgeUnpaused, error) {
	event := new(CrossBridgeUnpaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossBridge contract.
type CrossBridgeUpgradedIterator struct {
	Event *CrossBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUpgraded)
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
		it.Event = new(CrossBridgeUpgraded)
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
func (it *CrossBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUpgraded represents a Upgraded event raised by the CrossBridge contract.
type CrossBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUpgradedIterator{contract: _CrossBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUpgraded)
				if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUpgraded(log types.Log) (*CrossBridgeUpgraded, error) {
	event := new(CrossBridgeUpgraded)
	if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
