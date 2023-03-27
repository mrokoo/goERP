package domain

import "go.mongodb.org/mongo-driver/mongo"

type CheckingWarehouseValidityService struct {
	warehouseRepository Repository
}

func NewCheckingWarehouseValidityService(warehouseRepository Repository) *CheckingWarehouseValidityService {
	return &CheckingWarehouseValidityService{
		warehouseRepository: warehouseRepository,
	}
}

func (ds *CheckingWarehouseValidityService) IsValidated(warehouse Warehouse) bool {
	// ID唯一性校验
	_, err := ds.warehouseRepository.Get(warehouse.ID)
	return err == mongo.ErrNoDocuments
}
