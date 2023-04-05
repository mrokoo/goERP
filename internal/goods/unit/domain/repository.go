package domain

import (
	"github.com/google/uuid"
)

const (
	CollectionUnit = "units"
)

type Repository interface {
	GetAll() ([]*Unit, error)
	GetByID(unitID uuid.UUID) (*Unit, error)
	Save(unit *Unit) error
	Replace(unit *Unit) error
	Delete(unitID uuid.UUID) error
}
