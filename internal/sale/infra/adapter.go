package infra

import (
	"github.com/google/uuid"
	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/sale/domain"
	customer "github.com/mrokoo/goERP/internal/share/customer/domain"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type MySQLSaleOrder struct {
	ID          string              `gorm:"primaryKey; size:191;"`
	WarehouseID string              `gorm:"size:191;"`
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CustomerID  string              `gorm:"size:191;"`
	Customer    customer.Customer   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      string              `gorm:"size:191;"`
	User        user.User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt   string
	Basic       string // 只对return有效
	Items       []MySQLItem
	Kind        domain.Kind
}

func (s *MySQLSaleOrder) TableName() string {
	return "sale_orders"
}

type MySQLItem struct {
	ID               string               `gorm:"primaryKey; size:191;"`
	MySQLSaleOrderID string               `gorm:"size:191;"`
	ProductID        string               `gorm:"size:191;"`
	Product          product.MySQLProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount           int
	Price            float64
}

func (i *MySQLItem) TableName() string {
	return "sale_items"
}

func toMySQLSaleOrder(order *domain.SaleOrder) *MySQLSaleOrder {
	return &MySQLSaleOrder{
		ID:          order.ID,
		WarehouseID: order.WarehouseID,
		CustomerID:  order.CustomerID,
		UserID:      order.UserID,
		CreatedAt:   order.CreatedAt,
		Basic:       order.Basic,
		Items:       toMySQLItems(order.Items),
		Kind:        order.Kind,
	}
}

func toMySQLItems(items []domain.Item) []MySQLItem {
	mysqlItems := make([]MySQLItem, len(items))
	for i, item := range items {
		mysqlItems[i] = MySQLItem{
			ID:        uuid.New().String(),
			ProductID: item.ProductID,
			Amount:    item.Amount,
			Price:     item.Price,
		}
	}
	return mysqlItems
}

func (s *MySQLSaleOrder) toSaleOrder() *domain.SaleOrder {
	return &domain.SaleOrder{
		ID:          s.ID,
		WarehouseID: s.WarehouseID,
		CustomerID:  s.CustomerID,
		UserID:      s.UserID,
		CreatedAt:   s.CreatedAt,
		Basic:       s.Basic,
		Items:       toItems(s.Items),
		Kind:        s.Kind,
	}
}

func (s *MySQLItem) toItem() domain.Item {
	return domain.Item{
		ProductID: s.ProductID,
		Amount:    s.Amount,
		Price:     s.Price,
	}
}

func toItems(items []MySQLItem) []domain.Item {
	mysqlItems := make([]domain.Item, len(items))
	for i, item := range items {
		mysqlItems[i] = item.toItem()
	}
	return mysqlItems
}
