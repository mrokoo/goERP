package repository

import (
	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Supplier = domain.Supplier

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	db.AutoMigrate(&Supplier{}) //自动迁移
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) GetAll() ([]*Supplier, error) {
	var suppliers []Supplier
	result := r.db.Find(&suppliers)
	if err := result.Error; err != nil {
		return nil, err
	}
	var suppliersp []*Supplier
	for _, v := range suppliers {
		suppliersp = append(suppliersp, &v)
	}
	return suppliersp, nil
}

func (r *SupplierRepository) GetByID(supplierID string) (*Supplier, error) {
	supplier := Supplier{
		ID: supplierID,
	}
	result := r.db.First(&supplier)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *SupplierRepository) Save(supplier *domain.Supplier) error {
	result := r.db.Create(supplier)
	return result.Error
}

func (r *SupplierRepository) Replace(supplier *domain.Supplier) error {
	result := r.db.Save(supplier)
	return result.Error
}

func (r *SupplierRepository) Delete(supplierID string) error {
	result := r.db.Delete(&Supplier{
		ID: supplierID,
	})
	return result.Error
}
