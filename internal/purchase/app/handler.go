package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/purchase/domain"
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
		ID           string        `json:"id" binding:"required"`
		SupplierID   string        `json:"supplier_id" binding:"required"`
		WarehouseID  string        `json:"warehouse_id" binding:"required"`
		UserID       string        `json:"user_id" binding:"required"`
		Items        []domain.Item `json:"items" binding:"required"`
		AccountID    string        `json:"account_id" binding:"required"`
		OtherCost    float64       `json:"other_cost" binding:"gte=0"`
		ActalPayment float64       `json:"actal_payment" binding:"gte=0"`
		Basic        string        `json:"basic"`
		Kind         string        `json:"kind" binding:"required,oneof=Order ReturnOrder"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := domain.NewPurchaseOrder(req.ID, req.WarehouseID, req.SupplierID, req.UserID, req.AccountID, req.OtherCost, req.ActalPayment, req.Basic, req.Items, domain.Kind(req.Kind))
	err := h.PurchaseService.AddPurchaseOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    order,
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

func (h *PurchaseHandler) AddPurchaseReturnOrder(ctx *gin.Context) {
	var req struct {
		ID           string        `json:"id" binding:"required"`
		SupplierID   string        `json:"supplier_id" binding:"required"`
		WarehouseID  string        `json:"warehouse_id" binding:"required"`
		UserID       string        `json:"user_id" binding:"required"`
		Items        []domain.Item `json:"items" binding:"required"`
		AccountID    string        `json:"account_id" binding:"required"`
		OtherCost    float64       `json:"other_cost" binding:"gte=0"`
		ActalPayment float64       `json:"actal_payment" binding:"gte=0"`
		Basic        string        `json:"basic"`
		Kind         string        `json:"kind" binding:"required,oneof=Order ReturnOrder"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := domain.NewPurchaseOrder(req.ID, req.WarehouseID, req.SupplierID, req.UserID, req.AccountID, req.OtherCost, req.ActalPayment, req.Basic, req.Items, domain.Kind(req.Kind))
	err := h.PurchaseService.AddPurchaseReturnOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    order,
	})
}

func (h *PurchaseHandler) InvalidatePurchaseReturnOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.PurchaseService.InvalidatePurchaseReturnOrder(id)
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
