package record

import (
	"time"

	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/item"
)

type InRecord struct {
	UserID string
	Date   time.Time
	Note   string
	Items  []item.InItem
}

type OutRecord struct {
	UserID string
	Date   time.Time
	Note   string
	Items  []item.OutItem
}
