package take_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(take *take.Take) *model.Take {
	var items []model.TakeItem
	for _, item := range take.Items {
		items = append(items, *toModelItem(&item))
	}
	return &model.Take{
		ID:          take.ID,
		WarehouseID: take.WarehouseID,
		UserID:      take.UserID,
		CreateAt:    take.CreateAt,
		Items:       items,
	}
}

func toDomain(t *model.Take) *take.Take {
	var items []take.Item
	for _, item := range t.Items {
		items = append(items, *toDomianItem(&item))
	}
	return &take.Take{
		ID:          t.ID,
		WarehouseID: t.WarehouseID,
		UserID:      t.UserID,
		CreateAt:    t.CreateAt,
		Items:       items,
	}
}

func toModelItem(item *take.Item) *model.TakeItem {
	return &model.TakeItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
}

func toDomianItem(i *model.TakeItem) *take.Item {
	return &take.Item{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}
