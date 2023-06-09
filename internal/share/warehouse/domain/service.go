package domain

import (
	repository "github.com/mrokoo/goERP/internal/share/supplier/infra"
)

type CheckingWarehouseValidityService struct {
	warehouseRepository Repository
}

func NewCheckingWarehouseValidityService(warehouseRepository Repository) *CheckingWarehouseValidityService {
	return &CheckingWarehouseValidityService{
		warehouseRepository: warehouseRepository,
	}
}

func (ds *CheckingWarehouseValidityService) IsValidated(warehouse *Warehouse) bool {
	// ID唯一性校验
	_, err := ds.warehouseRepository.GetByID(warehouse.ID)
	return err == repository.ErrNotFound
}
