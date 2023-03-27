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
	_, err := r.e.RemoveFilteredPolicy(0, name)
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

func (r *MongoRepository) FindaAll() ([]*domain.Role, error) {
	var allrole []*domain.Role
	allSubjects := r.e.GetAllSubjects()
	for _, v := range allSubjects {
		role, err := r.FindByName(v)
		if err != nil {
			return nil, err
		}
		allrole = append(allrole, role)
	}

	return allrole, nil
}

func toRole(p [][]string) []domain.PermissionItem {
	var pp []domain.PermissionItem
	for _, v := range p {
		pp = append(pp, v[1])
	}
	return removeRepByMap(pp)
}

// slice去重
func removeRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}
