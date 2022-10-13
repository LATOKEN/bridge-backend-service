// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package laHandler

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
)

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC20HandlerDepositRecord struct {
	TokenAddress                common.Address
	DestinationChainID          [8]byte
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	Amount                      *big.Int
}

// LaHandlerMetaData contains all meta data concerning the LaHandler contract.
var LaHandlerMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"_bridgeAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_resourceIDToTokenContractAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_contractWhitelist\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"_burnList\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"setResource\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setBurnable\",\"type\":\"function\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"_isInitialised\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"ownableInit\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"implementation\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"proxyOwner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_dexAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_WETH\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositRecords\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structERC20Handler.DepositRecord\",\"components\":[{\"name\":\"_tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_destinationRecipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"initialise\",\"type\":\"function\",\"inputs\":[{\"name\":\"initBridgeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initDEXAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initWETHAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminChangeBridgeAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBridgeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminChangeDEXAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newDEX\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminChangeWETHAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newWETH\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getDepositRecord\",\"type\":\"function\",\"inputs\":[{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destId\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structERC20Handler.DepositRecord\",\"components\":[{\"name\":\"_tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_destinationRecipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"deposit\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"executeProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"withdraw\",\"type\":\"function\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"depositFunds\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"approve\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferExtraLA\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false}]",
}

// LaHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use LaHandlerMetaData.ABI instead.
var LaHandlerABI = LaHandlerMetaData.ABI

// LaHandler is an auto generated Go binding around an Ethereum contract.
type LaHandler struct {
	LaHandlerCaller     // Read-only binding to the contract
	LaHandlerTransactor // Write-only binding to the contract
	LaHandlerFilterer   // Log filterer for contract events
}

// LaHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaHandlerSession struct {
	Contract     *LaHandler        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaHandlerCallerSession struct {
	Contract *LaHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LaHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaHandlerTransactorSession struct {
	Contract     *LaHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LaHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaHandlerRaw struct {
	Contract *LaHandler // Generic contract binding to access the raw methods on
}

// LaHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaHandlerCallerRaw struct {
	Contract *LaHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// LaHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaHandlerTransactorRaw struct {
	Contract *LaHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaHandler creates a new instance of LaHandler, bound to a specific deployed contract.
func NewLaHandler(address common.Address, backend bind.ContractBackend) (*LaHandler, error) {
	contract, err := bindLaHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LaHandler{LaHandlerCaller: LaHandlerCaller{contract: contract}, LaHandlerTransactor: LaHandlerTransactor{contract: contract}, LaHandlerFilterer: LaHandlerFilterer{contract: contract}}, nil
}

// NewLaHandlerCaller creates a new read-only instance of LaHandler, bound to a specific deployed contract.
func NewLaHandlerCaller(address common.Address, caller bind.ContractCaller) (*LaHandlerCaller, error) {
	contract, err := bindLaHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaHandlerCaller{contract: contract}, nil
}

// NewLaHandlerTransactor creates a new write-only instance of LaHandler, bound to a specific deployed contract.
func NewLaHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*LaHandlerTransactor, error) {
	contract, err := bindLaHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaHandlerTransactor{contract: contract}, nil
}

// NewLaHandlerFilterer creates a new log filterer instance of LaHandler, bound to a specific deployed contract.
func NewLaHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*LaHandlerFilterer, error) {
	contract, err := bindLaHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaHandlerFilterer{contract: contract}, nil
}

// bindLaHandler binds a generic wrapper to an already deployed contract.
func bindLaHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LaHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaHandler *LaHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaHandler.Contract.LaHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaHandler *LaHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaHandler.Contract.LaHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaHandler *LaHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaHandler.Contract.LaHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaHandler *LaHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaHandler *LaHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaHandler *LaHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaHandler.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_LaHandler *LaHandlerCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_LaHandler *LaHandlerSession) WETH() (common.Address, error) {
	return _LaHandler.Contract.WETH(&_LaHandler.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_LaHandler *LaHandlerCallerSession) WETH() (common.Address, error) {
	return _LaHandler.Contract.WETH(&_LaHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_LaHandler *LaHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_LaHandler *LaHandlerSession) BridgeAddress() (common.Address, error) {
	return _LaHandler.Contract.BridgeAddress(&_LaHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_LaHandler *LaHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _LaHandler.Contract.BridgeAddress(&_LaHandler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_LaHandler *LaHandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_LaHandler *LaHandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _LaHandler.Contract.BurnList(&_LaHandler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_LaHandler *LaHandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _LaHandler.Contract.BurnList(&_LaHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_LaHandler *LaHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_LaHandler *LaHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _LaHandler.Contract.ContractWhitelist(&_LaHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_LaHandler *LaHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _LaHandler.Contract.ContractWhitelist(&_LaHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 [8]byte, arg1 uint64) (ERC20HandlerDepositRecord, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new(ERC20HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20HandlerDepositRecord)).(*ERC20HandlerDepositRecord)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerSession) DepositRecords(arg0 [8]byte, arg1 uint64) (ERC20HandlerDepositRecord, error) {
	return _LaHandler.Contract.DepositRecords(&_LaHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerCallerSession) DepositRecords(arg0 [8]byte, arg1 uint64) (ERC20HandlerDepositRecord, error) {
	return _LaHandler.Contract.DepositRecords(&_LaHandler.CallOpts, arg0, arg1)
}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_LaHandler *LaHandlerCaller) DexAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_dexAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_LaHandler *LaHandlerSession) DexAddress() (common.Address, error) {
	return _LaHandler.Contract.DexAddress(&_LaHandler.CallOpts)
}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_LaHandler *LaHandlerCallerSession) DexAddress() (common.Address, error) {
	return _LaHandler.Contract.DexAddress(&_LaHandler.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaHandler *LaHandlerCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaHandler *LaHandlerSession) IsInitialised() (bool, error) {
	return _LaHandler.Contract.IsInitialised(&_LaHandler.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaHandler *LaHandlerCallerSession) IsInitialised() (bool, error) {
	return _LaHandler.Contract.IsInitialised(&_LaHandler.CallOpts)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_LaHandler *LaHandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_LaHandler *LaHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _LaHandler.Contract.ResourceIDToTokenContractAddress(&_LaHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_LaHandler *LaHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _LaHandler.Contract.ResourceIDToTokenContractAddress(&_LaHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC20HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20HandlerDepositRecord)).(*ERC20HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerSession) GetDepositRecord(depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	return _LaHandler.Contract.GetDepositRecord(&_LaHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_LaHandler *LaHandlerCallerSession) GetDepositRecord(depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	return _LaHandler.Contract.GetDepositRecord(&_LaHandler.CallOpts, depositNonce, destId)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaHandler *LaHandlerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaHandler *LaHandlerSession) Implementation() (common.Address, error) {
	return _LaHandler.Contract.Implementation(&_LaHandler.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaHandler *LaHandlerCallerSession) Implementation() (common.Address, error) {
	return _LaHandler.Contract.Implementation(&_LaHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaHandler *LaHandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaHandler *LaHandlerSession) Owner() (common.Address, error) {
	return _LaHandler.Contract.Owner(&_LaHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaHandler *LaHandlerCallerSession) Owner() (common.Address, error) {
	return _LaHandler.Contract.Owner(&_LaHandler.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaHandler *LaHandlerCaller) ProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaHandler.contract.Call(opts, &out, "proxyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaHandler *LaHandlerSession) ProxyOwner() (common.Address, error) {
	return _LaHandler.Contract.ProxyOwner(&_LaHandler.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaHandler *LaHandlerCallerSession) ProxyOwner() (common.Address, error) {
	return _LaHandler.Contract.ProxyOwner(&_LaHandler.CallOpts)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_LaHandler *LaHandlerTransactor) AdminChangeBridgeAddress(opts *bind.TransactOpts, newBridgeAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "adminChangeBridgeAddress", newBridgeAddress)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_LaHandler *LaHandlerSession) AdminChangeBridgeAddress(newBridgeAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeBridgeAddress(&_LaHandler.TransactOpts, newBridgeAddress)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_LaHandler *LaHandlerTransactorSession) AdminChangeBridgeAddress(newBridgeAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeBridgeAddress(&_LaHandler.TransactOpts, newBridgeAddress)
}

// AdminChangeDEXAddress is a paid mutator transaction binding the contract method 0x311e0c67.
//
// Solidity: function adminChangeDEXAddress(address newDEX) returns()
func (_LaHandler *LaHandlerTransactor) AdminChangeDEXAddress(opts *bind.TransactOpts, newDEX common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "adminChangeDEXAddress", newDEX)
}

// AdminChangeDEXAddress is a paid mutator transaction binding the contract method 0x311e0c67.
//
// Solidity: function adminChangeDEXAddress(address newDEX) returns()
func (_LaHandler *LaHandlerSession) AdminChangeDEXAddress(newDEX common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeDEXAddress(&_LaHandler.TransactOpts, newDEX)
}

// AdminChangeDEXAddress is a paid mutator transaction binding the contract method 0x311e0c67.
//
// Solidity: function adminChangeDEXAddress(address newDEX) returns()
func (_LaHandler *LaHandlerTransactorSession) AdminChangeDEXAddress(newDEX common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeDEXAddress(&_LaHandler.TransactOpts, newDEX)
}

// AdminChangeWETHAddress is a paid mutator transaction binding the contract method 0xf6eb24fa.
//
// Solidity: function adminChangeWETHAddress(address newWETH) returns()
func (_LaHandler *LaHandlerTransactor) AdminChangeWETHAddress(opts *bind.TransactOpts, newWETH common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "adminChangeWETHAddress", newWETH)
}

// AdminChangeWETHAddress is a paid mutator transaction binding the contract method 0xf6eb24fa.
//
// Solidity: function adminChangeWETHAddress(address newWETH) returns()
func (_LaHandler *LaHandlerSession) AdminChangeWETHAddress(newWETH common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeWETHAddress(&_LaHandler.TransactOpts, newWETH)
}

// AdminChangeWETHAddress is a paid mutator transaction binding the contract method 0xf6eb24fa.
//
// Solidity: function adminChangeWETHAddress(address newWETH) returns()
func (_LaHandler *LaHandlerTransactorSession) AdminChangeWETHAddress(newWETH common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.AdminChangeWETHAddress(&_LaHandler.TransactOpts, newWETH)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_LaHandler *LaHandlerTransactor) Approve(opts *bind.TransactOpts, resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "approve", resourceID, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_LaHandler *LaHandlerSession) Approve(resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.Contract.Approve(&_LaHandler.TransactOpts, resourceID, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_LaHandler *LaHandlerTransactorSession) Approve(resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.Contract.Approve(&_LaHandler.TransactOpts, resourceID, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_LaHandler *LaHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_LaHandler *LaHandlerSession) Deposit(resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.Deposit(&_LaHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_LaHandler *LaHandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.Deposit(&_LaHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaHandler *LaHandlerTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaHandler *LaHandlerSession) DepositFunds() (*types.Transaction, error) {
	return _LaHandler.Contract.DepositFunds(&_LaHandler.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaHandler *LaHandlerTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _LaHandler.Contract.DepositFunds(&_LaHandler.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_LaHandler *LaHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "executeProposal", resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_LaHandler *LaHandlerSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.ExecuteProposal(&_LaHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_LaHandler *LaHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.ExecuteProposal(&_LaHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// Initialise is a paid mutator transaction binding the contract method 0x994731da.
//
// Solidity: function initialise(address initBridgeAddress, address initDEXAddress, address initWETHAddress, address ownerAddress) returns()
func (_LaHandler *LaHandlerTransactor) Initialise(opts *bind.TransactOpts, initBridgeAddress common.Address, initDEXAddress common.Address, initWETHAddress common.Address, ownerAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "initialise", initBridgeAddress, initDEXAddress, initWETHAddress, ownerAddress)
}

// Initialise is a paid mutator transaction binding the contract method 0x994731da.
//
// Solidity: function initialise(address initBridgeAddress, address initDEXAddress, address initWETHAddress, address ownerAddress) returns()
func (_LaHandler *LaHandlerSession) Initialise(initBridgeAddress common.Address, initDEXAddress common.Address, initWETHAddress common.Address, ownerAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.Initialise(&_LaHandler.TransactOpts, initBridgeAddress, initDEXAddress, initWETHAddress, ownerAddress)
}

// Initialise is a paid mutator transaction binding the contract method 0x994731da.
//
// Solidity: function initialise(address initBridgeAddress, address initDEXAddress, address initWETHAddress, address ownerAddress) returns()
func (_LaHandler *LaHandlerTransactorSession) Initialise(initBridgeAddress common.Address, initDEXAddress common.Address, initWETHAddress common.Address, ownerAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.Initialise(&_LaHandler.TransactOpts, initBridgeAddress, initDEXAddress, initWETHAddress, ownerAddress)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaHandler *LaHandlerTransactor) OwnableInit(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "ownableInit", owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaHandler *LaHandlerSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.OwnableInit(&_LaHandler.TransactOpts, owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaHandler *LaHandlerTransactorSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.OwnableInit(&_LaHandler.TransactOpts, owner_)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_LaHandler *LaHandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_LaHandler *LaHandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.SetBurnable(&_LaHandler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_LaHandler *LaHandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.SetBurnable(&_LaHandler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_LaHandler *LaHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_LaHandler *LaHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.SetResource(&_LaHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_LaHandler *LaHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.SetResource(&_LaHandler.TransactOpts, resourceID, contractAddress)
}

// TransferExtraLA is a paid mutator transaction binding the contract method 0x9488e9fb.
//
// Solidity: function transferExtraLA(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(bool)
func (_LaHandler *LaHandlerTransactor) TransferExtraLA(opts *bind.TransactOpts, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "transferExtraLA", resourceID, recipientAddress, amount, params)
}

// TransferExtraLA is a paid mutator transaction binding the contract method 0x9488e9fb.
//
// Solidity: function transferExtraLA(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(bool)
func (_LaHandler *LaHandlerSession) TransferExtraLA(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.TransferExtraLA(&_LaHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// TransferExtraLA is a paid mutator transaction binding the contract method 0x9488e9fb.
//
// Solidity: function transferExtraLA(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(bool)
func (_LaHandler *LaHandlerTransactorSession) TransferExtraLA(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaHandler.Contract.TransferExtraLA(&_LaHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaHandler *LaHandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaHandler *LaHandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.TransferOwnership(&_LaHandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaHandler *LaHandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaHandler.Contract.TransferOwnership(&_LaHandler.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_LaHandler *LaHandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_LaHandler *LaHandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.Contract.Withdraw(&_LaHandler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_LaHandler *LaHandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaHandler.Contract.Withdraw(&_LaHandler.TransactOpts, tokenAddress, recipient, amount)
}

// LaHandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LaHandler contract.
type LaHandlerOwnershipTransferredIterator struct {
	Event *LaHandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LaHandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaHandlerOwnershipTransferred)
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
		it.Event = new(LaHandlerOwnershipTransferred)
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
func (it *LaHandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaHandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaHandlerOwnershipTransferred represents a OwnershipTransferred event raised by the LaHandler contract.
type LaHandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaHandler *LaHandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LaHandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaHandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LaHandlerOwnershipTransferredIterator{contract: _LaHandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaHandler *LaHandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LaHandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaHandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaHandlerOwnershipTransferred)
				if err := _LaHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LaHandler *LaHandlerFilterer) ParseOwnershipTransferred(log types.Log) (*LaHandlerOwnershipTransferred, error) {
	event := new(LaHandlerOwnershipTransferred)
	if err := _LaHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
