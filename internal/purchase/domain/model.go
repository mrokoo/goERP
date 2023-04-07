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
	UserID      string
	Date        time.Time
	Items       []item.PurchaseOrderItem
	Note        string
	biling.Biling
	IsValidated bool
}
