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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeEthereumBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"supportPermit\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportPermit\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotSupportPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthereumBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthereumBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
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
		"90546099": "initializeEthereumBridge(address,address,uint8,uint256,address,uint256)",
		"91cf6d3e": "initializedAt()",
		"aa20ed47": "manualReleasePending(uint256,uint256)",
		"48a00ed8": "manualReleasePendingWithRecipient(uint256,uint256,address)",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"1313996b": "permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"85862d87": "registerToken(uint256,bool,bool,address,address)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"9f9b4f1c": "releasePendingBatch(uint256[],uint256[])",
		"8ae87c5c": "removePending(uint256,uint256)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"0b1d8d06": "setBridgeVerifier(address)",
		"b8aa837e": "setChainPause(uint256,bool)",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"bedb86fb": "setPause(bool)",
		"94ddc8c6": "setTokenPause(uint256,address,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615eaf6100f95f395f81816133360152818161335f01526134a40152615eaf5ff3fe6080604052600436106102b7575f3560e01c80638ae87c5c1161016b578063b2b49e2e116100c9578063d477f05f11610083578063d477f05f14610906578063d547741f14610925578063d5717fc514610944578063e2af6cd714610963578063f0702e8e14610982578063f4509643146109a1578063f698da25146109c0575f5ffd5b8063b2b49e2e1461084a578063b33eb36e14610869578063b7f3358d14610895578063b8aa837e146108b4578063bedb86fb146108d3578063cf56118e146108f2575f5ffd5b80639f9b4f1c116101255780639f9b4f1c14610752578063a217fddf14610771578063a3246ad314610784578063aa1bd0bc146107b0578063aa20ed47146107cf578063ad3cb1cc146107ee578063ae6893f81461082b575f5ffd5b80638ae87c5c146106a557806390546099146106c457806391cca3db146106e357806391cf6d3e1461070057806391d148541461071457806394ddc8c614610733575f5ffd5b80634f1ef28611610218578063670e6268116101d2578063670e6268146104fb578063792148741461051a578063814914b51461054657806384b0196e1461062d57806385862d871461065457806388d67d6d1461067357806389232a0014610686575f5ffd5b80634f1ef28614610462578063502cc5c01461047557806352d1902d146104945780635b605f5c146104a85780635c975abb146104d45780635fd262de146104e8575f5ffd5b806336568abe1161027457806336568abe14610382578063365d71e4146103a157806342cde4e8146103c057806348a00ed8146103e65780634be13f83146104055780634d5d0056146104305780634ee078ba14610443575f5ffd5b806301ffc9a7146102bb5780630b1d8d06146102ef5780631313996b146103105780631938e0f214610323578063248a9ca3146103365780632f2ff15d14610363575b5f5ffd5b3480156102c6575f5ffd5b506102da6102d5366004614b27565b6109d4565b60405190151581526020015b60405180910390f35b3480156102fa575f5ffd5b5061030e610309366004614b62565b610a0a565b005b61030e61031e366004614bc4565b610a92565b6102da610331366004614dcc565b610c25565b348015610341575f5ffd5b50610355610350366004614e85565b610f48565b6040519081526020016102e6565b34801561036e575f5ffd5b5061030e61037d366004614e9c565b610f68565b34801561038d575f5ffd5b5061030e61039c366004614e9c565b610f8a565b3480156103ac575f5ffd5b5061030e6103bb366004614eca565b610fc2565b3480156103cb575f5ffd5b506103d4611033565b60405160ff90911681526020016102e6565b3480156103f1575f5ffd5b5061030e610400366004614f8f565b61104e565b348015610410575f5ffd5b505f54610423906001600160a01b031681565b6040516102e69190614fc5565b6102da61043e366004615016565b6110f9565b34801561044e575f5ffd5b5061030e61045d3660046150b8565b6113dd565b61030e61047036600461514e565b611648565b348015610480575f5ffd5b5061030e61048f36600461519a565b611663565b34801561049f575f5ffd5b5061035561172f565b3480156104b3575f5ffd5b506104c76104c2366004614e85565b61174a565b6040516102e69190615216565b3480156104df575f5ffd5b506102da6118cf565b6102da6104f6366004615263565b6118e4565b348015610506575f5ffd5b506104236105153660046152eb565b6119be565b348015610525575f5ffd5b50610539610534366004614e85565b611a6c565b6040516102e6919061539d565b348015610551575f5ffd5b50610620610560366004614e9c565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152505f9182526006602090815260408084206001600160a01b039384168552825292839020835160e081018552815484168152600182015493841692810192909252600160a01b830460ff908116151594830194909452600160a81b9092049092161515606083015260028101546080830152600381015460a08301526004015460c082015290565b6040516102e691906153af565b348015610638575f5ffd5b50610641611a85565b6040516102e697969594939291906153eb565b34801561065f575f5ffd5b5061030e61066e366004615467565b611b2e565b6102da610681366004615549565b611b6d565b348015610691575f5ffd5b5061030e6106a0366004615679565b611c47565b3480156106b0575f5ffd5b5061030e6106bf3660046150b8565b611d57565b3480156106cf575f5ffd5b5061030e6106de3660046156bd565b611db5565b3480156106ee575f5ffd5b506033546001600160a01b0316610423565b34801561070b575f5ffd5b50603454610355565b34801561071f575f5ffd5b506102da61072e366004614e9c565b611f2f565b34801561073e575f5ffd5b5061030e61074d366004615727565b611f65565b34801561075d575f5ffd5b5061030e61076c36600461575b565b612047565b34801561077c575f5ffd5b506103555f81565b34801561078f575f5ffd5b506107a361079e366004614e85565b6120b8565b6040516102e691906157b4565b3480156107bb575f5ffd5b5061030e6107ca366004614e85565b6120fb565b3480156107da575f5ffd5b5061030e6107e93660046150b8565b61214e565b3480156107f9575f5ffd5b5061081e604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516102e691906157f4565b348015610836575f5ffd5b50610355610845366004614e85565b612193565b348015610855575f5ffd5b5061030e610864366004614eca565b6121af565b348015610874575f5ffd5b506108886108833660046150b8565b612220565b6040516102e6919061583a565b3480156108a0575f5ffd5b5061030e6108af3660046158ce565b61236d565b3480156108bf575f5ffd5b5061030e6108ce3660046158e7565b6123fc565b3480156108de575f5ffd5b5061030e6108ed36600461590a565b6124ae565b3480156108fd575f5ffd5b506105396124ee565b348015610911575f5ffd5b5061030e610920366004614b62565b6124ff565b348015610930575f5ffd5b5061030e61093f366004614e9c565b612587565b34801561094f575f5ffd5b5061035561095e366004614e85565b6125a3565b34801561096e575f5ffd5b5061030e61097d366004614b62565b6125bf565b34801561098d575f5ffd5b50603254610423906001600160a01b031681565b3480156109ac575f5ffd5b5061030e6109bb366004614e9c565b612671565b3480156109cb575f5ffd5b50610355612744565b5f6001600160e01b03198216637965db0b60e01b1480610a0457506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f516020615e5a5f395f51905f52610a218161274d565b6001600160a01b038216610a4857604051638219abe360e01b815260040160405180910390fd5b603280546001600160a01b0319166001600160a01b0384169081179091556040517fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3905f90a25050565b828114610ab2576040516329f517fd60e01b815260040160405180910390fd5b5f5b83811015610c1e57610c15858583818110610ad157610ad1615925565b9050602002810190610ae39190615939565b35868684818110610af657610af6615925565b9050602002810190610b089190615939565b610b19906040810190602001614b62565b878785818110610b2b57610b2b615925565b9050602002810190610b3d9190615939565b610b4e906080810190606001614b62565b888886818110610b6057610b60615925565b9050602002810190610b729190615939565b60800135898987818110610b8857610b88615925565b9050602002810190610b9a9190615939565b60a001358a8a88818110610bb057610bb0615925565b9050602002810190610bc29190615939565b60c001358b8b89818110610bd857610bd8615925565b9050602002810190610bea9190615939565b610bf89060e0810190615957565b8b8b8b818110610c0a57610c0a615925565b905060e002016110f9565b50600101610ab4565b5050505050565b5f610c2e61275a565b610c36612782565b610c5e610c496060870160408801614b62565b86355f908152600560205260409020906127cc565b610c6e6060870160408801614b62565b90610c9657604051633142cb1160e11b8152600401610c8d9190614fc5565b60405180910390fd5b505f348015610cc15760405163943892eb60e01b815260048101929092526024820152604401610c8d565b505084355f9081526004602052604081206002018054600101908190559050806020870135808214610d0f5760405163a6ab65c560e01b815260048101929092526024820152604401610c8d565b505f90507fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191487356020890135610d4b60608b0160408c01614b62565b610d5b60808c0160608d01614b62565b60808c0135610d6d60a08e018e615957565b604051602001610d84989796959493929190615999565b604051602081830303815290604052805190602001209050610da8818787876127ed565b5f80610dca8935610dbf60608c0160408d01614b62565b8b608001355f61295c565b90925090506001826009811115610de357610de3615806565b03610e1a57610e178935610dfd60608c0160408d01614b62565b610e0d60808d0160608e01614b62565b8c60800135612acc565b91505b6001826009811115610e2e57610e2e615806565b03610eb457610e4360608a0160408b01614b62565b6001600160a01b031660208a01358a357f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b610e8460808e0160608f01614b62565b604080516001600160a01b03909216825260808f01356020830152429082015260600160405180910390a4610f30565b610ebf898383612c0d565b610ecf60608a0160408b01614b62565b6001600160a01b031660208a01358a357f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f610f1060808e0160608f01614b62565b8d608001354288604051610f279493929190615a02565b60405180910390a45b6001945050505050610f40612dbf565b949350505050565b5f9081525f516020615e1a5f395f51905f52602052604090206001015490565b610f7182610f48565b610f7a8161274d565b610f848383612de5565b50505050565b6001600160a01b0381163314610fb35760405163334bd91960e11b815260040160405180910390fd5b610fbd8282612e5d565b505050565b8051825114610fe45760405163031206d560e51b815260040160405180910390fd5b5f5b8151811015610fbd5761102b83828151811061100457611004615925565b602002602001015183838151811061101e5761101e615925565b6020026020010151612587565b600101610fe6565b5f805f516020615dba5f395f51905f525b5460ff1692915050565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea096110788161274d565b611080612782565b5f8481526007602052604090206110979084612ecb565b83906110b9576040516373a1390160e11b8152600401610c8d91815260200190565b506110c5848484612ee2565b604051839085907fb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1905f90a3610f84612dbf565b5f61110261275a565b898961110e8282612fce565b611116612782565b6001600160a01b038b165f9081526009602052604090205460ff161561114f576040516373de95e960e11b815260040160405180910390fd5b61115c6020850185614b62565b6001600160a01b038c81169116148b6111786020870187614b62565b90916111aa576040516313f7c32b60e31b81526001600160a01b03928316600482015291166024820152604401610c8d565b506111bd90506040850160208601614b62565b6001600160a01b03168a6001600160a01b0316146111ee57604051630672aec160e01b815260040160405180910390fd5b6111fb8c8c8b8b8b61308e565b90985096508661120b898b615a41565b6112159190615a41565b60408501351015876112278a8c615a41565b6112319190615a41565b8560400135909161125e5760405163943892eb60e01b815260048101929092526024820152604401610c8d565b50506001600160a01b038b1663d505accf61127f6040870160208801614b62565b306040880135606089013561129a60a08b0160808c016158ce565b6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529490931660248501526044840191909152606483015260ff16608482015260a087013560a482015260c087013560c482015260e4015f604051808303815f87803b15801561130a575f5ffd5b505af115801561131c573d5f5f3e3d5ffd5b505050506113c26040518061010001604052808e81526020018d6001600160a01b031681526020018660200160208101906113579190614b62565b6001600160a01b031681526020018c6001600160a01b031681526020018b81526020018a815260200189815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505091525061323a565b600192506113ce612dbf565b50509998505050505050505050565b6113e561275a565b6113ed612782565b5f8281526007602052604090206114049082612ecb565b8190611426576040516373a1390160e11b8152600401610c8d91815260200190565b505f828152600860209081526040808320848452909152808220815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e08401526005820180548492918491610100850191906114a290615a54565b80601f01602080910402602001604051908101604052809291908181526020018280546114ce90615a54565b80156115195780601f106114f057610100808354040283529160200191611519565b820191905f5260205f20905b8154815290600101906020018083116114fc57829003601f168201915b505050919092525050508152600682015460209091019060ff16600981111561154457611544615806565b600981111561155557611555615806565b815260200160078201548152505090505f61158184835f015160400151845f015160800151600161295c565b509050600181600981111561159857611598615806565b148160098111156115ab576115ab615806565b6040516020016115bd91815260200190565b604051602081830303815290604052906115ea5760405162461bcd60e51b8152600401610c8d91906157f4565b50604082015115806115ff5750428260400151105b826040015142909161162d57604051637a88099160e11b815260048101929092526024820152604401610c8d565b505061163a84845f612ee2565b5050611644612dbf565b5050565b61165061332b565b611659826133cf565b61164482826133e6565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0961168d8161274d565b5f8481526007602052604090206116a49084612ecb565b83906116c6576040516373a1390160e11b8152600401610c8d91815260200190565b505f6116d28342615a41565b5f8681526008602090815260408083208884528252918290206007018390559051828152919250859187917fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949910160405180910390a35050505050565b5f611738613499565b505f516020615dfa5f395f51905f5290565b5f81815260056020526040812060609190611764906134e2565b90505f81516001600160401b0381111561178057611780614c5d565b6040519080825280602002602001820160405280156117e557816020015b6040805160e0810182525f8082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282525f1990920191018161179e5790505b5090505f5b82518110156118c75760065f8681526020019081526020015f205f84838151811061181757611817615925565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f20825160e08101845281548516815260018201549485169281019290925260ff600160a01b85048116151593830193909352600160a81b9093049091161515606082015260028201546080820152600382015460a082015260049091015460c082015282518390839081106118b4576118b4615925565b60209081029190910101526001016117ea565b509392505050565b5f805f516020615e3a5f395f51905f52611044565b5f6118ed61275a565b88886118f98282612fce565b611901612782565b61190e8b8b8a8a8a61308e565b60408051610100810182528e81526001600160a01b038e1660208201529299509097506119a491908101336001600160a01b031681526020018b6001600160a01b031681526020018a815260200189815260200188815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505091525061323a565b600192506119b0612dbf565b505098975050505050505050565b5f80546001600160a01b03166119e7576040516315aeca0d60e11b815260040160405180910390fd5b5f54604051637c469ea160e11b81526001600160a01b039091169063f88d3d4290611a1c908890889088908890600401615a8c565b6020604051808303815f875af1158015611a38573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a5c9190615ac9565b9050610f40855f60018488611b2e565b5f818152600760205260409020606090610a04906134e2565b5f60608082808083815f516020615dda5f395f51905f528054909150158015611ab057506001810154155b611af45760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606401610c8d565b611afc6134ee565b611b046135ae565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b7f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c611b588161274d565b611b6586868686866135ec565b505050505050565b82515f90859081148015611b815750808451145b8015611b8d5750808351145b611baa576040516329f517fd60e01b815260040160405180910390fd5b5f5b81811015611c3757611c2e888883818110611bc957611bc9615925565b9050602002810190611bdb9190615ae4565b878381518110611bed57611bed615925565b6020026020010151878481518110611c0757611c07615925565b6020026020010151878581518110611c2157611c21615925565b6020026020010151610c25565b50600101611bac565b5060019150505b95945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611c8b5750825b90505f826001600160401b03166001148015611ca65750303b155b905081158015611cb4575080155b15611cd25760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611cfc57845460ff60401b1916600160401b1785555b611d078888886138a7565b8315611d4d57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f516020615e5a5f395f51905f52611d6e8161274d565b611d76612782565b611d8083836139a1565b50604051829084907fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3905f90a3610fbd612dbf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611df95750825b90505f826001600160401b03166001148015611e145750303b155b905081158015611e22575080155b15611e405760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611e6a57845460ff60401b1916600160401b1785555b875f03611e8a57604051637d53fccd60e01b815260040160405180910390fd5b6001600160a01b038716611eb157604051635da4f0bf60e01b815260040160405180910390fd5b611ebc8b8b8b6138a7565b611ecb886001808a60016135ec565b8515611edc57611edc888888613b98565b8315611f2257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b5f9182525f516020615e1a5f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929611f8f8161274d565b5f848152600560205260409020611fa690846127cc565b8390611fc65760405163153096f360e11b8152600401610c8d9190614fc5565b505f8481526006602090815260408083206001600160a01b0387168085529252918290206001018054851515600160a81b0260ff60a81b19909116179055905185907f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea9061203990861515815260200190565b60405180910390a350505050565b8051825114612069576040516329f517fd60e01b815260040160405180910390fd5b5f5b8151811015610fbd576120b083828151811061208957612089615925565b60200260200101518383815181106120a3576120a3615925565b60200260200101516113dd565b60010161206b565b5f8181527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0060208190526040909120606091906120f4906134e2565b9392505050565b5f516020615e5a5f395f51905f526121128161274d565b60018290556040518281527fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b39060200160405180910390a15050565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea096121788161274d565b612180612782565b61218b83835f61104e565b610fbd612dbf565b5f818152600460205260408120600190810154610a0491615a41565b80518251146121d15760405163031206d560e51b815260040160405180910390fd5b5f5b8151811015610fbd576122188382815181106121f1576121f1615925565b602002602001015183838151811061220b5761220b615925565b6020026020010151610f68565b6001016121d3565b612228614a75565b5f83815260086020908152604080832085845290915290819020815161012081019092528054606083019081526001820154608084015260028201546001600160a01b0390811660a085015260038301541660c0840152600482015460e08401526005820180548492918491610100850191906122a490615a54565b80601f01602080910402602001604051908101604052809291908181526020018280546122d090615a54565b801561231b5780601f106122f25761010080835404028352916020019161231b565b820191905f5260205f20905b8154815290600101906020018083116122fe57829003601f168201915b505050919092525050508152600682015460209091019060ff16600981111561234657612346615806565b600981111561235757612357615806565b8152602001600782015481525050905092915050565b5f516020615e5a5f395f51905f526123848161274d565b8160ff165f036123a75760405163f0f15b9160e01b815260040160405180910390fd5b5f516020615dba5f395f51905f52805460ff841660ff199091168117825560408051918252517f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9181900360200190a1505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9296124268161274d565b612431600284612ecb565b83906124535760405163ac7dbbfd60e01b8152600401610c8d91815260200190565b505f83815260046020908152604091829020600301805460ff1916851515908117909155915191825284917f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675910160405180910390a2505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9296124d88161274d565b81156124e657611644613bfe565b611644613c5a565b60606124fa60026134e2565b905090565b5f516020615e5a5f395f51905f526125168161274d565b6001600160a01b03821661253d57604051638219abe360e01b815260040160405180910390fd5b603380546001600160a01b0319166001600160a01b0384169081179091556040517f76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871905f90a25050565b61259082610f48565b6125998161274d565b610f848383612e5d565b5f81815260046020526040812060020154610a04906001615a41565b5f516020615e5a5f395f51905f526125d68161274d565b5f546001600160a01b031680156126015760405163f6c75f8f60e01b8152600401610c8d9190614fc5565b506001600160a01b038216612629576040516302bfb33d60e51b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b038416908117825560405190917fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee91a25050565b5f516020615e5a5f395f51905f526126888161274d565b5f83815260056020526040902061269f9083613c9f565b82906126bf5760405163153096f360e11b8152600401610c8d9190614fc5565b505f8381526006602090815260408083206001600160a01b038616808552925280832080546001600160a01b03191681556001810180546001600160b01b0319169055600281018490556003810184905560040183905551909185917fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d9190a3505050565b5f6124fa613cb3565b6127578133613cbc565b50565b6127626118cf565b156127805760405163d93c066560e01b815260040160405180910390fd5b565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016127c657604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381165f90815260018301602052604081205415156120f4565b825182515f516020615dba5f395f51905f5291908114801561280f5750825181145b61282c5760405163b3f07a3960e01b815260040160405180910390fd5b8154819060ff1681101561285657604051631aebd17960e11b8152600401610c8d91815260200190565b505f80805b83811015612926575f6128c689838151811061287957612879615925565b602002602001015189848151811061289357612893615925565b60200260200101518985815181106128ad576128ad615925565b60200260200101516128be8e613ce7565b929190613d13565b9050806001600160a01b0316836001600160a01b031610801561290e575061290e7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892682611f2f565b1561291d578360010193508092505b5060010161285b565b508354829060ff1681101561295157604051631aebd17960e11b8152600401610c8d91815260200190565b505050505050505050565b6032545f9081906001600160a01b0316612989576040516330d45fb160e01b815260040160405180910390fd5b5f8681526006602090815260408083206001600160a01b03808a16855290835292819020815160e08101835281548516815260018201549485169381019390935260ff600160a01b85048116151592840192909252600160a81b90930416158015606083015260028301546080830152600383015460a083015260049092015460c082015290612a205760025f9250925050612ac3565b83612abc57603254604051633f4bdec560e01b81526001600160a01b0390911690633f4bdec590612a579089908990600401615af8565b6020604051808303815f875af1158015612a73573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a979190615b11565b92506001836009811115612aad57612aad615806565b14612ab757600191505b612ac1565b600192505b505b94509492505050565b5f5f196001600160a01b03851601612b1f575f612af8848460405180602001604052805f815250613d3f565b905080612b09576005915050610f40565b612b1586600185613e13565b6001915050610f40565b8115610f40575f8581526006602090815260408083206001600160a01b0388168452909152902060010154606090600160a01b900460ff1615612ba4578383604051602401612b6f929190615af8565b60408051601f198184030181529190526020810180516001600160e01b031663a9059cbb60e01b179052600592509050612be8565b8383604051602401612bb7929190615af8565b60408051601f198184030181529190526020810180516001600160e01b03166340c10f1960e01b1790526006925090505b612bf3855f83613d3f565b15612c075760019150612c07868685613e13565b50610f40565b82355f908152600760209081526040909120612c2b91850135613e6e565b836020013590612c51576040516307a5c91d60e31b8152600401610c8d91815260200190565b50604051806060016040528084612c6790615b2f565b8152602001836009811115612c7e57612c7e615806565b815260200182612c8e575f612c9b565b600154612c9b9042615a41565b905283355f908152600860209081526040808320828801358452825291829020835180518255918201516001820155918101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a081015182906005820190612d239082615bf9565b505050602082015160068201805460ff19166001836009811115612d4957612d49615806565b021790555060409182015160079091015583355f908152600660205281812090918290612d7c9060608801908801614b62565b6001600160a01b03166001600160a01b031681526020019081526020015f2090508360800135816002015f828254612db49190615a41565b909155505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f612df08383613e79565b90508015610a04575f8381527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0060208190526040909120612e319084613f1a565b83859091612e545760405163d180cb3160e01b8152600401610c8d929190615af8565b50505092915050565b5f612e688383613f2e565b90508015610a04575f8381527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0060208190526040909120612ea99084613c9f565b83859091612e545760405162a95f1b60e31b8152600401610c8d929190615af8565b5f81815260018301602052604081205415156120f4565b5f612eed84846139a1565b90506001600160a01b038216612f0557806060015191505b5f612f1d825f01518360400151858560800151612acc565b90506001816009811115612f3357612f33615806565b148190612f545760405163290cd55f60e01b8152600401610c8d9190615cb3565b5081604001516001600160a01b03168260200151835f01517f94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b86866080015142604051612fbf939291906001600160a01b039390931683526020830191909152604082015260600190565b60405180910390a45050505050565b5f828152600560205260409020612fe590826127cc565b81906130055760405163153096f360e11b8152600401610c8d9190614fc5565b505f82815260046020526040902060030154829060ff161561303d57604051636fc24b7960e11b8152600401610c8d91815260200190565b505f8281526006602090815260408083206001600160a01b03851684529091529020600101548190600160a81b900460ff1615610fbd576040516338384f6f60e11b8152600401610c8d9190614fc5565b5f8581526006602090815260408083206001600160a01b038089168552908352818420825160e08101845281548316815260018201549283169481019490945260ff600160a01b830481161515938501849052600160a81b9092049091161515606084015260028101546080840152600381015460a08401526004015460c083015282919061314b5760c08101518690818110156131485760405163943892eb60e01b815260048101929092526024820152604401610c8d565b50505b6032546001600160a01b0316613174576040516330d45fb160e01b815260040160405180910390fd5b6032546040516337dba1f760e21b8152600481018a90526001600160a01b038981166024830152604482018990525f92169063df6e87dc90606401606060405180830381865afa1580156131ca573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131ee9190615cc1565b909550935090508087108015906132055750838610155b80156132115750828510155b61322e576040516358e8878560e01b815260040160405180910390fd5b50509550959350505050565b80515f90815260046020908152604080832060019081018054820190819055855185526006845282852084870180516001600160a01b0390811688529190955294839020909101548551935192860151608087015160c088015160a0890151949793909316956132b695909490936132b191615a41565b613fa7565b82604001516001600160a01b031682845f01517f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708866020015185886060015189608001518a60a001518b60c001518c60e001514260405161331e989796959493929190615cec565b60405180910390a4505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806133b157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166133a55f516020615dfa5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156127805760405163703e46dd60e11b815260040160405180910390fd5b5f516020615e5a5f395f51905f526116448161274d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613440575060408051601f3d908101601f1916820190925261343d91810190615d4a565b60015b61345f5781604051634c9c8ce360e01b8152600401610c8d9190614fc5565b5f516020615dfa5f395f51905f52811461348f57604051632a87526960e21b815260048101829052602401610c8d565b610fbd83836141ab565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146127805760405163703e46dd60e11b815260040160405180910390fd5b60605f6120f483614200565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020615dda5f395f51905f529161352c90615a54565b80601f016020809104026020016040519081016040528092919081815260200182805461355890615a54565b80156135a35780601f1061357a576101008083540402835291602001916135a3565b820191905f5260205f20905b81548152906001019060200180831161358657829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020615dda5f395f51905f529161352c90615a54565b845f0361360c576040516302bfb33d60e51b815260040160405180910390fd5b6001600160a01b038216613633576040516302bfb33d60e51b815260040160405180910390fd5b6001600160a01b03811661365a576040516302bfb33d60e51b815260040160405180910390fd5b613665600286613e6e565b156136c257604080516080810182528681525f6020808301828152838501838152606085018481528b855260049093529490922092518355905160018301559151600282015590516003909101805460ff19169115159190911790555b5f8581526005602052604090206136d99083613f1a565b82906136f9576040516311dd05f360e31b8152600401610c8d9190614fc5565b506040518060e00160405280836001600160a01b03168152602001826001600160a01b0316815260200185151581526020015f151581526020015f81526020015f81526020015f81525060065f8781526020019081526020015f205f846001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055506080820151816002015560a0820151816003015560c0820151816004015590505082613854576001600160a01b0382165f908152600960205260409020805460ff191660011790555b806001600160a01b0316826001600160a01b0316867f6509d37a38735e68576c1146ef1653f847a0054869b85d51d170dd430899359d8787604051612fbf92919091151582521515602082015260400190565b6138af614259565b6001600160a01b0383166138d657604051638219abe360e01b815260040160405180910390fd5b6001600160a01b0382166138fd57604051638219abe360e01b815260040160405180910390fd5b8060ff165f03613920576040516338854f1160e21b815260040160405180910390fd5b6139286142a2565b6139306142aa565b6139386142ba565b613941836142ca565b61394a816142db565b613952614370565b603380546001600160a01b0319166001600160a01b0384169081179091556040517f76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871905f90a250504360345550565b6139a9614a9a565b5f8381526007602052604090206139c09083614381565b82906139e257604051637f11bea960e01b8152600401610c8d91815260200190565b505f838152600860209081526040808320858452825291829020825160c0810184528154815260018201549281019290925260028101546001600160a01b0390811693830193909352600381015490921660608201526004820154608082015260058201805491929160a084019190613a5a90615a54565b80601f0160208091040260200160405190810160405280929190818152602001828054613a8690615a54565b8015613ad15780601f10613aa857610100808354040283529160200191613ad1565b820191905f5260205f20905b815481529060010190602001808311613ab457829003601f168201915b505050919092525050505f848152600660209081526040808320818501516001600160a01b03168452909152812060808301516002820180549495509193909290613b1d908490615d61565b90915550505f8481526008602090815260408083208684529091528120818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055908181613b7b6005830182614add565b50505060068101805460ff191690555f6007909101555092915050565b5f8381526006602090815260408083206001600160a01b038616845290915290206001810154600160a01b900460ff1615613beb5781816003015f828254613be09190615a41565b90915550610f849050565b81816004015f828254612db49190615d61565b613c0661275a565b5f516020615e3a5f395f51905f52805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b604051613c4f9190614fc5565b60405180910390a150565b613c6261438c565b5f516020615e3a5f395f51905f52805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613c42565b5f6120f4836001600160a01b0384166143b1565b5f6124fa61448b565b613cc68282611f2f565b61164457808260405163e2517d3f60e01b8152600401610c8d929190615af8565b5f610a04613cf3613cb3565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613d23888888886144fe565b925092509250613d3382826145c6565b50909695505050505050565b5f82471015613d6157604051632b60b36f60e21b815260040160405180910390fd5b6060846001600160a01b03168484604051613d7c9190615d74565b5f6040518083038185875af1925050503d805f8114613db6576040519150601f19603f3d011682016040523d82523d5f602084013e613dbb565b606091505b50909250905081613dcf575f9150506120f4565b835f03613e085780515f03613dfd57846001600160a01b03163b5f03613df8575f9150506120f4565b613e08565b6020015190506120f4565b506001949350505050565b5f8381526006602090815260408083206001600160a01b038616845290915290206001810154600160a01b900460ff1615613e5b5781816003015f828254613be09190615d61565b81816004015f828254612db49190615a41565b5f6120f4838361467e565b5f5f516020615e1a5f395f51905f52613e928484611f2f565b613f11575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613ec73390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a04565b5f915050610a04565b5f6120f4836001600160a01b03841661467e565b5f5f516020615e1a5f395f51905f52613f478484611f2f565b15613f11575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a04565b5f196001600160a01b0385160161405057613fc28183615a41565b3414613fce8284615a41565b349091613ff75760405163943892eb60e01b815260048101929092526024820152604401610c8d565b5050801561404b5760335460408051602081019091525f80825291614029916001600160a01b03909116908490613d3f565b9050806140495760405163559d141b60e11b815260040160405180910390fd5b505b6141a0565b5f34801561407a5760405163943892eb60e01b815260048101929092526024820152604401610c8d565b5061409e9050833061408c8486615a41565b6001600160a01b0388169291906146ca565b80156140be576033546140be906001600160a01b03868116911683614731565b5f8581526006602090815260408083206001600160a01b0388168452909152902060010154600160a01b900460ff166141a057604051632770a7eb60e21b81526001600160a01b03851690639dc29fac9061411f9030908690600401615af8565b6020604051808303815f875af115801561413b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061415f9190615d8a565b84848490919261419c57604051639ac2f56d60e01b81526001600160a01b0393841660048201529290911660248301526044820152606401610c8d565b5050505b610c1e858584613b98565b6141b482614757565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156141f857610fbd82826147b1565b61164461481a565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561424d57602002820191905f5260205f20905b815481526020019060010190808311614239575b50505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661278057604051631afcd79f60e31b815260040160405180910390fd5b612780614259565b6142b2614259565b612780614839565b6142c2614259565b612780614859565b6142d2614259565b61275781614861565b6142e3614259565b8060ff165f036143065760405163f0f15b9160e01b815260040160405180910390fd5b61434e604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061487b565b5f516020615dba5f395f51905f52805460ff191660ff92909216919091179055565b614378614259565b62015180600155565b5f6120f483836143b1565b6143946118cf565b61278057604051638dfc202b60e01b815260040160405180910390fd5b5f8181526001830160205260408120548015613f11575f6143d3600183615d61565b85549091505f906143e690600190615d61565b9050808214614445575f865f01828154811061440457614404615925565b905f5260205f200154905080875f01848154811061442457614424615925565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061445657614456615da5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a04565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6144b561488d565b6144bd6148f5565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561453757505f915060039050826145bc565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614588573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166145b357505f9250600191508290506145bc565b92505f91508190505b9450945094915050565b5f8260038111156145d9576145d9615806565b036145e2575050565b60018260038111156145f6576145f6615806565b036146145760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561462857614628615806565b036146495760405163fce698f760e01b815260048101829052602401610c8d565b600382600381111561465d5761465d615806565b03611644576040516335e2f38360e21b815260048101829052602401610c8d565b5f8181526001830160205260408120546146c357508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a04565b505f610a04565b6040516001600160a01b038481166024830152838116604483015260648201839052610f849186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614937565b610fbd83846001600160a01b031663a9059cbb85856040516024016146ff929190615af8565b806001600160a01b03163b5f036147835780604051634c9c8ce360e01b8152600401610c8d9190614fc5565b5f516020615dfa5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516147cd9190615d74565b5f60405180830381855af49150503d805f8114614805576040519150601f19603f3d011682016040523d82523d5f602084013e61480a565b606091505b5091509150611c3e85838361499a565b34156127805760405163b398979f60e01b815260040160405180910390fd5b614841614259565b5f516020615e3a5f395f51905f52805460ff19169055565b612dbf614259565b614869614259565b6148716142a2565b6116445f82612de5565b614883614259565b61164482826149ed565b5f5f516020615dda5f395f51905f52816148a56134ee565b8051909150156148bd57805160209091012092915050565b815480156148cc579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f516020615dda5f395f51905f528161490d6135ae565b80519091501561492557805160209091012092915050565b600182015480156148cc579392505050565b5f5f60205f8451602086015f885af180614956576040513d5f823e3d81fd5b50505f513d9150811561496d57806001141561497a565b6001600160a01b0384163b155b15610f845783604051635274afe760e01b8152600401610c8d9190614fc5565b6060826149af576149aa82614a4c565b6120f4565b81511580156149c657506001600160a01b0384163b155b156149e65783604051639996b31560e01b8152600401610c8d9190614fc5565b50806120f4565b6149f5614259565b5f516020615dda5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102614a2e8482615bf9565b5060038101614a3d8382615bf9565b505f8082556001909101555050565b805115614a5c5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b6040518060600160405280614a88614a9a565b81526020015f81526020015f81525090565b6040518060c001604052805f81526020015f81526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f8152602001606081525090565b508054614ae990615a54565b5f825580601f10614af8575050565b601f0160209004905f5260205f209081019061275791905b80821115614b23575f8155600101614b10565b5090565b5f60208284031215614b37575f5ffd5b81356001600160e01b0319811681146120f4575f5ffd5b6001600160a01b0381168114612757575f5ffd5b5f60208284031215614b72575f5ffd5b81356120f481614b4e565b5f5f83601f840112614b8d575f5ffd5b5081356001600160401b03811115614ba3575f5ffd5b6020830191508360208260051b8501011115614bbd575f5ffd5b9250929050565b5f5f5f5f60408587031215614bd7575f5ffd5b84356001600160401b03811115614bec575f5ffd5b614bf887828801614b7d565b90955093505060208501356001600160401b03811115614c16575f5ffd5b8501601f81018713614c26575f5ffd5b80356001600160401b03811115614c3b575f5ffd5b87602060e083028401011115614c4f575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715614c9357614c93614c5d565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614cc157614cc1614c5d565b604052919050565b5f6001600160401b03821115614ce157614ce1614c5d565b5060051b60200190565b803560ff81168114614cfb575f5ffd5b919050565b5f82601f830112614d0f575f5ffd5b8135614d22614d1d82614cc9565b614c99565b8082825260208201915060208360051b860101925085831115614d43575f5ffd5b602085015b83811015614d6757614d5981614ceb565b835260209283019201614d48565b5095945050505050565b5f82601f830112614d80575f5ffd5b8135614d8e614d1d82614cc9565b8082825260208201915060208360051b860101925085831115614daf575f5ffd5b602085015b83811015614d67578035835260209283019201614db4565b5f5f5f5f60808587031215614ddf575f5ffd5b84356001600160401b03811115614df4575f5ffd5b850160c08188031215614e05575f5ffd5b935060208501356001600160401b03811115614e1f575f5ffd5b614e2b87828801614d00565b93505060408501356001600160401b03811115614e46575f5ffd5b614e5287828801614d71565b92505060608501356001600160401b03811115614e6d575f5ffd5b614e7987828801614d71565b91505092959194509250565b5f60208284031215614e95575f5ffd5b5035919050565b5f5f60408385031215614ead575f5ffd5b823591506020830135614ebf81614b4e565b809150509250929050565b5f5f60408385031215614edb575f5ffd5b82356001600160401b03811115614ef0575f5ffd5b614efc85828601614d71565b92505060208301356001600160401b03811115614f17575f5ffd5b8301601f81018513614f27575f5ffd5b8035614f35614d1d82614cc9565b8082825260208201915060208360051b850101925087831115614f56575f5ffd5b6020840193505b82841015614f81578335614f7081614b4e565b825260209384019390910190614f5d565b809450505050509250929050565b5f5f5f60608486031215614fa1575f5ffd5b83359250602084013591506040840135614fba81614b4e565b809150509250925092565b6001600160a01b0391909116815260200190565b5f5f83601f840112614fe9575f5ffd5b5081356001600160401b03811115614fff575f5ffd5b602083019150836020828501011115614bbd575f5ffd5b5f5f5f5f5f5f5f5f5f898b036101c0811215615030575f5ffd5b8a35995060208b013561504281614b4e565b985060408b013561505281614b4e565b975060608b0135965060808b0135955060a08b0135945060c08b01356001600160401b03811115615081575f5ffd5b61508d8d828e01614fd9565b90955093505060e060df19820112156150a4575f5ffd5b5060e08a0190509295985092959850929598565b5f5f604083850312156150c9575f5ffd5b50508035926020909101359150565b5f5f6001600160401b038411156150f1576150f1614c5d565b50601f8301601f191660200161510681614c99565b91505082815283838301111561511a575f5ffd5b828260208301375f602084830101529392505050565b5f82601f83011261513f575f5ffd5b6120f4838335602085016150d8565b5f5f6040838503121561515f575f5ffd5b823561516a81614b4e565b915060208301356001600160401b03811115615184575f5ffd5b61519085828601615130565b9150509250929050565b5f5f5f606084860312156151ac575f5ffd5b505081359360208301359350604090920135919050565b80516001600160a01b039081168352602080830151909116908301526040808201511515908301526060808201511515908301526080808201519083015260a0818101519083015260c090810151910152565b602080825282518282018190525f918401906040840190835b81811015615258576152428385516151c3565b6020939093019260e0929092019160010161522f565b509095945050505050565b5f5f5f5f5f5f5f5f60e0898b03121561527a575f5ffd5b88359750602089013561528c81614b4e565b9650604089013561529c81614b4e565b9550606089013594506080890135935060a0890135925060c08901356001600160401b038111156152cb575f5ffd5b6152d78b828c01614fd9565b999c989b5096995094979396929594505050565b5f5f5f5f608085870312156152fe575f5ffd5b84359350602085013561531081614b4e565b925060408501356001600160401b0381111561532a575f5ffd5b8501601f8101871361533a575f5ffd5b615349878235602084016150d8565b92505061535860608601614ceb565b905092959194509250565b5f8151808452602084019350602083015f5b82811015615393578151865260209586019590910190600101615375565b5093949350505050565b602081525f6120f46020830184615363565b60e08101610a0482846151c3565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61540960e08301896153bd565b828103604084015261541b81896153bd565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061544c8185615363565b9a9950505050505050505050565b8015158114612757575f5ffd5b5f5f5f5f5f60a0868803121561547b575f5ffd5b85359450602086013561548d8161545a565b9350604086013561549d8161545a565b925060608601356154ad81614b4e565b915060808601356154bd81614b4e565b809150509295509295909350565b5f82601f8301126154da575f5ffd5b81356154e8614d1d82614cc9565b8082825260208201915060208360051b860101925085831115615509575f5ffd5b602085015b83811015614d675780356001600160401b0381111561552b575f5ffd5b61553a886020838a0101614d71565b8452506020928301920161550e565b5f5f5f5f5f6080868803121561555d575f5ffd5b85356001600160401b03811115615572575f5ffd5b61557e88828901614b7d565b90965094505060208601356001600160401b0381111561559c575f5ffd5b8601601f810188136155ac575f5ffd5b80356155ba614d1d82614cc9565b8082825260208201915060208360051b85010192508a8311156155db575f5ffd5b602084015b8381101561561b5780356001600160401b038111156155fd575f5ffd5b61560c8d602083890101614d00565b845250602092830192016155e0565b50955050505060408601356001600160401b03811115615639575f5ffd5b615645888289016154cb565b92505060608601356001600160401b03811115615660575f5ffd5b61566c888289016154cb565b9150509295509295909350565b5f5f5f6060848603121561568b575f5ffd5b833561569681614b4e565b925060208401356156a681614b4e565b91506156b460408501614ceb565b90509250925092565b5f5f5f5f5f5f60c087890312156156d2575f5ffd5b86356156dd81614b4e565b955060208701356156ed81614b4e565b94506156fb60408801614ceb565b935060608701359250608087013561571281614b4e565b9598949750929591949360a090920135925050565b5f5f5f60608486031215615739575f5ffd5b83359250602084013561574b81614b4e565b91506040840135614fba8161545a565b5f5f6040838503121561576c575f5ffd5b82356001600160401b03811115615781575f5ffd5b61578d85828601614d71565b92505060208301356001600160401b038111156157a8575f5ffd5b61519085828601614d71565b602080825282518282018190525f918401906040840190835b818110156152585783516001600160a01b03168352602093840193909201916001016157cd565b602081525f6120f460208301846153bd565b634e487b7160e01b5f52602160045260245ffd5b600a811061583657634e487b7160e01b5f52602160045260245ffd5b9052565b602081525f82516060602084015280516080840152602081015160a084015260018060a01b0360408201511660c084015260018060a01b0360608201511660e0840152608081015161010084015260a0810151905060c06101208401526158a56101408401826153bd565b905060208401516158b9604085018261581a565b50604084015160608401528091505092915050565b5f602082840312156158de575f5ffd5b6120f482614ceb565b5f5f604083850312156158f8575f5ffd5b823591506020830135614ebf8161545a565b5f6020828403121561591a575f5ffd5b81356120f48161545a565b634e487b7160e01b5f52603260045260245ffd5b5f823560fe1983360301811261594d575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261596c575f5ffd5b8301803591506001600160401b03821115615985575f5ffd5b602001915036819003821315614bbd575f5ffd5b88815260208101889052604081018790526001600160a01b0386811660608301528516608082015260a0810184905260e060c08201819052810182905281836101008301375f81830161010090810191909152601f909201601f19160101979650505050505050565b6001600160a01b0385168152602081018490526040810183905260808101611c3e606083018461581a565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610a0457610a04615a2d565b600181811c90821680615a6857607f821691505b602082108103615a8657634e487b7160e01b5f52602260045260245ffd5b50919050565b8481526001600160a01b03841660208201526080604082018190525f90615ab5908301856153bd565b905060ff8316606083015295945050505050565b5f60208284031215615ad9575f5ffd5b81516120f481614b4e565b5f823560be1983360301811261594d575f5ffd5b6001600160a01b03929092168252602082015260400190565b5f60208284031215615b21575f5ffd5b8151600a81106120f4575f5ffd5b5f60c08236031215615b3f575f5ffd5b615b47614c71565b82358152602080840135908201526040830135615b6381614b4e565b60408201526060830135615b7681614b4e565b60608201526080838101359082015260a08301356001600160401b03811115615b9d575f5ffd5b615ba936828601615130565b60a08301525092915050565b601f821115610fbd57805f5260205f20601f840160051c81016020851015615bda5750805b601f840160051c820191505b81811015610c1e575f8155600101615be6565b81516001600160401b03811115615c1257615c12614c5d565b615c2681615c208454615a54565b84615bb5565b6020601f821160018114615c58575f8315615c415750848201515b5f19600385901b1c1916600184901b178455610c1e565b5f84815260208120601f198516915b82811015615c875787850151825560209485019460019092019101615c67565b5084821015615ca457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b60208101610a04828461581a565b5f5f5f60608486031215615cd3575f5ffd5b5050815160208301516040909301519094929350919050565b6001600160a01b038981168252888116602083015287166040820152606081018690526080810185905260a0810184905261010060c082018190525f90615d35908301856153bd565b90508260e08301529998505050505050505050565b5f60208284031215615d5a575f5ffd5b5051919050565b81810381811115610a0457610a04615a2d565b5f82518060208501845e5f920191825250919050565b5f60208284031215615d9a575f5ffd5b81516120f48161545a565b634e487b7160e01b5f52603160045260245ffdfe303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220570dcae3b6fe234495e18c136a30e22064e9dd805e92754477dd90591d08fdab64736f6c634300081c0033",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EthereumBridge.Contract.DEFAULTADMINROLE(&_EthereumBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EthereumBridge.Contract.DEFAULTADMINROLE(&_EthereumBridge.CallOpts)
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

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_EthereumBridge *EthereumBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_EthereumBridge *EthereumBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _EthereumBridge.Contract.AllPendingIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _EthereumBridge.Contract.AllPendingIndex(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
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
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_EthereumBridge *EthereumBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.AllTokenPairs(&_EthereumBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_EthereumBridge *EthereumBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.AllTokenPairs(&_EthereumBridge.CallOpts, remoteChainID)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) BridgeVerifier() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeVerifier(&_EthereumBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _EthereumBridge.Contract.BridgeVerifier(&_EthereumBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _EthereumBridge.Contract.CrossMintableERC20Code(&_EthereumBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _EthereumBridge.Contract.CrossMintableERC20Code(&_EthereumBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_EthereumBridge *EthereumBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_EthereumBridge *EthereumBridgeSession) Dev() (common.Address, error) {
	return _EthereumBridge.Contract.Dev(&_EthereumBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_EthereumBridge *EthereumBridgeCallerSession) Dev() (common.Address, error) {
	return _EthereumBridge.Contract.Dev(&_EthereumBridge.CallOpts)
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

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_EthereumBridge *EthereumBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_EthereumBridge *EthereumBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _EthereumBridge.Contract.GetPendingArguments(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_EthereumBridge *EthereumBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _EthereumBridge.Contract.GetPendingArguments(&_EthereumBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthereumBridge *EthereumBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EthereumBridge.Contract.GetRoleAdmin(&_EthereumBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EthereumBridge *EthereumBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EthereumBridge.Contract.GetRoleAdmin(&_EthereumBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EthereumBridge *EthereumBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EthereumBridge *EthereumBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EthereumBridge.Contract.GetRoleMembers(&_EthereumBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_EthereumBridge *EthereumBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _EthereumBridge.Contract.GetRoleMembers(&_EthereumBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
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
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_EthereumBridge *EthereumBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.GetTokenPair(&_EthereumBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_EthereumBridge *EthereumBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _EthereumBridge.Contract.GetTokenPair(&_EthereumBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EthereumBridge.Contract.HasRole(&_EthereumBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EthereumBridge.Contract.HasRole(&_EthereumBridge.CallOpts, role, account)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EthereumBridge *EthereumBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EthereumBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EthereumBridge *EthereumBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EthereumBridge.Contract.SupportsInterface(&_EthereumBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EthereumBridge *EthereumBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EthereumBridge.Contract.SupportsInterface(&_EthereumBridge.CallOpts, interfaceId)
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

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.BridgeToken(&_EthereumBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _EthereumBridge.Contract.BridgeToken(&_EthereumBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.GrantRole(&_EthereumBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.GrantRole(&_EthereumBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.GrantRoleBatch(&_EthereumBridge.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.GrantRoleBatch(&_EthereumBridge.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner_, address dev_, uint8 _threshold) returns()
func (_EthereumBridge *EthereumBridgeTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, dev_ common.Address, _threshold uint8) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "initialize", owner_, dev_, _threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner_, address dev_, uint8 _threshold) returns()
func (_EthereumBridge *EthereumBridgeSession) Initialize(owner_ common.Address, dev_ common.Address, _threshold uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, owner_, dev_, _threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner_, address dev_, uint8 _threshold) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) Initialize(owner_ common.Address, dev_ common.Address, _threshold uint8) (*types.Transaction, error) {
	return _EthereumBridge.Contract.Initialize(&_EthereumBridge.TransactOpts, owner_, dev_, _threshold)
}

// InitializeEthereumBridge is a paid mutator transaction binding the contract method 0x90546099.
//
// Solidity: function initializeEthereumBridge(address owner_, address dev_, uint8 _threshold, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_EthereumBridge *EthereumBridgeTransactor) InitializeEthereumBridge(opts *bind.TransactOpts, owner_ common.Address, dev_ common.Address, _threshold uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "initializeEthereumBridge", owner_, dev_, _threshold, crossChainID, cross, crossInitialSupply)
}

// InitializeEthereumBridge is a paid mutator transaction binding the contract method 0x90546099.
//
// Solidity: function initializeEthereumBridge(address owner_, address dev_, uint8 _threshold, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_EthereumBridge *EthereumBridgeSession) InitializeEthereumBridge(owner_ common.Address, dev_ common.Address, _threshold uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.InitializeEthereumBridge(&_EthereumBridge.TransactOpts, owner_, dev_, _threshold, crossChainID, cross, crossInitialSupply)
}

// InitializeEthereumBridge is a paid mutator transaction binding the contract method 0x90546099.
//
// Solidity: function initializeEthereumBridge(address owner_, address dev_, uint8 _threshold, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) InitializeEthereumBridge(owner_ common.Address, dev_ common.Address, _threshold uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.InitializeEthereumBridge(&_EthereumBridge.TransactOpts, owner_, dev_, _threshold, crossChainID, cross, crossInitialSupply)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ManualReleasePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ManualReleasePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_EthereumBridge *EthereumBridgeSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ManualReleasePendingWithRecipient(&_EthereumBridge.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ManualReleasePendingWithRecipient(&_EthereumBridge.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_EthereumBridge *EthereumBridgeSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeToken(&_EthereumBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_EthereumBridge *EthereumBridgeTransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeToken(&_EthereumBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_EthereumBridge *EthereumBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_EthereumBridge *EthereumBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeTokenBatch(&_EthereumBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _EthereumBridge.Contract.PermitBridgeTokenBatch(&_EthereumBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x85862d87.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, bool supportPermit, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, supportPermit bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, supportPermit, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x85862d87.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, bool supportPermit, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, supportPermit bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RegisterToken(&_EthereumBridge.TransactOpts, remoteChainID, isOrigin, supportPermit, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x85862d87.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, bool supportPermit, address localToken, address remoteToken) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, supportPermit bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RegisterToken(&_EthereumBridge.TransactOpts, remoteChainID, isOrigin, supportPermit, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ReleasePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ReleasePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_EthereumBridge *EthereumBridgeTransactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "releasePendingBatch", remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_EthereumBridge *EthereumBridgeSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ReleasePendingBatch(&_EthereumBridge.TransactOpts, remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.ReleasePendingBatch(&_EthereumBridge.TransactOpts, remoteChainIDs, indexes)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemovePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RemovePending(&_EthereumBridge.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EthereumBridge *EthereumBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RenounceRole(&_EthereumBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RenounceRole(&_EthereumBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RevokeRole(&_EthereumBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RevokeRole(&_EthereumBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RevokeRoleBatch(&_EthereumBridge.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.RevokeRoleBatch(&_EthereumBridge.TransactOpts, roles, accounts)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_EthereumBridge *EthereumBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetBridgeVerifier(&_EthereumBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetBridgeVerifier(&_EthereumBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_EthereumBridge *EthereumBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetChainPause(&_EthereumBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetChainPause(&_EthereumBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_EthereumBridge *EthereumBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetCrossMintableERC20Code(&_EthereumBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetCrossMintableERC20Code(&_EthereumBridge.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_EthereumBridge *EthereumBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetDev(&_EthereumBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetDev(&_EthereumBridge.TransactOpts, dev_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_EthereumBridge *EthereumBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetPause(&_EthereumBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetPause(&_EthereumBridge.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_EthereumBridge *EthereumBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetTokenPause(&_EthereumBridge.TransactOpts, remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetTokenPause(&_EthereumBridge.TransactOpts, remoteChainID, token, pause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetVerificationDelay(&_EthereumBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetVerificationDelay(&_EthereumBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetVerificationDelayExpiration(&_EthereumBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_EthereumBridge *EthereumBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _EthereumBridge.Contract.SetVerificationDelayExpiration(&_EthereumBridge.TransactOpts, remoteChainID, index, delay)
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*EthereumBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeFinalizedIterator{contract: _EthereumBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
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
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*EthereumBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeInitiatedIterator{contract: _EthereumBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeInitiated(log types.Log) (*EthereumBridgeBridgeInitiated, error) {
	event := new(EthereumBridgeBridgeInitiated)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the EthereumBridge contract.
type EthereumBridgeBridgePendingIterator struct {
	Event *EthereumBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgePending)
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
		it.Event = new(EthereumBridgeBridgePending)
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
func (it *EthereumBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgePending represents a BridgePending event raised by the EthereumBridge contract.
type EthereumBridgeBridgePending struct {
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
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*EthereumBridgeBridgePendingIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgePendingIterator{contract: _EthereumBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgePending)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgePending(log types.Log) (*EthereumBridgeBridgePending, error) {
	event := new(EthereumBridgeBridgePending)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the EthereumBridge contract.
type EthereumBridgeBridgeVerifierSetIterator struct {
	Event *EthereumBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeBridgeVerifierSet)
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
		it.Event = new(EthereumBridgeBridgeVerifierSet)
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
func (it *EthereumBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the EthereumBridge contract.
type EthereumBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_EthereumBridge *EthereumBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*EthereumBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeBridgeVerifierSetIterator{contract: _EthereumBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_EthereumBridge *EthereumBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeBridgeVerifierSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*EthereumBridgeBridgeVerifierSet, error) {
	event := new(EthereumBridgeBridgeVerifierSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the EthereumBridge contract.
type EthereumBridgeChainPauseSetIterator struct {
	Event *EthereumBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeChainPauseSet)
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
		it.Event = new(EthereumBridgeChainPauseSet)
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
func (it *EthereumBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeChainPauseSet represents a ChainPauseSet event raised by the EthereumBridge contract.
type EthereumBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_EthereumBridge *EthereumBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*EthereumBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeChainPauseSetIterator{contract: _EthereumBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_EthereumBridge *EthereumBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeChainPauseSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseChainPauseSet(log types.Log) (*EthereumBridgeChainPauseSet, error) {
	event := new(EthereumBridgeChainPauseSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the EthereumBridge contract.
type EthereumBridgeCrossMintableERC20CodeSetIterator struct {
	Event *EthereumBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(EthereumBridgeCrossMintableERC20CodeSet)
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
func (it *EthereumBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the EthereumBridge contract.
type EthereumBridgeCrossMintableERC20CodeSet struct {
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_EthereumBridge *EthereumBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*EthereumBridgeCrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeCrossMintableERC20CodeSetIterator{contract: _EthereumBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_EthereumBridge *EthereumBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeCrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeCrossMintableERC20CodeSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_EthereumBridge *EthereumBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*EthereumBridgeCrossMintableERC20CodeSet, error) {
	event := new(EthereumBridgeCrossMintableERC20CodeSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeDevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the EthereumBridge contract.
type EthereumBridgeDevSetIterator struct {
	Event *EthereumBridgeDevSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeDevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeDevSet)
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
		it.Event = new(EthereumBridgeDevSet)
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
func (it *EthereumBridgeDevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeDevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeDevSet represents a DevSet event raised by the EthereumBridge contract.
type EthereumBridgeDevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_EthereumBridge *EthereumBridgeFilterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*EthereumBridgeDevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeDevSetIterator{contract: _EthereumBridge.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_EthereumBridge *EthereumBridgeFilterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeDevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeDevSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseDevSet(log types.Log) (*EthereumBridgeDevSet, error) {
	event := new(EthereumBridgeDevSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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

// EthereumBridgeManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the EthereumBridge contract.
type EthereumBridgeManualReleasedIterator struct {
	Event *EthereumBridgeManualReleased // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeManualReleased)
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
		it.Event = new(EthereumBridgeManualReleased)
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
func (it *EthereumBridgeManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeManualReleased represents a ManualReleased event raised by the EthereumBridge contract.
type EthereumBridgeManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*EthereumBridgeManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeManualReleasedIterator{contract: _EthereumBridge.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *EthereumBridgeManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeManualReleased)
				if err := _EthereumBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseManualReleased(log types.Log) (*EthereumBridgeManualReleased, error) {
	event := new(EthereumBridgeManualReleased)
	if err := _EthereumBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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

// EthereumBridgePendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the EthereumBridge contract.
type EthereumBridgePendingRemovedIterator struct {
	Event *EthereumBridgePendingRemoved // Event containing the contract specifics and raw log

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
func (it *EthereumBridgePendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgePendingRemoved)
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
		it.Event = new(EthereumBridgePendingRemoved)
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
func (it *EthereumBridgePendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgePendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgePendingRemoved represents a PendingRemoved event raised by the EthereumBridge contract.
type EthereumBridgePendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*EthereumBridgePendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgePendingRemovedIterator{contract: _EthereumBridge.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_EthereumBridge *EthereumBridgeFilterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *EthereumBridgePendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgePendingRemoved)
				if err := _EthereumBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParsePendingRemoved(log types.Log) (*EthereumBridgePendingRemoved, error) {
	event := new(EthereumBridgePendingRemoved)
	if err := _EthereumBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EthereumBridge contract.
type EthereumBridgeRoleAdminChangedIterator struct {
	Event *EthereumBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeRoleAdminChanged)
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
		it.Event = new(EthereumBridgeRoleAdminChanged)
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
func (it *EthereumBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the EthereumBridge contract.
type EthereumBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EthereumBridge *EthereumBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EthereumBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeRoleAdminChangedIterator{contract: _EthereumBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EthereumBridge *EthereumBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EthereumBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeRoleAdminChanged)
				if err := _EthereumBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*EthereumBridgeRoleAdminChanged, error) {
	event := new(EthereumBridgeRoleAdminChanged)
	if err := _EthereumBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EthereumBridge contract.
type EthereumBridgeRoleGrantedIterator struct {
	Event *EthereumBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeRoleGranted)
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
		it.Event = new(EthereumBridgeRoleGranted)
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
func (it *EthereumBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeRoleGranted represents a RoleGranted event raised by the EthereumBridge contract.
type EthereumBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthereumBridge *EthereumBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthereumBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeRoleGrantedIterator{contract: _EthereumBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthereumBridge *EthereumBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EthereumBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeRoleGranted)
				if err := _EthereumBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseRoleGranted(log types.Log) (*EthereumBridgeRoleGranted, error) {
	event := new(EthereumBridgeRoleGranted)
	if err := _EthereumBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EthereumBridge contract.
type EthereumBridgeRoleRevokedIterator struct {
	Event *EthereumBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeRoleRevoked)
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
		it.Event = new(EthereumBridgeRoleRevoked)
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
func (it *EthereumBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeRoleRevoked represents a RoleRevoked event raised by the EthereumBridge contract.
type EthereumBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthereumBridge *EthereumBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthereumBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeRoleRevokedIterator{contract: _EthereumBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EthereumBridge *EthereumBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EthereumBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeRoleRevoked)
				if err := _EthereumBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseRoleRevoked(log types.Log) (*EthereumBridgeRoleRevoked, error) {
	event := new(EthereumBridgeRoleRevoked)
	if err := _EthereumBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	SupportPermit bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0x6509d37a38735e68576c1146ef1653f847a0054869b85d51d170dd430899359d.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin, bool supportPermit)
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

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0x6509d37a38735e68576c1146ef1653f847a0054869b85d51d170dd430899359d.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin, bool supportPermit)
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

// ParseTokenPairRegistered is a log parse operation binding the contract event 0x6509d37a38735e68576c1146ef1653f847a0054869b85d51d170dd430899359d.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin, bool supportPermit)
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*EthereumBridgeTokenPairRegistered, error) {
	event := new(EthereumBridgeTokenPairRegistered)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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

// EthereumBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the EthereumBridge contract.
type EthereumBridgeTokenPauseSetIterator struct {
	Event *EthereumBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeTokenPauseSet)
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
		it.Event = new(EthereumBridgeTokenPauseSet)
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
func (it *EthereumBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeTokenPauseSet represents a TokenPauseSet event raised by the EthereumBridge contract.
type EthereumBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_EthereumBridge *EthereumBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*EthereumBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeTokenPauseSetIterator{contract: _EthereumBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_EthereumBridge *EthereumBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeTokenPauseSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_EthereumBridge *EthereumBridgeFilterer) ParseTokenPauseSet(log types.Log) (*EthereumBridgeTokenPauseSet, error) {
	event := new(EthereumBridgeTokenPauseSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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

// EthereumBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the EthereumBridge contract.
type EthereumBridgeVerificationDelayExpirationSetIterator struct {
	Event *EthereumBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeVerificationDelayExpirationSet)
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
		it.Event = new(EthereumBridgeVerificationDelayExpirationSet)
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
func (it *EthereumBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the EthereumBridge contract.
type EthereumBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_EthereumBridge *EthereumBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*EthereumBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeVerificationDelayExpirationSetIterator{contract: _EthereumBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_EthereumBridge *EthereumBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeVerificationDelayExpirationSet)
				if err := _EthereumBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_EthereumBridge *EthereumBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*EthereumBridgeVerificationDelayExpirationSet, error) {
	event := new(EthereumBridgeVerificationDelayExpirationSet)
	if err := _EthereumBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the EthereumBridge contract.
type EthereumBridgeVerificationDelaySetIterator struct {
	Event *EthereumBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *EthereumBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBridgeVerificationDelaySet)
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
		it.Event = new(EthereumBridgeVerificationDelaySet)
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
func (it *EthereumBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the EthereumBridge contract.
type EthereumBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_EthereumBridge *EthereumBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*EthereumBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _EthereumBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &EthereumBridgeVerificationDelaySetIterator{contract: _EthereumBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_EthereumBridge *EthereumBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *EthereumBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _EthereumBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBridgeVerificationDelaySet)
				if err := _EthereumBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_EthereumBridge *EthereumBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*EthereumBridgeVerificationDelaySet, error) {
	event := new(EthereumBridgeVerificationDelaySet)
	if err := _EthereumBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
