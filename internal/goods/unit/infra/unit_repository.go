package repository

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Unit = domain.Unit

type UnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) *UnitRepository {
	db.AutoMigrate(&Unit{}) //自动迁移
	return &UnitRepository{
		db: db,
	}
}

func (r *UnitRepository) GetAll() ([]*Unit, error) {
	var units []Unit
	result := r.db.Find(&units)
	if err := result.Error; err != nil {
		return nil, err
	}
	var unitsp []*Unit
	for _, v := range units {
		unitsp = append(unitsp, &v)
	}
	return unitsp, nil
}

func (r *UnitRepository) GetByID(unitID uuid.UUID) (*Unit, error) {
	unit := Unit{
		ID: unitID,
	}
	result := r.db.First(&unit)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &unit, nil
}

func (r *UnitRepository) Save(unit *domain.Unit) error {
	result := r.db.Create(unit)
	return result.Error
}

func (r *UnitRepository) Replace(unit *domain.Unit) error {
	result := r.db.Save(unit)
	return result.Error
}

func (r *UnitRepository) Delete(unitID uuid.UUID) error {
	result := r.db.Delete(&Unit{
		ID: unitID,
	})
	return result.Error
}
