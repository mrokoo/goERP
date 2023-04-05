package domain

import (
	"github.com/google/uuid"
)

type Budget struct {
	ID   uuid.UUID
	Type BudgetType
	Note string
}

type BudgetType string

const (
	BUDGET_OUT = "out"
	BUDGET_IN  = "in"
)
