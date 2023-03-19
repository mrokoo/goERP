package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/account"
	"github.com/mrokoo/goERP/internal/share/budget"
	"github.com/mrokoo/goERP/internal/share/customer"
	"github.com/mrokoo/goERP/internal/share/supplier"
	"github.com/mrokoo/goERP/internal/share/warehouse"
)

func main() {
	router := gin.Default()
	customer.LoadCustomerRouter(router)
	supplier.LoadSupplierRouter(router)
	warehouse.LoadWarehouseRouter(router)
	budget.LoadBudgetRouter(router)
	account.LoadAccountRouter(router)
	router.Run()
}
