package inventoryflow_repository

import (
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(flow *flowrecord.InventoryFlow) *model.InventoryFlow {
	return &model.InventoryFlow{
		ID:          flow.ID,
		TaskID:      flow.TaskID,
		TakeID:      flow.TakeID,
		ProductID:   flow.ProductID,
		WarehouseID: flow.WarehouseID,
		Flow:        string(flow.Flow),
		Previous:    flow.Previous,
		Change:      flow.Change,
		Present:     flow.Present,
		CreatedAt:   flow.CreateAt,
	}
}

func toDomain(flow *model.InventoryFlow) *flowrecord.InventoryFlow {
	return &flowrecord.InventoryFlow{
		ID:          flow.ID,
		TaskID:      flow.TaskID,
		TakeID:      flow.TakeID,
		ProductID:   flow.ProductID,
		WarehouseID: flow.WarehouseID,
		Flow:        flowrecord.FlowType(flow.Flow),
		Previous:    flow.Previous,
		Change:      flow.Change,
		Present:     flow.Present,
		CreateAt:    flow.CreatedAt,
	}
}
