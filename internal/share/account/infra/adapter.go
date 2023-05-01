package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

func toDmain(a *model.Account) *domain.Account {
	return &domain.Account{
		ID:      a.ID,
		Name:    a.Name,
		Type:    domain.PayType(a.Type),
		Holder:  a.Holder,
		Number:  a.Number,
		State:   state.State(a.State),
		Note:    a.Note,
		Balance: a.Balance,
	}
}

func toModel(a *domain.Account) *model.Account {
	return &model.Account{
		ID:      a.ID,
		Name:    a.Name,
		Type:    string(a.Type),
		Holder:  a.Holder,
		Number:  a.Number,
		State:   string(a.State),
		Note:    a.Note,
		Balance: a.Balance,
	}
}
