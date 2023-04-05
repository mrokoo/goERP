package domain

const (
	CollectionSupplier = "suppliers"
)

type Repository interface {
	GetAll() ([]*Supplier, error)
	GetByID(supplierID string) (*Supplier, error)
	Save(supplier *Supplier) error
	Replace(supplier *Supplier) error
	Delete(supplierID string) error
}
