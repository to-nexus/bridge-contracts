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

// CrossBridgeV2MetaData contains all meta data concerning the CrossBridgeV2 contract.
var CrossBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeTokenForwarded\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupplyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExtraDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"setCrossSupplyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMaxExtraDataLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"CrossSupplyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MaxExtraDataLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeExtraDataTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"computed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ctxValue\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeForwardAmountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardContextInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardNotExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardNotOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Disabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
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
	Bin: "0x60a080604052346100c257306080525f5160206156db5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161561490816100c782396080518181816124ae01526125d90152f35b6001600160401b0319166001600160401b039081175f5160206156db5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60a0806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a714613033575080630b1d8d0614612fc75780631313996b14612c275780631e24abdb14612c0a578063230f4cda14612bdb578063248a9ca314612bbd5780632f2ff15d14612b5357806336568abe14612b0f578063365d71e414612ab557806342cde4e814612a8957806348a00ed81461295e5780634bc51672146128f15780634be13f83146128ca5780634ee078ba146127695780634f1ef2861461258b578063502cc5c01461250257806352d1902d1461249c57806355fd07761461215a57806356ff54b01461210e5780635b605f5c14611f4e5780635c975abb14611f205780635fd262de14611e1a578063670e626814611b6f57806369ceb2f114611b52578063761fe2d814611b095780637921487414611aa5578063814914b5146119ef57806384b0196e146118f657806388d67d6d14610fd457806389232a0014610f995780638ae87c5c14610f3957806391cca3db14610f1157806391cf6d3e14610ef457806391d1485414610e9f578063a10bab0b14610e77578063a217fddf14610e5d578063a3246ad314610da2578063aa1bd0bc14610d56578063aa20ed4714610ce4578063ad3cb1cc14610c9d578063ae6893f814610c64578063b2b49e2e14610bb3578063b33eb36e14610ab3578063b7f3358d14610a26578063b8886e1a146109da578063b8aa837e1461094b578063bedb86fb14610862578063bfbf67651461078a578063cba25e4b14610724578063cd95970614610702578063cf56118e14610676578063d477f05f146105ff578063d547741f146105c7578063d5717fc51461057a578063e2af6cd714610507578063edbbea3914610330578063f0702e8e146103085763f698da251461029a575f80fd5b34610304575f3660031901126103045760206102b46151c7565b6102bc61521e565b60405190838201925f51602061559f5f395f51905f528452604083015260608201524660808201523060a082015260a081526102f960c082613180565b519020604051908152f35b5f80fd5b34610304575f366003190112610304576032546040516001600160a01b039091168152602090f35b346103045760803660031901126103045760043561034c6134af565b6103546130cc565b6064356001600160a01b038116939190849003610304576103736138a0565b81156104f8576001600160a01b03169182156104f85783156104f85761039882614de6565b6104a3575b815f5260056020526103b28360405f20614e51565b156104905760205f51602061555f5f395f51905f52916040516103d48161312f565b85815282810187815291151560408083018281525f606085018181526080860182815260a080880184815260c089018581528d865260068c528786208f87528c5296909420975188546001600160a01b0319166001600160a01b03918216178955985160018901805496516001600160a81b031990971691909a1617941515901b60ff60a01b169390931787555193956004949192916104779190151590613785565b51600285015551600384015551910155604051908152a4005b826311dd05f360e31b5f5260045260245ffd5b6104f36040516104b28161314a565b8381526003602082015f8152604083015f815260608401915f8352875f52600460205260405f2094518555516001850155516002840155511515910161374c565b61039d565b6302bfb33d60e51b5f5260045ffd5b34610304576020366003190112610304576004356001600160a01b03811690819003610304576105356137c2565b80156104f8575f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b34610304576020366003190112610304576004355f526004602052600260405f200154600181018091116105b357602090604051908152f35b634e487b7160e01b5f52601160045260245ffd5b34610304576040366003190112610304576105fd6004356105e66130b6565b906105f86105f382613526565b61397e565b613a3e565b005b34610304576020366003190112610304576106186130e2565b6106206137c2565b6001600160a01b0316801561066757603380546001600160a01b031916821790557f76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc508715f80a2005b638219abe360e01b5f5260045ffd5b34610304575f36600319011261030457604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106106ec576106e8856106d481870382613180565b6040519182916020835260208301906133bf565b0390f35b82548452602090930192600192830192016106bd565b34610304575f36600319011261030457602061071c6137a2565b604051908152f35b34610304576020366003190112610304576004356001600160a01b03811690819003610304576107526137c2565b603580546001600160a01b031916821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b34610304576080366003190112610304576004356107a66130b6565b9060443580151580820361030457606435908115158083036103045761085560409361083c7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c966107f561390f565b5f89815260056020528790206001600160a01b038b169a6108209161081b908d90614575565b61375d565b885f526006602052865f208a5f526020526001875f2001613785565b865f526009602052845f20885f52602052845f2061374c565b82519182526020820152a3005b346103045760203660031901126103045760043580151581036103045761088761390f565b156108e5576108946139c4565b600160ff195f51602061551f5f395f51905f525416175f51602061551f5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f51602061551f5f395f51905f525460ff81161561093c5760ff19165f51602061551f5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b34610304576040366003190112610304576004356109676134af565b61096f61390f565b815f52600360205260405f2054156109c75760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526109bc81600360405f200161374c565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b34610304576020366003190112610304577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435610a196137c2565b80606455604051908152a1005b346103045760203660031901126103045760043560ff811680910361030457610a4d6137c2565b8015610aa4576020817f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9260ff195f51602061541f5f395f51905f525416175f51602061541f5f395f51905f5255604051908152a1005b63f0f15b9160e01b5f5260045ffd5b3461030457610ac1366132c6565b905f60408051610ad081613114565b610ad861371c565b815282602082015201525f52600860205260405f20905f5260205260a060405f2060405190610b0682613114565b610b0f816135c4565b8252600760ff60068301541691610b2a6020850193846136b9565b015460408301908152610ba9610b9d60405195869560208752516060602088015280516080880152602081015182880152600180831b0360408201511660c0880152600180831b0360608201511660e08801526080810151610100880152015160c06101208701526101408601906133f2565b925160408501906134a2565b5160608301520390f35b3461030457610bc136613217565b908051825103610c55575f5b82518110156105fd57610be08183613544565b516001600160a01b03610bf38386613544565b5116610c016105f383613526565b610c0b8183614435565b610c1a575b5050600101610bcd565b815f525f5160206154bf5f395f51905f52602052610c3b8160405f20614e51565b610c10575b63d180cb3160e01b5f5260045260245260445ffd5b63031206d560e51b5f5260045ffd5b34610304576020366003190112610304576004355f526004602052600160405f200154600181018091116105b357602090604051908152f35b34610304575f366003190112610304576106e8604051610cbe604082613180565b60058152640352e302e360dc1b60208201526040519182916020835260208301906133f2565b3461030457610cf2366132c6565b90610cfb613831565b610d03613831565b610d0b6139eb565b805f526007602052610d2982610d248160405f20614575565b613558565b610d338282613abb565b5f5160206153df5f395f51905f525f80a35f5f51602061553f5f395f51905f525d005b34610304576020366003190112610304577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435610d956137c2565b80600155604051908152a1005b34610304576020366003190112610304576004355f525f5160206154bf5f395f51905f5260205260405f20604051806020835491828152019081935f5260205f20905f5b818110610e475750505081610dfc910382613180565b604051918291602083019060208452518091526040830191905f5b818110610e25575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610e17565b8254845260209093019260019283019201610de6565b34610304575f3660031901126103045760206040515f8152f35b34610304575f366003190112610304576035546040516001600160a01b039091168152602090f35b3461030457604036600319011261030457610eb86130b6565b6004355f525f5160206154ff5f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b34610304575f366003190112610304576020603454604051908152f35b34610304575f366003190112610304576033546040516001600160a01b039091168152602090f35b3461030457610f47366132c6565b90610f506137c2565b610f586139eb565b610f628282614309565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f51602061553f5f395f51905f525d005b3461030457606036600319011261030457610fb26130e2565b50610fbb6130b6565b50610fc46133af565b50633ac4266d60e11b5f5260045ffd5b6080366003190112610304576004356001600160401b03811161030457610fff903690600401613086565b60243591906001600160401b038311610304573660238401121561030457826004013561102b816131a3565b936110396040519586613180565b8185526024602086019260051b820101903682116103045760248101925b82841061186c5750506044359150506001600160401b03811161030457611082903690600401613416565b926064356001600160401b038111610304576110a2903690600401613416565b916110ab6139c4565b6110b36139eb565b8082511480611862575b80611858575b15611849573461183257909184925f9460be1981360301905b84871015611817578660051b81013582811215610304578101966111008185613544565b519761110c8289613544565b51906111188388613544565b5199813590815f52600460205261113a8260ff600360405f2001541615613572565b815f52600560205260405f2090611166604085019260018060a01b0361115f856134e0565b1690614575565b6001600160a01b03611177846134e0565b1690156118055750825f526004602052600260405f20019360018554018095556020810135948086036117ee57506001600160a01b036111b6846134e0565b1695606082019e8f6111c7906134e0565b9360808401359860a08501956111dd87876134f4565b36906111e8926132f7565b805190602001206040519160208301937fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191485528b60408501528c60608501526080840152600160a01b600190031660a08301528b60c083015260e082015260e0815261125661010082613180565b51902081519184518314806117e4575b156117d55760ff5f51602061541f5f395f51905f5254169361128b8486811015614db7565b5f955f935f5b8681106116bf5750505050505050909d9e98999a9b9c9d6112b491811015614db7565b5f6112c8876112c2866134e0565b87613baa565b93908093600a82101561158b578b9260018b931461167e575b50505050600a82101561158b57600182036113465750505060019561131c6113165f51602061543f5f395f51905f52936134e0565b916134e0565b906113366040519283928a8060a01b031697429184613a9d565b0390a45b019594939291906110dc565b909195845f52600760205261135e8660405f20614e51565b1561166b5781156116635750965b611375836134e0565b9161137f826134e0565b604051602098916113908a83613180565b5f8252604051956113a087613165565b8887528a87018a90526001600160a01b039081166040880152166060860152608085018b905260a08501521561165c576113dc600154426136c5565b905b604051936113eb85613114565b8452878401996113fb828c6136b9565b60408581019384525f88815260088b528181208a82528b52819020955180518755808b01516001880155908101516002870180546001600160a01b039283166001600160a01b03199182161790915560608301516003890180549190931691161790556080810151600487015560a00151805160058701916001600160401b03821161164857819061148d845461358c565b8d601f821161160e575b50508c90601f83116001146115aa575f9261159f575b50508160011b915f199060031b1c19161790555b600685019a5192600a84101561158b576115567f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f97608097600760019f9461155c9561157d9960ff8019835416911617905551910155895f5260068c5260405f208e8060a01b03611531836134e0565b168f8060a01b03165f528c52600260405f200161154f8582546136c5565b90556134e0565b946134e0565b98604051998c8060a01b03168a5289015242604089015260608801906134a2565b878060a01b031694a461133a565b634e487b7160e01b5f52602160045260245ffd5b015190505f806114ad565b5f8581528e81209350601f198516908f5b8282106115f75750509084600195949392106115df575b505050811b0190556114c1565b01515f1960f88460031b161c191690555f80806115d2565b60018596829396860151815501950193018f6115bb565b61163791865f52815f2090601f860160051c820192861061163e575b601f0160051c01906142f3565b5f8d611497565b909150819061162a565b634e487b7160e01b5f52604160045260245ffd5b5f906113de565b90509661136c565b856307a5c91d60e31b5f5260045260245ffd5b6116b695506116ae91929394506116a7906116a161169b8a6134e0565b966134e0565b936134f4565b36916132f7565b9288886145ea565b88875f806112e1565b61176461175b60426116cf6151c7565b6116d761521e565b6040519060208201925f51602061559f5f395f51905f528452604083015260608201524660808201523060a082015260a0815261171560c082613180565b5190206040519061190160f01b825260028201528560228201522060ff61173c858a613544565b51166117488588613544565b5190611754868a613544565b5192615306565b9092919261537e565b6001600160a01b038181169088161080611795575b611787575b50600101611291565b6001998a019990965061177e565b506001600160a01b0381165f9081527fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa5602052604090205460ff16611779565b63b3f07a3960e01b5f5260045ffd5b5083518314611266565b859063a6ab65c560e01b5f5260045260245260445ffd5b633142cb1160e11b5f5260045260245ffd5b5f5f51602061553f5f395f51905f525d602060405160018152f35b63943892eb60e01b5f525f6004523460245260445ffd5b6329f517fd60e01b5f5260045ffd5b50808351146110c3565b50808551146110bd565b83356001600160401b03811161030457820136604382011215610304576024810135611897816131a3565b916118a56040519384613180565b818352602060248185019360051b830101019036821161030457604401915b8183106118dc57505050815260209384019301611057565b823560ff81168103610304578152602092830192016118c4565b34610304575f366003190112610304575f51602061547f5f395f51905f525415806119d9575b1561199c5761196e61192c61416a565b6106e8611937614239565b61197c60405191611949602084613180565b5f83525f368137604051958695600f60f81b875260e0602088015260e08701906133f2565b9085820360408701526133f2565b904660608501523060808501525f60a085015283820360c08501526133bf565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f51602061557f5f395f51905f52541561191c565b3461030457604036600319011261030457611a086130b6565b611a106136e6565b506004355f52600660205260405f209060018060a01b03165f5260205260e060405f20600460405191611a428361312f565b60018060a01b03815416835260ff600182015460018060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c0820152611aa3604051809261335a565bf35b34610304576020366003190112610304576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b818110611af3576106e8856106d481870382613180565b8254845260209093019260019283019201611adc565b3461030457604036600319011261030457611b226130b6565b6004355f52600960205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b34610304575f366003190112610304576020606454604051908152f35b3461030457608036600319011261030457600435611b8b6130b6565b6044356001600160401b038111610304573660238201121561030457611bbb9036906024816004013591016132f7565b906064359060ff8216809203610304575f546001600160a01b0316908115611e0b57602091611c1e915f604051809681958294637c469ea160e11b84528b600485015260018060a01b0316998a60248501526080604485015260848401906133f2565b90606483015203925af1908115611e00575f91611dc6575b50611c3f6138a0565b82156104f8576001600160a01b03169081156104f85780156104f857611c6483614de6565b611d71575b825f526005602052611c7e8260405f20614e51565b15611d5e5781602093604051611c938161312f565b8281528581018481525f6040808401828152606085018381526080860184815260a080880186815260c089018781528a885260068f528688208c89528f5295909620975188546001600160a01b0319166001600160a01b03918216178955965160018901805495516001600160a81b03199096169190981617931515901b60ff60a01b1692909217855551600494929392611d3091151590613785565b516002850155516003840155519101555f51602061555f5f395f51905f52856040515f8152a4604051908152f35b506311dd05f360e31b5f5260045260245ffd5b611dc1604051611d808161314a565b8481526003602082015f8152604083015f815260608401915f8352885f52600460205260405f2094518555516001850155516002840155511515910161374c565b611c69565b90506020813d602011611df8575b81611de160209383613180565b8101031261030457611df2906136d2565b83611c36565b3d9150611dd4565b6040513d5f823e3d90fd5b6315aeca0d60e11b5f5260045ffd5b60e0366003190112610304576024356001600160a01b038116906004359082810361030457611e476130cc565b6064359360c4356001600160401b03811161030457611e6a90369060040161332d565b919092611e756139c4565b611e7f8287613bf7565b611e876139eb565b6001600160a01b031690811561066757600a548015908115611f15575b5015611f065761181796611ec4611efc9660a4359083608435918b613c73565b93909260405198611ed48a6130f8565b895260208901523360408901526060880152608087015260a086015260c085015236916132f7565b60e0820152613e19565b633953b33f60e11b5f5260045ffd5b905083111588611ea4565b34610304575f36600319011261030457602060ff5f51602061551f5f395f51905f5254166040519015158152f35b3461030457602036600319011261030457600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106120f5575050611f9d92500383613180565b815191611fa9836131a3565b92611fb76040519485613180565b808452611fc6601f19916131a3565b015f5b8181106120de5750505f5b815181101561208e575f838152600660205260409020600191906001600160a01b036120008386613544565b5116838060a01b03165f5260205260405f206004604051916120218361312f565b848060a01b03815416835260ff85820154868060a01b0381166020860152818160a01c161515604086015260a81c161515606084015260028101546080840152600381015460a0840152015460c082015261207c8287613544565b526120878186613544565b5001611fd4565b836040518091602082016020835281518091526020604084019201905f5b8181106120ba575050500390f35b91935091602060e0826120d0600194885161335a565b0194019101918493926120ac565b6020906120e96136e6565b82828801015201611fc9565b8454835260019485019487945060209093019201611f88565b34610304576020366003190112610304577f1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1602060043561214d6138a0565b80600a55604051908152a1005b60c0366003190112610304576004356121716130b6565b606435906084359060443560a4356001600160401b0381116103045761219b90369060040161332d565b60355491939092916001600160a01b03168015159081612492575b50156124845773__$66aabcaa73966043eae8d94048b21b06c8$__96873b156103045760405163a3f11c4f60e01b81525f816004818c5af48015611e0057612474575b50604051634cfbad0360e01b8152906080826004818c5af4968715611e00575f985f945f945f9a61241e575b5061222e6139c4565b6122388685613bf7565b6001600160a01b031694851561066757600a548015908115612413575b5015611f06576001600160a01b0316935f1985016123b4575b806122828461227d858b6136c5565b6136c5565b146122918461227d858b6136c5565b901561239f5750506122fb916122a991878686613c73565b60355460405191999298916116a7916001600160a01b03166122ca826130f8565b86825287602083015260408201528760608201528860808201528960a08201528a60c0820152611efc3684866132f7565b6020815191012096893b156103045760405198631587d80760e21b8a5260048a0152602489015260448801526064870152608486015260a485015260c484015260e48301526101048201525f8161012481855af48015611e005761238f575b50803b15610304575f60049160405192838092632f60d07160e01b82525af48015611e005761238557005b5f6105fd91613180565b5f61239991613180565b8161235a565b639df556d160e01b5f5260045260245260445ffd5b8a5f52600660205260405f20855f5260205260ff600160405f20015460a01c16806123ed575b61226e5763be75e8c760e01b5f5260045ffd5b50835f52600660205260405f20855f5260205260ff600160405f20015460a01c166123da565b90508811158d612255565b95509850985091506080833d60801161246c575b8161243f60809383613180565b8101031261030457825191612456602085016136d2565b986060604086015195015193999493988c612225565b3d9150612432565b5f61247e91613180565b886121f9565b62a81d6360e11b5f5260045ffd5b90503314886121b6565b34610304575f366003190112610304577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036124f35760206040515f51602061549f5f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b3461030457606036600319011261030457602435600435612521613831565b805f52600760205261253a82610d248160405f20614575565b7fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db79853659496020612569604435426136c5565b835f526008825260405f20855f52825280600760405f200155604051908152a3005b60403660031901126103045761259f6130e2565b6024356001600160401b0381116103045736602382011215610304576125cf9036906024816004013591016132f7565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115612747575b506124f3576126116137c2565b6040516352d1902d60e01b81526001600160a01b0383169290602081600481875afa5f9181612713575b506126535783634c9c8ce360e01b5f5260045260245ffd5b805f51602061549f5f395f51905f528592036127015750813b156126ef575f51602061549f5f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28151156126d7575f808360206105fd95519101845af46126d1614588565b91615250565b5050346126e057005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d60201161273f575b8161272f60209383613180565b810103126103045751908561263b565b3d9150612722565b5f51602061549f5f395f51905f52546001600160a01b03161415905083612604565b3461030457612777366132c6565b61277f6139c4565b6127876139eb565b815f5260046020526127a48260ff600360405f2001541615613572565b815f5260076020526127bd81610d248160405f20614575565b815f52600860205260405f20815f5260205260405f2061282b6007604051926127e584613114565b6127ee816135c4565b845261280460ff600683015416602086016136b9565b0154604080840191825292519283015160809093015190926001600160a01b031685613b4d565b50600a81101561158b5760016040519180602084015260208352612850604084613180565b036128a25750519182158015612899575b15612882576128709250613abb565b5f5f51602061553f5f395f51905f525d005b82637a88099160e11b5f526004524260245260445ffd5b50428310612861565b60405162461bcd60e51b8152602060048201529081906128c69060248301906133f2565b0390fd5b34610304575f366003190112610304575f546040516001600160a01b039091168152602090f35b346103045760203660031901126103045760043561290d6137c2565b61c350811061294f5760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91603654908060365582519182526020820152a1005b63050779f960e21b5f5260045ffd5b346103045760603660031901126103045760243560043561297d6130cc565b612985613831565b61298d6139eb565b815f5260076020526129a683610d248160405f20614575565b806129b18484614309565b916001600160a01b031615612a75575b81519060208301908151936129f3604082019460018060a01b038651169684608085019889519260a0870151946145ea565b50600a81101561158b5760018103612a5d5750905f51602061543f5f395f51905f52915192519360018060a01b039051169451612a37604051928392429184613a9d565b0390a45f5160206153df5f395f51905f525f80a35f5f51602061553f5f395f51905f525d005b63290cd55f60e01b5f52612a7090613494565b60245ffd5b5060608101516001600160a01b03166129c1565b34610304575f36600319011261030457602060ff5f51602061541f5f395f51905f525416604051908152f35b3461030457612ac336613217565b908051825103610c55575f5b82518110156105fd5780612b08612ae860019385613544565b51838060a01b03612af98488613544565b5116906105f86105f382613526565b5001612acf565b3461030457604036600319011261030457612b286130b6565b336001600160a01b03821603612b44576105fd90600435613a3e565b63334bd91960e11b5f5260045ffd5b3461030457604036600319011261030457600435612b6f6130b6565b612b7b6105f383613526565b612b858183614435565b612b8b57005b815f525f5160206154bf5f395f51905f52602052612bb660405f209160018060a01b03168092614e51565b15610c4057005b3461030457602036600319011261030457602061071c600435613526565b34610304575f3660031901126103045760365480612c0257506020620249f0604051908152f35b60209061071c565b34610304575f366003190112610304576020600a54604051908152f35b6040366003190112610304576004356001600160401b03811161030457612c52903690600401613086565b6024356001600160401b03811161030457366023820112156103045760048101356001600160401b0381116103045736602460e083028401011161030457612c986139c4565b612ca06139eb565b335f9081527f7bf51035c163aa501aba0b1731707347fc5168978f3fc3c84ea475ee63d4b83f602052604090205460ff1615612f9057808303611849575f5b838110612cf8575f5f51602061553f5f395f51905f525d005b612d038185876134be565b3590612d1b6020612d1583888a6134be565b016134e0565b612d2b6060612d1584898b6134be565b926080612d3984898b6134be565b01359360a0612d49858a8c6134be565b01359260c0612d59868b8d6134be565b01359087861015612f7c578a968a93612d8a612d80898d60e082020197602489019c6134be565b60e08101906134f4565b6001600160a01b0385169a9195909491612da48c8a613bf7565b6001600160a01b03612db5826134e0565b168c14906001600160a01b0390612dcb906134e0565b169015612f66575060448701986001600160a01b03612de98b6134e0565b6001600160a01b0390951694168403612f5757600a548015908115612f4c575b5015611f065784612e1a938a613c73565b929091606487013596612e4c88612e358761227d88886136c5565b811015612e468861227d89896136c5565b90613a20565b6032546001600160a01b031690612e628b6134e0565b9160a48201359160ff831680930361030457813b15610304578e926040519b8c958695637f40ec1760e11b87526004870152600160a01b6001900316602486015230604486015260648501526084820135608485015260a484015260c481013560c484015260e4013560e48301525a925f610104928195f19a8b15611e0057612efa612f369a611efc9960019e612f3c575b506134e0565b9060405199612f088b6130f8565b8a5260208a01528b8060a01b031660408901526060880152608087015260a086015260c085015236916132f7565b01612cdf565b5f612f4691613180565b5f612ef4565b90508611155f612e09565b630672aec160e01b5f5260045ffd5b8b6313f7c32b60e31b5f5260045260245260445ffd5b634e487b7160e01b5f52603260045260245ffd5b63e2517d3f60e01b5f52336004527f6b8b15f1c11543d8280deaa7c24d12fffba6a357e4428e8c43e4234790186bff60245260445ffd5b34610304576020366003190112610304576004356001600160a01b0381169081900361030457612ff56137c2565b801561066757603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b34610304576020366003190112610304576004359063ffffffff60e01b821680920361030457602091637965db0b60e01b8114908115613075575b5015158152f35b6301ffc9a760e01b1490508361306e565b9181601f84011215610304578235916001600160401b038311610304576020808501948460051b01011161030457565b602435906001600160a01b038216820361030457565b604435906001600160a01b038216820361030457565b600435906001600160a01b038216820361030457565b61010081019081106001600160401b0382111761164857604052565b606081019081106001600160401b0382111761164857604052565b60e081019081106001600160401b0382111761164857604052565b608081019081106001600160401b0382111761164857604052565b60c081019081106001600160401b0382111761164857604052565b601f909101601f19168101906001600160401b0382119082101761164857604052565b6001600160401b0381116116485760051b60200190565b9080601f830112156103045781356131d1816131a3565b926131df6040519485613180565b81845260208085019260051b82010192831161030457602001905b8282106132075750505090565b81358152602091820191016131fa565b906040600319830112610304576004356001600160401b0381116103045782613242916004016131ba565b91602435906001600160401b03821161030457806023830112156103045781600401359061326f826131a3565b9261327d6040519485613180565b8284526024602085019360051b82010191821161030457602401915b8183106132a65750505090565b82356001600160a01b038116810361030457815260209283019201613299565b6040906003190112610304576004359060243590565b6001600160401b03811161164857601f01601f191660200190565b929192613303826132dc565b916133116040519384613180565b829481845281830111610304578281602093845f960137010152565b9181601f84011215610304578235916001600160401b038311610304576020838186019501011161030457565b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b6044359060ff8216820361030457565b90602080835192838152019201905f5b8181106133dc5750505090565b82518452602093840193909201916001016133cf565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9080601f8301121561030457813561342d816131a3565b9261343b6040519485613180565b81845260208085019260051b820101918383116103045760208201905b83821061346757505050505090565b81356001600160401b03811161030457602091613489878480948801016131ba565b815201910190613458565b600a81101561158b57600452565b90600a82101561158b5752565b60243590811515820361030457565b9190811015612f7c5760051b8101359060fe1981360301821215610304570190565b356001600160a01b03811681036103045790565b903590601e198136030182121561030457018035906001600160401b0382116103045760200191813603831361030457565b5f525f5160206154ff5f395f51905f52602052600160405f20015490565b8051821015612f7c5760209160051b010190565b156135605750565b6373a1390160e11b5f5260045260245ffd5b1561357a5750565b636fc24b7960e11b5f5260045260245ffd5b90600182811c921680156135ba575b60208310146135a657565b634e487b7160e01b5f52602260045260245ffd5b91607f169161359b565b906040516135d181613165565b60058193805483526001810154602084015260018060a01b03600282015416604084015260018060a01b03600382015416606084015260048101546080840152019060405180925f908054906136268261358c565b80855291600181169081156136925750600114613653575b505060a0929161364f910384613180565b0152565b5f908152602081209092505b81831061367657505081016020018161364f61363e565b602091935080600191548385890101520191019091849261365f565b60ff191660208681019190915292151560051b8501909201925083915061364f905061363e565b600a82101561158b5752565b919082018092116105b357565b51906001600160a01b038216820361030457565b604051906136f38261312f565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b6040519061372982613165565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b9060ff801983541691151516179055565b156137655750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6065545f52600660205260405f2060015f52602052600460405f20015490565b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff16156137fa57565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b335f9081527fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143602052604090205460ff161561386957565b63e2517d3f60e01b5f52336004527f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0960245260445ffd5b335f9081527f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f602052604090205460ff16156138d857565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b335f9081527f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456602052604090205460ff161561394757565b63e2517d3f60e01b5f52336004527f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92960245260445ffd5b5f8181525f5160206154ff5f395f51905f526020908152604080832033845290915290205460ff16156139ae5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f51602061551f5f395f51905f5254166139dc57565b63d93c066560e01b5f5260045ffd5b5f51602061553f5f395f51905f525c613a115760015f51602061553f5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b15613a29575050565b63943892eb60e01b5f5260045260245260445ffd5b919091613a4b83826144d9565b9283613a55575050565b815f525f5160206154bf5f395f51905f52602052613a8060405f209160018060a01b03168092614e98565b15613a89575050565b62a95f1b60e31b5f5260045260245260445ffd5b604091949392606082019560018060a01b0316825260208201520152565b90613ac591614309565b60018060a01b036060820151168151906020830190815193613b04604082019460018060a01b038651169684608085019889519260a0870151946145ea565b50600a81101561158b5760018103612a5d5750905f51602061543f5f395f51905f52915192519360018060a01b039051169451613b48604051928392429184613a9d565b0390a4565b9091906001600160a01b03831660011480613b9f575b613b79575b90613b7592600192614cd0565b9091565b90613b868161227d6137a2565b60645410613b945790613b68565b505050600990600190565b506065548114613b63565b9091906001600160a01b03831660011480613bec575b613bd1575b90613b75925f92614cd0565b90613bde8161227d6137a2565b60645410613b945790613bc5565b506065548114613bc0565b5f8181526005602052604090206001600160a01b03831692613c1e9161081b908590614575565b805f526004602052613c3b8160ff600360405f2001541615613572565b5f52600660205260405f20815f5260205260ff600160405f20015460a81c16613c615750565b6338384f6f60e11b5f5260045260245ffd5b5f8181526006602090815260408083206001600160a01b0390951680845294909152908190209051929694959260c09190613cad8161312f565b81546001600160a01b0390811682526001830154908116602083015260a081811c60ff908116158015604086015260a89390931c1615156060840152600284015460808401526003840154908301526004909201549201829052613e05575b506032546001600160a01b0316958615613df657606460609260405198899384926337dba1f760e21b8452600484015260248301528660448301525afa908115611e00575f955f905f93613db3575b50829581978515159586613da8575b505084613d9d575b505082613d92575b505015613d8357565b6358e8878560e01b5f5260045ffd5b101590505f80613d7a565b101592505f80613d72565b101594505f80613d6a565b96505090506060853d606011613dee575b81613dd160609383613180565b81010312610304578451906040602087015196015191955f613d5b565b3d9150613dc4565b6330d45fb160e01b5f5260045ffd5b808480613e13931015613a20565b5f613d0c565b80515f526004602052600160405f2001600181540180915581515f52600660205260405f20602083019060018060a01b0382511660018060a01b03165f5260205260405f206001808060a01b03910154169280519360018060a01b0383511694604083019060018060a01b03825116966080850191825160a08701998a5193613ea860c08a01958651906136c5565b93805f52600660205260405f2060018060a01b0383165f5260205260405f2060ff600182015460a01c165f1461415557600301613ee68582546136c5565b90555b60018203613fe25750505080613f0283613f0f936136c5565b3414612e468434936136c5565b80613f9c575b50905f5160206155bf5f395f51905f5294613f91925b81519760018060a01b039051169460018060a01b039051169960e060018060a01b03606085015116945191519251930151936040519788978852602088015260408701526060860152608085015260a084015261010060c08401526101008301906133f2565b4260e08301520390a4565b603354604051929392613fc79290916001600160a01b0316613fbf602084613180565b5f835261516d565b15613fd357905f613f15565b63559d141b60e11b5f5260045ffd5b9091929398959697983461183257614034613ffd82876136c5565b604051906323b872dd60e01b602083015286602483015230604483015260648201526064815261402e608482613180565b846152ae565b80614105575b505f52600660205260405f2060018060a01b0382165f5260205260ff600160405f20015460a01c1615614086575b50505091613f91915f5160206155bf5f395f51905f52959493613f2b565b604051632770a7eb60e21b8152602081806140a58730600484016145cf565b03815f865af1908115611e00575f916140d6575b5061406857639ac2f56d60e01b5f5260045260245260445260645ffd5b6140f8915060203d6020116140fe575b6140f08183613180565b8101906145b7565b5f6140b9565b503d6140e6565b60335460405163a9059cbb60e01b602082015261414f92909161414991839161413b91906001600160a01b0316602484016145cf565b03601f198101835282613180565b836152ae565b5f61403a565b6004016141638582546142e6565b9055613ee9565b604051905f825f5160206153ff5f395f51905f5254916141898361358c565b808352926001811690811561421a57506001146141af575b6141ad92500383613180565b565b505f5160206153ff5f395f51905f525f90815290917f42ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5b8183106141fe5750509060206141ad928201016141a1565b60209193508060019154838589010152019101909184926141e6565b602092506141ad94915060ff191682840152151560051b8201016141a1565b604051905f825f51602061545f5f395f51905f5254916142588361358c565b808352926001811690811561421a575060011461427b576141ad92500383613180565b505f51602061545f5f395f51905f525f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b8183106142ca5750509060206141ad928201016141a1565b60209193508060019154838589010152019101909184926142b2565b919082039182116105b357565b8181106142fe575050565b5f81556001016142f3565b919061431361371c565b50825f5260076020526143298160405f20614e98565b1561442357825f52600860205260405f20815f5260205261434c60405f206135c4565b5f848152600660209081526040808320818501516001600160a01b031684529091529020608082015160029091018054929592909161438a916142e6565b90555f52600860205260405f20905f526020525f600760408220828155826001820155826002820155826003820155826004820155600581016143cd815461358c565b90816143e0575b50508260068201550155565b81601f8693116001146143f75750555b5f806143d4565b8183526020832061441391601f0160051c8101906001016142f3565b80825281602081209155556143f0565b637f11bea960e01b5f5260045260245ffd5b5f8181525f5160206154ff5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff166144d3575f8181525f5160206154ff5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f8181525f5160206154ff5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16156144d3575f8181525f5160206154ff5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001915f520160205260405f2054151590565b3d156145b2573d90614599826132dc565b916145a76040519384613180565b82523d5f602084013e565b606090565b90816020910312610304575180151581036103045790565b6001600160a01b039091168152602081019190915260400190565b94919290936080525f935f90865f52600660205260405f2060018060a01b0360018060a01b038716165f5260205260ff600160405f20015460a01c16935f9160018060a01b03603554169160188151101580614cc7575b80614cbd575b6146bc575b5050506146b3575060805161466292918461505d565b92600a84101561469f576001841461467e575050509060805190565b6080519294935061469992916001600160a01b03169061500d565b60019190565b634e487b7160e01b83526021600452602483fd5b95939450505050565b602081015160601c908a6034820151918a604051966146da8861314a565b5f88525f60208901525f6040890152606080890152614707858760805160018060a01b0386168b89614f5d565b151588525f5f80604051602081019062483e5560e91b82528a602482015260248152614734604482613180565b5190855afa614741614588565b9080614cb1575b614ca5575b50614800575b5050505050508151614766575b8061464c565b602082015115156060604084015115159301519173__$66aabcaa73966043eae8d94048b21b06c8$__90813b15610304578b6147d65f9560405197889687958695631b24662760e11b8752600487015260248601526044850152606484015260a0608484015260a48301906133f2565b03915af48015611e00576147eb575b80614760565b6147f89196505f90613180565b5f945f6147e5565b6001600160a01b03821660011460408901819052928315614b57575b8b60408a015161482e575b5050614753565b89865f93859c9e9a9c85948d9f9b9d6148c69060365480158914614b525750620249f05b805a108914614b3e57506001602089985b01528b158814614b325761413b88955b6040519485936020850199631eeed51360e01b8b52602486015260448501528d60018060a01b0316606485015260018060a01b0316608484015260805160a484015260c060c484015260e48301906133f2565b5193f16148d1614588565b60608c0152151560408b01528215614ae3575b50506040880151151580614ad3575b156149ec57505050506060840151908151820191604081602085019403126103045760208101516040820151926001600160401b0384116103045784603f858501011215610304576020848401015161494b816132dc565b956149596040519788613180565b8187526040858701830101116103045787955f8f96836149bf9460406020935f5160206154df5f395f51905f529a01018386015e8301015260405193849363ffffffff60e01b1684526001602085015260408401526080606084015260808301906133f2565b0390a46080516149d9906001600160a01b0387168961500d565b60019182915b5f808a818a81808b614827565b9384879598969498979397159015614acb575b614a52575b5f5160206154df5f395f51905f5290606087015190614a4a60405192839263ffffffff60e01b1683525f60208401525f60408401526080606084015260808301906133f2565b0390a46149df565b915050604051632770a7eb60e21b81526020818b815f81614a7960805130600484016145cf565b03926001600160a01b03165af18015611e005784928d925f5160206154df5f395f51905f5292614aac575b509050614a04565b614ac49060203d6020116140fe576140f08183613180565b505f614aa4565b5089156149ff565b50604060608901515110156148f3565b5f91829182604051602081019263095ea7b360e01b8452602482015281604482015260448152614b14606482613180565b51926001600160a01b03165af150614b2a614588565b508a5f6148e4565b61413b60805195614873565b6020614b4c6001925a6142e6565b98614863565b614852565b8c1580614c25575b5f80604051602081019063095ea7b360e01b8252614b878161413b6080518a602484016145cf565b5190826001600160a01b0389165af1614b9e614588565b501590811560408c015281614c1d575b501561481c5791935060206040518095632770a7eb60e21b8252815f81614bdb60805130600484016145cf565b03926001600160a01b03165af1918215611e00578f948e93614bfe575b5061481c565b614c169060203d6020116140fe576140f08183613180565b505f614bf8565b90505f614bae565b92909450602060405180926340c10f1960e01b8252815f81614c4d60805130600484016145cf565b03926001600160a01b03165af1908115611e00575f91614c86575b5015614c77578e938d92614b5f565b632b60b36f60e21b5f5260045ffd5b614c9f915060203d6020116140fe576140f08183613180565b5f614c68565b6020915001515f61474d565b50602081511015614748565b50823b1515614647565b50821515614641565b6032546001600160a01b0316945f94909392918615613df6575f9081526009602090815260408083206001600160a01b039095168084529490915290205460ff16614dab57614da3579060446020925f6040519788948593633f4bdec560e01b8552600485015260248401525af1928315611e00575f93614d66575b5082600a81101561158b57600103614d6057565b60019150565b9092506020813d602011614d9b575b81614d8260209383613180565b810103126103045751600a81101561030457915f614d4c565b3d9150614d75565b506001935050565b50600294505f93505050565b15614dbf5750565b631aebd17960e11b5f5260045260245ffd5b8054821015612f7c575f5260205f2001905f90565b805f52600360205260405f2054155f14614e4c57600254600160401b81101561164857614e35614e1f8260018594016002556002614dd1565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b614e5b8282614575565b6144d357805490600160401b8210156116485782614e83614e1f846001809601855584614dd1565b90558054925f520160205260405f2055600190565b906001820191815f528260205260405f20548015155f14614f55575f1981018181116105b35782545f198101919082116105b357818103614f20575b50505080548015614f0c575f190190614eed8282614dd1565b8154905f199060031b1b19169055555f526020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b614f40614f30614e1f9386614dd1565b90549060031b1c92839286614dd1565b90555f528360205260405f20555f8080614ed4565b505050505f90565b92949293919290916001600160a01b0316301480159190614ff1575b50614f555773__$66aabcaa73966043eae8d94048b21b06c8$__92833b15610304575f936084926040519687958694631ff65d3360e31b8652600486015260018060a01b03166024850152604484015260648301525af48015611e0057614fe1575b50600190565b5f614feb91613180565b5f614fdb565b6001600160e01b031916632afe83bb60e11b141590505f614f79565b5f52600660205260405f209060018060a01b03165f5260205260405f2060ff600182015460a01c165f1461504e57600361504a91019182546142e6565b9055565b600461504a91019182546136c5565b6001600160a01b0316929190600184146150ef5781156150e6576150ac92156150ba5760405163a9059cbb60e01b6020820152916150a491839161413b91602484016145cf565b600592615124565b156150b75750600190565b90565b6040516340c10f1960e01b6020820152916150de91839161413b91602484016145cf565b600692615124565b50505050600190565b60405161511594509250615104602084613180565b5f83526001600160a01b031661516d565b1561511f57600190565b600590565b81516001600160a01b03909116915f9182916020018285620186a0f1615148614588565b90156144d357805161516457503b1561516057600190565b5f90565b60209150015190565b814710614c775782516001600160a01b03909116925f9182916020018486620186a0f190615199614588565b91156151c057156151ac575b5050600190565b805161516457503b15615160575f806151a5565b5050505f90565b6151cf61416a565b80519081156151df576020012090565b50505f51602061547f5f395f51905f525480156151f95790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615226614239565b8051908115615236576020012090565b50505f51602061557f5f395f51905f525480156151f95790565b90615274575080511561526557805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806152a5575b615285575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561527d565b905f602091828151910182855af115611e00575f513d6152fd57506001600160a01b0381163b155b6152dd5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156152d6565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615373579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15611e00575f516001600160a01b0381161561536957905f905f90565b505f906001905f90565b5050505f9160039190565b600481101561158b5780615390575050565b600181036153a75763f645eedf60e01b5f5260045ffd5b600281036153c2575063fce698f760e01b5f5260045260245ffd5b6003146153cc5750565b6335e2f38360e21b5f5260045260245ffdfeb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00c4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e4002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5dba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a2646970667358221220cdbdc532fb2e7e7aadad7881bbb32120035ed4ce4328b2d0d9936b2cdbcbaaa564736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// CrossBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeV2MetaData.ABI instead.
var CrossBridgeV2ABI = CrossBridgeV2MetaData.ABI

// Deprecated: Use CrossBridgeV2MetaData.Sigs instead.
// CrossBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var CrossBridgeV2FuncSigs = CrossBridgeV2MetaData.Sigs

// CrossBridgeV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossBridgeV2MetaData.Bin instead.
var CrossBridgeV2Bin = CrossBridgeV2MetaData.Bin

// DeployCrossBridgeV2 deploys a new Ethereum contract, binding an instance of CrossBridgeV2 to it.
func DeployCrossBridgeV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossBridgeV2, error) {
	parsed, err := CrossBridgeV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossBridgeV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossBridgeV2{CrossBridgeV2Caller: CrossBridgeV2Caller{contract: contract}, CrossBridgeV2Transactor: CrossBridgeV2Transactor{contract: contract}, CrossBridgeV2Filterer: CrossBridgeV2Filterer{contract: contract}}, nil
}

// CrossBridgeV2 is an auto generated Go binding around an Ethereum contract.
type CrossBridgeV2 struct {
	CrossBridgeV2Caller     // Read-only binding to the contract
	CrossBridgeV2Transactor // Write-only binding to the contract
	CrossBridgeV2Filterer   // Log filterer for contract events
}

// CrossBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossBridgeV2Session struct {
	Contract     *CrossBridgeV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossBridgeV2CallerSession struct {
	Contract *CrossBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CrossBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossBridgeV2TransactorSession struct {
	Contract     *CrossBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrossBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossBridgeV2Raw struct {
	Contract *CrossBridgeV2 // Generic contract binding to access the raw methods on
}

// CrossBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossBridgeV2CallerRaw struct {
	Contract *CrossBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// CrossBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossBridgeV2TransactorRaw struct {
	Contract *CrossBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossBridgeV2 creates a new instance of CrossBridgeV2, bound to a specific deployed contract.
func NewCrossBridgeV2(address common.Address, backend bind.ContractBackend) (*CrossBridgeV2, error) {
	contract, err := bindCrossBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2{CrossBridgeV2Caller: CrossBridgeV2Caller{contract: contract}, CrossBridgeV2Transactor: CrossBridgeV2Transactor{contract: contract}, CrossBridgeV2Filterer: CrossBridgeV2Filterer{contract: contract}}, nil
}

// NewCrossBridgeV2Caller creates a new read-only instance of CrossBridgeV2, bound to a specific deployed contract.
func NewCrossBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*CrossBridgeV2Caller, error) {
	contract, err := bindCrossBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2Caller{contract: contract}, nil
}

// NewCrossBridgeV2Transactor creates a new write-only instance of CrossBridgeV2, bound to a specific deployed contract.
func NewCrossBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossBridgeV2Transactor, error) {
	contract, err := bindCrossBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2Transactor{contract: contract}, nil
}

// NewCrossBridgeV2Filterer creates a new log filterer instance of CrossBridgeV2, bound to a specific deployed contract.
func NewCrossBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossBridgeV2Filterer, error) {
	contract, err := bindCrossBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2Filterer{contract: contract}, nil
}

// bindCrossBridgeV2 binds a generic wrapper to an already deployed contract.
func bindCrossBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridgeV2 *CrossBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridgeV2.Contract.CrossBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridgeV2 *CrossBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.CrossBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridgeV2 *CrossBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.CrossBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridgeV2 *CrossBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridgeV2 *CrossBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridgeV2 *CrossBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridgeV2.Contract.DEFAULTADMINROLE(&_CrossBridgeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridgeV2.Contract.DEFAULTADMINROLE(&_CrossBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridgeV2 *CrossBridgeV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridgeV2 *CrossBridgeV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_CrossBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_CrossBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2Caller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2Session) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridgeV2.Contract.AllChainIDs(&_CrossBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridgeV2.Contract.AllChainIDs(&_CrossBridgeV2.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2Caller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2Session) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridgeV2.Contract.AllPendingIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridgeV2.Contract.AllPendingIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridgeV2 *CrossBridgeV2Caller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridgeV2 *CrossBridgeV2Session) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridgeV2.Contract.AllTokenPairs(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridgeV2.Contract.AllTokenPairs(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Caller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Session) BridgeExecutor() (common.Address, error) {
	return _CrossBridgeV2.Contract.BridgeExecutor(&_CrossBridgeV2.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) BridgeExecutor() (common.Address, error) {
	return _CrossBridgeV2.Contract.BridgeExecutor(&_CrossBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Caller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Session) BridgeVerifier() (common.Address, error) {
	return _CrossBridgeV2.Contract.BridgeVerifier(&_CrossBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) BridgeVerifier() (common.Address, error) {
	return _CrossBridgeV2.Contract.BridgeVerifier(&_CrossBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Caller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Session) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridgeV2.Contract.CrossMintableERC20Code(&_CrossBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridgeV2.Contract.CrossMintableERC20Code(&_CrossBridgeV2.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) CrossSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "crossSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) CrossSupply() (*big.Int, error) {
	return _CrossBridgeV2.Contract.CrossSupply(&_CrossBridgeV2.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) CrossSupply() (*big.Int, error) {
	return _CrossBridgeV2.Contract.CrossSupply(&_CrossBridgeV2.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) CrossSupplyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "crossSupplyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridgeV2.Contract.CrossSupplyLimit(&_CrossBridgeV2.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridgeV2.Contract.CrossSupplyLimit(&_CrossBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Caller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2Session) Dev() (common.Address, error) {
	return _CrossBridgeV2.Contract.Dev(&_CrossBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) Dev() (common.Address, error) {
	return _CrossBridgeV2.Contract.Dev(&_CrossBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Session) DomainSeparator() ([32]byte, error) {
	return _CrossBridgeV2.Contract.DomainSeparator(&_CrossBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridgeV2.Contract.DomainSeparator(&_CrossBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridgeV2 *CrossBridgeV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossBridgeV2 *CrossBridgeV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridgeV2.Contract.Eip712Domain(&_CrossBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridgeV2.Contract.Eip712Domain(&_CrossBridgeV2.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridgeV2.Contract.GetNextFinalizeIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridgeV2.Contract.GetNextFinalizeIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridgeV2.Contract.GetNextInitiateIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridgeV2.Contract.GetNextInitiateIndex(&_CrossBridgeV2.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridgeV2 *CrossBridgeV2Session) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridgeV2.Contract.GetPendingArguments(&_CrossBridgeV2.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridgeV2.Contract.GetPendingArguments(&_CrossBridgeV2.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridgeV2.Contract.GetRoleAdmin(&_CrossBridgeV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridgeV2.Contract.GetRoleAdmin(&_CrossBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridgeV2 *CrossBridgeV2Session) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridgeV2.Contract.GetRoleMembers(&_CrossBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridgeV2.Contract.GetRoleMembers(&_CrossBridgeV2.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridgeV2 *CrossBridgeV2Caller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridgeV2 *CrossBridgeV2Session) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridgeV2.Contract.GetTokenPair(&_CrossBridgeV2.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridgeV2.Contract.GetTokenPair(&_CrossBridgeV2.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridgeV2.Contract.HasRole(&_CrossBridgeV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridgeV2.Contract.HasRole(&_CrossBridgeV2.CallOpts, role, account)
}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridgeV2 *CrossBridgeV2Caller) Initialize(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "initialize", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) Initialize(arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	return _CrossBridgeV2.Contract.Initialize(&_CrossBridgeV2.CallOpts, arg0, arg1, arg2)
}

// Initialize is a free data retrieval call binding the contract method 0x89232a00.
//
// Solidity: function initialize(address , address , uint8 ) pure returns()
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) Initialize(arg0 common.Address, arg1 common.Address, arg2 uint8) error {
	return _CrossBridgeV2.Contract.Initialize(&_CrossBridgeV2.CallOpts, arg0, arg1, arg2)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) InitializedAt() (*big.Int, error) {
	return _CrossBridgeV2.Contract.InitializedAt(&_CrossBridgeV2.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) InitializedAt() (*big.Int, error) {
	return _CrossBridgeV2.Contract.InitializedAt(&_CrossBridgeV2.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Caller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridgeV2.Contract.IsTokenFinalizePaused(&_CrossBridgeV2.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridgeV2.Contract.IsTokenFinalizePaused(&_CrossBridgeV2.CallOpts, remoteChainID, token)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) MaxExtraDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "maxExtraDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) MaxExtraDataLength() (*big.Int, error) {
	return _CrossBridgeV2.Contract.MaxExtraDataLength(&_CrossBridgeV2.CallOpts)
}

// MaxExtraDataLength is a free data retrieval call binding the contract method 0x1e24abdb.
//
// Solidity: function maxExtraDataLength() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) MaxExtraDataLength() (*big.Int, error) {
	return _CrossBridgeV2.Contract.MaxExtraDataLength(&_CrossBridgeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) Paused() (bool, error) {
	return _CrossBridgeV2.Contract.Paused(&_CrossBridgeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) Paused() (bool, error) {
	return _CrossBridgeV2.Contract.Paused(&_CrossBridgeV2.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Caller) PostCallGasReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "postCallGasReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2Session) PostCallGasReserve() (*big.Int, error) {
	return _CrossBridgeV2.Contract.PostCallGasReserve(&_CrossBridgeV2.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) PostCallGasReserve() (*big.Int, error) {
	return _CrossBridgeV2.Contract.PostCallGasReserve(&_CrossBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2Session) ProxiableUUID() ([32]byte, error) {
	return _CrossBridgeV2.Contract.ProxiableUUID(&_CrossBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridgeV2.Contract.ProxiableUUID(&_CrossBridgeV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridgeV2.Contract.SupportsInterface(&_CrossBridgeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridgeV2.Contract.SupportsInterface(&_CrossBridgeV2.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridgeV2 *CrossBridgeV2Caller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossBridgeV2.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridgeV2 *CrossBridgeV2Session) Threshold() (uint8, error) {
	return _CrossBridgeV2.Contract.Threshold(&_CrossBridgeV2.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridgeV2 *CrossBridgeV2CallerSession) Threshold() (uint8, error) {
	return _CrossBridgeV2.Contract.Threshold(&_CrossBridgeV2.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Transactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.BridgeToken(&_CrossBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.BridgeToken(&_CrossBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) BridgeTokenForwarded(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "bridgeTokenForwarded", toChainID, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) BridgeTokenForwarded(toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.BridgeTokenForwarded(&_CrossBridgeV2.TransactOpts, toChainID, to, value, networkFee, exFee, extraData)
}

// BridgeTokenForwarded is a paid mutator transaction binding the contract method 0x55fd0776.
//
// Solidity: function bridgeTokenForwarded(uint256 toChainID, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) BridgeTokenForwarded(toChainID *big.Int, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.BridgeTokenForwarded(&_CrossBridgeV2.TransactOpts, toChainID, to, value, networkFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ChangeThreshold(&_CrossBridgeV2.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ChangeThreshold(&_CrossBridgeV2.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridgeV2 *CrossBridgeV2Transactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridgeV2 *CrossBridgeV2Session) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.CreateToken(&_CrossBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.CreateToken(&_CrossBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Transactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2Session) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.FinalizeBridgeBatch(&_CrossBridgeV2.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.FinalizeBridgeBatch(&_CrossBridgeV2.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.GrantRole(&_CrossBridgeV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.GrantRole(&_CrossBridgeV2.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.GrantRoleBatch(&_CrossBridgeV2.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.GrantRoleBatch(&_CrossBridgeV2.TransactOpts, roles, accounts)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ManualReleasePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ManualReleasePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ManualReleasePendingWithRecipient(&_CrossBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ManualReleasePendingWithRecipient(&_CrossBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.PermitBridgeTokenBatch(&_CrossBridgeV2.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.PermitBridgeTokenBatch(&_CrossBridgeV2.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RegisterToken(&_CrossBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RegisterToken(&_CrossBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ReleasePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.ReleasePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RemovePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RemovePending(&_CrossBridgeV2.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RenounceRole(&_CrossBridgeV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RenounceRole(&_CrossBridgeV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RevokeRole(&_CrossBridgeV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RevokeRole(&_CrossBridgeV2.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RevokeRoleBatch(&_CrossBridgeV2.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.RevokeRoleBatch(&_CrossBridgeV2.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetBridgeExecutor(&_CrossBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetBridgeExecutor(&_CrossBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetBridgeVerifier(&_CrossBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetBridgeVerifier(&_CrossBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetChainPause(&_CrossBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetChainPause(&_CrossBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetCrossMintableERC20Code(&_CrossBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetCrossMintableERC20Code(&_CrossBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetCrossSupplyLimit(opts *bind.TransactOpts, _crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setCrossSupplyLimit", _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetCrossSupplyLimit(&_CrossBridgeV2.TransactOpts, _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetCrossSupplyLimit(&_CrossBridgeV2.TransactOpts, _crossSupplyLimit)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetDev(&_CrossBridgeV2.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetDev(&_CrossBridgeV2.TransactOpts, dev_)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetMaxExtraDataLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setMaxExtraDataLength", length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetMaxExtraDataLength(&_CrossBridgeV2.TransactOpts, length)
}

// SetMaxExtraDataLength is a paid mutator transaction binding the contract method 0x56ff54b0.
//
// Solidity: function setMaxExtraDataLength(uint256 length) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetMaxExtraDataLength(length *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetMaxExtraDataLength(&_CrossBridgeV2.TransactOpts, length)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetPause(&_CrossBridgeV2.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetPause(&_CrossBridgeV2.TransactOpts, set)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetPostCallGasReserve(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setPostCallGasReserve", value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetPostCallGasReserve(&_CrossBridgeV2.TransactOpts, value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetPostCallGasReserve(&_CrossBridgeV2.TransactOpts, value)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetTokenPause(&_CrossBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetTokenPause(&_CrossBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetVerificationDelay(&_CrossBridgeV2.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetVerificationDelay(&_CrossBridgeV2.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetVerificationDelayExpiration(&_CrossBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.SetVerificationDelayExpiration(&_CrossBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.UpgradeToAndCall(&_CrossBridgeV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridgeV2 *CrossBridgeV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridgeV2.Contract.UpgradeToAndCall(&_CrossBridgeV2.TransactOpts, newImplementation, data)
}

// CrossBridgeV2BridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeExecutorSetIterator struct {
	Event *CrossBridgeV2BridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2BridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2BridgeExecutorSet)
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
		it.Event = new(CrossBridgeV2BridgeExecutorSet)
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
func (it *CrossBridgeV2BridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2BridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2BridgeExecutorSet represents a BridgeExecutorSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*CrossBridgeV2BridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2BridgeExecutorSetIterator{contract: _CrossBridgeV2.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2BridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2BridgeExecutorSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseBridgeExecutorSet(log types.Log) (*CrossBridgeV2BridgeExecutorSet, error) {
	event := new(CrossBridgeV2BridgeExecutorSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2BridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeFinalizedIterator struct {
	Event *CrossBridgeV2BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2BridgeFinalized)
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
		it.Event = new(CrossBridgeV2BridgeFinalized)
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
func (it *CrossBridgeV2BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2BridgeFinalized represents a BridgeFinalized event raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeFinalized struct {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeV2BridgeFinalizedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2BridgeFinalizedIterator{contract: _CrossBridgeV2.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2BridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2BridgeFinalized)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseBridgeFinalized(log types.Log) (*CrossBridgeV2BridgeFinalized, error) {
	event := new(CrossBridgeV2BridgeFinalized)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2BridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeInitiatedIterator struct {
	Event *CrossBridgeV2BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2BridgeInitiated)
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
		it.Event = new(CrossBridgeV2BridgeInitiated)
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
func (it *CrossBridgeV2BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2BridgeInitiated represents a BridgeInitiated event raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeInitiated struct {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*CrossBridgeV2BridgeInitiatedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2BridgeInitiatedIterator{contract: _CrossBridgeV2.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2BridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2BridgeInitiated)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseBridgeInitiated(log types.Log) (*CrossBridgeV2BridgeInitiated, error) {
	event := new(CrossBridgeV2BridgeInitiated)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2BridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgePendingIterator struct {
	Event *CrossBridgeV2BridgePending // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2BridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2BridgePending)
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
		it.Event = new(CrossBridgeV2BridgePending)
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
func (it *CrossBridgeV2BridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2BridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2BridgePending represents a BridgePending event raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgePending struct {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeV2BridgePendingIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2BridgePendingIterator{contract: _CrossBridgeV2.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2BridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2BridgePending)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseBridgePending(log types.Log) (*CrossBridgeV2BridgePending, error) {
	event := new(CrossBridgeV2BridgePending)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2BridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeVerifierSetIterator struct {
	Event *CrossBridgeV2BridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2BridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2BridgeVerifierSet)
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
		it.Event = new(CrossBridgeV2BridgeVerifierSet)
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
func (it *CrossBridgeV2BridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2BridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2BridgeVerifierSet represents a BridgeVerifierSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2BridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*CrossBridgeV2BridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2BridgeVerifierSetIterator{contract: _CrossBridgeV2.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2BridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2BridgeVerifierSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseBridgeVerifierSet(log types.Log) (*CrossBridgeV2BridgeVerifierSet, error) {
	event := new(CrossBridgeV2BridgeVerifierSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2ChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2ChainPauseSetIterator struct {
	Event *CrossBridgeV2ChainPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2ChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2ChainPauseSet)
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
		it.Event = new(CrossBridgeV2ChainPauseSet)
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
func (it *CrossBridgeV2ChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2ChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2ChainPauseSet represents a ChainPauseSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2ChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*CrossBridgeV2ChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2ChainPauseSetIterator{contract: _CrossBridgeV2.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2ChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2ChainPauseSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseChainPauseSet(log types.Log) (*CrossBridgeV2ChainPauseSet, error) {
	event := new(CrossBridgeV2ChainPauseSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2CrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2CrossMintableERC20CodeSetIterator struct {
	Event *CrossBridgeV2CrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2CrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2CrossMintableERC20CodeSet)
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
		it.Event = new(CrossBridgeV2CrossMintableERC20CodeSet)
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
func (it *CrossBridgeV2CrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2CrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2CrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2CrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*CrossBridgeV2CrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2CrossMintableERC20CodeSetIterator{contract: _CrossBridgeV2.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2CrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2CrossMintableERC20CodeSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseCrossMintableERC20CodeSet(log types.Log) (*CrossBridgeV2CrossMintableERC20CodeSet, error) {
	event := new(CrossBridgeV2CrossMintableERC20CodeSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2CrossSupplyLimitSetIterator is returned from FilterCrossSupplyLimitSet and is used to iterate over the raw logs and unpacked data for CrossSupplyLimitSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2CrossSupplyLimitSetIterator struct {
	Event *CrossBridgeV2CrossSupplyLimitSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2CrossSupplyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2CrossSupplyLimitSet)
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
		it.Event = new(CrossBridgeV2CrossSupplyLimitSet)
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
func (it *CrossBridgeV2CrossSupplyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2CrossSupplyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2CrossSupplyLimitSet represents a CrossSupplyLimitSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2CrossSupplyLimitSet struct {
	CrossSupplyLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCrossSupplyLimitSet is a free log retrieval operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterCrossSupplyLimitSet(opts *bind.FilterOpts) (*CrossBridgeV2CrossSupplyLimitSetIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2CrossSupplyLimitSetIterator{contract: _CrossBridgeV2.contract, event: "CrossSupplyLimitSet", logs: logs, sub: sub}, nil
}

// WatchCrossSupplyLimitSet is a free log subscription operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchCrossSupplyLimitSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2CrossSupplyLimitSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2CrossSupplyLimitSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseCrossSupplyLimitSet(log types.Log) (*CrossBridgeV2CrossSupplyLimitSet, error) {
	event := new(CrossBridgeV2CrossSupplyLimitSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2DevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2DevSetIterator struct {
	Event *CrossBridgeV2DevSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2DevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2DevSet)
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
		it.Event = new(CrossBridgeV2DevSet)
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
func (it *CrossBridgeV2DevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2DevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2DevSet represents a DevSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2DevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*CrossBridgeV2DevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2DevSetIterator{contract: _CrossBridgeV2.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2DevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2DevSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseDevSet(log types.Log) (*CrossBridgeV2DevSet, error) {
	event := new(CrossBridgeV2DevSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossBridgeV2 contract.
type CrossBridgeV2EIP712DomainChangedIterator struct {
	Event *CrossBridgeV2EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2EIP712DomainChanged)
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
		it.Event = new(CrossBridgeV2EIP712DomainChanged)
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
func (it *CrossBridgeV2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossBridgeV2 contract.
type CrossBridgeV2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossBridgeV2EIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2EIP712DomainChangedIterator{contract: _CrossBridgeV2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2EIP712DomainChanged)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseEIP712DomainChanged(log types.Log) (*CrossBridgeV2EIP712DomainChanged, error) {
	event := new(CrossBridgeV2EIP712DomainChanged)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2ExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the CrossBridgeV2 contract.
type CrossBridgeV2ExtraCallExecutedIterator struct {
	Event *CrossBridgeV2ExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2ExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2ExtraCallExecuted)
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
		it.Event = new(CrossBridgeV2ExtraCallExecuted)
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
func (it *CrossBridgeV2ExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2ExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2ExtraCallExecuted represents a ExtraCallExecuted event raised by the CrossBridgeV2 contract.
type CrossBridgeV2ExtraCallExecuted struct {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (*CrossBridgeV2ExtraCallExecutedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2ExtraCallExecutedIterator{contract: _CrossBridgeV2.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 consumed, bytes returnData)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2ExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2ExtraCallExecuted)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseExtraCallExecuted(log types.Log) (*CrossBridgeV2ExtraCallExecuted, error) {
	event := new(CrossBridgeV2ExtraCallExecuted)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossBridgeV2 contract.
type CrossBridgeV2InitializedIterator struct {
	Event *CrossBridgeV2Initialized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2Initialized)
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
		it.Event = new(CrossBridgeV2Initialized)
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
func (it *CrossBridgeV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2Initialized represents a Initialized event raised by the CrossBridgeV2 contract.
type CrossBridgeV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CrossBridgeV2InitializedIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2InitializedIterator{contract: _CrossBridgeV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2Initialized) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2Initialized)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseInitialized(log types.Log) (*CrossBridgeV2Initialized, error) {
	event := new(CrossBridgeV2Initialized)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2ManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the CrossBridgeV2 contract.
type CrossBridgeV2ManualReleasedIterator struct {
	Event *CrossBridgeV2ManualReleased // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2ManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2ManualReleased)
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
		it.Event = new(CrossBridgeV2ManualReleased)
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
func (it *CrossBridgeV2ManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2ManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2ManualReleased represents a ManualReleased event raised by the CrossBridgeV2 contract.
type CrossBridgeV2ManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgeV2ManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2ManualReleasedIterator{contract: _CrossBridgeV2.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2ManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2ManualReleased)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseManualReleased(log types.Log) (*CrossBridgeV2ManualReleased, error) {
	event := new(CrossBridgeV2ManualReleased)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2MaxExtraDataLengthSetIterator is returned from FilterMaxExtraDataLengthSet and is used to iterate over the raw logs and unpacked data for MaxExtraDataLengthSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2MaxExtraDataLengthSetIterator struct {
	Event *CrossBridgeV2MaxExtraDataLengthSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2MaxExtraDataLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2MaxExtraDataLengthSet)
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
		it.Event = new(CrossBridgeV2MaxExtraDataLengthSet)
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
func (it *CrossBridgeV2MaxExtraDataLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2MaxExtraDataLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2MaxExtraDataLengthSet represents a MaxExtraDataLengthSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2MaxExtraDataLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxExtraDataLengthSet is a free log retrieval operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterMaxExtraDataLengthSet(opts *bind.FilterOpts) (*CrossBridgeV2MaxExtraDataLengthSetIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2MaxExtraDataLengthSetIterator{contract: _CrossBridgeV2.contract, event: "MaxExtraDataLengthSet", logs: logs, sub: sub}, nil
}

// WatchMaxExtraDataLengthSet is a free log subscription operation binding the contract event 0x1e7958e6a49c3baf1f1acef49309376a5222055229e2c6a41e70fa5fabdf48f1.
//
// Solidity: event MaxExtraDataLengthSet(uint256 length)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchMaxExtraDataLengthSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2MaxExtraDataLengthSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "MaxExtraDataLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2MaxExtraDataLengthSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseMaxExtraDataLengthSet(log types.Log) (*CrossBridgeV2MaxExtraDataLengthSet, error) {
	event := new(CrossBridgeV2MaxExtraDataLengthSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "MaxExtraDataLengthSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossBridgeV2 contract.
type CrossBridgeV2PausedIterator struct {
	Event *CrossBridgeV2Paused // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2Paused)
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
		it.Event = new(CrossBridgeV2Paused)
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
func (it *CrossBridgeV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2Paused represents a Paused event raised by the CrossBridgeV2 contract.
type CrossBridgeV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterPaused(opts *bind.FilterOpts) (*CrossBridgeV2PausedIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2PausedIterator{contract: _CrossBridgeV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2Paused) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2Paused)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParsePaused(log types.Log) (*CrossBridgeV2Paused, error) {
	event := new(CrossBridgeV2Paused)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2PendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the CrossBridgeV2 contract.
type CrossBridgeV2PendingRemovedIterator struct {
	Event *CrossBridgeV2PendingRemoved // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2PendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2PendingRemoved)
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
		it.Event = new(CrossBridgeV2PendingRemoved)
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
func (it *CrossBridgeV2PendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2PendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2PendingRemoved represents a PendingRemoved event raised by the CrossBridgeV2 contract.
type CrossBridgeV2PendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgeV2PendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2PendingRemovedIterator{contract: _CrossBridgeV2.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2PendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2PendingRemoved)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParsePendingRemoved(log types.Log) (*CrossBridgeV2PendingRemoved, error) {
	event := new(CrossBridgeV2PendingRemoved)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2PostCallGasReserveSetIterator is returned from FilterPostCallGasReserveSet and is used to iterate over the raw logs and unpacked data for PostCallGasReserveSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2PostCallGasReserveSetIterator struct {
	Event *CrossBridgeV2PostCallGasReserveSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2PostCallGasReserveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2PostCallGasReserveSet)
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
		it.Event = new(CrossBridgeV2PostCallGasReserveSet)
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
func (it *CrossBridgeV2PostCallGasReserveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2PostCallGasReserveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2PostCallGasReserveSet represents a PostCallGasReserveSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2PostCallGasReserveSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostCallGasReserveSet is a free log retrieval operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterPostCallGasReserveSet(opts *bind.FilterOpts) (*CrossBridgeV2PostCallGasReserveSetIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2PostCallGasReserveSetIterator{contract: _CrossBridgeV2.contract, event: "PostCallGasReserveSet", logs: logs, sub: sub}, nil
}

// WatchPostCallGasReserveSet is a free log subscription operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchPostCallGasReserveSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2PostCallGasReserveSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2PostCallGasReserveSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParsePostCallGasReserveSet(log types.Log) (*CrossBridgeV2PostCallGasReserveSet, error) {
	event := new(CrossBridgeV2PostCallGasReserveSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleAdminChangedIterator struct {
	Event *CrossBridgeV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2RoleAdminChanged)
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
		it.Event = new(CrossBridgeV2RoleAdminChanged)
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
func (it *CrossBridgeV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2RoleAdminChanged represents a RoleAdminChanged event raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossBridgeV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2RoleAdminChangedIterator{contract: _CrossBridgeV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2RoleAdminChanged)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseRoleAdminChanged(log types.Log) (*CrossBridgeV2RoleAdminChanged, error) {
	event := new(CrossBridgeV2RoleAdminChanged)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleGrantedIterator struct {
	Event *CrossBridgeV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2RoleGranted)
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
		it.Event = new(CrossBridgeV2RoleGranted)
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
func (it *CrossBridgeV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2RoleGranted represents a RoleGranted event raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeV2RoleGrantedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2RoleGrantedIterator{contract: _CrossBridgeV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2RoleGranted)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseRoleGranted(log types.Log) (*CrossBridgeV2RoleGranted, error) {
	event := new(CrossBridgeV2RoleGranted)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleRevokedIterator struct {
	Event *CrossBridgeV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2RoleRevoked)
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
		it.Event = new(CrossBridgeV2RoleRevoked)
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
func (it *CrossBridgeV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2RoleRevoked represents a RoleRevoked event raised by the CrossBridgeV2 contract.
type CrossBridgeV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeV2RoleRevokedIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2RoleRevokedIterator{contract: _CrossBridgeV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2RoleRevoked)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseRoleRevoked(log types.Log) (*CrossBridgeV2RoleRevoked, error) {
	event := new(CrossBridgeV2RoleRevoked)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2ThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossBridgeV2 contract.
type CrossBridgeV2ThresholdChangedIterator struct {
	Event *CrossBridgeV2ThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2ThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2ThresholdChanged)
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
		it.Event = new(CrossBridgeV2ThresholdChanged)
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
func (it *CrossBridgeV2ThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2ThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2ThresholdChanged represents a ThresholdChanged event raised by the CrossBridgeV2 contract.
type CrossBridgeV2ThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossBridgeV2ThresholdChangedIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2ThresholdChangedIterator{contract: _CrossBridgeV2.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2ThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2ThresholdChanged)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseThresholdChanged(log types.Log) (*CrossBridgeV2ThresholdChanged, error) {
	event := new(CrossBridgeV2ThresholdChanged)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2TokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the CrossBridgeV2 contract.
type CrossBridgeV2TokenPairRegisteredIterator struct {
	Event *CrossBridgeV2TokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2TokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2TokenPairRegistered)
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
		it.Event = new(CrossBridgeV2TokenPairRegistered)
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
func (it *CrossBridgeV2TokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2TokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2TokenPairRegistered represents a TokenPairRegistered event raised by the CrossBridgeV2 contract.
type CrossBridgeV2TokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*CrossBridgeV2TokenPairRegisteredIterator, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2TokenPairRegisteredIterator{contract: _CrossBridgeV2.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2TokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2TokenPairRegistered)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseTokenPairRegistered(log types.Log) (*CrossBridgeV2TokenPairRegistered, error) {
	event := new(CrossBridgeV2TokenPairRegistered)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2TokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2TokenPauseSetIterator struct {
	Event *CrossBridgeV2TokenPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2TokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2TokenPauseSet)
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
		it.Event = new(CrossBridgeV2TokenPauseSet)
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
func (it *CrossBridgeV2TokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2TokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2TokenPauseSet represents a TokenPauseSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2TokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeV2TokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2TokenPauseSetIterator{contract: _CrossBridgeV2.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2TokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2TokenPauseSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseTokenPauseSet(log types.Log) (*CrossBridgeV2TokenPauseSet, error) {
	event := new(CrossBridgeV2TokenPauseSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossBridgeV2 contract.
type CrossBridgeV2UnpausedIterator struct {
	Event *CrossBridgeV2Unpaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2Unpaused)
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
		it.Event = new(CrossBridgeV2Unpaused)
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
func (it *CrossBridgeV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2Unpaused represents a Unpaused event raised by the CrossBridgeV2 contract.
type CrossBridgeV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossBridgeV2UnpausedIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2UnpausedIterator{contract: _CrossBridgeV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2Unpaused)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseUnpaused(log types.Log) (*CrossBridgeV2Unpaused, error) {
	event := new(CrossBridgeV2Unpaused)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossBridgeV2 contract.
type CrossBridgeV2UpgradedIterator struct {
	Event *CrossBridgeV2Upgraded // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2Upgraded)
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
		it.Event = new(CrossBridgeV2Upgraded)
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
func (it *CrossBridgeV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2Upgraded represents a Upgraded event raised by the CrossBridgeV2 contract.
type CrossBridgeV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossBridgeV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2UpgradedIterator{contract: _CrossBridgeV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2Upgraded)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseUpgraded(log types.Log) (*CrossBridgeV2Upgraded, error) {
	event := new(CrossBridgeV2Upgraded)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2VerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2VerificationDelayExpirationSetIterator struct {
	Event *CrossBridgeV2VerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2VerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2VerificationDelayExpirationSet)
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
		it.Event = new(CrossBridgeV2VerificationDelayExpirationSet)
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
func (it *CrossBridgeV2VerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2VerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2VerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2VerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*CrossBridgeV2VerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2VerificationDelayExpirationSetIterator{contract: _CrossBridgeV2.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 expiration)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2VerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2VerificationDelayExpirationSet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseVerificationDelayExpirationSet(log types.Log) (*CrossBridgeV2VerificationDelayExpirationSet, error) {
	event := new(CrossBridgeV2VerificationDelayExpirationSet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeV2VerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the CrossBridgeV2 contract.
type CrossBridgeV2VerificationDelaySetIterator struct {
	Event *CrossBridgeV2VerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeV2VerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeV2VerificationDelaySet)
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
		it.Event = new(CrossBridgeV2VerificationDelaySet)
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
func (it *CrossBridgeV2VerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeV2VerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeV2VerificationDelaySet represents a VerificationDelaySet event raised by the CrossBridgeV2 contract.
type CrossBridgeV2VerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*CrossBridgeV2VerificationDelaySetIterator, error) {

	logs, sub, err := _CrossBridgeV2.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeV2VerificationDelaySetIterator{contract: _CrossBridgeV2.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridgeV2 *CrossBridgeV2Filterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *CrossBridgeV2VerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _CrossBridgeV2.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeV2VerificationDelaySet)
				if err := _CrossBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_CrossBridgeV2 *CrossBridgeV2Filterer) ParseVerificationDelaySet(log types.Log) (*CrossBridgeV2VerificationDelaySet, error) {
	event := new(CrossBridgeV2VerificationDelaySet)
	if err := _CrossBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
