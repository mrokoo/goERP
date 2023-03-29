package infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mrokoo/goERP/internal/system/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	users *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	users := client.Database("goERP").Collection("users")
	return &MongoUserRepository{
		users: users,
	}
}

func (r *MongoUserRepository) Create(user *domain.User) error {
	_, err := r.users.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("fail to persist user: %w", err)
	}
	return nil
}

func (r *MongoUserRepository) Save(user *domain.User) error {
	filter := bson.D{{"id", user.ID}}
	_, err := r.users.ReplaceOne(context.Background(), filter, user)
	if err != nil {
		return fmt.Errorf("fail to persist user: %w", err)
	}
	return nil
}

func (r *MongoUserRepository) Get(userId *uuid.UUID) (*domain.User, error) {
	filter := bson.D{{"id", userId}}
	user := &domain.User{}
	err := r.users.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("fail to find user: %w", err)
	}
	return user, nil
}

func (r *MongoUserRepository) Delete(userId *uuid.UUID) error {
	filter := bson.D{{"id", userId}}
	_, err := r.users.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("fail to delete user: %w", err)
	}
	return nil
}

func (r *MongoUserRepository) GetAll() ([]domain.User, error) {
	cursor, err := r.users.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("fail to get users: %w", err)
	}
	var results []domain.User
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("fail to get users: %w", err)
	}
	return results, nil
}
