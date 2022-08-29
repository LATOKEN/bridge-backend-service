package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	ethBr "github.com/latoken/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "github.com/latoken/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/la"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ProposalEventName = "ProposalEvent"
	DepositEventName  = "Deposit"
)

// Event Hash (SHA3)
var (
	ProposalEventHash = common.HexToHash("0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d")
	DepositEventHash  = common.HexToHash("0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f")
)

var txStatus = make(map[string]uint8)

// ProposalEvent represents a ProposalEvent event raised by the Bridge.sol contract.
type ProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

type DepositEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Params             [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// monitors swap status till 5 mins then throw error
func setTxMonitor(SwapID string, Status uint8) {
	if Status == 1 {
		txStatus[SwapID] = Status
		go func(SwapID string, Status uint8) {
			time.Sleep(5 * 60 * time.Second)
			if NewStatus, ok := txStatus[SwapID]; ok {
				if NewStatus < 3 {
					fmt.Printf("ERROR[%s] SwapID %s stuck in status %d\n\n", time.Now().Format(time.RFC3339Nano), SwapID, NewStatus)
				}
				delete(txStatus, SwapID)
			}
		}(SwapID, Status)
	} else {
		if HashStatus, ok := txStatus[SwapID]; ok {
			if HashStatus < Status {
				txStatus[SwapID] = Status
			}
		}
	}
}

func ParseLAProposalEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ProposalEvent
	if err := abi.UnpackIntoInterface(&ev, ProposalEventName, log.Data); err != nil {
		return nil, err
	}

	var event_time = time.Now().Format(time.RFC3339Nano)
	var SwapID = ev.CalcutateSwapID()
	fmt.Printf("INFO[%s] SwapID: %s\n", event_time, SwapID)
	fmt.Printf("INFO[%s] ProposalEvent\n", event_time)
	fmt.Printf("INFO[%s] origin chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("INFO[%s] destination chain ID: 0x%s\n", event_time, common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("INFO[%s] deposit nonce: %d\n", event_time, ev.DepositNonce)
	fmt.Printf("INFO[%s] status: %d\n", event_time, ev.Status)
	fmt.Printf("INFO[%s] resource ID: 0x%s\n", event_time, common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("INFO[%s] recipient: 0x%s\n", event_time, common.Bytes2Hex(ev.RecipientAddress[:]))
	fmt.Printf("INFO[%s] amount: %s\n\n", event_time, ev.Amount.String())

	setTxMonitor(SwapID, ev.Status)

	return ev, nil
}

// ParseLaDepositEvent ...
func ParseLaDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(laBr.LaBrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	var SwapID = utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))
	fmt.Printf("[%s] Deposited\n", SwapID)
	fmt.Printf("[%s] destination chain ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("[%s] resource ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("[%s] deposit nonce: %d\n", SwapID, ev.DepositNonce)
	fmt.Printf("[%s] depositor address: %s\n", SwapID, ev.Depositor.Hex())
	fmt.Printf("[%s] recipient address: %s\n", SwapID, ev.RecipientAddress.Hex())
	fmt.Printf("[%s] token address: %s\n", SwapID, ev.TokenAddress.Hex())
	fmt.Printf("[%s] amount : %s\n", SwapID, ev.Amount.String())
	fmt.Printf("[%s] Params : %s\n", SwapID, common.Bytes2Hex(ev.Params[:]))

	return ev, nil
}

// ParseDepositEvent ...
func ParseEthDepositEvent(log *types.Log) (ContractEvent, error) {
	var ev DepositEvent
	abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
	if err := abi.UnpackIntoInterface(&ev, DepositEventName, log.Data); err != nil {
		return nil, err
	}

	ev.DestinationChainID = utils.BytesToBytes8(log.Topics[1].Bytes())
	ev.ResourceID = utils.BytesToBytes32(log.Topics[2].Bytes())
	ev.DepositNonce = big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64()

	var SwapID = utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))
	fmt.Printf("[%s] Deposited\n", SwapID)
	fmt.Printf("[%s] destination chain ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("[%s] resource ID: 0x%s\n", SwapID, common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("[%s] deposit nonce: %d\n", SwapID, ev.DepositNonce)
	fmt.Printf("[%s] depositor address: %s\n", SwapID, ev.Depositor.Hex())
	fmt.Printf("[%s] recipient address: %s\n", SwapID, ev.RecipientAddress.Hex())
	fmt.Printf("[%s] token address: %s\n", SwapID, ev.TokenAddress.Hex())
	fmt.Printf("[%s] amount : %s\n", SwapID, ev.Amount.String())
	fmt.Printf("[%s] Params : %s\n", SwapID, common.Bytes2Hex(ev.Params[:]))

	return ev, nil
}

// !!! TODO !!!
func (ev ProposalEvent) CalcutateSwapID() string {
	return utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce))
}

// ToTxLog ...
func (ev ProposalEvent) ToTxLog(chain string) *storage.TxLog {
	// if status == 2 -> already claimed -> mint
	// if status == 3-> already minted(executed)
	txlog := &storage.TxLog{
		// Chain:              chain,
		TxType:             storage.TxTypeClaim,
		ReceiverAddr:       ev.RecipientAddress.String(),
		OutAmount:          ev.Amount.String(),
		SwapID:             ev.CalcutateSwapID(),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginChainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		DepositNonce:       ev.DepositNonce,
		SwapStatus:         ev.Status,
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		EventStatus:        storage.EventStatusClaimConfirmed,
	}

	if ev.Status == uint8(2) {
		txlog.TxType = storage.TxTypePassed
		txlog.EventStatus = storage.EventStatusPassedInit
	} else if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeSpend
		txlog.EventStatus = storage.EventStatusSpendConfirmed
	} else if ev.Status == uint8(4) {
		txlog.TxType = storage.TxTypeExpired
		txlog.EventStatus = storage.EventStatusExpiredConfirmed
	}

	return txlog
}

// ToTxLog ...
func (ev DepositEvent) ToTxLog(chain string) *storage.TxLog {
	return &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeDeposit,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginChainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		SwapID:             utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		DepositNonce:       ev.DepositNonce,
		SenderAddr:         ev.Depositor.Hex(),
		ReceiverAddr:       ev.RecipientAddress.Hex(),
		InAmount:           ev.Amount.String(),
		Params:             common.Bytes2Hex(ev.Params[:]),

	}
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ProposalEventHash[:]) {
		if w.GetChainName() == "LA" {
			abi, _ := abi.JSON(strings.NewReader(laBr.LaBrABI))
			return ParseLAProposalEvent(&abi, log)
		}
	}
	if bytes.Equal(log.Topics[0][:], DepositEventHash[:]) {
		if w.chainName == "LA" {
			return ParseLaDepositEvent(log)
		} else {
			return ParseEthDepositEvent(log)
		}
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
