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
	Items       []item.OrderItem
	biling.Biling
	IsValidated bool
	Note        string
	Date        time.Time
}

type PurchaseReturnOrder struct {
	ID              string
	PurchaseOrderID string // 并不强制要求
	SupplierID      string
	WarehouseID     string
	UserID          string
	Items           []item.ReturnOrderItem
	biling.Biling
	IsValidated bool
	Note        string
	Date        time.Time
}
