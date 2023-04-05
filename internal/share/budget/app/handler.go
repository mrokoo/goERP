package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	repository "github.com/mrokoo/goERP/internal/share/budget/infra"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type BudgetHandler struct {
	BudgetService BudgetService
}

func NewBudgetHandler(budgetService BudgetService) *BudgetHandler {
	return &BudgetHandler{
		BudgetService: budgetService,
	}
}

func (h *BudgetHandler) GetBudgetList(ctx *gin.Context) {
	budgets, err := h.BudgetService.GetBudgetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    budgets,
	})
}

func (h *BudgetHandler) GetBudget(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := uuid.MustParse(id)
	budget, err := h.BudgetService.GetBudget(uid)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Budget not found with the given id",
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
		Data:    budget,
	})
}

func (h *BudgetHandler) AddBudget(ctx *gin.Context) {
	var req struct {
		Type string `json:"type" binding:"oneof=out in"`
		Note string `json:"note" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	budget := domain.Budget{
		ID:   id,
		Type: domain.BudgetType(req.Type),
		Note: req.Note,
	}
	err = h.BudgetService.AddBudget(&budget)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, reponse.Reponse{
		Data: budget,
	})
}

func (h *BudgetHandler) ReplaceBudget(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Type string `json:"type" binding:"oneof=out in"`
		Note string `json:"note" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	budget := domain.Budget{
		ID:   uuid.MustParse(id),
		Type: domain.BudgetType(req.Type),
		Note: req.Note,
	}
	err := h.BudgetService.ReplaceBudget(&budget)
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

func (h *BudgetHandler) DeleteBudget(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := uuid.MustParse(id)
	if err := h.BudgetService.DeleteBudget(uid); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Budget not found with the given id",
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
