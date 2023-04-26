package domain

import (
	"github.com/google/uuid"
)

type Budget struct {
	ID   uuid.UUID `gorm:"primaryKey;<-:create"`
	Name string
	Type BudgetType `gorm:"default:in"`
	Note string
}

type BudgetType string

const (
	BUDGET_OUT = "out"
	BUDGET_IN  = "in"
)
