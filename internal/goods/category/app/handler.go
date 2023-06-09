package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	repository "github.com/mrokoo/goERP/internal/goods/category/infra"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type CategoryHandler struct {
	CategoryService CategoryService
}

func NewCategoryHandler(categoryService CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
	}
}

func (h *CategoryHandler) GetCategoryList(ctx *gin.Context) {
	var categories []*domain.Category
	categories, err := h.CategoryService.GetCategoryList()
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

func (h *CategoryHandler) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := h.CategoryService.GetCategory(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Category not found with the given id",
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
		Data:    category,
	})
}

func (h *CategoryHandler) AddCategory(ctx *gin.Context) {
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

	category := domain.Category{
		ID:   uuid.New().String(),
		Name: req.Name,
		Note: req.Note,
	}
	err := h.CategoryService.AddCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, reponse.Reponse{
		Data: category,
	})
}

func (h *CategoryHandler) ReplaceCategory(ctx *gin.Context) {
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
	category := domain.Category{
		ID:   id,
		Name: req.Name,
		Note: req.Note,
	}
	err := h.CategoryService.ReplaceCategory(&category)
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

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.CategoryService.DeleteCategory(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Category not found with the given id",
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
