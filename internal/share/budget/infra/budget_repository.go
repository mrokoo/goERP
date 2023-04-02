package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	budgets *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	budgets := db.Collection(collection)
	return &MongoRepository{
		budgets: budgets,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Budget, error) {
	cursor, err := r.budgets.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Budget
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(budgetID uuid.UUID) (*domain.Budget, error) {
	filter := bson.D{{Key: "id", Value: budgetID}}
	var budget domain.Budget
	err := r.budgets.FindOne(context.Background(), filter).Decode(&budget)
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *MongoRepository) Save(budget *domain.Budget) error {
	_, err := r.budgets.InsertOne(context.Background(), budget)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(budget *domain.Budget) error {
	filter := bson.D{{Key: "id", Value: budget.ID}}
	err := r.budgets.FindOneAndReplace(context.Background(), filter, budget).Decode(nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(budgetID uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: budgetID}}
	err := r.budgets.FindOneAndDelete(context.Background(), filter).Decode(nil)
	if err != nil {
		return err
	}
	return nil
}
