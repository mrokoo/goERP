package flow

import "time"

type InventoryFlow struct {
	ProductID   string
	WarehouseID string
	FlowType    string
	Previous    int // Previous Quantity
	Change      int // Change Quantity
	Date        time.Time
	Basic       string // 依据
}
