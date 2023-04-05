package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/shopspring/decimal"
)

type Supplier struct {
	ID      string          `json:"id" binding:"required"`
	Name    string          `json:"name" binding:"required"`
	Contact string          `json:"contact" binding:"-"`
	Email   string          `json:"email" binding:"-"`
	Address string          `json:"address" binding:"-"`
	Account string          `json:"account" binding:"-"`
	Bank    string          `json:"bank" binding:"-"`
	Note    string          `json:"note" binding:"-"`
	State   state.State     `json:"state" binding:"-"`
	Debt    decimal.Decimal `json:"debt" binding:"-"`
}
