package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrokoo/goERP/internal/share/customer/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	customers *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	customers := client.Database("goERP").Collection("customers")
	return &MongoRepository{
		customers: customers,
	}
}

// 插入新文档
func (r *MongoRepository) Save(customer domain.Customer) error {
	_, err := r.customers.InsertOne(context.Background(), customer)
	if err != nil {
		return fmt.Errorf("fail to persist customer: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(customer domain.Customer) error {
	filter := bson.D{{Key: "id", Value: customer.ID}}
	_, err := r.customers.ReplaceOne(context.Background(), filter, customer)
	if err != nil {
		return fmt.Errorf("fail to change customer: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(customerId domain.CustomerId) (*domain.Customer, error) {
	filter := bson.D{{Key: "id", Value: customerId}}
	var customer domain.Customer
	err := r.customers.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.Customer, error) {
	cursor, err := r.customers.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.Customer{}, fmt.Errorf("fail to fetch customers: %w", err)
	}

	var results []domain.Customer
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.Customer{}, fmt.Errorf("fail to fetch customers: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(customerID domain.CustomerId) error {
	filter := bson.D{{Key: "id", Value: customerID}}
	result, err := r.customers.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the customer")
	}
	if err != nil {
		return fmt.Errorf("fail to delete customer: %w", err)
	}
	return nil
}
