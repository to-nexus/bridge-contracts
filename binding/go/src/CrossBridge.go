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

// CrossBridgeMetaData contains all meta data concerning the CrossBridge contract.
var CrossBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeTokenForwarded\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupplyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"setCrossSupplyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"CrossSupplyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"computed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ctxValue\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeForwardAmountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardContextInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardNotExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardNotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staged\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeForwardStagedInsufficient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossSupplyAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"CrossBridgeManualReleaseExceedsSupplyLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Disabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"a10bab0b": "bridgeExecutor()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"55fd0776": "bridgeTokenForwarded(uint256,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"cd959706": "crossSupply()",
		"69ceb2f1": "crossSupplyLimit()",
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
		"b8886e1a": "setCrossSupplyLimit(uint256)",
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
	Bin: "0x60a080604052346100c257306080525f516020615d1f5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615c5890816100c782396080518181816129090152612a340152f35b6001600160401b0319166001600160401b039081175f516020615d1f5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a0806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714613585575080630b1d8d06146135195780631313996b146131735780631e24abdb14613156578063230f4cda14613127578063248a9ca3146131095780632f2ff15d1461309f57806336568abe1461305b578063365d71e41461300157806342cde4e814612fd557806348a00ed814612e155780634bc5167214612da85780634be13f8314612d815780634ee078ba14612bc45780634f1ef286146129e6578063502cc5c01461295d57806352d1902d146128f757806355fd07761461242057806356ff54b0146123d45780635b605f5c146122145780635c975abb146121e65780635fd262de1461204b578063670e626814611da057806369ceb2f114611d83578063761fe2d814611d3a5780637921487414611cd6578063814914b514611c2057806384b0196e146119c957806388d67d6d1461108357806389232a00146110485780638ae87c5c14610fe857806391cca3db14610fc057806391cf6d3e14610fa357806391d1485414610f4e578063a10bab0b14610f26578063a217fddf14610f0c578063a3246ad314610e51578063aa1bd0bc14610e05578063aa20ed4714610ce4578063ad3cb1cc14610c9d578063ae6893f814610c64578063b2b49e2e14610bb3578063b33eb36e14610ab3578063b7f3358d14610a26578063b8886e1a146109da578063b8aa837e1461094b578063bedb86fb14610862578063bfbf67651461078a578063cba25e4b14610724578063cd95970614610702578063cf56118e14610676578063d477f05f146105ff578063d547741f146105c7578063d5717fc51461057a578063e2af6cd714610507578063edbbea3914610330578063f0702e8e146103085763f698da251461029a575f80fd5b34610304575f3660031901126103045760206102b46154bd565b6102bc6155c4565b60405190838201925f516020615bec5f395f51905f528452604083015260608201524660808201523060a082015260a081526102f960c0826136d2565b519020604051908152f35b5f80fd5b34610304575f366003190112610304576032546040516001600160a01b039091168152602090f35b346103045760803660031901126103045760043561034c613a01565b61035461361e565b6064356001600160a01b03811693919084900361030457610373613e38565b81156104f8576001600160a01b03169182156104f85783156104f85761039882614c6f565b6104a3575b815f5260056020526103b28360405f20614cda565b156104905760205f516020615bac5f395f51905f52916040516103d481613681565b85815282810187815291151560408083018281525f606085018181526080860182815260a080880184815260c089018581528d865260068c528786208f87528c5296909420975188546001600160a01b0319166001600160a01b03918216178955985160018901805496516001600160a81b031990971691909a1617941515901b60ff60a01b169390931787555193956004949192916104779190151590613d1d565b51600285015551600384015551910155604051908152a4005b826311dd05f360e31b5f5260045260245ffd5b6104f36040516104b28161369c565b8381526003602082015f8152604083015f815260608401915f8352875f52600460205260405f20945185555160018501555160028401555115159101613ce4565b61039d565b6302bfb33d60e51b5f5260045ffd5b34610304576020366003190112610304576004356001600160a01b0381169081900361030457610535613d5a565b80156104f8575f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b34610304576020366003190112610304576004355f526004602052600260405f200154600181018091116105b357602090604051908152f35b634e487b7160e01b5f52601160045260245ffd5b34610304576040366003190112610304576105fd6004356105e6613608565b906105f86105f382613a78565b613f16565b613fb8565b005b3461030457602036600319011261030457610618613634565b610620613d5a565b6001600160a01b0316801561066757603380546001600160a01b031916821790557f76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc508715f80a2005b638219abe360e01b5f5260045ffd5b34610304575f36600319011261030457604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106106ec576106e8856106d4818703826136d2565b604051918291602083526020830190613911565b0390f35b82548452602090930192600192830192016106bd565b34610304575f36600319011261030457602061071c613d3a565b604051908152f35b34610304576020366003190112610304576004356001600160a01b0381169081900361030457610752613d5a565b603580546001600160a01b031916821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b34610304576080366003190112610304576004356107a6613608565b9060443580151580820361030457606435908115158083036103045761085560409361083c7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c966107f5613ea7565b5f89815260056020528790206001600160a01b038b169a6108209161081b908d90614a25565b613cf5565b885f526006602052865f208a5f526020526001875f2001613d1d565b865f526009602052845f20885f52602052845f20613ce4565b82519182526020820152a3005b3461030457602036600319011261030457600435801515810361030457610887613ea7565b156108e557610894613f5c565b600160ff195f516020615b6c5f395f51905f525416175f516020615b6c5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615b6c5f395f51905f525460ff81161561093c5760ff19165f516020615b6c5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b3461030457604036600319011261030457600435610967613a01565b61096f613ea7565b815f52600360205260405f2054156109c75760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526109bc81600360405f2001613ce4565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b34610304576020366003190112610304577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435610a19613d5a565b80606455604051908152a1005b346103045760203660031901126103045760043560ff811680910361030457610a4d613d5a565b8015610aa4576020817f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9260ff195f516020615a4c5f395f51905f525416175f516020615a4c5f395f51905f5255604051908152a1005b63f0f15b9160e01b5f5260045ffd5b3461030457610ac136613818565b905f60408051610ad081613666565b610ad8613cb4565b815282602082015201525f52600860205260405f20905f5260205260a060405f2060405190610b0682613666565b610b0f81613afc565b8252600760ff60068301541691610b2a602085019384613bf1565b015460408301908152610ba9610b9d60405195869560208752516060602088015280516080880152602081015182880152600180831b0360408201511660c0880152600180831b0360608201511660e08801526080810151610100880152015160c0610120870152610140860190613944565b925160408501906139f4565b5160608301520390f35b3461030457610bc136613769565b908051825103610c55575f5b82518110156105fd57610be08183613a96565b516001600160a01b03610bf38386613a96565b5116610c016105f383613a78565b610c0b81836148eb565b610c1a575b5050600101610bcd565b815f525f516020615b0c5f395f51905f52602052610c3b8160405f20614cda565b610c10575b63d180cb3160e01b5f5260045260245260445ffd5b63031206d560e51b5f5260045ffd5b34610304576020366003190112610304576004355f526004602052600160405f200154600181018091116105b357602090604051908152f35b34610304575f366003190112610304576106e8604051610cbe6040826136d2565b60058152640352e302e360dc1b6020820152604051918291602083526020830190613944565b3461030457610cf236613818565b90610cfb613dc9565b610d03613dc9565b610d0b613f83565b805f526007602052610d2982610d248160405f20614a25565b613aaa565b805f52600860205260405f20825f5260205260405f20600760405191610d4e83613666565b610d5781613afc565b8352610d6d60ff60068301541660208501613bf1565b015460408201525160016080818060a01b03604084015116920151911480610dfa575b610dc3575b50610da08282614035565b5f516020615a0c5f395f51905f525f80a35f5f516020615b8c5f395f51905f525d005b610dd490610dcf613d3a565b613c17565b60645490818111610de55750610d95565b6331d1ec8f60e11b5f5260045260245260445ffd5b506065548214610d90565b34610304576020366003190112610304577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435610e44613d5a565b80600155604051908152a1005b34610304576020366003190112610304576004355f525f516020615b0c5f395f51905f5260205260405f20604051806020835491828152019081935f5260205f20905f5b818110610ef65750505081610eab9103826136d2565b604051918291602083019060208452518091526040830191905f5b818110610ed4575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610ec6565b8254845260209093019260019283019201610e95565b34610304575f3660031901126103045760206040515f8152f35b34610304575f366003190112610304576035546040516001600160a01b039091168152602090f35b3461030457604036600319011261030457610f67613608565b6004355f525f516020615b4c5f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b34610304575f366003190112610304576020603454604051908152f35b34610304575f366003190112610304576033546040516001600160a01b039091168152602090f35b3461030457610ff636613818565b90610fff613d5a565b611007613f83565b61101182826147bf565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615b8c5f395f51905f525d005b3461030457606036600319011261030457611061613634565b5061106a613608565b50611073613901565b50633ac4266d60e11b5f5260045ffd5b6080366003190112610304576004356001600160401b038111610304576110ae9036906004016135d8565b60243591906001600160401b03831161030457366023840112156103045782600401356110da816136f5565b936110e860405195866136d2565b8185526024602086019260051b820101903682116103045760248101925b82841061193f5750506044359150506001600160401b03811161030457611131903690600401613968565b926064356001600160401b03811161030457611151903690600401613968565b9161115a613f5c565b611162613f83565b8082511480611935575b8061192b575b1561191c573461190557909184925f9460be1981360301905b848710156118ea578660051b81013582811215610304578101966111af8185613a96565b51976111bb8289613a96565b51906111c78388613a96565b5199813590815f5260046020526111e98260ff600360405f2001541615613bfd565b815f52600560205260405f2090611215604085019260018060a01b0361120e85613a32565b1690614a25565b6001600160a01b0361122684613a32565b1690156118d85750825f526004602052600260405f20019360018554018095556020810135948086036118c157506001600160a01b0361126584613a32565b1695606082019e8f61127690613a32565b9360808401359860a085019561128c8787613a46565b369061129792613849565b805190602001206040519160208301937fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191485528b60408501528c60608501526080840152600160a01b600190031660a08301528b60c083015260e082015260e08152611305610100826136d2565b51902081519184518314806118b7575b156118a85760ff5f516020615a4c5f395f51905f5254169361133a8486811015614c40565b5f955f935f5b8681106117925750505050505050909d9e98999a9b9c9d61136391811015614c40565b5f61136e8383613a46565b15159050611789576113926113828a613a32565b8861138c87613a32565b886140c7565b93908093600d821015611655578b9260018b9314611748575b50505050600d8210156116555760018203611410575050506001956113e66113e05f516020615a6c5f395f51905f5293613a32565b91613a32565b906114006040519283928a8060a01b031697429184614017565b0390a45b0195949392919061118b565b909195845f5260076020526114288660405f20614cda565b1561173557811561172d5750965b61143f83613a32565b9161144982613a32565b6040516020989161145a8a836136d2565b5f82526040519561146a876136b7565b8887528a87018a90526001600160a01b039081166040880152166060860152608085018b905260a085015215611726576114a660015442613c17565b905b604051936114b585613666565b8452878401996114c5828c613bf1565b60408581019384525f88815260088b528181208a82528b52819020955180518755808b01516001880155908101516002870180546001600160a01b039283166001600160a01b03199182161790915560608301516003890180549190931691161790556080810151600487015560a00151805160058701916001600160401b0382116117125781906115578454613ac4565b8d601f82116116d8575b50508c90601f8311600114611674575f92611669575b50508160011b915f199060031b1c19161790555b600685019a5192600d841015611655576116207f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f97608097600760019f94611626956116479960ff8019835416911617905551910155895f5260068c5260405f208e8060a01b036115fb83613a32565b168f8060a01b03165f528c52600260405f2001611619858254613c17565b9055613a32565b94613a32565b98604051998c8060a01b03168a5289015242604089015260608801906139f4565b878060a01b031694a4611404565b634e487b7160e01b5f52602160045260245ffd5b015190505f80611577565b5f8581528e81209350601f198516908f5b8282106116c15750509084600195949392106116a9575b505050811b01905561158b565b01515f1960f88460031b161c191690555f808061169c565b60018596829396860151815501950193018f611685565b61170191865f52815f2090601f860160051c8201928610611708575b601f0160051c01906147a9565b5f8d611561565b90915081906116f4565b634e487b7160e01b5f52604160045260245ffd5b5f906114a8565b905096611436565b856307a5c91d60e31b5f5260045260245ffd5b611780955061177891929394506117719061176b6117658a613a32565b96613a32565b93613a46565b3691613849565b928888614a38565b88875f806113ab565b6113925f611382565b61183761182e60426117a26154bd565b6117aa6155c4565b6040519060208201925f516020615bec5f395f51905f528452604083015260608201524660808201523060a082015260a081526117e860c0826136d2565b5190206040519061190160f01b825260028201528560228201522060ff61180f858a613a96565b511661181b8588613a96565b5190611827868a613a96565b5192615933565b909291926159ab565b6001600160a01b038181169088161080611868575b61185a575b50600101611340565b6001998a0199909650611851565b506001600160a01b0381165f9081527fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa5602052604090205460ff1661184c565b63b3f07a3960e01b5f5260045ffd5b5083518314611315565b859063a6ab65c560e01b5f5260045260245260445ffd5b633142cb1160e11b5f5260045260245ffd5b5f5f516020615b8c5f395f51905f525d602060405160018152f35b63943892eb60e01b5f525f6004523460245260445ffd5b6329f517fd60e01b5f5260045ffd5b5080835114611172565b508085511461116c565b83356001600160401b0381116103045782013660438201121561030457602481013561196a816136f5565b9161197860405193846136d2565b818352602060248185019360051b830101019036821161030457604401915b8183106119af57505050815260209384019301611106565b823560ff8116810361030457815260209283019201611997565b34610304575f366003190112610304575f516020615aac5f395f51905f52541580611c0a575b15611bcd576040515f516020615a2c5f395f51905f5254815f611a1183613ac4565b8083529260018116908115611bae5750600114611b56575b611a35925003826136d2565b604051905f5f516020615a8c5f395f51905f5254611a5281613ac4565b8085529060018116908115611b345750600114611ae3575b5082611a7e611ab5946106e89303826136d2565b611ac360405191611a906020846136d2565b5f83525f368137604051958695600f60f81b875260e0602088015260e0870190613944565b908582036040870152613944565b904660608501523060808501525f60a085015283820360c0850152613911565b5f516020615a8c5f395f51905f525f90815291505f516020615c0c5f395f51905f525b818310611b1a575050820160200182611a6a565b600181602092949394548385890101520191019190611b06565b60ff191660208087019190915291151560051b85019091019150839050611a6a565b505f516020615a2c5f395f51905f525f90815290915f516020615acc5f395f51905f525b818310611b92575050906020611a3592820101611a29565b6020919350806001915483858801015201910190918392611b7a565b60209250611a3594915060ff191682840152151560051b820101611a29565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615bcc5f395f51905f5254156119ef565b3461030457604036600319011261030457611c39613608565b611c41613c71565b506004355f52600660205260405f209060018060a01b03165f5260205260e060405f20600460405191611c7383613681565b60018060a01b03815416835260ff600182015460018060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c0820152611cd460405180926138ac565bf35b34610304576020366003190112610304576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b818110611d24576106e8856106d4818703826136d2565b8254845260209093019260019283019201611d0d565b3461030457604036600319011261030457611d53613608565b6004355f52600960205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b34610304575f366003190112610304576020606454604051908152f35b3461030457608036600319011261030457600435611dbc613608565b6044356001600160401b038111610304573660238201121561030457611dec903690602481600401359101613849565b906064359060ff8216809203610304575f546001600160a01b031690811561203c57602091611e4f915f604051809681958294637c469ea160e11b84528b600485015260018060a01b0316998a6024850152608060448501526084840190613944565b90606483015203925af1908115612031575f91611ff7575b50611e70613e38565b82156104f8576001600160a01b03169081156104f85780156104f857611e9583614c6f565b611fa2575b825f526005602052611eaf8260405f20614cda565b15611f8f5781602093604051611ec481613681565b8281528581018481525f6040808401828152606085018381526080860184815260a080880186815260c089018781528a885260068f528688208c89528f5295909620975188546001600160a01b0319166001600160a01b03918216178955965160018901805495516001600160a81b03199096169190981617931515901b60ff60a01b1692909217855551600494929392611f6191151590613d1d565b516002850155516003840155519101555f516020615bac5f395f51905f52856040515f8152a4604051908152f35b506311dd05f360e31b5f5260045260245ffd5b611ff2604051611fb18161369c565b8481526003602082015f8152604083015f815260608401915f8352885f52600460205260405f20945185555160018501555160028401555115159101613ce4565b611e9a565b90506020813d602011612029575b81612012602093836136d2565b810103126103045761202390613c24565b83611e67565b3d9150612005565b6040513d5f823e3d90fd5b6315aeca0d60e11b5f5260045ffd5b60e0366003190112610304576024356001600160a01b03811690600435908281036103045761207861361e565b6064359390919060c4356001600160401b0381116103045761209e90369060040161387f565b936120a7613f5c565b6120b1838761417a565b6120b9613f83565b6001600160a01b0316948515610667573086146121d757600a5480159081156121cc575b50156121bd5761213b612145926120ff60019660a435908b60843591876141f6565b9790986040519461210f8661364a565b855286602086015233604086015260608501528960808501528860a08501528760c08501523691613849565b60e082015261437f565b14612165575f5f516020615b8c5f395f51905f525d602060405160018152f35b610dcf6121759261217b94613c17565b34613ca7565b80612188575b80806118ea565b6040516121a39161219a6020836136d2565b5f825233614740565b156121ae5780612181565b63559d141b60e11b5f5260045ffd5b633953b33f60e11b5f5260045ffd5b9050851115886120dd565b6304c821ad60e31b5f5260045ffd5b34610304575f36600319011261030457602060ff5f516020615b6c5f395f51905f5254166040519015158152f35b3461030457602036600319011261030457600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106123bb575050612263925003836136d2565b81519161226f836136f5565b9261227d60405194856136d2565b80845261228c601f19916136f5565b015f5b8181106123a45750505f5b8151811015612354575f838152600660205260409020600191906001600160a01b036122c68386613a96565b5116838060a01b03165f5260205260405f206004604051916122e783613681565b848060a01b03815416835260ff85820154868060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c08201526123428287613a96565b5261234d8186613a96565b500161229a565b836040518091602082016020835281518091526020604084019201905f5b818110612380575050500390f35b91935091602060e08261239660019488516138ac565b019401910191849392612372565b6020906123af613c71565b8282880101520161228f565b845483526001948501948794506020909301920161224e565b34610304576020366003190112610304577f1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f16020600435612413613e38565b80600a55604051908152a1005b60c036600319011261030457600435612437613608565b6064359160443591906084359060a4356001600160401b0381116103045761246390369060040161387f565b60355490959194906001600160a01b031680151590816128ed575b50156128df5773__$66aabcaa73966043eae8d94048b21b06c8$__96873b156103045760405163a3f11c4f60e01b81525f816004818c5af48015612031576128cf575b50604051634cfbad0360e01b8152936080856004818c5af4968715612031575f985f945f975f9a612879575b506124f6613f5c565b612500868561417a565b6001600160a01b0316948515610667573086146121d757600a54801590811561286e575b50156121bd576001600160a01b031660018114989094908882898c15612803575b9161255b84610dcf610dcf946125639796613c17565b14938a613c17565b90156127ec575060018060a01b03603254169560405197633199576360e11b89528460048a01528560248a01526060896044818b5afa978815612031575f995f996127ad575b50906020600492604051938480926396ce079560e01b82525afa918215612031575f92612779575b50806125dd8b85613c17565b116125e88b85613c17565b906127645750916126166126279261260f6126088d612689999897613ca7565b9b82613c17565b908b614b5b565b908082111561275b57508098613ca7565b9861272d575b61263a89898988886141f6565b5050603554604051611771916001600160a01b03166126588261364a565b86825287602083015260408201528760608201528860808201528960a08201528a60c082015261213b368486613849565b6020815191012096893b156103045760405198631587d80760e21b8a5260048a0152602489015260448801526064870152608486015260a485015260c484015260e48301526101048201525f8161012481855af480156120315761271d575b50803b15610304575f60049160405192838092632f60d07160e01b82525af480156120315761271357005b5f6105fd916136d2565b5f612727916136d2565b816126e8565b61275661273e8a610dcf8b8b613c17565b341434906127508c610dcf8d8d613c17565b90613c53565b61262d565b90508098613ca7565b638399ac8d60e01b5f5260045260245260445ffd5b9091506020813d6020116127a5575b81612795602093836136d2565b810103126103045751908e6125d1565b3d9150612788565b60209a506004929199506127d89060603d6060116127e5575b6127d081836136d2565b810190613c38565b9b91509a999192506125a9565b503d6127c6565b8790639df556d160e01b5f5260045260245260445ffd5b5050508b5f52600660205260405f20865f5260205260ff600160405f20015460a01c1680612848575b1561283957888289612545565b63be75e8c760e01b5f5260045ffd5b50845f52600660205260405f20865f5260205260ff600160405f20015460a01c1661282c565b90508211158d612524565b95509850985094506080833d6080116128c7575b8161289a608093836136d2565b81010312610304578251946128b160208501613c24565b986060604086015195015196999496988c6124ed565b3d915061288d565b5f6128d9916136d2565b886124c1565b62a81d6360e11b5f5260045ffd5b905033148861247e565b34610304575f366003190112610304577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361294e5760206040515f516020615aec5f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b346103045760603660031901126103045760243560043561297c613dc9565b805f52600760205261299582610d248160405f20614a25565b7fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594960206129c460443542613c17565b835f526008825260405f20855f52825280600760405f200155604051908152a3005b6040366003190112610304576129fa613634565b6024356001600160401b038111610304573660238201121561030457612a2a903690602481600401359101613849565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115612ba2575b5061294e57612a6c613d5a565b6040516352d1902d60e01b81526001600160a01b0383169290602081600481875afa5f9181612b6e575b50612aae5783634c9c8ce360e01b5f5260045260245ffd5b805f516020615aec5f395f51905f52859203612b5c5750813b15612b4a575f516020615aec5f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115612b32575f808360206105fd95519101845af4612b2c6146c2565b9161587d565b505034612b3b57005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d602011612b9a575b81612b8a602093836136d2565b8101031261030457519085612a96565b3d9150612b7d565b5f516020615aec5f395f51905f52546001600160a01b03161415905083612a5f565b3461030457612bd236613818565b612bda613f5c565b612be2613f83565b815f526004602052612bff8260ff600360405f2001541615613bfd565b815f526007602052612c1881610d248160405f20614a25565b815f52600860205260405f20815f5260205260405f20604051612c3a81613666565b612c4382613afc565b8152600760ff60068401541692612c5e602084019485613bf1565b015491604082019283528051600d81101561165557600214159081612d6a575b81612d50575b50905160408101516080820151606090920151612cb193926001600160a01b039182169290911687614128565b50600d8110156116555760016040519180602084015260208352612cd66040846136d2565b03612d285750519182158015612d1f575b15612d0857612cf69250614035565b5f5f516020615b8c5f395f51905f525d005b82637a88099160e11b5f526004524260245260445ffd5b50428310612ce7565b60405162461bcd60e51b815260206004820152908190612d4c906024830190613944565b0390fd5b90505190600d8210156116555790600a1415612cb1612c84565b80915051600d811015611655576009141590612c7e565b34610304575f366003190112610304575f546040516001600160a01b039091168152602090f35b3461030457602036600319011261030457600435612dc4613d5a565b61c3508110612e065760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91603654908060365582519182526020820152a1005b63050779f960e21b5f5260045ffd5b3461030457606036600319011261030457602435600435612e3461361e565b612e3c613dc9565b612e44613f83565b815f526007602052612e5d83610d248160405f20614a25565b815f52600860205260405f20835f5260205260405f20600760405191612e8283613666565b612e8b81613afc565b8352612ea160ff60068301541660208501613bf1565b015460408201525160016080818060a01b03604084015116920151911480612fca575b612fad575b5080612ed584846147bf565b916001600160a01b031615612f99575b8151906020830190815193612f17604082019460018060a01b038651169684608085019889519260a087015194614a38565b50600d8110156116555760018103612f815750905f516020615a6c5f395f51905f52915192519360018060a01b039051169451612f5b604051928392429184614017565b0390a45f516020615a0c5f395f51905f525f80a35f5f516020615b8c5f395f51905f525d005b63290cd55f60e01b5f52612f94906139e6565b60245ffd5b5060608101516001600160a01b0316612ee5565b612fb990610dcf613d3a565b60645490818111610de55750612ec9565b506065548314612ec4565b34610304575f36600319011261030457602060ff5f516020615a4c5f395f51905f525416604051908152f35b346103045761300f36613769565b908051825103610c55575f5b82518110156105fd578061305461303460019385613a96565b51838060a01b036130458488613a96565b5116906105f86105f382613a78565b500161301b565b3461030457604036600319011261030457613074613608565b336001600160a01b03821603613090576105fd90600435613fb8565b63334bd91960e11b5f5260045ffd5b34610304576040366003190112610304576004356130bb613608565b6130c76105f383613a78565b6130d181836148eb565b6130d757005b815f525f516020615b0c5f395f51905f5260205261310260405f209160018060a01b03168092614cda565b15610c4057005b3461030457602036600319011261030457602061071c600435613a78565b34610304575f366003190112610304576036548061314e57506020620249f0604051908152f35b60209061071c565b34610304575f366003190112610304576020600a54604051908152f35b6040366003190112610304576004356001600160401b0381116103045761319e9036906004016135d8565b6024356001600160401b03811161030457366023820112156103045760048101356001600160401b0381116103045736602460e0830284010111610304576131e4613f5c565b6131ec613f83565b335f9081527f7bf51035c163aa501aba0b1731707347fc5168978f3fc3c84ea475ee63d4b83f602052604090205460ff16156134e257346119055780830361191c575f5b838110613249575f5f516020615b8c5f395f51905f525d005b613254818587613a10565b359061326c602061326683888a613a10565b01613a32565b61327c606061326684898b613a10565b92608061328a84898b613a10565b01359360a061329a858a8c613a10565b01359260c06132aa868b8d613a10565b013590878610156134ce578a968a936132db6132d1898d60e082020197602489019c613a10565b60e0810190613a46565b6001600160a01b0385169a91959094916132f58c8a61417a565b6001600160a01b0361330682613a32565b168c14906001600160a01b039061331c90613a32565b1690156134b8575060448701986001600160a01b0361333a8b613a32565b6001600160a01b03909516941684036134a9573084146121d757600a54801590811561349e575b50156121bd5784613372938a6141f6565b92909160648701359661339e8861338d87610dcf8888613c17565b81101561275088610dcf8989613c17565b6032546001600160a01b0316906133b48b613a32565b9160a48201359160ff831680930361030457813b15610304578e926040519b8c958695637f40ec1760e11b87526004870152600160a01b6001900316602486015230604486015260648501526084820135608485015260a484015260c481013560c484015260e4013560e48301525a925f610104928195f19a8b156120315761344c6134889a61213b9960019e61348e575b50613a32565b906040519961345a8b61364a565b8a5260208a01528b8060a01b031660408901526060880152608087015260a086015260c08501523691613849565b01613230565b5f613498916136d2565b5f613446565b90508611155f613361565b630672aec160e01b5f5260045ffd5b8b6313f7c32b60e31b5f5260045260245260445ffd5b634e487b7160e01b5f52603260045260245ffd5b63e2517d3f60e01b5f52336004527f6b8b15f1c11543d8280deaa7c24d12fffba6a357e4428e8c43e4234790186bff60245260445ffd5b34610304576020366003190112610304576004356001600160a01b0381169081900361030457613547613d5a565b801561066757603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b34610304576020366003190112610304576004359063ffffffff60e01b821680920361030457602091637965db0b60e01b81149081156135c7575b5015158152f35b6301ffc9a760e01b149050836135c0565b9181601f84011215610304578235916001600160401b038311610304576020808501948460051b01011161030457565b602435906001600160a01b038216820361030457565b604435906001600160a01b038216820361030457565b600435906001600160a01b038216820361030457565b61010081019081106001600160401b0382111761171257604052565b606081019081106001600160401b0382111761171257604052565b60e081019081106001600160401b0382111761171257604052565b608081019081106001600160401b0382111761171257604052565b60c081019081106001600160401b0382111761171257604052565b601f909101601f19168101906001600160401b0382119082101761171257604052565b6001600160401b0381116117125760051b60200190565b9080601f83011215610304578135613723816136f5565b9261373160405194856136d2565b81845260208085019260051b82010192831161030457602001905b8282106137595750505090565b813581526020918201910161374c565b906040600319830112610304576004356001600160401b03811161030457826137949160040161370c565b91602435906001600160401b0382116103045780602383011215610304578160040135906137c1826136f5565b926137cf60405194856136d2565b8284526024602085019360051b82010191821161030457602401915b8183106137f85750505090565b82356001600160a01b0381168103610304578152602092830192016137eb565b6040906003190112610304576004359060243590565b6001600160401b03811161171257601f01601f191660200190565b9291926138558261382e565b9161386360405193846136d2565b829481845281830111610304578281602093845f960137010152565b9181601f84011215610304578235916001600160401b038311610304576020838186019501011161030457565b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b6044359060ff8216820361030457565b90602080835192838152019201905f5b81811061392e5750505090565b8251845260209384019390920191600101613921565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9080601f8301121561030457813561397f816136f5565b9261398d60405194856136d2565b81845260208085019260051b820101918383116103045760208201905b8382106139b957505050505090565b81356001600160401b038111610304576020916139db8784809488010161370c565b8152019101906139aa565b600d81101561165557600452565b90600d8210156116555752565b60243590811515820361030457565b91908110156134ce5760051b8101359060fe1981360301821215610304570190565b356001600160a01b03811681036103045790565b903590601e198136030182121561030457018035906001600160401b0382116103045760200191813603831361030457565b5f525f516020615b4c5f395f51905f52602052600160405f20015490565b80518210156134ce5760209160051b010190565b15613ab25750565b6373a1390160e11b5f5260045260245ffd5b90600182811c92168015613af2575b6020831014613ade57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613ad3565b90604051613b09816136b7565b60058193805483526001810154602084015260018060a01b03600282015416604084015260018060a01b03600382015416606084015260048101546080840152019060405180925f90805490613b5e82613ac4565b8085529160018116908115613bca5750600114613b8b575b505060a09291613b879103846136d2565b0152565b5f908152602081209092505b818310613bae575050810160200181613b87613b76565b6020919350806001915483858901015201910190918492613b97565b60ff191660208681019190915292151560051b85019092019250839150613b879050613b76565b600d8210156116555752565b15613c055750565b636fc24b7960e11b5f5260045260245ffd5b919082018092116105b357565b51906001600160a01b038216820361030457565b90816060910312610304578051916040602083015192015190565b15613c5c575050565b63943892eb60e01b5f5260045260245260445ffd5b60405190613c7e82613681565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b919082039182116105b357565b60405190613cc1826136b7565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b9060ff801983541691151516179055565b15613cfd5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6065545f52600660205260405f2060015f52602052600460405f20015490565b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff1615613d9257565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b335f9081527fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143602052604090205460ff1615613e0157565b63e2517d3f60e01b5f52336004527f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0960245260445ffd5b335f9081527f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f602052604090205460ff1615613e7057565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b335f9081527f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456602052604090205460ff1615613edf57565b63e2517d3f60e01b5f52336004527f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92960245260445ffd5b5f8181525f516020615b4c5f395f51905f526020908152604080832033845290915290205460ff1615613f465750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615b6c5f395f51905f525416613f7457565b63d93c066560e01b5f5260045ffd5b5f516020615b8c5f395f51905f525c613fa95760015f516020615b8c5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b919091613fc58382614989565b9283613fcf575050565b815f525f516020615b0c5f395f51905f52602052613ffa60405f209160018060a01b03168092614d21565b15614003575050565b62a95f1b60e31b5f5260045260245260445ffd5b604091949392606082019560018060a01b0316825260208201520152565b9061403f916147bf565b60018060a01b03606082015116815190602083019081519361407e604082019460018060a01b038651169684608085019889519260a087015194614a38565b50600d8110156116555760018103612f815750905f516020615a6c5f395f51905f52915192519360018060a01b0390511694516140c2604051928392429184614017565b0390a4565b909291906001600160a01b0384166001148061411d575b6140f4575b906140f0935f9392614a64565b9091565b919061410281610dcf613d3a565b606454106141115790916140e3565b50505050600990600190565b5060655481146140de565b939291906001600160a01b0381166001148061416f575b61414d575b6140f094614a64565b61415982610dcf613d3a565b6064541015614144575050505050600990600190565b50606554851461413f565b5f8181526005602052604090206001600160a01b038316926141a19161081b908590614a25565b805f5260046020526141be8160ff600360405f2001541615613bfd565b5f52600660205260405f20815f5260205260ff600160405f20015460a81c166141e45750565b6338384f6f60e11b5f5260045260245ffd5b5f8181526006602090815260408083206001600160a01b0390951680845294909152908190209051929694959260c0919061423081613681565b81546001600160a01b0390811682526001830154908116602083015260a081811c60ff908116158015604086015260a89390931c161515606084015260028401546080840152600384015490830152600490920154920182905261436b575b506032546001600160a01b031695861561435c57606460609260405198899384926337dba1f760e21b8452600484015260248301528660448301525afa908115612031575f955f905f93614336575b5082958197851515958661432b575b505084614320575b505082614315575b50501561430657565b6358e8878560e01b5f5260045ffd5b101590505f806142fd565b101592505f806142f5565b101594505f806142ed565b91965050614353915060603d6060116127e5576127d081836136d2565b9195905f6142de565b6330d45fb160e01b5f5260045ffd5b808480614379931015613c53565b5f61428f565b80515f526004602052600160405f2001600181540180915581515f52600660205260405f20602083019060018060a01b0382511660018060a01b03165f5260205260405f206001808060a01b03910154169280519360018060a01b0383511694604083019060018060a01b03825116966080850191825160a08701998a519361440e60c08a0195865190613c17565b93805f52600660205260405f2060018060a01b0383165f5260205260405f2060ff600182015460a01c165f146146ad5760030161444c858254613c17565b90555b6001820361453a57505050806144688361447693613c17565b341015612750843493613c17565b80614503575b50905f516020615c2c5f395f51905f52946144f8925b81519760018060a01b039051169460018060a01b039051169960e060018060a01b03606085015116945191519251930151936040519788978852602088015260408701526060860152608085015260a084015261010060c0840152610100830190613944565b4260e08301520390a4565b60335460405192939261452e9290916001600160a01b03166145266020846136d2565b5f8352614740565b156121ae57905f61447c565b909192939895969798346119055761458c6145558287613c17565b604051906323b872dd60e01b60208301528660248301523060448301526064820152606481526145866084826136d2565b846158db565b8061465d575b505f52600660205260405f2060018060a01b0382165f5260205260ff600160405f20015460a01c16156145de575b505050916144f8915f516020615c2c5f395f51905f52959493614492565b604051632770a7eb60e21b8152602081806145fd873060048401614c25565b03815f865af1908115612031575f9161462e575b506145c057639ac2f56d60e01b5f5260045260245260445260645ffd5b614650915060203d602011614656575b61464881836136d2565b810190614c0d565b5f614611565b503d61463e565b60335460405163a9059cbb60e01b60208201526146a79290916146a191839161469391906001600160a01b031660248401614c25565b03601f1981018352826136d2565b836158db565b5f614592565b6004016146bb858254613ca7565b905561444f565b3d156146ec573d906146d38261382e565b916146e160405193846136d2565b82523d5f602084013e565b606090565b81516001600160a01b03909116915f9182916020018285620186a0f16147156146c2565b901561473a57805161473157503b1561472d57600190565b5f90565b60209150015190565b50505f90565b81471061479a5782516001600160a01b03909116925f9182916020018486620186a0f19061476c6146c2565b9115614793571561477f575b5050600190565b805161473157503b1561472d575f80614778565b5050505f90565b632b60b36f60e21b5f5260045ffd5b8181106147b4575050565b5f81556001016147a9565b91906147c9613cb4565b50825f5260076020526147df8160405f20614d21565b156148d957825f52600860205260405f20815f5260205261480260405f20613afc565b5f848152600660209081526040808320818501516001600160a01b031684529091529020608082015160029091018054929592909161484091613ca7565b90555f52600860205260405f20905f526020525f600760408220828155826001820155826002820155826003820155826004820155600581016148838154613ac4565b9081614896575b50508260068201550155565b81601f8693116001146148ad5750555b5f8061488a565b818352602083206148c991601f0160051c8101906001016147a9565b80825281602081209155556148a6565b637f11bea960e01b5f5260045260245ffd5b5f8181525f516020615b4c5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1661473a575f8181525f516020615b4c5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5f8181525f516020615b4c5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff161561473a575f8181525f516020615b4c5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001915f520160205260405f2054151590565b9594939291906001600160a01b0383163014614a58576140f09596614de6565b505050925050600c9190565b60325493955f9590946001600160a01b031691821561435c575f9081526009602090815260408083206001600160a01b039096168084529590915290205460ff16614b4e57614b4557602092915f60649260405198899586946319ab07bb60e21b86526004860152602485015260018060a01b031660448401525af1928315612031575f93614b08575b5082600d81101561165557600103614b0257565b60019150565b9092506020813d602011614b3d575b81614b24602093836136d2565b810103126103045751600d81101561030457915f614aee565b3d9150614b17565b50600194505050565b50600295505f9450505050565b91818302915f1981850993838086109503948086039514614beb5784831115614bd35790829109815f0382168092046002816003021880820260020302808202600203028082026002030280820260020302808202600203028091026002030293600183805f03040190848311900302920304170290565b82634e487b715f52156003026011186020526024601cfd5b505080925015614bf9570490565b634e487b7160e01b5f52601260045260245ffd5b90816020910312610304575180151581036103045790565b6001600160a01b039091168152602081019190915260400190565b15614c485750565b631aebd17960e11b5f5260045260245ffd5b80548210156134ce575f5260205f2001905f90565b805f52600360205260405f2054155f14614cd557600254600160401b81101561171257614cbe614ca88260018594016002556002614c5a565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b614ce48282614a25565b61473a57805490600160401b8210156117125782614d0c614ca8846001809601855584614c5a565b90558054925f520160205260405f2055600190565b906001820191815f528260205260405f20548015155f14614dde575f1981018181116105b35782545f198101919082116105b357818103614da9575b50505080548015614d95575f190190614d768282614c5a565b8154905f199060031b1b19169055555f526020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b614dc9614db9614ca89386614c5a565b90549060031b1c92839286614c5a565b90555f528360205260405f20555f8080614d5d565b505050505f90565b94919290936080525f935f90865f52600660205260405f2060018060a01b0360018060a01b038716165f5260205260ff600160405f20015460a01c16935f9160018060a01b036035541691601881511015806154b4575b806154aa575b614eb8575b505050614eaf5750608051614e5e9291846157a6565b92600d841015614e9b5760018414614e7a575050509060805190565b60805192949350614e9592916001600160a01b031690615756565b60019190565b634e487b7160e01b83526021600452602483fd5b95939450505050565b602081015160601c908a6034820151918a60405196614ed68861369c565b5f88525f60208901525f6040890152606080890152614f03858760805160018060a01b0386168b896156a6565b151588525f5f80604051602081019062483e5560e91b82528a602482015260248152614f306044826136d2565b5190855afa614f3d6146c2565b908061549e575b615492575b50614ffc575b5050505050508151614f62575b80614e48565b602082015115156060604084015115159301519173__$66aabcaa73966043eae8d94048b21b06c8$__90813b15610304578b614fd25f9560405197889687958695631b24662760e11b8752600487015260248601526044850152606484015260a0608484015260a4830190613944565b03915af4801561203157614fe7575b80614f5c565b614ff49196505f906136d2565b5f945f614fe1565b6001600160a01b03821660011460408901819052928315615353575b8b60408a015161502a575b5050614f4f565b89865f93859c9e9a9c85948d9f9b9d6150c2906036548015891461534e5750620249f05b805a10891461533a57506001602089985b01528b15881461532e5761469388955b6040519485936020850199631eeed51360e01b8b52602486015260448501528d60018060a01b0316606485015260018060a01b0316608484015260805160a484015260c060c484015260e4830190613944565b5193f16150cd6146c2565b60608c0152151560408b015282156152df575b505060408801511515806152cf575b156151e857505050506060840151908151820191604081602085019403126103045760208101516040820151926001600160401b0384116103045784603f85850101121561030457602084840101516151478161382e565b9561515560405197886136d2565b8187526040858701830101116103045787955f8f96836151bb9460406020935f516020615b2c5f395f51905f529a01018386015e8301015260405193849363ffffffff60e01b168452600160208501526040840152608060608401526080830190613944565b0390a46080516151d5906001600160a01b03871689615756565b60019182915b5f808a818a81808b615023565b93848795989694989793971590156152c7575b61524e575b5f516020615b2c5f395f51905f529060608701519061524660405192839263ffffffff60e01b1683525f60208401525f6040840152608060608401526080830190613944565b0390a46151db565b915050604051632770a7eb60e21b81526020818b815f816152756080513060048401614c25565b03926001600160a01b03165af180156120315784928d925f516020615b2c5f395f51905f52926152a8575b509050615200565b6152c09060203d6020116146565761464881836136d2565b505f6152a0565b5089156151fb565b50604060608901515110156150ef565b5f91829182604051602081019263095ea7b360e01b84526024820152816044820152604481526153106064826136d2565b51926001600160a01b03165af1506153266146c2565b508a5f6150e0565b6146936080519561506f565b60206153486001925a613ca7565b9861505f565b61504e565b8c1580615421575b5f80604051602081019063095ea7b360e01b8252615383816146936080518a60248401614c25565b5190826001600160a01b0389165af161539a6146c2565b501590811560408c015281615419575b50156150185791935060206040518095632770a7eb60e21b8252815f816153d76080513060048401614c25565b03926001600160a01b03165af1918215612031578f948e936153fa575b50615018565b6154129060203d6020116146565761464881836136d2565b505f6153f4565b90505f6153aa565b92909450602060405180926340c10f1960e01b8252815f816154496080513060048401614c25565b03926001600160a01b03165af1908115612031575f91615473575b501561479a578e938d9261535b565b61548c915060203d6020116146565761464881836136d2565b5f615464565b6020915001515f614f49565b50602081511015614f44565b50823b1515614e43565b50821515614e3d565b6040515f516020615a2c5f395f51905f5254905f816154db84613ac4565b9182825260208201946001811690815f146155a85750600114615550575b615505925003826136d2565b51908115615511572090565b50505f516020615aac5f395f51905f5254801561552b5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b505f516020615a2c5f395f51905f525f90815290915f516020615acc5f395f51905f525b81831061558c575050906020615505928201016154f9565b6020919350806001915483858801015201910190918392615574565b60ff191686525061550592151560051b820160200190506154f9565b6040515f516020615a8c5f395f51905f5254905f816155e284613ac4565b9182825260208201946001811690815f1461568a5750600114615632575b61560c925003826136d2565b51908115615618572090565b50505f516020615bcc5f395f51905f5254801561552b5790565b505f516020615a8c5f395f51905f525f90815290915f516020615c0c5f395f51905f525b81831061566e57505090602061560c92820101615600565b6020919350806001915483858801015201910190918392615656565b60ff191686525061560c92151560051b82016020019050615600565b92949293919290916001600160a01b031630148015919061573a575b50614dde5773__$66aabcaa73966043eae8d94048b21b06c8$__92833b15610304575f936084926040519687958694631ff65d3360e31b8652600486015260018060a01b03166024850152604484015260648301525af480156120315761572a575b50600190565b5f615734916136d2565b5f615724565b6001600160e01b031916632afe83bb60e11b141590505f6156c2565b5f52600660205260405f209060018060a01b03165f5260205260405f2060ff600182015460a01c165f146157975760036157939101918254613ca7565b9055565b60046157939101918254613c17565b6001600160a01b03169291906001841461583857811561582f576157f592156158035760405163a9059cbb60e01b6020820152916157ed9183916146939160248401614c25565b6005926146f1565b156158005750600190565b90565b6040516340c10f1960e01b6020820152916158279183916146939160248401614c25565b6006926146f1565b50505050600190565b915080925047106158765761586791604051916158566020846136d2565b5f83526001600160a01b0316614740565b1561587157600190565b600590565b5050600b90565b906158a1575080511561589257805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806158d2575b6158b2575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156158aa565b905f602091828151910182855af115612031575f513d61592a57506001600160a01b0381163b155b61590a5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615903565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116159a0579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612031575f516001600160a01b0381161561599657905f905f90565b505f906001905f90565b5050505f9160039190565b600481101561165557806159bd575050565b600181036159d45763f645eedf60e01b5f5260045ffd5b600281036159ef575063fce698f760e01b5f5260045260245ffd5b6003146159f95750565b6335e2f38360e21b5f5260045260245ffdfeb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5dba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a164736f6c634300081c000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// CrossBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeMetaData.ABI instead.
var CrossBridgeABI = CrossBridgeMetaData.ABI

// Deprecated: Use CrossBridgeMetaData.Sigs instead.
// CrossBridgeFuncSigs maps the 4-byte function signature to its string representation.
var CrossBridgeFuncSigs = CrossBridgeMetaData.Sigs

// CrossBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossBridgeMetaData.Bin instead.
var CrossBridgeBin = CrossBridgeMetaData.Bin

// DeployCrossBridge deploys a new Ethereum contract, binding an instance of CrossBridge to it.
func DeployCrossBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossBridge, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// CrossBridge is an auto generated Go binding around an Ethereum contract.
type CrossBridge struct {
	CrossBridgeCaller     // Read-only binding to the contract
	CrossBridgeTransactor // Write-only binding to the contract
	CrossBridgeFilterer   // Log filterer for contract events
}

// CrossBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossBridgeSession struct {
	Contract     *CrossBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossBridgeCallerSession struct {
	Contract *CrossBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrossBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossBridgeTransactorSession struct {
	Contract     *CrossBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossBridgeRaw struct {
	Contract *CrossBridge // Generic contract binding to access the raw methods on
}

// CrossBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossBridgeCallerRaw struct {
	Contract *CrossBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossBridgeTransactorRaw struct {
	Contract *CrossBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossBridge creates a new instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridge(address common.Address, backend bind.ContractBackend) (*CrossBridge, error) {
	contract, err := bindCrossBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// NewCrossBridgeCaller creates a new read-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeCaller(address common.Address, caller bind.ContractCaller) (*CrossBridgeCaller, error) {
	contract, err := bindCrossBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCaller{contract: contract}, nil
}

// NewCrossBridgeTransactor creates a new write-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossBridgeTransactor, error) {
	contract, err := bindCrossBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTransactor{contract: contract}, nil
}

// NewCrossBridgeFilterer creates a new log filterer instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossBridgeFilterer, error) {
	contract, err := bindCrossBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeFilterer{contract: contract}, nil
}

// bindCrossBridge binds a generic wrapper to an already deployed contract.
func bindCrossBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.CrossBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridge.Contract.DEFAULTADMINROLE(&_CrossBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridge.Contract.DEFAULTADMINROLE(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeExecutor() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecutor(&_CrossBridge.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeExecutor() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecutor(&_CrossBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeVerifier() (common.Address, error) {
	return _CrossBridge.Contract.BridgeVerifier(&_CrossBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _CrossBridge.Contract.BridgeVerifier(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) CrossSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) CrossSupply() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupply(&_CrossBridge.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) CrossSupply() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupply(&_CrossBridge.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) CrossSupplyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossSupplyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupplyLimit(&_CrossBridge.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupplyLimit(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossBridge *CrossBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridge.Contract.GetRoleAdmin(&_CrossBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridge.Contract.GetRoleAdmin(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridge *CrossBridgeCaller) Initialize(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "initialize", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridge *CrossBridgeSession) Initialize(arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.CallOpts, arg0, arg1, arg2)
}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridge *CrossBridgeCallerSession) Initialize(arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.CallOpts, arg0, arg1, arg2)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridge.Contract.IsTokenFinalizePaused(&_CrossBridge.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridge.Contract.IsTokenFinalizePaused(&_CrossBridge.CallOpts, remoteChainID, token)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) MaxExtraDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "maxExtraDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) MaxExtraDataLength() (*big.Int, error) {
	return _CrossBridge.Contract.MaxExtraDataLength(&_CrossBridge.CallOpts)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) MaxExtraDataLength() (*big.Int, error) {
	return _CrossBridge.Contract.MaxExtraDataLength(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) PostCallGasReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "postCallGasReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) PostCallGasReserve() (*big.Int, error) {
	return _CrossBridge.Contract.PostCallGasReserve(&_CrossBridge.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) PostCallGasReserve() (*big.Int, error) {
	return _CrossBridge.Contract.PostCallGasReserve(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridge.Contract.SupportsInterface(&_CrossBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridge.Contract.SupportsInterface(&_CrossBridge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCallerSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridge *CrossBridgeTransactor) BridgeTokenForwarded(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "bridgeTokenForwarded", toChainID, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridge *CrossBridgeSession) BridgeTokenForwarded(toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeTokenForwarded(&_CrossBridge.TransactOpts, toChainID, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) BridgeTokenForwarded(toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeTokenForwarded(&_CrossBridge.TransactOpts, toChainID, to, value, networkFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRole(&_CrossBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRole(&_CrossBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeTransactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePendingWithRecipient(&_CrossBridge.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePendingWithRecipient(&_CrossBridge.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RemovePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RemovePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceRole(&_CrossBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceRole(&_CrossBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRole(&_CrossBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRole(&_CrossBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeTransactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecutor(&_CrossBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecutor(&_CrossBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeVerifier(&_CrossBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeVerifier(&_CrossBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetChainPause(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetChainPause(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossSupplyLimit(opts *bind.TransactOpts, _crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossSupplyLimit", _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossSupplyLimit(&_CrossBridge.TransactOpts, _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossSupplyLimit(&_CrossBridge.TransactOpts, _crossSupplyLimit)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridge *CrossBridgeTransactor) SetMaxExtraDataLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setMaxExtraDataLength", length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridge *CrossBridgeSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetMaxExtraDataLength(&_CrossBridge.TransactOpts, length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetMaxExtraDataLength(&_CrossBridge.TransactOpts, length)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPostCallGasReserve(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPostCallGasReserve", value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridge *CrossBridgeSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPostCallGasReserve(&_CrossBridge.TransactOpts, value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPostCallGasReserve(&_CrossBridge.TransactOpts, value)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelayExpiration(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelayExpiration(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// CrossBridgeBridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the CrossBridge contract.
type CrossBridgeBridgeExecutorSetIterator struct {
	Event *CrossBridgeBridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeExecutorSet)
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
		it.Event = new(CrossBridgeBridgeExecutorSet)
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
func (it *CrossBridgeBridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeExecutorSet represents a BridgeExecutorSet event raised by the CrossBridge contract.
type CrossBridgeBridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*CrossBridgeBridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeExecutorSetIterator{contract: _CrossBridge.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeExecutorSet)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeExecutorSet(log types.Log) (*CrossBridgeBridgeExecutorSet, error) {
	event := new(CrossBridgeBridgeExecutorSet)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizedIterator struct {
	Event *CrossBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFinalized)
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
		it.Event = new(CrossBridgeBridgeFinalized)
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
func (it *CrossBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFinalized represents a BridgeFinalized event raised by the CrossBridge contract.
type CrossBridgeBridgeFinalized struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFinalizedIterator{contract: _CrossBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFinalized)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFinalized(log types.Log) (*CrossBridgeBridgeFinalized, error) {
	event := new(CrossBridgeBridgeFinalized)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the CrossBridge contract.
type CrossBridgeBridgeInitiatedIterator struct {
	Event *CrossBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeInitiated)
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
		it.Event = new(CrossBridgeBridgeInitiated)
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
func (it *CrossBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeInitiated represents a BridgeInitiated event raised by the CrossBridge contract.
type CrossBridgeBridgeInitiated struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*CrossBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeInitiatedIterator{contract: _CrossBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeInitiated)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeInitiated(log types.Log) (*CrossBridgeBridgeInitiated, error) {
	event := new(CrossBridgeBridgeInitiated)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the CrossBridge contract.
type CrossBridgeBridgePendingIterator struct {
	Event *CrossBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgePending)
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
		it.Event = new(CrossBridgeBridgePending)
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
func (it *CrossBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgePending represents a BridgePending event raised by the CrossBridge contract.
type CrossBridgeBridgePending struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgePendingIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgePendingIterator{contract: _CrossBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgePending)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgePending(log types.Log) (*CrossBridgeBridgePending, error) {
	event := new(CrossBridgeBridgePending)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the CrossBridge contract.
type CrossBridgeBridgeVerifierSetIterator struct {
	Event *CrossBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeVerifierSet)
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
		it.Event = new(CrossBridgeBridgeVerifierSet)
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
func (it *CrossBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the CrossBridge contract.
type CrossBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*CrossBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeVerifierSetIterator{contract: _CrossBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeVerifierSet)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*CrossBridgeBridgeVerifierSet, error) {
	event := new(CrossBridgeBridgeVerifierSet)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the CrossBridge contract.
type CrossBridgeChainPauseSetIterator struct {
	Event *CrossBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeChainPauseSet)
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
		it.Event = new(CrossBridgeChainPauseSet)
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
func (it *CrossBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeChainPauseSet represents a ChainPauseSet event raised by the CrossBridge contract.
type CrossBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*CrossBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeChainPauseSetIterator{contract: _CrossBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeChainPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseChainPauseSet(log types.Log) (*CrossBridgeChainPauseSet, error) {
	event := new(CrossBridgeChainPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSetIterator struct {
	Event *CrossBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
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
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*CrossBridgeCrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossMintableERC20CodeSetIterator{contract: _CrossBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossMintableERC20CodeSet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*CrossBridgeCrossMintableERC20CodeSet, error) {
	event := new(CrossBridgeCrossMintableERC20CodeSet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossSupplyLimitSetIterator is returned from FilterCrossSupplyLimitSet and is used to iterate over the raw logs and unpacked data for CrossSupplyLimitSet events raised by the CrossBridge contract.
type CrossBridgeCrossSupplyLimitSetIterator struct {
	Event *CrossBridgeCrossSupplyLimitSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeCrossSupplyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossSupplyLimitSet)
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
		it.Event = new(CrossBridgeCrossSupplyLimitSet)
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
func (it *CrossBridgeCrossSupplyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossSupplyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossSupplyLimitSet represents a CrossSupplyLimitSet event raised by the CrossBridge contract.
type CrossBridgeCrossSupplyLimitSet struct {
	CrossSupplyLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCrossSupplyLimitSet is a free log retrieval operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossSupplyLimitSet(opts *bind.FilterOpts) (*CrossBridgeCrossSupplyLimitSetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossSupplyLimitSetIterator{contract: _CrossBridge.contract, event: "CrossSupplyLimitSet", logs: logs, sub: sub}, nil
}

// WatchCrossSupplyLimitSet is a free log subscription operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossSupplyLimitSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossSupplyLimitSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossSupplyLimitSet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
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

// ParseCrossSupplyLimitSet is a log parse operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) ParseCrossSupplyLimitSet(log types.Log) (*CrossBridgeCrossSupplyLimitSet, error) {
	event := new(CrossBridgeCrossSupplyLimitSet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeDevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the CrossBridge contract.
type CrossBridgeDevSetIterator struct {
	Event *CrossBridgeDevSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeDevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeDevSet)
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
		it.Event = new(CrossBridgeDevSet)
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
func (it *CrossBridgeDevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeDevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeDevSet represents a DevSet event raised by the CrossBridge contract.
type CrossBridgeDevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridge *CrossBridgeFilterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*CrossBridgeDevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeDevSetIterator{contract: _CrossBridge.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridge *CrossBridgeFilterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeDevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeDevSet)
				if err := _CrossBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseDevSet(log types.Log) (*CrossBridgeDevSet, error) {
	event := new(CrossBridgeDevSet)
	if err := _CrossBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChangedIterator struct {
	Event *CrossBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeEIP712DomainChanged)
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
		it.Event = new(CrossBridgeEIP712DomainChanged)
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
func (it *CrossBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeEIP712DomainChangedIterator{contract: _CrossBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeEIP712DomainChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*CrossBridgeEIP712DomainChanged, error) {
	event := new(CrossBridgeEIP712DomainChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the CrossBridge contract.
type CrossBridgeExtraCallExecutedIterator struct {
	Event *CrossBridgeExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeExtraCallExecuted)
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
		it.Event = new(CrossBridgeExtraCallExecuted)
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
func (it *CrossBridgeExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeExtraCallExecuted represents a ExtraCallExecuted event raised by the CrossBridge contract.
type CrossBridgeExtraCallExecuted struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (*CrossBridgeExtraCallExecutedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeExtraCallExecutedIterator{contract: _CrossBridge.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_CrossBridge *CrossBridgeFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *CrossBridgeExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeExtraCallExecuted)
				if err := _CrossBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseExtraCallExecuted(log types.Log) (*CrossBridgeExtraCallExecuted, error) {
	event := new(CrossBridgeExtraCallExecuted)
	if err := _CrossBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossBridge contract.
type CrossBridgeInitializedIterator struct {
	Event *CrossBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeInitialized)
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
		it.Event = new(CrossBridgeInitialized)
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
func (it *CrossBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeInitialized represents a Initialized event raised by the CrossBridge contract.
type CrossBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossBridgeInitializedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeInitializedIterator{contract: _CrossBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeInitialized)
				if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseInitialized(log types.Log) (*CrossBridgeInitialized, error) {
	event := new(CrossBridgeInitialized)
	if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the CrossBridge contract.
type CrossBridgeManualReleasedIterator struct {
	Event *CrossBridgeManualReleased // Event containing the contract specifics and raw log

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
func (it *CrossBridgeManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeManualReleased)
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
		it.Event = new(CrossBridgeManualReleased)
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
func (it *CrossBridgeManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeManualReleased represents a ManualReleased event raised by the CrossBridge contract.
type CrossBridgeManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgeManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeManualReleasedIterator{contract: _CrossBridge.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *CrossBridgeManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeManualReleased)
				if err := _CrossBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseManualReleased(log types.Log) (*CrossBridgeManualReleased, error) {
	event := new(CrossBridgeManualReleased)
	if err := _CrossBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeMaxExtraDataLengthSetIterator is returned from FilterMaxExtraDataLengthSet and is used to iterate over the raw logs and unpacked data for MaxExtraDataLengthSet events raised by the CrossBridge contract.
type CrossBridgeMaxExtraDataLengthSetIterator struct {
	Event *CrossBridgeMaxExtraDataLengthSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeMaxExtraDataLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeMaxExtraDataLengthSet)
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
		it.Event = new(CrossBridgeMaxExtraDataLengthSet)
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
func (it *CrossBridgeMaxExtraDataLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeMaxExtraDataLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeMaxExtraDataLengthSet represents a MaxExtraDataLengthSet event raised by the CrossBridge contract.
type CrossBridgeMaxExtraDataLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxExtraDataLengthSet is a free log retrieval operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_CrossBridge *CrossBridgeFilterer) FilterMaxExtraDataLengthSet(opts *bind.FilterOpts) (*CrossBridgeMaxExtraDataLengthSetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeMaxExtraDataLengthSetIterator{contract: _CrossBridge.contract, event: "MaxExtraDataLengthSet", logs: logs, sub: sub}, nil
}

// WatchMaxExtraDataLengthSet is a free log subscription operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_CrossBridge *CrossBridgeFilterer) WatchMaxExtraDataLengthSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeMaxExtraDataLengthSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeMaxExtraDataLengthSet)
				if err := _CrossBridge.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseMaxExtraDataLengthSet(log types.Log) (*CrossBridgeMaxExtraDataLengthSet, error) {
	event := new(CrossBridgeMaxExtraDataLengthSet)
	if err := _CrossBridge.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossBridge contract.
type CrossBridgePausedIterator struct {
	Event *CrossBridgePaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePaused)
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
		it.Event = new(CrossBridgePaused)
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
func (it *CrossBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePaused represents a Paused event raised by the CrossBridge contract.
type CrossBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossBridgePausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgePausedIterator{contract: _CrossBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossBridgePaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePaused(log types.Log) (*CrossBridgePaused, error) {
	event := new(CrossBridgePaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the CrossBridge contract.
type CrossBridgePendingRemovedIterator struct {
	Event *CrossBridgePendingRemoved // Event containing the contract specifics and raw log

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
func (it *CrossBridgePendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePendingRemoved)
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
		it.Event = new(CrossBridgePendingRemoved)
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
func (it *CrossBridgePendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePendingRemoved represents a PendingRemoved event raised by the CrossBridge contract.
type CrossBridgePendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgePendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgePendingRemovedIterator{contract: _CrossBridge.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *CrossBridgePendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePendingRemoved)
				if err := _CrossBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePendingRemoved(log types.Log) (*CrossBridgePendingRemoved, error) {
	event := new(CrossBridgePendingRemoved)
	if err := _CrossBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePostCallGasReserveSetIterator is returned from FilterPostCallGasReserveSet and is used to iterate over the raw logs and unpacked data for PostCallGasReserveSet events raised by the CrossBridge contract.
type CrossBridgePostCallGasReserveSetIterator struct {
	Event *CrossBridgePostCallGasReserveSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgePostCallGasReserveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePostCallGasReserveSet)
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
		it.Event = new(CrossBridgePostCallGasReserveSet)
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
func (it *CrossBridgePostCallGasReserveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePostCallGasReserveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePostCallGasReserveSet represents a PostCallGasReserveSet event raised by the CrossBridge contract.
type CrossBridgePostCallGasReserveSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostCallGasReserveSet is a free log retrieval operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_CrossBridge *CrossBridgeFilterer) FilterPostCallGasReserveSet(opts *bind.FilterOpts) (*CrossBridgePostCallGasReserveSetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgePostCallGasReserveSetIterator{contract: _CrossBridge.contract, event: "PostCallGasReserveSet", logs: logs, sub: sub}, nil
}

// WatchPostCallGasReserveSet is a free log subscription operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_CrossBridge *CrossBridgeFilterer) WatchPostCallGasReserveSet(opts *bind.WatchOpts, sink chan<- *CrossBridgePostCallGasReserveSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePostCallGasReserveSet)
				if err := _CrossBridge.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePostCallGasReserveSet(log types.Log) (*CrossBridgePostCallGasReserveSet, error) {
	event := new(CrossBridgePostCallGasReserveSet)
	if err := _CrossBridge.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossBridge contract.
type CrossBridgeRoleAdminChangedIterator struct {
	Event *CrossBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleAdminChanged)
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
		it.Event = new(CrossBridgeRoleAdminChanged)
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
func (it *CrossBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the CrossBridge contract.
type CrossBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleAdminChangedIterator{contract: _CrossBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleAdminChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*CrossBridgeRoleAdminChanged, error) {
	event := new(CrossBridgeRoleAdminChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossBridge contract.
type CrossBridgeRoleGrantedIterator struct {
	Event *CrossBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleGranted)
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
		it.Event = new(CrossBridgeRoleGranted)
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
func (it *CrossBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleGranted represents a RoleGranted event raised by the CrossBridge contract.
type CrossBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleGrantedIterator{contract: _CrossBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleGranted)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleGranted(log types.Log) (*CrossBridgeRoleGranted, error) {
	event := new(CrossBridgeRoleGranted)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossBridge contract.
type CrossBridgeRoleRevokedIterator struct {
	Event *CrossBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleRevoked)
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
		it.Event = new(CrossBridgeRoleRevoked)
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
func (it *CrossBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleRevoked represents a RoleRevoked event raised by the CrossBridge contract.
type CrossBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleRevokedIterator{contract: _CrossBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleRevoked)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleRevoked(log types.Log) (*CrossBridgeRoleRevoked, error) {
	event := new(CrossBridgeRoleRevoked)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossBridge contract.
type CrossBridgeThresholdChangedIterator struct {
	Event *CrossBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeThresholdChanged)
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
		it.Event = new(CrossBridgeThresholdChanged)
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
func (it *CrossBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeThresholdChanged represents a ThresholdChanged event raised by the CrossBridge contract.
type CrossBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossBridgeThresholdChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeThresholdChangedIterator{contract: _CrossBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeThresholdChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseThresholdChanged(log types.Log) (*CrossBridgeThresholdChanged, error) {
	event := new(CrossBridgeThresholdChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairRegisteredIterator struct {
	Event *CrossBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairRegistered)
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
		it.Event = new(CrossBridgeTokenPairRegistered)
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
func (it *CrossBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*CrossBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairRegisteredIterator{contract: _CrossBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairRegistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*CrossBridgeTokenPairRegistered, error) {
	event := new(CrossBridgeTokenPairRegistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the CrossBridge contract.
type CrossBridgeTokenPauseSetIterator struct {
	Event *CrossBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPauseSet)
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
		it.Event = new(CrossBridgeTokenPauseSet)
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
func (it *CrossBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPauseSet represents a TokenPauseSet event raised by the CrossBridge contract.
type CrossBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPauseSetIterator{contract: _CrossBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPauseSet(log types.Log) (*CrossBridgeTokenPauseSet, error) {
	event := new(CrossBridgeTokenPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossBridge contract.
type CrossBridgeUnpausedIterator struct {
	Event *CrossBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUnpaused)
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
		it.Event = new(CrossBridgeUnpaused)
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
func (it *CrossBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUnpaused represents a Unpaused event raised by the CrossBridge contract.
type CrossBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossBridgeUnpausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUnpausedIterator{contract: _CrossBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUnpaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUnpaused(log types.Log) (*CrossBridgeUnpaused, error) {
	event := new(CrossBridgeUnpaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossBridge contract.
type CrossBridgeUpgradedIterator struct {
	Event *CrossBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUpgraded)
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
		it.Event = new(CrossBridgeUpgraded)
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
func (it *CrossBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUpgraded represents a Upgraded event raised by the CrossBridge contract.
type CrossBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUpgradedIterator{contract: _CrossBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUpgraded)
				if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUpgraded(log types.Log) (*CrossBridgeUpgraded, error) {
	event := new(CrossBridgeUpgraded)
	if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the CrossBridge contract.
type CrossBridgeVerificationDelayExpirationSetIterator struct {
	Event *CrossBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeVerificationDelayExpirationSet)
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
		it.Event = new(CrossBridgeVerificationDelayExpirationSet)
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
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the CrossBridge contract.
type CrossBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_CrossBridge *CrossBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*CrossBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeVerificationDelayExpirationSetIterator{contract: _CrossBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_CrossBridge *CrossBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeVerificationDelayExpirationSet)
				if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*CrossBridgeVerificationDelayExpirationSet, error) {
	event := new(CrossBridgeVerificationDelayExpirationSet)
	if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySetIterator struct {
	Event *CrossBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeVerificationDelaySet)
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
		it.Event = new(CrossBridgeVerificationDelaySet)
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
func (it *CrossBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*CrossBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeVerificationDelaySetIterator{contract: _CrossBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *CrossBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeVerificationDelaySet)
				if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*CrossBridgeVerificationDelaySet, error) {
	event := new(CrossBridgeVerificationDelaySet)
	if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
