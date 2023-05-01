package domain

type Repository interface {
	GetAll() ([]*Category, error)
	GetByID(categoryID string) (*Category, error)
	Save(category *Category) error
	Delete(categoryID string) error
}
