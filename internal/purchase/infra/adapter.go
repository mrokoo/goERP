package repository

import (
	"time"

	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/biling"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/item"
	supplier "github.com/mrokoo/goERP/internal/share/supplier/domain"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type PurchaseOrder struct {
	ID          string              `gorm:"primaryKey"`
	SupplierID  string              `gorm:"size:191"`
	Supplier    supplier.Supplier   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WarehouseID string              `gorm:"size:191"`
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      string              `gorm:"size:191"`
	User        user.User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Date        time.Time
	Items       []PurchaseOrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Note        string
	biling.Biling
	IsValidated bool `gorm:"default:false"`
}

func toMySQLPurchaseOrder(purchaseorder *domain.PurchaseOrder) *PurchaseOrder {
	pots := toMySQLPurchaseOrderItem(purchaseorder.Items)
	return &PurchaseOrder{
		ID:          purchaseorder.ID,
		SupplierID:  purchaseorder.SupplierID,
		WarehouseID: purchaseorder.WarehouseID,
		UserID:      purchaseorder.UserID,
		Date:        purchaseorder.Date,
		Items:       pots,
		Note:        purchaseorder.Note,
		Biling:      purchaseorder.Biling,
		IsValidated: purchaseorder.IsValidated,
	}
}

func (p *PurchaseOrder) toPurchaseOrder() *domain.PurchaseOrder {
	var pots []item.PurchaseOrderItem
	for _, poi := range p.Items {
		pots = append(pots, poi.toPurchaseOrderItem())
	}

	return &domain.PurchaseOrder{
		ID:          p.ID,
		SupplierID:  p.SupplierID,
		WarehouseID: p.WarehouseID,
		UserID:      p.UserID,
		Date:        p.Date,
		Items:       pots,
		Note:        p.Note,
		Biling:      p.Biling,
		IsValidated: p.IsValidated,
	}
}

type PurchaseOrderItem struct {
	PurchaseOrderID string          `gorm:"primaryKey;size:191"`
	ProductID       string          `gorm:"primaryKey;size:191"`
	Product         product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity        int
	Price           float64
	TotalAmount     float64
}

func toMySQLPurchaseOrderItem(items []item.PurchaseOrderItem) []PurchaseOrderItem {
	var pot []PurchaseOrderItem
	for _, item := range items {
		p := PurchaseOrderItem{
			PurchaseOrderID: item.PurchaseOrderID,
			ProductID:       item.ProductID,
			Quantity:        item.Quantity,
			Price:           item.Price,
			TotalAmount:     item.TotalAmount,
		}
		pot = append(pot, p)
	}
	return pot
}

func (pot *PurchaseOrderItem) toPurchaseOrderItem() item.PurchaseOrderItem {
	return item.PurchaseOrderItem{
		PurchaseOrderID: pot.PurchaseOrderID,
		ProductID:       pot.ProductID,
		Quantity:        pot.Quantity,
		Price:           pot.Price,
		TotalAmount:     pot.TotalAmount,
	}
}
