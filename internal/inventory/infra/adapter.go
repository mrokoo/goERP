package repository

import (
	"time"

	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

type Notificaion struct {
	ID          string
	Name        string
	Type        bool
	WarehouseID string
	Warehouse   warehouse.Warehouse
	Basis       string
	State       bool   // true表示完成入库或出库; false表示未完成入库或出库操作
	Status      string // true表示正常;false表示已作废
	Records     []Record
}

type Record struct {
	NotificaionID string
	UserID        string
	Date          time.Time
	Note          string
	Items         []Item
}

type Item struct {
	ProductID string
	Quantity  int
}
