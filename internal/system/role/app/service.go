package app

import "github.com/mrokoo/goERP/internal/system/role/domain"

type RoleServcie interface {
	GetRole(string) (*domain.Role, error)
	GetRoleList() ([]*domain.Role, error)
	AddRole(*domain.Role) error
	UpdateRole(*domain.Role) error
	DeleteRole(string) error
}

type RoleServiceImpl struct {
	roleRepository domain.Repository
}

func NewRoleServiceImpl(roleRepository domain.Repository) *RoleServiceImpl {
	return &RoleServiceImpl{
		roleRepository: roleRepository,
	}
}

func (s *RoleServiceImpl) GetRole(name string) (*domain.Role, error) {
	role, err := s.roleRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleServiceImpl) GetRoleList() ([]*domain.Role, error) {
	rolelist, err := s.roleRepository.FindaAll()
	if err != nil {
		return nil, err
	}
	return rolelist, nil
}

func (s *RoleServiceImpl) AddRole(role *domain.Role) error {
	err := s.roleRepository.Save(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleServiceImpl) UpdateRole(role *domain.Role) error {
	err := s.roleRepository.Update(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleServiceImpl) DeleteRole(name string) error {
	err := s.roleRepository.Delete(name)
	if err != nil {
		return err
	}
	return nil
}
