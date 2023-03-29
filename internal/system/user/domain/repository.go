package domain

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) error
	Save(user *User) error
	Get(userId *uuid.UUID) (*User, error)
	GetAll() ([]User, error)
	Delete(userId *uuid.UUID) error
}
