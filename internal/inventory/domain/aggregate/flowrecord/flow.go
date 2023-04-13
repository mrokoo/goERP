package flowrecord

import "time"

type InventoryFlow struct {
	ID          string
	ProductID   string
	WarehouseID string
	FlowType    FlowType
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Date        time.Time
	Basic
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
	InTaskID  *string
	OutTaskID *string
	// to do 盘点和调拨
}
