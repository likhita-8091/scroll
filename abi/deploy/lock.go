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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"when\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040516105d83803806105d8833981810160405281019061002591906100f0565b804210610067576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161005e906101a0565b60405180910390fd5b8060008190555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506101c0565b600080fd5b6000819050919050565b6100cd816100ba565b81146100d857600080fd5b50565b6000815190506100ea816100c4565b92915050565b600060208284031215610106576101056100b5565b5b6000610114848285016100db565b91505092915050565b600082825260208201905092915050565b7f556e6c6f636b2074696d652073686f756c6420626520696e207468652066757460008201527f7572650000000000000000000000000000000000000000000000000000000000602082015250565b600061018a60238361011d565b91506101958261012e565b604082019050919050565b600060208201905081810360008301526101b98161017d565b9050919050565b610409806101cf6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063251c1aa3146100465780633ccfd60b146100645780638da5cb5b1461006e575b600080fd5b61004e61008c565b60405161005b919061024a565b60405180910390f35b61006c610092565b005b61007661020b565b60405161008391906102a6565b60405180910390f35b60005481565b6000544210156100d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ce9061031e565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015e9061038a565b60405180910390fd5b7fbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b9347426040516101989291906103aa565b60405180910390a1600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610208573d6000803e3d6000fd5b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000819050919050565b61024481610231565b82525050565b600060208201905061025f600083018461023b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061029082610265565b9050919050565b6102a081610285565b82525050565b60006020820190506102bb6000830184610297565b92915050565b600082825260208201905092915050565b7f596f752063616e27742077697468647261772079657400000000000000000000600082015250565b60006103086016836102c1565b9150610313826102d2565b602082019050919050565b60006020820190508181036000830152610337816102fb565b9050919050565b7f596f75206172656e277420746865206f776e6572000000000000000000000000600082015250565b60006103746014836102c1565b915061037f8261033e565b602082019050919050565b600060208201905081810360008301526103a381610367565b9050919050565b60006040820190506103bf600083018561023b565b6103cc602083018461023b565b939250505056fea26469706673582212204dd4a7a41dee611b487ff8b238c573552a81d92630114c4cef6b1c122840b9bc64736f6c63430008130033",
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
