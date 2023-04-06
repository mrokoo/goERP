package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	repository "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var r *repository.ProductRepository

func TestMain(m *testing.M) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r = repository.NewProductRepository(db)
}
func TestNewProductRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repository.NewProductRepository(db)
}

func TestProductRepository_Save(t *testing.T) {
	product := domain.Product{
		ID:     "P001",
		Name:   "chanpin1",
		UnitID: uuid.MustParse("2ce0d254-7ce1-4257-b3be-4a9dc08218e3"),
		OpeningStock: []stock.Stock{
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
	err := r.Save(&product)
	assert := assert.New(t)
	assert.NoError(err)
}

func TestProductRepository_GetAll(t *testing.T) {
	_, err := r.GetAll()
	assert := assert.New(t)
	assert.NoError(err)
}

func TestProductRepository_Delete(t *testing.T) {
	assert := assert.New(t)
	err := r.Delete("P001")
	assert.NoError(err)
}
