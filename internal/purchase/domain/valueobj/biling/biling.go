package biling

type Biling struct {
	AccountID  string
	OtherCosts float64
	TotalCosts float64
	AmountPaid float64
	Debt       float64
}

func New(accountID string, othercosts float64, cost float64, amountPaid float64) Biling {
	totalCosts := othercosts + cost
	return Biling{
		AccountID:  accountID,
		OtherCosts: othercosts,
		TotalCosts: totalCosts,
		AmountPaid: amountPaid,
		Debt:       totalCosts - amountPaid,
	}
}
