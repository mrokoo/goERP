package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
)

type UnitService interface {
	CreateUnit(name string, note string) (*domain.Unit, error)
	ChangeUnit(unitId *uuid.UUID, name string, note string) (*domain.Unit, error)
	GetAllUnits() ([]domain.Unit, error)
	DeleteUnit(unitId *uuid.UUID) error
}

type UnitHandler struct {
	unitService UnitService
}

func NewUnitHandler(s UnitService) *UnitHandler {
	return &UnitHandler{
		unitService: s,
	}
}

func (h *UnitHandler) GetAllUnits(ctx *gin.Context) {
	units, err := h.unitService.GetAllUnits()
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
		"data":     units,
	})
}

func (h *UnitHandler) CreateUnit(ctx *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Note string `json:"note"`
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

	c, err := h.unitService.CreateUnit(req.Name, req.Note)
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
		"data":     c,
	})
}

func (h *UnitHandler) DeleteUnit(ctx *gin.Context) {
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
	var id uuid.UUID
	id.UnmarshalText([]byte(req.ID))
	if err := h.unitService.DeleteUnit(&id); err != nil {
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

func (h *UnitHandler) ChangeUnit(ctx *gin.Context) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Note string `json:"note"`
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	var id uuid.UUID
	id.UnmarshalText([]byte(req.ID))
	c, err := h.unitService.ChangeUnit(&id, req.Name, req.Note)
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
		"data":     c,
	})
}
