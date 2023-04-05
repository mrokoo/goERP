package price

import "github.com/shopspring/decimal"

type Price struct {
	Purchase decimal.Decimal `json:"purchase"`
	Retail   decimal.Decimal `json:"retail"`
	Grade1   decimal.Decimal `json:"grade1"`
	Grade2   decimal.Decimal `json:"grade2"`
	Grade3   decimal.Decimal `json:"grade3"`
}
