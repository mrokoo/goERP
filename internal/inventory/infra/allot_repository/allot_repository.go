package allot_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	"gorm.io/gorm"
)

type AllotRepository struct {
	db *gorm.DB
}

func NewAllotRepository(db *gorm.DB) *AllotRepository {
	db.AutoMigrate(&MySQLAllot{})
	return &AllotRepository{
		db: db,
	}
}

func (r *AllotRepository) GetAll() ([]*allot.Allot, error) {
	var allots []*allot.Allot
	if err := r.db.Find(&allots).Error; err != nil {
		return nil, err
	}
	return allots, nil
}

func (r *AllotRepository) GetByID(ID string) (*allot.Allot, error) {
	var allot allot.Allot
	if err := r.db.First(&allot, ID).Error; err != nil {
		return nil, err
	}
	return &allot, nil
}

func (r *AllotRepository) Save(allot *allot.Allot) error {
	allot_ := toMySQLAllot(allot)
	return r.db.Save(allot_).Error
}
