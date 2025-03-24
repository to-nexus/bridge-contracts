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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPeriodTotalValueThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCurrentScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMovementHistory\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"history\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getTokenPriceWithValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerificationAmountThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"setDefaultTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setExFeeRateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"setPeriodTotalValueThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasPrices\",\"type\":\"uint256[]\"}],\"name\":\"updateGasPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"DefaultTokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"ExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"FinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"GasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"PeriodTotalValueThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"TimeWindowSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"VerificationAmountThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeVerifierCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeVerifierChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"BridgeVerifierDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierMissmatchLength\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
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
		"be716b0a": "getVerificationAmountThreshold()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"b2b49e2e": "grantRoleBatch(bytes32[],address[])",
		"91d14854": "hasRole(bytes32,address)",
		"741bef1a": "priceFeed()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
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
	},
	Bin: "0x608060405234801561000f575f5ffd5b506040516126a03803806126a083398101604081905261002e9161039b565b865f03610077576040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b60448201526064015b60405180910390fd5b6001600160a01b038a166100bd576040516323056b1360e01b815260206004820152600c60248201526b34b734ba34b0b627bbb732b960a11b604482015260640161006e565b6001600160a01b0389166100fd576040516323056b1360e01b815260206004820152600660248201526562726964676560d01b604482015260640161006e565b6001600160a01b038816610141576040516323056b1360e01b815260206004820152600a60248201526917dc1c9a58d95199595960b21b604482015260640161006e565b61014b5f8b610214565b506101767f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c8b610214565b506101a17fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a8b610214565b506101cc7f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f8a610214565b50600180546001600160a01b0319166001600160a01b039990991698909817909755600295909555600493909355600391909155600555600655600755600855506104419050565b5f61021f8383610279565b90508015610273575f838152600d6020526040902061023e9083610320565b8284909161027057604051637a854c7360e11b81526001600160a01b039092166004830152602482015260440161006e565b50505b92915050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610319575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556102d13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610273565b505f610273565b5f610334836001600160a01b03841661033b565b9392505050565b5f81815260018301602052604081205461031957508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610273565b80516001600160a01b0381168114610396575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f5f6101408b8d0312156103b5575f5ffd5b6103be8b610380565b99506103cc60208c01610380565b98506103da60408c01610380565b97505f60608c01519050809750505f60808c01519050809650505f60a08c01519050809550505f60c08c01519050809450505f60e08c01519050809350505f6101008c01519050809250505f6101208c01519050809150509295989b9194979a5092959850565b6122528061044e5f395ff3fe608060405234801561000f575f5ffd5b506004361061021e575f3560e01c80638d3319961161012a578063ad1434ca116100b4578063d02641a011610079578063d02641a0146104ee578063d1c0154314610501578063d547741f14610520578063df6e87dc14610533578063f289d3ba14610546575f5ffd5b8063ad1434ca14610483578063b2b49e2e14610496578063ba8d25f8146104a9578063be716b0a146104d3578063c3de108b146104db575f5ffd5b8063979eb82d116100fa578063979eb82d1461042e5780639dc696cc14610436578063a217fddf14610449578063a3246ad314610450578063ab509f3a14610470575f5ffd5b80638d331996146103d857806391d14854146103eb57806396ce0795146103fe5780639745e4ba14610406575f5ffd5b8063365d71e4116101ab5780635ec8f51b1161017b5780635ec8f51b146103465780636332aec614610359578063724e78da14610387578063741bef1a1461039a578063751ccd79146103c5575f5ffd5b8063365d71e4146102f85780633f4bdec51461030b57806349c9215e1461032b5780635566ca091461033e575f5ffd5b8063248a9ca3116101f1578063248a9ca3146102925780632a4d2f32146102c25780632f2ff15d146102ca5780632fe3b6cf146102dd57806336568abe146102e5575f5ffd5b806301ffc9a714610222578063020a22121461024a5780630e4e96bf1461025f5780631f96131e1461027f575b5f5ffd5b610235610230366004611b49565b610559565b60405190151581526020015b60405180910390f35b61025d610258366004611c56565b61058f565b005b61027261026d366004611d1b565b610605565b6040516102419190611d36565b61025d61028d366004611d78565b6106c7565b6102b46102a0366004611d98565b5f9081526020819052604090206001015490565b604051908152602001610241565b6005546102b4565b61025d6102d8366004611daf565b610742565b6102b461076c565b61025d6102f3366004611daf565b6107df565b61025d610306366004611e41565b610812565b61031e610319366004611ef3565b610883565b6040516102419190611f1d565b61025d610339366004611d98565b610ad6565b6008546102b4565b61025d610354366004611ef3565b610b2a565b61036c610367366004611daf565b610bd1565b60408051938452602084019290925290820152606001610241565b61025d610395366004611d1b565b610d59565b6001546103ad906001600160a01b031681565b6040516001600160a01b039091168152602001610241565b61025d6103d3366004611f43565b610df1565b61025d6103e6366004611d98565b610f04565b6102356103f9366004611daf565b610f43565b6103e86102b4565b6102b4610414366004611d1b565b6001600160a01b03165f908152600b602052604090205490565b6007546102b4565b61025d610444366004611d98565b610f6b565b6102b45f81565b61046361045e366004611d98565b610fb7565b6040516102419190611f92565b61025d61047e366004611d98565b610fd0565b61025d610491366004611d98565b61100f565b61025d6104a4366004611e41565b611064565b6104bc6104b7366004611ef3565b6110d5565b604080519215158352602083019190915201610241565b6006546102b4565b61025d6104e9366004611d98565b6111f0565b6104bc6104fc366004611d1b565b61122f565b6102b461050f366004611d98565b5f908152600a602052604090205490565b61025d61052e366004611daf565b6112c7565b61036c610541366004611fd2565b6112eb565b61025d610554366004611d98565b611323565b5f6001600160e01b03198216637965db0b60e01b148061058957506301ffc9a760e01b6001600160e01b03198316145b92915050565b80518251146105b15760405163322473c360e01b815260040160405180910390fd5b5f5b8251811015610600576105f88382815181106105d1576105d1612007565b60200260200101518383815181106105eb576105eb612007565b6020026020010151610b2a565b6001016105b3565b505050565b6001600160a01b0381165f908152600c6020526040812060609190610629906113a6565b90508067ffffffffffffffff81111561064457610644611b70565b60405190808252806020026020018201604052801561066d578160200160208202803683370190505b5091505f5b818110156106c0576001600160a01b0384165f908152600c6020526040902061069b90826113c4565b8382815181106106ad576106ad612007565b6020908102919091010152600101610672565b5050919050565b7fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a6106f181611405565b5f838152600a6020526040908190208390555183907f4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3906107359085815260200190565b60405180910390a2505050565b5f8281526020819052604090206001015461075c81611405565b6107668383611412565b50505050565b60015460408051632fe3b6cf60e01b815290515f926001600160a01b031691632fe3b6cf9160048083019260209291908290030181865afa1580156107b3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107d7919061201b565b60ff16905090565b6001600160a01b03811633146108085760405163334bd91960e11b815260040160405180910390fd5b6106008282611476565b8051825114610834576040516374ddb37160e01b815260040160405180910390fd5b5f5b81518110156106005761087b83828151811061085457610854612007565b602002602001015183838151811061086e5761086e612007565b60200260200101516112c7565b600101610836565b5f7f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f6108ae81611405565b5f6108b985856110d5565b9150506001600160c01b038111156108d5576007925050610acf565b600654158015906108e7575080600654105b156108f6576003925050610acf565b6001600160a01b0385165f908152600c602052604081209061091a610e1042612063565b90505f610e106008548361092e9190612082565b6109389190612063565b90505b5f610945846113a6565b11156109d5575f610955846114d2565b905060c081901c828110156109c7576001600160a01b038a165f908152600b60205260409020546001600160c01b03831690818111610994575f61099e565b61099e8282612082565b6001600160a01b038d165f908152600b60205260409020556109bf8761151f565b5050506109ce565b50506109d5565b505061093b565b6001600160a01b0388165f908152600b60205260409020546109fe856001600160c01b03612082565b1015610a11576008955050505050610acf565b6001600160a01b0388165f908152600b602052604081208054869290610a38908490612095565b909155505060075415801590610a6657506007546001600160a01b0389165f908152600b6020526040902054115b15610a78576004955050505050610acf565b610a81836113a6565b15801590610a9a57508160c0610a968561158c565b901c145b15610ab4576001600160c01b03610ab0846115e6565b1693505b60c082901b8417610ac58482611644565b6001965050505050505b5092915050565b5f5160206121fd5f395f51905f52610aed81611405565b60048290556040518281527f1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a906020015b60405180910390a15050565b5f5160206121fd5f395f51905f52610b4181611405565b6001600160a01b038316610b85576040516323056b1360e01b81526020600482015260056024820152643a37b5b2b760d91b60448201526064015b60405180910390fd5b6001600160a01b0383165f8181526009602052604090819020849055517f386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1d906107359085815260200190565b60015460405163e588566f60e01b81526001600160a01b0383811660048301525f92839283928392839291169063e588566f90602401606060405180830381865afa158015610c22573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c4691906120a8565b509150915081610c5557506004545b604051636a24d41960e11b81526001600160a01b03871660048201525f9073__$50392a774d9f10cbbfcb17b0a031deadbc$__9063d449a83290602401602060405180830381865af4158015610cad573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cd1919061201b565b610cdc90600a6121c3565b9050815f03610ced57809550610d09565b8181600554610cfc91906121d1565b610d069190612063565b95505b6001600160a01b0387165f90815260096020526040902054935060018401610d33575f9350610d40565b835f03610d405760035493505b610d4a88886116a5565b50809550505050509250925092565b5f610d6381611405565b6001600160a01b038216610da7576040516323056b1360e01b815260206004820152600a60248201526917dc1c9a58d95199595960b21b6044820152606401610b7c565b600180546001600160a01b0319166001600160a01b0384169081179091556040517fe5b20b8497e4f3e2435ef9c20e2e26b47497ee13745ce1c681ad6640653119e6905f90a25050565b7fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a610e1b81611405565b8151835114610e3d5760405163322473c360e01b815260040160405180910390fd5b5f5b835181101561076657828181518110610e5a57610e5a612007565b6020026020010151600a5f868481518110610e7757610e77612007565b602002602001015181526020019081526020015f2081905550838181518110610ea257610ea2612007565b60200260200101517f4afe44d7dc150ca7bf2407ab84e694efb1898578aa5008d8d4265039f411e4c3848381518110610edd57610edd612007565b6020026020010151604051610ef491815260200190565b60405180910390a2600101610e3f565b5f610f0e81611405565b60088290556040518281527f52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a490602001610b1e565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f5160206121fd5f395f51905f52610f8281611405565b60058290556040518281527f0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e990602001610b1e565b5f818152600d6020526040902060609061058990611791565b5f610fda81611405565b60078290556040518281527f7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e290602001610b1e565b5f5160206121fd5f395f51905f5261102681611405565b60038290556040518281525f907f386d4c8c407b3a159372ee393f741c507db421400eb3e74d287584f7c0d3da1d9060200160405180910390a25050565b8051825114611086576040516374ddb37160e01b815260040160405180910390fd5b5f5b8151811015610600576110cd8382815181106110a6576110a6612007565b60200260200101518383815181106110c0576110c0612007565b6020026020010151610742565b600101611088565b60015460405163e588566f60e01b81526001600160a01b0384811660048301525f92839291169063e588566f90602401606060405180830381865afa158015611120573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061114491906120a8565b5090925090508161115457506004545b604051636a24d41960e11b81526001600160a01b03851660048201526111e790849073__$50392a774d9f10cbbfcb17b0a031deadbc$__9063d449a83290602401602060405180830381865af41580156111b0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111d4919061201b565b6111df90600a6121c3565b8391906117a4565b90509250929050565b5f6111fa81611405565b60068290556040518281527ff905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a290602001610b1e565b604051636a24d41960e11b81526001600160a01b03821660048201525f9081906112be90849073__$50392a774d9f10cbbfcb17b0a031deadbc$__9063d449a83290602401602060405180830381865af415801561128f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112b3919061201b565b6104b790600a6121c3565b91509150915091565b5f828152602081905260409020600101546112e181611405565b6107668383611476565b5f5f5f5f6112f98787610bd1565b919550935090506103e861130d82876121d1565b6113179190612063565b91505093509350939050565b5f61132d81611405565b815f03611371576040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606401610b7c565b60028290556040518281527fbd4381b5028f9389a5ec0a63a18bae9a1aaee356f563a421c8edd2c09f530a9b90602001610b1e565b546001600160801b03808216600160801b9092048116919091031690565b5f6113ce836113a6565b82106113de576113de603261185a565b5081546001600160801b039081168201165f90815260018301602052604090205492915050565b61140f813361186b565b50565b5f61141d83836118a8565b90508015610589575f838152600d6020526040902061143c9083611937565b8284909161146e57604051637a854c7360e11b81526001600160a01b0390921660048301526024820152604401610b7c565b505092915050565b5f611481838361194b565b90508015610589575f838152600d602052604090206114a090836119b4565b8284909161146e57604051639d98e84160e01b81526001600160a01b0390921660048301526024820152604401610b7c565b5f6114f08254600160801b81046001600160801b0390811691161490565b156114ff576114ff603261185a565b5080546001600160801b03165f9081526001909101602052604090205490565b80545f906001600160801b0380821691600160801b900416810361154757611547603161185a565b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6115aa8254600160801b81046001600160801b0390811691161490565b156115b9576115b9603261185a565b5080545f196001600160801b03600160801b909204821601165f9081526001909101602052604090205490565b80545f906001600160801b03600160801b820481169116810361160d5761160d603161185a565b5f19016001600160801b039081165f818152600185016020526040812080549190558454909216600160801b909102179092555090565b81546001600160801b03600160801b82048116918116600183019091160361167057611670604161185a565b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f828152600a602052604081205481908082036116c8575f5f925092505061178a565b60015460025473__$50392a774d9f10cbbfcb17b0a031deadbc$__916381d889b4916001600160a01b0390911690889088906117059087906121d1565b6040516001600160e01b031960e087901b1681526001600160a01b0394851660048201526024810193909352921660448201526064810191909152608401606060405180830381865af415801561175e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061178291906120a8565b909450925050505b9250929050565b60605f61179d836119c8565b9392505050565b5f838302815f1985870982811083820303915050805f036117d8578382816117ce576117ce61203b565b049250505061179d565b8084116117ef576117ef600385150260111861185a565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6118758282610f43565b6118a45760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610b7c565b5050565b5f6118b38383610f43565b611930575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556118e83390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610589565b505f610589565b5f61179d836001600160a01b038416611a21565b5f6119568383610f43565b15611930575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610589565b5f61179d836001600160a01b038416611a66565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611a1557602002820191905f5260205f20905b815481526020019060010190808311611a01575b50505050509050919050565b5f81815260018301602052604081205461193057508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610589565b5f8181526001830160205260408120548015611b40575f611a88600183612082565b85549091505f90611a9b90600190612082565b9050808214611afa575f865f018281548110611ab957611ab9612007565b905f5260205f200154905080875f018481548110611ad957611ad9612007565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611b0b57611b0b6121e8565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610589565b5f915050610589565b5f60208284031215611b59575f5ffd5b81356001600160e01b03198116811461179d575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611bad57611bad611b70565b604052919050565b5f67ffffffffffffffff821115611bce57611bce611b70565b5060051b60200190565b6001600160a01b038116811461140f575f5ffd5b5f82601f830112611bfb575f5ffd5b8135611c0e611c0982611bb5565b611b84565b8082825260208201915060208360051b860101925085831115611c2f575f5ffd5b602085015b83811015611c4c578035835260209283019201611c34565b5095945050505050565b5f5f60408385031215611c67575f5ffd5b823567ffffffffffffffff811115611c7d575f5ffd5b8301601f81018513611c8d575f5ffd5b8035611c9b611c0982611bb5565b8082825260208201915060208360051b850101925087831115611cbc575f5ffd5b6020840193505b82841015611ce7578335611cd681611bd8565b825260209384019390910190611cc3565b9450505050602083013567ffffffffffffffff811115611d05575f5ffd5b611d1185828601611bec565b9150509250929050565b5f60208284031215611d2b575f5ffd5b813561179d81611bd8565b602080825282518282018190525f918401906040840190835b81811015611d6d578351835260209384019390920191600101611d4f565b509095945050505050565b5f5f60408385031215611d89575f5ffd5b50508035926020909101359150565b5f60208284031215611da8575f5ffd5b5035919050565b5f5f60408385031215611dc0575f5ffd5b823591506020830135611dd281611bd8565b809150509250929050565b5f82601f830112611dec575f5ffd5b8135611dfa611c0982611bb5565b8082825260208201915060208360051b860101925085831115611e1b575f5ffd5b602085015b83811015611c4c578035611e3381611bd8565b835260209283019201611e20565b5f5f60408385031215611e52575f5ffd5b823567ffffffffffffffff811115611e68575f5ffd5b8301601f81018513611e78575f5ffd5b8035611e86611c0982611bb5565b8082825260208201915060208360051b850101925087831115611ea7575f5ffd5b6020840193505b82841015611ec9578335825260209384019390910190611eae565b9450505050602083013567ffffffffffffffff811115611ee7575f5ffd5b611d1185828601611ddd565b5f5f60408385031215611f04575f5ffd5b8235611f0f81611bd8565b946020939093013593505050565b60208101600a8310611f3d57634e487b7160e01b5f52602160045260245ffd5b91905290565b5f5f60408385031215611f54575f5ffd5b823567ffffffffffffffff811115611f6a575f5ffd5b611f7685828601611bec565b925050602083013567ffffffffffffffff811115611d05575f5ffd5b602080825282518282018190525f918401906040840190835b81811015611d6d5783516001600160a01b0316835260209384019390920191600101611fab565b5f5f5f60608486031215611fe4575f5ffd5b833592506020840135611ff681611bd8565b929592945050506040919091013590565b634e487b7160e01b5f52603260045260245ffd5b5f6020828403121561202b575f5ffd5b815160ff8116811461179d575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f8261207d57634e487b7160e01b5f52601260045260245ffd5b500490565b818103818111156105895761058961204f565b808201808211156105895761058961204f565b5f5f5f606084860312156120ba575f5ffd5b835180151581146120c9575f5ffd5b602085015160409095015190969495509392505050565b6001815b600184111561211b578085048111156120ff576120ff61204f565b600184161561210d57908102905b60019390931c9280026120e4565b935093915050565b5f8261213157506001610589565b8161213d57505f610589565b8160018114612153576002811461215d57612179565b6001915050610589565b60ff84111561216e5761216e61204f565b50506001821b610589565b5060208310610133831016604e8410600b841016171561219c575081810a610589565b6121a85f1984846120e0565b805f19048211156121bb576121bb61204f565b029392505050565b5f61179d60ff841683612123565b80820281158282048414176105895761058961204f565b634e487b7160e01b5f52603160045260245ffdfe21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9ca26469706673582212206e75e023bb84660c75398bdd57748906513adfab0a4b926ce9e33c71eed8497b64736f6c634300081c0033",
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
func DeployBridgeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, bridge common.Address, _priceFeed common.Address, finalizeBridgeGas *big.Int, defaultTokenPrice *big.Int, defaultExFeeRate *big.Int, minimumTokenValue *big.Int, verificationAmountThreshold *big.Int, periodTotalValueThreshold *big.Int, timeWindow *big.Int) (common.Address, *types.Transaction, *BridgeVerifier, error) {
	parsed, err := BridgeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeVerifierBin), backend, initialOwner, bridge, _priceFeed, finalizeBridgeGas, defaultTokenPrice, defaultExFeeRate, minimumTokenValue, verificationAmountThreshold, periodTotalValueThreshold, timeWindow)
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
