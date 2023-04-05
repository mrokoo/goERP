package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
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
		if err == mongo.ErrNoDocuments {
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

func (h *ProductHandler) AddProduct(ctx *gin.Context) {
	// var req struct {
	// 	ID
	// }
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
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

	if err := h.productService.DeleteProduct(req.ID); err != nil {
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

func (h *ProductHandler) ChangeProduct(ctx *gin.Context) {
	var req domain.Product
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.productService.UpdateProduct(&req)
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
		"data":     req,
	})
}

