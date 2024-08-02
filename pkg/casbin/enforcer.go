package casbin

import (
    "github.com/casbin/casbin/v2"
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "xiongmao/pkg/db"
    "xiongmao/pkg/log"
)

var EnforcerVar *casbin.Enforcer

func Init() {

    a, err := gormadapter.NewAdapterByDB(db.Gorm)

    if err != nil {
        log.SugarVar.Error("casbin适配器初始化失败")
    }

    e, err := casbin.NewEnforcer("pkg/casbin/model.conf", a)

    if err != nil {
        log.SugarVar.Error("casbin执行器初始化失败", err)
    }

    err = e.LoadPolicy()

    if err != nil {
        log.SugarVar.Error("casbin策略加载失败", err)
    }

    EnforcerVar = e
}
