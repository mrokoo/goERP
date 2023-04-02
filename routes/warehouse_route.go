package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/warehouse/app"
	domain "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	repository "github.com/mrokoo/goERP/internal/share/warehouse/infra"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewWarehouseRouter(db *mongo.Database, group *gin.RouterGroup) {
	m := repository.NewMongoRepository(db, domain.CollectionWarehouse)
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

