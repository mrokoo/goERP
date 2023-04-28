package app

import (
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
)

type UnitService interface {
	GetUnit(unitID string) (*domain.Unit, error)
	GetUnitList() ([]*domain.Unit, error)
	AddUnit(unit *domain.Unit) error
	ReplaceUnit(unit *domain.Unit) error
	DeleteUnit(unitID string) error
}

type UnitServiceImpl struct {
	unitRepository domain.Repository
}

func NewUnitServiceImpl(unitRepository domain.Repository) *UnitServiceImpl {
	return &UnitServiceImpl{
		unitRepository: unitRepository,
	}
}

func (s *UnitServiceImpl) GetUnitList() ([]*domain.Unit, error) {
	categories, err := s.unitRepository.GetAll()
	return categories, err
}

func (s *UnitServiceImpl) GetUnit(unitID string) (*domain.Unit, error) {
	unit, err := s.unitRepository.GetByID(unitID)
	return unit, err
}

func (s *UnitServiceImpl) AddUnit(unit *domain.Unit) error {
	if err := s.unitRepository.Save(unit); err != nil {
		return err
	}
	return nil
}

func (s *UnitServiceImpl) ReplaceUnit(unit *domain.Unit) error {
	err := s.unitRepository.Replace(unit)
	if err != nil {
		return err
	}
	return nil
}

func (s *UnitServiceImpl) DeleteUnit(unitID string) error {
	if err := s.unitRepository.Delete(unitID); err != nil {
		return err
	}
	return nil
}
