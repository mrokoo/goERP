//go:generate mockgen -destination=./mock/mock_article_repository.go -package=mock github.com/rectcircle/go-test-demo/02-mock/domain ArticleRepository
package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID      string
	Name    string
	Type    PayType
	Holder  string
	Number  string
	Note    string
	State   state.State
	Balance decimal.Decimal
}


type PayType string

const (
	TYPE_CASH   PayType = "cash"
	TYPE_WEIPAY PayType = "weipay"
	TYPE_ALiPAY PayType = "alipay"
	TYPE_OTHER  PayType = "other"
)

func (p *PayType) String() string {
	switch *p {
	case TYPE_CASH:
		return "cash"
	case TYPE_WEIPAY:
		return "weipay"
	case TYPE_ALiPAY:
		return "alipay"
	default:
		return "other"
	}
}
