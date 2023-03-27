package domain

type Repository interface {
	Save(role *Role) error
	Update(role *Role) error
	Delete(name string) error
	FindByName(name string) (*Role, error)
	FindaAll() ([]*Role, error)
}
