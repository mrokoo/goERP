package domain

type SaleOrder struct {
	ID          string
	WarehouseID string
	CustomerID  string
	UserID      string
	CreatedAt   string
	Basic       string // 只对return有效
	Items       []Item
	Kind        Kind
}

func NewSaleOrder(ID string, warehouse string, customer string, user string, createTime, basic string, items []Item, kind Kind) SaleOrder {
	return SaleOrder{
		ID:          ID,
		WarehouseID: warehouse,
		CustomerID:  customer,
		UserID:      user,
		CreatedAt:   createTime,
		Basic:       basic,
		Items:       items,
		Kind:        kind,
	}
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
