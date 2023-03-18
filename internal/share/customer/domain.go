package customer

import (
	"errors"
	"fmt"
	"regexp"
)

type Customer struct {
	ID          CustomerId  `json:"id"`
	Name        Name        `json:"name"`
	Grade       GradeType   `json:"grade"`
	Contact     ContactName `json:"contact"`
	PhoneNumber PhoneNumber `json:"phoneNumber"`
	Address     Address     `json:"address"`
	Note        string      `json:"note"`
	State       StateType   `json:"state"`
}

type CustomerCMD struct {
	ID          string
	Name        string
	Grade       int
	Contact     string
	PhoneNumber string
	Address     string
	Note        string
	State       int
}

func NewCustomer(cmd CustomerCMD) (_ Customer, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("new customer: %w", err)
		}
	}()
	customer := Customer{}

	// to do id 验证
	customer.ID, err = NewCustomerID(cmd.ID)
	if err != nil {
		return Customer{}, err
	}
	customer.Name, err = NewName(cmd.Name)
	if err != nil {
		return Customer{}, err
	}
	customer.Grade, err = NewGrade(cmd.Grade)
	if err != nil {
		return Customer{}, err
	}
	customer.Contact, err = NewContact(cmd.Contact)
	if err != nil {
		return Customer{}, err
	}
	customer.PhoneNumber, err = NewPhoneNumber(cmd.PhoneNumber)
	if err != nil {
		return Customer{}, err
	}
	customer.Address, err = NewAddress(cmd.Address)
	if err != nil {
		return Customer{}, err
	}
	customer.Note = cmd.Note
	customer.State, err = NewState(cmd.State)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

type CustomerId string

func (c *CustomerId) CheckCustomerID() error {
	regRuler := "^C[0-9]+$"
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(string(*c)) {
		return errors.New("the customerID is wrong")
	}
	return nil
}

func NewCustomerID(id string) (CustomerId, error) {
	cid := CustomerId(id)
	if err := cid.CheckCustomerID(); err != nil {
		return "", err
	}
	return cid, nil
}

type Name string

func NewName(name string) (Name, error) {
	if l := len(name); l < 0 || l > 50 {
		return "", errors.New(" the name length does not meet the requirements")
	}
	return Name(name), nil
}

type GradeType int

const (
	GRADE_INVAILD = iota
	GRADE_HIGH
	GRADE_MEDIUM
	GRADE_LOW
)

func NewGrade(grade int) (GradeType, error) {
	if grade < 1 || grade > 3 {
		return GRADE_INVAILD, errors.New("the grade is invaild")
	}
	return GradeType(grade), nil
}

type ContactName string

func NewContact(contactName string) (ContactName, error) {
	if l := len(contactName); l < 0 || l > 50 {
		return "", errors.New(" the contact name length does not meet the requirements")
	}
	return ContactName(contactName), nil
}

type PhoneNumber string

func NewPhoneNumber(number string) (PhoneNumber, error) {
	regRuler := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(number) {
		return "", errors.New("not a phone number")
	}
	return PhoneNumber(number), nil
}

type Address string

func NewAddress(addr string) (Address, error) {
	if len(addr) > 50 {
		return "", errors.New("the address length is too length")
	}
	return Address(addr), nil
}

type StateType int

const (
	STATE_INVAILD StateType = iota
	STATE_ACTIVE
	STATE_FREEZE
)

func NewState(state int) (StateType, error) {
	if state < 1 || state > 2 {
		return GRADE_INVAILD, errors.New("the state is invaild")
	}
	return StateType(state), nil
}
