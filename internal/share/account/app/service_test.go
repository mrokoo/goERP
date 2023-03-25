//go:generate mockgen -destination=./mock/mock_account_repository.go -package=mock github.com/mrokoo/goERP/internal/share/account/domain Repository
package app_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrokoo/goERP/internal/share/account/app"
	"github.com/mrokoo/goERP/internal/share/account/app/mock"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/stretchr/testify/assert"
)

var ErrNotFound = errors.New("the docment is not found")

func TestAccountService(t *testing.T) {
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

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)

	m.EXPECT().Get("A001").Return(&want, nil)
	m.EXPECT().Get("A002").Return(nil, ErrNotFound)

	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)

	t.Run("success", func(t *testing.T) {
		got, _ := s.GetAccount("A001")
		assert.Equal(want, *got, "this is equal")
	})

	t.Run("failure", func(t *testing.T) {
		_, err := s.GetAccount("A002")
		assert.ErrorIs(err, ErrNotFound, "")
	})
}
