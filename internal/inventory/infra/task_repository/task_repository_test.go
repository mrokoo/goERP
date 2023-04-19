package task_repository

import (
	"testing"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewTaskRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewTaskRepository(db)
}

func TestTaskRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewTaskRepository(db)
	items := []task.TaskItem{
		task.NewTaskItem("P002", 200),
	}
	task_ := task.NewTask("W001", task.IN_PURCHASE, "PO001", items)
	it := record.NewRecordItem("P002", 100)
	r := record.NewRecord("W001", "U001", []record.RecordItem{it})
	task_.AddRecord(r)
	if err := repo.Save(&task_); err != nil {
		t.Error(err)
	}
}

func TestTaskRepository_Save2(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewTaskRepository(db)

	task_, err := repo.GetByID("bfeadef3-6bab-4518-b004-9da5cbed56ee")
	if err != nil {
		assert.Error(err)
	}
	it := record.NewRecordItem("P002", 100)
	r := record.NewRecord("W001", "U001", []record.RecordItem{it})
	task_.AddRecord(r)
	if err := repo.Save(task_); err != nil {
		assert.Error(err)
	}
}

func TestTaskRepository_GetByID(t *testing.T) {
	id := "867f5482-c535-4f32-92d0-e483b01be063"
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewTaskRepository(db)
	task, err := repo.GetByID(id)
	if err != nil {
		t.Error(err)
	}
	t.Log(*task)
}

func TestTaskRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewTaskRepository(db)
	tasks, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(tasks)
}
