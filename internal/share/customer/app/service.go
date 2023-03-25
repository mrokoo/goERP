package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/customer/domain"
)

var ErrNotFound = errors.New("the docment is not found")
var ErrCustomerInVaildated = errors.New("the validity check fails")

type CustomerService interface {
	GetCustomer(customerId domain.CustomerId) (*domain.Customer, error)
	GetCustomerList() ([]domain.Customer, error)
	AddCustomer(customer domain.Customer) error
	UpdateCustomer(customer domain.Customer) error
	DeleteCustomer(customerId domain.CustomerId) error
}

type CustomerServiceImpl struct {
	checkCustomerValidityService *domain.CheckingCustomerValidityService
	repo                        domain.Repository
}

func NewCustomerServiceImpl(checkCustomerValidityService *domain.CheckingCustomerValidityService, repo domain.Repository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		checkCustomerValidityService: checkCustomerValidityService,
		repo:                        repo,
	}
}

func (s *CustomerServiceImpl) GetCustomer(customerId domain.CustomerId) (*domain.Customer, error) {
	customer, err := s.repo.Get(customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerServiceImpl) GetCustomerList() ([]domain.Customer, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *CustomerServiceImpl) AddCustomer(customer domain.Customer) error {
	// 检查Customer是否符合要求
	if !s.checkCustomerValidityService.IsValidated(customer) {
		return ErrCustomerInVaildated
	}
	err := s.repo.Save(customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerServiceImpl) UpdateCustomer(customer domain.Customer) error {
	if err := s.repo.Update(customer); err != nil {
		return err
	}
	return nil
}

func (s *CustomerServiceImpl) DeleteCustomer(customerId domain.CustomerId) error {
	if err := s.repo.Delete(customerId); err != nil {
		return err
	}
	return nil
}
