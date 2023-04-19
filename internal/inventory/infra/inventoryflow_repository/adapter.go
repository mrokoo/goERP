package inventoryflow_repository

import (
	"time"

	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

type MySQLInventoryFlow struct {
	ID          string              `gorm:"primaryKey;size:191;"`
	TaskID      *string             `gorm:"size:191;"`
	Task        task.Task           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TakeID      *string             `gorm:"size:191;"`
	ProductID   string              `gorm:"size:191;"`
	Product     product.Product     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WarehouseID string              `gorm:"size:191;"`
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Flow        flowrecord.FlowType
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Present     int // Present Quantity
	Date        time.Time
}

func toMySQLInventoryFlow(flow flowrecord.InventoryFlow) *MySQLInventoryFlow {
	return &MySQLInventoryFlow{
		ID:          flow.ID,
		TaskID:      flow.TaskID,
		TakeID:      flow.TakeID,
		ProductID:   flow.ProductID,
		WarehouseID: flow.WarehouseID,
		Flow:        flow.Flow,
		Previous:    flow.Previous,
		Change:      flow.Change,
		Present:     flow.Present,
		Date:        flow.Date,
	}
}

func (f MySQLInventoryFlow) toInventoryFlow() flowrecord.InventoryFlow {
	return flowrecord.InventoryFlow{
		ID:          f.ID,
		TaskID:      f.TaskID,
		TakeID:      f.TakeID,
		ProductID:   f.ProductID,
		WarehouseID: f.WarehouseID,
		Flow:        f.Flow,
		Previous:    f.Previous,
		Change:      f.Change,
		Present:     f.Present,
		Date:        f.Date,
	}
}
