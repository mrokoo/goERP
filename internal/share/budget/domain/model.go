package domain

type Budget struct {
	ID   string     `json:"id" gorm:"primaryKey;<-:create"`
	Name string     `json:"name"`
	Type BudgetType `json:"type" gorm:"default:in"`
	Note string     `json:"note"`
}

type BudgetType string

const (
	BUDGET_OUT = "out"
	BUDGET_IN  = "in"
)
