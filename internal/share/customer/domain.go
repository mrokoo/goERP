package customer

import (
	"errors"

	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Customer struct {
	ID      CustomerId         `json:"id" binding:"required"`
	Name    valueobj.Name      `json:"name" binding:"required"`
	Grade   GradeType          `json:"grade" binding:"-"`
	Contact valueobj.Contact   `json:"contact" binding:"-"`
	Phone   valueobj.Phone     `json:"phone" binding:"-"`
	Email   valueobj.Email     `json:"email" binding:"-"`
	Address valueobj.Address   `json:"address" binding:"-"`
	Note    valueobj.Note      `json:"note" binding:"-"`
	State   valueobj.StateType `json:"state" binding:"-"`
	Debt    valueobj.Money     `json:"debt" binding:"-"`
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
