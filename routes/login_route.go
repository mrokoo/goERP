package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	repository "github.com/mrokoo/goERP/internal/system/user/infra"
	"github.com/mrokoo/goERP/pkg/jwta"
	"github.com/mrokoo/goERP/pkg/reponse"

	"gorm.io/gorm"
)

func NewLoginRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewUserRepository(db)

	group.POST("/login", func(c *gin.Context) {
		var req struct {
			ID       string `json:"id" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, reponse.Reponse{
				Message: "Request parameter verification failed",
			})
			return
		}

		user, err := m.GetByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, reponse.Reponse{
					Message: "User not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, reponse.Reponse{
				Message: err.Error(),
			})
			return
		}
		if user.Password != req.Password {
			c.JSON(http.StatusUnauthorized, reponse.Reponse{
				Message: "Incorrect password",
			})
			return
		}

		token, err := jwta.GenerateToken(jwta.UserInfo{
			ID:     user.ID,
			Name:   user.Name,
			Gender: user.Gender,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, reponse.Reponse{
				Message: err.Error(),
			})
		}
		c.JSON(http.StatusOK, reponse.Reponse{
			Message: "",
			Data: map[string]string{
				token: token,
			},
		})
	})
}
