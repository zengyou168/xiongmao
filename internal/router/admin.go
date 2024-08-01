// Package router 管理员路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/handler"
)

func admin(app *fiber.App) {

	adminRouter := app.Group("admin")
	{
		adminRouter.Post("login", handler.AdminLogin) // 管理员登录
		adminRouter.Post("add", handler.AdminAdd)     // 管理员添加
	}

}
