package customer

import (
	"context"

	"github.com/gin-gonic/gin"
)

type CustomerApplicationService struct {
	repo Repository
}

func (c CustomerApplicationService) AddCustomer(ctx *gin.Context) {
	var customer Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	if err := c.repo.SaveCustomer(context.Background(), customer); err != nil {
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
}

func (c CustomerApplicationService) DeleteCustomer(ctx *gin.Context) {
	var req struct {
		CustomerId string `json:"id" bson:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := c.repo.DeleteCustomer(ctx, CustomerId(req.CustomerId)); err != nil {
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
}

func (c CustomerApplicationService) UpdateCustomer(ctx *gin.Context) {
	var customer Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := c.repo.ChangeCustomer(context.Background(), customer); err != nil {
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
}

func (c CustomerApplicationService) GetCustomerList(ctx *gin.Context) {
	customerList, err := c.repo.FetchAllCustomers(context.Background())
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
}

func NewCustomerApplicationService(repo Repository) CustomerApplicationService {
	return CustomerApplicationService{
		repo: repo,
	}
}

func LoadCustomerRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewCustomerApplicationService(repo)
	r := e.Group("/customer")
	{
		r.POST("/addCustomer", service.AddCustomer)
		r.GET("/getCustomerList", service.GetCustomerList)
		r.DELETE("/deleteCustomer", service.DeleteCustomer)
		r.PUT("/updateCustomer", service.UpdateCustomer)
	}
}
