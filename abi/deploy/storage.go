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

// StorageABIMetaData contains all meta data concerning the StorageABI contract.
var StorageABIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea26469706673582212202d1e6a790360e1ea57d40e5b95d84472b35aeaeb43c4a9c381a5cb3f091b3f1c64736f6c63430008130033",
}

// StorageABIABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageABIMetaData.ABI instead.
var StorageABIABI = StorageABIMetaData.ABI

// StorageABIBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageABIMetaData.Bin instead.
var StorageABIBin = StorageABIMetaData.Bin

// DeployStorageABI deploys a new Ethereum contract, binding an instance of StorageABI to it.
func DeployStorageABI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageABI, error) {
	parsed, err := StorageABIMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageABIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageABI{StorageABICaller: StorageABICaller{contract: contract}, StorageABITransactor: StorageABITransactor{contract: contract}, StorageABIFilterer: StorageABIFilterer{contract: contract}}, nil
}

// StorageABI is an auto generated Go binding around an Ethereum contract.
type StorageABI struct {
	StorageABICaller     // Read-only binding to the contract
	StorageABITransactor // Write-only binding to the contract
	StorageABIFilterer   // Log filterer for contract events
}

// StorageABICaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageABISession struct {
	Contract     *StorageABI       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageABICallerSession struct {
	Contract *StorageABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StorageABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageABITransactorSession struct {
	Contract     *StorageABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StorageABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageABIRaw struct {
	Contract *StorageABI // Generic contract binding to access the raw methods on
}

// StorageABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageABICallerRaw struct {
	Contract *StorageABICaller // Generic read-only contract binding to access the raw methods on
}

// StorageABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageABITransactorRaw struct {
	Contract *StorageABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageABI creates a new instance of StorageABI, bound to a specific deployed contract.
func NewStorageABI(address common.Address, backend bind.ContractBackend) (*StorageABI, error) {
	contract, err := bindStorageABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageABI{StorageABICaller: StorageABICaller{contract: contract}, StorageABITransactor: StorageABITransactor{contract: contract}, StorageABIFilterer: StorageABIFilterer{contract: contract}}, nil
}

// NewStorageABICaller creates a new read-only instance of StorageABI, bound to a specific deployed contract.
func NewStorageABICaller(address common.Address, caller bind.ContractCaller) (*StorageABICaller, error) {
	contract, err := bindStorageABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageABICaller{contract: contract}, nil
}

// NewStorageABITransactor creates a new write-only instance of StorageABI, bound to a specific deployed contract.
func NewStorageABITransactor(address common.Address, transactor bind.ContractTransactor) (*StorageABITransactor, error) {
	contract, err := bindStorageABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageABITransactor{contract: contract}, nil
}

// NewStorageABIFilterer creates a new log filterer instance of StorageABI, bound to a specific deployed contract.
func NewStorageABIFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageABIFilterer, error) {
	contract, err := bindStorageABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageABIFilterer{contract: contract}, nil
}

// bindStorageABI binds a generic wrapper to an already deployed contract.
func bindStorageABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageABI *StorageABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageABI.Contract.StorageABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageABI *StorageABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageABI.Contract.StorageABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageABI *StorageABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageABI.Contract.StorageABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageABI *StorageABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageABI *StorageABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageABI *StorageABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageABI.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageABI *StorageABICaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageABI.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageABI *StorageABISession) Retrieve() (*big.Int, error) {
	return _StorageABI.Contract.Retrieve(&_StorageABI.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageABI *StorageABICallerSession) Retrieve() (*big.Int, error) {
	return _StorageABI.Contract.Retrieve(&_StorageABI.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageABI *StorageABITransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _StorageABI.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageABI *StorageABISession) Store(num *big.Int) (*types.Transaction, error) {
	return _StorageABI.Contract.Store(&_StorageABI.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageABI *StorageABITransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _StorageABI.Contract.Store(&_StorageABI.TransactOpts, num)
}
