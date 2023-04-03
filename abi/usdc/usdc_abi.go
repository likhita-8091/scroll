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

// USDCABIMetaData contains all meta data concerning the USDCABI contract.
var USDCABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"claim\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"claimAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"received\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setClaimAmount\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"usdc\",\"inputs\":[]}]",
}

// USDCABIABI is the input ABI used to generate the binding from.
// Deprecated: Use USDCABIMetaData.ABI instead.
var USDCABIABI = USDCABIMetaData.ABI

// USDCABI is an auto generated Go binding around an Ethereum contract.
type USDCABI struct {
	USDCABICaller     // Read-only binding to the contract
	USDCABITransactor // Write-only binding to the contract
	USDCABIFilterer   // Log filterer for contract events
}

// USDCABICaller is an auto generated read-only Go binding around an Ethereum contract.
type USDCABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDCABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDCABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDCABISession struct {
	Contract     *USDCABI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDCABICallerSession struct {
	Contract *USDCABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// USDCABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDCABITransactorSession struct {
	Contract     *USDCABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// USDCABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDCABIRaw struct {
	Contract *USDCABI // Generic contract binding to access the raw methods on
}

// USDCABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDCABICallerRaw struct {
	Contract *USDCABICaller // Generic read-only contract binding to access the raw methods on
}

// USDCABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDCABITransactorRaw struct {
	Contract *USDCABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDCABI creates a new instance of USDCABI, bound to a specific deployed contract.
func NewUSDCABI(address common.Address, backend bind.ContractBackend) (*USDCABI, error) {
	contract, err := bindUSDCABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCABI{USDCABICaller: USDCABICaller{contract: contract}, USDCABITransactor: USDCABITransactor{contract: contract}, USDCABIFilterer: USDCABIFilterer{contract: contract}}, nil
}

// NewUSDCABICaller creates a new read-only instance of USDCABI, bound to a specific deployed contract.
func NewUSDCABICaller(address common.Address, caller bind.ContractCaller) (*USDCABICaller, error) {
	contract, err := bindUSDCABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCABICaller{contract: contract}, nil
}

// NewUSDCABITransactor creates a new write-only instance of USDCABI, bound to a specific deployed contract.
func NewUSDCABITransactor(address common.Address, transactor bind.ContractTransactor) (*USDCABITransactor, error) {
	contract, err := bindUSDCABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCABITransactor{contract: contract}, nil
}

// NewUSDCABIFilterer creates a new log filterer instance of USDCABI, bound to a specific deployed contract.
func NewUSDCABIFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCABIFilterer, error) {
	contract, err := bindUSDCABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCABIFilterer{contract: contract}, nil
}

// bindUSDCABI binds a generic wrapper to an already deployed contract.
func bindUSDCABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDCABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCABI *USDCABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCABI.Contract.USDCABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCABI *USDCABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCABI.Contract.USDCABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCABI *USDCABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCABI.Contract.USDCABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCABI *USDCABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCABI *USDCABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCABI *USDCABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCABI.Contract.contract.Transact(opts, method, params...)
}

// ClaimAmount is a free data retrieval call binding the contract method 0x830953ab.
//
// Solidity: function claimAmount() view returns(uint256)
func (_USDCABI *USDCABICaller) ClaimAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCABI.contract.Call(opts, &out, "claimAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimAmount is a free data retrieval call binding the contract method 0x830953ab.
//
// Solidity: function claimAmount() view returns(uint256)
func (_USDCABI *USDCABISession) ClaimAmount() (*big.Int, error) {
	return _USDCABI.Contract.ClaimAmount(&_USDCABI.CallOpts)
}

// ClaimAmount is a free data retrieval call binding the contract method 0x830953ab.
//
// Solidity: function claimAmount() view returns(uint256)
func (_USDCABI *USDCABICallerSession) ClaimAmount() (*big.Int, error) {
	return _USDCABI.Contract.ClaimAmount(&_USDCABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCABI *USDCABICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCABI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCABI *USDCABISession) Owner() (common.Address, error) {
	return _USDCABI.Contract.Owner(&_USDCABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCABI *USDCABICallerSession) Owner() (common.Address, error) {
	return _USDCABI.Contract.Owner(&_USDCABI.CallOpts)
}

// Received is a free data retrieval call binding the contract method 0xdf0cb934.
//
// Solidity: function received(address ) view returns(bool)
func (_USDCABI *USDCABICaller) Received(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _USDCABI.contract.Call(opts, &out, "received", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Received is a free data retrieval call binding the contract method 0xdf0cb934.
//
// Solidity: function received(address ) view returns(bool)
func (_USDCABI *USDCABISession) Received(arg0 common.Address) (bool, error) {
	return _USDCABI.Contract.Received(&_USDCABI.CallOpts, arg0)
}

// Received is a free data retrieval call binding the contract method 0xdf0cb934.
//
// Solidity: function received(address ) view returns(bool)
func (_USDCABI *USDCABICallerSession) Received(arg0 common.Address) (bool, error) {
	return _USDCABI.Contract.Received(&_USDCABI.CallOpts, arg0)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCABI *USDCABICaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCABI.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCABI *USDCABISession) Usdc() (common.Address, error) {
	return _USDCABI.Contract.Usdc(&_USDCABI.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCABI *USDCABICallerSession) Usdc() (common.Address, error) {
	return _USDCABI.Contract.Usdc(&_USDCABI.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_USDCABI *USDCABITransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCABI.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_USDCABI *USDCABISession) Claim() (*types.Transaction, error) {
	return _USDCABI.Contract.Claim(&_USDCABI.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_USDCABI *USDCABITransactorSession) Claim() (*types.Transaction, error) {
	return _USDCABI.Contract.Claim(&_USDCABI.TransactOpts)
}

// SetClaimAmount is a paid mutator transaction binding the contract method 0xb1c7ef0c.
//
// Solidity: function setClaimAmount(uint256 amount) returns()
func (_USDCABI *USDCABITransactor) SetClaimAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _USDCABI.contract.Transact(opts, "setClaimAmount", amount)
}

// SetClaimAmount is a paid mutator transaction binding the contract method 0xb1c7ef0c.
//
// Solidity: function setClaimAmount(uint256 amount) returns()
func (_USDCABI *USDCABISession) SetClaimAmount(amount *big.Int) (*types.Transaction, error) {
	return _USDCABI.Contract.SetClaimAmount(&_USDCABI.TransactOpts, amount)
}

// SetClaimAmount is a paid mutator transaction binding the contract method 0xb1c7ef0c.
//
// Solidity: function setClaimAmount(uint256 amount) returns()
func (_USDCABI *USDCABITransactorSession) SetClaimAmount(amount *big.Int) (*types.Transaction, error) {
	return _USDCABI.Contract.SetClaimAmount(&_USDCABI.TransactOpts, amount)
}
