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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addTokenDeploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"bridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenInfo\",\"outputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"finalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"sigs\",\"type\":\"bytes[][]\"}],\"name\":\"finalizeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossMintableERC20Code\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardWallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BridgeFeeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isValidToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBridgeStandard.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"permitBridgeTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"internalType\":\"structIBridgeStandard.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeTokenInfo\",\"name\":\"_bridgeTokenInfo\",\"type\":\"address\"}],\"name\":\"setTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenInfoLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xcross\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"pairToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"extraData\",\"type\":\"bytes[]\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"PairUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInsufficientValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeCrossInvalidValueUnit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeStandardCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeStandardFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedService\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualService\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeStandardNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5476bd72": "addToken(address,address)",
		"f7684ec4": "addTokenDeploy(address,string,uint8)",
		"c97682f8": "allPairs()",
		"9fdf1c6a": "allTokenInfo()",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"37d10075": "bridge(address,uint256,uint256,uint256,bytes[])",
		"3f5d3b5d": "bridgeTo(address,address,uint256,uint256,uint256,bytes[])",
		"9c1b65a9": "bridgeTokenInfo()",
		"6e908ca3": "calculate(address,uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"5dbe47e8": "contains(address)",
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
		"211c8ac0": "xcross()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614ee36100f95f395f81816123440152818161236d01526124bc0152614ee35ff3fe608060405260043610610368575f3560e01c80637cfed602116101c8578063c0c53b8b116100fd578063f120c4001161009d578063f7684ec41161006d578063f7684ec4146109c9578063facd743b146109e8578063fb75b2c714610a07578063fe2b8da614610a24575f5ffd5b8063f120c4001461096f578063f2fde38b14610982578063f30589c3146109a1578063f698da25146109b5575f5ffd5b8063cbae5958116100d8578063cbae595814610909578063d92fc67b14610928578063e1af7f501461093c578063e70a1b2614610950575f5ffd5b8063c0c53b8b146108b7578063c1876453146108d6578063c97682f8146108f5575f5ffd5b806396ce079511610168578063a982318311610143578063a982318314610828578063ad3cb1cc14610847578063aed1d40314610884578063b7f3358d14610898575f5ffd5b806396ce0795146107d45780639c1b65a9146107e85780639fdf1c6a14610807575f5ffd5b806385547884116101a357806385547884146107515780638da5cb5b1461076557806391cf6d3e146107a15780639300c926146107b5575f5ffd5b80637cfed602146107025780638456cb591461071657806384b0196e1461072a575f5ffd5b80634f1ef2861161029e5780635fa7b5841161023e5780637021fd0e116102195780637021fd0e1461066c578063715018a61461069857806371c59d7b146106ac5780637c41ad2c146106e3575f5ffd5b80635fa7b584146105f25780636e908ca3146106115780636ff97f1d1461064b575f5ffd5b806352d1902d1161027957806352d1902d1461057d5780635476bd72146105915780635c975abb146105b05780635dbe47e8146105d3575f5ffd5b80634f1ef286146105385780634f6ccce71461054b57806351c455791461056a575f5ffd5b80632f9b59d1116103095780633f4ba83a116102e45780633f4ba83a146104d15780633f5d3b5d146104e557806340a141ff146104f857806342cde4e814610517575f5ffd5b80632f9b59d11461048057806337d100751461049f5780633b3bff0f146104b2575f5ffd5b80631d40f0d8116103445780631d40f0d8146103dc5780631f69565f146103fb578063211c8ac01461042757806327938c811461045e575f5ffd5b80628bd028146103825780631327d3d8146103aa578063174991ab146103c9575f5ffd5b3661037e575f341161037c5761037c613c53565b005b5f5ffd5b610395610390366004613e48565b610a43565b60405190151581526020015b60405180910390f35b3480156103b5575f5ffd5b5061037c6103c4366004613f47565b610b6f565b6103956103d7366004613ff1565b610b7d565b3480156103e7575f5ffd5b5061037c6103f6366004614097565b610c6d565b348015610406575f5ffd5b5061041a610415366004613f47565b610ca7565b6040516103a1919061413c565b348015610432575f5ffd5b5060fd54610446906001600160a01b031681565b6040516001600160a01b0390911681526020016103a1565b348015610469575f5ffd5b50610472610d4a565b6040519081526020016103a1565b34801561048b575f5ffd5b5061037c61049a366004613f47565b610d5f565b6103956104ad366004614170565b610dd2565b3480156104bd575f5ffd5b5061037c6104cc366004613f47565b610ded565b3480156104dc575f5ffd5b5061037c610dfe565b6103956104f33660046141dd565b610e10565b348015610503575f5ffd5b5061037c610512366004613f47565b610eff565b348015610522575f5ffd5b5060995460405160ff90911681526020016103a1565b61037c61054636600461425c565b610f09565b348015610556575f5ffd5b506104466105653660046142a8565b610f24565b6103956105783660046142bf565b610f2f565b348015610588575f5ffd5b50610472610f4e565b34801561059c575f5ffd5b5061037c6105ab366004614351565b610f6a565b3480156105bb575f5ffd5b505f516020614e6e5f395f51905f525460ff16610395565b3480156105de575f5ffd5b506103956105ed366004613f47565b610f7c565b3480156105fd575f5ffd5b5061037c61060c366004613f47565b610f87565b34801561061c575f5ffd5b5061063061062b366004614388565b610f98565b604080519384526020840192909252908201526060016103a1565b348015610656575f5ffd5b5061065f61101f565b6040516103a191906143b2565b348015610677575f5ffd5b5061068b6106863660046142a8565b611080565b6040516103a1919061442b565b3480156106a3575f5ffd5b5061037c6111e0565b3480156106b7575f5ffd5b506104466106c6366004613f47565b6001600160a01b039081165f908152603360205260409020541690565b3480156106ee575f5ffd5b5061037c6106fd366004613f47565b6111f1565b34801561070d575f5ffd5b50610472611202565b348015610721575f5ffd5b5061037c61126d565b348015610735575f5ffd5b5061073e61127d565b6040516103a197969594939291906144d0565b34801561075c575f5ffd5b5061037c611326565b348015610770575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610446565b3480156107ac575f5ffd5b5060cb54610472565b3480156107c0575f5ffd5b5061037c6107cf366004614097565b611340565b3480156107df575f5ffd5b50610472611377565b3480156107f3575f5ffd5b5060ca54610446906001600160a01b031681565b348015610812575f5ffd5b5061081b6113be565b6040516103a19190614566565b348015610833575f5ffd5b506103956108423660046142a8565b61142c565b348015610852575f5ffd5b50610877604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516103a191906145cf565b34801561088f575f5ffd5b506104726116eb565b3480156108a3575f5ffd5b5061037c6108b23660046145e1565b6116f6565b3480156108c2575f5ffd5b5061037c6108d13660046145fa565b611747565b3480156108e1575f5ffd5b506103956108f0366004613f47565b611880565b348015610900575f5ffd5b5061065f6118b0565b348015610914575f5ffd5b506104466109233660046142a8565b611996565b348015610933575f5ffd5b506104726119a2565b348015610947575f5ffd5b506104726119ac565b34801561095b575f5ffd5b5061041a61096a3660046142a8565b6119bc565b61039561097d366004614642565b611a1d565b34801561098d575f5ffd5b5061037c61099c366004613f47565b611ce3565b3480156109ac575f5ffd5b5061065f611d1d565b3480156109c0575f5ffd5b50610472611d29565b3480156109d4575f5ffd5b506104466109e33660046146e0565b611d32565b3480156109f3575f5ffd5b50610395610a02366004613f47565b611e72565b348015610a12575f5ffd5b5060cc546001600160a01b0316610446565b348015610a2f575f5ffd5b50610877610a3e3660046142a8565b611e7e565b5f805b83811015610b6257610b59858583818110610a6357610a6361474e565b9050602002810190610a759190614762565b35868684818110610a8857610a8861474e565b9050602002810190610a9a9190614762565b610aab906040810190602001613f47565b878785818110610abd57610abd61474e565b9050602002810190610acf9190614762565b610ae0906060810190604001613f47565b888886818110610af257610af261474e565b9050602002810190610b049190614762565b60600135898987818110610b1a57610b1a61474e565b9050602002810190610b2c9190614762565b610b3a906080810190614780565b898881518110610b4c57610b4c61474e565b6020026020010151611a1d565b50600101610a46565b50600190505b9392505050565b610b7a816001611f1d565b50565b5f610b86611fe6565b89610b9081611880565b8190610bc057604051631fa1cbb560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610bc9612016565b5f5f5f5f610bd98f8d8d8d61204d565b929650909450925090508383838e8e8e86610c2a576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610bb7565b505050505050610c418f8f8f8f87878f8f8f612120565b6001955050505050610c5f60015f516020614e8e5f395f51905f5255565b509998505050505050505050565b5f5b8151811015610ca357610c9b828281518110610c8d57610c8d61474e565b60200260200101515f611f1d565b600101610c6f565b5050565b610cd760405180608001604052805f6001600160a01b031681526020015f81526020015f81526020015f81525090565b60ca54604051631f69565f60e01b81526001600160a01b03848116600483015290911690631f69565f906024015b608060405180830381865afa158015610d20573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d449190614831565b92915050565b5f6066546001610d5a919061485f565b905090565b610d676121a9565b6001600160a01b038116610db057604051633effd40360e21b815260206004820152600f60248201526e627269646765546f6b656e496e666f60881b6044820152606401610bb7565b60ca80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610de287338888888888610e10565b979650505050505050565b610df56121a9565b610b7a81612204565b610e066121a9565b610e0e6122a4565b565b5f610e19611fe6565b87610e2381611880565b8190610e4e57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bb7565b50610e57612016565b5f5f5f5f610e678d8c8c8c61204d565b929650909450925090508383838d8d8d86610eb8576040516334191f1d60e21b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610bb7565b505050505050610ed58d610ec93390565b8e8e87875f8f8f6122fd565b6001955050505050610ef360015f516020614e8e5f395f51905f5255565b50979650505050505050565b610b7a815f611f1d565b610f11612339565b610f1a826123dd565b610ca382826123e5565b5f610d4481836124a6565b5f610f4189898a8a8a8a8a8a8a610b7d565b9998505050505050505050565b5f610f576124b1565b505f516020614e4e5f395f51905f525b90565b610f726121a9565b610ca382826124fa565b5f610d448183612569565b610f8f6121a9565b610b7a8161258a565b60ca54604051636e908ca360e01b81526001600160a01b038481166004830152602482018490525f928392839290911690636e908ca390604401606060405180830381865afa158015610fed573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110119190614872565b919790965090945092505050565b60605f61102b5f6125f2565b9050806001600160401b0381111561104557611045613cae565b60405190808252806020026020018201604052801561106e578160200160208202803683370190505b50915061107a5f6125fb565b91505090565b6110c06040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919592946080870194939192919084015b828210156111d2578382905f5260205f200180546111479061489d565b80601f01602080910402602001604051908101604052809291908181526020018280546111739061489d565b80156111be5780601f10611195576101008083540402835291602001916111be565b820191905f5260205f20905b8154815290600101906020018083116111a157829003601f168201915b50505050508152602001906001019061112a565b505050915250909392505050565b6111e86121a9565b610e0e5f612607565b6111f96121a9565b610b7a81612677565b60ca546040805163d92fc67b60e01b815290515f926001600160a01b03169163d92fc67b9160048083019260209291908290030181865afa158015611249573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d5a91906148d5565b6112756121a9565b610e0e6126f7565b5f60608082808083815f516020614e2e5f395f51905f5280549091501580156112a857506001810154155b6112ec5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610bb7565b6112f461273f565b6112fc6127ff565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61132e6121a9565b60ca80546001600160a01b0319169055565b5f5b8151811015610ca35761136f8282815181106113605761136061474e565b60200260200101516001611f1d565b600101611342565b60ca54604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015611249573d5f5f3e3d5ffd5b60ca5460408051634fef8e3560e11b815290516060926001600160a01b031691639fdf1c6a916004808301925f9291908290030181865afa158015611405573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d5a91908101906148ec565b5f611435611fe6565b61143e33611e72565b3390611469576040516338905e7160e11b81526001600160a01b039091166004820152602401610bb7565b50611472612016565b61147d60cf8361283d565b829061149f5760405163473978bf60e01b8152600401610bb791815260200190565b505f82815260cd60209081526040808320815160a0810183528154815260018201546001600160a01b03908116828601526002830154168184015260038201546060820152600482018054845181870281018701909552808552919492936080860193909290879084015b828210156115b2578382905f5260205f200180546115279061489d565b80601f01602080910402602001604051908101604052809291908181526020018280546115539061489d565b801561159e5780601f106115755761010080835404028352916020019161159e565b820191905f5260205f20905b81548152906001019060200180831161158157829003601f168201915b50505050508152602001906001019061150a565b505050508152505090505f5f6115d5836020015184604001518560600151612854565b915091508181906115f95760405162461bcd60e51b8152600401610bb791906145cf565b5061160560cf86612a2a565b505f85815260ce6020526040812061161c91613b7d565b5f85815260cd602052604081208181556001810180546001600160a01b0319908116909155600282018054909116905560038101829055906116616004830182613bb4565b505082604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738660600151426040516116c1929190918252602082015260400190565b60405180910390a4600193505050506116e660015f516020614e8e5f395f51905f5255565b919050565b5f610d5a60976125f2565b6116fe6121a9565b6099805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561178b5750825b90505f826001600160401b031660011480156117a65750303b155b9050811580156117b4575080155b156117d25760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156117fc57845460ff60401b1916600160401b1785555b6118068787612a35565b60fd805460016001600160a01b03199182161790915560fe80549091166001600160a01b038a16179055831561187657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f61188a82610f7c565b8015610d445750506001600160a01b03165f9081526034602052604090205460ff161590565b60605f6118bb61101f565b90505f81516001600160401b038111156118d7576118d7613cae565b604051908082528060200260200182016040528015611900578160200160208202803683370190505b5090505f5b825181101561198f5760335f8483815181106119235761192361474e565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f205f9054906101000a90046001600160a01b031682828151811061196f5761196f61474e565b6001600160a01b0390921660209283029190910190910152600101611905565b5092915050565b5f610d446097836124a6565b5f610d5a5f6125f2565b5f6065546001610d5a919061485f565b6119ec60405180608001604052805f6001600160a01b031681526020015f81526020015f81526020015f81525090565b60ca546040516373850d9360e11b8152600481018490526001600160a01b039091169063e70a1b2690602401610d05565b5f611a26611fe6565b86611a3081611880565b8190611a5b57604051631fa1cbb560e01b81526001600160a01b039091166004820152602401610bb7565b50611a64612016565b5f611a6d610d4a565b9050808a808214611a9a57604051634982205b60e11b815260048101929092526024820152604401610bb7565b50505f7fd9b9f4a2428e1ffc6e0db614f2e1cbb0e18b038805f974bc76371611511d1ded8b8b8b8b8b8b604051602001611ada9796959493929190614a44565b604051602081830303815290604052805190602001209050611afc8186612ae1565b8b90611b1e57604051635b777d1160e11b8152600401610bb791815260200190565b50611b2d606680546001019055565b5f5f611b3a8c8c8c612854565b915091508115611ba0578a6001600160a01b03168c6001600160a01b03168e7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738d42604051611b93929190918252602082015260400190565b60405180910390a4610ed5565b611bab60cf8e612c2b565b8d90611bcd576040516368db995b60e11b8152600401610bb791815260200190565b505f8d815260ce60205260409020611be58282614ac8565b506040518060a001604052808e81526020018d6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a8a90611c289190614b82565b90525f8e815260cd602090815260409182902083518155838201516001820180546001600160a01b03199081166001600160a01b0393841617909155938501516002830180549095169116179092556060830151600383015560808301518051611c989260048501920190613bcf565b50506040518e91507ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a26001955050505050610ef360015f516020614e8e5f395f51905f5255565b611ceb6121a9565b6001600160a01b038116611d1457604051631e4fbdf760e01b81525f6004820152602401610bb7565b610b7a81612607565b6060610d5a60976125fb565b5f610d5a612c36565b5f611d3b6121a9565b5f83604051602001611d4d9190614ba5565b60408051601f19818403018152908290526bffffffffffffffffffffffff19606088901b16602083015291505f906034016040516020818303038152906040528051906020012090505f60fe5f9054906101000a90046001600160a01b03166001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015611de6573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611e0d9190810190614bc6565b838787604051602001611e2293929190614c45565b60408051601f1981840301815290829052611e409291602001614c7d565b6040516020818303038152906040529050611e5c5f8383612c3f565b9350611e688488610f6a565b5050509392505050565b5f610d44609783612569565b5f81815260ce60205260409020805460609190611e9a9061489d565b80601f0160208091040260200160405190810160405280929190818152602001828054611ec69061489d565b8015611f115780601f10611ee857610100808354040283529160200191611f11565b820191905f5260205f20905b815481529060010190602001808311611ef457829003601f168201915b50505050509050919050565b611f256121a9565b8015611f6757611f36609783612cd3565b8290611f615760405163082cdf5560e21b81526001600160a01b039091166004820152602401610bb7565b50611f9f565b611f72609783612ce7565b8290611f9d5760405163e491024560e01b81526001600160a01b039091166004820152602401610bb7565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f516020614e6e5f395f51905f525460ff1615610e0e5760405163d93c066560e01b815260040160405180910390fd5b5f516020614e8e5f395f51905f5280546001190161204757604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60ca545f908190819081906001600160a01b031661207657505f92508291508190506001612115565b60ca54604051636e908ca360e01b81526001600160a01b038a81166004830152602482018a905290911690636e908ca390604401606060405180830381865afa1580156120c5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120e99190614872565b919550935091508286108015906121005750818510155b801561210c5750838710155b15612115575060015b945094509450949050565b8361212b868861485f565b612135919061485f565b836040015110158987909161216e57604051638d424aa160e01b81526001600160a01b0390921660048301526024820152604401610bb7565b505061217983612cfb565b61218b898989898989600189896122fd565b505050505050505050565b60015f516020614e8e5f395f51905f5255565b336121db7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610e0e5760405163118cdaa760e01b8152336004820152602401610bb7565b61220d81610f7c565b801561223057506001600160a01b0381165f9081526034602052604090205460ff165b819061225b576040516340da71e560e01b81526001600160a01b039091166004820152602401610bb7565b506001600160a01b0381165f81815260346020526040808220805460ff19169055517ff38578ed892ce2ce655ca8ae03c73464ad74915a1331a9b4085e637534daeedf9190a250565b6122ac612dff565b5f516020614e6e5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200161173c565b61231289898861230d888a61485f565b612e2e565b61232b61231d6119ac565b8a8a8a8a8a8a8a8a8a612fd4565b61218b606580546001019055565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806123bf57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166123b35f516020614e4e5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610e0e5760405163703e46dd60e11b815260040160405180910390fd5b610b7a6121a9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561243f575060408051601f3d908101601f1916820190925261243c918101906148d5565b60015b61246757604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610bb7565b5f516020614e4e5f395f51905f52811461249757604051632a87526960e21b815260048101829052602401610bb7565b6124a1838361309a565b505050565b5f610b6883836130ef565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e0e5760405163703e46dd60e11b815260040160405180910390fd5b61250382613115565b6001600160a01b038281165f8181526033602090815260409182902080546001600160a01b03191694861694851790559051600181527fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce91015b60405180910390a35050565b6001600160a01b0381165f9081526001830160205260408120541515610b68565b612593816131c2565b6001600160a01b038181165f81815260336020908152604080832080546001600160a01b031981169091559051928352909316928392917fa0be238839c65552a544da5e66b7a224fe553eccb19681bc2ddd043f4ec863ce910161255d565b5f610d44825490565b60605f610b688361326f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61268081611880565b81906126ab576040516340da71e560e01b81526001600160a01b039091166004820152602401610bb7565b506001600160a01b0381165f81815260346020526040808220805460ff19166001179055517ff017c0de579727a3cd3ee18077ee8b4c43bf21892985952d1d5a0d52f983502d9190a250565b6126ff611fe6565b5f516020614e6e5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336122e5565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020614e2e5f395f51905f529161277d9061489d565b80601f01602080910402602001604051908101604052809291908181526020018280546127a99061489d565b80156127f45780601f106127cb576101008083540402835291602001916127f4565b820191905f5260205f20905b8154815290600101906020018083116127d757829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020614e2e5f395f51905f529161277d9061489d565b5f8181526001830160205260408120541515610b68565b5f60608215612a225760fd546001600160a01b03908116908616036128a957612891612881606485614c99565b6001600160a01b038616906132c7565b505060408051602081019091525f8152600190612a22565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528616906340c10f19906044016020604051808303815f875af1925050508015612913575060408051601f3d908101601f1916820190925261291091810190614cb0565b60015b6129c45761291f614ccf565b806308c379a0036129485750612933614ce7565b8061293e575061298b565b5f92509050612a22565b634e487b710361298b5761295a614d69565b90612965575061298b565b60408051602081018390525f945001604051602081830303815290604052915050612a22565b3d8080156129b4576040519150601f19603f3d011682016040523d82523d5f602084013e6129b9565b606091505b505f92509050612a22565b80156129e4576001925060405180602001604052805f8152509150612a20565b5f92506040518060400160405280601881526020017f42726964676543726f73733a206d696e74206661696c6564000000000000000081525091505b505b935093915050565b5f610b688383613353565b612a3d613436565b612a463361347f565b612a4e613490565b612a56613498565b612a5e6134a8565b612a666134b8565b6001600160a01b038216612aac57604051633effd40360e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610bb7565b4360cb5560ca80546001600160a01b039283166001600160a01b03199182161790915560cc8054939092169216919091179055565b80516099545f919060ff16811015612afc575f915050610d44565b5f80826001600160401b03811115612b1657612b16613cae565b604051908082528060200260200182016040528015612b3f578160200160208202803683370190505b5090505f5b8551811015612c19575f868281518110612b6057612b6061474e565b60200260200101519050604181511015612b81575f95505050505050610d44565b5f612b9582612b8f8b613517565b90613543565b9050612ba081611e72565b612bb2575f9650505050505050610d44565b5f805b8551811015612c0157858181518110612bd057612bd061474e565b60200260200101516001600160a01b0316836001600160a01b031603612bf95760019150612c01565b600101612bb5565b5080612c0e578560010195505b505050600101612b44565b505060995460ff161115949350505050565b5f610b68838361356b565b5f610d5a6135b7565b5f83471015612c6a5760405163cf47918160e01b815247600482015260248101859052604401610bb7565b81515f03612c8b57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615612cac576040513d5f823e3d81fd5b6001600160a01b038116610b685760405163b06ebf3d60e01b815260040160405180910390fd5b5f610b68836001600160a01b03841661356b565b5f610b68836001600160a01b038416613353565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180612da7576040513d5f823e3d81fd5b50505f513d91508115612dbe578060011415612dcc565b84516001600160a01b03163b155b15612df8578451604051637ec77ed960e11b81526001600160a01b039091166004820152602401610bb7565b5050505050565b5f516020614e6e5f395f51905f525460ff16610e0e57604051638dfc202b60e01b815260040160405180910390fd5b60fd546001600160a01b0390811690851603612eef57816064612e518183614d86565b612e5b9190614c99565b143490612e7e576040516308dc47c960e41b8152600401610bb791815260200190565b50612e89818361485f565b3414612e95828461485f565b349091612ebd576040516290c38760e71b815260048101929092526024820152604401610bb7565b50508015612eea57612eea81612edb60cc546001600160a01b031690565b6001600160a01b0316906132c7565b612fce565b8015612f1d57612f1d83612f0b60cc546001600160a01b031690565b6001600160a01b03871691908461362a565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201849052851690639dc29fac906044016020604051808303815f875af1158015612f69573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f8d9190614cb0565b848484909192612fca576040516336b52daf60e21b81526001600160a01b0393841660048201529290911660248301526044820152606401610bb7565b5050505b50505050565b6001600160a01b038981165f81815260336020526040902054828b16928d917f321422e869147e17ea7650432359e138192fd7c8e8b8d3e0b4a1d7f9fe85409391168b8b428a8a8a60405161302f9796959493929190614da5565b60405180910390a4876001600160a01b0316896001600160a01b03168b7f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8888604051613086929190918252602082015260400190565b60405180910390a450505050505050505050565b6130a382613684565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156130e7576124a182826136e7565b610ca3613759565b5f825f0182815481106131045761310461474e565b905f5260205f200154905092915050565b806001600160a01b038116613155576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bb7565b61315f5f83612cd3565b829061318a576040516351eccfe160e11b81526001600160a01b039091166004820152602401610bb7565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b806001600160a01b038116613202576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610bb7565b61320c5f83612ce7565b8290613237576040516340da71e560e01b81526001600160a01b039091166004820152602401610bb7565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611f1157602002820191905f5260205f20905b8154815260200190600101908083116132a85750505050509050919050565b804710156132f15760405163cf47918160e01b815247600482015260248101829052604401610bb7565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461333b576040519150601f19603f3d011682016040523d82523d5f602084013e613340565b606091505b509150915081612fce57612fce81613778565b5f818152600183016020526040812054801561342d575f613375600183614de7565b85549091505f9061338890600190614de7565b90508082146133e7575f865f0182815481106133a6576133a661474e565b905f5260205f200154905080875f0184815481106133c6576133c661474e565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806133f8576133f8614dfa565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610d44565b5f915050610d44565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e0e57604051631afcd79f60e31b815260040160405180910390fd5b613487613436565b610b7a816137a1565b610e0e613436565b6134a0613436565b610e0e6137a9565b6134b0613436565b610e0e6137c9565b6134c0613436565b613508604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506137d1565b6099805460ff19166003179055565b5f610d44613523612c36565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f61355186866137e3565b925092509250613561828261382c565b5090949350505050565b5f8181526001830160205260408120546135b057508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610d44565b505f610d44565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6135e16138e4565b6135e961394c565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612fce90859061398e565b806001600160a01b03163b5f036136b957604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610bb7565b5f516020614e4e5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516137039190614e0e565b5f60405180830381855af49150503d805f811461373b576040519150601f19603f3d011682016040523d82523d5f602084013e613740565b606091505b50915091506137508583836139fa565b95945050505050565b3415610e0e5760405163b398979f60e01b815260040160405180910390fd5b8051156137885780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b611ceb613436565b6137b1613436565b5f516020614e6e5f395f51905f52805460ff19169055565b612196613436565b6137d9613436565b610ca38282613a56565b5f5f5f835160410361381a576020840151604085015160608601515f1a61380c88828585613ab5565b955095509550505050613825565b505081515f91506002905b9250925092565b5f82600381111561383f5761383f614e19565b03613848575050565b600182600381111561385c5761385c614e19565b0361387a5760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561388e5761388e614e19565b036138af5760405163fce698f760e01b815260048101829052602401610bb7565b60038260038111156138c3576138c3614e19565b03610ca3576040516335e2f38360e21b815260048101829052602401610bb7565b5f5f516020614e2e5f395f51905f52816138fc61273f565b80519091501561391457805160209091012092915050565b81548015613923579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020614e2e5f395f51905f52816139646127ff565b80519091501561397c57805160209091012092915050565b60018201548015613923579392505050565b5f5f60205f8451602086015f885af1806139ad576040513d5f823e3d81fd5b50505f513d915081156139c45780600114156139d1565b6001600160a01b0384163b155b15612fce57604051635274afe760e01b81526001600160a01b0385166004820152602401610bb7565b606082613a0f57613a0a82613778565b610b68565b8151158015613a2657506001600160a01b0384163b155b15613a4f57604051639996b31560e01b81526001600160a01b0385166004820152602401610bb7565b5080610b68565b613a5e613436565b5f516020614e2e5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613a978482614ac8565b5060038101613aa68382614ac8565b505f8082556001909101555050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613aee57505f91506003905082613b73565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613b3f573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613b6a57505f925060019150829050613b73565b92505f91508190505b9450945094915050565b508054613b899061489d565b5f825580601f10613b98575050565b601f0160209004905f5260205f2090810190610b7a9190613c23565b5080545f8255905f5260205f2090810190610b7a9190613c37565b828054828255905f5260205f20908101928215613c13579160200282015b82811115613c135782518290613c039082614ac8565b5091602001919060010190613bed565b50613c1f929150613c37565b5090565b5b80821115613c1f575f8155600101613c24565b80821115613c1f575f613c4a8282613b7d565b50600101613c37565b634e487b7160e01b5f52600160045260245ffd5b5f5f83601f840112613c77575f5ffd5b5081356001600160401b03811115613c8d575f5ffd5b6020830191508360208260051b8501011115613ca7575f5ffd5b9250929050565b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b0382111715613ce157613ce1613cae565b60405250565b601f8201601f191681016001600160401b0381118282101715613d0c57613d0c613cae565b6040525050565b5f6001600160401b03821115613d2b57613d2b613cae565b5060051b60200190565b5f6001600160401b03821115613d4d57613d4d613cae565b50601f01601f191660200190565b5f613d6583613d35565b604051613d728282613ce7565b809250848152858585011115613d86575f5ffd5b848460208301375f6020868301015250509392505050565b5f82601f830112613dad575f5ffd5b610b6883833560208501613d5b565b5f613dc683613d13565b604051613dd38282613ce7565b84815291505060208101600584901b830185811115613df0575f5ffd5b835b81811015611e685780356001600160401b03811115613e0f575f5ffd5b613e1b88828801613d9e565b84525060209283019201613df2565b5f82601f830112613e39575f5ffd5b610b6883833560208501613dbc565b5f5f5f60408486031215613e5a575f5ffd5b83356001600160401b03811115613e6f575f5ffd5b613e7b86828701613c67565b90945092505060208401356001600160401b03811115613e99575f5ffd5b8401601f81018613613ea9575f5ffd5b8035613eb481613d13565b604051613ec18282613ce7565b80915082815260208101915060208360051b850101925088831115613ee4575f5ffd5b602084015b83811015613f245780356001600160401b03811115613f06575f5ffd5b613f158b602083890101613e2a565b84525060209283019201613ee9565b50809450505050509250925092565b6001600160a01b0381168114610b7a575f5ffd5b5f60208284031215613f57575f5ffd5b8135610b6881613f33565b803560ff811681146116e6575f5ffd5b5f60e08284031215613f82575f5ffd5b604051613f8e81613cc2565b8091508235613f9c81613f33565b81526020830135613fac81613f33565b60208201526040838101359082015260608084013590820152613fd160808401613f62565b608082015260a0838101359082015260c092830135920191909152919050565b5f5f5f5f5f5f5f5f5f6101c08a8c03121561400a575f5ffd5b893561401581613f33565b985060208a013561402581613f33565b975060408a013561403581613f33565b965060608a0135955060808a0135945060a08a013593506140598b60c08c01613f72565b92506101a08a01356001600160401b03811115614074575f5ffd5b6140808c828d01613c67565b915080935050809150509295985092959850929598565b5f602082840312156140a7575f5ffd5b81356001600160401b038111156140bc575f5ffd5b8201601f810184136140cc575f5ffd5b80356140d781613d13565b6040516140e48282613ce7565b80915082815260208101915060208360051b850101925086831115614107575f5ffd5b6020840193505b8284101561413257833561412181613f33565b82526020938401939091019061410e565b9695505050505050565b81516001600160a01b0316815260208083015190820152604080830151908201526060808301519082015260808101610d44565b5f5f5f5f5f5f60a08789031215614185575f5ffd5b863561419081613f33565b955060208701359450604087013593506060870135925060808701356001600160401b038111156141bf575f5ffd5b6141cb89828a01613c67565b979a9699509497509295939492505050565b5f5f5f5f5f5f5f60c0888a0312156141f3575f5ffd5b87356141fe81613f33565b9650602088013561420e81613f33565b955060408801359450606088013593506080880135925060a08801356001600160401b0381111561423d575f5ffd5b6142498a828b01613c67565b989b979a50959850939692959293505050565b5f5f6040838503121561426d575f5ffd5b823561427881613f33565b915060208301356001600160401b03811115614292575f5ffd5b61429e85828601613d9e565b9150509250929050565b5f602082840312156142b8575f5ffd5b5035919050565b5f5f5f5f5f5f5f5f6101a0898b0312156142d7575f5ffd5b88356142e281613f33565b975060208901356142f281613f33565b96506040890135955060608901359450608089013593506143168a60a08b01613f72565b92506101808901356001600160401b03811115614331575f5ffd5b61433d8b828c01613c67565b999c989b5096995094979396929594505050565b5f5f60408385031215614362575f5ffd5b823561436d81613f33565b9150602083013561437d81613f33565b809150509250929050565b5f5f60408385031215614399575f5ffd5b82356143a481613f33565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b818110156143f25783516001600160a01b03168352602093840193909201916001016143cb565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c082018351602084015260018060a01b03602085015116604084015260018060a01b03604085015116606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156144c45760df198786030183526144af8585516143fd565b94506020938401939290920191600101614493565b50929695505050505050565b60ff60f81b8816815260e060208201525f6144ee60e08301896143fd565b828103604084015261450081896143fd565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614555578351835260209384019390920191600101614537565b50909b9a5050505050505050505050565b602080825282518282018190525f918401906040840190835b818110156143f2576145b983855180516001600160a01b031682526020808201519083015260408082015190830152606090810151910152565b602093909301926080929092019160010161457f565b602081525f610b6860208301846143fd565b5f602082840312156145f1575f5ffd5b610b6882613f62565b5f5f5f6060848603121561460c575f5ffd5b833561461781613f33565b9250602084013561462781613f33565b9150604084013561463781613f33565b809150509250925092565b5f5f5f5f5f5f5f60c0888a031215614658575f5ffd5b87359650602088013561466a81613f33565b9550604088013561467a81613f33565b94506060880135935060808801356001600160401b0381111561469b575f5ffd5b6146a78a828b01613c67565b90945092505060a08801356001600160401b038111156146c5575f5ffd5b6146d18a828b01613e2a565b91505092959891949750929550565b5f5f5f606084860312156146f2575f5ffd5b83356146fd81613f33565b925060208401356001600160401b03811115614717575f5ffd5b8401601f81018613614727575f5ffd5b61473686823560208401613d5b565b92505061474560408501613f62565b90509250925092565b634e487b7160e01b5f52603260045260245ffd5b5f8235609e19833603018112614776575f5ffd5b9190910192915050565b5f5f8335601e19843603018112614795575f5ffd5b8301803591506001600160401b038211156147ae575f5ffd5b6020019150600581901b3603821315613ca7575f5ffd5b5f608082840312156147d5575f5ffd5b604051608081018181106001600160401b03821117156147f7576147f7613cae565b8060405250809150825161480a81613f33565b81526020838101519082015260408084015190820152606092830151920191909152919050565b5f60808284031215614841575f5ffd5b610b6883836147c5565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d4457610d4461484b565b5f5f5f60608486031215614884575f5ffd5b5050815160208301516040909301519094929350919050565b600181811c908216806148b157607f821691505b6020821081036148cf57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156148e5575f5ffd5b5051919050565b5f602082840312156148fc575f5ffd5b81516001600160401b03811115614911575f5ffd5b8201601f81018413614921575f5ffd5b805161492c81613d13565b6040516149398282613ce7565b80915082815260208101915060208360071b85010192508683111561495c575f5ffd5b6020840193505b828410156141325761497587856147c5565b8252602082019150608084019350614963565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208501945060208460051b820101835f5b86811015614a3857838303601f19018852813536879003601e190181126149ec575f5ffd5b86016020810190356001600160401b03811115614a07575f5ffd5b803603821315614a15575f5ffd5b614a20858284614988565b60209a8b019a909550939093019250506001016149c7565b50909695505050505050565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90610f4190830184866149b0565b601f8211156124a157805f5260205f20601f840160051c81016020851015614aa95750805b601f840160051c820191505b81811015612df8575f8155600101614ab5565b81516001600160401b03811115614ae157614ae1613cae565b614af581614aef845461489d565b84614a84565b6020601f821160018114614b27575f8315614b105750848201515b5f19600385901b1c1916600184901b178455612df8565b5f84815260208120601f198516915b82811015614b565787850151825560209485019460019092019101614b36565b5084821015614b7357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f610b68368484613dbc565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f610b68600d830184614b8e565b5f60208284031215614bd6575f5ffd5b81516001600160401b03811115614beb575f5ffd5b8201601f81018413614bfb575f5ffd5b8051614c0681613d35565b604051614c138282613ce7565b828152866020848601011115614c27575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b606081525f614c5760608301866143fd565b8281036020840152614c6981866143fd565b91505060ff83166040830152949350505050565b5f614c91614c8b8386614b8e565b84614b8e565b949350505050565b8082028115828204841417610d4457610d4461484b565b5f60208284031215614cc0575f5ffd5b81518015158114610b68575f5ffd5b5f60033d1115610f675760045f5f3e505f5160e01c90565b5f60443d1015614cf45790565b6040513d600319016004823e80513d60248201116001600160401b0382111715614d1d57505090565b80820180516001600160401b03811115614d38575050505090565b3d8401600319018282016020011115614d52575050505090565b614d6160208285010185613ce7565b509392505050565b5f5f60233d1115614d8257602060045f3e50505f516001905b9091565b5f82614da057634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160a01b038881168252871660208201526040810186905260608101859052831515608082015260c060a082018190525f90610f4190830184866149b0565b81810381811115610d4457610d4461484b565b634e487b7160e01b5f52603160045260245ffd5b5f610b688284614b8e565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220baa8d6928d30a005b3c608856f26c436a03a5cba80c79cff334b0671a5e3611b64736f6c634300081c0033",
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

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeCross *BridgeCrossCaller) AllTokenInfo(opts *bind.CallOpts) ([]IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "allTokenInfo")

	if err != nil {
		return *new([]IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeTokenInfoTokenInfo)).(*[]IBridgeTokenInfoTokenInfo)

	return out0, err

}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeCross *BridgeCrossSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.AllTokenInfo(&_BridgeCross.CallOpts)
}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeCross *BridgeCrossCallerSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.AllTokenInfo(&_BridgeCross.CallOpts)
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
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeCross *BridgeCrossCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "calculate", token, value)

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
func (_BridgeCross *BridgeCrossSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeCross.Contract.Calculate(&_BridgeCross.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeCross *BridgeCrossCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
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

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossCaller) GetTokenInfo(opts *bind.CallOpts, token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "getTokenInfo", token)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.GetTokenInfo(&_BridgeCross.CallOpts, token)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossCallerSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.GetTokenInfo(&_BridgeCross.CallOpts, token)
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

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossCaller) TokenInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "tokenInfoByIndex", index)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.TokenInfoByIndex(&_BridgeCross.CallOpts, index)
}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeCross *BridgeCrossCallerSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeCross.Contract.TokenInfoByIndex(&_BridgeCross.CallOpts, index)
}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCaller) TokenInfoLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeCross.contract.Call(opts, &out, "tokenInfoLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeCross *BridgeCrossSession) TokenInfoLength() (*big.Int, error) {
	return _BridgeCross.Contract.TokenInfoLength(&_BridgeCross.CallOpts)
}

// TokenInfoLength is a free data retrieval call binding the contract method 0x7cfed602.
//
// Solidity: function tokenInfoLength() view returns(uint256)
func (_BridgeCross *BridgeCrossCallerSession) TokenInfoLength() (*big.Int, error) {
	return _BridgeCross.Contract.TokenInfoLength(&_BridgeCross.CallOpts)
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
