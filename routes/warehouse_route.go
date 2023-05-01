package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/warehouse/app"
	domain "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	repository "github.com/mrokoo/goERP/internal/share/warehouse/infra"
	"gorm.io/gorm"
)

func NewWarehouseRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewWarehouseRepository(db)
	ds := domain.NewCheckingWarehouseValidityService(m)
	s := app.NewWarehouseServiceImpl(ds, m)
	h := app.NewWarehouseHandler(s)
	group.GET("/warehouses", h.GetWarehouseList)
	group.GET("/warehouses/:id", h.GetWarehouse)
	group.POST("/warehouses", h.AddWarehouse)
	group.PUT("/warehouses/:id", h.ReplaceWarehouse)
	group.PATCH("/warehouses/:id", h.ReplaceWarehouse)
	group.DELETE("/warehouses/:id", h.DeleteWarehouse)
}
