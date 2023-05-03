package record

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
)

type Record struct {
	ID          string       `json:"id"`
	WarehouseID string       `json:"warehouse_id"`
	UserID      string       `json:"user_id"`
	State       state.State  `json:"state"`
	Items       []RecordItem `json:"items"`
	CreatedAt   time.Time    `json:"created_at"`
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
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func NewRecordItem(productID string, quantity int) RecordItem {
	return RecordItem{
		ProductID: productID,
		Quantity:  quantity,
	}
}
