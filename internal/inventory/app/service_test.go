package app

import (
	"testing"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	allotRepo "github.com/mrokoo/goERP/internal/inventory/infra/allot_repository"
	flowRepo "github.com/mrokoo/goERP/internal/inventory/infra/inventoryflow_repository"
	takeRepo "github.com/mrokoo/goERP/internal/inventory/infra/take_repository"
	taskRepo "github.com/mrokoo/goERP/internal/inventory/infra/task_repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInventoryServiceImpl_GetTaskList(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := NewInventoryServiceImpl(flow, task, allot, take)
	list, err := m3.GetTaskList()
	if err != nil {
		t.Log(list)
	}
}

func TestInventoryServiceImpl_AddRecord(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := NewInventoryServiceImpl(flow, task, allot, take)
	err = m3.AddRecord("d55c4283-4d01-49ea-a1af-bb04b2f5e6bd", record.NewRecord("W002", "U001", []record.RecordItem{
		record.NewRecordItem("P001", 1),
		record.NewRecordItem("P002", 2),
	}))
	if err != nil {
		t.Error(err)
	}
}
