package take_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	"gorm.io/gorm"
)

type TakeRepository struct {
	db *gorm.DB
}

func NewTakeRepository(db *gorm.DB) *TakeRepository {
	db.AutoMigrate(&MySQLTake{})
	db.AutoMigrate(&MySQLItem{})
	return &TakeRepository{
		db: db,
	}
}

func (r *TakeRepository) GetAll() ([]*take.Take, error) {
	var takes []*MySQLTake
	if err := r.db.Preload("Items").Find(&takes).Error; err != nil {
		return nil, err
	}
	return toTakes(takes), nil
}

func (r *TakeRepository) GetByID(ID string) (*take.Take, error) {
	var take *MySQLTake
	if err := r.db.Preload("Items").First(&take, ID).Error; err != nil {
		return nil, err
	}
	take_ := take.toTake()
	return &take_, nil
}

func (r *TakeRepository) Save(take *take.Take) error {
	take_ := toMySQLTake(*take)
	return r.db.Create(&take_).Error
}
