package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/account/app"
	domain "github.com/mrokoo/goERP/internal/share/account/domain"
	infra "github.com/mrokoo/goERP/internal/share/account/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func AccountRoutes(r *gin.Engine, client *mongo.Client) {
	account := r.Group("/account")
	{
		m := infra.NewMongoRepository(client)
		ds := domain.NewCheckingAccountValidityService(m)
		s := app.NewAccountServiceImpl(ds, m)
		h := app.NewAccountHandler(s)
		account.GET("/getAccountList", h.GetAccountList)
		account.POST("/addAccountList", h.AddAccount)
		account.DELETE("/deleteAccount", h.DeleteAccount)
		account.PUT("/updateAccount", h.UpdateAccount)
	}
}
