package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("v1")
	NewAccountRouter(db, publicRouter)
	NewBudgetRouter(db, publicRouter)
	NewCustomerRouter(db, publicRouter)
	NewSupplierRouter(db, publicRouter)
	NewWarehouseRouter(db, publicRouter)
}
