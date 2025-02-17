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

// BridgeFeeManagerMetaData contains all meta data concerning the BridgeFeeManager contract.
var BridgeFeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"addFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"feeInfoList\",\"type\":\"tuple[]\"}],\"name\":\"addFeeInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allFeeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"service\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"feeInfoByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenFeeList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"removeFeeInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"updateFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeFeeManager.FeeInfo[]\",\"name\":\"feeInfoList\",\"type\":\"tuple[]\"}],\"name\":\"updateFeeInfoMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"FeeInfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"FeeInfoRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"}],\"name\":\"FeeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BridgeFeeManagerInvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b88bcb7b": "addFeeInfo(address,uint256,uint256)",
		"5adbbe58": "addFeeInfoMany((address,uint256,uint256)[])",
		"9d0dc76e": "allFeeInfo()",
		"6ff97f1d": "allTokens()",
		"8b28ab1e": "calculateFee(address,uint256)",
		"5dbe47e8": "contains(address)",
		"96ce0795": "denominator()",
		"751b4c9c": "feeInfoByIndex(uint256)",
		"252154fa": "getTokenFee(address)",
		"cd3eda6a": "getTokenFeeList(address[])",
		"8da5cb5b": "owner()",
		"523fcfad": "removeFeeInfo(address)",
		"61dae8be": "removeFeeInfoMany(address[])",
		"715018a6": "renounceOwnership()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"d92fc67b": "tokensLength()",
		"f2fde38b": "transferOwnership(address)",
		"e5ec03ce": "updateFeeInfo(address,uint256,uint256)",
		"0092af81": "updateFeeInfoMany((address,uint256,uint256)[])",
	},
	Bin: "0x60806040526103e86035553480156014575f5ffd5b503380603957604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6040816045565b506094565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611172806100a15f395ff3fe608060405234801561000f575f5ffd5b506004361061011b575f3560e01c8063751b4c9c116100a9578063b88bcb7b1161006e578063b88bcb7b14610273578063cd3eda6a14610286578063d92fc67b14610299578063e5ec03ce146102a1578063f2fde38b146102b4575f5ffd5b8063751b4c9c146102015780638b28ab1e146102145780638da5cb5b1461023c57806396ce07951461024c5780639d0dc76e1461025e575f5ffd5b80635adbbe58116100ef5780635adbbe581461019b5780635dbe47e8146101ae57806361dae8be146101d15780636ff97f1d146101e4578063715018a6146101f9575f5ffd5b806292af811461011f578063252154fa146101345780634f6ccce71461015d578063523fcfad14610188575b5f5ffd5b61013261012d366004610e01565b6102c7565b005b610147610142366004610ecd565b61033f565b6040516101549190610ee6565b60405180910390f35b61017061016b366004610f10565b6103bf565b6040516001600160a01b039091168152602001610154565b610132610196366004610ecd565b6103d1565b6101326101a9366004610e01565b61044f565b6101c16101bc366004610ecd565b6104c3565b6040519015158152602001610154565b6101326101df366004610f27565b6104cf565b6101ec610504565b6040516101549190610fb5565b610132610568565b61014761020f366004610f10565b61057b565b610227610222366004611000565b6105d7565b60408051928352602083019190915201610154565b5f546001600160a01b0316610170565b6035545b604051908152602001610154565b610266610613565b6040516101549190611028565b610132610281366004611087565b6106fc565b610266610294366004610f27565b61079b565b610250610857565b6101326102af366004611087565b610867565b6101326102c2366004610ecd565b6108fc565b5f5b815181101561033b576103338282815181106102e7576102e76110b7565b60200260200101515f0151838381518110610304576103046110b7565b602002602001015160200151848481518110610322576103226110b7565b602002602001015160400151610867565b6001016102c9565b5050565b610347610d2e565b610350826104c3565b1561039857506001600160a01b039081165f90815260346020908152604091829020825160608101845281549094168452600181015491840191909152600201549082015290565b50604080516060810182526001600160a01b0390921682525f602083018190529082015290565b5f6103cb600183610939565b92915050565b6103da8161094b565b6001600160a01b0381165f8181526034602052604080822080546001600160a01b031916815560018101839055600201829055517f4a1f88682f04ee93bea00d2449d18d59d1183f9fc7d7b7d4ff94ed799ae398b691610444918190918252602082015260400190565b60405180910390a250565b5f5b815181101561033b576104bb82828151811061046f5761046f6110b7565b60200260200101515f015183838151811061048c5761048c6110b7565b6020026020010151602001518484815181106104aa576104aa6110b7565b6020026020010151604001516106fc565b600101610451565b5f6103cb6001836109f9565b5f5b815181101561033b576104fc8282815181106104ef576104ef6110b7565b60200260200101516103d1565b6001016104d1565b60605f6105116001610a1a565b90508067ffffffffffffffff81111561052c5761052c610d55565b604051908082528060200260200182016040528015610555578160200160208202803683370190505b5091506105626001610a23565b91505090565b610570610a2f565b6105795f610a5b565b565b610583610d2e565b60345f61058f846103bf565b6001600160a01b03908116825260208083019390935260409182015f208251606081018452815490921682526001810154938201939093526002909201549082015292915050565b5f5f5f6105e38561033f565b905080602001516035548260400151866105fd91906110df565b61060791906110f6565b92509250509250929050565b60605f61061e610857565b90505f8167ffffffffffffffff81111561063a5761063a610d55565b60405190808252806020026020018201604052801561067357816020015b610660610d2e565b8152602001906001900390816106585790505b5090505f5b828110156106f55760345f61068c836103bf565b6001600160a01b03908116825260208083019390935260409182015f208251606081018452815490921682526001810154938201939093526002909201549082015282518390839081106106e2576106e26110b7565b6020908102919091010152600101610678565b5092915050565b61070583610aaa565b604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260348452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fdbcb9310f07861333e121598ee5abc54194fac343c952ce1e175ad7fdbe87df291015b60405180910390a2505050565b60605f6107a6610857565b90505f8167ffffffffffffffff8111156107c2576107c2610d55565b6040519080825280602002602001820160405280156107fb57816020015b6107e8610d2e565b8152602001906001900390816107e05790505b5090505f5b8281101561084f5761082a85828151811061081d5761081d6110b7565b602002602001015161033f565b82828151811061083c5761083c6110b7565b6020908102919091010152600101610800565b509392505050565b5f6108626001610a1a565b905090565b610870836104c3565b83906108a057604051636fbda9d360e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b506001600160a01b0383165f818152603460209081526040918290206001810186905560020184905581518581529081018490527fd2d2f80ba194d563310f6dc5da5a24bfd705119757f191ad58d61dafd314ec28910161078e565b610904610a2f565b6001600160a01b03811661092d57604051631e4fbdf760e01b81525f6004820152602401610897565b61093681610a5b565b50565b5f6109448383610b58565b9392505050565b806001600160a01b03811661098b576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610897565b610996600183610b7e565b82906109c1576040516340da71e560e01b81526001600160a01b039091166004820152602401610897565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b6001600160a01b0381165f9081526001830160205260408120541515610944565b5f6103cb825490565b60605f61094483610b92565b5f546001600160a01b031633146105795760405163118cdaa760e01b8152336004820152602401610897565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b806001600160a01b038116610aea576040516330de3edf60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610897565b610af5600183610beb565b8290610b20576040516351eccfe160e11b81526001600160a01b039091166004820152602401610897565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b5f825f018281548110610b6d57610b6d6110b7565b905f5260205f200154905092915050565b5f610944836001600160a01b038416610bff565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610bdf57602002820191905f5260205f20905b815481526020019060010190808311610bcb575b50505050509050919050565b5f610944836001600160a01b038416610ce2565b5f8181526001830160205260408120548015610cd9575f610c21600183611115565b85549091505f90610c3490600190611115565b9050808214610c93575f865f018281548110610c5257610c526110b7565b905f5260205f200154905080875f018481548110610c7257610c726110b7565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610ca457610ca4611128565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506103cb565b5f9150506103cb565b5f818152600183016020526040812054610d2757508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556103cb565b505f6103cb565b60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff81118282101715610d8c57610d8c610d55565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610dbb57610dbb610d55565b604052919050565b5f67ffffffffffffffff821115610ddc57610ddc610d55565b5060051b60200190565b80356001600160a01b0381168114610dfc575f5ffd5b919050565b5f60208284031215610e11575f5ffd5b813567ffffffffffffffff811115610e27575f5ffd5b8201601f81018413610e37575f5ffd5b8035610e4a610e4582610dc3565b610d92565b80828252602082019150602060608402850101925086831115610e6b575f5ffd5b6020840193505b82841015610ec35760608488031215610e89575f5ffd5b610e91610d69565b610e9a85610de6565b815260208581013581830152604080870135908301529083526060909401939190910190610e72565b9695505050505050565b5f60208284031215610edd575f5ffd5b61094482610de6565b81516001600160a01b031681526020808301519082015260408083015190820152606081016103cb565b5f60208284031215610f20575f5ffd5b5035919050565b5f60208284031215610f37575f5ffd5b813567ffffffffffffffff811115610f4d575f5ffd5b8201601f81018413610f5d575f5ffd5b8035610f6b610e4582610dc3565b8082825260208201915060208360051b850101925086831115610f8c575f5ffd5b6020840193505b82841015610ec357610fa484610de6565b825260209384019390910190610f93565b602080825282518282018190525f918401906040840190835b81811015610ff55783516001600160a01b0316835260209384019390920191600101610fce565b509095945050505050565b5f5f60408385031215611011575f5ffd5b61101a83610de6565b946020939093013593505050565b602080825282518282018190525f918401906040840190835b81811015610ff55761107183855180516001600160a01b0316825260208082015190830152604090810151910152565b6020939093019260609290920191600101611041565b5f5f5f60608486031215611099575f5ffd5b6110a284610de6565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176103cb576103cb6110cb565b5f8261111057634e487b7160e01b5f52601260045260245ffd5b500490565b818103818111156103cb576103cb6110cb565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220d666061cde14e3029e499afaa8d0c8f568ff1052c7daaa8999739b3c85d7574b64736f6c634300081c0033",
}

// BridgeFeeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeFeeManagerMetaData.ABI instead.
var BridgeFeeManagerABI = BridgeFeeManagerMetaData.ABI

// Deprecated: Use BridgeFeeManagerMetaData.Sigs instead.
// BridgeFeeManagerFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFeeManagerFuncSigs = BridgeFeeManagerMetaData.Sigs

// BridgeFeeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeFeeManagerMetaData.Bin instead.
var BridgeFeeManagerBin = BridgeFeeManagerMetaData.Bin

// DeployBridgeFeeManager deploys a new Ethereum contract, binding an instance of BridgeFeeManager to it.
func DeployBridgeFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeFeeManager, error) {
	parsed, err := BridgeFeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeFeeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeFeeManager{BridgeFeeManagerCaller: BridgeFeeManagerCaller{contract: contract}, BridgeFeeManagerTransactor: BridgeFeeManagerTransactor{contract: contract}, BridgeFeeManagerFilterer: BridgeFeeManagerFilterer{contract: contract}}, nil
}

// BridgeFeeManager is an auto generated Go binding around an Ethereum contract.
type BridgeFeeManager struct {
	BridgeFeeManagerCaller     // Read-only binding to the contract
	BridgeFeeManagerTransactor // Write-only binding to the contract
	BridgeFeeManagerFilterer   // Log filterer for contract events
}

// BridgeFeeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeFeeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeFeeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFeeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeFeeManagerSession struct {
	Contract     *BridgeFeeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeFeeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeFeeManagerCallerSession struct {
	Contract *BridgeFeeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BridgeFeeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeFeeManagerTransactorSession struct {
	Contract     *BridgeFeeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BridgeFeeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeFeeManagerRaw struct {
	Contract *BridgeFeeManager // Generic contract binding to access the raw methods on
}

// BridgeFeeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeFeeManagerCallerRaw struct {
	Contract *BridgeFeeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeFeeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeFeeManagerTransactorRaw struct {
	Contract *BridgeFeeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeFeeManager creates a new instance of BridgeFeeManager, bound to a specific deployed contract.
func NewBridgeFeeManager(address common.Address, backend bind.ContractBackend) (*BridgeFeeManager, error) {
	contract, err := bindBridgeFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManager{BridgeFeeManagerCaller: BridgeFeeManagerCaller{contract: contract}, BridgeFeeManagerTransactor: BridgeFeeManagerTransactor{contract: contract}, BridgeFeeManagerFilterer: BridgeFeeManagerFilterer{contract: contract}}, nil
}

// NewBridgeFeeManagerCaller creates a new read-only instance of BridgeFeeManager, bound to a specific deployed contract.
func NewBridgeFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*BridgeFeeManagerCaller, error) {
	contract, err := bindBridgeFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerCaller{contract: contract}, nil
}

// NewBridgeFeeManagerTransactor creates a new write-only instance of BridgeFeeManager, bound to a specific deployed contract.
func NewBridgeFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeFeeManagerTransactor, error) {
	contract, err := bindBridgeFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerTransactor{contract: contract}, nil
}

// NewBridgeFeeManagerFilterer creates a new log filterer instance of BridgeFeeManager, bound to a specific deployed contract.
func NewBridgeFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFeeManagerFilterer, error) {
	contract, err := bindBridgeFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerFilterer{contract: contract}, nil
}

// bindBridgeFeeManager binds a generic wrapper to an already deployed contract.
func bindBridgeFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeFeeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFeeManager *BridgeFeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFeeManager.Contract.BridgeFeeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFeeManager *BridgeFeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.BridgeFeeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFeeManager *BridgeFeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.BridgeFeeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFeeManager *BridgeFeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFeeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFeeManager *BridgeFeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFeeManager *BridgeFeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.contract.Transact(opts, method, params...)
}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerCaller) AllFeeInfo(opts *bind.CallOpts) ([]IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "allFeeInfo")

	if err != nil {
		return *new([]IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeFeeManagerFeeInfo)).(*[]IBridgeFeeManagerFeeInfo)

	return out0, err

}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.AllFeeInfo(&_BridgeFeeManager.CallOpts)
}

// AllFeeInfo is a free data retrieval call binding the contract method 0x9d0dc76e.
//
// Solidity: function allFeeInfo() view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) AllFeeInfo() ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.AllFeeInfo(&_BridgeFeeManager.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeFeeManager *BridgeFeeManagerCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeFeeManager *BridgeFeeManagerSession) AllTokens() ([]common.Address, error) {
	return _BridgeFeeManager.Contract.AllTokens(&_BridgeFeeManager.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) AllTokens() ([]common.Address, error) {
	return _BridgeFeeManager.Contract.AllTokens(&_BridgeFeeManager.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeFeeManager *BridgeFeeManagerCaller) CalculateFee(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "calculateFee", token, value)

	outstruct := new(struct {
		Gas     *big.Int
		Service *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gas = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Service = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeFeeManager *BridgeFeeManagerSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeFeeManager.Contract.CalculateFee(&_BridgeFeeManager.CallOpts, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 value) view returns(uint256 gas, uint256 service)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) CalculateFee(token common.Address, value *big.Int) (struct {
	Gas     *big.Int
	Service *big.Int
}, error) {
	return _BridgeFeeManager.Contract.CalculateFee(&_BridgeFeeManager.CallOpts, token, value)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeFeeManager *BridgeFeeManagerCaller) Contains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "contains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeFeeManager *BridgeFeeManagerSession) Contains(token common.Address) (bool, error) {
	return _BridgeFeeManager.Contract.Contains(&_BridgeFeeManager.CallOpts, token)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) Contains(token common.Address) (bool, error) {
	return _BridgeFeeManager.Contract.Contains(&_BridgeFeeManager.CallOpts, token)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerSession) Denominator() (*big.Int, error) {
	return _BridgeFeeManager.Contract.Denominator(&_BridgeFeeManager.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) Denominator() (*big.Int, error) {
	return _BridgeFeeManager.Contract.Denominator(&_BridgeFeeManager.CallOpts)
}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerCaller) FeeInfoByIndex(opts *bind.CallOpts, index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "feeInfoByIndex", index)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.FeeInfoByIndex(&_BridgeFeeManager.CallOpts, index)
}

// FeeInfoByIndex is a free data retrieval call binding the contract method 0x751b4c9c.
//
// Solidity: function feeInfoByIndex(uint256 index) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) FeeInfoByIndex(index *big.Int) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.FeeInfoByIndex(&_BridgeFeeManager.CallOpts, index)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerCaller) GetTokenFee(opts *bind.CallOpts, token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "getTokenFee", token)

	if err != nil {
		return *new(IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeManagerFeeInfo)).(*IBridgeFeeManagerFeeInfo)

	return out0, err

}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.GetTokenFee(&_BridgeFeeManager.CallOpts, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x252154fa.
//
// Solidity: function getTokenFee(address token) view returns((address,uint256,uint256))
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) GetTokenFee(token common.Address) (IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.GetTokenFee(&_BridgeFeeManager.CallOpts, token)
}

// GetTokenFeeList is a free data retrieval call binding the contract method 0xcd3eda6a.
//
// Solidity: function getTokenFeeList(address[] tokens) view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerCaller) GetTokenFeeList(opts *bind.CallOpts, tokens []common.Address) ([]IBridgeFeeManagerFeeInfo, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "getTokenFeeList", tokens)

	if err != nil {
		return *new([]IBridgeFeeManagerFeeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeFeeManagerFeeInfo)).(*[]IBridgeFeeManagerFeeInfo)

	return out0, err

}

// GetTokenFeeList is a free data retrieval call binding the contract method 0xcd3eda6a.
//
// Solidity: function getTokenFeeList(address[] tokens) view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerSession) GetTokenFeeList(tokens []common.Address) ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.GetTokenFeeList(&_BridgeFeeManager.CallOpts, tokens)
}

// GetTokenFeeList is a free data retrieval call binding the contract method 0xcd3eda6a.
//
// Solidity: function getTokenFeeList(address[] tokens) view returns((address,uint256,uint256)[])
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) GetTokenFeeList(tokens []common.Address) ([]IBridgeFeeManagerFeeInfo, error) {
	return _BridgeFeeManager.Contract.GetTokenFeeList(&_BridgeFeeManager.CallOpts, tokens)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerSession) Owner() (common.Address, error) {
	return _BridgeFeeManager.Contract.Owner(&_BridgeFeeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) Owner() (common.Address, error) {
	return _BridgeFeeManager.Contract.Owner(&_BridgeFeeManager.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerCaller) TokenByIndex(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "tokenByIndex", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeFeeManager.Contract.TokenByIndex(&_BridgeFeeManager.CallOpts, i)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _BridgeFeeManager.Contract.TokenByIndex(&_BridgeFeeManager.CallOpts, i)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerSession) TokensLength() (*big.Int, error) {
	return _BridgeFeeManager.Contract.TokensLength(&_BridgeFeeManager.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) TokensLength() (*big.Int, error) {
	return _BridgeFeeManager.Contract.TokensLength(&_BridgeFeeManager.CallOpts)
}

// AddFeeInfo is a paid mutator transaction binding the contract method 0xb88bcb7b.
//
// Solidity: function addFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) AddFeeInfo(opts *bind.TransactOpts, token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "addFeeInfo", token, gasFee, serviceFee)
}

// AddFeeInfo is a paid mutator transaction binding the contract method 0xb88bcb7b.
//
// Solidity: function addFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) AddFeeInfo(token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.AddFeeInfo(&_BridgeFeeManager.TransactOpts, token, gasFee, serviceFee)
}

// AddFeeInfo is a paid mutator transaction binding the contract method 0xb88bcb7b.
//
// Solidity: function addFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) AddFeeInfo(token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.AddFeeInfo(&_BridgeFeeManager.TransactOpts, token, gasFee, serviceFee)
}

// AddFeeInfoMany is a paid mutator transaction binding the contract method 0x5adbbe58.
//
// Solidity: function addFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) AddFeeInfoMany(opts *bind.TransactOpts, feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "addFeeInfoMany", feeInfoList)
}

// AddFeeInfoMany is a paid mutator transaction binding the contract method 0x5adbbe58.
//
// Solidity: function addFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) AddFeeInfoMany(feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.AddFeeInfoMany(&_BridgeFeeManager.TransactOpts, feeInfoList)
}

// AddFeeInfoMany is a paid mutator transaction binding the contract method 0x5adbbe58.
//
// Solidity: function addFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) AddFeeInfoMany(feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.AddFeeInfoMany(&_BridgeFeeManager.TransactOpts, feeInfoList)
}

// RemoveFeeInfo is a paid mutator transaction binding the contract method 0x523fcfad.
//
// Solidity: function removeFeeInfo(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) RemoveFeeInfo(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "removeFeeInfo", token)
}

// RemoveFeeInfo is a paid mutator transaction binding the contract method 0x523fcfad.
//
// Solidity: function removeFeeInfo(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) RemoveFeeInfo(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveFeeInfo(&_BridgeFeeManager.TransactOpts, token)
}

// RemoveFeeInfo is a paid mutator transaction binding the contract method 0x523fcfad.
//
// Solidity: function removeFeeInfo(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) RemoveFeeInfo(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveFeeInfo(&_BridgeFeeManager.TransactOpts, token)
}

// RemoveFeeInfoMany is a paid mutator transaction binding the contract method 0x61dae8be.
//
// Solidity: function removeFeeInfoMany(address[] token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) RemoveFeeInfoMany(opts *bind.TransactOpts, token []common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "removeFeeInfoMany", token)
}

// RemoveFeeInfoMany is a paid mutator transaction binding the contract method 0x61dae8be.
//
// Solidity: function removeFeeInfoMany(address[] token) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) RemoveFeeInfoMany(token []common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveFeeInfoMany(&_BridgeFeeManager.TransactOpts, token)
}

// RemoveFeeInfoMany is a paid mutator transaction binding the contract method 0x61dae8be.
//
// Solidity: function removeFeeInfoMany(address[] token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) RemoveFeeInfoMany(token []common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveFeeInfoMany(&_BridgeFeeManager.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RenounceOwnership(&_BridgeFeeManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RenounceOwnership(&_BridgeFeeManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.TransferOwnership(&_BridgeFeeManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.TransferOwnership(&_BridgeFeeManager.TransactOpts, newOwner)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0xe5ec03ce.
//
// Solidity: function updateFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) UpdateFeeInfo(opts *bind.TransactOpts, token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "updateFeeInfo", token, gasFee, serviceFee)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0xe5ec03ce.
//
// Solidity: function updateFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) UpdateFeeInfo(token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateFeeInfo(&_BridgeFeeManager.TransactOpts, token, gasFee, serviceFee)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0xe5ec03ce.
//
// Solidity: function updateFeeInfo(address token, uint256 gasFee, uint256 serviceFee) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) UpdateFeeInfo(token common.Address, gasFee *big.Int, serviceFee *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateFeeInfo(&_BridgeFeeManager.TransactOpts, token, gasFee, serviceFee)
}

// UpdateFeeInfoMany is a paid mutator transaction binding the contract method 0x0092af81.
//
// Solidity: function updateFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) UpdateFeeInfoMany(opts *bind.TransactOpts, feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "updateFeeInfoMany", feeInfoList)
}

// UpdateFeeInfoMany is a paid mutator transaction binding the contract method 0x0092af81.
//
// Solidity: function updateFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) UpdateFeeInfoMany(feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateFeeInfoMany(&_BridgeFeeManager.TransactOpts, feeInfoList)
}

// UpdateFeeInfoMany is a paid mutator transaction binding the contract method 0x0092af81.
//
// Solidity: function updateFeeInfoMany((address,uint256,uint256)[] feeInfoList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) UpdateFeeInfoMany(feeInfoList []IBridgeFeeManagerFeeInfo) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateFeeInfoMany(&_BridgeFeeManager.TransactOpts, feeInfoList)
}

// BridgeFeeManagerFeeInfoAddedIterator is returned from FilterFeeInfoAdded and is used to iterate over the raw logs and unpacked data for FeeInfoAdded events raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoAddedIterator struct {
	Event *BridgeFeeManagerFeeInfoAdded // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerFeeInfoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerFeeInfoAdded)
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
		it.Event = new(BridgeFeeManagerFeeInfoAdded)
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
func (it *BridgeFeeManagerFeeInfoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerFeeInfoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerFeeInfoAdded represents a FeeInfoAdded event raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoAdded struct {
	Token      common.Address
	GasFee     *big.Int
	ServiceFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeInfoAdded is a free log retrieval operation binding the contract event 0xdbcb9310f07861333e121598ee5abc54194fac343c952ce1e175ad7fdbe87df2.
//
// Solidity: event FeeInfoAdded(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterFeeInfoAdded(opts *bind.FilterOpts, token []common.Address) (*BridgeFeeManagerFeeInfoAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "FeeInfoAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerFeeInfoAddedIterator{contract: _BridgeFeeManager.contract, event: "FeeInfoAdded", logs: logs, sub: sub}, nil
}

// WatchFeeInfoAdded is a free log subscription operation binding the contract event 0xdbcb9310f07861333e121598ee5abc54194fac343c952ce1e175ad7fdbe87df2.
//
// Solidity: event FeeInfoAdded(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchFeeInfoAdded(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerFeeInfoAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "FeeInfoAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerFeeInfoAdded)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoAdded", log); err != nil {
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

// ParseFeeInfoAdded is a log parse operation binding the contract event 0xdbcb9310f07861333e121598ee5abc54194fac343c952ce1e175ad7fdbe87df2.
//
// Solidity: event FeeInfoAdded(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseFeeInfoAdded(log types.Log) (*BridgeFeeManagerFeeInfoAdded, error) {
	event := new(BridgeFeeManagerFeeInfoAdded)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerFeeInfoRemovedIterator is returned from FilterFeeInfoRemoved and is used to iterate over the raw logs and unpacked data for FeeInfoRemoved events raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoRemovedIterator struct {
	Event *BridgeFeeManagerFeeInfoRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerFeeInfoRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerFeeInfoRemoved)
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
		it.Event = new(BridgeFeeManagerFeeInfoRemoved)
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
func (it *BridgeFeeManagerFeeInfoRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerFeeInfoRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerFeeInfoRemoved represents a FeeInfoRemoved event raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoRemoved struct {
	Token      common.Address
	GasFee     *big.Int
	ServiceFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeInfoRemoved is a free log retrieval operation binding the contract event 0x4a1f88682f04ee93bea00d2449d18d59d1183f9fc7d7b7d4ff94ed799ae398b6.
//
// Solidity: event FeeInfoRemoved(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterFeeInfoRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeFeeManagerFeeInfoRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "FeeInfoRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerFeeInfoRemovedIterator{contract: _BridgeFeeManager.contract, event: "FeeInfoRemoved", logs: logs, sub: sub}, nil
}

// WatchFeeInfoRemoved is a free log subscription operation binding the contract event 0x4a1f88682f04ee93bea00d2449d18d59d1183f9fc7d7b7d4ff94ed799ae398b6.
//
// Solidity: event FeeInfoRemoved(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchFeeInfoRemoved(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerFeeInfoRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "FeeInfoRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerFeeInfoRemoved)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoRemoved", log); err != nil {
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

// ParseFeeInfoRemoved is a log parse operation binding the contract event 0x4a1f88682f04ee93bea00d2449d18d59d1183f9fc7d7b7d4ff94ed799ae398b6.
//
// Solidity: event FeeInfoRemoved(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseFeeInfoRemoved(log types.Log) (*BridgeFeeManagerFeeInfoRemoved, error) {
	event := new(BridgeFeeManagerFeeInfoRemoved)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerFeeInfoUpdatedIterator is returned from FilterFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for FeeInfoUpdated events raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoUpdatedIterator struct {
	Event *BridgeFeeManagerFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerFeeInfoUpdated)
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
		it.Event = new(BridgeFeeManagerFeeInfoUpdated)
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
func (it *BridgeFeeManagerFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerFeeInfoUpdated represents a FeeInfoUpdated event raised by the BridgeFeeManager contract.
type BridgeFeeManagerFeeInfoUpdated struct {
	Token      common.Address
	GasFee     *big.Int
	ServiceFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeInfoUpdated is a free log retrieval operation binding the contract event 0xd2d2f80ba194d563310f6dc5da5a24bfd705119757f191ad58d61dafd314ec28.
//
// Solidity: event FeeInfoUpdated(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterFeeInfoUpdated(opts *bind.FilterOpts, token []common.Address) (*BridgeFeeManagerFeeInfoUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "FeeInfoUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerFeeInfoUpdatedIterator{contract: _BridgeFeeManager.contract, event: "FeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeInfoUpdated is a free log subscription operation binding the contract event 0xd2d2f80ba194d563310f6dc5da5a24bfd705119757f191ad58d61dafd314ec28.
//
// Solidity: event FeeInfoUpdated(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerFeeInfoUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "FeeInfoUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerFeeInfoUpdated)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoUpdated", log); err != nil {
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

// ParseFeeInfoUpdated is a log parse operation binding the contract event 0xd2d2f80ba194d563310f6dc5da5a24bfd705119757f191ad58d61dafd314ec28.
//
// Solidity: event FeeInfoUpdated(address indexed token, uint256 gasFee, uint256 serviceFee)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseFeeInfoUpdated(log types.Log) (*BridgeFeeManagerFeeInfoUpdated, error) {
	event := new(BridgeFeeManagerFeeInfoUpdated)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "FeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeFeeManager contract.
type BridgeFeeManagerOwnershipTransferredIterator struct {
	Event *BridgeFeeManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerOwnershipTransferred)
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
		it.Event = new(BridgeFeeManagerOwnershipTransferred)
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
func (it *BridgeFeeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeFeeManager contract.
type BridgeFeeManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeFeeManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerOwnershipTransferredIterator{contract: _BridgeFeeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerOwnershipTransferred)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeFeeManagerOwnershipTransferred, error) {
	event := new(BridgeFeeManagerOwnershipTransferred)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the BridgeFeeManager contract.
type BridgeFeeManagerTokenAddedIterator struct {
	Event *BridgeFeeManagerTokenAdded // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerTokenAdded)
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
		it.Event = new(BridgeFeeManagerTokenAdded)
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
func (it *BridgeFeeManagerTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerTokenAdded represents a TokenAdded event raised by the BridgeFeeManager contract.
type BridgeFeeManagerTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*BridgeFeeManagerTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerTokenAddedIterator{contract: _BridgeFeeManager.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerTokenAdded)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseTokenAdded(log types.Log) (*BridgeFeeManagerTokenAdded, error) {
	event := new(BridgeFeeManagerTokenAdded)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the BridgeFeeManager contract.
type BridgeFeeManagerTokenRemovedIterator struct {
	Event *BridgeFeeManagerTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerTokenRemoved)
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
		it.Event = new(BridgeFeeManagerTokenRemoved)
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
func (it *BridgeFeeManagerTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerTokenRemoved represents a TokenRemoved event raised by the BridgeFeeManager contract.
type BridgeFeeManagerTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*BridgeFeeManagerTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerTokenRemovedIterator{contract: _BridgeFeeManager.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerTokenRemoved)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseTokenRemoved(log types.Log) (*BridgeFeeManagerTokenRemoved, error) {
	event := new(BridgeFeeManagerTokenRemoved)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
