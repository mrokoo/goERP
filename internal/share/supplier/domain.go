package supplier

import "github.com/Rhymond/go-money"

type Supplier struct {
	ID      SupplierId
	Name    Name
	Contact ContactName
	Email   Email
	Address Address
	Account Account
	Bank    BankName
	Note    string
	State   StateType
	Debt    money.Money
}

type SupplierId string
type Name string
type ContactName string
type Email string
type Address string
type Account string
type BankName string

type StateType int

const (
	STATE_INVAILD = iota
	STATE_ACTIVE
	STATE_FREEZE
)
