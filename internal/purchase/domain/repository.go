package domain

type PurchaseOrderRepository interface {
	GetByID(orderID string) (*PurchaseOrder, error)
	GetAll() ([]*PurchaseOrder, error)
	Save(purchaseOrder PurchaseOrder) error
	Invalidated(orderID string) error
}

type PurchaseReturnOrderRepository interface {
	GetByID(orderID string) (*PurchaseReturnOrder, error)
	GetAll() ([]*PurchaseReturnOrder, error)
	Save(purchaseReturnOrder PurchaseReturnOrder) error
	Invalidated(orderID string) error
}
