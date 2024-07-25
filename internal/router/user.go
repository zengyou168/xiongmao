// Package router 用户路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/handler"
)

func user(app *fiber.App) {

	userRouter := app.Group("user")

	{
		userRouter.Post("create", handler.Add) // 用户登录
	}
}
