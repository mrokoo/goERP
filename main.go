package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gin := gin.Default()
	gin.Use(cors.Default())
	dsn := "root:123456@tcp(127.0.0.1:3306)/goerp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	routes.Setup(db, gin)
	gin.Run(":8000")
}
