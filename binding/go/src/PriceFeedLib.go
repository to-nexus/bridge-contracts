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

// PriceFeedLibMetaData contains all meta data concerning the PriceFeedLib contract.
var PriceFeedLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"}],\"name\":\"allPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"prices\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"convertAtoB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"convertAtoBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"999d5cf4": "allPrices(IPriceFeed)",
		"126e5b3a": "convertAtoB(IPriceFeed,address,address,uint256)",
		"957bb22e": "convertAtoBWithPrice(uint256,uint256,uint256,uint8,uint8)",
	},
	Bin: "0x610cd5610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c8063126e5b3a1461004e578063957bb22e1461007b578063999d5cf41461009c575b5f5ffd5b61006161005c36600461077e565b6100bc565b604080519283529015156020830152015b60405180910390f35b61008e6100893660046107da565b6103c6565b604051908152602001610072565b6100af6100aa36600461082c565b61049b565b6040516100729190610847565b5f5f836001600160a01b0316856001600160a01b0316036100e2575081905060016103bd565b5f866001600160a01b03166329dcb0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561011f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061014391906108ab565b6040805160028082526060820183529293505f92909160208301908036833701905050905086815f8151811061017b5761017b6108d6565b60200260200101906001600160a01b031690816001600160a01b03168152505085816001815181106101af576101af6108d6565b6001600160a01b0392831660209182029290920101526040516347dad24160e11b81525f918a1690638fb5a482906101eb9085906004016108ea565b5f60405180830381865afa158015610205573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261022c91908101906109a7565b9050876001600160a01b0316815f8151811061024a5761024a6108d6565b60200260200101515f01516001600160a01b03161480156102995750866001600160a01b031681600181518110610283576102836108d6565b60200260200101515f01516001600160a01b0316145b6102a1575f5ffd5b5f19815f815181106102b5576102b56108d6565b60200260200101516040015114806102f557504283825f815181106102dc576102dc6108d6565b6020026020010151604001516102f29190610a89565b10155b801561035157505f1981600181518110610311576103116108d6565b60200260200101516040015114806103515750428382600181518110610339576103396108d6565b60200260200101516040015161034f9190610a89565b115b93505f5f61035f8b8b610571565b6103698c8b610571565b915091506103b588845f81518110610383576103836108d6565b602002602001015160200151856001815181106103a2576103a26108d6565b60200260200101516020015185856103c6565b965050505050505b94509492505050565b5f845f0361040457604051633a0b0f2960e01b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff16101561044c57610447866104218587610a9c565b61042c90600a610b98565b6104369088610bba565b6104409190610bba565b8890610659565b61046f565b61046f61044061045c8686610a9c565b61046790600a610b98565b8790896106a0565b925090508061049157604051633961e4cf60e11b815260040160405180910390fd5b5095945050505050565b6060816001600160a01b0316638fb5a482836001600160a01b0316636ff97f1d6040518163ffffffff1660e01b81526004015f60405180830381865afa1580156104e7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261050e9190810190610bd9565b6040518263ffffffff1660e01b815260040161052a91906108ea565b5f60405180830381865afa158015610544573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261056b91908101906109a7565b92915050565b5f826001600160a01b03166311df99956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ae573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105d29190610c69565b6001600160a01b0316826001600160a01b03161461064f57816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610626573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064a9190610c84565b610652565b60125b9392505050565b5f5f835f0361066d5750600190505f610699565b8383028385828161068057610680610ba6565b0414610692575f5f9250925050610699565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036106d4578382816106ca576106ca610ba6565b0492505050610652565b8084116106eb576106eb6003851502601118610756565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6001600160a01b038116811461077b575f5ffd5b50565b5f5f5f5f60808587031215610791575f5ffd5b843561079c81610767565b935060208501356107ac81610767565b925060408501356107bc81610767565b9396929550929360600135925050565b60ff8116811461077b575f5ffd5b5f5f5f5f5f60a086880312156107ee575f5ffd5b853594506020860135935060408601359250606086013561080e816107cc565b9150608086013561081e816107cc565b809150509295509295909350565b5f6020828403121561083c575f5ffd5b813561065281610767565b602080825282518282018190525f918401906040840190835b818110156108a057835180516001600160a01b03168452602080820151818601526040918201519185019190915290930192606090920191600101610860565b509095945050505050565b5f602082840312156108bb575f5ffd5b5051919050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b818110156108a05783516001600160a01b0316835260209384019390920191600101610903565b6040516060810167ffffffffffffffff8111828210171561094d5761094d6108c2565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561097c5761097c6108c2565b604052919050565b5f67ffffffffffffffff82111561099d5761099d6108c2565b5060051b60200190565b5f602082840312156109b7575f5ffd5b815167ffffffffffffffff8111156109cd575f5ffd5b8201601f810184136109dd575f5ffd5b80516109f06109eb82610984565b610953565b80828252602082019150602060608402850101925086831115610a11575f5ffd5b6020840193505b82841015610a6b5760608488031215610a2f575f5ffd5b610a3761092a565b8451610a4281610767565b815260208581015181830152604080870151908301529083526060909401939190910190610a18565b9695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561056b5761056b610a75565b60ff828116828216039081111561056b5761056b610a75565b6001815b6001841115610af057808504811115610ad457610ad4610a75565b6001841615610ae257908102905b60019390931c928002610ab9565b935093915050565b5f82610b065750600161056b565b81610b1257505f61056b565b8160018114610b285760028114610b3257610b4e565b600191505061056b565b60ff841115610b4357610b43610a75565b50506001821b61056b565b5060208310610133831016604e8410600b8410161715610b71575081810a61056b565b610b7d5f198484610ab5565b805f1904821115610b9057610b90610a75565b029392505050565b5f61065260ff841683610af8565b634e487b7160e01b5f52601260045260245ffd5b5f82610bd457634e487b7160e01b5f52601260045260245ffd5b500490565b5f60208284031215610be9575f5ffd5b815167ffffffffffffffff811115610bff575f5ffd5b8201601f81018413610c0f575f5ffd5b8051610c1d6109eb82610984565b8082825260208201915060208360051b850101925086831115610c3e575f5ffd5b6020840193505b82841015610a6b578351610c5881610767565b825260209384019390910190610c45565b5f60208284031215610c79575f5ffd5b815161065281610767565b5f60208284031215610c94575f5ffd5b8151610652816107cc56fea26469706673582212205945f5c7e8fec02ab1e970ec4476c064dce4001d68eefe1f52e998762562bf8b64736f6c634300081c0033",
}

// PriceFeedLibABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedLibMetaData.ABI instead.
var PriceFeedLibABI = PriceFeedLibMetaData.ABI

// Deprecated: Use PriceFeedLibMetaData.Sigs instead.
// PriceFeedLibFuncSigs maps the 4-byte function signature to its string representation.
var PriceFeedLibFuncSigs = PriceFeedLibMetaData.Sigs

// PriceFeedLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedLibMetaData.Bin instead.
var PriceFeedLibBin = PriceFeedLibMetaData.Bin

// DeployPriceFeedLib deploys a new Ethereum contract, binding an instance of PriceFeedLib to it.
func DeployPriceFeedLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeedLib, error) {
	parsed, err := PriceFeedLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeedLib{PriceFeedLibCaller: PriceFeedLibCaller{contract: contract}, PriceFeedLibTransactor: PriceFeedLibTransactor{contract: contract}, PriceFeedLibFilterer: PriceFeedLibFilterer{contract: contract}}, nil
}

// PriceFeedLib is an auto generated Go binding around an Ethereum contract.
type PriceFeedLib struct {
	PriceFeedLibCaller     // Read-only binding to the contract
	PriceFeedLibTransactor // Write-only binding to the contract
	PriceFeedLibFilterer   // Log filterer for contract events
}

// PriceFeedLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedLibSession struct {
	Contract     *PriceFeedLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedLibCallerSession struct {
	Contract *PriceFeedLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PriceFeedLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedLibTransactorSession struct {
	Contract     *PriceFeedLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PriceFeedLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedLibRaw struct {
	Contract *PriceFeedLib // Generic contract binding to access the raw methods on
}

// PriceFeedLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedLibCallerRaw struct {
	Contract *PriceFeedLibCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedLibTransactorRaw struct {
	Contract *PriceFeedLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedLib creates a new instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLib(address common.Address, backend bind.ContractBackend) (*PriceFeedLib, error) {
	contract, err := bindPriceFeedLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLib{PriceFeedLibCaller: PriceFeedLibCaller{contract: contract}, PriceFeedLibTransactor: PriceFeedLibTransactor{contract: contract}, PriceFeedLibFilterer: PriceFeedLibFilterer{contract: contract}}, nil
}

// NewPriceFeedLibCaller creates a new read-only instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedLibCaller, error) {
	contract, err := bindPriceFeedLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibCaller{contract: contract}, nil
}

// NewPriceFeedLibTransactor creates a new write-only instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedLibTransactor, error) {
	contract, err := bindPriceFeedLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibTransactor{contract: contract}, nil
}

// NewPriceFeedLibFilterer creates a new log filterer instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedLibFilterer, error) {
	contract, err := bindPriceFeedLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibFilterer{contract: contract}, nil
}

// bindPriceFeedLib binds a generic wrapper to an already deployed contract.
func bindPriceFeedLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedLib *PriceFeedLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedLib.Contract.PriceFeedLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedLib *PriceFeedLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.PriceFeedLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedLib *PriceFeedLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.PriceFeedLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedLib *PriceFeedLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedLib *PriceFeedLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedLib *PriceFeedLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.contract.Transact(opts, method, params...)
}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibCaller) AllPrices(opts *bind.CallOpts, feed common.Address) ([]IPriceFeedPriceData, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "allPrices", feed)

	if err != nil {
		return *new([]IPriceFeedPriceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPriceFeedPriceData)).(*[]IPriceFeedPriceData)

	return out0, err

}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibSession) AllPrices(feed common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedLib.Contract.AllPrices(&_PriceFeedLib.CallOpts, feed)
}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibCallerSession) AllPrices(feed common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedLib.Contract.AllPrices(&_PriceFeedLib.CallOpts, feed)
}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibCaller) ConvertAtoB(opts *bind.CallOpts, feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "convertAtoB", feed, tokenA, tokenB, amountA)

	outstruct := new(struct {
		AmountB *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountB = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibSession) ConvertAtoB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	return _PriceFeedLib.Contract.ConvertAtoB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibCallerSession) ConvertAtoB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	return _PriceFeedLib.Contract.ConvertAtoB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCaller) ConvertAtoBWithPrice(opts *bind.CallOpts, amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "convertAtoBWithPrice", amountA, priceA, priceB, decimalA, decimalB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibSession) ConvertAtoBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.ConvertAtoBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCallerSession) ConvertAtoBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.ConvertAtoBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}
