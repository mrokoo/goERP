package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/budget/app"
	infra "github.com/mrokoo/goERP/internal/share/budget/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func BudgetRoutes(r *gin.Engine, client *mongo.Client) {
	budget := r.Group("/budget")
	{
		m := infra.NewMongoRepository(client)
		s := app.NewBudgetServiceImpl(m)
		h := app.NewBudgetHandler(s)
		budget.GET("/getBudgetList", h.GetBudgetList)
		budget.POST("/addBudget", h.AddBudget)
		budget.DELETE("/deleteBudget", h.DeleteBudget)
		budget.PUT("/updateBudget", h.UpdateBudget)
	}
}
