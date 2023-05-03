package domain

import (
	category "github.com/mrokoo/goERP/internal/goods/category/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	unit "github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Category = category.Category
type Unit = unit.Unit
type Product struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	CategoryID   *string       `json:"category_id"`
	UnitID       *string       `json:"unit_id"`
	OpeningStock []stock.Stock `json:"openStock"`
	State        state.State   `json:"state"`	
	Note         string        `json:"note"`
	Img          string        `json:"img"`
	Intro        string        `json:"intro"`
	Purchase     float64       `json:"purchase"`
	Retail       float64       `json:"retail"`
	Grade1       float64       `json:"grade1"`
	Grade2       float64       `json:"grade2"`
	Grade3       float64       `json:"grade3"`
}
