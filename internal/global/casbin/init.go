package casbin

import (
	"fmt"
	"gin/configs"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 定义casbin的model
const (
	casModel = `
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
)

var Enforce *MyEnforce

// 初始化casbin
func Init() {
	var err error
	Enforce, err = createEnforcer()
	if err != nil {
		return
	}
	err1 := Enforce.LoadPolicy()
	if err1 != nil {
		return
	}
	Enforce.EnableAutoSave(true)
	Enforce.AddBasedPolicies()
}

// 创建casbin的enforcer
func createEnforcer() (*MyEnforce, error) {
	m, err := model.NewModelFromString(casModel)
	if err != nil {
		return nil, err
	}
	//使用config文件里面的参数创建适配器
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		configs.DbSettings.Root,
		configs.DbSettings.Password,
		configs.DbSettings.Host,
		configs.DbSettings.Port,
		configs.DbSettings.Dbname,
	)
	adapter, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	return &MyEnforce{e}, nil
}
