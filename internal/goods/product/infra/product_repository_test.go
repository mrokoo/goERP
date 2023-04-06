package repository_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	repository "github.com/mrokoo/goERP/internal/goods/product/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewProductRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repository.NewProductRepository(db)
}

func TestProductRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewProductRepository(db)
	product := domain.Product{
		ID:     "P001",
		Name:   "chanpin1",
		UnitID: uuid.MustParse("2ce0d254-7ce1-4257-b3be-4a9dc08218e3"),
		OpeningStock: []repository.Stock{
			{
				ProductID:   "P001",
				WarehouseID: "W001",
				Amount:      100,
			},
			{
				ProductID:   "P001",
				WarehouseID: "W002",
				Amount:      23,
			},
		},
	}
	if err := r.Save(&product); err != nil {
		t.Error(err)
	}
}

func TestProductRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewProductRepository(db)

	products, err := r.GetAll()
	if err != nil {
		t.Error(err)
	}
	for _, p := range products {
		fmt.Printf("%#v", p)
	}
}
