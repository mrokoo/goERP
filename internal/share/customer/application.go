package customer

import (
	"context"

	"github.com/gin-gonic/gin"
)

type CustomerApplicationService struct {
	repo Respository
}

func (c CustomerApplicationService) AddCustomer(ctx *gin.Context) {
	var err error
	var customer Customer
	// 最后处理错误
	defer func() {
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":     -1,
				"showMsg":  "failure",
				"errorMsg": err.Error(),
				"data":     nil,
			})
		} else {
			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     nil,
			})
		}
	}()
	// 转换请求
	var req struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Grade       int    `json:"grade"`
		Contact     string `json:"contact"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		Note        string `json:"note"`
		State       int    `json:"state"`
	}
	if err = ctx.BindJSON(&req); err != nil {
		return
	}
	// 创建新customer
	if customer, err = NewCustomer(struct {
		ID          string
		Name        string
		Grade       int
		Contact     string
		PhoneNumber string
		Address     string
		Note        string
		State       int
	}(req)); err != nil {
		return
	}

	if err = c.repo.SaveCustomer(context.Background(), customer); err != nil {
		return
	}
}

func (c CustomerApplicationService) DeleteCustomer(ctx *gin.Context) {
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
	var err error
	var req struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Grade       int    `json:"grade"`
		Contact     string `json:"contact"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		Note        string `json:"note"`
		State       int    `json:"state"`
	}
	defer func() {
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":     -1,
				"showMsg":  "failure",
				"errorMsg": err.Error(),
				"data":     nil,
			})
		} else {
			ctx.JSON(200, gin.H{
				"code":     1,
				"showMsg":  "success",
				"errorMsg": "",
				"data":     nil,
			})
		}
	}()

	if err = ctx.BindJSON(&req); err != nil {
		return
	}
	if customer, err = NewCustomer(
		CustomerCMD{
			ID:          req.ID,
			Name:        req.Name,
			Grade:       req.Grade,
			Contact:     req.Contact,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Note:        req.Note,
			State:       req.State,
		},
	); err != nil {
		return
	}

	if err = c.repo.ChangeCustomer(context.Background(), customer); err != nil {
		return
	}
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

func NewCustomerApplicationService(repo Respository) CustomerApplicationService {
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
