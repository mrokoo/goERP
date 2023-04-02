package domain

const (
	CollectionWarehouse = "warehouses"
)

type Repository interface {
	GetAll() ([]*Warehouse, error)
	GetByID(warehouseID string) (*Warehouse, error)
	Save(warehouse *Warehouse) error
	Replace(warehouse *Warehouse) error
	Delete(warehouseID string) error
}
