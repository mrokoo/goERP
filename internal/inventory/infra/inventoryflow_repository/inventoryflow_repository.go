package inventoryflow_repository

import (
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

type InventoryFlowRepository struct {
	db *gorm.DB
}

func NewInventoryFlowRepository(db *gorm.DB) *InventoryFlowRepository {
	return &InventoryFlowRepository{
		db: db,
	}
}

func (r InventoryFlowRepository) GetAll() ([]*flowrecord.InventoryFlow, error) {
	var list []*model.InventoryFlow
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var flows []*flowrecord.InventoryFlow
	for i := range list {
		flows = append(flows, toDomain(list[i]))
	}
	return flows, nil
}

func (r InventoryFlowRepository) GetByID(ID string) (*flowrecord.InventoryFlow, error) {
	var flow model.InventoryFlow
	result := r.db.First(&flow, ID)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&flow), nil
}

func (r InventoryFlowRepository) Save(flowRecord *flowrecord.InventoryFlow) error {
	flow := toModel(flowRecord)
	result := r.db.Create(flow)
	return result.Error
}

func (r InventoryFlowRepository) GetByProductIDAndWarehouseID(productID string, warehouseID string) (*flowrecord.InventoryFlow, error) {
	var flow model.InventoryFlow
	result := r.db.Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).Order("created_at desc").First(&flow)
	if err := result.Error; err != nil {
		return nil, err
	}

	return toDomain(&flow), nil
}
