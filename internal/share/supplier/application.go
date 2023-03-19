package supplier

import (
	"context"

	"github.com/gin-gonic/gin"
)

type SupplierApplicationService struct {
	repo Repository
}

func (s SupplierApplicationService) UpdateSupplier(ctx *gin.Context) {
	var supplier Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	if err := s.repo.ChangeSupplier(context.Background(), supplier); err != nil {
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

func (s SupplierApplicationService) AddSupplier(ctx *gin.Context) {
	var supplier Supplier
	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	if err := s.repo.SaveSupplier(context.Background(), supplier); err != nil {
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

func (s SupplierApplicationService) DeleteSupplier(ctx *gin.Context) {
	var req struct {
		SupplierId string `json:"id"`
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

	if err := s.repo.DeleteSupplier(ctx, req.SupplierId); err != nil {
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

func (s SupplierApplicationService) GetSupplierList(ctx *gin.Context) {
	supplierList, err := s.repo.FetchAllSuppliers(context.Background())
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
}

func NewSupplierApplicationService(repo Repository) SupplierApplicationService {
	return SupplierApplicationService{repo: repo}
}

func LoadSupplierRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewSupplierApplicationService(repo)
	r := e.Group("/supplier")
	{
		r.POST("/addSupplier", service.AddSupplier)
		r.PUT("/updateSupplier", service.UpdateSupplier)
		r.DELETE("/deleteSupplier", service.DeleteSupplier)
		r.GET("/getSupplierList", service.GetSupplierList)
	}
}
