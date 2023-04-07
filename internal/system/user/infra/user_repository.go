package repository

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type User = domain.User

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&User{}) //自动迁移
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]*User, error) {
	var users []User
	result := r.db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	var usersp []*User
	for i := range users {
		usersp = append(usersp, &users[i])
	}
	return usersp, nil
}

func (r *UserRepository) GetByID(userID uuid.UUID) (*User, error) {
	user := User{
		ID: userID,
	}
	result := r.db.First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Save(user *domain.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) Replace(user *domain.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *UserRepository) Delete(userID uuid.UUID) error {
	result := r.db.Delete(&User{
		ID: userID,
	})
	return result.Error
}
