// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ScrollABIMetaData contains all meta data concerning the ScrollABI contract.
var ScrollABIMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"SetDefaultERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"SetERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethGateway\",\"type\":\"address\"}],\"name\":\"SetETHGateway\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Address\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"setDefaultERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateways\",\"type\":\"address[]\"}],\"name\":\"setERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethGateway\",\"type\":\"address\"}],\"name\":\"setETHGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ScrollABIABI is the input ABI used to generate the binding from.
// Deprecated: Use ScrollABIMetaData.ABI instead.
var ScrollABIABI = ScrollABIMetaData.ABI

// ScrollABI is an auto generated Go binding around an Ethereum contract.
type ScrollABI struct {
	ScrollABICaller     // Read-only binding to the contract
	ScrollABITransactor // Write-only binding to the contract
	ScrollABIFilterer   // Log filterer for contract events
}

// ScrollABICaller is an auto generated read-only Go binding around an Ethereum contract.
type ScrollABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScrollABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScrollABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScrollABISession struct {
	Contract     *ScrollABI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScrollABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScrollABICallerSession struct {
	Contract *ScrollABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ScrollABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScrollABITransactorSession struct {
	Contract     *ScrollABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ScrollABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScrollABIRaw struct {
	Contract *ScrollABI // Generic contract binding to access the raw methods on
}

// ScrollABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScrollABICallerRaw struct {
	Contract *ScrollABICaller // Generic read-only contract binding to access the raw methods on
}

// ScrollABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScrollABITransactorRaw struct {
	Contract *ScrollABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewScrollABI creates a new instance of ScrollABI, bound to a specific deployed contract.
func NewScrollABI(address common.Address, backend bind.ContractBackend) (*ScrollABI, error) {
	contract, err := bindScrollABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollABI{ScrollABICaller: ScrollABICaller{contract: contract}, ScrollABITransactor: ScrollABITransactor{contract: contract}, ScrollABIFilterer: ScrollABIFilterer{contract: contract}}, nil
}

// NewScrollABICaller creates a new read-only instance of ScrollABI, bound to a specific deployed contract.
func NewScrollABICaller(address common.Address, caller bind.ContractCaller) (*ScrollABICaller, error) {
	contract, err := bindScrollABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollABICaller{contract: contract}, nil
}

// NewScrollABITransactor creates a new write-only instance of ScrollABI, bound to a specific deployed contract.
func NewScrollABITransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollABITransactor, error) {
	contract, err := bindScrollABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollABITransactor{contract: contract}, nil
}

// NewScrollABIFilterer creates a new log filterer instance of ScrollABI, bound to a specific deployed contract.
func NewScrollABIFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollABIFilterer, error) {
	contract, err := bindScrollABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollABIFilterer{contract: contract}, nil
}

// bindScrollABI binds a generic wrapper to an already deployed contract.
func bindScrollABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScrollABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollABI *ScrollABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollABI.Contract.ScrollABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollABI *ScrollABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollABI.Contract.ScrollABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollABI *ScrollABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollABI.Contract.ScrollABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollABI *ScrollABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollABI *ScrollABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollABI *ScrollABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollABI.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_ScrollABI *ScrollABICaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_ScrollABI *ScrollABISession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _ScrollABI.Contract.ERC20Gateway(&_ScrollABI.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_ScrollABI *ScrollABICallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _ScrollABI.Contract.ERC20Gateway(&_ScrollABI.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_ScrollABI *ScrollABICaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_ScrollABI *ScrollABISession) DefaultERC20Gateway() (common.Address, error) {
	return _ScrollABI.Contract.DefaultERC20Gateway(&_ScrollABI.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_ScrollABI *ScrollABICallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _ScrollABI.Contract.DefaultERC20Gateway(&_ScrollABI.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_ScrollABI *ScrollABICaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_ScrollABI *ScrollABISession) EthGateway() (common.Address, error) {
	return _ScrollABI.Contract.EthGateway(&_ScrollABI.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_ScrollABI *ScrollABICallerSession) EthGateway() (common.Address, error) {
	return _ScrollABI.Contract.EthGateway(&_ScrollABI.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_ScrollABI *ScrollABICaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_ScrollABI *ScrollABISession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _ScrollABI.Contract.GetERC20Gateway(&_ScrollABI.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_ScrollABI *ScrollABICallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _ScrollABI.Contract.GetERC20Gateway(&_ScrollABI.CallOpts, _token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_ScrollABI *ScrollABICaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "getL2ERC20Address", _l1Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_ScrollABI *ScrollABISession) GetL2ERC20Address(_l1Address common.Address) (common.Address, error) {
	return _ScrollABI.Contract.GetL2ERC20Address(&_ScrollABI.CallOpts, _l1Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Address) view returns(address)
func (_ScrollABI *ScrollABICallerSession) GetL2ERC20Address(_l1Address common.Address) (common.Address, error) {
	return _ScrollABI.Contract.GetL2ERC20Address(&_ScrollABI.CallOpts, _l1Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollABI *ScrollABICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollABI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollABI *ScrollABISession) Owner() (common.Address, error) {
	return _ScrollABI.Contract.Owner(&_ScrollABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollABI *ScrollABICallerSession) Owner() (common.Address, error) {
	return _ScrollABI.Contract.Owner(&_ScrollABI.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC20(&_ScrollABI.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC20(&_ScrollABI.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC200(&_ScrollABI.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC200(&_ScrollABI.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC20AndCall(&_ScrollABI.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositERC20AndCall(&_ScrollABI.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositETH(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositETH", _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETH(&_ScrollABI.TransactOpts, _amount, _gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositETH(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETH(&_ScrollABI.TransactOpts, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositETH0(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositETH0", _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETH0(&_ScrollABI.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositETH0(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETH0(&_ScrollABI.TransactOpts, _to, _amount, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactor) DepositETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "depositETHAndCall", _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABISession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETHAndCall(&_ScrollABI.TransactOpts, _to, _amount, _data, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_ScrollABI *ScrollABITransactorSession) DepositETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollABI.Contract.DepositETHAndCall(&_ScrollABI.TransactOpts, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABITransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABISession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _ScrollABI.Contract.FinalizeWithdrawERC20(&_ScrollABI.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABITransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _ScrollABI.Contract.FinalizeWithdrawERC20(&_ScrollABI.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABITransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "finalizeWithdrawETH", arg0, arg1, arg2, arg3)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABISession) FinalizeWithdrawETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ScrollABI.Contract.FinalizeWithdrawETH(&_ScrollABI.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address , address , uint256 , bytes ) payable returns()
func (_ScrollABI *ScrollABITransactorSession) FinalizeWithdrawETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ScrollABI.Contract.FinalizeWithdrawETH(&_ScrollABI.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABITransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABISession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.Initialize(&_ScrollABI.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABITransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.Initialize(&_ScrollABI.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollABI *ScrollABITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollABI *ScrollABISession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollABI.Contract.RenounceOwnership(&_ScrollABI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollABI *ScrollABITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollABI.Contract.RenounceOwnership(&_ScrollABI.TransactOpts)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABITransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "setDefaultERC20Gateway", _defaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABISession) SetDefaultERC20Gateway(_defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetDefaultERC20Gateway(&_ScrollABI.TransactOpts, _defaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_ScrollABI *ScrollABITransactorSession) SetDefaultERC20Gateway(_defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetDefaultERC20Gateway(&_ScrollABI.TransactOpts, _defaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_ScrollABI *ScrollABITransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_ScrollABI *ScrollABISession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetERC20Gateway(&_ScrollABI.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_ScrollABI *ScrollABITransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetERC20Gateway(&_ScrollABI.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_ScrollABI *ScrollABITransactor) SetETHGateway(opts *bind.TransactOpts, _ethGateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "setETHGateway", _ethGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_ScrollABI *ScrollABISession) SetETHGateway(_ethGateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetETHGateway(&_ScrollABI.TransactOpts, _ethGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_ScrollABI *ScrollABITransactorSession) SetETHGateway(_ethGateway common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.SetETHGateway(&_ScrollABI.TransactOpts, _ethGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollABI *ScrollABITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ScrollABI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollABI *ScrollABISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.TransferOwnership(&_ScrollABI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollABI *ScrollABITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollABI.Contract.TransferOwnership(&_ScrollABI.TransactOpts, newOwner)
}

// ScrollABIDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the ScrollABI contract.
type ScrollABIDepositERC20Iterator struct {
	Event *ScrollABIDepositERC20 // Event containing the contract specifics and raw log

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
func (it *ScrollABIDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABIDepositERC20)
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
		it.Event = new(ScrollABIDepositERC20)
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
func (it *ScrollABIDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABIDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABIDepositERC20 represents a DepositERC20 event raised by the ScrollABI contract.
type ScrollABIDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*ScrollABIDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABIDepositERC20Iterator{contract: _ScrollABI.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *ScrollABIDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABIDepositERC20)
				if err := _ScrollABI.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) ParseDepositERC20(log types.Log) (*ScrollABIDepositERC20, error) {
	event := new(ScrollABIDepositERC20)
	if err := _ScrollABI.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABIDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the ScrollABI contract.
type ScrollABIDepositETHIterator struct {
	Event *ScrollABIDepositETH // Event containing the contract specifics and raw log

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
func (it *ScrollABIDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABIDepositETH)
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
		it.Event = new(ScrollABIDepositETH)
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
func (it *ScrollABIDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABIDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABIDepositETH represents a DepositETH event raised by the ScrollABI contract.
type ScrollABIDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollABIDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABIDepositETHIterator{contract: _ScrollABI.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *ScrollABIDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABIDepositETH)
				if err := _ScrollABI.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) ParseDepositETH(log types.Log) (*ScrollABIDepositETH, error) {
	event := new(ScrollABIDepositETH)
	if err := _ScrollABI.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABIFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the ScrollABI contract.
type ScrollABIFinalizeWithdrawERC20Iterator struct {
	Event *ScrollABIFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *ScrollABIFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABIFinalizeWithdrawERC20)
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
		it.Event = new(ScrollABIFinalizeWithdrawERC20)
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
func (it *ScrollABIFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABIFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABIFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the ScrollABI contract.
type ScrollABIFinalizeWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC20 is a free log retrieval operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*ScrollABIFinalizeWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABIFinalizeWithdrawERC20Iterator{contract: _ScrollABI.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *ScrollABIFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABIFinalizeWithdrawERC20)
				if err := _ScrollABI.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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

// ParseFinalizeWithdrawERC20 is a log parse operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*ScrollABIFinalizeWithdrawERC20, error) {
	event := new(ScrollABIFinalizeWithdrawERC20)
	if err := _ScrollABI.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABIFinalizeWithdrawETHIterator is returned from FilterFinalizeWithdrawETH and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawETH events raised by the ScrollABI contract.
type ScrollABIFinalizeWithdrawETHIterator struct {
	Event *ScrollABIFinalizeWithdrawETH // Event containing the contract specifics and raw log

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
func (it *ScrollABIFinalizeWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABIFinalizeWithdrawETH)
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
		it.Event = new(ScrollABIFinalizeWithdrawETH)
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
func (it *ScrollABIFinalizeWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABIFinalizeWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABIFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the ScrollABI contract.
type ScrollABIFinalizeWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawETH is a free log retrieval operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) FilterFinalizeWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollABIFinalizeWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABIFinalizeWithdrawETHIterator{contract: _ScrollABI.contract, event: "FinalizeWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawETH is a free log subscription operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) WatchFinalizeWithdrawETH(opts *bind.WatchOpts, sink chan<- *ScrollABIFinalizeWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABIFinalizeWithdrawETH)
				if err := _ScrollABI.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
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

// ParseFinalizeWithdrawETH is a log parse operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_ScrollABI *ScrollABIFilterer) ParseFinalizeWithdrawETH(log types.Log) (*ScrollABIFinalizeWithdrawETH, error) {
	event := new(ScrollABIFinalizeWithdrawETH)
	if err := _ScrollABI.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ScrollABI contract.
type ScrollABIOwnershipTransferredIterator struct {
	Event *ScrollABIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ScrollABIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABIOwnershipTransferred)
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
		it.Event = new(ScrollABIOwnershipTransferred)
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
func (it *ScrollABIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABIOwnershipTransferred represents a OwnershipTransferred event raised by the ScrollABI contract.
type ScrollABIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollABI *ScrollABIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ScrollABIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABIOwnershipTransferredIterator{contract: _ScrollABI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollABI *ScrollABIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollABIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABIOwnershipTransferred)
				if err := _ScrollABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ScrollABI *ScrollABIFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollABIOwnershipTransferred, error) {
	event := new(ScrollABIOwnershipTransferred)
	if err := _ScrollABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABISetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the ScrollABI contract.
type ScrollABISetDefaultERC20GatewayIterator struct {
	Event *ScrollABISetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *ScrollABISetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABISetDefaultERC20Gateway)
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
		it.Event = new(ScrollABISetDefaultERC20Gateway)
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
func (it *ScrollABISetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABISetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABISetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the ScrollABI contract.
type ScrollABISetDefaultERC20Gateway struct {
	DefaultERC20Gateway common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x08338857eef8e29b906267c37965aff1fdcdb2c18d0f7b52de3b2eb71474d35c.
//
// Solidity: event SetDefaultERC20Gateway(address indexed defaultERC20Gateway)
func (_ScrollABI *ScrollABIFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, defaultERC20Gateway []common.Address) (*ScrollABISetDefaultERC20GatewayIterator, error) {

	var defaultERC20GatewayRule []interface{}
	for _, defaultERC20GatewayItem := range defaultERC20Gateway {
		defaultERC20GatewayRule = append(defaultERC20GatewayRule, defaultERC20GatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "SetDefaultERC20Gateway", defaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABISetDefaultERC20GatewayIterator{contract: _ScrollABI.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x08338857eef8e29b906267c37965aff1fdcdb2c18d0f7b52de3b2eb71474d35c.
//
// Solidity: event SetDefaultERC20Gateway(address indexed defaultERC20Gateway)
func (_ScrollABI *ScrollABIFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *ScrollABISetDefaultERC20Gateway, defaultERC20Gateway []common.Address) (event.Subscription, error) {

	var defaultERC20GatewayRule []interface{}
	for _, defaultERC20GatewayItem := range defaultERC20Gateway {
		defaultERC20GatewayRule = append(defaultERC20GatewayRule, defaultERC20GatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "SetDefaultERC20Gateway", defaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABISetDefaultERC20Gateway)
				if err := _ScrollABI.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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

// ParseSetDefaultERC20Gateway is a log parse operation binding the contract event 0x08338857eef8e29b906267c37965aff1fdcdb2c18d0f7b52de3b2eb71474d35c.
//
// Solidity: event SetDefaultERC20Gateway(address indexed defaultERC20Gateway)
func (_ScrollABI *ScrollABIFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*ScrollABISetDefaultERC20Gateway, error) {
	event := new(ScrollABISetDefaultERC20Gateway)
	if err := _ScrollABI.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABISetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the ScrollABI contract.
type ScrollABISetERC20GatewayIterator struct {
	Event *ScrollABISetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *ScrollABISetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABISetERC20Gateway)
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
		it.Event = new(ScrollABISetERC20Gateway)
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
func (it *ScrollABISetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABISetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABISetERC20Gateway represents a SetERC20Gateway event raised by the ScrollABI contract.
type ScrollABISetERC20Gateway struct {
	Token   common.Address
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x5b0c89ecf574aa07194121c4f4dd1cfbaaf37d303c22522c9498a0aaf678668d.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed gateway)
func (_ScrollABI *ScrollABIFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, gateway []common.Address) (*ScrollABISetERC20GatewayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABISetERC20GatewayIterator{contract: _ScrollABI.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x5b0c89ecf574aa07194121c4f4dd1cfbaaf37d303c22522c9498a0aaf678668d.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed gateway)
func (_ScrollABI *ScrollABIFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *ScrollABISetERC20Gateway, token []common.Address, gateway []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABISetERC20Gateway)
				if err := _ScrollABI.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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

// ParseSetERC20Gateway is a log parse operation binding the contract event 0x5b0c89ecf574aa07194121c4f4dd1cfbaaf37d303c22522c9498a0aaf678668d.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed gateway)
func (_ScrollABI *ScrollABIFilterer) ParseSetERC20Gateway(log types.Log) (*ScrollABISetERC20Gateway, error) {
	event := new(ScrollABISetERC20Gateway)
	if err := _ScrollABI.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollABISetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the ScrollABI contract.
type ScrollABISetETHGatewayIterator struct {
	Event *ScrollABISetETHGateway // Event containing the contract specifics and raw log

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
func (it *ScrollABISetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollABISetETHGateway)
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
		it.Event = new(ScrollABISetETHGateway)
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
func (it *ScrollABISetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollABISetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollABISetETHGateway represents a SetETHGateway event raised by the ScrollABI contract.
type ScrollABISetETHGateway struct {
	EthGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xadf5b643c52781d8c742887e4a7a282a5bf2beb794ae3582fdf97e05df003018.
//
// Solidity: event SetETHGateway(address indexed ethGateway)
func (_ScrollABI *ScrollABIFilterer) FilterSetETHGateway(opts *bind.FilterOpts, ethGateway []common.Address) (*ScrollABISetETHGatewayIterator, error) {

	var ethGatewayRule []interface{}
	for _, ethGatewayItem := range ethGateway {
		ethGatewayRule = append(ethGatewayRule, ethGatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.FilterLogs(opts, "SetETHGateway", ethGatewayRule)
	if err != nil {
		return nil, err
	}
	return &ScrollABISetETHGatewayIterator{contract: _ScrollABI.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xadf5b643c52781d8c742887e4a7a282a5bf2beb794ae3582fdf97e05df003018.
//
// Solidity: event SetETHGateway(address indexed ethGateway)
func (_ScrollABI *ScrollABIFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *ScrollABISetETHGateway, ethGateway []common.Address) (event.Subscription, error) {

	var ethGatewayRule []interface{}
	for _, ethGatewayItem := range ethGateway {
		ethGatewayRule = append(ethGatewayRule, ethGatewayItem)
	}

	logs, sub, err := _ScrollABI.contract.WatchLogs(opts, "SetETHGateway", ethGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollABISetETHGateway)
				if err := _ScrollABI.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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

// ParseSetETHGateway is a log parse operation binding the contract event 0xadf5b643c52781d8c742887e4a7a282a5bf2beb794ae3582fdf97e05df003018.
//
// Solidity: event SetETHGateway(address indexed ethGateway)
func (_ScrollABI *ScrollABIFilterer) ParseSetETHGateway(log types.Log) (*ScrollABISetETHGateway, error) {
	event := new(ScrollABISetETHGateway)
	if err := _ScrollABI.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
