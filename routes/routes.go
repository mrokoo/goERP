package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("v1")
	NewAccountRouter(db, publicRouter)
}
