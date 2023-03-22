package domain

type ProductRepository interface {
	Create(product *Product) error
	Save(product *Product) error
	Get(productId string) (*Product, error)
	GetAll() ([]Product, error)
	Delete(productId string) error
}
