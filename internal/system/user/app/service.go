package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
)

var ErrNotFound = errors.New("the docment is not found")

type UserService interface {
	GetUser(userId uuid.UUID) (*domain.User, error)
	GetUserList() ([]domain.User, error)
	AddUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(userId uuid.UUID) error
}

type UserServiceImpl struct {
	repo domain.Repository
}

func NewUserServiceImpl(repo domain.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetUser(userId uuid.UUID) (*domain.User, error) {
	user, err := s.repo.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUserList() ([]domain.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) AddUser(user domain.User) error {
	err := s.repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) UpdateUser(user domain.User) error {
	if err := s.repo.Update(user); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(userId uuid.UUID) error {
	if err := s.repo.Delete(userId); err != nil {
		return err
	}
	return nil
}
