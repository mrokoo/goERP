package returnorder

import (
	"time"

	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/biling"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/item"
	"github.com/mrokoo/goERP/internal/purchase/infra/order"
	supplier "github.com/mrokoo/goERP/internal/share/supplier/domain"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type PurchaseOrder = order.PurchaseOrder
type PurchaseReturnOrder struct {
	ID              string                    `gorm:"primaryKey"`
	PurchaseOrderID string                    `gorm:"default:null;"` // 并不强制要求
	PurchaseOrder   PurchaseOrder             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SupplierID      string                    `gorm:"size:191"`
	Supplier        supplier.Supplier         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WarehouseID     string                    `gorm:"size:191"`
	Warehouse       warehouse.Warehouse       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID          string                    `gorm:"size:191"`
	User            user.User                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Items           []PurchaseReturnOrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	biling.Biling
	IsValidated bool `gorm:"default:false"`
	Note        string
	Date        time.Time
}

func toMySQLPurchaseReturnOrder(order *domain.PurchaseReturnOrder) *PurchaseReturnOrder {
	poris := toMySQLPurchaseReturnOrderItem(order.ID, order.Items)
	return &PurchaseReturnOrder{
		ID:              order.ID,
		PurchaseOrderID: order.PurchaseOrderID,
		SupplierID:      order.SupplierID,
		WarehouseID:     order.WarehouseID,
		UserID:          order.UserID,
		Date:            order.Date,
		Items:           poris,
		Note:            order.Note,
		Biling:          order.Biling,
		IsValidated:     order.IsValidated,
	}
}

func (p *PurchaseReturnOrder) toPurchaseReturnOrder() *domain.PurchaseReturnOrder {
	var pois []item.ReturnOrderItem
	for _, poi := range p.Items {
		pois = append(pois, poi.toPurchaseReturnOrderItem())
	}
	return &domain.PurchaseReturnOrder{
		ID:              p.ID,
		PurchaseOrderID: p.PurchaseOrderID,
		SupplierID:      p.SupplierID,
		WarehouseID:     p.WarehouseID,
		UserID:          p.UserID,
		Date:            p.Date,
		Items:           pois,
		Note:            p.Note,
		Biling:          p.Biling,
		IsValidated:     p.IsValidated,
	}
}

type PurchaseReturnOrderItem struct {
	PurchaseReturnOrderID string               `gorm:"primaryKey;size:191"`
	ProductID             string               `gorm:"primaryKey;size:191"`
	Product               product.MySQLProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity              int
	Price                 float64
	TotalCollection       float64
}

func toMySQLPurchaseReturnOrderItem(OrderID string, items []item.ReturnOrderItem) []PurchaseReturnOrderItem {
	var poi []PurchaseReturnOrderItem
	for _, item := range items {
		p := PurchaseReturnOrderItem{
			PurchaseReturnOrderID: OrderID,
			ProductID:             item.ProductID,
			Quantity:              item.Quantity,
			Price:                 item.Price,
			TotalCollection:       item.TotalCollection,
		}
		poi = append(poi, p)
	}
	return poi
}

func (pot *PurchaseReturnOrderItem) toPurchaseReturnOrderItem() item.ReturnOrderItem {
	return item.ReturnOrderItem{
		ProductID:       pot.ProductID,
		Quantity:        pot.Quantity,
		Price:           pot.Price,
		TotalCollection: pot.TotalCollection,
	}
}
