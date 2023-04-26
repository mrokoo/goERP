package infra

import (
	"testing"

	"github.com/mrokoo/goERP/internal/sale/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewSaleOrderRepository(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewSaleOrderRepository(db)
}

func TestSaleOrderRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewSaleOrderRepository(db)
	item1 := domain.NewItem("P002", 10, 10.0)
	order := domain.NewSaleOrder("S003", "W001", "bA55586B-1846-b2ec-dbFd-2C1Fde73bE4f", "U001", "2022-01-01", "", []domain.Item{item1}, domain.Order)
	err = r.Save(&order)
	if err != nil {
		t.Error(err)
	}
}

func TestSaleOrderRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewSaleOrderRepository(db)
	orders, err := r.GetAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(orders)
}

func TestSaleOrderRepository_GetByID(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewSaleOrderRepository(db)
	order, err := r.GetByID("S003")
	if err != nil {
		t.Error(err)
	}

	t.Log(*order)
}
