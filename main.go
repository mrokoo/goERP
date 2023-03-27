package main

import (
	"context"

	"github.com/gin-gonic/gin"
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

	routes.GoodsRoutes(router, client)
	routes.AccountRoutes(router, client)
	routes.BudgetRoutes(router, client)
	routes.CustomerRoutes(router, client)
	routes.SupplierRoutes(router, client)
	routes.WarehouseRoutes(router, client)

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
