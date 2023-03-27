package infra

import (
	"errors"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/mrokoo/goERP/internal/system/role/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrFailUpdate = errors.New("fail to update role")

type MongoRepository struct {
	e *casbin.Enforcer
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	a, err := mongodbadapter.NewAdapterByDB(client, &mongodbadapter.AdapterConfig{
		DatabaseName:   "goERP",
		CollectionName: "roles",
	})
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("../../../../../config/model.conf", a)
	if err != nil {
		panic(err)
	}
	return &MongoRepository{
		e: e,
	}
}

func (r *MongoRepository) Save(role *domain.Role) error {
	for _, v := range role.Permission {
		rules := [][]string{
			{role.Name, v, "read"},
			{role.Name, v, "write"},
		}
		_, err := r.e.AddPolicies(rules)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MongoRepository) Update(role *domain.Role) error {
	if err := r.Delete(role.Name); err != nil {
		return ErrFailUpdate
	}
	if err := r.Save(role); err != nil {
		return ErrFailUpdate
	}
	return nil
}

func (r *MongoRepository) Delete(name string) error {
	ok, err := r.e.RemoveFilteredGroupingPolicy(0, name)
	println(ok)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) FindByName(name string) (*domain.Role, error) {
	filteredPolicy := r.e.GetFilteredPolicy(0, name)
	pp := toRole(filteredPolicy)
	return &domain.Role{
		Name:       name,
		Permission: pp,
	}, nil
}

// func (r *MongoRepository) FindaAll() ([]*domain.Role, error) {

// }

func toRole(p [][]string) []domain.PermissionItem {
	var pp []domain.PermissionItem
	for _, v := range p {
		pp = append(pp, v[1])
	}
	return pp
}
