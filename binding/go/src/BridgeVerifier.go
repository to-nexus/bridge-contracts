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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPeriodTotalValueThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCurrentScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMovementHistory\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"history\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getTokenPriceWithValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerificationAmountThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"setDefaultTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setExFeeRateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"setPeriodTotalValueThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasPrices\",\"type\":\"uint256[]\"}],\"name\":\"updateGasPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"BridgeVerifierExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"BridgeVerifierFinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"BridgeVerifierGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"BridgeVerifierPriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"DefaultTokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"PeriodTotalValueThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"TimeWindowSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"VerificationAmountThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeVerifierCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeVerifierChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeVerifierInvalidLength\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"96ce0795": "denominator()",
		"d1c01543": "getGasPrice(uint256)",
		"2a4d2f32": "getMinimumTokenValue()",
		"979eb82d": "getPeriodTotalValueThreshold()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"5566ca09": "getTimeWindow()",
		"6332aec6": "getTokenConfig(uint256,address)",
		"9745e4ba": "getTokenCurrentScore(address)",
		"0e4e96bf": "getTokenMovementHistory(address)",
		"d02641a0": "getTokenPrice(address)",
		"ba8d25f8": "getTokenPriceWithValue(address,uint256)",
		"be716b0a": "getVerificationAmountThreshold()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"ce2e5c66": "removePriceFeed()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
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
	Bin: "0x608060405234801561000f575f5ffd5b50604051611f6b380380611f6b83398101604081905261002e9161021a565b865f03610077576040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b60448201526064015b60405180910390fd5b6001600160a01b0389166100bd576040516323056b1360e01b815260206004820152600c60248201526b34b734ba34b0b627bbb732b960a11b604482015260640161006e565b6001600160a01b0388166100fe576040516323056b1360e01b81526020600482015260076024820152665f62726964676560c81b604482015260640161006e565b6101085f8a610156565b5061011b6420a226a4a760d91b8a610156565b5061012f6542524944474560d01b89610156565b50600296909655600494909455600392909255600555600655600855600755506102ae9050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166101f6575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101ae3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101f9565b505f5b92915050565b80516001600160a01b0381168114610215575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610233575f5ffd5b61023c8a6101ff565b985061024a60208b016101ff565b97505f60408b01519050809750505f60608b01519050809650505f60808b01519050809550505f60a08b01519050809450505f60c08b01519050809350505f60e08b01519050809250505f6101008b01519050809150509295985092959850929598565b611cb0806102bb5f395ff3fe608060405234801561000f575f5ffd5b50600436106101f2575f3560e01c806391d1485411610114578063ba8d25f8116100a9578063d02641a011610079578063d02641a014610451578063d1c0154314610464578063d547741f14610483578063df6e87dc14610496578063f289d3ba146104a9575f5ffd5b8063ba8d25f814610404578063be716b0a1461042e578063c3de108b14610436578063ce2e5c6614610449575f5ffd5b80639dc696cc116100e45780639dc696cc146103c4578063a217fddf146103d7578063ab509f3a146103de578063ad1434ca146103f1575f5ffd5b806391d148541461037957806396ce07951461038c5780639745e4ba14610394578063979eb82d146103bc575f5ffd5b80633f4bdec51161018a5780636332aec61161015a5780636332aec614610312578063724e78da14610340578063751ccd79146103535780638d33199614610366575f5ffd5b80633f4bdec5146102c457806349c9215e146102e45780635566ca09146102f75780635ec8f51b146102ff575f5ffd5b8063248a9ca3116101c5578063248a9ca3146102665780632a4d2f32146102965780632f2ff15d1461029e57806336568abe146102b1575f5ffd5b806301ffc9a7146101f6578063020a22121461021e5780630e4e96bf146102335780631f96131e14610253575b5f5ffd5b610209610204366004611731565b6104bc565b60405190151581526020015b60405180910390f35b61023161022c36600461183e565b6104f2565b005b610246610241366004611903565b61057b565b604051610215919061191e565b610231610261366004611960565b61063d565b610288610274366004611980565b5f9081526020819052604090206001015490565b604051908152602001610215565b600554610288565b6102316102ac366004611997565b6106a2565b6102316102bf366004611997565b6106c6565b6102d76102d23660046119c5565b6106fe565b60405161021591906119ef565b6102316102f2366004611980565b610922565b600754610288565b61023161030d3660046119c5565b610971565b610325610320366004611997565b610a13565b60408051938452602084019290925290820152606001610215565b61023161034e366004611903565b610bc1565b610231610361366004611a15565b610c60565b610231610374366004611980565b610d5d565b610209610387366004611997565b610da4565b6103e8610288565b6102886103a2366004611903565b6001600160a01b03165f908152600b602052604090205490565b600854610288565b6102316103d2366004611980565b610dcc565b6102885f81565b6102316103ec366004611980565b610e13565b6102316103ff366004611980565b610e5a565b6104176104123660046119c5565b610eaa565b604080519215158352602083019190915201610215565b600654610288565b610231610444366004611980565b610fdf565b610231611026565b61041761045f366004611903565b611075565b610288610472366004611980565b5f908152600a602052604090205490565b610231610491366004611997565b61110d565b6103256104a4366004611a64565b611131565b6102316104b7366004611980565b611169565b5f6001600160e01b03198216637965db0b60e01b14806104ec57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6420a226a4a760d91b610504816111f4565b81518351146105265760405163322473c360e01b815260040160405180910390fd5b5f5b83518110156105755761056d84828151811061054657610546611a99565b602002602001015184838151811061056057610560611a99565b6020026020010151610971565b600101610528565b50505050565b6001600160a01b0381165f908152600c602052604081206060919061059f90611201565b90508067ffffffffffffffff8111156105ba576105ba611758565b6040519080825280602002602001820160405280156105e3578160200160208202803683370190505b5091505f5b81811015610636576001600160a01b0384165f908152600c60205260409020610611908261121f565b83828151811061062357610623611a99565b60209081029190910101526001016105e8565b5050919050565b662aa82220aa27a960c91b610651816111f4565b5f838152600a6020526040908190208390555183907f3fc45c54c1bb058234a4c562d7bcc2bb829ce9f5842d6d0bcab221506f10a320906106959085815260200190565b60405180910390a2505050565b5f828152602081905260409020600101546106bc816111f4565b6105758383611260565b6001600160a01b03811633146106ef5760405163334bd91960e11b815260040160405180910390fd5b6106f982826112ef565b505050565b5f6542524944474560d01b610712816111f4565b5f61071d8585610eaa565b9150506006545f14158015610733575080600654105b1561074257600392505061091b565b6001600160a01b0385165f908152600c6020526040812090610766610e1042611ad5565b90505f610e106007548361077a9190611af4565b6107849190611ad5565b90505b5f61079184611201565b1115610821575f6107a184611358565b905060c081901c82811015610813576001600160a01b038a165f908152600b60205260409020546001600160c01b038316908181116107e0575f6107ea565b6107ea8282611af4565b6001600160a01b038d165f908152600b602052604090205561080b876113a5565b50505061081a565b5050610821565b5050610787565b6001600160a01b0388165f908152600b602052604090205461084a856001600160c01b03611af4565b101561085d57600595505050505061091b565b6008541580159061088657506008546001600160a01b0389165f908152600b6020526040902054115b1561089857600495505050505061091b565b6001600160a01b0388165f908152600b6020526040812080548692906108bf908490611b07565b909155506108ce905083611201565b158015906108e757508160c06108e385611412565b901c145b15610901576001600160c01b036108fd8461146c565b1693505b60c082901b841761091284826114ca565b5f965050505050505b5092915050565b6420a226a4a760d91b610934816111f4565b60048290556040518281527f1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a906020015b60405180910390a15050565b6420a226a4a760d91b610983816111f4565b6001600160a01b0383166109c7576040516323056b1360e01b81526020600482015260056024820152643a37b5b2b760d91b60448201526064015b60405180910390fd5b6001600160a01b0383165f8181526009602052604090819020849055517f21fe8cef36ad2d44c2fbaf4388d98bfc0e42794abac7e40f17b9a04dda7d60a1906106959085815260200190565b6001545f9081908190819081906001600160a01b0316610a395750506004545f90610ab5565b60015460405163e588566f60e01b81526001600160a01b0388811660048301529091169063e588566f90602401606060405180830381865afa158015610a81573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aa59190611b1a565b50909250905081610ab557506004545b805f03610acc57670de0b6b3a76400009450610b72565b604051636a24d41960e11b81526001600160a01b03871660048201525f9073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9063d449a83290602401602060405180830381865af4158015610b24573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b489190611b52565b610b5390600a611c55565b90508181600554610b649190611c63565b610b6e9190611ad5565b9550505b6001600160a01b0386165f90815260096020526040902054925060018301610b9c575f9250610ba9565b825f03610ba95760035492505b610bb3878761152b565b508094505050509250925092565b6420a226a4a760d91b610bd3816111f4565b6001600160a01b038216610c16576040516323056b1360e01b81526020600482015260096024820152681c1c9a58d95199595960ba1b60448201526064016109be565b600180546001600160a01b0319166001600160a01b0384169081179091556040517f931b9726a5ca07d2f1c9028e8d50480d8de967200ef6fa68e87e099004002766905f90a25050565b662aa82220aa27a960c91b610c74816111f4565b8151835114610c965760405163322473c360e01b815260040160405180910390fd5b5f5b835181101561057557828181518110610cb357610cb3611a99565b6020026020010151600a5f868481518110610cd057610cd0611a99565b602002602001015181526020019081526020015f2081905550838181518110610cfb57610cfb611a99565b60200260200101517f3fc45c54c1bb058234a4c562d7bcc2bb829ce9f5842d6d0bcab221506f10a320848381518110610d3657610d36611a99565b6020026020010151604051610d4d91815260200190565b60405180910390a2600101610c98565b6420a226a4a760d91b610d6f816111f4565b60078290556040518281527f52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a490602001610965565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b6420a226a4a760d91b610dde816111f4565b60058290556040518281527f0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e990602001610965565b6420a226a4a760d91b610e25816111f4565b60088290556040518281527f7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e290602001610965565b6420a226a4a760d91b610e6c816111f4565b60038290556040518281525f907f21fe8cef36ad2d44c2fbaf4388d98bfc0e42794abac7e40f17b9a04dda7d60a19060200160405180910390a25050565b6001545f9081906001600160a01b0316610ec6575f9150610f39565b60015460405163e588566f60e01b81526001600160a01b0386811660048301529091169063e588566f90602401606060405180830381865afa158015610f0e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f329190611b1a565b5090925090505b81610f4357506004545b604051636a24d41960e11b81526001600160a01b0385166004820152610fd690849073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9063d449a83290602401602060405180830381865af4158015610f9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc39190611b52565b610fce90600a611c55565b83919061162c565b90509250929050565b6420a226a4a760d91b610ff1816111f4565b60068290556040518281527ff905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a290602001610965565b6420a226a4a760d91b611038816111f4565b600180546001600160a01b03191690556040515f907f931b9726a5ca07d2f1c9028e8d50480d8de967200ef6fa68e87e099004002766908290a250565b604051636a24d41960e11b81526001600160a01b03821660048201525f90819061110490849073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9063d449a83290602401602060405180830381865af41580156110d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110f99190611b52565b61041290600a611c55565b91509150915091565b5f82815260208190526040902060010154611127816111f4565b61057583836112ef565b5f5f5f5f61113f8787610a13565b919550935090506103e86111538287611c63565b61115d9190611ad5565b91505093509350939050565b6420a226a4a760d91b61117b816111f4565b815f036111bf576040516323056b1360e01b815260206004820152601160248201527066696e616c697a6542726964676547617360781b60448201526064016109be565b60028290556040518281527f428221d3b9d5714a54e6899f3d9fa1bf85bfe37c725d781b09ec5f399412fbc990602001610965565b6111fe81336116e3565b50565b546001600160801b03808216600160801b9092048116919091031690565b5f61122983611201565b8210611239576112396032611720565b5081546001600160801b039081168201165f90815260018301602052604090205492915050565b5f61126b8383610da4565b6112e8575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556112a03390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104ec565b505f6104ec565b5f6112fa8383610da4565b156112e8575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104ec565b5f6113768254600160801b81046001600160801b0390811691161490565b15611385576113856032611720565b5080546001600160801b03165f9081526001909101602052604090205490565b80545f906001600160801b0380821691600160801b90041681036113cd576113cd6031611720565b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6114308254600160801b81046001600160801b0390811691161490565b1561143f5761143f6032611720565b5080545f196001600160801b03600160801b909204821601165f9081526001909101602052604090205490565b80545f906001600160801b03600160801b8204811691168103611493576114936031611720565b5f19016001600160801b039081165f818152600185016020526040812080549190558454909216600160801b909102179092555090565b81546001600160801b03600160801b8204811691811660018301909116036114f6576114f66041611720565b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f828152600a60205260408120546001548291906001600160a01b03161580611552575080155b15611563575f5f9250925050611625565b60015460025473__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9163889ad9e0916001600160a01b0390911690889088906115a0908790611c63565b6040516001600160e01b031960e087901b1681526001600160a01b0394851660048201526024810193909352921660448201526064810191909152608401606060405180830381865af41580156115f9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061161d9190611b1a565b909450925050505b9250929050565b5f838302815f1985870982811083820303915050805f036116605783828161165657611656611aad565b04925050506116dc565b808411611677576116776003851502601118611720565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b6116ed8282610da4565b61171c5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016109be565b5050565b634e487b715f52806020526024601cfd5b5f60208284031215611741575f5ffd5b81356001600160e01b0319811681146116dc575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561179557611795611758565b604052919050565b5f67ffffffffffffffff8211156117b6576117b6611758565b5060051b60200190565b6001600160a01b03811681146111fe575f5ffd5b5f82601f8301126117e3575f5ffd5b81356117f66117f18261179d565b61176c565b8082825260208201915060208360051b860101925085831115611817575f5ffd5b602085015b8381101561183457803583526020928301920161181c565b5095945050505050565b5f5f6040838503121561184f575f5ffd5b823567ffffffffffffffff811115611865575f5ffd5b8301601f81018513611875575f5ffd5b80356118836117f18261179d565b8082825260208201915060208360051b8501019250878311156118a4575f5ffd5b6020840193505b828410156118cf5783356118be816117c0565b8252602093840193909101906118ab565b9450505050602083013567ffffffffffffffff8111156118ed575f5ffd5b6118f9858286016117d4565b9150509250929050565b5f60208284031215611913575f5ffd5b81356116dc816117c0565b602080825282518282018190525f918401906040840190835b81811015611955578351835260209384019390920191600101611937565b509095945050505050565b5f5f60408385031215611971575f5ffd5b50508035926020909101359150565b5f60208284031215611990575f5ffd5b5035919050565b5f5f604083850312156119a8575f5ffd5b8235915060208301356119ba816117c0565b809150509250929050565b5f5f604083850312156119d6575f5ffd5b82356119e1816117c0565b946020939093013593505050565b6020810160088310611a0f57634e487b7160e01b5f52602160045260245ffd5b91905290565b5f5f60408385031215611a26575f5ffd5b823567ffffffffffffffff811115611a3c575f5ffd5b611a48858286016117d4565b925050602083013567ffffffffffffffff8111156118ed575f5ffd5b5f5f5f60608486031215611a76575f5ffd5b833592506020840135611a88816117c0565b929592945050506040919091013590565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f82611aef57634e487b7160e01b5f52601260045260245ffd5b500490565b818103818111156104ec576104ec611ac1565b808201808211156104ec576104ec611ac1565b5f5f5f60608486031215611b2c575f5ffd5b83518015158114611b3b575f5ffd5b602085015160409095015190969495509392505050565b5f60208284031215611b62575f5ffd5b815160ff811681146116dc575f5ffd5b6001815b6001841115611bad57808504811115611b9157611b91611ac1565b6001841615611b9f57908102905b60019390931c928002611b76565b935093915050565b5f82611bc3575060016104ec565b81611bcf57505f6104ec565b8160018114611be55760028114611bef57611c0b565b60019150506104ec565b60ff841115611c0057611c00611ac1565b50506001821b6104ec565b5060208310610133831016604e8410600b8410161715611c2e575081810a6104ec565b611c3a5f198484611b72565b805f1904821115611c4d57611c4d611ac1565b029392505050565b5f6116dc60ff841683611bb5565b80820281158282048414176104ec576104ec611ac156fea26469706673582212200ebf73748eec9571577c73139525e13a9c62335ec8122e50d9db64d84bbeb65d64736f6c634300081c0033",
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
func DeployBridgeVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, _bridge common.Address, finalizeBridgeGas *big.Int, defaultTokenPrice *big.Int, defaultExFeeRate *big.Int, minimumTokenValue *big.Int, verificationAmountThreshold *big.Int, periodTotalValueThreshold *big.Int, timeWindow *big.Int) (common.Address, *types.Transaction, *BridgeVerifier, error) {
	parsed, err := BridgeVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeVerifierBin), backend, initialOwner, _bridge, finalizeBridgeGas, defaultTokenPrice, defaultExFeeRate, minimumTokenValue, verificationAmountThreshold, periodTotalValueThreshold, timeWindow)
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
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumValue *big.Int
		GasFee       *big.Int
		ExFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeVerifier.Contract.CalculateFee(&_BridgeVerifier.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeVerifier *BridgeVerifierCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
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
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierCaller) GetTokenConfig(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	var out []interface{}
	err := _BridgeVerifier.contract.Call(opts, &out, "getTokenConfig", remoteChainID, token)

	outstruct := new(struct {
		MinimumValue *big.Int
		GasFee       *big.Int
		ExFeeRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFeeRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeVerifier.Contract.GetTokenConfig(&_BridgeVerifier.CallOpts, remoteChainID, token)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierCallerSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
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

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeVerifier *BridgeVerifierTransactor) RemovePriceFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "removePriceFeed")
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeVerifier *BridgeVerifierSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemovePriceFeed(&_BridgeVerifier.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeVerifier.Contract.RemovePriceFeed(&_BridgeVerifier.TransactOpts)
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
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPriceFeed(&_BridgeVerifier.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeVerifier *BridgeVerifierTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeVerifier.Contract.SetPriceFeed(&_BridgeVerifier.TransactOpts, priceFeed)
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

// BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator is returned from FilterBridgeVerifierExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeVerifierExchangeFeeUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator struct {
	Event *BridgeVerifierBridgeVerifierExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierBridgeVerifierExchangeFeeUpdated)
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
		it.Event = new(BridgeVerifierBridgeVerifierExchangeFeeUpdated)
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
func (it *BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierBridgeVerifierExchangeFeeUpdated represents a BridgeVerifierExchangeFeeUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierExchangeFeeUpdated struct {
	Token     common.Address
	ExFeeRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierExchangeFeeUpdated is a free log retrieval operation binding the contract event 0x21fe8cef36ad2d44c2fbaf4388d98bfc0e42794abac7e40f17b9a04dda7d60a1.
//
// Solidity: event BridgeVerifierExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterBridgeVerifierExchangeFeeUpdated(opts *bind.FilterOpts, token []common.Address) (*BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "BridgeVerifierExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierBridgeVerifierExchangeFeeUpdatedIterator{contract: _BridgeVerifier.contract, event: "BridgeVerifierExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierExchangeFeeUpdated is a free log subscription operation binding the contract event 0x21fe8cef36ad2d44c2fbaf4388d98bfc0e42794abac7e40f17b9a04dda7d60a1.
//
// Solidity: event BridgeVerifierExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchBridgeVerifierExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierBridgeVerifierExchangeFeeUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "BridgeVerifierExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierBridgeVerifierExchangeFeeUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierExchangeFeeUpdated", log); err != nil {
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

// ParseBridgeVerifierExchangeFeeUpdated is a log parse operation binding the contract event 0x21fe8cef36ad2d44c2fbaf4388d98bfc0e42794abac7e40f17b9a04dda7d60a1.
//
// Solidity: event BridgeVerifierExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseBridgeVerifierExchangeFeeUpdated(log types.Log) (*BridgeVerifierBridgeVerifierExchangeFeeUpdated, error) {
	event := new(BridgeVerifierBridgeVerifierExchangeFeeUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator is returned from FilterBridgeVerifierFinalizeBridgeGasSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierFinalizeBridgeGasSet events raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator struct {
	Event *BridgeVerifierBridgeVerifierFinalizeBridgeGasSet // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierBridgeVerifierFinalizeBridgeGasSet)
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
		it.Event = new(BridgeVerifierBridgeVerifierFinalizeBridgeGasSet)
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
func (it *BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierBridgeVerifierFinalizeBridgeGasSet represents a BridgeVerifierFinalizeBridgeGasSet event raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierFinalizeBridgeGasSet struct {
	FinalizeBridgeGas *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierFinalizeBridgeGasSet is a free log retrieval operation binding the contract event 0x428221d3b9d5714a54e6899f3d9fa1bf85bfe37c725d781b09ec5f399412fbc9.
//
// Solidity: event BridgeVerifierFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterBridgeVerifierFinalizeBridgeGasSet(opts *bind.FilterOpts) (*BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator, error) {

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "BridgeVerifierFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierBridgeVerifierFinalizeBridgeGasSetIterator{contract: _BridgeVerifier.contract, event: "BridgeVerifierFinalizeBridgeGasSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierFinalizeBridgeGasSet is a free log subscription operation binding the contract event 0x428221d3b9d5714a54e6899f3d9fa1bf85bfe37c725d781b09ec5f399412fbc9.
//
// Solidity: event BridgeVerifierFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchBridgeVerifierFinalizeBridgeGasSet(opts *bind.WatchOpts, sink chan<- *BridgeVerifierBridgeVerifierFinalizeBridgeGasSet) (event.Subscription, error) {

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "BridgeVerifierFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierBridgeVerifierFinalizeBridgeGasSet)
				if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierFinalizeBridgeGasSet", log); err != nil {
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

// ParseBridgeVerifierFinalizeBridgeGasSet is a log parse operation binding the contract event 0x428221d3b9d5714a54e6899f3d9fa1bf85bfe37c725d781b09ec5f399412fbc9.
//
// Solidity: event BridgeVerifierFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseBridgeVerifierFinalizeBridgeGasSet(log types.Log) (*BridgeVerifierBridgeVerifierFinalizeBridgeGasSet, error) {
	event := new(BridgeVerifierBridgeVerifierFinalizeBridgeGasSet)
	if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierFinalizeBridgeGasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierBridgeVerifierGasPriceUpdatedIterator is returned from FilterBridgeVerifierGasPriceUpdated and is used to iterate over the raw logs and unpacked data for BridgeVerifierGasPriceUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierGasPriceUpdatedIterator struct {
	Event *BridgeVerifierBridgeVerifierGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierBridgeVerifierGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierBridgeVerifierGasPriceUpdated)
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
		it.Event = new(BridgeVerifierBridgeVerifierGasPriceUpdated)
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
func (it *BridgeVerifierBridgeVerifierGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierBridgeVerifierGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierBridgeVerifierGasPriceUpdated represents a BridgeVerifierGasPriceUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierGasPriceUpdated struct {
	RemoteChainID *big.Int
	GasPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierGasPriceUpdated is a free log retrieval operation binding the contract event 0x3fc45c54c1bb058234a4c562d7bcc2bb829ce9f5842d6d0bcab221506f10a320.
//
// Solidity: event BridgeVerifierGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterBridgeVerifierGasPriceUpdated(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BridgeVerifierBridgeVerifierGasPriceUpdatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "BridgeVerifierGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierBridgeVerifierGasPriceUpdatedIterator{contract: _BridgeVerifier.contract, event: "BridgeVerifierGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierGasPriceUpdated is a free log subscription operation binding the contract event 0x3fc45c54c1bb058234a4c562d7bcc2bb829ce9f5842d6d0bcab221506f10a320.
//
// Solidity: event BridgeVerifierGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchBridgeVerifierGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierBridgeVerifierGasPriceUpdated, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "BridgeVerifierGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierBridgeVerifierGasPriceUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierGasPriceUpdated", log); err != nil {
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

// ParseBridgeVerifierGasPriceUpdated is a log parse operation binding the contract event 0x3fc45c54c1bb058234a4c562d7bcc2bb829ce9f5842d6d0bcab221506f10a320.
//
// Solidity: event BridgeVerifierGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseBridgeVerifierGasPriceUpdated(log types.Log) (*BridgeVerifierBridgeVerifierGasPriceUpdated, error) {
	event := new(BridgeVerifierBridgeVerifierGasPriceUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator is returned from FilterBridgeVerifierPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for BridgeVerifierPriceFeedUpdated events raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator struct {
	Event *BridgeVerifierBridgeVerifierPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVerifierBridgeVerifierPriceFeedUpdated)
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
		it.Event = new(BridgeVerifierBridgeVerifierPriceFeedUpdated)
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
func (it *BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVerifierBridgeVerifierPriceFeedUpdated represents a BridgeVerifierPriceFeedUpdated event raised by the BridgeVerifier contract.
type BridgeVerifierBridgeVerifierPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierPriceFeedUpdated is a free log retrieval operation binding the contract event 0x931b9726a5ca07d2f1c9028e8d50480d8de967200ef6fa68e87e099004002766.
//
// Solidity: event BridgeVerifierPriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) FilterBridgeVerifierPriceFeedUpdated(opts *bind.FilterOpts, priceFeed []common.Address) (*BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeVerifier.contract.FilterLogs(opts, "BridgeVerifierPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVerifierBridgeVerifierPriceFeedUpdatedIterator{contract: _BridgeVerifier.contract, event: "BridgeVerifierPriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierPriceFeedUpdated is a free log subscription operation binding the contract event 0x931b9726a5ca07d2f1c9028e8d50480d8de967200ef6fa68e87e099004002766.
//
// Solidity: event BridgeVerifierPriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) WatchBridgeVerifierPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeVerifierBridgeVerifierPriceFeedUpdated, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeVerifier.contract.WatchLogs(opts, "BridgeVerifierPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVerifierBridgeVerifierPriceFeedUpdated)
				if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierPriceFeedUpdated", log); err != nil {
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

// ParseBridgeVerifierPriceFeedUpdated is a log parse operation binding the contract event 0x931b9726a5ca07d2f1c9028e8d50480d8de967200ef6fa68e87e099004002766.
//
// Solidity: event BridgeVerifierPriceFeedUpdated(address indexed priceFeed)
func (_BridgeVerifier *BridgeVerifierFilterer) ParseBridgeVerifierPriceFeedUpdated(log types.Log) (*BridgeVerifierBridgeVerifierPriceFeedUpdated, error) {
	event := new(BridgeVerifierBridgeVerifierPriceFeedUpdated)
	if err := _BridgeVerifier.contract.UnpackLog(event, "BridgeVerifierPriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
