package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/info"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/price"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj/stock"
	repository "github.com/mrokoo/goERP/internal/goods/product/infra"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type ProductHandler struct {
	productService ProductService
}

func NewProductHandler(s ProductService) *ProductHandler {
	return &ProductHandler{
		productService: s,
	}
}

func (h *ProductHandler) GetProductList(ctx *gin.Context) {
	products, err := h.productService.GetProductList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    products,
	})
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.productService.GetProduct(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Product not found with the given id",
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
		Data:    product,
	})
}

func (h *ProductHandler) AddProduct(ctx *gin.Context) {
	var req struct {
		ID           string    `json:"id" binding:"required"`
		Name         string    `json:"name" binding:"required"`
		CategoryID   uuid.UUID `json:"category_id" binding:"-"`
		UnitID       uuid.UUID `json:"unit_id" binding:"-"`
		OpeningStock []stock.Stock
		State        state.State `json:"state" binding:"oneof=active freeze"`
		Note         string      `json:"note" binding:"-"`
		price.Price
		info.Info
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	product := domain.Product(req)
	err := h.productService.AddProduct(&product)
	if err != nil {
		if err == ErrProductInVaildated {
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
		Data: product,
	})
}

func (h *ProductHandler) ReplaceProduct(ctx *gin.Context) {
	var req struct {
		ID           string    `json:"id" binding:"required"`
		Name         string    `json:"name" binding:"required"`
		CategoryID   uuid.UUID `json:"category_id" binding:"-"`
		UnitID       uuid.UUID `json:"unit_id" binding:"-"`
		OpeningStock []stock.Stock
		State        state.State `json:"state" binding:"oneof=active freeze"`
		Note         string      `json:"note" binding:"-"`
		price.Price
		info.Info
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	product := domain.Product(req)

	err := h.productService.ReplaceProduct(&product)
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

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.productService.DeleteProduct(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Product not found with the given id",
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
