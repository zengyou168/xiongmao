package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"xiongmao/pkg/db"
)

var jwtKey = []byte("your_secret_key")
var enforcer *casbin.Enforcer

func Init() {

	// 初始化 Casbin
	adapter, _ := gormadapter.NewAdapterByDB(db.Gorm)
	enforcer, _ = casbin.NewEnforcer("model.conf", adapter)
	_ = enforcer.LoadPolicy()

}
