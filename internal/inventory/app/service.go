package app

import (
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	domain "github.com/mrokoo/goERP/internal/inventory/domain/repository"
)

type InventoryService interface {
	GetTaskList() ([]*task.Task, error)
	CreateTask(warehouseID string, kind task.Kind, basic string, items []task.TaskItem) error
	CreateTaskItem(productID string, total int) task.TaskItem
	InvalidateTask(taskID string) error
	AddRecord(taskID string, record record.Record) error
	CreateRecord(warehouseID string, userID string, items []record.RecordItem) (record.Record, error)
	CreateRecordItem(productID string, total int) record.RecordItem
	InvalidateRecord(taskID string, recordID string) error
	GetFlowList() ([]*flowrecord.InventoryFlow, error)
	CreateFlow(basicID string, productID string, warehouseID string, flow flowrecord.FlowType, previous int, change int) error
	GetPreviousProductQuantity(productID string, warehouseID string) (int, error)
}

type InventoryServiceImpl struct {
	InventoryFlowRepository domain.InventoryFlowRepository
	TaskRepository          domain.TaskRepository
}

func (i InventoryServiceImpl) GetTaskList() ([]*task.Task, error) {
	list, err := i.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 不直接暴露，需要在其他包中调用
func (i InventoryServiceImpl) CreateTask(warehouseID string, kind task.Kind, basic string, items []task.TaskItem) error {
	t := task.NewTask(warehouseID, kind, basic, items)
	err := i.TaskRepository.Save(&t)
	return err
}

func (i InventoryServiceImpl) CreateTaskItem(productID string, total int) task.TaskItem {
	item := task.NewTaskItem(productID, total)
	return item
}

func (i InventoryServiceImpl) InvalidateTask(taskID string) error {
	ID := taskID
	t, err := i.TaskRepository.GetByID(ID)
	if err != nil {
		return err
	}
	if err := t.InvalidateTask(); err != nil {
		return err
	}
	if err := i.TaskRepository.Save(t); err != nil {
		return err
	}
	return nil
}

func (i InventoryServiceImpl) AddRecord(taskID string, record record.Record) error {
	t, err := i.TaskRepository.GetByID(taskID)
	if err != nil {
		return err
	}
	if err := t.AddRecord(record); err != nil {
		return err
	}
	return nil
}

func (i InventoryServiceImpl) CreateRecord(warehouseID string, userID string, items []record.RecordItem) (record.Record, error) {
	r := record.NewRecord(warehouseID, userID, items)
	return r, nil
}

func (i InventoryServiceImpl) CreateRecordItem(productID string, total int) record.RecordItem {
	item := record.NewRecordItem(productID, total)
	return item
}

func (i InventoryServiceImpl) InvalidateRecord(taskID string, recordID string) error {
	t, err := i.TaskRepository.GetByID(taskID)
	if err != nil {
		return err
	}
	if err := t.InvalidateRecord(recordID); err != nil {
		return err
	}
	if err := i.TaskRepository.Save(t); err != nil {
		return err
	}
	return nil
}

func (i InventoryServiceImpl) GetFlowList() ([]*flowrecord.InventoryFlow, error) {
	list, err := i.InventoryFlowRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i InventoryServiceImpl) CreateFlow(basicID string, productID string, warehouseID string, flow flowrecord.FlowType, previous int, change int) error {
	f := flowrecord.NewInventoryFlow(basicID, productID, warehouseID, flow, previous, change)
	if err := i.InventoryFlowRepository.Save(&f); err != nil {
		return err
	}
	return nil
}

func (i InventoryServiceImpl) GetPreviousProductQuantity(productID string, warehouseID string) (int, error) {
	f, err := i.InventoryFlowRepository.GetByProductIDAndWarehouseID(productID, warehouseID)
	if err != nil {
		return 0, err
	}
	return f.Present, nil
}
