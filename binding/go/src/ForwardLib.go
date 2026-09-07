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

// ForwardLibMetaData contains all meta data concerning the ForwardLib contract.
var ForwardLibMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"finalizeIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumForwardLib.ForwardFailureCode\",\"name\":\"code\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reasonHash\",\"type\":\"bytes32\"}],\"name\":\"ForwardFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"finalizeIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"extraDataHash\",\"type\":\"bytes32\"}],\"name\":\"ForwardInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BaseBridgeForwardContextInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwardReentrantCall\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a3f11c4f": "acquireGuard()",
		"4cfbad03": "consumeCtx()",
		"561f601c": "emitInitiated(uint256,uint256,uint256,address,address,uint256,uint256,uint256,bytes32)",
		"3648cc4e": "endAndReport(uint256,uint256,bool,bool,bytes)",
		"2f60d071": "releaseGuard()",
		"ffb2e998": "setCtx(uint256,address,uint256,uint256)",
	},
	Bin: "0x60808060405234601b5761056590816100208239308160080152f35b5f80fdfe608080604052307f000000000000000000000000000000000000000000000000000000000000000014906004361015610036575f80fd5b5f3560e01c9081632f60d071146103575781633648cc4e14610294575080634cfbad03146101e0578063561f601c14610146578063a3f11c4f146101015763ffb2e99814610082575f80fd5b6100fd5760803660031901126100fd576024356001600160a01b03811681036100fd576004355f5160206105105f395f51905f525d5f5160206104b05f395f51905f525d6044355f5160206104d05f395f51905f525d6064355f5160206104f05f395f51905f525d60015f5160206104705f395f51905f525d005b5f80fd5b506100fd575f3660031901126100fd575f5160206104905f395f51905f525c6101375760015f5160206104905f395f51905f525d005b63ccd78ee760e01b5f5260045ffd5b506100fd576101203660031901126100fd576064356001600160a01b038116908190036100fd576084356001600160a01b038116908190036100fd57604051918252602082015260a435604082015260c435606082015260e43560808201526101043560a082015260443590602435907fd9c6b693f2a3033b5f515be724c03b8f30368cc278e689735e5edb7c6aa9bc5660c060043592a4005b506100fd575f3660031901126100fd575f5160206104705f395f51905f525c156102855760805f5160206105105f395f51905f52805c905f5160206104b05f395f51905f52805c5f5f5160206104d05f395f51905f529281845c94815f5160206104f05f395f51905f529381855c99815f5160206104705f395f51905f525d5d5d5d5d604080519485526001600160a01b039190911660208501528301526060820152f35b6306e6bb5b60e51b5f5260045ffd5b826100fd5760a03660031901126100fd576044359081151582036100fd5760643580151581036100fd57608435926001600160401b0384116100fd57366023850112156100fd5760048401356001600160401b03811161034357601f8101601f19908116603f011684016001600160401b038111858210176103435760405280845236602482870101116100fd576020815f92602461034198018388013785010152602435600435610379565b005b634e487b7160e01b5f52604160045260245ffd5b826100fd575f3660031901126100fd575f5f5160206104905f395f51905f525d005b929390935f5160206104705f395f51905f525c1515905f5f5160206104705f395f51905f525d5f5f5160206105105f395f51905f525d5f5f5160206104b05f395f51905f525d5f5f5160206104d05f395f51905f525d5f5f5160206104f05f395f51905f525d8361044f5750506001915b156104475760208151910120905b6040516004821015610433577f606a6de1b36fbe9108fe6f544d5941e752ff8bd36b5a01f08068cf7943e7e3df9260409282526020820152a3565b634e487b7160e01b5f52602160045260245ffd5b505f906103f8565b61045c57506002916103ea565b15610469576003916103ea565b5050505056fef985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f40096c5fec86ccff8d1ee06d7ba5a223396fbf7f905e4727220bd20cf4e693d0500f985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f402f985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f403f985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f404f985886670fe7bb8de789a16f5a5fac4e3acb178f2710001e0f3cfde4fc8f401a2646970667358221220045090a1702a006ab1e4b52cc98e8f846a03271a3650e4762177e378a09ac58b64736f6c634300081c0033",
}

// ForwardLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwardLibMetaData.ABI instead.
var ForwardLibABI = ForwardLibMetaData.ABI

// Deprecated: Use ForwardLibMetaData.Sigs instead.
// ForwardLibFuncSigs maps the 4-byte function signature to its string representation.
var ForwardLibFuncSigs = ForwardLibMetaData.Sigs

// ForwardLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForwardLibMetaData.Bin instead.
var ForwardLibBin = ForwardLibMetaData.Bin

// DeployForwardLib deploys a new Ethereum contract, binding an instance of ForwardLib to it.
func DeployForwardLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ForwardLib, error) {
	parsed, err := ForwardLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwardLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForwardLib{ForwardLibCaller: ForwardLibCaller{contract: contract}, ForwardLibTransactor: ForwardLibTransactor{contract: contract}, ForwardLibFilterer: ForwardLibFilterer{contract: contract}}, nil
}

// ForwardLib is an auto generated Go binding around an Ethereum contract.
type ForwardLib struct {
	ForwardLibCaller     // Read-only binding to the contract
	ForwardLibTransactor // Write-only binding to the contract
	ForwardLibFilterer   // Log filterer for contract events
}

// ForwardLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwardLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwardLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwardLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwardLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwardLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwardLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwardLibSession struct {
	Contract     *ForwardLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwardLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwardLibCallerSession struct {
	Contract *ForwardLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ForwardLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwardLibTransactorSession struct {
	Contract     *ForwardLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ForwardLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwardLibRaw struct {
	Contract *ForwardLib // Generic contract binding to access the raw methods on
}

// ForwardLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwardLibCallerRaw struct {
	Contract *ForwardLibCaller // Generic read-only contract binding to access the raw methods on
}

// ForwardLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwardLibTransactorRaw struct {
	Contract *ForwardLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwardLib creates a new instance of ForwardLib, bound to a specific deployed contract.
func NewForwardLib(address common.Address, backend bind.ContractBackend) (*ForwardLib, error) {
	contract, err := bindForwardLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForwardLib{ForwardLibCaller: ForwardLibCaller{contract: contract}, ForwardLibTransactor: ForwardLibTransactor{contract: contract}, ForwardLibFilterer: ForwardLibFilterer{contract: contract}}, nil
}

// NewForwardLibCaller creates a new read-only instance of ForwardLib, bound to a specific deployed contract.
func NewForwardLibCaller(address common.Address, caller bind.ContractCaller) (*ForwardLibCaller, error) {
	contract, err := bindForwardLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwardLibCaller{contract: contract}, nil
}

// NewForwardLibTransactor creates a new write-only instance of ForwardLib, bound to a specific deployed contract.
func NewForwardLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwardLibTransactor, error) {
	contract, err := bindForwardLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwardLibTransactor{contract: contract}, nil
}

// NewForwardLibFilterer creates a new log filterer instance of ForwardLib, bound to a specific deployed contract.
func NewForwardLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwardLibFilterer, error) {
	contract, err := bindForwardLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwardLibFilterer{contract: contract}, nil
}

// bindForwardLib binds a generic wrapper to an already deployed contract.
func bindForwardLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForwardLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwardLib *ForwardLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForwardLib.Contract.ForwardLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwardLib *ForwardLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwardLib.Contract.ForwardLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwardLib *ForwardLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwardLib.Contract.ForwardLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwardLib *ForwardLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForwardLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwardLib *ForwardLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwardLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwardLib *ForwardLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwardLib.Contract.contract.Transact(opts, method, params...)
}

// ForwardLibForwardFailedIterator is returned from FilterForwardFailed and is used to iterate over the raw logs and unpacked data for ForwardFailed events raised by the ForwardLib contract.
type ForwardLibForwardFailedIterator struct {
	Event *ForwardLibForwardFailed // Event containing the contract specifics and raw log

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
func (it *ForwardLibForwardFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwardLibForwardFailed)
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
		it.Event = new(ForwardLibForwardFailed)
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
func (it *ForwardLibForwardFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwardLibForwardFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwardLibForwardFailed represents a ForwardFailed event raised by the ForwardLib contract.
type ForwardLibForwardFailed struct {
	FromChainID   *big.Int
	FinalizeIndex *big.Int
	Code          uint8
	ReasonHash    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterForwardFailed is a free log retrieval operation binding the contract event 0x606a6de1b36fbe9108fe6f544d5941e752ff8bd36b5a01f08068cf7943e7e3df.
//
// Solidity: event ForwardFailed(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint8 code, bytes32 reasonHash)
func (_ForwardLib *ForwardLibFilterer) FilterForwardFailed(opts *bind.FilterOpts, fromChainID []*big.Int, finalizeIndex []*big.Int) (*ForwardLibForwardFailedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var finalizeIndexRule []interface{}
	for _, finalizeIndexItem := range finalizeIndex {
		finalizeIndexRule = append(finalizeIndexRule, finalizeIndexItem)
	}

	logs, sub, err := _ForwardLib.contract.FilterLogs(opts, "ForwardFailed", fromChainIDRule, finalizeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ForwardLibForwardFailedIterator{contract: _ForwardLib.contract, event: "ForwardFailed", logs: logs, sub: sub}, nil
}

// WatchForwardFailed is a free log subscription operation binding the contract event 0x606a6de1b36fbe9108fe6f544d5941e752ff8bd36b5a01f08068cf7943e7e3df.
//
// Solidity: event ForwardFailed(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint8 code, bytes32 reasonHash)
func (_ForwardLib *ForwardLibFilterer) WatchForwardFailed(opts *bind.WatchOpts, sink chan<- *ForwardLibForwardFailed, fromChainID []*big.Int, finalizeIndex []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var finalizeIndexRule []interface{}
	for _, finalizeIndexItem := range finalizeIndex {
		finalizeIndexRule = append(finalizeIndexRule, finalizeIndexItem)
	}

	logs, sub, err := _ForwardLib.contract.WatchLogs(opts, "ForwardFailed", fromChainIDRule, finalizeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwardLibForwardFailed)
				if err := _ForwardLib.contract.UnpackLog(event, "ForwardFailed", log); err != nil {
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

// ParseForwardFailed is a log parse operation binding the contract event 0x606a6de1b36fbe9108fe6f544d5941e752ff8bd36b5a01f08068cf7943e7e3df.
//
// Solidity: event ForwardFailed(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint8 code, bytes32 reasonHash)
func (_ForwardLib *ForwardLibFilterer) ParseForwardFailed(log types.Log) (*ForwardLibForwardFailed, error) {
	event := new(ForwardLibForwardFailed)
	if err := _ForwardLib.contract.UnpackLog(event, "ForwardFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwardLibForwardInitiatedIterator is returned from FilterForwardInitiated and is used to iterate over the raw logs and unpacked data for ForwardInitiated events raised by the ForwardLib contract.
type ForwardLibForwardInitiatedIterator struct {
	Event *ForwardLibForwardInitiated // Event containing the contract specifics and raw log

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
func (it *ForwardLibForwardInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwardLibForwardInitiated)
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
		it.Event = new(ForwardLibForwardInitiated)
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
func (it *ForwardLibForwardInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwardLibForwardInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwardLibForwardInitiated represents a ForwardInitiated event raised by the ForwardLib contract.
type ForwardLibForwardInitiated struct {
	FromChainID   *big.Int
	FinalizeIndex *big.Int
	ToChainID     *big.Int
	FromToken     common.Address
	To            common.Address
	Value         *big.Int
	NetworkFee    *big.Int
	ExFee         *big.Int
	ExtraDataHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterForwardInitiated is a free log retrieval operation binding the contract event 0xd9c6b693f2a3033b5f515be724c03b8f30368cc278e689735e5edb7c6aa9bc56.
//
// Solidity: event ForwardInitiated(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint256 indexed toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes32 extraDataHash)
func (_ForwardLib *ForwardLibFilterer) FilterForwardInitiated(opts *bind.FilterOpts, fromChainID []*big.Int, finalizeIndex []*big.Int, toChainID []*big.Int) (*ForwardLibForwardInitiatedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var finalizeIndexRule []interface{}
	for _, finalizeIndexItem := range finalizeIndex {
		finalizeIndexRule = append(finalizeIndexRule, finalizeIndexItem)
	}
	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}

	logs, sub, err := _ForwardLib.contract.FilterLogs(opts, "ForwardInitiated", fromChainIDRule, finalizeIndexRule, toChainIDRule)
	if err != nil {
		return nil, err
	}
	return &ForwardLibForwardInitiatedIterator{contract: _ForwardLib.contract, event: "ForwardInitiated", logs: logs, sub: sub}, nil
}

// WatchForwardInitiated is a free log subscription operation binding the contract event 0xd9c6b693f2a3033b5f515be724c03b8f30368cc278e689735e5edb7c6aa9bc56.
//
// Solidity: event ForwardInitiated(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint256 indexed toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes32 extraDataHash)
func (_ForwardLib *ForwardLibFilterer) WatchForwardInitiated(opts *bind.WatchOpts, sink chan<- *ForwardLibForwardInitiated, fromChainID []*big.Int, finalizeIndex []*big.Int, toChainID []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var finalizeIndexRule []interface{}
	for _, finalizeIndexItem := range finalizeIndex {
		finalizeIndexRule = append(finalizeIndexRule, finalizeIndexItem)
	}
	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}

	logs, sub, err := _ForwardLib.contract.WatchLogs(opts, "ForwardInitiated", fromChainIDRule, finalizeIndexRule, toChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwardLibForwardInitiated)
				if err := _ForwardLib.contract.UnpackLog(event, "ForwardInitiated", log); err != nil {
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

// ParseForwardInitiated is a log parse operation binding the contract event 0xd9c6b693f2a3033b5f515be724c03b8f30368cc278e689735e5edb7c6aa9bc56.
//
// Solidity: event ForwardInitiated(uint256 indexed fromChainID, uint256 indexed finalizeIndex, uint256 indexed toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes32 extraDataHash)
func (_ForwardLib *ForwardLibFilterer) ParseForwardInitiated(log types.Log) (*ForwardLibForwardInitiated, error) {
	event := new(ForwardLibForwardInitiated)
	if err := _ForwardLib.contract.UnpackLog(event, "ForwardInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
