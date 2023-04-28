package model

type SaleOrder struct {
	ID          string    `gorm:"primaryKey; size:191;"`
	WarehouseID string    `gorm:"size:191;"`
	Warehouse   Warehouse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CustomerID  string    `gorm:"size:191;"`
	Customer    Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID      string    `gorm:"size:191;"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt   string
	Basic       string
	Items       []SaleOrderItem
	Kind        string
}

type SaleOrderItem struct {
	ID          string  `gorm:"primaryKey; size:191;"`
	SaleOrderID string  `gorm:"size:191;"`
	ProductID   string  `gorm:"size:191;"`
	Product     Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity    int
	Price       float64
}
