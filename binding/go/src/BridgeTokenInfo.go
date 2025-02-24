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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wcoin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"}],\"name\":\"addTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"}],\"name\":\"calculateMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"}],\"name\":\"changeEx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeGas\",\"type\":\"uint256\"}],\"name\":\"updateFinalizeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"name\":\"BridgeTokenInfoExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"BridgeTokenInfoGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"BridgeTokenInfoPriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeTokenInfoCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeTokenInfoCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eth\",\"type\":\"address\"}],\"name\":\"BridgeTokenInfoNotFoundETHPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"22f136ef": "_tokenInfo(address)",
		"453a9c9b": "addTokenInfo(address,uint256,uint256)",
		"6e908ca3": "calculate(address,uint256)",
		"8ad46104": "calculateMax(address,uint256)",
		"40d48d36": "changeEx(uint256)",
		"96ce0795": "denominator()",
		"1f69565f": "getTokenInfo(address)",
		"8da5cb5b": "owner()",
		"741bef1a": "priceFeed()",
		"ce2e5c66": "removePriceFeed()",
		"95e4c77f": "removeTokenInfo(address)",
		"715018a6": "renounceOwnership()",
		"724e78da": "setPriceFeed(address)",
		"f2fde38b": "transferOwnership(address)",
		"46899c00": "updateFinalizeGas(uint256)",
		"204cff81": "updateGasPrice(uint256)",
		"600422b3": "wcoin()",
	},
	Bin: "0x6080604052600a60035564012a05f200600455620186a0600555348015610024575f5ffd5b50604051610d7e380380610d7e83398101604081905261004391610127565b338061006957604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610072816100d8565b506001600160a01b0381166100b35760405163f4bf80a960e01b81526020600482015260066024820152652fbbb1b7b4b760d11b6044820152606401610060565b600280546001600160a01b0319166001600160a01b0392909216919091179055610154565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215610137575f5ffd5b81516001600160a01b038116811461014d575f5ffd5b9392505050565b610c1d806101615f395ff3fe608060405234801561000f575f5ffd5b5060043610610106575f3560e01c8063715018a61161009e5780638da5cb5b1161006e5780638da5cb5b1461026e57806395e4c77f1461027e57806396ce079514610291578063ce2e5c66146102a1578063f2fde38b146102a9575f5ffd5b8063715018a61461022d578063724e78da14610235578063741bef1a146102485780638ad461041461025b575f5ffd5b8063453a9c9b116100d9578063453a9c9b146101c957806346899c00146101dc578063600422b3146101ef5780636e908ca31461021a575f5ffd5b80631f69565f1461010a578063204cff811461014457806322f136ef1461015957806340d48d36146101b6575b5f5ffd5b61011d610118366004610a97565b6102bc565b60408051948552602085019390935291830152151560608201526080015b60405180910390f35b610157610152366004610ab2565b610347565b005b610191610167366004610a97565b60066020525f90815260409020805460018201546002909201546001600160a01b03909116919083565b604080516001600160a01b03909416845260208401929092529082015260600161013b565b6101576101c4366004610ab2565b6103cb565b6101576101d7366004610ac9565b610408565b6101576101ea366004610ab2565b6104a7565b600254610202906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61011d610228366004610afb565b610522565b610157610560565b610157610243366004610a97565b610573565b600154610202906001600160a01b031681565b61011d610269366004610afb565b6106bb565b5f546001600160a01b0316610202565b61015761028c366004610a97565b610738565b6040516103e8815260200161013b565b6101576107d4565b6101576102b7366004610a97565b610821565b6001600160a01b038082165f908152600660209081526040808320815160608101835281549095168552600181015492850192909252600290910154908301529081908190819061030c8661085b565b825191955092506001600160a01b0387811691161461032e575f600354610339565b806020015181604001515b909794965094509092915050565b61034f610940565b805f0361038f5760405163beadd5a760e01b8152602060048201526008602482015267676173507269636560c01b60448201526064015b60405180910390fd5b60048190556040518181527f43635a24cdd8735af9fc48ff755c3da38b453a3ae4d6e1894315316d948e4076906020015b60405180910390a150565b6103d3610940565b60038190556040518181527fc71b628d8981ab4bd652a87d1e08186c47c5dd5d92f8b6d7fa599a0e2d02baf9906020016103c0565b610410610940565b6001600160a01b03831661044f5760405163f4bf80a960e01b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610386565b604080516060810182526001600160a01b0394851680825260208083019586528284019485525f9182526006905291909120905181546001600160a01b03191694169390931783559051600183015551600290910155565b6104af610940565b805f036104ed5760405163beadd5a760e01b815260206004820152600b60248201526a66696e616c697a6547617360a81b6044820152606401610386565b60058190556040518181527f43635a24cdd8735af9fc48ff755c3da38b453a3ae4d6e1894315316d948e4076906020016103c0565b5f5f5f5f61052f866102bc565b9296509094509250905061054685836103e861096c565b91508080156105555750838510155b905092959194509250565b610568610940565b6105715f610a23565b565b61057b610940565b6001600160a01b0381166105bf5760405163f4bf80a960e01b815260206004820152600a60248201526917dc1c9a58d95199595960b21b6044820152606401610386565b600254604051630bb7c8fd60e31b81526001600160a01b0391821660048201525f91831690635dbe47e890602401602060405180830381865afa158015610608573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061062c9190610b39565b6002549091506001600160a01b03168161066557604051632bfcb04f60e11b81526001600160a01b039091166004820152602401610386565b50600180546001600160a01b0319166001600160a01b0384169081179091556040519081527f9de4f0e361ff2892765ceacf345aee96b96513207cd9ebee35bd9ee67849b0089060200160405180910390a15050565b5f5f5f5f5f6106c9876102bc565b9196509450925090508115806106df5750838611155b156106f6575f5f5f5f94509450945094505061072f565b5f6107018588610b66565b905061071b6103e86107138682610b7f565b83919061096c565b955061072a86856103e861096c565b935050505b92959194509250565b610740610940565b6001600160a01b03811661077f5760405163f4bf80a960e01b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610386565b6001600160a01b038082165f81815260066020526040902054909116036107d1576001600160a01b0381165f90815260066020526040812080546001600160a01b031916815560018101829055600201555b50565b6107dc610940565b600180546001600160a01b03191690556040515f81527f9de4f0e361ff2892765ceacf345aee96b96513207cd9ebee35bd9ee67849b0089060200160405180910390a1565b610829610940565b6001600160a01b03811661085257604051631e4fbdf760e01b81525f6004820152602401610386565b6107d181610a23565b6001545f9081906001600160a01b0316156109385760015460025460055460045473__$79c24c681325f76413d25d3c06c8219b6e$__9363126e5b3a936001600160a01b039182169391169188916108b291610b92565b6040516001600160e01b031960e087901b1681526001600160a01b03948516600482015292841660248401529216604482015260648101919091526084016040805180830381865af415801561090a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092e9190610ba9565b9092509050915091565b506001915091565b5f546001600160a01b031633146105715760405163118cdaa760e01b8152336004820152602401610386565b5f838302815f1985870982811083820303915050805f036109a05783828161099657610996610bd3565b0492505050610a1c565b8084116109b7576109b76003851502601118610a72565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b715f52806020526024601cfd5b6001600160a01b03811681146107d1575f5ffd5b5f60208284031215610aa7575f5ffd5b8135610a1c81610a83565b5f60208284031215610ac2575f5ffd5b5035919050565b5f5f5f60608486031215610adb575f5ffd5b8335610ae681610a83565b95602085013595506040909401359392505050565b5f5f60408385031215610b0c575f5ffd5b8235610b1781610a83565b946020939093013593505050565b80518015158114610b34575f5ffd5b919050565b5f60208284031215610b49575f5ffd5b610a1c82610b25565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610b7957610b79610b52565b92915050565b80820180821115610b7957610b79610b52565b8082028115828204841417610b7957610b79610b52565b5f5f60408385031215610bba575f5ffd5b82519150610bca60208401610b25565b90509250929050565b634e487b7160e01b5f52601260045260245ffdfea2646970667358221220f780f9e58f3d0bb2f14171aba30b9a95bf509eb4ff8acb9d0ab9be1aabb4ece364736f6c634300081c0033",
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
func DeployBridgeTokenInfo(auth *bind.TransactOpts, backend bind.ContractBackend, _wcoin common.Address) (common.Address, *types.Transaction, *BridgeTokenInfo, error) {
	parsed, err := BridgeTokenInfoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTokenInfoBin), backend, _wcoin)
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

// TokenInfo is a free data retrieval call binding the contract method 0x22f136ef.
//
// Solidity: function _tokenInfo(address ) view returns(address token, uint256 minimumAmount, uint256 exchangeFee)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token         common.Address
	MinimumAmount *big.Int
	ExchangeFee   *big.Int
}, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "_tokenInfo", arg0)

	outstruct := new(struct {
		Token         common.Address
		MinimumAmount *big.Int
		ExchangeFee   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinimumAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExchangeFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0x22f136ef.
//
// Solidity: function _tokenInfo(address ) view returns(address token, uint256 minimumAmount, uint256 exchangeFee)
func (_BridgeTokenInfo *BridgeTokenInfoSession) TokenInfo(arg0 common.Address) (struct {
	Token         common.Address
	MinimumAmount *big.Int
	ExchangeFee   *big.Int
}, error) {
	return _BridgeTokenInfo.Contract.TokenInfo(&_BridgeTokenInfo.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0x22f136ef.
//
// Solidity: function _tokenInfo(address ) view returns(address token, uint256 minimumAmount, uint256 exchangeFee)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) TokenInfo(arg0 common.Address) (struct {
	Token         common.Address
	MinimumAmount *big.Int
	ExchangeFee   *big.Int
}, error) {
	return _BridgeTokenInfo.Contract.TokenInfo(&_BridgeTokenInfo.CallOpts, arg0)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Calculate(opts *bind.CallOpts, token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "calculate", token, value)

	outstruct := new(struct {
		Minimum *big.Int
		Gas     *big.Int
		Ex      *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minimum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.Calculate(&_BridgeTokenInfo.CallOpts, token, value)
}

// Calculate is a free data retrieval call binding the contract method 0x6e908ca3.
//
// Solidity: function calculate(address token, uint256 value) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Calculate(token common.Address, value *big.Int) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.Calculate(&_BridgeTokenInfo.CallOpts, token, value)
}

// CalculateMax is a free data retrieval call binding the contract method 0x8ad46104.
//
// Solidity: function calculateMax(address token, uint256 totalValue) view returns(uint256 value, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) CalculateMax(opts *bind.CallOpts, token common.Address, totalValue *big.Int) (struct {
	Value   *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "calculateMax", token, totalValue)

	outstruct := new(struct {
		Value   *big.Int
		Gas     *big.Int
		Ex      *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// CalculateMax is a free data retrieval call binding the contract method 0x8ad46104.
//
// Solidity: function calculateMax(address token, uint256 totalValue) view returns(uint256 value, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoSession) CalculateMax(token common.Address, totalValue *big.Int) (struct {
	Value   *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.CalculateMax(&_BridgeTokenInfo.CallOpts, token, totalValue)
}

// CalculateMax is a free data retrieval call binding the contract method 0x8ad46104.
//
// Solidity: function calculateMax(address token, uint256 totalValue) view returns(uint256 value, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) CalculateMax(token common.Address, totalValue *big.Int) (struct {
	Value   *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.CalculateMax(&_BridgeTokenInfo.CallOpts, token, totalValue)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
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
// Solidity: function denominator() pure returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Denominator() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.Denominator(&_BridgeTokenInfo.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Denominator() (*big.Int, error) {
	return _BridgeTokenInfo.Contract.Denominator(&_BridgeTokenInfo.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) GetTokenInfo(opts *bind.CallOpts, token common.Address) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "getTokenInfo", token)

	outstruct := new(struct {
		Minimum *big.Int
		Gas     *big.Int
		Ex      *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minimum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoSession) GetTokenInfo(token common.Address) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfo(&_BridgeTokenInfo.CallOpts, token)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 minimum, uint256 gas, uint256 ex, bool isValid)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) GetTokenInfo(token common.Address) (struct {
	Minimum *big.Int
	Gas     *big.Int
	Ex      *big.Int
	IsValid bool
}, error) {
	return _BridgeTokenInfo.Contract.GetTokenInfo(&_BridgeTokenInfo.CallOpts, token)
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

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoSession) PriceFeed() (common.Address, error) {
	return _BridgeTokenInfo.Contract.PriceFeed(&_BridgeTokenInfo.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) PriceFeed() (common.Address, error) {
	return _BridgeTokenInfo.Contract.PriceFeed(&_BridgeTokenInfo.CallOpts)
}

// Wcoin is a free data retrieval call binding the contract method 0x600422b3.
//
// Solidity: function wcoin() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCaller) Wcoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTokenInfo.contract.Call(opts, &out, "wcoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcoin is a free data retrieval call binding the contract method 0x600422b3.
//
// Solidity: function wcoin() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoSession) Wcoin() (common.Address, error) {
	return _BridgeTokenInfo.Contract.Wcoin(&_BridgeTokenInfo.CallOpts)
}

// Wcoin is a free data retrieval call binding the contract method 0x600422b3.
//
// Solidity: function wcoin() view returns(address)
func (_BridgeTokenInfo *BridgeTokenInfoCallerSession) Wcoin() (common.Address, error) {
	return _BridgeTokenInfo.Contract.Wcoin(&_BridgeTokenInfo.CallOpts)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0x453a9c9b.
//
// Solidity: function addTokenInfo(address token, uint256 minimum, uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) AddTokenInfo(opts *bind.TransactOpts, token common.Address, minimum *big.Int, ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "addTokenInfo", token, minimum, ex)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0x453a9c9b.
//
// Solidity: function addTokenInfo(address token, uint256 minimum, uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) AddTokenInfo(token common.Address, minimum *big.Int, ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimum, ex)
}

// AddTokenInfo is a paid mutator transaction binding the contract method 0x453a9c9b.
//
// Solidity: function addTokenInfo(address token, uint256 minimum, uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) AddTokenInfo(token common.Address, minimum *big.Int, ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.AddTokenInfo(&_BridgeTokenInfo.TransactOpts, token, minimum, ex)
}

// ChangeEx is a paid mutator transaction binding the contract method 0x40d48d36.
//
// Solidity: function changeEx(uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) ChangeEx(opts *bind.TransactOpts, ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "changeEx", ex)
}

// ChangeEx is a paid mutator transaction binding the contract method 0x40d48d36.
//
// Solidity: function changeEx(uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) ChangeEx(ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.ChangeEx(&_BridgeTokenInfo.TransactOpts, ex)
}

// ChangeEx is a paid mutator transaction binding the contract method 0x40d48d36.
//
// Solidity: function changeEx(uint256 ex) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) ChangeEx(ex *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.ChangeEx(&_BridgeTokenInfo.TransactOpts, ex)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) RemovePriceFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "removePriceFeed")
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemovePriceFeed(&_BridgeTokenInfo.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.RemovePriceFeed(&_BridgeTokenInfo.TransactOpts)
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

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.SetPriceFeed(&_BridgeTokenInfo.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.SetPriceFeed(&_BridgeTokenInfo.TransactOpts, _priceFeed)
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

// UpdateFinalizeGas is a paid mutator transaction binding the contract method 0x46899c00.
//
// Solidity: function updateFinalizeGas(uint256 finalizeGas) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) UpdateFinalizeGas(opts *bind.TransactOpts, finalizeGas *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "updateFinalizeGas", finalizeGas)
}

// UpdateFinalizeGas is a paid mutator transaction binding the contract method 0x46899c00.
//
// Solidity: function updateFinalizeGas(uint256 finalizeGas) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) UpdateFinalizeGas(finalizeGas *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateFinalizeGas(&_BridgeTokenInfo.TransactOpts, finalizeGas)
}

// UpdateFinalizeGas is a paid mutator transaction binding the contract method 0x46899c00.
//
// Solidity: function updateFinalizeGas(uint256 finalizeGas) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) UpdateFinalizeGas(finalizeGas *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateFinalizeGas(&_BridgeTokenInfo.TransactOpts, finalizeGas)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x204cff81.
//
// Solidity: function updateGasPrice(uint256 gasPrice) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactor) UpdateGasPrice(opts *bind.TransactOpts, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.contract.Transact(opts, "updateGasPrice", gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x204cff81.
//
// Solidity: function updateGasPrice(uint256 gasPrice) returns()
func (_BridgeTokenInfo *BridgeTokenInfoSession) UpdateGasPrice(gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateGasPrice(&_BridgeTokenInfo.TransactOpts, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x204cff81.
//
// Solidity: function updateGasPrice(uint256 gasPrice) returns()
func (_BridgeTokenInfo *BridgeTokenInfoTransactorSession) UpdateGasPrice(gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeTokenInfo.Contract.UpdateGasPrice(&_BridgeTokenInfo.TransactOpts, gasPrice)
}

// BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator is returned from FilterBridgeTokenInfoExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeTokenInfoExchangeFeeUpdated events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator struct {
	Event *BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated)
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
		it.Event = new(BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated)
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
func (it *BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated represents a BridgeTokenInfoExchangeFeeUpdated event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated struct {
	ExFee *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokenInfoExchangeFeeUpdated is a free log retrieval operation binding the contract event 0xc71b628d8981ab4bd652a87d1e08186c47c5dd5d92f8b6d7fa599a0e2d02baf9.
//
// Solidity: event BridgeTokenInfoExchangeFeeUpdated(uint256 exFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterBridgeTokenInfoExchangeFeeUpdated(opts *bind.FilterOpts) (*BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator, error) {

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "BridgeTokenInfoExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdatedIterator{contract: _BridgeTokenInfo.contract, event: "BridgeTokenInfoExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTokenInfoExchangeFeeUpdated is a free log subscription operation binding the contract event 0xc71b628d8981ab4bd652a87d1e08186c47c5dd5d92f8b6d7fa599a0e2d02baf9.
//
// Solidity: event BridgeTokenInfoExchangeFeeUpdated(uint256 exFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchBridgeTokenInfoExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "BridgeTokenInfoExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoExchangeFeeUpdated", log); err != nil {
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

// ParseBridgeTokenInfoExchangeFeeUpdated is a log parse operation binding the contract event 0xc71b628d8981ab4bd652a87d1e08186c47c5dd5d92f8b6d7fa599a0e2d02baf9.
//
// Solidity: event BridgeTokenInfoExchangeFeeUpdated(uint256 exFee)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseBridgeTokenInfoExchangeFeeUpdated(log types.Log) (*BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated, error) {
	event := new(BridgeTokenInfoBridgeTokenInfoExchangeFeeUpdated)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator is returned from FilterBridgeTokenInfoGasPriceUpdated and is used to iterate over the raw logs and unpacked data for BridgeTokenInfoGasPriceUpdated events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator struct {
	Event *BridgeTokenInfoBridgeTokenInfoGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoBridgeTokenInfoGasPriceUpdated)
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
		it.Event = new(BridgeTokenInfoBridgeTokenInfoGasPriceUpdated)
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
func (it *BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoBridgeTokenInfoGasPriceUpdated represents a BridgeTokenInfoGasPriceUpdated event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoGasPriceUpdated struct {
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokenInfoGasPriceUpdated is a free log retrieval operation binding the contract event 0x43635a24cdd8735af9fc48ff755c3da38b453a3ae4d6e1894315316d948e4076.
//
// Solidity: event BridgeTokenInfoGasPriceUpdated(uint256 gasPrice)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterBridgeTokenInfoGasPriceUpdated(opts *bind.FilterOpts) (*BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator, error) {

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "BridgeTokenInfoGasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoBridgeTokenInfoGasPriceUpdatedIterator{contract: _BridgeTokenInfo.contract, event: "BridgeTokenInfoGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTokenInfoGasPriceUpdated is a free log subscription operation binding the contract event 0x43635a24cdd8735af9fc48ff755c3da38b453a3ae4d6e1894315316d948e4076.
//
// Solidity: event BridgeTokenInfoGasPriceUpdated(uint256 gasPrice)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchBridgeTokenInfoGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoBridgeTokenInfoGasPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "BridgeTokenInfoGasPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoBridgeTokenInfoGasPriceUpdated)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoGasPriceUpdated", log); err != nil {
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

// ParseBridgeTokenInfoGasPriceUpdated is a log parse operation binding the contract event 0x43635a24cdd8735af9fc48ff755c3da38b453a3ae4d6e1894315316d948e4076.
//
// Solidity: event BridgeTokenInfoGasPriceUpdated(uint256 gasPrice)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseBridgeTokenInfoGasPriceUpdated(log types.Log) (*BridgeTokenInfoBridgeTokenInfoGasPriceUpdated, error) {
	event := new(BridgeTokenInfoBridgeTokenInfoGasPriceUpdated)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator is returned from FilterBridgeTokenInfoPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for BridgeTokenInfoPriceFeedUpdated events raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator struct {
	Event *BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated)
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
		it.Event = new(BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated)
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
func (it *BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated represents a BridgeTokenInfoPriceFeedUpdated event raised by the BridgeTokenInfo contract.
type BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokenInfoPriceFeedUpdated is a free log retrieval operation binding the contract event 0x9de4f0e361ff2892765ceacf345aee96b96513207cd9ebee35bd9ee67849b008.
//
// Solidity: event BridgeTokenInfoPriceFeedUpdated(address priceFeed)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) FilterBridgeTokenInfoPriceFeedUpdated(opts *bind.FilterOpts) (*BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator, error) {

	logs, sub, err := _BridgeTokenInfo.contract.FilterLogs(opts, "BridgeTokenInfoPriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenInfoBridgeTokenInfoPriceFeedUpdatedIterator{contract: _BridgeTokenInfo.contract, event: "BridgeTokenInfoPriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTokenInfoPriceFeedUpdated is a free log subscription operation binding the contract event 0x9de4f0e361ff2892765ceacf345aee96b96513207cd9ebee35bd9ee67849b008.
//
// Solidity: event BridgeTokenInfoPriceFeedUpdated(address priceFeed)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) WatchBridgeTokenInfoPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTokenInfo.contract.WatchLogs(opts, "BridgeTokenInfoPriceFeedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated)
				if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoPriceFeedUpdated", log); err != nil {
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

// ParseBridgeTokenInfoPriceFeedUpdated is a log parse operation binding the contract event 0x9de4f0e361ff2892765ceacf345aee96b96513207cd9ebee35bd9ee67849b008.
//
// Solidity: event BridgeTokenInfoPriceFeedUpdated(address priceFeed)
func (_BridgeTokenInfo *BridgeTokenInfoFilterer) ParseBridgeTokenInfoPriceFeedUpdated(log types.Log) (*BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated, error) {
	event := new(BridgeTokenInfoBridgeTokenInfoPriceFeedUpdated)
	if err := _BridgeTokenInfo.contract.UnpackLog(event, "BridgeTokenInfoPriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
