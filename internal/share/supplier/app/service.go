package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/supplier/domain"
)

var ErrSupplierInVaildated = errors.New("供应商ID检验无效")

type SupplierService interface {
	GetSupplier(supplierID string) (*domain.Supplier, error)
	GetSupplierList() ([]*domain.Supplier, error)
	AddSupplier(supplier *domain.Supplier) error
	ReplaceSupplier(supplier *domain.Supplier) error
	DeleteSupplier(supplierID string) error
}

type SupplierServiceImpl struct {
	checkSupplierValidityService *domain.CheckingSupplierValidityService
	repo                         domain.Repository
}

func NewSupplierServiceImpl(checkSupplierValidityService *domain.CheckingSupplierValidityService, repo domain.Repository) *SupplierServiceImpl {
	return &SupplierServiceImpl{
		checkSupplierValidityService: checkSupplierValidityService,
		repo:                         repo,
	}
}

func (s *SupplierServiceImpl) GetSupplier(supplierID string) (*domain.Supplier, error) {
	supplier, err := s.repo.GetByID(supplierID)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierServiceImpl) GetSupplierList() ([]*domain.Supplier, error) {
	suppliers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *SupplierServiceImpl) AddSupplier(supplier *domain.Supplier) error {

	if !s.checkSupplierValidityService.IsValidated(supplier) {
		return ErrSupplierInVaildated
	}
	err := s.repo.Save(supplier)
	if err != nil {
		return err
	}
	return nil
}

func (s *SupplierServiceImpl) ReplaceSupplier(supplier *domain.Supplier) error {
	if err := s.repo.Replace(supplier); err != nil {
		return err
	}
	return nil
}

func (s *SupplierServiceImpl) DeleteSupplier(supplierID string) error {
	if err := s.repo.Delete(supplierID); err != nil {
		return err
	}
	return nil
}
