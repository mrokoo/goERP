package app

import (
	"errors"

	"github.com/mrokoo/goERP/internal/goods/product/domain"
)

var ErrProductInVaildated = errors.New("客户ID检验无效")

type ProductService interface {
	GetProduct(productID string) (*domain.Product, error)
	GetProductList() ([]*domain.Product, error)
	AddProduct(product *domain.Product) error
	ReplaceProduct(product *domain.Product) error
	DeleteProduct(productID string) error
}

type ProductServiceImpl struct {
	checkProductValidityService *domain.CheckingProductValidityService
	repo                         domain.Repository
}

func NewProductServiceImpl(checkProductValidityService *domain.CheckingProductValidityService, repo domain.Repository) *ProductServiceImpl {
	return &ProductServiceImpl{
		checkProductValidityService: checkProductValidityService,
		repo:                         repo,
	}
}

func (s *ProductServiceImpl) GetProduct(productID string) (*domain.Product, error) {
	product, err := s.repo.GetByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetProductList() ([]*domain.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductServiceImpl) AddProduct(product *domain.Product) error {

	if !s.checkProductValidityService.IsValidated(product) {
		return ErrProductInVaildated
	}
	err := s.repo.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImpl) ReplaceProduct(product *domain.Product) error {
	if err := s.repo.Replace(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImpl) DeleteProduct(productID string) error {
	if err := s.repo.Delete(productID); err != nil {
		return err
	}
	return nil
}
