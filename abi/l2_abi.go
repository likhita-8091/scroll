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

// L2ABIMetaData contains all meta data concerning the L2ABI contract.
var L2ABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"FinalizeDepositERC20\",\"inputs\":[{\"type\":\"address\",\"name\":\"l1Token\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"l2Token\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizeDepositETH\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDefaultERC20Gateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"defaultERC20Gateway\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetERC20Gateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"gateway\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetETHGateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"ethGateway\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawERC20\",\"inputs\":[{\"type\":\"address\",\"name\":\"l1Token\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"l2Token\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawETH\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ERC20Gateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"defaultERC20Gateway\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ethGateway\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"finalizeDepositERC20\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"finalizeDepositETH\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getERC20Gateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getL1ERC20Address\",\"inputs\":[{\"type\":\"address\",\"name\":\"_l2Address\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getL2ERC20Address\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ethGateway\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_defaultERC20Gateway\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setDefaultERC20Gateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"_defaultERC20Gateway\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setERC20Gateway\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"_tokens\",\"internalType\":\"address[]\"},{\"type\":\"address[]\",\"name\":\"_gateways\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setETHGateway\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ethGateway\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawERC20\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawERC20\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawERC20AndCall\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawETH\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawETH\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"withdrawETHAndCall\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"},{\"type\":\"uint256\",\"name\":\"_gasLimit\",\"internalType\":\"uint256\"}]}]",
}

// L2ABIABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ABIMetaData.ABI instead.
var L2ABIABI = L2ABIMetaData.ABI

// L2ABI is an auto generated Go binding around an Ethereum contract.
type L2ABI struct {
	L2ABICaller     // Read-only binding to the contract
	L2ABITransactor // Write-only binding to the contract
	L2ABIFilterer   // Log filterer for contract events
}

// L2ABICaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ABISession struct {
	Contract     *L2ABI            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ABICallerSession struct {
	Contract *L2ABICaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L2ABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ABITransactorSession struct {
	Contract     *L2ABITransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ABIRaw struct {
	Contract *L2ABI // Generic contract binding to access the raw methods on
}

// L2ABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ABICallerRaw struct {
	Contract *L2ABICaller // Generic read-only contract binding to access the raw methods on
}

// L2ABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ABITransactorRaw struct {
	Contract *L2ABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ABI creates a new instance of L2ABI, bound to a specific deployed contract.
func NewL2ABI(address common.Address, backend bind.ContractBackend) (*L2ABI, error) {
	contract, err := bindL2ABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ABI{L2ABICaller: L2ABICaller{contract: contract}, L2ABITransactor: L2ABITransactor{contract: contract}, L2ABIFilterer: L2ABIFilterer{contract: contract}}, nil
}

// NewL2ABICaller creates a new read-only instance of L2ABI, bound to a specific deployed contract.
func NewL2ABICaller(address common.Address, caller bind.ContractCaller) (*L2ABICaller, error) {
	contract, err := bindL2ABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ABICaller{contract: contract}, nil
}

// NewL2ABITransactor creates a new write-only instance of L2ABI, bound to a specific deployed contract.
func NewL2ABITransactor(address common.Address, transactor bind.ContractTransactor) (*L2ABITransactor, error) {
	contract, err := bindL2ABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ABITransactor{contract: contract}, nil
}

// NewL2ABIFilterer creates a new log filterer instance of L2ABI, bound to a specific deployed contract.
func NewL2ABIFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ABIFilterer, error) {
	contract, err := bindL2ABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ABIFilterer{contract: contract}, nil
}

// bindL2ABI binds a generic wrapper to an already deployed contract.
func bindL2ABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ABI *L2ABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ABI.Contract.L2ABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ABI *L2ABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ABI.Contract.L2ABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ABI *L2ABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ABI.Contract.L2ABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ABI *L2ABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ABI *L2ABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ABI *L2ABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ABI.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2ABI *L2ABICaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2ABI *L2ABISession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2ABI.Contract.ERC20Gateway(&_L2ABI.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2ABI *L2ABICallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2ABI.Contract.ERC20Gateway(&_L2ABI.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2ABI *L2ABICaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2ABI *L2ABISession) DefaultERC20Gateway() (common.Address, error) {
	return _L2ABI.Contract.DefaultERC20Gateway(&_L2ABI.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2ABI *L2ABICallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2ABI.Contract.DefaultERC20Gateway(&_L2ABI.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2ABI *L2ABICaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2ABI *L2ABISession) EthGateway() (common.Address, error) {
	return _L2ABI.Contract.EthGateway(&_L2ABI.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2ABI *L2ABICallerSession) EthGateway() (common.Address, error) {
	return _L2ABI.Contract.EthGateway(&_L2ABI.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2ABI *L2ABICaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2ABI *L2ABISession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetERC20Gateway(&_L2ABI.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2ABI *L2ABICallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetERC20Gateway(&_L2ABI.CallOpts, _token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2ABI *L2ABICaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "getL1ERC20Address", _l2Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2ABI *L2ABISession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetL1ERC20Address(&_L2ABI.CallOpts, _l2Address)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2ABI *L2ABICallerSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetL1ERC20Address(&_L2ABI.CallOpts, _l2Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ABI *L2ABICaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ABI *L2ABISession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetL2ERC20Address(&_L2ABI.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2ABI *L2ABICallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2ABI.Contract.GetL2ERC20Address(&_L2ABI.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ABI *L2ABICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ABI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ABI *L2ABISession) Owner() (common.Address, error) {
	return _L2ABI.Contract.Owner(&_L2ABI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ABI *L2ABICallerSession) Owner() (common.Address, error) {
	return _L2ABI.Contract.Owner(&_L2ABI.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABITransactor) FinalizeDepositERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "finalizeDepositERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABISession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2ABI.Contract.FinalizeDepositERC20(&_L2ABI.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABITransactorSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2ABI.Contract.FinalizeDepositERC20(&_L2ABI.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABITransactor) FinalizeDepositETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "finalizeDepositETH", arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABISession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ABI.Contract.FinalizeDepositETH(&_L2ABI.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2ABI *L2ABITransactorSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2ABI.Contract.FinalizeDepositETH(&_L2ABI.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABITransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABISession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.Initialize(&_L2ABI.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABITransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.Initialize(&_L2ABI.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ABI *L2ABITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ABI *L2ABISession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ABI.Contract.RenounceOwnership(&_L2ABI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ABI *L2ABITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2ABI.Contract.RenounceOwnership(&_L2ABI.TransactOpts)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABITransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "setDefaultERC20Gateway", _defaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABISession) SetDefaultERC20Gateway(_defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetDefaultERC20Gateway(&_L2ABI.TransactOpts, _defaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _defaultERC20Gateway) returns()
func (_L2ABI *L2ABITransactorSession) SetDefaultERC20Gateway(_defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetDefaultERC20Gateway(&_L2ABI.TransactOpts, _defaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2ABI *L2ABITransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2ABI *L2ABISession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetERC20Gateway(&_L2ABI.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2ABI *L2ABITransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetERC20Gateway(&_L2ABI.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_L2ABI *L2ABITransactor) SetETHGateway(opts *bind.TransactOpts, _ethGateway common.Address) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "setETHGateway", _ethGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_L2ABI *L2ABISession) SetETHGateway(_ethGateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetETHGateway(&_L2ABI.TransactOpts, _ethGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _ethGateway) returns()
func (_L2ABI *L2ABITransactorSession) SetETHGateway(_ethGateway common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.SetETHGateway(&_L2ABI.TransactOpts, _ethGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ABI *L2ABITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ABI *L2ABISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.TransferOwnership(&_L2ABI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ABI *L2ABITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2ABI.Contract.TransferOwnership(&_L2ABI.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC20(&_L2ABI.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC20(&_L2ABI.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC200(&_L2ABI.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC200(&_L2ABI.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC20AndCall(&_L2ABI.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawERC20AndCall(&_L2ABI.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETH(&_L2ABI.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETH(&_L2ABI.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETH0(&_L2ABI.TransactOpts, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETH0(&_L2ABI.TransactOpts, _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABISession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETHAndCall(&_L2ABI.TransactOpts, _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ABI *L2ABITransactorSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ABI.Contract.WithdrawETHAndCall(&_L2ABI.TransactOpts, _to, _amount, _data, _gasLimit)
}

// L2ABIFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2ABI contract.
type L2ABIFinalizeDepositERC20Iterator struct {
	Event *L2ABIFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2ABIFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABIFinalizeDepositERC20)
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
		it.Event = new(L2ABIFinalizeDepositERC20)
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
func (it *L2ABIFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABIFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABIFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2ABI contract.
type L2ABIFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ABIFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ABIFinalizeDepositERC20Iterator{contract: _L2ABI.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2ABIFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABIFinalizeDepositERC20)
				if err := _L2ABI.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2ABIFinalizeDepositERC20, error) {
	event := new(L2ABIFinalizeDepositERC20)
	if err := _L2ABI.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABIFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the L2ABI contract.
type L2ABIFinalizeDepositETHIterator struct {
	Event *L2ABIFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *L2ABIFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABIFinalizeDepositETH)
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
		it.Event = new(L2ABIFinalizeDepositETH)
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
func (it *L2ABIFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABIFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABIFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2ABI contract.
type L2ABIFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2ABIFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2ABIFinalizeDepositETHIterator{contract: _L2ABI.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *L2ABIFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABIFinalizeDepositETH)
				if err := _L2ABI.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) ParseFinalizeDepositETH(log types.Log) (*L2ABIFinalizeDepositETH, error) {
	event := new(L2ABIFinalizeDepositETH)
	if err := _L2ABI.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2ABI contract.
type L2ABIOwnershipTransferredIterator struct {
	Event *L2ABIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2ABIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABIOwnershipTransferred)
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
		it.Event = new(L2ABIOwnershipTransferred)
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
func (it *L2ABIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABIOwnershipTransferred represents a OwnershipTransferred event raised by the L2ABI contract.
type L2ABIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ABI *L2ABIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2ABIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2ABIOwnershipTransferredIterator{contract: _L2ABI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2ABI *L2ABIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2ABIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABIOwnershipTransferred)
				if err := _L2ABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2ABI *L2ABIFilterer) ParseOwnershipTransferred(log types.Log) (*L2ABIOwnershipTransferred, error) {
	event := new(L2ABIOwnershipTransferred)
	if err := _L2ABI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABISetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the L2ABI contract.
type L2ABISetDefaultERC20GatewayIterator struct {
	Event *L2ABISetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2ABISetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABISetDefaultERC20Gateway)
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
		it.Event = new(L2ABISetDefaultERC20Gateway)
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
func (it *L2ABISetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABISetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABISetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the L2ABI contract.
type L2ABISetDefaultERC20Gateway struct {
	DefaultERC20Gateway common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x08338857eef8e29b906267c37965aff1fdcdb2c18d0f7b52de3b2eb71474d35c.
//
// Solidity: event SetDefaultERC20Gateway(address indexed defaultERC20Gateway)
func (_L2ABI *L2ABIFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, defaultERC20Gateway []common.Address) (*L2ABISetDefaultERC20GatewayIterator, error) {

	var defaultERC20GatewayRule []interface{}
	for _, defaultERC20GatewayItem := range defaultERC20Gateway {
		defaultERC20GatewayRule = append(defaultERC20GatewayRule, defaultERC20GatewayItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "SetDefaultERC20Gateway", defaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2ABISetDefaultERC20GatewayIterator{contract: _L2ABI.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x08338857eef8e29b906267c37965aff1fdcdb2c18d0f7b52de3b2eb71474d35c.
//
// Solidity: event SetDefaultERC20Gateway(address indexed defaultERC20Gateway)
func (_L2ABI *L2ABIFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2ABISetDefaultERC20Gateway, defaultERC20Gateway []common.Address) (event.Subscription, error) {

	var defaultERC20GatewayRule []interface{}
	for _, defaultERC20GatewayItem := range defaultERC20Gateway {
		defaultERC20GatewayRule = append(defaultERC20GatewayRule, defaultERC20GatewayItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "SetDefaultERC20Gateway", defaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABISetDefaultERC20Gateway)
				if err := _L2ABI.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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
func (_L2ABI *L2ABIFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*L2ABISetDefaultERC20Gateway, error) {
	event := new(L2ABISetDefaultERC20Gateway)
	if err := _L2ABI.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABISetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the L2ABI contract.
type L2ABISetERC20GatewayIterator struct {
	Event *L2ABISetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2ABISetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABISetERC20Gateway)
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
		it.Event = new(L2ABISetERC20Gateway)
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
func (it *L2ABISetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABISetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABISetERC20Gateway represents a SetERC20Gateway event raised by the L2ABI contract.
type L2ABISetERC20Gateway struct {
	Token   common.Address
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x5b0c89ecf574aa07194121c4f4dd1cfbaaf37d303c22522c9498a0aaf678668d.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed gateway)
func (_L2ABI *L2ABIFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, gateway []common.Address) (*L2ABISetERC20GatewayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2ABISetERC20GatewayIterator{contract: _L2ABI.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x5b0c89ecf574aa07194121c4f4dd1cfbaaf37d303c22522c9498a0aaf678668d.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed gateway)
func (_L2ABI *L2ABIFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2ABISetERC20Gateway, token []common.Address, gateway []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABISetERC20Gateway)
				if err := _L2ABI.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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
func (_L2ABI *L2ABIFilterer) ParseSetERC20Gateway(log types.Log) (*L2ABISetERC20Gateway, error) {
	event := new(L2ABISetERC20Gateway)
	if err := _L2ABI.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABISetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the L2ABI contract.
type L2ABISetETHGatewayIterator struct {
	Event *L2ABISetETHGateway // Event containing the contract specifics and raw log

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
func (it *L2ABISetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABISetETHGateway)
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
		it.Event = new(L2ABISetETHGateway)
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
func (it *L2ABISetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABISetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABISetETHGateway represents a SetETHGateway event raised by the L2ABI contract.
type L2ABISetETHGateway struct {
	EthGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xadf5b643c52781d8c742887e4a7a282a5bf2beb794ae3582fdf97e05df003018.
//
// Solidity: event SetETHGateway(address indexed ethGateway)
func (_L2ABI *L2ABIFilterer) FilterSetETHGateway(opts *bind.FilterOpts, ethGateway []common.Address) (*L2ABISetETHGatewayIterator, error) {

	var ethGatewayRule []interface{}
	for _, ethGatewayItem := range ethGateway {
		ethGatewayRule = append(ethGatewayRule, ethGatewayItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "SetETHGateway", ethGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2ABISetETHGatewayIterator{contract: _L2ABI.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xadf5b643c52781d8c742887e4a7a282a5bf2beb794ae3582fdf97e05df003018.
//
// Solidity: event SetETHGateway(address indexed ethGateway)
func (_L2ABI *L2ABIFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *L2ABISetETHGateway, ethGateway []common.Address) (event.Subscription, error) {

	var ethGatewayRule []interface{}
	for _, ethGatewayItem := range ethGateway {
		ethGatewayRule = append(ethGatewayRule, ethGatewayItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "SetETHGateway", ethGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABISetETHGateway)
				if err := _L2ABI.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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
func (_L2ABI *L2ABIFilterer) ParseSetETHGateway(log types.Log) (*L2ABISetETHGateway, error) {
	event := new(L2ABISetETHGateway)
	if err := _L2ABI.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABIWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2ABI contract.
type L2ABIWithdrawERC20Iterator struct {
	Event *L2ABIWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2ABIWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABIWithdrawERC20)
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
		it.Event = new(L2ABIWithdrawERC20)
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
func (it *L2ABIWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABIWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABIWithdrawERC20 represents a WithdrawERC20 event raised by the L2ABI contract.
type L2ABIWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2ABIWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ABIWithdrawERC20Iterator{contract: _L2ABI.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2ABIWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABIWithdrawERC20)
				if err := _L2ABI.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) ParseWithdrawERC20(log types.Log) (*L2ABIWithdrawERC20, error) {
	event := new(L2ABIWithdrawERC20)
	if err := _L2ABI.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ABIWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the L2ABI contract.
type L2ABIWithdrawETHIterator struct {
	Event *L2ABIWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L2ABIWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ABIWithdrawETH)
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
		it.Event = new(L2ABIWithdrawETH)
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
func (it *L2ABIWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ABIWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ABIWithdrawETH represents a WithdrawETH event raised by the L2ABI contract.
type L2ABIWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2ABIWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ABI.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2ABIWithdrawETHIterator{contract: _L2ABI.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *L2ABIWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ABI.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ABIWithdrawETH)
				if err := _L2ABI.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2ABI *L2ABIFilterer) ParseWithdrawETH(log types.Log) (*L2ABIWithdrawETH, error) {
	event := new(L2ABIWithdrawETH)
	if err := _L2ABI.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
