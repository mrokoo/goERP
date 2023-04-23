package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/record"
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type InventoryHandler struct {
	inventoryService InventoryService
}

func NewInventoryHandler(inventoryService InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

func (h *InventoryHandler) GetTaskList(ctx *gin.Context) {
	inventorys, err := h.inventoryService.GetTaskList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    inventorys,
	})
}

func (h *InventoryHandler) InvalidateTask(ctx *gin.Context) {
	err := h.inventoryService.InvalidateTask(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    nil,
	})
}

func (h *InventoryHandler) AddRecord(ctx *gin.Context) {
	type item struct {
		ProductID string
		Quantity  int
	}
	var req struct {
		WarehouseID string `json:"warehouseID" binding:"required"`
		UserID      string `json:"userID" binding:"required"`
		Items       []item
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var items []record.RecordItem
	for _, item := range req.Items {
		c := h.inventoryService.CreateRecordItem(item.ProductID, item.Quantity)
		items = append(items, c)
	}

	record := record.NewRecord(req.WarehouseID, req.UserID, items)
	err := h.inventoryService.AddRecord(ctx.Param("id"), record)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    nil,
	})
}

func (h *InventoryHandler) InvalidateRecord(ctx *gin.Context) {
	err := h.inventoryService.InvalidateRecord(ctx.Param("id"), ctx.Param("rid"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    nil,
	})
}

func (h *InventoryHandler) GetTakeList(ctx *gin.Context) {
	takes, err := h.inventoryService.GetTakeList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    takes,
	})
}

func (h *InventoryHandler) AddTake(ctx *gin.Context) {
	type item struct {
		ProductID string
		Quantity  int
	}
	var req struct {
		WarehouseID string `json:"warehouseID" binding:"required"`
		UserID      string `json:"userID" binding:"required"`
		Items       []item
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var takeItems []take.Item
	for _, item := range req.Items {
		takeItems = append(takeItems, take.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	err := h.inventoryService.CreateTake(req.WarehouseID, req.UserID, takeItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    nil,
	})
}

func (h *InventoryHandler) GetAllotList(ctx *gin.Context) {
	allots, err := h.inventoryService.GetAllotList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    allots,
	})
}

func (h *InventoryHandler) AddAllot(ctx *gin.Context) {
	type Item struct {
		ProductID string
		Quantity  int
	}

	var req struct {
		InWarehouseID  string `json:"inWarehouseID" binding:"required"`
		OutWarehouseID string `json:"outWarehouseID" binding:"required"`
		UserID         string `json:"userID" binding:"required"`
		Items          []Item
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var allotItems []allot.Item
	for _, item := range req.Items {
		allotItems = append(allotItems, allot.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	err := h.inventoryService.CreateAllot(req.InWarehouseID, req.OutWarehouseID, req.UserID, allotItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    nil,
	})
}

func (h *InventoryHandler) GetInventoryFlowList(ctx *gin.Context) {
	inventoryFlows, err := h.inventoryService.GetFlowList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    inventoryFlows,
	})
}
