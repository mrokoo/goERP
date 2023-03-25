//go:generate mockgen -destination=./mock/mock_account_repository.go -package=mock github.com/mrokoo/goERP/internal/share/account/domain Repository
package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrokoo/goERP/internal/share/account/app"
	"github.com/mrokoo/goERP/internal/share/account/app/mock"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotFound = mongo.ErrNoDocuments
var ErrAccountInVaildated = app.ErrAccountInVaildated

func TestAccountService_GetAccount(t *testing.T) {
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
		assert.Equal(want, *got)
	})

	t.Run("notfound", func(t *testing.T) {
		got, _ := s.GetAccount("A002")
		assert.Nil(got)
	})
}

func TestAccountService_GetAccountList(t *testing.T) {
	want := []domain.Account{
		{
			ID:      "A001",
			Name:    "账号1",
			Type:    3,
			Holder:  "张三",
			Number:  "402901000226",
			Note:    "测试实例",
			State:   2,
			Balance: 2000,
		},

		{
			ID:      "A002",
			Name:    "账号2",
			Type:    2,
			Holder:  "张三",
			Number:  "402901000225",
			Note:    "测试实例2",
			State:   2,
			Balance: 1000,
		},
	}

	assert := assert.New(t)

	ctrl := gomock.NewController(t)

	m := mock.NewMockRepository(ctrl)

	m.EXPECT().GetAll().Return(want, nil)

	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)
	t.Run("success", func(t *testing.T) {
		got, _ := s.GetAccountList()
		assert.Equal(want, got)
	})
}

func TestAccountService_UpdateAccount(t *testing.T) {
	p := domain.Account{
		ID:      "A001",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "	402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	p2 := domain.Account{
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
	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)

	m.EXPECT().Update(p).Return(nil)
	m.EXPECT().Update(p2).Return(ErrNotFound)

	t.Run("success", func(t *testing.T) {
		err := s.UpdateAccount(p)
		assert.NoError(err)
	})

	// 根据传入account的ID属性没有查找到对应的文档，导致错误
	t.Run("failure", func(t *testing.T) {
		err := s.UpdateAccount(p2)
		assert.ErrorIs(err, ErrNotFound)
	})
}

func TestAccountService_AddAccount(t *testing.T) {

	p := domain.Account{
		ID:      "A001",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	p2 := domain.Account{
		ID:      "A002",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)
	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)

	t.Run("success", func(t *testing.T) {
		m.EXPECT().Get("A001").Return(nil, ErrNotFound)
		m.EXPECT().Save(p).Return(nil)
		err := s.AddAccount(p)
		assert.NoError(err)
	})

	t.Run("failure", func(t *testing.T) {
		m.EXPECT().Get("A002").Return(&p2, nil)
		err := s.AddAccount(p2)
		assert.ErrorIs(err, ErrAccountInVaildated)
	})
}

func TestAccountService_DeleteAccount(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)
	ds := domain.NewCheckingAccountValidityService(m)
	s := app.NewAccountServiceImpl(ds, m)

	t.Run("success", func(t *testing.T) {
		m.EXPECT().Delete("A001").Return(nil)
		err := s.DeleteAccount("A001")
		assert.NoError(err)
	})

	t.Run("failure", func(t *testing.T) {
		m.EXPECT().Delete("A002").Return(ErrNotFound)
		err := s.DeleteAccount("A002")
		assert.ErrorIs(err, ErrNotFound)
	})

}
