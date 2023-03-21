package infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/unit/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUnitRepository struct {
	units *mongo.Collection
}

func NewMongoUnitRepository(client *mongo.Client) *MongoUnitRepository {
	units := client.Database("goERP").Collection("units")
	return &MongoUnitRepository{
		units: units,
	}
}

func (r *MongoUnitRepository) Create(unit *domain.Unit) error {
	_, err := r.units.InsertOne(context.Background(), unit)
	if err != nil {
		return fmt.Errorf("fail to persist unit: %w", err)
	}
	return nil
}

func (r *MongoUnitRepository) Save(unit *domain.Unit) error {
	filter := bson.D{{"id", unit.ID}}
	_, err := r.units.ReplaceOne(context.Background(), filter, unit)
	if err != nil {
		return fmt.Errorf("fail to persist unit: %w", err)
	}
	return nil
}

func (r *MongoUnitRepository) Get(unitId *uuid.UUID) (*domain.Unit, error) {
	filter := bson.D{{"id", unitId}}
	var unit *domain.Unit
	err := r.units.FindOne(context.Background(), filter).Decode(unit)
	if err != nil {
		return nil, fmt.Errorf("fail to find unit: %w", err)
	}
	return unit, nil
}

func (r *MongoUnitRepository) Delete(unitId *uuid.UUID) error {
	filter := bson.D{{"id", unitId}}
	_, err := r.units.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("fail to delete unit: %w", err)
	}
	return nil
}

func (r *MongoUnitRepository) GetAll() ([]domain.Unit, error) {
	cursor, err := r.units.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("fail to get units: %w", err)
	}
	var results []domain.Unit
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("fail to get units: %w", err)
	}

	for _, result := range results {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("fail to get units: %w", err)
		}
	}
	return results, nil
}
