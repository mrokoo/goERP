package customer

import "context"

type CustomerApplicationService struct {
	repo Respository
}

func (c CustomerApplicationService) SaveCustomer(ctx context.Context, req interface{}) error {
	// to do 检查req有效性
	if err := c.repo.SaveCustomer(ctx, req.(Customer)); err != nil {
		return err
	}
	return nil
}

func (c CustomerApplicationService) DeleteCustomer(ctx context.Context, id string) error {
	if err := c.repo.DeleteCustomer(ctx, CustomerId(id)); err != nil {
		return err
	}
	return nil
}

func (c CustomerApplicationService) UpdateCustomer(ctx context.Context, req interface{}) error {
	// to do 验证 req有效性。
	if err := c.repo.ChangeCustomer(ctx, req.(Customer)); err != nil {
		return err
	}
	return nil
}

func (c CustomerApplicationService) GetCustomerList(ctx context.Context) ([]Customer, error) {
	return c.repo.FetchAllCustomers(ctx)
}

func NewCustomerApplicationService(repo Respository) CustomerApplicationService {
	return CustomerApplicationService{
		repo: repo,
	}
}
