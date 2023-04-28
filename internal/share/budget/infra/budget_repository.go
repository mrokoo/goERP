package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{
		db: db,
	}
}

func (r *BudgetRepository) GetAll() ([]*domain.Budget, error) {
	var list []model.Budget
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var budgets []*domain.Budget
	for i := range list {
		budgets = append(budgets, toDomain(&list[i]))
	}
	return budgets, nil
}

func (r *BudgetRepository) GetByID(ID string) (*domain.Budget, error) {
	i := model.Budget{
		ID: ID,
	}
	result := r.db.First(&i)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&i), nil
}

func (r *BudgetRepository) Save(budget *domain.Budget) error {
	i := toModel(budget)
	result := r.db.Create(i)
	return result.Error
}

func (r *BudgetRepository) Replace(budget *domain.Budget) error {
	i := toModel(budget)
	result := r.db.Save(i)
	return result.Error
}

func (r *BudgetRepository) Delete(ID string) error {
	result := r.db.Delete(&model.Budget{
		ID: ID,
	})
	return result.Error
}
