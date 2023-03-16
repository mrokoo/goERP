package customer

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the customerId is not unique")

// customer 领域层存储库接口
type Respository interface {
	SaveCustomer(ctx context.Context, customer Customer) error
	DeleteCustomer(ctx context.Context, customerID CustomerId) error
	ChangeCustomer(ctx context.Context, customer Customer) error
	FetchAllCustomers(ctx context.Context) error
}

// mongo存储库
type MongoRespository struct {
	customer *mongo.Collection
}

func (mr *MongoRespository) SaveCustomer(ctx context.Context, customer Customer) error {
	mongoCustomer := toMongoCustomer(customer)
	_, err := mr.customer.InsertOne(ctx, mongoCustomer)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("fail to persist customer: %w", ErrNotUID)
		}

		return fmt.Errorf("fail to persist customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) DeleteCustomer(ctx context.Context, customerID CustomerId) error {
	filter := bson.D{{Key: "ID", Value: customerID}}
	_, err := mr.customer.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("fail to delete customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) ChangeCustomer(ctx context.Context, customer Customer) error {
	mongoCustomer := toMongoCustomer(customer)
	filter := bson.D{{Key: "ID", Value: customer.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: mongoCustomer.Name},
		{Key: "grade", Value: mongoCustomer.Grade},
		{Key: "contact", Value: mongoCustomer.Contact},
		{Key: "phone", Value: mongoCustomer.PhoneNumber},
		{Key: "address", Value: mongoCustomer.Address},
		{Key: "note", Value: mongoCustomer.Note},
		{Key: "state", Value: mongoCustomer.State},
	}}}
	_, err := mr.customer.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("fail to change customer: %w", err)
	}
	return nil
}

func (mr *MongoRespository) FetchAllCustomers(ctx context.Context) error {
	return fmt.Errorf("this is a error")
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRespository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	customer := client.Database("goERP").Collection("customers")
	return &MongoRespository{
		customer: customer,
	}, nil
}

type mongoCustomer struct {
	ID          CustomerId
	Name        Name
	Grade       GradeType
	Contact     ContactName
	PhoneNumber PhoneNumber
	Address     Address
	Note        string
	State       StateType
}

func toMongoCustomer(c Customer) mongoCustomer {
	return mongoCustomer{
		ID:          c.ID,
		Name:        c.Name,
		Grade:       c.Grade,
		Contact:     c.Contact,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		Note:        c.Note,
		State:       c.State,
	}
}
