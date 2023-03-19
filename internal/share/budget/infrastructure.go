package budget

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	SaveBudget(ctx context.Context, budget Budget) error
	DeleteBudget(ctx context.Context, budgetId uuid.UUID) error
	FetchAllBudget(ctx context.Context) ([]Budget, error)
	ChangeBudget(ctx context.Context, budget Budget) error
	LoadBudget(ctx context.Context, budgetId uuid.UUID) (Budget, error)
}

type MongoRepo struct {
	budgets *mongo.Collection
}

type MongoBudget struct {
	ID   uuid.UUID
	Type BudgetType
	Note string
	Time time.Time
}

func (mr *MongoRepo) SaveBudget(ctx context.Context, budget Budget) error {
	mongoB := toMongoBudget(budget)
	_, err := mr.budgets.InsertOne(ctx, mongoB)
	if err != nil {
		return fmt.Errorf("fail to persist budget: %w", err)
	}
	return nil
}

func (mr *MongoRepo) DeleteBudget(ctx context.Context, budgetId uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: budgetId}}
	result, err := mr.budgets.DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the budget")
	}
	if err != nil {
		return fmt.Errorf("fail to delete budget: %w", err)
	}
	return nil
}

func (mr *MongoRepo) ChangeBudget(ctx context.Context, budget Budget) error {
	mongoB := toMongoBudget(budget)
	filter := bson.D{{Key: "id", Value: budget.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "type", Value: mongoB.Type},
		{Key: "note", Value: mongoB.Note},
	}}}
	result, err := mr.budgets.UpdateOne(ctx, filter, update)
	if result.MatchedCount == 0 {
		return errors.New("fail to query the budget")
	}
	if err != nil {
		return fmt.Errorf("fail to change budget: %w", err)
	}
	return nil
}

func (mr *MongoRepo) FetchAllBudget(ctx context.Context) ([]Budget, error) {
	cursor, err := mr.budgets.Find(ctx, bson.D{})

	if err != nil {
		return []Budget{}, fmt.Errorf("fail to fetch budgets: %w", err)
	}

	var results []Budget
	if err = cursor.All(ctx, &results); err != nil {
		return []Budget{}, fmt.Errorf("fail to fetch budgets: %w", err)
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

func (mr *MongoRepo) LoadBudget(ctx context.Context, budgetId uuid.UUID) (Budget, error) {
	var budget Budget
	filter := bson.D{{Key: "id", Value: budgetId}}
	if err := mr.budgets.FindOne(ctx, filter).Decode(&budget); err != nil {
		return Budget{}, err
	}
	return budget, nil
}

func toMongoBudget(b Budget) MongoBudget {
	return MongoBudget{
		ID:   b.ID,
		Type: b.Type,
		Note: b.Note,
		Time: time.Now(),
	}
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	budgets := client.Database("goERP").Collection("budgets")
	return &MongoRepo{
		budgets: budgets,
	}, nil
}
