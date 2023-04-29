package infra

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/sale/domain"
)

func toModel(order *domain.SaleOrder) *model.SaleOrder {
	var items []model.SaleOrderItem
	for _, item := range order.Items {
		items = append(items, *toModelItem(&item))
	}
	return &model.SaleOrder{
		ID:          order.ID,
		WarehouseID: order.WarehouseID,
		CustomerID:  order.CustomerID,
		UserID:      order.UserID,
		CreatedAt:   order.CreatedAt,
		Basic:       order.Basic,
		Items:       items,
		Kind:        string(order.Kind),
	}
}

func toDomain(order *model.SaleOrder) *domain.SaleOrder {
	var items []domain.Item
	for _, item := range order.Items {
		items = append(items, *toDomainItem(&item))
	}
	return &domain.SaleOrder{
		ID:          order.ID,
		WarehouseID: order.WarehouseID,
		CustomerID:  order.CustomerID,
		UserID:      order.UserID,
		CreatedAt:   order.CreatedAt,
		Basic:       order.Basic,
		Items:       items,
		Kind:        domain.Kind(order.Kind),
	}
}

func toModelItem(item *domain.Item) *model.SaleOrderItem {
	return &model.SaleOrderItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}

func toDomainItem(item *model.SaleOrderItem) *domain.Item {
	return &domain.Item{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}
