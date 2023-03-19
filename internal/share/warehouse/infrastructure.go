package warehouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mrokoo/goERP/internal/share/valueobj"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the warehouseId is not unique")

type Repository interface {
	SaveWarehouse(ctx context.Context, warehouse Warehouse) error
	DeleteWarehouse(ctx context.Context, warehouseId WarehouseId) error
	FetchAllWarehouse(ctx context.Context) ([]Warehouse, error)
	ChangeWarehouse(ctx context.Context, warehouse Warehouse) error
}

type MongoRepo struct {
	warehouses *mongo.Collection
}

type MongoWarehouse struct {
	ID      WarehouseId
	Name    valueobj.Name
	Admin   string
	Phone   valueobj.Phone
	Address valueobj.Address
	Note    string
	State   valueobj.StateType
	Time    time.Time
}

func (mr *MongoRepo) SaveWarehouse(ctx context.Context, warehouse Warehouse) error {
	mongoW := toMongoWarehouse(warehouse)
	_, err := mr.warehouses.InsertOne(ctx, mongoW)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("fail to persist warehouse: %w", ErrNotUID)
		}
		return fmt.Errorf("fail to persist warehouse: %w", err)
	}
	return nil
}

func (mr *MongoRepo) DeleteWarehouse(ctx context.Context, warehouseId WarehouseId) error {
	filter := bson.D{{Key: "id", Value: warehouseId}}
	result, err := mr.warehouses.DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the supplier")
	}
	if err != nil {
		return fmt.Errorf("fail to delete supplier: %w", err)
	}
	return nil
}

func (mr *MongoRepo) ChangeWarehouse(ctx context.Context, warehouse Warehouse) error {
	mongoW := toMongoWarehouse(warehouse)
	filter := bson.D{{Key: "id", Value: warehouse.ID}}
	result, err := mr.warehouses.ReplaceOne(ctx, filter, mongoW)
	if result.MatchedCount == 0 {
		return errors.New("fail to query the warehouse")
	}
	if err != nil {
		return fmt.Errorf("fail to change warehouse: %w", err)
	}
	return nil
}

func (mr *MongoRepo) FetchAllWarehouse(ctx context.Context) ([]Warehouse, error) {
	cursor, err := mr.warehouses.Find(ctx, bson.D{})

	if err != nil {
		return []Warehouse{}, fmt.Errorf("fail to fetch warehouses: %w", err)
	}

	var results []Warehouse
	if err = cursor.All(ctx, &results); err != nil {
		return []Warehouse{}, fmt.Errorf("fail to fetch warehouses: %w", err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		_, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	}
	return results, nil
}

func toMongoWarehouse(w Warehouse) MongoWarehouse {
	return MongoWarehouse{
		ID:      w.ID,
		Name:    w.Name,
		Admin:   w.Admin,
		Phone:   w.Phone,
		Address: w.Address,
		Note:    w.Note,
		State:   w.State,
		Time:    time.Now(),
	}
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	warehouses := client.Database("goERP").Collection("warehouses")
	return &MongoRepo{
		warehouses: warehouses,
	}, nil
}
