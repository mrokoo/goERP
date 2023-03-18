package budget

import (
	"errors"

	"github.com/google/uuid"
)

type Budget struct {
	ID   uuid.UUID
	Type BudgetType
	Note string
}

type BudgetCMD struct {
	Type int
	Note string
}

func NewBudget(cmd BudgetCMD) (Budget, error) {
	var budget Budget
	var err error
	budget.ID = uuid.New()
	budget.Type, err = NewType(cmd.Type)
	if err != nil {
		return Budget{}, err
	}

	budget.Note = cmd.Note
	return budget, nil
}

const (
	BUDGET_INVALID = iota
	BUDGET_OUT
	BUDGET_IN
)

type BudgetType int

func NewType(t int) (BudgetType, error) {
	if t < 0 || t > 2 {
		return BUDGET_INVALID, errors.New("the type is wrong")
	}
	return BudgetType(t), nil
}
