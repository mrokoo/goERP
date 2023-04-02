package domain

import (
	"github.com/google/uuid"
)

const (
	CollectionCategory = "categorys"
)

type Repository interface {
	GetAll() ([]*Category, error)
	GetByID(categoryID uuid.UUID) (*Category, error)
	Save(category *Category) error
	Replace(category *Category) error
	Delete(categoryID uuid.UUID) error
}
