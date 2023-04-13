package task_repository

import (
	"time"

	"github.com/google/uuid"
	product "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/task"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/item"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/record"
	"github.com/mrokoo/goERP/internal/purchase/infra/order"
	"github.com/mrokoo/goERP/internal/purchase/infra/returnorder"
	user "github.com/mrokoo/goERP/internal/system/user/domain"
)

// -------------入库任务MySQL模型---------------------

type InTask struct {
	ID   string `gorm:"primaryKey;size:191;"`
	Type string
	InBasis
	State   string
	Status  string
	Records []InRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// toMySQLInTask 转换为InTask
func toMySQLInTask(task *task.InTask) *InTask {
	return &InTask{
		ID:      task.ID,
		Type:    task.Type,
		InBasis: toMySQLInBasis(task.InBasis),
		State:   task.State,
		Status:  task.Status,
		Records: toMySQLInRecords(task.Records),
	}
}

func (i *InTask) toTask() *task.InTask {
	var records []record.InRecord
	for _, ir := range i.Records {
		records = append(records, ir.toInRecord())
	}
	return &task.InTask{
		ID:      i.ID,
		Type:    i.Type,
		InBasis: i.InBasis.toInBasis(),
		State:   i.State,
		Status:  i.Status,
		Records: records,
	}
}

type InBasis struct {
	PurchaseOrderID string              `gorm:"size:191"`
	PurchaseOrder   order.PurchaseOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// to do 调拨、销售退货
}

func toMySQLInBasis(inBasis task.InBasis) InBasis {
	return InBasis{
		PurchaseOrderID: inBasis.PurchaseOrderID,
	}
}

func (i InBasis) toInBasis() task.InBasis {
	return task.InBasis{
		PurchaseOrderID: i.PurchaseOrderID,
	}
}

type InRecord struct {
	ID       string    `gorm:"primaryKey;size:191;"`
	InTaskID string    `gorm:"size:191"`
	UserID   string    `gorm:"size:191"`
	User     user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Date     time.Time
	Note     string
	Items    []InItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func toMySQLInRecords(inRecords []record.InRecord) []InRecord {
	var records []InRecord
	for _, inRecord := range inRecords {
		records = append(records, toMySQLInRecord(inRecord))
	}
	return records
}

func toMySQLInRecord(inRecord record.InRecord) InRecord {
	return InRecord{
		ID:     uuid.New().String(),
		UserID: inRecord.UserID,
		Date:   inRecord.Date,
		Note:   inRecord.Note,
		Items:  toMySQLInItems(inRecord.Items),
	}
}

func (i InRecord) toInRecord() record.InRecord {
	var items []item.InItem
	for _, v := range i.Items {
		items = append(items, v.toInItem())
	}
	return record.InRecord{
		UserID: i.UserID,
		Date:   i.Date,
		Note:   i.Note,
		Items:  items,
	}
}

type InItem struct {
	ID         string          `gorm:"primaryKey;size:191;"`
	InRecordID string          `gorm:"size:191"`
	ProductID  string          `gorm:"size:191"`
	Product    product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity   int
}

func toMySQLInItems(inItems []item.InItem) []InItem {
	var items []InItem
	for _, inItem := range inItems {
		items = append(items, toMySQLInItem(inItem))
	}
	return items
}

func toMySQLInItem(inItem item.InItem) InItem {
	return InItem{
		ID:        uuid.New().String(),
		ProductID: inItem.ProductID,
		Quantity:  inItem.Quantity,
	}
}

func (i *InItem) toInItem() item.InItem {
	return item.InItem{
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}

// -------------出库任务MySQL模型---------------------

type OutTask struct {
	ID   string `gorm:"primaryKey;size:191;"`
	Type string
	OutBasis
	State   string
	Status  string
	Records []OutRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func toMySQLOutTask(task *task.OutTask) *OutTask {
	return &OutTask{
		ID:       task.ID,
		Type:     task.Type,
		OutBasis: toMySQLOutBasis(task.OutBasis),
		State:    task.State,
		Status:   task.Status,
		Records:  toMySQLOutRecords(task.Records),
	}
}

func (o OutTask) toTask() *task.OutTask {
	var records []record.OutRecord
	for _, outRecord := range o.Records {
		records = append(records, outRecord.toOutRecord())
	}
	return &task.OutTask{
		ID:       o.ID,
		Type:     o.Type,
		OutBasis: o.OutBasis.toOutBasis(),
		State:    o.State,
		Status:   o.Status,
		Records:  records,
	}
}

type OutBasis struct {
	PurchaseReturnOrderID string                          `gorm:"size:191"`
	PurchaseReturnOrder   returnorder.PurchaseReturnOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// to do 调拨、销售退货
}

func toMySQLOutBasis(outBasis task.OutBasis) OutBasis {
	return OutBasis{
		PurchaseReturnOrderID: outBasis.PurchaseReturnOrderID,
	}
}

func (o OutBasis) toOutBasis() task.OutBasis {
	return task.OutBasis{
		PurchaseReturnOrderID: o.PurchaseReturnOrderID,
	}
}

type OutRecord struct {
	ID        string    `gorm:"primaryKey;size:191;"`
	OutTaskID string    `gorm:"size:191"`
	UserID    string    `gorm:"size:191"`
	User      user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Date      time.Time
	Note      string
	Items     []OutItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func toMySQLOutRecords(outRecords []record.OutRecord) []OutRecord {
	var records []OutRecord
	for _, outRecord := range outRecords {
		records = append(records, toMySQLOutRecord(outRecord))
	}
	return records
}

func toMySQLOutRecord(outRecord record.OutRecord) OutRecord {
	return OutRecord{
		ID:     uuid.New().String(),
		UserID: outRecord.UserID,
		Date:   outRecord.Date,
		Note:   outRecord.Note,
		Items:  toMySQLOutItems(outRecord.Items),
	}
}

func (o *OutRecord) toOutRecord() record.OutRecord {
	var items []item.OutItem
	for _, v := range o.Items {
		items = append(items, v.toOutItem())
	}
	return record.OutRecord{
		UserID: o.UserID,
		Date:   o.Date,
		Note:   o.Note,
		Items:  items,
	}
}

type OutItem struct {
	ID          string          `gorm:"primaryKey;size:191;"`
	OutRecordID string          `gorm:"size:191"`
	ProductID   string          `gorm:"size:191"`
	Product     product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity    int
}

func toMySQLOutItems(outItems []item.OutItem) []OutItem {
	var items []OutItem
	for _, outItem := range outItems {
		items = append(items, toMySQLOutItem(outItem))
	}
	return items
}

func toMySQLOutItem(outItem item.OutItem) OutItem {
	return OutItem{
		ID:        uuid.New().String(),
		ProductID: outItem.ProductID,
		Quantity:  outItem.Quantity,
	}
}

func (o *OutItem) toOutItem() item.OutItem {
	return item.OutItem{
		ProductID: o.ProductID,
		Quantity:  o.Quantity,
	}
}
