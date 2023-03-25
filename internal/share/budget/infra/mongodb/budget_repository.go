package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/share/budget/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	budgets *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	budgets := client.Database("goERP").Collection("budgets")
	return &MongoRepository{
		budgets: budgets,
	}
}

// 插入新文档
func (r *MongoRepository) Save(budget domain.Budget) error {
	_, err := r.budgets.InsertOne(context.Background(), budget)
	if err != nil {
		return fmt.Errorf("fail to persist budget: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(budget domain.Budget) error {
	filter := bson.D{{Key: "id", Value: budget.ID}}
	_, err := r.budgets.ReplaceOne(context.Background(), filter, budget)
	if err != nil {
		return fmt.Errorf("fail to change budget: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(budgetId uuid.UUID) (*domain.Budget, error) {
	filter := bson.D{{Key: "id", Value: budgetId}}
	var budget domain.Budget
	err := r.budgets.FindOne(context.Background(), filter).Decode(&budget)
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.Budget, error) {
	cursor, err := r.budgets.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.Budget{}, fmt.Errorf("fail to fetch budgets: %w", err)
	}

	var results []domain.Budget
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.Budget{}, fmt.Errorf("fail to fetch budgets: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(budgetID uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: budgetID}}
	result, err := r.budgets.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the budget")
	}
	if err != nil {
		return fmt.Errorf("fail to delete budget: %w", err)
	}
	return nil
}
