// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package initialize

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

// InitializeMetaData contains all meta data concerning the Initialize contract.
var InitializeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountId\",\"type\":\"uint256\"}],\"name\":\"accountEvent\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"internalType\":\"contractAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"callAccountInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThisAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountId\",\"type\":\"uint256\"}],\"name\":\"newAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accountId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"newAccountWithEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// InitializeABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializeMetaData.ABI instead.
var InitializeABI = InitializeMetaData.ABI

// Initialize is an auto generated Go binding around an Ethereum contract.
type Initialize struct {
	InitializeCaller     // Read-only binding to the contract
	InitializeTransactor // Write-only binding to the contract
	InitializeFilterer   // Log filterer for contract events
}

// InitializeCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializeSession struct {
	Contract     *Initialize       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializeCallerSession struct {
	Contract *InitializeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InitializeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializeTransactorSession struct {
	Contract     *InitializeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InitializeRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializeRaw struct {
	Contract *Initialize // Generic contract binding to access the raw methods on
}

// InitializeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializeCallerRaw struct {
	Contract *InitializeCaller // Generic read-only contract binding to access the raw methods on
}

// InitializeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializeTransactorRaw struct {
	Contract *InitializeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitialize creates a new instance of Initialize, bound to a specific deployed contract.
func NewInitialize(address common.Address, backend bind.ContractBackend) (*Initialize, error) {
	contract, err := bindInitialize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initialize{InitializeCaller: InitializeCaller{contract: contract}, InitializeTransactor: InitializeTransactor{contract: contract}, InitializeFilterer: InitializeFilterer{contract: contract}}, nil
}

// NewInitializeCaller creates a new read-only instance of Initialize, bound to a specific deployed contract.
func NewInitializeCaller(address common.Address, caller bind.ContractCaller) (*InitializeCaller, error) {
	contract, err := bindInitialize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializeCaller{contract: contract}, nil
}

// NewInitializeTransactor creates a new write-only instance of Initialize, bound to a specific deployed contract.
func NewInitializeTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializeTransactor, error) {
	contract, err := bindInitialize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializeTransactor{contract: contract}, nil
}

// NewInitializeFilterer creates a new log filterer instance of Initialize, bound to a specific deployed contract.
func NewInitializeFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializeFilterer, error) {
	contract, err := bindInitialize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializeFilterer{contract: contract}, nil
}

// bindInitialize binds a generic wrapper to an already deployed contract.
func bindInitialize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initialize *InitializeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initialize.Contract.InitializeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initialize *InitializeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initialize.Contract.InitializeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initialize *InitializeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initialize.Contract.InitializeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initialize *InitializeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initialize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initialize *InitializeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initialize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initialize *InitializeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initialize.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_Initialize *InitializeCaller) Account(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Initialize.contract.Call(opts, &out, "account")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_Initialize *InitializeSession) Account() (common.Address, error) {
	return _Initialize.Contract.Account(&_Initialize.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(address)
func (_Initialize *InitializeCallerSession) Account() (common.Address, error) {
	return _Initialize.Contract.Account(&_Initialize.CallOpts)
}

// GetAccountId is a free data retrieval call binding the contract method 0x6d1dca0b.
//
// Solidity: function getAccountId() view returns(uint256)
func (_Initialize *InitializeCaller) GetAccountId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Initialize.contract.Call(opts, &out, "getAccountId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountId is a free data retrieval call binding the contract method 0x6d1dca0b.
//
// Solidity: function getAccountId() view returns(uint256)
func (_Initialize *InitializeSession) GetAccountId() (*big.Int, error) {
	return _Initialize.Contract.GetAccountId(&_Initialize.CallOpts)
}

// GetAccountId is a free data retrieval call binding the contract method 0x6d1dca0b.
//
// Solidity: function getAccountId() view returns(uint256)
func (_Initialize *InitializeCallerSession) GetAccountId() (*big.Int, error) {
	return _Initialize.Contract.GetAccountId(&_Initialize.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Initialize *InitializeCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Initialize.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Initialize *InitializeSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Initialize.Contract.GetBalance(&_Initialize.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Initialize *InitializeCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Initialize.Contract.GetBalance(&_Initialize.CallOpts, addr)
}

// GetThisAddr is a free data retrieval call binding the contract method 0xc39911cf.
//
// Solidity: function getThisAddr() view returns(address)
func (_Initialize *InitializeCaller) GetThisAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Initialize.contract.Call(opts, &out, "getThisAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetThisAddr is a free data retrieval call binding the contract method 0xc39911cf.
//
// Solidity: function getThisAddr() view returns(address)
func (_Initialize *InitializeSession) GetThisAddr() (common.Address, error) {
	return _Initialize.Contract.GetThisAddr(&_Initialize.CallOpts)
}

// GetThisAddr is a free data retrieval call binding the contract method 0xc39911cf.
//
// Solidity: function getThisAddr() view returns(address)
func (_Initialize *InitializeCallerSession) GetThisAddr() (common.Address, error) {
	return _Initialize.Contract.GetThisAddr(&_Initialize.CallOpts)
}

// CallAccountInfo is a paid mutator transaction binding the contract method 0xaca85530.
//
// Solidity: function callAccountInfo(address _contractAddress) payable returns()
func (_Initialize *InitializeTransactor) CallAccountInfo(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _Initialize.contract.Transact(opts, "callAccountInfo", _contractAddress)
}

// CallAccountInfo is a paid mutator transaction binding the contract method 0xaca85530.
//
// Solidity: function callAccountInfo(address _contractAddress) payable returns()
func (_Initialize *InitializeSession) CallAccountInfo(_contractAddress common.Address) (*types.Transaction, error) {
	return _Initialize.Contract.CallAccountInfo(&_Initialize.TransactOpts, _contractAddress)
}

// CallAccountInfo is a paid mutator transaction binding the contract method 0xaca85530.
//
// Solidity: function callAccountInfo(address _contractAddress) payable returns()
func (_Initialize *InitializeTransactorSession) CallAccountInfo(_contractAddress common.Address) (*types.Transaction, error) {
	return _Initialize.Contract.CallAccountInfo(&_Initialize.TransactOpts, _contractAddress)
}

// NewAccount is a paid mutator transaction binding the contract method 0xf003abfe.
//
// Solidity: function newAccount(uint256 accountId) returns()
func (_Initialize *InitializeTransactor) NewAccount(opts *bind.TransactOpts, accountId *big.Int) (*types.Transaction, error) {
	return _Initialize.contract.Transact(opts, "newAccount", accountId)
}

// NewAccount is a paid mutator transaction binding the contract method 0xf003abfe.
//
// Solidity: function newAccount(uint256 accountId) returns()
func (_Initialize *InitializeSession) NewAccount(accountId *big.Int) (*types.Transaction, error) {
	return _Initialize.Contract.NewAccount(&_Initialize.TransactOpts, accountId)
}

// NewAccount is a paid mutator transaction binding the contract method 0xf003abfe.
//
// Solidity: function newAccount(uint256 accountId) returns()
func (_Initialize *InitializeTransactorSession) NewAccount(accountId *big.Int) (*types.Transaction, error) {
	return _Initialize.Contract.NewAccount(&_Initialize.TransactOpts, accountId)
}

// NewAccountWithEther is a paid mutator transaction binding the contract method 0xa699a9b6.
//
// Solidity: function newAccountWithEther(uint256 accountId, uint256 amount) payable returns()
func (_Initialize *InitializeTransactor) NewAccountWithEther(opts *bind.TransactOpts, accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Initialize.contract.Transact(opts, "newAccountWithEther", accountId, amount)
}

// NewAccountWithEther is a paid mutator transaction binding the contract method 0xa699a9b6.
//
// Solidity: function newAccountWithEther(uint256 accountId, uint256 amount) payable returns()
func (_Initialize *InitializeSession) NewAccountWithEther(accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Initialize.Contract.NewAccountWithEther(&_Initialize.TransactOpts, accountId, amount)
}

// NewAccountWithEther is a paid mutator transaction binding the contract method 0xa699a9b6.
//
// Solidity: function newAccountWithEther(uint256 accountId, uint256 amount) payable returns()
func (_Initialize *InitializeTransactorSession) NewAccountWithEther(accountId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Initialize.Contract.NewAccountWithEther(&_Initialize.TransactOpts, accountId, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Initialize *InitializeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Initialize.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Initialize *InitializeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Initialize.Contract.Fallback(&_Initialize.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Initialize *InitializeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Initialize.Contract.Fallback(&_Initialize.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Initialize *InitializeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initialize.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Initialize *InitializeSession) Receive() (*types.Transaction, error) {
	return _Initialize.Contract.Receive(&_Initialize.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Initialize *InitializeTransactorSession) Receive() (*types.Transaction, error) {
	return _Initialize.Contract.Receive(&_Initialize.TransactOpts)
}

// InitializeAccountEventIterator is returned from FilterAccountEvent and is used to iterate over the raw logs and unpacked data for AccountEvent events raised by the Initialize contract.
type InitializeAccountEventIterator struct {
	Event *InitializeAccountEvent // Event containing the contract specifics and raw log

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
func (it *InitializeAccountEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializeAccountEvent)
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
		it.Event = new(InitializeAccountEvent)
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
func (it *InitializeAccountEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializeAccountEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializeAccountEvent represents a AccountEvent event raised by the Initialize contract.
type InitializeAccountEvent struct {
	AccountId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountEvent is a free log retrieval operation binding the contract event 0xf985ff11c4bb6f31dbf12ed9f397d5625b0b28aba0e5fde753950bec73b189c4.
//
// Solidity: event accountEvent(uint256 accountId)
func (_Initialize *InitializeFilterer) FilterAccountEvent(opts *bind.FilterOpts) (*InitializeAccountEventIterator, error) {

	logs, sub, err := _Initialize.contract.FilterLogs(opts, "accountEvent")
	if err != nil {
		return nil, err
	}
	return &InitializeAccountEventIterator{contract: _Initialize.contract, event: "accountEvent", logs: logs, sub: sub}, nil
}

// WatchAccountEvent is a free log subscription operation binding the contract event 0xf985ff11c4bb6f31dbf12ed9f397d5625b0b28aba0e5fde753950bec73b189c4.
//
// Solidity: event accountEvent(uint256 accountId)
func (_Initialize *InitializeFilterer) WatchAccountEvent(opts *bind.WatchOpts, sink chan<- *InitializeAccountEvent) (event.Subscription, error) {

	logs, sub, err := _Initialize.contract.WatchLogs(opts, "accountEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializeAccountEvent)
				if err := _Initialize.contract.UnpackLog(event, "accountEvent", log); err != nil {
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

// ParseAccountEvent is a log parse operation binding the contract event 0xf985ff11c4bb6f31dbf12ed9f397d5625b0b28aba0e5fde753950bec73b189c4.
//
// Solidity: event accountEvent(uint256 accountId)
func (_Initialize *InitializeFilterer) ParseAccountEvent(log types.Log) (*InitializeAccountEvent, error) {
	event := new(InitializeAccountEvent)
	if err := _Initialize.contract.UnpackLog(event, "accountEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
