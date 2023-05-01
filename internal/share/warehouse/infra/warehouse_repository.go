package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}

func (r *WarehouseRepository) GetAll() ([]*domain.Warehouse, error) {
	var list []*model.Warehouse
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var warehouses []*domain.Warehouse
	for i := range list {
		warehouses = append(warehouses, toDomain(list[i]))
	}
	return warehouses, nil
}

func (r *WarehouseRepository) GetByID(ID string) (*domain.Warehouse, error) {
	warehouse := model.Warehouse{
		ID: ID,
	}
	result := r.db.First(&warehouse)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&warehouse), nil
}

func (r *WarehouseRepository) Save(warehouse *domain.Warehouse) error {
	i := toModel(warehouse)
	result := r.db.Create(i)
	return result.Error
}

func (r *WarehouseRepository) Replace(warehouse *domain.Warehouse) error {
	i := toModel(warehouse)
	result := r.db.Save(i)
	return result.Error
}

func (r *WarehouseRepository) Delete(ID string) error {
	result := r.db.Delete(&model.Warehouse{
		ID: ID,
	})
	return result.Error
}
