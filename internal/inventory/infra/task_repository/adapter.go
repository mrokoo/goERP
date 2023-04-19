package task_repository

import (
	"time"

	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/state"
	"github.com/mrokoo/goERP/internal/purchase/infra/order"
	"github.com/mrokoo/goERP/internal/purchase/infra/returnorder"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

type MySQLTask struct {
	ID                    string              `gorm:"primaryKey;size:191;"`
	WarehouseID           string              `gorm:"size:191; not null"`
	Warehouse             warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	Kind                  task.Kind
	State                 state.State
	Items                 []MySQLTaskItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Recrods               []MySQLRecord   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;default:null;"`
	IO                    bool
	PurchaseOrderID       *string                         `gorm:"size:191;default:null;"`
	PurchaseOrder         order.PurchaseOrder             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	PurchaseReturnOrderID *string                         `gorm:"size:191;default:null;"`
	PurchaseReturnOrder   returnorder.PurchaseReturnOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	SaleOrderID           *string                         `gorm:"size:191;"`
	SaleReturnOrderID     *string                         `gorm:"size:191;"`
	AllocationOrderID     *string                         `gorm:"size:191;"`
	CreatedAt             time.Time
}

type MySQLTaskItem struct {
	ID          string          `gorm:"size:191;primaryKey"`
	MySQLTaskID string          `gorm:"size:191;primaryKey"` // 外键约束
	ProductID   string          `gorm:"size:191;primaryKey"`
	Product     product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Total       int
	Quantity    int
}

type MySQLRecord struct {
	ID          string              `gorm:"size:191;primaryKey"`
	MySQLTaskID string              `gorm:"size:191;primaryKey"` // 外键约束
	WarehouseID string              `gorm:"size:191;"`
	Warehouse   warehouse.Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` //外键约束
	UserID      string              `gorm:"size:191;"`
	User        user.User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	State       state.State
	Items       []MySQLRecordItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
}

type MySQLRecordItem struct {
	ID            string          `gorm:"size:191;primaryKey"`
	MySQLRecordID string          `gorm:"size:191;primaryKey"` // 外键约束
	ProductID     string          `gorm:"size:191;primaryKey"`
	Product       product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"` // 外键约束
	Quantity      int
}

func (t MySQLTask) TableName() string {
	return "tasks"
}

func (t MySQLTaskItem) TableName() string {
	return "task_items"
}

func (t MySQLRecord) TableName() string {
	return "task_records"
}

func (t MySQLRecordItem) TableName() string {
	return "task_record_items"
}

func toMySQLTask(task task.Task) MySQLTask {
	return MySQLTask{
		ID:                    task.ID,
		WarehouseID:           task.WarehouseID,
		Kind:                  task.Kind,
		State:                 task.State,
		Items:                 toMySQLTaskItems(task.ID, task.Items),
		Recrods:               toMySQLRecords(task.ID, task.Recrods),
		IO:                    task.IO,
		PurchaseOrderID:       task.PurchaseOrderID,
		PurchaseReturnOrderID: task.PurchaseReturnOrderID,
		SaleOrderID:           task.SaleOrderID,
		SaleReturnOrderID:     task.SaleReturnOrderID,
		AllocationOrderID:     task.AllocationOrderID,
		CreatedAt:             task.CreatedAt,
	}
}

func toMySQLTaskItems(MySQLTaskID string, items []task.TaskItem) []MySQLTaskItem {
	var taskItems []MySQLTaskItem
	for _, item := range items {
		taskItems = append(taskItems, MySQLTaskItem{
			ID:          item.ID,
			MySQLTaskID: MySQLTaskID,
			ProductID:   item.ProductID,
			Total:       item.Total,
			Quantity:    item.Quantity,
		})
	}
	return taskItems
}

func toMySQLRecords(MySQLTaskID string, records []record.Record) []MySQLRecord {
	var records_ []MySQLRecord
	for _, r := range records {
		records_ = append(records_, MySQLRecord{
			ID:          r.ID,
			MySQLTaskID: MySQLTaskID,
			WarehouseID: r.WarehouseID,
			UserID:      r.UserID,
			State:       r.State,
			Items:       toMySQLRecordItems(r.ID, r.Items),
			CreatedAt:   r.CreatedAt,
		})
	}
	return records_
}

func toMySQLRecordItems(MySQLRecordID string, items []record.RecordItem) []MySQLRecordItem {
	var recordItems []MySQLRecordItem
	for _, item := range items {
		recordItems = append(recordItems, MySQLRecordItem{
			ID:            item.ID,
			MySQLRecordID: MySQLRecordID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
		})
	}
	return recordItems
}

func (t *MySQLTask) toTask() task.Task {
	return task.Task{
		ID:                    t.ID,
		WarehouseID:           t.WarehouseID,
		Kind:                  t.Kind,
		State:                 t.State,
		Items:                 toTaskItems(t.Items),
		Recrods:               toRecords(t.Recrods),
		IO:                    t.IO,
		PurchaseOrderID:       t.PurchaseOrderID,
		PurchaseReturnOrderID: t.PurchaseReturnOrderID,
		SaleOrderID:           t.SaleOrderID,
		SaleReturnOrderID:     t.SaleReturnOrderID,
		AllocationOrderID:     t.AllocationOrderID,
		CreatedAt:             t.CreatedAt,
	}
}

func (t *MySQLTaskItem) toTaskItem() task.TaskItem {
	return task.TaskItem{
		ID:        t.ID,
		ProductID: t.ProductID,
		Total:     t.Total,
		Quantity:  t.Quantity,
	}
}

func toTaskItems(items []MySQLTaskItem) []task.TaskItem {
	var taskItems []task.TaskItem
	for _, item := range items {
		taskItems = append(taskItems, item.toTaskItem())
	}
	return taskItems
}

func (r *MySQLRecord) toRecord() record.Record {
	return record.Record{
		ID:          r.ID,
		WarehouseID: r.WarehouseID,
		UserID:      r.UserID,
		State:       r.State,
		Items:       toRecordItems(r.Items),
		CreatedAt:   r.CreatedAt,
	}
}

func toRecords(records []MySQLRecord) []record.Record {
	var records2 []record.Record
	for _, r := range records {
		records2 = append(records2, r.toRecord())
	}
	return records2
}

func (ri *MySQLRecordItem) toRecordItem() record.RecordItem {
	return record.RecordItem{
		ID:        ri.ID,
		ProductID: ri.ProductID,
		Quantity:  ri.Quantity,
	}
}

func toRecordItems(items []MySQLRecordItem) []record.RecordItem {
	var recordItems []record.RecordItem
	for _, item := range items {
		recordItems = append(recordItems, item.toRecordItem())
	}
	return recordItems
}
