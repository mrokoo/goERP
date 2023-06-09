package flowrecord

import (
	"time"

	"github.com/google/uuid"
)

type InventoryFlow struct {
	ID          string
	TaskID      *string
	TakeID      *string
	ProductID   string
	WarehouseID string
	Flow        FlowType
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Present     int // Present Quantity
	CreateAt    time.Time
}

func NewInventoryFlow(basicID string, productID string, warehouseID string, flow FlowType, previous int, change int) InventoryFlow {
	var taskID *string
	var takeID *string
	if flow == FLOWTYPE_PANDIAN {
		takeID = &basicID
	} else {
		taskID = &basicID
	}
	return InventoryFlow{
		ID:          uuid.New().String(),
		TaskID:      taskID,
		TakeID:      takeID,
		ProductID:   productID,
		WarehouseID: warehouseID,
		Flow:        flow,
		Previous:    previous,
		Change:      change,
		Present:     previous + change,
		CreateAt:    time.Now(),
	}
}

type FlowType string

const (
	FLOWTYPE_RUKU          FlowType = "入库"
	FLOWTYPE_ZUOFEIRUKU    FlowType = "作废入库"
	FLOWTYPE_CHUKU         FlowType = "出库"
	FLOWTYPE_ZUOFEICHUKU   FlowType = "作废出库"
	FLOWTYPE_DIAOBO        FlowType = "调拨"
	FLOWTYPE_ZUOFEIDIAOBO  FlowType = "作废调拨"
	FLOWTYPE_PANDIAN       FlowType = "盘点"
	FLOWTYPE_ZUOFEIPANDIAN FlowType = "作废盘点"
)
