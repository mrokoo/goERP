package domain

type SaleOrderRepository interface {
	GetAll() ([]*SaleOrder, error)
	GetByID(id string) (*SaleOrder, error)
	Save(saleOrder *SaleOrder) error
}
