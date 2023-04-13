package inventoryflow_repository

import (
	"time"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flowrecord"
	"github.com/mrokoo/goERP/internal/inventory/infra/task_repository"
)

type InventoryFlow struct {
	ID          string
	ProductID   string `gorm:"size:191;primaryKey"`
	WarehouseID string `gorm:"size:191;primaryKey"`
	FlowType    FlowType
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Date        time.Time
	Basic
}

func toMySQLInventoryFlow(flow *flowrecord.InventoryFlow) *InventoryFlow {
	return &InventoryFlow{
		ID:          flow.ID,
		ProductID:   flow.ProductID,
		WarehouseID: flow.WarehouseID,
		FlowType:    FlowType(flow.FlowType),
		Previous:    flow.Previous,
		Change:      flow.Change,
		Date:        flow.Date,
		Basic:       toMySQLBasic(flow.Basic),
	}
}

func (i InventoryFlow) toInventoryFlow() *flowrecord.InventoryFlow {
	return &flowrecord.InventoryFlow{
		ID:          i.ID,
		ProductID:   i.ProductID,
		WarehouseID: i.WarehouseID,
		FlowType:    flowrecord.FlowType(i.FlowType),
		Previous:    i.Previous,
		Change:      i.Change,
		Date:        i.Date,
		Basic:       i.Basic.toBasic(),
	}
}

type FlowType string

const (
	FLOWTYPE_RUKU        FlowType = "入库"
	FLOWTYPE_ZUOFEIRUKU  FlowType = "作废入库"
	FLOWTYPE_CHUKU       FlowType = "出库"
	FLOWTYPE_ZUOFEICHUKU FlowType = "作废出库"
	FLOWTYPE_PANDIAN     FlowType = "盘点"
	FLOWTYPE_DIAOBO      FlowType = "调拨"
)

type Basic struct {
	InTaskID *string
	InTask   task_repository.InTask

	OutTaskID *string
	OutTask   task_repository.OutTask
	// to do 盘点和调拨
}

func toMySQLBasic(basic flowrecord.Basic) Basic {
	return Basic{
		InTaskID:  basic.InTaskID,
		OutTaskID: basic.OutTaskID,
	}
}

func (b Basic) toBasic() flowrecord.Basic {
	return flowrecord.Basic{
		InTaskID:  b.InTaskID,
		OutTaskID: b.OutTaskID,
	}
}
