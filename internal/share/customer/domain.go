package customer

import (
	"errors"
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

// to do 验证结构体有效性。用于请求参数

type CustomerId string

type Name string

func NewName(name string) Name {
	return Name(name)
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
		return GRADE_INVAILD, errors.New("this grade is invaild")
	}
	return GradeType(grade), nil
}

type ContactName string

type PhoneNumber string

func NewPhoneNumber(number string) PhoneNumber {
	return PhoneNumber(number)
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
		return GRADE_INVAILD, errors.New("this state is invaild")
	}
	return StateType(state), nil
}
