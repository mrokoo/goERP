package domain

type Repository interface {
	GetAll() ([]*Customer, error)
	GetByID(customerID string) (*Customer, error)
	Save(customer *Customer) error
	Replace(customer *Customer) error
	Delete(customerID string) error
}
