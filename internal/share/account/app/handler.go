package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/pkg/reponse"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/mongo"
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
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    accounts,
	})
}

func (h *AccountHandler) GetAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	account, err := h.AccountService.GetAccount(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Account not found with the given id",
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
		Data:    account,
	})
}

func (h *AccountHandler) AddAccount(ctx *gin.Context) {
	var req struct {
		ID      string `json:"id" binding:"required"`
		Name    string `json:"name" binding:"required"`
		Type    string `json:"type" binding:"oneof=cash  weipay alipay other"`
		Holder  string `json:"holder" binding:"-"`
		Number  string `json:"number" binding:"-"`
		Note    string `json:"note" binding:"-"`
		State   string `json:"state" binding:"oneof=active freeze"`
		Balance string `json:"balance" binding:"numeric"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	balance, err := decimal.NewFromString(req.Balance)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	account := domain.Account{
		ID:      req.ID,
		Name:    req.Name,
		Type:    domain.PayType(req.Type),
		Holder:  req.Holder,
		Number:  req.Number,
		Note:    req.Note,
		State:   state.State(req.State),
		Balance: balance,
	}

	err = h.AccountService.AddAccount(&account)
	if err != nil {
		if err == ErrAccountInVaildated {
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
	ctx.JSON(http.StatusCreated, reponse.Reponse{})
}

func (h *AccountHandler) ReplaceAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Name    string `json:"name" binding:"required"`
		Type    string `json:"type" binding:"oneof=cash  weipay alipay other"`
		Holder  string `json:"holder" binding:"-"`
		Number  string `json:"number" binding:"-"`
		Note    string `json:"note" binding:"-"`
		State   string `json:"state" binding:"oneof=active freeze"`
		Balance string `json:"balance" binding:"numeric"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	balance, err := decimal.NewFromString(req.Balance)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	account := domain.Account{
		ID:      id,
		Name:    req.Name,
		Type:    domain.PayType(req.Type),
		Holder:  req.Holder,
		Number:  req.Number,
		Note:    req.Note,
		State:   state.State(req.State),
		Balance: balance,
	}

	err = h.AccountService.ReplaceAccount(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
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

func (h *AccountHandler) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.AccountService.DeleteAccount(id); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Account not found with the given id",
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
