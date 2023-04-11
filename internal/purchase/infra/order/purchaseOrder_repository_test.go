package order

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrokoo/goERP/internal/purchase/domain"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/item"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewProductRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewPurchaseOrderRepository(db)
}

func TestPurchaseOrderRepository_Save(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewPurchaseOrderRepository(db)

	purchaseOrder := domain.PurchaseOrder{
		ID:          "PO001",
		SupplierID:  "S001",
		WarehouseID: "W001",
		UserID:      "U001",
		Date:        time.Now(),
		Items: []item.OrderItem{
			item.NewOrderItem("P002", 12, 3.21),
		},
	}
	err = r.Save(&purchaseOrder)
	assert.NoError(err)
}

func TestPurchaseOrderRepository_GetAll(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewPurchaseOrderRepository(db)
	_, err = r.GetAll()
	assert.NoError(err)
}

func TestPurchaseOrderRepository_GetByID(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewPurchaseOrderRepository(db)
	d, err := r.GetByID("PO001")
	fmt.Printf("d: %v\n", d)
	assert.NoError(err)
}

func TestPurchaseOrderRepository_InValidate(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewPurchaseOrderRepository(db)
	err = r.InValidate("PO001")
	assert.NoError(err)
}
