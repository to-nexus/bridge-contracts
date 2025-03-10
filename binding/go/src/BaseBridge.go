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
	Args         IBridgeRegistryFinalizeArguments
	SafeDeadline *big.Int
	Reason       []byte
}

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRegistryTokenPair struct {
	LocalToken                  common.Address
	RemoteToken                 common.Address
	Paused                      bool
	IsOrigin                    bool
	VerificationAmountThreshold *big.Int
	Deposited                   *big.Int
	PendingAmount               *big.Int
}

// BaseBridgeMetaData contains all meta data concerning the BaseBridge contract.
var BaseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"clearPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasExpiredPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualProcessPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"releaseExpiredPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"resetRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"_bridgeFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setSafetyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"setVerificationDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"SafetyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"conversionRatio\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"VerificationDeadlineSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"safeDeadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"name\":\"RegistryBalanceLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPauseNotChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleAlreadyGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"25bd5666": "bridgeFeeManager()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"0b43c02c": "clearPending(uint256)",
		"d016d625": "createToken(uint256,address,int256,uint256,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
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
		"7f4ab9f5": "manualProcessPending(uint256,uint256)",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"d605665b": "permitBridgeTokenBatch((uint256,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"1e7bf215": "registerToken(uint256,bool,address,address,int256,uint256)",
		"c1cb05dd": "releaseExpiredPending(uint256)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"6b0e4821": "releasePendingBatch(uint256,uint256[])",
		"715018a6": "renounceOwnership()",
		"2d87b7ee": "resetRole(bytes32,address[])",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"472d35b9": "setFeeManager(address)",
		"bedb86fb": "setPause(bool)",
		"6160751f": "setPauseChain(uint256,bool)",
		"4d3f0da9": "setPauseToken(uint256,address,bool)",
		"d4bf502a": "setRole(bytes32,address[],bool)",
		"39a621f3": "setSafetyLimit(uint256,address,uint256)",
		"3ca5b39f": "setVerificationDeadline(uint256,uint256,uint256)",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615faf6100f95f395f81816137e60152818161380f01526139d60152615faf5ff3fe6080604052600436106102c2575f3560e01c8063814914b51161016f578063bedb86fb116100d8578063d5717fc511610092578063e2af6cd71161006d578063e2af6cd714610973578063f2fde38b14610992578063f4509643146109b1578063f698da25146109d0575f5ffd5b8063d5717fc514610907578063d605665b14610926578063df6e87dc14610939575f5ffd5b8063bedb86fb14610858578063c1cb05dd14610877578063cf56118e14610896578063d016d625146108aa578063d477f05f146108c9578063d4bf502a146108e8575f5ffd5b806391d148541161012957806391d1485414610766578063a3246ad314610785578063ad3cb1cc146107b1578063ae6893f8146107ee578063b33eb36e1461080d578063b7f3358d14610839575f5ffd5b8063814914b51461060157806384b0196e146106e757806388d67d6d1461070e5780638da5cb5b1461072157806391cca3db1461073557806391cf6d3e14610752575f5ffd5b80634d5d00561161022b5780635c975abb116101e55780636b0e4821116101c05780636b0e482114610583578063715018a6146105a257806379214874146105b65780637f4ab9f5146105e2575f5ffd5b80635c975abb1461052e5780635fd262de146105515780636160751f14610564575f5ffd5b80634d5d00561461047c5780634ee078ba1461048f5780634f1ef286146104ae5780635187599d146104c157806352d1902d146104e05780635b605f5c14610502575f5ffd5b80633ca5b39f1161027c5780633ca5b39f146103c85780633d507c5e146103e757806342cde4e8146103fe578063472d35b91461041f5780634be13f831461043e5780634d3f0da91461045d575f5ffd5b80630b43c02c146102ed5780631938e0f21461030c5780631e7bf2151461033457806325bd5666146103535780632d87b7ee1461038a57806339a621f3146103a9575f5ffd5b366102e957345f036102e7576040516304a4cdd960e51b815260040160405180910390fd5b005b5f5ffd5b3480156102f8575f5ffd5b506102e7610307366004614c9f565b6109e4565b61031f61031a366004614e25565b610a40565b60405190151581526020015b60405180910390f35b34801561033f575f5ffd5b506102e761034e366004614eff565b610e5b565b34801561035e575f5ffd5b50609654610372906001600160a01b031681565b6040516001600160a01b03909116815260200161032b565b348015610395575f5ffd5b506102e76103a4366004614fc4565b61110b565b3480156103b4575f5ffd5b506102e76103c3366004615007565b61113d565b3480156103d3575f5ffd5b506102e76103e236600461503c565b611230565b3480156103f2575f5ffd5b5060335442111561031f565b348015610409575f5ffd5b5060645460405160ff909116815260200161032b565b34801561042a575f5ffd5b506102e7610439366004615065565b611309565b348015610449575f5ffd5b50603254610372906001600160a01b031681565b348015610468575f5ffd5b506102e7610477366004615080565b61137a565b61031f61048a366004615103565b61147e565b34801561049a575f5ffd5b5061031f6104a93660046151a5565b6116a0565b6102e76104bc36600461523b565b6119ab565b3480156104cc575f5ffd5b506102e76104db36600461527d565b6119c6565b3480156104eb575f5ffd5b506104f4611ad4565b60405190815260200161032b565b34801561050d575f5ffd5b5061052161051c366004614c9f565b611aef565b60405161032b9190615305565b348015610539575f5ffd5b505f516020615f3a5f395f51905f525460ff1661031f565b61031f61055f366004615352565b611c74565b34801561056f575f5ffd5b506102e761057e3660046153da565b611cdd565b34801561058e575f5ffd5b5061031f61059d3660046153fd565b611db4565b3480156105ad575f5ffd5b506102e7611df6565b3480156105c1575f5ffd5b506105d56105d0366004614c9f565b611e09565b60405161032b91906154d3565b3480156105ed575f5ffd5b5061031f6105fc3660046151a5565b611e22565b34801561060c575f5ffd5b506106da61061b3660046154e5565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526039602090815260408084206001600160a01b039384168552825292839020835160e08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b60405161032b9190615508565b3480156106f2575f5ffd5b506106fb611ebd565b60405161032b9796959493929190615544565b61031f61071c366004615671565b611f66565b34801561072c575f5ffd5b50610372612001565b348015610740575f5ffd5b506097546001600160a01b0316610372565b34801561075d575f5ffd5b506098546104f4565b348015610771575f5ffd5b5061031f6107803660046154e5565b61202f565b348015610790575f5ffd5b506107a461079f366004614c9f565b612046565b60405161032b91906157a1565b3480156107bc575f5ffd5b506107e1604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161032b91906157e1565b3480156107f9575f5ffd5b506104f4610808366004614c9f565b61205f565b348015610818575f5ffd5b5061082c6108273660046151a5565b61207b565b60405161032b91906157f3565b348015610844575f5ffd5b506102e7610853366004615885565b612225565b348015610863575f5ffd5b506102e761087236600461589e565b612276565b348015610882575f5ffd5b506102e7610891366004614c9f565b612297565b3480156108a1575f5ffd5b506105d561257f565b3480156108b5575f5ffd5b506103726108c43660046158b9565b612590565b3480156108d4575f5ffd5b506102e76108e3366004615065565b61264b565b3480156108f3575f5ffd5b506102e7610902366004615942565b61269c565b348015610912575f5ffd5b506104f4610921366004614c9f565b6126d3565b6102e761093436600461598d565b6126ef565b348015610944575f5ffd5b50610958610953366004615007565b612882565b6040805193845260208401929092529082015260600161032b565b34801561097e575f5ffd5b506102e761098d366004615065565b612911565b34801561099d575f5ffd5b506102e76109ac366004615065565b6129c1565b3480156109bc575f5ffd5b506102e76109cb3660046154e5565b6129fb565b3480156109db575f5ffd5b506104f4612ae0565b6109ec612ae9565b5f818152603b60205260408120610a0290612b1b565b90505f5b8151811015610a3b57610a3283838381518110610a2557610a25615a26565b6020026020010151612b27565b50600101610a06565b505050565b5f610a49612d34565b610a51612d64565b610a79610a646060870160408801615065565b86355f90815260386020526040902090612d9b565b610a896060870160408801615065565b90610ab857604051633142cb1160e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f348015610ae3576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505084355f9081526037602052604081206002018054600101908190559050806020870135808214610b315760405163a6ab65c560e01b815260048101929092526024820152604401610aaf565b50610bc390507fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219146020880135610b6d60608a0160408b01615065565b610b7d60808b0160608c01615065565b60808b0135610b8f60a08d018d615a3a565b604051602001610ba59796959493929190615aa4565b60405160208183030381529060405280519060200120868686612dbc565b85355f90815260396020526040808220606091839182918290610beb908d8701908e01615065565b6001600160a01b03908116825260208083019390935260409182015f20825160e08101845281548316815260018201549283169481019490945260ff600160a01b8304811615801594860194909452600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c0830152909150610c9c576040518060400160405280600c81526020016b1d1bdad95b881c185d5cd95960a21b8152509250610d25565b608081015115801590610cb6575089608001358160800151105b15610cf057604051806040016040528060118152602001701bdd995c881cd859995d1e481b1a5b5a5d607a1b815250925060019150610d25565b610d1f8a35610d0560608d0160408e01615065565b610d1560808e0160608f01615065565b8d6080013561300c565b90945092505b508215610dad57610d3c60608a0160408b01615065565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610d7d60808e0160608f01615065565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610e35565b610db88983836130fa565b610dc860608a0160408b01615065565b6001600160a01b031660208a01358a357fd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7610e0960808e0160608f01615065565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a45b6001945050505050610e5360015f516020615f5a5f395f51905f5255565b949350505050565b610e63612ae9565b610e6e6035876132ba565b15610ecb57604080516080810182528781525f6020808301828152838501838152606085018481528c855260379093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b6001600160a01b038416610ef257604051636ca1fdd760e01b815260040160405180910390fd5b5f868152603860205260409020610f0990856132c5565b8490610f34576040516311dd05f360e31b81526001600160a01b039091166004820152602401610aaf565b506040518060e00160405280856001600160a01b03168152602001846001600160a01b031681526020015f1515815260200186151581526020018281526020015f81526020015f81525060395f8881526020019081526020015f205f866001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c08201518160040155905050600182138061107657505f1982125b156110a0575f868152603a602090815260408083206001600160a01b038816845290915290208290555b826001600160a01b0316846001600160a01b0316877fffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a6685858a6040516110fb9392919092835260208301919091521515604082015260600190565b60405180910390a4505050505050565b61112d826111275f5f8681526020019081526020015f20612b1b565b5f61269c565b6111398282600161269c565b5050565b61114f6420a226a4a760d91b3361202f565b6420a226a4a760d91b33909161118a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206111a39083612d9b565b82906111ce5760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b038616808552908352928190206002018490555183815285917fd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c891015b60405180910390a3505050565b6112426420a226a4a760d91b3361202f565b6420a226a4a760d91b33909161127d5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b6020526040902061129690836132d9565b82906112b8576040516373a1390160e11b8152600401610aaf91815260200190565b505f838152603c602090815260408083208584528252918290206006018390559051828152839185917ffcfd532dc7387ad29a63519b09004fa6a982a1853a71522d0d5fe94f3ef76d0d9101611223565b61131b6420a226a4a760d91b3361202f565b6420a226a4a760d91b3390916113565760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b5050609680546001600160a01b0319166001600160a01b0392909216919091179055565b61138c6420a226a4a760d91b3361202f565b6420a226a4a760d91b3390916113c75760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f8381526038602052604090206113e09083612d9b565b829061140b5760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8381526039602090815260408083206001600160a01b0386168085529252918290206001018054841515600160a01b0260ff60a01b19909116179055905184907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061122390851515815260200190565b5f611487612d34565b898961149382826132f0565b61149b612d64565b6114a86020850185615065565b6001600160a01b038c81169116148b6114c46020870187615065565b90916114f6576040516313f7c32b60e31b81526001600160a01b03928316600482015291166024820152604401610aaf565b50506115058c8c8b8b8b6133c6565b909850965086611515898b615b05565b61151f9190615b05565b60408501351015876115318a8c615b05565b61153b9190615b05565b856040013590916115685760405163943892eb60e01b815260048101929092526024820152604401610aaf565b505f905061157c6040860160208701615065565b306040870135606088013561159760a08a0160808b01615885565b6040516001600160a01b0395861660248201529490931660448501526064840191909152608483015260ff1660a482015260a086013560c482015260c086013560e48201526101040160408051601f19818403018152919052602080820180516001600160e01b031663d505accf60e01b1790529091505f9081906116299061162290890189615065565b5f856134aa565b9150915081819061164e5760405163e8b5c4bb60e01b8152600401610aaf91906157e1565b505050506116778c8c86602001602081019061166a9190615065565b8d8d8d8d60018e8e613627565b6001925061169160015f516020615f5a5f395f51905f5255565b50509998505050505050505050565b5f6116a9612d34565b6116b1612d64565b5f838152603b602052604090206116c890836132d9565b82906116ea576040516373a1390160e11b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e084015260058201805484929184916101008501919061176690615b18565b80601f016020809104026020016040519081016040528092919081815260200182805461179290615b18565b80156117dd5780601f106117b4576101008083540402835291602001916117dd565b820191905f5260205f20905b8154815290600101906020018083116117c057829003601f168201915b50505050508152505081526020016006820154815260200160078201805461180490615b18565b80601f016020809104026020016040519081016040528092919081815260200182805461183090615b18565b801561187b5780601f106118525761010080835404028352916020019161187b565b820191905f5260205f20905b81548152906001019060200180831161185e57829003601f168201915b505050919092525050505f85815260396020908152604080832084518201516001600160a01b03908116855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b8504811615801585850152600160a81b909504161515606084015260028101546080840152600381015460a08401526004015460c083015283510151929350919061193c576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b50602082015115806119515750428260200151105b826020015142909161197f57604051637a88099160e11b815260048101929092526024820152604401610aaf565b505061198b8585613709565b925050506119a560015f516020615f5a5f395f51905f5255565b92915050565b6119b36137db565b6119bc8261387f565b6111398282613887565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611a0a5750825b90505f826001600160401b03166001148015611a255750303b155b905081158015611a33575080155b15611a515760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611a7b57845460ff60401b1916600160401b1785555b611a858787613943565b8315611acb57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b5f611add6139cb565b505f516020615f1a5f395f51905f5290565b5f81815260386020526040812060609190611b0990612b1b565b90505f81516001600160401b03811115611b2557611b25614cb6565b604051908082528060200260200182016040528015611b8a57816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f19909201910181611b435790505b5090505f5b8251811015611c6c5760395f8681526020019081526020015f205f848381518110611bbc57611bbc615a26565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c08201528251839083908110611c5957611c59615a26565b6020908102919091010152600101611b8f565b509392505050565b5f611c7d612d34565b8888611c8982826132f0565b611c91612d64565b611c9e8b8b8a8a8a6133c6565b9097509550611cb58b8b338c8c8c8c5f8d8d613627565b60019250611ccf60015f516020615f5a5f395f51905f5255565b505098975050505050505050565b611cef6420a226a4a760d91b3361202f565b6420a226a4a760d91b339091611d2a5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50611d3890506035836132d9565b8290611d5a5760405163ac7dbbfd60e01b8152600401610aaf91815260200190565b505f82815260376020908152604091829020600301805460ff1916841515908117909155915191825283917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a25050565b5f805b8251811015611dec57611de384848381518110611dd657611dd6615a26565b60200260200101516116a0565b50600101611db7565b5060019392505050565b611dfe612ae9565b611e075f613a14565b565b5f818152603b602052604090206060906119a590612b1b565b5f611e356420a226a4a760d91b3361202f565b6420a226a4a760d91b339091611e705760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b50505f838152603b60205260409020611e8990836132d9565b8290611eab576040516373a1390160e11b8152600401610aaf91815260200190565b50611eb68383613709565b9392505050565b5f60608082808083815f516020615efa5f395f51905f528054909150158015611ee857506001810154155b611f2c5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610aaf565b611f34613a84565b611f3c613b44565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f805b85811015611ff457611feb878783818110611f8657611f86615a26565b9050602002810190611f989190615b50565b868381518110611faa57611faa615a26565b6020026020010151868481518110611fc457611fc4615a26565b6020026020010151868581518110611fde57611fde615a26565b6020026020010151610a40565b50600101611f69565b5060019695505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f828152602081905260408120611eb69083612d9b565b5f8181526020819052604090206060906119a590612b1b565b5f8181526037602052604081206001908101546119a591615b05565b612083614bec565b5f838152603c6020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e08401526005820180548492918491610100850191906120ff90615b18565b80601f016020809104026020016040519081016040528092919081815260200182805461212b90615b18565b80156121765780601f1061214d57610100808354040283529160200191612176565b820191905f5260205f20905b81548152906001019060200180831161215957829003601f168201915b50505050508152505081526020016006820154815260200160078201805461219d90615b18565b80601f01602080910402602001604051908101604052809291908181526020018280546121c990615b18565b80156122145780601f106121eb57610100808354040283529160200191612214565b820191905f5260205f20905b8154815290600101906020018083116121f757829003601f168201915b505050505081525050905092915050565b61222d612ae9565b6064805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b61227e612ae9565b801561228f5761228c613b82565b50565b61228c613bde565b4260335410156122a45750565b5f806122b06035612b1b565b90505f5b8151811015612579575f8282815181106122d0576122d0615a26565b6020908102919091018101515f818152603790925260409091206003015490915060ff16156122ff5750612571565b5f818152603b6020526040812061231590612b1b565b90505f5b815181101561256d575f82828151811061233557612335615a26565b6020908102919091018101515f868152603c83526040808220838352909352828120835161012081019094528054606085019081526001820154608086015260028201546001600160a01b0390811660a087015260038301541660c0860152600482015460e08601526005820180549496509294939192849284916101008501916123bf90615b18565b80601f01602080910402602001604051908101604052809291908181526020018280546123eb90615b18565b80156124365780601f1061240d57610100808354040283529160200191612436565b820191905f5260205f20905b81548152906001019060200180831161241957829003601f168201915b50505050508152505081526020016006820154815260200160078201805461245d90615b18565b80601f016020809104026020016040519081016040528092919081815260200182805461248990615b18565b80156124d45780601f106124ab576101008083540402835291602001916124d4565b820191905f5260205f20905b8154815290600101906020018083116124b757829003601f168201915b5050509190925250508151515f90815260396020908152604080832085518201516001600160a01b0316845290915290206001015491925050600160a01b900460ff161580156125275750602081015115155b80156125365750428160200151105b15612563576125458583613709565b508861255089615b6e565b9850881061256357505050505050505050565b5050600101612319565b5050505b6001016122b4565b50505050565b606061258b6035612b1b565b905090565b6032545f906001600160a01b03166125bb576040516315aeca0d60e11b815260040160405180910390fd5b603254604051637c469ea160e11b81526001600160a01b039091169063f88d3d42906125f1908a908a9088908890600401615b86565b6020604051808303815f875af115801561260d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126319190615bc3565b9050612641875f83898989610e5b565b9695505050505050565b612653612ae9565b6001600160a01b03811661267a57604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b0392909216919091179055565b5f5b8251811015612579576126cb848483815181106126bd576126bd615a26565b602002602001015184613c23565b60010161269e565b5f818152603760205260408120600201546119a5906001615b05565b82811461270f576040516329f517fd60e01b815260040160405180910390fd5b5f5b8381101561287b5761287285858381811061272e5761272e615a26565b90506020028101906127409190615bde565b3586868481811061275357612753615a26565b90506020028101906127659190615bde565b612776906040810190602001615065565b87878581811061278857612788615a26565b905060200281019061279a9190615bde565b6127ab906060810190604001615065565b8888868181106127bd576127bd615a26565b90506020028101906127cf9190615bde565b606001358989878181106127e5576127e5615a26565b90506020028101906127f79190615bde565b608001358a8a8881811061280d5761280d615a26565b905060200281019061281f9190615bde565b60a001358b8b8981811061283557612835615a26565b90506020028101906128479190615bde565b6128559060c0810190615a3a565b8b8b8b81811061286757612867615a26565b905060e0020161147e565b50600101612711565b5050505050565b6096546040516337dba1f760e21b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063df6e87dc90606401606060405180830381865afa1580156128de573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129029190615bf2565b91989097509095509350505050565b612919612ae9565b6032546001600160a01b031680156129505760405163f6c75f8f60e01b81526001600160a01b039091166004820152602401610aaf565b506001600160a01b03811661297857604051636ca1fdd760e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0383169081179091556040517fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee905f90a250565b6129c9612ae9565b6001600160a01b0381166129f257604051631e4fbdf760e01b81525f6004820152602401610aaf565b61228c81613a14565b612a03612ae9565b5f828152603860205260409020612a1a9082613d52565b8190612a455760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f8281526039602090815260408083206001600160a01b03851680855290835281842080546001600160a01b03191681556001810180546001600160b01b03191690556002810185905560038101859055600401849055858452603a835281842081855290925280832083905551909184917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a35050565b5f61258b613d66565b33612af2612001565b6001600160a01b031614611e075760405163118cdaa760e01b8152336004820152602401610aaf565b60605f611eb683613d6f565b612b2f614c12565b5f838152603b60205260409020612b469083613dc8565b8290612b6857604051637f11bea960e01b8152600401610aaf91815260200190565b505f838152603c60209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190612be090615b18565b80601f0160208091040260200160405190810160405280929190818152602001828054612c0c90615b18565b8015612c575780601f10612c2e57610100808354040283529160200191612c57565b820191905f5260205f20905b815481529060010190602001808311612c3a57829003601f168201915b505050919092525050505f848152603960209081526040808320848201516001600160a01b031684529091529020600181015491925090600160a81b900460ff1615612cba578160800151816004015f828254612cb49190615c1d565b90915550505b5f848152603c602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181612d136005830182614c55565b5050600682015f9055600782015f612d2b9190614c55565b50505092915050565b5f516020615f3a5f395f51905f525460ff1615611e075760405163d93c066560e01b815260040160405180910390fd5b5f516020615f5a5f395f51905f52805460011901612d9557604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f9081526001830160205260408120541515611eb6565b8251825181148015612dce5750815181145b835183518392612e02576040516324227ae160e21b8152600481019390935260248301919091526044820152606401610aaf565b505060645482915060ff16811015612e3057604051631aebd17960e11b8152600401610aaf91815260200190565b505f80826001600160401b03811115612e4b57612e4b614cb6565b604051908082528060200260200182016040528015612e74578160200160208202803683370190505b5090505f5b83811015612fd6575f612ee4888381518110612e9757612e97615a26565b6020026020010151888481518110612eb157612eb1615a26565b6020026020010151888581518110612ecb57612ecb615a26565b6020026020010151612edc8d613dd3565b929190613dff565b9050612efc682b20a624a220aa27a960b91b8261202f565b682b20a624a220aa27a960b91b829091612f3b5760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610aaf565b505f9050805b85811015612f8c57848181518110612f5b57612f5b615a26565b60200260200101516001600160a01b0316836001600160a01b031603612f845760019150612f8c565b600101612f41565b5080612fcc5781848681518110612fa557612fa5615a26565b60200260200101906001600160a01b031690816001600160a01b0316815250508460010194505b5050600101612e79565b50606454829060ff1681101561300257604051631aebd17960e11b8152600401610aaf91815260200190565b5050505050505050565b5f848152603a602090815260408083206001600160a01b0387168452909152812054606090600181131561304b576130448185615c30565b935061306a565b5f1981121561306a5761305d81615c47565b6130679085615c75565b93505b5f196001600160a01b0387160161309e57613094858560405180602001604052805f8152506134aa565b92509250506130f1565b83156130ef575f8781526039602090815260408083206001600160a01b038a168452909152902060010154600160a81b900460ff16156130e45761309487878787613e2b565b613094868686613f21565b505b94509492505050565b82355f908152603b60209081526040909120613118918501356132ba565b83602001359061313e576040516307a5c91d60e31b8152600401610aaf91815260200190565b505f8115613167576034546131539042615b05565b9050603454426131639190615b05565b6033555b60405180606001604052808561317c90615c88565b81526020808201849052604091820186905286355f908152603c82528281208883013582528252829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a08101518290600582019061320e9082615d52565b505050602082015160068201556040820151600782019061322f9082615d52565b50505083355f908152603960205260408082209082906132559060608901908901615065565b6001600160a01b0316815260208101919091526040015f206001810154909150600160a81b900460ff161561287b578460800135816004015f82825461329b9190615b05565b90915550505050505050565b60015f516020615f5a5f395f51905f5255565b5f611eb68383614002565b5f611eb6836001600160a01b038416614002565b5f8181526001830160205260408120541515611eb6565b5f8281526038602052604090206133079082612d9b565b81906133325760405163153096f360e11b81526001600160a01b039091166004820152602401610aaf565b505f82815260376020526040902060030154829060ff161561336a57604051636fc24b7960e11b8152600401610aaf91815260200190565b505f8281526039602090815260408083206001600160a01b03851684529091529020600101548190600160a01b900460ff1615610a3b576040516338384f6f60e11b81526001600160a01b039091166004820152602401610aaf565b6096545f9081906001600160a01b03166133e457505f9050806134a0565b6096546040516337dba1f760e21b8152600481018990526001600160a01b038881166024830152604482018890525f92169063df6e87dc90606401606060405180830381865afa15801561343a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061345e9190615bf2565b909450925090508086108015906134755750828510155b80156134815750818410155b61349e576040516358e8878560e01b815260040160405180910390fd5b505b9550959350505050565b5f60608347478211156134d957604051632f2a246160e21b815260048101929092526024820152604401610aaf565b50506060856001600160a01b031685856040516134f69190615e0c565b5f6040518083038185875af1925050503d805f8114613530576040519150601f19603f3d011682016040523d82523d5f602084013e613535565b606091505b5090935090508261357a57805115613551575f9250905061361f565b50506040805180820190915260088152671c995d995c9d195960c21b60208201525f915061361f565b845f0361360a5780515f036135c857856001600160a01b03163b5f036135c35750506040805180820190915260088152676e6f7420636f646560c01b60208201525f915061361f565b61360a565b6020810151600181151514613608575f6040518060400160405280600c81526020016b72657475726e2066616c736560a01b81525093509350505061361f565b505b505060408051602081019091525f8152600191505b935093915050565b61363d8a8a8a89613638898b615b05565b61404e565b5f8a815260376020526040812060019081018054909101908190558190915060395f8d81526020019081526020015f205f8c6001600160a01b03166001600160a01b031681526020019081526020015f206001015f9054906101000a90046001600160a01b03169050896001600160a01b0316828d7f732adeda688b3a630ef0effe2e3fc34e22b5ae8c24cdb7d3ba5a0fb3a004db508e858e8e8e8e428f8f8f6040516136f39a99989796959493929190615e22565b60405180910390a4505050505050505050505050565b5f5f6137158484612b27565b90505f5f613734835f015184604001518560600151866080015161300c565b915091508181906137585760405162461bcd60e51b8152600401610aaf91906157e1565b5082604001516001600160a01b03168360200151845f01517f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b86606001518760800151426040516137c7939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a450600195945050505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061386157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166138555f516020615f1a5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15611e075760405163703e46dd60e11b815260040160405180910390fd5b61228c612ae9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156138e1575060408051601f3d908101601f191682019092526138de91810190615e8c565b60015b61390957604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610aaf565b5f516020615f1a5f395f51905f52811461393957604051632a87526960e21b815260048101829052602401610aaf565b610a3b83836142cd565b61394b614322565b6139543361436b565b61395c61437c565b613964614384565b61396c614394565b613975826143a4565b61397d61440a565b6001600160a01b0381166139a457604051638219abe360e01b815260040160405180910390fd5b609780546001600160a01b0319166001600160a01b03929092169190911790555043609855565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611e075760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615efa5f395f51905f5291613ac290615b18565b80601f0160208091040260200160405190810160405280929190818152602001828054613aee90615b18565b8015613b395780601f10613b1057610100808354040283529160200191613b39565b820191905f5260205f20905b815481529060010190602001808311613b1c57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615efa5f395f51905f5291613ac290615b18565b613b8a612d34565b5f516020615f3a5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200161226b565b613be6614436565b5f516020615f3a5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613bc6565b613c2b612ae9565b6001600160a01b038216613c5257604051635989b9d360e01b815260040160405180910390fd5b82613c7057604051630e1b39f960e31b815260040160405180910390fd5b8015613cc7575f838152602081905260409020613c8d90836132c5565b83839091613cc057604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610aaf565b5050613d14565b5f838152602081905260409020613cde9083613d52565b83839091613d1157604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610aaf565b50505b82826001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f383604051611223911515815260200190565b5f611eb6836001600160a01b038416614465565b5f61258b614548565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613dbc57602002820191905f5260205f20905b815481526020019060010190808311613da8575b50505050509050919050565b5f611eb68383614465565b5f6119a5613ddf613d66565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613e0f888888886145bb565b925092509250613e1f8282614683565b50909695505050505050565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613e9a575060408051601f3d908101601f19168201909252613e9791810190615ea3565b60015b613ed4573d808015613ec7576040519150601f19603f3d011682016040523d82523d5f602084013e613ecc565b606091505b5090506130f1565b8092508015613eed57613ee887878661473b565b6130ef565b6040518060400160405280600f81526020016e1d1c985b9cd9995c8819985a5b1959608a1b81525091505094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613f90575060408051601f3d908101601f19168201909252613f8d91810190615ea3565b60015b613fca573d808015613fbd576040519150601f19603f3d011682016040523d82523d5f602084013e613fc2565b606091505b50905061361f565b80925080613ff9576040518060400160405280600b81526020016a1b5a5b9d0819985a5b195960aa1b81525091505b50935093915050565b5f81815260018301602052604081205461404757508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556119a5565b505f6119a5565b5f858152603a602090815260408083206001600160a01b038816845290915290205460018113156140b9576140838184615ebe565b85908490156140b6576040516305330e5560e11b81526001600160a01b0390921660048301526024820152604401610aaf565b50505b5f196001600160a01b03861601614169576140d48284615b05565b34146140e08385615b05565b349091614109576040516362624a5160e11b815260048101929092526024820152604401610aaf565b505081156141645760975460408051602081019091525f80825291829161413b916001600160a01b03169086906134aa565b915091508181906141605760405163e8b5c4bb60e01b8152600401610aaf91906157e1565b5050505b6142c5565b5f348015614193576040516362624a5160e11b815260048101929092526024820152604401610aaf565b506141b7905084306141a58587615b05565b6001600160a01b0389169291906147e6565b81156141d7576097546141d7906001600160a01b0387811691168461484d565b5f8681526039602090815260408083206001600160a01b0389168452909152902060010154600160a81b900460ff16156142165761416486868561487e565b604051632770a7eb60e21b8152306004820152602481018490526001600160a01b03861690639dc29fac906044016020604051808303815f875af1158015614260573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142849190615ea3565b8585859091926142c157604051639ac2f56d60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610aaf565b5050505b505050505050565b6142d6826148bc565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561431a57610a3b828261491f565b611139614991565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16611e0757604051631afcd79f60e31b815260040160405180910390fd5b614373614322565b61228c816149b0565b611e07614322565b61438c614322565b611e076149b8565b61439c614322565b611e076149d8565b6143ac614322565b6143f4604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506149e0565b6064805460ff191660ff92909216919091179055565b614412614322565b61442d6420a226a4a760d91b614426612001565b6001613c23565b62015180603455565b5f516020615f3a5f395f51905f525460ff16611e0757604051638dfc202b60e01b815260040160405180910390fd5b5f818152600183016020526040812054801561453f575f614487600183615c1d565b85549091505f9061449a90600190615c1d565b90508082146144f9575f865f0182815481106144b8576144b8615a26565b905f5260205f200154905080875f0184815481106144d8576144d8615a26565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061450a5761450a615ed1565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506119a5565b5f9150506119a5565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6145726149f2565b61457a614a5a565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156145f457505f91506003905082614679565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614645573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661467057505f925060019150829050614679565b92505f91508190505b9450945094915050565b5f82600381111561469657614696615ee5565b0361469f575050565b60018260038111156146b3576146b3615ee5565b036146d15760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156146e5576146e5615ee5565b036147065760405163fce698f760e01b815260048101829052602401610aaf565b600382600381111561471a5761471a615ee5565b03611139576040516335e2f38360e21b815260048101829052602401610aaf565b5f8381526039602090815260408083206001600160a01b03861684529091529020600481015461476b9083615b05565b6003820154600483015486928692869290918210156147c3576040516352c2db3360e01b815260048101959095526001600160a01b03909316602485015260448401919091526064830152608482015260a401610aaf565b505050505081816003015f8282546147db9190615c1d565b909155505050505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526125799186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614a9c565b6040516001600160a01b03838116602483015260448201839052610a3b91859182169063a9059cbb9060640161481b565b5f8381526039602090815260408083206001600160a01b0386168452909152812060030180548392906148b2908490615b05565b9091555050505050565b806001600160a01b03163b5f036148f157604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610aaf565b5f516020615f1a5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161493b9190615e0c565b5f60405180830381855af49150503d805f8114614973576040519150601f19603f3d011682016040523d82523d5f602084013e614978565b606091505b5091509150614988858383614b08565b95945050505050565b3415611e075760405163b398979f60e01b815260040160405180910390fd5b6129c9614322565b6149c0614322565b5f516020615f3a5f395f51905f52805460ff19169055565b6132a7614322565b6149e8614322565b6111398282614b64565b5f5f516020615efa5f395f51905f5281614a0a613a84565b805190915015614a2257805160209091012092915050565b81548015614a31579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615efa5f395f51905f5281614a72613b44565b805190915015614a8a57805160209091012092915050565b60018201548015614a31579392505050565b5f5f60205f8451602086015f885af180614abb576040513d5f823e3d81fd5b50505f513d91508115614ad2578060011415614adf565b6001600160a01b0384163b155b1561257957604051635274afe760e01b81526001600160a01b0385166004820152602401610aaf565b606082614b1d57614b1882614bc3565b611eb6565b8151158015614b3457506001600160a01b0384163b155b15614b5d57604051639996b31560e01b81526001600160a01b0385166004820152602401610aaf565b5092915050565b614b6c614322565b5f516020615efa5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614ba58482615d52565b5060038101614bb48382615d52565b505f8082556001909101555050565b805115614bd35780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614bff614c12565b81526020015f8152602001606081525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614c6190615b18565b5f825580601f10614c70575050565b601f0160209004905f5260205f209081019061228c91905b80821115614c9b575f8155600101614c88565b5090565b5f60208284031215614caf575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614cec57614cec614cb6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614d1a57614d1a614cb6565b604052919050565b5f6001600160401b03821115614d3a57614d3a614cb6565b5060051b60200190565b803560ff81168114614d54575f5ffd5b919050565b5f82601f830112614d68575f5ffd5b8135614d7b614d7682614d22565b614cf2565b8082825260208201915060208360051b860101925085831115614d9c575f5ffd5b602085015b83811015614dc057614db281614d44565b835260209283019201614da1565b5095945050505050565b5f82601f830112614dd9575f5ffd5b8135614de7614d7682614d22565b8082825260208201915060208360051b860101925085831115614e08575f5ffd5b602085015b83811015614dc0578035835260209283019201614e0d565b5f5f5f5f60808587031215614e38575f5ffd5b84356001600160401b03811115614e4d575f5ffd5b850160c08188031215614e5e575f5ffd5b935060208501356001600160401b03811115614e78575f5ffd5b614e8487828801614d59565b93505060408501356001600160401b03811115614e9f575f5ffd5b614eab87828801614dca565b92505060608501356001600160401b03811115614ec6575f5ffd5b614ed287828801614dca565b91505092959194509250565b801515811461228c575f5ffd5b6001600160a01b038116811461228c575f5ffd5b5f5f5f5f5f5f60c08789031215614f14575f5ffd5b863595506020870135614f2681614ede565b94506040870135614f3681614eeb565b93506060870135614f4681614eeb565b9598949750929560808101359460a0909101359350915050565b5f82601f830112614f6f575f5ffd5b8135614f7d614d7682614d22565b8082825260208201915060208360051b860101925085831115614f9e575f5ffd5b602085015b83811015614dc0578035614fb681614eeb565b835260209283019201614fa3565b5f5f60408385031215614fd5575f5ffd5b8235915060208301356001600160401b03811115614ff1575f5ffd5b614ffd85828601614f60565b9150509250929050565b5f5f5f60608486031215615019575f5ffd5b83359250602084013561502b81614eeb565b929592945050506040919091013590565b5f5f5f6060848603121561504e575f5ffd5b505081359360208301359350604090920135919050565b5f60208284031215615075575f5ffd5b8135611eb681614eeb565b5f5f5f60608486031215615092575f5ffd5b8335925060208401356150a481614eeb565b915060408401356150b481614ede565b809150509250925092565b5f5f83601f8401126150cf575f5ffd5b5081356001600160401b038111156150e5575f5ffd5b6020830191508360208285010111156150fc575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f898b036101c081121561511d575f5ffd5b8a35995060208b013561512f81614eeb565b985060408b013561513f81614eeb565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b0381111561516e575f5ffd5b61517a8d828e016150bf565b90955093505060e060df1982011215615191575f5ffd5b5060e08a0190509295985092959850929598565b5f5f604083850312156151b6575f5ffd5b50508035926020909101359150565b5f5f6001600160401b038411156151de576151de614cb6565b50601f8301601f19166020016151f381614cf2565b915050828152838383011115615207575f5ffd5b828260208301375f602084830101529392505050565b5f82601f83011261522c575f5ffd5b611eb6838335602085016151c5565b5f5f6040838503121561524c575f5ffd5b823561525781614eeb565b915060208301356001600160401b03811115615271575f5ffd5b614ffd8582860161521d565b5f5f6040838503121561528e575f5ffd5b61529783614d44565b915060208301356152a781614eeb565b809150509250929050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b81811015615347576153318385516152b2565b6020939093019260e0929092019160010161531e565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215615369575f5ffd5b88359750602089013561537b81614eeb565b9650604089013561538b81614eeb565b9550606089013594506080890135935060a0890135925060c08901356001600160401b038111156153ba575f5ffd5b6153c68b828c016150bf565b999c989b5096995094979396929594505050565b5f5f604083850312156153eb575f5ffd5b8235915060208301356152a781614ede565b5f5f6040838503121561540e575f5ffd5b8235915060208301356001600160401b0381111561542a575f5ffd5b8301601f8101851361543a575f5ffd5b8035615448614d7682614d22565b8082825260208201915060208360051b850101925087831115615469575f5ffd5b6020840193505b8284101561548b578335825260209384019390910190615470565b809450505050509250929050565b5f8151808452602084019350602083015f5b828110156154c95781518652602095860195909101906001016154ab565b5093949350505050565b602081525f611eb66020830184615499565b5f5f604083850312156154f6575f5ffd5b8235915060208301356152a781614eeb565b60e081016119a582846152b2565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61556260e0830189615516565b82810360408401526155748189615516565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506155a58185615499565b9a9950505050505050505050565b5f5f83601f8401126155c3575f5ffd5b5081356001600160401b038111156155d9575f5ffd5b6020830191508360208260051b85010111156150fc575f5ffd5b5f82601f830112615602575f5ffd5b8135615610614d7682614d22565b8082825260208201915060208360051b860101925085831115615631575f5ffd5b602085015b83811015614dc05780356001600160401b03811115615653575f5ffd5b615662886020838a0101614dca565b84525060209283019201615636565b5f5f5f5f5f60808688031215615685575f5ffd5b85356001600160401b0381111561569a575f5ffd5b6156a6888289016155b3565b90965094505060208601356001600160401b038111156156c4575f5ffd5b8601601f810188136156d4575f5ffd5b80356156e2614d7682614d22565b8082825260208201915060208360051b85010192508a831115615703575f5ffd5b602084015b838110156157435780356001600160401b03811115615725575f5ffd5b6157348d602083890101614d59565b84525060209283019201615708565b50955050505060408601356001600160401b03811115615761575f5ffd5b61576d888289016155f3565b92505060608601356001600160401b03811115615788575f5ffd5b615794888289016155f3565b9150509295509295909350565b602080825282518282018190525f918401906040840190835b818110156153475783516001600160a01b03168352602093840193909201916001016157ba565b602081525f611eb66020830184615516565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c061012084015261585e610140840182615516565b9050602084015160408401526040840151601f198483030160608501526149888282615516565b5f60208284031215615895575f5ffd5b611eb682614d44565b5f602082840312156158ae575f5ffd5b8135611eb681614ede565b5f5f5f5f5f5f60c087890312156158ce575f5ffd5b8635955060208701356158e081614eeb565b9450604087013593506060870135925060808701356001600160401b03811115615908575f5ffd5b8701601f81018913615918575f5ffd5b615927898235602084016151c5565b92505061593660a08801614d44565b90509295509295509295565b5f5f5f60608486031215615954575f5ffd5b8335925060208401356001600160401b03811115615970575f5ffd5b61597c86828701614f60565b92505060408401356150b481614ede565b5f5f5f5f604085870312156159a0575f5ffd5b84356001600160401b038111156159b5575f5ffd5b6159c1878288016155b3565b90955093505060208501356001600160401b038111156159df575f5ffd5b8501601f810187136159ef575f5ffd5b80356001600160401b03811115615a04575f5ffd5b87602060e083028401011115615a18575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112615a4f575f5ffd5b8301803591506001600160401b03821115615a68575f5ffd5b6020019150368190038213156150fc575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90615ae49083018486615a7c565b9998505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156119a5576119a5615af1565b600181811c90821680615b2c57607f821691505b602082108103615b4a57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be19833603018112615b64575f5ffd5b9190910192915050565b5f60018201615b7f57615b7f615af1565b5060010190565b8481526001600160a01b03841660208201526080604082018190525f90615baf90830185615516565b905060ff8316606083015295945050505050565b5f60208284031215615bd3575f5ffd5b8151611eb681614eeb565b5f823560de19833603018112615b64575f5ffd5b5f5f5f60608486031215615c04575f5ffd5b5050815160208301516040909301519094929350919050565b818103818111156119a5576119a5615af1565b80820281158282048414176119a5576119a5615af1565b5f600160ff1b8201615c5b57615c5b615af1565b505f0390565b634e487b7160e01b5f52601260045260245ffd5b5f82615c8357615c83615c61565b500490565b5f60c08236031215615c98575f5ffd5b615ca0614cca565b82358152602080840135908201526040830135615cbc81614eeb565b60408201526060830135615ccf81614eeb565b60608201526080838101359082015260a08301356001600160401b03811115615cf6575f5ffd5b615d023682860161521d565b60a08301525092915050565b601f821115610a3b57805f5260205f20601f840160051c81016020851015615d335750805b601f840160051c820191505b8181101561287b575f8155600101615d3f565b81516001600160401b03811115615d6b57615d6b614cb6565b615d7f81615d798454615b18565b84615d0e565b6020601f821160018114615db1575f8315615d9a5750848201515b5f19600385901b1c1916600184901b17845561287b565b5f84815260208120601f198516915b82811015615de05787850151825560209485019460019092019101615dc0565b5084821015615dfd57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b6001600160a01b038b811682528a8116602083015289166040820152606081018890526080810187905260a0810186905260c0810185905283151560e082015261012061010082018190525f90615e7c9083018486615a7c565b9c9b505050505050505050505050565b5f60208284031215615e9c575f5ffd5b5051919050565b5f60208284031215615eb3575f5ffd5b8151611eb681614ede565b5f82615ecc57615ecc615c61565b500690565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220f57cde9e0771dc753d073f2ed82c2fccef88850e062a827640e184b07527b6cc64736f6c634300081c0033",
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BaseBridge *BaseBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BaseBridge *BaseBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_BaseBridge *BaseBridgeCaller) BridgeFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "bridgeFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_BaseBridge *BaseBridgeSession) BridgeFeeManager() (common.Address, error) {
	return _BaseBridge.Contract.BridgeFeeManager(&_BaseBridge.CallOpts)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0x25bd5666.
//
// Solidity: function bridgeFeeManager() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) BridgeFeeManager() (common.Address, error) {
	return _BaseBridge.Contract.BridgeFeeManager(&_BaseBridge.CallOpts)
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
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
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
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_BaseBridge *BaseBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint256,bytes))
func (_BaseBridge *BaseBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BaseBridge *BaseBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BaseBridge *BaseBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_BaseBridge *BaseBridgeCaller) HasExpiredPending(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "hasExpiredPending")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_BaseBridge *BaseBridgeSession) HasExpiredPending() (bool, error) {
	return _BaseBridge.Contract.HasExpiredPending(&_BaseBridge.CallOpts)
}

// HasExpiredPending is a free data retrieval call binding the contract method 0x3d507c5e.
//
// Solidity: function hasExpiredPending() view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) HasExpiredPending() (bool, error) {
	return _BaseBridge.Contract.HasExpiredPending(&_BaseBridge.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseBridge *BaseBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseBridge *BaseBridgeSession) Owner() (common.Address, error) {
	return _BaseBridge.Contract.Owner(&_BaseBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) Owner() (common.Address, error) {
	return _BaseBridge.Contract.Owner(&_BaseBridge.CallOpts)
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

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_BaseBridge *BaseBridgeTransactor) ClearPending(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "clearPending", remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_BaseBridge *BaseBridgeSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ClearPending(&_BaseBridge.TransactOpts, remoteChainID)
}

// ClearPending is a paid mutator transaction binding the contract method 0x0b43c02c.
//
// Solidity: function clearPending(uint256 remoteChainID) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ClearPending(remoteChainID *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ClearPending(&_BaseBridge.TransactOpts, remoteChainID)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0xd016d625.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, conversionRatio, verificationAmountThreshold, symbol, decimals)
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

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "initialize", _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeSession) Initialize(_threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, _threshold, dev_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5187599d.
//
// Solidity: function initialize(uint8 _threshold, address dev_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) Initialize(_threshold uint8, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, _threshold, dev_)
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

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x1e7bf215.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address, conversionRatio *big.Int, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken, conversionRatio, verificationAmountThreshold)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_BaseBridge *BaseBridgeTransactor) ReleaseExpiredPending(opts *bind.TransactOpts, maxCount *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "releaseExpiredPending", maxCount)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_BaseBridge *BaseBridgeSession) ReleaseExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleaseExpiredPending(&_BaseBridge.TransactOpts, maxCount)
}

// ReleaseExpiredPending is a paid mutator transaction binding the contract method 0xc1cb05dd.
//
// Solidity: function releaseExpiredPending(uint256 maxCount) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ReleaseExpiredPending(maxCount *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleaseExpiredPending(&_BaseBridge.TransactOpts, maxCount)
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

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x6b0e4821.
//
// Solidity: function releasePendingBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeTransactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "releasePendingBatch", remoteChainID, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x6b0e4821.
//
// Solidity: function releasePendingBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeSession) ReleasePendingBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePendingBatch(&_BaseBridge.TransactOpts, remoteChainID, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x6b0e4821.
//
// Solidity: function releasePendingBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) ReleasePendingBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePendingBatch(&_BaseBridge.TransactOpts, remoteChainID, indexes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseBridge *BaseBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseBridge *BaseBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceOwnership(&_BaseBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaseBridge *BaseBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceOwnership(&_BaseBridge.TransactOpts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_BaseBridge *BaseBridgeTransactor) ResetRole(opts *bind.TransactOpts, role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "resetRole", role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_BaseBridge *BaseBridgeSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.ResetRole(&_BaseBridge.TransactOpts, role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.ResetRole(&_BaseBridge.TransactOpts, role, newAccounts)
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

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_BaseBridge *BaseBridgeTransactor) SetFeeManager(opts *bind.TransactOpts, _bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setFeeManager", _bridgeFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_BaseBridge *BaseBridgeSession) SetFeeManager(_bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetFeeManager(&_BaseBridge.TransactOpts, _bridgeFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _bridgeFeeManager) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetFeeManager(_bridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetFeeManager(&_BaseBridge.TransactOpts, _bridgeFeeManager)
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

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetPauseChain(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setPauseChain", remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPauseChain(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseChain is a paid mutator transaction binding the contract method 0x6160751f.
//
// Solidity: function setPauseChain(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetPauseChain(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPauseChain(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetPauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setPauseToken", remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPauseToken(&_BaseBridge.TransactOpts, remoteChainID, token, pause)
}

// SetPauseToken is a paid mutator transaction binding the contract method 0x4d3f0da9.
//
// Solidity: function setPauseToken(uint256 remoteChainID, address token, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetPauseToken(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPauseToken(&_BaseBridge.TransactOpts, remoteChainID, token, pause)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_BaseBridge *BaseBridgeTransactor) SetRole(opts *bind.TransactOpts, role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setRole", role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_BaseBridge *BaseBridgeSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetRole(&_BaseBridge.TransactOpts, role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetRole(&_BaseBridge.TransactOpts, role, accounts, set)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeTransactor) SetSafetyLimit(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setSafetyLimit", remoteChainID, token, verificationAmountThreshold)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetSafetyLimit(&_BaseBridge.TransactOpts, remoteChainID, token, verificationAmountThreshold)
}

// SetSafetyLimit is a paid mutator transaction binding the contract method 0x39a621f3.
//
// Solidity: function setSafetyLimit(uint256 remoteChainID, address token, uint256 verificationAmountThreshold) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetSafetyLimit(remoteChainID *big.Int, token common.Address, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetSafetyLimit(&_BaseBridge.TransactOpts, remoteChainID, token, verificationAmountThreshold)
}

// SetVerificationDeadline is a paid mutator transaction binding the contract method 0x3ca5b39f.
//
// Solidity: function setVerificationDeadline(uint256 remoteChainID, uint256 index, uint256 deadline) returns()
func (_BaseBridge *BaseBridgeTransactor) SetVerificationDeadline(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setVerificationDeadline", remoteChainID, index, deadline)
}

// SetVerificationDeadline is a paid mutator transaction binding the contract method 0x3ca5b39f.
//
// Solidity: function setVerificationDeadline(uint256 remoteChainID, uint256 index, uint256 deadline) returns()
func (_BaseBridge *BaseBridgeSession) SetVerificationDeadline(remoteChainID *big.Int, index *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDeadline(&_BaseBridge.TransactOpts, remoteChainID, index, deadline)
}

// SetVerificationDeadline is a paid mutator transaction binding the contract method 0x3ca5b39f.
//
// Solidity: function setVerificationDeadline(uint256 remoteChainID, uint256 index, uint256 deadline) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetVerificationDeadline(remoteChainID *big.Int, index *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDeadline(&_BaseBridge.TransactOpts, remoteChainID, index, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseBridge *BaseBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseBridge *BaseBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.TransferOwnership(&_BaseBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BaseBridge *BaseBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.TransferOwnership(&_BaseBridge.TransactOpts, newOwner)
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
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgePending is a free log retrieval operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
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

// WatchBridgePending is a free log subscription operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
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

// ParseBridgePending is a log parse operation binding the contract event 0xd3122beb443f4b66faef07d45daf96faf0ffc23e050550cf1bbcbe6105f6bee7.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) ParseBridgePending(log types.Log) (*BaseBridgeBridgePending, error) {
	event := new(BaseBridgeBridgePending)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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

// BaseBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaseBridge contract.
type BaseBridgeOwnershipTransferredIterator struct {
	Event *BaseBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaseBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeOwnershipTransferred)
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
		it.Event = new(BaseBridgeOwnershipTransferred)
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
func (it *BaseBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the BaseBridge contract.
type BaseBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseBridge *BaseBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaseBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeOwnershipTransferredIterator{contract: _BaseBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BaseBridge *BaseBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaseBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeOwnershipTransferred)
				if err := _BaseBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BaseBridgeOwnershipTransferred, error) {
	event := new(BaseBridgeOwnershipTransferred)
	if err := _BaseBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaseBridgeRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the BaseBridge contract.
type BaseBridgeRoleUpdatedIterator struct {
	Event *BaseBridgeRoleUpdated // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleUpdated)
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
		it.Event = new(BaseBridgeRoleUpdated)
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
func (it *BaseBridgeRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleUpdated represents a RoleUpdated event raised by the BaseBridge contract.
type BaseBridgeRoleUpdated struct {
	Account common.Address
	Role    [32]byte
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleUpdated(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*BaseBridgeRoleUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleUpdated", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleUpdatedIterator{contract: _BaseBridge.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleUpdated, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleUpdated", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleUpdated)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool status)
func (_BaseBridge *BaseBridgeFilterer) ParseRoleUpdated(log types.Log) (*BaseBridgeRoleUpdated, error) {
	event := new(BaseBridgeRoleUpdated)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeSafetyLimitSetIterator is returned from FilterSafetyLimitSet and is used to iterate over the raw logs and unpacked data for SafetyLimitSet events raised by the BaseBridge contract.
type BaseBridgeSafetyLimitSetIterator struct {
	Event *BaseBridgeSafetyLimitSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeSafetyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeSafetyLimitSet)
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
		it.Event = new(BaseBridgeSafetyLimitSet)
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
func (it *BaseBridgeSafetyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeSafetyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeSafetyLimitSet represents a SafetyLimitSet event raised by the BaseBridge contract.
type BaseBridgeSafetyLimitSet struct {
	RemoteChainID               *big.Int
	Token                       common.Address
	VerificationAmountThreshold *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSafetyLimitSet is a free log retrieval operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_BaseBridge *BaseBridgeFilterer) FilterSafetyLimitSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BaseBridgeSafetyLimitSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeSafetyLimitSetIterator{contract: _BaseBridge.contract, event: "SafetyLimitSet", logs: logs, sub: sub}, nil
}

// WatchSafetyLimitSet is a free log subscription operation binding the contract event 0xd54edc9ee33780fc813e88b7fcafd9163ae63168837ee363914a2f892670d7c8.
//
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_BaseBridge *BaseBridgeFilterer) WatchSafetyLimitSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeSafetyLimitSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "SafetyLimitSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeSafetyLimitSet)
				if err := _BaseBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
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
// Solidity: event SafetyLimitSet(uint256 indexed remoteChainID, address indexed token, uint256 verificationAmountThreshold)
func (_BaseBridge *BaseBridgeFilterer) ParseSafetyLimitSet(log types.Log) (*BaseBridgeSafetyLimitSet, error) {
	event := new(BaseBridgeSafetyLimitSet)
	if err := _BaseBridge.contract.UnpackLog(event, "SafetyLimitSet", log); err != nil {
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
	RemoteChainID               *big.Int
	LocalToken                  common.Address
	RemoteToken                 common.Address
	ConversionRatio             *big.Int
	VerificationAmountThreshold *big.Int
	IsOrigin                    bool
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
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

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xffd50b3f28a6ab2baf4fb67d8aaf0a6c918767d3e790535a54eb250c503b2a66.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, int256 conversionRatio, uint256 verificationAmountThreshold, bool isOrigin)
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

// BaseBridgeVerificationDeadlineSetIterator is returned from FilterVerificationDeadlineSet and is used to iterate over the raw logs and unpacked data for VerificationDeadlineSet events raised by the BaseBridge contract.
type BaseBridgeVerificationDeadlineSetIterator struct {
	Event *BaseBridgeVerificationDeadlineSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeVerificationDeadlineSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeVerificationDeadlineSet)
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
		it.Event = new(BaseBridgeVerificationDeadlineSet)
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
func (it *BaseBridgeVerificationDeadlineSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeVerificationDeadlineSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeVerificationDeadlineSet represents a VerificationDeadlineSet event raised by the BaseBridge contract.
type BaseBridgeVerificationDeadlineSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Deadline    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDeadlineSet is a free log retrieval operation binding the contract event 0xfcfd532dc7387ad29a63519b09004fa6a982a1853a71522d0d5fe94f3ef76d0d.
//
// Solidity: event VerificationDeadlineSet(uint256 indexed fromChainID, uint256 indexed index, uint256 deadline)
func (_BaseBridge *BaseBridgeFilterer) FilterVerificationDeadlineSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BaseBridgeVerificationDeadlineSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "VerificationDeadlineSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeVerificationDeadlineSetIterator{contract: _BaseBridge.contract, event: "VerificationDeadlineSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDeadlineSet is a free log subscription operation binding the contract event 0xfcfd532dc7387ad29a63519b09004fa6a982a1853a71522d0d5fe94f3ef76d0d.
//
// Solidity: event VerificationDeadlineSet(uint256 indexed fromChainID, uint256 indexed index, uint256 deadline)
func (_BaseBridge *BaseBridgeFilterer) WatchVerificationDeadlineSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeVerificationDeadlineSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "VerificationDeadlineSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeVerificationDeadlineSet)
				if err := _BaseBridge.contract.UnpackLog(event, "VerificationDeadlineSet", log); err != nil {
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

// ParseVerificationDeadlineSet is a log parse operation binding the contract event 0xfcfd532dc7387ad29a63519b09004fa6a982a1853a71522d0d5fe94f3ef76d0d.
//
// Solidity: event VerificationDeadlineSet(uint256 indexed fromChainID, uint256 indexed index, uint256 deadline)
func (_BaseBridge *BaseBridgeFilterer) ParseVerificationDeadlineSet(log types.Log) (*BaseBridgeVerificationDeadlineSet, error) {
	event := new(BaseBridgeVerificationDeadlineSet)
	if err := _BaseBridge.contract.UnpackLog(event, "VerificationDeadlineSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
