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

// BridgeCrossMetaData contains all meta data concerning the BridgeCross contract.
var BridgeCrossMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BridgeFeeManagerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addTokenDeploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFeeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"feeInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossMintableERC20Code\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xcross\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInsufficientValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e3a446eb": "BridgeFeeManager()",
		"114ee544": "BridgeFeeManagerLength()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"f7684ec4": "addTokenDeploy(address,string,uint8)",
		"9d0dc76e": "allFeeInfo()",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"8b28ab1e": "calculateFee(address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
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
		"211c8ac0": "xcross()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614bf76100f95f395f81816121150152818161213e01526123bf0152614bf75ff3fe6080604052600436106102f2575f3560e01c8063751b4c9c11610189578063c0c53b8b116100d8578063f120c40011610092578063f7684ec41161006d578063f7684ec4146108bf578063facd743b146108de578063fb75b2c7146108fd578063fe2b8da61461091a575f5ffd5b8063f120c40014610879578063f2fde38b1461088c578063f698da25146108ab575f5ffd5b8063c0c53b8b146107e0578063c1876453146107ff578063c97682f81461081e578063d92fc67b14610832578063e1af7f5014610846578063e3a446eb1461085a575f5ffd5b806391cf6d3e116101435780639d0dc76e1161011e5780639d0dc76e14610744578063a982318314610765578063ad3cb1cc14610784578063b7f3358d146107c1575f5ffd5b806391cf6d3e146106fd5780639300c9261461071157806396ce079514610730575f5ffd5b8063751b4c9c146106145780637c41ad2c146106335780638456cb591461065257806384b0196e146106665780638b28ab1e1461068d5780638da5cb5b146106c1575f5ffd5b80634f1ef286116102455780635dbe47e8116101ff5780636ff97f1d116101da5780636ff97f1d1461057c5780637021fd0e1461059d578063715018a6146105c957806371c59d7b146105dd575f5ffd5b80635dbe47e81461052b5780635fa7b5841461054a57806369a3318b14610569575f5ffd5b80634f1ef286146104905780634f6ccce7146104a357806350d6fb48146104c257806352d1902d146104d55780635476bd72146104e95780635c975abb14610508575f5ffd5b806327938c81116102b05780633f4ba83a1161028b5780633f4ba83a146104295780633f5d3b5d1461043d57806340a141ff1461045057806342cde4e81461046f575f5ffd5b806327938c81146103e357806337d10075146103f75780633b3bff0f1461040a575f5ffd5b80628bd028146102f6578063114ee5441461031e5780631327d3d8146103405780631d40f0d814610361578063211c8ac014610380578063252154fa146103b7575b5f5ffd5b610309610304366004613ba1565b610939565b60405190151581526020015b60405180910390f35b348015610329575f5ffd5b50610332610a65565b604051908152602001610315565b34801561034b575f5ffd5b5061035f61035a366004613cab565b610ad5565b005b34801561036c575f5ffd5b5061035f61037b366004613cc6565b610ae3565b34801561038b575f5ffd5b5060fe5461039f906001600160a01b031681565b6040516001600160a01b039091168152602001610315565b3480156103c2575f5ffd5b506103d66103d1366004613cab565b610b1d565b6040516103159190613d6b565b3480156103ee575f5ffd5b50610332610bba565b610309610405366004613d95565b610bca565b348015610415575f5ffd5b5061035f610424366004613cab565b610be5565b348015610434575f5ffd5b5061035f610bf6565b61030961044b366004613e02565b610c08565b34801561045b575f5ffd5b5061035f61046a366004613cab565b610ce5565b34801561047a575f5ffd5b5060995460405160ff9091168152602001610315565b61035f61049e366004613e81565b610cef565b3480156104ae575f5ffd5b5061039f6104bd366004613ecd565b610d0a565b6103096104d0366004613ee4565b610d15565b3480156104e0575f5ffd5b50610332610dec565b3480156104f4575f5ffd5b5061035f610503366004613fa8565b610e08565b348015610513575f5ffd5b505f516020614b825f395f51905f525460ff16610309565b348015610536575f5ffd5b50610309610545366004613cab565b610e1a565b348015610555575f5ffd5b5061035f610564366004613cab565b610e25565b610309610577366004613fdf565b610e36565b348015610587575f5ffd5b50610590610e57565b6040516103159190614093565b3480156105a8575f5ffd5b506105bc6105b7366004613ecd565b610eb8565b604051610315919061410c565b3480156105d4575f5ffd5b5061035f611018565b3480156105e8575f5ffd5b5061039f6105f7366004613cab565b6001600160a01b039081165f908152603360205260409020541690565b34801561061f575f5ffd5b506103d661062e366004613ecd565b611029565b34801561063e575f5ffd5b5061035f61064d366004613cab565b611084565b34801561065d575f5ffd5b5061035f611095565b348015610671575f5ffd5b5061067a6110a5565b60405161031597969594939291906141b1565b348015610698575f5ffd5b506106ac6106a7366004614247565b61114e565b60408051928352602083019190915201610315565b3480156106cc575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661039f565b348015610708575f5ffd5b5060cb54610332565b34801561071c575f5ffd5b5061035f61072b366004613cc6565b6111ce565b34801561073b575f5ffd5b50610332611205565b34801561074f575f5ffd5b5061075861124c565b6040516103159190614271565b348015610770575f5ffd5b5061030961077f366004613ecd565b6112ba565b34801561078f575f5ffd5b506107b4604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161031591906142d0565b3480156107cc575f5ffd5b5061035f6107db3660046142f2565b611579565b3480156107eb575f5ffd5b5061035f6107fa36600461430b565b6115ca565b34801561080a575f5ffd5b50610309610819366004613cab565b611703565b348015610829575f5ffd5b50610590611733565b34801561083d575f5ffd5b50610332611819565b348015610851575f5ffd5b50610332611823565b348015610865575f5ffd5b5060ca5461039f906001600160a01b031681565b610309610887366004614353565b611833565b348015610897575f5ffd5b5061035f6108a6366004613cab565b611b2f565b3480156108b6575f5ffd5b50610332611b69565b3480156108ca575f5ffd5b5061039f6108d93660046143f1565b611b72565b3480156108e9575f5ffd5b506103096108f8366004613cab565b611cb2565b348015610908575f5ffd5b5060cc546001600160a01b031661039f565b348015610925575f5ffd5b506107b4610934366004613ecd565b611cbe565b5f805b83811015610a5857610a4f8585838181106109595761095961445f565b905060200281019061096b9190614473565b3586868481811061097e5761097e61445f565b90506020028101906109909190614473565b6109a1906040810190602001613cab565b8787858181106109b3576109b361445f565b90506020028101906109c59190614473565b6109d6906060810190604001613cab565b8888868181106109e8576109e861445f565b90506020028101906109fa9190614473565b60600135898987818110610a1057610a1061445f565b9050602002810190610a229190614473565b610a30906080810190614491565b898881518110610a4257610a4261445f565b6020026020010151611833565b5060010161093c565b50600190505b9392505050565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa158015610aac573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad091906144d6565b905090565b610ae0816001611d5d565b50565b5f5b8151811015610b1957610b11828281518110610b0357610b0361445f565b60200260200101515f611d5d565b600101610ae5565b5050565b610b4760405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631290aa7d60e11b81526001600160a01b0384811660048301529091169063252154fa906024015b606060405180830381865afa158015610b90573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb4919061454f565b92915050565b5f6066546001610ad0919061457d565b5f610bda87338888888888610c08565b979650505050505050565b610bed611e26565b610ae081611e81565b610bfe611e26565b610c06611f21565b565b5f610c11611f7a565b87610c1b81611703565b8190610c4b57604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610c54611faa565b5f5f5f610c638c8b8b8b611fe1565b9194509250905082828a8a84610ca2576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610c42565b50505050610cbc8c610cb13390565b8d8d87878d8d6120bc565b60019450505050610cd960015f516020614ba25f395f51905f5255565b50979650505050505050565b610ae0815f611d5d565b610cf761210a565b610d00826121ae565b610b1982826121b6565b5f610bb48183612277565b5f610d1e611f7a565b8a610d2881611703565b8190610d5357604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610c42565b50610d5c611faa565b5f5f5f610d6b8f8d8d8d611fe1565b9194509250905082828c8c84610daa576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610c42565b50505050610dc08f8f8f8f87878f8f8f8f612282565b60019450505050610ddd60015f516020614ba25f395f51905f5255565b509a9950505050505050505050565b5f610df56123b4565b505f516020614b625f395f51905f525b90565b610e10611e26565b610b1982826123fd565b5f610bb4818361246c565b610e2d611e26565b610ae08161248d565b5f610e498a8a8b8b8b8b8b8b8b8b610d15565b9a9950505050505050505050565b60605f610e635f6124f5565b9050806001600160401b03811115610e7d57610e7d613a2c565b604051908082528060200260200182016040528015610ea6578160200160208202803683370190505b509150610eb25f6124fe565b91505090565b610ef86040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b8282101561100a578382905f5260205f20018054610f7f90614590565b80601f0160208091040260200160405190810160405280929190818152602001828054610fab90614590565b8015610ff65780601f10610fcd57610100808354040283529160200191610ff6565b820191905f5260205f20905b815481529060010190602001808311610fd957829003601f168201915b505050505081526020019060010190610f62565b505050915250909392505050565b611020611e26565b610c065f61250a565b61105360405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631d46d32760e21b8152600481018490526001600160a01b039091169063751b4c9c90602401610b75565b61108c611e26565b610ae08161257a565b61109d611e26565b610c066125fa565b5f60608082808083815f516020614b425f395f51905f5280549091501580156110d057506001810154155b6111145760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610c42565b61111c612642565b611124612702565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60ca54604051634594558f60e11b81526001600160a01b038481166004830152602482018490525f928392911690638b28ab1e906044016040805180830381865afa15801561119f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c391906145c8565b909590945092505050565b5f5b8151811015610b19576111fd8282815181106111ee576111ee61445f565b60200260200101516001611d5d565b6001016111d0565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015610aac573d5f5f3e3d5ffd5b60ca5460408051634e86e3b760e11b815290516060926001600160a01b031691639d0dc76e916004808301925f9291908290030181865afa158015611293573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ad091908101906145ea565b5f6112c3611f7a565b6112cc33611cb2565b33906112f7576040516338905e7160e11b81526001600160a01b039091166004820152602401610c42565b50611300611faa565b61130b60cf83612740565b829061132d5760405163473978bf60e01b8152600401610c4291815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b82821015611440578382905f5260205f200180546113b590614590565b80601f01602080910402602001604051908101604052809291908181526020018280546113e190614590565b801561142c5780601f106114035761010080835404028352916020019161142c565b820191905f5260205f20905b81548152906001019060200180831161140f57829003601f168201915b505050505081526020019060010190611398565b505050508152505090505f5f611463836020015184604001518560600151612757565b915091508181906114875760405162461bcd60e51b8152600401610c4291906142d0565b5061149360cf8661291a565b505f85815260ce602052604081206114aa9161390f565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906114ef6004830182613946565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a37386606001514260405161154f929190918252602082015260400190565b60405180910390a46001935050505061157460015f516020614ba25f395f51905f5255565b919050565b611581611e26565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561160e5750825b90505f826001600160401b031660011480156116295750303b155b905081158015611637575080155b156116555760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561167f57845460ff60401b1916600160401b1785555b6116898787612925565b60fe805460016001600160a01b03199182161790915560ff80549091166001600160a01b038a1617905583156116f957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f61170d82610e1a565b8015610bb45750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f61173e610e57565b90505f81516001600160401b0381111561175a5761175a613a2c565b604051908082528060200260200182016040528015611783578160200160208202803683370190505b5090505f5b82518110156118125760335f8483815181106117a6576117a661445f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b03168282815181106117f2576117f261445f565b6001600160a01b0390921660209283029190910190910152600101611788565b5092915050565b5f610ad05f6124f5565b5f6065546001610ad0919061457d565b5f61183c611f7a565b8661184681611703565b819061187157604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610c42565b5061187b33611cb2565b33906118a6576040516338905e7160e11b81526001600160a01b039091166004820152602401610c42565b506118af611faa565b5f6118b8610bba565b9050808a8082146118e557604051634982205b60e11b815260048101929092526024820152604401610c42565b50505f7f35c408b4dafbe32b291806a386c3fe69830d0a7fc55f5b2767772398a035e5318b8b8b8b8b8b604051602001611925979695949392919061473e565b60405160208183030381529060405280519060200120905061194781866129d1565b8b90611969576040516303cac6fb60e11b8152600401610c4291815260200190565b50611978606680546001019055565b5f5f6119858c8c8c612757565b9150915081156119eb578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d426040516119de929190918252602082015260400190565b60405180910390a4611b11565b6119f660cf8e612b1b565b8d90611a18576040516368db995b60e11b8152600401610c4291815260200190565b505f8d815260ce60205260409020611a3082826147d6565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611a739190614890565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611ae39260048501920190613961565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b6001955050505050610cd960015f516020614ba25f395f51905f5255565b611b37611e26565b6001600160a01b038116611b6057604051631e4fbdf760e01b81525f6004820152602401610c42565b610ae08161250a565b5f610ad0612b26565b5f611b7b611e26565b5f83604051602001611b8d91906148b3565b60408051601f19818403018152908290526bffffffffffffffffffffffff19606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f60ff5f9054906101000a90046001600160a01b03166001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611c26573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611c4d91908101906148d4565b838787604051602001611c6293929190614953565b60408051601f1981840301815290829052611c80929160200161498b565b6040516020818303038152906040529050611c9c5f8383612b2f565b9350611ca88488610e08565b5050509392505050565b5f610bb460978361246c565b5f81815260ce60205260409020805460609190611cda90614590565b80601f0160208091040260200160405190810160405280929190818152602001828054611d0690614590565b8015611d515780601f10611d2857610100808354040283529160200191611d51565b820191905f5260205f20905b815481529060010190602001808311611d3457829003601f168201915b50505050509050919050565b611d65611e26565b8015611da757611d76609783612bc3565b8290611da15760405163082cdf5560e21b81526001600160a01b039091166004820152602401610c42565b50611ddf565b611db2609783612bd7565b8290611ddd5760405163e491024560e01b81526001600160a01b039091166004820152602401610c42565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b33611e587f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610c065760405163118cdaa760e01b8152336004820152602401610c42565b611e8a81610e1a565b8015611ead57506001600160a01b0381165f9081526034602052604090205460ff165b8190611ed8576040516340da71e560e01b81526001600160a01b039091166004820152602401610c42565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b611f29612beb565b5f516020614b825f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016115bf565b5f516020614b825f395f51905f525460ff1615610c065760405163d93c066560e01b815260040160405180910390fd5b5f516020614ba25f395f51905f52805460011901611fdb57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f90819081906001600160a01b031661200557505f915081905060016120b2565b60ca54604051634594558f60e11b81526001600160a01b038981166004830152602482018990525f928392911690638b28ab1e906044016040805180830381865afa158015612056573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061207a91906145c8565b91509150818710156120935790935091505f90506120b2565b808610156120a85790935091505f90506120b2565b9093509150600190505b9450945094915050565b6120d18888876120cc878961457d565b612c1a565b6120e96120dc611823565b8989898989898989612dc0565b6116f9606580546001019055565b60015f516020614ba25f395f51905f5255565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061219057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166121845f516020614b625f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610c065760405163703e46dd60e11b815260040160405180910390fd5b610ae0611e26565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612210575060408051601f3d908101601f1916820190925261220d918101906144d6565b60015b61223857604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610c42565b5f516020614b625f395f51905f52811461226857604051632a87526960e21b815260048101829052602401610c42565b6122728383612e30565b505050565b5f610a5e8383612e85565b82516041146122d35760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207065726d6974207369676e617475726500000000000000006044820152606401610c42565b5f5f5f602086015192506040860151915060608601515f1a90508c6001600160a01b031663d505accf8d308b8d8f61230b919061457d565b612315919061457d565b6040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606481018a905260ff8416608482015260a4810186905260c4810185905260e4015f604051808303815f87803b15801561237f575f5ffd5b505af1158015612391573d5f5f3e3d5ffd5b505050506123a58d8d8d8d8d8d8b8b6120bc565b50505050505050505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c065760405163703e46dd60e11b815260040160405180910390fd5b61240682612eab565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610a5e565b61249681612f58565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce9101612460565b5f610bb4825490565b60605f610a5e83613005565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61258381611703565b81906125ae576040516340da71e560e01b81526001600160a01b039091166004820152602401610c42565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b612602611f7a565b5f516020614b825f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611f62565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614b425f395f51905f529161268090614590565b80601f01602080910402602001604051908101604052809291908181526020018280546126ac90614590565b80156126f75780601f106126ce576101008083540402835291602001916126f7565b820191905f5260205f20905b8154815290600101906020018083116126da57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614b425f395f51905f529161268090614590565b5f8181526001830160205260408120541515610a5e565b5f606082156129125760fe546001600160a01b0390811690861603612799576127946127846064856149a7565b6001600160a01b0386169061305d565b612912565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528616906340c10f19906044016020604051808303815f875af1925050508015612803575060408051601f3d908101601f19168201909252612800918101906149be565b60015b6128b45761280f6149dd565b806308c379a00361283857506128236149f5565b8061282e575061287b565b5f92509050612912565b634e487b710361287b5761284a614a77565b90612855575061287b565b60408051602081018390525f945001604051602081830303815290604052915050612912565b3d8080156128a4576040519150601f19603f3d011682016040523d82523d5f602084013e6128a9565b606091505b505f92509050612912565b80156128d4576001925060405180602001604052805f8152509150612910565b5f92506040518060400160405280601881526020017f42726964676543726f73733a206d696e74206661696c6564000000000000000081525091505b505b935093915050565b5f610a5e83836130e9565b61292d6131cc565b61293633613215565b61293e613226565b61294661322e565b61294e61323e565b61295661324e565b6001600160a01b03821661299c57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610c42565b4260cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff168110156129ec575f915050610bb4565b5f80826001600160401b03811115612a0657612a06613a2c565b604051908082528060200260200182016040528015612a2f578160200160208202803683370190505b5090505f5b8551811015612b09575f868281518110612a5057612a5061445f565b60200260200101519050604181511015612a71575f95505050505050610bb4565b5f612a8582612a7f8b6132ad565b906132d9565b9050612a9081611cb2565b612aa2575f9650505050505050610bb4565b5f805b8551811015612af157858181518110612ac057612ac061445f565b60200260200101516001600160a01b0316836001600160a01b031603612ae95760019150612af1565b600101612aa5565b5080612afe578560010195505b505050600101612a34565b505060995460ff161115949350505050565b5f610a5e8383613301565b5f610ad061334d565b5f83471015612b5a5760405163cf47918160e01b815247600482015260248101859052604401610c42565b81515f03612b7b57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615612b9c576040513d5f823e3d81fd5b6001600160a01b038116610a5e5760405163b06ebf3d60e01b815260040160405180910390fd5b5f610a5e836001600160a01b038416613301565b5f610a5e836001600160a01b0384166130e9565b5f516020614b825f395f51905f525460ff16610c0657604051638dfc202b60e01b815260040160405180910390fd5b60fe546001600160a01b0390811690851603612cdb57346064612c3d8183614a94565b612c4791906149a7565b143490612c6a576040516308dc47c960e41b8152600401610c4291815260200190565b50612c75818361457d565b3414612c81828461457d565b349091612ca9576040516290c38760e71b815260048101929092526024820152604401610c42565b50508015612cd657612cd681612cc760cc546001600160a01b031690565b6001600160a01b03169061305d565b612dba565b8015612d0957612d0983612cf760cc546001600160a01b031690565b6001600160a01b0387169190846133c0565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201849052851690639dc29fac906044016020604051808303815f875af1158015612d55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d7991906149be565b848484909192612db6576040516336b52daf60e21b81526001600160a01b0393841660048201529290911660248301526044820152606401610c42565b5050505b50505050565b6001600160a01b038881165f81815260336020526040902054828a16928c917fd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c91168a8a8a8a428b8b604051612e1d989796959493929190614ab3565b60405180910390a4505050505050505050565b612e398261341a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612e7d57612272828261347d565b610b196134ef565b5f825f018281548110612e9a57612e9a61445f565b905f5260205f200154905092915050565b806001600160a01b038116612eeb576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610c42565b612ef55f83612bc3565b8290612f20576040516351eccfe160e11b81526001600160a01b039091166004820152602401610c42565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116612f98576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610c42565b612fa25f83612bd7565b8290612fcd576040516340da71e560e01b81526001600160a01b039091166004820152602401610c42565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611d5157602002820191905f5260205f20905b81548152602001906001019080831161303e5750505050509050919050565b804710156130875760405163cf47918160e01b815247600482015260248101829052604401610c42565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f81146130d1576040519150601f19603f3d011682016040523d82523d5f602084013e6130d6565b606091505b509150915081612dba57612dba8161350e565b5f81815260018301602052604081205480156131c3575f61310b600183614afb565b85549091505f9061311e90600190614afb565b905080821461317d575f865f01828154811061313c5761313c61445f565b905f5260205f200154905080875f01848154811061315c5761315c61445f565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061318e5761318e614b0e565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610bb4565b5f915050610bb4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610c0657604051631afcd79f60e31b815260040160405180910390fd5b61321d6131cc565b610ae081613537565b610c066131cc565b6132366131cc565b610c0661353f565b6132466131cc565b610c0661355f565b6132566131cc565b61329e604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613567565b6099805460ff19166003179055565b5f610bb46132b9612b26565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6132e78686613579565b9250925092506132f782826135c2565b5090949350505050565b5f81815260018301602052604081205461334657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610bb4565b505f610bb4565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61337761367a565b61337f6136e2565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612dba908590613724565b806001600160a01b03163b5f0361344f57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610c42565b5f516020614b625f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516134999190614b22565b5f60405180830381855af49150503d805f81146134d1576040519150601f19603f3d011682016040523d82523d5f602084013e6134d6565b606091505b50915091506134e6858383613790565b95945050505050565b3415610c065760405163b398979f60e01b815260040160405180910390fd5b80511561351e5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b611b376131cc565b6135476131cc565b5f516020614b825f395f51905f52805460ff19169055565b6120f76131cc565b61356f6131cc565b610b1982826137ec565b5f5f5f83516041036135b0576020840151604085015160608601515f1a6135a28882858561384b565b9550955095505050506135bb565b505081515f91506002905b9250925092565b5f8260038111156135d5576135d5614b2d565b036135de575050565b60018260038111156135f2576135f2614b2d565b036136105760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561362457613624614b2d565b036136455760405163fce698f760e01b815260048101829052602401610c42565b600382600381111561365957613659614b2d565b03610b19576040516335e2f38360e21b815260048101829052602401610c42565b5f5f516020614b425f395f51905f5281613692612642565b8051909150156136aa57805160209091012092915050565b815480156136b9579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614b425f395f51905f52816136fa612702565b80519091501561371257805160209091012092915050565b600182015480156136b9579392505050565b5f5f60205f8451602086015f885af180613743576040513d5f823e3d81fd5b50505f513d9150811561375a578060011415613767565b6001600160a01b0384163b155b15612dba57604051635274afe760e01b81526001600160a01b0385166004820152602401610c42565b6060826137a5576137a08261350e565b610a5e565b81511580156137bc57506001600160a01b0384163b155b156137e557604051639996b31560e01b81526001600160a01b0385166004820152602401610c42565b5080610a5e565b6137f46131cc565b5f516020614b425f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10261382d84826147d6565b506003810161383c83826147d6565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561388457505f915060039050826120b2565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156138d5573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661390057505f9250600191508290506120b2565b975f9750879650945050505050565b50805461391b90614590565b5f825580601f1061392a575050565b601f0160209004905f5260205f2090810190610ae091906139b5565b5080545f8255905f5260205f2090810190610ae091906139c9565b828054828255905f5260205f209081019282156139a5579160200282015b828111156139a5578251829061399590826147d6565b509160200191906001019061397f565b506139b19291506139c9565b5090565b5b808211156139b1575f81556001016139b6565b808211156139b1575f6139dc828261390f565b506001016139c9565b5f5f83601f8401126139f5575f5ffd5b5081356001600160401b03811115613a0b575f5ffd5b6020830191508360208260051b8501011115613a25575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f191681016001600160401b0381118282101715613a6557613a65613a2c565b6040525050565b5f6001600160401b03821115613a8457613a84613a2c565b5060051b60200190565b5f6001600160401b03821115613aa657613aa6613a2c565b50601f01601f191660200190565b5f613abe83613a8e565b604051613acb8282613a40565b809250848152858585011115613adf575f5ffd5b848460208301375f6020868301015250509392505050565b5f82601f830112613b06575f5ffd5b610a5e83833560208501613ab4565b5f613b1f83613a6c565b604051613b2c8282613a40565b84815291505060208101600584901b830185811115613b49575f5ffd5b835b81811015611ca85780356001600160401b03811115613b68575f5ffd5b613b7488828801613af7565b84525060209283019201613b4b565b5f82601f830112613b92575f5ffd5b610a5e83833560208501613b15565b5f5f5f60408486031215613bb3575f5ffd5b83356001600160401b03811115613bc8575f5ffd5b613bd4868287016139e5565b90945092505060208401356001600160401b03811115613bf2575f5ffd5b8401601f81018613613c02575f5ffd5b8035613c0d81613a6c565b604051613c1a8282613a40565b80915082815260208101915060208360051b850101925088831115613c3d575f5ffd5b602084015b83811015613c7d5780356001600160401b03811115613c5f575f5ffd5b613c6e8b602083890101613b83565b84525060209283019201613c42565b50809450505050509250925092565b6001600160a01b0381168114610ae0575f5ffd5b803561157481613c8c565b5f60208284031215613cbb575f5ffd5b8135610a5e81613c8c565b5f60208284031215613cd6575f5ffd5b81356001600160401b03811115613ceb575f5ffd5b8201601f81018413613cfb575f5ffd5b8035613d0681613a6c565b604051613d138282613a40565b80915082815260208101915060208360051b850101925086831115613d36575f5ffd5b6020840193505b82841015613d61578335613d5081613c8c565b825260209384019390910190613d3d565b9695505050505050565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610bb4565b5f5f5f5f5f5f60a08789031215613daa575f5ffd5b8635613db581613c8c565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613de4575f5ffd5b613df089828a016139e5565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613e18575f5ffd5b8735613e2381613c8c565b96506020880135613e3381613c8c565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115613e62575f5ffd5b613e6e8a828b016139e5565b989b979a50959850939692959293505050565b5f5f60408385031215613e92575f5ffd5b8235613e9d81613c8c565b915060208301356001600160401b03811115613eb7575f5ffd5b613ec385828601613af7565b9150509250929050565b5f60208284031215613edd575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f5f5f6101208b8d031215613efe575f5ffd5b8a35613f0981613c8c565b9950613f1760208c01613ca0565b9850613f2560408c01613ca0565b975060608b0135965060808b0135955060a08b0135945060c08b0135935060e08b01356001600160401b03811115613f5b575f5ffd5b613f678d828e01613af7565b9350506101008b01356001600160401b03811115613f83575f5ffd5b613f8f8d828e016139e5565b915080935050809150509295989b9194979a5092959850565b5f5f60408385031215613fb9575f5ffd5b8235613fc481613c8c565b91506020830135613fd481613c8c565b809150509250929050565b5f5f5f5f5f5f5f5f5f6101008a8c031215613ff8575f5ffd5b893561400381613c8c565b985060208a013561401381613c8c565b975060408a0135965060608a0135955060808a0135945060a08a0135935060c08a01356001600160401b03811115614049575f5ffd5b6140558c828d01613af7565b93505060e08a01356001600160401b03811115614070575f5ffd5b61407c8c828d016139e5565b915080935050809150509295985092959850929598565b602080825282518282018190525f918401906040840190835b818110156140d35783516001600160a01b03168352602093840193909201916001016140ac565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156141a55760df198786030183526141908585516140de565b94506020938401939290920191600101614174565b50929695505050505050565b60ff60f81b8816815260e060208201525f6141cf60e08301896140de565b82810360408401526141e181896140de565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614236578351835260209384019390920191600101614218565b50909b9a5050505050505050505050565b5f5f60408385031215614258575f5ffd5b823561426381613c8c565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156140d3576142ba83855180516001600160a01b0316825260208082015190830152604090810151910152565b602093909301926060929092019160010161428a565b602081525f610a5e60208301846140de565b803560ff81168114611574575f5ffd5b5f60208284031215614302575f5ffd5b610a5e826142e2565b5f5f5f6060848603121561431d575f5ffd5b833561432881613c8c565b9250602084013561433881613c8c565b9150604084013561434881613c8c565b809150509250925092565b5f5f5f5f5f5f5f60c0888a031215614369575f5ffd5b87359650602088013561437b81613c8c565b9550604088013561438b81613c8c565b94506060880135935060808801356001600160401b038111156143ac575f5ffd5b6143b88a828b016139e5565b90945092505060a08801356001600160401b038111156143d6575f5ffd5b6143e28a828b01613b83565b91505092959891949750929550565b5f5f5f60608486031215614403575f5ffd5b833561440e81613c8c565b925060208401356001600160401b03811115614428575f5ffd5b8401601f81018613614438575f5ffd5b61444786823560208401613ab4565b925050614456604085016142e2565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e19833603018112614487575f5ffd5b9190910192915050565b5f5f8335601e198436030181126144a6575f5ffd5b8301803591506001600160401b038211156144bf575f5ffd5b6020019150600581901b3603821315613a25575f5ffd5b5f602082840312156144e6575f5ffd5b5051919050565b5f606082840312156144fd575f5ffd5b604051606081018181106001600160401b038211171561451f5761451f613a2c565b8060405250809150825161453281613c8c565b815260208381015190820152604092830151920191909152919050565b5f6060828403121561455f575f5ffd5b610a5e83836144ed565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610bb457610bb4614569565b600181811c908216806145a457607f821691505b6020821081036145c257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f604083850312156145d9575f5ffd5b505080516020909101519092909150565b5f602082840312156145fa575f5ffd5b81516001600160401b0381111561460f575f5ffd5b8201601f8101841361461f575f5ffd5b805161462a81613a6c565b6040516146378282613a40565b8281526020606090930284018301928101915086831115614656575f5ffd5b6020840193505b82841015613d615761466f87856144ed565b825260208201915060608401935061465d565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b8681101561473257838303601f19018852813536879003601e190181126146e6575f5ffd5b86016020810190356001600160401b03811115614701575f5ffd5b80360382131561470f575f5ffd5b61471a858284614682565b60209a8b019a909550939093019250506001016146c1565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f9061477e90830184866146aa565b9998505050505050505050565b601f82111561227257805f5260205f20601f840160051c810160208510156147b05750805b601f840160051c820191505b818110156147cf575f81556001016147bc565b5050505050565b81516001600160401b038111156147ef576147ef613a2c565b614803816147fd8454614590565b8461478b565b6020601f821160018114614835575f831561481e5750848201515b5f19600385901b1c1916600184901b1784556147cf565b5f84815260208120601f198516915b828110156148645787850151825560209485019460019092019101614844565b508482101561488157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610a5e368484613b15565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f610a5e600d83018461489c565b5f602082840312156148e4575f5ffd5b81516001600160401b038111156148f9575f5ffd5b8201601f81018413614909575f5ffd5b805161491481613a8e565b6040516149218282613a40565b828152866020848601011115614935575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b606081525f61496560608301866140de565b828103602084015261497781866140de565b91505060ff83166040830152949350505050565b5f61499f614999838661489c565b8461489c565b949350505050565b8082028115828204841417610bb457610bb4614569565b5f602082840312156149ce575f5ffd5b81518015158114610a5e575f5ffd5b5f60033d1115610e055760045f5f3e505f5160e01c90565b5f60443d1015614a025790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614a2b57505090565b80820180516001600160401b03811115614a46575050505090565b3d8401600319018282016020011115614a60575050505090565b614a6f60208285010185613a40565b509392505050565b5f5f60233d1115614a9057602060045f3e50505f516001905b9091565b5f82614aae57634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160a01b0389811682528816602082015260408101879052606081018690526080810185905260a0810184905260e060c082018190525f90610e4990830184866146aa565b81810381811115610bb457610bb4614569565b634e487b7160e01b5f52603160045260245ffd5b5f610a5e828461489c565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220763da0e696c66db5271b9deeb5d54c654122129a91900c41cb7c6e6894e7a70364736f6c634300081c0033",
}

// BridgeCrossABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeCrossMetaData.ABI instead.
var BridgeCrossABI = BridgeCrossMetaData.ABI

// Deprecated: Use BridgeCrossMetaData.Sigs instead.
// BridgeCrossFuncSigs maps the 4-byte function signature to its string representation.
var BridgeCrossFuncSigs = BridgeCrossMetaData.Sigs

// BridgeCrossBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeCrossMetaData.Bin instead.
var BridgeCrossBin = BridgeCrossMetaData.Bin

// DeployBridgeCross deploys a new Ethereum contract, binding an instance of BridgeCross to it.
func DeployBridgeCross(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeCross, error) {
	parsed, err := BridgeCrossMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeCrossBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeCross{BridgeCrossCaller: BridgeCrossCaller{contract: contract}, BridgeCrossTransactor: BridgeCrossTransactor{contract: contract}, BridgeCrossFilterer: BridgeCrossFilterer{contract: contract}}, nil
}

// BridgeCross is an auto generated Go binding around an Ethereum contract.
type BridgeCross struct {
	BridgeCrossCaller     // Read-only binding to the contract
	BridgeCrossTransactor // Write-only binding to the contract
	BridgeCrossFilterer   // Log filterer for contract events
}

// BridgeCrossCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCrossCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCrossTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeCrossTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCrossFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeCrossFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCrossSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeCrossSession struct {
	Contract     *BridgeCross      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCrossCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCrossCallerSession struct {
	Contract *BridgeCrossCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeCrossTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeCrossTransactorSession struct {
	Contract     *BridgeCrossTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeCrossRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeCrossRaw struct {
	Contract *BridgeCross // Generic contract binding to access the raw methods on
}

// BridgeCrossCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCrossCallerRaw struct {
	Contract *BridgeCrossCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeCrossTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeCrossTransactorRaw struct {
	Contract *BridgeCrossTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeCross creates a new instance of BridgeCross, bound to a specific deployed contract.
func NewBridgeCross(address common.Address, backend bind.ContractBackend) (*BridgeCross, error) {
	contract, err := bindBridgeCross(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeCross{BridgeCrossCaller: BridgeCrossCaller{contract: contract}, BridgeCrossTransactor: BridgeCrossTransactor{contract: contract}, BridgeCrossFilterer: BridgeCrossFilterer{contract: contract}}, nil
}

// NewBridgeCrossCaller creates a new read-only instance of BridgeCross, bound to a specific deployed contract.
func NewBridgeCrossCaller(address common.Address, caller bind.ContractCaller) (*BridgeCrossCaller, error) {
	contract, err := bindBridgeCross(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossCaller{contract: contract}, nil
}

// NewBridgeCrossTransactor creates a new write-only instance of BridgeCross, bound to a specific deployed contract.
func NewBridgeCrossTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeCrossTransactor, error) {
	contract, err := bindBridgeCross(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossTransactor{contract: contract}, nil
}

// NewBridgeCrossFilterer creates a new log filterer instance of BridgeCross, bound to a specific deployed contract.
func NewBridgeCrossFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeCrossFilterer, error) {
	contract, err := bindBridgeCross(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossFilterer{contract: contract}, nil
}

// bindBridgeCross binds a generic wrapper to an already deployed contract.
func bindBridgeCross(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeCrossMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCross *BridgeCrossRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCross.Contract.BridgeCrossCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCross *BridgeCrossRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeCrossTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCross *BridgeCrossRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeCrossTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCross *BridgeCrossCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCross.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCross *BridgeCrossTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCross *BridgeCrossTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCross.Contract.contract.Transact(opts, method, params...)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeCross *BridgeCrossCaller) BridgeFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "BridgeFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeCross *BridgeCrossSession) BridgeFeeManager() (common.Address, error) {
	return _BridgeCross.Contract.BridgeFeeManager(&_BridgeCross.CallOpts)
}

// BridgeFeeManager is a free data retrieval call binding the contract method 0xe3a446eb.
//
// Solidity: function BridgeFeeManager() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) BridgeFeeManager() (common.Address, error) {
	return _BridgeCross.Contract.BridgeFeeManager(&_BridgeCross.CallOpts)
}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) BridgeFeeManagerLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "BridgeFeeManagerLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) BridgeFeeManagerLength() (*big.Int, error) {
	return _BridgeCross.Contract.BridgeFeeManagerLength(&_BridgeCross.CallOpts)
}

// BridgeFeeManagerLength is a free data retrieval call binding the contract method 0x114ee544.
//
// Solidity: function BridgeFeeManagerLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) BridgeFeeManagerLength() (*big.Int, error) {
	return _BridgeCross.Contract.BridgeFeeManagerLength(&_BridgeCross.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeCross *BridgeCrossCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeCross *BridgeCrossSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeCross.Contract.UPGRADEINTERFACEVERSION(&_BridgeCross.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeCross *BridgeCrossCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeCross.Contract.UPGRADEINTERFACEVERSION(&_BridgeCross.CallOpts)
}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeCross *BridgeCrossCaller) AllFeeInfo(opts *bind.CallOpts) ([]IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "allFeeInfo")

	if err != nil {
		return *new([]IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeFeeManagerFeeInfo)).(*[]IBridgeFeeManagerFeeInfo)

	return out0, err

}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeCross *BridgeCrossSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.AllFeeInfo(&_BridgeCross.CallOpts)
}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeCross *BridgeCrossCallerSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.AllFeeInfo(&_BridgeCross.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeCross *BridgeCrossCaller) AllPairs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "allPairs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeCross *BridgeCrossSession) AllPairs() ([]common.Address, error) {
	return _BridgeCross.Contract.AllPairs(&_BridgeCross.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0xc97682f8.
//
// Solidity: function allPairs() view returns(address[])
func (_BridgeCross *BridgeCrossCallerSession) AllPairs() ([]common.Address, error) {
	return _BridgeCross.Contract.AllPairs(&_BridgeCross.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeCross *BridgeCrossCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeCross *BridgeCrossSession) AllTokens() ([]common.Address, error) {
	return _BridgeCross.Contract.AllTokens(&_BridgeCross.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeCross *BridgeCrossCallerSession) AllTokens() ([]common.Address, error) {
	return _BridgeCross.Contract.AllTokens(&_BridgeCross.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeCross *BridgeCrossCaller) CalculateFee(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "calculateFee", token, value)

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
func (_BridgeCross *BridgeCrossSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeCross.Contract.CalculateFee(&_BridgeCross.CallOpts, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeCross *BridgeCrossCallerSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeCross.Contract.CalculateFee(&_BridgeCross.CallOpts, token, value)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeCross *BridgeCrossCaller) Contains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "contains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeCross *BridgeCrossSession) Contains(token common.Address) (bool, error) {
	return _BridgeCross.Contract.Contains(&_BridgeCross.CallOpts, token)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeCross *BridgeCrossCallerSession) Contains(token common.Address) (bool, error) {
	return _BridgeCross.Contract.Contains(&_BridgeCross.CallOpts, token)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) Denominator() (*big.Int, error) {
	return _BridgeCross.Contract.Denominator(&_BridgeCross.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) Denominator() (*big.Int, error) {
	return _BridgeCross.Contract.Denominator(&_BridgeCross.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeCross *BridgeCrossCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeCross *BridgeCrossSession) DomainSeparator() ([32]byte, error) {
	return _BridgeCross.Contract.DomainSeparator(&_BridgeCross.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeCross *BridgeCrossCallerSession) DomainSeparator() ([32]byte, error) {
	return _BridgeCross.Contract.DomainSeparator(&_BridgeCross.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeCross *BridgeCrossCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "eip712Domain")

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
func (_BridgeCross *BridgeCrossSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeCross.Contract.Eip712Domain(&_BridgeCross.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BridgeCross *BridgeCrossCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BridgeCross.Contract.Eip712Domain(&_BridgeCross.CallOpts)
}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossCaller) FeeInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "feeInfoByIndex", index)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.FeeInfoByIndex(&_BridgeCross.CallOpts, index)
}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossCallerSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.FeeInfoByIndex(&_BridgeCross.CallOpts, index)
}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeCross *BridgeCrossCaller) GetPairToken(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "getPairToken", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeCross *BridgeCrossSession) GetPairToken(token common.Address) (common.Address, error) {
	return _BridgeCross.Contract.GetPairToken(&_BridgeCross.CallOpts, token)
}

// GetPairToken is a free data retrieval call binding the contract method 0x71c59d7b.
//
// Solidity: function getPairToken(address token) view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) GetPairToken(token common.Address) (common.Address, error) {
	return _BridgeCross.Contract.GetPairToken(&_BridgeCross.CallOpts, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossCaller) GetTokenFee(opts *bind.CallOpts, token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "getTokenFee", token)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.GetTokenFee(&_BridgeCross.CallOpts, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeCross *BridgeCrossCallerSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeCross.Contract.GetTokenFee(&_BridgeCross.CallOpts, token)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) InitializedAt() (*big.Int, error) {
	return _BridgeCross.Contract.InitializedAt(&_BridgeCross.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) InitializedAt() (*big.Int, error) {
	return _BridgeCross.Contract.InitializedAt(&_BridgeCross.CallOpts)
}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeCross *BridgeCrossCaller) IsValidToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "isValidToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeCross *BridgeCrossSession) IsValidToken(token common.Address) (bool, error) {
	return _BridgeCross.Contract.IsValidToken(&_BridgeCross.CallOpts, token)
}

// IsValidToken is a free data retrieval call binding the contract method 0xc1876453.
//
// Solidity: function isValidToken(address token) view returns(bool)
func (_BridgeCross *BridgeCrossCallerSession) IsValidToken(token common.Address) (bool, error) {
	return _BridgeCross.Contract.IsValidToken(&_BridgeCross.CallOpts, token)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeCross *BridgeCrossCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeCross *BridgeCrossSession) IsValidator(validator common.Address) (bool, error) {
	return _BridgeCross.Contract.IsValidator(&_BridgeCross.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BridgeCross *BridgeCrossCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _BridgeCross.Contract.IsValidator(&_BridgeCross.CallOpts, validator)
}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) NextFinalizeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "nextFinalizeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) NextFinalizeIndex() (*big.Int, error) {
	return _BridgeCross.Contract.NextFinalizeIndex(&_BridgeCross.CallOpts)
}

// NextFinalizeIndex is a free data retrieval call binding the contract method 0x27938c81.
//
// Solidity: function nextFinalizeIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) NextFinalizeIndex() (*big.Int, error) {
	return _BridgeCross.Contract.NextFinalizeIndex(&_BridgeCross.CallOpts)
}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) NextInitiateIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "nextInitiateIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) NextInitiateIndex() (*big.Int, error) {
	return _BridgeCross.Contract.NextInitiateIndex(&_BridgeCross.CallOpts)
}

// NextInitiateIndex is a free data retrieval call binding the contract method 0xe1af7f50.
//
// Solidity: function nextInitiateIndex() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) NextInitiateIndex() (*big.Int, error) {
	return _BridgeCross.Contract.NextInitiateIndex(&_BridgeCross.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCross *BridgeCrossCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCross *BridgeCrossSession) Owner() (common.Address, error) {
	return _BridgeCross.Contract.Owner(&_BridgeCross.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) Owner() (common.Address, error) {
	return _BridgeCross.Contract.Owner(&_BridgeCross.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeCross *BridgeCrossCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeCross *BridgeCrossSession) Paused() (bool, error) {
	return _BridgeCross.Contract.Paused(&_BridgeCross.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeCross *BridgeCrossCallerSession) Paused() (bool, error) {
	return _BridgeCross.Contract.Paused(&_BridgeCross.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeCross *BridgeCrossCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeCross *BridgeCrossSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeCross.Contract.ProxiableUUID(&_BridgeCross.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeCross *BridgeCrossCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeCross.Contract.ProxiableUUID(&_BridgeCross.CallOpts)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeCross *BridgeCrossCaller) RevertedArguments(opts *bind.CallOpts, index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "revertedArguments", index)

	if err != nil {
		return *new(IBridgeStandardFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeStandardFinalizeArguments)).(*IBridgeStandardFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeCross *BridgeCrossSession) RevertedArguments(index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	return _BridgeCross.Contract.RevertedArguments(&_BridgeCross.CallOpts, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x7021fd0e.
//
// Solidity: function revertedArguments(uint256 index) view returns((uint256,address,address,uint256,bytes[]))
func (_BridgeCross *BridgeCrossCallerSession) RevertedArguments(index *big.Int) (IBridgeStandardFinalizeArguments, error) {
	return _BridgeCross.Contract.RevertedArguments(&_BridgeCross.CallOpts, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeCross *BridgeCrossCaller) RevertedReason(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "revertedReason", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeCross *BridgeCrossSession) RevertedReason(index *big.Int) ([]byte, error) {
	return _BridgeCross.Contract.RevertedReason(&_BridgeCross.CallOpts, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0xfe2b8da6.
//
// Solidity: function revertedReason(uint256 index) view returns(bytes)
func (_BridgeCross *BridgeCrossCallerSession) RevertedReason(index *big.Int) ([]byte, error) {
	return _BridgeCross.Contract.RevertedReason(&_BridgeCross.CallOpts, index)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeCross *BridgeCrossCaller) RewardWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "rewardWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeCross *BridgeCrossSession) RewardWallet() (common.Address, error) {
	return _BridgeCross.Contract.RewardWallet(&_BridgeCross.CallOpts)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) RewardWallet() (common.Address, error) {
	return _BridgeCross.Contract.RewardWallet(&_BridgeCross.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeCross *BridgeCrossCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeCross *BridgeCrossSession) Threshold() (uint8, error) {
	return _BridgeCross.Contract.Threshold(&_BridgeCross.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BridgeCross *BridgeCrossCallerSession) Threshold() (uint8, error) {
	return _BridgeCross.Contract.Threshold(&_BridgeCross.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeCross *BridgeCrossCaller) TokenByIndex(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "tokenByIndex", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeCross *BridgeCrossSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeCross.Contract.TokenByIndex(&_BridgeCross.CallOpts, i)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeCross.Contract.TokenByIndex(&_BridgeCross.CallOpts, i)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) TokensLength() (*big.Int, error) {
	return _BridgeCross.Contract.TokensLength(&_BridgeCross.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) TokensLength() (*big.Int, error) {
	return _BridgeCross.Contract.TokensLength(&_BridgeCross.CallOpts)
}

// Xcross is a free data retrieval call binding the contract method 0x211c8ac0.
//
// Solidity: function xcross() view returns(address)
func (_BridgeCross *BridgeCrossCaller) Xcross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "xcross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xcross is a free data retrieval call binding the contract method 0x211c8ac0.
//
// Solidity: function xcross() view returns(address)
func (_BridgeCross *BridgeCrossSession) Xcross() (common.Address, error) {
	return _BridgeCross.Contract.Xcross(&_BridgeCross.CallOpts)
}

// Xcross is a free data retrieval call binding the contract method 0x211c8ac0.
//
// Solidity: function xcross() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) Xcross() (common.Address, error) {
	return _BridgeCross.Contract.Xcross(&_BridgeCross.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeCross *BridgeCrossTransactor) AddToken(opts *bind.TransactOpts, token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "addToken", token, pair)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeCross *BridgeCrossSession) AddToken(token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.AddToken(&_BridgeCross.TransactOpts, token, pair)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address pair) returns()
func (_BridgeCross *BridgeCrossTransactorSession) AddToken(token common.Address, pair common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.AddToken(&_BridgeCross.TransactOpts, token, pair)
}

// AddTokenDeploy is a paid mutator transaction binding the contract method 0xf7684ec4.
//
// Solidity: function addTokenDeploy(address pair, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BridgeCross *BridgeCrossTransactor) AddTokenDeploy(opts *bind.TransactOpts, pair common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "addTokenDeploy", pair, symbol, decimals)
}

// AddTokenDeploy is a paid mutator transaction binding the contract method 0xf7684ec4.
//
// Solidity: function addTokenDeploy(address pair, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BridgeCross *BridgeCrossSession) AddTokenDeploy(pair common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeCross.Contract.AddTokenDeploy(&_BridgeCross.TransactOpts, pair, symbol, decimals)
}

// AddTokenDeploy is a paid mutator transaction binding the contract method 0xf7684ec4.
//
// Solidity: function addTokenDeploy(address pair, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BridgeCross *BridgeCrossTransactorSession) AddTokenDeploy(pair common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeCross.Contract.AddTokenDeploy(&_BridgeCross.TransactOpts, pair, symbol, decimals)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) Bridge(opts *bind.TransactOpts, token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "bridge", token, value, gas, service, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) Bridge(token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Bridge(&_BridgeCross.TransactOpts, token, value, gas, service, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) Bridge(token common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Bridge(&_BridgeCross.TransactOpts, token, value, gas, service, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) BridgeTo(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "bridgeTo", token, to, value, gas, service, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeTo(&_BridgeCross.TransactOpts, token, to, value, gas, service, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 service, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeTo(&_BridgeCross.TransactOpts, token, to, value, gas, service, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeCross *BridgeCrossTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeCross *BridgeCrossSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BridgeCross.Contract.ChangeThreshold(&_BridgeCross.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BridgeCross *BridgeCrossTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BridgeCross.Contract.ChangeThreshold(&_BridgeCross.TransactOpts, threshold_)
}

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) Finalize(opts *bind.TransactOpts, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "finalize", index, token, to, value, extraData, sigs)
}

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Finalize(&_BridgeCross.TransactOpts, index, token, to, value, extraData, sigs)
}

// Finalize is a paid mutator transaction binding the contract method 0xf120c400.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, bytes[] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, sigs [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Finalize(&_BridgeCross.TransactOpts, index, token, to, value, extraData, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) FinalizeBatch(opts *bind.TransactOpts, args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "finalizeBatch", args, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.FinalizeBatch(&_BridgeCross.TransactOpts, args, sigs)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x008bd028.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, bytes[][] sigs) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, sigs [][][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.FinalizeBatch(&_BridgeCross.TransactOpts, args, sigs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeCross *BridgeCrossTransactor) Initialize(opts *bind.TransactOpts, crossMintableERC20Code common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "initialize", crossMintableERC20Code, rewardWallet_, BridgeFeeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeCross *BridgeCrossSession) Initialize(crossMintableERC20Code common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, rewardWallet_, BridgeFeeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address BridgeFeeManager) returns()
func (_BridgeCross *BridgeCrossTransactorSession) Initialize(crossMintableERC20Code common.Address, rewardWallet_ common.Address, BridgeFeeManager common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, rewardWallet_, BridgeFeeManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeCross *BridgeCrossTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeCross *BridgeCrossSession) Pause() (*types.Transaction, error) {
	return _BridgeCross.Contract.Pause(&_BridgeCross.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeCross *BridgeCrossTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeCross.Contract.Pause(&_BridgeCross.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactor) PauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "pauseToken", token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeCross *BridgeCrossSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.PauseToken(&_BridgeCross.TransactOpts, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactorSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.PauseToken(&_BridgeCross.TransactOpts, token)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridge", token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x69a3318b.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, service, deadline, permitSig, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x50d6fb48.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, uint256 deadline, bytes permitSig, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, deadline *big.Int, permitSig []byte, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, service, deadline, permitSig, extraData)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeCross *BridgeCrossSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveToken(&_BridgeCross.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveToken(&_BridgeCross.TransactOpts, token)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeCross *BridgeCrossTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeCross *BridgeCrossSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveValidator(&_BridgeCross.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BridgeCross *BridgeCrossTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveValidator(&_BridgeCross.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveValidators(&_BridgeCross.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveValidators(&_BridgeCross.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCross *BridgeCrossTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCross *BridgeCrossSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeCross.Contract.RenounceOwnership(&_BridgeCross.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeCross *BridgeCrossTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeCross.Contract.RenounceOwnership(&_BridgeCross.TransactOpts)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeCross *BridgeCrossTransactor) RetryFinalize(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "retryFinalize", index)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeCross *BridgeCrossSession) RetryFinalize(index *big.Int) (*types.Transaction, error) {
	return _BridgeCross.Contract.RetryFinalize(&_BridgeCross.TransactOpts, index)
}

// RetryFinalize is a paid mutator transaction binding the contract method 0xa9823183.
//
// Solidity: function retryFinalize(uint256 index) returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) RetryFinalize(index *big.Int) (*types.Transaction, error) {
	return _BridgeCross.Contract.RetryFinalize(&_BridgeCross.TransactOpts, index)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeCross *BridgeCrossTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeCross *BridgeCrossSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetValidator(&_BridgeCross.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BridgeCross *BridgeCrossTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetValidator(&_BridgeCross.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetValidators(&_BridgeCross.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetValidators(&_BridgeCross.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCross *BridgeCrossTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCross *BridgeCrossSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.TransferOwnership(&_BridgeCross.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeCross *BridgeCrossTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.TransferOwnership(&_BridgeCross.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeCross *BridgeCrossTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeCross *BridgeCrossSession) Unpause() (*types.Transaction, error) {
	return _BridgeCross.Contract.Unpause(&_BridgeCross.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeCross *BridgeCrossTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeCross.Contract.Unpause(&_BridgeCross.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactor) UnpauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "unpauseToken", token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeCross *BridgeCrossSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.UnpauseToken(&_BridgeCross.TransactOpts, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_BridgeCross *BridgeCrossTransactorSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.UnpauseToken(&_BridgeCross.TransactOpts, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeCross *BridgeCrossTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeCross *BridgeCrossSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.UpgradeToAndCall(&_BridgeCross.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeCross *BridgeCrossTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.UpgradeToAndCall(&_BridgeCross.TransactOpts, newImplementation, data)
}

// BridgeCrossBridgeFinalizeRevertedIterator is returned from FilterBridgeFinalizeReverted and is used to iterate over the raw logs and unpacked data for BridgeFinalizeReverted events raised by the BridgeCross contract.
type BridgeCrossBridgeFinalizeRevertedIterator struct {
	Event *BridgeCrossBridgeFinalizeReverted // Event containing the contract specifics and raw log

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
func (it *BridgeCrossBridgeFinalizeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossBridgeFinalizeReverted)
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
		it.Event = new(BridgeCrossBridgeFinalizeReverted)
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
func (it *BridgeCrossBridgeFinalizeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossBridgeFinalizeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossBridgeFinalizeReverted represents a BridgeFinalizeReverted event raised by the BridgeCross contract.
type BridgeCrossBridgeFinalizeReverted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeReverted is a free log retrieval operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BridgeCross *BridgeCrossFilterer) FilterBridgeFinalizeReverted(opts *bind.FilterOpts, index []*big.Int) (*BridgeCrossBridgeFinalizeRevertedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossBridgeFinalizeRevertedIterator{contract: _BridgeCross.contract, event: "BridgeFinalizeReverted", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeReverted is a free log subscription operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BridgeCross *BridgeCrossFilterer) WatchBridgeFinalizeReverted(opts *bind.WatchOpts, sink chan<- *BridgeCrossBridgeFinalizeReverted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossBridgeFinalizeReverted)
				if err := _BridgeCross.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseBridgeFinalizeReverted(log types.Log) (*BridgeCrossBridgeFinalizeReverted, error) {
	event := new(BridgeCrossBridgeFinalizeReverted)
	if err := _BridgeCross.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BridgeCross contract.
type BridgeCrossBridgeFinalizedIterator struct {
	Event *BridgeCrossBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeCrossBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossBridgeFinalized)
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
		it.Event = new(BridgeCrossBridgeFinalized)
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
func (it *BridgeCrossBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossBridgeFinalized represents a BridgeFinalized event raised by the BridgeCross contract.
type BridgeCrossBridgeFinalized struct {
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
func (_BridgeCross *BridgeCrossFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, index []*big.Int, token []common.Address, to []common.Address) (*BridgeCrossBridgeFinalizedIterator, error) {

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

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossBridgeFinalizedIterator{contract: _BridgeCross.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_BridgeCross *BridgeCrossFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BridgeCrossBridgeFinalized, index []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossBridgeFinalized)
				if err := _BridgeCross.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseBridgeFinalized(log types.Log) (*BridgeCrossBridgeFinalized, error) {
	event := new(BridgeCrossBridgeFinalized)
	if err := _BridgeCross.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BridgeCross contract.
type BridgeCrossBridgeInitiatedIterator struct {
	Event *BridgeCrossBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BridgeCrossBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossBridgeInitiated)
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
		it.Event = new(BridgeCrossBridgeInitiated)
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
func (it *BridgeCrossBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossBridgeInitiated represents a BridgeInitiated event raised by the BridgeCross contract.
type BridgeCrossBridgeInitiated struct {
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
func (_BridgeCross *BridgeCrossFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, index []*big.Int, token []common.Address, from []common.Address) (*BridgeCrossBridgeInitiatedIterator, error) {

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

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "BridgeInitiated", indexRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossBridgeInitiatedIterator{contract: _BridgeCross.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0xd0155ebd9e7010a412821ba42bff3fb5901d17f75508f0bc19e468ce743c235c.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 gasFee, uint256 serviceFee, uint256 time, bytes[] extraData)
func (_BridgeCross *BridgeCrossFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BridgeCrossBridgeInitiated, index []*big.Int, token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "BridgeInitiated", indexRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossBridgeInitiated)
				if err := _BridgeCross.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseBridgeInitiated(log types.Log) (*BridgeCrossBridgeInitiated, error) {
	event := new(BridgeCrossBridgeInitiated)
	if err := _BridgeCross.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BridgeCross contract.
type BridgeCrossEIP712DomainChangedIterator struct {
	Event *BridgeCrossEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BridgeCrossEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossEIP712DomainChanged)
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
		it.Event = new(BridgeCrossEIP712DomainChanged)
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
func (it *BridgeCrossEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossEIP712DomainChanged represents a EIP712DomainChanged event raised by the BridgeCross contract.
type BridgeCrossEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeCross *BridgeCrossFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BridgeCrossEIP712DomainChangedIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossEIP712DomainChangedIterator{contract: _BridgeCross.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BridgeCross *BridgeCrossFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BridgeCrossEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossEIP712DomainChanged)
				if err := _BridgeCross.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseEIP712DomainChanged(log types.Log) (*BridgeCrossEIP712DomainChanged, error) {
	event := new(BridgeCrossEIP712DomainChanged)
	if err := _BridgeCross.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeCross contract.
type BridgeCrossInitializedIterator struct {
	Event *BridgeCrossInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeCrossInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossInitialized)
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
		it.Event = new(BridgeCrossInitialized)
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
func (it *BridgeCrossInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossInitialized represents a Initialized event raised by the BridgeCross contract.
type BridgeCrossInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeCross *BridgeCrossFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeCrossInitializedIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossInitializedIterator{contract: _BridgeCross.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeCross *BridgeCrossFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeCrossInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossInitialized)
				if err := _BridgeCross.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseInitialized(log types.Log) (*BridgeCrossInitialized, error) {
	event := new(BridgeCrossInitialized)
	if err := _BridgeCross.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeCross contract.
type BridgeCrossOwnershipTransferredIterator struct {
	Event *BridgeCrossOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeCrossOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossOwnershipTransferred)
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
		it.Event = new(BridgeCrossOwnershipTransferred)
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
func (it *BridgeCrossOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeCross contract.
type BridgeCrossOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeCross *BridgeCrossFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeCrossOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossOwnershipTransferredIterator{contract: _BridgeCross.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeCross *BridgeCrossFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeCrossOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossOwnershipTransferred)
				if err := _BridgeCross.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeCrossOwnershipTransferred, error) {
	event := new(BridgeCrossOwnershipTransferred)
	if err := _BridgeCross.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossPairUpdatedIterator is returned from FilterPairUpdated and is used to iterate over the raw logs and unpacked data for PairUpdated events raised by the BridgeCross contract.
type BridgeCrossPairUpdatedIterator struct {
	Event *BridgeCrossPairUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeCrossPairUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossPairUpdated)
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
		it.Event = new(BridgeCrossPairUpdated)
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
func (it *BridgeCrossPairUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossPairUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossPairUpdated represents a PairUpdated event raised by the BridgeCross contract.
type BridgeCrossPairUpdated struct {
	Token      common.Address
	Pair       common.Address
	Registered bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPairUpdated is a free log retrieval operation binding the contract event 0xa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce.
//
// Solidity: event PairUpdated(address indexed token, address indexed pair, bool registered)
func (_BridgeCross *BridgeCrossFilterer) FilterPairUpdated(opts *bind.FilterOpts, token []common.Address, pair []common.Address) (*BridgeCrossPairUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "PairUpdated", tokenRule, pairRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossPairUpdatedIterator{contract: _BridgeCross.contract, event: "PairUpdated", logs: logs, sub: sub}, nil
}

// WatchPairUpdated is a free log subscription operation binding the contract event 0xa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce.
//
// Solidity: event PairUpdated(address indexed token, address indexed pair, bool registered)
func (_BridgeCross *BridgeCrossFilterer) WatchPairUpdated(opts *bind.WatchOpts, sink chan<- *BridgeCrossPairUpdated, token []common.Address, pair []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "PairUpdated", tokenRule, pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossPairUpdated)
				if err := _BridgeCross.contract.UnpackLog(event, "PairUpdated", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParsePairUpdated(log types.Log) (*BridgeCrossPairUpdated, error) {
	event := new(BridgeCrossPairUpdated)
	if err := _BridgeCross.contract.UnpackLog(event, "PairUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeCross contract.
type BridgeCrossPausedIterator struct {
	Event *BridgeCrossPaused // Event containing the contract specifics and raw log

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
func (it *BridgeCrossPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossPaused)
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
		it.Event = new(BridgeCrossPaused)
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
func (it *BridgeCrossPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossPaused represents a Paused event raised by the BridgeCross contract.
type BridgeCrossPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeCross *BridgeCrossFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeCrossPausedIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossPausedIterator{contract: _BridgeCross.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeCross *BridgeCrossFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeCrossPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossPaused)
				if err := _BridgeCross.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParsePaused(log types.Log) (*BridgeCrossPaused, error) {
	event := new(BridgeCrossPaused)
	if err := _BridgeCross.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BridgeCross contract.
type BridgeCrossThresholdChangedIterator struct {
	Event *BridgeCrossThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeCrossThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossThresholdChanged)
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
		it.Event = new(BridgeCrossThresholdChanged)
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
func (it *BridgeCrossThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossThresholdChanged represents a ThresholdChanged event raised by the BridgeCross contract.
type BridgeCrossThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BridgeCross *BridgeCrossFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BridgeCrossThresholdChangedIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossThresholdChangedIterator{contract: _BridgeCross.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BridgeCross *BridgeCrossFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeCrossThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossThresholdChanged)
				if err := _BridgeCross.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseThresholdChanged(log types.Log) (*BridgeCrossThresholdChanged, error) {
	event := new(BridgeCrossThresholdChanged)
	if err := _BridgeCross.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the BridgeCross contract.
type BridgeCrossTokenAddedIterator struct {
	Event *BridgeCrossTokenAdded // Event containing the contract specifics and raw log

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
func (it *BridgeCrossTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossTokenAdded)
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
		it.Event = new(BridgeCrossTokenAdded)
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
func (it *BridgeCrossTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossTokenAdded represents a TokenAdded event raised by the BridgeCross contract.
type BridgeCrossTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*BridgeCrossTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossTokenAddedIterator{contract: _BridgeCross.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeCrossTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossTokenAdded)
				if err := _BridgeCross.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseTokenAdded(log types.Log) (*BridgeCrossTokenAdded, error) {
	event := new(BridgeCrossTokenAdded)
	if err := _BridgeCross.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossTokenPausedIterator is returned from FilterTokenPaused and is used to iterate over the raw logs and unpacked data for TokenPaused events raised by the BridgeCross contract.
type BridgeCrossTokenPausedIterator struct {
	Event *BridgeCrossTokenPaused // Event containing the contract specifics and raw log

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
func (it *BridgeCrossTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossTokenPaused)
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
		it.Event = new(BridgeCrossTokenPaused)
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
func (it *BridgeCrossTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossTokenPaused represents a TokenPaused event raised by the BridgeCross contract.
type BridgeCrossTokenPaused struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenPaused is a free log retrieval operation binding the contract event 0xf017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d.
//
// Solidity: event TokenPaused(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) FilterTokenPaused(opts *bind.FilterOpts, token []common.Address) (*BridgeCrossTokenPausedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "TokenPaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossTokenPausedIterator{contract: _BridgeCross.contract, event: "TokenPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPaused is a free log subscription operation binding the contract event 0xf017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d.
//
// Solidity: event TokenPaused(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) WatchTokenPaused(opts *bind.WatchOpts, sink chan<- *BridgeCrossTokenPaused, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "TokenPaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossTokenPaused)
				if err := _BridgeCross.contract.UnpackLog(event, "TokenPaused", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseTokenPaused(log types.Log) (*BridgeCrossTokenPaused, error) {
	event := new(BridgeCrossTokenPaused)
	if err := _BridgeCross.contract.UnpackLog(event, "TokenPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the BridgeCross contract.
type BridgeCrossTokenRemovedIterator struct {
	Event *BridgeCrossTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeCrossTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossTokenRemoved)
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
		it.Event = new(BridgeCrossTokenRemoved)
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
func (it *BridgeCrossTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossTokenRemoved represents a TokenRemoved event raised by the BridgeCross contract.
type BridgeCrossTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeCrossTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossTokenRemovedIterator{contract: _BridgeCross.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeCrossTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossTokenRemoved)
				if err := _BridgeCross.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseTokenRemoved(log types.Log) (*BridgeCrossTokenRemoved, error) {
	event := new(BridgeCrossTokenRemoved)
	if err := _BridgeCross.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossTokenUnpausedIterator is returned from FilterTokenUnpaused and is used to iterate over the raw logs and unpacked data for TokenUnpaused events raised by the BridgeCross contract.
type BridgeCrossTokenUnpausedIterator struct {
	Event *BridgeCrossTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeCrossTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossTokenUnpaused)
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
		it.Event = new(BridgeCrossTokenUnpaused)
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
func (it *BridgeCrossTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossTokenUnpaused represents a TokenUnpaused event raised by the BridgeCross contract.
type BridgeCrossTokenUnpaused struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnpaused is a free log retrieval operation binding the contract event 0xf38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf.
//
// Solidity: event TokenUnpaused(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) FilterTokenUnpaused(opts *bind.FilterOpts, token []common.Address) (*BridgeCrossTokenUnpausedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "TokenUnpaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossTokenUnpausedIterator{contract: _BridgeCross.contract, event: "TokenUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenUnpaused is a free log subscription operation binding the contract event 0xf38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf.
//
// Solidity: event TokenUnpaused(address indexed token)
func (_BridgeCross *BridgeCrossFilterer) WatchTokenUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeCrossTokenUnpaused, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "TokenUnpaused", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossTokenUnpaused)
				if err := _BridgeCross.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseTokenUnpaused(log types.Log) (*BridgeCrossTokenUnpaused, error) {
	event := new(BridgeCrossTokenUnpaused)
	if err := _BridgeCross.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeCross contract.
type BridgeCrossUnpausedIterator struct {
	Event *BridgeCrossUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeCrossUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossUnpaused)
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
		it.Event = new(BridgeCrossUnpaused)
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
func (it *BridgeCrossUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossUnpaused represents a Unpaused event raised by the BridgeCross contract.
type BridgeCrossUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeCross *BridgeCrossFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeCrossUnpausedIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossUnpausedIterator{contract: _BridgeCross.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeCross *BridgeCrossFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeCrossUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossUnpaused)
				if err := _BridgeCross.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseUnpaused(log types.Log) (*BridgeCrossUnpaused, error) {
	event := new(BridgeCrossUnpaused)
	if err := _BridgeCross.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeCross contract.
type BridgeCrossUpgradedIterator struct {
	Event *BridgeCrossUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeCrossUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossUpgraded)
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
		it.Event = new(BridgeCrossUpgraded)
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
func (it *BridgeCrossUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossUpgraded represents a Upgraded event raised by the BridgeCross contract.
type BridgeCrossUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeCross *BridgeCrossFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeCrossUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossUpgradedIterator{contract: _BridgeCross.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeCross *BridgeCrossFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeCrossUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossUpgraded)
				if err := _BridgeCross.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseUpgraded(log types.Log) (*BridgeCrossUpgraded, error) {
	event := new(BridgeCrossUpgraded)
	if err := _BridgeCross.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the BridgeCross contract.
type BridgeCrossValidatorSetIterator struct {
	Event *BridgeCrossValidatorSet // Event containing the contract specifics and raw log

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
func (it *BridgeCrossValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossValidatorSet)
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
		it.Event = new(BridgeCrossValidatorSet)
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
func (it *BridgeCrossValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossValidatorSet represents a ValidatorSet event raised by the BridgeCross contract.
type BridgeCrossValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BridgeCross *BridgeCrossFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*BridgeCrossValidatorSetIterator, error) {

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossValidatorSetIterator{contract: _BridgeCross.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BridgeCross *BridgeCrossFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *BridgeCrossValidatorSet) (event.Subscription, error) {

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossValidatorSet)
				if err := _BridgeCross.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseValidatorSet(log types.Log) (*BridgeCrossValidatorSet, error) {
	event := new(BridgeCrossValidatorSet)
	if err := _BridgeCross.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
