// Package router 路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"sort"
	"strconv"
	"strings"
	"xiongmao/config"
	"xiongmao/pkg/casbin"
	"xiongmao/pkg/respond"
)

func Init() {

	app := fiber.New(fiber.Config{
		ErrorHandler: respond.ErrorHandler, // 设置全局错误处理函数
	})

	// JWT身份验证中间件
	app.Use(func(c *fiber.Ctx) error {

		// 全局捕获信息并返回，注意：在此后执行的方法才会捕获，如果有此之前执行的方法，想捕获时，把此方法放在要捕获的方法之前，或者复制一份过去
		defer func() {
			if r := recover(); r != nil {
				_ = c.JSON(r)
			}
		}()

		url := c.OriginalURL()
		avoidLogin := []string{"/admin/login"} // 免登录

		// 排序数组
		sort.Strings(avoidLogin)

		// 使用二分查找
		search := sort.SearchStrings(avoidLogin, url)

		if search < len(avoidLogin) && avoidLogin[search] == url {
			return c.Next()
		}

		tokenStr := c.Get("Authorization")

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			return respond.ErrorCode(respond.TokenExpire, tokenStr)
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return config.CasbinSecretKey, nil
		})

		if err != nil || !token.Valid {
			return respond.ErrorCode(respond.TokenExpire, respond.TokenExpireMsg)
		}

		casbin.Init()

		c.Locals("user", claims)

		return c.Next()

		/*	// 检查权限
			sub := "alice"  // 用户
			obj := "/data1" // 路由
			act := "GET"    // 操作

			res, err := e.Enforce(sub, obj, act)

			if err != nil {
				log.SugarVar.Error("强制执行错误", err)
			}

			if res {
				log.SugarVar.Error("授予访问权限")
			} else {
				log.SugarVar.Error("访问被拒绝")
			}*/
	})

	admin(app)

	err := app.Listen(":" + strconv.Itoa(config.ServerVar.Port))

	if err != nil {
		panic(err)
	}
}
