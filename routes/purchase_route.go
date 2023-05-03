package routes

import (
	"github.com/gin-gonic/gin"
	inventory "github.com/mrokoo/goERP/internal/inventory/app"
	allotRepo "github.com/mrokoo/goERP/internal/inventory/infra/allot_repository"
	flowRepo "github.com/mrokoo/goERP/internal/inventory/infra/inventoryflow_repository"
	takeRepo "github.com/mrokoo/goERP/internal/inventory/infra/take_repository"
	taskRepo "github.com/mrokoo/goERP/internal/inventory/infra/task_repository"
	"github.com/mrokoo/goERP/internal/purchase/app"
	order "github.com/mrokoo/goERP/internal/purchase/infra"
	"gorm.io/gorm"
)

func NewPurchaseRouter(db *gorm.DB, group *gin.RouterGroup) {
	m1 := order.NewPurchaseOrderRepository(db)
	take := takeRepo.NewTakeRepository(db)
	task := taskRepo.NewTaskRepository(db)
	flow := flowRepo.NewInventoryFlowRepository(db)
	allot := allotRepo.NewAllotRepository(db)

	m3 := inventory.NewInventoryServiceImpl(flow, task, allot, take)

	s := app.NewPurchaseServiceImpl(m1, m3)
	h := app.NewPurchaseHandler(s)
	group.GET("/purchaseOrders", h.GetPurchaseOrderList)
	group.POST("/purchaseOrders", h.AddPurchaseOrder)
	group.PUT("/purchaseOrders/:id", h.InvalidatePurchaseOrder)
	group.POST("/purchaseReturnOrders", h.AddPurchaseReturnOrder)
	group.PUT("/purchaseReturnOrders/:id", h.InvalidatePurchaseReturnOrder)
}
