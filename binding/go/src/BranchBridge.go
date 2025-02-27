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

// BranchBridgeMetaData contains all meta data concerning the BranchBridge contract.
var BranchBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIChainManager.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIChainManager.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIChainManager.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossMintableERC20FactoryCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIChainManager.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"setChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"ChainSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibAleadyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errChainLibCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errChainLibDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errChainLibInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errChainLibNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errChainManagerChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errChainManagerChainNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errChainManagerInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"errChainManagerTokenFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainManagerTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errStandardBridgeFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"5b605f5c": "allTokenPairs(uint256)",
		"f30589c3": "allValidators()",
		"47666cb1": "bridgeFeeStation()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"8f517c17": "crossMintableERC20Factory()",
		"96ce0795": "denominator()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"ae766389": "estimateFee(uint256,address,uint256)",
		"15eb39ef": "finalizeBridge(uint256,uint256,address,address,uint256,bytes,uint8[],bytes32[],bytes32[])",
		"7c7505dc": "finalizeBridgeBatch(uint256,(uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"814914b5": "getTokenPair(uint256,address)",
		"607375a1": "initialize(uint8,address,address,address)",
		"91cf6d3e": "initializedAt()",
		"facd743b": "isValidator(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"d2ff130d": "pauseToken(uint256,address)",
		"5c975abb": "paused()",
		"7f260a75": "permitBridgeToken(uint256,address,address,address,uint256,uint256,uint256,(address,address,uint256,uint256,uint8,bytes32,bytes32),bytes)",
		"52d1902d": "proxiableUUID()",
		"edbbea39": "registerToken(uint256,bool,address,address)",
		"d7c82f32": "removeFeeStation()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
		"3960e787": "retryFinalizeBridge(uint256,uint256)",
		"030372c3": "retryFinalizeBridgeBatch(uint256,uint256[])",
		"952a95de": "revertedArguments(uint256,uint256)",
		"8415a385": "revertedReason(uint256,uint256)",
		"fb75b2c7": "rewardWallet()",
		"1b44fd15": "setChain(uint256)",
		"54db0126": "setFeeStation(address)",
		"5958621e": "setRewardWallet(address)",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"84d58d42": "unpauseToken(uint256,address)",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516154016100f95f395f818161204d0152818161207601526121ba01526154015ff3fe6080604052600436106102c2575f3560e01c80638456cb591161016f578063aed1d403116100d8578063edbbea3911610092578063f45096431161006d578063f45096431461093c578063f698da251461095b578063facd743b1461096f578063fb75b2c71461098e575f5ffd5b8063edbbea39146108dd578063f2fde38b146108fc578063f30589c31461091b575f5ffd5b8063aed1d40314610839578063b7f3358d1461084d578063cbae59581461086c578063d2ff130d1461088b578063d5717fc5146108aa578063d7c82f32146108c9575f5ffd5b80639300c926116101295780639300c92614610751578063952a95de1461077057806396ce07951461079c578063ad3cb1cc146107b0578063ae6893f8146107e0578063ae766389146107ff575f5ffd5b80638456cb591461068957806384b0196e1461069d57806384d58d42146106c45780638da5cb5b146106e35780638f517c171461071f57806391cf6d3e1461073d575f5ffd5b806354db01261161022b578063670e6268116101e55780637c7505dc116101c05780637c7505dc146105705780637f260a7514610583578063814914b5146105965780638415a3851461065d575f5ffd5b8063670e62681461051e5780637101fcd31461053d578063715018a61461055c575f5ffd5b806354db01261461045f5780635958621e1461047e5780635b605f5c1461049d5780635c975abb146104c95780635fd262de146104ec578063607375a1146104ff575f5ffd5b80633f4ba83a1161027c5780633f4ba83a1461039f57806340a141ff146103b357806342cde4e8146103d257806347666cb1146103f35780634f1ef2861461042a57806352d1902d1461043d575f5ffd5b8063030372c3146102dc5780631327d3d81461031057806315eb39ef1461032f5780631b44fd15146103425780631d40f0d8146103615780633960e78714610380575f5ffd5b366102d857345f036102d6576102d661418e565b005b5f5ffd5b3480156102e7575f5ffd5b506102fb6102f6366004614229565b6109ab565b60405190151581526020015b60405180910390f35b34801561031b575f5ffd5b506102d661032a3660046142f5565b6109ef565b6102fb61033d36600461443c565b6109fd565b34801561034d575f5ffd5b506102d661035c366004614533565b610c76565b34801561036c575f5ffd5b506102d661037b36600461454a565b610caf565b34801561038b575f5ffd5b506102fb61039a3660046145ef565b610ce5565b3480156103aa575f5ffd5b506102d6610dcf565b3480156103be575f5ffd5b506102d66103cd3660046142f5565b610de1565b3480156103dd575f5ffd5b5060355460405160ff9091168152602001610307565b3480156103fe575f5ffd5b50606654610412906001600160a01b031681565b6040516001600160a01b039091168152602001610307565b6102d6610438366004614678565b610deb565b348015610448575f5ffd5b50610451610e06565b604051908152602001610307565b34801561046a575f5ffd5b506102d66104793660046142f5565b610e22565b348015610489575f5ffd5b506102d66104983660046142f5565b610e96565b3480156104a8575f5ffd5b506104bc6104b7366004614533565b610f06565b6040516103079190614716565b3480156104d4575f5ffd5b505f51602061538c5f395f51905f525460ff166102fb565b6102fb6104fa366004614763565b61105f565b34801561050a575f5ffd5b506102d66105193660046147eb565b61116e565b348015610529575f5ffd5b50610412610538366004614842565b611280565b348015610548575f5ffd5b506102d661055736600461454a565b611391565b348015610567575f5ffd5b506102d66113a7565b6102fb61057e3660046149ce565b6113b8565b6102fb610591366004614ac1565b61151a565b3480156105a1575f5ffd5b506106506105b0366004614be5565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9182526003602090815260408084206001600160a01b039384168552600901825292839020835160a08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260020154608082015290565b6040516103079190614c13565b348015610668575f5ffd5b5061067c6106773660046145ef565b6115f9565b6040516103079190614c4f565b348015610694575f5ffd5b506102d661161a565b3480156106a8575f5ffd5b506106b161162a565b6040516103079796959493929190614c61565b3480156106cf575f5ffd5b506102d66106de366004614be5565b6116d3565b3480156106ee575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610412565b34801561072a575f5ffd5b505f54610412906001600160a01b031681565b348015610748575f5ffd5b50606854610451565b34801561075c575f5ffd5b506102d661076b36600461454a565b61172b565b34801561077b575f5ffd5b5061078f61078a3660046145ef565b611762565b6040516103079190614cf7565b3480156107a7575f5ffd5b50610451611781565b3480156107bb575f5ffd5b5061067c604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156107eb575f5ffd5b506104516107fa366004614533565b6117f1565b34801561080a575f5ffd5b5061081e610819366004614d49565b611807565b60408051938452602084019290925290820152606001610307565b348015610844575f5ffd5b50610451611896565b348015610858575f5ffd5b506102d6610867366004614d7e565b6118a1565b348015610877575f5ffd5b50610412610886366004614533565b6118f2565b348015610896575f5ffd5b506102d66108a5366004614be5565b6118fe565b3480156108b5575f5ffd5b506104516108c4366004614533565b611956565b3480156108d4575f5ffd5b506102d661196c565b3480156108e8575f5ffd5b506102d66108f7366004614da4565b611986565b348015610907575f5ffd5b506102d66109163660046142f5565b611a2b565b348015610926575f5ffd5b5061092f611a65565b6040516103079190614dc9565b348015610947575f5ffd5b506102d6610956366004614be5565b611a71565b348015610966575f5ffd5b50610451611af7565b34801561097a575f5ffd5b506102fb6109893660046142f5565b611b00565b348015610999575f5ffd5b506067546001600160a01b0316610412565b5f805b82518110156109e3576109da848483815181106109cd576109cd614e09565b6020026020010151610ce5565b506001016109ae565b50600190505b92915050565b6109fa816001611b0c565b50565b5f610a06611bd5565b5f8b81526003602052604090208b908a90610a249060070182611c05565b8190610a545760405163c4bfa74960e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff1615610ab35760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b50610abc611c26565b5f348015610ae65760405163341382c560e11b815260048101929092526024820152604401610a4b565b5050610af18d611956565b8c14610afc8e611956565b8d9091610b25576040516329846bd360e21b815260048101929092526024820152604401610a4b565b5050610b857f8ee949cbdad2722efdf3806ebb3e900a8822e1e18aea7c94f0ce625ff090b6d18d8d8d8d8d8d604051602001610b679796959493929190614e45565b60405160208183030381529060405280519060200120878787611c5d565b5f8d81526003602052604081206002018054600101905580610ba98f8e8e8e611e58565b915091508115610c0f578b6001600160a01b03168d6001600160a01b03168f7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738e42604051610c02929190918252602082015260400190565b60405180910390a4610c4a565b610c1f8f8f8f8f8f8f8f88611ed5565b6040518e907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610c6660015f5160206153ac5f395f51905f5255565b50509a9950505050505050505050565b610c7e611f6c565b610c89600182611fc7565b8190610cab5760405163021aa6c760e21b8152600401610a4b91815260200190565b5050565b5f5b8151811015610cab57610cdd828281518110610ccf57610ccf614e09565b60200260200101515f611b0c565b600101610cb1565b5f610cee611bd5565b610cf6611c26565b5f610d018484611762565b90505f5f610d1d86846020015185604001518660600151611e58565b91509150818190610d415760405162461bcd60e51b8152600401610a4b9190614c4f565b50610d4c8686611fd2565b82604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373866060015142604051610daa929190918252602082015260400190565b60405180910390a4600193505050506109e960015f5160206153ac5f395f51905f5255565b610dd7611f6c565b610ddf611fe9565b565b6109fa815f611b0c565b610df3612042565b610dfc826120e6565b610cab82826120ee565b5f610e0f6121af565b505f51602061536c5f395f51905f525b90565b610e2a611f6c565b6001600160a01b038116610e74576040516211211f60e21b81526020600482015260116024820152702fb13934b233b2a332b2a9ba30ba34b7b760791b6044820152606401610a4b565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b610e9e611f6c565b6001600160a01b038116610ee4576040516211211f60e21b815260206004820152600d60248201526c72657761726457616c6c65745f60981b6044820152606401610a4b565b606780546001600160a01b0319166001600160a01b0392909216919091179055565b5f818152600360205260408120606091610f22600783016121f8565b90505f81516001600160401b03811115610f3e57610f3e6141a2565b604051908082528060200260200182016040528015610f9557816020015b6040805160a0810182525f808252602080830182905292820181905260608201819052608082015282525f19909201910181610f5c5790505b5090505f5b825181101561105657836009015f848381518110610fba57610fba614e09565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160a08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b909304909116151560608201526002909101546080820152825183908390811061104357611043614e09565b6020908102919091010152600101610f9a565b50949350505050565b5f611068611bd5565b5f898152600360205260409020899089906110869060070182611c05565b81906110b15760405163c4bfa74960e01b81526001600160a01b039091166004820152602401610a4b565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156111105760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b50611119611c26565b5f5f6111288d8d8c8c8c612204565b915091506111448d8d6111383390565b8e8e87875f8f8f61231d565b60019450505061116060015f5160206153ac5f395f51905f5255565b505098975050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156111b25750825b90505f826001600160401b031660011480156111cd5750303b155b9050811580156111db575080155b156111f95760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561122357845460ff60401b1916600160401b1785555b61122f8989898961236b565b831561127557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f611289611f6c565b5f546001600160a01b03166112b1576040516380cf12ed60e01b815260040160405180910390fd5b5f546040516bffffffffffffffffffffffff19606087901b1660208201526001600160a01b0390911690634804a04190603401604051602081830303815290604052805190602001208560405160200161130b9190614ea9565b60405160208183030381529060405286866040518563ffffffff1660e01b815260040161133b9493929190614eca565b6020604051808303815f875af1158015611357573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061137b9190614f09565b9050611389855f8387611986565b949350505050565b61139e61037b60336121f8565b6109fa8161172b565b6113af611f6c565b610ddf5f61243f565b5f805b8581101561150c57611503888888848181106113d9576113d9614e09565b90506020028101906113eb9190614f24565b358989858181106113fe576113fe614e09565b90506020028101906114109190614f24565b6114219060408101906020016142f5565b8a8a8681811061143357611433614e09565b90506020028101906114459190614f24565b6114569060608101906040016142f5565b8b8b8781811061146857611468614e09565b905060200281019061147a9190614f24565b606001358c8c8881811061149057611490614e09565b90506020028101906114a29190614f24565b6114b0906080810190614f42565b8c89815181106114c2576114c2614e09565b60200260200101518c8a815181106114dc576114dc614e09565b60200260200101518c8b815181106114f6576114f6614e09565b60200260200101516109fd565b506001016113bb565b506001979650505050505050565b5f611523611bd5565b5f8b81526003602052604090208b908b906115419060070182611c05565b819061156c5760405163c4bfa74960e01b81526001600160a01b039091166004820152602401610a4b565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156115cb5760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b506115d4611c26565b5f5f6115e38f8f8d8d8d612204565b91509150610c4a8f8f8f8f8f87878f8f8f6124af565b5f828152600360205260409020606090611613908361251a565b9392505050565b611622611f6c565b610ddf6125bc565b5f60608082808083815f51602061534c5f395f51905f52805490915015801561165557506001810154155b6116995760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a4b565b6116a1612604565b6116a96126c4565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6116db611f6c565b5f8281526003602052604090206116f29082612702565b6040516001600160a01b0382169083907fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9905f90a35050565b5f5b8151811015610cab5761175a82828151811061174b5761174b614e09565b60200260200101516001611b0c565b60010161172d565b61176a614107565b5f83815260036020526040902061161390836127b7565b606654604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa1580156117c8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117ec9190614f84565b905090565b5f8181526003602052604081206109e9906128ae565b60665460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa158015611863573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118879190614f9b565b91989097509095509350505050565b5f6117ec60336128c0565b6118a9611f6c565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f6109e96033836128c9565b611906611f6c565b5f82815260036020526040902061191d90826128d4565b6040516001600160a01b0382169083907ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3905f90a35050565b5f8181526003602052604081206109e990612990565b611974611f6c565b606680546001600160a01b0319169055565b61198e611f6c565b6119996001856129a2565b84906119bb57604051636292216560e11b8152600401610a4b91815260200190565b505f8481526003602052604090206119d5908484846129b9565b806001600160a01b0316826001600160a01b0316857f78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c86604051611a1d911515815260200190565b60405180910390a450505050565b611a33611f6c565b6001600160a01b038116611a5c57604051631e4fbdf760e01b81525f6004820152602401610a4b565b6109fa8161243f565b60606117ec60336121f8565b611a79611f6c565b611a846001836129a2565b8290611aa657604051636292216560e11b8152600401610a4b91815260200190565b505f828152600360205260409020611abe9082612ae5565b6040516001600160a01b0382169083907fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d905f90a35050565b5f6117ec612ba0565b5f6109e9603383611c05565b611b14611f6c565b8015611b5657611b25603383612ba9565b8290611b50576040516329a04e7760e21b81526001600160a01b039091166004820152602401610a4b565b50611b8e565b611b61603383612bbd565b8290611b8c5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610a4b565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f51602061538c5f395f51905f525460ff1615610ddf5760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f9081526001830160205260408120541515611613565b5f5160206153ac5f395f51905f52805460011901611c5757604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b8251825181148015611c6f5750815181145b835183518392611ca3576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610a4b565b505060355482915060ff16811015611cd157604051632fcba65760e11b8152600401610a4b91815260200190565b505f80826001600160401b03811115611cec57611cec6141a2565b604051908082528060200260200182016040528015611d15578160200160208202803683370190505b5090505f5b83811015611e22575f611d85888381518110611d3857611d38614e09565b6020026020010151888481518110611d5257611d52614e09565b6020026020010151888581518110611d6c57611d6c614e09565b6020026020010151611d7d8d612bd1565b929190612bfd565b9050611d9081611b00565b8190611dbb5760405163845a09e760e01b81526001600160a01b039091166004820152602401610a4b565b505f805b8451811015611e0b57848181518110611dda57611dda614e09565b60200260200101516001600160a01b0316836001600160a01b031603611e035760019150611e0b565b600101611dbf565b5080611e18578460010194505b5050600101611d1a565b50603554829060ff16811015611e4e57604051632fcba65760e11b8152600401610a4b91815260200190565b5050505050505050565b5f60605f196001600160a01b03861601611e9757611e7f6001600160a01b03851684612c29565b505060408051602081019091525f8152600190611ecc565b8215611ecc57611ea78686612cbb565b15611ec157611eb886868686612cf3565b91509150611ecc565b611eb8858585612ea5565b94509492505050565b611e4e6040518060a00160405280898152602001886001600160a01b03168152602001876001600160a01b0316815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509390945250508b81526003602052604090209190508361304c565b60015f5160206153ac5f395f51905f5255565b33611f9e7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610ddf5760405163118cdaa760e01b8152336004820152602401610a4b565b5f6116138383613111565b5f828152600360205260409020610cab908261315d565b611ff16131ec565b5f51602061538c5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020016118e7565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806120c857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166120bc5f51602061536c5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610ddf5760405163703e46dd60e11b815260040160405180910390fd5b6109fa611f6c565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612148575060408051601f3d908101601f1916820190925261214591810190614f84565b60015b61217057604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a4b565b5f51602061536c5f395f51905f5281146121a057604051632a87526960e21b815260048101829052602401610a4b565b6121aa838361321b565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ddf5760405163703e46dd60e11b815260040160405180910390fd5b60605f61161383613270565b6066545f9081906001600160a01b031661222257505f905080612313565b60665460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa158015612278573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061229c9190614f9b565b909450925090508086108015906122b35750828510155b80156122bf5750818410155b81848489898990919293949561230b5760405163769296ad60e01b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a4b565b505050505050505b9550959350505050565b6123338a8a8a8961232e898b614fda565b6132c9565b6123458a8a8a8a8a8a8a8a8a8a61346f565b5f8a81526003602052604090206001908101805490910190555b50505050505050505050565b612373613567565b61237c336135b0565b612385846135c1565b61238e82613627565b612396613765565b61239e61376d565b6123a661377d565b6001600160a01b038116156123d157606680546001600160a01b0319166001600160a01b0383161790555b6001600160a01b038316612416576040516211211f60e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a4b565b5050606780546001600160a01b0319166001600160a01b03929092169190911790555043606855565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b836124ba8688614fda565b6124c49190614fda565b83604001511015898790916124fc576040516224b63960e71b81526001600160a01b0390921660048301526024820152604401610a4b565b50506125078361378d565b61235f8a8a8a8a8a8a8a60018a8a61231d565b5f818152600483016020526040902080546060919061253890614fed565b80601f016020809104026020016040519081016040528092919081815260200182805461256490614fed565b80156125af5780601f10612586576101008083540402835291602001916125af565b820191905f5260205f20905b81548152906001019060200180831161259257829003601f168201915b5050505050905092915050565b6125c4611bd5565b5f51602061538c5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361202a565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061534c5f395f51905f529161264290614fed565b80601f016020809104026020016040519081016040528092919081815260200182805461266e90614fed565b80156126b95780601f10612690576101008083540402835291602001916126b9565b820191905f5260205f20905b81548152906001019060200180831161269c57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061534c5f395f51905f529161264290614fed565b61270f6007830182611c05565b819061273a57604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b0381165f9081526009830160205260409020600101548190600160a81b900460ff1661278d57604051630196157d60e31b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f908152600990910160205260409020600101805460ff60a81b19169055565b6127bf614107565b5f82815260038085016020908152604092839020835160a0810185528154815260018201546001600160a01b039081169382019390935260028201549092169382019390935290820154606082015260048201805491929160808401919061282690614fed565b80601f016020809104026020016040519081016040528092919081815260200182805461285290614fed565b801561289d5780601f106128745761010080835404028352916020019161289d565b820191905f5260205f20905b81548152906001019060200180831161288057829003601f168201915b505050505081525050905092915050565b6001808201545f916109e99190614fda565b5f6109e9825490565b5f611613838361388a565b6128e16007830182611c05565b819061290c57604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b0381165f9081526009830160205260409020600101548190600160a81b900460ff161561296057604051635e31435560e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f908152600990910160205260409020600101805460ff60a81b1916600160a81b179055565b60028101545f906109e9906001614fda565b5f8181526001830160205260408120541515611613565b6001600160a01b0382166129fc5760405162d9127960e71b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a4b565b612a096007850183612ba9565b8290612a3457604051632f644c4360e21b81526001600160a01b039091166004820152602401610a4b565b506040805160a0810182526001600160a01b0393841680825292841660208083019182529515158284019081525f6060840181815260808501828152968252600990990190975292909520905181549085166001600160a01b03199091161781559351600185018054925196511515600160a81b0260ff60a81b19971515600160a01b026001600160a81b031990941692909516919091179190911794909416919091179092559051600290910155565b6001600160a01b038116612b285760405162d9127960e71b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a4b565b612b356007830182612bbd565b8190612b6057604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f90815260099091016020526040812080546001600160a01b03191681556001810180546001600160b01b031916905560020155565b5f6117ec6138b0565b5f611613836001600160a01b038416613111565b5f611613836001600160a01b038416613923565b5f6109e9612bdd612ba0565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612c0d88888888613a06565b925092509250612c1d8282613ace565b50909695505050505050565b80471015612c535760405163cf47918160e01b815247600482015260248101829052604401610a4b565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114612c9d576040519150601f19603f3d011682016040523d82523d5f602084013e612ca2565b606091505b509150915081612cb557612cb581613b86565b50505050565b5f9182526003602090815260408084206001600160a01b0393909316845260099092019052902060010154600160a01b900460ff1690565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015612d62575060408051601f3d908101601f19168201909252612d5f91810190615025565b60015b612e3457612d6e615040565b806308c379a003612d975750612d82615058565b80612d8d5750612df1565b5f92509050611ecc565b634e487b7103612df157612da96150da565b90612db45750612df1565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b604051602081830303815290604052915050611ecc565b3d808015612e1a576040519150601f19603f3d011682016040523d82523d5f602084013e612e1f565b606091505b505f925080604051602001612dda91906150f7565b8015612e5f576001925060405180602001604052805f8152509150612e5a878786613baf565b612e9b565b5f92506040518060400160405280601f81526020017f5374616e646172644272696467653a207472616e73666572206661696c65640081525091505b5094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015612f14575060408051601f3d908101601f19168201909252612f1191810190615025565b60015b612fe657612f20615040565b806308c379a003612f495750612f34615058565b80612f3f5750612fa3565b5f92509050613044565b634e487b7103612fa357612f5b6150da565b90612f665750612fa3565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b604051602081830303815290604052915050613044565b3d808015612fcc576040519150601f19603f3d011682016040523d82523d5f602084013e612fd1565b606091505b505f925080604051602001612f8c91906150f7565b8015613006576001925060405180602001604052805f8152509150613042565b5f92506040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b815161305b6005850182611fc7565b819061307d5760405163980242c960e01b8152600401610a4b91815260200190565b505f8181526003858101602090815260409283902086518155908601516001820180546001600160a01b03199081166001600160a01b039384161790915593870151600283018054909516911617909255606085015190820155608084015184919060048201906130ee9082615160565b5050505f818152600485016020526040902061310a8382615160565b5050505050565b5f81815260018301602052604081205461315657508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109e9565b505f6109e9565b61316a6005830182613bc7565b819061318c57604051632240947160e21b8152600401610a4b91815260200190565b505f81815260048301602052604081206131a591614144565b5f8181526003808401602052604082208281556001810180546001600160a01b0319908116909155600282018054909116905590810182905590612cb56004830182614144565b5f51602061538c5f395f51905f525460ff16610ddf57604051638dfc202b60e01b815260040160405180910390fd5b61322482613bd2565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613268576121aa8282613c35565b610cab613ca7565b6060815f018054806020026020016040519081016040528092919081815260200182805480156132bd57602002820191905f5260205f20905b8154815260200190600101908083116132a9575b50505050509050919050565b5f196001600160a01b0385160161333c576132e48183614fda565b34146132f08284614fda565b3490916133195760405163341382c560e11b815260048101929092526024820152604401610a4b565b5050801561333757606754613337906001600160a01b031682612c29565b61310a565b5f3480156133665760405163341382c560e11b815260048101929092526024820152604401610a4b565b5061338a905083306133788486614fda565b6001600160a01b038816929190613cc6565b80156133aa576067546133aa906001600160a01b03868116911683613d2d565b6133b48585612cbb565b156133c457613337858584613d5e565b604051632770a7eb60e21b8152306004820152602481018390526001600160a01b03851690639dc29fac906044016020604051808303815f875af115801561340e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134329190615025565b848484909192611e4e5760405163bc8ecf7b60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610a4b565b5f5f61347a8c6117f1565b91506134af8c8c5f9182526003602090815260408084206001600160a01b039384168552600901909152909120600101541690565b9050896001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d78e858e8e428d8d8d6040516134fa98979695949392919061521a565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a604051613551929190918252602082015260400190565b60405180910390a4505050505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610ddf57604051631afcd79f60e31b815260040160405180910390fd5b6135b8613567565b6109fa81613d76565b6135c9613567565b613611604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613d7e565b6035805460ff191660ff92909216919091179055565b6001600160a01b0381166136385750565b5f816001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015613674573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261369b9190810190615272565b604080513060208201520160408051601f19818403018152908290526136c492916020016152f1565b60408051601f1981840301815282825246602084015292505f91613701918391016040516020818303038152906040528051906020012084613d90565b90506001600160a01b0381166137425760405162461bcd60e51b8152600401610a4b906020808252600490820152637a65726f60e01b604082015260600190565b5f80546001600160a01b0319166001600160a01b03929092169190911790555050565b610ddf613567565b613775613567565b610ddf613e24565b613785613567565b610ddf613e44565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180613839576040513d5f823e3d81fd5b50505f513d9150811561385057806001141561385e565b84516001600160a01b03163b155b1561310a57845160405163733eb19560e01b81526001600160a01b039091166004820152602401610a4b565b5f825f01828154811061389f5761389f614e09565b905f5260205f200154905092915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6138da613e4c565b6138e2613eb4565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f81815260018301602052604081205480156139fd575f613945600183615305565b85549091505f9061395890600190615305565b90508082146139b7575f865f01828154811061397657613976614e09565b905f5260205f200154905080875f01848154811061399657613996614e09565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806139c8576139c8615318565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109e9565b5f9150506109e9565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613a3f57505f91506003905082613ac4565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613a90573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613abb57505f925060019150829050613ac4565b92505f91508190505b9450945094915050565b5f826003811115613ae157613ae161532c565b03613aea575050565b6001826003811115613afe57613afe61532c565b03613b1c5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613b3057613b3061532c565b03613b515760405163fce698f760e01b815260048101829052602401610a4b565b6003826003811115613b6557613b6561532c565b03610cab576040516335e2f38360e21b815260048101829052602401610a4b565b805115613b965780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f8381526003602052604090206121aa908383613ef6565b5f6116138383613923565b806001600160a01b03163b5f03613c0757604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a4b565b5f51602061536c5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613c519190615340565b5f60405180830381855af49150503d805f8114613c89576040519150601f19603f3d011682016040523d82523d5f602084013e613c8e565b606091505b5091509150613c9e858383613f86565b95945050505050565b3415610ddf5760405163b398979f60e01b815260040160405180910390fd5b6040516001600160a01b038481166024830152838116604483015260648201839052612cb59186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613fe2565b6040516001600160a01b038381166024830152604482018390526121aa91859182169063a9059cbb90606401613cfb565b5f8381526003602052604090206121aa90838361404e565b611a33613567565b613d86613567565b610cab8282614084565b5f83471015613dbb5760405163cf47918160e01b815247600482015260248101859052604401610a4b565b81515f03613ddc57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615613dfd576040513d5f823e3d81fd5b6001600160a01b0381166116135760405163b06ebf3d60e01b815260040160405180910390fd5b613e2c613567565b5f51602061538c5f395f51905f52805460ff19169055565b611f59613567565b5f5f51602061534c5f395f51905f5281613e64612604565b805190915015613e7c57805160209091012092915050565b81548015613e8b579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061534c5f395f51905f5281613ecc6126c4565b805190915015613ee457805160209091012092915050565b60018201548015613e8b579392505050565b6001600160a01b0382165f9081526009840160205260408120600201548190613f1f90846140e3565b86549193509150848484613f5f576040516302ea38e160e51b815260048101939093526001600160a01b0390911660248301526044820152606401610a4b565b5050506001600160a01b039093165f90815260099094016020525050604090912060020155565b606082613f9b57613f9682613b86565b611613565b8151158015613fb257506001600160a01b0384163b155b15613fdb57604051639996b31560e01b81526001600160a01b0385166004820152602401610a4b565b5080611613565b5f5f60205f8451602086015f885af180614001576040513d5f823e3d81fd5b50505f513d91508115614018578060011415614025565b6001600160a01b0384163b155b15612cb557604051635274afe760e01b81526001600160a01b0385166004820152602401610a4b565b6001600160a01b0382165f9081526009840160205260408120600201805483929061407a908490614fda565b9091555050505050565b61408c613567565b5f51602061534c5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026140c58482615160565b50600381016140d48382615160565b505f8082556001909101555050565b5f5f838311156140f757505f905080614100565b50600190508183035b9250929050565b6040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b50805461415090614fed565b5f825580601f1061415f575050565b601f0160209004905f5260205f20908101906109fa91905b8082111561418a575f8155600101614177565b5090565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b03821117156141d5576141d56141a2565b60405250565b601f8201601f191681016001600160401b0381118282101715614200576142006141a2565b6040525050565b5f6001600160401b0382111561421f5761421f6141a2565b5060051b60200190565b5f5f6040838503121561423a575f5ffd5b8235915060208301356001600160401b03811115614256575f5ffd5b8301601f81018513614266575f5ffd5b803561427181614207565b60405161427e82826141db565b80915082815260208101915060208360051b8501019250878311156142a1575f5ffd5b6020840193505b828410156142c35783358252602093840193909101906142a8565b809450505050509250929050565b6001600160a01b03811681146109fa575f5ffd5b80356142f0816142d1565b919050565b5f60208284031215614305575f5ffd5b8135611613816142d1565b5f5f83601f840112614320575f5ffd5b5081356001600160401b03811115614336575f5ffd5b602083019150836020828501011115614100575f5ffd5b803560ff811681146142f0575f5ffd5b5f82601f83011261436c575f5ffd5b813561437781614207565b60405161438482826141db565b80915082815260208101915060208360051b8601019250858311156143a7575f5ffd5b602085015b838110156143cb576143bd8161434d565b8352602092830192016143ac565b5095945050505050565b5f82601f8301126143e4575f5ffd5b81356143ef81614207565b6040516143fc82826141db565b80915082815260208101915060208360051b86010192508583111561441f575f5ffd5b602085015b838110156143cb578035835260209283019201614424565b5f5f5f5f5f5f5f5f5f5f6101208b8d031215614456575f5ffd5b8a35995060208b0135985061446d60408c016142e5565b975061447b60608c016142e5565b965060808b0135955060a08b01356001600160401b0381111561449c575f5ffd5b6144a88d828e01614310565b90965094505060c08b01356001600160401b038111156144c6575f5ffd5b6144d28d828e0161435d565b93505060e08b01356001600160401b038111156144ed575f5ffd5b6144f98d828e016143d5565b9250506101008b01356001600160401b03811115614515575f5ffd5b6145218d828e016143d5565b9150509295989b9194979a5092959850565b5f60208284031215614543575f5ffd5b5035919050565b5f6020828403121561455a575f5ffd5b81356001600160401b0381111561456f575f5ffd5b8201601f8101841361457f575f5ffd5b803561458a81614207565b60405161459782826141db565b80915082815260208101915060208360051b8501019250868311156145ba575f5ffd5b6020840193505b828410156145e55783356145d4816142d1565b8252602093840193909101906145c1565b9695505050505050565b5f5f60408385031215614600575f5ffd5b50508035926020909101359150565b5f6001600160401b03821115614627576146276141a2565b50601f01601f191660200190565b5f61463f8361460f565b60405161464c82826141db565b809250848152858585011115614660575f5ffd5b848460208301375f6020868301015250509392505050565b5f5f60408385031215614689575f5ffd5b8235614694816142d1565b915060208301356001600160401b038111156146ae575f5ffd5b8301601f810185136146be575f5ffd5b6146cd85823560208401614635565b9150509250929050565b80516001600160a01b03908116835260208083015190911690830152604080820151151590830152606080820151151590830152608090810151910152565b602080825282518282018190525f918401906040840190835b81811015614758576147428385516146d7565b6020939093019260a0929092019160010161472f565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b03121561477a575f5ffd5b88359750602089013561478c816142d1565b9650604089013561479c816142d1565b9550606089013594506080890135935060a0890135925060c08901356001600160401b038111156147cb575f5ffd5b6147d78b828c01614310565b999c989b5096995094979396929594505050565b5f5f5f5f608085870312156147fe575f5ffd5b6148078561434d565b93506020850135614817816142d1565b92506040850135614827816142d1565b91506060850135614837816142d1565b939692955090935050565b5f5f5f5f60808587031215614855575f5ffd5b843593506020850135614867816142d1565b925060408501356001600160401b03811115614881575f5ffd5b8501601f81018713614891575f5ffd5b6148a087823560208401614635565b9250506148af6060860161434d565b905092959194509250565b5f82601f8301126148c9575f5ffd5b81356148d481614207565b6040516148e182826141db565b80915082815260208101915060208360051b860101925085831115614904575f5ffd5b602085015b838110156143cb5780356001600160401b03811115614926575f5ffd5b614935886020838a010161435d565b84525060209283019201614909565b5f82601f830112614953575f5ffd5b813561495e81614207565b60405161496b82826141db565b80915082815260208101915060208360051b86010192508583111561498e575f5ffd5b602085015b838110156143cb5780356001600160401b038111156149b0575f5ffd5b6149bf886020838a01016143d5565b84525060209283019201614993565b5f5f5f5f5f5f60a087890312156149e3575f5ffd5b8635955060208701356001600160401b038111156149ff575f5ffd5b8701601f81018913614a0f575f5ffd5b80356001600160401b03811115614a24575f5ffd5b8960208260051b8401011115614a38575f5ffd5b6020919091019550935060408701356001600160401b03811115614a5a575f5ffd5b614a6689828a016148ba565b93505060608701356001600160401b03811115614a81575f5ffd5b614a8d89828a01614944565b92505060808701356001600160401b03811115614aa8575f5ffd5b614ab489828a01614944565b9150509295509295509295565b5f5f5f5f5f5f5f5f5f5f8a8c036101e0811215614adc575f5ffd5b8b359a5060208c0135614aee816142d1565b995060408c0135614afe816142d1565b985060608c0135614b0e816142d1565b975060808c0135965060a08c0135955060c08c0135945060e060df1982011215614b36575f5ffd5b50604051614b43816141b6565b60e08c0135614b51816142d1565b81526101008c0135614b62816142d1565b60208201526101208c013560408201526101408c01356060820152614b8a6101608d0161434d565b60808201526101808c013560a08201526101a08c013560c082015292506101c08b01356001600160401b03811115614bc0575f5ffd5b614bcc8d828e01614310565b915080935050809150509295989b9194979a5092959850565b5f5f60408385031215614bf6575f5ffd5b823591506020830135614c08816142d1565b809150509250929050565b60a081016109e982846146d7565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6116136020830184614c21565b60ff60f81b8816815260e060208201525f614c7f60e0830189614c21565b8281036040840152614c918189614c21565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614ce6578351835260209384019390920191600101614cc8565b50909b9a5050505050505050505050565b602081528151602082015260018060a01b03602083015116604082015260018060a01b036040830151166060820152606082015160808201525f608083015160a08084015261138960c0840182614c21565b5f5f5f60608486031215614d5b575f5ffd5b833592506020840135614d6d816142d1565b929592945050506040919091013590565b5f60208284031215614d8e575f5ffd5b6116138261434d565b80151581146109fa575f5ffd5b5f5f5f5f60808587031215614db7575f5ffd5b84359350602085013561481781614d97565b602080825282518282018190525f918401906040840190835b818110156147585783516001600160a01b0316835260209384019390920191600101614de2565b634e487b7160e01b5f52603260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90614e859083018486614e1d565b9998505050505050505050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f611613600d830184614e92565b848152608060208201525f614ee26080830186614c21565b8281036040840152614ef48186614c21565b91505060ff8316606083015295945050505050565b5f60208284031215614f19575f5ffd5b8151611613816142d1565b5f8235609e19833603018112614f38575f5ffd5b9190910192915050565b5f5f8335601e19843603018112614f57575f5ffd5b8301803591506001600160401b03821115614f70575f5ffd5b602001915036819003821315614100575f5ffd5b5f60208284031215614f94575f5ffd5b5051919050565b5f5f5f60608486031215614fad575f5ffd5b5050815160208301516040909301519094929350919050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156109e9576109e9614fc6565b600181811c9082168061500157607f821691505b60208210810361501f57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215615035575f5ffd5b815161161381614d97565b5f60033d1115610e1f5760045f5f3e505f5160e01c90565b5f60443d10156150655790565b6040513d600319016004823e80513d60248201116001600160401b038211171561508e57505090565b80820180516001600160401b038111156150a9575050505090565b3d84016003190182820160200111156150c3575050505090565b6150d2602082850101856141db565b509392505050565b5f5f60233d11156150f357602060045f3e50505f516001905b9091565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f6116136011830184614e92565b601f8211156121aa57805f5260205f20601f840160051c810160208510156151415750805b601f840160051c820191505b8181101561310a575f815560010161514d565b81516001600160401b03811115615179576151796141a2565b61518d816151878454614fed565b8461511c565b6020601f8211600181146151bf575f83156151a85750848201515b5f19600385901b1c1916600184901b17845561310a565b5f84815260208120601f198516915b828110156151ee57878501518255602094850194600190920191016151ce565b508482101561520b57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f906152649083018486614e1d565b9a9950505050505050505050565b5f60208284031215615282575f5ffd5b81516001600160401b03811115615297575f5ffd5b8201601f810184136152a7575f5ffd5b80516152b28161460f565b6040516152bf82826141db565b8281528660208486010111156152d3575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b5f6113896152ff8386614e92565b84614e92565b818103818111156109e9576109e9614fc6565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f6116138284614e9256fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212204c0deb44de96340598fc4a0ee9c5ba4f4ca873916bba713eb6eb36b2a0daea4564736f6c634300081c0033",
}

// BranchBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BranchBridgeMetaData.ABI instead.
var BranchBridgeABI = BranchBridgeMetaData.ABI

// Deprecated: Use BranchBridgeMetaData.Sigs instead.
// BranchBridgeFuncSigs maps the 4-byte function signature to its string representation.
var BranchBridgeFuncSigs = BranchBridgeMetaData.Sigs

// BranchBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BranchBridgeMetaData.Bin instead.
var BranchBridgeBin = BranchBridgeMetaData.Bin

// DeployBranchBridge deploys a new Ethereum contract, binding an instance of BranchBridge to it.
func DeployBranchBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BranchBridge, error) {
	parsed, err := BranchBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BranchBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BranchBridge{BranchBridgeCaller: BranchBridgeCaller{contract: contract}, BranchBridgeTransactor: BranchBridgeTransactor{contract: contract}, BranchBridgeFilterer: BranchBridgeFilterer{contract: contract}}, nil
}

// BranchBridge is an auto generated Go binding around an Ethereum contract.
type BranchBridge struct {
	BranchBridgeCaller     // Read-only binding to the contract
	BranchBridgeTransactor // Write-only binding to the contract
	BranchBridgeFilterer   // Log filterer for contract events
}

// BranchBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BranchBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BranchBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BranchBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BranchBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BranchBridgeSession struct {
	Contract     *BranchBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BranchBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BranchBridgeCallerSession struct {
	Contract *BranchBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BranchBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BranchBridgeTransactorSession struct {
	Contract     *BranchBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BranchBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BranchBridgeRaw struct {
	Contract *BranchBridge // Generic contract binding to access the raw methods on
}

// BranchBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BranchBridgeCallerRaw struct {
	Contract *BranchBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BranchBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BranchBridgeTransactorRaw struct {
	Contract *BranchBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBranchBridge creates a new instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridge(address common.Address, backend bind.ContractBackend) (*BranchBridge, error) {
	contract, err := bindBranchBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BranchBridge{BranchBridgeCaller: BranchBridgeCaller{contract: contract}, BranchBridgeTransactor: BranchBridgeTransactor{contract: contract}, BranchBridgeFilterer: BranchBridgeFilterer{contract: contract}}, nil
}

// NewBranchBridgeCaller creates a new read-only instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeCaller(address common.Address, caller bind.ContractCaller) (*BranchBridgeCaller, error) {
	contract, err := bindBranchBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeCaller{contract: contract}, nil
}

// NewBranchBridgeTransactor creates a new write-only instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BranchBridgeTransactor, error) {
	contract, err := bindBranchBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTransactor{contract: contract}, nil
}

// NewBranchBridgeFilterer creates a new log filterer instance of BranchBridge, bound to a specific deployed contract.
func NewBranchBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BranchBridgeFilterer, error) {
	contract, err := bindBranchBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeFilterer{contract: contract}, nil
}

// bindBranchBridge binds a generic wrapper to an already deployed contract.
func bindBranchBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BranchBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BranchBridge *BranchBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BranchBridge.Contract.BranchBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BranchBridge *BranchBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.Contract.BranchBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BranchBridge *BranchBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BranchBridge.Contract.BranchBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BranchBridge *BranchBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BranchBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BranchBridge *BranchBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BranchBridge *BranchBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BranchBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BranchBridge.Contract.UPGRADEINTERFACEVERSION(&_BranchBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BranchBridge *BranchBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BranchBridge.Contract.UPGRADEINTERFACEVERSION(&_BranchBridge.CallOpts)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IChainManagerTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IChainManagerTokenPair)).(*[]IChainManagerTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeSession) AllValidators() ([]common.Address, error) {
	return _BranchBridge.Contract.AllValidators(&_BranchBridge.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_BranchBridge *BranchBridgeCallerSession) AllValidators() ([]common.Address, error) {
	return _BranchBridge.Contract.AllValidators(&_BranchBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeCaller) BridgeFeeStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "bridgeFeeStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeSession) BridgeFeeStation() (common.Address, error) {
	return _BranchBridge.Contract.BridgeFeeStation(&_BranchBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) BridgeFeeStation() (common.Address, error) {
	return _BranchBridge.Contract.BridgeFeeStation(&_BranchBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeCaller) CrossMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "crossMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeSession) CrossMintableERC20Factory() (common.Address, error) {
	return _BranchBridge.Contract.CrossMintableERC20Factory(&_BranchBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) CrossMintableERC20Factory() (common.Address, error) {
	return _BranchBridge.Contract.CrossMintableERC20Factory(&_BranchBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) Denominator() (*big.Int, error) {
	return _BranchBridge.Contract.Denominator(&_BranchBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) Denominator() (*big.Int, error) {
	return _BranchBridge.Contract.Denominator(&_BranchBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeSession) DomainSeparator() ([32]byte, error) {
	return _BranchBridge.Contract.DomainSeparator(&_BranchBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BranchBridge *BranchBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _BranchBridge.Contract.DomainSeparator(&_BranchBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BranchBridge *BranchBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_BranchBridge *BranchBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BranchBridge.Contract.Eip712Domain(&_BranchBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BranchBridge *BranchBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BranchBridge.Contract.Eip712Domain(&_BranchBridge.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumAmount *big.Int
		GasFee        *big.Int
		ExFee         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BranchBridge.Contract.EstimateFee(&_BranchBridge.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BranchBridge.Contract.EstimateFee(&_BranchBridge.CallOpts, remoteChainID, token, value)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextFinalizeIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextFinalizeIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextInitiateIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BranchBridge.Contract.GetNextInitiateIndex(&_BranchBridge.CallOpts, remoteChainID)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IChainManagerTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IChainManagerTokenPair)).(*IChainManagerTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) InitializedAt() (*big.Int, error) {
	return _BranchBridge.Contract.InitializedAt(&_BranchBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _BranchBridge.Contract.InitializedAt(&_BranchBridge.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeSession) IsValidator(validator common.Address) (bool, error) {
	return _BranchBridge.Contract.IsValidator(&_BranchBridge.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_BranchBridge *BranchBridgeCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _BranchBridge.Contract.IsValidator(&_BranchBridge.CallOpts, validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeSession) Owner() (common.Address, error) {
	return _BranchBridge.Contract.Owner(&_BranchBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) Owner() (common.Address, error) {
	return _BranchBridge.Contract.Owner(&_BranchBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeSession) Paused() (bool, error) {
	return _BranchBridge.Contract.Paused(&_BranchBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BranchBridge *BranchBridgeCallerSession) Paused() (bool, error) {
	return _BranchBridge.Contract.Paused(&_BranchBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _BranchBridge.Contract.ProxiableUUID(&_BranchBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BranchBridge *BranchBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BranchBridge.Contract.ProxiableUUID(&_BranchBridge.CallOpts)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCaller) RevertedArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "revertedArguments", remoteChainID, index)

	if err != nil {
		return *new(IChainManagerFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IChainManagerFinalizeArguments)).(*IChainManagerFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
	return _BranchBridge.Contract.RevertedArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCallerSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
	return _BranchBridge.Contract.RevertedArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCaller) RevertedReason(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (string, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "revertedReason", remoteChainID, index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.RevertedReason(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_BranchBridge *BranchBridgeCallerSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _BranchBridge.Contract.RevertedReason(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeCaller) RewardWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "rewardWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeSession) RewardWallet() (common.Address, error) {
	return _BranchBridge.Contract.RewardWallet(&_BranchBridge.CallOpts)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) RewardWallet() (common.Address, error) {
	return _BranchBridge.Contract.RewardWallet(&_BranchBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeSession) Threshold() (uint8, error) {
	return _BranchBridge.Contract.Threshold(&_BranchBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BranchBridge *BranchBridgeCallerSession) Threshold() (uint8, error) {
	return _BranchBridge.Contract.Threshold(&_BranchBridge.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BranchBridge.Contract.ValidatorByIndex(&_BranchBridge.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_BranchBridge *BranchBridgeCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _BranchBridge.Contract.ValidatorByIndex(&_BranchBridge.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeSession) ValidatorLength() (*big.Int, error) {
	return _BranchBridge.Contract.ValidatorLength(&_BranchBridge.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_BranchBridge *BranchBridgeCallerSession) ValidatorLength() (*big.Int, error) {
	return _BranchBridge.Contract.ValidatorLength(&_BranchBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) BridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "bridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.BridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.ChangeThreshold(&_BranchBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BranchBridge *BranchBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.ChangeThreshold(&_BranchBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BranchBridge *BranchBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BranchBridge.Contract.CreateToken(&_BranchBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridge", fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridge(fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridge(fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridgeBatch", fromChainID, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridgeBatch(fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, fromChainID, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridgeBatch(fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, fromChainID, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "initialize", _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeSession) Initialize(_threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Initialize is a paid mutator transaction binding the contract method 0x607375a1.
//
// Solidity: function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactorSession) Initialize(_threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.Initialize(&_BranchBridge.TransactOpts, _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeSession) Pause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Pause(&_BranchBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BranchBridge *BranchBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Pause(&_BranchBridge.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) PauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "pauseToken", remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.PauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.PauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.PermitBridgeToken(&_BranchBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BranchBridge *BranchBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BranchBridge *BranchBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RegisterToken(&_BranchBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveFeeStation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeFeeStation")
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeSession) RemoveFeeStation() (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveFeeStation(&_BranchBridge.TransactOpts)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveFeeStation() (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveFeeStation(&_BranchBridge.TransactOpts)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidator(&_BranchBridge.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidator(&_BranchBridge.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidators(&_BranchBridge.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.RemoveValidators(&_BranchBridge.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BranchBridge.Contract.RenounceOwnership(&_BranchBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BranchBridge *BranchBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BranchBridge.Contract.RenounceOwnership(&_BranchBridge.TransactOpts)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.ResetValidators(&_BranchBridge.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.ResetValidators(&_BranchBridge.TransactOpts, validators)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridge", fromChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridge(fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, fromChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridge(fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, fromChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", fromChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridgeBatch(fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, fromChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridgeBatch(fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, fromChainID, indexes)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_BranchBridge *BranchBridgeTransactor) SetChain(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setChain", remoteChainID)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_BranchBridge *BranchBridgeSession) SetChain(remoteChainID *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetChain(&_BranchBridge.TransactOpts, remoteChainID)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetChain(remoteChainID *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetChain(&_BranchBridge.TransactOpts, remoteChainID)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactor) SetFeeStation(opts *bind.TransactOpts, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setFeeStation", _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetFeeStation(&_BranchBridge.TransactOpts, _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetFeeStation(&_BranchBridge.TransactOpts, _bridgeFeeStation)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeTransactor) SetRewardWallet(opts *bind.TransactOpts, rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setRewardWallet", rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetRewardWallet(&_BranchBridge.TransactOpts, rewardWallet_)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidator(&_BranchBridge.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidator(&_BranchBridge.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidators(&_BranchBridge.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_BranchBridge *BranchBridgeTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.SetValidators(&_BranchBridge.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.TransferOwnership(&_BranchBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BranchBridge *BranchBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.TransferOwnership(&_BranchBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeSession) Unpause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Unpause(&_BranchBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BranchBridge *BranchBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _BranchBridge.Contract.Unpause(&_BranchBridge.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) UnpauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unpauseToken", remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnpauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnpauseToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnregisterToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BranchBridge *BranchBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BranchBridge.Contract.UnregisterToken(&_BranchBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.UpgradeToAndCall(&_BranchBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.UpgradeToAndCall(&_BranchBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BranchBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeSession) Receive() (*types.Transaction, error) {
	return _BranchBridge.Contract.Receive(&_BranchBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BranchBridge *BranchBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BranchBridge.Contract.Receive(&_BranchBridge.TransactOpts)
}

// BranchBridgeBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the BranchBridge contract.
type BranchBridgeBridgeFeeChargedIterator struct {
	Event *BranchBridgeBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFeeCharged)
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
		it.Event = new(BranchBridgeBridgeFeeCharged)
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
func (it *BranchBridgeBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFeeCharged represents a BridgeFeeCharged event raised by the BranchBridge contract.
type BranchBridgeBridgeFeeCharged struct {
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
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*BranchBridgeBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFeeChargedIterator{contract: _BranchBridge.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFeeCharged)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFeeCharged(log types.Log) (*BranchBridgeBridgeFeeCharged, error) {
	event := new(BranchBridgeBridgeFeeCharged)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeFinalizeRevertedIterator is returned from FilterBridgeFinalizeReverted and is used to iterate over the raw logs and unpacked data for BridgeFinalizeReverted events raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizeRevertedIterator struct {
	Event *BranchBridgeBridgeFinalizeReverted // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFinalizeReverted)
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
		it.Event = new(BranchBridgeBridgeFinalizeReverted)
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
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFinalizeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFinalizeReverted represents a BridgeFinalizeReverted event raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizeReverted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeReverted is a free log retrieval operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalizeReverted(opts *bind.FilterOpts, index []*big.Int) (*BranchBridgeBridgeFinalizeRevertedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizeRevertedIterator{contract: _BranchBridge.contract, event: "BridgeFinalizeReverted", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeReverted is a free log subscription operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalizeReverted(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalizeReverted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFinalizeReverted)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFinalizeReverted(log types.Log) (*BranchBridgeBridgeFinalizeReverted, error) {
	event := new(BranchBridgeBridgeFinalizeReverted)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BranchBridge contract.
type BranchBridgeBridgeFinalizedIterator struct {
	Event *BranchBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeFinalized)
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
		it.Event = new(BranchBridgeBridgeFinalized)
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
func (it *BranchBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeFinalized represents a BridgeFinalized event raised by the BranchBridge contract.
type BranchBridgeBridgeFinalized struct {
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
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, index []*big.Int, token []common.Address, to []common.Address) (*BranchBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizedIterator{contract: _BranchBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalized, index []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeFinalized)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeFinalized(log types.Log) (*BranchBridgeBridgeFinalized, error) {
	event := new(BranchBridgeBridgeFinalized)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BranchBridge contract.
type BranchBridgeBridgeInitiatedIterator struct {
	Event *BranchBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BranchBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeBridgeInitiated)
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
		it.Event = new(BranchBridgeBridgeInitiated)
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
func (it *BranchBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeBridgeInitiated represents a BridgeInitiated event raised by the BranchBridge contract.
type BranchBridgeBridgeInitiated struct {
	RemoteChainID *big.Int
	Index         *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Time          *big.Int
	Permit        bool
	ExtraData     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address localToken, address remoteToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, from []common.Address) (*BranchBridgeBridgeInitiatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeInitiatedIterator{contract: _BranchBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address localToken, address remoteToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeInitiated, remoteChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeBridgeInitiated)
				if err := _BranchBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address localToken, address remoteToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_BranchBridge *BranchBridgeFilterer) ParseBridgeInitiated(log types.Log) (*BranchBridgeBridgeInitiated, error) {
	event := new(BranchBridgeBridgeInitiated)
	if err := _BranchBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeChainSetIterator is returned from FilterChainSet and is used to iterate over the raw logs and unpacked data for ChainSet events raised by the BranchBridge contract.
type BranchBridgeChainSetIterator struct {
	Event *BranchBridgeChainSet // Event containing the contract specifics and raw log

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
func (it *BranchBridgeChainSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeChainSet)
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
		it.Event = new(BranchBridgeChainSet)
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
func (it *BranchBridgeChainSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeChainSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeChainSet represents a ChainSet event raised by the BranchBridge contract.
type BranchBridgeChainSet struct {
	RemoteChainID *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainSet is a free log retrieval operation binding the contract event 0x0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a.
//
// Solidity: event ChainSet(uint256 indexed remoteChainID)
func (_BranchBridge *BranchBridgeFilterer) FilterChainSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BranchBridgeChainSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ChainSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeChainSetIterator{contract: _BranchBridge.contract, event: "ChainSet", logs: logs, sub: sub}, nil
}

// WatchChainSet is a free log subscription operation binding the contract event 0x0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a.
//
// Solidity: event ChainSet(uint256 indexed remoteChainID)
func (_BranchBridge *BranchBridgeFilterer) WatchChainSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeChainSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ChainSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeChainSet)
				if err := _BranchBridge.contract.UnpackLog(event, "ChainSet", log); err != nil {
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

// ParseChainSet is a log parse operation binding the contract event 0x0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a.
//
// Solidity: event ChainSet(uint256 indexed remoteChainID)
func (_BranchBridge *BranchBridgeFilterer) ParseChainSet(log types.Log) (*BranchBridgeChainSet, error) {
	event := new(BranchBridgeChainSet)
	if err := _BranchBridge.contract.UnpackLog(event, "ChainSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BranchBridge contract.
type BranchBridgeEIP712DomainChangedIterator struct {
	Event *BranchBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BranchBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeEIP712DomainChanged)
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
		it.Event = new(BranchBridgeEIP712DomainChanged)
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
func (it *BranchBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the BranchBridge contract.
type BranchBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BranchBridge *BranchBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BranchBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeEIP712DomainChangedIterator{contract: _BranchBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BranchBridge *BranchBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BranchBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeEIP712DomainChanged)
				if err := _BranchBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BranchBridgeEIP712DomainChanged, error) {
	event := new(BranchBridgeEIP712DomainChanged)
	if err := _BranchBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BranchBridge contract.
type BranchBridgeInitializedIterator struct {
	Event *BranchBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BranchBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeInitialized)
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
		it.Event = new(BranchBridgeInitialized)
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
func (it *BranchBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeInitialized represents a Initialized event raised by the BranchBridge contract.
type BranchBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BranchBridge *BranchBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BranchBridgeInitializedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeInitializedIterator{contract: _BranchBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BranchBridge *BranchBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BranchBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeInitialized)
				if err := _BranchBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseInitialized(log types.Log) (*BranchBridgeInitialized, error) {
	event := new(BranchBridgeInitialized)
	if err := _BranchBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BranchBridge contract.
type BranchBridgeOwnershipTransferredIterator struct {
	Event *BranchBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BranchBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeOwnershipTransferred)
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
		it.Event = new(BranchBridgeOwnershipTransferred)
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
func (it *BranchBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the BranchBridge contract.
type BranchBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BranchBridge *BranchBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BranchBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeOwnershipTransferredIterator{contract: _BranchBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BranchBridge *BranchBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BranchBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeOwnershipTransferred)
				if err := _BranchBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BranchBridgeOwnershipTransferred, error) {
	event := new(BranchBridgeOwnershipTransferred)
	if err := _BranchBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BranchBridge contract.
type BranchBridgePausedIterator struct {
	Event *BranchBridgePaused // Event containing the contract specifics and raw log

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
func (it *BranchBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgePaused)
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
		it.Event = new(BranchBridgePaused)
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
func (it *BranchBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgePaused represents a Paused event raised by the BranchBridge contract.
type BranchBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BranchBridge *BranchBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BranchBridgePausedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BranchBridgePausedIterator{contract: _BranchBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BranchBridge *BranchBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BranchBridgePaused) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgePaused)
				if err := _BranchBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParsePaused(log types.Log) (*BranchBridgePaused, error) {
	event := new(BranchBridgePaused)
	if err := _BranchBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BranchBridge contract.
type BranchBridgeThresholdChangedIterator struct {
	Event *BranchBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BranchBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeThresholdChanged)
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
		it.Event = new(BranchBridgeThresholdChanged)
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
func (it *BranchBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeThresholdChanged represents a ThresholdChanged event raised by the BranchBridge contract.
type BranchBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BranchBridge *BranchBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BranchBridgeThresholdChangedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeThresholdChangedIterator{contract: _BranchBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BranchBridge *BranchBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BranchBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeThresholdChanged)
				if err := _BranchBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseThresholdChanged(log types.Log) (*BranchBridgeThresholdChanged, error) {
	event := new(BranchBridgeThresholdChanged)
	if err := _BranchBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairPausedIterator is returned from FilterTokenPairPaused and is used to iterate over the raw logs and unpacked data for TokenPairPaused events raised by the BranchBridge contract.
type BranchBridgeTokenPairPausedIterator struct {
	Event *BranchBridgeTokenPairPaused // Event containing the contract specifics and raw log

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
func (it *BranchBridgeTokenPairPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairPaused)
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
		it.Event = new(BranchBridgeTokenPairPaused)
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
func (it *BranchBridgeTokenPairPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairPaused represents a TokenPairPaused event raised by the BranchBridge contract.
type BranchBridgeTokenPairPaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairPaused is a free log retrieval operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairPaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BranchBridgeTokenPairPausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairPausedIterator{contract: _BranchBridge.contract, event: "TokenPairPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairPaused is a free log subscription operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairPaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairPaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairPaused)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
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

// ParseTokenPairPaused is a log parse operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairPaused(log types.Log) (*BranchBridgeTokenPairPaused, error) {
	event := new(BranchBridgeTokenPairPaused)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BranchBridge contract.
type BranchBridgeTokenPairRegisteredIterator struct {
	Event *BranchBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *BranchBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairRegistered)
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
		it.Event = new(BranchBridgeTokenPairRegistered)
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
func (it *BranchBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the BranchBridge contract.
type BranchBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	IsOrigin      bool
	LocalToken    common.Address
	RemoteToken   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0x78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BranchBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairRegisteredIterator{contract: _BranchBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0x78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairRegistered)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0x78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*BranchBridgeTokenPairRegistered, error) {
	event := new(BranchBridgeTokenPairRegistered)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairUnpausedIterator is returned from FilterTokenPairUnpaused and is used to iterate over the raw logs and unpacked data for TokenPairUnpaused events raised by the BranchBridge contract.
type BranchBridgeTokenPairUnpausedIterator struct {
	Event *BranchBridgeTokenPairUnpaused // Event containing the contract specifics and raw log

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
func (it *BranchBridgeTokenPairUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairUnpaused)
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
		it.Event = new(BranchBridgeTokenPairUnpaused)
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
func (it *BranchBridgeTokenPairUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairUnpaused represents a TokenPairUnpaused event raised by the BranchBridge contract.
type BranchBridgeTokenPairUnpaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnpaused is a free log retrieval operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairUnpaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BranchBridgeTokenPairUnpausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairUnpausedIterator{contract: _BranchBridge.contract, event: "TokenPairUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnpaused is a free log subscription operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairUnpaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairUnpaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairUnpaused)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
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

// ParseTokenPairUnpaused is a log parse operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairUnpaused(log types.Log) (*BranchBridgeTokenPairUnpaused, error) {
	event := new(BranchBridgeTokenPairUnpaused)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the BranchBridge contract.
type BranchBridgeTokenPairUnregisteredIterator struct {
	Event *BranchBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *BranchBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeTokenPairUnregistered)
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
		it.Event = new(BranchBridgeTokenPairUnregistered)
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
func (it *BranchBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the BranchBridge contract.
type BranchBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BranchBridge *BranchBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*BranchBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeTokenPairUnregisteredIterator{contract: _BranchBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BranchBridge *BranchBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *BranchBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeTokenPairUnregistered)
				if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*BranchBridgeTokenPairUnregistered, error) {
	event := new(BranchBridgeTokenPairUnregistered)
	if err := _BranchBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BranchBridge contract.
type BranchBridgeUnpausedIterator struct {
	Event *BranchBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BranchBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeUnpaused)
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
		it.Event = new(BranchBridgeUnpaused)
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
func (it *BranchBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeUnpaused represents a Unpaused event raised by the BranchBridge contract.
type BranchBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BranchBridge *BranchBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BranchBridgeUnpausedIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeUnpausedIterator{contract: _BranchBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BranchBridge *BranchBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BranchBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeUnpaused)
				if err := _BranchBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseUnpaused(log types.Log) (*BranchBridgeUnpaused, error) {
	event := new(BranchBridgeUnpaused)
	if err := _BranchBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BranchBridge contract.
type BranchBridgeUpgradedIterator struct {
	Event *BranchBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BranchBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeUpgraded)
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
		it.Event = new(BranchBridgeUpgraded)
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
func (it *BranchBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeUpgraded represents a Upgraded event raised by the BranchBridge contract.
type BranchBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BranchBridge *BranchBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BranchBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeUpgradedIterator{contract: _BranchBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BranchBridge *BranchBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BranchBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeUpgraded)
				if err := _BranchBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseUpgraded(log types.Log) (*BranchBridgeUpgraded, error) {
	event := new(BranchBridgeUpgraded)
	if err := _BranchBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BranchBridgeValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the BranchBridge contract.
type BranchBridgeValidatorSetIterator struct {
	Event *BranchBridgeValidatorSet // Event containing the contract specifics and raw log

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
func (it *BranchBridgeValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeValidatorSet)
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
		it.Event = new(BranchBridgeValidatorSet)
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
func (it *BranchBridgeValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeValidatorSet represents a ValidatorSet event raised by the BranchBridge contract.
type BranchBridgeValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BranchBridge *BranchBridgeFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*BranchBridgeValidatorSetIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeValidatorSetIterator{contract: _BranchBridge.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_BranchBridge *BranchBridgeFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeValidatorSet) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeValidatorSet)
				if err := _BranchBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseValidatorSet(log types.Log) (*BranchBridgeValidatorSet, error) {
	event := new(BranchBridgeValidatorSet)
	if err := _BranchBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
