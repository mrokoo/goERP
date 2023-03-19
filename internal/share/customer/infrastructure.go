package customer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the customerId is not unique")

type Repository interface {
	LoadCustomer(ctx context.Context, customerID CustomerId) (Customer, error)
	SaveCustomer(ctx context.Context, customer Customer) error
	DeleteCustomer(ctx context.Context, customerID CustomerId) error
	ChangeCustomer(ctx context.Context, customer Customer) error
	FetchAllCustomers(ctx context.Context) ([]Customer, error)

	CheckCustomerID(ctx context.Context, customerID CustomerId) bool
}

// mongo存储库
type MongoRespository struct {
	customers *mongo.Collection
}

func (mr *MongoRespository) LoadCustomer(ctx context.Context, customerID CustomerId) (Customer, error) {
	filter := bson.D{{Key: "id", Value: customerID}}
	result := mr.customers.FindOne(ctx, filter)
	var customer Customer

	if err := result.Decode(customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return Customer{}, mongo.ErrNoDocuments
		} else {
			return Customer{}, err
		}
	}
	return customer, nil
}

func (mr *MongoRespository) SaveCustomer(ctx context.Context, customer Customer) error {
	mongoCustomer := toMongoCustomer(customer)
	_, err := mr.customers.InsertOne(ctx, mongoCustomer)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("fail to persist customer: %w", ErrNotUID)
		}

		return fmt.Errorf("fail to persist customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) DeleteCustomer(ctx context.Context, customerID CustomerId) error {
	filter := bson.D{{Key: "id", Value: customerID}}
	result, err := mr.customers.DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the customer")
	}
	if err != nil {
		return fmt.Errorf("fail to delete customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) ChangeCustomer(ctx context.Context, customer Customer) error {
	mongoCustomer := toMongoCustomer(customer)
	filter := bson.D{{Key: "id", Value: customer.ID}}
	result, err := mr.customers.ReplaceOne(ctx, filter, mongoCustomer)

	if result.MatchedCount == 0 {
		return errors.New("fail to query the customer")
	}
	if err != nil {
		return fmt.Errorf("fail to change customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) FetchAllCustomers(ctx context.Context) ([]Customer, error) {
	cursor, err := mr.customers.Find(ctx, bson.D{})

	if err != nil {
		return []Customer{}, fmt.Errorf("fail to fetch customers: %w", err)
	}

	var results []Customer
	if err = cursor.All(ctx, &results); err != nil {
		return []Customer{}, fmt.Errorf("fail to fetch customers: %w", err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}
	return results, nil
}

func (mr *MongoRespository) CheckCustomerID(ctx context.Context, customerId CustomerId) bool {
	filter := bson.D{{Key: "id", Value: customerId}}
	_, err := mr.customers.Find(ctx, filter)
	return err == nil
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRespository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	customers := client.Database("goERP").Collection("customers")
	return &MongoRespository{
		customers: customers,
	}, nil
}

type mongoCustomer struct {
	Customer `bson:"inline"`
	Time     time.Time
}

func toMongoCustomer(c Customer) mongoCustomer {
	return mongoCustomer{
		Customer: c,
		Time:     time.Now(),
	}
}
