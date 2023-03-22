package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNotUID = errors.New("the accountId is not unique")

type Repository interface {
	Save(ctx context.Context, account Account) error
	Change(ctx context.Context, account Account) error
	FetchAll(ctx context.Context) ([]Account, error)
	Delete(ctx context.Context, accountID AccountId) error
}

type MongoRepo struct {
	accounts *mongo.Collection
}

func (mr *MongoRepo) Save(ctx context.Context, account Account) error {
	ma := toMongoAccount(account)
	_, err := mr.accounts.InsertOne(ctx, ma)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("fail to persist account: %w", ErrNotUID)
		}
		return fmt.Errorf("fail to persist account: %w", err)
	}
	return nil
}

func (mr *MongoRepo) Change(ctx context.Context, account Account) error {
	mongoS := toMongoAccount(account)
	filter := bson.D{{Key: "id", Value: account.ID}}

	result, err := mr.accounts.ReplaceOne(ctx, filter, mongoS)
	if result.MatchedCount == 0 {
		return errors.New("fail to query the account")
	}
	if err != nil {
		return fmt.Errorf("fail to change account: %w", err)
	}
	return nil
}

func (mr *MongoRepo) FetchAll(ctx context.Context) ([]Account, error) {
	cursor, err := mr.accounts.Find(ctx, bson.D{})

	if err != nil {
		return []Account{}, fmt.Errorf("fail to fetch accounts: %w", err)
	}

	var results []Account
	if err = cursor.All(ctx, &results); err != nil {
		return []Account{}, fmt.Errorf("fail to fetch accounts: %w", err)
	}
	return results, nil
}

func (mr *MongoRepo) Delete(ctx context.Context, accountID AccountId) error {
	filter := bson.D{{Key: "id", Value: accountID}}
	result, err := mr.accounts.DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("fail to query the account")
	}
	if err != nil {
		return fmt.Errorf("fail to delete account: %w", err)
	}
	return nil
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	accounts := client.Database("goERP").Collection("suppliers")
	return &MongoRepo{
		accounts: accounts,
	}, nil
}

type MongoAccount struct {
	Account `bson:"inline"`
	Time    time.Time
}

func toMongoAccount(a Account) MongoAccount {
	return MongoAccount{
		Account: a,
		Time:    time.Now(),
	}
}
