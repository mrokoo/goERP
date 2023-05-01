package take

import (
	"time"

	"github.com/google/uuid"
)

type Take struct {
	ID          string
	WarehouseID string
	UserID      string
	CreateAt    time.Time
	Items       []Item
}

func NewTake(warehouseID string, userID string, items []Item) Take {
	return Take{
		ID:          uuid.New().String(),
		WarehouseID: warehouseID,
		UserID:      userID,
		CreateAt:    time.Now(),
		Items:       items,
	}
}

type Item struct {
	ProductID string
	Quantity  int
}
