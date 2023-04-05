package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/info"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/price"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

var ErrInvalidDate = errors.New("the date is invalid")

type Product struct {
	ID         string      `json:"id" binding:"required"`
	Name       string      `json:"name" binding:"required"`
	CategoryID uuid.UUID   `json:"categoryId"`
	UnitID     uuid.UUID   `json:"unitId"`
	State      state.State `json:"state" binding:"oneof=active freeze"`
	Note       string      `json:"note"`
	price.Price
	info.Info
	OpeningStock []stock.Stock `json:"openingStock" bson:"opening_stock"`
}
