package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/supplier/app"
	domain "github.com/mrokoo/goERP/internal/share/supplier/domain"
	infra "github.com/mrokoo/goERP/internal/share/supplier/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func SupplierRoutes(r *gin.Engine, client *mongo.Client) {
	supplier := r.Group("/supplier")
	{
		m := infra.NewMongoRepository(client)
		ds := domain.NewCheckingSupplierValidityService(m)
		s := app.NewSupplierServiceImpl(ds, m)
		h := app.NewSupplierHandler(s)
		supplier.GET("/getSupplierList", h.GetSupplierList)
		supplier.POST("/addSupplier", h.AddSupplier)
		supplier.DELETE("/deleteSupplier", h.DeleteSupplier)
		supplier.PUT("/updateSupplier", h.UpdateSupplier)
	}
}
