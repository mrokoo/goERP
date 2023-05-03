package app

import (
	"testing"

	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	allotRepo "github.com/mrokoo/goERP/internal/inventory/infra/allot_repository"
	flowRepo "github.com/mrokoo/goERP/internal/inventory/infra/inventoryflow_repository"
	takeRepo "github.com/mrokoo/goERP/internal/inventory/infra/take_repository"
	taskRepo "github.com/mrokoo/goERP/internal/inventory/infra/task_repository"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	order "github.com/mrokoo/goERP/internal/purchase/infra"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPurchaseServiceImpl_AddPurchaseOrder(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	m1 := order.NewPurchaseOrderRepository(db)
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := inventory.NewInventoryServiceImpl(flow, task, allot, take)

	s := NewPurchaseServiceImpl(m1, m3)
	order := domain.NewPurchaseOrder("P002", "W002", "S001", "U001", "A001", 0, 123, "", []domain.Item{
		{ProductID: "P001", Quantity: 10},
		{ProductID: "P002", Quantity: 20},
	}, domain.Order)
	if err := s.AddPurchaseOrder(&order); err != nil {
		t.Error(err)
	}
}

func TestPurchaseServiceImpl_GetPurchaseOrderList(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	m1 := order.NewPurchaseOrderRepository(db)
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := inventory.NewInventoryServiceImpl(flow, task, allot, take)

	s := NewPurchaseServiceImpl(m1, m3)
	list, err := s.GetPurchaseOrderList()
	if err != nil {
		t.Log(list)
	}
}

func TestPurchaseServiceImpl_InvalidatePurchaseOrder(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	m1 := order.NewPurchaseOrderRepository(db)
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := inventory.NewInventoryServiceImpl(flow, task, allot, take)

	s := NewPurchaseServiceImpl(m1, m3)
	if err := s.InvalidatePurchaseOrder("P002"); err != nil {
		assert.Error(err)
	}

}
