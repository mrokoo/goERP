package routes

import (
	"github.com/gin-gonic/gin"
	categoryapp "github.com/mrokoo/goERP/internal/goods/category/app"
)

func GoodsRoutes(r *gin.Engine, h *categoryapp.CategoryHandler) {
	category := r.Group("/category")
	{
		category.GET("/getCategoryList", h.GetAllCategories)
		category.POST("/addCategory", h.CreateCategory)
		category.DELETE("/deleteCategory", h.DeleteCategory)
		category.PUT("/updateCategory", h.ChangeCategory)
	}

	// r.Group("unit")
	// {
	// 	r.GET("/getUnitList", unitHandler.GetAllUnits)
	// 	r.POST("/addUnit", unitHandler.CreateUnit)
	// 	r.DELETE("/deleteUnit", unitHandler.DeleteUnit)
	// 	r.PUT("/updateUnit", unitHandler.ChangeUnit)
	// }
}
