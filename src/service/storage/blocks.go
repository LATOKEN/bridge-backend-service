package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
- CREATE - SaveBlockAndTxs
- GET - GetCurrentBlockLog
- UPDATE - UpdateConfirmedNum
- DELETE - DeleteBlockAndTxs
*/

// SaveBlockAndTxs saves block header and block's txs(=txLogs) into database
// Block header and height into 'block_logs'
// Txs into 'tx_logs'
// TxLogs contains transactions with 'our' events hashes
func (d *DataBase) SaveBlockAndTxs(chain string, blockLog *BlockLog, txLogs []*TxLog) error {
	tx := d.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("chain = ? and type = ?", chain, BlockTypeParent).
		Delete(BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(BlockLog{}).Where("chain = ? and type = ?", chain, BlockTypeCurrent).Updates(
		map[string]interface{}{
			"type": BlockTypeParent,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(blockLog).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, txLog := range txLogs {
		txLog.CreateTime = time.Now().Unix()
		if err := tx.Create(txLog).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeleteBlockAndTxs deletes from 'block_logs' and 'tx_logs' block and txs with
// current chain and height of block
func (d *DataBase) DeleteBlockAndTxs(chain string, height int64) error {
	tx := d.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("height = ? and chain = ?", height, chain).Delete(BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("height = ? and chain = ? and status = ?", height, chain, TxStatusInit).Delete(TxLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateConfirmedNum updates number of tx confirmation
func (d *DataBase) UpdateConfirmedNum(chain string, height int64) error {
	return d.db.Model(TxLog{}).Where("chain = ? and status = ?", chain, TxStatusInit).Updates(
		map[string]interface{}{
			"confirmed_num": gorm.Expr("? - height", height),
			"update_time":   time.Now().Unix(),
		}).Error
}

// GetCurrentBlockLog returns current block's logs
func (d *DataBase) GetCurrentBlockLog(chainID string) (logs BlockLog) {
	d.db.Where("chain = ?", chainID).Order("height desc").First(&logs)
	return logs
}
