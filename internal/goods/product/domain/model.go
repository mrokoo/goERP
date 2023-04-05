package domain

import (
	"errors"

	"github.com/google/uuid"
	category "github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/info"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/price"
	unit "github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

var ErrInvalidDate = errors.New("the date is invalid")

type Category = category.Category
type Unit = unit.Unit
type Product struct {
	ID         string      `gorm:"primaryKey"`
	Name       string      `gorm:"not null"`
	CategoryID uuid.UUID   `gorm:"default:null"`
	Category   Category    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitID     uuid.UUID   `gorm:"default:null"`
	Unit       Unit        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	State      state.State `gorm:"default:active"`
	Note       string
	price.Price
	info.Info
}
