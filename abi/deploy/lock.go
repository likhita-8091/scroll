// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deploy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// LockABIMetaData contains all meta data concerning the LockABI contract.
var LockABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"payable\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlockTime\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"when\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"addresspayable\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"unlockTime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[]}]",
	Bin: "0x60806040526040516102a53803806102a58339810160408190526100229161009b565b8042106100815760405162461bcd60e51b815260206004820152602360248201527f556e6c6f636b2074696d652073686f756c6420626520696e207468652066757460448201526275726560e81b606482015260840160405180910390fd5b600055600180546001600160a01b031916331790556100b4565b6000602082840312156100ad57600080fd5b5051919050565b6101e2806100c36000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063251c1aa3146100465780633ccfd60b146100625780638da5cb5b1461006c575b600080fd5b61004f60005481565b6040519081526020015b60405180910390f35b61006a610097565b005b60015461007f906001600160a01b031681565b6040516001600160a01b039091168152602001610059565b6000544210156100e75760405162461bcd60e51b8152602060048201526016602482015275165bdd4818d85b89dd081dda5d1a191c985dc81e595d60521b60448201526064015b60405180910390fd5b6001546001600160a01b031633146101385760405162461bcd60e51b81526020600482015260146024820152732cb7ba9030b932b713ba103a34329037bbb732b960611b60448201526064016100de565b604080514781524260208201527fbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93910160405180910390a16001546040516001600160a01b03909116904780156108fc02916000818181858888f193505050501580156101a9573d6000803e3d6000fd5b5056fea2646970667358221220b8a609dc05ff5065404cb77d883e8ca88cbd8dd1ac1fc633075c57da09f3f92864736f6c6343000809003300000000000000000000000000000000000000000000000000000000660f58a3",
}

// LockABIABI is the input ABI used to generate the binding from.
// Deprecated: Use LockABIMetaData.ABI instead.
var LockABIABI = LockABIMetaData.ABI

// LockABIBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LockABIMetaData.Bin instead.
var LockABIBin = LockABIMetaData.Bin

// DeployLockABI deploys a new Ethereum contract, binding an instance of LockABI to it.
func DeployLockABI(auth *bind.TransactOpts, backend bind.ContractBackend, _unlockTime *big.Int) (common.Address, *types.Transaction, *LockABI, error) {
	parsed, err := LockABIMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockABIBin), backend, _unlockTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockABI{LockABICaller: LockABICaller{contract: contract}, LockABITransactor: LockABITransactor{contract: contract}, LockABIFilterer: LockABIFilterer{contract: contract}}, nil
}

// LockABI is an auto generated Go binding around an Ethereum contract.
type LockABI struct {
	LockABICaller     // Read-only binding to the contract
	LockABITransactor // Write-only binding to the contract
	LockABIFilterer   // Log filterer for contract events
}

// LockABICaller is an auto generated read-only Go binding around an Ethereum contract.
type LockABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockABISession struct {
	Contract     *LockABI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockABICallerSession struct {
	Contract *LockABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LockABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockABITransactorSession struct {
	Contract     *LockABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LockABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockABIRaw struct {
	Contract *LockABI // Generic contract binding to access the raw methods on
}

// LockABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockABICallerRaw struct {
	Contract *LockABICaller // Generic read-only contract binding to access the raw methods on
}

// LockABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockABITransactorRaw struct {
	Contract *LockABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockABI creates a new instance of LockABI, bound to a specific deployed contract.
func NewLockABI(address common.Address, backend bind.ContractBackend) (*LockABI, error) {
	contract, err := bindLockABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockABI{LockABICaller: LockABICaller{contract: contract}, LockABITransactor: LockABITransactor{contract: contract}, LockABIFilterer: LockABIFilterer{contract: contract}}, nil
}

// NewLockABICaller creates a new read-only instance of LockABI, bound to a specific deployed contract.
func NewLockABICaller(address common.Address, caller bind.ContractCaller) (*LockABICaller, error) {
	contract, err := bindLockABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockABICaller{contract: contract}, nil
}

// NewLockABITransactor creates a new write-only instance of LockABI, bound to a specific deployed contract.
func NewLockABITransactor(address common.Address, transactor bind.ContractTransactor) (*LockABITransactor, error) {
	contract, err := bindLockABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockABITransactor{contract: contract}, nil
}

// NewLockABIFilterer creates a new log filterer instance of LockABI, bound to a specific deployed contract.
func NewLockABIFilterer(address common.Address, filterer bind.ContractFilterer) (*LockABIFilterer, error) {
	contract, err := bindLockABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockABIFilterer{contract: contract}, nil
}

// bindLockABI binds a generic wrapper to an already deployed contract.
func bindLockABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockABI *LockABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockABI.Contract.LockABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockABI *LockABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockABI.Contract.LockABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockABI *LockABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockABI.Contract.LockABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockABI *LockABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockABI *LockABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockABI *LockABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockABI.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockABI *LockABICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockABI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockABI *LockABISession) Owner() (common.Address, error) {
	return _LockABI.Contract.Owner(&_LockABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LockABI *LockABICallerSession) Owner() (common.Address, error) {
	return _LockABI.Contract.Owner(&_LockABI.CallOpts)
}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_LockABI *LockABICaller) UnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LockABI.contract.Call(opts, &out, "unlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_LockABI *LockABISession) UnlockTime() (*big.Int, error) {
	return _LockABI.Contract.UnlockTime(&_LockABI.CallOpts)
}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_LockABI *LockABICallerSession) UnlockTime() (*big.Int, error) {
	return _LockABI.Contract.UnlockTime(&_LockABI.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LockABI *LockABITransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockABI.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LockABI *LockABISession) Withdraw() (*types.Transaction, error) {
	return _LockABI.Contract.Withdraw(&_LockABI.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LockABI *LockABITransactorSession) Withdraw() (*types.Transaction, error) {
	return _LockABI.Contract.Withdraw(&_LockABI.TransactOpts)
}

// LockABIWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the LockABI contract.
type LockABIWithdrawalIterator struct {
	Event *LockABIWithdrawal // Event containing the contract specifics and raw log

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
func (it *LockABIWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockABIWithdrawal)
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
		it.Event = new(LockABIWithdrawal)
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
func (it *LockABIWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockABIWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockABIWithdrawal represents a Withdrawal event raised by the LockABI contract.
type LockABIWithdrawal struct {
	Amount *big.Int
	When   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_LockABI *LockABIFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*LockABIWithdrawalIterator, error) {

	logs, sub, err := _LockABI.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &LockABIWithdrawalIterator{contract: _LockABI.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_LockABI *LockABIFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *LockABIWithdrawal) (event.Subscription, error) {

	logs, sub, err := _LockABI.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockABIWithdrawal)
				if err := _LockABI.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (_LockABI *LockABIFilterer) ParseWithdrawal(log types.Log) (*LockABIWithdrawal, error) {
	event := new(LockABIWithdrawal)
	if err := _LockABI.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
