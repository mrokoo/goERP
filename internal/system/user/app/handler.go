package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
	repository "github.com/mrokoo/goERP/internal/system/user/infra"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUserList(ctx *gin.Context) {
	users, err := h.UserService.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    users,
	})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := uuid.MustParse(id)
	user, err := h.UserService.GetUser(uid)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "User not found with the given id",
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
		Data:    user,
	})
}

func (h *UserHandler) AddUser(ctx *gin.Context) {
	var req struct {
		// ID     string `json:"id" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Phone  string `json:"phone" binding:"-"`
		Email  string `json:"email" binding:"-"`
		Gender string `json:"gender" binding:"-"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	user := domain.User{
		ID:     id,
		Name:   req.Name,
		Phone:  req.Phone,
		Email:  req.Email,
		Gender: req.Gender,
	}
	err = h.UserService.AddUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, reponse.Reponse{
		Data: user,
	})
}

func (h *UserHandler) ReplaceUser(ctx *gin.Context) {

	var req struct {
		// ID     string `json:"id" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Phone  string `json:"phone" binding:"-"`
		Email  string `json:"email" binding:"-"`
		Gender string `json:"gender" binding:"-"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, reponse.Reponse{
			Message: "Request parameter verification failed",
		})
		return
	}
	id := ctx.Param("id")
	user := domain.User{
		ID:     uuid.MustParse(id),
		Name:   req.Name,
		Phone:  req.Phone,
		Email:  req.Email,
		Gender: req.Gender,
	}
	err := h.UserService.ReplaceUser(&user)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusBadRequest, reponse.Reponse{
				Message: "User not found with the given id",
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

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid := uuid.MustParse(id)
	if err := h.UserService.DeleteUser(uid); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, reponse.Reponse{
				Message: "User not found with the given id",
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
