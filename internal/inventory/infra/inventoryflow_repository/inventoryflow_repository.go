package inventoryflow_repository

import (
	flowrecord "github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"gorm.io/gorm"
)

type InventoryFlowRepository struct {
	db *gorm.DB
}

func NewInventoryFlowRepository(db *gorm.DB) *InventoryFlowRepository {
	db.AutoMigrate(&MySQLInventoryFlow{})
	return &InventoryFlowRepository{
		db: db,
	}
}

func (r InventoryFlowRepository) GetAll() ([]*flowrecord.InventoryFlow, error) {
	var flows []MySQLInventoryFlow
	result := r.db.Find(&flows)
	if err := result.Error; err != nil {
		return nil, err
	}
	flows_ := make([]*flowrecord.InventoryFlow, len(flows))
	for i, flow := range flows {
		p := flow.toInventoryFlow()
		flows_[i] = &p
	}
	return flows_, nil
}

func (r InventoryFlowRepository) GetByID(ID string) (*flowrecord.InventoryFlow, error) {
	var flow MySQLInventoryFlow
	result := r.db.First(&flow, ID)
	if err := result.Error; err != nil {
		return nil, err
	}
	flow_ := flow.toInventoryFlow()
	return &flow_, nil
}

func (r InventoryFlowRepository) Save(flowRecord *flowrecord.InventoryFlow) error {
	flow := toMySQLInventoryFlow(*flowRecord)
	result := r.db.Create(flow)
	return result.Error
}

func (r InventoryFlowRepository) GetByProductIDAndWarehouseID(productID string, warehouseID string) (*flowrecord.InventoryFlow, error) {
	var flow MySQLInventoryFlow
	result := r.db.Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).Order("date desc").First(&flow)
	if err := result.Error; err != nil {
		return nil, err
	}
	flow_ := flow.toInventoryFlow()
	return &flow_, nil
}
