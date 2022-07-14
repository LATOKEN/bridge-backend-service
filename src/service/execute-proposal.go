package rlr

import (
	"fmt"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	workers "github.com/latoken/bridge-backend-service/src/service/workers"
	"github.com/latoken/bridge-backend-service/src/service/workers/utils"
)

// emitRegistreted ...
func (r *BridgeSRV) emitProposal(worker workers.IWorker) {
	for {
		events := r.storage.GetEventsByTypeAndStatuses([]storage.EventStatus{storage.EventStatusPassedInit, storage.EventStatusPassedInitConfrimed, storage.EventStatusPassedSentFailed})
		for _, event := range events {
			if event.Status == storage.EventStatusPassedInitConfrimed &&
				worker.GetDestinationID() == event.DestinationChainID { //send tx where dest chainID matches
				r.logger.Infoln("attempting to send execute proposal")
				if _, err := r.sendExecuteProposal(worker, event); err != nil {
					r.logger.Errorf("submit claim failed: %s", err)
				}
			} else {
				r.handleTxSent(worker.GetChainName(), event, storage.TxTypePassed,
					storage.EventStatusPassedInitConfrimed, storage.EventStatusPassedFailed)
			}
			time.Sleep(2 * time.Second)
		}
		time.Sleep(10 * time.Second)
	}
}

// ethSendClaim ...
func (r *BridgeSRV) sendExecuteProposal(worker workers.IWorker, event *storage.Event) (txHash string, err error) {
	txSent := &storage.TxSent{
		Chain:      worker.GetChainName(),
		Type:       storage.TxTypePassed,
		SwapID:     event.SwapID,
		CreateTime: time.Now().Unix(),
	}

	r.logger.Infof("Execute parameters:  depositNonce(%d) | sender(%s) | outAmount(%s) | resourceID(%s) | chainID(%s)\n",
		event.DepositNonce, event.ReceiverAddr, event.OutAmount, event.ResourceID, worker.GetChainName())
	if worker.GetChainName() == "LA" {
		// to update liquidity index inside lachain for aave tokens
		if event.ResourceID == r.storage.FetchResourceIDByName("amToken").ID {
			wor := r.Workers["POS"]
			liquidity, _ := wor.GetLiquidityIndex(wor.GetConfig().AmTokenHandlerAddress, wor.GetConfig().AMUSDTContractAddr)
			txHash, err = worker.ExecuteProposalLa(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
				event.ReceiverAddr, event.OutAmount, liquidity)
		} else {
			txHash, err = worker.ExecuteProposalLa(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
				event.ReceiverAddr, event.OutAmount, nil)
		}
	} else {
		txHash, err = worker.ExecuteProposalEth(event.DepositNonce, utils.StringToBytes8(event.OriginChainID), utils.StringToBytes8(event.DestinationChainID), utils.StringToBytes32(event.ResourceID),
			event.ReceiverAddr, event.OutAmount)
	}
	if err != nil {
		txSent.ErrMsg = err.Error()
		txSent.TxHash = txHash
		r.storage.UpdateEventStatus(event, storage.EventStatusPassedSentFailed)
		r.storage.CreateTxSent(txSent)
		return "", fmt.Errorf("could not send claim tx: %w", err)
	}
	txSent.TxHash = txHash
	r.storage.UpdateEventStatus(event, storage.EventStatusPassedSent)
	r.logger.Infof("send execute proposal tx success | chain=%s, tx_hash=%s", worker.GetChainName(), txSent.TxHash)
	// create new tx(claimed)
	err = r.storage.CreateTxSent(txSent)
	if err != nil {
		println("db txsent err", err)
	}

	return txSent.TxHash, nil

}
