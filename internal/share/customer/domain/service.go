package domain

import (
	repository "github.com/mrokoo/goERP/internal/goods/category/infra"
)

type CheckingCustomerValidityService struct {
	customerRepository Repository
}

func NewCheckingCustomerValidityService(customerRepository Repository) *CheckingCustomerValidityService {
	return &CheckingCustomerValidityService{
		customerRepository: customerRepository,
	}
}

func (ds *CheckingCustomerValidityService) IsValidated(customer *Customer) bool {
	// ID唯一性校验
	_, err := ds.customerRepository.GetByID(customer.ID)
	return err == repository.ErrNotFound
}
