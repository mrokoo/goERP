package domain

import "github.com/google/uuid"

type UnitRepository interface {
	Create(unit *Unit) error
	Save(unit *Unit) error
	Get(unitId *uuid.UUID) (*Unit, error)
	GetAll() ([]Unit, error)
	Delete(unitId *uuid.UUID) error
}
