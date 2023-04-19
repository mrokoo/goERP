package allot_repository

import (
	"testing"

	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAllotRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewAllotRepository(db)
	a := allot.NewAllot("W001", "W002", "U001", []allot.Item{
		{
			ProductID: "P002",
			Quantity:  10,
		},
	})
	err = repo.Save(a)
	if err != nil {
		t.Error(err)
	}
}

func TestAllotRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo := NewAllotRepository(db)
	allots, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(allots)
}
