package budget

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BudgetApplicationService struct {
	repo Repository
}

func (b BudgetApplicationService) UpdateBudget(ctx *gin.Context) {
	var err error
	var budget Budget
	var req struct {
		ID   string `json:"id"`
		Type int    `json:"type"`
		Note string `json:"note"`
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

	id := []byte(req.ID)
	budget, err = b.repo.LoadBudget(context.Background(), uuid.UUID(id))
	if err != nil {
		return
	}

	budget.Type, err = NewType(req.Type)
	if err != nil {
		return
	}
	budget.Note = req.Note
	err = b.repo.SaveBudget(context.Background(), budget)
	if err != nil {
		return
	}
}

func (b BudgetApplicationService) AddBudget(ctx *gin.Context) {
	var err error
	var budget Budget
	var req struct {
		Type int    `json:"type"`
		Note string `json:"note"`
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

	budget, err = NewBudget(BudgetCMD{
		Note: req.Note,
		Type: req.Type,
	})
	if err != nil {
		return
	}

	err = b.repo.SaveBudget(context.Background(), budget)
	if err != nil {
		return
	}
}

func (b BudgetApplicationService) DeleteBudget(ctx *gin.Context) {
	var req struct {
		ID string `json:"id"`
	}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"code":     -1,
			"showMsg":  "failure",
			"errorMsg": err.Error(),
			"data":     nil,
		})
		return
	}

	if err := b.repo.DeleteBudget(context.Background(), uuid.UUID([]byte(req.ID))); err != nil {
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

func (b BudgetApplicationService) GetBudgetList(ctx *gin.Context) {
	budgetList, err := b.repo.FetchAllBudget(context.Background())
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
		"data":     budgetList,
	})
}

func NewBudgetApplicationService(repo Repository) BudgetApplicationService {
	return BudgetApplicationService{
		repo: repo,
	}
}
func LoadCustomerRouter(e *gin.Engine) {
	mongoConString := "mongodb://localhost:27017/"
	repo, err := NewMongoRepo(context.Background(), mongoConString)
	if err != nil {
		panic(err)
	}
	service := NewBudgetApplicationService(repo)
	r := e.Group("/budget")
	{
		r.POST("/addBudget", service.AddBudget)
		r.GET("/getBudgetList", service.GetBudgetList)
		r.DELETE("/deleteBudget", service.DeleteBudget)
		r.PUT("/updateBudget", service.UpdateBudget)
	}
}
