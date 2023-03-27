package infra

import (
	"context"
	"testing"

	"github.com/mrokoo/goERP/internal/system/role/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var r *MongoRepository

func TestMongoRepository_Save(t *testing.T) {
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

func TestMongoRepository_Update(t *testing.T) {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r = NewMongoRepository(client)
	role := &domain.Role{
		Name:       "admin",
		Permission: []domain.PermissionItem{"account", "customer"},
	}
	err = r.Update(role)
	if err != nil {
		t.Error(err)
	}
}

func TestMongoRepository_FindByName(t *testing.T) {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r = NewMongoRepository(client)
	role, err := r.FindByName("admin")
	if err != nil {
		t.Error(err)
	}
	role2 := &domain.Role{
		Name:       "admin",
		Permission: []domain.PermissionItem{"account", "customer"},
	}
	assert.Equal(t, *role2, *role)
}

func TestMongoRepository_FindAll(t *testing.T) {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r = NewMongoRepository(client)
	role2 := &domain.Role{
		Name:       "admin",
		Permission: []domain.PermissionItem{"account", "customer"},
	}
	result, err := r.FindaAll()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, *role2, *result[0])
}

func TestMongoRepository_Delete(t *testing.T) {
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
