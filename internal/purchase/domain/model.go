package domain

import (
	"time"

	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/biling"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/item"
)

type PurchaseOrder struct {
	ID          string
	SupplierID  string
	WarehouseID string
	HandlerID   string
	CreatedTime time.Time
	Items       []item.Item
	Biling      biling.Biling
}
