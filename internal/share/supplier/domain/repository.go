package domain

type Repository interface {
	Get(supplierID SupplierId) (*Supplier, error)
	GetAll() ([]Supplier, error)
	Update(supplier Supplier) error
	Save(supplier Supplier) error
	Delete(supplierID SupplierId) error
}
