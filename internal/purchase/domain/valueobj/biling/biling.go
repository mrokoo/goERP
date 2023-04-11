package biling

// 订单的账单信息：务必使用New
type Biling struct {
	AccountID  string  // 账户ID
	OtherCosts float64 // 其他费用
	TotalCosts float64 // 总费用
	Amount     float64 // 总收付款或收款
	Debt       float64 // 欠款（debt = totalCosts - amount）
}

func New(accountID string, othercosts float64, cost float64, amount float64) Biling {
	totalCosts := othercosts + cost
	return Biling{
		AccountID:  accountID,
		OtherCosts: othercosts,
		TotalCosts: totalCosts,
		Amount:     amount,
		Debt:       totalCosts - amount,
	}
}
