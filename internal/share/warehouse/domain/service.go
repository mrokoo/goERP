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

// IsValidated 检查仓库 ID 的唯一性。
//
// 参数:
// - warehouse: 指向 Warehouse 结构体的指针。
//
// 返回值:
// - bool: 布尔值，表示仓库 ID 是否唯一。
func (ds *CheckingWarehouseValidityService) IsValidated(warehouse *Warehouse) bool {
	// ID唯一性校验
	_, err := ds.warehouseRepository.GetByID(warehouse.ID)
	return err == repository.ErrNotFound
}
