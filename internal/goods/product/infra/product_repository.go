package repository

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Product = domain.Product

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {

	db.AutoMigrate(&Product{}) //自动迁移
	db.AutoMigrate(&stock.Stock{})
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]*Product, error) {
	var products []Product
	result := r.db.Find(&products)
	if err := result.Error; err != nil {
		return nil, err
	}
	var productsp []*Product
	for _, v := range products {
		productsp = append(productsp, &v)
	}
	return productsp, nil
}

func (r *ProductRepository) GetByID(productID string) (*Product, error) {
	product := Product{
		ID: productID,
	}
	result := r.db.First(&product)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Save(product *domain.Product) error {
	result := r.db.Create(product)
	return result.Error
}

func (r *ProductRepository) Replace(product *domain.Product) error {
	result := r.db.Save(product)
	return result.Error
}

func (r *ProductRepository) Delete(productID string) error {
	result := r.db.Delete(&Product{
		ID: productID,
	})
	return result.Error
}
