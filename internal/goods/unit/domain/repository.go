package domain

type Repository interface {
	GetAll() ([]*Unit, error)
	GetByID(unitID string) (*Unit, error)
	Save(unit *Unit) error
	Replace(unit *Unit) error
	Delete(unitID string) error
}
