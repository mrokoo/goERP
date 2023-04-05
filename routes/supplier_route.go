package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/supplier/app"
	domain "github.com/mrokoo/goERP/internal/share/supplier/domain"
	repository "github.com/mrokoo/goERP/internal/share/supplier/infra"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSupplierRouter(db *mongo.Database, group *gin.RouterGroup) {
	m := repository.NewMongoRepository(db, domain.CollectionSupplier)
	ds := domain.NewCheckingSupplierValidityService(m)
	s := app.NewSupplierServiceImpl(ds, m)
	h := app.NewSupplierHandler(s)
	group.GET("/suppliers", h.GetSupplierList)
	group.GET("/suppliers/:id", h.GetSupplier)
	group.POST("/suppliers", h.AddSupplier)
	group.PUT("/suppliers/:id", h.ReplaceSupplier)
	group.PATCH("/suppliers/:id", h.ReplaceSupplier)
	group.DELETE("/suppliers/:id", h.DeleteSupplier)
}
