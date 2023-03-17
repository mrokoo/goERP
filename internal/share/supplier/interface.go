package supplier

import (
	"context"
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
)

func CreateSupplierRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewSupplierApplicationService(repo)
	c := e.Group("/supplier")
	{
		c.PUT("/updateSupplier", func(ctx *gin.Context) {
			// to do 更改参数类型
			var req struct {
				ID      SupplierId  `json:"id"`
				Name    Name        `json:"name"`
				Contact ContactName `json:"contact"`
				Email   Email       `json:"email"`
				Address Address     `json:"address"`
				Account Account     `json:"account"`
				Bank    BankName    `json:"bank"`
				Note    string      `json:"note"`
				State   StateType   `json:"state"`
				Debt    money.Money `json:"debt"`
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
			if err := service.UpdateSupplier(context.Background(), req); err != nil {
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
		c.GET("/getSupplierList", func(ctx *gin.Context) {
			supplierList, err := service.GetSupplierList(context.Background())
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
				"data":     supplierList,
			})
		})
		c.POST("/saveSupplier", func(ctx *gin.Context) {
			var supplier Supplier
			err := ctx.BindJSON(&supplier)
			if err != nil {
				ctx.JSON(400, gin.H{
					"code":     -1,
					"showMsg":  "failure",
					"errorMsg": err.Error(),
					"data":     nil,
				})
				return
			}

			if err := service.SaveSupplier(context.Background(), supplier); err != nil {
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
		c.DELETE("/deleteSupplier", func(ctx *gin.Context) {
			var req struct {
				SupplierId string `json:"supplierId"`
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

			if err := service.DeleteSupplier(ctx, req.SupplierId); err != nil {
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
