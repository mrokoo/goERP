package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/system/user/app"
	repository "github.com/mrokoo/goERP/internal/system/user/infra"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewUserRepository(db)
	s := app.NewUserServiceImpl(m)
	h := app.NewUserHandler(s)
	group.GET("/users", h.GetUserList)
	group.GET("/users/:id", h.GetUser)
	group.POST("/users", h.AddUser)
	group.PUT("/users/:id", h.ReplaceUser)
	group.PATCH("/users/:id", h.ReplaceUser)
	group.DELETE("/users/:id", h.DeleteUser)
}
