package repository

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

func toModel(product *domain.Product) *model.Product {
	var openingStock []model.OpeningStock
	for i := range product.OpeningStock {
		openingStock = append(openingStock, *toModelStock(&product.OpeningStock[i]))
	}
	return &model.Product{
		ID:           product.ID,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		UnitID:       product.UnitID,
		OpeningStock: openingStock,
		State:        string(product.State),
		Note:         product.Note,
		Img:          product.Img,
		Intro:        product.Intro,
		Purchase:     product.Purchase,
		Retail:       product.Retail,
		Grade1:       product.Grade1,
		Grade2:       product.Grade2,
		Grade3:       product.Grade3,
	}
}

func toDomain(product *model.Product) *domain.Product {
	var openingStock []stock.Stock
	for i := range product.OpeningStock {
		openingStock = append(openingStock, *toDomainStock(&product.OpeningStock[i]))
	}
	return &domain.Product{
		ID:           product.ID,
		Name:         product.Name,
		CategoryID:   product.CategoryID,
		UnitID:       product.UnitID,
		OpeningStock: openingStock,
		State:        state.State(product.State),
		Note:         product.Note,
		Img:          product.Img,
		Intro:        product.Intro,
		Purchase:     product.Purchase,
		Retail:       product.Retail,
		Grade1:       product.Grade1,
		Grade2:       product.Grade2,
		Grade3:       product.Grade3,
	}
}

func toModelStock(stock *stock.Stock) *model.OpeningStock {
	return &model.OpeningStock{
		WarehouseID: stock.WarehouseID,
		Amount:      stock.Amount,
	}
}

func toDomainStock(s *model.OpeningStock) *stock.Stock {
	return &stock.Stock{
		WarehouseID: s.WarehouseID,
		Amount:      s.Amount,
	}
}
