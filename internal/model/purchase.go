package model

import "time"

type PurchaseOrder struct {
	ID           string              `gorm:"primaryKey"`
	SupplierID   string              `gorm:"size:191"`
	Supplier     Supplier            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WarehouseID  string              `gorm:"size:191"`
	Warehouse    Warehouse           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID       string              `gorm:"size:191"`
	User         User                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	IsValidated  bool                `gorm:"default:false"`
	Items        []PurchaseOrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountID    string              `gorm:"size:191"`
	Account      Account             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	OtherCost    float64
	TotalCost    float64
	ActalPayment float64
	Debt         float64
	CreatedAt    time.Time
	Basic        string // 用于退货单
	Kind         string
}

type PurchaseOrderItem struct {
	PurchaseOrderID string  `gorm:"primaryKey;size:191"`
	ProductID       string  `gorm:"primaryKey;size:191"`
	Product         Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity        int
	Price           float64
}
