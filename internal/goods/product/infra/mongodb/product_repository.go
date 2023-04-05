package repository

import (
	"context"

	"github.com/mrokoo/goERP/internal/goods/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	products *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	products := db.Collection(collection)
	return &MongoRepository{
		products: products,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Product, error) {
	cursor, err := r.products.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Product
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(productID string) (*domain.Product, error) {
	filter := bson.D{{Key: "id", Value: productID}}
	var product domain.Product
	err := r.products.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MongoRepository) Save(product *domain.Product) error {
	_, err := r.products.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(product *domain.Product) error {
	filter := bson.D{{Key: "id", Value: product.ID}}
	var c domain.Product
	err := r.products.FindOneAndReplace(context.Background(), filter, product).Decode(&c)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(productID string) error {
	filter := bson.D{{Key: "id", Value: productID}}
	var c domain.Product
	err := r.products.FindOneAndDelete(context.Background(), filter).Decode(&c)
	if err != nil {
		return err
	}
	return nil
}
