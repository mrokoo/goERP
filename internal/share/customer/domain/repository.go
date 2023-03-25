package domain

type Repository interface {
	Get(customerID CustomerId) (*Customer, error)
	GetAll() ([]Customer, error)
	Update(customer Customer) error
	Save(customer Customer) error
	Delete(customerID CustomerId) error
}
