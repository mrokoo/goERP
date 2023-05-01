package repository

import (
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(unit *domain.Unit) *model.Unit {
	return &model.Unit{
		ID:   unit.ID,
		Name: unit.Name,
		Note: unit.Note,
	}
}

func toDomain(unit *model.Unit) *domain.Unit {
	return &domain.Unit{
		ID:   unit.ID,
		Name: unit.Name,
		Note: unit.Note,
	}
}
