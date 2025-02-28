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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossMintableERC20FactoryCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"setChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"ChainSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"feeStation\",\"type\":\"address\"}],\"name\":\"FeeStationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RewardWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryAleadyExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryAleadyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errBridgeRegistryCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryInsufficientBridgeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errBridgeRegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"errBridgeRegistryTokenFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errBridgeRegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errStandardBridgeFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161550c6100f95f395f81816126cb015281816126f40152612838015261550c5ff3fe6080604052600436106102dc575f3560e01c806384b0196e11610189578063b7f3358d116100d8578063edbbea3911610092578063f45096431161006d578063f450964314610977578063f698da2514610996578063facd743b146109aa578063fb75b2c7146109c9575f5ffd5b8063edbbea3914610918578063f2fde38b14610937578063f30589c314610956575f5ffd5b8063b7f3358d14610867578063cbae595814610886578063cf56118e146108a5578063d2ff130d146108c6578063d5717fc5146108e5578063d7c82f3214610904575f5ffd5b80639300c92611610143578063ad3cb1cc1161011e578063ad3cb1cc146107ca578063ae6893f8146107fa578063ae76638914610819578063aed1d40314610853575f5ffd5b80639300c9261461076b578063952a95de1461078a57806396ce0795146107b6575f5ffd5b806384b0196e146106a457806384d58d42146106cb57806388d67d6d146106ea5780638da5cb5b146106fd5780638f517c171461073957806391cf6d3e14610757575f5ffd5b806354db012611610245578063670e6268116101ff5780637f260a75116101da5780637f260a751461058a578063814914b51461059d5780638415a385146106645780638456cb5914610690575f5ffd5b8063670e6268146105385780637101fcd314610557578063715018a614610576575f5ffd5b806354db0126146104795780635958621e146104985780635b605f5c146104b75780635c975abb146104e35780635fd262de14610506578063607375a114610519575f5ffd5b80633f4ba83a116102965780633f4ba83a146103b957806340a141ff146103cd57806342cde4e8146103ec57806347666cb11461040d5780634f1ef2861461044457806352d1902d14610457575f5ffd5b8063030372c3146102f65780631327d3d81461032a5780631938e0f2146103495780631b44fd151461035c5780631d40f0d81461037b5780633960e7871461039a575f5ffd5b366102f257345f036102f0576102f0614153565b005b5f5ffd5b348015610301575f5ffd5b506103156103103660046141ee565b6109e6565b60405190151581526020015b60405180910390f35b348015610335575f5ffd5b506102f06103443660046142aa565b610a2a565b6103156103573660046143b9565b610a38565b348015610367575f5ffd5b506102f0610376366004614472565b610d44565b348015610386575f5ffd5b506102f0610395366004614489565b610da7565b3480156103a5575f5ffd5b506103156103b436600461452e565b610de1565b3480156103c4575f5ffd5b506102f0610ed8565b3480156103d8575f5ffd5b506102f06103e73660046142aa565b610eea565b3480156103f7575f5ffd5b5060355460405160ff9091168152602001610321565b348015610418575f5ffd5b5060665461042c906001600160a01b031681565b6040516001600160a01b039091168152602001610321565b6102f06104523660046145b7565b610ef4565b348015610462575f5ffd5b5061046b610f0f565b604051908152602001610321565b348015610484575f5ffd5b506102f06104933660046142aa565b610f2b565b3480156104a3575f5ffd5b506102f06104b23660046142aa565b610fd2565b3480156104c2575f5ffd5b506104d66104d1366004614472565b61106e565b6040516103219190614655565b3480156104ee575f5ffd5b505f5160206154975f395f51905f525460ff16610315565b6103156105143660046146df565b6111c7565b348015610524575f5ffd5b506102f0610533366004614767565b6112d6565b348015610543575f5ffd5b5061042c6105523660046147be565b6113e8565b348015610562575f5ffd5b506102f0610571366004614489565b6114f9565b348015610581575f5ffd5b506102f061150f565b610315610598366004614836565b611520565b3480156105a8575f5ffd5b506106576105b736600461495a565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9182526003602090815260408084206001600160a01b039384168552600901825292839020835160a08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260020154608082015290565b6040516103219190614988565b34801561066f575f5ffd5b5061068361067e36600461452e565b61162b565b60405161032191906149c4565b34801561069b575f5ffd5b506102f06116d9565b3480156106af575f5ffd5b506106b86116e9565b6040516103219796959493929190614a10565b3480156106d6575f5ffd5b506102f06106e536600461495a565b611792565b6103156106f8366004614b93565b611886565b348015610708575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661042c565b348015610744575f5ffd5b505f5461042c906001600160a01b031681565b348015610762575f5ffd5b5060685461046b565b348015610776575f5ffd5b506102f0610785366004614489565b611921565b348015610795575f5ffd5b506107a96107a436600461452e565b611958565b6040516103219190614c7b565b3480156107c1575f5ffd5b5061046b611a9b565b3480156107d5575f5ffd5b50610683604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610805575f5ffd5b5061046b610814366004614472565b611b0b565b348015610824575f5ffd5b50610838610833366004614cd7565b611b27565b60408051938452602084019290925290820152606001610321565b34801561085e575f5ffd5b5061046b611bb6565b348015610872575f5ffd5b506102f0610881366004614d0c565b611bc1565b348015610891575f5ffd5b5061042c6108a0366004614472565b611c0b565b3480156108b0575f5ffd5b506108b9611c17565b6040516103219190614d25565b3480156108d1575f5ffd5b506102f06108e036600461495a565b611c23565b3480156108f0575f5ffd5b5061046b6108ff366004614472565b611d1e565b34801561090f575f5ffd5b506102f0611d3a565b348015610923575f5ffd5b506102f0610932366004614d44565b611dd3565b348015610942575f5ffd5b506102f06109513660046142aa565b611f7a565b348015610961575f5ffd5b5061096a611fb4565b6040516103219190614d69565b348015610982575f5ffd5b506102f061099136600461495a565b611fc0565b3480156109a1575f5ffd5b5061046b6120e6565b3480156109b5575f5ffd5b506103156109c43660046142aa565b6120ef565b3480156109d4575f5ffd5b506067546001600160a01b031661042c565b5f805b8251811015610a1e57610a1584848381518110610a0857610a08614da9565b6020026020010151610de1565b506001016109e9565b50600190505b92915050565b610a358160016120fb565b50565b5f610a416121c4565b8435610a5360608701604088016142aa565b5f828152600360205260409020610a6d90600701826121f4565b8190610a9d576040516367240f3b60e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff1615610afc5760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b50610b05612218565b5f348015610b2f5760405163341382c560e11b815260048101929092526024820152604401610a94565b50610b3c90508735611d1e565b602088013514610b4c8835611d1e565b88602001359091610b79576040516329846bd360e21b815260048101929092526024820152604401610a94565b50610c0b90507fb2b56073c3812af4a57f2830cbc00b1dd751f01c9c75ccee5c7f4efa28f8d89f6020890135610bb560608b0160408c016142aa565b610bc560808c0160608d016142aa565b60808c0135610bd760a08e018e614dbd565b604051602001610bed9796959493929190614e27565b6040516020818303038152906040528051906020012087878761224f565b610c2987355f90815260036020526040902060020180546001019055565b5f80610c5a8935610c4060608c0160408d016142aa565b610c5060808d0160608e016142aa565b8c6080013561244a565b915091508115610ce557610c7460808a0160608b016142aa565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610cb560608e0160408f016142aa565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610d1e565b610cef89826124c7565b60405160208a0135907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610d3a60015f5160206154b75f395f51905f5255565b5050949350505050565b610d4c61255d565b610d576001826125b8565b8190610d7957604051633d1bcfa960e01b8152600401610a9491815260200190565b5060405181907f0e394b8ef4f476f993ade67cc80cccc25089b07af7684fc27fecd74c3bc97d1a905f90a250565b5f5b8151811015610ddd57610dd5828281518110610dc757610dc7614da9565b60200260200101515f6120fb565b600101610da9565b5050565b5f610dea6121c4565b610df2612218565b5f610dfd8484611958565b90505f5f610e198684604001518560600151866080015161244a565b91509150818190610e3d5760405162461bcd60e51b8152600401610a9491906149c4565b50610e4886866125c3565b82606001516001600160a01b03168360200151877f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b8660400151876080015142604051610eb3939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a460019350505050610a2460015f5160206154b75f395f51905f5255565b610ee061255d565b610ee8612667565b565b610a35815f6120fb565b610efc6126c0565b610f0582612764565b610ddd828261276c565b5f610f1861282d565b505f5160206154775f395f51905f525b90565b610f3361255d565b6001600160a01b038116610f7d576040516211211f60e21b81526020600482015260116024820152702fb13934b233b2a332b2a9ba30ba34b7b760791b6044820152606401610a94565b606680546001600160a01b0319166001600160a01b0383169081179091556040519081527f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444906020015b60405180910390a150565b610fda61255d565b6001600160a01b038116611020576040516211211f60e21b815260206004820152600d60248201526c72657761726457616c6c65745f60981b6044820152606401610a94565b606780546001600160a01b0319166001600160a01b0383169081179091556040519081527f5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b90602001610fc7565b5f81815260036020526040812060609161108a60078301612876565b90505f81516001600160401b038111156110a6576110a6614167565b6040519080825280602002602001820160405280156110fd57816020015b6040805160a0810182525f808252602080830182905292820181905260608201819052608082015282525f199092019101816110c45790505b5090505f5b82518110156111be57836009015f84838151811061112257611122614da9565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160a08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b90930490911615156060820152600290910154608082015282518390839081106111ab576111ab614da9565b6020908102919091010152600101611102565b50949350505050565b5f6111d06121c4565b5f898152600360205260409020899089906111ee90600701826121f4565b8190611219576040516367240f3b60e01b81526001600160a01b039091166004820152602401610a94565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156112785760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b50611281612218565b5f5f6112908d8d8c8c8c612882565b915091506112ac8d8d6112a03390565b8e8e87875f8f8f61299b565b6001945050506112c860015f5160206154b75f395f51905f5255565b505098975050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561131a5750825b90505f826001600160401b031660011480156113355750303b155b905081158015611343575080155b156113615760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561138b57845460ff60401b1916600160401b1785555b611397898989896129e9565b83156113dd57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f6113f161255d565b5f546001600160a01b031661141957604051636274ce8960e11b815260040160405180910390fd5b5f546040516bffffffffffffffffffffffff19606087901b1660208201526001600160a01b0390911690634804a0419060340160405160208183030381529060405280519060200120856040516020016114739190614e8b565b60405160208183030381529060405286866040518563ffffffff1660e01b81526004016114a39493929190614eac565b6020604051808303815f875af11580156114bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114e39190614eeb565b90506114f1855f8387611dd3565b949350505050565b6115066103956033612876565b610a3581611921565b61151761255d565b610ee85f612abd565b5f6115296121c4565b5f8b81526003602052604090208b908b9061154790600701826121f4565b8190611572576040516367240f3b60e01b81526001600160a01b039091166004820152602401610a94565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156115d15760405163620ea8c560e01b81526001600160a01b039091166004820152602401610a94565b506115da612218565b5f5f6115e98f8f8d8d8d612882565b915091506115ff8f8f8f8f8f87878f8f8f612b2d565b60019450505061161b60015f5160206154b75f395f51905f5255565b50509a9950505050505050505050565b5f828152600360209081526040808320848452600401909152902080546060919061165590614f06565b80601f016020809104026020016040519081016040528092919081815260200182805461168190614f06565b80156116cc5780601f106116a3576101008083540402835291602001916116cc565b820191905f5260205f20905b8154815290600101906020018083116116af57829003601f168201915b5050505050905092915050565b6116e161255d565b610ee8612b98565b5f60608082808083815f5160206154575f395f51905f52805490915015801561171457506001810154155b6117585760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a94565b611760612be0565b611768612ca0565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61179a61255d565b5f8281526003602052604090206117b460078201836121f4565b82906117df57604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff166118325760405163273bf77b60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b191690555185917fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e991a3505050565b5f805b858110156119145761190b8787838181106118a6576118a6614da9565b90506020028101906118b89190614f3e565b8683815181106118ca576118ca614da9565b60200260200101518684815181106118e4576118e4614da9565b60200260200101518685815181106118fe576118fe614da9565b6020026020010151610a38565b50600101611889565b5060019695505050505050565b5f5b8151811015610ddd5761195082828151811061194157611941614da9565b602002602001015160016120fb565b600101611923565b61199e6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5f8381526003602081815260408084208685528301825292839020835160c0810185528154815260018201549281019290925260028101546001600160a01b03908116948301949094529182015490921660608301526004810154608083015260058101805460a084019190611a1390614f06565b80601f0160208091040260200160405190810160405280929190818152602001828054611a3f90614f06565b8015611a8a5780601f10611a6157610100808354040283529160200191611a8a565b820191905f5260205f20905b815481529060010190602001808311611a6d57829003601f168201915b505050505081525050905092915050565b606654604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa158015611ae2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b069190614f5c565b905090565b5f818152600360205260408120600190810154610a2491614f87565b60665460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa158015611b83573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ba79190614f9a565b91989097509095509350505050565b5f611b066033612cde565b611bc961255d565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff90602001610fc7565b5f610a24603383612ce7565b6060611b066001612876565b611c2b61255d565b5f828152600360205260409020611c4560078201836121f4565b8290611c7057604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f9081526009820160205260409020600101548290600160a81b900460ff1615611cc45760405163283e1c6360e21b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f818152600983016020526040808220600101805460ff60a81b1916600160a81b1790555185917ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a391a3505050565b5f81815260036020526040812060020154610a24906001614f87565b611d4261255d565b6066546001600160a01b0316611d8e57604051635f6c3f4f60e01b815260206004820152601060248201526f313934b233b2a332b2a9ba30ba34b7b760811b6044820152606401610a94565b606680546001600160a01b03191690556040515f81527f68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea4449060200160405180910390a1565b611ddb61255d565b611de6600185612cf2565b8490611e085760405163c45f2c2360e01b8152600401610a9491815260200190565b505f8481526003602052604090206001600160a01b038316611e5a57604051636e0b3c5f60e11b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a94565b611e676007820184612d09565b8390611e9257604051634d2df28360e01b81526001600160a01b039091166004820152602401610a94565b506040805160a0810182526001600160a01b0385811680835285821660208085018281528a15158688018181525f6060890181815260808a0182815288835260098d018752918b902099518a546001600160a01b031916908a16178a55935160018a01805493519551919099166001600160a81b031990931692909217600160a01b941515949094029390931760ff60a81b1916600160a81b91151591909102179095555160029095019490945593519182529188917f78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c910160405180910390a45050505050565b611f8261255d565b6001600160a01b038116611fab57604051631e4fbdf760e01b81525f6004820152602401610a94565b610a3581612abd565b6060611b066033612876565b611fc861255d565b611fd3600183612cf2565b8290611ff55760405163c45f2c2360e01b8152600401610a9491815260200190565b506001600160a01b03811661203557604051636e0b3c5f60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610a94565b5f82815260036020526040902061204f6007820183612d1d565b829061207a57604051632e91678d60e11b81526001600160a01b039091166004820152602401610a94565b506001600160a01b0382165f81815260098301602052604080822080546001600160a01b03191681556001810180546001600160b01b03191690556002018290555185917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d91a3505050565b5f611b06612d31565b5f610a246033836121f4565b61210361255d565b801561214557612114603383612d09565b829061213f576040516329a04e7760e21b81526001600160a01b039091166004820152602401610a94565b5061217d565b612150603383612d1d565b829061217b5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610a94565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f5160206154975f395f51905f525460ff1615610ee85760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f90815260018301602052604081205415155b9392505050565b5f5160206154b75f395f51905f5280546001190161224957604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b82518251811480156122615750815181145b835183518392612295576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610a94565b505060355482915060ff168110156122c357604051632fcba65760e11b8152600401610a9491815260200190565b505f80826001600160401b038111156122de576122de614167565b604051908082528060200260200182016040528015612307578160200160208202803683370190505b5090505f5b83811015612414575f61237788838151811061232a5761232a614da9565b602002602001015188848151811061234457612344614da9565b602002602001015188858151811061235e5761235e614da9565b602002602001015161236f8d612d3a565b929190612d66565b9050612382816120ef565b81906123ad5760405163845a09e760e01b81526001600160a01b039091166004820152602401610a94565b505f805b84518110156123fd578481815181106123cc576123cc614da9565b60200260200101516001600160a01b0316836001600160a01b0316036123f557600191506123fd565b6001016123b1565b508061240a578460010194505b505060010161230c565b50603554829060ff1681101561244057604051632fcba65760e11b8152600401610a9491815260200190565b5050505050505050565b5f60605f196001600160a01b03861601612489576124716001600160a01b03851684612d92565b505060408051602081019091525f81526001906124be565b82156124be576124998686612e24565b156124b3576124aa86868686612e5c565b915091506124be565b6124aa85858561300e565b94509492505050565b81355f908152600360209081526040909120908301356124ea60058301826125b8565b819061250c576040516306bfa19360e31b8152600401610a9491815260200190565b505f8181526003830160205260409020849061252882826150e2565b50505f81815260048301602052604090206125438482615182565b5050505050565b60015f5160206154b75f395f51905f5255565b3361258f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610ee85760405163118cdaa760e01b8152336004820152602401610a94565b5f61221183836131b5565b5f8281526003602052604090206125dd6005820183613201565b82906125ff5760405163052b63c160e11b8152600401610a9491815260200190565b505f828152600482016020526040812061261891614109565b5f828152600380830160205260408220828155600181018390556002810180546001600160a01b0319908116909155918101805490921690915560048101829055906125436005830182614109565b61266f61320c565b5f5160206154975f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610fc7565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061274657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661273a5f5160206154775f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610ee85760405163703e46dd60e11b815260040160405180910390fd5b610a3561255d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156127c6575060408051601f3d908101601f191682019092526127c391810190614f5c565b60015b6127ee57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a94565b5f5160206154775f395f51905f52811461281e57604051632a87526960e21b815260048101829052602401610a94565b612828838361323b565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ee85760405163703e46dd60e11b815260040160405180910390fd5b60605f61221183613290565b6066545f9081906001600160a01b03166128a057505f905080612991565b60665460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa1580156128f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061291a9190614f9a565b909450925090508086108015906129315750828510155b801561293d5750818410155b8184848989899091929394956129895760405163769296ad60e01b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a94565b505050505050505b9550959350505050565b6129b18a8a8a896129ac898b614f87565b6132e9565b6129c38a8a8a8a8a8a8a8a8a8a61348f565b5f8a81526003602052604090206001908101805490910190555b50505050505050505050565b6129f1613587565b6129fa336135d0565b612a03846135e1565b612a0c82613647565b612a14613785565b612a1c61378d565b612a2461379d565b6001600160a01b03811615612a4f57606680546001600160a01b0319166001600160a01b0383161790555b6001600160a01b038316612a94576040516211211f60e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a94565b5050606780546001600160a01b0319166001600160a01b03929092169190911790555043606855565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b83612b388688614f87565b612b429190614f87565b8360400151101589879091612b7a576040516224b63960e71b81526001600160a01b0390921660048301526024820152604401610a94565b5050612b85836137ad565b6129dd8a8a8a8a8a8a8a60018a8a61299b565b612ba06121c4565b5f5160206154975f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336126a8565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206154575f395f51905f5291612c1e90614f06565b80601f0160208091040260200160405190810160405280929190818152602001828054612c4a90614f06565b8015612c955780601f10612c6c57610100808354040283529160200191612c95565b820191905f5260205f20905b815481529060010190602001808311612c7857829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206154575f395f51905f5291612c1e90614f06565b5f610a24825490565b5f61221183836138aa565b5f8181526001830160205260408120541515612211565b5f612211836001600160a01b0384166131b5565b5f612211836001600160a01b0384166138d0565b5f611b066139b3565b5f610a24612d46612d31565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612d7688888888613a26565b925092509250612d868282613aee565b50909695505050505050565b80471015612dbc5760405163cf47918160e01b815247600482015260248101829052604401610a94565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f8114612e06576040519150601f19603f3d011682016040523d82523d5f602084013e612e0b565b606091505b509150915081612e1e57612e1e81613ba6565b50505050565b5f9182526003602090815260408084206001600160a01b0393909316845260099092019052902060010154600160a01b900460ff1690565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015612ecb575060408051601f3d908101601f19168201909252612ec89181019061523c565b60015b612f9d57612ed7615257565b806308c379a003612f005750612eeb61526f565b80612ef65750612f5a565b5f925090506124be565b634e487b7103612f5a57612f126152f1565b90612f1d5750612f5a565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b6040516020818303038152906040529150506124be565b3d808015612f83576040519150601f19603f3d011682016040523d82523d5f602084013e612f88565b606091505b505f925080604051602001612f43919061530e565b8015612fc8576001925060405180602001604052805f8152509150612fc3878786613bcf565b613004565b5f92506040518060400160405280601f81526020017f5374616e646172644272696467653a207472616e73666572206661696c65640081525091505b5094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af192505050801561307d575060408051601f3d908101601f1916820190925261307a9181019061523c565b60015b61314f57613089615257565b806308c379a0036130b2575061309d61526f565b806130a8575061310c565b5f925090506131ad565b634e487b710361310c576130c46152f1565b906130cf575061310c565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b6040516020818303038152906040529150506131ad565b3d808015613135576040519150601f19603f3d011682016040523d82523d5f602084013e61313a565b606091505b505f9250806040516020016130f5919061530e565b801561316f576001925060405180602001604052805f81525091506131ab565b5f92506040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b5f8181526001830160205260408120546131fa57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a24565b505f610a24565b5f61221183836138d0565b5f5160206154975f395f51905f525460ff16610ee857604051638dfc202b60e01b815260040160405180910390fd5b61324482613c71565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613288576128288282613cd4565b610ddd613d46565b6060815f018054806020026020016040519081016040528092919081815260200182805480156132dd57602002820191905f5260205f20905b8154815260200190600101908083116132c9575b50505050509050919050565b5f196001600160a01b0385160161335c576133048183614f87565b34146133108284614f87565b3490916133395760405163341382c560e11b815260048101929092526024820152604401610a94565b5050801561335757606754613357906001600160a01b031682612d92565b612543565b5f3480156133865760405163341382c560e11b815260048101929092526024820152604401610a94565b506133aa905083306133988486614f87565b6001600160a01b038816929190613d65565b80156133ca576067546133ca906001600160a01b03868116911683613dcc565b6133d48585612e24565b156133e457613357858584613dfd565b604051632770a7eb60e21b8152306004820152602481018390526001600160a01b03851690639dc29fac906044016020604051808303815f875af115801561342e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613452919061523c565b8484849091926124405760405163bc8ecf7b60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610a94565b5f5f61349a8c611b0b565b91506134cf8c8c5f9182526003602090815260408084206001600160a01b039384168552600901909152909120600101541690565b9050896001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d78e858e8e428d8d8d60405161351a989796959493929190615333565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a604051613571929190918252602082015260400190565b60405180910390a4505050505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610ee857604051631afcd79f60e31b815260040160405180910390fd5b6135d8613587565b610a3581613e3e565b6135e9613587565b613631604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613e46565b6035805460ff191660ff92909216919091179055565b6001600160a01b0381166136585750565b5f816001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015613694573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136bb919081019061537d565b604080513060208201520160408051601f19818403018152908290526136e492916020016153fc565b60408051601f1981840301815282825246602084015292505f91613721918391016040516020818303038152906040528051906020012084613e58565b90506001600160a01b0381166137625760405162461bcd60e51b8152600401610a94906020808252600490820152637a65726f60e01b604082015260600190565b5f80546001600160a01b0319166001600160a01b03929092169190911790555050565b610ee8613587565b613795613587565b610ee8613eec565b6137a5613587565b610ee8613f0c565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af180613859576040513d5f823e3d81fd5b50505f513d9150811561387057806001141561387e565b84516001600160a01b03163b155b1561254357845160405163733eb19560e01b81526001600160a01b039091166004820152602401610a94565b5f825f0182815481106138bf576138bf614da9565b905f5260205f200154905092915050565b5f81815260018301602052604081205480156139aa575f6138f2600183615410565b85549091505f9061390590600190615410565b9050808214613964575f865f01828154811061392357613923614da9565b905f5260205f200154905080875f01848154811061394357613943614da9565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061397557613975615423565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a24565b5f915050610a24565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6139dd613f14565b6139e5613f7c565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613a5f57505f91506003905082613ae4565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613ab0573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613adb57505f925060019150829050613ae4565b92505f91508190505b9450945094915050565b5f826003811115613b0157613b01615437565b03613b0a575050565b6001826003811115613b1e57613b1e615437565b03613b3c5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613b5057613b50615437565b03613b715760405163fce698f760e01b815260048101829052602401610a94565b6003826003811115613b8557613b85615437565b03610ddd576040516335e2f38360e21b815260048101829052602401610a94565b805115613bb65780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f8381526003602090815260408083206001600160a01b0386168452600981019092528220600201549091908190613c079085613fbe565b84549193509150858584613c4757604051630106202f60e71b815260048101939093526001600160a01b0390911660248301526044820152606401610a94565b5050506001600160a01b039094165f90815260099092016020525060409020600201919091555050565b806001600160a01b03163b5f03613ca657604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a94565b5f5160206154775f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613cf0919061544b565b5f60405180830381855af49150503d805f8114613d28576040519150601f19603f3d011682016040523d82523d5f602084013e613d2d565b606091505b5091509150613d3d858383613fe2565b95945050505050565b3415610ee85760405163b398979f60e01b815260040160405180910390fd5b6040516001600160a01b038481166024830152838116604483015260648201839052612e1e9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061403e565b6040516001600160a01b0383811660248301526044820183905261282891859182169063a9059cbb90606401613d9a565b5f8381526003602090815260408083206001600160a01b038616845260090190915281206002018054839290613e34908490614f87565b9091555050505050565b611f82613587565b613e4e613587565b610ddd82826140aa565b5f83471015613e835760405163cf47918160e01b815247600482015260248101859052604401610a94565b81515f03613ea457604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615613ec5576040513d5f823e3d81fd5b6001600160a01b0381166122115760405163b06ebf3d60e01b815260040160405180910390fd5b613ef4613587565b5f5160206154975f395f51905f52805460ff19169055565b61254a613587565b5f5f5160206154575f395f51905f5281613f2c612be0565b805190915015613f4457805160209091012092915050565b81548015613f53579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206154575f395f51905f5281613f94612ca0565b805190915015613fac57805160209091012092915050565b60018201548015613f53579392505050565b5f5f83831115613fd257505f905080613fdb565b50600190508183035b9250929050565b606082613ff757613ff282613ba6565b612211565b815115801561400e57506001600160a01b0384163b155b1561403757604051639996b31560e01b81526001600160a01b0385166004820152602401610a94565b5080612211565b5f5f60205f8451602086015f885af18061405d576040513d5f823e3d81fd5b50505f513d91508115614074578060011415614081565b6001600160a01b0384163b155b15612e1e57604051635274afe760e01b81526001600160a01b0385166004820152602401610a94565b6140b2613587565b5f5160206154575f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026140eb8482615182565b50600381016140fa8382615182565b505f8082556001909101555050565b50805461411590614f06565b5f825580601f10614124575050565b601f0160209004905f5260205f2090810190610a3591905b8082111561414f575f815560010161413c565b5090565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b038211171561419a5761419a614167565b60405250565b601f8201601f191681016001600160401b03811182821017156141c5576141c5614167565b6040525050565b5f6001600160401b038211156141e4576141e4614167565b5060051b60200190565b5f5f604083850312156141ff575f5ffd5b8235915060208301356001600160401b0381111561421b575f5ffd5b8301601f8101851361422b575f5ffd5b8035614236816141cc565b60405161424382826141a0565b80915082815260208101915060208360051b850101925087831115614266575f5ffd5b6020840193505b8284101561428857833582526020938401939091019061426d565b809450505050509250929050565b6001600160a01b0381168114610a35575f5ffd5b5f602082840312156142ba575f5ffd5b813561221181614296565b803560ff811681146142d5575f5ffd5b919050565b5f82601f8301126142e9575f5ffd5b81356142f4816141cc565b60405161430182826141a0565b80915082815260208101915060208360051b860101925085831115614324575f5ffd5b602085015b838110156143485761433a816142c5565b835260209283019201614329565b5095945050505050565b5f82601f830112614361575f5ffd5b813561436c816141cc565b60405161437982826141a0565b80915082815260208101915060208360051b86010192508583111561439c575f5ffd5b602085015b838110156143485780358352602092830192016143a1565b5f5f5f5f608085870312156143cc575f5ffd5b84356001600160401b038111156143e1575f5ffd5b850160c081880312156143f2575f5ffd5b935060208501356001600160401b0381111561440c575f5ffd5b614418878288016142da565b93505060408501356001600160401b03811115614433575f5ffd5b61443f87828801614352565b92505060608501356001600160401b0381111561445a575f5ffd5b61446687828801614352565b91505092959194509250565b5f60208284031215614482575f5ffd5b5035919050565b5f60208284031215614499575f5ffd5b81356001600160401b038111156144ae575f5ffd5b8201601f810184136144be575f5ffd5b80356144c9816141cc565b6040516144d682826141a0565b80915082815260208101915060208360051b8501019250868311156144f9575f5ffd5b6020840193505b8284101561452457833561451381614296565b825260209384019390910190614500565b9695505050505050565b5f5f6040838503121561453f575f5ffd5b50508035926020909101359150565b5f6001600160401b0382111561456657614566614167565b50601f01601f191660200190565b5f61457e8361454e565b60405161458b82826141a0565b80925084815285858501111561459f575f5ffd5b848460208301375f6020868301015250509392505050565b5f5f604083850312156145c8575f5ffd5b82356145d381614296565b915060208301356001600160401b038111156145ed575f5ffd5b8301601f810185136145fd575f5ffd5b61460c85823560208401614574565b9150509250929050565b80516001600160a01b03908116835260208083015190911690830152604080820151151590830152606080820151151590830152608090810151910152565b602080825282518282018190525f918401906040840190835b8181101561469757614681838551614616565b6020939093019260a0929092019160010161466e565b509095945050505050565b5f5f83601f8401126146b2575f5ffd5b5081356001600160401b038111156146c8575f5ffd5b602083019150836020828501011115613fdb575f5ffd5b5f5f5f5f5f5f5f5f60e0898b0312156146f6575f5ffd5b88359750602089013561470881614296565b9650604089013561471881614296565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115614747575f5ffd5b6147538b828c016146a2565b999c989b5096995094979396929594505050565b5f5f5f5f6080858703121561477a575f5ffd5b614783856142c5565b9350602085013561479381614296565b925060408501356147a381614296565b915060608501356147b381614296565b939692955090935050565b5f5f5f5f608085870312156147d1575f5ffd5b8435935060208501356147e381614296565b925060408501356001600160401b038111156147fd575f5ffd5b8501601f8101871361480d575f5ffd5b61481c87823560208401614574565b92505061482b606086016142c5565b905092959194509250565b5f5f5f5f5f5f5f5f5f5f8a8c036101e0811215614851575f5ffd5b8b359a5060208c013561486381614296565b995060408c013561487381614296565b985060608c013561488381614296565b975060808c0135965060a08c0135955060c08c0135945060e060df19820112156148ab575f5ffd5b506040516148b88161417b565b60e08c01356148c681614296565b81526101008c01356148d781614296565b60208201526101208c013560408201526101408c013560608201526148ff6101608d016142c5565b60808201526101808c013560a08201526101a08c013560c082015292506101c08b01356001600160401b03811115614935575f5ffd5b6149418d828e016146a2565b915080935050809150509295989b9194979a5092959850565b5f5f6040838503121561496b575f5ffd5b82359150602083013561497d81614296565b809150509250929050565b60a08101610a248284614616565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6122116020830184614996565b5f8151808452602084019350602083015f5b82811015614a065781518652602095860195909101906001016149e8565b5093949350505050565b60ff60f81b8816815260e060208201525f614a2e60e0830189614996565b8281036040840152614a408189614996565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050614a7181856149d6565b9a9950505050505050505050565b5f82601f830112614a8e575f5ffd5b8135614a99816141cc565b604051614aa682826141a0565b80915082815260208101915060208360051b860101925085831115614ac9575f5ffd5b602085015b838110156143485780356001600160401b03811115614aeb575f5ffd5b614afa886020838a01016142da565b84525060209283019201614ace565b5f82601f830112614b18575f5ffd5b8135614b23816141cc565b604051614b3082826141a0565b80915082815260208101915060208360051b860101925085831115614b53575f5ffd5b602085015b838110156143485780356001600160401b03811115614b75575f5ffd5b614b84886020838a0101614352565b84525060209283019201614b58565b5f5f5f5f5f60808688031215614ba7575f5ffd5b85356001600160401b03811115614bbc575f5ffd5b8601601f81018813614bcc575f5ffd5b80356001600160401b03811115614be1575f5ffd5b8860208260051b8401011115614bf5575f5ffd5b6020918201965094508601356001600160401b03811115614c14575f5ffd5b614c2088828901614a7f565b93505060408601356001600160401b03811115614c3b575f5ffd5b614c4788828901614b09565b92505060608601356001600160401b03811115614c62575f5ffd5b614c6e88828901614b09565b9150509295509295909350565b60208152815160208201526020820151604082015260018060a01b03604083015116606082015260018060a01b036060830151166080820152608082015160a08201525f60a083015160c0808401526114f160e0840182614996565b5f5f5f60608486031215614ce9575f5ffd5b833592506020840135614cfb81614296565b929592945050506040919091013590565b5f60208284031215614d1c575f5ffd5b612211826142c5565b602081525f61221160208301846149d6565b8015158114610a35575f5ffd5b5f5f5f5f60808587031215614d57575f5ffd5b84359350602085013561479381614d37565b602080825282518282018190525f918401906040840190835b818110156146975783516001600160a01b0316835260209384019390920191600101614d82565b634e487b7160e01b5f52603260045260245ffd5b5f5f8335601e19843603018112614dd2575f5ffd5b8301803591506001600160401b03821115614deb575f5ffd5b602001915036819003821315613fdb575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90614e679083018486614dff565b9998505050505050505050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f612211600d830184614e74565b848152608060208201525f614ec46080830186614996565b8281036040840152614ed68186614996565b91505060ff8316606083015295945050505050565b5f60208284031215614efb575f5ffd5b815161221181614296565b600181811c90821680614f1a57607f821691505b602082108103614f3857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f823560be19833603018112614f52575f5ffd5b9190910192915050565b5f60208284031215614f6c575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610a2457610a24614f73565b5f5f5f60608486031215614fac575f5ffd5b5050815160208301516040909301519094929350919050565b80546001600160a01b0319166001600160a01b0392909216919091179055565b601f82111561282857805f5260205f20601f840160051c8101602085101561500a5750805b601f840160051c820191505b81811015612543575f8155600101615016565b6001600160401b0383111561504057615040614167565b6150548361504e8354614f06565b83614fe5565b5f601f841160018114615085575f851561506e5750838201355b5f19600387901b1c1916600186901b178355612543565b5f83815260208120601f198716915b828110156150b45786850135825560209485019460019092019101615094565b50868210156150d0575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b813581556020820135600182015560408201356150fe81614296565b61510b8160028401614fc5565b50606082013561511a81614296565b6151278160038401614fc5565b506080820135600482015560a082013536839003601e19018112615149575f5ffd5b820180356001600160401b03811115615160575f5ffd5b602082019150803603821315615174575f5ffd5b612e1e818360058601615029565b81516001600160401b0381111561519b5761519b614167565b6151af816151a98454614f06565b84614fe5565b6020601f8211600181146151e1575f83156151ca5750848201515b5f19600385901b1c1916600184901b178455612543565b5f84815260208120601f198516915b8281101561521057878501518255602094850194600190920191016151f0565b508482101561522d57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6020828403121561524c575f5ffd5b815161221181614d37565b5f60033d1115610f285760045f5f3e505f5160e01c90565b5f60443d101561527c5790565b6040513d600319016004823e80513d60248201116001600160401b03821117156152a557505090565b80820180516001600160401b038111156152c0575050505090565b3d84016003190182820160200111156152da575050505090565b6152e9602082850101856141a0565b509392505050565b5f5f60233d111561530a57602060045f3e50505f516001905b9091565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f6122116011830184614e74565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f90614a719083018486614dff565b5f6020828403121561538d575f5ffd5b81516001600160401b038111156153a2575f5ffd5b8201601f810184136153b2575f5ffd5b80516153bd8161454e565b6040516153ca82826141a0565b8281528660208486010111156153de575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b5f6114f161540a8386614e74565b84614e74565b81810381811115610a2457610a24614f73565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f6122118284614e7456fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122049bbb8f50b731c7ec16865fad10a349f7c3f48d75a1a02b8510f9668de689b4064736f6c634300081c0033",
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

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _BranchBridge.Contract.AllChainIDs(&_BranchBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BranchBridge *BranchBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BranchBridge.Contract.AllChainIDs(&_BranchBridge.CallOpts)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.AllTokenPairs(&_BranchBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_BranchBridge *BranchBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
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
func (_BranchBridge *BranchBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BranchBridge.Contract.GetTokenPair(&_BranchBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_BranchBridge *BranchBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
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
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCaller) RevertedArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	var out []interface{}
	err := _BranchBridge.contract.Call(opts, &out, "revertedArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryFinalizeArguments)).(*IBridgeRegistryFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
	return _BranchBridge.Contract.RevertedArguments(&_BranchBridge.CallOpts, remoteChainID, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,uint256,address,address,uint256,bytes))
func (_BranchBridge *BranchBridgeCallerSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryFinalizeArguments, error) {
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

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridge(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BranchBridge.Contract.FinalizeBridgeBatch(&_BranchBridge.TransactOpts, args, v, r, s)
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
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridge", remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 remoteChainID, uint256 index) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridge(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridge(&_BranchBridge.TransactOpts, remoteChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, remoteChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 remoteChainID, uint256[] indexes) returns(bool)
func (_BranchBridge *BranchBridgeTransactorSession) RetryFinalizeBridgeBatch(remoteChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BranchBridge.Contract.RetryFinalizeBridgeBatch(&_BranchBridge.TransactOpts, remoteChainID, indexes)
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
func (_BranchBridge *BranchBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (*BranchBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BranchBridgeBridgeFinalizedIterator{contract: _BranchBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
func (_BranchBridge *BranchBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BranchBridgeBridgeFinalized, remoteChainID []*big.Int, index []*big.Int, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "BridgeFinalized", remoteChainIDRule, indexRule, toRule)
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

// ParseBridgeFinalized is a log parse operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed remoteChainID, uint256 indexed index, address token, address indexed to, uint256 value, uint256 time)
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

// BranchBridgeFeeStationSetIterator is returned from FilterFeeStationSet and is used to iterate over the raw logs and unpacked data for FeeStationSet events raised by the BranchBridge contract.
type BranchBridgeFeeStationSetIterator struct {
	Event *BranchBridgeFeeStationSet // Event containing the contract specifics and raw log

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
func (it *BranchBridgeFeeStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeFeeStationSet)
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
		it.Event = new(BranchBridgeFeeStationSet)
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
func (it *BranchBridgeFeeStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeFeeStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeFeeStationSet represents a FeeStationSet event raised by the BranchBridge contract.
type BranchBridgeFeeStationSet struct {
	FeeStation common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeStationSet is a free log retrieval operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address feeStation)
func (_BranchBridge *BranchBridgeFilterer) FilterFeeStationSet(opts *bind.FilterOpts) (*BranchBridgeFeeStationSetIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "FeeStationSet")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeFeeStationSetIterator{contract: _BranchBridge.contract, event: "FeeStationSet", logs: logs, sub: sub}, nil
}

// WatchFeeStationSet is a free log subscription operation binding the contract event 0x68e453b7f82baf4a13086202d4cfa2c41dc5c07b2f441c1f6da0d93bee2ea444.
//
// Solidity: event FeeStationSet(address feeStation)
func (_BranchBridge *BranchBridgeFilterer) WatchFeeStationSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeFeeStationSet) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "FeeStationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeFeeStationSet)
				if err := _BranchBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseFeeStationSet(log types.Log) (*BranchBridgeFeeStationSet, error) {
	event := new(BranchBridgeFeeStationSet)
	if err := _BranchBridge.contract.UnpackLog(event, "FeeStationSet", log); err != nil {
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

// BranchBridgeRewardWalletSetIterator is returned from FilterRewardWalletSet and is used to iterate over the raw logs and unpacked data for RewardWalletSet events raised by the BranchBridge contract.
type BranchBridgeRewardWalletSetIterator struct {
	Event *BranchBridgeRewardWalletSet // Event containing the contract specifics and raw log

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
func (it *BranchBridgeRewardWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BranchBridgeRewardWalletSet)
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
		it.Event = new(BranchBridgeRewardWalletSet)
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
func (it *BranchBridgeRewardWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BranchBridgeRewardWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BranchBridgeRewardWalletSet represents a RewardWalletSet event raised by the BranchBridge contract.
type BranchBridgeRewardWalletSet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardWalletSet is a free log retrieval operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address wallet)
func (_BranchBridge *BranchBridgeFilterer) FilterRewardWalletSet(opts *bind.FilterOpts) (*BranchBridgeRewardWalletSetIterator, error) {

	logs, sub, err := _BranchBridge.contract.FilterLogs(opts, "RewardWalletSet")
	if err != nil {
		return nil, err
	}
	return &BranchBridgeRewardWalletSetIterator{contract: _BranchBridge.contract, event: "RewardWalletSet", logs: logs, sub: sub}, nil
}

// WatchRewardWalletSet is a free log subscription operation binding the contract event 0x5fa7ec703bca1ae326c2fb6a283f918d39c66634fe7afeebbabf819628f13c9b.
//
// Solidity: event RewardWalletSet(address wallet)
func (_BranchBridge *BranchBridgeFilterer) WatchRewardWalletSet(opts *bind.WatchOpts, sink chan<- *BranchBridgeRewardWalletSet) (event.Subscription, error) {

	logs, sub, err := _BranchBridge.contract.WatchLogs(opts, "RewardWalletSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BranchBridgeRewardWalletSet)
				if err := _BranchBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
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
func (_BranchBridge *BranchBridgeFilterer) ParseRewardWalletSet(log types.Log) (*BranchBridgeRewardWalletSet, error) {
	event := new(BranchBridgeRewardWalletSet)
	if err := _BranchBridge.contract.UnpackLog(event, "RewardWalletSet", log); err != nil {
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
