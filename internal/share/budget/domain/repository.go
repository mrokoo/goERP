package domain

import (
	"github.com/google/uuid"
)

type Repository interface {
	Get(budgetID uuid.UUID) (*Budget, error)
	GetAll() ([]Budget, error)
	Update(budget Budget) error
	Save(budget Budget) error
	Delete(budgetID uuid.UUID) error
}
