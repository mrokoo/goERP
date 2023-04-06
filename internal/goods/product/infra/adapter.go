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

type Product = domain.Product
type Stock = stock.Stock
type Category = category.Category
type Unit = unit.Unit
type Warehouse = warehouse.Warehouse

// mysql product模型
type ProductModel struct {
	ID           string       `gorm:"primaryKey"`
	Name         string       `gorm:"not null"`
	CategoryID   uuid.UUID    `gorm:"default:null"`
	Category     Category     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UnitID       uuid.UUID    `gorm:"default:null"`
	Unit         Unit         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OpeningStock []StockModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	State        state.State  `gorm:"default:active"`
	Note         string
	price.Price
	info.Info
}

func toProductModel(product *Product) *ProductModel {
	openingStock := toStockModel(product.OpeningStock)
	return &ProductModel{
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

func (p *ProductModel) toProduct() *Product {
	var openingstock []Stock
	for _, s := range p.OpeningStock {
		openingstock = append(openingstock, s.toStock())
	}
	return &Product{
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
type StockModel struct {
	ProductID   string    `gorm:"primaryKey"`
	WarehouseID string    `gorm:"primaryKey"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:191"`
	Amount      int
}

func toStockModel(stocks []Stock) []StockModel {
	var stockModels []StockModel
	for _, stock := range stocks {
		s := StockModel{
			ProductID:   stock.ProductID,
			WarehouseID: stock.WarehouseID,
			Amount:      stock.Amount,
		}
		stockModels = append(stockModels, s)
	}
	return stockModels
}

func (s *StockModel) toStock() Stock {
	return Stock{
		ProductID:   s.ProductID,
		WarehouseID: s.WarehouseID,
		Amount:      s.Amount,
	}
}
