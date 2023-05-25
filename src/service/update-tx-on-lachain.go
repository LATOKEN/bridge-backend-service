package rlr

import (
	"fmt"
	"math/big"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"
)

// Updates withdraw swap status on lachain
func (b *BridgeSRV) UpdateTxOnLachain() {
	for {
		events := b.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusPassedFailed, storage.EventStatusPassedConfirmed})
		for _, event := range events {
			b.logger.Infoln("attempting to send confirmation tx")
			txHash, err := b.SendConfirmationLA(event)
			if err != nil {
				b.logger.Errorf("confirmation failed: %s | txHash: %s", err, txHash)
			}
			b.logger.Infoln("confirmation tx success")
		}
		time.Sleep(time.Minute)
	}
}

func (b *BridgeSRV) SendConfirmationLA(event *storage.Event) (string, error) {
	txSent := &storage.TxSent{
		Chain:      "LA",
		Type:       storage.TxTypeUpdate,
		SwapID:     event.SwapID,
		CreateTime: time.Now().Unix(),
	}
	var status uint8 = 3
	if event.Status == storage.EventStatusPassedFailed {
		status = 4
	}

	//no need to update on chain for deposit to lachain tx
	if event.DestinationChainID == b.laWorker.GetDestinationID() {
		if status == 3 {
			b.storage.UpdateEventStatus(event, storage.EventStatusUpdateConfirmed)
		} else {
			b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		}
		return "", nil
	}

	b.logger.Infof("Update status parameters:  depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) | inAmount(%s) \n",
		event.DepositNonce, event.ReceiverAddr, event.OutAmount, event.ResourceID, event.InAmount)

	if event.InAmount == "" || event.OutAmount == "" {
		err := fmt.Errorf("Error in finding amounts")
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.CreateTxSent(txSent)
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		return "", err
	}

	inAmount, _ := new(big.Int).SetString(event.InAmount, 10)
	outAmount, _ := new(big.Int).SetString(event.OutAmount, 10)

	txHash, nonce, err := b.laWorker.UpdateSwapStatusOnChain(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
		event.ReceiverAddr, outAmount, inAmount, utils.StringToBytes(event.Params), status)
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.Status = storage.TxSentStatusFailed
		b.storage.CreateTxSent(txSent)
		b.storage.UpdateEventStatus(event, storage.EventStatusUpdateFailed)
		return "", err
	}
	txSent.TxHash = txHash
	txSent.Nonce = nonce
	b.logger.Infof("Update status tx success: %s", txHash)
	b.storage.UpdateEventStatus(event, storage.EventStatusUpdateConfirmed)
	b.storage.CreateTxSent(txSent)
	return txHash, nil
}
