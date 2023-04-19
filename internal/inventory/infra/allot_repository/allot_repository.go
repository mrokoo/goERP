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
	db.AutoMigrate(&MySQLItem{})
	return &AllotRepository{
		db: db,
	}
}

func (r *AllotRepository) GetAll() ([]*allot.Allot, error) {
	var allots []*MySQLAllot
	if err := r.db.Preload("Items").Find(&allots).Error; err != nil {
		return nil, err
	}
	return toAllots(allots), nil
}

func (r *AllotRepository) GetByID(ID string) (*allot.Allot, error) {
	var allot *MySQLAllot
	if err := r.db.Preload("Items").First(&allot, ID).Error; err != nil {
		return nil, err
	}
	allot_ := allot.ToAllot()
	return allot_, nil
}

func (r *AllotRepository) Save(allot *allot.Allot) error {
	allot_ := toMySQLAllot(allot)
	return r.db.Create(allot_).Error
}
