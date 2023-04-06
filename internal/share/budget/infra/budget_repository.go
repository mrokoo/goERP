package repository

import (
	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Budget = domain.Budget

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	db.AutoMigrate(&Budget{}) //自动迁移
	return &BudgetRepository{
		db: db,
	}
}

func (r *BudgetRepository) GetAll() ([]*Budget, error) {
	var budgets []Budget
	result := r.db.Find(&budgets)
	if err := result.Error; err != nil {
		return nil, err
	}
	var budgetsp []*Budget
	for i := range budgets {
		budgetsp = append(budgetsp, &budgets[i])
	}
	return budgetsp, nil
}

func (r *BudgetRepository) GetByID(budgetID uuid.UUID) (*Budget, error) {
	budget := Budget{
		ID: budgetID,
	}
	result := r.db.First(&budget)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepository) Save(budget *domain.Budget) error {
	result := r.db.Create(budget)
	return result.Error
}

func (r *BudgetRepository) Replace(budget *domain.Budget) error {
	result := r.db.Save(budget)
	return result.Error
}

func (r *BudgetRepository) Delete(budgetID uuid.UUID) error {
	result := r.db.Delete(&Budget{
		ID: budgetID,
	})
	return result.Error
}
