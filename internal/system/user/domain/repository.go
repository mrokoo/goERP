package domain

import (
	"github.com/google/uuid"
)

type Repository interface {
	Get(userID uuid.UUID) (*User, error)
	GetAll() ([]User, error)
	Update(user User) error
	Save(user User) error
	Delete(userID uuid.UUID) error
}
