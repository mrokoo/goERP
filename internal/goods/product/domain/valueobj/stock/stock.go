package stock

type Stock struct {
	Warehouse string `json:"warehouse_id"`
	Amount    int    `json:"amount"`
}
