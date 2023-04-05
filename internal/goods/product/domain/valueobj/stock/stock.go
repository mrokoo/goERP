package stock

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	warehouse "github.com/mrokoo/goERP/internal/shared/warehouse/domain"
)

type Stock struct {
	ProductID   string
	Product     domain.Product
	WarehouseID string
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount      int
}
