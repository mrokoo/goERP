package allot

import (
	"time"

	"github.com/google/uuid"
)

type Allot struct {
	ID             string
	InWarehouseID  string
	OutWarehouseID string
	UserID         string
	CreatedAt      time.Time
	Items          []Item
}

func NewAllot(in string, out string, user string, items []Item) *Allot {
	return &Allot{
		ID:             uuid.New().String(),
		InWarehouseID:  in,
		OutWarehouseID: out,
		UserID:         user,
		CreatedAt:      time.Now(),
		Items:          items,
	}
}

type Item struct {
	ProductID string
	Quantity  int
}
