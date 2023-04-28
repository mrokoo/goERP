package repository

import (
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(category *domain.Category) *model.Category {
	return &model.Category{
		ID:   category.ID,
		Name: category.Name,
		Note: category.Note,
	}
}

func toDomain(category *model.Category) *domain.Category {
	return &domain.Category{
		ID:   category.ID,
		Name: category.Name,
		Note: category.Note,
	}
}
