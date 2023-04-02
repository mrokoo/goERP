package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	categoryapp "github.com/mrokoo/goERP/internal/goods/category/app"
// 	categoryinfra "github.com/mrokoo/goERP/internal/goods/category/infra/mongodb"
// 	productapp "github.com/mrokoo/goERP/internal/goods/product/app"
// 	productdomain "github.com/mrokoo/goERP/internal/goods/product/domain"
// 	productinfra "github.com/mrokoo/goERP/internal/goods/product/infra/mongodb"
// 	unitapp "github.com/mrokoo/goERP/internal/goods/unit/app"
// 	unitinfra "github.com/mrokoo/goERP/internal/goods/unit/infra/mongodb"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func GoodsRoutes(r *gin.Engine, client *mongo.Client) {

// 	category := r.Group("/category")
// 	{
// 		m := categoryinfra.NewMongoCategoryRepository(client)
// 		s := categoryapp.NewCategoryServiceImpl(m)
// 		h := categoryapp.NewCategoryHandler(s)
// 		category.GET("/getCategoryList", h.GetAllCategories)
// 		category.POST("/addCategory", h.CreateCategory)
// 		category.DELETE("/deleteCategory", h.DeleteCategory)
// 		category.PUT("/updateCategory", h.ChangeCategory)
// 	}

// 	unit := r.Group("/unit")
// 	{
// 		m := unitinfra.NewMongoUnitRepository(client)
// 		s := unitapp.NewUnitServiceImpl(m)
// 		h := unitapp.NewUnitHandler(s)
// 		unit.GET("/getUnitList", h.GetAllUnits)
// 		unit.POST("/addUnit", h.CreateUnit)
// 		unit.DELETE("/deleteUnit", h.DeleteUnit)
// 		unit.PUT("/updateUnit", h.ChangeUnit)
// 	}

// 	product := r.Group("/product")
// 	{
// 		m := productinfra.NewMongoProductRepository(client)
// 		ds := productdomain.NewCheckingProductValidityService(m)
// 		s := productapp.NewProductServiceImpl(m, *ds)
// 		h := productapp.NewProductHandler(s)
// 		product.GET("/getProductList", h.GetAllProducts)
// 		product.POST("/addProduct", h.CreateProduct)
// 		product.DELETE("/deleteProduct", h.DeleteProduct)
// 		product.PUT("/updateProduct", h.ChangeProduct)
// 	}
// }
