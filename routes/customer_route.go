package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/customer/app"
	domain "github.com/mrokoo/goERP/internal/share/customer/domain"
	repository "github.com/mrokoo/goERP/internal/share/customer/infra"
	"gorm.io/gorm"
)

func NewCustomerRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewCustomerRepository(db)
	ds := domain.NewCheckingCustomerValidityService(m)
	s := app.NewCustomerServiceImpl(ds, m)
	h := app.NewCustomerHandler(s)
	group.GET("/customers", h.GetCustomerList)
	group.GET("/customers/:id", h.GetCustomer)
	group.POST("/customers", h.AddCustomer)
	group.PUT("/customers/:id", h.ReplaceCustomer)
	group.PATCH("/customers/:id", h.ReplaceCustomer)
	group.DELETE("/customers/:id", h.DeleteCustomer)
}
