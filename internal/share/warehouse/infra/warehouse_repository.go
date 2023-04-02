package repository

import (
	"context"

	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	warehouses *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	warehouses := db.Collection(collection)
	return &MongoRepository{
		warehouses: warehouses,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Warehouse, error) {
	cursor, err := r.warehouses.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Warehouse
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(warehouseID string) (*domain.Warehouse, error) {
	filter := bson.D{{Key: "id", Value: warehouseID}}
	var warehouse domain.Warehouse
	err := r.warehouses.FindOne(context.Background(), filter).Decode(&warehouse)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *MongoRepository) Save(warehouse *domain.Warehouse) error {
	_, err := r.warehouses.InsertOne(context.Background(), warehouse)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(warehouse *domain.Warehouse) error {
	filter := bson.D{{Key: "id", Value: warehouse.ID}}
	var w domain.Warehouse
	err := r.warehouses.FindOneAndReplace(context.Background(), filter, warehouse).Decode(&w)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(warehouseID string) error {
	filter := bson.D{{Key: "id", Value: warehouseID}}
	var w domain.Warehouse
	err := r.warehouses.FindOneAndDelete(context.Background(), filter).Decode(&w)
	if err != nil {
		return err
	}
	return nil
}
