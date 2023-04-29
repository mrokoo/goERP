package allot_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/allot"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

type AllotRepository struct {
	db *gorm.DB
}

func NewAllotRepository(db *gorm.DB) *AllotRepository {
	return &AllotRepository{
		db: db,
	}
}

func (r *AllotRepository) GetAll() ([]*allot.Allot, error) {
	var list []*model.Allot
	if err := r.db.Preload("Items").Find(&list).Error; err != nil {
		return nil, err
	}
	var allots []*allot.Allot
	for i := range list {
		allots = append(allots, toDomain(list[i]))
	}
	return allots, nil
}

func (r *AllotRepository) GetByID(ID string) (*allot.Allot, error) {
	var allot *model.Allot
	if err := r.db.Preload("Items").First(&allot, ID).Error; err != nil {
		return nil, err
	}
	allot_ := toDomain(allot)
	return allot_, nil
}

func (r *AllotRepository) Save(allot *allot.Allot) error {
	allot_ := toModel(allot)
	return r.db.Create(allot_).Error
}
