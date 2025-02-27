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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIChainManager.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFeeStation\",\"outputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Factory\",\"outputs\":[{\"internalType\":\"contractCrossMintableERC20Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIChainManager.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"}],\"internalType\":\"structIChainManager.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_rewardWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossMintableERC20FactoryCode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIStandardBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"retryFinalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"retryFinalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedArguments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIChainManager.FinalizeArguments\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revertedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardWallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"setChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeFeeStation\",\"name\":\"_bridgeFeeStation\",\"type\":\"address\"}],\"name\":\"setFeeStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"rewardWallet_\",\"type\":\"address\"}],\"name\":\"setRewardWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"permit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"ChainSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenPairUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"EthereumBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"EthereumBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibAleadyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errChainLibCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errChainLibDuplicateIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errChainLibInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"errChainLibNotExistingIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainLibNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errChainManagerChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"errChainManagerChainNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errChainManagerInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"errChainManagerTokenFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errChainManagerTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"errStandardBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errStandardBridgeFailedPermit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedEx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualEx\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidPermitValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"errStandardBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516155146100f95f395f8181612086015281816120af01526121f301526155145ff3fe6080604052600436106102c2575f3560e01c8063846704f81161016f578063aed1d403116100d8578063edbbea3911610092578063f45096431161006d578063f45096431461093c578063f698da251461095b578063facd743b1461096f578063fb75b2c71461098e575f5ffd5b8063edbbea39146108dd578063f2fde38b146108fc578063f30589c31461091b575f5ffd5b8063aed1d40314610839578063b7f3358d1461084d578063cbae59581461086c578063d2ff130d1461088b578063d5717fc5146108aa578063d7c82f32146108c9575f5ffd5b80639300c926116101295780639300c92614610751578063952a95de1461077057806396ce07951461079c578063ad3cb1cc146107b0578063ae6893f8146107e0578063ae766389146107ff575f5ffd5b8063846704f81461067e57806384b0196e1461069d57806384d58d42146106c45780638da5cb5b146106e35780638f517c171461071f57806391cf6d3e1461073d575f5ffd5b806354db01261161022b5780637101fcd3116101e55780637f260a75116101c05780637f260a7514610564578063814914b5146105775780638415a3851461063e5780638456cb591461066a575f5ffd5b80637101fcd31461051e578063715018a61461053d5780637c7505dc14610551575f5ffd5b806354db01261461045f5780635958621e1461047e5780635b605f5c1461049d5780635c975abb146104c95780635fd262de146104ec578063670e6268146104ff575f5ffd5b80633f4ba83a1161027c5780633f4ba83a1461039f57806340a141ff146103b357806342cde4e8146103d257806347666cb1146103f35780634f1ef2861461042a57806352d1902d1461043d575f5ffd5b8063030372c3146102dc5780631327d3d81461031057806315eb39ef1461032f5780631b44fd15146103425780631d40f0d8146103615780633960e78714610380575f5ffd5b366102d857345f036102d6576102d661423b565b005b5f5ffd5b3480156102e7575f5ffd5b506102fb6102f63660046142d6565b6109ab565b60405190151581526020015b60405180910390f35b34801561031b575f5ffd5b506102d661032a3660046143a2565b6109ef565b6102fb61033d3660046144e9565b6109fd565b34801561034d575f5ffd5b506102d661035c3660046145e0565b610c76565b34801561036c575f5ffd5b506102d661037b3660046145f7565b610caf565b34801561038b575f5ffd5b506102fb61039a36600461469c565b610ce5565b3480156103aa575f5ffd5b506102d6610dcf565b3480156103be575f5ffd5b506102d66103cd3660046143a2565b610de1565b3480156103dd575f5ffd5b5060355460405160ff9091168152602001610307565b3480156103fe575f5ffd5b50606654610412906001600160a01b031681565b6040516001600160a01b039091168152602001610307565b6102d6610438366004614725565b610deb565b348015610448575f5ffd5b50610451610e06565b604051908152602001610307565b34801561046a575f5ffd5b506102d66104793660046143a2565b610e22565b348015610489575f5ffd5b506102d66104983660046143a2565b610e96565b3480156104a8575f5ffd5b506104bc6104b73660046145e0565b610f06565b60405161030791906147c3565b3480156104d4575f5ffd5b505f51602061549f5f395f51905f525460ff166102fb565b6102fb6104fa366004614810565b61105f565b34801561050a575f5ffd5b50610412610519366004614898565b61116e565b348015610529575f5ffd5b506102d66105383660046145f7565b61127f565b348015610548575f5ffd5b506102d6611295565b6102fb61055f366004614a24565b6112a6565b6102fb610572366004614b17565b611408565b348015610582575f5ffd5b50610631610591366004614c3b565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9182526003602090815260408084206001600160a01b039384168552600901825292839020835160a08101855281548416815260018201549384169281019290925260ff600160a01b84048116151594830194909452600160a81b9092049092161515606083015260020154608082015290565b6040516103079190614c69565b348015610649575f5ffd5b5061065d61065836600461469c565b6114e7565b6040516103079190614ca5565b348015610675575f5ffd5b506102d6611508565b348015610689575f5ffd5b506102d6610698366004614cb7565b611518565b3480156106a8575f5ffd5b506106b161169d565b6040516103079796959493929190614d2a565b3480156106cf575f5ffd5b506102d66106de366004614c3b565b611746565b3480156106ee575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610412565b34801561072a575f5ffd5b505f54610412906001600160a01b031681565b348015610748575f5ffd5b50606854610451565b34801561075c575f5ffd5b506102d661076b3660046145f7565b61179e565b34801561077b575f5ffd5b5061078f61078a36600461469c565b6117d5565b6040516103079190614dc0565b3480156107a7575f5ffd5b506104516117f4565b3480156107bb575f5ffd5b5061065d604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156107eb575f5ffd5b506104516107fa3660046145e0565b611864565b34801561080a575f5ffd5b5061081e610819366004614e12565b61187a565b60408051938452602084019290925290820152606001610307565b348015610844575f5ffd5b50610451611909565b348015610858575f5ffd5b506102d6610867366004614e47565b611914565b348015610877575f5ffd5b506104126108863660046145e0565b611965565b348015610896575f5ffd5b506102d66108a5366004614c3b565b611971565b3480156108b5575f5ffd5b506104516108c43660046145e0565b6119c9565b3480156108d4575f5ffd5b506102d66119df565b3480156108e8575f5ffd5b506102d66108f7366004614e6d565b6119f9565b348015610907575f5ffd5b506102d66109163660046143a2565b611a9e565b348015610926575f5ffd5b5061092f611ad8565b6040516103079190614ebd565b348015610947575f5ffd5b506102d6610956366004614c3b565b611ae4565b348015610966575f5ffd5b50610451611b6a565b34801561097a575f5ffd5b506102fb6109893660046143a2565b611b73565b348015610999575f5ffd5b506067546001600160a01b0316610412565b5f805b82518110156109e3576109da848483815181106109cd576109cd614efd565b6020026020010151610ce5565b506001016109ae565b50600190505b92915050565b6109fa816001611b7f565b50565b5f610a06611c48565b5f8b81526003602052604090208b908a90610a249060070182611c78565b8190610a545760405163c4bfa74960e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff1615610ab35760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b50610abc611c99565b5f348015610ae65760405163341382c560e11b815260048101929092526024820152604401610a4b565b5050610af18d6119c9565b8c14610afc8e6119c9565b8d9091610b25576040516329846bd360e21b815260048101929092526024820152604401610a4b565b5050610b857f8ee949cbdad2722efdf3806ebb3e900a8822e1e18aea7c94f0ce625ff090b6d18d8d8d8d8d8d604051602001610b679796959493929190614f39565b60405160208183030381529060405280519060200120878787611cd0565b5f8d81526003602052604081206002018054600101905580610ba98f8e8e8e611ecb565b915091508115610c0f578b6001600160a01b03168d6001600160a01b03168f7fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a3738e42604051610c02929190918252602082015260400190565b60405180910390a4610c4a565b610c1f8f8f8f8f8f8f8f88611f0e565b6040518e907ffb8da6f1288864ddb7fd0d3c022ac5e612bd23aeba60eb826188c5beb0e58e7f905f90a25b600194505050610c6660015f5160206154bf5f395f51905f5255565b50509a9950505050505050505050565b610c7e611fa5565b610c89600182612000565b8190610cab5760405163021aa6c760e21b8152600401610a4b91815260200190565b5050565b5f5b8151811015610cab57610cdd828281518110610ccf57610ccf614efd565b60200260200101515f611b7f565b600101610cb1565b5f610cee611c48565b610cf6611c99565b5f610d0184846117d5565b90505f5f610d1d86846020015185604001518660600151611ecb565b91509150818190610d415760405162461bcd60e51b8152600401610a4b9190614ca5565b50610d4c868661200b565b82604001516001600160a01b031683602001516001600160a01b0316845f01517fc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373866060015142604051610daa929190918252602082015260400190565b60405180910390a4600193505050506109e960015f5160206154bf5f395f51905f5255565b610dd7611fa5565b610ddf612022565b565b6109fa815f611b7f565b610df361207b565b610dfc8261211f565b610cab8282612127565b5f610e0f6121e8565b505f51602061547f5f395f51905f525b90565b610e2a611fa5565b6001600160a01b038116610e74576040516211211f60e21b81526020600482015260116024820152702fb13934b233b2a332b2a9ba30ba34b7b760791b6044820152606401610a4b565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b610e9e611fa5565b6001600160a01b038116610ee4576040516211211f60e21b815260206004820152600d60248201526c72657761726457616c6c65745f60981b6044820152606401610a4b565b606780546001600160a01b0319166001600160a01b0392909216919091179055565b5f818152600360205260408120606091610f2260078301612231565b90505f81516001600160401b03811115610f3e57610f3e61424f565b604051908082528060200260200182016040528015610f9557816020015b6040805160a0810182525f808252602080830182905292820181905260608201819052608082015282525f19909201910181610f5c5790505b5090505f5b825181101561105657836009015f848381518110610fba57610fba614efd565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160a08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b909304909116151560608201526002909101546080820152825183908390811061104357611043614efd565b6020908102919091010152600101610f9a565b50949350505050565b5f611068611c48565b5f898152600360205260409020899089906110869060070182611c78565b81906110b15760405163c4bfa74960e01b81526001600160a01b039091166004820152602401610a4b565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156111105760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b50611119611c99565b5f5f6111288d8d8c8c8c61223d565b915091506111448d8d6111383390565b8e8e87875f8f8f612356565b60019450505061116060015f5160206154bf5f395f51905f5255565b505098975050505050505050565b5f611177611fa5565b5f546001600160a01b031661119f576040516380cf12ed60e01b815260040160405180910390fd5b5f546040516bffffffffffffffffffffffff19606087901b1660208201526001600160a01b0390911690634804a0419060340160405160208183030381529060405280519060200120856040516020016111f99190614f9d565b60405160208183030381529060405286866040518563ffffffff1660e01b81526004016112299493929190614fbe565b6020604051808303815f875af1158015611245573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112699190614ffd565b9050611277855f83876119f9565b949350505050565b61128c61037b6033612231565b6109fa8161179e565b61129d611fa5565b610ddf5f6123a4565b5f805b858110156113fa576113f1888888848181106112c7576112c7614efd565b90506020028101906112d99190615018565b358989858181106112ec576112ec614efd565b90506020028101906112fe9190615018565b61130f9060408101906020016143a2565b8a8a8681811061132157611321614efd565b90506020028101906113339190615018565b6113449060608101906040016143a2565b8b8b8781811061135657611356614efd565b90506020028101906113689190615018565b606001358c8c8881811061137e5761137e614efd565b90506020028101906113909190615018565b61139e906080810190615036565b8c89815181106113b0576113b0614efd565b60200260200101518c8a815181106113ca576113ca614efd565b60200260200101518c8b815181106113e4576113e4614efd565b60200260200101516109fd565b506001016112a9565b506001979650505050505050565b5f611411611c48565b5f8b81526003602052604090208b908b9061142f9060070182611c78565b819061145a5760405163c4bfa74960e01b81526001600160a01b039091166004820152602401610a4b565b505f8281526003602090815260408083206001600160a01b03851684526009019091529020600101548190600160a81b900460ff16156114b95760405163026bfe9360e11b81526001600160a01b039091166004820152602401610a4b565b506114c2611c99565b5f5f6114d18f8f8d8d8d61223d565b91509150610c4a8f8f8f8f8f87878f8f8f612414565b5f828152600360205260409020606090611501908361247f565b9392505050565b611510611fa5565b610ddf612521565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561155c5750825b90505f826001600160401b031660011480156115775750303b155b905081158015611585575080155b156115a35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156115cd57845460ff60401b1916600160401b1785555b6115d989898989612569565b6001600160a01b038a166116185760405163634fe78160e11b815260206004820152600560248201526463726f737360d81b6044820152606401610a4b565b609880546001600160a01b0319166001600160a01b038c1617905561163c8b610c76565b61164a8b60018c60016119f9565b831561169057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b5f60608082808083815f51602061545f5f395f51905f5280549091501580156116c857506001810154155b61170c5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610a4b565b61171461263d565b61171c6126fd565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b61174e611fa5565b5f828152600360205260409020611765908261273b565b6040516001600160a01b0382169083907fac800026e47823d250f3265e869fac49cc159ebb8125df4559d01bfeb12f30e9905f90a35050565b5f5b8151811015610cab576117cd8282815181106117be576117be614efd565b60200260200101516001611b7f565b6001016117a0565b6117dd6141b4565b5f83815260036020526040902061150190836127f0565b606654604080516396ce079560e01b815290515f926001600160a01b0316916396ce07959160048083019260209291908290030181865afa15801561183b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061185f9190615078565b905090565b5f8181526003602052604081206109e9906128e7565b60665460405163ae76638960e01b8152600481018590526001600160a01b038481166024830152604482018490525f92839283929091169063ae76638990606401606060405180830381865afa1580156118d6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118fa919061508f565b91989097509095509350505050565b5f61185f60336128f9565b61191c611fa5565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff906020015b60405180910390a150565b5f6109e9603383612902565b611979611fa5565b5f828152600360205260409020611990908261290d565b6040516001600160a01b0382169083907ff98d95d31b49b957e140b6af179355984e4bccd450d9ff64c2c75ae111f9a1a3905f90a35050565b5f8181526003602052604081206109e9906129c9565b6119e7611fa5565b606680546001600160a01b0319169055565b611a01611fa5565b611a0c6001856129db565b8490611a2e57604051636292216560e11b8152600401610a4b91815260200190565b505f848152600360205260409020611a48908484846129f2565b806001600160a01b0316826001600160a01b0316857f78f7ef4e1e6dfd6d99bb7b0c064f736f5d3799f47af7b23242090631f05adb2c86604051611a90911515815260200190565b60405180910390a450505050565b611aa6611fa5565b6001600160a01b038116611acf57604051631e4fbdf760e01b81525f6004820152602401610a4b565b6109fa816123a4565b606061185f6033612231565b611aec611fa5565b611af76001836129db565b8290611b1957604051636292216560e11b8152600401610a4b91815260200190565b505f828152600360205260409020611b319082612b1e565b6040516001600160a01b0382169083907fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d905f90a35050565b5f61185f612bd9565b5f6109e9603383611c78565b611b87611fa5565b8015611bc957611b98603383612be2565b8290611bc3576040516329a04e7760e21b81526001600160a01b039091166004820152602401610a4b565b50611c01565b611bd4603383612bf6565b8290611bff5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610a4b565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f51602061549f5f395f51905f525460ff1615610ddf5760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381165f9081526001830160205260408120541515611501565b5f5160206154bf5f395f51905f52805460011901611cca57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b8251825181148015611ce25750815181145b835183518392611d16576040516337a9ac2560e01b8152600481019390935260248301919091526044820152606401610a4b565b505060355482915060ff16811015611d4457604051632fcba65760e11b8152600401610a4b91815260200190565b505f80826001600160401b03811115611d5f57611d5f61424f565b604051908082528060200260200182016040528015611d88578160200160208202803683370190505b5090505f5b83811015611e95575f611df8888381518110611dab57611dab614efd565b6020026020010151888481518110611dc557611dc5614efd565b6020026020010151888581518110611ddf57611ddf614efd565b6020026020010151611df08d612c0a565b929190612c36565b9050611e0381611b73565b8190611e2e5760405163845a09e760e01b81526001600160a01b039091166004820152602401610a4b565b505f805b8451811015611e7e57848181518110611e4d57611e4d614efd565b60200260200101516001600160a01b0316836001600160a01b031603611e765760019150611e7e565b600101611e32565b5080611e8b578460010194505b5050600101611d8d565b50603554829060ff16811015611ec157604051632fcba65760e11b8152600401610a4b91815260200190565b5050505050505050565b6098545f906060906001600160a01b0390811690861603611ef457611ef16064846150ce565b92505b611f0086868686612c62565b915091505b94509492505050565b611ec16040518060a00160405280898152602001886001600160a01b03168152602001876001600160a01b0316815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509390945250508b815260036020526040902091905083612cd6565b60015f5160206154bf5f395f51905f5255565b33611fd77f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610ddf5760405163118cdaa760e01b8152336004820152602401610a4b565b5f6115018383612d9b565b5f828152600360205260409020610cab9082612de7565b61202a612e7c565b5f51602061549f5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200161195a565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061210157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166120f55f51602061547f5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15610ddf5760405163703e46dd60e11b815260040160405180910390fd5b6109fa611fa5565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612181575060408051601f3d908101601f1916820190925261217e91810190615078565b60015b6121a957604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a4b565b5f51602061547f5f395f51905f5281146121d957604051632a87526960e21b815260048101829052602401610a4b565b6121e38383612eab565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ddf5760405163703e46dd60e11b815260040160405180910390fd5b60605f61150183612f00565b6066545f9081906001600160a01b031661225b57505f90508061234c565b60665460405163ae76638960e01b8152600481018990526001600160a01b038881166024830152604482018890525f92169063ae76638990606401606060405180830381865afa1580156122b1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122d5919061508f565b909450925090508086108015906122ec5750828510155b80156122f85750818410155b8184848989899091929394956123445760405163769296ad60e01b81526004810196909652602486019490945260448501929092526064840152608483015260a482015260c401610a4b565b505050505050505b9550959350505050565b61236c8a8a8a89612367898b6150ed565b612f59565b61237e8a8a8a8a8a8a8a8a8a8a6130ff565b5f8a81526003602052604090206001908101805490910190555b50505050505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b8361241f86886150ed565b61242991906150ed565b8360400151101589879091612461576040516224b63960e71b81526001600160a01b0390921660048301526024820152604401610a4b565b505061246c836131f7565b6123988a8a8a8a8a8a8a60018a8a612356565b5f818152600483016020526040902080546060919061249d90615100565b80601f01602080910402602001604051908101604052809291908181526020018280546124c990615100565b80156125145780601f106124eb57610100808354040283529160200191612514565b820191905f5260205f20905b8154815290600101906020018083116124f757829003601f168201915b5050505050905092915050565b612529611c48565b5f51602061549f5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612063565b6125716132f4565b61257a3361333d565b6125838461334e565b61258c826133b4565b6125946134f2565b61259c6134fa565b6125a461350a565b6001600160a01b038116156125cf57606680546001600160a01b0319166001600160a01b0383161790555b6001600160a01b038316612614576040516211211f60e21b815260206004820152600c60248201526b1c995dd85c9915d85b1b195d60a21b6044820152606401610a4b565b5050606780546001600160a01b0319166001600160a01b03929092169190911790555043606855565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061545f5f395f51905f529161267b90615100565b80601f01602080910402602001604051908101604052809291908181526020018280546126a790615100565b80156126f25780601f106126c9576101008083540402835291602001916126f2565b820191905f5260205f20905b8154815290600101906020018083116126d557829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061545f5f395f51905f529161267b90615100565b6127486007830182611c78565b819061277357604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b0381165f9081526009830160205260409020600101548190600160a81b900460ff166127c657604051630196157d60e31b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f908152600990910160205260409020600101805460ff60a81b19169055565b6127f86141b4565b5f82815260038085016020908152604092839020835160a0810185528154815260018201546001600160a01b039081169382019390935260028201549092169382019390935290820154606082015260048201805491929160808401919061285f90615100565b80601f016020809104026020016040519081016040528092919081815260200182805461288b90615100565b80156128d65780601f106128ad576101008083540402835291602001916128d6565b820191905f5260205f20905b8154815290600101906020018083116128b957829003601f168201915b505050505081525050905092915050565b6001808201545f916109e991906150ed565b5f6109e9825490565b5f611501838361351a565b61291a6007830182611c78565b819061294557604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b0381165f9081526009830160205260409020600101548190600160a81b900460ff161561299957604051635e31435560e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f908152600990910160205260409020600101805460ff60a81b1916600160a81b179055565b60028101545f906109e99060016150ed565b5f8181526001830160205260408120541515611501565b6001600160a01b038216612a355760405162d9127960e71b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a4b565b612a426007850183612be2565b8290612a6d57604051632f644c4360e21b81526001600160a01b039091166004820152602401610a4b565b506040805160a0810182526001600160a01b0393841680825292841660208083019182529515158284019081525f6060840181815260808501828152968252600990990190975292909520905181549085166001600160a01b03199091161781559351600185018054925196511515600160a81b0260ff60a81b19971515600160a01b026001600160a81b031990941692909516919091179190911794909416919091179092559051600290910155565b6001600160a01b038116612b615760405162d9127960e71b815260206004820152600a6024820152693637b1b0b62a37b5b2b760b11b6044820152606401610a4b565b612b6e6007830182612bf6565b8190612b9957604051637f7beaf760e01b81526001600160a01b039091166004820152602401610a4b565b506001600160a01b03165f90815260099091016020526040812080546001600160a01b03191681556001810180546001600160b01b031916905560020155565b5f61185f613540565b5f611501836001600160a01b038416612d9b565b5f611501836001600160a01b0384166135b3565b5f6109e9612c16612bd9565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612c4688888888613696565b925092509250612c56828261375e565b50909695505050505050565b5f60605f196001600160a01b03861601612ca157612c896001600160a01b03851684613816565b505060408051602081019091525f8152600190611f05565b8215611f0557612cb186866138a2565b15612ccb57612cc2868686866138da565b91509150611f05565b612cc2858585613a8c565b8151612ce56005850182612000565b8190612d075760405163980242c960e01b8152600401610a4b91815260200190565b505f8181526003858101602090815260409283902086518155908601516001820180546001600160a01b03199081166001600160a01b03938416179091559387015160028301805490951691161790925560608501519082015560808401518491906004820190612d78908261517c565b5050505f8181526004850160205260409020612d94838261517c565b5050505050565b5f818152600183016020526040812054612de057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109e9565b505f6109e9565b612df46005830182613c33565b8190612e1657604051632240947160e21b8152600401610a4b91815260200190565b505f8181526004830160205260408120612e2f916141f1565b5f8181526003808401602052604082208281556001810180546001600160a01b0319908116909155600282018054909116905590810182905590612e7660048301826141f1565b50505050565b5f51602061549f5f395f51905f525460ff16610ddf57604051638dfc202b60e01b815260040160405180910390fd5b612eb482613c3e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612ef8576121e38282613ca1565b610cab613d13565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612f4d57602002820191905f5260205f20905b815481526020019060010190808311612f39575b50505050509050919050565b5f196001600160a01b03851601612fcc57612f7481836150ed565b3414612f8082846150ed565b349091612fa95760405163341382c560e11b815260048101929092526024820152604401610a4b565b50508015612fc757606754612fc7906001600160a01b031682613816565b612d94565b5f348015612ff65760405163341382c560e11b815260048101929092526024820152604401610a4b565b5061301a9050833061300884866150ed565b6001600160a01b038816929190613d32565b801561303a5760675461303a906001600160a01b03868116911683613d99565b61304485856138a2565b1561305457612fc7858584613dca565b604051632770a7eb60e21b8152306004820152602481018390526001600160a01b03851690639dc29fac906044016020604051808303815f875af115801561309e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130c29190615236565b848484909192611ec15760405163bc8ecf7b60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610a4b565b5f5f61310a8c611864565b915061313f8c8c5f9182526003602090815260408084206001600160a01b039384168552600901909152909120600101541690565b9050896001600160a01b0316828d7f65b3e69df9006ac5b2f943ef9e7b59b447c491764e645bbf1dd8cf7f32eca5d78e858e8e428d8d8d60405161318a989796959493929190615251565b60405180910390a4896001600160a01b03168b6001600160a01b0316837f1f9715c5cdb86339f28cd3b93f6162d190a5c00b9305a65ee0889d1a7c00304d8a8a6040516131e1929190918252602082015260400190565b60405180910390a4505050505050505050505050565b6020818101516040808401516060850151608086015160a087015160c088015185516001600160a01b0390971660248801523060448801526064870194909452608486019290925260ff1660a485015260c484015260e48084019190915281518084039091018152610104909201905280820180516001600160e01b031663d505accf60e01b17815283518251929390925f92839291839182875af1806132a3576040513d5f823e3d81fd5b50505f513d915081156132ba5780600114156132c8565b84516001600160a01b03163b155b15612d9457845160405163733eb19560e01b81526001600160a01b039091166004820152602401610a4b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610ddf57604051631afcd79f60e31b815260040160405180910390fd5b6133456132f4565b6109fa81613de2565b6133566132f4565b61339e604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250613dea565b6035805460ff191660ff92909216919091179055565b6001600160a01b0381166133c55750565b5f816001600160a01b03166324c12bf66040518163ffffffff1660e01b81526004015f60405180830381865afa158015613401573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261342891908101906152a9565b604080513060208201520160408051601f19818403018152908290526134519291602001615328565b60408051601f1981840301815282825246602084015292505f9161348e918391016040516020818303038152906040528051906020012084613dfc565b90506001600160a01b0381166134cf5760405162461bcd60e51b8152600401610a4b906020808252600490820152637a65726f60e01b604082015260600190565b5f80546001600160a01b0319166001600160a01b03929092169190911790555050565b610ddf6132f4565b6135026132f4565b610ddf613e90565b6135126132f4565b610ddf613eb0565b5f825f01828154811061352f5761352f614efd565b905f5260205f200154905092915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61356a613eb8565b613572613f20565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f818152600183016020526040812054801561368d575f6135d560018361533c565b85549091505f906135e89060019061533c565b9050808214613647575f865f01828154811061360657613606614efd565b905f5260205f200154905080875f01848154811061362657613626614efd565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806136585761365861534f565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109e9565b5f9150506109e9565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156136cf57505f91506003905082613754565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613720573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661374b57505f925060019150829050613754565b92505f91508190505b9450945094915050565b5f82600381111561377157613771615363565b0361377a575050565b600182600381111561378e5761378e615363565b036137ac5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156137c0576137c0615363565b036137e15760405163fce698f760e01b815260048101829052602401610a4b565b60038260038111156137f5576137f5615363565b03610cab576040516335e2f38360e21b815260048101829052602401610a4b565b804710156138405760405163cf47918160e01b815247600482015260248101829052604401610a4b565b5f5f836001600160a01b0316836040515f6040518083038185875af1925050503d805f811461388a576040519150601f19603f3d011682016040523d82523d5f602084013e61388f565b606091505b509150915081612e7657612e7681613f62565b5f9182526003602090815260408084206001600160a01b0393909316845260099092019052902060010154600160a01b900460ff1690565b60405163a9059cbb60e01b81526001600160a01b038381166004830152602482018390525f9160609186169063a9059cbb906044016020604051808303815f875af1925050508015613949575060408051601f3d908101601f1916820190925261394691810190615236565b60015b613a1b57613955615377565b806308c379a00361397e575061396961538f565b8061397457506139d8565b5f92509050611f05565b634e487b71036139d857613990615411565b9061399b57506139d8565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b604051602081830303815290604052915050611f05565b3d808015613a01576040519150601f19603f3d011682016040523d82523d5f602084013e613a06565b606091505b505f9250806040516020016139c1919061542e565b8015613a46576001925060405180602001604052805f8152509150613a41878786613f8b565b613a82565b5f92506040518060400160405280601f81526020017f5374616e646172644272696467653a207472616e73666572206661696c65640081525091505b5094509492505050565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390525f916060918616906340c10f19906044016020604051808303815f875af1925050508015613afb575060408051601f3d908101601f19168201909252613af891810190615236565b60015b613bcd57613b07615377565b806308c379a003613b305750613b1b61538f565b80613b265750613b8a565b5f92509050613c2b565b634e487b7103613b8a57613b42615411565b90613b4d5750613b8a565b6040516b02830b734b19031b7b2329d160a51b6020820152602c81018290525f9350604c015b604051602081830303815290604052915050613c2b565b3d808015613bb3576040519150601f19603f3d011682016040523d82523d5f602084013e613bb8565b606091505b505f925080604051602001613b73919061542e565b8015613bed576001925060405180602001604052805f8152509150613c29565b5f92506040518060400160405280601b81526020017f5374616e646172644272696467653a206d696e74206661696c6564000000000081525091505b505b935093915050565b5f61150183836135b3565b806001600160a01b03163b5f03613c7357604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a4b565b5f51602061547f5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613cbd9190615453565b5f60405180830381855af49150503d805f8114613cf5576040519150601f19603f3d011682016040523d82523d5f602084013e613cfa565b606091505b5091509150613d0a858383613fa3565b95945050505050565b3415610ddf5760405163b398979f60e01b815260040160405180910390fd5b6040516001600160a01b038481166024830152838116604483015260648201839052612e769186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613fff565b6040516001600160a01b038381166024830152604482018390526121e391859182169063a9059cbb90606401613d67565b5f8381526003602052604090206121e390838361406b565b611aa66132f4565b613df26132f4565b610cab82826140a1565b5f83471015613e275760405163cf47918160e01b815247600482015260248101859052604401610a4b565b81515f03613e4857604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d151981151615613e69576040513d5f823e3d81fd5b6001600160a01b0381166115015760405163b06ebf3d60e01b815260040160405180910390fd5b613e986132f4565b5f51602061549f5f395f51905f52805460ff19169055565b611f926132f4565b5f5f51602061545f5f395f51905f5281613ed061263d565b805190915015613ee857805160209091012092915050565b81548015613ef7579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061545f5f395f51905f5281613f386126fd565b805190915015613f5057805160209091012092915050565b60018201548015613ef7579392505050565b805115613f725780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f8381526003602052604090206121e3908383614100565b606082613fb857613fb382613f62565b611501565b8151158015613fcf57506001600160a01b0384163b155b15613ff857604051639996b31560e01b81526001600160a01b0385166004820152602401610a4b565b5080611501565b5f5f60205f8451602086015f885af18061401e576040513d5f823e3d81fd5b50505f513d91508115614035578060011415614042565b6001600160a01b0384163b155b15612e7657604051635274afe760e01b81526001600160a01b0385166004820152602401610a4b565b6001600160a01b0382165f908152600984016020526040812060020180548392906140979084906150ed565b9091555050505050565b6140a96132f4565b5f51602061545f5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026140e2848261517c565b50600381016140f1838261517c565b505f8082556001909101555050565b6001600160a01b0382165f90815260098401602052604081206002015481906141299084614190565b86549193509150848484614169576040516302ea38e160e51b815260048101939093526001600160a01b0390911660248301526044820152606401610a4b565b5050506001600160a01b039093165f90815260099094016020525050604090912060020155565b5f5f838311156141a457505f9050806141ad565b50600190508183035b9250929050565b6040518060a001604052805f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b5080546141fd90615100565b5f825580601f1061420c575050565b601f0160209004905f5260205f20908101906109fa91905b80821115614237575f8155600101614224565b5090565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081018181106001600160401b03821117156142825761428261424f565b60405250565b601f8201601f191681016001600160401b03811182821017156142ad576142ad61424f565b6040525050565b5f6001600160401b038211156142cc576142cc61424f565b5060051b60200190565b5f5f604083850312156142e7575f5ffd5b8235915060208301356001600160401b03811115614303575f5ffd5b8301601f81018513614313575f5ffd5b803561431e816142b4565b60405161432b8282614288565b80915082815260208101915060208360051b85010192508783111561434e575f5ffd5b6020840193505b82841015614370578335825260209384019390910190614355565b809450505050509250929050565b6001600160a01b03811681146109fa575f5ffd5b803561439d8161437e565b919050565b5f602082840312156143b2575f5ffd5b81356115018161437e565b5f5f83601f8401126143cd575f5ffd5b5081356001600160401b038111156143e3575f5ffd5b6020830191508360208285010111156141ad575f5ffd5b803560ff8116811461439d575f5ffd5b5f82601f830112614419575f5ffd5b8135614424816142b4565b6040516144318282614288565b80915082815260208101915060208360051b860101925085831115614454575f5ffd5b602085015b838110156144785761446a816143fa565b835260209283019201614459565b5095945050505050565b5f82601f830112614491575f5ffd5b813561449c816142b4565b6040516144a98282614288565b80915082815260208101915060208360051b8601019250858311156144cc575f5ffd5b602085015b838110156144785780358352602092830192016144d1565b5f5f5f5f5f5f5f5f5f5f6101208b8d031215614503575f5ffd5b8a35995060208b0135985061451a60408c01614392565b975061452860608c01614392565b965060808b0135955060a08b01356001600160401b03811115614549575f5ffd5b6145558d828e016143bd565b90965094505060c08b01356001600160401b03811115614573575f5ffd5b61457f8d828e0161440a565b93505060e08b01356001600160401b0381111561459a575f5ffd5b6145a68d828e01614482565b9250506101008b01356001600160401b038111156145c2575f5ffd5b6145ce8d828e01614482565b9150509295989b9194979a5092959850565b5f602082840312156145f0575f5ffd5b5035919050565b5f60208284031215614607575f5ffd5b81356001600160401b0381111561461c575f5ffd5b8201601f8101841361462c575f5ffd5b8035614637816142b4565b6040516146448282614288565b80915082815260208101915060208360051b850101925086831115614667575f5ffd5b6020840193505b828410156146925783356146818161437e565b82526020938401939091019061466e565b9695505050505050565b5f5f604083850312156146ad575f5ffd5b50508035926020909101359150565b5f6001600160401b038211156146d4576146d461424f565b50601f01601f191660200190565b5f6146ec836146bc565b6040516146f98282614288565b80925084815285858501111561470d575f5ffd5b848460208301375f6020868301015250509392505050565b5f5f60408385031215614736575f5ffd5b82356147418161437e565b915060208301356001600160401b0381111561475b575f5ffd5b8301601f8101851361476b575f5ffd5b61477a858235602084016146e2565b9150509250929050565b80516001600160a01b03908116835260208083015190911690830152604080820151151590830152606080820151151590830152608090810151910152565b602080825282518282018190525f918401906040840190835b81811015614805576147ef838551614784565b6020939093019260a092909201916001016147dc565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b031215614827575f5ffd5b8835975060208901356148398161437e565b965060408901356148498161437e565b9550606089013594506080890135935060a0890135925060c08901356001600160401b03811115614878575f5ffd5b6148848b828c016143bd565b999c989b5096995094979396929594505050565b5f5f5f5f608085870312156148ab575f5ffd5b8435935060208501356148bd8161437e565b925060408501356001600160401b038111156148d7575f5ffd5b8501601f810187136148e7575f5ffd5b6148f6878235602084016146e2565b925050614905606086016143fa565b905092959194509250565b5f82601f83011261491f575f5ffd5b813561492a816142b4565b6040516149378282614288565b80915082815260208101915060208360051b86010192508583111561495a575f5ffd5b602085015b838110156144785780356001600160401b0381111561497c575f5ffd5b61498b886020838a010161440a565b8452506020928301920161495f565b5f82601f8301126149a9575f5ffd5b81356149b4816142b4565b6040516149c18282614288565b80915082815260208101915060208360051b8601019250858311156149e4575f5ffd5b602085015b838110156144785780356001600160401b03811115614a06575f5ffd5b614a15886020838a0101614482565b845250602092830192016149e9565b5f5f5f5f5f5f60a08789031215614a39575f5ffd5b8635955060208701356001600160401b03811115614a55575f5ffd5b8701601f81018913614a65575f5ffd5b80356001600160401b03811115614a7a575f5ffd5b8960208260051b8401011115614a8e575f5ffd5b6020919091019550935060408701356001600160401b03811115614ab0575f5ffd5b614abc89828a01614910565b93505060608701356001600160401b03811115614ad7575f5ffd5b614ae389828a0161499a565b92505060808701356001600160401b03811115614afe575f5ffd5b614b0a89828a0161499a565b9150509295509295509295565b5f5f5f5f5f5f5f5f5f5f8a8c036101e0811215614b32575f5ffd5b8b359a5060208c0135614b448161437e565b995060408c0135614b548161437e565b985060608c0135614b648161437e565b975060808c0135965060a08c0135955060c08c0135945060e060df1982011215614b8c575f5ffd5b50604051614b9981614263565b60e08c0135614ba78161437e565b81526101008c0135614bb88161437e565b60208201526101208c013560408201526101408c01356060820152614be06101608d016143fa565b60808201526101808c013560a08201526101a08c013560c082015292506101c08b01356001600160401b03811115614c16575f5ffd5b614c228d828e016143bd565b915080935050809150509295989b9194979a5092959850565b5f5f60408385031215614c4c575f5ffd5b823591506020830135614c5e8161437e565b809150509250929050565b60a081016109e98284614784565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6115016020830184614c77565b5f5f5f5f5f5f60c08789031215614ccc575f5ffd5b863595506020870135614cde8161437e565b9450614cec604088016143fa565b93506060870135614cfc8161437e565b92506080870135614d0c8161437e565b915060a0870135614d1c8161437e565b809150509295509295509295565b60ff60f81b8816815260e060208201525f614d4860e0830189614c77565b8281036040840152614d5a8189614c77565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015614daf578351835260209384019390920191600101614d91565b50909b9a5050505050505050505050565b602081528151602082015260018060a01b03602083015116604082015260018060a01b036040830151166060820152606082015160808201525f608083015160a08084015261127760c0840182614c77565b5f5f5f60608486031215614e24575f5ffd5b833592506020840135614e368161437e565b929592945050506040919091013590565b5f60208284031215614e57575f5ffd5b611501826143fa565b80151581146109fa575f5ffd5b5f5f5f5f60808587031215614e80575f5ffd5b843593506020850135614e9281614e60565b92506040850135614ea28161437e565b91506060850135614eb28161437e565b939692955090935050565b602080825282518282018190525f918401906040840190835b818110156148055783516001600160a01b0316835260209384019390920191600101614ed6565b634e487b7160e01b5f52603260045260245ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b878152602081018790526001600160a01b038681166040830152851660608201526080810184905260c060a082018190525f90614f799083018486614f11565b9998505050505050505050565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f611501600d830184614f86565b848152608060208201525f614fd66080830186614c77565b8281036040840152614fe88186614c77565b91505060ff8316606083015295945050505050565b5f6020828403121561500d575f5ffd5b81516115018161437e565b5f8235609e1983360301811261502c575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261504b575f5ffd5b8301803591506001600160401b03821115615064575f5ffd5b6020019150368190038213156141ad575f5ffd5b5f60208284031215615088575f5ffd5b5051919050565b5f5f5f606084860312156150a1575f5ffd5b5050815160208301516040909301519094929350919050565b634e487b7160e01b5f52601160045260245ffd5b5f826150e857634e487b7160e01b5f52601260045260245ffd5b500490565b808201808211156109e9576109e96150ba565b600181811c9082168061511457607f821691505b60208210810361513257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156121e357805f5260205f20601f840160051c8101602085101561515d5750805b601f840160051c820191505b81811015612d94575f8155600101615169565b81516001600160401b038111156151955761519561424f565b6151a9816151a38454615100565b84615138565b6020601f8211600181146151db575f83156151c45750848201515b5f19600385901b1c1916600184901b178455612d94565b5f84815260208120601f198516915b8281101561520a57878501518255602094850194600190920191016151ea565b508482101561522757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215615246575f5ffd5b815161150181614e60565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905283151560a082015260e060c082018190525f9061529b9083018486614f11565b9a9950505050505050505050565b5f602082840312156152b9575f5ffd5b81516001600160401b038111156152ce575f5ffd5b8201601f810184136152de575f5ffd5b80516152e9816146bc565b6040516152f68282614288565b82815286602084860101111561530a575f5ffd5b8260208501602083015e5f9281016020019290925250949350505050565b5f6112776153368386614f86565b84614f86565b818103818111156109e9576109e96150ba565b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b5f60033d1115610e1f5760045f5f3e505f5160e01c90565b5f60443d101561539c5790565b6040513d600319016004823e80513d60248201116001600160401b03821117156153c557505090565b80820180516001600160401b038111156153e0575050505090565b3d84016003190182820160200111156153fa575050505090565b61540960208285010185614288565b509392505050565b5f5f60233d111561542a57602060045f3e50505f516001905b9091565b7002637bb96b632bb32b61032b93937b91d1607d1b81525f6115016011830184614f86565b5f6115018284614f8656fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122031e1eefda7bcaf4372e1fd2157044cc34817e271f9dde3153549651e926841b464736f6c634300081c0033",
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

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IChainManagerTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IChainManagerTokenPair)).(*[]IChainManagerTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
	return _EthereumBridge.Contract.AllTokenPairs(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256)[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IChainManagerTokenPair, error) {
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
func (_EthereumBridge *EthereumBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IChainManagerTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IChainManagerTokenPair)).(*IChainManagerTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_EthereumBridge *EthereumBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
	return _EthereumBridge.Contract.GetTokenPair(&_EthereumBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256))
func (_EthereumBridge *EthereumBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IChainManagerTokenPair, error) {
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
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeCaller) RevertedArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "revertedArguments", remoteChainID, index)

	if err != nil {
		return *new(IChainManagerFinalizeArguments), err
	}

	out0 := *abi.ConvertType(out[0], new(IChainManagerFinalizeArguments)).(*IChainManagerFinalizeArguments)

	return out0, err

}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
	return _EthereumBridge.Contract.RevertedArguments(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// RevertedArguments is a free data retrieval call binding the contract method 0x952a95de.
//
// Solidity: function revertedArguments(uint256 remoteChainID, uint256 index) view returns((uint256,address,address,uint256,bytes))
func (_EthereumBridge *EthereumBridgeCallerSession) RevertedArguments(remoteChainID *big.Int, index *big.Int) (IChainManagerFinalizeArguments, error) {
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

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "finalizeBridge", fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) FinalizeBridge(fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridge(&_EthereumBridge.TransactOpts, fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x15eb39ef.
//
// Solidity: function finalizeBridge(uint256 fromChainID, uint256 index, address token, address to, uint256 value, bytes extraData, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) FinalizeBridge(fromChainID *big.Int, index *big.Int, token common.Address, to common.Address, value *big.Int, extraData []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridge(&_EthereumBridge.TransactOpts, fromChainID, index, token, to, value, extraData, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "finalizeBridgeBatch", fromChainID, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) FinalizeBridgeBatch(fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridgeBatch(&_EthereumBridge.TransactOpts, fromChainID, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x7c7505dc.
//
// Solidity: function finalizeBridgeBatch(uint256 fromChainID, (uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) FinalizeBridgeBatch(fromChainID *big.Int, args []IChainManagerFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.FinalizeBridgeBatch(&_EthereumBridge.TransactOpts, fromChainID, args, v, r, s)
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
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) RetryFinalizeBridge(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "retryFinalizeBridge", fromChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeSession) RetryFinalizeBridge(fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridge(&_EthereumBridge.TransactOpts, fromChainID, index)
}

// RetryFinalizeBridge is a paid mutator transaction binding the contract method 0x3960e787.
//
// Solidity: function retryFinalizeBridge(uint256 fromChainID, uint256 index) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) RetryFinalizeBridge(fromChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridge(&_EthereumBridge.TransactOpts, fromChainID, index)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) RetryFinalizeBridgeBatch(opts *bind.TransactOpts, fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "retryFinalizeBridgeBatch", fromChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeSession) RetryFinalizeBridgeBatch(fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridgeBatch(&_EthereumBridge.TransactOpts, fromChainID, indexes)
}

// RetryFinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x030372c3.
//
// Solidity: function retryFinalizeBridgeBatch(uint256 fromChainID, uint256[] indexes) returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) RetryFinalizeBridgeBatch(fromChainID *big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RetryFinalizeBridgeBatch(&_EthereumBridge.TransactOpts, fromChainID, indexes)
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, index []*big.Int, token []common.Address, to []common.Address) (*EthereumBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeFinalizedIterator{contract: _EthereumBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeFinalized, index []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeFinalized", indexRule, tokenRule, toRule)
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

// ParseBridgeFinalized is a log parse operation binding the contract event 0xc5e42566da2b6c721c96e7b791f192942f342ce351acc427f19def2881f5a373.
//
// Solidity: event BridgeFinalized(uint256 indexed index, address indexed token, address indexed to, uint256 value, uint256 time)
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
