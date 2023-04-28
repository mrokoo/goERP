package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/customer/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

func toModel(c *domain.Customer) *model.Customer {
	return &model.Customer{
		ID:      c.ID,
		Name:    c.Name,
		Grade:   string(c.Grade),
		Contact: c.Contact,
		Phone:   c.Phone,
		Email:   c.Email,
		Address: c.Address,
		Note:    c.Note,
		State:   string(c.State),
		Debt:    c.Debt,
	}
}

func toDomain(c *model.Customer) *domain.Customer {
	return &domain.Customer{
		ID:      c.ID,
		Name:    c.Name,
		Grade:   domain.GradeType(c.Grade),
		Contact: c.Contact,
		Phone:   c.Phone,
		Email:   c.Email,
		Address: c.Address,
		Note:    c.Note,
		State:   state.State(c.State),
		Debt:    c.Debt,
	}
}
