package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/customer"
)

func main() {
	router := gin.Default()
	customer.LoadCustomerRouter(router)
	// supplier.LoadSupplierRouter(router)
	router.Run()
}
