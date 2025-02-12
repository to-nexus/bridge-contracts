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
	ABI: "[{\"inputs\":[],\"name\":\"BridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BridgeFeeManagerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFeeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeToCross\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cross\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"feeInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_cross\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e3a446eb": "BridgeFeeManager()",
		"114ee544": "BridgeFeeManagerLength()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"9d0dc76e": "allFeeInfo()",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"aaa78020": "bridgeCross(address,uint256,uint256,bytes,bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"95d791e9": "bridgeToCross(address,address,uint256,uint256,bytes,bytes[])",
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
		"69a3318b": "permitBridge(address,address,uint256,uint256,uint256,uint256,bytes,bytes[])",
		"50d6fb48": "permitBridgeTo(address,address,address,uint256,uint256,uint256,uint256,bytes,bytes[])",
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
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516148266100f95f395f818161203b0152818161206401526122e501526148265ff3fe60806040526004361061030c575f3560e01c80637c41ad2c116101a3578063b7f3358d116100f2578063f120c40011610092578063fa074d031161006d578063fa074d03146108f8578063facd743b14610917578063fb75b2c714610936578063fe2b8da614610953575f5ffd5b8063f120c400146108b2578063f2fde38b146108c5578063f698da25146108e4575f5ffd5b8063c97682f8116100cd578063c97682f814610857578063d92fc67b1461086b578063e1af7f501461087f578063e3a446eb14610893575f5ffd5b8063b7f3358d146107fa578063c0c53b8b14610819578063c187645314610838575f5ffd5b80639300c9261161015d5780639d0dc76e116101385780639d0dc76e1461075e578063a98231831461077f578063aaa780201461079e578063ad3cb1cc146107bd575f5ffd5b80639300c9261461070c57806395d791e91461072b57806396ce07951461074a575f5ffd5b80637c41ad2c1461062e5780638456cb591461064d57806384b0196e146106615780638b28ab1e146106885780638da5cb5b146106bc57806391cf6d3e146106f8575f5ffd5b80634f6ccce71161025f5780635fa7b584116102195780637021fd0e116101f45780637021fd0e14610598578063715018a6146105c457806371c59d7b146105d8578063751b4c9c1461060f575f5ffd5b80635fa7b5841461054557806369a3318b146105645780636ff97f1d14610577575f5ffd5b80634f6ccce71461048657806350d6fb48146104bd57806352d1902d146104d05780635476bd72146104e45780635c975abb146105035780635dbe47e814610526575f5ffd5b806337d10075116102ca5780633f5d3b5d116102a55780633f5d3b5d1461042057806340a141ff1461043357806342cde4e8146104525780634f1ef28614610473575f5ffd5b806337d10075146103da5780633b3bff0f146103ed5780633f4ba83a1461040c575f5ffd5b80628bd02814610310578063114ee544146103385780631327d3d81461035a5780631d40f0d81461037b578063252154fa1461039a57806327938c81146103c6575b5f5ffd5b61032361031e36600461386f565b610972565b60405190151581526020015b60405180910390f35b348015610343575f5ffd5b5061034c610a9e565b60405190815260200161032f565b348015610365575f5ffd5b50610379610374366004613979565b610b0e565b005b348015610386575f5ffd5b50610379610395366004613994565b610b1c565b3480156103a5575f5ffd5b506103b96103b4366004613979565b610b56565b60405161032f9190613a39565b3480156103d1575f5ffd5b5061034c610bf3565b6103236103e8366004613a63565b610c03565b3480156103f8575f5ffd5b50610379610407366004613979565b610c1e565b348015610417575f5ffd5b50610379610c2f565b61032361042e366004613ad0565b610c41565b34801561043e575f5ffd5b5061037961044d366004613979565b610d1e565b34801561045d575f5ffd5b5060995460405160ff909116815260200161032f565b610379610481366004613b4f565b610d28565b348015610491575f5ffd5b506104a56104a0366004613b9b565b610d43565b6040516001600160a01b03909116815260200161032f565b6103236104cb366004613bb2565b610d4e565b3480156104db575f5ffd5b5061034c610e25565b3480156104ef575f5ffd5b506103796104fe366004613c76565b610e41565b34801561050e575f5ffd5b505f5160206147b15f395f51905f525460ff16610323565b348015610531575f5ffd5b50610323610540366004613979565b610e53565b348015610550575f5ffd5b5061037961055f366004613979565b610e5e565b610323610572366004613cad565b610e6f565b348015610582575f5ffd5b5061058b610e90565b60405161032f9190613d61565b3480156105a3575f5ffd5b506105b76105b2366004613b9b565b610ef1565b60405161032f9190613dda565b3480156105cf575f5ffd5b50610379611051565b3480156105e3575f5ffd5b506104a56105f2366004613979565b6001600160a01b039081165f908152603360205260409020541690565b34801561061a575f5ffd5b506103b9610629366004613b9b565b611062565b348015610639575f5ffd5b50610379610648366004613979565b6110bd565b348015610658575f5ffd5b506103796110ce565b34801561066c575f5ffd5b506106756110de565b60405161032f9796959493929190613e7f565b348015610693575f5ffd5b506106a76106a2366004613f15565b611187565b6040805192835260208301919091520161032f565b3480156106c7575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166104a5565b348015610703575f5ffd5b5060cb5461034c565b348015610717575f5ffd5b50610379610726366004613994565b611207565b348015610736575f5ffd5b50610323610745366004613f3f565b61123e565b348015610755575f5ffd5b5061034c61126a565b348015610769575f5ffd5b506107726112b1565b60405161032f9190613fbf565b34801561078a575f5ffd5b50610323610799366004613b9b565b61131f565b3480156107a9575f5ffd5b506103236107b836600461401e565b6115de565b3480156107c8575f5ffd5b506107ed604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161032f919061408d565b348015610805575f5ffd5b5061037961081436600461409f565b6115ee565b348015610824575f5ffd5b506103796108333660046140bf565b61163f565b348015610843575f5ffd5b50610323610852366004613979565b611769565b348015610862575f5ffd5b5061058b611799565b348015610876575f5ffd5b5061034c61187f565b34801561088a575f5ffd5b5061034c611889565b34801561089e575f5ffd5b5060ca546104a5906001600160a01b031681565b6103236108c0366004614107565b611899565b3480156108d0575f5ffd5b506103796108df366004613979565b611b95565b3480156108ef575f5ffd5b5061034c611bcf565b348015610903575f5ffd5b5060fe546104a5906001600160a01b031681565b348015610922575f5ffd5b50610323610931366004613979565b611bd8565b348015610941575f5ffd5b5060cc546001600160a01b03166104a5565b34801561095e575f5ffd5b506107ed61096d366004613b9b565b611be4565b5f805b83811015610a9157610a88858583818110610992576109926141a5565b90506020028101906109a491906141b9565b358686848181106109b7576109b76141a5565b90506020028101906109c991906141b9565b6109da906040810190602001613979565b8787858181106109ec576109ec6141a5565b90506020028101906109fe91906141b9565b610a0f906060810190604001613979565b888886818110610a2157610a216141a5565b9050602002810190610a3391906141b9565b60600135898987818110610a4957610a496141a5565b9050602002810190610a5b91906141b9565b610a699060808101906141d7565b898881518110610a7b57610a7b6141a5565b6020026020010151611899565b50600101610975565b50600190505b9392505050565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa158015610ae5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b09919061421c565b905090565b610b19816001611c83565b50565b5f5b8151811015610b5257610b4a828281518110610b3c57610b3c6141a5565b60200260200101515f611c83565b600101610b1e565b5050565b610b8060405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631290aa7d60e11b81526001600160a01b0384811660048301529091169063252154fa906024015b606060405180830381865afa158015610bc9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bed9190614295565b92915050565b5f6066546001610b0991906142c3565b5f610c1387338888888888610c41565b979650505050505050565b610c26611d4c565b610b1981611da7565b610c37611d4c565b610c3f611e47565b565b5f610c4a611ea0565b87610c5481611769565b8190610c8457604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610c8d611ed0565b5f5f5f610c9c8c8b8b8b611f07565b9194509250905082828a8a84610cdb576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610c7b565b50505050610cf58c610cea3390565b8d8d87878d8d611fe2565b60019450505050610d1260015f5160206147d15f395f51905f5255565b50979650505050505050565b610b19815f611c83565b610d30612030565b610d39826120d4565b610b5282826120dc565b5f610bed818361219d565b5f610d57611ea0565b8a610d6181611769565b8190610d8c57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610c7b565b50610d95611ed0565b5f5f5f610da48f8d8d8d611f07565b9194509250905082828c8c84610de3576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610c7b565b50505050610df98f8f8f8f87878f8f8f8f6121a8565b60019450505050610e1660015f5160206147d15f395f51905f5255565b509a9950505050505050505050565b5f610e2e6122da565b505f5160206147915f395f51905f525b90565b610e49611d4c565b610b528282612323565b5f610bed8183612392565b610e66611d4c565b610b19816123b3565b5f610e828a8a8b8b8b8b8b8b8b8b610d4e565b9a9950505050505050505050565b60605f610e9c5f61241b565b9050806001600160401b03811115610eb657610eb6613708565b604051908082528060200260200182016040528015610edf578160200160208202803683370190505b509150610eeb5f612424565b91505090565b610f316040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b82821015611043578382905f5260205f20018054610fb8906142d6565b80601f0160208091040260200160405190810160405280929190818152602001828054610fe4906142d6565b801561102f5780601f106110065761010080835404028352916020019161102f565b820191905f5260205f20905b81548152906001019060200180831161101257829003601f168201915b505050505081526020019060010190610f9b565b505050915250909392505050565b611059611d4c565b610c3f5f612430565b61108c60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631d46d32760e21b8152600481018490526001600160a01b039091169063751b4c9c90602401610bae565b6110c5611d4c565b610b19816124a0565b6110d6611d4c565b610c3f612520565b5f60608082808083815f5160206147715f395f51905f52805490915015801561110957506001810154155b61114d5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610c7b565b611155612568565b61115d612628565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60ca54604051634594558f60e11b81526001600160a01b038481166004830152602482018490525f928392911690638b28ab1e906044016040805180830381865afa1580156111d8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111fc919061430e565b909590945092505050565b5f5b8151811015610b5257611236828281518110611227576112276141a5565b60200260200101516001611c83565b600101611209565b60fe545f9061125e906001600160a01b031689898985808b8b8b8b610d4e565b98975050505050505050565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015610ae5573d5f5f3e3d5ffd5b60ca5460408051634e86e3b760e11b815290516060926001600160a01b031691639d0dc76e916004808301925f9291908290030181865afa1580156112f8573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b099190810190614330565b5f611328611ea0565b61133133611bd8565b339061135c576040516338905e7160e11b81526001600160a01b039091166004820152602401610c7b565b50611365611ed0565b61137060cf83612666565b82906113925760405163473978bf60e01b8152600401610c7b91815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156114a5578382905f5260205f2001805461141a906142d6565b80601f0160208091040260200160405190810160405280929190818152602001828054611446906142d6565b80156114915780601f1061146857610100808354040283529160200191611491565b820191905f5260205f20905b81548152906001019060200180831161147457829003601f168201915b5050505050815260200190600101906113fd565b505050508152505090505f5f6114c883602001518460400151856060015161267d565b915091508181906114ec5760405162461bcd60e51b8152600401610c7b919061408d565b506114f860cf8661282b565b505f85815260ce6020526040812061150f916135eb565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906115546004830182613622565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738660600151426040516115b4929190918252602082015260400190565b60405180910390a4600193505050506115d960015f5160206147d15f395f51905f5255565b919050565b5f610c138788888888888861123e565b6115f6611d4c565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156116835750825b90505f826001600160401b0316600114801561169e5750303b155b9050811580156116ac575080155b156116ca5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156116f457845460ff60401b1916600160401b1785555b6116fe8787612836565b60fe80546001600160a01b0319166001600160a01b038a16179055831561175f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f61177382610e53565b8015610bed5750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f6117a4610e90565b90505f81516001600160401b038111156117c0576117c0613708565b6040519080825280602002602001820160405280156117e9578160200160208202803683370190505b5090505f5b82518110156118785760335f84838151811061180c5761180c6141a5565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b0316828281518110611858576118586141a5565b6001600160a01b03909216602092830291909101909101526001016117ee565b5092915050565b5f610b095f61241b565b5f6065546001610b0991906142c3565b5f6118a2611ea0565b866118ac81611769565b81906118d757604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610c7b565b506118e133611bd8565b339061190c576040516338905e7160e11b81526001600160a01b039091166004820152602401610c7b565b50611915611ed0565b5f61191e610bf3565b9050808a80821461194b57604051634982205b60e11b815260048101929092526024820152604401610c7b565b50505f7f35c408b4dafbe32b291806a386c3fe69830d0a7fc55f5b2767772398a035e5318b8b8b8b8b8b60405160200161198b9796959493929190614484565b6040516020818303038152906040528051906020012090506119ad81866128e2565b8b906119cf576040516303cac6fb60e11b8152600401610c7b91815260200190565b506119de606680546001019055565b5f5f6119eb8c8c8c61267d565b915091508115611a51578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611a44929190918252602082015260400190565b60405180910390a4611b77565b611a5c60cf8e612a2c565b8d90611a7e576040516368db995b60e11b8152600401610c7b91815260200190565b505f8d815260ce60205260409020611a96828261451c565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611ad991906145d6565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611b49926004850192019061363d565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b6001955050505050610d1260015f5160206147d15f395f51905f5255565b611b9d611d4c565b6001600160a01b038116611bc657604051631e4fbdf760e01b81525f6004820152602401610c7b565b610b1981612430565b5f610b09612a37565b5f610bed609783612392565b5f81815260ce60205260409020805460609190611c00906142d6565b80601f0160208091040260200160405190810160405280929190818152602001828054611c2c906142d6565b8015611c775780601f10611c4e57610100808354040283529160200191611c77565b820191905f5260205f20905b815481529060010190602001808311611c5a57829003601f168201915b50505050509050919050565b611c8b611d4c565b8015611ccd57611c9c609783612a40565b8290611cc75760405163082cdf5560e21b81526001600160a01b039091166004820152602401610c7b565b50611d05565b611cd8609783612a54565b8290611d035760405163e491024560e01b81526001600160a01b039091166004820152602401610c7b565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b33611d7e7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610c3f5760405163118cdaa760e01b8152336004820152602401610c7b565b611db081610e53565b8015611dd357506001600160a01b0381165f9081526034602052604090205460ff165b8190611dfe576040516340da71e560e01b81526001600160a01b039091166004820152602401610c7b565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b611e4f612a68565b5f5160206147b15f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001611634565b5f5160206147b15f395f51905f525460ff1615610c3f5760405163d93c066560e01b815260040160405180910390fd5b5f5160206147d15f395f51905f52805460011901611f0157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f90819081906001600160a01b0316611f2b57505f91508190506001611fd8565b60ca54604051634594558f60e11b81526001600160a01b038981166004830152602482018990525f928392911690638b28ab1e906044016040805180830381865afa158015611f7c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fa0919061430e565b9150915081871015611fb95790935091505f9050611fd8565b80861015611fce5790935091505f9050611fd8565b9093509150600190505b9450945094915050565b611ff7888887611ff287896142c3565b612a97565b61200f612002611889565b8989898989898989612aea565b61175f606580546001019055565b60015f5160206147d15f395f51905f5255565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806120b657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166120aa5f5160206147915f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610c3f5760405163703e46dd60e11b815260040160405180910390fd5b610b19611d4c565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612136575060408051601f3d908101601f191682019092526121339181019061421c565b60015b61215e57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610c7b565b5f5160206147915f395f51905f52811461218e57604051632a87526960e21b815260048101829052602401610c7b565b6121988383612b5a565b505050565b5f610a978383612baf565b82516041146121f95760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207065726d6974207369676e617475726500000000000000006044820152606401610c7b565b5f5f5f602086015192506040860151915060608601515f1a90508c6001600160a01b031663d505accf8d308b8d8f61223191906142c3565b61223b91906142c3565b6040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606481018a905260ff8416608482015260a4810186905260c4810185905260e4015f604051808303815f87803b1580156122a5575f5ffd5b505af11580156122b7573d5f5f3e3d5ffd5b505050506122cb8d8d8d8d8d8d8b8b611fe2565b50505050505050505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c3f5760405163703e46dd60e11b815260040160405180910390fd5b61232c82612bd5565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610a97565b6123bc81612c82565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce9101612386565b5f610bed825490565b60605f610a9783612d2f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6124a981611769565b81906124d4576040516340da71e560e01b81526001600160a01b039091166004820152602401610c7b565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b612528611ea0565b5f5160206147b15f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611e88565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206147715f395f51905f52916125a6906142d6565b80601f01602080910402602001604051908101604052809291908181526020018280546125d2906142d6565b801561261d5780601f106125f45761010080835404028352916020019161261d565b820191905f5260205f20905b81548152906001019060200180831161260057829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206147715f395f51905f52916125a6906142d6565b5f8181526001830160205260408120541515610a97565b5f606082156128235760fe546001600160a01b03908116908616036126aa576126a76064846145e2565b92505b60405163a9059cbb60e01b81526001600160a01b0385811660048301526024820185905286169063a9059cbb906044016020604051808303815f875af1925050508015612714575060408051601f3d908101601f1916820190925261271191810190614601565b60015b6127c557612720614620565b806308c379a0036127495750612734614638565b8061273f575061278c565b5f92509050612823565b634e487b710361278c5761275b6146ba565b90612766575061278c565b60408051602081018390525f945001604051602081830303815290604052915050612823565b3d8080156127b5576040519150601f19603f3d011682016040523d82523d5f602084013e6127ba565b606091505b505f92509050612823565b80156127e5576001925060405180602001604052805f8152509150612821565b5f92506040518060400160405280601f81526020017f427269646765457468657265756d3a207472616e73666572206661696c65640081525091505b505b935093915050565b5f610a978383612d87565b61283e612e6a565b61284733612eb3565b61284f612ec4565b612857612ecc565b61285f612edc565b612867612eec565b6001600160a01b0382166128ad57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610c7b565b4260cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff168110156128fd575f915050610bed565b5f80826001600160401b0381111561291757612917613708565b604051908082528060200260200182016040528015612940578160200160208202803683370190505b5090505f5b8551811015612a1a575f868281518110612961576129616141a5565b60200260200101519050604181511015612982575f95505050505050610bed565b5f612996826129908b612f4b565b90612f77565b90506129a181611bd8565b6129b3575f9650505050505050610bed565b5f805b8551811015612a02578581815181106129d1576129d16141a5565b60200260200101516001600160a01b0316836001600160a01b0316036129fa5760019150612a02565b6001016129b6565b5080612a0f578560010195505b505050600101612945565b505060995460ff161115949350505050565b5f610a978383612f9f565b5f610b09612feb565b5f610a97836001600160a01b038416612f9f565b5f610a97836001600160a01b038416612d87565b5f5160206147b15f395f51905f525460ff16610c3f57604051638dfc202b60e01b815260040160405180910390fd5b612ab88330612aa684866142c3565b6001600160a01b03881692919061305e565b8015612ae457612ae4612ad360cc546001600160a01b031690565b6001600160a01b03861690836130c5565b50505050565b6001600160a01b038881165f81815260336020526040902054828a16928c917fd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c91168a8a8a8a428b8b604051612b479897969594939291906146d7565b60405180910390a4505050505050505050565b612b63826130f6565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612ba7576121988282613159565b610b526131cb565b5f825f018281548110612bc457612bc46141a5565b905f5260205f200154905092915050565b806001600160a01b038116612c15576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610c7b565b612c1f5f83612a40565b8290612c4a576040516351eccfe160e11b81526001600160a01b039091166004820152602401610c7b565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116612cc2576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610c7b565b612ccc5f83612a54565b8290612cf7576040516340da71e560e01b81526001600160a01b039091166004820152602401610c7b565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611c7757602002820191905f5260205f20905b815481526020019060010190808311612d685750505050509050919050565b5f8181526001830160205260408120548015612e61575f612da960018361471f565b85549091505f90612dbc9060019061471f565b9050808214612e1b575f865f018281548110612dda57612dda6141a5565b905f5260205f200154905080875f018481548110612dfa57612dfa6141a5565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612e2c57612e2c614732565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610bed565b5f915050610bed565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610c3f57604051631afcd79f60e31b815260040160405180910390fd5b612ebb612e6a565b610b19816131ea565b610c3f612e6a565b612ed4612e6a565b610c3f6131f2565b612ee4612e6a565b610c3f613212565b612ef4612e6a565b612f3c604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061321a565b6099805460ff19166003179055565b5f610bed612f57612a37565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612f85868661322c565b925092509250612f958282613275565b5090949350505050565b5f818152600183016020526040812054612fe457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610bed565b505f610bed565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61301561332d565b61301d613395565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6040516001600160a01b038481166024830152838116604483015260648201839052612ae49186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506133d7565b6040516001600160a01b0383811660248301526044820183905261219891859182169063a9059cbb90606401613093565b806001600160a01b03163b5f0361312b57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610c7b565b5f5160206147915f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516131759190614746565b5f60405180830381855af49150503d805f81146131ad576040519150601f19603f3d011682016040523d82523d5f602084013e6131b2565b606091505b50915091506131c2858383613443565b95945050505050565b3415610c3f5760405163b398979f60e01b815260040160405180910390fd5b611b9d612e6a565b6131fa612e6a565b5f5160206147b15f395f51905f52805460ff19169055565b61201d612e6a565b613222612e6a565b610b52828261349f565b5f5f5f8351604103613263576020840151604085015160608601515f1a613255888285856134fe565b95509550955050505061326e565b505081515f91506002905b9250925092565b5f8260038111156132885761328861475c565b03613291575050565b60018260038111156132a5576132a561475c565b036132c35760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156132d7576132d761475c565b036132f85760405163fce698f760e01b815260048101829052602401610c7b565b600382600381111561330c5761330c61475c565b03610b52576040516335e2f38360e21b815260048101829052602401610c7b565b5f5f5160206147715f395f51905f5281613345612568565b80519091501561335d57805160209091012092915050565b8154801561336c579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206147715f395f51905f52816133ad612628565b8051909150156133c557805160209091012092915050565b6001820154801561336c579392505050565b5f5f60205f8451602086015f885af1806133f6576040513d5f823e3d81fd5b50505f513d9150811561340d57806001141561341a565b6001600160a01b0384163b155b15612ae457604051635274afe760e01b81526001600160a01b0385166004820152602401610c7b565b60608261345857613453826135c2565b610a97565b815115801561346f57506001600160a01b0384163b155b1561349857604051639996b31560e01b81526001600160a01b0385166004820152602401610c7b565b5080610a97565b6134a7612e6a565b5f5160206147715f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026134e0848261451c565b50600381016134ef838261451c565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561353757505f91506003905082611fd8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613588573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166135b357505f925060019150829050611fd8565b975f9750879650945050505050565b8051156135d25780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5080546135f7906142d6565b5f825580601f10613606575050565b601f0160209004905f5260205f2090810190610b199190613691565b5080545f8255905f5260205f2090810190610b1991906136a5565b828054828255905f5260205f20908101928215613681579160200282015b828111156136815782518290613671908261451c565b509160200191906001019061365b565b5061368d9291506136a5565b5090565b5b8082111561368d575f8155600101613692565b8082111561368d575f6136b882826135eb565b506001016136a5565b5f5f83601f8401126136d1575f5ffd5b5081356001600160401b038111156136e7575f5ffd5b6020830191508360208260051b8501011115613701575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f191681016001600160401b038111828210171561374157613741613708565b6040525050565b5f6001600160401b0382111561376057613760613708565b5060051b60200190565b5f82601f830112613779575f5ffd5b81356001600160401b0381111561379257613792613708565b6040516137a9601f8301601f19166020018261371c565b8181528460208386010111156137bd575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6137e383613748565b6040516137f0828261371c565b84815291505060208101600584901b83018581111561380d575f5ffd5b835b818110156138475780356001600160401b0381111561382c575f5ffd5b6138388882880161376a565b8452506020928301920161380f565b5050509392505050565b5f82601f830112613860575f5ffd5b610a97838335602085016137d9565b5f5f5f60408486031215613881575f5ffd5b83356001600160401b03811115613896575f5ffd5b6138a2868287016136c1565b90945092505060208401356001600160401b038111156138c0575f5ffd5b8401601f810186136138d0575f5ffd5b80356138db81613748565b6040516138e8828261371c565b80915082815260208101915060208360051b85010192508883111561390b575f5ffd5b602084015b8381101561394b5780356001600160401b0381111561392d575f5ffd5b61393c8b602083890101613851565b84525060209283019201613910565b50809450505050509250925092565b6001600160a01b0381168114610b19575f5ffd5b80356115d98161395a565b5f60208284031215613989575f5ffd5b8135610a978161395a565b5f602082840312156139a4575f5ffd5b81356001600160401b038111156139b9575f5ffd5b8201601f810184136139c9575f5ffd5b80356139d481613748565b6040516139e1828261371c565b80915082815260208101915060208360051b850101925086831115613a04575f5ffd5b6020840193505b82841015613a2f578335613a1e8161395a565b825260209384019390910190613a0b565b9695505050505050565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610bed565b5f5f5f5f5f5f60a08789031215613a78575f5ffd5b8635613a838161395a565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613ab2575f5ffd5b613abe89828a016136c1565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613ae6575f5ffd5b8735613af18161395a565b96506020880135613b018161395a565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115613b30575f5ffd5b613b3c8a828b016136c1565b989b979a50959850939692959293505050565b5f5f60408385031215613b60575f5ffd5b8235613b6b8161395a565b915060208301356001600160401b03811115613b85575f5ffd5b613b918582860161376a565b9150509250929050565b5f60208284031215613bab575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f5f5f6101208b8d031215613bcc575f5ffd5b8a35613bd78161395a565b9950613be560208c0161396e565b9850613bf360408c0161396e565b975060608b0135965060808b0135955060a08b0135945060c08b0135935060e08b01356001600160401b03811115613c29575f5ffd5b613c358d828e0161376a565b9350506101008b01356001600160401b03811115613c51575f5ffd5b613c5d8d828e016136c1565b915080935050809150509295989b9194979a5092959850565b5f5f60408385031215613c87575f5ffd5b8235613c928161395a565b91506020830135613ca28161395a565b809150509250929050565b5f5f5f5f5f5f5f5f5f6101008a8c031215613cc6575f5ffd5b8935613cd18161395a565b985060208a0135613ce18161395a565b975060408a0135965060608a0135955060808a0135945060a08a0135935060c08a01356001600160401b03811115613d17575f5ffd5b613d238c828d0161376a565b93505060e08a01356001600160401b03811115613d3e575f5ffd5b613d4a8c828d016136c1565b915080935050809150509295985092959850929598565b602080825282518282018190525f918401906040840190835b81811015613da15783516001600160a01b0316835260209384019390920191600101613d7a565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b81811015613e735760df19878603018352613e5e858551613dac565b94506020938401939290920191600101613e42565b50929695505050505050565b60ff60f81b8816815260e060208201525f613e9d60e0830189613dac565b8281036040840152613eaf8189613dac565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015613f04578351835260209384019390920191600101613ee6565b50909b9a5050505050505050505050565b5f5f60408385031215613f26575f5ffd5b8235613f318161395a565b946020939093013593505050565b5f5f5f5f5f5f5f60c0888a031215613f55575f5ffd5b8735613f608161395a565b96506020880135613f708161395a565b9550604088013594506060880135935060808801356001600160401b03811115613f98575f5ffd5b613fa48a828b0161376a565b93505060a08801356001600160401b03811115613b30575f5ffd5b602080825282518282018190525f918401906040840190835b81811015613da15761400883855180516001600160a01b0316825260208082015190830152604090810151910152565b6020939093019260609290920191600101613fd8565b5f5f5f5f5f5f60a08789031215614033575f5ffd5b863561403e8161395a565b9550602087013594506040870135935060608701356001600160401b03811115614066575f5ffd5b61407289828a0161376a565b93505060808701356001600160401b03811115613ab2575f5ffd5b602081525f610a976020830184613dac565b5f602082840312156140af575f5ffd5b813560ff81168114610a97575f5ffd5b5f5f5f606084860312156140d1575f5ffd5b83356140dc8161395a565b925060208401356140ec8161395a565b915060408401356140fc8161395a565b809150509250925092565b5f5f5f5f5f5f5f60c0888a03121561411d575f5ffd5b87359650602088013561412f8161395a565b9550604088013561413f8161395a565b94506060880135935060808801356001600160401b03811115614160575f5ffd5b61416c8a828b016136c1565b90945092505060a08801356001600160401b0381111561418a575f5ffd5b6141968a828b01613851565b91505092959891949750929550565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e198336030181126141cd575f5ffd5b9190910192915050565b5f5f8335601e198436030181126141ec575f5ffd5b8301803591506001600160401b03821115614205575f5ffd5b6020019150600581901b3603821315613701575f5ffd5b5f6020828403121561422c575f5ffd5b5051919050565b5f60608284031215614243575f5ffd5b604051606081018181106001600160401b038211171561426557614265613708565b806040525080915082516142788161395a565b815260208381015190820152604092830151920191909152919050565b5f606082840312156142a5575f5ffd5b610a978383614233565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610bed57610bed6142af565b600181811c908216806142ea57607f821691505b60208210810361430857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f6040838503121561431f575f5ffd5b505080516020909101519092909150565b5f60208284031215614340575f5ffd5b81516001600160401b03811115614355575f5ffd5b8201601f81018413614365575f5ffd5b805161437081613748565b60405161437d828261371c565b828152602060609093028401830192810191508683111561439c575f5ffd5b6020840193505b82841015613a2f576143b58785614233565b82526020820191506060840193506143a3565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b8681101561447857838303601f19018852813536879003601e1901811261442c575f5ffd5b86016020810190356001600160401b03811115614447575f5ffd5b803603821315614455575f5ffd5b6144608582846143c8565b60209a8b019a90955093909301925050600101614407565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f906144c490830184866143f0565b9998505050505050505050565b601f82111561219857805f5260205f20601f840160051c810160208510156144f65750805b601f840160051c820191505b81811015614515575f8155600101614502565b5050505050565b81516001600160401b0381111561453557614535613708565b6145498161454384546142d6565b846144d1565b6020601f82116001811461457b575f83156145645750848201515b5f19600385901b1c1916600184901b178455614515565b5f84815260208120601f198516915b828110156145aa578785015182556020948501946001909201910161458a565b50848210156145c757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610a973684846137d9565b5f826145fc57634e487b7160e01b5f52601260045260245ffd5b500490565b5f60208284031215614611575f5ffd5b81518015158114610a97575f5ffd5b5f60033d1115610e3e5760045f5f3e505f5160e01c90565b5f60443d10156146455790565b6040513d600319016004823e80513d60248201116001600160401b038211171561466e57505090565b80820180516001600160401b03811115614689575050505090565b3d84016003190182820160200111156146a3575050505090565b6146b26020828501018561371c565b509392505050565b5f5f60233d11156146d357602060045f3e50505f516001905b9091565b6001600160a01b0389811682528816602082015260408101879052606081018690526080810185905260a0810184905260e060c082018190525f90610e8290830184866143f0565b81810381811115610bed57610bed6142af565b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f920191825250919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220e93aa279f355912b5315401180127fbb450d7fa2aac5be75f930b73dd3c2695d64736f6c634300081c0033",
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

// BridgeCross is a paid mutator transaction binding the contract method 0xaaa78020.
//
// Solidity: function bridgeCross(address account, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeCross(opts *bind.TransactOpts, account common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeCross", account, value, deadline, permitSig, extraData)
}

// BridgeCross is a paid mutator transaction binding the contract method 0xaaa78020.
//
// Solidity: function bridgeCross(address account, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeCross(account common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeCross(&_BridgeEthereum.TransactOpts, account, value, deadline, permitSig, extraData)
}

// BridgeCross is a paid mutator transaction binding the contract method 0xaaa78020.
//
// Solidity: function bridgeCross(address account, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeCross(account common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeCross(&_BridgeEthereum.TransactOpts, account, value, deadline, permitSig, extraData)
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

// BridgeToCross is a paid mutator transaction binding the contract method 0x95d791e9.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) BridgeToCross(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "bridgeToCross", from, to, value, deadline, permitSig, extraData)
}

// BridgeToCross is a paid mutator transaction binding the contract method 0x95d791e9.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) BridgeToCross(from common.Address, to common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeToCross(&_BridgeEthereum.TransactOpts, from, to, value, deadline, permitSig, extraData)
}

// BridgeToCross is a paid mutator transaction binding the contract method 0x95d791e9.
//
// Solidity: function bridgeToCross(address from, address to, uint256 value, uint256 deadline, bytes permitSig, bytes[] extraData) returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) BridgeToCross(from common.Address, to common.Address, value *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.BridgeToCross(&_BridgeEthereum.TransactOpts, from, to, value, deadline, permitSig, extraData)
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

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridge", token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridge(&_BridgeEthereum.TransactOpts, token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeEthereum *BridgeEthereumTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeEthereum.Contract.PermitBridgeTo(&_BridgeEthereum.TransactOpts, token, from, to, value, gas, service, deadline, permitSig, extraData)
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
	Index      *big.Int
	Token      common.Address
	PairToken  common.Address
	From       common.Address
	To         common.Address
	Value      *big.Int
	GasFee     *big.Int
	ServiceFee *big.Int
	Time       *big.Int
	ExtraData  [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0xd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 serviceFee, uint256 time, bytes[] extraData)
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

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0xd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 serviceFee, uint256 time, bytes[] extraData)
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0xd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 serviceFee, uint256 time, bytes[] extraData)
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
