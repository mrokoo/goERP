package app

import (
	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"

	"github.com/mrokoo/goERP/internal/sale/domain"
)

type SaleService interface {
	GetSaleOrder(id string) (*domain.SaleOrder, error)
	GetSaleOrderList() ([]*domain.SaleOrder, error)
	CreateSaleOrder(order *domain.SaleOrder) error
}

type SaleServiceImpl struct {
	repo             domain.SaleOrderRepository
	inventoryService inventory.InventoryService
}

func (s SaleServiceImpl) GetSaleOrder(id string) (*domain.SaleOrder, error) {
	order, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s SaleServiceImpl) GetSaleOrderList() ([]*domain.SaleOrder, error) {
	orders, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s SaleServiceImpl) CreateSaleOrder(order *domain.SaleOrder) error {
	if err := s.repo.Save(order); err != nil {
		return err
	}
	var taskItems []task.TaskItem
	for _, item := range order.Items {
		taskItem := task.NewTaskItem(item.ProductID, item.Quantity)
		taskItems = append(taskItems, taskItem)
	}
	switch order.Kind {
	case domain.Order:
		if err := s.inventoryService.CreateTask(order.WarehouseID, task.OUT_SALE, order.ID, taskItems); err != nil {
			return err
		}
	case domain.ReturnOrder:
		if err := s.inventoryService.CreateTask(order.WarehouseID, task.IN_SALE, order.ID, taskItems); err != nil {
			return err
		}
	}
	return nil
}
