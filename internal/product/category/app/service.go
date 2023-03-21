package app

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/product/category/domain"
)

type CategoryServiceImpl struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository domain.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryServiceImpl) CreateCategory(name string, note string) (*domain.Category, error) {
	category := &domain.Category{
		ID:   uuid.New(),
		Name: name,
		Note: note,
	}

	if err := s.categoryRepository.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryServiceImpl) ChangeCategory(categoryId *uuid.UUID, name string, note string) (*domain.Category, error) {
	c, err := s.categoryRepository.Get(categoryId)
	if err != nil {
		return nil, err
	}
	c.ChangeName(name)
	c.ChangeNote(note)
	if err := s.categoryRepository.Save(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryServiceImpl) GetAllCategories() ([]domain.Category, error) {
	categories, err := s.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryServiceImpl) DeleteCategory(categoryId *uuid.UUID) error {
	if err := s.categoryRepository.Delete(categoryId); err != nil {
		return err
	}
	return nil
}
