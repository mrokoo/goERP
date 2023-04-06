package repository

import (
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Account = domain.Account

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	db.AutoMigrate(&Account{}) //自动迁移
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetAll() ([]*Account, error) {
	var accounts []Account
	result := r.db.Find(&accounts)
	if err := result.Error; err != nil {
		return nil, err
	}
	var accountsp []*Account
	for i := range accounts {
		accountsp = append(accountsp, &accounts[i])
	}
	return accountsp, nil
}

func (r *AccountRepository) GetByID(accountID string) (*Account, error) {
	account := Account{
		ID: accountID,
	}
	result := r.db.First(&account)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) Save(account *domain.Account) error {
	result := r.db.Create(account)
	return result.Error
}

func (r *AccountRepository) Replace(account *domain.Account) error {
	result := r.db.Save(account)
	return result.Error
}

func (r *AccountRepository) Delete(accountID string) error {
	result := r.db.Delete(&Account{
		ID: accountID,
	})
	return result.Error
}
