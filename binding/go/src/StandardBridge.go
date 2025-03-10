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

// IBridgeRegistryFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryFinalizeArguments struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Token         common.Address
	To            common.Address
	Value         *big.Int
	ExtraData     []byte
}

// IBridgeRegistryPendingData is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryPendingData struct {
	Args         IBridgeRegistryFinalizeArguments
	SafeDeadline *big.Int
	Reason       []byte
}

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryTokenPair struct {
	LocalToken    common.Address
	RemoteToken   common.Address
	Paused        bool
	IsOrigin      bool
	SafetyLimit   *big.Int
	Deposited     *big.Int
	PendingAmount *big.Int
}

// IStandardBridgeBridgeTokenArguments is an auto generated low-level Go binding around an user-defined struct.
type IStandardBridgeBridgeTokenArguments struct {
	RemoteChainID *big.Int
	Token         common.Address
	To            common.Address
	Value         *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
	ExtraData     []byte
}

// IStandardBridgePermitArguments is an auto generated low-level Go binding around an user-defined struct.
type IStandardBridgePermitArguments struct {
	Token    common.Address
	Account  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// StandardBridgeMetaData contains all meta data concerning the StandardBridge contract.
var StandardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasExpiredPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"nexus_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lockPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualProcessPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nexus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIStandardBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"processExpiredPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"resetRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Factory\",\"name\":\"_crossMintableERC20Factory\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"nexus_\",\"type\":\"address\"}],\"name\":\"setNexus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"setSafetyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20FactorySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"NexusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"}],\"name\":\"SafetyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safetyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"RegistryExistFactory\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleAlreadyGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardFailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"StandardInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"StandardInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"StandardNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"StandardNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StandardNotExpiredSafeDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StandardNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"}]",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615fe66100f95f395f81816138630152818161388c0152613a530152615fe65ff3fe6080604052600436106102c2575f3560e01c80637f4ab9f51161016f578063ae766389116100d8578063d4bf502a11610092578063f17f6cb71161006d578063f17f6cb714610973578063f2fde38b14610992578063f4509643146109b1578063f698da25146109d0575f5ffd5b8063d4bf502a14610922578063d5717fc514610941578063d605665b14610960575f5ffd5b8063ae7663891461084b578063b33eb36e14610885578063b7f3358d146108b1578063bedb86fb146108d0578063cf56118e146108ef578063d016d62514610903575f5ffd5b806391cf6d3e1161012957806391cf6d3e1461077357806391d1485414610787578063a3246ad3146107a6578063a3f5c1d2146107d2578063ad3cb1cc146107ef578063ae6893f81461082c575f5ffd5b80637f4ab9f514610601578063814914b51461062057806384b0196e1461070657806388d67d6d1461072d5780638da5cb5b146107405780638f517c1714610754575f5ffd5b806347666cb11161022b57806354db0126116101e55780635fd262de116101c05780635fd262de1461058f5780636160751f146105a2578063715018a6146105c157806379214874146105d5575f5ffd5b806354db0126146105215780635b605f5c146105405780635c975abb1461056c575f5ffd5b806347666cb1146104645780634d3f0da91461049b5780634d5d0056146104ba5780634f1ef286146104cd5780635187599d146104e057806352d1902d146104ff575f5ffd5b80632670b8171161027c5780632670b817146103b05780632d87b7ee146103cf5780633960e787146103ee57806339a621f31461040d5780633d507c5e1461042c57806342cde4e814610443575f5ffd5b8063030372c3146102ed5780630b43c02c146103215780631089fd58146103405780631938e0f21461035f5780631a1aebbb146103725780631e7bf21514610391575f5ffd5b366102e957345f036102e7576040516365d14ce560e11b815260040160405180910390fd5b005b5f5ffd5b3480156102f8575f5ffd5b5061030c610307366004614d7e565b6109e4565b60405190151581526020015b60405180910390f35b34801561032c575f5ffd5b506102e761033b366004614e1f565b610a28565b34801561034b575f5ffd5b506102e761035a366004614e1f565b610a84565b61030c61036d366004614f12565b610d6c565b34801561037d575f5ffd5b506102e761038c366004614fdf565b611139565b34801561039c575f5ffd5b506102e76103ab366004615007565b6111e9565b3480156103bb575f5ffd5b506102e76103ca366004614fdf565b611499565b3480156103da575f5ffd5b506102e76103e93660046150cc565b611511565b3480156103f9575f5ffd5b5061030c61040836600461510f565b611543565b348015610418575f5ffd5b506102e761042736600461512f565b611848565b348015610437575f5ffd5b5060335442111561030c565b34801561044e575f5ffd5b5060645460405160ff9091168152602001610318565b34801561046f575f5ffd5b50609654610483906001600160a01b031681565b6040516001600160a01b039091168152602001610318565b3480156104a6575f5ffd5b506102e76104b5366004615164565b61193b565b61030c6104c83660046151e7565b611a79565b6102e76104db3660046152ff565b611c6c565b3480156104eb575f5ffd5b506102e76104fa366004615341565b611c87565b34801561050a575f5ffd5b50610513611d95565b604051908152602001610318565b34801561052c575f5ffd5b506102e761053b366004614fdf565b611db0565b34801561054b575f5ffd5b5061055f61055a366004614e1f565b611e48565b60405161031891906153c9565b348015610577575f5ffd5b505f516020615f715f395f51905f525460ff1661030c565b61030c61059d366004615416565b611fcd565b3480156105ad575f5ffd5b506102e76105bc36600461549e565b612036565b3480156105cc575f5ffd5b506102e761212b565b3480156105e0575f5ffd5b506105f46105ef366004614e1f565b61213e565b60405161031891906154fb565b34801561060c575f5ffd5b5061030c61061b36600461510f565b612157565b34801561062b575f5ffd5b506106f961063a36600461550d565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526039602090815260408084206001600160a01b039384168552825292839020835160e08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b6040516103189190615530565b348015610711575f5ffd5b5061071a6121f2565b604051610318979695949392919061556c565b61030c61073b366004615699565b61229b565b34801561074b575f5ffd5b50610483612336565b34801561075f575f5ffd5b50603254610483906001600160a01b031681565b34801561077e575f5ffd5b50609854610513565b348015610792575f5ffd5b5061030c6107a136600461550d565b612364565b3480156107b1575f5ffd5b506107c56107c0366004614e1f565b61237b565b60405161031891906157c9565b3480156107dd575f5ffd5b506097546001600160a01b0316610483565b3480156107fa575f5ffd5b5061081f604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516103189190615809565b348015610837575f5ffd5b50610513610846366004614e1f565b612394565b348015610856575f5ffd5b5061086a61086536600461512f565b6123b0565b60408051938452602084019290925290820152606001610318565b348015610890575f5ffd5b506108a461089f36600461510f565b61243f565b604051610318919061581b565b3480156108bc575f5ffd5b506102e76108cb3660046158ad565b6125e9565b3480156108db575f5ffd5b506102e76108ea3660046158c6565b61263a565b3480156108fa575f5ffd5b506105f461265b565b34801561090e575f5ffd5b5061048361091d3660046158e1565b61266c565b34801561092d575f5ffd5b506102e761093c36600461596a565b61277c565b34801561094c575f5ffd5b5061051361095b366004614e1f565b6127b3565b6102e761096e3660046159b5565b6127cf565b34801561097e575f5ffd5b506102e761098d36600461510f565b612962565b34801561099d575f5ffd5b506102e76109ac366004614fdf565b612a39565b3480156109bc575f5ffd5b506102e76109cb36600461550d565b612a73565b3480156109db575f5ffd5b50610513612b58565b5f805b8251811015610a1c57610a1384848381518110610a0657610a06615a4e565b6020026020010151611543565b506001016109e7565b50600190505b92915050565b610a30612b61565b5f818152603b60205260408120610a4690612b93565b90505f5b8151811015610a7f57610a7683838381518110610a6957610a69615a4e565b6020026020010151612b9f565b50600101610a4a565b505050565b426033541015610a915750565b5f80610a9d6035612b93565b90505f5b8151811015610d66575f828281518110610abd57610abd615a4e565b6020908102919091018101515f818152603790925260409091206003015490915060ff1615610aec5750610d5e565b5f818152603b60205260408120610b0290612b93565b90505f5b8151811015610d5a575f828281518110610b2257610b22615a4e565b6020908102919091018101515f868152603c83526040808220838352909352828120835161012081019094528054606085019081526001820154608086015260028201546001600160a01b0390811660a087015260038301541660c0860152600482015460e0860152600582018054949650929493919284928491610100850191610bac90615a62565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd890615a62565b8015610c235780601f10610bfa57610100808354040283529160200191610c23565b820191905f5260205f20905b815481529060010190602001808311610c0657829003601f168201915b505050505081525050815260200160068201548152602001600782018054610c4a90615a62565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7690615a62565b8015610cc15780601f10610c9857610100808354040283529160200191610cc1565b820191905f5260205f20905b815481529060010190602001808311610ca457829003601f168201915b5050509190925250508151515f90815260396020908152604080832085518201516001600160a01b0316845290915290206001015491925050600160a01b900460ff16158015610d145750602081015115155b8015610d235750428160200151105b15610d5057610d328583612dac565b5088610d3d89615aae565b98508810610d5057505050505050505050565b5050600101610b06565b5050505b600101610aa1565b50505050565b5f610d75612e7b565b610d7d612eab565b610da5610d906060870160408801614fdf565b86355f90815260386020526040902090612ee2565b610db56060870160408801614fdf565b90610de4576040516353b2527760e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610e0f576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b505084355f9081526037602052604081206002018054600101908190559050806020870135808214610e5d57604051631351db4160e31b815260048101929092526024820152604401610ddb565b50610eef90507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020880135610e9960608a0160408b01614fdf565b610ea960808b0160608c01614fdf565b60808b0135610ebb60a08d018d615ac6565b604051602001610ed19796959493929190615b30565b60405160208183030381529060405280519060200120868686612f03565b85355f90815260396020526040808220606091839182918290610f17908d8701908e01614fdf565b6001600160a01b03908116825260208083019390935260409182015f20825160e08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c0830152909150610fc8576040518060400160405280600c81526020016b1d1bdad95b881c185d5cd95960a21b8152509250611051565b608081015115801590610fe2575089608001358160800151105b1561101c57604051806040016040528060118152602001701bdd995c881cd859995d1e481b1a5b5a5d607a1b815250925060019150611051565b61104b8a3561103160608d0160408e01614fdf565b61104160808e0160608f01614fdf565b8d60800135613153565b90945092505b5082156110d95761106860808a0160608b01614fdf565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b6110a960608e0160408f01614fdf565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4611113565b6110e489838361324a565b60405160208a0135907f40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2905f90a25b600194505050505061113160015f516020615f915f395f51905f5255565b949350505050565b611141612b61565b6032546001600160a01b0316801561117857604051639ad61dbd60e01b81526001600160a01b039091166004820152602401610ddb565b506001600160a01b0381166111a057604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517f18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9905f90a250565b6111f1612b61565b6111fc60358761340a565b1561125957604080516080810182528781525f6020808301828152838501838152606085018481528c855260379093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b03841661128057604051636ca1fdd760e01b815260040160405180910390fd5b5f8681526038602052604090206112979085613415565b84906112c2576040516311dd05f360e31b81526001600160a01b039091166004820152602401610ddb565b506040518060e00160405280856001600160a01b03168152602001846001600160a01b031681526020015f1515815260200186151581526020018281526020015f81526020015f81525060395f8881526020019081526020015f205f866001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c08201518160040155905050600182138061140457505f1982125b1561142e575f868152603a602090815260408083206001600160a01b038816845290915290208290555b826001600160a01b0316846001600160a01b0316877fffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a6685858a6040516114899392919092835260208301919091521515604082015260600190565b60405180910390a4505050505050565b6114a1612b61565b6001600160a01b0381166114c857604051630fb9363360e41b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0383169081179091556040517f489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f905f90a250565b6115338261152d5f5f8681526020019081526020015f20612b93565b5f61277c565b61153f8282600161277c565b5050565b5f61154c612e7b565b611554612eab565b5f838152603b6020526040902061156b9083613429565b829061158d57604051630e0b575b60e21b8152600401610ddb91815260200190565b505f838152603c60209081526040808320858452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061160990615a62565b80601f016020809104026020016040519081016040528092919081815260200182805461163590615a62565b80156116805780601f1061165757610100808354040283529160200191611680565b820191905f5260205f20905b81548152906001019060200180831161166357829003601f168201915b5050505050815250508152602001600682015481526020016007820180546116a790615a62565b80601f01602080910402602001604051908101604052809291908181526020018280546116d390615a62565b801561171e5780601f106116f55761010080835404028352916020019161171e565b820191905f5260205f20905b81548152906001019060200180831161170157829003601f168201915b505050919092525050505f85815260396020908152604080832084518201516001600160a01b03908116855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b8504811615801585850152600160a81b909504161515606084015260028101546080840152600381015460a08401526004015460c08301528351015192935091906117df576040516338384f6f60e11b81526001600160a01b039091166004820152602401610ddb565b50602082015115806117f45750428260200151105b826020015142909161182257604051636d60854360e01b815260048101929092526024820152604401610ddb565b505061182e8585612dac565b92505050610a2260015f516020615f915f395f51905f5255565b61185a6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916118955760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f8381526038602052604090206118ae9083612ee2565b82906118d95760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8381526039602090815260408083206001600160a01b038616808552908352928190206002018490555183815285917fd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c891015b60405180910390a3505050565b61194d6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916119885760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f8381526038602052604090206119a19083612ee2565b82906119cc5760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8381526039602090815260408083206001600160a01b038616845290915290206001015460ff600160a01b90910416151581151514610a7f575f8381526039602090815260408083206001600160a01b0386168085529252918290206001018054841515600160a01b0260ff60a01b19909116179055905184907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061192e90851515815260200190565b5f611a82612e7b565b8989611a8e8282613440565b611a96612eab565b611aa36020850185614fdf565b6001600160a01b038c81169116148b611abf6020870187614fdf565b9091611af15760405163a9212f4360e01b81526001600160a01b03928316600482015291166024820152604401610ddb565b5050611b008c8c8b8b8b613516565b909850965086611b10898b615b7d565b611b1a9190615b7d565b6040850135101587611b2c8a8c615b7d565b611b369190615b7d565b85604001359091611b63576040516365efbabf60e11b815260048101929092526024820152604401610ddb565b505f9050611b776040860160208701614fdf565b3060408701356060880135611b9260a08a0160808b016158ad565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a086013560c482015260c086013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031663d505accf60e01b179052909150611c2090611c1990870187614fdf565b5f836135fa565b50611c438c8c611c366040880160208901614fdf565b8d8d8d8d60018e8e613723565b60019250611c5d60015f516020615f915f395f51905f5255565b50509998505050505050505050565b611c74613858565b611c7d826138fc565b61153f8282613904565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611ccb5750825b90505f826001600160401b03166001148015611ce65750303b155b905081158015611cf4575080155b15611d125760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611d3c57845460ff60401b1916600160401b1785555b611d4687876139c0565b8315611d8c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f611d9e613a48565b505f516020615f515f395f51905f5290565b611dc26420a226a4a760d91b33612364565b6420a226a4a760d91b339091611dfd5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b5050609680546001600160a01b0319166001600160a01b0383169081179091556040517f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444905f90a250565b5f81815260386020526040812060609190611e6290612b93565b90505f81516001600160401b03811115611e7e57611e7e614cf0565b604051908082528060200260200182016040528015611ee357816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f19909201910181611e9c5790505b5090505f5b8251811015611fc55760395f8681526020019081526020015f205f848381518110611f1557611f15615a4e565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c08201528251839083908110611fb257611fb2615a4e565b6020908102919091010152600101611ee8565b509392505050565b5f611fd6612e7b565b8888611fe28282613440565b611fea612eab565b611ff78b8b8a8a8a613516565b909750955061200e8b8b338c8c8c8c5f8d8d613723565b6001925061202860015f516020615f915f395f51905f5255565b505098975050505050505050565b6120486420a226a4a760d91b33612364565b6420a226a4a760d91b3390916120835760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b506120919050603583613429565b82906120b35760405163ac7dbbfd60e01b8152600401610ddb91815260200190565b505f8281526037602052604090206003015460ff1615158115151461153f575f82815260376020908152604091829020600301805460ff1916841515908117909155915191825283917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a25050565b612133612b61565b61213c5f613a91565b565b5f818152603b60205260409020606090610a2290612b93565b5f61216a6420a226a4a760d91b33612364565b6420a226a4a760d91b3390916121a55760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f838152603b602052604090206121be9083613429565b82906121e057604051630e0b575b60e21b8152600401610ddb91815260200190565b506121eb8383612dac565b9392505050565b5f60608082808083815f516020615f315f395f51905f52805490915015801561221d57506001810154155b6122615760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610ddb565b612269613b01565b612271613bc1565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015612329576123208787838181106122bb576122bb615a4e565b90506020028101906122cd9190615b90565b8683815181106122df576122df615a4e565b60200260200101518684815181106122f9576122f9615a4e565b602002602001015186858151811061231357612313615a4e565b6020026020010151610d6c565b5060010161229e565b5060019695505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f8281526020819052604081206121eb9083612ee2565b5f818152602081905260409020606090610a2290612b93565b5f818152603760205260408120600190810154610a2291615b7d565b60965460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa15801561240c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124309190615bae565b91989097509095509350505050565b612447614c3d565b5f838152603c6020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e08401526005820180548492918491610100850191906124c390615a62565b80601f01602080910402602001604051908101604052809291908181526020018280546124ef90615a62565b801561253a5780601f106125115761010080835404028352916020019161253a565b820191905f5260205f20905b81548152906001019060200180831161251d57829003601f168201915b50505050508152505081526020016006820154815260200160078201805461256190615a62565b80601f016020809104026020016040519081016040528092919081815260200182805461258d90615a62565b80156125d85780601f106125af576101008083540402835291602001916125d8565b820191905f5260205f20905b8154815290600101906020018083116125bb57829003601f168201915b505050505081525050905092915050565b6125f1612b61565b6064805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b612642612b61565b801561265357612650613bff565b50565b612650613c5b565b60606126676035612b93565b905090565b6032545f906001600160a01b0316612697576040516315aeca0d60e11b815260040160405180910390fd5b6032546040516bffffffffffffffffffffffff19606089901b1660208201526001600160a01b0390911690634804a0419060340160405160208183030381529060405280519060200120856040516020016126f29190615bf0565b60405160208183030381529060405286866040518563ffffffff1660e01b81526004016127229493929190615c11565b6020604051808303815f875af115801561273e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127629190615c50565b9050612772875f838989896111e9565b9695505050505050565b5f5b8251811015610d66576127ab8484838151811061279d5761279d615a4e565b602002602001015184613ca0565b60010161277e565b5f81815260376020526040812060020154610a22906001615b7d565b8281146127ef5760405163214485c960e01b815260040160405180910390fd5b5f5b8381101561295b5761295285858381811061280e5761280e615a4e565b90506020028101906128209190615c6b565b3586868481811061283357612833615a4e565b90506020028101906128459190615c6b565b612856906040810190602001614fdf565b87878581811061286857612868615a4e565b905060200281019061287a9190615c6b565b61288b906060810190604001614fdf565b88888681811061289d5761289d615a4e565b90506020028101906128af9190615c6b565b606001358989878181106128c5576128c5615a4e565b90506020028101906128d79190615c6b565b608001358a8a888181106128ed576128ed615a4e565b90506020028101906128ff9190615c6b565b60a001358b8b8981811061291557612915615a4e565b90506020028101906129279190615c6b565b6129359060c0810190615ac6565b8b8b8b81811061294757612947615a4e565b905060e00201611a79565b506001016127f1565b5050505050565b6129746420a226a4a760d91b33612364565b6420a226a4a760d91b3390916129af5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b50505f828152603b602052604090206129c89082613429565b81906129ea57604051630e0b575b60e21b8152600401610ddb91815260200190565b505f828152603c602090815260408083208484529091528082205f1960069091015551829184917fa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f69190a35050565b612a41612b61565b6001600160a01b038116612a6a57604051631e4fbdf760e01b81525f6004820152602401610ddb565b61265081613a91565b612a7b612b61565b5f828152603860205260409020612a929082613dd0565b8190612abd5760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f8281526039602090815260408083206001600160a01b03851680855290835281842080546001600160a01b03191681556001810180546001600160b01b03191690556002810185905560038101859055600401849055858452603a835281842081855290925280832083905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f612667613de4565b33612b6a612336565b6001600160a01b03161461213c5760405163118cdaa760e01b8152336004820152602401610ddb565b60605f6121eb83613ded565b612ba7614c63565b5f838152603b60205260409020612bbe9083613e46565b8290612be057604051637f11bea960e01b8152600401610ddb91815260200190565b505f838152603c60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190612c5890615a62565b80601f0160208091040260200160405190810160405280929190818152602001828054612c8490615a62565b8015612ccf5780601f10612ca657610100808354040283529160200191612ccf565b820191905f5260205f20905b815481529060010190602001808311612cb257829003601f168201915b505050919092525050505f848152603960209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615612d32578160800151816004015f828254612d2c9190615c7f565b90915550505b5f848152603c602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181612d8b6005830182614ca6565b5050600682015f9055600782015f612da39190614ca6565b50505092915050565b5f5f612db88484612b9f565b90505f5f612dd7835f0151846040015185606001518660800151613153565b91509150818190612dfb5760405162461bcd60e51b8152600401610ddb9190615809565b5082606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8660400151876080015142604051612e67939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b5f516020615f715f395f51905f525460ff161561213c5760405163d93c066560e01b815260040160405180910390fd5b5f516020615f915f395f51905f52805460011901612edc57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f90815260018301602052604081205415156121eb565b8251825181148015612f155750815181145b835183518392612f49576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610ddb565b505060645482915060ff16811015612f7757604051631aebd17960e11b8152600401610ddb91815260200190565b505f80826001600160401b03811115612f9257612f92614cf0565b604051908082528060200260200182016040528015612fbb578160200160208202803683370190505b5090505f5b8381101561311d575f61302b888381518110612fde57612fde615a4e565b6020026020010151888481518110612ff857612ff8615a4e565b602002602001015188858151811061301257613012615a4e565b60200260200101516130238d613e51565b929190613e7d565b9050613043682b20a624a220aa27a960b91b82612364565b682b20a624a220aa27a960b91b8290916130825760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610ddb565b505f9050805b858110156130d3578481815181106130a2576130a2615a4e565b60200260200101516001600160a01b0316836001600160a01b0316036130cb57600191506130d3565b600101613088565b508061311357818486815181106130ec576130ec615a4e565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b5050600101612fc0565b50606454829060ff1681101561314957604051631aebd17960e11b8152600401610ddb91815260200190565b5050505050505050565b5f848152603a602090815260408083206001600160a01b038716845290915281205460609060018113156131925761318b8185615c92565b93506131b1565b5f198112156131b1576131a481615ca9565b6131ae9085615cd7565b93505b5f196001600160a01b038716016131e4576131db858560405180602001604052805f8152506135fa565b6001925061323f565b831561323f575f8781526039602090815260408083206001600160a01b038a168452909152902060010154600160a81b900460ff16156132345761322a87878787613ea9565b9250925050613241565b61322a868686613f9f565b505b94509492505050565b82355f908152603b602090815260409091206132689185013561340a565b83602001359061328e576040516307a5c91d60e31b8152600401610ddb91815260200190565b505f81156132b7576034546132a39042615b7d565b9050603454426132b39190615b7d565b6033555b6040518060600160405280856132cc90615cea565b81526020808201849052604091820186905286355f908152603c82528281208883013582528252829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a08101518290600582019061335e9082615db4565b505050602082015160068201556040820151600782019061337f9082615db4565b50505083355f908152603960205260408082209082906133a59060608901908901614fdf565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff161561295b578460800135816004015f8282546133eb9190615b7d565b90915550505050505050565b60015f516020615f915f395f51905f5255565b5f6121eb8383614081565b5f6121eb836001600160a01b038416614081565b5f81815260018301602052604081205415156121eb565b5f8281526038602052604090206134579082612ee2565b81906134825760405163153096f360e11b81526001600160a01b039091166004820152602401610ddb565b505f82815260376020526040902060030154829060ff16156134ba57604051636fc24b7960e11b8152600401610ddb91815260200190565b505f8281526039602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610a7f576040516338384f6f60e11b81526001600160a01b039091166004820152602401610ddb565b6096545f9081906001600160a01b031661353457505f9050806135f0565b60965460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa15801561358a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135ae9190615bae565b909450925090508086108015906135c55750828510155b80156135d15750818410155b6135ee5760405163dc48c7f160e01b815260040160405180910390fd5b505b9550959350505050565b814747821115613626576040516371e4d07760e11b815260048101929092526024820152604401610ddb565b50505f5f846001600160a01b031684846040516136439190615e6e565b5f6040518083038185875af1925050503d805f811461367d576040519150601f19603f3d011682016040523d82523d5f602084013e613682565b606091505b5091509150816136b55780511561369c5780518082602001fd5b604051630daf67c960e11b815260040160405180910390fd5b835f0361295b5780515f036136f357846001600160a01b03163b5f036136ee57604051630daf67c960e11b815260040160405180910390fd5b61295b565b602081015160018115151461371b57604051630daf67c960e11b815260040160405180910390fd5b505050505050565b6137398a8a8a89613734898b615b7d565b6140cd565b5f8a815260376020526040812060019081018054909101908190558190915060395f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b031690508a6001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7848e8e8e428d8d8d6040516137eb989796959493929190615e79565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a604051613842929190918252602082015260400190565b60405180910390a4505050505050505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806138de57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166138d25f516020615f515f395f51905f52546001600160a01b031690565b6001600160a01b031614155b1561213c5760405163703e46dd60e11b815260040160405180910390fd5b612650612b61565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561395e575060408051601f3d908101601f1916820190925261395b91810190615ec3565b60015b61398657604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610ddb565b5f516020615f515f395f51905f5281146139b657604051632a87526960e21b815260048101829052602401610ddb565b610a7f838361431e565b6139c8614373565b6139d1336143bc565b6139d96143cd565b6139e16143d5565b6139e96143e5565b6139f2826143f5565b6139fa61445b565b6001600160a01b038116613a2157604051630fb9363360e41b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b03929092169190911790555043609855565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461213c5760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615f315f395f51905f5291613b3f90615a62565b80601f0160208091040260200160405190810160405280929190818152602001828054613b6b90615a62565b8015613bb65780601f10613b8d57610100808354040283529160200191613bb6565b820191905f5260205f20905b815481529060010190602001808311613b9957829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615f315f395f51905f5291613b3f90615a62565b613c07612e7b565b5f516020615f715f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200161262f565b613c63614487565b5f516020615f715f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613c43565b613ca8612b61565b6001600160a01b038216613ccf57604051635989b9d360e01b815260040160405180910390fd5b82613ced57604051630e1b39f960e31b815260040160405180910390fd5b8015613d44575f838152602081905260409020613d0a9083613415565b83839091613d3d57604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610ddb565b5050613d91565b5f838152602081905260409020613d5b9083613dd0565b83839091613d8e57604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610ddb565b50505b80151583836001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f360405160405180910390a4505050565b5f6121eb836001600160a01b0384166144b6565b5f612667614599565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613e3a57602002820191905f5260205f20905b815481526020019060010190808311613e26575b50505050509050919050565b5f6121eb83836144b6565b5f610a22613e5d613de4565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613e8d8888888861460c565b925092509250613e9d82826146d4565b50909695505050505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613f18575060408051601f3d908101601f19168201909252613f1591810190615eda565b60015b613f52573d808015613f45576040519150601f19603f3d011682016040523d82523d5f602084013e613f4a565b606091505b509050613241565b8092508015613f6b57613f6687878661478c565b61323f565b6040518060400160405280600f81526020016e1d1c985b9cd9995c8819985a5b1959608a1b81525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af192505050801561400e575060408051601f3d908101601f1916820190925261400b91810190615eda565b60015b614048573d80801561403b576040519150601f19603f3d011682016040523d82523d5f602084013e614040565b606091505b509050614079565b80925080614077576040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b81525091505b505b935093915050565b5f8181526001830160205260408120546140c657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a22565b505f610a22565b5f858152603a602090815260408083206001600160a01b03881684529091529020546001811315614138576141028184615ef5565b8590849015614135576040516348d53e0760e01b81526001600160a01b0390921660048301526024820152604401610ddb565b50505b5f196001600160a01b038616016141bb576141538284615b7d565b341461415f8385615b7d565b349091614188576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b505081156141b65760975460408051602081019091525f81526141b6916001600160a01b03169084906135fa565b61371b565b5f3480156141e5576040516329e2b03f60e21b815260048101929092526024820152604401610ddb565b50614209905084306141f78587615b7d565b6001600160a01b038916929190614837565b811561422957609754614229906001600160a01b0387811691168461489e565b5f8681526039602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff1615614268576141b68686856148cf565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af11580156142b2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142d69190615eda565b8585859091926143135760405163601bc95b60e11b81526001600160a01b0393841660048201529290911660248301526044820152606401610ddb565b505050505050505050565b6143278261490d565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561436b57610a7f8282614970565b61153f6149e2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661213c57604051631afcd79f60e31b815260040160405180910390fd5b6143c4614373565b61265081614a01565b61213c614373565b6143dd614373565b61213c614a09565b6143ed614373565b61213c614a29565b6143fd614373565b614445604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250614a31565b6064805460ff191660ff92909216919091179055565b614463614373565b61447e6420a226a4a760d91b614477612336565b6001613ca0565b62015180603455565b5f516020615f715f395f51905f525460ff1661213c57604051638dfc202b60e01b815260040160405180910390fd5b5f8181526001830160205260408120548015614590575f6144d8600183615c7f565b85549091505f906144eb90600190615c7f565b905080821461454a575f865f01828154811061450957614509615a4e565b905f5260205f200154905080875f01848154811061452957614529615a4e565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061455b5761455b615f08565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a22565b5f915050610a22565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6145c3614a43565b6145cb614aab565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561464557505f915060039050826146ca565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614696573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166146c157505f9250600191508290506146ca565b92505f91508190505b9450945094915050565b5f8260038111156146e7576146e7615f1c565b036146f0575050565b600182600381111561470457614704615f1c565b036147225760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561473657614736615f1c565b036147575760405163fce698f760e01b815260048101829052602401610ddb565b600382600381111561476b5761476b615f1c565b0361153f576040516335e2f38360e21b815260048101829052602401610ddb565b5f8381526039602090815260408083206001600160a01b0386168452909152902060048101546147bc9083615b7d565b600382015460048301548692869286929091821015614814576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610ddb565b505050505081816003015f82825461482c9190615c7f565b909155505050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610d669186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614aed565b6040516001600160a01b03838116602483015260448201839052610a7f91859182169063a9059cbb9060640161486c565b5f8381526039602090815260408083206001600160a01b038616845290915281206003018054839290614903908490615b7d565b9091555050505050565b806001600160a01b03163b5f0361494257604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610ddb565b5f516020615f515f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161498c9190615e6e565b5f60405180830381855af49150503d805f81146149c4576040519150601f19603f3d011682016040523d82523d5f602084013e6149c9565b606091505b50915091506149d9858383614b59565b95945050505050565b341561213c5760405163b398979f60e01b815260040160405180910390fd5b612a41614373565b614a11614373565b5f516020615f715f395f51905f52805460ff19169055565b6133f7614373565b614a39614373565b61153f8282614bb5565b5f5f516020615f315f395f51905f5281614a5b613b01565b805190915015614a7357805160209091012092915050565b81548015614a82579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615f315f395f51905f5281614ac3613bc1565b805190915015614adb57805160209091012092915050565b60018201548015614a82579392505050565b5f5f60205f8451602086015f885af180614b0c576040513d5f823e3d81fd5b50505f513d91508115614b23578060011415614b30565b6001600160a01b0384163b155b15610d6657604051635274afe760e01b81526001600160a01b0385166004820152602401610ddb565b606082614b6e57614b6982614c14565b6121eb565b8151158015614b8557506001600160a01b0384163b155b15614bae57604051639996b31560e01b81526001600160a01b0385166004820152602401610ddb565b5092915050565b614bbd614373565b5f516020615f315f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614bf68482615db4565b5060038101614c058382615db4565b505f8082556001909101555050565b805115614c245780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614c50614c63565b81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614cb290615a62565b5f825580601f10614cc1575050565b601f0160209004905f5260205f209081019061265091905b80821115614cec575f8155600101614cd9565b5090565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614d2657614d26614cf0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614d5457614d54614cf0565b604052919050565b5f6001600160401b03821115614d7457614d74614cf0565b5060051b60200190565b5f5f60408385031215614d8f575f5ffd5b8235915060208301356001600160401b03811115614dab575f5ffd5b8301601f81018513614dbb575f5ffd5b8035614dce614dc982614d5c565b614d2c565b8082825260208201915060208360051b850101925087831115614def575f5ffd5b6020840193505b82841015614e11578335825260209384019390910190614df6565b809450505050509250929050565b5f60208284031215614e2f575f5ffd5b5035919050565b803560ff81168114614e46575f5ffd5b919050565b5f82601f830112614e5a575f5ffd5b8135614e68614dc982614d5c565b8082825260208201915060208360051b860101925085831115614e89575f5ffd5b602085015b83811015614ead57614e9f81614e36565b835260209283019201614e8e565b5095945050505050565b5f82601f830112614ec6575f5ffd5b8135614ed4614dc982614d5c565b8082825260208201915060208360051b860101925085831115614ef5575f5ffd5b602085015b83811015614ead578035835260209283019201614efa565b5f5f5f5f60808587031215614f25575f5ffd5b84356001600160401b03811115614f3a575f5ffd5b850160c08188031215614f4b575f5ffd5b935060208501356001600160401b03811115614f65575f5ffd5b614f7187828801614e4b565b93505060408501356001600160401b03811115614f8c575f5ffd5b614f9887828801614eb7565b92505060608501356001600160401b03811115614fb3575f5ffd5b614fbf87828801614eb7565b91505092959194509250565b6001600160a01b0381168114612650575f5ffd5b5f60208284031215614fef575f5ffd5b81356121eb81614fcb565b8015158114612650575f5ffd5b5f5f5f5f5f5f60c0878903121561501c575f5ffd5b86359550602087013561502e81614ffa565b9450604087013561503e81614fcb565b9350606087013561504e81614fcb565b9598949750929560808101359460a0909101359350915050565b5f82601f830112615077575f5ffd5b8135615085614dc982614d5c565b8082825260208201915060208360051b8601019250858311156150a6575f5ffd5b602085015b83811015614ead5780356150be81614fcb565b8352602092830192016150ab565b5f5f604083850312156150dd575f5ffd5b8235915060208301356001600160401b038111156150f9575f5ffd5b61510585828601615068565b9150509250929050565b5f5f60408385031215615120575f5ffd5b50508035926020909101359150565b5f5f5f60608486031215615141575f5ffd5b83359250602084013561515381614fcb565b929592945050506040919091013590565b5f5f5f60608486031215615176575f5ffd5b83359250602084013561518881614fcb565b9150604084013561519881614ffa565b809150509250925092565b5f5f83601f8401126151b3575f5ffd5b5081356001600160401b038111156151c9575f5ffd5b6020830191508360208285010111156151e0575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c0811215615201575f5ffd5b8a35995060208b013561521381614fcb565b985060408b013561522381614fcb565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115615252575f5ffd5b61525e8d828e016151a3565b90955093505060e060df1982011215615275575f5ffd5b5060e08a0190509295985092959850929598565b5f5f6001600160401b038411156152a2576152a2614cf0565b50601f8301601f19166020016152b781614d2c565b9150508281528383830111156152cb575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126152f0575f5ffd5b6121eb83833560208501615289565b5f5f60408385031215615310575f5ffd5b823561531b81614fcb565b915060208301356001600160401b03811115615335575f5ffd5b615105858286016152e1565b5f5f60408385031215615352575f5ffd5b61535b83614e36565b9150602083013561536b81614fcb565b809150509250929050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b8181101561540b576153f5838551615376565b6020939093019260e092909201916001016153e2565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b03121561542d575f5ffd5b88359750602089013561543f81614fcb565b9650604089013561544f81614fcb565b9550606089013594506080890135935060a0890135925060c08901356001600160401b0381111561547e575f5ffd5b61548a8b828c016151a3565b999c989b5096995094979396929594505050565b5f5f604083850312156154af575f5ffd5b82359150602083013561536b81614ffa565b5f8151808452602084019350602083015f5b828110156154f15781518652602095860195909101906001016154d3565b5093949350505050565b602081525f6121eb60208301846154c1565b5f5f6040838503121561551e575f5ffd5b82359150602083013561536b81614fcb565b60e08101610a228284615376565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61558a60e083018961553e565b828103604084015261559c818961553e565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506155cd81856154c1565b9a9950505050505050505050565b5f5f83601f8401126155eb575f5ffd5b5081356001600160401b03811115615601575f5ffd5b6020830191508360208260051b85010111156151e0575f5ffd5b5f82601f83011261562a575f5ffd5b8135615638614dc982614d5c565b8082825260208201915060208360051b860101925085831115615659575f5ffd5b602085015b83811015614ead5780356001600160401b0381111561567b575f5ffd5b61568a886020838a0101614eb7565b8452506020928301920161565e565b5f5f5f5f5f608086880312156156ad575f5ffd5b85356001600160401b038111156156c2575f5ffd5b6156ce888289016155db565b90965094505060208601356001600160401b038111156156ec575f5ffd5b8601601f810188136156fc575f5ffd5b803561570a614dc982614d5c565b8082825260208201915060208360051b85010192508a83111561572b575f5ffd5b602084015b8381101561576b5780356001600160401b0381111561574d575f5ffd5b61575c8d602083890101614e4b565b84525060209283019201615730565b50955050505060408601356001600160401b03811115615789575f5ffd5b6157958882890161561b565b92505060608601356001600160401b038111156157b0575f5ffd5b6157bc8882890161561b565b9150509295509295909350565b602080825282518282018190525f918401906040840190835b8181101561540b5783516001600160a01b03168352602093840193909201916001016157e2565b602081525f6121eb602083018461553e565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c061012084015261588661014084018261553e565b9050602084015160408401526040840151601f198483030160608501526149d9828261553e565b5f602082840312156158bd575f5ffd5b6121eb82614e36565b5f602082840312156158d6575f5ffd5b81356121eb81614ffa565b5f5f5f5f5f5f60c087890312156158f6575f5ffd5b86359550602087013561590881614fcb565b9450604087013593506060870135925060808701356001600160401b03811115615930575f5ffd5b8701601f81018913615940575f5ffd5b61594f89823560208401615289565b92505061595e60a08801614e36565b90509295509295509295565b5f5f5f6060848603121561597c575f5ffd5b8335925060208401356001600160401b03811115615998575f5ffd5b6159a486828701615068565b925050604084013561519881614ffa565b5f5f5f5f604085870312156159c8575f5ffd5b84356001600160401b038111156159dd575f5ffd5b6159e9878288016155db565b90955093505060208501356001600160401b03811115615a07575f5ffd5b8501601f81018713615a17575f5ffd5b80356001600160401b03811115615a2c575f5ffd5b87602060e083028401011115615a40575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680615a7657607f821691505b602082108103615a9457634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201615abf57615abf615a9a565b5060010190565b5f5f8335601e19843603018112615adb575f5ffd5b8301803591506001600160401b03821115615af4575f5ffd5b6020019150368190038213156151e0575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615b709083018486615b08565b9998505050505050505050565b80820180821115610a2257610a22615a9a565b5f823560be19833603018112615ba4575f5ffd5b9190910192915050565b5f5f5f60608486031215615bc0575f5ffd5b5050815160208301516040909301519094929350919050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f6121eb600d830184615bd9565b848152608060208201525f615c29608083018661553e565b8281036040840152615c3b818661553e565b91505060ff8316606083015295945050505050565b5f60208284031215615c60575f5ffd5b81516121eb81614fcb565b5f823560de19833603018112615ba4575f5ffd5b81810381811115610a2257610a22615a9a565b8082028115828204841417610a2257610a22615a9a565b5f600160ff1b8201615cbd57615cbd615a9a565b505f0390565b634e487b7160e01b5f52601260045260245ffd5b5f82615ce557615ce5615cc3565b500490565b5f60c08236031215615cfa575f5ffd5b615d02614d04565b82358152602080840135908201526040830135615d1e81614fcb565b60408201526060830135615d3181614fcb565b60608201526080838101359082015260a08301356001600160401b03811115615d58575f5ffd5b615d64368286016152e1565b60a08301525092915050565b601f821115610a7f57805f5260205f20601f840160051c81016020851015615d955750805b601f840160051c820191505b8181101561295b575f8155600101615da1565b81516001600160401b03811115615dcd57615dcd614cf0565b615de181615ddb8454615a62565b84615d70565b6020601f821160018114615e13575f8315615dfc5750848201515b5f19600385901b1c1916600184901b17845561295b565b5f84815260208120601f198516915b82811015615e425787850151825560209485019460019092019101615e22565b5084821015615e5f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6121eb8284615bd9565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f906155cd9083018486615b08565b5f60208284031215615ed3575f5ffd5b5051919050565b5f60208284031215615eea575f5ffd5b81516121eb81614ffa565b5f82615f0357615f03615cc3565b500690565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220f6b427f6947e4f3d516a73b2cdf7e58fc8997a0ba67772ff82e9d275846b52df64736f6c634300081c0033",
}

// StandardBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardBridgeMetaData.ABI instead.
var StandardBridgeABI = StandardBridgeMetaData.ABI

// Deprecated: Use StandardBridgeMetaData.Sigs instead.
// StandardBridgeFuncSigs maps the 4-byte function signature to its string representation.
var StandardBridgeFuncSigs = StandardBridgeMetaData.Sigs

// StandardBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StandardBridgeMetaData.Bin instead.
var StandardBridgeBin = StandardBridgeMetaData.Bin

// DeployStandardBridge deploys a new Ethereum contract, binding an instance of StandardBridge to it.
func DeployStandardBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardBridge, error) {
	parsed, err := StandardBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StandardBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardBridge{StandardBridgeCaller: StandardBridgeCaller{contract: contract}, StandardBridgeTransactor: StandardBridgeTransactor{contract: contract}, StandardBridgeFilterer: StandardBridgeFilterer{contract: contract}}, nil
}

// StandardBridge is an auto generated Go binding around an Ethereum contract.
type StandardBridge struct {
	StandardBridgeCaller     // Read-only binding to the contract
	StandardBridgeTransactor // Write-only binding to the contract
	StandardBridgeFilterer   // Log filterer for contract events
}

// StandardBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardBridgeSession struct {
	Contract     *StandardBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardBridgeCallerSession struct {
	Contract *StandardBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StandardBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardBridgeTransactorSession struct {
	Contract     *StandardBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StandardBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardBridgeRaw struct {
	Contract *StandardBridge // Generic contract binding to access the raw methods on
}

// StandardBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardBridgeCallerRaw struct {
	Contract *StandardBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// StandardBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardBridgeTransactorRaw struct {
	Contract *StandardBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardBridge creates a new instance of StandardBridge, bound to a specific deployed contract.
func NewStandardBridge(address common.Address, backend bind.ContractBackend) (*StandardBridge, error) {
	contract, err := bindStandardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardBridge{StandardBridgeCaller: StandardBridgeCaller{contract: contract}, StandardBridgeTransactor: StandardBridgeTransactor{contract: contract}, StandardBridgeFilterer: StandardBridgeFilterer{contract: contract}}, nil
}

// NewStandardBridgeCaller creates a new read-only instance of StandardBridge, bound to a specific deployed contract.
func NewStandardBridgeCaller(address common.Address, caller bind.ContractCaller) (*StandardBridgeCaller, error) {
	contract, err := bindStandardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeCaller{contract: contract}, nil
}

// NewStandardBridgeTransactor creates a new write-only instance of StandardBridge, bound to a specific deployed contract.
func NewStandardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardBridgeTransactor, error) {
	contract, err := bindStandardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeTransactor{contract: contract}, nil
}

// NewStandardBridgeFilterer creates a new log filterer instance of StandardBridge, bound to a specific deployed contract.
func NewStandardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardBridgeFilterer, error) {
	contract, err := bindStandardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeFilterer{contract: contract}, nil
}

// bindStandardBridge binds a generic wrapper to an already deployed contract.
func bindStandardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardBridge *StandardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardBridge.Contract.StandardBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardBridge *StandardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardBridge.Contract.StandardBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardBridge *StandardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardBridge.Contract.StandardBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardBridge *StandardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardBridge *StandardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardBridge *StandardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StandardBridge *StandardBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StandardBridge *StandardBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StandardBridge.Contract.UPGRADEINTERFACEVERSION(&_StandardBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StandardBridge *StandardBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StandardBridge.Contract.UPGRADEINTERFACEVERSION(&_StandardBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_StandardBridge *StandardBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_StandardBridge *StandardBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _StandardBridge.Contract.AllChainIDs(&_StandardBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_StandardBridge *StandardBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _StandardBridge.Contract.AllChainIDs(&_StandardBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_StandardBridge *StandardBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_StandardBridge *StandardBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _StandardBridge.Contract.AllPendingIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_StandardBridge *StandardBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _StandardBridge.Contract.AllPendingIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_StandardBridge *StandardBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_StandardBridge *StandardBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _StandardBridge.Contract.AllTokenPairs(&_StandardBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_StandardBridge *StandardBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _StandardBridge.Contract.AllTokenPairs(&_StandardBridge.CallOpts, remoteChainID)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_StandardBridge *StandardBridgeCaller) BridgeFeeStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "bridgeFeeStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_StandardBridge *StandardBridgeSession) BridgeFeeStation() (common.Address, error) {
	return _StandardBridge.Contract.BridgeFeeStation(&_StandardBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_StandardBridge *StandardBridgeCallerSession) BridgeFeeStation() (common.Address, error) {
	return _StandardBridge.Contract.BridgeFeeStation(&_StandardBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_StandardBridge *StandardBridgeCaller) CrossMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "crossMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_StandardBridge *StandardBridgeSession) CrossMintableERC20Factory() (common.Address, error) {
	return _StandardBridge.Contract.CrossMintableERC20Factory(&_StandardBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_StandardBridge *StandardBridgeCallerSession) CrossMintableERC20Factory() (common.Address, error) {
	return _StandardBridge.Contract.CrossMintableERC20Factory(&_StandardBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StandardBridge *StandardBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StandardBridge *StandardBridgeSession) DomainSeparator() ([32]byte, error) {
	return _StandardBridge.Contract.DomainSeparator(&_StandardBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_StandardBridge *StandardBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _StandardBridge.Contract.DomainSeparator(&_StandardBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_StandardBridge *StandardBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_StandardBridge *StandardBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _StandardBridge.Contract.Eip712Domain(&_StandardBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_StandardBridge *StandardBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _StandardBridge.Contract.Eip712Domain(&_StandardBridge.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_StandardBridge *StandardBridgeCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

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
func (_StandardBridge *StandardBridgeSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _StandardBridge.Contract.EstimateFee(&_StandardBridge.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_StandardBridge *StandardBridgeCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _StandardBridge.Contract.EstimateFee(&_StandardBridge.CallOpts, remoteChainID, token, value)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _StandardBridge.Contract.GetNextFinalizeIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _StandardBridge.Contract.GetNextFinalizeIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _StandardBridge.Contract.GetNextInitiateIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_StandardBridge *StandardBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _StandardBridge.Contract.GetNextInitiateIndex(&_StandardBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_StandardBridge *StandardBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_StandardBridge *StandardBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _StandardBridge.Contract.GetPendingArguments(&_StandardBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_StandardBridge *StandardBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _StandardBridge.Contract.GetPendingArguments(&_StandardBridge.CallOpts, remoteChainID, index)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_StandardBridge *StandardBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_StandardBridge *StandardBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _StandardBridge.Contract.GetRoleMembers(&_StandardBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_StandardBridge *StandardBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _StandardBridge.Contract.GetRoleMembers(&_StandardBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_StandardBridge *StandardBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_StandardBridge *StandardBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _StandardBridge.Contract.GetTokenPair(&_StandardBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_StandardBridge *StandardBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _StandardBridge.Contract.GetTokenPair(&_StandardBridge.CallOpts, remoteChainID, token)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_StandardBridge *StandardBridgeCaller) HasExpiredPending(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "hasExpiredPending")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_StandardBridge *StandardBridgeSession) HasExpiredPending() (bool, error) {
	return _StandardBridge.Contract.HasExpiredPending(&_StandardBridge.CallOpts)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_StandardBridge *StandardBridgeCallerSession) HasExpiredPending() (bool, error) {
	return _StandardBridge.Contract.HasExpiredPending(&_StandardBridge.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StandardBridge *StandardBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StandardBridge *StandardBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StandardBridge.Contract.HasRole(&_StandardBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StandardBridge *StandardBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StandardBridge.Contract.HasRole(&_StandardBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_StandardBridge *StandardBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_StandardBridge *StandardBridgeSession) InitializedAt() (*big.Int, error) {
	return _StandardBridge.Contract.InitializedAt(&_StandardBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_StandardBridge *StandardBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _StandardBridge.Contract.InitializedAt(&_StandardBridge.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_StandardBridge *StandardBridgeCaller) Nexus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "nexus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_StandardBridge *StandardBridgeSession) Nexus() (common.Address, error) {
	return _StandardBridge.Contract.Nexus(&_StandardBridge.CallOpts)
}

// Nexus is a free data retrieval call binding the contract method 0xa3f5c1d2.
//
// Solidity: function nexus() view returns(address)
func (_StandardBridge *StandardBridgeCallerSession) Nexus() (common.Address, error) {
	return _StandardBridge.Contract.Nexus(&_StandardBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardBridge *StandardBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardBridge *StandardBridgeSession) Owner() (common.Address, error) {
	return _StandardBridge.Contract.Owner(&_StandardBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardBridge *StandardBridgeCallerSession) Owner() (common.Address, error) {
	return _StandardBridge.Contract.Owner(&_StandardBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StandardBridge *StandardBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StandardBridge *StandardBridgeSession) Paused() (bool, error) {
	return _StandardBridge.Contract.Paused(&_StandardBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StandardBridge *StandardBridgeCallerSession) Paused() (bool, error) {
	return _StandardBridge.Contract.Paused(&_StandardBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StandardBridge *StandardBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StandardBridge *StandardBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _StandardBridge.Contract.ProxiableUUID(&_StandardBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StandardBridge *StandardBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StandardBridge.Contract.ProxiableUUID(&_StandardBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StandardBridge *StandardBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StandardBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StandardBridge *StandardBridgeSession) Threshold() (uint8, error) {
	return _StandardBridge.Contract.Threshold(&_StandardBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_StandardBridge *StandardBridgeCallerSession) Threshold() (uint8, error) {
	return _StandardBridge.Contract.Threshold(&_StandardBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactor) BridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "bridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_StandardBridge *StandardBridgeSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.BridgeToken(&_StandardBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.BridgeToken(&_StandardBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_StandardBridge *StandardBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_StandardBridge *StandardBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _StandardBridge.Contract.ChangeThreshold(&_StandardBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_StandardBridge *StandardBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _StandardBridge.Contract.ChangeThreshold(&_StandardBridge.TransactOpts, threshold_)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_StandardBridge *StandardBridgeTransactor) ClearPending(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "clearPending", remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_StandardBridge *StandardBridgeSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ClearPending(&_StandardBridge.TransactOpts, remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_StandardBridge *StandardBridgeTransactorSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ClearPending(&_StandardBridge.TransactOpts, remoteChainID)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_StandardBridge *StandardBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_StandardBridge *StandardBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _StandardBridge.Contract.CreateToken(&_StandardBridge.TransactOpts, remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 exchangeRate, uint256 safetyLimit, string symbol, uint8 decimals) returns(address tokenAddress)
func (_StandardBridge *StandardBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _StandardBridge.Contract.CreateToken(&_StandardBridge.TransactOpts, remoteChainID, remoteToken, exchangeRate, safetyLimit, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_StandardBridge *StandardBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.FinalizeBridge(&_StandardBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.FinalizeBridge(&_StandardBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_StandardBridge *StandardBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.FinalizeBridgeBatch(&_StandardBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.FinalizeBridgeBatch(&_StandardBridge.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_StandardBridge *StandardBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "initialize", _threshold, nexus_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_StandardBridge *StandardBridgeSession) Initialize(_threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.Initialize(&_StandardBridge.TransactOpts, _threshold, nexus_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address nexus_) returns()
func (_StandardBridge *StandardBridgeTransactorSession) Initialize(_threshold uint8, nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.Initialize(&_StandardBridge.TransactOpts, _threshold, nexus_)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_StandardBridge *StandardBridgeTransactor) LockPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "lockPending", remoteChainID, index)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_StandardBridge *StandardBridgeSession) LockPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.LockPending(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// LockPending is a paid mutator transaction binding the contract method 0xf17f6cb7.
//
// Solidity: function lockPending(uint256 remoteChainID, uint256 index) returns()
func (_StandardBridge *StandardBridgeTransactorSession) LockPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.LockPending(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeTransactor) ManualProcessPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "manualProcessPending", remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ManualProcessPending(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ManualProcessPending(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_StandardBridge *StandardBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.Contract.PermitBridgeToken(&_StandardBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.Contract.PermitBridgeToken(&_StandardBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_StandardBridge *StandardBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_StandardBridge *StandardBridgeSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.Contract.PermitBridgeTokenBatch(&_StandardBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_StandardBridge *StandardBridgeTransactorSession) PermitBridgeTokenBatch(args []IStandardBridgeBridgeTokenArguments, permitArgs []IStandardBridgePermitArguments) (*types.Transaction, error) {
	return _StandardBridge.Contract.PermitBridgeTokenBatch(&_StandardBridge.TransactOpts, args, permitArgs)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_StandardBridge *StandardBridgeTransactor) ProcessExpiredPending(opts *bind.TransactOpts, maxCount *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "processExpiredPending", maxCount)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_StandardBridge *StandardBridgeSession) ProcessExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ProcessExpiredPending(&_StandardBridge.TransactOpts, maxCount)
}

// ProcessExpiredPending is a paid mutator transaction binding the contract method 0x1089fd58.
//
// Solidity: function processExpiredPending(uint256 maxCount) returns()
func (_StandardBridge *StandardBridgeTransactorSession) ProcessExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.ProcessExpiredPending(&_StandardBridge.TransactOpts, maxCount)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RegisterToken(&_StandardBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 exchangeRate, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, exchangeRate *big.Int, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RegisterToken(&_StandardBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, exchangeRate, safetyLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardBridge *StandardBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardBridge *StandardBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _StandardBridge.Contract.RenounceOwnership(&_StandardBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardBridge *StandardBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StandardBridge.Contract.RenounceOwnership(&_StandardBridge.TransactOpts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_StandardBridge *StandardBridgeTransactor) ResetRole(opts *bind.TransactOpts, role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "resetRole", role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_StandardBridge *StandardBridgeSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.ResetRole(&_StandardBridge.TransactOpts, role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_StandardBridge *StandardBridgeTransactorSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.ResetRole(&_StandardBridge.TransactOpts, role, newAccounts)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "retryFinalizeBridge", remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RetryFinalizeBridge(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RetryFinalizeBridge(&_StandardBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_StandardBridge *StandardBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_StandardBridge *StandardBridgeSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RetryFinalizeBridgeBatch(&_StandardBridge.TransactOpts, remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_StandardBridge *StandardBridgeTransactorSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.RetryFinalizeBridgeBatch(&_StandardBridge.TransactOpts, remoteChainID, indexes)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_StandardBridge *StandardBridgeTransactor) SetCrossMintableERC20Factory(opts *bind.TransactOpts, _crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setCrossMintableERC20Factory", _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_StandardBridge *StandardBridgeSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetCrossMintableERC20Factory(&_StandardBridge.TransactOpts, _crossMintableERC20Factory)
}

// SetCrossMintableERC20Factory is a paid mutator transaction binding the contract method 0x1a1aebbb.
//
// Solidity: function setCrossMintableERC20Factory(address _crossMintableERC20Factory) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetCrossMintableERC20Factory(_crossMintableERC20Factory common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetCrossMintableERC20Factory(&_StandardBridge.TransactOpts, _crossMintableERC20Factory)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_StandardBridge *StandardBridgeTransactor) SetFeeStation(opts *bind.TransactOpts, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setFeeStation", _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_StandardBridge *StandardBridgeSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetFeeStation(&_StandardBridge.TransactOpts, _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetFeeStation(&_StandardBridge.TransactOpts, _bridgeFeeStation)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_StandardBridge *StandardBridgeTransactor) SetNexus(opts *bind.TransactOpts, nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setNexus", nexus_)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_StandardBridge *StandardBridgeSession) SetNexus(nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetNexus(&_StandardBridge.TransactOpts, nexus_)
}

// SetNexus is a paid mutator transaction binding the contract method 0x2670b817.
//
// Solidity: function setNexus(address nexus_) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetNexus(nexus_ common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetNexus(&_StandardBridge.TransactOpts, nexus_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_StandardBridge *StandardBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_StandardBridge *StandardBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPause(&_StandardBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPause(&_StandardBridge.TransactOpts, set)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_StandardBridge *StandardBridgeTransactor) SetPauseChain(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setPauseChain", remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_StandardBridge *StandardBridgeSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPauseChain(&_StandardBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPauseChain(&_StandardBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_StandardBridge *StandardBridgeTransactor) SetPauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setPauseToken", remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_StandardBridge *StandardBridgeSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPauseToken(&_StandardBridge.TransactOpts, remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetPauseToken(&_StandardBridge.TransactOpts, remoteChainID, token, pause)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_StandardBridge *StandardBridgeTransactor) SetRole(opts *bind.TransactOpts, role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setRole", role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_StandardBridge *StandardBridgeSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetRole(&_StandardBridge.TransactOpts, role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetRole(&_StandardBridge.TransactOpts, role, accounts, set)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeTransactor) SetSafetyLimit(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "setSafetyLimit", remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetSafetyLimit(&_StandardBridge.TransactOpts, remoteChainID, token, safetyLimit)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 safetyLimit) returns()
func (_StandardBridge *StandardBridgeTransactorSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, safetyLimit *big.Int) (*types.Transaction, error) {
	return _StandardBridge.Contract.SetSafetyLimit(&_StandardBridge.TransactOpts, remoteChainID, token, safetyLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardBridge *StandardBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardBridge *StandardBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.TransferOwnership(&_StandardBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardBridge *StandardBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.TransferOwnership(&_StandardBridge.TransactOpts, newOwner)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_StandardBridge *StandardBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_StandardBridge *StandardBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.UnregisterToken(&_StandardBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_StandardBridge *StandardBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _StandardBridge.Contract.UnregisterToken(&_StandardBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StandardBridge *StandardBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StandardBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StandardBridge *StandardBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.UpgradeToAndCall(&_StandardBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StandardBridge *StandardBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StandardBridge.Contract.UpgradeToAndCall(&_StandardBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StandardBridge *StandardBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StandardBridge *StandardBridgeSession) Receive() (*types.Transaction, error) {
	return _StandardBridge.Contract.Receive(&_StandardBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StandardBridge *StandardBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _StandardBridge.Contract.Receive(&_StandardBridge.TransactOpts)
}

// StandardBridgeBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the StandardBridge contract.
type StandardBridgeBridgeFeeChargedIterator struct {
	Event *StandardBridgeBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *StandardBridgeBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeBridgeFeeCharged)
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
		it.Event = new(StandardBridgeBridgeFeeCharged)
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
func (it *StandardBridgeBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeBridgeFeeCharged represents a BridgeFeeCharged event raised by the StandardBridge contract.
type StandardBridgeBridgeFeeCharged struct {
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
func (_StandardBridge *StandardBridgeFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*StandardBridgeBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeBridgeFeeChargedIterator{contract: _StandardBridge.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_StandardBridge *StandardBridgeFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *StandardBridgeBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeBridgeFeeCharged)
				if err := _StandardBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseBridgeFeeCharged(log types.Log) (*StandardBridgeBridgeFeeCharged, error) {
	event := new(StandardBridgeBridgeFeeCharged)
	if err := _StandardBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeBridgeFinalizePendingIterator is returned from FilterBridgeFinalizePending and is used to iterate over the raw logs and unpacked data for BridgeFinalizePending events raised by the StandardBridge contract.
type StandardBridgeBridgeFinalizePendingIterator struct {
	Event *StandardBridgeBridgeFinalizePending // Event containing the contract specifics and raw log

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
func (it *StandardBridgeBridgeFinalizePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeBridgeFinalizePending)
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
		it.Event = new(StandardBridgeBridgeFinalizePending)
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
func (it *StandardBridgeBridgeFinalizePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeBridgeFinalizePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeBridgeFinalizePending represents a BridgeFinalizePending event raised by the StandardBridge contract.
type StandardBridgeBridgeFinalizePending struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizePending is a free log retrieval operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_StandardBridge *StandardBridgeFilterer) FilterBridgeFinalizePending(opts *bind.FilterOpts, index []*big.Int) (*StandardBridgeBridgeFinalizePendingIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeBridgeFinalizePendingIterator{contract: _StandardBridge.contract, event: "BridgeFinalizePending", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizePending is a free log subscription operation binding the contract event 0x40c1d3562756f3f19a0504cfdef1405cc3fcb9c8cf0660fff1d5f86a37d40fe2.
//
// Solidity: event BridgeFinalizePending(uint256 indexed index)
func (_StandardBridge *StandardBridgeFilterer) WatchBridgeFinalizePending(opts *bind.WatchOpts, sink chan<- *StandardBridgeBridgeFinalizePending, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "BridgeFinalizePending", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeBridgeFinalizePending)
				if err := _StandardBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseBridgeFinalizePending(log types.Log) (*StandardBridgeBridgeFinalizePending, error) {
	event := new(StandardBridgeBridgeFinalizePending)
	if err := _StandardBridge.contract.UnpackLog(event, "BridgeFinalizePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the StandardBridge contract.
type StandardBridgeBridgeFinalizedIterator struct {
	Event *StandardBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *StandardBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeBridgeFinalized)
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
		it.Event = new(StandardBridgeBridgeFinalized)
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
func (it *StandardBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeBridgeFinalized represents a BridgeFinalized event raised by the StandardBridge contract.
type StandardBridgeBridgeFinalized struct {
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
func (_StandardBridge *StandardBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (*StandardBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeBridgeFinalizedIterator{contract: _StandardBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_StandardBridge *StandardBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *StandardBridgeBridgeFinalized, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeBridgeFinalized)
				if err := _StandardBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseBridgeFinalized(log types.Log) (*StandardBridgeBridgeFinalized, error) {
	event := new(StandardBridgeBridgeFinalized)
	if err := _StandardBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the StandardBridge contract.
type StandardBridgeBridgeInitiatedIterator struct {
	Event *StandardBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *StandardBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeBridgeInitiated)
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
		it.Event = new(StandardBridgeBridgeInitiated)
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
func (it *StandardBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeBridgeInitiated represents a BridgeInitiated event raised by the StandardBridge contract.
type StandardBridgeBridgeInitiated struct {
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
func (_StandardBridge *StandardBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (*StandardBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeBridgeInitiatedIterator{contract: _StandardBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address indexed localToken, address remoteToken, address from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_StandardBridge *StandardBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *StandardBridgeBridgeInitiated, remoteChainID []*big.Int, index []*big.Int, localToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeBridgeInitiated)
				if err := _StandardBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseBridgeInitiated(log types.Log) (*StandardBridgeBridgeInitiated, error) {
	event := new(StandardBridgeBridgeInitiated)
	if err := _StandardBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the StandardBridge contract.
type StandardBridgeChainPauseSetIterator struct {
	Event *StandardBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeChainPauseSet)
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
		it.Event = new(StandardBridgeChainPauseSet)
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
func (it *StandardBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeChainPauseSet represents a ChainPauseSet event raised by the StandardBridge contract.
type StandardBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_StandardBridge *StandardBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*StandardBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeChainPauseSetIterator{contract: _StandardBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_StandardBridge *StandardBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *StandardBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeChainPauseSet)
				if err := _StandardBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseChainPauseSet(log types.Log) (*StandardBridgeChainPauseSet, error) {
	event := new(StandardBridgeChainPauseSet)
	if err := _StandardBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeCrossMintableERC20FactorySetIterator is returned from FilterCrossMintableERC20FactorySet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20FactorySet events raised by the StandardBridge contract.
type StandardBridgeCrossMintableERC20FactorySetIterator struct {
	Event *StandardBridgeCrossMintableERC20FactorySet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeCrossMintableERC20FactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeCrossMintableERC20FactorySet)
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
		it.Event = new(StandardBridgeCrossMintableERC20FactorySet)
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
func (it *StandardBridgeCrossMintableERC20FactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeCrossMintableERC20FactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeCrossMintableERC20FactorySet represents a CrossMintableERC20FactorySet event raised by the StandardBridge contract.
type StandardBridgeCrossMintableERC20FactorySet struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20FactorySet is a free log retrieval operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_StandardBridge *StandardBridgeFilterer) FilterCrossMintableERC20FactorySet(opts *bind.FilterOpts, factory []common.Address) (*StandardBridgeCrossMintableERC20FactorySetIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeCrossMintableERC20FactorySetIterator{contract: _StandardBridge.contract, event: "CrossMintableERC20FactorySet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20FactorySet is a free log subscription operation binding the contract event 0x18df1bf4db6b347f79d6775fd433461aadaf0c13ed6ac4e83873534f6705c1b9.
//
// Solidity: event CrossMintableERC20FactorySet(address indexed factory)
func (_StandardBridge *StandardBridgeFilterer) WatchCrossMintableERC20FactorySet(opts *bind.WatchOpts, sink chan<- *StandardBridgeCrossMintableERC20FactorySet, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "CrossMintableERC20FactorySet", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeCrossMintableERC20FactorySet)
				if err := _StandardBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseCrossMintableERC20FactorySet(log types.Log) (*StandardBridgeCrossMintableERC20FactorySet, error) {
	event := new(StandardBridgeCrossMintableERC20FactorySet)
	if err := _StandardBridge.contract.UnpackLog(event, "CrossMintableERC20FactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the StandardBridge contract.
type StandardBridgeEIP712DomainChangedIterator struct {
	Event *StandardBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *StandardBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeEIP712DomainChanged)
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
		it.Event = new(StandardBridgeEIP712DomainChanged)
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
func (it *StandardBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the StandardBridge contract.
type StandardBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_StandardBridge *StandardBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*StandardBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &StandardBridgeEIP712DomainChangedIterator{contract: _StandardBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_StandardBridge *StandardBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *StandardBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeEIP712DomainChanged)
				if err := _StandardBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*StandardBridgeEIP712DomainChanged, error) {
	event := new(StandardBridgeEIP712DomainChanged)
	if err := _StandardBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeFeeStationSetIterator is returned from FilterFeeStationSet and is used to iterate over the raw logs and unpacked data for FeeStationSet events raised by the StandardBridge contract.
type StandardBridgeFeeStationSetIterator struct {
	Event *StandardBridgeFeeStationSet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeFeeStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeFeeStationSet)
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
		it.Event = new(StandardBridgeFeeStationSet)
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
func (it *StandardBridgeFeeStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeFeeStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeFeeStationSet represents a FeeStationSet event raised by the StandardBridge contract.
type StandardBridgeFeeStationSet struct {
	FeeStation common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeStationSet is a free log retrieval operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_StandardBridge *StandardBridgeFilterer) FilterFeeStationSet(opts *bind.FilterOpts, feeStation []common.Address) (*StandardBridgeFeeStationSetIterator, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeFeeStationSetIterator{contract: _StandardBridge.contract, event: "FeeStationSet", logs: logs, sub: sub}, nil
}

// WatchFeeStationSet is a free log subscription operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address indexed feeStation)
func (_StandardBridge *StandardBridgeFilterer) WatchFeeStationSet(opts *bind.WatchOpts, sink chan<- *StandardBridgeFeeStationSet, feeStation []common.Address) (event.Subscription, error) {

	var feeStationRule []interface{}
	for _, feeStationItem := range feeStation {
		feeStationRule = append(feeStationRule, feeStationItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "FeeStationSet", feeStationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeFeeStationSet)
				if err := _StandardBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseFeeStationSet(log types.Log) (*StandardBridgeFeeStationSet, error) {
	event := new(StandardBridgeFeeStationSet)
	if err := _StandardBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StandardBridge contract.
type StandardBridgeInitializedIterator struct {
	Event *StandardBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *StandardBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeInitialized)
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
		it.Event = new(StandardBridgeInitialized)
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
func (it *StandardBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeInitialized represents a Initialized event raised by the StandardBridge contract.
type StandardBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StandardBridge *StandardBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*StandardBridgeInitializedIterator, error) {

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StandardBridgeInitializedIterator{contract: _StandardBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StandardBridge *StandardBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StandardBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeInitialized)
				if err := _StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseInitialized(log types.Log) (*StandardBridgeInitialized, error) {
	event := new(StandardBridgeInitialized)
	if err := _StandardBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeNexusSetIterator is returned from FilterNexusSet and is used to iterate over the raw logs and unpacked data for NexusSet events raised by the StandardBridge contract.
type StandardBridgeNexusSetIterator struct {
	Event *StandardBridgeNexusSet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeNexusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeNexusSet)
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
		it.Event = new(StandardBridgeNexusSet)
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
func (it *StandardBridgeNexusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeNexusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeNexusSet represents a NexusSet event raised by the StandardBridge contract.
type StandardBridgeNexusSet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNexusSet is a free log retrieval operation binding the contract event 0x489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f.
//
// Solidity: event NexusSet(address indexed wallet)
func (_StandardBridge *StandardBridgeFilterer) FilterNexusSet(opts *bind.FilterOpts, wallet []common.Address) (*StandardBridgeNexusSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "NexusSet", walletRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeNexusSetIterator{contract: _StandardBridge.contract, event: "NexusSet", logs: logs, sub: sub}, nil
}

// WatchNexusSet is a free log subscription operation binding the contract event 0x489b51edd4714861d63224d4ab9a10ed75c25102854b0dae7b961bfc094e1f0f.
//
// Solidity: event NexusSet(address indexed wallet)
func (_StandardBridge *StandardBridgeFilterer) WatchNexusSet(opts *bind.WatchOpts, sink chan<- *StandardBridgeNexusSet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "NexusSet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeNexusSet)
				if err := _StandardBridge.contract.UnpackLog(event, "NexusSet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseNexusSet(log types.Log) (*StandardBridgeNexusSet, error) {
	event := new(StandardBridgeNexusSet)
	if err := _StandardBridge.contract.UnpackLog(event, "NexusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StandardBridge contract.
type StandardBridgeOwnershipTransferredIterator struct {
	Event *StandardBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StandardBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeOwnershipTransferred)
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
		it.Event = new(StandardBridgeOwnershipTransferred)
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
func (it *StandardBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the StandardBridge contract.
type StandardBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StandardBridge *StandardBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StandardBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeOwnershipTransferredIterator{contract: _StandardBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StandardBridge *StandardBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StandardBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeOwnershipTransferred)
				if err := _StandardBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*StandardBridgeOwnershipTransferred, error) {
	event := new(StandardBridgeOwnershipTransferred)
	if err := _StandardBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StandardBridge contract.
type StandardBridgePausedIterator struct {
	Event *StandardBridgePaused // Event containing the contract specifics and raw log

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
func (it *StandardBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgePaused)
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
		it.Event = new(StandardBridgePaused)
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
func (it *StandardBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgePaused represents a Paused event raised by the StandardBridge contract.
type StandardBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StandardBridge *StandardBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*StandardBridgePausedIterator, error) {

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StandardBridgePausedIterator{contract: _StandardBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StandardBridge *StandardBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StandardBridgePaused) (event.Subscription, error) {

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgePaused)
				if err := _StandardBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParsePaused(log types.Log) (*StandardBridgePaused, error) {
	event := new(StandardBridgePaused)
	if err := _StandardBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgePendingLockedIterator is returned from FilterPendingLocked and is used to iterate over the raw logs and unpacked data for PendingLocked events raised by the StandardBridge contract.
type StandardBridgePendingLockedIterator struct {
	Event *StandardBridgePendingLocked // Event containing the contract specifics and raw log

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
func (it *StandardBridgePendingLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgePendingLocked)
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
		it.Event = new(StandardBridgePendingLocked)
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
func (it *StandardBridgePendingLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgePendingLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgePendingLocked represents a PendingLocked event raised by the StandardBridge contract.
type StandardBridgePendingLocked struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingLocked is a free log retrieval operation binding the contract event 0xa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f6.
//
// Solidity: event PendingLocked(uint256 indexed remoteChainID, uint256 indexed index)
func (_StandardBridge *StandardBridgeFilterer) FilterPendingLocked(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*StandardBridgePendingLockedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "PendingLocked", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgePendingLockedIterator{contract: _StandardBridge.contract, event: "PendingLocked", logs: logs, sub: sub}, nil
}

// WatchPendingLocked is a free log subscription operation binding the contract event 0xa4e22101eb77eafdff61069f49689ebfe0c71e0dab0640d5c0b7bb11d1abb4f6.
//
// Solidity: event PendingLocked(uint256 indexed remoteChainID, uint256 indexed index)
func (_StandardBridge *StandardBridgeFilterer) WatchPendingLocked(opts *bind.WatchOpts, sink chan<- *StandardBridgePendingLocked, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "PendingLocked", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgePendingLocked)
				if err := _StandardBridge.contract.UnpackLog(event, "PendingLocked", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParsePendingLocked(log types.Log) (*StandardBridgePendingLocked, error) {
	event := new(StandardBridgePendingLocked)
	if err := _StandardBridge.contract.UnpackLog(event, "PendingLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the StandardBridge contract.
type StandardBridgeRoleUpdatedIterator struct {
	Event *StandardBridgeRoleUpdated // Event containing the contract specifics and raw log

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
func (it *StandardBridgeRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeRoleUpdated)
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
		it.Event = new(StandardBridgeRoleUpdated)
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
func (it *StandardBridgeRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeRoleUpdated represents a RoleUpdated event raised by the StandardBridge contract.
type StandardBridgeRoleUpdated struct {
	Account common.Address
	Role    [32]byte
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_StandardBridge *StandardBridgeFilterer) FilterRoleUpdated(opts *bind.FilterOpts, account []common.Address, role [][32]byte, status []bool) (*StandardBridgeRoleUpdatedIterator, error) {

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

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeRoleUpdatedIterator{contract: _StandardBridge.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_StandardBridge *StandardBridgeFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *StandardBridgeRoleUpdated, account []common.Address, role [][32]byte, status []bool) (event.Subscription, error) {

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

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeRoleUpdated)
				if err := _StandardBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseRoleUpdated(log types.Log) (*StandardBridgeRoleUpdated, error) {
	event := new(StandardBridgeRoleUpdated)
	if err := _StandardBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeSafetyLimitSetIterator is returned from FilterSafetyLimitSet and is used to iterate over the raw logs and unpacked data for SafetyLimitSet events raised by the StandardBridge contract.
type StandardBridgeSafetyLimitSetIterator struct {
	Event *StandardBridgeSafetyLimitSet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeSafetyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeSafetyLimitSet)
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
		it.Event = new(StandardBridgeSafetyLimitSet)
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
func (it *StandardBridgeSafetyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeSafetyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeSafetyLimitSet represents a SafetyLimitSet event raised by the StandardBridge contract.
type StandardBridgeSafetyLimitSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	SafetyLimit   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSafetyLimitSet is a free log retrieval operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 safetyLimit)
func (_StandardBridge *StandardBridgeFilterer) FilterSafetyLimitSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*StandardBridgeSafetyLimitSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeSafetyLimitSetIterator{contract: _StandardBridge.contract, event: "SafetyLimitSet", logs: logs, sub: sub}, nil
}

// WatchSafetyLimitSet is a free log subscription operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 safetyLimit)
func (_StandardBridge *StandardBridgeFilterer) WatchSafetyLimitSet(opts *bind.WatchOpts, sink chan<- *StandardBridgeSafetyLimitSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeSafetyLimitSet)
				if err := _StandardBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseSafetyLimitSet(log types.Log) (*StandardBridgeSafetyLimitSet, error) {
	event := new(StandardBridgeSafetyLimitSet)
	if err := _StandardBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the StandardBridge contract.
type StandardBridgeThresholdChangedIterator struct {
	Event *StandardBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *StandardBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeThresholdChanged)
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
		it.Event = new(StandardBridgeThresholdChanged)
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
func (it *StandardBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeThresholdChanged represents a ThresholdChanged event raised by the StandardBridge contract.
type StandardBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_StandardBridge *StandardBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*StandardBridgeThresholdChangedIterator, error) {

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &StandardBridgeThresholdChangedIterator{contract: _StandardBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_StandardBridge *StandardBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *StandardBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeThresholdChanged)
				if err := _StandardBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseThresholdChanged(log types.Log) (*StandardBridgeThresholdChanged, error) {
	event := new(StandardBridgeThresholdChanged)
	if err := _StandardBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the StandardBridge contract.
type StandardBridgeTokenPairRegisteredIterator struct {
	Event *StandardBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *StandardBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeTokenPairRegistered)
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
		it.Event = new(StandardBridgeTokenPairRegistered)
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
func (it *StandardBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the StandardBridge contract.
type StandardBridgeTokenPairRegistered struct {
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
func (_StandardBridge *StandardBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*StandardBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeTokenPairRegisteredIterator{contract: _StandardBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 exchangeRate, uint256 safetyLimit, bool isOrigin)
func (_StandardBridge *StandardBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *StandardBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeTokenPairRegistered)
				if err := _StandardBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*StandardBridgeTokenPairRegistered, error) {
	event := new(StandardBridgeTokenPairRegistered)
	if err := _StandardBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the StandardBridge contract.
type StandardBridgeTokenPairUnregisteredIterator struct {
	Event *StandardBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *StandardBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeTokenPairUnregistered)
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
		it.Event = new(StandardBridgeTokenPairUnregistered)
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
func (it *StandardBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the StandardBridge contract.
type StandardBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_StandardBridge *StandardBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*StandardBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeTokenPairUnregisteredIterator{contract: _StandardBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_StandardBridge *StandardBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *StandardBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeTokenPairUnregistered)
				if err := _StandardBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*StandardBridgeTokenPairUnregistered, error) {
	event := new(StandardBridgeTokenPairUnregistered)
	if err := _StandardBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the StandardBridge contract.
type StandardBridgeTokenPauseSetIterator struct {
	Event *StandardBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *StandardBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeTokenPauseSet)
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
		it.Event = new(StandardBridgeTokenPauseSet)
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
func (it *StandardBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeTokenPauseSet represents a TokenPauseSet event raised by the StandardBridge contract.
type StandardBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_StandardBridge *StandardBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*StandardBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeTokenPauseSetIterator{contract: _StandardBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_StandardBridge *StandardBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *StandardBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeTokenPauseSet)
				if err := _StandardBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseTokenPauseSet(log types.Log) (*StandardBridgeTokenPauseSet, error) {
	event := new(StandardBridgeTokenPauseSet)
	if err := _StandardBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StandardBridge contract.
type StandardBridgeUnpausedIterator struct {
	Event *StandardBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *StandardBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeUnpaused)
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
		it.Event = new(StandardBridgeUnpaused)
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
func (it *StandardBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeUnpaused represents a Unpaused event raised by the StandardBridge contract.
type StandardBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StandardBridge *StandardBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StandardBridgeUnpausedIterator, error) {

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StandardBridgeUnpausedIterator{contract: _StandardBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StandardBridge *StandardBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StandardBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeUnpaused)
				if err := _StandardBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseUnpaused(log types.Log) (*StandardBridgeUnpaused, error) {
	event := new(StandardBridgeUnpaused)
	if err := _StandardBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StandardBridge contract.
type StandardBridgeUpgradedIterator struct {
	Event *StandardBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *StandardBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardBridgeUpgraded)
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
		it.Event = new(StandardBridgeUpgraded)
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
func (it *StandardBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardBridgeUpgraded represents a Upgraded event raised by the StandardBridge contract.
type StandardBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StandardBridge *StandardBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StandardBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StandardBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StandardBridgeUpgradedIterator{contract: _StandardBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StandardBridge *StandardBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StandardBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StandardBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardBridgeUpgraded)
				if err := _StandardBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StandardBridge *StandardBridgeFilterer) ParseUpgraded(log types.Log) (*StandardBridgeUpgraded, error) {
	event := new(StandardBridgeUpgraded)
	if err := _StandardBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
