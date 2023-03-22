package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
)

type ProductHandler struct {
	productService ProductService
}

func (h *ProductHandler) GetAllproducts(ctx *gin.Context) {
	products, err := h.productService.GetAllProducts()
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
		"data":     products,
	})
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req domain.Product
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	c, err := h.productService.CreateProduct(&req)
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

// change的检验有点小多，需要慎重。（其实create中的检验可以直接应用到change中，继续看看吧一会再写
// func (h *ProductHandler) ChangeProduct(ctx *gin.Context) {
// 	var req struct {
// 		ID   string `json:"id"`
// 		Name string `json:"name"`
// 		Note string `json:"note"`
// 	}
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		ctx.JSON(400, gin.H{
// 			"code":     -1,
// 			"showMsg":  "failure",
// 			"errorMsg": err.Error(),
// 			"data":     nil,
// 		})
// 		return
// 	}

// 	c, err := h.productService.ChangeProduct(&id, req.Name, req.Note)
// 	if err != nil {
// 		ctx.JSON(400, gin.H{
// 			"code":     -1,
// 			"showMsg":  "failure",
// 			"errorMsg": err.Error(),
// 			"data":     nil,
// 		})
// 		return
// 	}

// 	ctx.JSON(200, gin.H{
// 		"code":     1,
// 		"showMsg":  "success",
// 		"errorMsg": "",
// 		"data":     c,
// 	})
// }
