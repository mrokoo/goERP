package customer

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Customer struct {
	ID          CustomerId           `json:"id"`
	Name        valueobj.Name        `json:"name"`
	Grade       GradeType            `json:"grade"`
	Contact     valueobj.Contact     `json:"contact"`
	PhoneNumber valueobj.PhoneNumber `json:"phoneNumber"`
	Address     valueobj.Address     `json:"address"`
	Note        string               `json:"note"`
	State       valueobj.StateType   `json:"state"`
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
	customer.Name, err = valueobj.NewName(cmd.Name)
	if err != nil {
		return Customer{}, err
	}
	customer.Grade, err = NewGrade(cmd.Grade)
	if err != nil {
		return Customer{}, err
	}
	customer.Contact, err = valueobj.NewContact(cmd.Contact)
	if err != nil {
		return Customer{}, err
	}
	customer.PhoneNumber, err = valueobj.NewPhoneNumber(cmd.PhoneNumber)
	if err != nil {
		return Customer{}, err
	}
	customer.Address, err = valueobj.NewAddress(cmd.Address)
	if err != nil {
		return Customer{}, err
	}
	customer.Note = cmd.Note
	customer.State, err = valueobj.NewState(cmd.State)
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
