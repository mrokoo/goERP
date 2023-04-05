package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/customer/domain"
)

var ErrCustomerInVaildated = errors.New("客户ID检验无效")

type CustomerService interface {
	GetCustomer(customerID string) (*domain.Customer, error)
	GetCustomerList() ([]*domain.Customer, error)
	AddCustomer(customer *domain.Customer) error
	ReplaceCustomer(customer *domain.Customer) error
	DeleteCustomer(customerID string) error
}

type CustomerServiceImpl struct {
	checkCustomerValidityService *domain.CheckingCustomerValidityService
	repo                         domain.Repository
}

func NewCustomerServiceImpl(checkCustomerValidityService *domain.CheckingCustomerValidityService, repo domain.Repository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		checkCustomerValidityService: checkCustomerValidityService,
		repo:                         repo,
	}
}

func (s *CustomerServiceImpl) GetCustomer(customerID string) (*domain.Customer, error) {
	customer, err := s.repo.GetByID(customerID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerServiceImpl) GetCustomerList() ([]*domain.Customer, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *CustomerServiceImpl) AddCustomer(customer *domain.Customer) error {

	if !s.checkCustomerValidityService.IsValidated(customer) {
		return ErrCustomerInVaildated
	}
	err := s.repo.Save(customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerServiceImpl) ReplaceCustomer(customer *domain.Customer) error {
	if err := s.repo.Replace(customer); err != nil {
		return err
	}
	return nil
}

func (s *CustomerServiceImpl) DeleteCustomer(customerID string) error {
	if err := s.repo.Delete(customerID); err != nil {
		return err
	}
	return nil
}
