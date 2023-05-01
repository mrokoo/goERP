package repository

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

// 要将偶然复杂性消除掉，通过添加模型转换，或者适配器进行个隔离。
var ErrNotFound = gorm.ErrRecordNotFound

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]*domain.Product, error) {
	var list []*model.Product
	result := r.db.Preload("OpeningStock").Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var products []*domain.Product
	for i := range list {
		products = append(products, toDomain(list[i]))
	}
	return products, nil
}

func (r *ProductRepository) GetByID(productID string) (*domain.Product, error) {
	product := model.Product{
		ID: productID,
	}
	result := r.db.Preload("OpeningStock").First(&product)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&product), nil
}

func (r *ProductRepository) Save(product *domain.Product) error {
	i := toModel(product)
	result := r.db.Create(i)
	return result.Error
}

func (r *ProductRepository) Replace(product *domain.Product) error {
	i := toModel(product)
	result := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&i)
	return result.Error
}

func (r *ProductRepository) Delete(productID string) error {
	result := r.db.Delete(&model.Product{
		ID: productID,
	})
	return result.Error
}
