package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotUniqueID = errors.New("the id is not unique")

type CheckingProductValidityService struct {
	productRepository ProductRepository
}

func NewCheckingProductValidityService(productRepository ProductRepository) *CheckingProductValidityService {
	return &CheckingProductValidityService{
		productRepository: productRepository,
	}
}

func (ds *CheckingProductValidityService) IsValidated(product *Product) error {
	// ID唯一性校验
	_, err := ds.productRepository.Get(product.ID)
	if err != mongo.ErrNoDocuments {
		return ErrNotUniqueID
	}
	// 日期校验
	if err := CheckDate(product); err != nil {
		return err
	}
	return nil
}
