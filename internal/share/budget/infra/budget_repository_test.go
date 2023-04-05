package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestBudgetRepository_Save(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := NewBudgetRepository(db)
	id, _ := uuid.NewUUID()
	budget := &domain.Budget{
		ID:   id,
		Type: "in",
	}
	err = r.Save(budget)
	assert.NoError(err)
}
