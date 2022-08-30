package watcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/latoken/bridge-backend-service/src/service/storage"
	"github.com/latoken/bridge-backend-service/src/service/workers"

	"github.com/sirupsen/logrus"
)

// WatcherSRV ...
type WatcherSRV struct {
	logger  *logrus.Entry
	storage *storage.DataBase
	Workers map[string]workers.IWorker
}

// CreateNewWatcherSRV ...
func CreateNewWatcherSRV(logger *logrus.Logger, db *storage.DataBase, workers map[string]workers.IWorker) *WatcherSRV {
	return &WatcherSRV{
		logger:  logger.WithField("layer", "watcher"),
		storage: db,
		Workers: workers,
	}
}

// Run ...
func (w *WatcherSRV) Run() {
	for name, worker := range w.Workers {
		w.logger.Infoln("watcher | worker name: ", name)
		startHeight, err := worker.GetStartHeight()
		if err != nil {
			w.logger.Fatalf("err = %v", err)
			return
		}
		go w.collector(worker, worker.GetFetchInterval(), startHeight)
		time.Sleep(100 * time.Millisecond)
	}

}

func (w *WatcherSRV) collector(worker workers.IWorker, threshold time.Duration, startHeight int64) {
	for {
		curBlockLog := w.storage.GetCurrentBlockLog(worker.GetChainName())
		if curBlockLog.Height == 0 {
			w.logger.Warnf("%s current height: %d", worker.GetChainName(), curBlockLog.Height)
		} else {
			w.logger.Infof("%s current height: %d", worker.GetChainName(), curBlockLog.Height)
		}

		height := curBlockLog.Height
		if curBlockLog.Height == 0 && startHeight != 0 {
			height = startHeight
		}

		if err := w.getBlock(worker, height, curBlockLog.BlockHash); err != nil {
			normalizedErr := strings.ToLower(err.Error())
			if strings.Contains(normalizedErr, "height must be less than or equal to the current blockchain height") ||
				strings.Contains(normalizedErr, "not found") ||
				strings.Contains(normalizedErr, "block number out of range") {
				w.logger.Infof("try to get ahead block, chain=%s, height=%d", worker.GetChainName(), height)
			} else {
				w.logger.Error(normalizedErr)
			}
			time.Sleep(threshold)
		}

		// TODO
		time.Sleep(time.Second)
	}
}

func (w *WatcherSRV) getBlock(worker workers.IWorker, curHeight int64, curBlockHash string) error {
	blockAndTxLogs, err := worker.GetBlockAndTxs(curHeight)
	if err != nil {
		return fmt.Errorf("get %s block info error, height reached =%d, err=%s", worker.GetChainName(), curHeight, err.Error())
	}

	parentHash := blockAndTxLogs.ParentBlockHash
	// if curHeight != 0 && parentHash != curBlockHash {
	// 	w.logger.Infof("delete %s block at height %d, hash=%s(parent hash = %s)",
	// 		worker.GetChainName(), curHeight, curBlockHash, parentHash)
	// 	return w.storage.DeleteBlockAndTxs(worker.GetChainName(), curHeight)
	// }

	nextBlockLog := storage.BlockLog{
		Chain:      worker.GetChainName(),
		BlockHash:  blockAndTxLogs.BlockHash,
		ParentHash: parentHash,
		Height:     blockAndTxLogs.Height,
		BlockTime:  blockAndTxLogs.BlockTime,
		Type:       storage.BlockTypeCurrent,
		CreateTime: time.Now().Unix(),
	}

	// put block header and block txs into database
	if err := w.storage.SaveBlockAndTxs(worker.GetChainName(), &nextBlockLog, blockAndTxLogs.TxLogs); err != nil {
		return err
	}

	// update prev txs confirmed number(&& update_time)
	if err := w.storage.UpdateConfirmedNum(worker.GetChainName(), nextBlockLog.Height); err != nil {
		return err
	}

	return nil
}
