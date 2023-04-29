package take_repository

import (
	"github.com/mrokoo/goERP/internal/inventory/domain/aggregate/take"
	"github.com/mrokoo/goERP/internal/model"
	"gorm.io/gorm"
)

type TakeRepository struct {
	db *gorm.DB
}

func NewTakeRepository(db *gorm.DB) *TakeRepository {
	return &TakeRepository{
		db: db,
	}
}

func (r *TakeRepository) GetAll() ([]*take.Take, error) {
	var list []*model.Take
	if err := r.db.Preload("Items").Find(&list).Error; err != nil {
		return nil, err
	}
	var takes []*take.Take
	for _, take := range list {
		takes = append(takes, toDomain(take))
	}
	return takes, nil
}

func (r *TakeRepository) GetByID(ID string) (*take.Take, error) {
	var take *model.Take
	if err := r.db.Preload("Items").First(&take, ID).Error; err != nil {
		return nil, err
	}
	return toDomain(take), nil
}

func (r *TakeRepository) Save(take *take.Take) error {
	i := toModel(take)
	return r.db.Create(i).Error
}
