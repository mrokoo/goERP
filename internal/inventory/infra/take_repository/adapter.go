package take_repository

import (
	"time"

	"github.com/google/uuid"
	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type MySQLTake struct {
	ID          string              `gorm:"primaryKey"`
	WarehouseID string              `gorm:"size:191;"`
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      string              `gorm:"size:191;"`
	User        user.User           `gorm:"contraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreateAt    time.Time
	Items       []MySQLItem
}

func (t *MySQLTake) TableName() string {
	return "takes"
}

func (t *MySQLTake) toTake() take.Take {
	return take.Take{
		ID:          t.ID,
		WarehouseID: t.WarehouseID,
		UserID:      t.UserID,
		CreateAt:    t.CreateAt,
		Items:       toItems(t.Items),
	}
}

func toTakes(takes []*MySQLTake) []*take.Take {
	var mysqlTakes []*take.Take
	for _, take := range takes {
		take_ := take.toTake()
		mysqlTakes = append(mysqlTakes, &take_)
	}
	return mysqlTakes
}

type MySQLItem struct {
	ID          string          `gorm:"primaryKey"`
	MySQLTakeID string          `gorm:"size:191;"`
	ProductID   string          `gorm:"size:191;"`
	Product     product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity    int
}

func (i *MySQLItem) TableName() string {
	return "take_items"
}

func (i *MySQLItem) toItem() take.Item {
	return take.Item{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}

func toItems(items []MySQLItem) []take.Item {
	var mysqlItems []take.Item
	for _, item := range items {
		mysqlItems = append(mysqlItems, item.toItem())
	}
	return mysqlItems
}

func toMySQLTake(take take.Take) MySQLTake {
	return MySQLTake{
		ID:          take.ID,
		WarehouseID: take.WarehouseID,
		UserID:      take.UserID,
		CreateAt:    take.CreateAt,
		Items:       toMySQLItems(take.ID, take.Items),
	}
}

func toMySQLItems(takeID string, items []take.Item) []MySQLItem {
	var mysqlItems []MySQLItem
	for _, item := range items {
		mysqlItems = append(mysqlItems, MySQLItem{
			ID:          uuid.New().String(),
			MySQLTakeID: takeID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
		})
	}
	return mysqlItems
}
