package repository_test

import (
	"testing"

	repository "github.com/mrokoo/goERP/internal/system/user/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repository.NewUserRepository(db)
}
