package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	users *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	users := client.Database("goERP").Collection("users")
	return &MongoRepository{
		users: users,
	}
}

// 插入新文档
func (r *MongoRepository) Save(user domain.User) error {
	_, err := r.users.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("fail to persist user: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(user domain.User) error {
	filter := bson.D{{Key: "id", Value: user.ID}}
	_, err := r.users.ReplaceOne(context.Background(), filter, user)
	if err != nil {
		return fmt.Errorf("fail to change user: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(userId uuid.UUID) (*domain.User, error) {
	filter := bson.D{{Key: "id", Value: userId}}
	var user domain.User
	err := r.users.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.User, error) {
	cursor, err := r.users.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.User{}, fmt.Errorf("fail to fetch users: %w", err)
	}

	var results []domain.User
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.User{}, fmt.Errorf("fail to fetch users: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(userID uuid.UUID) error {
	filter := bson.D{{Key: "id", Value: userID}}
	result, err := r.users.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the user")
	}
	if err != nil {
		return fmt.Errorf("fail to delete user: %w", err)
	}
	return nil
}
