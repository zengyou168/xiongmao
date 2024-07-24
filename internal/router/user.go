// Package router 用户路由
package router

import (
	"github.com/gofiber/fiber/v2"
)

func user(app *fiber.App) {

	userRouter := app.Group("user")

	{
		userRouter.Post("create", api.Add) // 用户登录
	}
}
