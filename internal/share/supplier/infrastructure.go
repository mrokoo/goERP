package supplier

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the supplierId is not unique")

type Repository interface {
	SaveSupplier(ctx context.Context, supplier Supplier) error
	ChangeSupplier(ctx context.Context, supplier Supplier) error
	FetchAllSuppliers(ctx context.Context) ([]Supplier, error)
	DeleteSupplier(ctx context.Context, supplierID SupplierId) error
}

type MongoRepository struct {
	suppliers *mongo.Collection
}

func (mr *MongoRepository) SaveSupplier(ctx context.Context, supplier Supplier) error {
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

func (mr *MongoRepository) DeleteSupplier(ctx context.Context, supplierID SupplierId) error {
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

func (mr *MongoRepository) ChangeSupplier(ctx context.Context, supplier Supplier) error {
	mongoS := toMongoSupplier(supplier)
	filter := bson.D{{Key: "id", Value: supplier.ID}}

	result, err := mr.suppliers.ReplaceOne(ctx, filter, mongoS)
	if result.MatchedCount == 0 {
		return errors.New("fail to query the supplier")
	}
	if err != nil {
		return fmt.Errorf("fail to change supplier: %w", err)
	}
	return nil
}

func (mr *MongoRepository) FetchAllSuppliers(ctx context.Context) ([]Supplier, error) {
	cursor, err := mr.suppliers.Find(ctx, bson.D{})

	if err != nil {
		return []Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}

	var results []Supplier
	if err = cursor.All(ctx, &results); err != nil {
		return []Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}
	return results, nil
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	suppliers := client.Database("goERP").Collection("suppliers")
	return &MongoRepository{
		suppliers: suppliers,
	}, nil
}

type MongoSupplier struct {
	Supplier `bson:"inline"`
	Time     time.Time
}

func toMongoSupplier(s Supplier) MongoSupplier {
	return MongoSupplier{
		Supplier: s,
		Time:     time.Now(),
	}
}
