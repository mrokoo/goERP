package item

type Item struct {
	ProductID   string
	Quantity    int
	Price       float64
	TotalAmount float64
}

func New(productID string, quantity int, price float64) Item {
	return Item{
		ProductID:   productID,
		Quantity:    quantity,
		Price:       price,
		TotalAmount: float64(quantity) * price,
	}
}
