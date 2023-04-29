package domain

type PurchaseOrderRepository interface {
	GetByID(orderID string) (*PurchaseOrder, error)
	GetAll() ([]*PurchaseOrder, error)
	Save(purchaseOrder *PurchaseOrder) error
	Invalidated(orderID string) error
}
