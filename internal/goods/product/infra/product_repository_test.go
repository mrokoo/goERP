package repository

import (
	"testing"

	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestProductRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	repo := NewProductRepository(db)
	unit := new(string)
	*unit = "U001"
	product := domain.Product{
		ID:         "P002",
		Name:       "产品2",
		CategoryID: nil,
		UnitID:     unit,
		OpeningStock: []stock.Stock{
			{
				WarehouseID: "W001",
				Amount:      10,
			},
			{
				WarehouseID: "W002",
				Amount:      20,
			},
		},
	}

	err = repo.Save(&product)
	if err != nil {
		t.Error(err)
	}
}

func TestProductRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	repo := NewProductRepository(db)

	list, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}
