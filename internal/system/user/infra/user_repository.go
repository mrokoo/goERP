package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/system/user/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	var list []*model.User
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var users []*domain.User
	for i := range list {
		users = append(users, toDomain(list[i]))
	}
	return users, nil
}

func (r *UserRepository) GetByID(ID string) (*domain.User, error) {
	user := model.User{
		ID: ID,
	}
	result := r.db.First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&user), nil
}

func (r *UserRepository) Save(user *domain.User) error {
	i := toModel(user)
	result := r.db.Create(i)
	return result.Error
}

func (r *UserRepository) Replace(user *domain.User) error {
	i := toModel(user)
	result := r.db.Save(i)
	return result.Error
}

func (r *UserRepository) Delete(ID string) error {
	result := r.db.Delete(&model.User{
		ID: ID,
	})
	return result.Error
}

