package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj"
	warehouseDomain "github.com/mrokoo/goERP/internal/share/warehouse"
)

type Product struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	CategoryID         uuid.UUID          `json:"categoryId"`
	UnitID             uuid.UUID          `json:"unitId"`
	ExpirationDay      int                `json:"expirationDay" bson:"expiration_day"`
	AlertExpirationDay int                `json:"alertExpirationDay" bson:"alert_expiration_day"`
	State              valueobj.StateType `json:"state"`
	Note               string             `json:"note"`
	Price              valueobj.Price     `json:"price"`
	Info               valueobj.Info      `json:"info"`
	OpeningStock       []OpeningStock     `json:"openingStock"`
}

type OpeningStock struct {
	Warehouse warehouseDomain.WarehouseId
	Amount    int
}

func (p *Product) Validate() error {
	if p.ExpirationDay <= p.AlertExpirationDay {
		return errors.New("ExpirationDay must be greater than AlertExpirationDay")
	}
	return nil
}
