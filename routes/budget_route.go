package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/share/budget/app"
	repository "github.com/mrokoo/goERP/internal/share/budget/infra"
	"gorm.io/gorm"
)

func NewBudgetRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewBudgetRepository(db)
	s := app.NewBudgetServiceImpl(m)
	h := app.NewBudgetHandler(s)
	group.GET("/budgets", h.GetBudgetList)
	group.GET("/budgets/:id", h.GetBudget)
	group.POST("/budgets", h.AddBudget)
	group.PUT("/budgets/:id", h.ReplaceBudget)
	group.PATCH("/budgets/:id", h.ReplaceBudget)
	group.DELETE("/budgets/:id", h.DeleteBudget)
}
