// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc

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

// TestUsdcABIMetaData contains all meta data concerning the TestUsdcABI contract.
var TestUsdcABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"decreaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"increaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"addedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"nonces\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"permit\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]}]",
}

// TestUsdcABIABI is the input ABI used to generate the binding from.
// Deprecated: Use TestUsdcABIMetaData.ABI instead.
var TestUsdcABIABI = TestUsdcABIMetaData.ABI

// TestUsdcABI is an auto generated Go binding around an Ethereum contract.
type TestUsdcABI struct {
	TestUsdcABICaller     // Read-only binding to the contract
	TestUsdcABITransactor // Write-only binding to the contract
	TestUsdcABIFilterer   // Log filterer for contract events
}

// TestUsdcABICaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUsdcABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUsdcABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUsdcABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUsdcABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUsdcABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUsdcABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUsdcABISession struct {
	Contract     *TestUsdcABI      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestUsdcABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUsdcABICallerSession struct {
	Contract *TestUsdcABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestUsdcABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUsdcABITransactorSession struct {
	Contract     *TestUsdcABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestUsdcABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUsdcABIRaw struct {
	Contract *TestUsdcABI // Generic contract binding to access the raw methods on
}

// TestUsdcABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUsdcABICallerRaw struct {
	Contract *TestUsdcABICaller // Generic read-only contract binding to access the raw methods on
}

// TestUsdcABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUsdcABITransactorRaw struct {
	Contract *TestUsdcABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUsdcABI creates a new instance of TestUsdcABI, bound to a specific deployed contract.
func NewTestUsdcABI(address common.Address, backend bind.ContractBackend) (*TestUsdcABI, error) {
	contract, err := bindTestUsdcABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABI{TestUsdcABICaller: TestUsdcABICaller{contract: contract}, TestUsdcABITransactor: TestUsdcABITransactor{contract: contract}, TestUsdcABIFilterer: TestUsdcABIFilterer{contract: contract}}, nil
}

// NewTestUsdcABICaller creates a new read-only instance of TestUsdcABI, bound to a specific deployed contract.
func NewTestUsdcABICaller(address common.Address, caller bind.ContractCaller) (*TestUsdcABICaller, error) {
	contract, err := bindTestUsdcABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABICaller{contract: contract}, nil
}

// NewTestUsdcABITransactor creates a new write-only instance of TestUsdcABI, bound to a specific deployed contract.
func NewTestUsdcABITransactor(address common.Address, transactor bind.ContractTransactor) (*TestUsdcABITransactor, error) {
	contract, err := bindTestUsdcABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABITransactor{contract: contract}, nil
}

// NewTestUsdcABIFilterer creates a new log filterer instance of TestUsdcABI, bound to a specific deployed contract.
func NewTestUsdcABIFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUsdcABIFilterer, error) {
	contract, err := bindTestUsdcABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABIFilterer{contract: contract}, nil
}

// bindTestUsdcABI binds a generic wrapper to an already deployed contract.
func bindTestUsdcABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestUsdcABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUsdcABI *TestUsdcABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUsdcABI.Contract.TestUsdcABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUsdcABI *TestUsdcABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TestUsdcABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUsdcABI *TestUsdcABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TestUsdcABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUsdcABI *TestUsdcABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUsdcABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUsdcABI *TestUsdcABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUsdcABI *TestUsdcABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestUsdcABI *TestUsdcABICaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestUsdcABI *TestUsdcABISession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TestUsdcABI.Contract.DOMAINSEPARATOR(&_TestUsdcABI.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_TestUsdcABI *TestUsdcABICallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _TestUsdcABI.Contract.DOMAINSEPARATOR(&_TestUsdcABI.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestUsdcABI *TestUsdcABISession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.Allowance(&_TestUsdcABI.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.Allowance(&_TestUsdcABI.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestUsdcABI *TestUsdcABISession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.BalanceOf(&_TestUsdcABI.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.BalanceOf(&_TestUsdcABI.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestUsdcABI *TestUsdcABICaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestUsdcABI *TestUsdcABISession) Decimals() (uint8, error) {
	return _TestUsdcABI.Contract.Decimals(&_TestUsdcABI.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestUsdcABI *TestUsdcABICallerSession) Decimals() (uint8, error) {
	return _TestUsdcABI.Contract.Decimals(&_TestUsdcABI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestUsdcABI *TestUsdcABICaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestUsdcABI *TestUsdcABISession) Name() (string, error) {
	return _TestUsdcABI.Contract.Name(&_TestUsdcABI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestUsdcABI *TestUsdcABICallerSession) Name() (string, error) {
	return _TestUsdcABI.Contract.Name(&_TestUsdcABI.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestUsdcABI *TestUsdcABISession) Nonces(owner common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.Nonces(&_TestUsdcABI.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_TestUsdcABI *TestUsdcABICallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _TestUsdcABI.Contract.Nonces(&_TestUsdcABI.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestUsdcABI *TestUsdcABICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestUsdcABI *TestUsdcABISession) Owner() (common.Address, error) {
	return _TestUsdcABI.Contract.Owner(&_TestUsdcABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestUsdcABI *TestUsdcABICallerSession) Owner() (common.Address, error) {
	return _TestUsdcABI.Contract.Owner(&_TestUsdcABI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestUsdcABI *TestUsdcABICaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestUsdcABI *TestUsdcABISession) Symbol() (string, error) {
	return _TestUsdcABI.Contract.Symbol(&_TestUsdcABI.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestUsdcABI *TestUsdcABICallerSession) Symbol() (string, error) {
	return _TestUsdcABI.Contract.Symbol(&_TestUsdcABI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestUsdcABI *TestUsdcABICaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestUsdcABI.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestUsdcABI *TestUsdcABISession) TotalSupply() (*big.Int, error) {
	return _TestUsdcABI.Contract.TotalSupply(&_TestUsdcABI.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestUsdcABI *TestUsdcABICallerSession) TotalSupply() (*big.Int, error) {
	return _TestUsdcABI.Contract.TotalSupply(&_TestUsdcABI.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABISession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Approve(&_TestUsdcABI.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Approve(&_TestUsdcABI.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABISession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.DecreaseAllowance(&_TestUsdcABI.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.DecreaseAllowance(&_TestUsdcABI.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABISession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.IncreaseAllowance(&_TestUsdcABI.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.IncreaseAllowance(&_TestUsdcABI.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestUsdcABI *TestUsdcABITransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestUsdcABI *TestUsdcABISession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Permit(&_TestUsdcABI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TestUsdcABI *TestUsdcABITransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Permit(&_TestUsdcABI.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestUsdcABI *TestUsdcABITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestUsdcABI *TestUsdcABISession) RenounceOwnership() (*types.Transaction, error) {
	return _TestUsdcABI.Contract.RenounceOwnership(&_TestUsdcABI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestUsdcABI *TestUsdcABITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestUsdcABI.Contract.RenounceOwnership(&_TestUsdcABI.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABISession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Transfer(&_TestUsdcABI.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.Transfer(&_TestUsdcABI.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABISession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TransferFrom(&_TestUsdcABI.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestUsdcABI *TestUsdcABITransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TransferFrom(&_TestUsdcABI.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestUsdcABI *TestUsdcABITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestUsdcABI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestUsdcABI *TestUsdcABISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TransferOwnership(&_TestUsdcABI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestUsdcABI *TestUsdcABITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestUsdcABI.Contract.TransferOwnership(&_TestUsdcABI.TransactOpts, newOwner)
}

// TestUsdcABIApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestUsdcABI contract.
type TestUsdcABIApprovalIterator struct {
	Event *TestUsdcABIApproval // Event containing the contract specifics and raw log

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
func (it *TestUsdcABIApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUsdcABIApproval)
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
		it.Event = new(TestUsdcABIApproval)
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
func (it *TestUsdcABIApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUsdcABIApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUsdcABIApproval represents a Approval event raised by the TestUsdcABI contract.
type TestUsdcABIApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestUsdcABIApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestUsdcABI.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABIApprovalIterator{contract: _TestUsdcABI.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestUsdcABIApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestUsdcABI.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUsdcABIApproval)
				if err := _TestUsdcABI.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) ParseApproval(log types.Log) (*TestUsdcABIApproval, error) {
	event := new(TestUsdcABIApproval)
	if err := _TestUsdcABI.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUsdcABIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestUsdcABI contract.
type TestUsdcABIOwnershipTransferredIterator struct {
	Event *TestUsdcABIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestUsdcABIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUsdcABIOwnershipTransferred)
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
		it.Event = new(TestUsdcABIOwnershipTransferred)
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
func (it *TestUsdcABIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUsdcABIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUsdcABIOwnershipTransferred represents a OwnershipTransferred event raised by the TestUsdcABI contract.
type TestUsdcABIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestUsdcABI *TestUsdcABIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestUsdcABIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestUsdcABI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABIOwnershipTransferredIterator{contract: _TestUsdcABI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestUsdcABI *TestUsdcABIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestUsdcABIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestUsdcABI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUsdcABIOwnershipTransferred)
				if err := _TestUsdcABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestUsdcABI *TestUsdcABIFilterer) ParseOwnershipTransferred(log types.Log) (*TestUsdcABIOwnershipTransferred, error) {
	event := new(TestUsdcABIOwnershipTransferred)
	if err := _TestUsdcABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUsdcABITransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestUsdcABI contract.
type TestUsdcABITransferIterator struct {
	Event *TestUsdcABITransfer // Event containing the contract specifics and raw log

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
func (it *TestUsdcABITransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUsdcABITransfer)
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
		it.Event = new(TestUsdcABITransfer)
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
func (it *TestUsdcABITransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUsdcABITransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUsdcABITransfer represents a Transfer event raised by the TestUsdcABI contract.
type TestUsdcABITransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestUsdcABITransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestUsdcABI.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestUsdcABITransferIterator{contract: _TestUsdcABI.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestUsdcABITransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestUsdcABI.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUsdcABITransfer)
				if err := _TestUsdcABI.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestUsdcABI *TestUsdcABIFilterer) ParseTransfer(log types.Log) (*TestUsdcABITransfer, error) {
	event := new(TestUsdcABITransfer)
	if err := _TestUsdcABI.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
