package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/pkg/reponse"
)

type SaleHandler struct {
	SaleService SaleService
}

func NewSaleHandler(saleService SaleService) *SaleHandler {
	return &SaleHandler{
		SaleService: saleService,
	}
}

func (h *SaleHandler) GetSaleOrderList(ctx *gin.Context) {
	saleOrders, err := h.SaleService.GetSaleOrderList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, reponse.Reponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, reponse.Reponse{
		Message: "",
		Data:    saleOrders,
	})
}

func (h *SaleHandler) AddSaleOrder(ctx *gin.Context) {
	// var req struct {

	// }
}
