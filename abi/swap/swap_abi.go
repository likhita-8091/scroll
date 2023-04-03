// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// IApproveAndCallIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallIncreaseLiquidityParams struct {
	Token0     common.Address
	Token1     common.Address
	TokenId    *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
}

// IApproveAndCallMintParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallMintParams struct {
	Token0     common.Address
	Token1     common.Address
	Fee        *big.Int
	TickLower  *big.Int
	TickUpper  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Recipient  common.Address
}

// IV3SwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IV3SwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// SwapABIMetaData contains all meta data concerning the SwapABI contract.
var SwapABIMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_factoryV2\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"factoryV3\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_positionManager\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_WETH9\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"WETH9\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveMax\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveMaxMinusOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveZeroThenMax\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"approveZeroThenMaxMinusOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"callPositionManager\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[],\"name\":\"checkOracleSlippage\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"paths\",\"internalType\":\"bytes[]\"},{\"type\":\"uint128[]\",\"name\":\"amounts\",\"internalType\":\"uint128[]\"},{\"type\":\"uint24\",\"name\":\"maximumTickDivergence\",\"internalType\":\"uint24\"},{\"type\":\"uint32\",\"name\":\"secondsAgo\",\"internalType\":\"uint32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[],\"name\":\"checkOracleSlippage\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"uint24\",\"name\":\"maximumTickDivergence\",\"internalType\":\"uint24\"},{\"type\":\"uint32\",\"name\":\"secondsAgo\",\"internalType\":\"uint32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"exactInput\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"components\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMinimum\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"exactInputSingle\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"components\":[{\"type\":\"address\",\"name\":\"tokenIn\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenOut\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"internalType\":\"uint160\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"exactOutput\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"components\":[{\"type\":\"bytes\",\"name\":\"path\",\"internalType\":\"bytes\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMaximum\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"exactOutputSingle\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"components\":[{\"type\":\"address\",\"name\":\"tokenIn\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenOut\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMaximum\",\"internalType\":\"uint256\"},{\"type\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"internalType\":\"uint160\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factory\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factoryV2\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"enumIApproveAndCall.ApprovalType\"}],\"name\":\"getApprovalType\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"increaseLiquidity\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"components\":[{\"type\":\"address\",\"name\":\"token0\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"token1\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount0Min\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1Min\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"result\",\"internalType\":\"bytes\"}],\"name\":\"mint\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"params\",\"internalType\":\"structIApproveAndCall.MintParams\",\"components\":[{\"type\":\"address\",\"name\":\"token0\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"token1\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"int24\",\"name\":\"tickLower\",\"internalType\":\"int24\"},{\"type\":\"int24\",\"name\":\"tickUpper\",\"internalType\":\"int24\"},{\"type\":\"uint256\",\"name\":\"amount0Min\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount1Min\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"previousBlockhash\",\"internalType\":\"bytes32\"},{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"positionManager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"pull\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"refundETH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermit\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitAllowed\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"expiry\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitAllowedIfNecessary\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"nonce\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"expiry\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"selfPermitIfNecessary\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"v\",\"internalType\":\"uint8\"},{\"type\":\"bytes32\",\"name\":\"r\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"s\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountOutMin\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"path\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountIn\",\"internalType\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountOut\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountInMax\",\"internalType\":\"uint256\"},{\"type\":\"address[]\",\"name\":\"path\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepTokenWithFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepTokenWithFee\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"uniswapV3SwapCallback\",\"inputs\":[{\"type\":\"int256\",\"name\":\"amount0Delta\",\"internalType\":\"int256\"},{\"type\":\"int256\",\"name\":\"amount1Delta\",\"internalType\":\"int256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9WithFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9WithFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amountMinimum\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"feeBips\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeRecipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"wrapETH\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// SwapABIABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapABIMetaData.ABI instead.
var SwapABIABI = SwapABIMetaData.ABI

// SwapABI is an auto generated Go binding around an Ethereum contract.
type SwapABI struct {
	SwapABICaller     // Read-only binding to the contract
	SwapABITransactor // Write-only binding to the contract
	SwapABIFilterer   // Log filterer for contract events
}

// SwapABICaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapABISession struct {
	Contract     *SwapABI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapABICallerSession struct {
	Contract *SwapABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapABITransactorSession struct {
	Contract     *SwapABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapABIRaw struct {
	Contract *SwapABI // Generic contract binding to access the raw methods on
}

// SwapABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapABICallerRaw struct {
	Contract *SwapABICaller // Generic read-only contract binding to access the raw methods on
}

// SwapABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapABITransactorRaw struct {
	Contract *SwapABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapABI creates a new instance of SwapABI, bound to a specific deployed contract.
func NewSwapABI(address common.Address, backend bind.ContractBackend) (*SwapABI, error) {
	contract, err := bindSwapABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapABI{SwapABICaller: SwapABICaller{contract: contract}, SwapABITransactor: SwapABITransactor{contract: contract}, SwapABIFilterer: SwapABIFilterer{contract: contract}}, nil
}

// NewSwapABICaller creates a new read-only instance of SwapABI, bound to a specific deployed contract.
func NewSwapABICaller(address common.Address, caller bind.ContractCaller) (*SwapABICaller, error) {
	contract, err := bindSwapABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapABICaller{contract: contract}, nil
}

// NewSwapABITransactor creates a new write-only instance of SwapABI, bound to a specific deployed contract.
func NewSwapABITransactor(address common.Address, transactor bind.ContractTransactor) (*SwapABITransactor, error) {
	contract, err := bindSwapABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapABITransactor{contract: contract}, nil
}

// NewSwapABIFilterer creates a new log filterer instance of SwapABI, bound to a specific deployed contract.
func NewSwapABIFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapABIFilterer, error) {
	contract, err := bindSwapABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapABIFilterer{contract: contract}, nil
}

// bindSwapABI binds a generic wrapper to an already deployed contract.
func bindSwapABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapABI *SwapABIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapABI.Contract.SwapABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapABI *SwapABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapABI *SwapABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapABI *SwapABICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapABI *SwapABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapABI *SwapABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapABI.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapABI *SwapABICaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapABI *SwapABISession) WETH9() (common.Address, error) {
	return _SwapABI.Contract.WETH9(&_SwapABI.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SwapABI *SwapABICallerSession) WETH9() (common.Address, error) {
	return _SwapABI.Contract.WETH9(&_SwapABI.CallOpts)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABICaller) CheckOracleSlippage(opts *bind.CallOpts, paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABISession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SwapABI.Contract.CheckOracleSlippage(&_SwapABI.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABICallerSession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SwapABI.Contract.CheckOracleSlippage(&_SwapABI.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABICaller) CheckOracleSlippage0(opts *bind.CallOpts, path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABISession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SwapABI.Contract.CheckOracleSlippage0(&_SwapABI.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SwapABI *SwapABICallerSession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SwapABI.Contract.CheckOracleSlippage0(&_SwapABI.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapABI *SwapABICaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapABI *SwapABISession) Factory() (common.Address, error) {
	return _SwapABI.Contract.Factory(&_SwapABI.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapABI *SwapABICallerSession) Factory() (common.Address, error) {
	return _SwapABI.Contract.Factory(&_SwapABI.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SwapABI *SwapABICaller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SwapABI *SwapABISession) FactoryV2() (common.Address, error) {
	return _SwapABI.Contract.FactoryV2(&_SwapABI.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SwapABI *SwapABICallerSession) FactoryV2() (common.Address, error) {
	return _SwapABI.Contract.FactoryV2(&_SwapABI.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SwapABI *SwapABICaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapABI.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SwapABI *SwapABISession) PositionManager() (common.Address, error) {
	return _SwapABI.Contract.PositionManager(&_SwapABI.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SwapABI *SwapABICallerSession) PositionManager() (common.Address, error) {
	return _SwapABI.Contract.PositionManager(&_SwapABI.CallOpts)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SwapABI *SwapABITransactor) ApproveMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "approveMax", token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SwapABI *SwapABISession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveMax(&_SwapABI.TransactOpts, token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SwapABI *SwapABITransactorSession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveMax(&_SwapABI.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABITransactor) ApproveMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "approveMaxMinusOne", token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABISession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveMaxMinusOne(&_SwapABI.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABITransactorSession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveMaxMinusOne(&_SwapABI.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SwapABI *SwapABITransactor) ApproveZeroThenMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "approveZeroThenMax", token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SwapABI *SwapABISession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveZeroThenMax(&_SwapABI.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SwapABI *SwapABITransactorSession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveZeroThenMax(&_SwapABI.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABITransactor) ApproveZeroThenMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "approveZeroThenMaxMinusOne", token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABISession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveZeroThenMaxMinusOne(&_SwapABI.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SwapABI *SwapABITransactorSession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.ApproveZeroThenMaxMinusOne(&_SwapABI.TransactOpts, token)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SwapABI *SwapABITransactor) CallPositionManager(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "callPositionManager", data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SwapABI *SwapABISession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _SwapABI.Contract.CallPositionManager(&_SwapABI.TransactOpts, data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SwapABI *SwapABITransactorSession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _SwapABI.Contract.CallPositionManager(&_SwapABI.TransactOpts, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactor) ExactInput(opts *bind.TransactOpts, params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABISession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactInput(&_SwapABI.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactorSession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactInput(&_SwapABI.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactor) ExactInputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABISession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactInputSingle(&_SwapABI.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactorSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactInputSingle(&_SwapABI.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactor) ExactOutput(opts *bind.TransactOpts, params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABISession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactOutput(&_SwapABI.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactorSession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactOutput(&_SwapABI.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactor) ExactOutputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABISession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactOutputSingle(&_SwapABI.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactorSession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapABI.Contract.ExactOutputSingle(&_SwapABI.TransactOpts, params)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SwapABI *SwapABITransactor) GetApprovalType(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "getApprovalType", token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SwapABI *SwapABISession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.GetApprovalType(&_SwapABI.TransactOpts, token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SwapABI *SwapABITransactorSession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.GetApprovalType(&_SwapABI.TransactOpts, token, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SwapABI *SwapABITransactor) IncreaseLiquidity(opts *bind.TransactOpts, params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SwapABI *SwapABISession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SwapABI.Contract.IncreaseLiquidity(&_SwapABI.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SwapABI *SwapABITransactorSession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SwapABI.Contract.IncreaseLiquidity(&_SwapABI.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SwapABI *SwapABITransactor) Mint(opts *bind.TransactOpts, params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SwapABI *SwapABISession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SwapABI.Contract.Mint(&_SwapABI.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SwapABI *SwapABITransactorSession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SwapABI.Contract.Mint(&_SwapABI.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABITransactor) Multicall(opts *bind.TransactOpts, previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "multicall", previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABISession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall(&_SwapABI.TransactOpts, previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABITransactorSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall(&_SwapABI.TransactOpts, previousBlockhash, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABITransactor) Multicall0(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "multicall0", deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABISession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall0(&_SwapABI.TransactOpts, deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SwapABI *SwapABITransactorSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall0(&_SwapABI.TransactOpts, deadline, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapABI *SwapABITransactor) Multicall1(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "multicall1", data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapABI *SwapABISession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall1(&_SwapABI.TransactOpts, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapABI *SwapABITransactorSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.Multicall1(&_SwapABI.TransactOpts, data)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SwapABI *SwapABITransactor) Pull(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "pull", token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SwapABI *SwapABISession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.Pull(&_SwapABI.TransactOpts, token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SwapABI *SwapABITransactorSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.Pull(&_SwapABI.TransactOpts, token, value)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapABI *SwapABITransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapABI *SwapABISession) RefundETH() (*types.Transaction, error) {
	return _SwapABI.Contract.RefundETH(&_SwapABI.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapABI *SwapABITransactorSession) RefundETH() (*types.Transaction, error) {
	return _SwapABI.Contract.RefundETH(&_SwapABI.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABISession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermit(&_SwapABI.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermit(&_SwapABI.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABISession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitAllowed(&_SwapABI.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitAllowed(&_SwapABI.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABISession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitAllowedIfNecessary(&_SwapABI.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitAllowedIfNecessary(&_SwapABI.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABISession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitIfNecessary(&_SwapABI.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapABI *SwapABITransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapABI.Contract.SelfPermitIfNecessary(&_SwapABI.TransactOpts, token, value, deadline, v, r, s)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SwapABI *SwapABISession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapExactTokensForTokens(&_SwapABI.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SwapABI *SwapABITransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapExactTokensForTokens(&_SwapABI.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SwapABI *SwapABISession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapTokensForExactTokens(&_SwapABI.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SwapABI *SwapABITransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SwapTokensForExactTokens(&_SwapABI.TransactOpts, amountOut, amountInMax, path, to)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABITransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABISession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepToken(&_SwapABI.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABITransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepToken(&_SwapABI.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABITransactor) SweepToken0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "sweepToken0", token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABISession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepToken0(&_SwapABI.TransactOpts, token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABITransactorSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepToken0(&_SwapABI.TransactOpts, token, amountMinimum)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABISession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepTokenWithFee(&_SwapABI.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepTokenWithFee(&_SwapABI.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactor) SweepTokenWithFee0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABISession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepTokenWithFee0(&_SwapABI.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactorSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.SweepTokenWithFee0(&_SwapABI.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapABI *SwapABITransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapABI *SwapABISession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapABI.Contract.UniswapV3SwapCallback(&_SwapABI.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapABI *SwapABITransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapABI.Contract.UniswapV3SwapCallback(&_SwapABI.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABITransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABISession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9(&_SwapABI.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SwapABI *SwapABITransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9(&_SwapABI.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABITransactor) UnwrapWETH90(opts *bind.TransactOpts, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "unwrapWETH90", amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABISession) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH90(&_SwapABI.TransactOpts, amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_SwapABI *SwapABITransactorSession) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH90(&_SwapABI.TransactOpts, amountMinimum)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABISession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9WithFee(&_SwapABI.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9WithFee(&_SwapABI.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactor) UnwrapWETH9WithFee0(opts *bind.TransactOpts, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABISession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9WithFee0(&_SwapABI.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapABI *SwapABITransactorSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapABI.Contract.UnwrapWETH9WithFee0(&_SwapABI.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SwapABI *SwapABITransactor) WrapETH(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _SwapABI.contract.Transact(opts, "wrapETH", value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SwapABI *SwapABISession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.WrapETH(&_SwapABI.TransactOpts, value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SwapABI *SwapABITransactorSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _SwapABI.Contract.WrapETH(&_SwapABI.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapABI *SwapABITransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapABI.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapABI *SwapABISession) Receive() (*types.Transaction, error) {
	return _SwapABI.Contract.Receive(&_SwapABI.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapABI *SwapABITransactorSession) Receive() (*types.Transaction, error) {
	return _SwapABI.Contract.Receive(&_SwapABI.TransactOpts)
}
