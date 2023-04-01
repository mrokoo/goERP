package inbound

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/item"
)

type InboundNotificaion struct {
	ID          string
	Type        string
	WarehouseID string
	Basis       string
	Items       item.Item
	Date        string // 创建通知单的日期
}
