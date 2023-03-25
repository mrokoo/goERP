package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrokoo/goERP/internal/share/supplier/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	suppliers *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	suppliers := client.Database("goERP").Collection("suppliers")
	return &MongoRepository{
		suppliers: suppliers,
	}
}

// 插入新文档
func (r *MongoRepository) Save(supplier domain.Supplier) error {
	_, err := r.suppliers.InsertOne(context.Background(), supplier)
	if err != nil {
		return fmt.Errorf("fail to persist supplier: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(supplier domain.Supplier) error {
	filter := bson.D{{Key: "id", Value: supplier.ID}}
	_, err := r.suppliers.ReplaceOne(context.Background(), filter, supplier)
	if err != nil {
		return fmt.Errorf("fail to change supplier: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(supplierId domain.SupplierId) (*domain.Supplier, error) {
	filter := bson.D{{Key: "id", Value: supplierId}}
	var supplier domain.Supplier
	err := r.suppliers.FindOne(context.Background(), filter).Decode(&supplier)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.Supplier, error) {
	cursor, err := r.suppliers.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}

	var results []domain.Supplier
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.Supplier{}, fmt.Errorf("fail to fetch suppliers: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(supplierID domain.SupplierId) error {
	filter := bson.D{{Key: "id", Value: supplierID}}
	result, err := r.suppliers.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the supplier")
	}
	if err != nil {
		return fmt.Errorf("fail to delete supplier: %w", err)
	}
	return nil
}
