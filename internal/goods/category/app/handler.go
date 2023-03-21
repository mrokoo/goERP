package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/category/domain"
)

type CategoryService interface {
	CreateCategory(name string, note string) (*domain.Category, error)
	ChangeCategory(categoryId *uuid.UUID, name string, note string) (*domain.Category, error)
	GetAllCategories() ([]domain.Category, error)
	DeleteCategory(categoryId *uuid.UUID) error
}

type CategoryHandler struct {
	categoryService CategoryService
}

func (h *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
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
		"data":     categories,
	})
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
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

	c, err := h.categoryService.CreateCategory(req.Name, req.Note)
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

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
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
	if err := h.categoryService.DeleteCategory(&id); err != nil {
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

func (h *CategoryHandler) ChangeCategory(ctx *gin.Context) {
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
	c, err := h.categoryService.ChangeCategory(&id, req.Name, req.Note)
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
