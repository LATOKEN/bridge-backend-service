package rlr

import (
	"sync"
	"time"

	watcher "github.com/latoken/bridge-backend-service/src/service/blockchains-watcher"
	fetcher "github.com/latoken/bridge-backend-service/src/service/gas-price-fetcher"
	"github.com/latoken/bridge-backend-service/src/service/storage"
	workers "github.com/latoken/bridge-backend-service/src/service/workers"
	"github.com/latoken/bridge-backend-service/src/service/workers/eth-compatible"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/latoken/bridge-backend-service/src/models"
)

// BridgeSRV ...
type BridgeSRV struct {
	sync.RWMutex
	logger   *logrus.Logger
	Watcher  *watcher.WatcherSRV
	Fetcher  *fetcher.FetcherSrv
	laWorker workers.IWorker
	Workers  map[string]workers.IWorker
	storage  *storage.DataBase
}

// CreateNewBridgeSRV ...
func CreateNewBridgeSRV(logger *logrus.Logger, gormDB *gorm.DB, laConfig *models.WorkerConfig, chainCfgs []*models.WorkerConfig,
	chainFetCfgs []*models.FetcherConfig, resourceIDs []*storage.ResourceId) *BridgeSRV {
	// init database
	db, err := storage.InitStorage(gormDB)
	if err != nil {
		logger.Fatalf("Connect to DataBase: ", err)
	}

	// create Relayer instance
	inst := BridgeSRV{
		logger:   logger,
		storage:  db,
		laWorker: eth.NewErc20Worker(logger, laConfig, db),
		Workers:  make(map[string]workers.IWorker),
	}
	// create erc20 worker
	for _, cfg := range chainCfgs {
		inst.Workers[cfg.ChainName] = eth.NewErc20Worker(logger, cfg, db)
	}
	// create la worker
	inst.Workers["LA"] = inst.laWorker

	// check rules for workers(>=2, different chainIDs...)
	if len(inst.Workers) < 1 {
		logger.Fatalf("Num of workers must be > 1, but = %d", len(inst.Workers))
		return nil
	}
	inst.Watcher = watcher.CreateNewWatcherSRV(logger, db, inst.Workers)
	inst.Fetcher = fetcher.CreateFetcherSrv(logger, db, chainFetCfgs)

	db.SaveResourceIDs(resourceIDs)
	return &inst
}

// !!! TODO !!!

// Run ...
func (r *BridgeSRV) Run() {
	// start watcher
	r.Watcher.Run()
	//start fetcher
	r.Fetcher.Run()
	go r.UpdateTxOnLachain()
	// run Worker workers
	for _, worker := range r.Workers {
		go r.ConfirmWorkerTx(worker)
		go r.emitProposal(worker)
		go r.CheckTxSentRoutine(worker)
	}
}

// ConfirmWorkerTx ...
func (r *BridgeSRV) ConfirmWorkerTx(worker workers.IWorker) {
	for {
		txLogs, err := r.storage.FindTxLogs(worker.GetChainName(), worker.GetConfirmNum())
		if err != nil {
			r.logger.Errorf("ConfirmWorkerTx(), err = %s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		txHashes := make([]string, 0, len(txLogs))
		newEvents := make([]*storage.Event, 0)

		for _, txLog := range txLogs {
			// reject swap request if receiver addr and worker chain addr both are r addr
			// if worker.IsSameAddress(txLog.ReceiverAddr, worker.GetWorkerAddress()) &&
			// 	!r.laWorker.IsSameAddress(txLog.WorkerChainAddr, r.laWorker.GetWorkerAddress()) {
			// 	r.logger.Warnln("THE SAME")
			// }
			r.logger.Infoln("New Event")
			newEvent := &storage.Event{
				ReceiverAddr:       txLog.ReceiverAddr,
				DepositNonce:       txLog.DepositNonce,
				ResourceID:         txLog.ResourceID,
				ChainID:            txLog.Chain,
				DestinationChainID: txLog.DestinationChainID,
				OriginChainID:      txLog.OriginChainID,
				InAmount:           txLog.InAmount,
				OutAmount:          txLog.OutAmount,
				Height:             txLog.Height,
				SwapID:             txLog.SwapID,
				Status:             txLog.EventStatus,
				CreateTime:         time.Now().Unix(),
				Params:             txLog.Params,
			}
			newEvents = append(newEvents, newEvent)
			txHashes = append(txHashes, txLog.TxHash)
		}

		//
		if err := r.storage.ConfirmWorkerTx(worker.GetChainName(), txLogs, txHashes, newEvents); err != nil {
			r.logger.Errorf("compensate new swap tx error, err=%s", err)
		}

		time.Sleep(2 * time.Second)
	}
}

// CheckTxSentRoutine ...
func (r *BridgeSRV) CheckTxSentRoutine(worker workers.IWorker) {
	for {
		r.CheckTxSent(worker)
		time.Sleep(time.Second)
	}
}

// CheckTxSent ...
func (r *BridgeSRV) CheckTxSent(worker workers.IWorker) {
	txsSent, err := r.storage.GetTxsSentByStatus(worker.GetChainName())
	if err != nil {
		r.logger.WithFields(logrus.Fields{"function": "CheckTxSent() | GetTxsSentByStatus()"}).Errorln(err)
		return
	}

	for _, txSent := range txsSent {
		// Get status of tx from chain
		status := worker.GetSentTxStatus(txSent.TxHash, txSent.Nonce)
		if err := r.storage.UpdateTxSentStatus(txSent, status); err != nil {
			r.logger.WithFields(logrus.Fields{"function": "CheckTxSent() | UpdateTxSentStatus()"}).Errorln(err)
			return
		}
	}
}

func (r *BridgeSRV) handleTxSent(chain string, event *storage.Event, txType storage.TxType, backwardStatus storage.EventStatus,
	failedStatus storage.EventStatus, successStatus storage.EventStatus) {
	txsSent := r.storage.GetTxsSentByType(chain, txType, event)

	if len(txsSent) == 0 {
		r.storage.UpdateEventStatus(event, backwardStatus)
		return
	}
	latestTx := txsSent[0]
	timeElapsed := time.Now().Unix() - latestTx.CreateTime
	autoRetryTimeout, autoRetryNum := r.getAutoRetryConfig(chain)
	txStatus := latestTx.Status

	if txStatus == storage.TxSentStatusSuccess {
		r.storage.UpdateEventStatus(event, successStatus)
		return
	}
	if txStatus == storage.TxSentStatusFailed {
		r.storage.UpdateEventStatus(event, failedStatus)
		return
	}

	if timeElapsed > autoRetryTimeout && txStatus == storage.TxSentStatusInit {

		if len(txsSent) >= autoRetryNum {
			r.storage.UpdateEventStatus(event, failedStatus)
		} else {
			r.storage.UpdateEventStatus(event, backwardStatus)
		}
		r.storage.UpdateTxSentStatus(latestTx, storage.TxSentStatusLost)
	}
}

// !!! TODO !!!

func (r *BridgeSRV) getAutoRetryConfig(chain string) (int64, int) {
	return 1800, 1
}
