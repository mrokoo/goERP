package repository

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"gorm.io/gorm"
)

// 要将偶然复杂性消除掉，通过添加模型转换，或者适配器进行个隔离。
var ErrNotFound = gorm.ErrRecordNotFound

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	db.AutoMigrate(&Product{}) //自动迁移
	db.AutoMigrate(&Stock{})
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]*domain.Product, error) {
	var productms []Product
	result := r.db.Preload("OpeningStock").Find(&productms)
	if err := result.Error; err != nil {
		return nil, err
	}
	var products []*domain.Product
	for _, pm := range productms {
		products = append(products, pm.toProduct())
	}

	return products, nil
}

func (r *ProductRepository) GetByID(productID string) (*domain.Product, error) {
	pm := Product{
		ID: productID,
	}
	result := r.db.First(&pm)
	if err := result.Error; err != nil {
		return nil, err
	}
	product := pm.toProduct()
	return product, nil
}

func (r *ProductRepository) Save(product *domain.Product) error {
	pm := toMySQLProduct(product)
	result := r.db.Create(pm)
	return result.Error
}

func (r *ProductRepository) Replace(product *domain.Product) error {
	pm := toMySQLProduct(product)
	result := r.db.Save(pm)
	return result.Error
}

func (r *ProductRepository) Delete(productID string) error {
	result := r.db.Delete(&Product{
		ID: productID,
	})
	return result.Error
}
