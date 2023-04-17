package inventoryflow_repository

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewInventoryFlowRepository(t *testing.T) {
	// 帮我写这个测试
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	NewInventoryFlowRepository(db)
}
