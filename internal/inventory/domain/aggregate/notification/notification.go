package notification

import "github.com/mrokoo/goERP/internal/inventory/domain/valueobj/record"

type Notificaion struct {
	ID          string
	Name        string
	Type        bool // true is in; false is out;
	WarehouseID string
	Basis       string
	State       string
	Records     []record.Record
}
