package domain

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flowrecord"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
)

type InventoryFlowRepository interface {
	GetAll() ([]*flowrecord.InventoryFlow, error)
	GetByID(ID string) (*flowrecord.InventoryFlow, error)
	Save(flowRecord *flowrecord.InventoryFlow) error
}

type InTaskRepository interface {
	GetAll() ([]*task.InTask, error)
	GetByID() (*task.InTask, error)
	Save(inTask *task.InTask) error
}

type OutTaskRepository interface {
	GetAll() ([]*task.OutTask, error)
	GetByID() (*task.OutTask, error)
	Save(outTask *task.OutTask) error
}


