package domain

import "github.com/google/uuid"

type CategoryRepository interface {
	Create(category *Category) error
	Save(category *Category) error
	Get(categoryId *uuid.UUID) (*Category, error)
	GetAll() ([]Category, error)
	Delete(categoryId *uuid.UUID) error
}
