// Package router 路由
package router

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"panda/config"
	"panda/pkg/db"
	"panda/pkg/respond"
	"sort"
	"strconv"
	"strings"
)

var jwtKey = []byte("your_secret_key")
var enforcer *casbin.Enforcer

func Init() {

	app := fiber.New(fiber.Config{
		ErrorHandler: respond.ErrorHandler, // 设置全局错误处理函数
	})

	// 初始化 Casbin
	adapter, _ := gormadapter.NewAdapterByDB(db.Gorm)
	enforcer, _ = casbin.NewEnforcer("pkg/casbin/model.conf", adapter)
	_ = enforcer.LoadPolicy()

	// JWT身份验证中间件
	app.Use(func(c *fiber.Ctx) error {

		// 全局捕获信息并返回，注意：在此后执行的方法才会捕获，如果有此之前执行的方法，想捕获时，把此方法放在要捕获的方法之前，或者复制一份过去
		defer func() {
			if r := recover(); r != nil {
				_ = c.JSON(r)
			}
		}()

		url := c.OriginalURL()
		avoidLogin := []string{"/user/login"} // 免登录

		// 排序数组
		sort.Strings(avoidLogin)

		// 使用二分查找
		search := sort.SearchStrings(avoidLogin, url)

		if search < len(avoidLogin) && avoidLogin[search] == url {
			return c.Next()
		}

		tokenStr := c.Get("Authorization")

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			return respond.ErrorCode(respond.TokenExpire, respond.TokenExpireMsg)
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return respond.ErrorCode(respond.TokenExpire, respond.TokenExpireMsg)
		}

		c.Locals("user", claims)

		return c.Next()
	})

	user(app)

	err := app.Listen(":" + strconv.Itoa(config.ServerVar.Port))

	if err != nil {
		panic(err)
	}
}
