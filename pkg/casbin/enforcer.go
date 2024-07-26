package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"panda/pkg/db"
)

var jwtKey = []byte("your_secret_key")
var enforcer *casbin.Enforcer

func main() {

	// Initialize Casbin
	adapter, _ := gormadapter.NewAdapterByDB(db.Gorm)
	enforcer, _ = casbin.NewEnforcer("model.conf", adapter)
	enforcer.LoadPolicy()

}
