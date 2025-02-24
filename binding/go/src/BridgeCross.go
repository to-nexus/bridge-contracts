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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addTokenDeploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenInfo\",\"outputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"changeRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossMintableERC20Code\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
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
		"3fc8cef3": "weth()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614e746100eb5f395f6126cb0152614e745ff3fe60806040526004361061035e575f3560e01c80637c41ad2c116101bd578063b7f3358d116100f2578063f2fde38b11610092578063f7684ec41161006d578063f7684ec4146109a3578063facd743b146109c2578063fb75b2c7146109e1578063fe2b8da6146109fe575f5ffd5b8063f2fde38b1461095c578063f30589c31461097b578063f698da251461098f575f5ffd5b8063c97682f8116100cd578063c97682f814610901578063cbae595814610915578063d92fc67b14610934578063e1af7f5014610948575f5ffd5b8063b7f3358d146108a4578063c1876453146108c3578063c1ad8b95146108e2575f5ffd5b806396ce07951161015d578063a982318311610138578063a982318314610815578063abe110b614610834578063ad3cb1cc14610853578063aed1d40314610890575f5ffd5b806396ce0795146107c35780639c1b65a9146107d75780639c2c58f8146107f6575f5ffd5b8063855478841161019857806385547884146107405780638da5cb5b1461075457806391cf6d3e146107905780639300c926146107a4575f5ffd5b80637c41ad2c146106e65780638456cb591461070557806384b0196e14610719575f5ffd5b80634f1ef286116102935780635fa7b584116102335780637021fd0e1161020e5780637021fd0e146106505780637101fcd31461067c578063715018a61461069b57806371c59d7b146106af575f5ffd5b80635fa7b584146105cf5780636e908ca3146105ee5780636ff97f1d1461062f575f5ffd5b806352d1902d1161026e57806352d1902d1461055a5780635476bd721461056e5780635c975abb1461058d5780635dbe47e8146105b0575f5ffd5b80634f1ef286146105155780634f6ccce71461052857806351c4557914610547575f5ffd5b80633b3bff0f116102fe5780633fc8cef3116102d95780633fc8cef31461048b57806340a141ff146104c257806342cde4e8146104e157806349a495ec14610502575f5ffd5b80633b3bff0f146104455780633f4ba83a146104645780633f5d3b5d14610478575f5ffd5b806327938c811161033957806327938c81146103de5780632f63b7c8146104005780632f9b59d11461041357806337d1007514610432575f5ffd5b80631327d3d814610378578063174991ab146103975780631d40f0d8146103bf575f5ffd5b3661037457345f0361037257610372613b0b565b005b5f5ffd5b348015610383575f5ffd5b50610372610392366004613b3e565b610a1d565b6103aa6103a5366004613c94565b610a2b565b60405190151581526020015b60405180910390f35b3480156103ca575f5ffd5b506103726103d9366004613d5c565b610ac4565b3480156103e9575f5ffd5b506103f2610afe565b6040519081526020016103b6565b6103aa61040e366004613ee0565b610b13565b34801561041e575f5ffd5b5061037261042d366004613b3e565b610db7565b6103aa610440366004613fcc565b610e2b565b348015610450575f5ffd5b5061037261045f366004613b3e565b610e46565b34801561046f575f5ffd5b50610372610e57565b6103aa610486366004614039565b610e69565b348015610496575f5ffd5b5060fc546104aa906001600160a01b031681565b6040516001600160a01b0390911681526020016103b6565b3480156104cd575f5ffd5b506103726104dc366004613b3e565b610f01565b3480156104ec575f5ffd5b5060995460405160ff90911681526020016103b6565b6103aa610510366004614142565b610f0b565b610372610523366004614305565b61106b565b348015610533575f5ffd5b506104aa610542366004614351565b611086565b6103aa610555366004614368565b611097565b348015610565575f5ffd5b506103f26110b6565b348015610579575f5ffd5b506103726105883660046143fa565b6110d2565b348015610598575f5ffd5b505f516020614dff5f395f51905f525460ff166103aa565b3480156105bb575f5ffd5b506103aa6105ca366004613b3e565b6110e4565b3480156105da575f5ffd5b506103726105e9366004613b3e565b6110ef565b3480156105f9575f5ffd5b5061060d610608366004614431565b611100565b60408051948552602085019390935291830152151560608201526080016103b6565b34801561063a575f5ffd5b50610643611189565b6040516103b6919061445b565b34801561065b575f5ffd5b5061066f61066a366004614351565b6111ea565b6040516103b691906144d4565b348015610687575f5ffd5b50610372610696366004613d5c565b61134a565b3480156106a6575f5ffd5b50610372611360565b3480156106ba575f5ffd5b506104aa6106c9366004613b3e565b6001600160a01b039081165f908152603360205260409020541690565b3480156106f1575f5ffd5b50610372610700366004613b3e565b611371565b348015610710575f5ffd5b50610372611382565b348015610724575f5ffd5b5061072d611392565b6040516103b69796959493929190614579565b34801561074b575f5ffd5b5061037261143b565b34801561075f575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166104aa565b34801561079b575f5ffd5b5060cc546103f2565b3480156107af575f5ffd5b506103726107be366004613d5c565b611455565b3480156107ce575f5ffd5b506103f261148c565b3480156107e2575f5ffd5b5060ca546104aa906001600160a01b031681565b348015610801575f5ffd5b5061037261081036600461460f565b6114f7565b348015610820575f5ffd5b506103aa61082f366004614351565b61166b565b34801561083f575f5ffd5b5061037261084e366004613b3e565b6118f5565b34801561085e575f5ffd5b50610883604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516103b69190614666565b34801561089b575f5ffd5b506103f261191f565b3480156108af575f5ffd5b506103726108be366004614678565b61192a565b3480156108ce575f5ffd5b506103aa6108dd366004613b3e565b61197b565b3480156108ed575f5ffd5b506103aa6108fc366004614691565b6119ab565b34801561090c575f5ffd5b506106436119eb565b348015610920575f5ffd5b506104aa61092f366004614351565b611ad1565b34801561093f575f5ffd5b506103f2611add565b348015610953575f5ffd5b506103f2611ae7565b348015610967575f5ffd5b50610372610976366004613b3e565b611af7565b348015610986575f5ffd5b50610643611b31565b34801561099a575f5ffd5b506103f2611b3d565b3480156109ae575f5ffd5b506104aa6109bd366004614723565b611b46565b3480156109cd575f5ffd5b506103aa6109dc366004613b3e565b611c86565b3480156109ec575f5ffd5b5060cb546001600160a01b03166104aa565b348015610a09575f5ffd5b50610883610a18366004614351565b611c92565b610a28816001611d31565b50565b5f610a34611dfa565b89610a3e8161197b565b8190610a6e57604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610a77611e2a565b5f5f610a858d8b8b8b611e61565b91509150610a9a8d8d8d8d86868d8d8d611f81565b600193505050610ab660015f516020614e1f5f395f51905f5255565b509998505050505050505050565b5f5b8151811015610afa57610af2828281518110610ae457610ae4614791565b60200260200101515f611d31565b600101610ac6565b5050565b5f6066546001610b0e91906147b9565b905090565b5f610b1c611dfa565b88610b268161197b565b8190610b5157604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610a65565b50610b5a611e2a565b610b62610afe565b8b14610b6c610afe565b8c9091610b9557604051634982205b60e11b815260048101929092526024820152604401610a65565b5050610bf57fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8c8c8c8c8c8c604051602001610bd7979695949392919061487c565b60405160208183030381529060405280519060200120868686611fff565b610c03606680546001019055565b5f5f610c108c8c8c6121fa565b915091508115610c76578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051610c69929190918252602082015260400190565b60405180910390a4610a9a565b610c8160cf8e6123e4565b8d90610ca3576040516368db995b60e11b8152600401610a6591815260200190565b505f8d815260ce60205260409020610cbb8282614938565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90610cfe91906149f2565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051610d6e9260048501920190613a39565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a2600193505050610ab660015f516020614e1f5f395f51905f5255565b610dbf6123f6565b6001600160a01b038116610e0957604051633effd40360e21b815260206004820152601060248201526f5f627269646765546f6b656e496e666f60801b6044820152606401610a65565b60ca80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610e3b87338888888888610e69565b979650505050505050565b610e4e6123f6565b610a2881612451565b610e5f6123f6565b610e676124f1565b565b5f610e72611dfa565b87610e7c8161197b565b8190610ea757604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610a65565b50610eb0611e2a565b5f5f610ebe8b8a8a8a611e61565b91509150610ed98b610ecd3390565b8c8c86865f8d8d61254a565b600193505050610ef560015f516020614e1f5f395f51905f5255565b50979650505050505050565b610a28815f611d31565b5f805b8581101561105e57611055878783818110610f2b57610f2b614791565b9050602002810190610f3d9190614a5f565b35888884818110610f5057610f50614791565b9050602002810190610f629190614a5f565b610f73906040810190602001613b3e565b898985818110610f8557610f85614791565b9050602002810190610f979190614a5f565b610fa8906060810190604001613b3e565b8a8a86818110610fba57610fba614791565b9050602002810190610fcc9190614a5f565b606001358b8b87818110610fe257610fe2614791565b9050602002810190610ff49190614a5f565b611002906080810190614a7d565b8b888151811061101457611014614791565b60200260200101518b898151811061102e5761102e614791565b60200260200101518b8a8151811061104857611048614791565b6020026020010151610b13565b50600101610f0e565b5060019695505050505050565b611073612586565b61107c826125ec565b610afa82826125f4565b5f61109181836126b5565b92915050565b5f6110a989898a8a8a8a8a8a8a610a2b565b9998505050505050505050565b5f6110bf6126c0565b505f516020614ddf5f395f51905f525b90565b6110da6123f6565b610afa8282612709565b5f6110918183612778565b6110f76123f6565b610a2881612799565b60ca54604051636e908ca360e01b81526001600160a01b038481166004830152602482018490525f928392839283921690636e908ca390604401608060405180830381865afa158015611155573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111799190614ad1565b9299919850965090945092505050565b60605f6111955f612801565b9050806001600160401b038111156111af576111af613b59565b6040519080825280602002602001820160405280156111d8578160200160208202803683370190505b5091506111e45f61280a565b91505090565b61122a6040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b8282101561133c578382905f5260205f200180546112b1906148bc565b80601f01602080910402602001604051908101604052809291908181526020018280546112dd906148bc565b80156113285780601f106112ff57610100808354040283529160200191611328565b820191905f5260205f20905b81548152906001019060200180831161130b57829003601f168201915b505050505081526020019060010190611294565b505050915250909392505050565b6113576103d9609761280a565b610a2881611455565b6113686123f6565b610e675f612816565b6113796123f6565b610a2881612886565b61138a6123f6565b610e67612906565b5f60608082808083815f516020614dbf5f395f51905f5280549091501580156113bd57506001810154155b6114015760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a65565b61140961294e565b611411612a0e565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6114436123f6565b60ca80546001600160a01b0319169055565b5f5b8151811015610afa5761148482828151811061147557611475614791565b60200260200101516001611d31565b600101611457565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa1580156114d3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b0e9190614b0e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561153b5750825b90505f826001600160401b031660011480156115565750303b155b905081158015611564575080155b156115825760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156115ac57845460ff60401b1916600160401b1785555b6115b7888888612a4c565b60fd80546001600160a01b0319166001600160a01b038b1617905560408051808201909152600381526208aa8960eb1b60208201526115fa906001906012611b46565b60fc80546001600160a01b0319166001600160a01b0392909216919091179055831561166057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f611674611dfa565b61167c611e2a565b61168760cf83612b0f565b82906116a95760405163473978bf60e01b8152600401610a6591815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156117bc578382905f5260205f20018054611731906148bc565b80601f016020809104026020016040519081016040528092919081815260200182805461175d906148bc565b80156117a85780601f1061177f576101008083540402835291602001916117a8565b820191905f5260205f20905b81548152906001019060200180831161178b57829003601f168201915b505050505081526020019060010190611714565b505050508152505090505f5f6117df8360200151846040015185606001516121fa565b915091508181906118035760405162461bcd60e51b8152600401610a659190614666565b5061180f60cf86612b26565b505f85815260ce6020526040812061182691613a8d565b5f85815260cd602052604081208181556001810180546001600160a01b03199081169091556002820180549091169055600381018290559061186b6004830182613ac4565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738660600151426040516118cb929190918252602082015260400190565b60405180910390a4600193505050506118f060015f516020614e1f5f395f51905f5255565b919050565b6118fd6123f6565b60cb80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610b0e6097612801565b6119326123f6565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f611985826110e4565b80156110915750506001600160a01b03165f9081526034602052604090205460ff161590565b5f805b82518110156119e2576119d98382815181106119cc576119cc614791565b602002602001015161166b565b506001016119ae565b50600192915050565b60605f6119f6611189565b90505f81516001600160401b03811115611a1257611a12613b59565b604051908082528060200260200182016040528015611a3b578160200160208202803683370190505b5090505f5b8251811015611aca5760335f848381518110611a5e57611a5e614791565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b0316828281518110611aaa57611aaa614791565b6001600160a01b0390921660209283029190910190910152600101611a40565b5092915050565b5f6110916097836126b5565b5f610b0e5f612801565b5f6065546001610b0e91906147b9565b611aff6123f6565b6001600160a01b038116611b2857604051631e4fbdf760e01b81525f6004820152602401610a65565b610a2881612816565b6060610b0e609761280a565b5f610b0e612b31565b5f611b4f6123f6565b5f83604051602001611b619190614b3c565b60408051601f19818403018152908290526bffffffffffffffffffffffff19606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f60fd5f9054906101000a90046001600160a01b03166001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611bfa573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611c219190810190614b5d565b838787604051602001611c3693929190614bdc565b60408051601f1981840301815290829052611c549291602001614c14565b6040516020818303038152906040529050611c705f8383612b3a565b9350611c7c84886110d2565b5050509392505050565b5f611091609783612778565b5f81815260ce60205260409020805460609190611cae906148bc565b80601f0160208091040260200160405190810160405280929190818152602001828054611cda906148bc565b8015611d255780601f10611cfc57610100808354040283529160200191611d25565b820191905f5260205f20905b815481529060010190602001808311611d0857829003601f168201915b50505050509050919050565b611d396123f6565b8015611d7b57611d4a609783612bce565b8290611d755760405163082cdf5560e21b81526001600160a01b039091166004820152602401610a65565b50611db3565b611d86609783612be2565b8290611db15760405163e491024560e01b81526001600160a01b039091166004820152602401610a65565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614dff5f395f51905f525460ff1615610e675760405163d93c066560e01b815260040160405180910390fd5b5f516020614e1f5f395f51905f52805460011901611e5b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f9081906001600160a01b0316611e7f57505f905080611f78565b60ca54604051636e908ca360e01b81526001600160a01b038881166004830152602482018890525f928392911690636e908ca390604401608060405180830381865afa158015611ed1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ef59190614ad1565b919650945092509050808710801590611f0e5750838610155b8015611f1a5750828510155b8015611f235750815b8185858a8a8a909192939495611f6f576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a65565b50505050505050505b94509492505050565b83611f8c86886147b9565b611f9691906147b9565b8360400151101589879091611fcf57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610a65565b5050611fda83612bf6565b6116608989898989896001898961254a565b60015f516020614e1f5f395f51905f5255565b82518251811480156120115750815181145b835183518392612045576040516313a27e1960e21b8152600481019390935260248301919091526044820152606401610a65565b505060995482915060ff168110156120735760405163fedb0cbd60e01b8152600401610a6591815260200190565b505f80826001600160401b0381111561208e5761208e613b59565b6040519080825280602002602001820160405280156120b7578160200160208202803683370190505b5090505f5b838110156121c4575f6121278883815181106120da576120da614791565b60200260200101518884815181106120f4576120f4614791565b602002602001015188858151811061210e5761210e614791565b602002602001015161211f8d612cfa565b929190612d26565b905061213281611c86565b819061215d576040516338905e7160e11b81526001600160a01b039091166004820152602401610a65565b505f805b84518110156121ad5784818151811061217c5761217c614791565b60200260200101516001600160a01b0316836001600160a01b0316036121a557600191506121ad565b600101612161565b50806121ba578460010194505b50506001016120bc565b50609954829060ff168110156121f05760405163fedb0cbd60e01b8152600401610a6591815260200190565b5050505050505050565b5f6060825f0361221c57505060408051602081019091525f81526001906123dc565b5f196001600160a01b038616016122635761224b61223b606485614c30565b6001600160a01b03861690612d52565b505060408051602081019091525f81526001906123dc565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528616906340c10f19906044016020604051808303815f875af19250505080156122cd575060408051601f3d908101601f191682019092526122ca91810190614c47565b60015b61237e576122d9614c60565b806308c379a00361230257506122ed614c78565b806122f85750612345565b5f925090506123dc565b634e487b710361234557612314614cfa565b9061231f5750612345565b60408051602081018390525f9450016040516020818303038152906040529150506123dc565b3d80801561236e576040519150601f19603f3d011682016040523d82523d5f602084013e612373565b606091505b505f925090506123dc565b801561239e576001925060405180602001604052805f81525091506123da565b5f92506040518060400160405280601881526020017f42726964676543726f73733a206d696e74206661696c6564000000000000000081525091505b505b935093915050565b5f6123ef8383612de4565b9392505050565b336124287f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610e675760405163118cdaa760e01b8152336004820152602401610a65565b61245a816110e4565b801561247d57506001600160a01b0381165f9081526034602052604090205460ff165b81906124a8576040516340da71e560e01b81526001600160a01b039091166004820152602401610a65565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b6124f9612e30565b5f516020614dff5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001611970565b61255f89898861255a888a6147b9565b612e5f565b61257861256a611ae7565b8a8a8a8a8a8a8a8a8a612ff4565b611660606580546001019055565b30610101620956d760881b0114806125ce5750610101620956d760881b016125c25f516020614ddf5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e675760405163703e46dd60e11b815260040160405180910390fd5b610a286123f6565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561264e575060408051601f3d908101601f1916820190925261264b91810190614b0e565b60015b61267657604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a65565b5f516020614ddf5f395f51905f5281146126a657604051632a87526960e21b815260048101829052602401610a65565b6126b083836130ba565b505050565b5f6123ef838361310f565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e675760405163703e46dd60e11b815260040160405180910390fd5b61271282613135565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f90815260018301602052604081205415156123ef565b6127a2816131e2565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce910161276c565b5f611091825490565b60605f6123ef8361328f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61288f8161197b565b81906128ba576040516340da71e560e01b81526001600160a01b039091166004820152602401610a65565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b61290e611dfa565b5f516020614dff5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612532565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614dbf5f395f51905f529161298c906148bc565b80601f01602080910402602001604051908101604052809291908181526020018280546129b8906148bc565b8015612a035780601f106129da57610100808354040283529160200191612a03565b820191905f5260205f20905b8154815290600101906020018083116129e657829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614dbf5f395f51905f529161298c906148bc565b612a546132e7565b612a5e3384613330565b612a666133a0565b612a6e6133a8565b612a766133b8565b6001600160a01b038216612abc57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a65565b4360cc556001600160a01b03811615612aeb5760ca80546001600160a01b0319166001600160a01b0383161790555b5060cb80546001600160a01b0319166001600160a01b039290921691909117905550565b5f81815260018301602052604081205415156123ef565b5f6123ef83836133c8565b5f610b0e6134ab565b5f83471015612b655760405163cf47918160e01b815247600482015260248101859052604401610a65565b81515f03612b8657604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615612ba7576040513d5f823e3d81fd5b6001600160a01b0381166123ef5760405163b06ebf3d60e01b815260040160405180910390fd5b5f6123ef836001600160a01b038416612de4565b5f6123ef836001600160a01b0384166133c8565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612ca2576040513d5f823e3d81fd5b50505f513d91508115612cb9578060011415612cc7565b84516001600160a01b03163b155b15612cf3578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610a65565b5050505050565b5f611091612d06612b31565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612d368888888861351e565b925092509250612d4682826135e6565b50909695505050505050565b80471015612d7c5760405163cf47918160e01b815247600482015260248101829052604401610a65565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114612dc6576040519150601f19603f3d011682016040523d82523d5f602084013e612dcb565b606091505b509150915081612dde57612dde8161369e565b50505050565b5f818152600183016020526040812054612e2957508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611091565b505f611091565b5f516020614dff5f395f51905f525460ff16610e6757604051638dfc202b60e01b815260040160405180910390fd5b5f196001600160a01b03851601612f1057612e7b606483614d17565b349015612e9e576040516308dc47c960e41b8152600401610a6591815260200190565b50612ea981836147b9565b3414612eb582846147b9565b349091612ede57604051632711109560e11b815260048101929092526024820152604401610a65565b50508015612f0b57612f0b81612efc60cb546001600160a01b031690565b6001600160a01b031690612d52565b612dde565b8015612f3e57612f3e83612f2c60cb546001600160a01b031690565b6001600160a01b0387169190846136c7565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201849052851690639dc29fac906044016020604051808303815f875af1158015612f8a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fae9190614c47565b848484909192612feb576040516336b52daf60e21b81526001600160a01b0393841660048201529290911660248301526044820152606401610a65565b50505050505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a60405161304f9796959493929190614d36565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d88886040516130a6929190918252602082015260400190565b60405180910390a450505050505050505050565b6130c382613721565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613107576126b08282613784565b610afa6137f6565b5f825f01828154811061312457613124614791565b905f5260205f200154905092915050565b806001600160a01b038116613175576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a65565b61317f5f83612bce565b82906131aa576040516351eccfe160e11b81526001600160a01b039091166004820152602401610a65565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116613222576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a65565b61322c5f83612be2565b8290613257576040516340da71e560e01b81526001600160a01b039091166004820152602401610a65565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611d2557602002820191905f5260205f20905b8154815260200190600101908083116132c85750505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e6757604051631afcd79f60e31b815260040160405180910390fd5b6133386132e7565b61334182613815565b613389604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613826565b6099805460ff191660ff9290921691909117905550565b610e676132e7565b6133b06132e7565b610e67613838565b6133c06132e7565b610e67613858565b5f81815260018301602052604081205480156134a2575f6133ea600183614d78565b85549091505f906133fd90600190614d78565b905080821461345c575f865f01828154811061341b5761341b614791565b905f5260205f200154905080875f01848154811061343b5761343b614791565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061346d5761346d614d8b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611091565b5f915050611091565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6134d5613860565b6134dd6138c8565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561355757505f915060039050826135dc565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156135a8573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166135d357505f9250600191508290506135dc565b92505f91508190505b9450945094915050565b5f8260038111156135f9576135f9614d9f565b03613602575050565b600182600381111561361657613616614d9f565b036136345760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561364857613648614d9f565b036136695760405163fce698f760e01b815260048101829052602401610a65565b600382600381111561367d5761367d614d9f565b03610afa576040516335e2f38360e21b815260048101829052602401610a65565b8051156136ae5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612dde90859061390a565b806001600160a01b03163b5f0361375657604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a65565b5f516020614ddf5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516137a09190614db3565b5f60405180830381855af49150503d805f81146137d8576040519150601f19603f3d011682016040523d82523d5f602084013e6137dd565b606091505b50915091506137ed858383613976565b95945050505050565b3415610e675760405163b398979f60e01b815260040160405180910390fd5b61381d6132e7565b610a28816139d2565b61382e6132e7565b610afa82826139da565b6138406132e7565b5f516020614dff5f395f51905f52805460ff19169055565b611fec6132e7565b5f5f516020614dbf5f395f51905f528161387861294e565b80519091501561389057805160209091012092915050565b8154801561389f579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614dbf5f395f51905f52816138e0612a0e565b8051909150156138f857805160209091012092915050565b6001820154801561389f579392505050565b5f5f60205f8451602086015f885af180613929576040513d5f823e3d81fd5b50505f513d9150811561394057806001141561394d565b6001600160a01b0384163b155b15612dde57604051635274afe760e01b81526001600160a01b0385166004820152602401610a65565b60608261398b576139868261369e565b6123ef565b81511580156139a257506001600160a01b0384163b155b156139cb57604051639996b31560e01b81526001600160a01b0385166004820152602401610a65565b50806123ef565b611aff6132e7565b6139e26132e7565b5f516020614dbf5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613a1b8482614938565b5060038101613a2a8382614938565b505f8082556001909101555050565b828054828255905f5260205f20908101928215613a7d579160200282015b82811115613a7d5782518290613a6d9082614938565b5091602001919060010190613a57565b50613a89929150613adb565b5090565b508054613a99906148bc565b5f825580601f10613aa8575050565b601f0160209004905f5260205f2090810190610a289190613af7565b5080545f8255905f5260205f2090810190610a2891905b80821115613a89575f613aee8282613a8d565b50600101613adb565b5b80821115613a89575f8155600101613af8565b634e487b7160e01b5f52600160045260245ffd5b6001600160a01b0381168114610a28575f5ffd5b80356118f081613b1f565b5f60208284031215613b4e575f5ffd5b81356123ef81613b1f565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b0382111715613b8c57613b8c613b59565b60405250565b601f8201601f191681016001600160401b0381118282101715613bb757613bb7613b59565b6040525050565b803560ff811681146118f0575f5ffd5b5f60e08284031215613bde575f5ffd5b604051613bea81613b6d565b8091508235613bf881613b1f565b81526020830135613c0881613b1f565b60208201526040838101359082015260608084013590820152613c2d60808401613bbe565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f83601f840112613c5d575f5ffd5b5081356001600160401b03811115613c73575f5ffd5b6020830191508360208260051b8501011115613c8d575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f6101c08a8c031215613cad575f5ffd5b8935613cb881613b1f565b985060208a0135613cc881613b1f565b975060408a0135613cd881613b1f565b965060608a0135955060808a0135945060a08a01359350613cfc8b60c08c01613bce565b92506101a08a01356001600160401b03811115613d17575f5ffd5b613d238c828d01613c4d565b915080935050809150509295985092959850929598565b5f6001600160401b03821115613d5257613d52613b59565b5060051b60200190565b5f60208284031215613d6c575f5ffd5b81356001600160401b03811115613d81575f5ffd5b8201601f81018413613d91575f5ffd5b8035613d9c81613d3a565b604051613da98282613b92565b80915082815260208101915060208360051b850101925086831115613dcc575f5ffd5b6020840193505b82841015613df7578335613de681613b1f565b825260209384019390910190613dd3565b9695505050505050565b5f82601f830112613e10575f5ffd5b8135613e1b81613d3a565b604051613e288282613b92565b80915082815260208101915060208360051b860101925085831115613e4b575f5ffd5b602085015b83811015613e6f57613e6181613bbe565b835260209283019201613e50565b5095945050505050565b5f82601f830112613e88575f5ffd5b8135613e9381613d3a565b604051613ea08282613b92565b80915082815260208101915060208360051b860101925085831115613ec3575f5ffd5b602085015b83811015613e6f578035835260209283019201613ec8565b5f5f5f5f5f5f5f5f5f6101008a8c031215613ef9575f5ffd5b89359850613f0960208b01613b33565b9750613f1760408b01613b33565b965060608a0135955060808a01356001600160401b03811115613f38575f5ffd5b613f448c828d01613c4d565b90965094505060a08a01356001600160401b03811115613f62575f5ffd5b613f6e8c828d01613e01565b93505060c08a01356001600160401b03811115613f89575f5ffd5b613f958c828d01613e79565b92505060e08a01356001600160401b03811115613fb0575f5ffd5b613fbc8c828d01613e79565b9150509295985092959850929598565b5f5f5f5f5f5f60a08789031215613fe1575f5ffd5b8635613fec81613b1f565b955060208701359450604087013593506060870135925060808701356001600160401b0381111561401b575f5ffd5b61402789828a01613c4d565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a03121561404f575f5ffd5b873561405a81613b1f565b9650602088013561406a81613b1f565b955060408801359450606088013593506080880135925060a08801356001600160401b03811115614099575f5ffd5b6140a58a828b01613c4d565b989b979a50959850939692959293505050565b5f82601f8301126140c7575f5ffd5b81356140d281613d3a565b6040516140df8282613b92565b80915082815260208101915060208360051b860101925085831115614102575f5ffd5b602085015b83811015613e6f5780356001600160401b03811115614124575f5ffd5b614133886020838a0101613e79565b84525060209283019201614107565b5f5f5f5f5f60808688031215614156575f5ffd5b85356001600160401b0381111561416b575f5ffd5b61417788828901613c4d565b90965094505060208601356001600160401b03811115614195575f5ffd5b8601601f810188136141a5575f5ffd5b80356141b081613d3a565b6040516141bd8282613b92565b80915082815260208101915060208360051b85010192508a8311156141e0575f5ffd5b602084015b838110156142205780356001600160401b03811115614202575f5ffd5b6142118d602083890101613e01565b845250602092830192016141e5565b50955050505060408601356001600160401b0381111561423e575f5ffd5b61424a888289016140b8565b92505060608601356001600160401b03811115614265575f5ffd5b614271888289016140b8565b9150509295509295909350565b5f6001600160401b0382111561429657614296613b59565b50601f01601f191660200190565b5f6142ae8361427e565b6040516142bb8282613b92565b8092508481528585850111156142cf575f5ffd5b848460208301375f6020868301015250509392505050565b5f82601f8301126142f6575f5ffd5b6123ef838335602085016142a4565b5f5f60408385031215614316575f5ffd5b823561432181613b1f565b915060208301356001600160401b0381111561433b575f5ffd5b614347858286016142e7565b9150509250929050565b5f60208284031215614361575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b031215614380575f5ffd5b883561438b81613b1f565b9750602089013561439b81613b1f565b96506040890135955060608901359450608089013593506143bf8a60a08b01613bce565b92506101808901356001600160401b038111156143da575f5ffd5b6143e68b828c01613c4d565b999c989b5096995094979396929594505050565b5f5f6040838503121561440b575f5ffd5b823561441681613b1f565b9150602083013561442681613b1f565b809150509250929050565b5f5f60408385031215614442575f5ffd5b823561444d81613b1f565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b8181101561449b5783516001600160a01b0316835260209384019390920191600101614474565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b8181101561456d5760df198786030183526145588585516144a6565b9450602093840193929092019160010161453c565b50929695505050505050565b60ff60f81b8816815260e060208201525f61459760e08301896144a6565b82810360408401526145a981896144a6565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156145fe5783518352602093840193909201916001016145e0565b50909b9a5050505050505050505050565b5f5f5f5f60808587031215614622575f5ffd5b843561462d81613b1f565b935061463b60208601613bbe565b9250604085013561464b81613b1f565b9150606085013561465b81613b1f565b939692955090935050565b602081525f6123ef60208301846144a6565b5f60208284031215614688575f5ffd5b6123ef82613bbe565b5f602082840312156146a1575f5ffd5b81356001600160401b038111156146b6575f5ffd5b8201601f810184136146c6575f5ffd5b80356146d181613d3a565b6040516146de8282613b92565b80915082815260208101915060208360051b850101925086831115614701575f5ffd5b6020840193505b82841015613df7578335825260209384019390910190614708565b5f5f5f60608486031215614735575f5ffd5b833561474081613b1f565b925060208401356001600160401b0381111561475a575f5ffd5b8401601f8101861361476a575f5ffd5b614779868235602084016142a4565b92505061478860408501613bbe565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115611091576110916147a5565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b86811015612d4657838303601f19018852813536879003601e19018112614830575f5ffd5b86016020810190356001600160401b0381111561484b575f5ffd5b803603821315614859575f5ffd5b6148648582846147cc565b60209a8b019a9095509390930192505060010161480b565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f906110a990830184866147f4565b600181811c908216806148d057607f821691505b6020821081036148ee57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156126b057805f5260205f20601f840160051c810160208510156149195750805b601f840160051c820191505b81811015612cf3575f8155600101614925565b81516001600160401b0381111561495157614951613b59565b6149658161495f84546148bc565b846148f4565b6020601f821160018114614997575f83156149805750848201515b5f19600385901b1c1916600184901b178455612cf3565b5f84815260208120601f198516915b828110156149c657878501518255602094850194600190920191016149a6565b50848210156149e357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6149fc83613d3a565b604051614a098282613b92565b848152602081019150600585901b840136811115614a25575f5ffd5b845b8181101561449b5780356001600160401b03811115614a44575f5ffd5b614a50368289016142e7565b85525060209384019301614a27565b5f8235609e19833603018112614a73575f5ffd5b9190910192915050565b5f5f8335601e19843603018112614a92575f5ffd5b8301803591506001600160401b03821115614aab575f5ffd5b6020019150600581901b3603821315613c8d575f5ffd5b805180151581146118f0575f5ffd5b5f5f5f5f60808587031215614ae4575f5ffd5b84516020860151604087015191955093509150614b0360608601614ac2565b905092959194509250565b5f60208284031215614b1e575f5ffd5b5051919050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f6123ef600d830184614b25565b5f60208284031215614b6d575f5ffd5b81516001600160401b03811115614b82575f5ffd5b8201601f81018413614b92575f5ffd5b8051614b9d8161427e565b604051614baa8282613b92565b828152866020848601011115614bbe575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b606081525f614bee60608301866144a6565b8281036020840152614c0081866144a6565b91505060ff83166040830152949350505050565b5f614c28614c228386614b25565b84614b25565b949350505050565b8082028115828204841417611091576110916147a5565b5f60208284031215614c57575f5ffd5b6123ef82614ac2565b5f60033d11156110cf5760045f5f3e505f5160e01c90565b5f60443d1015614c855790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614cae57505090565b80820180516001600160401b03811115614cc9575050505090565b3d8401600319018282016020011115614ce3575050505090565b614cf260208285010185613b92565b509392505050565b5f5f60233d1115614d1357602060045f3e50505f516001905b9091565b5f82614d3157634e487b7160e01b5f52601260045260245ffd5b500690565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f906110a990830184866147f4565b81810381811115611091576110916147a5565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f6123ef8284614b2556fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212204dac763cd594718dee3415dd7f1e8464521d54c4a13f282733562a6f3a84dfb964736f6c634300081c0033",
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

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) Finalize(opts *bind.TransactOpts, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "finalize", index, token, to, value, extraData, v, r, s)
}

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Finalize(&_BridgeCross.TransactOpts, index, token, to, value, extraData, v, r, s)
}

// Finalize is a paid mutator transaction binding the contract method 0x2f63b7c8.
//
// Solidity: function finalize(uint256 index, address token, address to, uint256 value, bytes[] extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) Finalize(index *big.Int, token common.Address, to common.Address, value *big.Int, extraData [][]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.Finalize(&_BridgeCross.TransactOpts, index, token, to, value, extraData, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactor) FinalizeBatch(opts *bind.TransactOpts, args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "finalizeBatch", args, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeCross *BridgeCrossSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.FinalizeBatch(&_BridgeCross.TransactOpts, args, v, r, s)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x49a495ec.
//
// Solidity: function finalizeBatch((uint256,address,address,uint256,bytes[])[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BridgeCross *BridgeCrossTransactorSession) FinalizeBatch(args []IBridgeStandardFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BridgeCross.Contract.FinalizeBatch(&_BridgeCross.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address crossMintableERC20Code, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactor) Initialize(opts *bind.TransactOpts, crossMintableERC20Code common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "initialize", crossMintableERC20Code, _threshold, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address crossMintableERC20Code, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossSession) Initialize(crossMintableERC20Code common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, _threshold, rewardWallet_, _bridgeTokenInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address crossMintableERC20Code, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo) returns()
func (_BridgeCross *BridgeCrossTransactorSession) Initialize(crossMintableERC20Code common.Address, _threshold uint8, rewardWallet_ common.Address, _bridgeTokenInfo common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.Initialize(&_BridgeCross.TransactOpts, crossMintableERC20Code, _threshold, rewardWallet_, _bridgeTokenInfo)
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

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.ResetValidators(&_BridgeCross.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BridgeCross *BridgeCrossTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BridgeCross.Contract.ResetValidators(&_BridgeCross.TransactOpts, validators)
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
