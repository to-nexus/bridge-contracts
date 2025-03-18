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

// IBaseBridgeBridgeTokenArguments is an auto generated low-level Go binding around an user-defined struct.
type IBaseBridgeBridgeTokenArguments struct {
	ToChainID *big.Int
	FromToken common.Address
	To        common.Address
	Value     *big.Int
	GasFee    *big.Int
	ExFee     *big.Int
	ExtraData []byte
}

// IBaseBridgePermitArguments is an auto generated low-level Go binding around an user-defined struct.
type IBaseBridgePermitArguments struct {
	Token    common.Address
	Account  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// IBridgeRegistryFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryFinalizeArguments struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	ExtraData   []byte
}

// IBridgeRegistryPendingData is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryPendingData struct {
	Args            IBridgeRegistryFinalizeArguments
	Status          uint8
	DelayExpiration *big.Int
	Reason          []byte
}

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryTokenPair struct {
	LocalToken    common.Address
	RemoteToken   common.Address
	Paused        bool
	IsOrigin      bool
	Deposited     *big.Int
	PendingAmount *big.Int
}

// BaseBridgeMetaData contains all meta data concerning the BaseBridge contract.
var BaseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualProcessPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sLength\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"b33eb36e": "getPendingArguments(uint256,uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"814914b5": "getTokenPair(uint256,address)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"fd47852b": "grantRoleBatch(bytes32,address[])",
		"91d14854": "hasRole(bytes32,address)",
		"f35f9e45": "initialize(address,uint8,address)",
		"91cf6d3e": "initializedAt()",
		"8bb19439": "isPending(uint256,uint256)",
		"7f4ab9f5": "manualProcessPending(uint256,uint256)",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"edbbea39": "registerToken(uint256,bool,address,address)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"9f9b4f1c": "releasePendingBatch(uint256[],uint256[])",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"82f51fa6": "revokeRoleBatch(bytes32,address[])",
		"0b1d8d06": "setBridgeVerifier(address)",
		"b8aa837e": "setChainPause(uint256,bool)",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"bedb86fb": "setPause(bool)",
		"94ddc8c6": "setTokenPause(uint256,address,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615baa6100f95f395f818161315b0152818161318401526132c50152615baa5ff3fe6080604052600436106102dc575f3560e01c806391cf6d3e11610189578063cf56118e116100d8578063e2af6cd711610092578063f35f9e451161006d578063f35f9e45146109ad578063f4509643146109cc578063f698da25146109eb578063fd47852b146109ff575f5ffd5b8063e2af6cd714610950578063edbbea391461096f578063f0702e8e1461098e575f5ffd5b8063cf56118e14610892578063d477f05f146108a6578063d547741f146108c5578063d5717fc5146108e4578063d605665b14610903578063df6e87dc14610916575f5ffd5b8063aa1bd0bc11610143578063b33eb36e1161011e578063b33eb36e14610809578063b7f3358d14610835578063b8aa837e14610854578063bedb86fb14610873575f5ffd5b8063aa1bd0bc1461078e578063ad3cb1cc146107ad578063ae6893f8146107ea575f5ffd5b806391cf6d3e146106de57806391d14854146106f257806394ddc8c6146107115780639f9b4f1c14610730578063a217fddf1461074f578063a3246ad314610762575f5ffd5b806352d1902d116102455780637f4ab9f5116101ff57806384b0196e116101da57806384b0196e1461066857806388d67d6d1461068f5780638bb19439146106a257806391cca3db146106c1575f5ffd5b80637f4ab9f514610555578063814914b51461057457806382f51fa614610649575f5ffd5b806352d1902d146104945780635b605f5c146104a85780635c975abb146104d45780635fd262de146104f7578063670e62681461050a5780637921487414610529575f5ffd5b806342cde4e81161029657806342cde4e8146103d85780634be13f83146103f95780634d5d0056146104305780634ee078ba146104435780634f1ef28614610462578063502cc5c014610475575f5ffd5b806301ffc9a7146103075780630b1d8d061461033b5780631938e0f21461035a578063248a9ca31461036d5780632f2ff15d1461039a57806336568abe146103b9575f5ffd5b3661030357345f03610301576040516304a4cdd960e51b815260040160405180910390fd5b005b5f5ffd5b348015610312575f5ffd5b506103266103213660046148e1565b610a1e565b60405190151581526020015b60405180910390f35b348015610346575f5ffd5b5061030161035536600461491c565b610a54565b610326610368366004614aa6565b610acf565b348015610378575f5ffd5b5061038c610387366004614b5f565b610e25565b604051908152602001610332565b3480156103a5575f5ffd5b506103016103b4366004614b76565b610e45565b3480156103c4575f5ffd5b506103016103d3366004614b76565b610e67565b3480156103e3575f5ffd5b5060645460405160ff9091168152602001610332565b348015610404575f5ffd5b50603254610418906001600160a01b031681565b6040516001600160a01b039091168152602001610332565b61032661043e366004614be8565b610e9f565b34801561044e575f5ffd5b5061032661045d366004614c8a565b611099565b610301610470366004614d20565b6113a1565b348015610480575f5ffd5b5061030161048f366004614d6c565b6113c0565b34801561049f575f5ffd5b5061038c611477565b3480156104b3575f5ffd5b506104c76104c2366004614b5f565b611492565b6040516103329190614dde565b3480156104df575f5ffd5b505f516020615b355f395f51905f525460ff16610326565b610326610505366004614e2b565b611606565b348015610515575f5ffd5b50610418610524366004614eb3565b61166f565b348015610534575f5ffd5b50610548610543366004614b5f565b61171e565b6040516103329190614f65565b348015610560575f5ffd5b5061032661056f366004614c8a565b611737565b34801561057f575f5ffd5b5061063c61058e366004614b76565b6040805160c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810191909152505f9182526038602090815260408084206001600160a01b039384168552825292839020835160c08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b90920490921615156060830152600281015460808301526003015460a082015290565b6040516103329190614f77565b348015610654575f5ffd5b50610301610663366004614f85565b611791565b348015610673575f5ffd5b5061067c6117da565b6040516103329796959493929190615058565b61032661069d366004615185565b611883565b3480156106ad575f5ffd5b506103266106bc366004614c8a565b611920565b3480156106cc575f5ffd5b506097546001600160a01b0316610418565b3480156106e9575f5ffd5b5060985461038c565b3480156106fd575f5ffd5b5061032661070c366004614b76565b61193e565b34801561071c575f5ffd5b5061030161072b3660046152c2565b611974565b34801561073b575f5ffd5b5061032661074a366004615301565b611a4c565b34801561075a575f5ffd5b5061038c5f81565b34801561076d575f5ffd5b5061078161077c366004614b5f565b611ac9565b604051610332919061535a565b348015610799575f5ffd5b506103016107a8366004614b5f565b611ae2565b3480156107b8575f5ffd5b506107dd604051806040016040528060058152602001640352e302e360dc1b81525081565b604051610332919061539a565b3480156107f5575f5ffd5b5061038c610804366004614b5f565b611b31565b348015610814575f5ffd5b50610828610823366004614c8a565b611b4d565b60405161033291906153e0565b348015610840575f5ffd5b5061030161084f366004615486565b611d2b565b34801561085f575f5ffd5b5061030161086e36600461549f565b611d9a565b34801561087e575f5ffd5b5061030161088d3660046154c2565b611e37565b34801561089d575f5ffd5b50610548611e62565b3480156108b1575f5ffd5b506103016108c036600461491c565b611e73565b3480156108d0575f5ffd5b506103016108df366004614b76565b611ec7565b3480156108ef575f5ffd5b5061038c6108fe366004614b5f565b611ee3565b6103016109113660046154dd565b611eff565b348015610921575f5ffd5b50610935610930366004615576565b612092565b60408051938452602084019290925290820152606001610332565b34801561095b575f5ffd5b5061030161096a36600461491c565b612121565b34801561097a575f5ffd5b506103016109893660046155ab565b6121d4565b348015610999575f5ffd5b50609654610418906001600160a01b031681565b3480156109b8575f5ffd5b506103016109c73660046155fb565b612431565b3480156109d7575f5ffd5b506103016109e6366004614b76565b612541565b3480156109f6575f5ffd5b5061038c612613565b348015610a0a575f5ffd5b50610301610a19366004614f85565b61261c565b5f6001600160e01b03198216637965db0b60e01b1480610a4e57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f610a5e81612665565b6001600160a01b038216610a8557604051638219abe360e01b815260040160405180910390fd5b609680546001600160a01b0319166001600160a01b0384169081179091556040517fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3905f90a25050565b5f610ad8612672565b610ae06126a4565b610b08610af3606087016040880161491c565b86355f908152603760205260409020906126db565b610b18606087016040880161491c565b90610b4757604051633142cb1160e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610b72576040516362624a5160e11b815260048101929092526024820152604401610b3e565b505084355f9081526036602052604081206002018054600101908190559050806020870135808214610bc05760405163a6ab65c560e01b815260048101929092526024820152604401610b3e565b505f90507fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191487356020890135610bfc60608b0160408c0161491c565b610c0c60808c0160608d0161491c565b60808c0135610c1e60a08e018e615636565b604051602001610c359897969594939291906156a0565b604051602081830303815290604052805190602001209050610c59818787876126fc565b5f606081610c7c8a35610c718c850160408e0161491c565b8c608001355f612909565b90935090506001836009811115610c9557610c956153ac565b03610ccf57610cc98a35610caf60608d0160408e0161491c565b610cbf60808e0160608f0161491c565b8d60800135612a42565b90935091505b6001836009811115610ce357610ce36153ac565b03610d7c57610cf860608b0160408c0161491c565b6001600160a01b03168a602001358b5f01357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8d6060016020810190610d3e919061491c565b8e6080013542604051610d6f939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a4610dfe565b610d888a848484612ae4565b610d9860608b0160408c0161491c565b6001600160a01b03168a602001358b5f01357f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f8d6060016020810190610dde919061491c565b8e608001354289604051610df594939291906156e7565b60405180910390a45b600195505050505050610e1d60015f516020615b555f395f51905f5255565b949350505050565b5f9081525f516020615b155f395f51905f52602052604090206001015490565b610e4e82610e25565b610e5781612665565b610e618383612cd6565b50505050565b6001600160a01b0381163314610e905760405163334bd91960e11b815260040160405180910390fd5b610e9a8282612d3a565b505050565b5f610ea8612672565b8989610eb48282612d95565b610ebc6126a4565b610ec9602085018561491c565b6001600160a01b038c81169116148b610ee5602087018761491c565b9091610f17576040516313f7c32b60e31b81526001600160a01b03928316600482015291166024820152604401610b3e565b5050610f268c8c8b8b8b612e6b565b909850965086610f36898b615726565b610f409190615726565b6040850135101587610f528a8c615726565b610f5c9190615726565b85604001359091610f895760405163943892eb60e01b815260048101929092526024820152604401610b3e565b50506001600160a01b038b1663d505accf610faa604087016020880161491c565b3060408801356060890135610fc560a08b0160808c01615486565b6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529490931660248501526044840191909152606483015260ff16608482015260a087013560a482015260c087013560c482015260e4015f604051808303815f87803b158015611035575f5ffd5b505af1158015611047573d5f5f3e3d5ffd5b505050506110708c8c866020016020810190611063919061491c565b8d8d8d8d60018e8e612f4f565b6001925061108a60015f516020615b555f395f51905f5255565b50509998505050505050505050565b5f6110a2612672565b6110aa6126a4565b5f8381526039602052604090206110c19083613031565b82906110e3576040516373a1390160e11b8152600401610b3e91815260200190565b505f838152603a6020908152604080832085845290915280822081516101408101909252805460808301908152600182015460a084015260028201546001600160a01b0390811660c085015260038301541660e0840152600482015461010084015260058201805484929184916101208501919061116090615739565b80601f016020809104026020016040519081016040528092919081815260200182805461118c90615739565b80156111d75780601f106111ae576101008083540402835291602001916111d7565b820191905f5260205f20905b8154815290600101906020018083116111ba57829003601f168201915b505050919092525050508152600682015460209091019060ff166009811115611202576112026153ac565b6009811115611213576112136153ac565b81526020016007820154815260200160088201805461123190615739565b80601f016020809104026020016040519081016040528092919081815260200182805461125d90615739565b80156112a85780601f1061127f576101008083540402835291602001916112a8565b820191905f5260205f20905b81548152906001019060200180831161128b57829003601f168201915b50505050508152505090505f6112cf85835f015160400151845f0151608001516001612909565b50905060018160098111156112e6576112e66153ac565b148160098111156112f9576112f96153ac565b60405160200161130b91815260200190565b604051602081830303815290604052906113385760405162461bcd60e51b8152600401610b3e919061539a565b506040820151158061134d5750428260400151105b826040015142909161137b57604051637a88099160e11b815260048101929092526024820152604401610b3e565b50506113878585613048565b92505050610a4e60015f516020615b555f395f51905f5255565b6113a9613150565b6113b2826131f4565b6113bc82826131fe565b5050565b672b22a924a324a2a960c11b6113d581612665565b5f8481526039602052604090206113ec9084613031565b839061140e576040516373a1390160e11b8152600401610b3e91815260200190565b505f61141a8342615726565b5f868152603a602090815260408083208884528252918290206007018390559051828152919250859187917fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949910160405180910390a35050505050565b5f6114806132ba565b505f516020615af55f395f51905f5290565b5f818152603760205260408120606091906114ac90613303565b90505f81516001600160401b038111156114c8576114c8614937565b60405190808252806020026020018201604052801561152657816020015b6040805160c0810182525f8082526020808301829052928201819052606082018190526080820181905260a082015282525f199092019101816114e65790505b5090505f5b82518110156115fe5760385f8681526020019081526020015f205f84838151811061155857611558615771565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160c08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b909304909116151560608201526002820154608082015260039091015460a082015282518390839081106115eb576115eb615771565b602090810291909101015260010161152b565b509392505050565b5f61160f612672565b888861161b8282612d95565b6116236126a4565b6116308b8b8a8a8a612e6b565b90975095506116478b8b338c8c8c8c5f8d8d612f4f565b6001925061166160015f516020615b555f395f51905f5255565b505098975050505050505050565b6032545f906001600160a01b031661169a576040516315aeca0d60e11b815260040160405180910390fd5b603254604051637c469ea160e11b81526001600160a01b039091169063f88d3d42906116d0908890889088908890600401615785565b6020604051808303815f875af11580156116ec573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061171091906157c2565b9050610e1d855f83876121d4565b5f818152603960205260409020606090610a4e90613303565b5f672b22a924a324a2a960c11b61174d81612665565b5f8481526039602052604090206117649084613031565b8390611786576040516373a1390160e11b8152600401610b3e91815260200190565b50610e1d8484613048565b61179a82610e25565b6117a381612665565b5f5b8251811015610e61576117d1848483815181106117c4576117c4615771565b6020026020010151612d3a565b506001016117a5565b5f60608082808083815f516020615ad55f395f51905f52805490915015801561180557506001810154155b6118495760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610b3e565b61185161330f565b6118596133cf565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015611911576119088787838181106118a3576118a3615771565b90506020028101906118b591906157dd565b8683815181106118c7576118c7615771565b60200260200101518684815181106118e1576118e1615771565b60200260200101518685815181106118fb576118fb615771565b6020026020010151610acf565b50600101611886565b50600190505b95945050505050565b5f8281526039602052604081206119379083613031565b9392505050565b5f9182525f516020615b155f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6727a822a920aa27a960c11b61198981612665565b5f8481526037602052604090206119a090846126db565b83906119cb5760405163153096f360e11b81526001600160a01b039091166004820152602401610b3e565b505f8481526038602090815260408083206001600160a01b0387168085529252918290206001018054851515600160a01b0260ff60a01b19909116179055905185907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea90611a3e90861515815260200190565b60405180910390a350505050565b5f8151835114611a6f576040516329f517fd60e01b815260040160405180910390fd5b5f5b8251811015611abf57611ab6848281518110611a8f57611a8f615771565b6020026020010151848381518110611aa957611aa9615771565b6020026020010151611099565b50600101611a71565b5060019392505050565b5f818152602081905260409020606090610a4e90613303565b6420a226a4a760d91b611af481612665565b60338290556040518281527fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3906020015b60405180910390a15050565b5f818152603660205260408120600190810154610a4e91615726565b611b55614828565b5f838152603a602090815260408083208584529091529081902081516101408101909252805460808301908152600182015460a084015260028201546001600160a01b0390811660c085015260038301541660e08401526004820154610100840152600582018054849291849161012085019190611bd290615739565b80601f0160208091040260200160405190810160405280929190818152602001828054611bfe90615739565b8015611c495780601f10611c2057610100808354040283529160200191611c49565b820191905f5260205f20905b815481529060010190602001808311611c2c57829003601f168201915b505050919092525050508152600682015460209091019060ff166009811115611c7457611c746153ac565b6009811115611c8557611c856153ac565b815260200160078201548152602001600882018054611ca390615739565b80601f0160208091040260200160405190810160405280929190818152602001828054611ccf90615739565b8015611d1a5780601f10611cf157610100808354040283529160200191611d1a565b820191905f5260205f20905b815481529060010190602001808311611cfd57829003601f168201915b505050505081525050905092915050565b5f611d3581612665565b8160ff165f03611d585760405163f0f15b9160e01b815260040160405180910390fd5b6064805460ff191660ff84169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff90602001611b25565b6727a822a920aa27a960c11b611daf81612665565b611dba603484613031565b8390611ddc5760405163ac7dbbfd60e01b8152600401610b3e91815260200190565b505f83815260366020908152604091829020600301805460ff1916851515908117909155915191825284917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a2505050565b6727a822a920aa27a960c11b611e4c81612665565b8115611e5a576113bc61340d565b6113bc61346f565b6060611e6e6034613303565b905090565b5f611e7d81612665565b6001600160a01b038216611ea457604051638219abe360e01b815260040160405180910390fd5b50609780546001600160a01b0319166001600160a01b0392909216919091179055565b611ed082610e25565b611ed981612665565b610e618383612d3a565b5f81815260366020526040812060020154610a4e906001615726565b828114611f1f576040516329f517fd60e01b815260040160405180910390fd5b5f5b8381101561208b57612082858583818110611f3e57611f3e615771565b9050602002810190611f5091906157fb565b35868684818110611f6357611f63615771565b9050602002810190611f7591906157fb565b611f8690604081019060200161491c565b878785818110611f9857611f98615771565b9050602002810190611faa91906157fb565b611fbb90606081019060400161491c565b888886818110611fcd57611fcd615771565b9050602002810190611fdf91906157fb565b60600135898987818110611ff557611ff5615771565b905060200281019061200791906157fb565b608001358a8a8881811061201d5761201d615771565b905060200281019061202f91906157fb565b60a001358b8b8981811061204557612045615771565b905060200281019061205791906157fb565b6120659060c0810190615636565b8b8b8b81811061207757612077615771565b905060e00201610e9f565b50600101611f21565b5050505050565b6096546040516337dba1f760e21b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063df6e87dc90606401606060405180830381865afa1580156120ee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612112919061580f565b91989097509095509350505050565b5f61212b81612665565b6032546001600160a01b031680156121625760405163f6c75f8f60e01b81526001600160a01b039091166004820152602401610b3e565b506001600160a01b03821661218a57604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0384169081179091556040517fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee905f90a25050565b6420a226a4a760d91b6121e681612665565b6121f16034866134b4565b1561224e57604080516080810182528681525f6020808301828152838501838152606085018481528b855260369093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b03831661227557604051636ca1fdd760e01b815260040160405180910390fd5b5f85815260376020526040902061228c90846134bf565b83906122b7576040516311dd05f360e31b81526001600160a01b039091166004820152602401610b3e565b506040518060c00160405280846001600160a01b03168152602001836001600160a01b031681526020015f1515815260200185151581526020015f81526020015f81525060385f8781526020019081526020015f205f856001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a08201518160030155905050816001600160a01b0316836001600160a01b0316867fe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db87604051612422911515815260200190565b60405180910390a45050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156124755750825b90505f826001600160401b031660011480156124905750303b155b90508115801561249e575080155b156124bc5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156124e657845460ff60401b1916600160401b1785555b6124f18888886134d3565b831561253757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6420a226a4a760d91b61255381612665565b5f83815260376020526040902061256a908361355c565b82906125955760405163153096f360e11b81526001600160a01b039091166004820152602401610b3e565b505f8381526038602090815260408083206001600160a01b038616808552925280832080546001600160a01b03191681556001810180546001600160b01b03191690556002810184905560030183905551909185917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a3505050565b5f611e6e613570565b61262582610e25565b61262e81612665565b5f5b8251811015610e615761265c8484838151811061264f5761264f615771565b6020026020010151612cd6565b50600101612630565b61266f8133613579565b50565b5f516020615b355f395f51905f525460ff16156126a25760405163d93c066560e01b815260040160405180910390fd5b565b5f516020615b555f395f51905f528054600119016126d557604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f9081526001830160205260408120541515611937565b825182518114801561270e5750815181145b835183518392612742576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610b3e565b505060645482915060ff1681101561277057604051631aebd17960e11b8152600401610b3e91815260200190565b505f80826001600160401b0381111561278b5761278b614937565b6040519080825280602002602001820160405280156127b4578160200160208202803683370190505b5090505f5b838110156128dd575f6128248883815181106127d7576127d7615771565b60200260200101518884815181106127f1576127f1615771565b602002602001015188858151811061280b5761280b615771565b602002602001015161281c8d6135b2565b9291906135de565b90505f805b858110156128745784818151811061284357612843615771565b60200260200101516001600160a01b0316836001600160a01b03160361286c5760019150612874565b600101612829565b50801580156128945750612894682b20a624a220aa27a960b91b8361193e565b156128d357818486815181106128ac576128ac615771565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b50506001016127b9565b50606454829060ff1681101561253757604051631aebd17960e11b8152600401610b3e91815260200190565b5f8481526038602090815260408083206001600160a01b038088168552908352818420825160c08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b90920490911615156060840152600281015460808401526003015460a08301528291906129975760025f9250925050612a39565b83612a3057609654604051633f4bdec560e01b81526001600160a01b0388811660048301526024820188905290911690633f4bdec5906044016020604051808303815f875af11580156129ec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a10919061583a565b9250826009811115612a2457612a246153ac565b600114612a3057600191505b60015f92509250505b94509492505050565b5f60605f196001600160a01b03861601612a8a575f612a70858560405180602001604052805f81525061360a565b9250905080612a7f575f612a82565b60015b925050612a39565b8215612a39575f8681526038602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff1615612ad957612ad086868686613787565b91509150612a39565b612ad0858585613860565b83355f908152603960209081526040909120612b02918601356134b4565b846020013590612b28576040516307a5c91d60e31b8152600401610b3e91815260200190565b50604051806080016040528085612b3e90615858565b8152602001846009811115612b5557612b556153ac565b815260200182612b65575f612b72565b603354612b729042615726565b8152602090810184905285355f908152603a82526040808220888401358352835290819020835180518255928301516001820155908201516002820180546001600160a01b039283166001600160a01b03199182161790915560608401516003840180549190931691161790556080820151600482015560a082015190919082906005820190612c029082615922565b505050602082015160068201805460ff19166001836009811115612c2857612c286153ac565b02179055506040820151600782015560608201516008820190612c4b9082615922565b50505083355f90815260386020526040808220908290612c71906060890190890161491c565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff161561208b578460800135816003015f828254612cb79190615726565b90915550505050505050565b60015f516020615b555f395f51905f5255565b5f612ce18383613927565b90508015610a4e575f838152602081905260409020612d0090836134bf565b82849091612d325760405163d180cb3160e01b81526001600160a01b0390921660048301526024820152604401610b3e565b505092915050565b5f612d4583836139c8565b90508015610a4e575f838152602081905260409020612d64908361355c565b82849091612d325760405162a95f1b60e31b81526001600160a01b0390921660048301526024820152604401610b3e565b5f828152603760205260409020612dac90826126db565b8190612dd75760405163153096f360e11b81526001600160a01b039091166004820152602401610b3e565b505f82815260366020526040902060030154829060ff1615612e0f57604051636fc24b7960e11b8152600401610b3e91815260200190565b505f8281526038602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610e9a576040516338384f6f60e11b81526001600160a01b039091166004820152602401610b3e565b6096545f9081906001600160a01b0316612e8957505f905080612f45565b6096546040516337dba1f760e21b8152600481018990526001600160a01b038881166024830152604482018890525f92169063df6e87dc90606401606060405180830381865afa158015612edf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f03919061580f565b90945092509050808610801590612f1a5750828510155b8015612f265750818410155b612f43576040516358e8878560e01b815260040160405180910390fd5b505b9550959350505050565b612f658a8a8a89612f60898b615726565b613a41565b5f8a815260366020526040812060019081018054909101908190558190915060385f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b03169050896001600160a01b0316828d7f732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db508e858e8e8e8e428f8f8f60405161301b9a999897969594939291906159dc565b60405180910390a4505050505050505050505050565b5f8181526001830160205260408120541515611937565b5f5f6130548484613c49565b90505f5f613073835f0151846040015185606001518660800151612a42565b9092509050600182600981111561308c5761308c6153ac565b1482826040516020016130a0929190615a46565b604051602081830303815290604052906130cd5760405162461bcd60e51b8152600401610b3e919061539a565b5082604001516001600160a01b03168360200151845f01517f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b866060015187608001514260405161313c939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806131d657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166131ca5f516020615af55f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156126a25760405163703e46dd60e11b815260040160405180910390fd5b5f6113bc81612665565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613258575060408051601f3d908101601f1916820190925261325591810190615a65565b60015b61328057604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610b3e565b5f516020615af55f395f51905f5281146132b057604051632a87526960e21b815260048101829052602401610b3e565b610e9a8383613e62565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146126a25760405163703e46dd60e11b815260040160405180910390fd5b60605f61193783613eb7565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615ad55f395f51905f529161334d90615739565b80601f016020809104026020016040519081016040528092919081815260200182805461337990615739565b80156133c45780601f1061339b576101008083540402835291602001916133c4565b820191905f5260205f20905b8154815290600101906020018083116133a757829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615ad55f395f51905f529161334d90615739565b613415612672565b5f516020615b355f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b613477613f10565b5f516020615b355f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613451565b5f6119378383613f3f565b5f611937836001600160a01b038416613f3f565b6134db613f8b565b6134e3613fd4565b6134eb613fdc565b6134f3613fec565b6134fc83613ffc565b61350582614016565b61350d61409f565b6001600160a01b03811661353457604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0392909216919091179055505043609855565b5f611937836001600160a01b0384166140b0565b5f611e6e61418a565b613583828261193e565b6113bc5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610b3e565b5f610a4e6135be613570565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6135ee888888886141fd565b9250925092506135fe82826142c5565b50909695505050505050565b5f606083474782111561363957604051632f2a246160e21b815260048101929092526024820152604401610b3e565b50506060856001600160a01b031685856040516136569190615a7c565b5f6040518083038185875af1925050503d805f8114613690576040519150601f19603f3d011682016040523d82523d5f602084013e613695565b606091505b509093509050826136da578051156136b1575f9250905061377f565b50506040805180820190915260088152671c995d995c9d195960c21b60208201525f915061377f565b845f0361376a5780515f0361372857856001600160a01b03163b5f036137235750506040805180820190915260088152676e6f7420636f646560c01b60208201525f915061377f565b61376a565b6020810151600181151514613768575f6040518060400160405280600c81526020016b72657475726e2066616c736560a01b81525093509350505061377f565b505b505060408051602081019091525f8152600191505b935093915050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af19250505080156137f6575060408051601f3d908101601f191682019092526137f391810190615a92565b60015b613834573d808015613823576040519150601f19603f3d011682016040523d82523d5f602084013e613828565b606091505b50600592509050612a39565b80613840576005613843565b60015b925080156138565761385687878661437d565b5094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af19250505080156138cf575060408051601f3d908101601f191682019092526138cc91810190615a92565b60015b61390d573d8080156138fc576040519150601f19603f3d011682016040523d82523d5f602084013e613901565b606091505b5060069250905061377f565b8061391957600661391c565b60015b925050935093915050565b5f5f516020615b155f395f51905f52613940848461193e565b6139bf575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556139753390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a4e565b5f915050610a4e565b5f5f516020615b155f395f51905f526139e1848461193e565b156139bf575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a4e565b5f196001600160a01b03851601613af157613a5c8183615726565b3414613a688284615726565b349091613a91576040516362624a5160e11b815260048101929092526024820152604401610b3e565b50508015613aec5760975460408051602081019091525f808252918291613ac3916001600160a01b031690859061360a565b91509150818190613ae85760405163e8b5c4bb60e01b8152600401610b3e919061539a565b5050505b61208b565b5f348015613b1b576040516362624a5160e11b815260048101929092526024820152604401610b3e565b50613b3f90508330613b2d8486615726565b6001600160a01b038816929190614428565b8015613b5f57609754613b5f906001600160a01b0386811691168361448f565b5f8581526038602090815260408083206001600160a01b0388168452909152902060010154600160a81b900460ff1615613b9e57613aec8585846144c0565b604051632770a7eb60e21b8152306004820152602481018390526001600160a01b03851690639dc29fac906044016020604051808303815f875af1158015613be8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c0c9190615a92565b84848490919261253757604051639ac2f56d60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610b3e565b613c51614854565b5f838152603960205260409020613c6890836144fe565b8290613c8a57604051637f11bea960e01b8152600401610b3e91815260200190565b505f838152603a60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190613d0290615739565b80601f0160208091040260200160405190810160405280929190818152602001828054613d2e90615739565b8015613d795780601f10613d5057610100808354040283529160200191613d79565b820191905f5260205f20905b815481529060010190602001808311613d5c57829003601f168201915b505050919092525050505f848152603860209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615613ddc578160800151816003015f828254613dd69190615aad565b90915550505b5f848152603a602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181613e356005830182614897565b505060068201805460ff191690555f60078301819055613e59906008840190614897565b50505092915050565b613e6b82614509565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613eaf57610e9a828261456c565b6113bc6145d5565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613f0457602002820191905f5260205f20905b815481526020019060010190808311613ef0575b50505050509050919050565b5f516020615b355f395f51905f525460ff166126a257604051638dfc202b60e01b815260040160405180910390fd5b5f818152600183016020526040812054613f8457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a4e565b505f610a4e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166126a257604051631afcd79f60e31b815260040160405180910390fd5b6126a2613f8b565b613fe4613f8b565b6126a26145f4565b613ff4613f8b565b6126a2614614565b614004613f8b565b61400c613fd4565b6113bc5f82612cd6565b61401e613f8b565b8060ff165f036140415760405163f0f15b9160e01b815260040160405180910390fd5b614089604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061461c565b6064805460ff191660ff92909216919091179055565b6140a7613f8b565b62015180603355565b5f81815260018301602052604081205480156139bf575f6140d2600183615aad565b85549091505f906140e590600190615aad565b9050808214614144575f865f01828154811061410357614103615771565b905f5260205f200154905080875f01848154811061412357614123615771565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061415557614155615ac0565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a4e565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6141b461462e565b6141bc614696565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561423657505f915060039050826142bb565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614287573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166142b257505f9250600191508290506142bb565b92505f91508190505b9450945094915050565b5f8260038111156142d8576142d86153ac565b036142e1575050565b60018260038111156142f5576142f56153ac565b036143135760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115614327576143276153ac565b036143485760405163fce698f760e01b815260048101829052602401610b3e565b600382600381111561435c5761435c6153ac565b036113bc576040516335e2f38360e21b815260048101829052602401610b3e565b5f8381526038602090815260408083206001600160a01b0386168452909152902060038101546143ad9083615726565b600282015460038301548692869286929091821015614405576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610b3e565b505050505081816002015f82825461441d9190615aad565b909155505050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610e619186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506146d8565b6040516001600160a01b03838116602483015260448201839052610e9a91859182169063a9059cbb9060640161445d565b5f8381526038602090815260408083206001600160a01b0386168452909152812060020180548392906144f4908490615726565b9091555050505050565b5f61193783836140b0565b806001600160a01b03163b5f0361453e57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610b3e565b5f516020615af55f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516145889190615a7c565b5f60405180830381855af49150503d805f81146145c0576040519150601f19603f3d011682016040523d82523d5f602084013e6145c5565b606091505b5091509150611917858383614744565b34156126a25760405163b398979f60e01b815260040160405180910390fd5b6145fc613f8b565b5f516020615b355f395f51905f52805460ff19169055565b612cc3613f8b565b614624613f8b565b6113bc82826147a0565b5f5f516020615ad55f395f51905f528161464661330f565b80519091501561465e57805160209091012092915050565b8154801561466d579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615ad55f395f51905f52816146ae6133cf565b8051909150156146c657805160209091012092915050565b6001820154801561466d579392505050565b5f5f60205f8451602086015f885af1806146f7576040513d5f823e3d81fd5b50505f513d9150811561470e57806001141561471b565b6001600160a01b0384163b155b15610e6157604051635274afe760e01b81526001600160a01b0385166004820152602401610b3e565b60608261475957614754826147ff565b611937565b815115801561477057506001600160a01b0384163b155b1561479957604051639996b31560e01b81526001600160a01b0385166004820152602401610b3e565b5092915050565b6147a8613f8b565b5f516020615ad55f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026147e18482615922565b50600381016147f08382615922565b505f8082556001909101555050565b80511561480f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b604051806080016040528061483b614854565b81526020015f81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5080546148a390615739565b5f825580601f106148b2575050565b601f0160209004905f5260205f209081019061266f91905b808211156148dd575f81556001016148ca565b5090565b5f602082840312156148f1575f5ffd5b81356001600160e01b031981168114611937575f5ffd5b6001600160a01b038116811461266f575f5ffd5b5f6020828403121561492c575f5ffd5b813561193781614908565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b038111828210171561496d5761496d614937565b60405290565b604051601f8201601f191681016001600160401b038111828210171561499b5761499b614937565b604052919050565b5f6001600160401b038211156149bb576149bb614937565b5060051b60200190565b803560ff811681146149d5575f5ffd5b919050565b5f82601f8301126149e9575f5ffd5b81356149fc6149f7826149a3565b614973565b8082825260208201915060208360051b860101925085831115614a1d575f5ffd5b602085015b83811015614a4157614a33816149c5565b835260209283019201614a22565b5095945050505050565b5f82601f830112614a5a575f5ffd5b8135614a686149f7826149a3565b8082825260208201915060208360051b860101925085831115614a89575f5ffd5b602085015b83811015614a41578035835260209283019201614a8e565b5f5f5f5f60808587031215614ab9575f5ffd5b84356001600160401b03811115614ace575f5ffd5b850160c08188031215614adf575f5ffd5b935060208501356001600160401b03811115614af9575f5ffd5b614b05878288016149da565b93505060408501356001600160401b03811115614b20575f5ffd5b614b2c87828801614a4b565b92505060608501356001600160401b03811115614b47575f5ffd5b614b5387828801614a4b565b91505092959194509250565b5f60208284031215614b6f575f5ffd5b5035919050565b5f5f60408385031215614b87575f5ffd5b823591506020830135614b9981614908565b809150509250929050565b5f5f83601f840112614bb4575f5ffd5b5081356001600160401b03811115614bca575f5ffd5b602083019150836020828501011115614be1575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c0811215614c02575f5ffd5b8a35995060208b0135614c1481614908565b985060408b0135614c2481614908565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115614c53575f5ffd5b614c5f8d828e01614ba4565b90955093505060e060df1982011215614c76575f5ffd5b5060e08a0190509295985092959850929598565b5f5f60408385031215614c9b575f5ffd5b50508035926020909101359150565b5f5f6001600160401b03841115614cc357614cc3614937565b50601f8301601f1916602001614cd881614973565b915050828152838383011115614cec575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112614d11575f5ffd5b61193783833560208501614caa565b5f5f60408385031215614d31575f5ffd5b8235614d3c81614908565b915060208301356001600160401b03811115614d56575f5ffd5b614d6285828601614d02565b9150509250929050565b5f5f5f60608486031215614d7e575f5ffd5b505081359360208301359350604090920135919050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a090810151910152565b602080825282518282018190525f918401906040840190835b81811015614e2057614e0a838551614d95565b6020939093019260c09290920191600101614df7565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215614e42575f5ffd5b883597506020890135614e5481614908565b96506040890135614e6481614908565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115614e93575f5ffd5b614e9f8b828c01614ba4565b999c989b5096995094979396929594505050565b5f5f5f5f60808587031215614ec6575f5ffd5b843593506020850135614ed881614908565b925060408501356001600160401b03811115614ef2575f5ffd5b8501601f81018713614f02575f5ffd5b614f1187823560208401614caa565b925050614f20606086016149c5565b905092959194509250565b5f8151808452602084019350602083015f5b82811015614f5b578151865260209586019590910190600101614f3d565b5093949350505050565b602081525f6119376020830184614f2b565b60c08101610a4e8284614d95565b5f5f60408385031215614f96575f5ffd5b8235915060208301356001600160401b03811115614fb2575f5ffd5b8301601f81018513614fc2575f5ffd5b8035614fd06149f7826149a3565b8082825260208201915060208360051b850101925087831115614ff1575f5ffd5b6020840193505b8284101561501c57833561500b81614908565b825260209384019390910190614ff8565b809450505050509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61507660e083018961502a565b8281036040840152615088818961502a565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506150b98185614f2b565b9a9950505050505050505050565b5f5f83601f8401126150d7575f5ffd5b5081356001600160401b038111156150ed575f5ffd5b6020830191508360208260051b8501011115614be1575f5ffd5b5f82601f830112615116575f5ffd5b81356151246149f7826149a3565b8082825260208201915060208360051b860101925085831115615145575f5ffd5b602085015b83811015614a415780356001600160401b03811115615167575f5ffd5b615176886020838a0101614a4b565b8452506020928301920161514a565b5f5f5f5f5f60808688031215615199575f5ffd5b85356001600160401b038111156151ae575f5ffd5b6151ba888289016150c7565b90965094505060208601356001600160401b038111156151d8575f5ffd5b8601601f810188136151e8575f5ffd5b80356151f66149f7826149a3565b8082825260208201915060208360051b85010192508a831115615217575f5ffd5b602084015b838110156152575780356001600160401b03811115615239575f5ffd5b6152488d6020838901016149da565b8452506020928301920161521c565b50955050505060408601356001600160401b03811115615275575f5ffd5b61528188828901615107565b92505060608601356001600160401b0381111561529c575f5ffd5b6152a888828901615107565b9150509295509295909350565b801515811461266f575f5ffd5b5f5f5f606084860312156152d4575f5ffd5b8335925060208401356152e681614908565b915060408401356152f6816152b5565b809150509250925092565b5f5f60408385031215615312575f5ffd5b82356001600160401b03811115615327575f5ffd5b61533385828601614a4b565b92505060208301356001600160401b0381111561534e575f5ffd5b614d6285828601614a4b565b602080825282518282018190525f918401906040840190835b81811015614e205783516001600160a01b0316835260209384019390920191600101615373565b602081525f611937602083018461502a565b634e487b7160e01b5f52602160045260245ffd5b600a81106153dc57634e487b7160e01b5f52602160045260245ffd5b9052565b602081525f825160806020840152805160a0840152602081015160c084015260018060a01b0360408201511660e084015260018060a01b03606082015116610100840152608081015161012084015260a0810151905060c061014084015261544c61016084018261502a565b9050602084015161546060408501826153c0565b50604084015160608401526060840151601f19848303016080850152611917828261502a565b5f60208284031215615496575f5ffd5b611937826149c5565b5f5f604083850312156154b0575f5ffd5b823591506020830135614b99816152b5565b5f602082840312156154d2575f5ffd5b8135611937816152b5565b5f5f5f5f604085870312156154f0575f5ffd5b84356001600160401b03811115615505575f5ffd5b615511878288016150c7565b90955093505060208501356001600160401b0381111561552f575f5ffd5b8501601f8101871361553f575f5ffd5b80356001600160401b03811115615554575f5ffd5b87602060e083028401011115615568575f5ffd5b949793965060200194505050565b5f5f5f60608486031215615588575f5ffd5b83359250602084013561559a81614908565b929592945050506040919091013590565b5f5f5f5f608085870312156155be575f5ffd5b8435935060208501356155d0816152b5565b925060408501356155e081614908565b915060608501356155f081614908565b939692955090935050565b5f5f5f6060848603121561560d575f5ffd5b833561561881614908565b9250615626602085016149c5565b915060408401356152f681614908565b5f5f8335601e1984360301811261564b575f5ffd5b8301803591506001600160401b03821115615664575f5ffd5b602001915036819003821315614be1575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b88815260208101889052604081018790526001600160a01b0386811660608301528516608082015260a0810184905260e060c082018190525f906150b99083018486615678565b6001600160a01b038516815260208101849052604081018390526080810161191760608301846153c0565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610a4e57610a4e615712565b600181811c9082168061574d57607f821691505b60208210810361576b57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b8481526001600160a01b03841660208201526080604082018190525f906157ae9083018561502a565b905060ff8316606083015295945050505050565b5f602082840312156157d2575f5ffd5b815161193781614908565b5f823560be198336030181126157f1575f5ffd5b9190910192915050565b5f823560de198336030181126157f1575f5ffd5b5f5f5f60608486031215615821575f5ffd5b5050815160208301516040909301519094929350919050565b5f6020828403121561584a575f5ffd5b8151600a8110611937575f5ffd5b5f60c08236031215615868575f5ffd5b61587061494b565b8235815260208084013590820152604083013561588c81614908565b6040820152606083013561589f81614908565b60608201526080838101359082015260a08301356001600160401b038111156158c6575f5ffd5b6158d236828601614d02565b60a08301525092915050565b601f821115610e9a57805f5260205f20601f840160051c810160208510156159035750805b601f840160051c820191505b8181101561208b575f815560010161590f565b81516001600160401b0381111561593b5761593b614937565b61594f816159498454615739565b846158de565b6020601f821160018114615981575f831561596a5750848201515b5f19600385901b1c1916600184901b17845561208b565b5f84815260208120601f198516915b828110156159b05787850151825560209485019460019092019101615990565b50848210156159cd57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b038b811682528a8116602083015289166040820152606081018890526080810187905260a0810186905260c0810185905283151560e082015261012061010082018190525f90615a369083018486615678565b9c9b505050505050505050505050565b615a5081846153c0565b604060208201525f610e1d604083018461502a565b5f60208284031215615a75575f5ffd5b5051919050565b5f82518060208501845e5f920191825250919050565b5f60208284031215615aa2575f5ffd5b8151611937816152b5565b81810381811115610a4e57610a4e615712565b634e487b7160e01b5f52603160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212203f383118670b0057138f9a7c76cc86594529c4f3cc7d2dda78102681499793e564736f6c634300081c0033",
}

// BaseBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseBridgeMetaData.ABI instead.
var BaseBridgeABI = BaseBridgeMetaData.ABI

// Deprecated: Use BaseBridgeMetaData.Sigs instead.
// BaseBridgeFuncSigs maps the 4-byte function signature to its string representation.
var BaseBridgeFuncSigs = BaseBridgeMetaData.Sigs

// BaseBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseBridgeMetaData.Bin instead.
var BaseBridgeBin = BaseBridgeMetaData.Bin

// DeployBaseBridge deploys a new Ethereum contract, binding an instance of BaseBridge to it.
func DeployBaseBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseBridge, error) {
	parsed, err := BaseBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseBridge{BaseBridgeCaller: BaseBridgeCaller{contract: contract}, BaseBridgeTransactor: BaseBridgeTransactor{contract: contract}, BaseBridgeFilterer: BaseBridgeFilterer{contract: contract}}, nil
}

// BaseBridge is an auto generated Go binding around an Ethereum contract.
type BaseBridge struct {
	BaseBridgeCaller     // Read-only binding to the contract
	BaseBridgeTransactor // Write-only binding to the contract
	BaseBridgeFilterer   // Log filterer for contract events
}

// BaseBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseBridgeSession struct {
	Contract     *BaseBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseBridgeCallerSession struct {
	Contract *BaseBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BaseBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseBridgeTransactorSession struct {
	Contract     *BaseBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BaseBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseBridgeRaw struct {
	Contract *BaseBridge // Generic contract binding to access the raw methods on
}

// BaseBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseBridgeCallerRaw struct {
	Contract *BaseBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BaseBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseBridgeTransactorRaw struct {
	Contract *BaseBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseBridge creates a new instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridge(address common.Address, backend bind.ContractBackend) (*BaseBridge, error) {
	contract, err := bindBaseBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseBridge{BaseBridgeCaller: BaseBridgeCaller{contract: contract}, BaseBridgeTransactor: BaseBridgeTransactor{contract: contract}, BaseBridgeFilterer: BaseBridgeFilterer{contract: contract}}, nil
}

// NewBaseBridgeCaller creates a new read-only instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeCaller(address common.Address, caller bind.ContractCaller) (*BaseBridgeCaller, error) {
	contract, err := bindBaseBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeCaller{contract: contract}, nil
}

// NewBaseBridgeTransactor creates a new write-only instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseBridgeTransactor, error) {
	contract, err := bindBaseBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTransactor{contract: contract}, nil
}

// NewBaseBridgeFilterer creates a new log filterer instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseBridgeFilterer, error) {
	contract, err := bindBaseBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeFilterer{contract: contract}, nil
}

// bindBaseBridge binds a generic wrapper to an already deployed contract.
func bindBaseBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBridge *BaseBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseBridge.Contract.BaseBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBridge *BaseBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.Contract.BaseBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBridge *BaseBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBridge.Contract.BaseBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBridge *BaseBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBridge *BaseBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBridge *BaseBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BaseBridge.Contract.DEFAULTADMINROLE(&_BaseBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BaseBridge.Contract.DEFAULTADMINROLE(&_BaseBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BaseBridge.Contract.UPGRADEINTERFACEVERSION(&_BaseBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BaseBridge.Contract.UPGRADEINTERFACEVERSION(&_BaseBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _BaseBridge.Contract.AllChainIDs(&_BaseBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BaseBridge.Contract.AllChainIDs(&_BaseBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BaseBridge.Contract.AllPendingIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BaseBridge.Contract.AllPendingIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256)[])
func (_BaseBridge *BaseBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256)[])
func (_BaseBridge *BaseBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256)[])
func (_BaseBridge *BaseBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeSession) BridgeVerifier() (common.Address, error) {
	return _BaseBridge.Contract.BridgeVerifier(&_BaseBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _BaseBridge.Contract.BridgeVerifier(&_BaseBridge.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BaseBridge *BaseBridgeCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumValue *big.Int
		GasFee       *big.Int
		ExFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BaseBridge *BaseBridgeSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BaseBridge.Contract.CalculateFee(&_BaseBridge.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BaseBridge *BaseBridgeCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BaseBridge.Contract.CalculateFee(&_BaseBridge.CallOpts, remoteChainID, token, value)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _BaseBridge.Contract.CrossMintableERC20Code(&_BaseBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _BaseBridge.Contract.CrossMintableERC20Code(&_BaseBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeSession) Dev() (common.Address, error) {
	return _BaseBridge.Contract.Dev(&_BaseBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) Dev() (common.Address, error) {
	return _BaseBridge.Contract.Dev(&_BaseBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) DomainSeparator() ([32]byte, error) {
	return _BaseBridge.Contract.DomainSeparator(&_BaseBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _BaseBridge.Contract.DomainSeparator(&_BaseBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BaseBridge *BaseBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_BaseBridge *BaseBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BaseBridge.Contract.Eip712Domain(&_BaseBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BaseBridge *BaseBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BaseBridge.Contract.Eip712Domain(&_BaseBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextFinalizeIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextFinalizeIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextInitiateIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextInitiateIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256,bytes))
func (_BaseBridge *BaseBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256,bytes))
func (_BaseBridge *BaseBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256,bytes))
func (_BaseBridge *BaseBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BaseBridge.Contract.GetRoleAdmin(&_BaseBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BaseBridge.Contract.GetRoleAdmin(&_BaseBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BaseBridge.Contract.GetRoleMembers(&_BaseBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BaseBridge.Contract.GetRoleMembers(&_BaseBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256))
func (_BaseBridge *BaseBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256))
func (_BaseBridge *BaseBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256))
func (_BaseBridge *BaseBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BaseBridge.Contract.HasRole(&_BaseBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BaseBridge.Contract.HasRole(&_BaseBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeSession) InitializedAt() (*big.Int, error) {
	return _BaseBridge.Contract.InitializedAt(&_BaseBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _BaseBridge.Contract.InitializedAt(&_BaseBridge.CallOpts)
}

// IsPending is a free data retrieval call binding the contract method 0x8bb19439.
//
// Solidity: function isPending(uint256 remoteChainID, uint256 index) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) IsPending(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "isPending", remoteChainID, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPending is a free data retrieval call binding the contract method 0x8bb19439.
//
// Solidity: function isPending(uint256 remoteChainID, uint256 index) view returns(bool)
func (_BaseBridge *BaseBridgeSession) IsPending(remoteChainID *big.Int, index *big.Int) (bool, error) {
	return _BaseBridge.Contract.IsPending(&_BaseBridge.CallOpts, remoteChainID, index)
}

// IsPending is a free data retrieval call binding the contract method 0x8bb19439.
//
// Solidity: function isPending(uint256 remoteChainID, uint256 index) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) IsPending(remoteChainID *big.Int, index *big.Int) (bool, error) {
	return _BaseBridge.Contract.IsPending(&_BaseBridge.CallOpts, remoteChainID, index)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeSession) Paused() (bool, error) {
	return _BaseBridge.Contract.Paused(&_BaseBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) Paused() (bool, error) {
	return _BaseBridge.Contract.Paused(&_BaseBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _BaseBridge.Contract.ProxiableUUID(&_BaseBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BaseBridge.Contract.ProxiableUUID(&_BaseBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseBridge.Contract.SupportsInterface(&_BaseBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseBridge.Contract.SupportsInterface(&_BaseBridge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeSession) Threshold() (uint8, error) {
	return _BaseBridge.Contract.Threshold(&_BaseBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeCallerSession) Threshold() (uint8, error) {
	return _BaseBridge.Contract.Threshold(&_BaseBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.BridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.BridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.ChangeThreshold(&_BaseBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.ChangeThreshold(&_BaseBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridge(&_BaseBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridge(&_BaseBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridgeBatch(&_BaseBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridgeBatch(&_BaseBridge.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRole(&_BaseBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRole(&_BaseBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "grantRoleBatch", role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRoleBatch(&_BaseBridge.TransactOpts, role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactorSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRoleBatch(&_BaseBridge.TransactOpts, role, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf35f9e45.
//
// Solidity: function initialize(address owner_, uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, _threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "initialize", owner_, _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf35f9e45.
//
// Solidity: function initialize(address owner_, uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeSession) Initialize(owner_ common.Address, _threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, owner_, _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf35f9e45.
//
// Solidity: function initialize(address owner_, uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) Initialize(owner_ common.Address, _threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, owner_, _threshold, dev_)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeTransactor) ManualProcessPending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "manualProcessPending", remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualProcessPending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ManualProcessPending is a paid mutator transaction binding the contract method 0x7f4ab9f5.
//
// Solidity: function manualProcessPending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) ManualProcessPending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualProcessPending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, gasFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeTokenBatch(&_BaseBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0xd605665b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeTokenBatch(&_BaseBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeTransactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "releasePendingBatch", remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePendingBatch(&_BaseBridge.TransactOpts, remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePendingBatch(&_BaseBridge.TransactOpts, remoteChainIDs, indexes)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceRole(&_BaseBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceRole(&_BaseBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRole(&_BaseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRole(&_BaseBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "revokeRoleBatch", role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRoleBatch(&_BaseBridge.TransactOpts, role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRoleBatch(&_BaseBridge.TransactOpts, role, accounts)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeVerifier(&_BaseBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeVerifier(&_BaseBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetChainPause(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetChainPause(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetCrossMintableERC20Code(&_BaseBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetCrossMintableERC20Code(&_BaseBridge.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetDev(&_BaseBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetDev(&_BaseBridge.TransactOpts, dev_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPause(&_BaseBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPause(&_BaseBridge.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetTokenPause(&_BaseBridge.TransactOpts, remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetTokenPause(&_BaseBridge.TransactOpts, remoteChainID, token, pause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelay(&_BaseBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelay(&_BaseBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelayExpiration(&_BaseBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelayExpiration(&_BaseBridge.TransactOpts, remoteChainID, index, delay)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BaseBridge *BaseBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BaseBridge *BaseBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.UnregisterToken(&_BaseBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BaseBridge *BaseBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.UnregisterToken(&_BaseBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.UpgradeToAndCall(&_BaseBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.UpgradeToAndCall(&_BaseBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeSession) Receive() (*types.Transaction, error) {
	return _BaseBridge.Contract.Receive(&_BaseBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BaseBridge.Contract.Receive(&_BaseBridge.TransactOpts)
}

// BaseBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BaseBridge contract.
type BaseBridgeBridgeFinalizedIterator struct {
	Event *BaseBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeFinalized)
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
		it.Event = new(BaseBridgeBridgeFinalized)
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
func (it *BaseBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeFinalized represents a BridgeFinalized event raised by the BaseBridge contract.
type BaseBridgeBridgeFinalized struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BaseBridgeBridgeFinalizedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeFinalizedIterator{contract: _BaseBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeFinalized)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeFinalized(log types.Log) (*BaseBridgeBridgeFinalized, error) {
	event := new(BaseBridgeBridgeFinalized)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BaseBridge contract.
type BaseBridgeBridgeInitiatedIterator struct {
	Event *BaseBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeInitiated)
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
		it.Event = new(BaseBridgeBridgeInitiated)
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
func (it *BaseBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeInitiated represents a BridgeInitiated event raised by the BaseBridge contract.
type BaseBridgeBridgeInitiated struct {
	ToChainID *big.Int
	Index     *big.Int
	FromToken common.Address
	ToToken   common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	GasFee    *big.Int
	ExFee     *big.Int
	Time      *big.Int
	Permit    bool
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*BaseBridgeBridgeInitiatedIterator, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeInitiatedIterator{contract: _BaseBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeInitiated)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db50.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 exFee, uint256 time, bool permit, bytes extraData)
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeInitiated(log types.Log) (*BaseBridgeBridgeInitiated, error) {
	event := new(BaseBridgeBridgeInitiated)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the BaseBridge contract.
type BaseBridgeBridgePendingIterator struct {
	Event *BaseBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgePending)
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
		it.Event = new(BaseBridgeBridgePending)
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
func (it *BaseBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgePending represents a BridgePending event raised by the BaseBridge contract.
type BaseBridgeBridgePending struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgePending is a free log retrieval operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BaseBridgeBridgePendingIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgePendingIterator{contract: _BaseBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgePending)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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

// ParseBridgePending is a log parse operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BaseBridge *BaseBridgeFilterer) ParseBridgePending(log types.Log) (*BaseBridgeBridgePending, error) {
	event := new(BaseBridgeBridgePending)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the BaseBridge contract.
type BaseBridgeBridgeVerifierSetIterator struct {
	Event *BaseBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeVerifierSet)
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
		it.Event = new(BaseBridgeBridgeVerifierSet)
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
func (it *BaseBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the BaseBridge contract.
type BaseBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*BaseBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeVerifierSetIterator{contract: _BaseBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeVerifierSet)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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

// ParseBridgeVerifierSet is a log parse operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*BaseBridgeBridgeVerifierSet, error) {
	event := new(BaseBridgeBridgeVerifierSet)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the BaseBridge contract.
type BaseBridgeChainPauseSetIterator struct {
	Event *BaseBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeChainPauseSet)
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
		it.Event = new(BaseBridgeChainPauseSet)
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
func (it *BaseBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeChainPauseSet represents a ChainPauseSet event raised by the BaseBridge contract.
type BaseBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BaseBridge *BaseBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BaseBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeChainPauseSetIterator{contract: _BaseBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BaseBridge *BaseBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeChainPauseSet)
				if err := _BaseBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseChainPauseSet(log types.Log) (*BaseBridgeChainPauseSet, error) {
	event := new(BaseBridgeChainPauseSet)
	if err := _BaseBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the BaseBridge contract.
type BaseBridgeCrossMintableERC20CodeSetIterator struct {
	Event *BaseBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(BaseBridgeCrossMintableERC20CodeSet)
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
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the BaseBridge contract.
type BaseBridgeCrossMintableERC20CodeSet struct {
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BaseBridge *BaseBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*BaseBridgeCrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeCrossMintableERC20CodeSetIterator{contract: _BaseBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BaseBridge *BaseBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeCrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeCrossMintableERC20CodeSet)
				if err := _BaseBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BaseBridge *BaseBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*BaseBridgeCrossMintableERC20CodeSet, error) {
	event := new(BaseBridgeCrossMintableERC20CodeSet)
	if err := _BaseBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BaseBridge contract.
type BaseBridgeEIP712DomainChangedIterator struct {
	Event *BaseBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeEIP712DomainChanged)
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
		it.Event = new(BaseBridgeEIP712DomainChanged)
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
func (it *BaseBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the BaseBridge contract.
type BaseBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BaseBridge *BaseBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BaseBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeEIP712DomainChangedIterator{contract: _BaseBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BaseBridge *BaseBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeEIP712DomainChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BaseBridgeEIP712DomainChanged, error) {
	event := new(BaseBridgeEIP712DomainChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BaseBridge contract.
type BaseBridgeInitializedIterator struct {
	Event *BaseBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BaseBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeInitialized)
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
		it.Event = new(BaseBridgeInitialized)
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
func (it *BaseBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeInitialized represents a Initialized event raised by the BaseBridge contract.
type BaseBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BaseBridge *BaseBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BaseBridgeInitializedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeInitializedIterator{contract: _BaseBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BaseBridge *BaseBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BaseBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeInitialized)
				if err := _BaseBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseInitialized(log types.Log) (*BaseBridgeInitialized, error) {
	event := new(BaseBridgeInitialized)
	if err := _BaseBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BaseBridge contract.
type BaseBridgePausedIterator struct {
	Event *BaseBridgePaused // Event containing the contract specifics and raw log

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
func (it *BaseBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgePaused)
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
		it.Event = new(BaseBridgePaused)
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
func (it *BaseBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgePaused represents a Paused event raised by the BaseBridge contract.
type BaseBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseBridge *BaseBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BaseBridgePausedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BaseBridgePausedIterator{contract: _BaseBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseBridge *BaseBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BaseBridgePaused) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgePaused)
				if err := _BaseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParsePaused(log types.Log) (*BaseBridgePaused, error) {
	event := new(BaseBridgePaused)
	if err := _BaseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BaseBridge contract.
type BaseBridgeRoleAdminChangedIterator struct {
	Event *BaseBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleAdminChanged)
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
		it.Event = new(BaseBridgeRoleAdminChanged)
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
func (it *BaseBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the BaseBridge contract.
type BaseBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BaseBridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleAdminChangedIterator{contract: _BaseBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleAdminChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BaseBridge *BaseBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BaseBridgeRoleAdminChanged, error) {
	event := new(BaseBridgeRoleAdminChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BaseBridge contract.
type BaseBridgeRoleGrantedIterator struct {
	Event *BaseBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleGranted)
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
		it.Event = new(BaseBridgeRoleGranted)
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
func (it *BaseBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleGranted represents a RoleGranted event raised by the BaseBridge contract.
type BaseBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BaseBridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleGrantedIterator{contract: _BaseBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleGranted)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) ParseRoleGranted(log types.Log) (*BaseBridgeRoleGranted, error) {
	event := new(BaseBridgeRoleGranted)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BaseBridge contract.
type BaseBridgeRoleRevokedIterator struct {
	Event *BaseBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleRevoked)
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
		it.Event = new(BaseBridgeRoleRevoked)
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
func (it *BaseBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleRevoked represents a RoleRevoked event raised by the BaseBridge contract.
type BaseBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BaseBridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleRevokedIterator{contract: _BaseBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleRevoked)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) ParseRoleRevoked(log types.Log) (*BaseBridgeRoleRevoked, error) {
	event := new(BaseBridgeRoleRevoked)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BaseBridge contract.
type BaseBridgeThresholdChangedIterator struct {
	Event *BaseBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeThresholdChanged)
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
		it.Event = new(BaseBridgeThresholdChanged)
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
func (it *BaseBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeThresholdChanged represents a ThresholdChanged event raised by the BaseBridge contract.
type BaseBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BaseBridge *BaseBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BaseBridgeThresholdChangedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeThresholdChangedIterator{contract: _BaseBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BaseBridge *BaseBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeThresholdChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseThresholdChanged(log types.Log) (*BaseBridgeThresholdChanged, error) {
	event := new(BaseBridgeThresholdChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BaseBridge contract.
type BaseBridgeTokenPairRegisteredIterator struct {
	Event *BaseBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *BaseBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeTokenPairRegistered)
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
		it.Event = new(BaseBridgeTokenPairRegistered)
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
func (it *BaseBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the BaseBridge contract.
type BaseBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BaseBridge *BaseBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BaseBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTokenPairRegisteredIterator{contract: _BaseBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BaseBridge *BaseBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BaseBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeTokenPairRegistered)
				if err := _BaseBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BaseBridge *BaseBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*BaseBridgeTokenPairRegistered, error) {
	event := new(BaseBridgeTokenPairRegistered)
	if err := _BaseBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the BaseBridge contract.
type BaseBridgeTokenPairUnregisteredIterator struct {
	Event *BaseBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *BaseBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeTokenPairUnregistered)
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
		it.Event = new(BaseBridgeTokenPairUnregistered)
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
func (it *BaseBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the BaseBridge contract.
type BaseBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BaseBridge *BaseBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*BaseBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTokenPairUnregisteredIterator{contract: _BaseBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BaseBridge *BaseBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *BaseBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeTokenPairUnregistered)
				if err := _BaseBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*BaseBridgeTokenPairUnregistered, error) {
	event := new(BaseBridgeTokenPairUnregistered)
	if err := _BaseBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the BaseBridge contract.
type BaseBridgeTokenPauseSetIterator struct {
	Event *BaseBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeTokenPauseSet)
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
		it.Event = new(BaseBridgeTokenPauseSet)
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
func (it *BaseBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeTokenPauseSet represents a TokenPauseSet event raised by the BaseBridge contract.
type BaseBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_BaseBridge *BaseBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BaseBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTokenPauseSetIterator{contract: _BaseBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_BaseBridge *BaseBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeTokenPauseSet)
				if err := _BaseBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseTokenPauseSet(log types.Log) (*BaseBridgeTokenPauseSet, error) {
	event := new(BaseBridgeTokenPauseSet)
	if err := _BaseBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BaseBridge contract.
type BaseBridgeUnpausedIterator struct {
	Event *BaseBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BaseBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeUnpaused)
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
		it.Event = new(BaseBridgeUnpaused)
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
func (it *BaseBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeUnpaused represents a Unpaused event raised by the BaseBridge contract.
type BaseBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseBridge *BaseBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BaseBridgeUnpausedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeUnpausedIterator{contract: _BaseBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseBridge *BaseBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BaseBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeUnpaused)
				if err := _BaseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseUnpaused(log types.Log) (*BaseBridgeUnpaused, error) {
	event := new(BaseBridgeUnpaused)
	if err := _BaseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BaseBridge contract.
type BaseBridgeUpgradedIterator struct {
	Event *BaseBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BaseBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeUpgraded)
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
		it.Event = new(BaseBridgeUpgraded)
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
func (it *BaseBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeUpgraded represents a Upgraded event raised by the BaseBridge contract.
type BaseBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseBridge *BaseBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BaseBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeUpgradedIterator{contract: _BaseBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseBridge *BaseBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BaseBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeUpgraded)
				if err := _BaseBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseUpgraded(log types.Log) (*BaseBridgeUpgraded, error) {
	event := new(BaseBridgeUpgraded)
	if err := _BaseBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the BaseBridge contract.
type BaseBridgeVerificationDelayExpirationSetIterator struct {
	Event *BaseBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeVerificationDelayExpirationSet)
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
		it.Event = new(BaseBridgeVerificationDelayExpirationSet)
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
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the BaseBridge contract.
type BaseBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BaseBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeVerificationDelayExpirationSetIterator{contract: _BaseBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeVerificationDelayExpirationSet)
				if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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

// ParseVerificationDelayExpirationSet is a log parse operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*BaseBridgeVerificationDelayExpirationSet, error) {
	event := new(BaseBridgeVerificationDelayExpirationSet)
	if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the BaseBridge contract.
type BaseBridgeVerificationDelaySetIterator struct {
	Event *BaseBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeVerificationDelaySet)
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
		it.Event = new(BaseBridgeVerificationDelaySet)
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
func (it *BaseBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the BaseBridge contract.
type BaseBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*BaseBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeVerificationDelaySetIterator{contract: _BaseBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *BaseBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeVerificationDelaySet)
				if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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

// ParseVerificationDelaySet is a log parse operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*BaseBridgeVerificationDelaySet, error) {
	event := new(BaseBridgeVerificationDelaySet)
	if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
