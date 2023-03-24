package main

import (
	"context"

	"github.com/gin-gonic/gin"
	categoryapp "github.com/mrokoo/goERP/internal/goods/category/app"
	categoryinfra "github.com/mrokoo/goERP/internal/goods/category/infra/mongodb"
	"github.com/mrokoo/goERP/internal/share/budget"
	"github.com/mrokoo/goERP/internal/share/customer"
	"github.com/mrokoo/goERP/internal/share/supplier"
	"github.com/mrokoo/goERP/internal/share/warehouse"
	"github.com/mrokoo/goERP/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r := categoryinfra.NewMongoCategoryRepository(client)
	s := categoryapp.NewCategoryServiceImpl(r)
	categoryHandler := categoryapp.NewCategoryHandler(s)
	routes.GoodsRoutes(router, categoryHandler)
	customer.LoadCustomerRouter(router)
	supplier.LoadSupplierRouter(router)
	warehouse.LoadWarehouseRouter(router)
	budget.LoadBudgetRouter(router)
	account.LoadAccountRouter(router)

	return router
}
func main() {
	router := SetupRouter()
	router.Run()
}
