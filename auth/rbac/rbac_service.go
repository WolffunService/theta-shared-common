package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
var m, _ = model.NewModelFromString(text)

type AuthorizationService struct {
	e *casbin.Enforcer
}

var Service AuthorizationService

func NewService(db *mongo.Database, dbName string) (error, *AuthorizationService) {
	if len(dbName) == 0 {
		dbName = defaultCollectionName
	}

	a, err := NewAdapterWithDB(db, dbName)
	// Or you can use NewAdapterWithCollectionName for custom collection name.
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return err, nil
	}
	e.AddRoleForUser("admin", "3rd_default")

	return nil, &AuthorizationService{e: e}
}

func (s *AuthorizationService) LoadPolicy() error {
	return s.e.LoadPolicy()
}

func (s *AuthorizationService) GetEnforce() *casbin.Enforcer {
	return s.e
}

func (s *AuthorizationService) Enforce(subject, object, action string) (bool, error) {
	return s.e.Enforce(subject, object, action)
}
