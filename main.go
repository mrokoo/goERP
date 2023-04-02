package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	gin := gin.Default()
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	defer func() {
		client.Disconnect(context.Background())
	}()
	db := client.Database("goERP")
	if err != nil {
		panic(err)
	}
	routes.Setup(db, gin)
	gin.Run()
}
