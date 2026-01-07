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

// IBaseBridgeBridgeTokenArguments is an auto generated low-level Go binding around an user-defined struct.

// IBaseBridgePermitArguments is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryPendingData is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.

// BSCBridgeV2MetaData contains all meta data concerning the BSCBridgeV2 contract.
var BSCBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deadWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alreadyTransferred\",\"type\":\"bool\"}],\"name\":\"burnCrossToDeadWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross_\",\"type\":\"address\"}],\"name\":\"reinitializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeInvalidDeadWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"a10bab0b": "bridgeExecutor()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
		"86e9e175": "burnCrossToDeadWallet(address,uint256,bool)",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"b33eb36e": "getPendingArguments(uint256,uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"814914b5": "getTokenPair(uint256,address)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"b2b49e2e": "grantRoleBatch(bytes32[],address[])",
		"91d14854": "hasRole(bytes32,address)",
		"89232a00": "initialize(address,address,uint8)",
		"91cf6d3e": "initializedAt()",
		"761fe2d8": "isTokenFinalizePaused(uint256,address)",
		"aa20ed47": "manualReleasePending(uint256,uint256)",
		"48a00ed8": "manualReleasePendingWithRecipient(uint256,uint256,address)",
		"1e24abdb": "maxExtraDataLength()",
		"5c975abb": "paused()",
		"1313996b": "permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"edbbea39": "registerToken(uint256,bool,address,address)",
		"7b429ad8": "reinitializeBSCBridge(uint256,address)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"8ae87c5c": "removePending(uint256,uint256)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"cba25e4b": "setBridgeExecutor(address)",
		"0b1d8d06": "setBridgeVerifier(address)",
		"b8aa837e": "setChainPause(uint256,bool)",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"56ff54b0": "setMaxExtraDataLength(uint256)",
		"bedb86fb": "setPause(bool)",
		"bfbf6765": "setTokenPause(uint256,address,bool,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f516020615eb85f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615df190816100c78239608051818181610b810152610d3a0152f35b6001600160401b0319166001600160401b039081175f516020615eb85f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b610020612fcd565b005b5f3560e01c806301ffc9a7146103815780630b1d8d061461037c5780631313996b146103775780631e24abdb14610372578063248a9ca31461036d5780632f2ff15d1461036857806336568abe14610363578063365d71e41461035e57806342cde4e81461035957806348a00ed8146103545780634be13f831461034f5780634ee078ba1461034a5780634f1ef28614610345578063502cc5c01461034057806352d1902d1461033b57806356ff54b0146103365780635b605f5c146103315780635c975abb1461032c5780635fd262de14610327578063670e626814610322578063761fe2d81461031d57806379214874146103185780637b429ad814610313578063814914b51461030e57806384b0196e1461030957806386e9e1751461030457806388d67d6d146102ff57806389232a00146102fa5780638ae87c5c146102f557806391cca3db146102f057806391cf6d3e146102eb57806391d14854146102e6578063a10bab0b146102e1578063a217fddf146102dc578063a3246ad3146102d7578063aa1bd0bc146102d2578063aa20ed47146102cd578063ad3cb1cc146102c8578063ae6893f8146102c3578063b2b49e2e146102be578063b33eb36e146102b9578063b7f3358d146102b4578063b8aa837e146102af578063bedb86fb146102aa578063bfbf6765146102a5578063cba25e4b146102a0578063cf56118e1461029b578063d477f05f14610296578063d547741f14610291578063d5717fc51461028c578063e2af6cd714610287578063edbbea3914610282578063f0702e8e1461027d5763f698da250361000e576123eb565b6123c3565b61227d565b612204565b6121cb565b61219c565b61213d565b6120c9565b61205e565b611f85565b611e9b565b611dff565b611d7b565b611cea565b611bdc565b611ba3565b611b5c565b611ad1565b611a84565b611a08565b6119ac565b611984565b611942565b611925565b6118fd565b611894565b611798565b611631565b61147a565b6113b6565b611269565b611193565b611123565b61109a565b611023565b610f75565b610f47565b610e64565b610d7f565b610d28565b610c98565b610b2e565b6109c0565b610999565b61086e565b610842565b6107d9565b61059a565b61055f565b610539565b61051c565b61049a565b6103f8565b346103d75760203660031901126103d75760043563ffffffff60e01b81168091036103d757602090637965db0b60e01b81149081156103c6575b506040519015158152f35b6301ffc9a760e01b1490505f6103bb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103d757565b346103d75760203660031901126103d757600435610415816103e7565b61041e336139f8565b6001600160a01b0316610432811515612457565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103d7578235916001600160401b0383116103d7576020808501948460051b0101116103d757565b60403660031901126103d7576004356001600160401b0381116103d7576104c590369060040161046a565b602435916001600160401b0383116103d757366023840112156103d7576004830135916001600160401b0383116103d75736602460e08502860101116103d757602461002094019161248c565b5f9103126103d757565b346103d7575f3660031901126103d7576020600a54604051908152f35b346103d75760203660031901126103d757602061055760043561264d565b604051908152f35b346103d75760403660031901126103d757610020602435600435610582826103e7565b61059561058e8261264d565b3390613c18565b613306565b346103d75760403660031901126103d7576004356024356105ba816103e7565b336001600160a01b038216036105d35761002091613366565b63334bd91960e11b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b0382111761061157604052565b6105e2565b606081019081106001600160401b0382111761061157604052565b60c081019081106001600160401b0382111761061157604052565b601f909101601f19168101906001600160401b0382119082101761061157604052565b6040519061067e60808361064c565b565b6040519061067e60e08361064c565b6040519061067e6101008361064c565b6040519061067e60c08361064c565b6040519061067e60608361064c565b6001600160401b0381116106115760051b60200190565b9080601f830112156103d75781356106eb816106bd565b926106f9604051948561064c565b81845260208085019260051b8201019283116103d757602001905b8282106107215750505090565b8135815260209182019101610714565b9060406003198301126103d7576004356001600160401b0381116103d7578261075c916004016106d4565b91602435906001600160401b0382116103d757806023830112156103d7578160040135610788816106bd565b92610796604051948561064c565b8184526024602085019260051b8201019283116103d757602401905b8282106107bf5750505090565b6020809183356107ce816103e7565b8152019101906107b2565b346103d7576107e736610731565b906107f5815183511461266b565b5f5b8251811015610020578061083b61081060019385612681565b51838060a01b036108218488612681565b511690610836336108318361264d565b613c18565b613366565b50016107f7565b346103d7575f3660031901126103d757602060ff5f516020615adc5f395f51905f525416604051908152f35b346103d75760603660031901126103d757602435600435604435610891816103e7565b61089a33613a80565b6108a2613017565b815f5260076020526108c0836108bb8160405f206144e2565b612695565b806108cb848461436b565b916001600160a01b031615610985575b8151905f516020615afc5f395f51905f52602084019161093e8351956109346109216040830197885160018060a01b03169986608086019b8c519260a08801519461477b565b5061092b81611c4f565b600181146133c5565b51935194516103db565b94516040516001600160a01b039096169591829161095f91429190846133e8565b0390a45f516020615a9c5f395f51905f525f80a35f5f516020615c9c5f395f51905f525d005b5060608101516001600160a01b03166108db565b346103d7575f3660031901126103d7575f546040516001600160a01b039091168152602090f35b346103d75760403660031901126103d757610acb6024356004356109e2612ff0565b6109ea613017565b610a1081610a0b610a0760036109ff846125e1565b015460ff1690565b1590565b6126af565b610a2f826108bb81610a2a855f52600760205260405f2090565b6144e2565b610ac66040610a5f610a5a85610a4d865f52600860205260405f2090565b905f5260205260405f2090565b612814565b610ab3610a7f82516080610a75868301516103db565b91015190876134cf565b50610a8981611c4f565b610a9281611c4f565b83516020808201839052815290600190610aad60408461064c565b1461284d565b01518015908115610ad3575b429161287d565b613406565b61002061304c565b4281109150610abf565b6001600160401b03811161061157601f01601f191660200190565b929192610b0482610add565b91610b12604051938461064c565b8294818452818301116103d7578281602093845f960137010152565b60403660031901126103d757600435610b46816103e7565b6024356001600160401b0381116103d757366023820112156103d757610b76903690602481600401359101610af8565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610c76575b50610c6757610bba336139f8565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610c36575b50610c0357634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615b9c5f395f51905f528303610c22576100209250614a9e565b632a87526960e21b5f52600483905260245ffd5b610c5991945060203d602011610c60575b610c51818361064c565b8101906135f6565b925f610be2565b503d610c47565b63703e46dd60e11b5f5260045ffd5b5f516020615b9c5f395f51905f52546001600160a01b0316141590505f610bac565b346103d75760603660031901126103d757602435600435604435610cbb33613a80565b815f526007602052610cd4836108bb8160405f206144e2565b4201804211610d235760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b61289b565b346103d7575f3660031901126103d7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610c675760206040515f516020615b9c5f395f51905f528152f35b346103d75760203660031901126103d7577f1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f16020600435610dbf33613b08565b80600a55604051908152a1005b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610e445750505090565b909192602060e082610e596001948851610dcc565b019401929101610e37565b346103d75760203660031901126103d757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b818110610f2e575050610eb39250038361064c565b610ebd82516128f2565b915f5b8151811015610f1c57600190610f00610efb610edb866125ef565b610ef5610ee88588612681565b516001600160a01b031690565b90612941565b612956565b610f0a8287612681565b52610f158186612681565b5001610ec0565b60405180610f2a8682610e21565b0390f35b8454835260019485019487945060209093019201610e9e565b346103d7575f3660031901126103d757602060ff5f516020615c5c5f395f51905f5254166040519015158152f35b60e03660031901126103d757602435600435610f90826103e7565b604435610f9c816103e7565b60c4359160a4356084356064356001600160401b0386116103d757366023870112156103d7576004860135946001600160401b0386116103d75736602487890101116103d757610f2a976024610ff39801956129c9565b60405191829182901515815260200190565b60ff8116036103d757565b6001600160a01b03909116815260200190565b346103d75760803660031901126103d757602435600435611043826103e7565b604435906001600160401b0382116103d757366023830112156103d757610f2a9261107b61108e933690602481600401359101610af8565b906064359261108984611005565b612b48565b60405191829182611010565b346103d75760403660031901126103d757602060ff6110d06024356004356110c1826103e7565b5f526009845260405f20612941565b54166040519015158152f35b90602080835192838152019201905f5b8181106110f95750505090565b82518452602093840193909201916001016110ec565b9060206111209281815201906110dc565b90565b346103d75760203660031901126103d7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061117d57610f2a856111718187038261064c565b6040519182918261110f565b825484526020909301926001928301920161115a565b346103d75760403660031901126103d7576004356024356111b3816103e7565b6111bc336139f8565b5f516020615d1c5f395f51905f52549160ff8360401c168015611245575b611236576001600160401b03199092166002175f516020615d1c5f395f51905f525561120d91611208612c16565b612c79565b611215612c37565b604051600281525f516020615b1c5f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b03841610156111da565b60e08101929161067e9190610dcc565b346103d75760403660031901126103d757610f2a6112a860243560043561128f826103e7565b6112976128bc565b505f52600660205260405f20612941565b6004604051916112b7836105f6565b80546001600160a01b03168352600181015461130c90611303906112e66112dd826103db565b602088016126c9565b6112fa60a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611259565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161138c9061137e61112097959693600f60f81b865260e0602087015260e0860190611333565b908482036040860152611333565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526110dc565b346103d7575f3660031901126103d7575f516020615b5c5f395f51905f5254158061145a575b1561141d576113e96138a8565b6113f1613960565b90610f2a60405161140360208261064c565b5f8082523660208301376040519384933091469186611357565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615d3c5f395f51905f5254156113dc565b801515036103d757565b346103d75760603660031901126103d757600435611497816103e7565b6024356044356114a681611470565b6114ae612ff0565b6114b6613017565b6001600160a01b0383165f9081526066602052604090206114e2906114dd905b5460ff1690565b612d4f565b6114ed821515612c52565b6115a45761150881836115016065546103db565b3390613c64565b335b5f516020615d9c5f395f51905f52611523606454613cad565b926064549261158d6115366065546103db565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a461159861304c565b60405160018152602090f35b6115ad336139f8565b3061150a565b9080601f830112156103d75781356115ca816106bd565b926115d8604051948561064c565b81845260208085019260051b820101918383116103d75760208201905b83821061160457505050505090565b81356001600160401b0381116103d757602091611626878480948801016106d4565b8152019101906115f5565b60803660031901126103d7576004356001600160401b0381116103d75761165c90369060040161046a565b90602435906001600160401b0382116103d757366023830112156103d757816004013591611689836106bd565b92611697604051948561064c565b8084526024602085019160051b830101913683116103d75760248101915b83831061170d5750506044359150506001600160401b0381116103d7576116e09036906004016115b3565b606435926001600160401b0384116103d757610f2a94611707610ff39536906004016115b3565b93612d65565b82356001600160401b0381116103d7578201366043820112156103d757602481013590611739826106bd565b91611747604051938461064c565b808352602060248185019260051b84010101913683116103d757604401905b82821061177e575050508152602092830192016116b5565b60208091833561178d81611005565b815201910190611766565b346103d75760603660031901126103d7576004356117b5816103e7565b602435906117c2826103e7565b6044356117ce81611005565b5f516020615d1c5f395f51905f52549260ff604085901c1615936001600160401b03168015908161188c575b6001149081611882575b159081611879575b50611236575f516020615d1c5f395f51905f5280546001600160401b031916600117905561183e928461186c57613fc7565b61184457005b61184c612c37565b604051600181525f516020615b1c5f395f51905f52908060208101611231565b611874612c16565b613fc7565b9050155f61180c565b303b159150611804565b8591506117fa565b346103d75760403660031901126103d7576024356004356118b4336139f8565b6118bc613017565b6118c6828261436b565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615c9c5f395f51905f525d005b346103d7575f3660031901126103d7576033546040516001600160a01b039091168152602090f35b346103d7575f3660031901126103d7576020603454604051908152f35b346103d75760403660031901126103d757602060ff6110d0602435600435611969826103e7565b5f525f516020615c3c5f395f51905f52845260405f20612941565b346103d7575f3660031901126103d7576035546040516001600160a01b039091168152602090f35b346103d7575f3660031901126103d75760206040515f8152f35b60206040818301928281528451809452019201905f5b8181106119e95750505090565b82516001600160a01b03168452602093840193909201916001016119dc565b346103d75760203660031901126103d7576004355f525f516020615bbc5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611a6e57610f2a85611a628187038261064c565b604051918291826119c6565b8254845260209093019260019283019201611a4b565b346103d75760203660031901126103d7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611ac4336139f8565b80600155604051908152a1005b346103d75760403660031901126103d757602435600435611af133613a80565b611afa33613a80565b611b02613017565b805f526007602052611b1b826108bb8160405f206144e2565b611b258282613406565b5f516020615a9c5f395f51905f525f80a35f5f516020615c9c5f395f51905f525d005b60405190611b5760208361064c565b5f8252565b346103d7575f3660031901126103d757610f2a604051611b7d60408261064c565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611333565b346103d75760203660031901126103d7576004355f526004602052600160405f20015460018101809111610d2357602090604051908152f35b346103d757611bea36610731565b90611bf8815183511461266b565b5f5b82518110156100205780611c34611c1360019385612681565b51838060a01b03611c248488612681565b511690610595336108318361264d565b5001611bfa565b634e487b7160e01b5f52602160045260245ffd5b600b1115611c5957565b611c3b565b90600b821015611c595752565b6020815260606040611cd060a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611333565b93611ce2602082015183860190611c5e565b015191015290565b346103d75760403660031901126103d757600435602435905f60408051611d1081610616565b611d18612e42565b815282602082015201525f52600860205260405f20905f52602052610f2a60405f20600760405191611d4983610616565b611d5281612710565b8352611d6860ff60068301541660208501612808565b0154604082015260405191829182611c6b565b346103d75760203660031901126103d7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611dbd81611005565b611dc6336139f8565b16611dd2811515612e72565b8060ff195f516020615adc5f395f51905f525416175f516020615adc5f395f51905f5255604051908152a1005b346103d75760403660031901126103d757600435602435611e1f81611470565b611e2833613b90565b611e3d825f52600360205260405f2054151590565b15611e885760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611e7d81600360405f2001612c68565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103d75760203660031901126103d757600435611eb881611470565b611ec133613b90565b15611f1f57611ece612ff0565b600160ff195f516020615c5c5f395f51905f525416175f516020615c5c5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615c5c5f395f51905f525460ff811615611f765760ff19165f516020615c5c5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103d75760803660031901126103d757602435600435611fa5826103e7565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c6040604435611fd481611470565b606435611fe081611470565b611fe933613b90565b845f52600560205261204d81612048855f20986120188161201360018060a01b038216809d6144e2565b612e88565b885f526006602052612038866001612032848b5f20612941565b01612eb0565b885f526009602052865f20612941565b612c68565b8251911515825215156020820152a3005b346103d75760203660031901126103d75760043561207b816103e7565b612084336139f8565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103d7575f3660031901126103d757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b81811061212757610f2a856111718187038261064c565b8254845260209093019260019283019201612110565b346103d75760203660031901126103d75760043561215a816103e7565b612163336139f8565b6001600160a01b0316612177811515612457565b603380546001600160a01b031916821790555f516020615bdc5f395f51905f525f80a2005b346103d75760403660031901126103d7576100206024356004356121bf826103e7565b61083661058e8261264d565b346103d75760203660031901126103d7576004355f526004602052600260405f20015460018101809111610d2357602090604051908152f35b346103d75760203660031901126103d757600435612221816103e7565b61222a336139f8565b6001600160a01b031661223e811515612ecd565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346103d75760803660031901126103d75760043560243561229d81611470565b604435916122aa836103e7565b5f516020615cdc5f395f51905f5261238b606435936122c8856103e7565b6122d133613b08565b6122dc841515612ecd565b6001600160a01b03861694610ff3906122f6871515612ecd565b6001600160a01b03811697612386906123108a1515612ecd565b61231988615274565b612390575b612339816123348161232f8c6125fd565b614585565b61443b565b612358612344610680565b9361234f83866126c9565b602085016126c9565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152612381886125ef565b612941565b614463565b0390a4005b6123be61239b61066f565b8981525f60208201525f60408201525f60608201526123b98a6125e1565b614410565b61231e565b346103d7575f3660031901126103d7576032546040516001600160a01b039091168152602090f35b346103d7575f3660031901126103d7576124036157c1565b61240b615818565b6040519060208201925f516020615d5c5f395f51905f528452604083015260608201524660808201523060a082015260a0815261244960c08261064c565b519020604051908152602090f35b1561245e57565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b919092612497612ff0565b61249f613017565b6124aa818514612590565b5f5b8481106124c057505050505061067e61304c565b8061251a6125016124dd6124d76001958a8a6125ba565b356125fd565b6124fb6124f660206124f0878d8d6125ba565b0161260b565b6103db565b9061305e565b6125146124f660206124f0868c8c6125ba565b90612615565b61258a6125288288886125ba565b3561253960206124f0858b8b6125ba565b878961254b60606124f08884866125ba565b60c061257688608061255e8287896125ba565b01359460a061256e83838a6125ba565b0135966125ba565b013593612584888b8d61263d565b956130dc565b016124ac565b1561259757565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125dc5760051b8101359060fe19813603018212156103d7570190565b6125a6565b5f52600460205260405f2090565b5f52600660205260405f2090565b5f52600560205260405f2090565b35611120816103e7565b1561261d5750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b91908110156125dc5760e0020190565b5f525f516020615c3c5f395f51905f52602052600160405f20015490565b1561267257565b63031206d560e51b5f5260045ffd5b80518210156125dc5760209160051b010190565b1561269d5750565b6373a1390160e11b5f5260045260245ffd5b156126b75750565b636fc24b7960e11b5f5260045260245ffd5b6001600160a01b039091169052565b90600182811c92168015612706575b60208310146126f257565b634e487b7160e01b5f52602260045260245ffd5b91607f16916126e7565b9060405161271d81610631565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91612775826126d8565b80855291600181169081156127e157506001146127a2575b505060a0929161279e91038461064c565b0152565b5f908152602081209092505b8183106127c557505081016020018161279e61278d565b60209193508060019154838589010152019101909184926127ae565b60ff191660208681019190915292151560051b8501909201925083915061279e905061278d565b600b821015611c595752565b9060405161282181610616565b60406007829461283081612710565b845261284660ff60068301541660208601612808565b0154910152565b156128555750565b60405162461bcd60e51b815260206004820152908190612879906024830190611333565b0390fd5b15612886575050565b637a88099160e11b5f5260045260245260445ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610d2357565b604051906128c9826105f6565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906128fc826106bd565b612909604051918261064c565b828152809261291a601f19916106bd565b01905f5b82811061292a57505050565b6020906129356128bc565b8282850101520161291e565b9060018060a01b03165f5260205260405f2090565b90604051612963816105f6565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916129ae906129a5906112fa565b15156060860152565b60028101546080850152600381015460a08501520154910152565b97969594939291906129d9612ff0565b60018060a01b03811698805f5260056020526129fc8a6120138160405f206144e2565b805f526004602052612a198160ff600360405f20015416156126af565b805f52600660205260ff6001612a328c60405f20612941565b015460a81c16612a48576111209899505f612a66565b6338384f6f60e11b5f9081526001600160a01b038b16600452602490fd5b50959391949692612a75613017565b612a896001600160a01b0387161515612457565b600a548015908115612b1d575b5015612b0e57612ae3612af996612ab6612b039a612acd9685878d613636565b959094612ac161068f565b9a8b5260208b016126c9565b612ada3360408b016126c9565b606089016126c9565b608087015260a086015260c08501523691610af8565b60e08201526137d1565b60019061067e61304c565b633953b33f60e11b5f5260045ffd5b90508411155f612a96565b908160209103126103d75751611120816103e7565b6040513d5f823e3d90fd5b5f5490949392906001600160a01b03811615612c0757604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90612bb1906084870190611333565b931691015203925af18015612c025761067e925f91612bd3575b508094612ee3565b612bf5915060203d602011612bfb575b612bed818361064c565b810190612b28565b5f612bcb565b503d612be3565b612b3d565b6315aeca0d60e11b5f5260045ffd5b5f516020615d1c5f395f51905f52805460ff60401b1916600160401b179055565b5f516020615d1c5f395f51905f52805460ff60401b19169055565b15612c5957565b63ff5a231360e01b5f5260045ffd5b9060ff801983541691151516179055565b90612c85821515612c52565b6001600160a01b0316908115612d4057606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b15612d5657565b635e0c1f8960e01b5f5260045ffd5b919294939094612d73612ff0565b612d7b613017565b8584511480612e1a575b80612e10575b612d9790969596612590565b612da2343415612e24565b3683900360be1901955f5b86811015612dfb578060051b85013590888212156103d757612df5600192612dd5838a612681565b51612de08488612681565b5190612dec858a612681565b51928a01613d40565b01612dad565b5095505050505050612e0b61304c565b600190565b5081518614612d8b565b5085815114612d85565b15612e2c5750565b63943892eb60e01b5f525f60045260245260445ffd5b60405190612e4f82610631565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b15612e7957565b63f0f15b9160e01b5f5260045ffd5b15612e905750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b15612ed457565b6302bfb33d60e51b5f5260045ffd5b919091612eef33613b08565b612efa811515612ecd565b6001600160a01b03831691612f7f90612f14841515612ecd565b6001600160a01b0381169461238690612f2e871515612ecd565b612f3785615274565b612f9f575b612f4d816123348161232f896125fd565b612f58612344610680565b5f60408401525f60608401525f60808401525f60a08401525f60c0840152612381856125ef565b6040515f81525f516020615cdc5f395f51905f529080602081015b0390a4565b612fc8612faa61066f565b8681525f60208201525f60408201525f60608201526123b9876125e1565b612f3c565b6035546001600160a01b03163303612fe157565b63c82cebcb60e01b5f5260045ffd5b60ff5f516020615c5c5f395f51905f52541661300857565b63d93c066560e01b5f5260045ffd5b5f516020615c9c5f395f51905f525c61303d5760015f516020615c9c5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615c9c5f395f51905f525d565b611120916001600160a01b0316906144e2565b1561307a575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b156130a557565b630672aec160e01b5f5260045ffd5b156130bd575050565b63943892eb60e01b5f5260045260245260445ffd5b3561112081611005565b9491939291613140906131146001600160a01b0384166130ff6124f6808c61260b565b81149061310e6124f68c61260b565b91613071565b60208801946131386131286124f68861260b565b6001600160a01b0389161461309e565b848489613636565b94909360408801356131768161315f8961315a8a8a6128af565b6128af565b8110156131708a61315a8b8b6128af565b906130b4565b6131846124f66032546103db565b9061318e8361260b565b60608b01359a6131a0608082016130d2565b9060c08101359060a00135853b156103d757604051637f40ec1760e11b81526001600160a01b038a811660048301529490941660248501523060448501526064840194909452608483019c909c5260ff1660a482015260c481019190915260e481019990995288610104815a5f948591f1928315612c025761324a6132356132539361067e9b61325c97613273575b5061260b565b9161323e61068f565b998a5260208a016126c9565b604088016126c9565b606086016126c9565b608084015260a083015260c0820152612af9611b48565b806132815f6132879361064c565b80610512565b5f61322f565b90613298825f6144f5565b91826132a15750565b5f80525f516020615bbc5f395f51905f526020526001600160a01b03166132e8817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d6152dd565b156132f05750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161331383826144f5565b928361331d575050565b815f525f516020615bbc5f395f51905f5260205261334860405f209160018060a01b031680926152dd565b15613351575050565b63d180cb3160e01b5f5260045260245260445ffd5b9190916133738382614598565b928361337d575050565b815f525f516020615bbc5f395f51905f526020526133a860405f209160018060a01b0316809261535f565b156133b1575050565b62a95f1b60e31b5f5260045260245260445ffd5b156133cd5750565b63290cd55f60e01b5f52600b811015611c595760045260245ffd5b604091949392606082019560018060a01b0316825260208201520152565b906134109161436b565b60018060a01b036060820151168151905f516020615afc5f395f51905f5260208401918251946134606109216040830196875160018060a01b03169885608086019a8b519260a08801519461477b565b519251935194516040516001600160a01b0390961695918291612f9a91429190846133e8565b1561348d57565b6330d45fb160e01b5f5260045ffd5b908160209103126103d75751600b8110156103d75790565b6001600160a01b039091168152602081019190915260400190565b915061351060ff916134ff5f946134f160325460018060a01b03161515613486565b5f52600960205260405f2090565b6001600160a01b0390911690612941565b541661351b57600191565b506002905f90565b9190915f926135646114d661355461353f6124f66032546103db565b946134f16001600160a01b0387161515613486565b6001600160a01b03841690612941565b6135eb579160209161358e95935f604051809881958294633f4bdec560e01b8452600484016134b4565b03925af1928315612c02575f936135ba575b506001836135ad81611c4f565b036135b457565b60019150565b6135dd91935060203d6020116135e4575b6135d5818361064c565b81019061349c565b915f6135a0565b503d6135cb565b505050506002905f90565b908160209103126103d7575190565b908160609103126103d7578051916040602083015192015190565b1561362757565b6358e8878560e01b5f5260045ffd5b826060916136bf97959693613650610efb613554846125ef565b613660610a076040830151151590565b613761575b506136746124f66032546103db565b916136896001600160a01b0384161515613486565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612c02575f955f905f93613723575b5090829161067e94939681988515159586613718575b50508461370d575b505082613702575b5050613620565b101590505f806136fb565b101592505f806136f3565b101594505f806136eb565b905061374e91965061067e93925060603d60601161375a575b613746818361064c565b810190613605565b919692939192916136d5565b503d61373c565b60c061377391015184808210156130b4565b5f613665565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916137cc9190860190611333565b930152565b6137db8151613cad565b6137e582516125ef565b5f516020615d9c5f395f51905f526138176124f660016138106020880195610ef56124f688516103db565b01546103db565b93805190612f9a61382885516103db565b91604081019061383882516103db565b906138616080820196875160a084019485519861385b60c087019a8b51906128af565b93614b8a565b613877613870825199516103db565b93516103db565b9460e061388760608401516103db565b9751935191519201519260405197889760018060a01b03169c429689613779565b6040515f516020615abc5f395f51905f5254815f6138c5836126d8565b808352926001811690811561394157506001146138e9575b6111209250038261064c565b505f516020615abc5f395f51905f525f90815290915f516020615b7c5f395f51905f525b818310613925575050906020611120928201016138dd565b602091935080600191548385880101520191019091839261390d565b6020925061112094915060ff191682840152151560051b8201016138dd565b6040515f516020615b3c5f395f51905f5254815f61397d836126d8565b808352926001811690811561394157506001146139a0576111209250038261064c565b505f516020615b3c5f395f51905f525f90815290915f516020615d7c5f395f51905f525b8183106139dc575050906020611120928201016138dd565b60209193508060019154838588010152019101909183926139c4565b5f516020615cbc5f395f51905f525f525f516020615c3c5f395f51905f5260205260ff613a45827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c612941565b541615613a4f5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615cbc5f395f51905f52602452604490fd5b5f516020615cfc5f395f51905f525f525f516020615c3c5f395f51905f5260205260ff613acd827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143612941565b541615613ad75750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615cfc5f395f51905f52602452604490fd5b5f516020615c7c5f395f51905f525f525f516020615c3c5f395f51905f5260205260ff613b55827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f612941565b541615613b5f5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615c7c5f395f51905f52602452604490fd5b5f516020615bfc5f395f51905f525f525f516020615c3c5f395f51905f5260205260ff613bdd827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456612941565b541615613be75750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615bfc5f395f51905f52602452604490fd5b90815f525f516020615c3c5f395f51905f5260205260ff613c3c8260405f20612941565b541615613c47575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815261067e91613ca860848361064c565b614ccb565b5f526004602052600160405f2001600181540180915590565b15613ccf575050565b63a6ab65c560e01b5f5260045260245260445ffd5b903590601e19813603018212156103d757018035906001600160401b0382116103d7576020019181360383136103d757565b61067e93606092969593608083019760018060a01b03168352602083015260408201520190611c5e565b919282359182613d4f816125e1565b6003015460ff161590613d61916126af565b613d6a836125fd565b906040850191613d798361260b565b613d82906103db565b613d8b9161305e565b613d948361260b565b613d9d906103db565b613da691612615565b613dc3845f526004602052600260405f2001600181540180915590565b946020810135958681811491613dd892613cc6565b613de18361260b565b613dea906103db565b916060820194613df98661260b565b988860808501359a60a08601968a8d613e128a8a613ce4565b3690613e1d92610af8565b8051602091820120604080517fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219149381019384529081019490945260608401969096526001600160a01b0393841660808401529390921660a082015260c081019190915260e08082019390935291825290613e996101008261064c565b51902092613ea693614d53565b5f9087613eb28561260b565b90613ebd9188613523565b91909384613eca81611c4f565b600114613f85575b50613edc84611c4f565b60018403613f26575050505090613f09613f035f516020615afc5f395f51905f529361260b565b9161260b565b6040516001600160a01b03909216958291612f9a914291846133e8565b83949850613f61612f9a93613f6d937f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896613f67946150b0565b9861260b565b9261260b565b9260405193849360018060a01b031698429185613d16565b909350613fbe91925088613f988661260b565b91613fb6613faf613fa88a61260b565b9288613ce4565b3691610af8565b928a8a61477b565b9190925f613ed2565b91613fd06151cf565b613fe46001600160a01b0384161515612457565b6001600160a01b03821692613ffa841515612457565b60ff821615614278576140669061400f6151cf565b6140176151cf565b61401f6151cf565b60ff195f516020615c5c5f395f51905f5254165f516020615c5c5f395f51905f52556140496151cf565b6140516151cf565b6140596151cf565b6140616151cf565b61328d565b5061406f6151cf565b61407d60ff82161515612e72565b6040805161408b828261064c565b60098152682b30b634b230ba37b960b91b60208201526140ad8251928361064c565b60058252640312e302e360dc1b60208301526140c76151cf565b6140cf6151cf565b8051906001600160401b03821161061157614100826140fb5f516020615abc5f395f51905f52546126d8565b614ed5565b602090601f83116001146141e4579261412f836141bf97969461414394614195975f926141d9575b50506142c7565b5f516020615abc5f395f51905f52556159a6565b6141585f5f516020615b5c5f395f51905f5255565b61416d5f5f516020615d3c5f395f51905f5255565b60ff1660ff195f516020615adc5f395f51905f525416175f516020615adc5f395f51905f5255565b61419d6151fa565b603380546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615bdc5f395f51905f525f80a261067e43603455565b015190505f80614128565b5f516020615abc5f395f51905f525f52601f19831691905f516020615b7c5f395f51905f52925f5b818110614260575093614143936141959693600193836141bf9b9a9810614248575b505050811b015f516020615abc5f395f51905f52556159a6565b01515f1960f88460031b161c191690555f808061422e565b9293602060018192878601518155019501930161420c565b6338854f1160e21b5f5260045ffd5b91908203918211610d2357565b916142ad9183549060031b91821b915f19901b19161790565b9055565b8181106142bc575050565b5f81556001016142b1565b8160011b915f199060031b1c19161790565b60075f918281558260018201558260028201558260038201558260048201556005810161430681546126d8565b9081614319575b50508260068201550155565b601f821160011461433057849055505b5f8061430d565b614356614366926001601f614348855f5260205f2090565b920160051c820191016142b1565b5f81815260208120918190559055565b614329565b9190614375612e42565b50825f52600760205261438b8160405f2061535f565b156143fe576143f961067e91845f52600860205260405f20815f52602052610a4d6143b860405f20612710565b956143d56143c5826125ef565b610ef56124f660408b01516103db565b6143e9600260808a01519201918254614287565b90555f52600860205260405f2090565b6142d9565b637f11bea960e01b5f5260045260245ffd5b6003606061067e93805184556020810151600185015560408101516002850155015115159101612c68565b156144435750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c060049161447c60018060a01b038251168561246d565b60208101516144c790600186019061449d906001600160a01b03168261246d565b6040830151815460ff60a01b191690151560a01b60ff60a01b161781556060830151151590612eb0565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b805f525f516020615c3c5f395f51905f5260205260ff6145188360405f20612941565b541661457f57805f525f516020615c3c5f395f51905f5260205261453f8260405f20612941565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b611120916001600160a01b0316906152dd565b805f525f516020615c3c5f395f51905f5260205260ff6145bb8360405f20612941565b54161561457f57805f525f516020615c3c5f395f51905f526020526145e38260405f20612941565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b3d1561464a573d9061463182610add565b9161463f604051938461064c565b82523d5f602084013e565b606090565b908160209103126103d7575161112081611470565b1561466b57565b632b60b36f60e21b5f5260045ffd5b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a0820181905261112092910190611333565b91906040838203126103d7578251602084015190936001600160401b0382116103d7570181601f820112156103d7578051906146f482610add565b92614702604051948561064c565b828452602083830101116103d757815f9260208093018386015e8301015290565b611120939260809263ffffffff60e01b1682526001602083015260408201528160608201520190611333565b608090611120939263ffffffff60e01b1681525f60208201525f60408201528160608201520190611333565b9395949286926147aa9461478e876125ef565b6001600160a01b038316968792916147b5916001918590612941565b015460a01c60ff1690565b946147c46124f66035546103db565b9160188151101580614a8c575b80614a82575b614816575b50505050906147eb9391615534565b916147f583611c4f565b600183146148035750509190565b61480f939492506154eb565b6001905f90565b602081015160601c928a6034830151925f5f80604051602081019062483e5560e91b82526148598161484b8d60248301611010565b03601f19810183528261064c565b5190875afa614866614620565b9080614a76575b614a6a575b5061487f575050506147dc565b9087858a93899d60018e999a9b14159e8f998a61499b575b50505f9593919261484b8796946148d0938c89146149945788975b6040519687956020870199631eeed51360e01b8b526024880161467a565b5191855af1906148de614620565b93614984575b5080614978575b61492b5791895f516020615c1c5f395f51905f526147eb99989795938e97956149196040519283928361474f565b0390a484156155f357831594506155f3565b908997505f516020615c1c5f395f51905f5293965061497092955061480f9a9b99945080602080614961935183010191016146b9565b60409391935193849384614723565b0390a46154eb565b506040825110156148eb565b61498e9087615402565b5f6148e4565b80976148b2565b156149cd575b505f9593919261484b879694836149bc6148d0958c886154c2565b935093959750509294508f8e614897565b92509097505f9450602093506149fb92506040519485809481936340c10f1960e01b835230600484016134b4565b03925af1938415612c0257888f898f91898f995f97966148d093614a2b61484b938b9a8b91614a3b575b50614664565b935093959750509294508f6149a1565b614a5d915060203d602011614a63575b614a55818361064c565b81019061464f565b5f614a25565b503d614a4b565b6020915001515f614872565b5060208151101561486d565b50823b15156147d7565b506001600160a01b03831615156147d1565b90813b15614b1f575f516020615b9c5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115614b0757614b0491615666565b50565b505034614b1057565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b15614b4a57505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b15614b7b57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b0385169260018403614bf8575061067e9450614bc0614bb282866128af565b3414349061317084886128af565b80614bcc575b50615760565b614bed614bf291614bde6033546103db565b90614be7611b48565b91615701565b614b74565b5f614bc6565b90614c04343415612e24565b614c19614c1182876128af565b308489613c64565b80614cad575b50614c35610a0760016147aa86612381876125ef565b614c45575b5061067e9350615760565b604051632770a7eb60e21b815260208180614c648830600484016134b4565b03815f885af1918215612c025761067e96614c889387935f91614c8e575b50614b40565b5f614c3a565b614ca7915060203d602011614a6357614a55818361064c565b5f614c82565b614cc590614cbf6124f66033546103db565b87615683565b5f614c1f565b905f602091828151910182855af115612b3d575f513d614d1a57506001600160a01b0381163b155b614cfa5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415614cf3565b15614d2a57565b63b3f07a3960e01b5f5260045ffd5b15614d415750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480614eb1575b614d6c90614d23565b614d8d614d875f516020615adc5f395f51905f525460ff1690565b60ff1690565b95614d9b8488811015614d39565b5f945f935f5b868110614dbc575050505050505061067e9192811015614d39565b614e19614de883614dcb61520b565b6042916040519161190160f01b8352600283015260228201522090565b614dfc614df58489612681565b5160ff1690565b614e068487612681565b5190614e128589612681565b51926157a9565b6001600160a01b038181169088161080614e4a575b614e3c575b50600101614da1565b600198890198909650614e33565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615c3c5f395f51905f52602052614eac6114d6827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa5612941565b614e2e565b5085518314614d63565b15614ec35750565b6307a5c91d60e31b5f5260045260245ffd5b90601f8211614ee2575050565b61067e915f516020615abc5f395f51905f525f5260205f20906020601f840160051c83019310614f1a575b601f0160051c01906142b1565b9091508190614f0d565b9190601f8111614f3357505050565b61067e925f5260205f20906020601f840160051c83019310614f1a57601f0160051c01906142b1565b90600b811015611c595760ff80198354169116179055565b815180518255602081015160018301556040810151919291614fa2906001600160a01b03166002850161246d565b6060810151614fbd906001600160a01b03166003850161246d565b6080810151600484015560a00151805160058401916001600160401b03821161061157614ff482614fee85546126d8565b85614f24565b602090601f83116001146150405782600795936040959361501c935f926141d95750506142c7565b90555b615039602082015161503081611c4f565b60068601614f5c565b0151910155565b90601f19831691615054855f5260205f2090565b925f5b8181106150985750926001928592600798966040989610615080575b505050811b01905561501f565b01515f1960f88460031b161c191690555f8080615073565b92936020600181928786015181550195019301615057565b939290936150dd8135926150e26150cf855f52600760205260405f2090565b6020850135938480926152dd565b614ebb565b80156151c357935b849660408401916150fa8361260b565b946060016151079061260b565b61510f611b48565b9061511861069f565b968888528660208901526040880190615130916126c9565b61513d90606088016126c9565b87608087015260a08601525f14613f676142ad966151ab6002976151a66151b098610ef5976124f6976151b95761518f615179600154426128af565b915b6151836106ae565b95865260208601612808565b6040840152610a4d855f52600860205260405f2090565b614f74565b6125ef565b019182546128af565b61518f5f9161517b565b506080820135936150ea565b60ff5f516020615d1c5f395f51905f525460401c16156151eb57565b631afcd79f60e31b5f5260045ffd5b6152026151cf565b62015180600155565b6152136157c1565b61521b615818565b6040519060208201925f516020615d5c5f395f51905f528452604083015260608201524660808201523060a082015260a0815261525960c08261064c565b51902090565b80548210156125dc575f5260205f2001905f90565b5f818152600360205260409020546152d857600254600160401b811015610611576152c16152ab826001859401600255600261525f565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6152e782826144e2565b61457f57805490600160401b821015610611578261530f6152ab84600180960185558461525f565b90558054925f520160205260405f2055600190565b8054801561534b575f19019061533a828261525f565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146153fa575f198401848111610d235783545f19810194908511610d23575f9585836153ad97610a4d95036153b3575b505050615324565b55600190565b6153e36153dd916153d46153ca6153f1958861525f565b90549060031b1c90565b9283918761525f565b90614294565b85905f5260205260405f2090565b555f80806153a5565b505050505f90565b6040519060205f81840163095ea7b360e01b8152615436856154288489602484016134b4565b03601f19810187528661064c565b84519082855af15f51903d81615496575b501590505b61545557505050565b60405163095ea7b360e01b60208201526001600160a01b0390931660248401525f604484015261067e92613ca890615490816064810161484b565b82614ccb565b151590506154b6575061544c6001600160a01b0382163b15155b5f615447565b600161544c91146154b0565b91909160205f60405193615436856154288582019363095ea7b360e01b855289602484016134b4565b90615500915f52600660205260405f20612941565b600181015460a01c60ff1615615522576003018054918203918211610d235755565b6004018054918201809211610d235755565b6001600160a01b0316929190600184146155c35781156155ba57615583921561558e5760405163a9059cbb60e01b60208201529161557b91839161484b91602484016134b4565b6005926156b8565b156111205750600190565b6040516340c10f1960e01b6020820152916155b291839161484b91602484016134b4565b6006926156b8565b50505050600190565b906155e59350610a0792506155d6611b48565b916001600160a01b0316615701565b6155ee57600190565b600590565b6001600160a01b0316936001851461565257821561564857615583938115615640575b501561558e5760405163a9059cbb60e01b60208201529161557b91839161484b91602484016134b4565b90505f615616565b5050505050600190565b50906155e59350610a0792506155d6611b48565b5f8061112093602081519101845af461567d614620565b9161584a565b613ca861067e93926156aa60405194859263a9059cbb60e01b6020850152602484016134b4565b03601f19810184528361064c565b81516001600160a01b03909116915f9182916020018285620186a0f16156dc614620565b901561457f5780516156f857503b156156f457600190565b5f90565b60209150015190565b5f809161571084471015614664565b84516001600160a01b0391909116946020018486620186a0f190615732614620565b91156157595715615745575b5050600190565b80516156f857503b156156f4575f8061573e565b5050505f90565b90615775915f52600660205260405f20612941565b600181015460a01c60ff1615615797576003018054918201809211610d235755565b6004018054918203918211610d235755565b9161112093916157b8936158a8565b9092919261592a565b6157c96138a8565b80519081156157d9576020012090565b50505f516020615b5c5f395f51905f525480156157f35790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615820613960565b8051908115615830576020012090565b50505f516020615d3c5f395f51905f525480156157f35790565b9061586e575080511561585f57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061589f575b61587f575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615877565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615915579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612c02575f516001600160a01b0381161561590b57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611c5957565b61593381615920565b8061593c575050565b61594581615920565b6001810361595c5763f645eedf60e01b5f5260045ffd5b61596581615920565b60028103615980575063fce698f760e01b5f5260045260245ffd5b8061598c600392615920565b146159945750565b6335e2f38360e21b5f5260045260245ffd5b80519091906001600160401b038111610611576159e7816159d45f516020615b3c5f395f51905f52546126d8565b5f516020615b3c5f395f51905f52614f24565b602092601f8211600114615a1a57615a09929382915f926141d95750506142c7565b5f516020615b3c5f395f51905f5255565b5f516020615b3c5f395f51905f525f52601f198216935f516020615d7c5f395f51905f52915f5b868110615a835750836001959610615a6b575b505050811b015f516020615b3c5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615a54565b91926020600181928685015181550194019201615a4156feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a264697066735822122096da1c26f70818eb11ba183fbfb981e4cae47c8e411664a706175d54e551a6fe64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeV2MetaData.ABI instead.
var BSCBridgeV2ABI = BSCBridgeV2MetaData.ABI

// Deprecated: Use BSCBridgeV2MetaData.Sigs instead.
// BSCBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var BSCBridgeV2FuncSigs = BSCBridgeV2MetaData.Sigs

// BSCBridgeV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BSCBridgeV2MetaData.Bin instead.
var BSCBridgeV2Bin = BSCBridgeV2MetaData.Bin

// DeployBSCBridgeV2 deploys a new Ethereum contract, binding an instance of BSCBridgeV2 to it.
func DeployBSCBridgeV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCBridgeV2, error) {
	parsed, err := BSCBridgeV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BSCBridgeV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCBridgeV2{BSCBridgeV2Caller: BSCBridgeV2Caller{contract: contract}, BSCBridgeV2Transactor: BSCBridgeV2Transactor{contract: contract}, BSCBridgeV2Filterer: BSCBridgeV2Filterer{contract: contract}}, nil
}

// BSCBridgeV2 is an auto generated Go binding around an Ethereum contract.
type BSCBridgeV2 struct {
	BSCBridgeV2Caller     // Read-only binding to the contract
	BSCBridgeV2Transactor // Write-only binding to the contract
	BSCBridgeV2Filterer   // Log filterer for contract events
}

// BSCBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BSCBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCBridgeV2Session struct {
	Contract     *BSCBridgeV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCBridgeV2CallerSession struct {
	Contract *BSCBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BSCBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCBridgeV2TransactorSession struct {
	Contract     *BSCBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BSCBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BSCBridgeV2Raw struct {
	Contract *BSCBridgeV2 // Generic contract binding to access the raw methods on
}

// BSCBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCBridgeV2CallerRaw struct {
	Contract *BSCBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// BSCBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCBridgeV2TransactorRaw struct {
	Contract *BSCBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCBridgeV2 creates a new instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2(address common.Address, backend bind.ContractBackend) (*BSCBridgeV2, error) {
	contract, err := bindBSCBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2{BSCBridgeV2Caller: BSCBridgeV2Caller{contract: contract}, BSCBridgeV2Transactor: BSCBridgeV2Transactor{contract: contract}, BSCBridgeV2Filterer: BSCBridgeV2Filterer{contract: contract}}, nil
}

// NewBSCBridgeV2Caller creates a new read-only instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*BSCBridgeV2Caller, error) {
	contract, err := bindBSCBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Caller{contract: contract}, nil
}

// NewBSCBridgeV2Transactor creates a new write-only instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BSCBridgeV2Transactor, error) {
	contract, err := bindBSCBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Transactor{contract: contract}, nil
}

// NewBSCBridgeV2Filterer creates a new log filterer instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BSCBridgeV2Filterer, error) {
	contract, err := bindBSCBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Filterer{contract: contract}, nil
}

// bindBSCBridgeV2 binds a generic wrapper to an already deployed contract.
func bindBSCBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BSCBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeV2.Contract.BSCBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BSCBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BSCBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeV2 *BSCBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeV2 *BSCBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeV2 *BSCBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DEFAULTADMINROLE(&_BSCBridgeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DEFAULTADMINROLE(&_BSCBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_BSCBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_BSCBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllChainIDs(&_BSCBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllChainIDs(&_BSCBridgeV2.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllPendingIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllPendingIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.AllTokenPairs(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.AllTokenPairs(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeExecutor() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecutor(&_BSCBridgeV2.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) BridgeExecutor() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecutor(&_BSCBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeVerifier() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeVerifier(&_BSCBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) BridgeVerifier() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeVerifier(&_BSCBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridgeV2.Contract.CrossMintableERC20Code(&_BSCBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridgeV2.Contract.CrossMintableERC20Code(&_BSCBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) Dev() (common.Address, error) {
	return _BSCBridgeV2.Contract.Dev(&_BSCBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Dev() (common.Address, error) {
	return _BSCBridgeV2.Contract.Dev(&_BSCBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) DomainSeparator() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DomainSeparator(&_BSCBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) DomainSeparator() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DomainSeparator(&_BSCBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "eip712Domain")

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
func (_BSCBridgeV2 *BSCBridgeV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridgeV2.Contract.Eip712Domain(&_BSCBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridgeV2.Contract.Eip712Domain(&_BSCBridgeV2.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextFinalizeIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextFinalizeIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextInitiateIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextInitiateIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Session) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridgeV2.Contract.GetPendingArguments(&_BSCBridgeV2.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridgeV2.Contract.GetPendingArguments(&_BSCBridgeV2.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridgeV2.Contract.GetRoleAdmin(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridgeV2.Contract.GetRoleAdmin(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2Session) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridgeV2.Contract.GetRoleMembers(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridgeV2.Contract.GetRoleMembers(&_BSCBridgeV2.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Session) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.GetTokenPair(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.GetTokenPair(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.HasRole(&_BSCBridgeV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.HasRole(&_BSCBridgeV2.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) InitializedAt() (*big.Int, error) {
	return _BSCBridgeV2.Contract.InitializedAt(&_BSCBridgeV2.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) InitializedAt() (*big.Int, error) {
	return _BSCBridgeV2.Contract.InitializedAt(&_BSCBridgeV2.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.IsTokenFinalizePaused(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.IsTokenFinalizePaused(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) MaxExtraDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "maxExtraDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) MaxExtraDataLength() (*big.Int, error) {
	return _BSCBridgeV2.Contract.MaxExtraDataLength(&_BSCBridgeV2.CallOpts)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) MaxExtraDataLength() (*big.Int, error) {
	return _BSCBridgeV2.Contract.MaxExtraDataLength(&_BSCBridgeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) Paused() (bool, error) {
	return _BSCBridgeV2.Contract.Paused(&_BSCBridgeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Paused() (bool, error) {
	return _BSCBridgeV2.Contract.Paused(&_BSCBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) ProxiableUUID() ([32]byte, error) {
	return _BSCBridgeV2.Contract.ProxiableUUID(&_BSCBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _BSCBridgeV2.Contract.ProxiableUUID(&_BSCBridgeV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridgeV2.Contract.SupportsInterface(&_BSCBridgeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridgeV2.Contract.SupportsInterface(&_BSCBridgeV2.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2Session) Threshold() (uint8, error) {
	return _BSCBridgeV2.Contract.Threshold(&_BSCBridgeV2.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Threshold() (uint8, error) {
	return _BSCBridgeV2.Contract.Threshold(&_BSCBridgeV2.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) BurnCrossToDeadWallet(opts *bind.TransactOpts, deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "burnCrossToDeadWallet", deadWallet, amount, alreadyTransferred)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) BurnCrossToDeadWallet(deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BurnCrossToDeadWallet(&_BSCBridgeV2.TransactOpts, deadWallet, amount, alreadyTransferred)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) BurnCrossToDeadWallet(deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BurnCrossToDeadWallet(&_BSCBridgeV2.TransactOpts, deadWallet, amount, alreadyTransferred)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ChangeThreshold(&_BSCBridgeV2.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ChangeThreshold(&_BSCBridgeV2.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2Session) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.CreateToken(&_BSCBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.CreateToken(&_BSCBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridgeBatch(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridgeBatch(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) Initialize(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "initialize", owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Initialize(&_BSCBridgeV2.TransactOpts, owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Initialize(&_BSCBridgeV2.TransactOpts, owner, dev_, threshold_)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePendingWithRecipient(&_BSCBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePendingWithRecipient(&_BSCBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeTokenBatch(&_BSCBridgeV2.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeTokenBatch(&_BSCBridgeV2.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RegisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RegisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ReinitializeBSCBridge(opts *bind.TransactOpts, crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "reinitializeBSCBridge", crossChainID_, cross_)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ReinitializeBSCBridge(crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReinitializeBSCBridge(&_BSCBridgeV2.TransactOpts, crossChainID_, cross_)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ReinitializeBSCBridge(crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReinitializeBSCBridge(&_BSCBridgeV2.TransactOpts, crossChainID_, cross_)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RemovePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RemovePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RenounceRole(&_BSCBridgeV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RenounceRole(&_BSCBridgeV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecutor(&_BSCBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecutor(&_BSCBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeVerifier(&_BSCBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeVerifier(&_BSCBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetChainPause(&_BSCBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetChainPause(&_BSCBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetCrossMintableERC20Code(&_BSCBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetCrossMintableERC20Code(&_BSCBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetDev(&_BSCBridgeV2.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetDev(&_BSCBridgeV2.TransactOpts, dev_)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetMaxExtraDataLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setMaxExtraDataLength", length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetMaxExtraDataLength(&_BSCBridgeV2.TransactOpts, length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetMaxExtraDataLength(&_BSCBridgeV2.TransactOpts, length)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPause(&_BSCBridgeV2.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPause(&_BSCBridgeV2.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelay(&_BSCBridgeV2.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelay(&_BSCBridgeV2.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelayExpiration(&_BSCBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelayExpiration(&_BSCBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UpgradeToAndCall(&_BSCBridgeV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UpgradeToAndCall(&_BSCBridgeV2.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) Receive() (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Receive(&_BSCBridgeV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) Receive() (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Receive(&_BSCBridgeV2.TransactOpts)
}

// BSCBridgeV2BridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecutorSetIterator struct {
	Event *BSCBridgeV2BridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeExecutorSet)
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
		it.Event = new(BSCBridgeV2BridgeExecutorSet)
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
func (it *BSCBridgeV2BridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeExecutorSet represents a BridgeExecutorSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*BSCBridgeV2BridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeExecutorSetIterator{contract: _BSCBridgeV2.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeExecutorSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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

// ParseBridgeExecutorSet is a log parse operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeExecutorSet(log types.Log) (*BSCBridgeV2BridgeExecutorSet, error) {
	event := new(BSCBridgeV2BridgeExecutorSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeFinalizedIterator struct {
	Event *BSCBridgeV2BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeFinalized)
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
		it.Event = new(BSCBridgeV2BridgeFinalized)
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
func (it *BSCBridgeV2BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeFinalized represents a BridgeFinalized event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeFinalized struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeV2BridgeFinalizedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeFinalizedIterator{contract: _BSCBridgeV2.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeFinalized)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeFinalized(log types.Log) (*BSCBridgeV2BridgeFinalized, error) {
	event := new(BSCBridgeV2BridgeFinalized)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeInitiatedIterator struct {
	Event *BSCBridgeV2BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeInitiated)
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
		it.Event = new(BSCBridgeV2BridgeInitiated)
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
func (it *BSCBridgeV2BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeInitiated represents a BridgeInitiated event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeInitiated struct {
	ToChainID  *big.Int
	Index      *big.Int
	FromToken  common.Address
	ToToken    common.Address
	From       common.Address
	To         common.Address
	Value      *big.Int
	NetworkFee *big.Int
	ExFee      *big.Int
	ExtraData  []byte
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*BSCBridgeV2BridgeInitiatedIterator, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeInitiatedIterator{contract: _BSCBridgeV2.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeInitiated)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeInitiated(log types.Log) (*BSCBridgeV2BridgeInitiated, error) {
	event := new(BSCBridgeV2BridgeInitiated)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgePendingIterator struct {
	Event *BSCBridgeV2BridgePending // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgePending)
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
		it.Event = new(BSCBridgeV2BridgePending)
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
func (it *BSCBridgeV2BridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgePending represents a BridgePending event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgePending struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgePending is a free log retrieval operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeV2BridgePendingIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgePendingIterator{contract: _BSCBridgeV2.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgePending)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
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

// ParseBridgePending is a log parse operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgePending(log types.Log) (*BSCBridgeV2BridgePending, error) {
	event := new(BSCBridgeV2BridgePending)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeVerifierSetIterator struct {
	Event *BSCBridgeV2BridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeVerifierSet)
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
		it.Event = new(BSCBridgeV2BridgeVerifierSet)
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
func (it *BSCBridgeV2BridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeVerifierSet represents a BridgeVerifierSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*BSCBridgeV2BridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeVerifierSetIterator{contract: _BSCBridgeV2.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeVerifierSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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

// ParseBridgeVerifierSet is a log parse operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeVerifierSet(log types.Log) (*BSCBridgeV2BridgeVerifierSet, error) {
	event := new(BSCBridgeV2BridgeVerifierSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ChainPauseSetIterator struct {
	Event *BSCBridgeV2ChainPauseSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2ChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ChainPauseSet)
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
		it.Event = new(BSCBridgeV2ChainPauseSet)
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
func (it *BSCBridgeV2ChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ChainPauseSet represents a ChainPauseSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BSCBridgeV2ChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ChainPauseSetIterator{contract: _BSCBridgeV2.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ChainPauseSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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

// ParseChainPauseSet is a log parse operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseChainPauseSet(log types.Log) (*BSCBridgeV2ChainPauseSet, error) {
	event := new(BSCBridgeV2ChainPauseSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2CrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2CrossMintableERC20CodeSetIterator struct {
	Event *BSCBridgeV2CrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2CrossMintableERC20CodeSet)
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
		it.Event = new(BSCBridgeV2CrossMintableERC20CodeSet)
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
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2CrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2CrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*BSCBridgeV2CrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2CrossMintableERC20CodeSetIterator{contract: _BSCBridgeV2.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2CrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2CrossMintableERC20CodeSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseCrossMintableERC20CodeSet(log types.Log) (*BSCBridgeV2CrossMintableERC20CodeSet, error) {
	event := new(BSCBridgeV2CrossMintableERC20CodeSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2DevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2DevSetIterator struct {
	Event *BSCBridgeV2DevSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2DevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2DevSet)
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
		it.Event = new(BSCBridgeV2DevSet)
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
func (it *BSCBridgeV2DevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2DevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2DevSet represents a DevSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2DevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*BSCBridgeV2DevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2DevSetIterator{contract: _BSCBridgeV2.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2DevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2DevSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
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

// ParseDevSet is a log parse operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseDevSet(log types.Log) (*BSCBridgeV2DevSet, error) {
	event := new(BSCBridgeV2DevSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2EIP712DomainChangedIterator struct {
	Event *BSCBridgeV2EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2EIP712DomainChanged)
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
		it.Event = new(BSCBridgeV2EIP712DomainChanged)
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
func (it *BSCBridgeV2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BSCBridgeV2EIP712DomainChangedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2EIP712DomainChangedIterator{contract: _BSCBridgeV2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2EIP712DomainChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseEIP712DomainChanged(log types.Log) (*BSCBridgeV2EIP712DomainChanged, error) {
	event := new(BSCBridgeV2EIP712DomainChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ExtraCallExecutedIterator struct {
	Event *BSCBridgeV2ExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2ExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ExtraCallExecuted)
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
		it.Event = new(BSCBridgeV2ExtraCallExecuted)
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
func (it *BSCBridgeV2ExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ExtraCallExecuted represents a ExtraCallExecuted event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ExtraCallExecuted struct {
	FromChainID    *big.Int
	Index          *big.Int
	TargetContract common.Address
	MethodID       [4]byte
	Success        bool
	Consumed       *big.Int
	ReturnData     []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (*BSCBridgeV2ExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ExtraCallExecutedIterator{contract: _BSCBridgeV2.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ExtraCallExecuted)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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

// ParseExtraCallExecuted is a log parse operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseExtraCallExecuted(log types.Log) (*BSCBridgeV2ExtraCallExecuted, error) {
	event := new(BSCBridgeV2ExtraCallExecuted)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BSCBridgeV2 contract.
type BSCBridgeV2InitializedIterator struct {
	Event *BSCBridgeV2Initialized // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Initialized)
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
		it.Event = new(BSCBridgeV2Initialized)
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
func (it *BSCBridgeV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Initialized represents a Initialized event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*BSCBridgeV2InitializedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2InitializedIterator{contract: _BSCBridgeV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Initialized) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Initialized)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseInitialized(log types.Log) (*BSCBridgeV2Initialized, error) {
	event := new(BSCBridgeV2Initialized)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ManualReleasedIterator struct {
	Event *BSCBridgeV2ManualReleased // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2ManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ManualReleased)
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
		it.Event = new(BSCBridgeV2ManualReleased)
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
func (it *BSCBridgeV2ManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ManualReleased represents a ManualReleased event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgeV2ManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ManualReleasedIterator{contract: _BSCBridgeV2.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ManualReleased)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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

// ParseManualReleased is a log parse operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseManualReleased(log types.Log) (*BSCBridgeV2ManualReleased, error) {
	event := new(BSCBridgeV2ManualReleased)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2MaxExtraDataLengthSetIterator is returned from FilterMaxExtraDataLengthSet and is used to iterate over the raw logs and unpacked data for MaxExtraDataLengthSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2MaxExtraDataLengthSetIterator struct {
	Event *BSCBridgeV2MaxExtraDataLengthSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2MaxExtraDataLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2MaxExtraDataLengthSet)
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
		it.Event = new(BSCBridgeV2MaxExtraDataLengthSet)
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
func (it *BSCBridgeV2MaxExtraDataLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2MaxExtraDataLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2MaxExtraDataLengthSet represents a MaxExtraDataLengthSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2MaxExtraDataLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxExtraDataLengthSet is a free log retrieval operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterMaxExtraDataLengthSet(opts *bind.FilterOpts) (*BSCBridgeV2MaxExtraDataLengthSetIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2MaxExtraDataLengthSetIterator{contract: _BSCBridgeV2.contract, event: "MaxExtraDataLengthSet", logs: logs, sub: sub}, nil
}

// WatchMaxExtraDataLengthSet is a free log subscription operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchMaxExtraDataLengthSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2MaxExtraDataLengthSet) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2MaxExtraDataLengthSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
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

// ParseMaxExtraDataLengthSet is a log parse operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseMaxExtraDataLengthSet(log types.Log) (*BSCBridgeV2MaxExtraDataLengthSet, error) {
	event := new(BSCBridgeV2MaxExtraDataLengthSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BSCBridgeV2 contract.
type BSCBridgeV2PausedIterator struct {
	Event *BSCBridgeV2Paused // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Paused)
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
		it.Event = new(BSCBridgeV2Paused)
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
func (it *BSCBridgeV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Paused represents a Paused event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterPaused(opts *bind.FilterOpts) (*BSCBridgeV2PausedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2PausedIterator{contract: _BSCBridgeV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Paused) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Paused)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParsePaused(log types.Log) (*BSCBridgeV2Paused, error) {
	event := new(BSCBridgeV2Paused)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2PendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the BSCBridgeV2 contract.
type BSCBridgeV2PendingRemovedIterator struct {
	Event *BSCBridgeV2PendingRemoved // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2PendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2PendingRemoved)
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
		it.Event = new(BSCBridgeV2PendingRemoved)
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
func (it *BSCBridgeV2PendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2PendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2PendingRemoved represents a PendingRemoved event raised by the BSCBridgeV2 contract.
type BSCBridgeV2PendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgeV2PendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2PendingRemovedIterator{contract: _BSCBridgeV2.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2PendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2PendingRemoved)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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

// ParsePendingRemoved is a log parse operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParsePendingRemoved(log types.Log) (*BSCBridgeV2PendingRemoved, error) {
	event := new(BSCBridgeV2PendingRemoved)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleAdminChangedIterator struct {
	Event *BSCBridgeV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleAdminChanged)
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
		it.Event = new(BSCBridgeV2RoleAdminChanged)
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
func (it *BSCBridgeV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleAdminChanged represents a RoleAdminChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BSCBridgeV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleAdminChangedIterator{contract: _BSCBridgeV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleAdminChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleAdminChanged(log types.Log) (*BSCBridgeV2RoleAdminChanged, error) {
	event := new(BSCBridgeV2RoleAdminChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleGrantedIterator struct {
	Event *BSCBridgeV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleGranted)
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
		it.Event = new(BSCBridgeV2RoleGranted)
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
func (it *BSCBridgeV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleGranted represents a RoleGranted event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleGrantedIterator{contract: _BSCBridgeV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleGranted)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleGranted(log types.Log) (*BSCBridgeV2RoleGranted, error) {
	event := new(BSCBridgeV2RoleGranted)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleRevokedIterator struct {
	Event *BSCBridgeV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleRevoked)
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
		it.Event = new(BSCBridgeV2RoleRevoked)
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
func (it *BSCBridgeV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleRevoked represents a RoleRevoked event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleRevokedIterator{contract: _BSCBridgeV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleRevoked)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleRevoked(log types.Log) (*BSCBridgeV2RoleRevoked, error) {
	event := new(BSCBridgeV2RoleRevoked)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ThresholdChangedIterator struct {
	Event *BSCBridgeV2ThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2ThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ThresholdChanged)
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
		it.Event = new(BSCBridgeV2ThresholdChanged)
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
func (it *BSCBridgeV2ThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ThresholdChanged represents a ThresholdChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BSCBridgeV2ThresholdChangedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ThresholdChangedIterator{contract: _BSCBridgeV2.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ThresholdChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseThresholdChanged(log types.Log) (*BSCBridgeV2ThresholdChanged, error) {
	event := new(BSCBridgeV2ThresholdChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2TokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairRegisteredIterator struct {
	Event *BSCBridgeV2TokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2TokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2TokenPairRegistered)
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
		it.Event = new(BSCBridgeV2TokenPairRegistered)
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
func (it *BSCBridgeV2TokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2TokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2TokenPairRegistered represents a TokenPairRegistered event raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BSCBridgeV2TokenPairRegisteredIterator, error) {

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

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2TokenPairRegisteredIterator{contract: _BSCBridgeV2.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2TokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2TokenPairRegistered)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseTokenPairRegistered(log types.Log) (*BSCBridgeV2TokenPairRegistered, error) {
	event := new(BSCBridgeV2TokenPairRegistered)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2TokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPauseSetIterator struct {
	Event *BSCBridgeV2TokenPauseSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2TokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2TokenPauseSet)
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
		it.Event = new(BSCBridgeV2TokenPauseSet)
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
func (it *BSCBridgeV2TokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2TokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2TokenPauseSet represents a TokenPauseSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BSCBridgeV2TokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2TokenPauseSetIterator{contract: _BSCBridgeV2.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2TokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2TokenPauseSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseTokenPauseSet(log types.Log) (*BSCBridgeV2TokenPauseSet, error) {
	event := new(BSCBridgeV2TokenPauseSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BSCBridgeV2 contract.
type BSCBridgeV2UnpausedIterator struct {
	Event *BSCBridgeV2Unpaused // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Unpaused)
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
		it.Event = new(BSCBridgeV2Unpaused)
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
func (it *BSCBridgeV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Unpaused represents a Unpaused event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*BSCBridgeV2UnpausedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2UnpausedIterator{contract: _BSCBridgeV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Unpaused)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseUnpaused(log types.Log) (*BSCBridgeV2Unpaused, error) {
	event := new(BSCBridgeV2Unpaused)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BSCBridgeV2 contract.
type BSCBridgeV2UpgradedIterator struct {
	Event *BSCBridgeV2Upgraded // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Upgraded)
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
		it.Event = new(BSCBridgeV2Upgraded)
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
func (it *BSCBridgeV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Upgraded represents a Upgraded event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BSCBridgeV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2UpgradedIterator{contract: _BSCBridgeV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Upgraded)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseUpgraded(log types.Log) (*BSCBridgeV2Upgraded, error) {
	event := new(BSCBridgeV2Upgraded)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2VerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelayExpirationSetIterator struct {
	Event *BSCBridgeV2VerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2VerificationDelayExpirationSet)
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
		it.Event = new(BSCBridgeV2VerificationDelayExpirationSet)
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
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2VerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BSCBridgeV2VerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2VerificationDelayExpirationSetIterator{contract: _BSCBridgeV2.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2VerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2VerificationDelayExpirationSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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

// ParseVerificationDelayExpirationSet is a log parse operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseVerificationDelayExpirationSet(log types.Log) (*BSCBridgeV2VerificationDelayExpirationSet, error) {
	event := new(BSCBridgeV2VerificationDelayExpirationSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2VerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelaySetIterator struct {
	Event *BSCBridgeV2VerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2VerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2VerificationDelaySet)
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
		it.Event = new(BSCBridgeV2VerificationDelaySet)
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
func (it *BSCBridgeV2VerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2VerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2VerificationDelaySet represents a VerificationDelaySet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*BSCBridgeV2VerificationDelaySetIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2VerificationDelaySetIterator{contract: _BSCBridgeV2.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2VerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2VerificationDelaySet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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

// ParseVerificationDelaySet is a log parse operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseVerificationDelaySet(log types.Log) (*BSCBridgeV2VerificationDelaySet, error) {
	event := new(BSCBridgeV2VerificationDelaySet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
