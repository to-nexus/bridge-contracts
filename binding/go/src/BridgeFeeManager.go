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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"setGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minimumValueList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setTokenFeeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeManagerExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeManagerFinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeManagerGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"BridgeFeeManagerPriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeFeeManagerCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeManagerChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeFeeManagerInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"df6e87dc": "calculateFee(uint256,address,uint256)",
		"96ce0795": "denominator()",
		"d1c01543": "getGasPrice(uint256)",
		"09585b1c": "getTokenFee(uint256,address)",
		"8da5cb5b": "owner()",
		"ce2e5c66": "removePriceFeed()",
		"1b034c1c": "removeTokenFee(address)",
		"715018a6": "renounceOwnership()",
		"005d0f10": "setExFeeRate(uint256)",
		"f289d3ba": "setFinalizeBridgeGas(uint256)",
		"a7cb0507": "setGasPrice(uint256,uint256)",
		"724e78da": "setPriceFeed(address)",
		"dffa229e": "setTokenFee(address,uint256,uint256)",
		"6731d860": "setTokenFeeBatch(address[],uint256[],uint256[])",
		"f2fde38b": "transferOwnership(address)",
		"1f96131e": "updateGasPrice(uint256,uint256)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b50604051610e01380380610e0183398101604081905261002e916100fc565b338061005457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005d816100ad565b50805f036100a25760405163449c1d2d60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b604482015260640161004b565b60035560025561011e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5f6040838503121561010d575f5ffd5b505080516020909101519092909150565b610cd68061012b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100fa575f3560e01c806396ce079511610093578063df6e87dc11610063578063df6e87dc14610200578063dffa229e14610213578063f289d3ba14610226578063f2fde38b14610239575f5ffd5b806396ce0795146101b4578063a7cb0507146101c6578063ce2e5c66146101d9578063d1c01543146101e1575f5ffd5b80636731d860116100ce5780636731d8601461016c578063715018a61461017f578063724e78da146101875780638da5cb5b1461019a575f5ffd5b80625d0f10146100fe57806309585b1c146101135780631b034c1c146101465780631f96131e14610159575b5f5ffd5b61011161010c36600461094a565b61024c565b005b610126610121366004610975565b610290565b604080519384526020840192909252908201526060015b60405180910390f35b6101116101543660046109a3565b6102c0565b6101116101673660046109c5565b610361565b61011161017a366004610ab7565b6103b9565b61011161045f565b6101116101953660046109a3565b610472565b5f546040516001600160a01b03909116815260200161013d565b6103e85b60405190815260200161013d565b6101116101d43660046109c5565b610506565b61011161058c565b6101b86101ef36600461094a565b5f9081526005602052604090205490565b61012661020e366004610ba5565b6105d0565b610111610221366004610bda565b610608565b61011161023436600461094a565b6106a7565b6101116102473660046109a3565b610728565b610254610762565b60028190556040518181527fe455d85fca798ed469f404b1acfb31528a2dfc37ccf9f81ca0a667c6d5bae81a906020015b60405180910390a150565b5f5f5f61029c8461078e565b90935090506102ab85856107fa565b509150600181016102b957505f5b9250925092565b6102c8610762565b6001600160a01b03811661030c5760405163449c1d2d60e11b81526020600482015260056024820152643a37b5b2b760d91b60448201526064015b60405180910390fd5b6001600160a01b038082165f818152600460205260409020549091160361035e576001600160a01b0381165f90815260046020526040812080546001600160a01b031916815560018101829055600201555b50565b610369610762565b5f82815260056020526040908190208290555182907fd712f532f9c8595d0e4cc71ce07478db83f369b931351cc340eb33236cb7e050906103ad9084815260200190565b60405180910390a25050565b6103c1610762565b815183511480156103d3575080518351145b6103f0576040516371a485af60e01b815260040160405180910390fd5b5f5b83518110156104595761045184828151811061041057610410610c0c565b602002602001015184838151811061042a5761042a610c0c565b602002602001015184848151811061044457610444610c0c565b6020026020010151610608565b6001016103f2565b50505050565b610467610762565b6104705f6108fb565b565b61047a610762565b6001600160a01b0381166104bd5760405163449c1d2d60e11b81526020600482015260096024820152681c1c9a58d95199595960ba1b6044820152606401610303565b600180546001600160a01b0319166001600160a01b0383169081179091556040517fe27c56cb0e67232888202b1667d878cf050bc31c0710d2bcbbd5ba873972c649905f90a250565b61050e610762565b5f8281526005602052604090205482901561053f57604051630b088bef60e41b815260040161030391815260200190565b50805f0361057b5760405163449c1d2d60e11b8152602060048201526008602482015267676173507269636560c01b6044820152606401610303565b5f9182526005602052604090912055565b610594610762565b600180546001600160a01b03191690556040515f907fe27c56cb0e67232888202b1667d878cf050bc31c0710d2bcbbd5ba873972c649908290a2565b5f5f5f5f6105de8787610290565b919550935090506103e86105f28287610c20565b6105fc9190610c49565b91505093509350939050565b610610610762565b6001600160a01b03831661064f5760405163449c1d2d60e11b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610303565b604080516060810182526001600160a01b0394851680825260208083019586528284019485525f9182526004905291909120905181546001600160a01b03191694169390931783559051600183015551600290910155565b6106af610762565b805f036106f35760405163449c1d2d60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606401610303565b60038190556040518181527f5aaa51219688d46d1958215f2c48a8711727bf88d24f6e75af9acf9c5442072a90602001610285565b610730610762565b6001600160a01b03811661075957604051631e4fbdf760e01b81525f6004820152602401610303565b61035e816108fb565b5f546001600160a01b031633146104705760405163118cdaa760e01b8152336004820152602401610303565b6001600160a01b038082165f818152600460209081526040808320815160608101835281549096168087526001820154938701939093526002015490850152909283929091036107ec57806020015181604001519250925050915091565b5f6002549250925050915091565b5f828152600560205260408120546001548291906001600160a01b03161580610821575080155b15610832575f5f92509250506108f4565b60015460035473__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9163889ad9e0916001600160a01b03909116908890889061086f908790610c20565b6040516001600160e01b031960e087901b1681526001600160a01b0394851660048201526024810193909352921660448201526064810191909152608401606060405180830381865af41580156108c8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ec9190610c68565b909450925050505b9250929050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020828403121561095a575f5ffd5b5035919050565b6001600160a01b038116811461035e575f5ffd5b5f5f60408385031215610986575f5ffd5b82359150602083013561099881610961565b809150509250929050565b5f602082840312156109b3575f5ffd5b81356109be81610961565b9392505050565b5f5f604083850312156109d6575f5ffd5b50508035926020909101359150565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610a2257610a226109e5565b604052919050565b5f67ffffffffffffffff821115610a4357610a436109e5565b5060051b60200190565b5f82601f830112610a5c575f5ffd5b8135610a6f610a6a82610a2a565b6109f9565b8082825260208201915060208360051b860101925085831115610a90575f5ffd5b602085015b83811015610aad578035835260209283019201610a95565b5095945050505050565b5f5f5f60608486031215610ac9575f5ffd5b833567ffffffffffffffff811115610adf575f5ffd5b8401601f81018613610aef575f5ffd5b8035610afd610a6a82610a2a565b8082825260208201915060208360051b850101925088831115610b1e575f5ffd5b6020840193505b82841015610b49578335610b3881610961565b825260209384019390910190610b25565b9550505050602084013567ffffffffffffffff811115610b67575f5ffd5b610b7386828701610a4d565b925050604084013567ffffffffffffffff811115610b8f575f5ffd5b610b9b86828701610a4d565b9150509250925092565b5f5f5f60608486031215610bb7575f5ffd5b833592506020840135610bc981610961565b929592945050506040919091013590565b5f5f5f60608486031215610bec575f5ffd5b8335610bf781610961565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b8082028115828204841417610c4357634e487b7160e01b5f52601160045260245ffd5b92915050565b5f82610c6357634e487b7160e01b5f52601260045260245ffd5b500490565b5f5f5f60608486031215610c7a575f5ffd5b83518015158114610c89575f5ffd5b60208501516040909501519096949550939250505056fea2646970667358221220751e9c89b9270df96ce3665bdd0db63bf08b16f9b896f3b1821f78a2c0d3caf964736f6c634300081c0033",
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
func DeployBridgeFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend, exFeeRate *big.Int, finalizeBridgeGas *big.Int) (common.Address, *types.Transaction, *BridgeFeeManager, error) {
	parsed, err := BridgeFeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeFeeManagerBin), backend, exFeeRate, finalizeBridgeGas)
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

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeFeeManager *BridgeFeeManagerCaller) CalculateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "calculateFee", remoteChainID, token, value)

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
func (_BridgeFeeManager *BridgeFeeManagerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeFeeManager.Contract.CalculateFee(&_BridgeFeeManager.CallOpts, remoteChainID, token, value)
}

// CalculateFee is a free data retrieval call binding the contract method 0xdf6e87dc.
//
// Solidity: function calculateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFee)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) CalculateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFee        *big.Int
}, error) {
	return _BridgeFeeManager.Contract.CalculateFee(&_BridgeFeeManager.CallOpts, remoteChainID, token, value)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
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
// Solidity: function denominator() pure returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerSession) Denominator() (*big.Int, error) {
	return _BridgeFeeManager.Contract.Denominator(&_BridgeFeeManager.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) Denominator() (*big.Int, error) {
	return _BridgeFeeManager.Contract.Denominator(&_BridgeFeeManager.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCaller) GetGasPrice(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "getGasPrice", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeFeeManager.Contract.GetGasPrice(&_BridgeFeeManager.CallOpts, remoteChainID)
}

// GetGasPrice is a free data retrieval call binding the contract method 0xd1c01543.
//
// Solidity: function getGasPrice(uint256 remoteChainID) view returns(uint256)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) GetGasPrice(remoteChainID *big.Int) (*big.Int, error) {
	return _BridgeFeeManager.Contract.GetGasPrice(&_BridgeFeeManager.CallOpts, remoteChainID)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerCaller) GetTokenFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	var out []interface{}
	err := _BridgeFeeManager.contract.Call(opts, &out, "getTokenFee", remoteChainID, token)

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

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerSession) GetTokenFee(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeFeeManager.Contract.GetTokenFee(&_BridgeFeeManager.CallOpts, remoteChainID, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumValue, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerCallerSession) GetTokenFee(remoteChainID *big.Int, token common.Address) (struct {
	MinimumValue *big.Int
	GasFee       *big.Int
	ExFeeRate    *big.Int
}, error) {
	return _BridgeFeeManager.Contract.GetTokenFee(&_BridgeFeeManager.CallOpts, remoteChainID, token)
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

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) RemovePriceFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "removePriceFeed")
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemovePriceFeed(&_BridgeFeeManager.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemovePriceFeed(&_BridgeFeeManager.TransactOpts)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) RemoveTokenFee(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "removeTokenFee", token)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) RemoveTokenFee(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveTokenFee(&_BridgeFeeManager.TransactOpts, token)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) RemoveTokenFee(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.RemoveTokenFee(&_BridgeFeeManager.TransactOpts, token)
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

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetExFeeRate(opts *bind.TransactOpts, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setExFeeRate", exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetExFeeRate(&_BridgeFeeManager.TransactOpts, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetExFeeRate(&_BridgeFeeManager.TransactOpts, exFeeRate)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetFinalizeBridgeGas(opts *bind.TransactOpts, finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setFinalizeBridgeGas", finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetFinalizeBridgeGas(&_BridgeFeeManager.TransactOpts, finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetFinalizeBridgeGas(&_BridgeFeeManager.TransactOpts, finalizeBridgeGas)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetGasPrice(opts *bind.TransactOpts, remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setGasPrice", remoteChainID, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetGasPrice(&_BridgeFeeManager.TransactOpts, remoteChainID, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetGasPrice(&_BridgeFeeManager.TransactOpts, remoteChainID, gasPrice)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetPriceFeed(&_BridgeFeeManager.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetPriceFeed(&_BridgeFeeManager.TransactOpts, priceFeed)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumValue, uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, minimumValue *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setTokenFee", token, minimumValue, exFeeRate)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumValue, uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetTokenFee(token common.Address, minimumValue *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetTokenFee(&_BridgeFeeManager.TransactOpts, token, minimumValue, exFeeRate)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumValue, uint256 exFeeRate) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetTokenFee(token common.Address, minimumValue *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetTokenFee(&_BridgeFeeManager.TransactOpts, token, minimumValue, exFeeRate)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumValueList, uint256[] exFeeRateList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) SetTokenFeeBatch(opts *bind.TransactOpts, tokenList []common.Address, minimumValueList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "setTokenFeeBatch", tokenList, minimumValueList, exFeeRateList)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumValueList, uint256[] exFeeRateList) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) SetTokenFeeBatch(tokenList []common.Address, minimumValueList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetTokenFeeBatch(&_BridgeFeeManager.TransactOpts, tokenList, minimumValueList, exFeeRateList)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumValueList, uint256[] exFeeRateList) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) SetTokenFeeBatch(tokenList []common.Address, minimumValueList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.SetTokenFeeBatch(&_BridgeFeeManager.TransactOpts, tokenList, minimumValueList, exFeeRateList)
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

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactor) UpdateGasPrice(opts *bind.TransactOpts, remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.contract.Transact(opts, "updateGasPrice", remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateGasPrice(&_BridgeFeeManager.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeManager *BridgeFeeManagerTransactorSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeManager.Contract.UpdateGasPrice(&_BridgeFeeManager.TransactOpts, remoteChainID, gasPrice)
}

// BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator is returned from FilterBridgeFeeManagerExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeManagerExchangeFeeUpdated events raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator struct {
	Event *BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated)
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
		it.Event = new(BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated)
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
func (it *BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated represents a BridgeFeeManagerExchangeFeeUpdated event raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated struct {
	ExFeeRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeManagerExchangeFeeUpdated is a free log retrieval operation binding the contract event 0xe455d85fca798ed469f404b1acfb31528a2dfc37ccf9f81ca0a667c6d5bae81a.
//
// Solidity: event BridgeFeeManagerExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterBridgeFeeManagerExchangeFeeUpdated(opts *bind.FilterOpts) (*BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator, error) {

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "BridgeFeeManagerExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdatedIterator{contract: _BridgeFeeManager.contract, event: "BridgeFeeManagerExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeManagerExchangeFeeUpdated is a free log subscription operation binding the contract event 0xe455d85fca798ed469f404b1acfb31528a2dfc37ccf9f81ca0a667c6d5bae81a.
//
// Solidity: event BridgeFeeManagerExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchBridgeFeeManagerExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "BridgeFeeManagerExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerExchangeFeeUpdated", log); err != nil {
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

// ParseBridgeFeeManagerExchangeFeeUpdated is a log parse operation binding the contract event 0xe455d85fca798ed469f404b1acfb31528a2dfc37ccf9f81ca0a667c6d5bae81a.
//
// Solidity: event BridgeFeeManagerExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseBridgeFeeManagerExchangeFeeUpdated(log types.Log) (*BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated, error) {
	event := new(BridgeFeeManagerBridgeFeeManagerExchangeFeeUpdated)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator is returned from FilterBridgeFeeManagerFinalizeBridgeGasSet and is used to iterate over the raw logs and unpacked data for BridgeFeeManagerFinalizeBridgeGasSet events raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator struct {
	Event *BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet)
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
		it.Event = new(BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet)
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
func (it *BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet represents a BridgeFeeManagerFinalizeBridgeGasSet event raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet struct {
	FinalizeBridgeGas *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeManagerFinalizeBridgeGasSet is a free log retrieval operation binding the contract event 0x5aaa51219688d46d1958215f2c48a8711727bf88d24f6e75af9acf9c5442072a.
//
// Solidity: event BridgeFeeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterBridgeFeeManagerFinalizeBridgeGasSet(opts *bind.FilterOpts) (*BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator, error) {

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "BridgeFeeManagerFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSetIterator{contract: _BridgeFeeManager.contract, event: "BridgeFeeManagerFinalizeBridgeGasSet", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeManagerFinalizeBridgeGasSet is a free log subscription operation binding the contract event 0x5aaa51219688d46d1958215f2c48a8711727bf88d24f6e75af9acf9c5442072a.
//
// Solidity: event BridgeFeeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchBridgeFeeManagerFinalizeBridgeGasSet(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet) (event.Subscription, error) {

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "BridgeFeeManagerFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerFinalizeBridgeGasSet", log); err != nil {
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

// ParseBridgeFeeManagerFinalizeBridgeGasSet is a log parse operation binding the contract event 0x5aaa51219688d46d1958215f2c48a8711727bf88d24f6e75af9acf9c5442072a.
//
// Solidity: event BridgeFeeManagerFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseBridgeFeeManagerFinalizeBridgeGasSet(log types.Log) (*BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet, error) {
	event := new(BridgeFeeManagerBridgeFeeManagerFinalizeBridgeGasSet)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerFinalizeBridgeGasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator is returned from FilterBridgeFeeManagerGasPriceUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeManagerGasPriceUpdated events raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator struct {
	Event *BridgeFeeManagerBridgeFeeManagerGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerBridgeFeeManagerGasPriceUpdated)
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
		it.Event = new(BridgeFeeManagerBridgeFeeManagerGasPriceUpdated)
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
func (it *BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerBridgeFeeManagerGasPriceUpdated represents a BridgeFeeManagerGasPriceUpdated event raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerGasPriceUpdated struct {
	RemoteChainID *big.Int
	GasPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeManagerGasPriceUpdated is a free log retrieval operation binding the contract event 0xd712f532f9c8595d0e4cc71ce07478db83f369b931351cc340eb33236cb7e050.
//
// Solidity: event BridgeFeeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterBridgeFeeManagerGasPriceUpdated(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "BridgeFeeManagerGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerBridgeFeeManagerGasPriceUpdatedIterator{contract: _BridgeFeeManager.contract, event: "BridgeFeeManagerGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeManagerGasPriceUpdated is a free log subscription operation binding the contract event 0xd712f532f9c8595d0e4cc71ce07478db83f369b931351cc340eb33236cb7e050.
//
// Solidity: event BridgeFeeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchBridgeFeeManagerGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerBridgeFeeManagerGasPriceUpdated, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "BridgeFeeManagerGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerBridgeFeeManagerGasPriceUpdated)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerGasPriceUpdated", log); err != nil {
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

// ParseBridgeFeeManagerGasPriceUpdated is a log parse operation binding the contract event 0xd712f532f9c8595d0e4cc71ce07478db83f369b931351cc340eb33236cb7e050.
//
// Solidity: event BridgeFeeManagerGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseBridgeFeeManagerGasPriceUpdated(log types.Log) (*BridgeFeeManagerBridgeFeeManagerGasPriceUpdated, error) {
	event := new(BridgeFeeManagerBridgeFeeManagerGasPriceUpdated)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator is returned from FilterBridgeFeeManagerPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeManagerPriceFeedUpdated events raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator struct {
	Event *BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated)
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
		it.Event = new(BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated)
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
func (it *BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated represents a BridgeFeeManagerPriceFeedUpdated event raised by the BridgeFeeManager contract.
type BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeManagerPriceFeedUpdated is a free log retrieval operation binding the contract event 0xe27c56cb0e67232888202b1667d878cf050bc31c0710d2bcbbd5ba873972c649.
//
// Solidity: event BridgeFeeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) FilterBridgeFeeManagerPriceFeedUpdated(opts *bind.FilterOpts, priceFeed []common.Address) (*BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.FilterLogs(opts, "BridgeFeeManagerPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeManagerBridgeFeeManagerPriceFeedUpdatedIterator{contract: _BridgeFeeManager.contract, event: "BridgeFeeManagerPriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeManagerPriceFeedUpdated is a free log subscription operation binding the contract event 0xe27c56cb0e67232888202b1667d878cf050bc31c0710d2bcbbd5ba873972c649.
//
// Solidity: event BridgeFeeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) WatchBridgeFeeManagerPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeFeeManager.contract.WatchLogs(opts, "BridgeFeeManagerPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated)
				if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerPriceFeedUpdated", log); err != nil {
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

// ParseBridgeFeeManagerPriceFeedUpdated is a log parse operation binding the contract event 0xe27c56cb0e67232888202b1667d878cf050bc31c0710d2bcbbd5ba873972c649.
//
// Solidity: event BridgeFeeManagerPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeManager *BridgeFeeManagerFilterer) ParseBridgeFeeManagerPriceFeedUpdated(log types.Log) (*BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated, error) {
	event := new(BridgeFeeManagerBridgeFeeManagerPriceFeedUpdated)
	if err := _BridgeFeeManager.contract.UnpackLog(event, "BridgeFeeManagerPriceFeedUpdated", log); err != nil {
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
