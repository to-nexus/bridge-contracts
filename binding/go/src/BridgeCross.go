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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BridgeFeeManager\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BridgeFeeManagerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addTokenDeploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFeeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"feeInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossMintableERC20Code\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xcross\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInsufficientValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e3a446eb": "BridgeFeeManager()",
		"114ee544": "BridgeFeeManagerLength()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"f7684ec4": "addTokenDeploy(address,string,uint8)",
		"9d0dc76e": "allFeeInfo()",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
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
		"211c8ac0": "xcross()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614d636100f95f395f81816122060152818161222f015261237e0152614d635ff3fe608060405260043610610343575f3560e01c80637c41ad2c116101bd578063c1876453116100f2578063f2fde38b11610092578063f7684ec41161006d578063f7684ec41461096b578063facd743b1461098a578063fb75b2c7146109a9578063fe2b8da6146109c6575f5ffd5b8063f2fde38b14610924578063f30589c314610943578063f698da2514610957575f5ffd5b8063d92fc67b116100cd578063d92fc67b146108ca578063e1af7f50146108de578063e3a446eb146108f2578063f120c40014610911575f5ffd5b8063c187645314610878578063c97682f814610897578063cbae5958146108ab575f5ffd5b806396ce07951161015d578063ad3cb1cc11610138578063ad3cb1cc146107e9578063aed1d40314610826578063b7f3358d1461083a578063c0c53b8b14610859575f5ffd5b806396ce0795146107955780639d0dc76e146107a9578063a9823183146107ca575f5ffd5b80638b28ab1e116101985780638b28ab1e146106f25780638da5cb5b1461072657806391cf6d3e146107625780639300c92614610776575f5ffd5b80637c41ad2c146106985780638456cb59146106b757806384b0196e146106cb575f5ffd5b806342cde4e8116102935780635dbe47e8116102335780637021fd0e1161020e5780637021fd0e14610602578063715018a61461062e57806371c59d7b14610642578063751b4c9c14610679575f5ffd5b80635dbe47e8146105a35780635fa7b584146105c25780636ff97f1d146105e1575f5ffd5b806351c455791161026e57806351c455791461053a57806352d1902d1461054d5780635476bd72146105615780635c975abb14610580575f5ffd5b806342cde4e8146104e75780634f1ef286146105085780634f6ccce71461051b575f5ffd5b8063252154fa116102fe5780633b3bff0f116102d95780633b3bff0f146104825780633f4ba83a146104a15780633f5d3b5d146104b557806340a141ff146104c8575f5ffd5b8063252154fa1461042f57806327938c811461045b57806337d100751461046f575f5ffd5b80628bd0281461035d578063114ee544146103855780631327d3d8146103a7578063174991ab146103c65780631d40f0d8146103d9578063211c8ac0146103f8575f5ffd5b36610359575f341161035757610357613afe565b005b5f5ffd5b61037061036b366004613cf3565b6109e5565b60405190151581526020015b60405180910390f35b348015610390575f5ffd5b50610399610b11565b60405190815260200161037c565b3480156103b2575f5ffd5b506103576103c1366004613df2565b610b81565b6103706103d4366004613e9c565b610b8f565b3480156103e4575f5ffd5b506103576103f3366004613f42565b610c69565b348015610403575f5ffd5b5060fe54610417906001600160a01b031681565b6040516001600160a01b03909116815260200161037c565b34801561043a575f5ffd5b5061044e610449366004613df2565b610ca3565b60405161037c9190613fe7565b348015610466575f5ffd5b50610399610d40565b61037061047d366004614011565b610d50565b34801561048d575f5ffd5b5061035761049c366004613df2565b610d6b565b3480156104ac575f5ffd5b50610357610d7c565b6103706104c336600461407e565b610d8e565b3480156104d3575f5ffd5b506103576104e2366004613df2565b610e67565b3480156104f2575f5ffd5b5060995460405160ff909116815260200161037c565b6103576105163660046140fd565b610e71565b348015610526575f5ffd5b50610417610535366004614149565b610e8c565b610370610548366004614160565b610e97565b348015610558575f5ffd5b50610399610eb6565b34801561056c575f5ffd5b5061035761057b3660046141f2565b610ed2565b34801561058b575f5ffd5b505f516020614cee5f395f51905f525460ff16610370565b3480156105ae575f5ffd5b506103706105bd366004613df2565b610ee4565b3480156105cd575f5ffd5b506103576105dc366004613df2565b610eef565b3480156105ec575f5ffd5b506105f5610f00565b60405161037c9190614229565b34801561060d575f5ffd5b5061062161061c366004614149565b610f61565b60405161037c91906142a2565b348015610639575f5ffd5b506103576110c1565b34801561064d575f5ffd5b5061041761065c366004613df2565b6001600160a01b039081165f908152603360205260409020541690565b348015610684575f5ffd5b5061044e610693366004614149565b6110d2565b3480156106a3575f5ffd5b506103576106b2366004613df2565b61112d565b3480156106c2575f5ffd5b5061035761113e565b3480156106d6575f5ffd5b506106df61114e565b60405161037c9796959493929190614347565b3480156106fd575f5ffd5b5061071161070c3660046143dd565b6111f7565b6040805192835260208301919091520161037c565b348015610731575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610417565b34801561076d575f5ffd5b5060cb54610399565b348015610781575f5ffd5b50610357610790366004613f42565b611277565b3480156107a0575f5ffd5b506103996112ae565b3480156107b4575f5ffd5b506107bd6112f5565b60405161037c9190614407565b3480156107d5575f5ffd5b506103706107e4366004614149565b611363565b3480156107f4575f5ffd5b50610819604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161037c9190614466565b348015610831575f5ffd5b50610399611622565b348015610845575f5ffd5b50610357610854366004614478565b61162d565b348015610864575f5ffd5b50610357610873366004614491565b61167e565b348015610883575f5ffd5b50610370610892366004613df2565b6117b7565b3480156108a2575f5ffd5b506105f56117e7565b3480156108b6575f5ffd5b506104176108c5366004614149565b6118cd565b3480156108d5575f5ffd5b506103996118d9565b3480156108e9575f5ffd5b506103996118e3565b3480156108fd575f5ffd5b5060ca54610417906001600160a01b031681565b61037061091f3660046144d9565b6118f3565b34801561092f575f5ffd5b5061035761093e366004613df2565b611bba565b34801561094e575f5ffd5b506105f5611bf4565b348015610962575f5ffd5b50610399611c00565b348015610976575f5ffd5b50610417610985366004614577565b611c09565b348015610995575f5ffd5b506103706109a4366004613df2565b611d49565b3480156109b4575f5ffd5b5060cc546001600160a01b0316610417565b3480156109d1575f5ffd5b506108196109e0366004614149565b611d55565b5f805b83811015610b0457610afb858583818110610a0557610a056145e5565b9050602002810190610a1791906145f9565b35868684818110610a2a57610a2a6145e5565b9050602002810190610a3c91906145f9565b610a4d906040810190602001613df2565b878785818110610a5f57610a5f6145e5565b9050602002810190610a7191906145f9565b610a82906060810190604001613df2565b888886818110610a9457610a946145e5565b9050602002810190610aa691906145f9565b60600135898987818110610abc57610abc6145e5565b9050602002810190610ace91906145f9565b610adc906080810190614617565b898881518110610aee57610aee6145e5565b60200260200101516118f3565b506001016109e8565b50600190505b9392505050565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa158015610b58573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b7c919061465c565b905090565b610b8c816001611df4565b50565b5f610b98611ebd565b89610ba2816117b7565b8190610bd257604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610bdb611eed565b5f5f5f610bea8e8c8c8c611f24565b9194509250905082828b8b84610c29576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610bc9565b50505050610c3e8e8e8e8e87878e8e8e611fe2565b60019450505050610c5b60015f516020614d0e5f395f51905f5255565b509998505050505050505050565b5f5b8151811015610c9f57610c97828281518110610c8957610c896145e5565b60200260200101515f611df4565b600101610c6b565b5050565b610ccd60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631290aa7d60e11b81526001600160a01b0384811660048301529091169063252154fa906024015b606060405180830381865afa158015610d16573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d3a91906146d5565b92915050565b5f6066546001610b7c9190614703565b5f610d6087338888888888610d8e565b979650505050505050565b610d7361206b565b610b8c816120c6565b610d8461206b565b610d8c612166565b565b5f610d97611ebd565b87610da1816117b7565b8190610dcc57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bc9565b50610dd5611eed565b5f5f5f610de48c8b8b8b611f24565b9194509250905082828a8a84610e23576040516335c20af360e11b81526004810194909452602484019290925260448301526064820152608401610bc9565b50505050610e3e8c610e323390565b8d8d87875f8e8e6121bf565b60019450505050610e5b60015f516020614d0e5f395f51905f5255565b50979650505050505050565b610b8c815f611df4565b610e796121fb565b610e828261229f565b610c9f82826122a7565b5f610d3a8183612368565b5f610ea989898a8a8a8a8a8a8a610b8f565b9998505050505050505050565b5f610ebf612373565b505f516020614cce5f395f51905f525b90565b610eda61206b565b610c9f82826123bc565b5f610d3a818361242b565b610ef761206b565b610b8c8161244c565b60605f610f0c5f6124b4565b9050806001600160401b03811115610f2657610f26613b59565b604051908082528060200260200182016040528015610f4f578160200160208202803683370190505b509150610f5b5f6124bd565b91505090565b610fa16040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b828210156110b3578382905f5260205f2001805461102890614716565b80601f016020809104026020016040519081016040528092919081815260200182805461105490614716565b801561109f5780601f106110765761010080835404028352916020019161109f565b820191905f5260205f20905b81548152906001019060200180831161108257829003601f168201915b50505050508152602001906001019061100b565b505050915250909392505050565b6110c961206b565b610d8c5f6124c9565b6110fc60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b60ca54604051631d46d32760e21b8152600481018490526001600160a01b039091169063751b4c9c90602401610cfb565b61113561206b565b610b8c81612539565b61114661206b565b610d8c6125b9565b5f60608082808083815f516020614cae5f395f51905f52805490915015801561117957506001810154155b6111bd5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610bc9565b6111c5612601565b6111cd6126c1565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60ca54604051634594558f60e11b81526001600160a01b038481166004830152602482018490525f928392911690638b28ab1e906044016040805180830381865afa158015611248573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061126c919061474e565b909590945092505050565b5f5b8151811015610c9f576112a6828281518110611297576112976145e5565b60200260200101516001611df4565b600101611279565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015610b58573d5f5f3e3d5ffd5b60ca5460408051634e86e3b760e11b815290516060926001600160a01b031691639d0dc76e916004808301925f9291908290030181865afa15801561133c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610b7c9190810190614770565b5f61136c611ebd565b61137533611d49565b33906113a0576040516338905e7160e11b81526001600160a01b039091166004820152602401610bc9565b506113a9611eed565b6113b460cf836126ff565b82906113d65760405163473978bf60e01b8152600401610bc991815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156114e9578382905f5260205f2001805461145e90614716565b80601f016020809104026020016040519081016040528092919081815260200182805461148a90614716565b80156114d55780601f106114ac576101008083540402835291602001916114d5565b820191905f5260205f20905b8154815290600101906020018083116114b857829003601f168201915b505050505081526020019060010190611441565b505050508152505090505f5f61150c836020015184604001518560600151612716565b915091508181906115305760405162461bcd60e51b8152600401610bc99190614466565b5061153c60cf866128d9565b505f85815260ce6020526040812061155391613a28565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906115986004830182613a5f565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738660600151426040516115f8929190918252602082015260400190565b60405180910390a46001935050505061161d60015f516020614d0e5f395f51905f5255565b919050565b5f610b7c60976124b4565b61163561206b565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156116c25750825b90505f826001600160401b031660011480156116dd5750303b155b9050811580156116eb575080155b156117095760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561173357845460ff60401b1916600160401b1785555b61173d87876128e4565b60fe805460016001600160a01b03199182161790915560ff80549091166001600160a01b038a1617905583156117ad57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f6117c182610ee4565b8015610d3a5750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f6117f2610f00565b90505f81516001600160401b0381111561180e5761180e613b59565b604051908082528060200260200182016040528015611837578160200160208202803683370190505b5090505f5b82518110156118c65760335f84838151811061185a5761185a6145e5565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b03168282815181106118a6576118a66145e5565b6001600160a01b039092166020928302919091019091015260010161183c565b5092915050565b5f610d3a609783612368565b5f610b7c5f6124b4565b5f6065546001610b7c9190614703565b5f6118fc611ebd565b86611906816117b7565b819061193157604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bc9565b5061193a611eed565b5f611943610d40565b9050808a80821461197057604051634982205b60e11b815260048101929092526024820152604401610bc9565b50505f7fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8b8b8b8b8b8b6040516020016119b097969594939291906148c4565b6040516020818303038152906040528051906020012090506119d28186612990565b8b906119f457604051635b777d1160e11b8152600401610bc991815260200190565b50611a03606680546001019055565b5f5f611a108c8c8c612716565b915091508115611a76578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611a69929190918252602082015260400190565b60405180910390a4611b9c565b611a8160cf8e612ada565b8d90611aa3576040516368db995b60e11b8152600401610bc991815260200190565b505f8d815260ce60205260409020611abb8282614948565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611afe9190614a02565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611b6e9260048501920190613a7a565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b6001955050505050610e5b60015f516020614d0e5f395f51905f5255565b611bc261206b565b6001600160a01b038116611beb57604051631e4fbdf760e01b81525f6004820152602401610bc9565b610b8c816124c9565b6060610b7c60976124bd565b5f610b7c612ae5565b5f611c1261206b565b5f83604051602001611c249190614a25565b60408051601f19818403018152908290526bffffffffffffffffffffffff19606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f60ff5f9054906101000a90046001600160a01b03166001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611cbd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611ce49190810190614a46565b838787604051602001611cf993929190614ac5565b60408051601f1981840301815290829052611d179291602001614afd565b6040516020818303038152906040529050611d335f8383612aee565b9350611d3f8488610ed2565b5050509392505050565b5f610d3a60978361242b565b5f81815260ce60205260409020805460609190611d7190614716565b80601f0160208091040260200160405190810160405280929190818152602001828054611d9d90614716565b8015611de85780601f10611dbf57610100808354040283529160200191611de8565b820191905f5260205f20905b815481529060010190602001808311611dcb57829003601f168201915b50505050509050919050565b611dfc61206b565b8015611e3e57611e0d609783612b82565b8290611e385760405163082cdf5560e21b81526001600160a01b039091166004820152602401610bc9565b50611e76565b611e49609783612b96565b8290611e745760405163e491024560e01b81526001600160a01b039091166004820152602401610bc9565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614cee5f395f51905f525460ff1615610d8c5760405163d93c066560e01b815260040160405180910390fd5b5f516020614d0e5f395f51905f52805460011901611f1e57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f90819081906001600160a01b0316611f4857505f91508190506001611fd8565b60ca54604051634594558f60e11b81526001600160a01b0389811660048301526024820189905290911690638b28ab1e906044016040805180830381865afa158015611f96573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fba919061474e565b9093509150828510801590611fcf5750818410155b15611fd8575060015b9450945094915050565b83611fed8688614703565b611ff79190614703565b836040015110158987909161203057604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610bc9565b505061203b83612baa565b61204d898989898989600189896121bf565b505050505050505050565b60015f516020614d0e5f395f51905f5255565b3361209d7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610d8c5760405163118cdaa760e01b8152336004820152602401610bc9565b6120cf81610ee4565b80156120f257506001600160a01b0381165f9081526034602052604090205460ff165b819061211d576040516340da71e560e01b81526001600160a01b039091166004820152602401610bc9565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b61216e612cae565b5f516020614cee5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001611673565b6121d48989886121cf888a614703565b612cdd565b6121ed6121df6118e3565b8a8a8a8a8a8a8a8a8a612e83565b61204d606580546001019055565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061228157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166122755f516020614cce5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610d8c5760405163703e46dd60e11b815260040160405180910390fd5b610b8c61206b565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612301575060408051601f3d908101601f191682019092526122fe9181019061465c565b60015b61232957604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610bc9565b5f516020614cce5f395f51905f52811461235957604051632a87526960e21b815260048101829052602401610bc9565b6123638383612f49565b505050565b5f610b0a8383612f9e565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d8c5760405163703e46dd60e11b815260040160405180910390fd5b6123c582612fc4565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610b0a565b61245581613071565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce910161241f565b5f610d3a825490565b60605f610b0a8361311e565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b612542816117b7565b819061256d576040516340da71e560e01b81526001600160a01b039091166004820152602401610bc9565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b6125c1611ebd565b5f516020614cee5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336121a7565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614cae5f395f51905f529161263f90614716565b80601f016020809104026020016040519081016040528092919081815260200182805461266b90614716565b80156126b65780601f1061268d576101008083540402835291602001916126b6565b820191905f5260205f20905b81548152906001019060200180831161269957829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614cae5f395f51905f529161263f90614716565b5f8181526001830160205260408120541515610b0a565b5f606082156128d15760fe546001600160a01b039081169086160361275857612753612743606485614b19565b6001600160a01b03861690613176565b6128d1565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528616906340c10f19906044016020604051808303815f875af19250505080156127c2575060408051601f3d908101601f191682019092526127bf91810190614b30565b60015b612873576127ce614b4f565b806308c379a0036127f757506127e2614b67565b806127ed575061283a565b5f925090506128d1565b634e487b710361283a57612809614be9565b90612814575061283a565b60408051602081018390525f9450016040516020818303038152906040529150506128d1565b3d808015612863576040519150601f19603f3d011682016040523d82523d5f602084013e612868565b606091505b505f925090506128d1565b8015612893576001925060405180602001604052805f81525091506128cf565b5f92506040518060400160405280601881526020017f42726964676543726f73733a206d696e74206661696c6564000000000000000081525091505b505b935093915050565b5f610b0a8383613202565b6128ec6132e5565b6128f53361332e565b6128fd61333f565b612905613347565b61290d613357565b612915613367565b6001600160a01b03821661295b57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610bc9565b4360cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff168110156129ab575f915050610d3a565b5f80826001600160401b038111156129c5576129c5613b59565b6040519080825280602002602001820160405280156129ee578160200160208202803683370190505b5090505f5b8551811015612ac8575f868281518110612a0f57612a0f6145e5565b60200260200101519050604181511015612a30575f95505050505050610d3a565b5f612a4482612a3e8b6133c6565b906133f2565b9050612a4f81611d49565b612a61575f9650505050505050610d3a565b5f805b8551811015612ab057858181518110612a7f57612a7f6145e5565b60200260200101516001600160a01b0316836001600160a01b031603612aa85760019150612ab0565b600101612a64565b5080612abd578560010195505b5050506001016129f3565b505060995460ff161115949350505050565b5f610b0a838361341a565b5f610b7c613466565b5f83471015612b195760405163cf47918160e01b815247600482015260248101859052604401610bc9565b81515f03612b3a57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615612b5b576040513d5f823e3d81fd5b6001600160a01b038116610b0a5760405163b06ebf3d60e01b815260040160405180910390fd5b5f610b0a836001600160a01b03841661341a565b5f610b0a836001600160a01b038416613202565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612c56576040513d5f823e3d81fd5b50505f513d91508115612c6d578060011415612c7b565b84516001600160a01b03163b155b15612ca7578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610bc9565b5050505050565b5f516020614cee5f395f51905f525460ff16610d8c57604051638dfc202b60e01b815260040160405180910390fd5b60fe546001600160a01b0390811690851603612d9e57346064612d008183614c06565b612d0a9190614b19565b143490612d2d576040516308dc47c960e41b8152600401610bc991815260200190565b50612d388183614703565b3414612d448284614703565b349091612d6c576040516290c38760e71b815260048101929092526024820152604401610bc9565b50508015612d9957612d9981612d8a60cc546001600160a01b031690565b6001600160a01b031690613176565b612e7d565b8015612dcc57612dcc83612dba60cc546001600160a01b031690565b6001600160a01b0387169190846134d9565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201849052851690639dc29fac906044016020604051808303815f875af1158015612e18573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e3c9190614b30565b848484909192612e79576040516336b52daf60e21b81526001600160a01b0393841660048201529290911660248301526044820152606401610bc9565b5050505b50505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a604051612ede9796959493929190614c25565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051612f35929190918252602082015260400190565b60405180910390a450505050505050505050565b612f5282613533565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612f96576123638282613596565b610c9f613608565b5f825f018281548110612fb357612fb36145e5565b905f5260205f200154905092915050565b806001600160a01b038116613004576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bc9565b61300e5f83612b82565b8290613039576040516351eccfe160e11b81526001600160a01b039091166004820152602401610bc9565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b0381166130b1576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bc9565b6130bb5f83612b96565b82906130e6576040516340da71e560e01b81526001600160a01b039091166004820152602401610bc9565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611de857602002820191905f5260205f20905b8154815260200190600101908083116131575750505050509050919050565b804710156131a05760405163cf47918160e01b815247600482015260248101829052604401610bc9565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f81146131ea576040519150601f19603f3d011682016040523d82523d5f602084013e6131ef565b606091505b509150915081612e7d57612e7d81613627565b5f81815260018301602052604081205480156132dc575f613224600183614c67565b85549091505f9061323790600190614c67565b9050808214613296575f865f018281548110613255576132556145e5565b905f5260205f200154905080875f018481548110613275576132756145e5565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806132a7576132a7614c7a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610d3a565b5f915050610d3a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610d8c57604051631afcd79f60e31b815260040160405180910390fd5b6133366132e5565b610b8c81613650565b610d8c6132e5565b61334f6132e5565b610d8c613658565b61335f6132e5565b610d8c613678565b61336f6132e5565b6133b7604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613680565b6099805460ff19166003179055565b5f610d3a6133d2612ae5565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6134008686613692565b92509250925061341082826136db565b5090949350505050565b5f81815260018301602052604081205461345f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610d3a565b505f610d3a565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613490613793565b6134986137fb565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612e7d90859061383d565b806001600160a01b03163b5f0361356857604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610bc9565b5f516020614cce5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516135b29190614c8e565b5f60405180830381855af49150503d805f81146135ea576040519150601f19603f3d011682016040523d82523d5f602084013e6135ef565b606091505b50915091506135ff8583836138a9565b95945050505050565b3415610d8c5760405163b398979f60e01b815260040160405180910390fd5b8051156136375780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b611bc26132e5565b6136606132e5565b5f516020614cee5f395f51905f52805460ff19169055565b6120586132e5565b6136886132e5565b610c9f8282613905565b5f5f5f83516041036136c9576020840151604085015160608601515f1a6136bb88828585613964565b9550955095505050506136d4565b505081515f91506002905b9250925092565b5f8260038111156136ee576136ee614c99565b036136f7575050565b600182600381111561370b5761370b614c99565b036137295760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561373d5761373d614c99565b0361375e5760405163fce698f760e01b815260048101829052602401610bc9565b600382600381111561377257613772614c99565b03610c9f576040516335e2f38360e21b815260048101829052602401610bc9565b5f5f516020614cae5f395f51905f52816137ab612601565b8051909150156137c357805160209091012092915050565b815480156137d2579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614cae5f395f51905f52816138136126c1565b80519091501561382b57805160209091012092915050565b600182015480156137d2579392505050565b5f5f60205f8451602086015f885af18061385c576040513d5f823e3d81fd5b50505f513d91508115613873578060011415613880565b6001600160a01b0384163b155b15612e7d57604051635274afe760e01b81526001600160a01b0385166004820152602401610bc9565b6060826138be576138b982613627565b610b0a565b81511580156138d557506001600160a01b0384163b155b156138fe57604051639996b31560e01b81526001600160a01b0385166004820152602401610bc9565b5080610b0a565b61390d6132e5565b5f516020614cae5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026139468482614948565b50600381016139558382614948565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561399d57505f91506003905082611fd8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156139ee573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613a1957505f925060019150829050611fd8565b975f9750879650945050505050565b508054613a3490614716565b5f825580601f10613a43575050565b601f0160209004905f5260205f2090810190610b8c9190613ace565b5080545f8255905f5260205f2090810190610b8c9190613ae2565b828054828255905f5260205f20908101928215613abe579160200282015b82811115613abe5782518290613aae9082614948565b5091602001919060010190613a98565b50613aca929150613ae2565b5090565b5b80821115613aca575f8155600101613acf565b80821115613aca575f613af58282613a28565b50600101613ae2565b634e487b7160e01b5f52600160045260245ffd5b5f5f83601f840112613b22575f5ffd5b5081356001600160401b03811115613b38575f5ffd5b6020830191508360208260051b8501011115613b52575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b0382111715613b8c57613b8c613b59565b60405250565b601f8201601f191681016001600160401b0381118282101715613bb757613bb7613b59565b6040525050565b5f6001600160401b03821115613bd657613bd6613b59565b5060051b60200190565b5f6001600160401b03821115613bf857613bf8613b59565b50601f01601f191660200190565b5f613c1083613be0565b604051613c1d8282613b92565b809250848152858585011115613c31575f5ffd5b848460208301375f6020868301015250509392505050565b5f82601f830112613c58575f5ffd5b610b0a83833560208501613c06565b5f613c7183613bbe565b604051613c7e8282613b92565b84815291505060208101600584901b830185811115613c9b575f5ffd5b835b81811015611d3f5780356001600160401b03811115613cba575f5ffd5b613cc688828801613c49565b84525060209283019201613c9d565b5f82601f830112613ce4575f5ffd5b610b0a83833560208501613c67565b5f5f5f60408486031215613d05575f5ffd5b83356001600160401b03811115613d1a575f5ffd5b613d2686828701613b12565b90945092505060208401356001600160401b03811115613d44575f5ffd5b8401601f81018613613d54575f5ffd5b8035613d5f81613bbe565b604051613d6c8282613b92565b80915082815260208101915060208360051b850101925088831115613d8f575f5ffd5b602084015b83811015613dcf5780356001600160401b03811115613db1575f5ffd5b613dc08b602083890101613cd5565b84525060209283019201613d94565b50809450505050509250925092565b6001600160a01b0381168114610b8c575f5ffd5b5f60208284031215613e02575f5ffd5b8135610b0a81613dde565b803560ff8116811461161d575f5ffd5b5f60e08284031215613e2d575f5ffd5b604051613e3981613b6d565b8091508235613e4781613dde565b81526020830135613e5781613dde565b60208201526040838101359082015260608084013590820152613e7c60808401613e0d565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613eb5575f5ffd5b8935613ec081613dde565b985060208a0135613ed081613dde565b975060408a0135613ee081613dde565b965060608a0135955060808a0135945060a08a01359350613f048b60c08c01613e1d565b92506101a08a01356001600160401b03811115613f1f575f5ffd5b613f2b8c828d01613b12565b915080935050809150509295985092959850929598565b5f60208284031215613f52575f5ffd5b81356001600160401b03811115613f67575f5ffd5b8201601f81018413613f77575f5ffd5b8035613f8281613bbe565b604051613f8f8282613b92565b80915082815260208101915060208360051b850101925086831115613fb2575f5ffd5b6020840193505b82841015613fdd578335613fcc81613dde565b825260209384019390910190613fb9565b9695505050505050565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610d3a565b5f5f5f5f5f5f60a08789031215614026575f5ffd5b863561403181613dde565b955060208701359450604087013593506060870135925060808701356001600160401b03811115614060575f5ffd5b61406c89828a01613b12565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215614094575f5ffd5b873561409f81613dde565b965060208801356140af81613dde565b955060408801359450606088013593506080880135925060a08801356001600160401b038111156140de575f5ffd5b6140ea8a828b01613b12565b989b979a50959850939692959293505050565b5f5f6040838503121561410e575f5ffd5b823561411981613dde565b915060208301356001600160401b03811115614133575f5ffd5b61413f85828601613c49565b9150509250929050565b5f60208284031215614159575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b031215614178575f5ffd5b883561418381613dde565b9750602089013561419381613dde565b96506040890135955060608901359450608089013593506141b78a60a08b01613e1d565b92506101808901356001600160401b038111156141d2575f5ffd5b6141de8b828c01613b12565b999c989b5096995094979396929594505050565b5f5f60408385031215614203575f5ffd5b823561420e81613dde565b9150602083013561421e81613dde565b809150509250929050565b602080825282518282018190525f918401906040840190835b818110156142695783516001600160a01b0316835260209384019390920191600101614242565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b8181101561433b5760df19878603018352614326858551614274565b9450602093840193929092019160010161430a565b50929695505050505050565b60ff60f81b8816815260e060208201525f61436560e0830189614274565b82810360408401526143778189614274565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156143cc5783518352602093840193909201916001016143ae565b50909b9a5050505050505050505050565b5f5f604083850312156143ee575f5ffd5b82356143f981613dde565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156142695761445083855180516001600160a01b0316825260208082015190830152604090810151910152565b6020939093019260609290920191600101614420565b602081525f610b0a6020830184614274565b5f60208284031215614488575f5ffd5b610b0a82613e0d565b5f5f5f606084860312156144a3575f5ffd5b83356144ae81613dde565b925060208401356144be81613dde565b915060408401356144ce81613dde565b809150509250925092565b5f5f5f5f5f5f5f60c0888a0312156144ef575f5ffd5b87359650602088013561450181613dde565b9550604088013561451181613dde565b94506060880135935060808801356001600160401b03811115614532575f5ffd5b61453e8a828b01613b12565b90945092505060a08801356001600160401b0381111561455c575f5ffd5b6145688a828b01613cd5565b91505092959891949750929550565b5f5f5f60608486031215614589575f5ffd5b833561459481613dde565b925060208401356001600160401b038111156145ae575f5ffd5b8401601f810186136145be575f5ffd5b6145cd86823560208401613c06565b9250506145dc60408501613e0d565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e1983360301811261460d575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261462c575f5ffd5b8301803591506001600160401b03821115614645575f5ffd5b6020019150600581901b3603821315613b52575f5ffd5b5f6020828403121561466c575f5ffd5b5051919050565b5f60608284031215614683575f5ffd5b604051606081018181106001600160401b03821117156146a5576146a5613b59565b806040525080915082516146b881613dde565b815260208381015190820152604092830151920191909152919050565b5f606082840312156146e5575f5ffd5b610b0a8383614673565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d3a57610d3a6146ef565b600181811c9082168061472a57607f821691505b60208210810361474857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f5f6040838503121561475f575f5ffd5b505080516020909101519092909150565b5f60208284031215614780575f5ffd5b81516001600160401b03811115614795575f5ffd5b8201601f810184136147a5575f5ffd5b80516147b081613bbe565b6040516147bd8282613b92565b82815260206060909302840183019281019150868311156147dc575f5ffd5b6020840193505b82841015613fdd576147f58785614673565b82526020820191506060840193506147e3565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b868110156148b857838303601f19018852813536879003601e1901811261486c575f5ffd5b86016020810190356001600160401b03811115614887575f5ffd5b803603821315614895575f5ffd5b6148a0858284614808565b60209a8b019a90955093909301925050600101614847565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90610ea99083018486614830565b601f82111561236357805f5260205f20601f840160051c810160208510156149295750805b601f840160051c820191505b81811015612ca7575f8155600101614935565b81516001600160401b0381111561496157614961613b59565b6149758161496f8454614716565b84614904565b6020601f8211600181146149a7575f83156149905750848201515b5f19600385901b1c1916600184901b178455612ca7565b5f84815260208120601f198516915b828110156149d657878501518255602094850194600190920191016149b6565b50848210156149f357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610b0a368484613c67565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f610b0a600d830184614a0e565b5f60208284031215614a56575f5ffd5b81516001600160401b03811115614a6b575f5ffd5b8201601f81018413614a7b575f5ffd5b8051614a8681613be0565b604051614a938282613b92565b828152866020848601011115614aa7575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b606081525f614ad76060830186614274565b8281036020840152614ae98186614274565b91505060ff83166040830152949350505050565b5f614b11614b0b8386614a0e565b84614a0e565b949350505050565b8082028115828204841417610d3a57610d3a6146ef565b5f60208284031215614b40575f5ffd5b81518015158114610b0a575f5ffd5b5f60033d1115610ecf5760045f5f3e505f5160e01c90565b5f60443d1015614b745790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614b9d57505090565b80820180516001600160401b03811115614bb8575050505090565b3d8401600319018282016020011115614bd2575050505090565b614be160208285010185613b92565b509392505050565b5f5f60233d1115614c0257602060045f3e50505f516001905b9091565b5f82614c2057634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f90610ea99083018486614830565b81810381811115610d3a57610d3a6146ef565b634e487b7160e01b5f52603160045260245ffd5b5f610b0a8284614a0e565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220ede25f0218fcaf5a97dee90089dd00bdab2900d2ae7a386ad56ae7d8c707b6fa64736f6c634300081c0033",
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

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeCross *BridgeCrossCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeCross *BridgeCrossSession) AllValidators() ([]common.Address, error) {
	return _BridgeCross.Contract.AllValidators(&_BridgeCross.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BridgeCross *BridgeCrossCallerSession) AllValidators() ([]common.Address, error) {
	return _BridgeCross.Contract.AllValidators(&_BridgeCross.CallOpts)
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

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeCross *BridgeCrossCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeCross *BridgeCrossSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BridgeCross.Contract.ValidatorByIndex(&_BridgeCross.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BridgeCross.Contract.ValidatorByIndex(&_BridgeCross.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) ValidatorLength() (*big.Int, error) {
	return _BridgeCross.Contract.ValidatorLength(&_BridgeCross.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) ValidatorLength() (*big.Int, error) {
	return _BridgeCross.Contract.ValidatorLength(&_BridgeCross.CallOpts)
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

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridge", token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, service, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 service, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, service *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, service, permitArgs, extraData)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeCross *BridgeCrossTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeCross *BridgeCrossSession) Receive() (*types.Transaction, error) {
	return _BridgeCross.Contract.Receive(&_BridgeCross.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeCross *BridgeCrossTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeCross.Contract.Receive(&_BridgeCross.TransactOpts)
}

// BridgeCrossBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the BridgeCross contract.
type BridgeCrossBridgeFeeChargedIterator struct {
	Event *BridgeCrossBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *BridgeCrossBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossBridgeFeeCharged)
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
		it.Event = new(BridgeCrossBridgeFeeCharged)
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
func (it *BridgeCrossBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossBridgeFeeCharged represents a BridgeFeeCharged event raised by the BridgeCross contract.
type BridgeCrossBridgeFeeCharged struct {
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
func (_BridgeCross *BridgeCrossFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*BridgeCrossBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _BridgeCross.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCrossBridgeFeeChargedIterator{contract: _BridgeCross.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 serviceFee)
func (_BridgeCross *BridgeCrossFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *BridgeCrossBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeCross.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossBridgeFeeCharged)
				if err := _BridgeCross.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_BridgeCross *BridgeCrossFilterer) ParseBridgeFeeCharged(log types.Log) (*BridgeCrossBridgeFeeCharged, error) {
	event := new(BridgeCrossBridgeFeeCharged)
	if err := _BridgeCross.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe854093.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes[] extraData)
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe854093.
//
// Solidity: event BridgeInitiated(uint256 indexed index, address indexed token, address pairToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes[] extraData)
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
