// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethHandler

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

// EthHandlerMetaData contains all meta data concerning the EthHandler contract.
var EthHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes8\",\"name\":\"_destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isInitialised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_minDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBridgeAddress\",\"type\":\"address\"}],\"name\":\"adminChangeBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"getAddressFromResourceId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"destId\",\"type\":\"bytes8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes8\",\"name\":\"_destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_destinationRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"getMinDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownableInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EthHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use EthHandlerMetaData.ABI instead.
var EthHandlerABI = EthHandlerMetaData.ABI

// EthHandler is an auto generated Go binding around an Ethereum contract.
type EthHandler struct {
	EthHandlerCaller     // Read-only binding to the contract
	EthHandlerTransactor // Write-only binding to the contract
	EthHandlerFilterer   // Log filterer for contract events
}

// EthHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthHandlerSession struct {
	Contract     *EthHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthHandlerCallerSession struct {
	Contract *EthHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EthHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthHandlerTransactorSession struct {
	Contract     *EthHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthHandlerRaw struct {
	Contract *EthHandler // Generic contract binding to access the raw methods on
}

// EthHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthHandlerCallerRaw struct {
	Contract *EthHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// EthHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthHandlerTransactorRaw struct {
	Contract *EthHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthHandler creates a new instance of EthHandler, bound to a specific deployed contract.
func NewEthHandler(address common.Address, backend bind.ContractBackend) (*EthHandler, error) {
	contract, err := bindEthHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthHandler{EthHandlerCaller: EthHandlerCaller{contract: contract}, EthHandlerTransactor: EthHandlerTransactor{contract: contract}, EthHandlerFilterer: EthHandlerFilterer{contract: contract}}, nil
}

// NewEthHandlerCaller creates a new read-only instance of EthHandler, bound to a specific deployed contract.
func NewEthHandlerCaller(address common.Address, caller bind.ContractCaller) (*EthHandlerCaller, error) {
	contract, err := bindEthHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthHandlerCaller{contract: contract}, nil
}

// NewEthHandlerTransactor creates a new write-only instance of EthHandler, bound to a specific deployed contract.
func NewEthHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthHandlerTransactor, error) {
	contract, err := bindEthHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthHandlerTransactor{contract: contract}, nil
}

// NewEthHandlerFilterer creates a new log filterer instance of EthHandler, bound to a specific deployed contract.
func NewEthHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthHandlerFilterer, error) {
	contract, err := bindEthHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthHandlerFilterer{contract: contract}, nil
}

// bindEthHandler binds a generic wrapper to an already deployed contract.
func bindEthHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthHandler *EthHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthHandler.Contract.EthHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthHandler *EthHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthHandler.Contract.EthHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthHandler *EthHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthHandler.Contract.EthHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthHandler *EthHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthHandler *EthHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthHandler *EthHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_EthHandler *EthHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_EthHandler *EthHandlerSession) BridgeAddress() (common.Address, error) {
	return _EthHandler.Contract.BridgeAddress(&_EthHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_EthHandler *EthHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _EthHandler.Contract.BridgeAddress(&_EthHandler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_EthHandler *EthHandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_EthHandler *EthHandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _EthHandler.Contract.BurnList(&_EthHandler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_EthHandler *EthHandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _EthHandler.Contract.BurnList(&_EthHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_EthHandler *EthHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_EthHandler *EthHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _EthHandler.Contract.ContractWhitelist(&_EthHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_EthHandler *EthHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _EthHandler.Contract.ContractWhitelist(&_EthHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns(address _tokenAddress, bytes8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_EthHandler *EthHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 [8]byte, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          [8]byte
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	Amount                      *big.Int
}, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                common.Address
		DestinationChainID          [8]byte
		ResourceID                  [32]byte
		DestinationRecipientAddress common.Address
		Depositer                   common.Address
		Amount                      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DestinationChainID = *abi.ConvertType(out[1], new([8]byte)).(*[8]byte)
	outstruct.ResourceID = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DestinationRecipientAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Depositer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns(address _tokenAddress, bytes8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_EthHandler *EthHandlerSession) DepositRecords(arg0 [8]byte, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          [8]byte
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	Amount                      *big.Int
}, error) {
	return _EthHandler.Contract.DepositRecords(&_EthHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x78f6a94c.
//
// Solidity: function _depositRecords(bytes8 , uint64 ) view returns(address _tokenAddress, bytes8 _destinationChainID, bytes32 _resourceID, address _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_EthHandler *EthHandlerCallerSession) DepositRecords(arg0 [8]byte, arg1 uint64) (struct {
	TokenAddress                common.Address
	DestinationChainID          [8]byte
	ResourceID                  [32]byte
	DestinationRecipientAddress common.Address
	Depositer                   common.Address
	Amount                      *big.Int
}, error) {
	return _EthHandler.Contract.DepositRecords(&_EthHandler.CallOpts, arg0, arg1)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthHandler *EthHandlerCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthHandler *EthHandlerSession) IsInitialised() (bool, error) {
	return _EthHandler.Contract.IsInitialised(&_EthHandler.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthHandler *EthHandlerCallerSession) IsInitialised() (bool, error) {
	return _EthHandler.Contract.IsInitialised(&_EthHandler.CallOpts)
}

// MinDeposits is a free data retrieval call binding the contract method 0xd90445ca.
//
// Solidity: function _minDeposits(bytes32 ) view returns(uint256)
func (_EthHandler *EthHandlerCaller) MinDeposits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_minDeposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposits is a free data retrieval call binding the contract method 0xd90445ca.
//
// Solidity: function _minDeposits(bytes32 ) view returns(uint256)
func (_EthHandler *EthHandlerSession) MinDeposits(arg0 [32]byte) (*big.Int, error) {
	return _EthHandler.Contract.MinDeposits(&_EthHandler.CallOpts, arg0)
}

// MinDeposits is a free data retrieval call binding the contract method 0xd90445ca.
//
// Solidity: function _minDeposits(bytes32 ) view returns(uint256)
func (_EthHandler *EthHandlerCallerSession) MinDeposits(arg0 [32]byte) (*big.Int, error) {
	return _EthHandler.Contract.MinDeposits(&_EthHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_EthHandler *EthHandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_EthHandler *EthHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _EthHandler.Contract.ResourceIDToTokenContractAddress(&_EthHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_EthHandler *EthHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _EthHandler.Contract.ResourceIDToTokenContractAddress(&_EthHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_EthHandler *EthHandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_EthHandler *EthHandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _EthHandler.Contract.TokenContractAddressToResourceID(&_EthHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_EthHandler *EthHandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _EthHandler.Contract.TokenContractAddressToResourceID(&_EthHandler.CallOpts, arg0)
}

// GetAddressFromResourceId is a free data retrieval call binding the contract method 0xa7a3fb40.
//
// Solidity: function getAddressFromResourceId(bytes32 resourceID) view returns(address)
func (_EthHandler *EthHandlerCaller) GetAddressFromResourceId(opts *bind.CallOpts, resourceID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "getAddressFromResourceId", resourceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromResourceId is a free data retrieval call binding the contract method 0xa7a3fb40.
//
// Solidity: function getAddressFromResourceId(bytes32 resourceID) view returns(address)
func (_EthHandler *EthHandlerSession) GetAddressFromResourceId(resourceID [32]byte) (common.Address, error) {
	return _EthHandler.Contract.GetAddressFromResourceId(&_EthHandler.CallOpts, resourceID)
}

// GetAddressFromResourceId is a free data retrieval call binding the contract method 0xa7a3fb40.
//
// Solidity: function getAddressFromResourceId(bytes32 resourceID) view returns(address)
func (_EthHandler *EthHandlerCallerSession) GetAddressFromResourceId(resourceID [32]byte) (common.Address, error) {
	return _EthHandler.Contract.GetAddressFromResourceId(&_EthHandler.CallOpts, resourceID)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_EthHandler *EthHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC20HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20HandlerDepositRecord)).(*ERC20HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_EthHandler *EthHandlerSession) GetDepositRecord(depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	return _EthHandler.Contract.GetDepositRecord(&_EthHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0x34143cd7.
//
// Solidity: function getDepositRecord(uint64 depositNonce, bytes8 destId) view returns((address,bytes8,bytes32,address,address,uint256))
func (_EthHandler *EthHandlerCallerSession) GetDepositRecord(depositNonce uint64, destId [8]byte) (ERC20HandlerDepositRecord, error) {
	return _EthHandler.Contract.GetDepositRecord(&_EthHandler.CallOpts, depositNonce, destId)
}

// GetMinDeposits is a free data retrieval call binding the contract method 0x27e4b2f7.
//
// Solidity: function getMinDeposits(bytes32 resourceID) view returns(uint256)
func (_EthHandler *EthHandlerCaller) GetMinDeposits(opts *bind.CallOpts, resourceID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "getMinDeposits", resourceID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDeposits is a free data retrieval call binding the contract method 0x27e4b2f7.
//
// Solidity: function getMinDeposits(bytes32 resourceID) view returns(uint256)
func (_EthHandler *EthHandlerSession) GetMinDeposits(resourceID [32]byte) (*big.Int, error) {
	return _EthHandler.Contract.GetMinDeposits(&_EthHandler.CallOpts, resourceID)
}

// GetMinDeposits is a free data retrieval call binding the contract method 0x27e4b2f7.
//
// Solidity: function getMinDeposits(bytes32 resourceID) view returns(uint256)
func (_EthHandler *EthHandlerCallerSession) GetMinDeposits(resourceID [32]byte) (*big.Int, error) {
	return _EthHandler.Contract.GetMinDeposits(&_EthHandler.CallOpts, resourceID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthHandler *EthHandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthHandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthHandler *EthHandlerSession) Owner() (common.Address, error) {
	return _EthHandler.Contract.Owner(&_EthHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthHandler *EthHandlerCallerSession) Owner() (common.Address, error) {
	return _EthHandler.Contract.Owner(&_EthHandler.CallOpts)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_EthHandler *EthHandlerTransactor) AdminChangeBridgeAddress(opts *bind.TransactOpts, newBridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "adminChangeBridgeAddress", newBridgeAddress)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_EthHandler *EthHandlerSession) AdminChangeBridgeAddress(newBridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.AdminChangeBridgeAddress(&_EthHandler.TransactOpts, newBridgeAddress)
}

// AdminChangeBridgeAddress is a paid mutator transaction binding the contract method 0x292f3a68.
//
// Solidity: function adminChangeBridgeAddress(address newBridgeAddress) returns()
func (_EthHandler *EthHandlerTransactorSession) AdminChangeBridgeAddress(newBridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.AdminChangeBridgeAddress(&_EthHandler.TransactOpts, newBridgeAddress)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactor) Approve(opts *bind.TransactOpts, resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "approve", resourceID, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_EthHandler *EthHandlerSession) Approve(resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.Approve(&_EthHandler.TransactOpts, resourceID, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xbf1ed1eb.
//
// Solidity: function approve(bytes32 resourceID, address spender, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactorSession) Approve(resourceID [32]byte, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.Approve(&_EthHandler.TransactOpts, resourceID, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_EthHandler *EthHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_EthHandler *EthHandlerSession) Deposit(resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.Contract.Deposit(&_EthHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x75cd9cce.
//
// Solidity: function deposit(bytes32 resourceID, bytes8 destinationChainID, uint64 depositNonce, address depositer, address recipientAddress, uint256 amount, bytes params) returns(address, uint256)
func (_EthHandler *EthHandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID [8]byte, depositNonce uint64, depositer common.Address, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.Contract.Deposit(&_EthHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipientAddress, amount, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthHandler *EthHandlerTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthHandler *EthHandlerSession) DepositFunds() (*types.Transaction, error) {
	return _EthHandler.Contract.DepositFunds(&_EthHandler.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthHandler *EthHandlerTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _EthHandler.Contract.DepositFunds(&_EthHandler.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_EthHandler *EthHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "executeProposal", resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_EthHandler *EthHandlerSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.Contract.ExecuteProposal(&_EthHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns(uint256)
func (_EthHandler *EthHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthHandler.Contract.ExecuteProposal(&_EthHandler.TransactOpts, resourceID, recipientAddress, amount, params)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_EthHandler *EthHandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.FundERC20(&_EthHandler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.FundERC20(&_EthHandler.TransactOpts, tokenAddress, owner, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridgeAddress) returns()
func (_EthHandler *EthHandlerTransactor) Initialize(opts *bind.TransactOpts, bridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "initialize", bridgeAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridgeAddress) returns()
func (_EthHandler *EthHandlerSession) Initialize(bridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.Initialize(&_EthHandler.TransactOpts, bridgeAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address bridgeAddress) returns()
func (_EthHandler *EthHandlerTransactorSession) Initialize(bridgeAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.Initialize(&_EthHandler.TransactOpts, bridgeAddress)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthHandler *EthHandlerTransactor) OwnableInit(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "ownableInit", owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthHandler *EthHandlerSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.OwnableInit(&_EthHandler.TransactOpts, owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthHandler *EthHandlerTransactorSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.OwnableInit(&_EthHandler.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthHandler *EthHandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthHandler *EthHandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthHandler.Contract.RenounceOwnership(&_EthHandler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthHandler *EthHandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthHandler.Contract.RenounceOwnership(&_EthHandler.TransactOpts)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_EthHandler *EthHandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_EthHandler *EthHandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.SetBurnable(&_EthHandler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_EthHandler *EthHandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.SetBurnable(&_EthHandler.TransactOpts, contractAddress)
}

// SetMinDeposits is a paid mutator transaction binding the contract method 0x06a6906c.
//
// Solidity: function setMinDeposits(bytes32 resourceID, uint256 minAmount) returns()
func (_EthHandler *EthHandlerTransactor) SetMinDeposits(opts *bind.TransactOpts, resourceID [32]byte, minAmount *big.Int) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "setMinDeposits", resourceID, minAmount)
}

// SetMinDeposits is a paid mutator transaction binding the contract method 0x06a6906c.
//
// Solidity: function setMinDeposits(bytes32 resourceID, uint256 minAmount) returns()
func (_EthHandler *EthHandlerSession) SetMinDeposits(resourceID [32]byte, minAmount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.SetMinDeposits(&_EthHandler.TransactOpts, resourceID, minAmount)
}

// SetMinDeposits is a paid mutator transaction binding the contract method 0x06a6906c.
//
// Solidity: function setMinDeposits(bytes32 resourceID, uint256 minAmount) returns()
func (_EthHandler *EthHandlerTransactorSession) SetMinDeposits(resourceID [32]byte, minAmount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.SetMinDeposits(&_EthHandler.TransactOpts, resourceID, minAmount)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_EthHandler *EthHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_EthHandler *EthHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.SetResource(&_EthHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_EthHandler *EthHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.SetResource(&_EthHandler.TransactOpts, resourceID, contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthHandler *EthHandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthHandler *EthHandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.TransferOwnership(&_EthHandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthHandler *EthHandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthHandler.Contract.TransferOwnership(&_EthHandler.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_EthHandler *EthHandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.Withdraw(&_EthHandler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_EthHandler *EthHandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthHandler.Contract.Withdraw(&_EthHandler.TransactOpts, tokenAddress, recipient, amount)
}

// EthHandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthHandler contract.
type EthHandlerOwnershipTransferredIterator struct {
	Event *EthHandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthHandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthHandlerOwnershipTransferred)
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
		it.Event = new(EthHandlerOwnershipTransferred)
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
func (it *EthHandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthHandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthHandlerOwnershipTransferred represents a OwnershipTransferred event raised by the EthHandler contract.
type EthHandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthHandler *EthHandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthHandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthHandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthHandlerOwnershipTransferredIterator{contract: _EthHandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthHandler *EthHandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthHandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthHandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthHandlerOwnershipTransferred)
				if err := _EthHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthHandler *EthHandlerFilterer) ParseOwnershipTransferred(log types.Log) (*EthHandlerOwnershipTransferred, error) {
	event := new(EthHandlerOwnershipTransferred)
	if err := _EthHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
