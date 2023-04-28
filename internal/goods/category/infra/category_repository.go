package repository

import (
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrDuplicateEntry = gorm.ErrInvalidDB

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetAll() ([]*domain.Category, error) {
	var list []*model.Category
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var categories []*domain.Category
	for _, category := range list {
		categories = append(categories, toDomain(category))
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(ID string) (*domain.Category, error) {
	category := model.Category{
		ID: ID,
	}
	result := r.db.First(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&category), nil
}

func (r *CategoryRepository) Save(category *domain.Category) error {
	i := toModel(category)
	result := r.db.Save(i)
	return result.Error
}

func (r *CategoryRepository) Delete(categoryID string) error {
	result := r.db.Delete(&model.Category{
		ID: categoryID,
	})
	return result.Error
}
