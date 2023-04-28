package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) GetAll() ([]*domain.Supplier, error) {
	var list []*model.Supplier
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var suppliers []*domain.Supplier
	for i := range list {
		suppliers = append(suppliers, toDomain(list[i]))
	}
	return suppliers, nil
}

func (r *SupplierRepository) GetByID(ID string) (*domain.Supplier, error) {
	supplier := model.Supplier{
		ID: ID,
	}
	result := r.db.First(&supplier)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&supplier), nil
}

func (r *SupplierRepository) Save(supplier *domain.Supplier) error {
	i := toModel(supplier)
	result := r.db.Create(i)
	return result.Error
}

func (r *SupplierRepository) Replace(supplier *domain.Supplier) error {
	i := toModel(supplier)
	result := r.db.Save(i)
	return result.Error
}

func (r *SupplierRepository) Delete(ID string) error {
	result := r.db.Delete(&model.Supplier{
		ID: ID,
	})
	return result.Error
}
