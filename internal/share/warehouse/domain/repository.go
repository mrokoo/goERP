package domain

type Repository interface {
	Get(warehouseID WarehouseId) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	Update(warehouse Warehouse) error
	Save(warehouse Warehouse) error
	Delete(warehouseID WarehouseId) error
}
