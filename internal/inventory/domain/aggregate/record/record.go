package record

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
)

type Record struct {
	ID          string
	WarehouseID string
	UserID      string
	State       state.State
	Items       []RecordItem
}

func NewRecord(warehouseID string, userID string, items []RecordItem) Record {
	return Record{
		ID:          uuid.New().String(),
		WarehouseID: warehouseID,
		UserID:      userID,
		State:       state.NORMAL,
		Items:       items,
	}
}

type RecordItem struct {
	ID        string
	ProductID string
	Quantity  int
}

func NewRecordItem(productID string, quantity int) RecordItem {
	return RecordItem{
		ID:        uuid.New().String(),
		ProductID: productID,
		Quantity:  quantity,
	}
}
