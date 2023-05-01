package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	repository "github.com/mrokoo/goERP/internal/goods/unit/infra"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type UnitHandler struct {
	UnitService UnitService
}

func NewUnitHandler(unitService UnitService) *UnitHandler {
	return &UnitHandler{
		UnitService: unitService,
	}
}

func (h *UnitHandler) GetUnitList(ctx *gin.Context) {
	categories, err := h.UnitService.GetUnitList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    categories,
	})
}

func (h *UnitHandler) GetUnit(ctx *gin.Context) {
	id := ctx.Param("id")

	unit, err := h.UnitService.GetUnit(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Unit not found with the given id",
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
		Data:    unit,
	})
}

func (h *UnitHandler) AddUnit(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Note string `json:"note" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	unit := domain.Unit{
		ID:   uuid.New().String(),
		Name: req.Name,
		Note: req.Note,
	}
	err := h.UnitService.AddUnit(&unit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, reponse.Reponse{
		Data: unit,
	})
}

func (h *UnitHandler) ReplaceUnit(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Name string `json:"name" binding:"required"`
		Note string `json:"note" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	unit := domain.Unit{
		ID:   id,
		Name: req.Name,
		Note: req.Note,
	}
	err := h.UnitService.ReplaceUnit(&unit)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusBadRequest, reponse.Reponse{
				Message: "Budget not found with the given id",
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

func (h *UnitHandler) DeleteUnit(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.UnitService.DeleteUnit(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Unit not found with the given id",
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
