package warehouse

import (
	"context"

	"github.com/gin-gonic/gin"
)

type WarehouseApplicationService struct {
	repo Repository
}

func (w WarehouseApplicationService) UpdateWarehouse(ctx *gin.Context) {
	var warehouse Warehouse
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := w.repo.ChangeWarehouse(context.Background(), warehouse); err != nil {
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

func (w WarehouseApplicationService) AddWarehouse(ctx *gin.Context) {
	var warehouse Warehouse
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := w.repo.SaveWarehouse(context.Background(), warehouse); err != nil {
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

func (w WarehouseApplicationService) DeleteWarehouse(ctx *gin.Context) {
	var req struct {
		WarehouseId string `json:"id" binding:"required"`
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

	if err := w.repo.DeleteWarehouse(ctx, WarehouseId(req.WarehouseId)); err != nil {
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

func (w WarehouseApplicationService) GetWarehouseList(ctx *gin.Context) {
	warehouseList, err := w.repo.FetchAllWarehouse(context.Background())
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
		"data":     warehouseList,
	})
}

func NewWarehouseApplicationService(repo Repository) WarehouseApplicationService {
	return WarehouseApplicationService{repo: repo}
}

func LoadWarehouseRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewWarehouseApplicationService(repo)
	r := e.Group("/warehouse")
	{
		r.POST("/addWarehouse", service.AddWarehouse)
		r.PUT("/updateWarehouse", service.UpdateWarehouse)
		r.DELETE("/deleteWarehouse", service.DeleteWarehouse)
		r.GET("/getWarehouseList", service.GetWarehouseList)
	}
}
