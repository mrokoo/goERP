package app

import (
	"github.com/mrokoo/goERP/internal/goods/category/domain"
)

type CategoryService interface {
	GetCategory(categoryID string) (*domain.Category, error)
	GetCategoryList() ([]*domain.Category, error)
	AddCategory(category *domain.Category) error
	ReplaceCategory(category *domain.Category) error
	DeleteCategory(categoryID string) error
}

type CategoryServiceImpl struct {
	categoryRepository domain.Repository
}

func NewCategoryServiceImpl(categoryRepository domain.Repository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryServiceImpl) GetCategoryList() ([]*domain.Category, error) {
	var categories []*domain.Category
	categories, err := s.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryServiceImpl) GetCategory(categoryID string) (*domain.Category, error) {
	category, err := s.categoryRepository.GetByID(categoryID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryServiceImpl) AddCategory(category *domain.Category) error {
	if err := s.categoryRepository.Save(category); err != nil {
		return err
	}
	return nil
}

func (s *CategoryServiceImpl) ReplaceCategory(category *domain.Category) error {
	err := s.categoryRepository.Save(category)
	return err
}

func (s *CategoryServiceImpl) DeleteCategory(categoryID string) error {
	if err := s.categoryRepository.Delete(categoryID); err != nil {
		return err
	}
	return nil
}
