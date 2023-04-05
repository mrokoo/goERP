package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
	repository "github.com/mrokoo/goERP/internal/share/warehouse/infra"
	"github.com/mrokoo/goERP/pkg/reponse"
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
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    warehouses,
	})
}

func (h *WarehouseHandler) GetWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")
	warehouse, err := h.WarehouseService.GetWarehouse(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Warehouse not found with the given id",
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
		Data:    warehouse,
	})
}

func (h *WarehouseHandler) AddWarehouse(ctx *gin.Context) {
	var req struct {
		ID      string      `json:"id" binding:"required"`
		Name    string      `json:"name" binding:"required"`
		Admin   string      `json:"admin" binding:"-"`
		Phone   string      `json:"phone" binding:"-"`
		Address string      `json:"address" binding:"-"`
		Note    string      `json:"note" binding:"-"`
		State   state.State `json:"state" binding:"oneof=active freeze"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	warehouse := domain.Warehouse{
		ID:      req.ID,
		Name:    req.Name,
		Admin:   req.Admin,
		Phone:   req.Phone,
		Address: req.Address,
		Note:    req.Note,
		State:   req.State,
	}
	err := h.WarehouseService.AddWarehouse(&warehouse)
	if err != nil {
		if err == ErrWarehouseInVaildated {
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
		Data: warehouse,
	})
}

func (h *WarehouseHandler) ReplaceWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Name    string      `json:"name" binding:"required"`
		Admin   string      `json:"admin" binding:"-"`
		Phone   string      `json:"phone" binding:"-"`
		Address string      `json:"address" binding:"-"`
		Note    string      `json:"note" binding:"-"`
		State   state.State `json:"state" binding:"oneof=active freeze"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	warehouse := domain.Warehouse{
		ID:      id,
		Name:    req.Name,
		Admin:   req.Admin,
		Phone:   req.Phone,
		Address: req.Address,
		Note:    req.Note,
		State:   req.State,
	}

	err := h.WarehouseService.ReplaceWarehouse(&warehouse)
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

func (h *WarehouseHandler) DeleteWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.WarehouseService.DeleteWarehouse(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Warehouse not found with the given id",
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
