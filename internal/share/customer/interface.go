package customer

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

func CreateCustomerRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewCustomerApplicationService(repo)
	c := e.Group("/customer")
	{

		c.PUT("/updateCustomer", func(ctx *gin.Context) {
			var customer Customer
			err := ctx.BindJSON(&customer)
			if err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}
			if err := service.repo.ChangeCustomer(context.Background(), customer); err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}
			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     nil,
			})
		})
		c.GET("/getCustomerList", func(ctx *gin.Context) {
			customerList, err := service.repo.FetchAllCustomers(ctx)
			if err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}
			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     customerList,
			})
		})
		c.POST("/saveCustomer", func(ctx *gin.Context) {
			var customer Customer
			err := ctx.BindJSON(&customer)
			if err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}

			if err := service.SaveCustomer(context.Background(), customer); err != nil {
				if errors.Is(err, ErrNotUID) {
					ctx.JSON(400, gin.H{
						"code":     -1,
						"showMsg":  "failure",
						"errorMsg": err.Error(),
						"data":     nil,
					})
				}
				return
			}

			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     nil,
			})
		})
		c.DELETE("/deleteCustomer", func(ctx *gin.Context) {
			var req struct {
				CustomerId string `json:"customerId"`
			}

			if err := ctx.BindJSON(&req); err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}

			if err := service.repo.DeleteCustomer(ctx, CustomerId(req.CustomerId)); err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}
			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     nil,
			})
		})
	}
}
