package routes

import (
	"github.com/gin-gonic/gin"
	categoryapp "github.com/mrokoo/goERP/internal/goods/category/app"
	unitapp "github.com/mrokoo/goERP/internal/goods/unit/app"
)

func ProductRoutes(r *gin.Engine, categoryHandler categoryapp.CategoryHandler, unitHandler unitapp.UnitHandler) {
	r.Group("category")
	{
		r.GET("/getCategoryList", categoryHandler.GetAllCategories)
		r.POST("/addCategory", categoryHandler.CreateCategory)
		r.DELETE("/deleteCategory", categoryHandler.DeleteCategory)
		r.PUT("/updateCategory", categoryHandler.ChangeCategory)
	}

	r.Group("unit")
	{
		r.GET("/getUnitList", unitHandler.GetAllUnits)
		r.POST("/addUnit", unitHandler.CreateUnit)
		r.DELETE("/deleteUnit", unitHandler.DeleteUnit)
		r.PUT("/updateUnit", unitHandler.ChangeUnit)
	}
}
