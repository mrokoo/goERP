package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/account/domain"
)

var ErrAccountInVaildated = errors.New("账户ID检验无效")

type AccountService interface {
	GetAccount(accountID string) (*domain.Account, error)
	GetAccountList() ([]*domain.Account, error)
	AddAccount(account *domain.Account) error
	ReplaceAccount(account *domain.Account) error
	DeleteAccount(accountID string) error
}

type AccountServiceImpl struct {
	checkAccountValidityService *domain.CheckingAccountValidityService
	repo                        domain.Repository
}

func NewAccountServiceImpl(checkAccountValidityService *domain.CheckingAccountValidityService, repo domain.Repository) *AccountServiceImpl {
	return &AccountServiceImpl{
		checkAccountValidityService: checkAccountValidityService,
		repo:                        repo,
	}
}

func (s *AccountServiceImpl) GetAccount(accountID string) (*domain.Account, error) {
	account, err := s.repo.GetByID(accountID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountServiceImpl) GetAccountList() ([]*domain.Account, error) {
	accounts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountServiceImpl) AddAccount(account *domain.Account) error {

	if !s.checkAccountValidityService.IsValidated(account) {
		return ErrAccountInVaildated
	}
	err := s.repo.Save(account)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) ReplaceAccount(account *domain.Account) error {
	if err := s.repo.Replace(account); err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) DeleteAccount(accountID string) error {
	if err := s.repo.Delete(accountID); err != nil {
		return err
	}
	return nil
}
