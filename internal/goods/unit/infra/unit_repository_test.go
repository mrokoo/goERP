package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	repository "github.com/mrokoo/goERP/internal/goods/unit/infra"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewUnitRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repository.NewUnitRepository(db)
}

func TestUnitRepository_Save(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewUnitRepository(db)

	unit := domain.Unit{
		ID:   uuid.New(),
		Name: "个",
	}
	if err := r.Save(&unit); err != nil {
		assert.NoError(t, err)
	}
}
