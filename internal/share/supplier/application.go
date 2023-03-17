package supplier

import (
	"context"
)

type SupplierApplicationService struct {
	repo Repositiory
}

func (s SupplierApplicationService) SaveSupplier(ctx context.Context, req interface{}) error {
	return nil
}

func (s SupplierApplicationService) DeleteSupplier(ctx context.Context, id string) error {
	return nil
}

func (s SupplierApplicationService) UpdateSupplier(ctx context.Context, req interface{}) error {
	return nil
}

func (s SupplierApplicationService) GetSupplierList(ctx context.Context) ([]Supplier, error) {
	return []Supplier{}, nil
}

func NewSupplierApplicationService(repo Repositiory) SupplierApplicationService {
	return SupplierApplicationService{repo: repo}
}
