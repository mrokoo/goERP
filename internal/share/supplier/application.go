package supplier

import (
	"context"

	"github.com/gin-gonic/gin"
)

type SupplierApplicationService struct {
	repo Repositiory
}

func (s SupplierApplicationService) UpdateSupplier(ctx *gin.Context) {
	var err error
	var supplier Supplier
	var req struct {
		ID      string  `json:"id"`
		Name    string  `json:"name"`
		Contact string  `json:"contact"`
		Email   string  `json:"email"`
		Address string  `json:"address"`
		Account string  `json:"account"`
		Bank    string  `json:"bank"`
		Note    string  `json:"note"`
		State   int     `json:"state"`
		Debt    float64 `json:"debt"`
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

	if supplier, err = NewSupplier(
		SupplierCMD{
			ID:      req.ID,
			Name:    req.Name,
			Contact: req.Contact,
			Email:   req.Email,
			Address: req.Address,
			Account: req.Account,
			Bank:    req.Bank,
			Note:    req.Note,
			State:   req.State,
			Debt:    req.Debt,
		},
	); err != nil {
		return
	}

	if err = s.repo.ChangeSupplier(context.Background(), supplier); err != nil {
		return
	}
}

func (s SupplierApplicationService) AddSupplier(ctx *gin.Context) {
	var err error
	var supplier Supplier
	var req struct {
		ID      string  `json:"id"`
		Name    string  `json:"name"`
		Contact string  `json:"contact"`
		Email   string  `json:"email"`
		Address string  `json:"address"`
		Account string  `json:"account"`
		Bank    string  `json:"bank"`
		Note    string  `json:"note"`
		State   int     `json:"state"`
		Debt    float64 `json:"debt"`
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

	if supplier, err = NewSupplier(
		SupplierCMD{
			ID:      req.ID,
			Name:    req.Name,
			Contact: req.Contact,
			Email:   req.Email,
			Address: req.Address,
			Account: req.Account,
			Bank:    req.Bank,
			Note:    req.Note,
			State:   req.State,
			Debt:    req.Debt,
		},
	); err != nil {
		return
	}

	if err = s.repo.SaveSupplier(context.Background(), supplier); err != nil {
		return
	}
}

func (s SupplierApplicationService) DeleteSupplier(ctx *gin.Context) {

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

func NewSupplierApplicationService(repo Repositiory) SupplierApplicationService {
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
