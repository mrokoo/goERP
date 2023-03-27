package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/account/domain"
)

type AccountHandler struct {
	AccountService AccountService
}

func NewAccountHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

func (h *AccountHandler) GetAccountList(ctx *gin.Context) {
	accounts, err := h.AccountService.GetAccountList()
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
		"data":     accounts,
	})
}

func (h *AccountHandler) AddAccount(ctx *gin.Context) {
	var req domain.Account
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.AccountService.AddAccount(req)
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

func (h *AccountHandler) UpdateAccount(ctx *gin.Context) {
	var req domain.Account
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	err := h.AccountService.UpdateAccount(req)
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

func (h *AccountHandler) DeleteAccount(ctx *gin.Context) {
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

	if err := h.AccountService.DeleteAccount(req.ID); err != nil {
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
