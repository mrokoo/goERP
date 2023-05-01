package app

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	domain "github.com/mrokoo/goERP/internal/inventory/domain/repository"
)

type InventoryService interface {
	// task api
	GetTaskList() ([]*task.Task, error)
	GetTaskByPurchaseID(purchaseID string, kind task.Kind) (*task.Task, error)
	CreateTask(warehouseID string, kind task.Kind, basic string, items []task.TaskItem) error
	CreateTaskItem(productID string, total int) task.TaskItem
	InvalidateTask(taskID string) error
	AddRecord(taskID string, record record.Record) error
	CreateRecord(warehouseID string, userID string, items []record.RecordItem) (record.Record, error)
	CreateRecordItem(productID string, total int) record.RecordItem
	InvalidateRecord(taskID string, recordID string) error

	// flow api
	GetFlowList() ([]*flowrecord.InventoryFlow, error)
	CreateFlow(basicID string, productID string, warehouseID string, flow flowrecord.FlowType, change int) error
	GetPreviousProductQuantity(productID string, warehouseID string) (int, error)

	// allot api
	GetAllotList() ([]*allot.Allot, error)
	CreateAllot(in string, out string, userID string, items []allot.Item) error

	// take api
	GetTakeList() ([]*take.Take, error)
	CreateTake(warehouseID, userID string, items []take.Item) error
}

type InventoryServiceImpl struct {
	InventoryFlowRepository domain.InventoryFlowRepository
	TaskRepository          domain.TaskRepository
	AllotRepository         domain.AllotRepository
	TakeRepository          domain.TakeRepository
}

func NewInventoryServiceImpl(inventoryFlowRepository domain.InventoryFlowRepository, taskRepository domain.TaskRepository, allotRepository domain.AllotRepository, takeRepository domain.TakeRepository) InventoryServiceImpl {
	return InventoryServiceImpl{
		InventoryFlowRepository: inventoryFlowRepository,
		TaskRepository:          taskRepository,
		AllotRepository:         allotRepository,
		TakeRepository:          takeRepository,
	}
}

// Taskapi
func (i InventoryServiceImpl) GetTaskList() ([]*task.Task, error) {
	list, err := i.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i InventoryServiceImpl) GetTaskByPurchaseID(purchaseID string, kind task.Kind) (*task.Task, error) {
	t, err := i.TaskRepository.GetTaskByPurchaseID(purchaseID, kind)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// CreateTask 不直接暴露，需要在其他包中调用
func (i InventoryServiceImpl) CreateTask(warehouseID string, kind task.Kind, basic string, items []task.TaskItem) error {
	t := task.NewTask(warehouseID, kind, basic, items)
	err := i.TaskRepository.Save(&t)
	return err
}

func (i InventoryServiceImpl) CreateTaskItem(productID string, total int) task.TaskItem {
	item := task.NewTaskItem(productID, total)
	return item
}

// 不直接暴露，需要给其他包使用
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
	signal := 1
	var kind flowrecord.FlowType
	switch t.Kind {
	case task.IN_PURCHASE, task.IN_SALE:
		kind = flowrecord.FLOWTYPE_RUKU
		signal = 1
	case task.OUT_PURCHASE, task.OUT_SALE:
		kind = flowrecord.FLOWTYPE_CHUKU
		signal = -1
	case task.IN_ALLOCATION:
		kind = flowrecord.FLOWTYPE_DIAOBO
		signal = 1
	case task.OUT_ALLOCATION:
		kind = flowrecord.FLOWTYPE_DIAOBO
		signal = -1
	}
	for _, ri := range record.Items {
		if err := i.CreateFlow(taskID, ri.ProductID, t.WarehouseID, kind, signal*ri.Quantity); err != nil {
			return err
		}
	}

	if err := i.TaskRepository.Save(t); err != nil {
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
	signal := 1
	var kind flowrecord.FlowType
	switch t.Kind {
	case task.IN_PURCHASE, task.IN_SALE:
		kind = flowrecord.FLOWTYPE_ZUOFEIRUKU
		signal = -1
	case task.OUT_PURCHASE, task.OUT_SALE:
		kind = flowrecord.FLOWTYPE_ZUOFEICHUKU
		signal = 1
	case task.IN_ALLOCATION:
		kind = flowrecord.FLOWTYPE_ZUOFEIDIAOBO
		signal = -1
	case task.OUT_ALLOCATION:
		kind = flowrecord.FLOWTYPE_ZUOFEIDIAOBO
		signal = 1
	}
	for _, r := range t.Recrods {
		if r.ID == recordID {
			for _, ri := range r.Items {
				if err := i.CreateFlow(taskID, ri.ProductID, t.WarehouseID, kind, signal*ri.Quantity); err != nil {
					return err
				}
			}
		}
	}

	if err := i.TaskRepository.Save(t); err != nil {
		return err
	}
	return nil
}

// flow流水记录api
func (i InventoryServiceImpl) GetFlowList() ([]*flowrecord.InventoryFlow, error) {
	list, err := i.InventoryFlowRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i InventoryServiceImpl) CreateFlow(basicID string, productID string, warehouseID string, flow flowrecord.FlowType, change int) error {
	previous, err := i.GetPreviousProductQuantity(productID, warehouseID)
	if err != nil {
		return err
	}
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

// allot api
func (i InventoryServiceImpl) GetAllotList() ([]*allot.Allot, error) {
	list, err := i.AllotRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i InventoryServiceImpl) CreateAllot(in string, out string, userID string, items []allot.Item) error {
	a := allot.NewAllot(in, out, userID, items)
	var taskItems []task.TaskItem
	for _, item := range items {
		taskItems = append(taskItems, task.NewTaskItem(item.ProductID, item.Quantity))
	}
	if err := i.CreateTask(a.InWarehouseID, task.IN_ALLOCATION, a.ID, taskItems); err != nil {
		return err
	}
	if err := i.CreateTask(a.OutWarehouseID, task.OUT_ALLOCATION, a.ID, taskItems); err != nil {
		return err
	}
	if err := i.AllotRepository.Save(a); err != nil {
		return err
	}
	return nil
}

// take api
func (i InventoryServiceImpl) GetTakeList() ([]*take.Take, error) {
	list, err := i.TakeRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (i InventoryServiceImpl) CreateTake(warehouseID, userID string, items []take.Item) error {
	t := take.NewTake(warehouseID, userID, items)
	for _, item := range t.Items {
		if err := i.CreateFlow(t.ID, item.ProductID, t.WarehouseID, flowrecord.FLOWTYPE_PANDIAN, item.Quantity); err != nil {
			return err
		}
	}
	if err := i.TakeRepository.Save(&t); err != nil {
		return err
	}
	return nil
}
