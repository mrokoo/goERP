package supplier

import (
	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Supplier struct {
	ID      SupplierId         `json:"id" binding:"required"`
	Name    valueobj.Name      `json:"name" binding:"required"`
	Contact valueobj.Contact   `json:"contact" binding:"-"`
	Email   valueobj.Email     `json:"email" binding:"-"`
	Address valueobj.Address   `json:"address" binding:"-"`
	Account BankAccount        `json:"account" binding:"-"`
	Bank    Bank               `json:"bank" binding:"-"`
	Note    valueobj.Note      `json:"note" binding:"-"`
	State   valueobj.StateType `json:"state" binding:"-"`
	Debt    valueobj.Money     `json:"debt" binding:"-"`
}

type SupplierId = string

type BankAccount = string

type Bank = string
