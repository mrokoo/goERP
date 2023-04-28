package domain

const (
	CollectionBudget = "budgets"
)

type Repository interface {
	GetAll() ([]*Budget, error)
	GetByID(budgetID string) (*Budget, error)
	Save(budget *Budget) error
	Replace(budget *Budget) error
	Delete(budgetID string) error
}
