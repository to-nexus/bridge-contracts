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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addTokenDeploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenInfo\",\"outputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"changeRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossMintableERC20Code\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"f7684ec4": "addTokenDeploy(address,string,uint8)",
		"c97682f8": "allPairs()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"9c1b65a9": "bridgeTokenInfo()",
		"6e908ca3": "calculate(address,uint256)",
		"abe110b6": "changeRewardWallet(address)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"f120c400": "finalize(uint256,address,address,uint256,bytes[],bytes[])",
		"008bd028": "finalizeBatch((uint256,address,address,uint256,bytes[])[],bytes[][])",
		"71c59d7b": "getPairToken(address)",
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
		"3fc8cef3": "weth()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614bee6100f95f395f81816121420152818161216b01526122ba0152614bee5ff3fe608060405260043610610343575f3560e01c80638456cb59116101bd578063c1876453116100f2578063f2fde38b11610092578063f7684ec41161006d578063f7684ec414610969578063facd743b14610988578063fb75b2c7146109a7578063fe2b8da6146109c4575f5ffd5b8063f2fde38b14610922578063f30589c314610941578063f698da2514610955575f5ffd5b8063cbae5958116100cd578063cbae5958146108c8578063d92fc67b146108e7578063e1af7f50146108fb578063f120c4001461090f575f5ffd5b8063c187645314610876578063c1ad8b9514610895578063c97682f8146108b4575f5ffd5b80639c1b65a91161015d578063ad3cb1cc11610138578063ad3cb1cc146107e7578063aed1d40314610824578063b7f3358d14610838578063c0c53b8b14610857575f5ffd5b80639c1b65a91461078a578063a9823183146107a9578063abe110b6146107c8575f5ffd5b80638da5cb5b116101985780638da5cb5b1461070757806391cf6d3e146107435780639300c9261461075757806396ce079514610776575f5ffd5b80638456cb59146106b857806384b0196e146106cc57806385547884146106f3575f5ffd5b80634f1ef286116102935780635fa7b584116102335780637021fd0e1161020e5780637021fd0e14610622578063715018a61461064e57806371c59d7b146106625780637c41ad2c14610699575f5ffd5b80635fa7b584146105a15780636e908ca3146105c05780636ff97f1d14610601575f5ffd5b806352d1902d1161026e57806352d1902d1461052c5780635476bd72146105405780635c975abb1461055f5780635dbe47e814610582575f5ffd5b80634f1ef286146104e75780634f6ccce7146104fa57806351c4557914610519575f5ffd5b806337d10075116102fe5780633f5d3b5d116102d95780633f5d3b5d1461045d5780633fc8cef31461047057806340a141ff146104a757806342cde4e8146104c6575f5ffd5b806337d10075146104175780633b3bff0f1461042a5780633f4ba83a14610449575f5ffd5b80628bd0281461035d5780631327d3d814610385578063174991ab146103a45780631d40f0d8146103b757806327938c81146103d65780632f9b59d1146103f8575f5ffd5b3661035957345f0361035757610357613a70565b005b5f5ffd5b61037061036b366004613c65565b6109e3565b60405190151581526020015b60405180910390f35b348015610390575f5ffd5b5061035761039f366004613d64565b610b0f565b6103706103b2366004613e0e565b610b1d565b3480156103c2575f5ffd5b506103576103d1366004613eb4565b610bb6565b3480156103e1575f5ffd5b506103ea610bf0565b60405190815260200161037c565b348015610403575f5ffd5b50610357610412366004613d64565b610c05565b610370610425366004613f59565b610c79565b348015610435575f5ffd5b50610357610444366004613d64565b610c94565b348015610454575f5ffd5b50610357610ca5565b61037061046b366004613fc6565b610cb7565b34801561047b575f5ffd5b5060fd5461048f906001600160a01b031681565b6040516001600160a01b03909116815260200161037c565b3480156104b2575f5ffd5b506103576104c1366004613d64565b610d4f565b3480156104d1575f5ffd5b5060995460405160ff909116815260200161037c565b6103576104f5366004614045565b610d59565b348015610505575f5ffd5b5061048f610514366004614091565b610d74565b6103706105273660046140a8565b610d85565b348015610537575f5ffd5b506103ea610da4565b34801561054b575f5ffd5b5061035761055a36600461413a565b610dc0565b34801561056a575f5ffd5b505f516020614b795f395f51905f525460ff16610370565b34801561058d575f5ffd5b5061037061059c366004613d64565b610dd2565b3480156105ac575f5ffd5b506103576105bb366004613d64565b610ddd565b3480156105cb575f5ffd5b506105df6105da366004614171565b610dee565b604080519485526020850193909352918301521515606082015260800161037c565b34801561060c575f5ffd5b50610615610e77565b60405161037c919061419b565b34801561062d575f5ffd5b5061064161063c366004614091565b610ed8565b60405161037c9190614214565b348015610659575f5ffd5b50610357611038565b34801561066d575f5ffd5b5061048f61067c366004613d64565b6001600160a01b039081165f908152603360205260409020541690565b3480156106a4575f5ffd5b506103576106b3366004613d64565b611049565b3480156106c3575f5ffd5b5061035761105a565b3480156106d7575f5ffd5b506106e061106a565b60405161037c97969594939291906142b9565b3480156106fe575f5ffd5b50610357611113565b348015610712575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661048f565b34801561074e575f5ffd5b5060cd546103ea565b348015610762575f5ffd5b50610357610771366004613eb4565b61112d565b348015610781575f5ffd5b506103ea611164565b348015610795575f5ffd5b5060ca5461048f906001600160a01b031681565b3480156107b4575f5ffd5b506103706107c3366004614091565b6111cf565b3480156107d3575f5ffd5b506103576107e2366004613d64565b611459565b3480156107f2575f5ffd5b50610817604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161037c919061434f565b34801561082f575f5ffd5b506103ea611483565b348015610843575f5ffd5b50610357610852366004614361565b61148e565b348015610862575f5ffd5b5061035761087136600461437a565b6114df565b348015610881575f5ffd5b50610370610890366004613d64565b611651565b3480156108a0575f5ffd5b506103706108af3660046143c2565b611681565b3480156108bf575f5ffd5b506106156116c1565b3480156108d3575f5ffd5b5061048f6108e2366004614091565b6117a7565b3480156108f2575f5ffd5b506103ea6117b3565b348015610906575f5ffd5b506103ea6117bd565b61037061091d366004614454565b6117cd565b34801561092d575f5ffd5b5061035761093c366004613d64565b611a94565b34801561094c575f5ffd5b50610615611ace565b348015610960575f5ffd5b506103ea611ada565b348015610974575f5ffd5b5061048f6109833660046144f2565b611ae3565b348015610993575f5ffd5b506103706109a2366004613d64565b611c23565b3480156109b2575f5ffd5b5060cc546001600160a01b031661048f565b3480156109cf575f5ffd5b506108176109de366004614091565b611c2f565b5f805b83811015610b0257610af9858583818110610a0357610a03614560565b9050602002810190610a159190614574565b35868684818110610a2857610a28614560565b9050602002810190610a3a9190614574565b610a4b906040810190602001613d64565b878785818110610a5d57610a5d614560565b9050602002810190610a6f9190614574565b610a80906060810190604001613d64565b888886818110610a9257610a92614560565b9050602002810190610aa49190614574565b60600135898987818110610aba57610aba614560565b9050602002810190610acc9190614574565b610ada906080810190614592565b898881518110610aec57610aec614560565b60200260200101516117cd565b506001016109e6565b50600190505b9392505050565b610b1a816001611cce565b50565b5f610b26611d97565b89610b3081611651565b8190610b6057604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610b69611dc7565b5f5f610b778d8b8b8b611dfe565b91509150610b8c8d8d8d8d86868d8d8d611f1e565b600193505050610ba860015f516020614b995f395f51905f5255565b509998505050505050505050565b5f5b8151811015610bec57610be4828281518110610bd657610bd6614560565b60200260200101515f611cce565b600101610bb8565b5050565b5f6066546001610c0091906145eb565b905090565b610c0d611fa7565b6001600160a01b038116610c5757604051633effd40360e21b815260206004820152601060248201526f5f627269646765546f6b656e496e666f60801b6044820152606401610b57565b60ca80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610c8987338888888888610cb7565b979650505050505050565b610c9c611fa7565b610b1a81612002565b610cad611fa7565b610cb56120a2565b565b5f610cc0611d97565b87610cca81611651565b8190610cf557604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610b57565b50610cfe611dc7565b5f5f610d0c8b8a8a8a611dfe565b91509150610d278b610d1b3390565b8c8c86865f8d8d6120fb565b600193505050610d4360015f516020614b995f395f51905f5255565b50979650505050505050565b610b1a815f611cce565b610d61612137565b610d6a826121db565b610bec82826121e3565b5f610d7f81836122a4565b92915050565b5f610d9789898a8a8a8a8a8a8a610b1d565b9998505050505050505050565b5f610dad6122af565b505f516020614b595f395f51905f525b90565b610dc8611fa7565b610bec82826122f8565b5f610d7f8183612367565b610de5611fa7565b610b1a81612388565b60ca54604051636e908ca360e01b81526001600160a01b038481166004830152602482018490525f928392839283921690636e908ca390604401608060405180830381865afa158015610e43573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e67919061460d565b9299919850965090945092505050565b60605f610e835f6123f0565b9050806001600160401b03811115610e9d57610e9d613acb565b604051908082528060200260200182016040528015610ec6578160200160208202803683370190505b509150610ed25f6123f9565b91505090565b610f186040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260ce60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b8282101561102a578382905f5260205f20018054610f9f9061464a565b80601f0160208091040260200160405190810160405280929190818152602001828054610fcb9061464a565b80156110165780601f10610fed57610100808354040283529160200191611016565b820191905f5260205f20905b815481529060010190602001808311610ff957829003601f168201915b505050505081526020019060010190610f82565b505050915250909392505050565b611040611fa7565b610cb55f612405565b611051611fa7565b610b1a81612475565b611062611fa7565b610cb56124f5565b5f60608082808083815f516020614b395f395f51905f52805490915015801561109557506001810154155b6110d95760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610b57565b6110e161253d565b6110e96125fd565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61111b611fa7565b60ca80546001600160a01b0319169055565b5f5b8151811015610bec5761115c82828151811061114d5761114d614560565b60200260200101516001611cce565b60010161112f565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa1580156111ab573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c009190614682565b5f6111d8611d97565b6111e0611dc7565b6111eb60d08361263b565b829061120d5760405163473978bf60e01b8152600401610b5791815260200190565b505f82815260ce60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b82821015611320578382905f5260205f200180546112959061464a565b80601f01602080910402602001604051908101604052809291908181526020018280546112c19061464a565b801561130c5780601f106112e35761010080835404028352916020019161130c565b820191905f5260205f20905b8154815290600101906020018083116112ef57829003601f168201915b505050505081526020019060010190611278565b505050508152505090505f5f611343836020015184604001518560600151612652565b915091508181906113675760405162461bcd60e51b8152600401610b57919061434f565b5061137360d086612841565b505f85815260cf6020526040812061138a9161399a565b5f85815260ce602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906113cf60048301826139d1565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a37386606001514260405161142f929190918252602082015260400190565b60405180910390a46001935050505061145460015f516020614b995f395f51905f5255565b919050565b611461611fa7565b60cc80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610c0060976123f0565b611496611fa7565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156115235750825b90505f826001600160401b0316600114801561153e5750303b155b90508115801561154c575080155b1561156a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561159457845460ff60401b1916600160401b1785555b61159e878761284c565b60fe80546001600160a01b0319166001600160a01b038a1617905560408051808201909152600381526208aa8960eb1b60208201526115e1906001906012611ae3565b60fd80546001600160a01b0319166001600160a01b0392909216919091179055831561164757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f61165b82610dd2565b8015610d7f5750506001600160a01b03165f9081526034602052604090205460ff161590565b5f805b82518110156116b8576116af8382815181106116a2576116a2614560565b60200260200101516111cf565b50600101611684565b50600192915050565b60605f6116cc610e77565b90505f81516001600160401b038111156116e8576116e8613acb565b604051908082528060200260200182016040528015611711578160200160208202803683370190505b5090505f5b82518110156117a05760335f84838151811061173457611734614560565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b031682828151811061178057611780614560565b6001600160a01b0390921660209283029190910190910152600101611716565b5092915050565b5f610d7f6097836122a4565b5f610c005f6123f0565b5f6065546001610c0091906145eb565b5f6117d6611d97565b866117e081611651565b819061180b57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610b57565b50611814611dc7565b5f61181d610bf0565b9050808a80821461184a57604051634982205b60e11b815260048101929092526024820152604401610b57565b50505f7fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8b8b8b8b8b8b60405160200161188a9796959493929190614755565b6040516020818303038152906040528051906020012090506118ac81866128fc565b8b906118ce57604051635b777d1160e11b8152600401610b5791815260200190565b506118dd606680546001019055565b5f5f6118ea8c8c8c612652565b915091508115611950578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611943929190918252602082015260400190565b60405180910390a4611a76565b61195b60d08e612a49565b8d9061197d576040516368db995b60e11b8152600401610b5791815260200190565b505f8d815260cf6020526040902061199582826147d9565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a906119d89190614893565b90525f8e815260ce602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611a4892600485019201906139ec565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b6001955050505050610d4360015f516020614b995f395f51905f5255565b611a9c611fa7565b6001600160a01b038116611ac557604051631e4fbdf760e01b81525f6004820152602401610b57565b610b1a81612405565b6060610c0060976123f9565b5f610c00612a54565b5f611aec611fa7565b5f83604051602001611afe91906148b6565b60408051601f19818403018152908290526bffffffffffffffffffffffff19606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f60fe5f9054906101000a90046001600160a01b03166001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611b97573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611bbe91908101906148d7565b838787604051602001611bd393929190614956565b60408051601f1981840301815290829052611bf1929160200161498e565b6040516020818303038152906040529050611c0d5f8383612a5d565b9350611c198488610dc0565b5050509392505050565b5f610d7f609783612367565b5f81815260cf60205260409020805460609190611c4b9061464a565b80601f0160208091040260200160405190810160405280929190818152602001828054611c779061464a565b8015611cc25780601f10611c9957610100808354040283529160200191611cc2565b820191905f5260205f20905b815481529060010190602001808311611ca557829003601f168201915b50505050509050919050565b611cd6611fa7565b8015611d1857611ce7609783612af1565b8290611d125760405163082cdf5560e21b81526001600160a01b039091166004820152602401610b57565b50611d50565b611d23609783612b05565b8290611d4e5760405163e491024560e01b81526001600160a01b039091166004820152602401610b57565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614b795f395f51905f525460ff1615610cb55760405163d93c066560e01b815260040160405180910390fd5b5f516020614b995f395f51905f52805460011901611df857604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f9081906001600160a01b0316611e1c57505f905080611f15565b60ca54604051636e908ca360e01b81526001600160a01b038881166004830152602482018890525f928392911690636e908ca390604401608060405180830381865afa158015611e6e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e92919061460d565b919650945092509050808710801590611eab5750838610155b8015611eb75750828510155b8015611ec05750815b8185858a8a8a909192939495611f0c576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610b57565b50505050505050505b94509492505050565b83611f2986886145eb565b611f3391906145eb565b8360400151101589879091611f6c57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610b57565b5050611f7783612b19565b611f89898989898989600189896120fb565b505050505050505050565b60015f516020614b995f395f51905f5255565b33611fd97f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610cb55760405163118cdaa760e01b8152336004820152602401610b57565b61200b81610dd2565b801561202e57506001600160a01b0381165f9081526034602052604090205460ff165b8190612059576040516340da71e560e01b81526001600160a01b039091166004820152602401610b57565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b6120aa612c1d565b5f516020614b795f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016114d4565b61211089898861210b888a6145eb565b612c4c565b61212961211b6117bd565b8a8a8a8a8a8a8a8a8a612de7565b611f89606580546001019055565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806121bd57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166121b15f516020614b595f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610cb55760405163703e46dd60e11b815260040160405180910390fd5b610b1a611fa7565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561223d575060408051601f3d908101601f1916820190925261223a91810190614682565b60015b61226557604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610b57565b5f516020614b595f395f51905f52811461229557604051632a87526960e21b815260048101829052602401610b57565b61229f8383612ead565b505050565b5f610b088383612f02565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cb55760405163703e46dd60e11b815260040160405180910390fd5b61230182612f28565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610b08565b61239181612fd5565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce910161235b565b5f610d7f825490565b60605f610b0883613082565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61247e81611651565b81906124a9576040516340da71e560e01b81526001600160a01b039091166004820152602401610b57565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b6124fd611d97565b5f516020614b795f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336120e3565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614b395f395f51905f529161257b9061464a565b80601f01602080910402602001604051908101604052809291908181526020018280546125a79061464a565b80156125f25780601f106125c9576101008083540402835291602001916125f2565b820191905f5260205f20905b8154815290600101906020018083116125d557829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614b395f395f51905f529161257b9061464a565b5f8181526001830160205260408120541515610b08565b5f6060825f0361267457505060408051602081019091525f8152600190612839565b60cb546001600160a01b03908116908616036126c0576126a86126986064856149aa565b6001600160a01b038616906130da565b505060408051602081019091525f8152600190612839565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528616906340c10f19906044016020604051808303815f875af192505050801561272a575060408051601f3d908101601f19168201909252612727918101906149c1565b60015b6127db576127366149da565b806308c379a00361275f575061274a6149f2565b8061275557506127a2565b5f92509050612839565b634e487b71036127a257612771614a74565b9061277c57506127a2565b60408051602081018390525f945001604051602081830303815290604052915050612839565b3d8080156127cb576040519150601f19603f3d011682016040523d82523d5f602084013e6127d0565b606091505b505f92509050612839565b80156127fb576001925060405180602001604052805f8152509150612837565b5f92506040518060400160405280601881526020017f42726964676543726f73733a206d696e74206661696c6564000000000000000081525091505b505b935093915050565b5f610b088383613166565b612854613249565b61285d33613292565b6128656132fb565b61286d613303565b612875613313565b6001600160a01b0382166128bb57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610b57565b60cb80546001600160a01b03199081166001179091554360cd5560ca80546001600160a01b0393841690831617905560cc8054939092169216919091179055565b80516099545f919060ff16811015612917575f915050610d7f565b5f80826001600160401b0381111561293157612931613acb565b60405190808252806020026020018201604052801561295a578160200160208202803683370190505b5090505f5b8551811015612a37575f86828151811061297b5761297b614560565b6020026020010151905060418151101561299c575f95505050505050610d7f565b5f6129b0826129aa8b613323565b9061334f565b90506129bb81611c23565b6129cd575f9650505050505050610d7f565b5f805b8551811015612a1c578581815181106129eb576129eb614560565b60200260200101516001600160a01b0316836001600160a01b031603612a145760019150612a1c565b6001016129d0565b5080612a29578560010195505b50505080600101905061295f565b505060995460ff161115949350505050565b5f610b088383613377565b5f610c006133c3565b5f83471015612a885760405163cf47918160e01b815247600482015260248101859052604401610b57565b81515f03612aa957604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615612aca576040513d5f823e3d81fd5b6001600160a01b038116610b085760405163b06ebf3d60e01b815260040160405180910390fd5b5f610b08836001600160a01b038416613377565b5f610b08836001600160a01b038416613166565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612bc5576040513d5f823e3d81fd5b50505f513d91508115612bdc578060011415612bea565b84516001600160a01b03163b155b15612c16578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610b57565b5050505050565b5f516020614b795f395f51905f525460ff16610cb557604051638dfc202b60e01b815260040160405180910390fd5b60cb546001600160a01b0390811690851603612d0257612c6d606483614a91565b349015612c90576040516308dc47c960e41b8152600401610b5791815260200190565b50612c9b81836145eb565b3414612ca782846145eb565b349091612cd057604051632711109560e11b815260048101929092526024820152604401610b57565b50508015612cfd57612cfd81612cee60cc546001600160a01b031690565b6001600160a01b0316906130da565b612de1565b8015612d3057612d3083612d1e60cc546001600160a01b031690565b6001600160a01b038716919084613436565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201849052851690639dc29fac906044016020604051808303815f875af1158015612d7c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612da091906149c1565b848484909192612ddd576040516336b52daf60e21b81526001600160a01b0393841660048201529290911660248301526044820152606401610b57565b5050505b50505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a604051612e429796959493929190614ab0565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051612e99929190918252602082015260400190565b60405180910390a450505050505050505050565b612eb682613490565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612efa5761229f82826134f3565b610bec613565565b5f825f018281548110612f1757612f17614560565b905f5260205f200154905092915050565b806001600160a01b038116612f68576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610b57565b612f725f83612af1565b8290612f9d576040516351eccfe160e11b81526001600160a01b039091166004820152602401610b57565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116613015576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610b57565b61301f5f83612b05565b829061304a576040516340da71e560e01b81526001600160a01b039091166004820152602401610b57565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611cc257602002820191905f5260205f20905b8154815260200190600101908083116130bb5750505050509050919050565b804710156131045760405163cf47918160e01b815247600482015260248101829052604401610b57565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461314e576040519150601f19603f3d011682016040523d82523d5f602084013e613153565b606091505b509150915081612de157612de181613584565b5f8181526001830160205260408120548015613240575f613188600183614af2565b85549091505f9061319b90600190614af2565b90508082146131fa575f865f0182815481106131b9576131b9614560565b905f5260205f200154905080875f0184815481106131d9576131d9614560565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061320b5761320b614b05565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610d7f565b5f915050610d7f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610cb557604051631afcd79f60e31b815260040160405180910390fd5b61329a613249565b6132a3816135ad565b6132eb604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506135be565b506099805460ff19166003179055565b610cb5613249565b61330b613249565b610cb56135d0565b61331b613249565b610cb56135f0565b5f610d7f61332f612a54565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f61335d86866135f8565b92509250925061336d8282613641565b5090949350505050565b5f8181526001830160205260408120546133bc57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610d7f565b505f610d7f565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6133ed6136f9565b6133f5613761565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612de19085906137a3565b806001600160a01b03163b5f036134c557604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610b57565b5f516020614b595f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161350f9190614b19565b5f60405180830381855af49150503d805f8114613547576040519150601f19603f3d011682016040523d82523d5f602084013e61354c565b606091505b509150915061355c85838361380f565b95945050505050565b3415610cb55760405163b398979f60e01b815260040160405180910390fd5b8051156135945780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6135b5613249565b610b1a8161386b565b6135c6613249565b610bec8282613873565b6135d8613249565b5f516020614b795f395f51905f52805460ff19169055565b611f94613249565b5f5f5f835160410361362f576020840151604085015160608601515f1a613621888285856138d2565b95509550955050505061363a565b505081515f91506002905b9250925092565b5f82600381111561365457613654614b24565b0361365d575050565b600182600381111561367157613671614b24565b0361368f5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156136a3576136a3614b24565b036136c45760405163fce698f760e01b815260048101829052602401610b57565b60038260038111156136d8576136d8614b24565b03610bec576040516335e2f38360e21b815260048101829052602401610b57565b5f5f516020614b395f395f51905f528161371161253d565b80519091501561372957805160209091012092915050565b81548015613738579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614b395f395f51905f52816137796125fd565b80519091501561379157805160209091012092915050565b60018201548015613738579392505050565b5f5f60205f8451602086015f885af1806137c2576040513d5f823e3d81fd5b50505f513d915081156137d95780600114156137e6565b6001600160a01b0384163b155b15612de157604051635274afe760e01b81526001600160a01b0385166004820152602401610b57565b6060826138245761381f82613584565b610b08565b815115801561383b57506001600160a01b0384163b155b1561386457604051639996b31560e01b81526001600160a01b0385166004820152602401610b57565b5080610b08565b611a9c613249565b61387b613249565b5f516020614b395f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026138b484826147d9565b50600381016138c383826147d9565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561390b57505f91506003905082613990565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561395c573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661398757505f925060019150829050613990565b92505f91508190505b9450945094915050565b5080546139a69061464a565b5f825580601f106139b5575050565b601f0160209004905f5260205f2090810190610b1a9190613a40565b5080545f8255905f5260205f2090810190610b1a9190613a54565b828054828255905f5260205f20908101928215613a30579160200282015b82811115613a305782518290613a2090826147d9565b5091602001919060010190613a0a565b50613a3c929150613a54565b5090565b5b80821115613a3c575f8155600101613a41565b80821115613a3c575f613a67828261399a565b50600101613a54565b634e487b7160e01b5f52600160045260245ffd5b5f5f83601f840112613a94575f5ffd5b5081356001600160401b03811115613aaa575f5ffd5b6020830191508360208260051b8501011115613ac4575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b0382111715613afe57613afe613acb565b60405250565b601f8201601f191681016001600160401b0381118282101715613b2957613b29613acb565b6040525050565b5f6001600160401b03821115613b4857613b48613acb565b5060051b60200190565b5f6001600160401b03821115613b6a57613b6a613acb565b50601f01601f191660200190565b5f613b8283613b52565b604051613b8f8282613b04565b809250848152858585011115613ba3575f5ffd5b848460208301375f6020868301015250509392505050565b5f82601f830112613bca575f5ffd5b610b0883833560208501613b78565b5f613be383613b30565b604051613bf08282613b04565b84815291505060208101600584901b830185811115613c0d575f5ffd5b835b81811015611c195780356001600160401b03811115613c2c575f5ffd5b613c3888828801613bbb565b84525060209283019201613c0f565b5f82601f830112613c56575f5ffd5b610b0883833560208501613bd9565b5f5f5f60408486031215613c77575f5ffd5b83356001600160401b03811115613c8c575f5ffd5b613c9886828701613a84565b90945092505060208401356001600160401b03811115613cb6575f5ffd5b8401601f81018613613cc6575f5ffd5b8035613cd181613b30565b604051613cde8282613b04565b80915082815260208101915060208360051b850101925088831115613d01575f5ffd5b602084015b83811015613d415780356001600160401b03811115613d23575f5ffd5b613d328b602083890101613c47565b84525060209283019201613d06565b50809450505050509250925092565b6001600160a01b0381168114610b1a575f5ffd5b5f60208284031215613d74575f5ffd5b8135610b0881613d50565b803560ff81168114611454575f5ffd5b5f60e08284031215613d9f575f5ffd5b604051613dab81613adf565b8091508235613db981613d50565b81526020830135613dc981613d50565b60208201526040838101359082015260608084013590820152613dee60808401613d7f565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613e27575f5ffd5b8935613e3281613d50565b985060208a0135613e4281613d50565b975060408a0135613e5281613d50565b965060608a0135955060808a0135945060a08a01359350613e768b60c08c01613d8f565b92506101a08a01356001600160401b03811115613e91575f5ffd5b613e9d8c828d01613a84565b915080935050809150509295985092959850929598565b5f60208284031215613ec4575f5ffd5b81356001600160401b03811115613ed9575f5ffd5b8201601f81018413613ee9575f5ffd5b8035613ef481613b30565b604051613f018282613b04565b80915082815260208101915060208360051b850101925086831115613f24575f5ffd5b6020840193505b82841015613f4f578335613f3e81613d50565b825260209384019390910190613f2b565b9695505050505050565b5f5f5f5f5f5f60a08789031215613f6e575f5ffd5b8635613f7981613d50565b955060208701359450604087013593506060870135925060808701356001600160401b03811115613fa8575f5ffd5b613fb489828a01613a84565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a031215613fdc575f5ffd5b8735613fe781613d50565b96506020880135613ff781613d50565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115614026575f5ffd5b6140328a828b01613a84565b989b979a50959850939692959293505050565b5f5f60408385031215614056575f5ffd5b823561406181613d50565b915060208301356001600160401b0381111561407b575f5ffd5b61408785828601613bbb565b9150509250929050565b5f602082840312156140a1575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b0312156140c0575f5ffd5b88356140cb81613d50565b975060208901356140db81613d50565b96506040890135955060608901359450608089013593506140ff8a60a08b01613d8f565b92506101808901356001600160401b0381111561411a575f5ffd5b6141268b828c01613a84565b999c989b5096995094979396929594505050565b5f5f6040838503121561414b575f5ffd5b823561415681613d50565b9150602083013561416681613d50565b809150509250929050565b5f5f60408385031215614182575f5ffd5b823561418d81613d50565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156141db5783516001600160a01b03168352602093840193909201916001016141b4565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156142ad5760df198786030183526142988585516141e6565b9450602093840193929092019160010161427c565b50929695505050505050565b60ff60f81b8816815260e060208201525f6142d760e08301896141e6565b82810360408401526142e981896141e6565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561433e578351835260209384019390920191600101614320565b50909b9a5050505050505050505050565b602081525f610b0860208301846141e6565b5f60208284031215614371575f5ffd5b610b0882613d7f565b5f5f5f6060848603121561438c575f5ffd5b833561439781613d50565b925060208401356143a781613d50565b915060408401356143b781613d50565b809150509250925092565b5f602082840312156143d2575f5ffd5b81356001600160401b038111156143e7575f5ffd5b8201601f810184136143f7575f5ffd5b803561440281613b30565b60405161440f8282613b04565b80915082815260208101915060208360051b850101925086831115614432575f5ffd5b6020840193505b82841015613f4f578335825260209384019390910190614439565b5f5f5f5f5f5f5f60c0888a03121561446a575f5ffd5b87359650602088013561447c81613d50565b9550604088013561448c81613d50565b94506060880135935060808801356001600160401b038111156144ad575f5ffd5b6144b98a828b01613a84565b90945092505060a08801356001600160401b038111156144d7575f5ffd5b6144e38a828b01613c47565b91505092959891949750929550565b5f5f5f60608486031215614504575f5ffd5b833561450f81613d50565b925060208401356001600160401b03811115614529575f5ffd5b8401601f81018613614539575f5ffd5b61454886823560208401613b78565b92505061455760408501613d7f565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e19833603018112614588575f5ffd5b9190910192915050565b5f5f8335601e198436030181126145a7575f5ffd5b8301803591506001600160401b038211156145c0575f5ffd5b6020019150600581901b3603821315613ac4575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d7f57610d7f6145d7565b80518015158114611454575f5ffd5b5f5f5f5f60808587031215614620575f5ffd5b8451602086015160408701519195509350915061463f606086016145fe565b905092959194509250565b600181811c9082168061465e57607f821691505b60208210810361467c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215614692575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b8681101561474957838303601f19018852813536879003601e190181126146fd575f5ffd5b86016020810190356001600160401b03811115614718575f5ffd5b803603821315614726575f5ffd5b614731858284614699565b60209a8b019a909550939093019250506001016146d8565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90610d9790830184866146c1565b601f82111561229f57805f5260205f20601f840160051c810160208510156147ba5750805b601f840160051c820191505b81811015612c16575f81556001016147c6565b81516001600160401b038111156147f2576147f2613acb565b61480681614800845461464a565b84614795565b6020601f821160018114614838575f83156148215750848201515b5f19600385901b1c1916600184901b178455612c16565b5f84815260208120601f198516915b828110156148675787850151825560209485019460019092019101614847565b508482101561488457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610b08368484613bd9565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f610b08600d83018461489f565b5f602082840312156148e7575f5ffd5b81516001600160401b038111156148fc575f5ffd5b8201601f8101841361490c575f5ffd5b805161491781613b52565b6040516149248282613b04565b828152866020848601011115614938575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b606081525f61496860608301866141e6565b828103602084015261497a81866141e6565b91505060ff83166040830152949350505050565b5f6149a261499c838661489f565b8461489f565b949350505050565b8082028115828204841417610d7f57610d7f6145d7565b5f602082840312156149d1575f5ffd5b610b08826145fe565b5f60033d1115610dbd5760045f5f3e505f5160e01c90565b5f60443d10156149ff5790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614a2857505090565b80820180516001600160401b03811115614a43575050505090565b3d8401600319018282016020011115614a5d575050505090565b614a6c60208285010185613b04565b509392505050565b5f5f60233d1115614a8d57602060045f3e50505f516001905b9091565b5f82614aab57634e487b7160e01b5f52601260045260245ffd5b500690565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f90610d9790830184866146c1565b81810381811115610d7f57610d7f6145d7565b634e487b7160e01b5f52603160045260245ffd5b5f610b08828461489f565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220ff147271b934a51fd14242153ac98c2f7bf75aedb7408e30a63c7b701cf2f79564736f6c634300081c0033",
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

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeCross *BridgeCrossCaller) BridgeTokenInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "bridgeTokenInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeCross *BridgeCrossSession) BridgeTokenInfo() (common.Address, error) {
	return _BridgeCross.Contract.BridgeTokenInfo(&_BridgeCross.CallOpts)
}

// BridgeTokenInfo is a free data retrieval call binding the contract method 0x9c1b65a9.
//
// Solidity: function bridgeTokenInfo() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) BridgeTokenInfo() (common.Address, error) {
	return _BridgeCross.Contract.BridgeTokenInfo(&_BridgeCross.CallOpts)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeCross *BridgeCrossCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "calculate", token, value)

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
func (_BridgeCross *BridgeCrossSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeCross.Contract.Calculate(&_BridgeCross.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeCross *BridgeCrossCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeCross.Contract.Calculate(&_BridgeCross.CallOpts, token, value)
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

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BridgeCross *BridgeCrossCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BridgeCross *BridgeCrossSession) Weth() (common.Address, error) {
	return _BridgeCross.Contract.Weth(&_BridgeCross.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BridgeCross *BridgeCrossCallerSession) Weth() (common.Address, error) {
	return _BridgeCross.Contract.Weth(&_BridgeCross.CallOpts)
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
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) Bridge(opts *bind.TransactOpts, token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "bridge", token, value, gas, ex, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) Bridge(token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Bridge(&_BridgeCross.TransactOpts, token, value, gas, ex, extraData)
}

// Bridge is a paid mutator transaction binding the contract method 0x37d10075.
//
// Solidity: function bridge(address token, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) Bridge(token common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Bridge(&_BridgeCross.TransactOpts, token, value, gas, ex, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) BridgeTo(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "bridgeTo", token, to, value, gas, ex, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeTo(&_BridgeCross.TransactOpts, token, to, value, gas, ex, extraData)
}

// BridgeTo is a paid mutator transaction binding the contract method 0x3f5d3b5d.
//
// Solidity: function bridgeTo(address token, address to, uint256 value, uint256 gas, uint256 ex, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) BridgeTo(token common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.BridgeTo(&_BridgeCross.TransactOpts, token, to, value, gas, ex, extraData)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeCross *BridgeCrossTransactor) ChangeRewardWallet(opts *bind.TransactOpts, rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "changeRewardWallet", rewardWallet_)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeCross *BridgeCrossSession) ChangeRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.ChangeRewardWallet(&_BridgeCross.TransactOpts, rewardWallet_)
}

// ChangeRewardWallet is a paid mutator transaction binding the contract method 0xabe110b6.
//
// Solidity: function changeRewardWallet(address rewardWallet_) returns()
func (_BridgeCross *BridgeCrossTransactorSession) ChangeRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.ChangeRewardWallet(&_BridgeCross.TransactOpts, rewardWallet_)
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
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactor) Initialize(opts *bind.TransactOpts, crossMintableERC20Code common.Address, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "initialize", crossMintableERC20Code, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossSession) Initialize(crossMintableERC20Code common.Address, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address crossMintableERC20Code, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactorSession) Initialize(crossMintableERC20Code common.Address, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, rewardWallet_, _bridgeTokenInfo)
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
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridge(opts *bind.TransactOpts, token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridge", token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridge is a paid mutator transaction binding the contract method 0x51c45579.
//
// Solidity: function permitBridge(address token, address account, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridge(token common.Address, account common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridge(&_BridgeCross.TransactOpts, token, account, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) PermitBridgeTo(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "permitBridgeTo", token, from, to, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, ex, permitArgs, extraData)
}

// PermitBridgeTo is a paid mutator transaction binding the contract method 0x174991ab.
//
// Solidity: function permitBridgeTo(address token, address from, address to, uint256 value, uint256 gas, uint256 ex, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes[] extraData) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) PermitBridgeTo(token common.Address, from common.Address, to common.Address, value *big.Int, gas *big.Int, ex *big.Int, permitArgs IBridgeStandardPermitArguments, extraData [][]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.PermitBridgeTo(&_BridgeCross.TransactOpts, token, from, to, value, gas, ex, permitArgs, extraData)
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

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeCross *BridgeCrossTransactor) RemoveTokenInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "removeTokenInfo")
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeCross *BridgeCrossSession) RemoveTokenInfo() (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveTokenInfo(&_BridgeCross.TransactOpts)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x85547884.
//
// Solidity: function removeTokenInfo() returns()
func (_BridgeCross *BridgeCrossTransactorSession) RemoveTokenInfo() (*types.Transaction, error) {
	return _BridgeCross.Contract.RemoveTokenInfo(&_BridgeCross.TransactOpts)
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

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeCross *BridgeCrossTransactor) RetryFinalizeBatch(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "retryFinalizeBatch", indexes)
}

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeCross *BridgeCrossSession) RetryFinalizeBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeCross.Contract.RetryFinalizeBatch(&_BridgeCross.TransactOpts, indexes)
}

// RetryFinalizeBatch is a paid mutator transaction binding the contract method 0xc1ad8b95.
//
// Solidity: function retryFinalizeBatch(uint256[] indexes) returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) RetryFinalizeBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _BridgeCross.Contract.RetryFinalizeBatch(&_BridgeCross.TransactOpts, indexes)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactor) SetTokenInfo(opts *bind.TransactOpts, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "setTokenInfo", _bridgeTokenInfo)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossSession) SetTokenInfo(_bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetTokenInfo(&_BridgeCross.TransactOpts, _bridgeTokenInfo)
}

// SetTokenInfo is a paid mutator transaction binding the contract method 0x2f9b59d1.
//
// Solidity: function setTokenInfo(address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactorSession) SetTokenInfo(_bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.SetTokenInfo(&_BridgeCross.TransactOpts, _bridgeTokenInfo)
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
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
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
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
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
