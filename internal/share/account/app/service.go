package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/account/domain"
)

type AccountService interface {
	GetAccount(accountId domain.AccountId) (*domain.Account, error)
	GetAccountList() ([]domain.Account, error)
	AddAccount(account domain.Account) error
	UpdateAccount(account domain.Account) error
	DeleteAccount(accountId domain.AccountId) error
}

type AccountServiceImpl struct {
	checkAccountValidityService *domain.CheckingAccountValidityService
	repo                        domain.Repository
}

func NewAccountServiceImpl(checkAccountValidityService *domain.CheckingAccountValidityService, repo domain.Repository) *AccountServiceImpl {
	return &AccountServiceImpl{
		repo: repo,
	}
}

func (s *AccountServiceImpl) GetAccount(accountId domain.AccountId) (*domain.Account, error) {
	account, err := s.repo.Get(accountId)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountServiceImpl) GetAccountList() ([]domain.Account, error) {
	accounts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountServiceImpl) AddAccount(account domain.Account) error {
	// account验证逻辑
	if !s.checkAccountValidityService.IsValidated(account) {
		return errors.New("the validity check fails")
	}
	err := s.repo.Save(account)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) UpdateAccount(account domain.Account) error {
	if !s.checkAccountValidityService.IsValidated(account) {
		return errors.New("the validity check fails")
	}
	if err := s.repo.Update(account); err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) DeleteAccount(accountId domain.AccountId) error {
	if err := s.repo.Delete(accountId); err != nil {
		return err
	}
	return nil
}
