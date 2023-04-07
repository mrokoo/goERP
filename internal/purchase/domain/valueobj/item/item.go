package item

import (
	"github.com/shopspring/decimal"
)

type PurchaseOrderItem struct {
	PurchaseOrderID string
	ProductID       string
	Quantity        int
	Price           float64
	TotalAmount     float64
}

func New(purchaseID string, productID string, quantity int, price float64) PurchaseOrderItem {
	p := decimal.NewFromFloat(price)
	q := decimal.NewFromInt(int64(quantity))
	t := p.Mul(q)
	tf, _ := t.Round(2).Float64()
	return PurchaseOrderItem{
		PurchaseOrderID: purchaseID,
		ProductID:       productID,
		Quantity:        quantity,
		Price:           price,
		TotalAmount:     tf,
	}
}
