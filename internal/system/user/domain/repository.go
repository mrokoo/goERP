package domain

import (
	"github.com/google/uuid"
)

type Repository interface {
	GetAll() ([]*User, error)
	GetByID(userID uuid.UUID) (*User, error)
	Save(user *User) error
	Replace(user *User) error
	Delete(userID uuid.UUID) error
}
