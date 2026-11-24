// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// IBaseBridgeBridgeTokenArguments is an auto generated low-level Go binding around an user-defined struct.

// IBaseBridgePermitArguments is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryPendingData is an auto generated low-level Go binding around an user-defined struct.

// IBridgeRegistryTokenPair is an auto generated low-level Go binding around an user-defined struct.

// BaseBridgeMetaData contains all meta data concerning the BaseBridge contract.
var BaseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
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
		"230f4cda": "postCallGasReserve()",
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
		"4bc51672": "setPostCallGasReserve(uint256)",
		"bfbf6765": "setTokenPause(uint256,address,bool,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f5160206158ac5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b6040516157e590816100c7823960805181818161266801526127930152f35b6001600160401b0319166001600160401b039081175f5160206158ac5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a0806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714613269575080630b1d8d06146131fd5780631313996b14612e515780631e24abdb14612e34578063230f4cda14612e05578063248a9ca314612ddf5780632f2ff15d14612db357806336568abe14612d6f578063365d71e414612d1557806342cde4e814612ce957806348a00ed814612b745780634bc5167214612b075780634be13f8314612ae05780634ee078ba146129235780634f1ef28614612745578063502cc5c0146126bc57806352d1902d1461265657806356ff54b01461261d5780635b605f5c1461245d5780635c975abb1461242f5780635fd262de14612270578063670e626814611fbd578063761fe2d814611f745780637921487414611f10578063814914b514611e5a57806384b0196e14611d6157806388d67d6d1461144057806389232a0014610ef35780638ae87c5c14610e9357806391cca3db14610e6b57806391cf6d3e14610e4e57806391d1485414610df9578063a10bab0b14610dd1578063a217fddf14610db7578063a3246ad314610cfc578063aa1bd0bc14610cb0578063aa20ed4714610bf4578063ad3cb1cc14610bad578063ae6893f814610b74578063b2b49e2e14610b06578063b33eb36e14610a06578063b7f3358d14610979578063b8aa837e146108ea578063bedb86fb14610801578063bfbf676514610729578063cba25e4b146106c3578063cf56118e14610637578063d477f05f146105d3578063d547741f1461059b578063d5717fc51461054e578063e2af6cd7146104db578063edbbea3914610304578063f0702e8e146102dc5763f698da251461026e575f80fd5b346102d8575f3660031901126102d85760206102886151fd565b610290615254565b60405190838201925f5160206157795f395f51905f528452604083015260608201524660808201523060a082015260a081526102cd60c0826133b6565b519020604051908152f35b5f80fd5b346102d8575f3660031901126102d8576032546040516001600160a01b039091168152602090f35b346102d85760803660031901126102d8576004356103206136a8565b610328613302565b6064356001600160a01b0381169391908490036102d857610347613a72565b81156104cc576001600160a01b03169182156104cc5783156104cc5761036c82614a90565b610477575b815f5260056020526103868360405f20614afb565b156104645760205f5160206157195f395f51905f52916040516103a881613365565b85815282810187815291151560408083018281525f606085018181526080860182815260a080880184815260c089018581528d865260068c528786208f87528c5296909420975188546001600160a01b0319166001600160a01b03918216178955985160018901805496516001600160a81b031990971691909a1617941515901b60ff60a01b1693909317875551939560049491929161044b9190151590613977565b51600285015551600384015551910155604051908152a4005b826311dd05f360e31b5f5260045260245ffd5b6104c760405161048681613380565b8381526003602082015f8152604083015f815260608401915f8352875f52600460205260405f2094518555516001850155516002840155511515910161393e565b610371565b6302bfb33d60e51b5f5260045ffd5b346102d85760203660031901126102d8576004356001600160a01b038116908190036102d857610509613994565b80156104cc575f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346102d85760203660031901126102d8576004355f526004602052600260405f2001546001810180911161058757602090604051908152f35b634e487b7160e01b5f52601160045260245ffd5b346102d85760403660031901126102d8576105d16004356105ba6132ec565b906105cc6105c78261371f565b613b50565b613ce9565b005b346102d85760203660031901126102d8576105ec613318565b6105f4613994565b6001600160a01b0316801561062857603380546001600160a01b031916821790555f5160206156795f395f51905f525f80a2005b638219abe360e01b5f5260045ffd5b346102d8575f3660031901126102d857604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106106ad576106a985610695818703826133b6565b6040519182916020835260208301906135b8565b0390f35b825484526020909301926001928301920161067e565b346102d85760203660031901126102d8576004356001600160a01b038116908190036102d8576106f1613994565b603580546001600160a01b031916821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346102d85760803660031901126102d8576004356107456132ec565b906044358015158082036102d857606435908115158083036102d8576107f46040936107db7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c96610794613ae1565b5f89815260056020528790206001600160a01b038b169a6107bf916107ba908d906149c0565b61394f565b885f526006602052865f208a5f526020526001875f2001613977565b865f526009602052845f20885f52602052845f2061393e565b82519182526020820152a3005b346102d85760203660031901126102d85760043580151581036102d857610826613ae1565b1561088457610833613b96565b600160ff195f5160206156d95f395f51905f525416175f5160206156d95f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f5160206156d95f395f51905f525460ff8116156108db5760ff19165f5160206156d95f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346102d85760403660031901126102d8576004356109066136a8565b61090e613ae1565b815f52600360205260405f2054156109665760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f526004825261095b81600360405f200161393e565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346102d85760203660031901126102d85760043560ff81168091036102d8576109a0613994565b80156109f7576020817f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9260ff195f5160206155995f395f51905f525416175f5160206155995f395f51905f5255604051908152a1005b63f0f15b9160e01b5f5260045ffd5b346102d857610a14366134fc565b905f60408051610a238161334a565b610a2b61390e565b815282602082015201525f52600860205260405f20905f5260205260a060405f2060405190610a598261334a565b610a62816137a3565b8252600760ff60068301541691610a7d602085019384613898565b015460408301908152610afc610af060405195869560208752516060602088015280516080880152602081015182880152600180831b0360408201511660c0880152600180831b0360608201511660e08801526080810151610100880152015160c06101208701526101408601906135eb565b9251604085019061369b565b5160608301520390f35b346102d857610b143661344d565b908051825103610b65575f5b82518110156105d15780610b5e610b396001938561373d565b51838060a01b03610b4a848861373d565b511690610b596105c78261371f565b613c89565b5001610b20565b63031206d560e51b5f5260045ffd5b346102d85760203660031901126102d8576004355f526004602052600160405f2001546001810180911161058757602090604051908152f35b346102d8575f3660031901126102d8576106a9604051610bce6040826133b6565b60058152640352e302e360dc1b60208201526040519182916020835260208301906135eb565b346102d857610c02366134fc565b90610c0b613a03565b610c13613a03565b610c1b613bbd565b805f526007602052610c3982610c348160405f206149c0565b613751565b805f52600860205260405f20825f5260205260405f2060406007815192610c5f8461334a565b610c68816137a3565b8452610c7e60ff60068301541660208601613898565b0154910152610c8d8282613d66565b5f5160206155395f395f51905f525f80a35f5f5160206156f95f395f51905f525d005b346102d85760203660031901126102d8577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435610cef613994565b80600155604051908152a1005b346102d85760203660031901126102d8576004355f525f5160206156595f395f51905f5260205260405f20604051806020835491828152019081935f5260205f20905f5b818110610da15750505081610d569103826133b6565b604051918291602083019060208452518091526040830191905f5b818110610d7f575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610d71565b8254845260209093019260019283019201610d40565b346102d8575f3660031901126102d85760206040515f8152f35b346102d8575f3660031901126102d8576035546040516001600160a01b039091168152602090f35b346102d85760403660031901126102d857610e126132ec565b6004355f525f5160206156b95f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346102d8575f3660031901126102d8576020603454604051908152f35b346102d8575f3660031901126102d8576033546040516001600160a01b039091168152602090f35b346102d857610ea1366134fc565b90610eaa613994565b610eb2613bbd565b610ebc828261475a565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f5160206156f95f395f51905f525d005b346102d85760603660031901126102d857610f0c613318565b610f146132ec565b60443560ff81168091036102d8575f5160206157395f395f51905f5254604081901c60ff161592906001600160401b03811680159081611438575b600114908161142e575b159081611425575b50611416576001600160401b031981166001175f5160206157395f395f51905f5255836113ee575b50610f92614a50565b6001600160a01b03841615610628576001600160a01b03169283156106285781156113df5761101a90610fc3614a50565b610fcb614a50565b610fd3614a50565b60ff195f5160206156d95f395f51905f5254165f5160206156d95f395f51905f5255610ffd614a50565b611005614a50565b61100d614a50565b611015614a50565b613c10565b50611023614a50565b604090815161103283826133b6565b60098152682b30b634b230ba37b960b91b602082015282519061105584836133b6565b60058252640312e302e360dc1b602083015261106f614a50565b611077614a50565b8051906001600160401b0382116112fc5761109f5f5160206155795f395f51905f525461376b565b601f8111611398575b50602090601f831160011461131b576110cb92915f9183611310575b5050614748565b5f5160206155795f395f51905f52555b8051906001600160401b0382116112fc576111035f5160206155d95f395f51905f525461376b565b601f81116112aa575b50602090601f831160011461122d5761112e92915f9183611222575050614748565b5f5160206155d95f395f51905f52555b5f5f5160206155f95f395f51905f52555f5f5160206157595f395f51905f525560ff195f5160206155995f395f51905f525416175f5160206155995f395f51905f5255611189614a50565b62015180600155610400600a555f5160206155595f395f51905f52602082516104008152a1603380546001600160a01b0319168417905551915f5160206156795f395f51905f525f80a2436034556111dd57005b5f5160206157395f395f51905f52805460ff60401b19169055600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b0151905087806110c4565b90601f198316915f5160206155d95f395f51905f525f52815f20925f5b818110611292575090846001959493921061127a575b505050811b015f5160206155d95f395f51905f525561113e565b01515f1960f88460031b161c19169055868080611260565b9293602060018192878601518155019501930161124a565b5f5160206155d95f395f51905f525f526112ec905f5160206157995f395f51905f52601f850160051c810191602086106112f2575b601f0160051c0190614732565b8661110c565b90915081906112df565b634e487b7160e01b5f52604160045260245ffd5b0151905088806110c4565b90601f198316915f5160206155795f395f51905f525f52815f20925f5b8181106113805750908460019594939210611368575b505050811b015f5160206155795f395f51905f52556110db565b01515f1960f88460031b161c1916905587808061134e565b92936020600181928786015181550195019301611338565b5f5160206155795f395f51905f525f526113d9905f5160206156195f395f51905f52601f850160051c810191602086106112f257601f0160051c0190614732565b876110a8565b6338854f1160e21b5f5260045ffd5b6001600160481b0319166001600160401b01175f5160206157395f395f51905f525584610f89565b63f92ee8a960e01b5f5260045ffd5b90501586610f61565b303b159150610f59565b859150610f4f565b60803660031901126102d8576004356001600160401b0381116102d85761146b9036906004016132bc565b60243591906001600160401b0383116102d857366023840112156102d8578260040135611497816133d9565b936114a560405195866133b6565b8185526024602086019260051b820101903682116102d85760248101925b828410611cd75750506044359150506001600160401b0381116102d8576114ee90369060040161360f565b926064356001600160401b0381116102d85761150e90369060040161360f565b91611517613b96565b61151f613bbd565b8082511480611ccd575b80611cc3575b15611cb45734611c9d57909184925f9460be1981360301905b84871015611c82578660051b810135828112156102d85781019661156c818561373d565b5197611578828961373d565b5190611584838861373d565b5199813590815f5260046020526115a68260ff600360405f20015416156138a4565b815f52600560205260405f20906115d2604085019260018060a01b036115cb856136d9565b16906149c0565b6001600160a01b036115e3846136d9565b169015611c705750825f526004602052600260405f2001936001855401809555602081013594808603611c5957506001600160a01b03611622846136d9565b1695606082019e8f611633906136d9565b9360808401359860a085019561164987876136ed565b36906116549261352d565b805190602001206040519160208301937fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191485528b60408501528c60608501526080840152600160a01b600190031660a08301528b60c083015260e082015260e081526116c2610100826133b6565b5190208151918451831480611c4f575b15611c405760ff5f5160206155995f395f51905f525416936116f78486811015614a36565b5f955f935f5b868110611b2a5750505050505050909d9e98999a9b9c9d61172091811015614a36565b5f61172b83836136ed565b15159050611b215761174f61173f8a6136d9565b88611749876136d9565b88613df8565b93908093600d821015611a0b578b9260018b9314611ae0575b50505050600d821015611a0b57600182036117cd575050506001956117a361179d5f5160206155b95f395f51905f52936136d9565b916136d9565b906117bd6040519283928a8060a01b031697429184613d48565b0390a45b01959493929190611548565b909195845f5260076020526117e58660405f20614afb565b15611acd578115611ac55750965b6117fc836136d9565b91611806826136d9565b604051602098916118178a836133b6565b5f8252604051956118278761339b565b8887528a87018a90526001600160a01b039081166040880152166060860152608085018b905260a085015215611abe57611863600154426138be565b905b604051936118728561334a565b845287840199611882828c613898565b60408581019384525f88815260088b528181208a82528b52819020955180518755808b01516001880155908101516002870180546001600160a01b039283166001600160a01b03199182161790915560608301516003890180549190931691161790556080810151600487015560a00151805160058701916001600160401b0382116112fc57611912835461376b565b8c601f8211611a8f575b50508b90601f8311600114611a2a5761193e92915f9183611a1f575050614748565b90555b600685019a5192600d841015611a0b576119d67f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f97608097600760019f946119dc956119fd9960ff8019835416911617905551910155895f5260068c5260405f208e8060a01b036119b1836136d9565b168f8060a01b03165f528c52600260405f20016119cf8582546138be565b90556136d9565b946136d9565b98604051998c8060a01b03168a52890152426040890152606088019061369b565b878060a01b031694a46117c1565b634e487b7160e01b5f52602160045260245ffd5b015190505f806110c4565b5f8481528d8120929190601f198516908f5b828210611a78575050908460019594939210611a60575b505050811b019055611941565b01515f1960f88460031b161c191690555f8080611a53565b60018596829396860151815501950193018f611a3c565b611ab791855f52815f2090601f860160051c82019286106112f257601f0160051c0190614732565b5f8c61191c565b5f90611865565b9050966117f3565b856307a5c91d60e31b5f5260045260245ffd5b611b189550611b109192939450611b0990611b03611afd8a6136d9565b966136d9565b936136ed565b369161352d565b9288886149d3565b88875f80611768565b61174f5f61173f565b611bcf611bc66042611b3a6151fd565b611b42615254565b6040519060208201925f5160206157795f395f51905f528452604083015260608201524660808201523060a082015260a08152611b8060c0826133b6565b5190206040519061190160f01b825260028201528560228201522060ff611ba7858a61373d565b5116611bb3858861373d565b5190611bbf868a61373d565b5192615460565b909291926154d8565b6001600160a01b038181169088161080611c00575b611bf2575b506001016116fd565b6001998a0199909650611be9565b506001600160a01b0381165f9081527fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa5602052604090205460ff16611be4565b63b3f07a3960e01b5f5260045ffd5b50835183146116d2565b859063a6ab65c560e01b5f5260045260245260445ffd5b633142cb1160e11b5f5260045260245ffd5b5f5f5160206156f95f395f51905f525d602060405160018152f35b63943892eb60e01b5f525f6004523460245260445ffd5b6329f517fd60e01b5f5260045ffd5b508083511461152f565b5080855114611529565b83356001600160401b0381116102d8578201366043820112156102d8576024810135611d02816133d9565b91611d1060405193846133b6565b818352602060248185019360051b83010101903682116102d857604401915b818310611d47575050508152602093840193016114c3565b823560ff811681036102d857815260209283019201611d2f565b346102d8575f3660031901126102d8575f5160206155f95f395f51905f52541580611e44575b15611e0757611dd9611d976145df565b6106a9611da261469a565b611de760405191611db46020846133b6565b5f83525f368137604051958695600f60f81b875260e0602088015260e08701906135eb565b9085820360408701526135eb565b904660608501523060808501525f60a085015283820360c08501526135b8565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f5160206157595f395f51905f525415611d87565b346102d85760403660031901126102d857611e736132ec565b611e7b6138cb565b506004355f52600660205260405f209060018060a01b03165f5260205260e060405f20600460405191611ead83613365565b60018060a01b03815416835260ff600182015460018060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c0820152611f0e6040518092613563565bf35b346102d85760203660031901126102d8576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b818110611f5e576106a985610695818703826133b6565b8254845260209093019260019283019201611f47565b346102d85760403660031901126102d857611f8d6132ec565b6004355f52600960205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346102d85760803660031901126102d857600435611fd96132ec565b6044356001600160401b0381116102d857366023820112156102d85761200990369060248160040135910161352d565b906064359060ff82168092036102d8575f546001600160a01b03169081156122615760209161206c915f604051809681958294637c469ea160e11b84528b600485015260018060a01b0316998a60248501526080604485015260848401906135eb565b90606483015203925af1908115612256575f91612214575b5061208d613a72565b82156104cc576001600160a01b03169081156104cc5780156104cc576120b283614a90565b6121bf575b825f5260056020526120cc8260405f20614afb565b156121ac57816020936040516120e181613365565b8281528581018481525f6040808401828152606085018381526080860184815260a080880186815260c089018781528a885260068f528688208c89528f5295909620975188546001600160a01b0319166001600160a01b03918216178955965160018901805495516001600160a81b03199096169190981617931515901b60ff60a01b169290921785555160049492939261217e91151590613977565b516002850155516003840155519101555f5160206157195f395f51905f52856040515f8152a4604051908152f35b506311dd05f360e31b5f5260045260245ffd5b61220f6040516121ce81613380565b8481526003602082015f8152604083015f815260608401915f8352885f52600460205260405f2094518555516001850155516002840155511515910161393e565b6120b7565b90506020813d60201161224e575b8161222f602093836133b6565b810103126102d857516001600160a01b03811681036102d85783612084565b3d9150612222565b6040513d5f823e3d90fd5b6315aeca0d60e11b5f5260045ffd5b60e03660031901126102d8576024356001600160a01b03811690600435908281036102d85761229d613302565b6064359390919060c4356001600160401b0381116102d857366023820112156102d8576004810135936001600160401b0385116102d85736602486840101116102d8576122e8613b96565b6122f28387613fa2565b6122fa613bbd565b6001600160a01b03169485156106285730861461242057600a548015908115612415575b50156124065761237f6123899261234060019660a435908b608435918761401e565b979098604051946123508661332e565b855286602086015233604086015260608501528960808501528860a08501528760c0850152602436920161352d565b60e08201526141b5565b146123a9575f5f5160206156f95f395f51905f525d602060405160018152f35b6123b96123be926123c4946138be565b6138be565b34613901565b806123d1575b8080611c82565b6040516123ec916123e36020836133b6565b5f825233614576565b156123f757806123ca565b63559d141b60e11b5f5260045ffd5b633953b33f60e11b5f5260045ffd5b90508511158861231e565b6304c821ad60e31b5f5260045ffd5b346102d8575f3660031901126102d857602060ff5f5160206156d95f395f51905f5254166040519015158152f35b346102d85760203660031901126102d857600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106126045750506124ac925003836133b6565b8151916124b8836133d9565b926124c660405194856133b6565b8084526124d5601f19916133d9565b015f5b8181106125ed5750505f5b815181101561259d575f838152600660205260409020600191906001600160a01b0361250f838661373d565b5116838060a01b03165f5260205260405f2060046040519161253083613365565b848060a01b03815416835260ff85820154868060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c082015261258b828761373d565b52612596818661373d565b50016124e3565b836040518091602082016020835281518091526020604084019201905f5b8181106125c9575050500390f35b91935091602060e0826125df6001948851613563565b0194019101918493926125bb565b6020906125f86138cb565b828288010152016124d8565b8454835260019485019487945060209093019201612497565b346102d85760203660031901126102d8575f5160206155595f395f51905f526020600435612649613a72565b80600a55604051908152a1005b346102d8575f3660031901126102d8577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036126ad5760206040515f5160206156395f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b346102d85760603660031901126102d8576024356004356126db613a03565b805f5260076020526126f482610c348160405f206149c0565b7fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db79853659496020612723604435426138be565b835f526008825260405f20855f52825280600760405f200155604051908152a3005b60403660031901126102d857612759613318565b6024356001600160401b0381116102d857366023820112156102d85761278990369060248160040135910161352d565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115612901575b506126ad576127cb613994565b6040516352d1902d60e01b81526001600160a01b0383169290602081600481875afa5f91816128cd575b5061280d5783634c9c8ce360e01b5f5260045260245ffd5b805f5160206156395f395f51905f528592036128bb5750813b156128a9575f5160206156395f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115612891575f808360206105d195519101845af461288b6144f8565b916153aa565b50503461289a57005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d6020116128f9575b816128e9602093836133b6565b810103126102d8575190856127f5565b3d91506128dc565b5f5160206156395f395f51905f52546001600160a01b031614159050836127be565b346102d857612931366134fc565b612939613b96565b612941613bbd565b815f52600460205261295e8260ff600360405f20015416156138a4565b815f52600760205261297781610c348160405f206149c0565b815f52600860205260405f20815f5260205260405f206040516129998161334a565b6129a2826137a3565b8152600760ff600684015416926129bd602084019485613898565b015491604082019283528051600d811015611a0b57600214159081612ac9575b81612aaf575b50905160408101516080820151606090920151612a1093926001600160a01b039182169290911687613eef565b50600d811015611a0b5760016040519180602084015260208352612a356040846133b6565b03612a875750519182158015612a7e575b15612a6757612a559250613d66565b5f5f5160206156f95f395f51905f525d005b82637a88099160e11b5f526004524260245260445ffd5b50428310612a46565b60405162461bcd60e51b815260206004820152908190612aab9060248301906135eb565b0390fd5b90505190600d821015611a0b5790600a1415612a106129e3565b80915051600d811015611a0b5760091415906129dd565b346102d8575f3660031901126102d8575f546040516001600160a01b039091168152602090f35b346102d85760203660031901126102d857600435612b23613994565b61c3508110612b655760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91603654908060365582519182526020820152a1005b63050779f960e21b5f5260045ffd5b346102d85760603660031901126102d857602435600435612b93613302565b612b9b613a03565b612ba3613bbd565b815f526007602052612bbc83610c348160405f206149c0565b815f52600860205260405f20835f5260205260405f2060406007815192612be28461334a565b612beb816137a3565b8452612c0160ff60068301541660208601613898565b015491015280612c11848461475a565b916001600160a01b031615612cd5575b8151906020830190815193612c53604082019460018060a01b038651169684608085019889519260a0870151946149d3565b50600d811015611a0b5760018103612cbd5750905f5160206155b95f395f51905f52915192519360018060a01b039051169451612c97604051928392429184613d48565b0390a45f5160206155395f395f51905f525f80a35f5f5160206156f95f395f51905f525d005b63290cd55f60e01b5f52612cd09061368d565b60245ffd5b5060608101516001600160a01b0316612c21565b346102d8575f3660031901126102d857602060ff5f5160206155995f395f51905f525416604051908152f35b346102d857612d233661344d565b908051825103610b65575f5b82518110156105d15780612d68612d486001938561373d565b51838060a01b03612d59848861373d565b5116906105cc6105c78261371f565b5001612d2f565b346102d85760403660031901126102d857612d886132ec565b336001600160a01b03821603612da4576105d190600435613ce9565b63334bd91960e11b5f5260045ffd5b346102d85760403660031901126102d8576105d1600435612dd26132ec565b90610b596105c78261371f565b346102d85760203660031901126102d8576020612dfd60043561371f565b604051908152f35b346102d8575f3660031901126102d85760365480612e2c57506020620249f0604051908152f35b602090612dfd565b346102d8575f3660031901126102d8576020600a54604051908152f35b60403660031901126102d8576004356001600160401b0381116102d857612e7c9036906004016132bc565b6024356001600160401b0381116102d857366023820112156102d85760048101356001600160401b0381116102d85736602460e08302840101116102d857612ec2613b96565b612eca613bbd565b335f9081527f7bf51035c163aa501aba0b1731707347fc5168978f3fc3c84ea475ee63d4b83f602052604090205460ff16156131c65734611c9d57808303611cb4575f5b838110612f27575f5f5160206156f95f395f51905f525d005b612f328185876136b7565b3590612f4a6020612f4483888a6136b7565b016136d9565b612f5a6060612f4484898b6136b7565b926080612f6884898b6136b7565b01359360a0612f78858a8c6136b7565b01359260c0612f88868b8d6136b7565b013590878610156131b2578a968a93612fb9612faf898d60e082020197602489019c6136b7565b60e08101906136ed565b6001600160a01b0385169a9195909491612fd38c8a613fa2565b6001600160a01b03612fe4826136d9565b168c14906001600160a01b0390612ffa906136d9565b16901561319c575060448701986001600160a01b036130188b6136d9565b6001600160a01b039095169416840361318d5730841461242057600a548015908115613182575b50156124065784613050938a61401e565b9290916064870135966130828861306b876123b988886138be565b81101561307c886123b989896138be565b90613bf2565b6032546001600160a01b0316906130988b6136d9565b9160a48201359160ff83168093036102d857813b156102d8578e926040519b8c958695637f40ec1760e11b87526004870152600160a01b6001900316602486015230604486015260648501526084820135608485015260a484015260c481013560c484015260e4013560e48301525a925f610104928195f19a8b156122565761313061316c9a61237f9960019e613172575b506136d9565b906040519961313e8b61332e565b8a5260208a01528b8060a01b031660408901526060880152608087015260a086015260c0850152369161352d565b01612f0e565b5f61317c916133b6565b5f61312a565b90508611155f61303f565b630672aec160e01b5f5260045ffd5b8b6313f7c32b60e31b5f5260045260245260445ffd5b634e487b7160e01b5f52603260045260245ffd5b63e2517d3f60e01b5f52336004527f6b8b15f1c11543d8280deaa7c24d12fffba6a357e4428e8c43e4234790186bff60245260445ffd5b346102d85760203660031901126102d8576004356001600160a01b038116908190036102d85761322b613994565b801561062857603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b346102d85760203660031901126102d8576004359063ffffffff60e01b82168092036102d857602091637965db0b60e01b81149081156132ab575b5015158152f35b6301ffc9a760e01b149050836132a4565b9181601f840112156102d8578235916001600160401b0383116102d8576020808501948460051b0101116102d857565b602435906001600160a01b03821682036102d857565b604435906001600160a01b03821682036102d857565b600435906001600160a01b03821682036102d857565b61010081019081106001600160401b038211176112fc57604052565b606081019081106001600160401b038211176112fc57604052565b60e081019081106001600160401b038211176112fc57604052565b608081019081106001600160401b038211176112fc57604052565b60c081019081106001600160401b038211176112fc57604052565b601f909101601f19168101906001600160401b038211908210176112fc57604052565b6001600160401b0381116112fc5760051b60200190565b9080601f830112156102d8578135613407816133d9565b9261341560405194856133b6565b81845260208085019260051b8201019283116102d857602001905b82821061343d5750505090565b8135815260209182019101613430565b9060406003198301126102d8576004356001600160401b0381116102d85782613478916004016133f0565b91602435906001600160401b0382116102d857806023830112156102d8578160040135906134a5826133d9565b926134b360405194856133b6565b8284526024602085019360051b8201019182116102d857602401915b8183106134dc5750505090565b82356001600160a01b03811681036102d8578152602092830192016134cf565b60409060031901126102d8576004359060243590565b6001600160401b0381116112fc57601f01601f191660200190565b92919261353982613512565b9161354760405193846133b6565b8294818452818301116102d8578281602093845f960137010152565b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b90602080835192838152019201905f5b8181106135d55750505090565b82518452602093840193909201916001016135c8565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9080601f830112156102d8578135613626816133d9565b9261363460405194856133b6565b81845260208085019260051b820101918383116102d85760208201905b83821061366057505050505090565b81356001600160401b0381116102d857602091613682878480948801016133f0565b815201910190613651565b600d811015611a0b57600452565b90600d821015611a0b5752565b6024359081151582036102d857565b91908110156131b25760051b8101359060fe19813603018212156102d8570190565b356001600160a01b03811681036102d85790565b903590601e19813603018212156102d857018035906001600160401b0382116102d8576020019181360383136102d857565b5f525f5160206156b95f395f51905f52602052600160405f20015490565b80518210156131b25760209160051b010190565b156137595750565b6373a1390160e11b5f5260045260245ffd5b90600182811c92168015613799575b602083101461378557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161377a565b906040516137b08161339b565b60058193805483526001810154602084015260018060a01b03600282015416604084015260018060a01b03600382015416606084015260048101546080840152019060405180925f908054906138058261376b565b80855291600181169081156138715750600114613832575b505060a0929161382e9103846133b6565b0152565b5f908152602081209092505b81831061385557505081016020018161382e61381d565b602091935080600191548385890101520191019091849261383e565b60ff191660208681019190915292151560051b8501909201925083915061382e905061381d565b600d821015611a0b5752565b156138ac5750565b636fc24b7960e11b5f5260045260245ffd5b9190820180921161058757565b604051906138d882613365565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9190820391821161058757565b6040519061391b8261339b565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b9060ff801983541691151516179055565b156139575750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff16156139cc57565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b335f9081527fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143602052604090205460ff1615613a3b57565b63e2517d3f60e01b5f52336004527f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0960245260445ffd5b335f9081527f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f602052604090205460ff1615613aaa57565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b335f9081527f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456602052604090205460ff1615613b1957565b63e2517d3f60e01b5f52336004527f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92960245260445ffd5b5f8181525f5160206156b95f395f51905f526020908152604080832033845290915290205460ff1615613b805750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f5160206156d95f395f51905f525416613bae57565b63d93c066560e01b5f5260045ffd5b5f5160206156f95f395f51905f525c613be35760015f5160206156f95f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b15613bfb575050565b63943892eb60e01b5f5260045260245260445ffd5b90613c1b825f614886565b9182613c245750565b5f80525f5160206156595f395f51905f526020526001600160a01b0316613c6b817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d614afb565b15613c735750565b63d180cb3160e01b5f526004525f60245260445ffd5b919091613c968382614886565b9283613ca0575050565b815f525f5160206156595f395f51905f52602052613ccb60405f209160018060a01b03168092614afb565b15613cd4575050565b63d180cb3160e01b5f5260045260245260445ffd5b919091613cf68382614924565b9283613d00575050565b815f525f5160206156595f395f51905f52602052613d2b60405f209160018060a01b03168092614b42565b15613d34575050565b62a95f1b60e31b5f5260045260245260445ffd5b604091949392606082019560018060a01b0316825260208201520152565b90613d709161475a565b60018060a01b036060820151168151906020830190815193613daf604082019460018060a01b038651169684608085019889519260a0870151946149d3565b50600d811015611a0b5760018103612cbd5750905f5160206155b95f395f51905f52915192519360018060a01b039051169451613df3604051928392429184613d48565b0390a4565b929391905f9360018060a01b0360325416908115613ee0575f9081526009602090815260408083206001600160a01b039095168084529490915290205460ff16613ed457602092915f60649260405198899586946319ab07bb60e21b86526004860152602485015260018060a01b031660448401525af1928315612256575f93613e97575b5082600d811015611a0b57600103613e9157565b60019150565b9092506020813d602011613ecc575b81613eb3602093836133b6565b810103126102d85751600d8110156102d857915f613e7d565b3d9150613ea6565b50600294505f93505050565b6330d45fb160e01b5f5260045ffd5b60325493955f9590946001600160a01b0316918215613ee0575f9081526009602090815260408083206001600160a01b039096168084529590915290205460ff16613f9557613f8c57602092915f60649260405198899586946319ab07bb60e21b86526004860152602485015260018060a01b031660448401525af1928315612256575f93613e97575082600d811015611a0b57600103613e9157565b50600194505050565b50600295505f9450505050565b5f8181526005602052604090206001600160a01b03831692613fc9916107ba9085906149c0565b805f526004602052613fe68160ff600360405f20015416156138a4565b5f52600660205260405f20815f5260205260ff600160405f20015460a81c1661400c5750565b6338384f6f60e11b5f5260045260245ffd5b5f8181526006602090815260408083206001600160a01b0390951680845294909152908190209051929694959260c0919061405881613365565b81546001600160a01b0390811682526001830154908116602083015260a081811c60ff908116158015604086015260a89390931c16151560608401526002840154608084015260038401549083015260049092015492018290526141a1575b506032546001600160a01b0316958615613ee057606460609260405198899384926337dba1f760e21b8452600484015260248301528660448301525afa908115612256575f955f905f9361415e575b50829581978515159586614153575b505084614148575b50508261413d575b50501561412e57565b6358e8878560e01b5f5260045ffd5b101590505f80614125565b101592505f8061411d565b101594505f80614115565b96505090506060853d606011614199575b8161417c606093836133b6565b810103126102d8578451906040602087015196015191955f614106565b3d915061416f565b8084806141af931015613bf2565b5f6140b7565b80515f526004602052600160405f2001600181540180915581515f52600660205260405f20602083019060018060a01b0382511660018060a01b03165f5260205260405f206001808060a01b03910154169280519360018060a01b0383511694604083019060018060a01b03825116966080850191825160a08701998a519361424460c08a01958651906138be565b93805f52600660205260405f2060018060a01b0383165f5260205260405f2060ff600182015460a01c165f146144e3576003016142828582546138be565b90555b60018203614370575050508061429e836142ac936138be565b34101561307c8434936138be565b80614339575b50905f5160206157b95f395f51905f529461432e925b81519760018060a01b039051169460018060a01b039051169960e060018060a01b03606085015116945191519251930151936040519788978852602088015260408701526060860152608085015260a084015261010060c08401526101008301906135eb565b4260e08301520390a4565b6033546040519293926143649290916001600160a01b031661435c6020846133b6565b5f8352614576565b156123f757905f6142b2565b90919293989596979834611c9d576143c261438b82876138be565b604051906323b872dd60e01b60208301528660248301523060448301526064820152606481526143bc6084826133b6565b84615408565b80614493575b505f52600660205260405f2060018060a01b0382165f5260205260ff600160405f20015460a01c1615614414575b5050509161432e915f5160206157b95f395f51905f529594936142c8565b604051632770a7eb60e21b815260208180614433873060048401614a1b565b03815f865af1908115612256575f91614464575b506143f657639ac2f56d60e01b5f5260045260245260445260645ffd5b614486915060203d60201161448c575b61447e81836133b6565b810190614a03565b5f614447565b503d614474565b60335460405163a9059cbb60e01b60208201526144dd9290916144d79183916144c991906001600160a01b031660248401614a1b565b03601f1981018352826133b6565b83615408565b5f6143c8565b6004016144f1858254613901565b9055614285565b3d15614522573d9061450982613512565b9161451760405193846133b6565b82523d5f602084013e565b606090565b81516001600160a01b03909116915f9182916020018285620186a0f161454b6144f8565b901561457057805161456757503b1561456357600190565b5f90565b60209150015190565b50505f90565b8147106145d05782516001600160a01b03909116925f9182916020018486620186a0f1906145a26144f8565b91156145c957156145b5575b5050600190565b805161456757503b15614563575f806145ae565b5050505f90565b632b60b36f60e21b5f5260045ffd5b6040515f5160206155795f395f51905f5254815f6145fc8361376b565b808352926001811690811561467b5750600114614623575b614620925003826133b6565b90565b505f5160206155795f395f51905f525f90815290915f5160206156195f395f51905f525b81831061465f57505090602061462092820101614614565b6020919350806001915483858801015201910190918392614647565b6020925061462094915060ff191682840152151560051b820101614614565b6040515f5160206155d95f395f51905f5254815f6146b78361376b565b808352926001811690811561467b57506001146146da57614620925003826133b6565b505f5160206155d95f395f51905f525f90815290915f5160206157995f395f51905f525b81831061471657505090602061462092820101614614565b60209193508060019154838588010152019101909183926146fe565b81811061473d575050565b5f8155600101614732565b8160011b915f199060031b1c19161790565b919061476461390e565b50825f52600760205261477a8160405f20614b42565b1561487457825f52600860205260405f20815f5260205261479d60405f206137a3565b5f848152600660209081526040808320818501516001600160a01b03168452909152902060808201516002909101805492959290916147db91613901565b90555f52600860205260405f20905f526020525f6007604082208281558260018201558260028201558260038201558260048201556005810161481e815461376b565b9081614831575b50508260068201550155565b81601f8693116001146148485750555b5f80614825565b8183526020832061486491601f0160051c810190600101614732565b8082528160208120915555614841565b637f11bea960e01b5f5260045260245ffd5b5f8181525f5160206156b95f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16614570575f8181525f5160206156b95f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5f8181525f5160206156b95f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615614570575f8181525f5160206156b95f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001915f520160205260405f2054151590565b9594939291906001600160a01b03831630146149f7576149f39596614c07565b9091565b505050925050600c9190565b908160209103126102d8575180151581036102d85790565b6001600160a01b039091168152602081019190915260400190565b15614a3e5750565b631aebd17960e11b5f5260045260245ffd5b60ff5f5160206157395f395f51905f525460401c1615614a6c57565b631afcd79f60e31b5f5260045ffd5b80548210156131b2575f5260205f2001905f90565b805f52600360205260405f2054155f14614af657600254600160401b8110156112fc57614adf614ac98260018594016002556002614a7b565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b614b0582826149c0565b61457057805490600160401b8210156112fc5782614b2d614ac9846001809601855584614a7b565b90558054925f520160205260405f2055600190565b906001820191815f528260205260405f20548015155f14614bff575f1981018181116105875782545f1981019190821161058757818103614bca575b50505080548015614bb6575f190190614b978282614a7b565b8154905f199060031b1b19169055555f526020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b614bea614bda614ac99386614a7b565b90549060031b1c92839286614a7b565b90555f528360205260405f20555f8080614b7e565b505050505f90565b949291936080525f90855f52600660205260405f2060018060a01b0360018060a01b038716165f5260205260ff600160405f20015460a01c16935f9160018060a01b036035541691601881511015806151f4575b806151ea575b614cc0575b505050614cb65750608051614c7c9291846152d6565b91600d831015611a0b5760018314614c975750509060805190565b608051614caf9350916001600160a01b031690615286565b6001905f90565b9350505050905f90565b602081015160601c92896034830151928a60405193614cde85613380565b5f60208601525f60408601526060808601525f85525f5f80604051602081019062483e5560e91b82528c602482015260248152614d1c6044826133b6565b5190855afa614d296144f8565b90806151de575b6151d2575b50614d45575b5050505050614c66565b6001600160a01b03821660011460408601819052928315615096575b8b6040870151614d73575b5050614d3b565b86865f93859c9e9a9c85948d9f9b9d614e0b90603654801589146150915750620249f05b805a10891461507d57506001602089985b01528b158814615071576144c988955b6040519485936020850199631eeed51360e01b8b52602486015260448501528d60018060a01b0316606485015260018060a01b0316608484015260805160a484015260c060c484015260e48301906135eb565b5193f1614e166144f8565b6060890152151560408801528215615022575b50506040850151151580615012575b15614f2a575050505060600151908151820191604081602085019403126102d8576020810151604082015190916001600160401b0382116102d857019183603f840112156102d857602083015192614e8f84613512565b94614e9d60405196876133b6565b848652604082860101116102d8578a945f6020865f5160206156995f395f51905f52976040614efd96018386015e8301015260405193849363ffffffff60e01b1684526001602085015260408401526080606084015260808301906135eb565b0390a4608051614f17906001600160a01b03851686615286565b6001805b5f80808089818a81808b614d6c565b80979592969491939715901561500a575b614f8e575b9060605f5160206156995f395f51905f5292015190614f8660405192839263ffffffff60e01b1683525f60208401525f60408401526080606084015260808301906135eb565b0390a4614f1b565b915060405190632770a7eb60e21b82526020828b815f81614fb56080513060048401614a1b565b03926001600160a01b03165af18015612256578b935f5160206156995f395f51905f5293606092614feb575b5091925050614f40565b6150039060203d60201161448c5761447e81836133b6565b505f614fe1565b508815614f3b565b5060406060860151511015614e38565b5f91829182604051602081019263095ea7b360e01b84526024820152816044820152604481526150536064826133b6565b51926001600160a01b03165af1506150696144f8565b508a5f614e29565b6144c960805195614db8565b602061508b6001925a613901565b98614da8565b614d97565b8c15615161575b5f80604051602081019063095ea7b360e01b82526150c5816144c96080518960248401614a1b565b5190826001600160a01b0388165af16150dc6144f8565b50158015604088015280615159575b15614d615791935060206040518095632770a7eb60e21b8252815f816151176080513060048401614a1b565b03926001600160a01b03165af1918215612256578e948e9361513a575b50614d61565b6151529060203d60201161448c5761447e81836133b6565b505f615134565b508c156150eb565b91909350602060405180926340c10f1960e01b8252815f816151896080513060048401614a1b565b03926001600160a01b03165af1908115612256575f916151b3575b50156145d0578c928c9161509d565b6151cc915060203d60201161448c5761447e81836133b6565b5f6151a4565b6020915001515f614d35565b50602081511015614d30565b50823b1515614c61565b50821515614c5b565b6152056145df565b8051908115615215576020012090565b50505f5160206155f95f395f51905f5254801561522f5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61525c61469a565b805190811561526c576020012090565b50505f5160206157595f395f51905f5254801561522f5790565b5f52600660205260405f209060018060a01b03165f5260205260405f2060ff600182015460a01c165f146152c75760036152c39101918254613901565b9055565b60046152c391019182546138be565b6001600160a01b03169291906001841461536557811561535c5761532592156153305760405163a9059cbb60e01b60208201529161531d9183916144c99160248401614a1b565b600592614527565b156146205750600190565b6040516340c10f1960e01b6020820152916153549183916144c99160248401614a1b565b600692614527565b50505050600190565b915080925047106153a35761539491604051916153836020846133b6565b5f83526001600160a01b0316614576565b1561539e57600190565b600590565b5050600b90565b906153ce57508051156153bf57805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806153ff575b6153df575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156153d7565b905f602091828151910182855af115612256575f513d61545757506001600160a01b0381163b155b6154375750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615430565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116154cd579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612256575f516001600160a01b038116156154c357905f905f90565b505f906001905f90565b5050505f9160039190565b6004811015611a0b57806154ea575050565b600181036155015763f645eedf60e01b5f5260045ffd5b6002810361551c575063fce698f760e01b5f5260045260245ffd5b6003146155265750565b6335e2f38360e21b5f5260045260245ffdfeb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db11e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5dbf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a164736f6c634300081c000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BaseBridge *BaseBridgeCaller) PostCallGasReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseBridge.contract.Call(opts, &out, "postCallGasReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BaseBridge *BaseBridgeSession) PostCallGasReserve() (*big.Int, error) {
	return _BaseBridge.Contract.PostCallGasReserve(&_BaseBridge.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BaseBridge *BaseBridgeCallerSession) PostCallGasReserve() (*big.Int, error) {
	return _BaseBridge.Contract.PostCallGasReserve(&_BaseBridge.CallOpts)
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

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BaseBridge *BaseBridgeTransactor) SetPostCallGasReserve(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BaseBridge.contract.Transact(opts, "setPostCallGasReserve", value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BaseBridge *BaseBridgeSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPostCallGasReserve(&_BaseBridge.TransactOpts, value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BaseBridge *BaseBridgeTransactorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BaseBridge.Contract.SetPostCallGasReserve(&_BaseBridge.TransactOpts, value)
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

// BaseBridgePostCallGasReserveSetIterator is returned from FilterPostCallGasReserveSet and is used to iterate over the raw logs and unpacked data for PostCallGasReserveSet events raised by the BaseBridge contract.
type BaseBridgePostCallGasReserveSetIterator struct {
	Event *BaseBridgePostCallGasReserveSet // Event containing the contract specifics and raw log

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
func (it *BaseBridgePostCallGasReserveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseBridgePostCallGasReserveSet)
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
		it.Event = new(BaseBridgePostCallGasReserveSet)
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
func (it *BaseBridgePostCallGasReserveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseBridgePostCallGasReserveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseBridgePostCallGasReserveSet represents a PostCallGasReserveSet event raised by the BaseBridge contract.
type BaseBridgePostCallGasReserveSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostCallGasReserveSet is a free log retrieval operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BaseBridge *BaseBridgeFilterer) FilterPostCallGasReserveSet(opts *bind.FilterOpts) (*BaseBridgePostCallGasReserveSetIterator, error) {

	logs, sub, err := _BaseBridge.contract.FilterLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return &BaseBridgePostCallGasReserveSetIterator{contract: _BaseBridge.contract, event: "PostCallGasReserveSet", logs: logs, sub: sub}, nil
}

// WatchPostCallGasReserveSet is a free log subscription operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BaseBridge *BaseBridgeFilterer) WatchPostCallGasReserveSet(opts *bind.WatchOpts, sink chan<- *BaseBridgePostCallGasReserveSet) (event.Subscription, error) {

	logs, sub, err := _BaseBridge.contract.WatchLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseBridgePostCallGasReserveSet)
				if err := _BaseBridge.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
func (_BaseBridge *BaseBridgeFilterer) ParsePostCallGasReserveSet(log types.Log) (*BaseBridgePostCallGasReserveSet, error) {
	event := new(BaseBridgePostCallGasReserveSet)
	if err := _BaseBridge.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
