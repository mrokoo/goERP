package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetAll() ([]*domain.Account, error) {
	var list []*model.Account
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var accounts []*domain.Account
	for i := range list {
		accounts = append(accounts, toDmain(list[i]))
	}
	return accounts, nil
}

func (r *AccountRepository) GetByID(ID string) (*domain.Account, error) {
	i := model.Account{
		ID: ID,
	}
	result := r.db.First(&i)
	if err := result.Error; err != nil {
		return nil, err
	}
	account := toDmain(&i)
	return account, nil
}

func (r *AccountRepository) Save(account *domain.Account) error {
	a := toModel(account)
	result := r.db.Create(a)
	return result.Error
}

func (r *AccountRepository) Replace(account *domain.Account) error {
	a := toModel(account)
	result := r.db.Save(a)
	return result.Error
}

func (r *AccountRepository) Delete(accountID string) error {
	result := r.db.Delete(&model.Account{
		ID: accountID,
	})
	return result.Error
}
