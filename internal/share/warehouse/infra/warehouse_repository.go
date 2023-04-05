package repository

import (
	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Warehouse = domain.Warehouse

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	db.AutoMigrate(&Warehouse{}) //自动迁移
	return &WarehouseRepository{
		db: db,
	}
}

func (r *WarehouseRepository) GetAll() ([]*Warehouse, error) {
	var warehouses []Warehouse
	result := r.db.Find(&warehouses)
	if err := result.Error; err != nil {
		return nil, err
	}
	var warehousesp []*Warehouse
	for _, v := range warehouses {
		warehousesp = append(warehousesp, &v)
	}
	return warehousesp, nil
}

func (r *WarehouseRepository) GetByID(warehouseID string) (*Warehouse, error) {
	warehouse := Warehouse{
		ID: warehouseID,
	}
	result := r.db.First(&warehouse)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *WarehouseRepository) Save(warehouse *domain.Warehouse) error {
	result := r.db.Create(warehouse)
	return result.Error
}

func (r *WarehouseRepository) Replace(warehouse *domain.Warehouse) error {
	result := r.db.Save(warehouse)
	return result.Error
}

func (r *WarehouseRepository) Delete(warehouseID string) error {
	result := r.db.Delete(&Warehouse{
		ID: warehouseID,
	})
	return result.Error
}
