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

// BaseBridgeMetaData contains all meta data concerning the BaseBridge contract.
var BaseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"a10bab0b": "bridgeExecutor()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
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
	Bin: "0x60a080604052346100c257306080525f516020615b1a5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615a5390816100c78239608051818181610b570152610d0e0152f35b6001600160401b0319166001600160401b039081175f516020615b1a5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b610020612c9b565b005b5f3560e01c806301ffc9a7146103615780630b1d8d061461035c5780631313996b146103575780631e24abdb14610352578063248a9ca31461034d5780632f2ff15d1461034857806336568abe14610343578063365d71e41461033e57806342cde4e81461033957806348a00ed8146103345780634be13f831461032f5780634ee078ba1461032a5780634f1ef28614610325578063502cc5c01461032057806352d1902d1461031b57806356ff54b0146103165780635b605f5c146103115780635c975abb1461030c5780635fd262de14610307578063670e626814610302578063761fe2d8146102fd57806379214874146102f8578063814914b5146102f357806384b0196e146102ee57806388d67d6d146102e957806389232a00146102e45780638ae87c5c146102df57806391cca3db146102da57806391cf6d3e146102d557806391d14854146102d0578063a10bab0b146102cb578063a217fddf146102c6578063a3246ad3146102c1578063aa1bd0bc146102bc578063aa20ed47146102b7578063ad3cb1cc146102b2578063ae6893f8146102ad578063b2b49e2e146102a8578063b33eb36e146102a3578063b7f3358d1461029e578063b8aa837e14610299578063bedb86fb14610294578063bfbf67651461028f578063cba25e4b1461028a578063cf56118e14610285578063d477f05f14610280578063d547741f1461027b578063d5717fc514610276578063e2af6cd714610271578063edbbea391461026c578063f0702e8e146102675763f698da250361000e576121f8565b6121d0565b61208b565b612013565b611fda565b611fab565b611f4d565b611ed9565b611e6f565b611d97565b611cae565b611c13565b611b86565b611af5565b6119e8565b6119af565b611968565b6118df565b611893565b611817565b6117bb565b611793565b611751565b611734565b61170c565b6116a4565b611562565b6113fb565b6112c3565b611176565b6110f6565b61106d565b610ff6565b610f48565b610f1a565b610e37565b610d53565b610cfc565b610c6d565b610b04565b610996565b61096f565b610845565b610819565b6107b6565b610577565b61053e565b610518565b6104fb565b610479565b6103d8565b346103b75760203660031901126103b75760043563ffffffff60e01b81168091036103b757602090637965db0b60e01b81149081156103a6575b506040519015158152f35b6301ffc9a760e01b1490505f61039b565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103b757565b346103b75760203660031901126103b7576004356103f5816103c7565b6103fd612cbe565b6001600160a01b0316610411811515612264565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103b7578235916001600160401b0383116103b7576020808501948460051b0101116103b757565b60403660031901126103b7576004356001600160401b0381116103b7576104a4903690600401610449565b602435916001600160401b0383116103b757366023840112156103b7576004830135916001600160401b0383116103b75736602460e08502860101116103b7576024610020940191612299565b5f9103126103b757565b346103b7575f3660031901126103b7576020600a54604051908152f35b346103b75760203660031901126103b757602061053660043561245a565b604051908152f35b346103b75760403660031901126103b757610020602435600435610561826103c7565b61057261056d8261245a565b612ea6565b6131ff565b346103b75760403660031901126103b757600435602435610597816103c7565b336001600160a01b038216036105b0576100209161325f565b63334bd91960e11b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b038211176105ee57604052565b6105bf565b606081019081106001600160401b038211176105ee57604052565b60c081019081106001600160401b038211176105ee57604052565b601f909101601f19168101906001600160401b038211908210176105ee57604052565b6040519061065b608083610629565b565b6040519061065b60e083610629565b6040519061065b61010083610629565b6040519061065b60c083610629565b6040519061065b606083610629565b6001600160401b0381116105ee5760051b60200190565b9080601f830112156103b75781356106c88161069a565b926106d66040519485610629565b81845260208085019260051b8201019283116103b757602001905b8282106106fe5750505090565b81358152602091820191016106f1565b9060406003198301126103b7576004356001600160401b0381116103b75782610739916004016106b1565b91602435906001600160401b0382116103b757806023830112156103b75781600401356107658161069a565b926107736040519485610629565b8184526024602085019260051b8201019283116103b757602401905b82821061079c5750505090565b6020809183356107ab816103c7565b81520191019061078f565b346103b7576107c43661070e565b906107d28151835114612478565b5f5b825181101561002057806108126107ed6001938561248e565b51838060a01b036107fe848861248e565b51169061080d61056d8261245a565b61325f565b50016107d4565b346103b7575f3660031901126103b757602060ff5f51602061577e5f395f51905f525416604051908152f35b346103b75760603660031901126103b757602435600435604435610868816103c7565b610870612d38565b610878612f10565b815f526007602052610896836108918160405f2061413b565b6124a2565b806108a18484613fc4565b916001600160a01b03161561095b575b8151905f51602061579e5f395f51905f52602084019161091483519561090a6108f76040830197885160018060a01b03169986608086019b8c519260a0880151946143d4565b5061090181611a5a565b600181146132be565b51935194516103bb565b94516040516001600160a01b039096169591829161093591429190846132e1565b0390a45f51602061573e5f395f51905f525f80a35f5f51602061591e5f395f51905f525d005b5060608101516001600160a01b03166108b1565b346103b7575f3660031901126103b7575f546040516001600160a01b039091168152602090f35b346103b75760403660031901126103b757610aa16024356004356109b8612ee9565b6109c0612f10565b6109e6816109e16109dd60036109d5846123ee565b015460ff1690565b1590565b6124bc565b610a058261089181610a00855f52600760205260405f2090565b61413b565b610a9c6040610a35610a3085610a23865f52600860205260405f2090565b905f5260205260405f2090565b612621565b610a89610a5582516080610a4b868301516103bb565b91015190876133c8565b50610a5f81611a5a565b610a6881611a5a565b83516020808201839052815290600190610a83604084610629565b1461265a565b01518015908115610aa9575b429161268a565b6132ff565b610020612f45565b4281109150610a95565b6001600160401b0381116105ee57601f01601f191660200190565b929192610ada82610ab3565b91610ae86040519384610629565b8294818452818301116103b7578281602093845f960137010152565b60403660031901126103b757600435610b1c816103c7565b6024356001600160401b0381116103b757366023820112156103b757610b4c903690602481600401359101610ace565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610c4b575b50610c3c57610b8f612cbe565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610c0b575b50610bd857634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f51602061581e5f395f51905f528303610bf75761002092506146f7565b632a87526960e21b5f52600483905260245ffd5b610c2e91945060203d602011610c35575b610c268183610629565b8101906134f6565b925f610bb7565b503d610c1c565b63703e46dd60e11b5f5260045ffd5b5f51602061581e5f395f51905f52546001600160a01b0316141590505f610b82565b346103b75760603660031901126103b757602435600435604435610c8f612d38565b815f526007602052610ca8836108918160405f2061413b565b4201804211610cf75760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b6126a8565b346103b7575f3660031901126103b7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610c3c5760206040515f51602061581e5f395f51905f528152f35b346103b75760203660031901126103b7577f1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f16020600435610d92612db2565b80600a55604051908152a1005b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610e175750505090565b909192602060e082610e2c6001948851610d9f565b019401929101610e0a565b346103b75760203660031901126103b757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b818110610f01575050610e8692500383610629565b610e9082516126ff565b915f5b8151811015610eef57600190610ed3610ece610eae866123fc565b610ec8610ebb858861248e565b516001600160a01b031690565b9061274e565b612763565b610edd828761248e565b52610ee8818661248e565b5001610e93565b60405180610efd8682610df4565b0390f35b8454835260019485019487945060209093019201610e71565b346103b7575f3660031901126103b757602060ff5f5160206158de5f395f51905f5254166040519015158152f35b60e03660031901126103b757602435600435610f63826103c7565b604435610f6f816103c7565b60c4359160a4356084356064356001600160401b0386116103b757366023870112156103b7576004860135946001600160401b0386116103b75736602487890101116103b757610efd976024610fc69801956127d6565b60405191829182901515815260200190565b60ff8116036103b757565b6001600160a01b03909116815260200190565b346103b75760803660031901126103b757602435600435611016826103c7565b604435906001600160401b0382116103b757366023830112156103b757610efd9261104e611061933690602481600401359101610ace565b906064359261105c84610fd8565b612955565b60405191829182610fe3565b346103b75760403660031901126103b757602060ff6110a3602435600435611094826103c7565b5f526009845260405f2061274e565b54166040519015158152f35b90602080835192838152019201905f5b8181106110cc5750505090565b82518452602093840193909201916001016110bf565b9060206110f39281815201906110af565b90565b346103b75760203660031901126103b7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061115057610efd8561114481870382610629565b604051918291826110e2565b825484526020909301926001928301920161112d565b60e08101929161065b9190610d9f565b346103b75760403660031901126103b757610efd6111b560243560043561119c826103c7565b6111a46126c9565b505f52600660205260405f2061274e565b6004604051916111c4836105d3565b80546001600160a01b03168352600181015461121990611210906111f36111ea826103bb565b602088016124d6565b61120760a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611166565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916112999061128b6110f397959693600f60f81b865260e0602087015260e0860190611240565b908482036040860152611240565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526110af565b346103b7575f3660031901126103b7575f5160206157de5f395f51905f52541580611367575b1561132a576112f66137cf565b6112fe613887565b90610efd604051611310602082610629565b5f8082523660208301376040519384933091469186611264565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f5160206159be5f395f51905f5254156112e9565b9080601f830112156103b75781356113948161069a565b926113a26040519485610629565b81845260208085019260051b820101918383116103b75760208201905b8382106113ce57505050505090565b81356001600160401b0381116103b7576020916113f0878480948801016106b1565b8152019101906113bf565b60803660031901126103b7576004356001600160401b0381116103b757611426903690600401610449565b90602435906001600160401b0382116103b757366023830112156103b7578160040135916114538361069a565b926114616040519485610629565b8084526024602085019160051b830101913683116103b75760248101915b8383106114d75750506044359150506001600160401b0381116103b7576114aa90369060040161137d565b606435926001600160401b0384116103b757610efd946114d1610fc695369060040161137d565b93612a23565b82356001600160401b0381116103b7578201366043820112156103b7576024810135906115038261069a565b916115116040519384610629565b808352602060248185019260051b84010101913683116103b757604401905b8282106115485750505081526020928301920161147f565b60208091833561155781610fd8565b815201910190611530565b346103b75760603660031901126103b75760043561157f816103c7565b6024359061158c826103c7565b60443561159881610fd8565b5f51602061599e5f395f51905f52549260ff604085901c1615936001600160401b03168015908161169c575b6001149081611692575b159081611689575b5061167a575f51602061599e5f395f51905f5280546001600160401b0319166001179055611608928461165657613c20565b61160e57005b5f51602061599e5f395f51905f52805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b5f51602061599e5f395f51905f52805460ff60401b1916600160401b179055613c20565b63f92ee8a960e01b5f5260045ffd5b9050155f6115d6565b303b1591506115ce565b8591506115c4565b346103b75760403660031901126103b7576024356004356116c3612cbe565b6116cb612f10565b6116d58282613fc4565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f51602061591e5f395f51905f525d005b346103b7575f3660031901126103b7576033546040516001600160a01b039091168152602090f35b346103b7575f3660031901126103b7576020603454604051908152f35b346103b75760403660031901126103b757602060ff6110a3602435600435611778826103c7565b5f525f5160206158be5f395f51905f52845260405f2061274e565b346103b7575f3660031901126103b7576035546040516001600160a01b039091168152602090f35b346103b7575f3660031901126103b75760206040515f8152f35b60206040818301928281528451809452019201905f5b8181106117f85750505090565b82516001600160a01b03168452602093840193909201916001016117eb565b346103b75760203660031901126103b7576004355f525f51602061583e5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b81811061187d57610efd8561187181870382610629565b604051918291826117d5565b825484526020909301926001928301920161185a565b346103b75760203660031901126103b7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b360206004356118d2612cbe565b80600155604051908152a1005b346103b75760403660031901126103b7576024356004356118fe612d38565b611906612d38565b61190e612f10565b805f526007602052611927826108918160405f2061413b565b61193182826132ff565b5f51602061573e5f395f51905f525f80a35f5f51602061591e5f395f51905f525d005b60405190611963602083610629565b5f8252565b346103b7575f3660031901126103b757610efd604051611989604082610629565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611240565b346103b75760203660031901126103b7576004355f526004602052600160405f20015460018101809111610cf757602090604051908152f35b346103b7576119f63661070e565b90611a048151835114612478565b5f5b82518110156100205780611a3f611a1f6001938561248e565b51838060a01b03611a30848861248e565b51169061057261056d8261245a565b5001611a06565b634e487b7160e01b5f52602160045260245ffd5b600b1115611a6457565b611a46565b90600b821015611a645752565b6020815260606040611adb60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611240565b93611aed602082015183860190611a69565b015191015290565b346103b75760403660031901126103b757600435602435905f60408051611b1b816105f3565b611b23612b00565b815282602082015201525f52600860205260405f20905f52602052610efd60405f20600760405191611b54836105f3565b611b5d8161251d565b8352611b7360ff60068301541660208501612615565b0154604082015260405191829182611a76565b346103b75760203660031901126103b7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611bc881610fd8565b611bd0612cbe565b16611bdc811515612b30565b8060ff195f51602061577e5f395f51905f525416175f51602061577e5f395f51905f5255604051908152a1005b801515036103b757565b346103b75760403660031901126103b757600435602435611c3381611c09565b611c3b612e2c565b611c50825f52600360205260405f2054151590565b15611c9b5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611c9081600360405f2001612b46565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103b75760203660031901126103b757600435611ccb81611c09565b611cd3612e2c565b15611d3157611ce0612ee9565b600160ff195f5160206158de5f395f51905f525416175f5160206158de5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f5160206158de5f395f51905f525460ff811615611d885760ff19165f5160206158de5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103b75760803660031901126103b757602435600435611db7826103c7565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c6040604435611de681611c09565b606435611df281611c09565b611dfa612e2c565b845f526005602052611e5e81611e59855f2098611e2981611e2460018060a01b038216809d61413b565b612b57565b885f526006602052611e49866001611e43848b5f2061274e565b01612b7f565b885f526009602052865f2061274e565b612b46565b8251911515825215156020820152a3005b346103b75760203660031901126103b757600435611e8c816103c7565b611e94612cbe565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103b7575f3660031901126103b757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b818110611f3757610efd8561114481870382610629565b8254845260209093019260019283019201611f20565b346103b75760203660031901126103b757600435611f6a816103c7565b611f72612cbe565b6001600160a01b0316611f86811515612264565b603380546001600160a01b031916821790555f51602061585e5f395f51905f525f80a2005b346103b75760403660031901126103b757610020602435600435611fce826103c7565b61080d61056d8261245a565b346103b75760203660031901126103b7576004355f526004602052600260405f20015460018101809111610cf757602090604051908152f35b346103b75760203660031901126103b757600435612030816103c7565b612038612cbe565b6001600160a01b031661204c811515612b9c565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346103b75760803660031901126103b7576004356024356120ab81611c09565b604435916120b8836103c7565b5f51602061595e5f395f51905f52612198606435936120d6856103c7565b6120de612db2565b6120e9841515612b9c565b6001600160a01b03861694610fc690612103871515612b9c565b6001600160a01b038116976121939061211d8a1515612b9c565b61212688614e75565b61219d575b612146816121418161213c8c61240a565b6141de565b614094565b61216561215161065d565b9361215c83866124d6565b602085016124d6565b84151560408401525f60608401525f60808401525f60a08401525f60c084015261218e886123fc565b61274e565b6140bc565b0390a4005b6121cb6121a861064c565b8981525f60208201525f60408201525f60608201526121c68a6123ee565b614069565b61212b565b346103b7575f3660031901126103b7576032546040516001600160a01b039091168152602090f35b346103b7575f3660031901126103b75761221061540b565b612218615462565b6040519060208201925f5160206159de5f395f51905f528452604083015260608201524660808201523060a082015260a0815261225660c082610629565b519020604051908152602090f35b1561226b57565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b9190926122a4612ee9565b6122ac612f10565b6122b781851461239d565b5f5b8481106122cd57505050505061065b612f45565b8061232761230e6122ea6122e46001958a8a6123c7565b3561240a565b61230861230360206122fd878d8d6123c7565b01612418565b6103bb565b90612f57565b61232161230360206122fd868c8c6123c7565b90612422565b6123976123358288886123c7565b3561234660206122fd858b8b6123c7565b878961235860606122fd8884866123c7565b60c061238388608061236b8287896123c7565b01359460a061237b83838a6123c7565b0135966123c7565b013593612391888b8d61244a565b95612fd5565b016122b9565b156123a457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156123e95760051b8101359060fe19813603018212156103b7570190565b6123b3565b5f52600460205260405f2090565b5f52600660205260405f2090565b5f52600560205260405f2090565b356110f3816103c7565b1561242a5750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b91908110156123e95760e0020190565b5f525f5160206158be5f395f51905f52602052600160405f20015490565b1561247f57565b63031206d560e51b5f5260045ffd5b80518210156123e95760209160051b010190565b156124aa5750565b6373a1390160e11b5f5260045260245ffd5b156124c45750565b636fc24b7960e11b5f5260045260245ffd5b6001600160a01b039091169052565b90600182811c92168015612513575b60208310146124ff57565b634e487b7160e01b5f52602260045260245ffd5b91607f16916124f4565b9060405161252a8161060e565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91612582826124e5565b80855291600181169081156125ee57506001146125af575b505060a092916125ab910384610629565b0152565b5f908152602081209092505b8183106125d25750508101602001816125ab61259a565b60209193508060019154838589010152019101909184926125bb565b60ff191660208681019190915292151560051b850190920192508391506125ab905061259a565b600b821015611a645752565b9060405161262e816105f3565b60406007829461263d8161251d565b845261265360ff60068301541660208601612615565b0154910152565b156126625750565b60405162461bcd60e51b815260206004820152908190612686906024830190611240565b0390fd5b15612693575050565b637a88099160e11b5f5260045260245260445ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610cf757565b604051906126d6826105d3565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906127098261069a565b6127166040519182610629565b8281528092612727601f199161069a565b01905f5b82811061273757505050565b6020906127426126c9565b8282850101520161272b565b9060018060a01b03165f5260205260405f2090565b90604051612770816105d3565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916127bb906127b290611207565b15156060860152565b60028101546080850152600381015460a08501520154910152565b97969594939291906127e6612ee9565b60018060a01b03811698805f5260056020526128098a611e248160405f2061413b565b805f5260046020526128268160ff600360405f20015416156124bc565b805f52600660205260ff600161283f8c60405f2061274e565b015460a81c16612855576110f39899505f612873565b6338384f6f60e11b5f9081526001600160a01b038b16600452602490fd5b50959391949692612882612f10565b6128966001600160a01b0387161515612264565b600a54801590811561292a575b501561291b576128f0612906966128c36129109a6128da9685878d613536565b9590946128ce61066c565b9a8b5260208b016124d6565b6128e73360408b016124d6565b606089016124d6565b608087015260a086015260c08501523691610ace565b60e08201526136d1565b60019061065b612f45565b633953b33f60e11b5f5260045ffd5b90508411155f6128a3565b908160209103126103b757516110f3816103c7565b6040513d5f823e3d90fd5b5f5490949392906001600160a01b03811615612a1457604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906129be906084870190611240565b931691015203925af18015612a0f5761065b925f916129e0575b508094612bb2565b612a02915060203d602011612a08575b6129fa8183610629565b810190612935565b5f6129d8565b503d6129f0565b61294a565b6315aeca0d60e11b5f5260045ffd5b919294939094612a31612ee9565b612a39612f10565b8584511480612ad8575b80612ace575b612a559096959661239d565b612a60343415612ae2565b3683900360be1901955f5b86811015612ab9578060051b85013590888212156103b757612ab3600192612a93838a61248e565b51612a9e848861248e565b5190612aaa858a61248e565b51928a01613999565b01612a6b565b5095505050505050612ac9612f45565b600190565b5081518614612a49565b5085815114612a43565b15612aea5750565b63943892eb60e01b5f525f60045260245260445ffd5b60405190612b0d8261060e565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b15612b3757565b63f0f15b9160e01b5f5260045ffd5b9060ff801983541691151516179055565b15612b5f5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b15612ba357565b6302bfb33d60e51b5f5260045ffd5b919091612bbd612db2565b612bc8811515612b9c565b6001600160a01b03831691612c4d90612be2841515612b9c565b6001600160a01b0381169461219390612bfc871515612b9c565b612c0585614e75565b612c6d575b612c1b816121418161213c8961240a565b612c2661215161065d565b5f60408401525f60608401525f60808401525f60a08401525f60c084015261218e856123fc565b6040515f81525f51602061595e5f395f51905f529080602081015b0390a4565b612c96612c7861064c565b8681525f60208201525f60408201525f60608201526121c6876123ee565b612c0a565b6035546001600160a01b03163303612caf57565b63c82cebcb60e01b5f5260045ffd5b5f51602061593e5f395f51905f525f525f5160206158be5f395f51905f5260205260ff612d0b337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61274e565b541615612d1457565b63e2517d3f60e01b5f52336004525f51602061593e5f395f51905f5260245260445ffd5b5f51602061597e5f395f51905f525f525f5160206158be5f395f51905f5260205260ff612d85337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361274e565b541615612d8e57565b63e2517d3f60e01b5f52336004525f51602061597e5f395f51905f5260245260445ffd5b5f5160206158fe5f395f51905f525f525f5160206158be5f395f51905f5260205260ff612dff337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61274e565b541615612e0857565b63e2517d3f60e01b5f52336004525f5160206158fe5f395f51905f5260245260445ffd5b5f51602061587e5f395f51905f525f525f5160206158be5f395f51905f5260205260ff612e79337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661274e565b541615612e8257565b63e2517d3f60e01b5f52336004525f51602061587e5f395f51905f5260245260445ffd5b805f525f5160206158be5f395f51905f5260205260ff612ec93360405f2061274e565b541615612ed35750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f5160206158de5f395f51905f525416612f0157565b63d93c066560e01b5f5260045ffd5b5f51602061591e5f395f51905f525c612f365760015f51602061591e5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f51602061591e5f395f51905f525d565b6110f3916001600160a01b03169061413b565b15612f73575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612f9e57565b630672aec160e01b5f5260045ffd5b15612fb6575050565b63943892eb60e01b5f5260045260245260445ffd5b356110f381610fd8565b94919392916130399061300d6001600160a01b038416612ff8612303808c612418565b8114906130076123038c612418565b91612f6a565b602088019461303161302161230388612418565b6001600160a01b03891614612f97565b848489613536565b949093604088013561306f81613058896130538a8a6126bc565b6126bc565b8110156130698a6130538b8b6126bc565b90612fad565b61307d6123036032546103bb565b9061308783612418565b60608b01359a61309960808201612fcb565b9060c08101359060a00135853b156103b757604051637f40ec1760e11b81526001600160a01b038a811660048301529490941660248501523060448501526064840194909452608483019c909c5260ff1660a482015260c481019190915260e481019990995288610104815a5f948591f1928315612a0f5761314361312e61314c9361065b9b6131559761316c575b50612418565b9161313761066c565b998a5260208a016124d6565b604088016124d6565b606086016124d6565b608084015260a083015260c0820152612906611954565b8061317a5f61318093610629565b806104f1565b5f613128565b90613191825f61414e565b918261319a5750565b5f80525f51602061583e5f395f51905f526020526001600160a01b03166131e1817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d614ede565b156131e95750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161320c838261414e565b9283613216575050565b815f525f51602061583e5f395f51905f5260205261324160405f209160018060a01b03168092614ede565b1561324a575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161326c83826141f1565b9283613276575050565b815f525f51602061583e5f395f51905f526020526132a160405f209160018060a01b03168092614f60565b156132aa575050565b62a95f1b60e31b5f5260045260245260445ffd5b156132c65750565b63290cd55f60e01b5f52600b811015611a645760045260245ffd5b604091949392606082019560018060a01b0316825260208201520152565b9061330991613fc4565b60018060a01b036060820151168151905f51602061579e5f395f51905f5260208401918251946133596108f76040830196875160018060a01b03169885608086019a8b519260a0880151946143d4565b519251935194516040516001600160a01b0390961695918291612c6891429190846132e1565b1561338657565b6330d45fb160e01b5f5260045ffd5b908160209103126103b75751600b8110156103b75790565b6001600160a01b039091168152602081019190915260400190565b915061340960ff916133f85f946133ea60325460018060a01b0316151561337f565b5f52600960205260405f2090565b6001600160a01b039091169061274e565b541661341457600191565b506002905f90565b9190915f9261346461345d61344d6134386123036032546103bb565b946133ea6001600160a01b038716151561337f565b6001600160a01b0384169061274e565b5460ff1690565b6134eb579160209161348e95935f604051809881958294633f4bdec560e01b8452600484016133ad565b03925af1928315612a0f575f936134ba575b506001836134ad81611a5a565b036134b457565b60019150565b6134dd91935060203d6020116134e4575b6134d58183610629565b810190613395565b915f6134a0565b503d6134cb565b505050506002905f90565b908160209103126103b7575190565b908160609103126103b7578051916040602083015192015190565b1561352757565b6358e8878560e01b5f5260045ffd5b826060916135bf97959693613550610ece61344d846123fc565b6135606109dd6040830151151590565b613661575b506135746123036032546103bb565b916135896001600160a01b038416151561337f565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612a0f575f955f905f93613623575b5090829161065b94939681988515159586613618575b50508461360d575b505082613602575b5050613520565b101590505f806135fb565b101592505f806135f3565b101594505f806135eb565b905061364e91965061065b93925060603d60601161365a575b6136468183610629565b810190613505565b919692939192916135d5565b503d61363c565b60c06136739101518480821015612fad565b5f613565565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916136cc9190860190611240565b930152565b6136ef81515f526004602052600160405f2001600181540180915590565b6136f982516123fc565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff228070861373e61230360016137376020880195610ec861230388516103bb565b01546103bb565b93805190612c6861374f85516103bb565b91604081019061375f82516103bb565b906137886080820196875160a084019485519861378260c087019a8b51906126bc565b936147e3565b61379e613797825199516103bb565b93516103bb565b9460e06137ae60608401516103bb565b9751935191519201519260405197889760018060a01b03169c429689613679565b6040515f51602061575e5f395f51905f5254815f6137ec836124e5565b80835292600181169081156138685750600114613810575b6110f392500382610629565b505f51602061575e5f395f51905f525f90815290915f5160206157fe5f395f51905f525b81831061384c5750509060206110f392820101613804565b6020919350806001915483858801015201910190918392613834565b602092506110f394915060ff191682840152151560051b820101613804565b6040515f5160206157be5f395f51905f5254815f6138a4836124e5565b808352926001811690811561386857506001146138c7576110f392500382610629565b505f5160206157be5f395f51905f525f90815290915f5160206159fe5f395f51905f525b8183106139035750509060206110f392820101613804565b60209193508060019154838588010152019101909183926138eb565b15613928575050565b63a6ab65c560e01b5f5260045260245260445ffd5b903590601e19813603018212156103b757018035906001600160401b0382116103b7576020019181360383136103b757565b61065b93606092969593608083019760018060a01b03168352602083015260408201520190611a69565b9192823591826139a8816123ee565b6003015460ff1615906139ba916124bc565b6139c38361240a565b9060408501916139d283612418565b6139db906103bb565b6139e491612f57565b6139ed83612418565b6139f6906103bb565b6139ff91612422565b613a1c845f526004602052600260405f2001600181540180915590565b946020810135958681811491613a319261391f565b613a3a83612418565b613a43906103bb565b916060820194613a5286612418565b988860808501359a60a08601968a8d613a6b8a8a61393d565b3690613a7692610ace565b8051602091820120604080517fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219149381019384529081019490945260608401969096526001600160a01b0393841660808401529390921660a082015260c081019190915260e08082019390935291825290613af261010082610629565b51902092613aff93614954565b5f9087613b0b85612418565b90613b16918861341c565b91909384613b2381611a5a565b600114613bde575b50613b3584611a5a565b60018403613b7f575050505090613b62613b5c5f51602061579e5f395f51905f5293612418565b91612418565b6040516001600160a01b03909216958291612c68914291846132e1565b83949850613bba612c6893613bc6937f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896613bc094614cb1565b98612418565b92612418565b9260405193849360018060a01b03169842918561396f565b909350613c1791925088613bf186612418565b91613c0f613c08613c018a612418565b928861393d565b3691610ace565b928a8a6143d4565b9190925f613b2b565b91613c29614dd0565b613c3d6001600160a01b0384161515612264565b6001600160a01b03821692613c53841515612264565b60ff821615613ed157613cbf90613c68614dd0565b613c70614dd0565b613c78614dd0565b60ff195f5160206158de5f395f51905f5254165f5160206158de5f395f51905f5255613ca2614dd0565b613caa614dd0565b613cb2614dd0565b613cba614dd0565b613186565b50613cc8614dd0565b613cd660ff82161515612b30565b60408051613ce48282610629565b60098152682b30b634b230ba37b960b91b6020820152613d0682519283610629565b60058252640312e302e360dc1b6020830152613d20614dd0565b613d28614dd0565b8051906001600160401b0382116105ee57613d5982613d545f51602061575e5f395f51905f52546124e5565b614ad6565b602090601f8311600114613e3d5792613d8883613e18979694613d9c94613dee975f92613e32575b5050613f20565b5f51602061575e5f395f51905f5255615648565b613db15f5f5160206157de5f395f51905f5255565b613dc65f5f5160206159be5f395f51905f5255565b60ff1660ff195f51602061577e5f395f51905f525416175f51602061577e5f395f51905f5255565b613df6614dfb565b603380546001600160a01b0319166001600160a01b0392909216919091179055565b5f51602061585e5f395f51905f525f80a261065b43603455565b015190505f80613d81565b5f51602061575e5f395f51905f525f52601f19831691905f5160206157fe5f395f51905f52925f5b818110613eb9575093613d9c93613dee969360019383613e189b9a9810613ea1575b505050811b015f51602061575e5f395f51905f5255615648565b01515f1960f88460031b161c191690555f8080613e87565b92936020600181928786015181550195019301613e65565b6338854f1160e21b5f5260045ffd5b91908203918211610cf757565b91613f069183549060031b91821b915f19901b19161790565b9055565b818110613f15575050565b5f8155600101613f0a565b8160011b915f199060031b1c19161790565b60075f9182815582600182015582600282015582600382015582600482015560058101613f5f81546124e5565b9081613f72575b50508260068201550155565b601f8211600114613f8957849055505b5f80613f66565b613faf613fbf926001601f613fa1855f5260205f2090565b920160051c82019101613f0a565b5f81815260208120918190559055565b613f82565b9190613fce612b00565b50825f526007602052613fe48160405f20614f60565b156140575761405261065b91845f52600860205260405f20815f52602052610a2361401160405f2061251d565b9561402e61401e826123fc565b610ec861230360408b01516103bb565b614042600260808a01519201918254613ee0565b90555f52600860205260405f2090565b613f32565b637f11bea960e01b5f5260045260245ffd5b6003606061065b93805184556020810151600185015560408101516002850155015115159101612b46565b1561409c5750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916140d560018060a01b038251168561227a565b60208101516141209060018601906140f6906001600160a01b03168261227a565b6040830151815460ff60a01b191690151560a01b60ff60a01b161781556060830151151590612b7f565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b805f525f5160206158be5f395f51905f5260205260ff6141718360405f2061274e565b54166141d857805f525f5160206158be5f395f51905f526020526141988260405f2061274e565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b6110f3916001600160a01b031690614ede565b805f525f5160206158be5f395f51905f5260205260ff6142148360405f2061274e565b5416156141d857805f525f5160206158be5f395f51905f5260205261423c8260405f2061274e565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b3d156142a3573d9061428a82610ab3565b916142986040519384610629565b82523d5f602084013e565b606090565b908160209103126103b757516110f381611c09565b156142c457565b632b60b36f60e21b5f5260045ffd5b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a082018190526110f392910190611240565b91906040838203126103b7578251602084015190936001600160401b0382116103b7570181601f820112156103b75780519061434d82610ab3565b9261435b6040519485610629565b828452602083830101116103b757815f9260208093018386015e8301015290565b6110f3939260809263ffffffff60e01b1682526001602083015260408201528160608201520190611240565b6080906110f3939263ffffffff60e01b1681525f60208201525f60408201528160608201520190611240565b939594928692614403946143e7876123fc565b6001600160a01b0383169687929161440e91600191859061274e565b015460a01c60ff1690565b9461441d6123036035546103bb565b91601881511015806146e5575b806146db575b61446f575b5050505090614444939161513a565b9161444e83611a5a565b6001831461445c5750509190565b614468939492506150f1565b6001905f90565b602081015160601c928a6034830151925f5f80604051602081019062483e5560e91b82526144b2816144a48d60248301610fe3565b03601f198101835282610629565b5190875afa6144bf614279565b90806146cf575b6146c3575b506144d857505050614435565b9087858a93899d60018e999a9b14159e8f998a6145f4575b50505f959391926144a4879694614529938c89146145ed5788975b6040519687956020870199631eeed51360e01b8b52602488016142d3565b5191855af190614537614279565b936145dd575b50806145d1575b6145845791895f51602061589e5f395f51905f5261444499989795938e9795614572604051928392836143a8565b0390a484156151f957831594506151f9565b908997505f51602061589e5f395f51905f529396506145c99295506144689a9b999450806020806145ba93518301019101614312565b6040939193519384938461437c565b0390a46150f1565b50604082511015614544565b6145e79087615003565b5f61453d565b809761450b565b15614626575b505f959391926144a487969483614615614529958c886150c8565b935093959750509294508f8e6144f0565b92509097505f94506020935061465492506040519485809481936340c10f1960e01b835230600484016133ad565b03925af1938415612a0f57888f898f91898f995f9796614529936146846144a4938b9a8b91614694575b506142bd565b935093959750509294508f6145fa565b6146b6915060203d6020116146bc575b6146ae8183610629565b8101906142a8565b5f61467e565b503d6146a4565b6020915001515f6144cb565b506020815110156144c6565b50823b1515614430565b506001600160a01b038316151561442a565b90813b15614778575f51602061581e5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156147605761475d9161526c565b50565b50503461476957565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b156147a357505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156147d457565b63559d141b60e11b5f5260045ffd5b90936001600160a01b0385169260018403614851575061065b945061481961480b82866126bc565b3414349061306984886126bc565b80614825575b506153aa565b61484661484b916148376033546103bb565b90614840611954565b9161534b565b6147cd565b5f61481f565b9061485d343415612ae2565b61487261486a82876126bc565b308489615289565b80614906575b5061488e6109dd60016144038661218e876123fc565b61489e575b5061065b93506153aa565b604051632770a7eb60e21b8152602081806148bd8830600484016133ad565b03815f885af1918215612a0f5761065b966148e19387935f916148e7575b50614799565b5f614893565b614900915060203d6020116146bc576146ae8183610629565b5f6148db565b61491e906149186123036033546103bb565b876152cd565b5f614878565b1561492b57565b63b3f07a3960e01b5f5260045ffd5b156149425750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480614ab2575b61496d90614924565b61498e6149885f51602061577e5f395f51905f525460ff1690565b60ff1690565b9561499c848881101561493a565b5f945f935f5b8681106149bd575050505050505061065b919281101561493a565b614a1a6149e9836149cc614e0c565b6042916040519161190160f01b8352600283015260228201522090565b6149fd6149f6848961248e565b5160ff1690565b614a07848761248e565b5190614a13858961248e565b51926153f3565b6001600160a01b038181169088161080614a4b575b614a3d575b506001016149a2565b600198890198909650614a34565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f5160206158be5f395f51905f52602052614aad61345d827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561274e565b614a2f565b5085518314614964565b15614ac45750565b6307a5c91d60e31b5f5260045260245ffd5b90601f8211614ae3575050565b61065b915f51602061575e5f395f51905f525f5260205f20906020601f840160051c83019310614b1b575b601f0160051c0190613f0a565b9091508190614b0e565b9190601f8111614b3457505050565b61065b925f5260205f20906020601f840160051c83019310614b1b57601f0160051c0190613f0a565b90600b811015611a645760ff80198354169116179055565b815180518255602081015160018301556040810151919291614ba3906001600160a01b03166002850161227a565b6060810151614bbe906001600160a01b03166003850161227a565b6080810151600484015560a00151805160058401916001600160401b0382116105ee57614bf582614bef85546124e5565b85614b25565b602090601f8311600114614c4157826007959360409593614c1d935f92613e32575050613f20565b90555b614c3a6020820151614c3181611a5a565b60068601614b5d565b0151910155565b90601f19831691614c55855f5260205f2090565b925f5b818110614c995750926001928592600798966040989610614c81575b505050811b019055614c20565b01515f1960f88460031b161c191690555f8080614c74565b92936020600181928786015181550195019301614c58565b93929093614cde813592614ce3614cd0855f52600760205260405f2090565b602085013593848092614ede565b614abc565b8015614dc457935b84966040840191614cfb83612418565b94606001614d0890612418565b614d10611954565b90614d1961067c565b968888528660208901526040880190614d31916124d6565b614d3e90606088016124d6565b87608087015260a08601525f14613bc0613f0696614dac600297614da7614db198610ec89761230397614dba57614d90614d7a600154426126bc565b915b614d8461068b565b95865260208601612615565b6040840152610a23855f52600860205260405f2090565b614b75565b6123fc565b019182546126bc565b614d905f91614d7c565b50608082013593614ceb565b60ff5f51602061599e5f395f51905f525460401c1615614dec57565b631afcd79f60e31b5f5260045ffd5b614e03614dd0565b62015180600155565b614e1461540b565b614e1c615462565b6040519060208201925f5160206159de5f395f51905f528452604083015260608201524660808201523060a082015260a08152614e5a60c082610629565b51902090565b80548210156123e9575f5260205f2001905f90565b5f81815260036020526040902054614ed957600254600160401b8110156105ee57614ec2614eac8260018594016002556002614e60565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b614ee8828261413b565b6141d857805490600160401b8210156105ee5782614f10614eac846001809601855584614e60565b90558054925f520160205260405f2055600190565b80548015614f4c575f190190614f3b8282614e60565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14614ffb575f198401848111610cf75783545f19810194908511610cf7575f958583614fae97610a239503614fb4575b505050614f25565b55600190565b614fe4614fde91614fd5614fcb614ff29588614e60565b90549060031b1c90565b92839187614e60565b90613eed565b85905f5260205260405f2090565b555f8080614fa6565b505050505f90565b6040519060205f81840163095ea7b360e01b8152615037856150298489602484016133ad565b03601f198101875286610629565b84519082855af15f51903d8161509c575b501590505b61505657505050565b60405163095ea7b360e01b60208201526001600160a01b0390931660248401525f604484015261065b926150979061509181606481016144a4565b82615494565b615494565b151590506150bc575061504d6001600160a01b0382163b15155b5f615048565b600161504d91146150b6565b91909160205f60405193615037856150298582019363095ea7b360e01b855289602484016133ad565b90615106915f52600660205260405f2061274e565b600181015460a01c60ff1615615128576003018054918203918211610cf75755565b6004018054918201809211610cf75755565b6001600160a01b0316929190600184146151c95781156151c05761518992156151945760405163a9059cbb60e01b6020820152916151819183916144a491602484016133ad565b600592615302565b156110f35750600190565b6040516340c10f1960e01b6020820152916151b89183916144a491602484016133ad565b600692615302565b50505050600190565b906151eb93506109dd92506151dc611954565b916001600160a01b031661534b565b6151f457600190565b600590565b6001600160a01b0316936001851461525857821561524e57615189938115615246575b50156151945760405163a9059cbb60e01b6020820152916151819183916144a491602484016133ad565b90505f61521c565b5050505050600190565b50906151eb93506109dd92506151dc611954565b5f806110f393602081519101845af4615283614279565b916154ec565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815261065b91615097608483610629565b61509761065b93926152f460405194859263a9059cbb60e01b6020850152602484016133ad565b03601f198101845283610629565b81516001600160a01b03909116915f9182916020018285620186a0f1615326614279565b90156141d857805161534257503b1561533e57600190565b5f90565b60209150015190565b5f809161535a844710156142bd565b84516001600160a01b0391909116946020018486620186a0f19061537c614279565b91156153a3571561538f575b5050600190565b805161534257503b1561533e575f80615388565b5050505f90565b906153bf915f52600660205260405f2061274e565b600181015460a01c60ff16156153e1576003018054918201809211610cf75755565b6004018054918203918211610cf75755565b916110f393916154029361554a565b909291926155cc565b6154136137cf565b8051908115615423576020012090565b50505f5160206157de5f395f51905f5254801561543d5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61546a613887565b805190811561547a576020012090565b50505f5160206159be5f395f51905f5254801561543d5790565b905f602091828151910182855af11561294a575f513d6154e357506001600160a01b0381163b155b6154c35750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156154bc565b90615510575080511561550157805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615541575b615521575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615519565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116155b7579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612a0f575f516001600160a01b038116156155ad57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611a6457565b6155d5816155c2565b806155de575050565b6155e7816155c2565b600181036155fe5763f645eedf60e01b5f5260045ffd5b615607816155c2565b60028103615622575063fce698f760e01b5f5260045260245ffd5b8061562e6003926155c2565b146156365750565b6335e2f38360e21b5f5260045260245ffd5b80519091906001600160401b0381116105ee57615689816156765f5160206157be5f395f51905f52546124e5565b5f5160206157be5f395f51905f52614b25565b602092601f82116001146156bc576156ab929382915f92613e32575050613f20565b5f5160206157be5f395f51905f5255565b5f5160206157be5f395f51905f525f52601f198216935f5160206159fe5f395f51905f52915f5b868110615725575083600195961061570d575b505050811b015f5160206157be5f395f51905f5255565b01515f1960f88460031b161c191690555f80806156f6565b919260206001819286850151815501940192016156e356feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220cbcfb3f1e07a76393bb76de87220d3d5115e71c69f239210a485c23c9843d0a564736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BaseBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseBridgeMetaData.ABI instead.
var BaseBridgeABI = BaseBridgeMetaData.ABI

// Deprecated: Use BaseBridgeMetaData.Sigs instead.
// BaseBridgeFuncSigs maps the 4-byte function signature to its string representation.
var BaseBridgeFuncSigs = BaseBridgeMetaData.Sigs

// BaseBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseBridgeMetaData.Bin instead.
var BaseBridgeBin = BaseBridgeMetaData.Bin

// DeployBaseBridge deploys a new Ethereum contract, binding an instance of BaseBridge to it.
func DeployBaseBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseBridge, error) {
	parsed, err := BaseBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseBridge{BaseBridgeCaller: BaseBridgeCaller{contract: contract}, BaseBridgeTransactor: BaseBridgeTransactor{contract: contract}, BaseBridgeFilterer: BaseBridgeFilterer{contract: contract}}, nil
}

// BaseBridge is an auto generated Go binding around an Ethereum contract.
type BaseBridge struct {
	BaseBridgeCaller     // Read-only binding to the contract
	BaseBridgeTransactor // Write-only binding to the contract
	BaseBridgeFilterer   // Log filterer for contract events
}

// BaseBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseBridgeSession struct {
	Contract     *BaseBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseBridgeCallerSession struct {
	Contract *BaseBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BaseBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseBridgeTransactorSession struct {
	Contract     *BaseBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BaseBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseBridgeRaw struct {
	Contract *BaseBridge // Generic contract binding to access the raw methods on
}

// BaseBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseBridgeCallerRaw struct {
	Contract *BaseBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BaseBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseBridgeTransactorRaw struct {
	Contract *BaseBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseBridge creates a new instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridge(address common.Address, backend bind.ContractBackend) (*BaseBridge, error) {
	contract, err := bindBaseBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseBridge{BaseBridgeCaller: BaseBridgeCaller{contract: contract}, BaseBridgeTransactor: BaseBridgeTransactor{contract: contract}, BaseBridgeFilterer: BaseBridgeFilterer{contract: contract}}, nil
}

// NewBaseBridgeCaller creates a new read-only instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeCaller(address common.Address, caller bind.ContractCaller) (*BaseBridgeCaller, error) {
	contract, err := bindBaseBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeCaller{contract: contract}, nil
}

// NewBaseBridgeTransactor creates a new write-only instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseBridgeTransactor, error) {
	contract, err := bindBaseBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTransactor{contract: contract}, nil
}

// NewBaseBridgeFilterer creates a new log filterer instance of BaseBridge, bound to a specific deployed contract.
func NewBaseBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseBridgeFilterer, error) {
	contract, err := bindBaseBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeFilterer{contract: contract}, nil
}

// bindBaseBridge binds a generic wrapper to an already deployed contract.
func bindBaseBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBridge *BaseBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseBridge.Contract.BaseBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBridge *BaseBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.Contract.BaseBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBridge *BaseBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBridge.Contract.BaseBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseBridge *BaseBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseBridge *BaseBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseBridge *BaseBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BaseBridge.Contract.DEFAULTADMINROLE(&_BaseBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BaseBridge.Contract.DEFAULTADMINROLE(&_BaseBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BaseBridge.Contract.UPGRADEINTERFACEVERSION(&_BaseBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BaseBridge *BaseBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BaseBridge.Contract.UPGRADEINTERFACEVERSION(&_BaseBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _BaseBridge.Contract.AllChainIDs(&_BaseBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BaseBridge *BaseBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BaseBridge.Contract.AllChainIDs(&_BaseBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BaseBridge.Contract.AllPendingIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BaseBridge *BaseBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BaseBridge.Contract.AllPendingIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BaseBridge *BaseBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BaseBridge *BaseBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BaseBridge *BaseBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.AllTokenPairs(&_BaseBridge.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BaseBridge *BaseBridgeCaller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BaseBridge *BaseBridgeSession) BridgeExecutor() (common.Address, error) {
	return _BaseBridge.Contract.BridgeExecutor(&_BaseBridge.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) BridgeExecutor() (common.Address, error) {
	return _BaseBridge.Contract.BridgeExecutor(&_BaseBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeSession) BridgeVerifier() (common.Address, error) {
	return _BaseBridge.Contract.BridgeVerifier(&_BaseBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _BaseBridge.Contract.BridgeVerifier(&_BaseBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _BaseBridge.Contract.CrossMintableERC20Code(&_BaseBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _BaseBridge.Contract.CrossMintableERC20Code(&_BaseBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeSession) Dev() (common.Address, error) {
	return _BaseBridge.Contract.Dev(&_BaseBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BaseBridge *BaseBridgeCallerSession) Dev() (common.Address, error) {
	return _BaseBridge.Contract.Dev(&_BaseBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) DomainSeparator() ([32]byte, error) {
	return _BaseBridge.Contract.DomainSeparator(&_BaseBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _BaseBridge.Contract.DomainSeparator(&_BaseBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BaseBridge *BaseBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_BaseBridge *BaseBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BaseBridge.Contract.Eip712Domain(&_BaseBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BaseBridge *BaseBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BaseBridge.Contract.Eip712Domain(&_BaseBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextFinalizeIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextFinalizeIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextInitiateIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BaseBridge.Contract.GetNextInitiateIndex(&_BaseBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BaseBridge *BaseBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BaseBridge *BaseBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BaseBridge *BaseBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BaseBridge.Contract.GetPendingArguments(&_BaseBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BaseBridge.Contract.GetRoleAdmin(&_BaseBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BaseBridge.Contract.GetRoleAdmin(&_BaseBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BaseBridge.Contract.GetRoleMembers(&_BaseBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BaseBridge *BaseBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BaseBridge.Contract.GetRoleMembers(&_BaseBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BaseBridge *BaseBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BaseBridge *BaseBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BaseBridge *BaseBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BaseBridge.Contract.GetTokenPair(&_BaseBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BaseBridge.Contract.HasRole(&_BaseBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BaseBridge.Contract.HasRole(&_BaseBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeSession) InitializedAt() (*big.Int, error) {
	return _BaseBridge.Contract.InitializedAt(&_BaseBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _BaseBridge.Contract.InitializedAt(&_BaseBridge.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BaseBridge *BaseBridgeSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BaseBridge.Contract.IsTokenFinalizePaused(&_BaseBridge.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BaseBridge.Contract.IsTokenFinalizePaused(&_BaseBridge.CallOpts, remoteChainID, token)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) MaxExtraDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "maxExtraDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BaseBridge *BaseBridgeSession) MaxExtraDataLength() (*big.Int, error) {
	return _BaseBridge.Contract.MaxExtraDataLength(&_BaseBridge.CallOpts)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) MaxExtraDataLength() (*big.Int, error) {
	return _BaseBridge.Contract.MaxExtraDataLength(&_BaseBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeSession) Paused() (bool, error) {
	return _BaseBridge.Contract.Paused(&_BaseBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) Paused() (bool, error) {
	return _BaseBridge.Contract.Paused(&_BaseBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _BaseBridge.Contract.ProxiableUUID(&_BaseBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BaseBridge *BaseBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BaseBridge.Contract.ProxiableUUID(&_BaseBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseBridge.Contract.SupportsInterface(&_BaseBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BaseBridge *BaseBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BaseBridge.Contract.SupportsInterface(&_BaseBridge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeSession) Threshold() (uint8, error) {
	return _BaseBridge.Contract.Threshold(&_BaseBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BaseBridge *BaseBridgeCallerSession) Threshold() (uint8, error) {
	return _BaseBridge.Contract.Threshold(&_BaseBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.BridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.BridgeToken(&_BaseBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.ChangeThreshold(&_BaseBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.ChangeThreshold(&_BaseBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BaseBridge *BaseBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.CreateToken(&_BaseBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridgeBatch(&_BaseBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BaseBridge *BaseBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.FinalizeBridgeBatch(&_BaseBridge.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRole(&_BaseBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRole(&_BaseBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRoleBatch(&_BaseBridge.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.GrantRoleBatch(&_BaseBridge.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "initialize", owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BaseBridge.Contract.Initialize(&_BaseBridge.TransactOpts, owner, dev_, threshold_)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BaseBridge *BaseBridgeTransactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BaseBridge *BaseBridgeSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualReleasePendingWithRecipient(&_BaseBridge.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.ManualReleasePendingWithRecipient(&_BaseBridge.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeTokenBatch(&_BaseBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BaseBridge.Contract.PermitBridgeTokenBatch(&_BaseBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RegisterToken(&_BaseBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.ReleasePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.RemovePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.RemovePending(&_BaseBridge.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceRole(&_BaseBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RenounceRole(&_BaseBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRole(&_BaseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRole(&_BaseBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRoleBatch(&_BaseBridge.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BaseBridge *BaseBridgeTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.RevokeRoleBatch(&_BaseBridge.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BaseBridge *BaseBridgeTransactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BaseBridge *BaseBridgeSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeExecutor(&_BaseBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeExecutor(&_BaseBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeVerifier(&_BaseBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetBridgeVerifier(&_BaseBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetChainPause(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetChainPause(&_BaseBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetCrossMintableERC20Code(&_BaseBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetCrossMintableERC20Code(&_BaseBridge.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetDev(&_BaseBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetDev(&_BaseBridge.TransactOpts, dev_)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BaseBridge *BaseBridgeTransactor) SetMaxExtraDataLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setMaxExtraDataLength", length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BaseBridge *BaseBridgeSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetMaxExtraDataLength(&_BaseBridge.TransactOpts, length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetMaxExtraDataLength(&_BaseBridge.TransactOpts, length)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPause(&_BaseBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPause(&_BaseBridge.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BaseBridge *BaseBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BaseBridge *BaseBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetTokenPause(&_BaseBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetTokenPause(&_BaseBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelay(&_BaseBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelay(&_BaseBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelayExpiration(&_BaseBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetVerificationDelayExpiration(&_BaseBridge.TransactOpts, remoteChainID, index, delay)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.UpgradeToAndCall(&_BaseBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BaseBridge.Contract.UpgradeToAndCall(&_BaseBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeSession) Receive() (*types.Transaction, error) {
	return _BaseBridge.Contract.Receive(&_BaseBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseBridge *BaseBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BaseBridge.Contract.Receive(&_BaseBridge.TransactOpts)
}

// BaseBridgeBridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the BaseBridge contract.
type BaseBridgeBridgeExecutorSetIterator struct {
	Event *BaseBridgeBridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeExecutorSet)
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
		it.Event = new(BaseBridgeBridgeExecutorSet)
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
func (it *BaseBridgeBridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeExecutorSet represents a BridgeExecutorSet event raised by the BaseBridge contract.
type BaseBridgeBridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*BaseBridgeBridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeExecutorSetIterator{contract: _BaseBridge.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeExecutorSet)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeExecutorSet(log types.Log) (*BaseBridgeBridgeExecutorSet, error) {
	event := new(BaseBridgeBridgeExecutorSet)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BaseBridge contract.
type BaseBridgeBridgeFinalizedIterator struct {
	Event *BaseBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeFinalized)
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
		it.Event = new(BaseBridgeBridgeFinalized)
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
func (it *BaseBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeFinalized represents a BridgeFinalized event raised by the BaseBridge contract.
type BaseBridgeBridgeFinalized struct {
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
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BaseBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeFinalizedIterator{contract: _BaseBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeFinalized)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeFinalized(log types.Log) (*BaseBridgeBridgeFinalized, error) {
	event := new(BaseBridgeBridgeFinalized)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BaseBridge contract.
type BaseBridgeBridgeInitiatedIterator struct {
	Event *BaseBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeInitiated)
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
		it.Event = new(BaseBridgeBridgeInitiated)
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
func (it *BaseBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeInitiated represents a BridgeInitiated event raised by the BaseBridge contract.
type BaseBridgeBridgeInitiated struct {
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
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*BaseBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeInitiatedIterator{contract: _BaseBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeInitiated)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeInitiated(log types.Log) (*BaseBridgeBridgeInitiated, error) {
	event := new(BaseBridgeBridgeInitiated)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the BaseBridge contract.
type BaseBridgeBridgePendingIterator struct {
	Event *BaseBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgePending)
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
		it.Event = new(BaseBridgeBridgePending)
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
func (it *BaseBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgePending represents a BridgePending event raised by the BaseBridge contract.
type BaseBridgeBridgePending struct {
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
func (_BaseBridge *BaseBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BaseBridgeBridgePendingIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgePendingIterator{contract: _BaseBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgePending)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseBridgePending(log types.Log) (*BaseBridgeBridgePending, error) {
	event := new(BaseBridgeBridgePending)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the BaseBridge contract.
type BaseBridgeBridgeVerifierSetIterator struct {
	Event *BaseBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeBridgeVerifierSet)
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
		it.Event = new(BaseBridgeBridgeVerifierSet)
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
func (it *BaseBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the BaseBridge contract.
type BaseBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BaseBridge *BaseBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*BaseBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeBridgeVerifierSetIterator{contract: _BaseBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BaseBridge *BaseBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeBridgeVerifierSet)
				if err := _BaseBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*BaseBridgeBridgeVerifierSet, error) {
	event := new(BaseBridgeBridgeVerifierSet)
	if err := _BaseBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the BaseBridge contract.
type BaseBridgeChainPauseSetIterator struct {
	Event *BaseBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeChainPauseSet)
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
		it.Event = new(BaseBridgeChainPauseSet)
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
func (it *BaseBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeChainPauseSet represents a ChainPauseSet event raised by the BaseBridge contract.
type BaseBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BaseBridge *BaseBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BaseBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeChainPauseSetIterator{contract: _BaseBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BaseBridge *BaseBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeChainPauseSet)
				if err := _BaseBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseChainPauseSet(log types.Log) (*BaseBridgeChainPauseSet, error) {
	event := new(BaseBridgeChainPauseSet)
	if err := _BaseBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the BaseBridge contract.
type BaseBridgeCrossMintableERC20CodeSetIterator struct {
	Event *BaseBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(BaseBridgeCrossMintableERC20CodeSet)
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
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the BaseBridge contract.
type BaseBridgeCrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BaseBridge *BaseBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*BaseBridgeCrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeCrossMintableERC20CodeSetIterator{contract: _BaseBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BaseBridge *BaseBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeCrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeCrossMintableERC20CodeSet)
				if err := _BaseBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*BaseBridgeCrossMintableERC20CodeSet, error) {
	event := new(BaseBridgeCrossMintableERC20CodeSet)
	if err := _BaseBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeDevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the BaseBridge contract.
type BaseBridgeDevSetIterator struct {
	Event *BaseBridgeDevSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeDevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeDevSet)
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
		it.Event = new(BaseBridgeDevSet)
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
func (it *BaseBridgeDevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeDevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeDevSet represents a DevSet event raised by the BaseBridge contract.
type BaseBridgeDevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BaseBridge *BaseBridgeFilterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*BaseBridgeDevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeDevSetIterator{contract: _BaseBridge.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BaseBridge *BaseBridgeFilterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeDevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeDevSet)
				if err := _BaseBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseDevSet(log types.Log) (*BaseBridgeDevSet, error) {
	event := new(BaseBridgeDevSet)
	if err := _BaseBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BaseBridge contract.
type BaseBridgeEIP712DomainChangedIterator struct {
	Event *BaseBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeEIP712DomainChanged)
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
		it.Event = new(BaseBridgeEIP712DomainChanged)
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
func (it *BaseBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the BaseBridge contract.
type BaseBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BaseBridge *BaseBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BaseBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeEIP712DomainChangedIterator{contract: _BaseBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BaseBridge *BaseBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeEIP712DomainChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BaseBridgeEIP712DomainChanged, error) {
	event := new(BaseBridgeEIP712DomainChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BaseBridge contract.
type BaseBridgeExtraCallExecutedIterator struct {
	Event *BaseBridgeExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *BaseBridgeExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeExtraCallExecuted)
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
		it.Event = new(BaseBridgeExtraCallExecuted)
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
func (it *BaseBridgeExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeExtraCallExecuted represents a ExtraCallExecuted event raised by the BaseBridge contract.
type BaseBridgeExtraCallExecuted struct {
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
func (_BaseBridge *BaseBridgeFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (*BaseBridgeExtraCallExecutedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeExtraCallExecutedIterator{contract: _BaseBridge.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_BaseBridge *BaseBridgeFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BaseBridgeExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeExtraCallExecuted)
				if err := _BaseBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseExtraCallExecuted(log types.Log) (*BaseBridgeExtraCallExecuted, error) {
	event := new(BaseBridgeExtraCallExecuted)
	if err := _BaseBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BaseBridge contract.
type BaseBridgeInitializedIterator struct {
	Event *BaseBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BaseBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeInitialized)
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
		it.Event = new(BaseBridgeInitialized)
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
func (it *BaseBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeInitialized represents a Initialized event raised by the BaseBridge contract.
type BaseBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BaseBridge *BaseBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BaseBridgeInitializedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeInitializedIterator{contract: _BaseBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BaseBridge *BaseBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BaseBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeInitialized)
				if err := _BaseBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseInitialized(log types.Log) (*BaseBridgeInitialized, error) {
	event := new(BaseBridgeInitialized)
	if err := _BaseBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the BaseBridge contract.
type BaseBridgeManualReleasedIterator struct {
	Event *BaseBridgeManualReleased // Event containing the contract specifics and raw log

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
func (it *BaseBridgeManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeManualReleased)
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
		it.Event = new(BaseBridgeManualReleased)
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
func (it *BaseBridgeManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeManualReleased represents a ManualReleased event raised by the BaseBridge contract.
type BaseBridgeManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BaseBridge *BaseBridgeFilterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BaseBridgeManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeManualReleasedIterator{contract: _BaseBridge.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BaseBridge *BaseBridgeFilterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *BaseBridgeManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeManualReleased)
				if err := _BaseBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseManualReleased(log types.Log) (*BaseBridgeManualReleased, error) {
	event := new(BaseBridgeManualReleased)
	if err := _BaseBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeMaxExtraDataLengthSetIterator is returned from FilterMaxExtraDataLengthSet and is used to iterate over the raw logs and unpacked data for MaxExtraDataLengthSet events raised by the BaseBridge contract.
type BaseBridgeMaxExtraDataLengthSetIterator struct {
	Event *BaseBridgeMaxExtraDataLengthSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeMaxExtraDataLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeMaxExtraDataLengthSet)
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
		it.Event = new(BaseBridgeMaxExtraDataLengthSet)
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
func (it *BaseBridgeMaxExtraDataLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeMaxExtraDataLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeMaxExtraDataLengthSet represents a MaxExtraDataLengthSet event raised by the BaseBridge contract.
type BaseBridgeMaxExtraDataLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxExtraDataLengthSet is a free log retrieval operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_BaseBridge *BaseBridgeFilterer) FilterMaxExtraDataLengthSet(opts *bind.FilterOpts) (*BaseBridgeMaxExtraDataLengthSetIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeMaxExtraDataLengthSetIterator{contract: _BaseBridge.contract, event: "MaxExtraDataLengthSet", logs: logs, sub: sub}, nil
}

// WatchMaxExtraDataLengthSet is a free log subscription operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_BaseBridge *BaseBridgeFilterer) WatchMaxExtraDataLengthSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeMaxExtraDataLengthSet) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeMaxExtraDataLengthSet)
				if err := _BaseBridge.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseMaxExtraDataLengthSet(log types.Log) (*BaseBridgeMaxExtraDataLengthSet, error) {
	event := new(BaseBridgeMaxExtraDataLengthSet)
	if err := _BaseBridge.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BaseBridge contract.
type BaseBridgePausedIterator struct {
	Event *BaseBridgePaused // Event containing the contract specifics and raw log

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
func (it *BaseBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgePaused)
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
		it.Event = new(BaseBridgePaused)
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
func (it *BaseBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgePaused represents a Paused event raised by the BaseBridge contract.
type BaseBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseBridge *BaseBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BaseBridgePausedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BaseBridgePausedIterator{contract: _BaseBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BaseBridge *BaseBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BaseBridgePaused) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgePaused)
				if err := _BaseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParsePaused(log types.Log) (*BaseBridgePaused, error) {
	event := new(BaseBridgePaused)
	if err := _BaseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgePendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the BaseBridge contract.
type BaseBridgePendingRemovedIterator struct {
	Event *BaseBridgePendingRemoved // Event containing the contract specifics and raw log

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
func (it *BaseBridgePendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgePendingRemoved)
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
		it.Event = new(BaseBridgePendingRemoved)
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
func (it *BaseBridgePendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgePendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgePendingRemoved represents a PendingRemoved event raised by the BaseBridge contract.
type BaseBridgePendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BaseBridge *BaseBridgeFilterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BaseBridgePendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgePendingRemovedIterator{contract: _BaseBridge.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BaseBridge *BaseBridgeFilterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *BaseBridgePendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgePendingRemoved)
				if err := _BaseBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParsePendingRemoved(log types.Log) (*BaseBridgePendingRemoved, error) {
	event := new(BaseBridgePendingRemoved)
	if err := _BaseBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BaseBridge contract.
type BaseBridgeRoleAdminChangedIterator struct {
	Event *BaseBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleAdminChanged)
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
		it.Event = new(BaseBridgeRoleAdminChanged)
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
func (it *BaseBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the BaseBridge contract.
type BaseBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BaseBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleAdminChangedIterator{contract: _BaseBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleAdminChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BaseBridgeRoleAdminChanged, error) {
	event := new(BaseBridgeRoleAdminChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BaseBridge contract.
type BaseBridgeRoleGrantedIterator struct {
	Event *BaseBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleGranted)
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
		it.Event = new(BaseBridgeRoleGranted)
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
func (it *BaseBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleGranted represents a RoleGranted event raised by the BaseBridge contract.
type BaseBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BaseBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleGrantedIterator{contract: _BaseBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleGranted)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseRoleGranted(log types.Log) (*BaseBridgeRoleGranted, error) {
	event := new(BaseBridgeRoleGranted)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BaseBridge contract.
type BaseBridgeRoleRevokedIterator struct {
	Event *BaseBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BaseBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeRoleRevoked)
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
		it.Event = new(BaseBridgeRoleRevoked)
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
func (it *BaseBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeRoleRevoked represents a RoleRevoked event raised by the BaseBridge contract.
type BaseBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BaseBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeRoleRevokedIterator{contract: _BaseBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BaseBridge *BaseBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BaseBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeRoleRevoked)
				if err := _BaseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseRoleRevoked(log types.Log) (*BaseBridgeRoleRevoked, error) {
	event := new(BaseBridgeRoleRevoked)
	if err := _BaseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BaseBridge contract.
type BaseBridgeThresholdChangedIterator struct {
	Event *BaseBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BaseBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeThresholdChanged)
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
		it.Event = new(BaseBridgeThresholdChanged)
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
func (it *BaseBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeThresholdChanged represents a ThresholdChanged event raised by the BaseBridge contract.
type BaseBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BaseBridge *BaseBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BaseBridgeThresholdChangedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeThresholdChangedIterator{contract: _BaseBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BaseBridge *BaseBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BaseBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeThresholdChanged)
				if err := _BaseBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseThresholdChanged(log types.Log) (*BaseBridgeThresholdChanged, error) {
	event := new(BaseBridgeThresholdChanged)
	if err := _BaseBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BaseBridge contract.
type BaseBridgeTokenPairRegisteredIterator struct {
	Event *BaseBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *BaseBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeTokenPairRegistered)
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
		it.Event = new(BaseBridgeTokenPairRegistered)
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
func (it *BaseBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the BaseBridge contract.
type BaseBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BaseBridge *BaseBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BaseBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTokenPairRegisteredIterator{contract: _BaseBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BaseBridge *BaseBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BaseBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeTokenPairRegistered)
				if err := _BaseBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*BaseBridgeTokenPairRegistered, error) {
	event := new(BaseBridgeTokenPairRegistered)
	if err := _BaseBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the BaseBridge contract.
type BaseBridgeTokenPauseSetIterator struct {
	Event *BaseBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeTokenPauseSet)
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
		it.Event = new(BaseBridgeTokenPauseSet)
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
func (it *BaseBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeTokenPauseSet represents a TokenPauseSet event raised by the BaseBridge contract.
type BaseBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BaseBridge *BaseBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BaseBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeTokenPauseSetIterator{contract: _BaseBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BaseBridge *BaseBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeTokenPauseSet)
				if err := _BaseBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseTokenPauseSet(log types.Log) (*BaseBridgeTokenPauseSet, error) {
	event := new(BaseBridgeTokenPauseSet)
	if err := _BaseBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BaseBridge contract.
type BaseBridgeUnpausedIterator struct {
	Event *BaseBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BaseBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeUnpaused)
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
		it.Event = new(BaseBridgeUnpaused)
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
func (it *BaseBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeUnpaused represents a Unpaused event raised by the BaseBridge contract.
type BaseBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseBridge *BaseBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BaseBridgeUnpausedIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeUnpausedIterator{contract: _BaseBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BaseBridge *BaseBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BaseBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeUnpaused)
				if err := _BaseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseUnpaused(log types.Log) (*BaseBridgeUnpaused, error) {
	event := new(BaseBridgeUnpaused)
	if err := _BaseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BaseBridge contract.
type BaseBridgeUpgradedIterator struct {
	Event *BaseBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BaseBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeUpgraded)
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
		it.Event = new(BaseBridgeUpgraded)
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
func (it *BaseBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeUpgraded represents a Upgraded event raised by the BaseBridge contract.
type BaseBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseBridge *BaseBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BaseBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeUpgradedIterator{contract: _BaseBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BaseBridge *BaseBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BaseBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeUpgraded)
				if err := _BaseBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseUpgraded(log types.Log) (*BaseBridgeUpgraded, error) {
	event := new(BaseBridgeUpgraded)
	if err := _BaseBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the BaseBridge contract.
type BaseBridgeVerificationDelayExpirationSetIterator struct {
	Event *BaseBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeVerificationDelayExpirationSet)
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
		it.Event = new(BaseBridgeVerificationDelayExpirationSet)
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
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the BaseBridge contract.
type BaseBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_BaseBridge *BaseBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BaseBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BaseBridgeVerificationDelayExpirationSetIterator{contract: _BaseBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_BaseBridge *BaseBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *BaseBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeVerificationDelayExpirationSet)
				if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*BaseBridgeVerificationDelayExpirationSet, error) {
	event := new(BaseBridgeVerificationDelayExpirationSet)
	if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the BaseBridge contract.
type BaseBridgeVerificationDelaySetIterator struct {
	Event *BaseBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *BaseBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgeVerificationDelaySet)
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
		it.Event = new(BaseBridgeVerificationDelaySet)
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
func (it *BaseBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the BaseBridge contract.
type BaseBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*BaseBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &BaseBridgeVerificationDelaySetIterator{contract: _BaseBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BaseBridge *BaseBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *BaseBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgeVerificationDelaySet)
				if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*BaseBridgeVerificationDelaySet, error) {
	event := new(BaseBridgeVerificationDelaySet)
	if err := _BaseBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
