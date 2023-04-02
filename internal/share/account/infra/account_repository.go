package repository

import (
	"context"

	"github.com/mrokoo/goERP/internal/share/account/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	accounts *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collection string) *MongoRepository {
	accounts := db.Collection(collection)
	return &MongoRepository{
		accounts: accounts,
	}
}

func (r *MongoRepository) GetAll() ([]*domain.Account, error) {
	cursor, err := r.accounts.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*domain.Account
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *MongoRepository) GetByID(accountID string) (*domain.Account, error) {
	filter := bson.D{{Key: "id", Value: accountID}}
	var account domain.Account
	err := r.accounts.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *MongoRepository) Save(account *domain.Account) error {
	_, err := r.accounts.InsertOne(context.Background(), account)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Replace(account *domain.Account) error {
	filter := bson.D{{Key: "id", Value: account.ID}}
	err := r.accounts.FindOneAndReplace(context.Background(), filter, account).Decode(nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Delete(accountID string) error {
	filter := bson.D{{Key: "id", Value: accountID}}
	err := r.accounts.FindOneAndDelete(context.Background(), filter).Decode(nil)
	if err != nil {
		return err
	}
	return nil
}
