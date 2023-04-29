package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
)

type Task struct {
	ID                    string
	WarehouseID           string
	Kind                  Kind
	State                 state.State
	Items                 []TaskItem
	Recrods               []record.Record
	IO                    bool
	PurchaseOrderID       *string
	PurchaseReturnOrderID *string
	SaleOrderID           *string
	SaleReturnOrderID     *string
	AllotID               *string
	CreatedAt             time.Time
}

func NewTask(warehouseID string, kind Kind, basic string, items []TaskItem) Task {
	var io bool
	var PurchaseOrderID, PurchaseReturnOrderID, SaleOrderID, SaleReturnOrderID, AllocationOrderID *string
	switch kind {
	case IN_PURCHASE:
		io = true
		PurchaseOrderID = &basic
	case IN_SALE:
		io = true
		SaleOrderID = &basic
	case IN_ALLOCATION:
		io = true
		AllocationOrderID = &basic
	case OUT_PURCHASE:
		io = false
		PurchaseReturnOrderID = &basic
	case OUT_SALE:
		io = false
		SaleReturnOrderID = &basic
	case OUT_ALLOCATION:
		io = false
		AllocationOrderID = &basic
	}

	return Task{
		ID:                    uuid.New().String(),
		WarehouseID:           warehouseID,
		Kind:                  kind,
		State:                 state.NORMAL,
		Items:                 items,
		IO:                    io,
		PurchaseOrderID:       PurchaseOrderID,
		PurchaseReturnOrderID: PurchaseReturnOrderID,
		SaleOrderID:           SaleOrderID,
		SaleReturnOrderID:     SaleReturnOrderID,
		AllotID:               AllocationOrderID,
		CreatedAt:             time.Now(),
	}
}

type TaskItem struct {
	ProductID string
	Total     int
	Quantity  int
}

func NewTaskItem(total int) TaskItem {
	return TaskItem{
		Total:    total,
		Quantity: 0,
	}
}

type Kind string

const (
	IN_PURCHASE    Kind = "in_purchase"
	OUT_PURCHASE   Kind = "out_purchase"
	IN_SALE        Kind = "in_sale"
	OUT_SALE       Kind = "out_sale"
	IN_ALLOCATION  Kind = "in_allocation"
	OUT_ALLOCATION Kind = "out_allocation"
)

func (t *Task) GetID() string {
	return t.ID
}

func (t *Task) InvalidateTask() error {
	if len(t.Recrods) != 0 {
		return errors.New("cannot be invalidated")
	}
	t.State = state.INVALID
	return nil
}

func (t *Task) InvalidateRecord(id string) error {
	for i, r := range t.Recrods {
		if r.ID == id {
			t.Recrods[i].State = state.INVALID
			t.UpdateTaskItems()
			return nil
		}
	}
	return errors.New("record not found")
}

func (t *Task) AddRecord(r record.Record) error {
	if err := t.CheckTaskRecordItems(r.Items); err != nil {
		return err
	}
	t.Recrods = append(t.Recrods, r)
	t.UpdateTaskItems()
	return nil
}

func (t *Task) UpdateTaskItems() {
	for i := range t.Items {
		q := 0
		for _, r := range t.Recrods {
			for _, ri := range r.Items {
				if ri.ProductID == t.Items[i].ProductID {
					q += ri.Quantity
				}
			}
		}
		t.Items[i].Quantity = q
	}
}

func (t *Task) CheckTaskRecordItems(items []record.RecordItem) error {
	set := make(map[string]struct{}, len(t.Items))
	for _, v := range t.Items {
		set[v.ProductID] = struct{}{}
	}
	for _, ri := range items {
		if _, ok := set[ri.ProductID]; !ok {
			return errors.New("item not found")
		}
	}
	return nil
}
