package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
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
		ID           string  `json:"id" binding:"required"`
		Name         string  `json:"name" binding:"required"`
		CategoryID   *string `json:"category_id" binding:"-"`
		UnitID       *string `json:"unit_id" binding:"-"`
		OpeningStock []stock.Stock
		State        state.State `json:"state" binding:"oneof=active freeze"`
		Note         string      `json:"note" binding:"-"`
		Img          string      `json:"img" binding:"-"`
		Intro        string      `json:"intro" binding:"-"`
		Purchase     float64     `json:"purchase" binding:"-"`
		Retail       float64     `json:"retail" binding:"-"`
		Grade1       float64     `json:"grade1" binding:"-"`
		Grade2       float64     `json:"grade2" binding:"-"`
		Grade3       float64     `json:"grade3" binding:"-"`
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
	id := ctx.Param("id")
	var req struct {
		Name         string  `json:"name" binding:"required"`
		CategoryID   *string `json:"category_id" binding:"-"`
		UnitID       *string `json:"unit_id" binding:"-"`
		OpeningStock []stock.Stock
		State        state.State `json:"state" binding:"oneof=active freeze"`
		Note         string      `json:"note" binding:"-"`
		Img          string      `json:"img" binding:"-"`
		Intro        string      `json:"intro" binding:"-"`
		Purchase     float64     `json:"purchase" binding:"-"`
		Retail       float64     `json:"retail" binding:"-"`
		Grade1       float64     `json:"grade1" binding:"-"`
		Grade2       float64     `json:"grade2" binding:"-"`
		Grade3       float64     `json:"grade3" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	product := domain.Product{
		ID:           id,
		Name:         req.Name,
		CategoryID:   req.CategoryID,
		UnitID:       req.UnitID,
		OpeningStock: req.OpeningStock,
		State:        req.State,
		Note:         req.Note,
		Img:          req.Img,
		Intro:        req.Intro,
		Purchase:     req.Purchase,
		Retail:       req.Retail,
		Grade1:       req.Grade1,
		Grade2:       req.Grade2,
		Grade3:       req.Grade3,
	}

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
