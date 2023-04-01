package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

var ErrNotFound = errors.New("the docment is not found")
var ErrWarehouseInVaildated = errors.New("the validity check fails")

type WarehouseService interface {
	GetWarehouse(warehouseId domain.WarehouseId) (*domain.Warehouse, error)
	GetWarehouseList() ([]domain.Warehouse, error)
	AddWarehouse(warehouse domain.Warehouse) error
	UpdateWarehouse(warehouse domain.Warehouse) error
	DeleteWarehouse(warehouseId domain.WarehouseId) error
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

func (s *WarehouseServiceImpl) GetWarehouse(warehouseId domain.WarehouseId) (*domain.Warehouse, error) {
	warehouse, err := s.repo.Get(warehouseId)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (s *WarehouseServiceImpl) GetWarehouseList() ([]domain.Warehouse, error) {
	warehouses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (s *WarehouseServiceImpl) AddWarehouse(warehouse domain.Warehouse) error {
	// 检查Warehouse是否符合要求
	if !s.checkWarehouseValidityService.IsValidated(warehouse) {
		return ErrWarehouseInVaildated
	}
	err := s.repo.Save(warehouse)
	if err != nil {
		return err
	}
	return nil
}

func (s *WarehouseServiceImpl) UpdateWarehouse(warehouse domain.Warehouse) error {
	if err := s.repo.Update(warehouse); err != nil {
		return err
	}
	return nil
}

func (s *WarehouseServiceImpl) DeleteWarehouse(warehouseId domain.WarehouseId) error {
	if err := s.repo.Delete(warehouseId); err != nil {
		return err
	}
	return nil
}
