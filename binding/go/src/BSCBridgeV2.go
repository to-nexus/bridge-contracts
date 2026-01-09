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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deadWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alreadyTransferred\",\"type\":\"bool\"}],\"name\":\"burnCrossToDeadWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross_\",\"type\":\"address\"}],\"name\":\"reinitializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeInvalidDeadWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
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
		"230f4cda": "postCallGasReserve()",
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
		"4bc51672": "setPostCallGasReserve(uint256)",
		"bfbf6765": "setTokenPause(uint256,address,bool,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f516020615f285f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615e6190816100c78239608051818181610c2c0152610de50152f35b6001600160401b0319166001600160401b039081175f516020615f285f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146103945780630b1d8d061461038f5780631313996b1461038a5780631e24abdb14610385578063230f4cda14610380578063248a9ca31461037b5780632f2ff15d1461037657806336568abe14610371578063365d71e41461036c57806342cde4e81461036757806348a00ed8146103625780634bc516721461035d5780634be13f83146103585780634ee078ba146103535780634f1ef2861461034e578063502cc5c01461034957806352d1902d1461034457806356ff54b01461033f5780635b605f5c1461033a5780635c975abb146103355780635fd262de14610330578063670e62681461032b578063761fe2d81461032657806379214874146103215780637b429ad81461031c578063814914b51461031757806384b0196e1461031257806386e9e1751461030d57806388d67d6d1461030857806389232a00146103035780638ae87c5c146102fe57806391cca3db146102f957806391cf6d3e146102f457806391d14854146102ef578063a10bab0b146102ea578063a217fddf146102e5578063a3246ad3146102e0578063aa1bd0bc146102db578063aa20ed47146102d6578063ad3cb1cc146102d1578063ae6893f8146102cc578063b2b49e2e146102c7578063b33eb36e146102c2578063b7f3358d146102bd578063b8aa837e146102b8578063bedb86fb146102b3578063bfbf6765146102ae578063cba25e4b146102a9578063cf56118e146102a4578063d477f05f1461029f578063d547741f1461029a578063d5717fc514610295578063e2af6cd714610290578063edbbea391461028b578063f0702e8e146102865763f698da2514610281575f80fd5b61248f565b612467565b612318565b61229f565b612266565b612237565b6121d8565b612164565b6120f9565b612020565b611f36565b611e9a565b611e16565b611d85565b611c77565b611c3e565b611bf7565b611b6c565b611b1f565b611aa3565b611a47565b611a1f565b6119dd565b6119c0565b611998565b61192f565b611833565b6116cc565b611515565b611461565b611314565b61123e565b6111ce565b611145565b6110ce565b611020565b610ff2565b610f0f565b610e2a565b610dd3565b610d43565b610bd9565b610a6b565b610a44565b6109d6565b6108ab565b61087f565b610816565b6105d7565b61059c565b61057e565b61054e565b610531565b6104ad565b61040b565b346103ea5760203660031901126103ea5760043563ffffffff60e01b81168091036103ea57602090637965db0b60e01b81149081156103d9575b506040519015158152f35b6301ffc9a760e01b1490505f6103ce565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103ea57565b346103ea5760203660031901126103ea57600435610428816103fa565b61043133613a1d565b6001600160a01b03166104458115156124fb565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103ea578235916001600160401b0383116103ea576020808501948460051b0101116103ea57565b60403660031901126103ea576004356001600160401b0381116103ea576104d890369060040161047d565b602435916001600160401b0383116103ea57366023840112156103ea576004830135916001600160401b0383116103ea5736602460e08502860101116103ea576024610525940191612530565b005b5f9103126103ea57565b346103ea575f3660031901126103ea576020600a54604051908152f35b346103ea575f3660031901126103ea576036548061057657506020620249f05b604051908152f35b60209061056e565b346103ea5760203660031901126103ea57602061056e600435612669565b346103ea5760403660031901126103ea576105256024356004356105bf826103fa565b6105d26105cb82612669565b3390613c3d565b6132a3565b346103ea5760403660031901126103ea576004356024356105f7816103fa565b336001600160a01b038216036106105761052591613303565b63334bd91960e11b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b0382111761064e57604052565b61061f565b606081019081106001600160401b0382111761064e57604052565b60c081019081106001600160401b0382111761064e57604052565b601f909101601f19168101906001600160401b0382119082101761064e57604052565b604051906106bb608083610689565b565b604051906106bb60e083610689565b604051906106bb61010083610689565b604051906106bb60c083610689565b604051906106bb606083610689565b6001600160401b03811161064e5760051b60200190565b9080601f830112156103ea578135610728816106fa565b926107366040519485610689565b81845260208085019260051b8201019283116103ea57602001905b82821061075e5750505090565b8135815260209182019101610751565b9060406003198301126103ea576004356001600160401b0381116103ea578261079991600401610711565b91602435906001600160401b0382116103ea57806023830112156103ea5781600401356107c5816106fa565b926107d36040519485610689565b8184526024602085019260051b8201019283116103ea57602401905b8282106107fc5750505090565b60208091833561080b816103fa565b8152019101906107ef565b346103ea576108243661076e565b906108328151835114612687565b5f5b8251811015610525578061087861084d6001938561269d565b51838060a01b0361085e848861269d565b5116906108733361086e83612669565b613c3d565b613303565b5001610834565b346103ea575f3660031901126103ea57602060ff5f516020615b4c5f395f51905f525416604051908152f35b346103ea5760603660031901126103ea576024356004356044356108ce816103fa565b6108d733613aa5565b6108df612fb6565b815f5260076020526108fd836108f88160405f20614676565b6126b1565b8061090884846143c1565b916001600160a01b0316156109c2575b8151905f516020615b6c5f395f51905f52602084019161097b83519561097161095e6040830197885160018060a01b03169986608086019b8c519260a0880151946147e4565b5061096881611cea565b60018114613362565b51935194516103ee565b94516040516001600160a01b039096169591829161099c9142919084613385565b0390a45f516020615b0c5f395f51905f525f80a35f5f516020615d0c5f395f51905f525d005b5060608101516001600160a01b0316610918565b346103ea5760203660031901126103ea576004356109f333613a1d565b61c3508110610a355760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91603654908060365582519182526020820152a1005b63050779f960e21b5f5260045ffd5b346103ea575f3660031901126103ea575f546040516001600160a01b039091168152602090f35b346103ea5760403660031901126103ea57610b76602435600435610a8d612f8f565b610a95612fb6565b610abb81610ab6610ab26003610aaa8461264d565b015460ff1690565b1590565b6126cb565b610ada826108f881610ad5855f52600760205260405f2090565b614676565b610b716040610b0a610b0585610af8865f52600860205260405f2090565b905f5260205260405f2090565b612830565b610b5e610b2a82516080610b20868301516103ee565b910151908761346c565b50610b3481611cea565b610b3d81611cea565b83516020808201839052815290600190610b58604084610689565b14612869565b01518015908115610b7e575b4291612899565b6133a3565b610525612feb565b4281109150610b6a565b6001600160401b03811161064e57601f01601f191660200190565b929192610baf82610b88565b91610bbd6040519384610689565b8294818452818301116103ea578281602093845f960137010152565b60403660031901126103ea57600435610bf1816103fa565b6024356001600160401b0381116103ea57366023820112156103ea57610c21903690602481600401359101610ba3565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610d21575b50610d1257610c6533613a1d565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610ce1575b50610cae57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615c0c5f395f51905f528303610ccd576105259250614c6a565b632a87526960e21b5f52600483905260245ffd5b610d0491945060203d602011610d0b575b610cfc8183610689565b810190613593565b925f610c8d565b503d610cf2565b63703e46dd60e11b5f5260045ffd5b5f516020615c0c5f395f51905f52546001600160a01b0316141590505f610c57565b346103ea5760603660031901126103ea57602435600435604435610d6633613aa5565b815f526007602052610d7f836108f88160405f20614676565b4201804211610dce5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b6128b7565b346103ea575f3660031901126103ea577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610d125760206040515f516020615c0c5f395f51905f528152f35b346103ea5760203660031901126103ea577f1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f16020600435610e6a33613b2d565b80600a55604051908152a1005b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610eef5750505090565b909192602060e082610f046001948851610e77565b019401929101610ee2565b346103ea5760203660031901126103ea57600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b818110610fd9575050610f5e92500383610689565b610f68825161290e565b915f5b8151811015610fc757600190610fab610fa6610f868661265b565b610fa0610f93858861269d565b516001600160a01b031690565b9061295d565b612972565b610fb5828761269d565b52610fc0818661269d565b5001610f6b565b60405180610fd58682610ecc565b0390f35b8454835260019485019487945060209093019201610f49565b346103ea575f3660031901126103ea57602060ff5f516020615ccc5f395f51905f5254166040519015158152f35b60e03660031901126103ea5760243560043561103b826103fa565b604435611047816103fa565b60c4359160a4356084356064356001600160401b0386116103ea57366023870112156103ea576004860135946001600160401b0386116103ea5736602487890101116103ea57610fd597602461109e9801956129e5565b60405191829182901515815260200190565b60ff8116036103ea57565b6001600160a01b03909116815260200190565b346103ea5760803660031901126103ea576024356004356110ee826103fa565b604435906001600160401b0382116103ea57366023830112156103ea57610fd592611126611139933690602481600401359101610ba3565b9060643592611134846110b0565b612aeb565b604051918291826110bb565b346103ea5760403660031901126103ea57602060ff61117b60243560043561116c826103fa565b5f526009845260405f2061295d565b54166040519015158152f35b90602080835192838152019201905f5b8181106111a45750505090565b8251845260209384019390920191600101611197565b9060206111cb928181520190611187565b90565b346103ea5760203660031901126103ea576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061122857610fd58561121c81870382610689565b604051918291826111ba565b8254845260209093019260019283019201611205565b346103ea5760403660031901126103ea5760043560243561125e816103fa565b61126733613a1d565b5f516020615d8c5f395f51905f52549160ff8360401c1680156112f0575b6112e1576001600160401b03199092166002175f516020615d8c5f395f51905f52556112b8916112b3612bb9565b612c1c565b6112c0612bda565b604051600281525f516020615b8c5f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b0384161015611285565b60e0810192916106bb9190610e77565b346103ea5760403660031901126103ea57610fd561135360243560043561133a826103fa565b6113426128d8565b505f52600660205260405f2061295d565b60046040519161136283610633565b80546001600160a01b0316835260018101546113b7906113ae90611391611388826103ee565b602088016126e5565b6113a560a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611304565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b91611437906114296111cb97959693600f60f81b865260e0602087015260e08601906113de565b9084820360408601526113de565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611187565b346103ea575f3660031901126103ea575f516020615bcc5f395f51905f525415806114f5575b156114b8576114946138c9565b61149c613983565b90610fd56114a8612cf2565b6040519384933091469186611402565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615dac5f395f51905f525415611487565b801515036103ea57565b346103ea5760603660031901126103ea57600435611532816103fa565b6024356044356115418161150b565b611549612f8f565b611551612fb6565b6001600160a01b0383165f90815260666020526040902061157d90611578905b5460ff1690565b612d0d565b611588821515612bf5565b61163f576115a3818361159c6065546103ee565b3390613c89565b335b5f516020615e0c5f395f51905f526115be606454613cd2565b92606454926116286115d16065546103ee565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a4611633612feb565b60405160018152602090f35b61164833613a1d565b306115a5565b9080601f830112156103ea578135611665816106fa565b926116736040519485610689565b81845260208085019260051b820101918383116103ea5760208201905b83821061169f57505050505090565b81356001600160401b0381116103ea576020916116c187848094880101610711565b815201910190611690565b60803660031901126103ea576004356001600160401b0381116103ea576116f790369060040161047d565b90602435906001600160401b0382116103ea57366023830112156103ea57816004013591611724836106fa565b926117326040519485610689565b8084526024602085019160051b830101913683116103ea5760248101915b8383106117a85750506044359150506001600160401b0381116103ea5761177b90369060040161164e565b606435926001600160401b0384116103ea57610fd5946117a261109e95369060040161164e565b93612d23565b82356001600160401b0381116103ea578201366043820112156103ea576024810135906117d4826106fa565b916117e26040519384610689565b808352602060248185019260051b84010101913683116103ea57604401905b82821061181957505050815260209283019201611750565b602080918335611828816110b0565b815201910190611801565b346103ea5760603660031901126103ea57600435611850816103fa565b6024359061185d826103fa565b604435611869816110b0565b5f516020615d8c5f395f51905f52549260ff604085901c1615936001600160401b031680159081611927575b600114908161191d575b159081611914575b506112e1575f516020615d8c5f395f51905f5280546001600160401b03191660011790556118d992846119075761401d565b6118df57005b6118e7612bda565b604051600181525f516020615b8c5f395f51905f529080602081016112dc565b61190f612bb9565b61401d565b9050155f6118a7565b303b15915061189f565b859150611895565b346103ea5760403660031901126103ea5760243560043561194f33613a1d565b611957612fb6565b61196182826143c1565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615d0c5f395f51905f525d005b346103ea575f3660031901126103ea576033546040516001600160a01b039091168152602090f35b346103ea575f3660031901126103ea576020603454604051908152f35b346103ea5760403660031901126103ea57602060ff61117b602435600435611a04826103fa565b5f525f516020615cac5f395f51905f52845260405f2061295d565b346103ea575f3660031901126103ea576035546040516001600160a01b039091168152602090f35b346103ea575f3660031901126103ea5760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611a845750505090565b82516001600160a01b0316845260209384019390920191600101611a77565b346103ea5760203660031901126103ea576004355f525f516020615c2c5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611b0957610fd585611afd81870382610689565b60405191829182611a61565b8254845260209093019260019283019201611ae6565b346103ea5760203660031901126103ea577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611b5f33613a1d565b80600155604051908152a1005b346103ea5760403660031901126103ea57602435600435611b8c33613aa5565b611b9533613aa5565b611b9d612fb6565b805f526007602052611bb6826108f88160405f20614676565b611bc082826133a3565b5f516020615b0c5f395f51905f525f80a35f5f516020615d0c5f395f51905f525d005b60405190611bf2602083610689565b5f8252565b346103ea575f3660031901126103ea57610fd5604051611c18604082610689565b60058152640352e302e360dc1b60208201526040519182916020835260208301906113de565b346103ea5760203660031901126103ea576004355f526004602052600160405f20015460018101809111610dce57602090604051908152f35b346103ea57611c853661076e565b90611c938151835114612687565b5f5b82518110156105255780611ccf611cae6001938561269d565b51838060a01b03611cbf848861269d565b5116906105d23361086e83612669565b5001611c95565b634e487b7160e01b5f52602160045260245ffd5b600a1115611cf457565b611cd6565b90600a821015611cf45752565b6020815260606040611d6b60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c06101208601526101408501906113de565b93611d7d602082015183860190611cf9565b015191015290565b346103ea5760403660031901126103ea57600435602435905f60408051611dab81610653565b611db3612dfb565b815282602082015201525f52600860205260405f20905f52602052610fd560405f20600760405191611de483610653565b611ded8161272c565b8352611e0360ff60068301541660208501612824565b0154604082015260405191829182611d06565b346103ea5760203660031901126103ea577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611e58816110b0565b611e6133613a1d565b16611e6d811515612e2b565b8060ff195f516020615b4c5f395f51905f525416175f516020615b4c5f395f51905f5255604051908152a1005b346103ea5760403660031901126103ea57600435602435611eba8161150b565b611ec333613bb5565b611ed8825f52600360205260405f2054151590565b15611f235760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611f1881600360405f2001612c0b565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103ea5760203660031901126103ea57600435611f538161150b565b611f5c33613bb5565b15611fba57611f69612f8f565b600160ff195f516020615ccc5f395f51905f525416175f516020615ccc5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615ccc5f395f51905f525460ff8116156120115760ff19165f516020615ccc5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103ea5760803660031901126103ea57602435600435612040826103fa565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c604060443561206f8161150b565b60643561207b8161150b565b61208433613bb5565b845f5260056020526120e8816120e3855f20986120b3816120ae60018060a01b038216809d614676565b612e41565b885f5260066020526120d38660016120cd848b5f2061295d565b01612e69565b885f526009602052865f2061295d565b612c0b565b8251911515825215156020820152a3005b346103ea5760203660031901126103ea57600435612116816103fa565b61211f33613a1d565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103ea575f3660031901126103ea57604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106121c257610fd58561121c81870382610689565b82548452602090930192600192830192016121ab565b346103ea5760203660031901126103ea576004356121f5816103fa565b6121fe33613a1d565b6001600160a01b03166122128115156124fb565b603380546001600160a01b031916821790555f516020615c4c5f395f51905f525f80a2005b346103ea5760403660031901126103ea5761052560243560043561225a826103fa565b6108736105cb82612669565b346103ea5760203660031901126103ea576004355f526004602052600260405f20015460018101809111610dce57602090604051908152f35b346103ea5760203660031901126103ea576004356122bc816103fa565b6122c533613a1d565b6001600160a01b03166122d9811515612e86565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346103ea5760803660031901126103ea576004356024356123388161150b565b60443591612345836103fa565b5f516020615d4c5f395f51905f5261242f60643593612363856103fa565b61236c33613b2d565b612377841515612e86565b6001600160a01b0386169461109e90612391871515612e86565b6001600160a01b0381169761242a906123ab8a1515612e86565b6123b488615440565b612434575b6123dd816123d8816123d38c5f52600560205260405f2090565b6145db565b6144a4565b6123fc6123e86106bd565b936123f383866126e5565b602085016126e5565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526124258861265b565b61295d565b6144cc565b0390a4005b61246261243f6106ac565b8981525f60208201525f60408201525f606082015261245d8a61264d565b614479565b6123b9565b346103ea575f3660031901126103ea576032546040516001600160a01b039091168152602090f35b346103ea575f3660031901126103ea576124a7615831565b6124af615888565b6040519060208201925f516020615dcc5f395f51905f528452604083015260608201524660808201523060a082015260a081526124ed60c082610689565b519020604051908152602090f35b1561250257565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b9291909261253c612f8f565b612544612fb6565b61254f8385146125f7565b5f5b8481106125655750505050506106bb612feb565b612570818684612621565b3590602061257f828886612621565b013561258a816103fa565b6060612597838987612621565b0135926125a3846103fa565b60806125b0848a88612621565b013560a06125bf858b89612621565b013560c06125ce868c8a612621565b013591898610156125f2576001966125ec9560e088028b0195612ffd565b01612551565b61260d565b156125fe57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125f25760051b8101359060fe19813603018212156103ea570190565b356111cb816103fa565b5f52600460205260405f2090565b5f52600660205260405f2090565b5f525f516020615cac5f395f51905f52602052600160405f20015490565b1561268e57565b63031206d560e51b5f5260045ffd5b80518210156125f25760209160051b010190565b156126b95750565b6373a1390160e11b5f5260045260245ffd5b156126d35750565b636fc24b7960e11b5f5260045260245ffd5b6001600160a01b039091169052565b90600182811c92168015612722575b602083101461270e57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612703565b906040516127398161066e565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91612791826126f4565b80855291600181169081156127fd57506001146127be575b505060a092916127ba910384610689565b0152565b5f908152602081209092505b8183106127e15750508101602001816127ba6127a9565b60209193508060019154838589010152019101909184926127ca565b60ff191660208681019190915292151560051b850190920192508391506127ba90506127a9565b600a821015611cf45752565b9060405161283d81610653565b60406007829461284c8161272c565b845261286260ff60068301541660208601612824565b0154910152565b156128715750565b60405162461bcd60e51b8152602060048201529081906128959060248301906113de565b0390fd5b156128a2575050565b637a88099160e11b5f5260045260245260445ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610dce57565b604051906128e582610633565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b90612918826106fa565b6129256040519182610689565b8281528092612936601f19916106fa565b01905f5b82811061294657505050565b6020906129516128d8565b8282850101520161293a565b9060018060a01b03165f5260205260405f2090565b9060405161297f81610633565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916129ca906129c1906113a5565b15156060860152565b60028101546080850152600381015460a08501520154910152565b959392612a7d612a6793612a50612a9398612a9d9b97612a03612f8f565b612a166001600160a01b0388168d6135a2565b612a1e612fb6565b612a326001600160a01b03851615156124fb565b612a48600a548015908115612aaa575b50612ab5565b85878d613657565b959094612a5b6106cc565b9a8b5260208b016126e5565b612a743360408b016126e5565b606089016126e5565b608087015260a086015260c08501523691610ba3565b60e08201526137f2565b612aa5612feb565b600190565b90508a11155f612a42565b15612abc57565b633953b33f60e11b5f5260045ffd5b908160209103126103ea57516111cb816103fa565b6040513d5f823e3d90fd5b5f5490949392906001600160a01b03811615612baa57604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90612b549060848701906113de565b931691015203925af18015612ba5576106bb925f91612b76575b508094612e9c565b612b98915060203d602011612b9e575b612b908183610689565b810190612acb565b5f612b6e565b503d612b86565b612ae0565b6315aeca0d60e11b5f5260045ffd5b5f516020615d8c5f395f51905f52805460ff60401b1916600160401b179055565b5f516020615d8c5f395f51905f52805460ff60401b19169055565b15612bfc57565b63ff5a231360e01b5f5260045ffd5b9060ff801983541691151516179055565b90612c28821515612bf5565b6001600160a01b0316908115612ce357606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b60405190612d01602083610689565b5f808352366020840137565b15612d1457565b635e0c1f8960e01b5f5260045ffd5b919294939094612d31612f8f565b612d39612fb6565b8584511480612dd3575b80612dc9575b612d55909695966125f7565b612d60343415612ddd565b3683900360be1901955f5b86811015612db9578060051b85013590888212156103ea57612db3600192612d93838a61269d565b51612d9e848861269d565b5190612daa858a61269d565b51928a01613d8d565b01612d6b565b5095505050505050612aa5612feb565b5081518614612d49565b5085815114612d43565b15612de55750565b63943892eb60e01b5f525f60045260245260445ffd5b60405190612e088261066e565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b15612e3257565b63f0f15b9160e01b5f5260045ffd5b15612e495750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b15612e8d57565b6302bfb33d60e51b5f5260045ffd5b919091612ea833613b2d565b612eb3811515612e86565b6001600160a01b03831691612f4190612ecd841515612e86565b6001600160a01b0381169461242a90612ee7871515612e86565b612ef085615440565b612f61575b612f0f816123d8816123d3895f52600560205260405f2090565b612f1a6123e86106bd565b5f60408401525f60608401525f60808401525f60a08401525f60c08401526124258561265b565b6040515f81525f516020615d4c5f395f51905f529080602081015b0390a4565b612f8a612f6c6106ac565b8681525f60208201525f60408201525f606082015261245d8761264d565b612ef5565b60ff5f516020615ccc5f395f51905f525416612fa757565b63d93c066560e01b5f5260045ffd5b5f516020615d0c5f395f51905f525c612fdc5760015f516020615d0c5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615d0c5f395f51905f525d565b9491939291613072906130466001600160a01b03841661301d818a6135a2565b61303161302c61302c8c612643565b6103ee565b81149061304061302c8c612643565b916131bf565b602088019461306a61305a61302c88612643565b6001600160a01b038916146131ec565b848489613657565b94909360408801356130a8816130918961308c8a8a6128cb565b6128cb565b8110156130a28a61308c8b8b6128cb565b90613202565b6130b661302c6032546103ee565b906130c083612643565b60608b01359a6130d260808201613220565b9060c08101359060a00135853b156103ea57604051637f40ec1760e11b81526001600160a01b038a811660048301529490941660248501523060448501526064840194909452608483019c909c5260ff1660a482015260c481019190915260e481019990995288610104815a5f948591f1928315612ba55761317c613167613185936106bb9b61318e976131a5575b50612643565b916131706106cc565b998a5260208a016126e5565b604088016126e5565b606086016126e5565b608084015260a083015260c0820152612a93611be3565b806131b35f6131b993610689565b80610527565b5f613161565b156131c8575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b156131f357565b630672aec160e01b5f5260045ffd5b1561320b575050565b63943892eb60e01b5f5260045260245260445ffd5b356111cb816110b0565b90613235825f61454b565b918261323e5750565b5f80525f516020615c2c5f395f51905f526020526001600160a01b0316613285817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d6154a9565b1561328d5750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916132b0838261454b565b92836132ba575050565b815f525f516020615c2c5f395f51905f526020526132e560405f209160018060a01b031680926154a9565b156132ee575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161331083826145ee565b928361331a575050565b815f525f516020615c2c5f395f51905f5260205261334560405f209160018060a01b0316809261552b565b1561334e575050565b62a95f1b60e31b5f5260045260245260445ffd5b1561336a5750565b63290cd55f60e01b5f52600a811015611cf45760045260245ffd5b604091949392606082019560018060a01b0316825260208201520152565b906133ad916143c1565b60018060a01b036060820151168151905f516020615b6c5f395f51905f5260208401918251946133fd61095e6040830196875160018060a01b03169885608086019a8b519260a0880151946147e4565b519251935194516040516001600160a01b0390961695918291612f5c9142919084613385565b1561342a57565b6330d45fb160e01b5f5260045ffd5b908160209103126103ea5751600a8110156103ea5790565b6001600160a01b039091168152602081019190915260400190565b91506134ad60ff9161349c5f9461348e60325460018060a01b03161515613423565b5f52600960205260405f2090565b6001600160a01b039091169061295d565b54166134b857600191565b506002905f90565b9190915f926135016115716134f16134dc61302c6032546103ee565b9461348e6001600160a01b0387161515613423565b6001600160a01b0384169061295d565b613588579160209161352b95935f604051809881958294633f4bdec560e01b845260048401613451565b03925af1928315612ba5575f93613557575b5060018361354a81611cea565b0361355157565b60019150565b61357a91935060203d602011613581575b6135728183610689565b810190613439565b915f61353d565b503d613568565b505050506002905f90565b908160209103126103ea575190565b805f5260056020526135c4826120ae60405f2060018060a01b03831690614676565b805f5260046020526135e18160ff600360405f20015416156126cb565b5f52600660205260ff60016135f98360405f2061295d565b015460a81c166136065750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b908160609103126103ea578051916040602083015192015190565b1561364857565b6358e8878560e01b5f5260045ffd5b826060916136e097959693613671610fa66134f18461265b565b613681610ab26040830151151590565b613782575b5061369561302c6032546103ee565b916136aa6001600160a01b0384161515613423565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612ba5575f955f905f93613744575b509082916106bb94939681988515159586613739575b50508461372e575b505082613723575b5050613641565b101590505f8061371c565b101592505f80613714565b101594505f8061370c565b905061376f9196506106bb93925060603d60601161377b575b6137678183610689565b810190613626565b919692939192916136f6565b503d61375d565b60c06137949101518480821015613202565b5f613686565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916137ed91908601906113de565b930152565b6137fc8151613cd2565b613806825161265b565b5f516020615e0c5f395f51905f5261383861302c60016138316020880195610fa061302c88516103ee565b01546103ee565b93805190612f5c61384985516103ee565b91604081019061385982516103ee565b906138826080820196875160a084019485519861387c60c087019a8b51906128cb565b93614d56565b613898613891825199516103ee565b93516103ee565b9460e06138a860608401516103ee565b9751935191519201519260405197889760018060a01b03169c42968961379a565b604051905f825f516020615b2c5f395f51905f5254916138e8836126f4565b8083529260018116908115613964575060011461390c575b6106bb92500383610689565b505f516020615b2c5f395f51905f525f90815290915f516020615bec5f395f51905f525b8183106139485750509060206106bb92820101613900565b6020919350806001915483858901015201910190918492613930565b602092506106bb94915060ff191682840152151560051b820101613900565b604051905f825f516020615bac5f395f51905f5254916139a2836126f4565b808352926001811690811561396457506001146139c5576106bb92500383610689565b505f516020615bac5f395f51905f525f90815290915f516020615dec5f395f51905f525b818310613a015750509060206106bb92820101613900565b60209193508060019154838589010152019101909184926139e9565b5f516020615d2c5f395f51905f525f525f516020615cac5f395f51905f5260205260ff613a6a827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61295d565b541615613a745750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d2c5f395f51905f52602452604490fd5b5f516020615d6c5f395f51905f525f525f516020615cac5f395f51905f5260205260ff613af2827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361295d565b541615613afc5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d6c5f395f51905f52602452604490fd5b5f516020615cec5f395f51905f525f525f516020615cac5f395f51905f5260205260ff613b7a827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61295d565b541615613b845750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615cec5f395f51905f52602452604490fd5b5f516020615c6c5f395f51905f525f525f516020615cac5f395f51905f5260205260ff613c02827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661295d565b541615613c0c5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615c6c5f395f51905f52602452604490fd5b90815f525f516020615cac5f395f51905f5260205260ff613c618260405f2061295d565b541615613c6c575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526106bb91613ccd608483610689565b614e97565b5f526004602052600160405f2001600181540180915590565b15613cf35750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15613d1c575050565b63a6ab65c560e01b5f5260045260245260445ffd5b903590601e19813603018212156103ea57018035906001600160401b0382116103ea576020019181360383136103ea57565b6106bb93606092969593608083019760018060a01b03168352602083015260408201520190611cf9565b919282359182613d9c8161264d565b6003015460ff161590613dae916126cb565b613dc0835f52600560205260405f2090565b906040850191613dcf83612643565b613dd8906103ee565b613de191614466565b613dea83612643565b613df3906103ee565b613dfc91613ceb565b613e19845f526004602052600260405f2001600181540180915590565b946020810135958681811491613e2e92613d13565b613e3783612643565b613e40906103ee565b916060820194613e4f86612643565b988860808501359a60a08601968a8d613e688a8a613d31565b3690613e7392610ba3565b8051602091820120604080517fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b77219149381019384529081019490945260608401969096526001600160a01b0393841660808401529390921660a082015260c081019190915260e08082019390935291825290613eef61010082610689565b51902092613efc93614f1f565b5f9087613f0885612643565b90613f1391886134c0565b91909384613f2081611cea565b600114613fdb575b50613f3284611cea565b60018403613f7c575050505090613f5f613f595f516020615b6c5f395f51905f5293612643565b91612643565b6040516001600160a01b03909216958291612f5c91429184613385565b83949850613fb7612f5c93613fc3937f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896613fbd9461527c565b98612643565b92612643565b9260405193849360018060a01b031698429185613d63565b90935061401491925088613fee86612643565b9161400c614005613ffe8a612643565b9288613d31565b3691610ba3565b928a8a6147e4565b9190925f613f28565b9161402661539b565b61403a6001600160a01b03841615156124fb565b6001600160a01b038216926140508415156124fb565b60ff8216156142ce576140bc9061406561539b565b61406d61539b565b61407561539b565b60ff195f516020615ccc5f395f51905f5254165f516020615ccc5f395f51905f525561409f61539b565b6140a761539b565b6140af61539b565b6140b761539b565b61322a565b506140c561539b565b6140d360ff82161515612e2b565b604080516140e18282610689565b60098152682b30b634b230ba37b960b91b602082015261410382519283610689565b60058252640312e302e360dc1b602083015261411d61539b565b61412561539b565b8051906001600160401b03821161064e57614156826141515f516020615b2c5f395f51905f52546126f4565b6150a1565b602090601f831160011461423a579261418583614215979694614199946141eb975f9261422f575b505061431d565b5f516020615b2c5f395f51905f5255615a16565b6141ae5f5f516020615bcc5f395f51905f5255565b6141c35f5f516020615dac5f395f51905f5255565b60ff1660ff195f516020615b4c5f395f51905f525416175f516020615b4c5f395f51905f5255565b6141f36153c6565b603380546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615c4c5f395f51905f525f80a26106bb43603455565b015190505f8061417e565b5f516020615b2c5f395f51905f525f52601f19831691905f516020615bec5f395f51905f52925f5b8181106142b6575093614199936141eb9693600193836142159b9a981061429e575b505050811b015f516020615b2c5f395f51905f5255615a16565b01515f1960f88460031b161c191690555f8080614284565b92936020600181928786015181550195019301614262565b6338854f1160e21b5f5260045ffd5b91908203918211610dce57565b916143039183549060031b91821b915f19901b19161790565b9055565b818110614312575050565b5f8155600101614307565b8160011b915f199060031b1c19161790565b60075f918281558260018201558260028201558260038201558260048201556005810161435c81546126f4565b908161436f575b50508260068201550155565b601f821160011461438657849055505b5f80614363565b6143ac6143bc926001601f61439e855f5260205f2090565b920160051c82019101614307565b5f81815260208120918190559055565b61437f565b91906143cb612dfb565b50825f5260076020526143e18160405f2061552b565b156144545761444f6106bb91845f52600860205260405f20815f52602052610af861440e60405f2061272c565b9561442b61441b8261265b565b610fa061302c60408b01516103ee565b61443f600260808a015192019182546142dd565b90555f52600860205260405f2090565b61432f565b637f11bea960e01b5f5260045260245ffd5b6111cb916001600160a01b031690614676565b600360606106bb93805184556020810151600185015560408101516002850155015115159101612c0b565b156144ac5750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916144e560018060a01b0382511685612511565b6020810151614530906001860190614506906001600160a01b031682612511565b6040830151815460ff60a01b191690151560a01b60ff60a01b161781556060830151151590612e69565b6080810151600285015560a081015160038501550151910155565b805f525f516020615cac5f395f51905f5260205260ff61456e8360405f2061295d565b54166145d557805f525f516020615cac5f395f51905f526020526145958260405f2061295d565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b6111cb916001600160a01b0316906154a9565b805f525f516020615cac5f395f51905f5260205260ff6146118360405f2061295d565b5416156145d557805f525f516020615cac5f395f51905f526020526146398260405f2061295d565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b6001915f520160205260405f2054151590565b3d156146b3573d9061469a82610b88565b916146a86040519384610689565b82523d5f602084013e565b606090565b908160209103126103ea57516111cb8161150b565b156146d457565b632b60b36f60e21b5f5260045ffd5b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a082018190526111cb929101906113de565b91906040838203126103ea578251602084015190936001600160401b0382116103ea570181601f820112156103ea5780519061475d82610b88565b9261476b6040519485610689565b828452602083830101116103ea57815f9260208093018386015e8301015290565b6111cb939260809263ffffffff60e01b16825260016020830152604082015281606082015201906113de565b6080906111cb939263ffffffff60e01b1681525f60208201525f604082015281606082015201906113de565b938096959192936147f48661265b565b94614817600161480c818060a01b038416809961295d565b015460a01c60ff1690565b9361482661302c6035546103ee565b9560188251101580614c58575b80614c4e575b614875575b505061484a9450615617565b9161485483611cea565b600183146148625750509190565b61486e939492506155ce565b6001905f90565b602082015160601c966034830151905f5f808b6040516148b7816148a9602082019462483e5560e91b8652602483016110bb565b03601f198101835282610689565b5190855afa6148c4614689565b9080614c42575b614c36575b506148dd575b505061483e565b898b60018214958615809790614b0e575b6148fb575b5050506148d6565b865f928a86839f9b614958819e9f9b6148a98998829f9d9e60365480158c14614b095750620249f05b805a108c14614af957508a995b8b14614af2578a975b604051968795602087019b631eeed51360e01b8d52602488016146e3565b5193f191614964614689565b9a614a9f575b505080614a93575b614a47578b95949392918a9180614a3f575b6149ba575b61484a985f516020615c8c5f395f51905f52916149ab604051928392836147b8565b0390a45f808080898b826148f3565b956149de915060209060405180938192632770a7eb60e21b83523060048401613451565b03815f8d5af18015612ba55761484a988d978c935f516020615c8c5f395f51905f5293614a10575b5091509850614989565b614a319060203d602011614a38575b614a298183610689565b8101906146b8565b505f614a06565b503d614a1f565b508715614984565b505f516020615c8c5f395f51905f52929550614a8b919450889661486e9a9b99945080602080614a7c93518301019101614722565b6040939193519384938461478c565b0390a46155ce565b50604088511015614972565b60405163095ea7b360e01b602082019081526001600160a01b0390921660248201525f6044820181905292839290918390614add81606481016148a9565b51925af150614aea614689565b50895f61496a565b809761493a565b614b03905a6142dd565b99614931565b614924565b508a1580614bc8575b5f808c604051614b3d816148a9602082019463095ea7b360e01b86528c60248401613451565b519082885af190614b4c614689565b5081159081614bc0575b50156148ee579960209192505f93614b8591604051958680948193632770a7eb60e21b83523060048401613451565b03925af1908115612ba5578f998d938f93614ba1575b506148ee565b614bb99060203d602011614a3857614a298183610689565b505f614b9b565b90505f614b56565b9960209192505f93614bf1916040519586809481936340c10f1960e01b83523060048401613451565b03925af1908115612ba5578f99614c128f938f955f91614c17575b506146cd565b614b17565b614c30915060203d602011614a3857614a298183610689565b5f614c0c565b6020915001515f6148d0565b506020815110156148cb565b50863b1515614839565b506001600160a01b0387161515614833565b90813b15614ceb575f516020615c0c5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115614cd357614cd0916156d6565b50565b505034614cdc57565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b15614d1657505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b15614d4757565b63559d141b60e11b5f5260045ffd5b90936001600160a01b0385169260018403614dc457506106bb9450614d8c614d7e82866128cb565b341434906130a284886128cb565b80614d98575b506157d0565b614db9614dbe91614daa6033546103ee565b90614db3611be3565b91615771565b614d40565b5f614d92565b90614dd0343415612ddd565b614de5614ddd82876128cb565b308489613c89565b80614e79575b50614e01610ab2600161480c866124258761265b565b614e11575b506106bb93506157d0565b604051632770a7eb60e21b815260208180614e30883060048401613451565b03815f885af1918215612ba5576106bb96614e549387935f91614e5a575b50614d0c565b5f614e06565b614e73915060203d602011614a3857614a298183610689565b5f614e4e565b614e9190614e8b61302c6033546103ee565b876156f3565b5f614deb565b905f602091828151910182855af115612ae0575f513d614ee657506001600160a01b0381163b155b614ec65750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415614ebf565b15614ef657565b63b3f07a3960e01b5f5260045ffd5b15614f0d5750565b631aebd17960e11b5f5260045260245ffd5b939293815191835183148061507d575b614f3890614eef565b614f59614f535f516020615b4c5f395f51905f525460ff1690565b60ff1690565b95614f678488811015614f05565b5f945f935f5b868110614f8857505050505050506106bb9192811015614f05565b614fe5614fb483614f976153d7565b6042916040519161190160f01b8352600283015260228201522090565b614fc8614fc1848961269d565b5160ff1690565b614fd2848761269d565b5190614fde858961269d565b5192615819565b6001600160a01b038181169088161080615016575b615008575b50600101614f6d565b600198890198909650614fff565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615cac5f395f51905f52602052615078611571827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561295d565b614ffa565b5085518314614f2f565b1561508f5750565b6307a5c91d60e31b5f5260045260245ffd5b90601f82116150ae575050565b6106bb915f516020615b2c5f395f51905f525f5260205f20906020601f840160051c830193106150e6575b601f0160051c0190614307565b90915081906150d9565b9190601f81116150ff57505050565b6106bb925f5260205f20906020601f840160051c830193106150e657601f0160051c0190614307565b90600a811015611cf45760ff80198354169116179055565b81518051825560208101516001830155604081015191929161516e906001600160a01b031660028501612511565b6060810151615189906001600160a01b031660038501612511565b6080810151600484015560a00151805160058401916001600160401b03821161064e576151c0826151ba85546126f4565b856150f0565b602090601f831160011461520c578260079593604095936151e8935f9261422f57505061431d565b90555b61520560208201516151fc81611cea565b60068601615128565b0151910155565b90601f19831691615220855f5260205f2090565b925f5b818110615264575092600192859260079896604098961061524c575b505050811b0190556151eb565b01515f1960f88460031b161c191690555f808061523f565b92936020600181928786015181550195019301615223565b939290936152a98135926152ae61529b855f52600760205260405f2090565b6020850135938480926154a9565b615087565b801561538f57935b849660408401916152c683612643565b946060016152d390612643565b6152db611be3565b906152e46106dc565b9688885286602089015260408801906152fc916126e5565b61530990606088016126e5565b87608087015260a08601525f14613fbd6143039661537760029761537261537c98610fa09761302c976153855761535b615345600154426128cb565b915b61534f6106eb565b95865260208601612824565b6040840152610af8855f52600860205260405f2090565b615140565b61265b565b019182546128cb565b61535b5f91615347565b506080820135936152b6565b60ff5f516020615d8c5f395f51905f525460401c16156153b757565b631afcd79f60e31b5f5260045ffd5b6153ce61539b565b62015180600155565b6153df615831565b6153e7615888565b6040519060208201925f516020615dcc5f395f51905f528452604083015260608201524660808201523060a082015260a0815261542560c082610689565b51902090565b80548210156125f2575f5260205f2001905f90565b5f818152600360205260409020546154a457600254600160401b81101561064e5761548d615477826001859401600255600261542b565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6154b38282614676565b6145d557805490600160401b82101561064e57826154db61547784600180960185558461542b565b90558054925f520160205260405f2055600190565b80548015615517575f190190615506828261542b565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146155c6575f198401848111610dce5783545f19810194908511610dce575f95858361557997610af8950361557f575b5050506154f0565b55600190565b6155af6155a9916155a06155966155bd958861542b565b90549060031b1c90565b9283918761542b565b906142ea565b85905f5260205260405f2090565b555f8080615571565b505050505f90565b906155e3915f52600660205260405f2061295d565b600181015460a01c60ff1615615605576003018054918203918211610dce5755565b6004018054918201809211610dce5755565b6001600160a01b0316929190600184146156a657811561569d5761566692156156715760405163a9059cbb60e01b60208201529161565e9183916148a99160248401613451565b600592615728565b156111cb5750600190565b6040516340c10f1960e01b6020820152916156959183916148a99160248401613451565b600692615728565b50505050600190565b906156c89350610ab292506156b9611be3565b916001600160a01b0316615771565b6156d157600190565b600590565b5f806111cb93602081519101845af46156ed614689565b916158ba565b613ccd6106bb939261571a60405194859263a9059cbb60e01b602085015260248401613451565b03601f198101845283610689565b81516001600160a01b03909116915f9182916020018285620186a0f161574c614689565b90156145d557805161576857503b1561576457600190565b5f90565b60209150015190565b5f8091615780844710156146cd565b84516001600160a01b0391909116946020018486620186a0f1906157a2614689565b91156157c957156157b5575b5050600190565b805161576857503b15615764575f806157ae565b5050505f90565b906157e5915f52600660205260405f2061295d565b600181015460a01c60ff1615615807576003018054918201809211610dce5755565b6004018054918203918211610dce5755565b916111cb939161582893615918565b9092919261599a565b6158396138c9565b8051908115615849576020012090565b50505f516020615bcc5f395f51905f525480156158635790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615890613983565b80519081156158a0576020012090565b50505f516020615dac5f395f51905f525480156158635790565b906158de57508051156158cf57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061590f575b6158ef575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156158e7565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615985579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612ba5575f516001600160a01b0381161561597b57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611cf457565b6159a381615990565b806159ac575050565b6159b581615990565b600181036159cc5763f645eedf60e01b5f5260045ffd5b6159d581615990565b600281036159f0575063fce698f760e01b5f5260045260245ffd5b806159fc600392615990565b14615a045750565b6335e2f38360e21b5f5260045260245ffd5b80519091906001600160401b03811161064e57615a5781615a445f516020615bac5f395f51905f52546126f4565b5f516020615bac5f395f51905f526150f0565b602092601f8211600114615a8a57615a79929382915f9261422f57505061431d565b5f516020615bac5f395f51905f5255565b5f516020615bac5f395f51905f525f52601f198216935f516020615dec5f395f51905f52915f5b868110615af35750836001959610615adb575b505050811b015f516020615bac5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615ac4565b91926020600181928685015181550194019201615ab156feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a2646970667358221220110f0bad5a3b9ae61270a190e3ea6cd044ae3b0e9560f9a8b3e99038b07485a064736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) PostCallGasReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "postCallGasReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) PostCallGasReserve() (*big.Int, error) {
	return _BSCBridgeV2.Contract.PostCallGasReserve(&_BSCBridgeV2.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) PostCallGasReserve() (*big.Int, error) {
	return _BSCBridgeV2.Contract.PostCallGasReserve(&_BSCBridgeV2.CallOpts)
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

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetPostCallGasReserve(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setPostCallGasReserve", value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPostCallGasReserve(&_BSCBridgeV2.TransactOpts, value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPostCallGasReserve(&_BSCBridgeV2.TransactOpts, value)
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

// BSCBridgeV2PostCallGasReserveSetIterator is returned from FilterPostCallGasReserveSet and is used to iterate over the raw logs and unpacked data for PostCallGasReserveSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2PostCallGasReserveSetIterator struct {
	Event *BSCBridgeV2PostCallGasReserveSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2PostCallGasReserveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2PostCallGasReserveSet)
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
		it.Event = new(BSCBridgeV2PostCallGasReserveSet)
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
func (it *BSCBridgeV2PostCallGasReserveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2PostCallGasReserveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2PostCallGasReserveSet represents a PostCallGasReserveSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2PostCallGasReserveSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostCallGasReserveSet is a free log retrieval operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterPostCallGasReserveSet(opts *bind.FilterOpts) (*BSCBridgeV2PostCallGasReserveSetIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2PostCallGasReserveSetIterator{contract: _BSCBridgeV2.contract, event: "PostCallGasReserveSet", logs: logs, sub: sub}, nil
}

// WatchPostCallGasReserveSet is a free log subscription operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchPostCallGasReserveSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2PostCallGasReserveSet) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2PostCallGasReserveSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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

// ParsePostCallGasReserveSet is a log parse operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParsePostCallGasReserveSet(log types.Log) (*BSCBridgeV2PostCallGasReserveSet, error) {
	event := new(BSCBridgeV2PostCallGasReserveSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
