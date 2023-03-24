package routes

import (
	"github.com/gin-gonic/gin"
	categoryapp "github.com/mrokoo/goERP/internal/goods/category/app"
	categoryinfra "github.com/mrokoo/goERP/internal/goods/category/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func GoodsRoutes(r *gin.Engine, client *mongo.Client) {

	category := r.Group("/category")
	{
		m := categoryinfra.NewMongoCategoryRepository(client)
		s := categoryapp.NewCategoryServiceImpl(m)
		h := categoryapp.NewCategoryHandler(s)
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
