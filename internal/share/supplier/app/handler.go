package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	repository "github.com/mrokoo/goERP/internal/share/supplier/infra"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type SupplierHandler struct {
	SupplierService SupplierService
}

func NewSupplierHandler(supplierService SupplierService) *SupplierHandler {
	return &SupplierHandler{
		SupplierService: supplierService,
	}
}

func (h *SupplierHandler) GetSupplierList(ctx *gin.Context) {
	suppliers, err := h.SupplierService.GetSupplierList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    suppliers,
	})
}

func (h *SupplierHandler) GetSupplier(ctx *gin.Context) {
	id := ctx.Param("id")
	spplier, err := h.SupplierService.GetSupplier(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Supplier not found with the given id",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    spplier,
	})
}

func (h *SupplierHandler) AddSupplier(ctx *gin.Context) {
	var req struct {
		ID      string      `json:"id" binding:"required"`
		Name    string      `json:"name" binding:"required"`
		Contact string      `json:"contact" binding:"-"`
		Email   string      `json:"email" binding:"-"`
		Address string      `json:"address" binding:"-"`
		Account string      `json:"account" binding:"-"`
		Bank    string      `json:"bank" binding:"-"`
		Note    string      `json:"note" binding:"-"`
		State   state.State `json:"state" binding:"oneof=active freeze"`
		Debt    float64     `json:"debt"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	supplier := domain.Supplier{
		ID:      req.ID,
		Name:    req.Name,
		Contact: req.Contact,
		Email:   req.Email,
		Address: req.Address,
		Account: req.Account,
		Bank:    req.Bank,
		Note:    req.Note,
		State:   req.State,
		Debt:    req.Debt,
	}
	err := h.SupplierService.AddSupplier(&supplier)
	if err != nil {
		if err == ErrSupplierInVaildated {
			ctx.JSON(http.StatusBadRequest, reponse.Reponse{
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, reponse.Reponse{
		Data: supplier,
	})
}

func (h *SupplierHandler) ReplaceSupplier(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Name    string      `json:"name" binding:"required"`
		Contact string      `json:"contact" binding:"-"`
		Email   string      `json:"email" binding:"-"`
		Address string      `json:"address" binding:"-"`
		Account string      `json:"account" binding:"-"`
		Bank    string      `json:"bank" binding:"-"`
		Note    string      `json:"note" binding:"-"`
		State   state.State `json:"state" binding:"oneof=active freeze"`
		Debt    float64     `json:"debt"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	supplier := domain.Supplier{
		ID:      id,
		Name:    req.Name,
		Contact: req.Contact,
		Email:   req.Email,
		Address: req.Address,
		Account: req.Account,
		Bank:    req.Bank,
		Note:    req.Note,
		State:   req.State,
		Debt:    req.Debt,
	}
	err := h.SupplierService.ReplaceSupplier(&supplier)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusBadRequest, reponse.Reponse{
				Message: "Account not found with the given id",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, reponse.Reponse{})
}

func (h *SupplierHandler) DeleteSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.SupplierService.DeleteSupplier(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Supplier not found with the given id",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, reponse.Reponse{})
}
