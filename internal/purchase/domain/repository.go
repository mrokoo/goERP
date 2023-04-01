package domain

type Repository interface {
	Get(orderID string) (*PurchaseOrder, error)
	GetAll() ([]*PurchaseOrder, error)
	Save(purchaseOrder PurchaseOrder) error
	Invalidated(orderID string) error
}
