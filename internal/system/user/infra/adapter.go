package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/system/user/domain"
)

func toModel(user *domain.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Gender:   user.Gender,
		Password: user.Password,
	}
}

func toDomain(user *model.User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		Gender:   user.Gender,
		Password: user.Password,
	}
}
