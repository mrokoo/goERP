package infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCategoryRepository struct {
	categories *mongo.Collection
}

func NewMongoCategoryRepository(client *mongo.Client) *MongoCategoryRepository {
	categories := client.Database("goERP").Collection("categories")
	return &MongoCategoryRepository{
		categories: categories,
	}
}

func (r *MongoCategoryRepository) Create(category *domain.Category) error {
	_, err := r.categories.InsertOne(context.Background(), category)
	if err != nil {
		return fmt.Errorf("fail to persist category: %w", err)
	}
	return nil
}

func (r *MongoCategoryRepository) Save(category *domain.Category) error {
	filter := bson.D{{"id", category.ID}}
	_, err := r.categories.ReplaceOne(context.Background(), filter, category)
	if err != nil {
		return fmt.Errorf("fail to persist category: %w", err)
	}
	return nil
}

func (r *MongoCategoryRepository) Get(categoryId *uuid.UUID) (*domain.Category, error) {
	filter := bson.D{{"id", categoryId}}
	var category *domain.Category
	err := r.categories.FindOne(context.Background(), filter).Decode(category)
	if err != nil {
		return nil, fmt.Errorf("fail to find category: %w", err)
	}
	return category, nil
}

func (r *MongoCategoryRepository) Delete(categoryId *uuid.UUID) error {
	filter := bson.D{{"id", categoryId}}
	_, err := r.categories.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("fail to delete category: %w", err)
	}
	return nil
}

func (r *MongoCategoryRepository) GetAll() ([]domain.Category, error) {
	cursor, err := r.categories.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("fail to get categories: %w", err)
	}
	var results []domain.Category
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("fail to get categories: %w", err)
	}

	for _, result := range results {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("fail to get categories: %w", err)
		}
	}
	return results, nil
}
