package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/supplier/domain"
)

var ErrNotFound = errors.New("the docment is not found")
var ErrSupplierInVaildated = errors.New("the validity check fails")

type SupplierService interface {
	GetSupplier(supplierId domain.SupplierId) (*domain.Supplier, error)
	GetSupplierList() ([]domain.Supplier, error)
	AddSupplier(supplier domain.Supplier) error
	UpdateSupplier(supplier domain.Supplier) error
	DeleteSupplier(supplierId domain.SupplierId) error
}

type SupplierServiceImpl struct {
	checkSupplierValidityService *domain.CheckingSupplierValidityService
	repo                        domain.Repository
}

func NewSupplierServiceImpl(checkSupplierValidityService *domain.CheckingSupplierValidityService, repo domain.Repository) *SupplierServiceImpl {
	return &SupplierServiceImpl{
		checkSupplierValidityService: checkSupplierValidityService,
		repo:                        repo,
	}
}

func (s *SupplierServiceImpl) GetSupplier(supplierId domain.SupplierId) (*domain.Supplier, error) {
	supplier, err := s.repo.Get(supplierId)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierServiceImpl) GetSupplierList() ([]domain.Supplier, error) {
	suppliers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *SupplierServiceImpl) AddSupplier(supplier domain.Supplier) error {
	// 检查Supplier是否符合要求
	if !s.checkSupplierValidityService.IsValidated(supplier) {
		return ErrSupplierInVaildated
	}
	err := s.repo.Save(supplier)
	if err != nil {
		return err
	}
	return nil
}

func (s *SupplierServiceImpl) UpdateSupplier(supplier domain.Supplier) error {
	if err := s.repo.Update(supplier); err != nil {
		return err
	}
	return nil
}

func (s *SupplierServiceImpl) DeleteSupplier(supplierId domain.SupplierId) error {
	if err := s.repo.Delete(supplierId); err != nil {
		return err
	}
	return nil
}
