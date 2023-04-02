package app

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
)

type UnitService interface {
	GetUnit(unitID uuid.UUID) (*domain.Unit, error)
	GetUnitList() ([]*domain.Unit, error)
	AddUnit(unit *domain.Unit) error
	ReplaceUnit(unit *domain.Unit) error
	DeleteUnit(unitID uuid.UUID) error
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
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *UnitServiceImpl) GetUnit(unitID uuid.UUID) (*domain.Unit, error) {
	unit, err := s.unitRepository.GetByID(unitID)
	if err != nil {
		return nil, err
	}
	return unit, nil
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

func (s *UnitServiceImpl) DeleteUnit(unitID uuid.UUID) error {
	if err := s.unitRepository.Delete(unitID); err != nil {
		return err
	}
	return nil
}
