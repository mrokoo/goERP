package domain

import (
	"time"
)

type PurchaseOrder struct {
	ID           string    `json:"id"`
	WarehouseID  string    `json:"warehouse_id"`
	SupplierID   string    `json:"supplier_id"`
	UserID       string    `json:"user_id"`
	AccountID    string    `json:"account_id"`
	OtherCost    float64   `json:"other_cost"`
	TotalCost    float64   `json:"total_cost"`
	ActalPayment float64   `json:"actal_payment"`
	Debt         float64   `json:"debt"`
	CreatedAt    time.Time `json:"created_at"`
	IsValidated  bool      `json:"is_validated"`
	Basic        string    `json:"basic"`
	Items        []Item    `json:"items"`
	Kind         Kind      `json:"kind"`
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
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
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
