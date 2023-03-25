package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/customer/app"
	domain "github.com/mrokoo/goERP/internal/share/customer/domain"
	infra "github.com/mrokoo/goERP/internal/share/customer/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func CustomerRoutes(r *gin.Engine, client *mongo.Client) {
	customer := r.Group("/customer")
	{
		m := infra.NewMongoRepository(client)
		ds := domain.NewCheckingCustomerValidityService(m)
		s := app.NewCustomerServiceImpl(ds, m)
		h := app.NewCustomerHandler(s)
		customer.GET("/getCustomerList", h.GetCustomerList)
		customer.POST("/addCustomer", h.AddCustomer)
		customer.DELETE("/deleteCustomer", h.DeleteCustomer)
		customer.PUT("/updateCustomer", h.UpdateCustomer)
	}
}
