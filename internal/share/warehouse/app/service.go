package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

var ErrWarehouseInVaildated = errors.New("仓库ID检验无效")

type WarehouseService interface {
	GetWarehouse(warehouseID string) (*domain.Warehouse, error)
	GetWarehouseList() ([]*domain.Warehouse, error)
	AddWarehouse(warehouse *domain.Warehouse) error
	ReplaceWarehouse(warehouse *domain.Warehouse) error
	DeleteWarehouse(warehouseID string) error
}

type WarehouseServiceImpl struct {
	checkWarehouseValidityService *domain.CheckingWarehouseValidityService
	repo                          domain.Repository
}

func NewWarehouseServiceImpl(checkWarehouseValidityService *domain.CheckingWarehouseValidityService, repo domain.Repository) *WarehouseServiceImpl {
	return &WarehouseServiceImpl{
		checkWarehouseValidityService: checkWarehouseValidityService,
		repo:                          repo,
	}
}

func (s *WarehouseServiceImpl) GetWarehouse(warehouseID string) (*domain.Warehouse, error) {
	warehouse, err := s.repo.GetByID(warehouseID)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (s *WarehouseServiceImpl) GetWarehouseList() ([]*domain.Warehouse, error) {
	warehouses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (s *WarehouseServiceImpl) AddWarehouse(warehouse *domain.Warehouse) error {

	if !s.checkWarehouseValidityService.IsValidated(warehouse) {
		return ErrWarehouseInVaildated
	}
	err := s.repo.Save(warehouse)
	if err != nil {
		return err
	}
	return nil
}

func (s *WarehouseServiceImpl) ReplaceWarehouse(warehouse *domain.Warehouse) error {
	if err := s.repo.Replace(warehouse); err != nil {
		return err
	}
	return nil
}

func (s *WarehouseServiceImpl) DeleteWarehouse(warehouseID string) error {
	if err := s.repo.Delete(warehouseID); err != nil {
		return err
	}
	return nil
}
