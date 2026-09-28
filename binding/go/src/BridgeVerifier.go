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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addValueLimitWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addValueLimitWhitelistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMinimumTokenValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPeriodTotalValueThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCurrentScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMovementHistory\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"history\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getTokenPriceWithValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getValueLimitWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"page\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValueLimitWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValueLimitWhitelistLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerificationAmountThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValueLimitWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeValueLimitWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"removeValueLimitWhitelistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"safePermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"setDefaultTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setExFeeRateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValueOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minimumTokenValueList\",\"type\":\"uint256[]\"}],\"name\":\"setMinimumTokenValueOfBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"setPeriodTotalValueThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasPrices\",\"type\":\"uint256[]\"}],\"name\":\"updateGasPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"DefaultTokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"ExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"FinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"GasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueOfSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"PeriodTotalValueThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"TimeWindowSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValueLimitWhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueLimitWhitelistBypassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValueLimitWhitelistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"VerificationAmountThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeVerifierCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidExFeeRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidTimeWindow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BridgeVerifierWhitelistAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BridgeVerifierWhitelistNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"75a4463e": "addValueLimitWhitelist(address)",
		"bd574e83": "addValueLimitWhitelistBatch(address[])",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"96ce0795": "denominator()",
		"2fe3b6cf": "dollarDecimals()",
		"d1c01543": "getGasPrice(uint256)",
		"2a4d2f32": "getMinimumTokenValue()",
		"c5fe3bf4": "getMinimumTokenValueOf(address)",
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
		"3d949a66": "setMinimumTokenValueOf(address,uint256)",
		"edba2c77": "setMinimumTokenValueOfBatch(address[],uint256[])",
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
	Bin: "0x60803461027457601f61339238819003918201601f19168301916001600160401b038311848410176102785780849261014094604052833981010312610274576100488161028c565b906100556020820161028c565b916100626040830161028c565b926060830151608084015160a08501519160c08601519360e08701519561012061010089015198015198831561023a576001600160a01b03821615610205576001600160a01b038316156101d6576001600160a01b03169182156101a357841561016957612710861161015a57610e108a066101315789158015610140575b15610131576100f26100f8926102a0565b5061030d565b50600180546001600160a01b031916919091179055600255600455600355600555600655600755600855604051612ed8908161049a8239f35b6316a6929960e31b5f5260045ffd5b50610e108a101580156100e1575062093a808a11156100e1565b63aa80145d60e01b5f5260045ffd5b6040516323056b1360e01b815260206004820152601160248201527064656661756c74546f6b656e507269636560781b6044820152606490fd5b6040516323056b1360e01b815260206004820152600a6024820152697072696365466565645f60b01b6044820152606490fd5b6040516323056b1360e01b815260206004820152600660248201526562726964676560d01b6044820152606490fd5b6040516323056b1360e01b815260206004820152600c60248201526b34b734ba34b0b627bbb732b960a11b6044820152606490fd5b6040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361027457565b906102ab825f6103a1565b91826102b45750565b5f8052600d6020526001600160a01b03166102ef817f81955a0a11e65eac625c29e8882660bae4e165a75d72780094acae8ece9a29ee610429565b156102f75750565b637a854c7360e11b5f526004525f60245260445ffd5b90610325825f5160206133725f395f51905f526103a1565b918261032e5750565b5f5160206133725f395f51905f525f52600d6020526001600160a01b0316610376817fd06032cb802eef06aa279a4e954daed9c87dafcc32cbe2c5419da429e2c42a5f610429565b1561037e5750565b637a854c7360e11b5f526004525f5160206133725f395f51905f5260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610423575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b6001810190825f528160205260405f2054155f1461049257805468010000000000000000811015610278576001810180835581101561047e578390825f5260205f20015554915f5260205260405f2055600190565b634e487b7160e01b5f52603260045260245ffd5b5050505f9056fe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461133f57508063020a22121461129f5780630e4e96bf1461118857806317cd9cdd146111465780631f96131e146111075780631fd4d631146110ed578063248a9ca3146110c75780632a4d2f32146110aa5780632f2ff15d1461104c5780632fe3b6cf14610fc257806336568abe14610f7e578063365d71e414610f245780633d949a6614610ec45780633f4bdec514610e9457806349c9215e14610e075780635566ca0914610dea5780635b42895f14610d9f5780635ec8f51b14610d225780636332aec614610cf457806366ac1eec14610caf578063724e78da14610c10578063741bef1a14610be8578063751ccd7914610b2557806375a4463e14610ad45780638d33199614610a4757806391d14854146109ff57806396ce0795146109e35780639745e4ba146109ab578063979eb82d1461098e5780639dc696cc14610942578063a217fddf14610928578063a3246ad3146108c4578063ab509f3a14610878578063ad1434ca14610825578063b2b49e2e14610780578063ba8d25f814610738578063bd574e8314610686578063be716b0a14610669578063c3de108b1461061d578063c5fe3bf4146105e5578063ce35dd7e146105c8578063d02641a01461055f578063d1c0154314610535578063d547741f146104ff578063df6e87dc146104b8578063edba2c7714610420578063f1f1263d1461039f578063f289d3ba14610312578063fb1da429146102a15763fe81d82e14610242575f80fd5b3461029d5761010036600319011261029d5761025c6113e0565b6102646113f6565b9061026d611526565b60a43560ff8116810361029d5761029b936102866124a2565b60e4359360c435936084359260643592612025565b005b5f80fd5b3461029d575f36600319011261029d57604051600e80548083525f91825260208301915f516020612e035f395f51905f5291905b8181106102fc576102f8856102ec81870382611392565b60405191829182611552565b0390f35b82548452602090930192600192830192016102d5565b3461029d57602036600319011261029d5760043561032e612511565b8015610365576020817fbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b92600255604051908152a1005b6040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606490fd5b3461029d576103ad366116cb565b906103b6612511565b5f5b8281106103c157005b6103d46103cf828585611fe7565b611ff7565b6103dc612511565b6001600160a01b03166103ee81612906565b1561040e57906001915f516020612e635f395f51905f525f80a2016103b8565b63a60867e560e01b5f5260045260245ffd5b3461029d5761042e3661147d565b6104366123c4565b81518151036104a9575f5b825181101561029b576001906001600160a01b0361045f828661171d565b51165f516020612e835f395f51905f52602061047b848761171d565b516104846123c4565b61048f841515611874565b835f52601082528060405f2055604051908152a201610441565b63322473c360e01b5f5260045ffd5b3461029d57606036600319011261029d576127106104f26102f86104e56104dd6113f6565b60043561190f565b60449592949195356118de565b04604051938493846116b5565b3461029d57604036600319011261029d5761029b60043561051e6113f6565b9061053061052b8261184a565b612580565b6125b8565b3461029d57602036600319011261029d576004355f52600a602052602060405f2054604051908152f35b3461029d57602036600319011261029d576105786113e0565b6001600160a01b038116906105b49061059b906105ae6105a66105a0848761260c565b6118cd565b926126b4565b93909561260c565b91612729565b604080519215158352602083019190915290f35b3461029d575f36600319011261029d576020600e54604051908152f35b3461029d57602036600319011261029d576001600160a01b036106066113e0565b165f526010602052602060405f2054604051908152f35b3461029d57602036600319011261029d577ff905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2602060043561065c612511565b80600655604051908152a1005b3461029d575f36600319011261029d576020600654604051908152f35b3461029d57610694366116cb565b9061069d612511565b5f5b8281106106a857005b6106b66103cf828585611fe7565b6106be612511565b6001600160a01b03168015610708576106d681612b90565b156106f657906001915f516020612e235f395f51905f525f80a20161069f565b635bf16f2960e11b5f5260045260245ffd5b6040516323056b1360e01b81526020600482015260076024820152661858d8dbdd5b9d60ca1b6044820152606490fd5b3461029d57604036600319011261029d576107516113e0565b6105b461059b610776610763846126b4565b90949092906001600160a01b031661260c565b9060243590612729565b3461029d5761078e36611594565b908051825103610816575f5b825181101561029b576107ad818361171d565b516001600160a01b036107c0838661171d565b51166107ce61052b8361184a565b6107d881836127fe565b6107e7575b505060010161079a565b815f52600d6020526107fc8160405f20612be5565b6107dd575b637a854c7360e11b5f5260045260245260445ffd5b6374ddb37160e01b5f5260045ffd5b3461029d57602036600319011261029d576004356108416123c4565b6127108111610869575f516020612e435f395f51905f526020825f93600355604051908152a2005b63aa80145d60e01b5f5260045ffd5b3461029d57602036600319011261029d577f7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e260206004356108b7612511565b80600755604051908152a1005b3461029d57602036600319011261029d576004355f52600d60205260405f206040519081602082549182815201915f5260205f20905f5b818110610912576102f8856102ec81870382611392565b82548452602090930192600192830192016108fb565b3461029d575f36600319011261029d5760206040515f8152f35b3461029d57602036600319011261029d577f0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e960206004356109816123c4565b80600555604051908152a1005b3461029d575f36600319011261029d576020600754604051908152f35b3461029d57602036600319011261029d576001600160a01b036109cc6113e0565b165f52600b602052602060405f2054604051908152f35b3461029d575f36600319011261029d5760206040516127108152f35b3461029d57604036600319011261029d57610a186113f6565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461029d57602036600319011261029d57600435610a63612511565b610e108106610aab5780158015610aba575b15610aab576020817f52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a492600855604051908152a1005b6316a6929960e31b5f5260045ffd5b50610e108110158015610a75575062093a80811115610a75565b3461029d57602036600319011261029d57610aed6113e0565b610af5612511565b6001600160a01b0316801561070857610b0d81612b90565b156106f6575f516020612e235f395f51905f525f80a2005b3461029d57604036600319011261029d576004356001600160401b03811161029d57610b55903690600401611420565b6024356001600160401b03811161029d57610b74903690600401611420565b610b7c612433565b81518151036104a9575f5b825181101561029b5780610b9d6001928461171d565b51610ba8828661171d565b515f52600a60205260405f2055610bbf818561171d565b515f516020612de35f395f51905f526020610bda848761171d565b51604051908152a201610b87565b3461029d575f36600319011261029d576001546040516001600160a01b039091168152602090f35b3461029d57602036600319011261029d576004356001600160a01b0381169081900361029d57610c3e612511565b8015610c7c57600180546001600160a01b031916821790557fe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e65f80a2005b6040516323056b1360e01b815260206004820152600a60248201526917dc1c9a58d95199595960b21b6044820152606490fd5b3461029d57606036600319011261029d576102f8610ce8610cce6113e0565b610cd6611526565b90610cdf6124a2565b60243590611d87565b6040519182918261168e565b3461029d57604036600319011261029d576102f8610d136104dd6113f6565b604093919351938493846116b5565b3461029d57604036600319011261029d57610d3b6113e0565b60243590610d476123c4565b6001600160a01b031690610d5c821515611874565b6127108111801590610d95575b156108695760205f516020612e435f395f51905f5291835f52600982528060405f2055604051908152a2005b505f198114610d69565b3461029d57602036600319011261029d57610db86113e0565b610dc0612511565b6001600160a01b0316610dd281612906565b1561040e575f516020612e635f395f51905f525f80a2005b3461029d575f36600319011261029d576020600854604051908152f35b3461029d57602036600319011261029d57600435610e236123c4565b8015610e5a576020817f1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a92600455604051908152a1005b6040516323056b1360e01b815260206004820152601160248201527064656661756c74546f6b656e507269636560781b6044820152606490fd5b3461029d57604036600319011261029d576102f8610ce8610eb36113e0565b610ebb6124a2565b60243590611a20565b3461029d57604036600319011261029d57610edd6113e0565b5f516020612e835f395f51905f52602060243592610ef96123c4565b6001600160a01b031692610f0e841515611874565b835f52601082528060405f2055604051908152a2005b3461029d57610f3236611594565b908051825103610816575f5b825181101561029b5780610f77610f576001938561171d565b51838060a01b03610f68848861171d565b51169061053061052b8261184a565b5001610f3e565b3461029d57604036600319011261029d57610f976113f6565b336001600160a01b03821603610fb35761029b906004356125b8565b63334bd91960e11b5f5260045ffd5b3461029d575f36600319011261029d57600154604051632fe3b6cf60e01b815290602090829060049082906001600160a01b03165afa8015611041576020915f91611014575b5060ff60405191168152f35b6110349150823d841161103a575b61102c8183611392565b81019061185b565b82611008565b503d611022565b6040513d5f823e3d90fd5b3461029d57604036600319011261029d576004356110686113f6565b61107461052b8361184a565b61107e81836127fe565b61108457005b815f52600d6020526110a360405f209160018060a01b03168092612be5565b1561080157005b3461029d575f36600319011261029d576020600554604051908152f35b3461029d57602036600319011261029d5760206110e560043561184a565b604051908152f35b3461029d576102f86102ec6111013661153c565b90611788565b3461029d5760205f516020612de35f395f51905f526111253661153c565b611130939193612433565b835f52600a82528060405f2055604051908152a2005b3461029d57602036600319011261029d57602061117e6001600160a01b0361116c6113e0565b165f52600f60205260405f2054151590565b6040519015158152f35b3461029d57602036600319011261029d576001600160a01b036111a96113e0565b16805f52600c6020526111be60405f20612395565b906111c8826113c9565b916111d66040519384611392565b8083526111e2816113c9565b602084019290601f19013684375f5b82811061123c578385604051918291602083019060208452518091526040830191905f5b818110611223575050500390f35b8251845285945060209384019390920191600101611215565b815f52600c60205260405f209061125282612395565b81101561128d5781546001600160801b0390811682821601165f90815260019283016020526040902054611286828861171d565b52016111f1565b634e487b715f5260326020526024601cfd5b3461029d576112ad3661147d565b9080518251036104a9575f5b815181101561029b576001600160a01b036112d4828461171d565b5116906112e1818561171d565b51916112eb6123c4565b6112f6811515611874565b6127108311801590611335575b15610869575f516020612e435f395f51905f526020600194835f52600982528060405f2055604051908152a2016112b9565b505f198314611303565b3461029d57602036600319011261029d576004359063ffffffff60e01b821680920361029d57602091637965db0b60e01b8114908115611381575b5015158152f35b6301ffc9a760e01b1490508361137a565b601f909101601f19168101906001600160401b038211908210176113b557604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b0381116113b55760051b60200190565b600435906001600160a01b038216820361029d57565b602435906001600160a01b038216820361029d57565b35906001600160a01b038216820361029d57565b9080601f8301121561029d578135611437816113c9565b926114456040519485611392565b81845260208085019260051b82010192831161029d57602001905b82821061146d5750505090565b8135815260209182019101611460565b90604060031983011261029d576004356001600160401b03811161029d578260238201121561029d578060040135906114b5826113c9565b916114c36040519384611392565b8083526024602084019160051b8301019185831161029d57602401905b82821061150e57509193602435925090506001600160401b03821161029d5761150b91600401611420565b90565b6020809161151b8461140c565b8152019101906114e0565b604435906001600160a01b038216820361029d57565b604090600319011261029d576004359060243590565b60206040818301928281528451809452019201905f5b8181106115755750505090565b82516001600160a01b0316845260209384019390920191600101611568565b90604060031983011261029d576004356001600160401b03811161029d578260238201121561029d578060040135906115cc826113c9565b916115da6040519384611392565b8083526024602084019160051b8301019185831161029d57602401905b82821061167e57509193602435925090506001600160401b03821161029d578060238301121561029d57816004013561162f816113c9565b9261163d6040519485611392565b8184526024602085019260051b82010192831161029d57602401905b8282106116665750505090565b602080916116738461140c565b815201910190611659565b81358152602091820191016115f7565b91906020830192600c8210156116a15752565b634e487b7160e01b5f52602160045260245ffd5b6040919493926060820195825260208201520152565b90602060031983011261029d576004356001600160401b03811161029d578260238201121561029d576004810135926001600160401b03841161029d5760248460051b8301011161029d576024019190565b80518210156117315760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b6001600160c01b039081039190821161175a57565b634e487b7160e01b5f52601160045260245ffd5b9190820391821161175a57565b9190820180921161175a57565b9190600e54808410801590611842575b61183a57836117a69161176e565b908180821091180218916117b9836113c9565b926117c76040519485611392565b808452601f196117d6826113c9565b01366020860137835f92600e54935b8381106117f3575050505050565b6117fd818361177b565b908582101561173157600e5f525f516020612e035f395f51905f5290910154600191906001600160a01b0316611833828661171d565b52016117e5565b506060925050565b508115611798565b5f525f602052600160405f20015490565b9081602091031261029d575160ff8116810361029d5790565b1561187b57565b6040516323056b1360e01b81526020600482015260056024820152643a37b5b2b760d91b6044820152606490fd5b9081606091031261029d578051801515810361029d57916040602083015192015190565b60ff16604d811161175a57600a0a90565b8181029291811591840414171561175a57565b81156118fb570490565b634e487b7160e01b5f52601260045260245ffd5b60015460405163e588566f60e01b81526001600160a01b038085166004830181905295949260609183916024918391165afa8015611041575f915f916119ee575b5090156119e5575b61196461059b8661260c565b90855f52601060205260405f205480156119dc575b816119c4575050935b5f52600960205260405f2054905f1982145f146119a9576119a591505f92612665565b5091565b81156119ba575b6119a59192612665565b60035491506119b0565b6119d6926119d1916118de565b6118f1565b93611982565b50600554611979565b50600454611958565b9050611a12915060603d606011611a19575b611a0a8183611392565b8101906118a9565b505f611950565b503d611a00565b9190611a2b836126b4565b600654948515801593508390611d70575b82159081611d68575b50611d5e576001600160a01b031692611a6591906105ae61059b8661260c565b936001600160c01b038511611d5557849082611d4b575b5050611d4457600854158015611d3a575b611d3357611a9a83611745565b815f52600b60205260405f205411611d2c57805f52600b60205260405f20611ac384825461177b565b9055805f52600c60205260405f2092610e10420492610e10840293808504610e10149015171561175a57611af96008548561176e565b9160018601945b611b0987612395565b15611d1e578654608081901c6001600160801b039091161461128d5786546001600160801b03165f9081526020879052604090205460c081901c851115611bf0575f868152600b6020526040902054906001600160c01b031680821115611be857611b739161176e565b5f868152600b602052604090205586546001600160801b0381169060801c8114611bd657611ba181886123af565b506001600160801b038181165f9081526020899052604081205588546001600160801b03191660019290920116178755611b00565b634e487b715f5260316020526024601cfd5b50505f611b73565b509092509490949392935b611c0482612395565b151580611cdd575b611c81575b8154608081901c946001600160801b036001870181169592168514611c6f57611c519560018060801b03165f5260205260405f209160c01b1790556127c6565b5f52600b60205260405f205460075410611c6a57600190565b600490565b634e487b715f5260416020526024601cfd5b81549092608082901c916001600160801b03168214611bd6575f199091016001600160801b03165f8181526020869052604081208054919055611cd79291611cc990856127c6565b6001600160c01b031661177b565b91611c11565b508154608081901c6001600160801b039091161461128d57815460801c5f19016001600160801b03165f9081526020859052604090205460c01c8114611c0c565b909250949094939293611bfb565b5060089150565b5060019150565b5060075415611a8d565b5060039150565b109050835f611a7c565b50600793505050565b50600a9450505050565b90505f611a45565b5060085415158015611a3c57506007541515611a3c565b92916001600160a01b031680151580611fcc575b611f8f5750611da9836126b4565b600654948515801593508390611f78575b82159081611f70575b50611d5e576001600160a01b031692611de391906105ae61059b8661260c565b936001600160c01b038511611d5557849082611f66575b5050611d4457600854158015611f5c575b611d3357611e1883611745565b815f52600b60205260405f205411611d2c57805f52600b60205260405f20611e4184825461177b565b9055805f52600c60205260405f2092610e10420492610e10840293808504610e10149015171561175a57611e776008548561176e565b9160018601945b611e8787612395565b15611d1e578654608081901c6001600160801b039091161461128d5786546001600160801b03165f9081526020879052604090205460c081901c851115611bf0575f868152600b6020526040902054906001600160c01b031680821115611f5457611ef19161176e565b5f868152600b602052604090205586546001600160801b0381169060801c8114611bd657611f1f81886123af565b506001600160801b038181165f9081526020899052604081205588546001600160801b03191660019290920116178755611e7e565b50505f611ef1565b5060075415611e0b565b109050835f611dfa565b90505f611dc3565b5060085415158015611dba57506007541515611dba565b60405191825291926001600160a01b0316907fdfe196174d835a3614fc87874a6a97dde9719b32d7d503610251887094dccfce90602090a3600190565b50611fe2815f52600f60205260405f2054151590565b611d9b565b91908110156117315760051b0190565b356001600160a01b038116810361029d5790565b6001600160a01b0391821681529116602082015260400190565b604051636eb1769f60e11b81529791969195919491936001600160a01b0390911692916020898061205a868c6004840161200b565b0381875afa8015611041575f995f91612363575b508281101561235757843b1561029d5760405163d505accf60e01b81526001600160a01b038a81166004830181905290861660248301819052604483018690526064830185905260ff8916608484015260a483018a905260c483018b9052909a9095915f8160e481838c5af180156110415761233a575b50604051636eb1769f60e11b8152916020918391829161210991906004840161200b565b0381895afa90811561232f578b916122fd575b508381036122ee5714612135575b505050505050505050565b8042116122df57604051623f675f60e91b815260048101899052602081602481885afa9081156122d4578a916122a2575b5080156122935791600494939160209360405192858401947f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c986528c6040860152606085015260808401525f190160a083015260c082015260c081526121cd60e082611392565b5190209160405193848092633644e51560e01b82525afa91821561228857879261224f575b5091604261221c949261222596946040519161190160f01b83526002830152602282015220612c32565b90929192612caa565b6001600160a01b03160361224057808080808080808061212a565b6313f3ff5560e31b8152600490fd5b9091506020813d602011612280575b8161226b60209383611392565b8101031261227c57519060426121f2565b8680fd5b3d915061225e565b6040513d89823e3d90fd5b6304de863b60e21b8a5260048afd5b90506020813d6020116122cc575b816122bd60209383611392565b8101031261029d57515f612166565b3d91506122b0565b6040513d8c823e3d90fd5b63219386ff60e11b8952600489fd5b6304de863b60e21b8b5260048bfd5b90506020813d602011612327575b8161231860209383611392565b8101031261029d57515f61211c565b3d915061230b565b6040513d8d823e3d90fd5b6020919c509161234d5f61210994611392565b5f9c9150916120e5565b50505050505050505050565b90506020813d60201161238d575b8161237e60209383611392565b8101031261029d57515f61206e565b3d9150612371565b546001600160801b0380821660809290921c919091031690565b9060018060801b03165f5260205260405f2090565b335f9081527fa74d116c0406588b87fcebe22b8d03a12bf1aa08ea523c2004e00dd791abafe3602052604090205460ff16156123fc57565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b335f9081527fe70f4a856a97bd1ebdcb74b506f0a385e96f16da6133e1d2983a79b65bc91dc0602052604090205460ff161561246b57565b63e2517d3f60e01b5f52336004527fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a60245260445ffd5b335f9081527f7dcd276b72586271b4623598c6b3ecb5f3594881d1da5ce0d28e8f8adbf67787602052604090205460ff16156124da57565b63e2517d3f60e01b5f52336004527f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f60245260445ffd5b335f9081527f7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ec602052604090205460ff161561254957565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b5f8181526020818152604080832033845290915290205460ff16156125a25750565b63e2517d3f60e01b5f523360045260245260445ffd5b9190916125c58382612886565b92836125cf575050565b815f52600d6020526125ee60405f209160018060a01b031680926129db565b156125f7575050565b639d98e84160e01b5f5260045260245260445ffd5b6001600160a01b0316600181036126235750601290565b60206004916040519283809263313ce56760e01b82525afa908115611041575f9161264c575090565b61150b915060203d60201161103a5761102c8183611392565b90815f52600a60205260405f205480156126ab576126a49261269460018060a01b0360015416926002546118de565b926001600160a01b031691612a8c565b9091509091565b5050505f905f90565b60015460405163e588566f60e01b81526001600160a01b039283166004820152929160609184916024918391165afa8015611041575f925f91612703575b509082156126fc57565b6004549150565b905061271f91925060603d606011611a1957611a0a8183611392565b509190915f6126f2565b91818302915f19818509938380861095039480860395146127b957848311156127a15790829109815f0382168092046002816003021880820260020302808202600203028082026002030280820260020302808202600203028091026002030293600183805f03040190848311900302920304170290565b82634e487b715f52156003026011186020526024601cfd5b50509061150b92506118f1565b80546001600160801b031660809290921b6001600160801b031916919091179055565b8054821015611731575f5260205f2001905f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16612880575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615612880575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b5f818152600f60205260409020548015612880575f19810181811161175a57600e545f1981019190821161175a5781810361298d575b505050600e548015612979575f190161295681600e6127e9565b8154905f199060031b1b19169055600e555f52600f6020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b6129c561299e6129af93600e6127e9565b90549060031b1c928392600e6127e9565b819391549060031b91821b915f19901b19161790565b90555f52600f60205260405f20555f808061293c565b906001820191815f528260205260405f20548015155f14612a84575f19810181811161175a5782545f1981019190821161175a57818103612a4f575b50505080548015612979575f190190612a3082826127e9565b8154905f199060031b1b19169055555f526020525f6040812055600190565b612a6f612a5f6129af93866127e9565b90549060031b1c928392866127e9565b90555f528360205260405f20555f8080612a17565b505050505f90565b604051630ebec55160e21b81526004810192909252909392906001600160a01b0316606085602481845afa948515611041575f905f96612b6b575b5015612b5e5760405163e588566f60e01b81526001600160a01b038416600482015290606090829060249082905afa948515611041575f955f925f91612b37575b50938615612b285790612b1f612b2594939261260c565b92612d1e565b91565b505050505090505f905f905f90565b91965050612b54915060603d606011611a1957611a0a8183611392565b919590915f612b08565b50505090505f905f905f90565b9050612b8791955060603d606011611a1957611a0a8183611392565b95919095612ac7565b805f52600f60205260405f2054155f14612be057600e54600160401b8110156113b557612bc96129af826001859401600e55600e6127e9565b9055600e54905f52600f60205260405f2055600190565b505f90565b5f82815260018201602052604090205461288057805490600160401b8210156113b55782612c1d6129af8460018096018555846127e9565b90558054925f520160205260405f2055600190565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411612c9f579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15611041575f516001600160a01b03811615612c9557905f905f90565b505f906001905f90565b5050505f9160039190565b60048110156116a15780612cbc575050565b60018103612cd35763f645eedf60e01b5f5260045ffd5b60028103612cee575063fce698f760e01b5f5260045260245ffd5b600314612cf85750565b6335e2f38360e21b5f5260045260245ffd5b9060ff8091169116039060ff821161175a57565b91928115612db3578315612d845760ff811680601214612d785760121015612d635761150b939291612d5761059b6012612d5d94612d0a565b906118de565b90612729565b61150b93612d5761059b6105ae936012612d0a565b505061150b9291612729565b6040516349671cd160e11b8152602060048201526006602482015265383934b1b2a160d11b6044820152606490fd5b6040516349671cd160e11b815260206004820152600660248201526570726963654160d01b6044820152606490fdfe4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3bb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fde1d1777e4c7896f6e13f47c189990ed90f2210cb88f3836cd7bd0cb1f820945b386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1dcf6ae774f84d549d12db18368181005de6e396adc8740a42b1df7a36352edad3d0dc86d2f8e36fe5e3affad59f28b7616d50e357ee97837a5e2b97024f5a2f05a264697066735822122010342a39df83700924a119c5a367fcad9a06673a92f3e43c368c5b0a14d0f86164736f6c634300081c003352ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f",
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

// GetMinimumTokenValueOf is a free data retrieval call binding the contract method 0xc5fe3bf4.
//
// Solidity: function getMinimumTokenValueOf(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCaller) GetMinimumTokenValueOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getMinimumTokenValueOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTokenValueOf is a free data retrieval call binding the contract method 0xc5fe3bf4.
//
// Solidity: function getMinimumTokenValueOf(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierSession) GetMinimumTokenValueOf(token common.Address) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetMinimumTokenValueOf(&_BridgeVerifier.CallOpts, token)
}

// GetMinimumTokenValueOf is a free data retrieval call binding the contract method 0xc5fe3bf4.
//
// Solidity: function getMinimumTokenValueOf(address token) view returns(uint256)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetMinimumTokenValueOf(token common.Address) (*big.Int, error) {
	return _BridgeVerifier.Contract.GetMinimumTokenValueOf(&_BridgeVerifier.CallOpts, token)
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

// SetMinimumTokenValueOf is a paid mutator transaction binding the contract method 0x3d949a66.
//
// Solidity: function setMinimumTokenValueOf(address token, uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetMinimumTokenValueOf(opts *bind.TransactOpts, token common.Address, minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setMinimumTokenValueOf", token, minimumTokenValue)
}

// SetMinimumTokenValueOf is a paid mutator transaction binding the contract method 0x3d949a66.
//
// Solidity: function setMinimumTokenValueOf(address token, uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetMinimumTokenValueOf(token common.Address, minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValueOf(&_BridgeVerifier.TransactOpts, token, minimumTokenValue)
}

// SetMinimumTokenValueOf is a paid mutator transaction binding the contract method 0x3d949a66.
//
// Solidity: function setMinimumTokenValueOf(address token, uint256 minimumTokenValue) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetMinimumTokenValueOf(token common.Address, minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValueOf(&_BridgeVerifier.TransactOpts, token, minimumTokenValue)
}

// SetMinimumTokenValueOfBatch is a paid mutator transaction binding the contract method 0xedba2c77.
//
// Solidity: function setMinimumTokenValueOfBatch(address[] tokenList, uint256[] minimumTokenValueList) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetMinimumTokenValueOfBatch(opts *bind.TransactOpts, tokenList []common.Address, minimumTokenValueList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setMinimumTokenValueOfBatch", tokenList, minimumTokenValueList)
}

// SetMinimumTokenValueOfBatch is a paid mutator transaction binding the contract method 0xedba2c77.
//
// Solidity: function setMinimumTokenValueOfBatch(address[] tokenList, uint256[] minimumTokenValueList) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetMinimumTokenValueOfBatch(tokenList []common.Address, minimumTokenValueList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValueOfBatch(&_BridgeVerifier.TransactOpts, tokenList, minimumTokenValueList)
}

// SetMinimumTokenValueOfBatch is a paid mutator transaction binding the contract method 0xedba2c77.
//
// Solidity: function setMinimumTokenValueOfBatch(address[] tokenList, uint256[] minimumTokenValueList) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetMinimumTokenValueOfBatch(tokenList []common.Address, minimumTokenValueList []*big.Int) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetMinimumTokenValueOfBatch(&_BridgeVerifier.TransactOpts, tokenList, minimumTokenValueList)
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

// BridgeVerifierMinimumTokenValueOfSetIterator is returned from FilterMinimumTokenValueOfSet and is used to iterate over the raw logs and unpacked data for MinimumTokenValueOfSet events raised by the BridgeVerifier contract.
type BridgeVerifierMinimumTokenValueOfSetIterator struct {
	Event *BridgeVerifierMinimumTokenValueOfSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierMinimumTokenValueOfSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierMinimumTokenValueOfSet)
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
		it.Event = new(BridgeVerifierMinimumTokenValueOfSet)
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
func (it *BridgeVerifierMinimumTokenValueOfSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierMinimumTokenValueOfSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierMinimumTokenValueOfSet represents a MinimumTokenValueOfSet event raised by the BridgeVerifier contract.
type BridgeVerifierMinimumTokenValueOfSet struct {
	Token             common.Address
	MinimumTokenValue *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMinimumTokenValueOfSet is a free log retrieval operation binding the contract event 0xd0dc86d2f8e36fe5e3affad59f28b7616d50e357ee97837a5e2b97024f5a2f05.
//
// Solidity: event MinimumTokenValueOfSet(address indexed token, uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterMinimumTokenValueOfSet(opts *bind.FilterOpts, token []common.Address) (*BridgeVerifierMinimumTokenValueOfSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "MinimumTokenValueOfSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierMinimumTokenValueOfSetIterator{contract: _BridgeVerifier.contract, event: "MinimumTokenValueOfSet", logs: logs, sub: sub}, nil
}

// WatchMinimumTokenValueOfSet is a free log subscription operation binding the contract event 0xd0dc86d2f8e36fe5e3affad59f28b7616d50e357ee97837a5e2b97024f5a2f05.
//
// Solidity: event MinimumTokenValueOfSet(address indexed token, uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchMinimumTokenValueOfSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierMinimumTokenValueOfSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "MinimumTokenValueOfSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierMinimumTokenValueOfSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "MinimumTokenValueOfSet", log); err != nil {
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

// ParseMinimumTokenValueOfSet is a log parse operation binding the contract event 0xd0dc86d2f8e36fe5e3affad59f28b7616d50e357ee97837a5e2b97024f5a2f05.
//
// Solidity: event MinimumTokenValueOfSet(address indexed token, uint256 minimumTokenValue)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseMinimumTokenValueOfSet(log types.Log) (*BridgeVerifierMinimumTokenValueOfSet, error) {
	event := new(BridgeVerifierMinimumTokenValueOfSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "MinimumTokenValueOfSet", log); err != nil {
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
