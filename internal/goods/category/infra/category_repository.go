package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/goods/category/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	categorys *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	categorys := db.Collection(collection)
	return &MongoRepository{
		categorys: categorys,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Category, error) {
	cursor, err := r.categorys.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Category
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(categoryID uuid.UUID) (*domain.Category, error) {
	filter := bson.D{{Key: "id", Value: categoryID}}
	var category domain.Category
	err := r.categorys.FindOne(context.Background(), filter).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *MongoRepository) Save(category *domain.Category) error {
	_, err := r.categorys.InsertOne(context.Background(), category)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(category *domain.Category) error {
	filter := bson.D{{Key: "id", Value: category.ID}}
	var b domain.Category
	err := r.categorys.FindOneAndReplace(context.Background(), filter, category).Decode(&b)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(categoryID uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: categoryID}}
	var b domain.Category
	err := r.categorys.FindOneAndDelete(context.Background(), filter).Decode(&b)
	if err != nil {
		return err
	}
	return nil
}
