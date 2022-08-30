// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethBr

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

// EthBrMetaData contains all meta data concerning the EthBr contract.
var EthBrMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"ExtraFeeSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_backendSrvAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_balancerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_executedProposals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isInitialised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nativeResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminCollectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceIDOwner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"resourceIDSpender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"approveSpending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"autobounceExtraLA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes8\",\"name\":\"chainId\",\"type\":\"bytes8\"}],\"name\":\"createMesssageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToLA\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"handlers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initBackendSrvAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initBalancerAddress_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"internalDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownableInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBackendSrv\",\"type\":\"address\"}],\"name\":\"setBackendSrv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBalancer\",\"type\":\"address\"}],\"name\":\"setBalancerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setNativeResourceID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// EthBrABI is the input ABI used to generate the binding from.
// Deprecated: Use EthBrMetaData.ABI instead.
var EthBrABI = EthBrMetaData.ABI

// EthBr is an auto generated Go binding around an Ethereum contract.
type EthBr struct {
	EthBrCaller     // Read-only binding to the contract
	EthBrTransactor // Write-only binding to the contract
	EthBrFilterer   // Log filterer for contract events
}

// EthBrCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthBrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthBrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthBrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthBrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthBrSession struct {
	Contract     *EthBr            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthBrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthBrCallerSession struct {
	Contract *EthBrCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthBrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthBrTransactorSession struct {
	Contract     *EthBrTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthBrRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthBrRaw struct {
	Contract *EthBr // Generic contract binding to access the raw methods on
}

// EthBrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthBrCallerRaw struct {
	Contract *EthBrCaller // Generic read-only contract binding to access the raw methods on
}

// EthBrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthBrTransactorRaw struct {
	Contract *EthBrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthBr creates a new instance of EthBr, bound to a specific deployed contract.
func NewEthBr(address common.Address, backend bind.ContractBackend) (*EthBr, error) {
	contract, err := bindEthBr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthBr{EthBrCaller: EthBrCaller{contract: contract}, EthBrTransactor: EthBrTransactor{contract: contract}, EthBrFilterer: EthBrFilterer{contract: contract}}, nil
}

// NewEthBrCaller creates a new read-only instance of EthBr, bound to a specific deployed contract.
func NewEthBrCaller(address common.Address, caller bind.ContractCaller) (*EthBrCaller, error) {
	contract, err := bindEthBr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthBrCaller{contract: contract}, nil
}

// NewEthBrTransactor creates a new write-only instance of EthBr, bound to a specific deployed contract.
func NewEthBrTransactor(address common.Address, transactor bind.ContractTransactor) (*EthBrTransactor, error) {
	contract, err := bindEthBr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthBrTransactor{contract: contract}, nil
}

// NewEthBrFilterer creates a new log filterer instance of EthBr, bound to a specific deployed contract.
func NewEthBrFilterer(address common.Address, filterer bind.ContractFilterer) (*EthBrFilterer, error) {
	contract, err := bindEthBr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthBrFilterer{contract: contract}, nil
}

// bindEthBr binds a generic wrapper to an already deployed contract.
func bindEthBr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthBrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBr *EthBrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBr.Contract.EthBrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBr *EthBrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.Contract.EthBrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBr *EthBrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBr.Contract.EthBrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthBr *EthBrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthBr *EthBrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthBr *EthBrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBr.Contract.contract.Transact(opts, method, params...)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_EthBr *EthBrCaller) BackendSrvAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_backendSrvAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_EthBr *EthBrSession) BackendSrvAddress() (common.Address, error) {
	return _EthBr.Contract.BackendSrvAddress(&_EthBr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_EthBr *EthBrCallerSession) BackendSrvAddress() (common.Address, error) {
	return _EthBr.Contract.BackendSrvAddress(&_EthBr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_EthBr *EthBrCaller) BalancerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_balancerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_EthBr *EthBrSession) BalancerAddress() (common.Address, error) {
	return _EthBr.Contract.BalancerAddress(&_EthBr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_EthBr *EthBrCallerSession) BalancerAddress() (common.Address, error) {
	return _EthBr.Contract.BalancerAddress(&_EthBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrSession) ChainID() ([8]byte, error) {
	return _EthBr.Contract.ChainID(&_EthBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_EthBr *EthBrCallerSession) ChainID() ([8]byte, error) {
	return _EthBr.Contract.ChainID(&_EthBr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _EthBr.Contract.DepositCounts(&_EthBr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_EthBr *EthBrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _EthBr.Contract.DepositCounts(&_EthBr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _EthBr.Contract.DepositRecords(&_EthBr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_EthBr *EthBrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _EthBr.Contract.DepositRecords(&_EthBr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrCaller) ExecutedProposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_executedProposals", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _EthBr.Contract.ExecutedProposals(&_EthBr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_EthBr *EthBrCallerSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _EthBr.Contract.ExecutedProposals(&_EthBr.CallOpts, arg0, arg1)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrSession) Fee() (*big.Int, error) {
	return _EthBr.Contract.Fee(&_EthBr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_EthBr *EthBrCallerSession) Fee() (*big.Int, error) {
	return _EthBr.Contract.Fee(&_EthBr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthBr *EthBrCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthBr *EthBrSession) IsInitialised() (bool, error) {
	return _EthBr.Contract.IsInitialised(&_EthBr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_EthBr *EthBrCallerSession) IsInitialised() (bool, error) {
	return _EthBr.Contract.IsInitialised(&_EthBr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_EthBr *EthBrCaller) NativeResourceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_nativeResourceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_EthBr *EthBrSession) NativeResourceID() ([32]byte, error) {
	return _EthBr.Contract.NativeResourceID(&_EthBr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_EthBr *EthBrCallerSession) NativeResourceID() ([32]byte, error) {
	return _EthBr.Contract.NativeResourceID(&_EthBr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _EthBr.Contract.ResourceIDToHandlerAddress(&_EthBr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_EthBr *EthBrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _EthBr.Contract.ResourceIDToHandlerAddress(&_EthBr.CallOpts, arg0)
}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_EthBr *EthBrCaller) CreateMesssageHash(opts *bind.CallOpts, amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "createMesssageHash", amount, recipient, chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_EthBr *EthBrSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _EthBr.Contract.CreateMesssageHash(&_EthBr.CallOpts, amount, recipient, chainId)
}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_EthBr *EthBrCallerSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _EthBr.Contract.CreateMesssageHash(&_EthBr.CallOpts, amount, recipient, chainId)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_EthBr *EthBrCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_EthBr *EthBrSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _EthBr.Contract.GetEthSignedMessageHash(&_EthBr.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_EthBr *EthBrCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _EthBr.Contract.GetEthSignedMessageHash(&_EthBr.CallOpts, _messageHash)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_EthBr *EthBrCaller) GetSigner(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "getSigner", messageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_EthBr *EthBrSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _EthBr.Contract.GetSigner(&_EthBr.CallOpts, messageHash, signature)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_EthBr *EthBrCallerSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _EthBr.Contract.GetSigner(&_EthBr.CallOpts, messageHash, signature)
}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_EthBr *EthBrCaller) Handlers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "handlers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_EthBr *EthBrSession) Handlers(arg0 common.Address) (bool, error) {
	return _EthBr.Contract.Handlers(&_EthBr.CallOpts, arg0)
}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_EthBr *EthBrCallerSession) Handlers(arg0 common.Address) (bool, error) {
	return _EthBr.Contract.Handlers(&_EthBr.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthBr *EthBrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthBr *EthBrSession) Owner() (common.Address, error) {
	return _EthBr.Contract.Owner(&_EthBr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthBr *EthBrCallerSession) Owner() (common.Address, error) {
	return _EthBr.Contract.Owner(&_EthBr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrSession) Paused() (bool, error) {
	return _EthBr.Contract.Paused(&_EthBr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthBr *EthBrCallerSession) Paused() (bool, error) {
	return _EthBr.Contract.Paused(&_EthBr.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _EthBr.Contract.Sub(&_EthBr.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_EthBr *EthBrCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _EthBr.Contract.Sub(&_EthBr.CallOpts, a, b)
}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_EthBr *EthBrCaller) Verify(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	var out []interface{}
	err := _EthBr.contract.Call(opts, &out, "verify", _ethSignedMessageHash, _signer, _signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_EthBr *EthBrSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _EthBr.Contract.Verify(&_EthBr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_EthBr *EthBrCallerSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _EthBr.Contract.Verify(&_EthBr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_EthBr *EthBrTransactor) AdminCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminCollectFees", recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_EthBr *EthBrSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminCollectFees(&_EthBr.TransactOpts, recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_EthBr *EthBrTransactorSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminCollectFees(&_EthBr.TransactOpts, recipient, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminPauseTransfers(&_EthBr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_EthBr *EthBrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminPauseTransfers(&_EthBr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminUnpauseTransfers(&_EthBr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_EthBr *EthBrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _EthBr.Contract.AdminUnpauseTransfers(&_EthBr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminWithdraw(&_EthBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.AdminWithdraw(&_EthBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactor) ApproveSpending(opts *bind.TransactOpts, resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "approveSpending", resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ApproveSpending(&_EthBr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_EthBr *EthBrTransactorSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ApproveSpending(&_EthBr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// AutobounceExtraLA is a paid mutator transaction binding the contract method 0xc09e63a8.
//
// Solidity: function autobounceExtraLA(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_EthBr *EthBrTransactor) AutobounceExtraLA(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "autobounceExtraLA", originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// AutobounceExtraLA is a paid mutator transaction binding the contract method 0xc09e63a8.
//
// Solidity: function autobounceExtraLA(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_EthBr *EthBrSession) AutobounceExtraLA(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.AutobounceExtraLA(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// AutobounceExtraLA is a paid mutator transaction binding the contract method 0xc09e63a8.
//
// Solidity: function autobounceExtraLA(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_EthBr *EthBrTransactorSession) AutobounceExtraLA(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.AutobounceExtraLA(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_EthBr *EthBrTransactor) ChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "changeFee", newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_EthBr *EthBrSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ChangeFee(&_EthBr.TransactOpts, newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_EthBr *EthBrTransactorSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.ChangeFee(&_EthBr.TransactOpts, newFee)
}

// Deposit is a paid mutator transaction binding the contract method 0x2c392157.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes signature, bytes params) payable returns()
func (_EthBr *EthBrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, signature []byte, params []byte) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress, amountToLA, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x2c392157.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes signature, bytes params) payable returns()
func (_EthBr *EthBrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, signature []byte, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.Deposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, amountToLA, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x2c392157.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes signature, bytes params) payable returns()
func (_EthBr *EthBrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, signature []byte, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.Deposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, amountToLA, signature, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthBr *EthBrTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthBr *EthBrSession) DepositFunds() (*types.Transaction, error) {
	return _EthBr.Contract.DepositFunds(&_EthBr.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_EthBr *EthBrTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _EthBr.Contract.DepositFunds(&_EthBr.TransactOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x6717ec72.
//
// Solidity: function depositNativeToken(bytes32 resourceID, uint256 amount) payable returns()
func (_EthBr *EthBrTransactor) DepositNativeToken(opts *bind.TransactOpts, resourceID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "depositNativeToken", resourceID, amount)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x6717ec72.
//
// Solidity: function depositNativeToken(bytes32 resourceID, uint256 amount) payable returns()
func (_EthBr *EthBrSession) DepositNativeToken(resourceID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.DepositNativeToken(&_EthBr.TransactOpts, resourceID, amount)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x6717ec72.
//
// Solidity: function depositNativeToken(bytes32 resourceID, uint256 amount) payable returns()
func (_EthBr *EthBrTransactorSession) DepositNativeToken(resourceID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EthBr.Contract.DepositNativeToken(&_EthBr.TransactOpts, resourceID, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_EthBr *EthBrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_EthBr *EthBrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.ExecuteProposal(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_EthBr *EthBrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _EthBr.Contract.ExecuteProposal(&_EthBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x711f4a41.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress, address initBalancerAddress_) returns()
func (_EthBr *EthBrTransactor) Initialize(opts *bind.TransactOpts, chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "initialize", chainID, fee, initBackendSrvAddress, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x711f4a41.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress, address initBalancerAddress_) returns()
func (_EthBr *EthBrSession) Initialize(chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.Initialize(&_EthBr.TransactOpts, chainID, fee, initBackendSrvAddress, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x711f4a41.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress, address initBalancerAddress_) returns()
func (_EthBr *EthBrTransactorSession) Initialize(chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.Initialize(&_EthBr.TransactOpts, chainID, fee, initBackendSrvAddress, initBalancerAddress_)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_EthBr *EthBrTransactor) InternalDeposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "internalDeposit", destinationChainID, resourceID, amount, recipientAddress)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_EthBr *EthBrSession) InternalDeposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.InternalDeposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_EthBr *EthBrTransactorSession) InternalDeposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.InternalDeposit(&_EthBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthBr *EthBrTransactor) OwnableInit(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "ownableInit", owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthBr *EthBrSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.OwnableInit(&_EthBr.TransactOpts, owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_EthBr *EthBrTransactorSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.OwnableInit(&_EthBr.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthBr *EthBrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthBr *EthBrSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthBr.Contract.RenounceOwnership(&_EthBr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthBr *EthBrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthBr.Contract.RenounceOwnership(&_EthBr.TransactOpts)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrTransactor) SetBackendSrv(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setBackendSrv", newBackendSrv)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrSession) SetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBackendSrv(&_EthBr.TransactOpts, newBackendSrv)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_EthBr *EthBrTransactorSession) SetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBackendSrv(&_EthBr.TransactOpts, newBackendSrv)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_EthBr *EthBrTransactor) SetBalancerAddress(opts *bind.TransactOpts, newBalancer common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setBalancerAddress", newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_EthBr *EthBrSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBalancerAddress(&_EthBr.TransactOpts, newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_EthBr *EthBrTransactorSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBalancerAddress(&_EthBr.TransactOpts, newBalancer)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrTransactor) SetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setBurnable", handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBurnable(&_EthBr.TransactOpts, handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_EthBr *EthBrTransactorSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetBurnable(&_EthBr.TransactOpts, handlerAddress, tokenAddress)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_EthBr *EthBrTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, value bool) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setHandler", _handler, value)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_EthBr *EthBrSession) SetHandler(_handler common.Address, value bool) (*types.Transaction, error) {
	return _EthBr.Contract.SetHandler(&_EthBr.TransactOpts, _handler, value)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_EthBr *EthBrTransactorSession) SetHandler(_handler common.Address, value bool) (*types.Transaction, error) {
	return _EthBr.Contract.SetHandler(&_EthBr.TransactOpts, _handler, value)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_EthBr *EthBrTransactor) SetNativeResourceID(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setNativeResourceID", resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_EthBr *EthBrSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _EthBr.Contract.SetNativeResourceID(&_EthBr.TransactOpts, resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_EthBr *EthBrTransactorSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _EthBr.Contract.SetNativeResourceID(&_EthBr.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrTransactor) SetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "setResource", handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetResource(&_EthBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_EthBr *EthBrTransactorSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.SetResource(&_EthBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthBr *EthBrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthBr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthBr *EthBrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.TransferOwnership(&_EthBr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthBr *EthBrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthBr.Contract.TransferOwnership(&_EthBr.TransactOpts, newOwner)
}

// EthBrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EthBr contract.
type EthBrDepositIterator struct {
	Event *EthBrDeposit // Event containing the contract specifics and raw log

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
func (it *EthBrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrDeposit)
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
		it.Event = new(EthBrDeposit)
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
func (it *EthBrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrDeposit represents a Deposit event raised by the EthBr contract.
type EthBrDeposit struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Params             []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_EthBr *EthBrFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (*EthBrDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &EthBrDepositIterator{contract: _EthBr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_EthBr *EthBrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthBrDeposit, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrDeposit)
				if err := _EthBr.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_EthBr *EthBrFilterer) ParseDeposit(log types.Log) (*EthBrDeposit, error) {
	event := new(EthBrDeposit)
	if err := _EthBr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrExtraFeeSuppliedIterator is returned from FilterExtraFeeSupplied and is used to iterate over the raw logs and unpacked data for ExtraFeeSupplied events raised by the EthBr contract.
type EthBrExtraFeeSuppliedIterator struct {
	Event *EthBrExtraFeeSupplied // Event containing the contract specifics and raw log

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
func (it *EthBrExtraFeeSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrExtraFeeSupplied)
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
		it.Event = new(EthBrExtraFeeSupplied)
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
func (it *EthBrExtraFeeSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrExtraFeeSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrExtraFeeSupplied represents a ExtraFeeSupplied event raised by the EthBr contract.
type EthBrExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResourceID         [32]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Status             uint8
	Params             []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtraFeeSupplied is a free log retrieval operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_EthBr *EthBrFilterer) FilterExtraFeeSupplied(opts *bind.FilterOpts) (*EthBrExtraFeeSuppliedIterator, error) {

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return &EthBrExtraFeeSuppliedIterator{contract: _EthBr.contract, event: "ExtraFeeSupplied", logs: logs, sub: sub}, nil
}

// WatchExtraFeeSupplied is a free log subscription operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_EthBr *EthBrFilterer) WatchExtraFeeSupplied(opts *bind.WatchOpts, sink chan<- *EthBrExtraFeeSupplied) (event.Subscription, error) {

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrExtraFeeSupplied)
				if err := _EthBr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
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

// ParseExtraFeeSupplied is a log parse operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_EthBr *EthBrFilterer) ParseExtraFeeSupplied(log types.Log) (*EthBrExtraFeeSupplied, error) {
	event := new(EthBrExtraFeeSupplied)
	if err := _EthBr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthBr contract.
type EthBrOwnershipTransferredIterator struct {
	Event *EthBrOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthBrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrOwnershipTransferred)
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
		it.Event = new(EthBrOwnershipTransferred)
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
func (it *EthBrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrOwnershipTransferred represents a OwnershipTransferred event raised by the EthBr contract.
type EthBrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthBr *EthBrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthBrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthBrOwnershipTransferredIterator{contract: _EthBr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthBr *EthBrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthBrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrOwnershipTransferred)
				if err := _EthBr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthBr *EthBrFilterer) ParseOwnershipTransferred(log types.Log) (*EthBrOwnershipTransferred, error) {
	event := new(EthBrOwnershipTransferred)
	if err := _EthBr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthBr contract.
type EthBrPausedIterator struct {
	Event *EthBrPaused // Event containing the contract specifics and raw log

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
func (it *EthBrPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrPaused)
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
		it.Event = new(EthBrPaused)
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
func (it *EthBrPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrPaused represents a Paused event raised by the EthBr contract.
type EthBrPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthBr *EthBrFilterer) FilterPaused(opts *bind.FilterOpts) (*EthBrPausedIterator, error) {

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthBrPausedIterator{contract: _EthBr.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthBr *EthBrFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthBrPaused) (event.Subscription, error) {

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrPaused)
				if err := _EthBr.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthBr *EthBrFilterer) ParsePaused(log types.Log) (*EthBrPaused, error) {
	event := new(EthBrPaused)
	if err := _EthBr.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the EthBr contract.
type EthBrProposalEventIterator struct {
	Event *EthBrProposalEvent // Event containing the contract specifics and raw log

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
func (it *EthBrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrProposalEvent)
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
		it.Event = new(EthBrProposalEvent)
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
func (it *EthBrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrProposalEvent represents a ProposalEvent event raised by the EthBr contract.
type EthBrProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_EthBr *EthBrFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (*EthBrProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return &EthBrProposalEventIterator{contract: _EthBr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_EthBr *EthBrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *EthBrProposalEvent, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrProposalEvent)
				if err := _EthBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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

// ParseProposalEvent is a log parse operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_EthBr *EthBrFilterer) ParseProposalEvent(log types.Log) (*EthBrProposalEvent, error) {
	event := new(EthBrProposalEvent)
	if err := _EthBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthBrUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthBr contract.
type EthBrUnpausedIterator struct {
	Event *EthBrUnpaused // Event containing the contract specifics and raw log

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
func (it *EthBrUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBrUnpaused)
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
		it.Event = new(EthBrUnpaused)
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
func (it *EthBrUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBrUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBrUnpaused represents a Unpaused event raised by the EthBr contract.
type EthBrUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthBr *EthBrFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthBrUnpausedIterator, error) {

	logs, sub, err := _EthBr.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthBrUnpausedIterator{contract: _EthBr.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthBr *EthBrFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthBrUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthBr.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBrUnpaused)
				if err := _EthBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthBr *EthBrFilterer) ParseUnpaused(log types.Log) (*EthBrUnpaused, error) {
	event := new(EthBrUnpaused)
	if err := _EthBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
