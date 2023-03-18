package budget

type BudgetApplicationService struct {
	repo Repository
}

// func (b BudgetApplicationService) UpdateBudget(ctx *gin.Context) {
// 	var err error
// 	var budget Budget
// 	var req struct {
// 		ID   string `json:"id"`
// 		Type int    `json:"type"`
// 		Note string `json:"note"`
// 	}
// 	defer func() {
// 		if err != nil {
// 			ctx.JSON(400, gin.H{
// 				"code":     -1,
// 				"showMsg":  "failure",
// 				"errorMsg": err.Error(),
// 				"data":     nil,
// 			})
// 		} else {
// 			ctx.JSON(200, gin.H{
// 				"code":     1,
// 				"showMsg":  "success",
// 				"errorMsg": "",
// 				"data":     nil,
// 			})
// 		}
// 	}()

// 	if err = ctx.BindJSON(&req); err != nil {
// 		return
// 	}
// 	if budget, err = (
// 		WarehouseCMD{
// 			ID:          req.ID,
// 			Name:        req.Name,
// 			Admin:       req.Admin,
// 			PhoneNumber: req.PhoneNumber,
// 			Address:     req.Address,
// 			Note:        req.Note,
// 			State:       req.State,
// 		},
// 	); err != nil {
// 		return
// 	}

// 	if err = w.repo.ChangeWarehouse(context.Background(), warehouse); err != nil {
// 		return
// 	}
// }
