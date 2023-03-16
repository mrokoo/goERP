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
	app := NewCustomerApplicationService(repo)
	c := e.Group("/customer")
	{
		c.POST("/change", func(ctx *gin.Context) {
			var customer Customer
			err := ctx.BindJSON(&customer)
			if err != nil {
				// to do
			}
			app.repo.ChangeCustomer(context.Background(), customer)
			// to do
		})
		c.POST("/fetchAllcustomer", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "helloworld",
			})
		})

		c.POST("/saveCustomer", func(ctx *gin.Context) {
			var customer Customer
			err := ctx.BindJSON(&customer)
			if err != nil {
				ctx.JSON(400, gin.H{
					"message": "错误",
				})
			}

			if err := app.SaveCustomer(context.Background(), customer); err != nil {
				if errors.Is(err, ErrNotUID) {
					ctx.JSON(300, gin.H{
						"message": "不是唯一ID",
					})
				}

			} else {
				ctx.JSON(200, gin.H{
					"message": "添加OK",
				})
			}
		})
		// to do
	}
}
