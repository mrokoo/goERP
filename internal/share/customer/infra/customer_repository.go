package repository

import (
	"context"

	"github.com/mrokoo/goERP/internal/share/customer/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	customers *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	customers := db.Collection(collection)
	return &MongoRepository{
		customers: customers,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Customer, error) {
	cursor, err := r.customers.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Customer
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(customerID string) (*domain.Customer, error) {
	filter := bson.D{{Key: "id", Value: customerID}}
	var customer domain.Customer
	err := r.customers.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *MongoRepository) Save(customer *domain.Customer) error {
	_, err := r.customers.InsertOne(context.Background(), customer)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(customer *domain.Customer) error {
	filter := bson.D{{Key: "id", Value: customer.ID}}
	var c domain.Customer
	err := r.customers.FindOneAndReplace(context.Background(), filter, customer).Decode(&c)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(customerID string) error {
	filter := bson.D{{Key: "id", Value: customerID}}
	var c domain.Customer
	err := r.customers.FindOneAndDelete(context.Background(), filter).Decode(&c)
	if err != nil {
		return err
	}
	return nil
}
