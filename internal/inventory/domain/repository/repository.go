package domain

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/flow"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/notification"
	"github.com/mrokoo/goERP/internal/inventory/domain/valueobj/record"
)

type Repository[T flow.InventoryFlow | record.Record | notification.Notificaion] interface {
	GetAll() ([]*T, error)
	GetByID(ID string) (*T, error)
	Save(e *T) error
	Replace(e *T) error
	Delete(ID string) error
}
