package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

func toModel(w *domain.Warehouse) *model.Warehouse {
	return &model.Warehouse{
		ID:      w.ID,
		Name:    w.Name,
		Admin:   w.Admin,
		Phone:   w.Phone,
		Address: w.Address,
		Note:    w.Note,
		State:   string(w.State),
	}
}

func toDomain(w *model.Warehouse) *domain.Warehouse {
	return &domain.Warehouse{
		ID:      w.ID,
		Name:    w.Name,
		Admin:   w.Admin,
		Phone:   w.Phone,
		Address: w.Address,
		Note:    w.Note,
		State:   state.State(w.State),
	}
}
