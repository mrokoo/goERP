package allot_repository

import (
	"time"

	"github.com/google/uuid"
	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type MySQLAllot struct {
	ID             string `gorm:"primaryKey;"`
	InWarehouseID  string
	OutWarehouseID string
	UserID         string    `gorm:"size:191;"`
	User           user.User `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt      time.Time
	Items          []MySQLItem
}

func (a *MySQLAllot) TableName() string {
	return "allots"
}

type MySQLItem struct {
	ID           string               `gorm:"primaryKey;"`
	MySQLAllotID string               `gorm:"size:191;"`
	ProductID    string               `gorm:"size:191;"`
	Product      product.MySQLProduct `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity     int
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
			ID:           uuid.New().String(),
			MySQLAllotID: allotID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
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

func toAllots(allots []*MySQLAllot) []*allot.Allot {
	var mysqlAllots []*allot.Allot
	for _, allot := range allots {
		mysqlAllots = append(mysqlAllots, allot.ToAllot())
	}
	return mysqlAllots
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
