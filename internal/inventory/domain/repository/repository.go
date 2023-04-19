package domain

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
)

type InventoryFlowRepository interface {
	GetAll() ([]*flowrecord.InventoryFlow, error)
	GetByID(ID string) (*flowrecord.InventoryFlow, error)
	Save(flowRecord *flowrecord.InventoryFlow) error
	GetByProductIDAndWarehouseID(productID string, warehouseID string) (*flowrecord.InventoryFlow, error)
}

type TaskRepository interface {
	GetAll() ([]*task.Task, error)
	GetByID(ID string) (*task.Task, error)
	Save(task *task.Task) error
}

type AllotRepository interface {
	GetAll() ([]*allot.Allot, error)
	GetByID(ID string) (*allot.Allot, error)
	Save(allot *allot.Allot) error
}
