package item

import (
	"github.com/shopspring/decimal"
)

// 订单项：务必使用New创建
type OrderItem struct {
	ProductID    string  // 产品ID
	Quantity     int     // 产品数量
	Price        float64 // 产品价格
	TotalPayment float64 // 产品付款 = 产品数量 * 产品价格
}

func NewOrderItem(productID string, quantity int, price float64) OrderItem {
	p := decimal.NewFromFloat(price)
	q := decimal.NewFromInt(int64(quantity))
	t := p.Mul(q)
	tf, _ := t.Round(2).Float64()
	return OrderItem{
		ProductID:    productID,
		Quantity:     quantity,
		Price:        price,
		TotalPayment: tf,
	}
}

type ReturnOrderItem struct {
	ProductID       string  // 产品ID
	Quantity        int     // 产品数量
	Price           float64 // 产品价格
	TotalCollection float64 // 产品收款 = 产品数量 * 产品价格
}

func NewReturnOrderItem(productID string, quantity int, price float64) ReturnOrderItem {
	p := decimal.NewFromFloat(price)
	q := decimal.NewFromInt(int64(quantity))
	t := p.Mul(q)
	tf, _ := t.Round(2).Float64()
	return ReturnOrderItem{
		ProductID:       productID,
		Quantity:        quantity,
		Price:           price,
		TotalCollection: tf,
	}
}
