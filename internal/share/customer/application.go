package customer

import "context"

type CustomerApplicationService struct {
	repo Respository
}

func (c CustomerApplicationService) SaveCustomer(ctx context.Context, customer Customer) error {
	if err := c.repo.SaveCustomer(ctx, customer); err != nil {
		return err
	}
	return nil
}

func NewCustomerApplicationService(repo Respository) CustomerApplicationService {
	return CustomerApplicationService{
		repo: repo,
	}
}
