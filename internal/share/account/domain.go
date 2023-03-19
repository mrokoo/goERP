package account

import "github.com/mrokoo/goERP/internal/share/valueobj"

type Account struct {
	ID      AccountId          `json:"id" binding:"required"`
	Name    valueobj.Name      `json:"name" binding:"required"`
	Type    Type               `json:"type"`
	Holder  string             `json:"holder"`
	Number  string             `json:"number"`
	Note    valueobj.Note      `json:"note"`
	State   valueobj.StateType `json:"state"`
	Balance valueobj.Balance   `json:"balance"`
}

type AccountId = string

type Type int

const (
	TYPE_INVALID = iota
	TYPE_CASH
	TYPE_WEIPAY
	TYPE_ALiPAY
	TYPE_OTHER
)
