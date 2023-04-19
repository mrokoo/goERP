package inventoryflow_repository

import (
	"testing"

	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInventoryFlowRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewInventoryFlowRepository(db)
	flows, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(flows)
}

func TestInventoryFlowRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewInventoryFlowRepository(db)
	flow := flowrecord.NewInventoryFlow("bfeadef3-6bab-4518-b004-9da5cbed56ee", "P002", "W001", flowrecord.FLOWTYPE_RUKU, 10, 10)
	err = repo.Save(&flow)
	if err != nil {
		t.Error(err)
	}
}
