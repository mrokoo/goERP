package repository

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Category = domain.Category

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	db.AutoMigrate(&Category{}) //自动迁移
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAll() ([]*Category, error) {
	var categorys []Category
	result := r.db.Find(&categorys)
	if err := result.Error; err != nil {
		return nil, err
	}
	var categorysp []*Category
	for _, v := range categorys {
		categorysp = append(categorysp, &v)
	}
	return categorysp, nil
}

func (r *CategoryRepository) GetByID(categoryID uuid.UUID) (*Category, error) {
	category := Category{
		ID: categoryID,
	}
	result := r.db.First(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Save(category *domain.Category) error {
	result := r.db.Create(category)
	return result.Error
}

func (r *CategoryRepository) Replace(category *domain.Category) error {
	result := r.db.Save(category)
	return result.Error
}

func (r *CategoryRepository) Delete(categoryID uuid.UUID) error {
	result := r.db.Delete(&Category{
		ID: categoryID,
	})
	return result.Error
}
