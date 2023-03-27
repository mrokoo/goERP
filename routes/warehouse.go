package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/warehouse/app"
	domain "github.com/mrokoo/goERP/internal/share/warehouse/domain"
	infra "github.com/mrokoo/goERP/internal/share/warehouse/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func WarehouseRoutes(r *gin.Engine, client *mongo.Client) {
	warehouse := r.Group("/warehouse")
	{
		m := infra.NewMongoRepository(client)
		ds := domain.NewCheckingWarehouseValidityService(m)
		s := app.NewWarehouseServiceImpl(ds, m)
		h := app.NewWarehouseHandler(s)
		warehouse.GET("/getWarehouseList", h.GetWarehouseList)
		warehouse.POST("/addWarehouse", h.AddWarehouse)
		warehouse.DELETE("/deleteWarehouse", h.DeleteWarehouse)
		warehouse.PUT("/updateWarehouse", h.UpdateWarehouse)
	}
}
