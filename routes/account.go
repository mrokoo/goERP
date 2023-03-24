package routes

import (
	"github.com/gin-gonic/gin"
	accountapp "github.com/mrokoo/goERP/internal/share/account/app"
)

func AccountRoutes(r *gin.Engine, h accountapp.AccountHandler) {
	account := r.Group("/account")
	{
		account.GET("/getAccountList", h.GetAccountList)
		account.POST("/addAccountList", h.AddAccount)
		account.DELETE("/deleteAccount", h.DeleteAccount)
		account.PUT("/updateAccount", h.UpdateAccount)
	}
}
