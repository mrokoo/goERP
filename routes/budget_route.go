package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/budget/app"
	domain "github.com/mrokoo/goERP/internal/share/budget/domain"
	repository "github.com/mrokoo/goERP/internal/share/budget/infra"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBudgetRouter(db *mongo.Database, group *gin.RouterGroup) {
	m := repository.NewMongoRepository(db, domain.CollectionBudget)
	s := app.NewBudgetServiceImpl(m)
	h := app.NewBudgetHandler(s)
	group.GET("/budgets", h.GetBudgetList)
	group.GET("/budgets/:id", h.GetBudget)
	group.POST("/budgets", h.AddBudget)
	group.PUT("/budgets/:id", h.ReplaceBudget)
	group.PATCH("/budgets/:id", h.ReplaceBudget)
	group.DELETE("/budgets/:id", h.DeleteBudget)
}
