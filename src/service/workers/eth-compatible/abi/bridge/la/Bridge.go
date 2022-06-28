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
	ABI: "[{\"name\":\"paused\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"sub\",\"type\":\"function\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"name\":\"_isInitialised\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"ownableInit\",\"type\":\"function\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"transferOwnership\",\"type\":\"function\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"verify\",\"type\":\"function\",\"inputs\":[{\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"name\":\"getSigner\",\"type\":\"function\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"name\":\"createMesssageHash\",\"type\":\"function\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"name\":\"implementation\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"proxyOwner\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_chainID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerThreshold\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_fee\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_expiry\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"name\":\"_relayerHubAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_backendSrvAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_balancerAddress\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_nativeResourceID\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"name\":\"_resourceIDToHandlerAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"name\":\"_depositRecords\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"name\":\"_proposals\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"extraLATransferred\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"initialize\",\"type\":\"function\",\"inputs\":[{\"name\":\"chainID_\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"relayerThreshold_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initBackendSrvAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initBalancerAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setRelayerHub\",\"type\":\"function\",\"inputs\":[{\"name\":\"newRelayerHub\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setBackendSrvAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBackendSrv\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setBalancerAddress\",\"type\":\"function\",\"inputs\":[{\"name\":\"newBalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminPauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminUnpauseTransfers\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"changeRelayerThreshold\",\"type\":\"function\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setResource\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setBurnable\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"setNativeResourceID\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBridge.Proposal\",\"components\":[{\"name\":\"_resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_yesVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_noVotes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"_proposedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"name\":\"changeFee\",\"type\":\"function\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminWithdraw\",\"type\":\"function\",\"inputs\":[{\"name\":\"handlerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountOrTokenID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"deposit\",\"type\":\"function\",\"inputs\":[{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"voteProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"cancelProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"executeProposal\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"updateExternalTx\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"proposalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountToTransfer\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"adminCollectFees\",\"type\":\"function\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"backendSrvCollectFees\",\"type\":\"function\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"approveSpending\",\"type\":\"function\",\"inputs\":[{\"name\":\"resourceIDOwner\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resourceIDSpender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountOrTokenID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"relayerCollectReward\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"getExtraLATransferred\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"name\":\"transferExtraFee\",\"type\":\"function\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"name\":\"depositFunds\",\"type\":\"function\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"name\":\"Paused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"Unpaused\",\"type\":\"event\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"anonymous\":false},{\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"name\":\"RelayerThresholdChanged\",\"type\":\"event\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"Deposit\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalEvent\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"recipientAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"ProposalVote\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBridge.ProposalStatus\"},{\"name\":\"resourceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"name\":\"RewardCollected\",\"type\":\"event\",\"inputs\":[{\"name\":\"relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false},{\"name\":\"ExtraFeeTransferred\",\"type\":\"event\",\"inputs\":[{\"name\":\"originChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"destinationChainID\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"},{\"name\":\"depositNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"resouceID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"anonymous\":false}]",
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

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_LaBr *LaBrCaller) BalancerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_balancerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_LaBr *LaBrSession) BalancerAddress() (common.Address, error) {
	return _LaBr.Contract.BalancerAddress(&_LaBr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_LaBr *LaBrCallerSession) BalancerAddress() (common.Address, error) {
	return _LaBr.Contract.BalancerAddress(&_LaBr.CallOpts)
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

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaBr *LaBrCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaBr *LaBrSession) IsInitialised() (bool, error) {
	return _LaBr.Contract.IsInitialised(&_LaBr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_LaBr *LaBrCallerSession) IsInitialised() (bool, error) {
	return _LaBr.Contract.IsInitialised(&_LaBr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_LaBr *LaBrCaller) NativeResourceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "_nativeResourceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_LaBr *LaBrSession) NativeResourceID() ([32]byte, error) {
	return _LaBr.Contract.NativeResourceID(&_LaBr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_LaBr *LaBrCallerSession) NativeResourceID() ([32]byte, error) {
	return _LaBr.Contract.NativeResourceID(&_LaBr.CallOpts)
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

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_LaBr *LaBrCaller) CreateMesssageHash(opts *bind.CallOpts, amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "createMesssageHash", amount, recipient, chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_LaBr *LaBrSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _LaBr.Contract.CreateMesssageHash(&_LaBr.CallOpts, amount, recipient, chainId)
}

// CreateMesssageHash is a free data retrieval call binding the contract method 0xc79564f1.
//
// Solidity: function createMesssageHash(uint256 amount, address recipient, bytes8 chainId) pure returns(bytes32)
func (_LaBr *LaBrCallerSession) CreateMesssageHash(amount *big.Int, recipient common.Address, chainId [8]byte) ([32]byte, error) {
	return _LaBr.Contract.CreateMesssageHash(&_LaBr.CallOpts, amount, recipient, chainId)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_LaBr *LaBrCaller) ExtraLATransferred(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "extraLATransferred", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_LaBr *LaBrSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _LaBr.Contract.ExtraLATransferred(&_LaBr.CallOpts, arg0, arg1)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_LaBr *LaBrCallerSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _LaBr.Contract.ExtraLATransferred(&_LaBr.CallOpts, arg0, arg1)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_LaBr *LaBrCaller) GetExtraLATransferred(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "getExtraLATransferred", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_LaBr *LaBrSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	return _LaBr.Contract.GetExtraLATransferred(&_LaBr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns(bool)
func (_LaBr *LaBrCallerSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (bool, error) {
	return _LaBr.Contract.GetExtraLATransferred(&_LaBr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCaller) GetProposal(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "getProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.GetProposal(&_LaBr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_LaBr *LaBrCallerSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (BridgeProposal, error) {
	return _LaBr.Contract.GetProposal(&_LaBr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_LaBr *LaBrCaller) GetSigner(opts *bind.CallOpts, messageHash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "getSigner", messageHash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_LaBr *LaBrSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _LaBr.Contract.GetSigner(&_LaBr.CallOpts, messageHash, signature)
}

// GetSigner is a free data retrieval call binding the contract method 0xf7b2ec0d.
//
// Solidity: function getSigner(bytes32 messageHash, bytes signature) pure returns(address)
func (_LaBr *LaBrCallerSession) GetSigner(messageHash [32]byte, signature []byte) (common.Address, error) {
	return _LaBr.Contract.GetSigner(&_LaBr.CallOpts, messageHash, signature)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaBr *LaBrCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaBr *LaBrSession) Implementation() (common.Address, error) {
	return _LaBr.Contract.Implementation(&_LaBr.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LaBr *LaBrCallerSession) Implementation() (common.Address, error) {
	return _LaBr.Contract.Implementation(&_LaBr.CallOpts)
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

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaBr *LaBrCaller) ProxyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "proxyOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaBr *LaBrSession) ProxyOwner() (common.Address, error) {
	return _LaBr.Contract.ProxyOwner(&_LaBr.CallOpts)
}

// ProxyOwner is a free data retrieval call binding the contract method 0x025313a2.
//
// Solidity: function proxyOwner() view returns(address)
func (_LaBr *LaBrCallerSession) ProxyOwner() (common.Address, error) {
	return _LaBr.Contract.ProxyOwner(&_LaBr.CallOpts)
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

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_LaBr *LaBrCaller) Verify(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	var out []interface{}
	err := _LaBr.contract.Call(opts, &out, "verify", _ethSignedMessageHash, _signer, _signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_LaBr *LaBrSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _LaBr.Contract.Verify(&_LaBr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xf5cf2ce1.
//
// Solidity: function verify(bytes32 _ethSignedMessageHash, address _signer, bytes _signature) pure returns(bool)
func (_LaBr *LaBrCallerSession) Verify(_ethSignedMessageHash [32]byte, _signer common.Address, _signature []byte) (bool, error) {
	return _LaBr.Contract.Verify(&_LaBr.CallOpts, _ethSignedMessageHash, _signer, _signature)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactor) AdminCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "adminCollectFees", recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminCollectFees(&_LaBr.TransactOpts, recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.AdminCollectFees(&_LaBr.TransactOpts, recipient, amount)
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

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrTransactor) ApproveSpending(opts *bind.TransactOpts, resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "approveSpending", resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ApproveSpending(&_LaBr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_LaBr *LaBrTransactorSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ApproveSpending(&_LaBr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactor) BackendSrvCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "backendSrvCollectFees", recipient, amount)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrSession) BackendSrvCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.BackendSrvCollectFees(&_LaBr.TransactOpts, recipient, amount)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) BackendSrvCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.BackendSrvCollectFees(&_LaBr.TransactOpts, recipient, amount)
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

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_LaBr *LaBrTransactor) ChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "changeFee", newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_LaBr *LaBrSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ChangeFee(&_LaBr.TransactOpts, newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_LaBr *LaBrTransactorSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ChangeFee(&_LaBr.TransactOpts, newFee)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrTransactor) ChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "changeRelayerThreshold", newThreshold)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrSession) ChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ChangeRelayerThreshold(&_LaBr.TransactOpts, newThreshold)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_LaBr *LaBrTransactorSession) ChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.ChangeRelayerThreshold(&_LaBr.TransactOpts, newThreshold)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_LaBr *LaBrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_LaBr *LaBrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _LaBr.Contract.Deposit(&_LaBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_LaBr *LaBrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _LaBr.Contract.Deposit(&_LaBr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaBr *LaBrTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaBr *LaBrSession) DepositFunds() (*types.Transaction, error) {
	return _LaBr.Contract.DepositFunds(&_LaBr.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_LaBr *LaBrTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _LaBr.Contract.DepositFunds(&_LaBr.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_LaBr *LaBrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_LaBr *LaBrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaBr.Contract.ExecuteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_LaBr *LaBrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _LaBr.Contract.ExecuteProposal(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_LaBr *LaBrTransactor) Initialize(opts *bind.TransactOpts, chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "initialize", chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_LaBr *LaBrSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Initialize(&_LaBr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_LaBr *LaBrTransactorSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.Initialize(&_LaBr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaBr *LaBrTransactor) OwnableInit(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "ownableInit", owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaBr *LaBrSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.OwnableInit(&_LaBr.TransactOpts, owner_)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner_) returns()
func (_LaBr *LaBrTransactorSession) OwnableInit(owner_ common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.OwnableInit(&_LaBr.TransactOpts, owner_)
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

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrTransactor) SetBackendSrvAddress(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setBackendSrvAddress", newBackendSrv)
}

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrSession) SetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBackendSrvAddress(&_LaBr.TransactOpts, newBackendSrv)
}

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_LaBr *LaBrTransactorSession) SetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBackendSrvAddress(&_LaBr.TransactOpts, newBackendSrv)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_LaBr *LaBrTransactor) SetBalancerAddress(opts *bind.TransactOpts, newBalancer common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setBalancerAddress", newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_LaBr *LaBrSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBalancerAddress(&_LaBr.TransactOpts, newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_LaBr *LaBrTransactorSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBalancerAddress(&_LaBr.TransactOpts, newBalancer)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrTransactor) SetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setBurnable", handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBurnable(&_LaBr.TransactOpts, handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_LaBr *LaBrTransactorSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetBurnable(&_LaBr.TransactOpts, handlerAddress, tokenAddress)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_LaBr *LaBrTransactor) SetNativeResourceID(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setNativeResourceID", resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_LaBr *LaBrSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.Contract.SetNativeResourceID(&_LaBr.TransactOpts, resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_LaBr *LaBrTransactorSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _LaBr.Contract.SetNativeResourceID(&_LaBr.TransactOpts, resourceID)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrTransactor) SetRelayerHub(opts *bind.TransactOpts, newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setRelayerHub", newRelayerHub)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrSession) SetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetRelayerHub(&_LaBr.TransactOpts, newRelayerHub)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_LaBr *LaBrTransactorSession) SetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetRelayerHub(&_LaBr.TransactOpts, newRelayerHub)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrTransactor) SetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "setResource", handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetResource(&_LaBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_LaBr *LaBrTransactorSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _LaBr.Contract.SetResource(&_LaBr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactor) TransferExtraFee(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "transferExtraFee", originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_LaBr *LaBrSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.TransferExtraFee(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xc2a43f17.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) returns()
func (_LaBr *LaBrTransactorSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaBr.Contract.TransferExtraFee(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount)
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

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_LaBr *LaBrTransactor) UpdateExternalTx(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _LaBr.contract.Transact(opts, "updateExternalTx", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
}

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_LaBr *LaBrSession) UpdateExternalTx(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _LaBr.Contract.UpdateExternalTx(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
}

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_LaBr *LaBrTransactorSession) UpdateExternalTx(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _LaBr.Contract.UpdateExternalTx(&_LaBr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
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

// LaBrExtraFeeTransferredIterator is returned from FilterExtraFeeTransferred and is used to iterate over the raw logs and unpacked data for ExtraFeeTransferred events raised by the LaBr contract.
type LaBrExtraFeeTransferredIterator struct {
	Event *LaBrExtraFeeTransferred // Event containing the contract specifics and raw log

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
func (it *LaBrExtraFeeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaBrExtraFeeTransferred)
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
		it.Event = new(LaBrExtraFeeTransferred)
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
func (it *LaBrExtraFeeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaBrExtraFeeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaBrExtraFeeTransferred represents a ExtraFeeTransferred event raised by the LaBr contract.
type LaBrExtraFeeTransferred struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResouceID          [32]byte
	Recipient          common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtraFeeTransferred is a free log retrieval operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_LaBr *LaBrFilterer) FilterExtraFeeTransferred(opts *bind.FilterOpts) (*LaBrExtraFeeTransferredIterator, error) {

	logs, sub, err := _LaBr.contract.FilterLogs(opts, "ExtraFeeTransferred")
	if err != nil {
		return nil, err
	}
	return &LaBrExtraFeeTransferredIterator{contract: _LaBr.contract, event: "ExtraFeeTransferred", logs: logs, sub: sub}, nil
}

// WatchExtraFeeTransferred is a free log subscription operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_LaBr *LaBrFilterer) WatchExtraFeeTransferred(opts *bind.WatchOpts, sink chan<- *LaBrExtraFeeTransferred) (event.Subscription, error) {

	logs, sub, err := _LaBr.contract.WatchLogs(opts, "ExtraFeeTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaBrExtraFeeTransferred)
				if err := _LaBr.contract.UnpackLog(event, "ExtraFeeTransferred", log); err != nil {
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

// ParseExtraFeeTransferred is a log parse operation binding the contract event 0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd.
//
// Solidity: event ExtraFeeTransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resouceID, address recipient, uint256 amount)
func (_LaBr *LaBrFilterer) ParseExtraFeeTransferred(log types.Log) (*LaBrExtraFeeTransferred, error) {
	event := new(LaBrExtraFeeTransferred)
	if err := _LaBr.contract.UnpackLog(event, "ExtraFeeTransferred", log); err != nil {
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
