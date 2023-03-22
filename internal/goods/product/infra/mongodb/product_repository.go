package infra

import (
	"context"
	"fmt"

	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	products *mongo.Collection
}

func NewMongoProductRepository(client *mongo.Client) *MongoProductRepository {
	products := client.Database("goERP").Collection("products")
	return &MongoProductRepository{
		products: products,
	}
}

func (r *MongoProductRepository) Create(product *domain.Product) error {
	_, err := r.products.InsertOne(context.Background(), product)
	if err != nil {
		return fmt.Errorf("fail to persist product: %w", err)
	}
	return nil
}

func (r *MongoProductRepository) Save(product *domain.Product) error {
	filter := bson.D{{"id", product.ID}}
	_, err := r.products.ReplaceOne(context.Background(), filter, product)
	if err != nil {
		return fmt.Errorf("fail to persist product: %w", err)
	}
	return nil
}

func (r *MongoProductRepository) Get(productId string) (*domain.Product, error) {
	filter := bson.D{{"id", productId}}
	var product *domain.Product
	err := r.products.FindOne(context.Background(), filter).Decode(product)
	if err != nil {
		return nil, fmt.Errorf("fail to find product: %w", err)
	}
	return product, nil
}

func (r *MongoProductRepository) Delete(productId string) error {
	filter := bson.D{{"id", productId}}
	_, err := r.products.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("fail to delete product: %w", err)
	}
	return nil
}

func (r *MongoProductRepository) GetAll() ([]domain.Product, error) {
	cursor, err := r.products.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("fail to get products: %w", err)
	}
	var results []domain.Product
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("fail to get products: %w", err)
	}

	for _, result := range results {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("fail to get products: %w", err)
		}
	}
	return results, nil
}
