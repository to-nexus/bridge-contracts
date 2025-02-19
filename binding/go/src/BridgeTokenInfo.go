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

// BridgeTokenInfoMetaData contains all meta data concerning the BridgeTokenInfo contract.
var BridgeTokenInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"addTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"tokenInfoList\",\"type\":\"tuple[]\"}],\"name\":\"addTokenInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenInfoList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"removeTokenInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"updateTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeTokenInfo.TokenInfo[]\",\"name\":\"tokenInfoList\",\"type\":\"tuple[]\"}],\"name\":\"updateTokenInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenInfoRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"TokenInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeTokenInfoInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"aea97f0a": "addTokenInfo(address,uint256,uint256,uint256)",
		"0e17b19e": "addTokenInfoMany((address,uint256,uint256,uint256)[])",
		"9fdf1c6a": "allTokenInfo()",
		"6ff97f1d": "allTokens()",
		"6e908ca3": "calculate(address,uint256)",
		"5dbe47e8": "contains(address)",
		"96ce0795": "denominator()",
		"1f69565f": "getTokenInfo(address)",
		"1281cfdb": "getTokenInfoList(address[])",
		"8da5cb5b": "owner()",
		"95e4c77f": "removeTokenInfo(address)",
		"163f39ff": "removeTokenInfoMany(address[])",
		"715018a6": "renounceOwnership()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"e70a1b26": "tokenInfoByIndex(uint256)",
		"d92fc67b": "tokensLength()",
		"f2fde38b": "transferOwnership(address)",
		"5e78c722": "updateTokenInfo(address,uint256,uint256,uint256)",
		"bc45d21c": "updateTokenInfoMany((address,uint256,uint256,uint256)[])",
	},
	Bin: "0x60806040526103e86035553480156014575f5ffd5b503380603957604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6040816045565b506094565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611225806100a15f395ff3fe608060405234801561000f575f5ffd5b506004361061011c575f3560e01c8063715018a6116100a9578063aea97f0a1161006e578063aea97f0a1461027a578063bc45d21c1461028d578063d92fc67b146102a0578063e70a1b26146102a8578063f2fde38b146102bb575f5ffd5b8063715018a6146102355780638da5cb5b1461023d57806395e4c77f1461024d57806396ce0795146102605780639fdf1c6a14610272575f5ffd5b80634f6ccce7116100ef5780634f6ccce7146101915780635dbe47e8146101bc5780635e78c722146101df5780636e908ca3146101f25780636ff97f1d14610220575f5ffd5b80630e17b19e146101205780631281cfdb14610135578063163f39ff1461015e5780631f69565f14610171575b5f5ffd5b61013361012e366004610e90565b6102ce565b005b610148610143366004610f66565b610364565b6040516101559190610ff4565b60405180910390f35b61013361016c366004610f66565b610420565b61018461017f366004611068565b610455565b6040516101559190611081565b6101a461019f3660046110b5565b6104e9565b6040516001600160a01b039091168152602001610155565b6101cf6101ca366004611068565b6104fb565b6040519015158152602001610155565b6101336101ed3660046110cc565b610507565b610205610200366004611102565b6105b5565b60408051938452602084019290925290820152606001610155565b6102286105f9565b604051610155919061112a565b61013361065d565b5f546001600160a01b03166101a4565b61013361025b366004611068565b610670565b6035545b604051908152602001610155565b6101486106db565b6101336102883660046110cc565b6107d1565b61013361029b366004610e90565b61087a565b61026461090c565b6101846102b63660046110b5565b61091c565b6101336102c9366004611068565b610985565b5f5b8151811015610360576103588282815181106102ee576102ee61116a565b60200260200101515f015183838151811061030b5761030b61116a565b6020026020010151602001518484815181106103295761032961116a565b6020026020010151604001518585815181106103475761034761116a565b6020026020010151606001516107d1565b6001016102d0565b5050565b60605f61036f61090c565b90505f8167ffffffffffffffff81111561038b5761038b610de4565b6040519080825280602002602001820160405280156103c457816020015b6103b1610db7565b8152602001906001900390816103a95790505b5090505f5b82811015610418576103f38582815181106103e6576103e661116a565b6020026020010151610455565b8282815181106104055761040561116a565b60209081029190910101526001016103c9565b509392505050565b5f5b81518110156103605761044d8282815181106104405761044061116a565b6020026020010151610670565b600101610422565b61045d610db7565b610466826104fb565b156104bb57506001600160a01b039081165f9081526034602090815260409182902082516080810184528154909416845260018101549184019190915260028101549183019190915260030154606082015290565b50604080516080810182526001600160a01b0390921682525f60208301819052908201819052606082015290565b5f6104f56001836109c2565b92915050565b5f6104f56001836109d4565b610510846104fb565b8490610540576040516321180f2d60e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b506001600160a01b0384165f81815260346020908152604091829020600181018790556002810186905560030184905581518681529081018590529081018390527fe7bb7f0ee932b644e52c34df47e4bf3b30372d6455e8c87d54517ac43a3a123a906060015b60405180910390a250505050565b5f5f5f5f6105c286610455565b9050806020015181604001516035548360600151886105e19190611192565b6105eb91906111a9565b935093509350509250925092565b60605f61060660016109f5565b90508067ffffffffffffffff81111561062157610621610de4565b60405190808252806020026020018201604052801561064a578160200160208202803683370190505b50915061065760016109fe565b91505090565b610665610a0a565b61066e5f610a36565b565b61067981610a85565b6001600160a01b0381165f8181526034602052604080822080546001600160a01b03191681556001810183905560028101839055600301829055517fce718809ef43b78b8d60e18e71da49dfe61d60cc8dc80c68a0574d393c13d8459190a250565b60605f6106e661090c565b90505f8167ffffffffffffffff81111561070257610702610de4565b60405190808252806020026020018201604052801561073b57816020015b610728610db7565b8152602001906001900390816107205790505b5090505f5b828110156107ca5760345f610754836104e9565b6001600160a01b03908116825260208083019390935260409182015f20825160808101845281549092168252600181015493820193909352600283015491810191909152600390910154606082015282518390839081106107b7576107b761116a565b6020908102919091010152600101610740565b5092915050565b6107da84610b33565b604080516080810182526001600160a01b03868116808352602080840188815284860188815260608087018981525f86815260348652899020975188546001600160a01b031916971696909617875591516001870155516002860155925160039094019390935583518781529283018690529282018490527fe7bb7f0ee932b644e52c34df47e4bf3b30372d6455e8c87d54517ac43a3a123a91016105a7565b5f5b81518110156103605761090482828151811061089a5761089a61116a565b60200260200101515f01518383815181106108b7576108b761116a565b6020026020010151602001518484815181106108d5576108d561116a565b6020026020010151604001518585815181106108f3576108f361116a565b602002602001015160600151610507565b60010161087c565b5f61091760016109f5565b905090565b610924610db7565b60345f610930846104e9565b6001600160a01b03908116825260208083019390935260409182015f20825160808101845281549092168252600181015493820193909352600283015491810191909152600390910154606082015292915050565b61098d610a0a565b6001600160a01b0381166109b657604051631e4fbdf760e01b81525f6004820152602401610537565b6109bf81610a36565b50565b5f6109cd8383610be1565b9392505050565b6001600160a01b0381165f90815260018301602052604081205415156109cd565b5f6104f5825490565b60605f6109cd83610c07565b5f546001600160a01b0316331461066e5760405163118cdaa760e01b8152336004820152602401610537565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b806001600160a01b038116610ac5576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610537565b610ad0600183610c60565b8290610afb576040516340da71e560e01b81526001600160a01b039091166004820152602401610537565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b806001600160a01b038116610b73576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610537565b610b7e600183610c74565b8290610ba9576040516351eccfe160e11b81526001600160a01b039091166004820152602401610537565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b5f825f018281548110610bf657610bf661116a565b905f5260205f200154905092915050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610c5457602002820191905f5260205f20905b815481526020019060010190808311610c40575b50505050509050919050565b5f6109cd836001600160a01b038416610c88565b5f6109cd836001600160a01b038416610d6b565b5f8181526001830160205260408120548015610d62575f610caa6001836111c8565b85549091505f90610cbd906001906111c8565b9050808214610d1c575f865f018281548110610cdb57610cdb61116a565b905f5260205f200154905080875f018481548110610cfb57610cfb61116a565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610d2d57610d2d6111db565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506104f5565b5f9150506104f5565b5f818152600183016020526040812054610db057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556104f5565b505f6104f5565b60405180608001604052805f6001600160a01b031681526020015f81526020015f81526020015f81525090565b634e487b7160e01b5f52604160045260245ffd5b6040516080810167ffffffffffffffff81118282101715610e1b57610e1b610de4565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e4a57610e4a610de4565b604052919050565b5f67ffffffffffffffff821115610e6b57610e6b610de4565b5060051b60200190565b80356001600160a01b0381168114610e8b575f5ffd5b919050565b5f60208284031215610ea0575f5ffd5b813567ffffffffffffffff811115610eb6575f5ffd5b8201601f81018413610ec6575f5ffd5b8035610ed9610ed482610e52565b610e21565b8082825260208201915060208360071b850101925086831115610efa575f5ffd5b6020840193505b82841015610f5c5760808488031215610f18575f5ffd5b610f20610df8565b610f2985610e75565b81526020858101358183015260408087013590830152606080870135908301529083526080909401939190910190610f01565b9695505050505050565b5f60208284031215610f76575f5ffd5b813567ffffffffffffffff811115610f8c575f5ffd5b8201601f81018413610f9c575f5ffd5b8035610faa610ed482610e52565b8082825260208201915060208360051b850101925086831115610fcb575f5ffd5b6020840193505b82841015610f5c57610fe384610e75565b825260209384019390910190610fd2565b602080825282518282018190525f918401906040840190835b8181101561105d5761104783855180516001600160a01b031682526020808201519083015260408082015190830152606090810151910152565b602093909301926080929092019160010161100d565b509095945050505050565b5f60208284031215611078575f5ffd5b6109cd82610e75565b81516001600160a01b03168152602080830151908201526040808301519082015260608083015190820152608081016104f5565b5f602082840312156110c5575f5ffd5b5035919050565b5f5f5f5f608085870312156110df575f5ffd5b6110e885610e75565b966020860135965060408601359560600135945092505050565b5f5f60408385031215611113575f5ffd5b61111c83610e75565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b8181101561105d5783516001600160a01b0316835260209384019390920191600101611143565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176104f5576104f561117e565b5f826111c357634e487b7160e01b5f52601260045260245ffd5b500490565b818103818111156104f5576104f561117e565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220be4aae32ed0f88282ff6207224f9e8d57ffcfd0c70e9c2762adbd8ba6cccf55d64736f6c634300081c0033",
}

// BridgeTokenInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTokenInfoMetaData.ABI instead.
var BridgeTokenInfoABI = BridgeTokenInfoMetaData.ABI

// Deprecated: Use BridgeTokenInfoMetaData.Sigs instead.
// BridgeTokenInfoFuncSigs maps the 4-byte function signature to its string representation.
var BridgeTokenInfoFuncSigs = BridgeTokenInfoMetaData.Sigs

// BridgeTokenInfoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTokenInfoMetaData.Bin instead.
var BridgeTokenInfoBin = BridgeTokenInfoMetaData.Bin

// DeployBridgeTokenInfo deploys a new Ethereum contract, binding an instance of BridgeTokenInfo to it.
func DeployBridgeTokenInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTokenInfo, error) {
	parsed, err := BridgeTokenInfoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTokenInfoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTokenInfo{BridgeTokenInfoCaller: BridgeTokenInfoCaller{contract: contract}, BridgeTokenInfoTransactor: BridgeTokenInfoTransactor{contract: contract}, BridgeTokenInfoFilterer: BridgeTokenInfoFilterer{contract: contract}}, nil
}

// BridgeTokenInfo is an auto generated Go binding around an Ethereum contract.
type BridgeTokenInfo struct {
	BridgeTokenInfoCaller     // Read-only binding to the contract
	BridgeTokenInfoTransactor // Write-only binding to the contract
	BridgeTokenInfoFilterer   // Log filterer for contract events
}

// BridgeTokenInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTokenInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTokenInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTokenInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTokenInfoSession struct {
	Contract     *BridgeTokenInfo  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTokenInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTokenInfoCallerSession struct {
	Contract *BridgeTokenInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BridgeTokenInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTokenInfoTransactorSession struct {
	Contract     *BridgeTokenInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgeTokenInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTokenInfoRaw struct {
	Contract *BridgeTokenInfo // Generic contract binding to access the raw methods on
}

// BridgeTokenInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTokenInfoCallerRaw struct {
	Contract *BridgeTokenInfoCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTokenInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTokenInfoTransactorRaw struct {
	Contract *BridgeTokenInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTokenInfo creates a new instance of BridgeTokenInfo, bound to a specific deployed contract.
func NewBridgeTokenInfo(address common.Address, backend bind.ContractBackend) (*BridgeTokenInfo, error) {
	contract, err := bindBridgeTokenInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfo{BridgeTokenInfoCaller: BridgeTokenInfoCaller{contract: contract}, BridgeTokenInfoTransactor: BridgeTokenInfoTransactor{contract: contract}, BridgeTokenInfoFilterer: BridgeTokenInfoFilterer{contract: contract}}, nil
}

// NewBridgeTokenInfoCaller creates a new read-only instance of BridgeTokenInfo, bound to a specific deployed contract.
func NewBridgeTokenInfoCaller(address common.Address, caller bind.ContractCaller) (*BridgeTokenInfoCaller, error) {
	contract, err := bindBridgeTokenInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoCaller{contract: contract}, nil
}

// NewBridgeTokenInfoTransactor creates a new write-only instance of BridgeTokenInfo, bound to a specific deployed contract.
func NewBridgeTokenInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTokenInfoTransactor, error) {
	contract, err := bindBridgeTokenInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoTransactor{contract: contract}, nil
}

// NewBridgeTokenInfoFilterer creates a new log filterer instance of BridgeTokenInfo, bound to a specific deployed contract.
func NewBridgeTokenInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTokenInfoFilterer, error) {
	contract, err := bindBridgeTokenInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoFilterer{contract: contract}, nil
}

// bindBridgeTokenInfo binds a generic wrapper to an already deployed contract.
func bindBridgeTokenInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTokenInfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTokenInfo *BridgeTokenInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTokenInfo.Contract.BridgeTokenInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTokenInfo *BridgeTokenInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.BridgeTokenInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTokenInfo *BridgeTokenInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.BridgeTokenInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTokenInfo *BridgeTokenInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTokenInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTokenInfo *BridgeTokenInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTokenInfo *BridgeTokenInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.contract.Transact(opts, method, params...)
}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoCaller) AllTokenInfo(opts *bind.CallOpts) ([]IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "allTokenInfo")

	if err != nil {
		return *new([]IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeTokenInfoTokenInfo)).(*[]IBridgeTokenInfoTokenInfo)

	return out0, err

}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.AllTokenInfo(&_BridgeTokenInfo.CallOpts)
}

// AllTokenInfo is a free data retrieval call binding the contract method 0x9fdf1c6a.
//
// Solidity: function allTokenInfo() view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) AllTokenInfo() ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.AllTokenInfo(&_BridgeTokenInfo.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeTokenInfo *BridgeTokenInfoSession) AllTokens() ([]common.Address, error) {
	return _BridgeTokenInfo.Contract.AllTokens(&_BridgeTokenInfo.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) AllTokens() ([]common.Address, error) {
	return _BridgeTokenInfo.Contract.AllTokens(&_BridgeTokenInfo.CallOpts)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "calculate", token, value)

	outstruct := new(struct {
		Minimum *big.Int
		Gas     *big.Int
		Service *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minimum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Service = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeTokenInfo.Contract.Calculate(&_BridgeTokenInfo.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 service)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeTokenInfo.Contract.Calculate(&_BridgeTokenInfo.CallOpts, token, value)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Contains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "contains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Contains(token common.Address) (bool, error) {
	return _BridgeTokenInfo.Contract.Contains(&_BridgeTokenInfo.CallOpts, token)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Contains(token common.Address) (bool, error) {
	return _BridgeTokenInfo.Contract.Contains(&_BridgeTokenInfo.CallOpts, token)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Denominator() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.Denominator(&_BridgeTokenInfo.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Denominator() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.Denominator(&_BridgeTokenInfo.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoCaller) GetTokenInfo(opts *bind.CallOpts, token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "getTokenInfo", token)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfo(&_BridgeTokenInfo.CallOpts, token)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) GetTokenInfo(token common.Address) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfo(&_BridgeTokenInfo.CallOpts, token)
}

// GetTokenInfoList is a free data retrieval call binding the contract method 0x1281cfdb.
//
// Solidity: function getTokenInfoList(address[] tokens) view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoCaller) GetTokenInfoList(opts *bind.CallOpts, tokens []common.Address) ([]IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "getTokenInfoList", tokens)

	if err != nil {
		return *new([]IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeTokenInfoTokenInfo)).(*[]IBridgeTokenInfoTokenInfo)

	return out0, err

}

// GetTokenInfoList is a free data retrieval call binding the contract method 0x1281cfdb.
//
// Solidity: function getTokenInfoList(address[] tokens) view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoSession) GetTokenInfoList(tokens []common.Address) ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfoList(&_BridgeTokenInfo.CallOpts, tokens)
}

// GetTokenInfoList is a free data retrieval call binding the contract method 0x1281cfdb.
//
// Solidity: function getTokenInfoList(address[] tokens) view returns((address,uint256,uint256,uint256)[])
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) GetTokenInfoList(tokens []common.Address) ([]IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfoList(&_BridgeTokenInfo.CallOpts, tokens)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Owner() (common.Address, error) {
	return _BridgeTokenInfo.Contract.Owner(&_BridgeTokenInfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Owner() (common.Address, error) {
	return _BridgeTokenInfo.Contract.Owner(&_BridgeTokenInfo.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) TokenByIndex(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "tokenByIndex", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeTokenInfo.Contract.TokenByIndex(&_BridgeTokenInfo.CallOpts, i)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeTokenInfo.Contract.TokenByIndex(&_BridgeTokenInfo.CallOpts, i)
}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoCaller) TokenInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "tokenInfoByIndex", index)

	if err != nil {
		return *new(IBridgeTokenInfoTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeTokenInfoTokenInfo)).(*IBridgeTokenInfoTokenInfo)

	return out0, err

}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.TokenInfoByIndex(&_BridgeTokenInfo.CallOpts, index)
}

// TokenInfoByIndex is a free data retrieval call binding the contract method 0xe70a1b26.
//
// Solidity: function tokenInfoByIndex(uint256 index) view returns((address,uint256,uint256,uint256))
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) TokenInfoByIndex(index *big.Int) (IBridgeTokenInfoTokenInfo, error) {
	return _BridgeTokenInfo.Contract.TokenInfoByIndex(&_BridgeTokenInfo.CallOpts, index)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoSession) TokensLength() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.TokensLength(&_BridgeTokenInfo.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) TokensLength() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.TokensLength(&_BridgeTokenInfo.CallOpts)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xaea97f0a.
//
// Solidity: function addTokenInfo(address token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) AddTokenInfo(opts *bind.TransactOpts, token common.Address, minimumValue *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "addTokenInfo", token, minimumValue, gasFee, serviceFee)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xaea97f0a.
//
// Solidity: function addTokenInfo(address token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) AddTokenInfo(token common.Address, minimumValue *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimumValue, gasFee, serviceFee)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0xaea97f0a.
//
// Solidity: function addTokenInfo(address token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) AddTokenInfo(token common.Address, minimumValue *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimumValue, gasFee, serviceFee)
}

// AddTokenInfoMany is a paid mutator transaction binding the contract method 0x0e17b19e.
//
// Solidity: function addTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) AddTokenInfoMany(opts *bind.TransactOpts, tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "addTokenInfoMany", tokenInfoList)
}

// AddTokenInfoMany is a paid mutator transaction binding the contract method 0x0e17b19e.
//
// Solidity: function addTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) AddTokenInfoMany(tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfoMany(&_BridgeTokenInfo.TransactOpts, tokenInfoList)
}

// AddTokenInfoMany is a paid mutator transaction binding the contract method 0x0e17b19e.
//
// Solidity: function addTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) AddTokenInfoMany(tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfoMany(&_BridgeTokenInfo.TransactOpts, tokenInfoList)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) RemoveTokenInfo(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "removeTokenInfo", token)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) RemoveTokenInfo(token common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemoveTokenInfo(&_BridgeTokenInfo.TransactOpts, token)
}

// RemoveTokenInfo is a paid mutator transaction binding the contract method 0x95e4c77f.
//
// Solidity: function removeTokenInfo(address token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) RemoveTokenInfo(token common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemoveTokenInfo(&_BridgeTokenInfo.TransactOpts, token)
}

// RemoveTokenInfoMany is a paid mutator transaction binding the contract method 0x163f39ff.
//
// Solidity: function removeTokenInfoMany(address[] token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) RemoveTokenInfoMany(opts *bind.TransactOpts, token []common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "removeTokenInfoMany", token)
}

// RemoveTokenInfoMany is a paid mutator transaction binding the contract method 0x163f39ff.
//
// Solidity: function removeTokenInfoMany(address[] token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) RemoveTokenInfoMany(token []common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemoveTokenInfoMany(&_BridgeTokenInfo.TransactOpts, token)
}

// RemoveTokenInfoMany is a paid mutator transaction binding the contract method 0x163f39ff.
//
// Solidity: function removeTokenInfoMany(address[] token) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) RemoveTokenInfoMany(token []common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemoveTokenInfoMany(&_BridgeTokenInfo.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RenounceOwnership(&_BridgeTokenInfo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RenounceOwnership(&_BridgeTokenInfo.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.TransferOwnership(&_BridgeTokenInfo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.TransferOwnership(&_BridgeTokenInfo.TransactOpts, newOwner)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x5e78c722.
//
// Solidity: function updateTokenInfo(address token, uint256 minimum, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) UpdateTokenInfo(opts *bind.TransactOpts, token common.Address, minimum *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "updateTokenInfo", token, minimum, gasFee, serviceFee)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x5e78c722.
//
// Solidity: function updateTokenInfo(address token, uint256 minimum, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) UpdateTokenInfo(token common.Address, minimum *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimum, gasFee, serviceFee)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x5e78c722.
//
// Solidity: function updateTokenInfo(address token, uint256 minimum, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) UpdateTokenInfo(token common.Address, minimum *big.Int, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimum, gasFee, serviceFee)
}

// UpdateTokenInfoMany is a paid mutator transaction binding the contract method 0xbc45d21c.
//
// Solidity: function updateTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) UpdateTokenInfoMany(opts *bind.TransactOpts, tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "updateTokenInfoMany", tokenInfoList)
}

// UpdateTokenInfoMany is a paid mutator transaction binding the contract method 0xbc45d21c.
//
// Solidity: function updateTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) UpdateTokenInfoMany(tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateTokenInfoMany(&_BridgeTokenInfo.TransactOpts, tokenInfoList)
}

// UpdateTokenInfoMany is a paid mutator transaction binding the contract method 0xbc45d21c.
//
// Solidity: function updateTokenInfoMany((address,uint256,uint256,uint256)[] tokenInfoList) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) UpdateTokenInfoMany(tokenInfoList []IBridgeTokenInfoTokenInfo) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateTokenInfoMany(&_BridgeTokenInfo.TransactOpts, tokenInfoList)
}

// BridgeTokenInfoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoOwnershipTransferredIterator struct {
	Event *BridgeTokenInfoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoOwnershipTransferred)
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
		it.Event = new(BridgeTokenInfoOwnershipTransferred)
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
func (it *BridgeTokenInfoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeTokenInfoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoOwnershipTransferredIterator{contract: _BridgeTokenInfo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoOwnershipTransferred)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeTokenInfoOwnershipTransferred, error) {
	event := new(BridgeTokenInfoOwnershipTransferred)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenAddedIterator struct {
	Event *BridgeTokenInfoTokenAdded // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoTokenAdded)
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
		it.Event = new(BridgeTokenInfoTokenAdded)
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
func (it *BridgeTokenInfoTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoTokenAdded represents a TokenAdded event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenInfoTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoTokenAddedIterator{contract: _BridgeTokenInfo.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoTokenAdded)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseTokenAdded(log types.Log) (*BridgeTokenInfoTokenAdded, error) {
	event := new(BridgeTokenInfoTokenAdded)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoTokenInfoRemovedIterator is returned from FilterTokenInfoRemoved and is used to iterate over the raw logs and unpacked data for TokenInfoRemoved events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenInfoRemovedIterator struct {
	Event *BridgeTokenInfoTokenInfoRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoTokenInfoRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoTokenInfoRemoved)
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
		it.Event = new(BridgeTokenInfoTokenInfoRemoved)
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
func (it *BridgeTokenInfoTokenInfoRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoTokenInfoRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoTokenInfoRemoved represents a TokenInfoRemoved event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenInfoRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenInfoRemoved is a free log retrieval operation binding the contract event 0xce718809ef43b78b8d60e18e71da49dfe61d60cc8dc80c68a0574d393c13d845.
//
// Solidity: event TokenInfoRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterTokenInfoRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenInfoTokenInfoRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "TokenInfoRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoTokenInfoRemovedIterator{contract: _BridgeTokenInfo.contract, event: "TokenInfoRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenInfoRemoved is a free log subscription operation binding the contract event 0xce718809ef43b78b8d60e18e71da49dfe61d60cc8dc80c68a0574d393c13d845.
//
// Solidity: event TokenInfoRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchTokenInfoRemoved(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoTokenInfoRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "TokenInfoRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoTokenInfoRemoved)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenInfoRemoved", log); err != nil {
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

// ParseTokenInfoRemoved is a log parse operation binding the contract event 0xce718809ef43b78b8d60e18e71da49dfe61d60cc8dc80c68a0574d393c13d845.
//
// Solidity: event TokenInfoRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseTokenInfoRemoved(log types.Log) (*BridgeTokenInfoTokenInfoRemoved, error) {
	event := new(BridgeTokenInfoTokenInfoRemoved)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenInfoRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoTokenInfoUpdatedIterator is returned from FilterTokenInfoUpdated and is used to iterate over the raw logs and unpacked data for TokenInfoUpdated events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenInfoUpdatedIterator struct {
	Event *BridgeTokenInfoTokenInfoUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoTokenInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoTokenInfoUpdated)
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
		it.Event = new(BridgeTokenInfoTokenInfoUpdated)
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
func (it *BridgeTokenInfoTokenInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoTokenInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoTokenInfoUpdated represents a TokenInfoUpdated event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenInfoUpdated struct {
	Token        common.Address
	MinimumValue *big.Int
	GasFee       *big.Int
	ServiceFee   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenInfoUpdated is a free log retrieval operation binding the contract event 0xe7bb7f0ee932b644e52c34df47e4bf3b30372d6455e8c87d54517ac43a3a123a.
//
// Solidity: event TokenInfoUpdated(address indexed token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterTokenInfoUpdated(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenInfoTokenInfoUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "TokenInfoUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoTokenInfoUpdatedIterator{contract: _BridgeTokenInfo.contract, event: "TokenInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenInfoUpdated is a free log subscription operation binding the contract event 0xe7bb7f0ee932b644e52c34df47e4bf3b30372d6455e8c87d54517ac43a3a123a.
//
// Solidity: event TokenInfoUpdated(address indexed token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchTokenInfoUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoTokenInfoUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "TokenInfoUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoTokenInfoUpdated)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
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

// ParseTokenInfoUpdated is a log parse operation binding the contract event 0xe7bb7f0ee932b644e52c34df47e4bf3b30372d6455e8c87d54517ac43a3a123a.
//
// Solidity: event TokenInfoUpdated(address indexed token, uint256 minimumValue, uint256 gasFee, uint256 serviceFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseTokenInfoUpdated(log types.Log) (*BridgeTokenInfoTokenInfoUpdated, error) {
	event := new(BridgeTokenInfoTokenInfoUpdated)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenRemovedIterator struct {
	Event *BridgeTokenInfoTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoTokenRemoved)
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
		it.Event = new(BridgeTokenInfoTokenRemoved)
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
func (it *BridgeTokenInfoTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoTokenRemoved represents a TokenRemoved event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeTokenInfoTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoTokenRemovedIterator{contract: _BridgeTokenInfo.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoTokenRemoved)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseTokenRemoved(log types.Log) (*BridgeTokenInfoTokenRemoved, error) {
	event := new(BridgeTokenInfoTokenRemoved)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
