package supplier

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Supplier struct {
	ID      SupplierId
	Name    valueobj.Name
	Contact valueobj.Contact
	Email   valueobj.Email
	Address valueobj.Address
	Account Account
	Bank    BankName
	Note    string
	State   valueobj.StateType
	Debt    float64
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
	supplier.Name, err = valueobj.NewName(cmd.Name)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Contact, err = valueobj.NewContact(cmd.Contact)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Email, err = valueobj.NewEmail(cmd.Email)
	if err != nil {
		return Supplier{}, err
	}
	supplier.Address, err = valueobj.NewAddress(cmd.Address)
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
	supplier.State, err = valueobj.NewState(cmd.State)
	if err != nil {
		return Supplier{}, err
	}

	supplier.Debt = cmd.Debt
	return supplier, nil
}
