package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/product/category/app"
)

func ProductRoutes(r *gin.Engine, categoryHandler app.CategoryHandler) {
	r.Group("category")
	{
		r.GET("/getCategoryList", categoryHandler.GetAllCategories)
		r.POST("/addCategory", categoryHandler.CreateCategory)
		r.DELETE("/deleteCategory", categoryHandler.DeleteCategory)
		r.PUT("/updateCategory", categoryHandler.ChangeCategory)
	}
}
