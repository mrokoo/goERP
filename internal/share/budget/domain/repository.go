package domain

import (
	"github.com/google/uuid"
)

const (
	CollectionBudget = "budgets"
)

type Repository interface {
	GetAll() ([]*Budget, error)
	GetByID(budgetID uuid.UUID) (*Budget, error)
	Save(budget *Budget) error
	Replace(budget *Budget) error
	Delete(budgetID uuid.UUID) error
}
