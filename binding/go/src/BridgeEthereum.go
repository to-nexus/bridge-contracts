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
	ABI: "[{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeToCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenInfo\",\"outputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cross\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_cross\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenInfoLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"c97682f8": "allPairs()",
		"9fdf1c6a": "allTokenInfo()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"d86c62dc": "bridgeCross(address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"a01be2c7": "bridgeToCross(address,address,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes[])",
		"9c1b65a9": "bridgeTokenInfo()",
		"6e908ca3": "calculate(address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
		"fa074d03": "cross()",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"f120c400": "finalize(uint256,address,address,uint256,bytes[],bytes[])",
		"008bd028": "finalizeBatch((uint256,address,address,uint256,bytes[])[],bytes[][])",
		"71c59d7b": "getPairToken(address)",
		"1f69565f": "getTokenInfo(address)",
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
		"85547884": "removeTokenInfo()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"a9823183": "retryFinalize(uint256)",
		"7021fd0e": "revertedArguments(uint256)",
		"fe2b8da6": "revertedReason(uint256)",
		"fb75b2c7": "rewardWallet()",
		"2f9b59d1": "setTokenInfo(address)",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"e70a1b26": "tokenInfoByIndex(uint256)",
		"7cfed602": "tokenInfoLength()",
		"d92fc67b": "tokensLength()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"3b3bff0f": "unpauseToken(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614aa16100f95f395f81816122400152818161226901526123b80152614aa15ff3fe608060405260043610610370575f3560e01c806384b0196e116101c8578063c1876453116100fd578063f120c4001161009d578063fa074d031161006d578063fa074d03146109dc578063facd743b146109fb578063fb75b2c714610a1a578063fe2b8da614610a37575f5ffd5b8063f120c40014610982578063f2fde38b14610995578063f30589c3146109b4578063f698da25146109c8575f5ffd5b8063d86c62dc116100d8578063d86c62dc1461091c578063d92fc67b1461093b578063e1af7f501461094f578063e70a1b2614610963575f5ffd5b8063c1876453146108ca578063c97682f8146108e9578063cbae5958146108fd575f5ffd5b80639fdf1c6a11610168578063ad3cb1cc11610143578063ad3cb1cc1461083b578063aed1d40314610878578063b7f3358d1461088c578063c0c53b8b146108ab575f5ffd5b80639fdf1c6a146107dc578063a01be2c7146107fd578063a98231831461081c575f5ffd5b806391cf6d3e116101a357806391cf6d3e146107765780639300c9261461078a57806396ce0795146107a95780639c1b65a9146107bd575f5ffd5b806384b0196e146106ff57806385547884146107265780638da5cb5b1461073a575f5ffd5b80634f6ccce7116102a95780636e908ca31161024957806371c59d7b1161021957806371c59d7b146106815780637c41ad2c146106b85780637cfed602146106d75780638456cb59146106eb575f5ffd5b80636e908ca3146105e65780636ff97f1d146106205780637021fd0e14610641578063715018a61461066d575f5ffd5b80635476bd72116102845780635476bd72146105665780635c975abb146105855780635dbe47e8146105a85780635fa7b584146105c7575f5ffd5b80634f6ccce71461050857806351c455791461053f57806352d1902d14610552575f5ffd5b806337d10075116103145780633f5d3b5d116102ef5780633f5d3b5d146104a257806340a141ff146104b557806342cde4e8146104d45780634f1ef286146104f5575f5ffd5b806337d100751461045c5780633b3bff0f1461046f5780633f4ba83a1461048e575f5ffd5b80631d40f0d81161034f5780631d40f0d8146103d05780631f69565f146103ef57806327938c811461041b5780632f9b59d11461043d575f5ffd5b80628bd028146103745780631327d3d81461039c578063174991ab146103bd575b5f5ffd5b610387610382366004613ac5565b610a56565b60405190151581526020015b60405180910390f35b3480156103a7575f5ffd5b506103bb6103b6366004613bc4565b610b82565b005b6103876103cb366004613c6e565b610b90565b3480156103db575f5ffd5b506103bb6103ea366004613d14565b610c80565b3480156103fa575f5ffd5b5061040e610409366004613bc4565b610cba565b6040516103939190613daf565b348015610426575f5ffd5b5061042f610d5d565b604051908152602001610393565b348015610448575f5ffd5b506103bb610457366004613bc4565b610d72565b61038761046a366004613de3565b610de5565b34801561047a575f5ffd5b506103bb610489366004613bc4565b610e00565b348015610499575f5ffd5b506103bb610e11565b6103876104b0366004613e50565b610e23565b3480156104c0575f5ffd5b506103bb6104cf366004613bc4565b610f12565b3480156104df575f5ffd5b5060995460405160ff9091168152602001610393565b6103bb610503366004613ecf565b610f1c565b348015610513575f5ffd5b50610527610522366004613f1b565b610f37565b6040516001600160a01b039091168152602001610393565b61038761054d366004613f32565b610f42565b34801561055d575f5ffd5b5061042f610f61565b348015610571575f5ffd5b506103bb610580366004613fc4565b610f7d565b348015610590575f5ffd5b505f516020614a2c5f395f51905f525460ff16610387565b3480156105b3575f5ffd5b506103876105c2366004613bc4565b610f8f565b3480156105d2575f5ffd5b506103bb6105e1366004613bc4565b610f9a565b3480156105f1575f5ffd5b50610605610600366004613ffb565b610fab565b60408051938452602084019290925290820152606001610393565b34801561062b575f5ffd5b50610634611032565b6040516103939190614025565b34801561064c575f5ffd5b5061066061065b366004613f1b565b611093565b604051610393919061409e565b348015610678575f5ffd5b506103bb6111f3565b34801561068c575f5ffd5b5061052761069b366004613bc4565b6001600160a01b039081165f908152603360205260409020541690565b3480156106c3575f5ffd5b506103bb6106d2366004613bc4565b611204565b3480156106e2575f5ffd5b5061042f611215565b3480156106f6575f5ffd5b506103bb611280565b34801561070a575f5ffd5b50610713611290565b6040516103939796959493929190614143565b348015610731575f5ffd5b506103bb611339565b348015610745575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610527565b348015610781575f5ffd5b5060cb5461042f565b348015610795575f5ffd5b506103bb6107a4366004613d14565b611353565b3480156107b4575f5ffd5b5061042f61138a565b3480156107c8575f5ffd5b5060ca54610527906001600160a01b031681565b3480156107e7575f5ffd5b506107f06113d1565b60405161039391906141d9565b348015610808575f5ffd5b50610387610817366004614242565b61143f565b348015610827575f5ffd5b50610387610836366004613f1b565b61145e565b348015610846575f5ffd5b5061086b604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161039391906142a4565b348015610883575f5ffd5b5061042f61171d565b348015610897575f5ffd5b506103bb6108a63660046142b6565b611728565b3480156108b6575f5ffd5b506103bb6108c53660046142cf565b611779565b3480156108d5575f5ffd5b506103876108e4366004613bc4565b6118a3565b3480156108f4575f5ffd5b506106346118d3565b348015610908575f5ffd5b50610527610917366004613f1b565b6119b9565b348015610927575f5ffd5b50610387610936366004614317565b6119c5565b348015610946575f5ffd5b5061042f6119de565b34801561095a575f5ffd5b5061042f6119e8565b34801561096e575f5ffd5b5061040e61097d366004613f1b565b6119f8565b610387610990366004614385565b611a59565b3480156109a0575f5ffd5b506103bb6109af366004613bc4565b611d1f565b3480156109bf575f5ffd5b50610634611d59565b3480156109d3575f5ffd5b5061042f611d65565b3480156109e7575f5ffd5b5060fd54610527906001600160a01b031681565b348015610a06575f5ffd5b50610387610a15366004613bc4565b611d6e565b348015610a25575f5ffd5b5060cc546001600160a01b0316610527565b348015610a42575f5ffd5b5061086b610a51366004613f1b565b611d7a565b5f805b83811015610b7557610b6c858583818110610a7657610a76614423565b9050602002810190610a889190614437565b35868684818110610a9b57610a9b614423565b9050602002810190610aad9190614437565b610abe906040810190602001613bc4565b878785818110610ad057610ad0614423565b9050602002810190610ae29190614437565b610af3906060810190604001613bc4565b888886818110610b0557610b05614423565b9050602002810190610b179190614437565b60600135898987818110610b2d57610b2d614423565b9050602002810190610b3f9190614437565b610b4d906080810190614455565b898881518110610b5f57610b5f614423565b6020026020010151611a59565b50600101610a59565b50600190505b9392505050565b610b8d816001611e19565b50565b5f610b99611ee2565b89610ba3816118a3565b8190610bd357604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610bdc611f12565b5f5f5f5f610bec8f8d8d8d611f49565b929650909450925090508383838e8e8e86610c3d576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610bca565b505050505050610c548f8f8f8f87878f8f8f61201c565b6001955050505050610c7260015f516020614a4c5f395f51905f5255565b509998505050505050505050565b5f5b8151811015610cb657610cae828281518110610ca057610ca0614423565b60200260200101515f611e19565b600101610c82565b5050565b610cea60405180608001604052805f6001600160a01b031681526020015f81526020015f81526020015f81525090565b60ca54604051631f69565f60e01b81526001600160a01b03848116600483015290911690631f69565f906024015b608060405180830381865afa158015610d33573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d579190614506565b92915050565b5f6066546001610d6d9190614534565b905090565b610d7a6120a5565b6001600160a01b038116610dc357604051633effd40360e21b815260206004820152600f60248201526e627269646765546f6b656e496e666f60881b6044820152606401610bca565b60ca80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610df587338888888888610e23565b979650505050505050565b610e086120a5565b610b8d81612100565b610e196120a5565b610e216121a0565b565b5f610e2c611ee2565b87610e36816118a3565b8190610e6157604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bca565b50610e6a611f12565b5f5f5f5f610e7a8d8c8c8c611f49565b929650909450925090508383838d8d8d86610ecb576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610bca565b505050505050610ee88d610edc3390565b8e8e87875f8f8f6121f9565b6001955050505050610f0660015f516020614a4c5f395f51905f5255565b50979650505050505050565b610b8d815f611e19565b610f24612235565b610f2d826122d9565b610cb682826122e1565b5f610d5781836123a2565b5f610f5489898a8a8a8a8a8a8a610b90565b9998505050505050505050565b5f610f6a6123ad565b505f516020614a0c5f395f51905f525b90565b610f856120a5565b610cb682826123f6565b5f610d578183612465565b610fa26120a5565b610b8d81612486565b60ca54604051636e908ca360e01b81526001600160a01b038481166004830152602482018490525f928392839290911690636e908ca390604401606060405180830381865afa158015611000573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110249190614547565b919790965090945092505050565b60605f61103e5f6124ee565b9050806001600160401b0381111561105857611058613939565b604051908082528060200260200182016040528015611081578160200160208202803683370190505b50915061108d5f6124f7565b91505090565b6110d36040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b828210156111e5578382905f5260205f2001805461115a90614572565b80601f016020809104026020016040519081016040528092919081815260200182805461118690614572565b80156111d15780601f106111a8576101008083540402835291602001916111d1565b820191905f5260205f20905b8154815290600101906020018083116111b457829003601f168201915b50505050508152602001906001019061113d565b505050915250909392505050565b6111fb6120a5565b610e215f612503565b61120c6120a5565b610b8d81612573565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa15801561125c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6d91906145aa565b6112886120a5565b610e216125f3565b5f60608082808083815f5160206149ec5f395f51905f5280549091501580156112bb57506001810154155b6112ff5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610bca565b61130761263b565b61130f6126fb565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6113416120a5565b60ca80546001600160a01b0319169055565b5f5b8151811015610cb65761138282828151811061137357611373614423565b60200260200101516001611e19565b600101611355565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa15801561125c573d5f5f3e3d5ffd5b60ca5460408051634fef8e3560e11b815290516060926001600160a01b031691639fdf1c6a916004808301925f9291908290030181865afa158015611418573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d6d91908101906145c1565b60fd545f90610df5906001600160a01b031688888885808a8a8a610b90565b5f611467611ee2565b61147033611d6e565b339061149b576040516338905e7160e11b81526001600160a01b039091166004820152602401610bca565b506114a4611f12565b6114af60cf83612739565b82906114d15760405163473978bf60e01b8152600401610bca91815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156115e4578382905f5260205f2001805461155990614572565b80601f016020809104026020016040519081016040528092919081815260200182805461158590614572565b80156115d05780601f106115a7576101008083540402835291602001916115d0565b820191905f5260205f20905b8154815290600101906020018083116115b357829003601f168201915b50505050508152602001906001019061153c565b505050508152505090505f5f611607836020015184604001518560600151612750565b9150915081819061162b5760405162461bcd60e51b8152600401610bca91906142a4565b5061163760cf866128fe565b505f85815260ce6020526040812061164e9161381c565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906116936004830182613853565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738660600151426040516116f3929190918252602082015260400190565b60405180910390a46001935050505061171860015f516020614a4c5f395f51905f5255565b919050565b5f610d6d60976124ee565b6117306120a5565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156117bd5750825b90505f826001600160401b031660011480156117d85750303b155b9050811580156117e6575080155b156118045760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561182e57845460ff60401b1916600160401b1785555b6118388787612909565b60fd80546001600160a01b0319166001600160a01b038a16179055831561189957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f6118ad82610f8f565b8015610d575750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f6118de611032565b90505f81516001600160401b038111156118fa576118fa613939565b604051908082528060200260200182016040528015611923578160200160208202803683370190505b5090505f5b82518110156119b25760335f84838151811061194657611946614423565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b031682828151811061199257611992614423565b6001600160a01b0390921660209283029190910190910152600101611928565b5092915050565b5f610d576097836123a2565b5f6119d486878787878761143f565b9695505050505050565b5f610d6d5f6124ee565b5f6065546001610d6d9190614534565b611a2860405180608001604052805f6001600160a01b031681526020015f81526020015f81526020015f81525090565b60ca546040516373850d9360e11b8152600481018490526001600160a01b039091169063e70a1b2690602401610d18565b5f611a62611ee2565b86611a6c816118a3565b8190611a9757604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bca565b50611aa0611f12565b5f611aa9610d5d565b9050808a808214611ad657604051634982205b60e11b815260048101929092526024820152604401610bca565b50505f7fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8b8b8b8b8b8b604051602001611b169796959493929190614719565b604051602081830303815290604052805190602001209050611b3881866129b5565b8b90611b5a57604051635b777d1160e11b8152600401610bca91815260200190565b50611b69606680546001019055565b5f5f611b768c8c8c612750565b915091508115611bdc578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611bcf929190918252602082015260400190565b60405180910390a4610ee8565b611be760cf8e612aff565b8d90611c09576040516368db995b60e11b8152600401610bca91815260200190565b505f8d815260ce60205260409020611c21828261479d565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611c649190614857565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611cd4926004850192019061386e565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a26001955050505050610f0660015f516020614a4c5f395f51905f5255565b611d276120a5565b6001600160a01b038116611d5057604051631e4fbdf760e01b81525f6004820152602401610bca565b610b8d81612503565b6060610d6d60976124f7565b5f610d6d612b0a565b5f610d57609783612465565b5f81815260ce60205260409020805460609190611d9690614572565b80601f0160208091040260200160405190810160405280929190818152602001828054611dc290614572565b8015611e0d5780601f10611de457610100808354040283529160200191611e0d565b820191905f5260205f20905b815481529060010190602001808311611df057829003601f168201915b50505050509050919050565b611e216120a5565b8015611e6357611e32609783612b13565b8290611e5d5760405163082cdf5560e21b81526001600160a01b039091166004820152602401610bca565b50611e9b565b611e6e609783612b27565b8290611e995760405163e491024560e01b81526001600160a01b039091166004820152602401610bca565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614a2c5f395f51905f525460ff1615610e215760405163d93c066560e01b815260040160405180910390fd5b5f516020614a4c5f395f51905f52805460011901611f4357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f908190819081906001600160a01b0316611f7257505f92508291508190506001612011565b60ca54604051636e908ca360e01b81526001600160a01b038a81166004830152602482018a905290911690636e908ca390604401606060405180830381865afa158015611fc1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fe59190614547565b91955093509150828610801590611ffc5750818510155b80156120085750838710155b15612011575060015b945094509450949050565b836120278688614534565b6120319190614534565b836040015110158987909161206a57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610bca565b505061207583612b3b565b612087898989898989600189896121f9565b505050505050505050565b60015f516020614a4c5f395f51905f5255565b336120d77f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610e215760405163118cdaa760e01b8152336004820152602401610bca565b61210981610f8f565b801561212c57506001600160a01b0381165f9081526034602052604090205460ff165b8190612157576040516340da71e560e01b81526001600160a01b039091166004820152602401610bca565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b6121a8612c3f565b5f516020614a2c5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200161176e565b61220e898988612209888a614534565b612c6e565b6122276122196119e8565b8a8a8a8a8a8a8a8a8a612cc1565b612087606580546001019055565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806122bb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166122af5f516020614a0c5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e215760405163703e46dd60e11b815260040160405180910390fd5b610b8d6120a5565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561233b575060408051601f3d908101601f19168201909252612338918101906145aa565b60015b61236357604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610bca565b5f516020614a0c5f395f51905f52811461239357604051632a87526960e21b815260048101829052602401610bca565b61239d8383612d87565b505050565b5f610b7b8383612ddc565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e215760405163703e46dd60e11b815260040160405180910390fd5b6123ff82612e02565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610b7b565b61248f81612eaf565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce9101612459565b5f610d57825490565b60605f610b7b83612f5c565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61257c816118a3565b81906125a7576040516340da71e560e01b81526001600160a01b039091166004820152602401610bca565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b6125fb611ee2565b5f516020614a2c5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336121e1565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206149ec5f395f51905f529161267990614572565b80601f01602080910402602001604051908101604052809291908181526020018280546126a590614572565b80156126f05780601f106126c7576101008083540402835291602001916126f0565b820191905f5260205f20905b8154815290600101906020018083116126d357829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206149ec5f395f51905f529161267990614572565b5f8181526001830160205260408120541515610b7b565b5f606082156128f65760fd546001600160a01b039081169086160361277d5761277a606484614863565b92505b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820185905286169063a9059cbb906044016020604051808303815f875af19250505080156127e7575060408051601f3d908101601f191682019092526127e491810190614882565b60015b612898576127f36148a1565b806308c379a00361281c57506128076148b9565b80612812575061285f565b5f925090506128f6565b634e487b710361285f5761282e61493b565b90612839575061285f565b60408051602081018390525f9450016040516020818303038152906040529150506128f6565b3d808015612888576040519150601f19603f3d011682016040523d82523d5f602084013e61288d565b606091505b505f925090506128f6565b80156128b8576001925060405180602001604052805f81525091506128f4565b5f92506040518060400160405280601f81526020017f427269646765457468657265756d3a207472616e73666572206661696c65640081525091505b505b935093915050565b5f610b7b8383612fb4565b612911613097565b61291a336130e0565b6129226130f1565b61292a6130f9565b612932613109565b61293a613119565b6001600160a01b03821661298057604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610bca565b4360cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff168110156129d0575f915050610d57565b5f80826001600160401b038111156129ea576129ea613939565b604051908082528060200260200182016040528015612a13578160200160208202803683370190505b5090505f5b8551811015612aed575f868281518110612a3457612a34614423565b60200260200101519050604181511015612a55575f95505050505050610d57565b5f612a6982612a638b613178565b906131a4565b9050612a7481611d6e565b612a86575f9650505050505050610d57565b5f805b8551811015612ad557858181518110612aa457612aa4614423565b60200260200101516001600160a01b0316836001600160a01b031603612acd5760019150612ad5565b600101612a89565b5080612ae2578560010195505b505050600101612a18565b505060995460ff161115949350505050565b5f610b7b83836131cc565b5f610d6d613218565b5f610b7b836001600160a01b0384166131cc565b5f610b7b836001600160a01b038416612fb4565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612be7576040513d5f823e3d81fd5b50505f513d91508115612bfe578060011415612c0c565b84516001600160a01b03163b155b15612c38578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610bca565b5050505050565b5f516020614a2c5f395f51905f525460ff16610e2157604051638dfc202b60e01b815260040160405180910390fd5b612c8f8330612c7d8486614534565b6001600160a01b03881692919061328b565b8015612cbb57612cbb612caa60cc546001600160a01b031690565b6001600160a01b03861690836132f2565b50505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a604051612d1c9796959493929190614958565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051612d73929190918252602082015260400190565b60405180910390a450505050505050505050565b612d9082613323565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612dd45761239d8282613386565b610cb66133f8565b5f825f018281548110612df157612df1614423565b905f5260205f200154905092915050565b806001600160a01b038116612e42576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bca565b612e4c5f83612b13565b8290612e77576040516351eccfe160e11b81526001600160a01b039091166004820152602401610bca565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116612eef576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bca565b612ef95f83612b27565b8290612f24576040516340da71e560e01b81526001600160a01b039091166004820152602401610bca565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611e0d57602002820191905f5260205f20905b815481526020019060010190808311612f955750505050509050919050565b5f818152600183016020526040812054801561308e575f612fd660018361499a565b85549091505f90612fe99060019061499a565b9050808214613048575f865f01828154811061300757613007614423565b905f5260205f200154905080875f01848154811061302757613027614423565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080613059576130596149ad565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610d57565b5f915050610d57565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e2157604051631afcd79f60e31b815260040160405180910390fd5b6130e8613097565b610b8d81613417565b610e21613097565b613101613097565b610e2161341f565b613111613097565b610e2161343f565b613121613097565b613169604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613447565b6099805460ff19166003179055565b5f610d57613184612b0a565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6131b28686613459565b9250925092506131c282826134a2565b5090949350505050565b5f81815260018301602052604081205461321157508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610d57565b505f610d57565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61324261355a565b61324a6135c2565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6040516001600160a01b038481166024830152838116604483015260648201839052612cbb9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613604565b6040516001600160a01b0383811660248301526044820183905261239d91859182169063a9059cbb906064016132c0565b806001600160a01b03163b5f0361335857604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610bca565b5f516020614a0c5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516133a291906149c1565b5f60405180830381855af49150503d805f81146133da576040519150601f19603f3d011682016040523d82523d5f602084013e6133df565b606091505b50915091506133ef858383613670565b95945050505050565b3415610e215760405163b398979f60e01b815260040160405180910390fd5b611d27613097565b613427613097565b5f516020614a2c5f395f51905f52805460ff19169055565b612092613097565b61344f613097565b610cb682826136cc565b5f5f5f8351604103613490576020840151604085015160608601515f1a6134828882858561372b565b95509550955050505061349b565b505081515f91506002905b9250925092565b5f8260038111156134b5576134b56149d7565b036134be575050565b60018260038111156134d2576134d26149d7565b036134f05760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613504576135046149d7565b036135255760405163fce698f760e01b815260048101829052602401610bca565b6003826003811115613539576135396149d7565b03610cb6576040516335e2f38360e21b815260048101829052602401610bca565b5f5f5160206149ec5f395f51905f528161357261263b565b80519091501561358a57805160209091012092915050565b81548015613599579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206149ec5f395f51905f52816135da6126fb565b8051909150156135f257805160209091012092915050565b60018201548015613599579392505050565b5f5f60205f8451602086015f885af180613623576040513d5f823e3d81fd5b50505f513d9150811561363a578060011415613647565b6001600160a01b0384163b155b15612cbb57604051635274afe760e01b81526001600160a01b0385166004820152602401610bca565b60608261368557613680826137f3565b610b7b565b815115801561369c57506001600160a01b0384163b155b156136c557604051639996b31560e01b81526001600160a01b0385166004820152602401610bca565b5080610b7b565b6136d4613097565b5f5160206149ec5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10261370d848261479d565b506003810161371c838261479d565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561376457505f915060039050826137e9565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156137b5573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166137e057505f9250600191508290506137e9565b92505f91508190505b9450945094915050565b8051156138035780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b50805461382890614572565b5f825580601f10613837575050565b601f0160209004905f5260205f2090810190610b8d91906138c2565b5080545f8255905f5260205f2090810190610b8d91906138d6565b828054828255905f5260205f209081019282156138b2579160200282015b828111156138b257825182906138a2908261479d565b509160200191906001019061388c565b506138be9291506138d6565b5090565b5b808211156138be575f81556001016138c3565b808211156138be575f6138e9828261381c565b506001016138d6565b5f5f83601f840112613902575f5ffd5b5081356001600160401b03811115613918575f5ffd5b6020830191508360208260051b8501011115613932575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b038211171561396c5761396c613939565b60405250565b601f8201601f191681016001600160401b038111828210171561399757613997613939565b6040525050565b5f6001600160401b038211156139b6576139b6613939565b5060051b60200190565b5f82601f8301126139cf575f5ffd5b81356001600160401b038111156139e8576139e8613939565b6040516139ff601f8301601f191660200182613972565b818152846020838601011115613a13575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f613a398361399e565b604051613a468282613972565b84815291505060208101600584901b830185811115613a63575f5ffd5b835b81811015613a9d5780356001600160401b03811115613a82575f5ffd5b613a8e888288016139c0565b84525060209283019201613a65565b5050509392505050565b5f82601f830112613ab6575f5ffd5b610b7b83833560208501613a2f565b5f5f5f60408486031215613ad7575f5ffd5b83356001600160401b03811115613aec575f5ffd5b613af8868287016138f2565b90945092505060208401356001600160401b03811115613b16575f5ffd5b8401601f81018613613b26575f5ffd5b8035613b318161399e565b604051613b3e8282613972565b80915082815260208101915060208360051b850101925088831115613b61575f5ffd5b602084015b83811015613ba15780356001600160401b03811115613b83575f5ffd5b613b928b602083890101613aa7565b84525060209283019201613b66565b50809450505050509250925092565b6001600160a01b0381168114610b8d575f5ffd5b5f60208284031215613bd4575f5ffd5b8135610b7b81613bb0565b803560ff81168114611718575f5ffd5b5f60e08284031215613bff575f5ffd5b604051613c0b8161394d565b8091508235613c1981613bb0565b81526020830135613c2981613bb0565b60208201526040838101359082015260608084013590820152613c4e60808401613bdf565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613c87575f5ffd5b8935613c9281613bb0565b985060208a0135613ca281613bb0565b975060408a0135613cb281613bb0565b965060608a0135955060808a0135945060a08a01359350613cd68b60c08c01613bef565b92506101a08a01356001600160401b03811115613cf1575f5ffd5b613cfd8c828d016138f2565b915080935050809150509295985092959850929598565b5f60208284031215613d24575f5ffd5b81356001600160401b03811115613d39575f5ffd5b8201601f81018413613d49575f5ffd5b8035613d548161399e565b604051613d618282613972565b80915082815260208101915060208360051b850101925086831115613d84575f5ffd5b6020840193505b828410156119d4578335613d9e81613bb0565b825260209384019390910190613d8b565b81516001600160a01b0316815260208083015190820152604080830151908201526060808301519082015260808101610d57565b5f5f5f5f5f5f60a08789031215613df8575f5ffd5b8635613e0381613bb0565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613e32575f5ffd5b613e3e89828a016138f2565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613e66575f5ffd5b8735613e7181613bb0565b96506020880135613e8181613bb0565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115613eb0575f5ffd5b613ebc8a828b016138f2565b989b979a50959850939692959293505050565b5f5f60408385031215613ee0575f5ffd5b8235613eeb81613bb0565b915060208301356001600160401b03811115613f05575f5ffd5b613f11858286016139c0565b9150509250929050565b5f60208284031215613f2b575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b031215613f4a575f5ffd5b8835613f5581613bb0565b97506020890135613f6581613bb0565b9650604089013595506060890135945060808901359350613f898a60a08b01613bef565b92506101808901356001600160401b03811115613fa4575f5ffd5b613fb08b828c016138f2565b999c989b5096995094979396929594505050565b5f5f60408385031215613fd5575f5ffd5b8235613fe081613bb0565b91506020830135613ff081613bb0565b809150509250929050565b5f5f6040838503121561400c575f5ffd5b823561401781613bb0565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156140655783516001600160a01b031683526020938401939092019160010161403e565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156141375760df19878603018352614122858551614070565b94506020938401939290920191600101614106565b50929695505050505050565b60ff60f81b8816815260e060208201525f61416160e0830189614070565b82810360408401526141738189614070565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156141c85783518352602093840193909201916001016141aa565b50909b9a5050505050505050505050565b602080825282518282018190525f918401906040840190835b818110156140655761422c83855180516001600160a01b031682526020808201519083015260408082015190830152606090810151910152565b60209390930192608092909201916001016141f2565b5f5f5f5f5f5f6101608789031215614258575f5ffd5b863561426381613bb0565b9550602087013561427381613bb0565b9450604087013593506142898860608901613bef565b92506101408701356001600160401b03811115613e32575f5ffd5b602081525f610b7b6020830184614070565b5f602082840312156142c6575f5ffd5b610b7b82613bdf565b5f5f5f606084860312156142e1575f5ffd5b83356142ec81613bb0565b925060208401356142fc81613bb0565b9150604084013561430c81613bb0565b809150509250925092565b5f5f5f5f5f610140868803121561432c575f5ffd5b853561433781613bb0565b94506020860135935061434d8760408801613bef565b92506101208601356001600160401b03811115614368575f5ffd5b614374888289016138f2565b969995985093965092949392505050565b5f5f5f5f5f5f5f60c0888a03121561439b575f5ffd5b8735965060208801356143ad81613bb0565b955060408801356143bd81613bb0565b94506060880135935060808801356001600160401b038111156143de575f5ffd5b6143ea8a828b016138f2565b90945092505060a08801356001600160401b03811115614408575f5ffd5b6144148a828b01613aa7565b91505092959891949750929550565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e1983360301811261444b575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261446a575f5ffd5b8301803591506001600160401b03821115614483575f5ffd5b6020019150600581901b3603821315613932575f5ffd5b5f608082840312156144aa575f5ffd5b604051608081018181106001600160401b03821117156144cc576144cc613939565b806040525080915082516144df81613bb0565b81526020838101519082015260408084015190820152606092830151920191909152919050565b5f60808284031215614516575f5ffd5b610b7b838361449a565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d5757610d57614520565b5f5f5f60608486031215614559575f5ffd5b5050815160208301516040909301519094929350919050565b600181811c9082168061458657607f821691505b6020821081036145a457634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156145ba575f5ffd5b5051919050565b5f602082840312156145d1575f5ffd5b81516001600160401b038111156145e6575f5ffd5b8201601f810184136145f6575f5ffd5b80516146018161399e565b60405161460e8282613972565b80915082815260208101915060208360071b850101925086831115614631575f5ffd5b6020840193505b828410156119d45761464a878561449a565b8252602082019150608084019350614638565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b8681101561470d57838303601f19018852813536879003601e190181126146c1575f5ffd5b86016020810190356001600160401b038111156146dc575f5ffd5b8036038213156146ea575f5ffd5b6146f585828461465d565b60209a8b019a9095509390930192505060010161469c565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90610f549083018486614685565b601f82111561239d57805f5260205f20601f840160051c8101602085101561477e5750805b601f840160051c820191505b81811015612c38575f815560010161478a565b81516001600160401b038111156147b6576147b6613939565b6147ca816147c48454614572565b84614759565b6020601f8211600181146147fc575f83156147e55750848201515b5f19600385901b1c1916600184901b178455612c38565b5f84815260208120601f198516915b8281101561482b578785015182556020948501946001909201910161480b565b508482101561484857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610b7b368484613a2f565b5f8261487d57634e487b7160e01b5f52601260045260245ffd5b500490565b5f60208284031215614892575f5ffd5b81518015158114610b7b575f5ffd5b5f60033d1115610f7a5760045f5f3e505f5160e01c90565b5f60443d10156148c65790565b6040513d600319016004823e80513d60248201116001600160401b03821117156148ef57505090565b80820180516001600160401b0381111561490a575050505090565b3d8401600319018282016020011115614924575050505090565b61493360208285010185613972565b509392505050565b5f5f60233d111561495457602060045f3e50505f516001905b9091565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f90610f549083018486614685565b81810381811115610d5757610d57614520565b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f920191825250919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220b4670e28233ea80a7faedb8dde3a030968013b2eb6309652af3a549b30cca24d64736f6c634300081c0033",
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

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumCaller) AllTokenInfo(opts *bind.CallOpts) ([]IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "allTokenInfo")

	if err != nil {
		return *new([]IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeTokenInfoTokenInfo)).(*[]IBridgeTokenInfoTokenInfo)

	return out0, err

}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.AllTokenInfo(&_BridgeEthereum.CallOpts)
}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeEthereum *BridgeEthereumCallerSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.AllTokenInfo(&_BridgeEthereum.CallOpts)
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
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "calculate", token, value)

	outstruct := new(struct {
		Minimum *big.Int
		Gas     *big.Int
		Service *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minimum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Service = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeEthereum.Contract.Calculate(&_BridgeEthereum.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeEthereum *BridgeEthereumCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
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

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCaller) GetTokenInfo(opts *bind.CallOpts, token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "getTokenInfo", token)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.GetTokenInfo(&_BridgeEthereum.CallOpts, token)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCallerSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.GetTokenInfo(&_BridgeEthereum.CallOpts, token)
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

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCaller) TokenInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "tokenInfoByIndex", index)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.TokenInfoByIndex(&_BridgeEthereum.CallOpts, index)
}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeEthereum *BridgeEthereumCallerSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeEthereum.Contract.TokenInfoByIndex(&_BridgeEthereum.CallOpts, index)
}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCaller) TokenInfoLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeEthereum.contract.Call(opts, &out, "tokenInfoLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumSession) TokenInfoLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.TokenInfoLength(&_BridgeEthereum.CallOpts)
}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeEthereum *BridgeEthereumCallerSession) TokenInfoLength() (*big.Int, error) {
	return _BridgeEthereum.Contract.TokenInfoLength(&_BridgeEthereum.CallOpts)
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
