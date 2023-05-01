package domain

import (
	"time"
)

type PurchaseOrder struct {
	ID           string
	WarehouseID  string
	SupplierID   string
	UserID       string
	AccountID    string
	OtherCost    float64
	TotalCost    float64
	ActalPayment float64
	Debt         float64
	CreatedAt    time.Time
	IsValidated  bool
	Basic        string
	Items        []Item
	Kind         Kind
}

func NewPurchaseOrder(ID string, warehouse string, supplier string, user string, account string, otherCost float64, actalPayment float64, basic string, items []Item, kind Kind) PurchaseOrder {
	totalCost := 0.0
	for _, item := range items {
		totalCost += item.Price * float64(item.Quantity)
	}
	debt := actalPayment - totalCost
	return PurchaseOrder{
		ID:           ID,
		WarehouseID:  warehouse,
		SupplierID:   supplier,
		UserID:       user,
		AccountID:    account,
		OtherCost:    otherCost,
		TotalCost:    totalCost,
		ActalPayment: actalPayment,
		Debt:         debt,
		CreatedAt:    time.Now(),
		IsValidated:  false,
		Basic:        basic,
		Items:        items,
		Kind:         kind,
	}
}

func (p *PurchaseOrder) GetTotalCost() float64 {
	c := 0.0
	for _, item := range p.Items {
		c += item.Price * float64(item.Quantity)
	}
	return p.OtherCost + c
}

func (p *PurchaseOrder) GetDebt() float64 {
	return p.ActalPayment - p.GetTotalCost()
}

type Item struct {
	ProductID string
	Quantity  int
	Price     float64
}

type Kind string

const (
	Order       Kind = "Order"
	ReturnOrder Kind = "ReturnOrder"
)

func NewItem(productID string, quantity int, price float64) Item {
	return Item{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}
