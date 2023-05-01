package repository

import (
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type UnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) *UnitRepository {
	return &UnitRepository{
		db: db,
	}
}

func (r *UnitRepository) GetAll() ([]*domain.Unit, error) {
	var list []*model.Unit
	result := r.db.Find(&list)
	var units []*domain.Unit
	for i := range list {
		units = append(units, toDomain(list[i]))
	}
	return units, result.Error
}

func (r *UnitRepository) GetByID(unitID string) (*domain.Unit, error) {
	unit := model.Unit{
		ID: unitID,
	}
	result := r.db.First(&unit)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&unit), nil
}

func (r *UnitRepository) Save(unit *domain.Unit) error {
	i := toModel(unit)
	result := r.db.Create(i)
	return result.Error
}

func (r *UnitRepository) Replace(unit *domain.Unit) error {
	i := toModel(unit)
	result := r.db.Save(i)
	return result.Error
}

func (r *UnitRepository) Delete(unitID string) error {
	result := r.db.Delete(&model.Unit{
		ID: unitID,
	})
	return result.Error
}
