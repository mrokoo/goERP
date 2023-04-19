package allot_repository

import (
	"time"

	"github.com/google/uuid"
	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
)

type MySQLAllot struct {
	ID             string
	InWarehouseID  string
	OutWarehouseID string
	UserID         string
	CreatedAt      time.Time
	Items          []MySQLItem
}

func (a *MySQLAllot) TableName() string {
	return "allots"
}

type MySQLItem struct {
	ID        string
	AllotID   string
	ProductID string
	Product   product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity  int
}

func (i *MySQLItem) TableName() string {
	return "allot_items"
}

func toMySQLAllot(a *allot.Allot) *MySQLAllot {
	return &MySQLAllot{
		ID:             a.ID,
		InWarehouseID:  a.InWarehouseID,
		OutWarehouseID: a.OutWarehouseID,
		UserID:         a.UserID,
		CreatedAt:      a.CreatedAt,
		Items:          toMySQLItems(a.ID, a.Items),
	}
}

func toMySQLItems(allotID string, items []allot.Item) []MySQLItem {
	var mysqlItems []MySQLItem
	for _, item := range items {
		mysqlItems = append(mysqlItems, MySQLItem{
			ID:        uuid.New().String(),
			AllotID:   allotID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	return mysqlItems
}

func (a *MySQLAllot) ToAllot() *allot.Allot {
	return &allot.Allot{
		ID:             a.ID,
		InWarehouseID:  a.InWarehouseID,
		OutWarehouseID: a.OutWarehouseID,
		UserID:         a.UserID,
		CreatedAt:      a.CreatedAt,
		Items:          toAllotItems(a.Items),
	}
}

func (i *MySQLItem) ToItem() allot.Item {
	return allot.Item{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}

func toAllotItems(items []MySQLItem) []allot.Item {
	var mysqlItems []allot.Item
	for _, item := range items {
		mysqlItems = append(mysqlItems, item.ToItem())
	}
	return mysqlItems
}
