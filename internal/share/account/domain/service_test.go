//go:generate mockgen -destination=./mock/mock_account_repository.go -package=mock github.com/mrokoo/goERP/internal/share/account/domain Repository

package domain_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/mrokoo/goERP/internal/share/account/domain/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCheckingAccountValidityService(t *testing.T) {
	want := domain.Account{
		ID:      "A001",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "	402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	want2 := domain.Account{
		ID:      "A002",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "	402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)
	m.EXPECT().Get("A001").Return(&want, nil)
	m.EXPECT().Get("A002").Return(nil, mongo.ErrNoDocuments)
	ds := domain.NewCheckingAccountValidityService(m)
	t.Run("sucess", func(t *testing.T) {
		assert.False(ds.IsValidated(want))
	})

	t.Run(("failure"), func(t *testing.T) {
		assert.True(ds.IsValidated(want2))
	})
}
