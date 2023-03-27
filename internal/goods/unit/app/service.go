package app

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
)

type UnitServiceImpl struct {
	unitRepository domain.UnitRepository
}

func NewUnitServiceImpl(unitRepository domain.UnitRepository) *UnitServiceImpl {
	return &UnitServiceImpl{
		unitRepository: unitRepository,
	}
}

func (s *UnitServiceImpl) CreateUnit(name string, note string) (*domain.Unit, error) {
	unit := &domain.Unit{
		ID:   uuid.New(),
		Name: name,
		Note: note,
	}

	if err := s.unitRepository.Create(unit); err != nil {
		return nil, err
	}

	return unit, nil
}

func (s *UnitServiceImpl) ChangeUnit(unitId *uuid.UUID, name string, note string) (*domain.Unit, error) {
	c, err := s.unitRepository.Get(unitId)
	if err != nil {
		return nil, err
	}
	c.ChangeName(name)
	c.ChangeNote(note)
	if err := s.unitRepository.Save(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *UnitServiceImpl) GetAllUnits() ([]domain.Unit, error) {
	units, err := s.unitRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return units, nil
}

func (s *UnitServiceImpl) DeleteUnit(unitId *uuid.UUID) error {
	if err := s.unitRepository.Delete(unitId); err != nil {
		return err
	}
	return nil
}
