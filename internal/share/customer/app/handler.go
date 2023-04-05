package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/customer/domain"
	repository "github.com/mrokoo/goERP/internal/share/customer/infra"
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/mrokoo/goERP/pkg/reponse"
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
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    customers,
	})
}

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := h.CustomerService.GetCustomer(id)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Customer not found with the given id",
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
		Data:    customer,
	})
}

func (h *CustomerHandler) AddCustomer(ctx *gin.Context) {
	var req struct {
		ID      string  `json:"id" binding:"required"`
		Name    string  `json:"name" binding:"required"`
		Grade   string  `json:"grade" binding:"oneof=high medium low"`
		Contact string  `json:"contact" binding:"-"`
		Phone   string  `json:"phone" binding:"-"`
		Email   string  `json:"email" binding:"-"`
		Address string  `json:"address" binding:"-"`
		Note    string  `json:"note" binding:"-"`
		State   string  `json:"state" binding:"oneof=active freeze"`
		Debt    float64 `json:"debt"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	customer := domain.Customer{
		ID:      req.ID,
		Name:    req.Name,
		Grade:   domain.GradeType(req.Grade),
		Contact: req.Contact,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		Note:    req.Note,
		State:   state.State(req.State),
		Debt:    req.Debt,
	}
	err := h.CustomerService.AddCustomer(&customer)
	if err != nil {
		if err == ErrCustomerInVaildated {
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
		Data: req,
	})
}

func (h *CustomerHandler) ReplaceCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		Name    string  `json:"name" binding:"required"`
		Grade   string  `json:"grade" binding:"oneof=high medium low"`
		Contact string  `json:"contact" binding:"-"`
		Phone   string  `json:"phone" binding:"-"`
		Email   string  `json:"email" binding:"-"`
		Address string  `json:"address" binding:"-"`
		Note    string  `json:"note" binding:"-"`
		State   string  `json:"state" binding:"oneof=active freeze"`
		Debt    float64 `json:"debt" binding:"numeric"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	customer := domain.Customer{
		ID:      id,
		Name:    req.Name,
		Grade:   domain.GradeType(req.Grade),
		Contact: req.Contact,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		Note:    req.Note,
		State:   state.State(req.State),
		Debt:    req.Debt,
	}

	err := h.CustomerService.ReplaceCustomer(&customer)
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

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.CustomerService.DeleteCustomer(id); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "Customer not found with the given id",
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
