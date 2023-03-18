package supplier

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the supplierId is not unique")

type Repositiory interface {
	SaveSupplier(ctx context.Context, supplier Supplier) error
	ChangeSupplier(ctx context.Context, supplier Supplier) error
	FetchAllSuppliers(ctx context.Context) ([]Supplier, error)
	DeleteSupplier(ctx context.Context, supplierID SupplierId) error
}

type MongoRespository struct {
	suppliers *mongo.Collection
}

func (mr *MongoRespository) SaveSupplier(ctx context.Context, supplier Supplier) error {
	mongoS := toMongoSupplier(supplier)
	_, err := mr.suppliers.InsertOne(ctx, mongoS)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("fail to persist supplier: %w", ErrNotUID)
		}
		return fmt.Errorf("fail to persist supplier: %w", err)
	}
	return nil
}

func (mr *MongoRespository) DeleteSupplier(ctx context.Context, supplierID SupplierId) error {
	filter := bson.D{{Key: "id", Value: supplierID}}
	result, err := mr.suppliers.DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the supplier")
	}
	if err != nil {
		return fmt.Errorf("fail to delete supplier: %w", err)
	}
	return nil
}

func (mr *MongoRespository) ChangeSupplier(ctx context.Context, supplier Supplier) error {
	mongoS := toMongoSupplier(supplier)
	filter := bson.D{{Key: "id", Value: supplier.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: mongoS.Name},
		{Key: "contact", Value: mongoS.Contact},
		{Key: "email", Value: mongoS.Email},
		{Key: "address", Value: mongoS.Address},
		{Key: "account", Value: mongoS.Account},
		{Key: "bank", Value: mongoS.Bank},
		{Key: "note", Value: mongoS.Note},
		{Key: "state", Value: mongoS.State},
		{Key: "debt", Value: mongoS.debt},
	}}}

	result, err := mr.suppliers.UpdateOne(ctx, filter, update)
	if result.MatchedCount == 0 {
		return errors.New("fail to query the supplier")
	}
	if err != nil {
		return fmt.Errorf("fail to change supplier: %w", err)
	}
	return nil
}

func (mr *MongoRespository) FetchAllSuppliers(ctx context.Context) ([]Supplier, error) {
	cursor, err := mr.suppliers.Find(ctx, bson.D{})

	if err != nil {
		return []Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}

	var results []Supplier
	if err = cursor.All(ctx, &results); err != nil {
		return []Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	return results, nil
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRespository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	suppliers := client.Database("goERP").Collection("suppliers")
	return &MongoRespository{
		suppliers: suppliers,
	}, nil
}

type MongoSupplier struct {
	ID      SupplierId
	Name    Name
	Contact ContactName
	Email   Email
	Address Address
	Account Account
	Bank    BankName
	Note    string
	State   StateType
	debt    money.Money
	time    time.Time
}

func toMongoSupplier(s Supplier) MongoSupplier {
	return MongoSupplier{
		ID:      s.ID,
		Name:    s.Name,
		Contact: s.Contact,
		Email:   s.Email,
		Address: s.Address,
		Account: s.Account,
		Bank:    s.Bank,
		Note:    s.Note,
		State:   s.State,
		debt:    s.Debt,
		time:    time.Now(),
	}
}
