package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	units *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	units := db.Collection(collection)
	return &MongoRepository{
		units: units,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Unit, error) {
	cursor, err := r.units.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Unit
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(unitID uuid.UUID) (*domain.Unit, error) {
	filter := bson.D{{Key: "id", Value: unitID}}
	var unit domain.Unit
	err := r.units.FindOne(context.Background(), filter).Decode(&unit)
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

func (r *MongoRepository) Save(unit *domain.Unit) error {
	_, err := r.units.InsertOne(context.Background(), unit)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(unit *domain.Unit) error {
	filter := bson.D{{Key: "id", Value: unit.ID}}
	var b domain.Unit
	err := r.units.FindOneAndReplace(context.Background(), filter, unit).Decode(&b)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(unitID uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: unitID}}
	var b domain.Unit
	err := r.units.FindOneAndDelete(context.Background(), filter).Decode(&b)
	if err != nil {
		return err
	}
	return nil
}
