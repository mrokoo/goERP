//go:generate mockgen -destination=./mock/mock_article_repository.go -package=mock github.com/rectcircle/go-test-demo/02-mock/domain ArticleRepository
package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Account struct {
	ID      string      `json:"id" gorm:"primaryKey;<-:create"`
	Name    string      `json:"name" gorm:"not null"`
	Type    PayType     `json:"type" gorm:"default:other"`
	Holder  string      `json:"holder"`
	Number  string      `json:"number"`
	Note    string      `json:"note"`
	State   state.State `json:"state" gorm:"default:active"`
	Balance float64     `json:"balance" gorm:"default:0"`
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
