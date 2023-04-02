package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Customer struct {
	ID      string         `json:"id" binding:"required"`
	Name    string         `json:"name" binding:"required"`
	Grade   GradeType      `json:"grade" binding:"-"`
	Contact string         `json:"contact" binding:"-"`
	Phone   string         `json:"phone" binding:"-"`
	Email   string         `json:"email" binding:"-"`
	Address string         `json:"address" binding:"-"`
	Note    string         `json:"note" binding:"-"`
	State   state.State    `json:"state" binding:"-"`
	Debt    valueobj.Money `json:"debt" binding:"-"`
}

type GradeType string

const (
	GRADE_HIGH   GradeType = "high"
	GRADE_MEDIUM GradeType = "medium"
	GRADE_LOW    GradeType = "low"
)
