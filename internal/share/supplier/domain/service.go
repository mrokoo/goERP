package domain

import (
	repository "github.com/mrokoo/goERP/internal/goods/category/infra"
)

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
	return err == repository.ErrNotFound
}
