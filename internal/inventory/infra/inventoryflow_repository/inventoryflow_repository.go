package inventoryflow_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flowrecord"
	"gorm.io/gorm"
)

type InventoryFlowRepository struct {
	db *gorm.DB
}

// The function returns a new instance of an InventoryFlowRepository with a given database connection.
func NewInventoryFlowRepository(db *gorm.DB) *InventoryFlowRepository {
	db.AutoMigrate(&flowrecord.InventoryFlow{})
	return &InventoryFlowRepository{
		db: db,
	}
}

// These are methods of the `InventoryFlowRepository` struct that interact with the database to perform
// CRUD operations on `flowrecord.InventoryFlow` entities.
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
