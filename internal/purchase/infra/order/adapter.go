package order

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
	biling.Biling
	IsValidated bool                `gorm:"default:false"`
	Items       []PurchaseOrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Note        string
	Date        time.Time
}

func toMySQLPurchaseOrder(order *domain.PurchaseOrder) *PurchaseOrder {
	pots := toMySQLPurchaseOrderItem(order.ID, order.Items)
	return &PurchaseOrder{
		ID:          order.ID,
		SupplierID:  order.SupplierID,
		WarehouseID: order.WarehouseID,
		UserID:      order.UserID,
		Items:       pots,
		Biling:      order.Biling,
		IsValidated: order.IsValidated,
		Note:        order.Note,
		Date:        order.Date,
	}
}

func (p *PurchaseOrder) toPurchaseOrder() *domain.PurchaseOrder {
	var pois []item.OrderItem
	for _, poi := range p.Items {
		pois = append(pois, poi.toPurchaseOrderItem())
	}

	return &domain.PurchaseOrder{
		ID:          p.ID,
		SupplierID:  p.SupplierID,
		WarehouseID: p.WarehouseID,
		UserID:      p.UserID,
		Date:        p.Date,
		Items:       pois,
		Note:        p.Note,
		Biling:      p.Biling,
		IsValidated: p.IsValidated,
	}
}

type PurchaseOrderItem struct {
	PurchaseOrderID string               `gorm:"primaryKey;size:191"`
	ProductID       string               `gorm:"primaryKey;size:191"`
	Product         product.MySQLProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity        int
	Price           float64
	TotalPayment    float64
}

func toMySQLPurchaseOrderItem(OrderID string, items []item.OrderItem) []PurchaseOrderItem {
	var poi []PurchaseOrderItem
	for _, item := range items {
		p := PurchaseOrderItem{
			PurchaseOrderID: OrderID,
			ProductID:       item.ProductID,
			Quantity:        item.Quantity,
			Price:           item.Price,
			TotalPayment:    item.TotalPayment,
		}
		poi = append(poi, p)
	}
	return poi
}

func (pot *PurchaseOrderItem) toPurchaseOrderItem() item.OrderItem {
	return item.OrderItem{
		ProductID:    pot.ProductID,
		Quantity:     pot.Quantity,
		Price:        pot.Price,
		TotalPayment: pot.TotalPayment,
	}
}
