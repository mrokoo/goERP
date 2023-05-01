package model

type Warehouse struct {
	ID      string `gorm:"primaryKey;<-:create;size:191"`
	Name    string `gorm:"not null"`
	Admin   string
	Phone   string
	Address string
	Note    string
	State   string `gorm:"default:active"`
}

type Supplier struct {
	ID      string `gorm:"primaryKey;<-:create"`
	Name    string `gorm:"not null"`
	Contact string
	Email   string
	Address string
	Account string
	Bank    string
	Note    string
	State   string `gorm:"default:active"`
	Debt    float64
}

type Customer struct {
	ID      string `gorm:"primaryKey;<-:create"`
	Name    string `gorm:"not null"`
	Grade   string `gorm:"default:medium"`
	Contact string
	Phone   string
	Email   string
	Address string
	Note    string
	State   string `gorm:"default:active"`
	Debt    float64
}

type Budget struct {
	ID   string `gorm:"primaryKey;<-:create"`
	Name string
	Type string `gorm:"default:in"`
	Note string
}

type Account struct {
	ID      string `gorm:"primaryKey;<-:create"`
	Name    string `gorm:"not null"`
	Type    string `gorm:"default:other"`
	Holder  string
	Number  string
	Note    string
	State   string `gorm:"default:active"`
	Balance float64
}
