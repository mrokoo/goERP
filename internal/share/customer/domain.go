package customer

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Customer struct {
	ID      CustomerId         `json:"id"`
	Name    valueobj.Name      `json:"name"`
	Grade   GradeType          `json:"grade"`
	Contact valueobj.Contact   `json:"contact"`
	Phone   valueobj.Phone     `json:"phone"`
	Email   valueobj.Email     `json:"email"`
	Address valueobj.Address   `json:"address"`
	Note    valueobj.Note      `json:"note"`
	State   valueobj.StateType `json:"state"`
	Debt    valueobj.Money     `json:"debt"`
}

type CustomerId = string

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
