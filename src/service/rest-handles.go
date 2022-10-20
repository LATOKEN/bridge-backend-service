package rlr

import (
	"github.com/latoken/bridge-backend-service/src/models"
)

// Status ...
func (r *BridgeSRV) StatusOfWorkers() (map[string]*models.WorkerStatus, error) {
	// get blockchain heights from workers and from database
	workers := make(map[string]*models.WorkerStatus)
	for _, w := range r.Workers {
		status, err := w.GetStatus()
		if err != nil {
			r.logger.Errorf("While get status for worker = %s, err = %v", w.GetChainName(), err)
			return nil, err
		}
		workers[w.GetChainName()] = status
	}

	for name, w := range workers {
		blocks := r.storage.GetCurrentBlockLog(name)
		w.SyncHeight = blocks.Height
	}

	return workers, nil
}

func (r *BridgeSRV) GetTxSent(txHash string) (string, error) {
	txSent, err := r.storage.GetTxSentByTxHash(txHash)
	if err != nil {
		r.logger.Errorf("GetTxSent, txHash: %v, failed with error: %v", txHash, err)
		return "", err
	}

	return txSent, nil
}

// CreateNewBindRequest ...
func (r *BridgeSRV) CreateNewBindRequest() {}

// GetRelayerAccountBalance ...
func (r *BridgeSRV) GetRelayerAccountBalance() {

}

// GetWorkers ...
func (r *BridgeSRV) GetWorkers() {
	//	r.Workers
}

func (r *BridgeSRV) GetGasPrice(name string) string {
	priceLog := r.storage.GetGasPrice(name)
	return priceLog.Price
}
