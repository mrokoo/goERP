package warehouse

import (
	"context"

	"github.com/gin-gonic/gin"
)

type WarehouseApplicationService struct {
	repo Repository
}

func (w WarehouseApplicationService) UpdateWarehouse(ctx *gin.Context) {
	var err error
	var warehouse Warehouse
	var req struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Admin       string `json:"admin"`
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
	if warehouse, err = NewWarehouse(
		WarehouseCMD{
			ID:          req.ID,
			Name:        req.Name,
			Admin:       req.Admin,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Note:        req.Note,
			State:       req.State,
		},
	); err != nil {
		return
	}

	if err = w.repo.ChangeWarehouse(context.Background(), warehouse); err != nil {
		return
	}
}

func (w WarehouseApplicationService) AddWarehouse(ctx *gin.Context) {
	var err error
	var warehouse Warehouse
	var req struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Admin       string `json:"admin"`
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
	if warehouse, err = NewWarehouse(
		WarehouseCMD{
			ID:          req.ID,
			Name:        req.Name,
			Admin:       req.Admin,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Note:        req.Note,
			State:       req.State,
		},
	); err != nil {
		return
	}

	if err = w.repo.SaveWarehouse(context.Background(), warehouse); err != nil {
		return
	}
}

func (w WarehouseApplicationService) DeleteWarehouse(ctx *gin.Context) {
	var req struct {
		WarehouseId string `json:"id"`
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
