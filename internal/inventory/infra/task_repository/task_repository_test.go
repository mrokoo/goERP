package task_repository

import (
	"testing"
	"time"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/item"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/record"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewInTaskRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewInTaskRepository(db)
}

func TestNewOutTaskRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewOutTaskRepository(db)
}

func TestInTaskRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewInTaskRepository(db)
	inTasks, err := r.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(inTasks)
}

func TestInTaskRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewInTaskRepository(db)
	// 帮我写一个task.Task的测试用例
	inTask := task.InTask{
		ID:   "test",
		Type: "test",
		InBasis: task.InBasis{
			PurchaseOrderID: "PO001",
		},
		State:  "test",
		Status: "test",
		Records: []record.InRecord{
			{
				UserID: "U001",
				Date:   time.Now(),
				Note:   "test",
				Items: []item.InItem{
					{
						ProductID: "P002",
						Quantity:  20,
					},
				},
			},
		},
	}
	err = r.Save(&inTask)
	if err != nil {
		t.Fatal(err)
	}
}
