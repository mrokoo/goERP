package routes

import (
	"github.com/gin-gonic/gin"
	app "github.com/mrokoo/goERP/internal/goods/category/app"
	repository "github.com/mrokoo/goERP/internal/goods/category/infra"
	"gorm.io/gorm"
)

func NewCategoryRouter(db *gorm.DB, group *gin.RouterGroup) {
	m := repository.NewCategoryRepository(db)
	s := app.NewCategoryServiceImpl(m)
	h := app.NewCategoryHandler(s)
	group.GET("/categories", h.GetCategoryList)
	group.GET("/categories/:id", h.GetCategory)
	group.POST("/categories", h.AddCategory)
	group.PUT("/categories/:id", h.ReplaceCategory)
	group.PATCH("/categories/:id", h.ReplaceCategory)
	group.DELETE("/categories/:id", h.DeleteCategory)
}
