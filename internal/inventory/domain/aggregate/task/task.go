package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
)

type Task struct {
	ID                    string          `json:"id"`
	WarehouseID           string          `json:"warehouse_id"`
	Kind                  Kind            `json:"kind"`
	State                 state.State     `json:"state"`
	Items                 []TaskItem      `json:"items"`
	Recrods               []record.Record `json:"records"`
	IO                    bool            `json:"io"`
	PurchaseOrderID       *string         `json:"purchase_order_id"`
	PurchaseReturnOrderID *string         `json:"purchase_return_order_id"`
	SaleOrderID           *string         `json:"sale_order_id"`
	SaleReturnOrderID     *string         `json:"sale_return_order_id"`
	AllotID               *string         `json:"allot_id"`
	CreatedAt             time.Time       `json:"created_at"`
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
	ProductID string `json:"product_id"`
	Total     int    `json:"total"`
	Quantity  int    `json:"quantity"`
}

func NewTaskItem(productID string, total int) TaskItem {
	return TaskItem{
		ProductID: productID,
		Total:     total,
		Quantity:  0,
	}
}

type Kind string

const (
	IN_PURCHASE    Kind = "in_purchase"    // 采购
	OUT_PURCHASE   Kind = "out_purchase"   // 采购退货
	IN_SALE        Kind = "in_sale"        // 销售退货
	OUT_SALE       Kind = "out_sale"       // 销售
	IN_ALLOCATION  Kind = "in_allocation"  // 调拨入库
	OUT_ALLOCATION Kind = "out_allocation" // 调拨出库
)

func (t *Task) InvalidateTask() error {
	if len(t.Recrods) != 0 {
		return errors.New("cannot be invalidated")
	}
	t.State = state.INVALID
	return nil
}

// InvalidateRecord:使指定Task记录(record)无效，内部会自动更新TaskItem状态
func (t *Task) InvalidateRecord(recordID string) error {
	for index, record := range t.Recrods {
		if record.ID == recordID {
			t.Recrods[index].State = state.INVALID
			t.UpdateTaskItems()
			return nil
		}
	}
	return errors.New("record not found")
}

// AddRecord:添加Task记录(record)，内部会自动更新TaskItem状态
func (t *Task) AddRecord(r record.Record) error {
	if err := t.CheckTaskRecordItems(r.Items); err != nil {
		return err
	}
	t.Recrods = append(t.Recrods, r)
	t.UpdateTaskItems()
	return nil
}

// UpdateTaskItems更新TaskItem状态(遍历所有Record中的Item)
func (t *Task) UpdateTaskItems() {
	for index := range t.Items {
		quantity := 0
		for _, record := range t.Recrods {
			// 无效Task记录会跳过
			if record.State == state.INVALID {
				continue
			}
			for _, recordItem := range record.Items {
				if recordItem.ProductID == t.Items[index].ProductID {
					quantity += recordItem.Quantity
				}
			}
		}
		t.Items[index].Quantity = quantity
	}
}

// CheckTaskRecordItems: 保证记录中不会出现TaskItem中不存在的产品项
func (t *Task) CheckTaskRecordItems(items []record.RecordItem) error {
	set := make(map[string]struct{}, len(t.Items))
	for _, item := range t.Items {
		set[item.ProductID] = struct{}{}
	}
	for _, item := range items {
		if _, ok := set[item.ProductID]; !ok {
			return errors.New("item not found")
		}
	}
	return nil
}
