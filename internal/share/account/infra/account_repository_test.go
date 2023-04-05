package repository_test

import (
	"fmt"
	"testing"

	"github.com/mrokoo/goERP/internal/share/account/domain"
	repository "github.com/mrokoo/goERP/internal/share/account/infra"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAccountRepository_Save(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewAccountRepository(db)
	account := domain.Account{
		ID:   "A120",
		Name: "wang",
	}
	if err := r.Save(&account); err != nil {
		t.Error(err)
	}
}
func TestAccountRepository_GetAll(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewAccountRepository(db)
	accounts, err := r.GetAll()
	fmt.Printf("accounts: %v\n", accounts)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountRepository_GetByID(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewAccountRepository(db)

	t.Run("not found record", func(t *testing.T) {
		_, err := r.GetByID("A121")
		assert.ErrorIs(err, repository.ErrNotFound)
	})

	t.Run("normal", func(t *testing.T) {
		account, _ := r.GetByID("A120")
		expect_account := domain.Account{
			ID:    "A120",
			Name:  "wang",
			Type:  "other",
			State: "active",
		}
		assert.Equal(&expect_account, account)
	})
}

func TestAccountRepository_Replace(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewAccountRepository(db)
	t.Run("normal", func(t *testing.T) {
		expect_account := domain.Account{
			ID:    "A120",
			Name:  "li",
			Type:  "cash",
			State: "active",
		}
		err := r.Replace(&expect_account)
		assert.NoError(err)
	})
}

func TestAccountRepository_Delete(t *testing.T) {
	assert := assert.New(t)
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := repository.NewAccountRepository(db)
	t.Run("normal", func(t *testing.T) {
		err := r.Delete("A120")
		assert.NoError(err)
	})
}
