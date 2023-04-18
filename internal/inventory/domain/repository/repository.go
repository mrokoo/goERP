package domain

import (
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
)

type InventoryFlowRepository interface {
	GetAll() ([]*flowrecord.InventoryFlow, error)
	GetByID(ID string) (*flowrecord.InventoryFlow, error)
	Save(flowRecord *flowrecord.InventoryFlow) error
}

type TaskRepository interface {
	GetAll() ([]*task.Task, error)
	GetByID(ID string) (*task.Task, error)
	Save(task *task.Task) error
}
