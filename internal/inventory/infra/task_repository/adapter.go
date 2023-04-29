package task_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
	"github.com/mrokoo/goERP/internal/model"
)

func toModel(t *task.Task) *model.Task {
	var items []model.TaskItem
	for _, item := range t.Items {
		items = append(items, *toModelItem(&item))
	}
	var records []model.TaskRecord
	for _, record := range t.Recrods {
		records = append(records, *toModelRecord(&record))
	}

	return &model.Task{
		ID:                    t.ID,
		WarehouseID:           t.WarehouseID,
		Kind:                  string(t.Kind),
		State:                 string(t.State),
		Items:                 items,
		Recrods:               records,
		IO:                    t.IO,
		PurchaseOrderID:       t.PurchaseOrderID,
		PurchaseReturnOrderID: t.PurchaseReturnOrderID,
		SaleOrderID:           t.SaleOrderID,
		SaleReturnOrderID:     t.SaleReturnOrderID,
		AllotID:               t.AllotID,
		CreatedAt:             t.CreatedAt,
	}
}

func toDomain(t *model.Task) *task.Task {
	var items []task.TaskItem
	for _, item := range t.Items {
		items = append(items, *toDomainItem(&item))
	}
	var records []record.Record
	for _, record := range t.Recrods {
		records = append(records, *toDomainRecord(&record))
	}
	return &task.Task{
		ID:                    t.ID,
		WarehouseID:           t.WarehouseID,
		Kind:                  task.Kind(t.Kind),
		State:                 state.State(t.State),
		Items:                 items,
		Recrods:               records,
		IO:                    t.IO,
		PurchaseOrderID:       t.PurchaseOrderID,
		PurchaseReturnOrderID: t.PurchaseReturnOrderID,
		SaleOrderID:           t.SaleOrderID,
		SaleReturnOrderID:     t.SaleReturnOrderID,
		AllotID:               t.AllotID,
		CreatedAt:             t.CreatedAt,
	}
}

func toModelItem(i *task.TaskItem) *model.TaskItem {
	return &model.TaskItem{
		ProductID: i.ProductID,
		Total:     i.Total,
		Quantity:  i.Quantity,
	}
}

func toDomainItem(i *model.TaskItem) *task.TaskItem {
	return &task.TaskItem{
		ProductID: i.ProductID,
		Total:     i.Total,
		Quantity:  i.Quantity,
	}
}

func toModelRecord(r *record.Record) *model.TaskRecord {
	var items []model.TaskRecordItem
	for _, item := range r.Items {
		items = append(items, *toModelRecordItem(&item))
	}
	return &model.TaskRecord{
		ID:          r.ID,
		WarehouseID: r.WarehouseID,
		UserID:      r.UserID,
		State:       string(r.State),
		Items:       items,
		CreatedAt:   r.CreatedAt,
	}
}

func toDomainRecord(r *model.TaskRecord) *record.Record {
	var items []record.RecordItem
	for _, item := range r.Items {
		items = append(items, *toDomainRecordItem(&item))
	}
	return &record.Record{
		ID:          r.ID,
		WarehouseID: r.WarehouseID,
		UserID:      r.UserID,
		State:       state.State(r.State),
		Items:       items,
		CreatedAt:   r.CreatedAt,
	}
}

func toModelRecordItem(i *record.RecordItem) *model.TaskRecordItem {
	return &model.TaskRecordItem{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}

func toDomainRecordItem(i *model.TaskRecordItem) *record.RecordItem {
	return &record.RecordItem{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}
