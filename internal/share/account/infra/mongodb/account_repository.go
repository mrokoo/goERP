package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrokoo/goERP/internal/share/account/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	accounts *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	accounts := client.Database("goERP").Collection("suppliers")
	return &MongoRepository{
		accounts: accounts,
	}
}

// 插入新文档
func (r *MongoRepository) Save(account domain.Account) error {
	_, err := r.accounts.InsertOne(context.Background(), account)
	if err != nil {
		return fmt.Errorf("fail to persist account: %w", err)
	}
	return nil
}

// 更新替换匹配文档
func (r *MongoRepository) Update(account domain.Account) error {
	filter := bson.D{{Key: "id", Value: account.ID}}
	_, err := r.accounts.ReplaceOne(context.Background(), filter, account)
	if err != nil {
		return fmt.Errorf("fail to change account: %w", err)
	}
	return nil
}

// 根据ID获取文档
func (r *MongoRepository) Get(accountId domain.AccountId) (*domain.Account, error) {
	filter := bson.D{{Key: "id", Value: accountId}}
	var account domain.Account
	err := r.accounts.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// 获取全部文档
func (r *MongoRepository) GetAll() ([]domain.Account, error) {
	cursor, err := r.accounts.Find(context.Background(), bson.D{})

	if err != nil {
		return []domain.Account{}, fmt.Errorf("fail to fetch accounts: %w", err)
	}

	var results []domain.Account
	if err = cursor.All(context.Background(), &results); err != nil {
		return []domain.Account{}, fmt.Errorf("fail to fetch accounts: %w", err)
	}
	return results, nil
}

// 根据ID删除文档
func (r *MongoRepository) Delete(accountID domain.AccountId) error {
	filter := bson.D{{Key: "id", Value: accountID}}
	result, err := r.accounts.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the account")
	}
	if err != nil {
		return fmt.Errorf("fail to delete account: %w", err)
	}
	return nil
}
