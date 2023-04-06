package domain

import (
	repository "github.com/mrokoo/goERP/internal/goods/unit/infra"
)

type CheckingProductValidityService struct {
	productRepository Repository
}

func NewCheckingProductValidityService(productRepository Repository) *CheckingProductValidityService {
	return &CheckingProductValidityService{
		productRepository: productRepository,
	}
}

func (ds *CheckingProductValidityService) IsValidated(product *Product) bool {
	// ID唯一性校验
	_, err := ds.productRepository.GetByID(product.ID)
	return err == repository.ErrNotFound
}
