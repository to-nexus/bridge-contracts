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

// BridgeManagerMetaData contains all meta data concerning the BridgeManager contract.
var BridgeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPeriodTotalValueThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCurrentScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMovementHistory\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"history\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getTokenPriceWithValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerificationAmountThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"setDefaultTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setExFeeRateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumTokenValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"setPeriodTotalValueThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"setTimeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"setVerificationAmountThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasPrices\",\"type\":\"uint256[]\"}],\"name\":\"updateGasPriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"validateBridgeTokenValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"BridgeManagerExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"BridgeManagerFinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"BridgeManagerGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"BridgeManagerPriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTokenPrice\",\"type\":\"uint256\"}],\"name\":\"DefaultTokenPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumTokenValue\",\"type\":\"uint256\"}],\"name\":\"MinimumTokenValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodTotalValueThreshold\",\"type\":\"uint256\"}],\"name\":\"PeriodTotalValueThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeWindow\",\"type\":\"uint256\"}],\"name\":\"TimeWindowSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationAmountThreshold\",\"type\":\"uint256\"}],\"name\":\"VerificationAmountThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeManagerCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeManagerChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeManagerInvalidLength\",\"type\":\"error\"}]",
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
	Bin: "0x608060405234801561000f575f5ffd5b50604051611ef6380380611ef683398101604081905261002e9161021a565b865f03610077576040516378eb4e6d60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b60448201526064015b60405180910390fd5b6001600160a01b0389166100bd576040516378eb4e6d60e11b815260206004820152600c60248201526b34b734ba34b0b627bbb732b960a11b604482015260640161006e565b6001600160a01b0388166100fe576040516378eb4e6d60e11b81526020600482015260076024820152665f62726964676560c81b604482015260640161006e565b6101085f8a610156565b5061011b6420a226a4a760d91b8a610156565b5061012f6542524944474560d01b89610156565b50600296909655600494909455600392909255600555600655600855600755506102ae9050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166101f6575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101ae3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101f9565b505f5b92915050565b80516001600160a01b0381168114610215575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610233575f5ffd5b61023c8a6101ff565b985061024a60208b016101ff565b97505f60408b01519050809750505f60608b01519050809650505f60808b01519050809550505f60a08b01519050809450505f60c08b01519050809350505f60e08b01519050809250505f6101008b01519050809150509295985092959850929598565b611c3b806102bb5f395ff3fe608060405234801561000f575f5ffd5b50600436106101f2575f3560e01c806391d1485411610114578063ba8d25f8116100a9578063d02641a011610079578063d02641a014610452578063d1c0154314610465578063d547741f14610484578063df6e87dc14610497578063f289d3ba146104aa575f5ffd5b8063ba8d25f814610405578063be716b0a1461042f578063c3de108b14610437578063ce2e5c661461044a575f5ffd5b80639dc696cc116100e45780639dc696cc146103c5578063a217fddf146103d8578063ab509f3a146103df578063ad1434ca146103f2575f5ffd5b806391d148541461037a57806396ce07951461038d5780639745e4ba14610395578063979eb82d146103bd575f5ffd5b80633f4bdec51161018a5780636332aec61161015a5780636332aec614610313578063724e78da14610341578063751ccd79146103545780638d33199614610367575f5ffd5b80633f4bdec5146102c457806349c9215e146102e55780635566ca09146102f85780635ec8f51b14610300575f5ffd5b8063248a9ca3116101c5578063248a9ca3146102665780632a4d2f32146102965780632f2ff15d1461029e57806336568abe146102b1575f5ffd5b806301ffc9a7146101f6578063020a22121461021e5780630e4e96bf146102335780631f96131e14610253575b5f5ffd5b610209610204366004611659565b6104bd565b60405190151581526020015b60405180910390f35b61023161022c366004611766565b6104f3565b005b61024661024136600461182b565b61057e565b6040516102159190611846565b610231610261366004611888565b610640565b6102886102743660046118a8565b5f9081526020819052604090206001015490565b604051908152602001610215565b600554610288565b6102316102ac3660046118bf565b6106a5565b6102316102bf3660046118bf565b6106c9565b6102d76102d23660046118ed565b610701565b604051610215929190611917565b6102316102f33660046118a8565b6108dc565b600754610288565b61023161030e3660046118ed565b61092b565b6103266103213660046118bf565b6109cf565b60408051938452602084019290925290820152606001610215565b61023161034f36600461182b565b610b89565b610231610362366004611955565b610c28565b6102316103753660046118a8565b610d25565b6102096103883660046118bf565b610d6c565b6103e8610288565b6102886103a336600461182b565b6001600160a01b03165f908152600b602052604090205490565b600854610288565b6102316103d33660046118a8565b610d94565b6102885f81565b6102316103ed3660046118a8565b610ddb565b6102316104003660046118a8565b610e22565b6104186104133660046118ed565b610e72565b604080519215158352602083019190915201610215565b600654610288565b6102316104453660046118a8565b610fb3565b610231610ffa565b61041861046036600461182b565b611049565b6102886104733660046118a8565b5f908152600a602052604090205490565b6102316104923660046118bf565b6110ed565b6103266104a53660046119a4565b611111565b6102316104b83660046118a8565b611149565b5f6001600160e01b03198216637965db0b60e01b14806104ed57506301ffc9a760e01b6001600160e01b03198316145b92915050565b662aa82220aa27a960c91b610507816111d4565b81518351146105295760405163d8d9560b60e01b815260040160405180910390fd5b5f5b835181101561057857610570848281518110610549576105496119d9565b6020026020010151848381518110610563576105636119d9565b602002602001015161092b565b60010161052b565b50505050565b6001600160a01b0381165f908152600c60205260408120606091906105a2906111e1565b90508067ffffffffffffffff8111156105bd576105bd611680565b6040519080825280602002602001820160405280156105e6578160200160208202803683370190505b5091505f5b81811015610639576001600160a01b0384165f908152600c6020526040902061061490826111ff565b838281518110610626576106266119d9565b60209081029190910101526001016105eb565b5050919050565b662aa82220aa27a960c91b610654816111d4565b5f838152600a6020526040908190208390555183907f250cb85ef4edc2a5afdc6bf99c82fd11536c37fe9ad56db2866db5a24b14d848906106989085815260200190565b60405180910390a2505050565b5f828152602081905260409020600101546106bf816111d4565b6105788383611240565b6001600160a01b03811633146106f25760405163334bd91960e11b815260040160405180910390fd5b6106fc82826112cf565b505050565b5f60606542524944474560d01b610717816111d4565b5f6107228686610e72565b9150506006545f14158015610738575080600654105b15610761575f604051806060016040528060268152602001611bbb6026913993509350506108d4565b6001600160a01b0386165f908152600c60205260408120600754909142916107899083611a01565b90505b5f610796846111e1565b1115610826575f6107a684611338565b905060c081901c82811015610818576001600160a01b038b165f908152600b60205260409020546001600160c01b038316908181116107e5575f6107ef565b6107ef8282611a01565b6001600160a01b038e165f908152600b602052604090205561081087611385565b50505061081f565b5050610826565b505061078c565b60c082901b841761083784826113f2565b6001600160a01b038a165f908152600b60205260408120805487929061085e908490611a14565b90915550506008541580159061088c57506008546001600160a01b038b165f908152600b6020526040902054115b156108b9575f604051806060016040528060258152602001611be1602591399750975050505050506108d4565b600160405180602001604052805f8152509750975050505050505b509250929050565b6420a226a4a760d91b6108ee816111d4565b60048290556040518281527f1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a906020015b60405180910390a15050565b662aa82220aa27a960c91b61093f816111d4565b6001600160a01b038316610983576040516378eb4e6d60e11b81526020600482015260056024820152643a37b5b2b760d91b60448201526064015b60405180910390fd5b6001600160a01b0383165f8181526009602052604090819020849055517fe02362db7339f5c8733f9cf708855b61c1a82ae71e49239db928fe4c3ddd802b906106989085815260200190565b6001545f9081908190819081906001600160a01b03166109f55750506004545f90610a71565b60015460405163e588566f60e01b81526001600160a01b0388811660048301529091169063e588566f90602401606060405180830381865afa158015610a3d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a619190611a27565b50909250905081610a7157506004545b805f03610a8857670de0b6b3a76400009450610b3a565b600154604051630fe74f3b60e01b81526001600160a01b03918216600482015290871660248201525f9073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__90630fe74f3b90604401602060405180830381865af4158015610aec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b109190611a5f565b610b1b90600a611b62565b90508181600554610b2c9190611b70565b610b369190611b9b565b9550505b6001600160a01b0386165f90815260096020526040902054925060018301610b64575f9250610b71565b825f03610b715760035492505b610b7b8787611453565b508094505050509250925092565b6420a226a4a760d91b610b9b816111d4565b6001600160a01b038216610bde576040516378eb4e6d60e11b81526020600482015260096024820152681c1c9a58d95199595960ba1b604482015260640161097a565b600180546001600160a01b0319166001600160a01b0384169081179091556040517f6433071fe185ca8d6b5d1f327404fcc1ab697b7e9f29635aa9b0575ec51e5eaf905f90a25050565b662aa82220aa27a960c91b610c3c816111d4565b8151835114610c5e5760405163d8d9560b60e01b815260040160405180910390fd5b5f5b835181101561057857828181518110610c7b57610c7b6119d9565b6020026020010151600a5f868481518110610c9857610c986119d9565b602002602001015181526020019081526020015f2081905550838181518110610cc357610cc36119d9565b60200260200101517f250cb85ef4edc2a5afdc6bf99c82fd11536c37fe9ad56db2866db5a24b14d848848381518110610cfe57610cfe6119d9565b6020026020010151604051610d1591815260200190565b60405180910390a2600101610c60565b6420a226a4a760d91b610d37816111d4565b60078290556040518281527f52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a49060200161091f565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b6420a226a4a760d91b610da6816111d4565b60058290556040518281527f0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e99060200161091f565b6420a226a4a760d91b610ded816111d4565b60088290556040518281527f7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e29060200161091f565b6420a226a4a760d91b610e34816111d4565b60038290556040518281525f907fe02362db7339f5c8733f9cf708855b61c1a82ae71e49239db928fe4c3ddd802b9060200160405180910390a25050565b6001545f9081906001600160a01b0316610e8e575f9150610f01565b60015460405163e588566f60e01b81526001600160a01b0386811660048301529091169063e588566f90602401606060405180830381865afa158015610ed6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610efa9190611a27565b5090925090505b81610f0b57506004545b600154604051630fe74f3b60e01b81526001600160a01b0391821660048201529085166024820152610faa90849073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__90630fe74f3b90604401602060405180830381865af4158015610f73573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f979190611a5f565b610fa290600a611b62565b839190611554565b90509250929050565b6420a226a4a760d91b610fc5816111d4565b60068290556040518281527ff905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a29060200161091f565b6420a226a4a760d91b61100c816111d4565b600180546001600160a01b03191690556040515f907f6433071fe185ca8d6b5d1f327404fcc1ab697b7e9f29635aa9b0575ec51e5eaf908290a250565b600154604051630fe74f3b60e01b81526001600160a01b03918216600482015290821660248201525f9081906110e490849073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__90630fe74f3b90604401602060405180830381865af41580156110b5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110d99190611a5f565b61041390600a611b62565b91509150915091565b5f82815260208190526040902060010154611107816111d4565b61057883836112cf565b5f5f5f5f61111f87876109cf565b919550935090506103e86111338287611b70565b61113d9190611b9b565b91505093509350939050565b6420a226a4a760d91b61115b816111d4565b815f0361119f576040516378eb4e6d60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b604482015260640161097a565b60028290556040518281527f92fa3d1383e1ab6627e21af8131fd78645136294a46271cbe6c0d22350685da79060200161091f565b6111de813361160b565b50565b546001600160801b03808216600160801b9092048116919091031690565b5f611209836111e1565b8210611219576112196032611648565b5081546001600160801b039081168201165f90815260018301602052604090205492915050565b5f61124b8383610d6c565b6112c8575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556112803390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104ed565b505f6104ed565b5f6112da8383610d6c565b156112c8575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104ed565b5f6113568254600160801b81046001600160801b0390811691161490565b15611365576113656032611648565b5080546001600160801b03165f9081526001909101602052604090205490565b80545f906001600160801b0380821691600160801b90041681036113ad576113ad6031611648565b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b81546001600160801b03600160801b82048116918116600183019091160361141e5761141e6041611648565b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b5f828152600a60205260408120546001548291906001600160a01b0316158061147a575080155b1561148b575f5f925092505061154d565b60015460025473__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9163889ad9e0916001600160a01b0390911690889088906114c8908790611b70565b6040516001600160e01b031960e087901b1681526001600160a01b0394851660048201526024810193909352921660448201526064810191909152608401606060405180830381865af4158015611521573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115459190611a27565b909450925050505b9250929050565b5f838302815f1985870982811083820303915050805f036115885783828161157e5761157e611b87565b0492505050611604565b80841161159f5761159f6003851502601118611648565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b6116158282610d6c565b6116445760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161097a565b5050565b634e487b715f52806020526024601cfd5b5f60208284031215611669575f5ffd5b81356001600160e01b031981168114611604575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156116bd576116bd611680565b604052919050565b5f67ffffffffffffffff8211156116de576116de611680565b5060051b60200190565b6001600160a01b03811681146111de575f5ffd5b5f82601f83011261170b575f5ffd5b813561171e611719826116c5565b611694565b8082825260208201915060208360051b86010192508583111561173f575f5ffd5b602085015b8381101561175c578035835260209283019201611744565b5095945050505050565b5f5f60408385031215611777575f5ffd5b823567ffffffffffffffff81111561178d575f5ffd5b8301601f8101851361179d575f5ffd5b80356117ab611719826116c5565b8082825260208201915060208360051b8501019250878311156117cc575f5ffd5b6020840193505b828410156117f75783356117e6816116e8565b8252602093840193909101906117d3565b9450505050602083013567ffffffffffffffff811115611815575f5ffd5b611821858286016116fc565b9150509250929050565b5f6020828403121561183b575f5ffd5b8135611604816116e8565b602080825282518282018190525f918401906040840190835b8181101561187d57835183526020938401939092019160010161185f565b509095945050505050565b5f5f60408385031215611899575f5ffd5b50508035926020909101359150565b5f602082840312156118b8575f5ffd5b5035919050565b5f5f604083850312156118d0575f5ffd5b8235915060208301356118e2816116e8565b809150509250929050565b5f5f604083850312156118fe575f5ffd5b8235611909816116e8565b946020939093013593505050565b8215158152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f8301168401019150509392505050565b5f5f60408385031215611966575f5ffd5b823567ffffffffffffffff81111561197c575f5ffd5b611988858286016116fc565b925050602083013567ffffffffffffffff811115611815575f5ffd5b5f5f5f606084860312156119b6575f5ffd5b8335925060208401356119c8816116e8565b929592945050506040919091013590565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b818103818111156104ed576104ed6119ed565b808201808211156104ed576104ed6119ed565b5f5f5f60608486031215611a39575f5ffd5b83518015158114611a48575f5ffd5b602085015160409095015190969495509392505050565b5f60208284031215611a6f575f5ffd5b815160ff81168114611604575f5ffd5b6001815b6001841115611aba57808504811115611a9e57611a9e6119ed565b6001841615611aac57908102905b60019390931c928002611a83565b935093915050565b5f82611ad0575060016104ed565b81611adc57505f6104ed565b8160018114611af25760028114611afc57611b18565b60019150506104ed565b60ff841115611b0d57611b0d6119ed565b50506001821b6104ed565b5060208310610133831016604e8410600b8410161715611b3b575081810a6104ed565b611b475f198484611a7f565b805f1904821115611b5a57611b5a6119ed565b029392505050565b5f61160460ff841683611ac2565b80820281158282048414176104ed576104ed6119ed565b634e487b7160e01b5f52601260045260245ffd5b5f82611bb557634e487b7160e01b5f52601260045260245ffd5b50049056fe766572696669636174696f6e20616d6f756e74207468726573686f6c64206578636565646564706572696f6420746f74616c2076616c7565207468726573686f6c64206578636565646564a26469706673582212200fc9c4cce68f95c0d66e3093dc6995ef3922c216c4e82a7158736c064f726d4b64736f6c634300081c0033",
}

// BridgeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeManagerMetaData.ABI instead.
var BridgeManagerABI = BridgeManagerMetaData.ABI

// Deprecated: Use BridgeManagerMetaData.Sigs instead.
// BridgeManagerFuncSigs maps the 4-byte function signature to its string representation.
var BridgeManagerFuncSigs = BridgeManagerMetaData.Sigs

// BridgeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeManagerMetaData.Bin instead.
var BridgeManagerBin = BridgeManagerMetaData.Bin

// DeployBridgeManager deploys a new Ethereum contract, binding an instance of BridgeManager to it.
func DeployBridgeManager(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, _bridge common.Address, finalizeBridgeGas *big.Int, defaultTokenPrice *big.Int, defaultExFeeRate *big.Int, minimumTokenValue *big.Int, verificationAmountThreshold *big.Int, periodTotalValueThreshold *big.Int, timeWindow *big.Int) (common.Address, *types.Transaction, *BridgeManager, error) {
	parsed, err := BridgeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeManagerBin), backend, initialOwner, _bridge, finalizeBridgeGas, defaultTokenPrice, defaultExFeeRate, minimumTokenValue, verificationAmountThreshold, periodTotalValueThreshold, timeWindow)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeManager{BridgeManagerCaller: BridgeManagerCaller{contract: contract}, BridgeManagerTransactor: BridgeManagerTransactor{contract: contract}, BridgeManagerFilterer: BridgeManagerFilterer{contract: contract}}, nil
}

// BridgeManager is an auto generated Go binding around an Ethereum contract.
type BridgeManager struct {
	BridgeManagerCaller     // Read-only binding to the contract
	BridgeManagerTransactor // Write-only binding to the contract
	BridgeManagerFilterer   // Log filterer for contract events
}

// BridgeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeManagerSession struct {
	Contract     *BridgeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeManagerCallerSession struct {
	Contract *BridgeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeManagerTransactorSession struct {
	Contract     *BridgeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeManagerRaw struct {
	Contract *BridgeManager // Generic contract binding to access the raw methods on
}

// BridgeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeManagerCallerRaw struct {
	Contract *BridgeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeManagerTransactorRaw struct {
	Contract *BridgeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeManager creates a new instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManager(address common.Address, backend bind.ContractBackend) (*BridgeManager, error) {
	contract, err := bindBridgeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeManager{BridgeManagerCaller: BridgeManagerCaller{contract: contract}, BridgeManagerTransactor: BridgeManagerTransactor{contract: contract}, BridgeManagerFilterer: BridgeManagerFilterer{contract: contract}}, nil
}

// NewBridgeManagerCaller creates a new read-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerCaller(address common.Address, caller bind.ContractCaller) (*BridgeManagerCaller, error) {
	contract, err := bindBridgeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerCaller{contract: contract}, nil
}

// NewBridgeManagerTransactor creates a new write-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeManagerTransactor, error) {
	contract, err := bindBridgeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerTransactor{contract: contract}, nil
}

// NewBridgeManagerFilterer creates a new log filterer instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeManagerFilterer, error) {
	contract, err := bindBridgeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerFilterer{contract: contract}, nil
}

// bindBridgeManager binds a generic wrapper to an already deployed contract.
func bindBridgeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.BridgeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeManager *BridgeManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeManager *BridgeManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeManager.Contract.DEFAULTADMINROLE(&_BridgeManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeManager *BridgeManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeManager.Contract.DEFAULTADMINROLE(&_BridgeManager.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeManager *BridgeManagerCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

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
func (_BridgeManager *BridgeManagerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeManager.Contract.CalculateFee(&_BridgeManager.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeManager *BridgeManagerCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeManager.Contract.CalculateFee(&_BridgeManager.CallOpts, remoteChainID, token, value)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeManager *BridgeManagerCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeManager *BridgeManagerSession) Denominator() (*big.Int, error) {
	return _BridgeManager.Contract.Denominator(&_BridgeManager.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) Denominator() (*big.Int, error) {
	return _BridgeManager.Contract.Denominator(&_BridgeManager.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetGasPrice(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGasPrice", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeManager.Contract.GetGasPrice(&_BridgeManager.CallOpts, remoteChainID)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeManager.Contract.GetGasPrice(&_BridgeManager.CallOpts, remoteChainID)
}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetMinimumTokenValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getMinimumTokenValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetMinimumTokenValue() (*big.Int, error) {
	return _BridgeManager.Contract.GetMinimumTokenValue(&_BridgeManager.CallOpts)
}

// GetMinimumTokenValue is a free data retrieval call binding the contract method 0x2a4d2f32.
//
// Solidity: function getMinimumTokenValue() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetMinimumTokenValue() (*big.Int, error) {
	return _BridgeManager.Contract.GetMinimumTokenValue(&_BridgeManager.CallOpts)
}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetPeriodTotalValueThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getPeriodTotalValueThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetPeriodTotalValueThreshold() (*big.Int, error) {
	return _BridgeManager.Contract.GetPeriodTotalValueThreshold(&_BridgeManager.CallOpts)
}

// GetPeriodTotalValueThreshold is a free data retrieval call binding the contract method 0x979eb82d.
//
// Solidity: function getPeriodTotalValueThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetPeriodTotalValueThreshold() (*big.Int, error) {
	return _BridgeManager.Contract.GetPeriodTotalValueThreshold(&_BridgeManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeManager *BridgeManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeManager *BridgeManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeManager.Contract.GetRoleAdmin(&_BridgeManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeManager *BridgeManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeManager.Contract.GetRoleAdmin(&_BridgeManager.CallOpts, role)
}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetTimeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTimeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetTimeWindow() (*big.Int, error) {
	return _BridgeManager.Contract.GetTimeWindow(&_BridgeManager.CallOpts)
}

// GetTimeWindow is a free data retrieval call binding the contract method 0x5566ca09.
//
// Solidity: function getTimeWindow() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetTimeWindow() (*big.Int, error) {
	return _BridgeManager.Contract.GetTimeWindow(&_BridgeManager.CallOpts)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeManager *BridgeManagerCaller) GetTokenConfig(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTokenConfig", remoteChainID, token)

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
func (_BridgeManager *BridgeManagerSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenConfig(&_BridgeManager.CallOpts, remoteChainID, token)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x6332aec6.
//
// Solidity: function getTokenConfig(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeManager *BridgeManagerCallerSession) GetTokenConfig(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenConfig(&_BridgeManager.CallOpts, remoteChainID, token)
}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetTokenCurrentScore(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTokenCurrentScore", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetTokenCurrentScore(token common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetTokenCurrentScore(&_BridgeManager.CallOpts, token)
}

// GetTokenCurrentScore is a free data retrieval call binding the contract method 0x9745e4ba.
//
// Solidity: function getTokenCurrentScore(address token) view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetTokenCurrentScore(token common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetTokenCurrentScore(&_BridgeManager.CallOpts, token)
}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeManager *BridgeManagerCaller) GetTokenMovementHistory(opts *bind.CallOpts, token common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTokenMovementHistory", token)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeManager *BridgeManagerSession) GetTokenMovementHistory(token common.Address) ([][32]byte, error) {
	return _BridgeManager.Contract.GetTokenMovementHistory(&_BridgeManager.CallOpts, token)
}

// GetTokenMovementHistory is a free data retrieval call binding the contract method 0x0e4e96bf.
//
// Solidity: function getTokenMovementHistory(address token) view returns(bytes32[] history)
func (_BridgeManager *BridgeManagerCallerSession) GetTokenMovementHistory(token common.Address) ([][32]byte, error) {
	return _BridgeManager.Contract.GetTokenMovementHistory(&_BridgeManager.CallOpts, token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address token) view returns(bool exist, uint256 price)
func (_BridgeManager *BridgeManagerCaller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTokenPrice", token)

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
func (_BridgeManager *BridgeManagerSession) GetTokenPrice(token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenPrice(&_BridgeManager.CallOpts, token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address token) view returns(bool exist, uint256 price)
func (_BridgeManager *BridgeManagerCallerSession) GetTokenPrice(token common.Address) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenPrice(&_BridgeManager.CallOpts, token)
}

// GetTokenPriceWithValue is a free data retrieval call binding the contract method 0xba8d25f8.
//
// Solidity: function getTokenPriceWithValue(address token, uint256 value) view returns(bool exist, uint256 price)
func (_BridgeManager *BridgeManagerCaller) GetTokenPriceWithValue(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTokenPriceWithValue", token, value)

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
func (_BridgeManager *BridgeManagerSession) GetTokenPriceWithValue(token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenPriceWithValue(&_BridgeManager.CallOpts, token, value)
}

// GetTokenPriceWithValue is a free data retrieval call binding the contract method 0xba8d25f8.
//
// Solidity: function getTokenPriceWithValue(address token, uint256 value) view returns(bool exist, uint256 price)
func (_BridgeManager *BridgeManagerCallerSession) GetTokenPriceWithValue(token common.Address, value *big.Int) (struct {
	Exist bool
	Price *big.Int
}, error) {
	return _BridgeManager.Contract.GetTokenPriceWithValue(&_BridgeManager.CallOpts, token, value)
}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetVerificationAmountThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getVerificationAmountThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetVerificationAmountThreshold() (*big.Int, error) {
	return _BridgeManager.Contract.GetVerificationAmountThreshold(&_BridgeManager.CallOpts)
}

// GetVerificationAmountThreshold is a free data retrieval call binding the contract method 0xbe716b0a.
//
// Solidity: function getVerificationAmountThreshold() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetVerificationAmountThreshold() (*big.Int, error) {
	return _BridgeManager.Contract.GetVerificationAmountThreshold(&_BridgeManager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeManager *BridgeManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeManager.Contract.HasRole(&_BridgeManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeManager.Contract.HasRole(&_BridgeManager.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeManager *BridgeManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeManager.Contract.SupportsInterface(&_BridgeManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeManager.Contract.SupportsInterface(&_BridgeManager.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.GrantRole(&_BridgeManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.GrantRole(&_BridgeManager.TransactOpts, role, account)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeManager *BridgeManagerTransactor) RemovePriceFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "removePriceFeed")
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeManager *BridgeManagerSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeManager.Contract.RemovePriceFeed(&_BridgeManager.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeManager *BridgeManagerTransactorSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeManager.Contract.RemovePriceFeed(&_BridgeManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeManager *BridgeManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeManager *BridgeManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RenounceRole(&_BridgeManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeManager *BridgeManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RenounceRole(&_BridgeManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RevokeRole(&_BridgeManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeManager *BridgeManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RevokeRole(&_BridgeManager.TransactOpts, role, account)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerTransactor) SetDefaultExFeeRate(opts *bind.TransactOpts, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setDefaultExFeeRate", exFeeRate)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerSession) SetDefaultExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetDefaultExFeeRate(&_BridgeManager.TransactOpts, exFeeRate)
}

// SetDefaultExFeeRate is a paid mutator transaction binding the contract method 0xad1434ca.
//
// Solidity: function setDefaultExFeeRate(uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetDefaultExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetDefaultExFeeRate(&_BridgeManager.TransactOpts, exFeeRate)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeManager *BridgeManagerTransactor) SetDefaultTokenPrice(opts *bind.TransactOpts, defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setDefaultTokenPrice", defaultTokenPrice)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeManager *BridgeManagerSession) SetDefaultTokenPrice(defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetDefaultTokenPrice(&_BridgeManager.TransactOpts, defaultTokenPrice)
}

// SetDefaultTokenPrice is a paid mutator transaction binding the contract method 0x49c9215e.
//
// Solidity: function setDefaultTokenPrice(uint256 defaultTokenPrice) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetDefaultTokenPrice(defaultTokenPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetDefaultTokenPrice(&_BridgeManager.TransactOpts, defaultTokenPrice)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerTransactor) SetExFeeRate(opts *bind.TransactOpts, token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setExFeeRate", token, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerSession) SetExFeeRate(token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetExFeeRate(&_BridgeManager.TransactOpts, token, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x5ec8f51b.
//
// Solidity: function setExFeeRate(address token, uint256 exFeeRate) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetExFeeRate(token common.Address, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetExFeeRate(&_BridgeManager.TransactOpts, token, exFeeRate)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeManager *BridgeManagerTransactor) SetExFeeRateBatch(opts *bind.TransactOpts, tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setExFeeRateBatch", tokenList, exFeeRateList)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeManager *BridgeManagerSession) SetExFeeRateBatch(tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetExFeeRateBatch(&_BridgeManager.TransactOpts, tokenList, exFeeRateList)
}

// SetExFeeRateBatch is a paid mutator transaction binding the contract method 0x020a2212.
//
// Solidity: function setExFeeRateBatch(address[] tokenList, uint256[] exFeeRateList) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetExFeeRateBatch(tokenList []common.Address, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetExFeeRateBatch(&_BridgeManager.TransactOpts, tokenList, exFeeRateList)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeManager *BridgeManagerTransactor) SetFinalizeBridgeGas(opts *bind.TransactOpts, finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setFinalizeBridgeGas", finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeManager *BridgeManagerSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetFinalizeBridgeGas(&_BridgeManager.TransactOpts, finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetFinalizeBridgeGas(&_BridgeManager.TransactOpts, finalizeBridgeGas)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeManager *BridgeManagerTransactor) SetMinimumTokenValue(opts *bind.TransactOpts, minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setMinimumTokenValue", minimumTokenValue)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeManager *BridgeManagerSession) SetMinimumTokenValue(minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetMinimumTokenValue(&_BridgeManager.TransactOpts, minimumTokenValue)
}

// SetMinimumTokenValue is a paid mutator transaction binding the contract method 0x9dc696cc.
//
// Solidity: function setMinimumTokenValue(uint256 minimumTokenValue) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetMinimumTokenValue(minimumTokenValue *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetMinimumTokenValue(&_BridgeManager.TransactOpts, minimumTokenValue)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeManager *BridgeManagerTransactor) SetPeriodTotalValueThreshold(opts *bind.TransactOpts, periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setPeriodTotalValueThreshold", periodTotalValueThreshold)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeManager *BridgeManagerSession) SetPeriodTotalValueThreshold(periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetPeriodTotalValueThreshold(&_BridgeManager.TransactOpts, periodTotalValueThreshold)
}

// SetPeriodTotalValueThreshold is a paid mutator transaction binding the contract method 0xab509f3a.
//
// Solidity: function setPeriodTotalValueThreshold(uint256 periodTotalValueThreshold) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetPeriodTotalValueThreshold(periodTotalValueThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetPeriodTotalValueThreshold(&_BridgeManager.TransactOpts, periodTotalValueThreshold)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeManager *BridgeManagerTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeManager *BridgeManagerSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetPriceFeed(&_BridgeManager.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetPriceFeed(&_BridgeManager.TransactOpts, priceFeed)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeManager *BridgeManagerTransactor) SetTimeWindow(opts *bind.TransactOpts, timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setTimeWindow", timeWindow)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeManager *BridgeManagerSession) SetTimeWindow(timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetTimeWindow(&_BridgeManager.TransactOpts, timeWindow)
}

// SetTimeWindow is a paid mutator transaction binding the contract method 0x8d331996.
//
// Solidity: function setTimeWindow(uint256 timeWindow) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetTimeWindow(timeWindow *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetTimeWindow(&_BridgeManager.TransactOpts, timeWindow)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeManager *BridgeManagerTransactor) SetVerificationAmountThreshold(opts *bind.TransactOpts, verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setVerificationAmountThreshold", verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeManager *BridgeManagerSession) SetVerificationAmountThreshold(verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetVerificationAmountThreshold(&_BridgeManager.TransactOpts, verificationAmountThreshold)
}

// SetVerificationAmountThreshold is a paid mutator transaction binding the contract method 0xc3de108b.
//
// Solidity: function setVerificationAmountThreshold(uint256 verificationAmountThreshold) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetVerificationAmountThreshold(verificationAmountThreshold *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetVerificationAmountThreshold(&_BridgeManager.TransactOpts, verificationAmountThreshold)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeManager *BridgeManagerTransactor) UpdateGasPrice(opts *bind.TransactOpts, remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "updateGasPrice", remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeManager *BridgeManagerSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateGasPrice(&_BridgeManager.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeManager *BridgeManagerTransactorSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateGasPrice(&_BridgeManager.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeManager *BridgeManagerTransactor) UpdateGasPriceBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "updateGasPriceBatch", remoteChainIDs, gasPrices)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeManager *BridgeManagerSession) UpdateGasPriceBatch(remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateGasPriceBatch(&_BridgeManager.TransactOpts, remoteChainIDs, gasPrices)
}

// UpdateGasPriceBatch is a paid mutator transaction binding the contract method 0x751ccd79.
//
// Solidity: function updateGasPriceBatch(uint256[] remoteChainIDs, uint256[] gasPrices) returns()
func (_BridgeManager *BridgeManagerTransactorSession) UpdateGasPriceBatch(remoteChainIDs []*big.Int, gasPrices []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateGasPriceBatch(&_BridgeManager.TransactOpts, remoteChainIDs, gasPrices)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(bool ok, bytes reason)
func (_BridgeManager *BridgeManagerTransactor) ValidateBridgeTokenValue(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "validateBridgeTokenValue", token, value)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(bool ok, bytes reason)
func (_BridgeManager *BridgeManagerSession) ValidateBridgeTokenValue(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.ValidateBridgeTokenValue(&_BridgeManager.TransactOpts, token, value)
}

// ValidateBridgeTokenValue is a paid mutator transaction binding the contract method 0x3f4bdec5.
//
// Solidity: function validateBridgeTokenValue(address token, uint256 value) returns(bool ok, bytes reason)
func (_BridgeManager *BridgeManagerTransactorSession) ValidateBridgeTokenValue(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.ValidateBridgeTokenValue(&_BridgeManager.TransactOpts, token, value)
}

// BridgeManagerBridgeManagerExchangeFeeUpdatedIterator is returned from FilterBridgeManagerExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeManagerExchangeFeeUpdated events raised by the BridgeManager contract.
type BridgeManagerBridgeManagerExchangeFeeUpdatedIterator struct {
	Event *BridgeManagerBridgeManagerExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeManagerExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeManagerExchangeFeeUpdated)
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
		it.Event = new(BridgeManagerBridgeManagerExchangeFeeUpdated)
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
func (it *BridgeManagerBridgeManagerExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeManagerExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeManagerExchangeFeeUpdated represents a BridgeManagerExchangeFeeUpdated event raised by the BridgeManager contract.
type BridgeManagerBridgeManagerExchangeFeeUpdated struct {
	Token     common.Address
	ExFeeRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeManagerExchangeFeeUpdated is a free log retrieval operation binding the contract event 0xe02362db7339f5c8733f9cf708855b61c1a82ae71e49239db928fe4c3ddd802b.
//
// Solidity: event BridgeManagerExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeManagerExchangeFeeUpdated(opts *bind.FilterOpts, token []common.Address) (*BridgeManagerBridgeManagerExchangeFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeManagerExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeManagerExchangeFeeUpdatedIterator{contract: _BridgeManager.contract, event: "BridgeManagerExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeManagerExchangeFeeUpdated is a free log subscription operation binding the contract event 0xe02362db7339f5c8733f9cf708855b61c1a82ae71e49239db928fe4c3ddd802b.
//
// Solidity: event BridgeManagerExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeManagerExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeManagerExchangeFeeUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeManagerExchangeFeeUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeManagerExchangeFeeUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerExchangeFeeUpdated", log); err != nil {
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

// ParseBridgeManagerExchangeFeeUpdated is a log parse operation binding the contract event 0xe02362db7339f5c8733f9cf708855b61c1a82ae71e49239db928fe4c3ddd802b.
//
// Solidity: event BridgeManagerExchangeFeeUpdated(address indexed token, uint256 exFeeRate)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeManagerExchangeFeeUpdated(log types.Log) (*BridgeManagerBridgeManagerExchangeFeeUpdated, error) {
	event := new(BridgeManagerBridgeManagerExchangeFeeUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator is returned from FilterBridgeManagerFinalizeBridgeGasSet and is used to iterate over the raw logs and unpacked data for BridgeManagerFinalizeBridgeGasSet events raised by the BridgeManager contract.
type BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator struct {
	Event *BridgeManagerBridgeManagerFinalizeBridgeGasSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeManagerFinalizeBridgeGasSet)
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
		it.Event = new(BridgeManagerBridgeManagerFinalizeBridgeGasSet)
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
func (it *BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeManagerFinalizeBridgeGasSet represents a BridgeManagerFinalizeBridgeGasSet event raised by the BridgeManager contract.
type BridgeManagerBridgeManagerFinalizeBridgeGasSet struct {
	FinalizeBridgeGas *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBridgeManagerFinalizeBridgeGasSet is a free log retrieval operation binding the contract event 0x92fa3d1383e1ab6627e21af8131fd78645136294a46271cbe6c0d22350685da7.
//
// Solidity: event BridgeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeManagerFinalizeBridgeGasSet(opts *bind.FilterOpts) (*BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeManagerFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeManagerFinalizeBridgeGasSetIterator{contract: _BridgeManager.contract, event: "BridgeManagerFinalizeBridgeGasSet", logs: logs, sub: sub}, nil
}

// WatchBridgeManagerFinalizeBridgeGasSet is a free log subscription operation binding the contract event 0x92fa3d1383e1ab6627e21af8131fd78645136294a46271cbe6c0d22350685da7.
//
// Solidity: event BridgeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeManagerFinalizeBridgeGasSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeManagerFinalizeBridgeGasSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeManagerFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeManagerFinalizeBridgeGasSet)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerFinalizeBridgeGasSet", log); err != nil {
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

// ParseBridgeManagerFinalizeBridgeGasSet is a log parse operation binding the contract event 0x92fa3d1383e1ab6627e21af8131fd78645136294a46271cbe6c0d22350685da7.
//
// Solidity: event BridgeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeManagerFinalizeBridgeGasSet(log types.Log) (*BridgeManagerBridgeManagerFinalizeBridgeGasSet, error) {
	event := new(BridgeManagerBridgeManagerFinalizeBridgeGasSet)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerFinalizeBridgeGasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerBridgeManagerGasPriceUpdatedIterator is returned from FilterBridgeManagerGasPriceUpdated and is used to iterate over the raw logs and unpacked data for BridgeManagerGasPriceUpdated events raised by the BridgeManager contract.
type BridgeManagerBridgeManagerGasPriceUpdatedIterator struct {
	Event *BridgeManagerBridgeManagerGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeManagerGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeManagerGasPriceUpdated)
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
		it.Event = new(BridgeManagerBridgeManagerGasPriceUpdated)
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
func (it *BridgeManagerBridgeManagerGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeManagerGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeManagerGasPriceUpdated represents a BridgeManagerGasPriceUpdated event raised by the BridgeManager contract.
type BridgeManagerBridgeManagerGasPriceUpdated struct {
	RemoteChainID *big.Int
	GasPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeManagerGasPriceUpdated is a free log retrieval operation binding the contract event 0x250cb85ef4edc2a5afdc6bf99c82fd11536c37fe9ad56db2866db5a24b14d848.
//
// Solidity: event BridgeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeManagerGasPriceUpdated(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BridgeManagerBridgeManagerGasPriceUpdatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeManagerGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeManagerGasPriceUpdatedIterator{contract: _BridgeManager.contract, event: "BridgeManagerGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeManagerGasPriceUpdated is a free log subscription operation binding the contract event 0x250cb85ef4edc2a5afdc6bf99c82fd11536c37fe9ad56db2866db5a24b14d848.
//
// Solidity: event BridgeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeManagerGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeManagerGasPriceUpdated, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeManagerGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeManagerGasPriceUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerGasPriceUpdated", log); err != nil {
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

// ParseBridgeManagerGasPriceUpdated is a log parse operation binding the contract event 0x250cb85ef4edc2a5afdc6bf99c82fd11536c37fe9ad56db2866db5a24b14d848.
//
// Solidity: event BridgeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeManagerGasPriceUpdated(log types.Log) (*BridgeManagerBridgeManagerGasPriceUpdated, error) {
	event := new(BridgeManagerBridgeManagerGasPriceUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerBridgeManagerPriceFeedUpdatedIterator is returned from FilterBridgeManagerPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for BridgeManagerPriceFeedUpdated events raised by the BridgeManager contract.
type BridgeManagerBridgeManagerPriceFeedUpdatedIterator struct {
	Event *BridgeManagerBridgeManagerPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeManagerPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeManagerPriceFeedUpdated)
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
		it.Event = new(BridgeManagerBridgeManagerPriceFeedUpdated)
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
func (it *BridgeManagerBridgeManagerPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeManagerPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeManagerPriceFeedUpdated represents a BridgeManagerPriceFeedUpdated event raised by the BridgeManager contract.
type BridgeManagerBridgeManagerPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeManagerPriceFeedUpdated is a free log retrieval operation binding the contract event 0x6433071fe185ca8d6b5d1f327404fcc1ab697b7e9f29635aa9b0575ec51e5eaf.
//
// Solidity: event BridgeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeManagerPriceFeedUpdated(opts *bind.FilterOpts, priceFeed []common.Address) (*BridgeManagerBridgeManagerPriceFeedUpdatedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeManagerPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeManagerPriceFeedUpdatedIterator{contract: _BridgeManager.contract, event: "BridgeManagerPriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeManagerPriceFeedUpdated is a free log subscription operation binding the contract event 0x6433071fe185ca8d6b5d1f327404fcc1ab697b7e9f29635aa9b0575ec51e5eaf.
//
// Solidity: event BridgeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeManagerPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeManagerPriceFeedUpdated, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeManagerPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeManagerPriceFeedUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerPriceFeedUpdated", log); err != nil {
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

// ParseBridgeManagerPriceFeedUpdated is a log parse operation binding the contract event 0x6433071fe185ca8d6b5d1f327404fcc1ab697b7e9f29635aa9b0575ec51e5eaf.
//
// Solidity: event BridgeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeManagerPriceFeedUpdated(log types.Log) (*BridgeManagerBridgeManagerPriceFeedUpdated, error) {
	event := new(BridgeManagerBridgeManagerPriceFeedUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeManagerPriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerDefaultTokenPriceSetIterator is returned from FilterDefaultTokenPriceSet and is used to iterate over the raw logs and unpacked data for DefaultTokenPriceSet events raised by the BridgeManager contract.
type BridgeManagerDefaultTokenPriceSetIterator struct {
	Event *BridgeManagerDefaultTokenPriceSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerDefaultTokenPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerDefaultTokenPriceSet)
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
		it.Event = new(BridgeManagerDefaultTokenPriceSet)
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
func (it *BridgeManagerDefaultTokenPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerDefaultTokenPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerDefaultTokenPriceSet represents a DefaultTokenPriceSet event raised by the BridgeManager contract.
type BridgeManagerDefaultTokenPriceSet struct {
	DefaultTokenPrice *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDefaultTokenPriceSet is a free log retrieval operation binding the contract event 0x1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a.
//
// Solidity: event DefaultTokenPriceSet(uint256 defaultTokenPrice)
func (_BridgeManager *BridgeManagerFilterer) FilterDefaultTokenPriceSet(opts *bind.FilterOpts) (*BridgeManagerDefaultTokenPriceSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "DefaultTokenPriceSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerDefaultTokenPriceSetIterator{contract: _BridgeManager.contract, event: "DefaultTokenPriceSet", logs: logs, sub: sub}, nil
}

// WatchDefaultTokenPriceSet is a free log subscription operation binding the contract event 0x1b0009bc3090df60db5b8e1a7cc239057117cccff0af0b1d39d8fba2af21610a.
//
// Solidity: event DefaultTokenPriceSet(uint256 defaultTokenPrice)
func (_BridgeManager *BridgeManagerFilterer) WatchDefaultTokenPriceSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerDefaultTokenPriceSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "DefaultTokenPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerDefaultTokenPriceSet)
				if err := _BridgeManager.contract.UnpackLog(event, "DefaultTokenPriceSet", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseDefaultTokenPriceSet(log types.Log) (*BridgeManagerDefaultTokenPriceSet, error) {
	event := new(BridgeManagerDefaultTokenPriceSet)
	if err := _BridgeManager.contract.UnpackLog(event, "DefaultTokenPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerMinimumTokenValueSetIterator is returned from FilterMinimumTokenValueSet and is used to iterate over the raw logs and unpacked data for MinimumTokenValueSet events raised by the BridgeManager contract.
type BridgeManagerMinimumTokenValueSetIterator struct {
	Event *BridgeManagerMinimumTokenValueSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerMinimumTokenValueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerMinimumTokenValueSet)
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
		it.Event = new(BridgeManagerMinimumTokenValueSet)
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
func (it *BridgeManagerMinimumTokenValueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerMinimumTokenValueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerMinimumTokenValueSet represents a MinimumTokenValueSet event raised by the BridgeManager contract.
type BridgeManagerMinimumTokenValueSet struct {
	MinimumTokenValue *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMinimumTokenValueSet is a free log retrieval operation binding the contract event 0x0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e9.
//
// Solidity: event MinimumTokenValueSet(uint256 minimumTokenValue)
func (_BridgeManager *BridgeManagerFilterer) FilterMinimumTokenValueSet(opts *bind.FilterOpts) (*BridgeManagerMinimumTokenValueSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "MinimumTokenValueSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerMinimumTokenValueSetIterator{contract: _BridgeManager.contract, event: "MinimumTokenValueSet", logs: logs, sub: sub}, nil
}

// WatchMinimumTokenValueSet is a free log subscription operation binding the contract event 0x0b24652f50d73c1d6eb8f51f1258ebae14d5b1427f0a11e3bae742b044eab1e9.
//
// Solidity: event MinimumTokenValueSet(uint256 minimumTokenValue)
func (_BridgeManager *BridgeManagerFilterer) WatchMinimumTokenValueSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerMinimumTokenValueSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "MinimumTokenValueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerMinimumTokenValueSet)
				if err := _BridgeManager.contract.UnpackLog(event, "MinimumTokenValueSet", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseMinimumTokenValueSet(log types.Log) (*BridgeManagerMinimumTokenValueSet, error) {
	event := new(BridgeManagerMinimumTokenValueSet)
	if err := _BridgeManager.contract.UnpackLog(event, "MinimumTokenValueSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerPeriodTotalValueThresholdSetIterator is returned from FilterPeriodTotalValueThresholdSet and is used to iterate over the raw logs and unpacked data for PeriodTotalValueThresholdSet events raised by the BridgeManager contract.
type BridgeManagerPeriodTotalValueThresholdSetIterator struct {
	Event *BridgeManagerPeriodTotalValueThresholdSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerPeriodTotalValueThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerPeriodTotalValueThresholdSet)
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
		it.Event = new(BridgeManagerPeriodTotalValueThresholdSet)
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
func (it *BridgeManagerPeriodTotalValueThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerPeriodTotalValueThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerPeriodTotalValueThresholdSet represents a PeriodTotalValueThresholdSet event raised by the BridgeManager contract.
type BridgeManagerPeriodTotalValueThresholdSet struct {
	PeriodTotalValueThreshold *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPeriodTotalValueThresholdSet is a free log retrieval operation binding the contract event 0x7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e2.
//
// Solidity: event PeriodTotalValueThresholdSet(uint256 periodTotalValueThreshold)
func (_BridgeManager *BridgeManagerFilterer) FilterPeriodTotalValueThresholdSet(opts *bind.FilterOpts) (*BridgeManagerPeriodTotalValueThresholdSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "PeriodTotalValueThresholdSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerPeriodTotalValueThresholdSetIterator{contract: _BridgeManager.contract, event: "PeriodTotalValueThresholdSet", logs: logs, sub: sub}, nil
}

// WatchPeriodTotalValueThresholdSet is a free log subscription operation binding the contract event 0x7b5e4eb1bcf6f11b51639eca267f8ef08f9780c190e091f79a5fb810018b34e2.
//
// Solidity: event PeriodTotalValueThresholdSet(uint256 periodTotalValueThreshold)
func (_BridgeManager *BridgeManagerFilterer) WatchPeriodTotalValueThresholdSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerPeriodTotalValueThresholdSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "PeriodTotalValueThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerPeriodTotalValueThresholdSet)
				if err := _BridgeManager.contract.UnpackLog(event, "PeriodTotalValueThresholdSet", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParsePeriodTotalValueThresholdSet(log types.Log) (*BridgeManagerPeriodTotalValueThresholdSet, error) {
	event := new(BridgeManagerPeriodTotalValueThresholdSet)
	if err := _BridgeManager.contract.UnpackLog(event, "PeriodTotalValueThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeManager contract.
type BridgeManagerRoleAdminChangedIterator struct {
	Event *BridgeManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerRoleAdminChanged)
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
		it.Event = new(BridgeManagerRoleAdminChanged)
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
func (it *BridgeManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeManager contract.
type BridgeManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeManager *BridgeManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerRoleAdminChangedIterator{contract: _BridgeManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeManager *BridgeManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerRoleAdminChanged)
				if err := _BridgeManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeManagerRoleAdminChanged, error) {
	event := new(BridgeManagerRoleAdminChanged)
	if err := _BridgeManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeManager contract.
type BridgeManagerRoleGrantedIterator struct {
	Event *BridgeManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerRoleGranted)
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
		it.Event = new(BridgeManagerRoleGranted)
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
func (it *BridgeManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerRoleGranted represents a RoleGranted event raised by the BridgeManager contract.
type BridgeManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeManager *BridgeManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerRoleGrantedIterator{contract: _BridgeManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeManager *BridgeManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerRoleGranted)
				if err := _BridgeManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseRoleGranted(log types.Log) (*BridgeManagerRoleGranted, error) {
	event := new(BridgeManagerRoleGranted)
	if err := _BridgeManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeManager contract.
type BridgeManagerRoleRevokedIterator struct {
	Event *BridgeManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerRoleRevoked)
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
		it.Event = new(BridgeManagerRoleRevoked)
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
func (it *BridgeManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerRoleRevoked represents a RoleRevoked event raised by the BridgeManager contract.
type BridgeManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeManager *BridgeManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerRoleRevokedIterator{contract: _BridgeManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeManager *BridgeManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerRoleRevoked)
				if err := _BridgeManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseRoleRevoked(log types.Log) (*BridgeManagerRoleRevoked, error) {
	event := new(BridgeManagerRoleRevoked)
	if err := _BridgeManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerTimeWindowSetIterator is returned from FilterTimeWindowSet and is used to iterate over the raw logs and unpacked data for TimeWindowSet events raised by the BridgeManager contract.
type BridgeManagerTimeWindowSetIterator struct {
	Event *BridgeManagerTimeWindowSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerTimeWindowSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerTimeWindowSet)
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
		it.Event = new(BridgeManagerTimeWindowSet)
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
func (it *BridgeManagerTimeWindowSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerTimeWindowSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerTimeWindowSet represents a TimeWindowSet event raised by the BridgeManager contract.
type BridgeManagerTimeWindowSet struct {
	TimeWindow *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTimeWindowSet is a free log retrieval operation binding the contract event 0x52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a4.
//
// Solidity: event TimeWindowSet(uint256 timeWindow)
func (_BridgeManager *BridgeManagerFilterer) FilterTimeWindowSet(opts *bind.FilterOpts) (*BridgeManagerTimeWindowSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "TimeWindowSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerTimeWindowSetIterator{contract: _BridgeManager.contract, event: "TimeWindowSet", logs: logs, sub: sub}, nil
}

// WatchTimeWindowSet is a free log subscription operation binding the contract event 0x52c788384b8860c02c4a87558ab1721141442a3daa4a338f8e455a5d74e749a4.
//
// Solidity: event TimeWindowSet(uint256 timeWindow)
func (_BridgeManager *BridgeManagerFilterer) WatchTimeWindowSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerTimeWindowSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "TimeWindowSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerTimeWindowSet)
				if err := _BridgeManager.contract.UnpackLog(event, "TimeWindowSet", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseTimeWindowSet(log types.Log) (*BridgeManagerTimeWindowSet, error) {
	event := new(BridgeManagerTimeWindowSet)
	if err := _BridgeManager.contract.UnpackLog(event, "TimeWindowSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerVerificationAmountThresholdSetIterator is returned from FilterVerificationAmountThresholdSet and is used to iterate over the raw logs and unpacked data for VerificationAmountThresholdSet events raised by the BridgeManager contract.
type BridgeManagerVerificationAmountThresholdSetIterator struct {
	Event *BridgeManagerVerificationAmountThresholdSet // Event containing the contract specifics and raw log

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
func (it *BridgeManagerVerificationAmountThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerVerificationAmountThresholdSet)
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
		it.Event = new(BridgeManagerVerificationAmountThresholdSet)
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
func (it *BridgeManagerVerificationAmountThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerVerificationAmountThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerVerificationAmountThresholdSet represents a VerificationAmountThresholdSet event raised by the BridgeManager contract.
type BridgeManagerVerificationAmountThresholdSet struct {
	VerificationAmountThreshold *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterVerificationAmountThresholdSet is a free log retrieval operation binding the contract event 0xf905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2.
//
// Solidity: event VerificationAmountThresholdSet(uint256 verificationAmountThreshold)
func (_BridgeManager *BridgeManagerFilterer) FilterVerificationAmountThresholdSet(opts *bind.FilterOpts) (*BridgeManagerVerificationAmountThresholdSetIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "VerificationAmountThresholdSet")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerVerificationAmountThresholdSetIterator{contract: _BridgeManager.contract, event: "VerificationAmountThresholdSet", logs: logs, sub: sub}, nil
}

// WatchVerificationAmountThresholdSet is a free log subscription operation binding the contract event 0xf905277f1e1bedd48bc8ef6326f0ba77b592484e9656377f943124f6ebd701a2.
//
// Solidity: event VerificationAmountThresholdSet(uint256 verificationAmountThreshold)
func (_BridgeManager *BridgeManagerFilterer) WatchVerificationAmountThresholdSet(opts *bind.WatchOpts, sink chan<- *BridgeManagerVerificationAmountThresholdSet) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "VerificationAmountThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerVerificationAmountThresholdSet)
				if err := _BridgeManager.contract.UnpackLog(event, "VerificationAmountThresholdSet", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseVerificationAmountThresholdSet(log types.Log) (*BridgeManagerVerificationAmountThresholdSet, error) {
	event := new(BridgeManagerVerificationAmountThresholdSet)
	if err := _BridgeManager.contract.UnpackLog(event, "VerificationAmountThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
