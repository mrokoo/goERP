package supplier

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/Rhymond/go-money"
)

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

type SupplierCMD struct {
	ID      string
	Name    string
	Contact string
	Email   string
	Address string
	Account string
	Bank    string
	Note    string
	State   int
	Debt    float64
}

type SupplierId string

func (s *SupplierId) CheckSupplierID() error {
	regRuler := "^S[0-9]+$"
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(string(*s)) {
		return errors.New("the supplierID is wrong")
	}
	return nil
}

func NewSupplierID(id string) (SupplierId, error) {
	sid := SupplierId(id)
	if err := sid.CheckSupplierID(); err != nil {
		return "", err
	}
	return sid, nil
}

type Name string

func NewName(name string) (Name, error) {
	if l := len(name); l < 0 || l > 50 {
		return "", errors.New(" the name length does not meet the requirements")
	}
	return Name(name), nil
}

type ContactName string

func NewContact(contactName string) (ContactName, error) {
	if l := len(contactName); l < 0 || l > 50 {
		return "", errors.New(" the contact name length does not meet the requirements")
	}
	return ContactName(contactName), nil
}

type Email string

func NewEmail(email string) (Email, error) {
	regRuler := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(email) {
		return "", errors.New("not a email")
	}
	return Email(email), nil
}

type Address string

func NewAddress(addr string) (Address, error) {
	if len(addr) > 50 {
		return "", errors.New("the address length is too length")
	}
	return Address(addr), nil
}

type Account string

func NewAccount(account string) (Account, error) {
	if len(account) > 50 {
		return "", errors.New("the account length is too length")
	}
	return Account(account), nil
}

type BankName string

func NewBank(bank string) (BankName, error) {
	if len(bank) > 50 {
		return "", errors.New("the bank length is too length")
	}
	return BankName(bank), nil
}

type StateType int

const (
	STATE_INVAILD = iota
	STATE_ACTIVE
	STATE_FREEZE
)

func NewState(state int) (StateType, error) {
	if state < 1 || state > 2 {
		return STATE_INVAILD, errors.New("the state is invaild")
	}
	return StateType(state), nil
}

func NewSupplier(cmd SupplierCMD) (_ Supplier, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("new supplier: %w", err)
		}
	}()
	supplier := Supplier{}

	// to do id 验证
	supplier.ID, err = NewSupplierID(cmd.ID)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Name, err = NewName(cmd.Name)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Contact, err = NewContact(cmd.Contact)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Email, err = NewEmail(cmd.Email)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Address, err = NewAddress(cmd.Address)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Account, err = NewAccount(cmd.Account)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Bank, err = NewBank(cmd.Bank)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Note = cmd.Note
	supplier.State, err = NewState(cmd.State)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Debt = *money.NewFromFloat(cmd.Debt, money.CNY)
	return supplier, nil
}
