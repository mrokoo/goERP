package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
)

func toModel(b *domain.Budget) *model.Budget {
	return &model.Budget{
		ID:   b.ID,
		Name: b.Name,
		Type: string(b.Type),
		Note: b.Note,
	}
}

func toDomain(b *model.Budget) *domain.Budget {
	return &domain.Budget{
		ID:   b.ID,
		Name: b.Name,
		Type: domain.BudgetType(b.Type),
		Note: b.Note,
	}
}
