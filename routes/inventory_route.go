package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/inventory/app"
	allotRepo "github.com/mrokoo/goERP/internal/inventory/infra/allot_repository"
	flowRepo "github.com/mrokoo/goERP/internal/inventory/infra/inventoryflow_repository"
	takeRepo "github.com/mrokoo/goERP/internal/inventory/infra/take_repository"
	taskRepo "github.com/mrokoo/goERP/internal/inventory/infra/task_repository"
	"gorm.io/gorm"
)

func NewInventoryRouter(db *gorm.DB, group *gin.RouterGroup) {
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)
	s := app.NewInventoryServiceImpl(flow, task, allot, take)
	h := app.NewInventoryHandler(s)

	group.GET("/tasks", h.GetTaskList)
	group.PATCH("/tasks/:id", h.InvalidateTask)
	group.POST("/tasks/:id/records", h.AddRecord)
	group.PATCH("/tasks/:id/records/:rid", h.InvalidateRecord)
	group.GET("/takes", h.GetTakeList)
	group.POST("/takes", h.AddTake)
	group.GET("/allots", h.GetAllotList)
	group.POST("/allots", h.AddAllot)
	group.GET("/inventoryflows", h.GetInventoryFlowList)
}
