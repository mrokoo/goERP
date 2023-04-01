package infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/purchase/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	purchaseOrders *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	purchaseOrders := client.Database("goERP").Collection("purchaseOrders")
	return &MongoRepository{
		purchaseOrders: purchaseOrders,
	}
}

// 插入新文档
func (r *MongoRepository) Save(purchaseOrder domain.PurchaseOrder) error {
	_, err := r.purchaseOrders.InsertOne(context.Background(), purchaseOrder)
	if err != nil {
		return fmt.Errorf("fail to persist purchaseOrder: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(purchaseOrderId uuid.UUID) (*domain.PurchaseOrder, error) {
	filter := bson.D{{Key: "id", Value: purchaseOrderId}}
	var purchaseOrder domain.PurchaseOrder
	err := r.purchaseOrders.FindOne(context.Background(), filter).Decode(&purchaseOrder)
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.PurchaseOrder, error) {
	cursor, err := r.purchaseOrders.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.PurchaseOrder{}, fmt.Errorf("fail to fetch purchaseOrders: %w", err)
	}

	var results []domain.PurchaseOrder
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.PurchaseOrder{}, fmt.Errorf("fail to fetch purchaseOrders: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Invalidated(purchaseOrderID uuid.UUID) error {
	// to do it
	return nil
}
