package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Supplier struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Contact string      `json:"contact"`
	Email   string      `json:"email"`
	Address string      `json:"address"`
	Account string      `json:"account"`
	Bank    string      `json:"bank"`
	Note    string      `json:"note"`
	State   state.State `json:"state"`
	Debt    float64     `json:"debt"`
}
