// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package laBr

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

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// LaBrMetaData contains all meta data concerning the LaBr contract.
var LaBrMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"paused\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"sub\",\"type\":\"function\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"renounceOwnership\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"_chainID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerThreshold\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_totalProposals\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_fee\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_expiry\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerHubAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_backendSrvAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_CHAINID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_RELAYERTHRESHOLD\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_FEE\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"INIT_EXPIRY\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositCounts\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"name\":\"_resourceIDToHandlerAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositRecords\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"name\":\"_proposals\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"_hasVotedOnProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"init\",\"type\":\"function\",\"inputs\":[{\"name\":\"initBackendSrvAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetRelayerHub\",\"type\":\"function\",\"inputs\":[{\"name\":\"newRelayerHub\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetBackendSrvAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBackendSrv\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminPauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminUnpauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminChangeRelayerThreshold\",\"type\":\"function\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetResource\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminSetBurnable\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"adminChangeFee\",\"type\":\"function\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminWithdraw\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountOrTokenID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"deposit\",\"type\":\"function\",\"inputs\":[{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"voteProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"cancelProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"executeProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminCollectFees\",\"type\":\"function\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"relayerCollectReward\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"Paused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"Unpaused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerThresholdChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"Deposit\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalEvent\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalVote\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"RewardCollected\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// LaBrABI is the input ABI used to generate the binding from.
// Deprecated: Use LaBrMetaData.ABI instead.
var LaBrABI = LaBrMetaData.ABI

// LaBr is an auto generated Go binding around an Ethereum contract.
type LaBr struct {
	LaBrCaller     // Read-only binding to the contract
	LaBrTransactor // Write-only binding to the contract
	LaBrFilterer   // Log filterer for contract events
}

// LaBrCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaBrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaBrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaBrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaBrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaBrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaBrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaBrSession struct {
	Contract     *LaBr             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaBrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaBrCallerSession struct {
	Contract *LaBrCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LaBrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaBrTransactorSession struct {
	Contract     *LaBrTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaBrRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaBrRaw struct {
	Contract *LaBr // Generic contract binding to access the raw methods on
}

// LaBrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaBrCallerRaw struct {
	Contract *LaBrCaller // Generic read-only contract binding to access the raw methods on
}

// LaBrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaBrTransactorRaw struct {
	Contract *LaBrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaBr creates a new instance of LaBr, bound to a specific deployed contract.
func NewLaBr(address common.Address, backend bind.ContractBackend) (*LaBr, error) {
	contract, err := bindLaBr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LaBr{LaBrCaller: LaBrCaller{contract: contract}, LaBrTransactor: LaBrTransactor{contract: contract}, LaBrFilterer: LaBrFilterer{contract: contract}}, nil
}

// NewLaBrCaller creates a new read-only instance of LaBr, bound to a specific deployed contract.
func NewLaBrCaller(address common.Address, caller bind.ContractCaller) (*LaBrCaller, error) {
	contract, err := bindLaBr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaBrCaller{contract: contract}, nil
}

// NewLaBrTransactor creates a new write-only instance of LaBr, bound to a specific deployed contract.
func NewLaBrTransactor(address common.Address, transactor bind.ContractTransactor) (*LaBrTransactor, error) {
	contract, err := bindLaBr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaBrTransactor{contract: contract}, nil
}

// NewLaBrFilterer creates a new log filterer instance of LaBr, bound to a specific deployed contract.
func NewLaBrFilterer(address common.Address, filterer bind.ContractFilterer) (*LaBrFilterer, error) {
	contract, err := bindLaBr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaBrFilterer{contract: contract}, nil
}

// bindLaBr binds a generic wrapper to an already deployed contract.
func bindLaBr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LaBrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaBr *LaBrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaBr.Contract.LaBrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaBr *LaBrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.Contract.LaBrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaBr *LaBrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaBr.Contract.LaBrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaBr *LaBrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaBr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaBr *LaBrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaBr *LaBrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaBr.Contract.contract.Transact(opts, method, params...)
}

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_LaBr *LaBrCaller) INITCHAINID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "INIT_CHAINID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_LaBr *LaBrSession) INITCHAINID() ([8]byte, error) {
	return _LaBr.Contract.INITCHAINID(&_LaBr.CallOpts)
}

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_LaBr *LaBrCallerSession) INITCHAINID() ([8]byte, error) {
	return _LaBr.Contract.INITCHAINID(&_LaBr.CallOpts)
}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_LaBr *LaBrCaller) INITEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "INIT_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_LaBr *LaBrSession) INITEXPIRY() (*big.Int, error) {
	return _LaBr.Contract.INITEXPIRY(&_LaBr.CallOpts)
}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_LaBr *LaBrCallerSession) INITEXPIRY() (*big.Int, error) {
	return _LaBr.Contract.INITEXPIRY(&_LaBr.CallOpts)
}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_LaBr *LaBrCaller) INITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "INIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_LaBr *LaBrSession) INITFEE() (*big.Int, error) {
	return _LaBr.Contract.INITFEE(&_LaBr.CallOpts)
}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_LaBr *LaBrCallerSession) INITFEE() (*big.Int, error) {
	return _LaBr.Contract.INITFEE(&_LaBr.CallOpts)
}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_LaBr *LaBrCaller) INITRELAYERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "INIT_RELAYERTHRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_LaBr *LaBrSession) INITRELAYERTHRESHOLD() (*big.Int, error) {
	return _LaBr.Contract.INITRELAYERTHRESHOLD(&_LaBr.CallOpts)
}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_LaBr *LaBrCallerSession) INITRELAYERTHRESHOLD() (*big.Int, error) {
	return _LaBr.Contract.INITRELAYERTHRESHOLD(&_LaBr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_LaBr *LaBrCaller) BackendSrvAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_backendSrvAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_LaBr *LaBrSession) BackendSrvAddress() (common.Address, error) {
	return _LaBr.Contract.BackendSrvAddress(&_LaBr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_LaBr *LaBrCallerSession) BackendSrvAddress() (common.Address, error) {
	return _LaBr.Contract.BackendSrvAddress(&_LaBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_LaBr *LaBrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_LaBr *LaBrSession) ChainID() ([8]byte, error) {
	return _LaBr.Contract.ChainID(&_LaBr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_LaBr *LaBrCallerSession) ChainID() ([8]byte, error) {
	return _LaBr.Contract.ChainID(&_LaBr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_LaBr *LaBrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_LaBr *LaBrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _LaBr.Contract.DepositCounts(&_LaBr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_LaBr *LaBrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _LaBr.Contract.DepositCounts(&_LaBr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_LaBr *LaBrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_LaBr *LaBrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _LaBr.Contract.DepositRecords(&_LaBr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_LaBr *LaBrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _LaBr.Contract.DepositRecords(&_LaBr.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_LaBr *LaBrCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_LaBr *LaBrSession) Expiry() (*big.Int, error) {
	return _LaBr.Contract.Expiry(&_LaBr.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_LaBr *LaBrCallerSession) Expiry() (*big.Int, error) {
	return _LaBr.Contract.Expiry(&_LaBr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_LaBr *LaBrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_LaBr *LaBrSession) Fee() (*big.Int, error) {
	return _LaBr.Contract.Fee(&_LaBr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_LaBr *LaBrCallerSession) Fee() (*big.Int, error) {
	return _LaBr.Contract.Fee(&_LaBr.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_LaBr *LaBrCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_LaBr *LaBrSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _LaBr.Contract.HasVotedOnProposal(&_LaBr.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_LaBr *LaBrCallerSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _LaBr.Contract.HasVotedOnProposal(&_LaBr.CallOpts, arg0, arg1, arg2)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_proposals", arg0, arg1)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrSession) Proposals(arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.Proposals(&_LaBr.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCallerSession) Proposals(arg0 [32]byte, arg1 [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.Proposals(&_LaBr.CallOpts, arg0, arg1)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_LaBr *LaBrCaller) RelayerHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_relayerHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_LaBr *LaBrSession) RelayerHubAddress() (common.Address, error) {
	return _LaBr.Contract.RelayerHubAddress(&_LaBr.CallOpts)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_LaBr *LaBrCallerSession) RelayerHubAddress() (common.Address, error) {
	return _LaBr.Contract.RelayerHubAddress(&_LaBr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_LaBr *LaBrCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_LaBr *LaBrSession) RelayerThreshold() (*big.Int, error) {
	return _LaBr.Contract.RelayerThreshold(&_LaBr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_LaBr *LaBrCallerSession) RelayerThreshold() (*big.Int, error) {
	return _LaBr.Contract.RelayerThreshold(&_LaBr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_LaBr *LaBrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_LaBr *LaBrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _LaBr.Contract.ResourceIDToHandlerAddress(&_LaBr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_LaBr *LaBrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _LaBr.Contract.ResourceIDToHandlerAddress(&_LaBr.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_LaBr *LaBrCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_LaBr *LaBrSession) TotalProposals() (*big.Int, error) {
	return _LaBr.Contract.TotalProposals(&_LaBr.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_LaBr *LaBrCallerSession) TotalProposals() (*big.Int, error) {
	return _LaBr.Contract.TotalProposals(&_LaBr.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCaller) GetProposal(opts *bind.CallOpts, originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "getProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrSession) GetProposal(originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.GetProposal(&_LaBr.CallOpts, originChainID, depositNonce, dataHash)
}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCallerSession) GetProposal(originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.GetProposal(&_LaBr.CallOpts, originChainID, depositNonce, dataHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaBr *LaBrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaBr *LaBrSession) Owner() (common.Address, error) {
	return _LaBr.Contract.Owner(&_LaBr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LaBr *LaBrCallerSession) Owner() (common.Address, error) {
	return _LaBr.Contract.Owner(&_LaBr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LaBr *LaBrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LaBr *LaBrSession) Paused() (bool, error) {
	return _LaBr.Contract.Paused(&_LaBr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LaBr *LaBrCallerSession) Paused() (bool, error) {
	return _LaBr.Contract.Paused(&_LaBr.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_LaBr *LaBrCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_LaBr *LaBrSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _LaBr.Contract.Sub(&_LaBr.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_LaBr *LaBrCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _LaBr.Contract.Sub(&_LaBr.CallOpts, a, b)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_LaBr *LaBrTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_LaBr *LaBrSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminChangeFee(&_LaBr.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_LaBr *LaBrTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminChangeFee(&_LaBr.TransactOpts, newFee)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminChangeRelayerThreshold(&_LaBr.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminChangeRelayerThreshold(&_LaBr.TransactOpts, newThreshold)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_LaBr *LaBrTransactor) AdminCollectFees(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminCollectFees", amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_LaBr *LaBrSession) AdminCollectFees(amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminCollectFees(&_LaBr.TransactOpts, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) AdminCollectFees(amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminCollectFees(&_LaBr.TransactOpts, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_LaBr *LaBrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_LaBr *LaBrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _LaBr.Contract.AdminPauseTransfers(&_LaBr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_LaBr *LaBrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _LaBr.Contract.AdminPauseTransfers(&_LaBr.TransactOpts)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrTransactor) AdminSetBackendSrvAddress(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminSetBackendSrvAddress", newBackendSrv)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrSession) AdminSetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetBackendSrvAddress(&_LaBr.TransactOpts, newBackendSrv)
}

// AdminSetBackendSrvAddress is a paid mutator transaction binding the contract method 0x23328f39.
//
// Solidity: function adminSetBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrTransactorSession) AdminSetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetBackendSrvAddress(&_LaBr.TransactOpts, newBackendSrv)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetBurnable(&_LaBr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetBurnable(&_LaBr.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrTransactor) AdminSetRelayerHub(opts *bind.TransactOpts, newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminSetRelayerHub", newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetRelayerHub(&_LaBr.TransactOpts, newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrTransactorSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetRelayerHub(&_LaBr.TransactOpts, newRelayerHub)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetResource(&_LaBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.AdminSetResource(&_LaBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_LaBr *LaBrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_LaBr *LaBrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _LaBr.Contract.AdminUnpauseTransfers(&_LaBr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_LaBr *LaBrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _LaBr.Contract.AdminUnpauseTransfers(&_LaBr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminWithdraw(&_LaBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminWithdraw(&_LaBr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_LaBr *LaBrTransactor) CancelProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "cancelProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_LaBr *LaBrSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.Contract.CancelProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_LaBr *LaBrTransactorSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.Contract.CancelProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_LaBr *LaBrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_LaBr *LaBrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Deposit(&_LaBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_LaBr *LaBrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Deposit(&_LaBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ExecuteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0da380e5.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ExecuteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBackendSrvAddress) returns()
func (_LaBr *LaBrTransactor) Init(opts *bind.TransactOpts, initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "init", initBackendSrvAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBackendSrvAddress) returns()
func (_LaBr *LaBrSession) Init(initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Init(&_LaBr.TransactOpts, initBackendSrvAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address initBackendSrvAddress) returns()
func (_LaBr *LaBrTransactorSession) Init(initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Init(&_LaBr.TransactOpts, initBackendSrvAddress)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_LaBr *LaBrTransactor) RelayerCollectReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "relayerCollectReward")
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_LaBr *LaBrSession) RelayerCollectReward() (*types.Transaction, error) {
	return _LaBr.Contract.RelayerCollectReward(&_LaBr.TransactOpts)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_LaBr *LaBrTransactorSession) RelayerCollectReward() (*types.Transaction, error) {
	return _LaBr.Contract.RelayerCollectReward(&_LaBr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaBr *LaBrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaBr *LaBrSession) RenounceOwnership() (*types.Transaction, error) {
	return _LaBr.Contract.RenounceOwnership(&_LaBr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LaBr *LaBrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LaBr.Contract.RenounceOwnership(&_LaBr.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaBr *LaBrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaBr *LaBrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.TransferOwnership(&_LaBr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LaBr *LaBrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.TransferOwnership(&_LaBr.TransactOpts, newOwner)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrTransactor) VoteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "voteProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.VoteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x776860da.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.VoteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// LaBrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LaBr contract.
type LaBrDepositIterator struct {
	Event *LaBrDeposit // Event containing the contract specifics and raw log

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
func (it *LaBrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrDeposit)
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
		it.Event = new(LaBrDeposit)
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
func (it *LaBrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrDeposit represents a Deposit event raised by the LaBr contract.
type LaBrDeposit struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_LaBr *LaBrFilterer) FilterDeposit(opts *bind.FilterOpts) (*LaBrDepositIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &LaBrDepositIterator{contract: _LaBr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_LaBr *LaBrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LaBrDeposit) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrDeposit)
				if err := _LaBr.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_LaBr *LaBrFilterer) ParseDeposit(log types.Log) (*LaBrDeposit, error) {
	event := new(LaBrDeposit)
	if err := _LaBr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LaBr contract.
type LaBrOwnershipTransferredIterator struct {
	Event *LaBrOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LaBrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrOwnershipTransferred)
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
		it.Event = new(LaBrOwnershipTransferred)
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
func (it *LaBrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrOwnershipTransferred represents a OwnershipTransferred event raised by the LaBr contract.
type LaBrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaBr *LaBrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LaBrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LaBrOwnershipTransferredIterator{contract: _LaBr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LaBr *LaBrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LaBrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrOwnershipTransferred)
				if err := _LaBr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LaBr *LaBrFilterer) ParseOwnershipTransferred(log types.Log) (*LaBrOwnershipTransferred, error) {
	event := new(LaBrOwnershipTransferred)
	if err := _LaBr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LaBr contract.
type LaBrPausedIterator struct {
	Event *LaBrPaused // Event containing the contract specifics and raw log

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
func (it *LaBrPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrPaused)
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
		it.Event = new(LaBrPaused)
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
func (it *LaBrPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrPaused represents a Paused event raised by the LaBr contract.
type LaBrPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LaBr *LaBrFilterer) FilterPaused(opts *bind.FilterOpts) (*LaBrPausedIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LaBrPausedIterator{contract: _LaBr.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LaBr *LaBrFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LaBrPaused) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrPaused)
				if err := _LaBr.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LaBr *LaBrFilterer) ParsePaused(log types.Log) (*LaBrPaused, error) {
	event := new(LaBrPaused)
	if err := _LaBr.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the LaBr contract.
type LaBrProposalEventIterator struct {
	Event *LaBrProposalEvent // Event containing the contract specifics and raw log

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
func (it *LaBrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrProposalEvent)
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
		it.Event = new(LaBrProposalEvent)
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
func (it *LaBrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrProposalEvent represents a ProposalEvent event raised by the LaBr contract.
type LaBrProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_LaBr *LaBrFilterer) FilterProposalEvent(opts *bind.FilterOpts) (*LaBrProposalEventIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return &LaBrProposalEventIterator{contract: _LaBr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_LaBr *LaBrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *LaBrProposalEvent) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrProposalEvent)
				if err := _LaBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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

// ParseProposalEvent is a log parse operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_LaBr *LaBrFilterer) ParseProposalEvent(log types.Log) (*LaBrProposalEvent, error) {
	event := new(LaBrProposalEvent)
	if err := _LaBr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the LaBr contract.
type LaBrProposalVoteIterator struct {
	Event *LaBrProposalVote // Event containing the contract specifics and raw log

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
func (it *LaBrProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrProposalVote)
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
		it.Event = new(LaBrProposalVote)
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
func (it *LaBrProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrProposalVote represents a ProposalVote event raised by the LaBr contract.
type LaBrProposalVote struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_LaBr *LaBrFilterer) FilterProposalVote(opts *bind.FilterOpts) (*LaBrProposalVoteIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &LaBrProposalVoteIterator{contract: _LaBr.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_LaBr *LaBrFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *LaBrProposalVote) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrProposalVote)
				if err := _LaBr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_LaBr *LaBrFilterer) ParseProposalVote(log types.Log) (*LaBrProposalVote, error) {
	event := new(LaBrProposalVote)
	if err := _LaBr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the LaBr contract.
type LaBrRelayerThresholdChangedIterator struct {
	Event *LaBrRelayerThresholdChanged // Event containing the contract specifics and raw log

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
func (it *LaBrRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrRelayerThresholdChanged)
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
		it.Event = new(LaBrRelayerThresholdChanged)
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
func (it *LaBrRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the LaBr contract.
type LaBrRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_LaBr *LaBrFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts) (*LaBrRelayerThresholdChangedIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &LaBrRelayerThresholdChangedIterator{contract: _LaBr.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_LaBr *LaBrFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *LaBrRelayerThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrRelayerThresholdChanged)
				if err := _LaBr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
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

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_LaBr *LaBrFilterer) ParseRelayerThresholdChanged(log types.Log) (*LaBrRelayerThresholdChanged, error) {
	event := new(LaBrRelayerThresholdChanged)
	if err := _LaBr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrRewardCollectedIterator is returned from FilterRewardCollected and is used to iterate over the raw logs and unpacked data for RewardCollected events raised by the LaBr contract.
type LaBrRewardCollectedIterator struct {
	Event *LaBrRewardCollected // Event containing the contract specifics and raw log

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
func (it *LaBrRewardCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrRewardCollected)
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
		it.Event = new(LaBrRewardCollected)
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
func (it *LaBrRewardCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrRewardCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrRewardCollected represents a RewardCollected event raised by the LaBr contract.
type LaBrRewardCollected struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardCollected is a free log retrieval operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_LaBr *LaBrFilterer) FilterRewardCollected(opts *bind.FilterOpts) (*LaBrRewardCollectedIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return &LaBrRewardCollectedIterator{contract: _LaBr.contract, event: "RewardCollected", logs: logs, sub: sub}, nil
}

// WatchRewardCollected is a free log subscription operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_LaBr *LaBrFilterer) WatchRewardCollected(opts *bind.WatchOpts, sink chan<- *LaBrRewardCollected) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrRewardCollected)
				if err := _LaBr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
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

// ParseRewardCollected is a log parse operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_LaBr *LaBrFilterer) ParseRewardCollected(log types.Log) (*LaBrRewardCollected, error) {
	event := new(LaBrRewardCollected)
	if err := _LaBr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaBrUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LaBr contract.
type LaBrUnpausedIterator struct {
	Event *LaBrUnpaused // Event containing the contract specifics and raw log

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
func (it *LaBrUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrUnpaused)
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
		it.Event = new(LaBrUnpaused)
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
func (it *LaBrUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrUnpaused represents a Unpaused event raised by the LaBr contract.
type LaBrUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LaBr *LaBrFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LaBrUnpausedIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LaBrUnpausedIterator{contract: _LaBr.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LaBr *LaBrFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LaBrUnpaused) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrUnpaused)
				if err := _LaBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LaBr *LaBrFilterer) ParseUnpaused(log types.Log) (*LaBrUnpaused, error) {
	event := new(LaBrUnpaused)
	if err := _LaBr.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}