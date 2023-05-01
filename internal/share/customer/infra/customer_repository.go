package repository

import (
	"github.com/mrokoo/goERP/internal/model"
	"github.com/mrokoo/goERP/internal/share/customer/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAll() ([]*domain.Customer, error) {
	var list []model.Customer
	result := r.db.Find(&list)
	if err := result.Error; err != nil {
		return nil, err
	}
	var customers []*domain.Customer
	for i := range list {
		customers = append(customers, toDomain(&list[i]))
	}
	return customers, nil
}

func (r *CustomerRepository) GetByID(ID string) (*domain.Customer, error) {
	customer := model.Customer{
		ID: ID,
	}
	result := r.db.First(&customer)
	if err := result.Error; err != nil {
		return nil, err
	}
	return toDomain(&customer), nil
}

func (r *CustomerRepository) Save(customer *domain.Customer) error {
	i := toModel(customer)
	result := r.db.Create(i)
	return result.Error
}

func (r *CustomerRepository) Replace(customer *domain.Customer) error {
	i := toModel(customer)
	result := r.db.Save(i)
	return result.Error
}

func (r *CustomerRepository) Delete(ID string) error {
	result := r.db.Delete(&model.Customer{
		ID: ID,
	})
	return result.Error
}
