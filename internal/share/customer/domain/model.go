package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Customer struct {
	ID      string      `json:"id" gorm:"primaryKey;<-:create"`
	Name    string      `json:"name" gorm:"not null"`
	Grade   GradeType   `json:"grade" gorm:"default:medium"`
	Contact string      `json:"contact"`
	Phone   string      `json:"phone"`
	Email   string      `json:"email"`
	Address string      `json:"address"`
	Note    string      `json:"note"`
	State   state.State `json:"state" gorm:"default:active"`
	Debt    float64     `json:"debt"`
}

type GradeType string

const (
	GRADE_HIGH   GradeType = "high"
	GRADE_MEDIUM GradeType = "medium"
	GRADE_LOW    GradeType = "low"
)
