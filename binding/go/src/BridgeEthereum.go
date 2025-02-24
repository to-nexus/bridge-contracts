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

// BridgeEthereumMetaData contains all meta data concerning the BridgeEthereum contract.
var BridgeEthereumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeToCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenInfo\",\"outputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"changeRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cross\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_cross\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeEthereumCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeEthereumInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"d86c62dc": "bridgeCross(address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"a01be2c7": "bridgeToCross(address,address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"9c1b65a9": "bridgeTokenInfo()",
		"6e908ca3": "calculate(address,uint256)",
		"abe110b6": "changeRewardWallet(address)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
		"fa074d03": "cross()",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"2f63b7c8": "finalize(uint256,address,address,uint256,bytes[],uint8[],bytes32[],bytes32[])",
		"49a495ec": "finalizeBatch((uint256,address,address,uint256,bytes[])[],uint8[][],bytes32[][],bytes32[][])",
		"71c59d7b": "getPairToken(address)",
		"9c2c58f8": "initialize(address,uint8,address,address)",
		"91cf6d3e": "initializedAt()",
		"c1876453": "isValidToken(address)",
		"facd743b": "isValidator(address)",
		"27938c81": "nextFinalizeIndex()",
		"e1af7f50": "nextInitiateIndex()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"7c41ad2c": "pauseToken(address)",
		"5c975abb": "paused()",
		"51c45579": "permitBridge(address,address,uint256,uint256,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"174991ab": "permitBridgeTo(address,address,address,uint256,uint256,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"52d1902d": "proxiableUUID()",
		"5fa7b584": "removeToken(address)",
		"85547884": "removeTokenInfo()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
		"a9823183": "retryFinalize(uint256)",
		"c1ad8b95": "retryFinalizeBatch(uint256[])",
		"7021fd0e": "revertedArguments(uint256)",
		"fe2b8da6": "revertedReason(uint256)",
		"fb75b2c7": "rewardWallet()",
		"2f9b59d1": "setTokenInfo(address)",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"d92fc67b": "tokensLength()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"3b3bff0f": "unpauseToken(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614bcb6100f95f395f81816124b8015281816124e101526126300152614bcb5ff3fe608060405260043610610366575f3560e01c80638456cb59116101c8578063b7f3358d116100fd578063e1af7f501161009d578063fa074d031161006d578063fa074d03146109b6578063facd743b146109d5578063fb75b2c7146109f4578063fe2b8da614610a11575f5ffd5b8063e1af7f501461095b578063f2fde38b1461096f578063f30589c31461098e578063f698da25146109a2575f5ffd5b8063c97682f8116100d8578063c97682f8146108f5578063cbae595814610909578063d86c62dc14610928578063d92fc67b14610947575f5ffd5b8063b7f3358d14610898578063c1876453146108b7578063c1ad8b95146108d6575f5ffd5b80639c1b65a911610168578063a982318311610143578063a982318314610809578063abe110b614610828578063ad3cb1cc14610847578063aed1d40314610884575f5ffd5b80639c1b65a9146107ac5780639c2c58f8146107cb578063a01be2c7146107ea575f5ffd5b80638da5cb5b116101a35780638da5cb5b1461072957806391cf6d3e146107655780639300c9261461077957806396ce079514610798575f5ffd5b80638456cb59146106da57806384b0196e146106ee5780638554788414610715575f5ffd5b80634f6ccce71161029e5780636e908ca31161023e5780637101fcd3116102195780637101fcd314610651578063715018a61461067057806371c59d7b146106845780637c41ad2c146106bb575f5ffd5b80636e908ca3146105c35780636ff97f1d146106045780637021fd0e14610625575f5ffd5b80635476bd72116102795780635476bd72146105435780635c975abb146105625780635dbe47e8146105855780635fa7b584146105a4575f5ffd5b80634f6ccce7146104e557806351c455791461051c57806352d1902d1461052f575f5ffd5b80633b3bff0f1161030957806340a141ff116102e457806340a141ff1461047f57806342cde4e81461049e57806349a495ec146104bf5780634f1ef286146104d2575f5ffd5b80633b3bff0f146104395780633f4ba83a146104585780633f5d3b5d1461046c575f5ffd5b806327938c811161034457806327938c81146103d25780632f63b7c8146103f45780632f9b59d11461040757806337d1007514610426575f5ffd5b80631327d3d81461036a578063174991ab1461038b5780631d40f0d8146103b3575b5f5ffd5b348015610375575f5ffd5b5061038961038436600461396c565b610a30565b005b61039e610399366004613ac2565b610a3e565b60405190151581526020015b60405180910390f35b3480156103be575f5ffd5b506103896103cd366004613b8a565b610ad7565b3480156103dd575f5ffd5b506103e6610b11565b6040519081526020016103aa565b61039e610402366004613d04565b610b26565b348015610412575f5ffd5b5061038961042136600461396c565b610dca565b61039e610434366004613df0565b610e3e565b348015610444575f5ffd5b5061038961045336600461396c565b610e59565b348015610463575f5ffd5b50610389610e6a565b61039e61047a366004613e5d565b610e7c565b34801561048a575f5ffd5b5061038961049936600461396c565b610f14565b3480156104a9575f5ffd5b5060995460405160ff90911681526020016103aa565b61039e6104cd366004613f66565b610f1e565b6103896104e0366004614111565b61107e565b3480156104f0575f5ffd5b506105046104ff36600461415d565b611099565b6040516001600160a01b0390911681526020016103aa565b61039e61052a366004614174565b6110aa565b34801561053a575f5ffd5b506103e66110c9565b34801561054e575f5ffd5b5061038961055d366004614206565b6110e5565b34801561056d575f5ffd5b505f516020614b565f395f51905f525460ff1661039e565b348015610590575f5ffd5b5061039e61059f36600461396c565b6110f7565b3480156105af575f5ffd5b506103896105be36600461396c565b611102565b3480156105ce575f5ffd5b506105e26105dd36600461423d565b611113565b60408051948552602085019390935291830152151560608201526080016103aa565b34801561060f575f5ffd5b5061061861119c565b6040516103aa9190614267565b348015610630575f5ffd5b5061064461063f36600461415d565b6111fd565b6040516103aa91906142e0565b34801561065c575f5ffd5b5061038961066b366004613b8a565b61135d565b34801561067b575f5ffd5b50610389611373565b34801561068f575f5ffd5b5061050461069e36600461396c565b6001600160a01b039081165f908152603360205260409020541690565b3480156106c6575f5ffd5b506103896106d536600461396c565b611384565b3480156106e5575f5ffd5b50610389611395565b3480156106f9575f5ffd5b506107026113a5565b6040516103aa9796959493929190614385565b348015610720575f5ffd5b5061038961144e565b348015610734575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610504565b348015610770575f5ffd5b5060cc546103e6565b348015610784575f5ffd5b50610389610793366004613b8a565b611468565b3480156107a3575f5ffd5b506103e661149f565b3480156107b7575f5ffd5b5060ca54610504906001600160a01b031681565b3480156107d6575f5ffd5b506103896107e536600461441b565b61150a565b3480156107f5575f5ffd5b5061039e610804366004614472565b611681565b348015610814575f5ffd5b5061039e61082336600461415d565b6116a0565b348015610833575f5ffd5b5061038961084236600461396c565b61192a565b348015610852575f5ffd5b50610877604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516103aa91906144d4565b34801561088f575f5ffd5b506103e6611954565b3480156108a3575f5ffd5b506103896108b23660046144e6565b61195f565b3480156108c2575f5ffd5b5061039e6108d136600461396c565b6119b0565b3480156108e1575f5ffd5b5061039e6108f03660046144ff565b6119e0565b348015610900575f5ffd5b50610618611a20565b348015610914575f5ffd5b5061050461092336600461415d565b611b06565b348015610933575f5ffd5b5061039e610942366004614591565b611b12565b348015610952575f5ffd5b506103e6611b2b565b348015610966575f5ffd5b506103e6611b35565b34801561097a575f5ffd5b5061038961098936600461396c565b611b45565b348015610999575f5ffd5b50610618611b7f565b3480156109ad575f5ffd5b506103e6611b8b565b3480156109c1575f5ffd5b5060fc54610504906001600160a01b031681565b3480156109e0575f5ffd5b5061039e6109ef36600461396c565b611b94565b3480156109ff575f5ffd5b5060cb546001600160a01b0316610504565b348015610a1c575f5ffd5b50610877610a2b36600461415d565b611ba0565b610a3b816001611c3f565b50565b5f610a47611d08565b89610a51816119b0565b8190610a8157604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610a8a611d38565b5f5f610a988d8b8b8b611d6f565b91509150610aad8d8d8d8d86868d8d8d611e8f565b600193505050610ac960015f516020614b765f395f51905f5255565b509998505050505050505050565b5f5b8151811015610b0d57610b05828281518110610af757610af76145ff565b60200260200101515f611c3f565b600101610ad9565b5050565b5f6066546001610b219190614627565b905090565b5f610b2f611d08565b88610b39816119b0565b8190610b6457604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610a78565b50610b6d611d38565b610b75610b11565b8b14610b7f610b11565b8c9091610ba857604051634982205b60e11b815260048101929092526024820152604401610a78565b5050610c087fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8c8c8c8c8c8c604051602001610bea97969594939291906146ea565b60405160208183030381529060405280519060200120868686611f0d565b610c16606680546001019055565b5f5f610c238c8c8c612108565b915091508115610c89578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051610c7c929190918252602082015260400190565b60405180910390a4610aad565b610c9460cf8e61230b565b8d90610cb6576040516368db995b60e11b8152600401610a7891815260200190565b505f8d815260ce60205260409020610cce82826147a6565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90610d119190614860565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051610d81926004850192019061387b565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a2600193505050610ac960015f516020614b765f395f51905f5255565b610dd261231d565b6001600160a01b038116610e1c57604051633effd40360e21b815260206004820152601060248201526f5f627269646765546f6b656e496e666f60801b6044820152606401610a78565b60ca80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610e4e87338888888888610e7c565b979650505050505050565b610e6161231d565b610a3b81612378565b610e7261231d565b610e7a612418565b565b5f610e85611d08565b87610e8f816119b0565b8190610eba57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610a78565b50610ec3611d38565b5f5f610ed18b8a8a8a611d6f565b91509150610eec8b610ee03390565b8c8c86865f8d8d612471565b600193505050610f0860015f516020614b765f395f51905f5255565b50979650505050505050565b610a3b815f611c3f565b5f805b8581101561107157611068878783818110610f3e57610f3e6145ff565b9050602002810190610f5091906148cd565b35888884818110610f6357610f636145ff565b9050602002810190610f7591906148cd565b610f8690604081019060200161396c565b898985818110610f9857610f986145ff565b9050602002810190610faa91906148cd565b610fbb90606081019060400161396c565b8a8a86818110610fcd57610fcd6145ff565b9050602002810190610fdf91906148cd565b606001358b8b87818110610ff557610ff56145ff565b905060200281019061100791906148cd565b6110159060808101906148eb565b8b8881518110611027576110276145ff565b60200260200101518b8981518110611041576110416145ff565b60200260200101518b8a8151811061105b5761105b6145ff565b6020026020010151610b26565b50600101610f21565b5060019695505050505050565b6110866124ad565b61108f82612551565b610b0d8282612559565b5f6110a4818361261a565b92915050565b5f6110bc89898a8a8a8a8a8a8a610a3e565b9998505050505050505050565b5f6110d2612625565b505f516020614b365f395f51905f525b90565b6110ed61231d565b610b0d828261266e565b5f6110a481836126dd565b61110a61231d565b610a3b816126fe565b60ca54604051636e908ca360e01b81526001600160a01b038481166004830152602482018490525f928392839283921690636e908ca390604401608060405180830381865afa158015611168573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061118c919061493f565b9299919850965090945092505050565b60605f6111a85f612766565b9050806001600160401b038111156111c2576111c2613987565b6040519080825280602002602001820160405280156111eb578160200160208202803683370190505b5091506111f75f61276f565b91505090565b61123d6040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b8282101561134f578382905f5260205f200180546112c49061472a565b80601f01602080910402602001604051908101604052809291908181526020018280546112f09061472a565b801561133b5780601f106113125761010080835404028352916020019161133b565b820191905f5260205f20905b81548152906001019060200180831161131e57829003601f168201915b5050505050815260200190600101906112a7565b505050915250909392505050565b61136a6103cd609761276f565b610a3b81611468565b61137b61231d565b610e7a5f61277b565b61138c61231d565b610a3b816127eb565b61139d61231d565b610e7a61286b565b5f60608082808083815f516020614b165f395f51905f5280549091501580156113d057506001810154155b6114145760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a78565b61141c6128b3565b611424612973565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61145661231d565b60ca80546001600160a01b0319169055565b5f5b8151811015610b0d57611497828281518110611488576114886145ff565b60200260200101516001611c3f565b60010161146a565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa1580156114e6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b21919061497c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561154e5750825b90505f826001600160401b031660011480156115695750303b155b905081158015611577575080155b156115955760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156115bf57845460ff60401b1916600160401b1785555b6001600160a01b0389166115ff576040516313fe523b60e11b81526020600482015260066024820152655f63726f737360d01b6044820152606401610a78565b61160a8888886129b1565b60fc80546001600160a01b0319166001600160a01b038b161790556116308960016110e5565b831561167657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60fc545f90610e4e906001600160a01b031688888885808a8a8a610a3e565b5f6116a9611d08565b6116b1611d38565b6116bc60cf83612a74565b82906116de5760405163473978bf60e01b8152600401610a7891815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156117f1578382905f5260205f200180546117669061472a565b80601f01602080910402602001604051908101604052809291908181526020018280546117929061472a565b80156117dd5780601f106117b4576101008083540402835291602001916117dd565b820191905f5260205f20905b8154815290600101906020018083116117c057829003601f168201915b505050505081526020019060010190611749565b505050508152505090505f5f611814836020015184604001518560600151612108565b915091508181906118385760405162461bcd60e51b8152600401610a7891906144d4565b5061184460cf86612a8b565b505f85815260ce6020526040812061185b916138cf565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906118a06004830182613906565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373866060015142604051611900929190918252602082015260400190565b60405180910390a46001935050505061192560015f516020614b765f395f51905f5255565b919050565b61193261231d565b60cb80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610b216097612766565b61196761231d565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f6119ba826110f7565b80156110a45750506001600160a01b03165f9081526034602052604090205460ff161590565b5f805b8251811015611a1757611a0e838281518110611a0157611a016145ff565b60200260200101516116a0565b506001016119e3565b50600192915050565b60605f611a2b61119c565b90505f81516001600160401b03811115611a4757611a47613987565b604051908082528060200260200182016040528015611a70578160200160208202803683370190505b5090505f5b8251811015611aff5760335f848381518110611a9357611a936145ff565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b0316828281518110611adf57611adf6145ff565b6001600160a01b0390921660209283029190910190910152600101611a75565b5092915050565b5f6110a460978361261a565b5f611b21868787878787611681565b9695505050505050565b5f610b215f612766565b5f6065546001610b219190614627565b611b4d61231d565b6001600160a01b038116611b7657604051631e4fbdf760e01b81525f6004820152602401610a78565b610a3b8161277b565b6060610b21609761276f565b5f610b21612a96565b5f6110a46097836126dd565b5f81815260ce60205260409020805460609190611bbc9061472a565b80601f0160208091040260200160405190810160405280929190818152602001828054611be89061472a565b8015611c335780601f10611c0a57610100808354040283529160200191611c33565b820191905f5260205f20905b815481529060010190602001808311611c1657829003601f168201915b50505050509050919050565b611c4761231d565b8015611c8957611c58609783612a9f565b8290611c835760405163082cdf5560e21b81526001600160a01b039091166004820152602401610a78565b50611cc1565b611c94609783612ab3565b8290611cbf5760405163e491024560e01b81526001600160a01b039091166004820152602401610a78565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614b565f395f51905f525460ff1615610e7a5760405163d93c066560e01b815260040160405180910390fd5b5f516020614b765f395f51905f52805460011901611d6957604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f9081906001600160a01b0316611d8d57505f905080611e86565b60ca54604051636e908ca360e01b81526001600160a01b038881166004830152602482018890525f928392911690636e908ca390604401608060405180830381865afa158015611ddf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e03919061493f565b919650945092509050808710801590611e1c5750838610155b8015611e285750828510155b8015611e315750815b8185858a8a8a909192939495611e7d576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a78565b50505050505050505b94509492505050565b83611e9a8688614627565b611ea49190614627565b8360400151101589879091611edd57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610a78565b5050611ee883612ac7565b61167689898989898960018989612471565b60015f516020614b765f395f51905f5255565b8251825181148015611f1f5750815181145b835183518392611f53576040516313a27e1960e21b8152600481019390935260248301919091526044820152606401610a78565b505060995482915060ff16811015611f815760405163fedb0cbd60e01b8152600401610a7891815260200190565b505f80826001600160401b03811115611f9c57611f9c613987565b604051908082528060200260200182016040528015611fc5578160200160208202803683370190505b5090505f5b838110156120d2575f612035888381518110611fe857611fe86145ff565b6020026020010151888481518110612002576120026145ff565b602002602001015188858151811061201c5761201c6145ff565b602002602001015161202d8d612bcb565b929190612bf7565b905061204081611b94565b819061206b576040516338905e7160e11b81526001600160a01b039091166004820152602401610a78565b505f805b84518110156120bb5784818151811061208a5761208a6145ff565b60200260200101516001600160a01b0316836001600160a01b0316036120b357600191506120bb565b60010161206f565b50806120c8578460010194505b5050600101611fca565b50609954829060ff168110156120fe5760405163fedb0cbd60e01b8152600401610a7891815260200190565b5050505050505050565b5f6060825f0361212a57505060408051602081019091525f8152600190612303565b5f196001600160a01b038616016121665761214e6001600160a01b03851684612c23565b505060408051602081019091525f8152600190612303565b60fc546001600160a01b039081169086160361218a57612187606484614993565b92505b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820185905286169063a9059cbb906044016020604051808303815f875af19250505080156121f4575060408051601f3d908101601f191682019092526121f1918101906149b2565b60015b6122a5576122006149cb565b806308c379a00361222957506122146149e3565b8061221f575061226c565b5f92509050612303565b634e487b710361226c5761223b614a65565b90612246575061226c565b60408051602081018390525f945001604051602081830303815290604052915050612303565b3d808015612295576040519150601f19603f3d011682016040523d82523d5f602084013e61229a565b606091505b505f92509050612303565b80156122c5576001925060405180602001604052805f8152509150612301565b5f92506040518060400160405280601f81526020017f427269646765457468657265756d3a207472616e73666572206661696c65640081525091505b505b935093915050565b5f6123168383612cb5565b9392505050565b3361234f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610e7a5760405163118cdaa760e01b8152336004820152602401610a78565b612381816110f7565b80156123a457506001600160a01b0381165f9081526034602052604090205460ff165b81906123cf576040516340da71e560e01b81526001600160a01b039091166004820152602401610a78565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b612420612d01565b5f516020614b565f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016119a5565b612486898988612481888a614627565b612d30565b61249f612491611b35565b8a8a8a8a8a8a8a8a8a612dff565b611676606580546001019055565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061253357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166125275f516020614b365f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e7a5760405163703e46dd60e11b815260040160405180910390fd5b610a3b61231d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156125b3575060408051601f3d908101601f191682019092526125b09181019061497c565b60015b6125db57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a78565b5f516020614b365f395f51905f52811461260b57604051632a87526960e21b815260048101829052602401610a78565b6126158383612ec5565b505050565b5f6123168383612f1a565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e7a5760405163703e46dd60e11b815260040160405180910390fd5b61267782612f40565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515612316565b61270781612fed565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91016126d1565b5f6110a4825490565b60605f6123168361309a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6127f4816119b0565b819061281f576040516340da71e560e01b81526001600160a01b039091166004820152602401610a78565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b612873611d08565b5f516020614b565f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612459565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614b165f395f51905f52916128f19061472a565b80601f016020809104026020016040519081016040528092919081815260200182805461291d9061472a565b80156129685780601f1061293f57610100808354040283529160200191612968565b820191905f5260205f20905b81548152906001019060200180831161294b57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614b165f395f51905f52916128f19061472a565b6129b96130f2565b6129c3338461313b565b6129cb6131ab565b6129d36131b3565b6129db6131c3565b6001600160a01b038216612a2157604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a78565b4360cc556001600160a01b03811615612a505760ca80546001600160a01b0319166001600160a01b0383161790555b5060cb80546001600160a01b0319166001600160a01b039290921691909117905550565b5f8181526001830160205260408120541515612316565b5f61231683836131d3565b5f610b216132b6565b5f612316836001600160a01b038416612cb5565b5f612316836001600160a01b0384166131d3565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612b73576040513d5f823e3d81fd5b50505f513d91508115612b8a578060011415612b98565b84516001600160a01b03163b155b15612bc4578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610a78565b5050505050565b5f6110a4612bd7612a96565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612c0788888888613329565b925092509250612c1782826133f1565b50909695505050505050565b80471015612c4d5760405163cf47918160e01b815247600482015260248101829052604401610a78565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114612c97576040519150601f19603f3d011682016040523d82523d5f602084013e612c9c565b606091505b509150915081612caf57612caf816134a9565b50505050565b5f818152600183016020526040812054612cfa57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556110a4565b505f6110a4565b5f516020614b565f395f51905f525460ff16610e7a57604051638dfc202b60e01b815260040160405180910390fd5b5f196001600160a01b03851601612db257612d4b8183614627565b3414612d578284614627565b349091612d805760405163cf67d7b560e01b815260048101929092526024820152604401610a78565b50508015612dad57612dad81612d9e60cb546001600160a01b031690565b6001600160a01b031690612c23565b612caf565b612dd38330612dc18486614627565b6001600160a01b0388169291906134d2565b8015612caf57612caf612dee60cb546001600160a01b031690565b6001600160a01b0386169083613539565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a604051612e5a9796959493929190614a82565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051612eb1929190918252602082015260400190565b60405180910390a450505050505050505050565b612ece8261356a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612f125761261582826135cd565b610b0d61363f565b5f825f018281548110612f2f57612f2f6145ff565b905f5260205f200154905092915050565b806001600160a01b038116612f80576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a78565b612f8a5f83612a9f565b8290612fb5576040516351eccfe160e11b81526001600160a01b039091166004820152602401610a78565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b03811661302d576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a78565b6130375f83612ab3565b8290613062576040516340da71e560e01b81526001600160a01b039091166004820152602401610a78565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611c3357602002820191905f5260205f20905b8154815260200190600101908083116130d35750505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e7a57604051631afcd79f60e31b815260040160405180910390fd5b6131436130f2565b61314c8261365e565b613194604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061366f565b6099805460ff191660ff9290921691909117905550565b610e7a6130f2565b6131bb6130f2565b610e7a613681565b6131cb6130f2565b610e7a6136a1565b5f81815260018301602052604081205480156132ad575f6131f5600183614ac4565b85549091505f9061320890600190614ac4565b9050808214613267575f865f018281548110613226576132266145ff565b905f5260205f200154905080875f018481548110613246576132466145ff565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061327857613278614ad7565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506110a4565b5f9150506110a4565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6132e06136a9565b6132e8613711565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561336257505f915060039050826133e7565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156133b3573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166133de57505f9250600191508290506133e7565b92505f91508190505b9450945094915050565b5f82600381111561340457613404614aeb565b0361340d575050565b600182600381111561342157613421614aeb565b0361343f5760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561345357613453614aeb565b036134745760405163fce698f760e01b815260048101829052602401610a78565b600382600381111561348857613488614aeb565b03610b0d576040516335e2f38360e21b815260048101829052602401610a78565b8051156134b95780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040516001600160a01b038481166024830152838116604483015260648201839052612caf9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613753565b6040516001600160a01b0383811660248301526044820183905261261591859182169063a9059cbb90606401613507565b806001600160a01b03163b5f0361359f57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a78565b5f516020614b365f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516135e99190614aff565b5f60405180830381855af49150503d805f8114613621576040519150601f19603f3d011682016040523d82523d5f602084013e613626565b606091505b50915091506136368583836137bf565b95945050505050565b3415610e7a5760405163b398979f60e01b815260040160405180910390fd5b6136666130f2565b610a3b81613814565b6136776130f2565b610b0d828261381c565b6136896130f2565b5f516020614b565f395f51905f52805460ff19169055565b611efa6130f2565b5f5f516020614b165f395f51905f52816136c16128b3565b8051909150156136d957805160209091012092915050565b815480156136e8579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614b165f395f51905f5281613729612973565b80519091501561374157805160209091012092915050565b600182015480156136e8579392505050565b5f5f60205f8451602086015f885af180613772576040513d5f823e3d81fd5b50505f513d91508115613789578060011415613796565b6001600160a01b0384163b155b15612caf57604051635274afe760e01b81526001600160a01b0385166004820152602401610a78565b6060826137d4576137cf826134a9565b612316565b81511580156137eb57506001600160a01b0384163b155b15611aff57604051639996b31560e01b81526001600160a01b0385166004820152602401610a78565b611b4d6130f2565b6138246130f2565b5f516020614b165f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10261385d84826147a6565b506003810161386c83826147a6565b505f8082556001909101555050565b828054828255905f5260205f209081019282156138bf579160200282015b828111156138bf57825182906138af90826147a6565b5091602001919060010190613899565b506138cb92915061391d565b5090565b5080546138db9061472a565b5f825580601f106138ea575050565b601f0160209004905f5260205f2090810190610a3b9190613939565b5080545f8255905f5260205f2090810190610a3b91905b808211156138cb575f61393082826138cf565b5060010161391d565b5b808211156138cb575f815560010161393a565b6001600160a01b0381168114610a3b575f5ffd5b80356119258161394d565b5f6020828403121561397c575f5ffd5b81356123168161394d565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b03821117156139ba576139ba613987565b60405250565b601f8201601f191681016001600160401b03811182821017156139e5576139e5613987565b6040525050565b803560ff81168114611925575f5ffd5b5f60e08284031215613a0c575f5ffd5b604051613a188161399b565b8091508235613a268161394d565b81526020830135613a368161394d565b60208201526040838101359082015260608084013590820152613a5b608084016139ec565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f83601f840112613a8b575f5ffd5b5081356001600160401b03811115613aa1575f5ffd5b6020830191508360208260051b8501011115613abb575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613adb575f5ffd5b8935613ae68161394d565b985060208a0135613af68161394d565b975060408a0135613b068161394d565b965060608a0135955060808a0135945060a08a01359350613b2a8b60c08c016139fc565b92506101a08a01356001600160401b03811115613b45575f5ffd5b613b518c828d01613a7b565b915080935050809150509295985092959850929598565b5f6001600160401b03821115613b8057613b80613987565b5060051b60200190565b5f60208284031215613b9a575f5ffd5b81356001600160401b03811115613baf575f5ffd5b8201601f81018413613bbf575f5ffd5b8035613bca81613b68565b604051613bd782826139c0565b80915082815260208101915060208360051b850101925086831115613bfa575f5ffd5b6020840193505b82841015611b21578335613c148161394d565b825260209384019390910190613c01565b5f82601f830112613c34575f5ffd5b8135613c3f81613b68565b604051613c4c82826139c0565b80915082815260208101915060208360051b860101925085831115613c6f575f5ffd5b602085015b83811015613c9357613c85816139ec565b835260209283019201613c74565b5095945050505050565b5f82601f830112613cac575f5ffd5b8135613cb781613b68565b604051613cc482826139c0565b80915082815260208101915060208360051b860101925085831115613ce7575f5ffd5b602085015b83811015613c93578035835260209283019201613cec565b5f5f5f5f5f5f5f5f5f6101008a8c031215613d1d575f5ffd5b89359850613d2d60208b01613961565b9750613d3b60408b01613961565b965060608a0135955060808a01356001600160401b03811115613d5c575f5ffd5b613d688c828d01613a7b565b90965094505060a08a01356001600160401b03811115613d86575f5ffd5b613d928c828d01613c25565b93505060c08a01356001600160401b03811115613dad575f5ffd5b613db98c828d01613c9d565b92505060e08a01356001600160401b03811115613dd4575f5ffd5b613de08c828d01613c9d565b9150509295985092959850929598565b5f5f5f5f5f5f60a08789031215613e05575f5ffd5b8635613e108161394d565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613e3f575f5ffd5b613e4b89828a01613a7b565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613e73575f5ffd5b8735613e7e8161394d565b96506020880135613e8e8161394d565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115613ebd575f5ffd5b613ec98a828b01613a7b565b989b979a50959850939692959293505050565b5f82601f830112613eeb575f5ffd5b8135613ef681613b68565b604051613f0382826139c0565b80915082815260208101915060208360051b860101925085831115613f26575f5ffd5b602085015b83811015613c935780356001600160401b03811115613f48575f5ffd5b613f57886020838a0101613c9d565b84525060209283019201613f2b565b5f5f5f5f5f60808688031215613f7a575f5ffd5b85356001600160401b03811115613f8f575f5ffd5b613f9b88828901613a7b565b90965094505060208601356001600160401b03811115613fb9575f5ffd5b8601601f81018813613fc9575f5ffd5b8035613fd481613b68565b604051613fe182826139c0565b80915082815260208101915060208360051b85010192508a831115614004575f5ffd5b602084015b838110156140445780356001600160401b03811115614026575f5ffd5b6140358d602083890101613c25565b84525060209283019201614009565b50955050505060408601356001600160401b03811115614062575f5ffd5b61406e88828901613edc565b92505060608601356001600160401b03811115614089575f5ffd5b61409588828901613edc565b9150509295509295909350565b5f82601f8301126140b1575f5ffd5b81356001600160401b038111156140ca576140ca613987565b6040516140e1601f8301601f1916602001826139c0565b8181528460208386010111156140f5575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f60408385031215614122575f5ffd5b823561412d8161394d565b915060208301356001600160401b03811115614147575f5ffd5b614153858286016140a2565b9150509250929050565b5f6020828403121561416d575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b03121561418c575f5ffd5b88356141978161394d565b975060208901356141a78161394d565b96506040890135955060608901359450608089013593506141cb8a60a08b016139fc565b92506101808901356001600160401b038111156141e6575f5ffd5b6141f28b828c01613a7b565b999c989b5096995094979396929594505050565b5f5f60408385031215614217575f5ffd5b82356142228161394d565b915060208301356142328161394d565b809150509250929050565b5f5f6040838503121561424e575f5ffd5b82356142598161394d565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156142a75783516001600160a01b0316835260209384019390920191600101614280565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156143795760df198786030183526143648585516142b2565b94506020938401939290920191600101614348565b50929695505050505050565b60ff60f81b8816815260e060208201525f6143a360e08301896142b2565b82810360408401526143b581896142b2565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561440a5783518352602093840193909201916001016143ec565b50909b9a5050505050505050505050565b5f5f5f5f6080858703121561442e575f5ffd5b84356144398161394d565b9350614447602086016139ec565b925060408501356144578161394d565b915060608501356144678161394d565b939692955090935050565b5f5f5f5f5f5f6101608789031215614488575f5ffd5b86356144938161394d565b955060208701356144a38161394d565b9450604087013593506144b988606089016139fc565b92506101408701356001600160401b03811115613e3f575f5ffd5b602081525f61231660208301846142b2565b5f602082840312156144f6575f5ffd5b612316826139ec565b5f6020828403121561450f575f5ffd5b81356001600160401b03811115614524575f5ffd5b8201601f81018413614534575f5ffd5b803561453f81613b68565b60405161454c82826139c0565b80915082815260208101915060208360051b85010192508683111561456f575f5ffd5b6020840193505b82841015611b21578335825260209384019390910190614576565b5f5f5f5f5f61014086880312156145a6575f5ffd5b85356145b18161394d565b9450602086013593506145c787604088016139fc565b92506101208601356001600160401b038111156145e2575f5ffd5b6145ee88828901613a7b565b969995985093965092949392505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808201808211156110a4576110a4614613565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b86811015612c1757838303601f19018852813536879003601e1901811261469e575f5ffd5b86016020810190356001600160401b038111156146b9575f5ffd5b8036038213156146c7575f5ffd5b6146d285828461463a565b60209a8b019a90955093909301925050600101614679565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f906110bc9083018486614662565b600181811c9082168061473e57607f821691505b60208210810361475c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561261557805f5260205f20601f840160051c810160208510156147875750805b601f840160051c820191505b81811015612bc4575f8155600101614793565b81516001600160401b038111156147bf576147bf613987565b6147d3816147cd845461472a565b84614762565b6020601f821160018114614805575f83156147ee5750848201515b5f19600385901b1c1916600184901b178455612bc4565b5f84815260208120601f198516915b828110156148345787850151825560209485019460019092019101614814565b508482101561485157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f61486a83613b68565b60405161487782826139c0565b848152602081019150600585901b840136811115614893575f5ffd5b845b818110156142a75780356001600160401b038111156148b2575f5ffd5b6148be368289016140a2565b85525060209384019301614895565b5f8235609e198336030181126148e1575f5ffd5b9190910192915050565b5f5f8335601e19843603018112614900575f5ffd5b8301803591506001600160401b03821115614919575f5ffd5b6020019150600581901b3603821315613abb575f5ffd5b80518015158114611925575f5ffd5b5f5f5f5f60808587031215614952575f5ffd5b8451602086015160408701519195509350915061497160608601614930565b905092959194509250565b5f6020828403121561498c575f5ffd5b5051919050565b5f826149ad57634e487b7160e01b5f52601260045260245ffd5b500490565b5f602082840312156149c2575f5ffd5b61231682614930565b5f60033d11156110e25760045f5f3e505f5160e01c90565b5f60443d10156149f05790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614a1957505090565b80820180516001600160401b03811115614a34575050505090565b3d8401600319018282016020011115614a4e575050505090565b614a5d602082850101856139c0565b509392505050565b5f5f60233d1115614a7e57602060045f3e50505f516001905b9091565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f906110bc9083018486614662565b818103818111156110a4576110a4614613565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f82518060208501845e5f92019182525091905056fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220d04a2e9de47dee880a3ec18f959ae58202dc388da09511ffa90f11ef92e787c864736f6c634300081c0033",
}

// BridgeEthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeEthereumMetaData.ABI instead.
var BridgeEthereumABI = BridgeEthereumMetaData.ABI

// Deprecated: Use BridgeEthereumMetaData.Sigs instead.
// BridgeEthereumFuncSigs maps the 4-byte function signature to its string representation.
var BridgeEthereumFuncSigs = BridgeEthereumMetaData.Sigs

// BridgeEthereumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeEthereumMetaData.Bin instead.
var BridgeEthereumBin = BridgeEthereumMetaData.Bin

// DeployBridgeEthereum deploys a new Ethereum contract, binding an instance of BridgeEthereum to it.
func DeployBridgeEthereum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeEthereum, error) {
	parsed, err := BridgeEthereumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeEthereumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeEthereum{BridgeEthereumCaller: BridgeEthereumCaller{contract: contract}, BridgeEthereumTransactor: BridgeEthereumTransactor{contract: contract}, BridgeEthereumFilterer: BridgeEthereumFilterer{contract: contract}}, nil
}

// BridgeEthereum is an auto generated Go binding around an Ethereum contract.
type BridgeEthereum struct {
	BridgeEthereumCaller     // Read-only binding to the contract
	BridgeEthereumTransactor // Write-only binding to the contract
	BridgeEthereumFilterer   // Log filterer for contract events
}

// BridgeEthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeEthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeEthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeEthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeEthereumSession struct {
	Contract     *BridgeEthereum   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeEthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeEthereumCallerSession struct {
	Contract *BridgeEthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeEthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeEthereumTransactorSession struct {
	Contract     *BridgeEthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeEthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeEthereumRaw struct {
	Contract *BridgeEthereum // Generic contract binding to access the raw methods on
}

// BridgeEthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeEthereumCallerRaw struct {
	Contract *BridgeEthereumCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeEthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeEthereumTransactorRaw struct {
	Contract *BridgeEthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeEthereum creates a new instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereum(address common.Address, backend bind.ContractBackend) (*BridgeEthereum, error) {
	contract, err := bindBridgeEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereum{BridgeEthereumCaller: BridgeEthereumCaller{contract: contract}, BridgeEthereumTransactor: BridgeEthereumTransactor{contract: contract}, BridgeEthereumFilterer: BridgeEthereumFilterer{contract: contract}}, nil
}

// NewBridgeEthereumCaller creates a new read-only instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumCaller(address common.Address, caller bind.ContractCaller) (*BridgeEthereumCaller, error) {
	contract, err := bindBridgeEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumCaller{contract: contract}, nil
}

// NewBridgeEthereumTransactor creates a new write-only instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeEthereumTransactor, error) {
	contract, err := bindBridgeEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTransactor{contract: contract}, nil
}

// NewBridgeEthereumFilterer creates a new log filterer instance of BridgeEthereum, bound to a specific deployed contract.
func NewBridgeEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeEthereumFilterer, error) {
	contract, err := bindBridgeEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumFilterer{contract: contract}, nil
}

// bindBridgeEthereum binds a generic wrapper to an already deployed contract.
func bindBridgeEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeEthereumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEthereum *BridgeEthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEthereum.Contract.BridgeEthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEthereum *BridgeEthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeEthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEthereum *BridgeEthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeEthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEthereum *BridgeEthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEthereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEthereum *BridgeEthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEthereum *BridgeEthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeEthereum *BridgeEthereumCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeEthereum *BridgeEthereumSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeEthereum.Contract.UPGRADEINTERFACEVERSION(&_BridgeEthereum.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeEthereum *BridgeEthereumCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeEthereum.Contract.UPGRADEINTERFACEVERSION(&_BridgeEthereum.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeEthereum *BridgeEthereumCaller) AllPairs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "allPairs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeEthereum *BridgeEthereumSession) AllPairs() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllPairs(&_BridgeEthereum.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeEthereum *BridgeEthereumCallerSession) AllPairs() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllPairs(&_BridgeEthereum.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeEthereum *BridgeEthereumCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeEthereum *BridgeEthereumSession) AllTokens() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllTokens(&_BridgeEthereum.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeEthereum *BridgeEthereumCallerSession) AllTokens() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllTokens(&_BridgeEthereum.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeEthereum *BridgeEthereumCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeEthereum *BridgeEthereumSession) AllValidators() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllValidators(&_BridgeEthereum.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeEthereum *BridgeEthereumCallerSession) AllValidators() ([]common.Address, error) {
	return _BridgeEthereum.Contract.AllValidators(&_BridgeEthereum.CallOpts)
}

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) BridgeTokenInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "bridgeTokenInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) BridgeTokenInfo() (common.Address, error) {
	return _BridgeEthereum.Contract.BridgeTokenInfo(&_BridgeEthereum.CallOpts)
}

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) BridgeTokenInfo() (common.Address, error) {
	return _BridgeEthereum.Contract.BridgeTokenInfo(&_BridgeEthereum.CallOpts)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeEthereum *BridgeEthereumCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "calculate", token, value)

	outstruct := new(struct {
		Minimum *big.Int
		Gas     *big.Int
		Ex      *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minimum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeEthereum *BridgeEthereumSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeEthereum.Contract.Calculate(&_BridgeEthereum.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeEthereum *BridgeEthereumCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeEthereum.Contract.Calculate(&_BridgeEthereum.CallOpts, token, value)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) Contains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "contains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Contains(token common.Address) (bool, error) {
	return _BridgeEthereum.Contract.Contains(&_BridgeEthereum.CallOpts, token)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) Contains(token common.Address) (bool, error) {
	return _BridgeEthereum.Contract.Contains(&_BridgeEthereum.CallOpts, token)
}

// Cross is a free data retrieval call binding the contract method 0xfa074d03.
//
// Solidity: function cross() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) Cross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "cross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cross is a free data retrieval call binding the contract method 0xfa074d03.
//
// Solidity: function cross() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) Cross() (common.Address, error) {
	return _BridgeEthereum.Contract.Cross(&_BridgeEthereum.CallOpts)
}

// Cross is a free data retrieval call binding the contract method 0xfa074d03.
//
// Solidity: function cross() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) Cross() (common.Address, error) {
	return _BridgeEthereum.Contract.Cross(&_BridgeEthereum.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) Denominator() (*big.Int, error) {
	return _BridgeEthereum.Contract.Denominator(&_BridgeEthereum.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) Denominator() (*big.Int, error) {
	return _BridgeEthereum.Contract.Denominator(&_BridgeEthereum.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumSession) DomainSeparator() ([32]byte, error) {
	return _BridgeEthereum.Contract.DomainSeparator(&_BridgeEthereum.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumCallerSession) DomainSeparator() ([32]byte, error) {
	return _BridgeEthereum.Contract.DomainSeparator(&_BridgeEthereum.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeEthereum *BridgeEthereumCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "eip712Domain")

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
func (_BridgeEthereum *BridgeEthereumSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeEthereum.Contract.Eip712Domain(&_BridgeEthereum.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeEthereum *BridgeEthereumCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeEthereum.Contract.Eip712Domain(&_BridgeEthereum.CallOpts)
}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) GetPairToken(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "getPairToken", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) GetPairToken(token common.Address) (common.Address, error) {
	return _BridgeEthereum.Contract.GetPairToken(&_BridgeEthereum.CallOpts, token)
}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) GetPairToken(token common.Address) (common.Address, error) {
	return _BridgeEthereum.Contract.GetPairToken(&_BridgeEthereum.CallOpts, token)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) InitializedAt() (*big.Int, error) {
	return _BridgeEthereum.Contract.InitializedAt(&_BridgeEthereum.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) InitializedAt() (*big.Int, error) {
	return _BridgeEthereum.Contract.InitializedAt(&_BridgeEthereum.CallOpts)
}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) IsValidToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "isValidToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) IsValidToken(token common.Address) (bool, error) {
	return _BridgeEthereum.Contract.IsValidToken(&_BridgeEthereum.CallOpts, token)
}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) IsValidToken(token common.Address) (bool, error) {
	return _BridgeEthereum.Contract.IsValidToken(&_BridgeEthereum.CallOpts, token)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) IsValidator(validator common.Address) (bool, error) {
	return _BridgeEthereum.Contract.IsValidator(&_BridgeEthereum.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _BridgeEthereum.Contract.IsValidator(&_BridgeEthereum.CallOpts, validator)
}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) NextFinalizeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "nextFinalizeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) NextFinalizeIndex() (*big.Int, error) {
	return _BridgeEthereum.Contract.NextFinalizeIndex(&_BridgeEthereum.CallOpts)
}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) NextFinalizeIndex() (*big.Int, error) {
	return _BridgeEthereum.Contract.NextFinalizeIndex(&_BridgeEthereum.CallOpts)
}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) NextInitiateIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "nextInitiateIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) NextInitiateIndex() (*big.Int, error) {
	return _BridgeEthereum.Contract.NextInitiateIndex(&_BridgeEthereum.CallOpts)
}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) NextInitiateIndex() (*big.Int, error) {
	return _BridgeEthereum.Contract.NextInitiateIndex(&_BridgeEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) Owner() (common.Address, error) {
	return _BridgeEthereum.Contract.Owner(&_BridgeEthereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) Owner() (common.Address, error) {
	return _BridgeEthereum.Contract.Owner(&_BridgeEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Paused() (bool, error) {
	return _BridgeEthereum.Contract.Paused(&_BridgeEthereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeEthereum *BridgeEthereumCallerSession) Paused() (bool, error) {
	return _BridgeEthereum.Contract.Paused(&_BridgeEthereum.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeEthereum.Contract.ProxiableUUID(&_BridgeEthereum.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeEthereum *BridgeEthereumCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeEthereum.Contract.ProxiableUUID(&_BridgeEthereum.CallOpts)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeEthereum *BridgeEthereumCaller) RevertedArguments(opts *bind.CallOpts, index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "revertedArguments", index)

	if err != nil {
		return *new(IBridgeStandardFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeStandardFinalizeArguments)).(*IBridgeStandardFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeEthereum *BridgeEthereumSession) RevertedArguments(index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	return _BridgeEthereum.Contract.RevertedArguments(&_BridgeEthereum.CallOpts, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeEthereum *BridgeEthereumCallerSession) RevertedArguments(index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	return _BridgeEthereum.Contract.RevertedArguments(&_BridgeEthereum.CallOpts, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeEthereum *BridgeEthereumCaller) RevertedReason(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "revertedReason", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeEthereum *BridgeEthereumSession) RevertedReason(index *big.Int) ([]byte, error) {
	return _BridgeEthereum.Contract.RevertedReason(&_BridgeEthereum.CallOpts, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeEthereum *BridgeEthereumCallerSession) RevertedReason(index *big.Int) ([]byte, error) {
	return _BridgeEthereum.Contract.RevertedReason(&_BridgeEthereum.CallOpts, index)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) RewardWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "rewardWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) RewardWallet() (common.Address, error) {
	return _BridgeEthereum.Contract.RewardWallet(&_BridgeEthereum.CallOpts)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) RewardWallet() (common.Address, error) {
	return _BridgeEthereum.Contract.RewardWallet(&_BridgeEthereum.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeEthereum *BridgeEthereumCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeEthereum *BridgeEthereumSession) Threshold() (uint8, error) {
	return _BridgeEthereum.Contract.Threshold(&_BridgeEthereum.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeEthereum *BridgeEthereumCallerSession) Threshold() (uint8, error) {
	return _BridgeEthereum.Contract.Threshold(&_BridgeEthereum.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) TokenByIndex(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "tokenByIndex", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeEthereum.Contract.TokenByIndex(&_BridgeEthereum.CallOpts, i)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeEthereum.Contract.TokenByIndex(&_BridgeEthereum.CallOpts, i)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) TokensLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.TokensLength(&_BridgeEthereum.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) TokensLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.TokensLength(&_BridgeEthereum.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BridgeEthereum.Contract.ValidatorByIndex(&_BridgeEthereum.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BridgeEthereum.Contract.ValidatorByIndex(&_BridgeEthereum.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) ValidatorLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.ValidatorLength(&_BridgeEthereum.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) ValidatorLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.ValidatorLength(&_BridgeEthereum.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) AddToken(opts *bind.TransactOpts, token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "addToken", token, pair)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeEthereum *BridgeEthereumSession) AddToken(token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.AddToken(&_BridgeEthereum.TransactOpts, token, pair)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) AddToken(token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.AddToken(&_BridgeEthereum.TransactOpts, token, pair)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) Bridge(opts *bind.TransactOpts, token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridge", token, value, gas, ex, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Bridge(token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Bridge(&_BridgeEthereum.TransactOpts, token, value, gas, ex, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) Bridge(token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Bridge(&_BridgeEthereum.TransactOpts, token, value, gas, ex, extraData)
}

// BridgeCross is a paid mutator transaction binding the contract method 0xd86c62dc.
//
// Solidity: function bridgeCross(address account, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeCross(opts *bind.TransactOpts, account common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeCross", account, value, permitArgs, extraData)
}

// BridgeCross is a paid mutator transaction binding the contract method 0xd86c62dc.
//
// Solidity: function bridgeCross(address account, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeCross(account common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeCross(&_BridgeEthereum.TransactOpts, account, value, permitArgs, extraData)
}

// BridgeCross is a paid mutator transaction binding the contract method 0xd86c62dc.
//
// Solidity: function bridgeCross(address account, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeCross(account common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeCross(&_BridgeEthereum.TransactOpts, account, value, permitArgs, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeTo(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeTo", token, to, value, gas, ex, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeTo(&_BridgeEthereum.TransactOpts, token, to, value, gas, ex, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeTo(&_BridgeEthereum.TransactOpts, token, to, value, gas, ex, extraData)
}

// BridgeToCross is a paid mutator transaction binding the contract method 0xa01be2c7.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeToCross(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeToCross", from, to, value, permitArgs, extraData)
}

// BridgeToCross is a paid mutator transaction binding the contract method 0xa01be2c7.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeToCross(from common.Address, to common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeToCross(&_BridgeEthereum.TransactOpts, from, to, value, permitArgs, extraData)
}

// BridgeToCross is a paid mutator transaction binding the contract method 0xa01be2c7.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeToCross(from common.Address, to common.Address, value *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeToCross(&_BridgeEthereum.TransactOpts, from, to, value, permitArgs, extraData)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) ChangeRewardWallet(opts *bind.TransactOpts, rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "changeRewardWallet", rewardWallet_)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeEthereum *BridgeEthereumSession) ChangeRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ChangeRewardWallet(&_BridgeEthereum.TransactOpts, rewardWallet_)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) ChangeRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ChangeRewardWallet(&_BridgeEthereum.TransactOpts, rewardWallet_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeEthereum *BridgeEthereumSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ChangeThreshold(&_BridgeEthereum.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ChangeThreshold(&_BridgeEthereum.TransactOpts, threshold_)
}

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) Finalize(opts *bind.TransactOpts, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "finalize", index, token, to, value, extraData, v, r, s)
}

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Finalize(&_BridgeEthereum.TransactOpts, index, token, to, value, extraData, v, r, s)
}

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Finalize(&_BridgeEthereum.TransactOpts, index, token, to, value, extraData, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) FinalizeBatch(opts *bind.TransactOpts, args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "finalizeBatch", args, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.FinalizeBatch(&_BridgeEthereum.TransactOpts, args, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.FinalizeBatch(&_BridgeEthereum.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _cross, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) Initialize(opts *bind.TransactOpts, _cross common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "initialize", _cross, _threshold, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _cross, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumSession) Initialize(_cross common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Initialize(&_BridgeEthereum.TransactOpts, _cross, _threshold, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _cross, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) Initialize(_cross common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Initialize(&_BridgeEthereum.TransactOpts, _cross, _threshold, rewardWallet_, _bridgeTokenInfo)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEthereum *BridgeEthereumSession) Pause() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Pause(&_BridgeEthereum.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Pause(&_BridgeEthereum.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) PauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "pauseToken", token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PauseToken(&_BridgeEthereum.TransactOpts, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PauseToken(&_BridgeEthereum.TransactOpts, token)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridge", token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, ex, permitArgs, extraData)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveToken(&_BridgeEthereum.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveToken(&_BridgeEthereum.TransactOpts, token)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RemoveTokenInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "removeTokenInfo")
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeEthereum *BridgeEthereumSession) RemoveTokenInfo() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveTokenInfo(&_BridgeEthereum.TransactOpts)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RemoveTokenInfo() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveTokenInfo(&_BridgeEthereum.TransactOpts)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveValidator(&_BridgeEthereum.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveValidator(&_BridgeEthereum.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveValidators(&_BridgeEthereum.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RemoveValidators(&_BridgeEthereum.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RenounceOwnership(&_BridgeEthereum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RenounceOwnership(&_BridgeEthereum.TransactOpts)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ResetValidators(&_BridgeEthereum.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.ResetValidators(&_BridgeEthereum.TransactOpts, validators)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) RetryFinalize(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "retryFinalize", index)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) RetryFinalize(index *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RetryFinalize(&_BridgeEthereum.TransactOpts, index)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) RetryFinalize(index *big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RetryFinalize(&_BridgeEthereum.TransactOpts, index)
}

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) RetryFinalizeBatch(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "retryFinalizeBatch", indexes)
}

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) RetryFinalizeBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RetryFinalizeBatch(&_BridgeEthereum.TransactOpts, indexes)
}

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) RetryFinalizeBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.RetryFinalizeBatch(&_BridgeEthereum.TransactOpts, indexes)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) SetTokenInfo(opts *bind.TransactOpts, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "setTokenInfo", _bridgeTokenInfo)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumSession) SetTokenInfo(_bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetTokenInfo(&_BridgeEthereum.TransactOpts, _bridgeTokenInfo)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) SetTokenInfo(_bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetTokenInfo(&_BridgeEthereum.TransactOpts, _bridgeTokenInfo)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetValidator(&_BridgeEthereum.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetValidator(&_BridgeEthereum.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetValidators(&_BridgeEthereum.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.SetValidators(&_BridgeEthereum.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.TransferOwnership(&_BridgeEthereum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.TransferOwnership(&_BridgeEthereum.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEthereum *BridgeEthereumTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEthereum *BridgeEthereumSession) Unpause() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Unpause(&_BridgeEthereum.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Unpause(&_BridgeEthereum.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) UnpauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "unpauseToken", token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UnpauseToken(&_BridgeEthereum.TransactOpts, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UnpauseToken(&_BridgeEthereum.TransactOpts, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeEthereum *BridgeEthereumTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeEthereum *BridgeEthereumSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UpgradeToAndCall(&_BridgeEthereum.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.UpgradeToAndCall(&_BridgeEthereum.TransactOpts, newImplementation, data)
}

// BridgeEthereumBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFeeChargedIterator struct {
	Event *BridgeEthereumBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgeFeeCharged)
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
		it.Event = new(BridgeEthereumBridgeFeeCharged)
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
func (it *BridgeEthereumBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgeFeeCharged represents a BridgeFeeCharged event raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFeeCharged struct {
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
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*BridgeEthereumBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgeFeeChargedIterator{contract: _BridgeEthereum.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgeFeeCharged)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgeFeeCharged(log types.Log) (*BridgeEthereumBridgeFeeCharged, error) {
	event := new(BridgeEthereumBridgeFeeCharged)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumBridgeFinalizeRevertedIterator is returned from FilterBridgeFinalizeReverted and is used to iterate over the raw logs and unpacked data for BridgeFinalizeReverted events raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFinalizeRevertedIterator struct {
	Event *BridgeEthereumBridgeFinalizeReverted // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgeFinalizeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgeFinalizeReverted)
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
		it.Event = new(BridgeEthereumBridgeFinalizeReverted)
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
func (it *BridgeEthereumBridgeFinalizeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgeFinalizeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgeFinalizeReverted represents a BridgeFinalizeReverted event raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFinalizeReverted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeReverted is a free log retrieval operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgeFinalizeReverted(opts *bind.FilterOpts, index []*big.Int) (*BridgeEthereumBridgeFinalizeRevertedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgeFinalizeRevertedIterator{contract: _BridgeEthereum.contract, event: "BridgeFinalizeReverted", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeReverted is a free log subscription operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgeFinalizeReverted(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgeFinalizeReverted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgeFinalizeReverted)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgeFinalizeReverted(log types.Log) (*BridgeEthereumBridgeFinalizeReverted, error) {
	event := new(BridgeEthereumBridgeFinalizeReverted)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFinalizedIterator struct {
	Event *BridgeEthereumBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgeFinalized)
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
		it.Event = new(BridgeEthereumBridgeFinalized)
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
func (it *BridgeEthereumBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgeFinalized represents a BridgeFinalized event raised by the BridgeEthereum contract.
type BridgeEthereumBridgeFinalized struct {
	Index *big.Int
	Token common.Address
	To    common.Address
	Value *big.Int
	Time  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, index []*big.Int, token []common.Address, to []common.Address) (*BridgeEthereumBridgeFinalizedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgeFinalizedIterator{contract: _BridgeEthereum.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgeFinalized, index []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgeFinalized)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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

// ParseBridgeFinalized is a log parse operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgeFinalized(log types.Log) (*BridgeEthereumBridgeFinalized, error) {
	event := new(BridgeEthereumBridgeFinalized)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BridgeEthereum contract.
type BridgeEthereumBridgeInitiatedIterator struct {
	Event *BridgeEthereumBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumBridgeInitiated)
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
		it.Event = new(BridgeEthereumBridgeInitiated)
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
func (it *BridgeEthereumBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumBridgeInitiated represents a BridgeInitiated event raised by the BridgeEthereum contract.
type BridgeEthereumBridgeInitiated struct {
	Index     *big.Int
	Token     common.Address
	PairToken common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	Time      *big.Int
	Permit    bool
	ExtraData [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe854093.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes[] extraData)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, index []*big.Int, token []common.Address, from []common.Address) (*BridgeEthereumBridgeInitiatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "BridgeInitiated", indexRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumBridgeInitiatedIterator{contract: _BridgeEthereum.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe854093.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes[] extraData)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BridgeEthereumBridgeInitiated, index []*big.Int, token []common.Address, from []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "BridgeInitiated", indexRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumBridgeInitiated)
				if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe854093.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes[] extraData)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseBridgeInitiated(log types.Log) (*BridgeEthereumBridgeInitiated, error) {
	event := new(BridgeEthereumBridgeInitiated)
	if err := _BridgeEthereum.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BridgeEthereum contract.
type BridgeEthereumEIP712DomainChangedIterator struct {
	Event *BridgeEthereumEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumEIP712DomainChanged)
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
		it.Event = new(BridgeEthereumEIP712DomainChanged)
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
func (it *BridgeEthereumEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumEIP712DomainChanged represents a EIP712DomainChanged event raised by the BridgeEthereum contract.
type BridgeEthereumEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeEthereum *BridgeEthereumFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BridgeEthereumEIP712DomainChangedIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumEIP712DomainChangedIterator{contract: _BridgeEthereum.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeEthereum *BridgeEthereumFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BridgeEthereumEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumEIP712DomainChanged)
				if err := _BridgeEthereum.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseEIP712DomainChanged(log types.Log) (*BridgeEthereumEIP712DomainChanged, error) {
	event := new(BridgeEthereumEIP712DomainChanged)
	if err := _BridgeEthereum.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeEthereum contract.
type BridgeEthereumInitializedIterator struct {
	Event *BridgeEthereumInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumInitialized)
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
		it.Event = new(BridgeEthereumInitialized)
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
func (it *BridgeEthereumInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumInitialized represents a Initialized event raised by the BridgeEthereum contract.
type BridgeEthereumInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeEthereumInitializedIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumInitializedIterator{contract: _BridgeEthereum.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeEthereumInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumInitialized)
				if err := _BridgeEthereum.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseInitialized(log types.Log) (*BridgeEthereumInitialized, error) {
	event := new(BridgeEthereumInitialized)
	if err := _BridgeEthereum.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeEthereum contract.
type BridgeEthereumOwnershipTransferredIterator struct {
	Event *BridgeEthereumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumOwnershipTransferred)
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
		it.Event = new(BridgeEthereumOwnershipTransferred)
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
func (it *BridgeEthereumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeEthereum contract.
type BridgeEthereumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeEthereumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumOwnershipTransferredIterator{contract: _BridgeEthereum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeEthereumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumOwnershipTransferred)
				if err := _BridgeEthereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeEthereumOwnershipTransferred, error) {
	event := new(BridgeEthereumOwnershipTransferred)
	if err := _BridgeEthereum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumPairUpdatedIterator is returned from FilterPairUpdated and is used to iterate over the raw logs and unpacked data for PairUpdated events raised by the BridgeEthereum contract.
type BridgeEthereumPairUpdatedIterator struct {
	Event *BridgeEthereumPairUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumPairUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumPairUpdated)
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
		it.Event = new(BridgeEthereumPairUpdated)
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
func (it *BridgeEthereumPairUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumPairUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumPairUpdated represents a PairUpdated event raised by the BridgeEthereum contract.
type BridgeEthereumPairUpdated struct {
	Token      common.Address
	Pair       common.Address
	Registered bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPairUpdated is a free log retrieval operation binding the contract event 0xa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce.
//
// Solidity: event PairUpdated(address indexed token, address indexed pair, bool registered)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterPairUpdated(opts *bind.FilterOpts, token []common.Address, pair []common.Address) (*BridgeEthereumPairUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "PairUpdated", tokenRule, pairRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumPairUpdatedIterator{contract: _BridgeEthereum.contract, event: "PairUpdated", logs: logs, sub: sub}, nil
}

// WatchPairUpdated is a free log subscription operation binding the contract event 0xa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce.
//
// Solidity: event PairUpdated(address indexed token, address indexed pair, bool registered)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchPairUpdated(opts *bind.WatchOpts, sink chan<- *BridgeEthereumPairUpdated, token []common.Address, pair []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "PairUpdated", tokenRule, pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumPairUpdated)
				if err := _BridgeEthereum.contract.UnpackLog(event, "PairUpdated", log); err != nil {
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

// ParsePairUpdated is a log parse operation binding the contract event 0xa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce.
//
// Solidity: event PairUpdated(address indexed token, address indexed pair, bool registered)
func (_BridgeEthereum *BridgeEthereumFilterer) ParsePairUpdated(log types.Log) (*BridgeEthereumPairUpdated, error) {
	event := new(BridgeEthereumPairUpdated)
	if err := _BridgeEthereum.contract.UnpackLog(event, "PairUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeEthereum contract.
type BridgeEthereumPausedIterator struct {
	Event *BridgeEthereumPaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumPaused)
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
		it.Event = new(BridgeEthereumPaused)
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
func (it *BridgeEthereumPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumPaused represents a Paused event raised by the BridgeEthereum contract.
type BridgeEthereumPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeEthereumPausedIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumPausedIterator{contract: _BridgeEthereum.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumPaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParsePaused(log types.Log) (*BridgeEthereumPaused, error) {
	event := new(BridgeEthereumPaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BridgeEthereum contract.
type BridgeEthereumThresholdChangedIterator struct {
	Event *BridgeEthereumThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumThresholdChanged)
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
		it.Event = new(BridgeEthereumThresholdChanged)
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
func (it *BridgeEthereumThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumThresholdChanged represents a ThresholdChanged event raised by the BridgeEthereum contract.
type BridgeEthereumThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BridgeEthereumThresholdChangedIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumThresholdChangedIterator{contract: _BridgeEthereum.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeEthereumThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumThresholdChanged)
				if err := _BridgeEthereum.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseThresholdChanged(log types.Log) (*BridgeEthereumThresholdChanged, error) {
	event := new(BridgeEthereumThresholdChanged)
	if err := _BridgeEthereum.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the BridgeEthereum contract.
type BridgeEthereumTokenAddedIterator struct {
	Event *BridgeEthereumTokenAdded // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumTokenAdded)
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
		it.Event = new(BridgeEthereumTokenAdded)
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
func (it *BridgeEthereumTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumTokenAdded represents a TokenAdded event raised by the BridgeEthereum contract.
type BridgeEthereumTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*BridgeEthereumTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTokenAddedIterator{contract: _BridgeEthereum.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeEthereumTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumTokenAdded)
				if err := _BridgeEthereum.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseTokenAdded(log types.Log) (*BridgeEthereumTokenAdded, error) {
	event := new(BridgeEthereumTokenAdded)
	if err := _BridgeEthereum.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumTokenPausedIterator is returned from FilterTokenPaused and is used to iterate over the raw logs and unpacked data for TokenPaused events raised by the BridgeEthereum contract.
type BridgeEthereumTokenPausedIterator struct {
	Event *BridgeEthereumTokenPaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumTokenPaused)
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
		it.Event = new(BridgeEthereumTokenPaused)
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
func (it *BridgeEthereumTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumTokenPaused represents a TokenPaused event raised by the BridgeEthereum contract.
type BridgeEthereumTokenPaused struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenPaused is a free log retrieval operation binding the contract event 0xf017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d.
//
// Solidity: event TokenPaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterTokenPaused(opts *bind.FilterOpts, token []common.Address) (*BridgeEthereumTokenPausedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "TokenPaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTokenPausedIterator{contract: _BridgeEthereum.contract, event: "TokenPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPaused is a free log subscription operation binding the contract event 0xf017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d.
//
// Solidity: event TokenPaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchTokenPaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumTokenPaused, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "TokenPaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumTokenPaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "TokenPaused", log); err != nil {
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

// ParseTokenPaused is a log parse operation binding the contract event 0xf017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d.
//
// Solidity: event TokenPaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseTokenPaused(log types.Log) (*BridgeEthereumTokenPaused, error) {
	event := new(BridgeEthereumTokenPaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "TokenPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the BridgeEthereum contract.
type BridgeEthereumTokenRemovedIterator struct {
	Event *BridgeEthereumTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumTokenRemoved)
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
		it.Event = new(BridgeEthereumTokenRemoved)
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
func (it *BridgeEthereumTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumTokenRemoved represents a TokenRemoved event raised by the BridgeEthereum contract.
type BridgeEthereumTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeEthereumTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTokenRemovedIterator{contract: _BridgeEthereum.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeEthereumTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumTokenRemoved)
				if err := _BridgeEthereum.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseTokenRemoved(log types.Log) (*BridgeEthereumTokenRemoved, error) {
	event := new(BridgeEthereumTokenRemoved)
	if err := _BridgeEthereum.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumTokenUnpausedIterator is returned from FilterTokenUnpaused and is used to iterate over the raw logs and unpacked data for TokenUnpaused events raised by the BridgeEthereum contract.
type BridgeEthereumTokenUnpausedIterator struct {
	Event *BridgeEthereumTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumTokenUnpaused)
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
		it.Event = new(BridgeEthereumTokenUnpaused)
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
func (it *BridgeEthereumTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumTokenUnpaused represents a TokenUnpaused event raised by the BridgeEthereum contract.
type BridgeEthereumTokenUnpaused struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnpaused is a free log retrieval operation binding the contract event 0xf38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf.
//
// Solidity: event TokenUnpaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterTokenUnpaused(opts *bind.FilterOpts, token []common.Address) (*BridgeEthereumTokenUnpausedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "TokenUnpaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumTokenUnpausedIterator{contract: _BridgeEthereum.contract, event: "TokenUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenUnpaused is a free log subscription operation binding the contract event 0xf38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf.
//
// Solidity: event TokenUnpaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchTokenUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumTokenUnpaused, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "TokenUnpaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumTokenUnpaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
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

// ParseTokenUnpaused is a log parse operation binding the contract event 0xf38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf.
//
// Solidity: event TokenUnpaused(address indexed token)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseTokenUnpaused(log types.Log) (*BridgeEthereumTokenUnpaused, error) {
	event := new(BridgeEthereumTokenUnpaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeEthereum contract.
type BridgeEthereumUnpausedIterator struct {
	Event *BridgeEthereumUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumUnpaused)
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
		it.Event = new(BridgeEthereumUnpaused)
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
func (it *BridgeEthereumUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumUnpaused represents a Unpaused event raised by the BridgeEthereum contract.
type BridgeEthereumUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeEthereumUnpausedIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumUnpausedIterator{contract: _BridgeEthereum.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeEthereumUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumUnpaused)
				if err := _BridgeEthereum.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseUnpaused(log types.Log) (*BridgeEthereumUnpaused, error) {
	event := new(BridgeEthereumUnpaused)
	if err := _BridgeEthereum.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeEthereum contract.
type BridgeEthereumUpgradedIterator struct {
	Event *BridgeEthereumUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumUpgraded)
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
		it.Event = new(BridgeEthereumUpgraded)
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
func (it *BridgeEthereumUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumUpgraded represents a Upgraded event raised by the BridgeEthereum contract.
type BridgeEthereumUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeEthereumUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumUpgradedIterator{contract: _BridgeEthereum.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeEthereumUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumUpgraded)
				if err := _BridgeEthereum.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BridgeEthereum *BridgeEthereumFilterer) ParseUpgraded(log types.Log) (*BridgeEthereumUpgraded, error) {
	event := new(BridgeEthereumUpgraded)
	if err := _BridgeEthereum.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEthereumValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the BridgeEthereum contract.
type BridgeEthereumValidatorSetIterator struct {
	Event *BridgeEthereumValidatorSet // Event containing the contract specifics and raw log

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
func (it *BridgeEthereumValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEthereumValidatorSet)
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
		it.Event = new(BridgeEthereumValidatorSet)
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
func (it *BridgeEthereumValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEthereumValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEthereumValidatorSet represents a ValidatorSet event raised by the BridgeEthereum contract.
type BridgeEthereumValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BridgeEthereum *BridgeEthereumFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*BridgeEthereumValidatorSetIterator, error) {

	logs, sub, err := _BridgeEthereum.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &BridgeEthereumValidatorSetIterator{contract: _BridgeEthereum.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BridgeEthereum *BridgeEthereumFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *BridgeEthereumValidatorSet) (event.Subscription, error) {

	logs, sub, err := _BridgeEthereum.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEthereumValidatorSet)
				if err := _BridgeEthereum.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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
// Solidity: event ValidatorSet(address validators, bool status)
func (_BridgeEthereum *BridgeEthereumFilterer) ParseValidatorSet(log types.Log) (*BridgeEthereumValidatorSet, error) {
	event := new(BridgeEthereumValidatorSet)
	if err := _BridgeEthereum.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
