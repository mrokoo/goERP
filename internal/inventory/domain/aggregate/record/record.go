package record

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
)

type Record struct {
	ID          string
	WarehouseID string
	UserID      string
	State       state.State
	Items       []RecordItem
	CreatedAt   time.Time
}

func NewRecord(warehouseID string, userID string, items []RecordItem) Record {
	return Record{
		ID:          uuid.New().String(),
		WarehouseID: warehouseID,
		UserID:      userID,
		State:       state.NORMAL,
		Items:       items,
		CreatedAt:   time.Now(),
	}
}

type RecordItem struct {
	ProductID string
	Quantity  int
}

func NewRecordItem(productID string, quantity int) RecordItem {
	return RecordItem{
		ProductID: productID,
		Quantity:  quantity,
	}
}
