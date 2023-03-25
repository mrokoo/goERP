package domain

import "go.mongodb.org/mongo-driver/mongo"

type CheckingCustomerValidityService struct {
	customerRepository Repository
}

func NewCheckingCustomerValidityService(customerRepository Repository) *CheckingCustomerValidityService {
	return &CheckingCustomerValidityService{
		customerRepository: customerRepository,
	}
}

func (ds *CheckingCustomerValidityService) IsValidated(customer Customer) bool {
	// ID唯一性校验
	_, err := ds.customerRepository.Get(customer.ID)
	return err == mongo.ErrNoDocuments
}
