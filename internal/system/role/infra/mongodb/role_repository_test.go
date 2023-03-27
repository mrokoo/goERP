package infra

import (
	"context"
	"testing"

	"github.com/mrokoo/goERP/internal/system/role/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var r *MongoRepository

func Test_MongoRepository_Save(t *testing.T) {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r = NewMongoRepository(client)
	role := &domain.Role{
		Name:       "admin",
		Permission: []domain.PermissionItem{"supplier", "customer"},
	}

	err = r.Save(role)
	if err != nil {
		t.Error(err)
	}
}

func Test_MongoRepository_Delete(t *testing.T) {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r = NewMongoRepository(client)
	err = r.Delete("admin")
	if err != nil {
		t.Error(err)
	}
}
