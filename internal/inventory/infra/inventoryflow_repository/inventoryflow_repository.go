package inventoryflow_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flowrecord"
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
	var flows []InventoryFlow
	result := r.db.Find(&flows)
	if err := result.Error; err != nil {
		return nil, err
	}
	var flows_ []*flowrecord.InventoryFlow
	for _, f := range flows {
		flows_ = append(flows_, f.toInventoryFlow())
	}
	return flows_, nil
}

func (r InventoryFlowRepository) GetByID(ID string) (*flowrecord.InventoryFlow, error) {
	flow := InventoryFlow{
		ID: ID,
	}
	result := r.db.First(&flow)
	if err := result.Error; err != nil {
		return nil, err
	}
	return flow.toInventoryFlow(), nil
}

func (r InventoryFlowRepository) Save(flowRecord *flowrecord.InventoryFlow) error {
	flow := toMySQLInventoryFlow(flowRecord)
	result := r.db.Create(flow)
	return result.Error
}
