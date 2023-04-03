// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weth

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

// WEthABIMetaData contains all meta data concerning the WEthABI contract.
var WEthABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"src\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"guy\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"dst\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"src\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"dst\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"type\":\"address\",\"name\":\"src\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"guy\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"deposit\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"dst\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"src\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"dst\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"wad\",\"internalType\":\"uint256\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// WEthABIABI is the input ABI used to generate the binding from.
// Deprecated: Use WEthABIMetaData.ABI instead.
var WEthABIABI = WEthABIMetaData.ABI

// WEthABI is an auto generated Go binding around an Ethereum contract.
type WEthABI struct {
	WEthABICaller     // Read-only binding to the contract
	WEthABITransactor // Write-only binding to the contract
	WEthABIFilterer   // Log filterer for contract events
}

// WEthABICaller is an auto generated read-only Go binding around an Ethereum contract.
type WEthABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEthABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type WEthABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEthABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WEthABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WEthABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WEthABISession struct {
	Contract     *WEthABI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WEthABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WEthABICallerSession struct {
	Contract *WEthABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WEthABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WEthABITransactorSession struct {
	Contract     *WEthABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WEthABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type WEthABIRaw struct {
	Contract *WEthABI // Generic contract binding to access the raw methods on
}

// WEthABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WEthABICallerRaw struct {
	Contract *WEthABICaller // Generic read-only contract binding to access the raw methods on
}

// WEthABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WEthABITransactorRaw struct {
	Contract *WEthABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewWEthABI creates a new instance of WEthABI, bound to a specific deployed contract.
func NewWEthABI(address common.Address, backend bind.ContractBackend) (*WEthABI, error) {
	contract, err := bindWEthABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WEthABI{WEthABICaller: WEthABICaller{contract: contract}, WEthABITransactor: WEthABITransactor{contract: contract}, WEthABIFilterer: WEthABIFilterer{contract: contract}}, nil
}

// NewWEthABICaller creates a new read-only instance of WEthABI, bound to a specific deployed contract.
func NewWEthABICaller(address common.Address, caller bind.ContractCaller) (*WEthABICaller, error) {
	contract, err := bindWEthABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WEthABICaller{contract: contract}, nil
}

// NewWEthABITransactor creates a new write-only instance of WEthABI, bound to a specific deployed contract.
func NewWEthABITransactor(address common.Address, transactor bind.ContractTransactor) (*WEthABITransactor, error) {
	contract, err := bindWEthABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WEthABITransactor{contract: contract}, nil
}

// NewWEthABIFilterer creates a new log filterer instance of WEthABI, bound to a specific deployed contract.
func NewWEthABIFilterer(address common.Address, filterer bind.ContractFilterer) (*WEthABIFilterer, error) {
	contract, err := bindWEthABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WEthABIFilterer{contract: contract}, nil
}

// bindWEthABI binds a generic wrapper to an already deployed contract.
func bindWEthABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WEthABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEthABI *WEthABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEthABI.Contract.WEthABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEthABI *WEthABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEthABI.Contract.WEthABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEthABI *WEthABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEthABI.Contract.WEthABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WEthABI *WEthABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WEthABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WEthABI *WEthABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEthABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WEthABI *WEthABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WEthABI.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WEthABI *WEthABICaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WEthABI *WEthABISession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WEthABI.Contract.Allowance(&_WEthABI.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WEthABI *WEthABICallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WEthABI.Contract.Allowance(&_WEthABI.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WEthABI *WEthABICaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WEthABI *WEthABISession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WEthABI.Contract.BalanceOf(&_WEthABI.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WEthABI *WEthABICallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WEthABI.Contract.BalanceOf(&_WEthABI.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEthABI *WEthABICaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEthABI *WEthABISession) Decimals() (uint8, error) {
	return _WEthABI.Contract.Decimals(&_WEthABI.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WEthABI *WEthABICallerSession) Decimals() (uint8, error) {
	return _WEthABI.Contract.Decimals(&_WEthABI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEthABI *WEthABICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEthABI *WEthABISession) Name() (string, error) {
	return _WEthABI.Contract.Name(&_WEthABI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WEthABI *WEthABICallerSession) Name() (string, error) {
	return _WEthABI.Contract.Name(&_WEthABI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEthABI *WEthABICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEthABI *WEthABISession) Symbol() (string, error) {
	return _WEthABI.Contract.Symbol(&_WEthABI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WEthABI *WEthABICallerSession) Symbol() (string, error) {
	return _WEthABI.Contract.Symbol(&_WEthABI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEthABI *WEthABICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WEthABI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEthABI *WEthABISession) TotalSupply() (*big.Int, error) {
	return _WEthABI.Contract.TotalSupply(&_WEthABI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WEthABI *WEthABICallerSession) TotalSupply() (*big.Int, error) {
	return _WEthABI.Contract.TotalSupply(&_WEthABI.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WEthABI *WEthABISession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Approve(&_WEthABI.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Approve(&_WEthABI.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WEthABI *WEthABITransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEthABI.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WEthABI *WEthABISession) Deposit() (*types.Transaction, error) {
	return _WEthABI.Contract.Deposit(&_WEthABI.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WEthABI *WEthABITransactorSession) Deposit() (*types.Transaction, error) {
	return _WEthABI.Contract.Deposit(&_WEthABI.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABISession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Transfer(&_WEthABI.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Transfer(&_WEthABI.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABISession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.TransferFrom(&_WEthABI.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WEthABI *WEthABITransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.TransferFrom(&_WEthABI.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WEthABI *WEthABITransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WEthABI *WEthABISession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Withdraw(&_WEthABI.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WEthABI *WEthABITransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WEthABI.Contract.Withdraw(&_WEthABI.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEthABI *WEthABITransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WEthABI.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEthABI *WEthABISession) Receive() (*types.Transaction, error) {
	return _WEthABI.Contract.Receive(&_WEthABI.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WEthABI *WEthABITransactorSession) Receive() (*types.Transaction, error) {
	return _WEthABI.Contract.Receive(&_WEthABI.TransactOpts)
}

// WEthABIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WEthABI contract.
type WEthABIApprovalIterator struct {
	Event *WEthABIApproval // Event containing the contract specifics and raw log

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
func (it *WEthABIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEthABIApproval)
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
		it.Event = new(WEthABIApproval)
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
func (it *WEthABIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEthABIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEthABIApproval represents a Approval event raised by the WEthABI contract.
type WEthABIApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WEthABI *WEthABIFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WEthABIApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WEthABI.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WEthABIApprovalIterator{contract: _WEthABI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WEthABI *WEthABIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WEthABIApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WEthABI.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEthABIApproval)
				if err := _WEthABI.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WEthABI *WEthABIFilterer) ParseApproval(log types.Log) (*WEthABIApproval, error) {
	event := new(WEthABIApproval)
	if err := _WEthABI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEthABIDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WEthABI contract.
type WEthABIDepositIterator struct {
	Event *WEthABIDeposit // Event containing the contract specifics and raw log

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
func (it *WEthABIDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEthABIDeposit)
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
		it.Event = new(WEthABIDeposit)
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
func (it *WEthABIDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEthABIDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEthABIDeposit represents a Deposit event raised by the WEthABI contract.
type WEthABIDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WEthABIDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WEthABI.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WEthABIDepositIterator{contract: _WEthABI.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WEthABIDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WEthABI.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEthABIDeposit)
				if err := _WEthABI.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) ParseDeposit(log types.Log) (*WEthABIDeposit, error) {
	event := new(WEthABIDeposit)
	if err := _WEthABI.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEthABITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WEthABI contract.
type WEthABITransferIterator struct {
	Event *WEthABITransfer // Event containing the contract specifics and raw log

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
func (it *WEthABITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEthABITransfer)
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
		it.Event = new(WEthABITransfer)
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
func (it *WEthABITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEthABITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEthABITransfer represents a Transfer event raised by the WEthABI contract.
type WEthABITransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WEthABITransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WEthABI.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WEthABITransferIterator{contract: _WEthABI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WEthABITransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WEthABI.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEthABITransfer)
				if err := _WEthABI.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WEthABI *WEthABIFilterer) ParseTransfer(log types.Log) (*WEthABITransfer, error) {
	event := new(WEthABITransfer)
	if err := _WEthABI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WEthABIWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WEthABI contract.
type WEthABIWithdrawalIterator struct {
	Event *WEthABIWithdrawal // Event containing the contract specifics and raw log

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
func (it *WEthABIWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WEthABIWithdrawal)
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
		it.Event = new(WEthABIWithdrawal)
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
func (it *WEthABIWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WEthABIWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WEthABIWithdrawal represents a Withdrawal event raised by the WEthABI contract.
type WEthABIWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WEthABI *WEthABIFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WEthABIWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WEthABI.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WEthABIWithdrawalIterator{contract: _WEthABI.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WEthABI *WEthABIFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WEthABIWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WEthABI.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WEthABIWithdrawal)
				if err := _WEthABI.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WEthABI *WEthABIFilterer) ParseWithdrawal(log types.Log) (*WEthABIWithdrawal, error) {
	event := new(WEthABIWithdrawal)
	if err := _WEthABI.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
