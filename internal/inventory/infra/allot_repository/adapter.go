package allot_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(allot *allot.Allot) *model.Allot {
	var items []model.AllotItem
	for i := range allot.Items {
		items = append(items, *toModelItem(&allot.Items[i]))
	}
	return &model.Allot{
		ID:             allot.ID,
		InWarehouseID:  allot.InWarehouseID,
		OutWarehouseID: allot.OutWarehouseID,
		UserID:         allot.UserID,
		CreatedAt:      allot.CreatedAt,
		Items:          items,
	}
}

func toDomain(a *model.Allot) *allot.Allot {
	var items []allot.Item
	for i := range a.Items {
		items = append(items, *toDomainItem(&a.Items[i]))
	}
	return &allot.Allot{
		ID:             a.ID,
		InWarehouseID:  a.InWarehouseID,
		OutWarehouseID: a.OutWarehouseID,
		UserID:         a.UserID,
		CreatedAt:      a.CreatedAt,
		Items:          items,
	}
}

func toModelItem(item *allot.Item) *model.AllotItem {
	return &model.AllotItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
}

func toDomainItem(item *model.AllotItem) *allot.Item {
	return &allot.Item{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
}
