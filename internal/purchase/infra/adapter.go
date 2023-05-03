package purchaseorder_repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/purchase/domain"
)

func toModel(order *domain.PurchaseOrder) *model.PurchaseOrder {
	var items []model.PurchaseOrderItem
	for _, item := range order.Items {
		items = append(items, *toModelItem(&item))
	}
	return &model.PurchaseOrder{
		ID:           order.ID,
		SupplierID:   order.SupplierID,
		WarehouseID:  order.WarehouseID,
		UserID:       order.UserID,
		IsValidated:  order.IsValidated,
		Items:        items,
		AccountID:    order.AccountID,
		OtherCost:    order.OtherCost,
		TotalCost:    order.TotalCost,
		ActalPayment: order.ActalPayment,
		Debt:         order.Debt,
		CreatedAt:    order.CreatedAt,
		Basic:        order.Basic,
		Kind:         string(order.Kind),
	}
}

func toDomain(order *model.PurchaseOrder) *domain.PurchaseOrder {
	var items []domain.Item
	for _, item := range order.Items {
		items = append(items, *toDomainItem(&item))
	}
	return &domain.PurchaseOrder{
		ID:           order.ID,
		SupplierID:   order.SupplierID,
		WarehouseID:  order.WarehouseID,
		UserID:       order.UserID,
		IsValidated:  order.IsValidated,
		Items:        items,
		AccountID:    order.AccountID,
		OtherCost:    order.OtherCost,
		TotalCost:    order.TotalCost,
		ActalPayment: order.ActalPayment,
		Debt:         order.Debt,
		CreatedAt:    order.CreatedAt,
		Basic:        order.Basic,
		Kind:         domain.Kind(order.Kind),
	}
}

func toModelItem(item *domain.Item) *model.PurchaseOrderItem {
	return &model.PurchaseOrderItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}

func toDomainItem(item *model.PurchaseOrderItem) *domain.Item {
	return &domain.Item{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}
