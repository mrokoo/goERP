package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	repository "github.com/mrokoo/goERP/internal/goods/product/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	}
	if err := r.Save(&product); err != nil {
		t.Error(err)
	}
}
