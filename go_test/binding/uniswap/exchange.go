// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"func\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTokens\",\"type\":\"uint256\"}],\"name\":\"ethToTokenSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenSold\",\"type\":\"uint256\"}],\"name\":\"getEthAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inputReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputReserve\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethSold\",\"type\":\"uint256\"}],\"name\":\"getTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"}],\"name\":\"tokenToEthSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Exchange *ExchangeCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Exchange *ExchangeSession) AllAdmins() ([]common.Address, error) {
	return _Exchange.Contract.AllAdmins(&_Exchange.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Exchange *ExchangeCallerSession) AllAdmins() ([]common.Address, error) {
	return _Exchange.Contract.AllAdmins(&_Exchange.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Exchange *ExchangeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Exchange *ExchangeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Exchange.Contract.Allowance(&_Exchange.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Exchange *ExchangeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Exchange.Contract.Allowance(&_Exchange.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Exchange *ExchangeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Exchange *ExchangeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Exchange.Contract.BalanceOf(&_Exchange.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Exchange *ExchangeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Exchange.Contract.BalanceOf(&_Exchange.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Exchange *ExchangeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Exchange *ExchangeSession) Decimals() (uint8, error) {
	return _Exchange.Contract.Decimals(&_Exchange.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Exchange *ExchangeCallerSession) Decimals() (uint8, error) {
	return _Exchange.Contract.Decimals(&_Exchange.CallOpts)
}

// GetEthAmount is a free data retrieval call binding the contract method 0x8fc7a25d.
//
// Solidity: function getEthAmount(uint256 _tokenSold) view returns(uint256)
func (_Exchange *ExchangeCaller) GetEthAmount(opts *bind.CallOpts, _tokenSold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getEthAmount", _tokenSold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthAmount is a free data retrieval call binding the contract method 0x8fc7a25d.
//
// Solidity: function getEthAmount(uint256 _tokenSold) view returns(uint256)
func (_Exchange *ExchangeSession) GetEthAmount(_tokenSold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthAmount(&_Exchange.CallOpts, _tokenSold)
}

// GetEthAmount is a free data retrieval call binding the contract method 0x8fc7a25d.
//
// Solidity: function getEthAmount(uint256 _tokenSold) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetEthAmount(_tokenSold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthAmount(&_Exchange.CallOpts, _tokenSold)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 inputReserve, uint256 outputReserve) pure returns(uint256)
func (_Exchange *ExchangeCaller) GetPrice(opts *bind.CallOpts, inputReserve *big.Int, outputReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getPrice", inputReserve, outputReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 inputReserve, uint256 outputReserve) pure returns(uint256)
func (_Exchange *ExchangeSession) GetPrice(inputReserve *big.Int, outputReserve *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPrice(&_Exchange.CallOpts, inputReserve, outputReserve)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 inputReserve, uint256 outputReserve) pure returns(uint256)
func (_Exchange *ExchangeCallerSession) GetPrice(inputReserve *big.Int, outputReserve *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPrice(&_Exchange.CallOpts, inputReserve, outputReserve)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Exchange *ExchangeCaller) GetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Exchange *ExchangeSession) GetReserve() (*big.Int, error) {
	return _Exchange.Contract.GetReserve(&_Exchange.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetReserve() (*big.Int, error) {
	return _Exchange.Contract.GetReserve(&_Exchange.CallOpts)
}

// GetTokenAmount is a free data retrieval call binding the contract method 0xc2507ac1.
//
// Solidity: function getTokenAmount(uint256 _ethSold) view returns(uint256)
func (_Exchange *ExchangeCaller) GetTokenAmount(opts *bind.CallOpts, _ethSold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getTokenAmount", _ethSold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmount is a free data retrieval call binding the contract method 0xc2507ac1.
//
// Solidity: function getTokenAmount(uint256 _ethSold) view returns(uint256)
func (_Exchange *ExchangeSession) GetTokenAmount(_ethSold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenAmount(&_Exchange.CallOpts, _ethSold)
}

// GetTokenAmount is a free data retrieval call binding the contract method 0xc2507ac1.
//
// Solidity: function getTokenAmount(uint256 _ethSold) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetTokenAmount(_ethSold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenAmount(&_Exchange.CallOpts, _ethSold)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Exchange *ExchangeCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Exchange *ExchangeSession) IsAdmin(account common.Address) (bool, error) {
	return _Exchange.Contract.IsAdmin(&_Exchange.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _Exchange.Contract.IsAdmin(&_Exchange.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Exchange *ExchangeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Exchange *ExchangeSession) Name() (string, error) {
	return _Exchange.Contract.Name(&_Exchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Exchange *ExchangeCallerSession) Name() (string, error) {
	return _Exchange.Contract.Name(&_Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Exchange *ExchangeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Exchange *ExchangeSession) Symbol() (string, error) {
	return _Exchange.Contract.Symbol(&_Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Exchange *ExchangeCallerSession) Symbol() (string, error) {
	return _Exchange.Contract.Symbol(&_Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Exchange *ExchangeCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Exchange *ExchangeSession) TokenAddress() (common.Address, error) {
	return _Exchange.Contract.TokenAddress(&_Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_Exchange *ExchangeCallerSession) TokenAddress() (common.Address, error) {
	return _Exchange.Contract.TokenAddress(&_Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Exchange *ExchangeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Exchange *ExchangeSession) TotalSupply() (*big.Int, error) {
	return _Exchange.Contract.TotalSupply(&_Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Exchange *ExchangeCallerSession) TotalSupply() (*big.Int, error) {
	return _Exchange.Contract.TotalSupply(&_Exchange.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Exchange *ExchangeTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Exchange *ExchangeSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAdmin(&_Exchange.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Exchange *ExchangeTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AddAdmin(&_Exchange.TransactOpts, account)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _tokenAmount) payable returns(uint256)
func (_Exchange *ExchangeTransactor) AddLiquidity(opts *bind.TransactOpts, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addLiquidity", _tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _tokenAmount) payable returns(uint256)
func (_Exchange *ExchangeSession) AddLiquidity(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.AddLiquidity(&_Exchange.TransactOpts, _tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _tokenAmount) payable returns(uint256)
func (_Exchange *ExchangeTransactorSession) AddLiquidity(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.AddLiquidity(&_Exchange.TransactOpts, _tokenAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Exchange *ExchangeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Approve(&_Exchange.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Approve(&_Exchange.TransactOpts, spender, value)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Exchange *ExchangeTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Exchange *ExchangeSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.BatchAddAdmin(&_Exchange.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Exchange *ExchangeTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.BatchAddAdmin(&_Exchange.TransactOpts, amounts)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0xf979dcd7.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens) payable returns()
func (_Exchange *ExchangeTransactor) EthToTokenSwap(opts *bind.TransactOpts, _minTokens *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "ethToTokenSwap", _minTokens)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0xf979dcd7.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens) payable returns()
func (_Exchange *ExchangeSession) EthToTokenSwap(_minTokens *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwap(&_Exchange.TransactOpts, _minTokens)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0xf979dcd7.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens) payable returns()
func (_Exchange *ExchangeTransactorSession) EthToTokenSwap(_minTokens *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwap(&_Exchange.TransactOpts, _minTokens)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Exchange *ExchangeTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Exchange *ExchangeSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAdmin(&_Exchange.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAdmin(&_Exchange.TransactOpts, account)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _amount) returns(uint256, uint256)
func (_Exchange *ExchangeTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeLiquidity", _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _amount) returns(uint256, uint256)
func (_Exchange *ExchangeSession) RemoveLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveLiquidity(&_Exchange.TransactOpts, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _amount) returns(uint256, uint256)
func (_Exchange *ExchangeTransactorSession) RemoveLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveLiquidity(&_Exchange.TransactOpts, _amount)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Exchange *ExchangeTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Exchange *ExchangeSession) RenounceAdmin() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceAdmin(&_Exchange.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Exchange *ExchangeTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceAdmin(&_Exchange.TransactOpts)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xe0486c8a.
//
// Solidity: function tokenToEthSwap(uint256 _tokensSold, uint256 _minEth) returns()
func (_Exchange *ExchangeTransactor) TokenToEthSwap(opts *bind.TransactOpts, _tokensSold *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToEthSwap", _tokensSold, _minEth)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xe0486c8a.
//
// Solidity: function tokenToEthSwap(uint256 _tokensSold, uint256 _minEth) returns()
func (_Exchange *ExchangeSession) TokenToEthSwap(_tokensSold *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwap(&_Exchange.TransactOpts, _tokensSold, _minEth)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xe0486c8a.
//
// Solidity: function tokenToEthSwap(uint256 _tokensSold, uint256 _minEth) returns()
func (_Exchange *ExchangeTransactorSession) TokenToEthSwap(_tokensSold *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwap(&_Exchange.TransactOpts, _tokensSold, _minEth)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Exchange *ExchangeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Transfer(&_Exchange.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Transfer(&_Exchange.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Exchange *ExchangeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TransferFrom(&_Exchange.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Exchange *ExchangeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TransferFrom(&_Exchange.TransactOpts, from, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// ExchangeAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Exchange contract.
type ExchangeAdminAddedIterator struct {
	Event *ExchangeAdminAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeAdminAdded)
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
		it.Event = new(ExchangeAdminAdded)
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
func (it *ExchangeAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeAdminAdded represents a AdminAdded event raised by the Exchange contract.
type ExchangeAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Exchange *ExchangeFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*ExchangeAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeAdminAddedIterator{contract: _Exchange.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Exchange *ExchangeFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *ExchangeAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeAdminAdded)
				if err := _Exchange.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Exchange *ExchangeFilterer) ParseAdminAdded(log types.Log) (*ExchangeAdminAdded, error) {
	event := new(ExchangeAdminAdded)
	if err := _Exchange.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Exchange contract.
type ExchangeAdminRemovedIterator struct {
	Event *ExchangeAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeAdminRemoved)
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
		it.Event = new(ExchangeAdminRemoved)
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
func (it *ExchangeAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeAdminRemoved represents a AdminRemoved event raised by the Exchange contract.
type ExchangeAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Exchange *ExchangeFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*ExchangeAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeAdminRemovedIterator{contract: _Exchange.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Exchange *ExchangeFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeAdminRemoved)
				if err := _Exchange.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Exchange *ExchangeFilterer) ParseAdminRemoved(log types.Log) (*ExchangeAdminRemoved, error) {
	event := new(ExchangeAdminRemoved)
	if err := _Exchange.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Exchange contract.
type ExchangeApprovalIterator struct {
	Event *ExchangeApproval // Event containing the contract specifics and raw log

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
func (it *ExchangeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeApproval)
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
		it.Event = new(ExchangeApproval)
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
func (it *ExchangeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeApproval represents a Approval event raised by the Exchange contract.
type ExchangeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Exchange *ExchangeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExchangeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeApprovalIterator{contract: _Exchange.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Exchange *ExchangeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExchangeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeApproval)
				if err := _Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseApproval(log types.Log) (*ExchangeApproval, error) {
	event := new(ExchangeApproval)
	if err := _Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Exchange contract.
type ExchangeLogIterator struct {
	Event *ExchangeLog // Event containing the contract specifics and raw log

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
func (it *ExchangeLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLog)
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
		it.Event = new(ExchangeLog)
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
func (it *ExchangeLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLog represents a Log event raised by the Exchange contract.
type ExchangeLog struct {
	Arg0   string
	Sender common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xf7f75251dee7d7fc22deac3247729ebe7c86541f35930bf10c2a4207479a3b6c.
//
// Solidity: event Log(string func, address sender, uint256 value, bytes data)
func (_Exchange *ExchangeFilterer) FilterLog(opts *bind.FilterOpts) (*ExchangeLogIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogIterator{contract: _Exchange.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xf7f75251dee7d7fc22deac3247729ebe7c86541f35930bf10c2a4207479a3b6c.
//
// Solidity: event Log(string func, address sender, uint256 value, bytes data)
func (_Exchange *ExchangeFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ExchangeLog) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLog)
				if err := _Exchange.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0xf7f75251dee7d7fc22deac3247729ebe7c86541f35930bf10c2a4207479a3b6c.
//
// Solidity: event Log(string func, address sender, uint256 value, bytes data)
func (_Exchange *ExchangeFilterer) ParseLog(log types.Log) (*ExchangeLog, error) {
	event := new(ExchangeLog)
	if err := _Exchange.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Exchange contract.
type ExchangeTransferIterator struct {
	Event *ExchangeTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTransfer)
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
		it.Event = new(ExchangeTransfer)
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
func (it *ExchangeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTransfer represents a Transfer event raised by the Exchange contract.
type ExchangeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Exchange *ExchangeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExchangeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransferIterator{contract: _Exchange.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Exchange *ExchangeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTransfer)
				if err := _Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseTransfer(log types.Log) (*ExchangeTransfer, error) {
	event := new(ExchangeTransfer)
	if err := _Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
