package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/goods/product/app"
	domain "github.com/mrokoo/goERP/internal/goods/product/domain"
	repository "github.com/mrokoo/goERP/internal/goods/product/infra"
	"gorm.io/gorm"
)

func NewProductRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewProductRepository(db)
	ds := domain.NewCheckingProductValidityService(m)
	s := app.NewProductServiceImpl(ds, m)
	h := app.NewProductHandler(s)
	group.GET("/products", h.GetProductList)
	group.GET("/products/:id", h.GetProduct)
	group.POST("/products", h.AddProduct)
	group.PUT("/products/:id", h.ReplaceProduct)
	group.PATCH("/products/:id", h.ReplaceProduct)
	group.DELETE("/products/:id", h.DeleteProduct)
}
