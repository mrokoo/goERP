package repository

import (
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

func (r *ProductRepository) GetAll() ([]*Product, error) {
	var productms []ProductModel
	result := r.db.Preload("OpeningStock").Find(&productms)
	if err := result.Error; err != nil {
		return nil, err
	}
	var products []*Product
	for _, pm := range productms {
		products = append(products, pm.toProduct())
	}

	return products, nil
}

func (r *ProductRepository) GetByID(productID string) (*Product, error) {
	pm := ProductModel{
		ID: productID,
	}
	result := r.db.First(&pm)
	if err := result.Error; err != nil {
		return nil, err
	}
	product := pm.toProduct()
	return product, nil
}

func (r *ProductRepository) Save(product *Product) error {
	pm := toProductModel(product)
	result := r.db.Create(pm)
	return result.Error
}

func (r *ProductRepository) Replace(product *Product) error {
	pm := toProductModel(product)
	result := r.db.Save(pm)
	return result.Error
}

func (r *ProductRepository) Delete(productID string) error {
	result := r.db.Delete(&ProductModel{
		ID: productID,
	})
	return result.Error
}
