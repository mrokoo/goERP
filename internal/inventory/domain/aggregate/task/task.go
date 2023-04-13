package task

import "github.com/mrokoo/goERP/internal/inventory/domain/valueobj/record"

type InTask struct {
	ID   string
	Type string
	InBasis
	State   string
	Status  string
	Records []record.InRecord
}

type InBasis struct {
	PurchaseOrderID string
	// 调拨、销售退货
}

type OutTask struct {
	ID   string
	Type string
	OutBasis
	State   string
	Status  string
	Records []record.OutRecord
}

type OutBasis struct {
	PurchaseReturnOrderID string
	// 调拨、销售
}
