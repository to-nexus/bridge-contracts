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

// BridgeVerifierMetaData contains all meta data concerning the BridgeVerifier contract.
var BridgeVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addValueLimitWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addValueLimitWhitelistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPeriodTotalValueThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCurrentScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMovementHistory\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"history\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getTokenPriceWithValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getValueLimitWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"page\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValueLimitWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValueLimitWhitelistLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerificationAmountThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValueLimitWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeValueLimitWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"removeValueLimitWhitelistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"safePermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"setDefaultTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setExFeeRateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"setPeriodTotalValueThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasPrices\",\"type\":\"uint256[]\"}],\"name\":\"updateGasPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"DefaultTokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"ExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"FinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"GasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"PeriodTotalValueThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"TimeWindowSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValueLimitWhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueLimitWhitelistBypassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValueLimitWhitelistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"VerificationAmountThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeVerifierCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidExFeeRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidTimeWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BridgeVerifierWhitelistAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BridgeVerifierWhitelistNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"75a4463e": "addValueLimitWhitelist(address)",
		"bd574e83": "addValueLimitWhitelistBatch(address[])",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"96ce0795": "denominator()",
		"2fe3b6cf": "dollarDecimals()",
		"d1c01543": "getGasPrice(uint256)",
		"2a4d2f32": "getMinimumTokenValue()",
		"979eb82d": "getPeriodTotalValueThreshold()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"5566ca09": "getTimeWindow()",
		"6332aec6": "getTokenConfig(uint256,address)",
		"9745e4ba": "getTokenCurrentScore(address)",
		"0e4e96bf": "getTokenMovementHistory(address)",
		"d02641a0": "getTokenPrice(address)",
		"ba8d25f8": "getTokenPriceWithValue(address,uint256)",
		"fb1da429": "getValueLimitWhitelist()",
		"1fd4d631": "getValueLimitWhitelist(uint256,uint256)",
		"ce35dd7e": "getValueLimitWhitelistLength()",
		"be716b0a": "getVerificationAmountThreshold()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"b2b49e2e": "grantRoleBatch(bytes32[],address[])",
		"91d14854": "hasRole(bytes32,address)",
		"17cd9cdd": "isValueLimitWhitelisted(address)",
		"741bef1a": "priceFeed()",
		"5b42895f": "removeValueLimitWhitelist(address)",
		"f1f1263d": "removeValueLimitWhitelistBatch(address[])",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"fe81d82e": "safePermit(address,address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"ad1434ca": "setDefaultExFeeRate(uint256)",
		"49c9215e": "setDefaultTokenPrice(uint256)",
		"5ec8f51b": "setExFeeRate(address,uint256)",
		"020a2212": "setExFeeRateBatch(address[],uint256[])",
		"f289d3ba": "setFinalizeBridgeGas(uint256)",
		"9dc696cc": "setMinimumTokenValue(uint256)",
		"ab509f3a": "setPeriodTotalValueThreshold(uint256)",
		"724e78da": "setPriceFeed(address)",
		"8d331996": "setTimeWindow(uint256)",
		"c3de108b": "setVerificationAmountThreshold(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"1f96131e": "updateGasPrice(uint256,uint256)",
		"751ccd79": "updateGasPriceBatch(uint256[],uint256[])",
		"3f4bdec5": "validateBridgeTokenValue(address,uint256)",
		"66ac1eec": "validateBridgeTokenValue(address,uint256,address)",
	},
	Bin: "0x60803461027457601f6131f538819003918201601f19168301916001600160401b038311848410176102785780849261014094604052833981010312610274576100488161028c565b906100556020820161028c565b916100626040830161028c565b926060830151608084015160a08501519160c08601519360e08701519561012061010089015198015198831561023a576001600160a01b03821615610205576001600160a01b038316156101d6576001600160a01b03169182156101a357841561016957612710861161015a57610e108a066101315789158015610140575b15610131576100f26100f8926102a0565b5061030d565b50600180546001600160a01b031916919091179055600255600455600355600555600655600755600855604051612d3b908161049a8239f35b6316a6929960e31b5f5260045ffd5b50610e108a101580156100e1575062093a808a11156100e1565b63aa80145d60e01b5f5260045ffd5b6040516323056b1360e01b815260206004820152601160248201527064656661756c74546f6b656e507269636560781b6044820152606490fd5b6040516323056b1360e01b815260206004820152600a6024820152697072696365466565645f60b01b6044820152606490fd5b6040516323056b1360e01b815260206004820152600660248201526562726964676560d01b6044820152606490fd5b6040516323056b1360e01b815260206004820152600c60248201526b34b734ba34b0b627bbb732b960a11b6044820152606490fd5b6040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361027457565b906102ab825f6103a1565b91826102b45750565b5f8052600d6020526001600160a01b03166102ef817f81955a0a11e65eac625c29e8882660bae4e165a75d72780094acae8ece9a29ee610429565b156102f75750565b637a854c7360e11b5f526004525f60245260445ffd5b90610325825f5160206131d55f395f51905f526103a1565b918261032e5750565b5f5160206131d55f395f51905f525f52600d6020526001600160a01b0316610376817fd06032cb802eef06aa279a4e954daed9c87dafcc32cbe2c5419da429e2c42a5f610429565b1561037e5750565b637a854c7360e11b5f526004525f5160206131d55f395f51905f5260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610423575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b6001810190825f528160205260405f2054155f1461049257805468010000000000000000811015610278576001810180835581101561047e578390825f5260205f20015554915f5260205260405f2055600190565b634e487b7160e01b5f52603260045260245ffd5b5050505f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a7146112b957508063020a2212146111865780630e4e96bf1461106f57806317cd9cdd1461102d5780631f96131e14610fee5780631fd4d63114610fd4578063248a9ca314610fae5780632a4d2f3214610f915780632f2ff15d14610f335780632fe3b6cf14610ea957806336568abe14610e65578063365d71e414610e0b5780633f4bdec514610ddb57806349c9215e14610d4e5780635566ca0914610d315780635b42895f14610ce65780635ec8f51b14610c405780636332aec614610c1257806366ac1eec14610bcd578063724e78da14610b2e578063741bef1a14610b06578063751ccd7914610a3457806375a4463e146109e35780638d3319961461095657806391d148541461090e57806396ce0795146108f25780639745e4ba146108ba578063979eb82d1461089d5780639dc696cc14610851578063a217fddf14610837578063a3246ad3146107d3578063ab509f3a14610787578063ad1434ca14610734578063b2b49e2e1461068f578063ba8d25f814610647578063bd574e8314610595578063be716b0a14610578578063c3de108b1461052c578063ce35dd7e1461050f578063d02641a0146104a6578063d1c015431461047c578063d547741f14610446578063df6e87dc146103ff578063f1f1263d1461037e578063f289d3ba146102f1578063fb1da429146102805763fe81d82e14610221575f80fd5b3461027c5761010036600319011261027c5761023b61135a565b610243611370565b9061024c6113f7565b60a43560ff8116810361027c5761027a93610265612322565b60e4359360c435936084359260643592611ea5565b005b5f80fd5b3461027c575f36600319011261027c57604051600e80548083525f91825260208301915f516020612c865f395f51905f5291905b8181106102db576102d7856102cb8187038261130c565b60405191829182611423565b0390f35b82548452602090930192600192830192016102b4565b3461027c57602036600319011261027c5760043561030d612391565b8015610344576020817fbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b92600255604051908152a1005b6040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606490fd5b3461027c5761038c3661159c565b90610395612391565b5f5b8281106103a057005b6103b36103ae828585611e67565b611e77565b6103bb612391565b6001600160a01b03166103cd81612789565b156103ed57906001915f516020612ce65f395f51905f525f80a201610397565b63a60867e560e01b5f5260045260245ffd5b3461027c57606036600319011261027c576127106104396102d761042c610424611370565b6004356117ab565b604495929491953561177a565b0460405193849384611586565b3461027c57604036600319011261027c5761027a600435610465611370565b906104776104728261171b565b612400565b612438565b3461027c57602036600319011261027c576004355f52600a602052602060405f2054604051908152f35b3461027c57602036600319011261027c576104bf61135a565b6001600160a01b038116906104fb906104e2906104f56104ed6104e7848761248c565b611769565b92612537565b93909561248c565b916125ac565b604080519215158352602083019190915290f35b3461027c575f36600319011261027c576020600e54604051908152f35b3461027c57602036600319011261027c577ff905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2602060043561056b612391565b80600655604051908152a1005b3461027c575f36600319011261027c576020600654604051908152f35b3461027c576105a33661159c565b906105ac612391565b5f5b8281106105b757005b6105c56103ae828585611e67565b6105cd612391565b6001600160a01b03168015610617576105e581612a13565b1561060557906001915f516020612ca65f395f51905f525f80a2016105ae565b635bf16f2960e11b5f5260045260245ffd5b6040516323056b1360e01b81526020600482015260076024820152661858d8dbdd5b9d60ca1b6044820152606490fd5b3461027c57604036600319011261027c5761066061135a565b6104fb6104e261068561067284612537565b90949092906001600160a01b031661248c565b90602435906125ac565b3461027c5761069d36611465565b908051825103610725575f5b825181101561027a576106bc81836115ee565b516001600160a01b036106cf83866115ee565b51166106dd6104728361171b565b6106e78183612681565b6106f6575b50506001016106a9565b815f52600d60205261070b8160405f20612a68565b6106ec575b637a854c7360e11b5f5260045260245260445ffd5b6374ddb37160e01b5f5260045ffd5b3461027c57602036600319011261027c57600435610750612244565b6127108111610778575f516020612cc65f395f51905f526020825f93600355604051908152a2005b63aa80145d60e01b5f5260045ffd5b3461027c57602036600319011261027c577f7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e260206004356107c6612391565b80600755604051908152a1005b3461027c57602036600319011261027c576004355f52600d60205260405f206040519081602082549182815201915f5260205f20905f5b818110610821576102d7856102cb8187038261130c565b825484526020909301926001928301920161080a565b3461027c575f36600319011261027c5760206040515f8152f35b3461027c57602036600319011261027c577f0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e96020600435610890612244565b80600555604051908152a1005b3461027c575f36600319011261027c576020600754604051908152f35b3461027c57602036600319011261027c576001600160a01b036108db61135a565b165f52600b602052602060405f2054604051908152f35b3461027c575f36600319011261027c5760206040516127108152f35b3461027c57604036600319011261027c57610927611370565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461027c57602036600319011261027c57600435610972612391565b610e1081066109ba57801580156109c9575b156109ba576020817f52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a492600855604051908152a1005b6316a6929960e31b5f5260045ffd5b50610e108110158015610984575062093a80811115610984565b3461027c57602036600319011261027c576109fc61135a565b610a04612391565b6001600160a01b0316801561061757610a1c81612a13565b15610605575f516020612ca65f395f51905f525f80a2005b3461027c57604036600319011261027c576004356001600160401b03811161027c57610a6490369060040161139a565b6024356001600160401b03811161027c57610a8390369060040161139a565b610a8b6122b3565b8151815103610af7575f5b825181101561027a5780610aac600192846115ee565b51610ab782866115ee565b515f52600a60205260405f2055610ace81856115ee565b515f516020612c665f395f51905f526020610ae984876115ee565b51604051908152a201610a96565b63322473c360e01b5f5260045ffd5b3461027c575f36600319011261027c576001546040516001600160a01b039091168152602090f35b3461027c57602036600319011261027c576004356001600160a01b0381169081900361027c57610b5c612391565b8015610b9a57600180546001600160a01b031916821790557fe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e65f80a2005b6040516323056b1360e01b815260206004820152600a60248201526917dc1c9a58d95199595960b21b6044820152606490fd5b3461027c57606036600319011261027c576102d7610c06610bec61135a565b610bf46113f7565b90610bfd612322565b60243590611c07565b6040519182918261155f565b3461027c57604036600319011261027c576102d7610c31610424611370565b60409391935193849384611586565b3461027c57604036600319011261027c57610c5961135a565b60243590610c65612244565b6001600160a01b0316908115610cb8576127108111801590610cae575b156107785760205f516020612cc65f395f51905f5291835f52600982528060405f2055604051908152a2005b505f198114610c82565b6040516323056b1360e01b81526020600482015260056024820152643a37b5b2b760d91b6044820152606490fd5b3461027c57602036600319011261027c57610cff61135a565b610d07612391565b6001600160a01b0316610d1981612789565b156103ed575f516020612ce65f395f51905f525f80a2005b3461027c575f36600319011261027c576020600854604051908152f35b3461027c57602036600319011261027c57600435610d6a612244565b8015610da1576020817f1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a92600455604051908152a1005b6040516323056b1360e01b815260206004820152601160248201527064656661756c74546f6b656e507269636560781b6044820152606490fd5b3461027c57604036600319011261027c576102d7610c06610dfa61135a565b610e02612322565b602435906118a0565b3461027c57610e1936611465565b908051825103610725575f5b825181101561027a5780610e5e610e3e600193856115ee565b51838060a01b03610e4f84886115ee565b5116906104776104728261171b565b5001610e25565b3461027c57604036600319011261027c57610e7e611370565b336001600160a01b03821603610e9a5761027a90600435612438565b63334bd91960e11b5f5260045ffd5b3461027c575f36600319011261027c57600154604051632fe3b6cf60e01b815290602090829060049082906001600160a01b03165afa8015610f28576020915f91610efb575b5060ff60405191168152f35b610f1b9150823d8411610f21575b610f13818361130c565b81019061172c565b82610eef565b503d610f09565b6040513d5f823e3d90fd5b3461027c57604036600319011261027c57600435610f4f611370565b610f5b6104728361171b565b610f658183612681565b610f6b57005b815f52600d602052610f8a60405f209160018060a01b03168092612a68565b1561071057005b3461027c575f36600319011261027c576020600554604051908152f35b3461027c57602036600319011261027c576020610fcc60043561171b565b604051908152f35b3461027c576102d76102cb610fe83661140d565b90611659565b3461027c5760205f516020612c665f395f51905f5261100c3661140d565b6110179391936122b3565b835f52600a82528060405f2055604051908152a2005b3461027c57602036600319011261027c5760206110656001600160a01b0361105361135a565b165f52600f60205260405f2054151590565b6040519015158152f35b3461027c57602036600319011261027c576001600160a01b0361109061135a565b16805f52600c6020526110a560405f20612215565b906110af82611343565b916110bd604051938461130c565b8083526110c981611343565b602084019290601f19013684375f5b828110611123578385604051918291602083019060208452518091526040830191905f5b81811061110a575050500390f35b82518452859450602093840193909201916001016110fc565b815f52600c60205260405f209061113982612215565b8110156111745781546001600160801b0390811682821601165f9081526001928301602052604090205461116d82886115ee565b52016110d8565b634e487b715f5260326020526024601cfd5b3461027c57604036600319011261027c576004356001600160401b03811161027c573660238201121561027c578060040135906111c282611343565b916111d0604051938461130c565b8083526024602084019160051b8301019136831161027c57602401905b8282106112a157836024356001600160401b03811161027c5761121490369060040161139a565b908051825103610af7575f5b815181101561027a576001600160a01b0361123b82846115ee565b51169061124881856115ee565b5191611252612244565b8015610cb8576127108311801590611297575b15610778575f516020612cc65f395f51905f526020600194835f52600982528060405f2055604051908152a201611220565b505f198314611265565b602080916112ae84611386565b8152019101906111ed565b3461027c57602036600319011261027c576004359063ffffffff60e01b821680920361027c57602091637965db0b60e01b81149081156112fb575b5015158152f35b6301ffc9a760e01b149050836112f4565b601f909101601f19168101906001600160401b0382119082101761132f57604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161132f5760051b60200190565b600435906001600160a01b038216820361027c57565b602435906001600160a01b038216820361027c57565b35906001600160a01b038216820361027c57565b9080601f8301121561027c5781356113b181611343565b926113bf604051948561130c565b81845260208085019260051b82010192831161027c57602001905b8282106113e75750505090565b81358152602091820191016113da565b604435906001600160a01b038216820361027c57565b604090600319011261027c576004359060243590565b60206040818301928281528451809452019201905f5b8181106114465750505090565b82516001600160a01b0316845260209384019390920191600101611439565b90604060031983011261027c576004356001600160401b03811161027c578260238201121561027c5780600401359061149d82611343565b916114ab604051938461130c565b8083526024602084019160051b8301019185831161027c57602401905b82821061154f57509193602435925090506001600160401b03821161027c578060238301121561027c57816004013561150081611343565b9261150e604051948561130c565b8184526024602085019260051b82010192831161027c57602401905b8282106115375750505090565b6020809161154484611386565b81520191019061152a565b81358152602091820191016114c8565b91906020830192600c8210156115725752565b634e487b7160e01b5f52602160045260245ffd5b6040919493926060820195825260208201520152565b90602060031983011261027c576004356001600160401b03811161027c578260238201121561027c576004810135926001600160401b03841161027c5760248460051b8301011161027c576024019190565b80518210156116025760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b6001600160c01b039081039190821161162b57565b634e487b7160e01b5f52601160045260245ffd5b9190820391821161162b57565b9190820180921161162b57565b9190600e54808410801590611713575b61170b57836116779161163f565b9081808210911802189161168a83611343565b92611698604051948561130c565b808452601f196116a782611343565b01366020860137835f92600e54935b8381106116c4575050505050565b6116ce818361164c565b908582101561160257600e5f525f516020612c865f395f51905f5290910154600191906001600160a01b031661170482866115ee565b52016116b6565b506060925050565b508115611669565b5f525f602052600160405f20015490565b9081602091031261027c575160ff8116810361027c5790565b9081606091031261027c578051801515810361027c57916040602083015192015190565b60ff16604d811161162b57600a0a90565b8181029291811591840414171561162b57565b8115611797570490565b634e487b7160e01b5f52601260045260245ffd5b60015460405163e588566f60e01b81526001600160a01b038085166004830181905295949260609183916024918391165afa8015610f28575f915f9161186e575b509015611865575b6118006104e28661248c565b908061184b5750935b5f52600960205260405f2054905f1982145f146118305761182c91505f926124e8565b5091565b8115611841575b61182c91926124e8565b6003549150611837565b61185a61185f9260055461177a565b61178d565b93611809565b506004546117f4565b9050611892915060603d606011611899575b61188a818361130c565b810190611745565b505f6117ec565b503d611880565b91906118ab83612537565b600654948515801593508390611bf0575b82159081611be8575b50611bde576001600160a01b0316926118e591906104f56104e28661248c565b936001600160c01b038511611bd557849082611bcb575b5050611bc457600854158015611bba575b611bb35761191a83611616565b815f52600b60205260405f205411611bac57805f52600b60205260405f2061194384825461164c565b9055805f52600c60205260405f2092610e10420492610e10840293808504610e10149015171561162b576119796008548561163f565b9160018601945b61198987612215565b15611b9e578654608081901c6001600160801b03909116146111745786546001600160801b03165f9081526020879052604090205460c081901c851115611a70575f868152600b6020526040902054906001600160c01b031680821115611a68576119f39161163f565b5f868152600b602052604090205586546001600160801b0381169060801c8114611a5657611a21818861222f565b506001600160801b038181165f9081526020899052604081205588546001600160801b03191660019290920116178755611980565b634e487b715f5260316020526024601cfd5b50505f6119f3565b509092509490949392935b611a8482612215565b151580611b5d575b611b01575b8154608081901c946001600160801b036001870181169592168514611aef57611ad19560018060801b03165f5260205260405f209160c01b179055612649565b5f52600b60205260405f205460075410611aea57600190565b600490565b634e487b715f5260416020526024601cfd5b81549092608082901c916001600160801b03168214611a56575f199091016001600160801b03165f8181526020869052604081208054919055611b579291611b499085612649565b6001600160c01b031661164c565b91611a91565b508154608081901c6001600160801b039091161461117457815460801c5f19016001600160801b03165f9081526020859052604090205460c01c8114611a8c565b909250949094939293611a7b565b5060089150565b5060019150565b506007541561190d565b5060039150565b109050835f6118fc565b50600793505050565b50600a9450505050565b90505f6118c5565b50600854151580156118bc575060075415156118bc565b92916001600160a01b031680151580611e4c575b611e0f5750611c2983612537565b600654948515801593508390611df8575b82159081611df0575b50611bde576001600160a01b031692611c6391906104f56104e28661248c565b936001600160c01b038511611bd557849082611de6575b5050611bc457600854158015611ddc575b611bb357611c9883611616565b815f52600b60205260405f205411611bac57805f52600b60205260405f20611cc184825461164c565b9055805f52600c60205260405f2092610e10420492610e10840293808504610e10149015171561162b57611cf76008548561163f565b9160018601945b611d0787612215565b15611b9e578654608081901c6001600160801b03909116146111745786546001600160801b03165f9081526020879052604090205460c081901c851115611a70575f868152600b6020526040902054906001600160c01b031680821115611dd457611d719161163f565b5f868152600b602052604090205586546001600160801b0381169060801c8114611a5657611d9f818861222f565b506001600160801b038181165f9081526020899052604081205588546001600160801b03191660019290920116178755611cfe565b50505f611d71565b5060075415611c8b565b109050835f611c7a565b90505f611c43565b5060085415158015611c3a57506007541515611c3a565b60405191825291926001600160a01b0316907fdfe196174d835a3614fc87874a6a97dde9719b32d7d503610251887094dccfce90602090a3600190565b50611e62815f52600f60205260405f2054151590565b611c1b565b91908110156116025760051b0190565b356001600160a01b038116810361027c5790565b6001600160a01b0391821681529116602082015260400190565b604051636eb1769f60e11b81529791969195919491936001600160a01b03909116929160208980611eda868c60048401611e8b565b0381875afa8015610f28575f995f916121e3575b50828110156121d757843b1561027c5760405163d505accf60e01b81526001600160a01b038a81166004830181905290861660248301819052604483018690526064830185905260ff8916608484015260a483018a905260c483018b9052909a9095915f8160e481838c5af18015610f28576121ba575b50604051636eb1769f60e11b81529160209183918291611f89919060048401611e8b565b0381895afa9081156121af578b9161217d575b5083810361216e5714611fb5575b505050505050505050565b80421161215f57604051623f675f60e91b815260048101899052602081602481885afa908115612154578a91612122575b5080156121135791600494939160209360405192858401947f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c986528c6040860152606085015260808401525f190160a083015260c082015260c0815261204d60e08261130c565b5190209160405193848092633644e51560e01b82525afa9182156121085787926120cf575b5091604261209c94926120a596946040519161190160f01b83526002830152602282015220612ab5565b90929192612b2d565b6001600160a01b0316036120c0578080808080808080611faa565b6313f3ff5560e31b8152600490fd5b9091506020813d602011612100575b816120eb6020938361130c565b810103126120fc5751906042612072565b8680fd5b3d91506120de565b6040513d89823e3d90fd5b6304de863b60e21b8a5260048afd5b90506020813d60201161214c575b8161213d6020938361130c565b8101031261027c57515f611fe6565b3d9150612130565b6040513d8c823e3d90fd5b63219386ff60e11b8952600489fd5b6304de863b60e21b8b5260048bfd5b90506020813d6020116121a7575b816121986020938361130c565b8101031261027c57515f611f9c565b3d915061218b565b6040513d8d823e3d90fd5b6020919c50916121cd5f611f899461130c565b5f9c915091611f65565b50505050505050505050565b90506020813d60201161220d575b816121fe6020938361130c565b8101031261027c57515f611eee565b3d91506121f1565b546001600160801b0380821660809290921c919091031690565b9060018060801b03165f5260205260405f2090565b335f9081527fa74d116c0406588b87fcebe22b8d03a12bf1aa08ea523c2004e00dd791abafe3602052604090205460ff161561227c57565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b335f9081527fe70f4a856a97bd1ebdcb74b506f0a385e96f16da6133e1d2983a79b65bc91dc0602052604090205460ff16156122eb57565b63e2517d3f60e01b5f52336004527fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a60245260445ffd5b335f9081527f7dcd276b72586271b4623598c6b3ecb5f3594881d1da5ce0d28e8f8adbf67787602052604090205460ff161561235a57565b63e2517d3f60e01b5f52336004527f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f60245260445ffd5b335f9081527f7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ec602052604090205460ff16156123c957565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b5f8181526020818152604080832033845290915290205460ff16156124225750565b63e2517d3f60e01b5f523360045260245260445ffd5b9190916124458382612709565b928361244f575050565b815f52600d60205261246e60405f209160018060a01b0316809261285e565b15612477575050565b639d98e84160e01b5f5260045260245260445ffd5b6001600160a01b0316600181036124a35750601290565b60206004916040519283809263313ce56760e01b82525afa908115610f28575f916124cc575090565b6124e5915060203d602011610f2157610f13818361130c565b90565b90815f52600a60205260405f2054801561252e576125279261251760018060a01b03600154169260025461177a565b926001600160a01b03169161290f565b9091509091565b5050505f905f90565b60015460405163e588566f60e01b81526001600160a01b039283166004820152929160609184916024918391165afa8015610f28575f925f91612586575b5090821561257f57565b6004549150565b90506125a291925060603d6060116118995761188a818361130c565b509190915f612575565b91818302915f198185099383808610950394808603951461263c57848311156126245790829109815f0382168092046002816003021880820260020302808202600203028082026002030280820260020302808202600203028091026002030293600183805f03040190848311900302920304170290565b82634e487b715f52156003026011186020526024601cfd5b5050906124e5925061178d565b80546001600160801b031660809290921b6001600160801b031916919091179055565b8054821015611602575f5260205f2001905f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16612703575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615612703575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b5f818152600f60205260409020548015612703575f19810181811161162b57600e545f1981019190821161162b57818103612810575b505050600e5480156127fc575f19016127d981600e61266c565b8154905f199060031b1b19169055600e555f52600f6020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b61284861282161283293600e61266c565b90549060031b1c928392600e61266c565b819391549060031b91821b915f19901b19161790565b90555f52600f60205260405f20555f80806127bf565b906001820191815f528260205260405f20548015155f14612907575f19810181811161162b5782545f1981019190821161162b578181036128d2575b505050805480156127fc575f1901906128b3828261266c565b8154905f199060031b1b19169055555f526020525f6040812055600190565b6128f26128e2612832938661266c565b90549060031b1c9283928661266c565b90555f528360205260405f20555f808061289a565b505050505f90565b604051630ebec55160e21b81526004810192909252909392906001600160a01b0316606085602481845afa948515610f28575f905f966129ee575b50156129e15760405163e588566f60e01b81526001600160a01b038416600482015290606090829060249082905afa948515610f28575f955f925f916129ba575b509386156129ab57906129a26129a894939261248c565b92612ba1565b91565b505050505090505f905f905f90565b919650506129d7915060603d6060116118995761188a818361130c565b919590915f61298b565b50505090505f905f905f90565b9050612a0a91955060603d6060116118995761188a818361130c565b9591909561294a565b805f52600f60205260405f2054155f14612a6357600e54600160401b81101561132f57612a4c612832826001859401600e55600e61266c565b9055600e54905f52600f60205260405f2055600190565b505f90565b5f82815260018201602052604090205461270357805490600160401b82101561132f5782612aa061283284600180960185558461266c565b90558054925f520160205260405f2055600190565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411612b22579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610f28575f516001600160a01b03811615612b1857905f905f90565b505f906001905f90565b5050505f9160039190565b60048110156115725780612b3f575050565b60018103612b565763f645eedf60e01b5f5260045ffd5b60028103612b71575063fce698f760e01b5f5260045260245ffd5b600314612b7b5750565b6335e2f38360e21b5f5260045260245ffd5b9060ff8091169116039060ff821161162b57565b91928115612c36578315612c075760ff811680601214612bfb5760121015612be6576124e5939291612bda6104e26012612be094612b8d565b9061177a565b906125ac565b6124e593612bda6104e26104f5936012612b8d565b50506124e592916125ac565b6040516349671cd160e11b8152602060048201526006602482015265383934b1b2a160d11b6044820152606490fd5b6040516349671cd160e11b815260206004820152600660248201526570726963654160d01b6044820152606490fdfe4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3bb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fde1d1777e4c7896f6e13f47c189990ed90f2210cb88f3836cd7bd0cb1f820945b386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1dcf6ae774f84d549d12db18368181005de6e396adc8740a42b1df7a36352edad3a26469706673582212208965a39db8e1e27e172d9ced79e537596bf190382e9edf491b1a9034dd4362ed64736f6c634300081c003352ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f",
}

// BridgeVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeVerifierMetaData.ABI instead.
var BridgeVerifierABI = BridgeVerifierMetaData.ABI

// Deprecated: Use BridgeVerifierMetaData.Sigs instead.
// BridgeVerifierFuncSigs maps the 4-byte function signature to its string representation.
var BridgeVerifierFuncSigs = BridgeVerifierMetaData.Sigs

// BridgeVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeVerifierMetaData.Bin instead.
var BridgeVerifierBin = BridgeVerifierMetaData.Bin

// DeployBridgeVerifier deploys a new Ethereum contract, binding an instance of BridgeVerifier to it.
func DeployBridgeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, bridge common.Address, priceFeed_ common.Address, finalizeBridgeGas *big.Int, defaultTokenPrice *big.Int, defaultExFeeRate *big.Int, minimumTokenValue *big.Int, verificationAmountThreshold *big.Int, periodTotalValueThreshold *big.Int, timeWindow *big.Int) (common.Address, *types.Transaction, *BridgeVerifier, error) {
	parsed, err := BridgeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeVerifierBin), backend, initialOwner, bridge, priceFeed_, finalizeBridgeGas, defaultTokenPrice, defaultExFeeRate, minimumTokenValue, verificationAmountThreshold, periodTotalValueThreshold, timeWindow)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeVerifier{BridgeVerifierCaller: BridgeVerifierCaller{contract: contract}, BridgeVerifierTransactor: BridgeVerifierTransactor{contract: contract}, BridgeVerifierFilterer: BridgeVerifierFilterer{contract: contract}}, nil
}

// BridgeVerifier is an auto generated Go binding around an Ethereum contract.
type BridgeVerifier struct {
	BridgeVerifierCaller     // Read-only binding to the contract
	BridgeVerifierTransactor // Write-only binding to the contract
	BridgeVerifierFilterer   // Log filterer for contract events
}

// BridgeVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeVerifierSession struct {
	Contract     *BridgeVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeVerifierCallerSession struct {
	Contract *BridgeVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeVerifierTransactorSession struct {
	Contract     *BridgeVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeVerifierRaw struct {
	Contract *BridgeVerifier // Generic contract binding to access the raw methods on
}

// BridgeVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeVerifierCallerRaw struct {
	Contract *BridgeVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeVerifierTransactorRaw struct {
	Contract *BridgeVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeVerifier creates a new instance of BridgeVerifier, bound to a specific deployed contract.
func NewBridgeVerifier(address common.Address, backend bind.ContractBackend) (*BridgeVerifier, error) {
	contract, err := bindBridgeVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifier{BridgeVerifierCaller: BridgeVerifierCaller{contract: contract}, BridgeVerifierTransactor: BridgeVerifierTransactor{contract: contract}, BridgeVerifierFilterer: BridgeVerifierFilterer{contract: contract}}, nil
}

// NewBridgeVerifierCaller creates a new read-only instance of BridgeVerifier, bound to a specific deployed contract.
func NewBridgeVerifierCaller(address common.Address, caller bind.ContractCaller) (*BridgeVerifierCaller, error) {
	contract, err := bindBridgeVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierCaller{contract: contract}, nil
}

// NewBridgeVerifierTransactor creates a new write-only instance of BridgeVerifier, bound to a specific deployed contract.
func NewBridgeVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeVerifierTransactor, error) {
	contract, err := bindBridgeVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierTransactor{contract: contract}, nil
}

// NewBridgeVerifierFilterer creates a new log filterer instance of BridgeVerifier, bound to a specific deployed contract.
func NewBridgeVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeVerifierFilterer, error) {
	contract, err := bindBridgeVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierFilterer{contract: contract}, nil
}

// bindBridgeVerifier binds a generic wrapper to an already deployed contract.
func bindBridgeVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeVerifier *BridgeVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeVerifier.Contract.BridgeVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeVerifier *BridgeVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.BridgeVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeVerifier *BridgeVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.BridgeVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeVerifier *BridgeVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeVerifier *BridgeVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeVerifier *BridgeVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeVerifier.Contract.DEFAULTADMINROLE(&_BridgeVerifier.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeVerifier.Contract.DEFAULTADMINROLE(&_BridgeVerifier.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumValue *big.Int
		NetworkFee   *big.Int
		ExFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeVerifier.Contract.CalculateFee(&_BridgeVerifier.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeVerifier.Contract.CalculateFee(&_BridgeVerifier.CallOpts, remoteChainID, token, value)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) Denominator() (*big.Int, error) {
	return _BridgeVerifier.Contract.Denominator(&_BridgeVerifier.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) Denominator() (*big.Int, error) {
	return _BridgeVerifier.Contract.Denominator(&_BridgeVerifier.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) DollarDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "dollarDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) DollarDecimals() (*big.Int, error) {
	return _BridgeVerifier.Contract.DollarDecimals(&_BridgeVerifier.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) DollarDecimals() (*big.Int, error) {
	return _BridgeVerifier.Contract.DollarDecimals(&_BridgeVerifier.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetGasPrice(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getGasPrice", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetGasPrice(&_BridgeVerifier.CallOpts, remoteChainID)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetGasPrice(&_BridgeVerifier.CallOpts, remoteChainID)
}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetMinimumTokenValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getMinimumTokenValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetMinimumTokenValue() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetMinimumTokenValue(&_BridgeVerifier.CallOpts)
}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetMinimumTokenValue() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetMinimumTokenValue(&_BridgeVerifier.CallOpts)
}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetPeriodTotalValueThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getPeriodTotalValueThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetPeriodTotalValueThreshold() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetPeriodTotalValueThreshold(&_BridgeVerifier.CallOpts)
}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetPeriodTotalValueThreshold() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetPeriodTotalValueThreshold(&_BridgeVerifier.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeVerifier.Contract.GetRoleAdmin(&_BridgeVerifier.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeVerifier.Contract.GetRoleAdmin(&_BridgeVerifier.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BridgeVerifier *BridgeVerifierCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BridgeVerifier *BridgeVerifierSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetRoleMembers(&_BridgeVerifier.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BridgeVerifier *BridgeVerifierCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetRoleMembers(&_BridgeVerifier.CallOpts, role)
}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetTimeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTimeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetTimeWindow() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetTimeWindow(&_BridgeVerifier.CallOpts)
}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTimeWindow() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetTimeWindow(&_BridgeVerifier.CallOpts)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenConfig(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFeeRate    *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenConfig", remoteChainID, token)

	outstruct := new(struct {
		MinimumValue *big.Int
		NetworkFee   *big.Int
		ExFeeRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFeeRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenConfig(&_BridgeVerifier.CallOpts, remoteChainID, token)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 networkFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	NetworkFee   *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenConfig(&_BridgeVerifier.CallOpts, remoteChainID, token)
}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenCurrentScore(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenCurrentScore", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenCurrentScore(token common.Address) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetTokenCurrentScore(&_BridgeVerifier.CallOpts, token)
}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenCurrentScore(token common.Address) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetTokenCurrentScore(&_BridgeVerifier.CallOpts, token)
}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenMovementHistory(opts *bind.CallOpts, token common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenMovementHistory", token)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenMovementHistory(token common.Address) ([][32]byte, error) {
	return _BridgeVerifier.Contract.GetTokenMovementHistory(&_BridgeVerifier.CallOpts, token)
}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenMovementHistory(token common.Address) ([][32]byte, error) {
	return _BridgeVerifier.Contract.GetTokenMovementHistory(&_BridgeVerifier.CallOpts, token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address token) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenPrice", token)

	outstruct := new(struct {
		Exist bool
		Price *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address token) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenPrice(token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenPrice(&_BridgeVerifier.CallOpts, token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address token) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenPrice(token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenPrice(&_BridgeVerifier.CallOpts, token)
}

// GetTokenPriceWithValue is a free data retrieval call binding the contract method 0xba8d25f8.
//
// Solidity: function getTokenPriceWithValue(address token, uint256 value) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenPriceWithValue(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenPriceWithValue", token, value)

	outstruct := new(struct {
		Exist bool
		Price *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenPriceWithValue is a free data retrieval call binding the contract method 0xba8d25f8.
//
// Solidity: function getTokenPriceWithValue(address token, uint256 value) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenPriceWithValue(token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenPriceWithValue(&_BridgeVerifier.CallOpts, token, value)
}

// GetTokenPriceWithValue is a free data retrieval call binding the contract method 0xba8d25f8.
//
// Solidity: function getTokenPriceWithValue(address token, uint256 value) view returns(bool exist, uint256 price)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenPriceWithValue(token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenPriceWithValue(&_BridgeVerifier.CallOpts, token, value)
}

// GetValueLimitWhitelist is a free data retrieval call binding the contract method 0x1fd4d631.
//
// Solidity: function getValueLimitWhitelist(uint256 offset, uint256 limit) view returns(address[] page)
func (_BridgeVerifier *BridgeVerifierCaller) GetValueLimitWhitelist(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getValueLimitWhitelist", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValueLimitWhitelist is a free data retrieval call binding the contract method 0x1fd4d631.
//
// Solidity: function getValueLimitWhitelist(uint256 offset, uint256 limit) view returns(address[] page)
func (_BridgeVerifier *BridgeVerifierSession) GetValueLimitWhitelist(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelist(&_BridgeVerifier.CallOpts, offset, limit)
}

// GetValueLimitWhitelist is a free data retrieval call binding the contract method 0x1fd4d631.
//
// Solidity: function getValueLimitWhitelist(uint256 offset, uint256 limit) view returns(address[] page)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetValueLimitWhitelist(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelist(&_BridgeVerifier.CallOpts, offset, limit)
}

// GetValueLimitWhitelist0 is a free data retrieval call binding the contract method 0xfb1da429.
//
// Solidity: function getValueLimitWhitelist() view returns(address[])
func (_BridgeVerifier *BridgeVerifierCaller) GetValueLimitWhitelist0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getValueLimitWhitelist0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValueLimitWhitelist0 is a free data retrieval call binding the contract method 0xfb1da429.
//
// Solidity: function getValueLimitWhitelist() view returns(address[])
func (_BridgeVerifier *BridgeVerifierSession) GetValueLimitWhitelist0() ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelist0(&_BridgeVerifier.CallOpts)
}

// GetValueLimitWhitelist0 is a free data retrieval call binding the contract method 0xfb1da429.
//
// Solidity: function getValueLimitWhitelist() view returns(address[])
func (_BridgeVerifier *BridgeVerifierCallerSession) GetValueLimitWhitelist0() ([]common.Address, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelist0(&_BridgeVerifier.CallOpts)
}

// GetValueLimitWhitelistLength is a free data retrieval call binding the contract method 0xce35dd7e.
//
// Solidity: function getValueLimitWhitelistLength() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetValueLimitWhitelistLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getValueLimitWhitelistLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueLimitWhitelistLength is a free data retrieval call binding the contract method 0xce35dd7e.
//
// Solidity: function getValueLimitWhitelistLength() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetValueLimitWhitelistLength() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelistLength(&_BridgeVerifier.CallOpts)
}

// GetValueLimitWhitelistLength is a free data retrieval call binding the contract method 0xce35dd7e.
//
// Solidity: function getValueLimitWhitelistLength() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetValueLimitWhitelistLength() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetValueLimitWhitelistLength(&_BridgeVerifier.CallOpts)
}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetVerificationAmountThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getVerificationAmountThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetVerificationAmountThreshold() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetVerificationAmountThreshold(&_BridgeVerifier.CallOpts)
}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetVerificationAmountThreshold() (*big.Int, error) {
	return _BridgeVerifier.Contract.GetVerificationAmountThreshold(&_BridgeVerifier.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeVerifier.Contract.HasRole(&_BridgeVerifier.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeVerifier.Contract.HasRole(&_BridgeVerifier.CallOpts, role, account)
}

// IsValueLimitWhitelisted is a free data retrieval call binding the contract method 0x17cd9cdd.
//
// Solidity: function isValueLimitWhitelisted(address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCaller) IsValueLimitWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "isValueLimitWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValueLimitWhitelisted is a free data retrieval call binding the contract method 0x17cd9cdd.
//
// Solidity: function isValueLimitWhitelisted(address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierSession) IsValueLimitWhitelisted(account common.Address) (bool, error) {
	return _BridgeVerifier.Contract.IsValueLimitWhitelisted(&_BridgeVerifier.CallOpts, account)
}

// IsValueLimitWhitelisted is a free data retrieval call binding the contract method 0x17cd9cdd.
//
// Solidity: function isValueLimitWhitelisted(address account) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCallerSession) IsValueLimitWhitelisted(account common.Address) (bool, error) {
	return _BridgeVerifier.Contract.IsValueLimitWhitelisted(&_BridgeVerifier.CallOpts, account)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeVerifier *BridgeVerifierCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeVerifier *BridgeVerifierSession) PriceFeed() (common.Address, error) {
	return _BridgeVerifier.Contract.PriceFeed(&_BridgeVerifier.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeVerifier *BridgeVerifierCallerSession) PriceFeed() (common.Address, error) {
	return _BridgeVerifier.Contract.PriceFeed(&_BridgeVerifier.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeVerifier *BridgeVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeVerifier.Contract.SupportsInterface(&_BridgeVerifier.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeVerifier *BridgeVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeVerifier.Contract.SupportsInterface(&_BridgeVerifier.CallOpts, interfaceId)
}

// AddValueLimitWhitelist is a paid mutator transaction binding the contract method 0x75a4463e.
//
// Solidity: function addValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) AddValueLimitWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "addValueLimitWhitelist", account)
}

// AddValueLimitWhitelist is a paid mutator transaction binding the contract method 0x75a4463e.
//
// Solidity: function addValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierSession) AddValueLimitWhitelist(account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.AddValueLimitWhitelist(&_BridgeVerifier.TransactOpts, account)
}

// AddValueLimitWhitelist is a paid mutator transaction binding the contract method 0x75a4463e.
//
// Solidity: function addValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) AddValueLimitWhitelist(account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.AddValueLimitWhitelist(&_BridgeVerifier.TransactOpts, account)
}

// AddValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xbd574e83.
//
// Solidity: function addValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) AddValueLimitWhitelistBatch(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "addValueLimitWhitelistBatch", accounts)
}

// AddValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xbd574e83.
//
// Solidity: function addValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierSession) AddValueLimitWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.AddValueLimitWhitelistBatch(&_BridgeVerifier.TransactOpts, accounts)
}

// AddValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xbd574e83.
//
// Solidity: function addValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) AddValueLimitWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.AddValueLimitWhitelistBatch(&_BridgeVerifier.TransactOpts, accounts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.GrantRole(&_BridgeVerifier.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.GrantRole(&_BridgeVerifier.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.GrantRoleBatch(&_BridgeVerifier.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.GrantRoleBatch(&_BridgeVerifier.TransactOpts, roles, accounts)
}

// RemoveValueLimitWhitelist is a paid mutator transaction binding the contract method 0x5b42895f.
//
// Solidity: function removeValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RemoveValueLimitWhitelist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "removeValueLimitWhitelist", account)
}

// RemoveValueLimitWhitelist is a paid mutator transaction binding the contract method 0x5b42895f.
//
// Solidity: function removeValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierSession) RemoveValueLimitWhitelist(account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemoveValueLimitWhitelist(&_BridgeVerifier.TransactOpts, account)
}

// RemoveValueLimitWhitelist is a paid mutator transaction binding the contract method 0x5b42895f.
//
// Solidity: function removeValueLimitWhitelist(address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RemoveValueLimitWhitelist(account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemoveValueLimitWhitelist(&_BridgeVerifier.TransactOpts, account)
}

// RemoveValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xf1f1263d.
//
// Solidity: function removeValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RemoveValueLimitWhitelistBatch(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "removeValueLimitWhitelistBatch", accounts)
}

// RemoveValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xf1f1263d.
//
// Solidity: function removeValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierSession) RemoveValueLimitWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemoveValueLimitWhitelistBatch(&_BridgeVerifier.TransactOpts, accounts)
}

// RemoveValueLimitWhitelistBatch is a paid mutator transaction binding the contract method 0xf1f1263d.
//
// Solidity: function removeValueLimitWhitelistBatch(address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RemoveValueLimitWhitelistBatch(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemoveValueLimitWhitelistBatch(&_BridgeVerifier.TransactOpts, accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeVerifier *BridgeVerifierSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RenounceRole(&_BridgeVerifier.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RenounceRole(&_BridgeVerifier.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RevokeRole(&_BridgeVerifier.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RevokeRole(&_BridgeVerifier.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RevokeRoleBatch(&_BridgeVerifier.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RevokeRoleBatch(&_BridgeVerifier.TransactOpts, roles, accounts)
}

// SafePermit is a paid mutator transaction binding the contract method 0xfe81d82e.
//
// Solidity: function safePermit(address token, address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SafePermit(opts *bind.TransactOpts, token common.Address, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "safePermit", token, owner, spender, value, deadline, v, r, s)
}

// SafePermit is a paid mutator transaction binding the contract method 0xfe81d82e.
//
// Solidity: function safePermit(address token, address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeVerifier *BridgeVerifierSession) SafePermit(token common.Address, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SafePermit(&_BridgeVerifier.TransactOpts, token, owner, spender, value, deadline, v, r, s)
}

// SafePermit is a paid mutator transaction binding the contract method 0xfe81d82e.
//
// Solidity: function safePermit(address token, address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SafePermit(token common.Address, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SafePermit(&_BridgeVerifier.TransactOpts, token, owner, spender, value, deadline, v, r, s)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetDefaultExFeeRate(opts *bind.TransactOpts, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setDefaultExFeeRate", exFeeRate)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetDefaultExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetDefaultExFeeRate(&_BridgeVerifier.TransactOpts, exFeeRate)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetDefaultExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetDefaultExFeeRate(&_BridgeVerifier.TransactOpts, exFeeRate)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetDefaultTokenPrice(opts *bind.TransactOpts, defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setDefaultTokenPrice", defaultTokenPrice)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetDefaultTokenPrice(defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetDefaultTokenPrice(&_BridgeVerifier.TransactOpts, defaultTokenPrice)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetDefaultTokenPrice(defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetDefaultTokenPrice(&_BridgeVerifier.TransactOpts, defaultTokenPrice)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetExFeeRate(opts *bind.TransactOpts, token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setExFeeRate", token, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetExFeeRate(token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetExFeeRate(&_BridgeVerifier.TransactOpts, token, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetExFeeRate(token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetExFeeRate(&_BridgeVerifier.TransactOpts, token, exFeeRate)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetExFeeRateBatch(opts *bind.TransactOpts, tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setExFeeRateBatch", tokenList, exFeeRateList)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetExFeeRateBatch(tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetExFeeRateBatch(&_BridgeVerifier.TransactOpts, tokenList, exFeeRateList)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetExFeeRateBatch(tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetExFeeRateBatch(&_BridgeVerifier.TransactOpts, tokenList, exFeeRateList)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetFinalizeBridgeGas(opts *bind.TransactOpts, finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setFinalizeBridgeGas", finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetFinalizeBridgeGas(&_BridgeVerifier.TransactOpts, finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetFinalizeBridgeGas(&_BridgeVerifier.TransactOpts, finalizeBridgeGas)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetMinimumTokenValue(opts *bind.TransactOpts, minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setMinimumTokenValue", minimumTokenValue)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetMinimumTokenValue(minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValue(&_BridgeVerifier.TransactOpts, minimumTokenValue)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetMinimumTokenValue(minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValue(&_BridgeVerifier.TransactOpts, minimumTokenValue)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetPeriodTotalValueThreshold(opts *bind.TransactOpts, periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setPeriodTotalValueThreshold", periodTotalValueThreshold)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetPeriodTotalValueThreshold(periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPeriodTotalValueThreshold(&_BridgeVerifier.TransactOpts, periodTotalValueThreshold)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetPeriodTotalValueThreshold(periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPeriodTotalValueThreshold(&_BridgeVerifier.TransactOpts, periodTotalValueThreshold)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPriceFeed(&_BridgeVerifier.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPriceFeed(&_BridgeVerifier.TransactOpts, _priceFeed)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetTimeWindow(opts *bind.TransactOpts, timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setTimeWindow", timeWindow)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetTimeWindow(timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetTimeWindow(&_BridgeVerifier.TransactOpts, timeWindow)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetTimeWindow(timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetTimeWindow(&_BridgeVerifier.TransactOpts, timeWindow)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetVerificationAmountThreshold(opts *bind.TransactOpts, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setVerificationAmountThreshold", verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetVerificationAmountThreshold(verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetVerificationAmountThreshold(&_BridgeVerifier.TransactOpts, verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetVerificationAmountThreshold(verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetVerificationAmountThreshold(&_BridgeVerifier.TransactOpts, verificationAmountThreshold)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) UpdateGasPrice(opts *bind.TransactOpts, remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "updateGasPrice", remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeVerifier *BridgeVerifierSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.UpdateGasPrice(&_BridgeVerifier.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.UpdateGasPrice(&_BridgeVerifier.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) UpdateGasPriceBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "updateGasPriceBatch", remoteChainIDs, gasPrices)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeVerifier *BridgeVerifierSession) UpdateGasPriceBatch(remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.UpdateGasPriceBatch(&_BridgeVerifier.TransactOpts, remoteChainIDs, gasPrices)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) UpdateGasPriceBatch(remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.UpdateGasPriceBatch(&_BridgeVerifier.TransactOpts, remoteChainIDs, gasPrices)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierTransactor) ValidateBridgeTokenValue(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "validateBridgeTokenValue", token, value)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierSession) ValidateBridgeTokenValue(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.ValidateBridgeTokenValue(&_BridgeVerifier.TransactOpts, token, value)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierTransactorSession) ValidateBridgeTokenValue(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.ValidateBridgeTokenValue(&_BridgeVerifier.TransactOpts, token, value)
}

// ValidateBridgeTokenValue0 is a paid mutator transaction binding the contract method 0x66ac1eec.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value, address to) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierTransactor) ValidateBridgeTokenValue0(opts *bind.TransactOpts, token common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "validateBridgeTokenValue0", token, value, to)
}

// ValidateBridgeTokenValue0 is a paid mutator transaction binding the contract method 0x66ac1eec.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value, address to) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierSession) ValidateBridgeTokenValue0(token common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.ValidateBridgeTokenValue0(&_BridgeVerifier.TransactOpts, token, value, to)
}

// ValidateBridgeTokenValue0 is a paid mutator transaction binding the contract method 0x66ac1eec.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value, address to) returns(uint8 status)
func (_BridgeVerifier *BridgeVerifierTransactorSession) ValidateBridgeTokenValue0(token common.Address, value *big.Int, to common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.ValidateBridgeTokenValue0(&_BridgeVerifier.TransactOpts, token, value, to)
}

// BridgeVerifierDefaultTokenPriceSetIterator is returned from FilterDefaultTokenPriceSet and is used to iterate over the raw logs and unpacked data for DefaultTokenPriceSet events raised by the BridgeVerifier contract.
type BridgeVerifierDefaultTokenPriceSetIterator struct {
	Event *BridgeVerifierDefaultTokenPriceSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierDefaultTokenPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierDefaultTokenPriceSet)
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
		it.Event = new(BridgeVerifierDefaultTokenPriceSet)
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
func (it *BridgeVerifierDefaultTokenPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierDefaultTokenPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierDefaultTokenPriceSet represents a DefaultTokenPriceSet event raised by the BridgeVerifier contract.
type BridgeVerifierDefaultTokenPriceSet struct {
	DefaultTokenPrice *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDefaultTokenPriceSet is a free log retrieval operation binding the contract event 0x1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a.
//
// Solidity: event DefaultTokenPriceSet(uint256 defaultTokenPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterDefaultTokenPriceSet(opts *bind.FilterOpts) (*BridgeVerifierDefaultTokenPriceSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "DefaultTokenPriceSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierDefaultTokenPriceSetIterator{contract: _BridgeVerifier.contract, event: "DefaultTokenPriceSet", logs: logs, sub: sub}, nil
}

// WatchDefaultTokenPriceSet is a free log subscription operation binding the contract event 0x1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a.
//
// Solidity: event DefaultTokenPriceSet(uint256 defaultTokenPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchDefaultTokenPriceSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierDefaultTokenPriceSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "DefaultTokenPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierDefaultTokenPriceSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "DefaultTokenPriceSet", log); err != nil {
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

// ParseDefaultTokenPriceSet is a log parse operation binding the contract event 0x1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a.
//
// Solidity: event DefaultTokenPriceSet(uint256 defaultTokenPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseDefaultTokenPriceSet(log types.Log) (*BridgeVerifierDefaultTokenPriceSet, error) {
	event := new(BridgeVerifierDefaultTokenPriceSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "DefaultTokenPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierExchangeFeeUpdatedIterator is returned from FilterExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for ExchangeFeeUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierExchangeFeeUpdatedIterator struct {
	Event *BridgeVerifierExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierExchangeFeeUpdated)
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
		it.Event = new(BridgeVerifierExchangeFeeUpdated)
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
func (it *BridgeVerifierExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierExchangeFeeUpdated represents a ExchangeFeeUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierExchangeFeeUpdated struct {
	Token     common.Address
	ExFeeRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchangeFeeUpdated is a free log retrieval operation binding the contract event 0x386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1d.
//
// Solidity: event ExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterExchangeFeeUpdated(opts *bind.FilterOpts, token []common.Address) (*BridgeVerifierExchangeFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "ExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierExchangeFeeUpdatedIterator{contract: _BridgeVerifier.contract, event: "ExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeFeeUpdated is a free log subscription operation binding the contract event 0x386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1d.
//
// Solidity: event ExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierExchangeFeeUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "ExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierExchangeFeeUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "ExchangeFeeUpdated", log); err != nil {
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

// ParseExchangeFeeUpdated is a log parse operation binding the contract event 0x386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1d.
//
// Solidity: event ExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseExchangeFeeUpdated(log types.Log) (*BridgeVerifierExchangeFeeUpdated, error) {
	event := new(BridgeVerifierExchangeFeeUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "ExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierFinalizeBridgeGasSetIterator is returned from FilterFinalizeBridgeGasSet and is used to iterate over the raw logs and unpacked data for FinalizeBridgeGasSet events raised by the BridgeVerifier contract.
type BridgeVerifierFinalizeBridgeGasSetIterator struct {
	Event *BridgeVerifierFinalizeBridgeGasSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierFinalizeBridgeGasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierFinalizeBridgeGasSet)
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
		it.Event = new(BridgeVerifierFinalizeBridgeGasSet)
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
func (it *BridgeVerifierFinalizeBridgeGasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierFinalizeBridgeGasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierFinalizeBridgeGasSet represents a FinalizeBridgeGasSet event raised by the BridgeVerifier contract.
type BridgeVerifierFinalizeBridgeGasSet struct {
	FinalizeBridgeGas *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBridgeGasSet is a free log retrieval operation binding the contract event 0xbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b.
//
// Solidity: event FinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterFinalizeBridgeGasSet(opts *bind.FilterOpts) (*BridgeVerifierFinalizeBridgeGasSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "FinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierFinalizeBridgeGasSetIterator{contract: _BridgeVerifier.contract, event: "FinalizeBridgeGasSet", logs: logs, sub: sub}, nil
}

// WatchFinalizeBridgeGasSet is a free log subscription operation binding the contract event 0xbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b.
//
// Solidity: event FinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchFinalizeBridgeGasSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierFinalizeBridgeGasSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "FinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierFinalizeBridgeGasSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "FinalizeBridgeGasSet", log); err != nil {
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

// ParseFinalizeBridgeGasSet is a log parse operation binding the contract event 0xbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b.
//
// Solidity: event FinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseFinalizeBridgeGasSet(log types.Log) (*BridgeVerifierFinalizeBridgeGasSet, error) {
	event := new(BridgeVerifierFinalizeBridgeGasSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "FinalizeBridgeGasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierGasPriceUpdatedIterator is returned from FilterGasPriceUpdated and is used to iterate over the raw logs and unpacked data for GasPriceUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierGasPriceUpdatedIterator struct {
	Event *BridgeVerifierGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierGasPriceUpdated)
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
		it.Event = new(BridgeVerifierGasPriceUpdated)
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
func (it *BridgeVerifierGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierGasPriceUpdated represents a GasPriceUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierGasPriceUpdated struct {
	RemoteChainID *big.Int
	GasPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGasPriceUpdated is a free log retrieval operation binding the contract event 0x4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3.
//
// Solidity: event GasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterGasPriceUpdated(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BridgeVerifierGasPriceUpdatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "GasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierGasPriceUpdatedIterator{contract: _BridgeVerifier.contract, event: "GasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchGasPriceUpdated is a free log subscription operation binding the contract event 0x4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3.
//
// Solidity: event GasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierGasPriceUpdated, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "GasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierGasPriceUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "GasPriceUpdated", log); err != nil {
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

// ParseGasPriceUpdated is a log parse operation binding the contract event 0x4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3.
//
// Solidity: event GasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseGasPriceUpdated(log types.Log) (*BridgeVerifierGasPriceUpdated, error) {
	event := new(BridgeVerifierGasPriceUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "GasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierMinimumTokenValueSetIterator is returned from FilterMinimumTokenValueSet and is used to iterate over the raw logs and unpacked data for MinimumTokenValueSet events raised by the BridgeVerifier contract.
type BridgeVerifierMinimumTokenValueSetIterator struct {
	Event *BridgeVerifierMinimumTokenValueSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierMinimumTokenValueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierMinimumTokenValueSet)
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
		it.Event = new(BridgeVerifierMinimumTokenValueSet)
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
func (it *BridgeVerifierMinimumTokenValueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierMinimumTokenValueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierMinimumTokenValueSet represents a MinimumTokenValueSet event raised by the BridgeVerifier contract.
type BridgeVerifierMinimumTokenValueSet struct {
	MinimumTokenValue *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMinimumTokenValueSet is a free log retrieval operation binding the contract event 0x0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e9.
//
// Solidity: event MinimumTokenValueSet(uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterMinimumTokenValueSet(opts *bind.FilterOpts) (*BridgeVerifierMinimumTokenValueSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "MinimumTokenValueSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierMinimumTokenValueSetIterator{contract: _BridgeVerifier.contract, event: "MinimumTokenValueSet", logs: logs, sub: sub}, nil
}

// WatchMinimumTokenValueSet is a free log subscription operation binding the contract event 0x0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e9.
//
// Solidity: event MinimumTokenValueSet(uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchMinimumTokenValueSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierMinimumTokenValueSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "MinimumTokenValueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierMinimumTokenValueSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "MinimumTokenValueSet", log); err != nil {
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

// ParseMinimumTokenValueSet is a log parse operation binding the contract event 0x0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e9.
//
// Solidity: event MinimumTokenValueSet(uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseMinimumTokenValueSet(log types.Log) (*BridgeVerifierMinimumTokenValueSet, error) {
	event := new(BridgeVerifierMinimumTokenValueSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "MinimumTokenValueSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierPeriodTotalValueThresholdSetIterator is returned from FilterPeriodTotalValueThresholdSet and is used to iterate over the raw logs and unpacked data for PeriodTotalValueThresholdSet events raised by the BridgeVerifier contract.
type BridgeVerifierPeriodTotalValueThresholdSetIterator struct {
	Event *BridgeVerifierPeriodTotalValueThresholdSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierPeriodTotalValueThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierPeriodTotalValueThresholdSet)
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
		it.Event = new(BridgeVerifierPeriodTotalValueThresholdSet)
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
func (it *BridgeVerifierPeriodTotalValueThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierPeriodTotalValueThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierPeriodTotalValueThresholdSet represents a PeriodTotalValueThresholdSet event raised by the BridgeVerifier contract.
type BridgeVerifierPeriodTotalValueThresholdSet struct {
	PeriodTotalValueThreshold *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPeriodTotalValueThresholdSet is a free log retrieval operation binding the contract event 0x7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e2.
//
// Solidity: event PeriodTotalValueThresholdSet(uint256 periodTotalValueThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterPeriodTotalValueThresholdSet(opts *bind.FilterOpts) (*BridgeVerifierPeriodTotalValueThresholdSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "PeriodTotalValueThresholdSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierPeriodTotalValueThresholdSetIterator{contract: _BridgeVerifier.contract, event: "PeriodTotalValueThresholdSet", logs: logs, sub: sub}, nil
}

// WatchPeriodTotalValueThresholdSet is a free log subscription operation binding the contract event 0x7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e2.
//
// Solidity: event PeriodTotalValueThresholdSet(uint256 periodTotalValueThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchPeriodTotalValueThresholdSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierPeriodTotalValueThresholdSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "PeriodTotalValueThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierPeriodTotalValueThresholdSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "PeriodTotalValueThresholdSet", log); err != nil {
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

// ParsePeriodTotalValueThresholdSet is a log parse operation binding the contract event 0x7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e2.
//
// Solidity: event PeriodTotalValueThresholdSet(uint256 periodTotalValueThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) ParsePeriodTotalValueThresholdSet(log types.Log) (*BridgeVerifierPeriodTotalValueThresholdSet, error) {
	event := new(BridgeVerifierPeriodTotalValueThresholdSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "PeriodTotalValueThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierPriceFeedUpdatedIterator is returned from FilterPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierPriceFeedUpdatedIterator struct {
	Event *BridgeVerifierPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierPriceFeedUpdated)
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
		it.Event = new(BridgeVerifierPriceFeedUpdated)
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
func (it *BridgeVerifierPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierPriceFeedUpdated represents a PriceFeedUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdated is a free log retrieval operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterPriceFeedUpdated(opts *bind.FilterOpts, priceFeed []common.Address) (*BridgeVerifierPriceFeedUpdatedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "PriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierPriceFeedUpdatedIterator{contract: _BridgeVerifier.contract, event: "PriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdated is a free log subscription operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierPriceFeedUpdated, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "PriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierPriceFeedUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
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

// ParsePriceFeedUpdated is a log parse operation binding the contract event 0xe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6.
//
// Solidity: event PriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) ParsePriceFeedUpdated(log types.Log) (*BridgeVerifierPriceFeedUpdated, error) {
	event := new(BridgeVerifierPriceFeedUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeVerifier contract.
type BridgeVerifierRoleAdminChangedIterator struct {
	Event *BridgeVerifierRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierRoleAdminChanged)
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
		it.Event = new(BridgeVerifierRoleAdminChanged)
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
func (it *BridgeVerifierRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeVerifier contract.
type BridgeVerifierRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeVerifierRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierRoleAdminChangedIterator{contract: _BridgeVerifier.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeVerifierRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierRoleAdminChanged)
				if err := _BridgeVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeVerifier *BridgeVerifierFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeVerifierRoleAdminChanged, error) {
	event := new(BridgeVerifierRoleAdminChanged)
	if err := _BridgeVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeVerifier contract.
type BridgeVerifierRoleGrantedIterator struct {
	Event *BridgeVerifierRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierRoleGranted)
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
		it.Event = new(BridgeVerifierRoleGranted)
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
func (it *BridgeVerifierRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierRoleGranted represents a RoleGranted event raised by the BridgeVerifier contract.
type BridgeVerifierRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeVerifierRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierRoleGrantedIterator{contract: _BridgeVerifier.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeVerifierRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierRoleGranted)
				if err := _BridgeVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeVerifier *BridgeVerifierFilterer) ParseRoleGranted(log types.Log) (*BridgeVerifierRoleGranted, error) {
	event := new(BridgeVerifierRoleGranted)
	if err := _BridgeVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeVerifier contract.
type BridgeVerifierRoleRevokedIterator struct {
	Event *BridgeVerifierRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierRoleRevoked)
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
		it.Event = new(BridgeVerifierRoleRevoked)
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
func (it *BridgeVerifierRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierRoleRevoked represents a RoleRevoked event raised by the BridgeVerifier contract.
type BridgeVerifierRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeVerifierRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierRoleRevokedIterator{contract: _BridgeVerifier.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeVerifierRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierRoleRevoked)
				if err := _BridgeVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeVerifier *BridgeVerifierFilterer) ParseRoleRevoked(log types.Log) (*BridgeVerifierRoleRevoked, error) {
	event := new(BridgeVerifierRoleRevoked)
	if err := _BridgeVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierTimeWindowSetIterator is returned from FilterTimeWindowSet and is used to iterate over the raw logs and unpacked data for TimeWindowSet events raised by the BridgeVerifier contract.
type BridgeVerifierTimeWindowSetIterator struct {
	Event *BridgeVerifierTimeWindowSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierTimeWindowSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierTimeWindowSet)
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
		it.Event = new(BridgeVerifierTimeWindowSet)
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
func (it *BridgeVerifierTimeWindowSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierTimeWindowSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierTimeWindowSet represents a TimeWindowSet event raised by the BridgeVerifier contract.
type BridgeVerifierTimeWindowSet struct {
	TimeWindow *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTimeWindowSet is a free log retrieval operation binding the contract event 0x52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a4.
//
// Solidity: event TimeWindowSet(uint256 timeWindow)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterTimeWindowSet(opts *bind.FilterOpts) (*BridgeVerifierTimeWindowSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "TimeWindowSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierTimeWindowSetIterator{contract: _BridgeVerifier.contract, event: "TimeWindowSet", logs: logs, sub: sub}, nil
}

// WatchTimeWindowSet is a free log subscription operation binding the contract event 0x52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a4.
//
// Solidity: event TimeWindowSet(uint256 timeWindow)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchTimeWindowSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierTimeWindowSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "TimeWindowSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierTimeWindowSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "TimeWindowSet", log); err != nil {
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

// ParseTimeWindowSet is a log parse operation binding the contract event 0x52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a4.
//
// Solidity: event TimeWindowSet(uint256 timeWindow)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseTimeWindowSet(log types.Log) (*BridgeVerifierTimeWindowSet, error) {
	event := new(BridgeVerifierTimeWindowSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "TimeWindowSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierValueLimitWhitelistAddedIterator is returned from FilterValueLimitWhitelistAdded and is used to iterate over the raw logs and unpacked data for ValueLimitWhitelistAdded events raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistAddedIterator struct {
	Event *BridgeVerifierValueLimitWhitelistAdded // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierValueLimitWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierValueLimitWhitelistAdded)
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
		it.Event = new(BridgeVerifierValueLimitWhitelistAdded)
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
func (it *BridgeVerifierValueLimitWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierValueLimitWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierValueLimitWhitelistAdded represents a ValueLimitWhitelistAdded event raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValueLimitWhitelistAdded is a free log retrieval operation binding the contract event 0xe1d1777e4c7896f6e13f47c189990ed90f2210cb88f3836cd7bd0cb1f820945b.
//
// Solidity: event ValueLimitWhitelistAdded(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterValueLimitWhitelistAdded(opts *bind.FilterOpts, account []common.Address) (*BridgeVerifierValueLimitWhitelistAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "ValueLimitWhitelistAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierValueLimitWhitelistAddedIterator{contract: _BridgeVerifier.contract, event: "ValueLimitWhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchValueLimitWhitelistAdded is a free log subscription operation binding the contract event 0xe1d1777e4c7896f6e13f47c189990ed90f2210cb88f3836cd7bd0cb1f820945b.
//
// Solidity: event ValueLimitWhitelistAdded(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchValueLimitWhitelistAdded(opts *bind.WatchOpts, sink chan<- *BridgeVerifierValueLimitWhitelistAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "ValueLimitWhitelistAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierValueLimitWhitelistAdded)
				if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistAdded", log); err != nil {
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

// ParseValueLimitWhitelistAdded is a log parse operation binding the contract event 0xe1d1777e4c7896f6e13f47c189990ed90f2210cb88f3836cd7bd0cb1f820945b.
//
// Solidity: event ValueLimitWhitelistAdded(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseValueLimitWhitelistAdded(log types.Log) (*BridgeVerifierValueLimitWhitelistAdded, error) {
	event := new(BridgeVerifierValueLimitWhitelistAdded)
	if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierValueLimitWhitelistBypassedIterator is returned from FilterValueLimitWhitelistBypassed and is used to iterate over the raw logs and unpacked data for ValueLimitWhitelistBypassed events raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistBypassedIterator struct {
	Event *BridgeVerifierValueLimitWhitelistBypassed // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierValueLimitWhitelistBypassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierValueLimitWhitelistBypassed)
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
		it.Event = new(BridgeVerifierValueLimitWhitelistBypassed)
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
func (it *BridgeVerifierValueLimitWhitelistBypassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierValueLimitWhitelistBypassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierValueLimitWhitelistBypassed represents a ValueLimitWhitelistBypassed event raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistBypassed struct {
	Token common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValueLimitWhitelistBypassed is a free log retrieval operation binding the contract event 0xdfe196174d835a3614fc87874a6a97dde9719b32d7d503610251887094dccfce.
//
// Solidity: event ValueLimitWhitelistBypassed(address indexed token, address indexed to, uint256 value)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterValueLimitWhitelistBypassed(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeVerifierValueLimitWhitelistBypassedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "ValueLimitWhitelistBypassed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierValueLimitWhitelistBypassedIterator{contract: _BridgeVerifier.contract, event: "ValueLimitWhitelistBypassed", logs: logs, sub: sub}, nil
}

// WatchValueLimitWhitelistBypassed is a free log subscription operation binding the contract event 0xdfe196174d835a3614fc87874a6a97dde9719b32d7d503610251887094dccfce.
//
// Solidity: event ValueLimitWhitelistBypassed(address indexed token, address indexed to, uint256 value)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchValueLimitWhitelistBypassed(opts *bind.WatchOpts, sink chan<- *BridgeVerifierValueLimitWhitelistBypassed, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "ValueLimitWhitelistBypassed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierValueLimitWhitelistBypassed)
				if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistBypassed", log); err != nil {
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

// ParseValueLimitWhitelistBypassed is a log parse operation binding the contract event 0xdfe196174d835a3614fc87874a6a97dde9719b32d7d503610251887094dccfce.
//
// Solidity: event ValueLimitWhitelistBypassed(address indexed token, address indexed to, uint256 value)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseValueLimitWhitelistBypassed(log types.Log) (*BridgeVerifierValueLimitWhitelistBypassed, error) {
	event := new(BridgeVerifierValueLimitWhitelistBypassed)
	if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistBypassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierValueLimitWhitelistRemovedIterator is returned from FilterValueLimitWhitelistRemoved and is used to iterate over the raw logs and unpacked data for ValueLimitWhitelistRemoved events raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistRemovedIterator struct {
	Event *BridgeVerifierValueLimitWhitelistRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierValueLimitWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierValueLimitWhitelistRemoved)
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
		it.Event = new(BridgeVerifierValueLimitWhitelistRemoved)
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
func (it *BridgeVerifierValueLimitWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierValueLimitWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierValueLimitWhitelistRemoved represents a ValueLimitWhitelistRemoved event raised by the BridgeVerifier contract.
type BridgeVerifierValueLimitWhitelistRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValueLimitWhitelistRemoved is a free log retrieval operation binding the contract event 0xcf6ae774f84d549d12db18368181005de6e396adc8740a42b1df7a36352edad3.
//
// Solidity: event ValueLimitWhitelistRemoved(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterValueLimitWhitelistRemoved(opts *bind.FilterOpts, account []common.Address) (*BridgeVerifierValueLimitWhitelistRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "ValueLimitWhitelistRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierValueLimitWhitelistRemovedIterator{contract: _BridgeVerifier.contract, event: "ValueLimitWhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchValueLimitWhitelistRemoved is a free log subscription operation binding the contract event 0xcf6ae774f84d549d12db18368181005de6e396adc8740a42b1df7a36352edad3.
//
// Solidity: event ValueLimitWhitelistRemoved(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchValueLimitWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *BridgeVerifierValueLimitWhitelistRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "ValueLimitWhitelistRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierValueLimitWhitelistRemoved)
				if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistRemoved", log); err != nil {
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

// ParseValueLimitWhitelistRemoved is a log parse operation binding the contract event 0xcf6ae774f84d549d12db18368181005de6e396adc8740a42b1df7a36352edad3.
//
// Solidity: event ValueLimitWhitelistRemoved(address indexed account)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseValueLimitWhitelistRemoved(log types.Log) (*BridgeVerifierValueLimitWhitelistRemoved, error) {
	event := new(BridgeVerifierValueLimitWhitelistRemoved)
	if err := _BridgeVerifier.contract.UnpackLog(event, "ValueLimitWhitelistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierVerificationAmountThresholdSetIterator is returned from FilterVerificationAmountThresholdSet and is used to iterate over the raw logs and unpacked data for VerificationAmountThresholdSet events raised by the BridgeVerifier contract.
type BridgeVerifierVerificationAmountThresholdSetIterator struct {
	Event *BridgeVerifierVerificationAmountThresholdSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierVerificationAmountThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierVerificationAmountThresholdSet)
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
		it.Event = new(BridgeVerifierVerificationAmountThresholdSet)
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
func (it *BridgeVerifierVerificationAmountThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierVerificationAmountThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierVerificationAmountThresholdSet represents a VerificationAmountThresholdSet event raised by the BridgeVerifier contract.
type BridgeVerifierVerificationAmountThresholdSet struct {
	VerificationAmountThreshold *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterVerificationAmountThresholdSet is a free log retrieval operation binding the contract event 0xf905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2.
//
// Solidity: event VerificationAmountThresholdSet(uint256 verificationAmountThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterVerificationAmountThresholdSet(opts *bind.FilterOpts) (*BridgeVerifierVerificationAmountThresholdSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "VerificationAmountThresholdSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierVerificationAmountThresholdSetIterator{contract: _BridgeVerifier.contract, event: "VerificationAmountThresholdSet", logs: logs, sub: sub}, nil
}

// WatchVerificationAmountThresholdSet is a free log subscription operation binding the contract event 0xf905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2.
//
// Solidity: event VerificationAmountThresholdSet(uint256 verificationAmountThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchVerificationAmountThresholdSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierVerificationAmountThresholdSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "VerificationAmountThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierVerificationAmountThresholdSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "VerificationAmountThresholdSet", log); err != nil {
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

// ParseVerificationAmountThresholdSet is a log parse operation binding the contract event 0xf905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2.
//
// Solidity: event VerificationAmountThresholdSet(uint256 verificationAmountThreshold)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseVerificationAmountThresholdSet(log types.Log) (*BridgeVerifierVerificationAmountThresholdSet, error) {
	event := new(BridgeVerifierVerificationAmountThresholdSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "VerificationAmountThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
