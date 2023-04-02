package repository

import (
	"context"

	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	suppliers *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	suppliers := db.Collection(collection)
	return &MongoRepository{
		suppliers: suppliers,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Supplier, error) {
	cursor, err := r.suppliers.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Supplier
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(supplierID string) (*domain.Supplier, error) {
	filter := bson.D{{Key: "id", Value: supplierID}}
	var supplier domain.Supplier
	err := r.suppliers.FindOne(context.Background(), filter).Decode(&supplier)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *MongoRepository) Save(supplier *domain.Supplier) error {
	_, err := r.suppliers.InsertOne(context.Background(), supplier)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(supplier *domain.Supplier) error {
	filter := bson.D{{Key: "id", Value: supplier.ID}}
	var s domain.Supplier
	err := r.suppliers.FindOneAndReplace(context.Background(), filter, supplier).Decode(&s)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(supplierID string) error {
	filter := bson.D{{Key: "id", Value: supplierID}}
	var s domain.Supplier
	err := r.suppliers.FindOneAndDelete(context.Background(), filter).Decode(&s)
	if err != nil {
		return err
	}
	return nil
}
