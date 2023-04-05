package domain

const (
	CollectionProduct = "products"
)

type Repository interface {
	GetAll() ([]*Product, error)
	GetByID(productID string) (*Product, error)
	Save(product *Product) error
	Replace(product *Product) error
	Delete(productID string) error
}
