package returnorder

import (
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"gorm.io/gorm"
)

type PurchaseReturnRepository struct {
	db *gorm.DB
}

func NewPurchaseReturnRepository(db *gorm.DB) *PurchaseReturnRepository {
	db.AutoMigrate(&PurchaseReturnOrder{})
	db.AutoMigrate(&PurchaseReturnOrderItem{})
	return &PurchaseReturnRepository{
		db: db,
	}
}
func (r *PurchaseReturnRepository) GetAll() ([]*domain.PurchaseReturnOrder, error) {
	var pom []PurchaseReturnOrder
	result := r.db.Preload("Items").Find(&pom)
	if err := result.Error; err != nil {
		return nil, err
	}
	// 转换成领域层的模型
	var po []*domain.PurchaseReturnOrder
	for _, po2 := range pom {
		po = append(po, po2.toPurchaseReturnOrder())
	}
	return po, nil
}

func (r *PurchaseReturnRepository) GetByID(purchaseOrderID string) (*domain.PurchaseReturnOrder, error) {
	po := PurchaseReturnOrder{
		ID: purchaseOrderID,
	}

	result := r.db.First(&po)
	if err := result.Error; err != nil {
		return nil, err
	}
	purchaseOrder := po.toPurchaseReturnOrder()
	return purchaseOrder, nil
}

func (r *PurchaseReturnRepository) Save(purchaseOrder *domain.PurchaseReturnOrder) error {
	po := toMySQLPurchaseReturnOrder(purchaseOrder)
	result := r.db.Create(&po)
	return result.Error
}

func (r *PurchaseReturnRepository) InValidate(purchaseOrderID string) error {
	result := r.db.Model(&PurchaseReturnOrder{}).Where("id", purchaseOrderID).Update("is_validated", true)
	return result.Error
}
