package app

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
)

type UserService interface {
	GetUser(userID uuid.UUID) (*domain.User, error)
	GetUserList() ([]*domain.User, error)
	AddUser(user *domain.User) error
	ReplaceUser(user *domain.User) error
	DeleteUser(userID uuid.UUID) error
}

type UserServiceImpl struct {
	repo domain.Repository
}

func NewUserServiceImpl(repo domain.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetUser(userID uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUserList() ([]*domain.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) AddUser(user *domain.User) error {
	err := s.repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) ReplaceUser(user *domain.User) error {
	if err := s.repo.Replace(user); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(userID uuid.UUID) error {
	if err := s.repo.Delete(userID); err != nil {
		return err
	}
	return nil
}
