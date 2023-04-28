package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

func toModel(s *domain.Supplier) *model.Supplier {
	return &model.Supplier{
		ID:      s.ID,
		Name:    s.Name,
		Contact: s.Contact,
		Email:   s.Email,
		Address: s.Address,
		Account: s.Account,
		Bank:    s.Bank,
		Note:    s.Note,
		State:   s.State.String(),
		Debt:    s.Debt,
	}
}

func toDomain(s *model.Supplier) *domain.Supplier {
	return &domain.Supplier{
		ID:      s.ID,
		Name:    s.Name,
		Contact: s.Contact,
		Email:   s.Email,
		Address: s.Address,
		Account: s.Account,
		Bank:    s.Bank,
		Note:    s.Note,
		State:   state.State(s.State),
		Debt:    s.Debt,
	}
}
