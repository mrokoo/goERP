package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAutoMigrate(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = AutoMigrate(db)
	if err != nil {
		t.Error(err)
	}
}
