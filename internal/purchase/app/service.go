package app

import (
	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/purchase/domain"
)

type PurchaseService interface {
	GetPurchaseOrderList() ([]*domain.PurchaseOrder, error)
	AddPurchaeOrder(purchaseOrder *domain.PurchaseOrder) error
	InvalidatePurchaseOrder(purchaseOrder *domain.PurchaseOrder) error

	GetPurchaseReturnOrderList() ([]*domain.PurchaseReturnOrder, error)
	AddPurchaseReturnOrder(purchaseReturnOrder *domain.PurchaseReturnOrder) error
	InvalidatePurchaseReturnOrder(purchaseReturnOrder *domain.PurchaseReturnOrder) error
}

type PurchaseServiceImpl struct {
	orderRepository  domain.PurchaseOrderRepository
	returnRepository domain.PurchaseReturnOrderRepository
	inventoryService inventory.InventoryService
}

func NewPurchaseServiceImpl(orderRepository domain.PurchaseOrderRepository, returnRepository domain.PurchaseReturnOrderRepository) *PurchaseServiceImpl {
	return &PurchaseServiceImpl{
		orderRepository:  orderRepository,
		returnRepository: returnRepository,
	}
}

func (s *PurchaseServiceImpl) GetPurchaseOrderList() ([]*domain.PurchaseOrder, error) {
	return s.orderRepository.GetAll()
}

func (s *PurchaseServiceImpl) AddPurchaeOrder(purchaseOrder *domain.PurchaseOrder) error {
	if err := s.orderRepository.Save(*purchaseOrder); err != nil {
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

func (s *PurchaseServiceImpl) InvalidatePurchaseOrder(purchaseOrder *domain.PurchaseOrder) error {
	return s.orderRepository.Invalidated(purchaseOrder.ID)
}

func (s *PurchaseServiceImpl) GetPurchaseReturnOrderList() ([]*domain.PurchaseReturnOrder, error) {
	return s.returnRepository.GetAll()
}

func (s *PurchaseServiceImpl) AddPurchaseReturnOrder(purchaseReturnOrder *domain.PurchaseReturnOrder) error {
	if err := s.returnRepository.Save(*purchaseReturnOrder); err != nil {
		return err
	}
	p := purchaseReturnOrder
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

func (s *PurchaseServiceImpl) InvalidatePurchaseReturnOrder(purchaseReturnOrder *domain.PurchaseReturnOrder) error {
	return s.returnRepository.Invalidated(purchaseReturnOrder.ID)
}
