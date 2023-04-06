package repository

import (
	"github.com/mrokoo/goERP/internal/share/customer/domain"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound

type Customer = domain.Customer

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	db.AutoMigrate(&Customer{}) //自动迁移
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAll() ([]*Customer, error) {
	var customers []Customer
	result := r.db.Find(&customers)
	if err := result.Error; err != nil {
		return nil, err
	}
	var customersp []*Customer
	for i := range customers {
		customersp = append(customersp, &customers[i])
	}
	return customersp, nil
}

func (r *CustomerRepository) GetByID(customerID string) (*Customer, error) {
	customer := Customer{
		ID: customerID,
	}
	result := r.db.First(&customer)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) Save(customer *domain.Customer) error {
	result := r.db.Create(customer)
	return result.Error
}

func (r *CustomerRepository) Replace(customer *domain.Customer) error {
	result := r.db.Save(customer)
	return result.Error
}

func (r *CustomerRepository) Delete(customerID string) error {
	result := r.db.Delete(&Customer{
		ID: customerID,
	})
	return result.Error
}
