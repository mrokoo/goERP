package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/account/app"
	domain "github.com/mrokoo/goERP/internal/share/account/domain"
	repository "github.com/mrokoo/goERP/internal/share/account/infra"
	"gorm.io/gorm"
)

func NewAccountRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewAccountRepository(db)
	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)
	h := app.NewAccountHandler(s)
	group.GET("/accounts", h.GetAccountList)
	group.GET("/accounts/:id", h.GetAccount)
	group.POST("/accounts", h.AddAccount)
	group.PUT("/accounts/:id", h.ReplaceAccount)
	group.PATCH("/accounts/:id", h.ReplaceAccount)
	group.DELETE("/accounts/:id", h.DeleteAccount)
}
