package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/biling"
	"github.com/mrokoo/goERP/internal/purchase/domain/valueobj/item"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type PurchaseHandler struct {
	PurchaseService PurchaseService
}

func NewPurchaseHandler(purchaseService PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		PurchaseService: purchaseService,
	}
}

func (h *PurchaseHandler) GetPurchaseOrderList(ctx *gin.Context) {
	purchaseOrders, err := h.PurchaseService.GetPurchaseOrderList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    purchaseOrders,
	})
}

func (h *PurchaseHandler) AddPurchaseOrder(ctx *gin.Context) {
	var req struct {
		ID          string `json:"id" binding:"required"`
		SupplierID  string `json:"supplier_id" binding:"required"`
		WarehouseID string `json:"warehouse_id" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		Items       []item.OrderItem
		biling.Biling
		IsValidated bool
		Note        string
		Date        string
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	date, err := time.Parse("RFC3339Nano", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := domain.PurchaseOrder{
		ID:          req.ID,
		SupplierID:  req.SupplierID,
		WarehouseID: req.WarehouseID,
		UserID:      req.UserID,
		Items:       req.Items,
		Biling:      req.Biling,
		IsValidated: req.IsValidated,
		Note:        req.Note,
		Date:        date,
	}
	err = h.PurchaseService.AddPurchaseOrder(&order)
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

func (h *PurchaseHandler) InvalidatePurchaseOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.PurchaseService.InvalidatePurchaseOrder(id)
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

func (h *PurchaseHandler) GetPurchaseReturnOrderList(ctx *gin.Context) {
	purchaseReturnOrders, err := h.PurchaseService.GetPurchaseReturnOrderList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    purchaseReturnOrders,
	})
}

func (h *PurchaseHandler) AddPurchaseReturnOrder(ctx *gin.Context) {
	var req struct {
		ID          string `json:"id" binding:"required"`
		PurchaseID  string `json:"purchase_id" binding:"required"`
		SupplierID  string `json:"supplier_id" binding:"required"`
		WarehouseID string `json:"warehouse_id" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		Items       []item.ReturnOrderItem
		biling.Biling
		IsValidated bool
		Note        string
		Date        string
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	date, err := time.Parse("RFC3339Nano", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := domain.PurchaseReturnOrder{
		ID:              req.ID,
		PurchaseOrderID: req.PurchaseID,
		SupplierID:      req.SupplierID,
		WarehouseID:     req.WarehouseID,
		UserID:          req.UserID,
		Items:           req.Items,
		Biling:          req.Biling,
		IsValidated:     req.IsValidated,
		Note:            req.Note,
		Date:            date,
	}
	err = h.PurchaseService.AddPurchaseReturnOrder(&order)
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

func (h *PurchaseHandler) InvalidatePurchaseReturnOrder(ctx *gin.Context) {
	err := h.PurchaseService.InvalidatePurchaseReturnOrder(ctx.Param("id"))
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
