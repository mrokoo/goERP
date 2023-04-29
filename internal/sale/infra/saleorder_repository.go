package infra

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/sale/domain"
	"gorm.io/gorm"
)

type SaleOrderRepository struct {
	db *gorm.DB
}

func NewSaleOrderRepository(db *gorm.DB) *SaleOrderRepository {
	return &SaleOrderRepository{
		db: db,
	}
}

func (r *SaleOrderRepository) Save(saleOrder *domain.SaleOrder) error {
	order := toModel(saleOrder)
	return r.db.Create(order).Error
}

func (r *SaleOrderRepository) GetAll() ([]*domain.SaleOrder, error) {
	var orders []*model.SaleOrder
	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var saleOrders []*domain.SaleOrder
	for _, order := range orders {
		saleOrders = append(saleOrders, toDomain(order))
	}
	return saleOrders, nil
}

func (r *SaleOrderRepository) GetByID(id string) (*domain.SaleOrder, error) {
	var order *model.SaleOrder
	order.ID = id
	err := r.db.Preload("Items").First(&order).Error
	if err != nil {
		return nil, err
	}
	return toDomain(order), nil
}
