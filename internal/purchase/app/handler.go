package app

type PurchaseHandler struct {
	PurchaseService PurchaseService
}

func NewPurchaseHandler(purchaseService PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		PurchaseService: purchaseService,
	}
}


