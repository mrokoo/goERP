package app

import (
	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/purchase/domain"
)

type PurchaseService interface {
	GetPurchaseOrderList() ([]*domain.PurchaseOrder, error)
	AddPurchaseOrder(purchaseOrder *domain.PurchaseOrder) error
	AddPurchaseReturnOrder(purchaseOrder *domain.PurchaseOrder) error
	InvalidatePurchaseOrder(purchaseOrderID string) error
	InvalidatePurchaseReturnOrder(purchaseOrderID string) error
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

func (s *PurchaseServiceImpl) AddPurchaseReturnOrder(returnOrder *domain.PurchaseOrder) error {
	if err := s.orderRepository.Save(returnOrder); err != nil {
		return err
	}
	p := returnOrder
	var taskItems []task.TaskItem
	for _, item := range p.Items {
		taskItem := task.NewTaskItem(item.ProductID, item.Quantity)
		taskItems = append(taskItems, taskItem)
	}
	if err := s.inventoryService.CreateTask(p.WarehouseID, task.OUT_PURCHASE, p.ID, taskItems); err != nil {
		return err
	}
	return nil
}

func (s *PurchaseServiceImpl) InvalidatePurchaseOrder(purchaseOrderID string) error {
	task_, err := s.inventoryService.GetTaskByPurchaseID(purchaseOrderID, task.IN_PURCHASE)
	if err != nil {
		return err
	}
	if err := s.inventoryService.InvalidateTask(task_.ID); err != nil {
		return err
	}
	if err := s.orderRepository.Invalidated(purchaseOrderID); err != nil {
		return err
	}
	return nil
}

func (s *PurchaseServiceImpl) InvalidatePurchaseReturnOrder(purchaseOrderID string) error {
	task_, err := s.inventoryService.GetTaskByPurchaseID(purchaseOrderID, task.OUT_PURCHASE)
	if err != nil {
		return err
	}
	if err := s.inventoryService.InvalidateTask(task_.ID); err != nil {
		return err
	}
	if err := s.orderRepository.Invalidated(purchaseOrderID); err != nil {
		return err
	}
	return nil
}
