package app

import (
	"github.com/google/uuid"
	category "github.com/mrokoo/goERP/internal/goods/category/app"
	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"github.com/mrokoo/goERP/internal/goods/product/domain/valueobj"
)

type ProductService interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	DeleteProduct(productId string) error
	UpdateCategory(productID string, categoryID uuid.UUID) error
}

type ProductServiceImpl struct {
	productRepository domain.ProductRepository
	categoryService   category.CategoryService
}

func NewProductServiceImpl(productRepository domain.ProductRepository, categoryService category.CategoryService) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: productRepository,
		categoryService:   categoryService,
	}
}

func (s *ProductServiceImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	state, err := valueobj.NewState(int(product.State)) // Set state to STATE_ACTIVE
	if err != nil {
		return nil, err
	}
	product.State = state
	if err := product.Validate(); err != nil {
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

func (s *ProductServiceImpl) UpdateCategory(productID string, categoryID *uuid.UUID) error {
	product, err := s.productRepository.Get(productID)
	if err != nil {
		return err
	}
	if _, err := s.categoryService.GetCategory(categoryID); err != nil {
		return err
	}
	product.CategoryID = *categoryID
	if err := product.Validate(); err != nil {
		return err
	}
	return s.productRepository.Save(product)
}
