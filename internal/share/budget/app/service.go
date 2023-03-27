package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
)

var ErrNotFound = errors.New("the docment is not found")

type BudgetService interface {
	GetBudget(budgetId uuid.UUID) (*domain.Budget, error)
	GetBudgetList() ([]domain.Budget, error)
	AddBudget(budget domain.Budget) error
	UpdateBudget(budget domain.Budget) error
	DeleteBudget(budgetId uuid.UUID) error
}

type BudgetServiceImpl struct {
	repo domain.Repository
}

func NewBudgetServiceImpl(repo domain.Repository) *BudgetServiceImpl {
	return &BudgetServiceImpl{
		repo: repo,
	}
}

func (s *BudgetServiceImpl) GetBudget(budgetId uuid.UUID) (*domain.Budget, error) {
	budget, err := s.repo.Get(budgetId)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (s *BudgetServiceImpl) GetBudgetList() ([]domain.Budget, error) {
	budgets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (s *BudgetServiceImpl) AddBudget(budget domain.Budget) error {
	err := s.repo.Save(budget)
	if err != nil {
		return err
	}
	return nil
}

func (s *BudgetServiceImpl) UpdateBudget(budget domain.Budget) error {
	if err := s.repo.Update(budget); err != nil {
		return err
	}
	return nil
}

func (s *BudgetServiceImpl) DeleteBudget(budgetId uuid.UUID) error {
	if err := s.repo.Delete(budgetId); err != nil {
		return err
	}
	return nil
}
