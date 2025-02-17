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
	ABI: "[{\"inputs\":[],\"name\":\"BridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BridgeFeeManagerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFeeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeToCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cross\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"feeInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_cross\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e3a446eb": "BridgeFeeManager()",
		"114ee544": "BridgeFeeManagerLength()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"9d0dc76e": "allFeeInfo()",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"d86c62dc": "bridgeCross(address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"a01be2c7": "bridgeToCross(address,address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"8b28ab1e": "calculateFee(address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
		"fa074d03": "cross()",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"751b4c9c": "feeInfoByIndex(uint256)",
		"f120c400": "finalize(uint256,address,address,uint256,bytes[],bytes[])",
		"008bd028": "finalizeBatch((uint256,address,address,uint256,bytes[])[],bytes[][])",
		"71c59d7b": "getPairToken(address)",
		"252154fa": "getTokenFee(address)",
		"c0c53b8b": "initialize(address,address,address)",
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
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"a9823183": "retryFinalize(uint256)",
		"7021fd0e": "revertedArguments(uint256)",
		"fe2b8da6": "revertedReason(uint256)",
		"fb75b2c7": "rewardWallet()",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516149436100f95f395f81816121110152818161213a015261228901526149435ff3fe60806040526004361061035a575f3560e01c806384b0196e116101bd578063c97682f8116100f2578063f2fde38b11610092578063fa074d031161006d578063fa074d031461098d578063facd743b146109ac578063fb75b2c7146109cb578063fe2b8da6146109e8575f5ffd5b8063f2fde38b14610946578063f30589c314610965578063f698da2514610979575f5ffd5b8063d92fc67b116100cd578063d92fc67b146108ec578063e1af7f5014610900578063e3a446eb14610914578063f120c40014610933575f5ffd5b8063c97682f81461089a578063cbae5958146108ae578063d86c62dc146108cd575f5ffd5b8063a01be2c71161015d578063aed1d40311610138578063aed1d40314610829578063b7f3358d1461083d578063c0c53b8b1461085c578063c18764531461087b575f5ffd5b8063a01be2c7146107ae578063a9823183146107cd578063ad3cb1cc146107ec575f5ffd5b806391cf6d3e1161019857806391cf6d3e146107465780639300c9261461075a57806396ce0795146107795780639d0dc76e1461078d575f5ffd5b806384b0196e146106af5780638b28ab1e146106d65780638da5cb5b1461070a575f5ffd5b80634f6ccce7116102935780636ff97f1d1161023357806371c59d7b1161020e57806371c59d7b14610626578063751b4c9c1461065d5780637c41ad2c1461067c5780638456cb591461069b575f5ffd5b80636ff97f1d146105c55780637021fd0e146105e6578063715018a614610612575f5ffd5b80635476bd721161026e5780635476bd72146105455780635c975abb146105645780635dbe47e8146105875780635fa7b584146105a6575f5ffd5b80634f6ccce7146104e757806351c455791461051e57806352d1902d14610531575f5ffd5b806337d10075116102fe5780633f5d3b5d116102d95780633f5d3b5d1461048157806340a141ff1461049457806342cde4e8146104b35780634f1ef286146104d4575f5ffd5b806337d100751461043b5780633b3bff0f1461044e5780633f4ba83a1461046d575f5ffd5b8063174991ab11610339578063174991ab146103c95780631d40f0d8146103dc578063252154fa146103fb57806327938c8114610427575f5ffd5b80628bd0281461035e578063114ee544146103865780631327d3d8146103a8575b5f5ffd5b61037161036c366004613992565b610a07565b60405190151581526020015b60405180910390f35b348015610391575f5ffd5b5061039a610b33565b60405190815260200161037d565b3480156103b3575f5ffd5b506103c76103c2366004613a91565b610ba3565b005b6103716103d7366004613b3b565b610bb1565b3480156103e7575f5ffd5b506103c76103f6366004613be1565b610c8b565b348015610406575f5ffd5b5061041a610415366004613a91565b610cc5565b60405161037d9190613c7c565b348015610432575f5ffd5b5061039a610d62565b610371610449366004613ca6565b610d72565b348015610459575f5ffd5b506103c7610468366004613a91565b610d8d565b348015610478575f5ffd5b506103c7610d9e565b61037161048f366004613d13565b610db0565b34801561049f575f5ffd5b506103c76104ae366004613a91565b610e89565b3480156104be575f5ffd5b5060995460405160ff909116815260200161037d565b6103c76104e2366004613d92565b610e93565b3480156104f2575f5ffd5b50610506610501366004613dde565b610eae565b6040516001600160a01b03909116815260200161037d565b61037161052c366004613df5565b610eb9565b34801561053c575f5ffd5b5061039a610ed8565b348015610550575f5ffd5b506103c761055f366004613e87565b610ef4565b34801561056f575f5ffd5b505f5160206148ce5f395f51905f525460ff16610371565b348015610592575f5ffd5b506103716105a1366004613a91565b610f06565b3480156105b1575f5ffd5b506103c76105c0366004613a91565b610f11565b3480156105d0575f5ffd5b506105d9610f22565b60405161037d9190613ebe565b3480156105f1575f5ffd5b50610605610600366004613dde565b610f83565b60405161037d9190613f37565b34801561061d575f5ffd5b506103c76110e3565b348015610631575f5ffd5b50610506610640366004613a91565b6001600160a01b039081165f908152603360205260409020541690565b348015610668575f5ffd5b5061041a610677366004613dde565b6110f4565b348015610687575f5ffd5b506103c7610696366004613a91565b61114f565b3480156106a6575f5ffd5b506103c7611160565b3480156106ba575f5ffd5b506106c3611170565b60405161037d9796959493929190613fdc565b3480156106e1575f5ffd5b506106f56106f0366004614072565b611219565b6040805192835260208301919091520161037d565b348015610715575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610506565b348015610751575f5ffd5b5060cb5461039a565b348015610765575f5ffd5b506103c7610774366004613be1565b611299565b348015610784575f5ffd5b5061039a6112d0565b348015610798575f5ffd5b506107a1611317565b60405161037d919061409c565b3480156107b9575f5ffd5b506103716107c83660046140fb565b611385565b3480156107d8575f5ffd5b506103716107e7366004613dde565b6113a4565b3480156107f7575f5ffd5b5061081c604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161037d919061415d565b348015610834575f5ffd5b5061039a611663565b348015610848575f5ffd5b506103c761085736600461416f565b61166e565b348015610867575f5ffd5b506103c7610876366004614188565b6116bf565b348015610886575f5ffd5b50610371610895366004613a91565b6117e9565b3480156108a5575f5ffd5b506105d9611819565b3480156108b9575f5ffd5b506105066108c8366004613dde565b6118ff565b3480156108d8575f5ffd5b506103716108e73660046141d0565b61190b565b3480156108f7575f5ffd5b5061039a611924565b34801561090b575f5ffd5b5061039a61192e565b34801561091f575f5ffd5b5060ca54610506906001600160a01b031681565b61037161094136600461423e565b61193e565b348015610951575f5ffd5b506103c7610960366004613a91565b611c05565b348015610970575f5ffd5b506105d9611c3f565b348015610984575f5ffd5b5061039a611c4b565b348015610998575f5ffd5b5060fe54610506906001600160a01b031681565b3480156109b7575f5ffd5b506103716109c6366004613a91565b611c54565b3480156109d6575f5ffd5b5060cc546001600160a01b0316610506565b3480156109f3575f5ffd5b5061081c610a02366004613dde565b611c60565b5f805b83811015610b2657610b1d858583818110610a2757610a276142dc565b9050602002810190610a3991906142f0565b35868684818110610a4c57610a4c6142dc565b9050602002810190610a5e91906142f0565b610a6f906040810190602001613a91565b878785818110610a8157610a816142dc565b9050602002810190610a9391906142f0565b610aa4906060810190604001613a91565b888886818110610ab657610ab66142dc565b9050602002810190610ac891906142f0565b60600135898987818110610ade57610ade6142dc565b9050602002810190610af091906142f0565b610afe90608081019061430e565b898881518110610b1057610b106142dc565b602002602001015161193e565b50600101610a0a565b50600190505b9392505050565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa158015610b7a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b9e9190614353565b905090565b610bae816001611cff565b50565b5f610bba611dc8565b89610bc4816117e9565b8190610bf457604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610bfd611df8565b5f5f5f610c0c8e8c8c8c611e2f565b9194509250905082828b8b84610c4b576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610beb565b50505050610c608e8e8e8e87878e8e8e611eed565b60019450505050610c7d60015f5160206148ee5f395f51905f5255565b509998505050505050505050565b5f5b8151811015610cc157610cb9828281518110610cab57610cab6142dc565b60200260200101515f611cff565b600101610c8d565b5050565b610cef60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631290aa7d60e11b81526001600160a01b0384811660048301529091169063252154fa906024015b606060405180830381865afa158015610d38573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d5c91906143cc565b92915050565b5f6066546001610b9e91906143fa565b5f610d8287338888888888610db0565b979650505050505050565b610d95611f76565b610bae81611fd1565b610da6611f76565b610dae612071565b565b5f610db9611dc8565b87610dc3816117e9565b8190610dee57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610beb565b50610df7611df8565b5f5f5f610e068c8b8b8b611e2f565b9194509250905082828a8a84610e45576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610beb565b50505050610e608c610e543390565b8d8d87875f8e8e6120ca565b60019450505050610e7d60015f5160206148ee5f395f51905f5255565b50979650505050505050565b610bae815f611cff565b610e9b612106565b610ea4826121aa565b610cc182826121b2565b5f610d5c8183612273565b5f610ecb89898a8a8a8a8a8a8a610bb1565b9998505050505050505050565b5f610ee161227e565b505f5160206148ae5f395f51905f525b90565b610efc611f76565b610cc182826122c7565b5f610d5c8183612336565b610f19611f76565b610bae81612357565b60605f610f2e5f6123bf565b9050806001600160401b03811115610f4857610f48613806565b604051908082528060200260200182016040528015610f71578160200160208202803683370190505b509150610f7d5f6123c8565b91505090565b610fc36040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b828210156110d5578382905f5260205f2001805461104a9061440d565b80601f01602080910402602001604051908101604052809291908181526020018280546110769061440d565b80156110c15780601f10611098576101008083540402835291602001916110c1565b820191905f5260205f20905b8154815290600101906020018083116110a457829003601f168201915b50505050508152602001906001019061102d565b505050915250909392505050565b6110eb611f76565b610dae5f6123d4565b61111e60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631d46d32760e21b8152600481018490526001600160a01b039091169063751b4c9c90602401610d1d565b611157611f76565b610bae81612444565b611168611f76565b610dae6124c4565b5f60608082808083815f51602061488e5f395f51905f52805490915015801561119b57506001810154155b6111df5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610beb565b6111e761250c565b6111ef6125cc565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60ca54604051634594558f60e11b81526001600160a01b038481166004830152602482018490525f928392911690638b28ab1e906044016040805180830381865afa15801561126a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061128e9190614445565b909590945092505050565b5f5b8151811015610cc1576112c88282815181106112b9576112b96142dc565b60200260200101516001611cff565b60010161129b565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015610b7a573d5f5f3e3d5ffd5b60ca5460408051634e86e3b760e11b815290516060926001600160a01b031691639d0dc76e916004808301925f9291908290030181865afa15801561135e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b9e9190810190614467565b60fe545f90610d82906001600160a01b031688888885808a8a8a610bb1565b5f6113ad611dc8565b6113b633611c54565b33906113e1576040516338905e7160e11b81526001600160a01b039091166004820152602401610beb565b506113ea611df8565b6113f560cf8361260a565b82906114175760405163473978bf60e01b8152600401610beb91815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b8282101561152a578382905f5260205f2001805461149f9061440d565b80601f01602080910402602001604051908101604052809291908181526020018280546114cb9061440d565b80156115165780601f106114ed57610100808354040283529160200191611516565b820191905f5260205f20905b8154815290600101906020018083116114f957829003601f168201915b505050505081526020019060010190611482565b505050508152505090505f5f61154d836020015184604001518560600151612621565b915091508181906115715760405162461bcd60e51b8152600401610beb919061415d565b5061157d60cf866127cf565b505f85815260ce60205260408120611594916136e9565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906115d96004830182613720565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373866060015142604051611639929190918252602082015260400190565b60405180910390a46001935050505061165e60015f5160206148ee5f395f51905f5255565b919050565b5f610b9e60976123bf565b611676611f76565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156117035750825b90505f826001600160401b0316600114801561171e5750303b155b90508115801561172c575080155b1561174a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561177457845460ff60401b1916600160401b1785555b61177e87876127da565b60fe80546001600160a01b0319166001600160a01b038a1617905583156117df57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f6117f382610f06565b8015610d5c5750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f611824610f22565b90505f81516001600160401b0381111561184057611840613806565b604051908082528060200260200182016040528015611869578160200160208202803683370190505b5090505f5b82518110156118f85760335f84838151811061188c5761188c6142dc565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b03168282815181106118d8576118d86142dc565b6001600160a01b039092166020928302919091019091015260010161186e565b5092915050565b5f610d5c609783612273565b5f61191a868787878787611385565b9695505050505050565b5f610b9e5f6123bf565b5f6065546001610b9e91906143fa565b5f611947611dc8565b86611951816117e9565b819061197c57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610beb565b50611985611df8565b5f61198e610d62565b9050808a8082146119bb57604051634982205b60e11b815260048101929092526024820152604401610beb565b50505f7fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8b8b8b8b8b8b6040516020016119fb97969594939291906145bb565b604051602081830303815290604052805190602001209050611a1d8186612886565b8b90611a3f57604051635b777d1160e11b8152600401610beb91815260200190565b50611a4e606680546001019055565b5f5f611a5b8c8c8c612621565b915091508115611ac1578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611ab4929190918252602082015260400190565b60405180910390a4611be7565b611acc60cf8e6129d0565b8d90611aee576040516368db995b60e11b8152600401610beb91815260200190565b505f8d815260ce60205260409020611b06828261463f565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611b4991906146f9565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611bb9926004850192019061373b565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b6001955050505050610e7d60015f5160206148ee5f395f51905f5255565b611c0d611f76565b6001600160a01b038116611c3657604051631e4fbdf760e01b81525f6004820152602401610beb565b610bae816123d4565b6060610b9e60976123c8565b5f610b9e6129db565b5f610d5c609783612336565b5f81815260ce60205260409020805460609190611c7c9061440d565b80601f0160208091040260200160405190810160405280929190818152602001828054611ca89061440d565b8015611cf35780601f10611cca57610100808354040283529160200191611cf3565b820191905f5260205f20905b815481529060010190602001808311611cd657829003601f168201915b50505050509050919050565b611d07611f76565b8015611d4957611d186097836129e4565b8290611d435760405163082cdf5560e21b81526001600160a01b039091166004820152602401610beb565b50611d81565b611d546097836129f8565b8290611d7f5760405163e491024560e01b81526001600160a01b039091166004820152602401610beb565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f5160206148ce5f395f51905f525460ff1615610dae5760405163d93c066560e01b815260040160405180910390fd5b5f5160206148ee5f395f51905f52805460011901611e2957604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f90819081906001600160a01b0316611e5357505f91508190506001611ee3565b60ca54604051634594558f60e11b81526001600160a01b0389811660048301526024820189905290911690638b28ab1e906044016040805180830381865afa158015611ea1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ec59190614445565b9093509150828510801590611eda5750818410155b15611ee3575060015b9450945094915050565b83611ef886886143fa565b611f0291906143fa565b8360400151101589879091611f3b57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610beb565b5050611f4683612a0c565b611f58898989898989600189896120ca565b505050505050505050565b60015f5160206148ee5f395f51905f5255565b33611fa87f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610dae5760405163118cdaa760e01b8152336004820152602401610beb565b611fda81610f06565b8015611ffd57506001600160a01b0381165f9081526034602052604090205460ff165b8190612028576040516340da71e560e01b81526001600160a01b039091166004820152602401610beb565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b612079612b10565b5f5160206148ce5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016116b4565b6120df8989886120da888a6143fa565b612b3f565b6120f86120ea61192e565b8a8a8a8a8a8a8a8a8a612b92565b611f58606580546001019055565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061218c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166121805f5160206148ae5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610dae5760405163703e46dd60e11b815260040160405180910390fd5b610bae611f76565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561220c575060408051601f3d908101601f1916820190925261220991810190614353565b60015b61223457604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610beb565b5f5160206148ae5f395f51905f52811461226457604051632a87526960e21b815260048101829052602401610beb565b61226e8383612c58565b505050565b5f610b2c8383612cad565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610dae5760405163703e46dd60e11b815260040160405180910390fd5b6122d082612cd3565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610b2c565b61236081612d80565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce910161232a565b5f610d5c825490565b60605f610b2c83612e2d565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61244d816117e9565b8190612478576040516340da71e560e01b81526001600160a01b039091166004820152602401610beb565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b6124cc611dc8565b5f5160206148ce5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336120b2565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061488e5f395f51905f529161254a9061440d565b80601f01602080910402602001604051908101604052809291908181526020018280546125769061440d565b80156125c15780601f10612598576101008083540402835291602001916125c1565b820191905f5260205f20905b8154815290600101906020018083116125a457829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061488e5f395f51905f529161254a9061440d565b5f8181526001830160205260408120541515610b2c565b5f606082156127c75760fe546001600160a01b039081169086160361264e5761264b606484614705565b92505b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820185905286169063a9059cbb906044016020604051808303815f875af19250505080156126b8575060408051601f3d908101601f191682019092526126b591810190614724565b60015b612769576126c4614743565b806308c379a0036126ed57506126d861475b565b806126e35750612730565b5f925090506127c7565b634e487b7103612730576126ff6147dd565b9061270a5750612730565b60408051602081018390525f9450016040516020818303038152906040529150506127c7565b3d808015612759576040519150601f19603f3d011682016040523d82523d5f602084013e61275e565b606091505b505f925090506127c7565b8015612789576001925060405180602001604052805f81525091506127c5565b5f92506040518060400160405280601f81526020017f427269646765457468657265756d3a207472616e73666572206661696c65640081525091505b505b935093915050565b5f610b2c8383612e85565b6127e2612f68565b6127eb33612fb1565b6127f3612fc2565b6127fb612fca565b612803612fda565b61280b612fea565b6001600160a01b03821661285157604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610beb565b4360cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff168110156128a1575f915050610d5c565b5f80826001600160401b038111156128bb576128bb613806565b6040519080825280602002602001820160405280156128e4578160200160208202803683370190505b5090505f5b85518110156129be575f868281518110612905576129056142dc565b60200260200101519050604181511015612926575f95505050505050610d5c565b5f61293a826129348b613049565b90613075565b905061294581611c54565b612957575f9650505050505050610d5c565b5f805b85518110156129a657858181518110612975576129756142dc565b60200260200101516001600160a01b0316836001600160a01b03160361299e57600191506129a6565b60010161295a565b50806129b3578560010195505b5050506001016128e9565b505060995460ff161115949350505050565b5f610b2c838361309d565b5f610b9e6130e9565b5f610b2c836001600160a01b03841661309d565b5f610b2c836001600160a01b038416612e85565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612ab8576040513d5f823e3d81fd5b50505f513d91508115612acf578060011415612add565b84516001600160a01b03163b155b15612b09578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610beb565b5050505050565b5f5160206148ce5f395f51905f525460ff16610dae57604051638dfc202b60e01b815260040160405180910390fd5b612b608330612b4e84866143fa565b6001600160a01b03881692919061315c565b8015612b8c57612b8c612b7b60cc546001600160a01b031690565b6001600160a01b03861690836131c3565b50505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a604051612bed97969594939291906147fa565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051612c44929190918252602082015260400190565b60405180910390a450505050505050505050565b612c61826131f4565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612ca55761226e8282613257565b610cc16132c9565b5f825f018281548110612cc257612cc26142dc565b905f5260205f200154905092915050565b806001600160a01b038116612d13576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610beb565b612d1d5f836129e4565b8290612d48576040516351eccfe160e11b81526001600160a01b039091166004820152602401610beb565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116612dc0576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610beb565b612dca5f836129f8565b8290612df5576040516340da71e560e01b81526001600160a01b039091166004820152602401610beb565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611cf357602002820191905f5260205f20905b815481526020019060010190808311612e665750505050509050919050565b5f8181526001830160205260408120548015612f5f575f612ea760018361483c565b85549091505f90612eba9060019061483c565b9050808214612f19575f865f018281548110612ed857612ed86142dc565b905f5260205f200154905080875f018481548110612ef857612ef86142dc565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612f2a57612f2a61484f565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610d5c565b5f915050610d5c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610dae57604051631afcd79f60e31b815260040160405180910390fd5b612fb9612f68565b610bae816132e8565b610dae612f68565b612fd2612f68565b610dae6132f0565b612fe2612f68565b610dae613310565b612ff2612f68565b61303a604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613318565b6099805460ff19166003179055565b5f610d5c6130556129db565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613083868661332a565b9250925092506130938282613373565b5090949350505050565b5f8181526001830160205260408120546130e257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610d5c565b505f610d5c565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61311361342b565b61311b613493565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6040516001600160a01b038481166024830152838116604483015260648201839052612b8c9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506134d5565b6040516001600160a01b0383811660248301526044820183905261226e91859182169063a9059cbb90606401613191565b806001600160a01b03163b5f0361322957604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610beb565b5f5160206148ae5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516132739190614863565b5f60405180830381855af49150503d805f81146132ab576040519150601f19603f3d011682016040523d82523d5f602084013e6132b0565b606091505b50915091506132c0858383613541565b95945050505050565b3415610dae5760405163b398979f60e01b815260040160405180910390fd5b611c0d612f68565b6132f8612f68565b5f5160206148ce5f395f51905f52805460ff19169055565b611f63612f68565b613320612f68565b610cc1828261359d565b5f5f5f8351604103613361576020840151604085015160608601515f1a613353888285856135fc565b95509550955050505061336c565b505081515f91506002905b9250925092565b5f82600381111561338657613386614879565b0361338f575050565b60018260038111156133a3576133a3614879565b036133c15760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156133d5576133d5614879565b036133f65760405163fce698f760e01b815260048101829052602401610beb565b600382600381111561340a5761340a614879565b03610cc1576040516335e2f38360e21b815260048101829052602401610beb565b5f5f51602061488e5f395f51905f528161344361250c565b80519091501561345b57805160209091012092915050565b8154801561346a579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061488e5f395f51905f52816134ab6125cc565b8051909150156134c357805160209091012092915050565b6001820154801561346a579392505050565b5f5f60205f8451602086015f885af1806134f4576040513d5f823e3d81fd5b50505f513d9150811561350b578060011415613518565b6001600160a01b0384163b155b15612b8c57604051635274afe760e01b81526001600160a01b0385166004820152602401610beb565b60608261355657613551826136c0565b610b2c565b815115801561356d57506001600160a01b0384163b155b1561359657604051639996b31560e01b81526001600160a01b0385166004820152602401610beb565b5080610b2c565b6135a5612f68565b5f51602061488e5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026135de848261463f565b50600381016135ed838261463f565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561363557505f91506003905082611ee3565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613686573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166136b157505f925060019150829050611ee3565b975f9750879650945050505050565b8051156136d05780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5080546136f59061440d565b5f825580601f10613704575050565b601f0160209004905f5260205f2090810190610bae919061378f565b5080545f8255905f5260205f2090810190610bae91906137a3565b828054828255905f5260205f2090810192821561377f579160200282015b8281111561377f578251829061376f908261463f565b5091602001919060010190613759565b5061378b9291506137a3565b5090565b5b8082111561378b575f8155600101613790565b8082111561378b575f6137b682826136e9565b506001016137a3565b5f5f83601f8401126137cf575f5ffd5b5081356001600160401b038111156137e5575f5ffd5b6020830191508360208260051b85010111156137ff575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b038211171561383957613839613806565b60405250565b601f8201601f191681016001600160401b038111828210171561386457613864613806565b6040525050565b5f6001600160401b0382111561388357613883613806565b5060051b60200190565b5f82601f83011261389c575f5ffd5b81356001600160401b038111156138b5576138b5613806565b6040516138cc601f8301601f19166020018261383f565b8181528460208386010111156138e0575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6139068361386b565b604051613913828261383f565b84815291505060208101600584901b830185811115613930575f5ffd5b835b8181101561396a5780356001600160401b0381111561394f575f5ffd5b61395b8882880161388d565b84525060209283019201613932565b5050509392505050565b5f82601f830112613983575f5ffd5b610b2c838335602085016138fc565b5f5f5f604084860312156139a4575f5ffd5b83356001600160401b038111156139b9575f5ffd5b6139c5868287016137bf565b90945092505060208401356001600160401b038111156139e3575f5ffd5b8401601f810186136139f3575f5ffd5b80356139fe8161386b565b604051613a0b828261383f565b80915082815260208101915060208360051b850101925088831115613a2e575f5ffd5b602084015b83811015613a6e5780356001600160401b03811115613a50575f5ffd5b613a5f8b602083890101613974565b84525060209283019201613a33565b50809450505050509250925092565b6001600160a01b0381168114610bae575f5ffd5b5f60208284031215613aa1575f5ffd5b8135610b2c81613a7d565b803560ff8116811461165e575f5ffd5b5f60e08284031215613acc575f5ffd5b604051613ad88161381a565b8091508235613ae681613a7d565b81526020830135613af681613a7d565b60208201526040838101359082015260608084013590820152613b1b60808401613aac565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613b54575f5ffd5b8935613b5f81613a7d565b985060208a0135613b6f81613a7d565b975060408a0135613b7f81613a7d565b965060608a0135955060808a0135945060a08a01359350613ba38b60c08c01613abc565b92506101a08a01356001600160401b03811115613bbe575f5ffd5b613bca8c828d016137bf565b915080935050809150509295985092959850929598565b5f60208284031215613bf1575f5ffd5b81356001600160401b03811115613c06575f5ffd5b8201601f81018413613c16575f5ffd5b8035613c218161386b565b604051613c2e828261383f565b80915082815260208101915060208360051b850101925086831115613c51575f5ffd5b6020840193505b8284101561191a578335613c6b81613a7d565b825260209384019390910190613c58565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610d5c565b5f5f5f5f5f5f60a08789031215613cbb575f5ffd5b8635613cc681613a7d565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613cf5575f5ffd5b613d0189828a016137bf565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613d29575f5ffd5b8735613d3481613a7d565b96506020880135613d4481613a7d565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115613d73575f5ffd5b613d7f8a828b016137bf565b989b979a50959850939692959293505050565b5f5f60408385031215613da3575f5ffd5b8235613dae81613a7d565b915060208301356001600160401b03811115613dc8575f5ffd5b613dd48582860161388d565b9150509250929050565b5f60208284031215613dee575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b031215613e0d575f5ffd5b8835613e1881613a7d565b97506020890135613e2881613a7d565b9650604089013595506060890135945060808901359350613e4c8a60a08b01613abc565b92506101808901356001600160401b03811115613e67575f5ffd5b613e738b828c016137bf565b999c989b5096995094979396929594505050565b5f5f60408385031215613e98575f5ffd5b8235613ea381613a7d565b91506020830135613eb381613a7d565b809150509250929050565b602080825282518282018190525f918401906040840190835b81811015613efe5783516001600160a01b0316835260209384019390920191600101613ed7565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b81811015613fd05760df19878603018352613fbb858551613f09565b94506020938401939290920191600101613f9f565b50929695505050505050565b60ff60f81b8816815260e060208201525f613ffa60e0830189613f09565b828103604084015261400c8189613f09565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614061578351835260209384019390920191600101614043565b50909b9a5050505050505050505050565b5f5f60408385031215614083575f5ffd5b823561408e81613a7d565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b81811015613efe576140e583855180516001600160a01b0316825260208082015190830152604090810151910152565b60209390930192606092909201916001016140b5565b5f5f5f5f5f5f6101608789031215614111575f5ffd5b863561411c81613a7d565b9550602087013561412c81613a7d565b9450604087013593506141428860608901613abc565b92506101408701356001600160401b03811115613cf5575f5ffd5b602081525f610b2c6020830184613f09565b5f6020828403121561417f575f5ffd5b610b2c82613aac565b5f5f5f6060848603121561419a575f5ffd5b83356141a581613a7d565b925060208401356141b581613a7d565b915060408401356141c581613a7d565b809150509250925092565b5f5f5f5f5f61014086880312156141e5575f5ffd5b85356141f081613a7d565b9450602086013593506142068760408801613abc565b92506101208601356001600160401b03811115614221575f5ffd5b61422d888289016137bf565b969995985093965092949392505050565b5f5f5f5f5f5f5f60c0888a031215614254575f5ffd5b87359650602088013561426681613a7d565b9550604088013561427681613a7d565b94506060880135935060808801356001600160401b03811115614297575f5ffd5b6142a38a828b016137bf565b90945092505060a08801356001600160401b038111156142c1575f5ffd5b6142cd8a828b01613974565b91505092959891949750929550565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e19833603018112614304575f5ffd5b9190910192915050565b5f5f8335601e19843603018112614323575f5ffd5b8301803591506001600160401b0382111561433c575f5ffd5b6020019150600581901b36038213156137ff575f5ffd5b5f60208284031215614363575f5ffd5b5051919050565b5f6060828403121561437a575f5ffd5b604051606081018181106001600160401b038211171561439c5761439c613806565b806040525080915082516143af81613a7d565b815260208381015190820152604092830151920191909152919050565b5f606082840312156143dc575f5ffd5b610b2c838361436a565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d5c57610d5c6143e6565b600181811c9082168061442157607f821691505b60208210810361443f57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f60408385031215614456575f5ffd5b505080516020909101519092909150565b5f60208284031215614477575f5ffd5b81516001600160401b0381111561448c575f5ffd5b8201601f8101841361449c575f5ffd5b80516144a78161386b565b6040516144b4828261383f565b82815260206060909302840183019281019150868311156144d3575f5ffd5b6020840193505b8284101561191a576144ec878561436a565b82526020820191506060840193506144da565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b868110156145af57838303601f19018852813536879003601e19018112614563575f5ffd5b86016020810190356001600160401b0381111561457e575f5ffd5b80360382131561458c575f5ffd5b6145978582846144ff565b60209a8b019a9095509390930192505060010161453e565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90610ecb9083018486614527565b601f82111561226e57805f5260205f20601f840160051c810160208510156146205750805b601f840160051c820191505b81811015612b09575f815560010161462c565b81516001600160401b0381111561465857614658613806565b61466c81614666845461440d565b846145fb565b6020601f82116001811461469e575f83156146875750848201515b5f19600385901b1c1916600184901b178455612b09565b5f84815260208120601f198516915b828110156146cd57878501518255602094850194600190920191016146ad565b50848210156146ea57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610b2c3684846138fc565b5f8261471f57634e487b7160e01b5f52601260045260245ffd5b500490565b5f60208284031215614734575f5ffd5b81518015158114610b2c575f5ffd5b5f60033d1115610ef15760045f5f3e505f5160e01c90565b5f60443d10156147685790565b6040513d600319016004823e80513d60248201116001600160401b038211171561479157505090565b80820180516001600160401b038111156147ac575050505090565b3d84016003190182820160200111156147c6575050505090565b6147d56020828501018561383f565b509392505050565b5f5f60233d11156147f657602060045f3e50505f516001905b9091565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f90610ecb9083018486614527565b81810381811115610d5c57610d5c6143e6565b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f920191825250919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220ad0657e77cd34c80599278af5fb90b82252af439b3948f606b2ea320c53a0c3264736f6c634300081c0033",
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

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeEthereum *BridgeEthereumCaller) BridgeFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "BridgeFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeEthereum *BridgeEthereumSession) BridgeFeeManager() (common.Address, error) {
	return _BridgeEthereum.Contract.BridgeFeeManager(&_BridgeEthereum.CallOpts)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeEthereum *BridgeEthereumCallerSession) BridgeFeeManager() (common.Address, error) {
	return _BridgeEthereum.Contract.BridgeFeeManager(&_BridgeEthereum.CallOpts)
}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) BridgeFeeManagerLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "BridgeFeeManagerLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) BridgeFeeManagerLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.BridgeFeeManagerLength(&_BridgeEthereum.CallOpts)
}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) BridgeFeeManagerLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.BridgeFeeManagerLength(&_BridgeEthereum.CallOpts)
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

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumCaller) AllFeeInfo(opts *bind.CallOpts) ([]IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "allFeeInfo")

	if err != nil {
		return *new([]IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeFeeManagerFeeInfo)).(*[]IBridgeFeeManagerFeeInfo)

	return out0, err

}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.AllFeeInfo(&_BridgeEthereum.CallOpts)
}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumCallerSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.AllFeeInfo(&_BridgeEthereum.CallOpts)
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

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumCaller) CalculateFee(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "calculateFee", token, value)

	outstruct := new(struct {
		Gas     *big.Int
		Service *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Service = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeEthereum.Contract.CalculateFee(&_BridgeEthereum.CallOpts, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumCallerSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeEthereum.Contract.CalculateFee(&_BridgeEthereum.CallOpts, token, value)
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

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCaller) FeeInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "feeInfoByIndex", index)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.FeeInfoByIndex(&_BridgeEthereum.CallOpts, index)
}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCallerSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.FeeInfoByIndex(&_BridgeEthereum.CallOpts, index)
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

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCaller) GetTokenFee(opts *bind.CallOpts, token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "getTokenFee", token)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.GetTokenFee(&_BridgeEthereum.CallOpts, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCallerSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeEthereum.Contract.GetTokenFee(&_BridgeEthereum.CallOpts, token)
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
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) Bridge(opts *bind.TransactOpts, token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridge", token, value, gas, service, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Bridge(token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Bridge(&_BridgeEthereum.TransactOpts, token, value, gas, service, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) Bridge(token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Bridge(&_BridgeEthereum.TransactOpts, token, value, gas, service, extraData)
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
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeTo(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeTo", token, to, value, gas, service, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeTo(&_BridgeEthereum.TransactOpts, token, to, value, gas, service, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeTo(&_BridgeEthereum.TransactOpts, token, to, value, gas, service, extraData)
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

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) Finalize(opts *bind.TransactOpts, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "finalize", index, token, to, value, extraData, sigs)
}

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Finalize(&_BridgeEthereum.TransactOpts, index, token, to, value, extraData, sigs)
}

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Finalize(&_BridgeEthereum.TransactOpts, index, token, to, value, extraData, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) FinalizeBatch(opts *bind.TransactOpts, args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "finalizeBatch", args, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.FinalizeBatch(&_BridgeEthereum.TransactOpts, args, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.FinalizeBatch(&_BridgeEthereum.TransactOpts, args, sigs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _cross, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeEthereum *BridgeEthereumTransactor) Initialize(opts *bind.TransactOpts, _cross common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "initialize", _cross, rewardWallet_, BridgeFeeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _cross, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeEthereum *BridgeEthereumSession) Initialize(_cross common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Initialize(&_BridgeEthereum.TransactOpts, _cross, rewardWallet_, BridgeFeeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _cross, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeEthereum *BridgeEthereumTransactorSession) Initialize(_cross common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.Initialize(&_BridgeEthereum.TransactOpts, _cross, rewardWallet_, BridgeFeeManager)
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
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridge", token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, service, permitArgs, extraData)
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
	Index      *big.Int
	Token      common.Address
	Account    common.Address
	GasFee     *big.Int
	ServiceFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeCharged is a free log retrieval operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 serviceFee)
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
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 serviceFee)
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
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 serviceFee)
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
