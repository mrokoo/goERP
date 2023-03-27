
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
)

type WarehouseHandler struct {
	WarehouseService WarehouseService
}

func NewWarehouseHandler(warehouseService WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		WarehouseService: warehouseService,
	}
}

func (h *WarehouseHandler) GetWarehouseList(ctx *gin.Context) {
	warehouses, err := h.WarehouseService.GetWarehouseList()
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":     1,
		"showMsg":  "success",
		"errorMsg": "",
		"data":     warehouses,
	})
}

func (h *WarehouseHandler) AddWarehouse(ctx *gin.Context) {
	var req domain.Warehouse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.WarehouseService.AddWarehouse(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":     1,
		"showMsg":  "success",
		"errorMsg": "",
		"data":     nil,
	})
}

func (h *WarehouseHandler) UpdateWarehouse(ctx *gin.Context) {
	var req domain.Warehouse
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.WarehouseService.UpdateWarehouse(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":     1,
		"showMsg":  "success",
		"errorMsg": "",
		"data":     nil,
	})
}

func (h *WarehouseHandler) DeleteWarehouse(ctx *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := h.WarehouseService.DeleteWarehouse(req.ID); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":     1,
		"showMsg":  "success",
		"errorMsg": "",
		"data":     nil,
	})
}
