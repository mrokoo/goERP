package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrokoo/goERP/internal/share/warehouse/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrNotUID = errors.New("the warehouseId is not unique")

type MongoRepository struct {
	warehouses *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	warehouses := client.Database("goERP").Collection("warehouses")
	return &MongoRepository{
		warehouses: warehouses,
	}
}

// 插入新文档
func (r *MongoRepository) Save(warehouse domain.Warehouse) error {
	_, err := r.warehouses.InsertOne(context.Background(), warehouse)
	if err != nil {
		return fmt.Errorf("fail to persist warehouse: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(warehouse domain.Warehouse) error {
	filter := bson.D{{Key: "id", Value: warehouse.ID}}
	_, err := r.warehouses.ReplaceOne(context.Background(), filter, warehouse)
	if err != nil {
		return fmt.Errorf("fail to change warehouse: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(warehouseId domain.WarehouseId) (*domain.Warehouse, error) {
	filter := bson.D{{Key: "id", Value: warehouseId}}
	var warehouse domain.Warehouse
	err := r.warehouses.FindOne(context.Background(), filter).Decode(&warehouse)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.Warehouse, error) {
	cursor, err := r.warehouses.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.Warehouse{}, fmt.Errorf("fail to fetch warehouses: %w", err)
	}

	var results []domain.Warehouse
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.Warehouse{}, fmt.Errorf("fail to fetch warehouses: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(warehouseID domain.WarehouseId) error {
	filter := bson.D{{Key: "id", Value: warehouseID}}
	result, err := r.warehouses.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the warehouse")
	}
	if err != nil {
		return fmt.Errorf("fail to delete warehouse: %w", err)
	}
	return nil
}
