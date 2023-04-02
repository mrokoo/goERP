package domain

import "go.mongodb.org/mongo-driver/mongo"

type CheckingSupplierValidityService struct {
	supplierRepository Repository
}

func NewCheckingSupplierValidityService(supplierRepository Repository) *CheckingSupplierValidityService {
	return &CheckingSupplierValidityService{
		supplierRepository: supplierRepository,
	}
}

func (ds *CheckingSupplierValidityService) IsValidated(supplier *Supplier) bool {
	// ID唯一性校验
	_, err := ds.supplierRepository.GetByID(supplier.ID)
	return err == mongo.ErrNoDocuments
}
