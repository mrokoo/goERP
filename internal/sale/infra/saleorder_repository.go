package infra

import (
	"github.com/mrokoo/goERP/internal/sale/domain"
	"gorm.io/gorm"
)

type SaleOrderRepository struct {
	db *gorm.DB
}

func NewSaleOrderRepository(db *gorm.DB) *SaleOrderRepository {
	db.AutoMigrate(&MySQLSaleOrder{})
	db.AutoMigrate(&MySQLItem{})
	return &SaleOrderRepository{
		db: db,
	}
}

func (r *SaleOrderRepository) Save(saleOrder *domain.SaleOrder) error {
	order_ := toMySQLSaleOrder(saleOrder)
	return r.db.Create(order_).Error
}

func (r *SaleOrderRepository) GetAll() ([]*domain.SaleOrder, error) {
	var orders []*MySQLSaleOrder
	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var saleOrders []*domain.SaleOrder
	for _, order := range orders {
		saleOrders = append(saleOrders, order.toSaleOrder())
	}
	return saleOrders, nil
}

func (r *SaleOrderRepository) GetByID(id string) (*domain.SaleOrder, error) {
	var order MySQLSaleOrder
	order.ID = id
	err := r.db.Preload("Items").First(&order).Error
	if err != nil {
		return nil, err
	}
	return order.toSaleOrder(), nil
}
