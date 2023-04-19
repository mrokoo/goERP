package repository

import (
	"github.com/google/uuid"
	category "github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/info"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/price"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	unit "github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	warehouse "github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

type Category = category.Category
type Unit = unit.Unit
type Warehouse = warehouse.Warehouse

// mysql product模型
type Product struct {
	ID           string      `gorm:"primaryKey"`
	Name         string      `gorm:"not null"`
	CategoryID   uuid.UUID   `gorm:"default:null"`
	Category     Category    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitID       uuid.UUID   `gorm:"default:null"`
	Unit         Unit        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OpeningStock []Stock     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	State        state.State `gorm:"default:active"`
	Note         string
	price.Price
	info.Info
}

func toMySQLProduct(product *domain.Product) *Product {
	openingStock := toMySQLStock(product.OpeningStock)
	return &Product{
		ID:           product.ID,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		UnitID:       product.UnitID,
		OpeningStock: openingStock,
		State:        product.State,
		Note:         product.Note,
		Price:        product.Price,
		Info:         product.Info,
	}
}

func (p *Product) toProduct() *domain.Product {
	var openingstock []stock.Stock
	for _, s := range p.OpeningStock {
		openingstock = append(openingstock, s.toStock())
	}
	return &domain.Product{
		ID:           p.ID,
		Name:         p.Name,
		CategoryID:   p.CategoryID,
		UnitID:       p.UnitID,
		OpeningStock: openingstock,
		State:        p.State,
		Note:         p.Note,
		Price:        p.Price,
		Info:         p.Info,
	}
}

// mysql Stock模型
type Stock struct {
	ProductID   string    `gorm:"primaryKey"`
	WarehouseID string    `gorm:"primaryKey"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:191"`
	Amount      int
}

// 后期需要给Stock重新命名表名，明确其为开盘库存。

func toMySQLStock(stocks []stock.Stock) []Stock {
	var mysqlstock []Stock
	for _, stock := range stocks {
		s := Stock{
			ProductID:   stock.ProductID,
			WarehouseID: stock.WarehouseID,
			Amount:      stock.Amount,
		}
		mysqlstock = append(mysqlstock, s)
	}
	return mysqlstock
}

func (s *Stock) toStock() stock.Stock {
	return stock.Stock{
		ProductID:   s.ProductID,
		WarehouseID: s.WarehouseID,
		Amount:      s.Amount,
	}
}
