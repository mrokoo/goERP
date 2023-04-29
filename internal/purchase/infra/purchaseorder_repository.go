package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type PurchaseOrderRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{
		db: db,
	}
}

func (r *PurchaseOrderRepository) GetAll() ([]*domain.PurchaseOrder, error) {
	var list []*model.PurchaseOrder
	result := r.db.Preload("Items").Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	// 转换成领域层的模型
	var orders []*domain.PurchaseOrder
	for _, order := range list {
		orders = append(orders, toDomain(order))
	}
	return orders, nil
}

func (r *PurchaseOrderRepository) GetByID(purchaseOrderID string) (*domain.PurchaseOrder, error) {
	order := model.PurchaseOrder{
		ID: purchaseOrderID,
	}
	result := r.db.Preload("Items").First(&order)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&order), nil
}

func (r *PurchaseOrderRepository) Save(purchaseOrder *domain.PurchaseOrder) error {
	o := toModel(purchaseOrder)
	result := r.db.Create(o)
	return result.Error
}

func (r *PurchaseOrderRepository) Invalidated(purchaseOrderID string) error {
	result := r.db.Model(&model.PurchaseOrder{}).Where("id", purchaseOrderID).Update("is_validated", true)
	return result.Error
}
