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

// EthereumBridgeMetaData contains all meta data concerning the EthereumBridge contract.
var EthereumBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossMintableERC20FactoryCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"setChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"ChainSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RewardWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"EthereumBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"EthereumBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryAleadyExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryAleadyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errBridgeRegistryCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryInsufficientBridgeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"errBridgeRegistryTokenFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errStandardBridgeFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
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
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"814914b5": "getTokenPair(uint256,address)",
		"846704f8": "initialize(uint256,address,uint8,address,address,address)",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161561f6100f95f395f81816127040152818161272d0152612871015261561f5ff3fe6080604052600436106102dc575f3560e01c806384b0196e11610189578063b7f3358d116100d8578063edbbea3911610092578063f45096431161006d578063f450964314610977578063f698da2514610996578063facd743b146109aa578063fb75b2c7146109c9575f5ffd5b8063edbbea3914610918578063f2fde38b14610937578063f30589c314610956575f5ffd5b8063b7f3358d14610867578063cbae595814610886578063cf56118e146108a5578063d2ff130d146108c6578063d5717fc5146108e5578063d7c82f3214610904575f5ffd5b80639300c92611610143578063ad3cb1cc1161011e578063ad3cb1cc146107ca578063ae6893f8146107fa578063ae76638914610819578063aed1d40314610853575f5ffd5b80639300c9261461076b578063952a95de1461078a57806396ce0795146107b6575f5ffd5b806384b0196e146106a457806384d58d42146106cb57806388d67d6d146106ea5780638da5cb5b146106fd5780638f517c171461073957806391cf6d3e14610757575f5ffd5b806354db0126116102455780637101fcd3116101ff578063814914b5116101da578063814914b51461057e5780638415a385146106455780638456cb5914610671578063846704f814610685575f5ffd5b80637101fcd314610538578063715018a6146105575780637f260a751461056b575f5ffd5b806354db0126146104795780635958621e146104985780635b605f5c146104b75780635c975abb146104e35780635fd262de14610506578063670e626814610519575f5ffd5b80633f4ba83a116102965780633f4ba83a146103b957806340a141ff146103cd57806342cde4e8146103ec57806347666cb11461040d5780634f1ef2861461044457806352d1902d14610457575f5ffd5b8063030372c3146102f65780631327d3d81461032a5780631938e0f2146103495780631b44fd151461035c5780631d40f0d81461037b5780633960e7871461039a575f5ffd5b366102f257345f036102f0576102f0614200565b005b5f5ffd5b348015610301575f5ffd5b5061031561031036600461429b565b6109e6565b60405190151581526020015b60405180910390f35b348015610335575f5ffd5b506102f0610344366004614357565b610a2a565b610315610357366004614466565b610a38565b348015610367575f5ffd5b506102f061037636600461451f565b610d44565b348015610386575f5ffd5b506102f0610395366004614536565b610da7565b3480156103a5575f5ffd5b506103156103b43660046145db565b610de1565b3480156103c4575f5ffd5b506102f0610ed8565b3480156103d8575f5ffd5b506102f06103e7366004614357565b610eea565b3480156103f7575f5ffd5b5060355460405160ff9091168152602001610321565b348015610418575f5ffd5b5060665461042c906001600160a01b031681565b6040516001600160a01b039091168152602001610321565b6102f0610452366004614664565b610ef4565b348015610462575f5ffd5b5061046b610f0f565b604051908152602001610321565b348015610484575f5ffd5b506102f0610493366004614357565b610f2b565b3480156104a3575f5ffd5b506102f06104b2366004614357565b610fd2565b3480156104c2575f5ffd5b506104d66104d136600461451f565b61106e565b6040516103219190614702565b3480156104ee575f5ffd5b505f5160206155aa5f395f51905f525460ff16610315565b61031561051436600461478c565b6111c7565b348015610524575f5ffd5b5061042c610533366004614814565b6112d6565b348015610543575f5ffd5b506102f0610552366004614536565b6113e7565b348015610562575f5ffd5b506102f06113fd565b61031561057936600461488c565b61140e565b348015610589575f5ffd5b506106386105983660046149b0565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9182526003602090815260408084206001600160a01b039384168552600901825292839020835160a08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260020154608082015290565b60405161032191906149de565b348015610650575f5ffd5b5061066461065f3660046145db565b611519565b6040516103219190614a1a565b34801561067c575f5ffd5b506102f06115c7565b348015610690575f5ffd5b506102f061069f366004614a2c565b6115d7565b3480156106af575f5ffd5b506106b861175c565b6040516103219796959493929190614ad9565b3480156106d6575f5ffd5b506102f06106e53660046149b0565b611805565b6103156106f8366004614c5c565b6118f9565b348015610708575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661042c565b348015610744575f5ffd5b505f5461042c906001600160a01b031681565b348015610762575f5ffd5b5060685461046b565b348015610776575f5ffd5b506102f0610785366004614536565b611994565b348015610795575f5ffd5b506107a96107a43660046145db565b6119cb565b6040516103219190614d44565b3480156107c1575f5ffd5b5061046b611b0e565b3480156107d5575f5ffd5b50610664604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610805575f5ffd5b5061046b61081436600461451f565b611b7e565b348015610824575f5ffd5b50610838610833366004614da0565b611b9a565b60408051938452602084019290925290820152606001610321565b34801561085e575f5ffd5b5061046b611c29565b348015610872575f5ffd5b506102f0610881366004614dd5565b611c34565b348015610891575f5ffd5b5061042c6108a036600461451f565b611c7e565b3480156108b0575f5ffd5b506108b9611c8a565b6040516103219190614dee565b3480156108d1575f5ffd5b506102f06108e03660046149b0565b611c96565b3480156108f0575f5ffd5b5061046b6108ff36600461451f565b611d91565b34801561090f575f5ffd5b506102f0611dad565b348015610923575f5ffd5b506102f0610932366004614e0d565b611e46565b348015610942575f5ffd5b506102f0610951366004614357565b611fed565b348015610961575f5ffd5b5061096a612027565b6040516103219190614e5d565b348015610982575f5ffd5b506102f06109913660046149b0565b612033565b3480156109a1575f5ffd5b5061046b612159565b3480156109b5575f5ffd5b506103156109c4366004614357565b612162565b3480156109d4575f5ffd5b506067546001600160a01b031661042c565b5f805b8251811015610a1e57610a1584848381518110610a0857610a08614e9d565b6020026020010151610de1565b506001016109e9565b50600190505b92915050565b610a3581600161216e565b50565b5f610a41612237565b8435610a536060870160408801614357565b5f828152600360205260409020610a6d9060070182612267565b8190610a9d576040516367240f3b60e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff1615610afc5760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b50610b0561228b565b5f348015610b2f5760405163341382c560e11b815260048101929092526024820152604401610a94565b50610b3c90508735611d91565b602088013514610b4c8835611d91565b88602001359091610b79576040516329846bd360e21b815260048101929092526024820152604401610a94565b50610c0b90507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020890135610bb560608b0160408c01614357565b610bc560808c0160608d01614357565b60808c0135610bd760a08e018e614eb1565b604051602001610bed9796959493929190614f1b565b604051602081830303815290604052805190602001208787876122c2565b610c2987355f90815260036020526040902060020180546001019055565b5f80610c5a8935610c4060608c0160408d01614357565b610c5060808d0160608e01614357565b8c608001356124bd565b915091508115610ce557610c7460808a0160608b01614357565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610cb560608e0160408f01614357565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610d1e565b610cef8982612500565b60405160208a0135907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610d3a60015f5160206155ca5f395f51905f5255565b5050949350505050565b610d4c612596565b610d576001826125f1565b8190610d7957604051633d1bcfa960e01b8152600401610a9491815260200190565b5060405181907f0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a905f90a250565b5f5b8151811015610ddd57610dd5828281518110610dc757610dc7614e9d565b60200260200101515f61216e565b600101610da9565b5050565b5f610dea612237565b610df261228b565b5f610dfd84846119cb565b90505f5f610e19868460400151856060015186608001516124bd565b91509150818190610e3d5760405162461bcd60e51b8152600401610a949190614a1a565b50610e4886866125fc565b82606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8660400151876080015142604051610eb3939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a460019350505050610a2460015f5160206155ca5f395f51905f5255565b610ee0612596565b610ee86126a0565b565b610a35815f61216e565b610efc6126f9565b610f058261279d565b610ddd82826127a5565b5f610f18612866565b505f51602061558a5f395f51905f525b90565b610f33612596565b6001600160a01b038116610f7d576040516211211f60e21b81526020600482015260116024820152702fb13934b233b2a332b2a9ba30ba34b7b760791b6044820152606401610a94565b606680546001600160a01b0319166001600160a01b0383169081179091556040519081527f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444906020015b60405180910390a150565b610fda612596565b6001600160a01b038116611020576040516211211f60e21b815260206004820152600d60248201526c72657761726457616c6c65745f60981b6044820152606401610a94565b606780546001600160a01b0319166001600160a01b0383169081179091556040519081527f5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b90602001610fc7565b5f81815260036020526040812060609161108a600783016128af565b90505f81516001600160401b038111156110a6576110a6614214565b6040519080825280602002602001820160405280156110fd57816020015b6040805160a0810182525f808252602080830182905292820181905260608201819052608082015282525f199092019101816110c45790505b5090505f5b82518110156111be57836009015f84838151811061112257611122614e9d565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160a08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b90930490911615156060820152600290910154608082015282518390839081106111ab576111ab614e9d565b6020908102919091010152600101611102565b50949350505050565b5f6111d0612237565b5f898152600360205260409020899089906111ee9060070182612267565b8190611219576040516367240f3b60e01b81526001600160a01b039091166004820152602401610a94565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156112785760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b5061128161228b565b5f5f6112908d8d8c8c8c6128bb565b915091506112ac8d8d6112a03390565b8e8e87875f8f8f6129d4565b6001945050506112c860015f5160206155ca5f395f51905f5255565b505098975050505050505050565b5f6112df612596565b5f546001600160a01b031661130757604051636274ce8960e11b815260040160405180910390fd5b5f546040516bffffffffffffffffffffffff19606087901b1660208201526001600160a01b0390911690634804a0419060340160405160208183030381529060405280519060200120856040516020016113619190614f7f565b60405160208183030381529060405286866040518563ffffffff1660e01b81526004016113919493929190614fa0565b6020604051808303815f875af11580156113ad573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113d19190614fdf565b90506113df855f8387611e46565b949350505050565b6113f461039560336128af565b610a3581611994565b611405612596565b610ee85f612a22565b5f611417612237565b5f8b81526003602052604090208b908b906114359060070182612267565b8190611460576040516367240f3b60e01b81526001600160a01b039091166004820152602401610a94565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156114bf5760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b506114c861228b565b5f5f6114d78f8f8d8d8d6128bb565b915091506114ed8f8f8f8f8f87878f8f8f612a92565b60019450505061150960015f5160206155ca5f395f51905f5255565b50509a9950505050505050505050565b5f828152600360209081526040808320848452600401909152902080546060919061154390614ffa565b80601f016020809104026020016040519081016040528092919081815260200182805461156f90614ffa565b80156115ba5780601f10611591576101008083540402835291602001916115ba565b820191905f5260205f20905b81548152906001019060200180831161159d57829003601f168201915b5050505050905092915050565b6115cf612596565b610ee8612afd565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561161b5750825b90505f826001600160401b031660011480156116365750303b155b905081158015611644575080155b156116625760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561168c57845460ff60401b1916600160401b1785555b61169889898989612b45565b6001600160a01b038a166116d75760405163634fe78160e11b815260206004820152600560248201526463726f737360d81b6044820152606401610a94565b609880546001600160a01b0319166001600160a01b038c161790556116fb8b610d44565b6117098b60018c6001611e46565b831561174f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b5f60608082808083815f51602061556a5f395f51905f52805490915015801561178757506001810154155b6117cb5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a94565b6117d3612c19565b6117db612cd9565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61180d612596565b5f8281526003602052604090206118276007820183612267565b829061185257604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff166118a55760405163273bf77b60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b191690555185917fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e991a3505050565b5f805b858110156119875761197e87878381811061191957611919614e9d565b905060200281019061192b9190615032565b86838151811061193d5761193d614e9d565b602002602001015186848151811061195757611957614e9d565b602002602001015186858151811061197157611971614e9d565b6020026020010151610a38565b506001016118fc565b5060019695505050505050565b5f5b8151811015610ddd576119c38282815181106119b4576119b4614e9d565b6020026020010151600161216e565b600101611996565b611a116040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f8381526003602081815260408084208685528301825292839020835160c0810185528154815260018201549281019290925260028101546001600160a01b03908116948301949094529182015490921660608301526004810154608083015260058101805460a084019190611a8690614ffa565b80601f0160208091040260200160405190810160405280929190818152602001828054611ab290614ffa565b8015611afd5780601f10611ad457610100808354040283529160200191611afd565b820191905f5260205f20905b815481529060010190602001808311611ae057829003601f168201915b505050505081525050905092915050565b606654604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015611b55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b799190615050565b905090565b5f818152600360205260408120600190810154610a249161507b565b60665460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa158015611bf6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c1a919061508e565b91989097509095509350505050565b5f611b796033612d17565b611c3c612596565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff90602001610fc7565b5f610a24603383612d20565b6060611b7960016128af565b611c9e612596565b5f828152600360205260409020611cb86007820183612267565b8290611ce357604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff1615611d375760405163283e1c6360e21b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b1916600160a81b1790555185917ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a391a3505050565b5f81815260036020526040812060020154610a2490600161507b565b611db5612596565b6066546001600160a01b0316611e0157604051635f6c3f4f60e01b815260206004820152601060248201526f313934b233b2a332b2a9ba30ba34b7b760811b6044820152606401610a94565b606680546001600160a01b03191690556040515f81527f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea4449060200160405180910390a1565b611e4e612596565b611e59600185612d2b565b8490611e7b5760405163c45f2c2360e01b8152600401610a9491815260200190565b505f8481526003602052604090206001600160a01b038316611ecd57604051636e0b3c5f60e11b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a94565b611eda6007820184612d42565b8390611f0557604051634d2df28360e01b81526001600160a01b039091166004820152602401610a94565b506040805160a0810182526001600160a01b0385811680835285821660208085018281528a15158688018181525f6060890181815260808a0182815288835260098d018752918b902099518a546001600160a01b031916908a16178a55935160018a01805493519551919099166001600160a81b031990931692909217600160a01b941515949094029390931760ff60a81b1916600160a81b91151591909102179095555160029095019490945593519182529188917f78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c910160405180910390a45050505050565b611ff5612596565b6001600160a01b03811661201e57604051631e4fbdf760e01b81525f6004820152602401610a94565b610a3581612a22565b6060611b7960336128af565b61203b612596565b612046600183612d2b565b82906120685760405163c45f2c2360e01b8152600401610a9491815260200190565b506001600160a01b0381166120a857604051636e0b3c5f60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a94565b5f8281526003602052604090206120c26007820183612d56565b82906120ed57604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f81815260098301602052604080822080546001600160a01b03191681556001810180546001600160b01b03191690556002018290555185917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d91a3505050565b5f611b79612d6a565b5f610a24603383612267565b612176612596565b80156121b857612187603383612d42565b82906121b2576040516329a04e7760e21b81526001600160a01b039091166004820152602401610a94565b506121f0565b6121c3603383612d56565b82906121ee5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610a94565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f5160206155aa5f395f51905f525460ff1615610ee85760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f90815260018301602052604081205415155b9392505050565b5f5160206155ca5f395f51905f528054600119016122bc57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b82518251811480156122d45750815181145b835183518392612308576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610a94565b505060355482915060ff1681101561233657604051632fcba65760e11b8152600401610a9491815260200190565b505f80826001600160401b0381111561235157612351614214565b60405190808252806020026020018201604052801561237a578160200160208202803683370190505b5090505f5b83811015612487575f6123ea88838151811061239d5761239d614e9d565b60200260200101518884815181106123b7576123b7614e9d565b60200260200101518885815181106123d1576123d1614e9d565b60200260200101516123e28d612d73565b929190612d9f565b90506123f581612162565b81906124205760405163845a09e760e01b81526001600160a01b039091166004820152602401610a94565b505f805b84518110156124705784818151811061243f5761243f614e9d565b60200260200101516001600160a01b0316836001600160a01b0316036124685760019150612470565b600101612424565b508061247d578460010194505b505060010161237f565b50603554829060ff168110156124b357604051632fcba65760e11b8152600401610a9491815260200190565b5050505050505050565b6098545f906060906001600160a01b03908116908616036124e6576124e36064846150b9565b92505b6124f286868686612dcb565b915091505b94509492505050565b81355f9081526003602090815260409091209083013561252360058301826125f1565b8190612545576040516306bfa19360e31b8152600401610a9491815260200190565b505f8181526003830160205260409020849061256182826151f5565b50505f818152600483016020526040902061257c8482615295565b5050505050565b60015f5160206155ca5f395f51905f5255565b336125c87f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610ee85760405163118cdaa760e01b8152336004820152602401610a94565b5f6122848383612e3f565b5f8281526003602052604090206126166005820183612e8b565b82906126385760405163052b63c160e11b8152600401610a9491815260200190565b505f8281526004820160205260408120612651916141b6565b5f828152600380830160205260408220828155600181018390556002810180546001600160a01b03199081169091559181018054909216909155600481018290559061257c60058301826141b6565b6126a8612e96565b5f5160206155aa5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610fc7565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061277f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166127735f51602061558a5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610ee85760405163703e46dd60e11b815260040160405180910390fd5b610a35612596565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156127ff575060408051601f3d908101601f191682019092526127fc91810190615050565b60015b61282757604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a94565b5f51602061558a5f395f51905f52811461285757604051632a87526960e21b815260048101829052602401610a94565b6128618383612ec5565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ee85760405163703e46dd60e11b815260040160405180910390fd5b60605f61228483612f1a565b6066545f9081906001600160a01b03166128d957505f9050806129ca565b60665460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa15801561292f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612953919061508e565b9094509250905080861080159061296a5750828510155b80156129765750818410155b8184848989899091929394956129c25760405163769296ad60e01b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a94565b505050505050505b9550959350505050565b6129ea8a8a8a896129e5898b61507b565b612f73565b6129fc8a8a8a8a8a8a8a8a8a8a613119565b5f8a81526003602052604090206001908101805490910190555b50505050505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b83612a9d868861507b565b612aa7919061507b565b8360400151101589879091612adf576040516224b63960e71b81526001600160a01b0390921660048301526024820152604401610a94565b5050612aea83613211565b612a168a8a8a8a8a8a8a60018a8a6129d4565b612b05612237565b5f5160206155aa5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336126e1565b612b4d61330e565b612b5633613357565b612b5f84613368565b612b68826133ce565b612b7061350c565b612b78613514565b612b80613524565b6001600160a01b03811615612bab57606680546001600160a01b0319166001600160a01b0383161790555b6001600160a01b038316612bf0576040516211211f60e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a94565b5050606780546001600160a01b0319166001600160a01b03929092169190911790555043606855565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061556a5f395f51905f5291612c5790614ffa565b80601f0160208091040260200160405190810160405280929190818152602001828054612c8390614ffa565b8015612cce5780601f10612ca557610100808354040283529160200191612cce565b820191905f5260205f20905b815481529060010190602001808311612cb157829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061556a5f395f51905f5291612c5790614ffa565b5f610a24825490565b5f6122848383613534565b5f8181526001830160205260408120541515612284565b5f612284836001600160a01b038416612e3f565b5f612284836001600160a01b03841661355a565b5f611b7961363d565b5f610a24612d7f612d6a565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612daf888888886136b0565b925092509250612dbf8282613778565b50909695505050505050565b5f60605f196001600160a01b03861601612e0a57612df26001600160a01b03851684613830565b505060408051602081019091525f81526001906124f7565b82156124f757612e1a86866138c2565b15612e3457612e2b868686866138fa565b915091506124f7565b612e2b858585613aac565b5f818152600183016020526040812054612e8457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a24565b505f610a24565b5f612284838361355a565b5f5160206155aa5f395f51905f525460ff16610ee857604051638dfc202b60e01b815260040160405180910390fd5b612ece82613c53565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612f12576128618282613cb6565b610ddd613d28565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612f6757602002820191905f5260205f20905b815481526020019060010190808311612f53575b50505050509050919050565b5f196001600160a01b03851601612fe657612f8e818361507b565b3414612f9a828461507b565b349091612fc35760405163341382c560e11b815260048101929092526024820152604401610a94565b50508015612fe157606754612fe1906001600160a01b031682613830565b61257c565b5f3480156130105760405163341382c560e11b815260048101929092526024820152604401610a94565b5061303490508330613022848661507b565b6001600160a01b038816929190613d47565b801561305457606754613054906001600160a01b03868116911683613dae565b61305e85856138c2565b1561306e57612fe1858584613ddf565b604051632770a7eb60e21b8152306004820152602481018390526001600160a01b03851690639dc29fac906044016020604051808303815f875af11580156130b8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130dc919061534f565b8484849091926124b35760405163bc8ecf7b60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610a94565b5f5f6131248c611b7e565b91506131598c8c5f9182526003602090815260408084206001600160a01b039384168552600901909152909120600101541690565b9050896001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d78e858e8e428d8d8d6040516131a498979695949392919061536a565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a6040516131fb929190918252602082015260400190565b60405180910390a4505050505050505050505050565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af1806132bd576040513d5f823e3d81fd5b50505f513d915081156132d45780600114156132e2565b84516001600160a01b03163b155b1561257c57845160405163733eb19560e01b81526001600160a01b039091166004820152602401610a94565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610ee857604051631afcd79f60e31b815260040160405180910390fd5b61335f61330e565b610a3581613e20565b61337061330e565b6133b8604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613e28565b6035805460ff191660ff92909216919091179055565b6001600160a01b0381166133df5750565b5f816001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa15801561341b573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261344291908101906153b4565b604080513060208201520160408051601f198184030181529082905261346b9291602001615433565b60408051601f1981840301815282825246602084015292505f916134a8918391016040516020818303038152906040528051906020012084613e3a565b90506001600160a01b0381166134e95760405162461bcd60e51b8152600401610a94906020808252600490820152637a65726f60e01b604082015260600190565b5f80546001600160a01b0319166001600160a01b03929092169190911790555050565b610ee861330e565b61351c61330e565b610ee8613ece565b61352c61330e565b610ee8613eee565b5f825f01828154811061354957613549614e9d565b905f5260205f200154905092915050565b5f8181526001830160205260408120548015613634575f61357c600183615447565b85549091505f9061358f90600190615447565b90508082146135ee575f865f0182815481106135ad576135ad614e9d565b905f5260205f200154905080875f0184815481106135cd576135cd614e9d565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806135ff576135ff61545a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a24565b5f915050610a24565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613667613ef6565b61366f613f5e565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156136e957505f9150600390508261376e565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561373a573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661376557505f92506001915082905061376e565b92505f91508190505b9450945094915050565b5f82600381111561378b5761378b61546e565b03613794575050565b60018260038111156137a8576137a861546e565b036137c65760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156137da576137da61546e565b036137fb5760405163fce698f760e01b815260048101829052602401610a94565b600382600381111561380f5761380f61546e565b03610ddd576040516335e2f38360e21b815260048101829052602401610a94565b8047101561385a5760405163cf47918160e01b815247600482015260248101829052604401610a94565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f81146138a4576040519150601f19603f3d011682016040523d82523d5f602084013e6138a9565b606091505b5091509150816138bc576138bc81613fa0565b50505050565b5f9182526003602090815260408084206001600160a01b0393909316845260099092019052902060010154600160a01b900460ff1690565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613969575060408051601f3d908101601f191682019092526139669181019061534f565b60015b613a3b57613975615482565b806308c379a00361399e575061398961549a565b8061399457506139f8565b5f925090506124f7565b634e487b71036139f8576139b061551c565b906139bb57506139f8565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b6040516020818303038152906040529150506124f7565b3d808015613a21576040519150601f19603f3d011682016040523d82523d5f602084013e613a26565b606091505b505f9250806040516020016139e19190615539565b8015613a66576001925060405180602001604052805f8152509150613a61878786613fc9565b613aa2565b5f92506040518060400160405280601f81526020017f5374616e646172644272696467653a207472616e73666572206661696c65640081525091505b5094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613b1b575060408051601f3d908101601f19168201909252613b189181019061534f565b60015b613bed57613b27615482565b806308c379a003613b505750613b3b61549a565b80613b465750613baa565b5f92509050613c4b565b634e487b7103613baa57613b6261551c565b90613b6d5750613baa565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b604051602081830303815290604052915050613c4b565b3d808015613bd3576040519150601f19603f3d011682016040523d82523d5f602084013e613bd8565b606091505b505f925080604051602001613b939190615539565b8015613c0d576001925060405180602001604052805f8152509150613c49565b5f92506040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b806001600160a01b03163b5f03613c8857604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a94565b5f51602061558a5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613cd2919061555e565b5f60405180830381855af49150503d805f8114613d0a576040519150601f19603f3d011682016040523d82523d5f602084013e613d0f565b606091505b5091509150613d1f85838361406b565b95945050505050565b3415610ee85760405163b398979f60e01b815260040160405180910390fd5b6040516001600160a01b0384811660248301528381166044830152606482018390526138bc9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506140c7565b6040516001600160a01b0383811660248301526044820183905261286191859182169063a9059cbb90606401613d7c565b5f8381526003602090815260408083206001600160a01b038616845260090190915281206002018054839290613e1690849061507b565b9091555050505050565b611ff561330e565b613e3061330e565b610ddd8282614133565b5f83471015613e655760405163cf47918160e01b815247600482015260248101859052604401610a94565b81515f03613e8657604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615613ea7576040513d5f823e3d81fd5b6001600160a01b0381166122845760405163b06ebf3d60e01b815260040160405180910390fd5b613ed661330e565b5f5160206155aa5f395f51905f52805460ff19169055565b61258361330e565b5f5f51602061556a5f395f51905f5281613f0e612c19565b805190915015613f2657805160209091012092915050565b81548015613f35579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061556a5f395f51905f5281613f76612cd9565b805190915015613f8e57805160209091012092915050565b60018201548015613f35579392505050565b805115613fb05780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f8381526003602090815260408083206001600160a01b03861684526009810190925282206002015490919081906140019085614192565b8454919350915085858461404157604051630106202f60e71b815260048101939093526001600160a01b0390911660248301526044820152606401610a94565b5050506001600160a01b039094165f90815260099092016020525060409020600201919091555050565b6060826140805761407b82613fa0565b612284565b815115801561409757506001600160a01b0384163b155b156140c057604051639996b31560e01b81526001600160a01b0385166004820152602401610a94565b5080612284565b5f5f60205f8451602086015f885af1806140e6576040513d5f823e3d81fd5b50505f513d915081156140fd57806001141561410a565b6001600160a01b0384163b155b156138bc57604051635274afe760e01b81526001600160a01b0385166004820152602401610a94565b61413b61330e565b5f51602061556a5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026141748482615295565b50600381016141838382615295565b505f8082556001909101555050565b5f5f838311156141a657505f9050806141af565b50600190508183035b9250929050565b5080546141c290614ffa565b5f825580601f106141d1575050565b601f0160209004905f5260205f2090810190610a3591905b808211156141fc575f81556001016141e9565b5090565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b038211171561424757614247614214565b60405250565b601f8201601f191681016001600160401b038111828210171561427257614272614214565b6040525050565b5f6001600160401b0382111561429157614291614214565b5060051b60200190565b5f5f604083850312156142ac575f5ffd5b8235915060208301356001600160401b038111156142c8575f5ffd5b8301601f810185136142d8575f5ffd5b80356142e381614279565b6040516142f0828261424d565b80915082815260208101915060208360051b850101925087831115614313575f5ffd5b6020840193505b8284101561433557833582526020938401939091019061431a565b809450505050509250929050565b6001600160a01b0381168114610a35575f5ffd5b5f60208284031215614367575f5ffd5b813561228481614343565b803560ff81168114614382575f5ffd5b919050565b5f82601f830112614396575f5ffd5b81356143a181614279565b6040516143ae828261424d565b80915082815260208101915060208360051b8601019250858311156143d1575f5ffd5b602085015b838110156143f5576143e781614372565b8352602092830192016143d6565b5095945050505050565b5f82601f83011261440e575f5ffd5b813561441981614279565b604051614426828261424d565b80915082815260208101915060208360051b860101925085831115614449575f5ffd5b602085015b838110156143f557803583526020928301920161444e565b5f5f5f5f60808587031215614479575f5ffd5b84356001600160401b0381111561448e575f5ffd5b850160c0818803121561449f575f5ffd5b935060208501356001600160401b038111156144b9575f5ffd5b6144c587828801614387565b93505060408501356001600160401b038111156144e0575f5ffd5b6144ec878288016143ff565b92505060608501356001600160401b03811115614507575f5ffd5b614513878288016143ff565b91505092959194509250565b5f6020828403121561452f575f5ffd5b5035919050565b5f60208284031215614546575f5ffd5b81356001600160401b0381111561455b575f5ffd5b8201601f8101841361456b575f5ffd5b803561457681614279565b604051614583828261424d565b80915082815260208101915060208360051b8501019250868311156145a6575f5ffd5b6020840193505b828410156145d15783356145c081614343565b8252602093840193909101906145ad565b9695505050505050565b5f5f604083850312156145ec575f5ffd5b50508035926020909101359150565b5f6001600160401b0382111561461357614613614214565b50601f01601f191660200190565b5f61462b836145fb565b604051614638828261424d565b80925084815285858501111561464c575f5ffd5b848460208301375f6020868301015250509392505050565b5f5f60408385031215614675575f5ffd5b823561468081614343565b915060208301356001600160401b0381111561469a575f5ffd5b8301601f810185136146aa575f5ffd5b6146b985823560208401614621565b9150509250929050565b80516001600160a01b03908116835260208083015190911690830152604080820151151590830152606080820151151590830152608090810151910152565b602080825282518282018190525f918401906040840190835b818110156147445761472e8385516146c3565b6020939093019260a0929092019160010161471b565b509095945050505050565b5f5f83601f84011261475f575f5ffd5b5081356001600160401b03811115614775575f5ffd5b6020830191508360208285010111156141af575f5ffd5b5f5f5f5f5f5f5f5f60e0898b0312156147a3575f5ffd5b8835975060208901356147b581614343565b965060408901356147c581614343565b9550606089013594506080890135935060a0890135925060c08901356001600160401b038111156147f4575f5ffd5b6148008b828c0161474f565b999c989b5096995094979396929594505050565b5f5f5f5f60808587031215614827575f5ffd5b84359350602085013561483981614343565b925060408501356001600160401b03811115614853575f5ffd5b8501601f81018713614863575f5ffd5b61487287823560208401614621565b92505061488160608601614372565b905092959194509250565b5f5f5f5f5f5f5f5f5f5f8a8c036101e08112156148a7575f5ffd5b8b359a5060208c01356148b981614343565b995060408c01356148c981614343565b985060608c01356148d981614343565b975060808c0135965060a08c0135955060c08c0135945060e060df1982011215614901575f5ffd5b5060405161490e81614228565b60e08c013561491c81614343565b81526101008c013561492d81614343565b60208201526101208c013560408201526101408c013560608201526149556101608d01614372565b60808201526101808c013560a08201526101a08c013560c082015292506101c08b01356001600160401b0381111561498b575f5ffd5b6149978d828e0161474f565b915080935050809150509295989b9194979a5092959850565b5f5f604083850312156149c1575f5ffd5b8235915060208301356149d381614343565b809150509250929050565b60a08101610a2482846146c3565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61228460208301846149ec565b5f5f5f5f5f5f60c08789031215614a41575f5ffd5b863595506020870135614a5381614343565b9450614a6160408801614372565b93506060870135614a7181614343565b92506080870135614a8181614343565b915060a0870135614a9181614343565b809150509295509295509295565b5f8151808452602084019350602083015f5b82811015614acf578151865260209586019590910190600101614ab1565b5093949350505050565b60ff60f81b8816815260e060208201525f614af760e08301896149ec565b8281036040840152614b0981896149ec565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050614b3a8185614a9f565b9a9950505050505050505050565b5f82601f830112614b57575f5ffd5b8135614b6281614279565b604051614b6f828261424d565b80915082815260208101915060208360051b860101925085831115614b92575f5ffd5b602085015b838110156143f55780356001600160401b03811115614bb4575f5ffd5b614bc3886020838a0101614387565b84525060209283019201614b97565b5f82601f830112614be1575f5ffd5b8135614bec81614279565b604051614bf9828261424d565b80915082815260208101915060208360051b860101925085831115614c1c575f5ffd5b602085015b838110156143f55780356001600160401b03811115614c3e575f5ffd5b614c4d886020838a01016143ff565b84525060209283019201614c21565b5f5f5f5f5f60808688031215614c70575f5ffd5b85356001600160401b03811115614c85575f5ffd5b8601601f81018813614c95575f5ffd5b80356001600160401b03811115614caa575f5ffd5b8860208260051b8401011115614cbe575f5ffd5b6020918201965094508601356001600160401b03811115614cdd575f5ffd5b614ce988828901614b48565b93505060408601356001600160401b03811115614d04575f5ffd5b614d1088828901614bd2565b92505060608601356001600160401b03811115614d2b575f5ffd5b614d3788828901614bd2565b9150509295509295909350565b60208152815160208201526020820151604082015260018060a01b03604083015116606082015260018060a01b036060830151166080820152608082015160a08201525f60a083015160c0808401526113df60e08401826149ec565b5f5f5f60608486031215614db2575f5ffd5b833592506020840135614dc481614343565b929592945050506040919091013590565b5f60208284031215614de5575f5ffd5b61228482614372565b602081525f6122846020830184614a9f565b8015158114610a35575f5ffd5b5f5f5f5f60808587031215614e20575f5ffd5b843593506020850135614e3281614e00565b92506040850135614e4281614343565b91506060850135614e5281614343565b939692955090935050565b602080825282518282018190525f918401906040840190835b818110156147445783516001600160a01b0316835260209384019390920191600101614e76565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112614ec6575f5ffd5b8301803591506001600160401b03821115614edf575f5ffd5b6020019150368190038213156141af575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90614f5b9083018486614ef3565b9998505050505050505050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f612284600d830184614f68565b848152608060208201525f614fb860808301866149ec565b8281036040840152614fca81866149ec565b91505060ff8316606083015295945050505050565b5f60208284031215614fef575f5ffd5b815161228481614343565b600181811c9082168061500e57607f821691505b60208210810361502c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be19833603018112615046575f5ffd5b9190910192915050565b5f60208284031215615060575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610a2457610a24615067565b5f5f5f606084860312156150a0575f5ffd5b5050815160208301516040909301519094929350919050565b5f826150d357634e487b7160e01b5f52601260045260245ffd5b500490565b80546001600160a01b0319166001600160a01b0392909216919091179055565b601f82111561286157805f5260205f20601f840160051c8101602085101561511d5750805b601f840160051c820191505b8181101561257c575f8155600101615129565b6001600160401b0383111561515357615153614214565b615167836151618354614ffa565b836150f8565b5f601f841160018114615198575f85156151815750838201355b5f19600387901b1c1916600186901b17835561257c565b5f83815260208120601f198716915b828110156151c757868501358255602094850194600190920191016151a7565b50868210156151e3575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8135815560208201356001820155604082013561521181614343565b61521e81600284016150d8565b50606082013561522d81614343565b61523a81600384016150d8565b506080820135600482015560a082013536839003601e1901811261525c575f5ffd5b820180356001600160401b03811115615273575f5ffd5b602082019150803603821315615287575f5ffd5b6138bc81836005860161513c565b81516001600160401b038111156152ae576152ae614214565b6152c2816152bc8454614ffa565b846150f8565b6020601f8211600181146152f4575f83156152dd5750848201515b5f19600385901b1c1916600184901b17845561257c565b5f84815260208120601f198516915b828110156153235787850151825560209485019460019092019101615303565b508482101561534057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6020828403121561535f575f5ffd5b815161228481614e00565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f90614b3a9083018486614ef3565b5f602082840312156153c4575f5ffd5b81516001600160401b038111156153d9575f5ffd5b8201601f810184136153e9575f5ffd5b80516153f4816145fb565b604051615401828261424d565b828152866020848601011115615415575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b5f6113df6154418386614f68565b84614f68565b81810381811115610a2457610a24615067565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f60033d1115610f285760045f5f3e505f5160e01c90565b5f60443d10156154a75790565b6040513d600319016004823e80513d60248201116001600160401b03821117156154d057505090565b80820180516001600160401b038111156154eb575050505090565b3d8401600319018282016020011115615505575050505090565b6155146020828501018561424d565b509392505050565b5f5f60233d111561553557602060045f3e50505f516001905b9091565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f6122846011830184614f68565b5f6122848284614f6856fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122057412b7f954c51c7f17b0aea6f5c0e778b0396690ada7adb98c57cb77bf4ae7764736f6c634300081c0033",
}

// EthereumBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumBridgeMetaData.ABI instead.
var EthereumBridgeABI = EthereumBridgeMetaData.ABI

// Deprecated: Use EthereumBridgeMetaData.Sigs instead.
// EthereumBridgeFuncSigs maps the 4-byte function signature to its string representation.
var EthereumBridgeFuncSigs = EthereumBridgeMetaData.Sigs

// EthereumBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumBridgeMetaData.Bin instead.
var EthereumBridgeBin = EthereumBridgeMetaData.Bin

// DeployEthereumBridge deploys a new Ethereum contract, binding an instance of EthereumBridge to it.
func DeployEthereumBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthereumBridge, error) {
	parsed, err := EthereumBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthereumBridge{EthereumBridgeCaller: EthereumBridgeCaller{contract: contract}, EthereumBridgeTransactor: EthereumBridgeTransactor{contract: contract}, EthereumBridgeFilterer: EthereumBridgeFilterer{contract: contract}}, nil
}

// EthereumBridge is an auto generated Go binding around an Ethereum contract.
type EthereumBridge struct {
	EthereumBridgeCaller     // Read-only binding to the contract
	EthereumBridgeTransactor // Write-only binding to the contract
	EthereumBridgeFilterer   // Log filterer for contract events
}

// EthereumBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumBridgeSession struct {
	Contract     *EthereumBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumBridgeCallerSession struct {
	Contract *EthereumBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EthereumBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumBridgeTransactorSession struct {
	Contract     *EthereumBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EthereumBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumBridgeRaw struct {
	Contract *EthereumBridge // Generic contract binding to access the raw methods on
}

// EthereumBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumBridgeCallerRaw struct {
	Contract *EthereumBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumBridgeTransactorRaw struct {
	Contract *EthereumBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereumBridge creates a new instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridge(address common.Address, backend bind.ContractBackend) (*EthereumBridge, error) {
	contract, err := bindEthereumBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthereumBridge{EthereumBridgeCaller: EthereumBridgeCaller{contract: contract}, EthereumBridgeTransactor: EthereumBridgeTransactor{contract: contract}, EthereumBridgeFilterer: EthereumBridgeFilterer{contract: contract}}, nil
}

// NewEthereumBridgeCaller creates a new read-only instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeCaller(address common.Address, caller bind.ContractCaller) (*EthereumBridgeCaller, error) {
	contract, err := bindEthereumBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeCaller{contract: contract}, nil
}

// NewEthereumBridgeTransactor creates a new write-only instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumBridgeTransactor, error) {
	contract, err := bindEthereumBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTransactor{contract: contract}, nil
}

// NewEthereumBridgeFilterer creates a new log filterer instance of EthereumBridge, bound to a specific deployed contract.
func NewEthereumBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumBridgeFilterer, error) {
	contract, err := bindEthereumBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeFilterer{contract: contract}, nil
}

// bindEthereumBridge binds a generic wrapper to an already deployed contract.
func bindEthereumBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBridge *EthereumBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumBridge.Contract.EthereumBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBridge *EthereumBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.Contract.EthereumBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBridge *EthereumBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBridge.Contract.EthereumBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBridge *EthereumBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBridge *EthereumBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBridge *EthereumBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBridge.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EthereumBridge *EthereumBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EthereumBridge *EthereumBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EthereumBridge.Contract.UPGRADEINTERFACEVERSION(&_EthereumBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EthereumBridge *EthereumBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EthereumBridge.Contract.UPGRADEINTERFACEVERSION(&_EthereumBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_EthereumBridge *EthereumBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_EthereumBridge *EthereumBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _EthereumBridge.Contract.AllChainIDs(&_EthereumBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _EthereumBridge.Contract.AllChainIDs(&_EthereumBridge.CallOpts)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.AllTokenPairs(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.AllTokenPairs(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_EthereumBridge *EthereumBridgeCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_EthereumBridge *EthereumBridgeSession) AllValidators() ([]common.Address, error) {
	return _EthereumBridge.Contract.AllValidators(&_EthereumBridge.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllValidators() ([]common.Address, error) {
	return _EthereumBridge.Contract.AllValidators(&_EthereumBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) BridgeFeeStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "bridgeFeeStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) BridgeFeeStation() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeFeeStation(&_EthereumBridge.CallOpts)
}

// BridgeFeeStation is a free data retrieval call binding the contract method 0x47666cb1.
//
// Solidity: function bridgeFeeStation() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) BridgeFeeStation() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeFeeStation(&_EthereumBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) CrossMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "crossMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) CrossMintableERC20Factory() (common.Address, error) {
	return _EthereumBridge.Contract.CrossMintableERC20Factory(&_EthereumBridge.CallOpts)
}

// CrossMintableERC20Factory is a free data retrieval call binding the contract method 0x8f517c17.
//
// Solidity: function crossMintableERC20Factory() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) CrossMintableERC20Factory() (common.Address, error) {
	return _EthereumBridge.Contract.CrossMintableERC20Factory(&_EthereumBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) Denominator() (*big.Int, error) {
	return _EthereumBridge.Contract.Denominator(&_EthereumBridge.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) Denominator() (*big.Int, error) {
	return _EthereumBridge.Contract.Denominator(&_EthereumBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeSession) DomainSeparator() ([32]byte, error) {
	return _EthereumBridge.Contract.DomainSeparator(&_EthereumBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _EthereumBridge.Contract.DomainSeparator(&_EthereumBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EthereumBridge *EthereumBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_EthereumBridge *EthereumBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EthereumBridge.Contract.Eip712Domain(&_EthereumBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EthereumBridge *EthereumBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EthereumBridge.Contract.Eip712Domain(&_EthereumBridge.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_EthereumBridge *EthereumBridgeCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

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
func (_EthereumBridge *EthereumBridgeSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _EthereumBridge.Contract.EstimateFee(&_EthereumBridge.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_EthereumBridge *EthereumBridgeCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _EthereumBridge.Contract.EstimateFee(&_EthereumBridge.CallOpts, remoteChainID, token, value)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _EthereumBridge.Contract.GetNextFinalizeIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _EthereumBridge.Contract.GetNextFinalizeIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _EthereumBridge.Contract.GetNextInitiateIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _EthereumBridge.Contract.GetNextInitiateIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_EthereumBridge *EthereumBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_EthereumBridge *EthereumBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.GetTokenPair(&_EthereumBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_EthereumBridge *EthereumBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.GetTokenPair(&_EthereumBridge.CallOpts, remoteChainID, token)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) InitializedAt() (*big.Int, error) {
	return _EthereumBridge.Contract.InitializedAt(&_EthereumBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _EthereumBridge.Contract.InitializedAt(&_EthereumBridge.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) IsValidator(validator common.Address) (bool, error) {
	return _EthereumBridge.Contract.IsValidator(&_EthereumBridge.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _EthereumBridge.Contract.IsValidator(&_EthereumBridge.CallOpts, validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) Owner() (common.Address, error) {
	return _EthereumBridge.Contract.Owner(&_EthereumBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) Owner() (common.Address, error) {
	return _EthereumBridge.Contract.Owner(&_EthereumBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) Paused() (bool, error) {
	return _EthereumBridge.Contract.Paused(&_EthereumBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) Paused() (bool, error) {
	return _EthereumBridge.Contract.Paused(&_EthereumBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _EthereumBridge.Contract.ProxiableUUID(&_EthereumBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EthereumBridge.Contract.ProxiableUUID(&_EthereumBridge.CallOpts)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeCaller) RevertedArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "revertedArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryFinalizeArguments)).(*IBridgeRegistryFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _EthereumBridge.Contract.RevertedArguments(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeCallerSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _EthereumBridge.Contract.RevertedArguments(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_EthereumBridge *EthereumBridgeCaller) RevertedReason(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (string, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "revertedReason", remoteChainID, index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_EthereumBridge *EthereumBridgeSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _EthereumBridge.Contract.RevertedReason(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// RevertedReason is a free data retrieval call binding the contract method 0x8415a385.
//
// Solidity: function revertedReason(uint256 remoteChainID, uint256 index) view returns(string)
func (_EthereumBridge *EthereumBridgeCallerSession) RevertedReason(remoteChainID *big.Int, index *big.Int) (string, error) {
	return _EthereumBridge.Contract.RevertedReason(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) RewardWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "rewardWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) RewardWallet() (common.Address, error) {
	return _EthereumBridge.Contract.RewardWallet(&_EthereumBridge.CallOpts)
}

// RewardWallet is a free data retrieval call binding the contract method 0xfb75b2c7.
//
// Solidity: function rewardWallet() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) RewardWallet() (common.Address, error) {
	return _EthereumBridge.Contract.RewardWallet(&_EthereumBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_EthereumBridge *EthereumBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_EthereumBridge *EthereumBridgeSession) Threshold() (uint8, error) {
	return _EthereumBridge.Contract.Threshold(&_EthereumBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_EthereumBridge *EthereumBridgeCallerSession) Threshold() (uint8, error) {
	return _EthereumBridge.Contract.Threshold(&_EthereumBridge.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_EthereumBridge *EthereumBridgeSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _EthereumBridge.Contract.ValidatorByIndex(&_EthereumBridge.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _EthereumBridge.Contract.ValidatorByIndex(&_EthereumBridge.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_EthereumBridge *EthereumBridgeSession) ValidatorLength() (*big.Int, error) {
	return _EthereumBridge.Contract.ValidatorLength(&_EthereumBridge.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_EthereumBridge *EthereumBridgeCallerSession) ValidatorLength() (*big.Int, error) {
	return _EthereumBridge.Contract.ValidatorLength(&_EthereumBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) BridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "bridgeToken", remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.BridgeToken(&_EthereumBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 remoteChainID, address token, address to, uint256 value, uint256 gasFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) BridgeToken(remoteChainID *big.Int, token common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.BridgeToken(&_EthereumBridge.TransactOpts, remoteChainID, token, to, value, gasFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_EthereumBridge *EthereumBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ChangeThreshold(&_EthereumBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ChangeThreshold(&_EthereumBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_EthereumBridge *EthereumBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_EthereumBridge *EthereumBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.CreateToken(&_EthereumBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_EthereumBridge *EthereumBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.CreateToken(&_EthereumBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridge(&_EthereumBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridge(&_EthereumBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridgeBatch(&_EthereumBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridgeBatch(&_EthereumBridge.TransactOpts, args, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0x846704f8.
//
// Solidity: function initialize(uint256 crossChainID, address cross, uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeTransactor) Initialize(opts *bind.TransactOpts, crossChainID *big.Int, cross common.Address, _threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "initialize", crossChainID, cross, _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Initialize is a paid mutator transaction binding the contract method 0x846704f8.
//
// Solidity: function initialize(uint256 crossChainID, address cross, uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeSession) Initialize(crossChainID *big.Int, cross common.Address, _threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, crossChainID, cross, _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Initialize is a paid mutator transaction binding the contract method 0x846704f8.
//
// Solidity: function initialize(uint256 crossChainID, address cross, uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode, address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Initialize(crossChainID *big.Int, cross common.Address, _threshold uint8, _rewardWallet common.Address, _crossMintableERC20FactoryCode common.Address, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, crossChainID, cross, _threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EthereumBridge *EthereumBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EthereumBridge *EthereumBridgeSession) Pause() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Pause(&_EthereumBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Pause(&_EthereumBridge.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactor) PauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "pauseToken", remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PauseToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0xd2ff130d.
//
// Solidity: function pauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) PauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PauseToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "permitBridgeToken", remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeToken(&_EthereumBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x7f260a75.
//
// Solidity: function permitBridgeToken(uint256 remoteChainID, address token, address from, address to, uint256 value, uint256 gasFee, uint256 exFee, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) PermitBridgeToken(remoteChainID *big.Int, token common.Address, from common.Address, to common.Address, value *big.Int, gasFee *big.Int, exFee *big.Int, permitArgs IStandardBridgePermitArguments, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeToken(&_EthereumBridge.TransactOpts, remoteChainID, token, from, to, value, gasFee, exFee, permitArgs, extraData)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RegisterToken(&_EthereumBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RegisterToken(&_EthereumBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_EthereumBridge *EthereumBridgeTransactor) RemoveFeeStation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "removeFeeStation")
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_EthereumBridge *EthereumBridgeSession) RemoveFeeStation() (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveFeeStation(&_EthereumBridge.TransactOpts)
}

// RemoveFeeStation is a paid mutator transaction binding the contract method 0xd7c82f32.
//
// Solidity: function removeFeeStation() returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RemoveFeeStation() (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveFeeStation(&_EthereumBridge.TransactOpts)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveValidator(&_EthereumBridge.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveValidator(&_EthereumBridge.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveValidators(&_EthereumBridge.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemoveValidators(&_EthereumBridge.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumBridge *EthereumBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumBridge *EthereumBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthereumBridge.Contract.RenounceOwnership(&_EthereumBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthereumBridge.Contract.RenounceOwnership(&_EthereumBridge.TransactOpts)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ResetValidators(&_EthereumBridge.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ResetValidators(&_EthereumBridge.TransactOpts, validators)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "retryFinalizeBridge", remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridge(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridge(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridgeBatch(&_EthereumBridge.TransactOpts, remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridgeBatch(&_EthereumBridge.TransactOpts, remoteChainID, indexes)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetChain(opts *bind.TransactOpts, remoteChainID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setChain", remoteChainID)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_EthereumBridge *EthereumBridgeSession) SetChain(remoteChainID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetChain(&_EthereumBridge.TransactOpts, remoteChainID)
}

// SetChain is a paid mutator transaction binding the contract method 0x1b44fd15.
//
// Solidity: function setChain(uint256 remoteChainID) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetChain(remoteChainID *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetChain(&_EthereumBridge.TransactOpts, remoteChainID)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetFeeStation(opts *bind.TransactOpts, _bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setFeeStation", _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetFeeStation(&_EthereumBridge.TransactOpts, _bridgeFeeStation)
}

// SetFeeStation is a paid mutator transaction binding the contract method 0x54db0126.
//
// Solidity: function setFeeStation(address _bridgeFeeStation) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetFeeStation(_bridgeFeeStation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetFeeStation(&_EthereumBridge.TransactOpts, _bridgeFeeStation)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetRewardWallet(opts *bind.TransactOpts, rewardWallet_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setRewardWallet", rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_EthereumBridge *EthereumBridgeSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetRewardWallet(&_EthereumBridge.TransactOpts, rewardWallet_)
}

// SetRewardWallet is a paid mutator transaction binding the contract method 0x5958621e.
//
// Solidity: function setRewardWallet(address rewardWallet_) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetRewardWallet(rewardWallet_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetRewardWallet(&_EthereumBridge.TransactOpts, rewardWallet_)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetValidator(&_EthereumBridge.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetValidator(&_EthereumBridge.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetValidators(&_EthereumBridge.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetValidators(&_EthereumBridge.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumBridge *EthereumBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumBridge *EthereumBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.TransferOwnership(&_EthereumBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.TransferOwnership(&_EthereumBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EthereumBridge *EthereumBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EthereumBridge *EthereumBridgeSession) Unpause() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Unpause(&_EthereumBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Unpause(&_EthereumBridge.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactor) UnpauseToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "unpauseToken", remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UnpauseToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x84d58d42.
//
// Solidity: function unpauseToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) UnpauseToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UnpauseToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UnregisterToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UnregisterToken(&_EthereumBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthereumBridge *EthereumBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthereumBridge *EthereumBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UpgradeToAndCall(&_EthereumBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.UpgradeToAndCall(&_EthereumBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthereumBridge *EthereumBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthereumBridge *EthereumBridgeSession) Receive() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Receive(&_EthereumBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _EthereumBridge.Contract.Receive(&_EthereumBridge.TransactOpts)
}

// EthereumBridgeBridgeFeeChargedIterator is returned from FilterBridgeFeeCharged and is used to iterate over the raw logs and unpacked data for BridgeFeeCharged events raised by the EthereumBridge contract.
type EthereumBridgeBridgeFeeChargedIterator struct {
	Event *EthereumBridgeBridgeFeeCharged // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgeFeeCharged)
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
		it.Event = new(EthereumBridgeBridgeFeeCharged)
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
func (it *EthereumBridgeBridgeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgeFeeCharged represents a BridgeFeeCharged event raised by the EthereumBridge contract.
type EthereumBridgeBridgeFeeCharged struct {
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeFeeCharged(opts *bind.FilterOpts, index []*big.Int, token []common.Address, account []common.Address) (*EthereumBridgeBridgeFeeChargedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeFeeChargedIterator{contract: _EthereumBridge.contract, event: "BridgeFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeCharged is a free log subscription operation binding the contract event 0x1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d.
//
// Solidity: event BridgeFeeCharged(uint256 indexed index, address indexed token, address indexed account, uint256 gasFee, uint256 exFee)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeFeeCharged(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeFeeCharged, index []*big.Int, token []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeFeeCharged", indexRule, tokenRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgeFeeCharged)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeFeeCharged(log types.Log) (*EthereumBridgeBridgeFeeCharged, error) {
	event := new(EthereumBridgeBridgeFeeCharged)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeBridgeFinalizeRevertedIterator is returned from FilterBridgeFinalizeReverted and is used to iterate over the raw logs and unpacked data for BridgeFinalizeReverted events raised by the EthereumBridge contract.
type EthereumBridgeBridgeFinalizeRevertedIterator struct {
	Event *EthereumBridgeBridgeFinalizeReverted // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgeFinalizeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgeFinalizeReverted)
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
		it.Event = new(EthereumBridgeBridgeFinalizeReverted)
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
func (it *EthereumBridgeBridgeFinalizeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgeFinalizeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgeFinalizeReverted represents a BridgeFinalizeReverted event raised by the EthereumBridge contract.
type EthereumBridgeBridgeFinalizeReverted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalizeReverted is a free log retrieval operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeFinalizeReverted(opts *bind.FilterOpts, index []*big.Int) (*EthereumBridgeBridgeFinalizeRevertedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeFinalizeRevertedIterator{contract: _EthereumBridge.contract, event: "BridgeFinalizeReverted", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalizeReverted is a free log subscription operation binding the contract event 0xfb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f.
//
// Solidity: event BridgeFinalizeReverted(uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeFinalizeReverted(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeFinalizeReverted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeFinalizeReverted", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgeFinalizeReverted)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeFinalizeReverted(log types.Log) (*EthereumBridgeBridgeFinalizeReverted, error) {
	event := new(EthereumBridgeBridgeFinalizeReverted)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFinalizeReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the EthereumBridge contract.
type EthereumBridgeBridgeFinalizedIterator struct {
	Event *EthereumBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgeFinalized)
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
		it.Event = new(EthereumBridgeBridgeFinalized)
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
func (it *EthereumBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgeFinalized represents a BridgeFinalized event raised by the EthereumBridge contract.
type EthereumBridgeBridgeFinalized struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Token         common.Address
	To            common.Address
	Value         *big.Int
	Time          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (*EthereumBridgeBridgeFinalizedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeFinalizedIterator{contract: _EthereumBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeFinalized, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgeFinalized)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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

// ParseBridgeFinalized is a log parse operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeFinalized(log types.Log) (*EthereumBridgeBridgeFinalized, error) {
	event := new(EthereumBridgeBridgeFinalized)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the EthereumBridge contract.
type EthereumBridgeBridgeInitiatedIterator struct {
	Event *EthereumBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgeInitiated)
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
		it.Event = new(EthereumBridgeBridgeInitiated)
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
func (it *EthereumBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgeInitiated represents a BridgeInitiated event raised by the EthereumBridge contract.
type EthereumBridgeBridgeInitiated struct {
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, from []common.Address) (*EthereumBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeInitiatedIterator{contract: _EthereumBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d7.
//
// Solidity: event BridgeInitiated(uint256 indexed remoteChainID, uint256 indexed index, address localToken, address remoteToken, address indexed from, address to, uint256 value, uint256 time, bool permit, bytes extraData)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeInitiated, remoteChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeInitiated", remoteChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgeInitiated)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeInitiated(log types.Log) (*EthereumBridgeBridgeInitiated, error) {
	event := new(EthereumBridgeBridgeInitiated)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeChainSetIterator is returned from FilterChainSet and is used to iterate over the raw logs and unpacked data for ChainSet events raised by the EthereumBridge contract.
type EthereumBridgeChainSetIterator struct {
	Event *EthereumBridgeChainSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeChainSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeChainSet)
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
		it.Event = new(EthereumBridgeChainSet)
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
func (it *EthereumBridgeChainSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeChainSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeChainSet represents a ChainSet event raised by the EthereumBridge contract.
type EthereumBridgeChainSet struct {
	RemoteChainID *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainSet is a free log retrieval operation binding the contract event 0x0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a.
//
// Solidity: event ChainSet(uint256 indexed remoteChainID)
func (_EthereumBridge *EthereumBridgeFilterer) FilterChainSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*EthereumBridgeChainSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "ChainSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeChainSetIterator{contract: _EthereumBridge.contract, event: "ChainSet", logs: logs, sub: sub}, nil
}

// WatchChainSet is a free log subscription operation binding the contract event 0x0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a.
//
// Solidity: event ChainSet(uint256 indexed remoteChainID)
func (_EthereumBridge *EthereumBridgeFilterer) WatchChainSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeChainSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "ChainSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeChainSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "ChainSet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseChainSet(log types.Log) (*EthereumBridgeChainSet, error) {
	event := new(EthereumBridgeChainSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "ChainSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EthereumBridge contract.
type EthereumBridgeEIP712DomainChangedIterator struct {
	Event *EthereumBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeEIP712DomainChanged)
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
		it.Event = new(EthereumBridgeEIP712DomainChanged)
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
func (it *EthereumBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the EthereumBridge contract.
type EthereumBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EthereumBridge *EthereumBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EthereumBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeEIP712DomainChangedIterator{contract: _EthereumBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EthereumBridge *EthereumBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EthereumBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeEIP712DomainChanged)
				if err := _EthereumBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*EthereumBridgeEIP712DomainChanged, error) {
	event := new(EthereumBridgeEIP712DomainChanged)
	if err := _EthereumBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeFeeStationSetIterator is returned from FilterFeeStationSet and is used to iterate over the raw logs and unpacked data for FeeStationSet events raised by the EthereumBridge contract.
type EthereumBridgeFeeStationSetIterator struct {
	Event *EthereumBridgeFeeStationSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeFeeStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeFeeStationSet)
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
		it.Event = new(EthereumBridgeFeeStationSet)
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
func (it *EthereumBridgeFeeStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeFeeStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeFeeStationSet represents a FeeStationSet event raised by the EthereumBridge contract.
type EthereumBridgeFeeStationSet struct {
	FeeStation common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeStationSet is a free log retrieval operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address feeStation)
func (_EthereumBridge *EthereumBridgeFilterer) FilterFeeStationSet(opts *bind.FilterOpts) (*EthereumBridgeFeeStationSetIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "FeeStationSet")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeFeeStationSetIterator{contract: _EthereumBridge.contract, event: "FeeStationSet", logs: logs, sub: sub}, nil
}

// WatchFeeStationSet is a free log subscription operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address feeStation)
func (_EthereumBridge *EthereumBridgeFilterer) WatchFeeStationSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeFeeStationSet) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "FeeStationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeFeeStationSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
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

// ParseFeeStationSet is a log parse operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address feeStation)
func (_EthereumBridge *EthereumBridgeFilterer) ParseFeeStationSet(log types.Log) (*EthereumBridgeFeeStationSet, error) {
	event := new(EthereumBridgeFeeStationSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EthereumBridge contract.
type EthereumBridgeInitializedIterator struct {
	Event *EthereumBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeInitialized)
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
		it.Event = new(EthereumBridgeInitialized)
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
func (it *EthereumBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeInitialized represents a Initialized event raised by the EthereumBridge contract.
type EthereumBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EthereumBridge *EthereumBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*EthereumBridgeInitializedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeInitializedIterator{contract: _EthereumBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EthereumBridge *EthereumBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EthereumBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeInitialized)
				if err := _EthereumBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseInitialized(log types.Log) (*EthereumBridgeInitialized, error) {
	event := new(EthereumBridgeInitialized)
	if err := _EthereumBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthereumBridge contract.
type EthereumBridgeOwnershipTransferredIterator struct {
	Event *EthereumBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeOwnershipTransferred)
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
		it.Event = new(EthereumBridgeOwnershipTransferred)
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
func (it *EthereumBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the EthereumBridge contract.
type EthereumBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthereumBridge *EthereumBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthereumBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeOwnershipTransferredIterator{contract: _EthereumBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthereumBridge *EthereumBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthereumBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeOwnershipTransferred)
				if err := _EthereumBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*EthereumBridgeOwnershipTransferred, error) {
	event := new(EthereumBridgeOwnershipTransferred)
	if err := _EthereumBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthereumBridge contract.
type EthereumBridgePausedIterator struct {
	Event *EthereumBridgePaused // Event containing the contract specifics and raw log

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
func (it *EthereumBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgePaused)
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
		it.Event = new(EthereumBridgePaused)
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
func (it *EthereumBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgePaused represents a Paused event raised by the EthereumBridge contract.
type EthereumBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthereumBridge *EthereumBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*EthereumBridgePausedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgePausedIterator{contract: _EthereumBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthereumBridge *EthereumBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthereumBridgePaused) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgePaused)
				if err := _EthereumBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParsePaused(log types.Log) (*EthereumBridgePaused, error) {
	event := new(EthereumBridgePaused)
	if err := _EthereumBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeRewardWalletSetIterator is returned from FilterRewardWalletSet and is used to iterate over the raw logs and unpacked data for RewardWalletSet events raised by the EthereumBridge contract.
type EthereumBridgeRewardWalletSetIterator struct {
	Event *EthereumBridgeRewardWalletSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeRewardWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeRewardWalletSet)
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
		it.Event = new(EthereumBridgeRewardWalletSet)
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
func (it *EthereumBridgeRewardWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeRewardWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeRewardWalletSet represents a RewardWalletSet event raised by the EthereumBridge contract.
type EthereumBridgeRewardWalletSet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardWalletSet is a free log retrieval operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address wallet)
func (_EthereumBridge *EthereumBridgeFilterer) FilterRewardWalletSet(opts *bind.FilterOpts) (*EthereumBridgeRewardWalletSetIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "RewardWalletSet")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeRewardWalletSetIterator{contract: _EthereumBridge.contract, event: "RewardWalletSet", logs: logs, sub: sub}, nil
}

// WatchRewardWalletSet is a free log subscription operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address wallet)
func (_EthereumBridge *EthereumBridgeFilterer) WatchRewardWalletSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeRewardWalletSet) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "RewardWalletSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeRewardWalletSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
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

// ParseRewardWalletSet is a log parse operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address wallet)
func (_EthereumBridge *EthereumBridgeFilterer) ParseRewardWalletSet(log types.Log) (*EthereumBridgeRewardWalletSet, error) {
	event := new(EthereumBridgeRewardWalletSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the EthereumBridge contract.
type EthereumBridgeThresholdChangedIterator struct {
	Event *EthereumBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeThresholdChanged)
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
		it.Event = new(EthereumBridgeThresholdChanged)
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
func (it *EthereumBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeThresholdChanged represents a ThresholdChanged event raised by the EthereumBridge contract.
type EthereumBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_EthereumBridge *EthereumBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*EthereumBridgeThresholdChangedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeThresholdChangedIterator{contract: _EthereumBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_EthereumBridge *EthereumBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *EthereumBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeThresholdChanged)
				if err := _EthereumBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseThresholdChanged(log types.Log) (*EthereumBridgeThresholdChanged, error) {
	event := new(EthereumBridgeThresholdChanged)
	if err := _EthereumBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeTokenPairPausedIterator is returned from FilterTokenPairPaused and is used to iterate over the raw logs and unpacked data for TokenPairPaused events raised by the EthereumBridge contract.
type EthereumBridgeTokenPairPausedIterator struct {
	Event *EthereumBridgeTokenPairPaused // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeTokenPairPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeTokenPairPaused)
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
		it.Event = new(EthereumBridgeTokenPairPaused)
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
func (it *EthereumBridgeTokenPairPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeTokenPairPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeTokenPairPaused represents a TokenPairPaused event raised by the EthereumBridge contract.
type EthereumBridgeTokenPairPaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairPaused is a free log retrieval operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_EthereumBridge *EthereumBridgeFilterer) FilterTokenPairPaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*EthereumBridgeTokenPairPausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTokenPairPausedIterator{contract: _EthereumBridge.contract, event: "TokenPairPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairPaused is a free log subscription operation binding the contract event 0xf98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3.
//
// Solidity: event TokenPairPaused(uint256 indexed remoteChainID, address indexed token)
func (_EthereumBridge *EthereumBridgeFilterer) WatchTokenPairPaused(opts *bind.WatchOpts, sink chan<- *EthereumBridgeTokenPairPaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "TokenPairPaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeTokenPairPaused)
				if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPairPaused(log types.Log) (*EthereumBridgeTokenPairPaused, error) {
	event := new(EthereumBridgeTokenPairPaused)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the EthereumBridge contract.
type EthereumBridgeTokenPairRegisteredIterator struct {
	Event *EthereumBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeTokenPairRegistered)
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
		it.Event = new(EthereumBridgeTokenPairRegistered)
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
func (it *EthereumBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the EthereumBridge contract.
type EthereumBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	IsOrigin      bool
	LocalToken    common.Address
	RemoteToken   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0x78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken)
func (_EthereumBridge *EthereumBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*EthereumBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTokenPairRegisteredIterator{contract: _EthereumBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0x78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, bool isOrigin, address indexed localToken, address indexed remoteToken)
func (_EthereumBridge *EthereumBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *EthereumBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeTokenPairRegistered)
				if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*EthereumBridgeTokenPairRegistered, error) {
	event := new(EthereumBridgeTokenPairRegistered)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeTokenPairUnpausedIterator is returned from FilterTokenPairUnpaused and is used to iterate over the raw logs and unpacked data for TokenPairUnpaused events raised by the EthereumBridge contract.
type EthereumBridgeTokenPairUnpausedIterator struct {
	Event *EthereumBridgeTokenPairUnpaused // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeTokenPairUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeTokenPairUnpaused)
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
		it.Event = new(EthereumBridgeTokenPairUnpaused)
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
func (it *EthereumBridgeTokenPairUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeTokenPairUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeTokenPairUnpaused represents a TokenPairUnpaused event raised by the EthereumBridge contract.
type EthereumBridgeTokenPairUnpaused struct {
	RemoteChainID *big.Int
	Token         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnpaused is a free log retrieval operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_EthereumBridge *EthereumBridgeFilterer) FilterTokenPairUnpaused(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*EthereumBridgeTokenPairUnpausedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTokenPairUnpausedIterator{contract: _EthereumBridge.contract, event: "TokenPairUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnpaused is a free log subscription operation binding the contract event 0xac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9.
//
// Solidity: event TokenPairUnpaused(uint256 indexed remoteChainID, address indexed token)
func (_EthereumBridge *EthereumBridgeFilterer) WatchTokenPairUnpaused(opts *bind.WatchOpts, sink chan<- *EthereumBridgeTokenPairUnpaused, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "TokenPairUnpaused", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeTokenPairUnpaused)
				if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPairUnpaused(log types.Log) (*EthereumBridgeTokenPairUnpaused, error) {
	event := new(EthereumBridgeTokenPairUnpaused)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the EthereumBridge contract.
type EthereumBridgeTokenPairUnregisteredIterator struct {
	Event *EthereumBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeTokenPairUnregistered)
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
		it.Event = new(EthereumBridgeTokenPairUnregistered)
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
func (it *EthereumBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the EthereumBridge contract.
type EthereumBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_EthereumBridge *EthereumBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*EthereumBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTokenPairUnregisteredIterator{contract: _EthereumBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_EthereumBridge *EthereumBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *EthereumBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeTokenPairUnregistered)
				if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*EthereumBridgeTokenPairUnregistered, error) {
	event := new(EthereumBridgeTokenPairUnregistered)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthereumBridge contract.
type EthereumBridgeUnpausedIterator struct {
	Event *EthereumBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeUnpaused)
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
		it.Event = new(EthereumBridgeUnpaused)
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
func (it *EthereumBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeUnpaused represents a Unpaused event raised by the EthereumBridge contract.
type EthereumBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthereumBridge *EthereumBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthereumBridgeUnpausedIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeUnpausedIterator{contract: _EthereumBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthereumBridge *EthereumBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthereumBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeUnpaused)
				if err := _EthereumBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseUnpaused(log types.Log) (*EthereumBridgeUnpaused, error) {
	event := new(EthereumBridgeUnpaused)
	if err := _EthereumBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EthereumBridge contract.
type EthereumBridgeUpgradedIterator struct {
	Event *EthereumBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeUpgraded)
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
		it.Event = new(EthereumBridgeUpgraded)
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
func (it *EthereumBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeUpgraded represents a Upgraded event raised by the EthereumBridge contract.
type EthereumBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EthereumBridge *EthereumBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EthereumBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeUpgradedIterator{contract: _EthereumBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EthereumBridge *EthereumBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EthereumBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeUpgraded)
				if err := _EthereumBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseUpgraded(log types.Log) (*EthereumBridgeUpgraded, error) {
	event := new(EthereumBridgeUpgraded)
	if err := _EthereumBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the EthereumBridge contract.
type EthereumBridgeValidatorSetIterator struct {
	Event *EthereumBridgeValidatorSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeValidatorSet)
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
		it.Event = new(EthereumBridgeValidatorSet)
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
func (it *EthereumBridgeValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeValidatorSet represents a ValidatorSet event raised by the EthereumBridge contract.
type EthereumBridgeValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_EthereumBridge *EthereumBridgeFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*EthereumBridgeValidatorSetIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeValidatorSetIterator{contract: _EthereumBridge.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_EthereumBridge *EthereumBridgeFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeValidatorSet) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeValidatorSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseValidatorSet(log types.Log) (*EthereumBridgeValidatorSet, error) {
	event := new(EthereumBridgeValidatorSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
