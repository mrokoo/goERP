package stock

type Stock struct {
	ProductID   string `json:"product_id"`
	WarehouseID string `json:"warehouse_id"`
	Amount      int    `json:"amount"`
}
