package app

import (
	"github.com/mrokoo/goERP/internal/share/budget/domain"
)

type BudgetService interface {
	GetBudget(budgetID string) (*domain.Budget, error)
	GetBudgetList() ([]*domain.Budget, error)
	AddBudget(budget *domain.Budget) error
	ReplaceBudget(budget *domain.Budget) error
	DeleteBudget(budgetID string) error
}

type BudgetServiceImpl struct {
	repo domain.Repository
}

func NewBudgetServiceImpl(repo domain.Repository) *BudgetServiceImpl {
	return &BudgetServiceImpl{
		repo: repo,
	}
}

func (s *BudgetServiceImpl) GetBudget(budgetID string) (*domain.Budget, error) {
	budget, err := s.repo.GetByID(budgetID)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (s *BudgetServiceImpl) GetBudgetList() ([]*domain.Budget, error) {
	budgets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (s *BudgetServiceImpl) AddBudget(budget *domain.Budget) error {
	err := s.repo.Save(budget)
	if err != nil {
		return err
	}
	return nil
}

func (s *BudgetServiceImpl) ReplaceBudget(budget *domain.Budget) error {
	if err := s.repo.Replace(budget); err != nil {
		return err
	}
	return nil
}

func (s *BudgetServiceImpl) DeleteBudget(budgetID string) error {
	if err := s.repo.Delete(budgetID); err != nil {
		return err
	}
	return nil
}
