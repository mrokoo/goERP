package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/customer/domain"
)

type CustomerHandler struct {
	CustomerService CustomerService
}

func NewCustomerHandler(customerService CustomerService) *CustomerHandler {
	return &CustomerHandler{
		CustomerService: customerService,
	}
}

func (h *CustomerHandler) GetCustomerList(ctx *gin.Context) {
	customers, err := h.CustomerService.GetCustomerList()
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
		"data":     customers,
	})
}

func (h *CustomerHandler) AddCustomer(ctx *gin.Context) {
	var req domain.Customer
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.CustomerService.AddCustomer(req)
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
		"data":     nil,
	})
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	var req domain.Customer
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.CustomerService.UpdateCustomer(req)
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
		"data":     nil,
	})
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
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

	if err := h.CustomerService.DeleteCustomer(req.ID); err != nil {
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
