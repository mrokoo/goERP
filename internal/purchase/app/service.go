package app

import (
	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/purchase/domain"
)

type PurchaseService interface {
	GetPurchaseOrderList() ([]*domain.PurchaseOrder, error)
	AddPurchaseOrder(purchaseOrder *domain.PurchaseOrder) error
	// InvalidatePurchaseOrder(purchaseOrderID string) error
}

type PurchaseServiceImpl struct {
	orderRepository  domain.PurchaseOrderRepository
	inventoryService inventory.InventoryService
}

func NewPurchaseServiceImpl(orderRepository domain.PurchaseOrderRepository, inventoryService inventory.InventoryService) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		orderRepository:  orderRepository,
		inventoryService: inventoryService,
	}
}

func (s *PurchaseServiceImpl) GetPurchaseOrderList() ([]*domain.PurchaseOrder, error) {
	return s.orderRepository.GetAll()
}

func (s *PurchaseServiceImpl) AddPurchaseOrder(purchaseOrder *domain.PurchaseOrder) error {
	if err := s.orderRepository.Save(purchaseOrder); err != nil {
		return err
	}
	p := purchaseOrder
	var taskItems []task.TaskItem
	for _, item := range p.Items {
		taskItem := task.NewTaskItem(item.ProductID, item.Quantity)
		taskItems = append(taskItems, taskItem)
	}
	if err := s.inventoryService.CreateTask(p.WarehouseID, task.IN_PURCHASE, p.ID, taskItems); err != nil {
		return err
	}
	return nil
}

// func (s *PurchaseServiceImpl) InvalidatePurchaseOrder(purchaseOrderID string) error {
// 	return s.orderRepository.Invalidated(purchaseOrderID)
// }
