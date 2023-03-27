package app

import (
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj"
)

type ProductService interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	DeleteProduct(productId string) error
	UpdateProduct(product *domain.Product) error
}

type ProductServiceImpl struct {
	productRepository              domain.ProductRepository
	checkingProductValidityService domain.CheckingProductValidityService
}

func NewProductServiceImpl(productRepository domain.ProductRepository, s domain.CheckingProductValidityService) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository:              productRepository,
		checkingProductValidityService: s,
	}
}

func (s *ProductServiceImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	state, err := valueobj.NewState(int(product.State)) // Set state to STATE_ACTIVE
	if err != nil {
		return nil, err
	}
	product.State = state

	if err := s.checkingProductValidityService.IsValidated(product); err != nil {
		return nil, err
	}

	if err := s.productRepository.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetAllProducts() ([]domain.Product, error) {
	products, err := s.productRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductServiceImpl) DeleteProduct(productId string) error {
	if err := s.productRepository.Delete(productId); err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImpl) UpdateProduct(product *domain.Product) error {
	if err := domain.CheckDate(product); err != nil {
		return err
	}
	if err := s.productRepository.Save(product); err != nil {
		return err
	}
	return nil
}
