package server

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	Enforcer *casbin.Enforcer
}

func InitServer() *Server {
	dsn := "manager:123qweasd@tcp(127.0.0.1:3306)/market?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("init db fail. err: %s\n", err)
		return nil
	}

	gormadapter.TurnOffAutoMigrate(db)
	// 使用MySQL数据库初始化一个Xorm适配器
	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &gormadapter.CasbinRule{})
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
		return nil
	}

	m, err := model.NewModelFromString(`
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
`)

	if err != nil {
		log.Fatalf("error: model: %s", err)
		return nil
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
		return nil
	}

	// Load the policy from DB.
	e.LoadPolicy()

	result := &Server{Enforcer: e}
	return result
}
