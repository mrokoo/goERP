package domain

import (
	"github.com/google/uuid"
	category "github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/info"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/price"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	unit "github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Category = category.Category
type Unit = unit.Unit
type Product struct {
	ID           string
	Name         string
	CategoryID   uuid.UUID
	UnitID       uuid.UUID
	OpeningStock []stock.Stock
	State        state.State
	Note         string
	price.Price
	info.Info
}
