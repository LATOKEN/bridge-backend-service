package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
1. CREATE event IN ./tx.go|ConfirmWorkerTx(...)
2. GET event
3. UPDATE event
*/

func (d *DataBase) getEventBySwapID(swapID string) (event Event, err error) {
	if err = d.db.Model(Event{}).Where("swap_id = ?", swapID).First(&event).Error; err != nil {
		return event, err
	}
	return event, nil
}

// GetEventsByTypeAndStatuses ...
func (d *DataBase) GetEventsByTypeAndStatuses(statuses []EventStatus) []*Event {
	swaps := make([]*Event, 0)
	if err := d.db.Where("status in (?)", statuses).
		Find(&swaps).Error; err != nil {
		return nil
	}

	return swaps
}

// UpdateEventStatus ...
func (d *DataBase) UpdateEventStatus(event *Event, status EventStatus) {
	event.Status = status
	event.UpdateTime = time.Now().Unix()

	d.db.Model(Event{}).Where("swap_id = ?", event.SwapID).Update(event)
}

// CompensateNewEvent ...
func (d *DataBase) CompensateNewEvent(chain string, tx *gorm.DB, newEvents []*Event) error {
	for _, event := range newEvents {
		txLogs, err := d.GetConfirmedTxsLog(chain, event, tx)
		if err != nil {
			continue
		}

		if len(txLogs) == 0 {
			continue
		}

		if err = d.ConfirmTx(tx, txLogs[0]); err != nil {
			return err
		}
	}

	return nil
}

// UpdateEventStatusWhenConfirmTx ...
func (d *DataBase) UpdateEventStatusWhenConfirmTx(tx *gorm.DB, txLog *TxLog,
	inStatuses []EventStatus, notInStatuses []EventStatus, updateStatus EventStatus) error {
	query := tx.Model(Event{})

	query = query.Where("swap_id = ?", txLog.SwapID)

	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}

	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":      updateStatus,
		"update_time": time.Now().Unix(),
	}

	return query.Updates(toUpdate).Error
}

func (d *DataBase) UpdateEventStatusWithTxStatus(txSent *TxSent, status TxStatus, txType TxType) error {
	event, err := d.getEventBySwapID(txSent.SwapID)
	if err != nil {
		return err
	}

	switch txType {
	case TxTypePassed:
		if status == TxSentStatusFailed {
			d.UpdateParticularEventStatus(event, EventStatusPassedSentFailed, EventStatusPassedSent)
		} else if status == TxSentStatusSuccess {
			d.UpdateParticularEventStatus(event, EventStatusPassedConfirmed, EventStatusPassedSent)
		}
	}
	return nil
}

func (d *DataBase) UpdateParticularEventStatus(event Event, status EventStatus, inStatus EventStatus) error {
	query := d.db.Model(Event{})
	query = query.Where("swap_id = ? and status = ?", event.SwapID, inStatus)
	event.Status = status
	return query.Updates(event).Error
}
