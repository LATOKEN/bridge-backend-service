package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GetConfirmedTxsLog ...
func (d *DataBase) GetConfirmedTxsLog(chain string, event *Event, tx *gorm.DB) (txLogs []*TxLog, err error) {
	if err := tx.Where("chain = ? and status = ?", chain, TxStatusConfirmed).Find(&txLogs).Error; err != nil {
		return txLogs, err
	}

	return txLogs, nil
}

// FindTxLogs ...
func (d *DataBase) FindTxLogs(chainID string, confirmNum int64) (txLogs []*TxLog, err error) {
	if err := d.db.Where("chain = ? and status = ? and confirmed_num >= ?",
		chainID, TxStatusInit, confirmNum).Find(&txLogs).Error; err != nil {
		return nil, err
	}

	return txLogs, nil
}

// ConfirmWorkerTx ...
func (d *DataBase) ConfirmWorkerTx(chainID string, txLogs []*TxLog, txHashes []string, newEvents []*Event) error {
	tx := d.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := tx.Model(TxLog{}).Where("tx_hash in (?)", txHashes).Updates(
		map[string]interface{}{
			"status":      TxStatusConfirmed,
			"update_time": time.Now().Unix(),
		}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// create swap
	for _, swap := range newEvents {
		var previousSwap Event
		if tx.Model(Event{}).Where("swap_id = ?", swap.SwapID).Order("swap_id desc").First(&previousSwap); previousSwap.SwapID == "" {
			if err := tx.Create(swap).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			swap.Status = previousSwap.Status
			if err := tx.Model(Event{}).Where("swap_id = ?", swap.SwapID).Update(swap).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	for _, txLog := range txLogs {
		if err := d.ConfirmTx(tx, txLog); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := d.CompensateNewEvent(chainID, tx, newEvents); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// !!! TODO !!!

// ConfirmTx ...
func (d *DataBase) ConfirmTx(tx *gorm.DB, txLog *TxLog) error {
	switch txLog.TxType {
	case TxTypeDeposit:
		if err := d.UpdateEventStatusWhenConfirmTx(tx, txLog, []EventStatus{
			EventStatusDepositConfirmed},
			[]EventStatus{EventStatusClaimConfirmed, EventStatusPassedInit, EventStatusPassedInitConfrimed, EventStatusPassedSent, EventStatusPassedSentFailed, EventStatusPassedConfirmed, EventStatusPassedFailed, EventStatusUpdateConfirmed, EventStatusUpdateFailed, EventStatusSpendConfirmed, EventStatusExpiredConfirmed}, EventStatusClaimConfirmed); err != nil {
			return err
		}
	case TxTypeClaim:
		if err := d.UpdateEventStatusWhenConfirmTx(tx, txLog, []EventStatus{
			EventStatusDepositConfirmed, EventStatusClaimConfirmed},
			[]EventStatus{EventStatusPassedInit, EventStatusPassedInitConfrimed, EventStatusPassedSent, EventStatusPassedSentFailed, EventStatusPassedConfirmed, EventStatusPassedFailed, EventStatusUpdateConfirmed, EventStatusUpdateFailed, EventStatusSpendConfirmed, EventStatusExpiredConfirmed}, EventStatusClaimConfirmed); err != nil {
			return err
		}
	case TxTypePassed:
		if err := d.UpdateEventStatusWhenConfirmTx(tx, txLog, []EventStatus{
			EventStatusClaimConfirmed, EventStatusDepositConfirmed, EventStatusPassedInit},
			[]EventStatus{EventStatusPassedInitConfrimed, EventStatusPassedSent, EventStatusPassedSentFailed, EventStatusPassedConfirmed, EventStatusPassedFailed, EventStatusUpdateConfirmed, EventStatusUpdateFailed, EventStatusSpendConfirmed, EventStatusExpiredConfirmed}, EventStatusPassedInit); err != nil {
			return err
		}
	case TxTypeSpend:
		if err := d.UpdateEventStatusWhenConfirmTx(tx, txLog, []EventStatus{
			EventStatusDepositConfirmed, EventStatusClaimConfirmed, EventStatusPassedInit, EventStatusPassedInitConfrimed, EventStatusPassedSent, EventStatusPassedSentFailed, EventStatusPassedConfirmed, EventStatusUpdateConfirmed, EventStatusUpdateFailed},
			[]EventStatus{EventStatusExpiredConfirmed, EventStatusPassedFailed}, EventStatusSpendConfirmed); err != nil {
			return err
		}
	case TxTypeExpired:
		if err := d.UpdateEventStatusWhenConfirmTx(tx, txLog, []EventStatus{
			EventStatusDepositConfirmed, EventStatusClaimConfirmed, EventStatusPassedInit, EventStatusPassedInitConfrimed, EventStatusPassedSent, EventStatusPassedSentFailed, EventStatusPassedFailed, EventStatusUpdateConfirmed, EventStatusUpdateFailed},
			[]EventStatus{EventStatusSpendConfirmed, EventStatusPassedConfirmed}, EventStatusExpiredConfirmed); err != nil {
			return err
		}
	}

	return nil
}

// ------ TXSENT ------

// CreateTxSent ...
func (d *DataBase) CreateTxSent(txSent *TxSent) error {
	if txSent.Status == "" {
		txSent.Status = TxSentStatusInit
	}

	return d.db.Create(txSent).Error
}

// UpdateTxSentStatus ...
func (d *DataBase) UpdateTxSentStatus(txSent *TxSent, status TxStatus) error {
	return d.db.Model(TxSent{}).Where("id = ? and swap_id = ?", txSent.ID, txSent.SwapID).Update(
		map[string]interface{}{
			"status":      status,
			"update_time": time.Now().Unix(),
		}).Error
}

// GetTxsSentByStatus ...
func (d *DataBase) GetTxsSentByStatus(chain string) ([]*TxSent, error) {
	txsSent := make([]*TxSent, 0)
	status := []TxStatus{TxSentStatusInit, TxSentStatusNotFound, TxSentStatusPending}
	if err := d.db.Where("chain = ? and status in (?)", chain, status).Find(&txsSent).Error; err != nil {
		return nil, err
	}

	return txsSent, nil
}

// GetTxsSentByType ...
func (d *DataBase) GetTxsSentByType(chain string, txType TxType, event *Event) []*TxSent {
	txsSent := make([]*TxSent, 0)
	query := d.db.Where("swap_id = ? and type = ?", event.SwapID, txType)
	query.Order("id desc").Find(&txsSent)
	return txsSent
}

// GetTxSentByTxHash ...
func (d *DataBase) GetTxSentByTxHash(txHash string) (string, error) {
	txLog := &TxLog{}
	if err := d.db.Model(TxLog{}).Where("tx_hash = ? and tx_type = ?", txHash, TxTypeDeposit).
		Find(&txLog).Error; err != nil {
		return "", err
	}

	txSent := &TxSent{}
	if err := d.db.Model(TxSent{}).Where("swap_id = ? and type = ?", txLog.SwapID, TxTypePassed).
		Find(&txSent).Error; err != nil {
		return "", err
	}

	return txSent.TxHash, nil
}
