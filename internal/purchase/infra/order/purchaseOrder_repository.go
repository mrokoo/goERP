package order

import (
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type PurchaseOrderRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	db.AutoMigrate(&PurchaseOrder{}) //自动迁移
	db.AutoMigrate(&PurchaseOrderItem{})
	return &PurchaseOrderRepository{
		db: db,
	}
}

func (r *PurchaseOrderRepository) GetAll() ([]*domain.PurchaseOrder, error) {
	var pom []PurchaseOrder
	result := r.db.Preload("Items").Find(&pom)
	if err := result.Error; err != nil {
		return nil, err
	}
	// 转换成领域层的模型
	var po []*domain.PurchaseOrder
	for _, po2 := range pom {
		po = append(po, po2.toPurchaseOrder())
	}
	return po, nil
}

func (r *PurchaseOrderRepository) GetByID(purchaseOrderID string) (*domain.PurchaseOrder, error) {
	po := PurchaseOrder{
		ID: purchaseOrderID,
	}

	result := r.db.First(&po)
	if err := result.Error; err != nil {
		return nil, err
	}
	purchaseOrder := po.toPurchaseOrder()
	return purchaseOrder, nil
}

func (r *PurchaseOrderRepository) Save(purchaseOrder *domain.PurchaseOrder) error {
	po := toMySQLPurchaseOrder(purchaseOrder)
	result := r.db.Create(&po)
	return result.Error
}

func (r *PurchaseOrderRepository) Invalidated(purchaseOrderID string) error {
	result := r.db.Model(&PurchaseOrder{}).Where("id", purchaseOrderID).Update("is_validated", true)
	return result.Error
}
