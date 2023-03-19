package account

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AccountApplicationService struct {
	repo Repository
}

func (a AccountApplicationService) AddAccount(ctx *gin.Context) {
	var account Account
	if err := ctx.ShouldBindJSON(account); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}
	if err := a.repo.Save(context.Background(), account); err != nil {
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

func (a AccountApplicationService) UpdateAccount(ctx *gin.Context) {
	var account Account
	if err := ctx.ShouldBindJSON(account); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := a.repo.Change(context.Background(), account); err != nil {
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

func (a AccountApplicationService) DeleteAccount(ctx *gin.Context) {

}

func (a AccountApplicationService) GetAccountList(ctx *gin.Context) {
	var req struct {
		AccountId string `json:"id" binding:"required"`
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

	if err := a.repo.Delete(ctx, req.AccountId); err != nil {
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

func NewAccountApplicationService(repo Repository) AccountApplicationService {
	return AccountApplicationService{repo: repo}
}

func LoadAccountRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewAccountApplicationService(repo)
	r := e.Group("/account")
	{
		r.POST("/addSupplier", service.AddAccount)
		r.PUT("/updateSupplier", service.UpdateAccount)
		r.DELETE("/deleteSupplier", service.DeleteAccount)
		r.GET("/getSupplierList", service.GetAccountList)
	}
}
