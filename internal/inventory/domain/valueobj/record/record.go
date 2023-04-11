package record

import (
	"time"

	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/item"
)

type Record struct {
	NotificaionID string
	UserID        string
	Date          time.Time
	Note          string
	Items         []item.Item
}
